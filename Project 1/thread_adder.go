package main

import (
	//Used for Printouts

	"encoding/json"
	"fmt"
	"math"
	"os" //Takes in command line inputs
	"strconv"
	"strings"
	"sync"
)

// Json message to send from coordinator to worker
// datafile: fname, start: pos1 , end: pos2
// For whatever reason the member variables need to be capitalized
type WorkerMessage struct {
	Datafile string
	Start    int
	End      int
}

// From worker to coordinator
type CoordinatorMessage struct {
	Psum   int
	Pcount int
	Prefix string
	Suffix string
	Start  int
	End    int
}

// Channels between workers and coordinator
type ChannelWorker struct {
	Work       chan []byte
	Result     chan []byte
	Wait_Group *sync.WaitGroup
}

// Error checking dummy function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Coordinator
func multi_add(M float64, fname string) {
	wait_group := new(sync.WaitGroup)

	var worker_message WorkerMessage          //Message to send to worker
	var result_array = make([][]byte, int(M)) // Array to hold result arrays from partial_sum

	//Channels to send and recieve information between coordinator and worker
	work := make(chan []byte, int(M))
	result := make(chan []byte, int(M))

	// Reading file numbers into file variable
	number_file, err := os.Open(fname)
	check(err)
	defer number_file.Close()

	// Getting file size = #chars including white space
	file_stats, err := number_file.Stat()
	check(err)
	file_size := file_stats.Size()

	//Split into M fragments
	//May split unevenly because of floor function - extra numbers go to last fragment
	fragment_size := math.Floor(float64(file_size) / M)
	extra_numbers := file_size - int64(fragment_size)*int64(M)

	index := 0 //index to move window by fragment size
	for i := 0; i < int(M); i++ {

		//If i == M, then put these into that fragment else normal window
		if i == int(M)-1 {
			worker_message = WorkerMessage{fname, index, index + int(fragment_size) + int(extra_numbers)}
		} else {
			worker_message = WorkerMessage{fname, index, index + int(fragment_size)}
		}
		index = index + int(fragment_size) //iterating index

		//Sending job and message to worker
		send_message, _ := json.Marshal(worker_message)
		work <- send_message

		//Creating thread with worker
		worker := &ChannelWorker{work, result, wait_group}
		go worker.partial_sum()
		wait_group.Add(1)

		//Return value from partial_sum
		result_array[i] = <-result

		for i := 0; i < len(result_array); i++ {
			fmt.Println(string(result_array[i]))
		}

		//Logic after getting partial sums here
		//If intersection for adjacent prefix and suffix contains space character, then both are whole numbers
		//If first element, it is whole; if last element, it is whole

		//Probably should wait here for all threads
		//Computation of total sum here
	}

	//Wait here
	wait_group.Wait()
}

//	Worker
//
// {psum: 23, pcount: 3, prefix: '1224 ', suffix: ' 678', start:40, end:55}"
func (worker *ChannelWorker) partial_sum() {
	var coordinator_message CoordinatorMessage
	var psum int
	var pcount int
	var prefix string
	var suffix string

	//Get json file containing info
	var my_message WorkerMessage
	err := json.Unmarshal(<-worker.Work, &my_message)
	check(err)

	//Getting individual components
	filename := my_message.Datafile
	start := my_message.Start
	end := my_message.End

	// Reading file numbers into file variable
	number_file, err := os.ReadFile(filename)
	check(err)

	byte_file := []byte(number_file)                               //byte array of file contents
	fragment := string(byte_file[start:end])                       //Getting framgent of file as bytearray
	trimmed_framgent := strings.TrimSpace(fragment)                //Removing trailing whitespace
	fragment_array := strings.Split(string(trimmed_framgent), " ") //Turning string into array

	//Getting prefix suffix pcount and psum with special cases where there are only 1 or 2 numbers
	if len(fragment_array) == 1 {
		prefix = ""
		suffix = ""
		psum, _ = strconv.Atoi(fragment_array[0])
		pcount = 1
	} else if len(fragment_array) == 2 {
		prefix = fragment_array[0]
		suffix = fragment_array[1]
		psum = 0
		pcount = 0
	} else {
		//Getting prefix and suffix
		prefix = fragment_array[0]
		suffix = fragment_array[len(fragment_array)-1]

		//Getting psum
		psum = 0
		pcount = 0
		temp := 0
		for i := 1; i < len(fragment_array)-1; i++ {
			temp, err = strconv.Atoi(fragment_array[i])
			check(err)
			psum = psum + temp
			pcount = pcount + 1
		}
	}

	//Write to a json to return to coordinator
	coordinator_message = CoordinatorMessage{psum, pcount, prefix, suffix, start, end}

	//Return value through channel
	send_message, _ := json.Marshal(coordinator_message)

	<-worker.Wait_Group.Done()
	worker.Result <- send_message
}

// MAIN FUNCTION
//
// {psum: SUM OF ALL NUMBERS FULLY CONTAINED IN SLICE,
// pcount: NUMBER OF NUMBERS FULLY CONTAINED IN SLICE,
// prefix: FIRST NUMBER - COULD BE SPLICED,
// suffix: LAST NUMBER - COULD BE SPLICED,
// start: STARING BIT POSITION, end: ENDING BIT POSITION}
func main() {

	//Take in command line inputs here
	M_input := os.Args[1] //Should be integer for number of threads
	fname := os.Args[2]   //Should be directory adddress for file containing numbers

	M, err := strconv.ParseFloat(M_input, 64)
	check(err)

	multi_add(M, fname)
}

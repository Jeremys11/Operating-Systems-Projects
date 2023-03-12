package main

import (
	//Used for Printouts

	"encoding/json"
	"math"
	"os" //Takes in command line inputs
	"strconv"
)

// Json message to send to worker
// datafile: fname, start: pos1 , end: pos2
// For whatever reason the member variables need to be capitalized
type WorkerMessage struct {
	Datafile string
	Start    int
	End      int
}

// Error checking if file was read correctly
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func multi_add(M float64, fname string) {
	var worker_message WorkerMessage

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
			worker_message = WorkerMessage{fname, index, index + int(fragment_size) + int(extra_numbers) - 1}
		} else {
			worker_message = WorkerMessage{fname, index, index + int(fragment_size) - 1}
		}
		index = index + int(fragment_size) //iterating index

		//Sending job and message to worker
		send_message, _ := json.Marshal(worker_message)
		partial_sum(send_message)

		//Logic after getting partial sums here

		//Probably should wait here for all threads
		//Computation of total sum here
	}
}

func partial_sum(parameter_message []byte) {

	//Get json file containing info
	var my_message WorkerMessage
	err := json.Unmarshal(parameter_message, &my_message)
	check(err)

	//Getting individual components
	filename := my_message.Datafile
	start := my_message.Start
	end := my_message.End

	// Reading file numbers into file variable
	number_file, err := os.Open(filename)
	check(err)
	defer number_file.Close()

	//Logic to get start and end
	//Sum all non start and end
	//Write to a json to return to coordinator
	return
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

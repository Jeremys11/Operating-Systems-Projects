package main

import (
	pb "Project2/runserver"
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Class Token
type Token struct {
	ID              string
	NAME            string
	LOW             uint64
	MID             uint64
	HIGH            uint64
	PARTIAL_VALUE   uint64
	FINAL_VALUE     uint64
	TIME_LAST_WRITE time.Time
}

// Class YamlInfo
// Holds data parsed from Yaml file
type YamlInfo struct {
	TOKEN   string `yaml:"token"`
	WRITER  string `yaml:"writer"`
	READERS string `yaml:"readers"`
}

type ParsedYaml struct {
	TOKEN        string
	WRITER       string
	READER_ARRAY []string
}

// Class WRITE_RECORD
type WRITE_RECORD struct {
	TOKEN       string
	FINAL_VALUE uint64
	TIME        time.Time
}

// Holds records of all writes
var writeMap = make(map[string]WRITE_RECORD)

// Name of yaml file containing token information
var yaml_name = "token.yaml"

// Default port 50051
var default_port = flag.Int("port", 50051, "The server port")

// Holds database of tokens
var tokenMap = make(map[string]Token)

// Fail-Silent Emulation
var crashToken = "2"

// server is used to implement runserver.RunService
type server struct {
	pb.UnimplementedRunServiceServer
}

// readYaml(yaml_file string, token string)
// Reads yaml file with token id, writer, and reader list
// Returns info for token with id match
func readYaml(yaml_name string, token string) ParsedYaml {

	//Read yaml file
	yaml_file, err := ioutil.ReadFile(yaml_name)
	if err != nil {
		log.Fatal(err)
	}

	//Decoder to parse yaml file
	decoder := yaml.NewDecoder(bytes.NewBufferString(string(yaml_file)))

	//Parsing through yaml file
	for {
		var yamlObject YamlInfo
		err = decoder.Decode(&yamlObject)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(errors.New("failed to decode yaml file"))
		}

		//found token id match, returning token info
		if yamlObject.TOKEN == token {
			var readers []string
			var writers []string

			readers = strings.Split(yamlObject.READERS, " ")
			writers = strings.Split(yamlObject.WRITER, " ")

			var newReaders []string
			var newWriters []string

			//Parse for readers
			for _, addr := range readers {
				port := strings.Index(addr, ":")
				if addr[len(addr)-1:] == "," {
					newReaders = append(newReaders, addr[port+1:len(addr)-1])
				} else {
					newReaders = append(newReaders, addr[port+1:])
				}
			}

			//Parse for writer
			for _, addr := range writers {
				port := strings.Index(addr, ":")
				if addr[len(addr)-1:] == "," {
					newWriters = append(newWriters, addr[port+1:len(addr)-1])
				} else {
					newWriters = append(newWriters, addr[port+1:])
				}
			}

			var parsedYaml ParsedYaml
			parsedYaml.TOKEN = yamlObject.TOKEN
			parsedYaml.READER_ARRAY = newReaders
			parsedYaml.WRITER = newWriters[0]

			// Return the list of ports
			return parsedYaml
		}
	}
	return ParsedYaml{}
}

// onClose()
// Printout all token information
// Returns nothing
func onClose() {

	if len(tokenMap) == 0 {
		fmt.Println("No Tokens")
	} else {
		for key, value := range tokenMap {
			fmt.Println("Key:", key, "Name:", value.NAME, " ", "Low:", value.LOW, " ", "Mid:", value.MID, " ",
				"High:", value.HIGH, " ", "Partial value:", value.PARTIAL_VALUE, " ", "Final Value:", value.FINAL_VALUE)
		}
	}
	fmt.Println()
}

// Hash concatentates a message and a nonce and generates a hash value.
func Hash(name string, nonce uint64) uint64 {
	hasher := sha256.New()
	hasher.Write([]byte(fmt.Sprintf("%s %d", name, nonce)))
	return binary.BigEndian.Uint64(hasher.Sum(nil))
}

// Gets index of x that minimized Hash function in range of uint64 [start,stop]
func ArgMin(name string, start uint64, stop uint64) uint64 {
	var minVal uint64
	var minX uint64
	var hashVal uint64

	for i := start; i < stop; i++ {
		if i == start {
			minX = i
			minVal = Hash(name, i)
		} else {
			hashVal = Hash(name, i)
			if minVal > hashVal {
				minVal = hashVal
				minX = i
			}
		}
	}

	return minX
}

// Get final values and timing
//
// Returns token
func (s *server) GetFinalValue(ctx context.Context, in *pb.RPCHelper) (*pb.Write_Record, error) {
	for key := range writeMap {
		if key == in.GetID() {
			return &pb.Write_Record{
				FINAL_VALUE: writeMap[key].FINAL_VALUE,
				TIME:        timestamppb.New(writeMap[key].TIME),
			}, nil
		}
	}
	return &pb.Write_Record{}, nil
}

// create implements runserver.create
// Create a token with a with the given id
// Reset the token's sate to undefined/null
//
// Returns token and success or fail response
func (s *server) Create(ctx context.Context, in *pb.RPCHelper) (*pb.Token, error) {
	//Check membership
	for key := range tokenMap {
		if key == in.GetID() {
			onClose()
			return nil, errors.New("Token already in list")
		}
	}

	//Create new token
	newToken := Token{ID: in.GetID()}

	//Check if valid reader or writer

	//Get yaml info for reader/writers
	var parsedYaml = readYaml(yaml_name, in.GetID())

	//Add token
	tokenMap[in.GetID()] = newToken

	if in.SERVERTYPE == "Writer" {
		//iterating updates to reader servers
		for _, addr := range parsedYaml.READER_ARRAY {

			//Iterate through readers
			conn, err := grpc.Dial("localhost:"+addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("Could not connect: %v", err)
			}
			defer conn.Close()

			c := pb.NewRunServiceClient(conn)

			// Contact the server
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			c.Create(ctx, &pb.RPCHelper{
				ID:         in.GetID(),
				SERVERTYPE: "Reader",
			})
		}
	}

	onClose()
	return &pb.Token{ID: in.GetID()}, nil
}

// drop implements runserver.drop
// Delete token with given id
//
// Returns Token and error
func (s *server) Drop(ctx context.Context, in *pb.RPCHelper) (*pb.Token, error) {
	//Check membership
	for key := range tokenMap {
		if key == in.GetID() {

			//Remove membership
			delete(tokenMap, in.GetID())

			var parsedYaml = readYaml(yaml_name, in.GetID())
			//Telling writer servers to write same token
			if in.SERVERTYPE == "Writer" {
				for _, addr := range parsedYaml.READER_ARRAY {
					conn, err := grpc.Dial("localhost:"+addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
					if err != nil {
						log.Fatalf("Could not connect: %v", err)
					}
					defer conn.Close()

					c := pb.NewRunServiceClient(conn)

					// Contact the server
					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
					defer cancel()

					c.Drop(ctx, &pb.RPCHelper{
						ID:         in.GetID(),
						SERVERTYPE: "Reader",
					})
				}
			} //End Reader Calls

			onClose()
			return &pb.Token{ID: in.GetID()}, nil
		}
	}

	//Token not in list
	onClose()
	return nil, errors.New("Token not in list")
}

// write implements runserver.write
// Set name, low, mid, high based on id
// Compute the partial value of the toek as argmin_x H(name, x) for x in [low,mid) and
// Reset final value of token
//
// Return partial value on success or fail response
func (s *server) Write(ctx context.Context, in *pb.RPCHelper) (*pb.Token, error) {
	//Check membership
	for key, value := range tokenMap {
		if key == in.GetID() {

			//Emulate fail-silent for token = 2
			if in.GetID() == crashToken {
				time.Sleep(5 * time.Second)
				return &pb.Token{ID: in.GetID()}, nil
			} else {
				//Get yaml info for reader/writers
				var parsedYaml = readYaml(yaml_name, in.GetID())

				value.NAME = in.GetNAME()
				value.LOW = in.GetLOW()
				value.MID = in.GetMID()
				value.HIGH = in.GetHIGH()
				value.PARTIAL_VALUE = ArgMin(in.GetNAME(), in.GetLOW(), in.GetMID())
				value.FINAL_VALUE = 0

				//Getting final value
				temp := ArgMin(key, value.MID, value.HIGH)

				//Get min of temp final value and partial value and set final value accordingly
				if temp <= value.PARTIAL_VALUE {
					value.FINAL_VALUE = temp
				} else {
					value.FINAL_VALUE = value.PARTIAL_VALUE
				}

				tokenMap[key] = value //Reassign value back to key after update

				//Record of write operations
				writeMap[key] = WRITE_RECORD{TIME: time.Now(), FINAL_VALUE: value.FINAL_VALUE}

				//Telling writer servers to write same token
				if in.SERVERTYPE == "Writer" {
					for _, addr := range parsedYaml.READER_ARRAY {
						conn, err := grpc.Dial("localhost:"+addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
						if err != nil {
							log.Fatalf("Could not connect: %v", err)
						}
						defer conn.Close()

						c := pb.NewRunServiceClient(conn)

						// Contact the server
						ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
						defer cancel()

						c.Write(ctx, &pb.RPCHelper{
							ID:            in.GetID(),
							NAME:          in.GetNAME(),
							LOW:           in.GetLOW(),
							MID:           in.GetMID(),
							HIGH:          in.GetHIGH(),
							PARTIAL_VALUE: value.PARTIAL_VALUE,
							FINAL_VALUE:   value.FINAL_VALUE,
							SERVERTYPE:    "Reader",
						})
					}
				} //End Reader Calls

				//Return token and error
				onClose()
				return &pb.Token{
					ID:            in.GetID(),
					NAME:          in.GetNAME(),
					LOW:           in.GetLOW(),
					MID:           in.GetMID(),
					HIGH:          in.GetHIGH(),
					PARTIAL_VALUE: value.PARTIAL_VALUE,
					FINAL_VALUE:   value.FINAL_VALUE,
				}, nil
			}
		}
	}

	//Token not in list
	onClose()
	return nil, errors.New("Token not in list")
}

// read implements runserver.read
// Find argmin_x H(name,x) for x in [mid,high)
//
// Return token's final value on success or fail response
func (s *server) Read(ctx context.Context, in *pb.RPCHelper) (*pb.Token, error) {

	//Emulate fail-silent for token = 2
	if in.GetID() == crashToken {
		time.Sleep(5 * time.Second)
		return &pb.Token{ID: in.GetID(), FINAL_VALUE: in.FINAL_VALUE}, nil
	} else {
		//Check membership
		for key, value := range tokenMap {
			if key == in.GetID() {

				//Check for previous write operation performed
				if value.PARTIAL_VALUE != 0 {

					//Get yaml info for reader/writers
					var parsedYaml = readYaml(yaml_name, in.GetID())

					//Get final values from other readers
					if in.SERVERTYPE == "Writer" {
						var tempFinalVal = value.FINAL_VALUE
						var tempFinalTime = writeMap[in.ID].TIME
						for _, addr := range parsedYaml.READER_ARRAY {
							conn, err := grpc.Dial("localhost:"+addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
							if err != nil {
								log.Fatalf("Could not connect: %v", err)
							}
							defer conn.Close()

							c := pb.NewRunServiceClient(conn)

							// Contact the server
							ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
							defer cancel()
							temp2, _ := c.GetFinalValue(ctx, &pb.RPCHelper{ID: in.GetID()})

							if tempFinalTime.Before(temp2.TIME.AsTime()) {
								tempFinalVal = temp2.FINAL_VALUE
							}

						}
						value.FINAL_VALUE = tempFinalVal
					}
					//Return token and error
					onClose()
					return &pb.Token{
						ID:            in.GetID(),
						NAME:          value.NAME,
						LOW:           value.LOW,
						MID:           value.MID,
						HIGH:          value.HIGH,
						PARTIAL_VALUE: value.PARTIAL_VALUE,
						FINAL_VALUE:   value.FINAL_VALUE,
					}, nil
					//Error - no write before read
				} else {
					return nil, errors.New("must have written token at least once before read")
				}
			}
		}
		//Token not in list
		onClose()
		return nil, errors.New("Token not in list")
	}

}

// Main function brings server to life
// Code taken from helloworld example server file
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *default_port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRunServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

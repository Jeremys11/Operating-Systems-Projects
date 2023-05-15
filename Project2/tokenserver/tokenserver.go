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
	"sync"
	"time"

	"gopkg.in/yaml.v3"

	"google.golang.org/grpc"
)

// Class Token
type Token struct {
	ID            string
	NAME          string
	LOW           uint64
	MID           uint64
	HIGH          uint64
	PARTIAL_VALUE uint64
	FINAL_VALUE   uint64
}

// Class TokenMutex
type TokenMutex struct {
	TOKENMAP Token
	RWMUTEX  sync.RWMutex
}

// Class YamlInfo
// Holds data read from yaml file about token
type YamlInfo struct {
	TOKEN   string `yaml:"token"`
	WRITER  string `yaml:"writer"`
	READERS string `yaml:"readers"`
}

// Name of yaml file containing token information
var yaml_name = "token.yaml"

// Default port 50051
var default_port = flag.Int("port", 50051, "The server port")

// Holds database of tokens
var tokenMap = make(map[string]Token)

// Hold record of operations
var operationSlice []string

// Fail-Silent Emulation
// If crashBool = false, set crashBool to true and continue
// If crashBool = true, set crashBool to false and emulate fail-silent
var crashBool bool

// server is used to implement runserver.RunService
type server struct {
	pb.UnimplementedRunServiceServer
}

// readYaml(yaml_file string, token string)
// Reads yaml file with token id, writer, and reader list
// Returns info for token with id match
func readYaml(yaml_name string, token string) YamlInfo {

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
			return yamlObject
		}
	}
	return YamlInfo{}
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

// create implements runserver.create
// Create a token with a with the given id
// Reset the token's sate to undefined/null
//
// Returns token and success or fail response
func (s *server) Create(ctx context.Context, in *pb.Token) (*pb.Token, error) {
	//Check membership
	for key := range tokenMap {
		if key == in.GetID() {
			onClose()
			return nil, errors.New("Token already in list")
		}
	}

	//Create new token
	newToken := Token{ID: in.GetID()}

	//Lock token before adding
	//Get yaml info for reader/writers
	//Update token_list
	//Create reading servers if writer
	//Send rpc calls to update token_list for readers

	//Append new token to map
	tokenMap[in.GetID()] = newToken

	onClose()
	return &pb.Token{ID: in.GetID()}, nil
}

// drop implements runserver.drop
// Delete token with given id
//
// Returns Token and error
func (s *server) Drop(ctx context.Context, in *pb.Token) (*pb.Token, error) {
	//Check membership
	for key := range tokenMap {
		if key == in.GetID() {

			//Remove membership
			delete(tokenMap, in.GetID())

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
func (s *server) Write(ctx context.Context, in *pb.Token) (*pb.Token, error) {
	//Check membership
	for key, value := range tokenMap {
		if key == in.GetID() {

			//Lock token before modification
			//Get yaml info for reader/writers
			//Update token_list
			//Send rpc calls to update token_list for readers

			value.NAME = in.GetNAME()
			value.LOW = in.GetLOW()
			value.MID = in.GetMID()
			value.HIGH = in.GetHIGH()
			value.PARTIAL_VALUE = ArgMin(in.GetNAME(), in.GetLOW(), in.GetMID())
			value.FINAL_VALUE = 0

			tokenMap[key] = value //Reassign value back to key after update

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

	//Token not in list
	onClose()
	return nil, errors.New("Token not in list")
}

// read implements runserver.read
// Find argmin_x H(name,x) for x in [mid,high)
//
// Return token's final value on success or fail response
func (s *server) Read(ctx context.Context, in *pb.Token) (*pb.Token, error) {

	//Emulate fail-silent every second instruction
	if crashBool == true {
		crashBool = false
		time.Sleep(5 * time.Second)
		return &pb.Token{}, nil
	} else {
		//Reset CrashBool to true
		crashBool = true

		//Check membership
		for key, value := range tokenMap {
			if key == in.GetID() {

				//Check for previous write operation performed
				if value.PARTIAL_VALUE != 0 {
					//Lock token before modification
					//Get yaml info for reader/writers
					//Update token_list
					//Send rpc calls to update token_list for readers

					temp := ArgMin(key, value.MID, value.HIGH)

					//Get min of temp final value and partial value and set final value accordingly
					if temp <= value.PARTIAL_VALUE {
						value.FINAL_VALUE = temp
					} else {
						value.FINAL_VALUE = value.PARTIAL_VALUE
					}

					tokenMap[key] = value //Reassign value back to key after update

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
	}

	//Token not in list
	onClose()
	return nil, errors.New("Token not in list")
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

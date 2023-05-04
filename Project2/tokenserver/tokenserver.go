package main

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"

	pb "Project2/runserver"

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
	WRITER        map[string]string
	READER        map[string]string
}

/*
Extend each token to include the access points (IP address and port number) of its single writer
and multiple reader nodes; these nodes constitute the replication scheme of the token.
	token: <id>
	writer: <access-point>
	readers: array of <access-point>s
*/

var default_port = flag.Int("port", 50051, "The server port") //Default port 50051
// Try using Map for concurrency
// Rewrite:
// onClose
// rpc functions
var tokenMap = make(map[string]Token)

// Mutex for concurrency
//var mutex = &sync.Mutex{}

// server is used to implement runserver.RunService
type server struct {
	pb.UnimplementedRunServiceServer
}

// Test function
// Returns nothing
func (s *server) Test(ctx context.Context, in *pb.Token) (*pb.Token, error) {
	for key, value := range tokenMap {
		fmt.Println(key)
		fmt.Println(value.WRITER)
		fmt.Println(value.READER)
	}
	return nil, nil
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
	//mutex.Lock()
	//defer mutex.Unlock()
	for key := range tokenMap {
		if key == in.GetID() {
			onClose()
			return nil, errors.New("Token already in list")
		}
	}

	//Create new token
	newToken := Token{ID: in.GetID()}

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
	//mutex.Lock()
	//defer mutex.Unlock()
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
	//mutex.Lock()
	//defer mutex.Unlock()
	for key, value := range tokenMap {
		if key == in.GetID() {

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
	//Check membership
	//mutex.Lock()
	//defer mutex.Unlock()
	for key, value := range tokenMap {
		if key == in.GetID() {
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

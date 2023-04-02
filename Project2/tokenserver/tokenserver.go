package main

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"

	pb "Project2/runserver"
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

var (
	default_port = flag.Int("port", 50051, "The server port") //Default port 50051
	tokenSlice   = []Token{}                                  //Golang uses slices -- dynamic arrays
)

// server is used to implement runserver.RunService
type server struct {
	pb.UnimplementedRunServiceServer
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
func (s *server) create(ctx context.Context, in *pb.Token) (*pb.Token, error) {
	//Check membership
	for i := 0; i < len(tokenSlice); i++ {
		if tokenSlice[i].ID == in.GetID() {
			return nil, errors.New("Token already in list")
		}
	}

	//Create new token
	newToken := Token{ID: in.GetID()}

	//Append new token to slice
	tokenSlice = append(tokenSlice, newToken)

	return &pb.Token{ID: in.GetID()}, nil
}

// drop implements runserver.drop
// Delete token with given id
//
// Returns Token and error
func (s *server) drop(ctx context.Context, in *pb.Token) (*pb.Token, error) {
	//Check membership
	length := len(tokenSlice)
	for i := 0; i < length; i++ {
		if tokenSlice[i].ID == in.GetID() {
			//Remove membership and adjust slice size

			//Order unimportant
			//Replace the element to delete with the one at the end of the slice
			//Return the n-1 fist elements

			tokenSlice[i] = tokenSlice[length-1]
			tokenSlice = tokenSlice[:length-1]
			return &pb.Token{ID: in.GetID()}, nil
		}
	}

	//Token not in list
	return nil, errors.New("Token not in list")
}

// write implements runserver.write
// Set name, low, mid, high based on id
// Compute the partial value of the toek as argmin_x H(name, x) for x in [low,mid) and
// Reset final value of token
//
// Return partial value on success or fail response
func (s *server) write(ctx context.Context, in *pb.Token) (*pb.Token, error) {
	//Check membership
	for i := 0; i < len(tokenSlice); i++ {
		if tokenSlice[i].ID == in.GetID() {
			tokenSlice[i].NAME = in.GetNAME()
			tokenSlice[i].LOW = in.GetLOW()
			tokenSlice[i].MID = in.GetMID()
			tokenSlice[i].HIGH = in.GetHIGH()
			tokenSlice[i].PARTIAL_VALUE = ArgMin(in.GetNAME(), in.GetLOW(), in.GetMID())
			tokenSlice[i].FINAL_VALUE = 0

			//Return token and error
			return &pb.Token{
				ID:            in.GetID(),
				NAME:          in.GetNAME(),
				LOW:           in.GetLOW(),
				MID:           in.GetMID(),
				HIGH:          in.GetHIGH(),
				PARTIAL_VALUE: tokenSlice[i].PARTIAL_VALUE,
				FINAL_VALUE:   tokenSlice[i].FINAL_VALUE,
			}, nil
		}
	}

	//Token not in list
	return nil, errors.New("Token not in list")
}

// read implements runserver.read
// Find argmin_x H(name,x) for x in [mid,high)
//
// Return token's final value on success or fail response
func (s *server) read(ctx context.Context, in *pb.Token) (*pb.Token, error) {
	//Check membership
	for i := 0; i < len(tokenSlice); i++ {
		if tokenSlice[i].ID == in.GetID() {
			temp := ArgMin(tokenSlice[i].ID, tokenSlice[i].MID, tokenSlice[i].HIGH)

			//Get min of temp final value and partial value and set final value accordingly
			if temp <= tokenSlice[i].PARTIAL_VALUE {
				tokenSlice[i].FINAL_VALUE = temp
			} else {
				tokenSlice[i].FINAL_VALUE = tokenSlice[i].PARTIAL_VALUE
			}

			//Return token and error
			return &pb.Token{
				ID:            in.GetID(),
				NAME:          tokenSlice[i].NAME,
				LOW:           tokenSlice[i].LOW,
				MID:           tokenSlice[i].MID,
				HIGH:          tokenSlice[i].HIGH,
				PARTIAL_VALUE: tokenSlice[i].PARTIAL_VALUE,
				FINAL_VALUE:   tokenSlice[i].FINAL_VALUE,
			}, nil
		}
	}

	//Token not in list
	return nil, errors.New("Token not in list")
}

func main() {

	fmt.Println("Testing Printout")
}

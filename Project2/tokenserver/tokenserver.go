package main

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
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
	port       = flag.Int("port", 50051, "The server port")
	tokenSlice = []Token{} //Golang uses slices -- dynamic arrays
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
// Returns success or fail response
func (s *server) create(ctx context.Context, in *pb.Token) (*pb.Token, error) {}

// drop implements runserver.drop
// Delete token with given id
//
// Returns <NULL>
func (s *server) drop(ctx context.Context, in *pb.Token) (*pb.Token, error) {}

// write implements runserver.write
// Set name, low, mid, high based on id
// Compute the partial value of the toek as argmin_x H(name, x) for x in [low,mid) and
// Sreset final value of token
//
// Return partial value on success or fail response
func (s *server) write(ctx context.Context, in *pb.Token) (*pb.Token, error) {}

// read implements runserver.read
// Find argmin_x H(name,x) for x in [mid,high)
//
// Return token's final value on success or fail response
func (s *server) read(ctx context.Context, in *pb.Token) (*pb.Token, error) {}

func main() {

	fmt.Println("Testing Printout")
}

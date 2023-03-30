package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

// Class Token
type Token struct {
	ID   string
	NAME string
	LOW  uint64
	HIGH uint64
	MID  uint64
}

// Hash concatentates a message and a nonce and generates a hash value.
func Hash(name string, nonce uint64) uint64 {
	hasher := sha256.New()
	hasher.Write([]byte(fmt.Sprintf("%s %d", name, nonce)))
	return binary.BigEndian.Uint64(hasher.Sum(nil))
}

// create(id string)
// Create a token with a with the given id
// Reset the token's sate to undefined/null
//
// Returns success or fail response
func create(id string) {}

// drop(id string)
// Delete token with given id
//
// Returns <NULL>
func drop(id string) {}

// write(id string, name string, low uint64, high uint64, mid uint64)
// Set name, low, mid, high based on id
// Compute the partial value of the toek as argmin_x H(name, x) for x in [low,mid) and
// Sreset final value of token
//
// Return partial value on success or fail response
func write(argmin_x uint64) (id string, name string, low uint64, high uint64, mid uint64) {

	return
}

// read(id string)
// Find argmin_x H(name,x) for x in [mid,high)
//
// Return token's final value on success or fail response
func read(argmin_x uint64) (id string) {

	return
}

func main() {

	fmt.Println("Testing Printout")
}

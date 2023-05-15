package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "Project2/runserver"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	//localhost:50051
	port = flag.String("port", "50051", "the port address to connect to, default 50051")
	host = flag.String("host", "localhost", "the host address to connect to, default localhost")

	createOp = flag.Bool("create", false, "Create Op")
	dropOp   = flag.Bool("drop", false, "Drop Op")
	writeOp  = flag.Bool("write", false, "Write Op")
	readOp   = flag.Bool("read", false, "Read Op")

	idVal   = flag.String("id", "", "ID value")
	nameVal = flag.String("name", "", "Name value")
	lowVal  = flag.Uint64("low", 0, "Low value")
	midVal  = flag.Uint64("mid", 0, "Mid value")
	highVal = flag.Uint64("high", 0, "High value")
)

// Main function brings client to life
func main() {

	//Get command line inputs
	flag.Parse()

	addr := *host + ":" + *port
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRunServiceClient(conn)

	// Contact the server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//Listen for:
	//Create
	if *createOp {
		val, err := c.Create(ctx, &pb.Token{
			ID: *idVal,
		})
		if err != nil {
			log.Fatal("Failed to create", err)
		} else {
			log.Print("Successful Create", val.ID)
		}
	}
	//Drop
	if *dropOp {
		val, err := c.Drop(ctx, &pb.Token{
			ID: *idVal,
		})
		if err != nil {
			log.Fatal("Failed to drop", err)
		} else {
			log.Print("Successful Drop", val.ID)
		}
	}
	//Write
	if *writeOp {
		val, err := c.Write(ctx, &pb.Token{
			ID:   *idVal,
			NAME: *nameVal,
			LOW:  *lowVal,
			MID:  *midVal,
			HIGH: *highVal,
		})
		if err != nil {
			log.Fatal("Failed to write", err)
		} else {
			log.Print("Successful Write", val.ID, " Partial Value: ", val.PARTIAL_VALUE)
		}
	}
	//Read
	if *readOp {
		val, err := c.Read(ctx, &pb.Token{
			ID: *idVal,
		})
		if err != nil {
			log.Fatal("Failed to read", err)
		} else {
			log.Print("Successful Read", val.ID, " Partial Value: ", val.PARTIAL_VALUE)
		}
	}

}

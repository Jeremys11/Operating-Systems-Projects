package main

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"

	pb "Project2/runserver"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v3"
)

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

var (
	createOp = flag.Bool("create", false, "Create Op")
	dropOp   = flag.Bool("drop", false, "Drop Op")
	writeOp  = flag.Bool("write", false, "Write Op")
	readOp   = flag.Bool("read", false, "Read Op")

	idVal   = flag.String("id", "", "ID value")
	nameVal = flag.String("name", "", "Name value")
	lowVal  = flag.Uint64("low", 0, "Low value")
	midVal  = flag.Uint64("mid", 0, "Mid value")
	highVal = flag.Uint64("high", 0, "High value")

	// Name of yaml file containing token information
	yaml_name = "token.yaml"
)

// Main function brings client to life
func main() {

	//Get command line inputs
	flag.Parse()

	//Port depends on token
	var parsedYaml = readYaml(yaml_name, *idVal)

	var addr string
	if *readOp {
		randomIndex := rand.Intn(len(parsedYaml.READER_ARRAY))
		pick := parsedYaml.READER_ARRAY[randomIndex]
		addr = "localhost:" + pick
	} else {
		addr = "localhost:" + parsedYaml.WRITER
	}

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
		val, err := c.Create(ctx, &pb.RPCHelper{
			ID:         *idVal,
			SERVERTYPE: "Writer",
			ADDRESS:    addr,
		})
		if err != nil {
			log.Fatal("Failed to create", err)
		} else {
			log.Print("Successful Create", val.ID)
		}
	}
	//Drop
	if *dropOp {
		val, err := c.Drop(ctx, &pb.RPCHelper{
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
		val, err := c.Write(ctx, &pb.RPCHelper{
			ID:         *idVal,
			NAME:       *nameVal,
			LOW:        *lowVal,
			MID:        *midVal,
			HIGH:       *highVal,
			SERVERTYPE: "Writer",
			ADDRESS:    addr,
		})
		if err != nil {
			log.Fatal("Failed to write", err)
		} else {
			log.Print("Successful Write", val.ID, " Partial Value: ", val.PARTIAL_VALUE)
		}
	}
	//Read
	if *readOp {
		val, err := c.Read(ctx, &pb.RPCHelper{
			ID:         *idVal,
			SERVERTYPE: "Writer",
			ADDRESS:    addr,
		})
		if err != nil {
			log.Fatal("Failed to read", err)
		} else {
			log.Print("Successful Read", val.ID, " Partial Value: ", val.FINAL_VALUE)
		}
	}

}

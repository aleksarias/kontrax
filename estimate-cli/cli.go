// estimate-cli/cli.go
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/net/context"
	pb "kontrax/estimate-service/proto/estimate"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
)

const (
	defaultFilename = "estimate.json"
)

func parseFile(file string) (*pb.Estimate, error) {
	var estimate *pb.Estimate
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &estimate)
	if err != nil {
		return nil, err
	}
	return estimate, err
}

func main() {
	err := cmd.Init()
	if err != nil {
		log.Fatalf("Could not init program: %v", err)
	}

	// Create new greeter client
	client := pb.NewEstimateServiceClient("go.micro.srv.estimate", microclient.DefaultClient)

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateEstimate(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetEstimates(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Estimates {
		log.Println(v)
	}

}

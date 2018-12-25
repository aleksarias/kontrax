// estimate-service/main.go
package main

import (
	"fmt"
	"github.com/micro/go-micro"
	// Import the generated protobuf code
	pb "kontrax/estimate-service/proto/estimate"

	"golang.org/x/net/context"
)

const (
	port = ":50051"
)

type IRepository interface {
	Create(estimate *pb.Estimate) (*pb.Estimate, error)
	GetAll() []*pb.Estimate
}

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type Repository struct {
	estimates []*pb.Estimate
}

func (repo *Repository) Create(estimate *pb.Estimate) (*pb.Estimate, error) {
	updated := append(repo.estimates, estimate)
	repo.estimates = updated
	return estimate, nil
}

// Get all estimates
func (repo *Repository) GetAll() []*pb.Estimate {
	return repo.estimates
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo IRepository
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) CreateEstimate(ctx context.Context, req *pb.Estimate, res *pb.Response,)  error {

	// Save our consignment
	estimate, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Estimate = estimate

	// Return matching the `Response` message we created in our
	// protobuf definition.
	return nil
}

func (s *service) GetEstimates(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	estimates := s.repo.GetAll()
	res.Estimates = estimates

	return nil
}

func main() {

	repo := &Repository{}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.estimate"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	pb.RegisterEstimateServiceHandler(srv.Server(), &service{repo})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

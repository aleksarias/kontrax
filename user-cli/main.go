package main

import (
	"log"
	"os"

	pb "github.com/aleksarias/kontrax/user-service/proto/user"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"golang.org/x/net/context"
)

func main() {

	err := cmd.Init()
	if err != nil {
		log.Panic("Not able to initialize go-micro commandline interface")
	}

	// Create new greeter client
	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	// Define our flags
	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "You full name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "Your email",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
			cli.StringFlag{
				Name:  "organization",
				Usage: "Your organization id",
			},
		),
	)

	// Start as service
	service.Init(

		micro.Action(func(c *cli.Context) {

			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			organization := c.String("company")

			// Call our user service
			r, err := client.Create(context.TODO(), &pb.User{
				FirstName:    name,
				Email:        email,
				Password:     password,
				Organization: organization,
			})
			if err != nil {
				log.Fatalf("Could not create: %v", err)
			}
			log.Printf("Created: %s", r.User.Id)

			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				log.Fatalf("Could not list users: %v", err)
			}
			for _, v := range getAll.Users {
				log.Println(v)
			}

			os.Exit(0)
		}),
	)

	// Run the server
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}

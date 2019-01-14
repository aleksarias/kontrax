package main

import (
	"context"
	"fmt"
	pb "github.com/aleksarias/kontrax/user-service/proto/user"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Create bogus database for testing
func CreateBogusConnection() (*gorm.DB, error) {

	return gorm.Open("sqlite3", "/tmp/gorm.db")

}

func test_create_delete_user() {
	db, err := CreateBogusConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Automatically migrates the all structs into database columns/types etc. This will
	// check for changes and migrate them each time this service is restarted.
	db.AutoMigrate(&pb.User{}, &pb.Phone{}, &pb.PostalAddress{}, &pb.Organization{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	service := &service{repo, tokenService}

	address := pb.PostalAddress{
		Address: "112 NE 5th Street Unit 1",
		City:    "Fort Lauderdale",
		State:   "Florida",
		Zip:     "33301",
	}

	addresses := []*pb.PostalAddress{&address}

	user := pb.User{
		FirstName:  "Alex",
		LastName:   "Arias",
		MiddleName: "Luis",
		Addresses:  addresses,
	}

	response := pb.Response{}

	err = service.Create(context.TODO(), &user, &response)
	if err != nil {
		panic(err)
	}

	request := pb.Request{}

	err = service.GetAll(context.TODO(), &request, &response)

	fmt.Print(response)

}

func main() {
	test_create_delete_user()
}

package user

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"log"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	userUuid, err := uuid.NewV4()
	if err != nil {
		log.Panicf("Error during user creaion. Not able to generate new UUID for user.")
	}
	return scope.SetColumn("Id", userUuid.String())
}

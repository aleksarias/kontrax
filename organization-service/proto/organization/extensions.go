package go_micro_srv_organization

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"log"
)

func (model *Organization) BeforeCreate(scope *gorm.Scope) error {
	userUuid, err := uuid.NewV4()
	if err != nil {
		log.Panicf("Error during organization creaion. Not able to generate new UUID for organization.")
	}
	return scope.SetColumn("Id", userUuid.String())
}

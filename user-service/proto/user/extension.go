package go_micro_srv_user

import (
	uuid "github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("created uuid error: %v\n", err)
	}
	return scope.SetColumn("Id", uuid.String())
}

package db

import (
	"go-rest-api/models"

	"github.com/objectbox/objectbox-go/objectbox"
)

var OB *objectbox.ObjectBox

func InitDB() error {
	builder := objectbox.NewBuilder()
	builder.Model(models.ObjectBoxModel())

	var err error
	OB, err = builder.Build()
	if err != nil {
		return err
	}
	return nil
}

func CloseDB() {
	if OB != nil {
		OB.Close()
	}
} 
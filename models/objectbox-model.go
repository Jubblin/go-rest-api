// Code generated by ObjectBox; DO NOT EDIT.

package models

import (
	"github.com/objectbox/objectbox-go/objectbox"
)

// ObjectBoxModel declares and builds the model from all the entities in the package.
// It is usually used when setting-up ObjectBox as an argument to the Builder.Model() function.
func ObjectBoxModel() *objectbox.Model {
	model := objectbox.NewModel()
	model.GeneratorVersion(6)

	model.RegisterBinding(DeviceActivityBinding)
	model.LastEntityId(1, 2906110396233178886)
	model.LastIndexId(3, 5847086806180283837)

	return model
}

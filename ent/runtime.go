// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
	"training-ent/ent/group"
	"training-ent/ent/schema"
	"training-ent/ent/user"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[0].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = groupDescName.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[2].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
	// userDescActive is the schema descriptor for active field.
	userDescActive := userFields[3].Descriptor()
	// user.DefaultActive holds the default value on creation for the active field.
	user.DefaultActive = userDescActive.Default.(bool)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[7].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}

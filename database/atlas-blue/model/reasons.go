//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Reasons struct {
	ID                  string `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	EngagementID        string // (DC2Type:guid)
	ReasonForUnadjusted *string
	CreatedAt           *time.Time // (DC2Type:datetime_immutable)
}

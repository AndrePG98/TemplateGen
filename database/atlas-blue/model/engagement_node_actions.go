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

type EngagementNodeActions struct {
	ID               string `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	EngagementNodeID string // (DC2Type:guid)
	PerformedBy      string // (DC2Type:guid)
	ActionType       string
	PerformedByRole  *string
	PerformedAt      *time.Time // (DC2Type:datetime_immutable)
}

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

type EngagementsAuditTrails struct {
	ID             string `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	EngagementID   string // (DC2Type:guid)
	UserID         string // (DC2Type:guid)
	PackageID      *string
	OccurredAt     time.Time
	NodeID         *string
	Tab            *string
	Fieldset       *string
	Field          *string
	Action         *string
	OtherReference *string
	PreviousValue  *string
	NewValue       *string
	FormID         *string
}

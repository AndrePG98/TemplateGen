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

type WorkProgrammeProcedures struct {
	ID              string `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	WorkProgrammeID string // (DC2Type:guid)
	ProcedureID     string // (DC2Type:guid)
	Number          *int32
	IsDeleted       *bool
	CreatedAt       *time.Time // (DC2Type:date_immutable)
}

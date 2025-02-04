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

type FslrControls struct {
	ID                  string `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	EngagementID        string // (DC2Type:guid)
	FslrAreaID          *string
	OtherFslrID         *string
	Name                *string
	SourceID            *string
	ControlType         *string
	ControlFrequency    *string
	PreventOrDetect     *string
	ControlLevel        *string
	PlannedEvaluationDi *bool
	PlannedEvaluationOe *bool
	ItApplicationID     *string
	CreatedAt           time.Time
	IsDeleted           bool
}

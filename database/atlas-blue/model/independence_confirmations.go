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

type IndependenceConfirmations struct {
	ID                         string `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	EngagementID               string // (DC2Type:guid)
	HasGiftsOrHospitality      *bool
	GiftsOrHospitalityDetails  *string
	HasFinancialInterests      *bool
	FinancialInterestsDetails  *string
	HasSafeguards              *bool
	SafeguardsDetails          *string
	IsIndependent              *bool
	IndependenceConsiderations *string
	CompletionDate             *time.Time // (DC2Type:date_immutable)
	AmendmentDate              *time.Time // (DC2Type:date_immutable)
	Status                     bool
	EngagementTeamMemberID     string // (DC2Type:guid)
	Hidden                     *bool
}

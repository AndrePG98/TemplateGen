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

type Procedures struct {
	ID                                       string  `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	KbpControlID                             *string // (DC2Type:guid)
	ItControlID                              *string // (DC2Type:guid)
	FslrControlID                            *string // (DC2Type:guid)
	ProcedureTypeLovID                       *string
	ResponsibilityLovID                      *string
	Description                              *string
	DetailedDescription                      *string
	AddToCurrentEngagementInTheFuture        *bool
	SourceID                                 *string // Procedure id from customisation side
	IsMandatory                              *bool
	CreatedAt                                *time.Time // (DC2Type:date_immutable)
	TeamMemberResponsibleID                  *string    // (DC2Type:guid)
	DescriptionLabel                         *string    // (DC2Type:json)
	DateLastDescriptionLabelModified         *string    // (DC2Type:json)
	DetailedDescriptionLabel                 *string    // (DC2Type:json)
	DateLastDetailedDescriptionLabelModified *string    // (DC2Type:json)
}

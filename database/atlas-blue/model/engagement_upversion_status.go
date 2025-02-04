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

type EngagementUpversionStatus struct {
	ID                     int32  `sql:"primary_key"`
	EngagementID           string // (DC2Type:guid)
	UpversionedOrRefusedBy string // (DC2Type:guid)
	PreviousLocalPackage   string
	UpversionLocalPackage  string
	UpversionDate          time.Time // (DC2Type:datetime_immutable)
	UpversionStatus        string
	Comment                *string
	Changes                *string // (DC2Type:json)
}

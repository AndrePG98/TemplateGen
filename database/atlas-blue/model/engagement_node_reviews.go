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

type EngagementNodeReviews struct {
	ID                         string     `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	EngagementMasterfileNodeID string     // (DC2Type:guid)
	ReviewedBy                 *string    // (DC2Type:guid)
	ReversedBy                 *string    // (DC2Type:guid)
	ReviewDate                 *time.Time // (DC2Type:datetime_immutable)
	Reversed                   bool
	ReversedDate               *time.Time // (DC2Type:datetime_immutable)
	ReviewedByEqcr             bool
}

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

type EntitiesDeletionRequests struct {
	ID                         string  `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	EntityID                   string  // (DC2Type:guid)
	InitiatedBy                *string // (DC2Type:guid)
	ApprovedBy                 *string // (DC2Type:guid)
	ReasonForDeletionID        *int32
	AdditionnalComment         *string
	InitiatedOn                *time.Time
	DeletionStatusID           *int32
	ApprovedOn                 *time.Time
	ReasonForRejection         *string
	DeletionRetrievalRequestID *string
	Status                     *string
}

//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type EngagementProgress struct {
	ID              string `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	EngagementID    string // (DC2Type:guid)
	AuditWorkflow   string
	TotalNodesCount int32
	ValidNodesCount int32
	HasConflicts    bool
	IsSignedOff     bool
}

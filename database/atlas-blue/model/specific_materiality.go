//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type SpecificMateriality struct {
	ID                              string  `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	PlanningMaterialityID           *string // (DC2Type:guid)
	RevisedMaterialityID            *string // (DC2Type:guid)
	HasParticularClassesTransaction *bool
	Comment                         *string
}

//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type MasterGroups struct {
	ID       string `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	EntityID string // (DC2Type:guid)
	Type     string
	LovID    int32
}

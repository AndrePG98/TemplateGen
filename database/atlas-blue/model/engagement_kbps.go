//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type EngagementKbps struct {
	ID                       string `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	EngagementID             string // (DC2Type:guid)
	Name                     string
	Documentation            *string
	SourceID                 *string
	WalkThroughDocumentation *bool
	Comment                  *string
	Label                    *string // (DC2Type:json)
	DateLastLabelModified    *string // (DC2Type:json)
}

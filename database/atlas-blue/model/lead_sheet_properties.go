//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type LeadSheetProperties struct {
	UserID           string `sql:"primary_key"` // (DC2Type:guid)
	LeadSheetID      string `sql:"primary_key"` // (DC2Type:guid)
	RoundingLovID    *string
	AccountGrouping  *string
	Areas            *string // (DC2Type:array)
	ApplyToAllSheets *bool
}

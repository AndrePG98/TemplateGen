//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type LeadSheetTableProperties struct {
	LeadSheetTableID string `sql:"primary_key"` // (DC2Type:guid)
	ColumnName       string `sql:"primary_key"`
	Order            *int32
	Width            *float64
}

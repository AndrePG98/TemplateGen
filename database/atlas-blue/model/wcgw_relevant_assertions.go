//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type WcgwRelevantAssertions struct {
	ID                    string `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	CotabdKbpsID          string // (DC2Type:guid)
	RelevantAssertionType string
	RiskID                string
	MarkedAsSignificant   *bool
}

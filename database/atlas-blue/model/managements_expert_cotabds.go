//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type ManagementsExpertCotabds struct {
	ID                  string `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	ManagementsExpertID string // (DC2Type:guid)
	CotabdID            string // (DC2Type:guid)
}

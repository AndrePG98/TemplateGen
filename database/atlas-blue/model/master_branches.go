//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type MasterBranches struct {
	ID                      string  `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	EntityID                string  // (DC2Type:guid)
	LeadPartner             *string // (DC2Type:guid)
	Name                    string
	CountryCode             *string
	Address                 *string
	City                    *string
	State                   *string
	PostalCode              *string
	PrimaryPhoneNumber      *string
	PrimaryOtherPhoneNumber *string
	DescriptionActivity     *string
	NumberEmployes          *string
	UniqueGlobalNumber      *string
}

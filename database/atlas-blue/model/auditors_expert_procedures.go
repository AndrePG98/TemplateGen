//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type AuditorsExpertProcedures struct {
	ID               string `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	AuditorsExpertID string // (DC2Type:guid)
	ProcedureID      string
	RiskTypeID       *string
	FslrNameID       *string
}

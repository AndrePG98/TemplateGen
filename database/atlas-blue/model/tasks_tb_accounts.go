//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type TasksTbAccounts struct {
	ID          string `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	TbAccountID string // (DC2Type:guid)
	TaskID      string // (DC2Type:guid)
}

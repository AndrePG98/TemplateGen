//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type DmsProxy struct {
	ID          string `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	ReceiptDate time.Time
	Service     string
	Method      string
	Resource    string
	Version     string
	Headers     string  // (DC2Type:array)
	Content     *string // (DC2Type:array)
	ContentRaw  *string
}

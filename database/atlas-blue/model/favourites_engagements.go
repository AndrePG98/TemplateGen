//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type FavouritesEngagements struct {
	User       string `sql:"primary_key"` // (DC2Type:guid)
	Engagement string `sql:"primary_key"` // (DC2Type:guid)
}

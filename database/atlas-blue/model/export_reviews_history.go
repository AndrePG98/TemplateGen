//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type ExportReviewsHistory struct {
	ID                              int32 `sql:"primary_key"`
	DocumentID                      *int32
	DocumentVersion                 *int32
	UserID                          *string
	ReviewedBy                      *string
	ReviewStatusNumber              *bool
	ReviewedAt                      *string
	DocreviewOnChangeEntertimestamp *string
	UserEntertimestamp              *string
	EngagementID                    *int32
	AtlasEngagementID               *string
	DocumentAttachedAt              *string
	Imported                        *bool
}

//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type AnalyticalReviewFinal struct {
	ID                                             string `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	EngagementID                                   string // (DC2Type:guid)
	ConfirmUsingAnalyticalReview                   *bool
	ConfirmUsingAnalyticalReviewComment            *string
	AnalyticalIdentifyFluctuations                 *bool
	AnalyticalIdentifyFluctuationsDifferencesFound *string
	DocumentInvestigationDifferences               *string
	DidFindMisstatement                            *bool
	DidFindMisstatementComment                     *string
	IsValid                                        bool
}

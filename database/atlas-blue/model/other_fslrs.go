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

type OtherFslrs struct {
	ID                                              string `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	EngagementID                                    string // (DC2Type:guid)
	Name                                            *string
	EvaluateNatureAndExtentOfPervasiveEffectOnFs    *string
	DoesFsRiskAffectAssessmentRisksAtAssertionLevel *bool
	DocumentationOfHowFslrAffectsAssessmentOfRisks  *string
	IsThisRiskOfMaterialMisstatementDueToFraud      *bool
	OtherFslrAssessment                             *string
	CommunicationToCountryRiskManager               *bool
	IsCreated                                       time.Time
	IsDeleted                                       *bool
	AsoCr                                           *string
}

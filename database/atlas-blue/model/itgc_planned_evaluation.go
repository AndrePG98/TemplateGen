//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type ItgcPlannedEvaluation struct {
	ID                                       string `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	ItApplicationsItRisksItGeneralControlsID string // (DC2Type:guid)
	ContextID                                *string
	PlannedEvaluationOfOe                    *bool
	PlannedEvaluationOfDi                    *bool
}

//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type CompletionGoingConcernFinalConclusion struct {
	ID                                            string `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	EngagementID                                  string // (DC2Type:guid)
	NodeID                                        string
	EntityStatus                                  *string
	AccountingBasisChange                         *bool
	AccountingBasisPlanningStage                  *string
	AccountingBasisCompletionStage                *string
	AccountingBasisCompletionStageComments        *string
	GoingConcernAssessment                        *bool
	GoingConcernAssessmentComments                *string
	MaterialUncertainties                         *bool
	MaterialUncertaintiesComments                 *string
	Situation                                     *string
	ImplicationsReport                            *bool
	IsSituation                                   *bool
	ConfirmCompliedApplication                    *bool
	ConfirmCompliedImplicationComment             *string
	ExtendAssessment                              *bool
	ConsiderImplications                          *string
	EffectOnUnableObtainSufficientEvidence        *string
	EffectOnUnableObtainSufficientEvidenceComment *string
	SignificantDelayInApproval                    *string
	SignificantDelayInApprovalComment             *string
	DelayRelatedEventsConditions                  *bool
	ConfirmNecessityAdditionalAudit               *bool
	ConfirmNecessityAdditionalAuditComment        *string
	IsValid                                       *bool
	ConclusionImpactOnMaterialUncertainty         *bool
	DelayRelatedEventsConditionsComment           *string
}

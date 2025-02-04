//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/mysql"
)

var CompletionGoingConcernFinalConclusion = newCompletionGoingConcernFinalConclusionTable("atlas-blue", "completion_going_concern_final_conclusion", "")

type completionGoingConcernFinalConclusionTable struct {
	mysql.Table

	// Columns
	ID                                            mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	EngagementID                                  mysql.ColumnString // (DC2Type:guid)
	NodeID                                        mysql.ColumnString
	EntityStatus                                  mysql.ColumnString
	AccountingBasisChange                         mysql.ColumnBool
	AccountingBasisPlanningStage                  mysql.ColumnString
	AccountingBasisCompletionStage                mysql.ColumnString
	AccountingBasisCompletionStageComments        mysql.ColumnString
	GoingConcernAssessment                        mysql.ColumnBool
	GoingConcernAssessmentComments                mysql.ColumnString
	MaterialUncertainties                         mysql.ColumnBool
	MaterialUncertaintiesComments                 mysql.ColumnString
	Situation                                     mysql.ColumnString
	ImplicationsReport                            mysql.ColumnBool
	IsSituation                                   mysql.ColumnBool
	ConfirmCompliedApplication                    mysql.ColumnBool
	ConfirmCompliedImplicationComment             mysql.ColumnString
	ExtendAssessment                              mysql.ColumnBool
	ConsiderImplications                          mysql.ColumnString
	EffectOnUnableObtainSufficientEvidence        mysql.ColumnString
	EffectOnUnableObtainSufficientEvidenceComment mysql.ColumnString
	SignificantDelayInApproval                    mysql.ColumnString
	SignificantDelayInApprovalComment             mysql.ColumnString
	DelayRelatedEventsConditions                  mysql.ColumnBool
	ConfirmNecessityAdditionalAudit               mysql.ColumnBool
	ConfirmNecessityAdditionalAuditComment        mysql.ColumnString
	IsValid                                       mysql.ColumnBool
	ConclusionImpactOnMaterialUncertainty         mysql.ColumnBool
	DelayRelatedEventsConditionsComment           mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type CompletionGoingConcernFinalConclusionTable struct {
	completionGoingConcernFinalConclusionTable

	NEW completionGoingConcernFinalConclusionTable
}

// AS creates new CompletionGoingConcernFinalConclusionTable with assigned alias
func (a CompletionGoingConcernFinalConclusionTable) AS(alias string) *CompletionGoingConcernFinalConclusionTable {
	return newCompletionGoingConcernFinalConclusionTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CompletionGoingConcernFinalConclusionTable with assigned schema name
func (a CompletionGoingConcernFinalConclusionTable) FromSchema(schemaName string) *CompletionGoingConcernFinalConclusionTable {
	return newCompletionGoingConcernFinalConclusionTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CompletionGoingConcernFinalConclusionTable with assigned table prefix
func (a CompletionGoingConcernFinalConclusionTable) WithPrefix(prefix string) *CompletionGoingConcernFinalConclusionTable {
	return newCompletionGoingConcernFinalConclusionTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CompletionGoingConcernFinalConclusionTable with assigned table suffix
func (a CompletionGoingConcernFinalConclusionTable) WithSuffix(suffix string) *CompletionGoingConcernFinalConclusionTable {
	return newCompletionGoingConcernFinalConclusionTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCompletionGoingConcernFinalConclusionTable(schemaName, tableName, alias string) *CompletionGoingConcernFinalConclusionTable {
	return &CompletionGoingConcernFinalConclusionTable{
		completionGoingConcernFinalConclusionTable: newCompletionGoingConcernFinalConclusionTableImpl(schemaName, tableName, alias),
		NEW: newCompletionGoingConcernFinalConclusionTableImpl("", "new", ""),
	}
}

func newCompletionGoingConcernFinalConclusionTableImpl(schemaName, tableName, alias string) completionGoingConcernFinalConclusionTable {
	var (
		IDColumn                                            = mysql.StringColumn("id")
		EngagementIDColumn                                  = mysql.StringColumn("engagement_id")
		NodeIDColumn                                        = mysql.StringColumn("node_id")
		EntityStatusColumn                                  = mysql.StringColumn("entity_status")
		AccountingBasisChangeColumn                         = mysql.BoolColumn("accounting_basis_change")
		AccountingBasisPlanningStageColumn                  = mysql.StringColumn("accounting_basis_planning_stage")
		AccountingBasisCompletionStageColumn                = mysql.StringColumn("accounting_basis_completion_stage")
		AccountingBasisCompletionStageCommentsColumn        = mysql.StringColumn("accounting_basis_completion_stage_comments")
		GoingConcernAssessmentColumn                        = mysql.BoolColumn("going_concern_assessment")
		GoingConcernAssessmentCommentsColumn                = mysql.StringColumn("going_concern_assessment_comments")
		MaterialUncertaintiesColumn                         = mysql.BoolColumn("material_uncertainties")
		MaterialUncertaintiesCommentsColumn                 = mysql.StringColumn("material_uncertainties_comments")
		SituationColumn                                     = mysql.StringColumn("situation")
		ImplicationsReportColumn                            = mysql.BoolColumn("implications_report")
		IsSituationColumn                                   = mysql.BoolColumn("is_situation")
		ConfirmCompliedApplicationColumn                    = mysql.BoolColumn("confirm_complied_application")
		ConfirmCompliedImplicationCommentColumn             = mysql.StringColumn("confirm_complied_implication_comment")
		ExtendAssessmentColumn                              = mysql.BoolColumn("extend_assessment")
		ConsiderImplicationsColumn                          = mysql.StringColumn("consider_implications")
		EffectOnUnableObtainSufficientEvidenceColumn        = mysql.StringColumn("effect_on_unable_obtain_sufficient_evidence")
		EffectOnUnableObtainSufficientEvidenceCommentColumn = mysql.StringColumn("effect_on_unable_obtain_sufficient_evidence_comment")
		SignificantDelayInApprovalColumn                    = mysql.StringColumn("significant_delay_in_approval")
		SignificantDelayInApprovalCommentColumn             = mysql.StringColumn("significant_delay_in_approval_comment")
		DelayRelatedEventsConditionsColumn                  = mysql.BoolColumn("delay_related_events_conditions")
		ConfirmNecessityAdditionalAuditColumn               = mysql.BoolColumn("confirm_necessity_additional_audit")
		ConfirmNecessityAdditionalAuditCommentColumn        = mysql.StringColumn("confirm_necessity_additional_audit_comment")
		IsValidColumn                                       = mysql.BoolColumn("is_valid")
		ConclusionImpactOnMaterialUncertaintyColumn         = mysql.BoolColumn("conclusion_impact_on_material_uncertainty")
		DelayRelatedEventsConditionsCommentColumn           = mysql.StringColumn("delay_related_events_conditions_comment")
		allColumns                                          = mysql.ColumnList{IDColumn, EngagementIDColumn, NodeIDColumn, EntityStatusColumn, AccountingBasisChangeColumn, AccountingBasisPlanningStageColumn, AccountingBasisCompletionStageColumn, AccountingBasisCompletionStageCommentsColumn, GoingConcernAssessmentColumn, GoingConcernAssessmentCommentsColumn, MaterialUncertaintiesColumn, MaterialUncertaintiesCommentsColumn, SituationColumn, ImplicationsReportColumn, IsSituationColumn, ConfirmCompliedApplicationColumn, ConfirmCompliedImplicationCommentColumn, ExtendAssessmentColumn, ConsiderImplicationsColumn, EffectOnUnableObtainSufficientEvidenceColumn, EffectOnUnableObtainSufficientEvidenceCommentColumn, SignificantDelayInApprovalColumn, SignificantDelayInApprovalCommentColumn, DelayRelatedEventsConditionsColumn, ConfirmNecessityAdditionalAuditColumn, ConfirmNecessityAdditionalAuditCommentColumn, IsValidColumn, ConclusionImpactOnMaterialUncertaintyColumn, DelayRelatedEventsConditionsCommentColumn}
		mutableColumns                                      = mysql.ColumnList{EngagementIDColumn, NodeIDColumn, EntityStatusColumn, AccountingBasisChangeColumn, AccountingBasisPlanningStageColumn, AccountingBasisCompletionStageColumn, AccountingBasisCompletionStageCommentsColumn, GoingConcernAssessmentColumn, GoingConcernAssessmentCommentsColumn, MaterialUncertaintiesColumn, MaterialUncertaintiesCommentsColumn, SituationColumn, ImplicationsReportColumn, IsSituationColumn, ConfirmCompliedApplicationColumn, ConfirmCompliedImplicationCommentColumn, ExtendAssessmentColumn, ConsiderImplicationsColumn, EffectOnUnableObtainSufficientEvidenceColumn, EffectOnUnableObtainSufficientEvidenceCommentColumn, SignificantDelayInApprovalColumn, SignificantDelayInApprovalCommentColumn, DelayRelatedEventsConditionsColumn, ConfirmNecessityAdditionalAuditColumn, ConfirmNecessityAdditionalAuditCommentColumn, IsValidColumn, ConclusionImpactOnMaterialUncertaintyColumn, DelayRelatedEventsConditionsCommentColumn}
	)

	return completionGoingConcernFinalConclusionTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                                     IDColumn,
		EngagementID:                           EngagementIDColumn,
		NodeID:                                 NodeIDColumn,
		EntityStatus:                           EntityStatusColumn,
		AccountingBasisChange:                  AccountingBasisChangeColumn,
		AccountingBasisPlanningStage:           AccountingBasisPlanningStageColumn,
		AccountingBasisCompletionStage:         AccountingBasisCompletionStageColumn,
		AccountingBasisCompletionStageComments: AccountingBasisCompletionStageCommentsColumn,
		GoingConcernAssessment:                 GoingConcernAssessmentColumn,
		GoingConcernAssessmentComments:         GoingConcernAssessmentCommentsColumn,
		MaterialUncertainties:                  MaterialUncertaintiesColumn,
		MaterialUncertaintiesComments:          MaterialUncertaintiesCommentsColumn,
		Situation:                              SituationColumn,
		ImplicationsReport:                     ImplicationsReportColumn,
		IsSituation:                            IsSituationColumn,
		ConfirmCompliedApplication:             ConfirmCompliedApplicationColumn,
		ConfirmCompliedImplicationComment:      ConfirmCompliedImplicationCommentColumn,
		ExtendAssessment:                       ExtendAssessmentColumn,
		ConsiderImplications:                   ConsiderImplicationsColumn,
		EffectOnUnableObtainSufficientEvidence: EffectOnUnableObtainSufficientEvidenceColumn,
		EffectOnUnableObtainSufficientEvidenceComment: EffectOnUnableObtainSufficientEvidenceCommentColumn,
		SignificantDelayInApproval:                    SignificantDelayInApprovalColumn,
		SignificantDelayInApprovalComment:             SignificantDelayInApprovalCommentColumn,
		DelayRelatedEventsConditions:                  DelayRelatedEventsConditionsColumn,
		ConfirmNecessityAdditionalAudit:               ConfirmNecessityAdditionalAuditColumn,
		ConfirmNecessityAdditionalAuditComment:        ConfirmNecessityAdditionalAuditCommentColumn,
		IsValid:                                       IsValidColumn,
		ConclusionImpactOnMaterialUncertainty:         ConclusionImpactOnMaterialUncertaintyColumn,
		DelayRelatedEventsConditionsComment:           DelayRelatedEventsConditionsCommentColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

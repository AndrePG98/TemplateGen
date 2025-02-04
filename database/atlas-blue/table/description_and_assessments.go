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

var DescriptionAndAssessments = newDescriptionAndAssessmentsTable("atlas-blue", "description_and_assessments", "")

type descriptionAndAssessmentsTable struct {
	mysql.Table

	// Columns
	ID                                            mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	EngagementID                                  mysql.ColumnString // (DC2Type:guid)
	CreatedBy                                     mysql.ColumnString // (DC2Type:guid)
	EvaluatedAccountingPolicies                   mysql.ColumnString
	RiskDescription                               mysql.ColumnString
	ExpertRequired                                mysql.ColumnBool
	RiskAssessment                                mysql.ColumnString
	NeedOtherAuditProcedures                      mysql.ColumnBool
	Comment                                       mysql.ColumnString
	Status                                        mysql.ColumnString
	EffectOnTheFinancialStatement                 mysql.ColumnString
	DoesTheManagementOverrideAffectsTheAssessment mysql.ColumnBool
	HowTheManagementOverrideAffectsTheAssessment  mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type DescriptionAndAssessmentsTable struct {
	descriptionAndAssessmentsTable

	NEW descriptionAndAssessmentsTable
}

// AS creates new DescriptionAndAssessmentsTable with assigned alias
func (a DescriptionAndAssessmentsTable) AS(alias string) *DescriptionAndAssessmentsTable {
	return newDescriptionAndAssessmentsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new DescriptionAndAssessmentsTable with assigned schema name
func (a DescriptionAndAssessmentsTable) FromSchema(schemaName string) *DescriptionAndAssessmentsTable {
	return newDescriptionAndAssessmentsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new DescriptionAndAssessmentsTable with assigned table prefix
func (a DescriptionAndAssessmentsTable) WithPrefix(prefix string) *DescriptionAndAssessmentsTable {
	return newDescriptionAndAssessmentsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new DescriptionAndAssessmentsTable with assigned table suffix
func (a DescriptionAndAssessmentsTable) WithSuffix(suffix string) *DescriptionAndAssessmentsTable {
	return newDescriptionAndAssessmentsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newDescriptionAndAssessmentsTable(schemaName, tableName, alias string) *DescriptionAndAssessmentsTable {
	return &DescriptionAndAssessmentsTable{
		descriptionAndAssessmentsTable: newDescriptionAndAssessmentsTableImpl(schemaName, tableName, alias),
		NEW:                            newDescriptionAndAssessmentsTableImpl("", "new", ""),
	}
}

func newDescriptionAndAssessmentsTableImpl(schemaName, tableName, alias string) descriptionAndAssessmentsTable {
	var (
		IDColumn                                            = mysql.StringColumn("id")
		EngagementIDColumn                                  = mysql.StringColumn("engagement_id")
		CreatedByColumn                                     = mysql.StringColumn("created_by")
		EvaluatedAccountingPoliciesColumn                   = mysql.StringColumn("evaluated_accounting_policies")
		RiskDescriptionColumn                               = mysql.StringColumn("risk_description")
		ExpertRequiredColumn                                = mysql.BoolColumn("expert_required")
		RiskAssessmentColumn                                = mysql.StringColumn("risk_assessment")
		NeedOtherAuditProceduresColumn                      = mysql.BoolColumn("need_other_audit_procedures")
		CommentColumn                                       = mysql.StringColumn("comment")
		StatusColumn                                        = mysql.StringColumn("status")
		EffectOnTheFinancialStatementColumn                 = mysql.StringColumn("effect_on_the_financial_statement")
		DoesTheManagementOverrideAffectsTheAssessmentColumn = mysql.BoolColumn("does_the_management_override_affects_the_assessment")
		HowTheManagementOverrideAffectsTheAssessmentColumn  = mysql.StringColumn("how_the_management_override_affects_the_assessment")
		allColumns                                          = mysql.ColumnList{IDColumn, EngagementIDColumn, CreatedByColumn, EvaluatedAccountingPoliciesColumn, RiskDescriptionColumn, ExpertRequiredColumn, RiskAssessmentColumn, NeedOtherAuditProceduresColumn, CommentColumn, StatusColumn, EffectOnTheFinancialStatementColumn, DoesTheManagementOverrideAffectsTheAssessmentColumn, HowTheManagementOverrideAffectsTheAssessmentColumn}
		mutableColumns                                      = mysql.ColumnList{EngagementIDColumn, CreatedByColumn, EvaluatedAccountingPoliciesColumn, RiskDescriptionColumn, ExpertRequiredColumn, RiskAssessmentColumn, NeedOtherAuditProceduresColumn, CommentColumn, StatusColumn, EffectOnTheFinancialStatementColumn, DoesTheManagementOverrideAffectsTheAssessmentColumn, HowTheManagementOverrideAffectsTheAssessmentColumn}
	)

	return descriptionAndAssessmentsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                            IDColumn,
		EngagementID:                  EngagementIDColumn,
		CreatedBy:                     CreatedByColumn,
		EvaluatedAccountingPolicies:   EvaluatedAccountingPoliciesColumn,
		RiskDescription:               RiskDescriptionColumn,
		ExpertRequired:                ExpertRequiredColumn,
		RiskAssessment:                RiskAssessmentColumn,
		NeedOtherAuditProcedures:      NeedOtherAuditProceduresColumn,
		Comment:                       CommentColumn,
		Status:                        StatusColumn,
		EffectOnTheFinancialStatement: EffectOnTheFinancialStatementColumn,
		DoesTheManagementOverrideAffectsTheAssessment: DoesTheManagementOverrideAffectsTheAssessmentColumn,
		HowTheManagementOverrideAffectsTheAssessment:  HowTheManagementOverrideAffectsTheAssessmentColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

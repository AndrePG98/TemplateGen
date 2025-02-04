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

var FraudUnderstanding = newFraudUnderstandingTable("atlas-blue", "fraud_understanding", "")

type fraudUnderstandingTable struct {
	mysql.Table

	// Columns
	ID                                                mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	EngagementID                                      mysql.ColumnString // (DC2Type:guid)
	NodeID                                            mysql.ColumnString
	IsValid                                           mysql.ColumnBool
	InquiresWithTcwgConfirmed                         mysql.ColumnBool
	InquiresWithTcwgConfirmComment                    mysql.ColumnString
	InquiresWithManagementConfirmed                   mysql.ColumnBool
	InquiresWithManagementComment                     mysql.ColumnString
	IncludeAuditDocumentationConfirmed                mysql.ColumnBool
	IncludeAuditDocumentationComment                  mysql.ColumnString
	WillMaintainProfessionalScepticismConfirmed       mysql.ColumnBool
	WillMaintainProfessionalScepticismComment         mysql.ColumnString
	HaveObtainedChargeUnderstandingConfirmed          mysql.ColumnBool
	HaveObtainedChargeUnderstandingComment            mysql.ColumnString
	OtherRiskAssessmentInformationEvaluationConfirmed mysql.ColumnBool
	OtherRiskAssessmentInformationEvaluationComment   mysql.ColumnString
	IsThereRiskOfMaterialMisstatement                 mysql.ColumnBool
	RebuttalOfPresumptionOfRisk                       mysql.ColumnString
	TypesOfRevenuesEvaluatedConfirmed                 mysql.ColumnBool
	TypesOfRevenuesEvaluatedComment                   mysql.ColumnString
	CommunicationToCountryRiskManager                 mysql.ColumnBool
	OtherRisksOrFraudIdentified                       mysql.ColumnBool

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FraudUnderstandingTable struct {
	fraudUnderstandingTable

	NEW fraudUnderstandingTable
}

// AS creates new FraudUnderstandingTable with assigned alias
func (a FraudUnderstandingTable) AS(alias string) *FraudUnderstandingTable {
	return newFraudUnderstandingTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FraudUnderstandingTable with assigned schema name
func (a FraudUnderstandingTable) FromSchema(schemaName string) *FraudUnderstandingTable {
	return newFraudUnderstandingTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FraudUnderstandingTable with assigned table prefix
func (a FraudUnderstandingTable) WithPrefix(prefix string) *FraudUnderstandingTable {
	return newFraudUnderstandingTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FraudUnderstandingTable with assigned table suffix
func (a FraudUnderstandingTable) WithSuffix(suffix string) *FraudUnderstandingTable {
	return newFraudUnderstandingTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFraudUnderstandingTable(schemaName, tableName, alias string) *FraudUnderstandingTable {
	return &FraudUnderstandingTable{
		fraudUnderstandingTable: newFraudUnderstandingTableImpl(schemaName, tableName, alias),
		NEW:                     newFraudUnderstandingTableImpl("", "new", ""),
	}
}

func newFraudUnderstandingTableImpl(schemaName, tableName, alias string) fraudUnderstandingTable {
	var (
		IDColumn                                                = mysql.StringColumn("id")
		EngagementIDColumn                                      = mysql.StringColumn("engagement_id")
		NodeIDColumn                                            = mysql.StringColumn("node_id")
		IsValidColumn                                           = mysql.BoolColumn("is_valid")
		InquiresWithTcwgConfirmedColumn                         = mysql.BoolColumn("inquires_with_tcwg_confirmed")
		InquiresWithTcwgConfirmCommentColumn                    = mysql.StringColumn("inquires_with_tcwg_confirm_comment")
		InquiresWithManagementConfirmedColumn                   = mysql.BoolColumn("inquires_with_management_confirmed")
		InquiresWithManagementCommentColumn                     = mysql.StringColumn("inquires_with_management_comment")
		IncludeAuditDocumentationConfirmedColumn                = mysql.BoolColumn("include_audit_documentation_confirmed")
		IncludeAuditDocumentationCommentColumn                  = mysql.StringColumn("include_audit_documentation_comment")
		WillMaintainProfessionalScepticismConfirmedColumn       = mysql.BoolColumn("will_maintain_professional_scepticism_confirmed")
		WillMaintainProfessionalScepticismCommentColumn         = mysql.StringColumn("will_maintain_professional_scepticism_comment")
		HaveObtainedChargeUnderstandingConfirmedColumn          = mysql.BoolColumn("have_obtained_charge_understanding_confirmed")
		HaveObtainedChargeUnderstandingCommentColumn            = mysql.StringColumn("have_obtained_charge_understanding_comment")
		OtherRiskAssessmentInformationEvaluationConfirmedColumn = mysql.BoolColumn("other_risk_assessment_information_evaluation_confirmed")
		OtherRiskAssessmentInformationEvaluationCommentColumn   = mysql.StringColumn("other_risk_assessment_information_evaluation_comment")
		IsThereRiskOfMaterialMisstatementColumn                 = mysql.BoolColumn("is_there_risk_of_material_misstatement")
		RebuttalOfPresumptionOfRiskColumn                       = mysql.StringColumn("rebuttal_of_presumption_of_risk")
		TypesOfRevenuesEvaluatedConfirmedColumn                 = mysql.BoolColumn("types_of_revenues_evaluated_confirmed")
		TypesOfRevenuesEvaluatedCommentColumn                   = mysql.StringColumn("types_of_revenues_evaluated_comment")
		CommunicationToCountryRiskManagerColumn                 = mysql.BoolColumn("communication_to_country_risk_manager")
		OtherRisksOrFraudIdentifiedColumn                       = mysql.BoolColumn("other_risks_or_fraud_identified")
		allColumns                                              = mysql.ColumnList{IDColumn, EngagementIDColumn, NodeIDColumn, IsValidColumn, InquiresWithTcwgConfirmedColumn, InquiresWithTcwgConfirmCommentColumn, InquiresWithManagementConfirmedColumn, InquiresWithManagementCommentColumn, IncludeAuditDocumentationConfirmedColumn, IncludeAuditDocumentationCommentColumn, WillMaintainProfessionalScepticismConfirmedColumn, WillMaintainProfessionalScepticismCommentColumn, HaveObtainedChargeUnderstandingConfirmedColumn, HaveObtainedChargeUnderstandingCommentColumn, OtherRiskAssessmentInformationEvaluationConfirmedColumn, OtherRiskAssessmentInformationEvaluationCommentColumn, IsThereRiskOfMaterialMisstatementColumn, RebuttalOfPresumptionOfRiskColumn, TypesOfRevenuesEvaluatedConfirmedColumn, TypesOfRevenuesEvaluatedCommentColumn, CommunicationToCountryRiskManagerColumn, OtherRisksOrFraudIdentifiedColumn}
		mutableColumns                                          = mysql.ColumnList{EngagementIDColumn, NodeIDColumn, IsValidColumn, InquiresWithTcwgConfirmedColumn, InquiresWithTcwgConfirmCommentColumn, InquiresWithManagementConfirmedColumn, InquiresWithManagementCommentColumn, IncludeAuditDocumentationConfirmedColumn, IncludeAuditDocumentationCommentColumn, WillMaintainProfessionalScepticismConfirmedColumn, WillMaintainProfessionalScepticismCommentColumn, HaveObtainedChargeUnderstandingConfirmedColumn, HaveObtainedChargeUnderstandingCommentColumn, OtherRiskAssessmentInformationEvaluationConfirmedColumn, OtherRiskAssessmentInformationEvaluationCommentColumn, IsThereRiskOfMaterialMisstatementColumn, RebuttalOfPresumptionOfRiskColumn, TypesOfRevenuesEvaluatedConfirmedColumn, TypesOfRevenuesEvaluatedCommentColumn, CommunicationToCountryRiskManagerColumn, OtherRisksOrFraudIdentifiedColumn}
	)

	return fraudUnderstandingTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                                 IDColumn,
		EngagementID:                       EngagementIDColumn,
		NodeID:                             NodeIDColumn,
		IsValid:                            IsValidColumn,
		InquiresWithTcwgConfirmed:          InquiresWithTcwgConfirmedColumn,
		InquiresWithTcwgConfirmComment:     InquiresWithTcwgConfirmCommentColumn,
		InquiresWithManagementConfirmed:    InquiresWithManagementConfirmedColumn,
		InquiresWithManagementComment:      InquiresWithManagementCommentColumn,
		IncludeAuditDocumentationConfirmed: IncludeAuditDocumentationConfirmedColumn,
		IncludeAuditDocumentationComment:   IncludeAuditDocumentationCommentColumn,
		WillMaintainProfessionalScepticismConfirmed:       WillMaintainProfessionalScepticismConfirmedColumn,
		WillMaintainProfessionalScepticismComment:         WillMaintainProfessionalScepticismCommentColumn,
		HaveObtainedChargeUnderstandingConfirmed:          HaveObtainedChargeUnderstandingConfirmedColumn,
		HaveObtainedChargeUnderstandingComment:            HaveObtainedChargeUnderstandingCommentColumn,
		OtherRiskAssessmentInformationEvaluationConfirmed: OtherRiskAssessmentInformationEvaluationConfirmedColumn,
		OtherRiskAssessmentInformationEvaluationComment:   OtherRiskAssessmentInformationEvaluationCommentColumn,
		IsThereRiskOfMaterialMisstatement:                 IsThereRiskOfMaterialMisstatementColumn,
		RebuttalOfPresumptionOfRisk:                       RebuttalOfPresumptionOfRiskColumn,
		TypesOfRevenuesEvaluatedConfirmed:                 TypesOfRevenuesEvaluatedConfirmedColumn,
		TypesOfRevenuesEvaluatedComment:                   TypesOfRevenuesEvaluatedCommentColumn,
		CommunicationToCountryRiskManager:                 CommunicationToCountryRiskManagerColumn,
		OtherRisksOrFraudIdentified:                       OtherRisksOrFraudIdentifiedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

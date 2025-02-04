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

var EngagementEntityEnvironments = newEngagementEntityEnvironmentsTable("atlas-blue", "engagement_entity_environments", "")

type engagementEntityEnvironmentsTable struct {
	mysql.Table

	// Columns
	ID                                                 mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	EngagementID                                       mysql.ColumnString // (DC2Type:guid)
	RiskConclusion                                     mysql.ColumnString
	RelevantInformation                                mysql.ColumnString
	OrganizationalStructure                            mysql.ColumnString
	OwnershipGovernance                                mysql.ColumnString
	InvestmentType                                     mysql.ColumnString
	BusinessModel                                      mysql.ColumnString
	IndustryFactors                                    mysql.ColumnString
	RegulatoryFactors                                  mysql.ColumnString
	ExternalFactors                                    mysql.ColumnString
	Measures                                           mysql.ColumnString
	IsValid                                            mysql.ColumnBool
	ChargedWithGovernanceInvolved                      mysql.ColumnBool
	ConsiderAnyInformationObtainedFromOtherEngagements mysql.ColumnString
	GovernanceInvolvedComment                          mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type EngagementEntityEnvironmentsTable struct {
	engagementEntityEnvironmentsTable

	NEW engagementEntityEnvironmentsTable
}

// AS creates new EngagementEntityEnvironmentsTable with assigned alias
func (a EngagementEntityEnvironmentsTable) AS(alias string) *EngagementEntityEnvironmentsTable {
	return newEngagementEntityEnvironmentsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new EngagementEntityEnvironmentsTable with assigned schema name
func (a EngagementEntityEnvironmentsTable) FromSchema(schemaName string) *EngagementEntityEnvironmentsTable {
	return newEngagementEntityEnvironmentsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new EngagementEntityEnvironmentsTable with assigned table prefix
func (a EngagementEntityEnvironmentsTable) WithPrefix(prefix string) *EngagementEntityEnvironmentsTable {
	return newEngagementEntityEnvironmentsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new EngagementEntityEnvironmentsTable with assigned table suffix
func (a EngagementEntityEnvironmentsTable) WithSuffix(suffix string) *EngagementEntityEnvironmentsTable {
	return newEngagementEntityEnvironmentsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newEngagementEntityEnvironmentsTable(schemaName, tableName, alias string) *EngagementEntityEnvironmentsTable {
	return &EngagementEntityEnvironmentsTable{
		engagementEntityEnvironmentsTable: newEngagementEntityEnvironmentsTableImpl(schemaName, tableName, alias),
		NEW:                               newEngagementEntityEnvironmentsTableImpl("", "new", ""),
	}
}

func newEngagementEntityEnvironmentsTableImpl(schemaName, tableName, alias string) engagementEntityEnvironmentsTable {
	var (
		IDColumn                                                 = mysql.StringColumn("id")
		EngagementIDColumn                                       = mysql.StringColumn("engagement_id")
		RiskConclusionColumn                                     = mysql.StringColumn("risk_conclusion")
		RelevantInformationColumn                                = mysql.StringColumn("relevant_information")
		OrganizationalStructureColumn                            = mysql.StringColumn("organizational_structure")
		OwnershipGovernanceColumn                                = mysql.StringColumn("ownership_governance")
		InvestmentTypeColumn                                     = mysql.StringColumn("investment_type")
		BusinessModelColumn                                      = mysql.StringColumn("business_model")
		IndustryFactorsColumn                                    = mysql.StringColumn("industry_factors")
		RegulatoryFactorsColumn                                  = mysql.StringColumn("regulatory_factors")
		ExternalFactorsColumn                                    = mysql.StringColumn("external_factors")
		MeasuresColumn                                           = mysql.StringColumn("measures")
		IsValidColumn                                            = mysql.BoolColumn("is_valid")
		ChargedWithGovernanceInvolvedColumn                      = mysql.BoolColumn("charged_with_governance_involved")
		ConsiderAnyInformationObtainedFromOtherEngagementsColumn = mysql.StringColumn("consider_any_information_obtained_from_other_engagements")
		GovernanceInvolvedCommentColumn                          = mysql.StringColumn("governance_involved_comment")
		allColumns                                               = mysql.ColumnList{IDColumn, EngagementIDColumn, RiskConclusionColumn, RelevantInformationColumn, OrganizationalStructureColumn, OwnershipGovernanceColumn, InvestmentTypeColumn, BusinessModelColumn, IndustryFactorsColumn, RegulatoryFactorsColumn, ExternalFactorsColumn, MeasuresColumn, IsValidColumn, ChargedWithGovernanceInvolvedColumn, ConsiderAnyInformationObtainedFromOtherEngagementsColumn, GovernanceInvolvedCommentColumn}
		mutableColumns                                           = mysql.ColumnList{EngagementIDColumn, RiskConclusionColumn, RelevantInformationColumn, OrganizationalStructureColumn, OwnershipGovernanceColumn, InvestmentTypeColumn, BusinessModelColumn, IndustryFactorsColumn, RegulatoryFactorsColumn, ExternalFactorsColumn, MeasuresColumn, IsValidColumn, ChargedWithGovernanceInvolvedColumn, ConsiderAnyInformationObtainedFromOtherEngagementsColumn, GovernanceInvolvedCommentColumn}
	)

	return engagementEntityEnvironmentsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                            IDColumn,
		EngagementID:                  EngagementIDColumn,
		RiskConclusion:                RiskConclusionColumn,
		RelevantInformation:           RelevantInformationColumn,
		OrganizationalStructure:       OrganizationalStructureColumn,
		OwnershipGovernance:           OwnershipGovernanceColumn,
		InvestmentType:                InvestmentTypeColumn,
		BusinessModel:                 BusinessModelColumn,
		IndustryFactors:               IndustryFactorsColumn,
		RegulatoryFactors:             RegulatoryFactorsColumn,
		ExternalFactors:               ExternalFactorsColumn,
		Measures:                      MeasuresColumn,
		IsValid:                       IsValidColumn,
		ChargedWithGovernanceInvolved: ChargedWithGovernanceInvolvedColumn,
		ConsiderAnyInformationObtainedFromOtherEngagements: ConsiderAnyInformationObtainedFromOtherEngagementsColumn,
		GovernanceInvolvedComment:                          GovernanceInvolvedCommentColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

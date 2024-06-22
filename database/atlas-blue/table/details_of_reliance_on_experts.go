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

var DetailsOfRelianceOnExperts = newDetailsOfRelianceOnExpertsTable("atlas-blue", "details_of_reliance_on_experts", "")

type detailsOfRelianceOnExpertsTable struct {
	mysql.Table

	// Columns
	ID                                          mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	EngagementID                                mysql.ColumnString // (DC2Type:guid)
	WillYouRelyOnTheWorkOfAnAuditor             mysql.ColumnBool
	HasInformationThatIsToBeUsedAsAuditEvidence mysql.ColumnBool

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type DetailsOfRelianceOnExpertsTable struct {
	detailsOfRelianceOnExpertsTable

	NEW detailsOfRelianceOnExpertsTable
}

// AS creates new DetailsOfRelianceOnExpertsTable with assigned alias
func (a DetailsOfRelianceOnExpertsTable) AS(alias string) *DetailsOfRelianceOnExpertsTable {
	return newDetailsOfRelianceOnExpertsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new DetailsOfRelianceOnExpertsTable with assigned schema name
func (a DetailsOfRelianceOnExpertsTable) FromSchema(schemaName string) *DetailsOfRelianceOnExpertsTable {
	return newDetailsOfRelianceOnExpertsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new DetailsOfRelianceOnExpertsTable with assigned table prefix
func (a DetailsOfRelianceOnExpertsTable) WithPrefix(prefix string) *DetailsOfRelianceOnExpertsTable {
	return newDetailsOfRelianceOnExpertsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new DetailsOfRelianceOnExpertsTable with assigned table suffix
func (a DetailsOfRelianceOnExpertsTable) WithSuffix(suffix string) *DetailsOfRelianceOnExpertsTable {
	return newDetailsOfRelianceOnExpertsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newDetailsOfRelianceOnExpertsTable(schemaName, tableName, alias string) *DetailsOfRelianceOnExpertsTable {
	return &DetailsOfRelianceOnExpertsTable{
		detailsOfRelianceOnExpertsTable: newDetailsOfRelianceOnExpertsTableImpl(schemaName, tableName, alias),
		NEW:                             newDetailsOfRelianceOnExpertsTableImpl("", "new", ""),
	}
}

func newDetailsOfRelianceOnExpertsTableImpl(schemaName, tableName, alias string) detailsOfRelianceOnExpertsTable {
	var (
		IDColumn                                          = mysql.StringColumn("id")
		EngagementIDColumn                                = mysql.StringColumn("engagement_id")
		WillYouRelyOnTheWorkOfAnAuditorColumn             = mysql.BoolColumn("will_you_rely_on_the_work_of_an_auditor")
		HasInformationThatIsToBeUsedAsAuditEvidenceColumn = mysql.BoolColumn("has_information_that_is_to_be_used_as_audit_evidence")
		allColumns                                        = mysql.ColumnList{IDColumn, EngagementIDColumn, WillYouRelyOnTheWorkOfAnAuditorColumn, HasInformationThatIsToBeUsedAsAuditEvidenceColumn}
		mutableColumns                                    = mysql.ColumnList{EngagementIDColumn, WillYouRelyOnTheWorkOfAnAuditorColumn, HasInformationThatIsToBeUsedAsAuditEvidenceColumn}
	)

	return detailsOfRelianceOnExpertsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                              IDColumn,
		EngagementID:                    EngagementIDColumn,
		WillYouRelyOnTheWorkOfAnAuditor: WillYouRelyOnTheWorkOfAnAuditorColumn,
		HasInformationThatIsToBeUsedAsAuditEvidence: HasInformationThatIsToBeUsedAsAuditEvidenceColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

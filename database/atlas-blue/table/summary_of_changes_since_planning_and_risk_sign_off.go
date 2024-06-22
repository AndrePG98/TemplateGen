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

var SummaryOfChangesSincePlanningAndRiskSignOff = newSummaryOfChangesSincePlanningAndRiskSignOffTable("atlas-blue", "summary_of_changes_since_planning_and_risk_sign_off", "")

type summaryOfChangesSincePlanningAndRiskSignOffTable struct {
	mysql.Table

	// Columns
	ID           mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	EngagementID mysql.ColumnString // (DC2Type:guid)
	IsValid      mysql.ColumnBool
	AuditTrailID mysql.ColumnString // (DC2Type:guid)
	ChangeReason mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type SummaryOfChangesSincePlanningAndRiskSignOffTable struct {
	summaryOfChangesSincePlanningAndRiskSignOffTable

	NEW summaryOfChangesSincePlanningAndRiskSignOffTable
}

// AS creates new SummaryOfChangesSincePlanningAndRiskSignOffTable with assigned alias
func (a SummaryOfChangesSincePlanningAndRiskSignOffTable) AS(alias string) *SummaryOfChangesSincePlanningAndRiskSignOffTable {
	return newSummaryOfChangesSincePlanningAndRiskSignOffTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SummaryOfChangesSincePlanningAndRiskSignOffTable with assigned schema name
func (a SummaryOfChangesSincePlanningAndRiskSignOffTable) FromSchema(schemaName string) *SummaryOfChangesSincePlanningAndRiskSignOffTable {
	return newSummaryOfChangesSincePlanningAndRiskSignOffTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SummaryOfChangesSincePlanningAndRiskSignOffTable with assigned table prefix
func (a SummaryOfChangesSincePlanningAndRiskSignOffTable) WithPrefix(prefix string) *SummaryOfChangesSincePlanningAndRiskSignOffTable {
	return newSummaryOfChangesSincePlanningAndRiskSignOffTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SummaryOfChangesSincePlanningAndRiskSignOffTable with assigned table suffix
func (a SummaryOfChangesSincePlanningAndRiskSignOffTable) WithSuffix(suffix string) *SummaryOfChangesSincePlanningAndRiskSignOffTable {
	return newSummaryOfChangesSincePlanningAndRiskSignOffTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSummaryOfChangesSincePlanningAndRiskSignOffTable(schemaName, tableName, alias string) *SummaryOfChangesSincePlanningAndRiskSignOffTable {
	return &SummaryOfChangesSincePlanningAndRiskSignOffTable{
		summaryOfChangesSincePlanningAndRiskSignOffTable: newSummaryOfChangesSincePlanningAndRiskSignOffTableImpl(schemaName, tableName, alias),
		NEW: newSummaryOfChangesSincePlanningAndRiskSignOffTableImpl("", "new", ""),
	}
}

func newSummaryOfChangesSincePlanningAndRiskSignOffTableImpl(schemaName, tableName, alias string) summaryOfChangesSincePlanningAndRiskSignOffTable {
	var (
		IDColumn           = mysql.StringColumn("id")
		EngagementIDColumn = mysql.StringColumn("engagement_id")
		IsValidColumn      = mysql.BoolColumn("is_valid")
		AuditTrailIDColumn = mysql.StringColumn("audit_trail_id")
		ChangeReasonColumn = mysql.StringColumn("change_reason")
		allColumns         = mysql.ColumnList{IDColumn, EngagementIDColumn, IsValidColumn, AuditTrailIDColumn, ChangeReasonColumn}
		mutableColumns     = mysql.ColumnList{EngagementIDColumn, IsValidColumn, AuditTrailIDColumn, ChangeReasonColumn}
	)

	return summaryOfChangesSincePlanningAndRiskSignOffTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		EngagementID: EngagementIDColumn,
		IsValid:      IsValidColumn,
		AuditTrailID: AuditTrailIDColumn,
		ChangeReason: ChangeReasonColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

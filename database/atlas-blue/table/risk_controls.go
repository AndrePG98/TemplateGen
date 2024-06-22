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

var RiskControls = newRiskControlsTable("atlas-blue", "risk_controls", "")

type riskControlsTable struct {
	mysql.Table

	// Columns
	ID                    mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	EngagementID          mysql.ColumnString // (DC2Type:guid)
	Name                  mysql.ColumnString
	SourceID              mysql.ColumnString
	Label                 mysql.ColumnString // (DC2Type:json)
	DateLastLabelModified mysql.ColumnString // (DC2Type:json)

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type RiskControlsTable struct {
	riskControlsTable

	NEW riskControlsTable
}

// AS creates new RiskControlsTable with assigned alias
func (a RiskControlsTable) AS(alias string) *RiskControlsTable {
	return newRiskControlsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new RiskControlsTable with assigned schema name
func (a RiskControlsTable) FromSchema(schemaName string) *RiskControlsTable {
	return newRiskControlsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new RiskControlsTable with assigned table prefix
func (a RiskControlsTable) WithPrefix(prefix string) *RiskControlsTable {
	return newRiskControlsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new RiskControlsTable with assigned table suffix
func (a RiskControlsTable) WithSuffix(suffix string) *RiskControlsTable {
	return newRiskControlsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newRiskControlsTable(schemaName, tableName, alias string) *RiskControlsTable {
	return &RiskControlsTable{
		riskControlsTable: newRiskControlsTableImpl(schemaName, tableName, alias),
		NEW:               newRiskControlsTableImpl("", "new", ""),
	}
}

func newRiskControlsTableImpl(schemaName, tableName, alias string) riskControlsTable {
	var (
		IDColumn                    = mysql.StringColumn("id")
		EngagementIDColumn          = mysql.StringColumn("engagement_id")
		NameColumn                  = mysql.StringColumn("name")
		SourceIDColumn              = mysql.StringColumn("source_id")
		LabelColumn                 = mysql.StringColumn("label")
		DateLastLabelModifiedColumn = mysql.StringColumn("date_last_label_modified")
		allColumns                  = mysql.ColumnList{IDColumn, EngagementIDColumn, NameColumn, SourceIDColumn, LabelColumn, DateLastLabelModifiedColumn}
		mutableColumns              = mysql.ColumnList{EngagementIDColumn, NameColumn, SourceIDColumn, LabelColumn, DateLastLabelModifiedColumn}
	)

	return riskControlsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                    IDColumn,
		EngagementID:          EngagementIDColumn,
		Name:                  NameColumn,
		SourceID:              SourceIDColumn,
		Label:                 LabelColumn,
		DateLastLabelModified: DateLastLabelModifiedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

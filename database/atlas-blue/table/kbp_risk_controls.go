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

var KbpRiskControls = newKbpRiskControlsTable("atlas-blue", "kbp_risk_controls", "")

type kbpRiskControlsTable struct {
	mysql.Table

	// Columns
	ID        mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	KbpID     mysql.ColumnString // (DC2Type:guid)
	ControlID mysql.ColumnString // (DC2Type:guid)
	RiskID    mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type KbpRiskControlsTable struct {
	kbpRiskControlsTable

	NEW kbpRiskControlsTable
}

// AS creates new KbpRiskControlsTable with assigned alias
func (a KbpRiskControlsTable) AS(alias string) *KbpRiskControlsTable {
	return newKbpRiskControlsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new KbpRiskControlsTable with assigned schema name
func (a KbpRiskControlsTable) FromSchema(schemaName string) *KbpRiskControlsTable {
	return newKbpRiskControlsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new KbpRiskControlsTable with assigned table prefix
func (a KbpRiskControlsTable) WithPrefix(prefix string) *KbpRiskControlsTable {
	return newKbpRiskControlsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new KbpRiskControlsTable with assigned table suffix
func (a KbpRiskControlsTable) WithSuffix(suffix string) *KbpRiskControlsTable {
	return newKbpRiskControlsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newKbpRiskControlsTable(schemaName, tableName, alias string) *KbpRiskControlsTable {
	return &KbpRiskControlsTable{
		kbpRiskControlsTable: newKbpRiskControlsTableImpl(schemaName, tableName, alias),
		NEW:                  newKbpRiskControlsTableImpl("", "new", ""),
	}
}

func newKbpRiskControlsTableImpl(schemaName, tableName, alias string) kbpRiskControlsTable {
	var (
		IDColumn        = mysql.StringColumn("id")
		KbpIDColumn     = mysql.StringColumn("kbp_id")
		ControlIDColumn = mysql.StringColumn("control_id")
		RiskIDColumn    = mysql.StringColumn("risk_id")
		allColumns      = mysql.ColumnList{IDColumn, KbpIDColumn, ControlIDColumn, RiskIDColumn}
		mutableColumns  = mysql.ColumnList{KbpIDColumn, ControlIDColumn, RiskIDColumn}
	)

	return kbpRiskControlsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		KbpID:     KbpIDColumn,
		ControlID: ControlIDColumn,
		RiskID:    RiskIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

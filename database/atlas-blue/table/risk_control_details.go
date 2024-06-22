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

var RiskControlDetails = newRiskControlDetailsTable("atlas-blue", "risk_control_details", "")

type riskControlDetailsTable struct {
	mysql.Table

	// Columns
	ID                    mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	ControlID             mysql.ColumnString // (DC2Type:guid)
	ControlType           mysql.ColumnString
	ControlFrequency      mysql.ColumnString
	PreventOrDetect       mysql.ColumnString
	PlannedEvaluationDi   mysql.ColumnBool
	PlannedEvaluationOe   mysql.ColumnBool
	ItApplicationSelected mysql.ColumnBool
	ItApplicationID       mysql.ColumnString
	ControlLevel          mysql.ColumnString
	ControlDetailFormID   mysql.ColumnString // (DC2Type:guid)

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type RiskControlDetailsTable struct {
	riskControlDetailsTable

	NEW riskControlDetailsTable
}

// AS creates new RiskControlDetailsTable with assigned alias
func (a RiskControlDetailsTable) AS(alias string) *RiskControlDetailsTable {
	return newRiskControlDetailsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new RiskControlDetailsTable with assigned schema name
func (a RiskControlDetailsTable) FromSchema(schemaName string) *RiskControlDetailsTable {
	return newRiskControlDetailsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new RiskControlDetailsTable with assigned table prefix
func (a RiskControlDetailsTable) WithPrefix(prefix string) *RiskControlDetailsTable {
	return newRiskControlDetailsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new RiskControlDetailsTable with assigned table suffix
func (a RiskControlDetailsTable) WithSuffix(suffix string) *RiskControlDetailsTable {
	return newRiskControlDetailsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newRiskControlDetailsTable(schemaName, tableName, alias string) *RiskControlDetailsTable {
	return &RiskControlDetailsTable{
		riskControlDetailsTable: newRiskControlDetailsTableImpl(schemaName, tableName, alias),
		NEW:                     newRiskControlDetailsTableImpl("", "new", ""),
	}
}

func newRiskControlDetailsTableImpl(schemaName, tableName, alias string) riskControlDetailsTable {
	var (
		IDColumn                    = mysql.StringColumn("id")
		ControlIDColumn             = mysql.StringColumn("control_id")
		ControlTypeColumn           = mysql.StringColumn("control_type")
		ControlFrequencyColumn      = mysql.StringColumn("control_frequency")
		PreventOrDetectColumn       = mysql.StringColumn("prevent_or_detect")
		PlannedEvaluationDiColumn   = mysql.BoolColumn("planned_evaluation_di")
		PlannedEvaluationOeColumn   = mysql.BoolColumn("planned_evaluation_oe")
		ItApplicationSelectedColumn = mysql.BoolColumn("it_application_selected")
		ItApplicationIDColumn       = mysql.StringColumn("it_application_id")
		ControlLevelColumn          = mysql.StringColumn("control_level")
		ControlDetailFormIDColumn   = mysql.StringColumn("control_detail_form_id")
		allColumns                  = mysql.ColumnList{IDColumn, ControlIDColumn, ControlTypeColumn, ControlFrequencyColumn, PreventOrDetectColumn, PlannedEvaluationDiColumn, PlannedEvaluationOeColumn, ItApplicationSelectedColumn, ItApplicationIDColumn, ControlLevelColumn, ControlDetailFormIDColumn}
		mutableColumns              = mysql.ColumnList{ControlIDColumn, ControlTypeColumn, ControlFrequencyColumn, PreventOrDetectColumn, PlannedEvaluationDiColumn, PlannedEvaluationOeColumn, ItApplicationSelectedColumn, ItApplicationIDColumn, ControlLevelColumn, ControlDetailFormIDColumn}
	)

	return riskControlDetailsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                    IDColumn,
		ControlID:             ControlIDColumn,
		ControlType:           ControlTypeColumn,
		ControlFrequency:      ControlFrequencyColumn,
		PreventOrDetect:       PreventOrDetectColumn,
		PlannedEvaluationDi:   PlannedEvaluationDiColumn,
		PlannedEvaluationOe:   PlannedEvaluationOeColumn,
		ItApplicationSelected: ItApplicationSelectedColumn,
		ItApplicationID:       ItApplicationIDColumn,
		ControlLevel:          ControlLevelColumn,
		ControlDetailFormID:   ControlDetailFormIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

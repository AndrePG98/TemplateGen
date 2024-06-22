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

var ConclusionOnControlsFinancialStatementsLevelControls = newConclusionOnControlsFinancialStatementsLevelControlsTable("atlas-blue", "conclusion_on_controls_financial_statements_level_controls", "")

type conclusionOnControlsFinancialStatementsLevelControlsTable struct {
	mysql.Table

	// Columns
	ID                                        mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	ConclusionOnControlID                     mysql.ColumnString // (DC2Type:guid)
	FinancialStatementsLevelControlsControlID mysql.ColumnString // (DC2Type:guid)
	ControlRiskAfterEvaluatingOe              mysql.ColumnString
	ImpactOnAuditApproach                     mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ConclusionOnControlsFinancialStatementsLevelControlsTable struct {
	conclusionOnControlsFinancialStatementsLevelControlsTable

	NEW conclusionOnControlsFinancialStatementsLevelControlsTable
}

// AS creates new ConclusionOnControlsFinancialStatementsLevelControlsTable with assigned alias
func (a ConclusionOnControlsFinancialStatementsLevelControlsTable) AS(alias string) *ConclusionOnControlsFinancialStatementsLevelControlsTable {
	return newConclusionOnControlsFinancialStatementsLevelControlsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ConclusionOnControlsFinancialStatementsLevelControlsTable with assigned schema name
func (a ConclusionOnControlsFinancialStatementsLevelControlsTable) FromSchema(schemaName string) *ConclusionOnControlsFinancialStatementsLevelControlsTable {
	return newConclusionOnControlsFinancialStatementsLevelControlsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ConclusionOnControlsFinancialStatementsLevelControlsTable with assigned table prefix
func (a ConclusionOnControlsFinancialStatementsLevelControlsTable) WithPrefix(prefix string) *ConclusionOnControlsFinancialStatementsLevelControlsTable {
	return newConclusionOnControlsFinancialStatementsLevelControlsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ConclusionOnControlsFinancialStatementsLevelControlsTable with assigned table suffix
func (a ConclusionOnControlsFinancialStatementsLevelControlsTable) WithSuffix(suffix string) *ConclusionOnControlsFinancialStatementsLevelControlsTable {
	return newConclusionOnControlsFinancialStatementsLevelControlsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newConclusionOnControlsFinancialStatementsLevelControlsTable(schemaName, tableName, alias string) *ConclusionOnControlsFinancialStatementsLevelControlsTable {
	return &ConclusionOnControlsFinancialStatementsLevelControlsTable{
		conclusionOnControlsFinancialStatementsLevelControlsTable: newConclusionOnControlsFinancialStatementsLevelControlsTableImpl(schemaName, tableName, alias),
		NEW: newConclusionOnControlsFinancialStatementsLevelControlsTableImpl("", "new", ""),
	}
}

func newConclusionOnControlsFinancialStatementsLevelControlsTableImpl(schemaName, tableName, alias string) conclusionOnControlsFinancialStatementsLevelControlsTable {
	var (
		IDColumn                                        = mysql.StringColumn("id")
		ConclusionOnControlIDColumn                     = mysql.StringColumn("conclusion_on_control_id")
		FinancialStatementsLevelControlsControlIDColumn = mysql.StringColumn("financial_statements_level_controls_control_id")
		ControlRiskAfterEvaluatingOeColumn              = mysql.StringColumn("control_risk_after_evaluating_oe")
		ImpactOnAuditApproachColumn                     = mysql.StringColumn("impact_on_audit_approach")
		allColumns                                      = mysql.ColumnList{IDColumn, ConclusionOnControlIDColumn, FinancialStatementsLevelControlsControlIDColumn, ControlRiskAfterEvaluatingOeColumn, ImpactOnAuditApproachColumn}
		mutableColumns                                  = mysql.ColumnList{ConclusionOnControlIDColumn, FinancialStatementsLevelControlsControlIDColumn, ControlRiskAfterEvaluatingOeColumn, ImpactOnAuditApproachColumn}
	)

	return conclusionOnControlsFinancialStatementsLevelControlsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                    IDColumn,
		ConclusionOnControlID: ConclusionOnControlIDColumn,
		FinancialStatementsLevelControlsControlID: FinancialStatementsLevelControlsControlIDColumn,
		ControlRiskAfterEvaluatingOe:              ControlRiskAfterEvaluatingOeColumn,
		ImpactOnAuditApproach:                     ImpactOnAuditApproachColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

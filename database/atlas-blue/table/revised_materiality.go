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

var RevisedMateriality = newRevisedMaterialityTable("atlas-blue", "revised_materiality", "")

type revisedMaterialityTable struct {
	mysql.Table

	// Columns
	ID                                             mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	PlanningMaterialityID                          mysql.ColumnString // (DC2Type:guid)
	EntityStatus                                   mysql.ColumnString
	BasisOrAmountCalculationChanged                mysql.ColumnBool
	Rationale                                      mysql.ColumnString
	BasisLovID                                     mysql.ColumnString
	OtherBasisName                                 mysql.ColumnString
	FiguresBelowDisplayedInLovID                   mysql.ColumnString
	FinancialStatementQuantum                      mysql.ColumnFloat
	FinancialStatementPercentageSelected           mysql.ColumnFloat
	FinancialStatementRationale                    mysql.ColumnString
	ChosenFinancialStatementMateriality            mysql.ColumnFloat
	PerformanceStatementPercentageSelected         mysql.ColumnFloat
	PerformanceStatementRationale                  mysql.ColumnString
	ChosenPerformanceMateriality                   mysql.ColumnFloat
	ClearlyTrivialThresholdPercentageSelectedLovID mysql.ColumnString
	ChosenClearlyTrivialThreshold                  mysql.ColumnFloat
	UpdateAreasID                                  mysql.ColumnString
	Explanation                                    mysql.ColumnString
	ApplyAcrossTheAuditFile                        mysql.ColumnBool
	SelectedBenchmark                              mysql.ColumnString
	NodeID                                         mysql.ColumnString
	PeriodEndOfSourceDataUsed                      mysql.ColumnTimestamp // (DC2Type:datetime_immutable)
	RationaleRelatingToBenchmark                   mysql.ColumnString
	HeaderRationale                                mysql.ColumnString
	IsBasisOrAmountCalculationChangedReadonly      mysql.ColumnBool

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type RevisedMaterialityTable struct {
	revisedMaterialityTable

	NEW revisedMaterialityTable
}

// AS creates new RevisedMaterialityTable with assigned alias
func (a RevisedMaterialityTable) AS(alias string) *RevisedMaterialityTable {
	return newRevisedMaterialityTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new RevisedMaterialityTable with assigned schema name
func (a RevisedMaterialityTable) FromSchema(schemaName string) *RevisedMaterialityTable {
	return newRevisedMaterialityTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new RevisedMaterialityTable with assigned table prefix
func (a RevisedMaterialityTable) WithPrefix(prefix string) *RevisedMaterialityTable {
	return newRevisedMaterialityTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new RevisedMaterialityTable with assigned table suffix
func (a RevisedMaterialityTable) WithSuffix(suffix string) *RevisedMaterialityTable {
	return newRevisedMaterialityTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newRevisedMaterialityTable(schemaName, tableName, alias string) *RevisedMaterialityTable {
	return &RevisedMaterialityTable{
		revisedMaterialityTable: newRevisedMaterialityTableImpl(schemaName, tableName, alias),
		NEW:                     newRevisedMaterialityTableImpl("", "new", ""),
	}
}

func newRevisedMaterialityTableImpl(schemaName, tableName, alias string) revisedMaterialityTable {
	var (
		IDColumn                                             = mysql.StringColumn("id")
		PlanningMaterialityIDColumn                          = mysql.StringColumn("planning_materiality_id")
		EntityStatusColumn                                   = mysql.StringColumn("entity_status")
		BasisOrAmountCalculationChangedColumn                = mysql.BoolColumn("basis_or_amount_calculation_changed")
		RationaleColumn                                      = mysql.StringColumn("rationale")
		BasisLovIDColumn                                     = mysql.StringColumn("basis_lov_id")
		OtherBasisNameColumn                                 = mysql.StringColumn("other_basis_name")
		FiguresBelowDisplayedInLovIDColumn                   = mysql.StringColumn("figures_below_displayed_in_lov_id")
		FinancialStatementQuantumColumn                      = mysql.FloatColumn("financial_statement_quantum")
		FinancialStatementPercentageSelectedColumn           = mysql.FloatColumn("financial_statement_percentage_selected")
		FinancialStatementRationaleColumn                    = mysql.StringColumn("financial_statement_rationale")
		ChosenFinancialStatementMaterialityColumn            = mysql.FloatColumn("chosen_financial_statement_materiality")
		PerformanceStatementPercentageSelectedColumn         = mysql.FloatColumn("performance_statement_percentage_selected")
		PerformanceStatementRationaleColumn                  = mysql.StringColumn("performance_statement_rationale")
		ChosenPerformanceMaterialityColumn                   = mysql.FloatColumn("chosen_performance_materiality")
		ClearlyTrivialThresholdPercentageSelectedLovIDColumn = mysql.StringColumn("clearly_trivial_threshold_percentage_selected_lov_id")
		ChosenClearlyTrivialThresholdColumn                  = mysql.FloatColumn("chosen_clearly_trivial_threshold")
		UpdateAreasIDColumn                                  = mysql.StringColumn("update_areas_id")
		ExplanationColumn                                    = mysql.StringColumn("explanation")
		ApplyAcrossTheAuditFileColumn                        = mysql.BoolColumn("apply_across_the_audit_file")
		SelectedBenchmarkColumn                              = mysql.StringColumn("selected_benchmark")
		NodeIDColumn                                         = mysql.StringColumn("node_id")
		PeriodEndOfSourceDataUsedColumn                      = mysql.TimestampColumn("period_end_of_source_data_used")
		RationaleRelatingToBenchmarkColumn                   = mysql.StringColumn("rationale_relating_to_benchmark")
		HeaderRationaleColumn                                = mysql.StringColumn("header_rationale")
		IsBasisOrAmountCalculationChangedReadonlyColumn      = mysql.BoolColumn("is_basis_or_amount_calculation_changed_readonly")
		allColumns                                           = mysql.ColumnList{IDColumn, PlanningMaterialityIDColumn, EntityStatusColumn, BasisOrAmountCalculationChangedColumn, RationaleColumn, BasisLovIDColumn, OtherBasisNameColumn, FiguresBelowDisplayedInLovIDColumn, FinancialStatementQuantumColumn, FinancialStatementPercentageSelectedColumn, FinancialStatementRationaleColumn, ChosenFinancialStatementMaterialityColumn, PerformanceStatementPercentageSelectedColumn, PerformanceStatementRationaleColumn, ChosenPerformanceMaterialityColumn, ClearlyTrivialThresholdPercentageSelectedLovIDColumn, ChosenClearlyTrivialThresholdColumn, UpdateAreasIDColumn, ExplanationColumn, ApplyAcrossTheAuditFileColumn, SelectedBenchmarkColumn, NodeIDColumn, PeriodEndOfSourceDataUsedColumn, RationaleRelatingToBenchmarkColumn, HeaderRationaleColumn, IsBasisOrAmountCalculationChangedReadonlyColumn}
		mutableColumns                                       = mysql.ColumnList{PlanningMaterialityIDColumn, EntityStatusColumn, BasisOrAmountCalculationChangedColumn, RationaleColumn, BasisLovIDColumn, OtherBasisNameColumn, FiguresBelowDisplayedInLovIDColumn, FinancialStatementQuantumColumn, FinancialStatementPercentageSelectedColumn, FinancialStatementRationaleColumn, ChosenFinancialStatementMaterialityColumn, PerformanceStatementPercentageSelectedColumn, PerformanceStatementRationaleColumn, ChosenPerformanceMaterialityColumn, ClearlyTrivialThresholdPercentageSelectedLovIDColumn, ChosenClearlyTrivialThresholdColumn, UpdateAreasIDColumn, ExplanationColumn, ApplyAcrossTheAuditFileColumn, SelectedBenchmarkColumn, NodeIDColumn, PeriodEndOfSourceDataUsedColumn, RationaleRelatingToBenchmarkColumn, HeaderRationaleColumn, IsBasisOrAmountCalculationChangedReadonlyColumn}
	)

	return revisedMaterialityTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                                             IDColumn,
		PlanningMaterialityID:                          PlanningMaterialityIDColumn,
		EntityStatus:                                   EntityStatusColumn,
		BasisOrAmountCalculationChanged:                BasisOrAmountCalculationChangedColumn,
		Rationale:                                      RationaleColumn,
		BasisLovID:                                     BasisLovIDColumn,
		OtherBasisName:                                 OtherBasisNameColumn,
		FiguresBelowDisplayedInLovID:                   FiguresBelowDisplayedInLovIDColumn,
		FinancialStatementQuantum:                      FinancialStatementQuantumColumn,
		FinancialStatementPercentageSelected:           FinancialStatementPercentageSelectedColumn,
		FinancialStatementRationale:                    FinancialStatementRationaleColumn,
		ChosenFinancialStatementMateriality:            ChosenFinancialStatementMaterialityColumn,
		PerformanceStatementPercentageSelected:         PerformanceStatementPercentageSelectedColumn,
		PerformanceStatementRationale:                  PerformanceStatementRationaleColumn,
		ChosenPerformanceMateriality:                   ChosenPerformanceMaterialityColumn,
		ClearlyTrivialThresholdPercentageSelectedLovID: ClearlyTrivialThresholdPercentageSelectedLovIDColumn,
		ChosenClearlyTrivialThreshold:                  ChosenClearlyTrivialThresholdColumn,
		UpdateAreasID:                                  UpdateAreasIDColumn,
		Explanation:                                    ExplanationColumn,
		ApplyAcrossTheAuditFile:                        ApplyAcrossTheAuditFileColumn,
		SelectedBenchmark:                              SelectedBenchmarkColumn,
		NodeID:                                         NodeIDColumn,
		PeriodEndOfSourceDataUsed:                      PeriodEndOfSourceDataUsedColumn,
		RationaleRelatingToBenchmark:                   RationaleRelatingToBenchmarkColumn,
		HeaderRationale:                                HeaderRationaleColumn,
		IsBasisOrAmountCalculationChangedReadonly:      IsBasisOrAmountCalculationChangedReadonlyColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

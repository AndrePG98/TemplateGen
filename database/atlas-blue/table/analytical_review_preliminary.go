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

var AnalyticalReviewPreliminary = newAnalyticalReviewPreliminaryTable("atlas-blue", "analytical_review_preliminary", "")

type analyticalReviewPreliminaryTable struct {
	mysql.Table

	// Columns
	ID                                             mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	EngagementID                                   mysql.ColumnString // (DC2Type:guid)
	ConfirmUsingAnalyticalReview                   mysql.ColumnBool
	ConfirmUsingAnalyticalReviewComment            mysql.ColumnString
	DidIdentifyRisks                               mysql.ColumnBool
	DidIdentifyRisksComment                        mysql.ColumnString
	AnalyticalIdentifyFluctuations                 mysql.ColumnBool
	AnalyticalIdentifyFluctuationsDifferencesFound mysql.ColumnString
	DocumentInvestigationDifferences               mysql.ColumnString
	IsValid                                        mysql.ColumnBool

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type AnalyticalReviewPreliminaryTable struct {
	analyticalReviewPreliminaryTable

	NEW analyticalReviewPreliminaryTable
}

// AS creates new AnalyticalReviewPreliminaryTable with assigned alias
func (a AnalyticalReviewPreliminaryTable) AS(alias string) *AnalyticalReviewPreliminaryTable {
	return newAnalyticalReviewPreliminaryTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AnalyticalReviewPreliminaryTable with assigned schema name
func (a AnalyticalReviewPreliminaryTable) FromSchema(schemaName string) *AnalyticalReviewPreliminaryTable {
	return newAnalyticalReviewPreliminaryTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AnalyticalReviewPreliminaryTable with assigned table prefix
func (a AnalyticalReviewPreliminaryTable) WithPrefix(prefix string) *AnalyticalReviewPreliminaryTable {
	return newAnalyticalReviewPreliminaryTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AnalyticalReviewPreliminaryTable with assigned table suffix
func (a AnalyticalReviewPreliminaryTable) WithSuffix(suffix string) *AnalyticalReviewPreliminaryTable {
	return newAnalyticalReviewPreliminaryTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAnalyticalReviewPreliminaryTable(schemaName, tableName, alias string) *AnalyticalReviewPreliminaryTable {
	return &AnalyticalReviewPreliminaryTable{
		analyticalReviewPreliminaryTable: newAnalyticalReviewPreliminaryTableImpl(schemaName, tableName, alias),
		NEW:                              newAnalyticalReviewPreliminaryTableImpl("", "new", ""),
	}
}

func newAnalyticalReviewPreliminaryTableImpl(schemaName, tableName, alias string) analyticalReviewPreliminaryTable {
	var (
		IDColumn                                             = mysql.StringColumn("id")
		EngagementIDColumn                                   = mysql.StringColumn("engagement_id")
		ConfirmUsingAnalyticalReviewColumn                   = mysql.BoolColumn("confirm_using_analytical_review")
		ConfirmUsingAnalyticalReviewCommentColumn            = mysql.StringColumn("confirm_using_analytical_review_comment")
		DidIdentifyRisksColumn                               = mysql.BoolColumn("did_identify_risks")
		DidIdentifyRisksCommentColumn                        = mysql.StringColumn("did_identify_risks_comment")
		AnalyticalIdentifyFluctuationsColumn                 = mysql.BoolColumn("analytical_identify_fluctuations")
		AnalyticalIdentifyFluctuationsDifferencesFoundColumn = mysql.StringColumn("analytical_identify_fluctuations_differences_found")
		DocumentInvestigationDifferencesColumn               = mysql.StringColumn("document_investigation_differences")
		IsValidColumn                                        = mysql.BoolColumn("is_valid")
		allColumns                                           = mysql.ColumnList{IDColumn, EngagementIDColumn, ConfirmUsingAnalyticalReviewColumn, ConfirmUsingAnalyticalReviewCommentColumn, DidIdentifyRisksColumn, DidIdentifyRisksCommentColumn, AnalyticalIdentifyFluctuationsColumn, AnalyticalIdentifyFluctuationsDifferencesFoundColumn, DocumentInvestigationDifferencesColumn, IsValidColumn}
		mutableColumns                                       = mysql.ColumnList{EngagementIDColumn, ConfirmUsingAnalyticalReviewColumn, ConfirmUsingAnalyticalReviewCommentColumn, DidIdentifyRisksColumn, DidIdentifyRisksCommentColumn, AnalyticalIdentifyFluctuationsColumn, AnalyticalIdentifyFluctuationsDifferencesFoundColumn, DocumentInvestigationDifferencesColumn, IsValidColumn}
	)

	return analyticalReviewPreliminaryTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                                  IDColumn,
		EngagementID:                        EngagementIDColumn,
		ConfirmUsingAnalyticalReview:        ConfirmUsingAnalyticalReviewColumn,
		ConfirmUsingAnalyticalReviewComment: ConfirmUsingAnalyticalReviewCommentColumn,
		DidIdentifyRisks:                    DidIdentifyRisksColumn,
		DidIdentifyRisksComment:             DidIdentifyRisksCommentColumn,
		AnalyticalIdentifyFluctuations:      AnalyticalIdentifyFluctuationsColumn,
		AnalyticalIdentifyFluctuationsDifferencesFound: AnalyticalIdentifyFluctuationsDifferencesFoundColumn,
		DocumentInvestigationDifferences:               DocumentInvestigationDifferencesColumn,
		IsValid:                                        IsValidColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

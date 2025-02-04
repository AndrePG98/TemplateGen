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

var DocumentReviews = newDocumentReviewsTable("atlas-blue", "document_reviews", "")

type documentReviewsTable struct {
	mysql.Table

	// Columns
	ID             mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	ReviewedBy     mysql.ColumnString // (DC2Type:guid)
	ReversedBy     mysql.ColumnString // (DC2Type:guid)
	DocumentID     mysql.ColumnInteger
	ReviewDate     mysql.ColumnTimestamp // (DC2Type:datetime_immutable)
	Reversed       mysql.ColumnBool
	ReversedDate   mysql.ColumnTimestamp // (DC2Type:datetime_immutable)
	ReviewedByEqcr mysql.ColumnBool
	Preparers      mysql.ColumnString // (DC2Type:array)

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type DocumentReviewsTable struct {
	documentReviewsTable

	NEW documentReviewsTable
}

// AS creates new DocumentReviewsTable with assigned alias
func (a DocumentReviewsTable) AS(alias string) *DocumentReviewsTable {
	return newDocumentReviewsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new DocumentReviewsTable with assigned schema name
func (a DocumentReviewsTable) FromSchema(schemaName string) *DocumentReviewsTable {
	return newDocumentReviewsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new DocumentReviewsTable with assigned table prefix
func (a DocumentReviewsTable) WithPrefix(prefix string) *DocumentReviewsTable {
	return newDocumentReviewsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new DocumentReviewsTable with assigned table suffix
func (a DocumentReviewsTable) WithSuffix(suffix string) *DocumentReviewsTable {
	return newDocumentReviewsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newDocumentReviewsTable(schemaName, tableName, alias string) *DocumentReviewsTable {
	return &DocumentReviewsTable{
		documentReviewsTable: newDocumentReviewsTableImpl(schemaName, tableName, alias),
		NEW:                  newDocumentReviewsTableImpl("", "new", ""),
	}
}

func newDocumentReviewsTableImpl(schemaName, tableName, alias string) documentReviewsTable {
	var (
		IDColumn             = mysql.StringColumn("id")
		ReviewedByColumn     = mysql.StringColumn("reviewed_by")
		ReversedByColumn     = mysql.StringColumn("reversed_by")
		DocumentIDColumn     = mysql.IntegerColumn("document_id")
		ReviewDateColumn     = mysql.TimestampColumn("review_date")
		ReversedColumn       = mysql.BoolColumn("reversed")
		ReversedDateColumn   = mysql.TimestampColumn("reversed_date")
		ReviewedByEqcrColumn = mysql.BoolColumn("reviewed_by_eqcr")
		PreparersColumn      = mysql.StringColumn("preparers")
		allColumns           = mysql.ColumnList{IDColumn, ReviewedByColumn, ReversedByColumn, DocumentIDColumn, ReviewDateColumn, ReversedColumn, ReversedDateColumn, ReviewedByEqcrColumn, PreparersColumn}
		mutableColumns       = mysql.ColumnList{ReviewedByColumn, ReversedByColumn, DocumentIDColumn, ReviewDateColumn, ReversedColumn, ReversedDateColumn, ReviewedByEqcrColumn, PreparersColumn}
	)

	return documentReviewsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:             IDColumn,
		ReviewedBy:     ReviewedByColumn,
		ReversedBy:     ReversedByColumn,
		DocumentID:     DocumentIDColumn,
		ReviewDate:     ReviewDateColumn,
		Reversed:       ReversedColumn,
		ReversedDate:   ReversedDateColumn,
		ReviewedByEqcr: ReviewedByEqcrColumn,
		Preparers:      PreparersColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

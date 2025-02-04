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

var RelatedPartyCotabds = newRelatedPartyCotabdsTable("atlas-blue", "related_party_cotabds", "")

type relatedPartyCotabdsTable struct {
	mysql.Table

	// Columns
	ID             mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	RelatedPartyID mysql.ColumnString // (DC2Type:guid)
	CotabdID       mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type RelatedPartyCotabdsTable struct {
	relatedPartyCotabdsTable

	NEW relatedPartyCotabdsTable
}

// AS creates new RelatedPartyCotabdsTable with assigned alias
func (a RelatedPartyCotabdsTable) AS(alias string) *RelatedPartyCotabdsTable {
	return newRelatedPartyCotabdsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new RelatedPartyCotabdsTable with assigned schema name
func (a RelatedPartyCotabdsTable) FromSchema(schemaName string) *RelatedPartyCotabdsTable {
	return newRelatedPartyCotabdsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new RelatedPartyCotabdsTable with assigned table prefix
func (a RelatedPartyCotabdsTable) WithPrefix(prefix string) *RelatedPartyCotabdsTable {
	return newRelatedPartyCotabdsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new RelatedPartyCotabdsTable with assigned table suffix
func (a RelatedPartyCotabdsTable) WithSuffix(suffix string) *RelatedPartyCotabdsTable {
	return newRelatedPartyCotabdsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newRelatedPartyCotabdsTable(schemaName, tableName, alias string) *RelatedPartyCotabdsTable {
	return &RelatedPartyCotabdsTable{
		relatedPartyCotabdsTable: newRelatedPartyCotabdsTableImpl(schemaName, tableName, alias),
		NEW:                      newRelatedPartyCotabdsTableImpl("", "new", ""),
	}
}

func newRelatedPartyCotabdsTableImpl(schemaName, tableName, alias string) relatedPartyCotabdsTable {
	var (
		IDColumn             = mysql.StringColumn("id")
		RelatedPartyIDColumn = mysql.StringColumn("related_party_id")
		CotabdIDColumn       = mysql.StringColumn("cotabd_id")
		allColumns           = mysql.ColumnList{IDColumn, RelatedPartyIDColumn, CotabdIDColumn}
		mutableColumns       = mysql.ColumnList{RelatedPartyIDColumn, CotabdIDColumn}
	)

	return relatedPartyCotabdsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:             IDColumn,
		RelatedPartyID: RelatedPartyIDColumn,
		CotabdID:       CotabdIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

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

var TbAccountsCotabds = newTbAccountsCotabdsTable("atlas-blue", "tb_accounts_cotabds", "")

type tbAccountsCotabdsTable struct {
	mysql.Table

	// Columns
	ID           mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	EngagementID mysql.ColumnString // (DC2Type:guid)
	AccountID    mysql.ColumnString // (DC2Type:guid)
	CotabID      mysql.ColumnString // (DC2Type:guid)
	ParentID     mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type TbAccountsCotabdsTable struct {
	tbAccountsCotabdsTable

	NEW tbAccountsCotabdsTable
}

// AS creates new TbAccountsCotabdsTable with assigned alias
func (a TbAccountsCotabdsTable) AS(alias string) *TbAccountsCotabdsTable {
	return newTbAccountsCotabdsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TbAccountsCotabdsTable with assigned schema name
func (a TbAccountsCotabdsTable) FromSchema(schemaName string) *TbAccountsCotabdsTable {
	return newTbAccountsCotabdsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TbAccountsCotabdsTable with assigned table prefix
func (a TbAccountsCotabdsTable) WithPrefix(prefix string) *TbAccountsCotabdsTable {
	return newTbAccountsCotabdsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TbAccountsCotabdsTable with assigned table suffix
func (a TbAccountsCotabdsTable) WithSuffix(suffix string) *TbAccountsCotabdsTable {
	return newTbAccountsCotabdsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTbAccountsCotabdsTable(schemaName, tableName, alias string) *TbAccountsCotabdsTable {
	return &TbAccountsCotabdsTable{
		tbAccountsCotabdsTable: newTbAccountsCotabdsTableImpl(schemaName, tableName, alias),
		NEW:                    newTbAccountsCotabdsTableImpl("", "new", ""),
	}
}

func newTbAccountsCotabdsTableImpl(schemaName, tableName, alias string) tbAccountsCotabdsTable {
	var (
		IDColumn           = mysql.StringColumn("id")
		EngagementIDColumn = mysql.StringColumn("engagement_id")
		AccountIDColumn    = mysql.StringColumn("account_id")
		CotabIDColumn      = mysql.StringColumn("cotab_id")
		ParentIDColumn     = mysql.StringColumn("parent_id")
		allColumns         = mysql.ColumnList{IDColumn, EngagementIDColumn, AccountIDColumn, CotabIDColumn, ParentIDColumn}
		mutableColumns     = mysql.ColumnList{EngagementIDColumn, AccountIDColumn, CotabIDColumn, ParentIDColumn}
	)

	return tbAccountsCotabdsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		EngagementID: EngagementIDColumn,
		AccountID:    AccountIDColumn,
		CotabID:      CotabIDColumn,
		ParentID:     ParentIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

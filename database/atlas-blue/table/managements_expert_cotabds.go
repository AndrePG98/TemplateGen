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

var ManagementsExpertCotabds = newManagementsExpertCotabdsTable("atlas-blue", "managements_expert_cotabds", "")

type managementsExpertCotabdsTable struct {
	mysql.Table

	// Columns
	ID                  mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	ManagementsExpertID mysql.ColumnString // (DC2Type:guid)
	CotabdID            mysql.ColumnString // (DC2Type:guid)

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ManagementsExpertCotabdsTable struct {
	managementsExpertCotabdsTable

	NEW managementsExpertCotabdsTable
}

// AS creates new ManagementsExpertCotabdsTable with assigned alias
func (a ManagementsExpertCotabdsTable) AS(alias string) *ManagementsExpertCotabdsTable {
	return newManagementsExpertCotabdsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ManagementsExpertCotabdsTable with assigned schema name
func (a ManagementsExpertCotabdsTable) FromSchema(schemaName string) *ManagementsExpertCotabdsTable {
	return newManagementsExpertCotabdsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ManagementsExpertCotabdsTable with assigned table prefix
func (a ManagementsExpertCotabdsTable) WithPrefix(prefix string) *ManagementsExpertCotabdsTable {
	return newManagementsExpertCotabdsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ManagementsExpertCotabdsTable with assigned table suffix
func (a ManagementsExpertCotabdsTable) WithSuffix(suffix string) *ManagementsExpertCotabdsTable {
	return newManagementsExpertCotabdsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newManagementsExpertCotabdsTable(schemaName, tableName, alias string) *ManagementsExpertCotabdsTable {
	return &ManagementsExpertCotabdsTable{
		managementsExpertCotabdsTable: newManagementsExpertCotabdsTableImpl(schemaName, tableName, alias),
		NEW:                           newManagementsExpertCotabdsTableImpl("", "new", ""),
	}
}

func newManagementsExpertCotabdsTableImpl(schemaName, tableName, alias string) managementsExpertCotabdsTable {
	var (
		IDColumn                  = mysql.StringColumn("id")
		ManagementsExpertIDColumn = mysql.StringColumn("managements_expert_id")
		CotabdIDColumn            = mysql.StringColumn("cotabd_id")
		allColumns                = mysql.ColumnList{IDColumn, ManagementsExpertIDColumn, CotabdIDColumn}
		mutableColumns            = mysql.ColumnList{ManagementsExpertIDColumn, CotabdIDColumn}
	)

	return managementsExpertCotabdsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                  IDColumn,
		ManagementsExpertID: ManagementsExpertIDColumn,
		CotabdID:            CotabdIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

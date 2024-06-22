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

var ItGeneralControls = newItGeneralControlsTable("atlas-blue", "it_general_controls", "")

type itGeneralControlsTable struct {
	mysql.Table

	// Columns
	ID           mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	EngagementID mysql.ColumnString // (DC2Type:guid)
	NodeID       mysql.ColumnString
	Name         mysql.ColumnString
	IsDeleted    mysql.ColumnBool

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ItGeneralControlsTable struct {
	itGeneralControlsTable

	NEW itGeneralControlsTable
}

// AS creates new ItGeneralControlsTable with assigned alias
func (a ItGeneralControlsTable) AS(alias string) *ItGeneralControlsTable {
	return newItGeneralControlsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ItGeneralControlsTable with assigned schema name
func (a ItGeneralControlsTable) FromSchema(schemaName string) *ItGeneralControlsTable {
	return newItGeneralControlsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ItGeneralControlsTable with assigned table prefix
func (a ItGeneralControlsTable) WithPrefix(prefix string) *ItGeneralControlsTable {
	return newItGeneralControlsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ItGeneralControlsTable with assigned table suffix
func (a ItGeneralControlsTable) WithSuffix(suffix string) *ItGeneralControlsTable {
	return newItGeneralControlsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newItGeneralControlsTable(schemaName, tableName, alias string) *ItGeneralControlsTable {
	return &ItGeneralControlsTable{
		itGeneralControlsTable: newItGeneralControlsTableImpl(schemaName, tableName, alias),
		NEW:                    newItGeneralControlsTableImpl("", "new", ""),
	}
}

func newItGeneralControlsTableImpl(schemaName, tableName, alias string) itGeneralControlsTable {
	var (
		IDColumn           = mysql.StringColumn("id")
		EngagementIDColumn = mysql.StringColumn("engagement_id")
		NodeIDColumn       = mysql.StringColumn("node_id")
		NameColumn         = mysql.StringColumn("name")
		IsDeletedColumn    = mysql.BoolColumn("is_deleted")
		allColumns         = mysql.ColumnList{IDColumn, EngagementIDColumn, NodeIDColumn, NameColumn, IsDeletedColumn}
		mutableColumns     = mysql.ColumnList{EngagementIDColumn, NodeIDColumn, NameColumn, IsDeletedColumn}
	)

	return itGeneralControlsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		EngagementID: EngagementIDColumn,
		NodeID:       NodeIDColumn,
		Name:         NameColumn,
		IsDeleted:    IsDeletedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

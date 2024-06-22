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

var ItGeneralControlIdentifications = newItGeneralControlIdentificationsTable("atlas-blue", "it_general_control_identifications", "")

type itGeneralControlIdentificationsTable struct {
	mysql.Table

	// Columns
	ID           mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	EngagementID mysql.ColumnString // (DC2Type:guid)
	NodeID       mysql.ColumnString
	EntityStatus mysql.ColumnString
	Comment      mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ItGeneralControlIdentificationsTable struct {
	itGeneralControlIdentificationsTable

	NEW itGeneralControlIdentificationsTable
}

// AS creates new ItGeneralControlIdentificationsTable with assigned alias
func (a ItGeneralControlIdentificationsTable) AS(alias string) *ItGeneralControlIdentificationsTable {
	return newItGeneralControlIdentificationsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ItGeneralControlIdentificationsTable with assigned schema name
func (a ItGeneralControlIdentificationsTable) FromSchema(schemaName string) *ItGeneralControlIdentificationsTable {
	return newItGeneralControlIdentificationsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ItGeneralControlIdentificationsTable with assigned table prefix
func (a ItGeneralControlIdentificationsTable) WithPrefix(prefix string) *ItGeneralControlIdentificationsTable {
	return newItGeneralControlIdentificationsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ItGeneralControlIdentificationsTable with assigned table suffix
func (a ItGeneralControlIdentificationsTable) WithSuffix(suffix string) *ItGeneralControlIdentificationsTable {
	return newItGeneralControlIdentificationsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newItGeneralControlIdentificationsTable(schemaName, tableName, alias string) *ItGeneralControlIdentificationsTable {
	return &ItGeneralControlIdentificationsTable{
		itGeneralControlIdentificationsTable: newItGeneralControlIdentificationsTableImpl(schemaName, tableName, alias),
		NEW:                                  newItGeneralControlIdentificationsTableImpl("", "new", ""),
	}
}

func newItGeneralControlIdentificationsTableImpl(schemaName, tableName, alias string) itGeneralControlIdentificationsTable {
	var (
		IDColumn           = mysql.StringColumn("id")
		EngagementIDColumn = mysql.StringColumn("engagement_id")
		NodeIDColumn       = mysql.StringColumn("node_id")
		EntityStatusColumn = mysql.StringColumn("entity_status")
		CommentColumn      = mysql.StringColumn("comment")
		allColumns         = mysql.ColumnList{IDColumn, EngagementIDColumn, NodeIDColumn, EntityStatusColumn, CommentColumn}
		mutableColumns     = mysql.ColumnList{EngagementIDColumn, NodeIDColumn, EntityStatusColumn, CommentColumn}
	)

	return itGeneralControlIdentificationsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		EngagementID: EngagementIDColumn,
		NodeID:       NodeIDColumn,
		EntityStatus: EntityStatusColumn,
		Comment:      CommentColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

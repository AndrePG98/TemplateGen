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

var EngagementFieldValues = newEngagementFieldValuesTable("atlas-blue", "engagement_field_values", "")

type engagementFieldValuesTable struct {
	mysql.Table

	// Columns
	ID           mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	EngagementID mysql.ColumnString // (DC2Type:guid)
	FieldID      mysql.ColumnString
	Value        mysql.ColumnString // (DC2Type:array)
	Comment      mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type EngagementFieldValuesTable struct {
	engagementFieldValuesTable

	NEW engagementFieldValuesTable
}

// AS creates new EngagementFieldValuesTable with assigned alias
func (a EngagementFieldValuesTable) AS(alias string) *EngagementFieldValuesTable {
	return newEngagementFieldValuesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new EngagementFieldValuesTable with assigned schema name
func (a EngagementFieldValuesTable) FromSchema(schemaName string) *EngagementFieldValuesTable {
	return newEngagementFieldValuesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new EngagementFieldValuesTable with assigned table prefix
func (a EngagementFieldValuesTable) WithPrefix(prefix string) *EngagementFieldValuesTable {
	return newEngagementFieldValuesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new EngagementFieldValuesTable with assigned table suffix
func (a EngagementFieldValuesTable) WithSuffix(suffix string) *EngagementFieldValuesTable {
	return newEngagementFieldValuesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newEngagementFieldValuesTable(schemaName, tableName, alias string) *EngagementFieldValuesTable {
	return &EngagementFieldValuesTable{
		engagementFieldValuesTable: newEngagementFieldValuesTableImpl(schemaName, tableName, alias),
		NEW:                        newEngagementFieldValuesTableImpl("", "new", ""),
	}
}

func newEngagementFieldValuesTableImpl(schemaName, tableName, alias string) engagementFieldValuesTable {
	var (
		IDColumn           = mysql.StringColumn("id")
		EngagementIDColumn = mysql.StringColumn("engagement_id")
		FieldIDColumn      = mysql.StringColumn("field_id")
		ValueColumn        = mysql.StringColumn("value")
		CommentColumn      = mysql.StringColumn("comment")
		allColumns         = mysql.ColumnList{IDColumn, EngagementIDColumn, FieldIDColumn, ValueColumn, CommentColumn}
		mutableColumns     = mysql.ColumnList{EngagementIDColumn, FieldIDColumn, ValueColumn, CommentColumn}
	)

	return engagementFieldValuesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		EngagementID: EngagementIDColumn,
		FieldID:      FieldIDColumn,
		Value:        ValueColumn,
		Comment:      CommentColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

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

var Reasons = newReasonsTable("atlas-blue", "reasons", "")

type reasonsTable struct {
	mysql.Table

	// Columns
	ID                  mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	EngagementID        mysql.ColumnString // (DC2Type:guid)
	ReasonForUnadjusted mysql.ColumnString
	CreatedAt           mysql.ColumnTimestamp // (DC2Type:datetime_immutable)

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ReasonsTable struct {
	reasonsTable

	NEW reasonsTable
}

// AS creates new ReasonsTable with assigned alias
func (a ReasonsTable) AS(alias string) *ReasonsTable {
	return newReasonsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ReasonsTable with assigned schema name
func (a ReasonsTable) FromSchema(schemaName string) *ReasonsTable {
	return newReasonsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ReasonsTable with assigned table prefix
func (a ReasonsTable) WithPrefix(prefix string) *ReasonsTable {
	return newReasonsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ReasonsTable with assigned table suffix
func (a ReasonsTable) WithSuffix(suffix string) *ReasonsTable {
	return newReasonsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newReasonsTable(schemaName, tableName, alias string) *ReasonsTable {
	return &ReasonsTable{
		reasonsTable: newReasonsTableImpl(schemaName, tableName, alias),
		NEW:          newReasonsTableImpl("", "new", ""),
	}
}

func newReasonsTableImpl(schemaName, tableName, alias string) reasonsTable {
	var (
		IDColumn                  = mysql.StringColumn("id")
		EngagementIDColumn        = mysql.StringColumn("engagement_id")
		ReasonForUnadjustedColumn = mysql.StringColumn("reason_for_unadjusted")
		CreatedAtColumn           = mysql.TimestampColumn("created_at")
		allColumns                = mysql.ColumnList{IDColumn, EngagementIDColumn, ReasonForUnadjustedColumn, CreatedAtColumn}
		mutableColumns            = mysql.ColumnList{EngagementIDColumn, ReasonForUnadjustedColumn, CreatedAtColumn}
	)

	return reasonsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                  IDColumn,
		EngagementID:        EngagementIDColumn,
		ReasonForUnadjusted: ReasonForUnadjustedColumn,
		CreatedAt:           CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

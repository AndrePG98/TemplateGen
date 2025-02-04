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

var AuditTreeNodeConflicts = newAuditTreeNodeConflictsTable("atlas-blue", "audit_tree_node_conflicts", "")

type auditTreeNodeConflictsTable struct {
	mysql.Table

	// Columns
	ID               mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	AuditTreeNodesID mysql.ColumnInteger
	HasConflict      mysql.ColumnBool
	ConflictComment  mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type AuditTreeNodeConflictsTable struct {
	auditTreeNodeConflictsTable

	NEW auditTreeNodeConflictsTable
}

// AS creates new AuditTreeNodeConflictsTable with assigned alias
func (a AuditTreeNodeConflictsTable) AS(alias string) *AuditTreeNodeConflictsTable {
	return newAuditTreeNodeConflictsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AuditTreeNodeConflictsTable with assigned schema name
func (a AuditTreeNodeConflictsTable) FromSchema(schemaName string) *AuditTreeNodeConflictsTable {
	return newAuditTreeNodeConflictsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AuditTreeNodeConflictsTable with assigned table prefix
func (a AuditTreeNodeConflictsTable) WithPrefix(prefix string) *AuditTreeNodeConflictsTable {
	return newAuditTreeNodeConflictsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AuditTreeNodeConflictsTable with assigned table suffix
func (a AuditTreeNodeConflictsTable) WithSuffix(suffix string) *AuditTreeNodeConflictsTable {
	return newAuditTreeNodeConflictsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAuditTreeNodeConflictsTable(schemaName, tableName, alias string) *AuditTreeNodeConflictsTable {
	return &AuditTreeNodeConflictsTable{
		auditTreeNodeConflictsTable: newAuditTreeNodeConflictsTableImpl(schemaName, tableName, alias),
		NEW:                         newAuditTreeNodeConflictsTableImpl("", "new", ""),
	}
}

func newAuditTreeNodeConflictsTableImpl(schemaName, tableName, alias string) auditTreeNodeConflictsTable {
	var (
		IDColumn               = mysql.StringColumn("id")
		AuditTreeNodesIDColumn = mysql.IntegerColumn("audit_tree_nodes_id")
		HasConflictColumn      = mysql.BoolColumn("has_conflict")
		ConflictCommentColumn  = mysql.StringColumn("conflict_comment")
		allColumns             = mysql.ColumnList{IDColumn, AuditTreeNodesIDColumn, HasConflictColumn, ConflictCommentColumn}
		mutableColumns         = mysql.ColumnList{AuditTreeNodesIDColumn, HasConflictColumn, ConflictCommentColumn}
	)

	return auditTreeNodeConflictsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:               IDColumn,
		AuditTreeNodesID: AuditTreeNodesIDColumn,
		HasConflict:      HasConflictColumn,
		ConflictComment:  ConflictCommentColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

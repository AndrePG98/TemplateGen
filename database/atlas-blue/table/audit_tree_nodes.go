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

var AuditTreeNodes = newAuditTreeNodesTable("atlas-blue", "audit_tree_nodes", "")

type auditTreeNodesTable struct {
	mysql.Table

	// Columns
	ID            mysql.ColumnInteger
	EngagementID  mysql.ColumnString // (DC2Type:guid)
	NodeID        mysql.ColumnString
	ParentNodeID  mysql.ColumnString
	NodeLabel     mysql.ColumnString // (DC2Type:json)
	NodeType      mysql.ColumnString
	FormType      mysql.ColumnString
	NodeReference mysql.ColumnString
	Hidden        mysql.ColumnBool
	Position      mysql.ColumnInteger
	Disabled      mysql.ColumnBool
	AuditWorkflow mysql.ColumnString
	Rollforwarded mysql.ColumnBool

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type AuditTreeNodesTable struct {
	auditTreeNodesTable

	NEW auditTreeNodesTable
}

// AS creates new AuditTreeNodesTable with assigned alias
func (a AuditTreeNodesTable) AS(alias string) *AuditTreeNodesTable {
	return newAuditTreeNodesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AuditTreeNodesTable with assigned schema name
func (a AuditTreeNodesTable) FromSchema(schemaName string) *AuditTreeNodesTable {
	return newAuditTreeNodesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AuditTreeNodesTable with assigned table prefix
func (a AuditTreeNodesTable) WithPrefix(prefix string) *AuditTreeNodesTable {
	return newAuditTreeNodesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AuditTreeNodesTable with assigned table suffix
func (a AuditTreeNodesTable) WithSuffix(suffix string) *AuditTreeNodesTable {
	return newAuditTreeNodesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAuditTreeNodesTable(schemaName, tableName, alias string) *AuditTreeNodesTable {
	return &AuditTreeNodesTable{
		auditTreeNodesTable: newAuditTreeNodesTableImpl(schemaName, tableName, alias),
		NEW:                 newAuditTreeNodesTableImpl("", "new", ""),
	}
}

func newAuditTreeNodesTableImpl(schemaName, tableName, alias string) auditTreeNodesTable {
	var (
		IDColumn            = mysql.IntegerColumn("id")
		EngagementIDColumn  = mysql.StringColumn("engagement_id")
		NodeIDColumn        = mysql.StringColumn("node_id")
		ParentNodeIDColumn  = mysql.StringColumn("parent_node_id")
		NodeLabelColumn     = mysql.StringColumn("node_label")
		NodeTypeColumn      = mysql.StringColumn("node_type")
		FormTypeColumn      = mysql.StringColumn("form_type")
		NodeReferenceColumn = mysql.StringColumn("node_reference")
		HiddenColumn        = mysql.BoolColumn("hidden")
		PositionColumn      = mysql.IntegerColumn("position")
		DisabledColumn      = mysql.BoolColumn("disabled")
		AuditWorkflowColumn = mysql.StringColumn("audit_workflow")
		RollforwardedColumn = mysql.BoolColumn("rollforwarded")
		allColumns          = mysql.ColumnList{IDColumn, EngagementIDColumn, NodeIDColumn, ParentNodeIDColumn, NodeLabelColumn, NodeTypeColumn, FormTypeColumn, NodeReferenceColumn, HiddenColumn, PositionColumn, DisabledColumn, AuditWorkflowColumn, RollforwardedColumn}
		mutableColumns      = mysql.ColumnList{EngagementIDColumn, NodeIDColumn, ParentNodeIDColumn, NodeLabelColumn, NodeTypeColumn, FormTypeColumn, NodeReferenceColumn, HiddenColumn, PositionColumn, DisabledColumn, AuditWorkflowColumn, RollforwardedColumn}
	)

	return auditTreeNodesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:            IDColumn,
		EngagementID:  EngagementIDColumn,
		NodeID:        NodeIDColumn,
		ParentNodeID:  ParentNodeIDColumn,
		NodeLabel:     NodeLabelColumn,
		NodeType:      NodeTypeColumn,
		FormType:      FormTypeColumn,
		NodeReference: NodeReferenceColumn,
		Hidden:        HiddenColumn,
		Position:      PositionColumn,
		Disabled:      DisabledColumn,
		AuditWorkflow: AuditWorkflowColumn,
		Rollforwarded: RollforwardedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

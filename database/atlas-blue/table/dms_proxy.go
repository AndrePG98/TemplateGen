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

var DmsProxy = newDmsProxyTable("atlas-blue", "dms_proxy", "")

type dmsProxyTable struct {
	mysql.Table

	// Columns
	ID          mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	ReceiptDate mysql.ColumnTimestamp
	Service     mysql.ColumnString
	Method      mysql.ColumnString
	Resource    mysql.ColumnString
	Version     mysql.ColumnString
	Headers     mysql.ColumnString // (DC2Type:array)
	Content     mysql.ColumnString // (DC2Type:array)
	ContentRaw  mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type DmsProxyTable struct {
	dmsProxyTable

	NEW dmsProxyTable
}

// AS creates new DmsProxyTable with assigned alias
func (a DmsProxyTable) AS(alias string) *DmsProxyTable {
	return newDmsProxyTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new DmsProxyTable with assigned schema name
func (a DmsProxyTable) FromSchema(schemaName string) *DmsProxyTable {
	return newDmsProxyTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new DmsProxyTable with assigned table prefix
func (a DmsProxyTable) WithPrefix(prefix string) *DmsProxyTable {
	return newDmsProxyTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new DmsProxyTable with assigned table suffix
func (a DmsProxyTable) WithSuffix(suffix string) *DmsProxyTable {
	return newDmsProxyTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newDmsProxyTable(schemaName, tableName, alias string) *DmsProxyTable {
	return &DmsProxyTable{
		dmsProxyTable: newDmsProxyTableImpl(schemaName, tableName, alias),
		NEW:           newDmsProxyTableImpl("", "new", ""),
	}
}

func newDmsProxyTableImpl(schemaName, tableName, alias string) dmsProxyTable {
	var (
		IDColumn          = mysql.StringColumn("id")
		ReceiptDateColumn = mysql.TimestampColumn("receipt_date")
		ServiceColumn     = mysql.StringColumn("service")
		MethodColumn      = mysql.StringColumn("method")
		ResourceColumn    = mysql.StringColumn("resource")
		VersionColumn     = mysql.StringColumn("version")
		HeadersColumn     = mysql.StringColumn("headers")
		ContentColumn     = mysql.StringColumn("content")
		ContentRawColumn  = mysql.StringColumn("content_raw")
		allColumns        = mysql.ColumnList{IDColumn, ReceiptDateColumn, ServiceColumn, MethodColumn, ResourceColumn, VersionColumn, HeadersColumn, ContentColumn, ContentRawColumn}
		mutableColumns    = mysql.ColumnList{ReceiptDateColumn, ServiceColumn, MethodColumn, ResourceColumn, VersionColumn, HeadersColumn, ContentColumn, ContentRawColumn}
	)

	return dmsProxyTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		ReceiptDate: ReceiptDateColumn,
		Service:     ServiceColumn,
		Method:      MethodColumn,
		Resource:    ResourceColumn,
		Version:     VersionColumn,
		Headers:     HeadersColumn,
		Content:     ContentColumn,
		ContentRaw:  ContentRawColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

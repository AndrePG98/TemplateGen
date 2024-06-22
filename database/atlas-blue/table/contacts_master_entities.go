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

var ContactsMasterEntities = newContactsMasterEntitiesTable("atlas-blue", "contacts_master_entities", "")

type contactsMasterEntitiesTable struct {
	mysql.Table

	// Columns
	ContactID      mysql.ColumnString // (DC2Type:guid)
	MasterEntityID mysql.ColumnString // (DC2Type:guid)

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ContactsMasterEntitiesTable struct {
	contactsMasterEntitiesTable

	NEW contactsMasterEntitiesTable
}

// AS creates new ContactsMasterEntitiesTable with assigned alias
func (a ContactsMasterEntitiesTable) AS(alias string) *ContactsMasterEntitiesTable {
	return newContactsMasterEntitiesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ContactsMasterEntitiesTable with assigned schema name
func (a ContactsMasterEntitiesTable) FromSchema(schemaName string) *ContactsMasterEntitiesTable {
	return newContactsMasterEntitiesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ContactsMasterEntitiesTable with assigned table prefix
func (a ContactsMasterEntitiesTable) WithPrefix(prefix string) *ContactsMasterEntitiesTable {
	return newContactsMasterEntitiesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ContactsMasterEntitiesTable with assigned table suffix
func (a ContactsMasterEntitiesTable) WithSuffix(suffix string) *ContactsMasterEntitiesTable {
	return newContactsMasterEntitiesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newContactsMasterEntitiesTable(schemaName, tableName, alias string) *ContactsMasterEntitiesTable {
	return &ContactsMasterEntitiesTable{
		contactsMasterEntitiesTable: newContactsMasterEntitiesTableImpl(schemaName, tableName, alias),
		NEW:                         newContactsMasterEntitiesTableImpl("", "new", ""),
	}
}

func newContactsMasterEntitiesTableImpl(schemaName, tableName, alias string) contactsMasterEntitiesTable {
	var (
		ContactIDColumn      = mysql.StringColumn("contact_id")
		MasterEntityIDColumn = mysql.StringColumn("master_entity_id")
		allColumns           = mysql.ColumnList{ContactIDColumn, MasterEntityIDColumn}
		mutableColumns       = mysql.ColumnList{}
	)

	return contactsMasterEntitiesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ContactID:      ContactIDColumn,
		MasterEntityID: MasterEntityIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

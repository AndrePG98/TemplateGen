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

var UsersAppRoles = newUsersAppRolesTable("atlas-blue", "users_app_roles", "")

type usersAppRolesTable struct {
	mysql.Table

	// Columns
	UserID    mysql.ColumnString // (DC2Type:guid)
	AppRoleID mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type UsersAppRolesTable struct {
	usersAppRolesTable

	NEW usersAppRolesTable
}

// AS creates new UsersAppRolesTable with assigned alias
func (a UsersAppRolesTable) AS(alias string) *UsersAppRolesTable {
	return newUsersAppRolesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new UsersAppRolesTable with assigned schema name
func (a UsersAppRolesTable) FromSchema(schemaName string) *UsersAppRolesTable {
	return newUsersAppRolesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new UsersAppRolesTable with assigned table prefix
func (a UsersAppRolesTable) WithPrefix(prefix string) *UsersAppRolesTable {
	return newUsersAppRolesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new UsersAppRolesTable with assigned table suffix
func (a UsersAppRolesTable) WithSuffix(suffix string) *UsersAppRolesTable {
	return newUsersAppRolesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newUsersAppRolesTable(schemaName, tableName, alias string) *UsersAppRolesTable {
	return &UsersAppRolesTable{
		usersAppRolesTable: newUsersAppRolesTableImpl(schemaName, tableName, alias),
		NEW:                newUsersAppRolesTableImpl("", "new", ""),
	}
}

func newUsersAppRolesTableImpl(schemaName, tableName, alias string) usersAppRolesTable {
	var (
		UserIDColumn    = mysql.StringColumn("user_id")
		AppRoleIDColumn = mysql.StringColumn("app_role_id")
		allColumns      = mysql.ColumnList{UserIDColumn, AppRoleIDColumn}
		mutableColumns  = mysql.ColumnList{}
	)

	return usersAppRolesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID:    UserIDColumn,
		AppRoleID: AppRoleIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

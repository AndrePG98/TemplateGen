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

var OtherFslrBusinessRisks = newOtherFslrBusinessRisksTable("atlas-blue", "other_fslr_business_risks", "")

type otherFslrBusinessRisksTable struct {
	mysql.Table

	// Columns
	ID             mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	Name           mysql.ColumnString
	Description    mysql.ColumnString
	EngagementID   mysql.ColumnString // (DC2Type:guid)
	CreatedAt      mysql.ColumnTimestamp
	IsDeleted      mysql.ColumnBool
	BusinessRiskID mysql.ColumnString // (DC2Type:guid)

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type OtherFslrBusinessRisksTable struct {
	otherFslrBusinessRisksTable

	NEW otherFslrBusinessRisksTable
}

// AS creates new OtherFslrBusinessRisksTable with assigned alias
func (a OtherFslrBusinessRisksTable) AS(alias string) *OtherFslrBusinessRisksTable {
	return newOtherFslrBusinessRisksTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new OtherFslrBusinessRisksTable with assigned schema name
func (a OtherFslrBusinessRisksTable) FromSchema(schemaName string) *OtherFslrBusinessRisksTable {
	return newOtherFslrBusinessRisksTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new OtherFslrBusinessRisksTable with assigned table prefix
func (a OtherFslrBusinessRisksTable) WithPrefix(prefix string) *OtherFslrBusinessRisksTable {
	return newOtherFslrBusinessRisksTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new OtherFslrBusinessRisksTable with assigned table suffix
func (a OtherFslrBusinessRisksTable) WithSuffix(suffix string) *OtherFslrBusinessRisksTable {
	return newOtherFslrBusinessRisksTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newOtherFslrBusinessRisksTable(schemaName, tableName, alias string) *OtherFslrBusinessRisksTable {
	return &OtherFslrBusinessRisksTable{
		otherFslrBusinessRisksTable: newOtherFslrBusinessRisksTableImpl(schemaName, tableName, alias),
		NEW:                         newOtherFslrBusinessRisksTableImpl("", "new", ""),
	}
}

func newOtherFslrBusinessRisksTableImpl(schemaName, tableName, alias string) otherFslrBusinessRisksTable {
	var (
		IDColumn             = mysql.StringColumn("id")
		NameColumn           = mysql.StringColumn("name")
		DescriptionColumn    = mysql.StringColumn("description")
		EngagementIDColumn   = mysql.StringColumn("engagement_id")
		CreatedAtColumn      = mysql.TimestampColumn("created_at")
		IsDeletedColumn      = mysql.BoolColumn("is_deleted")
		BusinessRiskIDColumn = mysql.StringColumn("business_risk_id")
		allColumns           = mysql.ColumnList{IDColumn, NameColumn, DescriptionColumn, EngagementIDColumn, CreatedAtColumn, IsDeletedColumn, BusinessRiskIDColumn}
		mutableColumns       = mysql.ColumnList{NameColumn, DescriptionColumn, EngagementIDColumn, CreatedAtColumn, IsDeletedColumn, BusinessRiskIDColumn}
	)

	return otherFslrBusinessRisksTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:             IDColumn,
		Name:           NameColumn,
		Description:    DescriptionColumn,
		EngagementID:   EngagementIDColumn,
		CreatedAt:      CreatedAtColumn,
		IsDeleted:      IsDeletedColumn,
		BusinessRiskID: BusinessRiskIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

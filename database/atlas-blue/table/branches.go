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

var Branches = newBranchesTable("atlas-blue", "branches", "")

type branchesTable struct {
	mysql.Table

	// Columns
	ID                      mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	EntityID                mysql.ColumnString // (DC2Type:guid)
	LeadPartner             mysql.ColumnString // (DC2Type:guid)
	Name                    mysql.ColumnString
	CountryCode             mysql.ColumnString
	Address                 mysql.ColumnString
	City                    mysql.ColumnString
	State                   mysql.ColumnString
	PostalCode              mysql.ColumnString
	PrimaryPhoneNumber      mysql.ColumnString
	PrimaryOtherPhoneNumber mysql.ColumnString
	DescriptionActivity     mysql.ColumnString
	NumberEmployes          mysql.ColumnString
	UniqueGlobalNumber      mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type BranchesTable struct {
	branchesTable

	NEW branchesTable
}

// AS creates new BranchesTable with assigned alias
func (a BranchesTable) AS(alias string) *BranchesTable {
	return newBranchesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new BranchesTable with assigned schema name
func (a BranchesTable) FromSchema(schemaName string) *BranchesTable {
	return newBranchesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new BranchesTable with assigned table prefix
func (a BranchesTable) WithPrefix(prefix string) *BranchesTable {
	return newBranchesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new BranchesTable with assigned table suffix
func (a BranchesTable) WithSuffix(suffix string) *BranchesTable {
	return newBranchesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newBranchesTable(schemaName, tableName, alias string) *BranchesTable {
	return &BranchesTable{
		branchesTable: newBranchesTableImpl(schemaName, tableName, alias),
		NEW:           newBranchesTableImpl("", "new", ""),
	}
}

func newBranchesTableImpl(schemaName, tableName, alias string) branchesTable {
	var (
		IDColumn                      = mysql.StringColumn("id")
		EntityIDColumn                = mysql.StringColumn("entity_id")
		LeadPartnerColumn             = mysql.StringColumn("lead_partner")
		NameColumn                    = mysql.StringColumn("name")
		CountryCodeColumn             = mysql.StringColumn("country_code")
		AddressColumn                 = mysql.StringColumn("address")
		CityColumn                    = mysql.StringColumn("city")
		StateColumn                   = mysql.StringColumn("state")
		PostalCodeColumn              = mysql.StringColumn("postal_code")
		PrimaryPhoneNumberColumn      = mysql.StringColumn("primary_phone_number")
		PrimaryOtherPhoneNumberColumn = mysql.StringColumn("primary_other_phone_number")
		DescriptionActivityColumn     = mysql.StringColumn("description_activity")
		NumberEmployesColumn          = mysql.StringColumn("number_employes")
		UniqueGlobalNumberColumn      = mysql.StringColumn("unique_global_number")
		allColumns                    = mysql.ColumnList{IDColumn, EntityIDColumn, LeadPartnerColumn, NameColumn, CountryCodeColumn, AddressColumn, CityColumn, StateColumn, PostalCodeColumn, PrimaryPhoneNumberColumn, PrimaryOtherPhoneNumberColumn, DescriptionActivityColumn, NumberEmployesColumn, UniqueGlobalNumberColumn}
		mutableColumns                = mysql.ColumnList{EntityIDColumn, LeadPartnerColumn, NameColumn, CountryCodeColumn, AddressColumn, CityColumn, StateColumn, PostalCodeColumn, PrimaryPhoneNumberColumn, PrimaryOtherPhoneNumberColumn, DescriptionActivityColumn, NumberEmployesColumn, UniqueGlobalNumberColumn}
	)

	return branchesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                      IDColumn,
		EntityID:                EntityIDColumn,
		LeadPartner:             LeadPartnerColumn,
		Name:                    NameColumn,
		CountryCode:             CountryCodeColumn,
		Address:                 AddressColumn,
		City:                    CityColumn,
		State:                   StateColumn,
		PostalCode:              PostalCodeColumn,
		PrimaryPhoneNumber:      PrimaryPhoneNumberColumn,
		PrimaryOtherPhoneNumber: PrimaryOtherPhoneNumberColumn,
		DescriptionActivity:     DescriptionActivityColumn,
		NumberEmployes:          NumberEmployesColumn,
		UniqueGlobalNumber:      UniqueGlobalNumberColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

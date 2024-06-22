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

var MasterBranches = newMasterBranchesTable("atlas-blue", "master_branches", "")

type masterBranchesTable struct {
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

type MasterBranchesTable struct {
	masterBranchesTable

	NEW masterBranchesTable
}

// AS creates new MasterBranchesTable with assigned alias
func (a MasterBranchesTable) AS(alias string) *MasterBranchesTable {
	return newMasterBranchesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new MasterBranchesTable with assigned schema name
func (a MasterBranchesTable) FromSchema(schemaName string) *MasterBranchesTable {
	return newMasterBranchesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new MasterBranchesTable with assigned table prefix
func (a MasterBranchesTable) WithPrefix(prefix string) *MasterBranchesTable {
	return newMasterBranchesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new MasterBranchesTable with assigned table suffix
func (a MasterBranchesTable) WithSuffix(suffix string) *MasterBranchesTable {
	return newMasterBranchesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newMasterBranchesTable(schemaName, tableName, alias string) *MasterBranchesTable {
	return &MasterBranchesTable{
		masterBranchesTable: newMasterBranchesTableImpl(schemaName, tableName, alias),
		NEW:                 newMasterBranchesTableImpl("", "new", ""),
	}
}

func newMasterBranchesTableImpl(schemaName, tableName, alias string) masterBranchesTable {
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

	return masterBranchesTable{
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

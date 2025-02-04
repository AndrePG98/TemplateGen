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

var SharedServiceCenter = newSharedServiceCenterTable("atlas-blue", "shared_service_center", "")

type sharedServiceCenterTable struct {
	mysql.Table

	// Columns
	ID                            mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	EntityID                      mysql.ColumnString // (DC2Type:guid)
	LeadPartner                   mysql.ColumnString // (DC2Type:guid)
	Name                          mysql.ColumnString
	CountryCode                   mysql.ColumnString
	Address                       mysql.ColumnString
	City                          mysql.ColumnString
	State                         mysql.ColumnString
	PostalCode                    mysql.ColumnString
	PrimaryPhoneNumber            mysql.ColumnString
	PrimaryOtherPhoneNumber       mysql.ColumnString
	DescriptionActivity           mysql.ColumnString
	NumberEmployees               mysql.ColumnString
	UniqueGlobalNumber            mysql.ColumnString
	RelevantServicesProvided      mysql.ColumnString // (DC2Type:array)
	OtherRelevantServicesProvided mysql.ColumnString // (DC2Type:array)

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type SharedServiceCenterTable struct {
	sharedServiceCenterTable

	NEW sharedServiceCenterTable
}

// AS creates new SharedServiceCenterTable with assigned alias
func (a SharedServiceCenterTable) AS(alias string) *SharedServiceCenterTable {
	return newSharedServiceCenterTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SharedServiceCenterTable with assigned schema name
func (a SharedServiceCenterTable) FromSchema(schemaName string) *SharedServiceCenterTable {
	return newSharedServiceCenterTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SharedServiceCenterTable with assigned table prefix
func (a SharedServiceCenterTable) WithPrefix(prefix string) *SharedServiceCenterTable {
	return newSharedServiceCenterTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SharedServiceCenterTable with assigned table suffix
func (a SharedServiceCenterTable) WithSuffix(suffix string) *SharedServiceCenterTable {
	return newSharedServiceCenterTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSharedServiceCenterTable(schemaName, tableName, alias string) *SharedServiceCenterTable {
	return &SharedServiceCenterTable{
		sharedServiceCenterTable: newSharedServiceCenterTableImpl(schemaName, tableName, alias),
		NEW:                      newSharedServiceCenterTableImpl("", "new", ""),
	}
}

func newSharedServiceCenterTableImpl(schemaName, tableName, alias string) sharedServiceCenterTable {
	var (
		IDColumn                            = mysql.StringColumn("id")
		EntityIDColumn                      = mysql.StringColumn("entity_id")
		LeadPartnerColumn                   = mysql.StringColumn("lead_partner")
		NameColumn                          = mysql.StringColumn("name")
		CountryCodeColumn                   = mysql.StringColumn("country_code")
		AddressColumn                       = mysql.StringColumn("address")
		CityColumn                          = mysql.StringColumn("city")
		StateColumn                         = mysql.StringColumn("state")
		PostalCodeColumn                    = mysql.StringColumn("postal_code")
		PrimaryPhoneNumberColumn            = mysql.StringColumn("primary_phone_number")
		PrimaryOtherPhoneNumberColumn       = mysql.StringColumn("primary_other_phone_number")
		DescriptionActivityColumn           = mysql.StringColumn("description_activity")
		NumberEmployeesColumn               = mysql.StringColumn("number_employees")
		UniqueGlobalNumberColumn            = mysql.StringColumn("unique_global_number")
		RelevantServicesProvidedColumn      = mysql.StringColumn("relevant_services_provided")
		OtherRelevantServicesProvidedColumn = mysql.StringColumn("other_relevant_services_provided")
		allColumns                          = mysql.ColumnList{IDColumn, EntityIDColumn, LeadPartnerColumn, NameColumn, CountryCodeColumn, AddressColumn, CityColumn, StateColumn, PostalCodeColumn, PrimaryPhoneNumberColumn, PrimaryOtherPhoneNumberColumn, DescriptionActivityColumn, NumberEmployeesColumn, UniqueGlobalNumberColumn, RelevantServicesProvidedColumn, OtherRelevantServicesProvidedColumn}
		mutableColumns                      = mysql.ColumnList{EntityIDColumn, LeadPartnerColumn, NameColumn, CountryCodeColumn, AddressColumn, CityColumn, StateColumn, PostalCodeColumn, PrimaryPhoneNumberColumn, PrimaryOtherPhoneNumberColumn, DescriptionActivityColumn, NumberEmployeesColumn, UniqueGlobalNumberColumn, RelevantServicesProvidedColumn, OtherRelevantServicesProvidedColumn}
	)

	return sharedServiceCenterTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                            IDColumn,
		EntityID:                      EntityIDColumn,
		LeadPartner:                   LeadPartnerColumn,
		Name:                          NameColumn,
		CountryCode:                   CountryCodeColumn,
		Address:                       AddressColumn,
		City:                          CityColumn,
		State:                         StateColumn,
		PostalCode:                    PostalCodeColumn,
		PrimaryPhoneNumber:            PrimaryPhoneNumberColumn,
		PrimaryOtherPhoneNumber:       PrimaryOtherPhoneNumberColumn,
		DescriptionActivity:           DescriptionActivityColumn,
		NumberEmployees:               NumberEmployeesColumn,
		UniqueGlobalNumber:            UniqueGlobalNumberColumn,
		RelevantServicesProvided:      RelevantServicesProvidedColumn,
		OtherRelevantServicesProvided: OtherRelevantServicesProvidedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

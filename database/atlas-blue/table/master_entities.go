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

var MasterEntities = newMasterEntitiesTable("atlas-blue", "master_entities", "")

type masterEntitiesTable struct {
	mysql.Table

	// Columns
	ID                                 mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	ParentEntity                       mysql.ColumnString // (DC2Type:guid)
	LeadPartnerID                      mysql.ColumnString // (DC2Type:guid)
	TypeID                             mysql.ColumnInteger
	CountryCode                        mysql.ColumnString
	Name                               mysql.ColumnString
	AbbreviatedName                    mysql.ColumnString
	BranchSscName                      mysql.ColumnString
	StatusID                           mysql.ColumnInteger
	TestEntityCategoryID               mysql.ColumnInteger
	PrimaryAddress                     mysql.ColumnString
	PrimaryCity                        mysql.ColumnString
	PrimaryStateProvinceCountry        mysql.ColumnString
	PrimaryPostalCode                  mysql.ColumnString
	PrimaryPhoneNumber                 mysql.ColumnString
	PrimaryOtherPhoneNumber            mysql.ColumnString
	SecondaryAddress                   mysql.ColumnString
	SecondaryCity                      mysql.ColumnString
	SecondaryStateProvinceCountry      mysql.ColumnString
	SecondaryPostalCode                mysql.ColumnString
	SecondaryPhoneNumber               mysql.ColumnString
	SecondaryOtherPhoneNumber          mysql.ColumnString
	HasNationalRegistration            mysql.ColumnBool
	NationalRegistrationNumber         mysql.ColumnString
	UniqueGlobalNumber                 mysql.ColumnString
	Description                        mysql.ColumnString
	LegalFormID                        mysql.ColumnInteger
	DescriptionOfActivity              mysql.ColumnString
	ListedCompany                      mysql.ColumnBool
	StockExchangeListing               mysql.ColumnString // (DC2Type:array)
	CountriesOfListingCodes            mysql.ColumnString // (DC2Type:array)
	Pie                                mysql.ColumnBool
	OtherRelevantServicesProvidedBySsc mysql.ColumnString // (DC2Type:array)
	PeriodEndDate                      mysql.ColumnDate
	HasPreparingConsolidatedAccounts   mysql.ColumnBool
	HasBusinessDivisions               mysql.ColumnBool
	HasBranches                        mysql.ColumnBool
	HasSharedServicesCentres           mysql.ColumnBool
	FrequencyOfConsolidation           mysql.ColumnString
	NumberOfEmployees                  mysql.ColumnString
	IsMigrated                         mysql.ColumnBool
	PackageToTestIdentifier            mysql.ColumnString
	PackageToTestVersion               mysql.ColumnString
	PackageToTestLastCommitHash        mysql.ColumnString
	IsValid                            mysql.ColumnBool
	Locked                             mysql.ColumnBool
	LockedAt                           mysql.ColumnDate // (DC2Type:date_immutable)
	Deleted                            mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type MasterEntitiesTable struct {
	masterEntitiesTable

	NEW masterEntitiesTable
}

// AS creates new MasterEntitiesTable with assigned alias
func (a MasterEntitiesTable) AS(alias string) *MasterEntitiesTable {
	return newMasterEntitiesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new MasterEntitiesTable with assigned schema name
func (a MasterEntitiesTable) FromSchema(schemaName string) *MasterEntitiesTable {
	return newMasterEntitiesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new MasterEntitiesTable with assigned table prefix
func (a MasterEntitiesTable) WithPrefix(prefix string) *MasterEntitiesTable {
	return newMasterEntitiesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new MasterEntitiesTable with assigned table suffix
func (a MasterEntitiesTable) WithSuffix(suffix string) *MasterEntitiesTable {
	return newMasterEntitiesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newMasterEntitiesTable(schemaName, tableName, alias string) *MasterEntitiesTable {
	return &MasterEntitiesTable{
		masterEntitiesTable: newMasterEntitiesTableImpl(schemaName, tableName, alias),
		NEW:                 newMasterEntitiesTableImpl("", "new", ""),
	}
}

func newMasterEntitiesTableImpl(schemaName, tableName, alias string) masterEntitiesTable {
	var (
		IDColumn                                 = mysql.StringColumn("id")
		ParentEntityColumn                       = mysql.StringColumn("parent_entity")
		LeadPartnerIDColumn                      = mysql.StringColumn("lead_partner_id")
		TypeIDColumn                             = mysql.IntegerColumn("type_id")
		CountryCodeColumn                        = mysql.StringColumn("country_code")
		NameColumn                               = mysql.StringColumn("name")
		AbbreviatedNameColumn                    = mysql.StringColumn("abbreviated_name")
		BranchSscNameColumn                      = mysql.StringColumn("branch_ssc_name")
		StatusIDColumn                           = mysql.IntegerColumn("status_id")
		TestEntityCategoryIDColumn               = mysql.IntegerColumn("test_entity_category_id")
		PrimaryAddressColumn                     = mysql.StringColumn("primary_address")
		PrimaryCityColumn                        = mysql.StringColumn("primary_city")
		PrimaryStateProvinceCountryColumn        = mysql.StringColumn("primary_state_province_country")
		PrimaryPostalCodeColumn                  = mysql.StringColumn("primary_postal_code")
		PrimaryPhoneNumberColumn                 = mysql.StringColumn("primary_phone_number")
		PrimaryOtherPhoneNumberColumn            = mysql.StringColumn("primary_other_phone_number")
		SecondaryAddressColumn                   = mysql.StringColumn("secondary_address")
		SecondaryCityColumn                      = mysql.StringColumn("secondary_city")
		SecondaryStateProvinceCountryColumn      = mysql.StringColumn("secondary_state_province_country")
		SecondaryPostalCodeColumn                = mysql.StringColumn("secondary_postal_code")
		SecondaryPhoneNumberColumn               = mysql.StringColumn("secondary_phone_number")
		SecondaryOtherPhoneNumberColumn          = mysql.StringColumn("secondary_other_phone_number")
		HasNationalRegistrationColumn            = mysql.BoolColumn("has_national_registration")
		NationalRegistrationNumberColumn         = mysql.StringColumn("national_registration_number")
		UniqueGlobalNumberColumn                 = mysql.StringColumn("unique_global_number")
		DescriptionColumn                        = mysql.StringColumn("description")
		LegalFormIDColumn                        = mysql.IntegerColumn("legal_form_id")
		DescriptionOfActivityColumn              = mysql.StringColumn("description_of_activity")
		ListedCompanyColumn                      = mysql.BoolColumn("listed_company")
		StockExchangeListingColumn               = mysql.StringColumn("stock_exchange_listing")
		CountriesOfListingCodesColumn            = mysql.StringColumn("countries_of_listing_codes")
		PieColumn                                = mysql.BoolColumn("pie")
		OtherRelevantServicesProvidedBySscColumn = mysql.StringColumn("other_relevant_services_provided_by_ssc")
		PeriodEndDateColumn                      = mysql.DateColumn("period_end_date")
		HasPreparingConsolidatedAccountsColumn   = mysql.BoolColumn("has_preparing_consolidated_accounts")
		HasBusinessDivisionsColumn               = mysql.BoolColumn("has_business_divisions")
		HasBranchesColumn                        = mysql.BoolColumn("has_branches")
		HasSharedServicesCentresColumn           = mysql.BoolColumn("has_shared_services_centres")
		FrequencyOfConsolidationColumn           = mysql.StringColumn("frequency_of_consolidation")
		NumberOfEmployeesColumn                  = mysql.StringColumn("number_of_employees")
		IsMigratedColumn                         = mysql.BoolColumn("is_migrated")
		PackageToTestIdentifierColumn            = mysql.StringColumn("package_to_test_identifier")
		PackageToTestVersionColumn               = mysql.StringColumn("package_to_test_version")
		PackageToTestLastCommitHashColumn        = mysql.StringColumn("package_to_test_last_commit_hash")
		IsValidColumn                            = mysql.BoolColumn("is_valid")
		LockedColumn                             = mysql.BoolColumn("locked")
		LockedAtColumn                           = mysql.DateColumn("locked_at")
		DeletedColumn                            = mysql.IntegerColumn("deleted")
		allColumns                               = mysql.ColumnList{IDColumn, ParentEntityColumn, LeadPartnerIDColumn, TypeIDColumn, CountryCodeColumn, NameColumn, AbbreviatedNameColumn, BranchSscNameColumn, StatusIDColumn, TestEntityCategoryIDColumn, PrimaryAddressColumn, PrimaryCityColumn, PrimaryStateProvinceCountryColumn, PrimaryPostalCodeColumn, PrimaryPhoneNumberColumn, PrimaryOtherPhoneNumberColumn, SecondaryAddressColumn, SecondaryCityColumn, SecondaryStateProvinceCountryColumn, SecondaryPostalCodeColumn, SecondaryPhoneNumberColumn, SecondaryOtherPhoneNumberColumn, HasNationalRegistrationColumn, NationalRegistrationNumberColumn, UniqueGlobalNumberColumn, DescriptionColumn, LegalFormIDColumn, DescriptionOfActivityColumn, ListedCompanyColumn, StockExchangeListingColumn, CountriesOfListingCodesColumn, PieColumn, OtherRelevantServicesProvidedBySscColumn, PeriodEndDateColumn, HasPreparingConsolidatedAccountsColumn, HasBusinessDivisionsColumn, HasBranchesColumn, HasSharedServicesCentresColumn, FrequencyOfConsolidationColumn, NumberOfEmployeesColumn, IsMigratedColumn, PackageToTestIdentifierColumn, PackageToTestVersionColumn, PackageToTestLastCommitHashColumn, IsValidColumn, LockedColumn, LockedAtColumn, DeletedColumn}
		mutableColumns                           = mysql.ColumnList{ParentEntityColumn, LeadPartnerIDColumn, TypeIDColumn, CountryCodeColumn, NameColumn, AbbreviatedNameColumn, BranchSscNameColumn, StatusIDColumn, TestEntityCategoryIDColumn, PrimaryAddressColumn, PrimaryCityColumn, PrimaryStateProvinceCountryColumn, PrimaryPostalCodeColumn, PrimaryPhoneNumberColumn, PrimaryOtherPhoneNumberColumn, SecondaryAddressColumn, SecondaryCityColumn, SecondaryStateProvinceCountryColumn, SecondaryPostalCodeColumn, SecondaryPhoneNumberColumn, SecondaryOtherPhoneNumberColumn, HasNationalRegistrationColumn, NationalRegistrationNumberColumn, UniqueGlobalNumberColumn, DescriptionColumn, LegalFormIDColumn, DescriptionOfActivityColumn, ListedCompanyColumn, StockExchangeListingColumn, CountriesOfListingCodesColumn, PieColumn, OtherRelevantServicesProvidedBySscColumn, PeriodEndDateColumn, HasPreparingConsolidatedAccountsColumn, HasBusinessDivisionsColumn, HasBranchesColumn, HasSharedServicesCentresColumn, FrequencyOfConsolidationColumn, NumberOfEmployeesColumn, IsMigratedColumn, PackageToTestIdentifierColumn, PackageToTestVersionColumn, PackageToTestLastCommitHashColumn, IsValidColumn, LockedColumn, LockedAtColumn, DeletedColumn}
	)

	return masterEntitiesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                                 IDColumn,
		ParentEntity:                       ParentEntityColumn,
		LeadPartnerID:                      LeadPartnerIDColumn,
		TypeID:                             TypeIDColumn,
		CountryCode:                        CountryCodeColumn,
		Name:                               NameColumn,
		AbbreviatedName:                    AbbreviatedNameColumn,
		BranchSscName:                      BranchSscNameColumn,
		StatusID:                           StatusIDColumn,
		TestEntityCategoryID:               TestEntityCategoryIDColumn,
		PrimaryAddress:                     PrimaryAddressColumn,
		PrimaryCity:                        PrimaryCityColumn,
		PrimaryStateProvinceCountry:        PrimaryStateProvinceCountryColumn,
		PrimaryPostalCode:                  PrimaryPostalCodeColumn,
		PrimaryPhoneNumber:                 PrimaryPhoneNumberColumn,
		PrimaryOtherPhoneNumber:            PrimaryOtherPhoneNumberColumn,
		SecondaryAddress:                   SecondaryAddressColumn,
		SecondaryCity:                      SecondaryCityColumn,
		SecondaryStateProvinceCountry:      SecondaryStateProvinceCountryColumn,
		SecondaryPostalCode:                SecondaryPostalCodeColumn,
		SecondaryPhoneNumber:               SecondaryPhoneNumberColumn,
		SecondaryOtherPhoneNumber:          SecondaryOtherPhoneNumberColumn,
		HasNationalRegistration:            HasNationalRegistrationColumn,
		NationalRegistrationNumber:         NationalRegistrationNumberColumn,
		UniqueGlobalNumber:                 UniqueGlobalNumberColumn,
		Description:                        DescriptionColumn,
		LegalFormID:                        LegalFormIDColumn,
		DescriptionOfActivity:              DescriptionOfActivityColumn,
		ListedCompany:                      ListedCompanyColumn,
		StockExchangeListing:               StockExchangeListingColumn,
		CountriesOfListingCodes:            CountriesOfListingCodesColumn,
		Pie:                                PieColumn,
		OtherRelevantServicesProvidedBySsc: OtherRelevantServicesProvidedBySscColumn,
		PeriodEndDate:                      PeriodEndDateColumn,
		HasPreparingConsolidatedAccounts:   HasPreparingConsolidatedAccountsColumn,
		HasBusinessDivisions:               HasBusinessDivisionsColumn,
		HasBranches:                        HasBranchesColumn,
		HasSharedServicesCentres:           HasSharedServicesCentresColumn,
		FrequencyOfConsolidation:           FrequencyOfConsolidationColumn,
		NumberOfEmployees:                  NumberOfEmployeesColumn,
		IsMigrated:                         IsMigratedColumn,
		PackageToTestIdentifier:            PackageToTestIdentifierColumn,
		PackageToTestVersion:               PackageToTestVersionColumn,
		PackageToTestLastCommitHash:        PackageToTestLastCommitHashColumn,
		IsValid:                            IsValidColumn,
		Locked:                             LockedColumn,
		LockedAt:                           LockedAtColumn,
		Deleted:                            DeletedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

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

var EngagementsRollforwards = newEngagementsRollforwardsTable("atlas-blue", "engagements_rollforwards", "")

type engagementsRollforwardsTable struct {
	mysql.Table

	// Columns
	ID                            mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	EngagementID                  mysql.ColumnString // (DC2Type:guid)
	CreatedBy                     mysql.ColumnString // (DC2Type:guid)
	TargetEngagementLeadPartner   mysql.ColumnString // (DC2Type:guid)
	TargetEngagementType          mysql.ColumnString
	Status                        mysql.ColumnString
	CreatedAt                     mysql.ColumnTimestamp // (DC2Type:datetime_immutable)
	TargetEngagementPeriodEndDate mysql.ColumnTimestamp // (DC2Type:datetime_immutable)
	TargetEngagementID            mysql.ColumnString    // @JsonFormat(property="id") @JsonInclude(DC2Type:guid)

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type EngagementsRollforwardsTable struct {
	engagementsRollforwardsTable

	NEW engagementsRollforwardsTable
}

// AS creates new EngagementsRollforwardsTable with assigned alias
func (a EngagementsRollforwardsTable) AS(alias string) *EngagementsRollforwardsTable {
	return newEngagementsRollforwardsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new EngagementsRollforwardsTable with assigned schema name
func (a EngagementsRollforwardsTable) FromSchema(schemaName string) *EngagementsRollforwardsTable {
	return newEngagementsRollforwardsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new EngagementsRollforwardsTable with assigned table prefix
func (a EngagementsRollforwardsTable) WithPrefix(prefix string) *EngagementsRollforwardsTable {
	return newEngagementsRollforwardsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new EngagementsRollforwardsTable with assigned table suffix
func (a EngagementsRollforwardsTable) WithSuffix(suffix string) *EngagementsRollforwardsTable {
	return newEngagementsRollforwardsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newEngagementsRollforwardsTable(schemaName, tableName, alias string) *EngagementsRollforwardsTable {
	return &EngagementsRollforwardsTable{
		engagementsRollforwardsTable: newEngagementsRollforwardsTableImpl(schemaName, tableName, alias),
		NEW:                          newEngagementsRollforwardsTableImpl("", "new", ""),
	}
}

func newEngagementsRollforwardsTableImpl(schemaName, tableName, alias string) engagementsRollforwardsTable {
	var (
		IDColumn                            = mysql.StringColumn("id")
		EngagementIDColumn                  = mysql.StringColumn("engagement_id")
		CreatedByColumn                     = mysql.StringColumn("created_by")
		TargetEngagementLeadPartnerColumn   = mysql.StringColumn("target_engagement_lead_partner")
		TargetEngagementTypeColumn          = mysql.StringColumn("target_engagement_type")
		StatusColumn                        = mysql.StringColumn("status")
		CreatedAtColumn                     = mysql.TimestampColumn("created_at")
		TargetEngagementPeriodEndDateColumn = mysql.TimestampColumn("target_engagement_period_end_date")
		TargetEngagementIDColumn            = mysql.StringColumn("target_engagement_id")
		allColumns                          = mysql.ColumnList{IDColumn, EngagementIDColumn, CreatedByColumn, TargetEngagementLeadPartnerColumn, TargetEngagementTypeColumn, StatusColumn, CreatedAtColumn, TargetEngagementPeriodEndDateColumn, TargetEngagementIDColumn}
		mutableColumns                      = mysql.ColumnList{EngagementIDColumn, CreatedByColumn, TargetEngagementLeadPartnerColumn, TargetEngagementTypeColumn, StatusColumn, CreatedAtColumn, TargetEngagementPeriodEndDateColumn, TargetEngagementIDColumn}
	)

	return engagementsRollforwardsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                            IDColumn,
		EngagementID:                  EngagementIDColumn,
		CreatedBy:                     CreatedByColumn,
		TargetEngagementLeadPartner:   TargetEngagementLeadPartnerColumn,
		TargetEngagementType:          TargetEngagementTypeColumn,
		Status:                        StatusColumn,
		CreatedAt:                     CreatedAtColumn,
		TargetEngagementPeriodEndDate: TargetEngagementPeriodEndDateColumn,
		TargetEngagementID:            TargetEngagementIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

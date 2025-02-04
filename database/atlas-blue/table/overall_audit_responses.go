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

var OverallAuditResponses = newOverallAuditResponsesTable("atlas-blue", "overall_audit_responses", "")

type overallAuditResponsesTable struct {
	mysql.Table

	// Columns
	ID                                       mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	EngagementID                             mysql.ColumnString // (DC2Type:guid)
	CreatedBy                                mysql.ColumnString // (DC2Type:guid)
	DoYouNeedToEmphasize                     mysql.ColumnBool
	CommentEmphasize                         mysql.ColumnString
	DoYouNeedToAssignMoreExperiencedStaff    mysql.ColumnBool
	CommentExperiencedStaff                  mysql.ColumnString
	DoYouNeedToProvideMoreSupervision        mysql.ColumnBool
	CommentMoreSupervision                   mysql.ColumnString
	DoYouNeedToIncorporateAdditionalElements mysql.ColumnBool
	CommentAdditionalElements                mysql.ColumnString
	DoYouNeedToMakeGeneralChanges            mysql.ColumnBool
	CommentGeneralChanges                    mysql.ColumnString
	Status                                   mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type OverallAuditResponsesTable struct {
	overallAuditResponsesTable

	NEW overallAuditResponsesTable
}

// AS creates new OverallAuditResponsesTable with assigned alias
func (a OverallAuditResponsesTable) AS(alias string) *OverallAuditResponsesTable {
	return newOverallAuditResponsesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new OverallAuditResponsesTable with assigned schema name
func (a OverallAuditResponsesTable) FromSchema(schemaName string) *OverallAuditResponsesTable {
	return newOverallAuditResponsesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new OverallAuditResponsesTable with assigned table prefix
func (a OverallAuditResponsesTable) WithPrefix(prefix string) *OverallAuditResponsesTable {
	return newOverallAuditResponsesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new OverallAuditResponsesTable with assigned table suffix
func (a OverallAuditResponsesTable) WithSuffix(suffix string) *OverallAuditResponsesTable {
	return newOverallAuditResponsesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newOverallAuditResponsesTable(schemaName, tableName, alias string) *OverallAuditResponsesTable {
	return &OverallAuditResponsesTable{
		overallAuditResponsesTable: newOverallAuditResponsesTableImpl(schemaName, tableName, alias),
		NEW:                        newOverallAuditResponsesTableImpl("", "new", ""),
	}
}

func newOverallAuditResponsesTableImpl(schemaName, tableName, alias string) overallAuditResponsesTable {
	var (
		IDColumn                                       = mysql.StringColumn("id")
		EngagementIDColumn                             = mysql.StringColumn("engagement_id")
		CreatedByColumn                                = mysql.StringColumn("created_by")
		DoYouNeedToEmphasizeColumn                     = mysql.BoolColumn("do_you_need_to_emphasize")
		CommentEmphasizeColumn                         = mysql.StringColumn("comment_emphasize")
		DoYouNeedToAssignMoreExperiencedStaffColumn    = mysql.BoolColumn("do_you_need_to_assign_more_experienced_staff")
		CommentExperiencedStaffColumn                  = mysql.StringColumn("comment_experienced_staff")
		DoYouNeedToProvideMoreSupervisionColumn        = mysql.BoolColumn("do_you_need_to_provide_more_supervision")
		CommentMoreSupervisionColumn                   = mysql.StringColumn("comment_more_supervision")
		DoYouNeedToIncorporateAdditionalElementsColumn = mysql.BoolColumn("do_you_need_to_incorporate_additional_elements")
		CommentAdditionalElementsColumn                = mysql.StringColumn("comment_additional_elements")
		DoYouNeedToMakeGeneralChangesColumn            = mysql.BoolColumn("do_you_need_to_make_general_changes")
		CommentGeneralChangesColumn                    = mysql.StringColumn("comment_general_changes")
		StatusColumn                                   = mysql.StringColumn("status")
		allColumns                                     = mysql.ColumnList{IDColumn, EngagementIDColumn, CreatedByColumn, DoYouNeedToEmphasizeColumn, CommentEmphasizeColumn, DoYouNeedToAssignMoreExperiencedStaffColumn, CommentExperiencedStaffColumn, DoYouNeedToProvideMoreSupervisionColumn, CommentMoreSupervisionColumn, DoYouNeedToIncorporateAdditionalElementsColumn, CommentAdditionalElementsColumn, DoYouNeedToMakeGeneralChangesColumn, CommentGeneralChangesColumn, StatusColumn}
		mutableColumns                                 = mysql.ColumnList{EngagementIDColumn, CreatedByColumn, DoYouNeedToEmphasizeColumn, CommentEmphasizeColumn, DoYouNeedToAssignMoreExperiencedStaffColumn, CommentExperiencedStaffColumn, DoYouNeedToProvideMoreSupervisionColumn, CommentMoreSupervisionColumn, DoYouNeedToIncorporateAdditionalElementsColumn, CommentAdditionalElementsColumn, DoYouNeedToMakeGeneralChangesColumn, CommentGeneralChangesColumn, StatusColumn}
	)

	return overallAuditResponsesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                                       IDColumn,
		EngagementID:                             EngagementIDColumn,
		CreatedBy:                                CreatedByColumn,
		DoYouNeedToEmphasize:                     DoYouNeedToEmphasizeColumn,
		CommentEmphasize:                         CommentEmphasizeColumn,
		DoYouNeedToAssignMoreExperiencedStaff:    DoYouNeedToAssignMoreExperiencedStaffColumn,
		CommentExperiencedStaff:                  CommentExperiencedStaffColumn,
		DoYouNeedToProvideMoreSupervision:        DoYouNeedToProvideMoreSupervisionColumn,
		CommentMoreSupervision:                   CommentMoreSupervisionColumn,
		DoYouNeedToIncorporateAdditionalElements: DoYouNeedToIncorporateAdditionalElementsColumn,
		CommentAdditionalElements:                CommentAdditionalElementsColumn,
		DoYouNeedToMakeGeneralChanges:            DoYouNeedToMakeGeneralChangesColumn,
		CommentGeneralChanges:                    CommentGeneralChangesColumn,
		Status:                                   StatusColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

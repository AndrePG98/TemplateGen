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

var DeletionsRetrievalRequests = newDeletionsRetrievalRequestsTable("atlas-blue", "deletions_retrieval_requests", "")

type deletionsRetrievalRequestsTable struct {
	mysql.Table

	// Columns
	ID                         mysql.ColumnString // @UUID("v4")(DC2Type:guid)
	EntitiesDeletionRequestID  mysql.ColumnString // (DC2Type:guid)
	RequestPartnerApprovalFrom mysql.ColumnString // (DC2Type:guid)
	UserRequestingRetrieval    mysql.ColumnString // (DC2Type:guid)
	InitiatedBy                mysql.ColumnString // (DC2Type:guid)
	ApprovedBy                 mysql.ColumnString // (DC2Type:guid)
	ReasonForRetrievalID       mysql.ColumnInteger
	AdditionnalComment         mysql.ColumnString
	InitiatedOn                mysql.ColumnDate
	ApprovedOn                 mysql.ColumnDate
	ReasonForRejection         mysql.ColumnString
	ReasonForPartnerChange     mysql.ColumnString
	RetrievalStatus            mysql.ColumnInteger
	Status                     mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type DeletionsRetrievalRequestsTable struct {
	deletionsRetrievalRequestsTable

	NEW deletionsRetrievalRequestsTable
}

// AS creates new DeletionsRetrievalRequestsTable with assigned alias
func (a DeletionsRetrievalRequestsTable) AS(alias string) *DeletionsRetrievalRequestsTable {
	return newDeletionsRetrievalRequestsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new DeletionsRetrievalRequestsTable with assigned schema name
func (a DeletionsRetrievalRequestsTable) FromSchema(schemaName string) *DeletionsRetrievalRequestsTable {
	return newDeletionsRetrievalRequestsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new DeletionsRetrievalRequestsTable with assigned table prefix
func (a DeletionsRetrievalRequestsTable) WithPrefix(prefix string) *DeletionsRetrievalRequestsTable {
	return newDeletionsRetrievalRequestsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new DeletionsRetrievalRequestsTable with assigned table suffix
func (a DeletionsRetrievalRequestsTable) WithSuffix(suffix string) *DeletionsRetrievalRequestsTable {
	return newDeletionsRetrievalRequestsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newDeletionsRetrievalRequestsTable(schemaName, tableName, alias string) *DeletionsRetrievalRequestsTable {
	return &DeletionsRetrievalRequestsTable{
		deletionsRetrievalRequestsTable: newDeletionsRetrievalRequestsTableImpl(schemaName, tableName, alias),
		NEW:                             newDeletionsRetrievalRequestsTableImpl("", "new", ""),
	}
}

func newDeletionsRetrievalRequestsTableImpl(schemaName, tableName, alias string) deletionsRetrievalRequestsTable {
	var (
		IDColumn                         = mysql.StringColumn("id")
		EntitiesDeletionRequestIDColumn  = mysql.StringColumn("entities_deletion_request_id")
		RequestPartnerApprovalFromColumn = mysql.StringColumn("request_partner_approval_from")
		UserRequestingRetrievalColumn    = mysql.StringColumn("user_requesting_retrieval")
		InitiatedByColumn                = mysql.StringColumn("initiated_by")
		ApprovedByColumn                 = mysql.StringColumn("approved_by")
		ReasonForRetrievalIDColumn       = mysql.IntegerColumn("reason_for_retrieval_id")
		AdditionnalCommentColumn         = mysql.StringColumn("additionnal_comment")
		InitiatedOnColumn                = mysql.DateColumn("initiated_on")
		ApprovedOnColumn                 = mysql.DateColumn("approved_on")
		ReasonForRejectionColumn         = mysql.StringColumn("reason_for_rejection")
		ReasonForPartnerChangeColumn     = mysql.StringColumn("reason_for_partner_change")
		RetrievalStatusColumn            = mysql.IntegerColumn("retrieval_status")
		StatusColumn                     = mysql.StringColumn("status")
		allColumns                       = mysql.ColumnList{IDColumn, EntitiesDeletionRequestIDColumn, RequestPartnerApprovalFromColumn, UserRequestingRetrievalColumn, InitiatedByColumn, ApprovedByColumn, ReasonForRetrievalIDColumn, AdditionnalCommentColumn, InitiatedOnColumn, ApprovedOnColumn, ReasonForRejectionColumn, ReasonForPartnerChangeColumn, RetrievalStatusColumn, StatusColumn}
		mutableColumns                   = mysql.ColumnList{EntitiesDeletionRequestIDColumn, RequestPartnerApprovalFromColumn, UserRequestingRetrievalColumn, InitiatedByColumn, ApprovedByColumn, ReasonForRetrievalIDColumn, AdditionnalCommentColumn, InitiatedOnColumn, ApprovedOnColumn, ReasonForRejectionColumn, ReasonForPartnerChangeColumn, RetrievalStatusColumn, StatusColumn}
	)

	return deletionsRetrievalRequestsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                         IDColumn,
		EntitiesDeletionRequestID:  EntitiesDeletionRequestIDColumn,
		RequestPartnerApprovalFrom: RequestPartnerApprovalFromColumn,
		UserRequestingRetrieval:    UserRequestingRetrievalColumn,
		InitiatedBy:                InitiatedByColumn,
		ApprovedBy:                 ApprovedByColumn,
		ReasonForRetrievalID:       ReasonForRetrievalIDColumn,
		AdditionnalComment:         AdditionnalCommentColumn,
		InitiatedOn:                InitiatedOnColumn,
		ApprovedOn:                 ApprovedOnColumn,
		ReasonForRejection:         ReasonForRejectionColumn,
		ReasonForPartnerChange:     ReasonForPartnerChangeColumn,
		RetrievalStatus:            RetrievalStatusColumn,
		Status:                     StatusColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

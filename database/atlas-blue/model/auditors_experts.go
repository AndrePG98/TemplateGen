//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type AuditorsExperts struct {
	ID                                          string `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	EngagementID                                string // (DC2Type:guid)
	ExpertName                                  *string
	FieldOfExpertise                            *string
	NatureOfTheExpertsWork                      *string
	SignificanceOfTheExpert                     *string
	KnowledgeOfAndExperience                    *string
	IsTheExpertAnAuditorsInternal               *bool
	ExpertAnAuditorsInternalComment             *string
	IsTheExpertSubjectToOurOwnInternalQuality   *bool
	ExpertSubjectToOurOwnInternalQualityComment *string
	EvaluateCompetenceAndCapabilities           *string
	HaveYouInquiredOfInterestsRelationships     *bool
	InquiredOfInterestsRelationshipsComment     *string
	IsTheExpertObjective                        *bool
	DescribeObjectivityIssues                   *string
	NatureScopeObjectivesExpertsWork            *string
	HaveAgreed                                  *bool
	CommentOnAgreement                          *string
	Status                                      string
	Deleted                                     int32
	AnySafeguardsThatEliminateTheThreat         *bool
	DescribeTheSafeguards                       *string
}

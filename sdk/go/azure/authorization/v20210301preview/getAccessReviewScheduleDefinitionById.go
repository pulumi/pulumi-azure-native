


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccessReviewScheduleDefinitionById(ctx *pulumi.Context, args *LookupAccessReviewScheduleDefinitionByIdArgs, opts ...pulumi.InvokeOption) (*LookupAccessReviewScheduleDefinitionByIdResult, error) {
	var rv LookupAccessReviewScheduleDefinitionByIdResult
	err := ctx.Invoke("azure-native:authorization/v20210301preview:getAccessReviewScheduleDefinitionById", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessReviewScheduleDefinitionByIdArgs struct {
	ScheduleDefinitionId string `pulumi:"scheduleDefinitionId"`
}


type LookupAccessReviewScheduleDefinitionByIdResult struct {
	AssignmentState                 string                         `pulumi:"assignmentState"`
	AutoApplyDecisionsEnabled       *bool                          `pulumi:"autoApplyDecisionsEnabled"`
	BackupReviewers                 []AccessReviewReviewerResponse `pulumi:"backupReviewers"`
	DefaultDecision                 *string                        `pulumi:"defaultDecision"`
	DefaultDecisionEnabled          *bool                          `pulumi:"defaultDecisionEnabled"`
	DescriptionForAdmins            *string                        `pulumi:"descriptionForAdmins"`
	DescriptionForReviewers         *string                        `pulumi:"descriptionForReviewers"`
	DisplayName                     *string                        `pulumi:"displayName"`
	EndDate                         *string                        `pulumi:"endDate"`
	Id                              string                         `pulumi:"id"`
	InactiveDuration                *string                        `pulumi:"inactiveDuration"`
	InstanceDurationInDays          *int                           `pulumi:"instanceDurationInDays"`
	Instances                       []AccessReviewInstanceResponse `pulumi:"instances"`
	Interval                        *int                           `pulumi:"interval"`
	JustificationRequiredOnApproval *bool                          `pulumi:"justificationRequiredOnApproval"`
	MailNotificationsEnabled        *bool                          `pulumi:"mailNotificationsEnabled"`
	Name                            string                         `pulumi:"name"`
	NumberOfOccurrences             *int                           `pulumi:"numberOfOccurrences"`
	PrincipalId                     string                         `pulumi:"principalId"`
	PrincipalName                   string                         `pulumi:"principalName"`
	PrincipalType                   string                         `pulumi:"principalType"`
	RecommendationsEnabled          *bool                          `pulumi:"recommendationsEnabled"`
	ReminderNotificationsEnabled    *bool                          `pulumi:"reminderNotificationsEnabled"`
	ResourceId                      string                         `pulumi:"resourceId"`
	Reviewers                       []AccessReviewReviewerResponse `pulumi:"reviewers"`
	ReviewersType                   string                         `pulumi:"reviewersType"`
	RoleDefinitionId                string                         `pulumi:"roleDefinitionId"`
	StartDate                       *string                        `pulumi:"startDate"`
	Status                          string                         `pulumi:"status"`
	Type                            string                         `pulumi:"type"`
	UserPrincipalName               string                         `pulumi:"userPrincipalName"`
}

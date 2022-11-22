


package v20211116preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccessReviewScheduleDefinitionById(ctx *pulumi.Context, args *LookupAccessReviewScheduleDefinitionByIdArgs, opts ...pulumi.InvokeOption) (*LookupAccessReviewScheduleDefinitionByIdResult, error) {
	var rv LookupAccessReviewScheduleDefinitionByIdResult
	err := ctx.Invoke("azure-native:authorization/v20211116preview:getAccessReviewScheduleDefinitionById", args, &rv, opts...)
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
	ExpandNestedMemberships         *bool                          `pulumi:"expandNestedMemberships"`
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
	RecommendationLookBackDuration  *string                        `pulumi:"recommendationLookBackDuration"`
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

func LookupAccessReviewScheduleDefinitionByIdOutput(ctx *pulumi.Context, args LookupAccessReviewScheduleDefinitionByIdOutputArgs, opts ...pulumi.InvokeOption) LookupAccessReviewScheduleDefinitionByIdResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessReviewScheduleDefinitionByIdResult, error) {
			args := v.(LookupAccessReviewScheduleDefinitionByIdArgs)
			r, err := LookupAccessReviewScheduleDefinitionById(ctx, &args, opts...)
			var s LookupAccessReviewScheduleDefinitionByIdResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccessReviewScheduleDefinitionByIdResultOutput)
}

type LookupAccessReviewScheduleDefinitionByIdOutputArgs struct {
	ScheduleDefinitionId pulumi.StringInput `pulumi:"scheduleDefinitionId"`
}

func (LookupAccessReviewScheduleDefinitionByIdOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessReviewScheduleDefinitionByIdArgs)(nil)).Elem()
}


type LookupAccessReviewScheduleDefinitionByIdResultOutput struct{ *pulumi.OutputState }

func (LookupAccessReviewScheduleDefinitionByIdResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessReviewScheduleDefinitionByIdResult)(nil)).Elem()
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) ToLookupAccessReviewScheduleDefinitionByIdResultOutput() LookupAccessReviewScheduleDefinitionByIdResultOutput {
	return o
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) ToLookupAccessReviewScheduleDefinitionByIdResultOutputWithContext(ctx context.Context) LookupAccessReviewScheduleDefinitionByIdResultOutput {
	return o
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) AssignmentState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) string { return v.AssignmentState }).(pulumi.StringOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) AutoApplyDecisionsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) *bool { return v.AutoApplyDecisionsEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) BackupReviewers() AccessReviewReviewerResponseArrayOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) []AccessReviewReviewerResponse {
		return v.BackupReviewers
	}).(AccessReviewReviewerResponseArrayOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) DefaultDecision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) *string { return v.DefaultDecision }).(pulumi.StringPtrOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) DefaultDecisionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) *bool { return v.DefaultDecisionEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) DescriptionForAdmins() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) *string { return v.DescriptionForAdmins }).(pulumi.StringPtrOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) DescriptionForReviewers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) *string { return v.DescriptionForReviewers }).(pulumi.StringPtrOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) ExpandNestedMemberships() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) *bool { return v.ExpandNestedMemberships }).(pulumi.BoolPtrOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) InactiveDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) *string { return v.InactiveDuration }).(pulumi.StringPtrOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) InstanceDurationInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) *int { return v.InstanceDurationInDays }).(pulumi.IntPtrOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) Instances() AccessReviewInstanceResponseArrayOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) []AccessReviewInstanceResponse {
		return v.Instances
	}).(AccessReviewInstanceResponseArrayOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) JustificationRequiredOnApproval() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) *bool { return v.JustificationRequiredOnApproval }).(pulumi.BoolPtrOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) MailNotificationsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) *bool { return v.MailNotificationsEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) NumberOfOccurrences() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) *int { return v.NumberOfOccurrences }).(pulumi.IntPtrOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) string { return v.PrincipalName }).(pulumi.StringOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) string { return v.PrincipalType }).(pulumi.StringOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) RecommendationLookBackDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) *string {
		return v.RecommendationLookBackDuration
	}).(pulumi.StringPtrOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) RecommendationsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) *bool { return v.RecommendationsEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) ReminderNotificationsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) *bool { return v.ReminderNotificationsEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) Reviewers() AccessReviewReviewerResponseArrayOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) []AccessReviewReviewerResponse {
		return v.Reviewers
	}).(AccessReviewReviewerResponseArrayOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) ReviewersType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) string { return v.ReviewersType }).(pulumi.StringOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) *string { return v.StartDate }).(pulumi.StringPtrOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAccessReviewScheduleDefinitionByIdResultOutput) UserPrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessReviewScheduleDefinitionByIdResult) string { return v.UserPrincipalName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessReviewScheduleDefinitionByIdResultOutput{})
}

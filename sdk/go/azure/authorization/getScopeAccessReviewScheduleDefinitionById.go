


package authorization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScopeAccessReviewScheduleDefinitionById(ctx *pulumi.Context, args *LookupScopeAccessReviewScheduleDefinitionByIdArgs, opts ...pulumi.InvokeOption) (*LookupScopeAccessReviewScheduleDefinitionByIdResult, error) {
	var rv LookupScopeAccessReviewScheduleDefinitionByIdResult
	err := ctx.Invoke("azure-native:authorization:getScopeAccessReviewScheduleDefinitionById", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScopeAccessReviewScheduleDefinitionByIdArgs struct {
	ScheduleDefinitionId string `pulumi:"scheduleDefinitionId"`
	Scope                string `pulumi:"scope"`
}


type LookupScopeAccessReviewScheduleDefinitionByIdResult struct {
	AssignmentState                 string                         `pulumi:"assignmentState"`
	AutoApplyDecisionsEnabled       *bool                          `pulumi:"autoApplyDecisionsEnabled"`
	BackupReviewers                 []AccessReviewReviewerResponse `pulumi:"backupReviewers"`
	DefaultDecision                 *string                        `pulumi:"defaultDecision"`
	DefaultDecisionEnabled          *bool                          `pulumi:"defaultDecisionEnabled"`
	DescriptionForAdmins            *string                        `pulumi:"descriptionForAdmins"`
	DescriptionForReviewers         *string                        `pulumi:"descriptionForReviewers"`
	DisplayName                     *string                        `pulumi:"displayName"`
	EndDate                         *string                        `pulumi:"endDate"`
	ExcludeResourceId               *string                        `pulumi:"excludeResourceId"`
	ExcludeRoleDefinitionId         *string                        `pulumi:"excludeRoleDefinitionId"`
	ExpandNestedMemberships         *bool                          `pulumi:"expandNestedMemberships"`
	Id                              string                         `pulumi:"id"`
	InactiveDuration                *string                        `pulumi:"inactiveDuration"`
	IncludeAccessBelowResource      *bool                          `pulumi:"includeAccessBelowResource"`
	IncludeInheritedAccess          *bool                          `pulumi:"includeInheritedAccess"`
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

func LookupScopeAccessReviewScheduleDefinitionByIdOutput(ctx *pulumi.Context, args LookupScopeAccessReviewScheduleDefinitionByIdOutputArgs, opts ...pulumi.InvokeOption) LookupScopeAccessReviewScheduleDefinitionByIdResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScopeAccessReviewScheduleDefinitionByIdResult, error) {
			args := v.(LookupScopeAccessReviewScheduleDefinitionByIdArgs)
			r, err := LookupScopeAccessReviewScheduleDefinitionById(ctx, &args, opts...)
			var s LookupScopeAccessReviewScheduleDefinitionByIdResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScopeAccessReviewScheduleDefinitionByIdResultOutput)
}

type LookupScopeAccessReviewScheduleDefinitionByIdOutputArgs struct {
	ScheduleDefinitionId pulumi.StringInput `pulumi:"scheduleDefinitionId"`
	Scope                pulumi.StringInput `pulumi:"scope"`
}

func (LookupScopeAccessReviewScheduleDefinitionByIdOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScopeAccessReviewScheduleDefinitionByIdArgs)(nil)).Elem()
}


type LookupScopeAccessReviewScheduleDefinitionByIdResultOutput struct{ *pulumi.OutputState }

func (LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScopeAccessReviewScheduleDefinitionByIdResult)(nil)).Elem()
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) ToLookupScopeAccessReviewScheduleDefinitionByIdResultOutput() LookupScopeAccessReviewScheduleDefinitionByIdResultOutput {
	return o
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) ToLookupScopeAccessReviewScheduleDefinitionByIdResultOutputWithContext(ctx context.Context) LookupScopeAccessReviewScheduleDefinitionByIdResultOutput {
	return o
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) AssignmentState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) string { return v.AssignmentState }).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) AutoApplyDecisionsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *bool { return v.AutoApplyDecisionsEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) BackupReviewers() AccessReviewReviewerResponseArrayOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) []AccessReviewReviewerResponse {
		return v.BackupReviewers
	}).(AccessReviewReviewerResponseArrayOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) DefaultDecision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *string { return v.DefaultDecision }).(pulumi.StringPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) DefaultDecisionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *bool { return v.DefaultDecisionEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) DescriptionForAdmins() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *string { return v.DescriptionForAdmins }).(pulumi.StringPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) DescriptionForReviewers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *string { return v.DescriptionForReviewers }).(pulumi.StringPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) ExcludeResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *string { return v.ExcludeResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) ExcludeRoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *string { return v.ExcludeRoleDefinitionId }).(pulumi.StringPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) ExpandNestedMemberships() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *bool { return v.ExpandNestedMemberships }).(pulumi.BoolPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) InactiveDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *string { return v.InactiveDuration }).(pulumi.StringPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) IncludeAccessBelowResource() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *bool { return v.IncludeAccessBelowResource }).(pulumi.BoolPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) IncludeInheritedAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *bool { return v.IncludeInheritedAccess }).(pulumi.BoolPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) InstanceDurationInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *int { return v.InstanceDurationInDays }).(pulumi.IntPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) Instances() AccessReviewInstanceResponseArrayOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) []AccessReviewInstanceResponse {
		return v.Instances
	}).(AccessReviewInstanceResponseArrayOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) JustificationRequiredOnApproval() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *bool {
		return v.JustificationRequiredOnApproval
	}).(pulumi.BoolPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) MailNotificationsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *bool { return v.MailNotificationsEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) NumberOfOccurrences() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *int { return v.NumberOfOccurrences }).(pulumi.IntPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) string { return v.PrincipalName }).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) string { return v.PrincipalType }).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) RecommendationLookBackDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *string {
		return v.RecommendationLookBackDuration
	}).(pulumi.StringPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) RecommendationsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *bool { return v.RecommendationsEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) ReminderNotificationsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *bool {
		return v.ReminderNotificationsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) Reviewers() AccessReviewReviewerResponseArrayOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) []AccessReviewReviewerResponse {
		return v.Reviewers
	}).(AccessReviewReviewerResponseArrayOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) ReviewersType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) string { return v.ReviewersType }).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) *string { return v.StartDate }).(pulumi.StringPtrOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupScopeAccessReviewScheduleDefinitionByIdResultOutput) UserPrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAccessReviewScheduleDefinitionByIdResult) string { return v.UserPrincipalName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScopeAccessReviewScheduleDefinitionByIdResultOutput{})
}

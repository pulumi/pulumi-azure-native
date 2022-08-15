


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessReviewScheduleDefinitionById struct {
	pulumi.CustomResourceState

	AssignmentState                 pulumi.StringOutput                     `pulumi:"assignmentState"`
	AutoApplyDecisionsEnabled       pulumi.BoolPtrOutput                    `pulumi:"autoApplyDecisionsEnabled"`
	BackupReviewers                 AccessReviewReviewerResponseArrayOutput `pulumi:"backupReviewers"`
	DefaultDecision                 pulumi.StringPtrOutput                  `pulumi:"defaultDecision"`
	DefaultDecisionEnabled          pulumi.BoolPtrOutput                    `pulumi:"defaultDecisionEnabled"`
	DescriptionForAdmins            pulumi.StringPtrOutput                  `pulumi:"descriptionForAdmins"`
	DescriptionForReviewers         pulumi.StringPtrOutput                  `pulumi:"descriptionForReviewers"`
	DisplayName                     pulumi.StringPtrOutput                  `pulumi:"displayName"`
	EndDate                         pulumi.StringPtrOutput                  `pulumi:"endDate"`
	ExpandNestedMemberships         pulumi.BoolPtrOutput                    `pulumi:"expandNestedMemberships"`
	InactiveDuration                pulumi.StringPtrOutput                  `pulumi:"inactiveDuration"`
	InstanceDurationInDays          pulumi.IntPtrOutput                     `pulumi:"instanceDurationInDays"`
	Instances                       AccessReviewInstanceResponseArrayOutput `pulumi:"instances"`
	Interval                        pulumi.IntPtrOutput                     `pulumi:"interval"`
	JustificationRequiredOnApproval pulumi.BoolPtrOutput                    `pulumi:"justificationRequiredOnApproval"`
	MailNotificationsEnabled        pulumi.BoolPtrOutput                    `pulumi:"mailNotificationsEnabled"`
	Name                            pulumi.StringOutput                     `pulumi:"name"`
	NumberOfOccurrences             pulumi.IntPtrOutput                     `pulumi:"numberOfOccurrences"`
	PrincipalId                     pulumi.StringOutput                     `pulumi:"principalId"`
	PrincipalName                   pulumi.StringOutput                     `pulumi:"principalName"`
	PrincipalType                   pulumi.StringOutput                     `pulumi:"principalType"`
	RecommendationLookBackDuration  pulumi.StringPtrOutput                  `pulumi:"recommendationLookBackDuration"`
	RecommendationsEnabled          pulumi.BoolPtrOutput                    `pulumi:"recommendationsEnabled"`
	ReminderNotificationsEnabled    pulumi.BoolPtrOutput                    `pulumi:"reminderNotificationsEnabled"`
	ResourceId                      pulumi.StringOutput                     `pulumi:"resourceId"`
	Reviewers                       AccessReviewReviewerResponseArrayOutput `pulumi:"reviewers"`
	ReviewersType                   pulumi.StringOutput                     `pulumi:"reviewersType"`
	RoleDefinitionId                pulumi.StringOutput                     `pulumi:"roleDefinitionId"`
	StartDate                       pulumi.StringPtrOutput                  `pulumi:"startDate"`
	Status                          pulumi.StringOutput                     `pulumi:"status"`
	Type                            pulumi.StringOutput                     `pulumi:"type"`
	UserPrincipalName               pulumi.StringOutput                     `pulumi:"userPrincipalName"`
}


func NewAccessReviewScheduleDefinitionById(ctx *pulumi.Context,
	name string, args *AccessReviewScheduleDefinitionByIdArgs, opts ...pulumi.ResourceOption) (*AccessReviewScheduleDefinitionById, error) {
	if args == nil {
		args = &AccessReviewScheduleDefinitionByIdArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization:AccessReviewScheduleDefinitionById"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180501preview:AccessReviewScheduleDefinitionById"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20210301preview:AccessReviewScheduleDefinitionById"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20211116preview:AccessReviewScheduleDefinitionById"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20211201preview:AccessReviewScheduleDefinitionById"),
		},
	})
	opts = append(opts, aliases)
	var resource AccessReviewScheduleDefinitionById
	err := ctx.RegisterResource("azure-native:authorization/v20210701preview:AccessReviewScheduleDefinitionById", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAccessReviewScheduleDefinitionById(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessReviewScheduleDefinitionByIdState, opts ...pulumi.ResourceOption) (*AccessReviewScheduleDefinitionById, error) {
	var resource AccessReviewScheduleDefinitionById
	err := ctx.ReadResource("azure-native:authorization/v20210701preview:AccessReviewScheduleDefinitionById", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type accessReviewScheduleDefinitionByIdState struct {
}

type AccessReviewScheduleDefinitionByIdState struct {
}

func (AccessReviewScheduleDefinitionByIdState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessReviewScheduleDefinitionByIdState)(nil)).Elem()
}

type accessReviewScheduleDefinitionByIdArgs struct {
	AutoApplyDecisionsEnabled       *bool                  `pulumi:"autoApplyDecisionsEnabled"`
	BackupReviewers                 []AccessReviewReviewer `pulumi:"backupReviewers"`
	DefaultDecision                 *string                `pulumi:"defaultDecision"`
	DefaultDecisionEnabled          *bool                  `pulumi:"defaultDecisionEnabled"`
	DescriptionForAdmins            *string                `pulumi:"descriptionForAdmins"`
	DescriptionForReviewers         *string                `pulumi:"descriptionForReviewers"`
	DisplayName                     *string                `pulumi:"displayName"`
	EndDate                         *string                `pulumi:"endDate"`
	ExpandNestedMemberships         *bool                  `pulumi:"expandNestedMemberships"`
	InactiveDuration                *string                `pulumi:"inactiveDuration"`
	InstanceDurationInDays          *int                   `pulumi:"instanceDurationInDays"`
	Instances                       []AccessReviewInstance `pulumi:"instances"`
	Interval                        *int                   `pulumi:"interval"`
	JustificationRequiredOnApproval *bool                  `pulumi:"justificationRequiredOnApproval"`
	MailNotificationsEnabled        *bool                  `pulumi:"mailNotificationsEnabled"`
	NumberOfOccurrences             *int                   `pulumi:"numberOfOccurrences"`
	RecommendationLookBackDuration  *string                `pulumi:"recommendationLookBackDuration"`
	RecommendationsEnabled          *bool                  `pulumi:"recommendationsEnabled"`
	ReminderNotificationsEnabled    *bool                  `pulumi:"reminderNotificationsEnabled"`
	Reviewers                       []AccessReviewReviewer `pulumi:"reviewers"`
	ScheduleDefinitionId            *string                `pulumi:"scheduleDefinitionId"`
	StartDate                       *string                `pulumi:"startDate"`
	Type                            *string                `pulumi:"type"`
}


type AccessReviewScheduleDefinitionByIdArgs struct {
	AutoApplyDecisionsEnabled       pulumi.BoolPtrInput
	BackupReviewers                 AccessReviewReviewerArrayInput
	DefaultDecision                 pulumi.StringPtrInput
	DefaultDecisionEnabled          pulumi.BoolPtrInput
	DescriptionForAdmins            pulumi.StringPtrInput
	DescriptionForReviewers         pulumi.StringPtrInput
	DisplayName                     pulumi.StringPtrInput
	EndDate                         pulumi.StringPtrInput
	ExpandNestedMemberships         pulumi.BoolPtrInput
	InactiveDuration                pulumi.StringPtrInput
	InstanceDurationInDays          pulumi.IntPtrInput
	Instances                       AccessReviewInstanceArrayInput
	Interval                        pulumi.IntPtrInput
	JustificationRequiredOnApproval pulumi.BoolPtrInput
	MailNotificationsEnabled        pulumi.BoolPtrInput
	NumberOfOccurrences             pulumi.IntPtrInput
	RecommendationLookBackDuration  pulumi.StringPtrInput
	RecommendationsEnabled          pulumi.BoolPtrInput
	ReminderNotificationsEnabled    pulumi.BoolPtrInput
	Reviewers                       AccessReviewReviewerArrayInput
	ScheduleDefinitionId            pulumi.StringPtrInput
	StartDate                       pulumi.StringPtrInput
	Type                            pulumi.StringPtrInput
}

func (AccessReviewScheduleDefinitionByIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessReviewScheduleDefinitionByIdArgs)(nil)).Elem()
}

type AccessReviewScheduleDefinitionByIdInput interface {
	pulumi.Input

	ToAccessReviewScheduleDefinitionByIdOutput() AccessReviewScheduleDefinitionByIdOutput
	ToAccessReviewScheduleDefinitionByIdOutputWithContext(ctx context.Context) AccessReviewScheduleDefinitionByIdOutput
}

func (*AccessReviewScheduleDefinitionById) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessReviewScheduleDefinitionById)(nil)).Elem()
}

func (i *AccessReviewScheduleDefinitionById) ToAccessReviewScheduleDefinitionByIdOutput() AccessReviewScheduleDefinitionByIdOutput {
	return i.ToAccessReviewScheduleDefinitionByIdOutputWithContext(context.Background())
}

func (i *AccessReviewScheduleDefinitionById) ToAccessReviewScheduleDefinitionByIdOutputWithContext(ctx context.Context) AccessReviewScheduleDefinitionByIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessReviewScheduleDefinitionByIdOutput)
}

type AccessReviewScheduleDefinitionByIdOutput struct{ *pulumi.OutputState }

func (AccessReviewScheduleDefinitionByIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessReviewScheduleDefinitionById)(nil)).Elem()
}

func (o AccessReviewScheduleDefinitionByIdOutput) ToAccessReviewScheduleDefinitionByIdOutput() AccessReviewScheduleDefinitionByIdOutput {
	return o
}

func (o AccessReviewScheduleDefinitionByIdOutput) ToAccessReviewScheduleDefinitionByIdOutputWithContext(ctx context.Context) AccessReviewScheduleDefinitionByIdOutput {
	return o
}

func (o AccessReviewScheduleDefinitionByIdOutput) AssignmentState() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.AssignmentState }).(pulumi.StringOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) AutoApplyDecisionsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput { return v.AutoApplyDecisionsEnabled }).(pulumi.BoolPtrOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) BackupReviewers() AccessReviewReviewerResponseArrayOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) AccessReviewReviewerResponseArrayOutput {
		return v.BackupReviewers
	}).(AccessReviewReviewerResponseArrayOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) DefaultDecision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.DefaultDecision }).(pulumi.StringPtrOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) DefaultDecisionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput { return v.DefaultDecisionEnabled }).(pulumi.BoolPtrOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) DescriptionForAdmins() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.DescriptionForAdmins }).(pulumi.StringPtrOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) DescriptionForReviewers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.DescriptionForReviewers }).(pulumi.StringPtrOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) ExpandNestedMemberships() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput { return v.ExpandNestedMemberships }).(pulumi.BoolPtrOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) InactiveDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.InactiveDuration }).(pulumi.StringPtrOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) InstanceDurationInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.IntPtrOutput { return v.InstanceDurationInDays }).(pulumi.IntPtrOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) Instances() AccessReviewInstanceResponseArrayOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) AccessReviewInstanceResponseArrayOutput {
		return v.Instances
	}).(AccessReviewInstanceResponseArrayOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.IntPtrOutput { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) JustificationRequiredOnApproval() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput {
		return v.JustificationRequiredOnApproval
	}).(pulumi.BoolPtrOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) MailNotificationsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput { return v.MailNotificationsEnabled }).(pulumi.BoolPtrOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) NumberOfOccurrences() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.IntPtrOutput { return v.NumberOfOccurrences }).(pulumi.IntPtrOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.PrincipalName }).(pulumi.StringOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.PrincipalType }).(pulumi.StringOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) RecommendationLookBackDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringPtrOutput {
		return v.RecommendationLookBackDuration
	}).(pulumi.StringPtrOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) RecommendationsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput { return v.RecommendationsEnabled }).(pulumi.BoolPtrOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) ReminderNotificationsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput {
		return v.ReminderNotificationsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) Reviewers() AccessReviewReviewerResponseArrayOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) AccessReviewReviewerResponseArrayOutput {
		return v.Reviewers
	}).(AccessReviewReviewerResponseArrayOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) ReviewersType() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.ReviewersType }).(pulumi.StringOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.StartDate }).(pulumi.StringPtrOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AccessReviewScheduleDefinitionByIdOutput) UserPrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.UserPrincipalName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessReviewScheduleDefinitionByIdOutput{})
}

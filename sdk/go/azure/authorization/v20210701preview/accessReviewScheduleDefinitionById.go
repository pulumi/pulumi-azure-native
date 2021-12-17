


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

func init() {
	pulumi.RegisterOutputType(AccessReviewScheduleDefinitionByIdOutput{})
}

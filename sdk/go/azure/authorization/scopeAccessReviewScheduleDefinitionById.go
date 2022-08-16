


package authorization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ScopeAccessReviewScheduleDefinitionById struct {
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
	ExcludeResourceId               pulumi.StringPtrOutput                  `pulumi:"excludeResourceId"`
	ExcludeRoleDefinitionId         pulumi.StringPtrOutput                  `pulumi:"excludeRoleDefinitionId"`
	ExpandNestedMemberships         pulumi.BoolPtrOutput                    `pulumi:"expandNestedMemberships"`
	InactiveDuration                pulumi.StringPtrOutput                  `pulumi:"inactiveDuration"`
	IncludeAccessBelowResource      pulumi.BoolPtrOutput                    `pulumi:"includeAccessBelowResource"`
	IncludeInheritedAccess          pulumi.BoolPtrOutput                    `pulumi:"includeInheritedAccess"`
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


func NewScopeAccessReviewScheduleDefinitionById(ctx *pulumi.Context,
	name string, args *ScopeAccessReviewScheduleDefinitionByIdArgs, opts ...pulumi.ResourceOption) (*ScopeAccessReviewScheduleDefinitionById, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization/v20211201preview:ScopeAccessReviewScheduleDefinitionById"),
		},
	})
	opts = append(opts, aliases)
	var resource ScopeAccessReviewScheduleDefinitionById
	err := ctx.RegisterResource("azure-native:authorization:ScopeAccessReviewScheduleDefinitionById", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetScopeAccessReviewScheduleDefinitionById(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScopeAccessReviewScheduleDefinitionByIdState, opts ...pulumi.ResourceOption) (*ScopeAccessReviewScheduleDefinitionById, error) {
	var resource ScopeAccessReviewScheduleDefinitionById
	err := ctx.ReadResource("azure-native:authorization:ScopeAccessReviewScheduleDefinitionById", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type scopeAccessReviewScheduleDefinitionByIdState struct {
}

type ScopeAccessReviewScheduleDefinitionByIdState struct {
}

func (ScopeAccessReviewScheduleDefinitionByIdState) ElementType() reflect.Type {
	return reflect.TypeOf((*scopeAccessReviewScheduleDefinitionByIdState)(nil)).Elem()
}

type scopeAccessReviewScheduleDefinitionByIdArgs struct {
	AutoApplyDecisionsEnabled       *bool                  `pulumi:"autoApplyDecisionsEnabled"`
	BackupReviewers                 []AccessReviewReviewer `pulumi:"backupReviewers"`
	DefaultDecision                 *string                `pulumi:"defaultDecision"`
	DefaultDecisionEnabled          *bool                  `pulumi:"defaultDecisionEnabled"`
	DescriptionForAdmins            *string                `pulumi:"descriptionForAdmins"`
	DescriptionForReviewers         *string                `pulumi:"descriptionForReviewers"`
	DisplayName                     *string                `pulumi:"displayName"`
	EndDate                         *string                `pulumi:"endDate"`
	ExcludeResourceId               *string                `pulumi:"excludeResourceId"`
	ExcludeRoleDefinitionId         *string                `pulumi:"excludeRoleDefinitionId"`
	ExpandNestedMemberships         *bool                  `pulumi:"expandNestedMemberships"`
	InactiveDuration                *string                `pulumi:"inactiveDuration"`
	IncludeAccessBelowResource      *bool                  `pulumi:"includeAccessBelowResource"`
	IncludeInheritedAccess          *bool                  `pulumi:"includeInheritedAccess"`
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
	Scope                           string                 `pulumi:"scope"`
	StartDate                       *string                `pulumi:"startDate"`
	Type                            *string                `pulumi:"type"`
}


type ScopeAccessReviewScheduleDefinitionByIdArgs struct {
	AutoApplyDecisionsEnabled       pulumi.BoolPtrInput
	BackupReviewers                 AccessReviewReviewerArrayInput
	DefaultDecision                 pulumi.StringPtrInput
	DefaultDecisionEnabled          pulumi.BoolPtrInput
	DescriptionForAdmins            pulumi.StringPtrInput
	DescriptionForReviewers         pulumi.StringPtrInput
	DisplayName                     pulumi.StringPtrInput
	EndDate                         pulumi.StringPtrInput
	ExcludeResourceId               pulumi.StringPtrInput
	ExcludeRoleDefinitionId         pulumi.StringPtrInput
	ExpandNestedMemberships         pulumi.BoolPtrInput
	InactiveDuration                pulumi.StringPtrInput
	IncludeAccessBelowResource      pulumi.BoolPtrInput
	IncludeInheritedAccess          pulumi.BoolPtrInput
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
	Scope                           pulumi.StringInput
	StartDate                       pulumi.StringPtrInput
	Type                            pulumi.StringPtrInput
}

func (ScopeAccessReviewScheduleDefinitionByIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scopeAccessReviewScheduleDefinitionByIdArgs)(nil)).Elem()
}

type ScopeAccessReviewScheduleDefinitionByIdInput interface {
	pulumi.Input

	ToScopeAccessReviewScheduleDefinitionByIdOutput() ScopeAccessReviewScheduleDefinitionByIdOutput
	ToScopeAccessReviewScheduleDefinitionByIdOutputWithContext(ctx context.Context) ScopeAccessReviewScheduleDefinitionByIdOutput
}

func (*ScopeAccessReviewScheduleDefinitionById) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeAccessReviewScheduleDefinitionById)(nil)).Elem()
}

func (i *ScopeAccessReviewScheduleDefinitionById) ToScopeAccessReviewScheduleDefinitionByIdOutput() ScopeAccessReviewScheduleDefinitionByIdOutput {
	return i.ToScopeAccessReviewScheduleDefinitionByIdOutputWithContext(context.Background())
}

func (i *ScopeAccessReviewScheduleDefinitionById) ToScopeAccessReviewScheduleDefinitionByIdOutputWithContext(ctx context.Context) ScopeAccessReviewScheduleDefinitionByIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeAccessReviewScheduleDefinitionByIdOutput)
}

type ScopeAccessReviewScheduleDefinitionByIdOutput struct{ *pulumi.OutputState }

func (ScopeAccessReviewScheduleDefinitionByIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeAccessReviewScheduleDefinitionById)(nil)).Elem()
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) ToScopeAccessReviewScheduleDefinitionByIdOutput() ScopeAccessReviewScheduleDefinitionByIdOutput {
	return o
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) ToScopeAccessReviewScheduleDefinitionByIdOutputWithContext(ctx context.Context) ScopeAccessReviewScheduleDefinitionByIdOutput {
	return o
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) AssignmentState() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.AssignmentState }).(pulumi.StringOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) AutoApplyDecisionsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput {
		return v.AutoApplyDecisionsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) BackupReviewers() AccessReviewReviewerResponseArrayOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) AccessReviewReviewerResponseArrayOutput {
		return v.BackupReviewers
	}).(AccessReviewReviewerResponseArrayOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) DefaultDecision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.DefaultDecision }).(pulumi.StringPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) DefaultDecisionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput { return v.DefaultDecisionEnabled }).(pulumi.BoolPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) DescriptionForAdmins() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.DescriptionForAdmins }).(pulumi.StringPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) DescriptionForReviewers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringPtrOutput {
		return v.DescriptionForReviewers
	}).(pulumi.StringPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) ExcludeResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.ExcludeResourceId }).(pulumi.StringPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) ExcludeRoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringPtrOutput {
		return v.ExcludeRoleDefinitionId
	}).(pulumi.StringPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) ExpandNestedMemberships() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput {
		return v.ExpandNestedMemberships
	}).(pulumi.BoolPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) InactiveDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.InactiveDuration }).(pulumi.StringPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) IncludeAccessBelowResource() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput {
		return v.IncludeAccessBelowResource
	}).(pulumi.BoolPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) IncludeInheritedAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput { return v.IncludeInheritedAccess }).(pulumi.BoolPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) InstanceDurationInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.IntPtrOutput { return v.InstanceDurationInDays }).(pulumi.IntPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) Instances() AccessReviewInstanceResponseArrayOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) AccessReviewInstanceResponseArrayOutput {
		return v.Instances
	}).(AccessReviewInstanceResponseArrayOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.IntPtrOutput { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) JustificationRequiredOnApproval() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput {
		return v.JustificationRequiredOnApproval
	}).(pulumi.BoolPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) MailNotificationsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput {
		return v.MailNotificationsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) NumberOfOccurrences() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.IntPtrOutput { return v.NumberOfOccurrences }).(pulumi.IntPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.PrincipalName }).(pulumi.StringOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.PrincipalType }).(pulumi.StringOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) RecommendationLookBackDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringPtrOutput {
		return v.RecommendationLookBackDuration
	}).(pulumi.StringPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) RecommendationsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput { return v.RecommendationsEnabled }).(pulumi.BoolPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) ReminderNotificationsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput {
		return v.ReminderNotificationsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) Reviewers() AccessReviewReviewerResponseArrayOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) AccessReviewReviewerResponseArrayOutput {
		return v.Reviewers
	}).(AccessReviewReviewerResponseArrayOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) ReviewersType() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.ReviewersType }).(pulumi.StringOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.StartDate }).(pulumi.StringPtrOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ScopeAccessReviewScheduleDefinitionByIdOutput) UserPrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.UserPrincipalName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScopeAccessReviewScheduleDefinitionByIdOutput{})
}

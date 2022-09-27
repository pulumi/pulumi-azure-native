


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ScheduledQueryRule struct {
	pulumi.CustomResourceState

	Actions                               ActionsResponsePtrOutput                 `pulumi:"actions"`
	AutoMitigate                          pulumi.BoolPtrOutput                     `pulumi:"autoMitigate"`
	CheckWorkspaceAlertsStorageConfigured pulumi.BoolPtrOutput                     `pulumi:"checkWorkspaceAlertsStorageConfigured"`
	CreatedWithApiVersion                 pulumi.StringOutput                      `pulumi:"createdWithApiVersion"`
	Criteria                              ScheduledQueryRuleCriteriaResponseOutput `pulumi:"criteria"`
	Description                           pulumi.StringPtrOutput                   `pulumi:"description"`
	DisplayName                           pulumi.StringPtrOutput                   `pulumi:"displayName"`
	Enabled                               pulumi.BoolOutput                        `pulumi:"enabled"`
	Etag                                  pulumi.StringOutput                      `pulumi:"etag"`
	EvaluationFrequency                   pulumi.StringPtrOutput                   `pulumi:"evaluationFrequency"`
	IsLegacyLogAnalyticsRule              pulumi.BoolOutput                        `pulumi:"isLegacyLogAnalyticsRule"`
	IsWorkspaceAlertsStorageConfigured    pulumi.BoolOutput                        `pulumi:"isWorkspaceAlertsStorageConfigured"`
	Kind                                  pulumi.StringPtrOutput                   `pulumi:"kind"`
	Location                              pulumi.StringOutput                      `pulumi:"location"`
	MuteActionsDuration                   pulumi.StringPtrOutput                   `pulumi:"muteActionsDuration"`
	Name                                  pulumi.StringOutput                      `pulumi:"name"`
	OverrideQueryTimeRange                pulumi.StringPtrOutput                   `pulumi:"overrideQueryTimeRange"`
	Scopes                                pulumi.StringArrayOutput                 `pulumi:"scopes"`
	Severity                              pulumi.Float64PtrOutput                  `pulumi:"severity"`
	SkipQueryValidation                   pulumi.BoolPtrOutput                     `pulumi:"skipQueryValidation"`
	SystemData                            SystemDataResponseOutput                 `pulumi:"systemData"`
	Tags                                  pulumi.StringMapOutput                   `pulumi:"tags"`
	TargetResourceTypes                   pulumi.StringArrayOutput                 `pulumi:"targetResourceTypes"`
	Type                                  pulumi.StringOutput                      `pulumi:"type"`
	WindowSize                            pulumi.StringPtrOutput                   `pulumi:"windowSize"`
}


func NewScheduledQueryRule(ctx *pulumi.Context,
	name string, args *ScheduledQueryRuleArgs, opts ...pulumi.ResourceOption) (*ScheduledQueryRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Criteria == nil {
		return nil, errors.New("invalid value for required argument 'Criteria'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Scopes == nil {
		return nil, errors.New("invalid value for required argument 'Scopes'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:ScheduledQueryRule"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20180416:ScheduledQueryRule"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20200501preview:ScheduledQueryRule"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210801:ScheduledQueryRule"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20220615:ScheduledQueryRule"),
		},
	})
	opts = append(opts, aliases)
	var resource ScheduledQueryRule
	err := ctx.RegisterResource("azure-native:insights/v20210201preview:ScheduledQueryRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetScheduledQueryRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledQueryRuleState, opts ...pulumi.ResourceOption) (*ScheduledQueryRule, error) {
	var resource ScheduledQueryRule
	err := ctx.ReadResource("azure-native:insights/v20210201preview:ScheduledQueryRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type scheduledQueryRuleState struct {
}

type ScheduledQueryRuleState struct {
}

func (ScheduledQueryRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledQueryRuleState)(nil)).Elem()
}

type scheduledQueryRuleArgs struct {
	Actions                               *Actions                   `pulumi:"actions"`
	AutoMitigate                          *bool                      `pulumi:"autoMitigate"`
	CheckWorkspaceAlertsStorageConfigured *bool                      `pulumi:"checkWorkspaceAlertsStorageConfigured"`
	Criteria                              ScheduledQueryRuleCriteria `pulumi:"criteria"`
	Description                           *string                    `pulumi:"description"`
	DisplayName                           *string                    `pulumi:"displayName"`
	Enabled                               bool                       `pulumi:"enabled"`
	EvaluationFrequency                   *string                    `pulumi:"evaluationFrequency"`
	Kind                                  *string                    `pulumi:"kind"`
	Location                              *string                    `pulumi:"location"`
	MuteActionsDuration                   *string                    `pulumi:"muteActionsDuration"`
	OverrideQueryTimeRange                *string                    `pulumi:"overrideQueryTimeRange"`
	ResourceGroupName                     string                     `pulumi:"resourceGroupName"`
	RuleName                              *string                    `pulumi:"ruleName"`
	Scopes                                []string                   `pulumi:"scopes"`
	Severity                              *float64                   `pulumi:"severity"`
	SkipQueryValidation                   *bool                      `pulumi:"skipQueryValidation"`
	Tags                                  map[string]string          `pulumi:"tags"`
	TargetResourceTypes                   []string                   `pulumi:"targetResourceTypes"`
	WindowSize                            *string                    `pulumi:"windowSize"`
}


type ScheduledQueryRuleArgs struct {
	Actions                               ActionsPtrInput
	AutoMitigate                          pulumi.BoolPtrInput
	CheckWorkspaceAlertsStorageConfigured pulumi.BoolPtrInput
	Criteria                              ScheduledQueryRuleCriteriaInput
	Description                           pulumi.StringPtrInput
	DisplayName                           pulumi.StringPtrInput
	Enabled                               pulumi.BoolInput
	EvaluationFrequency                   pulumi.StringPtrInput
	Kind                                  pulumi.StringPtrInput
	Location                              pulumi.StringPtrInput
	MuteActionsDuration                   pulumi.StringPtrInput
	OverrideQueryTimeRange                pulumi.StringPtrInput
	ResourceGroupName                     pulumi.StringInput
	RuleName                              pulumi.StringPtrInput
	Scopes                                pulumi.StringArrayInput
	Severity                              pulumi.Float64PtrInput
	SkipQueryValidation                   pulumi.BoolPtrInput
	Tags                                  pulumi.StringMapInput
	TargetResourceTypes                   pulumi.StringArrayInput
	WindowSize                            pulumi.StringPtrInput
}

func (ScheduledQueryRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledQueryRuleArgs)(nil)).Elem()
}

type ScheduledQueryRuleInput interface {
	pulumi.Input

	ToScheduledQueryRuleOutput() ScheduledQueryRuleOutput
	ToScheduledQueryRuleOutputWithContext(ctx context.Context) ScheduledQueryRuleOutput
}

func (*ScheduledQueryRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledQueryRule)(nil)).Elem()
}

func (i *ScheduledQueryRule) ToScheduledQueryRuleOutput() ScheduledQueryRuleOutput {
	return i.ToScheduledQueryRuleOutputWithContext(context.Background())
}

func (i *ScheduledQueryRule) ToScheduledQueryRuleOutputWithContext(ctx context.Context) ScheduledQueryRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledQueryRuleOutput)
}

type ScheduledQueryRuleOutput struct{ *pulumi.OutputState }

func (ScheduledQueryRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledQueryRule)(nil)).Elem()
}

func (o ScheduledQueryRuleOutput) ToScheduledQueryRuleOutput() ScheduledQueryRuleOutput {
	return o
}

func (o ScheduledQueryRuleOutput) ToScheduledQueryRuleOutputWithContext(ctx context.Context) ScheduledQueryRuleOutput {
	return o
}

func (o ScheduledQueryRuleOutput) Actions() ActionsResponsePtrOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) ActionsResponsePtrOutput { return v.Actions }).(ActionsResponsePtrOutput)
}

func (o ScheduledQueryRuleOutput) AutoMitigate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.BoolPtrOutput { return v.AutoMitigate }).(pulumi.BoolPtrOutput)
}

func (o ScheduledQueryRuleOutput) CheckWorkspaceAlertsStorageConfigured() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.BoolPtrOutput { return v.CheckWorkspaceAlertsStorageConfigured }).(pulumi.BoolPtrOutput)
}

func (o ScheduledQueryRuleOutput) CreatedWithApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.StringOutput { return v.CreatedWithApiVersion }).(pulumi.StringOutput)
}

func (o ScheduledQueryRuleOutput) Criteria() ScheduledQueryRuleCriteriaResponseOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) ScheduledQueryRuleCriteriaResponseOutput { return v.Criteria }).(ScheduledQueryRuleCriteriaResponseOutput)
}

func (o ScheduledQueryRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ScheduledQueryRuleOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ScheduledQueryRuleOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o ScheduledQueryRuleOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ScheduledQueryRuleOutput) EvaluationFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.StringPtrOutput { return v.EvaluationFrequency }).(pulumi.StringPtrOutput)
}

func (o ScheduledQueryRuleOutput) IsLegacyLogAnalyticsRule() pulumi.BoolOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.BoolOutput { return v.IsLegacyLogAnalyticsRule }).(pulumi.BoolOutput)
}

func (o ScheduledQueryRuleOutput) IsWorkspaceAlertsStorageConfigured() pulumi.BoolOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.BoolOutput { return v.IsWorkspaceAlertsStorageConfigured }).(pulumi.BoolOutput)
}

func (o ScheduledQueryRuleOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ScheduledQueryRuleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ScheduledQueryRuleOutput) MuteActionsDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.StringPtrOutput { return v.MuteActionsDuration }).(pulumi.StringPtrOutput)
}

func (o ScheduledQueryRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ScheduledQueryRuleOutput) OverrideQueryTimeRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.StringPtrOutput { return v.OverrideQueryTimeRange }).(pulumi.StringPtrOutput)
}

func (o ScheduledQueryRuleOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

func (o ScheduledQueryRuleOutput) Severity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.Float64PtrOutput { return v.Severity }).(pulumi.Float64PtrOutput)
}

func (o ScheduledQueryRuleOutput) SkipQueryValidation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.BoolPtrOutput { return v.SkipQueryValidation }).(pulumi.BoolPtrOutput)
}

func (o ScheduledQueryRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ScheduledQueryRuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ScheduledQueryRuleOutput) TargetResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.StringArrayOutput { return v.TargetResourceTypes }).(pulumi.StringArrayOutput)
}

func (o ScheduledQueryRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ScheduledQueryRuleOutput) WindowSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledQueryRule) pulumi.StringPtrOutput { return v.WindowSize }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduledQueryRuleOutput{})
}

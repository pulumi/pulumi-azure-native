


package v20200501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ScheduledQueryRule struct {
	pulumi.CustomResourceState

	Actions                  ActionResponseArrayOutput                `pulumi:"actions"`
	CreatedWithApiVersion    pulumi.StringOutput                      `pulumi:"createdWithApiVersion"`
	Criteria                 ScheduledQueryRuleCriteriaResponseOutput `pulumi:"criteria"`
	Description              pulumi.StringPtrOutput                   `pulumi:"description"`
	DisplayName              pulumi.StringPtrOutput                   `pulumi:"displayName"`
	Enabled                  pulumi.BoolOutput                        `pulumi:"enabled"`
	Etag                     pulumi.StringOutput                      `pulumi:"etag"`
	EvaluationFrequency      pulumi.StringOutput                      `pulumi:"evaluationFrequency"`
	IsLegacyLogAnalyticsRule pulumi.BoolOutput                        `pulumi:"isLegacyLogAnalyticsRule"`
	Kind                     pulumi.StringOutput                      `pulumi:"kind"`
	Location                 pulumi.StringOutput                      `pulumi:"location"`
	MuteActionsDuration      pulumi.StringPtrOutput                   `pulumi:"muteActionsDuration"`
	Name                     pulumi.StringOutput                      `pulumi:"name"`
	OverrideQueryTimeRange   pulumi.StringPtrOutput                   `pulumi:"overrideQueryTimeRange"`
	Scopes                   pulumi.StringArrayOutput                 `pulumi:"scopes"`
	Severity                 pulumi.Float64Output                     `pulumi:"severity"`
	Tags                     pulumi.StringMapOutput                   `pulumi:"tags"`
	TargetResourceTypes      pulumi.StringArrayOutput                 `pulumi:"targetResourceTypes"`
	Type                     pulumi.StringOutput                      `pulumi:"type"`
	WindowSize               pulumi.StringOutput                      `pulumi:"windowSize"`
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
	if args.EvaluationFrequency == nil {
		return nil, errors.New("invalid value for required argument 'EvaluationFrequency'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Scopes == nil {
		return nil, errors.New("invalid value for required argument 'Scopes'")
	}
	if args.Severity == nil {
		return nil, errors.New("invalid value for required argument 'Severity'")
	}
	if args.WindowSize == nil {
		return nil, errors.New("invalid value for required argument 'WindowSize'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:insights/v20200501preview:ScheduledQueryRule"),
		},
		{
			Type: pulumi.String("azure-native:insights:ScheduledQueryRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights:ScheduledQueryRule"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20180416:ScheduledQueryRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights/v20180416:ScheduledQueryRule"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210201preview:ScheduledQueryRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights/v20210201preview:ScheduledQueryRule"),
		},
	})
	opts = append(opts, aliases)
	var resource ScheduledQueryRule
	err := ctx.RegisterResource("azure-native:insights/v20200501preview:ScheduledQueryRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetScheduledQueryRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledQueryRuleState, opts ...pulumi.ResourceOption) (*ScheduledQueryRule, error) {
	var resource ScheduledQueryRule
	err := ctx.ReadResource("azure-native:insights/v20200501preview:ScheduledQueryRule", name, id, state, &resource, opts...)
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
	Actions                []Action                   `pulumi:"actions"`
	Criteria               ScheduledQueryRuleCriteria `pulumi:"criteria"`
	Description            *string                    `pulumi:"description"`
	DisplayName            *string                    `pulumi:"displayName"`
	Enabled                bool                       `pulumi:"enabled"`
	EvaluationFrequency    string                     `pulumi:"evaluationFrequency"`
	Location               *string                    `pulumi:"location"`
	MuteActionsDuration    *string                    `pulumi:"muteActionsDuration"`
	OverrideQueryTimeRange *string                    `pulumi:"overrideQueryTimeRange"`
	ResourceGroupName      string                     `pulumi:"resourceGroupName"`
	RuleName               *string                    `pulumi:"ruleName"`
	Scopes                 []string                   `pulumi:"scopes"`
	Severity               float64                    `pulumi:"severity"`
	Tags                   map[string]string          `pulumi:"tags"`
	TargetResourceTypes    []string                   `pulumi:"targetResourceTypes"`
	WindowSize             string                     `pulumi:"windowSize"`
}


type ScheduledQueryRuleArgs struct {
	Actions                ActionArrayInput
	Criteria               ScheduledQueryRuleCriteriaInput
	Description            pulumi.StringPtrInput
	DisplayName            pulumi.StringPtrInput
	Enabled                pulumi.BoolInput
	EvaluationFrequency    pulumi.StringInput
	Location               pulumi.StringPtrInput
	MuteActionsDuration    pulumi.StringPtrInput
	OverrideQueryTimeRange pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	RuleName               pulumi.StringPtrInput
	Scopes                 pulumi.StringArrayInput
	Severity               pulumi.Float64Input
	Tags                   pulumi.StringMapInput
	TargetResourceTypes    pulumi.StringArrayInput
	WindowSize             pulumi.StringInput
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
	return reflect.TypeOf((*ScheduledQueryRule)(nil))
}

func (i *ScheduledQueryRule) ToScheduledQueryRuleOutput() ScheduledQueryRuleOutput {
	return i.ToScheduledQueryRuleOutputWithContext(context.Background())
}

func (i *ScheduledQueryRule) ToScheduledQueryRuleOutputWithContext(ctx context.Context) ScheduledQueryRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledQueryRuleOutput)
}

type ScheduledQueryRuleOutput struct{ *pulumi.OutputState }

func (ScheduledQueryRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryRule)(nil))
}

func (o ScheduledQueryRuleOutput) ToScheduledQueryRuleOutput() ScheduledQueryRuleOutput {
	return o
}

func (o ScheduledQueryRuleOutput) ToScheduledQueryRuleOutputWithContext(ctx context.Context) ScheduledQueryRuleOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledQueryRuleInput)(nil)).Elem(), &ScheduledQueryRule{})
	pulumi.RegisterOutputType(ScheduledQueryRuleOutput{})
}

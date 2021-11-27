


package v20180416

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ScheduledQueryRule struct {
	pulumi.CustomResourceState

	Action                   pulumi.AnyOutput          `pulumi:"action"`
	AutoMitigate             pulumi.BoolPtrOutput      `pulumi:"autoMitigate"`
	CreatedWithApiVersion    pulumi.StringOutput       `pulumi:"createdWithApiVersion"`
	Description              pulumi.StringPtrOutput    `pulumi:"description"`
	DisplayName              pulumi.StringPtrOutput    `pulumi:"displayName"`
	Enabled                  pulumi.StringPtrOutput    `pulumi:"enabled"`
	Etag                     pulumi.StringOutput       `pulumi:"etag"`
	IsLegacyLogAnalyticsRule pulumi.BoolOutput         `pulumi:"isLegacyLogAnalyticsRule"`
	Kind                     pulumi.StringOutput       `pulumi:"kind"`
	LastUpdatedTime          pulumi.StringOutput       `pulumi:"lastUpdatedTime"`
	Location                 pulumi.StringOutput       `pulumi:"location"`
	Name                     pulumi.StringOutput       `pulumi:"name"`
	ProvisioningState        pulumi.StringOutput       `pulumi:"provisioningState"`
	Schedule                 ScheduleResponsePtrOutput `pulumi:"schedule"`
	Source                   SourceResponseOutput      `pulumi:"source"`
	Tags                     pulumi.StringMapOutput    `pulumi:"tags"`
	Type                     pulumi.StringOutput       `pulumi:"type"`
}


func NewScheduledQueryRule(ctx *pulumi.Context,
	name string, args *ScheduledQueryRuleArgs, opts ...pulumi.ResourceOption) (*ScheduledQueryRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	if args.AutoMitigate == nil {
		args.AutoMitigate = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:ScheduledQueryRule"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20200501preview:ScheduledQueryRule"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210201preview:ScheduledQueryRule"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210801:ScheduledQueryRule"),
		},
	})
	opts = append(opts, aliases)
	var resource ScheduledQueryRule
	err := ctx.RegisterResource("azure-native:insights/v20180416:ScheduledQueryRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetScheduledQueryRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledQueryRuleState, opts ...pulumi.ResourceOption) (*ScheduledQueryRule, error) {
	var resource ScheduledQueryRule
	err := ctx.ReadResource("azure-native:insights/v20180416:ScheduledQueryRule", name, id, state, &resource, opts...)
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
	Action            interface{}       `pulumi:"action"`
	AutoMitigate      *bool             `pulumi:"autoMitigate"`
	Description       *string           `pulumi:"description"`
	DisplayName       *string           `pulumi:"displayName"`
	Enabled           *string           `pulumi:"enabled"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	RuleName          *string           `pulumi:"ruleName"`
	Schedule          *Schedule         `pulumi:"schedule"`
	Source            Source            `pulumi:"source"`
	Tags              map[string]string `pulumi:"tags"`
}


type ScheduledQueryRuleArgs struct {
	Action            pulumi.Input
	AutoMitigate      pulumi.BoolPtrInput
	Description       pulumi.StringPtrInput
	DisplayName       pulumi.StringPtrInput
	Enabled           pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	RuleName          pulumi.StringPtrInput
	Schedule          SchedulePtrInput
	Source            SourceInput
	Tags              pulumi.StringMapInput
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
	pulumi.RegisterOutputType(ScheduledQueryRuleOutput{})
}




package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EventHubAuthorizationRule struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput      `pulumi:"location"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	Rights     pulumi.StringArrayOutput `pulumi:"rights"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewEventHubAuthorizationRule(ctx *pulumi.Context,
	name string, args *EventHubAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*EventHubAuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventHubName == nil {
		return nil, errors.New("invalid value for required argument 'EventHubName'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Rights == nil {
		return nil, errors.New("invalid value for required argument 'Rights'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventhub:EventHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20140901:EventHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20150801:EventHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20170401:EventHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20180101preview:EventHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210101preview:EventHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210601preview:EventHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20211101:EventHubAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	var resource EventHubAuthorizationRule
	err := ctx.RegisterResource("azure-native:eventhub/v20220101preview:EventHubAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEventHubAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventHubAuthorizationRuleState, opts ...pulumi.ResourceOption) (*EventHubAuthorizationRule, error) {
	var resource EventHubAuthorizationRule
	err := ctx.ReadResource("azure-native:eventhub/v20220101preview:EventHubAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type eventHubAuthorizationRuleState struct {
}

type EventHubAuthorizationRuleState struct {
}

func (EventHubAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubAuthorizationRuleState)(nil)).Elem()
}

type eventHubAuthorizationRuleArgs struct {
	AuthorizationRuleName *string  `pulumi:"authorizationRuleName"`
	EventHubName          string   `pulumi:"eventHubName"`
	NamespaceName         string   `pulumi:"namespaceName"`
	ResourceGroupName     string   `pulumi:"resourceGroupName"`
	Rights                []string `pulumi:"rights"`
}


type EventHubAuthorizationRuleArgs struct {
	AuthorizationRuleName pulumi.StringPtrInput
	EventHubName          pulumi.StringInput
	NamespaceName         pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	Rights                pulumi.StringArrayInput
}

func (EventHubAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubAuthorizationRuleArgs)(nil)).Elem()
}

type EventHubAuthorizationRuleInput interface {
	pulumi.Input

	ToEventHubAuthorizationRuleOutput() EventHubAuthorizationRuleOutput
	ToEventHubAuthorizationRuleOutputWithContext(ctx context.Context) EventHubAuthorizationRuleOutput
}

func (*EventHubAuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHubAuthorizationRule)(nil)).Elem()
}

func (i *EventHubAuthorizationRule) ToEventHubAuthorizationRuleOutput() EventHubAuthorizationRuleOutput {
	return i.ToEventHubAuthorizationRuleOutputWithContext(context.Background())
}

func (i *EventHubAuthorizationRule) ToEventHubAuthorizationRuleOutputWithContext(ctx context.Context) EventHubAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubAuthorizationRuleOutput)
}

type EventHubAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (EventHubAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHubAuthorizationRule)(nil)).Elem()
}

func (o EventHubAuthorizationRuleOutput) ToEventHubAuthorizationRuleOutput() EventHubAuthorizationRuleOutput {
	return o
}

func (o EventHubAuthorizationRuleOutput) ToEventHubAuthorizationRuleOutputWithContext(ctx context.Context) EventHubAuthorizationRuleOutput {
	return o
}

func (o EventHubAuthorizationRuleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubAuthorizationRule) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o EventHubAuthorizationRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubAuthorizationRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EventHubAuthorizationRuleOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EventHubAuthorizationRule) pulumi.StringArrayOutput { return v.Rights }).(pulumi.StringArrayOutput)
}

func (o EventHubAuthorizationRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EventHubAuthorizationRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o EventHubAuthorizationRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubAuthorizationRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EventHubAuthorizationRuleOutput{})
}

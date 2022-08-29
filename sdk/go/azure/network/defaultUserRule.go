


package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DefaultUserRule struct {
	pulumi.CustomResourceState

	Description           pulumi.StringOutput                  `pulumi:"description"`
	DestinationPortRanges pulumi.StringArrayOutput             `pulumi:"destinationPortRanges"`
	Destinations          AddressPrefixItemResponseArrayOutput `pulumi:"destinations"`
	Direction             pulumi.StringOutput                  `pulumi:"direction"`
	DisplayName           pulumi.StringOutput                  `pulumi:"displayName"`
	Etag                  pulumi.StringOutput                  `pulumi:"etag"`
	Flag                  pulumi.StringPtrOutput               `pulumi:"flag"`
	Kind                  pulumi.StringOutput                  `pulumi:"kind"`
	Name                  pulumi.StringOutput                  `pulumi:"name"`
	Protocol              pulumi.StringOutput                  `pulumi:"protocol"`
	ProvisioningState     pulumi.StringOutput                  `pulumi:"provisioningState"`
	SourcePortRanges      pulumi.StringArrayOutput             `pulumi:"sourcePortRanges"`
	Sources               AddressPrefixItemResponseArrayOutput `pulumi:"sources"`
	SystemData            SystemDataResponseOutput             `pulumi:"systemData"`
	Type                  pulumi.StringOutput                  `pulumi:"type"`
}


func NewDefaultUserRule(ctx *pulumi.Context,
	name string, args *DefaultUserRuleArgs, opts ...pulumi.ResourceOption) (*DefaultUserRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.NetworkManagerName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RuleCollectionName == nil {
		return nil, errors.New("invalid value for required argument 'RuleCollectionName'")
	}
	args.Kind = pulumi.String("Default")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20210201preview:DefaultUserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:DefaultUserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:DefaultUserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:DefaultUserRule"),
		},
	})
	opts = append(opts, aliases)
	var resource DefaultUserRule
	err := ctx.RegisterResource("azure-native:network:DefaultUserRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDefaultUserRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultUserRuleState, opts ...pulumi.ResourceOption) (*DefaultUserRule, error) {
	var resource DefaultUserRule
	err := ctx.ReadResource("azure-native:network:DefaultUserRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type defaultUserRuleState struct {
}

type DefaultUserRuleState struct {
}

func (DefaultUserRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultUserRuleState)(nil)).Elem()
}

type defaultUserRuleArgs struct {
	ConfigurationName  string  `pulumi:"configurationName"`
	Flag               *string `pulumi:"flag"`
	Kind               string  `pulumi:"kind"`
	NetworkManagerName string  `pulumi:"networkManagerName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	RuleCollectionName string  `pulumi:"ruleCollectionName"`
	RuleName           *string `pulumi:"ruleName"`
}


type DefaultUserRuleArgs struct {
	ConfigurationName  pulumi.StringInput
	Flag               pulumi.StringPtrInput
	Kind               pulumi.StringInput
	NetworkManagerName pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	RuleCollectionName pulumi.StringInput
	RuleName           pulumi.StringPtrInput
}

func (DefaultUserRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultUserRuleArgs)(nil)).Elem()
}

type DefaultUserRuleInput interface {
	pulumi.Input

	ToDefaultUserRuleOutput() DefaultUserRuleOutput
	ToDefaultUserRuleOutputWithContext(ctx context.Context) DefaultUserRuleOutput
}

func (*DefaultUserRule) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultUserRule)(nil)).Elem()
}

func (i *DefaultUserRule) ToDefaultUserRuleOutput() DefaultUserRuleOutput {
	return i.ToDefaultUserRuleOutputWithContext(context.Background())
}

func (i *DefaultUserRule) ToDefaultUserRuleOutputWithContext(ctx context.Context) DefaultUserRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultUserRuleOutput)
}

type DefaultUserRuleOutput struct{ *pulumi.OutputState }

func (DefaultUserRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultUserRule)(nil)).Elem()
}

func (o DefaultUserRuleOutput) ToDefaultUserRuleOutput() DefaultUserRuleOutput {
	return o
}

func (o DefaultUserRuleOutput) ToDefaultUserRuleOutputWithContext(ctx context.Context) DefaultUserRuleOutput {
	return o
}

func (o DefaultUserRuleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultUserRule) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o DefaultUserRuleOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DefaultUserRule) pulumi.StringArrayOutput { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

func (o DefaultUserRuleOutput) Destinations() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v *DefaultUserRule) AddressPrefixItemResponseArrayOutput { return v.Destinations }).(AddressPrefixItemResponseArrayOutput)
}

func (o DefaultUserRuleOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultUserRule) pulumi.StringOutput { return v.Direction }).(pulumi.StringOutput)
}

func (o DefaultUserRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultUserRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o DefaultUserRuleOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultUserRule) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DefaultUserRuleOutput) Flag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultUserRule) pulumi.StringPtrOutput { return v.Flag }).(pulumi.StringPtrOutput)
}

func (o DefaultUserRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultUserRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o DefaultUserRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultUserRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DefaultUserRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultUserRule) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

func (o DefaultUserRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultUserRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DefaultUserRuleOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DefaultUserRule) pulumi.StringArrayOutput { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

func (o DefaultUserRuleOutput) Sources() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v *DefaultUserRule) AddressPrefixItemResponseArrayOutput { return v.Sources }).(AddressPrefixItemResponseArrayOutput)
}

func (o DefaultUserRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DefaultUserRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DefaultUserRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultUserRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DefaultUserRuleOutput{})
}




package v20220501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DefaultAdminRule struct {
	pulumi.CustomResourceState

	Access                pulumi.StringOutput                  `pulumi:"access"`
	Description           pulumi.StringOutput                  `pulumi:"description"`
	DestinationPortRanges pulumi.StringArrayOutput             `pulumi:"destinationPortRanges"`
	Destinations          AddressPrefixItemResponseArrayOutput `pulumi:"destinations"`
	Direction             pulumi.StringOutput                  `pulumi:"direction"`
	Etag                  pulumi.StringOutput                  `pulumi:"etag"`
	Flag                  pulumi.StringPtrOutput               `pulumi:"flag"`
	Kind                  pulumi.StringOutput                  `pulumi:"kind"`
	Name                  pulumi.StringOutput                  `pulumi:"name"`
	Priority              pulumi.IntOutput                     `pulumi:"priority"`
	Protocol              pulumi.StringOutput                  `pulumi:"protocol"`
	ProvisioningState     pulumi.StringOutput                  `pulumi:"provisioningState"`
	SourcePortRanges      pulumi.StringArrayOutput             `pulumi:"sourcePortRanges"`
	Sources               AddressPrefixItemResponseArrayOutput `pulumi:"sources"`
	SystemData            SystemDataResponseOutput             `pulumi:"systemData"`
	Type                  pulumi.StringOutput                  `pulumi:"type"`
}


func NewDefaultAdminRule(ctx *pulumi.Context,
	name string, args *DefaultAdminRuleArgs, opts ...pulumi.ResourceOption) (*DefaultAdminRule, error) {
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
			Type: pulumi.String("azure-native:network:DefaultAdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201preview:DefaultAdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:DefaultAdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:DefaultAdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:DefaultAdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:DefaultAdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:DefaultAdminRule"),
		},
	})
	opts = append(opts, aliases)
	var resource DefaultAdminRule
	err := ctx.RegisterResource("azure-native:network/v20220501:DefaultAdminRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDefaultAdminRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultAdminRuleState, opts ...pulumi.ResourceOption) (*DefaultAdminRule, error) {
	var resource DefaultAdminRule
	err := ctx.ReadResource("azure-native:network/v20220501:DefaultAdminRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type defaultAdminRuleState struct {
}

type DefaultAdminRuleState struct {
}

func (DefaultAdminRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultAdminRuleState)(nil)).Elem()
}

type defaultAdminRuleArgs struct {
	ConfigurationName  string  `pulumi:"configurationName"`
	Flag               *string `pulumi:"flag"`
	Kind               string  `pulumi:"kind"`
	NetworkManagerName string  `pulumi:"networkManagerName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	RuleCollectionName string  `pulumi:"ruleCollectionName"`
	RuleName           *string `pulumi:"ruleName"`
}


type DefaultAdminRuleArgs struct {
	ConfigurationName  pulumi.StringInput
	Flag               pulumi.StringPtrInput
	Kind               pulumi.StringInput
	NetworkManagerName pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	RuleCollectionName pulumi.StringInput
	RuleName           pulumi.StringPtrInput
}

func (DefaultAdminRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultAdminRuleArgs)(nil)).Elem()
}

type DefaultAdminRuleInput interface {
	pulumi.Input

	ToDefaultAdminRuleOutput() DefaultAdminRuleOutput
	ToDefaultAdminRuleOutputWithContext(ctx context.Context) DefaultAdminRuleOutput
}

func (*DefaultAdminRule) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultAdminRule)(nil)).Elem()
}

func (i *DefaultAdminRule) ToDefaultAdminRuleOutput() DefaultAdminRuleOutput {
	return i.ToDefaultAdminRuleOutputWithContext(context.Background())
}

func (i *DefaultAdminRule) ToDefaultAdminRuleOutputWithContext(ctx context.Context) DefaultAdminRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultAdminRuleOutput)
}

type DefaultAdminRuleOutput struct{ *pulumi.OutputState }

func (DefaultAdminRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultAdminRule)(nil)).Elem()
}

func (o DefaultAdminRuleOutput) ToDefaultAdminRuleOutput() DefaultAdminRuleOutput {
	return o
}

func (o DefaultAdminRuleOutput) ToDefaultAdminRuleOutputWithContext(ctx context.Context) DefaultAdminRuleOutput {
	return o
}

func (o DefaultAdminRuleOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringOutput { return v.Access }).(pulumi.StringOutput)
}

func (o DefaultAdminRuleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o DefaultAdminRuleOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringArrayOutput { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

func (o DefaultAdminRuleOutput) Destinations() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v *DefaultAdminRule) AddressPrefixItemResponseArrayOutput { return v.Destinations }).(AddressPrefixItemResponseArrayOutput)
}

func (o DefaultAdminRuleOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringOutput { return v.Direction }).(pulumi.StringOutput)
}

func (o DefaultAdminRuleOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DefaultAdminRuleOutput) Flag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringPtrOutput { return v.Flag }).(pulumi.StringPtrOutput)
}

func (o DefaultAdminRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o DefaultAdminRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DefaultAdminRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

func (o DefaultAdminRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

func (o DefaultAdminRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DefaultAdminRuleOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringArrayOutput { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

func (o DefaultAdminRuleOutput) Sources() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v *DefaultAdminRule) AddressPrefixItemResponseArrayOutput { return v.Sources }).(AddressPrefixItemResponseArrayOutput)
}

func (o DefaultAdminRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DefaultAdminRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DefaultAdminRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DefaultAdminRuleOutput{})
}

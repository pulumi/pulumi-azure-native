


package network

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
	DisplayName           pulumi.StringOutput                  `pulumi:"displayName"`
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
			Type: pulumi.String("azure-nextgen:network:DefaultAdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201preview:DefaultAdminRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210201preview:DefaultAdminRule"),
		},
	})
	opts = append(opts, aliases)
	var resource DefaultAdminRule
	err := ctx.RegisterResource("azure-native:network:DefaultAdminRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDefaultAdminRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultAdminRuleState, opts ...pulumi.ResourceOption) (*DefaultAdminRule, error) {
	var resource DefaultAdminRule
	err := ctx.ReadResource("azure-native:network:DefaultAdminRule", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*DefaultAdminRule)(nil))
}

func (i *DefaultAdminRule) ToDefaultAdminRuleOutput() DefaultAdminRuleOutput {
	return i.ToDefaultAdminRuleOutputWithContext(context.Background())
}

func (i *DefaultAdminRule) ToDefaultAdminRuleOutputWithContext(ctx context.Context) DefaultAdminRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultAdminRuleOutput)
}

type DefaultAdminRuleOutput struct{ *pulumi.OutputState }

func (DefaultAdminRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultAdminRule)(nil))
}

func (o DefaultAdminRuleOutput) ToDefaultAdminRuleOutput() DefaultAdminRuleOutput {
	return o
}

func (o DefaultAdminRuleOutput) ToDefaultAdminRuleOutputWithContext(ctx context.Context) DefaultAdminRuleOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultAdminRuleInput)(nil)).Elem(), &DefaultAdminRule{})
	pulumi.RegisterOutputType(DefaultAdminRuleOutput{})
}

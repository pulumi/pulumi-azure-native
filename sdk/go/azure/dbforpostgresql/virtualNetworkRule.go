


package dbforpostgresql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualNetworkRule struct {
	pulumi.CustomResourceState

	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrOutput `pulumi:"ignoreMissingVnetServiceEndpoint"`
	Name                             pulumi.StringOutput  `pulumi:"name"`
	State                            pulumi.StringOutput  `pulumi:"state"`
	Type                             pulumi.StringOutput  `pulumi:"type"`
	VirtualNetworkSubnetId           pulumi.StringOutput  `pulumi:"virtualNetworkSubnetId"`
}


func NewVirtualNetworkRule(ctx *pulumi.Context,
	name string, args *VirtualNetworkRuleArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.VirtualNetworkSubnetId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualNetworkSubnetId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20171201:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20171201preview:VirtualNetworkRule"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetworkRule
	err := ctx.RegisterResource("azure-native:dbforpostgresql:VirtualNetworkRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualNetworkRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkRuleState, opts ...pulumi.ResourceOption) (*VirtualNetworkRule, error) {
	var resource VirtualNetworkRule
	err := ctx.ReadResource("azure-native:dbforpostgresql:VirtualNetworkRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualNetworkRuleState struct {
}

type VirtualNetworkRuleState struct {
}

func (VirtualNetworkRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkRuleState)(nil)).Elem()
}

type virtualNetworkRuleArgs struct {
	IgnoreMissingVnetServiceEndpoint *bool   `pulumi:"ignoreMissingVnetServiceEndpoint"`
	ResourceGroupName                string  `pulumi:"resourceGroupName"`
	ServerName                       string  `pulumi:"serverName"`
	VirtualNetworkRuleName           *string `pulumi:"virtualNetworkRuleName"`
	VirtualNetworkSubnetId           string  `pulumi:"virtualNetworkSubnetId"`
}


type VirtualNetworkRuleArgs struct {
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrInput
	ResourceGroupName                pulumi.StringInput
	ServerName                       pulumi.StringInput
	VirtualNetworkRuleName           pulumi.StringPtrInput
	VirtualNetworkSubnetId           pulumi.StringInput
}

func (VirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkRuleArgs)(nil)).Elem()
}

type VirtualNetworkRuleInput interface {
	pulumi.Input

	ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput
	ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput
}

func (*VirtualNetworkRule) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkRule)(nil)).Elem()
}

func (i *VirtualNetworkRule) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return i.ToVirtualNetworkRuleOutputWithContext(context.Background())
}

func (i *VirtualNetworkRule) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleOutput)
}

type VirtualNetworkRuleOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) IgnoreMissingVnetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkRule) pulumi.BoolPtrOutput { return v.IgnoreMissingVnetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkRuleOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkRule) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o VirtualNetworkRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualNetworkRuleOutput) VirtualNetworkSubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkRule) pulumi.StringOutput { return v.VirtualNetworkSubnetId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
}

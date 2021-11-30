


package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IpFirewallRule struct {
	pulumi.CustomResourceState

	EndIpAddress      pulumi.StringPtrOutput `pulumi:"endIpAddress"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	StartIpAddress    pulumi.StringPtrOutput `pulumi:"startIpAddress"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewIpFirewallRule(ctx *pulumi.Context,
	name string, args *IpFirewallRuleArgs, opts ...pulumi.ResourceOption) (*IpFirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse:IpFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:IpFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:IpFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:IpFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210501:IpFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601:IpFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:IpFirewallRule"),
		},
	})
	opts = append(opts, aliases)
	var resource IpFirewallRule
	err := ctx.RegisterResource("azure-native:synapse/v20201201:IpFirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIpFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpFirewallRuleState, opts ...pulumi.ResourceOption) (*IpFirewallRule, error) {
	var resource IpFirewallRule
	err := ctx.ReadResource("azure-native:synapse/v20201201:IpFirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type ipFirewallRuleState struct {
}

type IpFirewallRuleState struct {
}

func (IpFirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipFirewallRuleState)(nil)).Elem()
}

type ipFirewallRuleArgs struct {
	EndIpAddress      *string `pulumi:"endIpAddress"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	RuleName          *string `pulumi:"ruleName"`
	StartIpAddress    *string `pulumi:"startIpAddress"`
	WorkspaceName     string  `pulumi:"workspaceName"`
}


type IpFirewallRuleArgs struct {
	EndIpAddress      pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	RuleName          pulumi.StringPtrInput
	StartIpAddress    pulumi.StringPtrInput
	WorkspaceName     pulumi.StringInput
}

func (IpFirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipFirewallRuleArgs)(nil)).Elem()
}

type IpFirewallRuleInput interface {
	pulumi.Input

	ToIpFirewallRuleOutput() IpFirewallRuleOutput
	ToIpFirewallRuleOutputWithContext(ctx context.Context) IpFirewallRuleOutput
}

func (*IpFirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((*IpFirewallRule)(nil))
}

func (i *IpFirewallRule) ToIpFirewallRuleOutput() IpFirewallRuleOutput {
	return i.ToIpFirewallRuleOutputWithContext(context.Background())
}

func (i *IpFirewallRule) ToIpFirewallRuleOutputWithContext(ctx context.Context) IpFirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpFirewallRuleOutput)
}

type IpFirewallRuleOutput struct{ *pulumi.OutputState }

func (IpFirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpFirewallRule)(nil))
}

func (o IpFirewallRuleOutput) ToIpFirewallRuleOutput() IpFirewallRuleOutput {
	return o
}

func (o IpFirewallRuleOutput) ToIpFirewallRuleOutputWithContext(ctx context.Context) IpFirewallRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IpFirewallRuleOutput{})
}

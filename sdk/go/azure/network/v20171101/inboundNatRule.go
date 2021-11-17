


package v20171101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InboundNatRule struct {
	pulumi.CustomResourceState

	BackendIPConfiguration  NetworkInterfaceIPConfigurationResponseOutput `pulumi:"backendIPConfiguration"`
	BackendPort             pulumi.IntPtrOutput                           `pulumi:"backendPort"`
	EnableFloatingIP        pulumi.BoolPtrOutput                          `pulumi:"enableFloatingIP"`
	Etag                    pulumi.StringPtrOutput                        `pulumi:"etag"`
	FrontendIPConfiguration SubResourceResponsePtrOutput                  `pulumi:"frontendIPConfiguration"`
	FrontendPort            pulumi.IntPtrOutput                           `pulumi:"frontendPort"`
	IdleTimeoutInMinutes    pulumi.IntPtrOutput                           `pulumi:"idleTimeoutInMinutes"`
	Name                    pulumi.StringPtrOutput                        `pulumi:"name"`
	Protocol                pulumi.StringPtrOutput                        `pulumi:"protocol"`
	ProvisioningState       pulumi.StringPtrOutput                        `pulumi:"provisioningState"`
}


func NewInboundNatRule(ctx *pulumi.Context,
	name string, args *InboundNatRuleArgs, opts ...pulumi.ResourceOption) (*InboundNatRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoadBalancerName == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:InboundNatRule"),
		},
	})
	opts = append(opts, aliases)
	var resource InboundNatRule
	err := ctx.RegisterResource("azure-native:network/v20171101:InboundNatRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetInboundNatRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InboundNatRuleState, opts ...pulumi.ResourceOption) (*InboundNatRule, error) {
	var resource InboundNatRule
	err := ctx.ReadResource("azure-native:network/v20171101:InboundNatRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type inboundNatRuleState struct {
}

type InboundNatRuleState struct {
}

func (InboundNatRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*inboundNatRuleState)(nil)).Elem()
}

type inboundNatRuleArgs struct {
	BackendPort             *int         `pulumi:"backendPort"`
	EnableFloatingIP        *bool        `pulumi:"enableFloatingIP"`
	Etag                    *string      `pulumi:"etag"`
	FrontendIPConfiguration *SubResource `pulumi:"frontendIPConfiguration"`
	FrontendPort            *int         `pulumi:"frontendPort"`
	Id                      *string      `pulumi:"id"`
	IdleTimeoutInMinutes    *int         `pulumi:"idleTimeoutInMinutes"`
	InboundNatRuleName      *string      `pulumi:"inboundNatRuleName"`
	LoadBalancerName        string       `pulumi:"loadBalancerName"`
	Name                    *string      `pulumi:"name"`
	Protocol                *string      `pulumi:"protocol"`
	ProvisioningState       *string      `pulumi:"provisioningState"`
	ResourceGroupName       string       `pulumi:"resourceGroupName"`
}


type InboundNatRuleArgs struct {
	BackendPort             pulumi.IntPtrInput
	EnableFloatingIP        pulumi.BoolPtrInput
	Etag                    pulumi.StringPtrInput
	FrontendIPConfiguration SubResourcePtrInput
	FrontendPort            pulumi.IntPtrInput
	Id                      pulumi.StringPtrInput
	IdleTimeoutInMinutes    pulumi.IntPtrInput
	InboundNatRuleName      pulumi.StringPtrInput
	LoadBalancerName        pulumi.StringInput
	Name                    pulumi.StringPtrInput
	Protocol                pulumi.StringPtrInput
	ProvisioningState       pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
}

func (InboundNatRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inboundNatRuleArgs)(nil)).Elem()
}

type InboundNatRuleInput interface {
	pulumi.Input

	ToInboundNatRuleOutput() InboundNatRuleOutput
	ToInboundNatRuleOutputWithContext(ctx context.Context) InboundNatRuleOutput
}

func (*InboundNatRule) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatRule)(nil))
}

func (i *InboundNatRule) ToInboundNatRuleOutput() InboundNatRuleOutput {
	return i.ToInboundNatRuleOutputWithContext(context.Background())
}

func (i *InboundNatRule) ToInboundNatRuleOutputWithContext(ctx context.Context) InboundNatRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundNatRuleOutput)
}

type InboundNatRuleOutput struct{ *pulumi.OutputState }

func (InboundNatRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatRule)(nil))
}

func (o InboundNatRuleOutput) ToInboundNatRuleOutput() InboundNatRuleOutput {
	return o
}

func (o InboundNatRuleOutput) ToInboundNatRuleOutputWithContext(ctx context.Context) InboundNatRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InboundNatRuleOutput{})
}

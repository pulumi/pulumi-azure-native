


package v20200701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LoadBalancerBackendAddressPool struct {
	pulumi.CustomResourceState

	BackendIPConfigurations      NetworkInterfaceIPConfigurationResponseArrayOutput `pulumi:"backendIPConfigurations"`
	Etag                         pulumi.StringOutput                                `pulumi:"etag"`
	LoadBalancerBackendAddresses LoadBalancerBackendAddressResponseArrayOutput      `pulumi:"loadBalancerBackendAddresses"`
	LoadBalancingRules           SubResourceResponseArrayOutput                     `pulumi:"loadBalancingRules"`
	Name                         pulumi.StringPtrOutput                             `pulumi:"name"`
	OutboundRule                 SubResourceResponseOutput                          `pulumi:"outboundRule"`
	OutboundRules                SubResourceResponseArrayOutput                     `pulumi:"outboundRules"`
	ProvisioningState            pulumi.StringOutput                                `pulumi:"provisioningState"`
	Type                         pulumi.StringOutput                                `pulumi:"type"`
}


func NewLoadBalancerBackendAddressPool(ctx *pulumi.Context,
	name string, args *LoadBalancerBackendAddressPoolArgs, opts ...pulumi.ResourceOption) (*LoadBalancerBackendAddressPool, error) {
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
			Type: pulumi.String("azure-nextgen:network/v20200701:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20201101:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210201:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210301:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210501:LoadBalancerBackendAddressPool"),
		},
	})
	opts = append(opts, aliases)
	var resource LoadBalancerBackendAddressPool
	err := ctx.RegisterResource("azure-native:network/v20200701:LoadBalancerBackendAddressPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLoadBalancerBackendAddressPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerBackendAddressPoolState, opts ...pulumi.ResourceOption) (*LoadBalancerBackendAddressPool, error) {
	var resource LoadBalancerBackendAddressPool
	err := ctx.ReadResource("azure-native:network/v20200701:LoadBalancerBackendAddressPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type loadBalancerBackendAddressPoolState struct {
}

type LoadBalancerBackendAddressPoolState struct {
}

func (LoadBalancerBackendAddressPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerBackendAddressPoolState)(nil)).Elem()
}

type loadBalancerBackendAddressPoolArgs struct {
	BackendAddressPoolName       *string                      `pulumi:"backendAddressPoolName"`
	Id                           *string                      `pulumi:"id"`
	LoadBalancerBackendAddresses []LoadBalancerBackendAddress `pulumi:"loadBalancerBackendAddresses"`
	LoadBalancerName             string                       `pulumi:"loadBalancerName"`
	Name                         *string                      `pulumi:"name"`
	ResourceGroupName            string                       `pulumi:"resourceGroupName"`
}


type LoadBalancerBackendAddressPoolArgs struct {
	BackendAddressPoolName       pulumi.StringPtrInput
	Id                           pulumi.StringPtrInput
	LoadBalancerBackendAddresses LoadBalancerBackendAddressArrayInput
	LoadBalancerName             pulumi.StringInput
	Name                         pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
}

func (LoadBalancerBackendAddressPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerBackendAddressPoolArgs)(nil)).Elem()
}

type LoadBalancerBackendAddressPoolInput interface {
	pulumi.Input

	ToLoadBalancerBackendAddressPoolOutput() LoadBalancerBackendAddressPoolOutput
	ToLoadBalancerBackendAddressPoolOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolOutput
}

func (*LoadBalancerBackendAddressPool) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerBackendAddressPool)(nil))
}

func (i *LoadBalancerBackendAddressPool) ToLoadBalancerBackendAddressPoolOutput() LoadBalancerBackendAddressPoolOutput {
	return i.ToLoadBalancerBackendAddressPoolOutputWithContext(context.Background())
}

func (i *LoadBalancerBackendAddressPool) ToLoadBalancerBackendAddressPoolOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerBackendAddressPoolOutput)
}

type LoadBalancerBackendAddressPoolOutput struct{ *pulumi.OutputState }

func (LoadBalancerBackendAddressPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerBackendAddressPool)(nil))
}

func (o LoadBalancerBackendAddressPoolOutput) ToLoadBalancerBackendAddressPoolOutput() LoadBalancerBackendAddressPoolOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolOutput) ToLoadBalancerBackendAddressPoolOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LoadBalancerBackendAddressPoolOutput{})
}




package v20201101

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
			Type: pulumi.String("azure-native:network:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:LoadBalancerBackendAddressPool"),
		},
	})
	opts = append(opts, aliases)
	var resource LoadBalancerBackendAddressPool
	err := ctx.RegisterResource("azure-native:network/v20201101:LoadBalancerBackendAddressPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLoadBalancerBackendAddressPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerBackendAddressPoolState, opts ...pulumi.ResourceOption) (*LoadBalancerBackendAddressPool, error) {
	var resource LoadBalancerBackendAddressPool
	err := ctx.ReadResource("azure-native:network/v20201101:LoadBalancerBackendAddressPool", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((**LoadBalancerBackendAddressPool)(nil)).Elem()
}

func (i *LoadBalancerBackendAddressPool) ToLoadBalancerBackendAddressPoolOutput() LoadBalancerBackendAddressPoolOutput {
	return i.ToLoadBalancerBackendAddressPoolOutputWithContext(context.Background())
}

func (i *LoadBalancerBackendAddressPool) ToLoadBalancerBackendAddressPoolOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerBackendAddressPoolOutput)
}

type LoadBalancerBackendAddressPoolOutput struct{ *pulumi.OutputState }

func (LoadBalancerBackendAddressPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerBackendAddressPool)(nil)).Elem()
}

func (o LoadBalancerBackendAddressPoolOutput) ToLoadBalancerBackendAddressPoolOutput() LoadBalancerBackendAddressPoolOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolOutput) ToLoadBalancerBackendAddressPoolOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolOutput) BackendIPConfigurations() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) NetworkInterfaceIPConfigurationResponseArrayOutput {
		return v.BackendIPConfigurations
	}).(NetworkInterfaceIPConfigurationResponseArrayOutput)
}

func (o LoadBalancerBackendAddressPoolOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o LoadBalancerBackendAddressPoolOutput) LoadBalancerBackendAddresses() LoadBalancerBackendAddressResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) LoadBalancerBackendAddressResponseArrayOutput {
		return v.LoadBalancerBackendAddresses
	}).(LoadBalancerBackendAddressResponseArrayOutput)
}

func (o LoadBalancerBackendAddressPoolOutput) LoadBalancingRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) SubResourceResponseArrayOutput { return v.LoadBalancingRules }).(SubResourceResponseArrayOutput)
}

func (o LoadBalancerBackendAddressPoolOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerBackendAddressPoolOutput) OutboundRule() SubResourceResponseOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) SubResourceResponseOutput { return v.OutboundRule }).(SubResourceResponseOutput)
}

func (o LoadBalancerBackendAddressPoolOutput) OutboundRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) SubResourceResponseArrayOutput { return v.OutboundRules }).(SubResourceResponseArrayOutput)
}

func (o LoadBalancerBackendAddressPoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LoadBalancerBackendAddressPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LoadBalancerBackendAddressPoolOutput{})
}

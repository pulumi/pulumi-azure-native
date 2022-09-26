


package v20190801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateLinkService struct {
	pulumi.CustomResourceState

	Alias                                pulumi.StringOutput                                       `pulumi:"alias"`
	AutoApproval                         PrivateLinkServicePropertiesResponseAutoApprovalPtrOutput `pulumi:"autoApproval"`
	Etag                                 pulumi.StringPtrOutput                                    `pulumi:"etag"`
	Fqdns                                pulumi.StringArrayOutput                                  `pulumi:"fqdns"`
	IpConfigurations                     PrivateLinkServiceIpConfigurationResponseArrayOutput      `pulumi:"ipConfigurations"`
	LoadBalancerFrontendIpConfigurations FrontendIPConfigurationResponseArrayOutput                `pulumi:"loadBalancerFrontendIpConfigurations"`
	Location                             pulumi.StringPtrOutput                                    `pulumi:"location"`
	Name                                 pulumi.StringOutput                                       `pulumi:"name"`
	NetworkInterfaces                    NetworkInterfaceResponseArrayOutput                       `pulumi:"networkInterfaces"`
	PrivateEndpointConnections           PrivateEndpointConnectionResponseArrayOutput              `pulumi:"privateEndpointConnections"`
	ProvisioningState                    pulumi.StringOutput                                       `pulumi:"provisioningState"`
	Tags                                 pulumi.StringMapOutput                                    `pulumi:"tags"`
	Type                                 pulumi.StringOutput                                       `pulumi:"type"`
	Visibility                           PrivateLinkServicePropertiesResponseVisibilityPtrOutput   `pulumi:"visibility"`
}


func NewPrivateLinkService(ctx *pulumi.Context,
	name string, args *PrivateLinkServiceArgs, opts ...pulumi.ResourceOption) (*PrivateLinkService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:PrivateLinkService"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateLinkService
	err := ctx.RegisterResource("azure-native:network/v20190801:PrivateLinkService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateLinkService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkServiceState, opts ...pulumi.ResourceOption) (*PrivateLinkService, error) {
	var resource PrivateLinkService
	err := ctx.ReadResource("azure-native:network/v20190801:PrivateLinkService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateLinkServiceState struct {
}

type PrivateLinkServiceState struct {
}

func (PrivateLinkServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServiceState)(nil)).Elem()
}

type privateLinkServiceArgs struct {
	AutoApproval                         *PrivateLinkServicePropertiesAutoApproval `pulumi:"autoApproval"`
	Fqdns                                []string                                  `pulumi:"fqdns"`
	Id                                   *string                                   `pulumi:"id"`
	IpConfigurations                     []PrivateLinkServiceIpConfiguration       `pulumi:"ipConfigurations"`
	LoadBalancerFrontendIpConfigurations []FrontendIPConfiguration                 `pulumi:"loadBalancerFrontendIpConfigurations"`
	Location                             *string                                   `pulumi:"location"`
	PrivateEndpointConnections           []PrivateEndpointConnection               `pulumi:"privateEndpointConnections"`
	ResourceGroupName                    string                                    `pulumi:"resourceGroupName"`
	ServiceName                          *string                                   `pulumi:"serviceName"`
	Tags                                 map[string]string                         `pulumi:"tags"`
	Visibility                           *PrivateLinkServicePropertiesVisibility   `pulumi:"visibility"`
}


type PrivateLinkServiceArgs struct {
	AutoApproval                         PrivateLinkServicePropertiesAutoApprovalPtrInput
	Fqdns                                pulumi.StringArrayInput
	Id                                   pulumi.StringPtrInput
	IpConfigurations                     PrivateLinkServiceIpConfigurationArrayInput
	LoadBalancerFrontendIpConfigurations FrontendIPConfigurationArrayInput
	Location                             pulumi.StringPtrInput
	PrivateEndpointConnections           PrivateEndpointConnectionArrayInput
	ResourceGroupName                    pulumi.StringInput
	ServiceName                          pulumi.StringPtrInput
	Tags                                 pulumi.StringMapInput
	Visibility                           PrivateLinkServicePropertiesVisibilityPtrInput
}

func (PrivateLinkServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServiceArgs)(nil)).Elem()
}

type PrivateLinkServiceInput interface {
	pulumi.Input

	ToPrivateLinkServiceOutput() PrivateLinkServiceOutput
	ToPrivateLinkServiceOutputWithContext(ctx context.Context) PrivateLinkServiceOutput
}

func (*PrivateLinkService) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkService)(nil)).Elem()
}

func (i *PrivateLinkService) ToPrivateLinkServiceOutput() PrivateLinkServiceOutput {
	return i.ToPrivateLinkServiceOutputWithContext(context.Background())
}

func (i *PrivateLinkService) ToPrivateLinkServiceOutputWithContext(ctx context.Context) PrivateLinkServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceOutput)
}

type PrivateLinkServiceOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkService)(nil)).Elem()
}

func (o PrivateLinkServiceOutput) ToPrivateLinkServiceOutput() PrivateLinkServiceOutput {
	return o
}

func (o PrivateLinkServiceOutput) ToPrivateLinkServiceOutputWithContext(ctx context.Context) PrivateLinkServiceOutput {
	return o
}

func (o PrivateLinkServiceOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkService) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceOutput) AutoApproval() PrivateLinkServicePropertiesResponseAutoApprovalPtrOutput {
	return o.ApplyT(func(v *PrivateLinkService) PrivateLinkServicePropertiesResponseAutoApprovalPtrOutput {
		return v.AutoApproval
	}).(PrivateLinkServicePropertiesResponseAutoApprovalPtrOutput)
}

func (o PrivateLinkServiceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkService) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceOutput) Fqdns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateLinkService) pulumi.StringArrayOutput { return v.Fqdns }).(pulumi.StringArrayOutput)
}

func (o PrivateLinkServiceOutput) IpConfigurations() PrivateLinkServiceIpConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *PrivateLinkService) PrivateLinkServiceIpConfigurationResponseArrayOutput {
		return v.IpConfigurations
	}).(PrivateLinkServiceIpConfigurationResponseArrayOutput)
}

func (o PrivateLinkServiceOutput) LoadBalancerFrontendIpConfigurations() FrontendIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *PrivateLinkService) FrontendIPConfigurationResponseArrayOutput {
		return v.LoadBalancerFrontendIpConfigurations
	}).(FrontendIPConfigurationResponseArrayOutput)
}

func (o PrivateLinkServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkService) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v *PrivateLinkService) NetworkInterfaceResponseArrayOutput { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

func (o PrivateLinkServiceOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *PrivateLinkService) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o PrivateLinkServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateLinkService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PrivateLinkServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceOutput) Visibility() PrivateLinkServicePropertiesResponseVisibilityPtrOutput {
	return o.ApplyT(func(v *PrivateLinkService) PrivateLinkServicePropertiesResponseVisibilityPtrOutput {
		return v.Visibility
	}).(PrivateLinkServicePropertiesResponseVisibilityPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkServiceOutput{})
}

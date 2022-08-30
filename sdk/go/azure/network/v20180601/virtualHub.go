


package v20180601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualHub struct {
	pulumi.CustomResourceState

	AddressPrefix                pulumi.StringPtrOutput                         `pulumi:"addressPrefix"`
	Etag                         pulumi.StringOutput                            `pulumi:"etag"`
	HubVirtualNetworkConnections HubVirtualNetworkConnectionResponseArrayOutput `pulumi:"hubVirtualNetworkConnections"`
	Location                     pulumi.StringOutput                            `pulumi:"location"`
	Name                         pulumi.StringOutput                            `pulumi:"name"`
	ProvisioningState            pulumi.StringOutput                            `pulumi:"provisioningState"`
	Tags                         pulumi.StringMapOutput                         `pulumi:"tags"`
	Type                         pulumi.StringOutput                            `pulumi:"type"`
	VirtualWan                   SubResourceResponsePtrOutput                   `pulumi:"virtualWan"`
}


func NewVirtualHub(ctx *pulumi.Context,
	name string, args *VirtualHubArgs, opts ...pulumi.ResourceOption) (*VirtualHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualHub"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualHub
	err := ctx.RegisterResource("azure-native:network/v20180601:VirtualHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualHubState, opts ...pulumi.ResourceOption) (*VirtualHub, error) {
	var resource VirtualHub
	err := ctx.ReadResource("azure-native:network/v20180601:VirtualHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualHubState struct {
}

type VirtualHubState struct {
}

func (VirtualHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubState)(nil)).Elem()
}

type virtualHubArgs struct {
	AddressPrefix                *string                       `pulumi:"addressPrefix"`
	HubVirtualNetworkConnections []HubVirtualNetworkConnection `pulumi:"hubVirtualNetworkConnections"`
	Id                           *string                       `pulumi:"id"`
	Location                     *string                       `pulumi:"location"`
	ResourceGroupName            string                        `pulumi:"resourceGroupName"`
	Tags                         map[string]string             `pulumi:"tags"`
	VirtualHubName               *string                       `pulumi:"virtualHubName"`
	VirtualWan                   *SubResource                  `pulumi:"virtualWan"`
}


type VirtualHubArgs struct {
	AddressPrefix                pulumi.StringPtrInput
	HubVirtualNetworkConnections HubVirtualNetworkConnectionArrayInput
	Id                           pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
	Tags                         pulumi.StringMapInput
	VirtualHubName               pulumi.StringPtrInput
	VirtualWan                   SubResourcePtrInput
}

func (VirtualHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubArgs)(nil)).Elem()
}

type VirtualHubInput interface {
	pulumi.Input

	ToVirtualHubOutput() VirtualHubOutput
	ToVirtualHubOutputWithContext(ctx context.Context) VirtualHubOutput
}

func (*VirtualHub) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHub)(nil)).Elem()
}

func (i *VirtualHub) ToVirtualHubOutput() VirtualHubOutput {
	return i.ToVirtualHubOutputWithContext(context.Background())
}

func (i *VirtualHub) ToVirtualHubOutputWithContext(ctx context.Context) VirtualHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubOutput)
}

type VirtualHubOutput struct{ *pulumi.OutputState }

func (VirtualHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHub)(nil)).Elem()
}

func (o VirtualHubOutput) ToVirtualHubOutput() VirtualHubOutput {
	return o
}

func (o VirtualHubOutput) ToVirtualHubOutputWithContext(ctx context.Context) VirtualHubOutput {
	return o
}

func (o VirtualHubOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringPtrOutput { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o VirtualHubOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o VirtualHubOutput) HubVirtualNetworkConnections() HubVirtualNetworkConnectionResponseArrayOutput {
	return o.ApplyT(func(v *VirtualHub) HubVirtualNetworkConnectionResponseArrayOutput {
		return v.HubVirtualNetworkConnections
	}).(HubVirtualNetworkConnectionResponseArrayOutput)
}

func (o VirtualHubOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualHubOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualHubOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualHubOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualHubOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualHubOutput) VirtualWan() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualHub) SubResourceResponsePtrOutput { return v.VirtualWan }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualHubOutput{})
}

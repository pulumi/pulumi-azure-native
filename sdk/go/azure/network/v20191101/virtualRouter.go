


package v20191101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualRouter struct {
	pulumi.CustomResourceState

	Etag              pulumi.StringOutput            `pulumi:"etag"`
	HostedGateway     SubResourceResponsePtrOutput   `pulumi:"hostedGateway"`
	HostedSubnet      SubResourceResponsePtrOutput   `pulumi:"hostedSubnet"`
	Location          pulumi.StringPtrOutput         `pulumi:"location"`
	Name              pulumi.StringOutput            `pulumi:"name"`
	Peerings          SubResourceResponseArrayOutput `pulumi:"peerings"`
	ProvisioningState pulumi.StringOutput            `pulumi:"provisioningState"`
	Tags              pulumi.StringMapOutput         `pulumi:"tags"`
	Type              pulumi.StringOutput            `pulumi:"type"`
	VirtualRouterAsn  pulumi.Float64PtrOutput        `pulumi:"virtualRouterAsn"`
	VirtualRouterIps  pulumi.StringArrayOutput       `pulumi:"virtualRouterIps"`
}


func NewVirtualRouter(ctx *pulumi.Context,
	name string, args *VirtualRouterArgs, opts ...pulumi.ResourceOption) (*VirtualRouter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:VirtualRouter"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualRouter
	err := ctx.RegisterResource("azure-native:network/v20191101:VirtualRouter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualRouter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualRouterState, opts ...pulumi.ResourceOption) (*VirtualRouter, error) {
	var resource VirtualRouter
	err := ctx.ReadResource("azure-native:network/v20191101:VirtualRouter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualRouterState struct {
}

type VirtualRouterState struct {
}

func (VirtualRouterState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualRouterState)(nil)).Elem()
}

type virtualRouterArgs struct {
	HostedGateway     *SubResource      `pulumi:"hostedGateway"`
	HostedSubnet      *SubResource      `pulumi:"hostedSubnet"`
	Id                *string           `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	VirtualRouterAsn  *float64          `pulumi:"virtualRouterAsn"`
	VirtualRouterIps  []string          `pulumi:"virtualRouterIps"`
	VirtualRouterName *string           `pulumi:"virtualRouterName"`
}


type VirtualRouterArgs struct {
	HostedGateway     SubResourcePtrInput
	HostedSubnet      SubResourcePtrInput
	Id                pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	VirtualRouterAsn  pulumi.Float64PtrInput
	VirtualRouterIps  pulumi.StringArrayInput
	VirtualRouterName pulumi.StringPtrInput
}

func (VirtualRouterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualRouterArgs)(nil)).Elem()
}

type VirtualRouterInput interface {
	pulumi.Input

	ToVirtualRouterOutput() VirtualRouterOutput
	ToVirtualRouterOutputWithContext(ctx context.Context) VirtualRouterOutput
}

func (*VirtualRouter) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualRouter)(nil)).Elem()
}

func (i *VirtualRouter) ToVirtualRouterOutput() VirtualRouterOutput {
	return i.ToVirtualRouterOutputWithContext(context.Background())
}

func (i *VirtualRouter) ToVirtualRouterOutputWithContext(ctx context.Context) VirtualRouterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualRouterOutput)
}

type VirtualRouterOutput struct{ *pulumi.OutputState }

func (VirtualRouterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualRouter)(nil)).Elem()
}

func (o VirtualRouterOutput) ToVirtualRouterOutput() VirtualRouterOutput {
	return o
}

func (o VirtualRouterOutput) ToVirtualRouterOutputWithContext(ctx context.Context) VirtualRouterOutput {
	return o
}

func (o VirtualRouterOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualRouter) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o VirtualRouterOutput) HostedGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualRouter) SubResourceResponsePtrOutput { return v.HostedGateway }).(SubResourceResponsePtrOutput)
}

func (o VirtualRouterOutput) HostedSubnet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualRouter) SubResourceResponsePtrOutput { return v.HostedSubnet }).(SubResourceResponsePtrOutput)
}

func (o VirtualRouterOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRouter) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualRouterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualRouter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualRouterOutput) Peerings() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *VirtualRouter) SubResourceResponseArrayOutput { return v.Peerings }).(SubResourceResponseArrayOutput)
}

func (o VirtualRouterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualRouter) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualRouterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualRouter) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualRouterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualRouter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualRouterOutput) VirtualRouterAsn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VirtualRouter) pulumi.Float64PtrOutput { return v.VirtualRouterAsn }).(pulumi.Float64PtrOutput)
}

func (o VirtualRouterOutput) VirtualRouterIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualRouter) pulumi.StringArrayOutput { return v.VirtualRouterIps }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualRouterOutput{})
}

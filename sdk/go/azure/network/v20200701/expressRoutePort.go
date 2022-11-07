


package v20200701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExpressRoutePort struct {
	pulumi.CustomResourceState

	AllocationDate             pulumi.StringOutput                     `pulumi:"allocationDate"`
	BandwidthInGbps            pulumi.IntPtrOutput                     `pulumi:"bandwidthInGbps"`
	Circuits                   SubResourceResponseArrayOutput          `pulumi:"circuits"`
	Encapsulation              pulumi.StringPtrOutput                  `pulumi:"encapsulation"`
	Etag                       pulumi.StringOutput                     `pulumi:"etag"`
	EtherType                  pulumi.StringOutput                     `pulumi:"etherType"`
	Identity                   ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	Links                      ExpressRouteLinkResponseArrayOutput     `pulumi:"links"`
	Location                   pulumi.StringPtrOutput                  `pulumi:"location"`
	Mtu                        pulumi.StringOutput                     `pulumi:"mtu"`
	Name                       pulumi.StringOutput                     `pulumi:"name"`
	PeeringLocation            pulumi.StringPtrOutput                  `pulumi:"peeringLocation"`
	ProvisionedBandwidthInGbps pulumi.Float64Output                    `pulumi:"provisionedBandwidthInGbps"`
	ProvisioningState          pulumi.StringOutput                     `pulumi:"provisioningState"`
	ResourceGuid               pulumi.StringOutput                     `pulumi:"resourceGuid"`
	Tags                       pulumi.StringMapOutput                  `pulumi:"tags"`
	Type                       pulumi.StringOutput                     `pulumi:"type"`
}


func NewExpressRoutePort(ctx *pulumi.Context,
	name string, args *ExpressRoutePortArgs, opts ...pulumi.ResourceOption) (*ExpressRoutePort, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:ExpressRoutePort"),
		},
	})
	opts = append(opts, aliases)
	var resource ExpressRoutePort
	err := ctx.RegisterResource("azure-native:network/v20200701:ExpressRoutePort", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExpressRoutePort(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRoutePortState, opts ...pulumi.ResourceOption) (*ExpressRoutePort, error) {
	var resource ExpressRoutePort
	err := ctx.ReadResource("azure-native:network/v20200701:ExpressRoutePort", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type expressRoutePortState struct {
}

type ExpressRoutePortState struct {
}

func (ExpressRoutePortState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRoutePortState)(nil)).Elem()
}

type expressRoutePortArgs struct {
	BandwidthInGbps      *int                    `pulumi:"bandwidthInGbps"`
	Encapsulation        *string                 `pulumi:"encapsulation"`
	ExpressRoutePortName *string                 `pulumi:"expressRoutePortName"`
	Id                   *string                 `pulumi:"id"`
	Identity             *ManagedServiceIdentity `pulumi:"identity"`
	Links                []ExpressRouteLink      `pulumi:"links"`
	Location             *string                 `pulumi:"location"`
	PeeringLocation      *string                 `pulumi:"peeringLocation"`
	ResourceGroupName    string                  `pulumi:"resourceGroupName"`
	Tags                 map[string]string       `pulumi:"tags"`
}


type ExpressRoutePortArgs struct {
	BandwidthInGbps      pulumi.IntPtrInput
	Encapsulation        pulumi.StringPtrInput
	ExpressRoutePortName pulumi.StringPtrInput
	Id                   pulumi.StringPtrInput
	Identity             ManagedServiceIdentityPtrInput
	Links                ExpressRouteLinkArrayInput
	Location             pulumi.StringPtrInput
	PeeringLocation      pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	Tags                 pulumi.StringMapInput
}

func (ExpressRoutePortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRoutePortArgs)(nil)).Elem()
}

type ExpressRoutePortInput interface {
	pulumi.Input

	ToExpressRoutePortOutput() ExpressRoutePortOutput
	ToExpressRoutePortOutputWithContext(ctx context.Context) ExpressRoutePortOutput
}

func (*ExpressRoutePort) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRoutePort)(nil)).Elem()
}

func (i *ExpressRoutePort) ToExpressRoutePortOutput() ExpressRoutePortOutput {
	return i.ToExpressRoutePortOutputWithContext(context.Background())
}

func (i *ExpressRoutePort) ToExpressRoutePortOutputWithContext(ctx context.Context) ExpressRoutePortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRoutePortOutput)
}

type ExpressRoutePortOutput struct{ *pulumi.OutputState }

func (ExpressRoutePortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRoutePort)(nil)).Elem()
}

func (o ExpressRoutePortOutput) ToExpressRoutePortOutput() ExpressRoutePortOutput {
	return o
}

func (o ExpressRoutePortOutput) ToExpressRoutePortOutputWithContext(ctx context.Context) ExpressRoutePortOutput {
	return o
}

func (o ExpressRoutePortOutput) AllocationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringOutput { return v.AllocationDate }).(pulumi.StringOutput)
}

func (o ExpressRoutePortOutput) BandwidthInGbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.IntPtrOutput { return v.BandwidthInGbps }).(pulumi.IntPtrOutput)
}

func (o ExpressRoutePortOutput) Circuits() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *ExpressRoutePort) SubResourceResponseArrayOutput { return v.Circuits }).(SubResourceResponseArrayOutput)
}

func (o ExpressRoutePortOutput) Encapsulation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringPtrOutput { return v.Encapsulation }).(pulumi.StringPtrOutput)
}

func (o ExpressRoutePortOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ExpressRoutePortOutput) EtherType() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringOutput { return v.EtherType }).(pulumi.StringOutput)
}

func (o ExpressRoutePortOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ExpressRoutePort) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o ExpressRoutePortOutput) Links() ExpressRouteLinkResponseArrayOutput {
	return o.ApplyT(func(v *ExpressRoutePort) ExpressRouteLinkResponseArrayOutput { return v.Links }).(ExpressRouteLinkResponseArrayOutput)
}

func (o ExpressRoutePortOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ExpressRoutePortOutput) Mtu() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringOutput { return v.Mtu }).(pulumi.StringOutput)
}

func (o ExpressRoutePortOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ExpressRoutePortOutput) PeeringLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringPtrOutput { return v.PeeringLocation }).(pulumi.StringPtrOutput)
}

func (o ExpressRoutePortOutput) ProvisionedBandwidthInGbps() pulumi.Float64Output {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.Float64Output { return v.ProvisionedBandwidthInGbps }).(pulumi.Float64Output)
}

func (o ExpressRoutePortOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ExpressRoutePortOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o ExpressRoutePortOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ExpressRoutePortOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ExpressRoutePortOutput{})
}

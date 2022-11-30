


package v20200601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateZone struct {
	pulumi.CustomResourceState

	Etag                                           pulumi.StringPtrOutput `pulumi:"etag"`
	InternalId                                     pulumi.StringOutput    `pulumi:"internalId"`
	Location                                       pulumi.StringPtrOutput `pulumi:"location"`
	MaxNumberOfRecordSets                          pulumi.Float64Output   `pulumi:"maxNumberOfRecordSets"`
	MaxNumberOfVirtualNetworkLinks                 pulumi.Float64Output   `pulumi:"maxNumberOfVirtualNetworkLinks"`
	MaxNumberOfVirtualNetworkLinksWithRegistration pulumi.Float64Output   `pulumi:"maxNumberOfVirtualNetworkLinksWithRegistration"`
	Name                                           pulumi.StringOutput    `pulumi:"name"`
	NumberOfRecordSets                             pulumi.Float64Output   `pulumi:"numberOfRecordSets"`
	NumberOfVirtualNetworkLinks                    pulumi.Float64Output   `pulumi:"numberOfVirtualNetworkLinks"`
	NumberOfVirtualNetworkLinksWithRegistration    pulumi.Float64Output   `pulumi:"numberOfVirtualNetworkLinksWithRegistration"`
	ProvisioningState                              pulumi.StringOutput    `pulumi:"provisioningState"`
	Tags                                           pulumi.StringMapOutput `pulumi:"tags"`
	Type                                           pulumi.StringOutput    `pulumi:"type"`
}


func NewPrivateZone(ctx *pulumi.Context,
	name string, args *PrivateZoneArgs, opts ...pulumi.ResourceOption) (*PrivateZone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:PrivateZone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180901:PrivateZone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200101:PrivateZone"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateZone
	err := ctx.RegisterResource("azure-native:network/v20200601:PrivateZone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateZoneState, opts ...pulumi.ResourceOption) (*PrivateZone, error) {
	var resource PrivateZone
	err := ctx.ReadResource("azure-native:network/v20200601:PrivateZone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateZoneState struct {
}

type PrivateZoneState struct {
}

func (PrivateZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateZoneState)(nil)).Elem()
}

type privateZoneArgs struct {
	Location          *string           `pulumi:"location"`
	PrivateZoneName   *string           `pulumi:"privateZoneName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type PrivateZoneArgs struct {
	Location          pulumi.StringPtrInput
	PrivateZoneName   pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (PrivateZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateZoneArgs)(nil)).Elem()
}

type PrivateZoneInput interface {
	pulumi.Input

	ToPrivateZoneOutput() PrivateZoneOutput
	ToPrivateZoneOutputWithContext(ctx context.Context) PrivateZoneOutput
}

func (*PrivateZone) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateZone)(nil)).Elem()
}

func (i *PrivateZone) ToPrivateZoneOutput() PrivateZoneOutput {
	return i.ToPrivateZoneOutputWithContext(context.Background())
}

func (i *PrivateZone) ToPrivateZoneOutputWithContext(ctx context.Context) PrivateZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateZoneOutput)
}

type PrivateZoneOutput struct{ *pulumi.OutputState }

func (PrivateZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateZone)(nil)).Elem()
}

func (o PrivateZoneOutput) ToPrivateZoneOutput() PrivateZoneOutput {
	return o
}

func (o PrivateZoneOutput) ToPrivateZoneOutputWithContext(ctx context.Context) PrivateZoneOutput {
	return o
}

func (o PrivateZoneOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateZone) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o PrivateZoneOutput) InternalId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateZone) pulumi.StringOutput { return v.InternalId }).(pulumi.StringOutput)
}

func (o PrivateZoneOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateZone) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o PrivateZoneOutput) MaxNumberOfRecordSets() pulumi.Float64Output {
	return o.ApplyT(func(v *PrivateZone) pulumi.Float64Output { return v.MaxNumberOfRecordSets }).(pulumi.Float64Output)
}

func (o PrivateZoneOutput) MaxNumberOfVirtualNetworkLinks() pulumi.Float64Output {
	return o.ApplyT(func(v *PrivateZone) pulumi.Float64Output { return v.MaxNumberOfVirtualNetworkLinks }).(pulumi.Float64Output)
}

func (o PrivateZoneOutput) MaxNumberOfVirtualNetworkLinksWithRegistration() pulumi.Float64Output {
	return o.ApplyT(func(v *PrivateZone) pulumi.Float64Output { return v.MaxNumberOfVirtualNetworkLinksWithRegistration }).(pulumi.Float64Output)
}

func (o PrivateZoneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateZone) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateZoneOutput) NumberOfRecordSets() pulumi.Float64Output {
	return o.ApplyT(func(v *PrivateZone) pulumi.Float64Output { return v.NumberOfRecordSets }).(pulumi.Float64Output)
}

func (o PrivateZoneOutput) NumberOfVirtualNetworkLinks() pulumi.Float64Output {
	return o.ApplyT(func(v *PrivateZone) pulumi.Float64Output { return v.NumberOfVirtualNetworkLinks }).(pulumi.Float64Output)
}

func (o PrivateZoneOutput) NumberOfVirtualNetworkLinksWithRegistration() pulumi.Float64Output {
	return o.ApplyT(func(v *PrivateZone) pulumi.Float64Output { return v.NumberOfVirtualNetworkLinksWithRegistration }).(pulumi.Float64Output)
}

func (o PrivateZoneOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateZone) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateZoneOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateZone) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PrivateZoneOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateZone) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateZoneOutput{})
}

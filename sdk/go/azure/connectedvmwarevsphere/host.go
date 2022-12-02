


package connectedvmwarevsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Host struct {
	pulumi.CustomResourceState

	CustomResourceName pulumi.StringOutput               `pulumi:"customResourceName"`
	ExtendedLocation   ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	InventoryItemId    pulumi.StringPtrOutput            `pulumi:"inventoryItemId"`
	Kind               pulumi.StringPtrOutput            `pulumi:"kind"`
	Location           pulumi.StringOutput               `pulumi:"location"`
	MoName             pulumi.StringOutput               `pulumi:"moName"`
	MoRefId            pulumi.StringPtrOutput            `pulumi:"moRefId"`
	Name               pulumi.StringOutput               `pulumi:"name"`
	ProvisioningState  pulumi.StringOutput               `pulumi:"provisioningState"`
	Statuses           ResourceStatusResponseArrayOutput `pulumi:"statuses"`
	SystemData         SystemDataResponseOutput          `pulumi:"systemData"`
	Tags               pulumi.StringMapOutput            `pulumi:"tags"`
	Type               pulumi.StringOutput               `pulumi:"type"`
	Uuid               pulumi.StringOutput               `pulumi:"uuid"`
	VCenterId          pulumi.StringPtrOutput            `pulumi:"vCenterId"`
}


func NewHost(ctx *pulumi.Context,
	name string, args *HostArgs, opts ...pulumi.ResourceOption) (*Host, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20201001preview:Host"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220110preview:Host"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220715preview:Host"),
		},
	})
	opts = append(opts, aliases)
	var resource Host
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere:Host", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostState, opts ...pulumi.ResourceOption) (*Host, error) {
	var resource Host
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere:Host", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type hostState struct {
}

type HostState struct {
}

func (HostState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostState)(nil)).Elem()
}

type hostArgs struct {
	ExtendedLocation  *ExtendedLocation `pulumi:"extendedLocation"`
	HostName          *string           `pulumi:"hostName"`
	InventoryItemId   *string           `pulumi:"inventoryItemId"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	MoRefId           *string           `pulumi:"moRefId"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	VCenterId         *string           `pulumi:"vCenterId"`
}


type HostArgs struct {
	ExtendedLocation  ExtendedLocationPtrInput
	HostName          pulumi.StringPtrInput
	InventoryItemId   pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	MoRefId           pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	VCenterId         pulumi.StringPtrInput
}

func (HostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostArgs)(nil)).Elem()
}

type HostInput interface {
	pulumi.Input

	ToHostOutput() HostOutput
	ToHostOutputWithContext(ctx context.Context) HostOutput
}

func (*Host) ElementType() reflect.Type {
	return reflect.TypeOf((**Host)(nil)).Elem()
}

func (i *Host) ToHostOutput() HostOutput {
	return i.ToHostOutputWithContext(context.Background())
}

func (i *Host) ToHostOutputWithContext(ctx context.Context) HostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostOutput)
}

type HostOutput struct{ *pulumi.OutputState }

func (HostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Host)(nil)).Elem()
}

func (o HostOutput) ToHostOutput() HostOutput {
	return o
}

func (o HostOutput) ToHostOutputWithContext(ctx context.Context) HostOutput {
	return o
}

func (o HostOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.CustomResourceName }).(pulumi.StringOutput)
}

func (o HostOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *Host) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o HostOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.StringPtrOutput { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

func (o HostOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o HostOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o HostOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.MoName }).(pulumi.StringOutput)
}

func (o HostOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.StringPtrOutput { return v.MoRefId }).(pulumi.StringPtrOutput)
}

func (o HostOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o HostOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o HostOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v *Host) ResourceStatusResponseArrayOutput { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

func (o HostOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Host) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o HostOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Host) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o HostOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o HostOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

func (o HostOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.StringPtrOutput { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(HostOutput{})
}

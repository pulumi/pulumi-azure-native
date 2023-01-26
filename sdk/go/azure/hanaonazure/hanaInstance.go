


package hanaonazure

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HanaInstance struct {
	pulumi.CustomResourceState

	HanaInstanceId          pulumi.StringOutput              `pulumi:"hanaInstanceId"`
	HardwareProfile         HardwareProfileResponsePtrOutput `pulumi:"hardwareProfile"`
	HwRevision              pulumi.StringOutput              `pulumi:"hwRevision"`
	Location                pulumi.StringPtrOutput           `pulumi:"location"`
	Name                    pulumi.StringOutput              `pulumi:"name"`
	NetworkProfile          NetworkProfileResponsePtrOutput  `pulumi:"networkProfile"`
	OsProfile               OSProfileResponsePtrOutput       `pulumi:"osProfile"`
	PartnerNodeId           pulumi.StringPtrOutput           `pulumi:"partnerNodeId"`
	PowerState              pulumi.StringOutput              `pulumi:"powerState"`
	ProvisioningState       pulumi.StringOutput              `pulumi:"provisioningState"`
	ProximityPlacementGroup pulumi.StringOutput              `pulumi:"proximityPlacementGroup"`
	StorageProfile          StorageProfileResponsePtrOutput  `pulumi:"storageProfile"`
	Tags                    pulumi.StringMapOutput           `pulumi:"tags"`
	Type                    pulumi.StringOutput              `pulumi:"type"`
}


func NewHanaInstance(ctx *pulumi.Context,
	name string, args *HanaInstanceArgs, opts ...pulumi.ResourceOption) (*HanaInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hanaonazure/v20171103preview:HanaInstance"),
		},
	})
	opts = append(opts, aliases)
	var resource HanaInstance
	err := ctx.RegisterResource("azure-native:hanaonazure:HanaInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHanaInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HanaInstanceState, opts ...pulumi.ResourceOption) (*HanaInstance, error) {
	var resource HanaInstance
	err := ctx.ReadResource("azure-native:hanaonazure:HanaInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type hanaInstanceState struct {
}

type HanaInstanceState struct {
}

func (HanaInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*hanaInstanceState)(nil)).Elem()
}

type hanaInstanceArgs struct {
	HanaInstanceName  *string           `pulumi:"hanaInstanceName"`
	Location          *string           `pulumi:"location"`
	NetworkProfile    *NetworkProfile   `pulumi:"networkProfile"`
	OsProfile         *OSProfile        `pulumi:"osProfile"`
	PartnerNodeId     *string           `pulumi:"partnerNodeId"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	StorageProfile    *StorageProfile   `pulumi:"storageProfile"`
	Tags              map[string]string `pulumi:"tags"`
}


type HanaInstanceArgs struct {
	HanaInstanceName  pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	NetworkProfile    NetworkProfilePtrInput
	OsProfile         OSProfilePtrInput
	PartnerNodeId     pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	StorageProfile    StorageProfilePtrInput
	Tags              pulumi.StringMapInput
}

func (HanaInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hanaInstanceArgs)(nil)).Elem()
}

type HanaInstanceInput interface {
	pulumi.Input

	ToHanaInstanceOutput() HanaInstanceOutput
	ToHanaInstanceOutputWithContext(ctx context.Context) HanaInstanceOutput
}

func (*HanaInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**HanaInstance)(nil)).Elem()
}

func (i *HanaInstance) ToHanaInstanceOutput() HanaInstanceOutput {
	return i.ToHanaInstanceOutputWithContext(context.Background())
}

func (i *HanaInstance) ToHanaInstanceOutputWithContext(ctx context.Context) HanaInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HanaInstanceOutput)
}

type HanaInstanceOutput struct{ *pulumi.OutputState }

func (HanaInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HanaInstance)(nil)).Elem()
}

func (o HanaInstanceOutput) ToHanaInstanceOutput() HanaInstanceOutput {
	return o
}

func (o HanaInstanceOutput) ToHanaInstanceOutputWithContext(ctx context.Context) HanaInstanceOutput {
	return o
}

func (o HanaInstanceOutput) HanaInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaInstance) pulumi.StringOutput { return v.HanaInstanceId }).(pulumi.StringOutput)
}

func (o HanaInstanceOutput) HardwareProfile() HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v *HanaInstance) HardwareProfileResponsePtrOutput { return v.HardwareProfile }).(HardwareProfileResponsePtrOutput)
}

func (o HanaInstanceOutput) HwRevision() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaInstance) pulumi.StringOutput { return v.HwRevision }).(pulumi.StringOutput)
}

func (o HanaInstanceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HanaInstance) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o HanaInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o HanaInstanceOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *HanaInstance) NetworkProfileResponsePtrOutput { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o HanaInstanceOutput) OsProfile() OSProfileResponsePtrOutput {
	return o.ApplyT(func(v *HanaInstance) OSProfileResponsePtrOutput { return v.OsProfile }).(OSProfileResponsePtrOutput)
}

func (o HanaInstanceOutput) PartnerNodeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HanaInstance) pulumi.StringPtrOutput { return v.PartnerNodeId }).(pulumi.StringPtrOutput)
}

func (o HanaInstanceOutput) PowerState() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaInstance) pulumi.StringOutput { return v.PowerState }).(pulumi.StringOutput)
}

func (o HanaInstanceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaInstance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o HanaInstanceOutput) ProximityPlacementGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaInstance) pulumi.StringOutput { return v.ProximityPlacementGroup }).(pulumi.StringOutput)
}

func (o HanaInstanceOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v *HanaInstance) StorageProfileResponsePtrOutput { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

func (o HanaInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HanaInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o HanaInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(HanaInstanceOutput{})
}

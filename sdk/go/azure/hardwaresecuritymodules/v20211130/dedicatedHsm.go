


package v20211130

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DedicatedHsm struct {
	pulumi.CustomResourceState

	Location                 pulumi.StringOutput             `pulumi:"location"`
	ManagementNetworkProfile NetworkProfileResponsePtrOutput `pulumi:"managementNetworkProfile"`
	Name                     pulumi.StringOutput             `pulumi:"name"`
	NetworkProfile           NetworkProfileResponsePtrOutput `pulumi:"networkProfile"`
	ProvisioningState        pulumi.StringOutput             `pulumi:"provisioningState"`
	Sku                      SkuResponseOutput               `pulumi:"sku"`
	StampId                  pulumi.StringPtrOutput          `pulumi:"stampId"`
	StatusMessage            pulumi.StringOutput             `pulumi:"statusMessage"`
	SystemData               SystemDataResponseOutput        `pulumi:"systemData"`
	Tags                     pulumi.StringMapOutput          `pulumi:"tags"`
	Type                     pulumi.StringOutput             `pulumi:"type"`
	Zones                    pulumi.StringArrayOutput        `pulumi:"zones"`
}


func NewDedicatedHsm(ctx *pulumi.Context,
	name string, args *DedicatedHsmArgs, opts ...pulumi.ResourceOption) (*DedicatedHsm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hardwaresecuritymodules:DedicatedHsm"),
		},
		{
			Type: pulumi.String("azure-native:hardwaresecuritymodules/v20181031preview:DedicatedHsm"),
		},
	})
	opts = append(opts, aliases)
	var resource DedicatedHsm
	err := ctx.RegisterResource("azure-native:hardwaresecuritymodules/v20211130:DedicatedHsm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDedicatedHsm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedHsmState, opts ...pulumi.ResourceOption) (*DedicatedHsm, error) {
	var resource DedicatedHsm
	err := ctx.ReadResource("azure-native:hardwaresecuritymodules/v20211130:DedicatedHsm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dedicatedHsmState struct {
}

type DedicatedHsmState struct {
}

func (DedicatedHsmState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHsmState)(nil)).Elem()
}

type dedicatedHsmArgs struct {
	Location                 *string           `pulumi:"location"`
	ManagementNetworkProfile *NetworkProfile   `pulumi:"managementNetworkProfile"`
	Name                     *string           `pulumi:"name"`
	NetworkProfile           *NetworkProfile   `pulumi:"networkProfile"`
	ResourceGroupName        string            `pulumi:"resourceGroupName"`
	Sku                      Sku               `pulumi:"sku"`
	StampId                  *string           `pulumi:"stampId"`
	Tags                     map[string]string `pulumi:"tags"`
	Zones                    []string          `pulumi:"zones"`
}


type DedicatedHsmArgs struct {
	Location                 pulumi.StringPtrInput
	ManagementNetworkProfile NetworkProfilePtrInput
	Name                     pulumi.StringPtrInput
	NetworkProfile           NetworkProfilePtrInput
	ResourceGroupName        pulumi.StringInput
	Sku                      SkuInput
	StampId                  pulumi.StringPtrInput
	Tags                     pulumi.StringMapInput
	Zones                    pulumi.StringArrayInput
}

func (DedicatedHsmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHsmArgs)(nil)).Elem()
}

type DedicatedHsmInput interface {
	pulumi.Input

	ToDedicatedHsmOutput() DedicatedHsmOutput
	ToDedicatedHsmOutputWithContext(ctx context.Context) DedicatedHsmOutput
}

func (*DedicatedHsm) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHsm)(nil)).Elem()
}

func (i *DedicatedHsm) ToDedicatedHsmOutput() DedicatedHsmOutput {
	return i.ToDedicatedHsmOutputWithContext(context.Background())
}

func (i *DedicatedHsm) ToDedicatedHsmOutputWithContext(ctx context.Context) DedicatedHsmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHsmOutput)
}

type DedicatedHsmOutput struct{ *pulumi.OutputState }

func (DedicatedHsmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHsm)(nil)).Elem()
}

func (o DedicatedHsmOutput) ToDedicatedHsmOutput() DedicatedHsmOutput {
	return o
}

func (o DedicatedHsmOutput) ToDedicatedHsmOutputWithContext(ctx context.Context) DedicatedHsmOutput {
	return o
}

func (o DedicatedHsmOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHsm) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DedicatedHsmOutput) ManagementNetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *DedicatedHsm) NetworkProfileResponsePtrOutput { return v.ManagementNetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o DedicatedHsmOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHsm) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DedicatedHsmOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *DedicatedHsm) NetworkProfileResponsePtrOutput { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o DedicatedHsmOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHsm) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DedicatedHsmOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *DedicatedHsm) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

func (o DedicatedHsmOutput) StampId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DedicatedHsm) pulumi.StringPtrOutput { return v.StampId }).(pulumi.StringPtrOutput)
}

func (o DedicatedHsmOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHsm) pulumi.StringOutput { return v.StatusMessage }).(pulumi.StringOutput)
}

func (o DedicatedHsmOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DedicatedHsm) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DedicatedHsmOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DedicatedHsm) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DedicatedHsmOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHsm) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o DedicatedHsmOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DedicatedHsm) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(DedicatedHsmOutput{})
}




package v20210401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DedicatedHost struct {
	pulumi.CustomResourceState

	AutoReplaceOnFailure pulumi.BoolPtrOutput                    `pulumi:"autoReplaceOnFailure"`
	HostId               pulumi.StringOutput                     `pulumi:"hostId"`
	InstanceView         DedicatedHostInstanceViewResponseOutput `pulumi:"instanceView"`
	LicenseType          pulumi.StringPtrOutput                  `pulumi:"licenseType"`
	Location             pulumi.StringOutput                     `pulumi:"location"`
	Name                 pulumi.StringOutput                     `pulumi:"name"`
	PlatformFaultDomain  pulumi.IntPtrOutput                     `pulumi:"platformFaultDomain"`
	ProvisioningState    pulumi.StringOutput                     `pulumi:"provisioningState"`
	ProvisioningTime     pulumi.StringOutput                     `pulumi:"provisioningTime"`
	Sku                  SkuResponseOutput                       `pulumi:"sku"`
	Tags                 pulumi.StringMapOutput                  `pulumi:"tags"`
	Type                 pulumi.StringOutput                     `pulumi:"type"`
	VirtualMachines      SubResourceReadOnlyResponseArrayOutput  `pulumi:"virtualMachines"`
}


func NewDedicatedHost(ctx *pulumi.Context,
	name string, args *DedicatedHostArgs, opts ...pulumi.ResourceOption) (*DedicatedHost, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostGroupName == nil {
		return nil, errors.New("invalid value for required argument 'HostGroupName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:DedicatedHost"),
		},
	})
	opts = append(opts, aliases)
	var resource DedicatedHost
	err := ctx.RegisterResource("azure-native:compute/v20210401:DedicatedHost", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDedicatedHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedHostState, opts ...pulumi.ResourceOption) (*DedicatedHost, error) {
	var resource DedicatedHost
	err := ctx.ReadResource("azure-native:compute/v20210401:DedicatedHost", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dedicatedHostState struct {
}

type DedicatedHostState struct {
}

func (DedicatedHostState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostState)(nil)).Elem()
}

type dedicatedHostArgs struct {
	AutoReplaceOnFailure *bool                      `pulumi:"autoReplaceOnFailure"`
	HostGroupName        string                     `pulumi:"hostGroupName"`
	HostName             *string                    `pulumi:"hostName"`
	LicenseType          *DedicatedHostLicenseTypes `pulumi:"licenseType"`
	Location             *string                    `pulumi:"location"`
	PlatformFaultDomain  *int                       `pulumi:"platformFaultDomain"`
	ResourceGroupName    string                     `pulumi:"resourceGroupName"`
	Sku                  Sku                        `pulumi:"sku"`
	Tags                 map[string]string          `pulumi:"tags"`
}


type DedicatedHostArgs struct {
	AutoReplaceOnFailure pulumi.BoolPtrInput
	HostGroupName        pulumi.StringInput
	HostName             pulumi.StringPtrInput
	LicenseType          DedicatedHostLicenseTypesPtrInput
	Location             pulumi.StringPtrInput
	PlatformFaultDomain  pulumi.IntPtrInput
	ResourceGroupName    pulumi.StringInput
	Sku                  SkuInput
	Tags                 pulumi.StringMapInput
}

func (DedicatedHostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostArgs)(nil)).Elem()
}

type DedicatedHostInput interface {
	pulumi.Input

	ToDedicatedHostOutput() DedicatedHostOutput
	ToDedicatedHostOutputWithContext(ctx context.Context) DedicatedHostOutput
}

func (*DedicatedHost) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHost)(nil)).Elem()
}

func (i *DedicatedHost) ToDedicatedHostOutput() DedicatedHostOutput {
	return i.ToDedicatedHostOutputWithContext(context.Background())
}

func (i *DedicatedHost) ToDedicatedHostOutputWithContext(ctx context.Context) DedicatedHostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostOutput)
}

type DedicatedHostOutput struct{ *pulumi.OutputState }

func (DedicatedHostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHost)(nil)).Elem()
}

func (o DedicatedHostOutput) ToDedicatedHostOutput() DedicatedHostOutput {
	return o
}

func (o DedicatedHostOutput) ToDedicatedHostOutputWithContext(ctx context.Context) DedicatedHostOutput {
	return o
}

func (o DedicatedHostOutput) AutoReplaceOnFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.BoolPtrOutput { return v.AutoReplaceOnFailure }).(pulumi.BoolPtrOutput)
}

func (o DedicatedHostOutput) HostId() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.HostId }).(pulumi.StringOutput)
}

func (o DedicatedHostOutput) InstanceView() DedicatedHostInstanceViewResponseOutput {
	return o.ApplyT(func(v *DedicatedHost) DedicatedHostInstanceViewResponseOutput { return v.InstanceView }).(DedicatedHostInstanceViewResponseOutput)
}

func (o DedicatedHostOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringPtrOutput { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o DedicatedHostOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DedicatedHostOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DedicatedHostOutput) PlatformFaultDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.IntPtrOutput { return v.PlatformFaultDomain }).(pulumi.IntPtrOutput)
}

func (o DedicatedHostOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DedicatedHostOutput) ProvisioningTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.ProvisioningTime }).(pulumi.StringOutput)
}

func (o DedicatedHostOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *DedicatedHost) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

func (o DedicatedHostOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DedicatedHostOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o DedicatedHostOutput) VirtualMachines() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v *DedicatedHost) SubResourceReadOnlyResponseArrayOutput { return v.VirtualMachines }).(SubResourceReadOnlyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(DedicatedHostOutput{})
}

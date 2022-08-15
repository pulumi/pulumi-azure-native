


package v20220415preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Appliance struct {
	pulumi.CustomResourceState

	Distro               pulumi.StringPtrOutput                                   `pulumi:"distro"`
	Identity             IdentityResponsePtrOutput                                `pulumi:"identity"`
	InfrastructureConfig AppliancePropertiesResponseInfrastructureConfigPtrOutput `pulumi:"infrastructureConfig"`
	Location             pulumi.StringOutput                                      `pulumi:"location"`
	Name                 pulumi.StringOutput                                      `pulumi:"name"`
	ProvisioningState    pulumi.StringOutput                                      `pulumi:"provisioningState"`
	PublicKey            pulumi.StringPtrOutput                                   `pulumi:"publicKey"`
	Status               pulumi.StringOutput                                      `pulumi:"status"`
	SystemData           SystemDataResponseOutput                                 `pulumi:"systemData"`
	Tags                 pulumi.StringMapOutput                                   `pulumi:"tags"`
	Type                 pulumi.StringOutput                                      `pulumi:"type"`
	Version              pulumi.StringPtrOutput                                   `pulumi:"version"`
}


func NewAppliance(ctx *pulumi.Context,
	name string, args *ApplianceArgs, opts ...pulumi.ResourceOption) (*Appliance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.Distro) {
		args.Distro = pulumi.StringPtr("AKSEdge")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resourceconnector:Appliance"),
		},
		{
			Type: pulumi.String("azure-native:resourceconnector/v20211031preview:Appliance"),
		},
	})
	opts = append(opts, aliases)
	var resource Appliance
	err := ctx.RegisterResource("azure-native:resourceconnector/v20220415preview:Appliance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAppliance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplianceState, opts ...pulumi.ResourceOption) (*Appliance, error) {
	var resource Appliance
	err := ctx.ReadResource("azure-native:resourceconnector/v20220415preview:Appliance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type applianceState struct {
}

type ApplianceState struct {
}

func (ApplianceState) ElementType() reflect.Type {
	return reflect.TypeOf((*applianceState)(nil)).Elem()
}

type applianceArgs struct {
	Distro               *string                                  `pulumi:"distro"`
	Identity             *Identity                                `pulumi:"identity"`
	InfrastructureConfig *AppliancePropertiesInfrastructureConfig `pulumi:"infrastructureConfig"`
	Location             *string                                  `pulumi:"location"`
	PublicKey            *string                                  `pulumi:"publicKey"`
	ResourceGroupName    string                                   `pulumi:"resourceGroupName"`
	ResourceName         *string                                  `pulumi:"resourceName"`
	Tags                 map[string]string                        `pulumi:"tags"`
	Version              *string                                  `pulumi:"version"`
}


type ApplianceArgs struct {
	Distro               pulumi.StringPtrInput
	Identity             IdentityPtrInput
	InfrastructureConfig AppliancePropertiesInfrastructureConfigPtrInput
	Location             pulumi.StringPtrInput
	PublicKey            pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	ResourceName         pulumi.StringPtrInput
	Tags                 pulumi.StringMapInput
	Version              pulumi.StringPtrInput
}

func (ApplianceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applianceArgs)(nil)).Elem()
}

type ApplianceInput interface {
	pulumi.Input

	ToApplianceOutput() ApplianceOutput
	ToApplianceOutputWithContext(ctx context.Context) ApplianceOutput
}

func (*Appliance) ElementType() reflect.Type {
	return reflect.TypeOf((**Appliance)(nil)).Elem()
}

func (i *Appliance) ToApplianceOutput() ApplianceOutput {
	return i.ToApplianceOutputWithContext(context.Background())
}

func (i *Appliance) ToApplianceOutputWithContext(ctx context.Context) ApplianceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplianceOutput)
}

type ApplianceOutput struct{ *pulumi.OutputState }

func (ApplianceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Appliance)(nil)).Elem()
}

func (o ApplianceOutput) ToApplianceOutput() ApplianceOutput {
	return o
}

func (o ApplianceOutput) ToApplianceOutputWithContext(ctx context.Context) ApplianceOutput {
	return o
}

func (o ApplianceOutput) Distro() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Appliance) pulumi.StringPtrOutput { return v.Distro }).(pulumi.StringPtrOutput)
}

func (o ApplianceOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Appliance) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o ApplianceOutput) InfrastructureConfig() AppliancePropertiesResponseInfrastructureConfigPtrOutput {
	return o.ApplyT(func(v *Appliance) AppliancePropertiesResponseInfrastructureConfigPtrOutput {
		return v.InfrastructureConfig
	}).(AppliancePropertiesResponseInfrastructureConfigPtrOutput)
}

func (o ApplianceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Appliance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ApplianceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Appliance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplianceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Appliance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ApplianceOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Appliance) pulumi.StringPtrOutput { return v.PublicKey }).(pulumi.StringPtrOutput)
}

func (o ApplianceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Appliance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o ApplianceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Appliance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ApplianceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Appliance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ApplianceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Appliance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ApplianceOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Appliance) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplianceOutput{})
}

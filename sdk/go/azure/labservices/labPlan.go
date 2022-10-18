


package labservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LabPlan struct {
	pulumi.CustomResourceState

	AllowedRegions             pulumi.StringArrayOutput               `pulumi:"allowedRegions"`
	DefaultAutoShutdownProfile AutoShutdownProfileResponsePtrOutput   `pulumi:"defaultAutoShutdownProfile"`
	DefaultConnectionProfile   ConnectionProfileResponsePtrOutput     `pulumi:"defaultConnectionProfile"`
	DefaultNetworkProfile      LabPlanNetworkProfileResponsePtrOutput `pulumi:"defaultNetworkProfile"`
	LinkedLmsInstance          pulumi.StringPtrOutput                 `pulumi:"linkedLmsInstance"`
	Location                   pulumi.StringOutput                    `pulumi:"location"`
	Name                       pulumi.StringOutput                    `pulumi:"name"`
	ProvisioningState          pulumi.StringOutput                    `pulumi:"provisioningState"`
	SharedGalleryId            pulumi.StringPtrOutput                 `pulumi:"sharedGalleryId"`
	SupportInfo                SupportInfoResponsePtrOutput           `pulumi:"supportInfo"`
	SystemData                 SystemDataResponseOutput               `pulumi:"systemData"`
	Tags                       pulumi.StringMapOutput                 `pulumi:"tags"`
	Type                       pulumi.StringOutput                    `pulumi:"type"`
}


func NewLabPlan(ctx *pulumi.Context,
	name string, args *LabPlanArgs, opts ...pulumi.ResourceOption) (*LabPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.DefaultAutoShutdownProfile != nil {
		args.DefaultAutoShutdownProfile = args.DefaultAutoShutdownProfile.ToAutoShutdownProfilePtrOutput().ApplyT(func(v *AutoShutdownProfile) *AutoShutdownProfile { return v.Defaults() }).(AutoShutdownProfilePtrOutput)
	}
	if args.DefaultConnectionProfile != nil {
		args.DefaultConnectionProfile = args.DefaultConnectionProfile.ToConnectionProfilePtrOutput().ApplyT(func(v *ConnectionProfile) *ConnectionProfile { return v.Defaults() }).(ConnectionProfilePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:labservices/v20211001preview:LabPlan"),
		},
		{
			Type: pulumi.String("azure-native:labservices/v20211115preview:LabPlan"),
		},
		{
			Type: pulumi.String("azure-native:labservices/v20220801:LabPlan"),
		},
	})
	opts = append(opts, aliases)
	var resource LabPlan
	err := ctx.RegisterResource("azure-native:labservices:LabPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLabPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabPlanState, opts ...pulumi.ResourceOption) (*LabPlan, error) {
	var resource LabPlan
	err := ctx.ReadResource("azure-native:labservices:LabPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type labPlanState struct {
}

type LabPlanState struct {
}

func (LabPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*labPlanState)(nil)).Elem()
}

type labPlanArgs struct {
	AllowedRegions             []string               `pulumi:"allowedRegions"`
	DefaultAutoShutdownProfile *AutoShutdownProfile   `pulumi:"defaultAutoShutdownProfile"`
	DefaultConnectionProfile   *ConnectionProfile     `pulumi:"defaultConnectionProfile"`
	DefaultNetworkProfile      *LabPlanNetworkProfile `pulumi:"defaultNetworkProfile"`
	LabPlanName                *string                `pulumi:"labPlanName"`
	LinkedLmsInstance          *string                `pulumi:"linkedLmsInstance"`
	Location                   *string                `pulumi:"location"`
	ResourceGroupName          string                 `pulumi:"resourceGroupName"`
	SharedGalleryId            *string                `pulumi:"sharedGalleryId"`
	SupportInfo                *SupportInfo           `pulumi:"supportInfo"`
	Tags                       map[string]string      `pulumi:"tags"`
}


type LabPlanArgs struct {
	AllowedRegions             pulumi.StringArrayInput
	DefaultAutoShutdownProfile AutoShutdownProfilePtrInput
	DefaultConnectionProfile   ConnectionProfilePtrInput
	DefaultNetworkProfile      LabPlanNetworkProfilePtrInput
	LabPlanName                pulumi.StringPtrInput
	LinkedLmsInstance          pulumi.StringPtrInput
	Location                   pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	SharedGalleryId            pulumi.StringPtrInput
	SupportInfo                SupportInfoPtrInput
	Tags                       pulumi.StringMapInput
}

func (LabPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labPlanArgs)(nil)).Elem()
}

type LabPlanInput interface {
	pulumi.Input

	ToLabPlanOutput() LabPlanOutput
	ToLabPlanOutputWithContext(ctx context.Context) LabPlanOutput
}

func (*LabPlan) ElementType() reflect.Type {
	return reflect.TypeOf((**LabPlan)(nil)).Elem()
}

func (i *LabPlan) ToLabPlanOutput() LabPlanOutput {
	return i.ToLabPlanOutputWithContext(context.Background())
}

func (i *LabPlan) ToLabPlanOutputWithContext(ctx context.Context) LabPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabPlanOutput)
}

type LabPlanOutput struct{ *pulumi.OutputState }

func (LabPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabPlan)(nil)).Elem()
}

func (o LabPlanOutput) ToLabPlanOutput() LabPlanOutput {
	return o
}

func (o LabPlanOutput) ToLabPlanOutputWithContext(ctx context.Context) LabPlanOutput {
	return o
}

func (o LabPlanOutput) AllowedRegions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LabPlan) pulumi.StringArrayOutput { return v.AllowedRegions }).(pulumi.StringArrayOutput)
}

func (o LabPlanOutput) DefaultAutoShutdownProfile() AutoShutdownProfileResponsePtrOutput {
	return o.ApplyT(func(v *LabPlan) AutoShutdownProfileResponsePtrOutput { return v.DefaultAutoShutdownProfile }).(AutoShutdownProfileResponsePtrOutput)
}

func (o LabPlanOutput) DefaultConnectionProfile() ConnectionProfileResponsePtrOutput {
	return o.ApplyT(func(v *LabPlan) ConnectionProfileResponsePtrOutput { return v.DefaultConnectionProfile }).(ConnectionProfileResponsePtrOutput)
}

func (o LabPlanOutput) DefaultNetworkProfile() LabPlanNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *LabPlan) LabPlanNetworkProfileResponsePtrOutput { return v.DefaultNetworkProfile }).(LabPlanNetworkProfileResponsePtrOutput)
}

func (o LabPlanOutput) LinkedLmsInstance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabPlan) pulumi.StringPtrOutput { return v.LinkedLmsInstance }).(pulumi.StringPtrOutput)
}

func (o LabPlanOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *LabPlan) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o LabPlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LabPlan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LabPlanOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *LabPlan) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LabPlanOutput) SharedGalleryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabPlan) pulumi.StringPtrOutput { return v.SharedGalleryId }).(pulumi.StringPtrOutput)
}

func (o LabPlanOutput) SupportInfo() SupportInfoResponsePtrOutput {
	return o.ApplyT(func(v *LabPlan) SupportInfoResponsePtrOutput { return v.SupportInfo }).(SupportInfoResponsePtrOutput)
}

func (o LabPlanOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *LabPlan) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LabPlanOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LabPlan) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LabPlanOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LabPlan) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LabPlanOutput{})
}




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

func init() {
	pulumi.RegisterOutputType(LabPlanOutput{})
}

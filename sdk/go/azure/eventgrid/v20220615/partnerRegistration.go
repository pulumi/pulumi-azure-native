


package v20220615

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PartnerRegistration struct {
	pulumi.CustomResourceState

	Location                       pulumi.StringOutput      `pulumi:"location"`
	Name                           pulumi.StringOutput      `pulumi:"name"`
	PartnerRegistrationImmutableId pulumi.StringPtrOutput   `pulumi:"partnerRegistrationImmutableId"`
	ProvisioningState              pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData                     SystemDataResponseOutput `pulumi:"systemData"`
	Tags                           pulumi.StringMapOutput   `pulumi:"tags"`
	Type                           pulumi.StringOutput      `pulumi:"type"`
}


func NewPartnerRegistration(ctx *pulumi.Context,
	name string, args *PartnerRegistrationArgs, opts ...pulumi.ResourceOption) (*PartnerRegistration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:PartnerRegistration"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200401preview:PartnerRegistration"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20201015preview:PartnerRegistration"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20210601preview:PartnerRegistration"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:PartnerRegistration"),
		},
	})
	opts = append(opts, aliases)
	var resource PartnerRegistration
	err := ctx.RegisterResource("azure-native:eventgrid/v20220615:PartnerRegistration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPartnerRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartnerRegistrationState, opts ...pulumi.ResourceOption) (*PartnerRegistration, error) {
	var resource PartnerRegistration
	err := ctx.ReadResource("azure-native:eventgrid/v20220615:PartnerRegistration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type partnerRegistrationState struct {
}

type PartnerRegistrationState struct {
}

func (PartnerRegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerRegistrationState)(nil)).Elem()
}

type partnerRegistrationArgs struct {
	Location                       *string           `pulumi:"location"`
	PartnerRegistrationImmutableId *string           `pulumi:"partnerRegistrationImmutableId"`
	PartnerRegistrationName        *string           `pulumi:"partnerRegistrationName"`
	ResourceGroupName              string            `pulumi:"resourceGroupName"`
	Tags                           map[string]string `pulumi:"tags"`
}


type PartnerRegistrationArgs struct {
	Location                       pulumi.StringPtrInput
	PartnerRegistrationImmutableId pulumi.StringPtrInput
	PartnerRegistrationName        pulumi.StringPtrInput
	ResourceGroupName              pulumi.StringInput
	Tags                           pulumi.StringMapInput
}

func (PartnerRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerRegistrationArgs)(nil)).Elem()
}

type PartnerRegistrationInput interface {
	pulumi.Input

	ToPartnerRegistrationOutput() PartnerRegistrationOutput
	ToPartnerRegistrationOutputWithContext(ctx context.Context) PartnerRegistrationOutput
}

func (*PartnerRegistration) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerRegistration)(nil)).Elem()
}

func (i *PartnerRegistration) ToPartnerRegistrationOutput() PartnerRegistrationOutput {
	return i.ToPartnerRegistrationOutputWithContext(context.Background())
}

func (i *PartnerRegistration) ToPartnerRegistrationOutputWithContext(ctx context.Context) PartnerRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerRegistrationOutput)
}

type PartnerRegistrationOutput struct{ *pulumi.OutputState }

func (PartnerRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerRegistration)(nil)).Elem()
}

func (o PartnerRegistrationOutput) ToPartnerRegistrationOutput() PartnerRegistrationOutput {
	return o
}

func (o PartnerRegistrationOutput) ToPartnerRegistrationOutputWithContext(ctx context.Context) PartnerRegistrationOutput {
	return o
}

func (o PartnerRegistrationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PartnerRegistrationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PartnerRegistrationOutput) PartnerRegistrationImmutableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringPtrOutput { return v.PartnerRegistrationImmutableId }).(pulumi.StringPtrOutput)
}

func (o PartnerRegistrationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PartnerRegistrationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PartnerRegistration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PartnerRegistrationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PartnerRegistrationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PartnerRegistrationOutput{})
}

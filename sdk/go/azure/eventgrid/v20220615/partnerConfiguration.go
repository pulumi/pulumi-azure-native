


package v20220615

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PartnerConfiguration struct {
	pulumi.CustomResourceState

	Location             pulumi.StringPtrOutput                `pulumi:"location"`
	Name                 pulumi.StringOutput                   `pulumi:"name"`
	PartnerAuthorization PartnerAuthorizationResponsePtrOutput `pulumi:"partnerAuthorization"`
	ProvisioningState    pulumi.StringPtrOutput                `pulumi:"provisioningState"`
	SystemData           SystemDataResponseOutput              `pulumi:"systemData"`
	Tags                 pulumi.StringMapOutput                `pulumi:"tags"`
	Type                 pulumi.StringOutput                   `pulumi:"type"`
}


func NewPartnerConfiguration(ctx *pulumi.Context,
	name string, args *PartnerConfigurationArgs, opts ...pulumi.ResourceOption) (*PartnerConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:PartnerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:PartnerConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource PartnerConfiguration
	err := ctx.RegisterResource("azure-native:eventgrid/v20220615:PartnerConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPartnerConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartnerConfigurationState, opts ...pulumi.ResourceOption) (*PartnerConfiguration, error) {
	var resource PartnerConfiguration
	err := ctx.ReadResource("azure-native:eventgrid/v20220615:PartnerConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type partnerConfigurationState struct {
}

type PartnerConfigurationState struct {
}

func (PartnerConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerConfigurationState)(nil)).Elem()
}

type partnerConfigurationArgs struct {
	Location             *string               `pulumi:"location"`
	PartnerAuthorization *PartnerAuthorization `pulumi:"partnerAuthorization"`
	ProvisioningState    *string               `pulumi:"provisioningState"`
	ResourceGroupName    string                `pulumi:"resourceGroupName"`
	Tags                 map[string]string     `pulumi:"tags"`
}


type PartnerConfigurationArgs struct {
	Location             pulumi.StringPtrInput
	PartnerAuthorization PartnerAuthorizationPtrInput
	ProvisioningState    pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	Tags                 pulumi.StringMapInput
}

func (PartnerConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerConfigurationArgs)(nil)).Elem()
}

type PartnerConfigurationInput interface {
	pulumi.Input

	ToPartnerConfigurationOutput() PartnerConfigurationOutput
	ToPartnerConfigurationOutputWithContext(ctx context.Context) PartnerConfigurationOutput
}

func (*PartnerConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerConfiguration)(nil)).Elem()
}

func (i *PartnerConfiguration) ToPartnerConfigurationOutput() PartnerConfigurationOutput {
	return i.ToPartnerConfigurationOutputWithContext(context.Background())
}

func (i *PartnerConfiguration) ToPartnerConfigurationOutputWithContext(ctx context.Context) PartnerConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerConfigurationOutput)
}

type PartnerConfigurationOutput struct{ *pulumi.OutputState }

func (PartnerConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerConfiguration)(nil)).Elem()
}

func (o PartnerConfigurationOutput) ToPartnerConfigurationOutput() PartnerConfigurationOutput {
	return o
}

func (o PartnerConfigurationOutput) ToPartnerConfigurationOutputWithContext(ctx context.Context) PartnerConfigurationOutput {
	return o
}

func (o PartnerConfigurationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerConfiguration) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o PartnerConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PartnerConfigurationOutput) PartnerAuthorization() PartnerAuthorizationResponsePtrOutput {
	return o.ApplyT(func(v *PartnerConfiguration) PartnerAuthorizationResponsePtrOutput { return v.PartnerAuthorization }).(PartnerAuthorizationResponsePtrOutput)
}

func (o PartnerConfigurationOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerConfiguration) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o PartnerConfigurationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PartnerConfiguration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PartnerConfigurationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PartnerConfiguration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PartnerConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PartnerConfigurationOutput{})
}




package v20211015preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PartnerDestination struct {
	pulumi.CustomResourceState

	ActivationState                 pulumi.StringPtrOutput   `pulumi:"activationState"`
	EndpointBaseUrl                 pulumi.StringPtrOutput   `pulumi:"endpointBaseUrl"`
	EndpointServiceContext          pulumi.StringPtrOutput   `pulumi:"endpointServiceContext"`
	ExpirationTimeIfNotActivatedUtc pulumi.StringPtrOutput   `pulumi:"expirationTimeIfNotActivatedUtc"`
	Location                        pulumi.StringOutput      `pulumi:"location"`
	MessageForActivation            pulumi.StringPtrOutput   `pulumi:"messageForActivation"`
	Name                            pulumi.StringOutput      `pulumi:"name"`
	PartnerRegistrationImmutableId  pulumi.StringPtrOutput   `pulumi:"partnerRegistrationImmutableId"`
	ProvisioningState               pulumi.StringPtrOutput   `pulumi:"provisioningState"`
	SystemData                      SystemDataResponseOutput `pulumi:"systemData"`
	Tags                            pulumi.StringMapOutput   `pulumi:"tags"`
	Type                            pulumi.StringOutput      `pulumi:"type"`
}


func NewPartnerDestination(ctx *pulumi.Context,
	name string, args *PartnerDestinationArgs, opts ...pulumi.ResourceOption) (*PartnerDestination, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:PartnerDestination"),
		},
	})
	opts = append(opts, aliases)
	var resource PartnerDestination
	err := ctx.RegisterResource("azure-native:eventgrid/v20211015preview:PartnerDestination", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPartnerDestination(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartnerDestinationState, opts ...pulumi.ResourceOption) (*PartnerDestination, error) {
	var resource PartnerDestination
	err := ctx.ReadResource("azure-native:eventgrid/v20211015preview:PartnerDestination", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type partnerDestinationState struct {
}

type PartnerDestinationState struct {
}

func (PartnerDestinationState) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerDestinationState)(nil)).Elem()
}

type partnerDestinationArgs struct {
	ActivationState                 *string           `pulumi:"activationState"`
	EndpointBaseUrl                 *string           `pulumi:"endpointBaseUrl"`
	EndpointServiceContext          *string           `pulumi:"endpointServiceContext"`
	ExpirationTimeIfNotActivatedUtc *string           `pulumi:"expirationTimeIfNotActivatedUtc"`
	Location                        *string           `pulumi:"location"`
	MessageForActivation            *string           `pulumi:"messageForActivation"`
	PartnerDestinationName          *string           `pulumi:"partnerDestinationName"`
	PartnerRegistrationImmutableId  *string           `pulumi:"partnerRegistrationImmutableId"`
	ProvisioningState               *string           `pulumi:"provisioningState"`
	ResourceGroupName               string            `pulumi:"resourceGroupName"`
	Tags                            map[string]string `pulumi:"tags"`
}


type PartnerDestinationArgs struct {
	ActivationState                 pulumi.StringPtrInput
	EndpointBaseUrl                 pulumi.StringPtrInput
	EndpointServiceContext          pulumi.StringPtrInput
	ExpirationTimeIfNotActivatedUtc pulumi.StringPtrInput
	Location                        pulumi.StringPtrInput
	MessageForActivation            pulumi.StringPtrInput
	PartnerDestinationName          pulumi.StringPtrInput
	PartnerRegistrationImmutableId  pulumi.StringPtrInput
	ProvisioningState               pulumi.StringPtrInput
	ResourceGroupName               pulumi.StringInput
	Tags                            pulumi.StringMapInput
}

func (PartnerDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerDestinationArgs)(nil)).Elem()
}

type PartnerDestinationInput interface {
	pulumi.Input

	ToPartnerDestinationOutput() PartnerDestinationOutput
	ToPartnerDestinationOutputWithContext(ctx context.Context) PartnerDestinationOutput
}

func (*PartnerDestination) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerDestination)(nil)).Elem()
}

func (i *PartnerDestination) ToPartnerDestinationOutput() PartnerDestinationOutput {
	return i.ToPartnerDestinationOutputWithContext(context.Background())
}

func (i *PartnerDestination) ToPartnerDestinationOutputWithContext(ctx context.Context) PartnerDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerDestinationOutput)
}

type PartnerDestinationOutput struct{ *pulumi.OutputState }

func (PartnerDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerDestination)(nil)).Elem()
}

func (o PartnerDestinationOutput) ToPartnerDestinationOutput() PartnerDestinationOutput {
	return o
}

func (o PartnerDestinationOutput) ToPartnerDestinationOutputWithContext(ctx context.Context) PartnerDestinationOutput {
	return o
}

func (o PartnerDestinationOutput) ActivationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerDestination) pulumi.StringPtrOutput { return v.ActivationState }).(pulumi.StringPtrOutput)
}

func (o PartnerDestinationOutput) EndpointBaseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerDestination) pulumi.StringPtrOutput { return v.EndpointBaseUrl }).(pulumi.StringPtrOutput)
}

func (o PartnerDestinationOutput) EndpointServiceContext() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerDestination) pulumi.StringPtrOutput { return v.EndpointServiceContext }).(pulumi.StringPtrOutput)
}

func (o PartnerDestinationOutput) ExpirationTimeIfNotActivatedUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerDestination) pulumi.StringPtrOutput { return v.ExpirationTimeIfNotActivatedUtc }).(pulumi.StringPtrOutput)
}

func (o PartnerDestinationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerDestination) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PartnerDestinationOutput) MessageForActivation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerDestination) pulumi.StringPtrOutput { return v.MessageForActivation }).(pulumi.StringPtrOutput)
}

func (o PartnerDestinationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerDestination) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PartnerDestinationOutput) PartnerRegistrationImmutableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerDestination) pulumi.StringPtrOutput { return v.PartnerRegistrationImmutableId }).(pulumi.StringPtrOutput)
}

func (o PartnerDestinationOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerDestination) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o PartnerDestinationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PartnerDestination) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PartnerDestinationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PartnerDestination) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PartnerDestinationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerDestination) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PartnerDestinationOutput{})
}

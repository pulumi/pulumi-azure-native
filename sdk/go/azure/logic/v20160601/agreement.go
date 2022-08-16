


package v20160601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Agreement struct {
	pulumi.CustomResourceState

	AgreementType pulumi.StringOutput            `pulumi:"agreementType"`
	ChangedTime   pulumi.StringOutput            `pulumi:"changedTime"`
	Content       AgreementContentResponseOutput `pulumi:"content"`
	CreatedTime   pulumi.StringOutput            `pulumi:"createdTime"`
	GuestIdentity BusinessIdentityResponseOutput `pulumi:"guestIdentity"`
	GuestPartner  pulumi.StringOutput            `pulumi:"guestPartner"`
	HostIdentity  BusinessIdentityResponseOutput `pulumi:"hostIdentity"`
	HostPartner   pulumi.StringOutput            `pulumi:"hostPartner"`
	Location      pulumi.StringPtrOutput         `pulumi:"location"`
	Metadata      pulumi.AnyOutput               `pulumi:"metadata"`
	Name          pulumi.StringOutput            `pulumi:"name"`
	Tags          pulumi.StringMapOutput         `pulumi:"tags"`
	Type          pulumi.StringOutput            `pulumi:"type"`
}


func NewAgreement(ctx *pulumi.Context,
	name string, args *AgreementArgs, opts ...pulumi.ResourceOption) (*Agreement, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgreementType == nil {
		return nil, errors.New("invalid value for required argument 'AgreementType'")
	}
	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.GuestIdentity == nil {
		return nil, errors.New("invalid value for required argument 'GuestIdentity'")
	}
	if args.GuestPartner == nil {
		return nil, errors.New("invalid value for required argument 'GuestPartner'")
	}
	if args.HostIdentity == nil {
		return nil, errors.New("invalid value for required argument 'HostIdentity'")
	}
	if args.HostPartner == nil {
		return nil, errors.New("invalid value for required argument 'HostPartner'")
	}
	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic:Agreement"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20150801preview:Agreement"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:Agreement"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:Agreement"),
		},
	})
	opts = append(opts, aliases)
	var resource Agreement
	err := ctx.RegisterResource("azure-native:logic/v20160601:Agreement", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAgreement(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgreementState, opts ...pulumi.ResourceOption) (*Agreement, error) {
	var resource Agreement
	err := ctx.ReadResource("azure-native:logic/v20160601:Agreement", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type agreementState struct {
}

type AgreementState struct {
}

func (AgreementState) ElementType() reflect.Type {
	return reflect.TypeOf((*agreementState)(nil)).Elem()
}

type agreementArgs struct {
	AgreementName          *string           `pulumi:"agreementName"`
	AgreementType          AgreementType     `pulumi:"agreementType"`
	Content                AgreementContent  `pulumi:"content"`
	GuestIdentity          BusinessIdentity  `pulumi:"guestIdentity"`
	GuestPartner           string            `pulumi:"guestPartner"`
	HostIdentity           BusinessIdentity  `pulumi:"hostIdentity"`
	HostPartner            string            `pulumi:"hostPartner"`
	IntegrationAccountName string            `pulumi:"integrationAccountName"`
	Location               *string           `pulumi:"location"`
	Metadata               interface{}       `pulumi:"metadata"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	Tags                   map[string]string `pulumi:"tags"`
}


type AgreementArgs struct {
	AgreementName          pulumi.StringPtrInput
	AgreementType          AgreementTypeInput
	Content                AgreementContentInput
	GuestIdentity          BusinessIdentityInput
	GuestPartner           pulumi.StringInput
	HostIdentity           BusinessIdentityInput
	HostPartner            pulumi.StringInput
	IntegrationAccountName pulumi.StringInput
	Location               pulumi.StringPtrInput
	Metadata               pulumi.Input
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
}

func (AgreementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agreementArgs)(nil)).Elem()
}

type AgreementInput interface {
	pulumi.Input

	ToAgreementOutput() AgreementOutput
	ToAgreementOutputWithContext(ctx context.Context) AgreementOutput
}

func (*Agreement) ElementType() reflect.Type {
	return reflect.TypeOf((**Agreement)(nil)).Elem()
}

func (i *Agreement) ToAgreementOutput() AgreementOutput {
	return i.ToAgreementOutputWithContext(context.Background())
}

func (i *Agreement) ToAgreementOutputWithContext(ctx context.Context) AgreementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgreementOutput)
}

type AgreementOutput struct{ *pulumi.OutputState }

func (AgreementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Agreement)(nil)).Elem()
}

func (o AgreementOutput) ToAgreementOutput() AgreementOutput {
	return o
}

func (o AgreementOutput) ToAgreementOutputWithContext(ctx context.Context) AgreementOutput {
	return o
}

func (o AgreementOutput) AgreementType() pulumi.StringOutput {
	return o.ApplyT(func(v *Agreement) pulumi.StringOutput { return v.AgreementType }).(pulumi.StringOutput)
}

func (o AgreementOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Agreement) pulumi.StringOutput { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o AgreementOutput) Content() AgreementContentResponseOutput {
	return o.ApplyT(func(v *Agreement) AgreementContentResponseOutput { return v.Content }).(AgreementContentResponseOutput)
}

func (o AgreementOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Agreement) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o AgreementOutput) GuestIdentity() BusinessIdentityResponseOutput {
	return o.ApplyT(func(v *Agreement) BusinessIdentityResponseOutput { return v.GuestIdentity }).(BusinessIdentityResponseOutput)
}

func (o AgreementOutput) GuestPartner() pulumi.StringOutput {
	return o.ApplyT(func(v *Agreement) pulumi.StringOutput { return v.GuestPartner }).(pulumi.StringOutput)
}

func (o AgreementOutput) HostIdentity() BusinessIdentityResponseOutput {
	return o.ApplyT(func(v *Agreement) BusinessIdentityResponseOutput { return v.HostIdentity }).(BusinessIdentityResponseOutput)
}

func (o AgreementOutput) HostPartner() pulumi.StringOutput {
	return o.ApplyT(func(v *Agreement) pulumi.StringOutput { return v.HostPartner }).(pulumi.StringOutput)
}

func (o AgreementOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Agreement) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o AgreementOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *Agreement) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

func (o AgreementOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Agreement) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AgreementOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Agreement) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AgreementOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Agreement) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AgreementOutput{})
}




package v20150801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationAccountAgreement struct {
	pulumi.CustomResourceState

	AgreementType pulumi.StringPtrOutput            `pulumi:"agreementType"`
	ChangedTime   pulumi.StringOutput               `pulumi:"changedTime"`
	Content       AgreementContentResponsePtrOutput `pulumi:"content"`
	CreatedTime   pulumi.StringOutput               `pulumi:"createdTime"`
	GuestIdentity BusinessIdentityResponsePtrOutput `pulumi:"guestIdentity"`
	GuestPartner  pulumi.StringPtrOutput            `pulumi:"guestPartner"`
	HostIdentity  BusinessIdentityResponsePtrOutput `pulumi:"hostIdentity"`
	HostPartner   pulumi.StringPtrOutput            `pulumi:"hostPartner"`
	Location      pulumi.StringPtrOutput            `pulumi:"location"`
	Metadata      pulumi.AnyOutput                  `pulumi:"metadata"`
	Name          pulumi.StringPtrOutput            `pulumi:"name"`
	Tags          pulumi.StringMapOutput            `pulumi:"tags"`
	Type          pulumi.StringPtrOutput            `pulumi:"type"`
}


func NewIntegrationAccountAgreement(ctx *pulumi.Context,
	name string, args *IntegrationAccountAgreementArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountAgreement, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:logic/v20150801preview:IntegrationAccountAgreement"),
		},
		{
			Type: pulumi.String("azure-native:logic:IntegrationAccountAgreement"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic:IntegrationAccountAgreement"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:IntegrationAccountAgreement"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20160601:IntegrationAccountAgreement"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:IntegrationAccountAgreement"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20180701preview:IntegrationAccountAgreement"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:IntegrationAccountAgreement"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20190501:IntegrationAccountAgreement"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationAccountAgreement
	err := ctx.RegisterResource("azure-native:logic/v20150801preview:IntegrationAccountAgreement", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIntegrationAccountAgreement(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountAgreementState, opts ...pulumi.ResourceOption) (*IntegrationAccountAgreement, error) {
	var resource IntegrationAccountAgreement
	err := ctx.ReadResource("azure-native:logic/v20150801preview:IntegrationAccountAgreement", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type integrationAccountAgreementState struct {
}

type IntegrationAccountAgreementState struct {
}

func (IntegrationAccountAgreementState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountAgreementState)(nil)).Elem()
}

type integrationAccountAgreementArgs struct {
	AgreementName          *string           `pulumi:"agreementName"`
	AgreementType          *AgreementType    `pulumi:"agreementType"`
	Content                *AgreementContent `pulumi:"content"`
	GuestIdentity          *BusinessIdentity `pulumi:"guestIdentity"`
	GuestPartner           *string           `pulumi:"guestPartner"`
	HostIdentity           *BusinessIdentity `pulumi:"hostIdentity"`
	HostPartner            *string           `pulumi:"hostPartner"`
	Id                     *string           `pulumi:"id"`
	IntegrationAccountName string            `pulumi:"integrationAccountName"`
	Location               *string           `pulumi:"location"`
	Metadata               interface{}       `pulumi:"metadata"`
	Name                   *string           `pulumi:"name"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	Tags                   map[string]string `pulumi:"tags"`
	Type                   *string           `pulumi:"type"`
}


type IntegrationAccountAgreementArgs struct {
	AgreementName          pulumi.StringPtrInput
	AgreementType          AgreementTypePtrInput
	Content                AgreementContentPtrInput
	GuestIdentity          BusinessIdentityPtrInput
	GuestPartner           pulumi.StringPtrInput
	HostIdentity           BusinessIdentityPtrInput
	HostPartner            pulumi.StringPtrInput
	Id                     pulumi.StringPtrInput
	IntegrationAccountName pulumi.StringInput
	Location               pulumi.StringPtrInput
	Metadata               pulumi.Input
	Name                   pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
	Type                   pulumi.StringPtrInput
}

func (IntegrationAccountAgreementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountAgreementArgs)(nil)).Elem()
}

type IntegrationAccountAgreementInput interface {
	pulumi.Input

	ToIntegrationAccountAgreementOutput() IntegrationAccountAgreementOutput
	ToIntegrationAccountAgreementOutputWithContext(ctx context.Context) IntegrationAccountAgreementOutput
}

func (*IntegrationAccountAgreement) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountAgreement)(nil))
}

func (i *IntegrationAccountAgreement) ToIntegrationAccountAgreementOutput() IntegrationAccountAgreementOutput {
	return i.ToIntegrationAccountAgreementOutputWithContext(context.Background())
}

func (i *IntegrationAccountAgreement) ToIntegrationAccountAgreementOutputWithContext(ctx context.Context) IntegrationAccountAgreementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountAgreementOutput)
}

type IntegrationAccountAgreementOutput struct{ *pulumi.OutputState }

func (IntegrationAccountAgreementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountAgreement)(nil))
}

func (o IntegrationAccountAgreementOutput) ToIntegrationAccountAgreementOutput() IntegrationAccountAgreementOutput {
	return o
}

func (o IntegrationAccountAgreementOutput) ToIntegrationAccountAgreementOutputWithContext(ctx context.Context) IntegrationAccountAgreementOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationAccountAgreementInput)(nil)).Elem(), &IntegrationAccountAgreement{})
	pulumi.RegisterOutputType(IntegrationAccountAgreementOutput{})
}

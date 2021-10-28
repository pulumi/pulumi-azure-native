


package v20150801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationAccountPartner struct {
	pulumi.CustomResourceState

	ChangedTime pulumi.StringOutput             `pulumi:"changedTime"`
	Content     PartnerContentResponsePtrOutput `pulumi:"content"`
	CreatedTime pulumi.StringOutput             `pulumi:"createdTime"`
	Location    pulumi.StringPtrOutput          `pulumi:"location"`
	Metadata    pulumi.AnyOutput                `pulumi:"metadata"`
	Name        pulumi.StringPtrOutput          `pulumi:"name"`
	PartnerType pulumi.StringPtrOutput          `pulumi:"partnerType"`
	Tags        pulumi.StringMapOutput          `pulumi:"tags"`
	Type        pulumi.StringPtrOutput          `pulumi:"type"`
}


func NewIntegrationAccountPartner(ctx *pulumi.Context,
	name string, args *IntegrationAccountPartnerArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountPartner, error) {
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
			Type: pulumi.String("azure-nextgen:logic/v20150801preview:IntegrationAccountPartner"),
		},
		{
			Type: pulumi.String("azure-native:logic:IntegrationAccountPartner"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic:IntegrationAccountPartner"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:IntegrationAccountPartner"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20160601:IntegrationAccountPartner"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:IntegrationAccountPartner"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20180701preview:IntegrationAccountPartner"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:IntegrationAccountPartner"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20190501:IntegrationAccountPartner"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationAccountPartner
	err := ctx.RegisterResource("azure-native:logic/v20150801preview:IntegrationAccountPartner", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIntegrationAccountPartner(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountPartnerState, opts ...pulumi.ResourceOption) (*IntegrationAccountPartner, error) {
	var resource IntegrationAccountPartner
	err := ctx.ReadResource("azure-native:logic/v20150801preview:IntegrationAccountPartner", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type integrationAccountPartnerState struct {
}

type IntegrationAccountPartnerState struct {
}

func (IntegrationAccountPartnerState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountPartnerState)(nil)).Elem()
}

type integrationAccountPartnerArgs struct {
	Content                *PartnerContent   `pulumi:"content"`
	Id                     *string           `pulumi:"id"`
	IntegrationAccountName string            `pulumi:"integrationAccountName"`
	Location               *string           `pulumi:"location"`
	Metadata               interface{}       `pulumi:"metadata"`
	Name                   *string           `pulumi:"name"`
	PartnerName            *string           `pulumi:"partnerName"`
	PartnerType            *PartnerType      `pulumi:"partnerType"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	Tags                   map[string]string `pulumi:"tags"`
	Type                   *string           `pulumi:"type"`
}


type IntegrationAccountPartnerArgs struct {
	Content                PartnerContentPtrInput
	Id                     pulumi.StringPtrInput
	IntegrationAccountName pulumi.StringInput
	Location               pulumi.StringPtrInput
	Metadata               pulumi.Input
	Name                   pulumi.StringPtrInput
	PartnerName            pulumi.StringPtrInput
	PartnerType            PartnerTypePtrInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
	Type                   pulumi.StringPtrInput
}

func (IntegrationAccountPartnerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountPartnerArgs)(nil)).Elem()
}

type IntegrationAccountPartnerInput interface {
	pulumi.Input

	ToIntegrationAccountPartnerOutput() IntegrationAccountPartnerOutput
	ToIntegrationAccountPartnerOutputWithContext(ctx context.Context) IntegrationAccountPartnerOutput
}

func (*IntegrationAccountPartner) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountPartner)(nil))
}

func (i *IntegrationAccountPartner) ToIntegrationAccountPartnerOutput() IntegrationAccountPartnerOutput {
	return i.ToIntegrationAccountPartnerOutputWithContext(context.Background())
}

func (i *IntegrationAccountPartner) ToIntegrationAccountPartnerOutputWithContext(ctx context.Context) IntegrationAccountPartnerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountPartnerOutput)
}

type IntegrationAccountPartnerOutput struct{ *pulumi.OutputState }

func (IntegrationAccountPartnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountPartner)(nil))
}

func (o IntegrationAccountPartnerOutput) ToIntegrationAccountPartnerOutput() IntegrationAccountPartnerOutput {
	return o
}

func (o IntegrationAccountPartnerOutput) ToIntegrationAccountPartnerOutputWithContext(ctx context.Context) IntegrationAccountPartnerOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationAccountPartnerInput)(nil)).Elem(), &IntegrationAccountPartner{})
	pulumi.RegisterOutputType(IntegrationAccountPartnerOutput{})
}




package v20160601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Partner struct {
	pulumi.CustomResourceState

	ChangedTime pulumi.StringOutput          `pulumi:"changedTime"`
	Content     PartnerContentResponseOutput `pulumi:"content"`
	CreatedTime pulumi.StringOutput          `pulumi:"createdTime"`
	Location    pulumi.StringPtrOutput       `pulumi:"location"`
	Metadata    pulumi.AnyOutput             `pulumi:"metadata"`
	Name        pulumi.StringOutput          `pulumi:"name"`
	PartnerType pulumi.StringOutput          `pulumi:"partnerType"`
	Tags        pulumi.StringMapOutput       `pulumi:"tags"`
	Type        pulumi.StringOutput          `pulumi:"type"`
}


func NewPartner(ctx *pulumi.Context,
	name string, args *PartnerArgs, opts ...pulumi.ResourceOption) (*Partner, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.PartnerType == nil {
		return nil, errors.New("invalid value for required argument 'PartnerType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:logic/v20160601:Partner"),
		},
		{
			Type: pulumi.String("azure-native:logic:Partner"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic:Partner"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20150801preview:Partner"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20150801preview:Partner"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:Partner"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20180701preview:Partner"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:Partner"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20190501:Partner"),
		},
	})
	opts = append(opts, aliases)
	var resource Partner
	err := ctx.RegisterResource("azure-native:logic/v20160601:Partner", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPartner(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartnerState, opts ...pulumi.ResourceOption) (*Partner, error) {
	var resource Partner
	err := ctx.ReadResource("azure-native:logic/v20160601:Partner", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type partnerState struct {
}

type PartnerState struct {
}

func (PartnerState) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerState)(nil)).Elem()
}

type partnerArgs struct {
	Content                PartnerContent    `pulumi:"content"`
	IntegrationAccountName string            `pulumi:"integrationAccountName"`
	Location               *string           `pulumi:"location"`
	Metadata               interface{}       `pulumi:"metadata"`
	PartnerName            *string           `pulumi:"partnerName"`
	PartnerType            PartnerType       `pulumi:"partnerType"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	Tags                   map[string]string `pulumi:"tags"`
}


type PartnerArgs struct {
	Content                PartnerContentInput
	IntegrationAccountName pulumi.StringInput
	Location               pulumi.StringPtrInput
	Metadata               pulumi.Input
	PartnerName            pulumi.StringPtrInput
	PartnerType            PartnerTypeInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
}

func (PartnerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerArgs)(nil)).Elem()
}

type PartnerInput interface {
	pulumi.Input

	ToPartnerOutput() PartnerOutput
	ToPartnerOutputWithContext(ctx context.Context) PartnerOutput
}

func (*Partner) ElementType() reflect.Type {
	return reflect.TypeOf((*Partner)(nil))
}

func (i *Partner) ToPartnerOutput() PartnerOutput {
	return i.ToPartnerOutputWithContext(context.Background())
}

func (i *Partner) ToPartnerOutputWithContext(ctx context.Context) PartnerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerOutput)
}

type PartnerOutput struct{ *pulumi.OutputState }

func (PartnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Partner)(nil))
}

func (o PartnerOutput) ToPartnerOutput() PartnerOutput {
	return o
}

func (o PartnerOutput) ToPartnerOutputWithContext(ctx context.Context) PartnerOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PartnerInput)(nil)).Elem(), &Partner{})
	pulumi.RegisterOutputType(PartnerOutput{})
}

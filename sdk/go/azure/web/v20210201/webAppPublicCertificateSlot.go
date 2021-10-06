


package v20210201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppPublicCertificateSlot struct {
	pulumi.CustomResourceState

	Blob                      pulumi.StringPtrOutput `pulumi:"blob"`
	Kind                      pulumi.StringPtrOutput `pulumi:"kind"`
	Name                      pulumi.StringOutput    `pulumi:"name"`
	PublicCertificateLocation pulumi.StringPtrOutput `pulumi:"publicCertificateLocation"`
	Thumbprint                pulumi.StringOutput    `pulumi:"thumbprint"`
	Type                      pulumi.StringOutput    `pulumi:"type"`
}


func NewWebAppPublicCertificateSlot(ctx *pulumi.Context,
	name string, args *WebAppPublicCertificateSlotArgs, opts ...pulumi.ResourceOption) (*WebAppPublicCertificateSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:WebAppPublicCertificateSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppPublicCertificateSlot
	err := ctx.RegisterResource("azure-native:web/v20210201:WebAppPublicCertificateSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppPublicCertificateSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppPublicCertificateSlotState, opts ...pulumi.ResourceOption) (*WebAppPublicCertificateSlot, error) {
	var resource WebAppPublicCertificateSlot
	err := ctx.ReadResource("azure-native:web/v20210201:WebAppPublicCertificateSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppPublicCertificateSlotState struct {
}

type WebAppPublicCertificateSlotState struct {
}

func (WebAppPublicCertificateSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPublicCertificateSlotState)(nil)).Elem()
}

type webAppPublicCertificateSlotArgs struct {
	Blob                      *string                    `pulumi:"blob"`
	Kind                      *string                    `pulumi:"kind"`
	Name                      string                     `pulumi:"name"`
	PublicCertificateLocation *PublicCertificateLocation `pulumi:"publicCertificateLocation"`
	PublicCertificateName     *string                    `pulumi:"publicCertificateName"`
	ResourceGroupName         string                     `pulumi:"resourceGroupName"`
	Slot                      string                     `pulumi:"slot"`
}


type WebAppPublicCertificateSlotArgs struct {
	Blob                      pulumi.StringPtrInput
	Kind                      pulumi.StringPtrInput
	Name                      pulumi.StringInput
	PublicCertificateLocation PublicCertificateLocationPtrInput
	PublicCertificateName     pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	Slot                      pulumi.StringInput
}

func (WebAppPublicCertificateSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPublicCertificateSlotArgs)(nil)).Elem()
}

type WebAppPublicCertificateSlotInput interface {
	pulumi.Input

	ToWebAppPublicCertificateSlotOutput() WebAppPublicCertificateSlotOutput
	ToWebAppPublicCertificateSlotOutputWithContext(ctx context.Context) WebAppPublicCertificateSlotOutput
}

func (*WebAppPublicCertificateSlot) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppPublicCertificateSlot)(nil))
}

func (i *WebAppPublicCertificateSlot) ToWebAppPublicCertificateSlotOutput() WebAppPublicCertificateSlotOutput {
	return i.ToWebAppPublicCertificateSlotOutputWithContext(context.Background())
}

func (i *WebAppPublicCertificateSlot) ToWebAppPublicCertificateSlotOutputWithContext(ctx context.Context) WebAppPublicCertificateSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppPublicCertificateSlotOutput)
}

type WebAppPublicCertificateSlotOutput struct{ *pulumi.OutputState }

func (WebAppPublicCertificateSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppPublicCertificateSlot)(nil))
}

func (o WebAppPublicCertificateSlotOutput) ToWebAppPublicCertificateSlotOutput() WebAppPublicCertificateSlotOutput {
	return o
}

func (o WebAppPublicCertificateSlotOutput) ToWebAppPublicCertificateSlotOutputWithContext(ctx context.Context) WebAppPublicCertificateSlotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppPublicCertificateSlotOutput{})
}

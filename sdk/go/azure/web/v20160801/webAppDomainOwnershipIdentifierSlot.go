


package v20160801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppDomainOwnershipIdentifierSlot struct {
	pulumi.CustomResourceState

	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	Name pulumi.StringOutput    `pulumi:"name"`
	Type pulumi.StringOutput    `pulumi:"type"`
}


func NewWebAppDomainOwnershipIdentifierSlot(ctx *pulumi.Context,
	name string, args *WebAppDomainOwnershipIdentifierSlotArgs, opts ...pulumi.ResourceOption) (*WebAppDomainOwnershipIdentifierSlot, error) {
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
			Type: pulumi.String("azure-native:web:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppDomainOwnershipIdentifierSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppDomainOwnershipIdentifierSlot
	err := ctx.RegisterResource("azure-native:web/v20160801:WebAppDomainOwnershipIdentifierSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppDomainOwnershipIdentifierSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppDomainOwnershipIdentifierSlotState, opts ...pulumi.ResourceOption) (*WebAppDomainOwnershipIdentifierSlot, error) {
	var resource WebAppDomainOwnershipIdentifierSlot
	err := ctx.ReadResource("azure-native:web/v20160801:WebAppDomainOwnershipIdentifierSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppDomainOwnershipIdentifierSlotState struct {
}

type WebAppDomainOwnershipIdentifierSlotState struct {
}

func (WebAppDomainOwnershipIdentifierSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppDomainOwnershipIdentifierSlotState)(nil)).Elem()
}

type webAppDomainOwnershipIdentifierSlotArgs struct {
	DomainOwnershipIdentifierName *string `pulumi:"domainOwnershipIdentifierName"`
	Id                            *string `pulumi:"id"`
	Kind                          *string `pulumi:"kind"`
	Name                          string  `pulumi:"name"`
	ResourceGroupName             string  `pulumi:"resourceGroupName"`
	Slot                          string  `pulumi:"slot"`
}


type WebAppDomainOwnershipIdentifierSlotArgs struct {
	DomainOwnershipIdentifierName pulumi.StringPtrInput
	Id                            pulumi.StringPtrInput
	Kind                          pulumi.StringPtrInput
	Name                          pulumi.StringInput
	ResourceGroupName             pulumi.StringInput
	Slot                          pulumi.StringInput
}

func (WebAppDomainOwnershipIdentifierSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppDomainOwnershipIdentifierSlotArgs)(nil)).Elem()
}

type WebAppDomainOwnershipIdentifierSlotInput interface {
	pulumi.Input

	ToWebAppDomainOwnershipIdentifierSlotOutput() WebAppDomainOwnershipIdentifierSlotOutput
	ToWebAppDomainOwnershipIdentifierSlotOutputWithContext(ctx context.Context) WebAppDomainOwnershipIdentifierSlotOutput
}

func (*WebAppDomainOwnershipIdentifierSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppDomainOwnershipIdentifierSlot)(nil)).Elem()
}

func (i *WebAppDomainOwnershipIdentifierSlot) ToWebAppDomainOwnershipIdentifierSlotOutput() WebAppDomainOwnershipIdentifierSlotOutput {
	return i.ToWebAppDomainOwnershipIdentifierSlotOutputWithContext(context.Background())
}

func (i *WebAppDomainOwnershipIdentifierSlot) ToWebAppDomainOwnershipIdentifierSlotOutputWithContext(ctx context.Context) WebAppDomainOwnershipIdentifierSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppDomainOwnershipIdentifierSlotOutput)
}

type WebAppDomainOwnershipIdentifierSlotOutput struct{ *pulumi.OutputState }

func (WebAppDomainOwnershipIdentifierSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppDomainOwnershipIdentifierSlot)(nil)).Elem()
}

func (o WebAppDomainOwnershipIdentifierSlotOutput) ToWebAppDomainOwnershipIdentifierSlotOutput() WebAppDomainOwnershipIdentifierSlotOutput {
	return o
}

func (o WebAppDomainOwnershipIdentifierSlotOutput) ToWebAppDomainOwnershipIdentifierSlotOutputWithContext(ctx context.Context) WebAppDomainOwnershipIdentifierSlotOutput {
	return o
}

func (o WebAppDomainOwnershipIdentifierSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppDomainOwnershipIdentifierSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppDomainOwnershipIdentifierSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppDomainOwnershipIdentifierSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppDomainOwnershipIdentifierSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppDomainOwnershipIdentifierSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppDomainOwnershipIdentifierSlotOutput{})
}

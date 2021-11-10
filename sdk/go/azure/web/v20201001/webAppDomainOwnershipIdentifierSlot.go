


package v20201001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppDomainOwnershipIdentifierSlot struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringPtrOutput   `pulumi:"kind"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
	Value      pulumi.StringPtrOutput   `pulumi:"value"`
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
			Type: pulumi.String("azure-native:web/v20160801:WebAppDomainOwnershipIdentifierSlot"),
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
	})
	opts = append(opts, aliases)
	var resource WebAppDomainOwnershipIdentifierSlot
	err := ctx.RegisterResource("azure-native:web/v20201001:WebAppDomainOwnershipIdentifierSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppDomainOwnershipIdentifierSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppDomainOwnershipIdentifierSlotState, opts ...pulumi.ResourceOption) (*WebAppDomainOwnershipIdentifierSlot, error) {
	var resource WebAppDomainOwnershipIdentifierSlot
	err := ctx.ReadResource("azure-native:web/v20201001:WebAppDomainOwnershipIdentifierSlot", name, id, state, &resource, opts...)
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
	Kind                          *string `pulumi:"kind"`
	Name                          string  `pulumi:"name"`
	ResourceGroupName             string  `pulumi:"resourceGroupName"`
	Slot                          string  `pulumi:"slot"`
	Value                         *string `pulumi:"value"`
}


type WebAppDomainOwnershipIdentifierSlotArgs struct {
	DomainOwnershipIdentifierName pulumi.StringPtrInput
	Kind                          pulumi.StringPtrInput
	Name                          pulumi.StringInput
	ResourceGroupName             pulumi.StringInput
	Slot                          pulumi.StringInput
	Value                         pulumi.StringPtrInput
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
	return reflect.TypeOf((*WebAppDomainOwnershipIdentifierSlot)(nil))
}

func (i *WebAppDomainOwnershipIdentifierSlot) ToWebAppDomainOwnershipIdentifierSlotOutput() WebAppDomainOwnershipIdentifierSlotOutput {
	return i.ToWebAppDomainOwnershipIdentifierSlotOutputWithContext(context.Background())
}

func (i *WebAppDomainOwnershipIdentifierSlot) ToWebAppDomainOwnershipIdentifierSlotOutputWithContext(ctx context.Context) WebAppDomainOwnershipIdentifierSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppDomainOwnershipIdentifierSlotOutput)
}

type WebAppDomainOwnershipIdentifierSlotOutput struct{ *pulumi.OutputState }

func (WebAppDomainOwnershipIdentifierSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppDomainOwnershipIdentifierSlot)(nil))
}

func (o WebAppDomainOwnershipIdentifierSlotOutput) ToWebAppDomainOwnershipIdentifierSlotOutput() WebAppDomainOwnershipIdentifierSlotOutput {
	return o
}

func (o WebAppDomainOwnershipIdentifierSlotOutput) ToWebAppDomainOwnershipIdentifierSlotOutputWithContext(ctx context.Context) WebAppDomainOwnershipIdentifierSlotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppDomainOwnershipIdentifierSlotOutput{})
}

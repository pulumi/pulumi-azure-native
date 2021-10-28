


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Product struct {
	pulumi.CustomResourceState

	ApprovalRequired     pulumi.BoolPtrOutput   `pulumi:"approvalRequired"`
	Description          pulumi.StringPtrOutput `pulumi:"description"`
	DisplayName          pulumi.StringOutput    `pulumi:"displayName"`
	Name                 pulumi.StringOutput    `pulumi:"name"`
	State                pulumi.StringPtrOutput `pulumi:"state"`
	SubscriptionRequired pulumi.BoolPtrOutput   `pulumi:"subscriptionRequired"`
	SubscriptionsLimit   pulumi.IntPtrOutput    `pulumi:"subscriptionsLimit"`
	Terms                pulumi.StringPtrOutput `pulumi:"terms"`
	Type                 pulumi.StringOutput    `pulumi:"type"`
}


func NewProduct(ctx *pulumi.Context,
	name string, args *ProductArgs, opts ...pulumi.ResourceOption) (*Product, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210401preview:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement:Product"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:Product"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20160707:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:Product"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20161010:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Product"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20170301:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:Product"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180101:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Product"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180601preview:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Product"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20190101:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Product"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Product"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Product"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Product"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20201201:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Product"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210101preview:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Product"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210801:Product"),
		},
	})
	opts = append(opts, aliases)
	var resource Product
	err := ctx.RegisterResource("azure-native:apimanagement/v20210401preview:Product", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProduct(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductState, opts ...pulumi.ResourceOption) (*Product, error) {
	var resource Product
	err := ctx.ReadResource("azure-native:apimanagement/v20210401preview:Product", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type productState struct {
}

type ProductState struct {
}

func (ProductState) ElementType() reflect.Type {
	return reflect.TypeOf((*productState)(nil)).Elem()
}

type productArgs struct {
	ApprovalRequired     *bool             `pulumi:"approvalRequired"`
	Description          *string           `pulumi:"description"`
	DisplayName          string            `pulumi:"displayName"`
	ProductId            *string           `pulumi:"productId"`
	ResourceGroupName    string            `pulumi:"resourceGroupName"`
	ServiceName          string            `pulumi:"serviceName"`
	State                *ProductStateEnum `pulumi:"state"`
	SubscriptionRequired *bool             `pulumi:"subscriptionRequired"`
	SubscriptionsLimit   *int              `pulumi:"subscriptionsLimit"`
	Terms                *string           `pulumi:"terms"`
}


type ProductArgs struct {
	ApprovalRequired     pulumi.BoolPtrInput
	Description          pulumi.StringPtrInput
	DisplayName          pulumi.StringInput
	ProductId            pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	ServiceName          pulumi.StringInput
	State                ProductStateEnumPtrInput
	SubscriptionRequired pulumi.BoolPtrInput
	SubscriptionsLimit   pulumi.IntPtrInput
	Terms                pulumi.StringPtrInput
}

func (ProductArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productArgs)(nil)).Elem()
}

type ProductInput interface {
	pulumi.Input

	ToProductOutput() ProductOutput
	ToProductOutputWithContext(ctx context.Context) ProductOutput
}

func (*Product) ElementType() reflect.Type {
	return reflect.TypeOf((*Product)(nil))
}

func (i *Product) ToProductOutput() ProductOutput {
	return i.ToProductOutputWithContext(context.Background())
}

func (i *Product) ToProductOutputWithContext(ctx context.Context) ProductOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductOutput)
}

type ProductOutput struct{ *pulumi.OutputState }

func (ProductOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Product)(nil))
}

func (o ProductOutput) ToProductOutput() ProductOutput {
	return o
}

func (o ProductOutput) ToProductOutputWithContext(ctx context.Context) ProductOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProductOutput{})
}

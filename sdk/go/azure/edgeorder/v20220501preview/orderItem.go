


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OrderItem struct {
	pulumi.CustomResourceState

	AddressDetails   AddressDetailsResponseOutput   `pulumi:"addressDetails"`
	Location         pulumi.StringOutput            `pulumi:"location"`
	Name             pulumi.StringOutput            `pulumi:"name"`
	OrderId          pulumi.StringOutput            `pulumi:"orderId"`
	OrderItemDetails OrderItemDetailsResponseOutput `pulumi:"orderItemDetails"`
	StartTime        pulumi.StringOutput            `pulumi:"startTime"`
	SystemData       SystemDataResponseOutput       `pulumi:"systemData"`
	Tags             pulumi.StringMapOutput         `pulumi:"tags"`
	Type             pulumi.StringOutput            `pulumi:"type"`
}


func NewOrderItem(ctx *pulumi.Context,
	name string, args *OrderItemArgs, opts ...pulumi.ResourceOption) (*OrderItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddressDetails == nil {
		return nil, errors.New("invalid value for required argument 'AddressDetails'")
	}
	if args.OrderId == nil {
		return nil, errors.New("invalid value for required argument 'OrderId'")
	}
	if args.OrderItemDetails == nil {
		return nil, errors.New("invalid value for required argument 'OrderItemDetails'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:edgeorder:OrderItem"),
		},
		{
			Type: pulumi.String("azure-native:edgeorder/v20201201preview:OrderItem"),
		},
		{
			Type: pulumi.String("azure-native:edgeorder/v20211201:OrderItem"),
		},
	})
	opts = append(opts, aliases)
	var resource OrderItem
	err := ctx.RegisterResource("azure-native:edgeorder/v20220501preview:OrderItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOrderItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrderItemState, opts ...pulumi.ResourceOption) (*OrderItem, error) {
	var resource OrderItem
	err := ctx.ReadResource("azure-native:edgeorder/v20220501preview:OrderItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type orderItemState struct {
}

type OrderItemState struct {
}

func (OrderItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*orderItemState)(nil)).Elem()
}

type orderItemArgs struct {
	AddressDetails    AddressDetails    `pulumi:"addressDetails"`
	Location          *string           `pulumi:"location"`
	OrderId           string            `pulumi:"orderId"`
	OrderItemDetails  OrderItemDetails  `pulumi:"orderItemDetails"`
	OrderItemName     *string           `pulumi:"orderItemName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type OrderItemArgs struct {
	AddressDetails    AddressDetailsInput
	Location          pulumi.StringPtrInput
	OrderId           pulumi.StringInput
	OrderItemDetails  OrderItemDetailsInput
	OrderItemName     pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (OrderItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orderItemArgs)(nil)).Elem()
}

type OrderItemInput interface {
	pulumi.Input

	ToOrderItemOutput() OrderItemOutput
	ToOrderItemOutputWithContext(ctx context.Context) OrderItemOutput
}

func (*OrderItem) ElementType() reflect.Type {
	return reflect.TypeOf((**OrderItem)(nil)).Elem()
}

func (i *OrderItem) ToOrderItemOutput() OrderItemOutput {
	return i.ToOrderItemOutputWithContext(context.Background())
}

func (i *OrderItem) ToOrderItemOutputWithContext(ctx context.Context) OrderItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderItemOutput)
}

type OrderItemOutput struct{ *pulumi.OutputState }

func (OrderItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrderItem)(nil)).Elem()
}

func (o OrderItemOutput) ToOrderItemOutput() OrderItemOutput {
	return o
}

func (o OrderItemOutput) ToOrderItemOutputWithContext(ctx context.Context) OrderItemOutput {
	return o
}

func (o OrderItemOutput) AddressDetails() AddressDetailsResponseOutput {
	return o.ApplyT(func(v *OrderItem) AddressDetailsResponseOutput { return v.AddressDetails }).(AddressDetailsResponseOutput)
}

func (o OrderItemOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderItem) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o OrderItemOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderItem) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OrderItemOutput) OrderId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderItem) pulumi.StringOutput { return v.OrderId }).(pulumi.StringOutput)
}

func (o OrderItemOutput) OrderItemDetails() OrderItemDetailsResponseOutput {
	return o.ApplyT(func(v *OrderItem) OrderItemDetailsResponseOutput { return v.OrderItemDetails }).(OrderItemDetailsResponseOutput)
}

func (o OrderItemOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderItem) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

func (o OrderItemOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *OrderItem) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o OrderItemOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OrderItem) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o OrderItemOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderItem) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OrderItemOutput{})
}

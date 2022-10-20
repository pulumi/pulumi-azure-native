


package v20201201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type OrderItemByName struct {
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


func NewOrderItemByName(ctx *pulumi.Context,
	name string, args *OrderItemByNameArgs, opts ...pulumi.ResourceOption) (*OrderItemByName, error) {
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
			Type: pulumi.String("azure-native:edgeorder:OrderItemByName"),
		},
		{
			Type: pulumi.String("azure-native:edgeorder/v20211201:OrderItemByName"),
		},
	})
	opts = append(opts, aliases)
	var resource OrderItemByName
	err := ctx.RegisterResource("azure-native:edgeorder/v20201201preview:OrderItemByName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOrderItemByName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrderItemByNameState, opts ...pulumi.ResourceOption) (*OrderItemByName, error) {
	var resource OrderItemByName
	err := ctx.ReadResource("azure-native:edgeorder/v20201201preview:OrderItemByName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type orderItemByNameState struct {
}

type OrderItemByNameState struct {
}

func (OrderItemByNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*orderItemByNameState)(nil)).Elem()
}

type orderItemByNameArgs struct {
	AddressDetails    AddressDetails    `pulumi:"addressDetails"`
	Location          *string           `pulumi:"location"`
	OrderId           string            `pulumi:"orderId"`
	OrderItemDetails  OrderItemDetails  `pulumi:"orderItemDetails"`
	OrderItemName     *string           `pulumi:"orderItemName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type OrderItemByNameArgs struct {
	AddressDetails    AddressDetailsInput
	Location          pulumi.StringPtrInput
	OrderId           pulumi.StringInput
	OrderItemDetails  OrderItemDetailsInput
	OrderItemName     pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (OrderItemByNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orderItemByNameArgs)(nil)).Elem()
}

type OrderItemByNameInput interface {
	pulumi.Input

	ToOrderItemByNameOutput() OrderItemByNameOutput
	ToOrderItemByNameOutputWithContext(ctx context.Context) OrderItemByNameOutput
}

func (*OrderItemByName) ElementType() reflect.Type {
	return reflect.TypeOf((**OrderItemByName)(nil)).Elem()
}

func (i *OrderItemByName) ToOrderItemByNameOutput() OrderItemByNameOutput {
	return i.ToOrderItemByNameOutputWithContext(context.Background())
}

func (i *OrderItemByName) ToOrderItemByNameOutputWithContext(ctx context.Context) OrderItemByNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderItemByNameOutput)
}

type OrderItemByNameOutput struct{ *pulumi.OutputState }

func (OrderItemByNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrderItemByName)(nil)).Elem()
}

func (o OrderItemByNameOutput) ToOrderItemByNameOutput() OrderItemByNameOutput {
	return o
}

func (o OrderItemByNameOutput) ToOrderItemByNameOutputWithContext(ctx context.Context) OrderItemByNameOutput {
	return o
}

func (o OrderItemByNameOutput) AddressDetails() AddressDetailsResponseOutput {
	return o.ApplyT(func(v *OrderItemByName) AddressDetailsResponseOutput { return v.AddressDetails }).(AddressDetailsResponseOutput)
}

func (o OrderItemByNameOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderItemByName) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o OrderItemByNameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderItemByName) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OrderItemByNameOutput) OrderId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderItemByName) pulumi.StringOutput { return v.OrderId }).(pulumi.StringOutput)
}

func (o OrderItemByNameOutput) OrderItemDetails() OrderItemDetailsResponseOutput {
	return o.ApplyT(func(v *OrderItemByName) OrderItemDetailsResponseOutput { return v.OrderItemDetails }).(OrderItemDetailsResponseOutput)
}

func (o OrderItemByNameOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderItemByName) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

func (o OrderItemByNameOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *OrderItemByName) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o OrderItemByNameOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OrderItemByName) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o OrderItemByNameOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderItemByName) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OrderItemByNameOutput{})
}




package v20180831

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type BudgetByResourceGroupName struct {
	pulumi.CustomResourceState

	Amount        pulumi.Float64Output           `pulumi:"amount"`
	Category      pulumi.StringOutput            `pulumi:"category"`
	CurrentSpend  CurrentSpendResponseOutput     `pulumi:"currentSpend"`
	ETag          pulumi.StringPtrOutput         `pulumi:"eTag"`
	Filters       FiltersResponsePtrOutput       `pulumi:"filters"`
	Name          pulumi.StringOutput            `pulumi:"name"`
	Notifications NotificationResponseMapOutput  `pulumi:"notifications"`
	TimeGrain     pulumi.StringOutput            `pulumi:"timeGrain"`
	TimePeriod    BudgetTimePeriodResponseOutput `pulumi:"timePeriod"`
	Type          pulumi.StringOutput            `pulumi:"type"`
}


func NewBudgetByResourceGroupName(ctx *pulumi.Context,
	name string, args *BudgetByResourceGroupNameArgs, opts ...pulumi.ResourceOption) (*BudgetByResourceGroupName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Amount == nil {
		return nil, errors.New("invalid value for required argument 'Amount'")
	}
	if args.Category == nil {
		return nil, errors.New("invalid value for required argument 'Category'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TimeGrain == nil {
		return nil, errors.New("invalid value for required argument 'TimeGrain'")
	}
	if args.TimePeriod == nil {
		return nil, errors.New("invalid value for required argument 'TimePeriod'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:consumption/v20180131:BudgetByResourceGroupName"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20180331:BudgetByResourceGroupName"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20180630:BudgetByResourceGroupName"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20181001:BudgetByResourceGroupName"),
		},
	})
	opts = append(opts, aliases)
	var resource BudgetByResourceGroupName
	err := ctx.RegisterResource("azure-native:consumption/v20180831:BudgetByResourceGroupName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBudgetByResourceGroupName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BudgetByResourceGroupNameState, opts ...pulumi.ResourceOption) (*BudgetByResourceGroupName, error) {
	var resource BudgetByResourceGroupName
	err := ctx.ReadResource("azure-native:consumption/v20180831:BudgetByResourceGroupName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type budgetByResourceGroupNameState struct {
}

type BudgetByResourceGroupNameState struct {
}

func (BudgetByResourceGroupNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetByResourceGroupNameState)(nil)).Elem()
}

type budgetByResourceGroupNameArgs struct {
	Amount            float64                 `pulumi:"amount"`
	BudgetName        *string                 `pulumi:"budgetName"`
	Category          string                  `pulumi:"category"`
	ETag              *string                 `pulumi:"eTag"`
	Filters           *Filters                `pulumi:"filters"`
	Notifications     map[string]Notification `pulumi:"notifications"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	TimeGrain         string                  `pulumi:"timeGrain"`
	TimePeriod        BudgetTimePeriod        `pulumi:"timePeriod"`
}


type BudgetByResourceGroupNameArgs struct {
	Amount            pulumi.Float64Input
	BudgetName        pulumi.StringPtrInput
	Category          pulumi.StringInput
	ETag              pulumi.StringPtrInput
	Filters           FiltersPtrInput
	Notifications     NotificationMapInput
	ResourceGroupName pulumi.StringInput
	TimeGrain         pulumi.StringInput
	TimePeriod        BudgetTimePeriodInput
}

func (BudgetByResourceGroupNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetByResourceGroupNameArgs)(nil)).Elem()
}

type BudgetByResourceGroupNameInput interface {
	pulumi.Input

	ToBudgetByResourceGroupNameOutput() BudgetByResourceGroupNameOutput
	ToBudgetByResourceGroupNameOutputWithContext(ctx context.Context) BudgetByResourceGroupNameOutput
}

func (*BudgetByResourceGroupName) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetByResourceGroupName)(nil)).Elem()
}

func (i *BudgetByResourceGroupName) ToBudgetByResourceGroupNameOutput() BudgetByResourceGroupNameOutput {
	return i.ToBudgetByResourceGroupNameOutputWithContext(context.Background())
}

func (i *BudgetByResourceGroupName) ToBudgetByResourceGroupNameOutputWithContext(ctx context.Context) BudgetByResourceGroupNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetByResourceGroupNameOutput)
}

type BudgetByResourceGroupNameOutput struct{ *pulumi.OutputState }

func (BudgetByResourceGroupNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetByResourceGroupName)(nil)).Elem()
}

func (o BudgetByResourceGroupNameOutput) ToBudgetByResourceGroupNameOutput() BudgetByResourceGroupNameOutput {
	return o
}

func (o BudgetByResourceGroupNameOutput) ToBudgetByResourceGroupNameOutputWithContext(ctx context.Context) BudgetByResourceGroupNameOutput {
	return o
}

func (o BudgetByResourceGroupNameOutput) Amount() pulumi.Float64Output {
	return o.ApplyT(func(v *BudgetByResourceGroupName) pulumi.Float64Output { return v.Amount }).(pulumi.Float64Output)
}

func (o BudgetByResourceGroupNameOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v *BudgetByResourceGroupName) pulumi.StringOutput { return v.Category }).(pulumi.StringOutput)
}

func (o BudgetByResourceGroupNameOutput) CurrentSpend() CurrentSpendResponseOutput {
	return o.ApplyT(func(v *BudgetByResourceGroupName) CurrentSpendResponseOutput { return v.CurrentSpend }).(CurrentSpendResponseOutput)
}

func (o BudgetByResourceGroupNameOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BudgetByResourceGroupName) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o BudgetByResourceGroupNameOutput) Filters() FiltersResponsePtrOutput {
	return o.ApplyT(func(v *BudgetByResourceGroupName) FiltersResponsePtrOutput { return v.Filters }).(FiltersResponsePtrOutput)
}

func (o BudgetByResourceGroupNameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BudgetByResourceGroupName) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BudgetByResourceGroupNameOutput) Notifications() NotificationResponseMapOutput {
	return o.ApplyT(func(v *BudgetByResourceGroupName) NotificationResponseMapOutput { return v.Notifications }).(NotificationResponseMapOutput)
}

func (o BudgetByResourceGroupNameOutput) TimeGrain() pulumi.StringOutput {
	return o.ApplyT(func(v *BudgetByResourceGroupName) pulumi.StringOutput { return v.TimeGrain }).(pulumi.StringOutput)
}

func (o BudgetByResourceGroupNameOutput) TimePeriod() BudgetTimePeriodResponseOutput {
	return o.ApplyT(func(v *BudgetByResourceGroupName) BudgetTimePeriodResponseOutput { return v.TimePeriod }).(BudgetTimePeriodResponseOutput)
}

func (o BudgetByResourceGroupNameOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BudgetByResourceGroupName) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BudgetByResourceGroupNameOutput{})
}

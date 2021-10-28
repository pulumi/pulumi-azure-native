


package v20180131

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
			Type: pulumi.String("azure-nextgen:consumption/v20180131:BudgetByResourceGroupName"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20180331:BudgetByResourceGroupName"),
		},
		{
			Type: pulumi.String("azure-nextgen:consumption/v20180331:BudgetByResourceGroupName"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20180630:BudgetByResourceGroupName"),
		},
		{
			Type: pulumi.String("azure-nextgen:consumption/v20180630:BudgetByResourceGroupName"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20180831:BudgetByResourceGroupName"),
		},
		{
			Type: pulumi.String("azure-nextgen:consumption/v20180831:BudgetByResourceGroupName"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20181001:BudgetByResourceGroupName"),
		},
		{
			Type: pulumi.String("azure-nextgen:consumption/v20181001:BudgetByResourceGroupName"),
		},
	})
	opts = append(opts, aliases)
	var resource BudgetByResourceGroupName
	err := ctx.RegisterResource("azure-native:consumption/v20180131:BudgetByResourceGroupName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBudgetByResourceGroupName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BudgetByResourceGroupNameState, opts ...pulumi.ResourceOption) (*BudgetByResourceGroupName, error) {
	var resource BudgetByResourceGroupName
	err := ctx.ReadResource("azure-native:consumption/v20180131:BudgetByResourceGroupName", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*BudgetByResourceGroupName)(nil))
}

func (i *BudgetByResourceGroupName) ToBudgetByResourceGroupNameOutput() BudgetByResourceGroupNameOutput {
	return i.ToBudgetByResourceGroupNameOutputWithContext(context.Background())
}

func (i *BudgetByResourceGroupName) ToBudgetByResourceGroupNameOutputWithContext(ctx context.Context) BudgetByResourceGroupNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetByResourceGroupNameOutput)
}

type BudgetByResourceGroupNameOutput struct{ *pulumi.OutputState }

func (BudgetByResourceGroupNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetByResourceGroupName)(nil))
}

func (o BudgetByResourceGroupNameOutput) ToBudgetByResourceGroupNameOutput() BudgetByResourceGroupNameOutput {
	return o
}

func (o BudgetByResourceGroupNameOutput) ToBudgetByResourceGroupNameOutputWithContext(ctx context.Context) BudgetByResourceGroupNameOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BudgetByResourceGroupNameInput)(nil)).Elem(), &BudgetByResourceGroupName{})
	pulumi.RegisterOutputType(BudgetByResourceGroupNameOutput{})
}

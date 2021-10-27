


package v20180331

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Budget struct {
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


func NewBudget(ctx *pulumi.Context,
	name string, args *BudgetArgs, opts ...pulumi.ResourceOption) (*Budget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Amount == nil {
		return nil, errors.New("invalid value for required argument 'Amount'")
	}
	if args.Category == nil {
		return nil, errors.New("invalid value for required argument 'Category'")
	}
	if args.TimeGrain == nil {
		return nil, errors.New("invalid value for required argument 'TimeGrain'")
	}
	if args.TimePeriod == nil {
		return nil, errors.New("invalid value for required argument 'TimePeriod'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:consumption/v20180331:Budget"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20171230preview:Budget"),
		},
		{
			Type: pulumi.String("azure-nextgen:consumption/v20171230preview:Budget"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20180131:Budget"),
		},
		{
			Type: pulumi.String("azure-nextgen:consumption/v20180131:Budget"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20180630:Budget"),
		},
		{
			Type: pulumi.String("azure-nextgen:consumption/v20180630:Budget"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20180831:Budget"),
		},
		{
			Type: pulumi.String("azure-nextgen:consumption/v20180831:Budget"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20181001:Budget"),
		},
		{
			Type: pulumi.String("azure-nextgen:consumption/v20181001:Budget"),
		},
	})
	opts = append(opts, aliases)
	var resource Budget
	err := ctx.RegisterResource("azure-native:consumption/v20180331:Budget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBudget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BudgetState, opts ...pulumi.ResourceOption) (*Budget, error) {
	var resource Budget
	err := ctx.ReadResource("azure-native:consumption/v20180331:Budget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type budgetState struct {
}

type BudgetState struct {
}

func (BudgetState) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetState)(nil)).Elem()
}

type budgetArgs struct {
	Amount        float64                 `pulumi:"amount"`
	BudgetName    *string                 `pulumi:"budgetName"`
	Category      string                  `pulumi:"category"`
	ETag          *string                 `pulumi:"eTag"`
	Filters       *Filters                `pulumi:"filters"`
	Notifications map[string]Notification `pulumi:"notifications"`
	TimeGrain     string                  `pulumi:"timeGrain"`
	TimePeriod    BudgetTimePeriod        `pulumi:"timePeriod"`
}


type BudgetArgs struct {
	Amount        pulumi.Float64Input
	BudgetName    pulumi.StringPtrInput
	Category      pulumi.StringInput
	ETag          pulumi.StringPtrInput
	Filters       FiltersPtrInput
	Notifications NotificationMapInput
	TimeGrain     pulumi.StringInput
	TimePeriod    BudgetTimePeriodInput
}

func (BudgetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetArgs)(nil)).Elem()
}

type BudgetInput interface {
	pulumi.Input

	ToBudgetOutput() BudgetOutput
	ToBudgetOutputWithContext(ctx context.Context) BudgetOutput
}

func (*Budget) ElementType() reflect.Type {
	return reflect.TypeOf((*Budget)(nil))
}

func (i *Budget) ToBudgetOutput() BudgetOutput {
	return i.ToBudgetOutputWithContext(context.Background())
}

func (i *Budget) ToBudgetOutputWithContext(ctx context.Context) BudgetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetOutput)
}

type BudgetOutput struct{ *pulumi.OutputState }

func (BudgetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Budget)(nil))
}

func (o BudgetOutput) ToBudgetOutput() BudgetOutput {
	return o
}

func (o BudgetOutput) ToBudgetOutputWithContext(ctx context.Context) BudgetOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BudgetInput)(nil)).Elem(), &Budget{})
	pulumi.RegisterOutputType(BudgetOutput{})
}

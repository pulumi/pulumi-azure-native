


package v20201101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SnapshotPolicy struct {
	pulumi.CustomResourceState

	DailySchedule     DailyScheduleResponsePtrOutput   `pulumi:"dailySchedule"`
	Enabled           pulumi.BoolPtrOutput             `pulumi:"enabled"`
	HourlySchedule    HourlyScheduleResponsePtrOutput  `pulumi:"hourlySchedule"`
	Location          pulumi.StringOutput              `pulumi:"location"`
	MonthlySchedule   MonthlyScheduleResponsePtrOutput `pulumi:"monthlySchedule"`
	Name              pulumi.StringOutput              `pulumi:"name"`
	ProvisioningState pulumi.StringOutput              `pulumi:"provisioningState"`
	Tags              pulumi.StringMapOutput           `pulumi:"tags"`
	Type              pulumi.StringOutput              `pulumi:"type"`
	WeeklySchedule    WeeklyScheduleResponsePtrOutput  `pulumi:"weeklySchedule"`
}


func NewSnapshotPolicy(ctx *pulumi.Context,
	name string, args *SnapshotPolicyArgs, opts ...pulumi.ResourceOption) (*SnapshotPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:netapp/v20201101:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200501:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200501:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200601:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200601:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200701:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200701:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200801:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200801:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200901:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200901:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201201:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20201201:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210201:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20210201:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20210401:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401preview:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20210401preview:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210601:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20210601:SnapshotPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource SnapshotPolicy
	err := ctx.RegisterResource("azure-native:netapp/v20201101:SnapshotPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSnapshotPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotPolicyState, opts ...pulumi.ResourceOption) (*SnapshotPolicy, error) {
	var resource SnapshotPolicy
	err := ctx.ReadResource("azure-native:netapp/v20201101:SnapshotPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type snapshotPolicyState struct {
}

type SnapshotPolicyState struct {
}

func (SnapshotPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotPolicyState)(nil)).Elem()
}

type snapshotPolicyArgs struct {
	AccountName        string            `pulumi:"accountName"`
	DailySchedule      *DailySchedule    `pulumi:"dailySchedule"`
	Enabled            *bool             `pulumi:"enabled"`
	HourlySchedule     *HourlySchedule   `pulumi:"hourlySchedule"`
	Location           *string           `pulumi:"location"`
	MonthlySchedule    *MonthlySchedule  `pulumi:"monthlySchedule"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	SnapshotPolicyName *string           `pulumi:"snapshotPolicyName"`
	Tags               map[string]string `pulumi:"tags"`
	WeeklySchedule     *WeeklySchedule   `pulumi:"weeklySchedule"`
}


type SnapshotPolicyArgs struct {
	AccountName        pulumi.StringInput
	DailySchedule      DailySchedulePtrInput
	Enabled            pulumi.BoolPtrInput
	HourlySchedule     HourlySchedulePtrInput
	Location           pulumi.StringPtrInput
	MonthlySchedule    MonthlySchedulePtrInput
	ResourceGroupName  pulumi.StringInput
	SnapshotPolicyName pulumi.StringPtrInput
	Tags               pulumi.StringMapInput
	WeeklySchedule     WeeklySchedulePtrInput
}

func (SnapshotPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotPolicyArgs)(nil)).Elem()
}

type SnapshotPolicyInput interface {
	pulumi.Input

	ToSnapshotPolicyOutput() SnapshotPolicyOutput
	ToSnapshotPolicyOutputWithContext(ctx context.Context) SnapshotPolicyOutput
}

func (*SnapshotPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotPolicy)(nil))
}

func (i *SnapshotPolicy) ToSnapshotPolicyOutput() SnapshotPolicyOutput {
	return i.ToSnapshotPolicyOutputWithContext(context.Background())
}

func (i *SnapshotPolicy) ToSnapshotPolicyOutputWithContext(ctx context.Context) SnapshotPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotPolicyOutput)
}

type SnapshotPolicyOutput struct{ *pulumi.OutputState }

func (SnapshotPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotPolicy)(nil))
}

func (o SnapshotPolicyOutput) ToSnapshotPolicyOutput() SnapshotPolicyOutput {
	return o
}

func (o SnapshotPolicyOutput) ToSnapshotPolicyOutputWithContext(ctx context.Context) SnapshotPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotPolicyInput)(nil)).Elem(), &SnapshotPolicy{})
	pulumi.RegisterOutputType(SnapshotPolicyOutput{})
}

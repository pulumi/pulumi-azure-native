


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedInstanceLongTermRetentionPolicy struct {
	pulumi.CustomResourceState

	MonthlyRetention pulumi.StringPtrOutput `pulumi:"monthlyRetention"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	Type             pulumi.StringOutput    `pulumi:"type"`
	WeekOfYear       pulumi.IntPtrOutput    `pulumi:"weekOfYear"`
	WeeklyRetention  pulumi.StringPtrOutput `pulumi:"weeklyRetention"`
	YearlyRetention  pulumi.StringPtrOutput `pulumi:"yearlyRetention"`
}


func NewManagedInstanceLongTermRetentionPolicy(ctx *pulumi.Context,
	name string, args *ManagedInstanceLongTermRetentionPolicyArgs, opts ...pulumi.ResourceOption) (*ManagedInstanceLongTermRetentionPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ManagedInstanceLongTermRetentionPolicy
	err := ctx.RegisterResource("azure-native:sql/v20220501preview:ManagedInstanceLongTermRetentionPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedInstanceLongTermRetentionPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedInstanceLongTermRetentionPolicyState, opts ...pulumi.ResourceOption) (*ManagedInstanceLongTermRetentionPolicy, error) {
	var resource ManagedInstanceLongTermRetentionPolicy
	err := ctx.ReadResource("azure-native:sql/v20220501preview:ManagedInstanceLongTermRetentionPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedInstanceLongTermRetentionPolicyState struct {
}

type ManagedInstanceLongTermRetentionPolicyState struct {
}

func (ManagedInstanceLongTermRetentionPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceLongTermRetentionPolicyState)(nil)).Elem()
}

type managedInstanceLongTermRetentionPolicyArgs struct {
	DatabaseName        string  `pulumi:"databaseName"`
	ManagedInstanceName string  `pulumi:"managedInstanceName"`
	MonthlyRetention    *string `pulumi:"monthlyRetention"`
	PolicyName          *string `pulumi:"policyName"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
	WeekOfYear          *int    `pulumi:"weekOfYear"`
	WeeklyRetention     *string `pulumi:"weeklyRetention"`
	YearlyRetention     *string `pulumi:"yearlyRetention"`
}


type ManagedInstanceLongTermRetentionPolicyArgs struct {
	DatabaseName        pulumi.StringInput
	ManagedInstanceName pulumi.StringInput
	MonthlyRetention    pulumi.StringPtrInput
	PolicyName          pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	WeekOfYear          pulumi.IntPtrInput
	WeeklyRetention     pulumi.StringPtrInput
	YearlyRetention     pulumi.StringPtrInput
}

func (ManagedInstanceLongTermRetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceLongTermRetentionPolicyArgs)(nil)).Elem()
}

type ManagedInstanceLongTermRetentionPolicyInput interface {
	pulumi.Input

	ToManagedInstanceLongTermRetentionPolicyOutput() ManagedInstanceLongTermRetentionPolicyOutput
	ToManagedInstanceLongTermRetentionPolicyOutputWithContext(ctx context.Context) ManagedInstanceLongTermRetentionPolicyOutput
}

func (*ManagedInstanceLongTermRetentionPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceLongTermRetentionPolicy)(nil)).Elem()
}

func (i *ManagedInstanceLongTermRetentionPolicy) ToManagedInstanceLongTermRetentionPolicyOutput() ManagedInstanceLongTermRetentionPolicyOutput {
	return i.ToManagedInstanceLongTermRetentionPolicyOutputWithContext(context.Background())
}

func (i *ManagedInstanceLongTermRetentionPolicy) ToManagedInstanceLongTermRetentionPolicyOutputWithContext(ctx context.Context) ManagedInstanceLongTermRetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceLongTermRetentionPolicyOutput)
}

type ManagedInstanceLongTermRetentionPolicyOutput struct{ *pulumi.OutputState }

func (ManagedInstanceLongTermRetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceLongTermRetentionPolicy)(nil)).Elem()
}

func (o ManagedInstanceLongTermRetentionPolicyOutput) ToManagedInstanceLongTermRetentionPolicyOutput() ManagedInstanceLongTermRetentionPolicyOutput {
	return o
}

func (o ManagedInstanceLongTermRetentionPolicyOutput) ToManagedInstanceLongTermRetentionPolicyOutputWithContext(ctx context.Context) ManagedInstanceLongTermRetentionPolicyOutput {
	return o
}

func (o ManagedInstanceLongTermRetentionPolicyOutput) MonthlyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceLongTermRetentionPolicy) pulumi.StringPtrOutput { return v.MonthlyRetention }).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceLongTermRetentionPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstanceLongTermRetentionPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedInstanceLongTermRetentionPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstanceLongTermRetentionPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedInstanceLongTermRetentionPolicyOutput) WeekOfYear() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceLongTermRetentionPolicy) pulumi.IntPtrOutput { return v.WeekOfYear }).(pulumi.IntPtrOutput)
}

func (o ManagedInstanceLongTermRetentionPolicyOutput) WeeklyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceLongTermRetentionPolicy) pulumi.StringPtrOutput { return v.WeeklyRetention }).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceLongTermRetentionPolicyOutput) YearlyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceLongTermRetentionPolicy) pulumi.StringPtrOutput { return v.YearlyRetention }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedInstanceLongTermRetentionPolicyOutput{})
}

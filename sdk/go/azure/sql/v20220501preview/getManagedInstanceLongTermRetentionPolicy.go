


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedInstanceLongTermRetentionPolicy(ctx *pulumi.Context, args *LookupManagedInstanceLongTermRetentionPolicyArgs, opts ...pulumi.InvokeOption) (*LookupManagedInstanceLongTermRetentionPolicyResult, error) {
	var rv LookupManagedInstanceLongTermRetentionPolicyResult
	err := ctx.Invoke("azure-native:sql/v20220501preview:getManagedInstanceLongTermRetentionPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedInstanceLongTermRetentionPolicyArgs struct {
	DatabaseName        string `pulumi:"databaseName"`
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	PolicyName          string `pulumi:"policyName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupManagedInstanceLongTermRetentionPolicyResult struct {
	Id               string  `pulumi:"id"`
	MonthlyRetention *string `pulumi:"monthlyRetention"`
	Name             string  `pulumi:"name"`
	Type             string  `pulumi:"type"`
	WeekOfYear       *int    `pulumi:"weekOfYear"`
	WeeklyRetention  *string `pulumi:"weeklyRetention"`
	YearlyRetention  *string `pulumi:"yearlyRetention"`
}

func LookupManagedInstanceLongTermRetentionPolicyOutput(ctx *pulumi.Context, args LookupManagedInstanceLongTermRetentionPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupManagedInstanceLongTermRetentionPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedInstanceLongTermRetentionPolicyResult, error) {
			args := v.(LookupManagedInstanceLongTermRetentionPolicyArgs)
			r, err := LookupManagedInstanceLongTermRetentionPolicy(ctx, &args, opts...)
			var s LookupManagedInstanceLongTermRetentionPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedInstanceLongTermRetentionPolicyResultOutput)
}

type LookupManagedInstanceLongTermRetentionPolicyOutputArgs struct {
	DatabaseName        pulumi.StringInput `pulumi:"databaseName"`
	ManagedInstanceName pulumi.StringInput `pulumi:"managedInstanceName"`
	PolicyName          pulumi.StringInput `pulumi:"policyName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedInstanceLongTermRetentionPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstanceLongTermRetentionPolicyArgs)(nil)).Elem()
}


type LookupManagedInstanceLongTermRetentionPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupManagedInstanceLongTermRetentionPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstanceLongTermRetentionPolicyResult)(nil)).Elem()
}

func (o LookupManagedInstanceLongTermRetentionPolicyResultOutput) ToLookupManagedInstanceLongTermRetentionPolicyResultOutput() LookupManagedInstanceLongTermRetentionPolicyResultOutput {
	return o
}

func (o LookupManagedInstanceLongTermRetentionPolicyResultOutput) ToLookupManagedInstanceLongTermRetentionPolicyResultOutputWithContext(ctx context.Context) LookupManagedInstanceLongTermRetentionPolicyResultOutput {
	return o
}

func (o LookupManagedInstanceLongTermRetentionPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceLongTermRetentionPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceLongTermRetentionPolicyResultOutput) MonthlyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceLongTermRetentionPolicyResult) *string { return v.MonthlyRetention }).(pulumi.StringPtrOutput)
}

func (o LookupManagedInstanceLongTermRetentionPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceLongTermRetentionPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceLongTermRetentionPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceLongTermRetentionPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceLongTermRetentionPolicyResultOutput) WeekOfYear() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceLongTermRetentionPolicyResult) *int { return v.WeekOfYear }).(pulumi.IntPtrOutput)
}

func (o LookupManagedInstanceLongTermRetentionPolicyResultOutput) WeeklyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceLongTermRetentionPolicyResult) *string { return v.WeeklyRetention }).(pulumi.StringPtrOutput)
}

func (o LookupManagedInstanceLongTermRetentionPolicyResultOutput) YearlyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceLongTermRetentionPolicyResult) *string { return v.YearlyRetention }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedInstanceLongTermRetentionPolicyResultOutput{})
}

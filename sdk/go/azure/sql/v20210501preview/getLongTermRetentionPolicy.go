


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLongTermRetentionPolicy(ctx *pulumi.Context, args *LookupLongTermRetentionPolicyArgs, opts ...pulumi.InvokeOption) (*LookupLongTermRetentionPolicyResult, error) {
	var rv LookupLongTermRetentionPolicyResult
	err := ctx.Invoke("azure-native:sql/v20210501preview:getLongTermRetentionPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLongTermRetentionPolicyArgs struct {
	DatabaseName      string `pulumi:"databaseName"`
	PolicyName        string `pulumi:"policyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupLongTermRetentionPolicyResult struct {
	Id               string  `pulumi:"id"`
	MonthlyRetention *string `pulumi:"monthlyRetention"`
	Name             string  `pulumi:"name"`
	Type             string  `pulumi:"type"`
	WeekOfYear       *int    `pulumi:"weekOfYear"`
	WeeklyRetention  *string `pulumi:"weeklyRetention"`
	YearlyRetention  *string `pulumi:"yearlyRetention"`
}

func LookupLongTermRetentionPolicyOutput(ctx *pulumi.Context, args LookupLongTermRetentionPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupLongTermRetentionPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLongTermRetentionPolicyResult, error) {
			args := v.(LookupLongTermRetentionPolicyArgs)
			r, err := LookupLongTermRetentionPolicy(ctx, &args, opts...)
			var s LookupLongTermRetentionPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLongTermRetentionPolicyResultOutput)
}

type LookupLongTermRetentionPolicyOutputArgs struct {
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	PolicyName        pulumi.StringInput `pulumi:"policyName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupLongTermRetentionPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLongTermRetentionPolicyArgs)(nil)).Elem()
}


type LookupLongTermRetentionPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupLongTermRetentionPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLongTermRetentionPolicyResult)(nil)).Elem()
}

func (o LookupLongTermRetentionPolicyResultOutput) ToLookupLongTermRetentionPolicyResultOutput() LookupLongTermRetentionPolicyResultOutput {
	return o
}

func (o LookupLongTermRetentionPolicyResultOutput) ToLookupLongTermRetentionPolicyResultOutputWithContext(ctx context.Context) LookupLongTermRetentionPolicyResultOutput {
	return o
}

func (o LookupLongTermRetentionPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLongTermRetentionPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLongTermRetentionPolicyResultOutput) MonthlyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLongTermRetentionPolicyResult) *string { return v.MonthlyRetention }).(pulumi.StringPtrOutput)
}

func (o LookupLongTermRetentionPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLongTermRetentionPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLongTermRetentionPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLongTermRetentionPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupLongTermRetentionPolicyResultOutput) WeekOfYear() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupLongTermRetentionPolicyResult) *int { return v.WeekOfYear }).(pulumi.IntPtrOutput)
}

func (o LookupLongTermRetentionPolicyResultOutput) WeeklyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLongTermRetentionPolicyResult) *string { return v.WeeklyRetention }).(pulumi.StringPtrOutput)
}

func (o LookupLongTermRetentionPolicyResultOutput) YearlyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLongTermRetentionPolicyResult) *string { return v.YearlyRetention }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLongTermRetentionPolicyResultOutput{})
}

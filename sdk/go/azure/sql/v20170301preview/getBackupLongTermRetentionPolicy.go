


package v20170301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackupLongTermRetentionPolicy(ctx *pulumi.Context, args *LookupBackupLongTermRetentionPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBackupLongTermRetentionPolicyResult, error) {
	var rv LookupBackupLongTermRetentionPolicyResult
	err := ctx.Invoke("azure-native:sql/v20170301preview:getBackupLongTermRetentionPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupLongTermRetentionPolicyArgs struct {
	DatabaseName      string `pulumi:"databaseName"`
	PolicyName        string `pulumi:"policyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupBackupLongTermRetentionPolicyResult struct {
	Id               string  `pulumi:"id"`
	MonthlyRetention *string `pulumi:"monthlyRetention"`
	Name             string  `pulumi:"name"`
	Type             string  `pulumi:"type"`
	WeekOfYear       *int    `pulumi:"weekOfYear"`
	WeeklyRetention  *string `pulumi:"weeklyRetention"`
	YearlyRetention  *string `pulumi:"yearlyRetention"`
}

func LookupBackupLongTermRetentionPolicyOutput(ctx *pulumi.Context, args LookupBackupLongTermRetentionPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupBackupLongTermRetentionPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackupLongTermRetentionPolicyResult, error) {
			args := v.(LookupBackupLongTermRetentionPolicyArgs)
			r, err := LookupBackupLongTermRetentionPolicy(ctx, &args, opts...)
			var s LookupBackupLongTermRetentionPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackupLongTermRetentionPolicyResultOutput)
}

type LookupBackupLongTermRetentionPolicyOutputArgs struct {
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	PolicyName        pulumi.StringInput `pulumi:"policyName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupBackupLongTermRetentionPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupLongTermRetentionPolicyArgs)(nil)).Elem()
}


type LookupBackupLongTermRetentionPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupBackupLongTermRetentionPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupLongTermRetentionPolicyResult)(nil)).Elem()
}

func (o LookupBackupLongTermRetentionPolicyResultOutput) ToLookupBackupLongTermRetentionPolicyResultOutput() LookupBackupLongTermRetentionPolicyResultOutput {
	return o
}

func (o LookupBackupLongTermRetentionPolicyResultOutput) ToLookupBackupLongTermRetentionPolicyResultOutputWithContext(ctx context.Context) LookupBackupLongTermRetentionPolicyResultOutput {
	return o
}

func (o LookupBackupLongTermRetentionPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupLongTermRetentionPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBackupLongTermRetentionPolicyResultOutput) MonthlyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackupLongTermRetentionPolicyResult) *string { return v.MonthlyRetention }).(pulumi.StringPtrOutput)
}

func (o LookupBackupLongTermRetentionPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupLongTermRetentionPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBackupLongTermRetentionPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupLongTermRetentionPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupBackupLongTermRetentionPolicyResultOutput) WeekOfYear() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBackupLongTermRetentionPolicyResult) *int { return v.WeekOfYear }).(pulumi.IntPtrOutput)
}

func (o LookupBackupLongTermRetentionPolicyResultOutput) WeeklyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackupLongTermRetentionPolicyResult) *string { return v.WeeklyRetention }).(pulumi.StringPtrOutput)
}

func (o LookupBackupLongTermRetentionPolicyResultOutput) YearlyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackupLongTermRetentionPolicyResult) *string { return v.YearlyRetention }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupLongTermRetentionPolicyResultOutput{})
}

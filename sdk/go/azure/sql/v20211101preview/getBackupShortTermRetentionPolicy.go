


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackupShortTermRetentionPolicy(ctx *pulumi.Context, args *LookupBackupShortTermRetentionPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBackupShortTermRetentionPolicyResult, error) {
	var rv LookupBackupShortTermRetentionPolicyResult
	err := ctx.Invoke("azure-native:sql/v20211101preview:getBackupShortTermRetentionPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupShortTermRetentionPolicyArgs struct {
	DatabaseName      string `pulumi:"databaseName"`
	PolicyName        string `pulumi:"policyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupBackupShortTermRetentionPolicyResult struct {
	DiffBackupIntervalInHours *int   `pulumi:"diffBackupIntervalInHours"`
	Id                        string `pulumi:"id"`
	Name                      string `pulumi:"name"`
	RetentionDays             *int   `pulumi:"retentionDays"`
	Type                      string `pulumi:"type"`
}

func LookupBackupShortTermRetentionPolicyOutput(ctx *pulumi.Context, args LookupBackupShortTermRetentionPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupBackupShortTermRetentionPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackupShortTermRetentionPolicyResult, error) {
			args := v.(LookupBackupShortTermRetentionPolicyArgs)
			r, err := LookupBackupShortTermRetentionPolicy(ctx, &args, opts...)
			var s LookupBackupShortTermRetentionPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackupShortTermRetentionPolicyResultOutput)
}

type LookupBackupShortTermRetentionPolicyOutputArgs struct {
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	PolicyName        pulumi.StringInput `pulumi:"policyName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupBackupShortTermRetentionPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupShortTermRetentionPolicyArgs)(nil)).Elem()
}


type LookupBackupShortTermRetentionPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupBackupShortTermRetentionPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupShortTermRetentionPolicyResult)(nil)).Elem()
}

func (o LookupBackupShortTermRetentionPolicyResultOutput) ToLookupBackupShortTermRetentionPolicyResultOutput() LookupBackupShortTermRetentionPolicyResultOutput {
	return o
}

func (o LookupBackupShortTermRetentionPolicyResultOutput) ToLookupBackupShortTermRetentionPolicyResultOutputWithContext(ctx context.Context) LookupBackupShortTermRetentionPolicyResultOutput {
	return o
}

func (o LookupBackupShortTermRetentionPolicyResultOutput) DiffBackupIntervalInHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBackupShortTermRetentionPolicyResult) *int { return v.DiffBackupIntervalInHours }).(pulumi.IntPtrOutput)
}

func (o LookupBackupShortTermRetentionPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupShortTermRetentionPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBackupShortTermRetentionPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupShortTermRetentionPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBackupShortTermRetentionPolicyResultOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBackupShortTermRetentionPolicyResult) *int { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

func (o LookupBackupShortTermRetentionPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupShortTermRetentionPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupShortTermRetentionPolicyResultOutput{})
}

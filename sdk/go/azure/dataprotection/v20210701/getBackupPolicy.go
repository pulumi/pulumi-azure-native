


package v20210701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackupPolicy(ctx *pulumi.Context, args *LookupBackupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBackupPolicyResult, error) {
	var rv LookupBackupPolicyResult
	err := ctx.Invoke("azure-native:dataprotection/v20210701:getBackupPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupPolicyArgs struct {
	BackupPolicyName  string `pulumi:"backupPolicyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VaultName         string `pulumi:"vaultName"`
}


type LookupBackupPolicyResult struct {
	Id         string               `pulumi:"id"`
	Name       string               `pulumi:"name"`
	Properties BackupPolicyResponse `pulumi:"properties"`
	SystemData SystemDataResponse   `pulumi:"systemData"`
	Type       string               `pulumi:"type"`
}

func LookupBackupPolicyOutput(ctx *pulumi.Context, args LookupBackupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupBackupPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackupPolicyResult, error) {
			args := v.(LookupBackupPolicyArgs)
			r, err := LookupBackupPolicy(ctx, &args, opts...)
			var s LookupBackupPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackupPolicyResultOutput)
}

type LookupBackupPolicyOutputArgs struct {
	BackupPolicyName  pulumi.StringInput `pulumi:"backupPolicyName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VaultName         pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupBackupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupPolicyArgs)(nil)).Elem()
}


type LookupBackupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupBackupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupPolicyResult)(nil)).Elem()
}

func (o LookupBackupPolicyResultOutput) ToLookupBackupPolicyResultOutput() LookupBackupPolicyResultOutput {
	return o
}

func (o LookupBackupPolicyResultOutput) ToLookupBackupPolicyResultOutputWithContext(ctx context.Context) LookupBackupPolicyResultOutput {
	return o
}

func (o LookupBackupPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBackupPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBackupPolicyResultOutput) Properties() BackupPolicyResponseOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) BackupPolicyResponse { return v.Properties }).(BackupPolicyResponseOutput)
}

func (o LookupBackupPolicyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupBackupPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupPolicyResultOutput{})
}

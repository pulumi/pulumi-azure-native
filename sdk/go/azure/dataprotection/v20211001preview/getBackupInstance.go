


package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackupInstance(ctx *pulumi.Context, args *LookupBackupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupBackupInstanceResult, error) {
	var rv LookupBackupInstanceResult
	err := ctx.Invoke("azure-native:dataprotection/v20211001preview:getBackupInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupInstanceArgs struct {
	BackupInstanceName string `pulumi:"backupInstanceName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	VaultName          string `pulumi:"vaultName"`
}


type LookupBackupInstanceResult struct {
	Id         string                 `pulumi:"id"`
	Name       string                 `pulumi:"name"`
	Properties BackupInstanceResponse `pulumi:"properties"`
	SystemData SystemDataResponse     `pulumi:"systemData"`
	Type       string                 `pulumi:"type"`
}

func LookupBackupInstanceOutput(ctx *pulumi.Context, args LookupBackupInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupBackupInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackupInstanceResult, error) {
			args := v.(LookupBackupInstanceArgs)
			r, err := LookupBackupInstance(ctx, &args, opts...)
			var s LookupBackupInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackupInstanceResultOutput)
}

type LookupBackupInstanceOutputArgs struct {
	BackupInstanceName pulumi.StringInput `pulumi:"backupInstanceName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	VaultName          pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupBackupInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupInstanceArgs)(nil)).Elem()
}


type LookupBackupInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupBackupInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupInstanceResult)(nil)).Elem()
}

func (o LookupBackupInstanceResultOutput) ToLookupBackupInstanceResultOutput() LookupBackupInstanceResultOutput {
	return o
}

func (o LookupBackupInstanceResultOutput) ToLookupBackupInstanceResultOutputWithContext(ctx context.Context) LookupBackupInstanceResultOutput {
	return o
}

func (o LookupBackupInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBackupInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBackupInstanceResultOutput) Properties() BackupInstanceResponseOutput {
	return o.ApplyT(func(v LookupBackupInstanceResult) BackupInstanceResponse { return v.Properties }).(BackupInstanceResponseOutput)
}

func (o LookupBackupInstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBackupInstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupBackupInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupInstanceResultOutput{})
}

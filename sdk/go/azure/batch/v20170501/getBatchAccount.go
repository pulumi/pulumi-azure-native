


package v20170501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupBatchAccount(ctx *pulumi.Context, args *LookupBatchAccountArgs, opts ...pulumi.InvokeOption) (*LookupBatchAccountResult, error) {
	var rv LookupBatchAccountResult
	err := ctx.Invoke("azure-native:batch/v20170501:getBatchAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBatchAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupBatchAccountResult struct {
	AccountEndpoint              string                        `pulumi:"accountEndpoint"`
	ActiveJobAndJobScheduleQuota int                           `pulumi:"activeJobAndJobScheduleQuota"`
	AutoStorage                  AutoStoragePropertiesResponse `pulumi:"autoStorage"`
	DedicatedCoreQuota           int                           `pulumi:"dedicatedCoreQuota"`
	Id                           string                        `pulumi:"id"`
	KeyVaultReference            KeyVaultReferenceResponse     `pulumi:"keyVaultReference"`
	Location                     string                        `pulumi:"location"`
	LowPriorityCoreQuota         int                           `pulumi:"lowPriorityCoreQuota"`
	Name                         string                        `pulumi:"name"`
	PoolAllocationMode           string                        `pulumi:"poolAllocationMode"`
	PoolQuota                    int                           `pulumi:"poolQuota"`
	ProvisioningState            string                        `pulumi:"provisioningState"`
	Tags                         map[string]string             `pulumi:"tags"`
	Type                         string                        `pulumi:"type"`
}

func LookupBatchAccountOutput(ctx *pulumi.Context, args LookupBatchAccountOutputArgs, opts ...pulumi.InvokeOption) LookupBatchAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBatchAccountResult, error) {
			args := v.(LookupBatchAccountArgs)
			r, err := LookupBatchAccount(ctx, &args, opts...)
			var s LookupBatchAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBatchAccountResultOutput)
}

type LookupBatchAccountOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBatchAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBatchAccountArgs)(nil)).Elem()
}


type LookupBatchAccountResultOutput struct{ *pulumi.OutputState }

func (LookupBatchAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBatchAccountResult)(nil)).Elem()
}

func (o LookupBatchAccountResultOutput) ToLookupBatchAccountResultOutput() LookupBatchAccountResultOutput {
	return o
}

func (o LookupBatchAccountResultOutput) ToLookupBatchAccountResultOutputWithContext(ctx context.Context) LookupBatchAccountResultOutput {
	return o
}

func (o LookupBatchAccountResultOutput) AccountEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) string { return v.AccountEndpoint }).(pulumi.StringOutput)
}

func (o LookupBatchAccountResultOutput) ActiveJobAndJobScheduleQuota() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) int { return v.ActiveJobAndJobScheduleQuota }).(pulumi.IntOutput)
}

func (o LookupBatchAccountResultOutput) AutoStorage() AutoStoragePropertiesResponseOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) AutoStoragePropertiesResponse { return v.AutoStorage }).(AutoStoragePropertiesResponseOutput)
}

func (o LookupBatchAccountResultOutput) DedicatedCoreQuota() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) int { return v.DedicatedCoreQuota }).(pulumi.IntOutput)
}

func (o LookupBatchAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBatchAccountResultOutput) KeyVaultReference() KeyVaultReferenceResponseOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) KeyVaultReferenceResponse { return v.KeyVaultReference }).(KeyVaultReferenceResponseOutput)
}

func (o LookupBatchAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupBatchAccountResultOutput) LowPriorityCoreQuota() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) int { return v.LowPriorityCoreQuota }).(pulumi.IntOutput)
}

func (o LookupBatchAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBatchAccountResultOutput) PoolAllocationMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) string { return v.PoolAllocationMode }).(pulumi.StringOutput)
}

func (o LookupBatchAccountResultOutput) PoolQuota() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) int { return v.PoolQuota }).(pulumi.IntOutput)
}

func (o LookupBatchAccountResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupBatchAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupBatchAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBatchAccountResultOutput{})
}

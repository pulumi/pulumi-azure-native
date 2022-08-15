


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlPool(ctx *pulumi.Context, args *LookupSqlPoolArgs, opts ...pulumi.InvokeOption) (*LookupSqlPoolResult, error) {
	var rv LookupSqlPoolResult
	err := ctx.Invoke("azure-native:synapse/v20210301:getSqlPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlPoolArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SqlPoolName       string `pulumi:"sqlPoolName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupSqlPoolResult struct {
	Collation             *string           `pulumi:"collation"`
	CreateMode            *string           `pulumi:"createMode"`
	CreationDate          *string           `pulumi:"creationDate"`
	Id                    string            `pulumi:"id"`
	Location              string            `pulumi:"location"`
	MaxSizeBytes          *float64          `pulumi:"maxSizeBytes"`
	Name                  string            `pulumi:"name"`
	ProvisioningState     *string           `pulumi:"provisioningState"`
	RecoverableDatabaseId *string           `pulumi:"recoverableDatabaseId"`
	RestorePointInTime    *string           `pulumi:"restorePointInTime"`
	Sku                   *SkuResponse      `pulumi:"sku"`
	SourceDatabaseId      *string           `pulumi:"sourceDatabaseId"`
	Status                *string           `pulumi:"status"`
	StorageAccountType    *string           `pulumi:"storageAccountType"`
	Tags                  map[string]string `pulumi:"tags"`
	Type                  string            `pulumi:"type"`
}

func LookupSqlPoolOutput(ctx *pulumi.Context, args LookupSqlPoolOutputArgs, opts ...pulumi.InvokeOption) LookupSqlPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlPoolResult, error) {
			args := v.(LookupSqlPoolArgs)
			r, err := LookupSqlPool(ctx, &args, opts...)
			var s LookupSqlPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlPoolResultOutput)
}

type LookupSqlPoolOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SqlPoolName       pulumi.StringInput `pulumi:"sqlPoolName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupSqlPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolArgs)(nil)).Elem()
}


type LookupSqlPoolResultOutput struct{ *pulumi.OutputState }

func (LookupSqlPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolResult)(nil)).Elem()
}

func (o LookupSqlPoolResultOutput) ToLookupSqlPoolResultOutput() LookupSqlPoolResultOutput {
	return o
}

func (o LookupSqlPoolResultOutput) ToLookupSqlPoolResultOutputWithContext(ctx context.Context) LookupSqlPoolResultOutput {
	return o
}

func (o LookupSqlPoolResultOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) *string { return v.Collation }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolResultOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolResultOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) *string { return v.CreationDate }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSqlPoolResultOutput) MaxSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) *float64 { return v.MaxSizeBytes }).(pulumi.Float64PtrOutput)
}

func (o LookupSqlPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlPoolResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolResultOutput) RecoverableDatabaseId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) *string { return v.RecoverableDatabaseId }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolResultOutput) RestorePointInTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) *string { return v.RestorePointInTime }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupSqlPoolResultOutput) SourceDatabaseId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) *string { return v.SourceDatabaseId }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolResultOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSqlPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlPoolResultOutput{})
}

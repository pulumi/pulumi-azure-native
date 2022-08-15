


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobStorageAccountDataSet(ctx *pulumi.Context, args *LookupBlobStorageAccountDataSetArgs, opts ...pulumi.InvokeOption) (*LookupBlobStorageAccountDataSetResult, error) {
	var rv LookupBlobStorageAccountDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20201001preview:getBlobStorageAccountDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobStorageAccountDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupBlobStorageAccountDataSetResult struct {
	DataSetId                string                           `pulumi:"dataSetId"`
	Id                       string                           `pulumi:"id"`
	Kind                     string                           `pulumi:"kind"`
	Location                 string                           `pulumi:"location"`
	Name                     string                           `pulumi:"name"`
	Paths                    []BlobStorageAccountPathResponse `pulumi:"paths"`
	StorageAccountResourceId string                           `pulumi:"storageAccountResourceId"`
	SystemData               SystemDataResponse               `pulumi:"systemData"`
	Type                     string                           `pulumi:"type"`
}

func LookupBlobStorageAccountDataSetOutput(ctx *pulumi.Context, args LookupBlobStorageAccountDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupBlobStorageAccountDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlobStorageAccountDataSetResult, error) {
			args := v.(LookupBlobStorageAccountDataSetArgs)
			r, err := LookupBlobStorageAccountDataSet(ctx, &args, opts...)
			var s LookupBlobStorageAccountDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlobStorageAccountDataSetResultOutput)
}

type LookupBlobStorageAccountDataSetOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DataSetName       pulumi.StringInput `pulumi:"dataSetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupBlobStorageAccountDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobStorageAccountDataSetArgs)(nil)).Elem()
}


type LookupBlobStorageAccountDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupBlobStorageAccountDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobStorageAccountDataSetResult)(nil)).Elem()
}

func (o LookupBlobStorageAccountDataSetResultOutput) ToLookupBlobStorageAccountDataSetResultOutput() LookupBlobStorageAccountDataSetResultOutput {
	return o
}

func (o LookupBlobStorageAccountDataSetResultOutput) ToLookupBlobStorageAccountDataSetResultOutputWithContext(ctx context.Context) LookupBlobStorageAccountDataSetResultOutput {
	return o
}

func (o LookupBlobStorageAccountDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupBlobStorageAccountDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBlobStorageAccountDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupBlobStorageAccountDataSetResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupBlobStorageAccountDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBlobStorageAccountDataSetResultOutput) Paths() BlobStorageAccountPathResponseArrayOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetResult) []BlobStorageAccountPathResponse { return v.Paths }).(BlobStorageAccountPathResponseArrayOutput)
}

func (o LookupBlobStorageAccountDataSetResultOutput) StorageAccountResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetResult) string { return v.StorageAccountResourceId }).(pulumi.StringOutput)
}

func (o LookupBlobStorageAccountDataSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupBlobStorageAccountDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlobStorageAccountDataSetResultOutput{})
}




package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupADLSGen2StorageAccountDataSet(ctx *pulumi.Context, args *LookupADLSGen2StorageAccountDataSetArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen2StorageAccountDataSetResult, error) {
	var rv LookupADLSGen2StorageAccountDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20201001preview:getADLSGen2StorageAccountDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen2StorageAccountDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupADLSGen2StorageAccountDataSetResult struct {
	DataSetId                string                               `pulumi:"dataSetId"`
	Id                       string                               `pulumi:"id"`
	Kind                     string                               `pulumi:"kind"`
	Location                 string                               `pulumi:"location"`
	Name                     string                               `pulumi:"name"`
	Paths                    []ADLSGen2StorageAccountPathResponse `pulumi:"paths"`
	StorageAccountResourceId string                               `pulumi:"storageAccountResourceId"`
	SystemData               SystemDataResponse                   `pulumi:"systemData"`
	Type                     string                               `pulumi:"type"`
}

func LookupADLSGen2StorageAccountDataSetOutput(ctx *pulumi.Context, args LookupADLSGen2StorageAccountDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupADLSGen2StorageAccountDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupADLSGen2StorageAccountDataSetResult, error) {
			args := v.(LookupADLSGen2StorageAccountDataSetArgs)
			r, err := LookupADLSGen2StorageAccountDataSet(ctx, &args, opts...)
			var s LookupADLSGen2StorageAccountDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupADLSGen2StorageAccountDataSetResultOutput)
}

type LookupADLSGen2StorageAccountDataSetOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DataSetName       pulumi.StringInput `pulumi:"dataSetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupADLSGen2StorageAccountDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen2StorageAccountDataSetArgs)(nil)).Elem()
}


type LookupADLSGen2StorageAccountDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupADLSGen2StorageAccountDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen2StorageAccountDataSetResult)(nil)).Elem()
}

func (o LookupADLSGen2StorageAccountDataSetResultOutput) ToLookupADLSGen2StorageAccountDataSetResultOutput() LookupADLSGen2StorageAccountDataSetResultOutput {
	return o
}

func (o LookupADLSGen2StorageAccountDataSetResultOutput) ToLookupADLSGen2StorageAccountDataSetResultOutputWithContext(ctx context.Context) LookupADLSGen2StorageAccountDataSetResultOutput {
	return o
}

func (o LookupADLSGen2StorageAccountDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupADLSGen2StorageAccountDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupADLSGen2StorageAccountDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupADLSGen2StorageAccountDataSetResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupADLSGen2StorageAccountDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupADLSGen2StorageAccountDataSetResultOutput) Paths() ADLSGen2StorageAccountPathResponseArrayOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetResult) []ADLSGen2StorageAccountPathResponse { return v.Paths }).(ADLSGen2StorageAccountPathResponseArrayOutput)
}

func (o LookupADLSGen2StorageAccountDataSetResultOutput) StorageAccountResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetResult) string { return v.StorageAccountResourceId }).(pulumi.StringOutput)
}

func (o LookupADLSGen2StorageAccountDataSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupADLSGen2StorageAccountDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupADLSGen2StorageAccountDataSetResultOutput{})
}

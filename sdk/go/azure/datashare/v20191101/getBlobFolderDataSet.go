


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobFolderDataSet(ctx *pulumi.Context, args *LookupBlobFolderDataSetArgs, opts ...pulumi.InvokeOption) (*LookupBlobFolderDataSetResult, error) {
	var rv LookupBlobFolderDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getBlobFolderDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobFolderDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupBlobFolderDataSetResult struct {
	ContainerName      string `pulumi:"containerName"`
	DataSetId          string `pulumi:"dataSetId"`
	Id                 string `pulumi:"id"`
	Kind               string `pulumi:"kind"`
	Name               string `pulumi:"name"`
	Prefix             string `pulumi:"prefix"`
	ResourceGroup      string `pulumi:"resourceGroup"`
	StorageAccountName string `pulumi:"storageAccountName"`
	SubscriptionId     string `pulumi:"subscriptionId"`
	Type               string `pulumi:"type"`
}

func LookupBlobFolderDataSetOutput(ctx *pulumi.Context, args LookupBlobFolderDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupBlobFolderDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlobFolderDataSetResult, error) {
			args := v.(LookupBlobFolderDataSetArgs)
			r, err := LookupBlobFolderDataSet(ctx, &args, opts...)
			var s LookupBlobFolderDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlobFolderDataSetResultOutput)
}

type LookupBlobFolderDataSetOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DataSetName       pulumi.StringInput `pulumi:"dataSetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupBlobFolderDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobFolderDataSetArgs)(nil)).Elem()
}


type LookupBlobFolderDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupBlobFolderDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobFolderDataSetResult)(nil)).Elem()
}

func (o LookupBlobFolderDataSetResultOutput) ToLookupBlobFolderDataSetResultOutput() LookupBlobFolderDataSetResultOutput {
	return o
}

func (o LookupBlobFolderDataSetResultOutput) ToLookupBlobFolderDataSetResultOutputWithContext(ctx context.Context) LookupBlobFolderDataSetResultOutput {
	return o
}

func (o LookupBlobFolderDataSetResultOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetResult) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o LookupBlobFolderDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupBlobFolderDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBlobFolderDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupBlobFolderDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBlobFolderDataSetResultOutput) Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetResult) string { return v.Prefix }).(pulumi.StringOutput)
}

func (o LookupBlobFolderDataSetResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupBlobFolderDataSetResultOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetResult) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o LookupBlobFolderDataSetResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o LookupBlobFolderDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlobFolderDataSetResultOutput{})
}

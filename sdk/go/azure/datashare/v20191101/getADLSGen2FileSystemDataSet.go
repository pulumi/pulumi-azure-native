


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupADLSGen2FileSystemDataSet(ctx *pulumi.Context, args *LookupADLSGen2FileSystemDataSetArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen2FileSystemDataSetResult, error) {
	var rv LookupADLSGen2FileSystemDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getADLSGen2FileSystemDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen2FileSystemDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupADLSGen2FileSystemDataSetResult struct {
	DataSetId          string `pulumi:"dataSetId"`
	FileSystem         string `pulumi:"fileSystem"`
	Id                 string `pulumi:"id"`
	Kind               string `pulumi:"kind"`
	Name               string `pulumi:"name"`
	ResourceGroup      string `pulumi:"resourceGroup"`
	StorageAccountName string `pulumi:"storageAccountName"`
	SubscriptionId     string `pulumi:"subscriptionId"`
	Type               string `pulumi:"type"`
}

func LookupADLSGen2FileSystemDataSetOutput(ctx *pulumi.Context, args LookupADLSGen2FileSystemDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupADLSGen2FileSystemDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupADLSGen2FileSystemDataSetResult, error) {
			args := v.(LookupADLSGen2FileSystemDataSetArgs)
			r, err := LookupADLSGen2FileSystemDataSet(ctx, &args, opts...)
			var s LookupADLSGen2FileSystemDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupADLSGen2FileSystemDataSetResultOutput)
}

type LookupADLSGen2FileSystemDataSetOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DataSetName       pulumi.StringInput `pulumi:"dataSetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupADLSGen2FileSystemDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen2FileSystemDataSetArgs)(nil)).Elem()
}


type LookupADLSGen2FileSystemDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupADLSGen2FileSystemDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen2FileSystemDataSetResult)(nil)).Elem()
}

func (o LookupADLSGen2FileSystemDataSetResultOutput) ToLookupADLSGen2FileSystemDataSetResultOutput() LookupADLSGen2FileSystemDataSetResultOutput {
	return o
}

func (o LookupADLSGen2FileSystemDataSetResultOutput) ToLookupADLSGen2FileSystemDataSetResultOutputWithContext(ctx context.Context) LookupADLSGen2FileSystemDataSetResultOutput {
	return o
}

func (o LookupADLSGen2FileSystemDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileSystemDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileSystemDataSetResultOutput) FileSystem() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileSystemDataSetResult) string { return v.FileSystem }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileSystemDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileSystemDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileSystemDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileSystemDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileSystemDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileSystemDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileSystemDataSetResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileSystemDataSetResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileSystemDataSetResultOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileSystemDataSetResult) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileSystemDataSetResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileSystemDataSetResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileSystemDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileSystemDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupADLSGen2FileSystemDataSetResultOutput{})
}

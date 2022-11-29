


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupADLSGen2FolderDataSet(ctx *pulumi.Context, args *LookupADLSGen2FolderDataSetArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen2FolderDataSetResult, error) {
	var rv LookupADLSGen2FolderDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20201001preview:getADLSGen2FolderDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen2FolderDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupADLSGen2FolderDataSetResult struct {
	DataSetId          string             `pulumi:"dataSetId"`
	FileSystem         string             `pulumi:"fileSystem"`
	FolderPath         string             `pulumi:"folderPath"`
	Id                 string             `pulumi:"id"`
	Kind               string             `pulumi:"kind"`
	Name               string             `pulumi:"name"`
	ResourceGroup      string             `pulumi:"resourceGroup"`
	StorageAccountName string             `pulumi:"storageAccountName"`
	SubscriptionId     string             `pulumi:"subscriptionId"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	Type               string             `pulumi:"type"`
}

func LookupADLSGen2FolderDataSetOutput(ctx *pulumi.Context, args LookupADLSGen2FolderDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupADLSGen2FolderDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupADLSGen2FolderDataSetResult, error) {
			args := v.(LookupADLSGen2FolderDataSetArgs)
			r, err := LookupADLSGen2FolderDataSet(ctx, &args, opts...)
			var s LookupADLSGen2FolderDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupADLSGen2FolderDataSetResultOutput)
}

type LookupADLSGen2FolderDataSetOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DataSetName       pulumi.StringInput `pulumi:"dataSetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupADLSGen2FolderDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen2FolderDataSetArgs)(nil)).Elem()
}


type LookupADLSGen2FolderDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupADLSGen2FolderDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen2FolderDataSetResult)(nil)).Elem()
}

func (o LookupADLSGen2FolderDataSetResultOutput) ToLookupADLSGen2FolderDataSetResultOutput() LookupADLSGen2FolderDataSetResultOutput {
	return o
}

func (o LookupADLSGen2FolderDataSetResultOutput) ToLookupADLSGen2FolderDataSetResultOutputWithContext(ctx context.Context) LookupADLSGen2FolderDataSetResultOutput {
	return o
}

func (o LookupADLSGen2FolderDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FolderDataSetResultOutput) FileSystem() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetResult) string { return v.FileSystem }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FolderDataSetResultOutput) FolderPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetResult) string { return v.FolderPath }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FolderDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FolderDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FolderDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FolderDataSetResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FolderDataSetResultOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetResult) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FolderDataSetResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FolderDataSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupADLSGen2FolderDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupADLSGen2FolderDataSetResultOutput{})
}

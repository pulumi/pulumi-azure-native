


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupADLSGen2FileDataSet(ctx *pulumi.Context, args *LookupADLSGen2FileDataSetArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen2FileDataSetResult, error) {
	var rv LookupADLSGen2FileDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20210801:getADLSGen2FileDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen2FileDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupADLSGen2FileDataSetResult struct {
	DataSetId          string             `pulumi:"dataSetId"`
	FilePath           string             `pulumi:"filePath"`
	FileSystem         string             `pulumi:"fileSystem"`
	Id                 string             `pulumi:"id"`
	Kind               string             `pulumi:"kind"`
	Name               string             `pulumi:"name"`
	ResourceGroup      string             `pulumi:"resourceGroup"`
	StorageAccountName string             `pulumi:"storageAccountName"`
	SubscriptionId     string             `pulumi:"subscriptionId"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	Type               string             `pulumi:"type"`
}

func LookupADLSGen2FileDataSetOutput(ctx *pulumi.Context, args LookupADLSGen2FileDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupADLSGen2FileDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupADLSGen2FileDataSetResult, error) {
			args := v.(LookupADLSGen2FileDataSetArgs)
			r, err := LookupADLSGen2FileDataSet(ctx, &args, opts...)
			var s LookupADLSGen2FileDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupADLSGen2FileDataSetResultOutput)
}

type LookupADLSGen2FileDataSetOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DataSetName       pulumi.StringInput `pulumi:"dataSetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupADLSGen2FileDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen2FileDataSetArgs)(nil)).Elem()
}


type LookupADLSGen2FileDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupADLSGen2FileDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen2FileDataSetResult)(nil)).Elem()
}

func (o LookupADLSGen2FileDataSetResultOutput) ToLookupADLSGen2FileDataSetResultOutput() LookupADLSGen2FileDataSetResultOutput {
	return o
}

func (o LookupADLSGen2FileDataSetResultOutput) ToLookupADLSGen2FileDataSetResultOutputWithContext(ctx context.Context) LookupADLSGen2FileDataSetResultOutput {
	return o
}

func (o LookupADLSGen2FileDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileDataSetResultOutput) FilePath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) string { return v.FilePath }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileDataSetResultOutput) FileSystem() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) string { return v.FileSystem }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileDataSetResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileDataSetResultOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileDataSetResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileDataSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupADLSGen2FileDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupADLSGen2FileDataSetResultOutput{})
}




package v20181101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupADLSGen1FileDataSet(ctx *pulumi.Context, args *LookupADLSGen1FileDataSetArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen1FileDataSetResult, error) {
	var rv LookupADLSGen1FileDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20181101preview:getADLSGen1FileDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen1FileDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupADLSGen1FileDataSetResult struct {
	AccountName    string `pulumi:"accountName"`
	DataSetId      string `pulumi:"dataSetId"`
	FileName       string `pulumi:"fileName"`
	FolderPath     string `pulumi:"folderPath"`
	Id             string `pulumi:"id"`
	Kind           string `pulumi:"kind"`
	Name           string `pulumi:"name"`
	ResourceGroup  string `pulumi:"resourceGroup"`
	SubscriptionId string `pulumi:"subscriptionId"`
	Type           string `pulumi:"type"`
}

func LookupADLSGen1FileDataSetOutput(ctx *pulumi.Context, args LookupADLSGen1FileDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupADLSGen1FileDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupADLSGen1FileDataSetResult, error) {
			args := v.(LookupADLSGen1FileDataSetArgs)
			r, err := LookupADLSGen1FileDataSet(ctx, &args, opts...)
			var s LookupADLSGen1FileDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupADLSGen1FileDataSetResultOutput)
}

type LookupADLSGen1FileDataSetOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DataSetName       pulumi.StringInput `pulumi:"dataSetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupADLSGen1FileDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen1FileDataSetArgs)(nil)).Elem()
}


type LookupADLSGen1FileDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupADLSGen1FileDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen1FileDataSetResult)(nil)).Elem()
}

func (o LookupADLSGen1FileDataSetResultOutput) ToLookupADLSGen1FileDataSetResultOutput() LookupADLSGen1FileDataSetResultOutput {
	return o
}

func (o LookupADLSGen1FileDataSetResultOutput) ToLookupADLSGen1FileDataSetResultOutputWithContext(ctx context.Context) LookupADLSGen1FileDataSetResultOutput {
	return o
}

func (o LookupADLSGen1FileDataSetResultOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.AccountName }).(pulumi.StringOutput)
}

func (o LookupADLSGen1FileDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupADLSGen1FileDataSetResultOutput) FileName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.FileName }).(pulumi.StringOutput)
}

func (o LookupADLSGen1FileDataSetResultOutput) FolderPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.FolderPath }).(pulumi.StringOutput)
}

func (o LookupADLSGen1FileDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupADLSGen1FileDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupADLSGen1FileDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupADLSGen1FileDataSetResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupADLSGen1FileDataSetResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o LookupADLSGen1FileDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupADLSGen1FileDataSetResultOutput{})
}

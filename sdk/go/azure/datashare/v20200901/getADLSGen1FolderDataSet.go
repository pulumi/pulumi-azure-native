


package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupADLSGen1FolderDataSet(ctx *pulumi.Context, args *LookupADLSGen1FolderDataSetArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen1FolderDataSetResult, error) {
	var rv LookupADLSGen1FolderDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20200901:getADLSGen1FolderDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen1FolderDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupADLSGen1FolderDataSetResult struct {
	AccountName    string             `pulumi:"accountName"`
	DataSetId      string             `pulumi:"dataSetId"`
	FolderPath     string             `pulumi:"folderPath"`
	Id             string             `pulumi:"id"`
	Kind           string             `pulumi:"kind"`
	Name           string             `pulumi:"name"`
	ResourceGroup  string             `pulumi:"resourceGroup"`
	SubscriptionId string             `pulumi:"subscriptionId"`
	SystemData     SystemDataResponse `pulumi:"systemData"`
	Type           string             `pulumi:"type"`
}

func LookupADLSGen1FolderDataSetOutput(ctx *pulumi.Context, args LookupADLSGen1FolderDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupADLSGen1FolderDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupADLSGen1FolderDataSetResult, error) {
			args := v.(LookupADLSGen1FolderDataSetArgs)
			r, err := LookupADLSGen1FolderDataSet(ctx, &args, opts...)
			var s LookupADLSGen1FolderDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupADLSGen1FolderDataSetResultOutput)
}

type LookupADLSGen1FolderDataSetOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DataSetName       pulumi.StringInput `pulumi:"dataSetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupADLSGen1FolderDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen1FolderDataSetArgs)(nil)).Elem()
}


type LookupADLSGen1FolderDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupADLSGen1FolderDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen1FolderDataSetResult)(nil)).Elem()
}

func (o LookupADLSGen1FolderDataSetResultOutput) ToLookupADLSGen1FolderDataSetResultOutput() LookupADLSGen1FolderDataSetResultOutput {
	return o
}

func (o LookupADLSGen1FolderDataSetResultOutput) ToLookupADLSGen1FolderDataSetResultOutputWithContext(ctx context.Context) LookupADLSGen1FolderDataSetResultOutput {
	return o
}

func (o LookupADLSGen1FolderDataSetResultOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FolderDataSetResult) string { return v.AccountName }).(pulumi.StringOutput)
}

func (o LookupADLSGen1FolderDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FolderDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupADLSGen1FolderDataSetResultOutput) FolderPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FolderDataSetResult) string { return v.FolderPath }).(pulumi.StringOutput)
}

func (o LookupADLSGen1FolderDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FolderDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupADLSGen1FolderDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FolderDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupADLSGen1FolderDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FolderDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupADLSGen1FolderDataSetResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FolderDataSetResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupADLSGen1FolderDataSetResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FolderDataSetResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o LookupADLSGen1FolderDataSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupADLSGen1FolderDataSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupADLSGen1FolderDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FolderDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupADLSGen1FolderDataSetResultOutput{})
}




package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobDataSet(ctx *pulumi.Context, args *LookupBlobDataSetArgs, opts ...pulumi.InvokeOption) (*LookupBlobDataSetResult, error) {
	var rv LookupBlobDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20201001preview:getBlobDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupBlobDataSetResult struct {
	ContainerName      string             `pulumi:"containerName"`
	DataSetId          string             `pulumi:"dataSetId"`
	FilePath           string             `pulumi:"filePath"`
	Id                 string             `pulumi:"id"`
	Kind               string             `pulumi:"kind"`
	Name               string             `pulumi:"name"`
	ResourceGroup      string             `pulumi:"resourceGroup"`
	StorageAccountName string             `pulumi:"storageAccountName"`
	SubscriptionId     string             `pulumi:"subscriptionId"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	Type               string             `pulumi:"type"`
}

func LookupBlobDataSetOutput(ctx *pulumi.Context, args LookupBlobDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupBlobDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlobDataSetResult, error) {
			args := v.(LookupBlobDataSetArgs)
			r, err := LookupBlobDataSet(ctx, &args, opts...)
			var s LookupBlobDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlobDataSetResultOutput)
}

type LookupBlobDataSetOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DataSetName       pulumi.StringInput `pulumi:"dataSetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupBlobDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobDataSetArgs)(nil)).Elem()
}


type LookupBlobDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupBlobDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobDataSetResult)(nil)).Elem()
}

func (o LookupBlobDataSetResultOutput) ToLookupBlobDataSetResultOutput() LookupBlobDataSetResultOutput {
	return o
}

func (o LookupBlobDataSetResultOutput) ToLookupBlobDataSetResultOutputWithContext(ctx context.Context) LookupBlobDataSetResultOutput {
	return o
}

func (o LookupBlobDataSetResultOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetResult) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o LookupBlobDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupBlobDataSetResultOutput) FilePath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetResult) string { return v.FilePath }).(pulumi.StringOutput)
}

func (o LookupBlobDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBlobDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupBlobDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBlobDataSetResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupBlobDataSetResultOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetResult) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o LookupBlobDataSetResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o LookupBlobDataSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBlobDataSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupBlobDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlobDataSetResultOutput{})
}

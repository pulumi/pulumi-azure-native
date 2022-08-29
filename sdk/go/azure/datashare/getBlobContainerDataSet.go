


package datashare

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobContainerDataSet(ctx *pulumi.Context, args *LookupBlobContainerDataSetArgs, opts ...pulumi.InvokeOption) (*LookupBlobContainerDataSetResult, error) {
	var rv LookupBlobContainerDataSetResult
	err := ctx.Invoke("azure-native:datashare:getBlobContainerDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobContainerDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupBlobContainerDataSetResult struct {
	ContainerName      string             `pulumi:"containerName"`
	DataSetId          string             `pulumi:"dataSetId"`
	Id                 string             `pulumi:"id"`
	Kind               string             `pulumi:"kind"`
	Name               string             `pulumi:"name"`
	ResourceGroup      string             `pulumi:"resourceGroup"`
	StorageAccountName string             `pulumi:"storageAccountName"`
	SubscriptionId     string             `pulumi:"subscriptionId"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	Type               string             `pulumi:"type"`
}

func LookupBlobContainerDataSetOutput(ctx *pulumi.Context, args LookupBlobContainerDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupBlobContainerDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlobContainerDataSetResult, error) {
			args := v.(LookupBlobContainerDataSetArgs)
			r, err := LookupBlobContainerDataSet(ctx, &args, opts...)
			var s LookupBlobContainerDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlobContainerDataSetResultOutput)
}

type LookupBlobContainerDataSetOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DataSetName       pulumi.StringInput `pulumi:"dataSetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupBlobContainerDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobContainerDataSetArgs)(nil)).Elem()
}


type LookupBlobContainerDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupBlobContainerDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobContainerDataSetResult)(nil)).Elem()
}

func (o LookupBlobContainerDataSetResultOutput) ToLookupBlobContainerDataSetResultOutput() LookupBlobContainerDataSetResultOutput {
	return o
}

func (o LookupBlobContainerDataSetResultOutput) ToLookupBlobContainerDataSetResultOutputWithContext(ctx context.Context) LookupBlobContainerDataSetResultOutput {
	return o
}

func (o LookupBlobContainerDataSetResultOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetResult) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o LookupBlobContainerDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupBlobContainerDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBlobContainerDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupBlobContainerDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBlobContainerDataSetResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupBlobContainerDataSetResultOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetResult) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o LookupBlobContainerDataSetResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o LookupBlobContainerDataSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupBlobContainerDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlobContainerDataSetResultOutput{})
}

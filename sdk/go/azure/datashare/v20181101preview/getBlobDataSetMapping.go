


package v20181101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobDataSetMapping(ctx *pulumi.Context, args *LookupBlobDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupBlobDataSetMappingResult, error) {
	var rv LookupBlobDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20181101preview:getBlobDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupBlobDataSetMappingResult struct {
	ContainerName        string  `pulumi:"containerName"`
	DataSetId            string  `pulumi:"dataSetId"`
	DataSetMappingStatus string  `pulumi:"dataSetMappingStatus"`
	FilePath             string  `pulumi:"filePath"`
	Id                   string  `pulumi:"id"`
	Kind                 string  `pulumi:"kind"`
	Name                 string  `pulumi:"name"`
	OutputType           *string `pulumi:"outputType"`
	ProvisioningState    string  `pulumi:"provisioningState"`
	ResourceGroup        string  `pulumi:"resourceGroup"`
	StorageAccountName   string  `pulumi:"storageAccountName"`
	SubscriptionId       string  `pulumi:"subscriptionId"`
	Type                 string  `pulumi:"type"`
}

func LookupBlobDataSetMappingOutput(ctx *pulumi.Context, args LookupBlobDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupBlobDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlobDataSetMappingResult, error) {
			args := v.(LookupBlobDataSetMappingArgs)
			r, err := LookupBlobDataSetMapping(ctx, &args, opts...)
			var s LookupBlobDataSetMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlobDataSetMappingResultOutput)
}

type LookupBlobDataSetMappingOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	DataSetMappingName    pulumi.StringInput `pulumi:"dataSetMappingName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupBlobDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobDataSetMappingArgs)(nil)).Elem()
}


type LookupBlobDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupBlobDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobDataSetMappingResult)(nil)).Elem()
}

func (o LookupBlobDataSetMappingResultOutput) ToLookupBlobDataSetMappingResultOutput() LookupBlobDataSetMappingResultOutput {
	return o
}

func (o LookupBlobDataSetMappingResultOutput) ToLookupBlobDataSetMappingResultOutputWithContext(ctx context.Context) LookupBlobDataSetMappingResultOutput {
	return o
}

func (o LookupBlobDataSetMappingResultOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetMappingResult) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o LookupBlobDataSetMappingResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetMappingResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupBlobDataSetMappingResultOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetMappingResult) string { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o LookupBlobDataSetMappingResultOutput) FilePath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetMappingResult) string { return v.FilePath }).(pulumi.StringOutput)
}

func (o LookupBlobDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBlobDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupBlobDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBlobDataSetMappingResultOutput) OutputType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlobDataSetMappingResult) *string { return v.OutputType }).(pulumi.StringPtrOutput)
}

func (o LookupBlobDataSetMappingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetMappingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupBlobDataSetMappingResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetMappingResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupBlobDataSetMappingResultOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetMappingResult) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o LookupBlobDataSetMappingResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetMappingResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o LookupBlobDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlobDataSetMappingResultOutput{})
}

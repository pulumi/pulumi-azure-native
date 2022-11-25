


package v20181101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupADLSGen2FileDataSetMapping(ctx *pulumi.Context, args *LookupADLSGen2FileDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen2FileDataSetMappingResult, error) {
	var rv LookupADLSGen2FileDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20181101preview:getADLSGen2FileDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen2FileDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupADLSGen2FileDataSetMappingResult struct {
	DataSetId            string  `pulumi:"dataSetId"`
	DataSetMappingStatus string  `pulumi:"dataSetMappingStatus"`
	FilePath             string  `pulumi:"filePath"`
	FileSystem           string  `pulumi:"fileSystem"`
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

func LookupADLSGen2FileDataSetMappingOutput(ctx *pulumi.Context, args LookupADLSGen2FileDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupADLSGen2FileDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupADLSGen2FileDataSetMappingResult, error) {
			args := v.(LookupADLSGen2FileDataSetMappingArgs)
			r, err := LookupADLSGen2FileDataSetMapping(ctx, &args, opts...)
			var s LookupADLSGen2FileDataSetMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupADLSGen2FileDataSetMappingResultOutput)
}

type LookupADLSGen2FileDataSetMappingOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	DataSetMappingName    pulumi.StringInput `pulumi:"dataSetMappingName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupADLSGen2FileDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen2FileDataSetMappingArgs)(nil)).Elem()
}


type LookupADLSGen2FileDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupADLSGen2FileDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen2FileDataSetMappingResult)(nil)).Elem()
}

func (o LookupADLSGen2FileDataSetMappingResultOutput) ToLookupADLSGen2FileDataSetMappingResultOutput() LookupADLSGen2FileDataSetMappingResultOutput {
	return o
}

func (o LookupADLSGen2FileDataSetMappingResultOutput) ToLookupADLSGen2FileDataSetMappingResultOutputWithContext(ctx context.Context) LookupADLSGen2FileDataSetMappingResultOutput {
	return o
}

func (o LookupADLSGen2FileDataSetMappingResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetMappingResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileDataSetMappingResultOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetMappingResult) string { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileDataSetMappingResultOutput) FilePath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetMappingResult) string { return v.FilePath }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileDataSetMappingResultOutput) FileSystem() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetMappingResult) string { return v.FileSystem }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileDataSetMappingResultOutput) OutputType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetMappingResult) *string { return v.OutputType }).(pulumi.StringPtrOutput)
}

func (o LookupADLSGen2FileDataSetMappingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetMappingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileDataSetMappingResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetMappingResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileDataSetMappingResultOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetMappingResult) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileDataSetMappingResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetMappingResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupADLSGen2FileDataSetMappingResultOutput{})
}

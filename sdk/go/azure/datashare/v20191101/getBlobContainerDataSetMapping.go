


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobContainerDataSetMapping(ctx *pulumi.Context, args *LookupBlobContainerDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupBlobContainerDataSetMappingResult, error) {
	var rv LookupBlobContainerDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getBlobContainerDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobContainerDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupBlobContainerDataSetMappingResult struct {
	ContainerName        string `pulumi:"containerName"`
	DataSetId            string `pulumi:"dataSetId"`
	DataSetMappingStatus string `pulumi:"dataSetMappingStatus"`
	Id                   string `pulumi:"id"`
	Kind                 string `pulumi:"kind"`
	Name                 string `pulumi:"name"`
	ProvisioningState    string `pulumi:"provisioningState"`
	ResourceGroup        string `pulumi:"resourceGroup"`
	StorageAccountName   string `pulumi:"storageAccountName"`
	SubscriptionId       string `pulumi:"subscriptionId"`
	Type                 string `pulumi:"type"`
}

func LookupBlobContainerDataSetMappingOutput(ctx *pulumi.Context, args LookupBlobContainerDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupBlobContainerDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlobContainerDataSetMappingResult, error) {
			args := v.(LookupBlobContainerDataSetMappingArgs)
			r, err := LookupBlobContainerDataSetMapping(ctx, &args, opts...)
			var s LookupBlobContainerDataSetMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlobContainerDataSetMappingResultOutput)
}

type LookupBlobContainerDataSetMappingOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	DataSetMappingName    pulumi.StringInput `pulumi:"dataSetMappingName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupBlobContainerDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobContainerDataSetMappingArgs)(nil)).Elem()
}


type LookupBlobContainerDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupBlobContainerDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobContainerDataSetMappingResult)(nil)).Elem()
}

func (o LookupBlobContainerDataSetMappingResultOutput) ToLookupBlobContainerDataSetMappingResultOutput() LookupBlobContainerDataSetMappingResultOutput {
	return o
}

func (o LookupBlobContainerDataSetMappingResultOutput) ToLookupBlobContainerDataSetMappingResultOutputWithContext(ctx context.Context) LookupBlobContainerDataSetMappingResultOutput {
	return o
}

func (o LookupBlobContainerDataSetMappingResultOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o LookupBlobContainerDataSetMappingResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupBlobContainerDataSetMappingResultOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o LookupBlobContainerDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBlobContainerDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupBlobContainerDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBlobContainerDataSetMappingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupBlobContainerDataSetMappingResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupBlobContainerDataSetMappingResultOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o LookupBlobContainerDataSetMappingResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o LookupBlobContainerDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlobContainerDataSetMappingResultOutput{})
}

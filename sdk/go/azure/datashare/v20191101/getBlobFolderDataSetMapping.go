


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobFolderDataSetMapping(ctx *pulumi.Context, args *LookupBlobFolderDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupBlobFolderDataSetMappingResult, error) {
	var rv LookupBlobFolderDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getBlobFolderDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobFolderDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupBlobFolderDataSetMappingResult struct {
	ContainerName        string `pulumi:"containerName"`
	DataSetId            string `pulumi:"dataSetId"`
	DataSetMappingStatus string `pulumi:"dataSetMappingStatus"`
	Id                   string `pulumi:"id"`
	Kind                 string `pulumi:"kind"`
	Name                 string `pulumi:"name"`
	Prefix               string `pulumi:"prefix"`
	ProvisioningState    string `pulumi:"provisioningState"`
	ResourceGroup        string `pulumi:"resourceGroup"`
	StorageAccountName   string `pulumi:"storageAccountName"`
	SubscriptionId       string `pulumi:"subscriptionId"`
	Type                 string `pulumi:"type"`
}

func LookupBlobFolderDataSetMappingOutput(ctx *pulumi.Context, args LookupBlobFolderDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupBlobFolderDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlobFolderDataSetMappingResult, error) {
			args := v.(LookupBlobFolderDataSetMappingArgs)
			r, err := LookupBlobFolderDataSetMapping(ctx, &args, opts...)
			var s LookupBlobFolderDataSetMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlobFolderDataSetMappingResultOutput)
}

type LookupBlobFolderDataSetMappingOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	DataSetMappingName    pulumi.StringInput `pulumi:"dataSetMappingName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupBlobFolderDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobFolderDataSetMappingArgs)(nil)).Elem()
}


type LookupBlobFolderDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupBlobFolderDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobFolderDataSetMappingResult)(nil)).Elem()
}

func (o LookupBlobFolderDataSetMappingResultOutput) ToLookupBlobFolderDataSetMappingResultOutput() LookupBlobFolderDataSetMappingResultOutput {
	return o
}

func (o LookupBlobFolderDataSetMappingResultOutput) ToLookupBlobFolderDataSetMappingResultOutputWithContext(ctx context.Context) LookupBlobFolderDataSetMappingResultOutput {
	return o
}

func (o LookupBlobFolderDataSetMappingResultOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o LookupBlobFolderDataSetMappingResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupBlobFolderDataSetMappingResultOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o LookupBlobFolderDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBlobFolderDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupBlobFolderDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBlobFolderDataSetMappingResultOutput) Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.Prefix }).(pulumi.StringOutput)
}

func (o LookupBlobFolderDataSetMappingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupBlobFolderDataSetMappingResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupBlobFolderDataSetMappingResultOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o LookupBlobFolderDataSetMappingResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o LookupBlobFolderDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlobFolderDataSetMappingResultOutput{})
}

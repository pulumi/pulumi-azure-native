


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupADLSGen2FileSystemDataSetMapping(ctx *pulumi.Context, args *LookupADLSGen2FileSystemDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen2FileSystemDataSetMappingResult, error) {
	var rv LookupADLSGen2FileSystemDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getADLSGen2FileSystemDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen2FileSystemDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupADLSGen2FileSystemDataSetMappingResult struct {
	DataSetId            string `pulumi:"dataSetId"`
	DataSetMappingStatus string `pulumi:"dataSetMappingStatus"`
	FileSystem           string `pulumi:"fileSystem"`
	Id                   string `pulumi:"id"`
	Kind                 string `pulumi:"kind"`
	Name                 string `pulumi:"name"`
	ProvisioningState    string `pulumi:"provisioningState"`
	ResourceGroup        string `pulumi:"resourceGroup"`
	StorageAccountName   string `pulumi:"storageAccountName"`
	SubscriptionId       string `pulumi:"subscriptionId"`
	Type                 string `pulumi:"type"`
}

func LookupADLSGen2FileSystemDataSetMappingOutput(ctx *pulumi.Context, args LookupADLSGen2FileSystemDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupADLSGen2FileSystemDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupADLSGen2FileSystemDataSetMappingResult, error) {
			args := v.(LookupADLSGen2FileSystemDataSetMappingArgs)
			r, err := LookupADLSGen2FileSystemDataSetMapping(ctx, &args, opts...)
			var s LookupADLSGen2FileSystemDataSetMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupADLSGen2FileSystemDataSetMappingResultOutput)
}

type LookupADLSGen2FileSystemDataSetMappingOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	DataSetMappingName    pulumi.StringInput `pulumi:"dataSetMappingName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupADLSGen2FileSystemDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen2FileSystemDataSetMappingArgs)(nil)).Elem()
}


type LookupADLSGen2FileSystemDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupADLSGen2FileSystemDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen2FileSystemDataSetMappingResult)(nil)).Elem()
}

func (o LookupADLSGen2FileSystemDataSetMappingResultOutput) ToLookupADLSGen2FileSystemDataSetMappingResultOutput() LookupADLSGen2FileSystemDataSetMappingResultOutput {
	return o
}

func (o LookupADLSGen2FileSystemDataSetMappingResultOutput) ToLookupADLSGen2FileSystemDataSetMappingResultOutputWithContext(ctx context.Context) LookupADLSGen2FileSystemDataSetMappingResultOutput {
	return o
}

func (o LookupADLSGen2FileSystemDataSetMappingResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileSystemDataSetMappingResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileSystemDataSetMappingResultOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileSystemDataSetMappingResult) string { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileSystemDataSetMappingResultOutput) FileSystem() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileSystemDataSetMappingResult) string { return v.FileSystem }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileSystemDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileSystemDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileSystemDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileSystemDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileSystemDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileSystemDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileSystemDataSetMappingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileSystemDataSetMappingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileSystemDataSetMappingResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileSystemDataSetMappingResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileSystemDataSetMappingResultOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileSystemDataSetMappingResult) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileSystemDataSetMappingResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileSystemDataSetMappingResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FileSystemDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileSystemDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupADLSGen2FileSystemDataSetMappingResultOutput{})
}

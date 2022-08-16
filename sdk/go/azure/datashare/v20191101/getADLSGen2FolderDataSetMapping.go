


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupADLSGen2FolderDataSetMapping(ctx *pulumi.Context, args *LookupADLSGen2FolderDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen2FolderDataSetMappingResult, error) {
	var rv LookupADLSGen2FolderDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getADLSGen2FolderDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen2FolderDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupADLSGen2FolderDataSetMappingResult struct {
	DataSetId            string `pulumi:"dataSetId"`
	DataSetMappingStatus string `pulumi:"dataSetMappingStatus"`
	FileSystem           string `pulumi:"fileSystem"`
	FolderPath           string `pulumi:"folderPath"`
	Id                   string `pulumi:"id"`
	Kind                 string `pulumi:"kind"`
	Name                 string `pulumi:"name"`
	ProvisioningState    string `pulumi:"provisioningState"`
	ResourceGroup        string `pulumi:"resourceGroup"`
	StorageAccountName   string `pulumi:"storageAccountName"`
	SubscriptionId       string `pulumi:"subscriptionId"`
	Type                 string `pulumi:"type"`
}

func LookupADLSGen2FolderDataSetMappingOutput(ctx *pulumi.Context, args LookupADLSGen2FolderDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupADLSGen2FolderDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupADLSGen2FolderDataSetMappingResult, error) {
			args := v.(LookupADLSGen2FolderDataSetMappingArgs)
			r, err := LookupADLSGen2FolderDataSetMapping(ctx, &args, opts...)
			var s LookupADLSGen2FolderDataSetMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupADLSGen2FolderDataSetMappingResultOutput)
}

type LookupADLSGen2FolderDataSetMappingOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	DataSetMappingName    pulumi.StringInput `pulumi:"dataSetMappingName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupADLSGen2FolderDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen2FolderDataSetMappingArgs)(nil)).Elem()
}


type LookupADLSGen2FolderDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupADLSGen2FolderDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen2FolderDataSetMappingResult)(nil)).Elem()
}

func (o LookupADLSGen2FolderDataSetMappingResultOutput) ToLookupADLSGen2FolderDataSetMappingResultOutput() LookupADLSGen2FolderDataSetMappingResultOutput {
	return o
}

func (o LookupADLSGen2FolderDataSetMappingResultOutput) ToLookupADLSGen2FolderDataSetMappingResultOutputWithContext(ctx context.Context) LookupADLSGen2FolderDataSetMappingResultOutput {
	return o
}

func (o LookupADLSGen2FolderDataSetMappingResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetMappingResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FolderDataSetMappingResultOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetMappingResult) string { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FolderDataSetMappingResultOutput) FileSystem() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetMappingResult) string { return v.FileSystem }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FolderDataSetMappingResultOutput) FolderPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetMappingResult) string { return v.FolderPath }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FolderDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FolderDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FolderDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FolderDataSetMappingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetMappingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FolderDataSetMappingResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetMappingResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FolderDataSetMappingResultOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetMappingResult) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FolderDataSetMappingResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetMappingResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o LookupADLSGen2FolderDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FolderDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupADLSGen2FolderDataSetMappingResultOutput{})
}

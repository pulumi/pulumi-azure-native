


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobStorageAccountDataSetMapping(ctx *pulumi.Context, args *LookupBlobStorageAccountDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupBlobStorageAccountDataSetMappingResult, error) {
	var rv LookupBlobStorageAccountDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20201001preview:getBlobStorageAccountDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobStorageAccountDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupBlobStorageAccountDataSetMappingResult struct {
	ContainerName            string             `pulumi:"containerName"`
	DataSetId                string             `pulumi:"dataSetId"`
	DataSetMappingStatus     string             `pulumi:"dataSetMappingStatus"`
	Folder                   string             `pulumi:"folder"`
	Id                       string             `pulumi:"id"`
	Kind                     string             `pulumi:"kind"`
	Location                 string             `pulumi:"location"`
	MountPath                *string            `pulumi:"mountPath"`
	Name                     string             `pulumi:"name"`
	ProvisioningState        string             `pulumi:"provisioningState"`
	StorageAccountResourceId string             `pulumi:"storageAccountResourceId"`
	SystemData               SystemDataResponse `pulumi:"systemData"`
	Type                     string             `pulumi:"type"`
}

func LookupBlobStorageAccountDataSetMappingOutput(ctx *pulumi.Context, args LookupBlobStorageAccountDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupBlobStorageAccountDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlobStorageAccountDataSetMappingResult, error) {
			args := v.(LookupBlobStorageAccountDataSetMappingArgs)
			r, err := LookupBlobStorageAccountDataSetMapping(ctx, &args, opts...)
			var s LookupBlobStorageAccountDataSetMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlobStorageAccountDataSetMappingResultOutput)
}

type LookupBlobStorageAccountDataSetMappingOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	DataSetMappingName    pulumi.StringInput `pulumi:"dataSetMappingName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupBlobStorageAccountDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobStorageAccountDataSetMappingArgs)(nil)).Elem()
}


type LookupBlobStorageAccountDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupBlobStorageAccountDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobStorageAccountDataSetMappingResult)(nil)).Elem()
}

func (o LookupBlobStorageAccountDataSetMappingResultOutput) ToLookupBlobStorageAccountDataSetMappingResultOutput() LookupBlobStorageAccountDataSetMappingResultOutput {
	return o
}

func (o LookupBlobStorageAccountDataSetMappingResultOutput) ToLookupBlobStorageAccountDataSetMappingResultOutputWithContext(ctx context.Context) LookupBlobStorageAccountDataSetMappingResultOutput {
	return o
}

func (o LookupBlobStorageAccountDataSetMappingResultOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o LookupBlobStorageAccountDataSetMappingResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupBlobStorageAccountDataSetMappingResultOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o LookupBlobStorageAccountDataSetMappingResultOutput) Folder() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.Folder }).(pulumi.StringOutput)
}

func (o LookupBlobStorageAccountDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBlobStorageAccountDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupBlobStorageAccountDataSetMappingResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupBlobStorageAccountDataSetMappingResultOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) *string { return v.MountPath }).(pulumi.StringPtrOutput)
}

func (o LookupBlobStorageAccountDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBlobStorageAccountDataSetMappingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupBlobStorageAccountDataSetMappingResultOutput) StorageAccountResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.StorageAccountResourceId }).(pulumi.StringOutput)
}

func (o LookupBlobStorageAccountDataSetMappingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupBlobStorageAccountDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlobStorageAccountDataSetMappingResultOutput{})
}

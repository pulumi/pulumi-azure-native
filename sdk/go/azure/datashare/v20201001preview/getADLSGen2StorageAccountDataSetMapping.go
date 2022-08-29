


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupADLSGen2StorageAccountDataSetMapping(ctx *pulumi.Context, args *LookupADLSGen2StorageAccountDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen2StorageAccountDataSetMappingResult, error) {
	var rv LookupADLSGen2StorageAccountDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20201001preview:getADLSGen2StorageAccountDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen2StorageAccountDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupADLSGen2StorageAccountDataSetMappingResult struct {
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

func LookupADLSGen2StorageAccountDataSetMappingOutput(ctx *pulumi.Context, args LookupADLSGen2StorageAccountDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupADLSGen2StorageAccountDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupADLSGen2StorageAccountDataSetMappingResult, error) {
			args := v.(LookupADLSGen2StorageAccountDataSetMappingArgs)
			r, err := LookupADLSGen2StorageAccountDataSetMapping(ctx, &args, opts...)
			var s LookupADLSGen2StorageAccountDataSetMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupADLSGen2StorageAccountDataSetMappingResultOutput)
}

type LookupADLSGen2StorageAccountDataSetMappingOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	DataSetMappingName    pulumi.StringInput `pulumi:"dataSetMappingName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupADLSGen2StorageAccountDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen2StorageAccountDataSetMappingArgs)(nil)).Elem()
}


type LookupADLSGen2StorageAccountDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupADLSGen2StorageAccountDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen2StorageAccountDataSetMappingResult)(nil)).Elem()
}

func (o LookupADLSGen2StorageAccountDataSetMappingResultOutput) ToLookupADLSGen2StorageAccountDataSetMappingResultOutput() LookupADLSGen2StorageAccountDataSetMappingResultOutput {
	return o
}

func (o LookupADLSGen2StorageAccountDataSetMappingResultOutput) ToLookupADLSGen2StorageAccountDataSetMappingResultOutputWithContext(ctx context.Context) LookupADLSGen2StorageAccountDataSetMappingResultOutput {
	return o
}

func (o LookupADLSGen2StorageAccountDataSetMappingResultOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetMappingResult) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o LookupADLSGen2StorageAccountDataSetMappingResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetMappingResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupADLSGen2StorageAccountDataSetMappingResultOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetMappingResult) string { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o LookupADLSGen2StorageAccountDataSetMappingResultOutput) Folder() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetMappingResult) string { return v.Folder }).(pulumi.StringOutput)
}

func (o LookupADLSGen2StorageAccountDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupADLSGen2StorageAccountDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupADLSGen2StorageAccountDataSetMappingResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetMappingResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupADLSGen2StorageAccountDataSetMappingResultOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetMappingResult) *string { return v.MountPath }).(pulumi.StringPtrOutput)
}

func (o LookupADLSGen2StorageAccountDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupADLSGen2StorageAccountDataSetMappingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetMappingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupADLSGen2StorageAccountDataSetMappingResultOutput) StorageAccountResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetMappingResult) string { return v.StorageAccountResourceId }).(pulumi.StringOutput)
}

func (o LookupADLSGen2StorageAccountDataSetMappingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetMappingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupADLSGen2StorageAccountDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2StorageAccountDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupADLSGen2StorageAccountDataSetMappingResultOutput{})
}

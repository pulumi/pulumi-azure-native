


package devtestlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomImage(ctx *pulumi.Context, args *LookupCustomImageArgs, opts ...pulumi.InvokeOption) (*LookupCustomImageResult, error) {
	var rv LookupCustomImageResult
	err := ctx.Invoke("azure-native:devtestlab:getCustomImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomImageArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupCustomImageResult struct {
	Author              *string                                `pulumi:"author"`
	CreationDate        string                                 `pulumi:"creationDate"`
	CustomImagePlan     *CustomImagePropertiesFromPlanResponse `pulumi:"customImagePlan"`
	DataDiskStorageInfo []DataDiskStorageTypeInfoResponse      `pulumi:"dataDiskStorageInfo"`
	Description         *string                                `pulumi:"description"`
	Id                  string                                 `pulumi:"id"`
	IsPlanAuthorized    *bool                                  `pulumi:"isPlanAuthorized"`
	Location            *string                                `pulumi:"location"`
	ManagedImageId      *string                                `pulumi:"managedImageId"`
	ManagedSnapshotId   *string                                `pulumi:"managedSnapshotId"`
	Name                string                                 `pulumi:"name"`
	ProvisioningState   string                                 `pulumi:"provisioningState"`
	Tags                map[string]string                      `pulumi:"tags"`
	Type                string                                 `pulumi:"type"`
	UniqueIdentifier    string                                 `pulumi:"uniqueIdentifier"`
	Vhd                 *CustomImagePropertiesCustomResponse   `pulumi:"vhd"`
	Vm                  *CustomImagePropertiesFromVmResponse   `pulumi:"vm"`
}

func LookupCustomImageOutput(ctx *pulumi.Context, args LookupCustomImageOutputArgs, opts ...pulumi.InvokeOption) LookupCustomImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomImageResult, error) {
			args := v.(LookupCustomImageArgs)
			r, err := LookupCustomImage(ctx, &args, opts...)
			var s LookupCustomImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomImageResultOutput)
}

type LookupCustomImageOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	LabName           pulumi.StringInput    `pulumi:"labName"`
	Name              pulumi.StringInput    `pulumi:"name"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupCustomImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomImageArgs)(nil)).Elem()
}


type LookupCustomImageResultOutput struct{ *pulumi.OutputState }

func (LookupCustomImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomImageResult)(nil)).Elem()
}

func (o LookupCustomImageResultOutput) ToLookupCustomImageResultOutput() LookupCustomImageResultOutput {
	return o
}

func (o LookupCustomImageResultOutput) ToLookupCustomImageResultOutputWithContext(ctx context.Context) LookupCustomImageResultOutput {
	return o
}

func (o LookupCustomImageResultOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResult) *string { return v.Author }).(pulumi.StringPtrOutput)
}

func (o LookupCustomImageResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomImageResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupCustomImageResultOutput) CustomImagePlan() CustomImagePropertiesFromPlanResponsePtrOutput {
	return o.ApplyT(func(v LookupCustomImageResult) *CustomImagePropertiesFromPlanResponse { return v.CustomImagePlan }).(CustomImagePropertiesFromPlanResponsePtrOutput)
}

func (o LookupCustomImageResultOutput) DataDiskStorageInfo() DataDiskStorageTypeInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupCustomImageResult) []DataDiskStorageTypeInfoResponse { return v.DataDiskStorageInfo }).(DataDiskStorageTypeInfoResponseArrayOutput)
}

func (o LookupCustomImageResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupCustomImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomImageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCustomImageResultOutput) IsPlanAuthorized() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResult) *bool { return v.IsPlanAuthorized }).(pulumi.BoolPtrOutput)
}

func (o LookupCustomImageResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupCustomImageResultOutput) ManagedImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResult) *string { return v.ManagedImageId }).(pulumi.StringPtrOutput)
}

func (o LookupCustomImageResultOutput) ManagedSnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResult) *string { return v.ManagedSnapshotId }).(pulumi.StringPtrOutput)
}

func (o LookupCustomImageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomImageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCustomImageResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomImageResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupCustomImageResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCustomImageResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCustomImageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomImageResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupCustomImageResultOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomImageResult) string { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func (o LookupCustomImageResultOutput) Vhd() CustomImagePropertiesCustomResponsePtrOutput {
	return o.ApplyT(func(v LookupCustomImageResult) *CustomImagePropertiesCustomResponse { return v.Vhd }).(CustomImagePropertiesCustomResponsePtrOutput)
}

func (o LookupCustomImageResultOutput) Vm() CustomImagePropertiesFromVmResponsePtrOutput {
	return o.ApplyT(func(v LookupCustomImageResult) *CustomImagePropertiesFromVmResponse { return v.Vm }).(CustomImagePropertiesFromVmResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomImageResultOutput{})
}




package devtestlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDisk(ctx *pulumi.Context, args *LookupDiskArgs, opts ...pulumi.InvokeOption) (*LookupDiskResult, error) {
	var rv LookupDiskResult
	err := ctx.Invoke("azure-native:devtestlab:getDisk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiskArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	UserName          string  `pulumi:"userName"`
}


type LookupDiskResult struct {
	CreatedDate       string            `pulumi:"createdDate"`
	DiskBlobName      *string           `pulumi:"diskBlobName"`
	DiskSizeGiB       *int              `pulumi:"diskSizeGiB"`
	DiskType          *string           `pulumi:"diskType"`
	DiskUri           *string           `pulumi:"diskUri"`
	HostCaching       *string           `pulumi:"hostCaching"`
	Id                string            `pulumi:"id"`
	LeasedByLabVmId   *string           `pulumi:"leasedByLabVmId"`
	Location          *string           `pulumi:"location"`
	ManagedDiskId     *string           `pulumi:"managedDiskId"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	StorageAccountId  *string           `pulumi:"storageAccountId"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
	UniqueIdentifier  string            `pulumi:"uniqueIdentifier"`
}

func LookupDiskOutput(ctx *pulumi.Context, args LookupDiskOutputArgs, opts ...pulumi.InvokeOption) LookupDiskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDiskResult, error) {
			args := v.(LookupDiskArgs)
			r, err := LookupDisk(ctx, &args, opts...)
			var s LookupDiskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDiskResultOutput)
}

type LookupDiskOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	LabName           pulumi.StringInput    `pulumi:"labName"`
	Name              pulumi.StringInput    `pulumi:"name"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	UserName          pulumi.StringInput    `pulumi:"userName"`
}

func (LookupDiskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskArgs)(nil)).Elem()
}


type LookupDiskResultOutput struct{ *pulumi.OutputState }

func (LookupDiskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskResult)(nil)).Elem()
}

func (o LookupDiskResultOutput) ToLookupDiskResultOutput() LookupDiskResultOutput {
	return o
}

func (o LookupDiskResultOutput) ToLookupDiskResultOutputWithContext(ctx context.Context) LookupDiskResultOutput {
	return o
}

func (o LookupDiskResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupDiskResultOutput) DiskBlobName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.DiskBlobName }).(pulumi.StringPtrOutput)
}

func (o LookupDiskResultOutput) DiskSizeGiB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *int { return v.DiskSizeGiB }).(pulumi.IntPtrOutput)
}

func (o LookupDiskResultOutput) DiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.DiskType }).(pulumi.StringPtrOutput)
}

func (o LookupDiskResultOutput) DiskUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.DiskUri }).(pulumi.StringPtrOutput)
}

func (o LookupDiskResultOutput) HostCaching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.HostCaching }).(pulumi.StringPtrOutput)
}

func (o LookupDiskResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDiskResultOutput) LeasedByLabVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.LeasedByLabVmId }).(pulumi.StringPtrOutput)
}

func (o LookupDiskResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDiskResultOutput) ManagedDiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.ManagedDiskId }).(pulumi.StringPtrOutput)
}

func (o LookupDiskResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDiskResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDiskResultOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

func (o LookupDiskResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDiskResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDiskResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDiskResultOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiskResultOutput{})
}

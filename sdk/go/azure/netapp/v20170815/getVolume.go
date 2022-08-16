


package v20170815

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("azure-native:netapp/v20170815:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVolumeArgs struct {
	AccountName       string `pulumi:"accountName"`
	PoolName          string `pulumi:"poolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VolumeName        string `pulumi:"volumeName"`
}


type LookupVolumeResult struct {
	CreationToken     string                                `pulumi:"creationToken"`
	ExportPolicy      *VolumePropertiesResponseExportPolicy `pulumi:"exportPolicy"`
	FileSystemId      string                                `pulumi:"fileSystemId"`
	Id                string                                `pulumi:"id"`
	Location          string                                `pulumi:"location"`
	Name              string                                `pulumi:"name"`
	ProvisioningState string                                `pulumi:"provisioningState"`
	ServiceLevel      string                                `pulumi:"serviceLevel"`
	SubnetId          *string                               `pulumi:"subnetId"`
	Tags              interface{}                           `pulumi:"tags"`
	Type              string                                `pulumi:"type"`
	UsageThreshold    *float64                              `pulumi:"usageThreshold"`
}


func (val *LookupVolumeResult) Defaults() *LookupVolumeResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ServiceLevel) {
		tmp.ServiceLevel = "Premium"
	}
	if isZero(tmp.UsageThreshold) {
		usageThreshold_ := 107374182400.0
		tmp.UsageThreshold = &usageThreshold_
	}
	return &tmp
}

func LookupVolumeOutput(ctx *pulumi.Context, args LookupVolumeOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeResult, error) {
			args := v.(LookupVolumeArgs)
			r, err := LookupVolume(ctx, &args, opts...)
			var s LookupVolumeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVolumeResultOutput)
}

type LookupVolumeOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	PoolName          pulumi.StringInput `pulumi:"poolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VolumeName        pulumi.StringInput `pulumi:"volumeName"`
}

func (LookupVolumeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeArgs)(nil)).Elem()
}


type LookupVolumeResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeResult)(nil)).Elem()
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutput() LookupVolumeResultOutput {
	return o
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutputWithContext(ctx context.Context) LookupVolumeResultOutput {
	return o
}

func (o LookupVolumeResultOutput) CreationToken() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.CreationToken }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) ExportPolicy() VolumePropertiesResponseExportPolicyPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *VolumePropertiesResponseExportPolicy { return v.ExportPolicy }).(VolumePropertiesResponseExportPolicyPtrOutput)
}

func (o LookupVolumeResultOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.FileSystemId }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) ServiceLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.ServiceLevel }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupVolumeResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func (o LookupVolumeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) UsageThreshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *float64 { return v.UsageThreshold }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeResultOutput{})
}

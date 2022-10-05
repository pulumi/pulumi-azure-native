


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("azure-native:netapp/v20191101:getVolume", args, &rv, opts...)
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
	BaremetalTenantId string                                  `pulumi:"baremetalTenantId"`
	CreationToken     string                                  `pulumi:"creationToken"`
	DataProtection    *VolumePropertiesResponseDataProtection `pulumi:"dataProtection"`
	ExportPolicy      *VolumePropertiesResponseExportPolicy   `pulumi:"exportPolicy"`
	FileSystemId      string                                  `pulumi:"fileSystemId"`
	Id                string                                  `pulumi:"id"`
	IsRestoring       *bool                                   `pulumi:"isRestoring"`
	Location          string                                  `pulumi:"location"`
	MountTargets      []MountTargetPropertiesResponse         `pulumi:"mountTargets"`
	Name              string                                  `pulumi:"name"`
	ProtocolTypes     []string                                `pulumi:"protocolTypes"`
	ProvisioningState string                                  `pulumi:"provisioningState"`
	ServiceLevel      *string                                 `pulumi:"serviceLevel"`
	SnapshotId        *string                                 `pulumi:"snapshotId"`
	SubnetId          string                                  `pulumi:"subnetId"`
	Tags              map[string]string                       `pulumi:"tags"`
	Type              string                                  `pulumi:"type"`
	UsageThreshold    float64                                 `pulumi:"usageThreshold"`
	UsedBytes         float64                                 `pulumi:"usedBytes"`
	VolumeType        *string                                 `pulumi:"volumeType"`
}


func (val *LookupVolumeResult) Defaults() *LookupVolumeResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ServiceLevel) {
		serviceLevel_ := "Premium"
		tmp.ServiceLevel = &serviceLevel_
	}
	if isZero(tmp.UsageThreshold) {
		tmp.UsageThreshold = 107374182400.0
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

func (o LookupVolumeResultOutput) BaremetalTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.BaremetalTenantId }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) CreationToken() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.CreationToken }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) DataProtection() VolumePropertiesResponseDataProtectionPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *VolumePropertiesResponseDataProtection { return v.DataProtection }).(VolumePropertiesResponseDataProtectionPtrOutput)
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

func (o LookupVolumeResultOutput) IsRestoring() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *bool { return v.IsRestoring }).(pulumi.BoolPtrOutput)
}

func (o LookupVolumeResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) MountTargets() MountTargetPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupVolumeResult) []MountTargetPropertiesResponse { return v.MountTargets }).(MountTargetPropertiesResponseArrayOutput)
}

func (o LookupVolumeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) ProtocolTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVolumeResult) []string { return v.ProtocolTypes }).(pulumi.StringArrayOutput)
}

func (o LookupVolumeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) ServiceLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.ServiceLevel }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeResultOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVolumeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVolumeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) UsageThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVolumeResult) float64 { return v.UsageThreshold }).(pulumi.Float64Output)
}

func (o LookupVolumeResultOutput) UsedBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVolumeResult) float64 { return v.UsedBytes }).(pulumi.Float64Output)
}

func (o LookupVolumeResultOutput) VolumeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.VolumeType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeResultOutput{})
}

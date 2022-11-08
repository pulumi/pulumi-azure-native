


package v20161001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupIscsiDisk(ctx *pulumi.Context, args *LookupIscsiDiskArgs, opts ...pulumi.InvokeOption) (*LookupIscsiDiskResult, error) {
	var rv LookupIscsiDiskResult
	err := ctx.Invoke("azure-native:storsimple/v20161001:getIscsiDisk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIscsiDiskArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	DiskName          string `pulumi:"diskName"`
	IscsiServerName   string `pulumi:"iscsiServerName"`
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupIscsiDiskResult struct {
	AccessControlRecords       []string `pulumi:"accessControlRecords"`
	DataPolicy                 string   `pulumi:"dataPolicy"`
	Description                *string  `pulumi:"description"`
	DiskStatus                 string   `pulumi:"diskStatus"`
	Id                         string   `pulumi:"id"`
	LocalUsedCapacityInBytes   float64  `pulumi:"localUsedCapacityInBytes"`
	MonitoringStatus           string   `pulumi:"monitoringStatus"`
	Name                       string   `pulumi:"name"`
	ProvisionedCapacityInBytes float64  `pulumi:"provisionedCapacityInBytes"`
	Type                       string   `pulumi:"type"`
	UsedCapacityInBytes        float64  `pulumi:"usedCapacityInBytes"`
}

func LookupIscsiDiskOutput(ctx *pulumi.Context, args LookupIscsiDiskOutputArgs, opts ...pulumi.InvokeOption) LookupIscsiDiskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIscsiDiskResult, error) {
			args := v.(LookupIscsiDiskArgs)
			r, err := LookupIscsiDisk(ctx, &args, opts...)
			var s LookupIscsiDiskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIscsiDiskResultOutput)
}

type LookupIscsiDiskOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	DiskName          pulumi.StringInput `pulumi:"diskName"`
	IscsiServerName   pulumi.StringInput `pulumi:"iscsiServerName"`
	ManagerName       pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIscsiDiskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIscsiDiskArgs)(nil)).Elem()
}


type LookupIscsiDiskResultOutput struct{ *pulumi.OutputState }

func (LookupIscsiDiskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIscsiDiskResult)(nil)).Elem()
}

func (o LookupIscsiDiskResultOutput) ToLookupIscsiDiskResultOutput() LookupIscsiDiskResultOutput {
	return o
}

func (o LookupIscsiDiskResultOutput) ToLookupIscsiDiskResultOutputWithContext(ctx context.Context) LookupIscsiDiskResultOutput {
	return o
}

func (o LookupIscsiDiskResultOutput) AccessControlRecords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIscsiDiskResult) []string { return v.AccessControlRecords }).(pulumi.StringArrayOutput)
}

func (o LookupIscsiDiskResultOutput) DataPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiDiskResult) string { return v.DataPolicy }).(pulumi.StringOutput)
}

func (o LookupIscsiDiskResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIscsiDiskResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupIscsiDiskResultOutput) DiskStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiDiskResult) string { return v.DiskStatus }).(pulumi.StringOutput)
}

func (o LookupIscsiDiskResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiDiskResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIscsiDiskResultOutput) LocalUsedCapacityInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupIscsiDiskResult) float64 { return v.LocalUsedCapacityInBytes }).(pulumi.Float64Output)
}

func (o LookupIscsiDiskResultOutput) MonitoringStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiDiskResult) string { return v.MonitoringStatus }).(pulumi.StringOutput)
}

func (o LookupIscsiDiskResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiDiskResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIscsiDiskResultOutput) ProvisionedCapacityInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupIscsiDiskResult) float64 { return v.ProvisionedCapacityInBytes }).(pulumi.Float64Output)
}

func (o LookupIscsiDiskResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiDiskResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupIscsiDiskResultOutput) UsedCapacityInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupIscsiDiskResult) float64 { return v.UsedCapacityInBytes }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(LookupIscsiDiskResultOutput{})
}

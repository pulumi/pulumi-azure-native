


package v20161001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupFileShare(ctx *pulumi.Context, args *LookupFileShareArgs, opts ...pulumi.InvokeOption) (*LookupFileShareResult, error) {
	var rv LookupFileShareResult
	err := ctx.Invoke("azure-native:storsimple/v20161001:getFileShare", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileShareArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	FileServerName    string `pulumi:"fileServerName"`
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupFileShareResult struct {
	AdminUser                  string  `pulumi:"adminUser"`
	DataPolicy                 string  `pulumi:"dataPolicy"`
	Description                *string `pulumi:"description"`
	Id                         string  `pulumi:"id"`
	LocalUsedCapacityInBytes   float64 `pulumi:"localUsedCapacityInBytes"`
	MonitoringStatus           string  `pulumi:"monitoringStatus"`
	Name                       string  `pulumi:"name"`
	ProvisionedCapacityInBytes float64 `pulumi:"provisionedCapacityInBytes"`
	ShareStatus                string  `pulumi:"shareStatus"`
	Type                       string  `pulumi:"type"`
	UsedCapacityInBytes        float64 `pulumi:"usedCapacityInBytes"`
}

func LookupFileShareOutput(ctx *pulumi.Context, args LookupFileShareOutputArgs, opts ...pulumi.InvokeOption) LookupFileShareResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFileShareResult, error) {
			args := v.(LookupFileShareArgs)
			r, err := LookupFileShare(ctx, &args, opts...)
			var s LookupFileShareResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFileShareResultOutput)
}

type LookupFileShareOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	FileServerName    pulumi.StringInput `pulumi:"fileServerName"`
	ManagerName       pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupFileShareOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileShareArgs)(nil)).Elem()
}


type LookupFileShareResultOutput struct{ *pulumi.OutputState }

func (LookupFileShareResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileShareResult)(nil)).Elem()
}

func (o LookupFileShareResultOutput) ToLookupFileShareResultOutput() LookupFileShareResultOutput {
	return o
}

func (o LookupFileShareResultOutput) ToLookupFileShareResultOutputWithContext(ctx context.Context) LookupFileShareResultOutput {
	return o
}

func (o LookupFileShareResultOutput) AdminUser() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.AdminUser }).(pulumi.StringOutput)
}

func (o LookupFileShareResultOutput) DataPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.DataPolicy }).(pulumi.StringOutput)
}

func (o LookupFileShareResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFileShareResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFileShareResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFileShareResultOutput) LocalUsedCapacityInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupFileShareResult) float64 { return v.LocalUsedCapacityInBytes }).(pulumi.Float64Output)
}

func (o LookupFileShareResultOutput) MonitoringStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.MonitoringStatus }).(pulumi.StringOutput)
}

func (o LookupFileShareResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFileShareResultOutput) ProvisionedCapacityInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupFileShareResult) float64 { return v.ProvisionedCapacityInBytes }).(pulumi.Float64Output)
}

func (o LookupFileShareResultOutput) ShareStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.ShareStatus }).(pulumi.StringOutput)
}

func (o LookupFileShareResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupFileShareResultOutput) UsedCapacityInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupFileShareResult) float64 { return v.UsedCapacityInBytes }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(LookupFileShareResultOutput{})
}

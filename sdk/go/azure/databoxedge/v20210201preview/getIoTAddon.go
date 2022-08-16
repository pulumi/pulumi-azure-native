


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIoTAddon(ctx *pulumi.Context, args *LookupIoTAddonArgs, opts ...pulumi.InvokeOption) (*LookupIoTAddonResult, error) {
	var rv LookupIoTAddonResult
	err := ctx.Invoke("azure-native:databoxedge/v20210201preview:getIoTAddon", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIoTAddonArgs struct {
	AddonName         string `pulumi:"addonName"`
	DeviceName        string `pulumi:"deviceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RoleName          string `pulumi:"roleName"`
}


type LookupIoTAddonResult struct {
	HostPlatform         string                `pulumi:"hostPlatform"`
	HostPlatformType     string                `pulumi:"hostPlatformType"`
	Id                   string                `pulumi:"id"`
	IoTDeviceDetails     IoTDeviceInfoResponse `pulumi:"ioTDeviceDetails"`
	IoTEdgeDeviceDetails IoTDeviceInfoResponse `pulumi:"ioTEdgeDeviceDetails"`
	Kind                 string                `pulumi:"kind"`
	Name                 string                `pulumi:"name"`
	ProvisioningState    string                `pulumi:"provisioningState"`
	SystemData           SystemDataResponse    `pulumi:"systemData"`
	Type                 string                `pulumi:"type"`
	Version              string                `pulumi:"version"`
}

func LookupIoTAddonOutput(ctx *pulumi.Context, args LookupIoTAddonOutputArgs, opts ...pulumi.InvokeOption) LookupIoTAddonResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIoTAddonResult, error) {
			args := v.(LookupIoTAddonArgs)
			r, err := LookupIoTAddon(ctx, &args, opts...)
			var s LookupIoTAddonResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIoTAddonResultOutput)
}

type LookupIoTAddonOutputArgs struct {
	AddonName         pulumi.StringInput `pulumi:"addonName"`
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RoleName          pulumi.StringInput `pulumi:"roleName"`
}

func (LookupIoTAddonOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIoTAddonArgs)(nil)).Elem()
}


type LookupIoTAddonResultOutput struct{ *pulumi.OutputState }

func (LookupIoTAddonResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIoTAddonResult)(nil)).Elem()
}

func (o LookupIoTAddonResultOutput) ToLookupIoTAddonResultOutput() LookupIoTAddonResultOutput {
	return o
}

func (o LookupIoTAddonResultOutput) ToLookupIoTAddonResultOutputWithContext(ctx context.Context) LookupIoTAddonResultOutput {
	return o
}

func (o LookupIoTAddonResultOutput) HostPlatform() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) string { return v.HostPlatform }).(pulumi.StringOutput)
}

func (o LookupIoTAddonResultOutput) HostPlatformType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) string { return v.HostPlatformType }).(pulumi.StringOutput)
}

func (o LookupIoTAddonResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIoTAddonResultOutput) IoTDeviceDetails() IoTDeviceInfoResponseOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) IoTDeviceInfoResponse { return v.IoTDeviceDetails }).(IoTDeviceInfoResponseOutput)
}

func (o LookupIoTAddonResultOutput) IoTEdgeDeviceDetails() IoTDeviceInfoResponseOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) IoTDeviceInfoResponse { return v.IoTEdgeDeviceDetails }).(IoTDeviceInfoResponseOutput)
}

func (o LookupIoTAddonResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupIoTAddonResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIoTAddonResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupIoTAddonResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupIoTAddonResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupIoTAddonResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIoTAddonResultOutput{})
}

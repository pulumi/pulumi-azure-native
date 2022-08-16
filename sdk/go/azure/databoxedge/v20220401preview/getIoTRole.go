


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIoTRole(ctx *pulumi.Context, args *LookupIoTRoleArgs, opts ...pulumi.InvokeOption) (*LookupIoTRoleResult, error) {
	var rv LookupIoTRoleResult
	err := ctx.Invoke("azure-native:databoxedge/v20220401preview:getIoTRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIoTRoleArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupIoTRoleResult struct {
	ComputeResource      *ComputeResourceResponse  `pulumi:"computeResource"`
	HostPlatform         string                    `pulumi:"hostPlatform"`
	HostPlatformType     string                    `pulumi:"hostPlatformType"`
	Id                   string                    `pulumi:"id"`
	IoTDeviceDetails     IoTDeviceInfoResponse     `pulumi:"ioTDeviceDetails"`
	IoTEdgeAgentInfo     *IoTEdgeAgentInfoResponse `pulumi:"ioTEdgeAgentInfo"`
	IoTEdgeDeviceDetails IoTDeviceInfoResponse     `pulumi:"ioTEdgeDeviceDetails"`
	Kind                 string                    `pulumi:"kind"`
	Name                 string                    `pulumi:"name"`
	RoleStatus           string                    `pulumi:"roleStatus"`
	ShareMappings        []MountPointMapResponse   `pulumi:"shareMappings"`
	SystemData           SystemDataResponse        `pulumi:"systemData"`
	Type                 string                    `pulumi:"type"`
}

func LookupIoTRoleOutput(ctx *pulumi.Context, args LookupIoTRoleOutputArgs, opts ...pulumi.InvokeOption) LookupIoTRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIoTRoleResult, error) {
			args := v.(LookupIoTRoleArgs)
			r, err := LookupIoTRole(ctx, &args, opts...)
			var s LookupIoTRoleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIoTRoleResultOutput)
}

type LookupIoTRoleOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIoTRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIoTRoleArgs)(nil)).Elem()
}


type LookupIoTRoleResultOutput struct{ *pulumi.OutputState }

func (LookupIoTRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIoTRoleResult)(nil)).Elem()
}

func (o LookupIoTRoleResultOutput) ToLookupIoTRoleResultOutput() LookupIoTRoleResultOutput {
	return o
}

func (o LookupIoTRoleResultOutput) ToLookupIoTRoleResultOutputWithContext(ctx context.Context) LookupIoTRoleResultOutput {
	return o
}

func (o LookupIoTRoleResultOutput) ComputeResource() ComputeResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupIoTRoleResult) *ComputeResourceResponse { return v.ComputeResource }).(ComputeResourceResponsePtrOutput)
}

func (o LookupIoTRoleResultOutput) HostPlatform() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTRoleResult) string { return v.HostPlatform }).(pulumi.StringOutput)
}

func (o LookupIoTRoleResultOutput) HostPlatformType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTRoleResult) string { return v.HostPlatformType }).(pulumi.StringOutput)
}

func (o LookupIoTRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIoTRoleResultOutput) IoTDeviceDetails() IoTDeviceInfoResponseOutput {
	return o.ApplyT(func(v LookupIoTRoleResult) IoTDeviceInfoResponse { return v.IoTDeviceDetails }).(IoTDeviceInfoResponseOutput)
}

func (o LookupIoTRoleResultOutput) IoTEdgeAgentInfo() IoTEdgeAgentInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupIoTRoleResult) *IoTEdgeAgentInfoResponse { return v.IoTEdgeAgentInfo }).(IoTEdgeAgentInfoResponsePtrOutput)
}

func (o LookupIoTRoleResultOutput) IoTEdgeDeviceDetails() IoTDeviceInfoResponseOutput {
	return o.ApplyT(func(v LookupIoTRoleResult) IoTDeviceInfoResponse { return v.IoTEdgeDeviceDetails }).(IoTDeviceInfoResponseOutput)
}

func (o LookupIoTRoleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTRoleResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupIoTRoleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTRoleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIoTRoleResultOutput) RoleStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTRoleResult) string { return v.RoleStatus }).(pulumi.StringOutput)
}

func (o LookupIoTRoleResultOutput) ShareMappings() MountPointMapResponseArrayOutput {
	return o.ApplyT(func(v LookupIoTRoleResult) []MountPointMapResponse { return v.ShareMappings }).(MountPointMapResponseArrayOutput)
}

func (o LookupIoTRoleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIoTRoleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupIoTRoleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTRoleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIoTRoleResultOutput{})
}

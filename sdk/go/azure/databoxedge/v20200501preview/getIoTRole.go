


package v20200501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIoTRole(ctx *pulumi.Context, args *LookupIoTRoleArgs, opts ...pulumi.InvokeOption) (*LookupIoTRoleResult, error) {
	var rv LookupIoTRoleResult
	err := ctx.Invoke("azure-native:databoxedge/v20200501preview:getIoTRole", args, &rv, opts...)
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
	Type                 string                    `pulumi:"type"`
}

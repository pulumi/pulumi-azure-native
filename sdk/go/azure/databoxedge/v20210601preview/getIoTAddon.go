


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIoTAddon(ctx *pulumi.Context, args *LookupIoTAddonArgs, opts ...pulumi.InvokeOption) (*LookupIoTAddonResult, error) {
	var rv LookupIoTAddonResult
	err := ctx.Invoke("azure-native:databoxedge/v20210601preview:getIoTAddon", args, &rv, opts...)
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

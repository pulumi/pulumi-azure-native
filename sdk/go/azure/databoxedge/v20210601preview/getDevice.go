


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDevice(ctx *pulumi.Context, args *LookupDeviceArgs, opts ...pulumi.InvokeOption) (*LookupDeviceResult, error) {
	var rv LookupDeviceResult
	err := ctx.Invoke("azure-native:databoxedge/v20210601preview:getDevice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeviceArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDeviceResult struct {
	ConfiguredRoleTypes     []string                    `pulumi:"configuredRoleTypes"`
	Culture                 string                      `pulumi:"culture"`
	DataBoxEdgeDeviceStatus *string                     `pulumi:"dataBoxEdgeDeviceStatus"`
	DataResidency           *DataResidencyResponse      `pulumi:"dataResidency"`
	Description             string                      `pulumi:"description"`
	DeviceHcsVersion        string                      `pulumi:"deviceHcsVersion"`
	DeviceLocalCapacity     float64                     `pulumi:"deviceLocalCapacity"`
	DeviceModel             string                      `pulumi:"deviceModel"`
	DeviceSoftwareVersion   string                      `pulumi:"deviceSoftwareVersion"`
	DeviceType              string                      `pulumi:"deviceType"`
	EdgeProfile             EdgeProfileResponse         `pulumi:"edgeProfile"`
	Etag                    *string                     `pulumi:"etag"`
	FriendlyName            string                      `pulumi:"friendlyName"`
	Id                      string                      `pulumi:"id"`
	Identity                *ResourceIdentityResponse   `pulumi:"identity"`
	Kind                    *string                     `pulumi:"kind"`
	Location                string                      `pulumi:"location"`
	ModelDescription        string                      `pulumi:"modelDescription"`
	Name                    string                      `pulumi:"name"`
	NodeCount               int                         `pulumi:"nodeCount"`
	ResourceMoveDetails     ResourceMoveDetailsResponse `pulumi:"resourceMoveDetails"`
	SerialNumber            string                      `pulumi:"serialNumber"`
	Sku                     *SkuResponse                `pulumi:"sku"`
	SystemData              SystemDataResponse          `pulumi:"systemData"`
	Tags                    map[string]string           `pulumi:"tags"`
	TimeZone                string                      `pulumi:"timeZone"`
	Type                    string                      `pulumi:"type"`
}

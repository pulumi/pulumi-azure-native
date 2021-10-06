


package hybridnetwork

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDevice(ctx *pulumi.Context, args *LookupDeviceArgs, opts ...pulumi.InvokeOption) (*LookupDeviceResult, error) {
	var rv LookupDeviceResult
	err := ctx.Invoke("azure-native:hybridnetwork:getDevice", args, &rv, opts...)
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
	AzureStackEdge    *SubResourceResponse  `pulumi:"azureStackEdge"`
	DeviceType        string                `pulumi:"deviceType"`
	Id                string                `pulumi:"id"`
	Location          string                `pulumi:"location"`
	Name              string                `pulumi:"name"`
	NetworkFunctions  []SubResourceResponse `pulumi:"networkFunctions"`
	ProvisioningState string                `pulumi:"provisioningState"`
	Status            string                `pulumi:"status"`
	Tags              map[string]string     `pulumi:"tags"`
	Type              string                `pulumi:"type"`
}




package v20180901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContainerGroup(ctx *pulumi.Context, args *LookupContainerGroupArgs, opts ...pulumi.InvokeOption) (*LookupContainerGroupResult, error) {
	var rv LookupContainerGroupResult
	err := ctx.Invoke("azure-native:containerinstance/v20180901:getContainerGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContainerGroupArgs struct {
	ContainerGroupName string `pulumi:"containerGroupName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupContainerGroupResult struct {
	Containers               []ContainerResponse                   `pulumi:"containers"`
	Diagnostics              *ContainerGroupDiagnosticsResponse    `pulumi:"diagnostics"`
	Id                       string                                `pulumi:"id"`
	ImageRegistryCredentials []ImageRegistryCredentialResponse     `pulumi:"imageRegistryCredentials"`
	InstanceView             ContainerGroupResponseInstanceView    `pulumi:"instanceView"`
	IpAddress                *IpAddressResponse                    `pulumi:"ipAddress"`
	Location                 *string                               `pulumi:"location"`
	Name                     string                                `pulumi:"name"`
	NetworkProfile           *ContainerGroupNetworkProfileResponse `pulumi:"networkProfile"`
	OsType                   string                                `pulumi:"osType"`
	ProvisioningState        string                                `pulumi:"provisioningState"`
	RestartPolicy            *string                               `pulumi:"restartPolicy"`
	Tags                     map[string]string                     `pulumi:"tags"`
	Type                     string                                `pulumi:"type"`
	Volumes                  []VolumeResponse                      `pulumi:"volumes"`
}

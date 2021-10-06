


package v20200301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegisteredServer(ctx *pulumi.Context, args *LookupRegisteredServerArgs, opts ...pulumi.InvokeOption) (*LookupRegisteredServerResult, error) {
	var rv LookupRegisteredServerResult
	err := ctx.Invoke("azure-native:storagesync/v20200301:getRegisteredServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegisteredServerArgs struct {
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	ServerId               string `pulumi:"serverId"`
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
}


type LookupRegisteredServerResult struct {
	AgentVersion              *string `pulumi:"agentVersion"`
	ClusterId                 *string `pulumi:"clusterId"`
	ClusterName               *string `pulumi:"clusterName"`
	DiscoveryEndpointUri      *string `pulumi:"discoveryEndpointUri"`
	FriendlyName              *string `pulumi:"friendlyName"`
	Id                        string  `pulumi:"id"`
	LastHeartBeat             *string `pulumi:"lastHeartBeat"`
	LastOperationName         *string `pulumi:"lastOperationName"`
	LastWorkflowId            *string `pulumi:"lastWorkflowId"`
	ManagementEndpointUri     *string `pulumi:"managementEndpointUri"`
	MonitoringConfiguration   *string `pulumi:"monitoringConfiguration"`
	MonitoringEndpointUri     *string `pulumi:"monitoringEndpointUri"`
	Name                      string  `pulumi:"name"`
	ProvisioningState         *string `pulumi:"provisioningState"`
	ResourceLocation          *string `pulumi:"resourceLocation"`
	ServerCertificate         *string `pulumi:"serverCertificate"`
	ServerId                  *string `pulumi:"serverId"`
	ServerManagementErrorCode *int    `pulumi:"serverManagementErrorCode"`
	ServerOSVersion           *string `pulumi:"serverOSVersion"`
	ServerRole                *string `pulumi:"serverRole"`
	ServiceLocation           *string `pulumi:"serviceLocation"`
	StorageSyncServiceUid     *string `pulumi:"storageSyncServiceUid"`
	Type                      string  `pulumi:"type"`
}

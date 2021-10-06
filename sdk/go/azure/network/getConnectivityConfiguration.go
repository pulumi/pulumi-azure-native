


package network

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectivityConfiguration(ctx *pulumi.Context, args *LookupConnectivityConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupConnectivityConfigurationResult, error) {
	var rv LookupConnectivityConfigurationResult
	err := ctx.Invoke("azure-native:network:getConnectivityConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectivityConfigurationArgs struct {
	ConfigurationName  string `pulumi:"configurationName"`
	NetworkManagerName string `pulumi:"networkManagerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupConnectivityConfigurationResult struct {
	AppliesToGroups       []ConnectivityGroupItemResponse `pulumi:"appliesToGroups"`
	ConnectivityTopology  string                          `pulumi:"connectivityTopology"`
	DeleteExistingPeering *string                         `pulumi:"deleteExistingPeering"`
	Description           *string                         `pulumi:"description"`
	DisplayName           *string                         `pulumi:"displayName"`
	Etag                  string                          `pulumi:"etag"`
	Hubs                  []HubResponse                   `pulumi:"hubs"`
	Id                    string                          `pulumi:"id"`
	IsGlobal              *string                         `pulumi:"isGlobal"`
	Name                  string                          `pulumi:"name"`
	ProvisioningState     string                          `pulumi:"provisioningState"`
	SystemData            SystemDataResponse              `pulumi:"systemData"`
	Type                  string                          `pulumi:"type"`
}




package v20190930preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOpenShiftManagedCluster(ctx *pulumi.Context, args *LookupOpenShiftManagedClusterArgs, opts ...pulumi.InvokeOption) (*LookupOpenShiftManagedClusterResult, error) {
	var rv LookupOpenShiftManagedClusterResult
	err := ctx.Invoke("azure-native:containerservice/v20190930preview:getOpenShiftManagedCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOpenShiftManagedClusterArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupOpenShiftManagedClusterResult struct {
	AgentPoolProfiles []OpenShiftManagedClusterAgentPoolProfileResponse `pulumi:"agentPoolProfiles"`
	AuthProfile       *OpenShiftManagedClusterAuthProfileResponse       `pulumi:"authProfile"`
	ClusterVersion    string                                            `pulumi:"clusterVersion"`
	Fqdn              string                                            `pulumi:"fqdn"`
	Id                string                                            `pulumi:"id"`
	Location          string                                            `pulumi:"location"`
	MasterPoolProfile *OpenShiftManagedClusterMasterPoolProfileResponse `pulumi:"masterPoolProfile"`
	MonitorProfile    *OpenShiftManagedClusterMonitorProfileResponse    `pulumi:"monitorProfile"`
	Name              string                                            `pulumi:"name"`
	NetworkProfile    *NetworkProfileResponse                           `pulumi:"networkProfile"`
	OpenShiftVersion  string                                            `pulumi:"openShiftVersion"`
	Plan              *PurchasePlanResponse                             `pulumi:"plan"`
	ProvisioningState string                                            `pulumi:"provisioningState"`
	PublicHostname    string                                            `pulumi:"publicHostname"`
	RouterProfiles    []OpenShiftRouterProfileResponse                  `pulumi:"routerProfiles"`
	Tags              map[string]string                                 `pulumi:"tags"`
	Type              string                                            `pulumi:"type"`
}




package v20200430

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOpenShiftCluster(ctx *pulumi.Context, args *LookupOpenShiftClusterArgs, opts ...pulumi.InvokeOption) (*LookupOpenShiftClusterResult, error) {
	var rv LookupOpenShiftClusterResult
	err := ctx.Invoke("azure-native:redhatopenshift/v20200430:getOpenShiftCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOpenShiftClusterArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupOpenShiftClusterResult struct {
	ApiserverProfile        *APIServerProfileResponse        `pulumi:"apiserverProfile"`
	ClusterProfile          *ClusterProfileResponse          `pulumi:"clusterProfile"`
	ConsoleProfile          *ConsoleProfileResponse          `pulumi:"consoleProfile"`
	Id                      string                           `pulumi:"id"`
	IngressProfiles         []IngressProfileResponse         `pulumi:"ingressProfiles"`
	Location                string                           `pulumi:"location"`
	MasterProfile           *MasterProfileResponse           `pulumi:"masterProfile"`
	Name                    string                           `pulumi:"name"`
	NetworkProfile          *NetworkProfileResponse          `pulumi:"networkProfile"`
	ProvisioningState       *string                          `pulumi:"provisioningState"`
	ServicePrincipalProfile *ServicePrincipalProfileResponse `pulumi:"servicePrincipalProfile"`
	Tags                    map[string]string                `pulumi:"tags"`
	Type                    string                           `pulumi:"type"`
	WorkerProfiles          []WorkerProfileResponse          `pulumi:"workerProfiles"`
}

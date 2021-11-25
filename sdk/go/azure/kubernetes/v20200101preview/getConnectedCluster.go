


package v20200101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectedCluster(ctx *pulumi.Context, args *LookupConnectedClusterArgs, opts ...pulumi.InvokeOption) (*LookupConnectedClusterResult, error) {
	var rv LookupConnectedClusterResult
	err := ctx.Invoke("azure-native:kubernetes/v20200101preview:getConnectedCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectedClusterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupConnectedClusterResult struct {
	AadProfile                               ConnectedClusterAADProfileResponse `pulumi:"aadProfile"`
	AgentPublicKeyCertificate                string                             `pulumi:"agentPublicKeyCertificate"`
	AgentVersion                             string                             `pulumi:"agentVersion"`
	ConnectivityStatus                       *string                            `pulumi:"connectivityStatus"`
	Distribution                             *string                            `pulumi:"distribution"`
	Id                                       string                             `pulumi:"id"`
	Identity                                 ConnectedClusterIdentityResponse   `pulumi:"identity"`
	Infrastructure                           *string                            `pulumi:"infrastructure"`
	KubernetesVersion                        string                             `pulumi:"kubernetesVersion"`
	LastConnectivityTime                     string                             `pulumi:"lastConnectivityTime"`
	Location                                 string                             `pulumi:"location"`
	ManagedIdentityCertificateExpirationTime string                             `pulumi:"managedIdentityCertificateExpirationTime"`
	Name                                     string                             `pulumi:"name"`
	Offering                                 string                             `pulumi:"offering"`
	ProvisioningState                        *string                            `pulumi:"provisioningState"`
	Tags                                     map[string]string                  `pulumi:"tags"`
	TotalCoreCount                           int                                `pulumi:"totalCoreCount"`
	TotalNodeCount                           int                                `pulumi:"totalNodeCount"`
	Type                                     string                             `pulumi:"type"`
}

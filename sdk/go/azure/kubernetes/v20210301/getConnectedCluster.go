


package v20210301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectedCluster(ctx *pulumi.Context, args *LookupConnectedClusterArgs, opts ...pulumi.InvokeOption) (*LookupConnectedClusterResult, error) {
	var rv LookupConnectedClusterResult
	err := ctx.Invoke("azure-native:kubernetes/v20210301:getConnectedCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupConnectedClusterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupConnectedClusterResult struct {
	AgentPublicKeyCertificate                string                           `pulumi:"agentPublicKeyCertificate"`
	AgentVersion                             string                           `pulumi:"agentVersion"`
	ConnectivityStatus                       string                           `pulumi:"connectivityStatus"`
	Distribution                             *string                          `pulumi:"distribution"`
	Id                                       string                           `pulumi:"id"`
	Identity                                 ConnectedClusterIdentityResponse `pulumi:"identity"`
	Infrastructure                           *string                          `pulumi:"infrastructure"`
	KubernetesVersion                        string                           `pulumi:"kubernetesVersion"`
	LastConnectivityTime                     string                           `pulumi:"lastConnectivityTime"`
	Location                                 string                           `pulumi:"location"`
	ManagedIdentityCertificateExpirationTime string                           `pulumi:"managedIdentityCertificateExpirationTime"`
	Name                                     string                           `pulumi:"name"`
	Offering                                 string                           `pulumi:"offering"`
	ProvisioningState                        *string                          `pulumi:"provisioningState"`
	SystemData                               SystemDataResponse               `pulumi:"systemData"`
	Tags                                     map[string]string                `pulumi:"tags"`
	TotalCoreCount                           int                              `pulumi:"totalCoreCount"`
	TotalNodeCount                           int                              `pulumi:"totalNodeCount"`
	Type                                     string                           `pulumi:"type"`
}


func (val *LookupConnectedClusterResult) Defaults() *LookupConnectedClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Identity = *tmp.Identity.Defaults()

	return &tmp
}

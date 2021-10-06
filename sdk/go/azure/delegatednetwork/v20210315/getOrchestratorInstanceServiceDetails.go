


package v20210315

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOrchestratorInstanceServiceDetails(ctx *pulumi.Context, args *LookupOrchestratorInstanceServiceDetailsArgs, opts ...pulumi.InvokeOption) (*LookupOrchestratorInstanceServiceDetailsResult, error) {
	var rv LookupOrchestratorInstanceServiceDetailsResult
	err := ctx.Invoke("azure-native:delegatednetwork/v20210315:getOrchestratorInstanceServiceDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrchestratorInstanceServiceDetailsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupOrchestratorInstanceServiceDetailsResult struct {
	ApiServerEndpoint     *string                       `pulumi:"apiServerEndpoint"`
	ClusterRootCA         *string                       `pulumi:"clusterRootCA"`
	ControllerDetails     ControllerDetailsResponse     `pulumi:"controllerDetails"`
	Id                    string                        `pulumi:"id"`
	Identity              *OrchestratorIdentityResponse `pulumi:"identity"`
	Kind                  string                        `pulumi:"kind"`
	Location              *string                       `pulumi:"location"`
	Name                  string                        `pulumi:"name"`
	OrchestratorAppId     *string                       `pulumi:"orchestratorAppId"`
	OrchestratorTenantId  *string                       `pulumi:"orchestratorTenantId"`
	PrivateLinkResourceId *string                       `pulumi:"privateLinkResourceId"`
	ProvisioningState     string                        `pulumi:"provisioningState"`
	ResourceGuid          string                        `pulumi:"resourceGuid"`
	Tags                  map[string]string             `pulumi:"tags"`
	Type                  string                        `pulumi:"type"`
}

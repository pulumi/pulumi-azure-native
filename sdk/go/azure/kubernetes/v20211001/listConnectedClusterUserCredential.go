


package v20211001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListConnectedClusterUserCredential(ctx *pulumi.Context, args *ListConnectedClusterUserCredentialArgs, opts ...pulumi.InvokeOption) (*ListConnectedClusterUserCredentialResult, error) {
	var rv ListConnectedClusterUserCredentialResult
	err := ctx.Invoke("azure-native:kubernetes/v20211001:listConnectedClusterUserCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConnectedClusterUserCredentialArgs struct {
	AuthenticationMethod string `pulumi:"authenticationMethod"`
	ClientProxy          bool   `pulumi:"clientProxy"`
	ClusterName          string `pulumi:"clusterName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type ListConnectedClusterUserCredentialResult struct {
	HybridConnectionConfig HybridConnectionConfigResponse `pulumi:"hybridConnectionConfig"`
	Kubeconfigs            []CredentialResultResponse     `pulumi:"kubeconfigs"`
}




package kubernetes

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListConnectedClusterUserCredentials(ctx *pulumi.Context, args *ListConnectedClusterUserCredentialsArgs, opts ...pulumi.InvokeOption) (*ListConnectedClusterUserCredentialsResult, error) {
	var rv ListConnectedClusterUserCredentialsResult
	err := ctx.Invoke("azure-native:kubernetes:listConnectedClusterUserCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConnectedClusterUserCredentialsArgs struct {
	AuthenticationMethod string `pulumi:"authenticationMethod"`
	ClientProxy          bool   `pulumi:"clientProxy"`
	ClusterName          string `pulumi:"clusterName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type ListConnectedClusterUserCredentialsResult struct {
	HybridConnectionConfig HybridConnectionConfigResponse `pulumi:"hybridConnectionConfig"`
	Kubeconfigs            []CredentialResultResponse     `pulumi:"kubeconfigs"`
}

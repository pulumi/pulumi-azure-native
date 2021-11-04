


package resourceconnector

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListApplianceClusterUserCredential(ctx *pulumi.Context, args *ListApplianceClusterUserCredentialArgs, opts ...pulumi.InvokeOption) (*ListApplianceClusterUserCredentialResult, error) {
	var rv ListApplianceClusterUserCredentialResult
	err := ctx.Invoke("azure-native:resourceconnector:listApplianceClusterUserCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListApplianceClusterUserCredentialArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListApplianceClusterUserCredentialResult struct {
	HybridConnectionConfig HybridConnectionConfigResponse          `pulumi:"hybridConnectionConfig"`
	Kubeconfigs            []ApplianceCredentialKubeconfigResponse `pulumi:"kubeconfigs"`
}




package v20200401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListManagedClusterUserCredentials(ctx *pulumi.Context, args *ListManagedClusterUserCredentialsArgs, opts ...pulumi.InvokeOption) (*ListManagedClusterUserCredentialsResult, error) {
	var rv ListManagedClusterUserCredentialsResult
	err := ctx.Invoke("azure-native:containerservice/v20200401:listManagedClusterUserCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagedClusterUserCredentialsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListManagedClusterUserCredentialsResult struct {
	Kubeconfigs []CredentialResultResponse `pulumi:"kubeconfigs"`
}

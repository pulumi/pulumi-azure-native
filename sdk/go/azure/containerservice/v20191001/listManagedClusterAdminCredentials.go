


package v20191001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListManagedClusterAdminCredentials(ctx *pulumi.Context, args *ListManagedClusterAdminCredentialsArgs, opts ...pulumi.InvokeOption) (*ListManagedClusterAdminCredentialsResult, error) {
	var rv ListManagedClusterAdminCredentialsResult
	err := ctx.Invoke("azure-native:containerservice/v20191001:listManagedClusterAdminCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagedClusterAdminCredentialsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListManagedClusterAdminCredentialsResult struct {
	Kubeconfigs []CredentialResultResponse `pulumi:"kubeconfigs"`
}

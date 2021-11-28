


package v20210901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListManagedClusterAdminCredentials(ctx *pulumi.Context, args *ListManagedClusterAdminCredentialsArgs, opts ...pulumi.InvokeOption) (*ListManagedClusterAdminCredentialsResult, error) {
	var rv ListManagedClusterAdminCredentialsResult
	err := ctx.Invoke("azure-native:containerservice/v20210901:listManagedClusterAdminCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagedClusterAdminCredentialsArgs struct {
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ResourceName      string  `pulumi:"resourceName"`
	ServerFqdn        *string `pulumi:"serverFqdn"`
}


type ListManagedClusterAdminCredentialsResult struct {
	Kubeconfigs []CredentialResultResponse `pulumi:"kubeconfigs"`
}

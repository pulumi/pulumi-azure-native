


package v20211001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListManagedClusterUserCredentials(ctx *pulumi.Context, args *ListManagedClusterUserCredentialsArgs, opts ...pulumi.InvokeOption) (*ListManagedClusterUserCredentialsResult, error) {
	var rv ListManagedClusterUserCredentialsResult
	err := ctx.Invoke("azure-native:containerservice/v20211001:listManagedClusterUserCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagedClusterUserCredentialsArgs struct {
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ResourceName      string  `pulumi:"resourceName"`
	ServerFqdn        *string `pulumi:"serverFqdn"`
}


type ListManagedClusterUserCredentialsResult struct {
	Kubeconfigs []CredentialResultResponse `pulumi:"kubeconfigs"`
}

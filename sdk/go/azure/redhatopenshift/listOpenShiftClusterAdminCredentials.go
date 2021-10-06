


package redhatopenshift

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListOpenShiftClusterAdminCredentials(ctx *pulumi.Context, args *ListOpenShiftClusterAdminCredentialsArgs, opts ...pulumi.InvokeOption) (*ListOpenShiftClusterAdminCredentialsResult, error) {
	var rv ListOpenShiftClusterAdminCredentialsResult
	err := ctx.Invoke("azure-native:redhatopenshift:listOpenShiftClusterAdminCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListOpenShiftClusterAdminCredentialsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListOpenShiftClusterAdminCredentialsResult struct {
	Kubeconfig *string `pulumi:"kubeconfig"`
}

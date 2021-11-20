


package v20210201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListManagedClusterMonitoringUserCredentials(ctx *pulumi.Context, args *ListManagedClusterMonitoringUserCredentialsArgs, opts ...pulumi.InvokeOption) (*ListManagedClusterMonitoringUserCredentialsResult, error) {
	var rv ListManagedClusterMonitoringUserCredentialsResult
	err := ctx.Invoke("azure-native:containerservice/v20210201:listManagedClusterMonitoringUserCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagedClusterMonitoringUserCredentialsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListManagedClusterMonitoringUserCredentialsResult struct {
	Kubeconfigs []CredentialResultResponse `pulumi:"kubeconfigs"`
}




package v20200201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListManagedClusterAccessProfile(ctx *pulumi.Context, args *ListManagedClusterAccessProfileArgs, opts ...pulumi.InvokeOption) (*ListManagedClusterAccessProfileResult, error) {
	var rv ListManagedClusterAccessProfileResult
	err := ctx.Invoke("azure-native:containerservice/v20200201:listManagedClusterAccessProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagedClusterAccessProfileArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
	RoleName          string `pulumi:"roleName"`
}


type ListManagedClusterAccessProfileResult struct {
	Id         string            `pulumi:"id"`
	KubeConfig *string           `pulumi:"kubeConfig"`
	Location   string            `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}

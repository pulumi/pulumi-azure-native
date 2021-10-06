


package v20200430

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListOpenShiftClusterCredentials(ctx *pulumi.Context, args *ListOpenShiftClusterCredentialsArgs, opts ...pulumi.InvokeOption) (*ListOpenShiftClusterCredentialsResult, error) {
	var rv ListOpenShiftClusterCredentialsResult
	err := ctx.Invoke("azure-native:redhatopenshift/v20200430:listOpenShiftClusterCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListOpenShiftClusterCredentialsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListOpenShiftClusterCredentialsResult struct {
	KubeadminPassword *string `pulumi:"kubeadminPassword"`
	KubeadminUsername *string `pulumi:"kubeadminUsername"`
}

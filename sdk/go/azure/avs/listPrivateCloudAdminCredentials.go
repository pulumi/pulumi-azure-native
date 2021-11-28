


package avs

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListPrivateCloudAdminCredentials(ctx *pulumi.Context, args *ListPrivateCloudAdminCredentialsArgs, opts ...pulumi.InvokeOption) (*ListPrivateCloudAdminCredentialsResult, error) {
	var rv ListPrivateCloudAdminCredentialsResult
	err := ctx.Invoke("azure-native:avs:listPrivateCloudAdminCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListPrivateCloudAdminCredentialsArgs struct {
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListPrivateCloudAdminCredentialsResult struct {
	NsxtPassword    string `pulumi:"nsxtPassword"`
	NsxtUsername    string `pulumi:"nsxtUsername"`
	VcenterPassword string `pulumi:"vcenterPassword"`
	VcenterUsername string `pulumi:"vcenterUsername"`
}




package v20210201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStaticSiteUsers(ctx *pulumi.Context, args *ListStaticSiteUsersArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteUsersResult, error) {
	var rv ListStaticSiteUsersResult
	err := ctx.Invoke("azure-native:web/v20210201:listStaticSiteUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteUsersArgs struct {
	Authprovider      string `pulumi:"authprovider"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListStaticSiteUsersResult struct {
	NextLink string                              `pulumi:"nextLink"`
	Value    []StaticSiteUserARMResourceResponse `pulumi:"value"`
}

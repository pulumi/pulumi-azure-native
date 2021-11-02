


package v20210101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStaticSiteConfiguredRoles(ctx *pulumi.Context, args *ListStaticSiteConfiguredRolesArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteConfiguredRolesResult, error) {
	var rv ListStaticSiteConfiguredRolesResult
	err := ctx.Invoke("azure-native:web/v20210101:listStaticSiteConfiguredRoles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteConfiguredRolesArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListStaticSiteConfiguredRolesResult struct {
	Id         string   `pulumi:"id"`
	Kind       *string  `pulumi:"kind"`
	Name       string   `pulumi:"name"`
	Properties []string `pulumi:"properties"`
	Type       string   `pulumi:"type"`
}

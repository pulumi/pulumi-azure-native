


package web

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppSiteBackups(ctx *pulumi.Context, args *ListWebAppSiteBackupsArgs, opts ...pulumi.InvokeOption) (*ListWebAppSiteBackupsResult, error) {
	var rv ListWebAppSiteBackupsResult
	err := ctx.Invoke("azure-native:web:listWebAppSiteBackups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppSiteBackupsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppSiteBackupsResult struct {
	NextLink string               `pulumi:"nextLink"`
	Value    []BackupItemResponse `pulumi:"value"`
}

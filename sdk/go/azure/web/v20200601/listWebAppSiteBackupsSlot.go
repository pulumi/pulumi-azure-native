


package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppSiteBackupsSlot(ctx *pulumi.Context, args *ListWebAppSiteBackupsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppSiteBackupsSlotResult, error) {
	var rv ListWebAppSiteBackupsSlotResult
	err := ctx.Invoke("azure-native:web/v20200601:listWebAppSiteBackupsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppSiteBackupsSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppSiteBackupsSlotResult struct {
	NextLink string               `pulumi:"nextLink"`
	Value    []BackupItemResponse `pulumi:"value"`
}

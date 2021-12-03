


package notebooks

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListNotebookProxyCredentials(ctx *pulumi.Context, args *ListNotebookProxyCredentialsArgs, opts ...pulumi.InvokeOption) (*ListNotebookProxyCredentialsResult, error) {
	var rv ListNotebookProxyCredentialsResult
	err := ctx.Invoke("azure-native:notebooks:listNotebookProxyCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNotebookProxyCredentialsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListNotebookProxyCredentialsResult struct {
	Hostname           *string `pulumi:"hostname"`
	PrimaryAccessKey   *string `pulumi:"primaryAccessKey"`
	ResourceId         *string `pulumi:"resourceId"`
	SecondaryAccessKey *string `pulumi:"secondaryAccessKey"`
}

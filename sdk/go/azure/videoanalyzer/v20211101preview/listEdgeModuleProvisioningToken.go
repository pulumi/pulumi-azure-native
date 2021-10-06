


package v20211101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListEdgeModuleProvisioningToken(ctx *pulumi.Context, args *ListEdgeModuleProvisioningTokenArgs, opts ...pulumi.InvokeOption) (*ListEdgeModuleProvisioningTokenResult, error) {
	var rv ListEdgeModuleProvisioningTokenResult
	err := ctx.Invoke("azure-native:videoanalyzer/v20211101preview:listEdgeModuleProvisioningToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEdgeModuleProvisioningTokenArgs struct {
	AccountName       string `pulumi:"accountName"`
	EdgeModuleName    string `pulumi:"edgeModuleName"`
	ExpirationDate    string `pulumi:"expirationDate"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListEdgeModuleProvisioningTokenResult struct {
	ExpirationDate string `pulumi:"expirationDate"`
	Token          string `pulumi:"token"`
}

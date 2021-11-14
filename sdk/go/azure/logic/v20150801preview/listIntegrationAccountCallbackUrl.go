


package v20150801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIntegrationAccountCallbackUrl(ctx *pulumi.Context, args *ListIntegrationAccountCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListIntegrationAccountCallbackUrlResult, error) {
	var rv ListIntegrationAccountCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20150801preview:listIntegrationAccountCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationAccountCallbackUrlArgs struct {
	IntegrationAccountName string  `pulumi:"integrationAccountName"`
	NotAfter               *string `pulumi:"notAfter"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
}

type ListIntegrationAccountCallbackUrlResult struct {
	Value *string `pulumi:"value"`
}

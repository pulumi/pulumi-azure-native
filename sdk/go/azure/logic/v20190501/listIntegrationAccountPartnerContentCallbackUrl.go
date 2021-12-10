


package v20190501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIntegrationAccountPartnerContentCallbackUrl(ctx *pulumi.Context, args *ListIntegrationAccountPartnerContentCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListIntegrationAccountPartnerContentCallbackUrlResult, error) {
	var rv ListIntegrationAccountPartnerContentCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20190501:listIntegrationAccountPartnerContentCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationAccountPartnerContentCallbackUrlArgs struct {
	IntegrationAccountName string  `pulumi:"integrationAccountName"`
	KeyType                *string `pulumi:"keyType"`
	NotAfter               *string `pulumi:"notAfter"`
	PartnerName            string  `pulumi:"partnerName"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
}


type ListIntegrationAccountPartnerContentCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

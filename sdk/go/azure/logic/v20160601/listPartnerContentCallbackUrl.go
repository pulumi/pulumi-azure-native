


package v20160601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListPartnerContentCallbackUrl(ctx *pulumi.Context, args *ListPartnerContentCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListPartnerContentCallbackUrlResult, error) {
	var rv ListPartnerContentCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20160601:listPartnerContentCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListPartnerContentCallbackUrlArgs struct {
	IntegrationAccountName string   `pulumi:"integrationAccountName"`
	KeyType                *KeyType `pulumi:"keyType"`
	NotAfter               *string  `pulumi:"notAfter"`
	PartnerName            string   `pulumi:"partnerName"`
	ResourceGroupName      string   `pulumi:"resourceGroupName"`
}


type ListPartnerContentCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}




package logic

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIntegrationAccountAgreementContentCallbackUrl(ctx *pulumi.Context, args *ListIntegrationAccountAgreementContentCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListIntegrationAccountAgreementContentCallbackUrlResult, error) {
	var rv ListIntegrationAccountAgreementContentCallbackUrlResult
	err := ctx.Invoke("azure-native:logic:listIntegrationAccountAgreementContentCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationAccountAgreementContentCallbackUrlArgs struct {
	AgreementName          string  `pulumi:"agreementName"`
	IntegrationAccountName string  `pulumi:"integrationAccountName"`
	KeyType                *string `pulumi:"keyType"`
	NotAfter               *string `pulumi:"notAfter"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
}


type ListIntegrationAccountAgreementContentCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

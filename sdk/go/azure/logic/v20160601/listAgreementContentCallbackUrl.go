


package v20160601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAgreementContentCallbackUrl(ctx *pulumi.Context, args *ListAgreementContentCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListAgreementContentCallbackUrlResult, error) {
	var rv ListAgreementContentCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20160601:listAgreementContentCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAgreementContentCallbackUrlArgs struct {
	AgreementName          string   `pulumi:"agreementName"`
	IntegrationAccountName string   `pulumi:"integrationAccountName"`
	KeyType                *KeyType `pulumi:"keyType"`
	NotAfter               *string  `pulumi:"notAfter"`
	ResourceGroupName      string   `pulumi:"resourceGroupName"`
}


type ListAgreementContentCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

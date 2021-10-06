


package v20200831

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIotHubResourceKeys(ctx *pulumi.Context, args *ListIotHubResourceKeysArgs, opts ...pulumi.InvokeOption) (*ListIotHubResourceKeysResult, error) {
	var rv ListIotHubResourceKeysResult
	err := ctx.Invoke("azure-native:devices/v20200831:listIotHubResourceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIotHubResourceKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListIotHubResourceKeysResult struct {
	NextLink string                                           `pulumi:"nextLink"`
	Value    []SharedAccessSignatureAuthorizationRuleResponse `pulumi:"value"`
}

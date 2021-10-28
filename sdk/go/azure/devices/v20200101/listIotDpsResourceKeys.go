


package v20200101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIotDpsResourceKeys(ctx *pulumi.Context, args *ListIotDpsResourceKeysArgs, opts ...pulumi.InvokeOption) (*ListIotDpsResourceKeysResult, error) {
	var rv ListIotDpsResourceKeysResult
	err := ctx.Invoke("azure-native:devices/v20200101:listIotDpsResourceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIotDpsResourceKeysArgs struct {
	ProvisioningServiceName string `pulumi:"provisioningServiceName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type ListIotDpsResourceKeysResult struct {
	NextLink string                                                                  `pulumi:"nextLink"`
	Value    []SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse `pulumi:"value"`
}

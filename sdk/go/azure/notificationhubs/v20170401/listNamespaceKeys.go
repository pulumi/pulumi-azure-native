


package v20170401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListNamespaceKeys(ctx *pulumi.Context, args *ListNamespaceKeysArgs, opts ...pulumi.InvokeOption) (*ListNamespaceKeysResult, error) {
	var rv ListNamespaceKeysResult
	err := ctx.Invoke("azure-native:notificationhubs/v20170401:listNamespaceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNamespaceKeysArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	NamespaceName         string `pulumi:"namespaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type ListNamespaceKeysResult struct {
	NextLink *string                                         `pulumi:"nextLink"`
	Value    []SharedAccessAuthorizationRuleResourceResponse `pulumi:"value"`
}

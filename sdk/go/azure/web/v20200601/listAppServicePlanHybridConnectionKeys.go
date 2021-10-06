


package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAppServicePlanHybridConnectionKeys(ctx *pulumi.Context, args *ListAppServicePlanHybridConnectionKeysArgs, opts ...pulumi.InvokeOption) (*ListAppServicePlanHybridConnectionKeysResult, error) {
	var rv ListAppServicePlanHybridConnectionKeysResult
	err := ctx.Invoke("azure-native:web/v20200601:listAppServicePlanHybridConnectionKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAppServicePlanHybridConnectionKeysArgs struct {
	Name              string `pulumi:"name"`
	NamespaceName     string `pulumi:"namespaceName"`
	RelayName         string `pulumi:"relayName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListAppServicePlanHybridConnectionKeysResult struct {
	Id           string  `pulumi:"id"`
	Kind         *string `pulumi:"kind"`
	Name         string  `pulumi:"name"`
	SendKeyName  string  `pulumi:"sendKeyName"`
	SendKeyValue string  `pulumi:"sendKeyValue"`
	Type         string  `pulumi:"type"`
}

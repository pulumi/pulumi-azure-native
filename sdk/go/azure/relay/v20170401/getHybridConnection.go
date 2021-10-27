


package v20170401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHybridConnection(ctx *pulumi.Context, args *LookupHybridConnectionArgs, opts ...pulumi.InvokeOption) (*LookupHybridConnectionResult, error) {
	var rv LookupHybridConnectionResult
	err := ctx.Invoke("azure-native:relay/v20170401:getHybridConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHybridConnectionArgs struct {
	HybridConnectionName string `pulumi:"hybridConnectionName"`
	NamespaceName        string `pulumi:"namespaceName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupHybridConnectionResult struct {
	CreatedAt                   string  `pulumi:"createdAt"`
	Id                          string  `pulumi:"id"`
	ListenerCount               int     `pulumi:"listenerCount"`
	Name                        string  `pulumi:"name"`
	RequiresClientAuthorization *bool   `pulumi:"requiresClientAuthorization"`
	Type                        string  `pulumi:"type"`
	UpdatedAt                   string  `pulumi:"updatedAt"`
	UserMetadata                *string `pulumi:"userMetadata"`
}

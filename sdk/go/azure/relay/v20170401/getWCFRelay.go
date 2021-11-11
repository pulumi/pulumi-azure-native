


package v20170401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWCFRelay(ctx *pulumi.Context, args *LookupWCFRelayArgs, opts ...pulumi.InvokeOption) (*LookupWCFRelayResult, error) {
	var rv LookupWCFRelayResult
	err := ctx.Invoke("azure-native:relay/v20170401:getWCFRelay", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWCFRelayArgs struct {
	NamespaceName     string `pulumi:"namespaceName"`
	RelayName         string `pulumi:"relayName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWCFRelayResult struct {
	CreatedAt                   string  `pulumi:"createdAt"`
	Id                          string  `pulumi:"id"`
	IsDynamic                   bool    `pulumi:"isDynamic"`
	ListenerCount               int     `pulumi:"listenerCount"`
	Name                        string  `pulumi:"name"`
	RelayType                   *string `pulumi:"relayType"`
	RequiresClientAuthorization *bool   `pulumi:"requiresClientAuthorization"`
	RequiresTransportSecurity   *bool   `pulumi:"requiresTransportSecurity"`
	Type                        string  `pulumi:"type"`
	UpdatedAt                   string  `pulumi:"updatedAt"`
	UserMetadata                *string `pulumi:"userMetadata"`
}

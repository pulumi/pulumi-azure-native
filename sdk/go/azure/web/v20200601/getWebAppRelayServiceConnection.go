


package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppRelayServiceConnection(ctx *pulumi.Context, args *LookupWebAppRelayServiceConnectionArgs, opts ...pulumi.InvokeOption) (*LookupWebAppRelayServiceConnectionResult, error) {
	var rv LookupWebAppRelayServiceConnectionResult
	err := ctx.Invoke("azure-native:web/v20200601:getWebAppRelayServiceConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppRelayServiceConnectionArgs struct {
	EntityName        string `pulumi:"entityName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWebAppRelayServiceConnectionResult struct {
	BiztalkUri               *string `pulumi:"biztalkUri"`
	EntityConnectionString   *string `pulumi:"entityConnectionString"`
	EntityName               *string `pulumi:"entityName"`
	Hostname                 *string `pulumi:"hostname"`
	Id                       string  `pulumi:"id"`
	Kind                     *string `pulumi:"kind"`
	Name                     string  `pulumi:"name"`
	Port                     *int    `pulumi:"port"`
	ResourceConnectionString *string `pulumi:"resourceConnectionString"`
	ResourceType             *string `pulumi:"resourceType"`
	Type                     string  `pulumi:"type"`
}

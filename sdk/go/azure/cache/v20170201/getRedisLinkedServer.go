


package v20170201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRedisLinkedServer(ctx *pulumi.Context, args *LookupRedisLinkedServerArgs, opts ...pulumi.InvokeOption) (*LookupRedisLinkedServerResult, error) {
	var rv LookupRedisLinkedServerResult
	err := ctx.Invoke("azure-native:cache/v20170201:getRedisLinkedServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRedisLinkedServerArgs struct {
	LinkedServerName  string `pulumi:"linkedServerName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRedisLinkedServerResult struct {
	Id                       string `pulumi:"id"`
	LinkedRedisCacheId       string `pulumi:"linkedRedisCacheId"`
	LinkedRedisCacheLocation string `pulumi:"linkedRedisCacheLocation"`
	Name                     string `pulumi:"name"`
	ProvisioningState        string `pulumi:"provisioningState"`
	ServerRole               string `pulumi:"serverRole"`
	Type                     string `pulumi:"type"`
}




package v20171001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLinkedServer(ctx *pulumi.Context, args *LookupLinkedServerArgs, opts ...pulumi.InvokeOption) (*LookupLinkedServerResult, error) {
	var rv LookupLinkedServerResult
	err := ctx.Invoke("azure-native:cache/v20171001:getLinkedServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkedServerArgs struct {
	LinkedServerName  string `pulumi:"linkedServerName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLinkedServerResult struct {
	Id                       string `pulumi:"id"`
	LinkedRedisCacheId       string `pulumi:"linkedRedisCacheId"`
	LinkedRedisCacheLocation string `pulumi:"linkedRedisCacheLocation"`
	Name                     string `pulumi:"name"`
	ProvisioningState        string `pulumi:"provisioningState"`
	ServerRole               string `pulumi:"serverRole"`
	Type                     string `pulumi:"type"`
}

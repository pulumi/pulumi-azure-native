


package v20211001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupResourceGuardProxy(ctx *pulumi.Context, args *LookupResourceGuardProxyArgs, opts ...pulumi.InvokeOption) (*LookupResourceGuardProxyResult, error) {
	var rv LookupResourceGuardProxyResult
	err := ctx.Invoke("azure-native:recoveryservices/v20211001:getResourceGuardProxy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceGuardProxyArgs struct {
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	ResourceGuardProxyName string `pulumi:"resourceGuardProxyName"`
	VaultName              string `pulumi:"vaultName"`
}

type LookupResourceGuardProxyResult struct {
	ETag       *string                        `pulumi:"eTag"`
	Id         string                         `pulumi:"id"`
	Location   *string                        `pulumi:"location"`
	Name       string                         `pulumi:"name"`
	Properties ResourceGuardProxyBaseResponse `pulumi:"properties"`
	Tags       map[string]string              `pulumi:"tags"`
	Type       string                         `pulumi:"type"`
}




package v20180901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecret(ctx *pulumi.Context, args *LookupSecretArgs, opts ...pulumi.InvokeOption) (*LookupSecretResult, error) {
	var rv LookupSecretResult
	err := ctx.Invoke("azure-native:servicefabricmesh/v20180901preview:getSecret", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecretArgs struct {
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	SecretResourceName string `pulumi:"secretResourceName"`
}


type LookupSecretResult struct {
	Id         string                           `pulumi:"id"`
	Location   string                           `pulumi:"location"`
	Name       string                           `pulumi:"name"`
	Properties SecretResourcePropertiesResponse `pulumi:"properties"`
	Tags       map[string]string                `pulumi:"tags"`
	Type       string                           `pulumi:"type"`
}

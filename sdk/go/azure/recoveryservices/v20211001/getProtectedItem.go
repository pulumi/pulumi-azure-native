


package v20211001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProtectedItem(ctx *pulumi.Context, args *LookupProtectedItemArgs, opts ...pulumi.InvokeOption) (*LookupProtectedItemResult, error) {
	var rv LookupProtectedItemResult
	err := ctx.Invoke("azure-native:recoveryservices/v20211001:getProtectedItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProtectedItemArgs struct {
	ContainerName     string  `pulumi:"containerName"`
	FabricName        string  `pulumi:"fabricName"`
	Filter            *string `pulumi:"filter"`
	ProtectedItemName string  `pulumi:"protectedItemName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	VaultName         string  `pulumi:"vaultName"`
}


type LookupProtectedItemResult struct {
	ETag       *string           `pulumi:"eTag"`
	Id         string            `pulumi:"id"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}




package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProtectionContainer(ctx *pulumi.Context, args *LookupProtectionContainerArgs, opts ...pulumi.InvokeOption) (*LookupProtectionContainerResult, error) {
	var rv LookupProtectionContainerResult
	err := ctx.Invoke("azure-native:recoveryservices/v20210201preview:getProtectionContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProtectionContainerArgs struct {
	ContainerName     string `pulumi:"containerName"`
	FabricName        string `pulumi:"fabricName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VaultName         string `pulumi:"vaultName"`
}


type LookupProtectionContainerResult struct {
	ETag       *string           `pulumi:"eTag"`
	Id         string            `pulumi:"id"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}




package v20210201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProtectionIntent(ctx *pulumi.Context, args *LookupProtectionIntentArgs, opts ...pulumi.InvokeOption) (*LookupProtectionIntentResult, error) {
	var rv LookupProtectionIntentResult
	err := ctx.Invoke("azure-native:recoveryservices/v20210201:getProtectionIntent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProtectionIntentArgs struct {
	FabricName        string `pulumi:"fabricName"`
	IntentObjectName  string `pulumi:"intentObjectName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VaultName         string `pulumi:"vaultName"`
}


type LookupProtectionIntentResult struct {
	ETag       *string           `pulumi:"eTag"`
	Id         string            `pulumi:"id"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}




package devtestlab

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecret(ctx *pulumi.Context, args *LookupSecretArgs, opts ...pulumi.InvokeOption) (*LookupSecretResult, error) {
	var rv LookupSecretResult
	err := ctx.Invoke("azure-native:devtestlab:getSecret", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecretArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	UserName          string  `pulumi:"userName"`
}


type LookupSecretResult struct {
	Id                string            `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
	UniqueIdentifier  string            `pulumi:"uniqueIdentifier"`
	Value             *string           `pulumi:"value"`
}

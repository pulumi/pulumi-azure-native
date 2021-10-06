


package devtestlab

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	var rv LookupUserResult
	err := ctx.Invoke("azure-native:devtestlab:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupUserResult struct {
	CreatedDate       string                   `pulumi:"createdDate"`
	Id                string                   `pulumi:"id"`
	Identity          *UserIdentityResponse    `pulumi:"identity"`
	Location          *string                  `pulumi:"location"`
	Name              string                   `pulumi:"name"`
	ProvisioningState string                   `pulumi:"provisioningState"`
	SecretStore       *UserSecretStoreResponse `pulumi:"secretStore"`
	Tags              map[string]string        `pulumi:"tags"`
	Type              string                   `pulumi:"type"`
	UniqueIdentifier  string                   `pulumi:"uniqueIdentifier"`
}

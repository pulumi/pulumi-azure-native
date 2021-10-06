


package v20210101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppDomainOwnershipIdentifier(ctx *pulumi.Context, args *LookupWebAppDomainOwnershipIdentifierArgs, opts ...pulumi.InvokeOption) (*LookupWebAppDomainOwnershipIdentifierResult, error) {
	var rv LookupWebAppDomainOwnershipIdentifierResult
	err := ctx.Invoke("azure-native:web/v20210101:getWebAppDomainOwnershipIdentifier", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppDomainOwnershipIdentifierArgs struct {
	DomainOwnershipIdentifierName string `pulumi:"domainOwnershipIdentifierName"`
	Name                          string `pulumi:"name"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
}


type LookupWebAppDomainOwnershipIdentifierResult struct {
	Id    string  `pulumi:"id"`
	Kind  *string `pulumi:"kind"`
	Name  string  `pulumi:"name"`
	Type  string  `pulumi:"type"`
	Value *string `pulumi:"value"`
}




package v20210101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppDomainOwnershipIdentifierSlot(ctx *pulumi.Context, args *LookupWebAppDomainOwnershipIdentifierSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppDomainOwnershipIdentifierSlotResult, error) {
	var rv LookupWebAppDomainOwnershipIdentifierSlotResult
	err := ctx.Invoke("azure-native:web/v20210101:getWebAppDomainOwnershipIdentifierSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppDomainOwnershipIdentifierSlotArgs struct {
	DomainOwnershipIdentifierName string `pulumi:"domainOwnershipIdentifierName"`
	Name                          string `pulumi:"name"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	Slot                          string `pulumi:"slot"`
}


type LookupWebAppDomainOwnershipIdentifierSlotResult struct {
	Id    string  `pulumi:"id"`
	Kind  *string `pulumi:"kind"`
	Name  string  `pulumi:"name"`
	Type  string  `pulumi:"type"`
	Value *string `pulumi:"value"`
}

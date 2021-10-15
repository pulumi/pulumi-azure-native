


package v20201001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDomainOwnershipIdentifier(ctx *pulumi.Context, args *LookupDomainOwnershipIdentifierArgs, opts ...pulumi.InvokeOption) (*LookupDomainOwnershipIdentifierResult, error) {
	var rv LookupDomainOwnershipIdentifierResult
	err := ctx.Invoke("azure-native:domainregistration/v20201001:getDomainOwnershipIdentifier", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDomainOwnershipIdentifierArgs struct {
	DomainName        string `pulumi:"domainName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDomainOwnershipIdentifierResult struct {
	Id          string             `pulumi:"id"`
	Kind        *string            `pulumi:"kind"`
	Name        string             `pulumi:"name"`
	OwnershipId *string            `pulumi:"ownershipId"`
	SystemData  SystemDataResponse `pulumi:"systemData"`
	Type        string             `pulumi:"type"`
}

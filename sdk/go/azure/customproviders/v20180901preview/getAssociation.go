


package v20180901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAssociation(ctx *pulumi.Context, args *LookupAssociationArgs, opts ...pulumi.InvokeOption) (*LookupAssociationResult, error) {
	var rv LookupAssociationResult
	err := ctx.Invoke("azure-native:customproviders/v20180901preview:getAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssociationArgs struct {
	AssociationName string `pulumi:"associationName"`
	Scope           string `pulumi:"scope"`
}


type LookupAssociationResult struct {
	Id                string  `pulumi:"id"`
	Name              string  `pulumi:"name"`
	ProvisioningState string  `pulumi:"provisioningState"`
	TargetResourceId  *string `pulumi:"targetResourceId"`
	Type              string  `pulumi:"type"`
}

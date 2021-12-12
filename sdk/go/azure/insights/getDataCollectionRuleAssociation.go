


package insights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataCollectionRuleAssociation(ctx *pulumi.Context, args *LookupDataCollectionRuleAssociationArgs, opts ...pulumi.InvokeOption) (*LookupDataCollectionRuleAssociationResult, error) {
	var rv LookupDataCollectionRuleAssociationResult
	err := ctx.Invoke("azure-native:insights:getDataCollectionRuleAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataCollectionRuleAssociationArgs struct {
	AssociationName string `pulumi:"associationName"`
	ResourceUri     string `pulumi:"resourceUri"`
}


type LookupDataCollectionRuleAssociationResult struct {
	DataCollectionRuleId *string `pulumi:"dataCollectionRuleId"`
	Description          *string `pulumi:"description"`
	Etag                 string  `pulumi:"etag"`
	Id                   string  `pulumi:"id"`
	Name                 string  `pulumi:"name"`
	ProvisioningState    string  `pulumi:"provisioningState"`
	Type                 string  `pulumi:"type"`
}

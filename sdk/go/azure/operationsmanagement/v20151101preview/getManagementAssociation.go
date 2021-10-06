


package v20151101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagementAssociation(ctx *pulumi.Context, args *LookupManagementAssociationArgs, opts ...pulumi.InvokeOption) (*LookupManagementAssociationResult, error) {
	var rv LookupManagementAssociationResult
	err := ctx.Invoke("azure-native:operationsmanagement/v20151101preview:getManagementAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementAssociationArgs struct {
	ManagementAssociationName string `pulumi:"managementAssociationName"`
	ProviderName              string `pulumi:"providerName"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	ResourceName              string `pulumi:"resourceName"`
	ResourceType              string `pulumi:"resourceType"`
}


type LookupManagementAssociationResult struct {
	Id         string                                  `pulumi:"id"`
	Location   *string                                 `pulumi:"location"`
	Name       string                                  `pulumi:"name"`
	Properties ManagementAssociationPropertiesResponse `pulumi:"properties"`
	Type       string                                  `pulumi:"type"`
}

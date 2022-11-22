


package v20220401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetApiManagementServiceDomainOwnershipIdentifier(ctx *pulumi.Context, args *GetApiManagementServiceDomainOwnershipIdentifierArgs, opts ...pulumi.InvokeOption) (*GetApiManagementServiceDomainOwnershipIdentifierResult, error) {
	var rv GetApiManagementServiceDomainOwnershipIdentifierResult
	err := ctx.Invoke("azure-native:apimanagement/v20220401preview:getApiManagementServiceDomainOwnershipIdentifier", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetApiManagementServiceDomainOwnershipIdentifierArgs struct {
}


type GetApiManagementServiceDomainOwnershipIdentifierResult struct {
	DomainOwnershipIdentifier string `pulumi:"domainOwnershipIdentifier"`
}

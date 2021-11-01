


package confluent

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOrganization(ctx *pulumi.Context, args *LookupOrganizationArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationResult, error) {
	var rv LookupOrganizationResult
	err := ctx.Invoke("azure-native:confluent:getOrganization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrganizationArgs struct {
	OrganizationName  string `pulumi:"organizationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupOrganizationResult struct {
	CreatedTime       string                                             `pulumi:"createdTime"`
	Id                string                                             `pulumi:"id"`
	Location          *string                                            `pulumi:"location"`
	Name              string                                             `pulumi:"name"`
	OfferDetail       *OrganizationResourcePropertiesResponseOfferDetail `pulumi:"offerDetail"`
	OrganizationId    string                                             `pulumi:"organizationId"`
	ProvisioningState string                                             `pulumi:"provisioningState"`
	SsoUrl            string                                             `pulumi:"ssoUrl"`
	Tags              map[string]string                                  `pulumi:"tags"`
	Type              string                                             `pulumi:"type"`
	UserDetail        *OrganizationResourcePropertiesResponseUserDetail  `pulumi:"userDetail"`
}

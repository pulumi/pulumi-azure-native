


package eventgrid

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPartnerRegistration(ctx *pulumi.Context, args *LookupPartnerRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupPartnerRegistrationResult, error) {
	var rv LookupPartnerRegistrationResult
	err := ctx.Invoke("azure-native:eventgrid:getPartnerRegistration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPartnerRegistrationArgs struct {
	PartnerRegistrationName string `pulumi:"partnerRegistrationName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupPartnerRegistrationResult struct {
	AuthorizedAzureSubscriptionIds  []string           `pulumi:"authorizedAzureSubscriptionIds"`
	CustomerServiceUri              *string            `pulumi:"customerServiceUri"`
	Id                              string             `pulumi:"id"`
	Location                        string             `pulumi:"location"`
	LogoUri                         *string            `pulumi:"logoUri"`
	LongDescription                 *string            `pulumi:"longDescription"`
	Name                            string             `pulumi:"name"`
	PartnerCustomerServiceExtension *string            `pulumi:"partnerCustomerServiceExtension"`
	PartnerCustomerServiceNumber    *string            `pulumi:"partnerCustomerServiceNumber"`
	PartnerName                     *string            `pulumi:"partnerName"`
	PartnerResourceTypeDescription  *string            `pulumi:"partnerResourceTypeDescription"`
	PartnerResourceTypeDisplayName  *string            `pulumi:"partnerResourceTypeDisplayName"`
	PartnerResourceTypeName         *string            `pulumi:"partnerResourceTypeName"`
	ProvisioningState               string             `pulumi:"provisioningState"`
	SetupUri                        *string            `pulumi:"setupUri"`
	SystemData                      SystemDataResponse `pulumi:"systemData"`
	Tags                            map[string]string  `pulumi:"tags"`
	Type                            string             `pulumi:"type"`
	VisibilityState                 *string            `pulumi:"visibilityState"`
}

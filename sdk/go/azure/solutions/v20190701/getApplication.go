


package v20190701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("azure-native:solutions/v20190701:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	ApplicationName   string `pulumi:"applicationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupApplicationResult struct {
	ApplicationDefinitionId *string                                     `pulumi:"applicationDefinitionId"`
	Artifacts               []ApplicationArtifactResponse               `pulumi:"artifacts"`
	Authorizations          []ApplicationAuthorizationResponse          `pulumi:"authorizations"`
	BillingDetails          ApplicationBillingDetailsDefinitionResponse `pulumi:"billingDetails"`
	CreatedBy               ApplicationClientDetailsResponse            `pulumi:"createdBy"`
	CustomerSupport         ApplicationPackageContactResponse           `pulumi:"customerSupport"`
	Id                      string                                      `pulumi:"id"`
	Identity                *IdentityResponse                           `pulumi:"identity"`
	JitAccessPolicy         *ApplicationJitAccessPolicyResponse         `pulumi:"jitAccessPolicy"`
	Kind                    string                                      `pulumi:"kind"`
	Location                *string                                     `pulumi:"location"`
	ManagedBy               *string                                     `pulumi:"managedBy"`
	ManagedResourceGroupId  *string                                     `pulumi:"managedResourceGroupId"`
	ManagementMode          string                                      `pulumi:"managementMode"`
	Name                    string                                      `pulumi:"name"`
	Outputs                 interface{}                                 `pulumi:"outputs"`
	Parameters              interface{}                                 `pulumi:"parameters"`
	Plan                    *PlanResponse                               `pulumi:"plan"`
	ProvisioningState       string                                      `pulumi:"provisioningState"`
	PublisherTenantId       string                                      `pulumi:"publisherTenantId"`
	Sku                     *SkuResponse                                `pulumi:"sku"`
	SupportUrls             ApplicationPackageSupportUrlsResponse       `pulumi:"supportUrls"`
	Tags                    map[string]string                           `pulumi:"tags"`
	Type                    string                                      `pulumi:"type"`
	UpdatedBy               ApplicationClientDetailsResponse            `pulumi:"updatedBy"`
}




package v20160901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplianceDefinition(ctx *pulumi.Context, args *LookupApplianceDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupApplianceDefinitionResult, error) {
	var rv LookupApplianceDefinitionResult
	err := ctx.Invoke("azure-native:solutions/v20160901preview:getApplianceDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplianceDefinitionArgs struct {
	ApplianceDefinitionName string `pulumi:"applianceDefinitionName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupApplianceDefinitionResult struct {
	Artifacts      []ApplianceArtifactResponse              `pulumi:"artifacts"`
	Authorizations []ApplianceProviderAuthorizationResponse `pulumi:"authorizations"`
	Description    *string                                  `pulumi:"description"`
	DisplayName    *string                                  `pulumi:"displayName"`
	Id             string                                   `pulumi:"id"`
	Identity       *IdentityResponse                        `pulumi:"identity"`
	Location       *string                                  `pulumi:"location"`
	LockLevel      string                                   `pulumi:"lockLevel"`
	ManagedBy      *string                                  `pulumi:"managedBy"`
	Name           string                                   `pulumi:"name"`
	PackageFileUri string                                   `pulumi:"packageFileUri"`
	Sku            *SkuResponse                             `pulumi:"sku"`
	Tags           map[string]string                        `pulumi:"tags"`
	Type           string                                   `pulumi:"type"`
}

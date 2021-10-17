


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupToken(ctx *pulumi.Context, args *LookupTokenArgs, opts ...pulumi.InvokeOption) (*LookupTokenResult, error) {
	var rv LookupTokenResult
	err := ctx.Invoke("azure-native:containerregistry/v20210601preview:getToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTokenArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TokenName         string `pulumi:"tokenName"`
}


type LookupTokenResult struct {
	CreationDate      string                              `pulumi:"creationDate"`
	Credentials       *TokenCredentialsPropertiesResponse `pulumi:"credentials"`
	Id                string                              `pulumi:"id"`
	Name              string                              `pulumi:"name"`
	ProvisioningState string                              `pulumi:"provisioningState"`
	ScopeMapId        *string                             `pulumi:"scopeMapId"`
	Status            *string                             `pulumi:"status"`
	SystemData        SystemDataResponse                  `pulumi:"systemData"`
	Type              string                              `pulumi:"type"`
}

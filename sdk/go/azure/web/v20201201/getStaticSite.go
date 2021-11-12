


package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStaticSite(ctx *pulumi.Context, args *LookupStaticSiteArgs, opts ...pulumi.InvokeOption) (*LookupStaticSiteResult, error) {
	var rv LookupStaticSiteResult
	err := ctx.Invoke("azure-native:web/v20201201:getStaticSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStaticSiteArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupStaticSiteResult struct {
	AllowConfigFileUpdates      *bool                                                            `pulumi:"allowConfigFileUpdates"`
	Branch                      *string                                                          `pulumi:"branch"`
	BuildProperties             *StaticSiteBuildPropertiesResponse                               `pulumi:"buildProperties"`
	ContentDistributionEndpoint string                                                           `pulumi:"contentDistributionEndpoint"`
	CustomDomains               []string                                                         `pulumi:"customDomains"`
	DefaultHostname             string                                                           `pulumi:"defaultHostname"`
	Id                          string                                                           `pulumi:"id"`
	Identity                    *ManagedServiceIdentityResponse                                  `pulumi:"identity"`
	KeyVaultReferenceIdentity   string                                                           `pulumi:"keyVaultReferenceIdentity"`
	Kind                        *string                                                          `pulumi:"kind"`
	Location                    string                                                           `pulumi:"location"`
	Name                        string                                                           `pulumi:"name"`
	PrivateEndpointConnections  []ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	Provider                    string                                                           `pulumi:"provider"`
	RepositoryToken             *string                                                          `pulumi:"repositoryToken"`
	RepositoryUrl               *string                                                          `pulumi:"repositoryUrl"`
	Sku                         *SkuDescriptionResponse                                          `pulumi:"sku"`
	StagingEnvironmentPolicy    *string                                                          `pulumi:"stagingEnvironmentPolicy"`
	Tags                        map[string]string                                                `pulumi:"tags"`
	TemplateProperties          *StaticSiteTemplateOptionsResponse                               `pulumi:"templateProperties"`
	Type                        string                                                           `pulumi:"type"`
	UserProvidedFunctionApps    []StaticSiteUserProvidedFunctionAppResponse                      `pulumi:"userProvidedFunctionApps"`
}

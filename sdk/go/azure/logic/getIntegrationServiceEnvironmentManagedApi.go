


package logic

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationServiceEnvironmentManagedApi(ctx *pulumi.Context, args *LookupIntegrationServiceEnvironmentManagedApiArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationServiceEnvironmentManagedApiResult, error) {
	var rv LookupIntegrationServiceEnvironmentManagedApiResult
	err := ctx.Invoke("azure-native:logic:getIntegrationServiceEnvironmentManagedApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationServiceEnvironmentManagedApiArgs struct {
	ApiName                           string `pulumi:"apiName"`
	IntegrationServiceEnvironmentName string `pulumi:"integrationServiceEnvironmentName"`
	ResourceGroup                     string `pulumi:"resourceGroup"`
}


type LookupIntegrationServiceEnvironmentManagedApiResult struct {
	ApiDefinitionUrl              string                                                               `pulumi:"apiDefinitionUrl"`
	ApiDefinitions                ApiResourceDefinitionsResponse                                       `pulumi:"apiDefinitions"`
	BackendService                ApiResourceBackendServiceResponse                                    `pulumi:"backendService"`
	Capabilities                  []string                                                             `pulumi:"capabilities"`
	Category                      string                                                               `pulumi:"category"`
	ConnectionParameters          map[string]interface{}                                               `pulumi:"connectionParameters"`
	DeploymentParameters          *IntegrationServiceEnvironmentManagedApiDeploymentParametersResponse `pulumi:"deploymentParameters"`
	GeneralInformation            ApiResourceGeneralInformationResponse                                `pulumi:"generalInformation"`
	Id                            string                                                               `pulumi:"id"`
	IntegrationServiceEnvironment *ResourceReferenceResponse                                           `pulumi:"integrationServiceEnvironment"`
	Location                      *string                                                              `pulumi:"location"`
	Metadata                      ApiResourceMetadataResponse                                          `pulumi:"metadata"`
	Name                          string                                                               `pulumi:"name"`
	Policies                      ApiResourcePoliciesResponse                                          `pulumi:"policies"`
	ProvisioningState             string                                                               `pulumi:"provisioningState"`
	RuntimeUrls                   []string                                                             `pulumi:"runtimeUrls"`
	Tags                          map[string]string                                                    `pulumi:"tags"`
	Type                          string                                                               `pulumi:"type"`
}

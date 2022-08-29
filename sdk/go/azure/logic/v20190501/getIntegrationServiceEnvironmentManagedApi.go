


package v20190501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationServiceEnvironmentManagedApi(ctx *pulumi.Context, args *LookupIntegrationServiceEnvironmentManagedApiArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationServiceEnvironmentManagedApiResult, error) {
	var rv LookupIntegrationServiceEnvironmentManagedApiResult
	err := ctx.Invoke("azure-native:logic/v20190501:getIntegrationServiceEnvironmentManagedApi", args, &rv, opts...)
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

func LookupIntegrationServiceEnvironmentManagedApiOutput(ctx *pulumi.Context, args LookupIntegrationServiceEnvironmentManagedApiOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationServiceEnvironmentManagedApiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntegrationServiceEnvironmentManagedApiResult, error) {
			args := v.(LookupIntegrationServiceEnvironmentManagedApiArgs)
			r, err := LookupIntegrationServiceEnvironmentManagedApi(ctx, &args, opts...)
			var s LookupIntegrationServiceEnvironmentManagedApiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntegrationServiceEnvironmentManagedApiResultOutput)
}

type LookupIntegrationServiceEnvironmentManagedApiOutputArgs struct {
	ApiName                           pulumi.StringInput `pulumi:"apiName"`
	IntegrationServiceEnvironmentName pulumi.StringInput `pulumi:"integrationServiceEnvironmentName"`
	ResourceGroup                     pulumi.StringInput `pulumi:"resourceGroup"`
}

func (LookupIntegrationServiceEnvironmentManagedApiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationServiceEnvironmentManagedApiArgs)(nil)).Elem()
}


type LookupIntegrationServiceEnvironmentManagedApiResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationServiceEnvironmentManagedApiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationServiceEnvironmentManagedApiResult)(nil)).Elem()
}

func (o LookupIntegrationServiceEnvironmentManagedApiResultOutput) ToLookupIntegrationServiceEnvironmentManagedApiResultOutput() LookupIntegrationServiceEnvironmentManagedApiResultOutput {
	return o
}

func (o LookupIntegrationServiceEnvironmentManagedApiResultOutput) ToLookupIntegrationServiceEnvironmentManagedApiResultOutputWithContext(ctx context.Context) LookupIntegrationServiceEnvironmentManagedApiResultOutput {
	return o
}

func (o LookupIntegrationServiceEnvironmentManagedApiResultOutput) ApiDefinitionUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentManagedApiResult) string { return v.ApiDefinitionUrl }).(pulumi.StringOutput)
}

func (o LookupIntegrationServiceEnvironmentManagedApiResultOutput) ApiDefinitions() ApiResourceDefinitionsResponseOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentManagedApiResult) ApiResourceDefinitionsResponse {
		return v.ApiDefinitions
	}).(ApiResourceDefinitionsResponseOutput)
}

func (o LookupIntegrationServiceEnvironmentManagedApiResultOutput) BackendService() ApiResourceBackendServiceResponseOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentManagedApiResult) ApiResourceBackendServiceResponse {
		return v.BackendService
	}).(ApiResourceBackendServiceResponseOutput)
}

func (o LookupIntegrationServiceEnvironmentManagedApiResultOutput) Capabilities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentManagedApiResult) []string { return v.Capabilities }).(pulumi.StringArrayOutput)
}

func (o LookupIntegrationServiceEnvironmentManagedApiResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentManagedApiResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o LookupIntegrationServiceEnvironmentManagedApiResultOutput) ConnectionParameters() pulumi.MapOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentManagedApiResult) map[string]interface{} {
		return v.ConnectionParameters
	}).(pulumi.MapOutput)
}

func (o LookupIntegrationServiceEnvironmentManagedApiResultOutput) DeploymentParameters() IntegrationServiceEnvironmentManagedApiDeploymentParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentManagedApiResult) *IntegrationServiceEnvironmentManagedApiDeploymentParametersResponse {
		return v.DeploymentParameters
	}).(IntegrationServiceEnvironmentManagedApiDeploymentParametersResponsePtrOutput)
}

func (o LookupIntegrationServiceEnvironmentManagedApiResultOutput) GeneralInformation() ApiResourceGeneralInformationResponseOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentManagedApiResult) ApiResourceGeneralInformationResponse {
		return v.GeneralInformation
	}).(ApiResourceGeneralInformationResponseOutput)
}

func (o LookupIntegrationServiceEnvironmentManagedApiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentManagedApiResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIntegrationServiceEnvironmentManagedApiResultOutput) IntegrationServiceEnvironment() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentManagedApiResult) *ResourceReferenceResponse {
		return v.IntegrationServiceEnvironment
	}).(ResourceReferenceResponsePtrOutput)
}

func (o LookupIntegrationServiceEnvironmentManagedApiResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentManagedApiResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationServiceEnvironmentManagedApiResultOutput) Metadata() ApiResourceMetadataResponseOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentManagedApiResult) ApiResourceMetadataResponse {
		return v.Metadata
	}).(ApiResourceMetadataResponseOutput)
}

func (o LookupIntegrationServiceEnvironmentManagedApiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentManagedApiResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIntegrationServiceEnvironmentManagedApiResultOutput) Policies() ApiResourcePoliciesResponseOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentManagedApiResult) ApiResourcePoliciesResponse {
		return v.Policies
	}).(ApiResourcePoliciesResponseOutput)
}

func (o LookupIntegrationServiceEnvironmentManagedApiResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentManagedApiResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupIntegrationServiceEnvironmentManagedApiResultOutput) RuntimeUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentManagedApiResult) []string { return v.RuntimeUrls }).(pulumi.StringArrayOutput)
}

func (o LookupIntegrationServiceEnvironmentManagedApiResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentManagedApiResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupIntegrationServiceEnvironmentManagedApiResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationServiceEnvironmentManagedApiResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationServiceEnvironmentManagedApiResultOutput{})
}

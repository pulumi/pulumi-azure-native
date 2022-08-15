


package v20220601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFhirService(ctx *pulumi.Context, args *LookupFhirServiceArgs, opts ...pulumi.InvokeOption) (*LookupFhirServiceResult, error) {
	var rv LookupFhirServiceResult
	err := ctx.Invoke("azure-native:healthcareapis/v20220601:getFhirService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFhirServiceArgs struct {
	FhirServiceName   string `pulumi:"fhirServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupFhirServiceResult struct {
	AccessPolicies                     []FhirServiceAccessPolicyEntryResponse          `pulumi:"accessPolicies"`
	AcrConfiguration                   *FhirServiceAcrConfigurationResponse            `pulumi:"acrConfiguration"`
	AuthenticationConfiguration        *FhirServiceAuthenticationConfigurationResponse `pulumi:"authenticationConfiguration"`
	CorsConfiguration                  *FhirServiceCorsConfigurationResponse           `pulumi:"corsConfiguration"`
	Etag                               *string                                         `pulumi:"etag"`
	EventState                         string                                          `pulumi:"eventState"`
	ExportConfiguration                *FhirServiceExportConfigurationResponse         `pulumi:"exportConfiguration"`
	Id                                 string                                          `pulumi:"id"`
	Identity                           *ServiceManagedIdentityResponseIdentity         `pulumi:"identity"`
	ImportConfiguration                *FhirServiceImportConfigurationResponse         `pulumi:"importConfiguration"`
	Kind                               *string                                         `pulumi:"kind"`
	Location                           *string                                         `pulumi:"location"`
	Name                               string                                          `pulumi:"name"`
	PrivateEndpointConnections         []PrivateEndpointConnectionResponse             `pulumi:"privateEndpointConnections"`
	ProvisioningState                  string                                          `pulumi:"provisioningState"`
	PublicNetworkAccess                string                                          `pulumi:"publicNetworkAccess"`
	ResourceVersionPolicyConfiguration *ResourceVersionPolicyConfigurationResponse     `pulumi:"resourceVersionPolicyConfiguration"`
	SystemData                         SystemDataResponse                              `pulumi:"systemData"`
	Tags                               map[string]string                               `pulumi:"tags"`
	Type                               string                                          `pulumi:"type"`
}

func LookupFhirServiceOutput(ctx *pulumi.Context, args LookupFhirServiceOutputArgs, opts ...pulumi.InvokeOption) LookupFhirServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFhirServiceResult, error) {
			args := v.(LookupFhirServiceArgs)
			r, err := LookupFhirService(ctx, &args, opts...)
			var s LookupFhirServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFhirServiceResultOutput)
}

type LookupFhirServiceOutputArgs struct {
	FhirServiceName   pulumi.StringInput `pulumi:"fhirServiceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupFhirServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFhirServiceArgs)(nil)).Elem()
}


type LookupFhirServiceResultOutput struct{ *pulumi.OutputState }

func (LookupFhirServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFhirServiceResult)(nil)).Elem()
}

func (o LookupFhirServiceResultOutput) ToLookupFhirServiceResultOutput() LookupFhirServiceResultOutput {
	return o
}

func (o LookupFhirServiceResultOutput) ToLookupFhirServiceResultOutputWithContext(ctx context.Context) LookupFhirServiceResultOutput {
	return o
}

func (o LookupFhirServiceResultOutput) AccessPolicies() FhirServiceAccessPolicyEntryResponseArrayOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) []FhirServiceAccessPolicyEntryResponse { return v.AccessPolicies }).(FhirServiceAccessPolicyEntryResponseArrayOutput)
}

func (o LookupFhirServiceResultOutput) AcrConfiguration() FhirServiceAcrConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *FhirServiceAcrConfigurationResponse { return v.AcrConfiguration }).(FhirServiceAcrConfigurationResponsePtrOutput)
}

func (o LookupFhirServiceResultOutput) AuthenticationConfiguration() FhirServiceAuthenticationConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *FhirServiceAuthenticationConfigurationResponse {
		return v.AuthenticationConfiguration
	}).(FhirServiceAuthenticationConfigurationResponsePtrOutput)
}

func (o LookupFhirServiceResultOutput) CorsConfiguration() FhirServiceCorsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *FhirServiceCorsConfigurationResponse { return v.CorsConfiguration }).(FhirServiceCorsConfigurationResponsePtrOutput)
}

func (o LookupFhirServiceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupFhirServiceResultOutput) EventState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) string { return v.EventState }).(pulumi.StringOutput)
}

func (o LookupFhirServiceResultOutput) ExportConfiguration() FhirServiceExportConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *FhirServiceExportConfigurationResponse { return v.ExportConfiguration }).(FhirServiceExportConfigurationResponsePtrOutput)
}

func (o LookupFhirServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFhirServiceResultOutput) Identity() ServiceManagedIdentityResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *ServiceManagedIdentityResponseIdentity { return v.Identity }).(ServiceManagedIdentityResponseIdentityPtrOutput)
}

func (o LookupFhirServiceResultOutput) ImportConfiguration() FhirServiceImportConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *FhirServiceImportConfigurationResponse { return v.ImportConfiguration }).(FhirServiceImportConfigurationResponsePtrOutput)
}

func (o LookupFhirServiceResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupFhirServiceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupFhirServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFhirServiceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupFhirServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupFhirServiceResultOutput) PublicNetworkAccess() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) string { return v.PublicNetworkAccess }).(pulumi.StringOutput)
}

func (o LookupFhirServiceResultOutput) ResourceVersionPolicyConfiguration() ResourceVersionPolicyConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *ResourceVersionPolicyConfigurationResponse {
		return v.ResourceVersionPolicyConfiguration
	}).(ResourceVersionPolicyConfigurationResponsePtrOutput)
}

func (o LookupFhirServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupFhirServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupFhirServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFhirServiceResultOutput{})
}

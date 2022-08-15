


package v20220601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDicomService(ctx *pulumi.Context, args *LookupDicomServiceArgs, opts ...pulumi.InvokeOption) (*LookupDicomServiceResult, error) {
	var rv LookupDicomServiceResult
	err := ctx.Invoke("azure-native:healthcareapis/v20220601:getDicomService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDicomServiceArgs struct {
	DicomServiceName  string `pulumi:"dicomServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupDicomServiceResult struct {
	AuthenticationConfiguration *DicomServiceAuthenticationConfigurationResponse `pulumi:"authenticationConfiguration"`
	CorsConfiguration           *CorsConfigurationResponse                       `pulumi:"corsConfiguration"`
	Etag                        *string                                          `pulumi:"etag"`
	Id                          string                                           `pulumi:"id"`
	Identity                    *ServiceManagedIdentityResponseIdentity          `pulumi:"identity"`
	Location                    *string                                          `pulumi:"location"`
	Name                        string                                           `pulumi:"name"`
	PrivateEndpointConnections  []PrivateEndpointConnectionResponse              `pulumi:"privateEndpointConnections"`
	ProvisioningState           string                                           `pulumi:"provisioningState"`
	PublicNetworkAccess         string                                           `pulumi:"publicNetworkAccess"`
	ServiceUrl                  string                                           `pulumi:"serviceUrl"`
	SystemData                  SystemDataResponse                               `pulumi:"systemData"`
	Tags                        map[string]string                                `pulumi:"tags"`
	Type                        string                                           `pulumi:"type"`
}

func LookupDicomServiceOutput(ctx *pulumi.Context, args LookupDicomServiceOutputArgs, opts ...pulumi.InvokeOption) LookupDicomServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDicomServiceResult, error) {
			args := v.(LookupDicomServiceArgs)
			r, err := LookupDicomService(ctx, &args, opts...)
			var s LookupDicomServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDicomServiceResultOutput)
}

type LookupDicomServiceOutputArgs struct {
	DicomServiceName  pulumi.StringInput `pulumi:"dicomServiceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupDicomServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDicomServiceArgs)(nil)).Elem()
}


type LookupDicomServiceResultOutput struct{ *pulumi.OutputState }

func (LookupDicomServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDicomServiceResult)(nil)).Elem()
}

func (o LookupDicomServiceResultOutput) ToLookupDicomServiceResultOutput() LookupDicomServiceResultOutput {
	return o
}

func (o LookupDicomServiceResultOutput) ToLookupDicomServiceResultOutputWithContext(ctx context.Context) LookupDicomServiceResultOutput {
	return o
}

func (o LookupDicomServiceResultOutput) AuthenticationConfiguration() DicomServiceAuthenticationConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) *DicomServiceAuthenticationConfigurationResponse {
		return v.AuthenticationConfiguration
	}).(DicomServiceAuthenticationConfigurationResponsePtrOutput)
}

func (o LookupDicomServiceResultOutput) CorsConfiguration() CorsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) *CorsConfigurationResponse { return v.CorsConfiguration }).(CorsConfigurationResponsePtrOutput)
}

func (o LookupDicomServiceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupDicomServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDicomServiceResultOutput) Identity() ServiceManagedIdentityResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) *ServiceManagedIdentityResponseIdentity { return v.Identity }).(ServiceManagedIdentityResponseIdentityPtrOutput)
}

func (o LookupDicomServiceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDicomServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDicomServiceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupDicomServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDicomServiceResultOutput) PublicNetworkAccess() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) string { return v.PublicNetworkAccess }).(pulumi.StringOutput)
}

func (o LookupDicomServiceResultOutput) ServiceUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) string { return v.ServiceUrl }).(pulumi.StringOutput)
}

func (o LookupDicomServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDicomServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDicomServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDicomServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDicomServiceResultOutput{})
}

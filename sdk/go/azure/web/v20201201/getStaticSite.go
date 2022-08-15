


package v20201201

import (
	"context"
	"reflect"

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

func LookupStaticSiteOutput(ctx *pulumi.Context, args LookupStaticSiteOutputArgs, opts ...pulumi.InvokeOption) LookupStaticSiteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStaticSiteResult, error) {
			args := v.(LookupStaticSiteArgs)
			r, err := LookupStaticSite(ctx, &args, opts...)
			var s LookupStaticSiteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStaticSiteResultOutput)
}

type LookupStaticSiteOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupStaticSiteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteArgs)(nil)).Elem()
}


type LookupStaticSiteResultOutput struct{ *pulumi.OutputState }

func (LookupStaticSiteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteResult)(nil)).Elem()
}

func (o LookupStaticSiteResultOutput) ToLookupStaticSiteResultOutput() LookupStaticSiteResultOutput {
	return o
}

func (o LookupStaticSiteResultOutput) ToLookupStaticSiteResultOutputWithContext(ctx context.Context) LookupStaticSiteResultOutput {
	return o
}

func (o LookupStaticSiteResultOutput) AllowConfigFileUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *bool { return v.AllowConfigFileUpdates }).(pulumi.BoolPtrOutput)
}

func (o LookupStaticSiteResultOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteResultOutput) BuildProperties() StaticSiteBuildPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *StaticSiteBuildPropertiesResponse { return v.BuildProperties }).(StaticSiteBuildPropertiesResponsePtrOutput)
}

func (o LookupStaticSiteResultOutput) ContentDistributionEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.ContentDistributionEndpoint }).(pulumi.StringOutput)
}

func (o LookupStaticSiteResultOutput) CustomDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) []string { return v.CustomDomains }).(pulumi.StringArrayOutput)
}

func (o LookupStaticSiteResultOutput) DefaultHostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.DefaultHostname }).(pulumi.StringOutput)
}

func (o LookupStaticSiteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStaticSiteResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupStaticSiteResultOutput) KeyVaultReferenceIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.KeyVaultReferenceIdentity }).(pulumi.StringOutput)
}

func (o LookupStaticSiteResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupStaticSiteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStaticSiteResultOutput) PrivateEndpointConnections() ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) []ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupStaticSiteResultOutput) Provider() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.Provider }).(pulumi.StringOutput)
}

func (o LookupStaticSiteResultOutput) RepositoryToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.RepositoryToken }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteResultOutput) RepositoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.RepositoryUrl }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteResultOutput) Sku() SkuDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *SkuDescriptionResponse { return v.Sku }).(SkuDescriptionResponsePtrOutput)
}

func (o LookupStaticSiteResultOutput) StagingEnvironmentPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.StagingEnvironmentPolicy }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupStaticSiteResultOutput) TemplateProperties() StaticSiteTemplateOptionsResponsePtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *StaticSiteTemplateOptionsResponse { return v.TemplateProperties }).(StaticSiteTemplateOptionsResponsePtrOutput)
}

func (o LookupStaticSiteResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupStaticSiteResultOutput) UserProvidedFunctionApps() StaticSiteUserProvidedFunctionAppResponseArrayOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) []StaticSiteUserProvidedFunctionAppResponse {
		return v.UserProvidedFunctionApps
	}).(StaticSiteUserProvidedFunctionAppResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStaticSiteResultOutput{})
}

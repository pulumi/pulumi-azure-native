


package v20150801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiEntity struct {
	ApiDefinitionUrl     *string                        `pulumi:"apiDefinitionUrl"`
	BackendService       *BackendServiceDefinition      `pulumi:"backendService"`
	Capabilities         []string                       `pulumi:"capabilities"`
	ChangedTime          *string                        `pulumi:"changedTime"`
	ConnectionParameters map[string]ConnectionParameter `pulumi:"connectionParameters"`
	CreatedTime          *string                        `pulumi:"createdTime"`
	GeneralInformation   *GeneralApiInformation         `pulumi:"generalInformation"`
	Id                   *string                        `pulumi:"id"`
	Kind                 *string                        `pulumi:"kind"`
	Location             string                         `pulumi:"location"`
	Metadata             interface{}                    `pulumi:"metadata"`
	Name                 *string                        `pulumi:"name"`
	Path                 *string                        `pulumi:"path"`
	Policies             *ApiPolicies                   `pulumi:"policies"`
	Protocols            []string                       `pulumi:"protocols"`
	RuntimeUrls          []string                       `pulumi:"runtimeUrls"`
	Tags                 map[string]string              `pulumi:"tags"`
	Type                 *string                        `pulumi:"type"`
}





type ApiEntityInput interface {
	pulumi.Input

	ToApiEntityOutput() ApiEntityOutput
	ToApiEntityOutputWithContext(context.Context) ApiEntityOutput
}

type ApiEntityArgs struct {
	ApiDefinitionUrl     pulumi.StringPtrInput            `pulumi:"apiDefinitionUrl"`
	BackendService       BackendServiceDefinitionPtrInput `pulumi:"backendService"`
	Capabilities         pulumi.StringArrayInput          `pulumi:"capabilities"`
	ChangedTime          pulumi.StringPtrInput            `pulumi:"changedTime"`
	ConnectionParameters ConnectionParameterMapInput      `pulumi:"connectionParameters"`
	CreatedTime          pulumi.StringPtrInput            `pulumi:"createdTime"`
	GeneralInformation   GeneralApiInformationPtrInput    `pulumi:"generalInformation"`
	Id                   pulumi.StringPtrInput            `pulumi:"id"`
	Kind                 pulumi.StringPtrInput            `pulumi:"kind"`
	Location             pulumi.StringInput               `pulumi:"location"`
	Metadata             pulumi.Input                     `pulumi:"metadata"`
	Name                 pulumi.StringPtrInput            `pulumi:"name"`
	Path                 pulumi.StringPtrInput            `pulumi:"path"`
	Policies             ApiPoliciesPtrInput              `pulumi:"policies"`
	Protocols            pulumi.StringArrayInput          `pulumi:"protocols"`
	RuntimeUrls          pulumi.StringArrayInput          `pulumi:"runtimeUrls"`
	Tags                 pulumi.StringMapInput            `pulumi:"tags"`
	Type                 pulumi.StringPtrInput            `pulumi:"type"`
}

func (ApiEntityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiEntity)(nil)).Elem()
}

func (i ApiEntityArgs) ToApiEntityOutput() ApiEntityOutput {
	return i.ToApiEntityOutputWithContext(context.Background())
}

func (i ApiEntityArgs) ToApiEntityOutputWithContext(ctx context.Context) ApiEntityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEntityOutput)
}

func (i ApiEntityArgs) ToApiEntityPtrOutput() ApiEntityPtrOutput {
	return i.ToApiEntityPtrOutputWithContext(context.Background())
}

func (i ApiEntityArgs) ToApiEntityPtrOutputWithContext(ctx context.Context) ApiEntityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEntityOutput).ToApiEntityPtrOutputWithContext(ctx)
}









type ApiEntityPtrInput interface {
	pulumi.Input

	ToApiEntityPtrOutput() ApiEntityPtrOutput
	ToApiEntityPtrOutputWithContext(context.Context) ApiEntityPtrOutput
}

type apiEntityPtrType ApiEntityArgs

func ApiEntityPtr(v *ApiEntityArgs) ApiEntityPtrInput {
	return (*apiEntityPtrType)(v)
}

func (*apiEntityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiEntity)(nil)).Elem()
}

func (i *apiEntityPtrType) ToApiEntityPtrOutput() ApiEntityPtrOutput {
	return i.ToApiEntityPtrOutputWithContext(context.Background())
}

func (i *apiEntityPtrType) ToApiEntityPtrOutputWithContext(ctx context.Context) ApiEntityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEntityPtrOutput)
}

type ApiEntityOutput struct{ *pulumi.OutputState }

func (ApiEntityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiEntity)(nil)).Elem()
}

func (o ApiEntityOutput) ToApiEntityOutput() ApiEntityOutput {
	return o
}

func (o ApiEntityOutput) ToApiEntityOutputWithContext(ctx context.Context) ApiEntityOutput {
	return o
}

func (o ApiEntityOutput) ToApiEntityPtrOutput() ApiEntityPtrOutput {
	return o.ToApiEntityPtrOutputWithContext(context.Background())
}

func (o ApiEntityOutput) ToApiEntityPtrOutputWithContext(ctx context.Context) ApiEntityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiEntity) *ApiEntity {
		return &v
	}).(ApiEntityPtrOutput)
}

func (o ApiEntityOutput) ApiDefinitionUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntity) *string { return v.ApiDefinitionUrl }).(pulumi.StringPtrOutput)
}

func (o ApiEntityOutput) BackendService() BackendServiceDefinitionPtrOutput {
	return o.ApplyT(func(v ApiEntity) *BackendServiceDefinition { return v.BackendService }).(BackendServiceDefinitionPtrOutput)
}

func (o ApiEntityOutput) Capabilities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApiEntity) []string { return v.Capabilities }).(pulumi.StringArrayOutput)
}

func (o ApiEntityOutput) ChangedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntity) *string { return v.ChangedTime }).(pulumi.StringPtrOutput)
}

func (o ApiEntityOutput) ConnectionParameters() ConnectionParameterMapOutput {
	return o.ApplyT(func(v ApiEntity) map[string]ConnectionParameter { return v.ConnectionParameters }).(ConnectionParameterMapOutput)
}

func (o ApiEntityOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntity) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

func (o ApiEntityOutput) GeneralInformation() GeneralApiInformationPtrOutput {
	return o.ApplyT(func(v ApiEntity) *GeneralApiInformation { return v.GeneralInformation }).(GeneralApiInformationPtrOutput)
}

func (o ApiEntityOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntity) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApiEntityOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntity) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ApiEntityOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ApiEntity) string { return v.Location }).(pulumi.StringOutput)
}

func (o ApiEntityOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v ApiEntity) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o ApiEntityOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntity) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApiEntityOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntity) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ApiEntityOutput) Policies() ApiPoliciesPtrOutput {
	return o.ApplyT(func(v ApiEntity) *ApiPolicies { return v.Policies }).(ApiPoliciesPtrOutput)
}

func (o ApiEntityOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApiEntity) []string { return v.Protocols }).(pulumi.StringArrayOutput)
}

func (o ApiEntityOutput) RuntimeUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApiEntity) []string { return v.RuntimeUrls }).(pulumi.StringArrayOutput)
}

func (o ApiEntityOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ApiEntity) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ApiEntityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ApiEntityPtrOutput struct{ *pulumi.OutputState }

func (ApiEntityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiEntity)(nil)).Elem()
}

func (o ApiEntityPtrOutput) ToApiEntityPtrOutput() ApiEntityPtrOutput {
	return o
}

func (o ApiEntityPtrOutput) ToApiEntityPtrOutputWithContext(ctx context.Context) ApiEntityPtrOutput {
	return o
}

func (o ApiEntityPtrOutput) Elem() ApiEntityOutput {
	return o.ApplyT(func(v *ApiEntity) ApiEntity {
		if v != nil {
			return *v
		}
		var ret ApiEntity
		return ret
	}).(ApiEntityOutput)
}

func (o ApiEntityPtrOutput) ApiDefinitionUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntity) *string {
		if v == nil {
			return nil
		}
		return v.ApiDefinitionUrl
	}).(pulumi.StringPtrOutput)
}

func (o ApiEntityPtrOutput) BackendService() BackendServiceDefinitionPtrOutput {
	return o.ApplyT(func(v *ApiEntity) *BackendServiceDefinition {
		if v == nil {
			return nil
		}
		return v.BackendService
	}).(BackendServiceDefinitionPtrOutput)
}

func (o ApiEntityPtrOutput) Capabilities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiEntity) []string {
		if v == nil {
			return nil
		}
		return v.Capabilities
	}).(pulumi.StringArrayOutput)
}

func (o ApiEntityPtrOutput) ChangedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntity) *string {
		if v == nil {
			return nil
		}
		return v.ChangedTime
	}).(pulumi.StringPtrOutput)
}

func (o ApiEntityPtrOutput) ConnectionParameters() ConnectionParameterMapOutput {
	return o.ApplyT(func(v *ApiEntity) map[string]ConnectionParameter {
		if v == nil {
			return nil
		}
		return v.ConnectionParameters
	}).(ConnectionParameterMapOutput)
}

func (o ApiEntityPtrOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntity) *string {
		if v == nil {
			return nil
		}
		return v.CreatedTime
	}).(pulumi.StringPtrOutput)
}

func (o ApiEntityPtrOutput) GeneralInformation() GeneralApiInformationPtrOutput {
	return o.ApplyT(func(v *ApiEntity) *GeneralApiInformation {
		if v == nil {
			return nil
		}
		return v.GeneralInformation
	}).(GeneralApiInformationPtrOutput)
}

func (o ApiEntityPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntity) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ApiEntityPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntity) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o ApiEntityPtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntity) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o ApiEntityPtrOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *ApiEntity) interface{} {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(pulumi.AnyOutput)
}

func (o ApiEntityPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntity) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApiEntityPtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntity) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o ApiEntityPtrOutput) Policies() ApiPoliciesPtrOutput {
	return o.ApplyT(func(v *ApiEntity) *ApiPolicies {
		if v == nil {
			return nil
		}
		return v.Policies
	}).(ApiPoliciesPtrOutput)
}

func (o ApiEntityPtrOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiEntity) []string {
		if v == nil {
			return nil
		}
		return v.Protocols
	}).(pulumi.StringArrayOutput)
}

func (o ApiEntityPtrOutput) RuntimeUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiEntity) []string {
		if v == nil {
			return nil
		}
		return v.RuntimeUrls
	}).(pulumi.StringArrayOutput)
}

func (o ApiEntityPtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApiEntity) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o ApiEntityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ApiEntityResponse struct {
	ApiDefinitionUrl     *string                                `pulumi:"apiDefinitionUrl"`
	BackendService       *BackendServiceDefinitionResponse      `pulumi:"backendService"`
	Capabilities         []string                               `pulumi:"capabilities"`
	ChangedTime          *string                                `pulumi:"changedTime"`
	ConnectionParameters map[string]ConnectionParameterResponse `pulumi:"connectionParameters"`
	CreatedTime          *string                                `pulumi:"createdTime"`
	GeneralInformation   *GeneralApiInformationResponse         `pulumi:"generalInformation"`
	Id                   *string                                `pulumi:"id"`
	Kind                 *string                                `pulumi:"kind"`
	Location             string                                 `pulumi:"location"`
	Metadata             interface{}                            `pulumi:"metadata"`
	Name                 *string                                `pulumi:"name"`
	Path                 *string                                `pulumi:"path"`
	Policies             *ApiPoliciesResponse                   `pulumi:"policies"`
	Protocols            []string                               `pulumi:"protocols"`
	RuntimeUrls          []string                               `pulumi:"runtimeUrls"`
	Tags                 map[string]string                      `pulumi:"tags"`
	Type                 *string                                `pulumi:"type"`
}





type ApiEntityResponseInput interface {
	pulumi.Input

	ToApiEntityResponseOutput() ApiEntityResponseOutput
	ToApiEntityResponseOutputWithContext(context.Context) ApiEntityResponseOutput
}

type ApiEntityResponseArgs struct {
	ApiDefinitionUrl     pulumi.StringPtrInput                    `pulumi:"apiDefinitionUrl"`
	BackendService       BackendServiceDefinitionResponsePtrInput `pulumi:"backendService"`
	Capabilities         pulumi.StringArrayInput                  `pulumi:"capabilities"`
	ChangedTime          pulumi.StringPtrInput                    `pulumi:"changedTime"`
	ConnectionParameters ConnectionParameterResponseMapInput      `pulumi:"connectionParameters"`
	CreatedTime          pulumi.StringPtrInput                    `pulumi:"createdTime"`
	GeneralInformation   GeneralApiInformationResponsePtrInput    `pulumi:"generalInformation"`
	Id                   pulumi.StringPtrInput                    `pulumi:"id"`
	Kind                 pulumi.StringPtrInput                    `pulumi:"kind"`
	Location             pulumi.StringInput                       `pulumi:"location"`
	Metadata             pulumi.Input                             `pulumi:"metadata"`
	Name                 pulumi.StringPtrInput                    `pulumi:"name"`
	Path                 pulumi.StringPtrInput                    `pulumi:"path"`
	Policies             ApiPoliciesResponsePtrInput              `pulumi:"policies"`
	Protocols            pulumi.StringArrayInput                  `pulumi:"protocols"`
	RuntimeUrls          pulumi.StringArrayInput                  `pulumi:"runtimeUrls"`
	Tags                 pulumi.StringMapInput                    `pulumi:"tags"`
	Type                 pulumi.StringPtrInput                    `pulumi:"type"`
}

func (ApiEntityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiEntityResponse)(nil)).Elem()
}

func (i ApiEntityResponseArgs) ToApiEntityResponseOutput() ApiEntityResponseOutput {
	return i.ToApiEntityResponseOutputWithContext(context.Background())
}

func (i ApiEntityResponseArgs) ToApiEntityResponseOutputWithContext(ctx context.Context) ApiEntityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEntityResponseOutput)
}

func (i ApiEntityResponseArgs) ToApiEntityResponsePtrOutput() ApiEntityResponsePtrOutput {
	return i.ToApiEntityResponsePtrOutputWithContext(context.Background())
}

func (i ApiEntityResponseArgs) ToApiEntityResponsePtrOutputWithContext(ctx context.Context) ApiEntityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEntityResponseOutput).ToApiEntityResponsePtrOutputWithContext(ctx)
}









type ApiEntityResponsePtrInput interface {
	pulumi.Input

	ToApiEntityResponsePtrOutput() ApiEntityResponsePtrOutput
	ToApiEntityResponsePtrOutputWithContext(context.Context) ApiEntityResponsePtrOutput
}

type apiEntityResponsePtrType ApiEntityResponseArgs

func ApiEntityResponsePtr(v *ApiEntityResponseArgs) ApiEntityResponsePtrInput {
	return (*apiEntityResponsePtrType)(v)
}

func (*apiEntityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiEntityResponse)(nil)).Elem()
}

func (i *apiEntityResponsePtrType) ToApiEntityResponsePtrOutput() ApiEntityResponsePtrOutput {
	return i.ToApiEntityResponsePtrOutputWithContext(context.Background())
}

func (i *apiEntityResponsePtrType) ToApiEntityResponsePtrOutputWithContext(ctx context.Context) ApiEntityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEntityResponsePtrOutput)
}

type ApiEntityResponseOutput struct{ *pulumi.OutputState }

func (ApiEntityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiEntityResponse)(nil)).Elem()
}

func (o ApiEntityResponseOutput) ToApiEntityResponseOutput() ApiEntityResponseOutput {
	return o
}

func (o ApiEntityResponseOutput) ToApiEntityResponseOutputWithContext(ctx context.Context) ApiEntityResponseOutput {
	return o
}

func (o ApiEntityResponseOutput) ToApiEntityResponsePtrOutput() ApiEntityResponsePtrOutput {
	return o.ToApiEntityResponsePtrOutputWithContext(context.Background())
}

func (o ApiEntityResponseOutput) ToApiEntityResponsePtrOutputWithContext(ctx context.Context) ApiEntityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiEntityResponse) *ApiEntityResponse {
		return &v
	}).(ApiEntityResponsePtrOutput)
}

func (o ApiEntityResponseOutput) ApiDefinitionUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntityResponse) *string { return v.ApiDefinitionUrl }).(pulumi.StringPtrOutput)
}

func (o ApiEntityResponseOutput) BackendService() BackendServiceDefinitionResponsePtrOutput {
	return o.ApplyT(func(v ApiEntityResponse) *BackendServiceDefinitionResponse { return v.BackendService }).(BackendServiceDefinitionResponsePtrOutput)
}

func (o ApiEntityResponseOutput) Capabilities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApiEntityResponse) []string { return v.Capabilities }).(pulumi.StringArrayOutput)
}

func (o ApiEntityResponseOutput) ChangedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntityResponse) *string { return v.ChangedTime }).(pulumi.StringPtrOutput)
}

func (o ApiEntityResponseOutput) ConnectionParameters() ConnectionParameterResponseMapOutput {
	return o.ApplyT(func(v ApiEntityResponse) map[string]ConnectionParameterResponse { return v.ConnectionParameters }).(ConnectionParameterResponseMapOutput)
}

func (o ApiEntityResponseOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntityResponse) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

func (o ApiEntityResponseOutput) GeneralInformation() GeneralApiInformationResponsePtrOutput {
	return o.ApplyT(func(v ApiEntityResponse) *GeneralApiInformationResponse { return v.GeneralInformation }).(GeneralApiInformationResponsePtrOutput)
}

func (o ApiEntityResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntityResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApiEntityResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntityResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ApiEntityResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ApiEntityResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o ApiEntityResponseOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v ApiEntityResponse) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o ApiEntityResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntityResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApiEntityResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntityResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ApiEntityResponseOutput) Policies() ApiPoliciesResponsePtrOutput {
	return o.ApplyT(func(v ApiEntityResponse) *ApiPoliciesResponse { return v.Policies }).(ApiPoliciesResponsePtrOutput)
}

func (o ApiEntityResponseOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApiEntityResponse) []string { return v.Protocols }).(pulumi.StringArrayOutput)
}

func (o ApiEntityResponseOutput) RuntimeUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApiEntityResponse) []string { return v.RuntimeUrls }).(pulumi.StringArrayOutput)
}

func (o ApiEntityResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ApiEntityResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ApiEntityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ApiEntityResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiEntityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiEntityResponse)(nil)).Elem()
}

func (o ApiEntityResponsePtrOutput) ToApiEntityResponsePtrOutput() ApiEntityResponsePtrOutput {
	return o
}

func (o ApiEntityResponsePtrOutput) ToApiEntityResponsePtrOutputWithContext(ctx context.Context) ApiEntityResponsePtrOutput {
	return o
}

func (o ApiEntityResponsePtrOutput) Elem() ApiEntityResponseOutput {
	return o.ApplyT(func(v *ApiEntityResponse) ApiEntityResponse {
		if v != nil {
			return *v
		}
		var ret ApiEntityResponse
		return ret
	}).(ApiEntityResponseOutput)
}

func (o ApiEntityResponsePtrOutput) ApiDefinitionUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApiDefinitionUrl
	}).(pulumi.StringPtrOutput)
}

func (o ApiEntityResponsePtrOutput) BackendService() BackendServiceDefinitionResponsePtrOutput {
	return o.ApplyT(func(v *ApiEntityResponse) *BackendServiceDefinitionResponse {
		if v == nil {
			return nil
		}
		return v.BackendService
	}).(BackendServiceDefinitionResponsePtrOutput)
}

func (o ApiEntityResponsePtrOutput) Capabilities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiEntityResponse) []string {
		if v == nil {
			return nil
		}
		return v.Capabilities
	}).(pulumi.StringArrayOutput)
}

func (o ApiEntityResponsePtrOutput) ChangedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.ChangedTime
	}).(pulumi.StringPtrOutput)
}

func (o ApiEntityResponsePtrOutput) ConnectionParameters() ConnectionParameterResponseMapOutput {
	return o.ApplyT(func(v *ApiEntityResponse) map[string]ConnectionParameterResponse {
		if v == nil {
			return nil
		}
		return v.ConnectionParameters
	}).(ConnectionParameterResponseMapOutput)
}

func (o ApiEntityResponsePtrOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedTime
	}).(pulumi.StringPtrOutput)
}

func (o ApiEntityResponsePtrOutput) GeneralInformation() GeneralApiInformationResponsePtrOutput {
	return o.ApplyT(func(v *ApiEntityResponse) *GeneralApiInformationResponse {
		if v == nil {
			return nil
		}
		return v.GeneralInformation
	}).(GeneralApiInformationResponsePtrOutput)
}

func (o ApiEntityResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ApiEntityResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o ApiEntityResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o ApiEntityResponsePtrOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *ApiEntityResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(pulumi.AnyOutput)
}

func (o ApiEntityResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApiEntityResponsePtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o ApiEntityResponsePtrOutput) Policies() ApiPoliciesResponsePtrOutput {
	return o.ApplyT(func(v *ApiEntityResponse) *ApiPoliciesResponse {
		if v == nil {
			return nil
		}
		return v.Policies
	}).(ApiPoliciesResponsePtrOutput)
}

func (o ApiEntityResponsePtrOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiEntityResponse) []string {
		if v == nil {
			return nil
		}
		return v.Protocols
	}).(pulumi.StringArrayOutput)
}

func (o ApiEntityResponsePtrOutput) RuntimeUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiEntityResponse) []string {
		if v == nil {
			return nil
		}
		return v.RuntimeUrls
	}).(pulumi.StringArrayOutput)
}

func (o ApiEntityResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApiEntityResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o ApiEntityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ApiOAuthSettings struct {
	ClientId         *string                              `pulumi:"clientId"`
	ClientSecret     *string                              `pulumi:"clientSecret"`
	CustomParameters map[string]ApiOAuthSettingsParameter `pulumi:"customParameters"`
	IdentityProvider *string                              `pulumi:"identityProvider"`
	Properties       interface{}                          `pulumi:"properties"`
	RedirectUrl      *string                              `pulumi:"redirectUrl"`
	Scopes           []string                             `pulumi:"scopes"`
}





type ApiOAuthSettingsInput interface {
	pulumi.Input

	ToApiOAuthSettingsOutput() ApiOAuthSettingsOutput
	ToApiOAuthSettingsOutputWithContext(context.Context) ApiOAuthSettingsOutput
}

type ApiOAuthSettingsArgs struct {
	ClientId         pulumi.StringPtrInput             `pulumi:"clientId"`
	ClientSecret     pulumi.StringPtrInput             `pulumi:"clientSecret"`
	CustomParameters ApiOAuthSettingsParameterMapInput `pulumi:"customParameters"`
	IdentityProvider pulumi.StringPtrInput             `pulumi:"identityProvider"`
	Properties       pulumi.Input                      `pulumi:"properties"`
	RedirectUrl      pulumi.StringPtrInput             `pulumi:"redirectUrl"`
	Scopes           pulumi.StringArrayInput           `pulumi:"scopes"`
}

func (ApiOAuthSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOAuthSettings)(nil)).Elem()
}

func (i ApiOAuthSettingsArgs) ToApiOAuthSettingsOutput() ApiOAuthSettingsOutput {
	return i.ToApiOAuthSettingsOutputWithContext(context.Background())
}

func (i ApiOAuthSettingsArgs) ToApiOAuthSettingsOutputWithContext(ctx context.Context) ApiOAuthSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOAuthSettingsOutput)
}

func (i ApiOAuthSettingsArgs) ToApiOAuthSettingsPtrOutput() ApiOAuthSettingsPtrOutput {
	return i.ToApiOAuthSettingsPtrOutputWithContext(context.Background())
}

func (i ApiOAuthSettingsArgs) ToApiOAuthSettingsPtrOutputWithContext(ctx context.Context) ApiOAuthSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOAuthSettingsOutput).ToApiOAuthSettingsPtrOutputWithContext(ctx)
}









type ApiOAuthSettingsPtrInput interface {
	pulumi.Input

	ToApiOAuthSettingsPtrOutput() ApiOAuthSettingsPtrOutput
	ToApiOAuthSettingsPtrOutputWithContext(context.Context) ApiOAuthSettingsPtrOutput
}

type apiOAuthSettingsPtrType ApiOAuthSettingsArgs

func ApiOAuthSettingsPtr(v *ApiOAuthSettingsArgs) ApiOAuthSettingsPtrInput {
	return (*apiOAuthSettingsPtrType)(v)
}

func (*apiOAuthSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiOAuthSettings)(nil)).Elem()
}

func (i *apiOAuthSettingsPtrType) ToApiOAuthSettingsPtrOutput() ApiOAuthSettingsPtrOutput {
	return i.ToApiOAuthSettingsPtrOutputWithContext(context.Background())
}

func (i *apiOAuthSettingsPtrType) ToApiOAuthSettingsPtrOutputWithContext(ctx context.Context) ApiOAuthSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOAuthSettingsPtrOutput)
}

type ApiOAuthSettingsOutput struct{ *pulumi.OutputState }

func (ApiOAuthSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOAuthSettings)(nil)).Elem()
}

func (o ApiOAuthSettingsOutput) ToApiOAuthSettingsOutput() ApiOAuthSettingsOutput {
	return o
}

func (o ApiOAuthSettingsOutput) ToApiOAuthSettingsOutputWithContext(ctx context.Context) ApiOAuthSettingsOutput {
	return o
}

func (o ApiOAuthSettingsOutput) ToApiOAuthSettingsPtrOutput() ApiOAuthSettingsPtrOutput {
	return o.ToApiOAuthSettingsPtrOutputWithContext(context.Background())
}

func (o ApiOAuthSettingsOutput) ToApiOAuthSettingsPtrOutputWithContext(ctx context.Context) ApiOAuthSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiOAuthSettings) *ApiOAuthSettings {
		return &v
	}).(ApiOAuthSettingsPtrOutput)
}

func (o ApiOAuthSettingsOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiOAuthSettings) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiOAuthSettings) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsOutput) CustomParameters() ApiOAuthSettingsParameterMapOutput {
	return o.ApplyT(func(v ApiOAuthSettings) map[string]ApiOAuthSettingsParameter { return v.CustomParameters }).(ApiOAuthSettingsParameterMapOutput)
}

func (o ApiOAuthSettingsOutput) IdentityProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiOAuthSettings) *string { return v.IdentityProvider }).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v ApiOAuthSettings) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o ApiOAuthSettingsOutput) RedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiOAuthSettings) *string { return v.RedirectUrl }).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApiOAuthSettings) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

type ApiOAuthSettingsPtrOutput struct{ *pulumi.OutputState }

func (ApiOAuthSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiOAuthSettings)(nil)).Elem()
}

func (o ApiOAuthSettingsPtrOutput) ToApiOAuthSettingsPtrOutput() ApiOAuthSettingsPtrOutput {
	return o
}

func (o ApiOAuthSettingsPtrOutput) ToApiOAuthSettingsPtrOutputWithContext(ctx context.Context) ApiOAuthSettingsPtrOutput {
	return o
}

func (o ApiOAuthSettingsPtrOutput) Elem() ApiOAuthSettingsOutput {
	return o.ApplyT(func(v *ApiOAuthSettings) ApiOAuthSettings {
		if v != nil {
			return *v
		}
		var ret ApiOAuthSettings
		return ret
	}).(ApiOAuthSettingsOutput)
}

func (o ApiOAuthSettingsPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiOAuthSettings) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsPtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiOAuthSettings) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsPtrOutput) CustomParameters() ApiOAuthSettingsParameterMapOutput {
	return o.ApplyT(func(v *ApiOAuthSettings) map[string]ApiOAuthSettingsParameter {
		if v == nil {
			return nil
		}
		return v.CustomParameters
	}).(ApiOAuthSettingsParameterMapOutput)
}

func (o ApiOAuthSettingsPtrOutput) IdentityProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiOAuthSettings) *string {
		if v == nil {
			return nil
		}
		return v.IdentityProvider
	}).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsPtrOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *ApiOAuthSettings) interface{} {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(pulumi.AnyOutput)
}

func (o ApiOAuthSettingsPtrOutput) RedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiOAuthSettings) *string {
		if v == nil {
			return nil
		}
		return v.RedirectUrl
	}).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsPtrOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiOAuthSettings) []string {
		if v == nil {
			return nil
		}
		return v.Scopes
	}).(pulumi.StringArrayOutput)
}

type ApiOAuthSettingsParameter struct {
	Options      interface{} `pulumi:"options"`
	UiDefinition interface{} `pulumi:"uiDefinition"`
	Value        *string     `pulumi:"value"`
}





type ApiOAuthSettingsParameterInput interface {
	pulumi.Input

	ToApiOAuthSettingsParameterOutput() ApiOAuthSettingsParameterOutput
	ToApiOAuthSettingsParameterOutputWithContext(context.Context) ApiOAuthSettingsParameterOutput
}

type ApiOAuthSettingsParameterArgs struct {
	Options      pulumi.Input          `pulumi:"options"`
	UiDefinition pulumi.Input          `pulumi:"uiDefinition"`
	Value        pulumi.StringPtrInput `pulumi:"value"`
}

func (ApiOAuthSettingsParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOAuthSettingsParameter)(nil)).Elem()
}

func (i ApiOAuthSettingsParameterArgs) ToApiOAuthSettingsParameterOutput() ApiOAuthSettingsParameterOutput {
	return i.ToApiOAuthSettingsParameterOutputWithContext(context.Background())
}

func (i ApiOAuthSettingsParameterArgs) ToApiOAuthSettingsParameterOutputWithContext(ctx context.Context) ApiOAuthSettingsParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOAuthSettingsParameterOutput)
}





type ApiOAuthSettingsParameterMapInput interface {
	pulumi.Input

	ToApiOAuthSettingsParameterMapOutput() ApiOAuthSettingsParameterMapOutput
	ToApiOAuthSettingsParameterMapOutputWithContext(context.Context) ApiOAuthSettingsParameterMapOutput
}

type ApiOAuthSettingsParameterMap map[string]ApiOAuthSettingsParameterInput

func (ApiOAuthSettingsParameterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiOAuthSettingsParameter)(nil)).Elem()
}

func (i ApiOAuthSettingsParameterMap) ToApiOAuthSettingsParameterMapOutput() ApiOAuthSettingsParameterMapOutput {
	return i.ToApiOAuthSettingsParameterMapOutputWithContext(context.Background())
}

func (i ApiOAuthSettingsParameterMap) ToApiOAuthSettingsParameterMapOutputWithContext(ctx context.Context) ApiOAuthSettingsParameterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOAuthSettingsParameterMapOutput)
}

type ApiOAuthSettingsParameterOutput struct{ *pulumi.OutputState }

func (ApiOAuthSettingsParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOAuthSettingsParameter)(nil)).Elem()
}

func (o ApiOAuthSettingsParameterOutput) ToApiOAuthSettingsParameterOutput() ApiOAuthSettingsParameterOutput {
	return o
}

func (o ApiOAuthSettingsParameterOutput) ToApiOAuthSettingsParameterOutputWithContext(ctx context.Context) ApiOAuthSettingsParameterOutput {
	return o
}

func (o ApiOAuthSettingsParameterOutput) Options() pulumi.AnyOutput {
	return o.ApplyT(func(v ApiOAuthSettingsParameter) interface{} { return v.Options }).(pulumi.AnyOutput)
}

func (o ApiOAuthSettingsParameterOutput) UiDefinition() pulumi.AnyOutput {
	return o.ApplyT(func(v ApiOAuthSettingsParameter) interface{} { return v.UiDefinition }).(pulumi.AnyOutput)
}

func (o ApiOAuthSettingsParameterOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiOAuthSettingsParameter) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ApiOAuthSettingsParameterMapOutput struct{ *pulumi.OutputState }

func (ApiOAuthSettingsParameterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiOAuthSettingsParameter)(nil)).Elem()
}

func (o ApiOAuthSettingsParameterMapOutput) ToApiOAuthSettingsParameterMapOutput() ApiOAuthSettingsParameterMapOutput {
	return o
}

func (o ApiOAuthSettingsParameterMapOutput) ToApiOAuthSettingsParameterMapOutputWithContext(ctx context.Context) ApiOAuthSettingsParameterMapOutput {
	return o
}

func (o ApiOAuthSettingsParameterMapOutput) MapIndex(k pulumi.StringInput) ApiOAuthSettingsParameterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApiOAuthSettingsParameter {
		return vs[0].(map[string]ApiOAuthSettingsParameter)[vs[1].(string)]
	}).(ApiOAuthSettingsParameterOutput)
}

type ApiOAuthSettingsParameterResponse struct {
	Options      interface{} `pulumi:"options"`
	UiDefinition interface{} `pulumi:"uiDefinition"`
	Value        *string     `pulumi:"value"`
}





type ApiOAuthSettingsParameterResponseInput interface {
	pulumi.Input

	ToApiOAuthSettingsParameterResponseOutput() ApiOAuthSettingsParameterResponseOutput
	ToApiOAuthSettingsParameterResponseOutputWithContext(context.Context) ApiOAuthSettingsParameterResponseOutput
}

type ApiOAuthSettingsParameterResponseArgs struct {
	Options      pulumi.Input          `pulumi:"options"`
	UiDefinition pulumi.Input          `pulumi:"uiDefinition"`
	Value        pulumi.StringPtrInput `pulumi:"value"`
}

func (ApiOAuthSettingsParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOAuthSettingsParameterResponse)(nil)).Elem()
}

func (i ApiOAuthSettingsParameterResponseArgs) ToApiOAuthSettingsParameterResponseOutput() ApiOAuthSettingsParameterResponseOutput {
	return i.ToApiOAuthSettingsParameterResponseOutputWithContext(context.Background())
}

func (i ApiOAuthSettingsParameterResponseArgs) ToApiOAuthSettingsParameterResponseOutputWithContext(ctx context.Context) ApiOAuthSettingsParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOAuthSettingsParameterResponseOutput)
}





type ApiOAuthSettingsParameterResponseMapInput interface {
	pulumi.Input

	ToApiOAuthSettingsParameterResponseMapOutput() ApiOAuthSettingsParameterResponseMapOutput
	ToApiOAuthSettingsParameterResponseMapOutputWithContext(context.Context) ApiOAuthSettingsParameterResponseMapOutput
}

type ApiOAuthSettingsParameterResponseMap map[string]ApiOAuthSettingsParameterResponseInput

func (ApiOAuthSettingsParameterResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiOAuthSettingsParameterResponse)(nil)).Elem()
}

func (i ApiOAuthSettingsParameterResponseMap) ToApiOAuthSettingsParameterResponseMapOutput() ApiOAuthSettingsParameterResponseMapOutput {
	return i.ToApiOAuthSettingsParameterResponseMapOutputWithContext(context.Background())
}

func (i ApiOAuthSettingsParameterResponseMap) ToApiOAuthSettingsParameterResponseMapOutputWithContext(ctx context.Context) ApiOAuthSettingsParameterResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOAuthSettingsParameterResponseMapOutput)
}

type ApiOAuthSettingsParameterResponseOutput struct{ *pulumi.OutputState }

func (ApiOAuthSettingsParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOAuthSettingsParameterResponse)(nil)).Elem()
}

func (o ApiOAuthSettingsParameterResponseOutput) ToApiOAuthSettingsParameterResponseOutput() ApiOAuthSettingsParameterResponseOutput {
	return o
}

func (o ApiOAuthSettingsParameterResponseOutput) ToApiOAuthSettingsParameterResponseOutputWithContext(ctx context.Context) ApiOAuthSettingsParameterResponseOutput {
	return o
}

func (o ApiOAuthSettingsParameterResponseOutput) Options() pulumi.AnyOutput {
	return o.ApplyT(func(v ApiOAuthSettingsParameterResponse) interface{} { return v.Options }).(pulumi.AnyOutput)
}

func (o ApiOAuthSettingsParameterResponseOutput) UiDefinition() pulumi.AnyOutput {
	return o.ApplyT(func(v ApiOAuthSettingsParameterResponse) interface{} { return v.UiDefinition }).(pulumi.AnyOutput)
}

func (o ApiOAuthSettingsParameterResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiOAuthSettingsParameterResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ApiOAuthSettingsParameterResponseMapOutput struct{ *pulumi.OutputState }

func (ApiOAuthSettingsParameterResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiOAuthSettingsParameterResponse)(nil)).Elem()
}

func (o ApiOAuthSettingsParameterResponseMapOutput) ToApiOAuthSettingsParameterResponseMapOutput() ApiOAuthSettingsParameterResponseMapOutput {
	return o
}

func (o ApiOAuthSettingsParameterResponseMapOutput) ToApiOAuthSettingsParameterResponseMapOutputWithContext(ctx context.Context) ApiOAuthSettingsParameterResponseMapOutput {
	return o
}

func (o ApiOAuthSettingsParameterResponseMapOutput) MapIndex(k pulumi.StringInput) ApiOAuthSettingsParameterResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApiOAuthSettingsParameterResponse {
		return vs[0].(map[string]ApiOAuthSettingsParameterResponse)[vs[1].(string)]
	}).(ApiOAuthSettingsParameterResponseOutput)
}

type ApiOAuthSettingsResponse struct {
	ClientId         *string                                      `pulumi:"clientId"`
	ClientSecret     *string                                      `pulumi:"clientSecret"`
	CustomParameters map[string]ApiOAuthSettingsParameterResponse `pulumi:"customParameters"`
	IdentityProvider *string                                      `pulumi:"identityProvider"`
	Properties       interface{}                                  `pulumi:"properties"`
	RedirectUrl      *string                                      `pulumi:"redirectUrl"`
	Scopes           []string                                     `pulumi:"scopes"`
}





type ApiOAuthSettingsResponseInput interface {
	pulumi.Input

	ToApiOAuthSettingsResponseOutput() ApiOAuthSettingsResponseOutput
	ToApiOAuthSettingsResponseOutputWithContext(context.Context) ApiOAuthSettingsResponseOutput
}

type ApiOAuthSettingsResponseArgs struct {
	ClientId         pulumi.StringPtrInput                     `pulumi:"clientId"`
	ClientSecret     pulumi.StringPtrInput                     `pulumi:"clientSecret"`
	CustomParameters ApiOAuthSettingsParameterResponseMapInput `pulumi:"customParameters"`
	IdentityProvider pulumi.StringPtrInput                     `pulumi:"identityProvider"`
	Properties       pulumi.Input                              `pulumi:"properties"`
	RedirectUrl      pulumi.StringPtrInput                     `pulumi:"redirectUrl"`
	Scopes           pulumi.StringArrayInput                   `pulumi:"scopes"`
}

func (ApiOAuthSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOAuthSettingsResponse)(nil)).Elem()
}

func (i ApiOAuthSettingsResponseArgs) ToApiOAuthSettingsResponseOutput() ApiOAuthSettingsResponseOutput {
	return i.ToApiOAuthSettingsResponseOutputWithContext(context.Background())
}

func (i ApiOAuthSettingsResponseArgs) ToApiOAuthSettingsResponseOutputWithContext(ctx context.Context) ApiOAuthSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOAuthSettingsResponseOutput)
}

func (i ApiOAuthSettingsResponseArgs) ToApiOAuthSettingsResponsePtrOutput() ApiOAuthSettingsResponsePtrOutput {
	return i.ToApiOAuthSettingsResponsePtrOutputWithContext(context.Background())
}

func (i ApiOAuthSettingsResponseArgs) ToApiOAuthSettingsResponsePtrOutputWithContext(ctx context.Context) ApiOAuthSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOAuthSettingsResponseOutput).ToApiOAuthSettingsResponsePtrOutputWithContext(ctx)
}









type ApiOAuthSettingsResponsePtrInput interface {
	pulumi.Input

	ToApiOAuthSettingsResponsePtrOutput() ApiOAuthSettingsResponsePtrOutput
	ToApiOAuthSettingsResponsePtrOutputWithContext(context.Context) ApiOAuthSettingsResponsePtrOutput
}

type apiOAuthSettingsResponsePtrType ApiOAuthSettingsResponseArgs

func ApiOAuthSettingsResponsePtr(v *ApiOAuthSettingsResponseArgs) ApiOAuthSettingsResponsePtrInput {
	return (*apiOAuthSettingsResponsePtrType)(v)
}

func (*apiOAuthSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiOAuthSettingsResponse)(nil)).Elem()
}

func (i *apiOAuthSettingsResponsePtrType) ToApiOAuthSettingsResponsePtrOutput() ApiOAuthSettingsResponsePtrOutput {
	return i.ToApiOAuthSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *apiOAuthSettingsResponsePtrType) ToApiOAuthSettingsResponsePtrOutputWithContext(ctx context.Context) ApiOAuthSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOAuthSettingsResponsePtrOutput)
}

type ApiOAuthSettingsResponseOutput struct{ *pulumi.OutputState }

func (ApiOAuthSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOAuthSettingsResponse)(nil)).Elem()
}

func (o ApiOAuthSettingsResponseOutput) ToApiOAuthSettingsResponseOutput() ApiOAuthSettingsResponseOutput {
	return o
}

func (o ApiOAuthSettingsResponseOutput) ToApiOAuthSettingsResponseOutputWithContext(ctx context.Context) ApiOAuthSettingsResponseOutput {
	return o
}

func (o ApiOAuthSettingsResponseOutput) ToApiOAuthSettingsResponsePtrOutput() ApiOAuthSettingsResponsePtrOutput {
	return o.ToApiOAuthSettingsResponsePtrOutputWithContext(context.Background())
}

func (o ApiOAuthSettingsResponseOutput) ToApiOAuthSettingsResponsePtrOutputWithContext(ctx context.Context) ApiOAuthSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiOAuthSettingsResponse) *ApiOAuthSettingsResponse {
		return &v
	}).(ApiOAuthSettingsResponsePtrOutput)
}

func (o ApiOAuthSettingsResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiOAuthSettingsResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsResponseOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiOAuthSettingsResponse) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsResponseOutput) CustomParameters() ApiOAuthSettingsParameterResponseMapOutput {
	return o.ApplyT(func(v ApiOAuthSettingsResponse) map[string]ApiOAuthSettingsParameterResponse {
		return v.CustomParameters
	}).(ApiOAuthSettingsParameterResponseMapOutput)
}

func (o ApiOAuthSettingsResponseOutput) IdentityProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiOAuthSettingsResponse) *string { return v.IdentityProvider }).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsResponseOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v ApiOAuthSettingsResponse) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o ApiOAuthSettingsResponseOutput) RedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiOAuthSettingsResponse) *string { return v.RedirectUrl }).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsResponseOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApiOAuthSettingsResponse) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

type ApiOAuthSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiOAuthSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiOAuthSettingsResponse)(nil)).Elem()
}

func (o ApiOAuthSettingsResponsePtrOutput) ToApiOAuthSettingsResponsePtrOutput() ApiOAuthSettingsResponsePtrOutput {
	return o
}

func (o ApiOAuthSettingsResponsePtrOutput) ToApiOAuthSettingsResponsePtrOutputWithContext(ctx context.Context) ApiOAuthSettingsResponsePtrOutput {
	return o
}

func (o ApiOAuthSettingsResponsePtrOutput) Elem() ApiOAuthSettingsResponseOutput {
	return o.ApplyT(func(v *ApiOAuthSettingsResponse) ApiOAuthSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ApiOAuthSettingsResponse
		return ret
	}).(ApiOAuthSettingsResponseOutput)
}

func (o ApiOAuthSettingsResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiOAuthSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsResponsePtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiOAuthSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsResponsePtrOutput) CustomParameters() ApiOAuthSettingsParameterResponseMapOutput {
	return o.ApplyT(func(v *ApiOAuthSettingsResponse) map[string]ApiOAuthSettingsParameterResponse {
		if v == nil {
			return nil
		}
		return v.CustomParameters
	}).(ApiOAuthSettingsParameterResponseMapOutput)
}

func (o ApiOAuthSettingsResponsePtrOutput) IdentityProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiOAuthSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.IdentityProvider
	}).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsResponsePtrOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *ApiOAuthSettingsResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(pulumi.AnyOutput)
}

func (o ApiOAuthSettingsResponsePtrOutput) RedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiOAuthSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.RedirectUrl
	}).(pulumi.StringPtrOutput)
}

func (o ApiOAuthSettingsResponsePtrOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiOAuthSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.Scopes
	}).(pulumi.StringArrayOutput)
}

type ApiPolicies struct {
	Content  *string           `pulumi:"content"`
	Id       *string           `pulumi:"id"`
	Kind     *string           `pulumi:"kind"`
	Location string            `pulumi:"location"`
	Name     *string           `pulumi:"name"`
	Tags     map[string]string `pulumi:"tags"`
	Type     *string           `pulumi:"type"`
}





type ApiPoliciesInput interface {
	pulumi.Input

	ToApiPoliciesOutput() ApiPoliciesOutput
	ToApiPoliciesOutputWithContext(context.Context) ApiPoliciesOutput
}

type ApiPoliciesArgs struct {
	Content  pulumi.StringPtrInput `pulumi:"content"`
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Kind     pulumi.StringPtrInput `pulumi:"kind"`
	Location pulumi.StringInput    `pulumi:"location"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Tags     pulumi.StringMapInput `pulumi:"tags"`
	Type     pulumi.StringPtrInput `pulumi:"type"`
}

func (ApiPoliciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiPolicies)(nil)).Elem()
}

func (i ApiPoliciesArgs) ToApiPoliciesOutput() ApiPoliciesOutput {
	return i.ToApiPoliciesOutputWithContext(context.Background())
}

func (i ApiPoliciesArgs) ToApiPoliciesOutputWithContext(ctx context.Context) ApiPoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPoliciesOutput)
}

func (i ApiPoliciesArgs) ToApiPoliciesPtrOutput() ApiPoliciesPtrOutput {
	return i.ToApiPoliciesPtrOutputWithContext(context.Background())
}

func (i ApiPoliciesArgs) ToApiPoliciesPtrOutputWithContext(ctx context.Context) ApiPoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPoliciesOutput).ToApiPoliciesPtrOutputWithContext(ctx)
}









type ApiPoliciesPtrInput interface {
	pulumi.Input

	ToApiPoliciesPtrOutput() ApiPoliciesPtrOutput
	ToApiPoliciesPtrOutputWithContext(context.Context) ApiPoliciesPtrOutput
}

type apiPoliciesPtrType ApiPoliciesArgs

func ApiPoliciesPtr(v *ApiPoliciesArgs) ApiPoliciesPtrInput {
	return (*apiPoliciesPtrType)(v)
}

func (*apiPoliciesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPolicies)(nil)).Elem()
}

func (i *apiPoliciesPtrType) ToApiPoliciesPtrOutput() ApiPoliciesPtrOutput {
	return i.ToApiPoliciesPtrOutputWithContext(context.Background())
}

func (i *apiPoliciesPtrType) ToApiPoliciesPtrOutputWithContext(ctx context.Context) ApiPoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPoliciesPtrOutput)
}

type ApiPoliciesOutput struct{ *pulumi.OutputState }

func (ApiPoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiPolicies)(nil)).Elem()
}

func (o ApiPoliciesOutput) ToApiPoliciesOutput() ApiPoliciesOutput {
	return o
}

func (o ApiPoliciesOutput) ToApiPoliciesOutputWithContext(ctx context.Context) ApiPoliciesOutput {
	return o
}

func (o ApiPoliciesOutput) ToApiPoliciesPtrOutput() ApiPoliciesPtrOutput {
	return o.ToApiPoliciesPtrOutputWithContext(context.Background())
}

func (o ApiPoliciesOutput) ToApiPoliciesPtrOutputWithContext(ctx context.Context) ApiPoliciesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiPolicies) *ApiPolicies {
		return &v
	}).(ApiPoliciesPtrOutput)
}

func (o ApiPoliciesOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPolicies) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o ApiPoliciesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPolicies) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApiPoliciesOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPolicies) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ApiPoliciesOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ApiPolicies) string { return v.Location }).(pulumi.StringOutput)
}

func (o ApiPoliciesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPolicies) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApiPoliciesOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ApiPolicies) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ApiPoliciesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPolicies) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ApiPoliciesPtrOutput struct{ *pulumi.OutputState }

func (ApiPoliciesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPolicies)(nil)).Elem()
}

func (o ApiPoliciesPtrOutput) ToApiPoliciesPtrOutput() ApiPoliciesPtrOutput {
	return o
}

func (o ApiPoliciesPtrOutput) ToApiPoliciesPtrOutputWithContext(ctx context.Context) ApiPoliciesPtrOutput {
	return o
}

func (o ApiPoliciesPtrOutput) Elem() ApiPoliciesOutput {
	return o.ApplyT(func(v *ApiPolicies) ApiPolicies {
		if v != nil {
			return *v
		}
		var ret ApiPolicies
		return ret
	}).(ApiPoliciesOutput)
}

func (o ApiPoliciesPtrOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPolicies) *string {
		if v == nil {
			return nil
		}
		return v.Content
	}).(pulumi.StringPtrOutput)
}

func (o ApiPoliciesPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPolicies) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ApiPoliciesPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPolicies) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o ApiPoliciesPtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPolicies) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o ApiPoliciesPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPolicies) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApiPoliciesPtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApiPolicies) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o ApiPoliciesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPolicies) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ApiPoliciesResponse struct {
	Content  *string           `pulumi:"content"`
	Id       *string           `pulumi:"id"`
	Kind     *string           `pulumi:"kind"`
	Location string            `pulumi:"location"`
	Name     *string           `pulumi:"name"`
	Tags     map[string]string `pulumi:"tags"`
	Type     *string           `pulumi:"type"`
}





type ApiPoliciesResponseInput interface {
	pulumi.Input

	ToApiPoliciesResponseOutput() ApiPoliciesResponseOutput
	ToApiPoliciesResponseOutputWithContext(context.Context) ApiPoliciesResponseOutput
}

type ApiPoliciesResponseArgs struct {
	Content  pulumi.StringPtrInput `pulumi:"content"`
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Kind     pulumi.StringPtrInput `pulumi:"kind"`
	Location pulumi.StringInput    `pulumi:"location"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Tags     pulumi.StringMapInput `pulumi:"tags"`
	Type     pulumi.StringPtrInput `pulumi:"type"`
}

func (ApiPoliciesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiPoliciesResponse)(nil)).Elem()
}

func (i ApiPoliciesResponseArgs) ToApiPoliciesResponseOutput() ApiPoliciesResponseOutput {
	return i.ToApiPoliciesResponseOutputWithContext(context.Background())
}

func (i ApiPoliciesResponseArgs) ToApiPoliciesResponseOutputWithContext(ctx context.Context) ApiPoliciesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPoliciesResponseOutput)
}

func (i ApiPoliciesResponseArgs) ToApiPoliciesResponsePtrOutput() ApiPoliciesResponsePtrOutput {
	return i.ToApiPoliciesResponsePtrOutputWithContext(context.Background())
}

func (i ApiPoliciesResponseArgs) ToApiPoliciesResponsePtrOutputWithContext(ctx context.Context) ApiPoliciesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPoliciesResponseOutput).ToApiPoliciesResponsePtrOutputWithContext(ctx)
}









type ApiPoliciesResponsePtrInput interface {
	pulumi.Input

	ToApiPoliciesResponsePtrOutput() ApiPoliciesResponsePtrOutput
	ToApiPoliciesResponsePtrOutputWithContext(context.Context) ApiPoliciesResponsePtrOutput
}

type apiPoliciesResponsePtrType ApiPoliciesResponseArgs

func ApiPoliciesResponsePtr(v *ApiPoliciesResponseArgs) ApiPoliciesResponsePtrInput {
	return (*apiPoliciesResponsePtrType)(v)
}

func (*apiPoliciesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPoliciesResponse)(nil)).Elem()
}

func (i *apiPoliciesResponsePtrType) ToApiPoliciesResponsePtrOutput() ApiPoliciesResponsePtrOutput {
	return i.ToApiPoliciesResponsePtrOutputWithContext(context.Background())
}

func (i *apiPoliciesResponsePtrType) ToApiPoliciesResponsePtrOutputWithContext(ctx context.Context) ApiPoliciesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPoliciesResponsePtrOutput)
}

type ApiPoliciesResponseOutput struct{ *pulumi.OutputState }

func (ApiPoliciesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiPoliciesResponse)(nil)).Elem()
}

func (o ApiPoliciesResponseOutput) ToApiPoliciesResponseOutput() ApiPoliciesResponseOutput {
	return o
}

func (o ApiPoliciesResponseOutput) ToApiPoliciesResponseOutputWithContext(ctx context.Context) ApiPoliciesResponseOutput {
	return o
}

func (o ApiPoliciesResponseOutput) ToApiPoliciesResponsePtrOutput() ApiPoliciesResponsePtrOutput {
	return o.ToApiPoliciesResponsePtrOutputWithContext(context.Background())
}

func (o ApiPoliciesResponseOutput) ToApiPoliciesResponsePtrOutputWithContext(ctx context.Context) ApiPoliciesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiPoliciesResponse) *ApiPoliciesResponse {
		return &v
	}).(ApiPoliciesResponsePtrOutput)
}

func (o ApiPoliciesResponseOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPoliciesResponse) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o ApiPoliciesResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPoliciesResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApiPoliciesResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPoliciesResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ApiPoliciesResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ApiPoliciesResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o ApiPoliciesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPoliciesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApiPoliciesResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ApiPoliciesResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ApiPoliciesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPoliciesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ApiPoliciesResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiPoliciesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPoliciesResponse)(nil)).Elem()
}

func (o ApiPoliciesResponsePtrOutput) ToApiPoliciesResponsePtrOutput() ApiPoliciesResponsePtrOutput {
	return o
}

func (o ApiPoliciesResponsePtrOutput) ToApiPoliciesResponsePtrOutputWithContext(ctx context.Context) ApiPoliciesResponsePtrOutput {
	return o
}

func (o ApiPoliciesResponsePtrOutput) Elem() ApiPoliciesResponseOutput {
	return o.ApplyT(func(v *ApiPoliciesResponse) ApiPoliciesResponse {
		if v != nil {
			return *v
		}
		var ret ApiPoliciesResponse
		return ret
	}).(ApiPoliciesResponseOutput)
}

func (o ApiPoliciesResponsePtrOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPoliciesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Content
	}).(pulumi.StringPtrOutput)
}

func (o ApiPoliciesResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPoliciesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ApiPoliciesResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPoliciesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o ApiPoliciesResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPoliciesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o ApiPoliciesResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPoliciesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApiPoliciesResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApiPoliciesResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o ApiPoliciesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPoliciesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ArmPlan struct {
	Name          *string `pulumi:"name"`
	Product       *string `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     *string `pulumi:"publisher"`
	Version       *string `pulumi:"version"`
}





type ArmPlanInput interface {
	pulumi.Input

	ToArmPlanOutput() ArmPlanOutput
	ToArmPlanOutputWithContext(context.Context) ArmPlanOutput
}

type ArmPlanArgs struct {
	Name          pulumi.StringPtrInput `pulumi:"name"`
	Product       pulumi.StringPtrInput `pulumi:"product"`
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	Publisher     pulumi.StringPtrInput `pulumi:"publisher"`
	Version       pulumi.StringPtrInput `pulumi:"version"`
}

func (ArmPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmPlan)(nil)).Elem()
}

func (i ArmPlanArgs) ToArmPlanOutput() ArmPlanOutput {
	return i.ToArmPlanOutputWithContext(context.Background())
}

func (i ArmPlanArgs) ToArmPlanOutputWithContext(ctx context.Context) ArmPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmPlanOutput)
}

func (i ArmPlanArgs) ToArmPlanPtrOutput() ArmPlanPtrOutput {
	return i.ToArmPlanPtrOutputWithContext(context.Background())
}

func (i ArmPlanArgs) ToArmPlanPtrOutputWithContext(ctx context.Context) ArmPlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmPlanOutput).ToArmPlanPtrOutputWithContext(ctx)
}









type ArmPlanPtrInput interface {
	pulumi.Input

	ToArmPlanPtrOutput() ArmPlanPtrOutput
	ToArmPlanPtrOutputWithContext(context.Context) ArmPlanPtrOutput
}

type armPlanPtrType ArmPlanArgs

func ArmPlanPtr(v *ArmPlanArgs) ArmPlanPtrInput {
	return (*armPlanPtrType)(v)
}

func (*armPlanPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmPlan)(nil)).Elem()
}

func (i *armPlanPtrType) ToArmPlanPtrOutput() ArmPlanPtrOutput {
	return i.ToArmPlanPtrOutputWithContext(context.Background())
}

func (i *armPlanPtrType) ToArmPlanPtrOutputWithContext(ctx context.Context) ArmPlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmPlanPtrOutput)
}

type ArmPlanOutput struct{ *pulumi.OutputState }

func (ArmPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmPlan)(nil)).Elem()
}

func (o ArmPlanOutput) ToArmPlanOutput() ArmPlanOutput {
	return o
}

func (o ArmPlanOutput) ToArmPlanOutputWithContext(ctx context.Context) ArmPlanOutput {
	return o
}

func (o ArmPlanOutput) ToArmPlanPtrOutput() ArmPlanPtrOutput {
	return o.ToArmPlanPtrOutputWithContext(context.Background())
}

func (o ArmPlanOutput) ToArmPlanPtrOutputWithContext(ctx context.Context) ArmPlanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ArmPlan) *ArmPlan {
		return &v
	}).(ArmPlanPtrOutput)
}

func (o ArmPlanOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmPlan) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ArmPlanOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmPlan) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o ArmPlanOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmPlan) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o ArmPlanOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmPlan) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ArmPlanOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmPlan) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ArmPlanPtrOutput struct{ *pulumi.OutputState }

func (ArmPlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmPlan)(nil)).Elem()
}

func (o ArmPlanPtrOutput) ToArmPlanPtrOutput() ArmPlanPtrOutput {
	return o
}

func (o ArmPlanPtrOutput) ToArmPlanPtrOutputWithContext(ctx context.Context) ArmPlanPtrOutput {
	return o
}

func (o ArmPlanPtrOutput) Elem() ArmPlanOutput {
	return o.ApplyT(func(v *ArmPlan) ArmPlan {
		if v != nil {
			return *v
		}
		var ret ArmPlan
		return ret
	}).(ArmPlanOutput)
}

func (o ArmPlanPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmPlan) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ArmPlanPtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmPlan) *string {
		if v == nil {
			return nil
		}
		return v.Product
	}).(pulumi.StringPtrOutput)
}

func (o ArmPlanPtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmPlan) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o ArmPlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmPlan) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o ArmPlanPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmPlan) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type ArmPlanResponse struct {
	Name          *string `pulumi:"name"`
	Product       *string `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     *string `pulumi:"publisher"`
	Version       *string `pulumi:"version"`
}





type ArmPlanResponseInput interface {
	pulumi.Input

	ToArmPlanResponseOutput() ArmPlanResponseOutput
	ToArmPlanResponseOutputWithContext(context.Context) ArmPlanResponseOutput
}

type ArmPlanResponseArgs struct {
	Name          pulumi.StringPtrInput `pulumi:"name"`
	Product       pulumi.StringPtrInput `pulumi:"product"`
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	Publisher     pulumi.StringPtrInput `pulumi:"publisher"`
	Version       pulumi.StringPtrInput `pulumi:"version"`
}

func (ArmPlanResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmPlanResponse)(nil)).Elem()
}

func (i ArmPlanResponseArgs) ToArmPlanResponseOutput() ArmPlanResponseOutput {
	return i.ToArmPlanResponseOutputWithContext(context.Background())
}

func (i ArmPlanResponseArgs) ToArmPlanResponseOutputWithContext(ctx context.Context) ArmPlanResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmPlanResponseOutput)
}

func (i ArmPlanResponseArgs) ToArmPlanResponsePtrOutput() ArmPlanResponsePtrOutput {
	return i.ToArmPlanResponsePtrOutputWithContext(context.Background())
}

func (i ArmPlanResponseArgs) ToArmPlanResponsePtrOutputWithContext(ctx context.Context) ArmPlanResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmPlanResponseOutput).ToArmPlanResponsePtrOutputWithContext(ctx)
}









type ArmPlanResponsePtrInput interface {
	pulumi.Input

	ToArmPlanResponsePtrOutput() ArmPlanResponsePtrOutput
	ToArmPlanResponsePtrOutputWithContext(context.Context) ArmPlanResponsePtrOutput
}

type armPlanResponsePtrType ArmPlanResponseArgs

func ArmPlanResponsePtr(v *ArmPlanResponseArgs) ArmPlanResponsePtrInput {
	return (*armPlanResponsePtrType)(v)
}

func (*armPlanResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmPlanResponse)(nil)).Elem()
}

func (i *armPlanResponsePtrType) ToArmPlanResponsePtrOutput() ArmPlanResponsePtrOutput {
	return i.ToArmPlanResponsePtrOutputWithContext(context.Background())
}

func (i *armPlanResponsePtrType) ToArmPlanResponsePtrOutputWithContext(ctx context.Context) ArmPlanResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmPlanResponsePtrOutput)
}

type ArmPlanResponseOutput struct{ *pulumi.OutputState }

func (ArmPlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmPlanResponse)(nil)).Elem()
}

func (o ArmPlanResponseOutput) ToArmPlanResponseOutput() ArmPlanResponseOutput {
	return o
}

func (o ArmPlanResponseOutput) ToArmPlanResponseOutputWithContext(ctx context.Context) ArmPlanResponseOutput {
	return o
}

func (o ArmPlanResponseOutput) ToArmPlanResponsePtrOutput() ArmPlanResponsePtrOutput {
	return o.ToArmPlanResponsePtrOutputWithContext(context.Background())
}

func (o ArmPlanResponseOutput) ToArmPlanResponsePtrOutputWithContext(ctx context.Context) ArmPlanResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ArmPlanResponse) *ArmPlanResponse {
		return &v
	}).(ArmPlanResponsePtrOutput)
}

func (o ArmPlanResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmPlanResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponseOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmPlanResponse) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponseOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmPlanResponse) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmPlanResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmPlanResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ArmPlanResponsePtrOutput struct{ *pulumi.OutputState }

func (ArmPlanResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmPlanResponse)(nil)).Elem()
}

func (o ArmPlanResponsePtrOutput) ToArmPlanResponsePtrOutput() ArmPlanResponsePtrOutput {
	return o
}

func (o ArmPlanResponsePtrOutput) ToArmPlanResponsePtrOutputWithContext(ctx context.Context) ArmPlanResponsePtrOutput {
	return o
}

func (o ArmPlanResponsePtrOutput) Elem() ArmPlanResponseOutput {
	return o.ApplyT(func(v *ArmPlanResponse) ArmPlanResponse {
		if v != nil {
			return *v
		}
		var ret ArmPlanResponse
		return ret
	}).(ArmPlanResponseOutput)
}

func (o ArmPlanResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponsePtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Product
	}).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponsePtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type BackendServiceDefinition struct {
	HostingEnvironmentServiceUrls []HostingEnvironmentServiceDescriptions `pulumi:"hostingEnvironmentServiceUrls"`
	Id                            *string                                 `pulumi:"id"`
	Kind                          *string                                 `pulumi:"kind"`
	Location                      string                                  `pulumi:"location"`
	Name                          *string                                 `pulumi:"name"`
	ServiceUrl                    *string                                 `pulumi:"serviceUrl"`
	Tags                          map[string]string                       `pulumi:"tags"`
	Type                          *string                                 `pulumi:"type"`
}





type BackendServiceDefinitionInput interface {
	pulumi.Input

	ToBackendServiceDefinitionOutput() BackendServiceDefinitionOutput
	ToBackendServiceDefinitionOutputWithContext(context.Context) BackendServiceDefinitionOutput
}

type BackendServiceDefinitionArgs struct {
	HostingEnvironmentServiceUrls HostingEnvironmentServiceDescriptionsArrayInput `pulumi:"hostingEnvironmentServiceUrls"`
	Id                            pulumi.StringPtrInput                           `pulumi:"id"`
	Kind                          pulumi.StringPtrInput                           `pulumi:"kind"`
	Location                      pulumi.StringInput                              `pulumi:"location"`
	Name                          pulumi.StringPtrInput                           `pulumi:"name"`
	ServiceUrl                    pulumi.StringPtrInput                           `pulumi:"serviceUrl"`
	Tags                          pulumi.StringMapInput                           `pulumi:"tags"`
	Type                          pulumi.StringPtrInput                           `pulumi:"type"`
}

func (BackendServiceDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendServiceDefinition)(nil)).Elem()
}

func (i BackendServiceDefinitionArgs) ToBackendServiceDefinitionOutput() BackendServiceDefinitionOutput {
	return i.ToBackendServiceDefinitionOutputWithContext(context.Background())
}

func (i BackendServiceDefinitionArgs) ToBackendServiceDefinitionOutputWithContext(ctx context.Context) BackendServiceDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceDefinitionOutput)
}

func (i BackendServiceDefinitionArgs) ToBackendServiceDefinitionPtrOutput() BackendServiceDefinitionPtrOutput {
	return i.ToBackendServiceDefinitionPtrOutputWithContext(context.Background())
}

func (i BackendServiceDefinitionArgs) ToBackendServiceDefinitionPtrOutputWithContext(ctx context.Context) BackendServiceDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceDefinitionOutput).ToBackendServiceDefinitionPtrOutputWithContext(ctx)
}









type BackendServiceDefinitionPtrInput interface {
	pulumi.Input

	ToBackendServiceDefinitionPtrOutput() BackendServiceDefinitionPtrOutput
	ToBackendServiceDefinitionPtrOutputWithContext(context.Context) BackendServiceDefinitionPtrOutput
}

type backendServiceDefinitionPtrType BackendServiceDefinitionArgs

func BackendServiceDefinitionPtr(v *BackendServiceDefinitionArgs) BackendServiceDefinitionPtrInput {
	return (*backendServiceDefinitionPtrType)(v)
}

func (*backendServiceDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServiceDefinition)(nil)).Elem()
}

func (i *backendServiceDefinitionPtrType) ToBackendServiceDefinitionPtrOutput() BackendServiceDefinitionPtrOutput {
	return i.ToBackendServiceDefinitionPtrOutputWithContext(context.Background())
}

func (i *backendServiceDefinitionPtrType) ToBackendServiceDefinitionPtrOutputWithContext(ctx context.Context) BackendServiceDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceDefinitionPtrOutput)
}

type BackendServiceDefinitionOutput struct{ *pulumi.OutputState }

func (BackendServiceDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendServiceDefinition)(nil)).Elem()
}

func (o BackendServiceDefinitionOutput) ToBackendServiceDefinitionOutput() BackendServiceDefinitionOutput {
	return o
}

func (o BackendServiceDefinitionOutput) ToBackendServiceDefinitionOutputWithContext(ctx context.Context) BackendServiceDefinitionOutput {
	return o
}

func (o BackendServiceDefinitionOutput) ToBackendServiceDefinitionPtrOutput() BackendServiceDefinitionPtrOutput {
	return o.ToBackendServiceDefinitionPtrOutputWithContext(context.Background())
}

func (o BackendServiceDefinitionOutput) ToBackendServiceDefinitionPtrOutputWithContext(ctx context.Context) BackendServiceDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackendServiceDefinition) *BackendServiceDefinition {
		return &v
	}).(BackendServiceDefinitionPtrOutput)
}

func (o BackendServiceDefinitionOutput) HostingEnvironmentServiceUrls() HostingEnvironmentServiceDescriptionsArrayOutput {
	return o.ApplyT(func(v BackendServiceDefinition) []HostingEnvironmentServiceDescriptions {
		return v.HostingEnvironmentServiceUrls
	}).(HostingEnvironmentServiceDescriptionsArrayOutput)
}

func (o BackendServiceDefinitionOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendServiceDefinition) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o BackendServiceDefinitionOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendServiceDefinition) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o BackendServiceDefinitionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v BackendServiceDefinition) string { return v.Location }).(pulumi.StringOutput)
}

func (o BackendServiceDefinitionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendServiceDefinition) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o BackendServiceDefinitionOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendServiceDefinition) *string { return v.ServiceUrl }).(pulumi.StringPtrOutput)
}

func (o BackendServiceDefinitionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v BackendServiceDefinition) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o BackendServiceDefinitionOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendServiceDefinition) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type BackendServiceDefinitionPtrOutput struct{ *pulumi.OutputState }

func (BackendServiceDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServiceDefinition)(nil)).Elem()
}

func (o BackendServiceDefinitionPtrOutput) ToBackendServiceDefinitionPtrOutput() BackendServiceDefinitionPtrOutput {
	return o
}

func (o BackendServiceDefinitionPtrOutput) ToBackendServiceDefinitionPtrOutputWithContext(ctx context.Context) BackendServiceDefinitionPtrOutput {
	return o
}

func (o BackendServiceDefinitionPtrOutput) Elem() BackendServiceDefinitionOutput {
	return o.ApplyT(func(v *BackendServiceDefinition) BackendServiceDefinition {
		if v != nil {
			return *v
		}
		var ret BackendServiceDefinition
		return ret
	}).(BackendServiceDefinitionOutput)
}

func (o BackendServiceDefinitionPtrOutput) HostingEnvironmentServiceUrls() HostingEnvironmentServiceDescriptionsArrayOutput {
	return o.ApplyT(func(v *BackendServiceDefinition) []HostingEnvironmentServiceDescriptions {
		if v == nil {
			return nil
		}
		return v.HostingEnvironmentServiceUrls
	}).(HostingEnvironmentServiceDescriptionsArrayOutput)
}

func (o BackendServiceDefinitionPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendServiceDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o BackendServiceDefinitionPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendServiceDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o BackendServiceDefinitionPtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendServiceDefinition) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o BackendServiceDefinitionPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendServiceDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o BackendServiceDefinitionPtrOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendServiceDefinition) *string {
		if v == nil {
			return nil
		}
		return v.ServiceUrl
	}).(pulumi.StringPtrOutput)
}

func (o BackendServiceDefinitionPtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BackendServiceDefinition) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o BackendServiceDefinitionPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendServiceDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type BackendServiceDefinitionResponse struct {
	HostingEnvironmentServiceUrls []HostingEnvironmentServiceDescriptionsResponse `pulumi:"hostingEnvironmentServiceUrls"`
	Id                            *string                                         `pulumi:"id"`
	Kind                          *string                                         `pulumi:"kind"`
	Location                      string                                          `pulumi:"location"`
	Name                          *string                                         `pulumi:"name"`
	ServiceUrl                    *string                                         `pulumi:"serviceUrl"`
	Tags                          map[string]string                               `pulumi:"tags"`
	Type                          *string                                         `pulumi:"type"`
}





type BackendServiceDefinitionResponseInput interface {
	pulumi.Input

	ToBackendServiceDefinitionResponseOutput() BackendServiceDefinitionResponseOutput
	ToBackendServiceDefinitionResponseOutputWithContext(context.Context) BackendServiceDefinitionResponseOutput
}

type BackendServiceDefinitionResponseArgs struct {
	HostingEnvironmentServiceUrls HostingEnvironmentServiceDescriptionsResponseArrayInput `pulumi:"hostingEnvironmentServiceUrls"`
	Id                            pulumi.StringPtrInput                                   `pulumi:"id"`
	Kind                          pulumi.StringPtrInput                                   `pulumi:"kind"`
	Location                      pulumi.StringInput                                      `pulumi:"location"`
	Name                          pulumi.StringPtrInput                                   `pulumi:"name"`
	ServiceUrl                    pulumi.StringPtrInput                                   `pulumi:"serviceUrl"`
	Tags                          pulumi.StringMapInput                                   `pulumi:"tags"`
	Type                          pulumi.StringPtrInput                                   `pulumi:"type"`
}

func (BackendServiceDefinitionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendServiceDefinitionResponse)(nil)).Elem()
}

func (i BackendServiceDefinitionResponseArgs) ToBackendServiceDefinitionResponseOutput() BackendServiceDefinitionResponseOutput {
	return i.ToBackendServiceDefinitionResponseOutputWithContext(context.Background())
}

func (i BackendServiceDefinitionResponseArgs) ToBackendServiceDefinitionResponseOutputWithContext(ctx context.Context) BackendServiceDefinitionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceDefinitionResponseOutput)
}

func (i BackendServiceDefinitionResponseArgs) ToBackendServiceDefinitionResponsePtrOutput() BackendServiceDefinitionResponsePtrOutput {
	return i.ToBackendServiceDefinitionResponsePtrOutputWithContext(context.Background())
}

func (i BackendServiceDefinitionResponseArgs) ToBackendServiceDefinitionResponsePtrOutputWithContext(ctx context.Context) BackendServiceDefinitionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceDefinitionResponseOutput).ToBackendServiceDefinitionResponsePtrOutputWithContext(ctx)
}









type BackendServiceDefinitionResponsePtrInput interface {
	pulumi.Input

	ToBackendServiceDefinitionResponsePtrOutput() BackendServiceDefinitionResponsePtrOutput
	ToBackendServiceDefinitionResponsePtrOutputWithContext(context.Context) BackendServiceDefinitionResponsePtrOutput
}

type backendServiceDefinitionResponsePtrType BackendServiceDefinitionResponseArgs

func BackendServiceDefinitionResponsePtr(v *BackendServiceDefinitionResponseArgs) BackendServiceDefinitionResponsePtrInput {
	return (*backendServiceDefinitionResponsePtrType)(v)
}

func (*backendServiceDefinitionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServiceDefinitionResponse)(nil)).Elem()
}

func (i *backendServiceDefinitionResponsePtrType) ToBackendServiceDefinitionResponsePtrOutput() BackendServiceDefinitionResponsePtrOutput {
	return i.ToBackendServiceDefinitionResponsePtrOutputWithContext(context.Background())
}

func (i *backendServiceDefinitionResponsePtrType) ToBackendServiceDefinitionResponsePtrOutputWithContext(ctx context.Context) BackendServiceDefinitionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceDefinitionResponsePtrOutput)
}

type BackendServiceDefinitionResponseOutput struct{ *pulumi.OutputState }

func (BackendServiceDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendServiceDefinitionResponse)(nil)).Elem()
}

func (o BackendServiceDefinitionResponseOutput) ToBackendServiceDefinitionResponseOutput() BackendServiceDefinitionResponseOutput {
	return o
}

func (o BackendServiceDefinitionResponseOutput) ToBackendServiceDefinitionResponseOutputWithContext(ctx context.Context) BackendServiceDefinitionResponseOutput {
	return o
}

func (o BackendServiceDefinitionResponseOutput) ToBackendServiceDefinitionResponsePtrOutput() BackendServiceDefinitionResponsePtrOutput {
	return o.ToBackendServiceDefinitionResponsePtrOutputWithContext(context.Background())
}

func (o BackendServiceDefinitionResponseOutput) ToBackendServiceDefinitionResponsePtrOutputWithContext(ctx context.Context) BackendServiceDefinitionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackendServiceDefinitionResponse) *BackendServiceDefinitionResponse {
		return &v
	}).(BackendServiceDefinitionResponsePtrOutput)
}

func (o BackendServiceDefinitionResponseOutput) HostingEnvironmentServiceUrls() HostingEnvironmentServiceDescriptionsResponseArrayOutput {
	return o.ApplyT(func(v BackendServiceDefinitionResponse) []HostingEnvironmentServiceDescriptionsResponse {
		return v.HostingEnvironmentServiceUrls
	}).(HostingEnvironmentServiceDescriptionsResponseArrayOutput)
}

func (o BackendServiceDefinitionResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendServiceDefinitionResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o BackendServiceDefinitionResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendServiceDefinitionResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o BackendServiceDefinitionResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v BackendServiceDefinitionResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o BackendServiceDefinitionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendServiceDefinitionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o BackendServiceDefinitionResponseOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendServiceDefinitionResponse) *string { return v.ServiceUrl }).(pulumi.StringPtrOutput)
}

func (o BackendServiceDefinitionResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v BackendServiceDefinitionResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o BackendServiceDefinitionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendServiceDefinitionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type BackendServiceDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (BackendServiceDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServiceDefinitionResponse)(nil)).Elem()
}

func (o BackendServiceDefinitionResponsePtrOutput) ToBackendServiceDefinitionResponsePtrOutput() BackendServiceDefinitionResponsePtrOutput {
	return o
}

func (o BackendServiceDefinitionResponsePtrOutput) ToBackendServiceDefinitionResponsePtrOutputWithContext(ctx context.Context) BackendServiceDefinitionResponsePtrOutput {
	return o
}

func (o BackendServiceDefinitionResponsePtrOutput) Elem() BackendServiceDefinitionResponseOutput {
	return o.ApplyT(func(v *BackendServiceDefinitionResponse) BackendServiceDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret BackendServiceDefinitionResponse
		return ret
	}).(BackendServiceDefinitionResponseOutput)
}

func (o BackendServiceDefinitionResponsePtrOutput) HostingEnvironmentServiceUrls() HostingEnvironmentServiceDescriptionsResponseArrayOutput {
	return o.ApplyT(func(v *BackendServiceDefinitionResponse) []HostingEnvironmentServiceDescriptionsResponse {
		if v == nil {
			return nil
		}
		return v.HostingEnvironmentServiceUrls
	}).(HostingEnvironmentServiceDescriptionsResponseArrayOutput)
}

func (o BackendServiceDefinitionResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendServiceDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o BackendServiceDefinitionResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendServiceDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o BackendServiceDefinitionResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendServiceDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o BackendServiceDefinitionResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendServiceDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o BackendServiceDefinitionResponsePtrOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendServiceDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceUrl
	}).(pulumi.StringPtrOutput)
}

func (o BackendServiceDefinitionResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BackendServiceDefinitionResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o BackendServiceDefinitionResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendServiceDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ConnectionError struct {
	Code     *string           `pulumi:"code"`
	Id       *string           `pulumi:"id"`
	Kind     *string           `pulumi:"kind"`
	Location string            `pulumi:"location"`
	Message  *string           `pulumi:"message"`
	Name     *string           `pulumi:"name"`
	Tags     map[string]string `pulumi:"tags"`
	Type     *string           `pulumi:"type"`
}





type ConnectionErrorInput interface {
	pulumi.Input

	ToConnectionErrorOutput() ConnectionErrorOutput
	ToConnectionErrorOutputWithContext(context.Context) ConnectionErrorOutput
}

type ConnectionErrorArgs struct {
	Code     pulumi.StringPtrInput `pulumi:"code"`
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Kind     pulumi.StringPtrInput `pulumi:"kind"`
	Location pulumi.StringInput    `pulumi:"location"`
	Message  pulumi.StringPtrInput `pulumi:"message"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Tags     pulumi.StringMapInput `pulumi:"tags"`
	Type     pulumi.StringPtrInput `pulumi:"type"`
}

func (ConnectionErrorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionError)(nil)).Elem()
}

func (i ConnectionErrorArgs) ToConnectionErrorOutput() ConnectionErrorOutput {
	return i.ToConnectionErrorOutputWithContext(context.Background())
}

func (i ConnectionErrorArgs) ToConnectionErrorOutputWithContext(ctx context.Context) ConnectionErrorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionErrorOutput)
}

func (i ConnectionErrorArgs) ToConnectionErrorPtrOutput() ConnectionErrorPtrOutput {
	return i.ToConnectionErrorPtrOutputWithContext(context.Background())
}

func (i ConnectionErrorArgs) ToConnectionErrorPtrOutputWithContext(ctx context.Context) ConnectionErrorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionErrorOutput).ToConnectionErrorPtrOutputWithContext(ctx)
}









type ConnectionErrorPtrInput interface {
	pulumi.Input

	ToConnectionErrorPtrOutput() ConnectionErrorPtrOutput
	ToConnectionErrorPtrOutputWithContext(context.Context) ConnectionErrorPtrOutput
}

type connectionErrorPtrType ConnectionErrorArgs

func ConnectionErrorPtr(v *ConnectionErrorArgs) ConnectionErrorPtrInput {
	return (*connectionErrorPtrType)(v)
}

func (*connectionErrorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionError)(nil)).Elem()
}

func (i *connectionErrorPtrType) ToConnectionErrorPtrOutput() ConnectionErrorPtrOutput {
	return i.ToConnectionErrorPtrOutputWithContext(context.Background())
}

func (i *connectionErrorPtrType) ToConnectionErrorPtrOutputWithContext(ctx context.Context) ConnectionErrorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionErrorPtrOutput)
}

type ConnectionErrorOutput struct{ *pulumi.OutputState }

func (ConnectionErrorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionError)(nil)).Elem()
}

func (o ConnectionErrorOutput) ToConnectionErrorOutput() ConnectionErrorOutput {
	return o
}

func (o ConnectionErrorOutput) ToConnectionErrorOutputWithContext(ctx context.Context) ConnectionErrorOutput {
	return o
}

func (o ConnectionErrorOutput) ToConnectionErrorPtrOutput() ConnectionErrorPtrOutput {
	return o.ToConnectionErrorPtrOutputWithContext(context.Background())
}

func (o ConnectionErrorOutput) ToConnectionErrorPtrOutputWithContext(ctx context.Context) ConnectionErrorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionError) *ConnectionError {
		return &v
	}).(ConnectionErrorPtrOutput)
}

func (o ConnectionErrorOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionError) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionError) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionError) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionError) string { return v.Location }).(pulumi.StringOutput)
}

func (o ConnectionErrorOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionError) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionError) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ConnectionError) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ConnectionErrorOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionError) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ConnectionErrorPtrOutput struct{ *pulumi.OutputState }

func (ConnectionErrorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionError)(nil)).Elem()
}

func (o ConnectionErrorPtrOutput) ToConnectionErrorPtrOutput() ConnectionErrorPtrOutput {
	return o
}

func (o ConnectionErrorPtrOutput) ToConnectionErrorPtrOutputWithContext(ctx context.Context) ConnectionErrorPtrOutput {
	return o
}

func (o ConnectionErrorPtrOutput) Elem() ConnectionErrorOutput {
	return o.ApplyT(func(v *ConnectionError) ConnectionError {
		if v != nil {
			return *v
		}
		var ret ConnectionError
		return ret
	}).(ConnectionErrorOutput)
}

func (o ConnectionErrorPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionError) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionError) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionError) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorPtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionError) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionError) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionError) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorPtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConnectionError) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o ConnectionErrorPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionError) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ConnectionErrorResponse struct {
	Code     *string           `pulumi:"code"`
	Id       *string           `pulumi:"id"`
	Kind     *string           `pulumi:"kind"`
	Location string            `pulumi:"location"`
	Message  *string           `pulumi:"message"`
	Name     *string           `pulumi:"name"`
	Tags     map[string]string `pulumi:"tags"`
	Type     *string           `pulumi:"type"`
}





type ConnectionErrorResponseInput interface {
	pulumi.Input

	ToConnectionErrorResponseOutput() ConnectionErrorResponseOutput
	ToConnectionErrorResponseOutputWithContext(context.Context) ConnectionErrorResponseOutput
}

type ConnectionErrorResponseArgs struct {
	Code     pulumi.StringPtrInput `pulumi:"code"`
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Kind     pulumi.StringPtrInput `pulumi:"kind"`
	Location pulumi.StringInput    `pulumi:"location"`
	Message  pulumi.StringPtrInput `pulumi:"message"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Tags     pulumi.StringMapInput `pulumi:"tags"`
	Type     pulumi.StringPtrInput `pulumi:"type"`
}

func (ConnectionErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionErrorResponse)(nil)).Elem()
}

func (i ConnectionErrorResponseArgs) ToConnectionErrorResponseOutput() ConnectionErrorResponseOutput {
	return i.ToConnectionErrorResponseOutputWithContext(context.Background())
}

func (i ConnectionErrorResponseArgs) ToConnectionErrorResponseOutputWithContext(ctx context.Context) ConnectionErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionErrorResponseOutput)
}

func (i ConnectionErrorResponseArgs) ToConnectionErrorResponsePtrOutput() ConnectionErrorResponsePtrOutput {
	return i.ToConnectionErrorResponsePtrOutputWithContext(context.Background())
}

func (i ConnectionErrorResponseArgs) ToConnectionErrorResponsePtrOutputWithContext(ctx context.Context) ConnectionErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionErrorResponseOutput).ToConnectionErrorResponsePtrOutputWithContext(ctx)
}









type ConnectionErrorResponsePtrInput interface {
	pulumi.Input

	ToConnectionErrorResponsePtrOutput() ConnectionErrorResponsePtrOutput
	ToConnectionErrorResponsePtrOutputWithContext(context.Context) ConnectionErrorResponsePtrOutput
}

type connectionErrorResponsePtrType ConnectionErrorResponseArgs

func ConnectionErrorResponsePtr(v *ConnectionErrorResponseArgs) ConnectionErrorResponsePtrInput {
	return (*connectionErrorResponsePtrType)(v)
}

func (*connectionErrorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionErrorResponse)(nil)).Elem()
}

func (i *connectionErrorResponsePtrType) ToConnectionErrorResponsePtrOutput() ConnectionErrorResponsePtrOutput {
	return i.ToConnectionErrorResponsePtrOutputWithContext(context.Background())
}

func (i *connectionErrorResponsePtrType) ToConnectionErrorResponsePtrOutputWithContext(ctx context.Context) ConnectionErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionErrorResponsePtrOutput)
}

type ConnectionErrorResponseOutput struct{ *pulumi.OutputState }

func (ConnectionErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionErrorResponse)(nil)).Elem()
}

func (o ConnectionErrorResponseOutput) ToConnectionErrorResponseOutput() ConnectionErrorResponseOutput {
	return o
}

func (o ConnectionErrorResponseOutput) ToConnectionErrorResponseOutputWithContext(ctx context.Context) ConnectionErrorResponseOutput {
	return o
}

func (o ConnectionErrorResponseOutput) ToConnectionErrorResponsePtrOutput() ConnectionErrorResponsePtrOutput {
	return o.ToConnectionErrorResponsePtrOutputWithContext(context.Background())
}

func (o ConnectionErrorResponseOutput) ToConnectionErrorResponsePtrOutputWithContext(ctx context.Context) ConnectionErrorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionErrorResponse) *ConnectionErrorResponse {
		return &v
	}).(ConnectionErrorResponsePtrOutput)
}

func (o ConnectionErrorResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionErrorResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionErrorResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionErrorResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionErrorResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o ConnectionErrorResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionErrorResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionErrorResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ConnectionErrorResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ConnectionErrorResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionErrorResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ConnectionErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectionErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionErrorResponse)(nil)).Elem()
}

func (o ConnectionErrorResponsePtrOutput) ToConnectionErrorResponsePtrOutput() ConnectionErrorResponsePtrOutput {
	return o
}

func (o ConnectionErrorResponsePtrOutput) ToConnectionErrorResponsePtrOutputWithContext(ctx context.Context) ConnectionErrorResponsePtrOutput {
	return o
}

func (o ConnectionErrorResponsePtrOutput) Elem() ConnectionErrorResponseOutput {
	return o.ApplyT(func(v *ConnectionErrorResponse) ConnectionErrorResponse {
		if v != nil {
			return *v
		}
		var ret ConnectionErrorResponse
		return ret
	}).(ConnectionErrorResponseOutput)
}

func (o ConnectionErrorResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionErrorResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionErrorResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConnectionErrorResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o ConnectionErrorResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ConnectionParameter struct {
	DefaultValue  interface{}              `pulumi:"defaultValue"`
	OAuthSettings *ApiOAuthSettings        `pulumi:"oAuthSettings"`
	Type          *ConnectionParameterType `pulumi:"type"`
	UiDefinition  interface{}              `pulumi:"uiDefinition"`
}





type ConnectionParameterInput interface {
	pulumi.Input

	ToConnectionParameterOutput() ConnectionParameterOutput
	ToConnectionParameterOutputWithContext(context.Context) ConnectionParameterOutput
}

type ConnectionParameterArgs struct {
	DefaultValue  pulumi.Input                    `pulumi:"defaultValue"`
	OAuthSettings ApiOAuthSettingsPtrInput        `pulumi:"oAuthSettings"`
	Type          ConnectionParameterTypePtrInput `pulumi:"type"`
	UiDefinition  pulumi.Input                    `pulumi:"uiDefinition"`
}

func (ConnectionParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionParameter)(nil)).Elem()
}

func (i ConnectionParameterArgs) ToConnectionParameterOutput() ConnectionParameterOutput {
	return i.ToConnectionParameterOutputWithContext(context.Background())
}

func (i ConnectionParameterArgs) ToConnectionParameterOutputWithContext(ctx context.Context) ConnectionParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionParameterOutput)
}





type ConnectionParameterMapInput interface {
	pulumi.Input

	ToConnectionParameterMapOutput() ConnectionParameterMapOutput
	ToConnectionParameterMapOutputWithContext(context.Context) ConnectionParameterMapOutput
}

type ConnectionParameterMap map[string]ConnectionParameterInput

func (ConnectionParameterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConnectionParameter)(nil)).Elem()
}

func (i ConnectionParameterMap) ToConnectionParameterMapOutput() ConnectionParameterMapOutput {
	return i.ToConnectionParameterMapOutputWithContext(context.Background())
}

func (i ConnectionParameterMap) ToConnectionParameterMapOutputWithContext(ctx context.Context) ConnectionParameterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionParameterMapOutput)
}

type ConnectionParameterOutput struct{ *pulumi.OutputState }

func (ConnectionParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionParameter)(nil)).Elem()
}

func (o ConnectionParameterOutput) ToConnectionParameterOutput() ConnectionParameterOutput {
	return o
}

func (o ConnectionParameterOutput) ToConnectionParameterOutputWithContext(ctx context.Context) ConnectionParameterOutput {
	return o
}

func (o ConnectionParameterOutput) DefaultValue() pulumi.AnyOutput {
	return o.ApplyT(func(v ConnectionParameter) interface{} { return v.DefaultValue }).(pulumi.AnyOutput)
}

func (o ConnectionParameterOutput) OAuthSettings() ApiOAuthSettingsPtrOutput {
	return o.ApplyT(func(v ConnectionParameter) *ApiOAuthSettings { return v.OAuthSettings }).(ApiOAuthSettingsPtrOutput)
}

func (o ConnectionParameterOutput) Type() ConnectionParameterTypePtrOutput {
	return o.ApplyT(func(v ConnectionParameter) *ConnectionParameterType { return v.Type }).(ConnectionParameterTypePtrOutput)
}

func (o ConnectionParameterOutput) UiDefinition() pulumi.AnyOutput {
	return o.ApplyT(func(v ConnectionParameter) interface{} { return v.UiDefinition }).(pulumi.AnyOutput)
}

type ConnectionParameterMapOutput struct{ *pulumi.OutputState }

func (ConnectionParameterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConnectionParameter)(nil)).Elem()
}

func (o ConnectionParameterMapOutput) ToConnectionParameterMapOutput() ConnectionParameterMapOutput {
	return o
}

func (o ConnectionParameterMapOutput) ToConnectionParameterMapOutputWithContext(ctx context.Context) ConnectionParameterMapOutput {
	return o
}

func (o ConnectionParameterMapOutput) MapIndex(k pulumi.StringInput) ConnectionParameterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ConnectionParameter {
		return vs[0].(map[string]ConnectionParameter)[vs[1].(string)]
	}).(ConnectionParameterOutput)
}

type ConnectionParameterResponse struct {
	DefaultValue  interface{}               `pulumi:"defaultValue"`
	OAuthSettings *ApiOAuthSettingsResponse `pulumi:"oAuthSettings"`
	Type          *string                   `pulumi:"type"`
	UiDefinition  interface{}               `pulumi:"uiDefinition"`
}





type ConnectionParameterResponseInput interface {
	pulumi.Input

	ToConnectionParameterResponseOutput() ConnectionParameterResponseOutput
	ToConnectionParameterResponseOutputWithContext(context.Context) ConnectionParameterResponseOutput
}

type ConnectionParameterResponseArgs struct {
	DefaultValue  pulumi.Input                     `pulumi:"defaultValue"`
	OAuthSettings ApiOAuthSettingsResponsePtrInput `pulumi:"oAuthSettings"`
	Type          pulumi.StringPtrInput            `pulumi:"type"`
	UiDefinition  pulumi.Input                     `pulumi:"uiDefinition"`
}

func (ConnectionParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionParameterResponse)(nil)).Elem()
}

func (i ConnectionParameterResponseArgs) ToConnectionParameterResponseOutput() ConnectionParameterResponseOutput {
	return i.ToConnectionParameterResponseOutputWithContext(context.Background())
}

func (i ConnectionParameterResponseArgs) ToConnectionParameterResponseOutputWithContext(ctx context.Context) ConnectionParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionParameterResponseOutput)
}





type ConnectionParameterResponseMapInput interface {
	pulumi.Input

	ToConnectionParameterResponseMapOutput() ConnectionParameterResponseMapOutput
	ToConnectionParameterResponseMapOutputWithContext(context.Context) ConnectionParameterResponseMapOutput
}

type ConnectionParameterResponseMap map[string]ConnectionParameterResponseInput

func (ConnectionParameterResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConnectionParameterResponse)(nil)).Elem()
}

func (i ConnectionParameterResponseMap) ToConnectionParameterResponseMapOutput() ConnectionParameterResponseMapOutput {
	return i.ToConnectionParameterResponseMapOutputWithContext(context.Background())
}

func (i ConnectionParameterResponseMap) ToConnectionParameterResponseMapOutputWithContext(ctx context.Context) ConnectionParameterResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionParameterResponseMapOutput)
}

type ConnectionParameterResponseOutput struct{ *pulumi.OutputState }

func (ConnectionParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionParameterResponse)(nil)).Elem()
}

func (o ConnectionParameterResponseOutput) ToConnectionParameterResponseOutput() ConnectionParameterResponseOutput {
	return o
}

func (o ConnectionParameterResponseOutput) ToConnectionParameterResponseOutputWithContext(ctx context.Context) ConnectionParameterResponseOutput {
	return o
}

func (o ConnectionParameterResponseOutput) DefaultValue() pulumi.AnyOutput {
	return o.ApplyT(func(v ConnectionParameterResponse) interface{} { return v.DefaultValue }).(pulumi.AnyOutput)
}

func (o ConnectionParameterResponseOutput) OAuthSettings() ApiOAuthSettingsResponsePtrOutput {
	return o.ApplyT(func(v ConnectionParameterResponse) *ApiOAuthSettingsResponse { return v.OAuthSettings }).(ApiOAuthSettingsResponsePtrOutput)
}

func (o ConnectionParameterResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionParameterResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ConnectionParameterResponseOutput) UiDefinition() pulumi.AnyOutput {
	return o.ApplyT(func(v ConnectionParameterResponse) interface{} { return v.UiDefinition }).(pulumi.AnyOutput)
}

type ConnectionParameterResponseMapOutput struct{ *pulumi.OutputState }

func (ConnectionParameterResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConnectionParameterResponse)(nil)).Elem()
}

func (o ConnectionParameterResponseMapOutput) ToConnectionParameterResponseMapOutput() ConnectionParameterResponseMapOutput {
	return o
}

func (o ConnectionParameterResponseMapOutput) ToConnectionParameterResponseMapOutputWithContext(ctx context.Context) ConnectionParameterResponseMapOutput {
	return o
}

func (o ConnectionParameterResponseMapOutput) MapIndex(k pulumi.StringInput) ConnectionParameterResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ConnectionParameterResponse {
		return vs[0].(map[string]ConnectionParameterResponse)[vs[1].(string)]
	}).(ConnectionParameterResponseOutput)
}

type ConnectionStatus struct {
	Error    *ConnectionError  `pulumi:"error"`
	Id       *string           `pulumi:"id"`
	Kind     *string           `pulumi:"kind"`
	Location string            `pulumi:"location"`
	Name     *string           `pulumi:"name"`
	Status   *string           `pulumi:"status"`
	Tags     map[string]string `pulumi:"tags"`
	Target   *string           `pulumi:"target"`
	Type     *string           `pulumi:"type"`
}





type ConnectionStatusInput interface {
	pulumi.Input

	ToConnectionStatusOutput() ConnectionStatusOutput
	ToConnectionStatusOutputWithContext(context.Context) ConnectionStatusOutput
}

type ConnectionStatusArgs struct {
	Error    ConnectionErrorPtrInput `pulumi:"error"`
	Id       pulumi.StringPtrInput   `pulumi:"id"`
	Kind     pulumi.StringPtrInput   `pulumi:"kind"`
	Location pulumi.StringInput      `pulumi:"location"`
	Name     pulumi.StringPtrInput   `pulumi:"name"`
	Status   pulumi.StringPtrInput   `pulumi:"status"`
	Tags     pulumi.StringMapInput   `pulumi:"tags"`
	Target   pulumi.StringPtrInput   `pulumi:"target"`
	Type     pulumi.StringPtrInput   `pulumi:"type"`
}

func (ConnectionStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStatus)(nil)).Elem()
}

func (i ConnectionStatusArgs) ToConnectionStatusOutput() ConnectionStatusOutput {
	return i.ToConnectionStatusOutputWithContext(context.Background())
}

func (i ConnectionStatusArgs) ToConnectionStatusOutputWithContext(ctx context.Context) ConnectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStatusOutput)
}





type ConnectionStatusArrayInput interface {
	pulumi.Input

	ToConnectionStatusArrayOutput() ConnectionStatusArrayOutput
	ToConnectionStatusArrayOutputWithContext(context.Context) ConnectionStatusArrayOutput
}

type ConnectionStatusArray []ConnectionStatusInput

func (ConnectionStatusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectionStatus)(nil)).Elem()
}

func (i ConnectionStatusArray) ToConnectionStatusArrayOutput() ConnectionStatusArrayOutput {
	return i.ToConnectionStatusArrayOutputWithContext(context.Background())
}

func (i ConnectionStatusArray) ToConnectionStatusArrayOutputWithContext(ctx context.Context) ConnectionStatusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStatusArrayOutput)
}

type ConnectionStatusOutput struct{ *pulumi.OutputState }

func (ConnectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStatus)(nil)).Elem()
}

func (o ConnectionStatusOutput) ToConnectionStatusOutput() ConnectionStatusOutput {
	return o
}

func (o ConnectionStatusOutput) ToConnectionStatusOutputWithContext(ctx context.Context) ConnectionStatusOutput {
	return o
}

func (o ConnectionStatusOutput) Error() ConnectionErrorPtrOutput {
	return o.ApplyT(func(v ConnectionStatus) *ConnectionError { return v.Error }).(ConnectionErrorPtrOutput)
}

func (o ConnectionStatusOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStatus) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ConnectionStatusOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStatus) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ConnectionStatusOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionStatus) string { return v.Location }).(pulumi.StringOutput)
}

func (o ConnectionStatusOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStatus) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConnectionStatusOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStatus) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ConnectionStatusOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ConnectionStatus) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ConnectionStatusOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStatus) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o ConnectionStatusOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStatus) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ConnectionStatusArrayOutput struct{ *pulumi.OutputState }

func (ConnectionStatusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectionStatus)(nil)).Elem()
}

func (o ConnectionStatusArrayOutput) ToConnectionStatusArrayOutput() ConnectionStatusArrayOutput {
	return o
}

func (o ConnectionStatusArrayOutput) ToConnectionStatusArrayOutputWithContext(ctx context.Context) ConnectionStatusArrayOutput {
	return o
}

func (o ConnectionStatusArrayOutput) Index(i pulumi.IntInput) ConnectionStatusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConnectionStatus {
		return vs[0].([]ConnectionStatus)[vs[1].(int)]
	}).(ConnectionStatusOutput)
}

type ConnectionStatusResponse struct {
	Error    *ConnectionErrorResponse `pulumi:"error"`
	Id       *string                  `pulumi:"id"`
	Kind     *string                  `pulumi:"kind"`
	Location string                   `pulumi:"location"`
	Name     *string                  `pulumi:"name"`
	Status   *string                  `pulumi:"status"`
	Tags     map[string]string        `pulumi:"tags"`
	Target   *string                  `pulumi:"target"`
	Type     *string                  `pulumi:"type"`
}





type ConnectionStatusResponseInput interface {
	pulumi.Input

	ToConnectionStatusResponseOutput() ConnectionStatusResponseOutput
	ToConnectionStatusResponseOutputWithContext(context.Context) ConnectionStatusResponseOutput
}

type ConnectionStatusResponseArgs struct {
	Error    ConnectionErrorResponsePtrInput `pulumi:"error"`
	Id       pulumi.StringPtrInput           `pulumi:"id"`
	Kind     pulumi.StringPtrInput           `pulumi:"kind"`
	Location pulumi.StringInput              `pulumi:"location"`
	Name     pulumi.StringPtrInput           `pulumi:"name"`
	Status   pulumi.StringPtrInput           `pulumi:"status"`
	Tags     pulumi.StringMapInput           `pulumi:"tags"`
	Target   pulumi.StringPtrInput           `pulumi:"target"`
	Type     pulumi.StringPtrInput           `pulumi:"type"`
}

func (ConnectionStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStatusResponse)(nil)).Elem()
}

func (i ConnectionStatusResponseArgs) ToConnectionStatusResponseOutput() ConnectionStatusResponseOutput {
	return i.ToConnectionStatusResponseOutputWithContext(context.Background())
}

func (i ConnectionStatusResponseArgs) ToConnectionStatusResponseOutputWithContext(ctx context.Context) ConnectionStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStatusResponseOutput)
}





type ConnectionStatusResponseArrayInput interface {
	pulumi.Input

	ToConnectionStatusResponseArrayOutput() ConnectionStatusResponseArrayOutput
	ToConnectionStatusResponseArrayOutputWithContext(context.Context) ConnectionStatusResponseArrayOutput
}

type ConnectionStatusResponseArray []ConnectionStatusResponseInput

func (ConnectionStatusResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectionStatusResponse)(nil)).Elem()
}

func (i ConnectionStatusResponseArray) ToConnectionStatusResponseArrayOutput() ConnectionStatusResponseArrayOutput {
	return i.ToConnectionStatusResponseArrayOutputWithContext(context.Background())
}

func (i ConnectionStatusResponseArray) ToConnectionStatusResponseArrayOutputWithContext(ctx context.Context) ConnectionStatusResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStatusResponseArrayOutput)
}

type ConnectionStatusResponseOutput struct{ *pulumi.OutputState }

func (ConnectionStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStatusResponse)(nil)).Elem()
}

func (o ConnectionStatusResponseOutput) ToConnectionStatusResponseOutput() ConnectionStatusResponseOutput {
	return o
}

func (o ConnectionStatusResponseOutput) ToConnectionStatusResponseOutputWithContext(ctx context.Context) ConnectionStatusResponseOutput {
	return o
}

func (o ConnectionStatusResponseOutput) Error() ConnectionErrorResponsePtrOutput {
	return o.ApplyT(func(v ConnectionStatusResponse) *ConnectionErrorResponse { return v.Error }).(ConnectionErrorResponsePtrOutput)
}

func (o ConnectionStatusResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStatusResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ConnectionStatusResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStatusResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ConnectionStatusResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionStatusResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o ConnectionStatusResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStatusResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConnectionStatusResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStatusResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ConnectionStatusResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ConnectionStatusResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ConnectionStatusResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStatusResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o ConnectionStatusResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStatusResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ConnectionStatusResponseArrayOutput struct{ *pulumi.OutputState }

func (ConnectionStatusResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectionStatusResponse)(nil)).Elem()
}

func (o ConnectionStatusResponseArrayOutput) ToConnectionStatusResponseArrayOutput() ConnectionStatusResponseArrayOutput {
	return o
}

func (o ConnectionStatusResponseArrayOutput) ToConnectionStatusResponseArrayOutputWithContext(ctx context.Context) ConnectionStatusResponseArrayOutput {
	return o
}

func (o ConnectionStatusResponseArrayOutput) Index(i pulumi.IntInput) ConnectionStatusResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConnectionStatusResponse {
		return vs[0].([]ConnectionStatusResponse)[vs[1].(int)]
	}).(ConnectionStatusResponseOutput)
}

type ConsentLinkInputParameter struct {
	ObjectId      *string        `pulumi:"objectId"`
	ParameterName *string        `pulumi:"parameterName"`
	PrincipalType *PrincipalType `pulumi:"principalType"`
	RedirectUrl   *string        `pulumi:"redirectUrl"`
	TenantId      *string        `pulumi:"tenantId"`
}





type ConsentLinkInputParameterInput interface {
	pulumi.Input

	ToConsentLinkInputParameterOutput() ConsentLinkInputParameterOutput
	ToConsentLinkInputParameterOutputWithContext(context.Context) ConsentLinkInputParameterOutput
}

type ConsentLinkInputParameterArgs struct {
	ObjectId      pulumi.StringPtrInput `pulumi:"objectId"`
	ParameterName pulumi.StringPtrInput `pulumi:"parameterName"`
	PrincipalType PrincipalTypePtrInput `pulumi:"principalType"`
	RedirectUrl   pulumi.StringPtrInput `pulumi:"redirectUrl"`
	TenantId      pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (ConsentLinkInputParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsentLinkInputParameter)(nil)).Elem()
}

func (i ConsentLinkInputParameterArgs) ToConsentLinkInputParameterOutput() ConsentLinkInputParameterOutput {
	return i.ToConsentLinkInputParameterOutputWithContext(context.Background())
}

func (i ConsentLinkInputParameterArgs) ToConsentLinkInputParameterOutputWithContext(ctx context.Context) ConsentLinkInputParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsentLinkInputParameterOutput)
}





type ConsentLinkInputParameterArrayInput interface {
	pulumi.Input

	ToConsentLinkInputParameterArrayOutput() ConsentLinkInputParameterArrayOutput
	ToConsentLinkInputParameterArrayOutputWithContext(context.Context) ConsentLinkInputParameterArrayOutput
}

type ConsentLinkInputParameterArray []ConsentLinkInputParameterInput

func (ConsentLinkInputParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConsentLinkInputParameter)(nil)).Elem()
}

func (i ConsentLinkInputParameterArray) ToConsentLinkInputParameterArrayOutput() ConsentLinkInputParameterArrayOutput {
	return i.ToConsentLinkInputParameterArrayOutputWithContext(context.Background())
}

func (i ConsentLinkInputParameterArray) ToConsentLinkInputParameterArrayOutputWithContext(ctx context.Context) ConsentLinkInputParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsentLinkInputParameterArrayOutput)
}

type ConsentLinkInputParameterOutput struct{ *pulumi.OutputState }

func (ConsentLinkInputParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsentLinkInputParameter)(nil)).Elem()
}

func (o ConsentLinkInputParameterOutput) ToConsentLinkInputParameterOutput() ConsentLinkInputParameterOutput {
	return o
}

func (o ConsentLinkInputParameterOutput) ToConsentLinkInputParameterOutputWithContext(ctx context.Context) ConsentLinkInputParameterOutput {
	return o
}

func (o ConsentLinkInputParameterOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsentLinkInputParameter) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o ConsentLinkInputParameterOutput) ParameterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsentLinkInputParameter) *string { return v.ParameterName }).(pulumi.StringPtrOutput)
}

func (o ConsentLinkInputParameterOutput) PrincipalType() PrincipalTypePtrOutput {
	return o.ApplyT(func(v ConsentLinkInputParameter) *PrincipalType { return v.PrincipalType }).(PrincipalTypePtrOutput)
}

func (o ConsentLinkInputParameterOutput) RedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsentLinkInputParameter) *string { return v.RedirectUrl }).(pulumi.StringPtrOutput)
}

func (o ConsentLinkInputParameterOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsentLinkInputParameter) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type ConsentLinkInputParameterArrayOutput struct{ *pulumi.OutputState }

func (ConsentLinkInputParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConsentLinkInputParameter)(nil)).Elem()
}

func (o ConsentLinkInputParameterArrayOutput) ToConsentLinkInputParameterArrayOutput() ConsentLinkInputParameterArrayOutput {
	return o
}

func (o ConsentLinkInputParameterArrayOutput) ToConsentLinkInputParameterArrayOutputWithContext(ctx context.Context) ConsentLinkInputParameterArrayOutput {
	return o
}

func (o ConsentLinkInputParameterArrayOutput) Index(i pulumi.IntInput) ConsentLinkInputParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConsentLinkInputParameter {
		return vs[0].([]ConsentLinkInputParameter)[vs[1].(int)]
	}).(ConsentLinkInputParameterOutput)
}

type ConsentLinkResponse struct {
	DisplayName        *string `pulumi:"displayName"`
	FirstPartyLoginUri *string `pulumi:"firstPartyLoginUri"`
	Link               *string `pulumi:"link"`
	Status             *string `pulumi:"status"`
}





type ConsentLinkResponseInput interface {
	pulumi.Input

	ToConsentLinkResponseOutput() ConsentLinkResponseOutput
	ToConsentLinkResponseOutputWithContext(context.Context) ConsentLinkResponseOutput
}

type ConsentLinkResponseArgs struct {
	DisplayName        pulumi.StringPtrInput `pulumi:"displayName"`
	FirstPartyLoginUri pulumi.StringPtrInput `pulumi:"firstPartyLoginUri"`
	Link               pulumi.StringPtrInput `pulumi:"link"`
	Status             pulumi.StringPtrInput `pulumi:"status"`
}

func (ConsentLinkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsentLinkResponse)(nil)).Elem()
}

func (i ConsentLinkResponseArgs) ToConsentLinkResponseOutput() ConsentLinkResponseOutput {
	return i.ToConsentLinkResponseOutputWithContext(context.Background())
}

func (i ConsentLinkResponseArgs) ToConsentLinkResponseOutputWithContext(ctx context.Context) ConsentLinkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsentLinkResponseOutput)
}





type ConsentLinkResponseArrayInput interface {
	pulumi.Input

	ToConsentLinkResponseArrayOutput() ConsentLinkResponseArrayOutput
	ToConsentLinkResponseArrayOutputWithContext(context.Context) ConsentLinkResponseArrayOutput
}

type ConsentLinkResponseArray []ConsentLinkResponseInput

func (ConsentLinkResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConsentLinkResponse)(nil)).Elem()
}

func (i ConsentLinkResponseArray) ToConsentLinkResponseArrayOutput() ConsentLinkResponseArrayOutput {
	return i.ToConsentLinkResponseArrayOutputWithContext(context.Background())
}

func (i ConsentLinkResponseArray) ToConsentLinkResponseArrayOutputWithContext(ctx context.Context) ConsentLinkResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsentLinkResponseArrayOutput)
}

type ConsentLinkResponseOutput struct{ *pulumi.OutputState }

func (ConsentLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsentLinkResponse)(nil)).Elem()
}

func (o ConsentLinkResponseOutput) ToConsentLinkResponseOutput() ConsentLinkResponseOutput {
	return o
}

func (o ConsentLinkResponseOutput) ToConsentLinkResponseOutputWithContext(ctx context.Context) ConsentLinkResponseOutput {
	return o
}

func (o ConsentLinkResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsentLinkResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ConsentLinkResponseOutput) FirstPartyLoginUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsentLinkResponse) *string { return v.FirstPartyLoginUri }).(pulumi.StringPtrOutput)
}

func (o ConsentLinkResponseOutput) Link() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsentLinkResponse) *string { return v.Link }).(pulumi.StringPtrOutput)
}

func (o ConsentLinkResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsentLinkResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ConsentLinkResponseArrayOutput struct{ *pulumi.OutputState }

func (ConsentLinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConsentLinkResponse)(nil)).Elem()
}

func (o ConsentLinkResponseArrayOutput) ToConsentLinkResponseArrayOutput() ConsentLinkResponseArrayOutput {
	return o
}

func (o ConsentLinkResponseArrayOutput) ToConsentLinkResponseArrayOutputWithContext(ctx context.Context) ConsentLinkResponseArrayOutput {
	return o
}

func (o ConsentLinkResponseArrayOutput) Index(i pulumi.IntInput) ConsentLinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConsentLinkResponse {
		return vs[0].([]ConsentLinkResponse)[vs[1].(int)]
	}).(ConsentLinkResponseOutput)
}

type CustomLoginSettingValue struct {
	Id       *string           `pulumi:"id"`
	Kind     *string           `pulumi:"kind"`
	Location string            `pulumi:"location"`
	Name     *string           `pulumi:"name"`
	Option   *string           `pulumi:"option"`
	Tags     map[string]string `pulumi:"tags"`
	Type     *string           `pulumi:"type"`
}





type CustomLoginSettingValueInput interface {
	pulumi.Input

	ToCustomLoginSettingValueOutput() CustomLoginSettingValueOutput
	ToCustomLoginSettingValueOutputWithContext(context.Context) CustomLoginSettingValueOutput
}

type CustomLoginSettingValueArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Kind     pulumi.StringPtrInput `pulumi:"kind"`
	Location pulumi.StringInput    `pulumi:"location"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Option   pulumi.StringPtrInput `pulumi:"option"`
	Tags     pulumi.StringMapInput `pulumi:"tags"`
	Type     pulumi.StringPtrInput `pulumi:"type"`
}

func (CustomLoginSettingValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLoginSettingValue)(nil)).Elem()
}

func (i CustomLoginSettingValueArgs) ToCustomLoginSettingValueOutput() CustomLoginSettingValueOutput {
	return i.ToCustomLoginSettingValueOutputWithContext(context.Background())
}

func (i CustomLoginSettingValueArgs) ToCustomLoginSettingValueOutputWithContext(ctx context.Context) CustomLoginSettingValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLoginSettingValueOutput)
}





type CustomLoginSettingValueMapInput interface {
	pulumi.Input

	ToCustomLoginSettingValueMapOutput() CustomLoginSettingValueMapOutput
	ToCustomLoginSettingValueMapOutputWithContext(context.Context) CustomLoginSettingValueMapOutput
}

type CustomLoginSettingValueMap map[string]CustomLoginSettingValueInput

func (CustomLoginSettingValueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomLoginSettingValue)(nil)).Elem()
}

func (i CustomLoginSettingValueMap) ToCustomLoginSettingValueMapOutput() CustomLoginSettingValueMapOutput {
	return i.ToCustomLoginSettingValueMapOutputWithContext(context.Background())
}

func (i CustomLoginSettingValueMap) ToCustomLoginSettingValueMapOutputWithContext(ctx context.Context) CustomLoginSettingValueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLoginSettingValueMapOutput)
}

type CustomLoginSettingValueOutput struct{ *pulumi.OutputState }

func (CustomLoginSettingValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLoginSettingValue)(nil)).Elem()
}

func (o CustomLoginSettingValueOutput) ToCustomLoginSettingValueOutput() CustomLoginSettingValueOutput {
	return o
}

func (o CustomLoginSettingValueOutput) ToCustomLoginSettingValueOutputWithContext(ctx context.Context) CustomLoginSettingValueOutput {
	return o
}

func (o CustomLoginSettingValueOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLoginSettingValue) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o CustomLoginSettingValueOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLoginSettingValue) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o CustomLoginSettingValueOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v CustomLoginSettingValue) string { return v.Location }).(pulumi.StringOutput)
}

func (o CustomLoginSettingValueOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLoginSettingValue) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CustomLoginSettingValueOutput) Option() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLoginSettingValue) *string { return v.Option }).(pulumi.StringPtrOutput)
}

func (o CustomLoginSettingValueOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v CustomLoginSettingValue) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o CustomLoginSettingValueOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLoginSettingValue) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type CustomLoginSettingValueMapOutput struct{ *pulumi.OutputState }

func (CustomLoginSettingValueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomLoginSettingValue)(nil)).Elem()
}

func (o CustomLoginSettingValueMapOutput) ToCustomLoginSettingValueMapOutput() CustomLoginSettingValueMapOutput {
	return o
}

func (o CustomLoginSettingValueMapOutput) ToCustomLoginSettingValueMapOutputWithContext(ctx context.Context) CustomLoginSettingValueMapOutput {
	return o
}

func (o CustomLoginSettingValueMapOutput) MapIndex(k pulumi.StringInput) CustomLoginSettingValueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CustomLoginSettingValue {
		return vs[0].(map[string]CustomLoginSettingValue)[vs[1].(string)]
	}).(CustomLoginSettingValueOutput)
}

type CustomLoginSettingValueResponse struct {
	Id       *string           `pulumi:"id"`
	Kind     *string           `pulumi:"kind"`
	Location string            `pulumi:"location"`
	Name     *string           `pulumi:"name"`
	Option   *string           `pulumi:"option"`
	Tags     map[string]string `pulumi:"tags"`
	Type     *string           `pulumi:"type"`
}





type CustomLoginSettingValueResponseInput interface {
	pulumi.Input

	ToCustomLoginSettingValueResponseOutput() CustomLoginSettingValueResponseOutput
	ToCustomLoginSettingValueResponseOutputWithContext(context.Context) CustomLoginSettingValueResponseOutput
}

type CustomLoginSettingValueResponseArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Kind     pulumi.StringPtrInput `pulumi:"kind"`
	Location pulumi.StringInput    `pulumi:"location"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Option   pulumi.StringPtrInput `pulumi:"option"`
	Tags     pulumi.StringMapInput `pulumi:"tags"`
	Type     pulumi.StringPtrInput `pulumi:"type"`
}

func (CustomLoginSettingValueResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLoginSettingValueResponse)(nil)).Elem()
}

func (i CustomLoginSettingValueResponseArgs) ToCustomLoginSettingValueResponseOutput() CustomLoginSettingValueResponseOutput {
	return i.ToCustomLoginSettingValueResponseOutputWithContext(context.Background())
}

func (i CustomLoginSettingValueResponseArgs) ToCustomLoginSettingValueResponseOutputWithContext(ctx context.Context) CustomLoginSettingValueResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLoginSettingValueResponseOutput)
}





type CustomLoginSettingValueResponseMapInput interface {
	pulumi.Input

	ToCustomLoginSettingValueResponseMapOutput() CustomLoginSettingValueResponseMapOutput
	ToCustomLoginSettingValueResponseMapOutputWithContext(context.Context) CustomLoginSettingValueResponseMapOutput
}

type CustomLoginSettingValueResponseMap map[string]CustomLoginSettingValueResponseInput

func (CustomLoginSettingValueResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomLoginSettingValueResponse)(nil)).Elem()
}

func (i CustomLoginSettingValueResponseMap) ToCustomLoginSettingValueResponseMapOutput() CustomLoginSettingValueResponseMapOutput {
	return i.ToCustomLoginSettingValueResponseMapOutputWithContext(context.Background())
}

func (i CustomLoginSettingValueResponseMap) ToCustomLoginSettingValueResponseMapOutputWithContext(ctx context.Context) CustomLoginSettingValueResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLoginSettingValueResponseMapOutput)
}

type CustomLoginSettingValueResponseOutput struct{ *pulumi.OutputState }

func (CustomLoginSettingValueResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLoginSettingValueResponse)(nil)).Elem()
}

func (o CustomLoginSettingValueResponseOutput) ToCustomLoginSettingValueResponseOutput() CustomLoginSettingValueResponseOutput {
	return o
}

func (o CustomLoginSettingValueResponseOutput) ToCustomLoginSettingValueResponseOutputWithContext(ctx context.Context) CustomLoginSettingValueResponseOutput {
	return o
}

func (o CustomLoginSettingValueResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLoginSettingValueResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o CustomLoginSettingValueResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLoginSettingValueResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o CustomLoginSettingValueResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v CustomLoginSettingValueResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o CustomLoginSettingValueResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLoginSettingValueResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CustomLoginSettingValueResponseOutput) Option() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLoginSettingValueResponse) *string { return v.Option }).(pulumi.StringPtrOutput)
}

func (o CustomLoginSettingValueResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v CustomLoginSettingValueResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o CustomLoginSettingValueResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLoginSettingValueResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type CustomLoginSettingValueResponseMapOutput struct{ *pulumi.OutputState }

func (CustomLoginSettingValueResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomLoginSettingValueResponse)(nil)).Elem()
}

func (o CustomLoginSettingValueResponseMapOutput) ToCustomLoginSettingValueResponseMapOutput() CustomLoginSettingValueResponseMapOutput {
	return o
}

func (o CustomLoginSettingValueResponseMapOutput) ToCustomLoginSettingValueResponseMapOutputWithContext(ctx context.Context) CustomLoginSettingValueResponseMapOutput {
	return o
}

func (o CustomLoginSettingValueResponseMapOutput) MapIndex(k pulumi.StringInput) CustomLoginSettingValueResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CustomLoginSettingValueResponse {
		return vs[0].(map[string]CustomLoginSettingValueResponse)[vs[1].(string)]
	}).(CustomLoginSettingValueResponseOutput)
}

type ExpandedParentApiEntity struct {
	Entity   *ResponseMessageEnvelopeApiEntity `pulumi:"entity"`
	Id       *string                           `pulumi:"id"`
	Kind     *string                           `pulumi:"kind"`
	Location string                            `pulumi:"location"`
	Name     *string                           `pulumi:"name"`
	Tags     map[string]string                 `pulumi:"tags"`
	Type     *string                           `pulumi:"type"`
}





type ExpandedParentApiEntityInput interface {
	pulumi.Input

	ToExpandedParentApiEntityOutput() ExpandedParentApiEntityOutput
	ToExpandedParentApiEntityOutputWithContext(context.Context) ExpandedParentApiEntityOutput
}

type ExpandedParentApiEntityArgs struct {
	Entity   ResponseMessageEnvelopeApiEntityPtrInput `pulumi:"entity"`
	Id       pulumi.StringPtrInput                    `pulumi:"id"`
	Kind     pulumi.StringPtrInput                    `pulumi:"kind"`
	Location pulumi.StringInput                       `pulumi:"location"`
	Name     pulumi.StringPtrInput                    `pulumi:"name"`
	Tags     pulumi.StringMapInput                    `pulumi:"tags"`
	Type     pulumi.StringPtrInput                    `pulumi:"type"`
}

func (ExpandedParentApiEntityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpandedParentApiEntity)(nil)).Elem()
}

func (i ExpandedParentApiEntityArgs) ToExpandedParentApiEntityOutput() ExpandedParentApiEntityOutput {
	return i.ToExpandedParentApiEntityOutputWithContext(context.Background())
}

func (i ExpandedParentApiEntityArgs) ToExpandedParentApiEntityOutputWithContext(ctx context.Context) ExpandedParentApiEntityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpandedParentApiEntityOutput)
}

func (i ExpandedParentApiEntityArgs) ToExpandedParentApiEntityPtrOutput() ExpandedParentApiEntityPtrOutput {
	return i.ToExpandedParentApiEntityPtrOutputWithContext(context.Background())
}

func (i ExpandedParentApiEntityArgs) ToExpandedParentApiEntityPtrOutputWithContext(ctx context.Context) ExpandedParentApiEntityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpandedParentApiEntityOutput).ToExpandedParentApiEntityPtrOutputWithContext(ctx)
}









type ExpandedParentApiEntityPtrInput interface {
	pulumi.Input

	ToExpandedParentApiEntityPtrOutput() ExpandedParentApiEntityPtrOutput
	ToExpandedParentApiEntityPtrOutputWithContext(context.Context) ExpandedParentApiEntityPtrOutput
}

type expandedParentApiEntityPtrType ExpandedParentApiEntityArgs

func ExpandedParentApiEntityPtr(v *ExpandedParentApiEntityArgs) ExpandedParentApiEntityPtrInput {
	return (*expandedParentApiEntityPtrType)(v)
}

func (*expandedParentApiEntityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpandedParentApiEntity)(nil)).Elem()
}

func (i *expandedParentApiEntityPtrType) ToExpandedParentApiEntityPtrOutput() ExpandedParentApiEntityPtrOutput {
	return i.ToExpandedParentApiEntityPtrOutputWithContext(context.Background())
}

func (i *expandedParentApiEntityPtrType) ToExpandedParentApiEntityPtrOutputWithContext(ctx context.Context) ExpandedParentApiEntityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpandedParentApiEntityPtrOutput)
}

type ExpandedParentApiEntityOutput struct{ *pulumi.OutputState }

func (ExpandedParentApiEntityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpandedParentApiEntity)(nil)).Elem()
}

func (o ExpandedParentApiEntityOutput) ToExpandedParentApiEntityOutput() ExpandedParentApiEntityOutput {
	return o
}

func (o ExpandedParentApiEntityOutput) ToExpandedParentApiEntityOutputWithContext(ctx context.Context) ExpandedParentApiEntityOutput {
	return o
}

func (o ExpandedParentApiEntityOutput) ToExpandedParentApiEntityPtrOutput() ExpandedParentApiEntityPtrOutput {
	return o.ToExpandedParentApiEntityPtrOutputWithContext(context.Background())
}

func (o ExpandedParentApiEntityOutput) ToExpandedParentApiEntityPtrOutputWithContext(ctx context.Context) ExpandedParentApiEntityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpandedParentApiEntity) *ExpandedParentApiEntity {
		return &v
	}).(ExpandedParentApiEntityPtrOutput)
}

func (o ExpandedParentApiEntityOutput) Entity() ResponseMessageEnvelopeApiEntityPtrOutput {
	return o.ApplyT(func(v ExpandedParentApiEntity) *ResponseMessageEnvelopeApiEntity { return v.Entity }).(ResponseMessageEnvelopeApiEntityPtrOutput)
}

func (o ExpandedParentApiEntityOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpandedParentApiEntity) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ExpandedParentApiEntityOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpandedParentApiEntity) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ExpandedParentApiEntityOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ExpandedParentApiEntity) string { return v.Location }).(pulumi.StringOutput)
}

func (o ExpandedParentApiEntityOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpandedParentApiEntity) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExpandedParentApiEntityOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ExpandedParentApiEntity) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ExpandedParentApiEntityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpandedParentApiEntity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExpandedParentApiEntityPtrOutput struct{ *pulumi.OutputState }

func (ExpandedParentApiEntityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpandedParentApiEntity)(nil)).Elem()
}

func (o ExpandedParentApiEntityPtrOutput) ToExpandedParentApiEntityPtrOutput() ExpandedParentApiEntityPtrOutput {
	return o
}

func (o ExpandedParentApiEntityPtrOutput) ToExpandedParentApiEntityPtrOutputWithContext(ctx context.Context) ExpandedParentApiEntityPtrOutput {
	return o
}

func (o ExpandedParentApiEntityPtrOutput) Elem() ExpandedParentApiEntityOutput {
	return o.ApplyT(func(v *ExpandedParentApiEntity) ExpandedParentApiEntity {
		if v != nil {
			return *v
		}
		var ret ExpandedParentApiEntity
		return ret
	}).(ExpandedParentApiEntityOutput)
}

func (o ExpandedParentApiEntityPtrOutput) Entity() ResponseMessageEnvelopeApiEntityPtrOutput {
	return o.ApplyT(func(v *ExpandedParentApiEntity) *ResponseMessageEnvelopeApiEntity {
		if v == nil {
			return nil
		}
		return v.Entity
	}).(ResponseMessageEnvelopeApiEntityPtrOutput)
}

func (o ExpandedParentApiEntityPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpandedParentApiEntity) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ExpandedParentApiEntityPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpandedParentApiEntity) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o ExpandedParentApiEntityPtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpandedParentApiEntity) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o ExpandedParentApiEntityPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpandedParentApiEntity) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExpandedParentApiEntityPtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ExpandedParentApiEntity) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o ExpandedParentApiEntityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpandedParentApiEntity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ExpandedParentApiEntityResponse struct {
	Entity   *ResponseMessageEnvelopeApiEntityResponse `pulumi:"entity"`
	Id       *string                                   `pulumi:"id"`
	Kind     *string                                   `pulumi:"kind"`
	Location string                                    `pulumi:"location"`
	Name     *string                                   `pulumi:"name"`
	Tags     map[string]string                         `pulumi:"tags"`
	Type     *string                                   `pulumi:"type"`
}





type ExpandedParentApiEntityResponseInput interface {
	pulumi.Input

	ToExpandedParentApiEntityResponseOutput() ExpandedParentApiEntityResponseOutput
	ToExpandedParentApiEntityResponseOutputWithContext(context.Context) ExpandedParentApiEntityResponseOutput
}

type ExpandedParentApiEntityResponseArgs struct {
	Entity   ResponseMessageEnvelopeApiEntityResponsePtrInput `pulumi:"entity"`
	Id       pulumi.StringPtrInput                            `pulumi:"id"`
	Kind     pulumi.StringPtrInput                            `pulumi:"kind"`
	Location pulumi.StringInput                               `pulumi:"location"`
	Name     pulumi.StringPtrInput                            `pulumi:"name"`
	Tags     pulumi.StringMapInput                            `pulumi:"tags"`
	Type     pulumi.StringPtrInput                            `pulumi:"type"`
}

func (ExpandedParentApiEntityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpandedParentApiEntityResponse)(nil)).Elem()
}

func (i ExpandedParentApiEntityResponseArgs) ToExpandedParentApiEntityResponseOutput() ExpandedParentApiEntityResponseOutput {
	return i.ToExpandedParentApiEntityResponseOutputWithContext(context.Background())
}

func (i ExpandedParentApiEntityResponseArgs) ToExpandedParentApiEntityResponseOutputWithContext(ctx context.Context) ExpandedParentApiEntityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpandedParentApiEntityResponseOutput)
}

func (i ExpandedParentApiEntityResponseArgs) ToExpandedParentApiEntityResponsePtrOutput() ExpandedParentApiEntityResponsePtrOutput {
	return i.ToExpandedParentApiEntityResponsePtrOutputWithContext(context.Background())
}

func (i ExpandedParentApiEntityResponseArgs) ToExpandedParentApiEntityResponsePtrOutputWithContext(ctx context.Context) ExpandedParentApiEntityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpandedParentApiEntityResponseOutput).ToExpandedParentApiEntityResponsePtrOutputWithContext(ctx)
}









type ExpandedParentApiEntityResponsePtrInput interface {
	pulumi.Input

	ToExpandedParentApiEntityResponsePtrOutput() ExpandedParentApiEntityResponsePtrOutput
	ToExpandedParentApiEntityResponsePtrOutputWithContext(context.Context) ExpandedParentApiEntityResponsePtrOutput
}

type expandedParentApiEntityResponsePtrType ExpandedParentApiEntityResponseArgs

func ExpandedParentApiEntityResponsePtr(v *ExpandedParentApiEntityResponseArgs) ExpandedParentApiEntityResponsePtrInput {
	return (*expandedParentApiEntityResponsePtrType)(v)
}

func (*expandedParentApiEntityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpandedParentApiEntityResponse)(nil)).Elem()
}

func (i *expandedParentApiEntityResponsePtrType) ToExpandedParentApiEntityResponsePtrOutput() ExpandedParentApiEntityResponsePtrOutput {
	return i.ToExpandedParentApiEntityResponsePtrOutputWithContext(context.Background())
}

func (i *expandedParentApiEntityResponsePtrType) ToExpandedParentApiEntityResponsePtrOutputWithContext(ctx context.Context) ExpandedParentApiEntityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpandedParentApiEntityResponsePtrOutput)
}

type ExpandedParentApiEntityResponseOutput struct{ *pulumi.OutputState }

func (ExpandedParentApiEntityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpandedParentApiEntityResponse)(nil)).Elem()
}

func (o ExpandedParentApiEntityResponseOutput) ToExpandedParentApiEntityResponseOutput() ExpandedParentApiEntityResponseOutput {
	return o
}

func (o ExpandedParentApiEntityResponseOutput) ToExpandedParentApiEntityResponseOutputWithContext(ctx context.Context) ExpandedParentApiEntityResponseOutput {
	return o
}

func (o ExpandedParentApiEntityResponseOutput) ToExpandedParentApiEntityResponsePtrOutput() ExpandedParentApiEntityResponsePtrOutput {
	return o.ToExpandedParentApiEntityResponsePtrOutputWithContext(context.Background())
}

func (o ExpandedParentApiEntityResponseOutput) ToExpandedParentApiEntityResponsePtrOutputWithContext(ctx context.Context) ExpandedParentApiEntityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpandedParentApiEntityResponse) *ExpandedParentApiEntityResponse {
		return &v
	}).(ExpandedParentApiEntityResponsePtrOutput)
}

func (o ExpandedParentApiEntityResponseOutput) Entity() ResponseMessageEnvelopeApiEntityResponsePtrOutput {
	return o.ApplyT(func(v ExpandedParentApiEntityResponse) *ResponseMessageEnvelopeApiEntityResponse { return v.Entity }).(ResponseMessageEnvelopeApiEntityResponsePtrOutput)
}

func (o ExpandedParentApiEntityResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpandedParentApiEntityResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ExpandedParentApiEntityResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpandedParentApiEntityResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ExpandedParentApiEntityResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ExpandedParentApiEntityResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o ExpandedParentApiEntityResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpandedParentApiEntityResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExpandedParentApiEntityResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ExpandedParentApiEntityResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ExpandedParentApiEntityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpandedParentApiEntityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExpandedParentApiEntityResponsePtrOutput struct{ *pulumi.OutputState }

func (ExpandedParentApiEntityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpandedParentApiEntityResponse)(nil)).Elem()
}

func (o ExpandedParentApiEntityResponsePtrOutput) ToExpandedParentApiEntityResponsePtrOutput() ExpandedParentApiEntityResponsePtrOutput {
	return o
}

func (o ExpandedParentApiEntityResponsePtrOutput) ToExpandedParentApiEntityResponsePtrOutputWithContext(ctx context.Context) ExpandedParentApiEntityResponsePtrOutput {
	return o
}

func (o ExpandedParentApiEntityResponsePtrOutput) Elem() ExpandedParentApiEntityResponseOutput {
	return o.ApplyT(func(v *ExpandedParentApiEntityResponse) ExpandedParentApiEntityResponse {
		if v != nil {
			return *v
		}
		var ret ExpandedParentApiEntityResponse
		return ret
	}).(ExpandedParentApiEntityResponseOutput)
}

func (o ExpandedParentApiEntityResponsePtrOutput) Entity() ResponseMessageEnvelopeApiEntityResponsePtrOutput {
	return o.ApplyT(func(v *ExpandedParentApiEntityResponse) *ResponseMessageEnvelopeApiEntityResponse {
		if v == nil {
			return nil
		}
		return v.Entity
	}).(ResponseMessageEnvelopeApiEntityResponsePtrOutput)
}

func (o ExpandedParentApiEntityResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpandedParentApiEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ExpandedParentApiEntityResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpandedParentApiEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o ExpandedParentApiEntityResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpandedParentApiEntityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o ExpandedParentApiEntityResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpandedParentApiEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExpandedParentApiEntityResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ExpandedParentApiEntityResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o ExpandedParentApiEntityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpandedParentApiEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type GeneralApiInformation struct {
	ConnectionDisplayName *string           `pulumi:"connectionDisplayName"`
	ConnectionPortalUrl   interface{}       `pulumi:"connectionPortalUrl"`
	Description           *string           `pulumi:"description"`
	DisplayName           *string           `pulumi:"displayName"`
	IconUrl               *string           `pulumi:"iconUrl"`
	Id                    *string           `pulumi:"id"`
	Kind                  *string           `pulumi:"kind"`
	Location              string            `pulumi:"location"`
	Name                  *string           `pulumi:"name"`
	Tags                  map[string]string `pulumi:"tags"`
	TermsOfUseUrl         *string           `pulumi:"termsOfUseUrl"`
	Type                  *string           `pulumi:"type"`
}





type GeneralApiInformationInput interface {
	pulumi.Input

	ToGeneralApiInformationOutput() GeneralApiInformationOutput
	ToGeneralApiInformationOutputWithContext(context.Context) GeneralApiInformationOutput
}

type GeneralApiInformationArgs struct {
	ConnectionDisplayName pulumi.StringPtrInput `pulumi:"connectionDisplayName"`
	ConnectionPortalUrl   pulumi.Input          `pulumi:"connectionPortalUrl"`
	Description           pulumi.StringPtrInput `pulumi:"description"`
	DisplayName           pulumi.StringPtrInput `pulumi:"displayName"`
	IconUrl               pulumi.StringPtrInput `pulumi:"iconUrl"`
	Id                    pulumi.StringPtrInput `pulumi:"id"`
	Kind                  pulumi.StringPtrInput `pulumi:"kind"`
	Location              pulumi.StringInput    `pulumi:"location"`
	Name                  pulumi.StringPtrInput `pulumi:"name"`
	Tags                  pulumi.StringMapInput `pulumi:"tags"`
	TermsOfUseUrl         pulumi.StringPtrInput `pulumi:"termsOfUseUrl"`
	Type                  pulumi.StringPtrInput `pulumi:"type"`
}

func (GeneralApiInformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GeneralApiInformation)(nil)).Elem()
}

func (i GeneralApiInformationArgs) ToGeneralApiInformationOutput() GeneralApiInformationOutput {
	return i.ToGeneralApiInformationOutputWithContext(context.Background())
}

func (i GeneralApiInformationArgs) ToGeneralApiInformationOutputWithContext(ctx context.Context) GeneralApiInformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeneralApiInformationOutput)
}

func (i GeneralApiInformationArgs) ToGeneralApiInformationPtrOutput() GeneralApiInformationPtrOutput {
	return i.ToGeneralApiInformationPtrOutputWithContext(context.Background())
}

func (i GeneralApiInformationArgs) ToGeneralApiInformationPtrOutputWithContext(ctx context.Context) GeneralApiInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeneralApiInformationOutput).ToGeneralApiInformationPtrOutputWithContext(ctx)
}









type GeneralApiInformationPtrInput interface {
	pulumi.Input

	ToGeneralApiInformationPtrOutput() GeneralApiInformationPtrOutput
	ToGeneralApiInformationPtrOutputWithContext(context.Context) GeneralApiInformationPtrOutput
}

type generalApiInformationPtrType GeneralApiInformationArgs

func GeneralApiInformationPtr(v *GeneralApiInformationArgs) GeneralApiInformationPtrInput {
	return (*generalApiInformationPtrType)(v)
}

func (*generalApiInformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GeneralApiInformation)(nil)).Elem()
}

func (i *generalApiInformationPtrType) ToGeneralApiInformationPtrOutput() GeneralApiInformationPtrOutput {
	return i.ToGeneralApiInformationPtrOutputWithContext(context.Background())
}

func (i *generalApiInformationPtrType) ToGeneralApiInformationPtrOutputWithContext(ctx context.Context) GeneralApiInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeneralApiInformationPtrOutput)
}

type GeneralApiInformationOutput struct{ *pulumi.OutputState }

func (GeneralApiInformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GeneralApiInformation)(nil)).Elem()
}

func (o GeneralApiInformationOutput) ToGeneralApiInformationOutput() GeneralApiInformationOutput {
	return o
}

func (o GeneralApiInformationOutput) ToGeneralApiInformationOutputWithContext(ctx context.Context) GeneralApiInformationOutput {
	return o
}

func (o GeneralApiInformationOutput) ToGeneralApiInformationPtrOutput() GeneralApiInformationPtrOutput {
	return o.ToGeneralApiInformationPtrOutputWithContext(context.Background())
}

func (o GeneralApiInformationOutput) ToGeneralApiInformationPtrOutputWithContext(ctx context.Context) GeneralApiInformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GeneralApiInformation) *GeneralApiInformation {
		return &v
	}).(GeneralApiInformationPtrOutput)
}

func (o GeneralApiInformationOutput) ConnectionDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GeneralApiInformation) *string { return v.ConnectionDisplayName }).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationOutput) ConnectionPortalUrl() pulumi.AnyOutput {
	return o.ApplyT(func(v GeneralApiInformation) interface{} { return v.ConnectionPortalUrl }).(pulumi.AnyOutput)
}

func (o GeneralApiInformationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GeneralApiInformation) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GeneralApiInformation) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationOutput) IconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GeneralApiInformation) *string { return v.IconUrl }).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GeneralApiInformation) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GeneralApiInformation) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GeneralApiInformation) string { return v.Location }).(pulumi.StringOutput)
}

func (o GeneralApiInformationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GeneralApiInformation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GeneralApiInformation) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GeneralApiInformationOutput) TermsOfUseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GeneralApiInformation) *string { return v.TermsOfUseUrl }).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GeneralApiInformation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type GeneralApiInformationPtrOutput struct{ *pulumi.OutputState }

func (GeneralApiInformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GeneralApiInformation)(nil)).Elem()
}

func (o GeneralApiInformationPtrOutput) ToGeneralApiInformationPtrOutput() GeneralApiInformationPtrOutput {
	return o
}

func (o GeneralApiInformationPtrOutput) ToGeneralApiInformationPtrOutputWithContext(ctx context.Context) GeneralApiInformationPtrOutput {
	return o
}

func (o GeneralApiInformationPtrOutput) Elem() GeneralApiInformationOutput {
	return o.ApplyT(func(v *GeneralApiInformation) GeneralApiInformation {
		if v != nil {
			return *v
		}
		var ret GeneralApiInformation
		return ret
	}).(GeneralApiInformationOutput)
}

func (o GeneralApiInformationPtrOutput) ConnectionDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeneralApiInformation) *string {
		if v == nil {
			return nil
		}
		return v.ConnectionDisplayName
	}).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationPtrOutput) ConnectionPortalUrl() pulumi.AnyOutput {
	return o.ApplyT(func(v *GeneralApiInformation) interface{} {
		if v == nil {
			return nil
		}
		return v.ConnectionPortalUrl
	}).(pulumi.AnyOutput)
}

func (o GeneralApiInformationPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeneralApiInformation) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeneralApiInformation) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationPtrOutput) IconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeneralApiInformation) *string {
		if v == nil {
			return nil
		}
		return v.IconUrl
	}).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeneralApiInformation) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeneralApiInformation) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationPtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeneralApiInformation) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeneralApiInformation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationPtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GeneralApiInformation) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o GeneralApiInformationPtrOutput) TermsOfUseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeneralApiInformation) *string {
		if v == nil {
			return nil
		}
		return v.TermsOfUseUrl
	}).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeneralApiInformation) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type GeneralApiInformationResponse struct {
	ConnectionDisplayName *string           `pulumi:"connectionDisplayName"`
	ConnectionPortalUrl   interface{}       `pulumi:"connectionPortalUrl"`
	Description           *string           `pulumi:"description"`
	DisplayName           *string           `pulumi:"displayName"`
	IconUrl               *string           `pulumi:"iconUrl"`
	Id                    *string           `pulumi:"id"`
	Kind                  *string           `pulumi:"kind"`
	Location              string            `pulumi:"location"`
	Name                  *string           `pulumi:"name"`
	Tags                  map[string]string `pulumi:"tags"`
	TermsOfUseUrl         *string           `pulumi:"termsOfUseUrl"`
	Type                  *string           `pulumi:"type"`
}





type GeneralApiInformationResponseInput interface {
	pulumi.Input

	ToGeneralApiInformationResponseOutput() GeneralApiInformationResponseOutput
	ToGeneralApiInformationResponseOutputWithContext(context.Context) GeneralApiInformationResponseOutput
}

type GeneralApiInformationResponseArgs struct {
	ConnectionDisplayName pulumi.StringPtrInput `pulumi:"connectionDisplayName"`
	ConnectionPortalUrl   pulumi.Input          `pulumi:"connectionPortalUrl"`
	Description           pulumi.StringPtrInput `pulumi:"description"`
	DisplayName           pulumi.StringPtrInput `pulumi:"displayName"`
	IconUrl               pulumi.StringPtrInput `pulumi:"iconUrl"`
	Id                    pulumi.StringPtrInput `pulumi:"id"`
	Kind                  pulumi.StringPtrInput `pulumi:"kind"`
	Location              pulumi.StringInput    `pulumi:"location"`
	Name                  pulumi.StringPtrInput `pulumi:"name"`
	Tags                  pulumi.StringMapInput `pulumi:"tags"`
	TermsOfUseUrl         pulumi.StringPtrInput `pulumi:"termsOfUseUrl"`
	Type                  pulumi.StringPtrInput `pulumi:"type"`
}

func (GeneralApiInformationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GeneralApiInformationResponse)(nil)).Elem()
}

func (i GeneralApiInformationResponseArgs) ToGeneralApiInformationResponseOutput() GeneralApiInformationResponseOutput {
	return i.ToGeneralApiInformationResponseOutputWithContext(context.Background())
}

func (i GeneralApiInformationResponseArgs) ToGeneralApiInformationResponseOutputWithContext(ctx context.Context) GeneralApiInformationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeneralApiInformationResponseOutput)
}

func (i GeneralApiInformationResponseArgs) ToGeneralApiInformationResponsePtrOutput() GeneralApiInformationResponsePtrOutput {
	return i.ToGeneralApiInformationResponsePtrOutputWithContext(context.Background())
}

func (i GeneralApiInformationResponseArgs) ToGeneralApiInformationResponsePtrOutputWithContext(ctx context.Context) GeneralApiInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeneralApiInformationResponseOutput).ToGeneralApiInformationResponsePtrOutputWithContext(ctx)
}









type GeneralApiInformationResponsePtrInput interface {
	pulumi.Input

	ToGeneralApiInformationResponsePtrOutput() GeneralApiInformationResponsePtrOutput
	ToGeneralApiInformationResponsePtrOutputWithContext(context.Context) GeneralApiInformationResponsePtrOutput
}

type generalApiInformationResponsePtrType GeneralApiInformationResponseArgs

func GeneralApiInformationResponsePtr(v *GeneralApiInformationResponseArgs) GeneralApiInformationResponsePtrInput {
	return (*generalApiInformationResponsePtrType)(v)
}

func (*generalApiInformationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GeneralApiInformationResponse)(nil)).Elem()
}

func (i *generalApiInformationResponsePtrType) ToGeneralApiInformationResponsePtrOutput() GeneralApiInformationResponsePtrOutput {
	return i.ToGeneralApiInformationResponsePtrOutputWithContext(context.Background())
}

func (i *generalApiInformationResponsePtrType) ToGeneralApiInformationResponsePtrOutputWithContext(ctx context.Context) GeneralApiInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeneralApiInformationResponsePtrOutput)
}

type GeneralApiInformationResponseOutput struct{ *pulumi.OutputState }

func (GeneralApiInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GeneralApiInformationResponse)(nil)).Elem()
}

func (o GeneralApiInformationResponseOutput) ToGeneralApiInformationResponseOutput() GeneralApiInformationResponseOutput {
	return o
}

func (o GeneralApiInformationResponseOutput) ToGeneralApiInformationResponseOutputWithContext(ctx context.Context) GeneralApiInformationResponseOutput {
	return o
}

func (o GeneralApiInformationResponseOutput) ToGeneralApiInformationResponsePtrOutput() GeneralApiInformationResponsePtrOutput {
	return o.ToGeneralApiInformationResponsePtrOutputWithContext(context.Background())
}

func (o GeneralApiInformationResponseOutput) ToGeneralApiInformationResponsePtrOutputWithContext(ctx context.Context) GeneralApiInformationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GeneralApiInformationResponse) *GeneralApiInformationResponse {
		return &v
	}).(GeneralApiInformationResponsePtrOutput)
}

func (o GeneralApiInformationResponseOutput) ConnectionDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GeneralApiInformationResponse) *string { return v.ConnectionDisplayName }).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationResponseOutput) ConnectionPortalUrl() pulumi.AnyOutput {
	return o.ApplyT(func(v GeneralApiInformationResponse) interface{} { return v.ConnectionPortalUrl }).(pulumi.AnyOutput)
}

func (o GeneralApiInformationResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GeneralApiInformationResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GeneralApiInformationResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationResponseOutput) IconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GeneralApiInformationResponse) *string { return v.IconUrl }).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GeneralApiInformationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GeneralApiInformationResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GeneralApiInformationResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o GeneralApiInformationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GeneralApiInformationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GeneralApiInformationResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GeneralApiInformationResponseOutput) TermsOfUseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GeneralApiInformationResponse) *string { return v.TermsOfUseUrl }).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GeneralApiInformationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type GeneralApiInformationResponsePtrOutput struct{ *pulumi.OutputState }

func (GeneralApiInformationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GeneralApiInformationResponse)(nil)).Elem()
}

func (o GeneralApiInformationResponsePtrOutput) ToGeneralApiInformationResponsePtrOutput() GeneralApiInformationResponsePtrOutput {
	return o
}

func (o GeneralApiInformationResponsePtrOutput) ToGeneralApiInformationResponsePtrOutputWithContext(ctx context.Context) GeneralApiInformationResponsePtrOutput {
	return o
}

func (o GeneralApiInformationResponsePtrOutput) Elem() GeneralApiInformationResponseOutput {
	return o.ApplyT(func(v *GeneralApiInformationResponse) GeneralApiInformationResponse {
		if v != nil {
			return *v
		}
		var ret GeneralApiInformationResponse
		return ret
	}).(GeneralApiInformationResponseOutput)
}

func (o GeneralApiInformationResponsePtrOutput) ConnectionDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeneralApiInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConnectionDisplayName
	}).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationResponsePtrOutput) ConnectionPortalUrl() pulumi.AnyOutput {
	return o.ApplyT(func(v *GeneralApiInformationResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.ConnectionPortalUrl
	}).(pulumi.AnyOutput)
}

func (o GeneralApiInformationResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeneralApiInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationResponsePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeneralApiInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationResponsePtrOutput) IconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeneralApiInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.IconUrl
	}).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeneralApiInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeneralApiInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeneralApiInformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeneralApiInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GeneralApiInformationResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o GeneralApiInformationResponsePtrOutput) TermsOfUseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeneralApiInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.TermsOfUseUrl
	}).(pulumi.StringPtrOutput)
}

func (o GeneralApiInformationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeneralApiInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type HostingEnvironmentServiceDescriptions struct {
	HostId               *string `pulumi:"hostId"`
	HostingEnvironmentId *string `pulumi:"hostingEnvironmentId"`
	ServiceUrl           *string `pulumi:"serviceUrl"`
	UseInternalRouting   *bool   `pulumi:"useInternalRouting"`
}





type HostingEnvironmentServiceDescriptionsInput interface {
	pulumi.Input

	ToHostingEnvironmentServiceDescriptionsOutput() HostingEnvironmentServiceDescriptionsOutput
	ToHostingEnvironmentServiceDescriptionsOutputWithContext(context.Context) HostingEnvironmentServiceDescriptionsOutput
}

type HostingEnvironmentServiceDescriptionsArgs struct {
	HostId               pulumi.StringPtrInput `pulumi:"hostId"`
	HostingEnvironmentId pulumi.StringPtrInput `pulumi:"hostingEnvironmentId"`
	ServiceUrl           pulumi.StringPtrInput `pulumi:"serviceUrl"`
	UseInternalRouting   pulumi.BoolPtrInput   `pulumi:"useInternalRouting"`
}

func (HostingEnvironmentServiceDescriptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentServiceDescriptions)(nil)).Elem()
}

func (i HostingEnvironmentServiceDescriptionsArgs) ToHostingEnvironmentServiceDescriptionsOutput() HostingEnvironmentServiceDescriptionsOutput {
	return i.ToHostingEnvironmentServiceDescriptionsOutputWithContext(context.Background())
}

func (i HostingEnvironmentServiceDescriptionsArgs) ToHostingEnvironmentServiceDescriptionsOutputWithContext(ctx context.Context) HostingEnvironmentServiceDescriptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentServiceDescriptionsOutput)
}





type HostingEnvironmentServiceDescriptionsArrayInput interface {
	pulumi.Input

	ToHostingEnvironmentServiceDescriptionsArrayOutput() HostingEnvironmentServiceDescriptionsArrayOutput
	ToHostingEnvironmentServiceDescriptionsArrayOutputWithContext(context.Context) HostingEnvironmentServiceDescriptionsArrayOutput
}

type HostingEnvironmentServiceDescriptionsArray []HostingEnvironmentServiceDescriptionsInput

func (HostingEnvironmentServiceDescriptionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostingEnvironmentServiceDescriptions)(nil)).Elem()
}

func (i HostingEnvironmentServiceDescriptionsArray) ToHostingEnvironmentServiceDescriptionsArrayOutput() HostingEnvironmentServiceDescriptionsArrayOutput {
	return i.ToHostingEnvironmentServiceDescriptionsArrayOutputWithContext(context.Background())
}

func (i HostingEnvironmentServiceDescriptionsArray) ToHostingEnvironmentServiceDescriptionsArrayOutputWithContext(ctx context.Context) HostingEnvironmentServiceDescriptionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentServiceDescriptionsArrayOutput)
}

type HostingEnvironmentServiceDescriptionsOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentServiceDescriptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentServiceDescriptions)(nil)).Elem()
}

func (o HostingEnvironmentServiceDescriptionsOutput) ToHostingEnvironmentServiceDescriptionsOutput() HostingEnvironmentServiceDescriptionsOutput {
	return o
}

func (o HostingEnvironmentServiceDescriptionsOutput) ToHostingEnvironmentServiceDescriptionsOutputWithContext(ctx context.Context) HostingEnvironmentServiceDescriptionsOutput {
	return o
}

func (o HostingEnvironmentServiceDescriptionsOutput) HostId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentServiceDescriptions) *string { return v.HostId }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentServiceDescriptionsOutput) HostingEnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentServiceDescriptions) *string { return v.HostingEnvironmentId }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentServiceDescriptionsOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentServiceDescriptions) *string { return v.ServiceUrl }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentServiceDescriptionsOutput) UseInternalRouting() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentServiceDescriptions) *bool { return v.UseInternalRouting }).(pulumi.BoolPtrOutput)
}

type HostingEnvironmentServiceDescriptionsArrayOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentServiceDescriptionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostingEnvironmentServiceDescriptions)(nil)).Elem()
}

func (o HostingEnvironmentServiceDescriptionsArrayOutput) ToHostingEnvironmentServiceDescriptionsArrayOutput() HostingEnvironmentServiceDescriptionsArrayOutput {
	return o
}

func (o HostingEnvironmentServiceDescriptionsArrayOutput) ToHostingEnvironmentServiceDescriptionsArrayOutputWithContext(ctx context.Context) HostingEnvironmentServiceDescriptionsArrayOutput {
	return o
}

func (o HostingEnvironmentServiceDescriptionsArrayOutput) Index(i pulumi.IntInput) HostingEnvironmentServiceDescriptionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HostingEnvironmentServiceDescriptions {
		return vs[0].([]HostingEnvironmentServiceDescriptions)[vs[1].(int)]
	}).(HostingEnvironmentServiceDescriptionsOutput)
}

type HostingEnvironmentServiceDescriptionsResponse struct {
	HostId               *string `pulumi:"hostId"`
	HostingEnvironmentId *string `pulumi:"hostingEnvironmentId"`
	ServiceUrl           *string `pulumi:"serviceUrl"`
	UseInternalRouting   *bool   `pulumi:"useInternalRouting"`
}





type HostingEnvironmentServiceDescriptionsResponseInput interface {
	pulumi.Input

	ToHostingEnvironmentServiceDescriptionsResponseOutput() HostingEnvironmentServiceDescriptionsResponseOutput
	ToHostingEnvironmentServiceDescriptionsResponseOutputWithContext(context.Context) HostingEnvironmentServiceDescriptionsResponseOutput
}

type HostingEnvironmentServiceDescriptionsResponseArgs struct {
	HostId               pulumi.StringPtrInput `pulumi:"hostId"`
	HostingEnvironmentId pulumi.StringPtrInput `pulumi:"hostingEnvironmentId"`
	ServiceUrl           pulumi.StringPtrInput `pulumi:"serviceUrl"`
	UseInternalRouting   pulumi.BoolPtrInput   `pulumi:"useInternalRouting"`
}

func (HostingEnvironmentServiceDescriptionsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentServiceDescriptionsResponse)(nil)).Elem()
}

func (i HostingEnvironmentServiceDescriptionsResponseArgs) ToHostingEnvironmentServiceDescriptionsResponseOutput() HostingEnvironmentServiceDescriptionsResponseOutput {
	return i.ToHostingEnvironmentServiceDescriptionsResponseOutputWithContext(context.Background())
}

func (i HostingEnvironmentServiceDescriptionsResponseArgs) ToHostingEnvironmentServiceDescriptionsResponseOutputWithContext(ctx context.Context) HostingEnvironmentServiceDescriptionsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentServiceDescriptionsResponseOutput)
}





type HostingEnvironmentServiceDescriptionsResponseArrayInput interface {
	pulumi.Input

	ToHostingEnvironmentServiceDescriptionsResponseArrayOutput() HostingEnvironmentServiceDescriptionsResponseArrayOutput
	ToHostingEnvironmentServiceDescriptionsResponseArrayOutputWithContext(context.Context) HostingEnvironmentServiceDescriptionsResponseArrayOutput
}

type HostingEnvironmentServiceDescriptionsResponseArray []HostingEnvironmentServiceDescriptionsResponseInput

func (HostingEnvironmentServiceDescriptionsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostingEnvironmentServiceDescriptionsResponse)(nil)).Elem()
}

func (i HostingEnvironmentServiceDescriptionsResponseArray) ToHostingEnvironmentServiceDescriptionsResponseArrayOutput() HostingEnvironmentServiceDescriptionsResponseArrayOutput {
	return i.ToHostingEnvironmentServiceDescriptionsResponseArrayOutputWithContext(context.Background())
}

func (i HostingEnvironmentServiceDescriptionsResponseArray) ToHostingEnvironmentServiceDescriptionsResponseArrayOutputWithContext(ctx context.Context) HostingEnvironmentServiceDescriptionsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentServiceDescriptionsResponseArrayOutput)
}

type HostingEnvironmentServiceDescriptionsResponseOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentServiceDescriptionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentServiceDescriptionsResponse)(nil)).Elem()
}

func (o HostingEnvironmentServiceDescriptionsResponseOutput) ToHostingEnvironmentServiceDescriptionsResponseOutput() HostingEnvironmentServiceDescriptionsResponseOutput {
	return o
}

func (o HostingEnvironmentServiceDescriptionsResponseOutput) ToHostingEnvironmentServiceDescriptionsResponseOutputWithContext(ctx context.Context) HostingEnvironmentServiceDescriptionsResponseOutput {
	return o
}

func (o HostingEnvironmentServiceDescriptionsResponseOutput) HostId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentServiceDescriptionsResponse) *string { return v.HostId }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentServiceDescriptionsResponseOutput) HostingEnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentServiceDescriptionsResponse) *string { return v.HostingEnvironmentId }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentServiceDescriptionsResponseOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentServiceDescriptionsResponse) *string { return v.ServiceUrl }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentServiceDescriptionsResponseOutput) UseInternalRouting() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentServiceDescriptionsResponse) *bool { return v.UseInternalRouting }).(pulumi.BoolPtrOutput)
}

type HostingEnvironmentServiceDescriptionsResponseArrayOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentServiceDescriptionsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostingEnvironmentServiceDescriptionsResponse)(nil)).Elem()
}

func (o HostingEnvironmentServiceDescriptionsResponseArrayOutput) ToHostingEnvironmentServiceDescriptionsResponseArrayOutput() HostingEnvironmentServiceDescriptionsResponseArrayOutput {
	return o
}

func (o HostingEnvironmentServiceDescriptionsResponseArrayOutput) ToHostingEnvironmentServiceDescriptionsResponseArrayOutputWithContext(ctx context.Context) HostingEnvironmentServiceDescriptionsResponseArrayOutput {
	return o
}

func (o HostingEnvironmentServiceDescriptionsResponseArrayOutput) Index(i pulumi.IntInput) HostingEnvironmentServiceDescriptionsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HostingEnvironmentServiceDescriptionsResponse {
		return vs[0].([]HostingEnvironmentServiceDescriptionsResponse)[vs[1].(int)]
	}).(HostingEnvironmentServiceDescriptionsResponseOutput)
}

type ParameterCustomLoginSettingValues struct {
	CustomParameters map[string]CustomLoginSettingValue `pulumi:"customParameters"`
	Id               *string                            `pulumi:"id"`
	Kind             *string                            `pulumi:"kind"`
	Location         string                             `pulumi:"location"`
	Name             *string                            `pulumi:"name"`
	Tags             map[string]string                  `pulumi:"tags"`
	Type             *string                            `pulumi:"type"`
}





type ParameterCustomLoginSettingValuesInput interface {
	pulumi.Input

	ToParameterCustomLoginSettingValuesOutput() ParameterCustomLoginSettingValuesOutput
	ToParameterCustomLoginSettingValuesOutputWithContext(context.Context) ParameterCustomLoginSettingValuesOutput
}

type ParameterCustomLoginSettingValuesArgs struct {
	CustomParameters CustomLoginSettingValueMapInput `pulumi:"customParameters"`
	Id               pulumi.StringPtrInput           `pulumi:"id"`
	Kind             pulumi.StringPtrInput           `pulumi:"kind"`
	Location         pulumi.StringInput              `pulumi:"location"`
	Name             pulumi.StringPtrInput           `pulumi:"name"`
	Tags             pulumi.StringMapInput           `pulumi:"tags"`
	Type             pulumi.StringPtrInput           `pulumi:"type"`
}

func (ParameterCustomLoginSettingValuesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterCustomLoginSettingValues)(nil)).Elem()
}

func (i ParameterCustomLoginSettingValuesArgs) ToParameterCustomLoginSettingValuesOutput() ParameterCustomLoginSettingValuesOutput {
	return i.ToParameterCustomLoginSettingValuesOutputWithContext(context.Background())
}

func (i ParameterCustomLoginSettingValuesArgs) ToParameterCustomLoginSettingValuesOutputWithContext(ctx context.Context) ParameterCustomLoginSettingValuesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterCustomLoginSettingValuesOutput)
}





type ParameterCustomLoginSettingValuesMapInput interface {
	pulumi.Input

	ToParameterCustomLoginSettingValuesMapOutput() ParameterCustomLoginSettingValuesMapOutput
	ToParameterCustomLoginSettingValuesMapOutputWithContext(context.Context) ParameterCustomLoginSettingValuesMapOutput
}

type ParameterCustomLoginSettingValuesMap map[string]ParameterCustomLoginSettingValuesInput

func (ParameterCustomLoginSettingValuesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterCustomLoginSettingValues)(nil)).Elem()
}

func (i ParameterCustomLoginSettingValuesMap) ToParameterCustomLoginSettingValuesMapOutput() ParameterCustomLoginSettingValuesMapOutput {
	return i.ToParameterCustomLoginSettingValuesMapOutputWithContext(context.Background())
}

func (i ParameterCustomLoginSettingValuesMap) ToParameterCustomLoginSettingValuesMapOutputWithContext(ctx context.Context) ParameterCustomLoginSettingValuesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterCustomLoginSettingValuesMapOutput)
}

type ParameterCustomLoginSettingValuesOutput struct{ *pulumi.OutputState }

func (ParameterCustomLoginSettingValuesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterCustomLoginSettingValues)(nil)).Elem()
}

func (o ParameterCustomLoginSettingValuesOutput) ToParameterCustomLoginSettingValuesOutput() ParameterCustomLoginSettingValuesOutput {
	return o
}

func (o ParameterCustomLoginSettingValuesOutput) ToParameterCustomLoginSettingValuesOutputWithContext(ctx context.Context) ParameterCustomLoginSettingValuesOutput {
	return o
}

func (o ParameterCustomLoginSettingValuesOutput) CustomParameters() CustomLoginSettingValueMapOutput {
	return o.ApplyT(func(v ParameterCustomLoginSettingValues) map[string]CustomLoginSettingValue {
		return v.CustomParameters
	}).(CustomLoginSettingValueMapOutput)
}

func (o ParameterCustomLoginSettingValuesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterCustomLoginSettingValues) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ParameterCustomLoginSettingValuesOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterCustomLoginSettingValues) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ParameterCustomLoginSettingValuesOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterCustomLoginSettingValues) string { return v.Location }).(pulumi.StringOutput)
}

func (o ParameterCustomLoginSettingValuesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterCustomLoginSettingValues) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ParameterCustomLoginSettingValuesOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ParameterCustomLoginSettingValues) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ParameterCustomLoginSettingValuesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterCustomLoginSettingValues) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ParameterCustomLoginSettingValuesMapOutput struct{ *pulumi.OutputState }

func (ParameterCustomLoginSettingValuesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterCustomLoginSettingValues)(nil)).Elem()
}

func (o ParameterCustomLoginSettingValuesMapOutput) ToParameterCustomLoginSettingValuesMapOutput() ParameterCustomLoginSettingValuesMapOutput {
	return o
}

func (o ParameterCustomLoginSettingValuesMapOutput) ToParameterCustomLoginSettingValuesMapOutputWithContext(ctx context.Context) ParameterCustomLoginSettingValuesMapOutput {
	return o
}

func (o ParameterCustomLoginSettingValuesMapOutput) MapIndex(k pulumi.StringInput) ParameterCustomLoginSettingValuesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterCustomLoginSettingValues {
		return vs[0].(map[string]ParameterCustomLoginSettingValues)[vs[1].(string)]
	}).(ParameterCustomLoginSettingValuesOutput)
}

type ParameterCustomLoginSettingValuesResponse struct {
	CustomParameters map[string]CustomLoginSettingValueResponse `pulumi:"customParameters"`
	Id               *string                                    `pulumi:"id"`
	Kind             *string                                    `pulumi:"kind"`
	Location         string                                     `pulumi:"location"`
	Name             *string                                    `pulumi:"name"`
	Tags             map[string]string                          `pulumi:"tags"`
	Type             *string                                    `pulumi:"type"`
}





type ParameterCustomLoginSettingValuesResponseInput interface {
	pulumi.Input

	ToParameterCustomLoginSettingValuesResponseOutput() ParameterCustomLoginSettingValuesResponseOutput
	ToParameterCustomLoginSettingValuesResponseOutputWithContext(context.Context) ParameterCustomLoginSettingValuesResponseOutput
}

type ParameterCustomLoginSettingValuesResponseArgs struct {
	CustomParameters CustomLoginSettingValueResponseMapInput `pulumi:"customParameters"`
	Id               pulumi.StringPtrInput                   `pulumi:"id"`
	Kind             pulumi.StringPtrInput                   `pulumi:"kind"`
	Location         pulumi.StringInput                      `pulumi:"location"`
	Name             pulumi.StringPtrInput                   `pulumi:"name"`
	Tags             pulumi.StringMapInput                   `pulumi:"tags"`
	Type             pulumi.StringPtrInput                   `pulumi:"type"`
}

func (ParameterCustomLoginSettingValuesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterCustomLoginSettingValuesResponse)(nil)).Elem()
}

func (i ParameterCustomLoginSettingValuesResponseArgs) ToParameterCustomLoginSettingValuesResponseOutput() ParameterCustomLoginSettingValuesResponseOutput {
	return i.ToParameterCustomLoginSettingValuesResponseOutputWithContext(context.Background())
}

func (i ParameterCustomLoginSettingValuesResponseArgs) ToParameterCustomLoginSettingValuesResponseOutputWithContext(ctx context.Context) ParameterCustomLoginSettingValuesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterCustomLoginSettingValuesResponseOutput)
}





type ParameterCustomLoginSettingValuesResponseMapInput interface {
	pulumi.Input

	ToParameterCustomLoginSettingValuesResponseMapOutput() ParameterCustomLoginSettingValuesResponseMapOutput
	ToParameterCustomLoginSettingValuesResponseMapOutputWithContext(context.Context) ParameterCustomLoginSettingValuesResponseMapOutput
}

type ParameterCustomLoginSettingValuesResponseMap map[string]ParameterCustomLoginSettingValuesResponseInput

func (ParameterCustomLoginSettingValuesResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterCustomLoginSettingValuesResponse)(nil)).Elem()
}

func (i ParameterCustomLoginSettingValuesResponseMap) ToParameterCustomLoginSettingValuesResponseMapOutput() ParameterCustomLoginSettingValuesResponseMapOutput {
	return i.ToParameterCustomLoginSettingValuesResponseMapOutputWithContext(context.Background())
}

func (i ParameterCustomLoginSettingValuesResponseMap) ToParameterCustomLoginSettingValuesResponseMapOutputWithContext(ctx context.Context) ParameterCustomLoginSettingValuesResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterCustomLoginSettingValuesResponseMapOutput)
}

type ParameterCustomLoginSettingValuesResponseOutput struct{ *pulumi.OutputState }

func (ParameterCustomLoginSettingValuesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterCustomLoginSettingValuesResponse)(nil)).Elem()
}

func (o ParameterCustomLoginSettingValuesResponseOutput) ToParameterCustomLoginSettingValuesResponseOutput() ParameterCustomLoginSettingValuesResponseOutput {
	return o
}

func (o ParameterCustomLoginSettingValuesResponseOutput) ToParameterCustomLoginSettingValuesResponseOutputWithContext(ctx context.Context) ParameterCustomLoginSettingValuesResponseOutput {
	return o
}

func (o ParameterCustomLoginSettingValuesResponseOutput) CustomParameters() CustomLoginSettingValueResponseMapOutput {
	return o.ApplyT(func(v ParameterCustomLoginSettingValuesResponse) map[string]CustomLoginSettingValueResponse {
		return v.CustomParameters
	}).(CustomLoginSettingValueResponseMapOutput)
}

func (o ParameterCustomLoginSettingValuesResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterCustomLoginSettingValuesResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ParameterCustomLoginSettingValuesResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterCustomLoginSettingValuesResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ParameterCustomLoginSettingValuesResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterCustomLoginSettingValuesResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o ParameterCustomLoginSettingValuesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterCustomLoginSettingValuesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ParameterCustomLoginSettingValuesResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ParameterCustomLoginSettingValuesResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ParameterCustomLoginSettingValuesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterCustomLoginSettingValuesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ParameterCustomLoginSettingValuesResponseMapOutput struct{ *pulumi.OutputState }

func (ParameterCustomLoginSettingValuesResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterCustomLoginSettingValuesResponse)(nil)).Elem()
}

func (o ParameterCustomLoginSettingValuesResponseMapOutput) ToParameterCustomLoginSettingValuesResponseMapOutput() ParameterCustomLoginSettingValuesResponseMapOutput {
	return o
}

func (o ParameterCustomLoginSettingValuesResponseMapOutput) ToParameterCustomLoginSettingValuesResponseMapOutputWithContext(ctx context.Context) ParameterCustomLoginSettingValuesResponseMapOutput {
	return o
}

func (o ParameterCustomLoginSettingValuesResponseMapOutput) MapIndex(k pulumi.StringInput) ParameterCustomLoginSettingValuesResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterCustomLoginSettingValuesResponse {
		return vs[0].(map[string]ParameterCustomLoginSettingValuesResponse)[vs[1].(string)]
	}).(ParameterCustomLoginSettingValuesResponseOutput)
}

type ResponseMessageEnvelopeApiEntity struct {
	Id         *string           `pulumi:"id"`
	Location   *string           `pulumi:"location"`
	Name       *string           `pulumi:"name"`
	Plan       *ArmPlan          `pulumi:"plan"`
	Properties *ApiEntity        `pulumi:"properties"`
	Sku        *SkuDescription   `pulumi:"sku"`
	Tags       map[string]string `pulumi:"tags"`
	Type       *string           `pulumi:"type"`
}





type ResponseMessageEnvelopeApiEntityInput interface {
	pulumi.Input

	ToResponseMessageEnvelopeApiEntityOutput() ResponseMessageEnvelopeApiEntityOutput
	ToResponseMessageEnvelopeApiEntityOutputWithContext(context.Context) ResponseMessageEnvelopeApiEntityOutput
}

type ResponseMessageEnvelopeApiEntityArgs struct {
	Id         pulumi.StringPtrInput  `pulumi:"id"`
	Location   pulumi.StringPtrInput  `pulumi:"location"`
	Name       pulumi.StringPtrInput  `pulumi:"name"`
	Plan       ArmPlanPtrInput        `pulumi:"plan"`
	Properties ApiEntityPtrInput      `pulumi:"properties"`
	Sku        SkuDescriptionPtrInput `pulumi:"sku"`
	Tags       pulumi.StringMapInput  `pulumi:"tags"`
	Type       pulumi.StringPtrInput  `pulumi:"type"`
}

func (ResponseMessageEnvelopeApiEntityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResponseMessageEnvelopeApiEntity)(nil)).Elem()
}

func (i ResponseMessageEnvelopeApiEntityArgs) ToResponseMessageEnvelopeApiEntityOutput() ResponseMessageEnvelopeApiEntityOutput {
	return i.ToResponseMessageEnvelopeApiEntityOutputWithContext(context.Background())
}

func (i ResponseMessageEnvelopeApiEntityArgs) ToResponseMessageEnvelopeApiEntityOutputWithContext(ctx context.Context) ResponseMessageEnvelopeApiEntityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponseMessageEnvelopeApiEntityOutput)
}

func (i ResponseMessageEnvelopeApiEntityArgs) ToResponseMessageEnvelopeApiEntityPtrOutput() ResponseMessageEnvelopeApiEntityPtrOutput {
	return i.ToResponseMessageEnvelopeApiEntityPtrOutputWithContext(context.Background())
}

func (i ResponseMessageEnvelopeApiEntityArgs) ToResponseMessageEnvelopeApiEntityPtrOutputWithContext(ctx context.Context) ResponseMessageEnvelopeApiEntityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponseMessageEnvelopeApiEntityOutput).ToResponseMessageEnvelopeApiEntityPtrOutputWithContext(ctx)
}









type ResponseMessageEnvelopeApiEntityPtrInput interface {
	pulumi.Input

	ToResponseMessageEnvelopeApiEntityPtrOutput() ResponseMessageEnvelopeApiEntityPtrOutput
	ToResponseMessageEnvelopeApiEntityPtrOutputWithContext(context.Context) ResponseMessageEnvelopeApiEntityPtrOutput
}

type responseMessageEnvelopeApiEntityPtrType ResponseMessageEnvelopeApiEntityArgs

func ResponseMessageEnvelopeApiEntityPtr(v *ResponseMessageEnvelopeApiEntityArgs) ResponseMessageEnvelopeApiEntityPtrInput {
	return (*responseMessageEnvelopeApiEntityPtrType)(v)
}

func (*responseMessageEnvelopeApiEntityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResponseMessageEnvelopeApiEntity)(nil)).Elem()
}

func (i *responseMessageEnvelopeApiEntityPtrType) ToResponseMessageEnvelopeApiEntityPtrOutput() ResponseMessageEnvelopeApiEntityPtrOutput {
	return i.ToResponseMessageEnvelopeApiEntityPtrOutputWithContext(context.Background())
}

func (i *responseMessageEnvelopeApiEntityPtrType) ToResponseMessageEnvelopeApiEntityPtrOutputWithContext(ctx context.Context) ResponseMessageEnvelopeApiEntityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponseMessageEnvelopeApiEntityPtrOutput)
}

type ResponseMessageEnvelopeApiEntityOutput struct{ *pulumi.OutputState }

func (ResponseMessageEnvelopeApiEntityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResponseMessageEnvelopeApiEntity)(nil)).Elem()
}

func (o ResponseMessageEnvelopeApiEntityOutput) ToResponseMessageEnvelopeApiEntityOutput() ResponseMessageEnvelopeApiEntityOutput {
	return o
}

func (o ResponseMessageEnvelopeApiEntityOutput) ToResponseMessageEnvelopeApiEntityOutputWithContext(ctx context.Context) ResponseMessageEnvelopeApiEntityOutput {
	return o
}

func (o ResponseMessageEnvelopeApiEntityOutput) ToResponseMessageEnvelopeApiEntityPtrOutput() ResponseMessageEnvelopeApiEntityPtrOutput {
	return o.ToResponseMessageEnvelopeApiEntityPtrOutputWithContext(context.Background())
}

func (o ResponseMessageEnvelopeApiEntityOutput) ToResponseMessageEnvelopeApiEntityPtrOutputWithContext(ctx context.Context) ResponseMessageEnvelopeApiEntityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResponseMessageEnvelopeApiEntity) *ResponseMessageEnvelopeApiEntity {
		return &v
	}).(ResponseMessageEnvelopeApiEntityPtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeApiEntity) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeApiEntity) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeApiEntity) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityOutput) Plan() ArmPlanPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeApiEntity) *ArmPlan { return v.Plan }).(ArmPlanPtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityOutput) Properties() ApiEntityPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeApiEntity) *ApiEntity { return v.Properties }).(ApiEntityPtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityOutput) Sku() SkuDescriptionPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeApiEntity) *SkuDescription { return v.Sku }).(SkuDescriptionPtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeApiEntity) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ResponseMessageEnvelopeApiEntityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeApiEntity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ResponseMessageEnvelopeApiEntityPtrOutput struct{ *pulumi.OutputState }

func (ResponseMessageEnvelopeApiEntityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResponseMessageEnvelopeApiEntity)(nil)).Elem()
}

func (o ResponseMessageEnvelopeApiEntityPtrOutput) ToResponseMessageEnvelopeApiEntityPtrOutput() ResponseMessageEnvelopeApiEntityPtrOutput {
	return o
}

func (o ResponseMessageEnvelopeApiEntityPtrOutput) ToResponseMessageEnvelopeApiEntityPtrOutputWithContext(ctx context.Context) ResponseMessageEnvelopeApiEntityPtrOutput {
	return o
}

func (o ResponseMessageEnvelopeApiEntityPtrOutput) Elem() ResponseMessageEnvelopeApiEntityOutput {
	return o.ApplyT(func(v *ResponseMessageEnvelopeApiEntity) ResponseMessageEnvelopeApiEntity {
		if v != nil {
			return *v
		}
		var ret ResponseMessageEnvelopeApiEntity
		return ret
	}).(ResponseMessageEnvelopeApiEntityOutput)
}

func (o ResponseMessageEnvelopeApiEntityPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResponseMessageEnvelopeApiEntity) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityPtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResponseMessageEnvelopeApiEntity) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResponseMessageEnvelopeApiEntity) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityPtrOutput) Plan() ArmPlanPtrOutput {
	return o.ApplyT(func(v *ResponseMessageEnvelopeApiEntity) *ArmPlan {
		if v == nil {
			return nil
		}
		return v.Plan
	}).(ArmPlanPtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityPtrOutput) Properties() ApiEntityPtrOutput {
	return o.ApplyT(func(v *ResponseMessageEnvelopeApiEntity) *ApiEntity {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(ApiEntityPtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityPtrOutput) Sku() SkuDescriptionPtrOutput {
	return o.ApplyT(func(v *ResponseMessageEnvelopeApiEntity) *SkuDescription {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(SkuDescriptionPtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityPtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResponseMessageEnvelopeApiEntity) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o ResponseMessageEnvelopeApiEntityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResponseMessageEnvelopeApiEntity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ResponseMessageEnvelopeApiEntityResponse struct {
	Id         *string                 `pulumi:"id"`
	Location   *string                 `pulumi:"location"`
	Name       *string                 `pulumi:"name"`
	Plan       *ArmPlanResponse        `pulumi:"plan"`
	Properties *ApiEntityResponse      `pulumi:"properties"`
	Sku        *SkuDescriptionResponse `pulumi:"sku"`
	Tags       map[string]string       `pulumi:"tags"`
	Type       *string                 `pulumi:"type"`
}





type ResponseMessageEnvelopeApiEntityResponseInput interface {
	pulumi.Input

	ToResponseMessageEnvelopeApiEntityResponseOutput() ResponseMessageEnvelopeApiEntityResponseOutput
	ToResponseMessageEnvelopeApiEntityResponseOutputWithContext(context.Context) ResponseMessageEnvelopeApiEntityResponseOutput
}

type ResponseMessageEnvelopeApiEntityResponseArgs struct {
	Id         pulumi.StringPtrInput          `pulumi:"id"`
	Location   pulumi.StringPtrInput          `pulumi:"location"`
	Name       pulumi.StringPtrInput          `pulumi:"name"`
	Plan       ArmPlanResponsePtrInput        `pulumi:"plan"`
	Properties ApiEntityResponsePtrInput      `pulumi:"properties"`
	Sku        SkuDescriptionResponsePtrInput `pulumi:"sku"`
	Tags       pulumi.StringMapInput          `pulumi:"tags"`
	Type       pulumi.StringPtrInput          `pulumi:"type"`
}

func (ResponseMessageEnvelopeApiEntityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResponseMessageEnvelopeApiEntityResponse)(nil)).Elem()
}

func (i ResponseMessageEnvelopeApiEntityResponseArgs) ToResponseMessageEnvelopeApiEntityResponseOutput() ResponseMessageEnvelopeApiEntityResponseOutput {
	return i.ToResponseMessageEnvelopeApiEntityResponseOutputWithContext(context.Background())
}

func (i ResponseMessageEnvelopeApiEntityResponseArgs) ToResponseMessageEnvelopeApiEntityResponseOutputWithContext(ctx context.Context) ResponseMessageEnvelopeApiEntityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponseMessageEnvelopeApiEntityResponseOutput)
}

func (i ResponseMessageEnvelopeApiEntityResponseArgs) ToResponseMessageEnvelopeApiEntityResponsePtrOutput() ResponseMessageEnvelopeApiEntityResponsePtrOutput {
	return i.ToResponseMessageEnvelopeApiEntityResponsePtrOutputWithContext(context.Background())
}

func (i ResponseMessageEnvelopeApiEntityResponseArgs) ToResponseMessageEnvelopeApiEntityResponsePtrOutputWithContext(ctx context.Context) ResponseMessageEnvelopeApiEntityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponseMessageEnvelopeApiEntityResponseOutput).ToResponseMessageEnvelopeApiEntityResponsePtrOutputWithContext(ctx)
}









type ResponseMessageEnvelopeApiEntityResponsePtrInput interface {
	pulumi.Input

	ToResponseMessageEnvelopeApiEntityResponsePtrOutput() ResponseMessageEnvelopeApiEntityResponsePtrOutput
	ToResponseMessageEnvelopeApiEntityResponsePtrOutputWithContext(context.Context) ResponseMessageEnvelopeApiEntityResponsePtrOutput
}

type responseMessageEnvelopeApiEntityResponsePtrType ResponseMessageEnvelopeApiEntityResponseArgs

func ResponseMessageEnvelopeApiEntityResponsePtr(v *ResponseMessageEnvelopeApiEntityResponseArgs) ResponseMessageEnvelopeApiEntityResponsePtrInput {
	return (*responseMessageEnvelopeApiEntityResponsePtrType)(v)
}

func (*responseMessageEnvelopeApiEntityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResponseMessageEnvelopeApiEntityResponse)(nil)).Elem()
}

func (i *responseMessageEnvelopeApiEntityResponsePtrType) ToResponseMessageEnvelopeApiEntityResponsePtrOutput() ResponseMessageEnvelopeApiEntityResponsePtrOutput {
	return i.ToResponseMessageEnvelopeApiEntityResponsePtrOutputWithContext(context.Background())
}

func (i *responseMessageEnvelopeApiEntityResponsePtrType) ToResponseMessageEnvelopeApiEntityResponsePtrOutputWithContext(ctx context.Context) ResponseMessageEnvelopeApiEntityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponseMessageEnvelopeApiEntityResponsePtrOutput)
}

type ResponseMessageEnvelopeApiEntityResponseOutput struct{ *pulumi.OutputState }

func (ResponseMessageEnvelopeApiEntityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResponseMessageEnvelopeApiEntityResponse)(nil)).Elem()
}

func (o ResponseMessageEnvelopeApiEntityResponseOutput) ToResponseMessageEnvelopeApiEntityResponseOutput() ResponseMessageEnvelopeApiEntityResponseOutput {
	return o
}

func (o ResponseMessageEnvelopeApiEntityResponseOutput) ToResponseMessageEnvelopeApiEntityResponseOutputWithContext(ctx context.Context) ResponseMessageEnvelopeApiEntityResponseOutput {
	return o
}

func (o ResponseMessageEnvelopeApiEntityResponseOutput) ToResponseMessageEnvelopeApiEntityResponsePtrOutput() ResponseMessageEnvelopeApiEntityResponsePtrOutput {
	return o.ToResponseMessageEnvelopeApiEntityResponsePtrOutputWithContext(context.Background())
}

func (o ResponseMessageEnvelopeApiEntityResponseOutput) ToResponseMessageEnvelopeApiEntityResponsePtrOutputWithContext(ctx context.Context) ResponseMessageEnvelopeApiEntityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResponseMessageEnvelopeApiEntityResponse) *ResponseMessageEnvelopeApiEntityResponse {
		return &v
	}).(ResponseMessageEnvelopeApiEntityResponsePtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeApiEntityResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeApiEntityResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeApiEntityResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityResponseOutput) Plan() ArmPlanResponsePtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeApiEntityResponse) *ArmPlanResponse { return v.Plan }).(ArmPlanResponsePtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityResponseOutput) Properties() ApiEntityResponsePtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeApiEntityResponse) *ApiEntityResponse { return v.Properties }).(ApiEntityResponsePtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityResponseOutput) Sku() SkuDescriptionResponsePtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeApiEntityResponse) *SkuDescriptionResponse { return v.Sku }).(SkuDescriptionResponsePtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeApiEntityResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ResponseMessageEnvelopeApiEntityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeApiEntityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ResponseMessageEnvelopeApiEntityResponsePtrOutput struct{ *pulumi.OutputState }

func (ResponseMessageEnvelopeApiEntityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResponseMessageEnvelopeApiEntityResponse)(nil)).Elem()
}

func (o ResponseMessageEnvelopeApiEntityResponsePtrOutput) ToResponseMessageEnvelopeApiEntityResponsePtrOutput() ResponseMessageEnvelopeApiEntityResponsePtrOutput {
	return o
}

func (o ResponseMessageEnvelopeApiEntityResponsePtrOutput) ToResponseMessageEnvelopeApiEntityResponsePtrOutputWithContext(ctx context.Context) ResponseMessageEnvelopeApiEntityResponsePtrOutput {
	return o
}

func (o ResponseMessageEnvelopeApiEntityResponsePtrOutput) Elem() ResponseMessageEnvelopeApiEntityResponseOutput {
	return o.ApplyT(func(v *ResponseMessageEnvelopeApiEntityResponse) ResponseMessageEnvelopeApiEntityResponse {
		if v != nil {
			return *v
		}
		var ret ResponseMessageEnvelopeApiEntityResponse
		return ret
	}).(ResponseMessageEnvelopeApiEntityResponseOutput)
}

func (o ResponseMessageEnvelopeApiEntityResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResponseMessageEnvelopeApiEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResponseMessageEnvelopeApiEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResponseMessageEnvelopeApiEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityResponsePtrOutput) Plan() ArmPlanResponsePtrOutput {
	return o.ApplyT(func(v *ResponseMessageEnvelopeApiEntityResponse) *ArmPlanResponse {
		if v == nil {
			return nil
		}
		return v.Plan
	}).(ArmPlanResponsePtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityResponsePtrOutput) Properties() ApiEntityResponsePtrOutput {
	return o.ApplyT(func(v *ResponseMessageEnvelopeApiEntityResponse) *ApiEntityResponse {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(ApiEntityResponsePtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityResponsePtrOutput) Sku() SkuDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *ResponseMessageEnvelopeApiEntityResponse) *SkuDescriptionResponse {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(SkuDescriptionResponsePtrOutput)
}

func (o ResponseMessageEnvelopeApiEntityResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResponseMessageEnvelopeApiEntityResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o ResponseMessageEnvelopeApiEntityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResponseMessageEnvelopeApiEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type SkuDescription struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     *string `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuDescriptionInput interface {
	pulumi.Input

	ToSkuDescriptionOutput() SkuDescriptionOutput
	ToSkuDescriptionOutputWithContext(context.Context) SkuDescriptionOutput
}

type SkuDescriptionArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuDescription)(nil)).Elem()
}

func (i SkuDescriptionArgs) ToSkuDescriptionOutput() SkuDescriptionOutput {
	return i.ToSkuDescriptionOutputWithContext(context.Background())
}

func (i SkuDescriptionArgs) ToSkuDescriptionOutputWithContext(ctx context.Context) SkuDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionOutput)
}

func (i SkuDescriptionArgs) ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput {
	return i.ToSkuDescriptionPtrOutputWithContext(context.Background())
}

func (i SkuDescriptionArgs) ToSkuDescriptionPtrOutputWithContext(ctx context.Context) SkuDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionOutput).ToSkuDescriptionPtrOutputWithContext(ctx)
}









type SkuDescriptionPtrInput interface {
	pulumi.Input

	ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput
	ToSkuDescriptionPtrOutputWithContext(context.Context) SkuDescriptionPtrOutput
}

type skuDescriptionPtrType SkuDescriptionArgs

func SkuDescriptionPtr(v *SkuDescriptionArgs) SkuDescriptionPtrInput {
	return (*skuDescriptionPtrType)(v)
}

func (*skuDescriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuDescription)(nil)).Elem()
}

func (i *skuDescriptionPtrType) ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput {
	return i.ToSkuDescriptionPtrOutputWithContext(context.Background())
}

func (i *skuDescriptionPtrType) ToSkuDescriptionPtrOutputWithContext(ctx context.Context) SkuDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionPtrOutput)
}

type SkuDescriptionOutput struct{ *pulumi.OutputState }

func (SkuDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuDescription)(nil)).Elem()
}

func (o SkuDescriptionOutput) ToSkuDescriptionOutput() SkuDescriptionOutput {
	return o
}

func (o SkuDescriptionOutput) ToSkuDescriptionOutputWithContext(ctx context.Context) SkuDescriptionOutput {
	return o
}

func (o SkuDescriptionOutput) ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput {
	return o.ToSkuDescriptionPtrOutputWithContext(context.Background())
}

func (o SkuDescriptionOutput) ToSkuDescriptionPtrOutputWithContext(ctx context.Context) SkuDescriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuDescription) *SkuDescription {
		return &v
	}).(SkuDescriptionPtrOutput)
}

func (o SkuDescriptionOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuDescription) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuDescriptionOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescription) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescription) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescription) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescription) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuDescriptionPtrOutput struct{ *pulumi.OutputState }

func (SkuDescriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuDescription)(nil)).Elem()
}

func (o SkuDescriptionPtrOutput) ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput {
	return o
}

func (o SkuDescriptionPtrOutput) ToSkuDescriptionPtrOutputWithContext(ctx context.Context) SkuDescriptionPtrOutput {
	return o
}

func (o SkuDescriptionPtrOutput) Elem() SkuDescriptionOutput {
	return o.ApplyT(func(v *SkuDescription) SkuDescription {
		if v != nil {
			return *v
		}
		var ret SkuDescription
		return ret
	}).(SkuDescriptionOutput)
}

func (o SkuDescriptionPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuDescriptionPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuDescriptionResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     *string `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuDescriptionResponseInput interface {
	pulumi.Input

	ToSkuDescriptionResponseOutput() SkuDescriptionResponseOutput
	ToSkuDescriptionResponseOutputWithContext(context.Context) SkuDescriptionResponseOutput
}

type SkuDescriptionResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuDescriptionResponse)(nil)).Elem()
}

func (i SkuDescriptionResponseArgs) ToSkuDescriptionResponseOutput() SkuDescriptionResponseOutput {
	return i.ToSkuDescriptionResponseOutputWithContext(context.Background())
}

func (i SkuDescriptionResponseArgs) ToSkuDescriptionResponseOutputWithContext(ctx context.Context) SkuDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionResponseOutput)
}

func (i SkuDescriptionResponseArgs) ToSkuDescriptionResponsePtrOutput() SkuDescriptionResponsePtrOutput {
	return i.ToSkuDescriptionResponsePtrOutputWithContext(context.Background())
}

func (i SkuDescriptionResponseArgs) ToSkuDescriptionResponsePtrOutputWithContext(ctx context.Context) SkuDescriptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionResponseOutput).ToSkuDescriptionResponsePtrOutputWithContext(ctx)
}









type SkuDescriptionResponsePtrInput interface {
	pulumi.Input

	ToSkuDescriptionResponsePtrOutput() SkuDescriptionResponsePtrOutput
	ToSkuDescriptionResponsePtrOutputWithContext(context.Context) SkuDescriptionResponsePtrOutput
}

type skuDescriptionResponsePtrType SkuDescriptionResponseArgs

func SkuDescriptionResponsePtr(v *SkuDescriptionResponseArgs) SkuDescriptionResponsePtrInput {
	return (*skuDescriptionResponsePtrType)(v)
}

func (*skuDescriptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuDescriptionResponse)(nil)).Elem()
}

func (i *skuDescriptionResponsePtrType) ToSkuDescriptionResponsePtrOutput() SkuDescriptionResponsePtrOutput {
	return i.ToSkuDescriptionResponsePtrOutputWithContext(context.Background())
}

func (i *skuDescriptionResponsePtrType) ToSkuDescriptionResponsePtrOutputWithContext(ctx context.Context) SkuDescriptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionResponsePtrOutput)
}

type SkuDescriptionResponseOutput struct{ *pulumi.OutputState }

func (SkuDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuDescriptionResponse)(nil)).Elem()
}

func (o SkuDescriptionResponseOutput) ToSkuDescriptionResponseOutput() SkuDescriptionResponseOutput {
	return o
}

func (o SkuDescriptionResponseOutput) ToSkuDescriptionResponseOutputWithContext(ctx context.Context) SkuDescriptionResponseOutput {
	return o
}

func (o SkuDescriptionResponseOutput) ToSkuDescriptionResponsePtrOutput() SkuDescriptionResponsePtrOutput {
	return o.ToSkuDescriptionResponsePtrOutputWithContext(context.Background())
}

func (o SkuDescriptionResponseOutput) ToSkuDescriptionResponsePtrOutputWithContext(ctx context.Context) SkuDescriptionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuDescriptionResponse) *SkuDescriptionResponse {
		return &v
	}).(SkuDescriptionResponsePtrOutput)
}

func (o SkuDescriptionResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuDescriptionResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuDescriptionResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuDescriptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuDescriptionResponse)(nil)).Elem()
}

func (o SkuDescriptionResponsePtrOutput) ToSkuDescriptionResponsePtrOutput() SkuDescriptionResponsePtrOutput {
	return o
}

func (o SkuDescriptionResponsePtrOutput) ToSkuDescriptionResponsePtrOutputWithContext(ctx context.Context) SkuDescriptionResponsePtrOutput {
	return o
}

func (o SkuDescriptionResponsePtrOutput) Elem() SkuDescriptionResponseOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) SkuDescriptionResponse {
		if v != nil {
			return *v
		}
		var ret SkuDescriptionResponse
		return ret
	}).(SkuDescriptionResponseOutput)
}

func (o SkuDescriptionResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiEntityInput)(nil)).Elem(), ApiEntityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiEntityPtrInput)(nil)).Elem(), ApiEntityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiEntityResponseInput)(nil)).Elem(), ApiEntityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiEntityResponsePtrInput)(nil)).Elem(), ApiEntityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiOAuthSettingsInput)(nil)).Elem(), ApiOAuthSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiOAuthSettingsPtrInput)(nil)).Elem(), ApiOAuthSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiOAuthSettingsParameterInput)(nil)).Elem(), ApiOAuthSettingsParameterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiOAuthSettingsParameterMapInput)(nil)).Elem(), ApiOAuthSettingsParameterMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiOAuthSettingsParameterResponseInput)(nil)).Elem(), ApiOAuthSettingsParameterResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiOAuthSettingsParameterResponseMapInput)(nil)).Elem(), ApiOAuthSettingsParameterResponseMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiOAuthSettingsResponseInput)(nil)).Elem(), ApiOAuthSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiOAuthSettingsResponsePtrInput)(nil)).Elem(), ApiOAuthSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiPoliciesInput)(nil)).Elem(), ApiPoliciesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiPoliciesPtrInput)(nil)).Elem(), ApiPoliciesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiPoliciesResponseInput)(nil)).Elem(), ApiPoliciesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiPoliciesResponsePtrInput)(nil)).Elem(), ApiPoliciesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArmPlanInput)(nil)).Elem(), ArmPlanArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArmPlanPtrInput)(nil)).Elem(), ArmPlanArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArmPlanResponseInput)(nil)).Elem(), ArmPlanResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArmPlanResponsePtrInput)(nil)).Elem(), ArmPlanResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendServiceDefinitionInput)(nil)).Elem(), BackendServiceDefinitionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendServiceDefinitionPtrInput)(nil)).Elem(), BackendServiceDefinitionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendServiceDefinitionResponseInput)(nil)).Elem(), BackendServiceDefinitionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendServiceDefinitionResponsePtrInput)(nil)).Elem(), BackendServiceDefinitionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionErrorInput)(nil)).Elem(), ConnectionErrorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionErrorPtrInput)(nil)).Elem(), ConnectionErrorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionErrorResponseInput)(nil)).Elem(), ConnectionErrorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionErrorResponsePtrInput)(nil)).Elem(), ConnectionErrorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionParameterInput)(nil)).Elem(), ConnectionParameterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionParameterMapInput)(nil)).Elem(), ConnectionParameterMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionParameterResponseInput)(nil)).Elem(), ConnectionParameterResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionParameterResponseMapInput)(nil)).Elem(), ConnectionParameterResponseMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionStatusInput)(nil)).Elem(), ConnectionStatusArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionStatusArrayInput)(nil)).Elem(), ConnectionStatusArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionStatusResponseInput)(nil)).Elem(), ConnectionStatusResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionStatusResponseArrayInput)(nil)).Elem(), ConnectionStatusResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsentLinkInputParameterInput)(nil)).Elem(), ConsentLinkInputParameterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsentLinkInputParameterArrayInput)(nil)).Elem(), ConsentLinkInputParameterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsentLinkResponseInput)(nil)).Elem(), ConsentLinkResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsentLinkResponseArrayInput)(nil)).Elem(), ConsentLinkResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomLoginSettingValueInput)(nil)).Elem(), CustomLoginSettingValueArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomLoginSettingValueMapInput)(nil)).Elem(), CustomLoginSettingValueMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomLoginSettingValueResponseInput)(nil)).Elem(), CustomLoginSettingValueResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomLoginSettingValueResponseMapInput)(nil)).Elem(), CustomLoginSettingValueResponseMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpandedParentApiEntityInput)(nil)).Elem(), ExpandedParentApiEntityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpandedParentApiEntityPtrInput)(nil)).Elem(), ExpandedParentApiEntityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpandedParentApiEntityResponseInput)(nil)).Elem(), ExpandedParentApiEntityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpandedParentApiEntityResponsePtrInput)(nil)).Elem(), ExpandedParentApiEntityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GeneralApiInformationInput)(nil)).Elem(), GeneralApiInformationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GeneralApiInformationPtrInput)(nil)).Elem(), GeneralApiInformationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GeneralApiInformationResponseInput)(nil)).Elem(), GeneralApiInformationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GeneralApiInformationResponsePtrInput)(nil)).Elem(), GeneralApiInformationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostingEnvironmentServiceDescriptionsInput)(nil)).Elem(), HostingEnvironmentServiceDescriptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostingEnvironmentServiceDescriptionsArrayInput)(nil)).Elem(), HostingEnvironmentServiceDescriptionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostingEnvironmentServiceDescriptionsResponseInput)(nil)).Elem(), HostingEnvironmentServiceDescriptionsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostingEnvironmentServiceDescriptionsResponseArrayInput)(nil)).Elem(), HostingEnvironmentServiceDescriptionsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterCustomLoginSettingValuesInput)(nil)).Elem(), ParameterCustomLoginSettingValuesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterCustomLoginSettingValuesMapInput)(nil)).Elem(), ParameterCustomLoginSettingValuesMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterCustomLoginSettingValuesResponseInput)(nil)).Elem(), ParameterCustomLoginSettingValuesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterCustomLoginSettingValuesResponseMapInput)(nil)).Elem(), ParameterCustomLoginSettingValuesResponseMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResponseMessageEnvelopeApiEntityInput)(nil)).Elem(), ResponseMessageEnvelopeApiEntityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResponseMessageEnvelopeApiEntityPtrInput)(nil)).Elem(), ResponseMessageEnvelopeApiEntityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResponseMessageEnvelopeApiEntityResponseInput)(nil)).Elem(), ResponseMessageEnvelopeApiEntityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResponseMessageEnvelopeApiEntityResponsePtrInput)(nil)).Elem(), ResponseMessageEnvelopeApiEntityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuDescriptionInput)(nil)).Elem(), SkuDescriptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuDescriptionPtrInput)(nil)).Elem(), SkuDescriptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuDescriptionResponseInput)(nil)).Elem(), SkuDescriptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuDescriptionResponsePtrInput)(nil)).Elem(), SkuDescriptionResponseArgs{})
	pulumi.RegisterOutputType(ApiEntityOutput{})
	pulumi.RegisterOutputType(ApiEntityPtrOutput{})
	pulumi.RegisterOutputType(ApiEntityResponseOutput{})
	pulumi.RegisterOutputType(ApiEntityResponsePtrOutput{})
	pulumi.RegisterOutputType(ApiOAuthSettingsOutput{})
	pulumi.RegisterOutputType(ApiOAuthSettingsPtrOutput{})
	pulumi.RegisterOutputType(ApiOAuthSettingsParameterOutput{})
	pulumi.RegisterOutputType(ApiOAuthSettingsParameterMapOutput{})
	pulumi.RegisterOutputType(ApiOAuthSettingsParameterResponseOutput{})
	pulumi.RegisterOutputType(ApiOAuthSettingsParameterResponseMapOutput{})
	pulumi.RegisterOutputType(ApiOAuthSettingsResponseOutput{})
	pulumi.RegisterOutputType(ApiOAuthSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(ApiPoliciesOutput{})
	pulumi.RegisterOutputType(ApiPoliciesPtrOutput{})
	pulumi.RegisterOutputType(ApiPoliciesResponseOutput{})
	pulumi.RegisterOutputType(ApiPoliciesResponsePtrOutput{})
	pulumi.RegisterOutputType(ArmPlanOutput{})
	pulumi.RegisterOutputType(ArmPlanPtrOutput{})
	pulumi.RegisterOutputType(ArmPlanResponseOutput{})
	pulumi.RegisterOutputType(ArmPlanResponsePtrOutput{})
	pulumi.RegisterOutputType(BackendServiceDefinitionOutput{})
	pulumi.RegisterOutputType(BackendServiceDefinitionPtrOutput{})
	pulumi.RegisterOutputType(BackendServiceDefinitionResponseOutput{})
	pulumi.RegisterOutputType(BackendServiceDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectionErrorOutput{})
	pulumi.RegisterOutputType(ConnectionErrorPtrOutput{})
	pulumi.RegisterOutputType(ConnectionErrorResponseOutput{})
	pulumi.RegisterOutputType(ConnectionErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectionParameterOutput{})
	pulumi.RegisterOutputType(ConnectionParameterMapOutput{})
	pulumi.RegisterOutputType(ConnectionParameterResponseOutput{})
	pulumi.RegisterOutputType(ConnectionParameterResponseMapOutput{})
	pulumi.RegisterOutputType(ConnectionStatusOutput{})
	pulumi.RegisterOutputType(ConnectionStatusArrayOutput{})
	pulumi.RegisterOutputType(ConnectionStatusResponseOutput{})
	pulumi.RegisterOutputType(ConnectionStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(ConsentLinkInputParameterOutput{})
	pulumi.RegisterOutputType(ConsentLinkInputParameterArrayOutput{})
	pulumi.RegisterOutputType(ConsentLinkResponseOutput{})
	pulumi.RegisterOutputType(ConsentLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(CustomLoginSettingValueOutput{})
	pulumi.RegisterOutputType(CustomLoginSettingValueMapOutput{})
	pulumi.RegisterOutputType(CustomLoginSettingValueResponseOutput{})
	pulumi.RegisterOutputType(CustomLoginSettingValueResponseMapOutput{})
	pulumi.RegisterOutputType(ExpandedParentApiEntityOutput{})
	pulumi.RegisterOutputType(ExpandedParentApiEntityPtrOutput{})
	pulumi.RegisterOutputType(ExpandedParentApiEntityResponseOutput{})
	pulumi.RegisterOutputType(ExpandedParentApiEntityResponsePtrOutput{})
	pulumi.RegisterOutputType(GeneralApiInformationOutput{})
	pulumi.RegisterOutputType(GeneralApiInformationPtrOutput{})
	pulumi.RegisterOutputType(GeneralApiInformationResponseOutput{})
	pulumi.RegisterOutputType(GeneralApiInformationResponsePtrOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentServiceDescriptionsOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentServiceDescriptionsArrayOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentServiceDescriptionsResponseOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentServiceDescriptionsResponseArrayOutput{})
	pulumi.RegisterOutputType(ParameterCustomLoginSettingValuesOutput{})
	pulumi.RegisterOutputType(ParameterCustomLoginSettingValuesMapOutput{})
	pulumi.RegisterOutputType(ParameterCustomLoginSettingValuesResponseOutput{})
	pulumi.RegisterOutputType(ParameterCustomLoginSettingValuesResponseMapOutput{})
	pulumi.RegisterOutputType(ResponseMessageEnvelopeApiEntityOutput{})
	pulumi.RegisterOutputType(ResponseMessageEnvelopeApiEntityPtrOutput{})
	pulumi.RegisterOutputType(ResponseMessageEnvelopeApiEntityResponseOutput{})
	pulumi.RegisterOutputType(ResponseMessageEnvelopeApiEntityResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuDescriptionOutput{})
	pulumi.RegisterOutputType(SkuDescriptionPtrOutput{})
	pulumi.RegisterOutputType(SkuDescriptionResponseOutput{})
	pulumi.RegisterOutputType(SkuDescriptionResponsePtrOutput{})
}

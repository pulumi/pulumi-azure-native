


package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DicomServiceAuthenticationConfigurationResponse struct {
	Audiences []string `pulumi:"audiences"`
	Authority string   `pulumi:"authority"`
}

type DicomServiceAuthenticationConfigurationResponseOutput struct{ *pulumi.OutputState }

func (DicomServiceAuthenticationConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DicomServiceAuthenticationConfigurationResponse)(nil)).Elem()
}

func (o DicomServiceAuthenticationConfigurationResponseOutput) ToDicomServiceAuthenticationConfigurationResponseOutput() DicomServiceAuthenticationConfigurationResponseOutput {
	return o
}

func (o DicomServiceAuthenticationConfigurationResponseOutput) ToDicomServiceAuthenticationConfigurationResponseOutputWithContext(ctx context.Context) DicomServiceAuthenticationConfigurationResponseOutput {
	return o
}

func (o DicomServiceAuthenticationConfigurationResponseOutput) Audiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DicomServiceAuthenticationConfigurationResponse) []string { return v.Audiences }).(pulumi.StringArrayOutput)
}

func (o DicomServiceAuthenticationConfigurationResponseOutput) Authority() pulumi.StringOutput {
	return o.ApplyT(func(v DicomServiceAuthenticationConfigurationResponse) string { return v.Authority }).(pulumi.StringOutput)
}

type DicomServiceAuthenticationConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (DicomServiceAuthenticationConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DicomServiceAuthenticationConfigurationResponse)(nil)).Elem()
}

func (o DicomServiceAuthenticationConfigurationResponsePtrOutput) ToDicomServiceAuthenticationConfigurationResponsePtrOutput() DicomServiceAuthenticationConfigurationResponsePtrOutput {
	return o
}

func (o DicomServiceAuthenticationConfigurationResponsePtrOutput) ToDicomServiceAuthenticationConfigurationResponsePtrOutputWithContext(ctx context.Context) DicomServiceAuthenticationConfigurationResponsePtrOutput {
	return o
}

func (o DicomServiceAuthenticationConfigurationResponsePtrOutput) Elem() DicomServiceAuthenticationConfigurationResponseOutput {
	return o.ApplyT(func(v *DicomServiceAuthenticationConfigurationResponse) DicomServiceAuthenticationConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret DicomServiceAuthenticationConfigurationResponse
		return ret
	}).(DicomServiceAuthenticationConfigurationResponseOutput)
}

func (o DicomServiceAuthenticationConfigurationResponsePtrOutput) Audiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DicomServiceAuthenticationConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.Audiences
	}).(pulumi.StringArrayOutput)
}

func (o DicomServiceAuthenticationConfigurationResponsePtrOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DicomServiceAuthenticationConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Authority
	}).(pulumi.StringPtrOutput)
}

type FhirServiceAccessPolicyEntry struct {
	ObjectId string `pulumi:"objectId"`
}





type FhirServiceAccessPolicyEntryInput interface {
	pulumi.Input

	ToFhirServiceAccessPolicyEntryOutput() FhirServiceAccessPolicyEntryOutput
	ToFhirServiceAccessPolicyEntryOutputWithContext(context.Context) FhirServiceAccessPolicyEntryOutput
}

type FhirServiceAccessPolicyEntryArgs struct {
	ObjectId pulumi.StringInput `pulumi:"objectId"`
}

func (FhirServiceAccessPolicyEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FhirServiceAccessPolicyEntry)(nil)).Elem()
}

func (i FhirServiceAccessPolicyEntryArgs) ToFhirServiceAccessPolicyEntryOutput() FhirServiceAccessPolicyEntryOutput {
	return i.ToFhirServiceAccessPolicyEntryOutputWithContext(context.Background())
}

func (i FhirServiceAccessPolicyEntryArgs) ToFhirServiceAccessPolicyEntryOutputWithContext(ctx context.Context) FhirServiceAccessPolicyEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirServiceAccessPolicyEntryOutput)
}





type FhirServiceAccessPolicyEntryArrayInput interface {
	pulumi.Input

	ToFhirServiceAccessPolicyEntryArrayOutput() FhirServiceAccessPolicyEntryArrayOutput
	ToFhirServiceAccessPolicyEntryArrayOutputWithContext(context.Context) FhirServiceAccessPolicyEntryArrayOutput
}

type FhirServiceAccessPolicyEntryArray []FhirServiceAccessPolicyEntryInput

func (FhirServiceAccessPolicyEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FhirServiceAccessPolicyEntry)(nil)).Elem()
}

func (i FhirServiceAccessPolicyEntryArray) ToFhirServiceAccessPolicyEntryArrayOutput() FhirServiceAccessPolicyEntryArrayOutput {
	return i.ToFhirServiceAccessPolicyEntryArrayOutputWithContext(context.Background())
}

func (i FhirServiceAccessPolicyEntryArray) ToFhirServiceAccessPolicyEntryArrayOutputWithContext(ctx context.Context) FhirServiceAccessPolicyEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirServiceAccessPolicyEntryArrayOutput)
}

type FhirServiceAccessPolicyEntryOutput struct{ *pulumi.OutputState }

func (FhirServiceAccessPolicyEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FhirServiceAccessPolicyEntry)(nil)).Elem()
}

func (o FhirServiceAccessPolicyEntryOutput) ToFhirServiceAccessPolicyEntryOutput() FhirServiceAccessPolicyEntryOutput {
	return o
}

func (o FhirServiceAccessPolicyEntryOutput) ToFhirServiceAccessPolicyEntryOutputWithContext(ctx context.Context) FhirServiceAccessPolicyEntryOutput {
	return o
}

func (o FhirServiceAccessPolicyEntryOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v FhirServiceAccessPolicyEntry) string { return v.ObjectId }).(pulumi.StringOutput)
}

type FhirServiceAccessPolicyEntryArrayOutput struct{ *pulumi.OutputState }

func (FhirServiceAccessPolicyEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FhirServiceAccessPolicyEntry)(nil)).Elem()
}

func (o FhirServiceAccessPolicyEntryArrayOutput) ToFhirServiceAccessPolicyEntryArrayOutput() FhirServiceAccessPolicyEntryArrayOutput {
	return o
}

func (o FhirServiceAccessPolicyEntryArrayOutput) ToFhirServiceAccessPolicyEntryArrayOutputWithContext(ctx context.Context) FhirServiceAccessPolicyEntryArrayOutput {
	return o
}

func (o FhirServiceAccessPolicyEntryArrayOutput) Index(i pulumi.IntInput) FhirServiceAccessPolicyEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FhirServiceAccessPolicyEntry {
		return vs[0].([]FhirServiceAccessPolicyEntry)[vs[1].(int)]
	}).(FhirServiceAccessPolicyEntryOutput)
}

type FhirServiceAccessPolicyEntryResponse struct {
	ObjectId string `pulumi:"objectId"`
}

type FhirServiceAccessPolicyEntryResponseOutput struct{ *pulumi.OutputState }

func (FhirServiceAccessPolicyEntryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FhirServiceAccessPolicyEntryResponse)(nil)).Elem()
}

func (o FhirServiceAccessPolicyEntryResponseOutput) ToFhirServiceAccessPolicyEntryResponseOutput() FhirServiceAccessPolicyEntryResponseOutput {
	return o
}

func (o FhirServiceAccessPolicyEntryResponseOutput) ToFhirServiceAccessPolicyEntryResponseOutputWithContext(ctx context.Context) FhirServiceAccessPolicyEntryResponseOutput {
	return o
}

func (o FhirServiceAccessPolicyEntryResponseOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v FhirServiceAccessPolicyEntryResponse) string { return v.ObjectId }).(pulumi.StringOutput)
}

type FhirServiceAccessPolicyEntryResponseArrayOutput struct{ *pulumi.OutputState }

func (FhirServiceAccessPolicyEntryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FhirServiceAccessPolicyEntryResponse)(nil)).Elem()
}

func (o FhirServiceAccessPolicyEntryResponseArrayOutput) ToFhirServiceAccessPolicyEntryResponseArrayOutput() FhirServiceAccessPolicyEntryResponseArrayOutput {
	return o
}

func (o FhirServiceAccessPolicyEntryResponseArrayOutput) ToFhirServiceAccessPolicyEntryResponseArrayOutputWithContext(ctx context.Context) FhirServiceAccessPolicyEntryResponseArrayOutput {
	return o
}

func (o FhirServiceAccessPolicyEntryResponseArrayOutput) Index(i pulumi.IntInput) FhirServiceAccessPolicyEntryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FhirServiceAccessPolicyEntryResponse {
		return vs[0].([]FhirServiceAccessPolicyEntryResponse)[vs[1].(int)]
	}).(FhirServiceAccessPolicyEntryResponseOutput)
}

type FhirServiceAcrConfiguration struct {
	LoginServers []string                  `pulumi:"loginServers"`
	OciArtifacts []ServiceOciArtifactEntry `pulumi:"ociArtifacts"`
}





type FhirServiceAcrConfigurationInput interface {
	pulumi.Input

	ToFhirServiceAcrConfigurationOutput() FhirServiceAcrConfigurationOutput
	ToFhirServiceAcrConfigurationOutputWithContext(context.Context) FhirServiceAcrConfigurationOutput
}

type FhirServiceAcrConfigurationArgs struct {
	LoginServers pulumi.StringArrayInput           `pulumi:"loginServers"`
	OciArtifacts ServiceOciArtifactEntryArrayInput `pulumi:"ociArtifacts"`
}

func (FhirServiceAcrConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FhirServiceAcrConfiguration)(nil)).Elem()
}

func (i FhirServiceAcrConfigurationArgs) ToFhirServiceAcrConfigurationOutput() FhirServiceAcrConfigurationOutput {
	return i.ToFhirServiceAcrConfigurationOutputWithContext(context.Background())
}

func (i FhirServiceAcrConfigurationArgs) ToFhirServiceAcrConfigurationOutputWithContext(ctx context.Context) FhirServiceAcrConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirServiceAcrConfigurationOutput)
}

func (i FhirServiceAcrConfigurationArgs) ToFhirServiceAcrConfigurationPtrOutput() FhirServiceAcrConfigurationPtrOutput {
	return i.ToFhirServiceAcrConfigurationPtrOutputWithContext(context.Background())
}

func (i FhirServiceAcrConfigurationArgs) ToFhirServiceAcrConfigurationPtrOutputWithContext(ctx context.Context) FhirServiceAcrConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirServiceAcrConfigurationOutput).ToFhirServiceAcrConfigurationPtrOutputWithContext(ctx)
}









type FhirServiceAcrConfigurationPtrInput interface {
	pulumi.Input

	ToFhirServiceAcrConfigurationPtrOutput() FhirServiceAcrConfigurationPtrOutput
	ToFhirServiceAcrConfigurationPtrOutputWithContext(context.Context) FhirServiceAcrConfigurationPtrOutput
}

type fhirServiceAcrConfigurationPtrType FhirServiceAcrConfigurationArgs

func FhirServiceAcrConfigurationPtr(v *FhirServiceAcrConfigurationArgs) FhirServiceAcrConfigurationPtrInput {
	return (*fhirServiceAcrConfigurationPtrType)(v)
}

func (*fhirServiceAcrConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirServiceAcrConfiguration)(nil)).Elem()
}

func (i *fhirServiceAcrConfigurationPtrType) ToFhirServiceAcrConfigurationPtrOutput() FhirServiceAcrConfigurationPtrOutput {
	return i.ToFhirServiceAcrConfigurationPtrOutputWithContext(context.Background())
}

func (i *fhirServiceAcrConfigurationPtrType) ToFhirServiceAcrConfigurationPtrOutputWithContext(ctx context.Context) FhirServiceAcrConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirServiceAcrConfigurationPtrOutput)
}

type FhirServiceAcrConfigurationOutput struct{ *pulumi.OutputState }

func (FhirServiceAcrConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FhirServiceAcrConfiguration)(nil)).Elem()
}

func (o FhirServiceAcrConfigurationOutput) ToFhirServiceAcrConfigurationOutput() FhirServiceAcrConfigurationOutput {
	return o
}

func (o FhirServiceAcrConfigurationOutput) ToFhirServiceAcrConfigurationOutputWithContext(ctx context.Context) FhirServiceAcrConfigurationOutput {
	return o
}

func (o FhirServiceAcrConfigurationOutput) ToFhirServiceAcrConfigurationPtrOutput() FhirServiceAcrConfigurationPtrOutput {
	return o.ToFhirServiceAcrConfigurationPtrOutputWithContext(context.Background())
}

func (o FhirServiceAcrConfigurationOutput) ToFhirServiceAcrConfigurationPtrOutputWithContext(ctx context.Context) FhirServiceAcrConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FhirServiceAcrConfiguration) *FhirServiceAcrConfiguration {
		return &v
	}).(FhirServiceAcrConfigurationPtrOutput)
}

func (o FhirServiceAcrConfigurationOutput) LoginServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FhirServiceAcrConfiguration) []string { return v.LoginServers }).(pulumi.StringArrayOutput)
}

func (o FhirServiceAcrConfigurationOutput) OciArtifacts() ServiceOciArtifactEntryArrayOutput {
	return o.ApplyT(func(v FhirServiceAcrConfiguration) []ServiceOciArtifactEntry { return v.OciArtifacts }).(ServiceOciArtifactEntryArrayOutput)
}

type FhirServiceAcrConfigurationPtrOutput struct{ *pulumi.OutputState }

func (FhirServiceAcrConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirServiceAcrConfiguration)(nil)).Elem()
}

func (o FhirServiceAcrConfigurationPtrOutput) ToFhirServiceAcrConfigurationPtrOutput() FhirServiceAcrConfigurationPtrOutput {
	return o
}

func (o FhirServiceAcrConfigurationPtrOutput) ToFhirServiceAcrConfigurationPtrOutputWithContext(ctx context.Context) FhirServiceAcrConfigurationPtrOutput {
	return o
}

func (o FhirServiceAcrConfigurationPtrOutput) Elem() FhirServiceAcrConfigurationOutput {
	return o.ApplyT(func(v *FhirServiceAcrConfiguration) FhirServiceAcrConfiguration {
		if v != nil {
			return *v
		}
		var ret FhirServiceAcrConfiguration
		return ret
	}).(FhirServiceAcrConfigurationOutput)
}

func (o FhirServiceAcrConfigurationPtrOutput) LoginServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FhirServiceAcrConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.LoginServers
	}).(pulumi.StringArrayOutput)
}

func (o FhirServiceAcrConfigurationPtrOutput) OciArtifacts() ServiceOciArtifactEntryArrayOutput {
	return o.ApplyT(func(v *FhirServiceAcrConfiguration) []ServiceOciArtifactEntry {
		if v == nil {
			return nil
		}
		return v.OciArtifacts
	}).(ServiceOciArtifactEntryArrayOutput)
}

type FhirServiceAcrConfigurationResponse struct {
	LoginServers []string                          `pulumi:"loginServers"`
	OciArtifacts []ServiceOciArtifactEntryResponse `pulumi:"ociArtifacts"`
}

type FhirServiceAcrConfigurationResponseOutput struct{ *pulumi.OutputState }

func (FhirServiceAcrConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FhirServiceAcrConfigurationResponse)(nil)).Elem()
}

func (o FhirServiceAcrConfigurationResponseOutput) ToFhirServiceAcrConfigurationResponseOutput() FhirServiceAcrConfigurationResponseOutput {
	return o
}

func (o FhirServiceAcrConfigurationResponseOutput) ToFhirServiceAcrConfigurationResponseOutputWithContext(ctx context.Context) FhirServiceAcrConfigurationResponseOutput {
	return o
}

func (o FhirServiceAcrConfigurationResponseOutput) LoginServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FhirServiceAcrConfigurationResponse) []string { return v.LoginServers }).(pulumi.StringArrayOutput)
}

func (o FhirServiceAcrConfigurationResponseOutput) OciArtifacts() ServiceOciArtifactEntryResponseArrayOutput {
	return o.ApplyT(func(v FhirServiceAcrConfigurationResponse) []ServiceOciArtifactEntryResponse { return v.OciArtifacts }).(ServiceOciArtifactEntryResponseArrayOutput)
}

type FhirServiceAcrConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (FhirServiceAcrConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirServiceAcrConfigurationResponse)(nil)).Elem()
}

func (o FhirServiceAcrConfigurationResponsePtrOutput) ToFhirServiceAcrConfigurationResponsePtrOutput() FhirServiceAcrConfigurationResponsePtrOutput {
	return o
}

func (o FhirServiceAcrConfigurationResponsePtrOutput) ToFhirServiceAcrConfigurationResponsePtrOutputWithContext(ctx context.Context) FhirServiceAcrConfigurationResponsePtrOutput {
	return o
}

func (o FhirServiceAcrConfigurationResponsePtrOutput) Elem() FhirServiceAcrConfigurationResponseOutput {
	return o.ApplyT(func(v *FhirServiceAcrConfigurationResponse) FhirServiceAcrConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret FhirServiceAcrConfigurationResponse
		return ret
	}).(FhirServiceAcrConfigurationResponseOutput)
}

func (o FhirServiceAcrConfigurationResponsePtrOutput) LoginServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FhirServiceAcrConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.LoginServers
	}).(pulumi.StringArrayOutput)
}

func (o FhirServiceAcrConfigurationResponsePtrOutput) OciArtifacts() ServiceOciArtifactEntryResponseArrayOutput {
	return o.ApplyT(func(v *FhirServiceAcrConfigurationResponse) []ServiceOciArtifactEntryResponse {
		if v == nil {
			return nil
		}
		return v.OciArtifacts
	}).(ServiceOciArtifactEntryResponseArrayOutput)
}

type FhirServiceAuthenticationConfiguration struct {
	Audience          *string `pulumi:"audience"`
	Authority         *string `pulumi:"authority"`
	SmartProxyEnabled *bool   `pulumi:"smartProxyEnabled"`
}





type FhirServiceAuthenticationConfigurationInput interface {
	pulumi.Input

	ToFhirServiceAuthenticationConfigurationOutput() FhirServiceAuthenticationConfigurationOutput
	ToFhirServiceAuthenticationConfigurationOutputWithContext(context.Context) FhirServiceAuthenticationConfigurationOutput
}

type FhirServiceAuthenticationConfigurationArgs struct {
	Audience          pulumi.StringPtrInput `pulumi:"audience"`
	Authority         pulumi.StringPtrInput `pulumi:"authority"`
	SmartProxyEnabled pulumi.BoolPtrInput   `pulumi:"smartProxyEnabled"`
}

func (FhirServiceAuthenticationConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FhirServiceAuthenticationConfiguration)(nil)).Elem()
}

func (i FhirServiceAuthenticationConfigurationArgs) ToFhirServiceAuthenticationConfigurationOutput() FhirServiceAuthenticationConfigurationOutput {
	return i.ToFhirServiceAuthenticationConfigurationOutputWithContext(context.Background())
}

func (i FhirServiceAuthenticationConfigurationArgs) ToFhirServiceAuthenticationConfigurationOutputWithContext(ctx context.Context) FhirServiceAuthenticationConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirServiceAuthenticationConfigurationOutput)
}

func (i FhirServiceAuthenticationConfigurationArgs) ToFhirServiceAuthenticationConfigurationPtrOutput() FhirServiceAuthenticationConfigurationPtrOutput {
	return i.ToFhirServiceAuthenticationConfigurationPtrOutputWithContext(context.Background())
}

func (i FhirServiceAuthenticationConfigurationArgs) ToFhirServiceAuthenticationConfigurationPtrOutputWithContext(ctx context.Context) FhirServiceAuthenticationConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirServiceAuthenticationConfigurationOutput).ToFhirServiceAuthenticationConfigurationPtrOutputWithContext(ctx)
}









type FhirServiceAuthenticationConfigurationPtrInput interface {
	pulumi.Input

	ToFhirServiceAuthenticationConfigurationPtrOutput() FhirServiceAuthenticationConfigurationPtrOutput
	ToFhirServiceAuthenticationConfigurationPtrOutputWithContext(context.Context) FhirServiceAuthenticationConfigurationPtrOutput
}

type fhirServiceAuthenticationConfigurationPtrType FhirServiceAuthenticationConfigurationArgs

func FhirServiceAuthenticationConfigurationPtr(v *FhirServiceAuthenticationConfigurationArgs) FhirServiceAuthenticationConfigurationPtrInput {
	return (*fhirServiceAuthenticationConfigurationPtrType)(v)
}

func (*fhirServiceAuthenticationConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirServiceAuthenticationConfiguration)(nil)).Elem()
}

func (i *fhirServiceAuthenticationConfigurationPtrType) ToFhirServiceAuthenticationConfigurationPtrOutput() FhirServiceAuthenticationConfigurationPtrOutput {
	return i.ToFhirServiceAuthenticationConfigurationPtrOutputWithContext(context.Background())
}

func (i *fhirServiceAuthenticationConfigurationPtrType) ToFhirServiceAuthenticationConfigurationPtrOutputWithContext(ctx context.Context) FhirServiceAuthenticationConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirServiceAuthenticationConfigurationPtrOutput)
}

type FhirServiceAuthenticationConfigurationOutput struct{ *pulumi.OutputState }

func (FhirServiceAuthenticationConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FhirServiceAuthenticationConfiguration)(nil)).Elem()
}

func (o FhirServiceAuthenticationConfigurationOutput) ToFhirServiceAuthenticationConfigurationOutput() FhirServiceAuthenticationConfigurationOutput {
	return o
}

func (o FhirServiceAuthenticationConfigurationOutput) ToFhirServiceAuthenticationConfigurationOutputWithContext(ctx context.Context) FhirServiceAuthenticationConfigurationOutput {
	return o
}

func (o FhirServiceAuthenticationConfigurationOutput) ToFhirServiceAuthenticationConfigurationPtrOutput() FhirServiceAuthenticationConfigurationPtrOutput {
	return o.ToFhirServiceAuthenticationConfigurationPtrOutputWithContext(context.Background())
}

func (o FhirServiceAuthenticationConfigurationOutput) ToFhirServiceAuthenticationConfigurationPtrOutputWithContext(ctx context.Context) FhirServiceAuthenticationConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FhirServiceAuthenticationConfiguration) *FhirServiceAuthenticationConfiguration {
		return &v
	}).(FhirServiceAuthenticationConfigurationPtrOutput)
}

func (o FhirServiceAuthenticationConfigurationOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FhirServiceAuthenticationConfiguration) *string { return v.Audience }).(pulumi.StringPtrOutput)
}

func (o FhirServiceAuthenticationConfigurationOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FhirServiceAuthenticationConfiguration) *string { return v.Authority }).(pulumi.StringPtrOutput)
}

func (o FhirServiceAuthenticationConfigurationOutput) SmartProxyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FhirServiceAuthenticationConfiguration) *bool { return v.SmartProxyEnabled }).(pulumi.BoolPtrOutput)
}

type FhirServiceAuthenticationConfigurationPtrOutput struct{ *pulumi.OutputState }

func (FhirServiceAuthenticationConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirServiceAuthenticationConfiguration)(nil)).Elem()
}

func (o FhirServiceAuthenticationConfigurationPtrOutput) ToFhirServiceAuthenticationConfigurationPtrOutput() FhirServiceAuthenticationConfigurationPtrOutput {
	return o
}

func (o FhirServiceAuthenticationConfigurationPtrOutput) ToFhirServiceAuthenticationConfigurationPtrOutputWithContext(ctx context.Context) FhirServiceAuthenticationConfigurationPtrOutput {
	return o
}

func (o FhirServiceAuthenticationConfigurationPtrOutput) Elem() FhirServiceAuthenticationConfigurationOutput {
	return o.ApplyT(func(v *FhirServiceAuthenticationConfiguration) FhirServiceAuthenticationConfiguration {
		if v != nil {
			return *v
		}
		var ret FhirServiceAuthenticationConfiguration
		return ret
	}).(FhirServiceAuthenticationConfigurationOutput)
}

func (o FhirServiceAuthenticationConfigurationPtrOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FhirServiceAuthenticationConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Audience
	}).(pulumi.StringPtrOutput)
}

func (o FhirServiceAuthenticationConfigurationPtrOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FhirServiceAuthenticationConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Authority
	}).(pulumi.StringPtrOutput)
}

func (o FhirServiceAuthenticationConfigurationPtrOutput) SmartProxyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FhirServiceAuthenticationConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.SmartProxyEnabled
	}).(pulumi.BoolPtrOutput)
}

type FhirServiceAuthenticationConfigurationResponse struct {
	Audience          *string `pulumi:"audience"`
	Authority         *string `pulumi:"authority"`
	SmartProxyEnabled *bool   `pulumi:"smartProxyEnabled"`
}

type FhirServiceAuthenticationConfigurationResponseOutput struct{ *pulumi.OutputState }

func (FhirServiceAuthenticationConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FhirServiceAuthenticationConfigurationResponse)(nil)).Elem()
}

func (o FhirServiceAuthenticationConfigurationResponseOutput) ToFhirServiceAuthenticationConfigurationResponseOutput() FhirServiceAuthenticationConfigurationResponseOutput {
	return o
}

func (o FhirServiceAuthenticationConfigurationResponseOutput) ToFhirServiceAuthenticationConfigurationResponseOutputWithContext(ctx context.Context) FhirServiceAuthenticationConfigurationResponseOutput {
	return o
}

func (o FhirServiceAuthenticationConfigurationResponseOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FhirServiceAuthenticationConfigurationResponse) *string { return v.Audience }).(pulumi.StringPtrOutput)
}

func (o FhirServiceAuthenticationConfigurationResponseOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FhirServiceAuthenticationConfigurationResponse) *string { return v.Authority }).(pulumi.StringPtrOutput)
}

func (o FhirServiceAuthenticationConfigurationResponseOutput) SmartProxyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FhirServiceAuthenticationConfigurationResponse) *bool { return v.SmartProxyEnabled }).(pulumi.BoolPtrOutput)
}

type FhirServiceAuthenticationConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (FhirServiceAuthenticationConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirServiceAuthenticationConfigurationResponse)(nil)).Elem()
}

func (o FhirServiceAuthenticationConfigurationResponsePtrOutput) ToFhirServiceAuthenticationConfigurationResponsePtrOutput() FhirServiceAuthenticationConfigurationResponsePtrOutput {
	return o
}

func (o FhirServiceAuthenticationConfigurationResponsePtrOutput) ToFhirServiceAuthenticationConfigurationResponsePtrOutputWithContext(ctx context.Context) FhirServiceAuthenticationConfigurationResponsePtrOutput {
	return o
}

func (o FhirServiceAuthenticationConfigurationResponsePtrOutput) Elem() FhirServiceAuthenticationConfigurationResponseOutput {
	return o.ApplyT(func(v *FhirServiceAuthenticationConfigurationResponse) FhirServiceAuthenticationConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret FhirServiceAuthenticationConfigurationResponse
		return ret
	}).(FhirServiceAuthenticationConfigurationResponseOutput)
}

func (o FhirServiceAuthenticationConfigurationResponsePtrOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FhirServiceAuthenticationConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Audience
	}).(pulumi.StringPtrOutput)
}

func (o FhirServiceAuthenticationConfigurationResponsePtrOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FhirServiceAuthenticationConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Authority
	}).(pulumi.StringPtrOutput)
}

func (o FhirServiceAuthenticationConfigurationResponsePtrOutput) SmartProxyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FhirServiceAuthenticationConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SmartProxyEnabled
	}).(pulumi.BoolPtrOutput)
}

type FhirServiceCorsConfiguration struct {
	AllowCredentials *bool    `pulumi:"allowCredentials"`
	Headers          []string `pulumi:"headers"`
	MaxAge           *int     `pulumi:"maxAge"`
	Methods          []string `pulumi:"methods"`
	Origins          []string `pulumi:"origins"`
}





type FhirServiceCorsConfigurationInput interface {
	pulumi.Input

	ToFhirServiceCorsConfigurationOutput() FhirServiceCorsConfigurationOutput
	ToFhirServiceCorsConfigurationOutputWithContext(context.Context) FhirServiceCorsConfigurationOutput
}

type FhirServiceCorsConfigurationArgs struct {
	AllowCredentials pulumi.BoolPtrInput     `pulumi:"allowCredentials"`
	Headers          pulumi.StringArrayInput `pulumi:"headers"`
	MaxAge           pulumi.IntPtrInput      `pulumi:"maxAge"`
	Methods          pulumi.StringArrayInput `pulumi:"methods"`
	Origins          pulumi.StringArrayInput `pulumi:"origins"`
}

func (FhirServiceCorsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FhirServiceCorsConfiguration)(nil)).Elem()
}

func (i FhirServiceCorsConfigurationArgs) ToFhirServiceCorsConfigurationOutput() FhirServiceCorsConfigurationOutput {
	return i.ToFhirServiceCorsConfigurationOutputWithContext(context.Background())
}

func (i FhirServiceCorsConfigurationArgs) ToFhirServiceCorsConfigurationOutputWithContext(ctx context.Context) FhirServiceCorsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirServiceCorsConfigurationOutput)
}

func (i FhirServiceCorsConfigurationArgs) ToFhirServiceCorsConfigurationPtrOutput() FhirServiceCorsConfigurationPtrOutput {
	return i.ToFhirServiceCorsConfigurationPtrOutputWithContext(context.Background())
}

func (i FhirServiceCorsConfigurationArgs) ToFhirServiceCorsConfigurationPtrOutputWithContext(ctx context.Context) FhirServiceCorsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirServiceCorsConfigurationOutput).ToFhirServiceCorsConfigurationPtrOutputWithContext(ctx)
}









type FhirServiceCorsConfigurationPtrInput interface {
	pulumi.Input

	ToFhirServiceCorsConfigurationPtrOutput() FhirServiceCorsConfigurationPtrOutput
	ToFhirServiceCorsConfigurationPtrOutputWithContext(context.Context) FhirServiceCorsConfigurationPtrOutput
}

type fhirServiceCorsConfigurationPtrType FhirServiceCorsConfigurationArgs

func FhirServiceCorsConfigurationPtr(v *FhirServiceCorsConfigurationArgs) FhirServiceCorsConfigurationPtrInput {
	return (*fhirServiceCorsConfigurationPtrType)(v)
}

func (*fhirServiceCorsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirServiceCorsConfiguration)(nil)).Elem()
}

func (i *fhirServiceCorsConfigurationPtrType) ToFhirServiceCorsConfigurationPtrOutput() FhirServiceCorsConfigurationPtrOutput {
	return i.ToFhirServiceCorsConfigurationPtrOutputWithContext(context.Background())
}

func (i *fhirServiceCorsConfigurationPtrType) ToFhirServiceCorsConfigurationPtrOutputWithContext(ctx context.Context) FhirServiceCorsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirServiceCorsConfigurationPtrOutput)
}

type FhirServiceCorsConfigurationOutput struct{ *pulumi.OutputState }

func (FhirServiceCorsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FhirServiceCorsConfiguration)(nil)).Elem()
}

func (o FhirServiceCorsConfigurationOutput) ToFhirServiceCorsConfigurationOutput() FhirServiceCorsConfigurationOutput {
	return o
}

func (o FhirServiceCorsConfigurationOutput) ToFhirServiceCorsConfigurationOutputWithContext(ctx context.Context) FhirServiceCorsConfigurationOutput {
	return o
}

func (o FhirServiceCorsConfigurationOutput) ToFhirServiceCorsConfigurationPtrOutput() FhirServiceCorsConfigurationPtrOutput {
	return o.ToFhirServiceCorsConfigurationPtrOutputWithContext(context.Background())
}

func (o FhirServiceCorsConfigurationOutput) ToFhirServiceCorsConfigurationPtrOutputWithContext(ctx context.Context) FhirServiceCorsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FhirServiceCorsConfiguration) *FhirServiceCorsConfiguration {
		return &v
	}).(FhirServiceCorsConfigurationPtrOutput)
}

func (o FhirServiceCorsConfigurationOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FhirServiceCorsConfiguration) *bool { return v.AllowCredentials }).(pulumi.BoolPtrOutput)
}

func (o FhirServiceCorsConfigurationOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FhirServiceCorsConfiguration) []string { return v.Headers }).(pulumi.StringArrayOutput)
}

func (o FhirServiceCorsConfigurationOutput) MaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FhirServiceCorsConfiguration) *int { return v.MaxAge }).(pulumi.IntPtrOutput)
}

func (o FhirServiceCorsConfigurationOutput) Methods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FhirServiceCorsConfiguration) []string { return v.Methods }).(pulumi.StringArrayOutput)
}

func (o FhirServiceCorsConfigurationOutput) Origins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FhirServiceCorsConfiguration) []string { return v.Origins }).(pulumi.StringArrayOutput)
}

type FhirServiceCorsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (FhirServiceCorsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirServiceCorsConfiguration)(nil)).Elem()
}

func (o FhirServiceCorsConfigurationPtrOutput) ToFhirServiceCorsConfigurationPtrOutput() FhirServiceCorsConfigurationPtrOutput {
	return o
}

func (o FhirServiceCorsConfigurationPtrOutput) ToFhirServiceCorsConfigurationPtrOutputWithContext(ctx context.Context) FhirServiceCorsConfigurationPtrOutput {
	return o
}

func (o FhirServiceCorsConfigurationPtrOutput) Elem() FhirServiceCorsConfigurationOutput {
	return o.ApplyT(func(v *FhirServiceCorsConfiguration) FhirServiceCorsConfiguration {
		if v != nil {
			return *v
		}
		var ret FhirServiceCorsConfiguration
		return ret
	}).(FhirServiceCorsConfigurationOutput)
}

func (o FhirServiceCorsConfigurationPtrOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FhirServiceCorsConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.AllowCredentials
	}).(pulumi.BoolPtrOutput)
}

func (o FhirServiceCorsConfigurationPtrOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FhirServiceCorsConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(pulumi.StringArrayOutput)
}

func (o FhirServiceCorsConfigurationPtrOutput) MaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FhirServiceCorsConfiguration) *int {
		if v == nil {
			return nil
		}
		return v.MaxAge
	}).(pulumi.IntPtrOutput)
}

func (o FhirServiceCorsConfigurationPtrOutput) Methods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FhirServiceCorsConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.Methods
	}).(pulumi.StringArrayOutput)
}

func (o FhirServiceCorsConfigurationPtrOutput) Origins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FhirServiceCorsConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.Origins
	}).(pulumi.StringArrayOutput)
}

type FhirServiceCorsConfigurationResponse struct {
	AllowCredentials *bool    `pulumi:"allowCredentials"`
	Headers          []string `pulumi:"headers"`
	MaxAge           *int     `pulumi:"maxAge"`
	Methods          []string `pulumi:"methods"`
	Origins          []string `pulumi:"origins"`
}

type FhirServiceCorsConfigurationResponseOutput struct{ *pulumi.OutputState }

func (FhirServiceCorsConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FhirServiceCorsConfigurationResponse)(nil)).Elem()
}

func (o FhirServiceCorsConfigurationResponseOutput) ToFhirServiceCorsConfigurationResponseOutput() FhirServiceCorsConfigurationResponseOutput {
	return o
}

func (o FhirServiceCorsConfigurationResponseOutput) ToFhirServiceCorsConfigurationResponseOutputWithContext(ctx context.Context) FhirServiceCorsConfigurationResponseOutput {
	return o
}

func (o FhirServiceCorsConfigurationResponseOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FhirServiceCorsConfigurationResponse) *bool { return v.AllowCredentials }).(pulumi.BoolPtrOutput)
}

func (o FhirServiceCorsConfigurationResponseOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FhirServiceCorsConfigurationResponse) []string { return v.Headers }).(pulumi.StringArrayOutput)
}

func (o FhirServiceCorsConfigurationResponseOutput) MaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FhirServiceCorsConfigurationResponse) *int { return v.MaxAge }).(pulumi.IntPtrOutput)
}

func (o FhirServiceCorsConfigurationResponseOutput) Methods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FhirServiceCorsConfigurationResponse) []string { return v.Methods }).(pulumi.StringArrayOutput)
}

func (o FhirServiceCorsConfigurationResponseOutput) Origins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FhirServiceCorsConfigurationResponse) []string { return v.Origins }).(pulumi.StringArrayOutput)
}

type FhirServiceCorsConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (FhirServiceCorsConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirServiceCorsConfigurationResponse)(nil)).Elem()
}

func (o FhirServiceCorsConfigurationResponsePtrOutput) ToFhirServiceCorsConfigurationResponsePtrOutput() FhirServiceCorsConfigurationResponsePtrOutput {
	return o
}

func (o FhirServiceCorsConfigurationResponsePtrOutput) ToFhirServiceCorsConfigurationResponsePtrOutputWithContext(ctx context.Context) FhirServiceCorsConfigurationResponsePtrOutput {
	return o
}

func (o FhirServiceCorsConfigurationResponsePtrOutput) Elem() FhirServiceCorsConfigurationResponseOutput {
	return o.ApplyT(func(v *FhirServiceCorsConfigurationResponse) FhirServiceCorsConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret FhirServiceCorsConfigurationResponse
		return ret
	}).(FhirServiceCorsConfigurationResponseOutput)
}

func (o FhirServiceCorsConfigurationResponsePtrOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FhirServiceCorsConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AllowCredentials
	}).(pulumi.BoolPtrOutput)
}

func (o FhirServiceCorsConfigurationResponsePtrOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FhirServiceCorsConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(pulumi.StringArrayOutput)
}

func (o FhirServiceCorsConfigurationResponsePtrOutput) MaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FhirServiceCorsConfigurationResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxAge
	}).(pulumi.IntPtrOutput)
}

func (o FhirServiceCorsConfigurationResponsePtrOutput) Methods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FhirServiceCorsConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.Methods
	}).(pulumi.StringArrayOutput)
}

func (o FhirServiceCorsConfigurationResponsePtrOutput) Origins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FhirServiceCorsConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.Origins
	}).(pulumi.StringArrayOutput)
}

type FhirServiceExportConfiguration struct {
	StorageAccountName *string `pulumi:"storageAccountName"`
}





type FhirServiceExportConfigurationInput interface {
	pulumi.Input

	ToFhirServiceExportConfigurationOutput() FhirServiceExportConfigurationOutput
	ToFhirServiceExportConfigurationOutputWithContext(context.Context) FhirServiceExportConfigurationOutput
}

type FhirServiceExportConfigurationArgs struct {
	StorageAccountName pulumi.StringPtrInput `pulumi:"storageAccountName"`
}

func (FhirServiceExportConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FhirServiceExportConfiguration)(nil)).Elem()
}

func (i FhirServiceExportConfigurationArgs) ToFhirServiceExportConfigurationOutput() FhirServiceExportConfigurationOutput {
	return i.ToFhirServiceExportConfigurationOutputWithContext(context.Background())
}

func (i FhirServiceExportConfigurationArgs) ToFhirServiceExportConfigurationOutputWithContext(ctx context.Context) FhirServiceExportConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirServiceExportConfigurationOutput)
}

func (i FhirServiceExportConfigurationArgs) ToFhirServiceExportConfigurationPtrOutput() FhirServiceExportConfigurationPtrOutput {
	return i.ToFhirServiceExportConfigurationPtrOutputWithContext(context.Background())
}

func (i FhirServiceExportConfigurationArgs) ToFhirServiceExportConfigurationPtrOutputWithContext(ctx context.Context) FhirServiceExportConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirServiceExportConfigurationOutput).ToFhirServiceExportConfigurationPtrOutputWithContext(ctx)
}









type FhirServiceExportConfigurationPtrInput interface {
	pulumi.Input

	ToFhirServiceExportConfigurationPtrOutput() FhirServiceExportConfigurationPtrOutput
	ToFhirServiceExportConfigurationPtrOutputWithContext(context.Context) FhirServiceExportConfigurationPtrOutput
}

type fhirServiceExportConfigurationPtrType FhirServiceExportConfigurationArgs

func FhirServiceExportConfigurationPtr(v *FhirServiceExportConfigurationArgs) FhirServiceExportConfigurationPtrInput {
	return (*fhirServiceExportConfigurationPtrType)(v)
}

func (*fhirServiceExportConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirServiceExportConfiguration)(nil)).Elem()
}

func (i *fhirServiceExportConfigurationPtrType) ToFhirServiceExportConfigurationPtrOutput() FhirServiceExportConfigurationPtrOutput {
	return i.ToFhirServiceExportConfigurationPtrOutputWithContext(context.Background())
}

func (i *fhirServiceExportConfigurationPtrType) ToFhirServiceExportConfigurationPtrOutputWithContext(ctx context.Context) FhirServiceExportConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirServiceExportConfigurationPtrOutput)
}

type FhirServiceExportConfigurationOutput struct{ *pulumi.OutputState }

func (FhirServiceExportConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FhirServiceExportConfiguration)(nil)).Elem()
}

func (o FhirServiceExportConfigurationOutput) ToFhirServiceExportConfigurationOutput() FhirServiceExportConfigurationOutput {
	return o
}

func (o FhirServiceExportConfigurationOutput) ToFhirServiceExportConfigurationOutputWithContext(ctx context.Context) FhirServiceExportConfigurationOutput {
	return o
}

func (o FhirServiceExportConfigurationOutput) ToFhirServiceExportConfigurationPtrOutput() FhirServiceExportConfigurationPtrOutput {
	return o.ToFhirServiceExportConfigurationPtrOutputWithContext(context.Background())
}

func (o FhirServiceExportConfigurationOutput) ToFhirServiceExportConfigurationPtrOutputWithContext(ctx context.Context) FhirServiceExportConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FhirServiceExportConfiguration) *FhirServiceExportConfiguration {
		return &v
	}).(FhirServiceExportConfigurationPtrOutput)
}

func (o FhirServiceExportConfigurationOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FhirServiceExportConfiguration) *string { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

type FhirServiceExportConfigurationPtrOutput struct{ *pulumi.OutputState }

func (FhirServiceExportConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirServiceExportConfiguration)(nil)).Elem()
}

func (o FhirServiceExportConfigurationPtrOutput) ToFhirServiceExportConfigurationPtrOutput() FhirServiceExportConfigurationPtrOutput {
	return o
}

func (o FhirServiceExportConfigurationPtrOutput) ToFhirServiceExportConfigurationPtrOutputWithContext(ctx context.Context) FhirServiceExportConfigurationPtrOutput {
	return o
}

func (o FhirServiceExportConfigurationPtrOutput) Elem() FhirServiceExportConfigurationOutput {
	return o.ApplyT(func(v *FhirServiceExportConfiguration) FhirServiceExportConfiguration {
		if v != nil {
			return *v
		}
		var ret FhirServiceExportConfiguration
		return ret
	}).(FhirServiceExportConfigurationOutput)
}

func (o FhirServiceExportConfigurationPtrOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FhirServiceExportConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountName
	}).(pulumi.StringPtrOutput)
}

type FhirServiceExportConfigurationResponse struct {
	StorageAccountName *string `pulumi:"storageAccountName"`
}

type FhirServiceExportConfigurationResponseOutput struct{ *pulumi.OutputState }

func (FhirServiceExportConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FhirServiceExportConfigurationResponse)(nil)).Elem()
}

func (o FhirServiceExportConfigurationResponseOutput) ToFhirServiceExportConfigurationResponseOutput() FhirServiceExportConfigurationResponseOutput {
	return o
}

func (o FhirServiceExportConfigurationResponseOutput) ToFhirServiceExportConfigurationResponseOutputWithContext(ctx context.Context) FhirServiceExportConfigurationResponseOutput {
	return o
}

func (o FhirServiceExportConfigurationResponseOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FhirServiceExportConfigurationResponse) *string { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

type FhirServiceExportConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (FhirServiceExportConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirServiceExportConfigurationResponse)(nil)).Elem()
}

func (o FhirServiceExportConfigurationResponsePtrOutput) ToFhirServiceExportConfigurationResponsePtrOutput() FhirServiceExportConfigurationResponsePtrOutput {
	return o
}

func (o FhirServiceExportConfigurationResponsePtrOutput) ToFhirServiceExportConfigurationResponsePtrOutputWithContext(ctx context.Context) FhirServiceExportConfigurationResponsePtrOutput {
	return o
}

func (o FhirServiceExportConfigurationResponsePtrOutput) Elem() FhirServiceExportConfigurationResponseOutput {
	return o.ApplyT(func(v *FhirServiceExportConfigurationResponse) FhirServiceExportConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret FhirServiceExportConfigurationResponse
		return ret
	}).(FhirServiceExportConfigurationResponseOutput)
}

func (o FhirServiceExportConfigurationResponsePtrOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FhirServiceExportConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountName
	}).(pulumi.StringPtrOutput)
}

type IotEventHubIngestionEndpointConfiguration struct {
	ConsumerGroup                   *string `pulumi:"consumerGroup"`
	EventHubName                    *string `pulumi:"eventHubName"`
	FullyQualifiedEventHubNamespace *string `pulumi:"fullyQualifiedEventHubNamespace"`
}





type IotEventHubIngestionEndpointConfigurationInput interface {
	pulumi.Input

	ToIotEventHubIngestionEndpointConfigurationOutput() IotEventHubIngestionEndpointConfigurationOutput
	ToIotEventHubIngestionEndpointConfigurationOutputWithContext(context.Context) IotEventHubIngestionEndpointConfigurationOutput
}

type IotEventHubIngestionEndpointConfigurationArgs struct {
	ConsumerGroup                   pulumi.StringPtrInput `pulumi:"consumerGroup"`
	EventHubName                    pulumi.StringPtrInput `pulumi:"eventHubName"`
	FullyQualifiedEventHubNamespace pulumi.StringPtrInput `pulumi:"fullyQualifiedEventHubNamespace"`
}

func (IotEventHubIngestionEndpointConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotEventHubIngestionEndpointConfiguration)(nil)).Elem()
}

func (i IotEventHubIngestionEndpointConfigurationArgs) ToIotEventHubIngestionEndpointConfigurationOutput() IotEventHubIngestionEndpointConfigurationOutput {
	return i.ToIotEventHubIngestionEndpointConfigurationOutputWithContext(context.Background())
}

func (i IotEventHubIngestionEndpointConfigurationArgs) ToIotEventHubIngestionEndpointConfigurationOutputWithContext(ctx context.Context) IotEventHubIngestionEndpointConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotEventHubIngestionEndpointConfigurationOutput)
}

func (i IotEventHubIngestionEndpointConfigurationArgs) ToIotEventHubIngestionEndpointConfigurationPtrOutput() IotEventHubIngestionEndpointConfigurationPtrOutput {
	return i.ToIotEventHubIngestionEndpointConfigurationPtrOutputWithContext(context.Background())
}

func (i IotEventHubIngestionEndpointConfigurationArgs) ToIotEventHubIngestionEndpointConfigurationPtrOutputWithContext(ctx context.Context) IotEventHubIngestionEndpointConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotEventHubIngestionEndpointConfigurationOutput).ToIotEventHubIngestionEndpointConfigurationPtrOutputWithContext(ctx)
}









type IotEventHubIngestionEndpointConfigurationPtrInput interface {
	pulumi.Input

	ToIotEventHubIngestionEndpointConfigurationPtrOutput() IotEventHubIngestionEndpointConfigurationPtrOutput
	ToIotEventHubIngestionEndpointConfigurationPtrOutputWithContext(context.Context) IotEventHubIngestionEndpointConfigurationPtrOutput
}

type iotEventHubIngestionEndpointConfigurationPtrType IotEventHubIngestionEndpointConfigurationArgs

func IotEventHubIngestionEndpointConfigurationPtr(v *IotEventHubIngestionEndpointConfigurationArgs) IotEventHubIngestionEndpointConfigurationPtrInput {
	return (*iotEventHubIngestionEndpointConfigurationPtrType)(v)
}

func (*iotEventHubIngestionEndpointConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IotEventHubIngestionEndpointConfiguration)(nil)).Elem()
}

func (i *iotEventHubIngestionEndpointConfigurationPtrType) ToIotEventHubIngestionEndpointConfigurationPtrOutput() IotEventHubIngestionEndpointConfigurationPtrOutput {
	return i.ToIotEventHubIngestionEndpointConfigurationPtrOutputWithContext(context.Background())
}

func (i *iotEventHubIngestionEndpointConfigurationPtrType) ToIotEventHubIngestionEndpointConfigurationPtrOutputWithContext(ctx context.Context) IotEventHubIngestionEndpointConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotEventHubIngestionEndpointConfigurationPtrOutput)
}

type IotEventHubIngestionEndpointConfigurationOutput struct{ *pulumi.OutputState }

func (IotEventHubIngestionEndpointConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotEventHubIngestionEndpointConfiguration)(nil)).Elem()
}

func (o IotEventHubIngestionEndpointConfigurationOutput) ToIotEventHubIngestionEndpointConfigurationOutput() IotEventHubIngestionEndpointConfigurationOutput {
	return o
}

func (o IotEventHubIngestionEndpointConfigurationOutput) ToIotEventHubIngestionEndpointConfigurationOutputWithContext(ctx context.Context) IotEventHubIngestionEndpointConfigurationOutput {
	return o
}

func (o IotEventHubIngestionEndpointConfigurationOutput) ToIotEventHubIngestionEndpointConfigurationPtrOutput() IotEventHubIngestionEndpointConfigurationPtrOutput {
	return o.ToIotEventHubIngestionEndpointConfigurationPtrOutputWithContext(context.Background())
}

func (o IotEventHubIngestionEndpointConfigurationOutput) ToIotEventHubIngestionEndpointConfigurationPtrOutputWithContext(ctx context.Context) IotEventHubIngestionEndpointConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IotEventHubIngestionEndpointConfiguration) *IotEventHubIngestionEndpointConfiguration {
		return &v
	}).(IotEventHubIngestionEndpointConfigurationPtrOutput)
}

func (o IotEventHubIngestionEndpointConfigurationOutput) ConsumerGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotEventHubIngestionEndpointConfiguration) *string { return v.ConsumerGroup }).(pulumi.StringPtrOutput)
}

func (o IotEventHubIngestionEndpointConfigurationOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotEventHubIngestionEndpointConfiguration) *string { return v.EventHubName }).(pulumi.StringPtrOutput)
}

func (o IotEventHubIngestionEndpointConfigurationOutput) FullyQualifiedEventHubNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotEventHubIngestionEndpointConfiguration) *string { return v.FullyQualifiedEventHubNamespace }).(pulumi.StringPtrOutput)
}

type IotEventHubIngestionEndpointConfigurationPtrOutput struct{ *pulumi.OutputState }

func (IotEventHubIngestionEndpointConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotEventHubIngestionEndpointConfiguration)(nil)).Elem()
}

func (o IotEventHubIngestionEndpointConfigurationPtrOutput) ToIotEventHubIngestionEndpointConfigurationPtrOutput() IotEventHubIngestionEndpointConfigurationPtrOutput {
	return o
}

func (o IotEventHubIngestionEndpointConfigurationPtrOutput) ToIotEventHubIngestionEndpointConfigurationPtrOutputWithContext(ctx context.Context) IotEventHubIngestionEndpointConfigurationPtrOutput {
	return o
}

func (o IotEventHubIngestionEndpointConfigurationPtrOutput) Elem() IotEventHubIngestionEndpointConfigurationOutput {
	return o.ApplyT(func(v *IotEventHubIngestionEndpointConfiguration) IotEventHubIngestionEndpointConfiguration {
		if v != nil {
			return *v
		}
		var ret IotEventHubIngestionEndpointConfiguration
		return ret
	}).(IotEventHubIngestionEndpointConfigurationOutput)
}

func (o IotEventHubIngestionEndpointConfigurationPtrOutput) ConsumerGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotEventHubIngestionEndpointConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.ConsumerGroup
	}).(pulumi.StringPtrOutput)
}

func (o IotEventHubIngestionEndpointConfigurationPtrOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotEventHubIngestionEndpointConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.EventHubName
	}).(pulumi.StringPtrOutput)
}

func (o IotEventHubIngestionEndpointConfigurationPtrOutput) FullyQualifiedEventHubNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotEventHubIngestionEndpointConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.FullyQualifiedEventHubNamespace
	}).(pulumi.StringPtrOutput)
}

type IotEventHubIngestionEndpointConfigurationResponse struct {
	ConsumerGroup                   *string `pulumi:"consumerGroup"`
	EventHubName                    *string `pulumi:"eventHubName"`
	FullyQualifiedEventHubNamespace *string `pulumi:"fullyQualifiedEventHubNamespace"`
}

type IotEventHubIngestionEndpointConfigurationResponseOutput struct{ *pulumi.OutputState }

func (IotEventHubIngestionEndpointConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotEventHubIngestionEndpointConfigurationResponse)(nil)).Elem()
}

func (o IotEventHubIngestionEndpointConfigurationResponseOutput) ToIotEventHubIngestionEndpointConfigurationResponseOutput() IotEventHubIngestionEndpointConfigurationResponseOutput {
	return o
}

func (o IotEventHubIngestionEndpointConfigurationResponseOutput) ToIotEventHubIngestionEndpointConfigurationResponseOutputWithContext(ctx context.Context) IotEventHubIngestionEndpointConfigurationResponseOutput {
	return o
}

func (o IotEventHubIngestionEndpointConfigurationResponseOutput) ConsumerGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotEventHubIngestionEndpointConfigurationResponse) *string { return v.ConsumerGroup }).(pulumi.StringPtrOutput)
}

func (o IotEventHubIngestionEndpointConfigurationResponseOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotEventHubIngestionEndpointConfigurationResponse) *string { return v.EventHubName }).(pulumi.StringPtrOutput)
}

func (o IotEventHubIngestionEndpointConfigurationResponseOutput) FullyQualifiedEventHubNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotEventHubIngestionEndpointConfigurationResponse) *string {
		return v.FullyQualifiedEventHubNamespace
	}).(pulumi.StringPtrOutput)
}

type IotEventHubIngestionEndpointConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (IotEventHubIngestionEndpointConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotEventHubIngestionEndpointConfigurationResponse)(nil)).Elem()
}

func (o IotEventHubIngestionEndpointConfigurationResponsePtrOutput) ToIotEventHubIngestionEndpointConfigurationResponsePtrOutput() IotEventHubIngestionEndpointConfigurationResponsePtrOutput {
	return o
}

func (o IotEventHubIngestionEndpointConfigurationResponsePtrOutput) ToIotEventHubIngestionEndpointConfigurationResponsePtrOutputWithContext(ctx context.Context) IotEventHubIngestionEndpointConfigurationResponsePtrOutput {
	return o
}

func (o IotEventHubIngestionEndpointConfigurationResponsePtrOutput) Elem() IotEventHubIngestionEndpointConfigurationResponseOutput {
	return o.ApplyT(func(v *IotEventHubIngestionEndpointConfigurationResponse) IotEventHubIngestionEndpointConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret IotEventHubIngestionEndpointConfigurationResponse
		return ret
	}).(IotEventHubIngestionEndpointConfigurationResponseOutput)
}

func (o IotEventHubIngestionEndpointConfigurationResponsePtrOutput) ConsumerGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotEventHubIngestionEndpointConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConsumerGroup
	}).(pulumi.StringPtrOutput)
}

func (o IotEventHubIngestionEndpointConfigurationResponsePtrOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotEventHubIngestionEndpointConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.EventHubName
	}).(pulumi.StringPtrOutput)
}

func (o IotEventHubIngestionEndpointConfigurationResponsePtrOutput) FullyQualifiedEventHubNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotEventHubIngestionEndpointConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.FullyQualifiedEventHubNamespace
	}).(pulumi.StringPtrOutput)
}

type IotMappingProperties struct {
	Content interface{} `pulumi:"content"`
}





type IotMappingPropertiesInput interface {
	pulumi.Input

	ToIotMappingPropertiesOutput() IotMappingPropertiesOutput
	ToIotMappingPropertiesOutputWithContext(context.Context) IotMappingPropertiesOutput
}

type IotMappingPropertiesArgs struct {
	Content pulumi.Input `pulumi:"content"`
}

func (IotMappingPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotMappingProperties)(nil)).Elem()
}

func (i IotMappingPropertiesArgs) ToIotMappingPropertiesOutput() IotMappingPropertiesOutput {
	return i.ToIotMappingPropertiesOutputWithContext(context.Background())
}

func (i IotMappingPropertiesArgs) ToIotMappingPropertiesOutputWithContext(ctx context.Context) IotMappingPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotMappingPropertiesOutput)
}

func (i IotMappingPropertiesArgs) ToIotMappingPropertiesPtrOutput() IotMappingPropertiesPtrOutput {
	return i.ToIotMappingPropertiesPtrOutputWithContext(context.Background())
}

func (i IotMappingPropertiesArgs) ToIotMappingPropertiesPtrOutputWithContext(ctx context.Context) IotMappingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotMappingPropertiesOutput).ToIotMappingPropertiesPtrOutputWithContext(ctx)
}









type IotMappingPropertiesPtrInput interface {
	pulumi.Input

	ToIotMappingPropertiesPtrOutput() IotMappingPropertiesPtrOutput
	ToIotMappingPropertiesPtrOutputWithContext(context.Context) IotMappingPropertiesPtrOutput
}

type iotMappingPropertiesPtrType IotMappingPropertiesArgs

func IotMappingPropertiesPtr(v *IotMappingPropertiesArgs) IotMappingPropertiesPtrInput {
	return (*iotMappingPropertiesPtrType)(v)
}

func (*iotMappingPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IotMappingProperties)(nil)).Elem()
}

func (i *iotMappingPropertiesPtrType) ToIotMappingPropertiesPtrOutput() IotMappingPropertiesPtrOutput {
	return i.ToIotMappingPropertiesPtrOutputWithContext(context.Background())
}

func (i *iotMappingPropertiesPtrType) ToIotMappingPropertiesPtrOutputWithContext(ctx context.Context) IotMappingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotMappingPropertiesPtrOutput)
}

type IotMappingPropertiesOutput struct{ *pulumi.OutputState }

func (IotMappingPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotMappingProperties)(nil)).Elem()
}

func (o IotMappingPropertiesOutput) ToIotMappingPropertiesOutput() IotMappingPropertiesOutput {
	return o
}

func (o IotMappingPropertiesOutput) ToIotMappingPropertiesOutputWithContext(ctx context.Context) IotMappingPropertiesOutput {
	return o
}

func (o IotMappingPropertiesOutput) ToIotMappingPropertiesPtrOutput() IotMappingPropertiesPtrOutput {
	return o.ToIotMappingPropertiesPtrOutputWithContext(context.Background())
}

func (o IotMappingPropertiesOutput) ToIotMappingPropertiesPtrOutputWithContext(ctx context.Context) IotMappingPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IotMappingProperties) *IotMappingProperties {
		return &v
	}).(IotMappingPropertiesPtrOutput)
}

func (o IotMappingPropertiesOutput) Content() pulumi.AnyOutput {
	return o.ApplyT(func(v IotMappingProperties) interface{} { return v.Content }).(pulumi.AnyOutput)
}

type IotMappingPropertiesPtrOutput struct{ *pulumi.OutputState }

func (IotMappingPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotMappingProperties)(nil)).Elem()
}

func (o IotMappingPropertiesPtrOutput) ToIotMappingPropertiesPtrOutput() IotMappingPropertiesPtrOutput {
	return o
}

func (o IotMappingPropertiesPtrOutput) ToIotMappingPropertiesPtrOutputWithContext(ctx context.Context) IotMappingPropertiesPtrOutput {
	return o
}

func (o IotMappingPropertiesPtrOutput) Elem() IotMappingPropertiesOutput {
	return o.ApplyT(func(v *IotMappingProperties) IotMappingProperties {
		if v != nil {
			return *v
		}
		var ret IotMappingProperties
		return ret
	}).(IotMappingPropertiesOutput)
}

func (o IotMappingPropertiesPtrOutput) Content() pulumi.AnyOutput {
	return o.ApplyT(func(v *IotMappingProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.Content
	}).(pulumi.AnyOutput)
}

type IotMappingPropertiesResponse struct {
	Content interface{} `pulumi:"content"`
}

type IotMappingPropertiesResponseOutput struct{ *pulumi.OutputState }

func (IotMappingPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotMappingPropertiesResponse)(nil)).Elem()
}

func (o IotMappingPropertiesResponseOutput) ToIotMappingPropertiesResponseOutput() IotMappingPropertiesResponseOutput {
	return o
}

func (o IotMappingPropertiesResponseOutput) ToIotMappingPropertiesResponseOutputWithContext(ctx context.Context) IotMappingPropertiesResponseOutput {
	return o
}

func (o IotMappingPropertiesResponseOutput) Content() pulumi.AnyOutput {
	return o.ApplyT(func(v IotMappingPropertiesResponse) interface{} { return v.Content }).(pulumi.AnyOutput)
}

type IotMappingPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (IotMappingPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotMappingPropertiesResponse)(nil)).Elem()
}

func (o IotMappingPropertiesResponsePtrOutput) ToIotMappingPropertiesResponsePtrOutput() IotMappingPropertiesResponsePtrOutput {
	return o
}

func (o IotMappingPropertiesResponsePtrOutput) ToIotMappingPropertiesResponsePtrOutputWithContext(ctx context.Context) IotMappingPropertiesResponsePtrOutput {
	return o
}

func (o IotMappingPropertiesResponsePtrOutput) Elem() IotMappingPropertiesResponseOutput {
	return o.ApplyT(func(v *IotMappingPropertiesResponse) IotMappingPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret IotMappingPropertiesResponse
		return ret
	}).(IotMappingPropertiesResponseOutput)
}

func (o IotMappingPropertiesResponsePtrOutput) Content() pulumi.AnyOutput {
	return o.ApplyT(func(v *IotMappingPropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Content
	}).(pulumi.AnyOutput)
}

type PrivateEndpointConnectionType struct {
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionTypeInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput
	ToPrivateEndpointConnectionTypeOutputWithContext(context.Context) PrivateEndpointConnectionTypeOutput
}

type PrivateEndpointConnectionTypeArgs struct {
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateEndpointConnectionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return i.ToPrivateEndpointConnectionTypeOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeOutput)
}





type PrivateEndpointConnectionTypeArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput
	ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Context) PrivateEndpointConnectionTypeArrayOutput
}

type PrivateEndpointConnectionTypeArray []PrivateEndpointConnectionTypeInput

func (PrivateEndpointConnectionTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return i.ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeArrayOutput)
}

type PrivateEndpointConnectionTypeOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) PrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateOutput)
}

type PrivateEndpointConnectionTypeArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionType {
		return vs[0].([]PrivateEndpointConnectionType)[vs[1].(int)]
	}).(PrivateEndpointConnectionTypeOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	Type                              string                                    `pulumi:"type"`
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ResourceVersionPolicyConfiguration struct {
	Default               *string           `pulumi:"default"`
	ResourceTypeOverrides map[string]string `pulumi:"resourceTypeOverrides"`
}





type ResourceVersionPolicyConfigurationInput interface {
	pulumi.Input

	ToResourceVersionPolicyConfigurationOutput() ResourceVersionPolicyConfigurationOutput
	ToResourceVersionPolicyConfigurationOutputWithContext(context.Context) ResourceVersionPolicyConfigurationOutput
}

type ResourceVersionPolicyConfigurationArgs struct {
	Default               pulumi.StringPtrInput `pulumi:"default"`
	ResourceTypeOverrides pulumi.StringMapInput `pulumi:"resourceTypeOverrides"`
}

func (ResourceVersionPolicyConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceVersionPolicyConfiguration)(nil)).Elem()
}

func (i ResourceVersionPolicyConfigurationArgs) ToResourceVersionPolicyConfigurationOutput() ResourceVersionPolicyConfigurationOutput {
	return i.ToResourceVersionPolicyConfigurationOutputWithContext(context.Background())
}

func (i ResourceVersionPolicyConfigurationArgs) ToResourceVersionPolicyConfigurationOutputWithContext(ctx context.Context) ResourceVersionPolicyConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceVersionPolicyConfigurationOutput)
}

func (i ResourceVersionPolicyConfigurationArgs) ToResourceVersionPolicyConfigurationPtrOutput() ResourceVersionPolicyConfigurationPtrOutput {
	return i.ToResourceVersionPolicyConfigurationPtrOutputWithContext(context.Background())
}

func (i ResourceVersionPolicyConfigurationArgs) ToResourceVersionPolicyConfigurationPtrOutputWithContext(ctx context.Context) ResourceVersionPolicyConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceVersionPolicyConfigurationOutput).ToResourceVersionPolicyConfigurationPtrOutputWithContext(ctx)
}









type ResourceVersionPolicyConfigurationPtrInput interface {
	pulumi.Input

	ToResourceVersionPolicyConfigurationPtrOutput() ResourceVersionPolicyConfigurationPtrOutput
	ToResourceVersionPolicyConfigurationPtrOutputWithContext(context.Context) ResourceVersionPolicyConfigurationPtrOutput
}

type resourceVersionPolicyConfigurationPtrType ResourceVersionPolicyConfigurationArgs

func ResourceVersionPolicyConfigurationPtr(v *ResourceVersionPolicyConfigurationArgs) ResourceVersionPolicyConfigurationPtrInput {
	return (*resourceVersionPolicyConfigurationPtrType)(v)
}

func (*resourceVersionPolicyConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceVersionPolicyConfiguration)(nil)).Elem()
}

func (i *resourceVersionPolicyConfigurationPtrType) ToResourceVersionPolicyConfigurationPtrOutput() ResourceVersionPolicyConfigurationPtrOutput {
	return i.ToResourceVersionPolicyConfigurationPtrOutputWithContext(context.Background())
}

func (i *resourceVersionPolicyConfigurationPtrType) ToResourceVersionPolicyConfigurationPtrOutputWithContext(ctx context.Context) ResourceVersionPolicyConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceVersionPolicyConfigurationPtrOutput)
}

type ResourceVersionPolicyConfigurationOutput struct{ *pulumi.OutputState }

func (ResourceVersionPolicyConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceVersionPolicyConfiguration)(nil)).Elem()
}

func (o ResourceVersionPolicyConfigurationOutput) ToResourceVersionPolicyConfigurationOutput() ResourceVersionPolicyConfigurationOutput {
	return o
}

func (o ResourceVersionPolicyConfigurationOutput) ToResourceVersionPolicyConfigurationOutputWithContext(ctx context.Context) ResourceVersionPolicyConfigurationOutput {
	return o
}

func (o ResourceVersionPolicyConfigurationOutput) ToResourceVersionPolicyConfigurationPtrOutput() ResourceVersionPolicyConfigurationPtrOutput {
	return o.ToResourceVersionPolicyConfigurationPtrOutputWithContext(context.Background())
}

func (o ResourceVersionPolicyConfigurationOutput) ToResourceVersionPolicyConfigurationPtrOutputWithContext(ctx context.Context) ResourceVersionPolicyConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceVersionPolicyConfiguration) *ResourceVersionPolicyConfiguration {
		return &v
	}).(ResourceVersionPolicyConfigurationPtrOutput)
}

func (o ResourceVersionPolicyConfigurationOutput) Default() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceVersionPolicyConfiguration) *string { return v.Default }).(pulumi.StringPtrOutput)
}

func (o ResourceVersionPolicyConfigurationOutput) ResourceTypeOverrides() pulumi.StringMapOutput {
	return o.ApplyT(func(v ResourceVersionPolicyConfiguration) map[string]string { return v.ResourceTypeOverrides }).(pulumi.StringMapOutput)
}

type ResourceVersionPolicyConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ResourceVersionPolicyConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceVersionPolicyConfiguration)(nil)).Elem()
}

func (o ResourceVersionPolicyConfigurationPtrOutput) ToResourceVersionPolicyConfigurationPtrOutput() ResourceVersionPolicyConfigurationPtrOutput {
	return o
}

func (o ResourceVersionPolicyConfigurationPtrOutput) ToResourceVersionPolicyConfigurationPtrOutputWithContext(ctx context.Context) ResourceVersionPolicyConfigurationPtrOutput {
	return o
}

func (o ResourceVersionPolicyConfigurationPtrOutput) Elem() ResourceVersionPolicyConfigurationOutput {
	return o.ApplyT(func(v *ResourceVersionPolicyConfiguration) ResourceVersionPolicyConfiguration {
		if v != nil {
			return *v
		}
		var ret ResourceVersionPolicyConfiguration
		return ret
	}).(ResourceVersionPolicyConfigurationOutput)
}

func (o ResourceVersionPolicyConfigurationPtrOutput) Default() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceVersionPolicyConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Default
	}).(pulumi.StringPtrOutput)
}

func (o ResourceVersionPolicyConfigurationPtrOutput) ResourceTypeOverrides() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResourceVersionPolicyConfiguration) map[string]string {
		if v == nil {
			return nil
		}
		return v.ResourceTypeOverrides
	}).(pulumi.StringMapOutput)
}

type ResourceVersionPolicyConfigurationResponse struct {
	Default               *string           `pulumi:"default"`
	ResourceTypeOverrides map[string]string `pulumi:"resourceTypeOverrides"`
}

type ResourceVersionPolicyConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ResourceVersionPolicyConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceVersionPolicyConfigurationResponse)(nil)).Elem()
}

func (o ResourceVersionPolicyConfigurationResponseOutput) ToResourceVersionPolicyConfigurationResponseOutput() ResourceVersionPolicyConfigurationResponseOutput {
	return o
}

func (o ResourceVersionPolicyConfigurationResponseOutput) ToResourceVersionPolicyConfigurationResponseOutputWithContext(ctx context.Context) ResourceVersionPolicyConfigurationResponseOutput {
	return o
}

func (o ResourceVersionPolicyConfigurationResponseOutput) Default() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceVersionPolicyConfigurationResponse) *string { return v.Default }).(pulumi.StringPtrOutput)
}

func (o ResourceVersionPolicyConfigurationResponseOutput) ResourceTypeOverrides() pulumi.StringMapOutput {
	return o.ApplyT(func(v ResourceVersionPolicyConfigurationResponse) map[string]string { return v.ResourceTypeOverrides }).(pulumi.StringMapOutput)
}

type ResourceVersionPolicyConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceVersionPolicyConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceVersionPolicyConfigurationResponse)(nil)).Elem()
}

func (o ResourceVersionPolicyConfigurationResponsePtrOutput) ToResourceVersionPolicyConfigurationResponsePtrOutput() ResourceVersionPolicyConfigurationResponsePtrOutput {
	return o
}

func (o ResourceVersionPolicyConfigurationResponsePtrOutput) ToResourceVersionPolicyConfigurationResponsePtrOutputWithContext(ctx context.Context) ResourceVersionPolicyConfigurationResponsePtrOutput {
	return o
}

func (o ResourceVersionPolicyConfigurationResponsePtrOutput) Elem() ResourceVersionPolicyConfigurationResponseOutput {
	return o.ApplyT(func(v *ResourceVersionPolicyConfigurationResponse) ResourceVersionPolicyConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ResourceVersionPolicyConfigurationResponse
		return ret
	}).(ResourceVersionPolicyConfigurationResponseOutput)
}

func (o ResourceVersionPolicyConfigurationResponsePtrOutput) Default() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceVersionPolicyConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Default
	}).(pulumi.StringPtrOutput)
}

func (o ResourceVersionPolicyConfigurationResponsePtrOutput) ResourceTypeOverrides() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResourceVersionPolicyConfigurationResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.ResourceTypeOverrides
	}).(pulumi.StringMapOutput)
}

type ServiceAccessPolicyEntry struct {
	ObjectId string `pulumi:"objectId"`
}





type ServiceAccessPolicyEntryInput interface {
	pulumi.Input

	ToServiceAccessPolicyEntryOutput() ServiceAccessPolicyEntryOutput
	ToServiceAccessPolicyEntryOutputWithContext(context.Context) ServiceAccessPolicyEntryOutput
}

type ServiceAccessPolicyEntryArgs struct {
	ObjectId pulumi.StringInput `pulumi:"objectId"`
}

func (ServiceAccessPolicyEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAccessPolicyEntry)(nil)).Elem()
}

func (i ServiceAccessPolicyEntryArgs) ToServiceAccessPolicyEntryOutput() ServiceAccessPolicyEntryOutput {
	return i.ToServiceAccessPolicyEntryOutputWithContext(context.Background())
}

func (i ServiceAccessPolicyEntryArgs) ToServiceAccessPolicyEntryOutputWithContext(ctx context.Context) ServiceAccessPolicyEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccessPolicyEntryOutput)
}





type ServiceAccessPolicyEntryArrayInput interface {
	pulumi.Input

	ToServiceAccessPolicyEntryArrayOutput() ServiceAccessPolicyEntryArrayOutput
	ToServiceAccessPolicyEntryArrayOutputWithContext(context.Context) ServiceAccessPolicyEntryArrayOutput
}

type ServiceAccessPolicyEntryArray []ServiceAccessPolicyEntryInput

func (ServiceAccessPolicyEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceAccessPolicyEntry)(nil)).Elem()
}

func (i ServiceAccessPolicyEntryArray) ToServiceAccessPolicyEntryArrayOutput() ServiceAccessPolicyEntryArrayOutput {
	return i.ToServiceAccessPolicyEntryArrayOutputWithContext(context.Background())
}

func (i ServiceAccessPolicyEntryArray) ToServiceAccessPolicyEntryArrayOutputWithContext(ctx context.Context) ServiceAccessPolicyEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccessPolicyEntryArrayOutput)
}

type ServiceAccessPolicyEntryOutput struct{ *pulumi.OutputState }

func (ServiceAccessPolicyEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAccessPolicyEntry)(nil)).Elem()
}

func (o ServiceAccessPolicyEntryOutput) ToServiceAccessPolicyEntryOutput() ServiceAccessPolicyEntryOutput {
	return o
}

func (o ServiceAccessPolicyEntryOutput) ToServiceAccessPolicyEntryOutputWithContext(ctx context.Context) ServiceAccessPolicyEntryOutput {
	return o
}

func (o ServiceAccessPolicyEntryOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceAccessPolicyEntry) string { return v.ObjectId }).(pulumi.StringOutput)
}

type ServiceAccessPolicyEntryArrayOutput struct{ *pulumi.OutputState }

func (ServiceAccessPolicyEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceAccessPolicyEntry)(nil)).Elem()
}

func (o ServiceAccessPolicyEntryArrayOutput) ToServiceAccessPolicyEntryArrayOutput() ServiceAccessPolicyEntryArrayOutput {
	return o
}

func (o ServiceAccessPolicyEntryArrayOutput) ToServiceAccessPolicyEntryArrayOutputWithContext(ctx context.Context) ServiceAccessPolicyEntryArrayOutput {
	return o
}

func (o ServiceAccessPolicyEntryArrayOutput) Index(i pulumi.IntInput) ServiceAccessPolicyEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceAccessPolicyEntry {
		return vs[0].([]ServiceAccessPolicyEntry)[vs[1].(int)]
	}).(ServiceAccessPolicyEntryOutput)
}

type ServiceAccessPolicyEntryResponse struct {
	ObjectId string `pulumi:"objectId"`
}

type ServiceAccessPolicyEntryResponseOutput struct{ *pulumi.OutputState }

func (ServiceAccessPolicyEntryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAccessPolicyEntryResponse)(nil)).Elem()
}

func (o ServiceAccessPolicyEntryResponseOutput) ToServiceAccessPolicyEntryResponseOutput() ServiceAccessPolicyEntryResponseOutput {
	return o
}

func (o ServiceAccessPolicyEntryResponseOutput) ToServiceAccessPolicyEntryResponseOutputWithContext(ctx context.Context) ServiceAccessPolicyEntryResponseOutput {
	return o
}

func (o ServiceAccessPolicyEntryResponseOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceAccessPolicyEntryResponse) string { return v.ObjectId }).(pulumi.StringOutput)
}

type ServiceAccessPolicyEntryResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceAccessPolicyEntryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceAccessPolicyEntryResponse)(nil)).Elem()
}

func (o ServiceAccessPolicyEntryResponseArrayOutput) ToServiceAccessPolicyEntryResponseArrayOutput() ServiceAccessPolicyEntryResponseArrayOutput {
	return o
}

func (o ServiceAccessPolicyEntryResponseArrayOutput) ToServiceAccessPolicyEntryResponseArrayOutputWithContext(ctx context.Context) ServiceAccessPolicyEntryResponseArrayOutput {
	return o
}

func (o ServiceAccessPolicyEntryResponseArrayOutput) Index(i pulumi.IntInput) ServiceAccessPolicyEntryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceAccessPolicyEntryResponse {
		return vs[0].([]ServiceAccessPolicyEntryResponse)[vs[1].(int)]
	}).(ServiceAccessPolicyEntryResponseOutput)
}

type ServiceAcrConfigurationInfo struct {
	LoginServers []string                  `pulumi:"loginServers"`
	OciArtifacts []ServiceOciArtifactEntry `pulumi:"ociArtifacts"`
}





type ServiceAcrConfigurationInfoInput interface {
	pulumi.Input

	ToServiceAcrConfigurationInfoOutput() ServiceAcrConfigurationInfoOutput
	ToServiceAcrConfigurationInfoOutputWithContext(context.Context) ServiceAcrConfigurationInfoOutput
}

type ServiceAcrConfigurationInfoArgs struct {
	LoginServers pulumi.StringArrayInput           `pulumi:"loginServers"`
	OciArtifacts ServiceOciArtifactEntryArrayInput `pulumi:"ociArtifacts"`
}

func (ServiceAcrConfigurationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAcrConfigurationInfo)(nil)).Elem()
}

func (i ServiceAcrConfigurationInfoArgs) ToServiceAcrConfigurationInfoOutput() ServiceAcrConfigurationInfoOutput {
	return i.ToServiceAcrConfigurationInfoOutputWithContext(context.Background())
}

func (i ServiceAcrConfigurationInfoArgs) ToServiceAcrConfigurationInfoOutputWithContext(ctx context.Context) ServiceAcrConfigurationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAcrConfigurationInfoOutput)
}

func (i ServiceAcrConfigurationInfoArgs) ToServiceAcrConfigurationInfoPtrOutput() ServiceAcrConfigurationInfoPtrOutput {
	return i.ToServiceAcrConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i ServiceAcrConfigurationInfoArgs) ToServiceAcrConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceAcrConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAcrConfigurationInfoOutput).ToServiceAcrConfigurationInfoPtrOutputWithContext(ctx)
}









type ServiceAcrConfigurationInfoPtrInput interface {
	pulumi.Input

	ToServiceAcrConfigurationInfoPtrOutput() ServiceAcrConfigurationInfoPtrOutput
	ToServiceAcrConfigurationInfoPtrOutputWithContext(context.Context) ServiceAcrConfigurationInfoPtrOutput
}

type serviceAcrConfigurationInfoPtrType ServiceAcrConfigurationInfoArgs

func ServiceAcrConfigurationInfoPtr(v *ServiceAcrConfigurationInfoArgs) ServiceAcrConfigurationInfoPtrInput {
	return (*serviceAcrConfigurationInfoPtrType)(v)
}

func (*serviceAcrConfigurationInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAcrConfigurationInfo)(nil)).Elem()
}

func (i *serviceAcrConfigurationInfoPtrType) ToServiceAcrConfigurationInfoPtrOutput() ServiceAcrConfigurationInfoPtrOutput {
	return i.ToServiceAcrConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i *serviceAcrConfigurationInfoPtrType) ToServiceAcrConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceAcrConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAcrConfigurationInfoPtrOutput)
}

type ServiceAcrConfigurationInfoOutput struct{ *pulumi.OutputState }

func (ServiceAcrConfigurationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAcrConfigurationInfo)(nil)).Elem()
}

func (o ServiceAcrConfigurationInfoOutput) ToServiceAcrConfigurationInfoOutput() ServiceAcrConfigurationInfoOutput {
	return o
}

func (o ServiceAcrConfigurationInfoOutput) ToServiceAcrConfigurationInfoOutputWithContext(ctx context.Context) ServiceAcrConfigurationInfoOutput {
	return o
}

func (o ServiceAcrConfigurationInfoOutput) ToServiceAcrConfigurationInfoPtrOutput() ServiceAcrConfigurationInfoPtrOutput {
	return o.ToServiceAcrConfigurationInfoPtrOutputWithContext(context.Background())
}

func (o ServiceAcrConfigurationInfoOutput) ToServiceAcrConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceAcrConfigurationInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceAcrConfigurationInfo) *ServiceAcrConfigurationInfo {
		return &v
	}).(ServiceAcrConfigurationInfoPtrOutput)
}

func (o ServiceAcrConfigurationInfoOutput) LoginServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceAcrConfigurationInfo) []string { return v.LoginServers }).(pulumi.StringArrayOutput)
}

func (o ServiceAcrConfigurationInfoOutput) OciArtifacts() ServiceOciArtifactEntryArrayOutput {
	return o.ApplyT(func(v ServiceAcrConfigurationInfo) []ServiceOciArtifactEntry { return v.OciArtifacts }).(ServiceOciArtifactEntryArrayOutput)
}

type ServiceAcrConfigurationInfoPtrOutput struct{ *pulumi.OutputState }

func (ServiceAcrConfigurationInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAcrConfigurationInfo)(nil)).Elem()
}

func (o ServiceAcrConfigurationInfoPtrOutput) ToServiceAcrConfigurationInfoPtrOutput() ServiceAcrConfigurationInfoPtrOutput {
	return o
}

func (o ServiceAcrConfigurationInfoPtrOutput) ToServiceAcrConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceAcrConfigurationInfoPtrOutput {
	return o
}

func (o ServiceAcrConfigurationInfoPtrOutput) Elem() ServiceAcrConfigurationInfoOutput {
	return o.ApplyT(func(v *ServiceAcrConfigurationInfo) ServiceAcrConfigurationInfo {
		if v != nil {
			return *v
		}
		var ret ServiceAcrConfigurationInfo
		return ret
	}).(ServiceAcrConfigurationInfoOutput)
}

func (o ServiceAcrConfigurationInfoPtrOutput) LoginServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceAcrConfigurationInfo) []string {
		if v == nil {
			return nil
		}
		return v.LoginServers
	}).(pulumi.StringArrayOutput)
}

func (o ServiceAcrConfigurationInfoPtrOutput) OciArtifacts() ServiceOciArtifactEntryArrayOutput {
	return o.ApplyT(func(v *ServiceAcrConfigurationInfo) []ServiceOciArtifactEntry {
		if v == nil {
			return nil
		}
		return v.OciArtifacts
	}).(ServiceOciArtifactEntryArrayOutput)
}

type ServiceAcrConfigurationInfoResponse struct {
	LoginServers []string                          `pulumi:"loginServers"`
	OciArtifacts []ServiceOciArtifactEntryResponse `pulumi:"ociArtifacts"`
}

type ServiceAcrConfigurationInfoResponseOutput struct{ *pulumi.OutputState }

func (ServiceAcrConfigurationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAcrConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceAcrConfigurationInfoResponseOutput) ToServiceAcrConfigurationInfoResponseOutput() ServiceAcrConfigurationInfoResponseOutput {
	return o
}

func (o ServiceAcrConfigurationInfoResponseOutput) ToServiceAcrConfigurationInfoResponseOutputWithContext(ctx context.Context) ServiceAcrConfigurationInfoResponseOutput {
	return o
}

func (o ServiceAcrConfigurationInfoResponseOutput) LoginServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceAcrConfigurationInfoResponse) []string { return v.LoginServers }).(pulumi.StringArrayOutput)
}

func (o ServiceAcrConfigurationInfoResponseOutput) OciArtifacts() ServiceOciArtifactEntryResponseArrayOutput {
	return o.ApplyT(func(v ServiceAcrConfigurationInfoResponse) []ServiceOciArtifactEntryResponse { return v.OciArtifacts }).(ServiceOciArtifactEntryResponseArrayOutput)
}

type ServiceAcrConfigurationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceAcrConfigurationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAcrConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceAcrConfigurationInfoResponsePtrOutput) ToServiceAcrConfigurationInfoResponsePtrOutput() ServiceAcrConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceAcrConfigurationInfoResponsePtrOutput) ToServiceAcrConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceAcrConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceAcrConfigurationInfoResponsePtrOutput) Elem() ServiceAcrConfigurationInfoResponseOutput {
	return o.ApplyT(func(v *ServiceAcrConfigurationInfoResponse) ServiceAcrConfigurationInfoResponse {
		if v != nil {
			return *v
		}
		var ret ServiceAcrConfigurationInfoResponse
		return ret
	}).(ServiceAcrConfigurationInfoResponseOutput)
}

func (o ServiceAcrConfigurationInfoResponsePtrOutput) LoginServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceAcrConfigurationInfoResponse) []string {
		if v == nil {
			return nil
		}
		return v.LoginServers
	}).(pulumi.StringArrayOutput)
}

func (o ServiceAcrConfigurationInfoResponsePtrOutput) OciArtifacts() ServiceOciArtifactEntryResponseArrayOutput {
	return o.ApplyT(func(v *ServiceAcrConfigurationInfoResponse) []ServiceOciArtifactEntryResponse {
		if v == nil {
			return nil
		}
		return v.OciArtifacts
	}).(ServiceOciArtifactEntryResponseArrayOutput)
}

type ServiceAuthenticationConfigurationInfo struct {
	Audience          *string `pulumi:"audience"`
	Authority         *string `pulumi:"authority"`
	SmartProxyEnabled *bool   `pulumi:"smartProxyEnabled"`
}





type ServiceAuthenticationConfigurationInfoInput interface {
	pulumi.Input

	ToServiceAuthenticationConfigurationInfoOutput() ServiceAuthenticationConfigurationInfoOutput
	ToServiceAuthenticationConfigurationInfoOutputWithContext(context.Context) ServiceAuthenticationConfigurationInfoOutput
}

type ServiceAuthenticationConfigurationInfoArgs struct {
	Audience          pulumi.StringPtrInput `pulumi:"audience"`
	Authority         pulumi.StringPtrInput `pulumi:"authority"`
	SmartProxyEnabled pulumi.BoolPtrInput   `pulumi:"smartProxyEnabled"`
}

func (ServiceAuthenticationConfigurationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAuthenticationConfigurationInfo)(nil)).Elem()
}

func (i ServiceAuthenticationConfigurationInfoArgs) ToServiceAuthenticationConfigurationInfoOutput() ServiceAuthenticationConfigurationInfoOutput {
	return i.ToServiceAuthenticationConfigurationInfoOutputWithContext(context.Background())
}

func (i ServiceAuthenticationConfigurationInfoArgs) ToServiceAuthenticationConfigurationInfoOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthenticationConfigurationInfoOutput)
}

func (i ServiceAuthenticationConfigurationInfoArgs) ToServiceAuthenticationConfigurationInfoPtrOutput() ServiceAuthenticationConfigurationInfoPtrOutput {
	return i.ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i ServiceAuthenticationConfigurationInfoArgs) ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthenticationConfigurationInfoOutput).ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(ctx)
}









type ServiceAuthenticationConfigurationInfoPtrInput interface {
	pulumi.Input

	ToServiceAuthenticationConfigurationInfoPtrOutput() ServiceAuthenticationConfigurationInfoPtrOutput
	ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(context.Context) ServiceAuthenticationConfigurationInfoPtrOutput
}

type serviceAuthenticationConfigurationInfoPtrType ServiceAuthenticationConfigurationInfoArgs

func ServiceAuthenticationConfigurationInfoPtr(v *ServiceAuthenticationConfigurationInfoArgs) ServiceAuthenticationConfigurationInfoPtrInput {
	return (*serviceAuthenticationConfigurationInfoPtrType)(v)
}

func (*serviceAuthenticationConfigurationInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAuthenticationConfigurationInfo)(nil)).Elem()
}

func (i *serviceAuthenticationConfigurationInfoPtrType) ToServiceAuthenticationConfigurationInfoPtrOutput() ServiceAuthenticationConfigurationInfoPtrOutput {
	return i.ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i *serviceAuthenticationConfigurationInfoPtrType) ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthenticationConfigurationInfoPtrOutput)
}

type ServiceAuthenticationConfigurationInfoOutput struct{ *pulumi.OutputState }

func (ServiceAuthenticationConfigurationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAuthenticationConfigurationInfo)(nil)).Elem()
}

func (o ServiceAuthenticationConfigurationInfoOutput) ToServiceAuthenticationConfigurationInfoOutput() ServiceAuthenticationConfigurationInfoOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoOutput) ToServiceAuthenticationConfigurationInfoOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoOutput) ToServiceAuthenticationConfigurationInfoPtrOutput() ServiceAuthenticationConfigurationInfoPtrOutput {
	return o.ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(context.Background())
}

func (o ServiceAuthenticationConfigurationInfoOutput) ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceAuthenticationConfigurationInfo) *ServiceAuthenticationConfigurationInfo {
		return &v
	}).(ServiceAuthenticationConfigurationInfoPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfigurationInfo) *string { return v.Audience }).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfigurationInfo) *string { return v.Authority }).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoOutput) SmartProxyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfigurationInfo) *bool { return v.SmartProxyEnabled }).(pulumi.BoolPtrOutput)
}

type ServiceAuthenticationConfigurationInfoPtrOutput struct{ *pulumi.OutputState }

func (ServiceAuthenticationConfigurationInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAuthenticationConfigurationInfo)(nil)).Elem()
}

func (o ServiceAuthenticationConfigurationInfoPtrOutput) ToServiceAuthenticationConfigurationInfoPtrOutput() ServiceAuthenticationConfigurationInfoPtrOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoPtrOutput) ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoPtrOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoPtrOutput) Elem() ServiceAuthenticationConfigurationInfoOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfo) ServiceAuthenticationConfigurationInfo {
		if v != nil {
			return *v
		}
		var ret ServiceAuthenticationConfigurationInfo
		return ret
	}).(ServiceAuthenticationConfigurationInfoOutput)
}

func (o ServiceAuthenticationConfigurationInfoPtrOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfo) *string {
		if v == nil {
			return nil
		}
		return v.Audience
	}).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoPtrOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfo) *string {
		if v == nil {
			return nil
		}
		return v.Authority
	}).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoPtrOutput) SmartProxyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfo) *bool {
		if v == nil {
			return nil
		}
		return v.SmartProxyEnabled
	}).(pulumi.BoolPtrOutput)
}

type ServiceAuthenticationConfigurationInfoResponse struct {
	Audience          *string `pulumi:"audience"`
	Authority         *string `pulumi:"authority"`
	SmartProxyEnabled *bool   `pulumi:"smartProxyEnabled"`
}

type ServiceAuthenticationConfigurationInfoResponseOutput struct{ *pulumi.OutputState }

func (ServiceAuthenticationConfigurationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAuthenticationConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceAuthenticationConfigurationInfoResponseOutput) ToServiceAuthenticationConfigurationInfoResponseOutput() ServiceAuthenticationConfigurationInfoResponseOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoResponseOutput) ToServiceAuthenticationConfigurationInfoResponseOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoResponseOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoResponseOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfigurationInfoResponse) *string { return v.Audience }).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoResponseOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfigurationInfoResponse) *string { return v.Authority }).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoResponseOutput) SmartProxyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfigurationInfoResponse) *bool { return v.SmartProxyEnabled }).(pulumi.BoolPtrOutput)
}

type ServiceAuthenticationConfigurationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceAuthenticationConfigurationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAuthenticationConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceAuthenticationConfigurationInfoResponsePtrOutput) ToServiceAuthenticationConfigurationInfoResponsePtrOutput() ServiceAuthenticationConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoResponsePtrOutput) ToServiceAuthenticationConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoResponsePtrOutput) Elem() ServiceAuthenticationConfigurationInfoResponseOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfoResponse) ServiceAuthenticationConfigurationInfoResponse {
		if v != nil {
			return *v
		}
		var ret ServiceAuthenticationConfigurationInfoResponse
		return ret
	}).(ServiceAuthenticationConfigurationInfoResponseOutput)
}

func (o ServiceAuthenticationConfigurationInfoResponsePtrOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Audience
	}).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoResponsePtrOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Authority
	}).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoResponsePtrOutput) SmartProxyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfoResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SmartProxyEnabled
	}).(pulumi.BoolPtrOutput)
}

type ServiceCorsConfigurationInfo struct {
	AllowCredentials *bool    `pulumi:"allowCredentials"`
	Headers          []string `pulumi:"headers"`
	MaxAge           *int     `pulumi:"maxAge"`
	Methods          []string `pulumi:"methods"`
	Origins          []string `pulumi:"origins"`
}





type ServiceCorsConfigurationInfoInput interface {
	pulumi.Input

	ToServiceCorsConfigurationInfoOutput() ServiceCorsConfigurationInfoOutput
	ToServiceCorsConfigurationInfoOutputWithContext(context.Context) ServiceCorsConfigurationInfoOutput
}

type ServiceCorsConfigurationInfoArgs struct {
	AllowCredentials pulumi.BoolPtrInput     `pulumi:"allowCredentials"`
	Headers          pulumi.StringArrayInput `pulumi:"headers"`
	MaxAge           pulumi.IntPtrInput      `pulumi:"maxAge"`
	Methods          pulumi.StringArrayInput `pulumi:"methods"`
	Origins          pulumi.StringArrayInput `pulumi:"origins"`
}

func (ServiceCorsConfigurationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorsConfigurationInfo)(nil)).Elem()
}

func (i ServiceCorsConfigurationInfoArgs) ToServiceCorsConfigurationInfoOutput() ServiceCorsConfigurationInfoOutput {
	return i.ToServiceCorsConfigurationInfoOutputWithContext(context.Background())
}

func (i ServiceCorsConfigurationInfoArgs) ToServiceCorsConfigurationInfoOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorsConfigurationInfoOutput)
}

func (i ServiceCorsConfigurationInfoArgs) ToServiceCorsConfigurationInfoPtrOutput() ServiceCorsConfigurationInfoPtrOutput {
	return i.ToServiceCorsConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i ServiceCorsConfigurationInfoArgs) ToServiceCorsConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorsConfigurationInfoOutput).ToServiceCorsConfigurationInfoPtrOutputWithContext(ctx)
}









type ServiceCorsConfigurationInfoPtrInput interface {
	pulumi.Input

	ToServiceCorsConfigurationInfoPtrOutput() ServiceCorsConfigurationInfoPtrOutput
	ToServiceCorsConfigurationInfoPtrOutputWithContext(context.Context) ServiceCorsConfigurationInfoPtrOutput
}

type serviceCorsConfigurationInfoPtrType ServiceCorsConfigurationInfoArgs

func ServiceCorsConfigurationInfoPtr(v *ServiceCorsConfigurationInfoArgs) ServiceCorsConfigurationInfoPtrInput {
	return (*serviceCorsConfigurationInfoPtrType)(v)
}

func (*serviceCorsConfigurationInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCorsConfigurationInfo)(nil)).Elem()
}

func (i *serviceCorsConfigurationInfoPtrType) ToServiceCorsConfigurationInfoPtrOutput() ServiceCorsConfigurationInfoPtrOutput {
	return i.ToServiceCorsConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i *serviceCorsConfigurationInfoPtrType) ToServiceCorsConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorsConfigurationInfoPtrOutput)
}

type ServiceCorsConfigurationInfoOutput struct{ *pulumi.OutputState }

func (ServiceCorsConfigurationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorsConfigurationInfo)(nil)).Elem()
}

func (o ServiceCorsConfigurationInfoOutput) ToServiceCorsConfigurationInfoOutput() ServiceCorsConfigurationInfoOutput {
	return o
}

func (o ServiceCorsConfigurationInfoOutput) ToServiceCorsConfigurationInfoOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoOutput {
	return o
}

func (o ServiceCorsConfigurationInfoOutput) ToServiceCorsConfigurationInfoPtrOutput() ServiceCorsConfigurationInfoPtrOutput {
	return o.ToServiceCorsConfigurationInfoPtrOutputWithContext(context.Background())
}

func (o ServiceCorsConfigurationInfoOutput) ToServiceCorsConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceCorsConfigurationInfo) *ServiceCorsConfigurationInfo {
		return &v
	}).(ServiceCorsConfigurationInfoPtrOutput)
}

func (o ServiceCorsConfigurationInfoOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfo) *bool { return v.AllowCredentials }).(pulumi.BoolPtrOutput)
}

func (o ServiceCorsConfigurationInfoOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfo) []string { return v.Headers }).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoOutput) MaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfo) *int { return v.MaxAge }).(pulumi.IntPtrOutput)
}

func (o ServiceCorsConfigurationInfoOutput) Methods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfo) []string { return v.Methods }).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoOutput) Origins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfo) []string { return v.Origins }).(pulumi.StringArrayOutput)
}

type ServiceCorsConfigurationInfoPtrOutput struct{ *pulumi.OutputState }

func (ServiceCorsConfigurationInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCorsConfigurationInfo)(nil)).Elem()
}

func (o ServiceCorsConfigurationInfoPtrOutput) ToServiceCorsConfigurationInfoPtrOutput() ServiceCorsConfigurationInfoPtrOutput {
	return o
}

func (o ServiceCorsConfigurationInfoPtrOutput) ToServiceCorsConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoPtrOutput {
	return o
}

func (o ServiceCorsConfigurationInfoPtrOutput) Elem() ServiceCorsConfigurationInfoOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfo) ServiceCorsConfigurationInfo {
		if v != nil {
			return *v
		}
		var ret ServiceCorsConfigurationInfo
		return ret
	}).(ServiceCorsConfigurationInfoOutput)
}

func (o ServiceCorsConfigurationInfoPtrOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfo) *bool {
		if v == nil {
			return nil
		}
		return v.AllowCredentials
	}).(pulumi.BoolPtrOutput)
}

func (o ServiceCorsConfigurationInfoPtrOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfo) []string {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoPtrOutput) MaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfo) *int {
		if v == nil {
			return nil
		}
		return v.MaxAge
	}).(pulumi.IntPtrOutput)
}

func (o ServiceCorsConfigurationInfoPtrOutput) Methods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfo) []string {
		if v == nil {
			return nil
		}
		return v.Methods
	}).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoPtrOutput) Origins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfo) []string {
		if v == nil {
			return nil
		}
		return v.Origins
	}).(pulumi.StringArrayOutput)
}

type ServiceCorsConfigurationInfoResponse struct {
	AllowCredentials *bool    `pulumi:"allowCredentials"`
	Headers          []string `pulumi:"headers"`
	MaxAge           *int     `pulumi:"maxAge"`
	Methods          []string `pulumi:"methods"`
	Origins          []string `pulumi:"origins"`
}

type ServiceCorsConfigurationInfoResponseOutput struct{ *pulumi.OutputState }

func (ServiceCorsConfigurationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorsConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceCorsConfigurationInfoResponseOutput) ToServiceCorsConfigurationInfoResponseOutput() ServiceCorsConfigurationInfoResponseOutput {
	return o
}

func (o ServiceCorsConfigurationInfoResponseOutput) ToServiceCorsConfigurationInfoResponseOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoResponseOutput {
	return o
}

func (o ServiceCorsConfigurationInfoResponseOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfoResponse) *bool { return v.AllowCredentials }).(pulumi.BoolPtrOutput)
}

func (o ServiceCorsConfigurationInfoResponseOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfoResponse) []string { return v.Headers }).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoResponseOutput) MaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfoResponse) *int { return v.MaxAge }).(pulumi.IntPtrOutput)
}

func (o ServiceCorsConfigurationInfoResponseOutput) Methods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfoResponse) []string { return v.Methods }).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoResponseOutput) Origins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfoResponse) []string { return v.Origins }).(pulumi.StringArrayOutput)
}

type ServiceCorsConfigurationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceCorsConfigurationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCorsConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) ToServiceCorsConfigurationInfoResponsePtrOutput() ServiceCorsConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) ToServiceCorsConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) Elem() ServiceCorsConfigurationInfoResponseOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfoResponse) ServiceCorsConfigurationInfoResponse {
		if v != nil {
			return *v
		}
		var ret ServiceCorsConfigurationInfoResponse
		return ret
	}).(ServiceCorsConfigurationInfoResponseOutput)
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfoResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AllowCredentials
	}).(pulumi.BoolPtrOutput)
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfoResponse) []string {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) MaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfoResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxAge
	}).(pulumi.IntPtrOutput)
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) Methods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfoResponse) []string {
		if v == nil {
			return nil
		}
		return v.Methods
	}).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) Origins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfoResponse) []string {
		if v == nil {
			return nil
		}
		return v.Origins
	}).(pulumi.StringArrayOutput)
}

type ServiceCosmosDbConfigurationInfo struct {
	KeyVaultKeyUri  *string `pulumi:"keyVaultKeyUri"`
	OfferThroughput *int    `pulumi:"offerThroughput"`
}





type ServiceCosmosDbConfigurationInfoInput interface {
	pulumi.Input

	ToServiceCosmosDbConfigurationInfoOutput() ServiceCosmosDbConfigurationInfoOutput
	ToServiceCosmosDbConfigurationInfoOutputWithContext(context.Context) ServiceCosmosDbConfigurationInfoOutput
}

type ServiceCosmosDbConfigurationInfoArgs struct {
	KeyVaultKeyUri  pulumi.StringPtrInput `pulumi:"keyVaultKeyUri"`
	OfferThroughput pulumi.IntPtrInput    `pulumi:"offerThroughput"`
}

func (ServiceCosmosDbConfigurationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCosmosDbConfigurationInfo)(nil)).Elem()
}

func (i ServiceCosmosDbConfigurationInfoArgs) ToServiceCosmosDbConfigurationInfoOutput() ServiceCosmosDbConfigurationInfoOutput {
	return i.ToServiceCosmosDbConfigurationInfoOutputWithContext(context.Background())
}

func (i ServiceCosmosDbConfigurationInfoArgs) ToServiceCosmosDbConfigurationInfoOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCosmosDbConfigurationInfoOutput)
}

func (i ServiceCosmosDbConfigurationInfoArgs) ToServiceCosmosDbConfigurationInfoPtrOutput() ServiceCosmosDbConfigurationInfoPtrOutput {
	return i.ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i ServiceCosmosDbConfigurationInfoArgs) ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCosmosDbConfigurationInfoOutput).ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(ctx)
}









type ServiceCosmosDbConfigurationInfoPtrInput interface {
	pulumi.Input

	ToServiceCosmosDbConfigurationInfoPtrOutput() ServiceCosmosDbConfigurationInfoPtrOutput
	ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(context.Context) ServiceCosmosDbConfigurationInfoPtrOutput
}

type serviceCosmosDbConfigurationInfoPtrType ServiceCosmosDbConfigurationInfoArgs

func ServiceCosmosDbConfigurationInfoPtr(v *ServiceCosmosDbConfigurationInfoArgs) ServiceCosmosDbConfigurationInfoPtrInput {
	return (*serviceCosmosDbConfigurationInfoPtrType)(v)
}

func (*serviceCosmosDbConfigurationInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCosmosDbConfigurationInfo)(nil)).Elem()
}

func (i *serviceCosmosDbConfigurationInfoPtrType) ToServiceCosmosDbConfigurationInfoPtrOutput() ServiceCosmosDbConfigurationInfoPtrOutput {
	return i.ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i *serviceCosmosDbConfigurationInfoPtrType) ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCosmosDbConfigurationInfoPtrOutput)
}

type ServiceCosmosDbConfigurationInfoOutput struct{ *pulumi.OutputState }

func (ServiceCosmosDbConfigurationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCosmosDbConfigurationInfo)(nil)).Elem()
}

func (o ServiceCosmosDbConfigurationInfoOutput) ToServiceCosmosDbConfigurationInfoOutput() ServiceCosmosDbConfigurationInfoOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoOutput) ToServiceCosmosDbConfigurationInfoOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoOutput) ToServiceCosmosDbConfigurationInfoPtrOutput() ServiceCosmosDbConfigurationInfoPtrOutput {
	return o.ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(context.Background())
}

func (o ServiceCosmosDbConfigurationInfoOutput) ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceCosmosDbConfigurationInfo) *ServiceCosmosDbConfigurationInfo {
		return &v
	}).(ServiceCosmosDbConfigurationInfoPtrOutput)
}

func (o ServiceCosmosDbConfigurationInfoOutput) KeyVaultKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceCosmosDbConfigurationInfo) *string { return v.KeyVaultKeyUri }).(pulumi.StringPtrOutput)
}

func (o ServiceCosmosDbConfigurationInfoOutput) OfferThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceCosmosDbConfigurationInfo) *int { return v.OfferThroughput }).(pulumi.IntPtrOutput)
}

type ServiceCosmosDbConfigurationInfoPtrOutput struct{ *pulumi.OutputState }

func (ServiceCosmosDbConfigurationInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCosmosDbConfigurationInfo)(nil)).Elem()
}

func (o ServiceCosmosDbConfigurationInfoPtrOutput) ToServiceCosmosDbConfigurationInfoPtrOutput() ServiceCosmosDbConfigurationInfoPtrOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoPtrOutput) ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoPtrOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoPtrOutput) Elem() ServiceCosmosDbConfigurationInfoOutput {
	return o.ApplyT(func(v *ServiceCosmosDbConfigurationInfo) ServiceCosmosDbConfigurationInfo {
		if v != nil {
			return *v
		}
		var ret ServiceCosmosDbConfigurationInfo
		return ret
	}).(ServiceCosmosDbConfigurationInfoOutput)
}

func (o ServiceCosmosDbConfigurationInfoPtrOutput) KeyVaultKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceCosmosDbConfigurationInfo) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultKeyUri
	}).(pulumi.StringPtrOutput)
}

func (o ServiceCosmosDbConfigurationInfoPtrOutput) OfferThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceCosmosDbConfigurationInfo) *int {
		if v == nil {
			return nil
		}
		return v.OfferThroughput
	}).(pulumi.IntPtrOutput)
}

type ServiceCosmosDbConfigurationInfoResponse struct {
	KeyVaultKeyUri  *string `pulumi:"keyVaultKeyUri"`
	OfferThroughput *int    `pulumi:"offerThroughput"`
}

type ServiceCosmosDbConfigurationInfoResponseOutput struct{ *pulumi.OutputState }

func (ServiceCosmosDbConfigurationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCosmosDbConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceCosmosDbConfigurationInfoResponseOutput) ToServiceCosmosDbConfigurationInfoResponseOutput() ServiceCosmosDbConfigurationInfoResponseOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoResponseOutput) ToServiceCosmosDbConfigurationInfoResponseOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoResponseOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoResponseOutput) KeyVaultKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceCosmosDbConfigurationInfoResponse) *string { return v.KeyVaultKeyUri }).(pulumi.StringPtrOutput)
}

func (o ServiceCosmosDbConfigurationInfoResponseOutput) OfferThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceCosmosDbConfigurationInfoResponse) *int { return v.OfferThroughput }).(pulumi.IntPtrOutput)
}

type ServiceCosmosDbConfigurationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceCosmosDbConfigurationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCosmosDbConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceCosmosDbConfigurationInfoResponsePtrOutput) ToServiceCosmosDbConfigurationInfoResponsePtrOutput() ServiceCosmosDbConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoResponsePtrOutput) ToServiceCosmosDbConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoResponsePtrOutput) Elem() ServiceCosmosDbConfigurationInfoResponseOutput {
	return o.ApplyT(func(v *ServiceCosmosDbConfigurationInfoResponse) ServiceCosmosDbConfigurationInfoResponse {
		if v != nil {
			return *v
		}
		var ret ServiceCosmosDbConfigurationInfoResponse
		return ret
	}).(ServiceCosmosDbConfigurationInfoResponseOutput)
}

func (o ServiceCosmosDbConfigurationInfoResponsePtrOutput) KeyVaultKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceCosmosDbConfigurationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultKeyUri
	}).(pulumi.StringPtrOutput)
}

func (o ServiceCosmosDbConfigurationInfoResponsePtrOutput) OfferThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceCosmosDbConfigurationInfoResponse) *int {
		if v == nil {
			return nil
		}
		return v.OfferThroughput
	}).(pulumi.IntPtrOutput)
}

type ServiceExportConfigurationInfo struct {
	StorageAccountName *string `pulumi:"storageAccountName"`
}





type ServiceExportConfigurationInfoInput interface {
	pulumi.Input

	ToServiceExportConfigurationInfoOutput() ServiceExportConfigurationInfoOutput
	ToServiceExportConfigurationInfoOutputWithContext(context.Context) ServiceExportConfigurationInfoOutput
}

type ServiceExportConfigurationInfoArgs struct {
	StorageAccountName pulumi.StringPtrInput `pulumi:"storageAccountName"`
}

func (ServiceExportConfigurationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceExportConfigurationInfo)(nil)).Elem()
}

func (i ServiceExportConfigurationInfoArgs) ToServiceExportConfigurationInfoOutput() ServiceExportConfigurationInfoOutput {
	return i.ToServiceExportConfigurationInfoOutputWithContext(context.Background())
}

func (i ServiceExportConfigurationInfoArgs) ToServiceExportConfigurationInfoOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceExportConfigurationInfoOutput)
}

func (i ServiceExportConfigurationInfoArgs) ToServiceExportConfigurationInfoPtrOutput() ServiceExportConfigurationInfoPtrOutput {
	return i.ToServiceExportConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i ServiceExportConfigurationInfoArgs) ToServiceExportConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceExportConfigurationInfoOutput).ToServiceExportConfigurationInfoPtrOutputWithContext(ctx)
}









type ServiceExportConfigurationInfoPtrInput interface {
	pulumi.Input

	ToServiceExportConfigurationInfoPtrOutput() ServiceExportConfigurationInfoPtrOutput
	ToServiceExportConfigurationInfoPtrOutputWithContext(context.Context) ServiceExportConfigurationInfoPtrOutput
}

type serviceExportConfigurationInfoPtrType ServiceExportConfigurationInfoArgs

func ServiceExportConfigurationInfoPtr(v *ServiceExportConfigurationInfoArgs) ServiceExportConfigurationInfoPtrInput {
	return (*serviceExportConfigurationInfoPtrType)(v)
}

func (*serviceExportConfigurationInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceExportConfigurationInfo)(nil)).Elem()
}

func (i *serviceExportConfigurationInfoPtrType) ToServiceExportConfigurationInfoPtrOutput() ServiceExportConfigurationInfoPtrOutput {
	return i.ToServiceExportConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i *serviceExportConfigurationInfoPtrType) ToServiceExportConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceExportConfigurationInfoPtrOutput)
}

type ServiceExportConfigurationInfoOutput struct{ *pulumi.OutputState }

func (ServiceExportConfigurationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceExportConfigurationInfo)(nil)).Elem()
}

func (o ServiceExportConfigurationInfoOutput) ToServiceExportConfigurationInfoOutput() ServiceExportConfigurationInfoOutput {
	return o
}

func (o ServiceExportConfigurationInfoOutput) ToServiceExportConfigurationInfoOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoOutput {
	return o
}

func (o ServiceExportConfigurationInfoOutput) ToServiceExportConfigurationInfoPtrOutput() ServiceExportConfigurationInfoPtrOutput {
	return o.ToServiceExportConfigurationInfoPtrOutputWithContext(context.Background())
}

func (o ServiceExportConfigurationInfoOutput) ToServiceExportConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceExportConfigurationInfo) *ServiceExportConfigurationInfo {
		return &v
	}).(ServiceExportConfigurationInfoPtrOutput)
}

func (o ServiceExportConfigurationInfoOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceExportConfigurationInfo) *string { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

type ServiceExportConfigurationInfoPtrOutput struct{ *pulumi.OutputState }

func (ServiceExportConfigurationInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceExportConfigurationInfo)(nil)).Elem()
}

func (o ServiceExportConfigurationInfoPtrOutput) ToServiceExportConfigurationInfoPtrOutput() ServiceExportConfigurationInfoPtrOutput {
	return o
}

func (o ServiceExportConfigurationInfoPtrOutput) ToServiceExportConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoPtrOutput {
	return o
}

func (o ServiceExportConfigurationInfoPtrOutput) Elem() ServiceExportConfigurationInfoOutput {
	return o.ApplyT(func(v *ServiceExportConfigurationInfo) ServiceExportConfigurationInfo {
		if v != nil {
			return *v
		}
		var ret ServiceExportConfigurationInfo
		return ret
	}).(ServiceExportConfigurationInfoOutput)
}

func (o ServiceExportConfigurationInfoPtrOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceExportConfigurationInfo) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountName
	}).(pulumi.StringPtrOutput)
}

type ServiceExportConfigurationInfoResponse struct {
	StorageAccountName *string `pulumi:"storageAccountName"`
}

type ServiceExportConfigurationInfoResponseOutput struct{ *pulumi.OutputState }

func (ServiceExportConfigurationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceExportConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceExportConfigurationInfoResponseOutput) ToServiceExportConfigurationInfoResponseOutput() ServiceExportConfigurationInfoResponseOutput {
	return o
}

func (o ServiceExportConfigurationInfoResponseOutput) ToServiceExportConfigurationInfoResponseOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoResponseOutput {
	return o
}

func (o ServiceExportConfigurationInfoResponseOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceExportConfigurationInfoResponse) *string { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

type ServiceExportConfigurationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceExportConfigurationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceExportConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceExportConfigurationInfoResponsePtrOutput) ToServiceExportConfigurationInfoResponsePtrOutput() ServiceExportConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceExportConfigurationInfoResponsePtrOutput) ToServiceExportConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceExportConfigurationInfoResponsePtrOutput) Elem() ServiceExportConfigurationInfoResponseOutput {
	return o.ApplyT(func(v *ServiceExportConfigurationInfoResponse) ServiceExportConfigurationInfoResponse {
		if v != nil {
			return *v
		}
		var ret ServiceExportConfigurationInfoResponse
		return ret
	}).(ServiceExportConfigurationInfoResponseOutput)
}

func (o ServiceExportConfigurationInfoResponsePtrOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceExportConfigurationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountName
	}).(pulumi.StringPtrOutput)
}

type ServiceManagedIdentityIdentity struct {
	Type                   string                 `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ServiceManagedIdentityIdentityInput interface {
	pulumi.Input

	ToServiceManagedIdentityIdentityOutput() ServiceManagedIdentityIdentityOutput
	ToServiceManagedIdentityIdentityOutputWithContext(context.Context) ServiceManagedIdentityIdentityOutput
}

type ServiceManagedIdentityIdentityArgs struct {
	Type                   pulumi.StringInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput    `pulumi:"userAssignedIdentities"`
}

func (ServiceManagedIdentityIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceManagedIdentityIdentity)(nil)).Elem()
}

func (i ServiceManagedIdentityIdentityArgs) ToServiceManagedIdentityIdentityOutput() ServiceManagedIdentityIdentityOutput {
	return i.ToServiceManagedIdentityIdentityOutputWithContext(context.Background())
}

func (i ServiceManagedIdentityIdentityArgs) ToServiceManagedIdentityIdentityOutputWithContext(ctx context.Context) ServiceManagedIdentityIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceManagedIdentityIdentityOutput)
}

func (i ServiceManagedIdentityIdentityArgs) ToServiceManagedIdentityIdentityPtrOutput() ServiceManagedIdentityIdentityPtrOutput {
	return i.ToServiceManagedIdentityIdentityPtrOutputWithContext(context.Background())
}

func (i ServiceManagedIdentityIdentityArgs) ToServiceManagedIdentityIdentityPtrOutputWithContext(ctx context.Context) ServiceManagedIdentityIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceManagedIdentityIdentityOutput).ToServiceManagedIdentityIdentityPtrOutputWithContext(ctx)
}









type ServiceManagedIdentityIdentityPtrInput interface {
	pulumi.Input

	ToServiceManagedIdentityIdentityPtrOutput() ServiceManagedIdentityIdentityPtrOutput
	ToServiceManagedIdentityIdentityPtrOutputWithContext(context.Context) ServiceManagedIdentityIdentityPtrOutput
}

type serviceManagedIdentityIdentityPtrType ServiceManagedIdentityIdentityArgs

func ServiceManagedIdentityIdentityPtr(v *ServiceManagedIdentityIdentityArgs) ServiceManagedIdentityIdentityPtrInput {
	return (*serviceManagedIdentityIdentityPtrType)(v)
}

func (*serviceManagedIdentityIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceManagedIdentityIdentity)(nil)).Elem()
}

func (i *serviceManagedIdentityIdentityPtrType) ToServiceManagedIdentityIdentityPtrOutput() ServiceManagedIdentityIdentityPtrOutput {
	return i.ToServiceManagedIdentityIdentityPtrOutputWithContext(context.Background())
}

func (i *serviceManagedIdentityIdentityPtrType) ToServiceManagedIdentityIdentityPtrOutputWithContext(ctx context.Context) ServiceManagedIdentityIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceManagedIdentityIdentityPtrOutput)
}

type ServiceManagedIdentityIdentityOutput struct{ *pulumi.OutputState }

func (ServiceManagedIdentityIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceManagedIdentityIdentity)(nil)).Elem()
}

func (o ServiceManagedIdentityIdentityOutput) ToServiceManagedIdentityIdentityOutput() ServiceManagedIdentityIdentityOutput {
	return o
}

func (o ServiceManagedIdentityIdentityOutput) ToServiceManagedIdentityIdentityOutputWithContext(ctx context.Context) ServiceManagedIdentityIdentityOutput {
	return o
}

func (o ServiceManagedIdentityIdentityOutput) ToServiceManagedIdentityIdentityPtrOutput() ServiceManagedIdentityIdentityPtrOutput {
	return o.ToServiceManagedIdentityIdentityPtrOutputWithContext(context.Background())
}

func (o ServiceManagedIdentityIdentityOutput) ToServiceManagedIdentityIdentityPtrOutputWithContext(ctx context.Context) ServiceManagedIdentityIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceManagedIdentityIdentity) *ServiceManagedIdentityIdentity {
		return &v
	}).(ServiceManagedIdentityIdentityPtrOutput)
}

func (o ServiceManagedIdentityIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceManagedIdentityIdentity) string { return v.Type }).(pulumi.StringOutput)
}

func (o ServiceManagedIdentityIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ServiceManagedIdentityIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ServiceManagedIdentityIdentityPtrOutput struct{ *pulumi.OutputState }

func (ServiceManagedIdentityIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceManagedIdentityIdentity)(nil)).Elem()
}

func (o ServiceManagedIdentityIdentityPtrOutput) ToServiceManagedIdentityIdentityPtrOutput() ServiceManagedIdentityIdentityPtrOutput {
	return o
}

func (o ServiceManagedIdentityIdentityPtrOutput) ToServiceManagedIdentityIdentityPtrOutputWithContext(ctx context.Context) ServiceManagedIdentityIdentityPtrOutput {
	return o
}

func (o ServiceManagedIdentityIdentityPtrOutput) Elem() ServiceManagedIdentityIdentityOutput {
	return o.ApplyT(func(v *ServiceManagedIdentityIdentity) ServiceManagedIdentityIdentity {
		if v != nil {
			return *v
		}
		var ret ServiceManagedIdentityIdentity
		return ret
	}).(ServiceManagedIdentityIdentityOutput)
}

func (o ServiceManagedIdentityIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceManagedIdentityIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ServiceManagedIdentityIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ServiceManagedIdentityIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ServiceManagedIdentityResponseIdentity struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   string                                  `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}

type ServiceManagedIdentityResponseIdentityOutput struct{ *pulumi.OutputState }

func (ServiceManagedIdentityResponseIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceManagedIdentityResponseIdentity)(nil)).Elem()
}

func (o ServiceManagedIdentityResponseIdentityOutput) ToServiceManagedIdentityResponseIdentityOutput() ServiceManagedIdentityResponseIdentityOutput {
	return o
}

func (o ServiceManagedIdentityResponseIdentityOutput) ToServiceManagedIdentityResponseIdentityOutputWithContext(ctx context.Context) ServiceManagedIdentityResponseIdentityOutput {
	return o
}

func (o ServiceManagedIdentityResponseIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceManagedIdentityResponseIdentity) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ServiceManagedIdentityResponseIdentityOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceManagedIdentityResponseIdentity) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ServiceManagedIdentityResponseIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceManagedIdentityResponseIdentity) string { return v.Type }).(pulumi.StringOutput)
}

func (o ServiceManagedIdentityResponseIdentityOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v ServiceManagedIdentityResponseIdentity) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ServiceManagedIdentityResponseIdentityPtrOutput struct{ *pulumi.OutputState }

func (ServiceManagedIdentityResponseIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceManagedIdentityResponseIdentity)(nil)).Elem()
}

func (o ServiceManagedIdentityResponseIdentityPtrOutput) ToServiceManagedIdentityResponseIdentityPtrOutput() ServiceManagedIdentityResponseIdentityPtrOutput {
	return o
}

func (o ServiceManagedIdentityResponseIdentityPtrOutput) ToServiceManagedIdentityResponseIdentityPtrOutputWithContext(ctx context.Context) ServiceManagedIdentityResponseIdentityPtrOutput {
	return o
}

func (o ServiceManagedIdentityResponseIdentityPtrOutput) Elem() ServiceManagedIdentityResponseIdentityOutput {
	return o.ApplyT(func(v *ServiceManagedIdentityResponseIdentity) ServiceManagedIdentityResponseIdentity {
		if v != nil {
			return *v
		}
		var ret ServiceManagedIdentityResponseIdentity
		return ret
	}).(ServiceManagedIdentityResponseIdentityOutput)
}

func (o ServiceManagedIdentityResponseIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceManagedIdentityResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ServiceManagedIdentityResponseIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceManagedIdentityResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ServiceManagedIdentityResponseIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceManagedIdentityResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ServiceManagedIdentityResponseIdentityPtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ServiceManagedIdentityResponseIdentity) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ServiceOciArtifactEntry struct {
	Digest      *string `pulumi:"digest"`
	ImageName   *string `pulumi:"imageName"`
	LoginServer *string `pulumi:"loginServer"`
}





type ServiceOciArtifactEntryInput interface {
	pulumi.Input

	ToServiceOciArtifactEntryOutput() ServiceOciArtifactEntryOutput
	ToServiceOciArtifactEntryOutputWithContext(context.Context) ServiceOciArtifactEntryOutput
}

type ServiceOciArtifactEntryArgs struct {
	Digest      pulumi.StringPtrInput `pulumi:"digest"`
	ImageName   pulumi.StringPtrInput `pulumi:"imageName"`
	LoginServer pulumi.StringPtrInput `pulumi:"loginServer"`
}

func (ServiceOciArtifactEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceOciArtifactEntry)(nil)).Elem()
}

func (i ServiceOciArtifactEntryArgs) ToServiceOciArtifactEntryOutput() ServiceOciArtifactEntryOutput {
	return i.ToServiceOciArtifactEntryOutputWithContext(context.Background())
}

func (i ServiceOciArtifactEntryArgs) ToServiceOciArtifactEntryOutputWithContext(ctx context.Context) ServiceOciArtifactEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOciArtifactEntryOutput)
}





type ServiceOciArtifactEntryArrayInput interface {
	pulumi.Input

	ToServiceOciArtifactEntryArrayOutput() ServiceOciArtifactEntryArrayOutput
	ToServiceOciArtifactEntryArrayOutputWithContext(context.Context) ServiceOciArtifactEntryArrayOutput
}

type ServiceOciArtifactEntryArray []ServiceOciArtifactEntryInput

func (ServiceOciArtifactEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceOciArtifactEntry)(nil)).Elem()
}

func (i ServiceOciArtifactEntryArray) ToServiceOciArtifactEntryArrayOutput() ServiceOciArtifactEntryArrayOutput {
	return i.ToServiceOciArtifactEntryArrayOutputWithContext(context.Background())
}

func (i ServiceOciArtifactEntryArray) ToServiceOciArtifactEntryArrayOutputWithContext(ctx context.Context) ServiceOciArtifactEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOciArtifactEntryArrayOutput)
}

type ServiceOciArtifactEntryOutput struct{ *pulumi.OutputState }

func (ServiceOciArtifactEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceOciArtifactEntry)(nil)).Elem()
}

func (o ServiceOciArtifactEntryOutput) ToServiceOciArtifactEntryOutput() ServiceOciArtifactEntryOutput {
	return o
}

func (o ServiceOciArtifactEntryOutput) ToServiceOciArtifactEntryOutputWithContext(ctx context.Context) ServiceOciArtifactEntryOutput {
	return o
}

func (o ServiceOciArtifactEntryOutput) Digest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceOciArtifactEntry) *string { return v.Digest }).(pulumi.StringPtrOutput)
}

func (o ServiceOciArtifactEntryOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceOciArtifactEntry) *string { return v.ImageName }).(pulumi.StringPtrOutput)
}

func (o ServiceOciArtifactEntryOutput) LoginServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceOciArtifactEntry) *string { return v.LoginServer }).(pulumi.StringPtrOutput)
}

type ServiceOciArtifactEntryArrayOutput struct{ *pulumi.OutputState }

func (ServiceOciArtifactEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceOciArtifactEntry)(nil)).Elem()
}

func (o ServiceOciArtifactEntryArrayOutput) ToServiceOciArtifactEntryArrayOutput() ServiceOciArtifactEntryArrayOutput {
	return o
}

func (o ServiceOciArtifactEntryArrayOutput) ToServiceOciArtifactEntryArrayOutputWithContext(ctx context.Context) ServiceOciArtifactEntryArrayOutput {
	return o
}

func (o ServiceOciArtifactEntryArrayOutput) Index(i pulumi.IntInput) ServiceOciArtifactEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceOciArtifactEntry {
		return vs[0].([]ServiceOciArtifactEntry)[vs[1].(int)]
	}).(ServiceOciArtifactEntryOutput)
}

type ServiceOciArtifactEntryResponse struct {
	Digest      *string `pulumi:"digest"`
	ImageName   *string `pulumi:"imageName"`
	LoginServer *string `pulumi:"loginServer"`
}

type ServiceOciArtifactEntryResponseOutput struct{ *pulumi.OutputState }

func (ServiceOciArtifactEntryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceOciArtifactEntryResponse)(nil)).Elem()
}

func (o ServiceOciArtifactEntryResponseOutput) ToServiceOciArtifactEntryResponseOutput() ServiceOciArtifactEntryResponseOutput {
	return o
}

func (o ServiceOciArtifactEntryResponseOutput) ToServiceOciArtifactEntryResponseOutputWithContext(ctx context.Context) ServiceOciArtifactEntryResponseOutput {
	return o
}

func (o ServiceOciArtifactEntryResponseOutput) Digest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceOciArtifactEntryResponse) *string { return v.Digest }).(pulumi.StringPtrOutput)
}

func (o ServiceOciArtifactEntryResponseOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceOciArtifactEntryResponse) *string { return v.ImageName }).(pulumi.StringPtrOutput)
}

func (o ServiceOciArtifactEntryResponseOutput) LoginServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceOciArtifactEntryResponse) *string { return v.LoginServer }).(pulumi.StringPtrOutput)
}

type ServiceOciArtifactEntryResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceOciArtifactEntryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceOciArtifactEntryResponse)(nil)).Elem()
}

func (o ServiceOciArtifactEntryResponseArrayOutput) ToServiceOciArtifactEntryResponseArrayOutput() ServiceOciArtifactEntryResponseArrayOutput {
	return o
}

func (o ServiceOciArtifactEntryResponseArrayOutput) ToServiceOciArtifactEntryResponseArrayOutputWithContext(ctx context.Context) ServiceOciArtifactEntryResponseArrayOutput {
	return o
}

func (o ServiceOciArtifactEntryResponseArrayOutput) Index(i pulumi.IntInput) ServiceOciArtifactEntryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceOciArtifactEntryResponse {
		return vs[0].([]ServiceOciArtifactEntryResponse)[vs[1].(int)]
	}).(ServiceOciArtifactEntryResponseOutput)
}

type ServicesProperties struct {
	AccessPolicies              []ServiceAccessPolicyEntry              `pulumi:"accessPolicies"`
	AcrConfiguration            *ServiceAcrConfigurationInfo            `pulumi:"acrConfiguration"`
	AuthenticationConfiguration *ServiceAuthenticationConfigurationInfo `pulumi:"authenticationConfiguration"`
	CorsConfiguration           *ServiceCorsConfigurationInfo           `pulumi:"corsConfiguration"`
	CosmosDbConfiguration       *ServiceCosmosDbConfigurationInfo       `pulumi:"cosmosDbConfiguration"`
	ExportConfiguration         *ServiceExportConfigurationInfo         `pulumi:"exportConfiguration"`
	PrivateEndpointConnections  []PrivateEndpointConnectionType         `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess         *string                                 `pulumi:"publicNetworkAccess"`
}





type ServicesPropertiesInput interface {
	pulumi.Input

	ToServicesPropertiesOutput() ServicesPropertiesOutput
	ToServicesPropertiesOutputWithContext(context.Context) ServicesPropertiesOutput
}

type ServicesPropertiesArgs struct {
	AccessPolicies              ServiceAccessPolicyEntryArrayInput             `pulumi:"accessPolicies"`
	AcrConfiguration            ServiceAcrConfigurationInfoPtrInput            `pulumi:"acrConfiguration"`
	AuthenticationConfiguration ServiceAuthenticationConfigurationInfoPtrInput `pulumi:"authenticationConfiguration"`
	CorsConfiguration           ServiceCorsConfigurationInfoPtrInput           `pulumi:"corsConfiguration"`
	CosmosDbConfiguration       ServiceCosmosDbConfigurationInfoPtrInput       `pulumi:"cosmosDbConfiguration"`
	ExportConfiguration         ServiceExportConfigurationInfoPtrInput         `pulumi:"exportConfiguration"`
	PrivateEndpointConnections  PrivateEndpointConnectionTypeArrayInput        `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess         pulumi.StringPtrInput                          `pulumi:"publicNetworkAccess"`
}

func (ServicesPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicesProperties)(nil)).Elem()
}

func (i ServicesPropertiesArgs) ToServicesPropertiesOutput() ServicesPropertiesOutput {
	return i.ToServicesPropertiesOutputWithContext(context.Background())
}

func (i ServicesPropertiesArgs) ToServicesPropertiesOutputWithContext(ctx context.Context) ServicesPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesPropertiesOutput)
}

func (i ServicesPropertiesArgs) ToServicesPropertiesPtrOutput() ServicesPropertiesPtrOutput {
	return i.ToServicesPropertiesPtrOutputWithContext(context.Background())
}

func (i ServicesPropertiesArgs) ToServicesPropertiesPtrOutputWithContext(ctx context.Context) ServicesPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesPropertiesOutput).ToServicesPropertiesPtrOutputWithContext(ctx)
}









type ServicesPropertiesPtrInput interface {
	pulumi.Input

	ToServicesPropertiesPtrOutput() ServicesPropertiesPtrOutput
	ToServicesPropertiesPtrOutputWithContext(context.Context) ServicesPropertiesPtrOutput
}

type servicesPropertiesPtrType ServicesPropertiesArgs

func ServicesPropertiesPtr(v *ServicesPropertiesArgs) ServicesPropertiesPtrInput {
	return (*servicesPropertiesPtrType)(v)
}

func (*servicesPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicesProperties)(nil)).Elem()
}

func (i *servicesPropertiesPtrType) ToServicesPropertiesPtrOutput() ServicesPropertiesPtrOutput {
	return i.ToServicesPropertiesPtrOutputWithContext(context.Background())
}

func (i *servicesPropertiesPtrType) ToServicesPropertiesPtrOutputWithContext(ctx context.Context) ServicesPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesPropertiesPtrOutput)
}

type ServicesPropertiesOutput struct{ *pulumi.OutputState }

func (ServicesPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicesProperties)(nil)).Elem()
}

func (o ServicesPropertiesOutput) ToServicesPropertiesOutput() ServicesPropertiesOutput {
	return o
}

func (o ServicesPropertiesOutput) ToServicesPropertiesOutputWithContext(ctx context.Context) ServicesPropertiesOutput {
	return o
}

func (o ServicesPropertiesOutput) ToServicesPropertiesPtrOutput() ServicesPropertiesPtrOutput {
	return o.ToServicesPropertiesPtrOutputWithContext(context.Background())
}

func (o ServicesPropertiesOutput) ToServicesPropertiesPtrOutputWithContext(ctx context.Context) ServicesPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServicesProperties) *ServicesProperties {
		return &v
	}).(ServicesPropertiesPtrOutput)
}

func (o ServicesPropertiesOutput) AccessPolicies() ServiceAccessPolicyEntryArrayOutput {
	return o.ApplyT(func(v ServicesProperties) []ServiceAccessPolicyEntry { return v.AccessPolicies }).(ServiceAccessPolicyEntryArrayOutput)
}

func (o ServicesPropertiesOutput) AcrConfiguration() ServiceAcrConfigurationInfoPtrOutput {
	return o.ApplyT(func(v ServicesProperties) *ServiceAcrConfigurationInfo { return v.AcrConfiguration }).(ServiceAcrConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesOutput) AuthenticationConfiguration() ServiceAuthenticationConfigurationInfoPtrOutput {
	return o.ApplyT(func(v ServicesProperties) *ServiceAuthenticationConfigurationInfo {
		return v.AuthenticationConfiguration
	}).(ServiceAuthenticationConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesOutput) CorsConfiguration() ServiceCorsConfigurationInfoPtrOutput {
	return o.ApplyT(func(v ServicesProperties) *ServiceCorsConfigurationInfo { return v.CorsConfiguration }).(ServiceCorsConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesOutput) CosmosDbConfiguration() ServiceCosmosDbConfigurationInfoPtrOutput {
	return o.ApplyT(func(v ServicesProperties) *ServiceCosmosDbConfigurationInfo { return v.CosmosDbConfiguration }).(ServiceCosmosDbConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesOutput) ExportConfiguration() ServiceExportConfigurationInfoPtrOutput {
	return o.ApplyT(func(v ServicesProperties) *ServiceExportConfigurationInfo { return v.ExportConfiguration }).(ServiceExportConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesOutput) PrivateEndpointConnections() PrivateEndpointConnectionTypeArrayOutput {
	return o.ApplyT(func(v ServicesProperties) []PrivateEndpointConnectionType { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionTypeArrayOutput)
}

func (o ServicesPropertiesOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicesProperties) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

type ServicesPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ServicesPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicesProperties)(nil)).Elem()
}

func (o ServicesPropertiesPtrOutput) ToServicesPropertiesPtrOutput() ServicesPropertiesPtrOutput {
	return o
}

func (o ServicesPropertiesPtrOutput) ToServicesPropertiesPtrOutputWithContext(ctx context.Context) ServicesPropertiesPtrOutput {
	return o
}

func (o ServicesPropertiesPtrOutput) Elem() ServicesPropertiesOutput {
	return o.ApplyT(func(v *ServicesProperties) ServicesProperties {
		if v != nil {
			return *v
		}
		var ret ServicesProperties
		return ret
	}).(ServicesPropertiesOutput)
}

func (o ServicesPropertiesPtrOutput) AccessPolicies() ServiceAccessPolicyEntryArrayOutput {
	return o.ApplyT(func(v *ServicesProperties) []ServiceAccessPolicyEntry {
		if v == nil {
			return nil
		}
		return v.AccessPolicies
	}).(ServiceAccessPolicyEntryArrayOutput)
}

func (o ServicesPropertiesPtrOutput) AcrConfiguration() ServiceAcrConfigurationInfoPtrOutput {
	return o.ApplyT(func(v *ServicesProperties) *ServiceAcrConfigurationInfo {
		if v == nil {
			return nil
		}
		return v.AcrConfiguration
	}).(ServiceAcrConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesPtrOutput) AuthenticationConfiguration() ServiceAuthenticationConfigurationInfoPtrOutput {
	return o.ApplyT(func(v *ServicesProperties) *ServiceAuthenticationConfigurationInfo {
		if v == nil {
			return nil
		}
		return v.AuthenticationConfiguration
	}).(ServiceAuthenticationConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesPtrOutput) CorsConfiguration() ServiceCorsConfigurationInfoPtrOutput {
	return o.ApplyT(func(v *ServicesProperties) *ServiceCorsConfigurationInfo {
		if v == nil {
			return nil
		}
		return v.CorsConfiguration
	}).(ServiceCorsConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesPtrOutput) CosmosDbConfiguration() ServiceCosmosDbConfigurationInfoPtrOutput {
	return o.ApplyT(func(v *ServicesProperties) *ServiceCosmosDbConfigurationInfo {
		if v == nil {
			return nil
		}
		return v.CosmosDbConfiguration
	}).(ServiceCosmosDbConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesPtrOutput) ExportConfiguration() ServiceExportConfigurationInfoPtrOutput {
	return o.ApplyT(func(v *ServicesProperties) *ServiceExportConfigurationInfo {
		if v == nil {
			return nil
		}
		return v.ExportConfiguration
	}).(ServiceExportConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesPtrOutput) PrivateEndpointConnections() PrivateEndpointConnectionTypeArrayOutput {
	return o.ApplyT(func(v *ServicesProperties) []PrivateEndpointConnectionType {
		if v == nil {
			return nil
		}
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionTypeArrayOutput)
}

func (o ServicesPropertiesPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

type ServicesPropertiesResponse struct {
	AccessPolicies              []ServiceAccessPolicyEntryResponse              `pulumi:"accessPolicies"`
	AcrConfiguration            *ServiceAcrConfigurationInfoResponse            `pulumi:"acrConfiguration"`
	AuthenticationConfiguration *ServiceAuthenticationConfigurationInfoResponse `pulumi:"authenticationConfiguration"`
	CorsConfiguration           *ServiceCorsConfigurationInfoResponse           `pulumi:"corsConfiguration"`
	CosmosDbConfiguration       *ServiceCosmosDbConfigurationInfoResponse       `pulumi:"cosmosDbConfiguration"`
	ExportConfiguration         *ServiceExportConfigurationInfoResponse         `pulumi:"exportConfiguration"`
	PrivateEndpointConnections  []PrivateEndpointConnectionResponse             `pulumi:"privateEndpointConnections"`
	ProvisioningState           string                                          `pulumi:"provisioningState"`
	PublicNetworkAccess         *string                                         `pulumi:"publicNetworkAccess"`
}

type ServicesPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ServicesPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicesPropertiesResponse)(nil)).Elem()
}

func (o ServicesPropertiesResponseOutput) ToServicesPropertiesResponseOutput() ServicesPropertiesResponseOutput {
	return o
}

func (o ServicesPropertiesResponseOutput) ToServicesPropertiesResponseOutputWithContext(ctx context.Context) ServicesPropertiesResponseOutput {
	return o
}

func (o ServicesPropertiesResponseOutput) AccessPolicies() ServiceAccessPolicyEntryResponseArrayOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) []ServiceAccessPolicyEntryResponse { return v.AccessPolicies }).(ServiceAccessPolicyEntryResponseArrayOutput)
}

func (o ServicesPropertiesResponseOutput) AcrConfiguration() ServiceAcrConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) *ServiceAcrConfigurationInfoResponse { return v.AcrConfiguration }).(ServiceAcrConfigurationInfoResponsePtrOutput)
}

func (o ServicesPropertiesResponseOutput) AuthenticationConfiguration() ServiceAuthenticationConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) *ServiceAuthenticationConfigurationInfoResponse {
		return v.AuthenticationConfiguration
	}).(ServiceAuthenticationConfigurationInfoResponsePtrOutput)
}

func (o ServicesPropertiesResponseOutput) CorsConfiguration() ServiceCorsConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) *ServiceCorsConfigurationInfoResponse { return v.CorsConfiguration }).(ServiceCorsConfigurationInfoResponsePtrOutput)
}

func (o ServicesPropertiesResponseOutput) CosmosDbConfiguration() ServiceCosmosDbConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) *ServiceCosmosDbConfigurationInfoResponse {
		return v.CosmosDbConfiguration
	}).(ServiceCosmosDbConfigurationInfoResponsePtrOutput)
}

func (o ServicesPropertiesResponseOutput) ExportConfiguration() ServiceExportConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) *ServiceExportConfigurationInfoResponse {
		return v.ExportConfiguration
	}).(ServiceExportConfigurationInfoResponsePtrOutput)
}

func (o ServicesPropertiesResponseOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o ServicesPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ServicesPropertiesResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

type ServicesResourceIdentity struct {
	Type *string `pulumi:"type"`
}





type ServicesResourceIdentityInput interface {
	pulumi.Input

	ToServicesResourceIdentityOutput() ServicesResourceIdentityOutput
	ToServicesResourceIdentityOutputWithContext(context.Context) ServicesResourceIdentityOutput
}

type ServicesResourceIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ServicesResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicesResourceIdentity)(nil)).Elem()
}

func (i ServicesResourceIdentityArgs) ToServicesResourceIdentityOutput() ServicesResourceIdentityOutput {
	return i.ToServicesResourceIdentityOutputWithContext(context.Background())
}

func (i ServicesResourceIdentityArgs) ToServicesResourceIdentityOutputWithContext(ctx context.Context) ServicesResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesResourceIdentityOutput)
}

func (i ServicesResourceIdentityArgs) ToServicesResourceIdentityPtrOutput() ServicesResourceIdentityPtrOutput {
	return i.ToServicesResourceIdentityPtrOutputWithContext(context.Background())
}

func (i ServicesResourceIdentityArgs) ToServicesResourceIdentityPtrOutputWithContext(ctx context.Context) ServicesResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesResourceIdentityOutput).ToServicesResourceIdentityPtrOutputWithContext(ctx)
}









type ServicesResourceIdentityPtrInput interface {
	pulumi.Input

	ToServicesResourceIdentityPtrOutput() ServicesResourceIdentityPtrOutput
	ToServicesResourceIdentityPtrOutputWithContext(context.Context) ServicesResourceIdentityPtrOutput
}

type servicesResourceIdentityPtrType ServicesResourceIdentityArgs

func ServicesResourceIdentityPtr(v *ServicesResourceIdentityArgs) ServicesResourceIdentityPtrInput {
	return (*servicesResourceIdentityPtrType)(v)
}

func (*servicesResourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicesResourceIdentity)(nil)).Elem()
}

func (i *servicesResourceIdentityPtrType) ToServicesResourceIdentityPtrOutput() ServicesResourceIdentityPtrOutput {
	return i.ToServicesResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *servicesResourceIdentityPtrType) ToServicesResourceIdentityPtrOutputWithContext(ctx context.Context) ServicesResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesResourceIdentityPtrOutput)
}

type ServicesResourceIdentityOutput struct{ *pulumi.OutputState }

func (ServicesResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicesResourceIdentity)(nil)).Elem()
}

func (o ServicesResourceIdentityOutput) ToServicesResourceIdentityOutput() ServicesResourceIdentityOutput {
	return o
}

func (o ServicesResourceIdentityOutput) ToServicesResourceIdentityOutputWithContext(ctx context.Context) ServicesResourceIdentityOutput {
	return o
}

func (o ServicesResourceIdentityOutput) ToServicesResourceIdentityPtrOutput() ServicesResourceIdentityPtrOutput {
	return o.ToServicesResourceIdentityPtrOutputWithContext(context.Background())
}

func (o ServicesResourceIdentityOutput) ToServicesResourceIdentityPtrOutputWithContext(ctx context.Context) ServicesResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServicesResourceIdentity) *ServicesResourceIdentity {
		return &v
	}).(ServicesResourceIdentityPtrOutput)
}

func (o ServicesResourceIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicesResourceIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ServicesResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ServicesResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicesResourceIdentity)(nil)).Elem()
}

func (o ServicesResourceIdentityPtrOutput) ToServicesResourceIdentityPtrOutput() ServicesResourceIdentityPtrOutput {
	return o
}

func (o ServicesResourceIdentityPtrOutput) ToServicesResourceIdentityPtrOutputWithContext(ctx context.Context) ServicesResourceIdentityPtrOutput {
	return o
}

func (o ServicesResourceIdentityPtrOutput) Elem() ServicesResourceIdentityOutput {
	return o.ApplyT(func(v *ServicesResourceIdentity) ServicesResourceIdentity {
		if v != nil {
			return *v
		}
		var ret ServicesResourceIdentity
		return ret
	}).(ServicesResourceIdentityOutput)
}

func (o ServicesResourceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ServicesResourceResponseIdentity struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type ServicesResourceResponseIdentityOutput struct{ *pulumi.OutputState }

func (ServicesResourceResponseIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicesResourceResponseIdentity)(nil)).Elem()
}

func (o ServicesResourceResponseIdentityOutput) ToServicesResourceResponseIdentityOutput() ServicesResourceResponseIdentityOutput {
	return o
}

func (o ServicesResourceResponseIdentityOutput) ToServicesResourceResponseIdentityOutputWithContext(ctx context.Context) ServicesResourceResponseIdentityOutput {
	return o
}

func (o ServicesResourceResponseIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ServicesResourceResponseIdentity) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ServicesResourceResponseIdentityOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ServicesResourceResponseIdentity) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ServicesResourceResponseIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicesResourceResponseIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ServicesResourceResponseIdentityPtrOutput struct{ *pulumi.OutputState }

func (ServicesResourceResponseIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicesResourceResponseIdentity)(nil)).Elem()
}

func (o ServicesResourceResponseIdentityPtrOutput) ToServicesResourceResponseIdentityPtrOutput() ServicesResourceResponseIdentityPtrOutput {
	return o
}

func (o ServicesResourceResponseIdentityPtrOutput) ToServicesResourceResponseIdentityPtrOutputWithContext(ctx context.Context) ServicesResourceResponseIdentityPtrOutput {
	return o
}

func (o ServicesResourceResponseIdentityPtrOutput) Elem() ServicesResourceResponseIdentityOutput {
	return o.ApplyT(func(v *ServicesResourceResponseIdentity) ServicesResourceResponseIdentity {
		if v != nil {
			return *v
		}
		var ret ServicesResourceResponseIdentity
		return ret
	}).(ServicesResourceResponseIdentityOutput)
}

func (o ServicesResourceResponseIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesResourceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ServicesResourceResponseIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesResourceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ServicesResourceResponseIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesResourceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type UserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

type WorkspaceResponseProperties struct {
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	PublicNetworkAccess        string                              `pulumi:"publicNetworkAccess"`
}

type WorkspaceResponsePropertiesOutput struct{ *pulumi.OutputState }

func (WorkspaceResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceResponseProperties)(nil)).Elem()
}

func (o WorkspaceResponsePropertiesOutput) ToWorkspaceResponsePropertiesOutput() WorkspaceResponsePropertiesOutput {
	return o
}

func (o WorkspaceResponsePropertiesOutput) ToWorkspaceResponsePropertiesOutputWithContext(ctx context.Context) WorkspaceResponsePropertiesOutput {
	return o
}

func (o WorkspaceResponsePropertiesOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v WorkspaceResponseProperties) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o WorkspaceResponsePropertiesOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceResponseProperties) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WorkspaceResponsePropertiesOutput) PublicNetworkAccess() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceResponseProperties) string { return v.PublicNetworkAccess }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DicomServiceAuthenticationConfigurationResponseOutput{})
	pulumi.RegisterOutputType(DicomServiceAuthenticationConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(FhirServiceAccessPolicyEntryOutput{})
	pulumi.RegisterOutputType(FhirServiceAccessPolicyEntryArrayOutput{})
	pulumi.RegisterOutputType(FhirServiceAccessPolicyEntryResponseOutput{})
	pulumi.RegisterOutputType(FhirServiceAccessPolicyEntryResponseArrayOutput{})
	pulumi.RegisterOutputType(FhirServiceAcrConfigurationOutput{})
	pulumi.RegisterOutputType(FhirServiceAcrConfigurationPtrOutput{})
	pulumi.RegisterOutputType(FhirServiceAcrConfigurationResponseOutput{})
	pulumi.RegisterOutputType(FhirServiceAcrConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(FhirServiceAuthenticationConfigurationOutput{})
	pulumi.RegisterOutputType(FhirServiceAuthenticationConfigurationPtrOutput{})
	pulumi.RegisterOutputType(FhirServiceAuthenticationConfigurationResponseOutput{})
	pulumi.RegisterOutputType(FhirServiceAuthenticationConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(FhirServiceCorsConfigurationOutput{})
	pulumi.RegisterOutputType(FhirServiceCorsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(FhirServiceCorsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(FhirServiceCorsConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(FhirServiceExportConfigurationOutput{})
	pulumi.RegisterOutputType(FhirServiceExportConfigurationPtrOutput{})
	pulumi.RegisterOutputType(FhirServiceExportConfigurationResponseOutput{})
	pulumi.RegisterOutputType(FhirServiceExportConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(IotEventHubIngestionEndpointConfigurationOutput{})
	pulumi.RegisterOutputType(IotEventHubIngestionEndpointConfigurationPtrOutput{})
	pulumi.RegisterOutputType(IotEventHubIngestionEndpointConfigurationResponseOutput{})
	pulumi.RegisterOutputType(IotEventHubIngestionEndpointConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(IotMappingPropertiesOutput{})
	pulumi.RegisterOutputType(IotMappingPropertiesPtrOutput{})
	pulumi.RegisterOutputType(IotMappingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(IotMappingPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(ResourceVersionPolicyConfigurationOutput{})
	pulumi.RegisterOutputType(ResourceVersionPolicyConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ResourceVersionPolicyConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ResourceVersionPolicyConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceAccessPolicyEntryOutput{})
	pulumi.RegisterOutputType(ServiceAccessPolicyEntryArrayOutput{})
	pulumi.RegisterOutputType(ServiceAccessPolicyEntryResponseOutput{})
	pulumi.RegisterOutputType(ServiceAccessPolicyEntryResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceAcrConfigurationInfoOutput{})
	pulumi.RegisterOutputType(ServiceAcrConfigurationInfoPtrOutput{})
	pulumi.RegisterOutputType(ServiceAcrConfigurationInfoResponseOutput{})
	pulumi.RegisterOutputType(ServiceAcrConfigurationInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceAuthenticationConfigurationInfoOutput{})
	pulumi.RegisterOutputType(ServiceAuthenticationConfigurationInfoPtrOutput{})
	pulumi.RegisterOutputType(ServiceAuthenticationConfigurationInfoResponseOutput{})
	pulumi.RegisterOutputType(ServiceAuthenticationConfigurationInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceCorsConfigurationInfoOutput{})
	pulumi.RegisterOutputType(ServiceCorsConfigurationInfoPtrOutput{})
	pulumi.RegisterOutputType(ServiceCorsConfigurationInfoResponseOutput{})
	pulumi.RegisterOutputType(ServiceCorsConfigurationInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceCosmosDbConfigurationInfoOutput{})
	pulumi.RegisterOutputType(ServiceCosmosDbConfigurationInfoPtrOutput{})
	pulumi.RegisterOutputType(ServiceCosmosDbConfigurationInfoResponseOutput{})
	pulumi.RegisterOutputType(ServiceCosmosDbConfigurationInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceExportConfigurationInfoOutput{})
	pulumi.RegisterOutputType(ServiceExportConfigurationInfoPtrOutput{})
	pulumi.RegisterOutputType(ServiceExportConfigurationInfoResponseOutput{})
	pulumi.RegisterOutputType(ServiceExportConfigurationInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceManagedIdentityIdentityOutput{})
	pulumi.RegisterOutputType(ServiceManagedIdentityIdentityPtrOutput{})
	pulumi.RegisterOutputType(ServiceManagedIdentityResponseIdentityOutput{})
	pulumi.RegisterOutputType(ServiceManagedIdentityResponseIdentityPtrOutput{})
	pulumi.RegisterOutputType(ServiceOciArtifactEntryOutput{})
	pulumi.RegisterOutputType(ServiceOciArtifactEntryArrayOutput{})
	pulumi.RegisterOutputType(ServiceOciArtifactEntryResponseOutput{})
	pulumi.RegisterOutputType(ServiceOciArtifactEntryResponseArrayOutput{})
	pulumi.RegisterOutputType(ServicesPropertiesOutput{})
	pulumi.RegisterOutputType(ServicesPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ServicesPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ServicesResourceIdentityOutput{})
	pulumi.RegisterOutputType(ServicesResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ServicesResourceResponseIdentityOutput{})
	pulumi.RegisterOutputType(ServicesResourceResponseIdentityPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(WorkspaceResponsePropertiesOutput{})
}

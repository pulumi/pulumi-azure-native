


package appplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiPortalCustomDomainProperties struct {
	Thumbprint *string `pulumi:"thumbprint"`
}





type ApiPortalCustomDomainPropertiesInput interface {
	pulumi.Input

	ToApiPortalCustomDomainPropertiesOutput() ApiPortalCustomDomainPropertiesOutput
	ToApiPortalCustomDomainPropertiesOutputWithContext(context.Context) ApiPortalCustomDomainPropertiesOutput
}

type ApiPortalCustomDomainPropertiesArgs struct {
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (ApiPortalCustomDomainPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiPortalCustomDomainProperties)(nil)).Elem()
}

func (i ApiPortalCustomDomainPropertiesArgs) ToApiPortalCustomDomainPropertiesOutput() ApiPortalCustomDomainPropertiesOutput {
	return i.ToApiPortalCustomDomainPropertiesOutputWithContext(context.Background())
}

func (i ApiPortalCustomDomainPropertiesArgs) ToApiPortalCustomDomainPropertiesOutputWithContext(ctx context.Context) ApiPortalCustomDomainPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPortalCustomDomainPropertiesOutput)
}

func (i ApiPortalCustomDomainPropertiesArgs) ToApiPortalCustomDomainPropertiesPtrOutput() ApiPortalCustomDomainPropertiesPtrOutput {
	return i.ToApiPortalCustomDomainPropertiesPtrOutputWithContext(context.Background())
}

func (i ApiPortalCustomDomainPropertiesArgs) ToApiPortalCustomDomainPropertiesPtrOutputWithContext(ctx context.Context) ApiPortalCustomDomainPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPortalCustomDomainPropertiesOutput).ToApiPortalCustomDomainPropertiesPtrOutputWithContext(ctx)
}









type ApiPortalCustomDomainPropertiesPtrInput interface {
	pulumi.Input

	ToApiPortalCustomDomainPropertiesPtrOutput() ApiPortalCustomDomainPropertiesPtrOutput
	ToApiPortalCustomDomainPropertiesPtrOutputWithContext(context.Context) ApiPortalCustomDomainPropertiesPtrOutput
}

type apiPortalCustomDomainPropertiesPtrType ApiPortalCustomDomainPropertiesArgs

func ApiPortalCustomDomainPropertiesPtr(v *ApiPortalCustomDomainPropertiesArgs) ApiPortalCustomDomainPropertiesPtrInput {
	return (*apiPortalCustomDomainPropertiesPtrType)(v)
}

func (*apiPortalCustomDomainPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPortalCustomDomainProperties)(nil)).Elem()
}

func (i *apiPortalCustomDomainPropertiesPtrType) ToApiPortalCustomDomainPropertiesPtrOutput() ApiPortalCustomDomainPropertiesPtrOutput {
	return i.ToApiPortalCustomDomainPropertiesPtrOutputWithContext(context.Background())
}

func (i *apiPortalCustomDomainPropertiesPtrType) ToApiPortalCustomDomainPropertiesPtrOutputWithContext(ctx context.Context) ApiPortalCustomDomainPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPortalCustomDomainPropertiesPtrOutput)
}

type ApiPortalCustomDomainPropertiesOutput struct{ *pulumi.OutputState }

func (ApiPortalCustomDomainPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiPortalCustomDomainProperties)(nil)).Elem()
}

func (o ApiPortalCustomDomainPropertiesOutput) ToApiPortalCustomDomainPropertiesOutput() ApiPortalCustomDomainPropertiesOutput {
	return o
}

func (o ApiPortalCustomDomainPropertiesOutput) ToApiPortalCustomDomainPropertiesOutputWithContext(ctx context.Context) ApiPortalCustomDomainPropertiesOutput {
	return o
}

func (o ApiPortalCustomDomainPropertiesOutput) ToApiPortalCustomDomainPropertiesPtrOutput() ApiPortalCustomDomainPropertiesPtrOutput {
	return o.ToApiPortalCustomDomainPropertiesPtrOutputWithContext(context.Background())
}

func (o ApiPortalCustomDomainPropertiesOutput) ToApiPortalCustomDomainPropertiesPtrOutputWithContext(ctx context.Context) ApiPortalCustomDomainPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiPortalCustomDomainProperties) *ApiPortalCustomDomainProperties {
		return &v
	}).(ApiPortalCustomDomainPropertiesPtrOutput)
}

func (o ApiPortalCustomDomainPropertiesOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPortalCustomDomainProperties) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type ApiPortalCustomDomainPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ApiPortalCustomDomainPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPortalCustomDomainProperties)(nil)).Elem()
}

func (o ApiPortalCustomDomainPropertiesPtrOutput) ToApiPortalCustomDomainPropertiesPtrOutput() ApiPortalCustomDomainPropertiesPtrOutput {
	return o
}

func (o ApiPortalCustomDomainPropertiesPtrOutput) ToApiPortalCustomDomainPropertiesPtrOutputWithContext(ctx context.Context) ApiPortalCustomDomainPropertiesPtrOutput {
	return o
}

func (o ApiPortalCustomDomainPropertiesPtrOutput) Elem() ApiPortalCustomDomainPropertiesOutput {
	return o.ApplyT(func(v *ApiPortalCustomDomainProperties) ApiPortalCustomDomainProperties {
		if v != nil {
			return *v
		}
		var ret ApiPortalCustomDomainProperties
		return ret
	}).(ApiPortalCustomDomainPropertiesOutput)
}

func (o ApiPortalCustomDomainPropertiesPtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPortalCustomDomainProperties) *string {
		if v == nil {
			return nil
		}
		return v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

type ApiPortalCustomDomainPropertiesResponse struct {
	Thumbprint *string `pulumi:"thumbprint"`
}

type ApiPortalCustomDomainPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ApiPortalCustomDomainPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiPortalCustomDomainPropertiesResponse)(nil)).Elem()
}

func (o ApiPortalCustomDomainPropertiesResponseOutput) ToApiPortalCustomDomainPropertiesResponseOutput() ApiPortalCustomDomainPropertiesResponseOutput {
	return o
}

func (o ApiPortalCustomDomainPropertiesResponseOutput) ToApiPortalCustomDomainPropertiesResponseOutputWithContext(ctx context.Context) ApiPortalCustomDomainPropertiesResponseOutput {
	return o
}

func (o ApiPortalCustomDomainPropertiesResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPortalCustomDomainPropertiesResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type ApiPortalInstanceResponse struct {
	Name   string `pulumi:"name"`
	Status string `pulumi:"status"`
}

type ApiPortalInstanceResponseOutput struct{ *pulumi.OutputState }

func (ApiPortalInstanceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiPortalInstanceResponse)(nil)).Elem()
}

func (o ApiPortalInstanceResponseOutput) ToApiPortalInstanceResponseOutput() ApiPortalInstanceResponseOutput {
	return o
}

func (o ApiPortalInstanceResponseOutput) ToApiPortalInstanceResponseOutputWithContext(ctx context.Context) ApiPortalInstanceResponseOutput {
	return o
}

func (o ApiPortalInstanceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApiPortalInstanceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ApiPortalInstanceResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ApiPortalInstanceResponse) string { return v.Status }).(pulumi.StringOutput)
}

type ApiPortalInstanceResponseArrayOutput struct{ *pulumi.OutputState }

func (ApiPortalInstanceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiPortalInstanceResponse)(nil)).Elem()
}

func (o ApiPortalInstanceResponseArrayOutput) ToApiPortalInstanceResponseArrayOutput() ApiPortalInstanceResponseArrayOutput {
	return o
}

func (o ApiPortalInstanceResponseArrayOutput) ToApiPortalInstanceResponseArrayOutputWithContext(ctx context.Context) ApiPortalInstanceResponseArrayOutput {
	return o
}

func (o ApiPortalInstanceResponseArrayOutput) Index(i pulumi.IntInput) ApiPortalInstanceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiPortalInstanceResponse {
		return vs[0].([]ApiPortalInstanceResponse)[vs[1].(int)]
	}).(ApiPortalInstanceResponseOutput)
}

type ApiPortalProperties struct {
	GatewayIds    []string       `pulumi:"gatewayIds"`
	HttpsOnly     *bool          `pulumi:"httpsOnly"`
	Public        *bool          `pulumi:"public"`
	SourceUrls    []string       `pulumi:"sourceUrls"`
	SsoProperties *SsoProperties `pulumi:"ssoProperties"`
}


func (val *ApiPortalProperties) Defaults() *ApiPortalProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HttpsOnly) {
		httpsOnly_ := false
		tmp.HttpsOnly = &httpsOnly_
	}
	if isZero(tmp.Public) {
		public_ := false
		tmp.Public = &public_
	}
	return &tmp
}





type ApiPortalPropertiesInput interface {
	pulumi.Input

	ToApiPortalPropertiesOutput() ApiPortalPropertiesOutput
	ToApiPortalPropertiesOutputWithContext(context.Context) ApiPortalPropertiesOutput
}

type ApiPortalPropertiesArgs struct {
	GatewayIds    pulumi.StringArrayInput `pulumi:"gatewayIds"`
	HttpsOnly     pulumi.BoolPtrInput     `pulumi:"httpsOnly"`
	Public        pulumi.BoolPtrInput     `pulumi:"public"`
	SourceUrls    pulumi.StringArrayInput `pulumi:"sourceUrls"`
	SsoProperties SsoPropertiesPtrInput   `pulumi:"ssoProperties"`
}


func (val *ApiPortalPropertiesArgs) Defaults() *ApiPortalPropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HttpsOnly) {
		tmp.HttpsOnly = pulumi.BoolPtr(false)
	}
	if isZero(tmp.Public) {
		tmp.Public = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (ApiPortalPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiPortalProperties)(nil)).Elem()
}

func (i ApiPortalPropertiesArgs) ToApiPortalPropertiesOutput() ApiPortalPropertiesOutput {
	return i.ToApiPortalPropertiesOutputWithContext(context.Background())
}

func (i ApiPortalPropertiesArgs) ToApiPortalPropertiesOutputWithContext(ctx context.Context) ApiPortalPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPortalPropertiesOutput)
}

func (i ApiPortalPropertiesArgs) ToApiPortalPropertiesPtrOutput() ApiPortalPropertiesPtrOutput {
	return i.ToApiPortalPropertiesPtrOutputWithContext(context.Background())
}

func (i ApiPortalPropertiesArgs) ToApiPortalPropertiesPtrOutputWithContext(ctx context.Context) ApiPortalPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPortalPropertiesOutput).ToApiPortalPropertiesPtrOutputWithContext(ctx)
}









type ApiPortalPropertiesPtrInput interface {
	pulumi.Input

	ToApiPortalPropertiesPtrOutput() ApiPortalPropertiesPtrOutput
	ToApiPortalPropertiesPtrOutputWithContext(context.Context) ApiPortalPropertiesPtrOutput
}

type apiPortalPropertiesPtrType ApiPortalPropertiesArgs

func ApiPortalPropertiesPtr(v *ApiPortalPropertiesArgs) ApiPortalPropertiesPtrInput {
	return (*apiPortalPropertiesPtrType)(v)
}

func (*apiPortalPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPortalProperties)(nil)).Elem()
}

func (i *apiPortalPropertiesPtrType) ToApiPortalPropertiesPtrOutput() ApiPortalPropertiesPtrOutput {
	return i.ToApiPortalPropertiesPtrOutputWithContext(context.Background())
}

func (i *apiPortalPropertiesPtrType) ToApiPortalPropertiesPtrOutputWithContext(ctx context.Context) ApiPortalPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPortalPropertiesPtrOutput)
}

type ApiPortalPropertiesOutput struct{ *pulumi.OutputState }

func (ApiPortalPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiPortalProperties)(nil)).Elem()
}

func (o ApiPortalPropertiesOutput) ToApiPortalPropertiesOutput() ApiPortalPropertiesOutput {
	return o
}

func (o ApiPortalPropertiesOutput) ToApiPortalPropertiesOutputWithContext(ctx context.Context) ApiPortalPropertiesOutput {
	return o
}

func (o ApiPortalPropertiesOutput) ToApiPortalPropertiesPtrOutput() ApiPortalPropertiesPtrOutput {
	return o.ToApiPortalPropertiesPtrOutputWithContext(context.Background())
}

func (o ApiPortalPropertiesOutput) ToApiPortalPropertiesPtrOutputWithContext(ctx context.Context) ApiPortalPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiPortalProperties) *ApiPortalProperties {
		return &v
	}).(ApiPortalPropertiesPtrOutput)
}

func (o ApiPortalPropertiesOutput) GatewayIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApiPortalProperties) []string { return v.GatewayIds }).(pulumi.StringArrayOutput)
}

func (o ApiPortalPropertiesOutput) HttpsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApiPortalProperties) *bool { return v.HttpsOnly }).(pulumi.BoolPtrOutput)
}

func (o ApiPortalPropertiesOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApiPortalProperties) *bool { return v.Public }).(pulumi.BoolPtrOutput)
}

func (o ApiPortalPropertiesOutput) SourceUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApiPortalProperties) []string { return v.SourceUrls }).(pulumi.StringArrayOutput)
}

func (o ApiPortalPropertiesOutput) SsoProperties() SsoPropertiesPtrOutput {
	return o.ApplyT(func(v ApiPortalProperties) *SsoProperties { return v.SsoProperties }).(SsoPropertiesPtrOutput)
}

type ApiPortalPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ApiPortalPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPortalProperties)(nil)).Elem()
}

func (o ApiPortalPropertiesPtrOutput) ToApiPortalPropertiesPtrOutput() ApiPortalPropertiesPtrOutput {
	return o
}

func (o ApiPortalPropertiesPtrOutput) ToApiPortalPropertiesPtrOutputWithContext(ctx context.Context) ApiPortalPropertiesPtrOutput {
	return o
}

func (o ApiPortalPropertiesPtrOutput) Elem() ApiPortalPropertiesOutput {
	return o.ApplyT(func(v *ApiPortalProperties) ApiPortalProperties {
		if v != nil {
			return *v
		}
		var ret ApiPortalProperties
		return ret
	}).(ApiPortalPropertiesOutput)
}

func (o ApiPortalPropertiesPtrOutput) GatewayIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiPortalProperties) []string {
		if v == nil {
			return nil
		}
		return v.GatewayIds
	}).(pulumi.StringArrayOutput)
}

func (o ApiPortalPropertiesPtrOutput) HttpsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiPortalProperties) *bool {
		if v == nil {
			return nil
		}
		return v.HttpsOnly
	}).(pulumi.BoolPtrOutput)
}

func (o ApiPortalPropertiesPtrOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiPortalProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Public
	}).(pulumi.BoolPtrOutput)
}

func (o ApiPortalPropertiesPtrOutput) SourceUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiPortalProperties) []string {
		if v == nil {
			return nil
		}
		return v.SourceUrls
	}).(pulumi.StringArrayOutput)
}

func (o ApiPortalPropertiesPtrOutput) SsoProperties() SsoPropertiesPtrOutput {
	return o.ApplyT(func(v *ApiPortalProperties) *SsoProperties {
		if v == nil {
			return nil
		}
		return v.SsoProperties
	}).(SsoPropertiesPtrOutput)
}

type ApiPortalPropertiesResponse struct {
	GatewayIds        []string                          `pulumi:"gatewayIds"`
	HttpsOnly         *bool                             `pulumi:"httpsOnly"`
	Instances         []ApiPortalInstanceResponse       `pulumi:"instances"`
	ProvisioningState string                            `pulumi:"provisioningState"`
	Public            *bool                             `pulumi:"public"`
	ResourceRequests  ApiPortalResourceRequestsResponse `pulumi:"resourceRequests"`
	SourceUrls        []string                          `pulumi:"sourceUrls"`
	SsoProperties     *SsoPropertiesResponse            `pulumi:"ssoProperties"`
	Url               string                            `pulumi:"url"`
}


func (val *ApiPortalPropertiesResponse) Defaults() *ApiPortalPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HttpsOnly) {
		httpsOnly_ := false
		tmp.HttpsOnly = &httpsOnly_
	}
	if isZero(tmp.Public) {
		public_ := false
		tmp.Public = &public_
	}
	return &tmp
}

type ApiPortalPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ApiPortalPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiPortalPropertiesResponse)(nil)).Elem()
}

func (o ApiPortalPropertiesResponseOutput) ToApiPortalPropertiesResponseOutput() ApiPortalPropertiesResponseOutput {
	return o
}

func (o ApiPortalPropertiesResponseOutput) ToApiPortalPropertiesResponseOutputWithContext(ctx context.Context) ApiPortalPropertiesResponseOutput {
	return o
}

func (o ApiPortalPropertiesResponseOutput) GatewayIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApiPortalPropertiesResponse) []string { return v.GatewayIds }).(pulumi.StringArrayOutput)
}

func (o ApiPortalPropertiesResponseOutput) HttpsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApiPortalPropertiesResponse) *bool { return v.HttpsOnly }).(pulumi.BoolPtrOutput)
}

func (o ApiPortalPropertiesResponseOutput) Instances() ApiPortalInstanceResponseArrayOutput {
	return o.ApplyT(func(v ApiPortalPropertiesResponse) []ApiPortalInstanceResponse { return v.Instances }).(ApiPortalInstanceResponseArrayOutput)
}

func (o ApiPortalPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ApiPortalPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ApiPortalPropertiesResponseOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApiPortalPropertiesResponse) *bool { return v.Public }).(pulumi.BoolPtrOutput)
}

func (o ApiPortalPropertiesResponseOutput) ResourceRequests() ApiPortalResourceRequestsResponseOutput {
	return o.ApplyT(func(v ApiPortalPropertiesResponse) ApiPortalResourceRequestsResponse { return v.ResourceRequests }).(ApiPortalResourceRequestsResponseOutput)
}

func (o ApiPortalPropertiesResponseOutput) SourceUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApiPortalPropertiesResponse) []string { return v.SourceUrls }).(pulumi.StringArrayOutput)
}

func (o ApiPortalPropertiesResponseOutput) SsoProperties() SsoPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ApiPortalPropertiesResponse) *SsoPropertiesResponse { return v.SsoProperties }).(SsoPropertiesResponsePtrOutput)
}

func (o ApiPortalPropertiesResponseOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v ApiPortalPropertiesResponse) string { return v.Url }).(pulumi.StringOutput)
}

type ApiPortalResourceRequestsResponse struct {
	Cpu    string `pulumi:"cpu"`
	Memory string `pulumi:"memory"`
}

type ApiPortalResourceRequestsResponseOutput struct{ *pulumi.OutputState }

func (ApiPortalResourceRequestsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiPortalResourceRequestsResponse)(nil)).Elem()
}

func (o ApiPortalResourceRequestsResponseOutput) ToApiPortalResourceRequestsResponseOutput() ApiPortalResourceRequestsResponseOutput {
	return o
}

func (o ApiPortalResourceRequestsResponseOutput) ToApiPortalResourceRequestsResponseOutputWithContext(ctx context.Context) ApiPortalResourceRequestsResponseOutput {
	return o
}

func (o ApiPortalResourceRequestsResponseOutput) Cpu() pulumi.StringOutput {
	return o.ApplyT(func(v ApiPortalResourceRequestsResponse) string { return v.Cpu }).(pulumi.StringOutput)
}

func (o ApiPortalResourceRequestsResponseOutput) Memory() pulumi.StringOutput {
	return o.ApplyT(func(v ApiPortalResourceRequestsResponse) string { return v.Memory }).(pulumi.StringOutput)
}

type AppResourceProperties struct {
	ActiveDeploymentName *string         `pulumi:"activeDeploymentName"`
	Fqdn                 *string         `pulumi:"fqdn"`
	HttpsOnly            *bool           `pulumi:"httpsOnly"`
	PersistentDisk       *PersistentDisk `pulumi:"persistentDisk"`
	Public               *bool           `pulumi:"public"`
	TemporaryDisk        *TemporaryDisk  `pulumi:"temporaryDisk"`
}


func (val *AppResourceProperties) Defaults() *AppResourceProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HttpsOnly) {
		httpsOnly_ := false
		tmp.HttpsOnly = &httpsOnly_
	}
	tmp.TemporaryDisk = tmp.TemporaryDisk.Defaults()

	return &tmp
}





type AppResourcePropertiesInput interface {
	pulumi.Input

	ToAppResourcePropertiesOutput() AppResourcePropertiesOutput
	ToAppResourcePropertiesOutputWithContext(context.Context) AppResourcePropertiesOutput
}

type AppResourcePropertiesArgs struct {
	ActiveDeploymentName pulumi.StringPtrInput  `pulumi:"activeDeploymentName"`
	Fqdn                 pulumi.StringPtrInput  `pulumi:"fqdn"`
	HttpsOnly            pulumi.BoolPtrInput    `pulumi:"httpsOnly"`
	PersistentDisk       PersistentDiskPtrInput `pulumi:"persistentDisk"`
	Public               pulumi.BoolPtrInput    `pulumi:"public"`
	TemporaryDisk        TemporaryDiskPtrInput  `pulumi:"temporaryDisk"`
}


func (val *AppResourcePropertiesArgs) Defaults() *AppResourcePropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HttpsOnly) {
		tmp.HttpsOnly = pulumi.BoolPtr(false)
	}

	return &tmp
}
func (AppResourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppResourceProperties)(nil)).Elem()
}

func (i AppResourcePropertiesArgs) ToAppResourcePropertiesOutput() AppResourcePropertiesOutput {
	return i.ToAppResourcePropertiesOutputWithContext(context.Background())
}

func (i AppResourcePropertiesArgs) ToAppResourcePropertiesOutputWithContext(ctx context.Context) AppResourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppResourcePropertiesOutput)
}

func (i AppResourcePropertiesArgs) ToAppResourcePropertiesPtrOutput() AppResourcePropertiesPtrOutput {
	return i.ToAppResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i AppResourcePropertiesArgs) ToAppResourcePropertiesPtrOutputWithContext(ctx context.Context) AppResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppResourcePropertiesOutput).ToAppResourcePropertiesPtrOutputWithContext(ctx)
}









type AppResourcePropertiesPtrInput interface {
	pulumi.Input

	ToAppResourcePropertiesPtrOutput() AppResourcePropertiesPtrOutput
	ToAppResourcePropertiesPtrOutputWithContext(context.Context) AppResourcePropertiesPtrOutput
}

type appResourcePropertiesPtrType AppResourcePropertiesArgs

func AppResourcePropertiesPtr(v *AppResourcePropertiesArgs) AppResourcePropertiesPtrInput {
	return (*appResourcePropertiesPtrType)(v)
}

func (*appResourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppResourceProperties)(nil)).Elem()
}

func (i *appResourcePropertiesPtrType) ToAppResourcePropertiesPtrOutput() AppResourcePropertiesPtrOutput {
	return i.ToAppResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *appResourcePropertiesPtrType) ToAppResourcePropertiesPtrOutputWithContext(ctx context.Context) AppResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppResourcePropertiesPtrOutput)
}

type AppResourcePropertiesOutput struct{ *pulumi.OutputState }

func (AppResourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppResourceProperties)(nil)).Elem()
}

func (o AppResourcePropertiesOutput) ToAppResourcePropertiesOutput() AppResourcePropertiesOutput {
	return o
}

func (o AppResourcePropertiesOutput) ToAppResourcePropertiesOutputWithContext(ctx context.Context) AppResourcePropertiesOutput {
	return o
}

func (o AppResourcePropertiesOutput) ToAppResourcePropertiesPtrOutput() AppResourcePropertiesPtrOutput {
	return o.ToAppResourcePropertiesPtrOutputWithContext(context.Background())
}

func (o AppResourcePropertiesOutput) ToAppResourcePropertiesPtrOutputWithContext(ctx context.Context) AppResourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppResourceProperties) *AppResourceProperties {
		return &v
	}).(AppResourcePropertiesPtrOutput)
}

func (o AppResourcePropertiesOutput) ActiveDeploymentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppResourceProperties) *string { return v.ActiveDeploymentName }).(pulumi.StringPtrOutput)
}

func (o AppResourcePropertiesOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppResourceProperties) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o AppResourcePropertiesOutput) HttpsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AppResourceProperties) *bool { return v.HttpsOnly }).(pulumi.BoolPtrOutput)
}

func (o AppResourcePropertiesOutput) PersistentDisk() PersistentDiskPtrOutput {
	return o.ApplyT(func(v AppResourceProperties) *PersistentDisk { return v.PersistentDisk }).(PersistentDiskPtrOutput)
}

func (o AppResourcePropertiesOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AppResourceProperties) *bool { return v.Public }).(pulumi.BoolPtrOutput)
}

func (o AppResourcePropertiesOutput) TemporaryDisk() TemporaryDiskPtrOutput {
	return o.ApplyT(func(v AppResourceProperties) *TemporaryDisk { return v.TemporaryDisk }).(TemporaryDiskPtrOutput)
}

type AppResourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AppResourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppResourceProperties)(nil)).Elem()
}

func (o AppResourcePropertiesPtrOutput) ToAppResourcePropertiesPtrOutput() AppResourcePropertiesPtrOutput {
	return o
}

func (o AppResourcePropertiesPtrOutput) ToAppResourcePropertiesPtrOutputWithContext(ctx context.Context) AppResourcePropertiesPtrOutput {
	return o
}

func (o AppResourcePropertiesPtrOutput) Elem() AppResourcePropertiesOutput {
	return o.ApplyT(func(v *AppResourceProperties) AppResourceProperties {
		if v != nil {
			return *v
		}
		var ret AppResourceProperties
		return ret
	}).(AppResourcePropertiesOutput)
}

func (o AppResourcePropertiesPtrOutput) ActiveDeploymentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ActiveDeploymentName
	}).(pulumi.StringPtrOutput)
}

func (o AppResourcePropertiesPtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Fqdn
	}).(pulumi.StringPtrOutput)
}

func (o AppResourcePropertiesPtrOutput) HttpsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppResourceProperties) *bool {
		if v == nil {
			return nil
		}
		return v.HttpsOnly
	}).(pulumi.BoolPtrOutput)
}

func (o AppResourcePropertiesPtrOutput) PersistentDisk() PersistentDiskPtrOutput {
	return o.ApplyT(func(v *AppResourceProperties) *PersistentDisk {
		if v == nil {
			return nil
		}
		return v.PersistentDisk
	}).(PersistentDiskPtrOutput)
}

func (o AppResourcePropertiesPtrOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppResourceProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Public
	}).(pulumi.BoolPtrOutput)
}

func (o AppResourcePropertiesPtrOutput) TemporaryDisk() TemporaryDiskPtrOutput {
	return o.ApplyT(func(v *AppResourceProperties) *TemporaryDisk {
		if v == nil {
			return nil
		}
		return v.TemporaryDisk
	}).(TemporaryDiskPtrOutput)
}

type AppResourcePropertiesResponse struct {
	ActiveDeploymentName *string                 `pulumi:"activeDeploymentName"`
	CreatedTime          string                  `pulumi:"createdTime"`
	Fqdn                 *string                 `pulumi:"fqdn"`
	HttpsOnly            *bool                   `pulumi:"httpsOnly"`
	PersistentDisk       *PersistentDiskResponse `pulumi:"persistentDisk"`
	ProvisioningState    string                  `pulumi:"provisioningState"`
	Public               *bool                   `pulumi:"public"`
	TemporaryDisk        *TemporaryDiskResponse  `pulumi:"temporaryDisk"`
	Url                  string                  `pulumi:"url"`
}


func (val *AppResourcePropertiesResponse) Defaults() *AppResourcePropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HttpsOnly) {
		httpsOnly_ := false
		tmp.HttpsOnly = &httpsOnly_
	}
	tmp.TemporaryDisk = tmp.TemporaryDisk.Defaults()

	return &tmp
}

type AppResourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (AppResourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppResourcePropertiesResponse)(nil)).Elem()
}

func (o AppResourcePropertiesResponseOutput) ToAppResourcePropertiesResponseOutput() AppResourcePropertiesResponseOutput {
	return o
}

func (o AppResourcePropertiesResponseOutput) ToAppResourcePropertiesResponseOutputWithContext(ctx context.Context) AppResourcePropertiesResponseOutput {
	return o
}

func (o AppResourcePropertiesResponseOutput) ActiveDeploymentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppResourcePropertiesResponse) *string { return v.ActiveDeploymentName }).(pulumi.StringPtrOutput)
}

func (o AppResourcePropertiesResponseOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v AppResourcePropertiesResponse) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o AppResourcePropertiesResponseOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppResourcePropertiesResponse) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o AppResourcePropertiesResponseOutput) HttpsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AppResourcePropertiesResponse) *bool { return v.HttpsOnly }).(pulumi.BoolPtrOutput)
}

func (o AppResourcePropertiesResponseOutput) PersistentDisk() PersistentDiskResponsePtrOutput {
	return o.ApplyT(func(v AppResourcePropertiesResponse) *PersistentDiskResponse { return v.PersistentDisk }).(PersistentDiskResponsePtrOutput)
}

func (o AppResourcePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AppResourcePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AppResourcePropertiesResponseOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AppResourcePropertiesResponse) *bool { return v.Public }).(pulumi.BoolPtrOutput)
}

func (o AppResourcePropertiesResponseOutput) TemporaryDisk() TemporaryDiskResponsePtrOutput {
	return o.ApplyT(func(v AppResourcePropertiesResponse) *TemporaryDiskResponse { return v.TemporaryDisk }).(TemporaryDiskResponsePtrOutput)
}

func (o AppResourcePropertiesResponseOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v AppResourcePropertiesResponse) string { return v.Url }).(pulumi.StringOutput)
}

type BindingResourceProperties struct {
	BindingParameters map[string]interface{} `pulumi:"bindingParameters"`
	Key               *string                `pulumi:"key"`
	ResourceId        *string                `pulumi:"resourceId"`
}





type BindingResourcePropertiesInput interface {
	pulumi.Input

	ToBindingResourcePropertiesOutput() BindingResourcePropertiesOutput
	ToBindingResourcePropertiesOutputWithContext(context.Context) BindingResourcePropertiesOutput
}

type BindingResourcePropertiesArgs struct {
	BindingParameters pulumi.MapInput       `pulumi:"bindingParameters"`
	Key               pulumi.StringPtrInput `pulumi:"key"`
	ResourceId        pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (BindingResourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BindingResourceProperties)(nil)).Elem()
}

func (i BindingResourcePropertiesArgs) ToBindingResourcePropertiesOutput() BindingResourcePropertiesOutput {
	return i.ToBindingResourcePropertiesOutputWithContext(context.Background())
}

func (i BindingResourcePropertiesArgs) ToBindingResourcePropertiesOutputWithContext(ctx context.Context) BindingResourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingResourcePropertiesOutput)
}

func (i BindingResourcePropertiesArgs) ToBindingResourcePropertiesPtrOutput() BindingResourcePropertiesPtrOutput {
	return i.ToBindingResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i BindingResourcePropertiesArgs) ToBindingResourcePropertiesPtrOutputWithContext(ctx context.Context) BindingResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingResourcePropertiesOutput).ToBindingResourcePropertiesPtrOutputWithContext(ctx)
}









type BindingResourcePropertiesPtrInput interface {
	pulumi.Input

	ToBindingResourcePropertiesPtrOutput() BindingResourcePropertiesPtrOutput
	ToBindingResourcePropertiesPtrOutputWithContext(context.Context) BindingResourcePropertiesPtrOutput
}

type bindingResourcePropertiesPtrType BindingResourcePropertiesArgs

func BindingResourcePropertiesPtr(v *BindingResourcePropertiesArgs) BindingResourcePropertiesPtrInput {
	return (*bindingResourcePropertiesPtrType)(v)
}

func (*bindingResourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BindingResourceProperties)(nil)).Elem()
}

func (i *bindingResourcePropertiesPtrType) ToBindingResourcePropertiesPtrOutput() BindingResourcePropertiesPtrOutput {
	return i.ToBindingResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *bindingResourcePropertiesPtrType) ToBindingResourcePropertiesPtrOutputWithContext(ctx context.Context) BindingResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingResourcePropertiesPtrOutput)
}

type BindingResourcePropertiesOutput struct{ *pulumi.OutputState }

func (BindingResourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BindingResourceProperties)(nil)).Elem()
}

func (o BindingResourcePropertiesOutput) ToBindingResourcePropertiesOutput() BindingResourcePropertiesOutput {
	return o
}

func (o BindingResourcePropertiesOutput) ToBindingResourcePropertiesOutputWithContext(ctx context.Context) BindingResourcePropertiesOutput {
	return o
}

func (o BindingResourcePropertiesOutput) ToBindingResourcePropertiesPtrOutput() BindingResourcePropertiesPtrOutput {
	return o.ToBindingResourcePropertiesPtrOutputWithContext(context.Background())
}

func (o BindingResourcePropertiesOutput) ToBindingResourcePropertiesPtrOutputWithContext(ctx context.Context) BindingResourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BindingResourceProperties) *BindingResourceProperties {
		return &v
	}).(BindingResourcePropertiesPtrOutput)
}

func (o BindingResourcePropertiesOutput) BindingParameters() pulumi.MapOutput {
	return o.ApplyT(func(v BindingResourceProperties) map[string]interface{} { return v.BindingParameters }).(pulumi.MapOutput)
}

func (o BindingResourcePropertiesOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BindingResourceProperties) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o BindingResourcePropertiesOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BindingResourceProperties) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type BindingResourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (BindingResourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BindingResourceProperties)(nil)).Elem()
}

func (o BindingResourcePropertiesPtrOutput) ToBindingResourcePropertiesPtrOutput() BindingResourcePropertiesPtrOutput {
	return o
}

func (o BindingResourcePropertiesPtrOutput) ToBindingResourcePropertiesPtrOutputWithContext(ctx context.Context) BindingResourcePropertiesPtrOutput {
	return o
}

func (o BindingResourcePropertiesPtrOutput) Elem() BindingResourcePropertiesOutput {
	return o.ApplyT(func(v *BindingResourceProperties) BindingResourceProperties {
		if v != nil {
			return *v
		}
		var ret BindingResourceProperties
		return ret
	}).(BindingResourcePropertiesOutput)
}

func (o BindingResourcePropertiesPtrOutput) BindingParameters() pulumi.MapOutput {
	return o.ApplyT(func(v *BindingResourceProperties) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.BindingParameters
	}).(pulumi.MapOutput)
}

func (o BindingResourcePropertiesPtrOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BindingResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Key
	}).(pulumi.StringPtrOutput)
}

func (o BindingResourcePropertiesPtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BindingResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

type BindingResourcePropertiesResponse struct {
	BindingParameters   map[string]interface{} `pulumi:"bindingParameters"`
	CreatedAt           string                 `pulumi:"createdAt"`
	GeneratedProperties string                 `pulumi:"generatedProperties"`
	Key                 *string                `pulumi:"key"`
	ResourceId          *string                `pulumi:"resourceId"`
	ResourceName        string                 `pulumi:"resourceName"`
	ResourceType        string                 `pulumi:"resourceType"`
	UpdatedAt           string                 `pulumi:"updatedAt"`
}

type BindingResourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (BindingResourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BindingResourcePropertiesResponse)(nil)).Elem()
}

func (o BindingResourcePropertiesResponseOutput) ToBindingResourcePropertiesResponseOutput() BindingResourcePropertiesResponseOutput {
	return o
}

func (o BindingResourcePropertiesResponseOutput) ToBindingResourcePropertiesResponseOutputWithContext(ctx context.Context) BindingResourcePropertiesResponseOutput {
	return o
}

func (o BindingResourcePropertiesResponseOutput) BindingParameters() pulumi.MapOutput {
	return o.ApplyT(func(v BindingResourcePropertiesResponse) map[string]interface{} { return v.BindingParameters }).(pulumi.MapOutput)
}

func (o BindingResourcePropertiesResponseOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v BindingResourcePropertiesResponse) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o BindingResourcePropertiesResponseOutput) GeneratedProperties() pulumi.StringOutput {
	return o.ApplyT(func(v BindingResourcePropertiesResponse) string { return v.GeneratedProperties }).(pulumi.StringOutput)
}

func (o BindingResourcePropertiesResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BindingResourcePropertiesResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o BindingResourcePropertiesResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BindingResourcePropertiesResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o BindingResourcePropertiesResponseOutput) ResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v BindingResourcePropertiesResponse) string { return v.ResourceName }).(pulumi.StringOutput)
}

func (o BindingResourcePropertiesResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v BindingResourcePropertiesResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o BindingResourcePropertiesResponseOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v BindingResourcePropertiesResponse) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

type BuilderProperties struct {
	BuildpackGroups []BuildpacksGroupProperties `pulumi:"buildpackGroups"`
	Stack           *StackProperties            `pulumi:"stack"`
}





type BuilderPropertiesInput interface {
	pulumi.Input

	ToBuilderPropertiesOutput() BuilderPropertiesOutput
	ToBuilderPropertiesOutputWithContext(context.Context) BuilderPropertiesOutput
}

type BuilderPropertiesArgs struct {
	BuildpackGroups BuildpacksGroupPropertiesArrayInput `pulumi:"buildpackGroups"`
	Stack           StackPropertiesPtrInput             `pulumi:"stack"`
}

func (BuilderPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BuilderProperties)(nil)).Elem()
}

func (i BuilderPropertiesArgs) ToBuilderPropertiesOutput() BuilderPropertiesOutput {
	return i.ToBuilderPropertiesOutputWithContext(context.Background())
}

func (i BuilderPropertiesArgs) ToBuilderPropertiesOutputWithContext(ctx context.Context) BuilderPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuilderPropertiesOutput)
}

func (i BuilderPropertiesArgs) ToBuilderPropertiesPtrOutput() BuilderPropertiesPtrOutput {
	return i.ToBuilderPropertiesPtrOutputWithContext(context.Background())
}

func (i BuilderPropertiesArgs) ToBuilderPropertiesPtrOutputWithContext(ctx context.Context) BuilderPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuilderPropertiesOutput).ToBuilderPropertiesPtrOutputWithContext(ctx)
}









type BuilderPropertiesPtrInput interface {
	pulumi.Input

	ToBuilderPropertiesPtrOutput() BuilderPropertiesPtrOutput
	ToBuilderPropertiesPtrOutputWithContext(context.Context) BuilderPropertiesPtrOutput
}

type builderPropertiesPtrType BuilderPropertiesArgs

func BuilderPropertiesPtr(v *BuilderPropertiesArgs) BuilderPropertiesPtrInput {
	return (*builderPropertiesPtrType)(v)
}

func (*builderPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BuilderProperties)(nil)).Elem()
}

func (i *builderPropertiesPtrType) ToBuilderPropertiesPtrOutput() BuilderPropertiesPtrOutput {
	return i.ToBuilderPropertiesPtrOutputWithContext(context.Background())
}

func (i *builderPropertiesPtrType) ToBuilderPropertiesPtrOutputWithContext(ctx context.Context) BuilderPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuilderPropertiesPtrOutput)
}

type BuilderPropertiesOutput struct{ *pulumi.OutputState }

func (BuilderPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuilderProperties)(nil)).Elem()
}

func (o BuilderPropertiesOutput) ToBuilderPropertiesOutput() BuilderPropertiesOutput {
	return o
}

func (o BuilderPropertiesOutput) ToBuilderPropertiesOutputWithContext(ctx context.Context) BuilderPropertiesOutput {
	return o
}

func (o BuilderPropertiesOutput) ToBuilderPropertiesPtrOutput() BuilderPropertiesPtrOutput {
	return o.ToBuilderPropertiesPtrOutputWithContext(context.Background())
}

func (o BuilderPropertiesOutput) ToBuilderPropertiesPtrOutputWithContext(ctx context.Context) BuilderPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BuilderProperties) *BuilderProperties {
		return &v
	}).(BuilderPropertiesPtrOutput)
}

func (o BuilderPropertiesOutput) BuildpackGroups() BuildpacksGroupPropertiesArrayOutput {
	return o.ApplyT(func(v BuilderProperties) []BuildpacksGroupProperties { return v.BuildpackGroups }).(BuildpacksGroupPropertiesArrayOutput)
}

func (o BuilderPropertiesOutput) Stack() StackPropertiesPtrOutput {
	return o.ApplyT(func(v BuilderProperties) *StackProperties { return v.Stack }).(StackPropertiesPtrOutput)
}

type BuilderPropertiesPtrOutput struct{ *pulumi.OutputState }

func (BuilderPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuilderProperties)(nil)).Elem()
}

func (o BuilderPropertiesPtrOutput) ToBuilderPropertiesPtrOutput() BuilderPropertiesPtrOutput {
	return o
}

func (o BuilderPropertiesPtrOutput) ToBuilderPropertiesPtrOutputWithContext(ctx context.Context) BuilderPropertiesPtrOutput {
	return o
}

func (o BuilderPropertiesPtrOutput) Elem() BuilderPropertiesOutput {
	return o.ApplyT(func(v *BuilderProperties) BuilderProperties {
		if v != nil {
			return *v
		}
		var ret BuilderProperties
		return ret
	}).(BuilderPropertiesOutput)
}

func (o BuilderPropertiesPtrOutput) BuildpackGroups() BuildpacksGroupPropertiesArrayOutput {
	return o.ApplyT(func(v *BuilderProperties) []BuildpacksGroupProperties {
		if v == nil {
			return nil
		}
		return v.BuildpackGroups
	}).(BuildpacksGroupPropertiesArrayOutput)
}

func (o BuilderPropertiesPtrOutput) Stack() StackPropertiesPtrOutput {
	return o.ApplyT(func(v *BuilderProperties) *StackProperties {
		if v == nil {
			return nil
		}
		return v.Stack
	}).(StackPropertiesPtrOutput)
}

type BuilderPropertiesResponse struct {
	BuildpackGroups   []BuildpacksGroupPropertiesResponse `pulumi:"buildpackGroups"`
	ProvisioningState string                              `pulumi:"provisioningState"`
	Stack             *StackPropertiesResponse            `pulumi:"stack"`
}

type BuilderPropertiesResponseOutput struct{ *pulumi.OutputState }

func (BuilderPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuilderPropertiesResponse)(nil)).Elem()
}

func (o BuilderPropertiesResponseOutput) ToBuilderPropertiesResponseOutput() BuilderPropertiesResponseOutput {
	return o
}

func (o BuilderPropertiesResponseOutput) ToBuilderPropertiesResponseOutputWithContext(ctx context.Context) BuilderPropertiesResponseOutput {
	return o
}

func (o BuilderPropertiesResponseOutput) BuildpackGroups() BuildpacksGroupPropertiesResponseArrayOutput {
	return o.ApplyT(func(v BuilderPropertiesResponse) []BuildpacksGroupPropertiesResponse { return v.BuildpackGroups }).(BuildpacksGroupPropertiesResponseArrayOutput)
}

func (o BuilderPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v BuilderPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o BuilderPropertiesResponseOutput) Stack() StackPropertiesResponsePtrOutput {
	return o.ApplyT(func(v BuilderPropertiesResponse) *StackPropertiesResponse { return v.Stack }).(StackPropertiesResponsePtrOutput)
}

type BuildpackBindingLaunchProperties struct {
	Properties map[string]string `pulumi:"properties"`
	Secrets    map[string]string `pulumi:"secrets"`
}





type BuildpackBindingLaunchPropertiesInput interface {
	pulumi.Input

	ToBuildpackBindingLaunchPropertiesOutput() BuildpackBindingLaunchPropertiesOutput
	ToBuildpackBindingLaunchPropertiesOutputWithContext(context.Context) BuildpackBindingLaunchPropertiesOutput
}

type BuildpackBindingLaunchPropertiesArgs struct {
	Properties pulumi.StringMapInput `pulumi:"properties"`
	Secrets    pulumi.StringMapInput `pulumi:"secrets"`
}

func (BuildpackBindingLaunchPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildpackBindingLaunchProperties)(nil)).Elem()
}

func (i BuildpackBindingLaunchPropertiesArgs) ToBuildpackBindingLaunchPropertiesOutput() BuildpackBindingLaunchPropertiesOutput {
	return i.ToBuildpackBindingLaunchPropertiesOutputWithContext(context.Background())
}

func (i BuildpackBindingLaunchPropertiesArgs) ToBuildpackBindingLaunchPropertiesOutputWithContext(ctx context.Context) BuildpackBindingLaunchPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildpackBindingLaunchPropertiesOutput)
}

func (i BuildpackBindingLaunchPropertiesArgs) ToBuildpackBindingLaunchPropertiesPtrOutput() BuildpackBindingLaunchPropertiesPtrOutput {
	return i.ToBuildpackBindingLaunchPropertiesPtrOutputWithContext(context.Background())
}

func (i BuildpackBindingLaunchPropertiesArgs) ToBuildpackBindingLaunchPropertiesPtrOutputWithContext(ctx context.Context) BuildpackBindingLaunchPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildpackBindingLaunchPropertiesOutput).ToBuildpackBindingLaunchPropertiesPtrOutputWithContext(ctx)
}









type BuildpackBindingLaunchPropertiesPtrInput interface {
	pulumi.Input

	ToBuildpackBindingLaunchPropertiesPtrOutput() BuildpackBindingLaunchPropertiesPtrOutput
	ToBuildpackBindingLaunchPropertiesPtrOutputWithContext(context.Context) BuildpackBindingLaunchPropertiesPtrOutput
}

type buildpackBindingLaunchPropertiesPtrType BuildpackBindingLaunchPropertiesArgs

func BuildpackBindingLaunchPropertiesPtr(v *BuildpackBindingLaunchPropertiesArgs) BuildpackBindingLaunchPropertiesPtrInput {
	return (*buildpackBindingLaunchPropertiesPtrType)(v)
}

func (*buildpackBindingLaunchPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildpackBindingLaunchProperties)(nil)).Elem()
}

func (i *buildpackBindingLaunchPropertiesPtrType) ToBuildpackBindingLaunchPropertiesPtrOutput() BuildpackBindingLaunchPropertiesPtrOutput {
	return i.ToBuildpackBindingLaunchPropertiesPtrOutputWithContext(context.Background())
}

func (i *buildpackBindingLaunchPropertiesPtrType) ToBuildpackBindingLaunchPropertiesPtrOutputWithContext(ctx context.Context) BuildpackBindingLaunchPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildpackBindingLaunchPropertiesPtrOutput)
}

type BuildpackBindingLaunchPropertiesOutput struct{ *pulumi.OutputState }

func (BuildpackBindingLaunchPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildpackBindingLaunchProperties)(nil)).Elem()
}

func (o BuildpackBindingLaunchPropertiesOutput) ToBuildpackBindingLaunchPropertiesOutput() BuildpackBindingLaunchPropertiesOutput {
	return o
}

func (o BuildpackBindingLaunchPropertiesOutput) ToBuildpackBindingLaunchPropertiesOutputWithContext(ctx context.Context) BuildpackBindingLaunchPropertiesOutput {
	return o
}

func (o BuildpackBindingLaunchPropertiesOutput) ToBuildpackBindingLaunchPropertiesPtrOutput() BuildpackBindingLaunchPropertiesPtrOutput {
	return o.ToBuildpackBindingLaunchPropertiesPtrOutputWithContext(context.Background())
}

func (o BuildpackBindingLaunchPropertiesOutput) ToBuildpackBindingLaunchPropertiesPtrOutputWithContext(ctx context.Context) BuildpackBindingLaunchPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BuildpackBindingLaunchProperties) *BuildpackBindingLaunchProperties {
		return &v
	}).(BuildpackBindingLaunchPropertiesPtrOutput)
}

func (o BuildpackBindingLaunchPropertiesOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v BuildpackBindingLaunchProperties) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o BuildpackBindingLaunchPropertiesOutput) Secrets() pulumi.StringMapOutput {
	return o.ApplyT(func(v BuildpackBindingLaunchProperties) map[string]string { return v.Secrets }).(pulumi.StringMapOutput)
}

type BuildpackBindingLaunchPropertiesPtrOutput struct{ *pulumi.OutputState }

func (BuildpackBindingLaunchPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildpackBindingLaunchProperties)(nil)).Elem()
}

func (o BuildpackBindingLaunchPropertiesPtrOutput) ToBuildpackBindingLaunchPropertiesPtrOutput() BuildpackBindingLaunchPropertiesPtrOutput {
	return o
}

func (o BuildpackBindingLaunchPropertiesPtrOutput) ToBuildpackBindingLaunchPropertiesPtrOutputWithContext(ctx context.Context) BuildpackBindingLaunchPropertiesPtrOutput {
	return o
}

func (o BuildpackBindingLaunchPropertiesPtrOutput) Elem() BuildpackBindingLaunchPropertiesOutput {
	return o.ApplyT(func(v *BuildpackBindingLaunchProperties) BuildpackBindingLaunchProperties {
		if v != nil {
			return *v
		}
		var ret BuildpackBindingLaunchProperties
		return ret
	}).(BuildpackBindingLaunchPropertiesOutput)
}

func (o BuildpackBindingLaunchPropertiesPtrOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BuildpackBindingLaunchProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(pulumi.StringMapOutput)
}

func (o BuildpackBindingLaunchPropertiesPtrOutput) Secrets() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BuildpackBindingLaunchProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.Secrets
	}).(pulumi.StringMapOutput)
}

type BuildpackBindingLaunchPropertiesResponse struct {
	Properties map[string]string `pulumi:"properties"`
	Secrets    map[string]string `pulumi:"secrets"`
}

type BuildpackBindingLaunchPropertiesResponseOutput struct{ *pulumi.OutputState }

func (BuildpackBindingLaunchPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildpackBindingLaunchPropertiesResponse)(nil)).Elem()
}

func (o BuildpackBindingLaunchPropertiesResponseOutput) ToBuildpackBindingLaunchPropertiesResponseOutput() BuildpackBindingLaunchPropertiesResponseOutput {
	return o
}

func (o BuildpackBindingLaunchPropertiesResponseOutput) ToBuildpackBindingLaunchPropertiesResponseOutputWithContext(ctx context.Context) BuildpackBindingLaunchPropertiesResponseOutput {
	return o
}

func (o BuildpackBindingLaunchPropertiesResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v BuildpackBindingLaunchPropertiesResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o BuildpackBindingLaunchPropertiesResponseOutput) Secrets() pulumi.StringMapOutput {
	return o.ApplyT(func(v BuildpackBindingLaunchPropertiesResponse) map[string]string { return v.Secrets }).(pulumi.StringMapOutput)
}

type BuildpackBindingLaunchPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (BuildpackBindingLaunchPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildpackBindingLaunchPropertiesResponse)(nil)).Elem()
}

func (o BuildpackBindingLaunchPropertiesResponsePtrOutput) ToBuildpackBindingLaunchPropertiesResponsePtrOutput() BuildpackBindingLaunchPropertiesResponsePtrOutput {
	return o
}

func (o BuildpackBindingLaunchPropertiesResponsePtrOutput) ToBuildpackBindingLaunchPropertiesResponsePtrOutputWithContext(ctx context.Context) BuildpackBindingLaunchPropertiesResponsePtrOutput {
	return o
}

func (o BuildpackBindingLaunchPropertiesResponsePtrOutput) Elem() BuildpackBindingLaunchPropertiesResponseOutput {
	return o.ApplyT(func(v *BuildpackBindingLaunchPropertiesResponse) BuildpackBindingLaunchPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret BuildpackBindingLaunchPropertiesResponse
		return ret
	}).(BuildpackBindingLaunchPropertiesResponseOutput)
}

func (o BuildpackBindingLaunchPropertiesResponsePtrOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BuildpackBindingLaunchPropertiesResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(pulumi.StringMapOutput)
}

func (o BuildpackBindingLaunchPropertiesResponsePtrOutput) Secrets() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BuildpackBindingLaunchPropertiesResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Secrets
	}).(pulumi.StringMapOutput)
}

type BuildpackBindingProperties struct {
	BindingType      *string                           `pulumi:"bindingType"`
	LaunchProperties *BuildpackBindingLaunchProperties `pulumi:"launchProperties"`
}





type BuildpackBindingPropertiesInput interface {
	pulumi.Input

	ToBuildpackBindingPropertiesOutput() BuildpackBindingPropertiesOutput
	ToBuildpackBindingPropertiesOutputWithContext(context.Context) BuildpackBindingPropertiesOutput
}

type BuildpackBindingPropertiesArgs struct {
	BindingType      pulumi.StringPtrInput                    `pulumi:"bindingType"`
	LaunchProperties BuildpackBindingLaunchPropertiesPtrInput `pulumi:"launchProperties"`
}

func (BuildpackBindingPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildpackBindingProperties)(nil)).Elem()
}

func (i BuildpackBindingPropertiesArgs) ToBuildpackBindingPropertiesOutput() BuildpackBindingPropertiesOutput {
	return i.ToBuildpackBindingPropertiesOutputWithContext(context.Background())
}

func (i BuildpackBindingPropertiesArgs) ToBuildpackBindingPropertiesOutputWithContext(ctx context.Context) BuildpackBindingPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildpackBindingPropertiesOutput)
}

func (i BuildpackBindingPropertiesArgs) ToBuildpackBindingPropertiesPtrOutput() BuildpackBindingPropertiesPtrOutput {
	return i.ToBuildpackBindingPropertiesPtrOutputWithContext(context.Background())
}

func (i BuildpackBindingPropertiesArgs) ToBuildpackBindingPropertiesPtrOutputWithContext(ctx context.Context) BuildpackBindingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildpackBindingPropertiesOutput).ToBuildpackBindingPropertiesPtrOutputWithContext(ctx)
}









type BuildpackBindingPropertiesPtrInput interface {
	pulumi.Input

	ToBuildpackBindingPropertiesPtrOutput() BuildpackBindingPropertiesPtrOutput
	ToBuildpackBindingPropertiesPtrOutputWithContext(context.Context) BuildpackBindingPropertiesPtrOutput
}

type buildpackBindingPropertiesPtrType BuildpackBindingPropertiesArgs

func BuildpackBindingPropertiesPtr(v *BuildpackBindingPropertiesArgs) BuildpackBindingPropertiesPtrInput {
	return (*buildpackBindingPropertiesPtrType)(v)
}

func (*buildpackBindingPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildpackBindingProperties)(nil)).Elem()
}

func (i *buildpackBindingPropertiesPtrType) ToBuildpackBindingPropertiesPtrOutput() BuildpackBindingPropertiesPtrOutput {
	return i.ToBuildpackBindingPropertiesPtrOutputWithContext(context.Background())
}

func (i *buildpackBindingPropertiesPtrType) ToBuildpackBindingPropertiesPtrOutputWithContext(ctx context.Context) BuildpackBindingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildpackBindingPropertiesPtrOutput)
}

type BuildpackBindingPropertiesOutput struct{ *pulumi.OutputState }

func (BuildpackBindingPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildpackBindingProperties)(nil)).Elem()
}

func (o BuildpackBindingPropertiesOutput) ToBuildpackBindingPropertiesOutput() BuildpackBindingPropertiesOutput {
	return o
}

func (o BuildpackBindingPropertiesOutput) ToBuildpackBindingPropertiesOutputWithContext(ctx context.Context) BuildpackBindingPropertiesOutput {
	return o
}

func (o BuildpackBindingPropertiesOutput) ToBuildpackBindingPropertiesPtrOutput() BuildpackBindingPropertiesPtrOutput {
	return o.ToBuildpackBindingPropertiesPtrOutputWithContext(context.Background())
}

func (o BuildpackBindingPropertiesOutput) ToBuildpackBindingPropertiesPtrOutputWithContext(ctx context.Context) BuildpackBindingPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BuildpackBindingProperties) *BuildpackBindingProperties {
		return &v
	}).(BuildpackBindingPropertiesPtrOutput)
}

func (o BuildpackBindingPropertiesOutput) BindingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BuildpackBindingProperties) *string { return v.BindingType }).(pulumi.StringPtrOutput)
}

func (o BuildpackBindingPropertiesOutput) LaunchProperties() BuildpackBindingLaunchPropertiesPtrOutput {
	return o.ApplyT(func(v BuildpackBindingProperties) *BuildpackBindingLaunchProperties { return v.LaunchProperties }).(BuildpackBindingLaunchPropertiesPtrOutput)
}

type BuildpackBindingPropertiesPtrOutput struct{ *pulumi.OutputState }

func (BuildpackBindingPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildpackBindingProperties)(nil)).Elem()
}

func (o BuildpackBindingPropertiesPtrOutput) ToBuildpackBindingPropertiesPtrOutput() BuildpackBindingPropertiesPtrOutput {
	return o
}

func (o BuildpackBindingPropertiesPtrOutput) ToBuildpackBindingPropertiesPtrOutputWithContext(ctx context.Context) BuildpackBindingPropertiesPtrOutput {
	return o
}

func (o BuildpackBindingPropertiesPtrOutput) Elem() BuildpackBindingPropertiesOutput {
	return o.ApplyT(func(v *BuildpackBindingProperties) BuildpackBindingProperties {
		if v != nil {
			return *v
		}
		var ret BuildpackBindingProperties
		return ret
	}).(BuildpackBindingPropertiesOutput)
}

func (o BuildpackBindingPropertiesPtrOutput) BindingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BuildpackBindingProperties) *string {
		if v == nil {
			return nil
		}
		return v.BindingType
	}).(pulumi.StringPtrOutput)
}

func (o BuildpackBindingPropertiesPtrOutput) LaunchProperties() BuildpackBindingLaunchPropertiesPtrOutput {
	return o.ApplyT(func(v *BuildpackBindingProperties) *BuildpackBindingLaunchProperties {
		if v == nil {
			return nil
		}
		return v.LaunchProperties
	}).(BuildpackBindingLaunchPropertiesPtrOutput)
}

type BuildpackBindingPropertiesResponse struct {
	BindingType       *string                                   `pulumi:"bindingType"`
	LaunchProperties  *BuildpackBindingLaunchPropertiesResponse `pulumi:"launchProperties"`
	ProvisioningState string                                    `pulumi:"provisioningState"`
}

type BuildpackBindingPropertiesResponseOutput struct{ *pulumi.OutputState }

func (BuildpackBindingPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildpackBindingPropertiesResponse)(nil)).Elem()
}

func (o BuildpackBindingPropertiesResponseOutput) ToBuildpackBindingPropertiesResponseOutput() BuildpackBindingPropertiesResponseOutput {
	return o
}

func (o BuildpackBindingPropertiesResponseOutput) ToBuildpackBindingPropertiesResponseOutputWithContext(ctx context.Context) BuildpackBindingPropertiesResponseOutput {
	return o
}

func (o BuildpackBindingPropertiesResponseOutput) BindingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BuildpackBindingPropertiesResponse) *string { return v.BindingType }).(pulumi.StringPtrOutput)
}

func (o BuildpackBindingPropertiesResponseOutput) LaunchProperties() BuildpackBindingLaunchPropertiesResponsePtrOutput {
	return o.ApplyT(func(v BuildpackBindingPropertiesResponse) *BuildpackBindingLaunchPropertiesResponse {
		return v.LaunchProperties
	}).(BuildpackBindingLaunchPropertiesResponsePtrOutput)
}

func (o BuildpackBindingPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v BuildpackBindingPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type BuildpackProperties struct {
	Id *string `pulumi:"id"`
}





type BuildpackPropertiesInput interface {
	pulumi.Input

	ToBuildpackPropertiesOutput() BuildpackPropertiesOutput
	ToBuildpackPropertiesOutputWithContext(context.Context) BuildpackPropertiesOutput
}

type BuildpackPropertiesArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (BuildpackPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildpackProperties)(nil)).Elem()
}

func (i BuildpackPropertiesArgs) ToBuildpackPropertiesOutput() BuildpackPropertiesOutput {
	return i.ToBuildpackPropertiesOutputWithContext(context.Background())
}

func (i BuildpackPropertiesArgs) ToBuildpackPropertiesOutputWithContext(ctx context.Context) BuildpackPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildpackPropertiesOutput)
}





type BuildpackPropertiesArrayInput interface {
	pulumi.Input

	ToBuildpackPropertiesArrayOutput() BuildpackPropertiesArrayOutput
	ToBuildpackPropertiesArrayOutputWithContext(context.Context) BuildpackPropertiesArrayOutput
}

type BuildpackPropertiesArray []BuildpackPropertiesInput

func (BuildpackPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BuildpackProperties)(nil)).Elem()
}

func (i BuildpackPropertiesArray) ToBuildpackPropertiesArrayOutput() BuildpackPropertiesArrayOutput {
	return i.ToBuildpackPropertiesArrayOutputWithContext(context.Background())
}

func (i BuildpackPropertiesArray) ToBuildpackPropertiesArrayOutputWithContext(ctx context.Context) BuildpackPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildpackPropertiesArrayOutput)
}

type BuildpackPropertiesOutput struct{ *pulumi.OutputState }

func (BuildpackPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildpackProperties)(nil)).Elem()
}

func (o BuildpackPropertiesOutput) ToBuildpackPropertiesOutput() BuildpackPropertiesOutput {
	return o
}

func (o BuildpackPropertiesOutput) ToBuildpackPropertiesOutputWithContext(ctx context.Context) BuildpackPropertiesOutput {
	return o
}

func (o BuildpackPropertiesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BuildpackProperties) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type BuildpackPropertiesArrayOutput struct{ *pulumi.OutputState }

func (BuildpackPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BuildpackProperties)(nil)).Elem()
}

func (o BuildpackPropertiesArrayOutput) ToBuildpackPropertiesArrayOutput() BuildpackPropertiesArrayOutput {
	return o
}

func (o BuildpackPropertiesArrayOutput) ToBuildpackPropertiesArrayOutputWithContext(ctx context.Context) BuildpackPropertiesArrayOutput {
	return o
}

func (o BuildpackPropertiesArrayOutput) Index(i pulumi.IntInput) BuildpackPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BuildpackProperties {
		return vs[0].([]BuildpackProperties)[vs[1].(int)]
	}).(BuildpackPropertiesOutput)
}

type BuildpackPropertiesResponse struct {
	Id *string `pulumi:"id"`
}

type BuildpackPropertiesResponseOutput struct{ *pulumi.OutputState }

func (BuildpackPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildpackPropertiesResponse)(nil)).Elem()
}

func (o BuildpackPropertiesResponseOutput) ToBuildpackPropertiesResponseOutput() BuildpackPropertiesResponseOutput {
	return o
}

func (o BuildpackPropertiesResponseOutput) ToBuildpackPropertiesResponseOutputWithContext(ctx context.Context) BuildpackPropertiesResponseOutput {
	return o
}

func (o BuildpackPropertiesResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BuildpackPropertiesResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type BuildpackPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (BuildpackPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BuildpackPropertiesResponse)(nil)).Elem()
}

func (o BuildpackPropertiesResponseArrayOutput) ToBuildpackPropertiesResponseArrayOutput() BuildpackPropertiesResponseArrayOutput {
	return o
}

func (o BuildpackPropertiesResponseArrayOutput) ToBuildpackPropertiesResponseArrayOutputWithContext(ctx context.Context) BuildpackPropertiesResponseArrayOutput {
	return o
}

func (o BuildpackPropertiesResponseArrayOutput) Index(i pulumi.IntInput) BuildpackPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BuildpackPropertiesResponse {
		return vs[0].([]BuildpackPropertiesResponse)[vs[1].(int)]
	}).(BuildpackPropertiesResponseOutput)
}

type BuildpacksGroupProperties struct {
	Buildpacks []BuildpackProperties `pulumi:"buildpacks"`
	Name       *string               `pulumi:"name"`
}





type BuildpacksGroupPropertiesInput interface {
	pulumi.Input

	ToBuildpacksGroupPropertiesOutput() BuildpacksGroupPropertiesOutput
	ToBuildpacksGroupPropertiesOutputWithContext(context.Context) BuildpacksGroupPropertiesOutput
}

type BuildpacksGroupPropertiesArgs struct {
	Buildpacks BuildpackPropertiesArrayInput `pulumi:"buildpacks"`
	Name       pulumi.StringPtrInput         `pulumi:"name"`
}

func (BuildpacksGroupPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildpacksGroupProperties)(nil)).Elem()
}

func (i BuildpacksGroupPropertiesArgs) ToBuildpacksGroupPropertiesOutput() BuildpacksGroupPropertiesOutput {
	return i.ToBuildpacksGroupPropertiesOutputWithContext(context.Background())
}

func (i BuildpacksGroupPropertiesArgs) ToBuildpacksGroupPropertiesOutputWithContext(ctx context.Context) BuildpacksGroupPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildpacksGroupPropertiesOutput)
}





type BuildpacksGroupPropertiesArrayInput interface {
	pulumi.Input

	ToBuildpacksGroupPropertiesArrayOutput() BuildpacksGroupPropertiesArrayOutput
	ToBuildpacksGroupPropertiesArrayOutputWithContext(context.Context) BuildpacksGroupPropertiesArrayOutput
}

type BuildpacksGroupPropertiesArray []BuildpacksGroupPropertiesInput

func (BuildpacksGroupPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BuildpacksGroupProperties)(nil)).Elem()
}

func (i BuildpacksGroupPropertiesArray) ToBuildpacksGroupPropertiesArrayOutput() BuildpacksGroupPropertiesArrayOutput {
	return i.ToBuildpacksGroupPropertiesArrayOutputWithContext(context.Background())
}

func (i BuildpacksGroupPropertiesArray) ToBuildpacksGroupPropertiesArrayOutputWithContext(ctx context.Context) BuildpacksGroupPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildpacksGroupPropertiesArrayOutput)
}

type BuildpacksGroupPropertiesOutput struct{ *pulumi.OutputState }

func (BuildpacksGroupPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildpacksGroupProperties)(nil)).Elem()
}

func (o BuildpacksGroupPropertiesOutput) ToBuildpacksGroupPropertiesOutput() BuildpacksGroupPropertiesOutput {
	return o
}

func (o BuildpacksGroupPropertiesOutput) ToBuildpacksGroupPropertiesOutputWithContext(ctx context.Context) BuildpacksGroupPropertiesOutput {
	return o
}

func (o BuildpacksGroupPropertiesOutput) Buildpacks() BuildpackPropertiesArrayOutput {
	return o.ApplyT(func(v BuildpacksGroupProperties) []BuildpackProperties { return v.Buildpacks }).(BuildpackPropertiesArrayOutput)
}

func (o BuildpacksGroupPropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BuildpacksGroupProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type BuildpacksGroupPropertiesArrayOutput struct{ *pulumi.OutputState }

func (BuildpacksGroupPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BuildpacksGroupProperties)(nil)).Elem()
}

func (o BuildpacksGroupPropertiesArrayOutput) ToBuildpacksGroupPropertiesArrayOutput() BuildpacksGroupPropertiesArrayOutput {
	return o
}

func (o BuildpacksGroupPropertiesArrayOutput) ToBuildpacksGroupPropertiesArrayOutputWithContext(ctx context.Context) BuildpacksGroupPropertiesArrayOutput {
	return o
}

func (o BuildpacksGroupPropertiesArrayOutput) Index(i pulumi.IntInput) BuildpacksGroupPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BuildpacksGroupProperties {
		return vs[0].([]BuildpacksGroupProperties)[vs[1].(int)]
	}).(BuildpacksGroupPropertiesOutput)
}

type BuildpacksGroupPropertiesResponse struct {
	Buildpacks []BuildpackPropertiesResponse `pulumi:"buildpacks"`
	Name       *string                       `pulumi:"name"`
}

type BuildpacksGroupPropertiesResponseOutput struct{ *pulumi.OutputState }

func (BuildpacksGroupPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildpacksGroupPropertiesResponse)(nil)).Elem()
}

func (o BuildpacksGroupPropertiesResponseOutput) ToBuildpacksGroupPropertiesResponseOutput() BuildpacksGroupPropertiesResponseOutput {
	return o
}

func (o BuildpacksGroupPropertiesResponseOutput) ToBuildpacksGroupPropertiesResponseOutputWithContext(ctx context.Context) BuildpacksGroupPropertiesResponseOutput {
	return o
}

func (o BuildpacksGroupPropertiesResponseOutput) Buildpacks() BuildpackPropertiesResponseArrayOutput {
	return o.ApplyT(func(v BuildpacksGroupPropertiesResponse) []BuildpackPropertiesResponse { return v.Buildpacks }).(BuildpackPropertiesResponseArrayOutput)
}

func (o BuildpacksGroupPropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BuildpacksGroupPropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type BuildpacksGroupPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (BuildpacksGroupPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BuildpacksGroupPropertiesResponse)(nil)).Elem()
}

func (o BuildpacksGroupPropertiesResponseArrayOutput) ToBuildpacksGroupPropertiesResponseArrayOutput() BuildpacksGroupPropertiesResponseArrayOutput {
	return o
}

func (o BuildpacksGroupPropertiesResponseArrayOutput) ToBuildpacksGroupPropertiesResponseArrayOutputWithContext(ctx context.Context) BuildpacksGroupPropertiesResponseArrayOutput {
	return o
}

func (o BuildpacksGroupPropertiesResponseArrayOutput) Index(i pulumi.IntInput) BuildpacksGroupPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BuildpacksGroupPropertiesResponse {
		return vs[0].([]BuildpacksGroupPropertiesResponse)[vs[1].(int)]
	}).(BuildpacksGroupPropertiesResponseOutput)
}

type CertificateProperties struct {
	CertVersion      *string `pulumi:"certVersion"`
	KeyVaultCertName string  `pulumi:"keyVaultCertName"`
	VaultUri         string  `pulumi:"vaultUri"`
}





type CertificatePropertiesInput interface {
	pulumi.Input

	ToCertificatePropertiesOutput() CertificatePropertiesOutput
	ToCertificatePropertiesOutputWithContext(context.Context) CertificatePropertiesOutput
}

type CertificatePropertiesArgs struct {
	CertVersion      pulumi.StringPtrInput `pulumi:"certVersion"`
	KeyVaultCertName pulumi.StringInput    `pulumi:"keyVaultCertName"`
	VaultUri         pulumi.StringInput    `pulumi:"vaultUri"`
}

func (CertificatePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateProperties)(nil)).Elem()
}

func (i CertificatePropertiesArgs) ToCertificatePropertiesOutput() CertificatePropertiesOutput {
	return i.ToCertificatePropertiesOutputWithContext(context.Background())
}

func (i CertificatePropertiesArgs) ToCertificatePropertiesOutputWithContext(ctx context.Context) CertificatePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesOutput)
}

func (i CertificatePropertiesArgs) ToCertificatePropertiesPtrOutput() CertificatePropertiesPtrOutput {
	return i.ToCertificatePropertiesPtrOutputWithContext(context.Background())
}

func (i CertificatePropertiesArgs) ToCertificatePropertiesPtrOutputWithContext(ctx context.Context) CertificatePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesOutput).ToCertificatePropertiesPtrOutputWithContext(ctx)
}









type CertificatePropertiesPtrInput interface {
	pulumi.Input

	ToCertificatePropertiesPtrOutput() CertificatePropertiesPtrOutput
	ToCertificatePropertiesPtrOutputWithContext(context.Context) CertificatePropertiesPtrOutput
}

type certificatePropertiesPtrType CertificatePropertiesArgs

func CertificatePropertiesPtr(v *CertificatePropertiesArgs) CertificatePropertiesPtrInput {
	return (*certificatePropertiesPtrType)(v)
}

func (*certificatePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateProperties)(nil)).Elem()
}

func (i *certificatePropertiesPtrType) ToCertificatePropertiesPtrOutput() CertificatePropertiesPtrOutput {
	return i.ToCertificatePropertiesPtrOutputWithContext(context.Background())
}

func (i *certificatePropertiesPtrType) ToCertificatePropertiesPtrOutputWithContext(ctx context.Context) CertificatePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesPtrOutput)
}

type CertificatePropertiesOutput struct{ *pulumi.OutputState }

func (CertificatePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateProperties)(nil)).Elem()
}

func (o CertificatePropertiesOutput) ToCertificatePropertiesOutput() CertificatePropertiesOutput {
	return o
}

func (o CertificatePropertiesOutput) ToCertificatePropertiesOutputWithContext(ctx context.Context) CertificatePropertiesOutput {
	return o
}

func (o CertificatePropertiesOutput) ToCertificatePropertiesPtrOutput() CertificatePropertiesPtrOutput {
	return o.ToCertificatePropertiesPtrOutputWithContext(context.Background())
}

func (o CertificatePropertiesOutput) ToCertificatePropertiesPtrOutputWithContext(ctx context.Context) CertificatePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificateProperties) *CertificateProperties {
		return &v
	}).(CertificatePropertiesPtrOutput)
}

func (o CertificatePropertiesOutput) CertVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateProperties) *string { return v.CertVersion }).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesOutput) KeyVaultCertName() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateProperties) string { return v.KeyVaultCertName }).(pulumi.StringOutput)
}

func (o CertificatePropertiesOutput) VaultUri() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateProperties) string { return v.VaultUri }).(pulumi.StringOutput)
}

type CertificatePropertiesPtrOutput struct{ *pulumi.OutputState }

func (CertificatePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateProperties)(nil)).Elem()
}

func (o CertificatePropertiesPtrOutput) ToCertificatePropertiesPtrOutput() CertificatePropertiesPtrOutput {
	return o
}

func (o CertificatePropertiesPtrOutput) ToCertificatePropertiesPtrOutputWithContext(ctx context.Context) CertificatePropertiesPtrOutput {
	return o
}

func (o CertificatePropertiesPtrOutput) Elem() CertificatePropertiesOutput {
	return o.ApplyT(func(v *CertificateProperties) CertificateProperties {
		if v != nil {
			return *v
		}
		var ret CertificateProperties
		return ret
	}).(CertificatePropertiesOutput)
}

func (o CertificatePropertiesPtrOutput) CertVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateProperties) *string {
		if v == nil {
			return nil
		}
		return v.CertVersion
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesPtrOutput) KeyVaultCertName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultCertName
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesPtrOutput) VaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateProperties) *string {
		if v == nil {
			return nil
		}
		return &v.VaultUri
	}).(pulumi.StringPtrOutput)
}

type CertificatePropertiesResponse struct {
	ActivateDate     string   `pulumi:"activateDate"`
	CertVersion      *string  `pulumi:"certVersion"`
	DnsNames         []string `pulumi:"dnsNames"`
	ExpirationDate   string   `pulumi:"expirationDate"`
	IssuedDate       string   `pulumi:"issuedDate"`
	Issuer           string   `pulumi:"issuer"`
	KeyVaultCertName string   `pulumi:"keyVaultCertName"`
	SubjectName      string   `pulumi:"subjectName"`
	Thumbprint       string   `pulumi:"thumbprint"`
	VaultUri         string   `pulumi:"vaultUri"`
}

type CertificatePropertiesResponseOutput struct{ *pulumi.OutputState }

func (CertificatePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificatePropertiesResponse)(nil)).Elem()
}

func (o CertificatePropertiesResponseOutput) ToCertificatePropertiesResponseOutput() CertificatePropertiesResponseOutput {
	return o
}

func (o CertificatePropertiesResponseOutput) ToCertificatePropertiesResponseOutputWithContext(ctx context.Context) CertificatePropertiesResponseOutput {
	return o
}

func (o CertificatePropertiesResponseOutput) ActivateDate() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.ActivateDate }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) CertVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) *string { return v.CertVersion }).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponseOutput) DnsNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) []string { return v.DnsNames }).(pulumi.StringArrayOutput)
}

func (o CertificatePropertiesResponseOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.ExpirationDate }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) IssuedDate() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.IssuedDate }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Issuer }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) KeyVaultCertName() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.KeyVaultCertName }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) SubjectName() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.SubjectName }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) VaultUri() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.VaultUri }).(pulumi.StringOutput)
}

type ClusterResourceProperties struct {
	NetworkProfile *NetworkProfile `pulumi:"networkProfile"`
}





type ClusterResourcePropertiesInput interface {
	pulumi.Input

	ToClusterResourcePropertiesOutput() ClusterResourcePropertiesOutput
	ToClusterResourcePropertiesOutputWithContext(context.Context) ClusterResourcePropertiesOutput
}

type ClusterResourcePropertiesArgs struct {
	NetworkProfile NetworkProfilePtrInput `pulumi:"networkProfile"`
}

func (ClusterResourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterResourceProperties)(nil)).Elem()
}

func (i ClusterResourcePropertiesArgs) ToClusterResourcePropertiesOutput() ClusterResourcePropertiesOutput {
	return i.ToClusterResourcePropertiesOutputWithContext(context.Background())
}

func (i ClusterResourcePropertiesArgs) ToClusterResourcePropertiesOutputWithContext(ctx context.Context) ClusterResourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterResourcePropertiesOutput)
}

func (i ClusterResourcePropertiesArgs) ToClusterResourcePropertiesPtrOutput() ClusterResourcePropertiesPtrOutput {
	return i.ToClusterResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i ClusterResourcePropertiesArgs) ToClusterResourcePropertiesPtrOutputWithContext(ctx context.Context) ClusterResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterResourcePropertiesOutput).ToClusterResourcePropertiesPtrOutputWithContext(ctx)
}









type ClusterResourcePropertiesPtrInput interface {
	pulumi.Input

	ToClusterResourcePropertiesPtrOutput() ClusterResourcePropertiesPtrOutput
	ToClusterResourcePropertiesPtrOutputWithContext(context.Context) ClusterResourcePropertiesPtrOutput
}

type clusterResourcePropertiesPtrType ClusterResourcePropertiesArgs

func ClusterResourcePropertiesPtr(v *ClusterResourcePropertiesArgs) ClusterResourcePropertiesPtrInput {
	return (*clusterResourcePropertiesPtrType)(v)
}

func (*clusterResourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterResourceProperties)(nil)).Elem()
}

func (i *clusterResourcePropertiesPtrType) ToClusterResourcePropertiesPtrOutput() ClusterResourcePropertiesPtrOutput {
	return i.ToClusterResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *clusterResourcePropertiesPtrType) ToClusterResourcePropertiesPtrOutputWithContext(ctx context.Context) ClusterResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterResourcePropertiesPtrOutput)
}

type ClusterResourcePropertiesOutput struct{ *pulumi.OutputState }

func (ClusterResourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterResourceProperties)(nil)).Elem()
}

func (o ClusterResourcePropertiesOutput) ToClusterResourcePropertiesOutput() ClusterResourcePropertiesOutput {
	return o
}

func (o ClusterResourcePropertiesOutput) ToClusterResourcePropertiesOutputWithContext(ctx context.Context) ClusterResourcePropertiesOutput {
	return o
}

func (o ClusterResourcePropertiesOutput) ToClusterResourcePropertiesPtrOutput() ClusterResourcePropertiesPtrOutput {
	return o.ToClusterResourcePropertiesPtrOutputWithContext(context.Background())
}

func (o ClusterResourcePropertiesOutput) ToClusterResourcePropertiesPtrOutputWithContext(ctx context.Context) ClusterResourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterResourceProperties) *ClusterResourceProperties {
		return &v
	}).(ClusterResourcePropertiesPtrOutput)
}

func (o ClusterResourcePropertiesOutput) NetworkProfile() NetworkProfilePtrOutput {
	return o.ApplyT(func(v ClusterResourceProperties) *NetworkProfile { return v.NetworkProfile }).(NetworkProfilePtrOutput)
}

type ClusterResourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ClusterResourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterResourceProperties)(nil)).Elem()
}

func (o ClusterResourcePropertiesPtrOutput) ToClusterResourcePropertiesPtrOutput() ClusterResourcePropertiesPtrOutput {
	return o
}

func (o ClusterResourcePropertiesPtrOutput) ToClusterResourcePropertiesPtrOutputWithContext(ctx context.Context) ClusterResourcePropertiesPtrOutput {
	return o
}

func (o ClusterResourcePropertiesPtrOutput) Elem() ClusterResourcePropertiesOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) ClusterResourceProperties {
		if v != nil {
			return *v
		}
		var ret ClusterResourceProperties
		return ret
	}).(ClusterResourcePropertiesOutput)
}

func (o ClusterResourcePropertiesPtrOutput) NetworkProfile() NetworkProfilePtrOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) *NetworkProfile {
		if v == nil {
			return nil
		}
		return v.NetworkProfile
	}).(NetworkProfilePtrOutput)
}

type ClusterResourcePropertiesResponse struct {
	NetworkProfile    *NetworkProfileResponse `pulumi:"networkProfile"`
	ProvisioningState string                  `pulumi:"provisioningState"`
	ServiceId         string                  `pulumi:"serviceId"`
	Version           int                     `pulumi:"version"`
}

type ClusterResourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ClusterResourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterResourcePropertiesResponse)(nil)).Elem()
}

func (o ClusterResourcePropertiesResponseOutput) ToClusterResourcePropertiesResponseOutput() ClusterResourcePropertiesResponseOutput {
	return o
}

func (o ClusterResourcePropertiesResponseOutput) ToClusterResourcePropertiesResponseOutputWithContext(ctx context.Context) ClusterResourcePropertiesResponseOutput {
	return o
}

func (o ClusterResourcePropertiesResponseOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v ClusterResourcePropertiesResponse) *NetworkProfileResponse { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o ClusterResourcePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterResourcePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ClusterResourcePropertiesResponseOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterResourcePropertiesResponse) string { return v.ServiceId }).(pulumi.StringOutput)
}

func (o ClusterResourcePropertiesResponseOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v ClusterResourcePropertiesResponse) int { return v.Version }).(pulumi.IntOutput)
}

type ConfigurationServiceGitProperty struct {
	Repositories []ConfigurationServiceGitRepository `pulumi:"repositories"`
}





type ConfigurationServiceGitPropertyInput interface {
	pulumi.Input

	ToConfigurationServiceGitPropertyOutput() ConfigurationServiceGitPropertyOutput
	ToConfigurationServiceGitPropertyOutputWithContext(context.Context) ConfigurationServiceGitPropertyOutput
}

type ConfigurationServiceGitPropertyArgs struct {
	Repositories ConfigurationServiceGitRepositoryArrayInput `pulumi:"repositories"`
}

func (ConfigurationServiceGitPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationServiceGitProperty)(nil)).Elem()
}

func (i ConfigurationServiceGitPropertyArgs) ToConfigurationServiceGitPropertyOutput() ConfigurationServiceGitPropertyOutput {
	return i.ToConfigurationServiceGitPropertyOutputWithContext(context.Background())
}

func (i ConfigurationServiceGitPropertyArgs) ToConfigurationServiceGitPropertyOutputWithContext(ctx context.Context) ConfigurationServiceGitPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationServiceGitPropertyOutput)
}

func (i ConfigurationServiceGitPropertyArgs) ToConfigurationServiceGitPropertyPtrOutput() ConfigurationServiceGitPropertyPtrOutput {
	return i.ToConfigurationServiceGitPropertyPtrOutputWithContext(context.Background())
}

func (i ConfigurationServiceGitPropertyArgs) ToConfigurationServiceGitPropertyPtrOutputWithContext(ctx context.Context) ConfigurationServiceGitPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationServiceGitPropertyOutput).ToConfigurationServiceGitPropertyPtrOutputWithContext(ctx)
}









type ConfigurationServiceGitPropertyPtrInput interface {
	pulumi.Input

	ToConfigurationServiceGitPropertyPtrOutput() ConfigurationServiceGitPropertyPtrOutput
	ToConfigurationServiceGitPropertyPtrOutputWithContext(context.Context) ConfigurationServiceGitPropertyPtrOutput
}

type configurationServiceGitPropertyPtrType ConfigurationServiceGitPropertyArgs

func ConfigurationServiceGitPropertyPtr(v *ConfigurationServiceGitPropertyArgs) ConfigurationServiceGitPropertyPtrInput {
	return (*configurationServiceGitPropertyPtrType)(v)
}

func (*configurationServiceGitPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationServiceGitProperty)(nil)).Elem()
}

func (i *configurationServiceGitPropertyPtrType) ToConfigurationServiceGitPropertyPtrOutput() ConfigurationServiceGitPropertyPtrOutput {
	return i.ToConfigurationServiceGitPropertyPtrOutputWithContext(context.Background())
}

func (i *configurationServiceGitPropertyPtrType) ToConfigurationServiceGitPropertyPtrOutputWithContext(ctx context.Context) ConfigurationServiceGitPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationServiceGitPropertyPtrOutput)
}

type ConfigurationServiceGitPropertyOutput struct{ *pulumi.OutputState }

func (ConfigurationServiceGitPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationServiceGitProperty)(nil)).Elem()
}

func (o ConfigurationServiceGitPropertyOutput) ToConfigurationServiceGitPropertyOutput() ConfigurationServiceGitPropertyOutput {
	return o
}

func (o ConfigurationServiceGitPropertyOutput) ToConfigurationServiceGitPropertyOutputWithContext(ctx context.Context) ConfigurationServiceGitPropertyOutput {
	return o
}

func (o ConfigurationServiceGitPropertyOutput) ToConfigurationServiceGitPropertyPtrOutput() ConfigurationServiceGitPropertyPtrOutput {
	return o.ToConfigurationServiceGitPropertyPtrOutputWithContext(context.Background())
}

func (o ConfigurationServiceGitPropertyOutput) ToConfigurationServiceGitPropertyPtrOutputWithContext(ctx context.Context) ConfigurationServiceGitPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationServiceGitProperty) *ConfigurationServiceGitProperty {
		return &v
	}).(ConfigurationServiceGitPropertyPtrOutput)
}

func (o ConfigurationServiceGitPropertyOutput) Repositories() ConfigurationServiceGitRepositoryArrayOutput {
	return o.ApplyT(func(v ConfigurationServiceGitProperty) []ConfigurationServiceGitRepository { return v.Repositories }).(ConfigurationServiceGitRepositoryArrayOutput)
}

type ConfigurationServiceGitPropertyPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationServiceGitPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationServiceGitProperty)(nil)).Elem()
}

func (o ConfigurationServiceGitPropertyPtrOutput) ToConfigurationServiceGitPropertyPtrOutput() ConfigurationServiceGitPropertyPtrOutput {
	return o
}

func (o ConfigurationServiceGitPropertyPtrOutput) ToConfigurationServiceGitPropertyPtrOutputWithContext(ctx context.Context) ConfigurationServiceGitPropertyPtrOutput {
	return o
}

func (o ConfigurationServiceGitPropertyPtrOutput) Elem() ConfigurationServiceGitPropertyOutput {
	return o.ApplyT(func(v *ConfigurationServiceGitProperty) ConfigurationServiceGitProperty {
		if v != nil {
			return *v
		}
		var ret ConfigurationServiceGitProperty
		return ret
	}).(ConfigurationServiceGitPropertyOutput)
}

func (o ConfigurationServiceGitPropertyPtrOutput) Repositories() ConfigurationServiceGitRepositoryArrayOutput {
	return o.ApplyT(func(v *ConfigurationServiceGitProperty) []ConfigurationServiceGitRepository {
		if v == nil {
			return nil
		}
		return v.Repositories
	}).(ConfigurationServiceGitRepositoryArrayOutput)
}

type ConfigurationServiceGitPropertyResponse struct {
	Repositories []ConfigurationServiceGitRepositoryResponse `pulumi:"repositories"`
}

type ConfigurationServiceGitPropertyResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationServiceGitPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationServiceGitPropertyResponse)(nil)).Elem()
}

func (o ConfigurationServiceGitPropertyResponseOutput) ToConfigurationServiceGitPropertyResponseOutput() ConfigurationServiceGitPropertyResponseOutput {
	return o
}

func (o ConfigurationServiceGitPropertyResponseOutput) ToConfigurationServiceGitPropertyResponseOutputWithContext(ctx context.Context) ConfigurationServiceGitPropertyResponseOutput {
	return o
}

func (o ConfigurationServiceGitPropertyResponseOutput) Repositories() ConfigurationServiceGitRepositoryResponseArrayOutput {
	return o.ApplyT(func(v ConfigurationServiceGitPropertyResponse) []ConfigurationServiceGitRepositoryResponse {
		return v.Repositories
	}).(ConfigurationServiceGitRepositoryResponseArrayOutput)
}

type ConfigurationServiceGitPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationServiceGitPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationServiceGitPropertyResponse)(nil)).Elem()
}

func (o ConfigurationServiceGitPropertyResponsePtrOutput) ToConfigurationServiceGitPropertyResponsePtrOutput() ConfigurationServiceGitPropertyResponsePtrOutput {
	return o
}

func (o ConfigurationServiceGitPropertyResponsePtrOutput) ToConfigurationServiceGitPropertyResponsePtrOutputWithContext(ctx context.Context) ConfigurationServiceGitPropertyResponsePtrOutput {
	return o
}

func (o ConfigurationServiceGitPropertyResponsePtrOutput) Elem() ConfigurationServiceGitPropertyResponseOutput {
	return o.ApplyT(func(v *ConfigurationServiceGitPropertyResponse) ConfigurationServiceGitPropertyResponse {
		if v != nil {
			return *v
		}
		var ret ConfigurationServiceGitPropertyResponse
		return ret
	}).(ConfigurationServiceGitPropertyResponseOutput)
}

func (o ConfigurationServiceGitPropertyResponsePtrOutput) Repositories() ConfigurationServiceGitRepositoryResponseArrayOutput {
	return o.ApplyT(func(v *ConfigurationServiceGitPropertyResponse) []ConfigurationServiceGitRepositoryResponse {
		if v == nil {
			return nil
		}
		return v.Repositories
	}).(ConfigurationServiceGitRepositoryResponseArrayOutput)
}

type ConfigurationServiceGitRepository struct {
	HostKey               *string  `pulumi:"hostKey"`
	HostKeyAlgorithm      *string  `pulumi:"hostKeyAlgorithm"`
	Label                 string   `pulumi:"label"`
	Name                  string   `pulumi:"name"`
	Password              *string  `pulumi:"password"`
	Patterns              []string `pulumi:"patterns"`
	PrivateKey            *string  `pulumi:"privateKey"`
	SearchPaths           []string `pulumi:"searchPaths"`
	StrictHostKeyChecking *bool    `pulumi:"strictHostKeyChecking"`
	Uri                   string   `pulumi:"uri"`
	Username              *string  `pulumi:"username"`
}





type ConfigurationServiceGitRepositoryInput interface {
	pulumi.Input

	ToConfigurationServiceGitRepositoryOutput() ConfigurationServiceGitRepositoryOutput
	ToConfigurationServiceGitRepositoryOutputWithContext(context.Context) ConfigurationServiceGitRepositoryOutput
}

type ConfigurationServiceGitRepositoryArgs struct {
	HostKey               pulumi.StringPtrInput   `pulumi:"hostKey"`
	HostKeyAlgorithm      pulumi.StringPtrInput   `pulumi:"hostKeyAlgorithm"`
	Label                 pulumi.StringInput      `pulumi:"label"`
	Name                  pulumi.StringInput      `pulumi:"name"`
	Password              pulumi.StringPtrInput   `pulumi:"password"`
	Patterns              pulumi.StringArrayInput `pulumi:"patterns"`
	PrivateKey            pulumi.StringPtrInput   `pulumi:"privateKey"`
	SearchPaths           pulumi.StringArrayInput `pulumi:"searchPaths"`
	StrictHostKeyChecking pulumi.BoolPtrInput     `pulumi:"strictHostKeyChecking"`
	Uri                   pulumi.StringInput      `pulumi:"uri"`
	Username              pulumi.StringPtrInput   `pulumi:"username"`
}

func (ConfigurationServiceGitRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationServiceGitRepository)(nil)).Elem()
}

func (i ConfigurationServiceGitRepositoryArgs) ToConfigurationServiceGitRepositoryOutput() ConfigurationServiceGitRepositoryOutput {
	return i.ToConfigurationServiceGitRepositoryOutputWithContext(context.Background())
}

func (i ConfigurationServiceGitRepositoryArgs) ToConfigurationServiceGitRepositoryOutputWithContext(ctx context.Context) ConfigurationServiceGitRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationServiceGitRepositoryOutput)
}





type ConfigurationServiceGitRepositoryArrayInput interface {
	pulumi.Input

	ToConfigurationServiceGitRepositoryArrayOutput() ConfigurationServiceGitRepositoryArrayOutput
	ToConfigurationServiceGitRepositoryArrayOutputWithContext(context.Context) ConfigurationServiceGitRepositoryArrayOutput
}

type ConfigurationServiceGitRepositoryArray []ConfigurationServiceGitRepositoryInput

func (ConfigurationServiceGitRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationServiceGitRepository)(nil)).Elem()
}

func (i ConfigurationServiceGitRepositoryArray) ToConfigurationServiceGitRepositoryArrayOutput() ConfigurationServiceGitRepositoryArrayOutput {
	return i.ToConfigurationServiceGitRepositoryArrayOutputWithContext(context.Background())
}

func (i ConfigurationServiceGitRepositoryArray) ToConfigurationServiceGitRepositoryArrayOutputWithContext(ctx context.Context) ConfigurationServiceGitRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationServiceGitRepositoryArrayOutput)
}

type ConfigurationServiceGitRepositoryOutput struct{ *pulumi.OutputState }

func (ConfigurationServiceGitRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationServiceGitRepository)(nil)).Elem()
}

func (o ConfigurationServiceGitRepositoryOutput) ToConfigurationServiceGitRepositoryOutput() ConfigurationServiceGitRepositoryOutput {
	return o
}

func (o ConfigurationServiceGitRepositoryOutput) ToConfigurationServiceGitRepositoryOutputWithContext(ctx context.Context) ConfigurationServiceGitRepositoryOutput {
	return o
}

func (o ConfigurationServiceGitRepositoryOutput) HostKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepository) *string { return v.HostKey }).(pulumi.StringPtrOutput)
}

func (o ConfigurationServiceGitRepositoryOutput) HostKeyAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepository) *string { return v.HostKeyAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ConfigurationServiceGitRepositoryOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepository) string { return v.Label }).(pulumi.StringOutput)
}

func (o ConfigurationServiceGitRepositoryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepository) string { return v.Name }).(pulumi.StringOutput)
}

func (o ConfigurationServiceGitRepositoryOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepository) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ConfigurationServiceGitRepositoryOutput) Patterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepository) []string { return v.Patterns }).(pulumi.StringArrayOutput)
}

func (o ConfigurationServiceGitRepositoryOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepository) *string { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

func (o ConfigurationServiceGitRepositoryOutput) SearchPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepository) []string { return v.SearchPaths }).(pulumi.StringArrayOutput)
}

func (o ConfigurationServiceGitRepositoryOutput) StrictHostKeyChecking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepository) *bool { return v.StrictHostKeyChecking }).(pulumi.BoolPtrOutput)
}

func (o ConfigurationServiceGitRepositoryOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepository) string { return v.Uri }).(pulumi.StringOutput)
}

func (o ConfigurationServiceGitRepositoryOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepository) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type ConfigurationServiceGitRepositoryArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationServiceGitRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationServiceGitRepository)(nil)).Elem()
}

func (o ConfigurationServiceGitRepositoryArrayOutput) ToConfigurationServiceGitRepositoryArrayOutput() ConfigurationServiceGitRepositoryArrayOutput {
	return o
}

func (o ConfigurationServiceGitRepositoryArrayOutput) ToConfigurationServiceGitRepositoryArrayOutputWithContext(ctx context.Context) ConfigurationServiceGitRepositoryArrayOutput {
	return o
}

func (o ConfigurationServiceGitRepositoryArrayOutput) Index(i pulumi.IntInput) ConfigurationServiceGitRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConfigurationServiceGitRepository {
		return vs[0].([]ConfigurationServiceGitRepository)[vs[1].(int)]
	}).(ConfigurationServiceGitRepositoryOutput)
}

type ConfigurationServiceGitRepositoryResponse struct {
	HostKey               *string  `pulumi:"hostKey"`
	HostKeyAlgorithm      *string  `pulumi:"hostKeyAlgorithm"`
	Label                 string   `pulumi:"label"`
	Name                  string   `pulumi:"name"`
	Password              *string  `pulumi:"password"`
	Patterns              []string `pulumi:"patterns"`
	PrivateKey            *string  `pulumi:"privateKey"`
	SearchPaths           []string `pulumi:"searchPaths"`
	StrictHostKeyChecking *bool    `pulumi:"strictHostKeyChecking"`
	Uri                   string   `pulumi:"uri"`
	Username              *string  `pulumi:"username"`
}

type ConfigurationServiceGitRepositoryResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationServiceGitRepositoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationServiceGitRepositoryResponse)(nil)).Elem()
}

func (o ConfigurationServiceGitRepositoryResponseOutput) ToConfigurationServiceGitRepositoryResponseOutput() ConfigurationServiceGitRepositoryResponseOutput {
	return o
}

func (o ConfigurationServiceGitRepositoryResponseOutput) ToConfigurationServiceGitRepositoryResponseOutputWithContext(ctx context.Context) ConfigurationServiceGitRepositoryResponseOutput {
	return o
}

func (o ConfigurationServiceGitRepositoryResponseOutput) HostKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepositoryResponse) *string { return v.HostKey }).(pulumi.StringPtrOutput)
}

func (o ConfigurationServiceGitRepositoryResponseOutput) HostKeyAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepositoryResponse) *string { return v.HostKeyAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ConfigurationServiceGitRepositoryResponseOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepositoryResponse) string { return v.Label }).(pulumi.StringOutput)
}

func (o ConfigurationServiceGitRepositoryResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepositoryResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ConfigurationServiceGitRepositoryResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepositoryResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ConfigurationServiceGitRepositoryResponseOutput) Patterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepositoryResponse) []string { return v.Patterns }).(pulumi.StringArrayOutput)
}

func (o ConfigurationServiceGitRepositoryResponseOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepositoryResponse) *string { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

func (o ConfigurationServiceGitRepositoryResponseOutput) SearchPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepositoryResponse) []string { return v.SearchPaths }).(pulumi.StringArrayOutput)
}

func (o ConfigurationServiceGitRepositoryResponseOutput) StrictHostKeyChecking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepositoryResponse) *bool { return v.StrictHostKeyChecking }).(pulumi.BoolPtrOutput)
}

func (o ConfigurationServiceGitRepositoryResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepositoryResponse) string { return v.Uri }).(pulumi.StringOutput)
}

func (o ConfigurationServiceGitRepositoryResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationServiceGitRepositoryResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type ConfigurationServiceGitRepositoryResponseArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationServiceGitRepositoryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationServiceGitRepositoryResponse)(nil)).Elem()
}

func (o ConfigurationServiceGitRepositoryResponseArrayOutput) ToConfigurationServiceGitRepositoryResponseArrayOutput() ConfigurationServiceGitRepositoryResponseArrayOutput {
	return o
}

func (o ConfigurationServiceGitRepositoryResponseArrayOutput) ToConfigurationServiceGitRepositoryResponseArrayOutputWithContext(ctx context.Context) ConfigurationServiceGitRepositoryResponseArrayOutput {
	return o
}

func (o ConfigurationServiceGitRepositoryResponseArrayOutput) Index(i pulumi.IntInput) ConfigurationServiceGitRepositoryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConfigurationServiceGitRepositoryResponse {
		return vs[0].([]ConfigurationServiceGitRepositoryResponse)[vs[1].(int)]
	}).(ConfigurationServiceGitRepositoryResponseOutput)
}

type ConfigurationServiceInstanceResponse struct {
	Name   string `pulumi:"name"`
	Status string `pulumi:"status"`
}

type ConfigurationServiceInstanceResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationServiceInstanceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationServiceInstanceResponse)(nil)).Elem()
}

func (o ConfigurationServiceInstanceResponseOutput) ToConfigurationServiceInstanceResponseOutput() ConfigurationServiceInstanceResponseOutput {
	return o
}

func (o ConfigurationServiceInstanceResponseOutput) ToConfigurationServiceInstanceResponseOutputWithContext(ctx context.Context) ConfigurationServiceInstanceResponseOutput {
	return o
}

func (o ConfigurationServiceInstanceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationServiceInstanceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ConfigurationServiceInstanceResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationServiceInstanceResponse) string { return v.Status }).(pulumi.StringOutput)
}

type ConfigurationServiceInstanceResponseArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationServiceInstanceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationServiceInstanceResponse)(nil)).Elem()
}

func (o ConfigurationServiceInstanceResponseArrayOutput) ToConfigurationServiceInstanceResponseArrayOutput() ConfigurationServiceInstanceResponseArrayOutput {
	return o
}

func (o ConfigurationServiceInstanceResponseArrayOutput) ToConfigurationServiceInstanceResponseArrayOutputWithContext(ctx context.Context) ConfigurationServiceInstanceResponseArrayOutput {
	return o
}

func (o ConfigurationServiceInstanceResponseArrayOutput) Index(i pulumi.IntInput) ConfigurationServiceInstanceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConfigurationServiceInstanceResponse {
		return vs[0].([]ConfigurationServiceInstanceResponse)[vs[1].(int)]
	}).(ConfigurationServiceInstanceResponseOutput)
}

type ConfigurationServiceProperties struct {
	Settings *ConfigurationServiceSettings `pulumi:"settings"`
}





type ConfigurationServicePropertiesInput interface {
	pulumi.Input

	ToConfigurationServicePropertiesOutput() ConfigurationServicePropertiesOutput
	ToConfigurationServicePropertiesOutputWithContext(context.Context) ConfigurationServicePropertiesOutput
}

type ConfigurationServicePropertiesArgs struct {
	Settings ConfigurationServiceSettingsPtrInput `pulumi:"settings"`
}

func (ConfigurationServicePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationServiceProperties)(nil)).Elem()
}

func (i ConfigurationServicePropertiesArgs) ToConfigurationServicePropertiesOutput() ConfigurationServicePropertiesOutput {
	return i.ToConfigurationServicePropertiesOutputWithContext(context.Background())
}

func (i ConfigurationServicePropertiesArgs) ToConfigurationServicePropertiesOutputWithContext(ctx context.Context) ConfigurationServicePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationServicePropertiesOutput)
}

func (i ConfigurationServicePropertiesArgs) ToConfigurationServicePropertiesPtrOutput() ConfigurationServicePropertiesPtrOutput {
	return i.ToConfigurationServicePropertiesPtrOutputWithContext(context.Background())
}

func (i ConfigurationServicePropertiesArgs) ToConfigurationServicePropertiesPtrOutputWithContext(ctx context.Context) ConfigurationServicePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationServicePropertiesOutput).ToConfigurationServicePropertiesPtrOutputWithContext(ctx)
}









type ConfigurationServicePropertiesPtrInput interface {
	pulumi.Input

	ToConfigurationServicePropertiesPtrOutput() ConfigurationServicePropertiesPtrOutput
	ToConfigurationServicePropertiesPtrOutputWithContext(context.Context) ConfigurationServicePropertiesPtrOutput
}

type configurationServicePropertiesPtrType ConfigurationServicePropertiesArgs

func ConfigurationServicePropertiesPtr(v *ConfigurationServicePropertiesArgs) ConfigurationServicePropertiesPtrInput {
	return (*configurationServicePropertiesPtrType)(v)
}

func (*configurationServicePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationServiceProperties)(nil)).Elem()
}

func (i *configurationServicePropertiesPtrType) ToConfigurationServicePropertiesPtrOutput() ConfigurationServicePropertiesPtrOutput {
	return i.ToConfigurationServicePropertiesPtrOutputWithContext(context.Background())
}

func (i *configurationServicePropertiesPtrType) ToConfigurationServicePropertiesPtrOutputWithContext(ctx context.Context) ConfigurationServicePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationServicePropertiesPtrOutput)
}

type ConfigurationServicePropertiesOutput struct{ *pulumi.OutputState }

func (ConfigurationServicePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationServiceProperties)(nil)).Elem()
}

func (o ConfigurationServicePropertiesOutput) ToConfigurationServicePropertiesOutput() ConfigurationServicePropertiesOutput {
	return o
}

func (o ConfigurationServicePropertiesOutput) ToConfigurationServicePropertiesOutputWithContext(ctx context.Context) ConfigurationServicePropertiesOutput {
	return o
}

func (o ConfigurationServicePropertiesOutput) ToConfigurationServicePropertiesPtrOutput() ConfigurationServicePropertiesPtrOutput {
	return o.ToConfigurationServicePropertiesPtrOutputWithContext(context.Background())
}

func (o ConfigurationServicePropertiesOutput) ToConfigurationServicePropertiesPtrOutputWithContext(ctx context.Context) ConfigurationServicePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationServiceProperties) *ConfigurationServiceProperties {
		return &v
	}).(ConfigurationServicePropertiesPtrOutput)
}

func (o ConfigurationServicePropertiesOutput) Settings() ConfigurationServiceSettingsPtrOutput {
	return o.ApplyT(func(v ConfigurationServiceProperties) *ConfigurationServiceSettings { return v.Settings }).(ConfigurationServiceSettingsPtrOutput)
}

type ConfigurationServicePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationServicePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationServiceProperties)(nil)).Elem()
}

func (o ConfigurationServicePropertiesPtrOutput) ToConfigurationServicePropertiesPtrOutput() ConfigurationServicePropertiesPtrOutput {
	return o
}

func (o ConfigurationServicePropertiesPtrOutput) ToConfigurationServicePropertiesPtrOutputWithContext(ctx context.Context) ConfigurationServicePropertiesPtrOutput {
	return o
}

func (o ConfigurationServicePropertiesPtrOutput) Elem() ConfigurationServicePropertiesOutput {
	return o.ApplyT(func(v *ConfigurationServiceProperties) ConfigurationServiceProperties {
		if v != nil {
			return *v
		}
		var ret ConfigurationServiceProperties
		return ret
	}).(ConfigurationServicePropertiesOutput)
}

func (o ConfigurationServicePropertiesPtrOutput) Settings() ConfigurationServiceSettingsPtrOutput {
	return o.ApplyT(func(v *ConfigurationServiceProperties) *ConfigurationServiceSettings {
		if v == nil {
			return nil
		}
		return v.Settings
	}).(ConfigurationServiceSettingsPtrOutput)
}

type ConfigurationServicePropertiesResponse struct {
	Instances         []ConfigurationServiceInstanceResponse       `pulumi:"instances"`
	ProvisioningState string                                       `pulumi:"provisioningState"`
	ResourceRequests  ConfigurationServiceResourceRequestsResponse `pulumi:"resourceRequests"`
	Settings          *ConfigurationServiceSettingsResponse        `pulumi:"settings"`
}

type ConfigurationServicePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationServicePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationServicePropertiesResponse)(nil)).Elem()
}

func (o ConfigurationServicePropertiesResponseOutput) ToConfigurationServicePropertiesResponseOutput() ConfigurationServicePropertiesResponseOutput {
	return o
}

func (o ConfigurationServicePropertiesResponseOutput) ToConfigurationServicePropertiesResponseOutputWithContext(ctx context.Context) ConfigurationServicePropertiesResponseOutput {
	return o
}

func (o ConfigurationServicePropertiesResponseOutput) Instances() ConfigurationServiceInstanceResponseArrayOutput {
	return o.ApplyT(func(v ConfigurationServicePropertiesResponse) []ConfigurationServiceInstanceResponse {
		return v.Instances
	}).(ConfigurationServiceInstanceResponseArrayOutput)
}

func (o ConfigurationServicePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationServicePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ConfigurationServicePropertiesResponseOutput) ResourceRequests() ConfigurationServiceResourceRequestsResponseOutput {
	return o.ApplyT(func(v ConfigurationServicePropertiesResponse) ConfigurationServiceResourceRequestsResponse {
		return v.ResourceRequests
	}).(ConfigurationServiceResourceRequestsResponseOutput)
}

func (o ConfigurationServicePropertiesResponseOutput) Settings() ConfigurationServiceSettingsResponsePtrOutput {
	return o.ApplyT(func(v ConfigurationServicePropertiesResponse) *ConfigurationServiceSettingsResponse {
		return v.Settings
	}).(ConfigurationServiceSettingsResponsePtrOutput)
}

type ConfigurationServiceResourceRequestsResponse struct {
	Cpu           string `pulumi:"cpu"`
	InstanceCount int    `pulumi:"instanceCount"`
	Memory        string `pulumi:"memory"`
}

type ConfigurationServiceResourceRequestsResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationServiceResourceRequestsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationServiceResourceRequestsResponse)(nil)).Elem()
}

func (o ConfigurationServiceResourceRequestsResponseOutput) ToConfigurationServiceResourceRequestsResponseOutput() ConfigurationServiceResourceRequestsResponseOutput {
	return o
}

func (o ConfigurationServiceResourceRequestsResponseOutput) ToConfigurationServiceResourceRequestsResponseOutputWithContext(ctx context.Context) ConfigurationServiceResourceRequestsResponseOutput {
	return o
}

func (o ConfigurationServiceResourceRequestsResponseOutput) Cpu() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationServiceResourceRequestsResponse) string { return v.Cpu }).(pulumi.StringOutput)
}

func (o ConfigurationServiceResourceRequestsResponseOutput) InstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v ConfigurationServiceResourceRequestsResponse) int { return v.InstanceCount }).(pulumi.IntOutput)
}

func (o ConfigurationServiceResourceRequestsResponseOutput) Memory() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationServiceResourceRequestsResponse) string { return v.Memory }).(pulumi.StringOutput)
}

type ConfigurationServiceSettings struct {
	GitProperty *ConfigurationServiceGitProperty `pulumi:"gitProperty"`
}





type ConfigurationServiceSettingsInput interface {
	pulumi.Input

	ToConfigurationServiceSettingsOutput() ConfigurationServiceSettingsOutput
	ToConfigurationServiceSettingsOutputWithContext(context.Context) ConfigurationServiceSettingsOutput
}

type ConfigurationServiceSettingsArgs struct {
	GitProperty ConfigurationServiceGitPropertyPtrInput `pulumi:"gitProperty"`
}

func (ConfigurationServiceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationServiceSettings)(nil)).Elem()
}

func (i ConfigurationServiceSettingsArgs) ToConfigurationServiceSettingsOutput() ConfigurationServiceSettingsOutput {
	return i.ToConfigurationServiceSettingsOutputWithContext(context.Background())
}

func (i ConfigurationServiceSettingsArgs) ToConfigurationServiceSettingsOutputWithContext(ctx context.Context) ConfigurationServiceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationServiceSettingsOutput)
}

func (i ConfigurationServiceSettingsArgs) ToConfigurationServiceSettingsPtrOutput() ConfigurationServiceSettingsPtrOutput {
	return i.ToConfigurationServiceSettingsPtrOutputWithContext(context.Background())
}

func (i ConfigurationServiceSettingsArgs) ToConfigurationServiceSettingsPtrOutputWithContext(ctx context.Context) ConfigurationServiceSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationServiceSettingsOutput).ToConfigurationServiceSettingsPtrOutputWithContext(ctx)
}









type ConfigurationServiceSettingsPtrInput interface {
	pulumi.Input

	ToConfigurationServiceSettingsPtrOutput() ConfigurationServiceSettingsPtrOutput
	ToConfigurationServiceSettingsPtrOutputWithContext(context.Context) ConfigurationServiceSettingsPtrOutput
}

type configurationServiceSettingsPtrType ConfigurationServiceSettingsArgs

func ConfigurationServiceSettingsPtr(v *ConfigurationServiceSettingsArgs) ConfigurationServiceSettingsPtrInput {
	return (*configurationServiceSettingsPtrType)(v)
}

func (*configurationServiceSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationServiceSettings)(nil)).Elem()
}

func (i *configurationServiceSettingsPtrType) ToConfigurationServiceSettingsPtrOutput() ConfigurationServiceSettingsPtrOutput {
	return i.ToConfigurationServiceSettingsPtrOutputWithContext(context.Background())
}

func (i *configurationServiceSettingsPtrType) ToConfigurationServiceSettingsPtrOutputWithContext(ctx context.Context) ConfigurationServiceSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationServiceSettingsPtrOutput)
}

type ConfigurationServiceSettingsOutput struct{ *pulumi.OutputState }

func (ConfigurationServiceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationServiceSettings)(nil)).Elem()
}

func (o ConfigurationServiceSettingsOutput) ToConfigurationServiceSettingsOutput() ConfigurationServiceSettingsOutput {
	return o
}

func (o ConfigurationServiceSettingsOutput) ToConfigurationServiceSettingsOutputWithContext(ctx context.Context) ConfigurationServiceSettingsOutput {
	return o
}

func (o ConfigurationServiceSettingsOutput) ToConfigurationServiceSettingsPtrOutput() ConfigurationServiceSettingsPtrOutput {
	return o.ToConfigurationServiceSettingsPtrOutputWithContext(context.Background())
}

func (o ConfigurationServiceSettingsOutput) ToConfigurationServiceSettingsPtrOutputWithContext(ctx context.Context) ConfigurationServiceSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationServiceSettings) *ConfigurationServiceSettings {
		return &v
	}).(ConfigurationServiceSettingsPtrOutput)
}

func (o ConfigurationServiceSettingsOutput) GitProperty() ConfigurationServiceGitPropertyPtrOutput {
	return o.ApplyT(func(v ConfigurationServiceSettings) *ConfigurationServiceGitProperty { return v.GitProperty }).(ConfigurationServiceGitPropertyPtrOutput)
}

type ConfigurationServiceSettingsPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationServiceSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationServiceSettings)(nil)).Elem()
}

func (o ConfigurationServiceSettingsPtrOutput) ToConfigurationServiceSettingsPtrOutput() ConfigurationServiceSettingsPtrOutput {
	return o
}

func (o ConfigurationServiceSettingsPtrOutput) ToConfigurationServiceSettingsPtrOutputWithContext(ctx context.Context) ConfigurationServiceSettingsPtrOutput {
	return o
}

func (o ConfigurationServiceSettingsPtrOutput) Elem() ConfigurationServiceSettingsOutput {
	return o.ApplyT(func(v *ConfigurationServiceSettings) ConfigurationServiceSettings {
		if v != nil {
			return *v
		}
		var ret ConfigurationServiceSettings
		return ret
	}).(ConfigurationServiceSettingsOutput)
}

func (o ConfigurationServiceSettingsPtrOutput) GitProperty() ConfigurationServiceGitPropertyPtrOutput {
	return o.ApplyT(func(v *ConfigurationServiceSettings) *ConfigurationServiceGitProperty {
		if v == nil {
			return nil
		}
		return v.GitProperty
	}).(ConfigurationServiceGitPropertyPtrOutput)
}

type ConfigurationServiceSettingsResponse struct {
	GitProperty *ConfigurationServiceGitPropertyResponse `pulumi:"gitProperty"`
}

type ConfigurationServiceSettingsResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationServiceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationServiceSettingsResponse)(nil)).Elem()
}

func (o ConfigurationServiceSettingsResponseOutput) ToConfigurationServiceSettingsResponseOutput() ConfigurationServiceSettingsResponseOutput {
	return o
}

func (o ConfigurationServiceSettingsResponseOutput) ToConfigurationServiceSettingsResponseOutputWithContext(ctx context.Context) ConfigurationServiceSettingsResponseOutput {
	return o
}

func (o ConfigurationServiceSettingsResponseOutput) GitProperty() ConfigurationServiceGitPropertyResponsePtrOutput {
	return o.ApplyT(func(v ConfigurationServiceSettingsResponse) *ConfigurationServiceGitPropertyResponse {
		return v.GitProperty
	}).(ConfigurationServiceGitPropertyResponsePtrOutput)
}

type ConfigurationServiceSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationServiceSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationServiceSettingsResponse)(nil)).Elem()
}

func (o ConfigurationServiceSettingsResponsePtrOutput) ToConfigurationServiceSettingsResponsePtrOutput() ConfigurationServiceSettingsResponsePtrOutput {
	return o
}

func (o ConfigurationServiceSettingsResponsePtrOutput) ToConfigurationServiceSettingsResponsePtrOutputWithContext(ctx context.Context) ConfigurationServiceSettingsResponsePtrOutput {
	return o
}

func (o ConfigurationServiceSettingsResponsePtrOutput) Elem() ConfigurationServiceSettingsResponseOutput {
	return o.ApplyT(func(v *ConfigurationServiceSettingsResponse) ConfigurationServiceSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ConfigurationServiceSettingsResponse
		return ret
	}).(ConfigurationServiceSettingsResponseOutput)
}

func (o ConfigurationServiceSettingsResponsePtrOutput) GitProperty() ConfigurationServiceGitPropertyResponsePtrOutput {
	return o.ApplyT(func(v *ConfigurationServiceSettingsResponse) *ConfigurationServiceGitPropertyResponse {
		if v == nil {
			return nil
		}
		return v.GitProperty
	}).(ConfigurationServiceGitPropertyResponsePtrOutput)
}

type CustomDomainProperties struct {
	CertName   *string `pulumi:"certName"`
	Thumbprint *string `pulumi:"thumbprint"`
}





type CustomDomainPropertiesInput interface {
	pulumi.Input

	ToCustomDomainPropertiesOutput() CustomDomainPropertiesOutput
	ToCustomDomainPropertiesOutputWithContext(context.Context) CustomDomainPropertiesOutput
}

type CustomDomainPropertiesArgs struct {
	CertName   pulumi.StringPtrInput `pulumi:"certName"`
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (CustomDomainPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomainProperties)(nil)).Elem()
}

func (i CustomDomainPropertiesArgs) ToCustomDomainPropertiesOutput() CustomDomainPropertiesOutput {
	return i.ToCustomDomainPropertiesOutputWithContext(context.Background())
}

func (i CustomDomainPropertiesArgs) ToCustomDomainPropertiesOutputWithContext(ctx context.Context) CustomDomainPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainPropertiesOutput)
}

func (i CustomDomainPropertiesArgs) ToCustomDomainPropertiesPtrOutput() CustomDomainPropertiesPtrOutput {
	return i.ToCustomDomainPropertiesPtrOutputWithContext(context.Background())
}

func (i CustomDomainPropertiesArgs) ToCustomDomainPropertiesPtrOutputWithContext(ctx context.Context) CustomDomainPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainPropertiesOutput).ToCustomDomainPropertiesPtrOutputWithContext(ctx)
}









type CustomDomainPropertiesPtrInput interface {
	pulumi.Input

	ToCustomDomainPropertiesPtrOutput() CustomDomainPropertiesPtrOutput
	ToCustomDomainPropertiesPtrOutputWithContext(context.Context) CustomDomainPropertiesPtrOutput
}

type customDomainPropertiesPtrType CustomDomainPropertiesArgs

func CustomDomainPropertiesPtr(v *CustomDomainPropertiesArgs) CustomDomainPropertiesPtrInput {
	return (*customDomainPropertiesPtrType)(v)
}

func (*customDomainPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomainProperties)(nil)).Elem()
}

func (i *customDomainPropertiesPtrType) ToCustomDomainPropertiesPtrOutput() CustomDomainPropertiesPtrOutput {
	return i.ToCustomDomainPropertiesPtrOutputWithContext(context.Background())
}

func (i *customDomainPropertiesPtrType) ToCustomDomainPropertiesPtrOutputWithContext(ctx context.Context) CustomDomainPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainPropertiesPtrOutput)
}

type CustomDomainPropertiesOutput struct{ *pulumi.OutputState }

func (CustomDomainPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomainProperties)(nil)).Elem()
}

func (o CustomDomainPropertiesOutput) ToCustomDomainPropertiesOutput() CustomDomainPropertiesOutput {
	return o
}

func (o CustomDomainPropertiesOutput) ToCustomDomainPropertiesOutputWithContext(ctx context.Context) CustomDomainPropertiesOutput {
	return o
}

func (o CustomDomainPropertiesOutput) ToCustomDomainPropertiesPtrOutput() CustomDomainPropertiesPtrOutput {
	return o.ToCustomDomainPropertiesPtrOutputWithContext(context.Background())
}

func (o CustomDomainPropertiesOutput) ToCustomDomainPropertiesPtrOutputWithContext(ctx context.Context) CustomDomainPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomDomainProperties) *CustomDomainProperties {
		return &v
	}).(CustomDomainPropertiesPtrOutput)
}

func (o CustomDomainPropertiesOutput) CertName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomDomainProperties) *string { return v.CertName }).(pulumi.StringPtrOutput)
}

func (o CustomDomainPropertiesOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomDomainProperties) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type CustomDomainPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CustomDomainPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomainProperties)(nil)).Elem()
}

func (o CustomDomainPropertiesPtrOutput) ToCustomDomainPropertiesPtrOutput() CustomDomainPropertiesPtrOutput {
	return o
}

func (o CustomDomainPropertiesPtrOutput) ToCustomDomainPropertiesPtrOutputWithContext(ctx context.Context) CustomDomainPropertiesPtrOutput {
	return o
}

func (o CustomDomainPropertiesPtrOutput) Elem() CustomDomainPropertiesOutput {
	return o.ApplyT(func(v *CustomDomainProperties) CustomDomainProperties {
		if v != nil {
			return *v
		}
		var ret CustomDomainProperties
		return ret
	}).(CustomDomainPropertiesOutput)
}

func (o CustomDomainPropertiesPtrOutput) CertName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDomainProperties) *string {
		if v == nil {
			return nil
		}
		return v.CertName
	}).(pulumi.StringPtrOutput)
}

func (o CustomDomainPropertiesPtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDomainProperties) *string {
		if v == nil {
			return nil
		}
		return v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

type CustomDomainPropertiesResponse struct {
	AppName    string  `pulumi:"appName"`
	CertName   *string `pulumi:"certName"`
	Thumbprint *string `pulumi:"thumbprint"`
}

type CustomDomainPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CustomDomainPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomainPropertiesResponse)(nil)).Elem()
}

func (o CustomDomainPropertiesResponseOutput) ToCustomDomainPropertiesResponseOutput() CustomDomainPropertiesResponseOutput {
	return o
}

func (o CustomDomainPropertiesResponseOutput) ToCustomDomainPropertiesResponseOutputWithContext(ctx context.Context) CustomDomainPropertiesResponseOutput {
	return o
}

func (o CustomDomainPropertiesResponseOutput) AppName() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDomainPropertiesResponse) string { return v.AppName }).(pulumi.StringOutput)
}

func (o CustomDomainPropertiesResponseOutput) CertName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomDomainPropertiesResponse) *string { return v.CertName }).(pulumi.StringPtrOutput)
}

func (o CustomDomainPropertiesResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomDomainPropertiesResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type DeploymentInstanceResponse struct {
	DiscoveryStatus string `pulumi:"discoveryStatus"`
	Name            string `pulumi:"name"`
	Reason          string `pulumi:"reason"`
	StartTime       string `pulumi:"startTime"`
	Status          string `pulumi:"status"`
}

type DeploymentInstanceResponseOutput struct{ *pulumi.OutputState }

func (DeploymentInstanceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentInstanceResponse)(nil)).Elem()
}

func (o DeploymentInstanceResponseOutput) ToDeploymentInstanceResponseOutput() DeploymentInstanceResponseOutput {
	return o
}

func (o DeploymentInstanceResponseOutput) ToDeploymentInstanceResponseOutputWithContext(ctx context.Context) DeploymentInstanceResponseOutput {
	return o
}

func (o DeploymentInstanceResponseOutput) DiscoveryStatus() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentInstanceResponse) string { return v.DiscoveryStatus }).(pulumi.StringOutput)
}

func (o DeploymentInstanceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentInstanceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeploymentInstanceResponseOutput) Reason() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentInstanceResponse) string { return v.Reason }).(pulumi.StringOutput)
}

func (o DeploymentInstanceResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentInstanceResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o DeploymentInstanceResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentInstanceResponse) string { return v.Status }).(pulumi.StringOutput)
}

type DeploymentInstanceResponseArrayOutput struct{ *pulumi.OutputState }

func (DeploymentInstanceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeploymentInstanceResponse)(nil)).Elem()
}

func (o DeploymentInstanceResponseArrayOutput) ToDeploymentInstanceResponseArrayOutput() DeploymentInstanceResponseArrayOutput {
	return o
}

func (o DeploymentInstanceResponseArrayOutput) ToDeploymentInstanceResponseArrayOutputWithContext(ctx context.Context) DeploymentInstanceResponseArrayOutput {
	return o
}

func (o DeploymentInstanceResponseArrayOutput) Index(i pulumi.IntInput) DeploymentInstanceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeploymentInstanceResponse {
		return vs[0].([]DeploymentInstanceResponse)[vs[1].(int)]
	}).(DeploymentInstanceResponseOutput)
}

type DeploymentResourceProperties struct {
	DeploymentSettings *DeploymentSettings `pulumi:"deploymentSettings"`
	Source             *UserSourceInfo     `pulumi:"source"`
}


func (val *DeploymentResourceProperties) Defaults() *DeploymentResourceProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DeploymentSettings = tmp.DeploymentSettings.Defaults()

	return &tmp
}





type DeploymentResourcePropertiesInput interface {
	pulumi.Input

	ToDeploymentResourcePropertiesOutput() DeploymentResourcePropertiesOutput
	ToDeploymentResourcePropertiesOutputWithContext(context.Context) DeploymentResourcePropertiesOutput
}

type DeploymentResourcePropertiesArgs struct {
	DeploymentSettings DeploymentSettingsPtrInput `pulumi:"deploymentSettings"`
	Source             UserSourceInfoPtrInput     `pulumi:"source"`
}


func (val *DeploymentResourcePropertiesArgs) Defaults() *DeploymentResourcePropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	return &tmp
}
func (DeploymentResourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentResourceProperties)(nil)).Elem()
}

func (i DeploymentResourcePropertiesArgs) ToDeploymentResourcePropertiesOutput() DeploymentResourcePropertiesOutput {
	return i.ToDeploymentResourcePropertiesOutputWithContext(context.Background())
}

func (i DeploymentResourcePropertiesArgs) ToDeploymentResourcePropertiesOutputWithContext(ctx context.Context) DeploymentResourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentResourcePropertiesOutput)
}

func (i DeploymentResourcePropertiesArgs) ToDeploymentResourcePropertiesPtrOutput() DeploymentResourcePropertiesPtrOutput {
	return i.ToDeploymentResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i DeploymentResourcePropertiesArgs) ToDeploymentResourcePropertiesPtrOutputWithContext(ctx context.Context) DeploymentResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentResourcePropertiesOutput).ToDeploymentResourcePropertiesPtrOutputWithContext(ctx)
}









type DeploymentResourcePropertiesPtrInput interface {
	pulumi.Input

	ToDeploymentResourcePropertiesPtrOutput() DeploymentResourcePropertiesPtrOutput
	ToDeploymentResourcePropertiesPtrOutputWithContext(context.Context) DeploymentResourcePropertiesPtrOutput
}

type deploymentResourcePropertiesPtrType DeploymentResourcePropertiesArgs

func DeploymentResourcePropertiesPtr(v *DeploymentResourcePropertiesArgs) DeploymentResourcePropertiesPtrInput {
	return (*deploymentResourcePropertiesPtrType)(v)
}

func (*deploymentResourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentResourceProperties)(nil)).Elem()
}

func (i *deploymentResourcePropertiesPtrType) ToDeploymentResourcePropertiesPtrOutput() DeploymentResourcePropertiesPtrOutput {
	return i.ToDeploymentResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *deploymentResourcePropertiesPtrType) ToDeploymentResourcePropertiesPtrOutputWithContext(ctx context.Context) DeploymentResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentResourcePropertiesPtrOutput)
}

type DeploymentResourcePropertiesOutput struct{ *pulumi.OutputState }

func (DeploymentResourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentResourceProperties)(nil)).Elem()
}

func (o DeploymentResourcePropertiesOutput) ToDeploymentResourcePropertiesOutput() DeploymentResourcePropertiesOutput {
	return o
}

func (o DeploymentResourcePropertiesOutput) ToDeploymentResourcePropertiesOutputWithContext(ctx context.Context) DeploymentResourcePropertiesOutput {
	return o
}

func (o DeploymentResourcePropertiesOutput) ToDeploymentResourcePropertiesPtrOutput() DeploymentResourcePropertiesPtrOutput {
	return o.ToDeploymentResourcePropertiesPtrOutputWithContext(context.Background())
}

func (o DeploymentResourcePropertiesOutput) ToDeploymentResourcePropertiesPtrOutputWithContext(ctx context.Context) DeploymentResourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentResourceProperties) *DeploymentResourceProperties {
		return &v
	}).(DeploymentResourcePropertiesPtrOutput)
}

func (o DeploymentResourcePropertiesOutput) DeploymentSettings() DeploymentSettingsPtrOutput {
	return o.ApplyT(func(v DeploymentResourceProperties) *DeploymentSettings { return v.DeploymentSettings }).(DeploymentSettingsPtrOutput)
}

func (o DeploymentResourcePropertiesOutput) Source() UserSourceInfoPtrOutput {
	return o.ApplyT(func(v DeploymentResourceProperties) *UserSourceInfo { return v.Source }).(UserSourceInfoPtrOutput)
}

type DeploymentResourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (DeploymentResourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentResourceProperties)(nil)).Elem()
}

func (o DeploymentResourcePropertiesPtrOutput) ToDeploymentResourcePropertiesPtrOutput() DeploymentResourcePropertiesPtrOutput {
	return o
}

func (o DeploymentResourcePropertiesPtrOutput) ToDeploymentResourcePropertiesPtrOutputWithContext(ctx context.Context) DeploymentResourcePropertiesPtrOutput {
	return o
}

func (o DeploymentResourcePropertiesPtrOutput) Elem() DeploymentResourcePropertiesOutput {
	return o.ApplyT(func(v *DeploymentResourceProperties) DeploymentResourceProperties {
		if v != nil {
			return *v
		}
		var ret DeploymentResourceProperties
		return ret
	}).(DeploymentResourcePropertiesOutput)
}

func (o DeploymentResourcePropertiesPtrOutput) DeploymentSettings() DeploymentSettingsPtrOutput {
	return o.ApplyT(func(v *DeploymentResourceProperties) *DeploymentSettings {
		if v == nil {
			return nil
		}
		return v.DeploymentSettings
	}).(DeploymentSettingsPtrOutput)
}

func (o DeploymentResourcePropertiesPtrOutput) Source() UserSourceInfoPtrOutput {
	return o.ApplyT(func(v *DeploymentResourceProperties) *UserSourceInfo {
		if v == nil {
			return nil
		}
		return v.Source
	}).(UserSourceInfoPtrOutput)
}

type DeploymentResourcePropertiesResponse struct {
	Active             bool                         `pulumi:"active"`
	AppName            string                       `pulumi:"appName"`
	CreatedTime        string                       `pulumi:"createdTime"`
	DeploymentSettings *DeploymentSettingsResponse  `pulumi:"deploymentSettings"`
	Instances          []DeploymentInstanceResponse `pulumi:"instances"`
	ProvisioningState  string                       `pulumi:"provisioningState"`
	Source             *UserSourceInfoResponse      `pulumi:"source"`
	Status             string                       `pulumi:"status"`
}


func (val *DeploymentResourcePropertiesResponse) Defaults() *DeploymentResourcePropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DeploymentSettings = tmp.DeploymentSettings.Defaults()

	return &tmp
}

type DeploymentResourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (DeploymentResourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentResourcePropertiesResponse)(nil)).Elem()
}

func (o DeploymentResourcePropertiesResponseOutput) ToDeploymentResourcePropertiesResponseOutput() DeploymentResourcePropertiesResponseOutput {
	return o
}

func (o DeploymentResourcePropertiesResponseOutput) ToDeploymentResourcePropertiesResponseOutputWithContext(ctx context.Context) DeploymentResourcePropertiesResponseOutput {
	return o
}

func (o DeploymentResourcePropertiesResponseOutput) Active() pulumi.BoolOutput {
	return o.ApplyT(func(v DeploymentResourcePropertiesResponse) bool { return v.Active }).(pulumi.BoolOutput)
}

func (o DeploymentResourcePropertiesResponseOutput) AppName() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentResourcePropertiesResponse) string { return v.AppName }).(pulumi.StringOutput)
}

func (o DeploymentResourcePropertiesResponseOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentResourcePropertiesResponse) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o DeploymentResourcePropertiesResponseOutput) DeploymentSettings() DeploymentSettingsResponsePtrOutput {
	return o.ApplyT(func(v DeploymentResourcePropertiesResponse) *DeploymentSettingsResponse { return v.DeploymentSettings }).(DeploymentSettingsResponsePtrOutput)
}

func (o DeploymentResourcePropertiesResponseOutput) Instances() DeploymentInstanceResponseArrayOutput {
	return o.ApplyT(func(v DeploymentResourcePropertiesResponse) []DeploymentInstanceResponse { return v.Instances }).(DeploymentInstanceResponseArrayOutput)
}

func (o DeploymentResourcePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentResourcePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DeploymentResourcePropertiesResponseOutput) Source() UserSourceInfoResponsePtrOutput {
	return o.ApplyT(func(v DeploymentResourcePropertiesResponse) *UserSourceInfoResponse { return v.Source }).(UserSourceInfoResponsePtrOutput)
}

func (o DeploymentResourcePropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentResourcePropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

type DeploymentSettings struct {
	Cpu                  *int              `pulumi:"cpu"`
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	JvmOptions           *string           `pulumi:"jvmOptions"`
	MemoryInGB           *int              `pulumi:"memoryInGB"`
	NetCoreMainEntryPath *string           `pulumi:"netCoreMainEntryPath"`
	RuntimeVersion       *string           `pulumi:"runtimeVersion"`
}


func (val *DeploymentSettings) Defaults() *DeploymentSettings {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Cpu) {
		cpu_ := 1
		tmp.Cpu = &cpu_
	}
	if isZero(tmp.MemoryInGB) {
		memoryInGB_ := 1
		tmp.MemoryInGB = &memoryInGB_
	}
	if isZero(tmp.RuntimeVersion) {
		runtimeVersion_ := "Java_8"
		tmp.RuntimeVersion = &runtimeVersion_
	}
	return &tmp
}





type DeploymentSettingsInput interface {
	pulumi.Input

	ToDeploymentSettingsOutput() DeploymentSettingsOutput
	ToDeploymentSettingsOutputWithContext(context.Context) DeploymentSettingsOutput
}

type DeploymentSettingsArgs struct {
	Cpu                  pulumi.IntPtrInput    `pulumi:"cpu"`
	EnvironmentVariables pulumi.StringMapInput `pulumi:"environmentVariables"`
	JvmOptions           pulumi.StringPtrInput `pulumi:"jvmOptions"`
	MemoryInGB           pulumi.IntPtrInput    `pulumi:"memoryInGB"`
	NetCoreMainEntryPath pulumi.StringPtrInput `pulumi:"netCoreMainEntryPath"`
	RuntimeVersion       pulumi.StringPtrInput `pulumi:"runtimeVersion"`
}


func (val *DeploymentSettingsArgs) Defaults() *DeploymentSettingsArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Cpu) {
		tmp.Cpu = pulumi.IntPtr(1)
	}
	if isZero(tmp.MemoryInGB) {
		tmp.MemoryInGB = pulumi.IntPtr(1)
	}
	if isZero(tmp.RuntimeVersion) {
		tmp.RuntimeVersion = pulumi.StringPtr("Java_8")
	}
	return &tmp
}
func (DeploymentSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentSettings)(nil)).Elem()
}

func (i DeploymentSettingsArgs) ToDeploymentSettingsOutput() DeploymentSettingsOutput {
	return i.ToDeploymentSettingsOutputWithContext(context.Background())
}

func (i DeploymentSettingsArgs) ToDeploymentSettingsOutputWithContext(ctx context.Context) DeploymentSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentSettingsOutput)
}

func (i DeploymentSettingsArgs) ToDeploymentSettingsPtrOutput() DeploymentSettingsPtrOutput {
	return i.ToDeploymentSettingsPtrOutputWithContext(context.Background())
}

func (i DeploymentSettingsArgs) ToDeploymentSettingsPtrOutputWithContext(ctx context.Context) DeploymentSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentSettingsOutput).ToDeploymentSettingsPtrOutputWithContext(ctx)
}









type DeploymentSettingsPtrInput interface {
	pulumi.Input

	ToDeploymentSettingsPtrOutput() DeploymentSettingsPtrOutput
	ToDeploymentSettingsPtrOutputWithContext(context.Context) DeploymentSettingsPtrOutput
}

type deploymentSettingsPtrType DeploymentSettingsArgs

func DeploymentSettingsPtr(v *DeploymentSettingsArgs) DeploymentSettingsPtrInput {
	return (*deploymentSettingsPtrType)(v)
}

func (*deploymentSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentSettings)(nil)).Elem()
}

func (i *deploymentSettingsPtrType) ToDeploymentSettingsPtrOutput() DeploymentSettingsPtrOutput {
	return i.ToDeploymentSettingsPtrOutputWithContext(context.Background())
}

func (i *deploymentSettingsPtrType) ToDeploymentSettingsPtrOutputWithContext(ctx context.Context) DeploymentSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentSettingsPtrOutput)
}

type DeploymentSettingsOutput struct{ *pulumi.OutputState }

func (DeploymentSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentSettings)(nil)).Elem()
}

func (o DeploymentSettingsOutput) ToDeploymentSettingsOutput() DeploymentSettingsOutput {
	return o
}

func (o DeploymentSettingsOutput) ToDeploymentSettingsOutputWithContext(ctx context.Context) DeploymentSettingsOutput {
	return o
}

func (o DeploymentSettingsOutput) ToDeploymentSettingsPtrOutput() DeploymentSettingsPtrOutput {
	return o.ToDeploymentSettingsPtrOutputWithContext(context.Background())
}

func (o DeploymentSettingsOutput) ToDeploymentSettingsPtrOutputWithContext(ctx context.Context) DeploymentSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentSettings) *DeploymentSettings {
		return &v
	}).(DeploymentSettingsPtrOutput)
}

func (o DeploymentSettingsOutput) Cpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeploymentSettings) *int { return v.Cpu }).(pulumi.IntPtrOutput)
}

func (o DeploymentSettingsOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v DeploymentSettings) map[string]string { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

func (o DeploymentSettingsOutput) JvmOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentSettings) *string { return v.JvmOptions }).(pulumi.StringPtrOutput)
}

func (o DeploymentSettingsOutput) MemoryInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeploymentSettings) *int { return v.MemoryInGB }).(pulumi.IntPtrOutput)
}

func (o DeploymentSettingsOutput) NetCoreMainEntryPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentSettings) *string { return v.NetCoreMainEntryPath }).(pulumi.StringPtrOutput)
}

func (o DeploymentSettingsOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentSettings) *string { return v.RuntimeVersion }).(pulumi.StringPtrOutput)
}

type DeploymentSettingsPtrOutput struct{ *pulumi.OutputState }

func (DeploymentSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentSettings)(nil)).Elem()
}

func (o DeploymentSettingsPtrOutput) ToDeploymentSettingsPtrOutput() DeploymentSettingsPtrOutput {
	return o
}

func (o DeploymentSettingsPtrOutput) ToDeploymentSettingsPtrOutputWithContext(ctx context.Context) DeploymentSettingsPtrOutput {
	return o
}

func (o DeploymentSettingsPtrOutput) Elem() DeploymentSettingsOutput {
	return o.ApplyT(func(v *DeploymentSettings) DeploymentSettings {
		if v != nil {
			return *v
		}
		var ret DeploymentSettings
		return ret
	}).(DeploymentSettingsOutput)
}

func (o DeploymentSettingsPtrOutput) Cpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DeploymentSettings) *int {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.IntPtrOutput)
}

func (o DeploymentSettingsPtrOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeploymentSettings) map[string]string {
		if v == nil {
			return nil
		}
		return v.EnvironmentVariables
	}).(pulumi.StringMapOutput)
}

func (o DeploymentSettingsPtrOutput) JvmOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentSettings) *string {
		if v == nil {
			return nil
		}
		return v.JvmOptions
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentSettingsPtrOutput) MemoryInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DeploymentSettings) *int {
		if v == nil {
			return nil
		}
		return v.MemoryInGB
	}).(pulumi.IntPtrOutput)
}

func (o DeploymentSettingsPtrOutput) NetCoreMainEntryPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentSettings) *string {
		if v == nil {
			return nil
		}
		return v.NetCoreMainEntryPath
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentSettingsPtrOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentSettings) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeVersion
	}).(pulumi.StringPtrOutput)
}

type DeploymentSettingsResponse struct {
	Cpu                  *int              `pulumi:"cpu"`
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	JvmOptions           *string           `pulumi:"jvmOptions"`
	MemoryInGB           *int              `pulumi:"memoryInGB"`
	NetCoreMainEntryPath *string           `pulumi:"netCoreMainEntryPath"`
	RuntimeVersion       *string           `pulumi:"runtimeVersion"`
}


func (val *DeploymentSettingsResponse) Defaults() *DeploymentSettingsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Cpu) {
		cpu_ := 1
		tmp.Cpu = &cpu_
	}
	if isZero(tmp.MemoryInGB) {
		memoryInGB_ := 1
		tmp.MemoryInGB = &memoryInGB_
	}
	if isZero(tmp.RuntimeVersion) {
		runtimeVersion_ := "Java_8"
		tmp.RuntimeVersion = &runtimeVersion_
	}
	return &tmp
}

type DeploymentSettingsResponseOutput struct{ *pulumi.OutputState }

func (DeploymentSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentSettingsResponse)(nil)).Elem()
}

func (o DeploymentSettingsResponseOutput) ToDeploymentSettingsResponseOutput() DeploymentSettingsResponseOutput {
	return o
}

func (o DeploymentSettingsResponseOutput) ToDeploymentSettingsResponseOutputWithContext(ctx context.Context) DeploymentSettingsResponseOutput {
	return o
}

func (o DeploymentSettingsResponseOutput) Cpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeploymentSettingsResponse) *int { return v.Cpu }).(pulumi.IntPtrOutput)
}

func (o DeploymentSettingsResponseOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v DeploymentSettingsResponse) map[string]string { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

func (o DeploymentSettingsResponseOutput) JvmOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentSettingsResponse) *string { return v.JvmOptions }).(pulumi.StringPtrOutput)
}

func (o DeploymentSettingsResponseOutput) MemoryInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeploymentSettingsResponse) *int { return v.MemoryInGB }).(pulumi.IntPtrOutput)
}

func (o DeploymentSettingsResponseOutput) NetCoreMainEntryPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentSettingsResponse) *string { return v.NetCoreMainEntryPath }).(pulumi.StringPtrOutput)
}

func (o DeploymentSettingsResponseOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentSettingsResponse) *string { return v.RuntimeVersion }).(pulumi.StringPtrOutput)
}

type DeploymentSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (DeploymentSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentSettingsResponse)(nil)).Elem()
}

func (o DeploymentSettingsResponsePtrOutput) ToDeploymentSettingsResponsePtrOutput() DeploymentSettingsResponsePtrOutput {
	return o
}

func (o DeploymentSettingsResponsePtrOutput) ToDeploymentSettingsResponsePtrOutputWithContext(ctx context.Context) DeploymentSettingsResponsePtrOutput {
	return o
}

func (o DeploymentSettingsResponsePtrOutput) Elem() DeploymentSettingsResponseOutput {
	return o.ApplyT(func(v *DeploymentSettingsResponse) DeploymentSettingsResponse {
		if v != nil {
			return *v
		}
		var ret DeploymentSettingsResponse
		return ret
	}).(DeploymentSettingsResponseOutput)
}

func (o DeploymentSettingsResponsePtrOutput) Cpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DeploymentSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.IntPtrOutput)
}

func (o DeploymentSettingsResponsePtrOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeploymentSettingsResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.EnvironmentVariables
	}).(pulumi.StringMapOutput)
}

func (o DeploymentSettingsResponsePtrOutput) JvmOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.JvmOptions
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentSettingsResponsePtrOutput) MemoryInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DeploymentSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.MemoryInGB
	}).(pulumi.IntPtrOutput)
}

func (o DeploymentSettingsResponsePtrOutput) NetCoreMainEntryPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.NetCoreMainEntryPath
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentSettingsResponsePtrOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeVersion
	}).(pulumi.StringPtrOutput)
}

type GatewayApiMetadataProperties struct {
	Description   *string `pulumi:"description"`
	Documentation *string `pulumi:"documentation"`
	ServerUrl     *string `pulumi:"serverUrl"`
	Title         *string `pulumi:"title"`
	Version       *string `pulumi:"version"`
}





type GatewayApiMetadataPropertiesInput interface {
	pulumi.Input

	ToGatewayApiMetadataPropertiesOutput() GatewayApiMetadataPropertiesOutput
	ToGatewayApiMetadataPropertiesOutputWithContext(context.Context) GatewayApiMetadataPropertiesOutput
}

type GatewayApiMetadataPropertiesArgs struct {
	Description   pulumi.StringPtrInput `pulumi:"description"`
	Documentation pulumi.StringPtrInput `pulumi:"documentation"`
	ServerUrl     pulumi.StringPtrInput `pulumi:"serverUrl"`
	Title         pulumi.StringPtrInput `pulumi:"title"`
	Version       pulumi.StringPtrInput `pulumi:"version"`
}

func (GatewayApiMetadataPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayApiMetadataProperties)(nil)).Elem()
}

func (i GatewayApiMetadataPropertiesArgs) ToGatewayApiMetadataPropertiesOutput() GatewayApiMetadataPropertiesOutput {
	return i.ToGatewayApiMetadataPropertiesOutputWithContext(context.Background())
}

func (i GatewayApiMetadataPropertiesArgs) ToGatewayApiMetadataPropertiesOutputWithContext(ctx context.Context) GatewayApiMetadataPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayApiMetadataPropertiesOutput)
}

func (i GatewayApiMetadataPropertiesArgs) ToGatewayApiMetadataPropertiesPtrOutput() GatewayApiMetadataPropertiesPtrOutput {
	return i.ToGatewayApiMetadataPropertiesPtrOutputWithContext(context.Background())
}

func (i GatewayApiMetadataPropertiesArgs) ToGatewayApiMetadataPropertiesPtrOutputWithContext(ctx context.Context) GatewayApiMetadataPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayApiMetadataPropertiesOutput).ToGatewayApiMetadataPropertiesPtrOutputWithContext(ctx)
}









type GatewayApiMetadataPropertiesPtrInput interface {
	pulumi.Input

	ToGatewayApiMetadataPropertiesPtrOutput() GatewayApiMetadataPropertiesPtrOutput
	ToGatewayApiMetadataPropertiesPtrOutputWithContext(context.Context) GatewayApiMetadataPropertiesPtrOutput
}

type gatewayApiMetadataPropertiesPtrType GatewayApiMetadataPropertiesArgs

func GatewayApiMetadataPropertiesPtr(v *GatewayApiMetadataPropertiesArgs) GatewayApiMetadataPropertiesPtrInput {
	return (*gatewayApiMetadataPropertiesPtrType)(v)
}

func (*gatewayApiMetadataPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayApiMetadataProperties)(nil)).Elem()
}

func (i *gatewayApiMetadataPropertiesPtrType) ToGatewayApiMetadataPropertiesPtrOutput() GatewayApiMetadataPropertiesPtrOutput {
	return i.ToGatewayApiMetadataPropertiesPtrOutputWithContext(context.Background())
}

func (i *gatewayApiMetadataPropertiesPtrType) ToGatewayApiMetadataPropertiesPtrOutputWithContext(ctx context.Context) GatewayApiMetadataPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayApiMetadataPropertiesPtrOutput)
}

type GatewayApiMetadataPropertiesOutput struct{ *pulumi.OutputState }

func (GatewayApiMetadataPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayApiMetadataProperties)(nil)).Elem()
}

func (o GatewayApiMetadataPropertiesOutput) ToGatewayApiMetadataPropertiesOutput() GatewayApiMetadataPropertiesOutput {
	return o
}

func (o GatewayApiMetadataPropertiesOutput) ToGatewayApiMetadataPropertiesOutputWithContext(ctx context.Context) GatewayApiMetadataPropertiesOutput {
	return o
}

func (o GatewayApiMetadataPropertiesOutput) ToGatewayApiMetadataPropertiesPtrOutput() GatewayApiMetadataPropertiesPtrOutput {
	return o.ToGatewayApiMetadataPropertiesPtrOutputWithContext(context.Background())
}

func (o GatewayApiMetadataPropertiesOutput) ToGatewayApiMetadataPropertiesPtrOutputWithContext(ctx context.Context) GatewayApiMetadataPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GatewayApiMetadataProperties) *GatewayApiMetadataProperties {
		return &v
	}).(GatewayApiMetadataPropertiesPtrOutput)
}

func (o GatewayApiMetadataPropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayApiMetadataProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GatewayApiMetadataPropertiesOutput) Documentation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayApiMetadataProperties) *string { return v.Documentation }).(pulumi.StringPtrOutput)
}

func (o GatewayApiMetadataPropertiesOutput) ServerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayApiMetadataProperties) *string { return v.ServerUrl }).(pulumi.StringPtrOutput)
}

func (o GatewayApiMetadataPropertiesOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayApiMetadataProperties) *string { return v.Title }).(pulumi.StringPtrOutput)
}

func (o GatewayApiMetadataPropertiesOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayApiMetadataProperties) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type GatewayApiMetadataPropertiesPtrOutput struct{ *pulumi.OutputState }

func (GatewayApiMetadataPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayApiMetadataProperties)(nil)).Elem()
}

func (o GatewayApiMetadataPropertiesPtrOutput) ToGatewayApiMetadataPropertiesPtrOutput() GatewayApiMetadataPropertiesPtrOutput {
	return o
}

func (o GatewayApiMetadataPropertiesPtrOutput) ToGatewayApiMetadataPropertiesPtrOutputWithContext(ctx context.Context) GatewayApiMetadataPropertiesPtrOutput {
	return o
}

func (o GatewayApiMetadataPropertiesPtrOutput) Elem() GatewayApiMetadataPropertiesOutput {
	return o.ApplyT(func(v *GatewayApiMetadataProperties) GatewayApiMetadataProperties {
		if v != nil {
			return *v
		}
		var ret GatewayApiMetadataProperties
		return ret
	}).(GatewayApiMetadataPropertiesOutput)
}

func (o GatewayApiMetadataPropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o GatewayApiMetadataPropertiesPtrOutput) Documentation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return v.Documentation
	}).(pulumi.StringPtrOutput)
}

func (o GatewayApiMetadataPropertiesPtrOutput) ServerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return v.ServerUrl
	}).(pulumi.StringPtrOutput)
}

func (o GatewayApiMetadataPropertiesPtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return v.Title
	}).(pulumi.StringPtrOutput)
}

func (o GatewayApiMetadataPropertiesPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type GatewayApiMetadataPropertiesResponse struct {
	Description   *string `pulumi:"description"`
	Documentation *string `pulumi:"documentation"`
	ServerUrl     *string `pulumi:"serverUrl"`
	Title         *string `pulumi:"title"`
	Version       *string `pulumi:"version"`
}

type GatewayApiMetadataPropertiesResponseOutput struct{ *pulumi.OutputState }

func (GatewayApiMetadataPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayApiMetadataPropertiesResponse)(nil)).Elem()
}

func (o GatewayApiMetadataPropertiesResponseOutput) ToGatewayApiMetadataPropertiesResponseOutput() GatewayApiMetadataPropertiesResponseOutput {
	return o
}

func (o GatewayApiMetadataPropertiesResponseOutput) ToGatewayApiMetadataPropertiesResponseOutputWithContext(ctx context.Context) GatewayApiMetadataPropertiesResponseOutput {
	return o
}

func (o GatewayApiMetadataPropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayApiMetadataPropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GatewayApiMetadataPropertiesResponseOutput) Documentation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayApiMetadataPropertiesResponse) *string { return v.Documentation }).(pulumi.StringPtrOutput)
}

func (o GatewayApiMetadataPropertiesResponseOutput) ServerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayApiMetadataPropertiesResponse) *string { return v.ServerUrl }).(pulumi.StringPtrOutput)
}

func (o GatewayApiMetadataPropertiesResponseOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayApiMetadataPropertiesResponse) *string { return v.Title }).(pulumi.StringPtrOutput)
}

func (o GatewayApiMetadataPropertiesResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayApiMetadataPropertiesResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type GatewayApiMetadataPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (GatewayApiMetadataPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayApiMetadataPropertiesResponse)(nil)).Elem()
}

func (o GatewayApiMetadataPropertiesResponsePtrOutput) ToGatewayApiMetadataPropertiesResponsePtrOutput() GatewayApiMetadataPropertiesResponsePtrOutput {
	return o
}

func (o GatewayApiMetadataPropertiesResponsePtrOutput) ToGatewayApiMetadataPropertiesResponsePtrOutputWithContext(ctx context.Context) GatewayApiMetadataPropertiesResponsePtrOutput {
	return o
}

func (o GatewayApiMetadataPropertiesResponsePtrOutput) Elem() GatewayApiMetadataPropertiesResponseOutput {
	return o.ApplyT(func(v *GatewayApiMetadataPropertiesResponse) GatewayApiMetadataPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret GatewayApiMetadataPropertiesResponse
		return ret
	}).(GatewayApiMetadataPropertiesResponseOutput)
}

func (o GatewayApiMetadataPropertiesResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o GatewayApiMetadataPropertiesResponsePtrOutput) Documentation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Documentation
	}).(pulumi.StringPtrOutput)
}

func (o GatewayApiMetadataPropertiesResponsePtrOutput) ServerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServerUrl
	}).(pulumi.StringPtrOutput)
}

func (o GatewayApiMetadataPropertiesResponsePtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Title
	}).(pulumi.StringPtrOutput)
}

func (o GatewayApiMetadataPropertiesResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type GatewayApiRoute struct {
	Description *string  `pulumi:"description"`
	Filters     []string `pulumi:"filters"`
	Order       *int     `pulumi:"order"`
	Predicates  []string `pulumi:"predicates"`
	SsoEnabled  *bool    `pulumi:"ssoEnabled"`
	Tags        []string `pulumi:"tags"`
	Title       *string  `pulumi:"title"`
	TokenRelay  *bool    `pulumi:"tokenRelay"`
	Uri         *string  `pulumi:"uri"`
}





type GatewayApiRouteInput interface {
	pulumi.Input

	ToGatewayApiRouteOutput() GatewayApiRouteOutput
	ToGatewayApiRouteOutputWithContext(context.Context) GatewayApiRouteOutput
}

type GatewayApiRouteArgs struct {
	Description pulumi.StringPtrInput   `pulumi:"description"`
	Filters     pulumi.StringArrayInput `pulumi:"filters"`
	Order       pulumi.IntPtrInput      `pulumi:"order"`
	Predicates  pulumi.StringArrayInput `pulumi:"predicates"`
	SsoEnabled  pulumi.BoolPtrInput     `pulumi:"ssoEnabled"`
	Tags        pulumi.StringArrayInput `pulumi:"tags"`
	Title       pulumi.StringPtrInput   `pulumi:"title"`
	TokenRelay  pulumi.BoolPtrInput     `pulumi:"tokenRelay"`
	Uri         pulumi.StringPtrInput   `pulumi:"uri"`
}

func (GatewayApiRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayApiRoute)(nil)).Elem()
}

func (i GatewayApiRouteArgs) ToGatewayApiRouteOutput() GatewayApiRouteOutput {
	return i.ToGatewayApiRouteOutputWithContext(context.Background())
}

func (i GatewayApiRouteArgs) ToGatewayApiRouteOutputWithContext(ctx context.Context) GatewayApiRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayApiRouteOutput)
}





type GatewayApiRouteArrayInput interface {
	pulumi.Input

	ToGatewayApiRouteArrayOutput() GatewayApiRouteArrayOutput
	ToGatewayApiRouteArrayOutputWithContext(context.Context) GatewayApiRouteArrayOutput
}

type GatewayApiRouteArray []GatewayApiRouteInput

func (GatewayApiRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GatewayApiRoute)(nil)).Elem()
}

func (i GatewayApiRouteArray) ToGatewayApiRouteArrayOutput() GatewayApiRouteArrayOutput {
	return i.ToGatewayApiRouteArrayOutputWithContext(context.Background())
}

func (i GatewayApiRouteArray) ToGatewayApiRouteArrayOutputWithContext(ctx context.Context) GatewayApiRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayApiRouteArrayOutput)
}

type GatewayApiRouteOutput struct{ *pulumi.OutputState }

func (GatewayApiRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayApiRoute)(nil)).Elem()
}

func (o GatewayApiRouteOutput) ToGatewayApiRouteOutput() GatewayApiRouteOutput {
	return o
}

func (o GatewayApiRouteOutput) ToGatewayApiRouteOutputWithContext(ctx context.Context) GatewayApiRouteOutput {
	return o
}

func (o GatewayApiRouteOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayApiRoute) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GatewayApiRouteOutput) Filters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GatewayApiRoute) []string { return v.Filters }).(pulumi.StringArrayOutput)
}

func (o GatewayApiRouteOutput) Order() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GatewayApiRoute) *int { return v.Order }).(pulumi.IntPtrOutput)
}

func (o GatewayApiRouteOutput) Predicates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GatewayApiRoute) []string { return v.Predicates }).(pulumi.StringArrayOutput)
}

func (o GatewayApiRouteOutput) SsoEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GatewayApiRoute) *bool { return v.SsoEnabled }).(pulumi.BoolPtrOutput)
}

func (o GatewayApiRouteOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GatewayApiRoute) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o GatewayApiRouteOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayApiRoute) *string { return v.Title }).(pulumi.StringPtrOutput)
}

func (o GatewayApiRouteOutput) TokenRelay() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GatewayApiRoute) *bool { return v.TokenRelay }).(pulumi.BoolPtrOutput)
}

func (o GatewayApiRouteOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayApiRoute) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type GatewayApiRouteArrayOutput struct{ *pulumi.OutputState }

func (GatewayApiRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GatewayApiRoute)(nil)).Elem()
}

func (o GatewayApiRouteArrayOutput) ToGatewayApiRouteArrayOutput() GatewayApiRouteArrayOutput {
	return o
}

func (o GatewayApiRouteArrayOutput) ToGatewayApiRouteArrayOutputWithContext(ctx context.Context) GatewayApiRouteArrayOutput {
	return o
}

func (o GatewayApiRouteArrayOutput) Index(i pulumi.IntInput) GatewayApiRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GatewayApiRoute {
		return vs[0].([]GatewayApiRoute)[vs[1].(int)]
	}).(GatewayApiRouteOutput)
}

type GatewayApiRouteResponse struct {
	Description *string  `pulumi:"description"`
	Filters     []string `pulumi:"filters"`
	Order       *int     `pulumi:"order"`
	Predicates  []string `pulumi:"predicates"`
	SsoEnabled  *bool    `pulumi:"ssoEnabled"`
	Tags        []string `pulumi:"tags"`
	Title       *string  `pulumi:"title"`
	TokenRelay  *bool    `pulumi:"tokenRelay"`
	Uri         *string  `pulumi:"uri"`
}

type GatewayApiRouteResponseOutput struct{ *pulumi.OutputState }

func (GatewayApiRouteResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayApiRouteResponse)(nil)).Elem()
}

func (o GatewayApiRouteResponseOutput) ToGatewayApiRouteResponseOutput() GatewayApiRouteResponseOutput {
	return o
}

func (o GatewayApiRouteResponseOutput) ToGatewayApiRouteResponseOutputWithContext(ctx context.Context) GatewayApiRouteResponseOutput {
	return o
}

func (o GatewayApiRouteResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayApiRouteResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GatewayApiRouteResponseOutput) Filters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GatewayApiRouteResponse) []string { return v.Filters }).(pulumi.StringArrayOutput)
}

func (o GatewayApiRouteResponseOutput) Order() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GatewayApiRouteResponse) *int { return v.Order }).(pulumi.IntPtrOutput)
}

func (o GatewayApiRouteResponseOutput) Predicates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GatewayApiRouteResponse) []string { return v.Predicates }).(pulumi.StringArrayOutput)
}

func (o GatewayApiRouteResponseOutput) SsoEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GatewayApiRouteResponse) *bool { return v.SsoEnabled }).(pulumi.BoolPtrOutput)
}

func (o GatewayApiRouteResponseOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GatewayApiRouteResponse) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o GatewayApiRouteResponseOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayApiRouteResponse) *string { return v.Title }).(pulumi.StringPtrOutput)
}

func (o GatewayApiRouteResponseOutput) TokenRelay() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GatewayApiRouteResponse) *bool { return v.TokenRelay }).(pulumi.BoolPtrOutput)
}

func (o GatewayApiRouteResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayApiRouteResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type GatewayApiRouteResponseArrayOutput struct{ *pulumi.OutputState }

func (GatewayApiRouteResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GatewayApiRouteResponse)(nil)).Elem()
}

func (o GatewayApiRouteResponseArrayOutput) ToGatewayApiRouteResponseArrayOutput() GatewayApiRouteResponseArrayOutput {
	return o
}

func (o GatewayApiRouteResponseArrayOutput) ToGatewayApiRouteResponseArrayOutputWithContext(ctx context.Context) GatewayApiRouteResponseArrayOutput {
	return o
}

func (o GatewayApiRouteResponseArrayOutput) Index(i pulumi.IntInput) GatewayApiRouteResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GatewayApiRouteResponse {
		return vs[0].([]GatewayApiRouteResponse)[vs[1].(int)]
	}).(GatewayApiRouteResponseOutput)
}

type GatewayCorsProperties struct {
	AllowCredentials *bool    `pulumi:"allowCredentials"`
	AllowedHeaders   []string `pulumi:"allowedHeaders"`
	AllowedMethods   []string `pulumi:"allowedMethods"`
	AllowedOrigins   []string `pulumi:"allowedOrigins"`
	ExposedHeaders   []string `pulumi:"exposedHeaders"`
	MaxAge           *int     `pulumi:"maxAge"`
}





type GatewayCorsPropertiesInput interface {
	pulumi.Input

	ToGatewayCorsPropertiesOutput() GatewayCorsPropertiesOutput
	ToGatewayCorsPropertiesOutputWithContext(context.Context) GatewayCorsPropertiesOutput
}

type GatewayCorsPropertiesArgs struct {
	AllowCredentials pulumi.BoolPtrInput     `pulumi:"allowCredentials"`
	AllowedHeaders   pulumi.StringArrayInput `pulumi:"allowedHeaders"`
	AllowedMethods   pulumi.StringArrayInput `pulumi:"allowedMethods"`
	AllowedOrigins   pulumi.StringArrayInput `pulumi:"allowedOrigins"`
	ExposedHeaders   pulumi.StringArrayInput `pulumi:"exposedHeaders"`
	MaxAge           pulumi.IntPtrInput      `pulumi:"maxAge"`
}

func (GatewayCorsPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayCorsProperties)(nil)).Elem()
}

func (i GatewayCorsPropertiesArgs) ToGatewayCorsPropertiesOutput() GatewayCorsPropertiesOutput {
	return i.ToGatewayCorsPropertiesOutputWithContext(context.Background())
}

func (i GatewayCorsPropertiesArgs) ToGatewayCorsPropertiesOutputWithContext(ctx context.Context) GatewayCorsPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayCorsPropertiesOutput)
}

func (i GatewayCorsPropertiesArgs) ToGatewayCorsPropertiesPtrOutput() GatewayCorsPropertiesPtrOutput {
	return i.ToGatewayCorsPropertiesPtrOutputWithContext(context.Background())
}

func (i GatewayCorsPropertiesArgs) ToGatewayCorsPropertiesPtrOutputWithContext(ctx context.Context) GatewayCorsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayCorsPropertiesOutput).ToGatewayCorsPropertiesPtrOutputWithContext(ctx)
}









type GatewayCorsPropertiesPtrInput interface {
	pulumi.Input

	ToGatewayCorsPropertiesPtrOutput() GatewayCorsPropertiesPtrOutput
	ToGatewayCorsPropertiesPtrOutputWithContext(context.Context) GatewayCorsPropertiesPtrOutput
}

type gatewayCorsPropertiesPtrType GatewayCorsPropertiesArgs

func GatewayCorsPropertiesPtr(v *GatewayCorsPropertiesArgs) GatewayCorsPropertiesPtrInput {
	return (*gatewayCorsPropertiesPtrType)(v)
}

func (*gatewayCorsPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayCorsProperties)(nil)).Elem()
}

func (i *gatewayCorsPropertiesPtrType) ToGatewayCorsPropertiesPtrOutput() GatewayCorsPropertiesPtrOutput {
	return i.ToGatewayCorsPropertiesPtrOutputWithContext(context.Background())
}

func (i *gatewayCorsPropertiesPtrType) ToGatewayCorsPropertiesPtrOutputWithContext(ctx context.Context) GatewayCorsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayCorsPropertiesPtrOutput)
}

type GatewayCorsPropertiesOutput struct{ *pulumi.OutputState }

func (GatewayCorsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayCorsProperties)(nil)).Elem()
}

func (o GatewayCorsPropertiesOutput) ToGatewayCorsPropertiesOutput() GatewayCorsPropertiesOutput {
	return o
}

func (o GatewayCorsPropertiesOutput) ToGatewayCorsPropertiesOutputWithContext(ctx context.Context) GatewayCorsPropertiesOutput {
	return o
}

func (o GatewayCorsPropertiesOutput) ToGatewayCorsPropertiesPtrOutput() GatewayCorsPropertiesPtrOutput {
	return o.ToGatewayCorsPropertiesPtrOutputWithContext(context.Background())
}

func (o GatewayCorsPropertiesOutput) ToGatewayCorsPropertiesPtrOutputWithContext(ctx context.Context) GatewayCorsPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GatewayCorsProperties) *GatewayCorsProperties {
		return &v
	}).(GatewayCorsPropertiesPtrOutput)
}

func (o GatewayCorsPropertiesOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GatewayCorsProperties) *bool { return v.AllowCredentials }).(pulumi.BoolPtrOutput)
}

func (o GatewayCorsPropertiesOutput) AllowedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GatewayCorsProperties) []string { return v.AllowedHeaders }).(pulumi.StringArrayOutput)
}

func (o GatewayCorsPropertiesOutput) AllowedMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GatewayCorsProperties) []string { return v.AllowedMethods }).(pulumi.StringArrayOutput)
}

func (o GatewayCorsPropertiesOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GatewayCorsProperties) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

func (o GatewayCorsPropertiesOutput) ExposedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GatewayCorsProperties) []string { return v.ExposedHeaders }).(pulumi.StringArrayOutput)
}

func (o GatewayCorsPropertiesOutput) MaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GatewayCorsProperties) *int { return v.MaxAge }).(pulumi.IntPtrOutput)
}

type GatewayCorsPropertiesPtrOutput struct{ *pulumi.OutputState }

func (GatewayCorsPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayCorsProperties)(nil)).Elem()
}

func (o GatewayCorsPropertiesPtrOutput) ToGatewayCorsPropertiesPtrOutput() GatewayCorsPropertiesPtrOutput {
	return o
}

func (o GatewayCorsPropertiesPtrOutput) ToGatewayCorsPropertiesPtrOutputWithContext(ctx context.Context) GatewayCorsPropertiesPtrOutput {
	return o
}

func (o GatewayCorsPropertiesPtrOutput) Elem() GatewayCorsPropertiesOutput {
	return o.ApplyT(func(v *GatewayCorsProperties) GatewayCorsProperties {
		if v != nil {
			return *v
		}
		var ret GatewayCorsProperties
		return ret
	}).(GatewayCorsPropertiesOutput)
}

func (o GatewayCorsPropertiesPtrOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GatewayCorsProperties) *bool {
		if v == nil {
			return nil
		}
		return v.AllowCredentials
	}).(pulumi.BoolPtrOutput)
}

func (o GatewayCorsPropertiesPtrOutput) AllowedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayCorsProperties) []string {
		if v == nil {
			return nil
		}
		return v.AllowedHeaders
	}).(pulumi.StringArrayOutput)
}

func (o GatewayCorsPropertiesPtrOutput) AllowedMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayCorsProperties) []string {
		if v == nil {
			return nil
		}
		return v.AllowedMethods
	}).(pulumi.StringArrayOutput)
}

func (o GatewayCorsPropertiesPtrOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayCorsProperties) []string {
		if v == nil {
			return nil
		}
		return v.AllowedOrigins
	}).(pulumi.StringArrayOutput)
}

func (o GatewayCorsPropertiesPtrOutput) ExposedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayCorsProperties) []string {
		if v == nil {
			return nil
		}
		return v.ExposedHeaders
	}).(pulumi.StringArrayOutput)
}

func (o GatewayCorsPropertiesPtrOutput) MaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GatewayCorsProperties) *int {
		if v == nil {
			return nil
		}
		return v.MaxAge
	}).(pulumi.IntPtrOutput)
}

type GatewayCorsPropertiesResponse struct {
	AllowCredentials *bool    `pulumi:"allowCredentials"`
	AllowedHeaders   []string `pulumi:"allowedHeaders"`
	AllowedMethods   []string `pulumi:"allowedMethods"`
	AllowedOrigins   []string `pulumi:"allowedOrigins"`
	ExposedHeaders   []string `pulumi:"exposedHeaders"`
	MaxAge           *int     `pulumi:"maxAge"`
}

type GatewayCorsPropertiesResponseOutput struct{ *pulumi.OutputState }

func (GatewayCorsPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayCorsPropertiesResponse)(nil)).Elem()
}

func (o GatewayCorsPropertiesResponseOutput) ToGatewayCorsPropertiesResponseOutput() GatewayCorsPropertiesResponseOutput {
	return o
}

func (o GatewayCorsPropertiesResponseOutput) ToGatewayCorsPropertiesResponseOutputWithContext(ctx context.Context) GatewayCorsPropertiesResponseOutput {
	return o
}

func (o GatewayCorsPropertiesResponseOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GatewayCorsPropertiesResponse) *bool { return v.AllowCredentials }).(pulumi.BoolPtrOutput)
}

func (o GatewayCorsPropertiesResponseOutput) AllowedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GatewayCorsPropertiesResponse) []string { return v.AllowedHeaders }).(pulumi.StringArrayOutput)
}

func (o GatewayCorsPropertiesResponseOutput) AllowedMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GatewayCorsPropertiesResponse) []string { return v.AllowedMethods }).(pulumi.StringArrayOutput)
}

func (o GatewayCorsPropertiesResponseOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GatewayCorsPropertiesResponse) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

func (o GatewayCorsPropertiesResponseOutput) ExposedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GatewayCorsPropertiesResponse) []string { return v.ExposedHeaders }).(pulumi.StringArrayOutput)
}

func (o GatewayCorsPropertiesResponseOutput) MaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GatewayCorsPropertiesResponse) *int { return v.MaxAge }).(pulumi.IntPtrOutput)
}

type GatewayCorsPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (GatewayCorsPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayCorsPropertiesResponse)(nil)).Elem()
}

func (o GatewayCorsPropertiesResponsePtrOutput) ToGatewayCorsPropertiesResponsePtrOutput() GatewayCorsPropertiesResponsePtrOutput {
	return o
}

func (o GatewayCorsPropertiesResponsePtrOutput) ToGatewayCorsPropertiesResponsePtrOutputWithContext(ctx context.Context) GatewayCorsPropertiesResponsePtrOutput {
	return o
}

func (o GatewayCorsPropertiesResponsePtrOutput) Elem() GatewayCorsPropertiesResponseOutput {
	return o.ApplyT(func(v *GatewayCorsPropertiesResponse) GatewayCorsPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret GatewayCorsPropertiesResponse
		return ret
	}).(GatewayCorsPropertiesResponseOutput)
}

func (o GatewayCorsPropertiesResponsePtrOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GatewayCorsPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AllowCredentials
	}).(pulumi.BoolPtrOutput)
}

func (o GatewayCorsPropertiesResponsePtrOutput) AllowedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayCorsPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedHeaders
	}).(pulumi.StringArrayOutput)
}

func (o GatewayCorsPropertiesResponsePtrOutput) AllowedMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayCorsPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedMethods
	}).(pulumi.StringArrayOutput)
}

func (o GatewayCorsPropertiesResponsePtrOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayCorsPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedOrigins
	}).(pulumi.StringArrayOutput)
}

func (o GatewayCorsPropertiesResponsePtrOutput) ExposedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayCorsPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.ExposedHeaders
	}).(pulumi.StringArrayOutput)
}

func (o GatewayCorsPropertiesResponsePtrOutput) MaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GatewayCorsPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxAge
	}).(pulumi.IntPtrOutput)
}

type GatewayCustomDomainProperties struct {
	Thumbprint *string `pulumi:"thumbprint"`
}





type GatewayCustomDomainPropertiesInput interface {
	pulumi.Input

	ToGatewayCustomDomainPropertiesOutput() GatewayCustomDomainPropertiesOutput
	ToGatewayCustomDomainPropertiesOutputWithContext(context.Context) GatewayCustomDomainPropertiesOutput
}

type GatewayCustomDomainPropertiesArgs struct {
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (GatewayCustomDomainPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayCustomDomainProperties)(nil)).Elem()
}

func (i GatewayCustomDomainPropertiesArgs) ToGatewayCustomDomainPropertiesOutput() GatewayCustomDomainPropertiesOutput {
	return i.ToGatewayCustomDomainPropertiesOutputWithContext(context.Background())
}

func (i GatewayCustomDomainPropertiesArgs) ToGatewayCustomDomainPropertiesOutputWithContext(ctx context.Context) GatewayCustomDomainPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayCustomDomainPropertiesOutput)
}

func (i GatewayCustomDomainPropertiesArgs) ToGatewayCustomDomainPropertiesPtrOutput() GatewayCustomDomainPropertiesPtrOutput {
	return i.ToGatewayCustomDomainPropertiesPtrOutputWithContext(context.Background())
}

func (i GatewayCustomDomainPropertiesArgs) ToGatewayCustomDomainPropertiesPtrOutputWithContext(ctx context.Context) GatewayCustomDomainPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayCustomDomainPropertiesOutput).ToGatewayCustomDomainPropertiesPtrOutputWithContext(ctx)
}









type GatewayCustomDomainPropertiesPtrInput interface {
	pulumi.Input

	ToGatewayCustomDomainPropertiesPtrOutput() GatewayCustomDomainPropertiesPtrOutput
	ToGatewayCustomDomainPropertiesPtrOutputWithContext(context.Context) GatewayCustomDomainPropertiesPtrOutput
}

type gatewayCustomDomainPropertiesPtrType GatewayCustomDomainPropertiesArgs

func GatewayCustomDomainPropertiesPtr(v *GatewayCustomDomainPropertiesArgs) GatewayCustomDomainPropertiesPtrInput {
	return (*gatewayCustomDomainPropertiesPtrType)(v)
}

func (*gatewayCustomDomainPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayCustomDomainProperties)(nil)).Elem()
}

func (i *gatewayCustomDomainPropertiesPtrType) ToGatewayCustomDomainPropertiesPtrOutput() GatewayCustomDomainPropertiesPtrOutput {
	return i.ToGatewayCustomDomainPropertiesPtrOutputWithContext(context.Background())
}

func (i *gatewayCustomDomainPropertiesPtrType) ToGatewayCustomDomainPropertiesPtrOutputWithContext(ctx context.Context) GatewayCustomDomainPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayCustomDomainPropertiesPtrOutput)
}

type GatewayCustomDomainPropertiesOutput struct{ *pulumi.OutputState }

func (GatewayCustomDomainPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayCustomDomainProperties)(nil)).Elem()
}

func (o GatewayCustomDomainPropertiesOutput) ToGatewayCustomDomainPropertiesOutput() GatewayCustomDomainPropertiesOutput {
	return o
}

func (o GatewayCustomDomainPropertiesOutput) ToGatewayCustomDomainPropertiesOutputWithContext(ctx context.Context) GatewayCustomDomainPropertiesOutput {
	return o
}

func (o GatewayCustomDomainPropertiesOutput) ToGatewayCustomDomainPropertiesPtrOutput() GatewayCustomDomainPropertiesPtrOutput {
	return o.ToGatewayCustomDomainPropertiesPtrOutputWithContext(context.Background())
}

func (o GatewayCustomDomainPropertiesOutput) ToGatewayCustomDomainPropertiesPtrOutputWithContext(ctx context.Context) GatewayCustomDomainPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GatewayCustomDomainProperties) *GatewayCustomDomainProperties {
		return &v
	}).(GatewayCustomDomainPropertiesPtrOutput)
}

func (o GatewayCustomDomainPropertiesOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayCustomDomainProperties) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type GatewayCustomDomainPropertiesPtrOutput struct{ *pulumi.OutputState }

func (GatewayCustomDomainPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayCustomDomainProperties)(nil)).Elem()
}

func (o GatewayCustomDomainPropertiesPtrOutput) ToGatewayCustomDomainPropertiesPtrOutput() GatewayCustomDomainPropertiesPtrOutput {
	return o
}

func (o GatewayCustomDomainPropertiesPtrOutput) ToGatewayCustomDomainPropertiesPtrOutputWithContext(ctx context.Context) GatewayCustomDomainPropertiesPtrOutput {
	return o
}

func (o GatewayCustomDomainPropertiesPtrOutput) Elem() GatewayCustomDomainPropertiesOutput {
	return o.ApplyT(func(v *GatewayCustomDomainProperties) GatewayCustomDomainProperties {
		if v != nil {
			return *v
		}
		var ret GatewayCustomDomainProperties
		return ret
	}).(GatewayCustomDomainPropertiesOutput)
}

func (o GatewayCustomDomainPropertiesPtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayCustomDomainProperties) *string {
		if v == nil {
			return nil
		}
		return v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

type GatewayCustomDomainPropertiesResponse struct {
	Thumbprint *string `pulumi:"thumbprint"`
}

type GatewayCustomDomainPropertiesResponseOutput struct{ *pulumi.OutputState }

func (GatewayCustomDomainPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayCustomDomainPropertiesResponse)(nil)).Elem()
}

func (o GatewayCustomDomainPropertiesResponseOutput) ToGatewayCustomDomainPropertiesResponseOutput() GatewayCustomDomainPropertiesResponseOutput {
	return o
}

func (o GatewayCustomDomainPropertiesResponseOutput) ToGatewayCustomDomainPropertiesResponseOutputWithContext(ctx context.Context) GatewayCustomDomainPropertiesResponseOutput {
	return o
}

func (o GatewayCustomDomainPropertiesResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayCustomDomainPropertiesResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type GatewayInstanceResponse struct {
	Name   string `pulumi:"name"`
	Status string `pulumi:"status"`
}

type GatewayInstanceResponseOutput struct{ *pulumi.OutputState }

func (GatewayInstanceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayInstanceResponse)(nil)).Elem()
}

func (o GatewayInstanceResponseOutput) ToGatewayInstanceResponseOutput() GatewayInstanceResponseOutput {
	return o
}

func (o GatewayInstanceResponseOutput) ToGatewayInstanceResponseOutputWithContext(ctx context.Context) GatewayInstanceResponseOutput {
	return o
}

func (o GatewayInstanceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayInstanceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o GatewayInstanceResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayInstanceResponse) string { return v.Status }).(pulumi.StringOutput)
}

type GatewayInstanceResponseArrayOutput struct{ *pulumi.OutputState }

func (GatewayInstanceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GatewayInstanceResponse)(nil)).Elem()
}

func (o GatewayInstanceResponseArrayOutput) ToGatewayInstanceResponseArrayOutput() GatewayInstanceResponseArrayOutput {
	return o
}

func (o GatewayInstanceResponseArrayOutput) ToGatewayInstanceResponseArrayOutputWithContext(ctx context.Context) GatewayInstanceResponseArrayOutput {
	return o
}

func (o GatewayInstanceResponseArrayOutput) Index(i pulumi.IntInput) GatewayInstanceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GatewayInstanceResponse {
		return vs[0].([]GatewayInstanceResponse)[vs[1].(int)]
	}).(GatewayInstanceResponseOutput)
}

type GatewayOperatorPropertiesResponse struct {
	Instances        []GatewayInstanceResponse               `pulumi:"instances"`
	ResourceRequests GatewayOperatorResourceRequestsResponse `pulumi:"resourceRequests"`
}

type GatewayOperatorPropertiesResponseOutput struct{ *pulumi.OutputState }

func (GatewayOperatorPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayOperatorPropertiesResponse)(nil)).Elem()
}

func (o GatewayOperatorPropertiesResponseOutput) ToGatewayOperatorPropertiesResponseOutput() GatewayOperatorPropertiesResponseOutput {
	return o
}

func (o GatewayOperatorPropertiesResponseOutput) ToGatewayOperatorPropertiesResponseOutputWithContext(ctx context.Context) GatewayOperatorPropertiesResponseOutput {
	return o
}

func (o GatewayOperatorPropertiesResponseOutput) Instances() GatewayInstanceResponseArrayOutput {
	return o.ApplyT(func(v GatewayOperatorPropertiesResponse) []GatewayInstanceResponse { return v.Instances }).(GatewayInstanceResponseArrayOutput)
}

func (o GatewayOperatorPropertiesResponseOutput) ResourceRequests() GatewayOperatorResourceRequestsResponseOutput {
	return o.ApplyT(func(v GatewayOperatorPropertiesResponse) GatewayOperatorResourceRequestsResponse {
		return v.ResourceRequests
	}).(GatewayOperatorResourceRequestsResponseOutput)
}

type GatewayOperatorResourceRequestsResponse struct {
	Cpu           string `pulumi:"cpu"`
	InstanceCount int    `pulumi:"instanceCount"`
	Memory        string `pulumi:"memory"`
}

type GatewayOperatorResourceRequestsResponseOutput struct{ *pulumi.OutputState }

func (GatewayOperatorResourceRequestsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayOperatorResourceRequestsResponse)(nil)).Elem()
}

func (o GatewayOperatorResourceRequestsResponseOutput) ToGatewayOperatorResourceRequestsResponseOutput() GatewayOperatorResourceRequestsResponseOutput {
	return o
}

func (o GatewayOperatorResourceRequestsResponseOutput) ToGatewayOperatorResourceRequestsResponseOutputWithContext(ctx context.Context) GatewayOperatorResourceRequestsResponseOutput {
	return o
}

func (o GatewayOperatorResourceRequestsResponseOutput) Cpu() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayOperatorResourceRequestsResponse) string { return v.Cpu }).(pulumi.StringOutput)
}

func (o GatewayOperatorResourceRequestsResponseOutput) InstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v GatewayOperatorResourceRequestsResponse) int { return v.InstanceCount }).(pulumi.IntOutput)
}

func (o GatewayOperatorResourceRequestsResponseOutput) Memory() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayOperatorResourceRequestsResponse) string { return v.Memory }).(pulumi.StringOutput)
}

type GatewayProperties struct {
	ApiMetadataProperties *GatewayApiMetadataProperties `pulumi:"apiMetadataProperties"`
	CorsProperties        *GatewayCorsProperties        `pulumi:"corsProperties"`
	HttpsOnly             *bool                         `pulumi:"httpsOnly"`
	Public                *bool                         `pulumi:"public"`
	ResourceRequests      *GatewayResourceRequests      `pulumi:"resourceRequests"`
	SsoProperties         *SsoProperties                `pulumi:"ssoProperties"`
}


func (val *GatewayProperties) Defaults() *GatewayProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HttpsOnly) {
		httpsOnly_ := false
		tmp.HttpsOnly = &httpsOnly_
	}
	if isZero(tmp.Public) {
		public_ := false
		tmp.Public = &public_
	}
	return &tmp
}





type GatewayPropertiesInput interface {
	pulumi.Input

	ToGatewayPropertiesOutput() GatewayPropertiesOutput
	ToGatewayPropertiesOutputWithContext(context.Context) GatewayPropertiesOutput
}

type GatewayPropertiesArgs struct {
	ApiMetadataProperties GatewayApiMetadataPropertiesPtrInput `pulumi:"apiMetadataProperties"`
	CorsProperties        GatewayCorsPropertiesPtrInput        `pulumi:"corsProperties"`
	HttpsOnly             pulumi.BoolPtrInput                  `pulumi:"httpsOnly"`
	Public                pulumi.BoolPtrInput                  `pulumi:"public"`
	ResourceRequests      GatewayResourceRequestsPtrInput      `pulumi:"resourceRequests"`
	SsoProperties         SsoPropertiesPtrInput                `pulumi:"ssoProperties"`
}


func (val *GatewayPropertiesArgs) Defaults() *GatewayPropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HttpsOnly) {
		tmp.HttpsOnly = pulumi.BoolPtr(false)
	}
	if isZero(tmp.Public) {
		tmp.Public = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (GatewayPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayProperties)(nil)).Elem()
}

func (i GatewayPropertiesArgs) ToGatewayPropertiesOutput() GatewayPropertiesOutput {
	return i.ToGatewayPropertiesOutputWithContext(context.Background())
}

func (i GatewayPropertiesArgs) ToGatewayPropertiesOutputWithContext(ctx context.Context) GatewayPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPropertiesOutput)
}

func (i GatewayPropertiesArgs) ToGatewayPropertiesPtrOutput() GatewayPropertiesPtrOutput {
	return i.ToGatewayPropertiesPtrOutputWithContext(context.Background())
}

func (i GatewayPropertiesArgs) ToGatewayPropertiesPtrOutputWithContext(ctx context.Context) GatewayPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPropertiesOutput).ToGatewayPropertiesPtrOutputWithContext(ctx)
}









type GatewayPropertiesPtrInput interface {
	pulumi.Input

	ToGatewayPropertiesPtrOutput() GatewayPropertiesPtrOutput
	ToGatewayPropertiesPtrOutputWithContext(context.Context) GatewayPropertiesPtrOutput
}

type gatewayPropertiesPtrType GatewayPropertiesArgs

func GatewayPropertiesPtr(v *GatewayPropertiesArgs) GatewayPropertiesPtrInput {
	return (*gatewayPropertiesPtrType)(v)
}

func (*gatewayPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayProperties)(nil)).Elem()
}

func (i *gatewayPropertiesPtrType) ToGatewayPropertiesPtrOutput() GatewayPropertiesPtrOutput {
	return i.ToGatewayPropertiesPtrOutputWithContext(context.Background())
}

func (i *gatewayPropertiesPtrType) ToGatewayPropertiesPtrOutputWithContext(ctx context.Context) GatewayPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPropertiesPtrOutput)
}

type GatewayPropertiesOutput struct{ *pulumi.OutputState }

func (GatewayPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayProperties)(nil)).Elem()
}

func (o GatewayPropertiesOutput) ToGatewayPropertiesOutput() GatewayPropertiesOutput {
	return o
}

func (o GatewayPropertiesOutput) ToGatewayPropertiesOutputWithContext(ctx context.Context) GatewayPropertiesOutput {
	return o
}

func (o GatewayPropertiesOutput) ToGatewayPropertiesPtrOutput() GatewayPropertiesPtrOutput {
	return o.ToGatewayPropertiesPtrOutputWithContext(context.Background())
}

func (o GatewayPropertiesOutput) ToGatewayPropertiesPtrOutputWithContext(ctx context.Context) GatewayPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GatewayProperties) *GatewayProperties {
		return &v
	}).(GatewayPropertiesPtrOutput)
}

func (o GatewayPropertiesOutput) ApiMetadataProperties() GatewayApiMetadataPropertiesPtrOutput {
	return o.ApplyT(func(v GatewayProperties) *GatewayApiMetadataProperties { return v.ApiMetadataProperties }).(GatewayApiMetadataPropertiesPtrOutput)
}

func (o GatewayPropertiesOutput) CorsProperties() GatewayCorsPropertiesPtrOutput {
	return o.ApplyT(func(v GatewayProperties) *GatewayCorsProperties { return v.CorsProperties }).(GatewayCorsPropertiesPtrOutput)
}

func (o GatewayPropertiesOutput) HttpsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GatewayProperties) *bool { return v.HttpsOnly }).(pulumi.BoolPtrOutput)
}

func (o GatewayPropertiesOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GatewayProperties) *bool { return v.Public }).(pulumi.BoolPtrOutput)
}

func (o GatewayPropertiesOutput) ResourceRequests() GatewayResourceRequestsPtrOutput {
	return o.ApplyT(func(v GatewayProperties) *GatewayResourceRequests { return v.ResourceRequests }).(GatewayResourceRequestsPtrOutput)
}

func (o GatewayPropertiesOutput) SsoProperties() SsoPropertiesPtrOutput {
	return o.ApplyT(func(v GatewayProperties) *SsoProperties { return v.SsoProperties }).(SsoPropertiesPtrOutput)
}

type GatewayPropertiesPtrOutput struct{ *pulumi.OutputState }

func (GatewayPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayProperties)(nil)).Elem()
}

func (o GatewayPropertiesPtrOutput) ToGatewayPropertiesPtrOutput() GatewayPropertiesPtrOutput {
	return o
}

func (o GatewayPropertiesPtrOutput) ToGatewayPropertiesPtrOutputWithContext(ctx context.Context) GatewayPropertiesPtrOutput {
	return o
}

func (o GatewayPropertiesPtrOutput) Elem() GatewayPropertiesOutput {
	return o.ApplyT(func(v *GatewayProperties) GatewayProperties {
		if v != nil {
			return *v
		}
		var ret GatewayProperties
		return ret
	}).(GatewayPropertiesOutput)
}

func (o GatewayPropertiesPtrOutput) ApiMetadataProperties() GatewayApiMetadataPropertiesPtrOutput {
	return o.ApplyT(func(v *GatewayProperties) *GatewayApiMetadataProperties {
		if v == nil {
			return nil
		}
		return v.ApiMetadataProperties
	}).(GatewayApiMetadataPropertiesPtrOutput)
}

func (o GatewayPropertiesPtrOutput) CorsProperties() GatewayCorsPropertiesPtrOutput {
	return o.ApplyT(func(v *GatewayProperties) *GatewayCorsProperties {
		if v == nil {
			return nil
		}
		return v.CorsProperties
	}).(GatewayCorsPropertiesPtrOutput)
}

func (o GatewayPropertiesPtrOutput) HttpsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GatewayProperties) *bool {
		if v == nil {
			return nil
		}
		return v.HttpsOnly
	}).(pulumi.BoolPtrOutput)
}

func (o GatewayPropertiesPtrOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GatewayProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Public
	}).(pulumi.BoolPtrOutput)
}

func (o GatewayPropertiesPtrOutput) ResourceRequests() GatewayResourceRequestsPtrOutput {
	return o.ApplyT(func(v *GatewayProperties) *GatewayResourceRequests {
		if v == nil {
			return nil
		}
		return v.ResourceRequests
	}).(GatewayResourceRequestsPtrOutput)
}

func (o GatewayPropertiesPtrOutput) SsoProperties() SsoPropertiesPtrOutput {
	return o.ApplyT(func(v *GatewayProperties) *SsoProperties {
		if v == nil {
			return nil
		}
		return v.SsoProperties
	}).(SsoPropertiesPtrOutput)
}

type GatewayPropertiesResponse struct {
	ApiMetadataProperties *GatewayApiMetadataPropertiesResponse `pulumi:"apiMetadataProperties"`
	CorsProperties        *GatewayCorsPropertiesResponse        `pulumi:"corsProperties"`
	HttpsOnly             *bool                                 `pulumi:"httpsOnly"`
	Instances             []GatewayInstanceResponse             `pulumi:"instances"`
	OperatorProperties    GatewayOperatorPropertiesResponse     `pulumi:"operatorProperties"`
	ProvisioningState     string                                `pulumi:"provisioningState"`
	Public                *bool                                 `pulumi:"public"`
	ResourceRequests      *GatewayResourceRequestsResponse      `pulumi:"resourceRequests"`
	SsoProperties         *SsoPropertiesResponse                `pulumi:"ssoProperties"`
	Url                   string                                `pulumi:"url"`
}


func (val *GatewayPropertiesResponse) Defaults() *GatewayPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HttpsOnly) {
		httpsOnly_ := false
		tmp.HttpsOnly = &httpsOnly_
	}
	if isZero(tmp.Public) {
		public_ := false
		tmp.Public = &public_
	}
	return &tmp
}

type GatewayPropertiesResponseOutput struct{ *pulumi.OutputState }

func (GatewayPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayPropertiesResponse)(nil)).Elem()
}

func (o GatewayPropertiesResponseOutput) ToGatewayPropertiesResponseOutput() GatewayPropertiesResponseOutput {
	return o
}

func (o GatewayPropertiesResponseOutput) ToGatewayPropertiesResponseOutputWithContext(ctx context.Context) GatewayPropertiesResponseOutput {
	return o
}

func (o GatewayPropertiesResponseOutput) ApiMetadataProperties() GatewayApiMetadataPropertiesResponsePtrOutput {
	return o.ApplyT(func(v GatewayPropertiesResponse) *GatewayApiMetadataPropertiesResponse {
		return v.ApiMetadataProperties
	}).(GatewayApiMetadataPropertiesResponsePtrOutput)
}

func (o GatewayPropertiesResponseOutput) CorsProperties() GatewayCorsPropertiesResponsePtrOutput {
	return o.ApplyT(func(v GatewayPropertiesResponse) *GatewayCorsPropertiesResponse { return v.CorsProperties }).(GatewayCorsPropertiesResponsePtrOutput)
}

func (o GatewayPropertiesResponseOutput) HttpsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GatewayPropertiesResponse) *bool { return v.HttpsOnly }).(pulumi.BoolPtrOutput)
}

func (o GatewayPropertiesResponseOutput) Instances() GatewayInstanceResponseArrayOutput {
	return o.ApplyT(func(v GatewayPropertiesResponse) []GatewayInstanceResponse { return v.Instances }).(GatewayInstanceResponseArrayOutput)
}

func (o GatewayPropertiesResponseOutput) OperatorProperties() GatewayOperatorPropertiesResponseOutput {
	return o.ApplyT(func(v GatewayPropertiesResponse) GatewayOperatorPropertiesResponse { return v.OperatorProperties }).(GatewayOperatorPropertiesResponseOutput)
}

func (o GatewayPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GatewayPropertiesResponseOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GatewayPropertiesResponse) *bool { return v.Public }).(pulumi.BoolPtrOutput)
}

func (o GatewayPropertiesResponseOutput) ResourceRequests() GatewayResourceRequestsResponsePtrOutput {
	return o.ApplyT(func(v GatewayPropertiesResponse) *GatewayResourceRequestsResponse { return v.ResourceRequests }).(GatewayResourceRequestsResponsePtrOutput)
}

func (o GatewayPropertiesResponseOutput) SsoProperties() SsoPropertiesResponsePtrOutput {
	return o.ApplyT(func(v GatewayPropertiesResponse) *SsoPropertiesResponse { return v.SsoProperties }).(SsoPropertiesResponsePtrOutput)
}

func (o GatewayPropertiesResponseOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayPropertiesResponse) string { return v.Url }).(pulumi.StringOutput)
}

type GatewayResourceRequests struct {
	Cpu    *string `pulumi:"cpu"`
	Memory *string `pulumi:"memory"`
}





type GatewayResourceRequestsInput interface {
	pulumi.Input

	ToGatewayResourceRequestsOutput() GatewayResourceRequestsOutput
	ToGatewayResourceRequestsOutputWithContext(context.Context) GatewayResourceRequestsOutput
}

type GatewayResourceRequestsArgs struct {
	Cpu    pulumi.StringPtrInput `pulumi:"cpu"`
	Memory pulumi.StringPtrInput `pulumi:"memory"`
}

func (GatewayResourceRequestsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayResourceRequests)(nil)).Elem()
}

func (i GatewayResourceRequestsArgs) ToGatewayResourceRequestsOutput() GatewayResourceRequestsOutput {
	return i.ToGatewayResourceRequestsOutputWithContext(context.Background())
}

func (i GatewayResourceRequestsArgs) ToGatewayResourceRequestsOutputWithContext(ctx context.Context) GatewayResourceRequestsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayResourceRequestsOutput)
}

func (i GatewayResourceRequestsArgs) ToGatewayResourceRequestsPtrOutput() GatewayResourceRequestsPtrOutput {
	return i.ToGatewayResourceRequestsPtrOutputWithContext(context.Background())
}

func (i GatewayResourceRequestsArgs) ToGatewayResourceRequestsPtrOutputWithContext(ctx context.Context) GatewayResourceRequestsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayResourceRequestsOutput).ToGatewayResourceRequestsPtrOutputWithContext(ctx)
}









type GatewayResourceRequestsPtrInput interface {
	pulumi.Input

	ToGatewayResourceRequestsPtrOutput() GatewayResourceRequestsPtrOutput
	ToGatewayResourceRequestsPtrOutputWithContext(context.Context) GatewayResourceRequestsPtrOutput
}

type gatewayResourceRequestsPtrType GatewayResourceRequestsArgs

func GatewayResourceRequestsPtr(v *GatewayResourceRequestsArgs) GatewayResourceRequestsPtrInput {
	return (*gatewayResourceRequestsPtrType)(v)
}

func (*gatewayResourceRequestsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayResourceRequests)(nil)).Elem()
}

func (i *gatewayResourceRequestsPtrType) ToGatewayResourceRequestsPtrOutput() GatewayResourceRequestsPtrOutput {
	return i.ToGatewayResourceRequestsPtrOutputWithContext(context.Background())
}

func (i *gatewayResourceRequestsPtrType) ToGatewayResourceRequestsPtrOutputWithContext(ctx context.Context) GatewayResourceRequestsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayResourceRequestsPtrOutput)
}

type GatewayResourceRequestsOutput struct{ *pulumi.OutputState }

func (GatewayResourceRequestsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayResourceRequests)(nil)).Elem()
}

func (o GatewayResourceRequestsOutput) ToGatewayResourceRequestsOutput() GatewayResourceRequestsOutput {
	return o
}

func (o GatewayResourceRequestsOutput) ToGatewayResourceRequestsOutputWithContext(ctx context.Context) GatewayResourceRequestsOutput {
	return o
}

func (o GatewayResourceRequestsOutput) ToGatewayResourceRequestsPtrOutput() GatewayResourceRequestsPtrOutput {
	return o.ToGatewayResourceRequestsPtrOutputWithContext(context.Background())
}

func (o GatewayResourceRequestsOutput) ToGatewayResourceRequestsPtrOutputWithContext(ctx context.Context) GatewayResourceRequestsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GatewayResourceRequests) *GatewayResourceRequests {
		return &v
	}).(GatewayResourceRequestsPtrOutput)
}

func (o GatewayResourceRequestsOutput) Cpu() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayResourceRequests) *string { return v.Cpu }).(pulumi.StringPtrOutput)
}

func (o GatewayResourceRequestsOutput) Memory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayResourceRequests) *string { return v.Memory }).(pulumi.StringPtrOutput)
}

type GatewayResourceRequestsPtrOutput struct{ *pulumi.OutputState }

func (GatewayResourceRequestsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayResourceRequests)(nil)).Elem()
}

func (o GatewayResourceRequestsPtrOutput) ToGatewayResourceRequestsPtrOutput() GatewayResourceRequestsPtrOutput {
	return o
}

func (o GatewayResourceRequestsPtrOutput) ToGatewayResourceRequestsPtrOutputWithContext(ctx context.Context) GatewayResourceRequestsPtrOutput {
	return o
}

func (o GatewayResourceRequestsPtrOutput) Elem() GatewayResourceRequestsOutput {
	return o.ApplyT(func(v *GatewayResourceRequests) GatewayResourceRequests {
		if v != nil {
			return *v
		}
		var ret GatewayResourceRequests
		return ret
	}).(GatewayResourceRequestsOutput)
}

func (o GatewayResourceRequestsPtrOutput) Cpu() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayResourceRequests) *string {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.StringPtrOutput)
}

func (o GatewayResourceRequestsPtrOutput) Memory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayResourceRequests) *string {
		if v == nil {
			return nil
		}
		return v.Memory
	}).(pulumi.StringPtrOutput)
}

type GatewayResourceRequestsResponse struct {
	Cpu    *string `pulumi:"cpu"`
	Memory *string `pulumi:"memory"`
}

type GatewayResourceRequestsResponseOutput struct{ *pulumi.OutputState }

func (GatewayResourceRequestsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayResourceRequestsResponse)(nil)).Elem()
}

func (o GatewayResourceRequestsResponseOutput) ToGatewayResourceRequestsResponseOutput() GatewayResourceRequestsResponseOutput {
	return o
}

func (o GatewayResourceRequestsResponseOutput) ToGatewayResourceRequestsResponseOutputWithContext(ctx context.Context) GatewayResourceRequestsResponseOutput {
	return o
}

func (o GatewayResourceRequestsResponseOutput) Cpu() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayResourceRequestsResponse) *string { return v.Cpu }).(pulumi.StringPtrOutput)
}

func (o GatewayResourceRequestsResponseOutput) Memory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayResourceRequestsResponse) *string { return v.Memory }).(pulumi.StringPtrOutput)
}

type GatewayResourceRequestsResponsePtrOutput struct{ *pulumi.OutputState }

func (GatewayResourceRequestsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayResourceRequestsResponse)(nil)).Elem()
}

func (o GatewayResourceRequestsResponsePtrOutput) ToGatewayResourceRequestsResponsePtrOutput() GatewayResourceRequestsResponsePtrOutput {
	return o
}

func (o GatewayResourceRequestsResponsePtrOutput) ToGatewayResourceRequestsResponsePtrOutputWithContext(ctx context.Context) GatewayResourceRequestsResponsePtrOutput {
	return o
}

func (o GatewayResourceRequestsResponsePtrOutput) Elem() GatewayResourceRequestsResponseOutput {
	return o.ApplyT(func(v *GatewayResourceRequestsResponse) GatewayResourceRequestsResponse {
		if v != nil {
			return *v
		}
		var ret GatewayResourceRequestsResponse
		return ret
	}).(GatewayResourceRequestsResponseOutput)
}

func (o GatewayResourceRequestsResponsePtrOutput) Cpu() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayResourceRequestsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.StringPtrOutput)
}

func (o GatewayResourceRequestsResponsePtrOutput) Memory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayResourceRequestsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Memory
	}).(pulumi.StringPtrOutput)
}

type GatewayRouteConfigProperties struct {
	AppResourceId *string           `pulumi:"appResourceId"`
	Routes        []GatewayApiRoute `pulumi:"routes"`
}





type GatewayRouteConfigPropertiesInput interface {
	pulumi.Input

	ToGatewayRouteConfigPropertiesOutput() GatewayRouteConfigPropertiesOutput
	ToGatewayRouteConfigPropertiesOutputWithContext(context.Context) GatewayRouteConfigPropertiesOutput
}

type GatewayRouteConfigPropertiesArgs struct {
	AppResourceId pulumi.StringPtrInput     `pulumi:"appResourceId"`
	Routes        GatewayApiRouteArrayInput `pulumi:"routes"`
}

func (GatewayRouteConfigPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayRouteConfigProperties)(nil)).Elem()
}

func (i GatewayRouteConfigPropertiesArgs) ToGatewayRouteConfigPropertiesOutput() GatewayRouteConfigPropertiesOutput {
	return i.ToGatewayRouteConfigPropertiesOutputWithContext(context.Background())
}

func (i GatewayRouteConfigPropertiesArgs) ToGatewayRouteConfigPropertiesOutputWithContext(ctx context.Context) GatewayRouteConfigPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteConfigPropertiesOutput)
}

func (i GatewayRouteConfigPropertiesArgs) ToGatewayRouteConfigPropertiesPtrOutput() GatewayRouteConfigPropertiesPtrOutput {
	return i.ToGatewayRouteConfigPropertiesPtrOutputWithContext(context.Background())
}

func (i GatewayRouteConfigPropertiesArgs) ToGatewayRouteConfigPropertiesPtrOutputWithContext(ctx context.Context) GatewayRouteConfigPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteConfigPropertiesOutput).ToGatewayRouteConfigPropertiesPtrOutputWithContext(ctx)
}









type GatewayRouteConfigPropertiesPtrInput interface {
	pulumi.Input

	ToGatewayRouteConfigPropertiesPtrOutput() GatewayRouteConfigPropertiesPtrOutput
	ToGatewayRouteConfigPropertiesPtrOutputWithContext(context.Context) GatewayRouteConfigPropertiesPtrOutput
}

type gatewayRouteConfigPropertiesPtrType GatewayRouteConfigPropertiesArgs

func GatewayRouteConfigPropertiesPtr(v *GatewayRouteConfigPropertiesArgs) GatewayRouteConfigPropertiesPtrInput {
	return (*gatewayRouteConfigPropertiesPtrType)(v)
}

func (*gatewayRouteConfigPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayRouteConfigProperties)(nil)).Elem()
}

func (i *gatewayRouteConfigPropertiesPtrType) ToGatewayRouteConfigPropertiesPtrOutput() GatewayRouteConfigPropertiesPtrOutput {
	return i.ToGatewayRouteConfigPropertiesPtrOutputWithContext(context.Background())
}

func (i *gatewayRouteConfigPropertiesPtrType) ToGatewayRouteConfigPropertiesPtrOutputWithContext(ctx context.Context) GatewayRouteConfigPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteConfigPropertiesPtrOutput)
}

type GatewayRouteConfigPropertiesOutput struct{ *pulumi.OutputState }

func (GatewayRouteConfigPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayRouteConfigProperties)(nil)).Elem()
}

func (o GatewayRouteConfigPropertiesOutput) ToGatewayRouteConfigPropertiesOutput() GatewayRouteConfigPropertiesOutput {
	return o
}

func (o GatewayRouteConfigPropertiesOutput) ToGatewayRouteConfigPropertiesOutputWithContext(ctx context.Context) GatewayRouteConfigPropertiesOutput {
	return o
}

func (o GatewayRouteConfigPropertiesOutput) ToGatewayRouteConfigPropertiesPtrOutput() GatewayRouteConfigPropertiesPtrOutput {
	return o.ToGatewayRouteConfigPropertiesPtrOutputWithContext(context.Background())
}

func (o GatewayRouteConfigPropertiesOutput) ToGatewayRouteConfigPropertiesPtrOutputWithContext(ctx context.Context) GatewayRouteConfigPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GatewayRouteConfigProperties) *GatewayRouteConfigProperties {
		return &v
	}).(GatewayRouteConfigPropertiesPtrOutput)
}

func (o GatewayRouteConfigPropertiesOutput) AppResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayRouteConfigProperties) *string { return v.AppResourceId }).(pulumi.StringPtrOutput)
}

func (o GatewayRouteConfigPropertiesOutput) Routes() GatewayApiRouteArrayOutput {
	return o.ApplyT(func(v GatewayRouteConfigProperties) []GatewayApiRoute { return v.Routes }).(GatewayApiRouteArrayOutput)
}

type GatewayRouteConfigPropertiesPtrOutput struct{ *pulumi.OutputState }

func (GatewayRouteConfigPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayRouteConfigProperties)(nil)).Elem()
}

func (o GatewayRouteConfigPropertiesPtrOutput) ToGatewayRouteConfigPropertiesPtrOutput() GatewayRouteConfigPropertiesPtrOutput {
	return o
}

func (o GatewayRouteConfigPropertiesPtrOutput) ToGatewayRouteConfigPropertiesPtrOutputWithContext(ctx context.Context) GatewayRouteConfigPropertiesPtrOutput {
	return o
}

func (o GatewayRouteConfigPropertiesPtrOutput) Elem() GatewayRouteConfigPropertiesOutput {
	return o.ApplyT(func(v *GatewayRouteConfigProperties) GatewayRouteConfigProperties {
		if v != nil {
			return *v
		}
		var ret GatewayRouteConfigProperties
		return ret
	}).(GatewayRouteConfigPropertiesOutput)
}

func (o GatewayRouteConfigPropertiesPtrOutput) AppResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayRouteConfigProperties) *string {
		if v == nil {
			return nil
		}
		return v.AppResourceId
	}).(pulumi.StringPtrOutput)
}

func (o GatewayRouteConfigPropertiesPtrOutput) Routes() GatewayApiRouteArrayOutput {
	return o.ApplyT(func(v *GatewayRouteConfigProperties) []GatewayApiRoute {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(GatewayApiRouteArrayOutput)
}

type GatewayRouteConfigPropertiesResponse struct {
	AppResourceId     *string                   `pulumi:"appResourceId"`
	ProvisioningState string                    `pulumi:"provisioningState"`
	Routes            []GatewayApiRouteResponse `pulumi:"routes"`
}

type GatewayRouteConfigPropertiesResponseOutput struct{ *pulumi.OutputState }

func (GatewayRouteConfigPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayRouteConfigPropertiesResponse)(nil)).Elem()
}

func (o GatewayRouteConfigPropertiesResponseOutput) ToGatewayRouteConfigPropertiesResponseOutput() GatewayRouteConfigPropertiesResponseOutput {
	return o
}

func (o GatewayRouteConfigPropertiesResponseOutput) ToGatewayRouteConfigPropertiesResponseOutputWithContext(ctx context.Context) GatewayRouteConfigPropertiesResponseOutput {
	return o
}

func (o GatewayRouteConfigPropertiesResponseOutput) AppResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayRouteConfigPropertiesResponse) *string { return v.AppResourceId }).(pulumi.StringPtrOutput)
}

func (o GatewayRouteConfigPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayRouteConfigPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GatewayRouteConfigPropertiesResponseOutput) Routes() GatewayApiRouteResponseArrayOutput {
	return o.ApplyT(func(v GatewayRouteConfigPropertiesResponse) []GatewayApiRouteResponse { return v.Routes }).(GatewayApiRouteResponseArrayOutput)
}

type ManagedIdentityProperties struct {
	PrincipalId *string `pulumi:"principalId"`
	TenantId    *string `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type ManagedIdentityPropertiesInput interface {
	pulumi.Input

	ToManagedIdentityPropertiesOutput() ManagedIdentityPropertiesOutput
	ToManagedIdentityPropertiesOutputWithContext(context.Context) ManagedIdentityPropertiesOutput
}

type ManagedIdentityPropertiesArgs struct {
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
	TenantId    pulumi.StringPtrInput `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (ManagedIdentityPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityProperties)(nil)).Elem()
}

func (i ManagedIdentityPropertiesArgs) ToManagedIdentityPropertiesOutput() ManagedIdentityPropertiesOutput {
	return i.ToManagedIdentityPropertiesOutputWithContext(context.Background())
}

func (i ManagedIdentityPropertiesArgs) ToManagedIdentityPropertiesOutputWithContext(ctx context.Context) ManagedIdentityPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityPropertiesOutput)
}

func (i ManagedIdentityPropertiesArgs) ToManagedIdentityPropertiesPtrOutput() ManagedIdentityPropertiesPtrOutput {
	return i.ToManagedIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i ManagedIdentityPropertiesArgs) ToManagedIdentityPropertiesPtrOutputWithContext(ctx context.Context) ManagedIdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityPropertiesOutput).ToManagedIdentityPropertiesPtrOutputWithContext(ctx)
}









type ManagedIdentityPropertiesPtrInput interface {
	pulumi.Input

	ToManagedIdentityPropertiesPtrOutput() ManagedIdentityPropertiesPtrOutput
	ToManagedIdentityPropertiesPtrOutputWithContext(context.Context) ManagedIdentityPropertiesPtrOutput
}

type managedIdentityPropertiesPtrType ManagedIdentityPropertiesArgs

func ManagedIdentityPropertiesPtr(v *ManagedIdentityPropertiesArgs) ManagedIdentityPropertiesPtrInput {
	return (*managedIdentityPropertiesPtrType)(v)
}

func (*managedIdentityPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityProperties)(nil)).Elem()
}

func (i *managedIdentityPropertiesPtrType) ToManagedIdentityPropertiesPtrOutput() ManagedIdentityPropertiesPtrOutput {
	return i.ToManagedIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i *managedIdentityPropertiesPtrType) ToManagedIdentityPropertiesPtrOutputWithContext(ctx context.Context) ManagedIdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityPropertiesPtrOutput)
}

type ManagedIdentityPropertiesOutput struct{ *pulumi.OutputState }

func (ManagedIdentityPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityProperties)(nil)).Elem()
}

func (o ManagedIdentityPropertiesOutput) ToManagedIdentityPropertiesOutput() ManagedIdentityPropertiesOutput {
	return o
}

func (o ManagedIdentityPropertiesOutput) ToManagedIdentityPropertiesOutputWithContext(ctx context.Context) ManagedIdentityPropertiesOutput {
	return o
}

func (o ManagedIdentityPropertiesOutput) ToManagedIdentityPropertiesPtrOutput() ManagedIdentityPropertiesPtrOutput {
	return o.ToManagedIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (o ManagedIdentityPropertiesOutput) ToManagedIdentityPropertiesPtrOutputWithContext(ctx context.Context) ManagedIdentityPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentityProperties) *ManagedIdentityProperties {
		return &v
	}).(ManagedIdentityPropertiesPtrOutput)
}

func (o ManagedIdentityPropertiesOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityProperties) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityPropertiesOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityProperties) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityPropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ManagedIdentityPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityProperties)(nil)).Elem()
}

func (o ManagedIdentityPropertiesPtrOutput) ToManagedIdentityPropertiesPtrOutput() ManagedIdentityPropertiesPtrOutput {
	return o
}

func (o ManagedIdentityPropertiesPtrOutput) ToManagedIdentityPropertiesPtrOutputWithContext(ctx context.Context) ManagedIdentityPropertiesPtrOutput {
	return o
}

func (o ManagedIdentityPropertiesPtrOutput) Elem() ManagedIdentityPropertiesOutput {
	return o.ApplyT(func(v *ManagedIdentityProperties) ManagedIdentityProperties {
		if v != nil {
			return *v
		}
		var ret ManagedIdentityProperties
		return ret
	}).(ManagedIdentityPropertiesOutput)
}

func (o ManagedIdentityPropertiesPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityPropertiesPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityPropertiesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ManagedIdentityPropertiesResponse struct {
	PrincipalId *string `pulumi:"principalId"`
	TenantId    *string `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type ManagedIdentityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ManagedIdentityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityPropertiesResponse)(nil)).Elem()
}

func (o ManagedIdentityPropertiesResponseOutput) ToManagedIdentityPropertiesResponseOutput() ManagedIdentityPropertiesResponseOutput {
	return o
}

func (o ManagedIdentityPropertiesResponseOutput) ToManagedIdentityPropertiesResponseOutputWithContext(ctx context.Context) ManagedIdentityPropertiesResponseOutput {
	return o
}

func (o ManagedIdentityPropertiesResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityPropertiesResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityPropertiesResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityPropertiesResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ManagedIdentityPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityPropertiesResponse)(nil)).Elem()
}

func (o ManagedIdentityPropertiesResponsePtrOutput) ToManagedIdentityPropertiesResponsePtrOutput() ManagedIdentityPropertiesResponsePtrOutput {
	return o
}

func (o ManagedIdentityPropertiesResponsePtrOutput) ToManagedIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityPropertiesResponsePtrOutput {
	return o
}

func (o ManagedIdentityPropertiesResponsePtrOutput) Elem() ManagedIdentityPropertiesResponseOutput {
	return o.ApplyT(func(v *ManagedIdentityPropertiesResponse) ManagedIdentityPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ManagedIdentityPropertiesResponse
		return ret
	}).(ManagedIdentityPropertiesResponseOutput)
}

func (o ManagedIdentityPropertiesResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityPropertiesResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityPropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type NetworkProfile struct {
	AppNetworkResourceGroup            *string `pulumi:"appNetworkResourceGroup"`
	AppSubnetId                        *string `pulumi:"appSubnetId"`
	ServiceCidr                        *string `pulumi:"serviceCidr"`
	ServiceRuntimeNetworkResourceGroup *string `pulumi:"serviceRuntimeNetworkResourceGroup"`
	ServiceRuntimeSubnetId             *string `pulumi:"serviceRuntimeSubnetId"`
}





type NetworkProfileInput interface {
	pulumi.Input

	ToNetworkProfileOutput() NetworkProfileOutput
	ToNetworkProfileOutputWithContext(context.Context) NetworkProfileOutput
}

type NetworkProfileArgs struct {
	AppNetworkResourceGroup            pulumi.StringPtrInput `pulumi:"appNetworkResourceGroup"`
	AppSubnetId                        pulumi.StringPtrInput `pulumi:"appSubnetId"`
	ServiceCidr                        pulumi.StringPtrInput `pulumi:"serviceCidr"`
	ServiceRuntimeNetworkResourceGroup pulumi.StringPtrInput `pulumi:"serviceRuntimeNetworkResourceGroup"`
	ServiceRuntimeSubnetId             pulumi.StringPtrInput `pulumi:"serviceRuntimeSubnetId"`
}

func (NetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (i NetworkProfileArgs) ToNetworkProfileOutput() NetworkProfileOutput {
	return i.ToNetworkProfileOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput)
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput).ToNetworkProfilePtrOutputWithContext(ctx)
}









type NetworkProfilePtrInput interface {
	pulumi.Input

	ToNetworkProfilePtrOutput() NetworkProfilePtrOutput
	ToNetworkProfilePtrOutputWithContext(context.Context) NetworkProfilePtrOutput
}

type networkProfilePtrType NetworkProfileArgs

func NetworkProfilePtr(v *NetworkProfileArgs) NetworkProfilePtrInput {
	return (*networkProfilePtrType)(v)
}

func (*networkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfilePtrOutput)
}

type NetworkProfileOutput struct{ *pulumi.OutputState }

func (NetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (o NetworkProfileOutput) ToNetworkProfileOutput() NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkProfile) *NetworkProfile {
		return &v
	}).(NetworkProfilePtrOutput)
}

func (o NetworkProfileOutput) AppNetworkResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.AppNetworkResourceGroup }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) AppSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.AppSubnetId }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.ServiceCidr }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) ServiceRuntimeNetworkResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.ServiceRuntimeNetworkResourceGroup }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) ServiceRuntimeSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.ServiceRuntimeSubnetId }).(pulumi.StringPtrOutput)
}

type NetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) Elem() NetworkProfileOutput {
	return o.ApplyT(func(v *NetworkProfile) NetworkProfile {
		if v != nil {
			return *v
		}
		var ret NetworkProfile
		return ret
	}).(NetworkProfileOutput)
}

func (o NetworkProfilePtrOutput) AppNetworkResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.AppNetworkResourceGroup
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfilePtrOutput) AppSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.AppSubnetId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfilePtrOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCidr
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfilePtrOutput) ServiceRuntimeNetworkResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.ServiceRuntimeNetworkResourceGroup
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfilePtrOutput) ServiceRuntimeSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.ServiceRuntimeSubnetId
	}).(pulumi.StringPtrOutput)
}

type NetworkProfileResponse struct {
	AppNetworkResourceGroup            *string                           `pulumi:"appNetworkResourceGroup"`
	AppSubnetId                        *string                           `pulumi:"appSubnetId"`
	OutboundIPs                        NetworkProfileResponseOutboundIPs `pulumi:"outboundIPs"`
	RequiredTraffics                   []RequiredTrafficResponse         `pulumi:"requiredTraffics"`
	ServiceCidr                        *string                           `pulumi:"serviceCidr"`
	ServiceRuntimeNetworkResourceGroup *string                           `pulumi:"serviceRuntimeNetworkResourceGroup"`
	ServiceRuntimeSubnetId             *string                           `pulumi:"serviceRuntimeSubnetId"`
}

type NetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutput() NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutputWithContext(ctx context.Context) NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) AppNetworkResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.AppNetworkResourceGroup }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponseOutput) AppSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.AppSubnetId }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponseOutput) OutboundIPs() NetworkProfileResponseOutboundIPsOutput {
	return o.ApplyT(func(v NetworkProfileResponse) NetworkProfileResponseOutboundIPs { return v.OutboundIPs }).(NetworkProfileResponseOutboundIPsOutput)
}

func (o NetworkProfileResponseOutput) RequiredTraffics() RequiredTrafficResponseArrayOutput {
	return o.ApplyT(func(v NetworkProfileResponse) []RequiredTrafficResponse { return v.RequiredTraffics }).(RequiredTrafficResponseArrayOutput)
}

func (o NetworkProfileResponseOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.ServiceCidr }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponseOutput) ServiceRuntimeNetworkResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.ServiceRuntimeNetworkResourceGroup }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponseOutput) ServiceRuntimeSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.ServiceRuntimeSubnetId }).(pulumi.StringPtrOutput)
}

type NetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutputWithContext(ctx context.Context) NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) Elem() NetworkProfileResponseOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) NetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret NetworkProfileResponse
		return ret
	}).(NetworkProfileResponseOutput)
}

func (o NetworkProfileResponsePtrOutput) AppNetworkResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppNetworkResourceGroup
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) AppSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppSubnetId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) OutboundIPs() NetworkProfileResponseOutboundIPsPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *NetworkProfileResponseOutboundIPs {
		if v == nil {
			return nil
		}
		return &v.OutboundIPs
	}).(NetworkProfileResponseOutboundIPsPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) RequiredTraffics() RequiredTrafficResponseArrayOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) []RequiredTrafficResponse {
		if v == nil {
			return nil
		}
		return v.RequiredTraffics
	}).(RequiredTrafficResponseArrayOutput)
}

func (o NetworkProfileResponsePtrOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCidr
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) ServiceRuntimeNetworkResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceRuntimeNetworkResourceGroup
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) ServiceRuntimeSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceRuntimeSubnetId
	}).(pulumi.StringPtrOutput)
}

type NetworkProfileResponseOutboundIPs struct {
	PublicIPs []string `pulumi:"publicIPs"`
}

type NetworkProfileResponseOutboundIPsOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponseOutboundIPsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfileResponseOutboundIPs)(nil)).Elem()
}

func (o NetworkProfileResponseOutboundIPsOutput) ToNetworkProfileResponseOutboundIPsOutput() NetworkProfileResponseOutboundIPsOutput {
	return o
}

func (o NetworkProfileResponseOutboundIPsOutput) ToNetworkProfileResponseOutboundIPsOutputWithContext(ctx context.Context) NetworkProfileResponseOutboundIPsOutput {
	return o
}

func (o NetworkProfileResponseOutboundIPsOutput) PublicIPs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkProfileResponseOutboundIPs) []string { return v.PublicIPs }).(pulumi.StringArrayOutput)
}

type NetworkProfileResponseOutboundIPsPtrOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponseOutboundIPsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfileResponseOutboundIPs)(nil)).Elem()
}

func (o NetworkProfileResponseOutboundIPsPtrOutput) ToNetworkProfileResponseOutboundIPsPtrOutput() NetworkProfileResponseOutboundIPsPtrOutput {
	return o
}

func (o NetworkProfileResponseOutboundIPsPtrOutput) ToNetworkProfileResponseOutboundIPsPtrOutputWithContext(ctx context.Context) NetworkProfileResponseOutboundIPsPtrOutput {
	return o
}

func (o NetworkProfileResponseOutboundIPsPtrOutput) Elem() NetworkProfileResponseOutboundIPsOutput {
	return o.ApplyT(func(v *NetworkProfileResponseOutboundIPs) NetworkProfileResponseOutboundIPs {
		if v != nil {
			return *v
		}
		var ret NetworkProfileResponseOutboundIPs
		return ret
	}).(NetworkProfileResponseOutboundIPsOutput)
}

func (o NetworkProfileResponseOutboundIPsPtrOutput) PublicIPs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkProfileResponseOutboundIPs) []string {
		if v == nil {
			return nil
		}
		return v.PublicIPs
	}).(pulumi.StringArrayOutput)
}

type PersistentDisk struct {
	MountPath *string `pulumi:"mountPath"`
	SizeInGB  *int    `pulumi:"sizeInGB"`
}





type PersistentDiskInput interface {
	pulumi.Input

	ToPersistentDiskOutput() PersistentDiskOutput
	ToPersistentDiskOutputWithContext(context.Context) PersistentDiskOutput
}

type PersistentDiskArgs struct {
	MountPath pulumi.StringPtrInput `pulumi:"mountPath"`
	SizeInGB  pulumi.IntPtrInput    `pulumi:"sizeInGB"`
}

func (PersistentDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PersistentDisk)(nil)).Elem()
}

func (i PersistentDiskArgs) ToPersistentDiskOutput() PersistentDiskOutput {
	return i.ToPersistentDiskOutputWithContext(context.Background())
}

func (i PersistentDiskArgs) ToPersistentDiskOutputWithContext(ctx context.Context) PersistentDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentDiskOutput)
}

func (i PersistentDiskArgs) ToPersistentDiskPtrOutput() PersistentDiskPtrOutput {
	return i.ToPersistentDiskPtrOutputWithContext(context.Background())
}

func (i PersistentDiskArgs) ToPersistentDiskPtrOutputWithContext(ctx context.Context) PersistentDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentDiskOutput).ToPersistentDiskPtrOutputWithContext(ctx)
}









type PersistentDiskPtrInput interface {
	pulumi.Input

	ToPersistentDiskPtrOutput() PersistentDiskPtrOutput
	ToPersistentDiskPtrOutputWithContext(context.Context) PersistentDiskPtrOutput
}

type persistentDiskPtrType PersistentDiskArgs

func PersistentDiskPtr(v *PersistentDiskArgs) PersistentDiskPtrInput {
	return (*persistentDiskPtrType)(v)
}

func (*persistentDiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistentDisk)(nil)).Elem()
}

func (i *persistentDiskPtrType) ToPersistentDiskPtrOutput() PersistentDiskPtrOutput {
	return i.ToPersistentDiskPtrOutputWithContext(context.Background())
}

func (i *persistentDiskPtrType) ToPersistentDiskPtrOutputWithContext(ctx context.Context) PersistentDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentDiskPtrOutput)
}

type PersistentDiskOutput struct{ *pulumi.OutputState }

func (PersistentDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PersistentDisk)(nil)).Elem()
}

func (o PersistentDiskOutput) ToPersistentDiskOutput() PersistentDiskOutput {
	return o
}

func (o PersistentDiskOutput) ToPersistentDiskOutputWithContext(ctx context.Context) PersistentDiskOutput {
	return o
}

func (o PersistentDiskOutput) ToPersistentDiskPtrOutput() PersistentDiskPtrOutput {
	return o.ToPersistentDiskPtrOutputWithContext(context.Background())
}

func (o PersistentDiskOutput) ToPersistentDiskPtrOutputWithContext(ctx context.Context) PersistentDiskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PersistentDisk) *PersistentDisk {
		return &v
	}).(PersistentDiskPtrOutput)
}

func (o PersistentDiskOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PersistentDisk) *string { return v.MountPath }).(pulumi.StringPtrOutput)
}

func (o PersistentDiskOutput) SizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PersistentDisk) *int { return v.SizeInGB }).(pulumi.IntPtrOutput)
}

type PersistentDiskPtrOutput struct{ *pulumi.OutputState }

func (PersistentDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistentDisk)(nil)).Elem()
}

func (o PersistentDiskPtrOutput) ToPersistentDiskPtrOutput() PersistentDiskPtrOutput {
	return o
}

func (o PersistentDiskPtrOutput) ToPersistentDiskPtrOutputWithContext(ctx context.Context) PersistentDiskPtrOutput {
	return o
}

func (o PersistentDiskPtrOutput) Elem() PersistentDiskOutput {
	return o.ApplyT(func(v *PersistentDisk) PersistentDisk {
		if v != nil {
			return *v
		}
		var ret PersistentDisk
		return ret
	}).(PersistentDiskOutput)
}

func (o PersistentDiskPtrOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PersistentDisk) *string {
		if v == nil {
			return nil
		}
		return v.MountPath
	}).(pulumi.StringPtrOutput)
}

func (o PersistentDiskPtrOutput) SizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PersistentDisk) *int {
		if v == nil {
			return nil
		}
		return v.SizeInGB
	}).(pulumi.IntPtrOutput)
}

type PersistentDiskResponse struct {
	MountPath *string `pulumi:"mountPath"`
	SizeInGB  *int    `pulumi:"sizeInGB"`
	UsedInGB  int     `pulumi:"usedInGB"`
}

type PersistentDiskResponseOutput struct{ *pulumi.OutputState }

func (PersistentDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PersistentDiskResponse)(nil)).Elem()
}

func (o PersistentDiskResponseOutput) ToPersistentDiskResponseOutput() PersistentDiskResponseOutput {
	return o
}

func (o PersistentDiskResponseOutput) ToPersistentDiskResponseOutputWithContext(ctx context.Context) PersistentDiskResponseOutput {
	return o
}

func (o PersistentDiskResponseOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PersistentDiskResponse) *string { return v.MountPath }).(pulumi.StringPtrOutput)
}

func (o PersistentDiskResponseOutput) SizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PersistentDiskResponse) *int { return v.SizeInGB }).(pulumi.IntPtrOutput)
}

func (o PersistentDiskResponseOutput) UsedInGB() pulumi.IntOutput {
	return o.ApplyT(func(v PersistentDiskResponse) int { return v.UsedInGB }).(pulumi.IntOutput)
}

type PersistentDiskResponsePtrOutput struct{ *pulumi.OutputState }

func (PersistentDiskResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistentDiskResponse)(nil)).Elem()
}

func (o PersistentDiskResponsePtrOutput) ToPersistentDiskResponsePtrOutput() PersistentDiskResponsePtrOutput {
	return o
}

func (o PersistentDiskResponsePtrOutput) ToPersistentDiskResponsePtrOutputWithContext(ctx context.Context) PersistentDiskResponsePtrOutput {
	return o
}

func (o PersistentDiskResponsePtrOutput) Elem() PersistentDiskResponseOutput {
	return o.ApplyT(func(v *PersistentDiskResponse) PersistentDiskResponse {
		if v != nil {
			return *v
		}
		var ret PersistentDiskResponse
		return ret
	}).(PersistentDiskResponseOutput)
}

func (o PersistentDiskResponsePtrOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PersistentDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.MountPath
	}).(pulumi.StringPtrOutput)
}

func (o PersistentDiskResponsePtrOutput) SizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PersistentDiskResponse) *int {
		if v == nil {
			return nil
		}
		return v.SizeInGB
	}).(pulumi.IntPtrOutput)
}

func (o PersistentDiskResponsePtrOutput) UsedInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PersistentDiskResponse) *int {
		if v == nil {
			return nil
		}
		return &v.UsedInGB
	}).(pulumi.IntPtrOutput)
}

type RequiredTrafficResponse struct {
	Direction string   `pulumi:"direction"`
	Fqdns     []string `pulumi:"fqdns"`
	Ips       []string `pulumi:"ips"`
	Port      int      `pulumi:"port"`
	Protocol  string   `pulumi:"protocol"`
}

type RequiredTrafficResponseOutput struct{ *pulumi.OutputState }

func (RequiredTrafficResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequiredTrafficResponse)(nil)).Elem()
}

func (o RequiredTrafficResponseOutput) ToRequiredTrafficResponseOutput() RequiredTrafficResponseOutput {
	return o
}

func (o RequiredTrafficResponseOutput) ToRequiredTrafficResponseOutputWithContext(ctx context.Context) RequiredTrafficResponseOutput {
	return o
}

func (o RequiredTrafficResponseOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v RequiredTrafficResponse) string { return v.Direction }).(pulumi.StringOutput)
}

func (o RequiredTrafficResponseOutput) Fqdns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RequiredTrafficResponse) []string { return v.Fqdns }).(pulumi.StringArrayOutput)
}

func (o RequiredTrafficResponseOutput) Ips() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RequiredTrafficResponse) []string { return v.Ips }).(pulumi.StringArrayOutput)
}

func (o RequiredTrafficResponseOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v RequiredTrafficResponse) int { return v.Port }).(pulumi.IntOutput)
}

func (o RequiredTrafficResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v RequiredTrafficResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

type RequiredTrafficResponseArrayOutput struct{ *pulumi.OutputState }

func (RequiredTrafficResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RequiredTrafficResponse)(nil)).Elem()
}

func (o RequiredTrafficResponseArrayOutput) ToRequiredTrafficResponseArrayOutput() RequiredTrafficResponseArrayOutput {
	return o
}

func (o RequiredTrafficResponseArrayOutput) ToRequiredTrafficResponseArrayOutputWithContext(ctx context.Context) RequiredTrafficResponseArrayOutput {
	return o
}

func (o RequiredTrafficResponseArrayOutput) Index(i pulumi.IntInput) RequiredTrafficResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RequiredTrafficResponse {
		return vs[0].([]RequiredTrafficResponse)[vs[1].(int)]
	}).(RequiredTrafficResponseOutput)
}

type ServiceRegistryInstanceResponse struct {
	Name   string `pulumi:"name"`
	Status string `pulumi:"status"`
}

type ServiceRegistryInstanceResponseOutput struct{ *pulumi.OutputState }

func (ServiceRegistryInstanceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceRegistryInstanceResponse)(nil)).Elem()
}

func (o ServiceRegistryInstanceResponseOutput) ToServiceRegistryInstanceResponseOutput() ServiceRegistryInstanceResponseOutput {
	return o
}

func (o ServiceRegistryInstanceResponseOutput) ToServiceRegistryInstanceResponseOutputWithContext(ctx context.Context) ServiceRegistryInstanceResponseOutput {
	return o
}

func (o ServiceRegistryInstanceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceRegistryInstanceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceRegistryInstanceResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceRegistryInstanceResponse) string { return v.Status }).(pulumi.StringOutput)
}

type ServiceRegistryInstanceResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceRegistryInstanceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceRegistryInstanceResponse)(nil)).Elem()
}

func (o ServiceRegistryInstanceResponseArrayOutput) ToServiceRegistryInstanceResponseArrayOutput() ServiceRegistryInstanceResponseArrayOutput {
	return o
}

func (o ServiceRegistryInstanceResponseArrayOutput) ToServiceRegistryInstanceResponseArrayOutputWithContext(ctx context.Context) ServiceRegistryInstanceResponseArrayOutput {
	return o
}

func (o ServiceRegistryInstanceResponseArrayOutput) Index(i pulumi.IntInput) ServiceRegistryInstanceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceRegistryInstanceResponse {
		return vs[0].([]ServiceRegistryInstanceResponse)[vs[1].(int)]
	}).(ServiceRegistryInstanceResponseOutput)
}

type ServiceRegistryPropertiesResponse struct {
	Instances         []ServiceRegistryInstanceResponse       `pulumi:"instances"`
	ProvisioningState string                                  `pulumi:"provisioningState"`
	ResourceRequests  ServiceRegistryResourceRequestsResponse `pulumi:"resourceRequests"`
}

type ServiceRegistryPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ServiceRegistryPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceRegistryPropertiesResponse)(nil)).Elem()
}

func (o ServiceRegistryPropertiesResponseOutput) ToServiceRegistryPropertiesResponseOutput() ServiceRegistryPropertiesResponseOutput {
	return o
}

func (o ServiceRegistryPropertiesResponseOutput) ToServiceRegistryPropertiesResponseOutputWithContext(ctx context.Context) ServiceRegistryPropertiesResponseOutput {
	return o
}

func (o ServiceRegistryPropertiesResponseOutput) Instances() ServiceRegistryInstanceResponseArrayOutput {
	return o.ApplyT(func(v ServiceRegistryPropertiesResponse) []ServiceRegistryInstanceResponse { return v.Instances }).(ServiceRegistryInstanceResponseArrayOutput)
}

func (o ServiceRegistryPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceRegistryPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ServiceRegistryPropertiesResponseOutput) ResourceRequests() ServiceRegistryResourceRequestsResponseOutput {
	return o.ApplyT(func(v ServiceRegistryPropertiesResponse) ServiceRegistryResourceRequestsResponse {
		return v.ResourceRequests
	}).(ServiceRegistryResourceRequestsResponseOutput)
}

type ServiceRegistryResourceRequestsResponse struct {
	Cpu           string `pulumi:"cpu"`
	InstanceCount int    `pulumi:"instanceCount"`
	Memory        string `pulumi:"memory"`
}

type ServiceRegistryResourceRequestsResponseOutput struct{ *pulumi.OutputState }

func (ServiceRegistryResourceRequestsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceRegistryResourceRequestsResponse)(nil)).Elem()
}

func (o ServiceRegistryResourceRequestsResponseOutput) ToServiceRegistryResourceRequestsResponseOutput() ServiceRegistryResourceRequestsResponseOutput {
	return o
}

func (o ServiceRegistryResourceRequestsResponseOutput) ToServiceRegistryResourceRequestsResponseOutputWithContext(ctx context.Context) ServiceRegistryResourceRequestsResponseOutput {
	return o
}

func (o ServiceRegistryResourceRequestsResponseOutput) Cpu() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceRegistryResourceRequestsResponse) string { return v.Cpu }).(pulumi.StringOutput)
}

func (o ServiceRegistryResourceRequestsResponseOutput) InstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v ServiceRegistryResourceRequestsResponse) int { return v.InstanceCount }).(pulumi.IntOutput)
}

func (o ServiceRegistryResourceRequestsResponseOutput) Memory() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceRegistryResourceRequestsResponse) string { return v.Memory }).(pulumi.StringOutput)
}

type Sku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SsoProperties struct {
	ClientId     *string  `pulumi:"clientId"`
	ClientSecret *string  `pulumi:"clientSecret"`
	IssuerUri    *string  `pulumi:"issuerUri"`
	Scope        []string `pulumi:"scope"`
}





type SsoPropertiesInput interface {
	pulumi.Input

	ToSsoPropertiesOutput() SsoPropertiesOutput
	ToSsoPropertiesOutputWithContext(context.Context) SsoPropertiesOutput
}

type SsoPropertiesArgs struct {
	ClientId     pulumi.StringPtrInput   `pulumi:"clientId"`
	ClientSecret pulumi.StringPtrInput   `pulumi:"clientSecret"`
	IssuerUri    pulumi.StringPtrInput   `pulumi:"issuerUri"`
	Scope        pulumi.StringArrayInput `pulumi:"scope"`
}

func (SsoPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SsoProperties)(nil)).Elem()
}

func (i SsoPropertiesArgs) ToSsoPropertiesOutput() SsoPropertiesOutput {
	return i.ToSsoPropertiesOutputWithContext(context.Background())
}

func (i SsoPropertiesArgs) ToSsoPropertiesOutputWithContext(ctx context.Context) SsoPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SsoPropertiesOutput)
}

func (i SsoPropertiesArgs) ToSsoPropertiesPtrOutput() SsoPropertiesPtrOutput {
	return i.ToSsoPropertiesPtrOutputWithContext(context.Background())
}

func (i SsoPropertiesArgs) ToSsoPropertiesPtrOutputWithContext(ctx context.Context) SsoPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SsoPropertiesOutput).ToSsoPropertiesPtrOutputWithContext(ctx)
}









type SsoPropertiesPtrInput interface {
	pulumi.Input

	ToSsoPropertiesPtrOutput() SsoPropertiesPtrOutput
	ToSsoPropertiesPtrOutputWithContext(context.Context) SsoPropertiesPtrOutput
}

type ssoPropertiesPtrType SsoPropertiesArgs

func SsoPropertiesPtr(v *SsoPropertiesArgs) SsoPropertiesPtrInput {
	return (*ssoPropertiesPtrType)(v)
}

func (*ssoPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SsoProperties)(nil)).Elem()
}

func (i *ssoPropertiesPtrType) ToSsoPropertiesPtrOutput() SsoPropertiesPtrOutput {
	return i.ToSsoPropertiesPtrOutputWithContext(context.Background())
}

func (i *ssoPropertiesPtrType) ToSsoPropertiesPtrOutputWithContext(ctx context.Context) SsoPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SsoPropertiesPtrOutput)
}

type SsoPropertiesOutput struct{ *pulumi.OutputState }

func (SsoPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SsoProperties)(nil)).Elem()
}

func (o SsoPropertiesOutput) ToSsoPropertiesOutput() SsoPropertiesOutput {
	return o
}

func (o SsoPropertiesOutput) ToSsoPropertiesOutputWithContext(ctx context.Context) SsoPropertiesOutput {
	return o
}

func (o SsoPropertiesOutput) ToSsoPropertiesPtrOutput() SsoPropertiesPtrOutput {
	return o.ToSsoPropertiesPtrOutputWithContext(context.Background())
}

func (o SsoPropertiesOutput) ToSsoPropertiesPtrOutputWithContext(ctx context.Context) SsoPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SsoProperties) *SsoProperties {
		return &v
	}).(SsoPropertiesPtrOutput)
}

func (o SsoPropertiesOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SsoProperties) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o SsoPropertiesOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SsoProperties) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o SsoPropertiesOutput) IssuerUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SsoProperties) *string { return v.IssuerUri }).(pulumi.StringPtrOutput)
}

func (o SsoPropertiesOutput) Scope() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SsoProperties) []string { return v.Scope }).(pulumi.StringArrayOutput)
}

type SsoPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SsoPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SsoProperties)(nil)).Elem()
}

func (o SsoPropertiesPtrOutput) ToSsoPropertiesPtrOutput() SsoPropertiesPtrOutput {
	return o
}

func (o SsoPropertiesPtrOutput) ToSsoPropertiesPtrOutputWithContext(ctx context.Context) SsoPropertiesPtrOutput {
	return o
}

func (o SsoPropertiesPtrOutput) Elem() SsoPropertiesOutput {
	return o.ApplyT(func(v *SsoProperties) SsoProperties {
		if v != nil {
			return *v
		}
		var ret SsoProperties
		return ret
	}).(SsoPropertiesOutput)
}

func (o SsoPropertiesPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SsoProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o SsoPropertiesPtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SsoProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

func (o SsoPropertiesPtrOutput) IssuerUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SsoProperties) *string {
		if v == nil {
			return nil
		}
		return v.IssuerUri
	}).(pulumi.StringPtrOutput)
}

func (o SsoPropertiesPtrOutput) Scope() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SsoProperties) []string {
		if v == nil {
			return nil
		}
		return v.Scope
	}).(pulumi.StringArrayOutput)
}

type SsoPropertiesResponse struct {
	ClientId     *string  `pulumi:"clientId"`
	ClientSecret *string  `pulumi:"clientSecret"`
	IssuerUri    *string  `pulumi:"issuerUri"`
	Scope        []string `pulumi:"scope"`
}

type SsoPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SsoPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SsoPropertiesResponse)(nil)).Elem()
}

func (o SsoPropertiesResponseOutput) ToSsoPropertiesResponseOutput() SsoPropertiesResponseOutput {
	return o
}

func (o SsoPropertiesResponseOutput) ToSsoPropertiesResponseOutputWithContext(ctx context.Context) SsoPropertiesResponseOutput {
	return o
}

func (o SsoPropertiesResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SsoPropertiesResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o SsoPropertiesResponseOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SsoPropertiesResponse) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o SsoPropertiesResponseOutput) IssuerUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SsoPropertiesResponse) *string { return v.IssuerUri }).(pulumi.StringPtrOutput)
}

func (o SsoPropertiesResponseOutput) Scope() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SsoPropertiesResponse) []string { return v.Scope }).(pulumi.StringArrayOutput)
}

type SsoPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SsoPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SsoPropertiesResponse)(nil)).Elem()
}

func (o SsoPropertiesResponsePtrOutput) ToSsoPropertiesResponsePtrOutput() SsoPropertiesResponsePtrOutput {
	return o
}

func (o SsoPropertiesResponsePtrOutput) ToSsoPropertiesResponsePtrOutputWithContext(ctx context.Context) SsoPropertiesResponsePtrOutput {
	return o
}

func (o SsoPropertiesResponsePtrOutput) Elem() SsoPropertiesResponseOutput {
	return o.ApplyT(func(v *SsoPropertiesResponse) SsoPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret SsoPropertiesResponse
		return ret
	}).(SsoPropertiesResponseOutput)
}

func (o SsoPropertiesResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SsoPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o SsoPropertiesResponsePtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SsoPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

func (o SsoPropertiesResponsePtrOutput) IssuerUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SsoPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.IssuerUri
	}).(pulumi.StringPtrOutput)
}

func (o SsoPropertiesResponsePtrOutput) Scope() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SsoPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.Scope
	}).(pulumi.StringArrayOutput)
}

type StackProperties struct {
	Id      *string `pulumi:"id"`
	Version *string `pulumi:"version"`
}





type StackPropertiesInput interface {
	pulumi.Input

	ToStackPropertiesOutput() StackPropertiesOutput
	ToStackPropertiesOutputWithContext(context.Context) StackPropertiesOutput
}

type StackPropertiesArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (StackPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StackProperties)(nil)).Elem()
}

func (i StackPropertiesArgs) ToStackPropertiesOutput() StackPropertiesOutput {
	return i.ToStackPropertiesOutputWithContext(context.Background())
}

func (i StackPropertiesArgs) ToStackPropertiesOutputWithContext(ctx context.Context) StackPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackPropertiesOutput)
}

func (i StackPropertiesArgs) ToStackPropertiesPtrOutput() StackPropertiesPtrOutput {
	return i.ToStackPropertiesPtrOutputWithContext(context.Background())
}

func (i StackPropertiesArgs) ToStackPropertiesPtrOutputWithContext(ctx context.Context) StackPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackPropertiesOutput).ToStackPropertiesPtrOutputWithContext(ctx)
}









type StackPropertiesPtrInput interface {
	pulumi.Input

	ToStackPropertiesPtrOutput() StackPropertiesPtrOutput
	ToStackPropertiesPtrOutputWithContext(context.Context) StackPropertiesPtrOutput
}

type stackPropertiesPtrType StackPropertiesArgs

func StackPropertiesPtr(v *StackPropertiesArgs) StackPropertiesPtrInput {
	return (*stackPropertiesPtrType)(v)
}

func (*stackPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StackProperties)(nil)).Elem()
}

func (i *stackPropertiesPtrType) ToStackPropertiesPtrOutput() StackPropertiesPtrOutput {
	return i.ToStackPropertiesPtrOutputWithContext(context.Background())
}

func (i *stackPropertiesPtrType) ToStackPropertiesPtrOutputWithContext(ctx context.Context) StackPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackPropertiesPtrOutput)
}

type StackPropertiesOutput struct{ *pulumi.OutputState }

func (StackPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StackProperties)(nil)).Elem()
}

func (o StackPropertiesOutput) ToStackPropertiesOutput() StackPropertiesOutput {
	return o
}

func (o StackPropertiesOutput) ToStackPropertiesOutputWithContext(ctx context.Context) StackPropertiesOutput {
	return o
}

func (o StackPropertiesOutput) ToStackPropertiesPtrOutput() StackPropertiesPtrOutput {
	return o.ToStackPropertiesPtrOutputWithContext(context.Background())
}

func (o StackPropertiesOutput) ToStackPropertiesPtrOutputWithContext(ctx context.Context) StackPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StackProperties) *StackProperties {
		return &v
	}).(StackPropertiesPtrOutput)
}

func (o StackPropertiesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StackProperties) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o StackPropertiesOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StackProperties) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type StackPropertiesPtrOutput struct{ *pulumi.OutputState }

func (StackPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StackProperties)(nil)).Elem()
}

func (o StackPropertiesPtrOutput) ToStackPropertiesPtrOutput() StackPropertiesPtrOutput {
	return o
}

func (o StackPropertiesPtrOutput) ToStackPropertiesPtrOutputWithContext(ctx context.Context) StackPropertiesPtrOutput {
	return o
}

func (o StackPropertiesPtrOutput) Elem() StackPropertiesOutput {
	return o.ApplyT(func(v *StackProperties) StackProperties {
		if v != nil {
			return *v
		}
		var ret StackProperties
		return ret
	}).(StackPropertiesOutput)
}

func (o StackPropertiesPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StackProperties) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o StackPropertiesPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StackProperties) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type StackPropertiesResponse struct {
	Id      *string `pulumi:"id"`
	Version *string `pulumi:"version"`
}

type StackPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StackPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StackPropertiesResponse)(nil)).Elem()
}

func (o StackPropertiesResponseOutput) ToStackPropertiesResponseOutput() StackPropertiesResponseOutput {
	return o
}

func (o StackPropertiesResponseOutput) ToStackPropertiesResponseOutputWithContext(ctx context.Context) StackPropertiesResponseOutput {
	return o
}

func (o StackPropertiesResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StackPropertiesResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o StackPropertiesResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StackPropertiesResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type StackPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (StackPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StackPropertiesResponse)(nil)).Elem()
}

func (o StackPropertiesResponsePtrOutput) ToStackPropertiesResponsePtrOutput() StackPropertiesResponsePtrOutput {
	return o
}

func (o StackPropertiesResponsePtrOutput) ToStackPropertiesResponsePtrOutputWithContext(ctx context.Context) StackPropertiesResponsePtrOutput {
	return o
}

func (o StackPropertiesResponsePtrOutput) Elem() StackPropertiesResponseOutput {
	return o.ApplyT(func(v *StackPropertiesResponse) StackPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret StackPropertiesResponse
		return ret
	}).(StackPropertiesResponseOutput)
}

func (o StackPropertiesResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StackPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o StackPropertiesResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StackPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type StorageAccount struct {
	AccountKey  string `pulumi:"accountKey"`
	AccountName string `pulumi:"accountName"`
	StorageType string `pulumi:"storageType"`
}





type StorageAccountInput interface {
	pulumi.Input

	ToStorageAccountOutput() StorageAccountOutput
	ToStorageAccountOutputWithContext(context.Context) StorageAccountOutput
}

type StorageAccountArgs struct {
	AccountKey  pulumi.StringInput `pulumi:"accountKey"`
	AccountName pulumi.StringInput `pulumi:"accountName"`
	StorageType pulumi.StringInput `pulumi:"storageType"`
}

func (StorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (i StorageAccountArgs) ToStorageAccountOutput() StorageAccountOutput {
	return i.ToStorageAccountOutputWithContext(context.Background())
}

func (i StorageAccountArgs) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput)
}

func (i StorageAccountArgs) ToStorageAccountPtrOutput() StorageAccountPtrOutput {
	return i.ToStorageAccountPtrOutputWithContext(context.Background())
}

func (i StorageAccountArgs) ToStorageAccountPtrOutputWithContext(ctx context.Context) StorageAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput).ToStorageAccountPtrOutputWithContext(ctx)
}









type StorageAccountPtrInput interface {
	pulumi.Input

	ToStorageAccountPtrOutput() StorageAccountPtrOutput
	ToStorageAccountPtrOutputWithContext(context.Context) StorageAccountPtrOutput
}

type storageAccountPtrType StorageAccountArgs

func StorageAccountPtr(v *StorageAccountArgs) StorageAccountPtrInput {
	return (*storageAccountPtrType)(v)
}

func (*storageAccountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccount)(nil)).Elem()
}

func (i *storageAccountPtrType) ToStorageAccountPtrOutput() StorageAccountPtrOutput {
	return i.ToStorageAccountPtrOutputWithContext(context.Background())
}

func (i *storageAccountPtrType) ToStorageAccountPtrOutputWithContext(ctx context.Context) StorageAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPtrOutput)
}

type StorageAccountOutput struct{ *pulumi.OutputState }

func (StorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (o StorageAccountOutput) ToStorageAccountOutput() StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountPtrOutput() StorageAccountPtrOutput {
	return o.ToStorageAccountPtrOutputWithContext(context.Background())
}

func (o StorageAccountOutput) ToStorageAccountPtrOutputWithContext(ctx context.Context) StorageAccountPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccount) *StorageAccount {
		return &v
	}).(StorageAccountPtrOutput)
}

func (o StorageAccountOutput) AccountKey() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccount) string { return v.AccountKey }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccount) string { return v.AccountName }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccount) string { return v.StorageType }).(pulumi.StringOutput)
}

type StorageAccountPtrOutput struct{ *pulumi.OutputState }

func (StorageAccountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccount)(nil)).Elem()
}

func (o StorageAccountPtrOutput) ToStorageAccountPtrOutput() StorageAccountPtrOutput {
	return o
}

func (o StorageAccountPtrOutput) ToStorageAccountPtrOutputWithContext(ctx context.Context) StorageAccountPtrOutput {
	return o
}

func (o StorageAccountPtrOutput) Elem() StorageAccountOutput {
	return o.ApplyT(func(v *StorageAccount) StorageAccount {
		if v != nil {
			return *v
		}
		var ret StorageAccount
		return ret
	}).(StorageAccountOutput)
}

func (o StorageAccountPtrOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) *string {
		if v == nil {
			return nil
		}
		return &v.AccountKey
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountPtrOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) *string {
		if v == nil {
			return nil
		}
		return &v.AccountName
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountPtrOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) *string {
		if v == nil {
			return nil
		}
		return &v.StorageType
	}).(pulumi.StringPtrOutput)
}

type StorageAccountResponse struct {
	AccountName string `pulumi:"accountName"`
	StorageType string `pulumi:"storageType"`
}

type StorageAccountResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutput() StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutputWithContext(ctx context.Context) StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountResponse) string { return v.AccountName }).(pulumi.StringOutput)
}

func (o StorageAccountResponseOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountResponse) string { return v.StorageType }).(pulumi.StringOutput)
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

type TemporaryDisk struct {
	MountPath *string `pulumi:"mountPath"`
	SizeInGB  *int    `pulumi:"sizeInGB"`
}


func (val *TemporaryDisk) Defaults() *TemporaryDisk {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MountPath) {
		mountPath_ := "/tmp"
		tmp.MountPath = &mountPath_
	}
	return &tmp
}





type TemporaryDiskInput interface {
	pulumi.Input

	ToTemporaryDiskOutput() TemporaryDiskOutput
	ToTemporaryDiskOutputWithContext(context.Context) TemporaryDiskOutput
}

type TemporaryDiskArgs struct {
	MountPath pulumi.StringPtrInput `pulumi:"mountPath"`
	SizeInGB  pulumi.IntPtrInput    `pulumi:"sizeInGB"`
}


func (val *TemporaryDiskArgs) Defaults() *TemporaryDiskArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MountPath) {
		tmp.MountPath = pulumi.StringPtr("/tmp")
	}
	return &tmp
}
func (TemporaryDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TemporaryDisk)(nil)).Elem()
}

func (i TemporaryDiskArgs) ToTemporaryDiskOutput() TemporaryDiskOutput {
	return i.ToTemporaryDiskOutputWithContext(context.Background())
}

func (i TemporaryDiskArgs) ToTemporaryDiskOutputWithContext(ctx context.Context) TemporaryDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemporaryDiskOutput)
}

func (i TemporaryDiskArgs) ToTemporaryDiskPtrOutput() TemporaryDiskPtrOutput {
	return i.ToTemporaryDiskPtrOutputWithContext(context.Background())
}

func (i TemporaryDiskArgs) ToTemporaryDiskPtrOutputWithContext(ctx context.Context) TemporaryDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemporaryDiskOutput).ToTemporaryDiskPtrOutputWithContext(ctx)
}









type TemporaryDiskPtrInput interface {
	pulumi.Input

	ToTemporaryDiskPtrOutput() TemporaryDiskPtrOutput
	ToTemporaryDiskPtrOutputWithContext(context.Context) TemporaryDiskPtrOutput
}

type temporaryDiskPtrType TemporaryDiskArgs

func TemporaryDiskPtr(v *TemporaryDiskArgs) TemporaryDiskPtrInput {
	return (*temporaryDiskPtrType)(v)
}

func (*temporaryDiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TemporaryDisk)(nil)).Elem()
}

func (i *temporaryDiskPtrType) ToTemporaryDiskPtrOutput() TemporaryDiskPtrOutput {
	return i.ToTemporaryDiskPtrOutputWithContext(context.Background())
}

func (i *temporaryDiskPtrType) ToTemporaryDiskPtrOutputWithContext(ctx context.Context) TemporaryDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemporaryDiskPtrOutput)
}

type TemporaryDiskOutput struct{ *pulumi.OutputState }

func (TemporaryDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemporaryDisk)(nil)).Elem()
}

func (o TemporaryDiskOutput) ToTemporaryDiskOutput() TemporaryDiskOutput {
	return o
}

func (o TemporaryDiskOutput) ToTemporaryDiskOutputWithContext(ctx context.Context) TemporaryDiskOutput {
	return o
}

func (o TemporaryDiskOutput) ToTemporaryDiskPtrOutput() TemporaryDiskPtrOutput {
	return o.ToTemporaryDiskPtrOutputWithContext(context.Background())
}

func (o TemporaryDiskOutput) ToTemporaryDiskPtrOutputWithContext(ctx context.Context) TemporaryDiskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TemporaryDisk) *TemporaryDisk {
		return &v
	}).(TemporaryDiskPtrOutput)
}

func (o TemporaryDiskOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TemporaryDisk) *string { return v.MountPath }).(pulumi.StringPtrOutput)
}

func (o TemporaryDiskOutput) SizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TemporaryDisk) *int { return v.SizeInGB }).(pulumi.IntPtrOutput)
}

type TemporaryDiskPtrOutput struct{ *pulumi.OutputState }

func (TemporaryDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemporaryDisk)(nil)).Elem()
}

func (o TemporaryDiskPtrOutput) ToTemporaryDiskPtrOutput() TemporaryDiskPtrOutput {
	return o
}

func (o TemporaryDiskPtrOutput) ToTemporaryDiskPtrOutputWithContext(ctx context.Context) TemporaryDiskPtrOutput {
	return o
}

func (o TemporaryDiskPtrOutput) Elem() TemporaryDiskOutput {
	return o.ApplyT(func(v *TemporaryDisk) TemporaryDisk {
		if v != nil {
			return *v
		}
		var ret TemporaryDisk
		return ret
	}).(TemporaryDiskOutput)
}

func (o TemporaryDiskPtrOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemporaryDisk) *string {
		if v == nil {
			return nil
		}
		return v.MountPath
	}).(pulumi.StringPtrOutput)
}

func (o TemporaryDiskPtrOutput) SizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TemporaryDisk) *int {
		if v == nil {
			return nil
		}
		return v.SizeInGB
	}).(pulumi.IntPtrOutput)
}

type TemporaryDiskResponse struct {
	MountPath *string `pulumi:"mountPath"`
	SizeInGB  *int    `pulumi:"sizeInGB"`
}


func (val *TemporaryDiskResponse) Defaults() *TemporaryDiskResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MountPath) {
		mountPath_ := "/tmp"
		tmp.MountPath = &mountPath_
	}
	return &tmp
}

type TemporaryDiskResponseOutput struct{ *pulumi.OutputState }

func (TemporaryDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemporaryDiskResponse)(nil)).Elem()
}

func (o TemporaryDiskResponseOutput) ToTemporaryDiskResponseOutput() TemporaryDiskResponseOutput {
	return o
}

func (o TemporaryDiskResponseOutput) ToTemporaryDiskResponseOutputWithContext(ctx context.Context) TemporaryDiskResponseOutput {
	return o
}

func (o TemporaryDiskResponseOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TemporaryDiskResponse) *string { return v.MountPath }).(pulumi.StringPtrOutput)
}

func (o TemporaryDiskResponseOutput) SizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TemporaryDiskResponse) *int { return v.SizeInGB }).(pulumi.IntPtrOutput)
}

type TemporaryDiskResponsePtrOutput struct{ *pulumi.OutputState }

func (TemporaryDiskResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemporaryDiskResponse)(nil)).Elem()
}

func (o TemporaryDiskResponsePtrOutput) ToTemporaryDiskResponsePtrOutput() TemporaryDiskResponsePtrOutput {
	return o
}

func (o TemporaryDiskResponsePtrOutput) ToTemporaryDiskResponsePtrOutputWithContext(ctx context.Context) TemporaryDiskResponsePtrOutput {
	return o
}

func (o TemporaryDiskResponsePtrOutput) Elem() TemporaryDiskResponseOutput {
	return o.ApplyT(func(v *TemporaryDiskResponse) TemporaryDiskResponse {
		if v != nil {
			return *v
		}
		var ret TemporaryDiskResponse
		return ret
	}).(TemporaryDiskResponseOutput)
}

func (o TemporaryDiskResponsePtrOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemporaryDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.MountPath
	}).(pulumi.StringPtrOutput)
}

func (o TemporaryDiskResponsePtrOutput) SizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TemporaryDiskResponse) *int {
		if v == nil {
			return nil
		}
		return v.SizeInGB
	}).(pulumi.IntPtrOutput)
}

type UserSourceInfo struct {
	ArtifactSelector *string `pulumi:"artifactSelector"`
	RelativePath     *string `pulumi:"relativePath"`
	Type             *string `pulumi:"type"`
	Version          *string `pulumi:"version"`
}





type UserSourceInfoInput interface {
	pulumi.Input

	ToUserSourceInfoOutput() UserSourceInfoOutput
	ToUserSourceInfoOutputWithContext(context.Context) UserSourceInfoOutput
}

type UserSourceInfoArgs struct {
	ArtifactSelector pulumi.StringPtrInput `pulumi:"artifactSelector"`
	RelativePath     pulumi.StringPtrInput `pulumi:"relativePath"`
	Type             pulumi.StringPtrInput `pulumi:"type"`
	Version          pulumi.StringPtrInput `pulumi:"version"`
}

func (UserSourceInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserSourceInfo)(nil)).Elem()
}

func (i UserSourceInfoArgs) ToUserSourceInfoOutput() UserSourceInfoOutput {
	return i.ToUserSourceInfoOutputWithContext(context.Background())
}

func (i UserSourceInfoArgs) ToUserSourceInfoOutputWithContext(ctx context.Context) UserSourceInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSourceInfoOutput)
}

func (i UserSourceInfoArgs) ToUserSourceInfoPtrOutput() UserSourceInfoPtrOutput {
	return i.ToUserSourceInfoPtrOutputWithContext(context.Background())
}

func (i UserSourceInfoArgs) ToUserSourceInfoPtrOutputWithContext(ctx context.Context) UserSourceInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSourceInfoOutput).ToUserSourceInfoPtrOutputWithContext(ctx)
}









type UserSourceInfoPtrInput interface {
	pulumi.Input

	ToUserSourceInfoPtrOutput() UserSourceInfoPtrOutput
	ToUserSourceInfoPtrOutputWithContext(context.Context) UserSourceInfoPtrOutput
}

type userSourceInfoPtrType UserSourceInfoArgs

func UserSourceInfoPtr(v *UserSourceInfoArgs) UserSourceInfoPtrInput {
	return (*userSourceInfoPtrType)(v)
}

func (*userSourceInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSourceInfo)(nil)).Elem()
}

func (i *userSourceInfoPtrType) ToUserSourceInfoPtrOutput() UserSourceInfoPtrOutput {
	return i.ToUserSourceInfoPtrOutputWithContext(context.Background())
}

func (i *userSourceInfoPtrType) ToUserSourceInfoPtrOutputWithContext(ctx context.Context) UserSourceInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSourceInfoPtrOutput)
}

type UserSourceInfoOutput struct{ *pulumi.OutputState }

func (UserSourceInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserSourceInfo)(nil)).Elem()
}

func (o UserSourceInfoOutput) ToUserSourceInfoOutput() UserSourceInfoOutput {
	return o
}

func (o UserSourceInfoOutput) ToUserSourceInfoOutputWithContext(ctx context.Context) UserSourceInfoOutput {
	return o
}

func (o UserSourceInfoOutput) ToUserSourceInfoPtrOutput() UserSourceInfoPtrOutput {
	return o.ToUserSourceInfoPtrOutputWithContext(context.Background())
}

func (o UserSourceInfoOutput) ToUserSourceInfoPtrOutputWithContext(ctx context.Context) UserSourceInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserSourceInfo) *UserSourceInfo {
		return &v
	}).(UserSourceInfoPtrOutput)
}

func (o UserSourceInfoOutput) ArtifactSelector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserSourceInfo) *string { return v.ArtifactSelector }).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserSourceInfo) *string { return v.RelativePath }).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserSourceInfo) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserSourceInfo) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type UserSourceInfoPtrOutput struct{ *pulumi.OutputState }

func (UserSourceInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSourceInfo)(nil)).Elem()
}

func (o UserSourceInfoPtrOutput) ToUserSourceInfoPtrOutput() UserSourceInfoPtrOutput {
	return o
}

func (o UserSourceInfoPtrOutput) ToUserSourceInfoPtrOutputWithContext(ctx context.Context) UserSourceInfoPtrOutput {
	return o
}

func (o UserSourceInfoPtrOutput) Elem() UserSourceInfoOutput {
	return o.ApplyT(func(v *UserSourceInfo) UserSourceInfo {
		if v != nil {
			return *v
		}
		var ret UserSourceInfo
		return ret
	}).(UserSourceInfoOutput)
}

func (o UserSourceInfoPtrOutput) ArtifactSelector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSourceInfo) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactSelector
	}).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoPtrOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSourceInfo) *string {
		if v == nil {
			return nil
		}
		return v.RelativePath
	}).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSourceInfo) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSourceInfo) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type UserSourceInfoResponse struct {
	ArtifactSelector *string `pulumi:"artifactSelector"`
	RelativePath     *string `pulumi:"relativePath"`
	Type             *string `pulumi:"type"`
	Version          *string `pulumi:"version"`
}

type UserSourceInfoResponseOutput struct{ *pulumi.OutputState }

func (UserSourceInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserSourceInfoResponse)(nil)).Elem()
}

func (o UserSourceInfoResponseOutput) ToUserSourceInfoResponseOutput() UserSourceInfoResponseOutput {
	return o
}

func (o UserSourceInfoResponseOutput) ToUserSourceInfoResponseOutputWithContext(ctx context.Context) UserSourceInfoResponseOutput {
	return o
}

func (o UserSourceInfoResponseOutput) ArtifactSelector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserSourceInfoResponse) *string { return v.ArtifactSelector }).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoResponseOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserSourceInfoResponse) *string { return v.RelativePath }).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserSourceInfoResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserSourceInfoResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type UserSourceInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (UserSourceInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSourceInfoResponse)(nil)).Elem()
}

func (o UserSourceInfoResponsePtrOutput) ToUserSourceInfoResponsePtrOutput() UserSourceInfoResponsePtrOutput {
	return o
}

func (o UserSourceInfoResponsePtrOutput) ToUserSourceInfoResponsePtrOutputWithContext(ctx context.Context) UserSourceInfoResponsePtrOutput {
	return o
}

func (o UserSourceInfoResponsePtrOutput) Elem() UserSourceInfoResponseOutput {
	return o.ApplyT(func(v *UserSourceInfoResponse) UserSourceInfoResponse {
		if v != nil {
			return *v
		}
		var ret UserSourceInfoResponse
		return ret
	}).(UserSourceInfoResponseOutput)
}

func (o UserSourceInfoResponsePtrOutput) ArtifactSelector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSourceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactSelector
	}).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoResponsePtrOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSourceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.RelativePath
	}).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSourceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSourceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiPortalCustomDomainPropertiesOutput{})
	pulumi.RegisterOutputType(ApiPortalCustomDomainPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ApiPortalCustomDomainPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ApiPortalInstanceResponseOutput{})
	pulumi.RegisterOutputType(ApiPortalInstanceResponseArrayOutput{})
	pulumi.RegisterOutputType(ApiPortalPropertiesOutput{})
	pulumi.RegisterOutputType(ApiPortalPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ApiPortalPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ApiPortalResourceRequestsResponseOutput{})
	pulumi.RegisterOutputType(AppResourcePropertiesOutput{})
	pulumi.RegisterOutputType(AppResourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AppResourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(BindingResourcePropertiesOutput{})
	pulumi.RegisterOutputType(BindingResourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(BindingResourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(BuilderPropertiesOutput{})
	pulumi.RegisterOutputType(BuilderPropertiesPtrOutput{})
	pulumi.RegisterOutputType(BuilderPropertiesResponseOutput{})
	pulumi.RegisterOutputType(BuildpackBindingLaunchPropertiesOutput{})
	pulumi.RegisterOutputType(BuildpackBindingLaunchPropertiesPtrOutput{})
	pulumi.RegisterOutputType(BuildpackBindingLaunchPropertiesResponseOutput{})
	pulumi.RegisterOutputType(BuildpackBindingLaunchPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(BuildpackBindingPropertiesOutput{})
	pulumi.RegisterOutputType(BuildpackBindingPropertiesPtrOutput{})
	pulumi.RegisterOutputType(BuildpackBindingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(BuildpackPropertiesOutput{})
	pulumi.RegisterOutputType(BuildpackPropertiesArrayOutput{})
	pulumi.RegisterOutputType(BuildpackPropertiesResponseOutput{})
	pulumi.RegisterOutputType(BuildpackPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(BuildpacksGroupPropertiesOutput{})
	pulumi.RegisterOutputType(BuildpacksGroupPropertiesArrayOutput{})
	pulumi.RegisterOutputType(BuildpacksGroupPropertiesResponseOutput{})
	pulumi.RegisterOutputType(BuildpacksGroupPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(CertificatePropertiesOutput{})
	pulumi.RegisterOutputType(CertificatePropertiesPtrOutput{})
	pulumi.RegisterOutputType(CertificatePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ClusterResourcePropertiesOutput{})
	pulumi.RegisterOutputType(ClusterResourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ClusterResourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationServiceGitPropertyOutput{})
	pulumi.RegisterOutputType(ConfigurationServiceGitPropertyPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationServiceGitPropertyResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationServiceGitPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(ConfigurationServiceGitRepositoryOutput{})
	pulumi.RegisterOutputType(ConfigurationServiceGitRepositoryArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationServiceGitRepositoryResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationServiceGitRepositoryResponseArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationServiceInstanceResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationServiceInstanceResponseArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationServicePropertiesOutput{})
	pulumi.RegisterOutputType(ConfigurationServicePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationServicePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationServiceResourceRequestsResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationServiceSettingsOutput{})
	pulumi.RegisterOutputType(ConfigurationServiceSettingsPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationServiceSettingsResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationServiceSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(CustomDomainPropertiesOutput{})
	pulumi.RegisterOutputType(CustomDomainPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CustomDomainPropertiesResponseOutput{})
	pulumi.RegisterOutputType(DeploymentInstanceResponseOutput{})
	pulumi.RegisterOutputType(DeploymentInstanceResponseArrayOutput{})
	pulumi.RegisterOutputType(DeploymentResourcePropertiesOutput{})
	pulumi.RegisterOutputType(DeploymentResourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(DeploymentResourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(DeploymentSettingsOutput{})
	pulumi.RegisterOutputType(DeploymentSettingsPtrOutput{})
	pulumi.RegisterOutputType(DeploymentSettingsResponseOutput{})
	pulumi.RegisterOutputType(DeploymentSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(GatewayApiMetadataPropertiesOutput{})
	pulumi.RegisterOutputType(GatewayApiMetadataPropertiesPtrOutput{})
	pulumi.RegisterOutputType(GatewayApiMetadataPropertiesResponseOutput{})
	pulumi.RegisterOutputType(GatewayApiMetadataPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(GatewayApiRouteOutput{})
	pulumi.RegisterOutputType(GatewayApiRouteArrayOutput{})
	pulumi.RegisterOutputType(GatewayApiRouteResponseOutput{})
	pulumi.RegisterOutputType(GatewayApiRouteResponseArrayOutput{})
	pulumi.RegisterOutputType(GatewayCorsPropertiesOutput{})
	pulumi.RegisterOutputType(GatewayCorsPropertiesPtrOutput{})
	pulumi.RegisterOutputType(GatewayCorsPropertiesResponseOutput{})
	pulumi.RegisterOutputType(GatewayCorsPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(GatewayCustomDomainPropertiesOutput{})
	pulumi.RegisterOutputType(GatewayCustomDomainPropertiesPtrOutput{})
	pulumi.RegisterOutputType(GatewayCustomDomainPropertiesResponseOutput{})
	pulumi.RegisterOutputType(GatewayInstanceResponseOutput{})
	pulumi.RegisterOutputType(GatewayInstanceResponseArrayOutput{})
	pulumi.RegisterOutputType(GatewayOperatorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(GatewayOperatorResourceRequestsResponseOutput{})
	pulumi.RegisterOutputType(GatewayPropertiesOutput{})
	pulumi.RegisterOutputType(GatewayPropertiesPtrOutput{})
	pulumi.RegisterOutputType(GatewayPropertiesResponseOutput{})
	pulumi.RegisterOutputType(GatewayResourceRequestsOutput{})
	pulumi.RegisterOutputType(GatewayResourceRequestsPtrOutput{})
	pulumi.RegisterOutputType(GatewayResourceRequestsResponseOutput{})
	pulumi.RegisterOutputType(GatewayResourceRequestsResponsePtrOutput{})
	pulumi.RegisterOutputType(GatewayRouteConfigPropertiesOutput{})
	pulumi.RegisterOutputType(GatewayRouteConfigPropertiesPtrOutput{})
	pulumi.RegisterOutputType(GatewayRouteConfigPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ManagedIdentityPropertiesOutput{})
	pulumi.RegisterOutputType(ManagedIdentityPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ManagedIdentityPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkProfileOutput{})
	pulumi.RegisterOutputType(NetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponseOutboundIPsOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponseOutboundIPsPtrOutput{})
	pulumi.RegisterOutputType(PersistentDiskOutput{})
	pulumi.RegisterOutputType(PersistentDiskPtrOutput{})
	pulumi.RegisterOutputType(PersistentDiskResponseOutput{})
	pulumi.RegisterOutputType(PersistentDiskResponsePtrOutput{})
	pulumi.RegisterOutputType(RequiredTrafficResponseOutput{})
	pulumi.RegisterOutputType(RequiredTrafficResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceRegistryInstanceResponseOutput{})
	pulumi.RegisterOutputType(ServiceRegistryInstanceResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceRegistryPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ServiceRegistryResourceRequestsResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SsoPropertiesOutput{})
	pulumi.RegisterOutputType(SsoPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SsoPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SsoPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(StackPropertiesOutput{})
	pulumi.RegisterOutputType(StackPropertiesPtrOutput{})
	pulumi.RegisterOutputType(StackPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StackPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageAccountOutput{})
	pulumi.RegisterOutputType(StorageAccountPtrOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TemporaryDiskOutput{})
	pulumi.RegisterOutputType(TemporaryDiskPtrOutput{})
	pulumi.RegisterOutputType(TemporaryDiskResponseOutput{})
	pulumi.RegisterOutputType(TemporaryDiskResponsePtrOutput{})
	pulumi.RegisterOutputType(UserSourceInfoOutput{})
	pulumi.RegisterOutputType(UserSourceInfoPtrOutput{})
	pulumi.RegisterOutputType(UserSourceInfoResponseOutput{})
	pulumi.RegisterOutputType(UserSourceInfoResponsePtrOutput{})
}




package v20180820preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceIdentity struct {
	Type *string `pulumi:"type"`
}





type ResourceIdentityInput interface {
	pulumi.Input

	ToResourceIdentityOutput() ResourceIdentityOutput
	ToResourceIdentityOutputWithContext(context.Context) ResourceIdentityOutput
}

type ResourceIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (i ResourceIdentityArgs) ToResourceIdentityOutput() ResourceIdentityOutput {
	return i.ToResourceIdentityOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput)
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput).ToResourceIdentityPtrOutputWithContext(ctx)
}









type ResourceIdentityPtrInput interface {
	pulumi.Input

	ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput
	ToResourceIdentityPtrOutputWithContext(context.Context) ResourceIdentityPtrOutput
}

type resourceIdentityPtrType ResourceIdentityArgs

func ResourceIdentityPtr(v *ResourceIdentityArgs) ResourceIdentityPtrInput {
	return (*resourceIdentityPtrType)(v)
}

func (*resourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityPtrOutput)
}

type ResourceIdentityOutput struct{ *pulumi.OutputState }

func (ResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityOutput) ToResourceIdentityOutput() ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentity) *ResourceIdentity {
		return &v
	}).(ResourceIdentityPtrOutput)
}

func (o ResourceIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) Elem() ResourceIdentityOutput {
	return o.ApplyT(func(v *ResourceIdentity) ResourceIdentity {
		if v != nil {
			return *v
		}
		var ret ResourceIdentity
		return ret
	}).(ResourceIdentityOutput)
}

func (o ResourceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ResourceResponseIdentity struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type ResourceResponseIdentityOutput struct{ *pulumi.OutputState }

func (ResourceResponseIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceResponseIdentity)(nil)).Elem()
}

func (o ResourceResponseIdentityOutput) ToResourceResponseIdentityOutput() ResourceResponseIdentityOutput {
	return o
}

func (o ResourceResponseIdentityOutput) ToResourceResponseIdentityOutputWithContext(ctx context.Context) ResourceResponseIdentityOutput {
	return o
}

func (o ResourceResponseIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceResponseIdentity) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ResourceResponseIdentityOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceResponseIdentity) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ResourceResponseIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceResponseIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ResourceResponseIdentityPtrOutput struct{ *pulumi.OutputState }

func (ResourceResponseIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceResponseIdentity)(nil)).Elem()
}

func (o ResourceResponseIdentityPtrOutput) ToResourceResponseIdentityPtrOutput() ResourceResponseIdentityPtrOutput {
	return o
}

func (o ResourceResponseIdentityPtrOutput) ToResourceResponseIdentityPtrOutputWithContext(ctx context.Context) ResourceResponseIdentityPtrOutput {
	return o
}

func (o ResourceResponseIdentityPtrOutput) Elem() ResourceResponseIdentityOutput {
	return o.ApplyT(func(v *ResourceResponseIdentity) ResourceResponseIdentity {
		if v != nil {
			return *v
		}
		var ret ResourceResponseIdentity
		return ret
	}).(ResourceResponseIdentityOutput)
}

func (o ResourceResponseIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceResponseIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceResponseIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
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
	OfferThroughput *int `pulumi:"offerThroughput"`
}





type ServiceCosmosDbConfigurationInfoInput interface {
	pulumi.Input

	ToServiceCosmosDbConfigurationInfoOutput() ServiceCosmosDbConfigurationInfoOutput
	ToServiceCosmosDbConfigurationInfoOutputWithContext(context.Context) ServiceCosmosDbConfigurationInfoOutput
}

type ServiceCosmosDbConfigurationInfoArgs struct {
	OfferThroughput pulumi.IntPtrInput `pulumi:"offerThroughput"`
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

func (o ServiceCosmosDbConfigurationInfoPtrOutput) OfferThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceCosmosDbConfigurationInfo) *int {
		if v == nil {
			return nil
		}
		return v.OfferThroughput
	}).(pulumi.IntPtrOutput)
}

type ServiceCosmosDbConfigurationInfoResponse struct {
	OfferThroughput *int `pulumi:"offerThroughput"`
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

func (o ServiceCosmosDbConfigurationInfoResponsePtrOutput) OfferThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceCosmosDbConfigurationInfoResponse) *int {
		if v == nil {
			return nil
		}
		return v.OfferThroughput
	}).(pulumi.IntPtrOutput)
}

type ServicesProperties struct {
	AccessPolicies              []ServiceAccessPolicyEntry              `pulumi:"accessPolicies"`
	AuthenticationConfiguration *ServiceAuthenticationConfigurationInfo `pulumi:"authenticationConfiguration"`
	CorsConfiguration           *ServiceCorsConfigurationInfo           `pulumi:"corsConfiguration"`
	CosmosDbConfiguration       *ServiceCosmosDbConfigurationInfo       `pulumi:"cosmosDbConfiguration"`
}





type ServicesPropertiesInput interface {
	pulumi.Input

	ToServicesPropertiesOutput() ServicesPropertiesOutput
	ToServicesPropertiesOutputWithContext(context.Context) ServicesPropertiesOutput
}

type ServicesPropertiesArgs struct {
	AccessPolicies              ServiceAccessPolicyEntryArrayInput             `pulumi:"accessPolicies"`
	AuthenticationConfiguration ServiceAuthenticationConfigurationInfoPtrInput `pulumi:"authenticationConfiguration"`
	CorsConfiguration           ServiceCorsConfigurationInfoPtrInput           `pulumi:"corsConfiguration"`
	CosmosDbConfiguration       ServiceCosmosDbConfigurationInfoPtrInput       `pulumi:"cosmosDbConfiguration"`
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

type ServicesPropertiesResponse struct {
	AccessPolicies              []ServiceAccessPolicyEntryResponse              `pulumi:"accessPolicies"`
	AuthenticationConfiguration *ServiceAuthenticationConfigurationInfoResponse `pulumi:"authenticationConfiguration"`
	CorsConfiguration           *ServiceCorsConfigurationInfoResponse           `pulumi:"corsConfiguration"`
	CosmosDbConfiguration       *ServiceCosmosDbConfigurationInfoResponse       `pulumi:"cosmosDbConfiguration"`
	ProvisioningState           string                                          `pulumi:"provisioningState"`
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

func (o ServicesPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceIdentityOutput{})
	pulumi.RegisterOutputType(ResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ResourceResponseIdentityOutput{})
	pulumi.RegisterOutputType(ResourceResponseIdentityPtrOutput{})
	pulumi.RegisterOutputType(ServiceAccessPolicyEntryOutput{})
	pulumi.RegisterOutputType(ServiceAccessPolicyEntryArrayOutput{})
	pulumi.RegisterOutputType(ServiceAccessPolicyEntryResponseOutput{})
	pulumi.RegisterOutputType(ServiceAccessPolicyEntryResponseArrayOutput{})
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
	pulumi.RegisterOutputType(ServicesPropertiesOutput{})
	pulumi.RegisterOutputType(ServicesPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ServicesPropertiesResponseOutput{})
}

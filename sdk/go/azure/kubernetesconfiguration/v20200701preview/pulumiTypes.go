


package v20200701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComplianceStatusResponse struct {
	ComplianceState   string  `pulumi:"complianceState"`
	LastConfigApplied *string `pulumi:"lastConfigApplied"`
	Message           *string `pulumi:"message"`
	MessageLevel      *string `pulumi:"messageLevel"`
}





type ComplianceStatusResponseInput interface {
	pulumi.Input

	ToComplianceStatusResponseOutput() ComplianceStatusResponseOutput
	ToComplianceStatusResponseOutputWithContext(context.Context) ComplianceStatusResponseOutput
}

type ComplianceStatusResponseArgs struct {
	ComplianceState   pulumi.StringInput    `pulumi:"complianceState"`
	LastConfigApplied pulumi.StringPtrInput `pulumi:"lastConfigApplied"`
	Message           pulumi.StringPtrInput `pulumi:"message"`
	MessageLevel      pulumi.StringPtrInput `pulumi:"messageLevel"`
}

func (ComplianceStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComplianceStatusResponse)(nil)).Elem()
}

func (i ComplianceStatusResponseArgs) ToComplianceStatusResponseOutput() ComplianceStatusResponseOutput {
	return i.ToComplianceStatusResponseOutputWithContext(context.Background())
}

func (i ComplianceStatusResponseArgs) ToComplianceStatusResponseOutputWithContext(ctx context.Context) ComplianceStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComplianceStatusResponseOutput)
}

func (i ComplianceStatusResponseArgs) ToComplianceStatusResponsePtrOutput() ComplianceStatusResponsePtrOutput {
	return i.ToComplianceStatusResponsePtrOutputWithContext(context.Background())
}

func (i ComplianceStatusResponseArgs) ToComplianceStatusResponsePtrOutputWithContext(ctx context.Context) ComplianceStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComplianceStatusResponseOutput).ToComplianceStatusResponsePtrOutputWithContext(ctx)
}









type ComplianceStatusResponsePtrInput interface {
	pulumi.Input

	ToComplianceStatusResponsePtrOutput() ComplianceStatusResponsePtrOutput
	ToComplianceStatusResponsePtrOutputWithContext(context.Context) ComplianceStatusResponsePtrOutput
}

type complianceStatusResponsePtrType ComplianceStatusResponseArgs

func ComplianceStatusResponsePtr(v *ComplianceStatusResponseArgs) ComplianceStatusResponsePtrInput {
	return (*complianceStatusResponsePtrType)(v)
}

func (*complianceStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ComplianceStatusResponse)(nil)).Elem()
}

func (i *complianceStatusResponsePtrType) ToComplianceStatusResponsePtrOutput() ComplianceStatusResponsePtrOutput {
	return i.ToComplianceStatusResponsePtrOutputWithContext(context.Background())
}

func (i *complianceStatusResponsePtrType) ToComplianceStatusResponsePtrOutputWithContext(ctx context.Context) ComplianceStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComplianceStatusResponsePtrOutput)
}

type ComplianceStatusResponseOutput struct{ *pulumi.OutputState }

func (ComplianceStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComplianceStatusResponse)(nil)).Elem()
}

func (o ComplianceStatusResponseOutput) ToComplianceStatusResponseOutput() ComplianceStatusResponseOutput {
	return o
}

func (o ComplianceStatusResponseOutput) ToComplianceStatusResponseOutputWithContext(ctx context.Context) ComplianceStatusResponseOutput {
	return o
}

func (o ComplianceStatusResponseOutput) ToComplianceStatusResponsePtrOutput() ComplianceStatusResponsePtrOutput {
	return o.ToComplianceStatusResponsePtrOutputWithContext(context.Background())
}

func (o ComplianceStatusResponseOutput) ToComplianceStatusResponsePtrOutputWithContext(ctx context.Context) ComplianceStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComplianceStatusResponse) *ComplianceStatusResponse {
		return &v
	}).(ComplianceStatusResponsePtrOutput)
}

func (o ComplianceStatusResponseOutput) ComplianceState() pulumi.StringOutput {
	return o.ApplyT(func(v ComplianceStatusResponse) string { return v.ComplianceState }).(pulumi.StringOutput)
}

func (o ComplianceStatusResponseOutput) LastConfigApplied() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComplianceStatusResponse) *string { return v.LastConfigApplied }).(pulumi.StringPtrOutput)
}

func (o ComplianceStatusResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComplianceStatusResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ComplianceStatusResponseOutput) MessageLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComplianceStatusResponse) *string { return v.MessageLevel }).(pulumi.StringPtrOutput)
}

type ComplianceStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (ComplianceStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComplianceStatusResponse)(nil)).Elem()
}

func (o ComplianceStatusResponsePtrOutput) ToComplianceStatusResponsePtrOutput() ComplianceStatusResponsePtrOutput {
	return o
}

func (o ComplianceStatusResponsePtrOutput) ToComplianceStatusResponsePtrOutputWithContext(ctx context.Context) ComplianceStatusResponsePtrOutput {
	return o
}

func (o ComplianceStatusResponsePtrOutput) Elem() ComplianceStatusResponseOutput {
	return o.ApplyT(func(v *ComplianceStatusResponse) ComplianceStatusResponse {
		if v != nil {
			return *v
		}
		var ret ComplianceStatusResponse
		return ret
	}).(ComplianceStatusResponseOutput)
}

func (o ComplianceStatusResponsePtrOutput) ComplianceState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComplianceStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ComplianceState
	}).(pulumi.StringPtrOutput)
}

func (o ComplianceStatusResponsePtrOutput) LastConfigApplied() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComplianceStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastConfigApplied
	}).(pulumi.StringPtrOutput)
}

func (o ComplianceStatusResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComplianceStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ComplianceStatusResponsePtrOutput) MessageLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComplianceStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.MessageLevel
	}).(pulumi.StringPtrOutput)
}

type ConfigurationIdentity struct {
	Type *ResourceIdentityType `pulumi:"type"`
}





type ConfigurationIdentityInput interface {
	pulumi.Input

	ToConfigurationIdentityOutput() ConfigurationIdentityOutput
	ToConfigurationIdentityOutputWithContext(context.Context) ConfigurationIdentityOutput
}

type ConfigurationIdentityArgs struct {
	Type ResourceIdentityTypePtrInput `pulumi:"type"`
}

func (ConfigurationIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationIdentity)(nil)).Elem()
}

func (i ConfigurationIdentityArgs) ToConfigurationIdentityOutput() ConfigurationIdentityOutput {
	return i.ToConfigurationIdentityOutputWithContext(context.Background())
}

func (i ConfigurationIdentityArgs) ToConfigurationIdentityOutputWithContext(ctx context.Context) ConfigurationIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationIdentityOutput)
}

func (i ConfigurationIdentityArgs) ToConfigurationIdentityPtrOutput() ConfigurationIdentityPtrOutput {
	return i.ToConfigurationIdentityPtrOutputWithContext(context.Background())
}

func (i ConfigurationIdentityArgs) ToConfigurationIdentityPtrOutputWithContext(ctx context.Context) ConfigurationIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationIdentityOutput).ToConfigurationIdentityPtrOutputWithContext(ctx)
}









type ConfigurationIdentityPtrInput interface {
	pulumi.Input

	ToConfigurationIdentityPtrOutput() ConfigurationIdentityPtrOutput
	ToConfigurationIdentityPtrOutputWithContext(context.Context) ConfigurationIdentityPtrOutput
}

type configurationIdentityPtrType ConfigurationIdentityArgs

func ConfigurationIdentityPtr(v *ConfigurationIdentityArgs) ConfigurationIdentityPtrInput {
	return (*configurationIdentityPtrType)(v)
}

func (*configurationIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationIdentity)(nil)).Elem()
}

func (i *configurationIdentityPtrType) ToConfigurationIdentityPtrOutput() ConfigurationIdentityPtrOutput {
	return i.ToConfigurationIdentityPtrOutputWithContext(context.Background())
}

func (i *configurationIdentityPtrType) ToConfigurationIdentityPtrOutputWithContext(ctx context.Context) ConfigurationIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationIdentityPtrOutput)
}

type ConfigurationIdentityOutput struct{ *pulumi.OutputState }

func (ConfigurationIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationIdentity)(nil)).Elem()
}

func (o ConfigurationIdentityOutput) ToConfigurationIdentityOutput() ConfigurationIdentityOutput {
	return o
}

func (o ConfigurationIdentityOutput) ToConfigurationIdentityOutputWithContext(ctx context.Context) ConfigurationIdentityOutput {
	return o
}

func (o ConfigurationIdentityOutput) ToConfigurationIdentityPtrOutput() ConfigurationIdentityPtrOutput {
	return o.ToConfigurationIdentityPtrOutputWithContext(context.Background())
}

func (o ConfigurationIdentityOutput) ToConfigurationIdentityPtrOutputWithContext(ctx context.Context) ConfigurationIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationIdentity) *ConfigurationIdentity {
		return &v
	}).(ConfigurationIdentityPtrOutput)
}

func (o ConfigurationIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v ConfigurationIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

type ConfigurationIdentityPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationIdentity)(nil)).Elem()
}

func (o ConfigurationIdentityPtrOutput) ToConfigurationIdentityPtrOutput() ConfigurationIdentityPtrOutput {
	return o
}

func (o ConfigurationIdentityPtrOutput) ToConfigurationIdentityPtrOutputWithContext(ctx context.Context) ConfigurationIdentityPtrOutput {
	return o
}

func (o ConfigurationIdentityPtrOutput) Elem() ConfigurationIdentityOutput {
	return o.ApplyT(func(v *ConfigurationIdentity) ConfigurationIdentity {
		if v != nil {
			return *v
		}
		var ret ConfigurationIdentity
		return ret
	}).(ConfigurationIdentityOutput)
}

func (o ConfigurationIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *ConfigurationIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type ConfigurationIdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type ConfigurationIdentityResponseInput interface {
	pulumi.Input

	ToConfigurationIdentityResponseOutput() ConfigurationIdentityResponseOutput
	ToConfigurationIdentityResponseOutputWithContext(context.Context) ConfigurationIdentityResponseOutput
}

type ConfigurationIdentityResponseArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (ConfigurationIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationIdentityResponse)(nil)).Elem()
}

func (i ConfigurationIdentityResponseArgs) ToConfigurationIdentityResponseOutput() ConfigurationIdentityResponseOutput {
	return i.ToConfigurationIdentityResponseOutputWithContext(context.Background())
}

func (i ConfigurationIdentityResponseArgs) ToConfigurationIdentityResponseOutputWithContext(ctx context.Context) ConfigurationIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationIdentityResponseOutput)
}

func (i ConfigurationIdentityResponseArgs) ToConfigurationIdentityResponsePtrOutput() ConfigurationIdentityResponsePtrOutput {
	return i.ToConfigurationIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ConfigurationIdentityResponseArgs) ToConfigurationIdentityResponsePtrOutputWithContext(ctx context.Context) ConfigurationIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationIdentityResponseOutput).ToConfigurationIdentityResponsePtrOutputWithContext(ctx)
}









type ConfigurationIdentityResponsePtrInput interface {
	pulumi.Input

	ToConfigurationIdentityResponsePtrOutput() ConfigurationIdentityResponsePtrOutput
	ToConfigurationIdentityResponsePtrOutputWithContext(context.Context) ConfigurationIdentityResponsePtrOutput
}

type configurationIdentityResponsePtrType ConfigurationIdentityResponseArgs

func ConfigurationIdentityResponsePtr(v *ConfigurationIdentityResponseArgs) ConfigurationIdentityResponsePtrInput {
	return (*configurationIdentityResponsePtrType)(v)
}

func (*configurationIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationIdentityResponse)(nil)).Elem()
}

func (i *configurationIdentityResponsePtrType) ToConfigurationIdentityResponsePtrOutput() ConfigurationIdentityResponsePtrOutput {
	return i.ToConfigurationIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *configurationIdentityResponsePtrType) ToConfigurationIdentityResponsePtrOutputWithContext(ctx context.Context) ConfigurationIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationIdentityResponsePtrOutput)
}

type ConfigurationIdentityResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationIdentityResponse)(nil)).Elem()
}

func (o ConfigurationIdentityResponseOutput) ToConfigurationIdentityResponseOutput() ConfigurationIdentityResponseOutput {
	return o
}

func (o ConfigurationIdentityResponseOutput) ToConfigurationIdentityResponseOutputWithContext(ctx context.Context) ConfigurationIdentityResponseOutput {
	return o
}

func (o ConfigurationIdentityResponseOutput) ToConfigurationIdentityResponsePtrOutput() ConfigurationIdentityResponsePtrOutput {
	return o.ToConfigurationIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ConfigurationIdentityResponseOutput) ToConfigurationIdentityResponsePtrOutputWithContext(ctx context.Context) ConfigurationIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationIdentityResponse) *ConfigurationIdentityResponse {
		return &v
	}).(ConfigurationIdentityResponsePtrOutput)
}

func (o ConfigurationIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ConfigurationIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ConfigurationIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ConfigurationIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationIdentityResponse)(nil)).Elem()
}

func (o ConfigurationIdentityResponsePtrOutput) ToConfigurationIdentityResponsePtrOutput() ConfigurationIdentityResponsePtrOutput {
	return o
}

func (o ConfigurationIdentityResponsePtrOutput) ToConfigurationIdentityResponsePtrOutputWithContext(ctx context.Context) ConfigurationIdentityResponsePtrOutput {
	return o
}

func (o ConfigurationIdentityResponsePtrOutput) Elem() ConfigurationIdentityResponseOutput {
	return o.ApplyT(func(v *ConfigurationIdentityResponse) ConfigurationIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ConfigurationIdentityResponse
		return ret
	}).(ConfigurationIdentityResponseOutput)
}

func (o ConfigurationIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ErrorDefinitionResponse struct {
	Code    string `pulumi:"code"`
	Message string `pulumi:"message"`
}





type ErrorDefinitionResponseInput interface {
	pulumi.Input

	ToErrorDefinitionResponseOutput() ErrorDefinitionResponseOutput
	ToErrorDefinitionResponseOutputWithContext(context.Context) ErrorDefinitionResponseOutput
}

type ErrorDefinitionResponseArgs struct {
	Code    pulumi.StringInput `pulumi:"code"`
	Message pulumi.StringInput `pulumi:"message"`
}

func (ErrorDefinitionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDefinitionResponse)(nil)).Elem()
}

func (i ErrorDefinitionResponseArgs) ToErrorDefinitionResponseOutput() ErrorDefinitionResponseOutput {
	return i.ToErrorDefinitionResponseOutputWithContext(context.Background())
}

func (i ErrorDefinitionResponseArgs) ToErrorDefinitionResponseOutputWithContext(ctx context.Context) ErrorDefinitionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDefinitionResponseOutput)
}

func (i ErrorDefinitionResponseArgs) ToErrorDefinitionResponsePtrOutput() ErrorDefinitionResponsePtrOutput {
	return i.ToErrorDefinitionResponsePtrOutputWithContext(context.Background())
}

func (i ErrorDefinitionResponseArgs) ToErrorDefinitionResponsePtrOutputWithContext(ctx context.Context) ErrorDefinitionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDefinitionResponseOutput).ToErrorDefinitionResponsePtrOutputWithContext(ctx)
}









type ErrorDefinitionResponsePtrInput interface {
	pulumi.Input

	ToErrorDefinitionResponsePtrOutput() ErrorDefinitionResponsePtrOutput
	ToErrorDefinitionResponsePtrOutputWithContext(context.Context) ErrorDefinitionResponsePtrOutput
}

type errorDefinitionResponsePtrType ErrorDefinitionResponseArgs

func ErrorDefinitionResponsePtr(v *ErrorDefinitionResponseArgs) ErrorDefinitionResponsePtrInput {
	return (*errorDefinitionResponsePtrType)(v)
}

func (*errorDefinitionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorDefinitionResponse)(nil)).Elem()
}

func (i *errorDefinitionResponsePtrType) ToErrorDefinitionResponsePtrOutput() ErrorDefinitionResponsePtrOutput {
	return i.ToErrorDefinitionResponsePtrOutputWithContext(context.Background())
}

func (i *errorDefinitionResponsePtrType) ToErrorDefinitionResponsePtrOutputWithContext(ctx context.Context) ErrorDefinitionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDefinitionResponsePtrOutput)
}

type ErrorDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ErrorDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDefinitionResponse)(nil)).Elem()
}

func (o ErrorDefinitionResponseOutput) ToErrorDefinitionResponseOutput() ErrorDefinitionResponseOutput {
	return o
}

func (o ErrorDefinitionResponseOutput) ToErrorDefinitionResponseOutputWithContext(ctx context.Context) ErrorDefinitionResponseOutput {
	return o
}

func (o ErrorDefinitionResponseOutput) ToErrorDefinitionResponsePtrOutput() ErrorDefinitionResponsePtrOutput {
	return o.ToErrorDefinitionResponsePtrOutputWithContext(context.Background())
}

func (o ErrorDefinitionResponseOutput) ToErrorDefinitionResponsePtrOutputWithContext(ctx context.Context) ErrorDefinitionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ErrorDefinitionResponse) *ErrorDefinitionResponse {
		return &v
	}).(ErrorDefinitionResponsePtrOutput)
}

func (o ErrorDefinitionResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorDefinitionResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) string { return v.Message }).(pulumi.StringOutput)
}

type ErrorDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (ErrorDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorDefinitionResponse)(nil)).Elem()
}

func (o ErrorDefinitionResponsePtrOutput) ToErrorDefinitionResponsePtrOutput() ErrorDefinitionResponsePtrOutput {
	return o
}

func (o ErrorDefinitionResponsePtrOutput) ToErrorDefinitionResponsePtrOutputWithContext(ctx context.Context) ErrorDefinitionResponsePtrOutput {
	return o
}

func (o ErrorDefinitionResponsePtrOutput) Elem() ErrorDefinitionResponseOutput {
	return o.ApplyT(func(v *ErrorDefinitionResponse) ErrorDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret ErrorDefinitionResponse
		return ret
	}).(ErrorDefinitionResponseOutput)
}

func (o ErrorDefinitionResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ErrorDefinitionResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

type ExtensionStatus struct {
	Code          *string `pulumi:"code"`
	DisplayStatus *string `pulumi:"displayStatus"`
	Level         *string `pulumi:"level"`
	Message       *string `pulumi:"message"`
	Time          *string `pulumi:"time"`
}


func (val *ExtensionStatus) Defaults() *ExtensionStatus {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Level) {
		level_ := "Information"
		tmp.Level = &level_
	}
	return &tmp
}





type ExtensionStatusInput interface {
	pulumi.Input

	ToExtensionStatusOutput() ExtensionStatusOutput
	ToExtensionStatusOutputWithContext(context.Context) ExtensionStatusOutput
}

type ExtensionStatusArgs struct {
	Code          pulumi.StringPtrInput `pulumi:"code"`
	DisplayStatus pulumi.StringPtrInput `pulumi:"displayStatus"`
	Level         pulumi.StringPtrInput `pulumi:"level"`
	Message       pulumi.StringPtrInput `pulumi:"message"`
	Time          pulumi.StringPtrInput `pulumi:"time"`
}

func (ExtensionStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionStatus)(nil)).Elem()
}

func (i ExtensionStatusArgs) ToExtensionStatusOutput() ExtensionStatusOutput {
	return i.ToExtensionStatusOutputWithContext(context.Background())
}

func (i ExtensionStatusArgs) ToExtensionStatusOutputWithContext(ctx context.Context) ExtensionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionStatusOutput)
}





type ExtensionStatusArrayInput interface {
	pulumi.Input

	ToExtensionStatusArrayOutput() ExtensionStatusArrayOutput
	ToExtensionStatusArrayOutputWithContext(context.Context) ExtensionStatusArrayOutput
}

type ExtensionStatusArray []ExtensionStatusInput

func (ExtensionStatusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtensionStatus)(nil)).Elem()
}

func (i ExtensionStatusArray) ToExtensionStatusArrayOutput() ExtensionStatusArrayOutput {
	return i.ToExtensionStatusArrayOutputWithContext(context.Background())
}

func (i ExtensionStatusArray) ToExtensionStatusArrayOutputWithContext(ctx context.Context) ExtensionStatusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionStatusArrayOutput)
}

type ExtensionStatusOutput struct{ *pulumi.OutputState }

func (ExtensionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionStatus)(nil)).Elem()
}

func (o ExtensionStatusOutput) ToExtensionStatusOutput() ExtensionStatusOutput {
	return o
}

func (o ExtensionStatusOutput) ToExtensionStatusOutputWithContext(ctx context.Context) ExtensionStatusOutput {
	return o
}

func (o ExtensionStatusOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatus) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatus) *string { return v.DisplayStatus }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatus) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatus) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatus) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type ExtensionStatusArrayOutput struct{ *pulumi.OutputState }

func (ExtensionStatusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtensionStatus)(nil)).Elem()
}

func (o ExtensionStatusArrayOutput) ToExtensionStatusArrayOutput() ExtensionStatusArrayOutput {
	return o
}

func (o ExtensionStatusArrayOutput) ToExtensionStatusArrayOutputWithContext(ctx context.Context) ExtensionStatusArrayOutput {
	return o
}

func (o ExtensionStatusArrayOutput) Index(i pulumi.IntInput) ExtensionStatusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExtensionStatus {
		return vs[0].([]ExtensionStatus)[vs[1].(int)]
	}).(ExtensionStatusOutput)
}

type ExtensionStatusResponse struct {
	Code          *string `pulumi:"code"`
	DisplayStatus *string `pulumi:"displayStatus"`
	Level         *string `pulumi:"level"`
	Message       *string `pulumi:"message"`
	Time          *string `pulumi:"time"`
}


func (val *ExtensionStatusResponse) Defaults() *ExtensionStatusResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Level) {
		level_ := "Information"
		tmp.Level = &level_
	}
	return &tmp
}





type ExtensionStatusResponseInput interface {
	pulumi.Input

	ToExtensionStatusResponseOutput() ExtensionStatusResponseOutput
	ToExtensionStatusResponseOutputWithContext(context.Context) ExtensionStatusResponseOutput
}

type ExtensionStatusResponseArgs struct {
	Code          pulumi.StringPtrInput `pulumi:"code"`
	DisplayStatus pulumi.StringPtrInput `pulumi:"displayStatus"`
	Level         pulumi.StringPtrInput `pulumi:"level"`
	Message       pulumi.StringPtrInput `pulumi:"message"`
	Time          pulumi.StringPtrInput `pulumi:"time"`
}

func (ExtensionStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionStatusResponse)(nil)).Elem()
}

func (i ExtensionStatusResponseArgs) ToExtensionStatusResponseOutput() ExtensionStatusResponseOutput {
	return i.ToExtensionStatusResponseOutputWithContext(context.Background())
}

func (i ExtensionStatusResponseArgs) ToExtensionStatusResponseOutputWithContext(ctx context.Context) ExtensionStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionStatusResponseOutput)
}





type ExtensionStatusResponseArrayInput interface {
	pulumi.Input

	ToExtensionStatusResponseArrayOutput() ExtensionStatusResponseArrayOutput
	ToExtensionStatusResponseArrayOutputWithContext(context.Context) ExtensionStatusResponseArrayOutput
}

type ExtensionStatusResponseArray []ExtensionStatusResponseInput

func (ExtensionStatusResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtensionStatusResponse)(nil)).Elem()
}

func (i ExtensionStatusResponseArray) ToExtensionStatusResponseArrayOutput() ExtensionStatusResponseArrayOutput {
	return i.ToExtensionStatusResponseArrayOutputWithContext(context.Background())
}

func (i ExtensionStatusResponseArray) ToExtensionStatusResponseArrayOutputWithContext(ctx context.Context) ExtensionStatusResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionStatusResponseArrayOutput)
}

type ExtensionStatusResponseOutput struct{ *pulumi.OutputState }

func (ExtensionStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionStatusResponse)(nil)).Elem()
}

func (o ExtensionStatusResponseOutput) ToExtensionStatusResponseOutput() ExtensionStatusResponseOutput {
	return o
}

func (o ExtensionStatusResponseOutput) ToExtensionStatusResponseOutputWithContext(ctx context.Context) ExtensionStatusResponseOutput {
	return o
}

func (o ExtensionStatusResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatusResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusResponseOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatusResponse) *string { return v.DisplayStatus }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusResponseOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatusResponse) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatusResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusResponseOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatusResponse) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type ExtensionStatusResponseArrayOutput struct{ *pulumi.OutputState }

func (ExtensionStatusResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtensionStatusResponse)(nil)).Elem()
}

func (o ExtensionStatusResponseArrayOutput) ToExtensionStatusResponseArrayOutput() ExtensionStatusResponseArrayOutput {
	return o
}

func (o ExtensionStatusResponseArrayOutput) ToExtensionStatusResponseArrayOutputWithContext(ctx context.Context) ExtensionStatusResponseArrayOutput {
	return o
}

func (o ExtensionStatusResponseArrayOutput) Index(i pulumi.IntInput) ExtensionStatusResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExtensionStatusResponse {
		return vs[0].([]ExtensionStatusResponse)[vs[1].(int)]
	}).(ExtensionStatusResponseOutput)
}

type HelmOperatorProperties struct {
	ChartValues  *string `pulumi:"chartValues"`
	ChartVersion *string `pulumi:"chartVersion"`
}





type HelmOperatorPropertiesInput interface {
	pulumi.Input

	ToHelmOperatorPropertiesOutput() HelmOperatorPropertiesOutput
	ToHelmOperatorPropertiesOutputWithContext(context.Context) HelmOperatorPropertiesOutput
}

type HelmOperatorPropertiesArgs struct {
	ChartValues  pulumi.StringPtrInput `pulumi:"chartValues"`
	ChartVersion pulumi.StringPtrInput `pulumi:"chartVersion"`
}

func (HelmOperatorPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HelmOperatorProperties)(nil)).Elem()
}

func (i HelmOperatorPropertiesArgs) ToHelmOperatorPropertiesOutput() HelmOperatorPropertiesOutput {
	return i.ToHelmOperatorPropertiesOutputWithContext(context.Background())
}

func (i HelmOperatorPropertiesArgs) ToHelmOperatorPropertiesOutputWithContext(ctx context.Context) HelmOperatorPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmOperatorPropertiesOutput)
}

func (i HelmOperatorPropertiesArgs) ToHelmOperatorPropertiesPtrOutput() HelmOperatorPropertiesPtrOutput {
	return i.ToHelmOperatorPropertiesPtrOutputWithContext(context.Background())
}

func (i HelmOperatorPropertiesArgs) ToHelmOperatorPropertiesPtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmOperatorPropertiesOutput).ToHelmOperatorPropertiesPtrOutputWithContext(ctx)
}









type HelmOperatorPropertiesPtrInput interface {
	pulumi.Input

	ToHelmOperatorPropertiesPtrOutput() HelmOperatorPropertiesPtrOutput
	ToHelmOperatorPropertiesPtrOutputWithContext(context.Context) HelmOperatorPropertiesPtrOutput
}

type helmOperatorPropertiesPtrType HelmOperatorPropertiesArgs

func HelmOperatorPropertiesPtr(v *HelmOperatorPropertiesArgs) HelmOperatorPropertiesPtrInput {
	return (*helmOperatorPropertiesPtrType)(v)
}

func (*helmOperatorPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HelmOperatorProperties)(nil)).Elem()
}

func (i *helmOperatorPropertiesPtrType) ToHelmOperatorPropertiesPtrOutput() HelmOperatorPropertiesPtrOutput {
	return i.ToHelmOperatorPropertiesPtrOutputWithContext(context.Background())
}

func (i *helmOperatorPropertiesPtrType) ToHelmOperatorPropertiesPtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmOperatorPropertiesPtrOutput)
}

type HelmOperatorPropertiesOutput struct{ *pulumi.OutputState }

func (HelmOperatorPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HelmOperatorProperties)(nil)).Elem()
}

func (o HelmOperatorPropertiesOutput) ToHelmOperatorPropertiesOutput() HelmOperatorPropertiesOutput {
	return o
}

func (o HelmOperatorPropertiesOutput) ToHelmOperatorPropertiesOutputWithContext(ctx context.Context) HelmOperatorPropertiesOutput {
	return o
}

func (o HelmOperatorPropertiesOutput) ToHelmOperatorPropertiesPtrOutput() HelmOperatorPropertiesPtrOutput {
	return o.ToHelmOperatorPropertiesPtrOutputWithContext(context.Background())
}

func (o HelmOperatorPropertiesOutput) ToHelmOperatorPropertiesPtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HelmOperatorProperties) *HelmOperatorProperties {
		return &v
	}).(HelmOperatorPropertiesPtrOutput)
}

func (o HelmOperatorPropertiesOutput) ChartValues() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmOperatorProperties) *string { return v.ChartValues }).(pulumi.StringPtrOutput)
}

func (o HelmOperatorPropertiesOutput) ChartVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmOperatorProperties) *string { return v.ChartVersion }).(pulumi.StringPtrOutput)
}

type HelmOperatorPropertiesPtrOutput struct{ *pulumi.OutputState }

func (HelmOperatorPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HelmOperatorProperties)(nil)).Elem()
}

func (o HelmOperatorPropertiesPtrOutput) ToHelmOperatorPropertiesPtrOutput() HelmOperatorPropertiesPtrOutput {
	return o
}

func (o HelmOperatorPropertiesPtrOutput) ToHelmOperatorPropertiesPtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesPtrOutput {
	return o
}

func (o HelmOperatorPropertiesPtrOutput) Elem() HelmOperatorPropertiesOutput {
	return o.ApplyT(func(v *HelmOperatorProperties) HelmOperatorProperties {
		if v != nil {
			return *v
		}
		var ret HelmOperatorProperties
		return ret
	}).(HelmOperatorPropertiesOutput)
}

func (o HelmOperatorPropertiesPtrOutput) ChartValues() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmOperatorProperties) *string {
		if v == nil {
			return nil
		}
		return v.ChartValues
	}).(pulumi.StringPtrOutput)
}

func (o HelmOperatorPropertiesPtrOutput) ChartVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmOperatorProperties) *string {
		if v == nil {
			return nil
		}
		return v.ChartVersion
	}).(pulumi.StringPtrOutput)
}

type HelmOperatorPropertiesResponse struct {
	ChartValues  *string `pulumi:"chartValues"`
	ChartVersion *string `pulumi:"chartVersion"`
}





type HelmOperatorPropertiesResponseInput interface {
	pulumi.Input

	ToHelmOperatorPropertiesResponseOutput() HelmOperatorPropertiesResponseOutput
	ToHelmOperatorPropertiesResponseOutputWithContext(context.Context) HelmOperatorPropertiesResponseOutput
}

type HelmOperatorPropertiesResponseArgs struct {
	ChartValues  pulumi.StringPtrInput `pulumi:"chartValues"`
	ChartVersion pulumi.StringPtrInput `pulumi:"chartVersion"`
}

func (HelmOperatorPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HelmOperatorPropertiesResponse)(nil)).Elem()
}

func (i HelmOperatorPropertiesResponseArgs) ToHelmOperatorPropertiesResponseOutput() HelmOperatorPropertiesResponseOutput {
	return i.ToHelmOperatorPropertiesResponseOutputWithContext(context.Background())
}

func (i HelmOperatorPropertiesResponseArgs) ToHelmOperatorPropertiesResponseOutputWithContext(ctx context.Context) HelmOperatorPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmOperatorPropertiesResponseOutput)
}

func (i HelmOperatorPropertiesResponseArgs) ToHelmOperatorPropertiesResponsePtrOutput() HelmOperatorPropertiesResponsePtrOutput {
	return i.ToHelmOperatorPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i HelmOperatorPropertiesResponseArgs) ToHelmOperatorPropertiesResponsePtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmOperatorPropertiesResponseOutput).ToHelmOperatorPropertiesResponsePtrOutputWithContext(ctx)
}









type HelmOperatorPropertiesResponsePtrInput interface {
	pulumi.Input

	ToHelmOperatorPropertiesResponsePtrOutput() HelmOperatorPropertiesResponsePtrOutput
	ToHelmOperatorPropertiesResponsePtrOutputWithContext(context.Context) HelmOperatorPropertiesResponsePtrOutput
}

type helmOperatorPropertiesResponsePtrType HelmOperatorPropertiesResponseArgs

func HelmOperatorPropertiesResponsePtr(v *HelmOperatorPropertiesResponseArgs) HelmOperatorPropertiesResponsePtrInput {
	return (*helmOperatorPropertiesResponsePtrType)(v)
}

func (*helmOperatorPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HelmOperatorPropertiesResponse)(nil)).Elem()
}

func (i *helmOperatorPropertiesResponsePtrType) ToHelmOperatorPropertiesResponsePtrOutput() HelmOperatorPropertiesResponsePtrOutput {
	return i.ToHelmOperatorPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *helmOperatorPropertiesResponsePtrType) ToHelmOperatorPropertiesResponsePtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmOperatorPropertiesResponsePtrOutput)
}

type HelmOperatorPropertiesResponseOutput struct{ *pulumi.OutputState }

func (HelmOperatorPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HelmOperatorPropertiesResponse)(nil)).Elem()
}

func (o HelmOperatorPropertiesResponseOutput) ToHelmOperatorPropertiesResponseOutput() HelmOperatorPropertiesResponseOutput {
	return o
}

func (o HelmOperatorPropertiesResponseOutput) ToHelmOperatorPropertiesResponseOutputWithContext(ctx context.Context) HelmOperatorPropertiesResponseOutput {
	return o
}

func (o HelmOperatorPropertiesResponseOutput) ToHelmOperatorPropertiesResponsePtrOutput() HelmOperatorPropertiesResponsePtrOutput {
	return o.ToHelmOperatorPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o HelmOperatorPropertiesResponseOutput) ToHelmOperatorPropertiesResponsePtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HelmOperatorPropertiesResponse) *HelmOperatorPropertiesResponse {
		return &v
	}).(HelmOperatorPropertiesResponsePtrOutput)
}

func (o HelmOperatorPropertiesResponseOutput) ChartValues() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmOperatorPropertiesResponse) *string { return v.ChartValues }).(pulumi.StringPtrOutput)
}

func (o HelmOperatorPropertiesResponseOutput) ChartVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmOperatorPropertiesResponse) *string { return v.ChartVersion }).(pulumi.StringPtrOutput)
}

type HelmOperatorPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (HelmOperatorPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HelmOperatorPropertiesResponse)(nil)).Elem()
}

func (o HelmOperatorPropertiesResponsePtrOutput) ToHelmOperatorPropertiesResponsePtrOutput() HelmOperatorPropertiesResponsePtrOutput {
	return o
}

func (o HelmOperatorPropertiesResponsePtrOutput) ToHelmOperatorPropertiesResponsePtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesResponsePtrOutput {
	return o
}

func (o HelmOperatorPropertiesResponsePtrOutput) Elem() HelmOperatorPropertiesResponseOutput {
	return o.ApplyT(func(v *HelmOperatorPropertiesResponse) HelmOperatorPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret HelmOperatorPropertiesResponse
		return ret
	}).(HelmOperatorPropertiesResponseOutput)
}

func (o HelmOperatorPropertiesResponsePtrOutput) ChartValues() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmOperatorPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ChartValues
	}).(pulumi.StringPtrOutput)
}

func (o HelmOperatorPropertiesResponsePtrOutput) ChartVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmOperatorPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ChartVersion
	}).(pulumi.StringPtrOutput)
}

type Scope struct {
	Cluster   *ScopeCluster   `pulumi:"cluster"`
	Namespace *ScopeNamespace `pulumi:"namespace"`
}





type ScopeInput interface {
	pulumi.Input

	ToScopeOutput() ScopeOutput
	ToScopeOutputWithContext(context.Context) ScopeOutput
}

type ScopeArgs struct {
	Cluster   ScopeClusterPtrInput   `pulumi:"cluster"`
	Namespace ScopeNamespacePtrInput `pulumi:"namespace"`
}

func (ScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Scope)(nil)).Elem()
}

func (i ScopeArgs) ToScopeOutput() ScopeOutput {
	return i.ToScopeOutputWithContext(context.Background())
}

func (i ScopeArgs) ToScopeOutputWithContext(ctx context.Context) ScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeOutput)
}

func (i ScopeArgs) ToScopePtrOutput() ScopePtrOutput {
	return i.ToScopePtrOutputWithContext(context.Background())
}

func (i ScopeArgs) ToScopePtrOutputWithContext(ctx context.Context) ScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeOutput).ToScopePtrOutputWithContext(ctx)
}









type ScopePtrInput interface {
	pulumi.Input

	ToScopePtrOutput() ScopePtrOutput
	ToScopePtrOutputWithContext(context.Context) ScopePtrOutput
}

type scopePtrType ScopeArgs

func ScopePtr(v *ScopeArgs) ScopePtrInput {
	return (*scopePtrType)(v)
}

func (*scopePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Scope)(nil)).Elem()
}

func (i *scopePtrType) ToScopePtrOutput() ScopePtrOutput {
	return i.ToScopePtrOutputWithContext(context.Background())
}

func (i *scopePtrType) ToScopePtrOutputWithContext(ctx context.Context) ScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopePtrOutput)
}

type ScopeOutput struct{ *pulumi.OutputState }

func (ScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Scope)(nil)).Elem()
}

func (o ScopeOutput) ToScopeOutput() ScopeOutput {
	return o
}

func (o ScopeOutput) ToScopeOutputWithContext(ctx context.Context) ScopeOutput {
	return o
}

func (o ScopeOutput) ToScopePtrOutput() ScopePtrOutput {
	return o.ToScopePtrOutputWithContext(context.Background())
}

func (o ScopeOutput) ToScopePtrOutputWithContext(ctx context.Context) ScopePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Scope) *Scope {
		return &v
	}).(ScopePtrOutput)
}

func (o ScopeOutput) Cluster() ScopeClusterPtrOutput {
	return o.ApplyT(func(v Scope) *ScopeCluster { return v.Cluster }).(ScopeClusterPtrOutput)
}

func (o ScopeOutput) Namespace() ScopeNamespacePtrOutput {
	return o.ApplyT(func(v Scope) *ScopeNamespace { return v.Namespace }).(ScopeNamespacePtrOutput)
}

type ScopePtrOutput struct{ *pulumi.OutputState }

func (ScopePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Scope)(nil)).Elem()
}

func (o ScopePtrOutput) ToScopePtrOutput() ScopePtrOutput {
	return o
}

func (o ScopePtrOutput) ToScopePtrOutputWithContext(ctx context.Context) ScopePtrOutput {
	return o
}

func (o ScopePtrOutput) Elem() ScopeOutput {
	return o.ApplyT(func(v *Scope) Scope {
		if v != nil {
			return *v
		}
		var ret Scope
		return ret
	}).(ScopeOutput)
}

func (o ScopePtrOutput) Cluster() ScopeClusterPtrOutput {
	return o.ApplyT(func(v *Scope) *ScopeCluster {
		if v == nil {
			return nil
		}
		return v.Cluster
	}).(ScopeClusterPtrOutput)
}

func (o ScopePtrOutput) Namespace() ScopeNamespacePtrOutput {
	return o.ApplyT(func(v *Scope) *ScopeNamespace {
		if v == nil {
			return nil
		}
		return v.Namespace
	}).(ScopeNamespacePtrOutput)
}

type ScopeCluster struct {
	ReleaseNamespace *string `pulumi:"releaseNamespace"`
}





type ScopeClusterInput interface {
	pulumi.Input

	ToScopeClusterOutput() ScopeClusterOutput
	ToScopeClusterOutputWithContext(context.Context) ScopeClusterOutput
}

type ScopeClusterArgs struct {
	ReleaseNamespace pulumi.StringPtrInput `pulumi:"releaseNamespace"`
}

func (ScopeClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeCluster)(nil)).Elem()
}

func (i ScopeClusterArgs) ToScopeClusterOutput() ScopeClusterOutput {
	return i.ToScopeClusterOutputWithContext(context.Background())
}

func (i ScopeClusterArgs) ToScopeClusterOutputWithContext(ctx context.Context) ScopeClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeClusterOutput)
}

func (i ScopeClusterArgs) ToScopeClusterPtrOutput() ScopeClusterPtrOutput {
	return i.ToScopeClusterPtrOutputWithContext(context.Background())
}

func (i ScopeClusterArgs) ToScopeClusterPtrOutputWithContext(ctx context.Context) ScopeClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeClusterOutput).ToScopeClusterPtrOutputWithContext(ctx)
}









type ScopeClusterPtrInput interface {
	pulumi.Input

	ToScopeClusterPtrOutput() ScopeClusterPtrOutput
	ToScopeClusterPtrOutputWithContext(context.Context) ScopeClusterPtrOutput
}

type scopeClusterPtrType ScopeClusterArgs

func ScopeClusterPtr(v *ScopeClusterArgs) ScopeClusterPtrInput {
	return (*scopeClusterPtrType)(v)
}

func (*scopeClusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeCluster)(nil)).Elem()
}

func (i *scopeClusterPtrType) ToScopeClusterPtrOutput() ScopeClusterPtrOutput {
	return i.ToScopeClusterPtrOutputWithContext(context.Background())
}

func (i *scopeClusterPtrType) ToScopeClusterPtrOutputWithContext(ctx context.Context) ScopeClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeClusterPtrOutput)
}

type ScopeClusterOutput struct{ *pulumi.OutputState }

func (ScopeClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeCluster)(nil)).Elem()
}

func (o ScopeClusterOutput) ToScopeClusterOutput() ScopeClusterOutput {
	return o
}

func (o ScopeClusterOutput) ToScopeClusterOutputWithContext(ctx context.Context) ScopeClusterOutput {
	return o
}

func (o ScopeClusterOutput) ToScopeClusterPtrOutput() ScopeClusterPtrOutput {
	return o.ToScopeClusterPtrOutputWithContext(context.Background())
}

func (o ScopeClusterOutput) ToScopeClusterPtrOutputWithContext(ctx context.Context) ScopeClusterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScopeCluster) *ScopeCluster {
		return &v
	}).(ScopeClusterPtrOutput)
}

func (o ScopeClusterOutput) ReleaseNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScopeCluster) *string { return v.ReleaseNamespace }).(pulumi.StringPtrOutput)
}

type ScopeClusterPtrOutput struct{ *pulumi.OutputState }

func (ScopeClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeCluster)(nil)).Elem()
}

func (o ScopeClusterPtrOutput) ToScopeClusterPtrOutput() ScopeClusterPtrOutput {
	return o
}

func (o ScopeClusterPtrOutput) ToScopeClusterPtrOutputWithContext(ctx context.Context) ScopeClusterPtrOutput {
	return o
}

func (o ScopeClusterPtrOutput) Elem() ScopeClusterOutput {
	return o.ApplyT(func(v *ScopeCluster) ScopeCluster {
		if v != nil {
			return *v
		}
		var ret ScopeCluster
		return ret
	}).(ScopeClusterOutput)
}

func (o ScopeClusterPtrOutput) ReleaseNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeCluster) *string {
		if v == nil {
			return nil
		}
		return v.ReleaseNamespace
	}).(pulumi.StringPtrOutput)
}

type ScopeClusterResponse struct {
	ReleaseNamespace *string `pulumi:"releaseNamespace"`
}





type ScopeClusterResponseInput interface {
	pulumi.Input

	ToScopeClusterResponseOutput() ScopeClusterResponseOutput
	ToScopeClusterResponseOutputWithContext(context.Context) ScopeClusterResponseOutput
}

type ScopeClusterResponseArgs struct {
	ReleaseNamespace pulumi.StringPtrInput `pulumi:"releaseNamespace"`
}

func (ScopeClusterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeClusterResponse)(nil)).Elem()
}

func (i ScopeClusterResponseArgs) ToScopeClusterResponseOutput() ScopeClusterResponseOutput {
	return i.ToScopeClusterResponseOutputWithContext(context.Background())
}

func (i ScopeClusterResponseArgs) ToScopeClusterResponseOutputWithContext(ctx context.Context) ScopeClusterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeClusterResponseOutput)
}

func (i ScopeClusterResponseArgs) ToScopeClusterResponsePtrOutput() ScopeClusterResponsePtrOutput {
	return i.ToScopeClusterResponsePtrOutputWithContext(context.Background())
}

func (i ScopeClusterResponseArgs) ToScopeClusterResponsePtrOutputWithContext(ctx context.Context) ScopeClusterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeClusterResponseOutput).ToScopeClusterResponsePtrOutputWithContext(ctx)
}









type ScopeClusterResponsePtrInput interface {
	pulumi.Input

	ToScopeClusterResponsePtrOutput() ScopeClusterResponsePtrOutput
	ToScopeClusterResponsePtrOutputWithContext(context.Context) ScopeClusterResponsePtrOutput
}

type scopeClusterResponsePtrType ScopeClusterResponseArgs

func ScopeClusterResponsePtr(v *ScopeClusterResponseArgs) ScopeClusterResponsePtrInput {
	return (*scopeClusterResponsePtrType)(v)
}

func (*scopeClusterResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeClusterResponse)(nil)).Elem()
}

func (i *scopeClusterResponsePtrType) ToScopeClusterResponsePtrOutput() ScopeClusterResponsePtrOutput {
	return i.ToScopeClusterResponsePtrOutputWithContext(context.Background())
}

func (i *scopeClusterResponsePtrType) ToScopeClusterResponsePtrOutputWithContext(ctx context.Context) ScopeClusterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeClusterResponsePtrOutput)
}

type ScopeClusterResponseOutput struct{ *pulumi.OutputState }

func (ScopeClusterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeClusterResponse)(nil)).Elem()
}

func (o ScopeClusterResponseOutput) ToScopeClusterResponseOutput() ScopeClusterResponseOutput {
	return o
}

func (o ScopeClusterResponseOutput) ToScopeClusterResponseOutputWithContext(ctx context.Context) ScopeClusterResponseOutput {
	return o
}

func (o ScopeClusterResponseOutput) ToScopeClusterResponsePtrOutput() ScopeClusterResponsePtrOutput {
	return o.ToScopeClusterResponsePtrOutputWithContext(context.Background())
}

func (o ScopeClusterResponseOutput) ToScopeClusterResponsePtrOutputWithContext(ctx context.Context) ScopeClusterResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScopeClusterResponse) *ScopeClusterResponse {
		return &v
	}).(ScopeClusterResponsePtrOutput)
}

func (o ScopeClusterResponseOutput) ReleaseNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScopeClusterResponse) *string { return v.ReleaseNamespace }).(pulumi.StringPtrOutput)
}

type ScopeClusterResponsePtrOutput struct{ *pulumi.OutputState }

func (ScopeClusterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeClusterResponse)(nil)).Elem()
}

func (o ScopeClusterResponsePtrOutput) ToScopeClusterResponsePtrOutput() ScopeClusterResponsePtrOutput {
	return o
}

func (o ScopeClusterResponsePtrOutput) ToScopeClusterResponsePtrOutputWithContext(ctx context.Context) ScopeClusterResponsePtrOutput {
	return o
}

func (o ScopeClusterResponsePtrOutput) Elem() ScopeClusterResponseOutput {
	return o.ApplyT(func(v *ScopeClusterResponse) ScopeClusterResponse {
		if v != nil {
			return *v
		}
		var ret ScopeClusterResponse
		return ret
	}).(ScopeClusterResponseOutput)
}

func (o ScopeClusterResponsePtrOutput) ReleaseNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeClusterResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReleaseNamespace
	}).(pulumi.StringPtrOutput)
}

type ScopeNamespace struct {
	TargetNamespace *string `pulumi:"targetNamespace"`
}





type ScopeNamespaceInput interface {
	pulumi.Input

	ToScopeNamespaceOutput() ScopeNamespaceOutput
	ToScopeNamespaceOutputWithContext(context.Context) ScopeNamespaceOutput
}

type ScopeNamespaceArgs struct {
	TargetNamespace pulumi.StringPtrInput `pulumi:"targetNamespace"`
}

func (ScopeNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeNamespace)(nil)).Elem()
}

func (i ScopeNamespaceArgs) ToScopeNamespaceOutput() ScopeNamespaceOutput {
	return i.ToScopeNamespaceOutputWithContext(context.Background())
}

func (i ScopeNamespaceArgs) ToScopeNamespaceOutputWithContext(ctx context.Context) ScopeNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeNamespaceOutput)
}

func (i ScopeNamespaceArgs) ToScopeNamespacePtrOutput() ScopeNamespacePtrOutput {
	return i.ToScopeNamespacePtrOutputWithContext(context.Background())
}

func (i ScopeNamespaceArgs) ToScopeNamespacePtrOutputWithContext(ctx context.Context) ScopeNamespacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeNamespaceOutput).ToScopeNamespacePtrOutputWithContext(ctx)
}









type ScopeNamespacePtrInput interface {
	pulumi.Input

	ToScopeNamespacePtrOutput() ScopeNamespacePtrOutput
	ToScopeNamespacePtrOutputWithContext(context.Context) ScopeNamespacePtrOutput
}

type scopeNamespacePtrType ScopeNamespaceArgs

func ScopeNamespacePtr(v *ScopeNamespaceArgs) ScopeNamespacePtrInput {
	return (*scopeNamespacePtrType)(v)
}

func (*scopeNamespacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeNamespace)(nil)).Elem()
}

func (i *scopeNamespacePtrType) ToScopeNamespacePtrOutput() ScopeNamespacePtrOutput {
	return i.ToScopeNamespacePtrOutputWithContext(context.Background())
}

func (i *scopeNamespacePtrType) ToScopeNamespacePtrOutputWithContext(ctx context.Context) ScopeNamespacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeNamespacePtrOutput)
}

type ScopeNamespaceOutput struct{ *pulumi.OutputState }

func (ScopeNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeNamespace)(nil)).Elem()
}

func (o ScopeNamespaceOutput) ToScopeNamespaceOutput() ScopeNamespaceOutput {
	return o
}

func (o ScopeNamespaceOutput) ToScopeNamespaceOutputWithContext(ctx context.Context) ScopeNamespaceOutput {
	return o
}

func (o ScopeNamespaceOutput) ToScopeNamespacePtrOutput() ScopeNamespacePtrOutput {
	return o.ToScopeNamespacePtrOutputWithContext(context.Background())
}

func (o ScopeNamespaceOutput) ToScopeNamespacePtrOutputWithContext(ctx context.Context) ScopeNamespacePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScopeNamespace) *ScopeNamespace {
		return &v
	}).(ScopeNamespacePtrOutput)
}

func (o ScopeNamespaceOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScopeNamespace) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

type ScopeNamespacePtrOutput struct{ *pulumi.OutputState }

func (ScopeNamespacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeNamespace)(nil)).Elem()
}

func (o ScopeNamespacePtrOutput) ToScopeNamespacePtrOutput() ScopeNamespacePtrOutput {
	return o
}

func (o ScopeNamespacePtrOutput) ToScopeNamespacePtrOutputWithContext(ctx context.Context) ScopeNamespacePtrOutput {
	return o
}

func (o ScopeNamespacePtrOutput) Elem() ScopeNamespaceOutput {
	return o.ApplyT(func(v *ScopeNamespace) ScopeNamespace {
		if v != nil {
			return *v
		}
		var ret ScopeNamespace
		return ret
	}).(ScopeNamespaceOutput)
}

func (o ScopeNamespacePtrOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeNamespace) *string {
		if v == nil {
			return nil
		}
		return v.TargetNamespace
	}).(pulumi.StringPtrOutput)
}

type ScopeNamespaceResponse struct {
	TargetNamespace *string `pulumi:"targetNamespace"`
}





type ScopeNamespaceResponseInput interface {
	pulumi.Input

	ToScopeNamespaceResponseOutput() ScopeNamespaceResponseOutput
	ToScopeNamespaceResponseOutputWithContext(context.Context) ScopeNamespaceResponseOutput
}

type ScopeNamespaceResponseArgs struct {
	TargetNamespace pulumi.StringPtrInput `pulumi:"targetNamespace"`
}

func (ScopeNamespaceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeNamespaceResponse)(nil)).Elem()
}

func (i ScopeNamespaceResponseArgs) ToScopeNamespaceResponseOutput() ScopeNamespaceResponseOutput {
	return i.ToScopeNamespaceResponseOutputWithContext(context.Background())
}

func (i ScopeNamespaceResponseArgs) ToScopeNamespaceResponseOutputWithContext(ctx context.Context) ScopeNamespaceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeNamespaceResponseOutput)
}

func (i ScopeNamespaceResponseArgs) ToScopeNamespaceResponsePtrOutput() ScopeNamespaceResponsePtrOutput {
	return i.ToScopeNamespaceResponsePtrOutputWithContext(context.Background())
}

func (i ScopeNamespaceResponseArgs) ToScopeNamespaceResponsePtrOutputWithContext(ctx context.Context) ScopeNamespaceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeNamespaceResponseOutput).ToScopeNamespaceResponsePtrOutputWithContext(ctx)
}









type ScopeNamespaceResponsePtrInput interface {
	pulumi.Input

	ToScopeNamespaceResponsePtrOutput() ScopeNamespaceResponsePtrOutput
	ToScopeNamespaceResponsePtrOutputWithContext(context.Context) ScopeNamespaceResponsePtrOutput
}

type scopeNamespaceResponsePtrType ScopeNamespaceResponseArgs

func ScopeNamespaceResponsePtr(v *ScopeNamespaceResponseArgs) ScopeNamespaceResponsePtrInput {
	return (*scopeNamespaceResponsePtrType)(v)
}

func (*scopeNamespaceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeNamespaceResponse)(nil)).Elem()
}

func (i *scopeNamespaceResponsePtrType) ToScopeNamespaceResponsePtrOutput() ScopeNamespaceResponsePtrOutput {
	return i.ToScopeNamespaceResponsePtrOutputWithContext(context.Background())
}

func (i *scopeNamespaceResponsePtrType) ToScopeNamespaceResponsePtrOutputWithContext(ctx context.Context) ScopeNamespaceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeNamespaceResponsePtrOutput)
}

type ScopeNamespaceResponseOutput struct{ *pulumi.OutputState }

func (ScopeNamespaceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeNamespaceResponse)(nil)).Elem()
}

func (o ScopeNamespaceResponseOutput) ToScopeNamespaceResponseOutput() ScopeNamespaceResponseOutput {
	return o
}

func (o ScopeNamespaceResponseOutput) ToScopeNamespaceResponseOutputWithContext(ctx context.Context) ScopeNamespaceResponseOutput {
	return o
}

func (o ScopeNamespaceResponseOutput) ToScopeNamespaceResponsePtrOutput() ScopeNamespaceResponsePtrOutput {
	return o.ToScopeNamespaceResponsePtrOutputWithContext(context.Background())
}

func (o ScopeNamespaceResponseOutput) ToScopeNamespaceResponsePtrOutputWithContext(ctx context.Context) ScopeNamespaceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScopeNamespaceResponse) *ScopeNamespaceResponse {
		return &v
	}).(ScopeNamespaceResponsePtrOutput)
}

func (o ScopeNamespaceResponseOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScopeNamespaceResponse) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

type ScopeNamespaceResponsePtrOutput struct{ *pulumi.OutputState }

func (ScopeNamespaceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeNamespaceResponse)(nil)).Elem()
}

func (o ScopeNamespaceResponsePtrOutput) ToScopeNamespaceResponsePtrOutput() ScopeNamespaceResponsePtrOutput {
	return o
}

func (o ScopeNamespaceResponsePtrOutput) ToScopeNamespaceResponsePtrOutputWithContext(ctx context.Context) ScopeNamespaceResponsePtrOutput {
	return o
}

func (o ScopeNamespaceResponsePtrOutput) Elem() ScopeNamespaceResponseOutput {
	return o.ApplyT(func(v *ScopeNamespaceResponse) ScopeNamespaceResponse {
		if v != nil {
			return *v
		}
		var ret ScopeNamespaceResponse
		return ret
	}).(ScopeNamespaceResponseOutput)
}

func (o ScopeNamespaceResponsePtrOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeNamespaceResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetNamespace
	}).(pulumi.StringPtrOutput)
}

type ScopeResponse struct {
	Cluster   *ScopeClusterResponse   `pulumi:"cluster"`
	Namespace *ScopeNamespaceResponse `pulumi:"namespace"`
}





type ScopeResponseInput interface {
	pulumi.Input

	ToScopeResponseOutput() ScopeResponseOutput
	ToScopeResponseOutputWithContext(context.Context) ScopeResponseOutput
}

type ScopeResponseArgs struct {
	Cluster   ScopeClusterResponsePtrInput   `pulumi:"cluster"`
	Namespace ScopeNamespaceResponsePtrInput `pulumi:"namespace"`
}

func (ScopeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeResponse)(nil)).Elem()
}

func (i ScopeResponseArgs) ToScopeResponseOutput() ScopeResponseOutput {
	return i.ToScopeResponseOutputWithContext(context.Background())
}

func (i ScopeResponseArgs) ToScopeResponseOutputWithContext(ctx context.Context) ScopeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeResponseOutput)
}

func (i ScopeResponseArgs) ToScopeResponsePtrOutput() ScopeResponsePtrOutput {
	return i.ToScopeResponsePtrOutputWithContext(context.Background())
}

func (i ScopeResponseArgs) ToScopeResponsePtrOutputWithContext(ctx context.Context) ScopeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeResponseOutput).ToScopeResponsePtrOutputWithContext(ctx)
}









type ScopeResponsePtrInput interface {
	pulumi.Input

	ToScopeResponsePtrOutput() ScopeResponsePtrOutput
	ToScopeResponsePtrOutputWithContext(context.Context) ScopeResponsePtrOutput
}

type scopeResponsePtrType ScopeResponseArgs

func ScopeResponsePtr(v *ScopeResponseArgs) ScopeResponsePtrInput {
	return (*scopeResponsePtrType)(v)
}

func (*scopeResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeResponse)(nil)).Elem()
}

func (i *scopeResponsePtrType) ToScopeResponsePtrOutput() ScopeResponsePtrOutput {
	return i.ToScopeResponsePtrOutputWithContext(context.Background())
}

func (i *scopeResponsePtrType) ToScopeResponsePtrOutputWithContext(ctx context.Context) ScopeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeResponsePtrOutput)
}

type ScopeResponseOutput struct{ *pulumi.OutputState }

func (ScopeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeResponse)(nil)).Elem()
}

func (o ScopeResponseOutput) ToScopeResponseOutput() ScopeResponseOutput {
	return o
}

func (o ScopeResponseOutput) ToScopeResponseOutputWithContext(ctx context.Context) ScopeResponseOutput {
	return o
}

func (o ScopeResponseOutput) ToScopeResponsePtrOutput() ScopeResponsePtrOutput {
	return o.ToScopeResponsePtrOutputWithContext(context.Background())
}

func (o ScopeResponseOutput) ToScopeResponsePtrOutputWithContext(ctx context.Context) ScopeResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScopeResponse) *ScopeResponse {
		return &v
	}).(ScopeResponsePtrOutput)
}

func (o ScopeResponseOutput) Cluster() ScopeClusterResponsePtrOutput {
	return o.ApplyT(func(v ScopeResponse) *ScopeClusterResponse { return v.Cluster }).(ScopeClusterResponsePtrOutput)
}

func (o ScopeResponseOutput) Namespace() ScopeNamespaceResponsePtrOutput {
	return o.ApplyT(func(v ScopeResponse) *ScopeNamespaceResponse { return v.Namespace }).(ScopeNamespaceResponsePtrOutput)
}

type ScopeResponsePtrOutput struct{ *pulumi.OutputState }

func (ScopeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeResponse)(nil)).Elem()
}

func (o ScopeResponsePtrOutput) ToScopeResponsePtrOutput() ScopeResponsePtrOutput {
	return o
}

func (o ScopeResponsePtrOutput) ToScopeResponsePtrOutputWithContext(ctx context.Context) ScopeResponsePtrOutput {
	return o
}

func (o ScopeResponsePtrOutput) Elem() ScopeResponseOutput {
	return o.ApplyT(func(v *ScopeResponse) ScopeResponse {
		if v != nil {
			return *v
		}
		var ret ScopeResponse
		return ret
	}).(ScopeResponseOutput)
}

func (o ScopeResponsePtrOutput) Cluster() ScopeClusterResponsePtrOutput {
	return o.ApplyT(func(v *ScopeResponse) *ScopeClusterResponse {
		if v == nil {
			return nil
		}
		return v.Cluster
	}).(ScopeClusterResponsePtrOutput)
}

func (o ScopeResponsePtrOutput) Namespace() ScopeNamespaceResponsePtrOutput {
	return o.ApplyT(func(v *ScopeResponse) *ScopeNamespaceResponse {
		if v == nil {
			return nil
		}
		return v.Namespace
	}).(ScopeNamespaceResponsePtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          string `pulumi:"createdAt"`
	CreatedBy          string `pulumi:"createdBy"`
	CreatedByType      string `pulumi:"createdByType"`
	LastModifiedAt     string `pulumi:"lastModifiedAt"`
	LastModifiedBy     string `pulumi:"lastModifiedBy"`
	LastModifiedByType string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
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

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.CreatedBy }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.CreatedByType }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.LastModifiedAt }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.LastModifiedBy }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.LastModifiedByType }).(pulumi.StringOutput)
}

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ComplianceStatusResponseOutput{})
	pulumi.RegisterOutputType(ComplianceStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(ConfigurationIdentityOutput{})
	pulumi.RegisterOutputType(ConfigurationIdentityPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationIdentityResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ErrorDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ErrorDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(ExtensionStatusOutput{})
	pulumi.RegisterOutputType(ExtensionStatusArrayOutput{})
	pulumi.RegisterOutputType(ExtensionStatusResponseOutput{})
	pulumi.RegisterOutputType(ExtensionStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(HelmOperatorPropertiesOutput{})
	pulumi.RegisterOutputType(HelmOperatorPropertiesPtrOutput{})
	pulumi.RegisterOutputType(HelmOperatorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(HelmOperatorPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ScopeOutput{})
	pulumi.RegisterOutputType(ScopePtrOutput{})
	pulumi.RegisterOutputType(ScopeClusterOutput{})
	pulumi.RegisterOutputType(ScopeClusterPtrOutput{})
	pulumi.RegisterOutputType(ScopeClusterResponseOutput{})
	pulumi.RegisterOutputType(ScopeClusterResponsePtrOutput{})
	pulumi.RegisterOutputType(ScopeNamespaceOutput{})
	pulumi.RegisterOutputType(ScopeNamespacePtrOutput{})
	pulumi.RegisterOutputType(ScopeNamespaceResponseOutput{})
	pulumi.RegisterOutputType(ScopeNamespaceResponsePtrOutput{})
	pulumi.RegisterOutputType(ScopeResponseOutput{})
	pulumi.RegisterOutputType(ScopeResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}




package automanage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccountIdentity struct {
	Type *ResourceIdentityType `pulumi:"type"`
}





type AccountIdentityInput interface {
	pulumi.Input

	ToAccountIdentityOutput() AccountIdentityOutput
	ToAccountIdentityOutputWithContext(context.Context) AccountIdentityOutput
}

type AccountIdentityArgs struct {
	Type ResourceIdentityTypePtrInput `pulumi:"type"`
}

func (AccountIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountIdentity)(nil)).Elem()
}

func (i AccountIdentityArgs) ToAccountIdentityOutput() AccountIdentityOutput {
	return i.ToAccountIdentityOutputWithContext(context.Background())
}

func (i AccountIdentityArgs) ToAccountIdentityOutputWithContext(ctx context.Context) AccountIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountIdentityOutput)
}

func (i AccountIdentityArgs) ToAccountIdentityPtrOutput() AccountIdentityPtrOutput {
	return i.ToAccountIdentityPtrOutputWithContext(context.Background())
}

func (i AccountIdentityArgs) ToAccountIdentityPtrOutputWithContext(ctx context.Context) AccountIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountIdentityOutput).ToAccountIdentityPtrOutputWithContext(ctx)
}









type AccountIdentityPtrInput interface {
	pulumi.Input

	ToAccountIdentityPtrOutput() AccountIdentityPtrOutput
	ToAccountIdentityPtrOutputWithContext(context.Context) AccountIdentityPtrOutput
}

type accountIdentityPtrType AccountIdentityArgs

func AccountIdentityPtr(v *AccountIdentityArgs) AccountIdentityPtrInput {
	return (*accountIdentityPtrType)(v)
}

func (*accountIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountIdentity)(nil)).Elem()
}

func (i *accountIdentityPtrType) ToAccountIdentityPtrOutput() AccountIdentityPtrOutput {
	return i.ToAccountIdentityPtrOutputWithContext(context.Background())
}

func (i *accountIdentityPtrType) ToAccountIdentityPtrOutputWithContext(ctx context.Context) AccountIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountIdentityPtrOutput)
}

type AccountIdentityOutput struct{ *pulumi.OutputState }

func (AccountIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountIdentity)(nil)).Elem()
}

func (o AccountIdentityOutput) ToAccountIdentityOutput() AccountIdentityOutput {
	return o
}

func (o AccountIdentityOutput) ToAccountIdentityOutputWithContext(ctx context.Context) AccountIdentityOutput {
	return o
}

func (o AccountIdentityOutput) ToAccountIdentityPtrOutput() AccountIdentityPtrOutput {
	return o.ToAccountIdentityPtrOutputWithContext(context.Background())
}

func (o AccountIdentityOutput) ToAccountIdentityPtrOutputWithContext(ctx context.Context) AccountIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccountIdentity) *AccountIdentity {
		return &v
	}).(AccountIdentityPtrOutput)
}

func (o AccountIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v AccountIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

type AccountIdentityPtrOutput struct{ *pulumi.OutputState }

func (AccountIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountIdentity)(nil)).Elem()
}

func (o AccountIdentityPtrOutput) ToAccountIdentityPtrOutput() AccountIdentityPtrOutput {
	return o
}

func (o AccountIdentityPtrOutput) ToAccountIdentityPtrOutputWithContext(ctx context.Context) AccountIdentityPtrOutput {
	return o
}

func (o AccountIdentityPtrOutput) Elem() AccountIdentityOutput {
	return o.ApplyT(func(v *AccountIdentity) AccountIdentity {
		if v != nil {
			return *v
		}
		var ret AccountIdentity
		return ret
	}).(AccountIdentityOutput)
}

func (o AccountIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *AccountIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type AccountIdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type AccountIdentityResponseOutput struct{ *pulumi.OutputState }

func (AccountIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountIdentityResponse)(nil)).Elem()
}

func (o AccountIdentityResponseOutput) ToAccountIdentityResponseOutput() AccountIdentityResponseOutput {
	return o
}

func (o AccountIdentityResponseOutput) ToAccountIdentityResponseOutputWithContext(ctx context.Context) AccountIdentityResponseOutput {
	return o
}

func (o AccountIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v AccountIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o AccountIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v AccountIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o AccountIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type AccountIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (AccountIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountIdentityResponse)(nil)).Elem()
}

func (o AccountIdentityResponsePtrOutput) ToAccountIdentityResponsePtrOutput() AccountIdentityResponsePtrOutput {
	return o
}

func (o AccountIdentityResponsePtrOutput) ToAccountIdentityResponsePtrOutputWithContext(ctx context.Context) AccountIdentityResponsePtrOutput {
	return o
}

func (o AccountIdentityResponsePtrOutput) Elem() AccountIdentityResponseOutput {
	return o.ApplyT(func(v *AccountIdentityResponse) AccountIdentityResponse {
		if v != nil {
			return *v
		}
		var ret AccountIdentityResponse
		return ret
	}).(AccountIdentityResponseOutput)
}

func (o AccountIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o AccountIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o AccountIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ConfigurationProfileAssignmentComplianceResponse struct {
	UpdateStatus string `pulumi:"updateStatus"`
}

type ConfigurationProfileAssignmentComplianceResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileAssignmentComplianceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileAssignmentComplianceResponse)(nil)).Elem()
}

func (o ConfigurationProfileAssignmentComplianceResponseOutput) ToConfigurationProfileAssignmentComplianceResponseOutput() ConfigurationProfileAssignmentComplianceResponseOutput {
	return o
}

func (o ConfigurationProfileAssignmentComplianceResponseOutput) ToConfigurationProfileAssignmentComplianceResponseOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentComplianceResponseOutput {
	return o
}

func (o ConfigurationProfileAssignmentComplianceResponseOutput) UpdateStatus() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentComplianceResponse) string { return v.UpdateStatus }).(pulumi.StringOutput)
}

type ConfigurationProfileAssignmentComplianceResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileAssignmentComplianceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileAssignmentComplianceResponse)(nil)).Elem()
}

func (o ConfigurationProfileAssignmentComplianceResponsePtrOutput) ToConfigurationProfileAssignmentComplianceResponsePtrOutput() ConfigurationProfileAssignmentComplianceResponsePtrOutput {
	return o
}

func (o ConfigurationProfileAssignmentComplianceResponsePtrOutput) ToConfigurationProfileAssignmentComplianceResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentComplianceResponsePtrOutput {
	return o
}

func (o ConfigurationProfileAssignmentComplianceResponsePtrOutput) Elem() ConfigurationProfileAssignmentComplianceResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentComplianceResponse) ConfigurationProfileAssignmentComplianceResponse {
		if v != nil {
			return *v
		}
		var ret ConfigurationProfileAssignmentComplianceResponse
		return ret
	}).(ConfigurationProfileAssignmentComplianceResponseOutput)
}

func (o ConfigurationProfileAssignmentComplianceResponsePtrOutput) UpdateStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentComplianceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpdateStatus
	}).(pulumi.StringPtrOutput)
}

type ConfigurationProfileAssignmentProperties struct {
	AccountId                        *string `pulumi:"accountId"`
	ConfigurationProfile             *string `pulumi:"configurationProfile"`
	ConfigurationProfilePreferenceId *string `pulumi:"configurationProfilePreferenceId"`
	TargetId                         *string `pulumi:"targetId"`
}





type ConfigurationProfileAssignmentPropertiesInput interface {
	pulumi.Input

	ToConfigurationProfileAssignmentPropertiesOutput() ConfigurationProfileAssignmentPropertiesOutput
	ToConfigurationProfileAssignmentPropertiesOutputWithContext(context.Context) ConfigurationProfileAssignmentPropertiesOutput
}

type ConfigurationProfileAssignmentPropertiesArgs struct {
	AccountId                        pulumi.StringPtrInput `pulumi:"accountId"`
	ConfigurationProfile             pulumi.StringPtrInput `pulumi:"configurationProfile"`
	ConfigurationProfilePreferenceId pulumi.StringPtrInput `pulumi:"configurationProfilePreferenceId"`
	TargetId                         pulumi.StringPtrInput `pulumi:"targetId"`
}

func (ConfigurationProfileAssignmentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileAssignmentProperties)(nil)).Elem()
}

func (i ConfigurationProfileAssignmentPropertiesArgs) ToConfigurationProfileAssignmentPropertiesOutput() ConfigurationProfileAssignmentPropertiesOutput {
	return i.ToConfigurationProfileAssignmentPropertiesOutputWithContext(context.Background())
}

func (i ConfigurationProfileAssignmentPropertiesArgs) ToConfigurationProfileAssignmentPropertiesOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentPropertiesOutput)
}

func (i ConfigurationProfileAssignmentPropertiesArgs) ToConfigurationProfileAssignmentPropertiesPtrOutput() ConfigurationProfileAssignmentPropertiesPtrOutput {
	return i.ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (i ConfigurationProfileAssignmentPropertiesArgs) ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentPropertiesOutput).ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(ctx)
}









type ConfigurationProfileAssignmentPropertiesPtrInput interface {
	pulumi.Input

	ToConfigurationProfileAssignmentPropertiesPtrOutput() ConfigurationProfileAssignmentPropertiesPtrOutput
	ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(context.Context) ConfigurationProfileAssignmentPropertiesPtrOutput
}

type configurationProfileAssignmentPropertiesPtrType ConfigurationProfileAssignmentPropertiesArgs

func ConfigurationProfileAssignmentPropertiesPtr(v *ConfigurationProfileAssignmentPropertiesArgs) ConfigurationProfileAssignmentPropertiesPtrInput {
	return (*configurationProfileAssignmentPropertiesPtrType)(v)
}

func (*configurationProfileAssignmentPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileAssignmentProperties)(nil)).Elem()
}

func (i *configurationProfileAssignmentPropertiesPtrType) ToConfigurationProfileAssignmentPropertiesPtrOutput() ConfigurationProfileAssignmentPropertiesPtrOutput {
	return i.ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (i *configurationProfileAssignmentPropertiesPtrType) ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentPropertiesPtrOutput)
}

type ConfigurationProfileAssignmentPropertiesOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileAssignmentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileAssignmentProperties)(nil)).Elem()
}

func (o ConfigurationProfileAssignmentPropertiesOutput) ToConfigurationProfileAssignmentPropertiesOutput() ConfigurationProfileAssignmentPropertiesOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesOutput) ToConfigurationProfileAssignmentPropertiesOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesOutput) ToConfigurationProfileAssignmentPropertiesPtrOutput() ConfigurationProfileAssignmentPropertiesPtrOutput {
	return o.ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (o ConfigurationProfileAssignmentPropertiesOutput) ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationProfileAssignmentProperties) *ConfigurationProfileAssignmentProperties {
		return &v
	}).(ConfigurationProfileAssignmentPropertiesPtrOutput)
}

func (o ConfigurationProfileAssignmentPropertiesOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentProperties) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfileAssignmentPropertiesOutput) ConfigurationProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentProperties) *string { return v.ConfigurationProfile }).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfileAssignmentPropertiesOutput) ConfigurationProfilePreferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentProperties) *string { return v.ConfigurationProfilePreferenceId }).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfileAssignmentPropertiesOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentProperties) *string { return v.TargetId }).(pulumi.StringPtrOutput)
}

type ConfigurationProfileAssignmentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileAssignmentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileAssignmentProperties)(nil)).Elem()
}

func (o ConfigurationProfileAssignmentPropertiesPtrOutput) ToConfigurationProfileAssignmentPropertiesPtrOutput() ConfigurationProfileAssignmentPropertiesPtrOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesPtrOutput) ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesPtrOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesPtrOutput) Elem() ConfigurationProfileAssignmentPropertiesOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentProperties) ConfigurationProfileAssignmentProperties {
		if v != nil {
			return *v
		}
		var ret ConfigurationProfileAssignmentProperties
		return ret
	}).(ConfigurationProfileAssignmentPropertiesOutput)
}

func (o ConfigurationProfileAssignmentPropertiesPtrOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.AccountId
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfileAssignmentPropertiesPtrOutput) ConfigurationProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.ConfigurationProfile
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfileAssignmentPropertiesPtrOutput) ConfigurationProfilePreferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.ConfigurationProfilePreferenceId
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfileAssignmentPropertiesPtrOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.TargetId
	}).(pulumi.StringPtrOutput)
}

type ConfigurationProfileAssignmentPropertiesResponse struct {
	AccountId                        *string                                           `pulumi:"accountId"`
	Compliance                       *ConfigurationProfileAssignmentComplianceResponse `pulumi:"compliance"`
	ConfigurationProfile             *string                                           `pulumi:"configurationProfile"`
	ConfigurationProfilePreferenceId *string                                           `pulumi:"configurationProfilePreferenceId"`
	ProvisioningState                string                                            `pulumi:"provisioningState"`
	TargetId                         *string                                           `pulumi:"targetId"`
}

type ConfigurationProfileAssignmentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileAssignmentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileAssignmentPropertiesResponse)(nil)).Elem()
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ToConfigurationProfileAssignmentPropertiesResponseOutput() ConfigurationProfileAssignmentPropertiesResponseOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ToConfigurationProfileAssignmentPropertiesResponseOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesResponseOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentPropertiesResponse) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) Compliance() ConfigurationProfileAssignmentComplianceResponsePtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentPropertiesResponse) *ConfigurationProfileAssignmentComplianceResponse {
		return v.Compliance
	}).(ConfigurationProfileAssignmentComplianceResponsePtrOutput)
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ConfigurationProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentPropertiesResponse) *string { return v.ConfigurationProfile }).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ConfigurationProfilePreferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentPropertiesResponse) *string {
		return v.ConfigurationProfilePreferenceId
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentPropertiesResponse) *string { return v.TargetId }).(pulumi.StringPtrOutput)
}

type ConfigurationProfilePreferenceAntiMalware struct {
	EnableRealTimeProtection *string     `pulumi:"enableRealTimeProtection"`
	Exclusions               interface{} `pulumi:"exclusions"`
	RunScheduledScan         *string     `pulumi:"runScheduledScan"`
	ScanDay                  *string     `pulumi:"scanDay"`
	ScanTimeInMinutes        *string     `pulumi:"scanTimeInMinutes"`
	ScanType                 *string     `pulumi:"scanType"`
}





type ConfigurationProfilePreferenceAntiMalwareInput interface {
	pulumi.Input

	ToConfigurationProfilePreferenceAntiMalwareOutput() ConfigurationProfilePreferenceAntiMalwareOutput
	ToConfigurationProfilePreferenceAntiMalwareOutputWithContext(context.Context) ConfigurationProfilePreferenceAntiMalwareOutput
}

type ConfigurationProfilePreferenceAntiMalwareArgs struct {
	EnableRealTimeProtection pulumi.StringPtrInput `pulumi:"enableRealTimeProtection"`
	Exclusions               pulumi.Input          `pulumi:"exclusions"`
	RunScheduledScan         pulumi.StringPtrInput `pulumi:"runScheduledScan"`
	ScanDay                  pulumi.StringPtrInput `pulumi:"scanDay"`
	ScanTimeInMinutes        pulumi.StringPtrInput `pulumi:"scanTimeInMinutes"`
	ScanType                 pulumi.StringPtrInput `pulumi:"scanType"`
}

func (ConfigurationProfilePreferenceAntiMalwareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfilePreferenceAntiMalware)(nil)).Elem()
}

func (i ConfigurationProfilePreferenceAntiMalwareArgs) ToConfigurationProfilePreferenceAntiMalwareOutput() ConfigurationProfilePreferenceAntiMalwareOutput {
	return i.ToConfigurationProfilePreferenceAntiMalwareOutputWithContext(context.Background())
}

func (i ConfigurationProfilePreferenceAntiMalwareArgs) ToConfigurationProfilePreferenceAntiMalwareOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceAntiMalwareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilePreferenceAntiMalwareOutput)
}

func (i ConfigurationProfilePreferenceAntiMalwareArgs) ToConfigurationProfilePreferenceAntiMalwarePtrOutput() ConfigurationProfilePreferenceAntiMalwarePtrOutput {
	return i.ToConfigurationProfilePreferenceAntiMalwarePtrOutputWithContext(context.Background())
}

func (i ConfigurationProfilePreferenceAntiMalwareArgs) ToConfigurationProfilePreferenceAntiMalwarePtrOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceAntiMalwarePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilePreferenceAntiMalwareOutput).ToConfigurationProfilePreferenceAntiMalwarePtrOutputWithContext(ctx)
}









type ConfigurationProfilePreferenceAntiMalwarePtrInput interface {
	pulumi.Input

	ToConfigurationProfilePreferenceAntiMalwarePtrOutput() ConfigurationProfilePreferenceAntiMalwarePtrOutput
	ToConfigurationProfilePreferenceAntiMalwarePtrOutputWithContext(context.Context) ConfigurationProfilePreferenceAntiMalwarePtrOutput
}

type configurationProfilePreferenceAntiMalwarePtrType ConfigurationProfilePreferenceAntiMalwareArgs

func ConfigurationProfilePreferenceAntiMalwarePtr(v *ConfigurationProfilePreferenceAntiMalwareArgs) ConfigurationProfilePreferenceAntiMalwarePtrInput {
	return (*configurationProfilePreferenceAntiMalwarePtrType)(v)
}

func (*configurationProfilePreferenceAntiMalwarePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfilePreferenceAntiMalware)(nil)).Elem()
}

func (i *configurationProfilePreferenceAntiMalwarePtrType) ToConfigurationProfilePreferenceAntiMalwarePtrOutput() ConfigurationProfilePreferenceAntiMalwarePtrOutput {
	return i.ToConfigurationProfilePreferenceAntiMalwarePtrOutputWithContext(context.Background())
}

func (i *configurationProfilePreferenceAntiMalwarePtrType) ToConfigurationProfilePreferenceAntiMalwarePtrOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceAntiMalwarePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilePreferenceAntiMalwarePtrOutput)
}

type ConfigurationProfilePreferenceAntiMalwareOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePreferenceAntiMalwareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfilePreferenceAntiMalware)(nil)).Elem()
}

func (o ConfigurationProfilePreferenceAntiMalwareOutput) ToConfigurationProfilePreferenceAntiMalwareOutput() ConfigurationProfilePreferenceAntiMalwareOutput {
	return o
}

func (o ConfigurationProfilePreferenceAntiMalwareOutput) ToConfigurationProfilePreferenceAntiMalwareOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceAntiMalwareOutput {
	return o
}

func (o ConfigurationProfilePreferenceAntiMalwareOutput) ToConfigurationProfilePreferenceAntiMalwarePtrOutput() ConfigurationProfilePreferenceAntiMalwarePtrOutput {
	return o.ToConfigurationProfilePreferenceAntiMalwarePtrOutputWithContext(context.Background())
}

func (o ConfigurationProfilePreferenceAntiMalwareOutput) ToConfigurationProfilePreferenceAntiMalwarePtrOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceAntiMalwarePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationProfilePreferenceAntiMalware) *ConfigurationProfilePreferenceAntiMalware {
		return &v
	}).(ConfigurationProfilePreferenceAntiMalwarePtrOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwareOutput) EnableRealTimeProtection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceAntiMalware) *string { return v.EnableRealTimeProtection }).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwareOutput) Exclusions() pulumi.AnyOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceAntiMalware) interface{} { return v.Exclusions }).(pulumi.AnyOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwareOutput) RunScheduledScan() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceAntiMalware) *string { return v.RunScheduledScan }).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwareOutput) ScanDay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceAntiMalware) *string { return v.ScanDay }).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwareOutput) ScanTimeInMinutes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceAntiMalware) *string { return v.ScanTimeInMinutes }).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwareOutput) ScanType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceAntiMalware) *string { return v.ScanType }).(pulumi.StringPtrOutput)
}

type ConfigurationProfilePreferenceAntiMalwarePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePreferenceAntiMalwarePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfilePreferenceAntiMalware)(nil)).Elem()
}

func (o ConfigurationProfilePreferenceAntiMalwarePtrOutput) ToConfigurationProfilePreferenceAntiMalwarePtrOutput() ConfigurationProfilePreferenceAntiMalwarePtrOutput {
	return o
}

func (o ConfigurationProfilePreferenceAntiMalwarePtrOutput) ToConfigurationProfilePreferenceAntiMalwarePtrOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceAntiMalwarePtrOutput {
	return o
}

func (o ConfigurationProfilePreferenceAntiMalwarePtrOutput) Elem() ConfigurationProfilePreferenceAntiMalwareOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceAntiMalware) ConfigurationProfilePreferenceAntiMalware {
		if v != nil {
			return *v
		}
		var ret ConfigurationProfilePreferenceAntiMalware
		return ret
	}).(ConfigurationProfilePreferenceAntiMalwareOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwarePtrOutput) EnableRealTimeProtection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceAntiMalware) *string {
		if v == nil {
			return nil
		}
		return v.EnableRealTimeProtection
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwarePtrOutput) Exclusions() pulumi.AnyOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceAntiMalware) interface{} {
		if v == nil {
			return nil
		}
		return v.Exclusions
	}).(pulumi.AnyOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwarePtrOutput) RunScheduledScan() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceAntiMalware) *string {
		if v == nil {
			return nil
		}
		return v.RunScheduledScan
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwarePtrOutput) ScanDay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceAntiMalware) *string {
		if v == nil {
			return nil
		}
		return v.ScanDay
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwarePtrOutput) ScanTimeInMinutes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceAntiMalware) *string {
		if v == nil {
			return nil
		}
		return v.ScanTimeInMinutes
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwarePtrOutput) ScanType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceAntiMalware) *string {
		if v == nil {
			return nil
		}
		return v.ScanType
	}).(pulumi.StringPtrOutput)
}

type ConfigurationProfilePreferenceAntiMalwareResponse struct {
	EnableRealTimeProtection *string     `pulumi:"enableRealTimeProtection"`
	Exclusions               interface{} `pulumi:"exclusions"`
	RunScheduledScan         *string     `pulumi:"runScheduledScan"`
	ScanDay                  *string     `pulumi:"scanDay"`
	ScanTimeInMinutes        *string     `pulumi:"scanTimeInMinutes"`
	ScanType                 *string     `pulumi:"scanType"`
}

type ConfigurationProfilePreferenceAntiMalwareResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePreferenceAntiMalwareResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfilePreferenceAntiMalwareResponse)(nil)).Elem()
}

func (o ConfigurationProfilePreferenceAntiMalwareResponseOutput) ToConfigurationProfilePreferenceAntiMalwareResponseOutput() ConfigurationProfilePreferenceAntiMalwareResponseOutput {
	return o
}

func (o ConfigurationProfilePreferenceAntiMalwareResponseOutput) ToConfigurationProfilePreferenceAntiMalwareResponseOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceAntiMalwareResponseOutput {
	return o
}

func (o ConfigurationProfilePreferenceAntiMalwareResponseOutput) EnableRealTimeProtection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceAntiMalwareResponse) *string { return v.EnableRealTimeProtection }).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwareResponseOutput) Exclusions() pulumi.AnyOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceAntiMalwareResponse) interface{} { return v.Exclusions }).(pulumi.AnyOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwareResponseOutput) RunScheduledScan() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceAntiMalwareResponse) *string { return v.RunScheduledScan }).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwareResponseOutput) ScanDay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceAntiMalwareResponse) *string { return v.ScanDay }).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwareResponseOutput) ScanTimeInMinutes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceAntiMalwareResponse) *string { return v.ScanTimeInMinutes }).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwareResponseOutput) ScanType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceAntiMalwareResponse) *string { return v.ScanType }).(pulumi.StringPtrOutput)
}

type ConfigurationProfilePreferenceAntiMalwareResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePreferenceAntiMalwareResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfilePreferenceAntiMalwareResponse)(nil)).Elem()
}

func (o ConfigurationProfilePreferenceAntiMalwareResponsePtrOutput) ToConfigurationProfilePreferenceAntiMalwareResponsePtrOutput() ConfigurationProfilePreferenceAntiMalwareResponsePtrOutput {
	return o
}

func (o ConfigurationProfilePreferenceAntiMalwareResponsePtrOutput) ToConfigurationProfilePreferenceAntiMalwareResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceAntiMalwareResponsePtrOutput {
	return o
}

func (o ConfigurationProfilePreferenceAntiMalwareResponsePtrOutput) Elem() ConfigurationProfilePreferenceAntiMalwareResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceAntiMalwareResponse) ConfigurationProfilePreferenceAntiMalwareResponse {
		if v != nil {
			return *v
		}
		var ret ConfigurationProfilePreferenceAntiMalwareResponse
		return ret
	}).(ConfigurationProfilePreferenceAntiMalwareResponseOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwareResponsePtrOutput) EnableRealTimeProtection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceAntiMalwareResponse) *string {
		if v == nil {
			return nil
		}
		return v.EnableRealTimeProtection
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwareResponsePtrOutput) Exclusions() pulumi.AnyOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceAntiMalwareResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Exclusions
	}).(pulumi.AnyOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwareResponsePtrOutput) RunScheduledScan() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceAntiMalwareResponse) *string {
		if v == nil {
			return nil
		}
		return v.RunScheduledScan
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwareResponsePtrOutput) ScanDay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceAntiMalwareResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScanDay
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwareResponsePtrOutput) ScanTimeInMinutes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceAntiMalwareResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScanTimeInMinutes
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceAntiMalwareResponsePtrOutput) ScanType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceAntiMalwareResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScanType
	}).(pulumi.StringPtrOutput)
}

type ConfigurationProfilePreferenceProperties struct {
	AntiMalware *ConfigurationProfilePreferenceAntiMalware `pulumi:"antiMalware"`
	VmBackup    *ConfigurationProfilePreferenceVmBackup    `pulumi:"vmBackup"`
}





type ConfigurationProfilePreferencePropertiesInput interface {
	pulumi.Input

	ToConfigurationProfilePreferencePropertiesOutput() ConfigurationProfilePreferencePropertiesOutput
	ToConfigurationProfilePreferencePropertiesOutputWithContext(context.Context) ConfigurationProfilePreferencePropertiesOutput
}

type ConfigurationProfilePreferencePropertiesArgs struct {
	AntiMalware ConfigurationProfilePreferenceAntiMalwarePtrInput `pulumi:"antiMalware"`
	VmBackup    ConfigurationProfilePreferenceVmBackupPtrInput    `pulumi:"vmBackup"`
}

func (ConfigurationProfilePreferencePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfilePreferenceProperties)(nil)).Elem()
}

func (i ConfigurationProfilePreferencePropertiesArgs) ToConfigurationProfilePreferencePropertiesOutput() ConfigurationProfilePreferencePropertiesOutput {
	return i.ToConfigurationProfilePreferencePropertiesOutputWithContext(context.Background())
}

func (i ConfigurationProfilePreferencePropertiesArgs) ToConfigurationProfilePreferencePropertiesOutputWithContext(ctx context.Context) ConfigurationProfilePreferencePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilePreferencePropertiesOutput)
}

func (i ConfigurationProfilePreferencePropertiesArgs) ToConfigurationProfilePreferencePropertiesPtrOutput() ConfigurationProfilePreferencePropertiesPtrOutput {
	return i.ToConfigurationProfilePreferencePropertiesPtrOutputWithContext(context.Background())
}

func (i ConfigurationProfilePreferencePropertiesArgs) ToConfigurationProfilePreferencePropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfilePreferencePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilePreferencePropertiesOutput).ToConfigurationProfilePreferencePropertiesPtrOutputWithContext(ctx)
}









type ConfigurationProfilePreferencePropertiesPtrInput interface {
	pulumi.Input

	ToConfigurationProfilePreferencePropertiesPtrOutput() ConfigurationProfilePreferencePropertiesPtrOutput
	ToConfigurationProfilePreferencePropertiesPtrOutputWithContext(context.Context) ConfigurationProfilePreferencePropertiesPtrOutput
}

type configurationProfilePreferencePropertiesPtrType ConfigurationProfilePreferencePropertiesArgs

func ConfigurationProfilePreferencePropertiesPtr(v *ConfigurationProfilePreferencePropertiesArgs) ConfigurationProfilePreferencePropertiesPtrInput {
	return (*configurationProfilePreferencePropertiesPtrType)(v)
}

func (*configurationProfilePreferencePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfilePreferenceProperties)(nil)).Elem()
}

func (i *configurationProfilePreferencePropertiesPtrType) ToConfigurationProfilePreferencePropertiesPtrOutput() ConfigurationProfilePreferencePropertiesPtrOutput {
	return i.ToConfigurationProfilePreferencePropertiesPtrOutputWithContext(context.Background())
}

func (i *configurationProfilePreferencePropertiesPtrType) ToConfigurationProfilePreferencePropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfilePreferencePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilePreferencePropertiesPtrOutput)
}

type ConfigurationProfilePreferencePropertiesOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePreferencePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfilePreferenceProperties)(nil)).Elem()
}

func (o ConfigurationProfilePreferencePropertiesOutput) ToConfigurationProfilePreferencePropertiesOutput() ConfigurationProfilePreferencePropertiesOutput {
	return o
}

func (o ConfigurationProfilePreferencePropertiesOutput) ToConfigurationProfilePreferencePropertiesOutputWithContext(ctx context.Context) ConfigurationProfilePreferencePropertiesOutput {
	return o
}

func (o ConfigurationProfilePreferencePropertiesOutput) ToConfigurationProfilePreferencePropertiesPtrOutput() ConfigurationProfilePreferencePropertiesPtrOutput {
	return o.ToConfigurationProfilePreferencePropertiesPtrOutputWithContext(context.Background())
}

func (o ConfigurationProfilePreferencePropertiesOutput) ToConfigurationProfilePreferencePropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfilePreferencePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationProfilePreferenceProperties) *ConfigurationProfilePreferenceProperties {
		return &v
	}).(ConfigurationProfilePreferencePropertiesPtrOutput)
}

func (o ConfigurationProfilePreferencePropertiesOutput) AntiMalware() ConfigurationProfilePreferenceAntiMalwarePtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceProperties) *ConfigurationProfilePreferenceAntiMalware {
		return v.AntiMalware
	}).(ConfigurationProfilePreferenceAntiMalwarePtrOutput)
}

func (o ConfigurationProfilePreferencePropertiesOutput) VmBackup() ConfigurationProfilePreferenceVmBackupPtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceProperties) *ConfigurationProfilePreferenceVmBackup {
		return v.VmBackup
	}).(ConfigurationProfilePreferenceVmBackupPtrOutput)
}

type ConfigurationProfilePreferencePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePreferencePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfilePreferenceProperties)(nil)).Elem()
}

func (o ConfigurationProfilePreferencePropertiesPtrOutput) ToConfigurationProfilePreferencePropertiesPtrOutput() ConfigurationProfilePreferencePropertiesPtrOutput {
	return o
}

func (o ConfigurationProfilePreferencePropertiesPtrOutput) ToConfigurationProfilePreferencePropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfilePreferencePropertiesPtrOutput {
	return o
}

func (o ConfigurationProfilePreferencePropertiesPtrOutput) Elem() ConfigurationProfilePreferencePropertiesOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceProperties) ConfigurationProfilePreferenceProperties {
		if v != nil {
			return *v
		}
		var ret ConfigurationProfilePreferenceProperties
		return ret
	}).(ConfigurationProfilePreferencePropertiesOutput)
}

func (o ConfigurationProfilePreferencePropertiesPtrOutput) AntiMalware() ConfigurationProfilePreferenceAntiMalwarePtrOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceProperties) *ConfigurationProfilePreferenceAntiMalware {
		if v == nil {
			return nil
		}
		return v.AntiMalware
	}).(ConfigurationProfilePreferenceAntiMalwarePtrOutput)
}

func (o ConfigurationProfilePreferencePropertiesPtrOutput) VmBackup() ConfigurationProfilePreferenceVmBackupPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceProperties) *ConfigurationProfilePreferenceVmBackup {
		if v == nil {
			return nil
		}
		return v.VmBackup
	}).(ConfigurationProfilePreferenceVmBackupPtrOutput)
}

type ConfigurationProfilePreferencePropertiesResponse struct {
	AntiMalware *ConfigurationProfilePreferenceAntiMalwareResponse `pulumi:"antiMalware"`
	VmBackup    *ConfigurationProfilePreferenceVmBackupResponse    `pulumi:"vmBackup"`
}

type ConfigurationProfilePreferencePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePreferencePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfilePreferencePropertiesResponse)(nil)).Elem()
}

func (o ConfigurationProfilePreferencePropertiesResponseOutput) ToConfigurationProfilePreferencePropertiesResponseOutput() ConfigurationProfilePreferencePropertiesResponseOutput {
	return o
}

func (o ConfigurationProfilePreferencePropertiesResponseOutput) ToConfigurationProfilePreferencePropertiesResponseOutputWithContext(ctx context.Context) ConfigurationProfilePreferencePropertiesResponseOutput {
	return o
}

func (o ConfigurationProfilePreferencePropertiesResponseOutput) AntiMalware() ConfigurationProfilePreferenceAntiMalwareResponsePtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferencePropertiesResponse) *ConfigurationProfilePreferenceAntiMalwareResponse {
		return v.AntiMalware
	}).(ConfigurationProfilePreferenceAntiMalwareResponsePtrOutput)
}

func (o ConfigurationProfilePreferencePropertiesResponseOutput) VmBackup() ConfigurationProfilePreferenceVmBackupResponsePtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferencePropertiesResponse) *ConfigurationProfilePreferenceVmBackupResponse {
		return v.VmBackup
	}).(ConfigurationProfilePreferenceVmBackupResponsePtrOutput)
}

type ConfigurationProfilePreferenceVmBackup struct {
	InstantRpRetentionRangeInDays *int    `pulumi:"instantRpRetentionRangeInDays"`
	RetentionPolicy               *string `pulumi:"retentionPolicy"`
	SchedulePolicy                *string `pulumi:"schedulePolicy"`
	TimeZone                      *string `pulumi:"timeZone"`
}





type ConfigurationProfilePreferenceVmBackupInput interface {
	pulumi.Input

	ToConfigurationProfilePreferenceVmBackupOutput() ConfigurationProfilePreferenceVmBackupOutput
	ToConfigurationProfilePreferenceVmBackupOutputWithContext(context.Context) ConfigurationProfilePreferenceVmBackupOutput
}

type ConfigurationProfilePreferenceVmBackupArgs struct {
	InstantRpRetentionRangeInDays pulumi.IntPtrInput    `pulumi:"instantRpRetentionRangeInDays"`
	RetentionPolicy               pulumi.StringPtrInput `pulumi:"retentionPolicy"`
	SchedulePolicy                pulumi.StringPtrInput `pulumi:"schedulePolicy"`
	TimeZone                      pulumi.StringPtrInput `pulumi:"timeZone"`
}

func (ConfigurationProfilePreferenceVmBackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfilePreferenceVmBackup)(nil)).Elem()
}

func (i ConfigurationProfilePreferenceVmBackupArgs) ToConfigurationProfilePreferenceVmBackupOutput() ConfigurationProfilePreferenceVmBackupOutput {
	return i.ToConfigurationProfilePreferenceVmBackupOutputWithContext(context.Background())
}

func (i ConfigurationProfilePreferenceVmBackupArgs) ToConfigurationProfilePreferenceVmBackupOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceVmBackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilePreferenceVmBackupOutput)
}

func (i ConfigurationProfilePreferenceVmBackupArgs) ToConfigurationProfilePreferenceVmBackupPtrOutput() ConfigurationProfilePreferenceVmBackupPtrOutput {
	return i.ToConfigurationProfilePreferenceVmBackupPtrOutputWithContext(context.Background())
}

func (i ConfigurationProfilePreferenceVmBackupArgs) ToConfigurationProfilePreferenceVmBackupPtrOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceVmBackupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilePreferenceVmBackupOutput).ToConfigurationProfilePreferenceVmBackupPtrOutputWithContext(ctx)
}









type ConfigurationProfilePreferenceVmBackupPtrInput interface {
	pulumi.Input

	ToConfigurationProfilePreferenceVmBackupPtrOutput() ConfigurationProfilePreferenceVmBackupPtrOutput
	ToConfigurationProfilePreferenceVmBackupPtrOutputWithContext(context.Context) ConfigurationProfilePreferenceVmBackupPtrOutput
}

type configurationProfilePreferenceVmBackupPtrType ConfigurationProfilePreferenceVmBackupArgs

func ConfigurationProfilePreferenceVmBackupPtr(v *ConfigurationProfilePreferenceVmBackupArgs) ConfigurationProfilePreferenceVmBackupPtrInput {
	return (*configurationProfilePreferenceVmBackupPtrType)(v)
}

func (*configurationProfilePreferenceVmBackupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfilePreferenceVmBackup)(nil)).Elem()
}

func (i *configurationProfilePreferenceVmBackupPtrType) ToConfigurationProfilePreferenceVmBackupPtrOutput() ConfigurationProfilePreferenceVmBackupPtrOutput {
	return i.ToConfigurationProfilePreferenceVmBackupPtrOutputWithContext(context.Background())
}

func (i *configurationProfilePreferenceVmBackupPtrType) ToConfigurationProfilePreferenceVmBackupPtrOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceVmBackupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilePreferenceVmBackupPtrOutput)
}

type ConfigurationProfilePreferenceVmBackupOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePreferenceVmBackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfilePreferenceVmBackup)(nil)).Elem()
}

func (o ConfigurationProfilePreferenceVmBackupOutput) ToConfigurationProfilePreferenceVmBackupOutput() ConfigurationProfilePreferenceVmBackupOutput {
	return o
}

func (o ConfigurationProfilePreferenceVmBackupOutput) ToConfigurationProfilePreferenceVmBackupOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceVmBackupOutput {
	return o
}

func (o ConfigurationProfilePreferenceVmBackupOutput) ToConfigurationProfilePreferenceVmBackupPtrOutput() ConfigurationProfilePreferenceVmBackupPtrOutput {
	return o.ToConfigurationProfilePreferenceVmBackupPtrOutputWithContext(context.Background())
}

func (o ConfigurationProfilePreferenceVmBackupOutput) ToConfigurationProfilePreferenceVmBackupPtrOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceVmBackupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationProfilePreferenceVmBackup) *ConfigurationProfilePreferenceVmBackup {
		return &v
	}).(ConfigurationProfilePreferenceVmBackupPtrOutput)
}

func (o ConfigurationProfilePreferenceVmBackupOutput) InstantRpRetentionRangeInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceVmBackup) *int { return v.InstantRpRetentionRangeInDays }).(pulumi.IntPtrOutput)
}

func (o ConfigurationProfilePreferenceVmBackupOutput) RetentionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceVmBackup) *string { return v.RetentionPolicy }).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceVmBackupOutput) SchedulePolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceVmBackup) *string { return v.SchedulePolicy }).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceVmBackupOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceVmBackup) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type ConfigurationProfilePreferenceVmBackupPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePreferenceVmBackupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfilePreferenceVmBackup)(nil)).Elem()
}

func (o ConfigurationProfilePreferenceVmBackupPtrOutput) ToConfigurationProfilePreferenceVmBackupPtrOutput() ConfigurationProfilePreferenceVmBackupPtrOutput {
	return o
}

func (o ConfigurationProfilePreferenceVmBackupPtrOutput) ToConfigurationProfilePreferenceVmBackupPtrOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceVmBackupPtrOutput {
	return o
}

func (o ConfigurationProfilePreferenceVmBackupPtrOutput) Elem() ConfigurationProfilePreferenceVmBackupOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceVmBackup) ConfigurationProfilePreferenceVmBackup {
		if v != nil {
			return *v
		}
		var ret ConfigurationProfilePreferenceVmBackup
		return ret
	}).(ConfigurationProfilePreferenceVmBackupOutput)
}

func (o ConfigurationProfilePreferenceVmBackupPtrOutput) InstantRpRetentionRangeInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceVmBackup) *int {
		if v == nil {
			return nil
		}
		return v.InstantRpRetentionRangeInDays
	}).(pulumi.IntPtrOutput)
}

func (o ConfigurationProfilePreferenceVmBackupPtrOutput) RetentionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceVmBackup) *string {
		if v == nil {
			return nil
		}
		return v.RetentionPolicy
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceVmBackupPtrOutput) SchedulePolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceVmBackup) *string {
		if v == nil {
			return nil
		}
		return v.SchedulePolicy
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceVmBackupPtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceVmBackup) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

type ConfigurationProfilePreferenceVmBackupResponse struct {
	InstantRpRetentionRangeInDays *int    `pulumi:"instantRpRetentionRangeInDays"`
	RetentionPolicy               *string `pulumi:"retentionPolicy"`
	SchedulePolicy                *string `pulumi:"schedulePolicy"`
	TimeZone                      *string `pulumi:"timeZone"`
}

type ConfigurationProfilePreferenceVmBackupResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePreferenceVmBackupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfilePreferenceVmBackupResponse)(nil)).Elem()
}

func (o ConfigurationProfilePreferenceVmBackupResponseOutput) ToConfigurationProfilePreferenceVmBackupResponseOutput() ConfigurationProfilePreferenceVmBackupResponseOutput {
	return o
}

func (o ConfigurationProfilePreferenceVmBackupResponseOutput) ToConfigurationProfilePreferenceVmBackupResponseOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceVmBackupResponseOutput {
	return o
}

func (o ConfigurationProfilePreferenceVmBackupResponseOutput) InstantRpRetentionRangeInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceVmBackupResponse) *int { return v.InstantRpRetentionRangeInDays }).(pulumi.IntPtrOutput)
}

func (o ConfigurationProfilePreferenceVmBackupResponseOutput) RetentionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceVmBackupResponse) *string { return v.RetentionPolicy }).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceVmBackupResponseOutput) SchedulePolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceVmBackupResponse) *string { return v.SchedulePolicy }).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceVmBackupResponseOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfilePreferenceVmBackupResponse) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type ConfigurationProfilePreferenceVmBackupResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePreferenceVmBackupResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfilePreferenceVmBackupResponse)(nil)).Elem()
}

func (o ConfigurationProfilePreferenceVmBackupResponsePtrOutput) ToConfigurationProfilePreferenceVmBackupResponsePtrOutput() ConfigurationProfilePreferenceVmBackupResponsePtrOutput {
	return o
}

func (o ConfigurationProfilePreferenceVmBackupResponsePtrOutput) ToConfigurationProfilePreferenceVmBackupResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceVmBackupResponsePtrOutput {
	return o
}

func (o ConfigurationProfilePreferenceVmBackupResponsePtrOutput) Elem() ConfigurationProfilePreferenceVmBackupResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceVmBackupResponse) ConfigurationProfilePreferenceVmBackupResponse {
		if v != nil {
			return *v
		}
		var ret ConfigurationProfilePreferenceVmBackupResponse
		return ret
	}).(ConfigurationProfilePreferenceVmBackupResponseOutput)
}

func (o ConfigurationProfilePreferenceVmBackupResponsePtrOutput) InstantRpRetentionRangeInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceVmBackupResponse) *int {
		if v == nil {
			return nil
		}
		return v.InstantRpRetentionRangeInDays
	}).(pulumi.IntPtrOutput)
}

func (o ConfigurationProfilePreferenceVmBackupResponsePtrOutput) RetentionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceVmBackupResponse) *string {
		if v == nil {
			return nil
		}
		return v.RetentionPolicy
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceVmBackupResponsePtrOutput) SchedulePolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceVmBackupResponse) *string {
		if v == nil {
			return nil
		}
		return v.SchedulePolicy
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationProfilePreferenceVmBackupResponsePtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreferenceVmBackupResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountIdentityOutput{})
	pulumi.RegisterOutputType(AccountIdentityPtrOutput{})
	pulumi.RegisterOutputType(AccountIdentityResponseOutput{})
	pulumi.RegisterOutputType(AccountIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentComplianceResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentComplianceResponsePtrOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentPropertiesOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePreferenceAntiMalwareOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePreferenceAntiMalwarePtrOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePreferenceAntiMalwareResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePreferenceAntiMalwareResponsePtrOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePreferencePropertiesOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePreferencePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePreferencePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePreferenceVmBackupOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePreferenceVmBackupPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePreferenceVmBackupResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePreferenceVmBackupResponsePtrOutput{})
}

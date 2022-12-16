


package v20220930preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureMonitorAlertSettings struct {
	AlertsForAllJobFailures *string `pulumi:"alertsForAllJobFailures"`
}





type AzureMonitorAlertSettingsInput interface {
	pulumi.Input

	ToAzureMonitorAlertSettingsOutput() AzureMonitorAlertSettingsOutput
	ToAzureMonitorAlertSettingsOutputWithContext(context.Context) AzureMonitorAlertSettingsOutput
}

type AzureMonitorAlertSettingsArgs struct {
	AlertsForAllJobFailures pulumi.StringPtrInput `pulumi:"alertsForAllJobFailures"`
}

func (AzureMonitorAlertSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMonitorAlertSettings)(nil)).Elem()
}

func (i AzureMonitorAlertSettingsArgs) ToAzureMonitorAlertSettingsOutput() AzureMonitorAlertSettingsOutput {
	return i.ToAzureMonitorAlertSettingsOutputWithContext(context.Background())
}

func (i AzureMonitorAlertSettingsArgs) ToAzureMonitorAlertSettingsOutputWithContext(ctx context.Context) AzureMonitorAlertSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMonitorAlertSettingsOutput)
}

func (i AzureMonitorAlertSettingsArgs) ToAzureMonitorAlertSettingsPtrOutput() AzureMonitorAlertSettingsPtrOutput {
	return i.ToAzureMonitorAlertSettingsPtrOutputWithContext(context.Background())
}

func (i AzureMonitorAlertSettingsArgs) ToAzureMonitorAlertSettingsPtrOutputWithContext(ctx context.Context) AzureMonitorAlertSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMonitorAlertSettingsOutput).ToAzureMonitorAlertSettingsPtrOutputWithContext(ctx)
}









type AzureMonitorAlertSettingsPtrInput interface {
	pulumi.Input

	ToAzureMonitorAlertSettingsPtrOutput() AzureMonitorAlertSettingsPtrOutput
	ToAzureMonitorAlertSettingsPtrOutputWithContext(context.Context) AzureMonitorAlertSettingsPtrOutput
}

type azureMonitorAlertSettingsPtrType AzureMonitorAlertSettingsArgs

func AzureMonitorAlertSettingsPtr(v *AzureMonitorAlertSettingsArgs) AzureMonitorAlertSettingsPtrInput {
	return (*azureMonitorAlertSettingsPtrType)(v)
}

func (*azureMonitorAlertSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureMonitorAlertSettings)(nil)).Elem()
}

func (i *azureMonitorAlertSettingsPtrType) ToAzureMonitorAlertSettingsPtrOutput() AzureMonitorAlertSettingsPtrOutput {
	return i.ToAzureMonitorAlertSettingsPtrOutputWithContext(context.Background())
}

func (i *azureMonitorAlertSettingsPtrType) ToAzureMonitorAlertSettingsPtrOutputWithContext(ctx context.Context) AzureMonitorAlertSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMonitorAlertSettingsPtrOutput)
}

type AzureMonitorAlertSettingsOutput struct{ *pulumi.OutputState }

func (AzureMonitorAlertSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMonitorAlertSettings)(nil)).Elem()
}

func (o AzureMonitorAlertSettingsOutput) ToAzureMonitorAlertSettingsOutput() AzureMonitorAlertSettingsOutput {
	return o
}

func (o AzureMonitorAlertSettingsOutput) ToAzureMonitorAlertSettingsOutputWithContext(ctx context.Context) AzureMonitorAlertSettingsOutput {
	return o
}

func (o AzureMonitorAlertSettingsOutput) ToAzureMonitorAlertSettingsPtrOutput() AzureMonitorAlertSettingsPtrOutput {
	return o.ToAzureMonitorAlertSettingsPtrOutputWithContext(context.Background())
}

func (o AzureMonitorAlertSettingsOutput) ToAzureMonitorAlertSettingsPtrOutputWithContext(ctx context.Context) AzureMonitorAlertSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureMonitorAlertSettings) *AzureMonitorAlertSettings {
		return &v
	}).(AzureMonitorAlertSettingsPtrOutput)
}

func (o AzureMonitorAlertSettingsOutput) AlertsForAllJobFailures() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMonitorAlertSettings) *string { return v.AlertsForAllJobFailures }).(pulumi.StringPtrOutput)
}

type AzureMonitorAlertSettingsPtrOutput struct{ *pulumi.OutputState }

func (AzureMonitorAlertSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureMonitorAlertSettings)(nil)).Elem()
}

func (o AzureMonitorAlertSettingsPtrOutput) ToAzureMonitorAlertSettingsPtrOutput() AzureMonitorAlertSettingsPtrOutput {
	return o
}

func (o AzureMonitorAlertSettingsPtrOutput) ToAzureMonitorAlertSettingsPtrOutputWithContext(ctx context.Context) AzureMonitorAlertSettingsPtrOutput {
	return o
}

func (o AzureMonitorAlertSettingsPtrOutput) Elem() AzureMonitorAlertSettingsOutput {
	return o.ApplyT(func(v *AzureMonitorAlertSettings) AzureMonitorAlertSettings {
		if v != nil {
			return *v
		}
		var ret AzureMonitorAlertSettings
		return ret
	}).(AzureMonitorAlertSettingsOutput)
}

func (o AzureMonitorAlertSettingsPtrOutput) AlertsForAllJobFailures() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMonitorAlertSettings) *string {
		if v == nil {
			return nil
		}
		return v.AlertsForAllJobFailures
	}).(pulumi.StringPtrOutput)
}

type AzureMonitorAlertSettingsResponse struct {
	AlertsForAllJobFailures *string `pulumi:"alertsForAllJobFailures"`
}

type AzureMonitorAlertSettingsResponseOutput struct{ *pulumi.OutputState }

func (AzureMonitorAlertSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMonitorAlertSettingsResponse)(nil)).Elem()
}

func (o AzureMonitorAlertSettingsResponseOutput) ToAzureMonitorAlertSettingsResponseOutput() AzureMonitorAlertSettingsResponseOutput {
	return o
}

func (o AzureMonitorAlertSettingsResponseOutput) ToAzureMonitorAlertSettingsResponseOutputWithContext(ctx context.Context) AzureMonitorAlertSettingsResponseOutput {
	return o
}

func (o AzureMonitorAlertSettingsResponseOutput) AlertsForAllJobFailures() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMonitorAlertSettingsResponse) *string { return v.AlertsForAllJobFailures }).(pulumi.StringPtrOutput)
}

type AzureMonitorAlertSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureMonitorAlertSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureMonitorAlertSettingsResponse)(nil)).Elem()
}

func (o AzureMonitorAlertSettingsResponsePtrOutput) ToAzureMonitorAlertSettingsResponsePtrOutput() AzureMonitorAlertSettingsResponsePtrOutput {
	return o
}

func (o AzureMonitorAlertSettingsResponsePtrOutput) ToAzureMonitorAlertSettingsResponsePtrOutputWithContext(ctx context.Context) AzureMonitorAlertSettingsResponsePtrOutput {
	return o
}

func (o AzureMonitorAlertSettingsResponsePtrOutput) Elem() AzureMonitorAlertSettingsResponseOutput {
	return o.ApplyT(func(v *AzureMonitorAlertSettingsResponse) AzureMonitorAlertSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AzureMonitorAlertSettingsResponse
		return ret
	}).(AzureMonitorAlertSettingsResponseOutput)
}

func (o AzureMonitorAlertSettingsResponsePtrOutput) AlertsForAllJobFailures() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMonitorAlertSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AlertsForAllJobFailures
	}).(pulumi.StringPtrOutput)
}

type ClassicAlertSettings struct {
	AlertsForCriticalOperations *string `pulumi:"alertsForCriticalOperations"`
}





type ClassicAlertSettingsInput interface {
	pulumi.Input

	ToClassicAlertSettingsOutput() ClassicAlertSettingsOutput
	ToClassicAlertSettingsOutputWithContext(context.Context) ClassicAlertSettingsOutput
}

type ClassicAlertSettingsArgs struct {
	AlertsForCriticalOperations pulumi.StringPtrInput `pulumi:"alertsForCriticalOperations"`
}

func (ClassicAlertSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClassicAlertSettings)(nil)).Elem()
}

func (i ClassicAlertSettingsArgs) ToClassicAlertSettingsOutput() ClassicAlertSettingsOutput {
	return i.ToClassicAlertSettingsOutputWithContext(context.Background())
}

func (i ClassicAlertSettingsArgs) ToClassicAlertSettingsOutputWithContext(ctx context.Context) ClassicAlertSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClassicAlertSettingsOutput)
}

func (i ClassicAlertSettingsArgs) ToClassicAlertSettingsPtrOutput() ClassicAlertSettingsPtrOutput {
	return i.ToClassicAlertSettingsPtrOutputWithContext(context.Background())
}

func (i ClassicAlertSettingsArgs) ToClassicAlertSettingsPtrOutputWithContext(ctx context.Context) ClassicAlertSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClassicAlertSettingsOutput).ToClassicAlertSettingsPtrOutputWithContext(ctx)
}









type ClassicAlertSettingsPtrInput interface {
	pulumi.Input

	ToClassicAlertSettingsPtrOutput() ClassicAlertSettingsPtrOutput
	ToClassicAlertSettingsPtrOutputWithContext(context.Context) ClassicAlertSettingsPtrOutput
}

type classicAlertSettingsPtrType ClassicAlertSettingsArgs

func ClassicAlertSettingsPtr(v *ClassicAlertSettingsArgs) ClassicAlertSettingsPtrInput {
	return (*classicAlertSettingsPtrType)(v)
}

func (*classicAlertSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClassicAlertSettings)(nil)).Elem()
}

func (i *classicAlertSettingsPtrType) ToClassicAlertSettingsPtrOutput() ClassicAlertSettingsPtrOutput {
	return i.ToClassicAlertSettingsPtrOutputWithContext(context.Background())
}

func (i *classicAlertSettingsPtrType) ToClassicAlertSettingsPtrOutputWithContext(ctx context.Context) ClassicAlertSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClassicAlertSettingsPtrOutput)
}

type ClassicAlertSettingsOutput struct{ *pulumi.OutputState }

func (ClassicAlertSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClassicAlertSettings)(nil)).Elem()
}

func (o ClassicAlertSettingsOutput) ToClassicAlertSettingsOutput() ClassicAlertSettingsOutput {
	return o
}

func (o ClassicAlertSettingsOutput) ToClassicAlertSettingsOutputWithContext(ctx context.Context) ClassicAlertSettingsOutput {
	return o
}

func (o ClassicAlertSettingsOutput) ToClassicAlertSettingsPtrOutput() ClassicAlertSettingsPtrOutput {
	return o.ToClassicAlertSettingsPtrOutputWithContext(context.Background())
}

func (o ClassicAlertSettingsOutput) ToClassicAlertSettingsPtrOutputWithContext(ctx context.Context) ClassicAlertSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClassicAlertSettings) *ClassicAlertSettings {
		return &v
	}).(ClassicAlertSettingsPtrOutput)
}

func (o ClassicAlertSettingsOutput) AlertsForCriticalOperations() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClassicAlertSettings) *string { return v.AlertsForCriticalOperations }).(pulumi.StringPtrOutput)
}

type ClassicAlertSettingsPtrOutput struct{ *pulumi.OutputState }

func (ClassicAlertSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClassicAlertSettings)(nil)).Elem()
}

func (o ClassicAlertSettingsPtrOutput) ToClassicAlertSettingsPtrOutput() ClassicAlertSettingsPtrOutput {
	return o
}

func (o ClassicAlertSettingsPtrOutput) ToClassicAlertSettingsPtrOutputWithContext(ctx context.Context) ClassicAlertSettingsPtrOutput {
	return o
}

func (o ClassicAlertSettingsPtrOutput) Elem() ClassicAlertSettingsOutput {
	return o.ApplyT(func(v *ClassicAlertSettings) ClassicAlertSettings {
		if v != nil {
			return *v
		}
		var ret ClassicAlertSettings
		return ret
	}).(ClassicAlertSettingsOutput)
}

func (o ClassicAlertSettingsPtrOutput) AlertsForCriticalOperations() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClassicAlertSettings) *string {
		if v == nil {
			return nil
		}
		return v.AlertsForCriticalOperations
	}).(pulumi.StringPtrOutput)
}

type ClassicAlertSettingsResponse struct {
	AlertsForCriticalOperations *string `pulumi:"alertsForCriticalOperations"`
}

type ClassicAlertSettingsResponseOutput struct{ *pulumi.OutputState }

func (ClassicAlertSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClassicAlertSettingsResponse)(nil)).Elem()
}

func (o ClassicAlertSettingsResponseOutput) ToClassicAlertSettingsResponseOutput() ClassicAlertSettingsResponseOutput {
	return o
}

func (o ClassicAlertSettingsResponseOutput) ToClassicAlertSettingsResponseOutputWithContext(ctx context.Context) ClassicAlertSettingsResponseOutput {
	return o
}

func (o ClassicAlertSettingsResponseOutput) AlertsForCriticalOperations() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClassicAlertSettingsResponse) *string { return v.AlertsForCriticalOperations }).(pulumi.StringPtrOutput)
}

type ClassicAlertSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ClassicAlertSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClassicAlertSettingsResponse)(nil)).Elem()
}

func (o ClassicAlertSettingsResponsePtrOutput) ToClassicAlertSettingsResponsePtrOutput() ClassicAlertSettingsResponsePtrOutput {
	return o
}

func (o ClassicAlertSettingsResponsePtrOutput) ToClassicAlertSettingsResponsePtrOutputWithContext(ctx context.Context) ClassicAlertSettingsResponsePtrOutput {
	return o
}

func (o ClassicAlertSettingsResponsePtrOutput) Elem() ClassicAlertSettingsResponseOutput {
	return o.ApplyT(func(v *ClassicAlertSettingsResponse) ClassicAlertSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ClassicAlertSettingsResponse
		return ret
	}).(ClassicAlertSettingsResponseOutput)
}

func (o ClassicAlertSettingsResponsePtrOutput) AlertsForCriticalOperations() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClassicAlertSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AlertsForCriticalOperations
	}).(pulumi.StringPtrOutput)
}

type CmkKekIdentity struct {
	UseSystemAssignedIdentity *bool   `pulumi:"useSystemAssignedIdentity"`
	UserAssignedIdentity      *string `pulumi:"userAssignedIdentity"`
}





type CmkKekIdentityInput interface {
	pulumi.Input

	ToCmkKekIdentityOutput() CmkKekIdentityOutput
	ToCmkKekIdentityOutputWithContext(context.Context) CmkKekIdentityOutput
}

type CmkKekIdentityArgs struct {
	UseSystemAssignedIdentity pulumi.BoolPtrInput   `pulumi:"useSystemAssignedIdentity"`
	UserAssignedIdentity      pulumi.StringPtrInput `pulumi:"userAssignedIdentity"`
}

func (CmkKekIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CmkKekIdentity)(nil)).Elem()
}

func (i CmkKekIdentityArgs) ToCmkKekIdentityOutput() CmkKekIdentityOutput {
	return i.ToCmkKekIdentityOutputWithContext(context.Background())
}

func (i CmkKekIdentityArgs) ToCmkKekIdentityOutputWithContext(ctx context.Context) CmkKekIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CmkKekIdentityOutput)
}

func (i CmkKekIdentityArgs) ToCmkKekIdentityPtrOutput() CmkKekIdentityPtrOutput {
	return i.ToCmkKekIdentityPtrOutputWithContext(context.Background())
}

func (i CmkKekIdentityArgs) ToCmkKekIdentityPtrOutputWithContext(ctx context.Context) CmkKekIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CmkKekIdentityOutput).ToCmkKekIdentityPtrOutputWithContext(ctx)
}









type CmkKekIdentityPtrInput interface {
	pulumi.Input

	ToCmkKekIdentityPtrOutput() CmkKekIdentityPtrOutput
	ToCmkKekIdentityPtrOutputWithContext(context.Context) CmkKekIdentityPtrOutput
}

type cmkKekIdentityPtrType CmkKekIdentityArgs

func CmkKekIdentityPtr(v *CmkKekIdentityArgs) CmkKekIdentityPtrInput {
	return (*cmkKekIdentityPtrType)(v)
}

func (*cmkKekIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CmkKekIdentity)(nil)).Elem()
}

func (i *cmkKekIdentityPtrType) ToCmkKekIdentityPtrOutput() CmkKekIdentityPtrOutput {
	return i.ToCmkKekIdentityPtrOutputWithContext(context.Background())
}

func (i *cmkKekIdentityPtrType) ToCmkKekIdentityPtrOutputWithContext(ctx context.Context) CmkKekIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CmkKekIdentityPtrOutput)
}

type CmkKekIdentityOutput struct{ *pulumi.OutputState }

func (CmkKekIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CmkKekIdentity)(nil)).Elem()
}

func (o CmkKekIdentityOutput) ToCmkKekIdentityOutput() CmkKekIdentityOutput {
	return o
}

func (o CmkKekIdentityOutput) ToCmkKekIdentityOutputWithContext(ctx context.Context) CmkKekIdentityOutput {
	return o
}

func (o CmkKekIdentityOutput) ToCmkKekIdentityPtrOutput() CmkKekIdentityPtrOutput {
	return o.ToCmkKekIdentityPtrOutputWithContext(context.Background())
}

func (o CmkKekIdentityOutput) ToCmkKekIdentityPtrOutputWithContext(ctx context.Context) CmkKekIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CmkKekIdentity) *CmkKekIdentity {
		return &v
	}).(CmkKekIdentityPtrOutput)
}

func (o CmkKekIdentityOutput) UseSystemAssignedIdentity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CmkKekIdentity) *bool { return v.UseSystemAssignedIdentity }).(pulumi.BoolPtrOutput)
}

func (o CmkKekIdentityOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CmkKekIdentity) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type CmkKekIdentityPtrOutput struct{ *pulumi.OutputState }

func (CmkKekIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CmkKekIdentity)(nil)).Elem()
}

func (o CmkKekIdentityPtrOutput) ToCmkKekIdentityPtrOutput() CmkKekIdentityPtrOutput {
	return o
}

func (o CmkKekIdentityPtrOutput) ToCmkKekIdentityPtrOutputWithContext(ctx context.Context) CmkKekIdentityPtrOutput {
	return o
}

func (o CmkKekIdentityPtrOutput) Elem() CmkKekIdentityOutput {
	return o.ApplyT(func(v *CmkKekIdentity) CmkKekIdentity {
		if v != nil {
			return *v
		}
		var ret CmkKekIdentity
		return ret
	}).(CmkKekIdentityOutput)
}

func (o CmkKekIdentityPtrOutput) UseSystemAssignedIdentity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CmkKekIdentity) *bool {
		if v == nil {
			return nil
		}
		return v.UseSystemAssignedIdentity
	}).(pulumi.BoolPtrOutput)
}

func (o CmkKekIdentityPtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CmkKekIdentity) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type CmkKekIdentityResponse struct {
	UseSystemAssignedIdentity *bool   `pulumi:"useSystemAssignedIdentity"`
	UserAssignedIdentity      *string `pulumi:"userAssignedIdentity"`
}

type CmkKekIdentityResponseOutput struct{ *pulumi.OutputState }

func (CmkKekIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CmkKekIdentityResponse)(nil)).Elem()
}

func (o CmkKekIdentityResponseOutput) ToCmkKekIdentityResponseOutput() CmkKekIdentityResponseOutput {
	return o
}

func (o CmkKekIdentityResponseOutput) ToCmkKekIdentityResponseOutputWithContext(ctx context.Context) CmkKekIdentityResponseOutput {
	return o
}

func (o CmkKekIdentityResponseOutput) UseSystemAssignedIdentity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CmkKekIdentityResponse) *bool { return v.UseSystemAssignedIdentity }).(pulumi.BoolPtrOutput)
}

func (o CmkKekIdentityResponseOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CmkKekIdentityResponse) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type CmkKekIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (CmkKekIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CmkKekIdentityResponse)(nil)).Elem()
}

func (o CmkKekIdentityResponsePtrOutput) ToCmkKekIdentityResponsePtrOutput() CmkKekIdentityResponsePtrOutput {
	return o
}

func (o CmkKekIdentityResponsePtrOutput) ToCmkKekIdentityResponsePtrOutputWithContext(ctx context.Context) CmkKekIdentityResponsePtrOutput {
	return o
}

func (o CmkKekIdentityResponsePtrOutput) Elem() CmkKekIdentityResponseOutput {
	return o.ApplyT(func(v *CmkKekIdentityResponse) CmkKekIdentityResponse {
		if v != nil {
			return *v
		}
		var ret CmkKekIdentityResponse
		return ret
	}).(CmkKekIdentityResponseOutput)
}

func (o CmkKekIdentityResponsePtrOutput) UseSystemAssignedIdentity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CmkKekIdentityResponse) *bool {
		if v == nil {
			return nil
		}
		return v.UseSystemAssignedIdentity
	}).(pulumi.BoolPtrOutput)
}

func (o CmkKekIdentityResponsePtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CmkKekIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type CmkKeyVaultProperties struct {
	KeyUri *string `pulumi:"keyUri"`
}





type CmkKeyVaultPropertiesInput interface {
	pulumi.Input

	ToCmkKeyVaultPropertiesOutput() CmkKeyVaultPropertiesOutput
	ToCmkKeyVaultPropertiesOutputWithContext(context.Context) CmkKeyVaultPropertiesOutput
}

type CmkKeyVaultPropertiesArgs struct {
	KeyUri pulumi.StringPtrInput `pulumi:"keyUri"`
}

func (CmkKeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CmkKeyVaultProperties)(nil)).Elem()
}

func (i CmkKeyVaultPropertiesArgs) ToCmkKeyVaultPropertiesOutput() CmkKeyVaultPropertiesOutput {
	return i.ToCmkKeyVaultPropertiesOutputWithContext(context.Background())
}

func (i CmkKeyVaultPropertiesArgs) ToCmkKeyVaultPropertiesOutputWithContext(ctx context.Context) CmkKeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CmkKeyVaultPropertiesOutput)
}

func (i CmkKeyVaultPropertiesArgs) ToCmkKeyVaultPropertiesPtrOutput() CmkKeyVaultPropertiesPtrOutput {
	return i.ToCmkKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i CmkKeyVaultPropertiesArgs) ToCmkKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) CmkKeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CmkKeyVaultPropertiesOutput).ToCmkKeyVaultPropertiesPtrOutputWithContext(ctx)
}









type CmkKeyVaultPropertiesPtrInput interface {
	pulumi.Input

	ToCmkKeyVaultPropertiesPtrOutput() CmkKeyVaultPropertiesPtrOutput
	ToCmkKeyVaultPropertiesPtrOutputWithContext(context.Context) CmkKeyVaultPropertiesPtrOutput
}

type cmkKeyVaultPropertiesPtrType CmkKeyVaultPropertiesArgs

func CmkKeyVaultPropertiesPtr(v *CmkKeyVaultPropertiesArgs) CmkKeyVaultPropertiesPtrInput {
	return (*cmkKeyVaultPropertiesPtrType)(v)
}

func (*cmkKeyVaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CmkKeyVaultProperties)(nil)).Elem()
}

func (i *cmkKeyVaultPropertiesPtrType) ToCmkKeyVaultPropertiesPtrOutput() CmkKeyVaultPropertiesPtrOutput {
	return i.ToCmkKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *cmkKeyVaultPropertiesPtrType) ToCmkKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) CmkKeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CmkKeyVaultPropertiesPtrOutput)
}

type CmkKeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (CmkKeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CmkKeyVaultProperties)(nil)).Elem()
}

func (o CmkKeyVaultPropertiesOutput) ToCmkKeyVaultPropertiesOutput() CmkKeyVaultPropertiesOutput {
	return o
}

func (o CmkKeyVaultPropertiesOutput) ToCmkKeyVaultPropertiesOutputWithContext(ctx context.Context) CmkKeyVaultPropertiesOutput {
	return o
}

func (o CmkKeyVaultPropertiesOutput) ToCmkKeyVaultPropertiesPtrOutput() CmkKeyVaultPropertiesPtrOutput {
	return o.ToCmkKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o CmkKeyVaultPropertiesOutput) ToCmkKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) CmkKeyVaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CmkKeyVaultProperties) *CmkKeyVaultProperties {
		return &v
	}).(CmkKeyVaultPropertiesPtrOutput)
}

func (o CmkKeyVaultPropertiesOutput) KeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CmkKeyVaultProperties) *string { return v.KeyUri }).(pulumi.StringPtrOutput)
}

type CmkKeyVaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CmkKeyVaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CmkKeyVaultProperties)(nil)).Elem()
}

func (o CmkKeyVaultPropertiesPtrOutput) ToCmkKeyVaultPropertiesPtrOutput() CmkKeyVaultPropertiesPtrOutput {
	return o
}

func (o CmkKeyVaultPropertiesPtrOutput) ToCmkKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) CmkKeyVaultPropertiesPtrOutput {
	return o
}

func (o CmkKeyVaultPropertiesPtrOutput) Elem() CmkKeyVaultPropertiesOutput {
	return o.ApplyT(func(v *CmkKeyVaultProperties) CmkKeyVaultProperties {
		if v != nil {
			return *v
		}
		var ret CmkKeyVaultProperties
		return ret
	}).(CmkKeyVaultPropertiesOutput)
}

func (o CmkKeyVaultPropertiesPtrOutput) KeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CmkKeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyUri
	}).(pulumi.StringPtrOutput)
}

type CmkKeyVaultPropertiesResponse struct {
	KeyUri *string `pulumi:"keyUri"`
}

type CmkKeyVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CmkKeyVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CmkKeyVaultPropertiesResponse)(nil)).Elem()
}

func (o CmkKeyVaultPropertiesResponseOutput) ToCmkKeyVaultPropertiesResponseOutput() CmkKeyVaultPropertiesResponseOutput {
	return o
}

func (o CmkKeyVaultPropertiesResponseOutput) ToCmkKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) CmkKeyVaultPropertiesResponseOutput {
	return o
}

func (o CmkKeyVaultPropertiesResponseOutput) KeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CmkKeyVaultPropertiesResponse) *string { return v.KeyUri }).(pulumi.StringPtrOutput)
}

type CmkKeyVaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CmkKeyVaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CmkKeyVaultPropertiesResponse)(nil)).Elem()
}

func (o CmkKeyVaultPropertiesResponsePtrOutput) ToCmkKeyVaultPropertiesResponsePtrOutput() CmkKeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o CmkKeyVaultPropertiesResponsePtrOutput) ToCmkKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) CmkKeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o CmkKeyVaultPropertiesResponsePtrOutput) Elem() CmkKeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v *CmkKeyVaultPropertiesResponse) CmkKeyVaultPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CmkKeyVaultPropertiesResponse
		return ret
	}).(CmkKeyVaultPropertiesResponseOutput)
}

func (o CmkKeyVaultPropertiesResponsePtrOutput) KeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CmkKeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyUri
	}).(pulumi.StringPtrOutput)
}

type IdentityData struct {
	Type                   string                 `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type IdentityDataInput interface {
	pulumi.Input

	ToIdentityDataOutput() IdentityDataOutput
	ToIdentityDataOutputWithContext(context.Context) IdentityDataOutput
}

type IdentityDataArgs struct {
	Type                   pulumi.StringInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput    `pulumi:"userAssignedIdentities"`
}

func (IdentityDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityData)(nil)).Elem()
}

func (i IdentityDataArgs) ToIdentityDataOutput() IdentityDataOutput {
	return i.ToIdentityDataOutputWithContext(context.Background())
}

func (i IdentityDataArgs) ToIdentityDataOutputWithContext(ctx context.Context) IdentityDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityDataOutput)
}

func (i IdentityDataArgs) ToIdentityDataPtrOutput() IdentityDataPtrOutput {
	return i.ToIdentityDataPtrOutputWithContext(context.Background())
}

func (i IdentityDataArgs) ToIdentityDataPtrOutputWithContext(ctx context.Context) IdentityDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityDataOutput).ToIdentityDataPtrOutputWithContext(ctx)
}









type IdentityDataPtrInput interface {
	pulumi.Input

	ToIdentityDataPtrOutput() IdentityDataPtrOutput
	ToIdentityDataPtrOutputWithContext(context.Context) IdentityDataPtrOutput
}

type identityDataPtrType IdentityDataArgs

func IdentityDataPtr(v *IdentityDataArgs) IdentityDataPtrInput {
	return (*identityDataPtrType)(v)
}

func (*identityDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityData)(nil)).Elem()
}

func (i *identityDataPtrType) ToIdentityDataPtrOutput() IdentityDataPtrOutput {
	return i.ToIdentityDataPtrOutputWithContext(context.Background())
}

func (i *identityDataPtrType) ToIdentityDataPtrOutputWithContext(ctx context.Context) IdentityDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityDataPtrOutput)
}

type IdentityDataOutput struct{ *pulumi.OutputState }

func (IdentityDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityData)(nil)).Elem()
}

func (o IdentityDataOutput) ToIdentityDataOutput() IdentityDataOutput {
	return o
}

func (o IdentityDataOutput) ToIdentityDataOutputWithContext(ctx context.Context) IdentityDataOutput {
	return o
}

func (o IdentityDataOutput) ToIdentityDataPtrOutput() IdentityDataPtrOutput {
	return o.ToIdentityDataPtrOutputWithContext(context.Background())
}

func (o IdentityDataOutput) ToIdentityDataPtrOutputWithContext(ctx context.Context) IdentityDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityData) *IdentityData {
		return &v
	}).(IdentityDataPtrOutput)
}

func (o IdentityDataOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityData) string { return v.Type }).(pulumi.StringOutput)
}

func (o IdentityDataOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v IdentityData) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type IdentityDataPtrOutput struct{ *pulumi.OutputState }

func (IdentityDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityData)(nil)).Elem()
}

func (o IdentityDataPtrOutput) ToIdentityDataPtrOutput() IdentityDataPtrOutput {
	return o
}

func (o IdentityDataPtrOutput) ToIdentityDataPtrOutputWithContext(ctx context.Context) IdentityDataPtrOutput {
	return o
}

func (o IdentityDataPtrOutput) Elem() IdentityDataOutput {
	return o.ApplyT(func(v *IdentityData) IdentityData {
		if v != nil {
			return *v
		}
		var ret IdentityData
		return ret
	}).(IdentityDataOutput)
}

func (o IdentityDataPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityData) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityDataPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *IdentityData) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type IdentityDataResponse struct {
	PrincipalId            string                          `pulumi:"principalId"`
	TenantId               string                          `pulumi:"tenantId"`
	Type                   string                          `pulumi:"type"`
	UserAssignedIdentities map[string]UserIdentityResponse `pulumi:"userAssignedIdentities"`
}

type IdentityDataResponseOutput struct{ *pulumi.OutputState }

func (IdentityDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityDataResponse)(nil)).Elem()
}

func (o IdentityDataResponseOutput) ToIdentityDataResponseOutput() IdentityDataResponseOutput {
	return o
}

func (o IdentityDataResponseOutput) ToIdentityDataResponseOutputWithContext(ctx context.Context) IdentityDataResponseOutput {
	return o
}

func (o IdentityDataResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityDataResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityDataResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityDataResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityDataResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityDataResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o IdentityDataResponseOutput) UserAssignedIdentities() UserIdentityResponseMapOutput {
	return o.ApplyT(func(v IdentityDataResponse) map[string]UserIdentityResponse { return v.UserAssignedIdentities }).(UserIdentityResponseMapOutput)
}

type IdentityDataResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityDataResponse)(nil)).Elem()
}

func (o IdentityDataResponsePtrOutput) ToIdentityDataResponsePtrOutput() IdentityDataResponsePtrOutput {
	return o
}

func (o IdentityDataResponsePtrOutput) ToIdentityDataResponsePtrOutputWithContext(ctx context.Context) IdentityDataResponsePtrOutput {
	return o
}

func (o IdentityDataResponsePtrOutput) Elem() IdentityDataResponseOutput {
	return o.ApplyT(func(v *IdentityDataResponse) IdentityDataResponse {
		if v != nil {
			return *v
		}
		var ret IdentityDataResponse
		return ret
	}).(IdentityDataResponseOutput)
}

func (o IdentityDataResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityDataResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityDataResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityDataResponsePtrOutput) UserAssignedIdentities() UserIdentityResponseMapOutput {
	return o.ApplyT(func(v *IdentityDataResponse) map[string]UserIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserIdentityResponseMapOutput)
}

type ImmutabilitySettings struct {
	State *string `pulumi:"state"`
}





type ImmutabilitySettingsInput interface {
	pulumi.Input

	ToImmutabilitySettingsOutput() ImmutabilitySettingsOutput
	ToImmutabilitySettingsOutputWithContext(context.Context) ImmutabilitySettingsOutput
}

type ImmutabilitySettingsArgs struct {
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (ImmutabilitySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImmutabilitySettings)(nil)).Elem()
}

func (i ImmutabilitySettingsArgs) ToImmutabilitySettingsOutput() ImmutabilitySettingsOutput {
	return i.ToImmutabilitySettingsOutputWithContext(context.Background())
}

func (i ImmutabilitySettingsArgs) ToImmutabilitySettingsOutputWithContext(ctx context.Context) ImmutabilitySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImmutabilitySettingsOutput)
}

func (i ImmutabilitySettingsArgs) ToImmutabilitySettingsPtrOutput() ImmutabilitySettingsPtrOutput {
	return i.ToImmutabilitySettingsPtrOutputWithContext(context.Background())
}

func (i ImmutabilitySettingsArgs) ToImmutabilitySettingsPtrOutputWithContext(ctx context.Context) ImmutabilitySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImmutabilitySettingsOutput).ToImmutabilitySettingsPtrOutputWithContext(ctx)
}









type ImmutabilitySettingsPtrInput interface {
	pulumi.Input

	ToImmutabilitySettingsPtrOutput() ImmutabilitySettingsPtrOutput
	ToImmutabilitySettingsPtrOutputWithContext(context.Context) ImmutabilitySettingsPtrOutput
}

type immutabilitySettingsPtrType ImmutabilitySettingsArgs

func ImmutabilitySettingsPtr(v *ImmutabilitySettingsArgs) ImmutabilitySettingsPtrInput {
	return (*immutabilitySettingsPtrType)(v)
}

func (*immutabilitySettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImmutabilitySettings)(nil)).Elem()
}

func (i *immutabilitySettingsPtrType) ToImmutabilitySettingsPtrOutput() ImmutabilitySettingsPtrOutput {
	return i.ToImmutabilitySettingsPtrOutputWithContext(context.Background())
}

func (i *immutabilitySettingsPtrType) ToImmutabilitySettingsPtrOutputWithContext(ctx context.Context) ImmutabilitySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImmutabilitySettingsPtrOutput)
}

type ImmutabilitySettingsOutput struct{ *pulumi.OutputState }

func (ImmutabilitySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImmutabilitySettings)(nil)).Elem()
}

func (o ImmutabilitySettingsOutput) ToImmutabilitySettingsOutput() ImmutabilitySettingsOutput {
	return o
}

func (o ImmutabilitySettingsOutput) ToImmutabilitySettingsOutputWithContext(ctx context.Context) ImmutabilitySettingsOutput {
	return o
}

func (o ImmutabilitySettingsOutput) ToImmutabilitySettingsPtrOutput() ImmutabilitySettingsPtrOutput {
	return o.ToImmutabilitySettingsPtrOutputWithContext(context.Background())
}

func (o ImmutabilitySettingsOutput) ToImmutabilitySettingsPtrOutputWithContext(ctx context.Context) ImmutabilitySettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImmutabilitySettings) *ImmutabilitySettings {
		return &v
	}).(ImmutabilitySettingsPtrOutput)
}

func (o ImmutabilitySettingsOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImmutabilitySettings) *string { return v.State }).(pulumi.StringPtrOutput)
}

type ImmutabilitySettingsPtrOutput struct{ *pulumi.OutputState }

func (ImmutabilitySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImmutabilitySettings)(nil)).Elem()
}

func (o ImmutabilitySettingsPtrOutput) ToImmutabilitySettingsPtrOutput() ImmutabilitySettingsPtrOutput {
	return o
}

func (o ImmutabilitySettingsPtrOutput) ToImmutabilitySettingsPtrOutputWithContext(ctx context.Context) ImmutabilitySettingsPtrOutput {
	return o
}

func (o ImmutabilitySettingsPtrOutput) Elem() ImmutabilitySettingsOutput {
	return o.ApplyT(func(v *ImmutabilitySettings) ImmutabilitySettings {
		if v != nil {
			return *v
		}
		var ret ImmutabilitySettings
		return ret
	}).(ImmutabilitySettingsOutput)
}

func (o ImmutabilitySettingsPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImmutabilitySettings) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type ImmutabilitySettingsResponse struct {
	State *string `pulumi:"state"`
}

type ImmutabilitySettingsResponseOutput struct{ *pulumi.OutputState }

func (ImmutabilitySettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImmutabilitySettingsResponse)(nil)).Elem()
}

func (o ImmutabilitySettingsResponseOutput) ToImmutabilitySettingsResponseOutput() ImmutabilitySettingsResponseOutput {
	return o
}

func (o ImmutabilitySettingsResponseOutput) ToImmutabilitySettingsResponseOutputWithContext(ctx context.Context) ImmutabilitySettingsResponseOutput {
	return o
}

func (o ImmutabilitySettingsResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImmutabilitySettingsResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type ImmutabilitySettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ImmutabilitySettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImmutabilitySettingsResponse)(nil)).Elem()
}

func (o ImmutabilitySettingsResponsePtrOutput) ToImmutabilitySettingsResponsePtrOutput() ImmutabilitySettingsResponsePtrOutput {
	return o
}

func (o ImmutabilitySettingsResponsePtrOutput) ToImmutabilitySettingsResponsePtrOutputWithContext(ctx context.Context) ImmutabilitySettingsResponsePtrOutput {
	return o
}

func (o ImmutabilitySettingsResponsePtrOutput) Elem() ImmutabilitySettingsResponseOutput {
	return o.ApplyT(func(v *ImmutabilitySettingsResponse) ImmutabilitySettingsResponse {
		if v != nil {
			return *v
		}
		var ret ImmutabilitySettingsResponse
		return ret
	}).(ImmutabilitySettingsResponseOutput)
}

func (o ImmutabilitySettingsResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImmutabilitySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type MonitoringSettings struct {
	AzureMonitorAlertSettings *AzureMonitorAlertSettings `pulumi:"azureMonitorAlertSettings"`
	ClassicAlertSettings      *ClassicAlertSettings      `pulumi:"classicAlertSettings"`
}





type MonitoringSettingsInput interface {
	pulumi.Input

	ToMonitoringSettingsOutput() MonitoringSettingsOutput
	ToMonitoringSettingsOutputWithContext(context.Context) MonitoringSettingsOutput
}

type MonitoringSettingsArgs struct {
	AzureMonitorAlertSettings AzureMonitorAlertSettingsPtrInput `pulumi:"azureMonitorAlertSettings"`
	ClassicAlertSettings      ClassicAlertSettingsPtrInput      `pulumi:"classicAlertSettings"`
}

func (MonitoringSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringSettings)(nil)).Elem()
}

func (i MonitoringSettingsArgs) ToMonitoringSettingsOutput() MonitoringSettingsOutput {
	return i.ToMonitoringSettingsOutputWithContext(context.Background())
}

func (i MonitoringSettingsArgs) ToMonitoringSettingsOutputWithContext(ctx context.Context) MonitoringSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringSettingsOutput)
}

func (i MonitoringSettingsArgs) ToMonitoringSettingsPtrOutput() MonitoringSettingsPtrOutput {
	return i.ToMonitoringSettingsPtrOutputWithContext(context.Background())
}

func (i MonitoringSettingsArgs) ToMonitoringSettingsPtrOutputWithContext(ctx context.Context) MonitoringSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringSettingsOutput).ToMonitoringSettingsPtrOutputWithContext(ctx)
}









type MonitoringSettingsPtrInput interface {
	pulumi.Input

	ToMonitoringSettingsPtrOutput() MonitoringSettingsPtrOutput
	ToMonitoringSettingsPtrOutputWithContext(context.Context) MonitoringSettingsPtrOutput
}

type monitoringSettingsPtrType MonitoringSettingsArgs

func MonitoringSettingsPtr(v *MonitoringSettingsArgs) MonitoringSettingsPtrInput {
	return (*monitoringSettingsPtrType)(v)
}

func (*monitoringSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringSettings)(nil)).Elem()
}

func (i *monitoringSettingsPtrType) ToMonitoringSettingsPtrOutput() MonitoringSettingsPtrOutput {
	return i.ToMonitoringSettingsPtrOutputWithContext(context.Background())
}

func (i *monitoringSettingsPtrType) ToMonitoringSettingsPtrOutputWithContext(ctx context.Context) MonitoringSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringSettingsPtrOutput)
}

type MonitoringSettingsOutput struct{ *pulumi.OutputState }

func (MonitoringSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringSettings)(nil)).Elem()
}

func (o MonitoringSettingsOutput) ToMonitoringSettingsOutput() MonitoringSettingsOutput {
	return o
}

func (o MonitoringSettingsOutput) ToMonitoringSettingsOutputWithContext(ctx context.Context) MonitoringSettingsOutput {
	return o
}

func (o MonitoringSettingsOutput) ToMonitoringSettingsPtrOutput() MonitoringSettingsPtrOutput {
	return o.ToMonitoringSettingsPtrOutputWithContext(context.Background())
}

func (o MonitoringSettingsOutput) ToMonitoringSettingsPtrOutputWithContext(ctx context.Context) MonitoringSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitoringSettings) *MonitoringSettings {
		return &v
	}).(MonitoringSettingsPtrOutput)
}

func (o MonitoringSettingsOutput) AzureMonitorAlertSettings() AzureMonitorAlertSettingsPtrOutput {
	return o.ApplyT(func(v MonitoringSettings) *AzureMonitorAlertSettings { return v.AzureMonitorAlertSettings }).(AzureMonitorAlertSettingsPtrOutput)
}

func (o MonitoringSettingsOutput) ClassicAlertSettings() ClassicAlertSettingsPtrOutput {
	return o.ApplyT(func(v MonitoringSettings) *ClassicAlertSettings { return v.ClassicAlertSettings }).(ClassicAlertSettingsPtrOutput)
}

type MonitoringSettingsPtrOutput struct{ *pulumi.OutputState }

func (MonitoringSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringSettings)(nil)).Elem()
}

func (o MonitoringSettingsPtrOutput) ToMonitoringSettingsPtrOutput() MonitoringSettingsPtrOutput {
	return o
}

func (o MonitoringSettingsPtrOutput) ToMonitoringSettingsPtrOutputWithContext(ctx context.Context) MonitoringSettingsPtrOutput {
	return o
}

func (o MonitoringSettingsPtrOutput) Elem() MonitoringSettingsOutput {
	return o.ApplyT(func(v *MonitoringSettings) MonitoringSettings {
		if v != nil {
			return *v
		}
		var ret MonitoringSettings
		return ret
	}).(MonitoringSettingsOutput)
}

func (o MonitoringSettingsPtrOutput) AzureMonitorAlertSettings() AzureMonitorAlertSettingsPtrOutput {
	return o.ApplyT(func(v *MonitoringSettings) *AzureMonitorAlertSettings {
		if v == nil {
			return nil
		}
		return v.AzureMonitorAlertSettings
	}).(AzureMonitorAlertSettingsPtrOutput)
}

func (o MonitoringSettingsPtrOutput) ClassicAlertSettings() ClassicAlertSettingsPtrOutput {
	return o.ApplyT(func(v *MonitoringSettings) *ClassicAlertSettings {
		if v == nil {
			return nil
		}
		return v.ClassicAlertSettings
	}).(ClassicAlertSettingsPtrOutput)
}

type MonitoringSettingsResponse struct {
	AzureMonitorAlertSettings *AzureMonitorAlertSettingsResponse `pulumi:"azureMonitorAlertSettings"`
	ClassicAlertSettings      *ClassicAlertSettingsResponse      `pulumi:"classicAlertSettings"`
}

type MonitoringSettingsResponseOutput struct{ *pulumi.OutputState }

func (MonitoringSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringSettingsResponse)(nil)).Elem()
}

func (o MonitoringSettingsResponseOutput) ToMonitoringSettingsResponseOutput() MonitoringSettingsResponseOutput {
	return o
}

func (o MonitoringSettingsResponseOutput) ToMonitoringSettingsResponseOutputWithContext(ctx context.Context) MonitoringSettingsResponseOutput {
	return o
}

func (o MonitoringSettingsResponseOutput) AzureMonitorAlertSettings() AzureMonitorAlertSettingsResponsePtrOutput {
	return o.ApplyT(func(v MonitoringSettingsResponse) *AzureMonitorAlertSettingsResponse {
		return v.AzureMonitorAlertSettings
	}).(AzureMonitorAlertSettingsResponsePtrOutput)
}

func (o MonitoringSettingsResponseOutput) ClassicAlertSettings() ClassicAlertSettingsResponsePtrOutput {
	return o.ApplyT(func(v MonitoringSettingsResponse) *ClassicAlertSettingsResponse { return v.ClassicAlertSettings }).(ClassicAlertSettingsResponsePtrOutput)
}

type MonitoringSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (MonitoringSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringSettingsResponse)(nil)).Elem()
}

func (o MonitoringSettingsResponsePtrOutput) ToMonitoringSettingsResponsePtrOutput() MonitoringSettingsResponsePtrOutput {
	return o
}

func (o MonitoringSettingsResponsePtrOutput) ToMonitoringSettingsResponsePtrOutputWithContext(ctx context.Context) MonitoringSettingsResponsePtrOutput {
	return o
}

func (o MonitoringSettingsResponsePtrOutput) Elem() MonitoringSettingsResponseOutput {
	return o.ApplyT(func(v *MonitoringSettingsResponse) MonitoringSettingsResponse {
		if v != nil {
			return *v
		}
		var ret MonitoringSettingsResponse
		return ret
	}).(MonitoringSettingsResponseOutput)
}

func (o MonitoringSettingsResponsePtrOutput) AzureMonitorAlertSettings() AzureMonitorAlertSettingsResponsePtrOutput {
	return o.ApplyT(func(v *MonitoringSettingsResponse) *AzureMonitorAlertSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AzureMonitorAlertSettings
	}).(AzureMonitorAlertSettingsResponsePtrOutput)
}

func (o MonitoringSettingsResponsePtrOutput) ClassicAlertSettings() ClassicAlertSettingsResponsePtrOutput {
	return o.ApplyT(func(v *MonitoringSettingsResponse) *ClassicAlertSettingsResponse {
		if v == nil {
			return nil
		}
		return v.ClassicAlertSettings
	}).(ClassicAlertSettingsResponsePtrOutput)
}

type PrivateEndpointConnectionVaultPropertiesResponse struct {
	Id         string                                 `pulumi:"id"`
	Location   string                                 `pulumi:"location"`
	Name       string                                 `pulumi:"name"`
	Properties VaultPrivateEndpointConnectionResponse `pulumi:"properties"`
	Type       string                                 `pulumi:"type"`
}

type PrivateEndpointConnectionVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionVaultPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionVaultPropertiesResponseOutput) ToPrivateEndpointConnectionVaultPropertiesResponseOutput() PrivateEndpointConnectionVaultPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionVaultPropertiesResponseOutput) ToPrivateEndpointConnectionVaultPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionVaultPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionVaultPropertiesResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionVaultPropertiesResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionVaultPropertiesResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionVaultPropertiesResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionVaultPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionVaultPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionVaultPropertiesResponseOutput) Properties() VaultPrivateEndpointConnectionResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionVaultPropertiesResponse) VaultPrivateEndpointConnectionResponse {
		return v.Properties
	}).(VaultPrivateEndpointConnectionResponseOutput)
}

func (o PrivateEndpointConnectionVaultPropertiesResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionVaultPropertiesResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionVaultPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionVaultPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionVaultPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionVaultPropertiesResponseArrayOutput) ToPrivateEndpointConnectionVaultPropertiesResponseArrayOutput() PrivateEndpointConnectionVaultPropertiesResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionVaultPropertiesResponseArrayOutput) ToPrivateEndpointConnectionVaultPropertiesResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionVaultPropertiesResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionVaultPropertiesResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionVaultPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionVaultPropertiesResponse {
		return vs[0].([]PrivateEndpointConnectionVaultPropertiesResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionVaultPropertiesResponseOutput)
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

type SecuritySettings struct {
	ImmutabilitySettings *ImmutabilitySettings `pulumi:"immutabilitySettings"`
}





type SecuritySettingsInput interface {
	pulumi.Input

	ToSecuritySettingsOutput() SecuritySettingsOutput
	ToSecuritySettingsOutputWithContext(context.Context) SecuritySettingsOutput
}

type SecuritySettingsArgs struct {
	ImmutabilitySettings ImmutabilitySettingsPtrInput `pulumi:"immutabilitySettings"`
}

func (SecuritySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecuritySettings)(nil)).Elem()
}

func (i SecuritySettingsArgs) ToSecuritySettingsOutput() SecuritySettingsOutput {
	return i.ToSecuritySettingsOutputWithContext(context.Background())
}

func (i SecuritySettingsArgs) ToSecuritySettingsOutputWithContext(ctx context.Context) SecuritySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecuritySettingsOutput)
}

func (i SecuritySettingsArgs) ToSecuritySettingsPtrOutput() SecuritySettingsPtrOutput {
	return i.ToSecuritySettingsPtrOutputWithContext(context.Background())
}

func (i SecuritySettingsArgs) ToSecuritySettingsPtrOutputWithContext(ctx context.Context) SecuritySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecuritySettingsOutput).ToSecuritySettingsPtrOutputWithContext(ctx)
}









type SecuritySettingsPtrInput interface {
	pulumi.Input

	ToSecuritySettingsPtrOutput() SecuritySettingsPtrOutput
	ToSecuritySettingsPtrOutputWithContext(context.Context) SecuritySettingsPtrOutput
}

type securitySettingsPtrType SecuritySettingsArgs

func SecuritySettingsPtr(v *SecuritySettingsArgs) SecuritySettingsPtrInput {
	return (*securitySettingsPtrType)(v)
}

func (*securitySettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecuritySettings)(nil)).Elem()
}

func (i *securitySettingsPtrType) ToSecuritySettingsPtrOutput() SecuritySettingsPtrOutput {
	return i.ToSecuritySettingsPtrOutputWithContext(context.Background())
}

func (i *securitySettingsPtrType) ToSecuritySettingsPtrOutputWithContext(ctx context.Context) SecuritySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecuritySettingsPtrOutput)
}

type SecuritySettingsOutput struct{ *pulumi.OutputState }

func (SecuritySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecuritySettings)(nil)).Elem()
}

func (o SecuritySettingsOutput) ToSecuritySettingsOutput() SecuritySettingsOutput {
	return o
}

func (o SecuritySettingsOutput) ToSecuritySettingsOutputWithContext(ctx context.Context) SecuritySettingsOutput {
	return o
}

func (o SecuritySettingsOutput) ToSecuritySettingsPtrOutput() SecuritySettingsPtrOutput {
	return o.ToSecuritySettingsPtrOutputWithContext(context.Background())
}

func (o SecuritySettingsOutput) ToSecuritySettingsPtrOutputWithContext(ctx context.Context) SecuritySettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecuritySettings) *SecuritySettings {
		return &v
	}).(SecuritySettingsPtrOutput)
}

func (o SecuritySettingsOutput) ImmutabilitySettings() ImmutabilitySettingsPtrOutput {
	return o.ApplyT(func(v SecuritySettings) *ImmutabilitySettings { return v.ImmutabilitySettings }).(ImmutabilitySettingsPtrOutput)
}

type SecuritySettingsPtrOutput struct{ *pulumi.OutputState }

func (SecuritySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecuritySettings)(nil)).Elem()
}

func (o SecuritySettingsPtrOutput) ToSecuritySettingsPtrOutput() SecuritySettingsPtrOutput {
	return o
}

func (o SecuritySettingsPtrOutput) ToSecuritySettingsPtrOutputWithContext(ctx context.Context) SecuritySettingsPtrOutput {
	return o
}

func (o SecuritySettingsPtrOutput) Elem() SecuritySettingsOutput {
	return o.ApplyT(func(v *SecuritySettings) SecuritySettings {
		if v != nil {
			return *v
		}
		var ret SecuritySettings
		return ret
	}).(SecuritySettingsOutput)
}

func (o SecuritySettingsPtrOutput) ImmutabilitySettings() ImmutabilitySettingsPtrOutput {
	return o.ApplyT(func(v *SecuritySettings) *ImmutabilitySettings {
		if v == nil {
			return nil
		}
		return v.ImmutabilitySettings
	}).(ImmutabilitySettingsPtrOutput)
}

type SecuritySettingsResponse struct {
	ImmutabilitySettings *ImmutabilitySettingsResponse `pulumi:"immutabilitySettings"`
}

type SecuritySettingsResponseOutput struct{ *pulumi.OutputState }

func (SecuritySettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecuritySettingsResponse)(nil)).Elem()
}

func (o SecuritySettingsResponseOutput) ToSecuritySettingsResponseOutput() SecuritySettingsResponseOutput {
	return o
}

func (o SecuritySettingsResponseOutput) ToSecuritySettingsResponseOutputWithContext(ctx context.Context) SecuritySettingsResponseOutput {
	return o
}

func (o SecuritySettingsResponseOutput) ImmutabilitySettings() ImmutabilitySettingsResponsePtrOutput {
	return o.ApplyT(func(v SecuritySettingsResponse) *ImmutabilitySettingsResponse { return v.ImmutabilitySettings }).(ImmutabilitySettingsResponsePtrOutput)
}

type SecuritySettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (SecuritySettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecuritySettingsResponse)(nil)).Elem()
}

func (o SecuritySettingsResponsePtrOutput) ToSecuritySettingsResponsePtrOutput() SecuritySettingsResponsePtrOutput {
	return o
}

func (o SecuritySettingsResponsePtrOutput) ToSecuritySettingsResponsePtrOutputWithContext(ctx context.Context) SecuritySettingsResponsePtrOutput {
	return o
}

func (o SecuritySettingsResponsePtrOutput) Elem() SecuritySettingsResponseOutput {
	return o.ApplyT(func(v *SecuritySettingsResponse) SecuritySettingsResponse {
		if v != nil {
			return *v
		}
		var ret SecuritySettingsResponse
		return ret
	}).(SecuritySettingsResponseOutput)
}

func (o SecuritySettingsResponsePtrOutput) ImmutabilitySettings() ImmutabilitySettingsResponsePtrOutput {
	return o.ApplyT(func(v *SecuritySettingsResponse) *ImmutabilitySettingsResponse {
		if v == nil {
			return nil
		}
		return v.ImmutabilitySettings
	}).(ImmutabilitySettingsResponsePtrOutput)
}

type Sku struct {
	Capacity *string `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.StringPtrInput `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
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

func (o SkuOutput) Capacity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Capacity }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Size }).(pulumi.StringPtrOutput)
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

func (o SkuPtrOutput) Capacity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Size
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
	Capacity *string `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
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

func (o SkuResponseOutput) Capacity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Capacity }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
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

func (o SkuResponsePtrOutput) Capacity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
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

type UpgradeDetailsResponse struct {
	EndTimeUtc         string `pulumi:"endTimeUtc"`
	LastUpdatedTimeUtc string `pulumi:"lastUpdatedTimeUtc"`
	Message            string `pulumi:"message"`
	OperationId        string `pulumi:"operationId"`
	PreviousResourceId string `pulumi:"previousResourceId"`
	StartTimeUtc       string `pulumi:"startTimeUtc"`
	Status             string `pulumi:"status"`
	TriggerType        string `pulumi:"triggerType"`
	UpgradedResourceId string `pulumi:"upgradedResourceId"`
}

type UpgradeDetailsResponseOutput struct{ *pulumi.OutputState }

func (UpgradeDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpgradeDetailsResponse)(nil)).Elem()
}

func (o UpgradeDetailsResponseOutput) ToUpgradeDetailsResponseOutput() UpgradeDetailsResponseOutput {
	return o
}

func (o UpgradeDetailsResponseOutput) ToUpgradeDetailsResponseOutputWithContext(ctx context.Context) UpgradeDetailsResponseOutput {
	return o
}

func (o UpgradeDetailsResponseOutput) EndTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.EndTimeUtc }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) LastUpdatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.LastUpdatedTimeUtc }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) OperationId() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.OperationId }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) PreviousResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.PreviousResourceId }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) StartTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.StartTimeUtc }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) TriggerType() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.TriggerType }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) UpgradedResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.UpgradedResourceId }).(pulumi.StringOutput)
}

type UpgradeDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (UpgradeDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradeDetailsResponse)(nil)).Elem()
}

func (o UpgradeDetailsResponsePtrOutput) ToUpgradeDetailsResponsePtrOutput() UpgradeDetailsResponsePtrOutput {
	return o
}

func (o UpgradeDetailsResponsePtrOutput) ToUpgradeDetailsResponsePtrOutputWithContext(ctx context.Context) UpgradeDetailsResponsePtrOutput {
	return o
}

func (o UpgradeDetailsResponsePtrOutput) Elem() UpgradeDetailsResponseOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) UpgradeDetailsResponse {
		if v != nil {
			return *v
		}
		var ret UpgradeDetailsResponse
		return ret
	}).(UpgradeDetailsResponseOutput)
}

func (o UpgradeDetailsResponsePtrOutput) EndTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EndTimeUtc
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) LastUpdatedTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastUpdatedTimeUtc
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OperationId
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) PreviousResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PreviousResourceId
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) StartTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StartTimeUtc
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) TriggerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TriggerType
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) UpgradedResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpgradedResourceId
	}).(pulumi.StringPtrOutput)
}

type UserIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type UserIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityResponse)(nil)).Elem()
}

func (o UserIdentityResponseOutput) ToUserIdentityResponseOutput() UserIdentityResponseOutput {
	return o
}

func (o UserIdentityResponseOutput) ToUserIdentityResponseOutputWithContext(ctx context.Context) UserIdentityResponseOutput {
	return o
}

func (o UserIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityResponse)(nil)).Elem()
}

func (o UserIdentityResponseMapOutput) ToUserIdentityResponseMapOutput() UserIdentityResponseMapOutput {
	return o
}

func (o UserIdentityResponseMapOutput) ToUserIdentityResponseMapOutputWithContext(ctx context.Context) UserIdentityResponseMapOutput {
	return o
}

func (o UserIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserIdentityResponse {
		return vs[0].(map[string]UserIdentityResponse)[vs[1].(string)]
	}).(UserIdentityResponseOutput)
}

type VaultPrivateEndpointConnectionResponse struct {
	GroupIds                          []string                                       `pulumi:"groupIds"`
	PrivateEndpoint                   PrivateEndpointResponse                        `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState VaultPrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                         `pulumi:"provisioningState"`
}

type VaultPrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (VaultPrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o VaultPrivateEndpointConnectionResponseOutput) ToVaultPrivateEndpointConnectionResponseOutput() VaultPrivateEndpointConnectionResponseOutput {
	return o
}

func (o VaultPrivateEndpointConnectionResponseOutput) ToVaultPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) VaultPrivateEndpointConnectionResponseOutput {
	return o
}

func (o VaultPrivateEndpointConnectionResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VaultPrivateEndpointConnectionResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o VaultPrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v VaultPrivateEndpointConnectionResponse) PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponseOutput)
}

func (o VaultPrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() VaultPrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v VaultPrivateEndpointConnectionResponse) VaultPrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(VaultPrivateLinkServiceConnectionStateResponseOutput)
}

func (o VaultPrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type VaultPrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired string `pulumi:"actionsRequired"`
	Description     string `pulumi:"description"`
	Status          string `pulumi:"status"`
}

type VaultPrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (VaultPrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o VaultPrivateLinkServiceConnectionStateResponseOutput) ToVaultPrivateLinkServiceConnectionStateResponseOutput() VaultPrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o VaultPrivateLinkServiceConnectionStateResponseOutput) ToVaultPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) VaultPrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o VaultPrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPrivateLinkServiceConnectionStateResponse) string { return v.ActionsRequired }).(pulumi.StringOutput)
}

func (o VaultPrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPrivateLinkServiceConnectionStateResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o VaultPrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPrivateLinkServiceConnectionStateResponse) string { return v.Status }).(pulumi.StringOutput)
}

type VaultProperties struct {
	Encryption          *VaultPropertiesEncryption `pulumi:"encryption"`
	MonitoringSettings  *MonitoringSettings        `pulumi:"monitoringSettings"`
	PublicNetworkAccess *string                    `pulumi:"publicNetworkAccess"`
	SecuritySettings    *SecuritySettings          `pulumi:"securitySettings"`
}





type VaultPropertiesInput interface {
	pulumi.Input

	ToVaultPropertiesOutput() VaultPropertiesOutput
	ToVaultPropertiesOutputWithContext(context.Context) VaultPropertiesOutput
}

type VaultPropertiesArgs struct {
	Encryption          VaultPropertiesEncryptionPtrInput `pulumi:"encryption"`
	MonitoringSettings  MonitoringSettingsPtrInput        `pulumi:"monitoringSettings"`
	PublicNetworkAccess pulumi.StringPtrInput             `pulumi:"publicNetworkAccess"`
	SecuritySettings    SecuritySettingsPtrInput          `pulumi:"securitySettings"`
}

func (VaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultProperties)(nil)).Elem()
}

func (i VaultPropertiesArgs) ToVaultPropertiesOutput() VaultPropertiesOutput {
	return i.ToVaultPropertiesOutputWithContext(context.Background())
}

func (i VaultPropertiesArgs) ToVaultPropertiesOutputWithContext(ctx context.Context) VaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesOutput)
}

func (i VaultPropertiesArgs) ToVaultPropertiesPtrOutput() VaultPropertiesPtrOutput {
	return i.ToVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i VaultPropertiesArgs) ToVaultPropertiesPtrOutputWithContext(ctx context.Context) VaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesOutput).ToVaultPropertiesPtrOutputWithContext(ctx)
}









type VaultPropertiesPtrInput interface {
	pulumi.Input

	ToVaultPropertiesPtrOutput() VaultPropertiesPtrOutput
	ToVaultPropertiesPtrOutputWithContext(context.Context) VaultPropertiesPtrOutput
}

type vaultPropertiesPtrType VaultPropertiesArgs

func VaultPropertiesPtr(v *VaultPropertiesArgs) VaultPropertiesPtrInput {
	return (*vaultPropertiesPtrType)(v)
}

func (*vaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultProperties)(nil)).Elem()
}

func (i *vaultPropertiesPtrType) ToVaultPropertiesPtrOutput() VaultPropertiesPtrOutput {
	return i.ToVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *vaultPropertiesPtrType) ToVaultPropertiesPtrOutputWithContext(ctx context.Context) VaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesPtrOutput)
}

type VaultPropertiesOutput struct{ *pulumi.OutputState }

func (VaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultProperties)(nil)).Elem()
}

func (o VaultPropertiesOutput) ToVaultPropertiesOutput() VaultPropertiesOutput {
	return o
}

func (o VaultPropertiesOutput) ToVaultPropertiesOutputWithContext(ctx context.Context) VaultPropertiesOutput {
	return o
}

func (o VaultPropertiesOutput) ToVaultPropertiesPtrOutput() VaultPropertiesPtrOutput {
	return o.ToVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o VaultPropertiesOutput) ToVaultPropertiesPtrOutputWithContext(ctx context.Context) VaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VaultProperties) *VaultProperties {
		return &v
	}).(VaultPropertiesPtrOutput)
}

func (o VaultPropertiesOutput) Encryption() VaultPropertiesEncryptionPtrOutput {
	return o.ApplyT(func(v VaultProperties) *VaultPropertiesEncryption { return v.Encryption }).(VaultPropertiesEncryptionPtrOutput)
}

func (o VaultPropertiesOutput) MonitoringSettings() MonitoringSettingsPtrOutput {
	return o.ApplyT(func(v VaultProperties) *MonitoringSettings { return v.MonitoringSettings }).(MonitoringSettingsPtrOutput)
}

func (o VaultPropertiesOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultProperties) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesOutput) SecuritySettings() SecuritySettingsPtrOutput {
	return o.ApplyT(func(v VaultProperties) *SecuritySettings { return v.SecuritySettings }).(SecuritySettingsPtrOutput)
}

type VaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (VaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultProperties)(nil)).Elem()
}

func (o VaultPropertiesPtrOutput) ToVaultPropertiesPtrOutput() VaultPropertiesPtrOutput {
	return o
}

func (o VaultPropertiesPtrOutput) ToVaultPropertiesPtrOutputWithContext(ctx context.Context) VaultPropertiesPtrOutput {
	return o
}

func (o VaultPropertiesPtrOutput) Elem() VaultPropertiesOutput {
	return o.ApplyT(func(v *VaultProperties) VaultProperties {
		if v != nil {
			return *v
		}
		var ret VaultProperties
		return ret
	}).(VaultPropertiesOutput)
}

func (o VaultPropertiesPtrOutput) Encryption() VaultPropertiesEncryptionPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *VaultPropertiesEncryption {
		if v == nil {
			return nil
		}
		return v.Encryption
	}).(VaultPropertiesEncryptionPtrOutput)
}

func (o VaultPropertiesPtrOutput) MonitoringSettings() MonitoringSettingsPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *MonitoringSettings {
		if v == nil {
			return nil
		}
		return v.MonitoringSettings
	}).(MonitoringSettingsPtrOutput)
}

func (o VaultPropertiesPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesPtrOutput) SecuritySettings() SecuritySettingsPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *SecuritySettings {
		if v == nil {
			return nil
		}
		return v.SecuritySettings
	}).(SecuritySettingsPtrOutput)
}

type VaultPropertiesEncryption struct {
	InfrastructureEncryption *string                `pulumi:"infrastructureEncryption"`
	KekIdentity              *CmkKekIdentity        `pulumi:"kekIdentity"`
	KeyVaultProperties       *CmkKeyVaultProperties `pulumi:"keyVaultProperties"`
}





type VaultPropertiesEncryptionInput interface {
	pulumi.Input

	ToVaultPropertiesEncryptionOutput() VaultPropertiesEncryptionOutput
	ToVaultPropertiesEncryptionOutputWithContext(context.Context) VaultPropertiesEncryptionOutput
}

type VaultPropertiesEncryptionArgs struct {
	InfrastructureEncryption pulumi.StringPtrInput         `pulumi:"infrastructureEncryption"`
	KekIdentity              CmkKekIdentityPtrInput        `pulumi:"kekIdentity"`
	KeyVaultProperties       CmkKeyVaultPropertiesPtrInput `pulumi:"keyVaultProperties"`
}

func (VaultPropertiesEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPropertiesEncryption)(nil)).Elem()
}

func (i VaultPropertiesEncryptionArgs) ToVaultPropertiesEncryptionOutput() VaultPropertiesEncryptionOutput {
	return i.ToVaultPropertiesEncryptionOutputWithContext(context.Background())
}

func (i VaultPropertiesEncryptionArgs) ToVaultPropertiesEncryptionOutputWithContext(ctx context.Context) VaultPropertiesEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesEncryptionOutput)
}

func (i VaultPropertiesEncryptionArgs) ToVaultPropertiesEncryptionPtrOutput() VaultPropertiesEncryptionPtrOutput {
	return i.ToVaultPropertiesEncryptionPtrOutputWithContext(context.Background())
}

func (i VaultPropertiesEncryptionArgs) ToVaultPropertiesEncryptionPtrOutputWithContext(ctx context.Context) VaultPropertiesEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesEncryptionOutput).ToVaultPropertiesEncryptionPtrOutputWithContext(ctx)
}









type VaultPropertiesEncryptionPtrInput interface {
	pulumi.Input

	ToVaultPropertiesEncryptionPtrOutput() VaultPropertiesEncryptionPtrOutput
	ToVaultPropertiesEncryptionPtrOutputWithContext(context.Context) VaultPropertiesEncryptionPtrOutput
}

type vaultPropertiesEncryptionPtrType VaultPropertiesEncryptionArgs

func VaultPropertiesEncryptionPtr(v *VaultPropertiesEncryptionArgs) VaultPropertiesEncryptionPtrInput {
	return (*vaultPropertiesEncryptionPtrType)(v)
}

func (*vaultPropertiesEncryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultPropertiesEncryption)(nil)).Elem()
}

func (i *vaultPropertiesEncryptionPtrType) ToVaultPropertiesEncryptionPtrOutput() VaultPropertiesEncryptionPtrOutput {
	return i.ToVaultPropertiesEncryptionPtrOutputWithContext(context.Background())
}

func (i *vaultPropertiesEncryptionPtrType) ToVaultPropertiesEncryptionPtrOutputWithContext(ctx context.Context) VaultPropertiesEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesEncryptionPtrOutput)
}

type VaultPropertiesEncryptionOutput struct{ *pulumi.OutputState }

func (VaultPropertiesEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPropertiesEncryption)(nil)).Elem()
}

func (o VaultPropertiesEncryptionOutput) ToVaultPropertiesEncryptionOutput() VaultPropertiesEncryptionOutput {
	return o
}

func (o VaultPropertiesEncryptionOutput) ToVaultPropertiesEncryptionOutputWithContext(ctx context.Context) VaultPropertiesEncryptionOutput {
	return o
}

func (o VaultPropertiesEncryptionOutput) ToVaultPropertiesEncryptionPtrOutput() VaultPropertiesEncryptionPtrOutput {
	return o.ToVaultPropertiesEncryptionPtrOutputWithContext(context.Background())
}

func (o VaultPropertiesEncryptionOutput) ToVaultPropertiesEncryptionPtrOutputWithContext(ctx context.Context) VaultPropertiesEncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VaultPropertiesEncryption) *VaultPropertiesEncryption {
		return &v
	}).(VaultPropertiesEncryptionPtrOutput)
}

func (o VaultPropertiesEncryptionOutput) InfrastructureEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultPropertiesEncryption) *string { return v.InfrastructureEncryption }).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesEncryptionOutput) KekIdentity() CmkKekIdentityPtrOutput {
	return o.ApplyT(func(v VaultPropertiesEncryption) *CmkKekIdentity { return v.KekIdentity }).(CmkKekIdentityPtrOutput)
}

func (o VaultPropertiesEncryptionOutput) KeyVaultProperties() CmkKeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v VaultPropertiesEncryption) *CmkKeyVaultProperties { return v.KeyVaultProperties }).(CmkKeyVaultPropertiesPtrOutput)
}

type VaultPropertiesEncryptionPtrOutput struct{ *pulumi.OutputState }

func (VaultPropertiesEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultPropertiesEncryption)(nil)).Elem()
}

func (o VaultPropertiesEncryptionPtrOutput) ToVaultPropertiesEncryptionPtrOutput() VaultPropertiesEncryptionPtrOutput {
	return o
}

func (o VaultPropertiesEncryptionPtrOutput) ToVaultPropertiesEncryptionPtrOutputWithContext(ctx context.Context) VaultPropertiesEncryptionPtrOutput {
	return o
}

func (o VaultPropertiesEncryptionPtrOutput) Elem() VaultPropertiesEncryptionOutput {
	return o.ApplyT(func(v *VaultPropertiesEncryption) VaultPropertiesEncryption {
		if v != nil {
			return *v
		}
		var ret VaultPropertiesEncryption
		return ret
	}).(VaultPropertiesEncryptionOutput)
}

func (o VaultPropertiesEncryptionPtrOutput) InfrastructureEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesEncryption) *string {
		if v == nil {
			return nil
		}
		return v.InfrastructureEncryption
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesEncryptionPtrOutput) KekIdentity() CmkKekIdentityPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesEncryption) *CmkKekIdentity {
		if v == nil {
			return nil
		}
		return v.KekIdentity
	}).(CmkKekIdentityPtrOutput)
}

func (o VaultPropertiesEncryptionPtrOutput) KeyVaultProperties() CmkKeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesEncryption) *CmkKeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(CmkKeyVaultPropertiesPtrOutput)
}

type VaultPropertiesResponse struct {
	BackupStorageVersion                string                                             `pulumi:"backupStorageVersion"`
	Encryption                          *VaultPropertiesResponseEncryption                 `pulumi:"encryption"`
	MonitoringSettings                  *MonitoringSettingsResponse                        `pulumi:"monitoringSettings"`
	MoveDetails                         *VaultPropertiesResponseMoveDetails                `pulumi:"moveDetails"`
	MoveState                           string                                             `pulumi:"moveState"`
	PrivateEndpointConnections          []PrivateEndpointConnectionVaultPropertiesResponse `pulumi:"privateEndpointConnections"`
	PrivateEndpointStateForBackup       string                                             `pulumi:"privateEndpointStateForBackup"`
	PrivateEndpointStateForSiteRecovery string                                             `pulumi:"privateEndpointStateForSiteRecovery"`
	ProvisioningState                   string                                             `pulumi:"provisioningState"`
	PublicNetworkAccess                 *string                                            `pulumi:"publicNetworkAccess"`
	RedundancySettings                  *VaultPropertiesResponseRedundancySettings         `pulumi:"redundancySettings"`
	SecuritySettings                    *SecuritySettingsResponse                          `pulumi:"securitySettings"`
	UpgradeDetails                      *UpgradeDetailsResponse                            `pulumi:"upgradeDetails"`
}

type VaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPropertiesResponse)(nil)).Elem()
}

func (o VaultPropertiesResponseOutput) ToVaultPropertiesResponseOutput() VaultPropertiesResponseOutput {
	return o
}

func (o VaultPropertiesResponseOutput) ToVaultPropertiesResponseOutputWithContext(ctx context.Context) VaultPropertiesResponseOutput {
	return o
}

func (o VaultPropertiesResponseOutput) BackupStorageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) string { return v.BackupStorageVersion }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseOutput) Encryption() VaultPropertiesResponseEncryptionPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *VaultPropertiesResponseEncryption { return v.Encryption }).(VaultPropertiesResponseEncryptionPtrOutput)
}

func (o VaultPropertiesResponseOutput) MonitoringSettings() MonitoringSettingsResponsePtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *MonitoringSettingsResponse { return v.MonitoringSettings }).(MonitoringSettingsResponsePtrOutput)
}

func (o VaultPropertiesResponseOutput) MoveDetails() VaultPropertiesResponseMoveDetailsPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *VaultPropertiesResponseMoveDetails { return v.MoveDetails }).(VaultPropertiesResponseMoveDetailsPtrOutput)
}

func (o VaultPropertiesResponseOutput) MoveState() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) string { return v.MoveState }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseOutput) PrivateEndpointConnections() PrivateEndpointConnectionVaultPropertiesResponseArrayOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) []PrivateEndpointConnectionVaultPropertiesResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionVaultPropertiesResponseArrayOutput)
}

func (o VaultPropertiesResponseOutput) PrivateEndpointStateForBackup() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) string { return v.PrivateEndpointStateForBackup }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseOutput) PrivateEndpointStateForSiteRecovery() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) string { return v.PrivateEndpointStateForSiteRecovery }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponseOutput) RedundancySettings() VaultPropertiesResponseRedundancySettingsPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *VaultPropertiesResponseRedundancySettings {
		return v.RedundancySettings
	}).(VaultPropertiesResponseRedundancySettingsPtrOutput)
}

func (o VaultPropertiesResponseOutput) SecuritySettings() SecuritySettingsResponsePtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *SecuritySettingsResponse { return v.SecuritySettings }).(SecuritySettingsResponsePtrOutput)
}

func (o VaultPropertiesResponseOutput) UpgradeDetails() UpgradeDetailsResponsePtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *UpgradeDetailsResponse { return v.UpgradeDetails }).(UpgradeDetailsResponsePtrOutput)
}

type VaultPropertiesResponseEncryption struct {
	InfrastructureEncryption *string                        `pulumi:"infrastructureEncryption"`
	KekIdentity              *CmkKekIdentityResponse        `pulumi:"kekIdentity"`
	KeyVaultProperties       *CmkKeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
}

type VaultPropertiesResponseEncryptionOutput struct{ *pulumi.OutputState }

func (VaultPropertiesResponseEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPropertiesResponseEncryption)(nil)).Elem()
}

func (o VaultPropertiesResponseEncryptionOutput) ToVaultPropertiesResponseEncryptionOutput() VaultPropertiesResponseEncryptionOutput {
	return o
}

func (o VaultPropertiesResponseEncryptionOutput) ToVaultPropertiesResponseEncryptionOutputWithContext(ctx context.Context) VaultPropertiesResponseEncryptionOutput {
	return o
}

func (o VaultPropertiesResponseEncryptionOutput) InfrastructureEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponseEncryption) *string { return v.InfrastructureEncryption }).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponseEncryptionOutput) KekIdentity() CmkKekIdentityResponsePtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponseEncryption) *CmkKekIdentityResponse { return v.KekIdentity }).(CmkKekIdentityResponsePtrOutput)
}

func (o VaultPropertiesResponseEncryptionOutput) KeyVaultProperties() CmkKeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponseEncryption) *CmkKeyVaultPropertiesResponse { return v.KeyVaultProperties }).(CmkKeyVaultPropertiesResponsePtrOutput)
}

type VaultPropertiesResponseEncryptionPtrOutput struct{ *pulumi.OutputState }

func (VaultPropertiesResponseEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultPropertiesResponseEncryption)(nil)).Elem()
}

func (o VaultPropertiesResponseEncryptionPtrOutput) ToVaultPropertiesResponseEncryptionPtrOutput() VaultPropertiesResponseEncryptionPtrOutput {
	return o
}

func (o VaultPropertiesResponseEncryptionPtrOutput) ToVaultPropertiesResponseEncryptionPtrOutputWithContext(ctx context.Context) VaultPropertiesResponseEncryptionPtrOutput {
	return o
}

func (o VaultPropertiesResponseEncryptionPtrOutput) Elem() VaultPropertiesResponseEncryptionOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseEncryption) VaultPropertiesResponseEncryption {
		if v != nil {
			return *v
		}
		var ret VaultPropertiesResponseEncryption
		return ret
	}).(VaultPropertiesResponseEncryptionOutput)
}

func (o VaultPropertiesResponseEncryptionPtrOutput) InfrastructureEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseEncryption) *string {
		if v == nil {
			return nil
		}
		return v.InfrastructureEncryption
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponseEncryptionPtrOutput) KekIdentity() CmkKekIdentityResponsePtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseEncryption) *CmkKekIdentityResponse {
		if v == nil {
			return nil
		}
		return v.KekIdentity
	}).(CmkKekIdentityResponsePtrOutput)
}

func (o VaultPropertiesResponseEncryptionPtrOutput) KeyVaultProperties() CmkKeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseEncryption) *CmkKeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(CmkKeyVaultPropertiesResponsePtrOutput)
}

type VaultPropertiesResponseMoveDetails struct {
	CompletionTimeUtc string `pulumi:"completionTimeUtc"`
	OperationId       string `pulumi:"operationId"`
	SourceResourceId  string `pulumi:"sourceResourceId"`
	StartTimeUtc      string `pulumi:"startTimeUtc"`
	TargetResourceId  string `pulumi:"targetResourceId"`
}

type VaultPropertiesResponseMoveDetailsOutput struct{ *pulumi.OutputState }

func (VaultPropertiesResponseMoveDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPropertiesResponseMoveDetails)(nil)).Elem()
}

func (o VaultPropertiesResponseMoveDetailsOutput) ToVaultPropertiesResponseMoveDetailsOutput() VaultPropertiesResponseMoveDetailsOutput {
	return o
}

func (o VaultPropertiesResponseMoveDetailsOutput) ToVaultPropertiesResponseMoveDetailsOutputWithContext(ctx context.Context) VaultPropertiesResponseMoveDetailsOutput {
	return o
}

func (o VaultPropertiesResponseMoveDetailsOutput) CompletionTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponseMoveDetails) string { return v.CompletionTimeUtc }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseMoveDetailsOutput) OperationId() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponseMoveDetails) string { return v.OperationId }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseMoveDetailsOutput) SourceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponseMoveDetails) string { return v.SourceResourceId }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseMoveDetailsOutput) StartTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponseMoveDetails) string { return v.StartTimeUtc }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseMoveDetailsOutput) TargetResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponseMoveDetails) string { return v.TargetResourceId }).(pulumi.StringOutput)
}

type VaultPropertiesResponseMoveDetailsPtrOutput struct{ *pulumi.OutputState }

func (VaultPropertiesResponseMoveDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultPropertiesResponseMoveDetails)(nil)).Elem()
}

func (o VaultPropertiesResponseMoveDetailsPtrOutput) ToVaultPropertiesResponseMoveDetailsPtrOutput() VaultPropertiesResponseMoveDetailsPtrOutput {
	return o
}

func (o VaultPropertiesResponseMoveDetailsPtrOutput) ToVaultPropertiesResponseMoveDetailsPtrOutputWithContext(ctx context.Context) VaultPropertiesResponseMoveDetailsPtrOutput {
	return o
}

func (o VaultPropertiesResponseMoveDetailsPtrOutput) Elem() VaultPropertiesResponseMoveDetailsOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseMoveDetails) VaultPropertiesResponseMoveDetails {
		if v != nil {
			return *v
		}
		var ret VaultPropertiesResponseMoveDetails
		return ret
	}).(VaultPropertiesResponseMoveDetailsOutput)
}

func (o VaultPropertiesResponseMoveDetailsPtrOutput) CompletionTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseMoveDetails) *string {
		if v == nil {
			return nil
		}
		return &v.CompletionTimeUtc
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponseMoveDetailsPtrOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseMoveDetails) *string {
		if v == nil {
			return nil
		}
		return &v.OperationId
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponseMoveDetailsPtrOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseMoveDetails) *string {
		if v == nil {
			return nil
		}
		return &v.SourceResourceId
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponseMoveDetailsPtrOutput) StartTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseMoveDetails) *string {
		if v == nil {
			return nil
		}
		return &v.StartTimeUtc
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponseMoveDetailsPtrOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseMoveDetails) *string {
		if v == nil {
			return nil
		}
		return &v.TargetResourceId
	}).(pulumi.StringPtrOutput)
}

type VaultPropertiesResponseRedundancySettings struct {
	CrossRegionRestore            string `pulumi:"crossRegionRestore"`
	StandardTierStorageRedundancy string `pulumi:"standardTierStorageRedundancy"`
}

type VaultPropertiesResponseRedundancySettingsOutput struct{ *pulumi.OutputState }

func (VaultPropertiesResponseRedundancySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPropertiesResponseRedundancySettings)(nil)).Elem()
}

func (o VaultPropertiesResponseRedundancySettingsOutput) ToVaultPropertiesResponseRedundancySettingsOutput() VaultPropertiesResponseRedundancySettingsOutput {
	return o
}

func (o VaultPropertiesResponseRedundancySettingsOutput) ToVaultPropertiesResponseRedundancySettingsOutputWithContext(ctx context.Context) VaultPropertiesResponseRedundancySettingsOutput {
	return o
}

func (o VaultPropertiesResponseRedundancySettingsOutput) CrossRegionRestore() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponseRedundancySettings) string { return v.CrossRegionRestore }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseRedundancySettingsOutput) StandardTierStorageRedundancy() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponseRedundancySettings) string { return v.StandardTierStorageRedundancy }).(pulumi.StringOutput)
}

type VaultPropertiesResponseRedundancySettingsPtrOutput struct{ *pulumi.OutputState }

func (VaultPropertiesResponseRedundancySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultPropertiesResponseRedundancySettings)(nil)).Elem()
}

func (o VaultPropertiesResponseRedundancySettingsPtrOutput) ToVaultPropertiesResponseRedundancySettingsPtrOutput() VaultPropertiesResponseRedundancySettingsPtrOutput {
	return o
}

func (o VaultPropertiesResponseRedundancySettingsPtrOutput) ToVaultPropertiesResponseRedundancySettingsPtrOutputWithContext(ctx context.Context) VaultPropertiesResponseRedundancySettingsPtrOutput {
	return o
}

func (o VaultPropertiesResponseRedundancySettingsPtrOutput) Elem() VaultPropertiesResponseRedundancySettingsOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseRedundancySettings) VaultPropertiesResponseRedundancySettings {
		if v != nil {
			return *v
		}
		var ret VaultPropertiesResponseRedundancySettings
		return ret
	}).(VaultPropertiesResponseRedundancySettingsOutput)
}

func (o VaultPropertiesResponseRedundancySettingsPtrOutput) CrossRegionRestore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseRedundancySettings) *string {
		if v == nil {
			return nil
		}
		return &v.CrossRegionRestore
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponseRedundancySettingsPtrOutput) StandardTierStorageRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseRedundancySettings) *string {
		if v == nil {
			return nil
		}
		return &v.StandardTierStorageRedundancy
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureMonitorAlertSettingsOutput{})
	pulumi.RegisterOutputType(AzureMonitorAlertSettingsPtrOutput{})
	pulumi.RegisterOutputType(AzureMonitorAlertSettingsResponseOutput{})
	pulumi.RegisterOutputType(AzureMonitorAlertSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(ClassicAlertSettingsOutput{})
	pulumi.RegisterOutputType(ClassicAlertSettingsPtrOutput{})
	pulumi.RegisterOutputType(ClassicAlertSettingsResponseOutput{})
	pulumi.RegisterOutputType(ClassicAlertSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(CmkKekIdentityOutput{})
	pulumi.RegisterOutputType(CmkKekIdentityPtrOutput{})
	pulumi.RegisterOutputType(CmkKekIdentityResponseOutput{})
	pulumi.RegisterOutputType(CmkKekIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(CmkKeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(CmkKeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CmkKeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CmkKeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityDataOutput{})
	pulumi.RegisterOutputType(IdentityDataPtrOutput{})
	pulumi.RegisterOutputType(IdentityDataResponseOutput{})
	pulumi.RegisterOutputType(IdentityDataResponsePtrOutput{})
	pulumi.RegisterOutputType(ImmutabilitySettingsOutput{})
	pulumi.RegisterOutputType(ImmutabilitySettingsPtrOutput{})
	pulumi.RegisterOutputType(ImmutabilitySettingsResponseOutput{})
	pulumi.RegisterOutputType(ImmutabilitySettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(MonitoringSettingsOutput{})
	pulumi.RegisterOutputType(MonitoringSettingsPtrOutput{})
	pulumi.RegisterOutputType(MonitoringSettingsResponseOutput{})
	pulumi.RegisterOutputType(MonitoringSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionVaultPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(SecuritySettingsOutput{})
	pulumi.RegisterOutputType(SecuritySettingsPtrOutput{})
	pulumi.RegisterOutputType(SecuritySettingsResponseOutput{})
	pulumi.RegisterOutputType(SecuritySettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UpgradeDetailsResponseOutput{})
	pulumi.RegisterOutputType(UpgradeDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(UserIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(VaultPrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(VaultPrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(VaultPropertiesOutput{})
	pulumi.RegisterOutputType(VaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VaultPropertiesEncryptionOutput{})
	pulumi.RegisterOutputType(VaultPropertiesEncryptionPtrOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseEncryptionOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseEncryptionPtrOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseMoveDetailsOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseMoveDetailsPtrOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseRedundancySettingsOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseRedundancySettingsPtrOutput{})
}

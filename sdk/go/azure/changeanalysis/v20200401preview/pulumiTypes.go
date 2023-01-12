


package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureMonitorWorkspaceProperties struct {
	IncludeChangeDetails *string `pulumi:"includeChangeDetails"`
	WorkspaceId          *string `pulumi:"workspaceId"`
	WorkspaceResourceId  *string `pulumi:"workspaceResourceId"`
}





type AzureMonitorWorkspacePropertiesInput interface {
	pulumi.Input

	ToAzureMonitorWorkspacePropertiesOutput() AzureMonitorWorkspacePropertiesOutput
	ToAzureMonitorWorkspacePropertiesOutputWithContext(context.Context) AzureMonitorWorkspacePropertiesOutput
}

type AzureMonitorWorkspacePropertiesArgs struct {
	IncludeChangeDetails pulumi.StringPtrInput `pulumi:"includeChangeDetails"`
	WorkspaceId          pulumi.StringPtrInput `pulumi:"workspaceId"`
	WorkspaceResourceId  pulumi.StringPtrInput `pulumi:"workspaceResourceId"`
}

func (AzureMonitorWorkspacePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMonitorWorkspaceProperties)(nil)).Elem()
}

func (i AzureMonitorWorkspacePropertiesArgs) ToAzureMonitorWorkspacePropertiesOutput() AzureMonitorWorkspacePropertiesOutput {
	return i.ToAzureMonitorWorkspacePropertiesOutputWithContext(context.Background())
}

func (i AzureMonitorWorkspacePropertiesArgs) ToAzureMonitorWorkspacePropertiesOutputWithContext(ctx context.Context) AzureMonitorWorkspacePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMonitorWorkspacePropertiesOutput)
}

func (i AzureMonitorWorkspacePropertiesArgs) ToAzureMonitorWorkspacePropertiesPtrOutput() AzureMonitorWorkspacePropertiesPtrOutput {
	return i.ToAzureMonitorWorkspacePropertiesPtrOutputWithContext(context.Background())
}

func (i AzureMonitorWorkspacePropertiesArgs) ToAzureMonitorWorkspacePropertiesPtrOutputWithContext(ctx context.Context) AzureMonitorWorkspacePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMonitorWorkspacePropertiesOutput).ToAzureMonitorWorkspacePropertiesPtrOutputWithContext(ctx)
}









type AzureMonitorWorkspacePropertiesPtrInput interface {
	pulumi.Input

	ToAzureMonitorWorkspacePropertiesPtrOutput() AzureMonitorWorkspacePropertiesPtrOutput
	ToAzureMonitorWorkspacePropertiesPtrOutputWithContext(context.Context) AzureMonitorWorkspacePropertiesPtrOutput
}

type azureMonitorWorkspacePropertiesPtrType AzureMonitorWorkspacePropertiesArgs

func AzureMonitorWorkspacePropertiesPtr(v *AzureMonitorWorkspacePropertiesArgs) AzureMonitorWorkspacePropertiesPtrInput {
	return (*azureMonitorWorkspacePropertiesPtrType)(v)
}

func (*azureMonitorWorkspacePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureMonitorWorkspaceProperties)(nil)).Elem()
}

func (i *azureMonitorWorkspacePropertiesPtrType) ToAzureMonitorWorkspacePropertiesPtrOutput() AzureMonitorWorkspacePropertiesPtrOutput {
	return i.ToAzureMonitorWorkspacePropertiesPtrOutputWithContext(context.Background())
}

func (i *azureMonitorWorkspacePropertiesPtrType) ToAzureMonitorWorkspacePropertiesPtrOutputWithContext(ctx context.Context) AzureMonitorWorkspacePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMonitorWorkspacePropertiesPtrOutput)
}

type AzureMonitorWorkspacePropertiesOutput struct{ *pulumi.OutputState }

func (AzureMonitorWorkspacePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMonitorWorkspaceProperties)(nil)).Elem()
}

func (o AzureMonitorWorkspacePropertiesOutput) ToAzureMonitorWorkspacePropertiesOutput() AzureMonitorWorkspacePropertiesOutput {
	return o
}

func (o AzureMonitorWorkspacePropertiesOutput) ToAzureMonitorWorkspacePropertiesOutputWithContext(ctx context.Context) AzureMonitorWorkspacePropertiesOutput {
	return o
}

func (o AzureMonitorWorkspacePropertiesOutput) ToAzureMonitorWorkspacePropertiesPtrOutput() AzureMonitorWorkspacePropertiesPtrOutput {
	return o.ToAzureMonitorWorkspacePropertiesPtrOutputWithContext(context.Background())
}

func (o AzureMonitorWorkspacePropertiesOutput) ToAzureMonitorWorkspacePropertiesPtrOutputWithContext(ctx context.Context) AzureMonitorWorkspacePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureMonitorWorkspaceProperties) *AzureMonitorWorkspaceProperties {
		return &v
	}).(AzureMonitorWorkspacePropertiesPtrOutput)
}

func (o AzureMonitorWorkspacePropertiesOutput) IncludeChangeDetails() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMonitorWorkspaceProperties) *string { return v.IncludeChangeDetails }).(pulumi.StringPtrOutput)
}

func (o AzureMonitorWorkspacePropertiesOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMonitorWorkspaceProperties) *string { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

func (o AzureMonitorWorkspacePropertiesOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMonitorWorkspaceProperties) *string { return v.WorkspaceResourceId }).(pulumi.StringPtrOutput)
}

type AzureMonitorWorkspacePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AzureMonitorWorkspacePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureMonitorWorkspaceProperties)(nil)).Elem()
}

func (o AzureMonitorWorkspacePropertiesPtrOutput) ToAzureMonitorWorkspacePropertiesPtrOutput() AzureMonitorWorkspacePropertiesPtrOutput {
	return o
}

func (o AzureMonitorWorkspacePropertiesPtrOutput) ToAzureMonitorWorkspacePropertiesPtrOutputWithContext(ctx context.Context) AzureMonitorWorkspacePropertiesPtrOutput {
	return o
}

func (o AzureMonitorWorkspacePropertiesPtrOutput) Elem() AzureMonitorWorkspacePropertiesOutput {
	return o.ApplyT(func(v *AzureMonitorWorkspaceProperties) AzureMonitorWorkspaceProperties {
		if v != nil {
			return *v
		}
		var ret AzureMonitorWorkspaceProperties
		return ret
	}).(AzureMonitorWorkspacePropertiesOutput)
}

func (o AzureMonitorWorkspacePropertiesPtrOutput) IncludeChangeDetails() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMonitorWorkspaceProperties) *string {
		if v == nil {
			return nil
		}
		return v.IncludeChangeDetails
	}).(pulumi.StringPtrOutput)
}

func (o AzureMonitorWorkspacePropertiesPtrOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMonitorWorkspaceProperties) *string {
		if v == nil {
			return nil
		}
		return v.WorkspaceId
	}).(pulumi.StringPtrOutput)
}

func (o AzureMonitorWorkspacePropertiesPtrOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMonitorWorkspaceProperties) *string {
		if v == nil {
			return nil
		}
		return v.WorkspaceResourceId
	}).(pulumi.StringPtrOutput)
}

type AzureMonitorWorkspacePropertiesResponse struct {
	IncludeChangeDetails *string `pulumi:"includeChangeDetails"`
	WorkspaceId          *string `pulumi:"workspaceId"`
	WorkspaceResourceId  *string `pulumi:"workspaceResourceId"`
}

type AzureMonitorWorkspacePropertiesResponseOutput struct{ *pulumi.OutputState }

func (AzureMonitorWorkspacePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMonitorWorkspacePropertiesResponse)(nil)).Elem()
}

func (o AzureMonitorWorkspacePropertiesResponseOutput) ToAzureMonitorWorkspacePropertiesResponseOutput() AzureMonitorWorkspacePropertiesResponseOutput {
	return o
}

func (o AzureMonitorWorkspacePropertiesResponseOutput) ToAzureMonitorWorkspacePropertiesResponseOutputWithContext(ctx context.Context) AzureMonitorWorkspacePropertiesResponseOutput {
	return o
}

func (o AzureMonitorWorkspacePropertiesResponseOutput) IncludeChangeDetails() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMonitorWorkspacePropertiesResponse) *string { return v.IncludeChangeDetails }).(pulumi.StringPtrOutput)
}

func (o AzureMonitorWorkspacePropertiesResponseOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMonitorWorkspacePropertiesResponse) *string { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

func (o AzureMonitorWorkspacePropertiesResponseOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMonitorWorkspacePropertiesResponse) *string { return v.WorkspaceResourceId }).(pulumi.StringPtrOutput)
}

type AzureMonitorWorkspacePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureMonitorWorkspacePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureMonitorWorkspacePropertiesResponse)(nil)).Elem()
}

func (o AzureMonitorWorkspacePropertiesResponsePtrOutput) ToAzureMonitorWorkspacePropertiesResponsePtrOutput() AzureMonitorWorkspacePropertiesResponsePtrOutput {
	return o
}

func (o AzureMonitorWorkspacePropertiesResponsePtrOutput) ToAzureMonitorWorkspacePropertiesResponsePtrOutputWithContext(ctx context.Context) AzureMonitorWorkspacePropertiesResponsePtrOutput {
	return o
}

func (o AzureMonitorWorkspacePropertiesResponsePtrOutput) Elem() AzureMonitorWorkspacePropertiesResponseOutput {
	return o.ApplyT(func(v *AzureMonitorWorkspacePropertiesResponse) AzureMonitorWorkspacePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AzureMonitorWorkspacePropertiesResponse
		return ret
	}).(AzureMonitorWorkspacePropertiesResponseOutput)
}

func (o AzureMonitorWorkspacePropertiesResponsePtrOutput) IncludeChangeDetails() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMonitorWorkspacePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.IncludeChangeDetails
	}).(pulumi.StringPtrOutput)
}

func (o AzureMonitorWorkspacePropertiesResponsePtrOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMonitorWorkspacePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.WorkspaceId
	}).(pulumi.StringPtrOutput)
}

func (o AzureMonitorWorkspacePropertiesResponsePtrOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMonitorWorkspacePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.WorkspaceResourceId
	}).(pulumi.StringPtrOutput)
}

type ConfigurationProfileResourceProperties struct {
	Notifications *NotificationSettings `pulumi:"notifications"`
}





type ConfigurationProfileResourcePropertiesInput interface {
	pulumi.Input

	ToConfigurationProfileResourcePropertiesOutput() ConfigurationProfileResourcePropertiesOutput
	ToConfigurationProfileResourcePropertiesOutputWithContext(context.Context) ConfigurationProfileResourcePropertiesOutput
}

type ConfigurationProfileResourcePropertiesArgs struct {
	Notifications NotificationSettingsPtrInput `pulumi:"notifications"`
}

func (ConfigurationProfileResourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileResourceProperties)(nil)).Elem()
}

func (i ConfigurationProfileResourcePropertiesArgs) ToConfigurationProfileResourcePropertiesOutput() ConfigurationProfileResourcePropertiesOutput {
	return i.ToConfigurationProfileResourcePropertiesOutputWithContext(context.Background())
}

func (i ConfigurationProfileResourcePropertiesArgs) ToConfigurationProfileResourcePropertiesOutputWithContext(ctx context.Context) ConfigurationProfileResourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileResourcePropertiesOutput)
}

func (i ConfigurationProfileResourcePropertiesArgs) ToConfigurationProfileResourcePropertiesPtrOutput() ConfigurationProfileResourcePropertiesPtrOutput {
	return i.ToConfigurationProfileResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i ConfigurationProfileResourcePropertiesArgs) ToConfigurationProfileResourcePropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfileResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileResourcePropertiesOutput).ToConfigurationProfileResourcePropertiesPtrOutputWithContext(ctx)
}









type ConfigurationProfileResourcePropertiesPtrInput interface {
	pulumi.Input

	ToConfigurationProfileResourcePropertiesPtrOutput() ConfigurationProfileResourcePropertiesPtrOutput
	ToConfigurationProfileResourcePropertiesPtrOutputWithContext(context.Context) ConfigurationProfileResourcePropertiesPtrOutput
}

type configurationProfileResourcePropertiesPtrType ConfigurationProfileResourcePropertiesArgs

func ConfigurationProfileResourcePropertiesPtr(v *ConfigurationProfileResourcePropertiesArgs) ConfigurationProfileResourcePropertiesPtrInput {
	return (*configurationProfileResourcePropertiesPtrType)(v)
}

func (*configurationProfileResourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileResourceProperties)(nil)).Elem()
}

func (i *configurationProfileResourcePropertiesPtrType) ToConfigurationProfileResourcePropertiesPtrOutput() ConfigurationProfileResourcePropertiesPtrOutput {
	return i.ToConfigurationProfileResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *configurationProfileResourcePropertiesPtrType) ToConfigurationProfileResourcePropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfileResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileResourcePropertiesPtrOutput)
}

type ConfigurationProfileResourcePropertiesOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileResourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileResourceProperties)(nil)).Elem()
}

func (o ConfigurationProfileResourcePropertiesOutput) ToConfigurationProfileResourcePropertiesOutput() ConfigurationProfileResourcePropertiesOutput {
	return o
}

func (o ConfigurationProfileResourcePropertiesOutput) ToConfigurationProfileResourcePropertiesOutputWithContext(ctx context.Context) ConfigurationProfileResourcePropertiesOutput {
	return o
}

func (o ConfigurationProfileResourcePropertiesOutput) ToConfigurationProfileResourcePropertiesPtrOutput() ConfigurationProfileResourcePropertiesPtrOutput {
	return o.ToConfigurationProfileResourcePropertiesPtrOutputWithContext(context.Background())
}

func (o ConfigurationProfileResourcePropertiesOutput) ToConfigurationProfileResourcePropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfileResourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationProfileResourceProperties) *ConfigurationProfileResourceProperties {
		return &v
	}).(ConfigurationProfileResourcePropertiesPtrOutput)
}

func (o ConfigurationProfileResourcePropertiesOutput) Notifications() NotificationSettingsPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileResourceProperties) *NotificationSettings { return v.Notifications }).(NotificationSettingsPtrOutput)
}

type ConfigurationProfileResourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileResourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileResourceProperties)(nil)).Elem()
}

func (o ConfigurationProfileResourcePropertiesPtrOutput) ToConfigurationProfileResourcePropertiesPtrOutput() ConfigurationProfileResourcePropertiesPtrOutput {
	return o
}

func (o ConfigurationProfileResourcePropertiesPtrOutput) ToConfigurationProfileResourcePropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfileResourcePropertiesPtrOutput {
	return o
}

func (o ConfigurationProfileResourcePropertiesPtrOutput) Elem() ConfigurationProfileResourcePropertiesOutput {
	return o.ApplyT(func(v *ConfigurationProfileResourceProperties) ConfigurationProfileResourceProperties {
		if v != nil {
			return *v
		}
		var ret ConfigurationProfileResourceProperties
		return ret
	}).(ConfigurationProfileResourcePropertiesOutput)
}

func (o ConfigurationProfileResourcePropertiesPtrOutput) Notifications() NotificationSettingsPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileResourceProperties) *NotificationSettings {
		if v == nil {
			return nil
		}
		return v.Notifications
	}).(NotificationSettingsPtrOutput)
}

type ConfigurationProfileResourcePropertiesResponse struct {
	Notifications *NotificationSettingsResponse `pulumi:"notifications"`
}

type ConfigurationProfileResourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileResourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileResourcePropertiesResponse)(nil)).Elem()
}

func (o ConfigurationProfileResourcePropertiesResponseOutput) ToConfigurationProfileResourcePropertiesResponseOutput() ConfigurationProfileResourcePropertiesResponseOutput {
	return o
}

func (o ConfigurationProfileResourcePropertiesResponseOutput) ToConfigurationProfileResourcePropertiesResponseOutputWithContext(ctx context.Context) ConfigurationProfileResourcePropertiesResponseOutput {
	return o
}

func (o ConfigurationProfileResourcePropertiesResponseOutput) Notifications() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v ConfigurationProfileResourcePropertiesResponse) *NotificationSettingsResponse {
		return v.Notifications
	}).(NotificationSettingsResponsePtrOutput)
}

type NotificationSettings struct {
	ActivationState                 *string                          `pulumi:"activationState"`
	AzureMonitorWorkspaceProperties *AzureMonitorWorkspaceProperties `pulumi:"azureMonitorWorkspaceProperties"`
}





type NotificationSettingsInput interface {
	pulumi.Input

	ToNotificationSettingsOutput() NotificationSettingsOutput
	ToNotificationSettingsOutputWithContext(context.Context) NotificationSettingsOutput
}

type NotificationSettingsArgs struct {
	ActivationState                 pulumi.StringPtrInput                   `pulumi:"activationState"`
	AzureMonitorWorkspaceProperties AzureMonitorWorkspacePropertiesPtrInput `pulumi:"azureMonitorWorkspaceProperties"`
}

func (NotificationSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationSettings)(nil)).Elem()
}

func (i NotificationSettingsArgs) ToNotificationSettingsOutput() NotificationSettingsOutput {
	return i.ToNotificationSettingsOutputWithContext(context.Background())
}

func (i NotificationSettingsArgs) ToNotificationSettingsOutputWithContext(ctx context.Context) NotificationSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationSettingsOutput)
}

func (i NotificationSettingsArgs) ToNotificationSettingsPtrOutput() NotificationSettingsPtrOutput {
	return i.ToNotificationSettingsPtrOutputWithContext(context.Background())
}

func (i NotificationSettingsArgs) ToNotificationSettingsPtrOutputWithContext(ctx context.Context) NotificationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationSettingsOutput).ToNotificationSettingsPtrOutputWithContext(ctx)
}









type NotificationSettingsPtrInput interface {
	pulumi.Input

	ToNotificationSettingsPtrOutput() NotificationSettingsPtrOutput
	ToNotificationSettingsPtrOutputWithContext(context.Context) NotificationSettingsPtrOutput
}

type notificationSettingsPtrType NotificationSettingsArgs

func NotificationSettingsPtr(v *NotificationSettingsArgs) NotificationSettingsPtrInput {
	return (*notificationSettingsPtrType)(v)
}

func (*notificationSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationSettings)(nil)).Elem()
}

func (i *notificationSettingsPtrType) ToNotificationSettingsPtrOutput() NotificationSettingsPtrOutput {
	return i.ToNotificationSettingsPtrOutputWithContext(context.Background())
}

func (i *notificationSettingsPtrType) ToNotificationSettingsPtrOutputWithContext(ctx context.Context) NotificationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationSettingsPtrOutput)
}

type NotificationSettingsOutput struct{ *pulumi.OutputState }

func (NotificationSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationSettings)(nil)).Elem()
}

func (o NotificationSettingsOutput) ToNotificationSettingsOutput() NotificationSettingsOutput {
	return o
}

func (o NotificationSettingsOutput) ToNotificationSettingsOutputWithContext(ctx context.Context) NotificationSettingsOutput {
	return o
}

func (o NotificationSettingsOutput) ToNotificationSettingsPtrOutput() NotificationSettingsPtrOutput {
	return o.ToNotificationSettingsPtrOutputWithContext(context.Background())
}

func (o NotificationSettingsOutput) ToNotificationSettingsPtrOutputWithContext(ctx context.Context) NotificationSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotificationSettings) *NotificationSettings {
		return &v
	}).(NotificationSettingsPtrOutput)
}

func (o NotificationSettingsOutput) ActivationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationSettings) *string { return v.ActivationState }).(pulumi.StringPtrOutput)
}

func (o NotificationSettingsOutput) AzureMonitorWorkspaceProperties() AzureMonitorWorkspacePropertiesPtrOutput {
	return o.ApplyT(func(v NotificationSettings) *AzureMonitorWorkspaceProperties {
		return v.AzureMonitorWorkspaceProperties
	}).(AzureMonitorWorkspacePropertiesPtrOutput)
}

type NotificationSettingsPtrOutput struct{ *pulumi.OutputState }

func (NotificationSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationSettings)(nil)).Elem()
}

func (o NotificationSettingsPtrOutput) ToNotificationSettingsPtrOutput() NotificationSettingsPtrOutput {
	return o
}

func (o NotificationSettingsPtrOutput) ToNotificationSettingsPtrOutputWithContext(ctx context.Context) NotificationSettingsPtrOutput {
	return o
}

func (o NotificationSettingsPtrOutput) Elem() NotificationSettingsOutput {
	return o.ApplyT(func(v *NotificationSettings) NotificationSettings {
		if v != nil {
			return *v
		}
		var ret NotificationSettings
		return ret
	}).(NotificationSettingsOutput)
}

func (o NotificationSettingsPtrOutput) ActivationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationSettings) *string {
		if v == nil {
			return nil
		}
		return v.ActivationState
	}).(pulumi.StringPtrOutput)
}

func (o NotificationSettingsPtrOutput) AzureMonitorWorkspaceProperties() AzureMonitorWorkspacePropertiesPtrOutput {
	return o.ApplyT(func(v *NotificationSettings) *AzureMonitorWorkspaceProperties {
		if v == nil {
			return nil
		}
		return v.AzureMonitorWorkspaceProperties
	}).(AzureMonitorWorkspacePropertiesPtrOutput)
}

type NotificationSettingsResponse struct {
	ActivationState                 *string                                  `pulumi:"activationState"`
	AzureMonitorWorkspaceProperties *AzureMonitorWorkspacePropertiesResponse `pulumi:"azureMonitorWorkspaceProperties"`
}

type NotificationSettingsResponseOutput struct{ *pulumi.OutputState }

func (NotificationSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationSettingsResponse)(nil)).Elem()
}

func (o NotificationSettingsResponseOutput) ToNotificationSettingsResponseOutput() NotificationSettingsResponseOutput {
	return o
}

func (o NotificationSettingsResponseOutput) ToNotificationSettingsResponseOutputWithContext(ctx context.Context) NotificationSettingsResponseOutput {
	return o
}

func (o NotificationSettingsResponseOutput) ActivationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationSettingsResponse) *string { return v.ActivationState }).(pulumi.StringPtrOutput)
}

func (o NotificationSettingsResponseOutput) AzureMonitorWorkspaceProperties() AzureMonitorWorkspacePropertiesResponsePtrOutput {
	return o.ApplyT(func(v NotificationSettingsResponse) *AzureMonitorWorkspacePropertiesResponse {
		return v.AzureMonitorWorkspaceProperties
	}).(AzureMonitorWorkspacePropertiesResponsePtrOutput)
}

type NotificationSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (NotificationSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationSettingsResponse)(nil)).Elem()
}

func (o NotificationSettingsResponsePtrOutput) ToNotificationSettingsResponsePtrOutput() NotificationSettingsResponsePtrOutput {
	return o
}

func (o NotificationSettingsResponsePtrOutput) ToNotificationSettingsResponsePtrOutputWithContext(ctx context.Context) NotificationSettingsResponsePtrOutput {
	return o
}

func (o NotificationSettingsResponsePtrOutput) Elem() NotificationSettingsResponseOutput {
	return o.ApplyT(func(v *NotificationSettingsResponse) NotificationSettingsResponse {
		if v != nil {
			return *v
		}
		var ret NotificationSettingsResponse
		return ret
	}).(NotificationSettingsResponseOutput)
}

func (o NotificationSettingsResponsePtrOutput) ActivationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActivationState
	}).(pulumi.StringPtrOutput)
}

func (o NotificationSettingsResponsePtrOutput) AzureMonitorWorkspaceProperties() AzureMonitorWorkspacePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *NotificationSettingsResponse) *AzureMonitorWorkspacePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.AzureMonitorWorkspaceProperties
	}).(AzureMonitorWorkspacePropertiesResponsePtrOutput)
}

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

type ResourceIdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type ResourceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutputWithContext(ctx context.Context) ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ResourceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ResourceIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ResourceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) Elem() ResourceIdentityResponseOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) ResourceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityResponse
		return ret
	}).(ResourceIdentityResponseOutput)
}

func (o ResourceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          string `pulumi:"createdAt"`
	CreatedBy          string `pulumi:"createdBy"`
	CreatedByType      string `pulumi:"createdByType"`
	LastModifiedAt     string `pulumi:"lastModifiedAt"`
	LastModifiedBy     string `pulumi:"lastModifiedBy"`
	LastModifiedByType string `pulumi:"lastModifiedByType"`
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
	pulumi.RegisterOutputType(AzureMonitorWorkspacePropertiesOutput{})
	pulumi.RegisterOutputType(AzureMonitorWorkspacePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AzureMonitorWorkspacePropertiesResponseOutput{})
	pulumi.RegisterOutputType(AzureMonitorWorkspacePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileResourcePropertiesOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileResourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileResourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(NotificationSettingsOutput{})
	pulumi.RegisterOutputType(NotificationSettingsPtrOutput{})
	pulumi.RegisterOutputType(NotificationSettingsResponseOutput{})
	pulumi.RegisterOutputType(NotificationSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityOutput{})
	pulumi.RegisterOutputType(ResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}

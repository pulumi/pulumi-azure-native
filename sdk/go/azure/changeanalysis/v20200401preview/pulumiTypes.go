


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





type AzureMonitorWorkspacePropertiesResponseInput interface {
	pulumi.Input

	ToAzureMonitorWorkspacePropertiesResponseOutput() AzureMonitorWorkspacePropertiesResponseOutput
	ToAzureMonitorWorkspacePropertiesResponseOutputWithContext(context.Context) AzureMonitorWorkspacePropertiesResponseOutput
}

type AzureMonitorWorkspacePropertiesResponseArgs struct {
	IncludeChangeDetails pulumi.StringPtrInput `pulumi:"includeChangeDetails"`
	WorkspaceId          pulumi.StringPtrInput `pulumi:"workspaceId"`
	WorkspaceResourceId  pulumi.StringPtrInput `pulumi:"workspaceResourceId"`
}

func (AzureMonitorWorkspacePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMonitorWorkspacePropertiesResponse)(nil)).Elem()
}

func (i AzureMonitorWorkspacePropertiesResponseArgs) ToAzureMonitorWorkspacePropertiesResponseOutput() AzureMonitorWorkspacePropertiesResponseOutput {
	return i.ToAzureMonitorWorkspacePropertiesResponseOutputWithContext(context.Background())
}

func (i AzureMonitorWorkspacePropertiesResponseArgs) ToAzureMonitorWorkspacePropertiesResponseOutputWithContext(ctx context.Context) AzureMonitorWorkspacePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMonitorWorkspacePropertiesResponseOutput)
}

func (i AzureMonitorWorkspacePropertiesResponseArgs) ToAzureMonitorWorkspacePropertiesResponsePtrOutput() AzureMonitorWorkspacePropertiesResponsePtrOutput {
	return i.ToAzureMonitorWorkspacePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i AzureMonitorWorkspacePropertiesResponseArgs) ToAzureMonitorWorkspacePropertiesResponsePtrOutputWithContext(ctx context.Context) AzureMonitorWorkspacePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMonitorWorkspacePropertiesResponseOutput).ToAzureMonitorWorkspacePropertiesResponsePtrOutputWithContext(ctx)
}









type AzureMonitorWorkspacePropertiesResponsePtrInput interface {
	pulumi.Input

	ToAzureMonitorWorkspacePropertiesResponsePtrOutput() AzureMonitorWorkspacePropertiesResponsePtrOutput
	ToAzureMonitorWorkspacePropertiesResponsePtrOutputWithContext(context.Context) AzureMonitorWorkspacePropertiesResponsePtrOutput
}

type azureMonitorWorkspacePropertiesResponsePtrType AzureMonitorWorkspacePropertiesResponseArgs

func AzureMonitorWorkspacePropertiesResponsePtr(v *AzureMonitorWorkspacePropertiesResponseArgs) AzureMonitorWorkspacePropertiesResponsePtrInput {
	return (*azureMonitorWorkspacePropertiesResponsePtrType)(v)
}

func (*azureMonitorWorkspacePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureMonitorWorkspacePropertiesResponse)(nil)).Elem()
}

func (i *azureMonitorWorkspacePropertiesResponsePtrType) ToAzureMonitorWorkspacePropertiesResponsePtrOutput() AzureMonitorWorkspacePropertiesResponsePtrOutput {
	return i.ToAzureMonitorWorkspacePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *azureMonitorWorkspacePropertiesResponsePtrType) ToAzureMonitorWorkspacePropertiesResponsePtrOutputWithContext(ctx context.Context) AzureMonitorWorkspacePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMonitorWorkspacePropertiesResponsePtrOutput)
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

func (o AzureMonitorWorkspacePropertiesResponseOutput) ToAzureMonitorWorkspacePropertiesResponsePtrOutput() AzureMonitorWorkspacePropertiesResponsePtrOutput {
	return o.ToAzureMonitorWorkspacePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o AzureMonitorWorkspacePropertiesResponseOutput) ToAzureMonitorWorkspacePropertiesResponsePtrOutputWithContext(ctx context.Context) AzureMonitorWorkspacePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureMonitorWorkspacePropertiesResponse) *AzureMonitorWorkspacePropertiesResponse {
		return &v
	}).(AzureMonitorWorkspacePropertiesResponsePtrOutput)
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





type ConfigurationProfileResourcePropertiesResponseInput interface {
	pulumi.Input

	ToConfigurationProfileResourcePropertiesResponseOutput() ConfigurationProfileResourcePropertiesResponseOutput
	ToConfigurationProfileResourcePropertiesResponseOutputWithContext(context.Context) ConfigurationProfileResourcePropertiesResponseOutput
}

type ConfigurationProfileResourcePropertiesResponseArgs struct {
	Notifications NotificationSettingsResponsePtrInput `pulumi:"notifications"`
}

func (ConfigurationProfileResourcePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileResourcePropertiesResponse)(nil)).Elem()
}

func (i ConfigurationProfileResourcePropertiesResponseArgs) ToConfigurationProfileResourcePropertiesResponseOutput() ConfigurationProfileResourcePropertiesResponseOutput {
	return i.ToConfigurationProfileResourcePropertiesResponseOutputWithContext(context.Background())
}

func (i ConfigurationProfileResourcePropertiesResponseArgs) ToConfigurationProfileResourcePropertiesResponseOutputWithContext(ctx context.Context) ConfigurationProfileResourcePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileResourcePropertiesResponseOutput)
}

func (i ConfigurationProfileResourcePropertiesResponseArgs) ToConfigurationProfileResourcePropertiesResponsePtrOutput() ConfigurationProfileResourcePropertiesResponsePtrOutput {
	return i.ToConfigurationProfileResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ConfigurationProfileResourcePropertiesResponseArgs) ToConfigurationProfileResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfileResourcePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileResourcePropertiesResponseOutput).ToConfigurationProfileResourcePropertiesResponsePtrOutputWithContext(ctx)
}









type ConfigurationProfileResourcePropertiesResponsePtrInput interface {
	pulumi.Input

	ToConfigurationProfileResourcePropertiesResponsePtrOutput() ConfigurationProfileResourcePropertiesResponsePtrOutput
	ToConfigurationProfileResourcePropertiesResponsePtrOutputWithContext(context.Context) ConfigurationProfileResourcePropertiesResponsePtrOutput
}

type configurationProfileResourcePropertiesResponsePtrType ConfigurationProfileResourcePropertiesResponseArgs

func ConfigurationProfileResourcePropertiesResponsePtr(v *ConfigurationProfileResourcePropertiesResponseArgs) ConfigurationProfileResourcePropertiesResponsePtrInput {
	return (*configurationProfileResourcePropertiesResponsePtrType)(v)
}

func (*configurationProfileResourcePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileResourcePropertiesResponse)(nil)).Elem()
}

func (i *configurationProfileResourcePropertiesResponsePtrType) ToConfigurationProfileResourcePropertiesResponsePtrOutput() ConfigurationProfileResourcePropertiesResponsePtrOutput {
	return i.ToConfigurationProfileResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *configurationProfileResourcePropertiesResponsePtrType) ToConfigurationProfileResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfileResourcePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileResourcePropertiesResponsePtrOutput)
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

func (o ConfigurationProfileResourcePropertiesResponseOutput) ToConfigurationProfileResourcePropertiesResponsePtrOutput() ConfigurationProfileResourcePropertiesResponsePtrOutput {
	return o.ToConfigurationProfileResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ConfigurationProfileResourcePropertiesResponseOutput) ToConfigurationProfileResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfileResourcePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationProfileResourcePropertiesResponse) *ConfigurationProfileResourcePropertiesResponse {
		return &v
	}).(ConfigurationProfileResourcePropertiesResponsePtrOutput)
}

func (o ConfigurationProfileResourcePropertiesResponseOutput) Notifications() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v ConfigurationProfileResourcePropertiesResponse) *NotificationSettingsResponse {
		return v.Notifications
	}).(NotificationSettingsResponsePtrOutput)
}

type ConfigurationProfileResourcePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileResourcePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileResourcePropertiesResponse)(nil)).Elem()
}

func (o ConfigurationProfileResourcePropertiesResponsePtrOutput) ToConfigurationProfileResourcePropertiesResponsePtrOutput() ConfigurationProfileResourcePropertiesResponsePtrOutput {
	return o
}

func (o ConfigurationProfileResourcePropertiesResponsePtrOutput) ToConfigurationProfileResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfileResourcePropertiesResponsePtrOutput {
	return o
}

func (o ConfigurationProfileResourcePropertiesResponsePtrOutput) Elem() ConfigurationProfileResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfileResourcePropertiesResponse) ConfigurationProfileResourcePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ConfigurationProfileResourcePropertiesResponse
		return ret
	}).(ConfigurationProfileResourcePropertiesResponseOutput)
}

func (o ConfigurationProfileResourcePropertiesResponsePtrOutput) Notifications() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileResourcePropertiesResponse) *NotificationSettingsResponse {
		if v == nil {
			return nil
		}
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





type NotificationSettingsResponseInput interface {
	pulumi.Input

	ToNotificationSettingsResponseOutput() NotificationSettingsResponseOutput
	ToNotificationSettingsResponseOutputWithContext(context.Context) NotificationSettingsResponseOutput
}

type NotificationSettingsResponseArgs struct {
	ActivationState                 pulumi.StringPtrInput                           `pulumi:"activationState"`
	AzureMonitorWorkspaceProperties AzureMonitorWorkspacePropertiesResponsePtrInput `pulumi:"azureMonitorWorkspaceProperties"`
}

func (NotificationSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationSettingsResponse)(nil)).Elem()
}

func (i NotificationSettingsResponseArgs) ToNotificationSettingsResponseOutput() NotificationSettingsResponseOutput {
	return i.ToNotificationSettingsResponseOutputWithContext(context.Background())
}

func (i NotificationSettingsResponseArgs) ToNotificationSettingsResponseOutputWithContext(ctx context.Context) NotificationSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationSettingsResponseOutput)
}

func (i NotificationSettingsResponseArgs) ToNotificationSettingsResponsePtrOutput() NotificationSettingsResponsePtrOutput {
	return i.ToNotificationSettingsResponsePtrOutputWithContext(context.Background())
}

func (i NotificationSettingsResponseArgs) ToNotificationSettingsResponsePtrOutputWithContext(ctx context.Context) NotificationSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationSettingsResponseOutput).ToNotificationSettingsResponsePtrOutputWithContext(ctx)
}









type NotificationSettingsResponsePtrInput interface {
	pulumi.Input

	ToNotificationSettingsResponsePtrOutput() NotificationSettingsResponsePtrOutput
	ToNotificationSettingsResponsePtrOutputWithContext(context.Context) NotificationSettingsResponsePtrOutput
}

type notificationSettingsResponsePtrType NotificationSettingsResponseArgs

func NotificationSettingsResponsePtr(v *NotificationSettingsResponseArgs) NotificationSettingsResponsePtrInput {
	return (*notificationSettingsResponsePtrType)(v)
}

func (*notificationSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationSettingsResponse)(nil)).Elem()
}

func (i *notificationSettingsResponsePtrType) ToNotificationSettingsResponsePtrOutput() NotificationSettingsResponsePtrOutput {
	return i.ToNotificationSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *notificationSettingsResponsePtrType) ToNotificationSettingsResponsePtrOutputWithContext(ctx context.Context) NotificationSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationSettingsResponsePtrOutput)
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

func (o NotificationSettingsResponseOutput) ToNotificationSettingsResponsePtrOutput() NotificationSettingsResponsePtrOutput {
	return o.ToNotificationSettingsResponsePtrOutputWithContext(context.Background())
}

func (o NotificationSettingsResponseOutput) ToNotificationSettingsResponsePtrOutputWithContext(ctx context.Context) NotificationSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotificationSettingsResponse) *NotificationSettingsResponse {
		return &v
	}).(NotificationSettingsResponsePtrOutput)
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





type ResourceIdentityResponseInput interface {
	pulumi.Input

	ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput
	ToResourceIdentityResponseOutputWithContext(context.Context) ResourceIdentityResponseOutput
}

type ResourceIdentityResponseArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (ResourceIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityResponse)(nil)).Elem()
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput {
	return i.ToResourceIdentityResponseOutputWithContext(context.Background())
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponseOutputWithContext(ctx context.Context) ResourceIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponseOutput)
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return i.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponseOutput).ToResourceIdentityResponsePtrOutputWithContext(ctx)
}









type ResourceIdentityResponsePtrInput interface {
	pulumi.Input

	ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput
	ToResourceIdentityResponsePtrOutputWithContext(context.Context) ResourceIdentityResponsePtrOutput
}

type resourceIdentityResponsePtrType ResourceIdentityResponseArgs

func ResourceIdentityResponsePtr(v *ResourceIdentityResponseArgs) ResourceIdentityResponsePtrInput {
	return (*resourceIdentityResponsePtrType)(v)
}

func (*resourceIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityResponse)(nil)).Elem()
}

func (i *resourceIdentityResponsePtrType) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return i.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *resourceIdentityResponsePtrType) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponsePtrOutput)
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

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return o.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityResponse) *ResourceIdentityResponse {
		return &v
	}).(ResourceIdentityResponsePtrOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*AzureMonitorWorkspacePropertiesInput)(nil)).Elem(), AzureMonitorWorkspacePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureMonitorWorkspacePropertiesPtrInput)(nil)).Elem(), AzureMonitorWorkspacePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureMonitorWorkspacePropertiesResponseInput)(nil)).Elem(), AzureMonitorWorkspacePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureMonitorWorkspacePropertiesResponsePtrInput)(nil)).Elem(), AzureMonitorWorkspacePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationProfileResourcePropertiesInput)(nil)).Elem(), ConfigurationProfileResourcePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationProfileResourcePropertiesPtrInput)(nil)).Elem(), ConfigurationProfileResourcePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationProfileResourcePropertiesResponseInput)(nil)).Elem(), ConfigurationProfileResourcePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationProfileResourcePropertiesResponsePtrInput)(nil)).Elem(), ConfigurationProfileResourcePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationSettingsInput)(nil)).Elem(), NotificationSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationSettingsPtrInput)(nil)).Elem(), NotificationSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationSettingsResponseInput)(nil)).Elem(), NotificationSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationSettingsResponsePtrInput)(nil)).Elem(), NotificationSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityInput)(nil)).Elem(), ResourceIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityPtrInput)(nil)).Elem(), ResourceIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityResponseInput)(nil)).Elem(), ResourceIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityResponsePtrInput)(nil)).Elem(), ResourceIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterOutputType(AzureMonitorWorkspacePropertiesOutput{})
	pulumi.RegisterOutputType(AzureMonitorWorkspacePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AzureMonitorWorkspacePropertiesResponseOutput{})
	pulumi.RegisterOutputType(AzureMonitorWorkspacePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileResourcePropertiesOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileResourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileResourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileResourcePropertiesResponsePtrOutput{})
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

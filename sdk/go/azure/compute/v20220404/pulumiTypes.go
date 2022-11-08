


package v20220404

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CloudServiceExtensionProfile struct {
	Extensions []Extension `pulumi:"extensions"`
}





type CloudServiceExtensionProfileInput interface {
	pulumi.Input

	ToCloudServiceExtensionProfileOutput() CloudServiceExtensionProfileOutput
	ToCloudServiceExtensionProfileOutputWithContext(context.Context) CloudServiceExtensionProfileOutput
}

type CloudServiceExtensionProfileArgs struct {
	Extensions ExtensionArrayInput `pulumi:"extensions"`
}

func (CloudServiceExtensionProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceExtensionProfile)(nil)).Elem()
}

func (i CloudServiceExtensionProfileArgs) ToCloudServiceExtensionProfileOutput() CloudServiceExtensionProfileOutput {
	return i.ToCloudServiceExtensionProfileOutputWithContext(context.Background())
}

func (i CloudServiceExtensionProfileArgs) ToCloudServiceExtensionProfileOutputWithContext(ctx context.Context) CloudServiceExtensionProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceExtensionProfileOutput)
}

func (i CloudServiceExtensionProfileArgs) ToCloudServiceExtensionProfilePtrOutput() CloudServiceExtensionProfilePtrOutput {
	return i.ToCloudServiceExtensionProfilePtrOutputWithContext(context.Background())
}

func (i CloudServiceExtensionProfileArgs) ToCloudServiceExtensionProfilePtrOutputWithContext(ctx context.Context) CloudServiceExtensionProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceExtensionProfileOutput).ToCloudServiceExtensionProfilePtrOutputWithContext(ctx)
}









type CloudServiceExtensionProfilePtrInput interface {
	pulumi.Input

	ToCloudServiceExtensionProfilePtrOutput() CloudServiceExtensionProfilePtrOutput
	ToCloudServiceExtensionProfilePtrOutputWithContext(context.Context) CloudServiceExtensionProfilePtrOutput
}

type cloudServiceExtensionProfilePtrType CloudServiceExtensionProfileArgs

func CloudServiceExtensionProfilePtr(v *CloudServiceExtensionProfileArgs) CloudServiceExtensionProfilePtrInput {
	return (*cloudServiceExtensionProfilePtrType)(v)
}

func (*cloudServiceExtensionProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceExtensionProfile)(nil)).Elem()
}

func (i *cloudServiceExtensionProfilePtrType) ToCloudServiceExtensionProfilePtrOutput() CloudServiceExtensionProfilePtrOutput {
	return i.ToCloudServiceExtensionProfilePtrOutputWithContext(context.Background())
}

func (i *cloudServiceExtensionProfilePtrType) ToCloudServiceExtensionProfilePtrOutputWithContext(ctx context.Context) CloudServiceExtensionProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceExtensionProfilePtrOutput)
}

type CloudServiceExtensionProfileOutput struct{ *pulumi.OutputState }

func (CloudServiceExtensionProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceExtensionProfile)(nil)).Elem()
}

func (o CloudServiceExtensionProfileOutput) ToCloudServiceExtensionProfileOutput() CloudServiceExtensionProfileOutput {
	return o
}

func (o CloudServiceExtensionProfileOutput) ToCloudServiceExtensionProfileOutputWithContext(ctx context.Context) CloudServiceExtensionProfileOutput {
	return o
}

func (o CloudServiceExtensionProfileOutput) ToCloudServiceExtensionProfilePtrOutput() CloudServiceExtensionProfilePtrOutput {
	return o.ToCloudServiceExtensionProfilePtrOutputWithContext(context.Background())
}

func (o CloudServiceExtensionProfileOutput) ToCloudServiceExtensionProfilePtrOutputWithContext(ctx context.Context) CloudServiceExtensionProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudServiceExtensionProfile) *CloudServiceExtensionProfile {
		return &v
	}).(CloudServiceExtensionProfilePtrOutput)
}

func (o CloudServiceExtensionProfileOutput) Extensions() ExtensionArrayOutput {
	return o.ApplyT(func(v CloudServiceExtensionProfile) []Extension { return v.Extensions }).(ExtensionArrayOutput)
}

type CloudServiceExtensionProfilePtrOutput struct{ *pulumi.OutputState }

func (CloudServiceExtensionProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceExtensionProfile)(nil)).Elem()
}

func (o CloudServiceExtensionProfilePtrOutput) ToCloudServiceExtensionProfilePtrOutput() CloudServiceExtensionProfilePtrOutput {
	return o
}

func (o CloudServiceExtensionProfilePtrOutput) ToCloudServiceExtensionProfilePtrOutputWithContext(ctx context.Context) CloudServiceExtensionProfilePtrOutput {
	return o
}

func (o CloudServiceExtensionProfilePtrOutput) Elem() CloudServiceExtensionProfileOutput {
	return o.ApplyT(func(v *CloudServiceExtensionProfile) CloudServiceExtensionProfile {
		if v != nil {
			return *v
		}
		var ret CloudServiceExtensionProfile
		return ret
	}).(CloudServiceExtensionProfileOutput)
}

func (o CloudServiceExtensionProfilePtrOutput) Extensions() ExtensionArrayOutput {
	return o.ApplyT(func(v *CloudServiceExtensionProfile) []Extension {
		if v == nil {
			return nil
		}
		return v.Extensions
	}).(ExtensionArrayOutput)
}

type CloudServiceExtensionProfileResponse struct {
	Extensions []ExtensionResponse `pulumi:"extensions"`
}

type CloudServiceExtensionProfileResponseOutput struct{ *pulumi.OutputState }

func (CloudServiceExtensionProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceExtensionProfileResponse)(nil)).Elem()
}

func (o CloudServiceExtensionProfileResponseOutput) ToCloudServiceExtensionProfileResponseOutput() CloudServiceExtensionProfileResponseOutput {
	return o
}

func (o CloudServiceExtensionProfileResponseOutput) ToCloudServiceExtensionProfileResponseOutputWithContext(ctx context.Context) CloudServiceExtensionProfileResponseOutput {
	return o
}

func (o CloudServiceExtensionProfileResponseOutput) Extensions() ExtensionResponseArrayOutput {
	return o.ApplyT(func(v CloudServiceExtensionProfileResponse) []ExtensionResponse { return v.Extensions }).(ExtensionResponseArrayOutput)
}

type CloudServiceExtensionProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudServiceExtensionProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceExtensionProfileResponse)(nil)).Elem()
}

func (o CloudServiceExtensionProfileResponsePtrOutput) ToCloudServiceExtensionProfileResponsePtrOutput() CloudServiceExtensionProfileResponsePtrOutput {
	return o
}

func (o CloudServiceExtensionProfileResponsePtrOutput) ToCloudServiceExtensionProfileResponsePtrOutputWithContext(ctx context.Context) CloudServiceExtensionProfileResponsePtrOutput {
	return o
}

func (o CloudServiceExtensionProfileResponsePtrOutput) Elem() CloudServiceExtensionProfileResponseOutput {
	return o.ApplyT(func(v *CloudServiceExtensionProfileResponse) CloudServiceExtensionProfileResponse {
		if v != nil {
			return *v
		}
		var ret CloudServiceExtensionProfileResponse
		return ret
	}).(CloudServiceExtensionProfileResponseOutput)
}

func (o CloudServiceExtensionProfileResponsePtrOutput) Extensions() ExtensionResponseArrayOutput {
	return o.ApplyT(func(v *CloudServiceExtensionProfileResponse) []ExtensionResponse {
		if v == nil {
			return nil
		}
		return v.Extensions
	}).(ExtensionResponseArrayOutput)
}

type CloudServiceExtensionProperties struct {
	AutoUpgradeMinorVersion       *bool                                `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag                *string                              `pulumi:"forceUpdateTag"`
	ProtectedSettings             interface{}                          `pulumi:"protectedSettings"`
	ProtectedSettingsFromKeyVault *CloudServiceVaultAndSecretReference `pulumi:"protectedSettingsFromKeyVault"`
	Publisher                     *string                              `pulumi:"publisher"`
	RolesAppliedTo                []string                             `pulumi:"rolesAppliedTo"`
	Settings                      interface{}                          `pulumi:"settings"`
	Type                          *string                              `pulumi:"type"`
	TypeHandlerVersion            *string                              `pulumi:"typeHandlerVersion"`
}





type CloudServiceExtensionPropertiesInput interface {
	pulumi.Input

	ToCloudServiceExtensionPropertiesOutput() CloudServiceExtensionPropertiesOutput
	ToCloudServiceExtensionPropertiesOutputWithContext(context.Context) CloudServiceExtensionPropertiesOutput
}

type CloudServiceExtensionPropertiesArgs struct {
	AutoUpgradeMinorVersion       pulumi.BoolPtrInput                         `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag                pulumi.StringPtrInput                       `pulumi:"forceUpdateTag"`
	ProtectedSettings             pulumi.Input                                `pulumi:"protectedSettings"`
	ProtectedSettingsFromKeyVault CloudServiceVaultAndSecretReferencePtrInput `pulumi:"protectedSettingsFromKeyVault"`
	Publisher                     pulumi.StringPtrInput                       `pulumi:"publisher"`
	RolesAppliedTo                pulumi.StringArrayInput                     `pulumi:"rolesAppliedTo"`
	Settings                      pulumi.Input                                `pulumi:"settings"`
	Type                          pulumi.StringPtrInput                       `pulumi:"type"`
	TypeHandlerVersion            pulumi.StringPtrInput                       `pulumi:"typeHandlerVersion"`
}

func (CloudServiceExtensionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceExtensionProperties)(nil)).Elem()
}

func (i CloudServiceExtensionPropertiesArgs) ToCloudServiceExtensionPropertiesOutput() CloudServiceExtensionPropertiesOutput {
	return i.ToCloudServiceExtensionPropertiesOutputWithContext(context.Background())
}

func (i CloudServiceExtensionPropertiesArgs) ToCloudServiceExtensionPropertiesOutputWithContext(ctx context.Context) CloudServiceExtensionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceExtensionPropertiesOutput)
}

func (i CloudServiceExtensionPropertiesArgs) ToCloudServiceExtensionPropertiesPtrOutput() CloudServiceExtensionPropertiesPtrOutput {
	return i.ToCloudServiceExtensionPropertiesPtrOutputWithContext(context.Background())
}

func (i CloudServiceExtensionPropertiesArgs) ToCloudServiceExtensionPropertiesPtrOutputWithContext(ctx context.Context) CloudServiceExtensionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceExtensionPropertiesOutput).ToCloudServiceExtensionPropertiesPtrOutputWithContext(ctx)
}









type CloudServiceExtensionPropertiesPtrInput interface {
	pulumi.Input

	ToCloudServiceExtensionPropertiesPtrOutput() CloudServiceExtensionPropertiesPtrOutput
	ToCloudServiceExtensionPropertiesPtrOutputWithContext(context.Context) CloudServiceExtensionPropertiesPtrOutput
}

type cloudServiceExtensionPropertiesPtrType CloudServiceExtensionPropertiesArgs

func CloudServiceExtensionPropertiesPtr(v *CloudServiceExtensionPropertiesArgs) CloudServiceExtensionPropertiesPtrInput {
	return (*cloudServiceExtensionPropertiesPtrType)(v)
}

func (*cloudServiceExtensionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceExtensionProperties)(nil)).Elem()
}

func (i *cloudServiceExtensionPropertiesPtrType) ToCloudServiceExtensionPropertiesPtrOutput() CloudServiceExtensionPropertiesPtrOutput {
	return i.ToCloudServiceExtensionPropertiesPtrOutputWithContext(context.Background())
}

func (i *cloudServiceExtensionPropertiesPtrType) ToCloudServiceExtensionPropertiesPtrOutputWithContext(ctx context.Context) CloudServiceExtensionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceExtensionPropertiesPtrOutput)
}

type CloudServiceExtensionPropertiesOutput struct{ *pulumi.OutputState }

func (CloudServiceExtensionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceExtensionProperties)(nil)).Elem()
}

func (o CloudServiceExtensionPropertiesOutput) ToCloudServiceExtensionPropertiesOutput() CloudServiceExtensionPropertiesOutput {
	return o
}

func (o CloudServiceExtensionPropertiesOutput) ToCloudServiceExtensionPropertiesOutputWithContext(ctx context.Context) CloudServiceExtensionPropertiesOutput {
	return o
}

func (o CloudServiceExtensionPropertiesOutput) ToCloudServiceExtensionPropertiesPtrOutput() CloudServiceExtensionPropertiesPtrOutput {
	return o.ToCloudServiceExtensionPropertiesPtrOutputWithContext(context.Background())
}

func (o CloudServiceExtensionPropertiesOutput) ToCloudServiceExtensionPropertiesPtrOutputWithContext(ctx context.Context) CloudServiceExtensionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudServiceExtensionProperties) *CloudServiceExtensionProperties {
		return &v
	}).(CloudServiceExtensionPropertiesPtrOutput)
}

func (o CloudServiceExtensionPropertiesOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloudServiceExtensionProperties) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o CloudServiceExtensionPropertiesOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceExtensionProperties) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o CloudServiceExtensionPropertiesOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v CloudServiceExtensionProperties) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o CloudServiceExtensionPropertiesOutput) ProtectedSettingsFromKeyVault() CloudServiceVaultAndSecretReferencePtrOutput {
	return o.ApplyT(func(v CloudServiceExtensionProperties) *CloudServiceVaultAndSecretReference {
		return v.ProtectedSettingsFromKeyVault
	}).(CloudServiceVaultAndSecretReferencePtrOutput)
}

func (o CloudServiceExtensionPropertiesOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceExtensionProperties) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o CloudServiceExtensionPropertiesOutput) RolesAppliedTo() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CloudServiceExtensionProperties) []string { return v.RolesAppliedTo }).(pulumi.StringArrayOutput)
}

func (o CloudServiceExtensionPropertiesOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v CloudServiceExtensionProperties) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o CloudServiceExtensionPropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceExtensionProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o CloudServiceExtensionPropertiesOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceExtensionProperties) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type CloudServiceExtensionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CloudServiceExtensionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceExtensionProperties)(nil)).Elem()
}

func (o CloudServiceExtensionPropertiesPtrOutput) ToCloudServiceExtensionPropertiesPtrOutput() CloudServiceExtensionPropertiesPtrOutput {
	return o
}

func (o CloudServiceExtensionPropertiesPtrOutput) ToCloudServiceExtensionPropertiesPtrOutputWithContext(ctx context.Context) CloudServiceExtensionPropertiesPtrOutput {
	return o
}

func (o CloudServiceExtensionPropertiesPtrOutput) Elem() CloudServiceExtensionPropertiesOutput {
	return o.ApplyT(func(v *CloudServiceExtensionProperties) CloudServiceExtensionProperties {
		if v != nil {
			return *v
		}
		var ret CloudServiceExtensionProperties
		return ret
	}).(CloudServiceExtensionPropertiesOutput)
}

func (o CloudServiceExtensionPropertiesPtrOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloudServiceExtensionProperties) *bool {
		if v == nil {
			return nil
		}
		return v.AutoUpgradeMinorVersion
	}).(pulumi.BoolPtrOutput)
}

func (o CloudServiceExtensionPropertiesPtrOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceExtensionProperties) *string {
		if v == nil {
			return nil
		}
		return v.ForceUpdateTag
	}).(pulumi.StringPtrOutput)
}

func (o CloudServiceExtensionPropertiesPtrOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *CloudServiceExtensionProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.ProtectedSettings
	}).(pulumi.AnyOutput)
}

func (o CloudServiceExtensionPropertiesPtrOutput) ProtectedSettingsFromKeyVault() CloudServiceVaultAndSecretReferencePtrOutput {
	return o.ApplyT(func(v *CloudServiceExtensionProperties) *CloudServiceVaultAndSecretReference {
		if v == nil {
			return nil
		}
		return v.ProtectedSettingsFromKeyVault
	}).(CloudServiceVaultAndSecretReferencePtrOutput)
}

func (o CloudServiceExtensionPropertiesPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceExtensionProperties) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o CloudServiceExtensionPropertiesPtrOutput) RolesAppliedTo() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudServiceExtensionProperties) []string {
		if v == nil {
			return nil
		}
		return v.RolesAppliedTo
	}).(pulumi.StringArrayOutput)
}

func (o CloudServiceExtensionPropertiesPtrOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v *CloudServiceExtensionProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.Settings
	}).(pulumi.AnyOutput)
}

func (o CloudServiceExtensionPropertiesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceExtensionProperties) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o CloudServiceExtensionPropertiesPtrOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceExtensionProperties) *string {
		if v == nil {
			return nil
		}
		return v.TypeHandlerVersion
	}).(pulumi.StringPtrOutput)
}

type CloudServiceExtensionPropertiesResponse struct {
	AutoUpgradeMinorVersion       *bool                                        `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag                *string                                      `pulumi:"forceUpdateTag"`
	ProtectedSettings             interface{}                                  `pulumi:"protectedSettings"`
	ProtectedSettingsFromKeyVault *CloudServiceVaultAndSecretReferenceResponse `pulumi:"protectedSettingsFromKeyVault"`
	ProvisioningState             string                                       `pulumi:"provisioningState"`
	Publisher                     *string                                      `pulumi:"publisher"`
	RolesAppliedTo                []string                                     `pulumi:"rolesAppliedTo"`
	Settings                      interface{}                                  `pulumi:"settings"`
	Type                          *string                                      `pulumi:"type"`
	TypeHandlerVersion            *string                                      `pulumi:"typeHandlerVersion"`
}

type CloudServiceExtensionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CloudServiceExtensionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceExtensionPropertiesResponse)(nil)).Elem()
}

func (o CloudServiceExtensionPropertiesResponseOutput) ToCloudServiceExtensionPropertiesResponseOutput() CloudServiceExtensionPropertiesResponseOutput {
	return o
}

func (o CloudServiceExtensionPropertiesResponseOutput) ToCloudServiceExtensionPropertiesResponseOutputWithContext(ctx context.Context) CloudServiceExtensionPropertiesResponseOutput {
	return o
}

func (o CloudServiceExtensionPropertiesResponseOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloudServiceExtensionPropertiesResponse) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o CloudServiceExtensionPropertiesResponseOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceExtensionPropertiesResponse) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o CloudServiceExtensionPropertiesResponseOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v CloudServiceExtensionPropertiesResponse) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o CloudServiceExtensionPropertiesResponseOutput) ProtectedSettingsFromKeyVault() CloudServiceVaultAndSecretReferenceResponsePtrOutput {
	return o.ApplyT(func(v CloudServiceExtensionPropertiesResponse) *CloudServiceVaultAndSecretReferenceResponse {
		return v.ProtectedSettingsFromKeyVault
	}).(CloudServiceVaultAndSecretReferenceResponsePtrOutput)
}

func (o CloudServiceExtensionPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v CloudServiceExtensionPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CloudServiceExtensionPropertiesResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceExtensionPropertiesResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o CloudServiceExtensionPropertiesResponseOutput) RolesAppliedTo() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CloudServiceExtensionPropertiesResponse) []string { return v.RolesAppliedTo }).(pulumi.StringArrayOutput)
}

func (o CloudServiceExtensionPropertiesResponseOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v CloudServiceExtensionPropertiesResponse) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o CloudServiceExtensionPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceExtensionPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o CloudServiceExtensionPropertiesResponseOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceExtensionPropertiesResponse) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type CloudServiceExtensionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudServiceExtensionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceExtensionPropertiesResponse)(nil)).Elem()
}

func (o CloudServiceExtensionPropertiesResponsePtrOutput) ToCloudServiceExtensionPropertiesResponsePtrOutput() CloudServiceExtensionPropertiesResponsePtrOutput {
	return o
}

func (o CloudServiceExtensionPropertiesResponsePtrOutput) ToCloudServiceExtensionPropertiesResponsePtrOutputWithContext(ctx context.Context) CloudServiceExtensionPropertiesResponsePtrOutput {
	return o
}

func (o CloudServiceExtensionPropertiesResponsePtrOutput) Elem() CloudServiceExtensionPropertiesResponseOutput {
	return o.ApplyT(func(v *CloudServiceExtensionPropertiesResponse) CloudServiceExtensionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CloudServiceExtensionPropertiesResponse
		return ret
	}).(CloudServiceExtensionPropertiesResponseOutput)
}

func (o CloudServiceExtensionPropertiesResponsePtrOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloudServiceExtensionPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AutoUpgradeMinorVersion
	}).(pulumi.BoolPtrOutput)
}

func (o CloudServiceExtensionPropertiesResponsePtrOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceExtensionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ForceUpdateTag
	}).(pulumi.StringPtrOutput)
}

func (o CloudServiceExtensionPropertiesResponsePtrOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *CloudServiceExtensionPropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.ProtectedSettings
	}).(pulumi.AnyOutput)
}

func (o CloudServiceExtensionPropertiesResponsePtrOutput) ProtectedSettingsFromKeyVault() CloudServiceVaultAndSecretReferenceResponsePtrOutput {
	return o.ApplyT(func(v *CloudServiceExtensionPropertiesResponse) *CloudServiceVaultAndSecretReferenceResponse {
		if v == nil {
			return nil
		}
		return v.ProtectedSettingsFromKeyVault
	}).(CloudServiceVaultAndSecretReferenceResponsePtrOutput)
}

func (o CloudServiceExtensionPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceExtensionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o CloudServiceExtensionPropertiesResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceExtensionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o CloudServiceExtensionPropertiesResponsePtrOutput) RolesAppliedTo() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudServiceExtensionPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.RolesAppliedTo
	}).(pulumi.StringArrayOutput)
}

func (o CloudServiceExtensionPropertiesResponsePtrOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v *CloudServiceExtensionPropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Settings
	}).(pulumi.AnyOutput)
}

func (o CloudServiceExtensionPropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceExtensionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o CloudServiceExtensionPropertiesResponsePtrOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceExtensionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TypeHandlerVersion
	}).(pulumi.StringPtrOutput)
}

type CloudServiceNetworkProfile struct {
	LoadBalancerConfigurations []LoadBalancerConfiguration `pulumi:"loadBalancerConfigurations"`
	SlotType                   *string                     `pulumi:"slotType"`
	SwappableCloudService      *SubResource                `pulumi:"swappableCloudService"`
}





type CloudServiceNetworkProfileInput interface {
	pulumi.Input

	ToCloudServiceNetworkProfileOutput() CloudServiceNetworkProfileOutput
	ToCloudServiceNetworkProfileOutputWithContext(context.Context) CloudServiceNetworkProfileOutput
}

type CloudServiceNetworkProfileArgs struct {
	LoadBalancerConfigurations LoadBalancerConfigurationArrayInput `pulumi:"loadBalancerConfigurations"`
	SlotType                   pulumi.StringPtrInput               `pulumi:"slotType"`
	SwappableCloudService      SubResourcePtrInput                 `pulumi:"swappableCloudService"`
}

func (CloudServiceNetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceNetworkProfile)(nil)).Elem()
}

func (i CloudServiceNetworkProfileArgs) ToCloudServiceNetworkProfileOutput() CloudServiceNetworkProfileOutput {
	return i.ToCloudServiceNetworkProfileOutputWithContext(context.Background())
}

func (i CloudServiceNetworkProfileArgs) ToCloudServiceNetworkProfileOutputWithContext(ctx context.Context) CloudServiceNetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceNetworkProfileOutput)
}

func (i CloudServiceNetworkProfileArgs) ToCloudServiceNetworkProfilePtrOutput() CloudServiceNetworkProfilePtrOutput {
	return i.ToCloudServiceNetworkProfilePtrOutputWithContext(context.Background())
}

func (i CloudServiceNetworkProfileArgs) ToCloudServiceNetworkProfilePtrOutputWithContext(ctx context.Context) CloudServiceNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceNetworkProfileOutput).ToCloudServiceNetworkProfilePtrOutputWithContext(ctx)
}









type CloudServiceNetworkProfilePtrInput interface {
	pulumi.Input

	ToCloudServiceNetworkProfilePtrOutput() CloudServiceNetworkProfilePtrOutput
	ToCloudServiceNetworkProfilePtrOutputWithContext(context.Context) CloudServiceNetworkProfilePtrOutput
}

type cloudServiceNetworkProfilePtrType CloudServiceNetworkProfileArgs

func CloudServiceNetworkProfilePtr(v *CloudServiceNetworkProfileArgs) CloudServiceNetworkProfilePtrInput {
	return (*cloudServiceNetworkProfilePtrType)(v)
}

func (*cloudServiceNetworkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceNetworkProfile)(nil)).Elem()
}

func (i *cloudServiceNetworkProfilePtrType) ToCloudServiceNetworkProfilePtrOutput() CloudServiceNetworkProfilePtrOutput {
	return i.ToCloudServiceNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *cloudServiceNetworkProfilePtrType) ToCloudServiceNetworkProfilePtrOutputWithContext(ctx context.Context) CloudServiceNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceNetworkProfilePtrOutput)
}

type CloudServiceNetworkProfileOutput struct{ *pulumi.OutputState }

func (CloudServiceNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceNetworkProfile)(nil)).Elem()
}

func (o CloudServiceNetworkProfileOutput) ToCloudServiceNetworkProfileOutput() CloudServiceNetworkProfileOutput {
	return o
}

func (o CloudServiceNetworkProfileOutput) ToCloudServiceNetworkProfileOutputWithContext(ctx context.Context) CloudServiceNetworkProfileOutput {
	return o
}

func (o CloudServiceNetworkProfileOutput) ToCloudServiceNetworkProfilePtrOutput() CloudServiceNetworkProfilePtrOutput {
	return o.ToCloudServiceNetworkProfilePtrOutputWithContext(context.Background())
}

func (o CloudServiceNetworkProfileOutput) ToCloudServiceNetworkProfilePtrOutputWithContext(ctx context.Context) CloudServiceNetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudServiceNetworkProfile) *CloudServiceNetworkProfile {
		return &v
	}).(CloudServiceNetworkProfilePtrOutput)
}

func (o CloudServiceNetworkProfileOutput) LoadBalancerConfigurations() LoadBalancerConfigurationArrayOutput {
	return o.ApplyT(func(v CloudServiceNetworkProfile) []LoadBalancerConfiguration { return v.LoadBalancerConfigurations }).(LoadBalancerConfigurationArrayOutput)
}

func (o CloudServiceNetworkProfileOutput) SlotType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceNetworkProfile) *string { return v.SlotType }).(pulumi.StringPtrOutput)
}

func (o CloudServiceNetworkProfileOutput) SwappableCloudService() SubResourcePtrOutput {
	return o.ApplyT(func(v CloudServiceNetworkProfile) *SubResource { return v.SwappableCloudService }).(SubResourcePtrOutput)
}

type CloudServiceNetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (CloudServiceNetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceNetworkProfile)(nil)).Elem()
}

func (o CloudServiceNetworkProfilePtrOutput) ToCloudServiceNetworkProfilePtrOutput() CloudServiceNetworkProfilePtrOutput {
	return o
}

func (o CloudServiceNetworkProfilePtrOutput) ToCloudServiceNetworkProfilePtrOutputWithContext(ctx context.Context) CloudServiceNetworkProfilePtrOutput {
	return o
}

func (o CloudServiceNetworkProfilePtrOutput) Elem() CloudServiceNetworkProfileOutput {
	return o.ApplyT(func(v *CloudServiceNetworkProfile) CloudServiceNetworkProfile {
		if v != nil {
			return *v
		}
		var ret CloudServiceNetworkProfile
		return ret
	}).(CloudServiceNetworkProfileOutput)
}

func (o CloudServiceNetworkProfilePtrOutput) LoadBalancerConfigurations() LoadBalancerConfigurationArrayOutput {
	return o.ApplyT(func(v *CloudServiceNetworkProfile) []LoadBalancerConfiguration {
		if v == nil {
			return nil
		}
		return v.LoadBalancerConfigurations
	}).(LoadBalancerConfigurationArrayOutput)
}

func (o CloudServiceNetworkProfilePtrOutput) SlotType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.SlotType
	}).(pulumi.StringPtrOutput)
}

func (o CloudServiceNetworkProfilePtrOutput) SwappableCloudService() SubResourcePtrOutput {
	return o.ApplyT(func(v *CloudServiceNetworkProfile) *SubResource {
		if v == nil {
			return nil
		}
		return v.SwappableCloudService
	}).(SubResourcePtrOutput)
}

type CloudServiceNetworkProfileResponse struct {
	LoadBalancerConfigurations []LoadBalancerConfigurationResponse `pulumi:"loadBalancerConfigurations"`
	SlotType                   *string                             `pulumi:"slotType"`
	SwappableCloudService      *SubResourceResponse                `pulumi:"swappableCloudService"`
}

type CloudServiceNetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (CloudServiceNetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceNetworkProfileResponse)(nil)).Elem()
}

func (o CloudServiceNetworkProfileResponseOutput) ToCloudServiceNetworkProfileResponseOutput() CloudServiceNetworkProfileResponseOutput {
	return o
}

func (o CloudServiceNetworkProfileResponseOutput) ToCloudServiceNetworkProfileResponseOutputWithContext(ctx context.Context) CloudServiceNetworkProfileResponseOutput {
	return o
}

func (o CloudServiceNetworkProfileResponseOutput) LoadBalancerConfigurations() LoadBalancerConfigurationResponseArrayOutput {
	return o.ApplyT(func(v CloudServiceNetworkProfileResponse) []LoadBalancerConfigurationResponse {
		return v.LoadBalancerConfigurations
	}).(LoadBalancerConfigurationResponseArrayOutput)
}

func (o CloudServiceNetworkProfileResponseOutput) SlotType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceNetworkProfileResponse) *string { return v.SlotType }).(pulumi.StringPtrOutput)
}

func (o CloudServiceNetworkProfileResponseOutput) SwappableCloudService() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v CloudServiceNetworkProfileResponse) *SubResourceResponse { return v.SwappableCloudService }).(SubResourceResponsePtrOutput)
}

type CloudServiceNetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudServiceNetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceNetworkProfileResponse)(nil)).Elem()
}

func (o CloudServiceNetworkProfileResponsePtrOutput) ToCloudServiceNetworkProfileResponsePtrOutput() CloudServiceNetworkProfileResponsePtrOutput {
	return o
}

func (o CloudServiceNetworkProfileResponsePtrOutput) ToCloudServiceNetworkProfileResponsePtrOutputWithContext(ctx context.Context) CloudServiceNetworkProfileResponsePtrOutput {
	return o
}

func (o CloudServiceNetworkProfileResponsePtrOutput) Elem() CloudServiceNetworkProfileResponseOutput {
	return o.ApplyT(func(v *CloudServiceNetworkProfileResponse) CloudServiceNetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret CloudServiceNetworkProfileResponse
		return ret
	}).(CloudServiceNetworkProfileResponseOutput)
}

func (o CloudServiceNetworkProfileResponsePtrOutput) LoadBalancerConfigurations() LoadBalancerConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *CloudServiceNetworkProfileResponse) []LoadBalancerConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.LoadBalancerConfigurations
	}).(LoadBalancerConfigurationResponseArrayOutput)
}

func (o CloudServiceNetworkProfileResponsePtrOutput) SlotType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.SlotType
	}).(pulumi.StringPtrOutput)
}

func (o CloudServiceNetworkProfileResponsePtrOutput) SwappableCloudService() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *CloudServiceNetworkProfileResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return v.SwappableCloudService
	}).(SubResourceResponsePtrOutput)
}

type CloudServiceOsProfile struct {
	Secrets []CloudServiceVaultSecretGroup `pulumi:"secrets"`
}





type CloudServiceOsProfileInput interface {
	pulumi.Input

	ToCloudServiceOsProfileOutput() CloudServiceOsProfileOutput
	ToCloudServiceOsProfileOutputWithContext(context.Context) CloudServiceOsProfileOutput
}

type CloudServiceOsProfileArgs struct {
	Secrets CloudServiceVaultSecretGroupArrayInput `pulumi:"secrets"`
}

func (CloudServiceOsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceOsProfile)(nil)).Elem()
}

func (i CloudServiceOsProfileArgs) ToCloudServiceOsProfileOutput() CloudServiceOsProfileOutput {
	return i.ToCloudServiceOsProfileOutputWithContext(context.Background())
}

func (i CloudServiceOsProfileArgs) ToCloudServiceOsProfileOutputWithContext(ctx context.Context) CloudServiceOsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceOsProfileOutput)
}

func (i CloudServiceOsProfileArgs) ToCloudServiceOsProfilePtrOutput() CloudServiceOsProfilePtrOutput {
	return i.ToCloudServiceOsProfilePtrOutputWithContext(context.Background())
}

func (i CloudServiceOsProfileArgs) ToCloudServiceOsProfilePtrOutputWithContext(ctx context.Context) CloudServiceOsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceOsProfileOutput).ToCloudServiceOsProfilePtrOutputWithContext(ctx)
}









type CloudServiceOsProfilePtrInput interface {
	pulumi.Input

	ToCloudServiceOsProfilePtrOutput() CloudServiceOsProfilePtrOutput
	ToCloudServiceOsProfilePtrOutputWithContext(context.Context) CloudServiceOsProfilePtrOutput
}

type cloudServiceOsProfilePtrType CloudServiceOsProfileArgs

func CloudServiceOsProfilePtr(v *CloudServiceOsProfileArgs) CloudServiceOsProfilePtrInput {
	return (*cloudServiceOsProfilePtrType)(v)
}

func (*cloudServiceOsProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceOsProfile)(nil)).Elem()
}

func (i *cloudServiceOsProfilePtrType) ToCloudServiceOsProfilePtrOutput() CloudServiceOsProfilePtrOutput {
	return i.ToCloudServiceOsProfilePtrOutputWithContext(context.Background())
}

func (i *cloudServiceOsProfilePtrType) ToCloudServiceOsProfilePtrOutputWithContext(ctx context.Context) CloudServiceOsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceOsProfilePtrOutput)
}

type CloudServiceOsProfileOutput struct{ *pulumi.OutputState }

func (CloudServiceOsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceOsProfile)(nil)).Elem()
}

func (o CloudServiceOsProfileOutput) ToCloudServiceOsProfileOutput() CloudServiceOsProfileOutput {
	return o
}

func (o CloudServiceOsProfileOutput) ToCloudServiceOsProfileOutputWithContext(ctx context.Context) CloudServiceOsProfileOutput {
	return o
}

func (o CloudServiceOsProfileOutput) ToCloudServiceOsProfilePtrOutput() CloudServiceOsProfilePtrOutput {
	return o.ToCloudServiceOsProfilePtrOutputWithContext(context.Background())
}

func (o CloudServiceOsProfileOutput) ToCloudServiceOsProfilePtrOutputWithContext(ctx context.Context) CloudServiceOsProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudServiceOsProfile) *CloudServiceOsProfile {
		return &v
	}).(CloudServiceOsProfilePtrOutput)
}

func (o CloudServiceOsProfileOutput) Secrets() CloudServiceVaultSecretGroupArrayOutput {
	return o.ApplyT(func(v CloudServiceOsProfile) []CloudServiceVaultSecretGroup { return v.Secrets }).(CloudServiceVaultSecretGroupArrayOutput)
}

type CloudServiceOsProfilePtrOutput struct{ *pulumi.OutputState }

func (CloudServiceOsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceOsProfile)(nil)).Elem()
}

func (o CloudServiceOsProfilePtrOutput) ToCloudServiceOsProfilePtrOutput() CloudServiceOsProfilePtrOutput {
	return o
}

func (o CloudServiceOsProfilePtrOutput) ToCloudServiceOsProfilePtrOutputWithContext(ctx context.Context) CloudServiceOsProfilePtrOutput {
	return o
}

func (o CloudServiceOsProfilePtrOutput) Elem() CloudServiceOsProfileOutput {
	return o.ApplyT(func(v *CloudServiceOsProfile) CloudServiceOsProfile {
		if v != nil {
			return *v
		}
		var ret CloudServiceOsProfile
		return ret
	}).(CloudServiceOsProfileOutput)
}

func (o CloudServiceOsProfilePtrOutput) Secrets() CloudServiceVaultSecretGroupArrayOutput {
	return o.ApplyT(func(v *CloudServiceOsProfile) []CloudServiceVaultSecretGroup {
		if v == nil {
			return nil
		}
		return v.Secrets
	}).(CloudServiceVaultSecretGroupArrayOutput)
}

type CloudServiceOsProfileResponse struct {
	Secrets []CloudServiceVaultSecretGroupResponse `pulumi:"secrets"`
}

type CloudServiceOsProfileResponseOutput struct{ *pulumi.OutputState }

func (CloudServiceOsProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceOsProfileResponse)(nil)).Elem()
}

func (o CloudServiceOsProfileResponseOutput) ToCloudServiceOsProfileResponseOutput() CloudServiceOsProfileResponseOutput {
	return o
}

func (o CloudServiceOsProfileResponseOutput) ToCloudServiceOsProfileResponseOutputWithContext(ctx context.Context) CloudServiceOsProfileResponseOutput {
	return o
}

func (o CloudServiceOsProfileResponseOutput) Secrets() CloudServiceVaultSecretGroupResponseArrayOutput {
	return o.ApplyT(func(v CloudServiceOsProfileResponse) []CloudServiceVaultSecretGroupResponse { return v.Secrets }).(CloudServiceVaultSecretGroupResponseArrayOutput)
}

type CloudServiceOsProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudServiceOsProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceOsProfileResponse)(nil)).Elem()
}

func (o CloudServiceOsProfileResponsePtrOutput) ToCloudServiceOsProfileResponsePtrOutput() CloudServiceOsProfileResponsePtrOutput {
	return o
}

func (o CloudServiceOsProfileResponsePtrOutput) ToCloudServiceOsProfileResponsePtrOutputWithContext(ctx context.Context) CloudServiceOsProfileResponsePtrOutput {
	return o
}

func (o CloudServiceOsProfileResponsePtrOutput) Elem() CloudServiceOsProfileResponseOutput {
	return o.ApplyT(func(v *CloudServiceOsProfileResponse) CloudServiceOsProfileResponse {
		if v != nil {
			return *v
		}
		var ret CloudServiceOsProfileResponse
		return ret
	}).(CloudServiceOsProfileResponseOutput)
}

func (o CloudServiceOsProfileResponsePtrOutput) Secrets() CloudServiceVaultSecretGroupResponseArrayOutput {
	return o.ApplyT(func(v *CloudServiceOsProfileResponse) []CloudServiceVaultSecretGroupResponse {
		if v == nil {
			return nil
		}
		return v.Secrets
	}).(CloudServiceVaultSecretGroupResponseArrayOutput)
}

type CloudServiceProperties struct {
	AllowModelOverride *bool                         `pulumi:"allowModelOverride"`
	Configuration      *string                       `pulumi:"configuration"`
	ConfigurationUrl   *string                       `pulumi:"configurationUrl"`
	ExtensionProfile   *CloudServiceExtensionProfile `pulumi:"extensionProfile"`
	NetworkProfile     *CloudServiceNetworkProfile   `pulumi:"networkProfile"`
	OsProfile          *CloudServiceOsProfile        `pulumi:"osProfile"`
	PackageUrl         *string                       `pulumi:"packageUrl"`
	RoleProfile        *CloudServiceRoleProfile      `pulumi:"roleProfile"`
	StartCloudService  *bool                         `pulumi:"startCloudService"`
	UpgradeMode        *string                       `pulumi:"upgradeMode"`
}





type CloudServicePropertiesInput interface {
	pulumi.Input

	ToCloudServicePropertiesOutput() CloudServicePropertiesOutput
	ToCloudServicePropertiesOutputWithContext(context.Context) CloudServicePropertiesOutput
}

type CloudServicePropertiesArgs struct {
	AllowModelOverride pulumi.BoolPtrInput                  `pulumi:"allowModelOverride"`
	Configuration      pulumi.StringPtrInput                `pulumi:"configuration"`
	ConfigurationUrl   pulumi.StringPtrInput                `pulumi:"configurationUrl"`
	ExtensionProfile   CloudServiceExtensionProfilePtrInput `pulumi:"extensionProfile"`
	NetworkProfile     CloudServiceNetworkProfilePtrInput   `pulumi:"networkProfile"`
	OsProfile          CloudServiceOsProfilePtrInput        `pulumi:"osProfile"`
	PackageUrl         pulumi.StringPtrInput                `pulumi:"packageUrl"`
	RoleProfile        CloudServiceRoleProfilePtrInput      `pulumi:"roleProfile"`
	StartCloudService  pulumi.BoolPtrInput                  `pulumi:"startCloudService"`
	UpgradeMode        pulumi.StringPtrInput                `pulumi:"upgradeMode"`
}

func (CloudServicePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceProperties)(nil)).Elem()
}

func (i CloudServicePropertiesArgs) ToCloudServicePropertiesOutput() CloudServicePropertiesOutput {
	return i.ToCloudServicePropertiesOutputWithContext(context.Background())
}

func (i CloudServicePropertiesArgs) ToCloudServicePropertiesOutputWithContext(ctx context.Context) CloudServicePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServicePropertiesOutput)
}

func (i CloudServicePropertiesArgs) ToCloudServicePropertiesPtrOutput() CloudServicePropertiesPtrOutput {
	return i.ToCloudServicePropertiesPtrOutputWithContext(context.Background())
}

func (i CloudServicePropertiesArgs) ToCloudServicePropertiesPtrOutputWithContext(ctx context.Context) CloudServicePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServicePropertiesOutput).ToCloudServicePropertiesPtrOutputWithContext(ctx)
}









type CloudServicePropertiesPtrInput interface {
	pulumi.Input

	ToCloudServicePropertiesPtrOutput() CloudServicePropertiesPtrOutput
	ToCloudServicePropertiesPtrOutputWithContext(context.Context) CloudServicePropertiesPtrOutput
}

type cloudServicePropertiesPtrType CloudServicePropertiesArgs

func CloudServicePropertiesPtr(v *CloudServicePropertiesArgs) CloudServicePropertiesPtrInput {
	return (*cloudServicePropertiesPtrType)(v)
}

func (*cloudServicePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceProperties)(nil)).Elem()
}

func (i *cloudServicePropertiesPtrType) ToCloudServicePropertiesPtrOutput() CloudServicePropertiesPtrOutput {
	return i.ToCloudServicePropertiesPtrOutputWithContext(context.Background())
}

func (i *cloudServicePropertiesPtrType) ToCloudServicePropertiesPtrOutputWithContext(ctx context.Context) CloudServicePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServicePropertiesPtrOutput)
}

type CloudServicePropertiesOutput struct{ *pulumi.OutputState }

func (CloudServicePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceProperties)(nil)).Elem()
}

func (o CloudServicePropertiesOutput) ToCloudServicePropertiesOutput() CloudServicePropertiesOutput {
	return o
}

func (o CloudServicePropertiesOutput) ToCloudServicePropertiesOutputWithContext(ctx context.Context) CloudServicePropertiesOutput {
	return o
}

func (o CloudServicePropertiesOutput) ToCloudServicePropertiesPtrOutput() CloudServicePropertiesPtrOutput {
	return o.ToCloudServicePropertiesPtrOutputWithContext(context.Background())
}

func (o CloudServicePropertiesOutput) ToCloudServicePropertiesPtrOutputWithContext(ctx context.Context) CloudServicePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudServiceProperties) *CloudServiceProperties {
		return &v
	}).(CloudServicePropertiesPtrOutput)
}

func (o CloudServicePropertiesOutput) AllowModelOverride() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloudServiceProperties) *bool { return v.AllowModelOverride }).(pulumi.BoolPtrOutput)
}

func (o CloudServicePropertiesOutput) Configuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceProperties) *string { return v.Configuration }).(pulumi.StringPtrOutput)
}

func (o CloudServicePropertiesOutput) ConfigurationUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceProperties) *string { return v.ConfigurationUrl }).(pulumi.StringPtrOutput)
}

func (o CloudServicePropertiesOutput) ExtensionProfile() CloudServiceExtensionProfilePtrOutput {
	return o.ApplyT(func(v CloudServiceProperties) *CloudServiceExtensionProfile { return v.ExtensionProfile }).(CloudServiceExtensionProfilePtrOutput)
}

func (o CloudServicePropertiesOutput) NetworkProfile() CloudServiceNetworkProfilePtrOutput {
	return o.ApplyT(func(v CloudServiceProperties) *CloudServiceNetworkProfile { return v.NetworkProfile }).(CloudServiceNetworkProfilePtrOutput)
}

func (o CloudServicePropertiesOutput) OsProfile() CloudServiceOsProfilePtrOutput {
	return o.ApplyT(func(v CloudServiceProperties) *CloudServiceOsProfile { return v.OsProfile }).(CloudServiceOsProfilePtrOutput)
}

func (o CloudServicePropertiesOutput) PackageUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceProperties) *string { return v.PackageUrl }).(pulumi.StringPtrOutput)
}

func (o CloudServicePropertiesOutput) RoleProfile() CloudServiceRoleProfilePtrOutput {
	return o.ApplyT(func(v CloudServiceProperties) *CloudServiceRoleProfile { return v.RoleProfile }).(CloudServiceRoleProfilePtrOutput)
}

func (o CloudServicePropertiesOutput) StartCloudService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloudServiceProperties) *bool { return v.StartCloudService }).(pulumi.BoolPtrOutput)
}

func (o CloudServicePropertiesOutput) UpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceProperties) *string { return v.UpgradeMode }).(pulumi.StringPtrOutput)
}

type CloudServicePropertiesPtrOutput struct{ *pulumi.OutputState }

func (CloudServicePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceProperties)(nil)).Elem()
}

func (o CloudServicePropertiesPtrOutput) ToCloudServicePropertiesPtrOutput() CloudServicePropertiesPtrOutput {
	return o
}

func (o CloudServicePropertiesPtrOutput) ToCloudServicePropertiesPtrOutputWithContext(ctx context.Context) CloudServicePropertiesPtrOutput {
	return o
}

func (o CloudServicePropertiesPtrOutput) Elem() CloudServicePropertiesOutput {
	return o.ApplyT(func(v *CloudServiceProperties) CloudServiceProperties {
		if v != nil {
			return *v
		}
		var ret CloudServiceProperties
		return ret
	}).(CloudServicePropertiesOutput)
}

func (o CloudServicePropertiesPtrOutput) AllowModelOverride() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloudServiceProperties) *bool {
		if v == nil {
			return nil
		}
		return v.AllowModelOverride
	}).(pulumi.BoolPtrOutput)
}

func (o CloudServicePropertiesPtrOutput) Configuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(pulumi.StringPtrOutput)
}

func (o CloudServicePropertiesPtrOutput) ConfigurationUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ConfigurationUrl
	}).(pulumi.StringPtrOutput)
}

func (o CloudServicePropertiesPtrOutput) ExtensionProfile() CloudServiceExtensionProfilePtrOutput {
	return o.ApplyT(func(v *CloudServiceProperties) *CloudServiceExtensionProfile {
		if v == nil {
			return nil
		}
		return v.ExtensionProfile
	}).(CloudServiceExtensionProfilePtrOutput)
}

func (o CloudServicePropertiesPtrOutput) NetworkProfile() CloudServiceNetworkProfilePtrOutput {
	return o.ApplyT(func(v *CloudServiceProperties) *CloudServiceNetworkProfile {
		if v == nil {
			return nil
		}
		return v.NetworkProfile
	}).(CloudServiceNetworkProfilePtrOutput)
}

func (o CloudServicePropertiesPtrOutput) OsProfile() CloudServiceOsProfilePtrOutput {
	return o.ApplyT(func(v *CloudServiceProperties) *CloudServiceOsProfile {
		if v == nil {
			return nil
		}
		return v.OsProfile
	}).(CloudServiceOsProfilePtrOutput)
}

func (o CloudServicePropertiesPtrOutput) PackageUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceProperties) *string {
		if v == nil {
			return nil
		}
		return v.PackageUrl
	}).(pulumi.StringPtrOutput)
}

func (o CloudServicePropertiesPtrOutput) RoleProfile() CloudServiceRoleProfilePtrOutput {
	return o.ApplyT(func(v *CloudServiceProperties) *CloudServiceRoleProfile {
		if v == nil {
			return nil
		}
		return v.RoleProfile
	}).(CloudServiceRoleProfilePtrOutput)
}

func (o CloudServicePropertiesPtrOutput) StartCloudService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloudServiceProperties) *bool {
		if v == nil {
			return nil
		}
		return v.StartCloudService
	}).(pulumi.BoolPtrOutput)
}

func (o CloudServicePropertiesPtrOutput) UpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceProperties) *string {
		if v == nil {
			return nil
		}
		return v.UpgradeMode
	}).(pulumi.StringPtrOutput)
}

type CloudServicePropertiesResponse struct {
	AllowModelOverride *bool                                 `pulumi:"allowModelOverride"`
	Configuration      *string                               `pulumi:"configuration"`
	ConfigurationUrl   *string                               `pulumi:"configurationUrl"`
	ExtensionProfile   *CloudServiceExtensionProfileResponse `pulumi:"extensionProfile"`
	NetworkProfile     *CloudServiceNetworkProfileResponse   `pulumi:"networkProfile"`
	OsProfile          *CloudServiceOsProfileResponse        `pulumi:"osProfile"`
	PackageUrl         *string                               `pulumi:"packageUrl"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	RoleProfile        *CloudServiceRoleProfileResponse      `pulumi:"roleProfile"`
	StartCloudService  *bool                                 `pulumi:"startCloudService"`
	UniqueId           string                                `pulumi:"uniqueId"`
	UpgradeMode        *string                               `pulumi:"upgradeMode"`
}

type CloudServicePropertiesResponseOutput struct{ *pulumi.OutputState }

func (CloudServicePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServicePropertiesResponse)(nil)).Elem()
}

func (o CloudServicePropertiesResponseOutput) ToCloudServicePropertiesResponseOutput() CloudServicePropertiesResponseOutput {
	return o
}

func (o CloudServicePropertiesResponseOutput) ToCloudServicePropertiesResponseOutputWithContext(ctx context.Context) CloudServicePropertiesResponseOutput {
	return o
}

func (o CloudServicePropertiesResponseOutput) AllowModelOverride() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloudServicePropertiesResponse) *bool { return v.AllowModelOverride }).(pulumi.BoolPtrOutput)
}

func (o CloudServicePropertiesResponseOutput) Configuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServicePropertiesResponse) *string { return v.Configuration }).(pulumi.StringPtrOutput)
}

func (o CloudServicePropertiesResponseOutput) ConfigurationUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServicePropertiesResponse) *string { return v.ConfigurationUrl }).(pulumi.StringPtrOutput)
}

func (o CloudServicePropertiesResponseOutput) ExtensionProfile() CloudServiceExtensionProfileResponsePtrOutput {
	return o.ApplyT(func(v CloudServicePropertiesResponse) *CloudServiceExtensionProfileResponse {
		return v.ExtensionProfile
	}).(CloudServiceExtensionProfileResponsePtrOutput)
}

func (o CloudServicePropertiesResponseOutput) NetworkProfile() CloudServiceNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v CloudServicePropertiesResponse) *CloudServiceNetworkProfileResponse { return v.NetworkProfile }).(CloudServiceNetworkProfileResponsePtrOutput)
}

func (o CloudServicePropertiesResponseOutput) OsProfile() CloudServiceOsProfileResponsePtrOutput {
	return o.ApplyT(func(v CloudServicePropertiesResponse) *CloudServiceOsProfileResponse { return v.OsProfile }).(CloudServiceOsProfileResponsePtrOutput)
}

func (o CloudServicePropertiesResponseOutput) PackageUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServicePropertiesResponse) *string { return v.PackageUrl }).(pulumi.StringPtrOutput)
}

func (o CloudServicePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v CloudServicePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CloudServicePropertiesResponseOutput) RoleProfile() CloudServiceRoleProfileResponsePtrOutput {
	return o.ApplyT(func(v CloudServicePropertiesResponse) *CloudServiceRoleProfileResponse { return v.RoleProfile }).(CloudServiceRoleProfileResponsePtrOutput)
}

func (o CloudServicePropertiesResponseOutput) StartCloudService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloudServicePropertiesResponse) *bool { return v.StartCloudService }).(pulumi.BoolPtrOutput)
}

func (o CloudServicePropertiesResponseOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v CloudServicePropertiesResponse) string { return v.UniqueId }).(pulumi.StringOutput)
}

func (o CloudServicePropertiesResponseOutput) UpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServicePropertiesResponse) *string { return v.UpgradeMode }).(pulumi.StringPtrOutput)
}

type CloudServiceRoleProfile struct {
	Roles []CloudServiceRoleProfileProperties `pulumi:"roles"`
}





type CloudServiceRoleProfileInput interface {
	pulumi.Input

	ToCloudServiceRoleProfileOutput() CloudServiceRoleProfileOutput
	ToCloudServiceRoleProfileOutputWithContext(context.Context) CloudServiceRoleProfileOutput
}

type CloudServiceRoleProfileArgs struct {
	Roles CloudServiceRoleProfilePropertiesArrayInput `pulumi:"roles"`
}

func (CloudServiceRoleProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceRoleProfile)(nil)).Elem()
}

func (i CloudServiceRoleProfileArgs) ToCloudServiceRoleProfileOutput() CloudServiceRoleProfileOutput {
	return i.ToCloudServiceRoleProfileOutputWithContext(context.Background())
}

func (i CloudServiceRoleProfileArgs) ToCloudServiceRoleProfileOutputWithContext(ctx context.Context) CloudServiceRoleProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceRoleProfileOutput)
}

func (i CloudServiceRoleProfileArgs) ToCloudServiceRoleProfilePtrOutput() CloudServiceRoleProfilePtrOutput {
	return i.ToCloudServiceRoleProfilePtrOutputWithContext(context.Background())
}

func (i CloudServiceRoleProfileArgs) ToCloudServiceRoleProfilePtrOutputWithContext(ctx context.Context) CloudServiceRoleProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceRoleProfileOutput).ToCloudServiceRoleProfilePtrOutputWithContext(ctx)
}









type CloudServiceRoleProfilePtrInput interface {
	pulumi.Input

	ToCloudServiceRoleProfilePtrOutput() CloudServiceRoleProfilePtrOutput
	ToCloudServiceRoleProfilePtrOutputWithContext(context.Context) CloudServiceRoleProfilePtrOutput
}

type cloudServiceRoleProfilePtrType CloudServiceRoleProfileArgs

func CloudServiceRoleProfilePtr(v *CloudServiceRoleProfileArgs) CloudServiceRoleProfilePtrInput {
	return (*cloudServiceRoleProfilePtrType)(v)
}

func (*cloudServiceRoleProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceRoleProfile)(nil)).Elem()
}

func (i *cloudServiceRoleProfilePtrType) ToCloudServiceRoleProfilePtrOutput() CloudServiceRoleProfilePtrOutput {
	return i.ToCloudServiceRoleProfilePtrOutputWithContext(context.Background())
}

func (i *cloudServiceRoleProfilePtrType) ToCloudServiceRoleProfilePtrOutputWithContext(ctx context.Context) CloudServiceRoleProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceRoleProfilePtrOutput)
}

type CloudServiceRoleProfileOutput struct{ *pulumi.OutputState }

func (CloudServiceRoleProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceRoleProfile)(nil)).Elem()
}

func (o CloudServiceRoleProfileOutput) ToCloudServiceRoleProfileOutput() CloudServiceRoleProfileOutput {
	return o
}

func (o CloudServiceRoleProfileOutput) ToCloudServiceRoleProfileOutputWithContext(ctx context.Context) CloudServiceRoleProfileOutput {
	return o
}

func (o CloudServiceRoleProfileOutput) ToCloudServiceRoleProfilePtrOutput() CloudServiceRoleProfilePtrOutput {
	return o.ToCloudServiceRoleProfilePtrOutputWithContext(context.Background())
}

func (o CloudServiceRoleProfileOutput) ToCloudServiceRoleProfilePtrOutputWithContext(ctx context.Context) CloudServiceRoleProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudServiceRoleProfile) *CloudServiceRoleProfile {
		return &v
	}).(CloudServiceRoleProfilePtrOutput)
}

func (o CloudServiceRoleProfileOutput) Roles() CloudServiceRoleProfilePropertiesArrayOutput {
	return o.ApplyT(func(v CloudServiceRoleProfile) []CloudServiceRoleProfileProperties { return v.Roles }).(CloudServiceRoleProfilePropertiesArrayOutput)
}

type CloudServiceRoleProfilePtrOutput struct{ *pulumi.OutputState }

func (CloudServiceRoleProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceRoleProfile)(nil)).Elem()
}

func (o CloudServiceRoleProfilePtrOutput) ToCloudServiceRoleProfilePtrOutput() CloudServiceRoleProfilePtrOutput {
	return o
}

func (o CloudServiceRoleProfilePtrOutput) ToCloudServiceRoleProfilePtrOutputWithContext(ctx context.Context) CloudServiceRoleProfilePtrOutput {
	return o
}

func (o CloudServiceRoleProfilePtrOutput) Elem() CloudServiceRoleProfileOutput {
	return o.ApplyT(func(v *CloudServiceRoleProfile) CloudServiceRoleProfile {
		if v != nil {
			return *v
		}
		var ret CloudServiceRoleProfile
		return ret
	}).(CloudServiceRoleProfileOutput)
}

func (o CloudServiceRoleProfilePtrOutput) Roles() CloudServiceRoleProfilePropertiesArrayOutput {
	return o.ApplyT(func(v *CloudServiceRoleProfile) []CloudServiceRoleProfileProperties {
		if v == nil {
			return nil
		}
		return v.Roles
	}).(CloudServiceRoleProfilePropertiesArrayOutput)
}

type CloudServiceRoleProfileProperties struct {
	Name *string              `pulumi:"name"`
	Sku  *CloudServiceRoleSku `pulumi:"sku"`
}





type CloudServiceRoleProfilePropertiesInput interface {
	pulumi.Input

	ToCloudServiceRoleProfilePropertiesOutput() CloudServiceRoleProfilePropertiesOutput
	ToCloudServiceRoleProfilePropertiesOutputWithContext(context.Context) CloudServiceRoleProfilePropertiesOutput
}

type CloudServiceRoleProfilePropertiesArgs struct {
	Name pulumi.StringPtrInput       `pulumi:"name"`
	Sku  CloudServiceRoleSkuPtrInput `pulumi:"sku"`
}

func (CloudServiceRoleProfilePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceRoleProfileProperties)(nil)).Elem()
}

func (i CloudServiceRoleProfilePropertiesArgs) ToCloudServiceRoleProfilePropertiesOutput() CloudServiceRoleProfilePropertiesOutput {
	return i.ToCloudServiceRoleProfilePropertiesOutputWithContext(context.Background())
}

func (i CloudServiceRoleProfilePropertiesArgs) ToCloudServiceRoleProfilePropertiesOutputWithContext(ctx context.Context) CloudServiceRoleProfilePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceRoleProfilePropertiesOutput)
}





type CloudServiceRoleProfilePropertiesArrayInput interface {
	pulumi.Input

	ToCloudServiceRoleProfilePropertiesArrayOutput() CloudServiceRoleProfilePropertiesArrayOutput
	ToCloudServiceRoleProfilePropertiesArrayOutputWithContext(context.Context) CloudServiceRoleProfilePropertiesArrayOutput
}

type CloudServiceRoleProfilePropertiesArray []CloudServiceRoleProfilePropertiesInput

func (CloudServiceRoleProfilePropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudServiceRoleProfileProperties)(nil)).Elem()
}

func (i CloudServiceRoleProfilePropertiesArray) ToCloudServiceRoleProfilePropertiesArrayOutput() CloudServiceRoleProfilePropertiesArrayOutput {
	return i.ToCloudServiceRoleProfilePropertiesArrayOutputWithContext(context.Background())
}

func (i CloudServiceRoleProfilePropertiesArray) ToCloudServiceRoleProfilePropertiesArrayOutputWithContext(ctx context.Context) CloudServiceRoleProfilePropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceRoleProfilePropertiesArrayOutput)
}

type CloudServiceRoleProfilePropertiesOutput struct{ *pulumi.OutputState }

func (CloudServiceRoleProfilePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceRoleProfileProperties)(nil)).Elem()
}

func (o CloudServiceRoleProfilePropertiesOutput) ToCloudServiceRoleProfilePropertiesOutput() CloudServiceRoleProfilePropertiesOutput {
	return o
}

func (o CloudServiceRoleProfilePropertiesOutput) ToCloudServiceRoleProfilePropertiesOutputWithContext(ctx context.Context) CloudServiceRoleProfilePropertiesOutput {
	return o
}

func (o CloudServiceRoleProfilePropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceRoleProfileProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CloudServiceRoleProfilePropertiesOutput) Sku() CloudServiceRoleSkuPtrOutput {
	return o.ApplyT(func(v CloudServiceRoleProfileProperties) *CloudServiceRoleSku { return v.Sku }).(CloudServiceRoleSkuPtrOutput)
}

type CloudServiceRoleProfilePropertiesArrayOutput struct{ *pulumi.OutputState }

func (CloudServiceRoleProfilePropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudServiceRoleProfileProperties)(nil)).Elem()
}

func (o CloudServiceRoleProfilePropertiesArrayOutput) ToCloudServiceRoleProfilePropertiesArrayOutput() CloudServiceRoleProfilePropertiesArrayOutput {
	return o
}

func (o CloudServiceRoleProfilePropertiesArrayOutput) ToCloudServiceRoleProfilePropertiesArrayOutputWithContext(ctx context.Context) CloudServiceRoleProfilePropertiesArrayOutput {
	return o
}

func (o CloudServiceRoleProfilePropertiesArrayOutput) Index(i pulumi.IntInput) CloudServiceRoleProfilePropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CloudServiceRoleProfileProperties {
		return vs[0].([]CloudServiceRoleProfileProperties)[vs[1].(int)]
	}).(CloudServiceRoleProfilePropertiesOutput)
}

type CloudServiceRoleProfilePropertiesResponse struct {
	Name *string                      `pulumi:"name"`
	Sku  *CloudServiceRoleSkuResponse `pulumi:"sku"`
}

type CloudServiceRoleProfilePropertiesResponseOutput struct{ *pulumi.OutputState }

func (CloudServiceRoleProfilePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceRoleProfilePropertiesResponse)(nil)).Elem()
}

func (o CloudServiceRoleProfilePropertiesResponseOutput) ToCloudServiceRoleProfilePropertiesResponseOutput() CloudServiceRoleProfilePropertiesResponseOutput {
	return o
}

func (o CloudServiceRoleProfilePropertiesResponseOutput) ToCloudServiceRoleProfilePropertiesResponseOutputWithContext(ctx context.Context) CloudServiceRoleProfilePropertiesResponseOutput {
	return o
}

func (o CloudServiceRoleProfilePropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceRoleProfilePropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CloudServiceRoleProfilePropertiesResponseOutput) Sku() CloudServiceRoleSkuResponsePtrOutput {
	return o.ApplyT(func(v CloudServiceRoleProfilePropertiesResponse) *CloudServiceRoleSkuResponse { return v.Sku }).(CloudServiceRoleSkuResponsePtrOutput)
}

type CloudServiceRoleProfilePropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (CloudServiceRoleProfilePropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudServiceRoleProfilePropertiesResponse)(nil)).Elem()
}

func (o CloudServiceRoleProfilePropertiesResponseArrayOutput) ToCloudServiceRoleProfilePropertiesResponseArrayOutput() CloudServiceRoleProfilePropertiesResponseArrayOutput {
	return o
}

func (o CloudServiceRoleProfilePropertiesResponseArrayOutput) ToCloudServiceRoleProfilePropertiesResponseArrayOutputWithContext(ctx context.Context) CloudServiceRoleProfilePropertiesResponseArrayOutput {
	return o
}

func (o CloudServiceRoleProfilePropertiesResponseArrayOutput) Index(i pulumi.IntInput) CloudServiceRoleProfilePropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CloudServiceRoleProfilePropertiesResponse {
		return vs[0].([]CloudServiceRoleProfilePropertiesResponse)[vs[1].(int)]
	}).(CloudServiceRoleProfilePropertiesResponseOutput)
}

type CloudServiceRoleProfileResponse struct {
	Roles []CloudServiceRoleProfilePropertiesResponse `pulumi:"roles"`
}

type CloudServiceRoleProfileResponseOutput struct{ *pulumi.OutputState }

func (CloudServiceRoleProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceRoleProfileResponse)(nil)).Elem()
}

func (o CloudServiceRoleProfileResponseOutput) ToCloudServiceRoleProfileResponseOutput() CloudServiceRoleProfileResponseOutput {
	return o
}

func (o CloudServiceRoleProfileResponseOutput) ToCloudServiceRoleProfileResponseOutputWithContext(ctx context.Context) CloudServiceRoleProfileResponseOutput {
	return o
}

func (o CloudServiceRoleProfileResponseOutput) Roles() CloudServiceRoleProfilePropertiesResponseArrayOutput {
	return o.ApplyT(func(v CloudServiceRoleProfileResponse) []CloudServiceRoleProfilePropertiesResponse { return v.Roles }).(CloudServiceRoleProfilePropertiesResponseArrayOutput)
}

type CloudServiceRoleProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudServiceRoleProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceRoleProfileResponse)(nil)).Elem()
}

func (o CloudServiceRoleProfileResponsePtrOutput) ToCloudServiceRoleProfileResponsePtrOutput() CloudServiceRoleProfileResponsePtrOutput {
	return o
}

func (o CloudServiceRoleProfileResponsePtrOutput) ToCloudServiceRoleProfileResponsePtrOutputWithContext(ctx context.Context) CloudServiceRoleProfileResponsePtrOutput {
	return o
}

func (o CloudServiceRoleProfileResponsePtrOutput) Elem() CloudServiceRoleProfileResponseOutput {
	return o.ApplyT(func(v *CloudServiceRoleProfileResponse) CloudServiceRoleProfileResponse {
		if v != nil {
			return *v
		}
		var ret CloudServiceRoleProfileResponse
		return ret
	}).(CloudServiceRoleProfileResponseOutput)
}

func (o CloudServiceRoleProfileResponsePtrOutput) Roles() CloudServiceRoleProfilePropertiesResponseArrayOutput {
	return o.ApplyT(func(v *CloudServiceRoleProfileResponse) []CloudServiceRoleProfilePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Roles
	}).(CloudServiceRoleProfilePropertiesResponseArrayOutput)
}

type CloudServiceRoleSku struct {
	Capacity *float64 `pulumi:"capacity"`
	Name     *string  `pulumi:"name"`
	Tier     *string  `pulumi:"tier"`
}





type CloudServiceRoleSkuInput interface {
	pulumi.Input

	ToCloudServiceRoleSkuOutput() CloudServiceRoleSkuOutput
	ToCloudServiceRoleSkuOutputWithContext(context.Context) CloudServiceRoleSkuOutput
}

type CloudServiceRoleSkuArgs struct {
	Capacity pulumi.Float64PtrInput `pulumi:"capacity"`
	Name     pulumi.StringPtrInput  `pulumi:"name"`
	Tier     pulumi.StringPtrInput  `pulumi:"tier"`
}

func (CloudServiceRoleSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceRoleSku)(nil)).Elem()
}

func (i CloudServiceRoleSkuArgs) ToCloudServiceRoleSkuOutput() CloudServiceRoleSkuOutput {
	return i.ToCloudServiceRoleSkuOutputWithContext(context.Background())
}

func (i CloudServiceRoleSkuArgs) ToCloudServiceRoleSkuOutputWithContext(ctx context.Context) CloudServiceRoleSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceRoleSkuOutput)
}

func (i CloudServiceRoleSkuArgs) ToCloudServiceRoleSkuPtrOutput() CloudServiceRoleSkuPtrOutput {
	return i.ToCloudServiceRoleSkuPtrOutputWithContext(context.Background())
}

func (i CloudServiceRoleSkuArgs) ToCloudServiceRoleSkuPtrOutputWithContext(ctx context.Context) CloudServiceRoleSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceRoleSkuOutput).ToCloudServiceRoleSkuPtrOutputWithContext(ctx)
}









type CloudServiceRoleSkuPtrInput interface {
	pulumi.Input

	ToCloudServiceRoleSkuPtrOutput() CloudServiceRoleSkuPtrOutput
	ToCloudServiceRoleSkuPtrOutputWithContext(context.Context) CloudServiceRoleSkuPtrOutput
}

type cloudServiceRoleSkuPtrType CloudServiceRoleSkuArgs

func CloudServiceRoleSkuPtr(v *CloudServiceRoleSkuArgs) CloudServiceRoleSkuPtrInput {
	return (*cloudServiceRoleSkuPtrType)(v)
}

func (*cloudServiceRoleSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceRoleSku)(nil)).Elem()
}

func (i *cloudServiceRoleSkuPtrType) ToCloudServiceRoleSkuPtrOutput() CloudServiceRoleSkuPtrOutput {
	return i.ToCloudServiceRoleSkuPtrOutputWithContext(context.Background())
}

func (i *cloudServiceRoleSkuPtrType) ToCloudServiceRoleSkuPtrOutputWithContext(ctx context.Context) CloudServiceRoleSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceRoleSkuPtrOutput)
}

type CloudServiceRoleSkuOutput struct{ *pulumi.OutputState }

func (CloudServiceRoleSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceRoleSku)(nil)).Elem()
}

func (o CloudServiceRoleSkuOutput) ToCloudServiceRoleSkuOutput() CloudServiceRoleSkuOutput {
	return o
}

func (o CloudServiceRoleSkuOutput) ToCloudServiceRoleSkuOutputWithContext(ctx context.Context) CloudServiceRoleSkuOutput {
	return o
}

func (o CloudServiceRoleSkuOutput) ToCloudServiceRoleSkuPtrOutput() CloudServiceRoleSkuPtrOutput {
	return o.ToCloudServiceRoleSkuPtrOutputWithContext(context.Background())
}

func (o CloudServiceRoleSkuOutput) ToCloudServiceRoleSkuPtrOutputWithContext(ctx context.Context) CloudServiceRoleSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudServiceRoleSku) *CloudServiceRoleSku {
		return &v
	}).(CloudServiceRoleSkuPtrOutput)
}

func (o CloudServiceRoleSkuOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v CloudServiceRoleSku) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o CloudServiceRoleSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceRoleSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CloudServiceRoleSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceRoleSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type CloudServiceRoleSkuPtrOutput struct{ *pulumi.OutputState }

func (CloudServiceRoleSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceRoleSku)(nil)).Elem()
}

func (o CloudServiceRoleSkuPtrOutput) ToCloudServiceRoleSkuPtrOutput() CloudServiceRoleSkuPtrOutput {
	return o
}

func (o CloudServiceRoleSkuPtrOutput) ToCloudServiceRoleSkuPtrOutputWithContext(ctx context.Context) CloudServiceRoleSkuPtrOutput {
	return o
}

func (o CloudServiceRoleSkuPtrOutput) Elem() CloudServiceRoleSkuOutput {
	return o.ApplyT(func(v *CloudServiceRoleSku) CloudServiceRoleSku {
		if v != nil {
			return *v
		}
		var ret CloudServiceRoleSku
		return ret
	}).(CloudServiceRoleSkuOutput)
}

func (o CloudServiceRoleSkuPtrOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CloudServiceRoleSku) *float64 {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.Float64PtrOutput)
}

func (o CloudServiceRoleSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceRoleSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o CloudServiceRoleSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceRoleSku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type CloudServiceRoleSkuResponse struct {
	Capacity *float64 `pulumi:"capacity"`
	Name     *string  `pulumi:"name"`
	Tier     *string  `pulumi:"tier"`
}

type CloudServiceRoleSkuResponseOutput struct{ *pulumi.OutputState }

func (CloudServiceRoleSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceRoleSkuResponse)(nil)).Elem()
}

func (o CloudServiceRoleSkuResponseOutput) ToCloudServiceRoleSkuResponseOutput() CloudServiceRoleSkuResponseOutput {
	return o
}

func (o CloudServiceRoleSkuResponseOutput) ToCloudServiceRoleSkuResponseOutputWithContext(ctx context.Context) CloudServiceRoleSkuResponseOutput {
	return o
}

func (o CloudServiceRoleSkuResponseOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v CloudServiceRoleSkuResponse) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o CloudServiceRoleSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceRoleSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CloudServiceRoleSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceRoleSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type CloudServiceRoleSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudServiceRoleSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceRoleSkuResponse)(nil)).Elem()
}

func (o CloudServiceRoleSkuResponsePtrOutput) ToCloudServiceRoleSkuResponsePtrOutput() CloudServiceRoleSkuResponsePtrOutput {
	return o
}

func (o CloudServiceRoleSkuResponsePtrOutput) ToCloudServiceRoleSkuResponsePtrOutputWithContext(ctx context.Context) CloudServiceRoleSkuResponsePtrOutput {
	return o
}

func (o CloudServiceRoleSkuResponsePtrOutput) Elem() CloudServiceRoleSkuResponseOutput {
	return o.ApplyT(func(v *CloudServiceRoleSkuResponse) CloudServiceRoleSkuResponse {
		if v != nil {
			return *v
		}
		var ret CloudServiceRoleSkuResponse
		return ret
	}).(CloudServiceRoleSkuResponseOutput)
}

func (o CloudServiceRoleSkuResponsePtrOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CloudServiceRoleSkuResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.Float64PtrOutput)
}

func (o CloudServiceRoleSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceRoleSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o CloudServiceRoleSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceRoleSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type CloudServiceVaultAndSecretReference struct {
	SecretUrl   *string      `pulumi:"secretUrl"`
	SourceVault *SubResource `pulumi:"sourceVault"`
}





type CloudServiceVaultAndSecretReferenceInput interface {
	pulumi.Input

	ToCloudServiceVaultAndSecretReferenceOutput() CloudServiceVaultAndSecretReferenceOutput
	ToCloudServiceVaultAndSecretReferenceOutputWithContext(context.Context) CloudServiceVaultAndSecretReferenceOutput
}

type CloudServiceVaultAndSecretReferenceArgs struct {
	SecretUrl   pulumi.StringPtrInput `pulumi:"secretUrl"`
	SourceVault SubResourcePtrInput   `pulumi:"sourceVault"`
}

func (CloudServiceVaultAndSecretReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceVaultAndSecretReference)(nil)).Elem()
}

func (i CloudServiceVaultAndSecretReferenceArgs) ToCloudServiceVaultAndSecretReferenceOutput() CloudServiceVaultAndSecretReferenceOutput {
	return i.ToCloudServiceVaultAndSecretReferenceOutputWithContext(context.Background())
}

func (i CloudServiceVaultAndSecretReferenceArgs) ToCloudServiceVaultAndSecretReferenceOutputWithContext(ctx context.Context) CloudServiceVaultAndSecretReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceVaultAndSecretReferenceOutput)
}

func (i CloudServiceVaultAndSecretReferenceArgs) ToCloudServiceVaultAndSecretReferencePtrOutput() CloudServiceVaultAndSecretReferencePtrOutput {
	return i.ToCloudServiceVaultAndSecretReferencePtrOutputWithContext(context.Background())
}

func (i CloudServiceVaultAndSecretReferenceArgs) ToCloudServiceVaultAndSecretReferencePtrOutputWithContext(ctx context.Context) CloudServiceVaultAndSecretReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceVaultAndSecretReferenceOutput).ToCloudServiceVaultAndSecretReferencePtrOutputWithContext(ctx)
}









type CloudServiceVaultAndSecretReferencePtrInput interface {
	pulumi.Input

	ToCloudServiceVaultAndSecretReferencePtrOutput() CloudServiceVaultAndSecretReferencePtrOutput
	ToCloudServiceVaultAndSecretReferencePtrOutputWithContext(context.Context) CloudServiceVaultAndSecretReferencePtrOutput
}

type cloudServiceVaultAndSecretReferencePtrType CloudServiceVaultAndSecretReferenceArgs

func CloudServiceVaultAndSecretReferencePtr(v *CloudServiceVaultAndSecretReferenceArgs) CloudServiceVaultAndSecretReferencePtrInput {
	return (*cloudServiceVaultAndSecretReferencePtrType)(v)
}

func (*cloudServiceVaultAndSecretReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceVaultAndSecretReference)(nil)).Elem()
}

func (i *cloudServiceVaultAndSecretReferencePtrType) ToCloudServiceVaultAndSecretReferencePtrOutput() CloudServiceVaultAndSecretReferencePtrOutput {
	return i.ToCloudServiceVaultAndSecretReferencePtrOutputWithContext(context.Background())
}

func (i *cloudServiceVaultAndSecretReferencePtrType) ToCloudServiceVaultAndSecretReferencePtrOutputWithContext(ctx context.Context) CloudServiceVaultAndSecretReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceVaultAndSecretReferencePtrOutput)
}

type CloudServiceVaultAndSecretReferenceOutput struct{ *pulumi.OutputState }

func (CloudServiceVaultAndSecretReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceVaultAndSecretReference)(nil)).Elem()
}

func (o CloudServiceVaultAndSecretReferenceOutput) ToCloudServiceVaultAndSecretReferenceOutput() CloudServiceVaultAndSecretReferenceOutput {
	return o
}

func (o CloudServiceVaultAndSecretReferenceOutput) ToCloudServiceVaultAndSecretReferenceOutputWithContext(ctx context.Context) CloudServiceVaultAndSecretReferenceOutput {
	return o
}

func (o CloudServiceVaultAndSecretReferenceOutput) ToCloudServiceVaultAndSecretReferencePtrOutput() CloudServiceVaultAndSecretReferencePtrOutput {
	return o.ToCloudServiceVaultAndSecretReferencePtrOutputWithContext(context.Background())
}

func (o CloudServiceVaultAndSecretReferenceOutput) ToCloudServiceVaultAndSecretReferencePtrOutputWithContext(ctx context.Context) CloudServiceVaultAndSecretReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudServiceVaultAndSecretReference) *CloudServiceVaultAndSecretReference {
		return &v
	}).(CloudServiceVaultAndSecretReferencePtrOutput)
}

func (o CloudServiceVaultAndSecretReferenceOutput) SecretUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceVaultAndSecretReference) *string { return v.SecretUrl }).(pulumi.StringPtrOutput)
}

func (o CloudServiceVaultAndSecretReferenceOutput) SourceVault() SubResourcePtrOutput {
	return o.ApplyT(func(v CloudServiceVaultAndSecretReference) *SubResource { return v.SourceVault }).(SubResourcePtrOutput)
}

type CloudServiceVaultAndSecretReferencePtrOutput struct{ *pulumi.OutputState }

func (CloudServiceVaultAndSecretReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceVaultAndSecretReference)(nil)).Elem()
}

func (o CloudServiceVaultAndSecretReferencePtrOutput) ToCloudServiceVaultAndSecretReferencePtrOutput() CloudServiceVaultAndSecretReferencePtrOutput {
	return o
}

func (o CloudServiceVaultAndSecretReferencePtrOutput) ToCloudServiceVaultAndSecretReferencePtrOutputWithContext(ctx context.Context) CloudServiceVaultAndSecretReferencePtrOutput {
	return o
}

func (o CloudServiceVaultAndSecretReferencePtrOutput) Elem() CloudServiceVaultAndSecretReferenceOutput {
	return o.ApplyT(func(v *CloudServiceVaultAndSecretReference) CloudServiceVaultAndSecretReference {
		if v != nil {
			return *v
		}
		var ret CloudServiceVaultAndSecretReference
		return ret
	}).(CloudServiceVaultAndSecretReferenceOutput)
}

func (o CloudServiceVaultAndSecretReferencePtrOutput) SecretUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceVaultAndSecretReference) *string {
		if v == nil {
			return nil
		}
		return v.SecretUrl
	}).(pulumi.StringPtrOutput)
}

func (o CloudServiceVaultAndSecretReferencePtrOutput) SourceVault() SubResourcePtrOutput {
	return o.ApplyT(func(v *CloudServiceVaultAndSecretReference) *SubResource {
		if v == nil {
			return nil
		}
		return v.SourceVault
	}).(SubResourcePtrOutput)
}

type CloudServiceVaultAndSecretReferenceResponse struct {
	SecretUrl   *string              `pulumi:"secretUrl"`
	SourceVault *SubResourceResponse `pulumi:"sourceVault"`
}

type CloudServiceVaultAndSecretReferenceResponseOutput struct{ *pulumi.OutputState }

func (CloudServiceVaultAndSecretReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceVaultAndSecretReferenceResponse)(nil)).Elem()
}

func (o CloudServiceVaultAndSecretReferenceResponseOutput) ToCloudServiceVaultAndSecretReferenceResponseOutput() CloudServiceVaultAndSecretReferenceResponseOutput {
	return o
}

func (o CloudServiceVaultAndSecretReferenceResponseOutput) ToCloudServiceVaultAndSecretReferenceResponseOutputWithContext(ctx context.Context) CloudServiceVaultAndSecretReferenceResponseOutput {
	return o
}

func (o CloudServiceVaultAndSecretReferenceResponseOutput) SecretUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceVaultAndSecretReferenceResponse) *string { return v.SecretUrl }).(pulumi.StringPtrOutput)
}

func (o CloudServiceVaultAndSecretReferenceResponseOutput) SourceVault() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v CloudServiceVaultAndSecretReferenceResponse) *SubResourceResponse { return v.SourceVault }).(SubResourceResponsePtrOutput)
}

type CloudServiceVaultAndSecretReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudServiceVaultAndSecretReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServiceVaultAndSecretReferenceResponse)(nil)).Elem()
}

func (o CloudServiceVaultAndSecretReferenceResponsePtrOutput) ToCloudServiceVaultAndSecretReferenceResponsePtrOutput() CloudServiceVaultAndSecretReferenceResponsePtrOutput {
	return o
}

func (o CloudServiceVaultAndSecretReferenceResponsePtrOutput) ToCloudServiceVaultAndSecretReferenceResponsePtrOutputWithContext(ctx context.Context) CloudServiceVaultAndSecretReferenceResponsePtrOutput {
	return o
}

func (o CloudServiceVaultAndSecretReferenceResponsePtrOutput) Elem() CloudServiceVaultAndSecretReferenceResponseOutput {
	return o.ApplyT(func(v *CloudServiceVaultAndSecretReferenceResponse) CloudServiceVaultAndSecretReferenceResponse {
		if v != nil {
			return *v
		}
		var ret CloudServiceVaultAndSecretReferenceResponse
		return ret
	}).(CloudServiceVaultAndSecretReferenceResponseOutput)
}

func (o CloudServiceVaultAndSecretReferenceResponsePtrOutput) SecretUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServiceVaultAndSecretReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.SecretUrl
	}).(pulumi.StringPtrOutput)
}

func (o CloudServiceVaultAndSecretReferenceResponsePtrOutput) SourceVault() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *CloudServiceVaultAndSecretReferenceResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return v.SourceVault
	}).(SubResourceResponsePtrOutput)
}

type CloudServiceVaultCertificate struct {
	CertificateUrl *string `pulumi:"certificateUrl"`
}





type CloudServiceVaultCertificateInput interface {
	pulumi.Input

	ToCloudServiceVaultCertificateOutput() CloudServiceVaultCertificateOutput
	ToCloudServiceVaultCertificateOutputWithContext(context.Context) CloudServiceVaultCertificateOutput
}

type CloudServiceVaultCertificateArgs struct {
	CertificateUrl pulumi.StringPtrInput `pulumi:"certificateUrl"`
}

func (CloudServiceVaultCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceVaultCertificate)(nil)).Elem()
}

func (i CloudServiceVaultCertificateArgs) ToCloudServiceVaultCertificateOutput() CloudServiceVaultCertificateOutput {
	return i.ToCloudServiceVaultCertificateOutputWithContext(context.Background())
}

func (i CloudServiceVaultCertificateArgs) ToCloudServiceVaultCertificateOutputWithContext(ctx context.Context) CloudServiceVaultCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceVaultCertificateOutput)
}





type CloudServiceVaultCertificateArrayInput interface {
	pulumi.Input

	ToCloudServiceVaultCertificateArrayOutput() CloudServiceVaultCertificateArrayOutput
	ToCloudServiceVaultCertificateArrayOutputWithContext(context.Context) CloudServiceVaultCertificateArrayOutput
}

type CloudServiceVaultCertificateArray []CloudServiceVaultCertificateInput

func (CloudServiceVaultCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudServiceVaultCertificate)(nil)).Elem()
}

func (i CloudServiceVaultCertificateArray) ToCloudServiceVaultCertificateArrayOutput() CloudServiceVaultCertificateArrayOutput {
	return i.ToCloudServiceVaultCertificateArrayOutputWithContext(context.Background())
}

func (i CloudServiceVaultCertificateArray) ToCloudServiceVaultCertificateArrayOutputWithContext(ctx context.Context) CloudServiceVaultCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceVaultCertificateArrayOutput)
}

type CloudServiceVaultCertificateOutput struct{ *pulumi.OutputState }

func (CloudServiceVaultCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceVaultCertificate)(nil)).Elem()
}

func (o CloudServiceVaultCertificateOutput) ToCloudServiceVaultCertificateOutput() CloudServiceVaultCertificateOutput {
	return o
}

func (o CloudServiceVaultCertificateOutput) ToCloudServiceVaultCertificateOutputWithContext(ctx context.Context) CloudServiceVaultCertificateOutput {
	return o
}

func (o CloudServiceVaultCertificateOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceVaultCertificate) *string { return v.CertificateUrl }).(pulumi.StringPtrOutput)
}

type CloudServiceVaultCertificateArrayOutput struct{ *pulumi.OutputState }

func (CloudServiceVaultCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudServiceVaultCertificate)(nil)).Elem()
}

func (o CloudServiceVaultCertificateArrayOutput) ToCloudServiceVaultCertificateArrayOutput() CloudServiceVaultCertificateArrayOutput {
	return o
}

func (o CloudServiceVaultCertificateArrayOutput) ToCloudServiceVaultCertificateArrayOutputWithContext(ctx context.Context) CloudServiceVaultCertificateArrayOutput {
	return o
}

func (o CloudServiceVaultCertificateArrayOutput) Index(i pulumi.IntInput) CloudServiceVaultCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CloudServiceVaultCertificate {
		return vs[0].([]CloudServiceVaultCertificate)[vs[1].(int)]
	}).(CloudServiceVaultCertificateOutput)
}

type CloudServiceVaultCertificateResponse struct {
	CertificateUrl *string `pulumi:"certificateUrl"`
}

type CloudServiceVaultCertificateResponseOutput struct{ *pulumi.OutputState }

func (CloudServiceVaultCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceVaultCertificateResponse)(nil)).Elem()
}

func (o CloudServiceVaultCertificateResponseOutput) ToCloudServiceVaultCertificateResponseOutput() CloudServiceVaultCertificateResponseOutput {
	return o
}

func (o CloudServiceVaultCertificateResponseOutput) ToCloudServiceVaultCertificateResponseOutputWithContext(ctx context.Context) CloudServiceVaultCertificateResponseOutput {
	return o
}

func (o CloudServiceVaultCertificateResponseOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServiceVaultCertificateResponse) *string { return v.CertificateUrl }).(pulumi.StringPtrOutput)
}

type CloudServiceVaultCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (CloudServiceVaultCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudServiceVaultCertificateResponse)(nil)).Elem()
}

func (o CloudServiceVaultCertificateResponseArrayOutput) ToCloudServiceVaultCertificateResponseArrayOutput() CloudServiceVaultCertificateResponseArrayOutput {
	return o
}

func (o CloudServiceVaultCertificateResponseArrayOutput) ToCloudServiceVaultCertificateResponseArrayOutputWithContext(ctx context.Context) CloudServiceVaultCertificateResponseArrayOutput {
	return o
}

func (o CloudServiceVaultCertificateResponseArrayOutput) Index(i pulumi.IntInput) CloudServiceVaultCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CloudServiceVaultCertificateResponse {
		return vs[0].([]CloudServiceVaultCertificateResponse)[vs[1].(int)]
	}).(CloudServiceVaultCertificateResponseOutput)
}

type CloudServiceVaultSecretGroup struct {
	SourceVault       *SubResource                   `pulumi:"sourceVault"`
	VaultCertificates []CloudServiceVaultCertificate `pulumi:"vaultCertificates"`
}





type CloudServiceVaultSecretGroupInput interface {
	pulumi.Input

	ToCloudServiceVaultSecretGroupOutput() CloudServiceVaultSecretGroupOutput
	ToCloudServiceVaultSecretGroupOutputWithContext(context.Context) CloudServiceVaultSecretGroupOutput
}

type CloudServiceVaultSecretGroupArgs struct {
	SourceVault       SubResourcePtrInput                    `pulumi:"sourceVault"`
	VaultCertificates CloudServiceVaultCertificateArrayInput `pulumi:"vaultCertificates"`
}

func (CloudServiceVaultSecretGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceVaultSecretGroup)(nil)).Elem()
}

func (i CloudServiceVaultSecretGroupArgs) ToCloudServiceVaultSecretGroupOutput() CloudServiceVaultSecretGroupOutput {
	return i.ToCloudServiceVaultSecretGroupOutputWithContext(context.Background())
}

func (i CloudServiceVaultSecretGroupArgs) ToCloudServiceVaultSecretGroupOutputWithContext(ctx context.Context) CloudServiceVaultSecretGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceVaultSecretGroupOutput)
}





type CloudServiceVaultSecretGroupArrayInput interface {
	pulumi.Input

	ToCloudServiceVaultSecretGroupArrayOutput() CloudServiceVaultSecretGroupArrayOutput
	ToCloudServiceVaultSecretGroupArrayOutputWithContext(context.Context) CloudServiceVaultSecretGroupArrayOutput
}

type CloudServiceVaultSecretGroupArray []CloudServiceVaultSecretGroupInput

func (CloudServiceVaultSecretGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudServiceVaultSecretGroup)(nil)).Elem()
}

func (i CloudServiceVaultSecretGroupArray) ToCloudServiceVaultSecretGroupArrayOutput() CloudServiceVaultSecretGroupArrayOutput {
	return i.ToCloudServiceVaultSecretGroupArrayOutputWithContext(context.Background())
}

func (i CloudServiceVaultSecretGroupArray) ToCloudServiceVaultSecretGroupArrayOutputWithContext(ctx context.Context) CloudServiceVaultSecretGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServiceVaultSecretGroupArrayOutput)
}

type CloudServiceVaultSecretGroupOutput struct{ *pulumi.OutputState }

func (CloudServiceVaultSecretGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceVaultSecretGroup)(nil)).Elem()
}

func (o CloudServiceVaultSecretGroupOutput) ToCloudServiceVaultSecretGroupOutput() CloudServiceVaultSecretGroupOutput {
	return o
}

func (o CloudServiceVaultSecretGroupOutput) ToCloudServiceVaultSecretGroupOutputWithContext(ctx context.Context) CloudServiceVaultSecretGroupOutput {
	return o
}

func (o CloudServiceVaultSecretGroupOutput) SourceVault() SubResourcePtrOutput {
	return o.ApplyT(func(v CloudServiceVaultSecretGroup) *SubResource { return v.SourceVault }).(SubResourcePtrOutput)
}

func (o CloudServiceVaultSecretGroupOutput) VaultCertificates() CloudServiceVaultCertificateArrayOutput {
	return o.ApplyT(func(v CloudServiceVaultSecretGroup) []CloudServiceVaultCertificate { return v.VaultCertificates }).(CloudServiceVaultCertificateArrayOutput)
}

type CloudServiceVaultSecretGroupArrayOutput struct{ *pulumi.OutputState }

func (CloudServiceVaultSecretGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudServiceVaultSecretGroup)(nil)).Elem()
}

func (o CloudServiceVaultSecretGroupArrayOutput) ToCloudServiceVaultSecretGroupArrayOutput() CloudServiceVaultSecretGroupArrayOutput {
	return o
}

func (o CloudServiceVaultSecretGroupArrayOutput) ToCloudServiceVaultSecretGroupArrayOutputWithContext(ctx context.Context) CloudServiceVaultSecretGroupArrayOutput {
	return o
}

func (o CloudServiceVaultSecretGroupArrayOutput) Index(i pulumi.IntInput) CloudServiceVaultSecretGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CloudServiceVaultSecretGroup {
		return vs[0].([]CloudServiceVaultSecretGroup)[vs[1].(int)]
	}).(CloudServiceVaultSecretGroupOutput)
}

type CloudServiceVaultSecretGroupResponse struct {
	SourceVault       *SubResourceResponse                   `pulumi:"sourceVault"`
	VaultCertificates []CloudServiceVaultCertificateResponse `pulumi:"vaultCertificates"`
}

type CloudServiceVaultSecretGroupResponseOutput struct{ *pulumi.OutputState }

func (CloudServiceVaultSecretGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServiceVaultSecretGroupResponse)(nil)).Elem()
}

func (o CloudServiceVaultSecretGroupResponseOutput) ToCloudServiceVaultSecretGroupResponseOutput() CloudServiceVaultSecretGroupResponseOutput {
	return o
}

func (o CloudServiceVaultSecretGroupResponseOutput) ToCloudServiceVaultSecretGroupResponseOutputWithContext(ctx context.Context) CloudServiceVaultSecretGroupResponseOutput {
	return o
}

func (o CloudServiceVaultSecretGroupResponseOutput) SourceVault() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v CloudServiceVaultSecretGroupResponse) *SubResourceResponse { return v.SourceVault }).(SubResourceResponsePtrOutput)
}

func (o CloudServiceVaultSecretGroupResponseOutput) VaultCertificates() CloudServiceVaultCertificateResponseArrayOutput {
	return o.ApplyT(func(v CloudServiceVaultSecretGroupResponse) []CloudServiceVaultCertificateResponse {
		return v.VaultCertificates
	}).(CloudServiceVaultCertificateResponseArrayOutput)
}

type CloudServiceVaultSecretGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (CloudServiceVaultSecretGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudServiceVaultSecretGroupResponse)(nil)).Elem()
}

func (o CloudServiceVaultSecretGroupResponseArrayOutput) ToCloudServiceVaultSecretGroupResponseArrayOutput() CloudServiceVaultSecretGroupResponseArrayOutput {
	return o
}

func (o CloudServiceVaultSecretGroupResponseArrayOutput) ToCloudServiceVaultSecretGroupResponseArrayOutputWithContext(ctx context.Context) CloudServiceVaultSecretGroupResponseArrayOutput {
	return o
}

func (o CloudServiceVaultSecretGroupResponseArrayOutput) Index(i pulumi.IntInput) CloudServiceVaultSecretGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CloudServiceVaultSecretGroupResponse {
		return vs[0].([]CloudServiceVaultSecretGroupResponse)[vs[1].(int)]
	}).(CloudServiceVaultSecretGroupResponseOutput)
}

type Extension struct {
	Name       *string                          `pulumi:"name"`
	Properties *CloudServiceExtensionProperties `pulumi:"properties"`
}





type ExtensionInput interface {
	pulumi.Input

	ToExtensionOutput() ExtensionOutput
	ToExtensionOutputWithContext(context.Context) ExtensionOutput
}

type ExtensionArgs struct {
	Name       pulumi.StringPtrInput                   `pulumi:"name"`
	Properties CloudServiceExtensionPropertiesPtrInput `pulumi:"properties"`
}

func (ExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Extension)(nil)).Elem()
}

func (i ExtensionArgs) ToExtensionOutput() ExtensionOutput {
	return i.ToExtensionOutputWithContext(context.Background())
}

func (i ExtensionArgs) ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionOutput)
}





type ExtensionArrayInput interface {
	pulumi.Input

	ToExtensionArrayOutput() ExtensionArrayOutput
	ToExtensionArrayOutputWithContext(context.Context) ExtensionArrayOutput
}

type ExtensionArray []ExtensionInput

func (ExtensionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Extension)(nil)).Elem()
}

func (i ExtensionArray) ToExtensionArrayOutput() ExtensionArrayOutput {
	return i.ToExtensionArrayOutputWithContext(context.Background())
}

func (i ExtensionArray) ToExtensionArrayOutputWithContext(ctx context.Context) ExtensionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionArrayOutput)
}

type ExtensionOutput struct{ *pulumi.OutputState }

func (ExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Extension)(nil)).Elem()
}

func (o ExtensionOutput) ToExtensionOutput() ExtensionOutput {
	return o
}

func (o ExtensionOutput) ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput {
	return o
}

func (o ExtensionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Extension) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtensionOutput) Properties() CloudServiceExtensionPropertiesPtrOutput {
	return o.ApplyT(func(v Extension) *CloudServiceExtensionProperties { return v.Properties }).(CloudServiceExtensionPropertiesPtrOutput)
}

type ExtensionArrayOutput struct{ *pulumi.OutputState }

func (ExtensionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Extension)(nil)).Elem()
}

func (o ExtensionArrayOutput) ToExtensionArrayOutput() ExtensionArrayOutput {
	return o
}

func (o ExtensionArrayOutput) ToExtensionArrayOutputWithContext(ctx context.Context) ExtensionArrayOutput {
	return o
}

func (o ExtensionArrayOutput) Index(i pulumi.IntInput) ExtensionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Extension {
		return vs[0].([]Extension)[vs[1].(int)]
	}).(ExtensionOutput)
}

type ExtensionResponse struct {
	Name       *string                                  `pulumi:"name"`
	Properties *CloudServiceExtensionPropertiesResponse `pulumi:"properties"`
}

type ExtensionResponseOutput struct{ *pulumi.OutputState }

func (ExtensionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionResponse)(nil)).Elem()
}

func (o ExtensionResponseOutput) ToExtensionResponseOutput() ExtensionResponseOutput {
	return o
}

func (o ExtensionResponseOutput) ToExtensionResponseOutputWithContext(ctx context.Context) ExtensionResponseOutput {
	return o
}

func (o ExtensionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtensionResponseOutput) Properties() CloudServiceExtensionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ExtensionResponse) *CloudServiceExtensionPropertiesResponse { return v.Properties }).(CloudServiceExtensionPropertiesResponsePtrOutput)
}

type ExtensionResponseArrayOutput struct{ *pulumi.OutputState }

func (ExtensionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtensionResponse)(nil)).Elem()
}

func (o ExtensionResponseArrayOutput) ToExtensionResponseArrayOutput() ExtensionResponseArrayOutput {
	return o
}

func (o ExtensionResponseArrayOutput) ToExtensionResponseArrayOutputWithContext(ctx context.Context) ExtensionResponseArrayOutput {
	return o
}

func (o ExtensionResponseArrayOutput) Index(i pulumi.IntInput) ExtensionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExtensionResponse {
		return vs[0].([]ExtensionResponse)[vs[1].(int)]
	}).(ExtensionResponseOutput)
}

type LoadBalancerConfiguration struct {
	Id         *string                             `pulumi:"id"`
	Name       string                              `pulumi:"name"`
	Properties LoadBalancerConfigurationProperties `pulumi:"properties"`
}





type LoadBalancerConfigurationInput interface {
	pulumi.Input

	ToLoadBalancerConfigurationOutput() LoadBalancerConfigurationOutput
	ToLoadBalancerConfigurationOutputWithContext(context.Context) LoadBalancerConfigurationOutput
}

type LoadBalancerConfigurationArgs struct {
	Id         pulumi.StringPtrInput                    `pulumi:"id"`
	Name       pulumi.StringInput                       `pulumi:"name"`
	Properties LoadBalancerConfigurationPropertiesInput `pulumi:"properties"`
}

func (LoadBalancerConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerConfiguration)(nil)).Elem()
}

func (i LoadBalancerConfigurationArgs) ToLoadBalancerConfigurationOutput() LoadBalancerConfigurationOutput {
	return i.ToLoadBalancerConfigurationOutputWithContext(context.Background())
}

func (i LoadBalancerConfigurationArgs) ToLoadBalancerConfigurationOutputWithContext(ctx context.Context) LoadBalancerConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerConfigurationOutput)
}





type LoadBalancerConfigurationArrayInput interface {
	pulumi.Input

	ToLoadBalancerConfigurationArrayOutput() LoadBalancerConfigurationArrayOutput
	ToLoadBalancerConfigurationArrayOutputWithContext(context.Context) LoadBalancerConfigurationArrayOutput
}

type LoadBalancerConfigurationArray []LoadBalancerConfigurationInput

func (LoadBalancerConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerConfiguration)(nil)).Elem()
}

func (i LoadBalancerConfigurationArray) ToLoadBalancerConfigurationArrayOutput() LoadBalancerConfigurationArrayOutput {
	return i.ToLoadBalancerConfigurationArrayOutputWithContext(context.Background())
}

func (i LoadBalancerConfigurationArray) ToLoadBalancerConfigurationArrayOutputWithContext(ctx context.Context) LoadBalancerConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerConfigurationArrayOutput)
}

type LoadBalancerConfigurationOutput struct{ *pulumi.OutputState }

func (LoadBalancerConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerConfiguration)(nil)).Elem()
}

func (o LoadBalancerConfigurationOutput) ToLoadBalancerConfigurationOutput() LoadBalancerConfigurationOutput {
	return o
}

func (o LoadBalancerConfigurationOutput) ToLoadBalancerConfigurationOutputWithContext(ctx context.Context) LoadBalancerConfigurationOutput {
	return o
}

func (o LoadBalancerConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerConfiguration) string { return v.Name }).(pulumi.StringOutput)
}

func (o LoadBalancerConfigurationOutput) Properties() LoadBalancerConfigurationPropertiesOutput {
	return o.ApplyT(func(v LoadBalancerConfiguration) LoadBalancerConfigurationProperties { return v.Properties }).(LoadBalancerConfigurationPropertiesOutput)
}

type LoadBalancerConfigurationArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerConfiguration)(nil)).Elem()
}

func (o LoadBalancerConfigurationArrayOutput) ToLoadBalancerConfigurationArrayOutput() LoadBalancerConfigurationArrayOutput {
	return o
}

func (o LoadBalancerConfigurationArrayOutput) ToLoadBalancerConfigurationArrayOutputWithContext(ctx context.Context) LoadBalancerConfigurationArrayOutput {
	return o
}

func (o LoadBalancerConfigurationArrayOutput) Index(i pulumi.IntInput) LoadBalancerConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancerConfiguration {
		return vs[0].([]LoadBalancerConfiguration)[vs[1].(int)]
	}).(LoadBalancerConfigurationOutput)
}

type LoadBalancerConfigurationProperties struct {
	FrontendIPConfigurations []LoadBalancerFrontendIPConfiguration `pulumi:"frontendIPConfigurations"`
}





type LoadBalancerConfigurationPropertiesInput interface {
	pulumi.Input

	ToLoadBalancerConfigurationPropertiesOutput() LoadBalancerConfigurationPropertiesOutput
	ToLoadBalancerConfigurationPropertiesOutputWithContext(context.Context) LoadBalancerConfigurationPropertiesOutput
}

type LoadBalancerConfigurationPropertiesArgs struct {
	FrontendIPConfigurations LoadBalancerFrontendIPConfigurationArrayInput `pulumi:"frontendIPConfigurations"`
}

func (LoadBalancerConfigurationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerConfigurationProperties)(nil)).Elem()
}

func (i LoadBalancerConfigurationPropertiesArgs) ToLoadBalancerConfigurationPropertiesOutput() LoadBalancerConfigurationPropertiesOutput {
	return i.ToLoadBalancerConfigurationPropertiesOutputWithContext(context.Background())
}

func (i LoadBalancerConfigurationPropertiesArgs) ToLoadBalancerConfigurationPropertiesOutputWithContext(ctx context.Context) LoadBalancerConfigurationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerConfigurationPropertiesOutput)
}

type LoadBalancerConfigurationPropertiesOutput struct{ *pulumi.OutputState }

func (LoadBalancerConfigurationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerConfigurationProperties)(nil)).Elem()
}

func (o LoadBalancerConfigurationPropertiesOutput) ToLoadBalancerConfigurationPropertiesOutput() LoadBalancerConfigurationPropertiesOutput {
	return o
}

func (o LoadBalancerConfigurationPropertiesOutput) ToLoadBalancerConfigurationPropertiesOutputWithContext(ctx context.Context) LoadBalancerConfigurationPropertiesOutput {
	return o
}

func (o LoadBalancerConfigurationPropertiesOutput) FrontendIPConfigurations() LoadBalancerFrontendIPConfigurationArrayOutput {
	return o.ApplyT(func(v LoadBalancerConfigurationProperties) []LoadBalancerFrontendIPConfiguration {
		return v.FrontendIPConfigurations
	}).(LoadBalancerFrontendIPConfigurationArrayOutput)
}

type LoadBalancerConfigurationPropertiesResponse struct {
	FrontendIPConfigurations []LoadBalancerFrontendIPConfigurationResponse `pulumi:"frontendIPConfigurations"`
}

type LoadBalancerConfigurationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancerConfigurationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerConfigurationPropertiesResponse)(nil)).Elem()
}

func (o LoadBalancerConfigurationPropertiesResponseOutput) ToLoadBalancerConfigurationPropertiesResponseOutput() LoadBalancerConfigurationPropertiesResponseOutput {
	return o
}

func (o LoadBalancerConfigurationPropertiesResponseOutput) ToLoadBalancerConfigurationPropertiesResponseOutputWithContext(ctx context.Context) LoadBalancerConfigurationPropertiesResponseOutput {
	return o
}

func (o LoadBalancerConfigurationPropertiesResponseOutput) FrontendIPConfigurations() LoadBalancerFrontendIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LoadBalancerConfigurationPropertiesResponse) []LoadBalancerFrontendIPConfigurationResponse {
		return v.FrontendIPConfigurations
	}).(LoadBalancerFrontendIPConfigurationResponseArrayOutput)
}

type LoadBalancerConfigurationResponse struct {
	Id         *string                                     `pulumi:"id"`
	Name       string                                      `pulumi:"name"`
	Properties LoadBalancerConfigurationPropertiesResponse `pulumi:"properties"`
}

type LoadBalancerConfigurationResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancerConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerConfigurationResponse)(nil)).Elem()
}

func (o LoadBalancerConfigurationResponseOutput) ToLoadBalancerConfigurationResponseOutput() LoadBalancerConfigurationResponseOutput {
	return o
}

func (o LoadBalancerConfigurationResponseOutput) ToLoadBalancerConfigurationResponseOutputWithContext(ctx context.Context) LoadBalancerConfigurationResponseOutput {
	return o
}

func (o LoadBalancerConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o LoadBalancerConfigurationResponseOutput) Properties() LoadBalancerConfigurationPropertiesResponseOutput {
	return o.ApplyT(func(v LoadBalancerConfigurationResponse) LoadBalancerConfigurationPropertiesResponse {
		return v.Properties
	}).(LoadBalancerConfigurationPropertiesResponseOutput)
}

type LoadBalancerConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerConfigurationResponse)(nil)).Elem()
}

func (o LoadBalancerConfigurationResponseArrayOutput) ToLoadBalancerConfigurationResponseArrayOutput() LoadBalancerConfigurationResponseArrayOutput {
	return o
}

func (o LoadBalancerConfigurationResponseArrayOutput) ToLoadBalancerConfigurationResponseArrayOutputWithContext(ctx context.Context) LoadBalancerConfigurationResponseArrayOutput {
	return o
}

func (o LoadBalancerConfigurationResponseArrayOutput) Index(i pulumi.IntInput) LoadBalancerConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancerConfigurationResponse {
		return vs[0].([]LoadBalancerConfigurationResponse)[vs[1].(int)]
	}).(LoadBalancerConfigurationResponseOutput)
}

type LoadBalancerFrontendIPConfiguration struct {
	Name       string                                        `pulumi:"name"`
	Properties LoadBalancerFrontendIPConfigurationProperties `pulumi:"properties"`
}





type LoadBalancerFrontendIPConfigurationInput interface {
	pulumi.Input

	ToLoadBalancerFrontendIPConfigurationOutput() LoadBalancerFrontendIPConfigurationOutput
	ToLoadBalancerFrontendIPConfigurationOutputWithContext(context.Context) LoadBalancerFrontendIPConfigurationOutput
}

type LoadBalancerFrontendIPConfigurationArgs struct {
	Name       pulumi.StringInput                                 `pulumi:"name"`
	Properties LoadBalancerFrontendIPConfigurationPropertiesInput `pulumi:"properties"`
}

func (LoadBalancerFrontendIPConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerFrontendIPConfiguration)(nil)).Elem()
}

func (i LoadBalancerFrontendIPConfigurationArgs) ToLoadBalancerFrontendIPConfigurationOutput() LoadBalancerFrontendIPConfigurationOutput {
	return i.ToLoadBalancerFrontendIPConfigurationOutputWithContext(context.Background())
}

func (i LoadBalancerFrontendIPConfigurationArgs) ToLoadBalancerFrontendIPConfigurationOutputWithContext(ctx context.Context) LoadBalancerFrontendIPConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerFrontendIPConfigurationOutput)
}





type LoadBalancerFrontendIPConfigurationArrayInput interface {
	pulumi.Input

	ToLoadBalancerFrontendIPConfigurationArrayOutput() LoadBalancerFrontendIPConfigurationArrayOutput
	ToLoadBalancerFrontendIPConfigurationArrayOutputWithContext(context.Context) LoadBalancerFrontendIPConfigurationArrayOutput
}

type LoadBalancerFrontendIPConfigurationArray []LoadBalancerFrontendIPConfigurationInput

func (LoadBalancerFrontendIPConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerFrontendIPConfiguration)(nil)).Elem()
}

func (i LoadBalancerFrontendIPConfigurationArray) ToLoadBalancerFrontendIPConfigurationArrayOutput() LoadBalancerFrontendIPConfigurationArrayOutput {
	return i.ToLoadBalancerFrontendIPConfigurationArrayOutputWithContext(context.Background())
}

func (i LoadBalancerFrontendIPConfigurationArray) ToLoadBalancerFrontendIPConfigurationArrayOutputWithContext(ctx context.Context) LoadBalancerFrontendIPConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerFrontendIPConfigurationArrayOutput)
}

type LoadBalancerFrontendIPConfigurationOutput struct{ *pulumi.OutputState }

func (LoadBalancerFrontendIPConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerFrontendIPConfiguration)(nil)).Elem()
}

func (o LoadBalancerFrontendIPConfigurationOutput) ToLoadBalancerFrontendIPConfigurationOutput() LoadBalancerFrontendIPConfigurationOutput {
	return o
}

func (o LoadBalancerFrontendIPConfigurationOutput) ToLoadBalancerFrontendIPConfigurationOutputWithContext(ctx context.Context) LoadBalancerFrontendIPConfigurationOutput {
	return o
}

func (o LoadBalancerFrontendIPConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerFrontendIPConfiguration) string { return v.Name }).(pulumi.StringOutput)
}

func (o LoadBalancerFrontendIPConfigurationOutput) Properties() LoadBalancerFrontendIPConfigurationPropertiesOutput {
	return o.ApplyT(func(v LoadBalancerFrontendIPConfiguration) LoadBalancerFrontendIPConfigurationProperties {
		return v.Properties
	}).(LoadBalancerFrontendIPConfigurationPropertiesOutput)
}

type LoadBalancerFrontendIPConfigurationArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerFrontendIPConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerFrontendIPConfiguration)(nil)).Elem()
}

func (o LoadBalancerFrontendIPConfigurationArrayOutput) ToLoadBalancerFrontendIPConfigurationArrayOutput() LoadBalancerFrontendIPConfigurationArrayOutput {
	return o
}

func (o LoadBalancerFrontendIPConfigurationArrayOutput) ToLoadBalancerFrontendIPConfigurationArrayOutputWithContext(ctx context.Context) LoadBalancerFrontendIPConfigurationArrayOutput {
	return o
}

func (o LoadBalancerFrontendIPConfigurationArrayOutput) Index(i pulumi.IntInput) LoadBalancerFrontendIPConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancerFrontendIPConfiguration {
		return vs[0].([]LoadBalancerFrontendIPConfiguration)[vs[1].(int)]
	}).(LoadBalancerFrontendIPConfigurationOutput)
}

type LoadBalancerFrontendIPConfigurationProperties struct {
	PrivateIPAddress *string      `pulumi:"privateIPAddress"`
	PublicIPAddress  *SubResource `pulumi:"publicIPAddress"`
	Subnet           *SubResource `pulumi:"subnet"`
}





type LoadBalancerFrontendIPConfigurationPropertiesInput interface {
	pulumi.Input

	ToLoadBalancerFrontendIPConfigurationPropertiesOutput() LoadBalancerFrontendIPConfigurationPropertiesOutput
	ToLoadBalancerFrontendIPConfigurationPropertiesOutputWithContext(context.Context) LoadBalancerFrontendIPConfigurationPropertiesOutput
}

type LoadBalancerFrontendIPConfigurationPropertiesArgs struct {
	PrivateIPAddress pulumi.StringPtrInput `pulumi:"privateIPAddress"`
	PublicIPAddress  SubResourcePtrInput   `pulumi:"publicIPAddress"`
	Subnet           SubResourcePtrInput   `pulumi:"subnet"`
}

func (LoadBalancerFrontendIPConfigurationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerFrontendIPConfigurationProperties)(nil)).Elem()
}

func (i LoadBalancerFrontendIPConfigurationPropertiesArgs) ToLoadBalancerFrontendIPConfigurationPropertiesOutput() LoadBalancerFrontendIPConfigurationPropertiesOutput {
	return i.ToLoadBalancerFrontendIPConfigurationPropertiesOutputWithContext(context.Background())
}

func (i LoadBalancerFrontendIPConfigurationPropertiesArgs) ToLoadBalancerFrontendIPConfigurationPropertiesOutputWithContext(ctx context.Context) LoadBalancerFrontendIPConfigurationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerFrontendIPConfigurationPropertiesOutput)
}

type LoadBalancerFrontendIPConfigurationPropertiesOutput struct{ *pulumi.OutputState }

func (LoadBalancerFrontendIPConfigurationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerFrontendIPConfigurationProperties)(nil)).Elem()
}

func (o LoadBalancerFrontendIPConfigurationPropertiesOutput) ToLoadBalancerFrontendIPConfigurationPropertiesOutput() LoadBalancerFrontendIPConfigurationPropertiesOutput {
	return o
}

func (o LoadBalancerFrontendIPConfigurationPropertiesOutput) ToLoadBalancerFrontendIPConfigurationPropertiesOutputWithContext(ctx context.Context) LoadBalancerFrontendIPConfigurationPropertiesOutput {
	return o
}

func (o LoadBalancerFrontendIPConfigurationPropertiesOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerFrontendIPConfigurationProperties) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerFrontendIPConfigurationPropertiesOutput) PublicIPAddress() SubResourcePtrOutput {
	return o.ApplyT(func(v LoadBalancerFrontendIPConfigurationProperties) *SubResource { return v.PublicIPAddress }).(SubResourcePtrOutput)
}

func (o LoadBalancerFrontendIPConfigurationPropertiesOutput) Subnet() SubResourcePtrOutput {
	return o.ApplyT(func(v LoadBalancerFrontendIPConfigurationProperties) *SubResource { return v.Subnet }).(SubResourcePtrOutput)
}

type LoadBalancerFrontendIPConfigurationPropertiesResponse struct {
	PrivateIPAddress *string              `pulumi:"privateIPAddress"`
	PublicIPAddress  *SubResourceResponse `pulumi:"publicIPAddress"`
	Subnet           *SubResourceResponse `pulumi:"subnet"`
}

type LoadBalancerFrontendIPConfigurationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancerFrontendIPConfigurationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerFrontendIPConfigurationPropertiesResponse)(nil)).Elem()
}

func (o LoadBalancerFrontendIPConfigurationPropertiesResponseOutput) ToLoadBalancerFrontendIPConfigurationPropertiesResponseOutput() LoadBalancerFrontendIPConfigurationPropertiesResponseOutput {
	return o
}

func (o LoadBalancerFrontendIPConfigurationPropertiesResponseOutput) ToLoadBalancerFrontendIPConfigurationPropertiesResponseOutputWithContext(ctx context.Context) LoadBalancerFrontendIPConfigurationPropertiesResponseOutput {
	return o
}

func (o LoadBalancerFrontendIPConfigurationPropertiesResponseOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerFrontendIPConfigurationPropertiesResponse) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerFrontendIPConfigurationPropertiesResponseOutput) PublicIPAddress() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LoadBalancerFrontendIPConfigurationPropertiesResponse) *SubResourceResponse {
		return v.PublicIPAddress
	}).(SubResourceResponsePtrOutput)
}

func (o LoadBalancerFrontendIPConfigurationPropertiesResponseOutput) Subnet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LoadBalancerFrontendIPConfigurationPropertiesResponse) *SubResourceResponse { return v.Subnet }).(SubResourceResponsePtrOutput)
}

type LoadBalancerFrontendIPConfigurationResponse struct {
	Name       string                                                `pulumi:"name"`
	Properties LoadBalancerFrontendIPConfigurationPropertiesResponse `pulumi:"properties"`
}

type LoadBalancerFrontendIPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancerFrontendIPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerFrontendIPConfigurationResponse)(nil)).Elem()
}

func (o LoadBalancerFrontendIPConfigurationResponseOutput) ToLoadBalancerFrontendIPConfigurationResponseOutput() LoadBalancerFrontendIPConfigurationResponseOutput {
	return o
}

func (o LoadBalancerFrontendIPConfigurationResponseOutput) ToLoadBalancerFrontendIPConfigurationResponseOutputWithContext(ctx context.Context) LoadBalancerFrontendIPConfigurationResponseOutput {
	return o
}

func (o LoadBalancerFrontendIPConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerFrontendIPConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o LoadBalancerFrontendIPConfigurationResponseOutput) Properties() LoadBalancerFrontendIPConfigurationPropertiesResponseOutput {
	return o.ApplyT(func(v LoadBalancerFrontendIPConfigurationResponse) LoadBalancerFrontendIPConfigurationPropertiesResponse {
		return v.Properties
	}).(LoadBalancerFrontendIPConfigurationPropertiesResponseOutput)
}

type LoadBalancerFrontendIPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerFrontendIPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerFrontendIPConfigurationResponse)(nil)).Elem()
}

func (o LoadBalancerFrontendIPConfigurationResponseArrayOutput) ToLoadBalancerFrontendIPConfigurationResponseArrayOutput() LoadBalancerFrontendIPConfigurationResponseArrayOutput {
	return o
}

func (o LoadBalancerFrontendIPConfigurationResponseArrayOutput) ToLoadBalancerFrontendIPConfigurationResponseArrayOutputWithContext(ctx context.Context) LoadBalancerFrontendIPConfigurationResponseArrayOutput {
	return o
}

func (o LoadBalancerFrontendIPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) LoadBalancerFrontendIPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancerFrontendIPConfigurationResponse {
		return vs[0].([]LoadBalancerFrontendIPConfigurationResponse)[vs[1].(int)]
	}).(LoadBalancerFrontendIPConfigurationResponseOutput)
}

type SubResource struct {
	Id *string `pulumi:"id"`
}





type SubResourceInput interface {
	pulumi.Input

	ToSubResourceOutput() SubResourceOutput
	ToSubResourceOutputWithContext(context.Context) SubResourceOutput
}

type SubResourceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (i SubResourceArgs) ToSubResourceOutput() SubResourceOutput {
	return i.ToSubResourceOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput)
}

func (i SubResourceArgs) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput).ToSubResourcePtrOutputWithContext(ctx)
}









type SubResourcePtrInput interface {
	pulumi.Input

	ToSubResourcePtrOutput() SubResourcePtrOutput
	ToSubResourcePtrOutputWithContext(context.Context) SubResourcePtrOutput
}

type subResourcePtrType SubResourceArgs

func SubResourcePtr(v *SubResourceArgs) SubResourcePtrInput {
	return (*subResourcePtrType)(v)
}

func (*subResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (i *subResourcePtrType) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i *subResourcePtrType) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourcePtrOutput)
}

type SubResourceOutput struct{ *pulumi.OutputState }

func (SubResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (o SubResourceOutput) ToSubResourceOutput() SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o.ToSubResourcePtrOutputWithContext(context.Background())
}

func (o SubResourceOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubResource) *SubResource {
		return &v
	}).(SubResourcePtrOutput)
}

func (o SubResourceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResource) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourcePtrOutput struct{ *pulumi.OutputState }

func (SubResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) Elem() SubResourceOutput {
	return o.ApplyT(func(v *SubResource) SubResource {
		if v != nil {
			return *v
		}
		var ret SubResource
		return ret
	}).(SubResourceOutput)
}

func (o SubResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResource) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceResponse struct {
	Id *string `pulumi:"id"`
}

type SubResourceResponseOutput struct{ *pulumi.OutputState }

func (SubResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResourceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourceResponsePtrOutput struct{ *pulumi.OutputState }

func (SubResourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) Elem() SubResourceResponseOutput {
	return o.ApplyT(func(v *SubResourceResponse) SubResourceResponse {
		if v != nil {
			return *v
		}
		var ret SubResourceResponse
		return ret
	}).(SubResourceResponseOutput)
}

func (o SubResourceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt      string `pulumi:"createdAt"`
	LastModifiedAt string `pulumi:"lastModifiedAt"`
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

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.LastModifiedAt }).(pulumi.StringOutput)
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

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CloudServiceExtensionProfileOutput{})
	pulumi.RegisterOutputType(CloudServiceExtensionProfilePtrOutput{})
	pulumi.RegisterOutputType(CloudServiceExtensionProfileResponseOutput{})
	pulumi.RegisterOutputType(CloudServiceExtensionProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(CloudServiceExtensionPropertiesOutput{})
	pulumi.RegisterOutputType(CloudServiceExtensionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CloudServiceExtensionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CloudServiceExtensionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CloudServiceNetworkProfileOutput{})
	pulumi.RegisterOutputType(CloudServiceNetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(CloudServiceNetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(CloudServiceNetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(CloudServiceOsProfileOutput{})
	pulumi.RegisterOutputType(CloudServiceOsProfilePtrOutput{})
	pulumi.RegisterOutputType(CloudServiceOsProfileResponseOutput{})
	pulumi.RegisterOutputType(CloudServiceOsProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(CloudServicePropertiesOutput{})
	pulumi.RegisterOutputType(CloudServicePropertiesPtrOutput{})
	pulumi.RegisterOutputType(CloudServicePropertiesResponseOutput{})
	pulumi.RegisterOutputType(CloudServiceRoleProfileOutput{})
	pulumi.RegisterOutputType(CloudServiceRoleProfilePtrOutput{})
	pulumi.RegisterOutputType(CloudServiceRoleProfilePropertiesOutput{})
	pulumi.RegisterOutputType(CloudServiceRoleProfilePropertiesArrayOutput{})
	pulumi.RegisterOutputType(CloudServiceRoleProfilePropertiesResponseOutput{})
	pulumi.RegisterOutputType(CloudServiceRoleProfilePropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(CloudServiceRoleProfileResponseOutput{})
	pulumi.RegisterOutputType(CloudServiceRoleProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(CloudServiceRoleSkuOutput{})
	pulumi.RegisterOutputType(CloudServiceRoleSkuPtrOutput{})
	pulumi.RegisterOutputType(CloudServiceRoleSkuResponseOutput{})
	pulumi.RegisterOutputType(CloudServiceRoleSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(CloudServiceVaultAndSecretReferenceOutput{})
	pulumi.RegisterOutputType(CloudServiceVaultAndSecretReferencePtrOutput{})
	pulumi.RegisterOutputType(CloudServiceVaultAndSecretReferenceResponseOutput{})
	pulumi.RegisterOutputType(CloudServiceVaultAndSecretReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(CloudServiceVaultCertificateOutput{})
	pulumi.RegisterOutputType(CloudServiceVaultCertificateArrayOutput{})
	pulumi.RegisterOutputType(CloudServiceVaultCertificateResponseOutput{})
	pulumi.RegisterOutputType(CloudServiceVaultCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(CloudServiceVaultSecretGroupOutput{})
	pulumi.RegisterOutputType(CloudServiceVaultSecretGroupArrayOutput{})
	pulumi.RegisterOutputType(CloudServiceVaultSecretGroupResponseOutput{})
	pulumi.RegisterOutputType(CloudServiceVaultSecretGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(ExtensionOutput{})
	pulumi.RegisterOutputType(ExtensionArrayOutput{})
	pulumi.RegisterOutputType(ExtensionResponseOutput{})
	pulumi.RegisterOutputType(ExtensionResponseArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerConfigurationOutput{})
	pulumi.RegisterOutputType(LoadBalancerConfigurationArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerConfigurationPropertiesOutput{})
	pulumi.RegisterOutputType(LoadBalancerConfigurationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancerConfigurationResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancerConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerFrontendIPConfigurationOutput{})
	pulumi.RegisterOutputType(LoadBalancerFrontendIPConfigurationArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerFrontendIPConfigurationPropertiesOutput{})
	pulumi.RegisterOutputType(LoadBalancerFrontendIPConfigurationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancerFrontendIPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancerFrontendIPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourcePtrOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SubResourceResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}

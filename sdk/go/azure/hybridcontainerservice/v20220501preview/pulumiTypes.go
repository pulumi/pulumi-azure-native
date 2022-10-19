


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AADProfile struct {
	AdminGroupObjectIDs []string `pulumi:"adminGroupObjectIDs"`
	ClientAppID         *string  `pulumi:"clientAppID"`
	EnableAzureRbac     *bool    `pulumi:"enableAzureRbac"`
	Managed             *bool    `pulumi:"managed"`
	ServerAppID         *string  `pulumi:"serverAppID"`
	ServerAppSecret     *string  `pulumi:"serverAppSecret"`
	TenantID            *string  `pulumi:"tenantID"`
}





type AADProfileInput interface {
	pulumi.Input

	ToAADProfileOutput() AADProfileOutput
	ToAADProfileOutputWithContext(context.Context) AADProfileOutput
}

type AADProfileArgs struct {
	AdminGroupObjectIDs pulumi.StringArrayInput `pulumi:"adminGroupObjectIDs"`
	ClientAppID         pulumi.StringPtrInput   `pulumi:"clientAppID"`
	EnableAzureRbac     pulumi.BoolPtrInput     `pulumi:"enableAzureRbac"`
	Managed             pulumi.BoolPtrInput     `pulumi:"managed"`
	ServerAppID         pulumi.StringPtrInput   `pulumi:"serverAppID"`
	ServerAppSecret     pulumi.StringPtrInput   `pulumi:"serverAppSecret"`
	TenantID            pulumi.StringPtrInput   `pulumi:"tenantID"`
}

func (AADProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AADProfile)(nil)).Elem()
}

func (i AADProfileArgs) ToAADProfileOutput() AADProfileOutput {
	return i.ToAADProfileOutputWithContext(context.Background())
}

func (i AADProfileArgs) ToAADProfileOutputWithContext(ctx context.Context) AADProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AADProfileOutput)
}

func (i AADProfileArgs) ToAADProfilePtrOutput() AADProfilePtrOutput {
	return i.ToAADProfilePtrOutputWithContext(context.Background())
}

func (i AADProfileArgs) ToAADProfilePtrOutputWithContext(ctx context.Context) AADProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AADProfileOutput).ToAADProfilePtrOutputWithContext(ctx)
}









type AADProfilePtrInput interface {
	pulumi.Input

	ToAADProfilePtrOutput() AADProfilePtrOutput
	ToAADProfilePtrOutputWithContext(context.Context) AADProfilePtrOutput
}

type aadprofilePtrType AADProfileArgs

func AADProfilePtr(v *AADProfileArgs) AADProfilePtrInput {
	return (*aadprofilePtrType)(v)
}

func (*aadprofilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AADProfile)(nil)).Elem()
}

func (i *aadprofilePtrType) ToAADProfilePtrOutput() AADProfilePtrOutput {
	return i.ToAADProfilePtrOutputWithContext(context.Background())
}

func (i *aadprofilePtrType) ToAADProfilePtrOutputWithContext(ctx context.Context) AADProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AADProfilePtrOutput)
}

type AADProfileOutput struct{ *pulumi.OutputState }

func (AADProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AADProfile)(nil)).Elem()
}

func (o AADProfileOutput) ToAADProfileOutput() AADProfileOutput {
	return o
}

func (o AADProfileOutput) ToAADProfileOutputWithContext(ctx context.Context) AADProfileOutput {
	return o
}

func (o AADProfileOutput) ToAADProfilePtrOutput() AADProfilePtrOutput {
	return o.ToAADProfilePtrOutputWithContext(context.Background())
}

func (o AADProfileOutput) ToAADProfilePtrOutputWithContext(ctx context.Context) AADProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AADProfile) *AADProfile {
		return &v
	}).(AADProfilePtrOutput)
}

func (o AADProfileOutput) AdminGroupObjectIDs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AADProfile) []string { return v.AdminGroupObjectIDs }).(pulumi.StringArrayOutput)
}

func (o AADProfileOutput) ClientAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AADProfile) *string { return v.ClientAppID }).(pulumi.StringPtrOutput)
}

func (o AADProfileOutput) EnableAzureRbac() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AADProfile) *bool { return v.EnableAzureRbac }).(pulumi.BoolPtrOutput)
}

func (o AADProfileOutput) Managed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AADProfile) *bool { return v.Managed }).(pulumi.BoolPtrOutput)
}

func (o AADProfileOutput) ServerAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AADProfile) *string { return v.ServerAppID }).(pulumi.StringPtrOutput)
}

func (o AADProfileOutput) ServerAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AADProfile) *string { return v.ServerAppSecret }).(pulumi.StringPtrOutput)
}

func (o AADProfileOutput) TenantID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AADProfile) *string { return v.TenantID }).(pulumi.StringPtrOutput)
}

type AADProfilePtrOutput struct{ *pulumi.OutputState }

func (AADProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AADProfile)(nil)).Elem()
}

func (o AADProfilePtrOutput) ToAADProfilePtrOutput() AADProfilePtrOutput {
	return o
}

func (o AADProfilePtrOutput) ToAADProfilePtrOutputWithContext(ctx context.Context) AADProfilePtrOutput {
	return o
}

func (o AADProfilePtrOutput) Elem() AADProfileOutput {
	return o.ApplyT(func(v *AADProfile) AADProfile {
		if v != nil {
			return *v
		}
		var ret AADProfile
		return ret
	}).(AADProfileOutput)
}

func (o AADProfilePtrOutput) AdminGroupObjectIDs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AADProfile) []string {
		if v == nil {
			return nil
		}
		return v.AdminGroupObjectIDs
	}).(pulumi.StringArrayOutput)
}

func (o AADProfilePtrOutput) ClientAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADProfile) *string {
		if v == nil {
			return nil
		}
		return v.ClientAppID
	}).(pulumi.StringPtrOutput)
}

func (o AADProfilePtrOutput) EnableAzureRbac() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AADProfile) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAzureRbac
	}).(pulumi.BoolPtrOutput)
}

func (o AADProfilePtrOutput) Managed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AADProfile) *bool {
		if v == nil {
			return nil
		}
		return v.Managed
	}).(pulumi.BoolPtrOutput)
}

func (o AADProfilePtrOutput) ServerAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADProfile) *string {
		if v == nil {
			return nil
		}
		return v.ServerAppID
	}).(pulumi.StringPtrOutput)
}

func (o AADProfilePtrOutput) ServerAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADProfile) *string {
		if v == nil {
			return nil
		}
		return v.ServerAppSecret
	}).(pulumi.StringPtrOutput)
}

func (o AADProfilePtrOutput) TenantID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADProfile) *string {
		if v == nil {
			return nil
		}
		return v.TenantID
	}).(pulumi.StringPtrOutput)
}

type AADProfileResponseResponse struct {
	AdminGroupObjectIDs []string `pulumi:"adminGroupObjectIDs"`
	ClientAppID         *string  `pulumi:"clientAppID"`
	EnableAzureRbac     *bool    `pulumi:"enableAzureRbac"`
	Managed             *bool    `pulumi:"managed"`
	ServerAppID         *string  `pulumi:"serverAppID"`
	TenantID            *string  `pulumi:"tenantID"`
}

type AADProfileResponseResponseOutput struct{ *pulumi.OutputState }

func (AADProfileResponseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AADProfileResponseResponse)(nil)).Elem()
}

func (o AADProfileResponseResponseOutput) ToAADProfileResponseResponseOutput() AADProfileResponseResponseOutput {
	return o
}

func (o AADProfileResponseResponseOutput) ToAADProfileResponseResponseOutputWithContext(ctx context.Context) AADProfileResponseResponseOutput {
	return o
}

func (o AADProfileResponseResponseOutput) AdminGroupObjectIDs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AADProfileResponseResponse) []string { return v.AdminGroupObjectIDs }).(pulumi.StringArrayOutput)
}

func (o AADProfileResponseResponseOutput) ClientAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AADProfileResponseResponse) *string { return v.ClientAppID }).(pulumi.StringPtrOutput)
}

func (o AADProfileResponseResponseOutput) EnableAzureRbac() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AADProfileResponseResponse) *bool { return v.EnableAzureRbac }).(pulumi.BoolPtrOutput)
}

func (o AADProfileResponseResponseOutput) Managed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AADProfileResponseResponse) *bool { return v.Managed }).(pulumi.BoolPtrOutput)
}

func (o AADProfileResponseResponseOutput) ServerAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AADProfileResponseResponse) *string { return v.ServerAppID }).(pulumi.StringPtrOutput)
}

func (o AADProfileResponseResponseOutput) TenantID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AADProfileResponseResponse) *string { return v.TenantID }).(pulumi.StringPtrOutput)
}

type AADProfileResponseResponsePtrOutput struct{ *pulumi.OutputState }

func (AADProfileResponseResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AADProfileResponseResponse)(nil)).Elem()
}

func (o AADProfileResponseResponsePtrOutput) ToAADProfileResponseResponsePtrOutput() AADProfileResponseResponsePtrOutput {
	return o
}

func (o AADProfileResponseResponsePtrOutput) ToAADProfileResponseResponsePtrOutputWithContext(ctx context.Context) AADProfileResponseResponsePtrOutput {
	return o
}

func (o AADProfileResponseResponsePtrOutput) Elem() AADProfileResponseResponseOutput {
	return o.ApplyT(func(v *AADProfileResponseResponse) AADProfileResponseResponse {
		if v != nil {
			return *v
		}
		var ret AADProfileResponseResponse
		return ret
	}).(AADProfileResponseResponseOutput)
}

func (o AADProfileResponseResponsePtrOutput) AdminGroupObjectIDs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AADProfileResponseResponse) []string {
		if v == nil {
			return nil
		}
		return v.AdminGroupObjectIDs
	}).(pulumi.StringArrayOutput)
}

func (o AADProfileResponseResponsePtrOutput) ClientAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADProfileResponseResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientAppID
	}).(pulumi.StringPtrOutput)
}

func (o AADProfileResponseResponsePtrOutput) EnableAzureRbac() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AADProfileResponseResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAzureRbac
	}).(pulumi.BoolPtrOutput)
}

func (o AADProfileResponseResponsePtrOutput) Managed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AADProfileResponseResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Managed
	}).(pulumi.BoolPtrOutput)
}

func (o AADProfileResponseResponsePtrOutput) ServerAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADProfileResponseResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServerAppID
	}).(pulumi.StringPtrOutput)
}

func (o AADProfileResponseResponsePtrOutput) TenantID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADProfileResponseResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantID
	}).(pulumi.StringPtrOutput)
}

type AddonProfiles struct {
	Config  map[string]string `pulumi:"config"`
	Enabled *bool             `pulumi:"enabled"`
}





type AddonProfilesInput interface {
	pulumi.Input

	ToAddonProfilesOutput() AddonProfilesOutput
	ToAddonProfilesOutputWithContext(context.Context) AddonProfilesOutput
}

type AddonProfilesArgs struct {
	Config  pulumi.StringMapInput `pulumi:"config"`
	Enabled pulumi.BoolPtrInput   `pulumi:"enabled"`
}

func (AddonProfilesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonProfiles)(nil)).Elem()
}

func (i AddonProfilesArgs) ToAddonProfilesOutput() AddonProfilesOutput {
	return i.ToAddonProfilesOutputWithContext(context.Background())
}

func (i AddonProfilesArgs) ToAddonProfilesOutputWithContext(ctx context.Context) AddonProfilesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddonProfilesOutput)
}





type AddonProfilesMapInput interface {
	pulumi.Input

	ToAddonProfilesMapOutput() AddonProfilesMapOutput
	ToAddonProfilesMapOutputWithContext(context.Context) AddonProfilesMapOutput
}

type AddonProfilesMap map[string]AddonProfilesInput

func (AddonProfilesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AddonProfiles)(nil)).Elem()
}

func (i AddonProfilesMap) ToAddonProfilesMapOutput() AddonProfilesMapOutput {
	return i.ToAddonProfilesMapOutputWithContext(context.Background())
}

func (i AddonProfilesMap) ToAddonProfilesMapOutputWithContext(ctx context.Context) AddonProfilesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddonProfilesMapOutput)
}

type AddonProfilesOutput struct{ *pulumi.OutputState }

func (AddonProfilesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonProfiles)(nil)).Elem()
}

func (o AddonProfilesOutput) ToAddonProfilesOutput() AddonProfilesOutput {
	return o
}

func (o AddonProfilesOutput) ToAddonProfilesOutputWithContext(ctx context.Context) AddonProfilesOutput {
	return o
}

func (o AddonProfilesOutput) Config() pulumi.StringMapOutput {
	return o.ApplyT(func(v AddonProfiles) map[string]string { return v.Config }).(pulumi.StringMapOutput)
}

func (o AddonProfilesOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AddonProfiles) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type AddonProfilesMapOutput struct{ *pulumi.OutputState }

func (AddonProfilesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AddonProfiles)(nil)).Elem()
}

func (o AddonProfilesMapOutput) ToAddonProfilesMapOutput() AddonProfilesMapOutput {
	return o
}

func (o AddonProfilesMapOutput) ToAddonProfilesMapOutputWithContext(ctx context.Context) AddonProfilesMapOutput {
	return o
}

func (o AddonProfilesMapOutput) MapIndex(k pulumi.StringInput) AddonProfilesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AddonProfiles {
		return vs[0].(map[string]AddonProfiles)[vs[1].(string)]
	}).(AddonProfilesOutput)
}

type AddonProfilesResponse struct {
	Config  map[string]string `pulumi:"config"`
	Enabled *bool             `pulumi:"enabled"`
}

type AddonProfilesResponseOutput struct{ *pulumi.OutputState }

func (AddonProfilesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonProfilesResponse)(nil)).Elem()
}

func (o AddonProfilesResponseOutput) ToAddonProfilesResponseOutput() AddonProfilesResponseOutput {
	return o
}

func (o AddonProfilesResponseOutput) ToAddonProfilesResponseOutputWithContext(ctx context.Context) AddonProfilesResponseOutput {
	return o
}

func (o AddonProfilesResponseOutput) Config() pulumi.StringMapOutput {
	return o.ApplyT(func(v AddonProfilesResponse) map[string]string { return v.Config }).(pulumi.StringMapOutput)
}

func (o AddonProfilesResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AddonProfilesResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type AddonProfilesResponseMapOutput struct{ *pulumi.OutputState }

func (AddonProfilesResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AddonProfilesResponse)(nil)).Elem()
}

func (o AddonProfilesResponseMapOutput) ToAddonProfilesResponseMapOutput() AddonProfilesResponseMapOutput {
	return o
}

func (o AddonProfilesResponseMapOutput) ToAddonProfilesResponseMapOutputWithContext(ctx context.Context) AddonProfilesResponseMapOutput {
	return o
}

func (o AddonProfilesResponseMapOutput) MapIndex(k pulumi.StringInput) AddonProfilesResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AddonProfilesResponse {
		return vs[0].(map[string]AddonProfilesResponse)[vs[1].(string)]
	}).(AddonProfilesResponseOutput)
}

type AddonStatusResponse struct {
	ErrorMessage *string `pulumi:"errorMessage"`
	Phase        *string `pulumi:"phase"`
	Ready        *bool   `pulumi:"ready"`
}

type AddonStatusResponseOutput struct{ *pulumi.OutputState }

func (AddonStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonStatusResponse)(nil)).Elem()
}

func (o AddonStatusResponseOutput) ToAddonStatusResponseOutput() AddonStatusResponseOutput {
	return o
}

func (o AddonStatusResponseOutput) ToAddonStatusResponseOutputWithContext(ctx context.Context) AddonStatusResponseOutput {
	return o
}

func (o AddonStatusResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddonStatusResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o AddonStatusResponseOutput) Phase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddonStatusResponse) *string { return v.Phase }).(pulumi.StringPtrOutput)
}

func (o AddonStatusResponseOutput) Ready() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AddonStatusResponse) *bool { return v.Ready }).(pulumi.BoolPtrOutput)
}

type AddonStatusResponseMapOutput struct{ *pulumi.OutputState }

func (AddonStatusResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AddonStatusResponse)(nil)).Elem()
}

func (o AddonStatusResponseMapOutput) ToAddonStatusResponseMapOutput() AddonStatusResponseMapOutput {
	return o
}

func (o AddonStatusResponseMapOutput) ToAddonStatusResponseMapOutputWithContext(ctx context.Context) AddonStatusResponseMapOutput {
	return o
}

func (o AddonStatusResponseMapOutput) MapIndex(k pulumi.StringInput) AddonStatusResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AddonStatusResponse {
		return vs[0].(map[string]AddonStatusResponse)[vs[1].(string)]
	}).(AddonStatusResponseOutput)
}

type AgentPoolExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type AgentPoolExtendedLocationInput interface {
	pulumi.Input

	ToAgentPoolExtendedLocationOutput() AgentPoolExtendedLocationOutput
	ToAgentPoolExtendedLocationOutputWithContext(context.Context) AgentPoolExtendedLocationOutput
}

type AgentPoolExtendedLocationArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (AgentPoolExtendedLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPoolExtendedLocation)(nil)).Elem()
}

func (i AgentPoolExtendedLocationArgs) ToAgentPoolExtendedLocationOutput() AgentPoolExtendedLocationOutput {
	return i.ToAgentPoolExtendedLocationOutputWithContext(context.Background())
}

func (i AgentPoolExtendedLocationArgs) ToAgentPoolExtendedLocationOutputWithContext(ctx context.Context) AgentPoolExtendedLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolExtendedLocationOutput)
}

func (i AgentPoolExtendedLocationArgs) ToAgentPoolExtendedLocationPtrOutput() AgentPoolExtendedLocationPtrOutput {
	return i.ToAgentPoolExtendedLocationPtrOutputWithContext(context.Background())
}

func (i AgentPoolExtendedLocationArgs) ToAgentPoolExtendedLocationPtrOutputWithContext(ctx context.Context) AgentPoolExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolExtendedLocationOutput).ToAgentPoolExtendedLocationPtrOutputWithContext(ctx)
}









type AgentPoolExtendedLocationPtrInput interface {
	pulumi.Input

	ToAgentPoolExtendedLocationPtrOutput() AgentPoolExtendedLocationPtrOutput
	ToAgentPoolExtendedLocationPtrOutputWithContext(context.Context) AgentPoolExtendedLocationPtrOutput
}

type agentPoolExtendedLocationPtrType AgentPoolExtendedLocationArgs

func AgentPoolExtendedLocationPtr(v *AgentPoolExtendedLocationArgs) AgentPoolExtendedLocationPtrInput {
	return (*agentPoolExtendedLocationPtrType)(v)
}

func (*agentPoolExtendedLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPoolExtendedLocation)(nil)).Elem()
}

func (i *agentPoolExtendedLocationPtrType) ToAgentPoolExtendedLocationPtrOutput() AgentPoolExtendedLocationPtrOutput {
	return i.ToAgentPoolExtendedLocationPtrOutputWithContext(context.Background())
}

func (i *agentPoolExtendedLocationPtrType) ToAgentPoolExtendedLocationPtrOutputWithContext(ctx context.Context) AgentPoolExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolExtendedLocationPtrOutput)
}

type AgentPoolExtendedLocationOutput struct{ *pulumi.OutputState }

func (AgentPoolExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPoolExtendedLocation)(nil)).Elem()
}

func (o AgentPoolExtendedLocationOutput) ToAgentPoolExtendedLocationOutput() AgentPoolExtendedLocationOutput {
	return o
}

func (o AgentPoolExtendedLocationOutput) ToAgentPoolExtendedLocationOutputWithContext(ctx context.Context) AgentPoolExtendedLocationOutput {
	return o
}

func (o AgentPoolExtendedLocationOutput) ToAgentPoolExtendedLocationPtrOutput() AgentPoolExtendedLocationPtrOutput {
	return o.ToAgentPoolExtendedLocationPtrOutputWithContext(context.Background())
}

func (o AgentPoolExtendedLocationOutput) ToAgentPoolExtendedLocationPtrOutputWithContext(ctx context.Context) AgentPoolExtendedLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AgentPoolExtendedLocation) *AgentPoolExtendedLocation {
		return &v
	}).(AgentPoolExtendedLocationPtrOutput)
}

func (o AgentPoolExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPoolExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o AgentPoolExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPoolExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type AgentPoolExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (AgentPoolExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPoolExtendedLocation)(nil)).Elem()
}

func (o AgentPoolExtendedLocationPtrOutput) ToAgentPoolExtendedLocationPtrOutput() AgentPoolExtendedLocationPtrOutput {
	return o
}

func (o AgentPoolExtendedLocationPtrOutput) ToAgentPoolExtendedLocationPtrOutputWithContext(ctx context.Context) AgentPoolExtendedLocationPtrOutput {
	return o
}

func (o AgentPoolExtendedLocationPtrOutput) Elem() AgentPoolExtendedLocationOutput {
	return o.ApplyT(func(v *AgentPoolExtendedLocation) AgentPoolExtendedLocation {
		if v != nil {
			return *v
		}
		var ret AgentPoolExtendedLocation
		return ret
	}).(AgentPoolExtendedLocationOutput)
}

func (o AgentPoolExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPoolExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o AgentPoolExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPoolExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type AgentPoolProvisioningStatusError struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}





type AgentPoolProvisioningStatusErrorInput interface {
	pulumi.Input

	ToAgentPoolProvisioningStatusErrorOutput() AgentPoolProvisioningStatusErrorOutput
	ToAgentPoolProvisioningStatusErrorOutputWithContext(context.Context) AgentPoolProvisioningStatusErrorOutput
}

type AgentPoolProvisioningStatusErrorArgs struct {
	Code    pulumi.StringPtrInput `pulumi:"code"`
	Message pulumi.StringPtrInput `pulumi:"message"`
}

func (AgentPoolProvisioningStatusErrorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPoolProvisioningStatusError)(nil)).Elem()
}

func (i AgentPoolProvisioningStatusErrorArgs) ToAgentPoolProvisioningStatusErrorOutput() AgentPoolProvisioningStatusErrorOutput {
	return i.ToAgentPoolProvisioningStatusErrorOutputWithContext(context.Background())
}

func (i AgentPoolProvisioningStatusErrorArgs) ToAgentPoolProvisioningStatusErrorOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusErrorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolProvisioningStatusErrorOutput)
}

func (i AgentPoolProvisioningStatusErrorArgs) ToAgentPoolProvisioningStatusErrorPtrOutput() AgentPoolProvisioningStatusErrorPtrOutput {
	return i.ToAgentPoolProvisioningStatusErrorPtrOutputWithContext(context.Background())
}

func (i AgentPoolProvisioningStatusErrorArgs) ToAgentPoolProvisioningStatusErrorPtrOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusErrorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolProvisioningStatusErrorOutput).ToAgentPoolProvisioningStatusErrorPtrOutputWithContext(ctx)
}









type AgentPoolProvisioningStatusErrorPtrInput interface {
	pulumi.Input

	ToAgentPoolProvisioningStatusErrorPtrOutput() AgentPoolProvisioningStatusErrorPtrOutput
	ToAgentPoolProvisioningStatusErrorPtrOutputWithContext(context.Context) AgentPoolProvisioningStatusErrorPtrOutput
}

type agentPoolProvisioningStatusErrorPtrType AgentPoolProvisioningStatusErrorArgs

func AgentPoolProvisioningStatusErrorPtr(v *AgentPoolProvisioningStatusErrorArgs) AgentPoolProvisioningStatusErrorPtrInput {
	return (*agentPoolProvisioningStatusErrorPtrType)(v)
}

func (*agentPoolProvisioningStatusErrorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPoolProvisioningStatusError)(nil)).Elem()
}

func (i *agentPoolProvisioningStatusErrorPtrType) ToAgentPoolProvisioningStatusErrorPtrOutput() AgentPoolProvisioningStatusErrorPtrOutput {
	return i.ToAgentPoolProvisioningStatusErrorPtrOutputWithContext(context.Background())
}

func (i *agentPoolProvisioningStatusErrorPtrType) ToAgentPoolProvisioningStatusErrorPtrOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusErrorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolProvisioningStatusErrorPtrOutput)
}

type AgentPoolProvisioningStatusErrorOutput struct{ *pulumi.OutputState }

func (AgentPoolProvisioningStatusErrorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPoolProvisioningStatusError)(nil)).Elem()
}

func (o AgentPoolProvisioningStatusErrorOutput) ToAgentPoolProvisioningStatusErrorOutput() AgentPoolProvisioningStatusErrorOutput {
	return o
}

func (o AgentPoolProvisioningStatusErrorOutput) ToAgentPoolProvisioningStatusErrorOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusErrorOutput {
	return o
}

func (o AgentPoolProvisioningStatusErrorOutput) ToAgentPoolProvisioningStatusErrorPtrOutput() AgentPoolProvisioningStatusErrorPtrOutput {
	return o.ToAgentPoolProvisioningStatusErrorPtrOutputWithContext(context.Background())
}

func (o AgentPoolProvisioningStatusErrorOutput) ToAgentPoolProvisioningStatusErrorPtrOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusErrorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AgentPoolProvisioningStatusError) *AgentPoolProvisioningStatusError {
		return &v
	}).(AgentPoolProvisioningStatusErrorPtrOutput)
}

func (o AgentPoolProvisioningStatusErrorOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPoolProvisioningStatusError) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o AgentPoolProvisioningStatusErrorOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPoolProvisioningStatusError) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type AgentPoolProvisioningStatusErrorPtrOutput struct{ *pulumi.OutputState }

func (AgentPoolProvisioningStatusErrorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPoolProvisioningStatusError)(nil)).Elem()
}

func (o AgentPoolProvisioningStatusErrorPtrOutput) ToAgentPoolProvisioningStatusErrorPtrOutput() AgentPoolProvisioningStatusErrorPtrOutput {
	return o
}

func (o AgentPoolProvisioningStatusErrorPtrOutput) ToAgentPoolProvisioningStatusErrorPtrOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusErrorPtrOutput {
	return o
}

func (o AgentPoolProvisioningStatusErrorPtrOutput) Elem() AgentPoolProvisioningStatusErrorOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusError) AgentPoolProvisioningStatusError {
		if v != nil {
			return *v
		}
		var ret AgentPoolProvisioningStatusError
		return ret
	}).(AgentPoolProvisioningStatusErrorOutput)
}

func (o AgentPoolProvisioningStatusErrorPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusError) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o AgentPoolProvisioningStatusErrorPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusError) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

type AgentPoolProvisioningStatusProvisioningStatus struct {
	Error       *AgentPoolProvisioningStatusError `pulumi:"error"`
	OperationId *string                           `pulumi:"operationId"`
	Phase       *string                           `pulumi:"phase"`
	Status      *string                           `pulumi:"status"`
}





type AgentPoolProvisioningStatusProvisioningStatusInput interface {
	pulumi.Input

	ToAgentPoolProvisioningStatusProvisioningStatusOutput() AgentPoolProvisioningStatusProvisioningStatusOutput
	ToAgentPoolProvisioningStatusProvisioningStatusOutputWithContext(context.Context) AgentPoolProvisioningStatusProvisioningStatusOutput
}

type AgentPoolProvisioningStatusProvisioningStatusArgs struct {
	Error       AgentPoolProvisioningStatusErrorPtrInput `pulumi:"error"`
	OperationId pulumi.StringPtrInput                    `pulumi:"operationId"`
	Phase       pulumi.StringPtrInput                    `pulumi:"phase"`
	Status      pulumi.StringPtrInput                    `pulumi:"status"`
}

func (AgentPoolProvisioningStatusProvisioningStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPoolProvisioningStatusProvisioningStatus)(nil)).Elem()
}

func (i AgentPoolProvisioningStatusProvisioningStatusArgs) ToAgentPoolProvisioningStatusProvisioningStatusOutput() AgentPoolProvisioningStatusProvisioningStatusOutput {
	return i.ToAgentPoolProvisioningStatusProvisioningStatusOutputWithContext(context.Background())
}

func (i AgentPoolProvisioningStatusProvisioningStatusArgs) ToAgentPoolProvisioningStatusProvisioningStatusOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusProvisioningStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolProvisioningStatusProvisioningStatusOutput)
}

func (i AgentPoolProvisioningStatusProvisioningStatusArgs) ToAgentPoolProvisioningStatusProvisioningStatusPtrOutput() AgentPoolProvisioningStatusProvisioningStatusPtrOutput {
	return i.ToAgentPoolProvisioningStatusProvisioningStatusPtrOutputWithContext(context.Background())
}

func (i AgentPoolProvisioningStatusProvisioningStatusArgs) ToAgentPoolProvisioningStatusProvisioningStatusPtrOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusProvisioningStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolProvisioningStatusProvisioningStatusOutput).ToAgentPoolProvisioningStatusProvisioningStatusPtrOutputWithContext(ctx)
}









type AgentPoolProvisioningStatusProvisioningStatusPtrInput interface {
	pulumi.Input

	ToAgentPoolProvisioningStatusProvisioningStatusPtrOutput() AgentPoolProvisioningStatusProvisioningStatusPtrOutput
	ToAgentPoolProvisioningStatusProvisioningStatusPtrOutputWithContext(context.Context) AgentPoolProvisioningStatusProvisioningStatusPtrOutput
}

type agentPoolProvisioningStatusProvisioningStatusPtrType AgentPoolProvisioningStatusProvisioningStatusArgs

func AgentPoolProvisioningStatusProvisioningStatusPtr(v *AgentPoolProvisioningStatusProvisioningStatusArgs) AgentPoolProvisioningStatusProvisioningStatusPtrInput {
	return (*agentPoolProvisioningStatusProvisioningStatusPtrType)(v)
}

func (*agentPoolProvisioningStatusProvisioningStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPoolProvisioningStatusProvisioningStatus)(nil)).Elem()
}

func (i *agentPoolProvisioningStatusProvisioningStatusPtrType) ToAgentPoolProvisioningStatusProvisioningStatusPtrOutput() AgentPoolProvisioningStatusProvisioningStatusPtrOutput {
	return i.ToAgentPoolProvisioningStatusProvisioningStatusPtrOutputWithContext(context.Background())
}

func (i *agentPoolProvisioningStatusProvisioningStatusPtrType) ToAgentPoolProvisioningStatusProvisioningStatusPtrOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusProvisioningStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolProvisioningStatusProvisioningStatusPtrOutput)
}

type AgentPoolProvisioningStatusProvisioningStatusOutput struct{ *pulumi.OutputState }

func (AgentPoolProvisioningStatusProvisioningStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPoolProvisioningStatusProvisioningStatus)(nil)).Elem()
}

func (o AgentPoolProvisioningStatusProvisioningStatusOutput) ToAgentPoolProvisioningStatusProvisioningStatusOutput() AgentPoolProvisioningStatusProvisioningStatusOutput {
	return o
}

func (o AgentPoolProvisioningStatusProvisioningStatusOutput) ToAgentPoolProvisioningStatusProvisioningStatusOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusProvisioningStatusOutput {
	return o
}

func (o AgentPoolProvisioningStatusProvisioningStatusOutput) ToAgentPoolProvisioningStatusProvisioningStatusPtrOutput() AgentPoolProvisioningStatusProvisioningStatusPtrOutput {
	return o.ToAgentPoolProvisioningStatusProvisioningStatusPtrOutputWithContext(context.Background())
}

func (o AgentPoolProvisioningStatusProvisioningStatusOutput) ToAgentPoolProvisioningStatusProvisioningStatusPtrOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusProvisioningStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AgentPoolProvisioningStatusProvisioningStatus) *AgentPoolProvisioningStatusProvisioningStatus {
		return &v
	}).(AgentPoolProvisioningStatusProvisioningStatusPtrOutput)
}

func (o AgentPoolProvisioningStatusProvisioningStatusOutput) Error() AgentPoolProvisioningStatusErrorPtrOutput {
	return o.ApplyT(func(v AgentPoolProvisioningStatusProvisioningStatus) *AgentPoolProvisioningStatusError {
		return v.Error
	}).(AgentPoolProvisioningStatusErrorPtrOutput)
}

func (o AgentPoolProvisioningStatusProvisioningStatusOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPoolProvisioningStatusProvisioningStatus) *string { return v.OperationId }).(pulumi.StringPtrOutput)
}

func (o AgentPoolProvisioningStatusProvisioningStatusOutput) Phase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPoolProvisioningStatusProvisioningStatus) *string { return v.Phase }).(pulumi.StringPtrOutput)
}

func (o AgentPoolProvisioningStatusProvisioningStatusOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPoolProvisioningStatusProvisioningStatus) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type AgentPoolProvisioningStatusProvisioningStatusPtrOutput struct{ *pulumi.OutputState }

func (AgentPoolProvisioningStatusProvisioningStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPoolProvisioningStatusProvisioningStatus)(nil)).Elem()
}

func (o AgentPoolProvisioningStatusProvisioningStatusPtrOutput) ToAgentPoolProvisioningStatusProvisioningStatusPtrOutput() AgentPoolProvisioningStatusProvisioningStatusPtrOutput {
	return o
}

func (o AgentPoolProvisioningStatusProvisioningStatusPtrOutput) ToAgentPoolProvisioningStatusProvisioningStatusPtrOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusProvisioningStatusPtrOutput {
	return o
}

func (o AgentPoolProvisioningStatusProvisioningStatusPtrOutput) Elem() AgentPoolProvisioningStatusProvisioningStatusOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusProvisioningStatus) AgentPoolProvisioningStatusProvisioningStatus {
		if v != nil {
			return *v
		}
		var ret AgentPoolProvisioningStatusProvisioningStatus
		return ret
	}).(AgentPoolProvisioningStatusProvisioningStatusOutput)
}

func (o AgentPoolProvisioningStatusProvisioningStatusPtrOutput) Error() AgentPoolProvisioningStatusErrorPtrOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusProvisioningStatus) *AgentPoolProvisioningStatusError {
		if v == nil {
			return nil
		}
		return v.Error
	}).(AgentPoolProvisioningStatusErrorPtrOutput)
}

func (o AgentPoolProvisioningStatusProvisioningStatusPtrOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.OperationId
	}).(pulumi.StringPtrOutput)
}

func (o AgentPoolProvisioningStatusProvisioningStatusPtrOutput) Phase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.Phase
	}).(pulumi.StringPtrOutput)
}

func (o AgentPoolProvisioningStatusProvisioningStatusPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type AgentPoolProvisioningStatusResponseError struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}

type AgentPoolProvisioningStatusResponseErrorOutput struct{ *pulumi.OutputState }

func (AgentPoolProvisioningStatusResponseErrorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPoolProvisioningStatusResponseError)(nil)).Elem()
}

func (o AgentPoolProvisioningStatusResponseErrorOutput) ToAgentPoolProvisioningStatusResponseErrorOutput() AgentPoolProvisioningStatusResponseErrorOutput {
	return o
}

func (o AgentPoolProvisioningStatusResponseErrorOutput) ToAgentPoolProvisioningStatusResponseErrorOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusResponseErrorOutput {
	return o
}

func (o AgentPoolProvisioningStatusResponseErrorOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPoolProvisioningStatusResponseError) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o AgentPoolProvisioningStatusResponseErrorOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPoolProvisioningStatusResponseError) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type AgentPoolProvisioningStatusResponseErrorPtrOutput struct{ *pulumi.OutputState }

func (AgentPoolProvisioningStatusResponseErrorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPoolProvisioningStatusResponseError)(nil)).Elem()
}

func (o AgentPoolProvisioningStatusResponseErrorPtrOutput) ToAgentPoolProvisioningStatusResponseErrorPtrOutput() AgentPoolProvisioningStatusResponseErrorPtrOutput {
	return o
}

func (o AgentPoolProvisioningStatusResponseErrorPtrOutput) ToAgentPoolProvisioningStatusResponseErrorPtrOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusResponseErrorPtrOutput {
	return o
}

func (o AgentPoolProvisioningStatusResponseErrorPtrOutput) Elem() AgentPoolProvisioningStatusResponseErrorOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusResponseError) AgentPoolProvisioningStatusResponseError {
		if v != nil {
			return *v
		}
		var ret AgentPoolProvisioningStatusResponseError
		return ret
	}).(AgentPoolProvisioningStatusResponseErrorOutput)
}

func (o AgentPoolProvisioningStatusResponseErrorPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusResponseError) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o AgentPoolProvisioningStatusResponseErrorPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusResponseError) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

type AgentPoolProvisioningStatusResponseProvisioningStatus struct {
	Error       *AgentPoolProvisioningStatusResponseError `pulumi:"error"`
	OperationId *string                                   `pulumi:"operationId"`
	Phase       *string                                   `pulumi:"phase"`
	Status      *string                                   `pulumi:"status"`
}

type AgentPoolProvisioningStatusResponseProvisioningStatusOutput struct{ *pulumi.OutputState }

func (AgentPoolProvisioningStatusResponseProvisioningStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPoolProvisioningStatusResponseProvisioningStatus)(nil)).Elem()
}

func (o AgentPoolProvisioningStatusResponseProvisioningStatusOutput) ToAgentPoolProvisioningStatusResponseProvisioningStatusOutput() AgentPoolProvisioningStatusResponseProvisioningStatusOutput {
	return o
}

func (o AgentPoolProvisioningStatusResponseProvisioningStatusOutput) ToAgentPoolProvisioningStatusResponseProvisioningStatusOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusResponseProvisioningStatusOutput {
	return o
}

func (o AgentPoolProvisioningStatusResponseProvisioningStatusOutput) Error() AgentPoolProvisioningStatusResponseErrorPtrOutput {
	return o.ApplyT(func(v AgentPoolProvisioningStatusResponseProvisioningStatus) *AgentPoolProvisioningStatusResponseError {
		return v.Error
	}).(AgentPoolProvisioningStatusResponseErrorPtrOutput)
}

func (o AgentPoolProvisioningStatusResponseProvisioningStatusOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPoolProvisioningStatusResponseProvisioningStatus) *string { return v.OperationId }).(pulumi.StringPtrOutput)
}

func (o AgentPoolProvisioningStatusResponseProvisioningStatusOutput) Phase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPoolProvisioningStatusResponseProvisioningStatus) *string { return v.Phase }).(pulumi.StringPtrOutput)
}

func (o AgentPoolProvisioningStatusResponseProvisioningStatusOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPoolProvisioningStatusResponseProvisioningStatus) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type AgentPoolProvisioningStatusResponseProvisioningStatusPtrOutput struct{ *pulumi.OutputState }

func (AgentPoolProvisioningStatusResponseProvisioningStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPoolProvisioningStatusResponseProvisioningStatus)(nil)).Elem()
}

func (o AgentPoolProvisioningStatusResponseProvisioningStatusPtrOutput) ToAgentPoolProvisioningStatusResponseProvisioningStatusPtrOutput() AgentPoolProvisioningStatusResponseProvisioningStatusPtrOutput {
	return o
}

func (o AgentPoolProvisioningStatusResponseProvisioningStatusPtrOutput) ToAgentPoolProvisioningStatusResponseProvisioningStatusPtrOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusResponseProvisioningStatusPtrOutput {
	return o
}

func (o AgentPoolProvisioningStatusResponseProvisioningStatusPtrOutput) Elem() AgentPoolProvisioningStatusResponseProvisioningStatusOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusResponseProvisioningStatus) AgentPoolProvisioningStatusResponseProvisioningStatus {
		if v != nil {
			return *v
		}
		var ret AgentPoolProvisioningStatusResponseProvisioningStatus
		return ret
	}).(AgentPoolProvisioningStatusResponseProvisioningStatusOutput)
}

func (o AgentPoolProvisioningStatusResponseProvisioningStatusPtrOutput) Error() AgentPoolProvisioningStatusResponseErrorPtrOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusResponseProvisioningStatus) *AgentPoolProvisioningStatusResponseError {
		if v == nil {
			return nil
		}
		return v.Error
	}).(AgentPoolProvisioningStatusResponseErrorPtrOutput)
}

func (o AgentPoolProvisioningStatusResponseProvisioningStatusPtrOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.OperationId
	}).(pulumi.StringPtrOutput)
}

func (o AgentPoolProvisioningStatusResponseProvisioningStatusPtrOutput) Phase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.Phase
	}).(pulumi.StringPtrOutput)
}

func (o AgentPoolProvisioningStatusResponseProvisioningStatusPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type AgentPoolProvisioningStatusResponseStatus struct {
	ErrorMessage       *string                                                `pulumi:"errorMessage"`
	ProvisioningStatus *AgentPoolProvisioningStatusResponseProvisioningStatus `pulumi:"provisioningStatus"`
	ReadyReplicas      *int                                                   `pulumi:"readyReplicas"`
	Replicas           *int                                                   `pulumi:"replicas"`
}

type AgentPoolProvisioningStatusResponseStatusOutput struct{ *pulumi.OutputState }

func (AgentPoolProvisioningStatusResponseStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPoolProvisioningStatusResponseStatus)(nil)).Elem()
}

func (o AgentPoolProvisioningStatusResponseStatusOutput) ToAgentPoolProvisioningStatusResponseStatusOutput() AgentPoolProvisioningStatusResponseStatusOutput {
	return o
}

func (o AgentPoolProvisioningStatusResponseStatusOutput) ToAgentPoolProvisioningStatusResponseStatusOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusResponseStatusOutput {
	return o
}

func (o AgentPoolProvisioningStatusResponseStatusOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPoolProvisioningStatusResponseStatus) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o AgentPoolProvisioningStatusResponseStatusOutput) ProvisioningStatus() AgentPoolProvisioningStatusResponseProvisioningStatusPtrOutput {
	return o.ApplyT(func(v AgentPoolProvisioningStatusResponseStatus) *AgentPoolProvisioningStatusResponseProvisioningStatus {
		return v.ProvisioningStatus
	}).(AgentPoolProvisioningStatusResponseProvisioningStatusPtrOutput)
}

func (o AgentPoolProvisioningStatusResponseStatusOutput) ReadyReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AgentPoolProvisioningStatusResponseStatus) *int { return v.ReadyReplicas }).(pulumi.IntPtrOutput)
}

func (o AgentPoolProvisioningStatusResponseStatusOutput) Replicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AgentPoolProvisioningStatusResponseStatus) *int { return v.Replicas }).(pulumi.IntPtrOutput)
}

type AgentPoolProvisioningStatusResponseStatusPtrOutput struct{ *pulumi.OutputState }

func (AgentPoolProvisioningStatusResponseStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPoolProvisioningStatusResponseStatus)(nil)).Elem()
}

func (o AgentPoolProvisioningStatusResponseStatusPtrOutput) ToAgentPoolProvisioningStatusResponseStatusPtrOutput() AgentPoolProvisioningStatusResponseStatusPtrOutput {
	return o
}

func (o AgentPoolProvisioningStatusResponseStatusPtrOutput) ToAgentPoolProvisioningStatusResponseStatusPtrOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusResponseStatusPtrOutput {
	return o
}

func (o AgentPoolProvisioningStatusResponseStatusPtrOutput) Elem() AgentPoolProvisioningStatusResponseStatusOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusResponseStatus) AgentPoolProvisioningStatusResponseStatus {
		if v != nil {
			return *v
		}
		var ret AgentPoolProvisioningStatusResponseStatus
		return ret
	}).(AgentPoolProvisioningStatusResponseStatusOutput)
}

func (o AgentPoolProvisioningStatusResponseStatusPtrOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusResponseStatus) *string {
		if v == nil {
			return nil
		}
		return v.ErrorMessage
	}).(pulumi.StringPtrOutput)
}

func (o AgentPoolProvisioningStatusResponseStatusPtrOutput) ProvisioningStatus() AgentPoolProvisioningStatusResponseProvisioningStatusPtrOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusResponseStatus) *AgentPoolProvisioningStatusResponseProvisioningStatus {
		if v == nil {
			return nil
		}
		return v.ProvisioningStatus
	}).(AgentPoolProvisioningStatusResponseProvisioningStatusPtrOutput)
}

func (o AgentPoolProvisioningStatusResponseStatusPtrOutput) ReadyReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusResponseStatus) *int {
		if v == nil {
			return nil
		}
		return v.ReadyReplicas
	}).(pulumi.IntPtrOutput)
}

func (o AgentPoolProvisioningStatusResponseStatusPtrOutput) Replicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusResponseStatus) *int {
		if v == nil {
			return nil
		}
		return v.Replicas
	}).(pulumi.IntPtrOutput)
}

type AgentPoolProvisioningStatusStatus struct {
	ErrorMessage       *string                                        `pulumi:"errorMessage"`
	ProvisioningStatus *AgentPoolProvisioningStatusProvisioningStatus `pulumi:"provisioningStatus"`
	ReadyReplicas      *int                                           `pulumi:"readyReplicas"`
	Replicas           *int                                           `pulumi:"replicas"`
}





type AgentPoolProvisioningStatusStatusInput interface {
	pulumi.Input

	ToAgentPoolProvisioningStatusStatusOutput() AgentPoolProvisioningStatusStatusOutput
	ToAgentPoolProvisioningStatusStatusOutputWithContext(context.Context) AgentPoolProvisioningStatusStatusOutput
}

type AgentPoolProvisioningStatusStatusArgs struct {
	ErrorMessage       pulumi.StringPtrInput                                 `pulumi:"errorMessage"`
	ProvisioningStatus AgentPoolProvisioningStatusProvisioningStatusPtrInput `pulumi:"provisioningStatus"`
	ReadyReplicas      pulumi.IntPtrInput                                    `pulumi:"readyReplicas"`
	Replicas           pulumi.IntPtrInput                                    `pulumi:"replicas"`
}

func (AgentPoolProvisioningStatusStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPoolProvisioningStatusStatus)(nil)).Elem()
}

func (i AgentPoolProvisioningStatusStatusArgs) ToAgentPoolProvisioningStatusStatusOutput() AgentPoolProvisioningStatusStatusOutput {
	return i.ToAgentPoolProvisioningStatusStatusOutputWithContext(context.Background())
}

func (i AgentPoolProvisioningStatusStatusArgs) ToAgentPoolProvisioningStatusStatusOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolProvisioningStatusStatusOutput)
}

func (i AgentPoolProvisioningStatusStatusArgs) ToAgentPoolProvisioningStatusStatusPtrOutput() AgentPoolProvisioningStatusStatusPtrOutput {
	return i.ToAgentPoolProvisioningStatusStatusPtrOutputWithContext(context.Background())
}

func (i AgentPoolProvisioningStatusStatusArgs) ToAgentPoolProvisioningStatusStatusPtrOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolProvisioningStatusStatusOutput).ToAgentPoolProvisioningStatusStatusPtrOutputWithContext(ctx)
}









type AgentPoolProvisioningStatusStatusPtrInput interface {
	pulumi.Input

	ToAgentPoolProvisioningStatusStatusPtrOutput() AgentPoolProvisioningStatusStatusPtrOutput
	ToAgentPoolProvisioningStatusStatusPtrOutputWithContext(context.Context) AgentPoolProvisioningStatusStatusPtrOutput
}

type agentPoolProvisioningStatusStatusPtrType AgentPoolProvisioningStatusStatusArgs

func AgentPoolProvisioningStatusStatusPtr(v *AgentPoolProvisioningStatusStatusArgs) AgentPoolProvisioningStatusStatusPtrInput {
	return (*agentPoolProvisioningStatusStatusPtrType)(v)
}

func (*agentPoolProvisioningStatusStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPoolProvisioningStatusStatus)(nil)).Elem()
}

func (i *agentPoolProvisioningStatusStatusPtrType) ToAgentPoolProvisioningStatusStatusPtrOutput() AgentPoolProvisioningStatusStatusPtrOutput {
	return i.ToAgentPoolProvisioningStatusStatusPtrOutputWithContext(context.Background())
}

func (i *agentPoolProvisioningStatusStatusPtrType) ToAgentPoolProvisioningStatusStatusPtrOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolProvisioningStatusStatusPtrOutput)
}

type AgentPoolProvisioningStatusStatusOutput struct{ *pulumi.OutputState }

func (AgentPoolProvisioningStatusStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPoolProvisioningStatusStatus)(nil)).Elem()
}

func (o AgentPoolProvisioningStatusStatusOutput) ToAgentPoolProvisioningStatusStatusOutput() AgentPoolProvisioningStatusStatusOutput {
	return o
}

func (o AgentPoolProvisioningStatusStatusOutput) ToAgentPoolProvisioningStatusStatusOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusStatusOutput {
	return o
}

func (o AgentPoolProvisioningStatusStatusOutput) ToAgentPoolProvisioningStatusStatusPtrOutput() AgentPoolProvisioningStatusStatusPtrOutput {
	return o.ToAgentPoolProvisioningStatusStatusPtrOutputWithContext(context.Background())
}

func (o AgentPoolProvisioningStatusStatusOutput) ToAgentPoolProvisioningStatusStatusPtrOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AgentPoolProvisioningStatusStatus) *AgentPoolProvisioningStatusStatus {
		return &v
	}).(AgentPoolProvisioningStatusStatusPtrOutput)
}

func (o AgentPoolProvisioningStatusStatusOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPoolProvisioningStatusStatus) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o AgentPoolProvisioningStatusStatusOutput) ProvisioningStatus() AgentPoolProvisioningStatusProvisioningStatusPtrOutput {
	return o.ApplyT(func(v AgentPoolProvisioningStatusStatus) *AgentPoolProvisioningStatusProvisioningStatus {
		return v.ProvisioningStatus
	}).(AgentPoolProvisioningStatusProvisioningStatusPtrOutput)
}

func (o AgentPoolProvisioningStatusStatusOutput) ReadyReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AgentPoolProvisioningStatusStatus) *int { return v.ReadyReplicas }).(pulumi.IntPtrOutput)
}

func (o AgentPoolProvisioningStatusStatusOutput) Replicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AgentPoolProvisioningStatusStatus) *int { return v.Replicas }).(pulumi.IntPtrOutput)
}

type AgentPoolProvisioningStatusStatusPtrOutput struct{ *pulumi.OutputState }

func (AgentPoolProvisioningStatusStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPoolProvisioningStatusStatus)(nil)).Elem()
}

func (o AgentPoolProvisioningStatusStatusPtrOutput) ToAgentPoolProvisioningStatusStatusPtrOutput() AgentPoolProvisioningStatusStatusPtrOutput {
	return o
}

func (o AgentPoolProvisioningStatusStatusPtrOutput) ToAgentPoolProvisioningStatusStatusPtrOutputWithContext(ctx context.Context) AgentPoolProvisioningStatusStatusPtrOutput {
	return o
}

func (o AgentPoolProvisioningStatusStatusPtrOutput) Elem() AgentPoolProvisioningStatusStatusOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusStatus) AgentPoolProvisioningStatusStatus {
		if v != nil {
			return *v
		}
		var ret AgentPoolProvisioningStatusStatus
		return ret
	}).(AgentPoolProvisioningStatusStatusOutput)
}

func (o AgentPoolProvisioningStatusStatusPtrOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusStatus) *string {
		if v == nil {
			return nil
		}
		return v.ErrorMessage
	}).(pulumi.StringPtrOutput)
}

func (o AgentPoolProvisioningStatusStatusPtrOutput) ProvisioningStatus() AgentPoolProvisioningStatusProvisioningStatusPtrOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusStatus) *AgentPoolProvisioningStatusProvisioningStatus {
		if v == nil {
			return nil
		}
		return v.ProvisioningStatus
	}).(AgentPoolProvisioningStatusProvisioningStatusPtrOutput)
}

func (o AgentPoolProvisioningStatusStatusPtrOutput) ReadyReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusStatus) *int {
		if v == nil {
			return nil
		}
		return v.ReadyReplicas
	}).(pulumi.IntPtrOutput)
}

func (o AgentPoolProvisioningStatusStatusPtrOutput) Replicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPoolProvisioningStatusStatus) *int {
		if v == nil {
			return nil
		}
		return v.Replicas
	}).(pulumi.IntPtrOutput)
}

type AgentPoolResponseExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type AgentPoolResponseExtendedLocationOutput struct{ *pulumi.OutputState }

func (AgentPoolResponseExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPoolResponseExtendedLocation)(nil)).Elem()
}

func (o AgentPoolResponseExtendedLocationOutput) ToAgentPoolResponseExtendedLocationOutput() AgentPoolResponseExtendedLocationOutput {
	return o
}

func (o AgentPoolResponseExtendedLocationOutput) ToAgentPoolResponseExtendedLocationOutputWithContext(ctx context.Context) AgentPoolResponseExtendedLocationOutput {
	return o
}

func (o AgentPoolResponseExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPoolResponseExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o AgentPoolResponseExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPoolResponseExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type AgentPoolResponseExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (AgentPoolResponseExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPoolResponseExtendedLocation)(nil)).Elem()
}

func (o AgentPoolResponseExtendedLocationPtrOutput) ToAgentPoolResponseExtendedLocationPtrOutput() AgentPoolResponseExtendedLocationPtrOutput {
	return o
}

func (o AgentPoolResponseExtendedLocationPtrOutput) ToAgentPoolResponseExtendedLocationPtrOutputWithContext(ctx context.Context) AgentPoolResponseExtendedLocationPtrOutput {
	return o
}

func (o AgentPoolResponseExtendedLocationPtrOutput) Elem() AgentPoolResponseExtendedLocationOutput {
	return o.ApplyT(func(v *AgentPoolResponseExtendedLocation) AgentPoolResponseExtendedLocation {
		if v != nil {
			return *v
		}
		var ret AgentPoolResponseExtendedLocation
		return ret
	}).(AgentPoolResponseExtendedLocationOutput)
}

func (o AgentPoolResponseExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPoolResponseExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o AgentPoolResponseExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPoolResponseExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ArcAgentProfile struct {
	AgentAutoUpgrade *string `pulumi:"agentAutoUpgrade"`
	AgentVersion     *string `pulumi:"agentVersion"`
}


func (val *ArcAgentProfile) Defaults() *ArcAgentProfile {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AgentAutoUpgrade) {
		agentAutoUpgrade_ := "Enabled"
		tmp.AgentAutoUpgrade = &agentAutoUpgrade_
	}
	return &tmp
}





type ArcAgentProfileInput interface {
	pulumi.Input

	ToArcAgentProfileOutput() ArcAgentProfileOutput
	ToArcAgentProfileOutputWithContext(context.Context) ArcAgentProfileOutput
}

type ArcAgentProfileArgs struct {
	AgentAutoUpgrade pulumi.StringPtrInput `pulumi:"agentAutoUpgrade"`
	AgentVersion     pulumi.StringPtrInput `pulumi:"agentVersion"`
}


func (val *ArcAgentProfileArgs) Defaults() *ArcAgentProfileArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AgentAutoUpgrade) {
		tmp.AgentAutoUpgrade = pulumi.StringPtr("Enabled")
	}
	return &tmp
}
func (ArcAgentProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArcAgentProfile)(nil)).Elem()
}

func (i ArcAgentProfileArgs) ToArcAgentProfileOutput() ArcAgentProfileOutput {
	return i.ToArcAgentProfileOutputWithContext(context.Background())
}

func (i ArcAgentProfileArgs) ToArcAgentProfileOutputWithContext(ctx context.Context) ArcAgentProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArcAgentProfileOutput)
}

func (i ArcAgentProfileArgs) ToArcAgentProfilePtrOutput() ArcAgentProfilePtrOutput {
	return i.ToArcAgentProfilePtrOutputWithContext(context.Background())
}

func (i ArcAgentProfileArgs) ToArcAgentProfilePtrOutputWithContext(ctx context.Context) ArcAgentProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArcAgentProfileOutput).ToArcAgentProfilePtrOutputWithContext(ctx)
}









type ArcAgentProfilePtrInput interface {
	pulumi.Input

	ToArcAgentProfilePtrOutput() ArcAgentProfilePtrOutput
	ToArcAgentProfilePtrOutputWithContext(context.Context) ArcAgentProfilePtrOutput
}

type arcAgentProfilePtrType ArcAgentProfileArgs

func ArcAgentProfilePtr(v *ArcAgentProfileArgs) ArcAgentProfilePtrInput {
	return (*arcAgentProfilePtrType)(v)
}

func (*arcAgentProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ArcAgentProfile)(nil)).Elem()
}

func (i *arcAgentProfilePtrType) ToArcAgentProfilePtrOutput() ArcAgentProfilePtrOutput {
	return i.ToArcAgentProfilePtrOutputWithContext(context.Background())
}

func (i *arcAgentProfilePtrType) ToArcAgentProfilePtrOutputWithContext(ctx context.Context) ArcAgentProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArcAgentProfilePtrOutput)
}

type ArcAgentProfileOutput struct{ *pulumi.OutputState }

func (ArcAgentProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArcAgentProfile)(nil)).Elem()
}

func (o ArcAgentProfileOutput) ToArcAgentProfileOutput() ArcAgentProfileOutput {
	return o
}

func (o ArcAgentProfileOutput) ToArcAgentProfileOutputWithContext(ctx context.Context) ArcAgentProfileOutput {
	return o
}

func (o ArcAgentProfileOutput) ToArcAgentProfilePtrOutput() ArcAgentProfilePtrOutput {
	return o.ToArcAgentProfilePtrOutputWithContext(context.Background())
}

func (o ArcAgentProfileOutput) ToArcAgentProfilePtrOutputWithContext(ctx context.Context) ArcAgentProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ArcAgentProfile) *ArcAgentProfile {
		return &v
	}).(ArcAgentProfilePtrOutput)
}

func (o ArcAgentProfileOutput) AgentAutoUpgrade() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcAgentProfile) *string { return v.AgentAutoUpgrade }).(pulumi.StringPtrOutput)
}

func (o ArcAgentProfileOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcAgentProfile) *string { return v.AgentVersion }).(pulumi.StringPtrOutput)
}

type ArcAgentProfilePtrOutput struct{ *pulumi.OutputState }

func (ArcAgentProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArcAgentProfile)(nil)).Elem()
}

func (o ArcAgentProfilePtrOutput) ToArcAgentProfilePtrOutput() ArcAgentProfilePtrOutput {
	return o
}

func (o ArcAgentProfilePtrOutput) ToArcAgentProfilePtrOutputWithContext(ctx context.Context) ArcAgentProfilePtrOutput {
	return o
}

func (o ArcAgentProfilePtrOutput) Elem() ArcAgentProfileOutput {
	return o.ApplyT(func(v *ArcAgentProfile) ArcAgentProfile {
		if v != nil {
			return *v
		}
		var ret ArcAgentProfile
		return ret
	}).(ArcAgentProfileOutput)
}

func (o ArcAgentProfilePtrOutput) AgentAutoUpgrade() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcAgentProfile) *string {
		if v == nil {
			return nil
		}
		return v.AgentAutoUpgrade
	}).(pulumi.StringPtrOutput)
}

func (o ArcAgentProfilePtrOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcAgentProfile) *string {
		if v == nil {
			return nil
		}
		return v.AgentVersion
	}).(pulumi.StringPtrOutput)
}

type ArcAgentProfileResponse struct {
	AgentAutoUpgrade *string `pulumi:"agentAutoUpgrade"`
	AgentVersion     *string `pulumi:"agentVersion"`
}


func (val *ArcAgentProfileResponse) Defaults() *ArcAgentProfileResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AgentAutoUpgrade) {
		agentAutoUpgrade_ := "Enabled"
		tmp.AgentAutoUpgrade = &agentAutoUpgrade_
	}
	return &tmp
}

type ArcAgentProfileResponseOutput struct{ *pulumi.OutputState }

func (ArcAgentProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArcAgentProfileResponse)(nil)).Elem()
}

func (o ArcAgentProfileResponseOutput) ToArcAgentProfileResponseOutput() ArcAgentProfileResponseOutput {
	return o
}

func (o ArcAgentProfileResponseOutput) ToArcAgentProfileResponseOutputWithContext(ctx context.Context) ArcAgentProfileResponseOutput {
	return o
}

func (o ArcAgentProfileResponseOutput) AgentAutoUpgrade() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcAgentProfileResponse) *string { return v.AgentAutoUpgrade }).(pulumi.StringPtrOutput)
}

func (o ArcAgentProfileResponseOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcAgentProfileResponse) *string { return v.AgentVersion }).(pulumi.StringPtrOutput)
}

type ArcAgentProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ArcAgentProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArcAgentProfileResponse)(nil)).Elem()
}

func (o ArcAgentProfileResponsePtrOutput) ToArcAgentProfileResponsePtrOutput() ArcAgentProfileResponsePtrOutput {
	return o
}

func (o ArcAgentProfileResponsePtrOutput) ToArcAgentProfileResponsePtrOutputWithContext(ctx context.Context) ArcAgentProfileResponsePtrOutput {
	return o
}

func (o ArcAgentProfileResponsePtrOutput) Elem() ArcAgentProfileResponseOutput {
	return o.ApplyT(func(v *ArcAgentProfileResponse) ArcAgentProfileResponse {
		if v != nil {
			return *v
		}
		var ret ArcAgentProfileResponse
		return ret
	}).(ArcAgentProfileResponseOutput)
}

func (o ArcAgentProfileResponsePtrOutput) AgentAutoUpgrade() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcAgentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AgentAutoUpgrade
	}).(pulumi.StringPtrOutput)
}

func (o ArcAgentProfileResponsePtrOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcAgentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AgentVersion
	}).(pulumi.StringPtrOutput)
}

type ArcAgentStatusResponse struct {
	AgentVersion                             *string  `pulumi:"agentVersion"`
	CoreCount                                *float64 `pulumi:"coreCount"`
	DeploymentState                          *string  `pulumi:"deploymentState"`
	ErrorMessage                             *string  `pulumi:"errorMessage"`
	LastConnectivityTime                     *string  `pulumi:"lastConnectivityTime"`
	ManagedIdentityCertificateExpirationTime *string  `pulumi:"managedIdentityCertificateExpirationTime"`
	OnboardingPublicKey                      *string  `pulumi:"onboardingPublicKey"`
}

type ArcAgentStatusResponseOutput struct{ *pulumi.OutputState }

func (ArcAgentStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArcAgentStatusResponse)(nil)).Elem()
}

func (o ArcAgentStatusResponseOutput) ToArcAgentStatusResponseOutput() ArcAgentStatusResponseOutput {
	return o
}

func (o ArcAgentStatusResponseOutput) ToArcAgentStatusResponseOutputWithContext(ctx context.Context) ArcAgentStatusResponseOutput {
	return o
}

func (o ArcAgentStatusResponseOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcAgentStatusResponse) *string { return v.AgentVersion }).(pulumi.StringPtrOutput)
}

func (o ArcAgentStatusResponseOutput) CoreCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ArcAgentStatusResponse) *float64 { return v.CoreCount }).(pulumi.Float64PtrOutput)
}

func (o ArcAgentStatusResponseOutput) DeploymentState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcAgentStatusResponse) *string { return v.DeploymentState }).(pulumi.StringPtrOutput)
}

func (o ArcAgentStatusResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcAgentStatusResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o ArcAgentStatusResponseOutput) LastConnectivityTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcAgentStatusResponse) *string { return v.LastConnectivityTime }).(pulumi.StringPtrOutput)
}

func (o ArcAgentStatusResponseOutput) ManagedIdentityCertificateExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcAgentStatusResponse) *string { return v.ManagedIdentityCertificateExpirationTime }).(pulumi.StringPtrOutput)
}

func (o ArcAgentStatusResponseOutput) OnboardingPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcAgentStatusResponse) *string { return v.OnboardingPublicKey }).(pulumi.StringPtrOutput)
}

type ArcAgentStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (ArcAgentStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArcAgentStatusResponse)(nil)).Elem()
}

func (o ArcAgentStatusResponsePtrOutput) ToArcAgentStatusResponsePtrOutput() ArcAgentStatusResponsePtrOutput {
	return o
}

func (o ArcAgentStatusResponsePtrOutput) ToArcAgentStatusResponsePtrOutputWithContext(ctx context.Context) ArcAgentStatusResponsePtrOutput {
	return o
}

func (o ArcAgentStatusResponsePtrOutput) Elem() ArcAgentStatusResponseOutput {
	return o.ApplyT(func(v *ArcAgentStatusResponse) ArcAgentStatusResponse {
		if v != nil {
			return *v
		}
		var ret ArcAgentStatusResponse
		return ret
	}).(ArcAgentStatusResponseOutput)
}

func (o ArcAgentStatusResponsePtrOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcAgentStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.AgentVersion
	}).(pulumi.StringPtrOutput)
}

func (o ArcAgentStatusResponsePtrOutput) CoreCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ArcAgentStatusResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.CoreCount
	}).(pulumi.Float64PtrOutput)
}

func (o ArcAgentStatusResponsePtrOutput) DeploymentState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcAgentStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentState
	}).(pulumi.StringPtrOutput)
}

func (o ArcAgentStatusResponsePtrOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcAgentStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.ErrorMessage
	}).(pulumi.StringPtrOutput)
}

func (o ArcAgentStatusResponsePtrOutput) LastConnectivityTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcAgentStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastConnectivityTime
	}).(pulumi.StringPtrOutput)
}

func (o ArcAgentStatusResponsePtrOutput) ManagedIdentityCertificateExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcAgentStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.ManagedIdentityCertificateExpirationTime
	}).(pulumi.StringPtrOutput)
}

func (o ArcAgentStatusResponsePtrOutput) OnboardingPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcAgentStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.OnboardingPublicKey
	}).(pulumi.StringPtrOutput)
}

type CloudProviderProfile struct {
	InfraNetworkProfile *CloudProviderProfileInfraNetworkProfile `pulumi:"infraNetworkProfile"`
	InfraStorageProfile *CloudProviderProfileInfraStorageProfile `pulumi:"infraStorageProfile"`
}





type CloudProviderProfileInput interface {
	pulumi.Input

	ToCloudProviderProfileOutput() CloudProviderProfileOutput
	ToCloudProviderProfileOutputWithContext(context.Context) CloudProviderProfileOutput
}

type CloudProviderProfileArgs struct {
	InfraNetworkProfile CloudProviderProfileInfraNetworkProfilePtrInput `pulumi:"infraNetworkProfile"`
	InfraStorageProfile CloudProviderProfileInfraStorageProfilePtrInput `pulumi:"infraStorageProfile"`
}

func (CloudProviderProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudProviderProfile)(nil)).Elem()
}

func (i CloudProviderProfileArgs) ToCloudProviderProfileOutput() CloudProviderProfileOutput {
	return i.ToCloudProviderProfileOutputWithContext(context.Background())
}

func (i CloudProviderProfileArgs) ToCloudProviderProfileOutputWithContext(ctx context.Context) CloudProviderProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProviderProfileOutput)
}

func (i CloudProviderProfileArgs) ToCloudProviderProfilePtrOutput() CloudProviderProfilePtrOutput {
	return i.ToCloudProviderProfilePtrOutputWithContext(context.Background())
}

func (i CloudProviderProfileArgs) ToCloudProviderProfilePtrOutputWithContext(ctx context.Context) CloudProviderProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProviderProfileOutput).ToCloudProviderProfilePtrOutputWithContext(ctx)
}









type CloudProviderProfilePtrInput interface {
	pulumi.Input

	ToCloudProviderProfilePtrOutput() CloudProviderProfilePtrOutput
	ToCloudProviderProfilePtrOutputWithContext(context.Context) CloudProviderProfilePtrOutput
}

type cloudProviderProfilePtrType CloudProviderProfileArgs

func CloudProviderProfilePtr(v *CloudProviderProfileArgs) CloudProviderProfilePtrInput {
	return (*cloudProviderProfilePtrType)(v)
}

func (*cloudProviderProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProviderProfile)(nil)).Elem()
}

func (i *cloudProviderProfilePtrType) ToCloudProviderProfilePtrOutput() CloudProviderProfilePtrOutput {
	return i.ToCloudProviderProfilePtrOutputWithContext(context.Background())
}

func (i *cloudProviderProfilePtrType) ToCloudProviderProfilePtrOutputWithContext(ctx context.Context) CloudProviderProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProviderProfilePtrOutput)
}

type CloudProviderProfileOutput struct{ *pulumi.OutputState }

func (CloudProviderProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudProviderProfile)(nil)).Elem()
}

func (o CloudProviderProfileOutput) ToCloudProviderProfileOutput() CloudProviderProfileOutput {
	return o
}

func (o CloudProviderProfileOutput) ToCloudProviderProfileOutputWithContext(ctx context.Context) CloudProviderProfileOutput {
	return o
}

func (o CloudProviderProfileOutput) ToCloudProviderProfilePtrOutput() CloudProviderProfilePtrOutput {
	return o.ToCloudProviderProfilePtrOutputWithContext(context.Background())
}

func (o CloudProviderProfileOutput) ToCloudProviderProfilePtrOutputWithContext(ctx context.Context) CloudProviderProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudProviderProfile) *CloudProviderProfile {
		return &v
	}).(CloudProviderProfilePtrOutput)
}

func (o CloudProviderProfileOutput) InfraNetworkProfile() CloudProviderProfileInfraNetworkProfilePtrOutput {
	return o.ApplyT(func(v CloudProviderProfile) *CloudProviderProfileInfraNetworkProfile { return v.InfraNetworkProfile }).(CloudProviderProfileInfraNetworkProfilePtrOutput)
}

func (o CloudProviderProfileOutput) InfraStorageProfile() CloudProviderProfileInfraStorageProfilePtrOutput {
	return o.ApplyT(func(v CloudProviderProfile) *CloudProviderProfileInfraStorageProfile { return v.InfraStorageProfile }).(CloudProviderProfileInfraStorageProfilePtrOutput)
}

type CloudProviderProfilePtrOutput struct{ *pulumi.OutputState }

func (CloudProviderProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProviderProfile)(nil)).Elem()
}

func (o CloudProviderProfilePtrOutput) ToCloudProviderProfilePtrOutput() CloudProviderProfilePtrOutput {
	return o
}

func (o CloudProviderProfilePtrOutput) ToCloudProviderProfilePtrOutputWithContext(ctx context.Context) CloudProviderProfilePtrOutput {
	return o
}

func (o CloudProviderProfilePtrOutput) Elem() CloudProviderProfileOutput {
	return o.ApplyT(func(v *CloudProviderProfile) CloudProviderProfile {
		if v != nil {
			return *v
		}
		var ret CloudProviderProfile
		return ret
	}).(CloudProviderProfileOutput)
}

func (o CloudProviderProfilePtrOutput) InfraNetworkProfile() CloudProviderProfileInfraNetworkProfilePtrOutput {
	return o.ApplyT(func(v *CloudProviderProfile) *CloudProviderProfileInfraNetworkProfile {
		if v == nil {
			return nil
		}
		return v.InfraNetworkProfile
	}).(CloudProviderProfileInfraNetworkProfilePtrOutput)
}

func (o CloudProviderProfilePtrOutput) InfraStorageProfile() CloudProviderProfileInfraStorageProfilePtrOutput {
	return o.ApplyT(func(v *CloudProviderProfile) *CloudProviderProfileInfraStorageProfile {
		if v == nil {
			return nil
		}
		return v.InfraStorageProfile
	}).(CloudProviderProfileInfraStorageProfilePtrOutput)
}

type CloudProviderProfileInfraNetworkProfile struct {
	VnetSubnetIds []string `pulumi:"vnetSubnetIds"`
}





type CloudProviderProfileInfraNetworkProfileInput interface {
	pulumi.Input

	ToCloudProviderProfileInfraNetworkProfileOutput() CloudProviderProfileInfraNetworkProfileOutput
	ToCloudProviderProfileInfraNetworkProfileOutputWithContext(context.Context) CloudProviderProfileInfraNetworkProfileOutput
}

type CloudProviderProfileInfraNetworkProfileArgs struct {
	VnetSubnetIds pulumi.StringArrayInput `pulumi:"vnetSubnetIds"`
}

func (CloudProviderProfileInfraNetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudProviderProfileInfraNetworkProfile)(nil)).Elem()
}

func (i CloudProviderProfileInfraNetworkProfileArgs) ToCloudProviderProfileInfraNetworkProfileOutput() CloudProviderProfileInfraNetworkProfileOutput {
	return i.ToCloudProviderProfileInfraNetworkProfileOutputWithContext(context.Background())
}

func (i CloudProviderProfileInfraNetworkProfileArgs) ToCloudProviderProfileInfraNetworkProfileOutputWithContext(ctx context.Context) CloudProviderProfileInfraNetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProviderProfileInfraNetworkProfileOutput)
}

func (i CloudProviderProfileInfraNetworkProfileArgs) ToCloudProviderProfileInfraNetworkProfilePtrOutput() CloudProviderProfileInfraNetworkProfilePtrOutput {
	return i.ToCloudProviderProfileInfraNetworkProfilePtrOutputWithContext(context.Background())
}

func (i CloudProviderProfileInfraNetworkProfileArgs) ToCloudProviderProfileInfraNetworkProfilePtrOutputWithContext(ctx context.Context) CloudProviderProfileInfraNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProviderProfileInfraNetworkProfileOutput).ToCloudProviderProfileInfraNetworkProfilePtrOutputWithContext(ctx)
}









type CloudProviderProfileInfraNetworkProfilePtrInput interface {
	pulumi.Input

	ToCloudProviderProfileInfraNetworkProfilePtrOutput() CloudProviderProfileInfraNetworkProfilePtrOutput
	ToCloudProviderProfileInfraNetworkProfilePtrOutputWithContext(context.Context) CloudProviderProfileInfraNetworkProfilePtrOutput
}

type cloudProviderProfileInfraNetworkProfilePtrType CloudProviderProfileInfraNetworkProfileArgs

func CloudProviderProfileInfraNetworkProfilePtr(v *CloudProviderProfileInfraNetworkProfileArgs) CloudProviderProfileInfraNetworkProfilePtrInput {
	return (*cloudProviderProfileInfraNetworkProfilePtrType)(v)
}

func (*cloudProviderProfileInfraNetworkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProviderProfileInfraNetworkProfile)(nil)).Elem()
}

func (i *cloudProviderProfileInfraNetworkProfilePtrType) ToCloudProviderProfileInfraNetworkProfilePtrOutput() CloudProviderProfileInfraNetworkProfilePtrOutput {
	return i.ToCloudProviderProfileInfraNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *cloudProviderProfileInfraNetworkProfilePtrType) ToCloudProviderProfileInfraNetworkProfilePtrOutputWithContext(ctx context.Context) CloudProviderProfileInfraNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProviderProfileInfraNetworkProfilePtrOutput)
}

type CloudProviderProfileInfraNetworkProfileOutput struct{ *pulumi.OutputState }

func (CloudProviderProfileInfraNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudProviderProfileInfraNetworkProfile)(nil)).Elem()
}

func (o CloudProviderProfileInfraNetworkProfileOutput) ToCloudProviderProfileInfraNetworkProfileOutput() CloudProviderProfileInfraNetworkProfileOutput {
	return o
}

func (o CloudProviderProfileInfraNetworkProfileOutput) ToCloudProviderProfileInfraNetworkProfileOutputWithContext(ctx context.Context) CloudProviderProfileInfraNetworkProfileOutput {
	return o
}

func (o CloudProviderProfileInfraNetworkProfileOutput) ToCloudProviderProfileInfraNetworkProfilePtrOutput() CloudProviderProfileInfraNetworkProfilePtrOutput {
	return o.ToCloudProviderProfileInfraNetworkProfilePtrOutputWithContext(context.Background())
}

func (o CloudProviderProfileInfraNetworkProfileOutput) ToCloudProviderProfileInfraNetworkProfilePtrOutputWithContext(ctx context.Context) CloudProviderProfileInfraNetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudProviderProfileInfraNetworkProfile) *CloudProviderProfileInfraNetworkProfile {
		return &v
	}).(CloudProviderProfileInfraNetworkProfilePtrOutput)
}

func (o CloudProviderProfileInfraNetworkProfileOutput) VnetSubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CloudProviderProfileInfraNetworkProfile) []string { return v.VnetSubnetIds }).(pulumi.StringArrayOutput)
}

type CloudProviderProfileInfraNetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (CloudProviderProfileInfraNetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProviderProfileInfraNetworkProfile)(nil)).Elem()
}

func (o CloudProviderProfileInfraNetworkProfilePtrOutput) ToCloudProviderProfileInfraNetworkProfilePtrOutput() CloudProviderProfileInfraNetworkProfilePtrOutput {
	return o
}

func (o CloudProviderProfileInfraNetworkProfilePtrOutput) ToCloudProviderProfileInfraNetworkProfilePtrOutputWithContext(ctx context.Context) CloudProviderProfileInfraNetworkProfilePtrOutput {
	return o
}

func (o CloudProviderProfileInfraNetworkProfilePtrOutput) Elem() CloudProviderProfileInfraNetworkProfileOutput {
	return o.ApplyT(func(v *CloudProviderProfileInfraNetworkProfile) CloudProviderProfileInfraNetworkProfile {
		if v != nil {
			return *v
		}
		var ret CloudProviderProfileInfraNetworkProfile
		return ret
	}).(CloudProviderProfileInfraNetworkProfileOutput)
}

func (o CloudProviderProfileInfraNetworkProfilePtrOutput) VnetSubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudProviderProfileInfraNetworkProfile) []string {
		if v == nil {
			return nil
		}
		return v.VnetSubnetIds
	}).(pulumi.StringArrayOutput)
}

type CloudProviderProfileInfraStorageProfile struct {
	StorageSpaceIds []string `pulumi:"storageSpaceIds"`
}





type CloudProviderProfileInfraStorageProfileInput interface {
	pulumi.Input

	ToCloudProviderProfileInfraStorageProfileOutput() CloudProviderProfileInfraStorageProfileOutput
	ToCloudProviderProfileInfraStorageProfileOutputWithContext(context.Context) CloudProviderProfileInfraStorageProfileOutput
}

type CloudProviderProfileInfraStorageProfileArgs struct {
	StorageSpaceIds pulumi.StringArrayInput `pulumi:"storageSpaceIds"`
}

func (CloudProviderProfileInfraStorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudProviderProfileInfraStorageProfile)(nil)).Elem()
}

func (i CloudProviderProfileInfraStorageProfileArgs) ToCloudProviderProfileInfraStorageProfileOutput() CloudProviderProfileInfraStorageProfileOutput {
	return i.ToCloudProviderProfileInfraStorageProfileOutputWithContext(context.Background())
}

func (i CloudProviderProfileInfraStorageProfileArgs) ToCloudProviderProfileInfraStorageProfileOutputWithContext(ctx context.Context) CloudProviderProfileInfraStorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProviderProfileInfraStorageProfileOutput)
}

func (i CloudProviderProfileInfraStorageProfileArgs) ToCloudProviderProfileInfraStorageProfilePtrOutput() CloudProviderProfileInfraStorageProfilePtrOutput {
	return i.ToCloudProviderProfileInfraStorageProfilePtrOutputWithContext(context.Background())
}

func (i CloudProviderProfileInfraStorageProfileArgs) ToCloudProviderProfileInfraStorageProfilePtrOutputWithContext(ctx context.Context) CloudProviderProfileInfraStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProviderProfileInfraStorageProfileOutput).ToCloudProviderProfileInfraStorageProfilePtrOutputWithContext(ctx)
}









type CloudProviderProfileInfraStorageProfilePtrInput interface {
	pulumi.Input

	ToCloudProviderProfileInfraStorageProfilePtrOutput() CloudProviderProfileInfraStorageProfilePtrOutput
	ToCloudProviderProfileInfraStorageProfilePtrOutputWithContext(context.Context) CloudProviderProfileInfraStorageProfilePtrOutput
}

type cloudProviderProfileInfraStorageProfilePtrType CloudProviderProfileInfraStorageProfileArgs

func CloudProviderProfileInfraStorageProfilePtr(v *CloudProviderProfileInfraStorageProfileArgs) CloudProviderProfileInfraStorageProfilePtrInput {
	return (*cloudProviderProfileInfraStorageProfilePtrType)(v)
}

func (*cloudProviderProfileInfraStorageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProviderProfileInfraStorageProfile)(nil)).Elem()
}

func (i *cloudProviderProfileInfraStorageProfilePtrType) ToCloudProviderProfileInfraStorageProfilePtrOutput() CloudProviderProfileInfraStorageProfilePtrOutput {
	return i.ToCloudProviderProfileInfraStorageProfilePtrOutputWithContext(context.Background())
}

func (i *cloudProviderProfileInfraStorageProfilePtrType) ToCloudProviderProfileInfraStorageProfilePtrOutputWithContext(ctx context.Context) CloudProviderProfileInfraStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProviderProfileInfraStorageProfilePtrOutput)
}

type CloudProviderProfileInfraStorageProfileOutput struct{ *pulumi.OutputState }

func (CloudProviderProfileInfraStorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudProviderProfileInfraStorageProfile)(nil)).Elem()
}

func (o CloudProviderProfileInfraStorageProfileOutput) ToCloudProviderProfileInfraStorageProfileOutput() CloudProviderProfileInfraStorageProfileOutput {
	return o
}

func (o CloudProviderProfileInfraStorageProfileOutput) ToCloudProviderProfileInfraStorageProfileOutputWithContext(ctx context.Context) CloudProviderProfileInfraStorageProfileOutput {
	return o
}

func (o CloudProviderProfileInfraStorageProfileOutput) ToCloudProviderProfileInfraStorageProfilePtrOutput() CloudProviderProfileInfraStorageProfilePtrOutput {
	return o.ToCloudProviderProfileInfraStorageProfilePtrOutputWithContext(context.Background())
}

func (o CloudProviderProfileInfraStorageProfileOutput) ToCloudProviderProfileInfraStorageProfilePtrOutputWithContext(ctx context.Context) CloudProviderProfileInfraStorageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudProviderProfileInfraStorageProfile) *CloudProviderProfileInfraStorageProfile {
		return &v
	}).(CloudProviderProfileInfraStorageProfilePtrOutput)
}

func (o CloudProviderProfileInfraStorageProfileOutput) StorageSpaceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CloudProviderProfileInfraStorageProfile) []string { return v.StorageSpaceIds }).(pulumi.StringArrayOutput)
}

type CloudProviderProfileInfraStorageProfilePtrOutput struct{ *pulumi.OutputState }

func (CloudProviderProfileInfraStorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProviderProfileInfraStorageProfile)(nil)).Elem()
}

func (o CloudProviderProfileInfraStorageProfilePtrOutput) ToCloudProviderProfileInfraStorageProfilePtrOutput() CloudProviderProfileInfraStorageProfilePtrOutput {
	return o
}

func (o CloudProviderProfileInfraStorageProfilePtrOutput) ToCloudProviderProfileInfraStorageProfilePtrOutputWithContext(ctx context.Context) CloudProviderProfileInfraStorageProfilePtrOutput {
	return o
}

func (o CloudProviderProfileInfraStorageProfilePtrOutput) Elem() CloudProviderProfileInfraStorageProfileOutput {
	return o.ApplyT(func(v *CloudProviderProfileInfraStorageProfile) CloudProviderProfileInfraStorageProfile {
		if v != nil {
			return *v
		}
		var ret CloudProviderProfileInfraStorageProfile
		return ret
	}).(CloudProviderProfileInfraStorageProfileOutput)
}

func (o CloudProviderProfileInfraStorageProfilePtrOutput) StorageSpaceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudProviderProfileInfraStorageProfile) []string {
		if v == nil {
			return nil
		}
		return v.StorageSpaceIds
	}).(pulumi.StringArrayOutput)
}

type CloudProviderProfileResponse struct {
	InfraNetworkProfile *CloudProviderProfileResponseInfraNetworkProfile `pulumi:"infraNetworkProfile"`
	InfraStorageProfile *CloudProviderProfileResponseInfraStorageProfile `pulumi:"infraStorageProfile"`
}

type CloudProviderProfileResponseOutput struct{ *pulumi.OutputState }

func (CloudProviderProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudProviderProfileResponse)(nil)).Elem()
}

func (o CloudProviderProfileResponseOutput) ToCloudProviderProfileResponseOutput() CloudProviderProfileResponseOutput {
	return o
}

func (o CloudProviderProfileResponseOutput) ToCloudProviderProfileResponseOutputWithContext(ctx context.Context) CloudProviderProfileResponseOutput {
	return o
}

func (o CloudProviderProfileResponseOutput) InfraNetworkProfile() CloudProviderProfileResponseInfraNetworkProfilePtrOutput {
	return o.ApplyT(func(v CloudProviderProfileResponse) *CloudProviderProfileResponseInfraNetworkProfile {
		return v.InfraNetworkProfile
	}).(CloudProviderProfileResponseInfraNetworkProfilePtrOutput)
}

func (o CloudProviderProfileResponseOutput) InfraStorageProfile() CloudProviderProfileResponseInfraStorageProfilePtrOutput {
	return o.ApplyT(func(v CloudProviderProfileResponse) *CloudProviderProfileResponseInfraStorageProfile {
		return v.InfraStorageProfile
	}).(CloudProviderProfileResponseInfraStorageProfilePtrOutput)
}

type CloudProviderProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudProviderProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProviderProfileResponse)(nil)).Elem()
}

func (o CloudProviderProfileResponsePtrOutput) ToCloudProviderProfileResponsePtrOutput() CloudProviderProfileResponsePtrOutput {
	return o
}

func (o CloudProviderProfileResponsePtrOutput) ToCloudProviderProfileResponsePtrOutputWithContext(ctx context.Context) CloudProviderProfileResponsePtrOutput {
	return o
}

func (o CloudProviderProfileResponsePtrOutput) Elem() CloudProviderProfileResponseOutput {
	return o.ApplyT(func(v *CloudProviderProfileResponse) CloudProviderProfileResponse {
		if v != nil {
			return *v
		}
		var ret CloudProviderProfileResponse
		return ret
	}).(CloudProviderProfileResponseOutput)
}

func (o CloudProviderProfileResponsePtrOutput) InfraNetworkProfile() CloudProviderProfileResponseInfraNetworkProfilePtrOutput {
	return o.ApplyT(func(v *CloudProviderProfileResponse) *CloudProviderProfileResponseInfraNetworkProfile {
		if v == nil {
			return nil
		}
		return v.InfraNetworkProfile
	}).(CloudProviderProfileResponseInfraNetworkProfilePtrOutput)
}

func (o CloudProviderProfileResponsePtrOutput) InfraStorageProfile() CloudProviderProfileResponseInfraStorageProfilePtrOutput {
	return o.ApplyT(func(v *CloudProviderProfileResponse) *CloudProviderProfileResponseInfraStorageProfile {
		if v == nil {
			return nil
		}
		return v.InfraStorageProfile
	}).(CloudProviderProfileResponseInfraStorageProfilePtrOutput)
}

type CloudProviderProfileResponseInfraNetworkProfile struct {
	VnetSubnetIds []string `pulumi:"vnetSubnetIds"`
}

type CloudProviderProfileResponseInfraNetworkProfileOutput struct{ *pulumi.OutputState }

func (CloudProviderProfileResponseInfraNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudProviderProfileResponseInfraNetworkProfile)(nil)).Elem()
}

func (o CloudProviderProfileResponseInfraNetworkProfileOutput) ToCloudProviderProfileResponseInfraNetworkProfileOutput() CloudProviderProfileResponseInfraNetworkProfileOutput {
	return o
}

func (o CloudProviderProfileResponseInfraNetworkProfileOutput) ToCloudProviderProfileResponseInfraNetworkProfileOutputWithContext(ctx context.Context) CloudProviderProfileResponseInfraNetworkProfileOutput {
	return o
}

func (o CloudProviderProfileResponseInfraNetworkProfileOutput) VnetSubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CloudProviderProfileResponseInfraNetworkProfile) []string { return v.VnetSubnetIds }).(pulumi.StringArrayOutput)
}

type CloudProviderProfileResponseInfraNetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (CloudProviderProfileResponseInfraNetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProviderProfileResponseInfraNetworkProfile)(nil)).Elem()
}

func (o CloudProviderProfileResponseInfraNetworkProfilePtrOutput) ToCloudProviderProfileResponseInfraNetworkProfilePtrOutput() CloudProviderProfileResponseInfraNetworkProfilePtrOutput {
	return o
}

func (o CloudProviderProfileResponseInfraNetworkProfilePtrOutput) ToCloudProviderProfileResponseInfraNetworkProfilePtrOutputWithContext(ctx context.Context) CloudProviderProfileResponseInfraNetworkProfilePtrOutput {
	return o
}

func (o CloudProviderProfileResponseInfraNetworkProfilePtrOutput) Elem() CloudProviderProfileResponseInfraNetworkProfileOutput {
	return o.ApplyT(func(v *CloudProviderProfileResponseInfraNetworkProfile) CloudProviderProfileResponseInfraNetworkProfile {
		if v != nil {
			return *v
		}
		var ret CloudProviderProfileResponseInfraNetworkProfile
		return ret
	}).(CloudProviderProfileResponseInfraNetworkProfileOutput)
}

func (o CloudProviderProfileResponseInfraNetworkProfilePtrOutput) VnetSubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudProviderProfileResponseInfraNetworkProfile) []string {
		if v == nil {
			return nil
		}
		return v.VnetSubnetIds
	}).(pulumi.StringArrayOutput)
}

type CloudProviderProfileResponseInfraStorageProfile struct {
	StorageSpaceIds []string `pulumi:"storageSpaceIds"`
}

type CloudProviderProfileResponseInfraStorageProfileOutput struct{ *pulumi.OutputState }

func (CloudProviderProfileResponseInfraStorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudProviderProfileResponseInfraStorageProfile)(nil)).Elem()
}

func (o CloudProviderProfileResponseInfraStorageProfileOutput) ToCloudProviderProfileResponseInfraStorageProfileOutput() CloudProviderProfileResponseInfraStorageProfileOutput {
	return o
}

func (o CloudProviderProfileResponseInfraStorageProfileOutput) ToCloudProviderProfileResponseInfraStorageProfileOutputWithContext(ctx context.Context) CloudProviderProfileResponseInfraStorageProfileOutput {
	return o
}

func (o CloudProviderProfileResponseInfraStorageProfileOutput) StorageSpaceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CloudProviderProfileResponseInfraStorageProfile) []string { return v.StorageSpaceIds }).(pulumi.StringArrayOutput)
}

type CloudProviderProfileResponseInfraStorageProfilePtrOutput struct{ *pulumi.OutputState }

func (CloudProviderProfileResponseInfraStorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProviderProfileResponseInfraStorageProfile)(nil)).Elem()
}

func (o CloudProviderProfileResponseInfraStorageProfilePtrOutput) ToCloudProviderProfileResponseInfraStorageProfilePtrOutput() CloudProviderProfileResponseInfraStorageProfilePtrOutput {
	return o
}

func (o CloudProviderProfileResponseInfraStorageProfilePtrOutput) ToCloudProviderProfileResponseInfraStorageProfilePtrOutputWithContext(ctx context.Context) CloudProviderProfileResponseInfraStorageProfilePtrOutput {
	return o
}

func (o CloudProviderProfileResponseInfraStorageProfilePtrOutput) Elem() CloudProviderProfileResponseInfraStorageProfileOutput {
	return o.ApplyT(func(v *CloudProviderProfileResponseInfraStorageProfile) CloudProviderProfileResponseInfraStorageProfile {
		if v != nil {
			return *v
		}
		var ret CloudProviderProfileResponseInfraStorageProfile
		return ret
	}).(CloudProviderProfileResponseInfraStorageProfileOutput)
}

func (o CloudProviderProfileResponseInfraStorageProfilePtrOutput) StorageSpaceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudProviderProfileResponseInfraStorageProfile) []string {
		if v == nil {
			return nil
		}
		return v.StorageSpaceIds
	}).(pulumi.StringArrayOutput)
}

type ControlPlaneEndpointProfileControlPlaneEndpoint struct {
	HostIP *string `pulumi:"hostIP"`
	Port   *string `pulumi:"port"`
}





type ControlPlaneEndpointProfileControlPlaneEndpointInput interface {
	pulumi.Input

	ToControlPlaneEndpointProfileControlPlaneEndpointOutput() ControlPlaneEndpointProfileControlPlaneEndpointOutput
	ToControlPlaneEndpointProfileControlPlaneEndpointOutputWithContext(context.Context) ControlPlaneEndpointProfileControlPlaneEndpointOutput
}

type ControlPlaneEndpointProfileControlPlaneEndpointArgs struct {
	HostIP pulumi.StringPtrInput `pulumi:"hostIP"`
	Port   pulumi.StringPtrInput `pulumi:"port"`
}

func (ControlPlaneEndpointProfileControlPlaneEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ControlPlaneEndpointProfileControlPlaneEndpoint)(nil)).Elem()
}

func (i ControlPlaneEndpointProfileControlPlaneEndpointArgs) ToControlPlaneEndpointProfileControlPlaneEndpointOutput() ControlPlaneEndpointProfileControlPlaneEndpointOutput {
	return i.ToControlPlaneEndpointProfileControlPlaneEndpointOutputWithContext(context.Background())
}

func (i ControlPlaneEndpointProfileControlPlaneEndpointArgs) ToControlPlaneEndpointProfileControlPlaneEndpointOutputWithContext(ctx context.Context) ControlPlaneEndpointProfileControlPlaneEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPlaneEndpointProfileControlPlaneEndpointOutput)
}

func (i ControlPlaneEndpointProfileControlPlaneEndpointArgs) ToControlPlaneEndpointProfileControlPlaneEndpointPtrOutput() ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput {
	return i.ToControlPlaneEndpointProfileControlPlaneEndpointPtrOutputWithContext(context.Background())
}

func (i ControlPlaneEndpointProfileControlPlaneEndpointArgs) ToControlPlaneEndpointProfileControlPlaneEndpointPtrOutputWithContext(ctx context.Context) ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPlaneEndpointProfileControlPlaneEndpointOutput).ToControlPlaneEndpointProfileControlPlaneEndpointPtrOutputWithContext(ctx)
}









type ControlPlaneEndpointProfileControlPlaneEndpointPtrInput interface {
	pulumi.Input

	ToControlPlaneEndpointProfileControlPlaneEndpointPtrOutput() ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput
	ToControlPlaneEndpointProfileControlPlaneEndpointPtrOutputWithContext(context.Context) ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput
}

type controlPlaneEndpointProfileControlPlaneEndpointPtrType ControlPlaneEndpointProfileControlPlaneEndpointArgs

func ControlPlaneEndpointProfileControlPlaneEndpointPtr(v *ControlPlaneEndpointProfileControlPlaneEndpointArgs) ControlPlaneEndpointProfileControlPlaneEndpointPtrInput {
	return (*controlPlaneEndpointProfileControlPlaneEndpointPtrType)(v)
}

func (*controlPlaneEndpointProfileControlPlaneEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ControlPlaneEndpointProfileControlPlaneEndpoint)(nil)).Elem()
}

func (i *controlPlaneEndpointProfileControlPlaneEndpointPtrType) ToControlPlaneEndpointProfileControlPlaneEndpointPtrOutput() ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput {
	return i.ToControlPlaneEndpointProfileControlPlaneEndpointPtrOutputWithContext(context.Background())
}

func (i *controlPlaneEndpointProfileControlPlaneEndpointPtrType) ToControlPlaneEndpointProfileControlPlaneEndpointPtrOutputWithContext(ctx context.Context) ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput)
}

type ControlPlaneEndpointProfileControlPlaneEndpointOutput struct{ *pulumi.OutputState }

func (ControlPlaneEndpointProfileControlPlaneEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ControlPlaneEndpointProfileControlPlaneEndpoint)(nil)).Elem()
}

func (o ControlPlaneEndpointProfileControlPlaneEndpointOutput) ToControlPlaneEndpointProfileControlPlaneEndpointOutput() ControlPlaneEndpointProfileControlPlaneEndpointOutput {
	return o
}

func (o ControlPlaneEndpointProfileControlPlaneEndpointOutput) ToControlPlaneEndpointProfileControlPlaneEndpointOutputWithContext(ctx context.Context) ControlPlaneEndpointProfileControlPlaneEndpointOutput {
	return o
}

func (o ControlPlaneEndpointProfileControlPlaneEndpointOutput) ToControlPlaneEndpointProfileControlPlaneEndpointPtrOutput() ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput {
	return o.ToControlPlaneEndpointProfileControlPlaneEndpointPtrOutputWithContext(context.Background())
}

func (o ControlPlaneEndpointProfileControlPlaneEndpointOutput) ToControlPlaneEndpointProfileControlPlaneEndpointPtrOutputWithContext(ctx context.Context) ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ControlPlaneEndpointProfileControlPlaneEndpoint) *ControlPlaneEndpointProfileControlPlaneEndpoint {
		return &v
	}).(ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput)
}

func (o ControlPlaneEndpointProfileControlPlaneEndpointOutput) HostIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ControlPlaneEndpointProfileControlPlaneEndpoint) *string { return v.HostIP }).(pulumi.StringPtrOutput)
}

func (o ControlPlaneEndpointProfileControlPlaneEndpointOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ControlPlaneEndpointProfileControlPlaneEndpoint) *string { return v.Port }).(pulumi.StringPtrOutput)
}

type ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput struct{ *pulumi.OutputState }

func (ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ControlPlaneEndpointProfileControlPlaneEndpoint)(nil)).Elem()
}

func (o ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput) ToControlPlaneEndpointProfileControlPlaneEndpointPtrOutput() ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput {
	return o
}

func (o ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput) ToControlPlaneEndpointProfileControlPlaneEndpointPtrOutputWithContext(ctx context.Context) ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput {
	return o
}

func (o ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput) Elem() ControlPlaneEndpointProfileControlPlaneEndpointOutput {
	return o.ApplyT(func(v *ControlPlaneEndpointProfileControlPlaneEndpoint) ControlPlaneEndpointProfileControlPlaneEndpoint {
		if v != nil {
			return *v
		}
		var ret ControlPlaneEndpointProfileControlPlaneEndpoint
		return ret
	}).(ControlPlaneEndpointProfileControlPlaneEndpointOutput)
}

func (o ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput) HostIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ControlPlaneEndpointProfileControlPlaneEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.HostIP
	}).(pulumi.StringPtrOutput)
}

func (o ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ControlPlaneEndpointProfileControlPlaneEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.StringPtrOutput)
}

type ControlPlaneEndpointProfileResponseControlPlaneEndpoint struct {
	HostIP *string `pulumi:"hostIP"`
	Port   *string `pulumi:"port"`
}

type ControlPlaneEndpointProfileResponseControlPlaneEndpointOutput struct{ *pulumi.OutputState }

func (ControlPlaneEndpointProfileResponseControlPlaneEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ControlPlaneEndpointProfileResponseControlPlaneEndpoint)(nil)).Elem()
}

func (o ControlPlaneEndpointProfileResponseControlPlaneEndpointOutput) ToControlPlaneEndpointProfileResponseControlPlaneEndpointOutput() ControlPlaneEndpointProfileResponseControlPlaneEndpointOutput {
	return o
}

func (o ControlPlaneEndpointProfileResponseControlPlaneEndpointOutput) ToControlPlaneEndpointProfileResponseControlPlaneEndpointOutputWithContext(ctx context.Context) ControlPlaneEndpointProfileResponseControlPlaneEndpointOutput {
	return o
}

func (o ControlPlaneEndpointProfileResponseControlPlaneEndpointOutput) HostIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ControlPlaneEndpointProfileResponseControlPlaneEndpoint) *string { return v.HostIP }).(pulumi.StringPtrOutput)
}

func (o ControlPlaneEndpointProfileResponseControlPlaneEndpointOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ControlPlaneEndpointProfileResponseControlPlaneEndpoint) *string { return v.Port }).(pulumi.StringPtrOutput)
}

type ControlPlaneEndpointProfileResponseControlPlaneEndpointPtrOutput struct{ *pulumi.OutputState }

func (ControlPlaneEndpointProfileResponseControlPlaneEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ControlPlaneEndpointProfileResponseControlPlaneEndpoint)(nil)).Elem()
}

func (o ControlPlaneEndpointProfileResponseControlPlaneEndpointPtrOutput) ToControlPlaneEndpointProfileResponseControlPlaneEndpointPtrOutput() ControlPlaneEndpointProfileResponseControlPlaneEndpointPtrOutput {
	return o
}

func (o ControlPlaneEndpointProfileResponseControlPlaneEndpointPtrOutput) ToControlPlaneEndpointProfileResponseControlPlaneEndpointPtrOutputWithContext(ctx context.Context) ControlPlaneEndpointProfileResponseControlPlaneEndpointPtrOutput {
	return o
}

func (o ControlPlaneEndpointProfileResponseControlPlaneEndpointPtrOutput) Elem() ControlPlaneEndpointProfileResponseControlPlaneEndpointOutput {
	return o.ApplyT(func(v *ControlPlaneEndpointProfileResponseControlPlaneEndpoint) ControlPlaneEndpointProfileResponseControlPlaneEndpoint {
		if v != nil {
			return *v
		}
		var ret ControlPlaneEndpointProfileResponseControlPlaneEndpoint
		return ret
	}).(ControlPlaneEndpointProfileResponseControlPlaneEndpointOutput)
}

func (o ControlPlaneEndpointProfileResponseControlPlaneEndpointPtrOutput) HostIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ControlPlaneEndpointProfileResponseControlPlaneEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.HostIP
	}).(pulumi.StringPtrOutput)
}

func (o ControlPlaneEndpointProfileResponseControlPlaneEndpointPtrOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ControlPlaneEndpointProfileResponseControlPlaneEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.StringPtrOutput)
}

type ControlPlaneProfile struct {
	AvailabilityZones    []string                                         `pulumi:"availabilityZones"`
	CloudProviderProfile *CloudProviderProfile                            `pulumi:"cloudProviderProfile"`
	ControlPlaneEndpoint *ControlPlaneEndpointProfileControlPlaneEndpoint `pulumi:"controlPlaneEndpoint"`
	Count                *int                                             `pulumi:"count"`
	LinuxProfile         *LinuxProfileProperties                          `pulumi:"linuxProfile"`
	MaxCount             *int                                             `pulumi:"maxCount"`
	MaxPods              *int                                             `pulumi:"maxPods"`
	MinCount             *int                                             `pulumi:"minCount"`
	Mode                 *string                                          `pulumi:"mode"`
	Name                 *string                                          `pulumi:"name"`
	NodeImageVersion     *string                                          `pulumi:"nodeImageVersion"`
	NodeLabels           map[string]string                                `pulumi:"nodeLabels"`
	NodeTaints           []string                                         `pulumi:"nodeTaints"`
	OsType               *string                                          `pulumi:"osType"`
	VmSize               *string                                          `pulumi:"vmSize"`
}


func (val *ControlPlaneProfile) Defaults() *ControlPlaneProfile {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		count_ := 1
		tmp.Count = &count_
	}
	if isZero(tmp.Mode) {
		mode_ := "User"
		tmp.Mode = &mode_
	}
	if isZero(tmp.OsType) {
		osType_ := "Linux"
		tmp.OsType = &osType_
	}
	return &tmp
}





type ControlPlaneProfileInput interface {
	pulumi.Input

	ToControlPlaneProfileOutput() ControlPlaneProfileOutput
	ToControlPlaneProfileOutputWithContext(context.Context) ControlPlaneProfileOutput
}

type ControlPlaneProfileArgs struct {
	AvailabilityZones    pulumi.StringArrayInput                                 `pulumi:"availabilityZones"`
	CloudProviderProfile CloudProviderProfilePtrInput                            `pulumi:"cloudProviderProfile"`
	ControlPlaneEndpoint ControlPlaneEndpointProfileControlPlaneEndpointPtrInput `pulumi:"controlPlaneEndpoint"`
	Count                pulumi.IntPtrInput                                      `pulumi:"count"`
	LinuxProfile         LinuxProfilePropertiesPtrInput                          `pulumi:"linuxProfile"`
	MaxCount             pulumi.IntPtrInput                                      `pulumi:"maxCount"`
	MaxPods              pulumi.IntPtrInput                                      `pulumi:"maxPods"`
	MinCount             pulumi.IntPtrInput                                      `pulumi:"minCount"`
	Mode                 pulumi.StringPtrInput                                   `pulumi:"mode"`
	Name                 pulumi.StringPtrInput                                   `pulumi:"name"`
	NodeImageVersion     pulumi.StringPtrInput                                   `pulumi:"nodeImageVersion"`
	NodeLabels           pulumi.StringMapInput                                   `pulumi:"nodeLabels"`
	NodeTaints           pulumi.StringArrayInput                                 `pulumi:"nodeTaints"`
	OsType               pulumi.StringPtrInput                                   `pulumi:"osType"`
	VmSize               pulumi.StringPtrInput                                   `pulumi:"vmSize"`
}


func (val *ControlPlaneProfileArgs) Defaults() *ControlPlaneProfileArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		tmp.Count = pulumi.IntPtr(1)
	}
	if isZero(tmp.Mode) {
		tmp.Mode = pulumi.StringPtr("User")
	}
	if isZero(tmp.OsType) {
		tmp.OsType = pulumi.StringPtr("Linux")
	}
	return &tmp
}
func (ControlPlaneProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ControlPlaneProfile)(nil)).Elem()
}

func (i ControlPlaneProfileArgs) ToControlPlaneProfileOutput() ControlPlaneProfileOutput {
	return i.ToControlPlaneProfileOutputWithContext(context.Background())
}

func (i ControlPlaneProfileArgs) ToControlPlaneProfileOutputWithContext(ctx context.Context) ControlPlaneProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPlaneProfileOutput)
}

func (i ControlPlaneProfileArgs) ToControlPlaneProfilePtrOutput() ControlPlaneProfilePtrOutput {
	return i.ToControlPlaneProfilePtrOutputWithContext(context.Background())
}

func (i ControlPlaneProfileArgs) ToControlPlaneProfilePtrOutputWithContext(ctx context.Context) ControlPlaneProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPlaneProfileOutput).ToControlPlaneProfilePtrOutputWithContext(ctx)
}









type ControlPlaneProfilePtrInput interface {
	pulumi.Input

	ToControlPlaneProfilePtrOutput() ControlPlaneProfilePtrOutput
	ToControlPlaneProfilePtrOutputWithContext(context.Context) ControlPlaneProfilePtrOutput
}

type controlPlaneProfilePtrType ControlPlaneProfileArgs

func ControlPlaneProfilePtr(v *ControlPlaneProfileArgs) ControlPlaneProfilePtrInput {
	return (*controlPlaneProfilePtrType)(v)
}

func (*controlPlaneProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ControlPlaneProfile)(nil)).Elem()
}

func (i *controlPlaneProfilePtrType) ToControlPlaneProfilePtrOutput() ControlPlaneProfilePtrOutput {
	return i.ToControlPlaneProfilePtrOutputWithContext(context.Background())
}

func (i *controlPlaneProfilePtrType) ToControlPlaneProfilePtrOutputWithContext(ctx context.Context) ControlPlaneProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPlaneProfilePtrOutput)
}

type ControlPlaneProfileOutput struct{ *pulumi.OutputState }

func (ControlPlaneProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ControlPlaneProfile)(nil)).Elem()
}

func (o ControlPlaneProfileOutput) ToControlPlaneProfileOutput() ControlPlaneProfileOutput {
	return o
}

func (o ControlPlaneProfileOutput) ToControlPlaneProfileOutputWithContext(ctx context.Context) ControlPlaneProfileOutput {
	return o
}

func (o ControlPlaneProfileOutput) ToControlPlaneProfilePtrOutput() ControlPlaneProfilePtrOutput {
	return o.ToControlPlaneProfilePtrOutputWithContext(context.Background())
}

func (o ControlPlaneProfileOutput) ToControlPlaneProfilePtrOutputWithContext(ctx context.Context) ControlPlaneProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ControlPlaneProfile) *ControlPlaneProfile {
		return &v
	}).(ControlPlaneProfilePtrOutput)
}

func (o ControlPlaneProfileOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ControlPlaneProfile) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o ControlPlaneProfileOutput) CloudProviderProfile() CloudProviderProfilePtrOutput {
	return o.ApplyT(func(v ControlPlaneProfile) *CloudProviderProfile { return v.CloudProviderProfile }).(CloudProviderProfilePtrOutput)
}

func (o ControlPlaneProfileOutput) ControlPlaneEndpoint() ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfile) *ControlPlaneEndpointProfileControlPlaneEndpoint {
		return v.ControlPlaneEndpoint
	}).(ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput)
}

func (o ControlPlaneProfileOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfile) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o ControlPlaneProfileOutput) LinuxProfile() LinuxProfilePropertiesPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfile) *LinuxProfileProperties { return v.LinuxProfile }).(LinuxProfilePropertiesPtrOutput)
}

func (o ControlPlaneProfileOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfile) *int { return v.MaxCount }).(pulumi.IntPtrOutput)
}

func (o ControlPlaneProfileOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfile) *int { return v.MaxPods }).(pulumi.IntPtrOutput)
}

func (o ControlPlaneProfileOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfile) *int { return v.MinCount }).(pulumi.IntPtrOutput)
}

func (o ControlPlaneProfileOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfile) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o ControlPlaneProfileOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfile) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ControlPlaneProfileOutput) NodeImageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfile) *string { return v.NodeImageVersion }).(pulumi.StringPtrOutput)
}

func (o ControlPlaneProfileOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v ControlPlaneProfile) map[string]string { return v.NodeLabels }).(pulumi.StringMapOutput)
}

func (o ControlPlaneProfileOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ControlPlaneProfile) []string { return v.NodeTaints }).(pulumi.StringArrayOutput)
}

func (o ControlPlaneProfileOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfile) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o ControlPlaneProfileOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfile) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type ControlPlaneProfilePtrOutput struct{ *pulumi.OutputState }

func (ControlPlaneProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ControlPlaneProfile)(nil)).Elem()
}

func (o ControlPlaneProfilePtrOutput) ToControlPlaneProfilePtrOutput() ControlPlaneProfilePtrOutput {
	return o
}

func (o ControlPlaneProfilePtrOutput) ToControlPlaneProfilePtrOutputWithContext(ctx context.Context) ControlPlaneProfilePtrOutput {
	return o
}

func (o ControlPlaneProfilePtrOutput) Elem() ControlPlaneProfileOutput {
	return o.ApplyT(func(v *ControlPlaneProfile) ControlPlaneProfile {
		if v != nil {
			return *v
		}
		var ret ControlPlaneProfile
		return ret
	}).(ControlPlaneProfileOutput)
}

func (o ControlPlaneProfilePtrOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ControlPlaneProfile) []string {
		if v == nil {
			return nil
		}
		return v.AvailabilityZones
	}).(pulumi.StringArrayOutput)
}

func (o ControlPlaneProfilePtrOutput) CloudProviderProfile() CloudProviderProfilePtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfile) *CloudProviderProfile {
		if v == nil {
			return nil
		}
		return v.CloudProviderProfile
	}).(CloudProviderProfilePtrOutput)
}

func (o ControlPlaneProfilePtrOutput) ControlPlaneEndpoint() ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfile) *ControlPlaneEndpointProfileControlPlaneEndpoint {
		if v == nil {
			return nil
		}
		return v.ControlPlaneEndpoint
	}).(ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput)
}

func (o ControlPlaneProfilePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfile) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o ControlPlaneProfilePtrOutput) LinuxProfile() LinuxProfilePropertiesPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfile) *LinuxProfileProperties {
		if v == nil {
			return nil
		}
		return v.LinuxProfile
	}).(LinuxProfilePropertiesPtrOutput)
}

func (o ControlPlaneProfilePtrOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfile) *int {
		if v == nil {
			return nil
		}
		return v.MaxCount
	}).(pulumi.IntPtrOutput)
}

func (o ControlPlaneProfilePtrOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfile) *int {
		if v == nil {
			return nil
		}
		return v.MaxPods
	}).(pulumi.IntPtrOutput)
}

func (o ControlPlaneProfilePtrOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfile) *int {
		if v == nil {
			return nil
		}
		return v.MinCount
	}).(pulumi.IntPtrOutput)
}

func (o ControlPlaneProfilePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfile) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

func (o ControlPlaneProfilePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfile) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ControlPlaneProfilePtrOutput) NodeImageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfile) *string {
		if v == nil {
			return nil
		}
		return v.NodeImageVersion
	}).(pulumi.StringPtrOutput)
}

func (o ControlPlaneProfilePtrOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ControlPlaneProfile) map[string]string {
		if v == nil {
			return nil
		}
		return v.NodeLabels
	}).(pulumi.StringMapOutput)
}

func (o ControlPlaneProfilePtrOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ControlPlaneProfile) []string {
		if v == nil {
			return nil
		}
		return v.NodeTaints
	}).(pulumi.StringArrayOutput)
}

func (o ControlPlaneProfilePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfile) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o ControlPlaneProfilePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfile) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type ControlPlaneProfileResponse struct {
	AvailabilityZones    []string                                                 `pulumi:"availabilityZones"`
	CloudProviderProfile *CloudProviderProfileResponse                            `pulumi:"cloudProviderProfile"`
	ControlPlaneEndpoint *ControlPlaneEndpointProfileResponseControlPlaneEndpoint `pulumi:"controlPlaneEndpoint"`
	Count                *int                                                     `pulumi:"count"`
	LinuxProfile         *LinuxProfilePropertiesResponse                          `pulumi:"linuxProfile"`
	MaxCount             *int                                                     `pulumi:"maxCount"`
	MaxPods              *int                                                     `pulumi:"maxPods"`
	MinCount             *int                                                     `pulumi:"minCount"`
	Mode                 *string                                                  `pulumi:"mode"`
	Name                 *string                                                  `pulumi:"name"`
	NodeImageVersion     *string                                                  `pulumi:"nodeImageVersion"`
	NodeLabels           map[string]string                                        `pulumi:"nodeLabels"`
	NodeTaints           []string                                                 `pulumi:"nodeTaints"`
	OsType               *string                                                  `pulumi:"osType"`
	VmSize               *string                                                  `pulumi:"vmSize"`
}


func (val *ControlPlaneProfileResponse) Defaults() *ControlPlaneProfileResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		count_ := 1
		tmp.Count = &count_
	}
	if isZero(tmp.Mode) {
		mode_ := "User"
		tmp.Mode = &mode_
	}
	if isZero(tmp.OsType) {
		osType_ := "Linux"
		tmp.OsType = &osType_
	}
	return &tmp
}

type ControlPlaneProfileResponseOutput struct{ *pulumi.OutputState }

func (ControlPlaneProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ControlPlaneProfileResponse)(nil)).Elem()
}

func (o ControlPlaneProfileResponseOutput) ToControlPlaneProfileResponseOutput() ControlPlaneProfileResponseOutput {
	return o
}

func (o ControlPlaneProfileResponseOutput) ToControlPlaneProfileResponseOutputWithContext(ctx context.Context) ControlPlaneProfileResponseOutput {
	return o
}

func (o ControlPlaneProfileResponseOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ControlPlaneProfileResponse) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o ControlPlaneProfileResponseOutput) CloudProviderProfile() CloudProviderProfileResponsePtrOutput {
	return o.ApplyT(func(v ControlPlaneProfileResponse) *CloudProviderProfileResponse { return v.CloudProviderProfile }).(CloudProviderProfileResponsePtrOutput)
}

func (o ControlPlaneProfileResponseOutput) ControlPlaneEndpoint() ControlPlaneEndpointProfileResponseControlPlaneEndpointPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfileResponse) *ControlPlaneEndpointProfileResponseControlPlaneEndpoint {
		return v.ControlPlaneEndpoint
	}).(ControlPlaneEndpointProfileResponseControlPlaneEndpointPtrOutput)
}

func (o ControlPlaneProfileResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfileResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o ControlPlaneProfileResponseOutput) LinuxProfile() LinuxProfilePropertiesResponsePtrOutput {
	return o.ApplyT(func(v ControlPlaneProfileResponse) *LinuxProfilePropertiesResponse { return v.LinuxProfile }).(LinuxProfilePropertiesResponsePtrOutput)
}

func (o ControlPlaneProfileResponseOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfileResponse) *int { return v.MaxCount }).(pulumi.IntPtrOutput)
}

func (o ControlPlaneProfileResponseOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfileResponse) *int { return v.MaxPods }).(pulumi.IntPtrOutput)
}

func (o ControlPlaneProfileResponseOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfileResponse) *int { return v.MinCount }).(pulumi.IntPtrOutput)
}

func (o ControlPlaneProfileResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfileResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o ControlPlaneProfileResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfileResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ControlPlaneProfileResponseOutput) NodeImageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfileResponse) *string { return v.NodeImageVersion }).(pulumi.StringPtrOutput)
}

func (o ControlPlaneProfileResponseOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v ControlPlaneProfileResponse) map[string]string { return v.NodeLabels }).(pulumi.StringMapOutput)
}

func (o ControlPlaneProfileResponseOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ControlPlaneProfileResponse) []string { return v.NodeTaints }).(pulumi.StringArrayOutput)
}

func (o ControlPlaneProfileResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfileResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o ControlPlaneProfileResponseOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ControlPlaneProfileResponse) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type ControlPlaneProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ControlPlaneProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ControlPlaneProfileResponse)(nil)).Elem()
}

func (o ControlPlaneProfileResponsePtrOutput) ToControlPlaneProfileResponsePtrOutput() ControlPlaneProfileResponsePtrOutput {
	return o
}

func (o ControlPlaneProfileResponsePtrOutput) ToControlPlaneProfileResponsePtrOutputWithContext(ctx context.Context) ControlPlaneProfileResponsePtrOutput {
	return o
}

func (o ControlPlaneProfileResponsePtrOutput) Elem() ControlPlaneProfileResponseOutput {
	return o.ApplyT(func(v *ControlPlaneProfileResponse) ControlPlaneProfileResponse {
		if v != nil {
			return *v
		}
		var ret ControlPlaneProfileResponse
		return ret
	}).(ControlPlaneProfileResponseOutput)
}

func (o ControlPlaneProfileResponsePtrOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ControlPlaneProfileResponse) []string {
		if v == nil {
			return nil
		}
		return v.AvailabilityZones
	}).(pulumi.StringArrayOutput)
}

func (o ControlPlaneProfileResponsePtrOutput) CloudProviderProfile() CloudProviderProfileResponsePtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfileResponse) *CloudProviderProfileResponse {
		if v == nil {
			return nil
		}
		return v.CloudProviderProfile
	}).(CloudProviderProfileResponsePtrOutput)
}

func (o ControlPlaneProfileResponsePtrOutput) ControlPlaneEndpoint() ControlPlaneEndpointProfileResponseControlPlaneEndpointPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfileResponse) *ControlPlaneEndpointProfileResponseControlPlaneEndpoint {
		if v == nil {
			return nil
		}
		return v.ControlPlaneEndpoint
	}).(ControlPlaneEndpointProfileResponseControlPlaneEndpointPtrOutput)
}

func (o ControlPlaneProfileResponsePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o ControlPlaneProfileResponsePtrOutput) LinuxProfile() LinuxProfilePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfileResponse) *LinuxProfilePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.LinuxProfile
	}).(LinuxProfilePropertiesResponsePtrOutput)
}

func (o ControlPlaneProfileResponsePtrOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxCount
	}).(pulumi.IntPtrOutput)
}

func (o ControlPlaneProfileResponsePtrOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxPods
	}).(pulumi.IntPtrOutput)
}

func (o ControlPlaneProfileResponsePtrOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.MinCount
	}).(pulumi.IntPtrOutput)
}

func (o ControlPlaneProfileResponsePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

func (o ControlPlaneProfileResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ControlPlaneProfileResponsePtrOutput) NodeImageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.NodeImageVersion
	}).(pulumi.StringPtrOutput)
}

func (o ControlPlaneProfileResponsePtrOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ControlPlaneProfileResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.NodeLabels
	}).(pulumi.StringMapOutput)
}

func (o ControlPlaneProfileResponsePtrOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ControlPlaneProfileResponse) []string {
		if v == nil {
			return nil
		}
		return v.NodeTaints
	}).(pulumi.StringArrayOutput)
}

func (o ControlPlaneProfileResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o ControlPlaneProfileResponsePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ControlPlaneProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type HttpProxyConfig struct {
	HttpProxy  *string  `pulumi:"httpProxy"`
	HttpsProxy *string  `pulumi:"httpsProxy"`
	NoProxy    []string `pulumi:"noProxy"`
	Password   *string  `pulumi:"password"`
	TrustedCa  *string  `pulumi:"trustedCa"`
	Username   *string  `pulumi:"username"`
}





type HttpProxyConfigInput interface {
	pulumi.Input

	ToHttpProxyConfigOutput() HttpProxyConfigOutput
	ToHttpProxyConfigOutputWithContext(context.Context) HttpProxyConfigOutput
}

type HttpProxyConfigArgs struct {
	HttpProxy  pulumi.StringPtrInput   `pulumi:"httpProxy"`
	HttpsProxy pulumi.StringPtrInput   `pulumi:"httpsProxy"`
	NoProxy    pulumi.StringArrayInput `pulumi:"noProxy"`
	Password   pulumi.StringPtrInput   `pulumi:"password"`
	TrustedCa  pulumi.StringPtrInput   `pulumi:"trustedCa"`
	Username   pulumi.StringPtrInput   `pulumi:"username"`
}

func (HttpProxyConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpProxyConfig)(nil)).Elem()
}

func (i HttpProxyConfigArgs) ToHttpProxyConfigOutput() HttpProxyConfigOutput {
	return i.ToHttpProxyConfigOutputWithContext(context.Background())
}

func (i HttpProxyConfigArgs) ToHttpProxyConfigOutputWithContext(ctx context.Context) HttpProxyConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpProxyConfigOutput)
}

func (i HttpProxyConfigArgs) ToHttpProxyConfigPtrOutput() HttpProxyConfigPtrOutput {
	return i.ToHttpProxyConfigPtrOutputWithContext(context.Background())
}

func (i HttpProxyConfigArgs) ToHttpProxyConfigPtrOutputWithContext(ctx context.Context) HttpProxyConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpProxyConfigOutput).ToHttpProxyConfigPtrOutputWithContext(ctx)
}









type HttpProxyConfigPtrInput interface {
	pulumi.Input

	ToHttpProxyConfigPtrOutput() HttpProxyConfigPtrOutput
	ToHttpProxyConfigPtrOutputWithContext(context.Context) HttpProxyConfigPtrOutput
}

type httpProxyConfigPtrType HttpProxyConfigArgs

func HttpProxyConfigPtr(v *HttpProxyConfigArgs) HttpProxyConfigPtrInput {
	return (*httpProxyConfigPtrType)(v)
}

func (*httpProxyConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpProxyConfig)(nil)).Elem()
}

func (i *httpProxyConfigPtrType) ToHttpProxyConfigPtrOutput() HttpProxyConfigPtrOutput {
	return i.ToHttpProxyConfigPtrOutputWithContext(context.Background())
}

func (i *httpProxyConfigPtrType) ToHttpProxyConfigPtrOutputWithContext(ctx context.Context) HttpProxyConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpProxyConfigPtrOutput)
}

type HttpProxyConfigOutput struct{ *pulumi.OutputState }

func (HttpProxyConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpProxyConfig)(nil)).Elem()
}

func (o HttpProxyConfigOutput) ToHttpProxyConfigOutput() HttpProxyConfigOutput {
	return o
}

func (o HttpProxyConfigOutput) ToHttpProxyConfigOutputWithContext(ctx context.Context) HttpProxyConfigOutput {
	return o
}

func (o HttpProxyConfigOutput) ToHttpProxyConfigPtrOutput() HttpProxyConfigPtrOutput {
	return o.ToHttpProxyConfigPtrOutputWithContext(context.Background())
}

func (o HttpProxyConfigOutput) ToHttpProxyConfigPtrOutputWithContext(ctx context.Context) HttpProxyConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpProxyConfig) *HttpProxyConfig {
		return &v
	}).(HttpProxyConfigPtrOutput)
}

func (o HttpProxyConfigOutput) HttpProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpProxyConfig) *string { return v.HttpProxy }).(pulumi.StringPtrOutput)
}

func (o HttpProxyConfigOutput) HttpsProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpProxyConfig) *string { return v.HttpsProxy }).(pulumi.StringPtrOutput)
}

func (o HttpProxyConfigOutput) NoProxy() pulumi.StringArrayOutput {
	return o.ApplyT(func(v HttpProxyConfig) []string { return v.NoProxy }).(pulumi.StringArrayOutput)
}

func (o HttpProxyConfigOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpProxyConfig) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o HttpProxyConfigOutput) TrustedCa() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpProxyConfig) *string { return v.TrustedCa }).(pulumi.StringPtrOutput)
}

func (o HttpProxyConfigOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpProxyConfig) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type HttpProxyConfigPtrOutput struct{ *pulumi.OutputState }

func (HttpProxyConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpProxyConfig)(nil)).Elem()
}

func (o HttpProxyConfigPtrOutput) ToHttpProxyConfigPtrOutput() HttpProxyConfigPtrOutput {
	return o
}

func (o HttpProxyConfigPtrOutput) ToHttpProxyConfigPtrOutputWithContext(ctx context.Context) HttpProxyConfigPtrOutput {
	return o
}

func (o HttpProxyConfigPtrOutput) Elem() HttpProxyConfigOutput {
	return o.ApplyT(func(v *HttpProxyConfig) HttpProxyConfig {
		if v != nil {
			return *v
		}
		var ret HttpProxyConfig
		return ret
	}).(HttpProxyConfigOutput)
}

func (o HttpProxyConfigPtrOutput) HttpProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpProxyConfig) *string {
		if v == nil {
			return nil
		}
		return v.HttpProxy
	}).(pulumi.StringPtrOutput)
}

func (o HttpProxyConfigPtrOutput) HttpsProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpProxyConfig) *string {
		if v == nil {
			return nil
		}
		return v.HttpsProxy
	}).(pulumi.StringPtrOutput)
}

func (o HttpProxyConfigPtrOutput) NoProxy() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HttpProxyConfig) []string {
		if v == nil {
			return nil
		}
		return v.NoProxy
	}).(pulumi.StringArrayOutput)
}

func (o HttpProxyConfigPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpProxyConfig) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o HttpProxyConfigPtrOutput) TrustedCa() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpProxyConfig) *string {
		if v == nil {
			return nil
		}
		return v.TrustedCa
	}).(pulumi.StringPtrOutput)
}

func (o HttpProxyConfigPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpProxyConfig) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type HttpProxyConfigResponseResponse struct {
	HttpProxy  *string  `pulumi:"httpProxy"`
	HttpsProxy *string  `pulumi:"httpsProxy"`
	NoProxy    []string `pulumi:"noProxy"`
	TrustedCa  *string  `pulumi:"trustedCa"`
	Username   *string  `pulumi:"username"`
}

type HttpProxyConfigResponseResponseOutput struct{ *pulumi.OutputState }

func (HttpProxyConfigResponseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpProxyConfigResponseResponse)(nil)).Elem()
}

func (o HttpProxyConfigResponseResponseOutput) ToHttpProxyConfigResponseResponseOutput() HttpProxyConfigResponseResponseOutput {
	return o
}

func (o HttpProxyConfigResponseResponseOutput) ToHttpProxyConfigResponseResponseOutputWithContext(ctx context.Context) HttpProxyConfigResponseResponseOutput {
	return o
}

func (o HttpProxyConfigResponseResponseOutput) HttpProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpProxyConfigResponseResponse) *string { return v.HttpProxy }).(pulumi.StringPtrOutput)
}

func (o HttpProxyConfigResponseResponseOutput) HttpsProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpProxyConfigResponseResponse) *string { return v.HttpsProxy }).(pulumi.StringPtrOutput)
}

func (o HttpProxyConfigResponseResponseOutput) NoProxy() pulumi.StringArrayOutput {
	return o.ApplyT(func(v HttpProxyConfigResponseResponse) []string { return v.NoProxy }).(pulumi.StringArrayOutput)
}

func (o HttpProxyConfigResponseResponseOutput) TrustedCa() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpProxyConfigResponseResponse) *string { return v.TrustedCa }).(pulumi.StringPtrOutput)
}

func (o HttpProxyConfigResponseResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpProxyConfigResponseResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type HttpProxyConfigResponseResponsePtrOutput struct{ *pulumi.OutputState }

func (HttpProxyConfigResponseResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpProxyConfigResponseResponse)(nil)).Elem()
}

func (o HttpProxyConfigResponseResponsePtrOutput) ToHttpProxyConfigResponseResponsePtrOutput() HttpProxyConfigResponseResponsePtrOutput {
	return o
}

func (o HttpProxyConfigResponseResponsePtrOutput) ToHttpProxyConfigResponseResponsePtrOutputWithContext(ctx context.Context) HttpProxyConfigResponseResponsePtrOutput {
	return o
}

func (o HttpProxyConfigResponseResponsePtrOutput) Elem() HttpProxyConfigResponseResponseOutput {
	return o.ApplyT(func(v *HttpProxyConfigResponseResponse) HttpProxyConfigResponseResponse {
		if v != nil {
			return *v
		}
		var ret HttpProxyConfigResponseResponse
		return ret
	}).(HttpProxyConfigResponseResponseOutput)
}

func (o HttpProxyConfigResponseResponsePtrOutput) HttpProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpProxyConfigResponseResponse) *string {
		if v == nil {
			return nil
		}
		return v.HttpProxy
	}).(pulumi.StringPtrOutput)
}

func (o HttpProxyConfigResponseResponsePtrOutput) HttpsProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpProxyConfigResponseResponse) *string {
		if v == nil {
			return nil
		}
		return v.HttpsProxy
	}).(pulumi.StringPtrOutput)
}

func (o HttpProxyConfigResponseResponsePtrOutput) NoProxy() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HttpProxyConfigResponseResponse) []string {
		if v == nil {
			return nil
		}
		return v.NoProxy
	}).(pulumi.StringArrayOutput)
}

func (o HttpProxyConfigResponseResponsePtrOutput) TrustedCa() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpProxyConfigResponseResponse) *string {
		if v == nil {
			return nil
		}
		return v.TrustedCa
	}).(pulumi.StringPtrOutput)
}

func (o HttpProxyConfigResponseResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpProxyConfigResponseResponse) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type LinuxProfileProperties struct {
	AdminUsername *string                    `pulumi:"adminUsername"`
	Ssh           *LinuxProfilePropertiesSsh `pulumi:"ssh"`
}





type LinuxProfilePropertiesInput interface {
	pulumi.Input

	ToLinuxProfilePropertiesOutput() LinuxProfilePropertiesOutput
	ToLinuxProfilePropertiesOutputWithContext(context.Context) LinuxProfilePropertiesOutput
}

type LinuxProfilePropertiesArgs struct {
	AdminUsername pulumi.StringPtrInput             `pulumi:"adminUsername"`
	Ssh           LinuxProfilePropertiesSshPtrInput `pulumi:"ssh"`
}

func (LinuxProfilePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxProfileProperties)(nil)).Elem()
}

func (i LinuxProfilePropertiesArgs) ToLinuxProfilePropertiesOutput() LinuxProfilePropertiesOutput {
	return i.ToLinuxProfilePropertiesOutputWithContext(context.Background())
}

func (i LinuxProfilePropertiesArgs) ToLinuxProfilePropertiesOutputWithContext(ctx context.Context) LinuxProfilePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxProfilePropertiesOutput)
}

func (i LinuxProfilePropertiesArgs) ToLinuxProfilePropertiesPtrOutput() LinuxProfilePropertiesPtrOutput {
	return i.ToLinuxProfilePropertiesPtrOutputWithContext(context.Background())
}

func (i LinuxProfilePropertiesArgs) ToLinuxProfilePropertiesPtrOutputWithContext(ctx context.Context) LinuxProfilePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxProfilePropertiesOutput).ToLinuxProfilePropertiesPtrOutputWithContext(ctx)
}









type LinuxProfilePropertiesPtrInput interface {
	pulumi.Input

	ToLinuxProfilePropertiesPtrOutput() LinuxProfilePropertiesPtrOutput
	ToLinuxProfilePropertiesPtrOutputWithContext(context.Context) LinuxProfilePropertiesPtrOutput
}

type linuxProfilePropertiesPtrType LinuxProfilePropertiesArgs

func LinuxProfilePropertiesPtr(v *LinuxProfilePropertiesArgs) LinuxProfilePropertiesPtrInput {
	return (*linuxProfilePropertiesPtrType)(v)
}

func (*linuxProfilePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxProfileProperties)(nil)).Elem()
}

func (i *linuxProfilePropertiesPtrType) ToLinuxProfilePropertiesPtrOutput() LinuxProfilePropertiesPtrOutput {
	return i.ToLinuxProfilePropertiesPtrOutputWithContext(context.Background())
}

func (i *linuxProfilePropertiesPtrType) ToLinuxProfilePropertiesPtrOutputWithContext(ctx context.Context) LinuxProfilePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxProfilePropertiesPtrOutput)
}

type LinuxProfilePropertiesOutput struct{ *pulumi.OutputState }

func (LinuxProfilePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxProfileProperties)(nil)).Elem()
}

func (o LinuxProfilePropertiesOutput) ToLinuxProfilePropertiesOutput() LinuxProfilePropertiesOutput {
	return o
}

func (o LinuxProfilePropertiesOutput) ToLinuxProfilePropertiesOutputWithContext(ctx context.Context) LinuxProfilePropertiesOutput {
	return o
}

func (o LinuxProfilePropertiesOutput) ToLinuxProfilePropertiesPtrOutput() LinuxProfilePropertiesPtrOutput {
	return o.ToLinuxProfilePropertiesPtrOutputWithContext(context.Background())
}

func (o LinuxProfilePropertiesOutput) ToLinuxProfilePropertiesPtrOutputWithContext(ctx context.Context) LinuxProfilePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinuxProfileProperties) *LinuxProfileProperties {
		return &v
	}).(LinuxProfilePropertiesPtrOutput)
}

func (o LinuxProfilePropertiesOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxProfileProperties) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o LinuxProfilePropertiesOutput) Ssh() LinuxProfilePropertiesSshPtrOutput {
	return o.ApplyT(func(v LinuxProfileProperties) *LinuxProfilePropertiesSsh { return v.Ssh }).(LinuxProfilePropertiesSshPtrOutput)
}

type LinuxProfilePropertiesPtrOutput struct{ *pulumi.OutputState }

func (LinuxProfilePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxProfileProperties)(nil)).Elem()
}

func (o LinuxProfilePropertiesPtrOutput) ToLinuxProfilePropertiesPtrOutput() LinuxProfilePropertiesPtrOutput {
	return o
}

func (o LinuxProfilePropertiesPtrOutput) ToLinuxProfilePropertiesPtrOutputWithContext(ctx context.Context) LinuxProfilePropertiesPtrOutput {
	return o
}

func (o LinuxProfilePropertiesPtrOutput) Elem() LinuxProfilePropertiesOutput {
	return o.ApplyT(func(v *LinuxProfileProperties) LinuxProfileProperties {
		if v != nil {
			return *v
		}
		var ret LinuxProfileProperties
		return ret
	}).(LinuxProfilePropertiesOutput)
}

func (o LinuxProfilePropertiesPtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxProfileProperties) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o LinuxProfilePropertiesPtrOutput) Ssh() LinuxProfilePropertiesSshPtrOutput {
	return o.ApplyT(func(v *LinuxProfileProperties) *LinuxProfilePropertiesSsh {
		if v == nil {
			return nil
		}
		return v.Ssh
	}).(LinuxProfilePropertiesSshPtrOutput)
}

type LinuxProfilePropertiesPublicKeys struct {
	KeyData *string `pulumi:"keyData"`
}





type LinuxProfilePropertiesPublicKeysInput interface {
	pulumi.Input

	ToLinuxProfilePropertiesPublicKeysOutput() LinuxProfilePropertiesPublicKeysOutput
	ToLinuxProfilePropertiesPublicKeysOutputWithContext(context.Context) LinuxProfilePropertiesPublicKeysOutput
}

type LinuxProfilePropertiesPublicKeysArgs struct {
	KeyData pulumi.StringPtrInput `pulumi:"keyData"`
}

func (LinuxProfilePropertiesPublicKeysArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxProfilePropertiesPublicKeys)(nil)).Elem()
}

func (i LinuxProfilePropertiesPublicKeysArgs) ToLinuxProfilePropertiesPublicKeysOutput() LinuxProfilePropertiesPublicKeysOutput {
	return i.ToLinuxProfilePropertiesPublicKeysOutputWithContext(context.Background())
}

func (i LinuxProfilePropertiesPublicKeysArgs) ToLinuxProfilePropertiesPublicKeysOutputWithContext(ctx context.Context) LinuxProfilePropertiesPublicKeysOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxProfilePropertiesPublicKeysOutput)
}





type LinuxProfilePropertiesPublicKeysArrayInput interface {
	pulumi.Input

	ToLinuxProfilePropertiesPublicKeysArrayOutput() LinuxProfilePropertiesPublicKeysArrayOutput
	ToLinuxProfilePropertiesPublicKeysArrayOutputWithContext(context.Context) LinuxProfilePropertiesPublicKeysArrayOutput
}

type LinuxProfilePropertiesPublicKeysArray []LinuxProfilePropertiesPublicKeysInput

func (LinuxProfilePropertiesPublicKeysArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinuxProfilePropertiesPublicKeys)(nil)).Elem()
}

func (i LinuxProfilePropertiesPublicKeysArray) ToLinuxProfilePropertiesPublicKeysArrayOutput() LinuxProfilePropertiesPublicKeysArrayOutput {
	return i.ToLinuxProfilePropertiesPublicKeysArrayOutputWithContext(context.Background())
}

func (i LinuxProfilePropertiesPublicKeysArray) ToLinuxProfilePropertiesPublicKeysArrayOutputWithContext(ctx context.Context) LinuxProfilePropertiesPublicKeysArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxProfilePropertiesPublicKeysArrayOutput)
}

type LinuxProfilePropertiesPublicKeysOutput struct{ *pulumi.OutputState }

func (LinuxProfilePropertiesPublicKeysOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxProfilePropertiesPublicKeys)(nil)).Elem()
}

func (o LinuxProfilePropertiesPublicKeysOutput) ToLinuxProfilePropertiesPublicKeysOutput() LinuxProfilePropertiesPublicKeysOutput {
	return o
}

func (o LinuxProfilePropertiesPublicKeysOutput) ToLinuxProfilePropertiesPublicKeysOutputWithContext(ctx context.Context) LinuxProfilePropertiesPublicKeysOutput {
	return o
}

func (o LinuxProfilePropertiesPublicKeysOutput) KeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxProfilePropertiesPublicKeys) *string { return v.KeyData }).(pulumi.StringPtrOutput)
}

type LinuxProfilePropertiesPublicKeysArrayOutput struct{ *pulumi.OutputState }

func (LinuxProfilePropertiesPublicKeysArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinuxProfilePropertiesPublicKeys)(nil)).Elem()
}

func (o LinuxProfilePropertiesPublicKeysArrayOutput) ToLinuxProfilePropertiesPublicKeysArrayOutput() LinuxProfilePropertiesPublicKeysArrayOutput {
	return o
}

func (o LinuxProfilePropertiesPublicKeysArrayOutput) ToLinuxProfilePropertiesPublicKeysArrayOutputWithContext(ctx context.Context) LinuxProfilePropertiesPublicKeysArrayOutput {
	return o
}

func (o LinuxProfilePropertiesPublicKeysArrayOutput) Index(i pulumi.IntInput) LinuxProfilePropertiesPublicKeysOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinuxProfilePropertiesPublicKeys {
		return vs[0].([]LinuxProfilePropertiesPublicKeys)[vs[1].(int)]
	}).(LinuxProfilePropertiesPublicKeysOutput)
}

type LinuxProfilePropertiesResponse struct {
	AdminUsername *string                            `pulumi:"adminUsername"`
	Ssh           *LinuxProfilePropertiesResponseSsh `pulumi:"ssh"`
}

type LinuxProfilePropertiesResponseOutput struct{ *pulumi.OutputState }

func (LinuxProfilePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxProfilePropertiesResponse)(nil)).Elem()
}

func (o LinuxProfilePropertiesResponseOutput) ToLinuxProfilePropertiesResponseOutput() LinuxProfilePropertiesResponseOutput {
	return o
}

func (o LinuxProfilePropertiesResponseOutput) ToLinuxProfilePropertiesResponseOutputWithContext(ctx context.Context) LinuxProfilePropertiesResponseOutput {
	return o
}

func (o LinuxProfilePropertiesResponseOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxProfilePropertiesResponse) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o LinuxProfilePropertiesResponseOutput) Ssh() LinuxProfilePropertiesResponseSshPtrOutput {
	return o.ApplyT(func(v LinuxProfilePropertiesResponse) *LinuxProfilePropertiesResponseSsh { return v.Ssh }).(LinuxProfilePropertiesResponseSshPtrOutput)
}

type LinuxProfilePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (LinuxProfilePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxProfilePropertiesResponse)(nil)).Elem()
}

func (o LinuxProfilePropertiesResponsePtrOutput) ToLinuxProfilePropertiesResponsePtrOutput() LinuxProfilePropertiesResponsePtrOutput {
	return o
}

func (o LinuxProfilePropertiesResponsePtrOutput) ToLinuxProfilePropertiesResponsePtrOutputWithContext(ctx context.Context) LinuxProfilePropertiesResponsePtrOutput {
	return o
}

func (o LinuxProfilePropertiesResponsePtrOutput) Elem() LinuxProfilePropertiesResponseOutput {
	return o.ApplyT(func(v *LinuxProfilePropertiesResponse) LinuxProfilePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret LinuxProfilePropertiesResponse
		return ret
	}).(LinuxProfilePropertiesResponseOutput)
}

func (o LinuxProfilePropertiesResponsePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxProfilePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o LinuxProfilePropertiesResponsePtrOutput) Ssh() LinuxProfilePropertiesResponseSshPtrOutput {
	return o.ApplyT(func(v *LinuxProfilePropertiesResponse) *LinuxProfilePropertiesResponseSsh {
		if v == nil {
			return nil
		}
		return v.Ssh
	}).(LinuxProfilePropertiesResponseSshPtrOutput)
}

type LinuxProfilePropertiesResponsePublicKeys struct {
	KeyData *string `pulumi:"keyData"`
}

type LinuxProfilePropertiesResponsePublicKeysOutput struct{ *pulumi.OutputState }

func (LinuxProfilePropertiesResponsePublicKeysOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxProfilePropertiesResponsePublicKeys)(nil)).Elem()
}

func (o LinuxProfilePropertiesResponsePublicKeysOutput) ToLinuxProfilePropertiesResponsePublicKeysOutput() LinuxProfilePropertiesResponsePublicKeysOutput {
	return o
}

func (o LinuxProfilePropertiesResponsePublicKeysOutput) ToLinuxProfilePropertiesResponsePublicKeysOutputWithContext(ctx context.Context) LinuxProfilePropertiesResponsePublicKeysOutput {
	return o
}

func (o LinuxProfilePropertiesResponsePublicKeysOutput) KeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxProfilePropertiesResponsePublicKeys) *string { return v.KeyData }).(pulumi.StringPtrOutput)
}

type LinuxProfilePropertiesResponsePublicKeysArrayOutput struct{ *pulumi.OutputState }

func (LinuxProfilePropertiesResponsePublicKeysArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinuxProfilePropertiesResponsePublicKeys)(nil)).Elem()
}

func (o LinuxProfilePropertiesResponsePublicKeysArrayOutput) ToLinuxProfilePropertiesResponsePublicKeysArrayOutput() LinuxProfilePropertiesResponsePublicKeysArrayOutput {
	return o
}

func (o LinuxProfilePropertiesResponsePublicKeysArrayOutput) ToLinuxProfilePropertiesResponsePublicKeysArrayOutputWithContext(ctx context.Context) LinuxProfilePropertiesResponsePublicKeysArrayOutput {
	return o
}

func (o LinuxProfilePropertiesResponsePublicKeysArrayOutput) Index(i pulumi.IntInput) LinuxProfilePropertiesResponsePublicKeysOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinuxProfilePropertiesResponsePublicKeys {
		return vs[0].([]LinuxProfilePropertiesResponsePublicKeys)[vs[1].(int)]
	}).(LinuxProfilePropertiesResponsePublicKeysOutput)
}

type LinuxProfilePropertiesResponseSsh struct {
	PublicKeys []LinuxProfilePropertiesResponsePublicKeys `pulumi:"publicKeys"`
}

type LinuxProfilePropertiesResponseSshOutput struct{ *pulumi.OutputState }

func (LinuxProfilePropertiesResponseSshOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxProfilePropertiesResponseSsh)(nil)).Elem()
}

func (o LinuxProfilePropertiesResponseSshOutput) ToLinuxProfilePropertiesResponseSshOutput() LinuxProfilePropertiesResponseSshOutput {
	return o
}

func (o LinuxProfilePropertiesResponseSshOutput) ToLinuxProfilePropertiesResponseSshOutputWithContext(ctx context.Context) LinuxProfilePropertiesResponseSshOutput {
	return o
}

func (o LinuxProfilePropertiesResponseSshOutput) PublicKeys() LinuxProfilePropertiesResponsePublicKeysArrayOutput {
	return o.ApplyT(func(v LinuxProfilePropertiesResponseSsh) []LinuxProfilePropertiesResponsePublicKeys {
		return v.PublicKeys
	}).(LinuxProfilePropertiesResponsePublicKeysArrayOutput)
}

type LinuxProfilePropertiesResponseSshPtrOutput struct{ *pulumi.OutputState }

func (LinuxProfilePropertiesResponseSshPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxProfilePropertiesResponseSsh)(nil)).Elem()
}

func (o LinuxProfilePropertiesResponseSshPtrOutput) ToLinuxProfilePropertiesResponseSshPtrOutput() LinuxProfilePropertiesResponseSshPtrOutput {
	return o
}

func (o LinuxProfilePropertiesResponseSshPtrOutput) ToLinuxProfilePropertiesResponseSshPtrOutputWithContext(ctx context.Context) LinuxProfilePropertiesResponseSshPtrOutput {
	return o
}

func (o LinuxProfilePropertiesResponseSshPtrOutput) Elem() LinuxProfilePropertiesResponseSshOutput {
	return o.ApplyT(func(v *LinuxProfilePropertiesResponseSsh) LinuxProfilePropertiesResponseSsh {
		if v != nil {
			return *v
		}
		var ret LinuxProfilePropertiesResponseSsh
		return ret
	}).(LinuxProfilePropertiesResponseSshOutput)
}

func (o LinuxProfilePropertiesResponseSshPtrOutput) PublicKeys() LinuxProfilePropertiesResponsePublicKeysArrayOutput {
	return o.ApplyT(func(v *LinuxProfilePropertiesResponseSsh) []LinuxProfilePropertiesResponsePublicKeys {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(LinuxProfilePropertiesResponsePublicKeysArrayOutput)
}

type LinuxProfilePropertiesSsh struct {
	PublicKeys []LinuxProfilePropertiesPublicKeys `pulumi:"publicKeys"`
}





type LinuxProfilePropertiesSshInput interface {
	pulumi.Input

	ToLinuxProfilePropertiesSshOutput() LinuxProfilePropertiesSshOutput
	ToLinuxProfilePropertiesSshOutputWithContext(context.Context) LinuxProfilePropertiesSshOutput
}

type LinuxProfilePropertiesSshArgs struct {
	PublicKeys LinuxProfilePropertiesPublicKeysArrayInput `pulumi:"publicKeys"`
}

func (LinuxProfilePropertiesSshArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxProfilePropertiesSsh)(nil)).Elem()
}

func (i LinuxProfilePropertiesSshArgs) ToLinuxProfilePropertiesSshOutput() LinuxProfilePropertiesSshOutput {
	return i.ToLinuxProfilePropertiesSshOutputWithContext(context.Background())
}

func (i LinuxProfilePropertiesSshArgs) ToLinuxProfilePropertiesSshOutputWithContext(ctx context.Context) LinuxProfilePropertiesSshOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxProfilePropertiesSshOutput)
}

func (i LinuxProfilePropertiesSshArgs) ToLinuxProfilePropertiesSshPtrOutput() LinuxProfilePropertiesSshPtrOutput {
	return i.ToLinuxProfilePropertiesSshPtrOutputWithContext(context.Background())
}

func (i LinuxProfilePropertiesSshArgs) ToLinuxProfilePropertiesSshPtrOutputWithContext(ctx context.Context) LinuxProfilePropertiesSshPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxProfilePropertiesSshOutput).ToLinuxProfilePropertiesSshPtrOutputWithContext(ctx)
}









type LinuxProfilePropertiesSshPtrInput interface {
	pulumi.Input

	ToLinuxProfilePropertiesSshPtrOutput() LinuxProfilePropertiesSshPtrOutput
	ToLinuxProfilePropertiesSshPtrOutputWithContext(context.Context) LinuxProfilePropertiesSshPtrOutput
}

type linuxProfilePropertiesSshPtrType LinuxProfilePropertiesSshArgs

func LinuxProfilePropertiesSshPtr(v *LinuxProfilePropertiesSshArgs) LinuxProfilePropertiesSshPtrInput {
	return (*linuxProfilePropertiesSshPtrType)(v)
}

func (*linuxProfilePropertiesSshPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxProfilePropertiesSsh)(nil)).Elem()
}

func (i *linuxProfilePropertiesSshPtrType) ToLinuxProfilePropertiesSshPtrOutput() LinuxProfilePropertiesSshPtrOutput {
	return i.ToLinuxProfilePropertiesSshPtrOutputWithContext(context.Background())
}

func (i *linuxProfilePropertiesSshPtrType) ToLinuxProfilePropertiesSshPtrOutputWithContext(ctx context.Context) LinuxProfilePropertiesSshPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxProfilePropertiesSshPtrOutput)
}

type LinuxProfilePropertiesSshOutput struct{ *pulumi.OutputState }

func (LinuxProfilePropertiesSshOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxProfilePropertiesSsh)(nil)).Elem()
}

func (o LinuxProfilePropertiesSshOutput) ToLinuxProfilePropertiesSshOutput() LinuxProfilePropertiesSshOutput {
	return o
}

func (o LinuxProfilePropertiesSshOutput) ToLinuxProfilePropertiesSshOutputWithContext(ctx context.Context) LinuxProfilePropertiesSshOutput {
	return o
}

func (o LinuxProfilePropertiesSshOutput) ToLinuxProfilePropertiesSshPtrOutput() LinuxProfilePropertiesSshPtrOutput {
	return o.ToLinuxProfilePropertiesSshPtrOutputWithContext(context.Background())
}

func (o LinuxProfilePropertiesSshOutput) ToLinuxProfilePropertiesSshPtrOutputWithContext(ctx context.Context) LinuxProfilePropertiesSshPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinuxProfilePropertiesSsh) *LinuxProfilePropertiesSsh {
		return &v
	}).(LinuxProfilePropertiesSshPtrOutput)
}

func (o LinuxProfilePropertiesSshOutput) PublicKeys() LinuxProfilePropertiesPublicKeysArrayOutput {
	return o.ApplyT(func(v LinuxProfilePropertiesSsh) []LinuxProfilePropertiesPublicKeys { return v.PublicKeys }).(LinuxProfilePropertiesPublicKeysArrayOutput)
}

type LinuxProfilePropertiesSshPtrOutput struct{ *pulumi.OutputState }

func (LinuxProfilePropertiesSshPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxProfilePropertiesSsh)(nil)).Elem()
}

func (o LinuxProfilePropertiesSshPtrOutput) ToLinuxProfilePropertiesSshPtrOutput() LinuxProfilePropertiesSshPtrOutput {
	return o
}

func (o LinuxProfilePropertiesSshPtrOutput) ToLinuxProfilePropertiesSshPtrOutputWithContext(ctx context.Context) LinuxProfilePropertiesSshPtrOutput {
	return o
}

func (o LinuxProfilePropertiesSshPtrOutput) Elem() LinuxProfilePropertiesSshOutput {
	return o.ApplyT(func(v *LinuxProfilePropertiesSsh) LinuxProfilePropertiesSsh {
		if v != nil {
			return *v
		}
		var ret LinuxProfilePropertiesSsh
		return ret
	}).(LinuxProfilePropertiesSshOutput)
}

func (o LinuxProfilePropertiesSshPtrOutput) PublicKeys() LinuxProfilePropertiesPublicKeysArrayOutput {
	return o.ApplyT(func(v *LinuxProfilePropertiesSsh) []LinuxProfilePropertiesPublicKeys {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(LinuxProfilePropertiesPublicKeysArrayOutput)
}

type LoadBalancerProfile struct {
	AvailabilityZones    []string                `pulumi:"availabilityZones"`
	CloudProviderProfile *CloudProviderProfile   `pulumi:"cloudProviderProfile"`
	Count                *int                    `pulumi:"count"`
	LinuxProfile         *LinuxProfileProperties `pulumi:"linuxProfile"`
	MaxCount             *int                    `pulumi:"maxCount"`
	MaxPods              *int                    `pulumi:"maxPods"`
	MinCount             *int                    `pulumi:"minCount"`
	Mode                 *string                 `pulumi:"mode"`
	Name                 *string                 `pulumi:"name"`
	NodeImageVersion     *string                 `pulumi:"nodeImageVersion"`
	NodeLabels           map[string]string       `pulumi:"nodeLabels"`
	NodeTaints           []string                `pulumi:"nodeTaints"`
	OsType               *string                 `pulumi:"osType"`
	VmSize               *string                 `pulumi:"vmSize"`
}


func (val *LoadBalancerProfile) Defaults() *LoadBalancerProfile {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		count_ := 1
		tmp.Count = &count_
	}
	if isZero(tmp.Mode) {
		mode_ := "User"
		tmp.Mode = &mode_
	}
	if isZero(tmp.OsType) {
		osType_ := "Linux"
		tmp.OsType = &osType_
	}
	return &tmp
}





type LoadBalancerProfileInput interface {
	pulumi.Input

	ToLoadBalancerProfileOutput() LoadBalancerProfileOutput
	ToLoadBalancerProfileOutputWithContext(context.Context) LoadBalancerProfileOutput
}

type LoadBalancerProfileArgs struct {
	AvailabilityZones    pulumi.StringArrayInput        `pulumi:"availabilityZones"`
	CloudProviderProfile CloudProviderProfilePtrInput   `pulumi:"cloudProviderProfile"`
	Count                pulumi.IntPtrInput             `pulumi:"count"`
	LinuxProfile         LinuxProfilePropertiesPtrInput `pulumi:"linuxProfile"`
	MaxCount             pulumi.IntPtrInput             `pulumi:"maxCount"`
	MaxPods              pulumi.IntPtrInput             `pulumi:"maxPods"`
	MinCount             pulumi.IntPtrInput             `pulumi:"minCount"`
	Mode                 pulumi.StringPtrInput          `pulumi:"mode"`
	Name                 pulumi.StringPtrInput          `pulumi:"name"`
	NodeImageVersion     pulumi.StringPtrInput          `pulumi:"nodeImageVersion"`
	NodeLabels           pulumi.StringMapInput          `pulumi:"nodeLabels"`
	NodeTaints           pulumi.StringArrayInput        `pulumi:"nodeTaints"`
	OsType               pulumi.StringPtrInput          `pulumi:"osType"`
	VmSize               pulumi.StringPtrInput          `pulumi:"vmSize"`
}


func (val *LoadBalancerProfileArgs) Defaults() *LoadBalancerProfileArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		tmp.Count = pulumi.IntPtr(1)
	}
	if isZero(tmp.Mode) {
		tmp.Mode = pulumi.StringPtr("User")
	}
	if isZero(tmp.OsType) {
		tmp.OsType = pulumi.StringPtr("Linux")
	}
	return &tmp
}
func (LoadBalancerProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerProfile)(nil)).Elem()
}

func (i LoadBalancerProfileArgs) ToLoadBalancerProfileOutput() LoadBalancerProfileOutput {
	return i.ToLoadBalancerProfileOutputWithContext(context.Background())
}

func (i LoadBalancerProfileArgs) ToLoadBalancerProfileOutputWithContext(ctx context.Context) LoadBalancerProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerProfileOutput)
}

func (i LoadBalancerProfileArgs) ToLoadBalancerProfilePtrOutput() LoadBalancerProfilePtrOutput {
	return i.ToLoadBalancerProfilePtrOutputWithContext(context.Background())
}

func (i LoadBalancerProfileArgs) ToLoadBalancerProfilePtrOutputWithContext(ctx context.Context) LoadBalancerProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerProfileOutput).ToLoadBalancerProfilePtrOutputWithContext(ctx)
}









type LoadBalancerProfilePtrInput interface {
	pulumi.Input

	ToLoadBalancerProfilePtrOutput() LoadBalancerProfilePtrOutput
	ToLoadBalancerProfilePtrOutputWithContext(context.Context) LoadBalancerProfilePtrOutput
}

type loadBalancerProfilePtrType LoadBalancerProfileArgs

func LoadBalancerProfilePtr(v *LoadBalancerProfileArgs) LoadBalancerProfilePtrInput {
	return (*loadBalancerProfilePtrType)(v)
}

func (*loadBalancerProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerProfile)(nil)).Elem()
}

func (i *loadBalancerProfilePtrType) ToLoadBalancerProfilePtrOutput() LoadBalancerProfilePtrOutput {
	return i.ToLoadBalancerProfilePtrOutputWithContext(context.Background())
}

func (i *loadBalancerProfilePtrType) ToLoadBalancerProfilePtrOutputWithContext(ctx context.Context) LoadBalancerProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerProfilePtrOutput)
}

type LoadBalancerProfileOutput struct{ *pulumi.OutputState }

func (LoadBalancerProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerProfile)(nil)).Elem()
}

func (o LoadBalancerProfileOutput) ToLoadBalancerProfileOutput() LoadBalancerProfileOutput {
	return o
}

func (o LoadBalancerProfileOutput) ToLoadBalancerProfileOutputWithContext(ctx context.Context) LoadBalancerProfileOutput {
	return o
}

func (o LoadBalancerProfileOutput) ToLoadBalancerProfilePtrOutput() LoadBalancerProfilePtrOutput {
	return o.ToLoadBalancerProfilePtrOutputWithContext(context.Background())
}

func (o LoadBalancerProfileOutput) ToLoadBalancerProfilePtrOutputWithContext(ctx context.Context) LoadBalancerProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoadBalancerProfile) *LoadBalancerProfile {
		return &v
	}).(LoadBalancerProfilePtrOutput)
}

func (o LoadBalancerProfileOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LoadBalancerProfile) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o LoadBalancerProfileOutput) CloudProviderProfile() CloudProviderProfilePtrOutput {
	return o.ApplyT(func(v LoadBalancerProfile) *CloudProviderProfile { return v.CloudProviderProfile }).(CloudProviderProfilePtrOutput)
}

func (o LoadBalancerProfileOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancerProfile) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o LoadBalancerProfileOutput) LinuxProfile() LinuxProfilePropertiesPtrOutput {
	return o.ApplyT(func(v LoadBalancerProfile) *LinuxProfileProperties { return v.LinuxProfile }).(LinuxProfilePropertiesPtrOutput)
}

func (o LoadBalancerProfileOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancerProfile) *int { return v.MaxCount }).(pulumi.IntPtrOutput)
}

func (o LoadBalancerProfileOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancerProfile) *int { return v.MaxPods }).(pulumi.IntPtrOutput)
}

func (o LoadBalancerProfileOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancerProfile) *int { return v.MinCount }).(pulumi.IntPtrOutput)
}

func (o LoadBalancerProfileOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerProfile) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerProfileOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerProfile) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerProfileOutput) NodeImageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerProfile) *string { return v.NodeImageVersion }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerProfileOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LoadBalancerProfile) map[string]string { return v.NodeLabels }).(pulumi.StringMapOutput)
}

func (o LoadBalancerProfileOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LoadBalancerProfile) []string { return v.NodeTaints }).(pulumi.StringArrayOutput)
}

func (o LoadBalancerProfileOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerProfile) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerProfileOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerProfile) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type LoadBalancerProfilePtrOutput struct{ *pulumi.OutputState }

func (LoadBalancerProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerProfile)(nil)).Elem()
}

func (o LoadBalancerProfilePtrOutput) ToLoadBalancerProfilePtrOutput() LoadBalancerProfilePtrOutput {
	return o
}

func (o LoadBalancerProfilePtrOutput) ToLoadBalancerProfilePtrOutputWithContext(ctx context.Context) LoadBalancerProfilePtrOutput {
	return o
}

func (o LoadBalancerProfilePtrOutput) Elem() LoadBalancerProfileOutput {
	return o.ApplyT(func(v *LoadBalancerProfile) LoadBalancerProfile {
		if v != nil {
			return *v
		}
		var ret LoadBalancerProfile
		return ret
	}).(LoadBalancerProfileOutput)
}

func (o LoadBalancerProfilePtrOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LoadBalancerProfile) []string {
		if v == nil {
			return nil
		}
		return v.AvailabilityZones
	}).(pulumi.StringArrayOutput)
}

func (o LoadBalancerProfilePtrOutput) CloudProviderProfile() CloudProviderProfilePtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfile) *CloudProviderProfile {
		if v == nil {
			return nil
		}
		return v.CloudProviderProfile
	}).(CloudProviderProfilePtrOutput)
}

func (o LoadBalancerProfilePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfile) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o LoadBalancerProfilePtrOutput) LinuxProfile() LinuxProfilePropertiesPtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfile) *LinuxProfileProperties {
		if v == nil {
			return nil
		}
		return v.LinuxProfile
	}).(LinuxProfilePropertiesPtrOutput)
}

func (o LoadBalancerProfilePtrOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfile) *int {
		if v == nil {
			return nil
		}
		return v.MaxCount
	}).(pulumi.IntPtrOutput)
}

func (o LoadBalancerProfilePtrOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfile) *int {
		if v == nil {
			return nil
		}
		return v.MaxPods
	}).(pulumi.IntPtrOutput)
}

func (o LoadBalancerProfilePtrOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfile) *int {
		if v == nil {
			return nil
		}
		return v.MinCount
	}).(pulumi.IntPtrOutput)
}

func (o LoadBalancerProfilePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfile) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

func (o LoadBalancerProfilePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfile) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o LoadBalancerProfilePtrOutput) NodeImageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfile) *string {
		if v == nil {
			return nil
		}
		return v.NodeImageVersion
	}).(pulumi.StringPtrOutput)
}

func (o LoadBalancerProfilePtrOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LoadBalancerProfile) map[string]string {
		if v == nil {
			return nil
		}
		return v.NodeLabels
	}).(pulumi.StringMapOutput)
}

func (o LoadBalancerProfilePtrOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LoadBalancerProfile) []string {
		if v == nil {
			return nil
		}
		return v.NodeTaints
	}).(pulumi.StringArrayOutput)
}

func (o LoadBalancerProfilePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfile) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o LoadBalancerProfilePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfile) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type LoadBalancerProfileResponse struct {
	AvailabilityZones    []string                        `pulumi:"availabilityZones"`
	CloudProviderProfile *CloudProviderProfileResponse   `pulumi:"cloudProviderProfile"`
	Count                *int                            `pulumi:"count"`
	LinuxProfile         *LinuxProfilePropertiesResponse `pulumi:"linuxProfile"`
	MaxCount             *int                            `pulumi:"maxCount"`
	MaxPods              *int                            `pulumi:"maxPods"`
	MinCount             *int                            `pulumi:"minCount"`
	Mode                 *string                         `pulumi:"mode"`
	Name                 *string                         `pulumi:"name"`
	NodeImageVersion     *string                         `pulumi:"nodeImageVersion"`
	NodeLabels           map[string]string               `pulumi:"nodeLabels"`
	NodeTaints           []string                        `pulumi:"nodeTaints"`
	OsType               *string                         `pulumi:"osType"`
	VmSize               *string                         `pulumi:"vmSize"`
}


func (val *LoadBalancerProfileResponse) Defaults() *LoadBalancerProfileResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		count_ := 1
		tmp.Count = &count_
	}
	if isZero(tmp.Mode) {
		mode_ := "User"
		tmp.Mode = &mode_
	}
	if isZero(tmp.OsType) {
		osType_ := "Linux"
		tmp.OsType = &osType_
	}
	return &tmp
}

type LoadBalancerProfileResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancerProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerProfileResponse)(nil)).Elem()
}

func (o LoadBalancerProfileResponseOutput) ToLoadBalancerProfileResponseOutput() LoadBalancerProfileResponseOutput {
	return o
}

func (o LoadBalancerProfileResponseOutput) ToLoadBalancerProfileResponseOutputWithContext(ctx context.Context) LoadBalancerProfileResponseOutput {
	return o
}

func (o LoadBalancerProfileResponseOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LoadBalancerProfileResponse) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o LoadBalancerProfileResponseOutput) CloudProviderProfile() CloudProviderProfileResponsePtrOutput {
	return o.ApplyT(func(v LoadBalancerProfileResponse) *CloudProviderProfileResponse { return v.CloudProviderProfile }).(CloudProviderProfileResponsePtrOutput)
}

func (o LoadBalancerProfileResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancerProfileResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o LoadBalancerProfileResponseOutput) LinuxProfile() LinuxProfilePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LoadBalancerProfileResponse) *LinuxProfilePropertiesResponse { return v.LinuxProfile }).(LinuxProfilePropertiesResponsePtrOutput)
}

func (o LoadBalancerProfileResponseOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancerProfileResponse) *int { return v.MaxCount }).(pulumi.IntPtrOutput)
}

func (o LoadBalancerProfileResponseOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancerProfileResponse) *int { return v.MaxPods }).(pulumi.IntPtrOutput)
}

func (o LoadBalancerProfileResponseOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancerProfileResponse) *int { return v.MinCount }).(pulumi.IntPtrOutput)
}

func (o LoadBalancerProfileResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerProfileResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerProfileResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerProfileResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerProfileResponseOutput) NodeImageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerProfileResponse) *string { return v.NodeImageVersion }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerProfileResponseOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LoadBalancerProfileResponse) map[string]string { return v.NodeLabels }).(pulumi.StringMapOutput)
}

func (o LoadBalancerProfileResponseOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LoadBalancerProfileResponse) []string { return v.NodeTaints }).(pulumi.StringArrayOutput)
}

func (o LoadBalancerProfileResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerProfileResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerProfileResponseOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerProfileResponse) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type LoadBalancerProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (LoadBalancerProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerProfileResponse)(nil)).Elem()
}

func (o LoadBalancerProfileResponsePtrOutput) ToLoadBalancerProfileResponsePtrOutput() LoadBalancerProfileResponsePtrOutput {
	return o
}

func (o LoadBalancerProfileResponsePtrOutput) ToLoadBalancerProfileResponsePtrOutputWithContext(ctx context.Context) LoadBalancerProfileResponsePtrOutput {
	return o
}

func (o LoadBalancerProfileResponsePtrOutput) Elem() LoadBalancerProfileResponseOutput {
	return o.ApplyT(func(v *LoadBalancerProfileResponse) LoadBalancerProfileResponse {
		if v != nil {
			return *v
		}
		var ret LoadBalancerProfileResponse
		return ret
	}).(LoadBalancerProfileResponseOutput)
}

func (o LoadBalancerProfileResponsePtrOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LoadBalancerProfileResponse) []string {
		if v == nil {
			return nil
		}
		return v.AvailabilityZones
	}).(pulumi.StringArrayOutput)
}

func (o LoadBalancerProfileResponsePtrOutput) CloudProviderProfile() CloudProviderProfileResponsePtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfileResponse) *CloudProviderProfileResponse {
		if v == nil {
			return nil
		}
		return v.CloudProviderProfile
	}).(CloudProviderProfileResponsePtrOutput)
}

func (o LoadBalancerProfileResponsePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o LoadBalancerProfileResponsePtrOutput) LinuxProfile() LinuxProfilePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfileResponse) *LinuxProfilePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.LinuxProfile
	}).(LinuxProfilePropertiesResponsePtrOutput)
}

func (o LoadBalancerProfileResponsePtrOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxCount
	}).(pulumi.IntPtrOutput)
}

func (o LoadBalancerProfileResponsePtrOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxPods
	}).(pulumi.IntPtrOutput)
}

func (o LoadBalancerProfileResponsePtrOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.MinCount
	}).(pulumi.IntPtrOutput)
}

func (o LoadBalancerProfileResponsePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

func (o LoadBalancerProfileResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o LoadBalancerProfileResponsePtrOutput) NodeImageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.NodeImageVersion
	}).(pulumi.StringPtrOutput)
}

func (o LoadBalancerProfileResponsePtrOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LoadBalancerProfileResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.NodeLabels
	}).(pulumi.StringMapOutput)
}

func (o LoadBalancerProfileResponsePtrOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LoadBalancerProfileResponse) []string {
		if v == nil {
			return nil
		}
		return v.NodeTaints
	}).(pulumi.StringArrayOutput)
}

func (o LoadBalancerProfileResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o LoadBalancerProfileResponsePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancerProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type NamedAgentPoolProfile struct {
	AvailabilityZones    []string              `pulumi:"availabilityZones"`
	CloudProviderProfile *CloudProviderProfile `pulumi:"cloudProviderProfile"`
	Count                *int                  `pulumi:"count"`
	MaxCount             *int                  `pulumi:"maxCount"`
	MaxPods              *int                  `pulumi:"maxPods"`
	MinCount             *int                  `pulumi:"minCount"`
	Mode                 *string               `pulumi:"mode"`
	Name                 *string               `pulumi:"name"`
	NodeImageVersion     *string               `pulumi:"nodeImageVersion"`
	NodeLabels           map[string]string     `pulumi:"nodeLabels"`
	NodeTaints           []string              `pulumi:"nodeTaints"`
	OsType               *string               `pulumi:"osType"`
	VmSize               *string               `pulumi:"vmSize"`
}


func (val *NamedAgentPoolProfile) Defaults() *NamedAgentPoolProfile {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		count_ := 1
		tmp.Count = &count_
	}
	if isZero(tmp.Mode) {
		mode_ := "User"
		tmp.Mode = &mode_
	}
	if isZero(tmp.OsType) {
		osType_ := "Linux"
		tmp.OsType = &osType_
	}
	return &tmp
}





type NamedAgentPoolProfileInput interface {
	pulumi.Input

	ToNamedAgentPoolProfileOutput() NamedAgentPoolProfileOutput
	ToNamedAgentPoolProfileOutputWithContext(context.Context) NamedAgentPoolProfileOutput
}

type NamedAgentPoolProfileArgs struct {
	AvailabilityZones    pulumi.StringArrayInput      `pulumi:"availabilityZones"`
	CloudProviderProfile CloudProviderProfilePtrInput `pulumi:"cloudProviderProfile"`
	Count                pulumi.IntPtrInput           `pulumi:"count"`
	MaxCount             pulumi.IntPtrInput           `pulumi:"maxCount"`
	MaxPods              pulumi.IntPtrInput           `pulumi:"maxPods"`
	MinCount             pulumi.IntPtrInput           `pulumi:"minCount"`
	Mode                 pulumi.StringPtrInput        `pulumi:"mode"`
	Name                 pulumi.StringPtrInput        `pulumi:"name"`
	NodeImageVersion     pulumi.StringPtrInput        `pulumi:"nodeImageVersion"`
	NodeLabels           pulumi.StringMapInput        `pulumi:"nodeLabels"`
	NodeTaints           pulumi.StringArrayInput      `pulumi:"nodeTaints"`
	OsType               pulumi.StringPtrInput        `pulumi:"osType"`
	VmSize               pulumi.StringPtrInput        `pulumi:"vmSize"`
}


func (val *NamedAgentPoolProfileArgs) Defaults() *NamedAgentPoolProfileArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		tmp.Count = pulumi.IntPtr(1)
	}
	if isZero(tmp.Mode) {
		tmp.Mode = pulumi.StringPtr("User")
	}
	if isZero(tmp.OsType) {
		tmp.OsType = pulumi.StringPtr("Linux")
	}
	return &tmp
}
func (NamedAgentPoolProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NamedAgentPoolProfile)(nil)).Elem()
}

func (i NamedAgentPoolProfileArgs) ToNamedAgentPoolProfileOutput() NamedAgentPoolProfileOutput {
	return i.ToNamedAgentPoolProfileOutputWithContext(context.Background())
}

func (i NamedAgentPoolProfileArgs) ToNamedAgentPoolProfileOutputWithContext(ctx context.Context) NamedAgentPoolProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedAgentPoolProfileOutput)
}





type NamedAgentPoolProfileArrayInput interface {
	pulumi.Input

	ToNamedAgentPoolProfileArrayOutput() NamedAgentPoolProfileArrayOutput
	ToNamedAgentPoolProfileArrayOutputWithContext(context.Context) NamedAgentPoolProfileArrayOutput
}

type NamedAgentPoolProfileArray []NamedAgentPoolProfileInput

func (NamedAgentPoolProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamedAgentPoolProfile)(nil)).Elem()
}

func (i NamedAgentPoolProfileArray) ToNamedAgentPoolProfileArrayOutput() NamedAgentPoolProfileArrayOutput {
	return i.ToNamedAgentPoolProfileArrayOutputWithContext(context.Background())
}

func (i NamedAgentPoolProfileArray) ToNamedAgentPoolProfileArrayOutputWithContext(ctx context.Context) NamedAgentPoolProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedAgentPoolProfileArrayOutput)
}

type NamedAgentPoolProfileOutput struct{ *pulumi.OutputState }

func (NamedAgentPoolProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamedAgentPoolProfile)(nil)).Elem()
}

func (o NamedAgentPoolProfileOutput) ToNamedAgentPoolProfileOutput() NamedAgentPoolProfileOutput {
	return o
}

func (o NamedAgentPoolProfileOutput) ToNamedAgentPoolProfileOutputWithContext(ctx context.Context) NamedAgentPoolProfileOutput {
	return o
}

func (o NamedAgentPoolProfileOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NamedAgentPoolProfile) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o NamedAgentPoolProfileOutput) CloudProviderProfile() CloudProviderProfilePtrOutput {
	return o.ApplyT(func(v NamedAgentPoolProfile) *CloudProviderProfile { return v.CloudProviderProfile }).(CloudProviderProfilePtrOutput)
}

func (o NamedAgentPoolProfileOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NamedAgentPoolProfile) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o NamedAgentPoolProfileOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NamedAgentPoolProfile) *int { return v.MaxCount }).(pulumi.IntPtrOutput)
}

func (o NamedAgentPoolProfileOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NamedAgentPoolProfile) *int { return v.MaxPods }).(pulumi.IntPtrOutput)
}

func (o NamedAgentPoolProfileOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NamedAgentPoolProfile) *int { return v.MinCount }).(pulumi.IntPtrOutput)
}

func (o NamedAgentPoolProfileOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamedAgentPoolProfile) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o NamedAgentPoolProfileOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamedAgentPoolProfile) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NamedAgentPoolProfileOutput) NodeImageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamedAgentPoolProfile) *string { return v.NodeImageVersion }).(pulumi.StringPtrOutput)
}

func (o NamedAgentPoolProfileOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v NamedAgentPoolProfile) map[string]string { return v.NodeLabels }).(pulumi.StringMapOutput)
}

func (o NamedAgentPoolProfileOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NamedAgentPoolProfile) []string { return v.NodeTaints }).(pulumi.StringArrayOutput)
}

func (o NamedAgentPoolProfileOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamedAgentPoolProfile) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o NamedAgentPoolProfileOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamedAgentPoolProfile) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type NamedAgentPoolProfileArrayOutput struct{ *pulumi.OutputState }

func (NamedAgentPoolProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamedAgentPoolProfile)(nil)).Elem()
}

func (o NamedAgentPoolProfileArrayOutput) ToNamedAgentPoolProfileArrayOutput() NamedAgentPoolProfileArrayOutput {
	return o
}

func (o NamedAgentPoolProfileArrayOutput) ToNamedAgentPoolProfileArrayOutputWithContext(ctx context.Context) NamedAgentPoolProfileArrayOutput {
	return o
}

func (o NamedAgentPoolProfileArrayOutput) Index(i pulumi.IntInput) NamedAgentPoolProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NamedAgentPoolProfile {
		return vs[0].([]NamedAgentPoolProfile)[vs[1].(int)]
	}).(NamedAgentPoolProfileOutput)
}

type NamedAgentPoolProfileResponse struct {
	AvailabilityZones    []string                      `pulumi:"availabilityZones"`
	CloudProviderProfile *CloudProviderProfileResponse `pulumi:"cloudProviderProfile"`
	Count                *int                          `pulumi:"count"`
	MaxCount             *int                          `pulumi:"maxCount"`
	MaxPods              *int                          `pulumi:"maxPods"`
	MinCount             *int                          `pulumi:"minCount"`
	Mode                 *string                       `pulumi:"mode"`
	Name                 *string                       `pulumi:"name"`
	NodeImageVersion     *string                       `pulumi:"nodeImageVersion"`
	NodeLabels           map[string]string             `pulumi:"nodeLabels"`
	NodeTaints           []string                      `pulumi:"nodeTaints"`
	OsType               *string                       `pulumi:"osType"`
	VmSize               *string                       `pulumi:"vmSize"`
}


func (val *NamedAgentPoolProfileResponse) Defaults() *NamedAgentPoolProfileResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		count_ := 1
		tmp.Count = &count_
	}
	if isZero(tmp.Mode) {
		mode_ := "User"
		tmp.Mode = &mode_
	}
	if isZero(tmp.OsType) {
		osType_ := "Linux"
		tmp.OsType = &osType_
	}
	return &tmp
}

type NamedAgentPoolProfileResponseOutput struct{ *pulumi.OutputState }

func (NamedAgentPoolProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamedAgentPoolProfileResponse)(nil)).Elem()
}

func (o NamedAgentPoolProfileResponseOutput) ToNamedAgentPoolProfileResponseOutput() NamedAgentPoolProfileResponseOutput {
	return o
}

func (o NamedAgentPoolProfileResponseOutput) ToNamedAgentPoolProfileResponseOutputWithContext(ctx context.Context) NamedAgentPoolProfileResponseOutput {
	return o
}

func (o NamedAgentPoolProfileResponseOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NamedAgentPoolProfileResponse) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o NamedAgentPoolProfileResponseOutput) CloudProviderProfile() CloudProviderProfileResponsePtrOutput {
	return o.ApplyT(func(v NamedAgentPoolProfileResponse) *CloudProviderProfileResponse { return v.CloudProviderProfile }).(CloudProviderProfileResponsePtrOutput)
}

func (o NamedAgentPoolProfileResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NamedAgentPoolProfileResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o NamedAgentPoolProfileResponseOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NamedAgentPoolProfileResponse) *int { return v.MaxCount }).(pulumi.IntPtrOutput)
}

func (o NamedAgentPoolProfileResponseOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NamedAgentPoolProfileResponse) *int { return v.MaxPods }).(pulumi.IntPtrOutput)
}

func (o NamedAgentPoolProfileResponseOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NamedAgentPoolProfileResponse) *int { return v.MinCount }).(pulumi.IntPtrOutput)
}

func (o NamedAgentPoolProfileResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamedAgentPoolProfileResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o NamedAgentPoolProfileResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamedAgentPoolProfileResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NamedAgentPoolProfileResponseOutput) NodeImageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamedAgentPoolProfileResponse) *string { return v.NodeImageVersion }).(pulumi.StringPtrOutput)
}

func (o NamedAgentPoolProfileResponseOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v NamedAgentPoolProfileResponse) map[string]string { return v.NodeLabels }).(pulumi.StringMapOutput)
}

func (o NamedAgentPoolProfileResponseOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NamedAgentPoolProfileResponse) []string { return v.NodeTaints }).(pulumi.StringArrayOutput)
}

func (o NamedAgentPoolProfileResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamedAgentPoolProfileResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o NamedAgentPoolProfileResponseOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamedAgentPoolProfileResponse) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type NamedAgentPoolProfileResponseArrayOutput struct{ *pulumi.OutputState }

func (NamedAgentPoolProfileResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamedAgentPoolProfileResponse)(nil)).Elem()
}

func (o NamedAgentPoolProfileResponseArrayOutput) ToNamedAgentPoolProfileResponseArrayOutput() NamedAgentPoolProfileResponseArrayOutput {
	return o
}

func (o NamedAgentPoolProfileResponseArrayOutput) ToNamedAgentPoolProfileResponseArrayOutputWithContext(ctx context.Context) NamedAgentPoolProfileResponseArrayOutput {
	return o
}

func (o NamedAgentPoolProfileResponseArrayOutput) Index(i pulumi.IntInput) NamedAgentPoolProfileResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NamedAgentPoolProfileResponse {
		return vs[0].([]NamedAgentPoolProfileResponse)[vs[1].(int)]
	}).(NamedAgentPoolProfileResponseOutput)
}

type NetworkProfile struct {
	DnsServiceIP        *string              `pulumi:"dnsServiceIP"`
	LoadBalancerProfile *LoadBalancerProfile `pulumi:"loadBalancerProfile"`
	LoadBalancerSku     *string              `pulumi:"loadBalancerSku"`
	NetworkPolicy       *string              `pulumi:"networkPolicy"`
	PodCidr             *string              `pulumi:"podCidr"`
	PodCidrs            []string             `pulumi:"podCidrs"`
	ServiceCidr         *string              `pulumi:"serviceCidr"`
	ServiceCidrs        []string             `pulumi:"serviceCidrs"`
}


func (val *NetworkProfile) Defaults() *NetworkProfile {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.LoadBalancerProfile = tmp.LoadBalancerProfile.Defaults()

	if isZero(tmp.LoadBalancerSku) {
		loadBalancerSku_ := "unmanaged"
		tmp.LoadBalancerSku = &loadBalancerSku_
	}
	if isZero(tmp.NetworkPolicy) {
		networkPolicy_ := "calico"
		tmp.NetworkPolicy = &networkPolicy_
	}
	return &tmp
}





type NetworkProfileInput interface {
	pulumi.Input

	ToNetworkProfileOutput() NetworkProfileOutput
	ToNetworkProfileOutputWithContext(context.Context) NetworkProfileOutput
}

type NetworkProfileArgs struct {
	DnsServiceIP        pulumi.StringPtrInput       `pulumi:"dnsServiceIP"`
	LoadBalancerProfile LoadBalancerProfilePtrInput `pulumi:"loadBalancerProfile"`
	LoadBalancerSku     pulumi.StringPtrInput       `pulumi:"loadBalancerSku"`
	NetworkPolicy       pulumi.StringPtrInput       `pulumi:"networkPolicy"`
	PodCidr             pulumi.StringPtrInput       `pulumi:"podCidr"`
	PodCidrs            pulumi.StringArrayInput     `pulumi:"podCidrs"`
	ServiceCidr         pulumi.StringPtrInput       `pulumi:"serviceCidr"`
	ServiceCidrs        pulumi.StringArrayInput     `pulumi:"serviceCidrs"`
}


func (val *NetworkProfileArgs) Defaults() *NetworkProfileArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	if isZero(tmp.LoadBalancerSku) {
		tmp.LoadBalancerSku = pulumi.StringPtr("unmanaged")
	}
	if isZero(tmp.NetworkPolicy) {
		tmp.NetworkPolicy = pulumi.StringPtr("calico")
	}
	return &tmp
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

func (o NetworkProfileOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.DnsServiceIP }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) LoadBalancerProfile() LoadBalancerProfilePtrOutput {
	return o.ApplyT(func(v NetworkProfile) *LoadBalancerProfile { return v.LoadBalancerProfile }).(LoadBalancerProfilePtrOutput)
}

func (o NetworkProfileOutput) LoadBalancerSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.LoadBalancerSku }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) NetworkPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.NetworkPolicy }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.PodCidr }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) PodCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkProfile) []string { return v.PodCidrs }).(pulumi.StringArrayOutput)
}

func (o NetworkProfileOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.ServiceCidr }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) ServiceCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkProfile) []string { return v.ServiceCidrs }).(pulumi.StringArrayOutput)
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

func (o NetworkProfilePtrOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.DnsServiceIP
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfilePtrOutput) LoadBalancerProfile() LoadBalancerProfilePtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *LoadBalancerProfile {
		if v == nil {
			return nil
		}
		return v.LoadBalancerProfile
	}).(LoadBalancerProfilePtrOutput)
}

func (o NetworkProfilePtrOutput) LoadBalancerSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancerSku
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfilePtrOutput) NetworkPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.NetworkPolicy
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfilePtrOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.PodCidr
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfilePtrOutput) PodCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkProfile) []string {
		if v == nil {
			return nil
		}
		return v.PodCidrs
	}).(pulumi.StringArrayOutput)
}

func (o NetworkProfilePtrOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCidr
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfilePtrOutput) ServiceCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkProfile) []string {
		if v == nil {
			return nil
		}
		return v.ServiceCidrs
	}).(pulumi.StringArrayOutput)
}

type NetworkProfileResponse struct {
	DnsServiceIP        *string                      `pulumi:"dnsServiceIP"`
	LoadBalancerProfile *LoadBalancerProfileResponse `pulumi:"loadBalancerProfile"`
	LoadBalancerSku     *string                      `pulumi:"loadBalancerSku"`
	NetworkPolicy       *string                      `pulumi:"networkPolicy"`
	PodCidr             *string                      `pulumi:"podCidr"`
	PodCidrs            []string                     `pulumi:"podCidrs"`
	ServiceCidr         *string                      `pulumi:"serviceCidr"`
	ServiceCidrs        []string                     `pulumi:"serviceCidrs"`
}


func (val *NetworkProfileResponse) Defaults() *NetworkProfileResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.LoadBalancerProfile = tmp.LoadBalancerProfile.Defaults()

	if isZero(tmp.LoadBalancerSku) {
		loadBalancerSku_ := "unmanaged"
		tmp.LoadBalancerSku = &loadBalancerSku_
	}
	if isZero(tmp.NetworkPolicy) {
		networkPolicy_ := "calico"
		tmp.NetworkPolicy = &networkPolicy_
	}
	return &tmp
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

func (o NetworkProfileResponseOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.DnsServiceIP }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponseOutput) LoadBalancerProfile() LoadBalancerProfileResponsePtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *LoadBalancerProfileResponse { return v.LoadBalancerProfile }).(LoadBalancerProfileResponsePtrOutput)
}

func (o NetworkProfileResponseOutput) LoadBalancerSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.LoadBalancerSku }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponseOutput) NetworkPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.NetworkPolicy }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponseOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.PodCidr }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponseOutput) PodCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkProfileResponse) []string { return v.PodCidrs }).(pulumi.StringArrayOutput)
}

func (o NetworkProfileResponseOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.ServiceCidr }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponseOutput) ServiceCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkProfileResponse) []string { return v.ServiceCidrs }).(pulumi.StringArrayOutput)
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

func (o NetworkProfileResponsePtrOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.DnsServiceIP
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) LoadBalancerProfile() LoadBalancerProfileResponsePtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *LoadBalancerProfileResponse {
		if v == nil {
			return nil
		}
		return v.LoadBalancerProfile
	}).(LoadBalancerProfileResponsePtrOutput)
}

func (o NetworkProfileResponsePtrOutput) LoadBalancerSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancerSku
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) NetworkPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.NetworkPolicy
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.PodCidr
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) PodCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) []string {
		if v == nil {
			return nil
		}
		return v.PodCidrs
	}).(pulumi.StringArrayOutput)
}

func (o NetworkProfileResponsePtrOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCidr
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) ServiceCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) []string {
		if v == nil {
			return nil
		}
		return v.ServiceCidrs
	}).(pulumi.StringArrayOutput)
}

type ProvisionedClusterIdentity struct {
	Type ResourceIdentityType `pulumi:"type"`
}





type ProvisionedClusterIdentityInput interface {
	pulumi.Input

	ToProvisionedClusterIdentityOutput() ProvisionedClusterIdentityOutput
	ToProvisionedClusterIdentityOutputWithContext(context.Context) ProvisionedClusterIdentityOutput
}

type ProvisionedClusterIdentityArgs struct {
	Type ResourceIdentityTypeInput `pulumi:"type"`
}

func (ProvisionedClusterIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisionedClusterIdentity)(nil)).Elem()
}

func (i ProvisionedClusterIdentityArgs) ToProvisionedClusterIdentityOutput() ProvisionedClusterIdentityOutput {
	return i.ToProvisionedClusterIdentityOutputWithContext(context.Background())
}

func (i ProvisionedClusterIdentityArgs) ToProvisionedClusterIdentityOutputWithContext(ctx context.Context) ProvisionedClusterIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionedClusterIdentityOutput)
}

func (i ProvisionedClusterIdentityArgs) ToProvisionedClusterIdentityPtrOutput() ProvisionedClusterIdentityPtrOutput {
	return i.ToProvisionedClusterIdentityPtrOutputWithContext(context.Background())
}

func (i ProvisionedClusterIdentityArgs) ToProvisionedClusterIdentityPtrOutputWithContext(ctx context.Context) ProvisionedClusterIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionedClusterIdentityOutput).ToProvisionedClusterIdentityPtrOutputWithContext(ctx)
}









type ProvisionedClusterIdentityPtrInput interface {
	pulumi.Input

	ToProvisionedClusterIdentityPtrOutput() ProvisionedClusterIdentityPtrOutput
	ToProvisionedClusterIdentityPtrOutputWithContext(context.Context) ProvisionedClusterIdentityPtrOutput
}

type provisionedClusterIdentityPtrType ProvisionedClusterIdentityArgs

func ProvisionedClusterIdentityPtr(v *ProvisionedClusterIdentityArgs) ProvisionedClusterIdentityPtrInput {
	return (*provisionedClusterIdentityPtrType)(v)
}

func (*provisionedClusterIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedClusterIdentity)(nil)).Elem()
}

func (i *provisionedClusterIdentityPtrType) ToProvisionedClusterIdentityPtrOutput() ProvisionedClusterIdentityPtrOutput {
	return i.ToProvisionedClusterIdentityPtrOutputWithContext(context.Background())
}

func (i *provisionedClusterIdentityPtrType) ToProvisionedClusterIdentityPtrOutputWithContext(ctx context.Context) ProvisionedClusterIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionedClusterIdentityPtrOutput)
}

type ProvisionedClusterIdentityOutput struct{ *pulumi.OutputState }

func (ProvisionedClusterIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisionedClusterIdentity)(nil)).Elem()
}

func (o ProvisionedClusterIdentityOutput) ToProvisionedClusterIdentityOutput() ProvisionedClusterIdentityOutput {
	return o
}

func (o ProvisionedClusterIdentityOutput) ToProvisionedClusterIdentityOutputWithContext(ctx context.Context) ProvisionedClusterIdentityOutput {
	return o
}

func (o ProvisionedClusterIdentityOutput) ToProvisionedClusterIdentityPtrOutput() ProvisionedClusterIdentityPtrOutput {
	return o.ToProvisionedClusterIdentityPtrOutputWithContext(context.Background())
}

func (o ProvisionedClusterIdentityOutput) ToProvisionedClusterIdentityPtrOutputWithContext(ctx context.Context) ProvisionedClusterIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProvisionedClusterIdentity) *ProvisionedClusterIdentity {
		return &v
	}).(ProvisionedClusterIdentityPtrOutput)
}

func (o ProvisionedClusterIdentityOutput) Type() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v ProvisionedClusterIdentity) ResourceIdentityType { return v.Type }).(ResourceIdentityTypeOutput)
}

type ProvisionedClusterIdentityPtrOutput struct{ *pulumi.OutputState }

func (ProvisionedClusterIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedClusterIdentity)(nil)).Elem()
}

func (o ProvisionedClusterIdentityPtrOutput) ToProvisionedClusterIdentityPtrOutput() ProvisionedClusterIdentityPtrOutput {
	return o
}

func (o ProvisionedClusterIdentityPtrOutput) ToProvisionedClusterIdentityPtrOutputWithContext(ctx context.Context) ProvisionedClusterIdentityPtrOutput {
	return o
}

func (o ProvisionedClusterIdentityPtrOutput) Elem() ProvisionedClusterIdentityOutput {
	return o.ApplyT(func(v *ProvisionedClusterIdentity) ProvisionedClusterIdentity {
		if v != nil {
			return *v
		}
		var ret ProvisionedClusterIdentity
		return ret
	}).(ProvisionedClusterIdentityOutput)
}

func (o ProvisionedClusterIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *ProvisionedClusterIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type ProvisionedClusterIdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}

type ProvisionedClusterIdentityResponseOutput struct{ *pulumi.OutputState }

func (ProvisionedClusterIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisionedClusterIdentityResponse)(nil)).Elem()
}

func (o ProvisionedClusterIdentityResponseOutput) ToProvisionedClusterIdentityResponseOutput() ProvisionedClusterIdentityResponseOutput {
	return o
}

func (o ProvisionedClusterIdentityResponseOutput) ToProvisionedClusterIdentityResponseOutputWithContext(ctx context.Context) ProvisionedClusterIdentityResponseOutput {
	return o
}

func (o ProvisionedClusterIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ProvisionedClusterIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ProvisionedClusterIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ProvisionedClusterIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ProvisionedClusterIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ProvisionedClusterIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ProvisionedClusterIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ProvisionedClusterIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedClusterIdentityResponse)(nil)).Elem()
}

func (o ProvisionedClusterIdentityResponsePtrOutput) ToProvisionedClusterIdentityResponsePtrOutput() ProvisionedClusterIdentityResponsePtrOutput {
	return o
}

func (o ProvisionedClusterIdentityResponsePtrOutput) ToProvisionedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ProvisionedClusterIdentityResponsePtrOutput {
	return o
}

func (o ProvisionedClusterIdentityResponsePtrOutput) Elem() ProvisionedClusterIdentityResponseOutput {
	return o.ApplyT(func(v *ProvisionedClusterIdentityResponse) ProvisionedClusterIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ProvisionedClusterIdentityResponse
		return ret
	}).(ProvisionedClusterIdentityResponseOutput)
}

func (o ProvisionedClusterIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionedClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ProvisionedClusterIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionedClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ProvisionedClusterIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionedClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ProvisionedClustersAllProperties struct {
	AadProfile           *AADProfile                                  `pulumi:"aadProfile"`
	AddonProfiles        map[string]AddonProfiles                     `pulumi:"addonProfiles"`
	AgentPoolProfiles    []NamedAgentPoolProfile                      `pulumi:"agentPoolProfiles"`
	CloudProviderProfile *CloudProviderProfile                        `pulumi:"cloudProviderProfile"`
	ControlPlane         *ControlPlaneProfile                         `pulumi:"controlPlane"`
	EnableRbac           *bool                                        `pulumi:"enableRbac"`
	Features             *ProvisionedClustersCommonPropertiesFeatures `pulumi:"features"`
	HttpProxyConfig      *HttpProxyConfig                             `pulumi:"httpProxyConfig"`
	KubernetesVersion    *string                                      `pulumi:"kubernetesVersion"`
	LinuxProfile         *LinuxProfileProperties                      `pulumi:"linuxProfile"`
	NetworkProfile       *NetworkProfile                              `pulumi:"networkProfile"`
	NodeResourceGroup    *string                                      `pulumi:"nodeResourceGroup"`
	WindowsProfile       *WindowsProfile                              `pulumi:"windowsProfile"`
}


func (val *ProvisionedClustersAllProperties) Defaults() *ProvisionedClustersAllProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ControlPlane = tmp.ControlPlane.Defaults()

	tmp.Features = tmp.Features.Defaults()

	tmp.NetworkProfile = tmp.NetworkProfile.Defaults()

	return &tmp
}





type ProvisionedClustersAllPropertiesInput interface {
	pulumi.Input

	ToProvisionedClustersAllPropertiesOutput() ProvisionedClustersAllPropertiesOutput
	ToProvisionedClustersAllPropertiesOutputWithContext(context.Context) ProvisionedClustersAllPropertiesOutput
}

type ProvisionedClustersAllPropertiesArgs struct {
	AadProfile           AADProfilePtrInput                                  `pulumi:"aadProfile"`
	AddonProfiles        AddonProfilesMapInput                               `pulumi:"addonProfiles"`
	AgentPoolProfiles    NamedAgentPoolProfileArrayInput                     `pulumi:"agentPoolProfiles"`
	CloudProviderProfile CloudProviderProfilePtrInput                        `pulumi:"cloudProviderProfile"`
	ControlPlane         ControlPlaneProfilePtrInput                         `pulumi:"controlPlane"`
	EnableRbac           pulumi.BoolPtrInput                                 `pulumi:"enableRbac"`
	Features             ProvisionedClustersCommonPropertiesFeaturesPtrInput `pulumi:"features"`
	HttpProxyConfig      HttpProxyConfigPtrInput                             `pulumi:"httpProxyConfig"`
	KubernetesVersion    pulumi.StringPtrInput                               `pulumi:"kubernetesVersion"`
	LinuxProfile         LinuxProfilePropertiesPtrInput                      `pulumi:"linuxProfile"`
	NetworkProfile       NetworkProfilePtrInput                              `pulumi:"networkProfile"`
	NodeResourceGroup    pulumi.StringPtrInput                               `pulumi:"nodeResourceGroup"`
	WindowsProfile       WindowsProfilePtrInput                              `pulumi:"windowsProfile"`
}


func (val *ProvisionedClustersAllPropertiesArgs) Defaults() *ProvisionedClustersAllPropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	return &tmp
}
func (ProvisionedClustersAllPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisionedClustersAllProperties)(nil)).Elem()
}

func (i ProvisionedClustersAllPropertiesArgs) ToProvisionedClustersAllPropertiesOutput() ProvisionedClustersAllPropertiesOutput {
	return i.ToProvisionedClustersAllPropertiesOutputWithContext(context.Background())
}

func (i ProvisionedClustersAllPropertiesArgs) ToProvisionedClustersAllPropertiesOutputWithContext(ctx context.Context) ProvisionedClustersAllPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionedClustersAllPropertiesOutput)
}

func (i ProvisionedClustersAllPropertiesArgs) ToProvisionedClustersAllPropertiesPtrOutput() ProvisionedClustersAllPropertiesPtrOutput {
	return i.ToProvisionedClustersAllPropertiesPtrOutputWithContext(context.Background())
}

func (i ProvisionedClustersAllPropertiesArgs) ToProvisionedClustersAllPropertiesPtrOutputWithContext(ctx context.Context) ProvisionedClustersAllPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionedClustersAllPropertiesOutput).ToProvisionedClustersAllPropertiesPtrOutputWithContext(ctx)
}









type ProvisionedClustersAllPropertiesPtrInput interface {
	pulumi.Input

	ToProvisionedClustersAllPropertiesPtrOutput() ProvisionedClustersAllPropertiesPtrOutput
	ToProvisionedClustersAllPropertiesPtrOutputWithContext(context.Context) ProvisionedClustersAllPropertiesPtrOutput
}

type provisionedClustersAllPropertiesPtrType ProvisionedClustersAllPropertiesArgs

func ProvisionedClustersAllPropertiesPtr(v *ProvisionedClustersAllPropertiesArgs) ProvisionedClustersAllPropertiesPtrInput {
	return (*provisionedClustersAllPropertiesPtrType)(v)
}

func (*provisionedClustersAllPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedClustersAllProperties)(nil)).Elem()
}

func (i *provisionedClustersAllPropertiesPtrType) ToProvisionedClustersAllPropertiesPtrOutput() ProvisionedClustersAllPropertiesPtrOutput {
	return i.ToProvisionedClustersAllPropertiesPtrOutputWithContext(context.Background())
}

func (i *provisionedClustersAllPropertiesPtrType) ToProvisionedClustersAllPropertiesPtrOutputWithContext(ctx context.Context) ProvisionedClustersAllPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionedClustersAllPropertiesPtrOutput)
}

type ProvisionedClustersAllPropertiesOutput struct{ *pulumi.OutputState }

func (ProvisionedClustersAllPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisionedClustersAllProperties)(nil)).Elem()
}

func (o ProvisionedClustersAllPropertiesOutput) ToProvisionedClustersAllPropertiesOutput() ProvisionedClustersAllPropertiesOutput {
	return o
}

func (o ProvisionedClustersAllPropertiesOutput) ToProvisionedClustersAllPropertiesOutputWithContext(ctx context.Context) ProvisionedClustersAllPropertiesOutput {
	return o
}

func (o ProvisionedClustersAllPropertiesOutput) ToProvisionedClustersAllPropertiesPtrOutput() ProvisionedClustersAllPropertiesPtrOutput {
	return o.ToProvisionedClustersAllPropertiesPtrOutputWithContext(context.Background())
}

func (o ProvisionedClustersAllPropertiesOutput) ToProvisionedClustersAllPropertiesPtrOutputWithContext(ctx context.Context) ProvisionedClustersAllPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProvisionedClustersAllProperties) *ProvisionedClustersAllProperties {
		return &v
	}).(ProvisionedClustersAllPropertiesPtrOutput)
}

func (o ProvisionedClustersAllPropertiesOutput) AadProfile() AADProfilePtrOutput {
	return o.ApplyT(func(v ProvisionedClustersAllProperties) *AADProfile { return v.AadProfile }).(AADProfilePtrOutput)
}

func (o ProvisionedClustersAllPropertiesOutput) AddonProfiles() AddonProfilesMapOutput {
	return o.ApplyT(func(v ProvisionedClustersAllProperties) map[string]AddonProfiles { return v.AddonProfiles }).(AddonProfilesMapOutput)
}

func (o ProvisionedClustersAllPropertiesOutput) AgentPoolProfiles() NamedAgentPoolProfileArrayOutput {
	return o.ApplyT(func(v ProvisionedClustersAllProperties) []NamedAgentPoolProfile { return v.AgentPoolProfiles }).(NamedAgentPoolProfileArrayOutput)
}

func (o ProvisionedClustersAllPropertiesOutput) CloudProviderProfile() CloudProviderProfilePtrOutput {
	return o.ApplyT(func(v ProvisionedClustersAllProperties) *CloudProviderProfile { return v.CloudProviderProfile }).(CloudProviderProfilePtrOutput)
}

func (o ProvisionedClustersAllPropertiesOutput) ControlPlane() ControlPlaneProfilePtrOutput {
	return o.ApplyT(func(v ProvisionedClustersAllProperties) *ControlPlaneProfile { return v.ControlPlane }).(ControlPlaneProfilePtrOutput)
}

func (o ProvisionedClustersAllPropertiesOutput) EnableRbac() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersAllProperties) *bool { return v.EnableRbac }).(pulumi.BoolPtrOutput)
}

func (o ProvisionedClustersAllPropertiesOutput) Features() ProvisionedClustersCommonPropertiesFeaturesPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersAllProperties) *ProvisionedClustersCommonPropertiesFeatures {
		return v.Features
	}).(ProvisionedClustersCommonPropertiesFeaturesPtrOutput)
}

func (o ProvisionedClustersAllPropertiesOutput) HttpProxyConfig() HttpProxyConfigPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersAllProperties) *HttpProxyConfig { return v.HttpProxyConfig }).(HttpProxyConfigPtrOutput)
}

func (o ProvisionedClustersAllPropertiesOutput) KubernetesVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersAllProperties) *string { return v.KubernetesVersion }).(pulumi.StringPtrOutput)
}

func (o ProvisionedClustersAllPropertiesOutput) LinuxProfile() LinuxProfilePropertiesPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersAllProperties) *LinuxProfileProperties { return v.LinuxProfile }).(LinuxProfilePropertiesPtrOutput)
}

func (o ProvisionedClustersAllPropertiesOutput) NetworkProfile() NetworkProfilePtrOutput {
	return o.ApplyT(func(v ProvisionedClustersAllProperties) *NetworkProfile { return v.NetworkProfile }).(NetworkProfilePtrOutput)
}

func (o ProvisionedClustersAllPropertiesOutput) NodeResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersAllProperties) *string { return v.NodeResourceGroup }).(pulumi.StringPtrOutput)
}

func (o ProvisionedClustersAllPropertiesOutput) WindowsProfile() WindowsProfilePtrOutput {
	return o.ApplyT(func(v ProvisionedClustersAllProperties) *WindowsProfile { return v.WindowsProfile }).(WindowsProfilePtrOutput)
}

type ProvisionedClustersAllPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ProvisionedClustersAllPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedClustersAllProperties)(nil)).Elem()
}

func (o ProvisionedClustersAllPropertiesPtrOutput) ToProvisionedClustersAllPropertiesPtrOutput() ProvisionedClustersAllPropertiesPtrOutput {
	return o
}

func (o ProvisionedClustersAllPropertiesPtrOutput) ToProvisionedClustersAllPropertiesPtrOutputWithContext(ctx context.Context) ProvisionedClustersAllPropertiesPtrOutput {
	return o
}

func (o ProvisionedClustersAllPropertiesPtrOutput) Elem() ProvisionedClustersAllPropertiesOutput {
	return o.ApplyT(func(v *ProvisionedClustersAllProperties) ProvisionedClustersAllProperties {
		if v != nil {
			return *v
		}
		var ret ProvisionedClustersAllProperties
		return ret
	}).(ProvisionedClustersAllPropertiesOutput)
}

func (o ProvisionedClustersAllPropertiesPtrOutput) AadProfile() AADProfilePtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersAllProperties) *AADProfile {
		if v == nil {
			return nil
		}
		return v.AadProfile
	}).(AADProfilePtrOutput)
}

func (o ProvisionedClustersAllPropertiesPtrOutput) AddonProfiles() AddonProfilesMapOutput {
	return o.ApplyT(func(v *ProvisionedClustersAllProperties) map[string]AddonProfiles {
		if v == nil {
			return nil
		}
		return v.AddonProfiles
	}).(AddonProfilesMapOutput)
}

func (o ProvisionedClustersAllPropertiesPtrOutput) AgentPoolProfiles() NamedAgentPoolProfileArrayOutput {
	return o.ApplyT(func(v *ProvisionedClustersAllProperties) []NamedAgentPoolProfile {
		if v == nil {
			return nil
		}
		return v.AgentPoolProfiles
	}).(NamedAgentPoolProfileArrayOutput)
}

func (o ProvisionedClustersAllPropertiesPtrOutput) CloudProviderProfile() CloudProviderProfilePtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersAllProperties) *CloudProviderProfile {
		if v == nil {
			return nil
		}
		return v.CloudProviderProfile
	}).(CloudProviderProfilePtrOutput)
}

func (o ProvisionedClustersAllPropertiesPtrOutput) ControlPlane() ControlPlaneProfilePtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersAllProperties) *ControlPlaneProfile {
		if v == nil {
			return nil
		}
		return v.ControlPlane
	}).(ControlPlaneProfilePtrOutput)
}

func (o ProvisionedClustersAllPropertiesPtrOutput) EnableRbac() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersAllProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnableRbac
	}).(pulumi.BoolPtrOutput)
}

func (o ProvisionedClustersAllPropertiesPtrOutput) Features() ProvisionedClustersCommonPropertiesFeaturesPtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersAllProperties) *ProvisionedClustersCommonPropertiesFeatures {
		if v == nil {
			return nil
		}
		return v.Features
	}).(ProvisionedClustersCommonPropertiesFeaturesPtrOutput)
}

func (o ProvisionedClustersAllPropertiesPtrOutput) HttpProxyConfig() HttpProxyConfigPtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersAllProperties) *HttpProxyConfig {
		if v == nil {
			return nil
		}
		return v.HttpProxyConfig
	}).(HttpProxyConfigPtrOutput)
}

func (o ProvisionedClustersAllPropertiesPtrOutput) KubernetesVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersAllProperties) *string {
		if v == nil {
			return nil
		}
		return v.KubernetesVersion
	}).(pulumi.StringPtrOutput)
}

func (o ProvisionedClustersAllPropertiesPtrOutput) LinuxProfile() LinuxProfilePropertiesPtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersAllProperties) *LinuxProfileProperties {
		if v == nil {
			return nil
		}
		return v.LinuxProfile
	}).(LinuxProfilePropertiesPtrOutput)
}

func (o ProvisionedClustersAllPropertiesPtrOutput) NetworkProfile() NetworkProfilePtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersAllProperties) *NetworkProfile {
		if v == nil {
			return nil
		}
		return v.NetworkProfile
	}).(NetworkProfilePtrOutput)
}

func (o ProvisionedClustersAllPropertiesPtrOutput) NodeResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersAllProperties) *string {
		if v == nil {
			return nil
		}
		return v.NodeResourceGroup
	}).(pulumi.StringPtrOutput)
}

func (o ProvisionedClustersAllPropertiesPtrOutput) WindowsProfile() WindowsProfilePtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersAllProperties) *WindowsProfile {
		if v == nil {
			return nil
		}
		return v.WindowsProfile
	}).(WindowsProfilePtrOutput)
}

type ProvisionedClustersCommonPropertiesFeatures struct {
	ArcAgentProfile *ArcAgentProfile `pulumi:"arcAgentProfile"`
}


func (val *ProvisionedClustersCommonPropertiesFeatures) Defaults() *ProvisionedClustersCommonPropertiesFeatures {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ArcAgentProfile = tmp.ArcAgentProfile.Defaults()

	return &tmp
}





type ProvisionedClustersCommonPropertiesFeaturesInput interface {
	pulumi.Input

	ToProvisionedClustersCommonPropertiesFeaturesOutput() ProvisionedClustersCommonPropertiesFeaturesOutput
	ToProvisionedClustersCommonPropertiesFeaturesOutputWithContext(context.Context) ProvisionedClustersCommonPropertiesFeaturesOutput
}

type ProvisionedClustersCommonPropertiesFeaturesArgs struct {
	ArcAgentProfile ArcAgentProfilePtrInput `pulumi:"arcAgentProfile"`
}


func (val *ProvisionedClustersCommonPropertiesFeaturesArgs) Defaults() *ProvisionedClustersCommonPropertiesFeaturesArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	return &tmp
}
func (ProvisionedClustersCommonPropertiesFeaturesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisionedClustersCommonPropertiesFeatures)(nil)).Elem()
}

func (i ProvisionedClustersCommonPropertiesFeaturesArgs) ToProvisionedClustersCommonPropertiesFeaturesOutput() ProvisionedClustersCommonPropertiesFeaturesOutput {
	return i.ToProvisionedClustersCommonPropertiesFeaturesOutputWithContext(context.Background())
}

func (i ProvisionedClustersCommonPropertiesFeaturesArgs) ToProvisionedClustersCommonPropertiesFeaturesOutputWithContext(ctx context.Context) ProvisionedClustersCommonPropertiesFeaturesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionedClustersCommonPropertiesFeaturesOutput)
}

func (i ProvisionedClustersCommonPropertiesFeaturesArgs) ToProvisionedClustersCommonPropertiesFeaturesPtrOutput() ProvisionedClustersCommonPropertiesFeaturesPtrOutput {
	return i.ToProvisionedClustersCommonPropertiesFeaturesPtrOutputWithContext(context.Background())
}

func (i ProvisionedClustersCommonPropertiesFeaturesArgs) ToProvisionedClustersCommonPropertiesFeaturesPtrOutputWithContext(ctx context.Context) ProvisionedClustersCommonPropertiesFeaturesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionedClustersCommonPropertiesFeaturesOutput).ToProvisionedClustersCommonPropertiesFeaturesPtrOutputWithContext(ctx)
}









type ProvisionedClustersCommonPropertiesFeaturesPtrInput interface {
	pulumi.Input

	ToProvisionedClustersCommonPropertiesFeaturesPtrOutput() ProvisionedClustersCommonPropertiesFeaturesPtrOutput
	ToProvisionedClustersCommonPropertiesFeaturesPtrOutputWithContext(context.Context) ProvisionedClustersCommonPropertiesFeaturesPtrOutput
}

type provisionedClustersCommonPropertiesFeaturesPtrType ProvisionedClustersCommonPropertiesFeaturesArgs

func ProvisionedClustersCommonPropertiesFeaturesPtr(v *ProvisionedClustersCommonPropertiesFeaturesArgs) ProvisionedClustersCommonPropertiesFeaturesPtrInput {
	return (*provisionedClustersCommonPropertiesFeaturesPtrType)(v)
}

func (*provisionedClustersCommonPropertiesFeaturesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedClustersCommonPropertiesFeatures)(nil)).Elem()
}

func (i *provisionedClustersCommonPropertiesFeaturesPtrType) ToProvisionedClustersCommonPropertiesFeaturesPtrOutput() ProvisionedClustersCommonPropertiesFeaturesPtrOutput {
	return i.ToProvisionedClustersCommonPropertiesFeaturesPtrOutputWithContext(context.Background())
}

func (i *provisionedClustersCommonPropertiesFeaturesPtrType) ToProvisionedClustersCommonPropertiesFeaturesPtrOutputWithContext(ctx context.Context) ProvisionedClustersCommonPropertiesFeaturesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionedClustersCommonPropertiesFeaturesPtrOutput)
}

type ProvisionedClustersCommonPropertiesFeaturesOutput struct{ *pulumi.OutputState }

func (ProvisionedClustersCommonPropertiesFeaturesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisionedClustersCommonPropertiesFeatures)(nil)).Elem()
}

func (o ProvisionedClustersCommonPropertiesFeaturesOutput) ToProvisionedClustersCommonPropertiesFeaturesOutput() ProvisionedClustersCommonPropertiesFeaturesOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesFeaturesOutput) ToProvisionedClustersCommonPropertiesFeaturesOutputWithContext(ctx context.Context) ProvisionedClustersCommonPropertiesFeaturesOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesFeaturesOutput) ToProvisionedClustersCommonPropertiesFeaturesPtrOutput() ProvisionedClustersCommonPropertiesFeaturesPtrOutput {
	return o.ToProvisionedClustersCommonPropertiesFeaturesPtrOutputWithContext(context.Background())
}

func (o ProvisionedClustersCommonPropertiesFeaturesOutput) ToProvisionedClustersCommonPropertiesFeaturesPtrOutputWithContext(ctx context.Context) ProvisionedClustersCommonPropertiesFeaturesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProvisionedClustersCommonPropertiesFeatures) *ProvisionedClustersCommonPropertiesFeatures {
		return &v
	}).(ProvisionedClustersCommonPropertiesFeaturesPtrOutput)
}

func (o ProvisionedClustersCommonPropertiesFeaturesOutput) ArcAgentProfile() ArcAgentProfilePtrOutput {
	return o.ApplyT(func(v ProvisionedClustersCommonPropertiesFeatures) *ArcAgentProfile { return v.ArcAgentProfile }).(ArcAgentProfilePtrOutput)
}

type ProvisionedClustersCommonPropertiesFeaturesPtrOutput struct{ *pulumi.OutputState }

func (ProvisionedClustersCommonPropertiesFeaturesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedClustersCommonPropertiesFeatures)(nil)).Elem()
}

func (o ProvisionedClustersCommonPropertiesFeaturesPtrOutput) ToProvisionedClustersCommonPropertiesFeaturesPtrOutput() ProvisionedClustersCommonPropertiesFeaturesPtrOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesFeaturesPtrOutput) ToProvisionedClustersCommonPropertiesFeaturesPtrOutputWithContext(ctx context.Context) ProvisionedClustersCommonPropertiesFeaturesPtrOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesFeaturesPtrOutput) Elem() ProvisionedClustersCommonPropertiesFeaturesOutput {
	return o.ApplyT(func(v *ProvisionedClustersCommonPropertiesFeatures) ProvisionedClustersCommonPropertiesFeatures {
		if v != nil {
			return *v
		}
		var ret ProvisionedClustersCommonPropertiesFeatures
		return ret
	}).(ProvisionedClustersCommonPropertiesFeaturesOutput)
}

func (o ProvisionedClustersCommonPropertiesFeaturesPtrOutput) ArcAgentProfile() ArcAgentProfilePtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersCommonPropertiesFeatures) *ArcAgentProfile {
		if v == nil {
			return nil
		}
		return v.ArcAgentProfile
	}).(ArcAgentProfilePtrOutput)
}

type ProvisionedClustersCommonPropertiesResponseError struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}

type ProvisionedClustersCommonPropertiesResponseErrorOutput struct{ *pulumi.OutputState }

func (ProvisionedClustersCommonPropertiesResponseErrorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisionedClustersCommonPropertiesResponseError)(nil)).Elem()
}

func (o ProvisionedClustersCommonPropertiesResponseErrorOutput) ToProvisionedClustersCommonPropertiesResponseErrorOutput() ProvisionedClustersCommonPropertiesResponseErrorOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesResponseErrorOutput) ToProvisionedClustersCommonPropertiesResponseErrorOutputWithContext(ctx context.Context) ProvisionedClustersCommonPropertiesResponseErrorOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesResponseErrorOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersCommonPropertiesResponseError) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ProvisionedClustersCommonPropertiesResponseErrorOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersCommonPropertiesResponseError) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type ProvisionedClustersCommonPropertiesResponseErrorPtrOutput struct{ *pulumi.OutputState }

func (ProvisionedClustersCommonPropertiesResponseErrorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedClustersCommonPropertiesResponseError)(nil)).Elem()
}

func (o ProvisionedClustersCommonPropertiesResponseErrorPtrOutput) ToProvisionedClustersCommonPropertiesResponseErrorPtrOutput() ProvisionedClustersCommonPropertiesResponseErrorPtrOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesResponseErrorPtrOutput) ToProvisionedClustersCommonPropertiesResponseErrorPtrOutputWithContext(ctx context.Context) ProvisionedClustersCommonPropertiesResponseErrorPtrOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesResponseErrorPtrOutput) Elem() ProvisionedClustersCommonPropertiesResponseErrorOutput {
	return o.ApplyT(func(v *ProvisionedClustersCommonPropertiesResponseError) ProvisionedClustersCommonPropertiesResponseError {
		if v != nil {
			return *v
		}
		var ret ProvisionedClustersCommonPropertiesResponseError
		return ret
	}).(ProvisionedClustersCommonPropertiesResponseErrorOutput)
}

func (o ProvisionedClustersCommonPropertiesResponseErrorPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersCommonPropertiesResponseError) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ProvisionedClustersCommonPropertiesResponseErrorPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersCommonPropertiesResponseError) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

type ProvisionedClustersCommonPropertiesResponseFeatures struct {
	ArcAgentProfile *ArcAgentProfileResponse `pulumi:"arcAgentProfile"`
}


func (val *ProvisionedClustersCommonPropertiesResponseFeatures) Defaults() *ProvisionedClustersCommonPropertiesResponseFeatures {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ArcAgentProfile = tmp.ArcAgentProfile.Defaults()

	return &tmp
}

type ProvisionedClustersCommonPropertiesResponseFeaturesOutput struct{ *pulumi.OutputState }

func (ProvisionedClustersCommonPropertiesResponseFeaturesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisionedClustersCommonPropertiesResponseFeatures)(nil)).Elem()
}

func (o ProvisionedClustersCommonPropertiesResponseFeaturesOutput) ToProvisionedClustersCommonPropertiesResponseFeaturesOutput() ProvisionedClustersCommonPropertiesResponseFeaturesOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesResponseFeaturesOutput) ToProvisionedClustersCommonPropertiesResponseFeaturesOutputWithContext(ctx context.Context) ProvisionedClustersCommonPropertiesResponseFeaturesOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesResponseFeaturesOutput) ArcAgentProfile() ArcAgentProfileResponsePtrOutput {
	return o.ApplyT(func(v ProvisionedClustersCommonPropertiesResponseFeatures) *ArcAgentProfileResponse {
		return v.ArcAgentProfile
	}).(ArcAgentProfileResponsePtrOutput)
}

type ProvisionedClustersCommonPropertiesResponseFeaturesPtrOutput struct{ *pulumi.OutputState }

func (ProvisionedClustersCommonPropertiesResponseFeaturesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedClustersCommonPropertiesResponseFeatures)(nil)).Elem()
}

func (o ProvisionedClustersCommonPropertiesResponseFeaturesPtrOutput) ToProvisionedClustersCommonPropertiesResponseFeaturesPtrOutput() ProvisionedClustersCommonPropertiesResponseFeaturesPtrOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesResponseFeaturesPtrOutput) ToProvisionedClustersCommonPropertiesResponseFeaturesPtrOutputWithContext(ctx context.Context) ProvisionedClustersCommonPropertiesResponseFeaturesPtrOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesResponseFeaturesPtrOutput) Elem() ProvisionedClustersCommonPropertiesResponseFeaturesOutput {
	return o.ApplyT(func(v *ProvisionedClustersCommonPropertiesResponseFeatures) ProvisionedClustersCommonPropertiesResponseFeatures {
		if v != nil {
			return *v
		}
		var ret ProvisionedClustersCommonPropertiesResponseFeatures
		return ret
	}).(ProvisionedClustersCommonPropertiesResponseFeaturesOutput)
}

func (o ProvisionedClustersCommonPropertiesResponseFeaturesPtrOutput) ArcAgentProfile() ArcAgentProfileResponsePtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersCommonPropertiesResponseFeatures) *ArcAgentProfileResponse {
		if v == nil {
			return nil
		}
		return v.ArcAgentProfile
	}).(ArcAgentProfileResponsePtrOutput)
}

type ProvisionedClustersCommonPropertiesResponseFeaturesStatus struct {
	ArcAgentStatus *ArcAgentStatusResponse `pulumi:"arcAgentStatus"`
}

type ProvisionedClustersCommonPropertiesResponseFeaturesStatusOutput struct{ *pulumi.OutputState }

func (ProvisionedClustersCommonPropertiesResponseFeaturesStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisionedClustersCommonPropertiesResponseFeaturesStatus)(nil)).Elem()
}

func (o ProvisionedClustersCommonPropertiesResponseFeaturesStatusOutput) ToProvisionedClustersCommonPropertiesResponseFeaturesStatusOutput() ProvisionedClustersCommonPropertiesResponseFeaturesStatusOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesResponseFeaturesStatusOutput) ToProvisionedClustersCommonPropertiesResponseFeaturesStatusOutputWithContext(ctx context.Context) ProvisionedClustersCommonPropertiesResponseFeaturesStatusOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesResponseFeaturesStatusOutput) ArcAgentStatus() ArcAgentStatusResponsePtrOutput {
	return o.ApplyT(func(v ProvisionedClustersCommonPropertiesResponseFeaturesStatus) *ArcAgentStatusResponse {
		return v.ArcAgentStatus
	}).(ArcAgentStatusResponsePtrOutput)
}

type ProvisionedClustersCommonPropertiesResponseFeaturesStatusPtrOutput struct{ *pulumi.OutputState }

func (ProvisionedClustersCommonPropertiesResponseFeaturesStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedClustersCommonPropertiesResponseFeaturesStatus)(nil)).Elem()
}

func (o ProvisionedClustersCommonPropertiesResponseFeaturesStatusPtrOutput) ToProvisionedClustersCommonPropertiesResponseFeaturesStatusPtrOutput() ProvisionedClustersCommonPropertiesResponseFeaturesStatusPtrOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesResponseFeaturesStatusPtrOutput) ToProvisionedClustersCommonPropertiesResponseFeaturesStatusPtrOutputWithContext(ctx context.Context) ProvisionedClustersCommonPropertiesResponseFeaturesStatusPtrOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesResponseFeaturesStatusPtrOutput) Elem() ProvisionedClustersCommonPropertiesResponseFeaturesStatusOutput {
	return o.ApplyT(func(v *ProvisionedClustersCommonPropertiesResponseFeaturesStatus) ProvisionedClustersCommonPropertiesResponseFeaturesStatus {
		if v != nil {
			return *v
		}
		var ret ProvisionedClustersCommonPropertiesResponseFeaturesStatus
		return ret
	}).(ProvisionedClustersCommonPropertiesResponseFeaturesStatusOutput)
}

func (o ProvisionedClustersCommonPropertiesResponseFeaturesStatusPtrOutput) ArcAgentStatus() ArcAgentStatusResponsePtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersCommonPropertiesResponseFeaturesStatus) *ArcAgentStatusResponse {
		if v == nil {
			return nil
		}
		return v.ArcAgentStatus
	}).(ArcAgentStatusResponsePtrOutput)
}

type ProvisionedClustersCommonPropertiesResponseProvisioningStatus struct {
	Error       *ProvisionedClustersCommonPropertiesResponseError `pulumi:"error"`
	OperationId *string                                           `pulumi:"operationId"`
	Phase       *string                                           `pulumi:"phase"`
	Status      *string                                           `pulumi:"status"`
}

type ProvisionedClustersCommonPropertiesResponseProvisioningStatusOutput struct{ *pulumi.OutputState }

func (ProvisionedClustersCommonPropertiesResponseProvisioningStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisionedClustersCommonPropertiesResponseProvisioningStatus)(nil)).Elem()
}

func (o ProvisionedClustersCommonPropertiesResponseProvisioningStatusOutput) ToProvisionedClustersCommonPropertiesResponseProvisioningStatusOutput() ProvisionedClustersCommonPropertiesResponseProvisioningStatusOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesResponseProvisioningStatusOutput) ToProvisionedClustersCommonPropertiesResponseProvisioningStatusOutputWithContext(ctx context.Context) ProvisionedClustersCommonPropertiesResponseProvisioningStatusOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesResponseProvisioningStatusOutput) Error() ProvisionedClustersCommonPropertiesResponseErrorPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersCommonPropertiesResponseProvisioningStatus) *ProvisionedClustersCommonPropertiesResponseError {
		return v.Error
	}).(ProvisionedClustersCommonPropertiesResponseErrorPtrOutput)
}

func (o ProvisionedClustersCommonPropertiesResponseProvisioningStatusOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersCommonPropertiesResponseProvisioningStatus) *string { return v.OperationId }).(pulumi.StringPtrOutput)
}

func (o ProvisionedClustersCommonPropertiesResponseProvisioningStatusOutput) Phase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersCommonPropertiesResponseProvisioningStatus) *string { return v.Phase }).(pulumi.StringPtrOutput)
}

func (o ProvisionedClustersCommonPropertiesResponseProvisioningStatusOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersCommonPropertiesResponseProvisioningStatus) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ProvisionedClustersCommonPropertiesResponseProvisioningStatusPtrOutput struct{ *pulumi.OutputState }

func (ProvisionedClustersCommonPropertiesResponseProvisioningStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedClustersCommonPropertiesResponseProvisioningStatus)(nil)).Elem()
}

func (o ProvisionedClustersCommonPropertiesResponseProvisioningStatusPtrOutput) ToProvisionedClustersCommonPropertiesResponseProvisioningStatusPtrOutput() ProvisionedClustersCommonPropertiesResponseProvisioningStatusPtrOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesResponseProvisioningStatusPtrOutput) ToProvisionedClustersCommonPropertiesResponseProvisioningStatusPtrOutputWithContext(ctx context.Context) ProvisionedClustersCommonPropertiesResponseProvisioningStatusPtrOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesResponseProvisioningStatusPtrOutput) Elem() ProvisionedClustersCommonPropertiesResponseProvisioningStatusOutput {
	return o.ApplyT(func(v *ProvisionedClustersCommonPropertiesResponseProvisioningStatus) ProvisionedClustersCommonPropertiesResponseProvisioningStatus {
		if v != nil {
			return *v
		}
		var ret ProvisionedClustersCommonPropertiesResponseProvisioningStatus
		return ret
	}).(ProvisionedClustersCommonPropertiesResponseProvisioningStatusOutput)
}

func (o ProvisionedClustersCommonPropertiesResponseProvisioningStatusPtrOutput) Error() ProvisionedClustersCommonPropertiesResponseErrorPtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersCommonPropertiesResponseProvisioningStatus) *ProvisionedClustersCommonPropertiesResponseError {
		if v == nil {
			return nil
		}
		return v.Error
	}).(ProvisionedClustersCommonPropertiesResponseErrorPtrOutput)
}

func (o ProvisionedClustersCommonPropertiesResponseProvisioningStatusPtrOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersCommonPropertiesResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.OperationId
	}).(pulumi.StringPtrOutput)
}

func (o ProvisionedClustersCommonPropertiesResponseProvisioningStatusPtrOutput) Phase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersCommonPropertiesResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.Phase
	}).(pulumi.StringPtrOutput)
}

func (o ProvisionedClustersCommonPropertiesResponseProvisioningStatusPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersCommonPropertiesResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ProvisionedClustersCommonPropertiesResponseStatus struct {
	AddonStatus        map[string]AddonStatusResponse                                 `pulumi:"addonStatus"`
	ErrorMessage       *string                                                        `pulumi:"errorMessage"`
	FeaturesStatus     *ProvisionedClustersCommonPropertiesResponseFeaturesStatus     `pulumi:"featuresStatus"`
	ProvisioningStatus *ProvisionedClustersCommonPropertiesResponseProvisioningStatus `pulumi:"provisioningStatus"`
}

type ProvisionedClustersCommonPropertiesResponseStatusOutput struct{ *pulumi.OutputState }

func (ProvisionedClustersCommonPropertiesResponseStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisionedClustersCommonPropertiesResponseStatus)(nil)).Elem()
}

func (o ProvisionedClustersCommonPropertiesResponseStatusOutput) ToProvisionedClustersCommonPropertiesResponseStatusOutput() ProvisionedClustersCommonPropertiesResponseStatusOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesResponseStatusOutput) ToProvisionedClustersCommonPropertiesResponseStatusOutputWithContext(ctx context.Context) ProvisionedClustersCommonPropertiesResponseStatusOutput {
	return o
}

func (o ProvisionedClustersCommonPropertiesResponseStatusOutput) AddonStatus() AddonStatusResponseMapOutput {
	return o.ApplyT(func(v ProvisionedClustersCommonPropertiesResponseStatus) map[string]AddonStatusResponse {
		return v.AddonStatus
	}).(AddonStatusResponseMapOutput)
}

func (o ProvisionedClustersCommonPropertiesResponseStatusOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersCommonPropertiesResponseStatus) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o ProvisionedClustersCommonPropertiesResponseStatusOutput) FeaturesStatus() ProvisionedClustersCommonPropertiesResponseFeaturesStatusPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersCommonPropertiesResponseStatus) *ProvisionedClustersCommonPropertiesResponseFeaturesStatus {
		return v.FeaturesStatus
	}).(ProvisionedClustersCommonPropertiesResponseFeaturesStatusPtrOutput)
}

func (o ProvisionedClustersCommonPropertiesResponseStatusOutput) ProvisioningStatus() ProvisionedClustersCommonPropertiesResponseProvisioningStatusPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersCommonPropertiesResponseStatus) *ProvisionedClustersCommonPropertiesResponseProvisioningStatus {
		return v.ProvisioningStatus
	}).(ProvisionedClustersCommonPropertiesResponseProvisioningStatusPtrOutput)
}

type ProvisionedClustersExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ProvisionedClustersExtendedLocationInput interface {
	pulumi.Input

	ToProvisionedClustersExtendedLocationOutput() ProvisionedClustersExtendedLocationOutput
	ToProvisionedClustersExtendedLocationOutputWithContext(context.Context) ProvisionedClustersExtendedLocationOutput
}

type ProvisionedClustersExtendedLocationArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ProvisionedClustersExtendedLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisionedClustersExtendedLocation)(nil)).Elem()
}

func (i ProvisionedClustersExtendedLocationArgs) ToProvisionedClustersExtendedLocationOutput() ProvisionedClustersExtendedLocationOutput {
	return i.ToProvisionedClustersExtendedLocationOutputWithContext(context.Background())
}

func (i ProvisionedClustersExtendedLocationArgs) ToProvisionedClustersExtendedLocationOutputWithContext(ctx context.Context) ProvisionedClustersExtendedLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionedClustersExtendedLocationOutput)
}

func (i ProvisionedClustersExtendedLocationArgs) ToProvisionedClustersExtendedLocationPtrOutput() ProvisionedClustersExtendedLocationPtrOutput {
	return i.ToProvisionedClustersExtendedLocationPtrOutputWithContext(context.Background())
}

func (i ProvisionedClustersExtendedLocationArgs) ToProvisionedClustersExtendedLocationPtrOutputWithContext(ctx context.Context) ProvisionedClustersExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionedClustersExtendedLocationOutput).ToProvisionedClustersExtendedLocationPtrOutputWithContext(ctx)
}









type ProvisionedClustersExtendedLocationPtrInput interface {
	pulumi.Input

	ToProvisionedClustersExtendedLocationPtrOutput() ProvisionedClustersExtendedLocationPtrOutput
	ToProvisionedClustersExtendedLocationPtrOutputWithContext(context.Context) ProvisionedClustersExtendedLocationPtrOutput
}

type provisionedClustersExtendedLocationPtrType ProvisionedClustersExtendedLocationArgs

func ProvisionedClustersExtendedLocationPtr(v *ProvisionedClustersExtendedLocationArgs) ProvisionedClustersExtendedLocationPtrInput {
	return (*provisionedClustersExtendedLocationPtrType)(v)
}

func (*provisionedClustersExtendedLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedClustersExtendedLocation)(nil)).Elem()
}

func (i *provisionedClustersExtendedLocationPtrType) ToProvisionedClustersExtendedLocationPtrOutput() ProvisionedClustersExtendedLocationPtrOutput {
	return i.ToProvisionedClustersExtendedLocationPtrOutputWithContext(context.Background())
}

func (i *provisionedClustersExtendedLocationPtrType) ToProvisionedClustersExtendedLocationPtrOutputWithContext(ctx context.Context) ProvisionedClustersExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionedClustersExtendedLocationPtrOutput)
}

type ProvisionedClustersExtendedLocationOutput struct{ *pulumi.OutputState }

func (ProvisionedClustersExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisionedClustersExtendedLocation)(nil)).Elem()
}

func (o ProvisionedClustersExtendedLocationOutput) ToProvisionedClustersExtendedLocationOutput() ProvisionedClustersExtendedLocationOutput {
	return o
}

func (o ProvisionedClustersExtendedLocationOutput) ToProvisionedClustersExtendedLocationOutputWithContext(ctx context.Context) ProvisionedClustersExtendedLocationOutput {
	return o
}

func (o ProvisionedClustersExtendedLocationOutput) ToProvisionedClustersExtendedLocationPtrOutput() ProvisionedClustersExtendedLocationPtrOutput {
	return o.ToProvisionedClustersExtendedLocationPtrOutputWithContext(context.Background())
}

func (o ProvisionedClustersExtendedLocationOutput) ToProvisionedClustersExtendedLocationPtrOutputWithContext(ctx context.Context) ProvisionedClustersExtendedLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProvisionedClustersExtendedLocation) *ProvisionedClustersExtendedLocation {
		return &v
	}).(ProvisionedClustersExtendedLocationPtrOutput)
}

func (o ProvisionedClustersExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ProvisionedClustersExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ProvisionedClustersExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (ProvisionedClustersExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedClustersExtendedLocation)(nil)).Elem()
}

func (o ProvisionedClustersExtendedLocationPtrOutput) ToProvisionedClustersExtendedLocationPtrOutput() ProvisionedClustersExtendedLocationPtrOutput {
	return o
}

func (o ProvisionedClustersExtendedLocationPtrOutput) ToProvisionedClustersExtendedLocationPtrOutputWithContext(ctx context.Context) ProvisionedClustersExtendedLocationPtrOutput {
	return o
}

func (o ProvisionedClustersExtendedLocationPtrOutput) Elem() ProvisionedClustersExtendedLocationOutput {
	return o.ApplyT(func(v *ProvisionedClustersExtendedLocation) ProvisionedClustersExtendedLocation {
		if v != nil {
			return *v
		}
		var ret ProvisionedClustersExtendedLocation
		return ret
	}).(ProvisionedClustersExtendedLocationOutput)
}

func (o ProvisionedClustersExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ProvisionedClustersExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ProvisionedClustersResponsePropertiesResponse struct {
	AadProfile           *AADProfileResponseResponse                          `pulumi:"aadProfile"`
	AddonProfiles        map[string]AddonProfilesResponse                     `pulumi:"addonProfiles"`
	AgentPoolProfiles    []NamedAgentPoolProfileResponse                      `pulumi:"agentPoolProfiles"`
	CloudProviderProfile *CloudProviderProfileResponse                        `pulumi:"cloudProviderProfile"`
	ControlPlane         *ControlPlaneProfileResponse                         `pulumi:"controlPlane"`
	EnableRbac           *bool                                                `pulumi:"enableRbac"`
	Features             *ProvisionedClustersCommonPropertiesResponseFeatures `pulumi:"features"`
	HttpProxyConfig      *HttpProxyConfigResponseResponse                     `pulumi:"httpProxyConfig"`
	KubernetesVersion    *string                                              `pulumi:"kubernetesVersion"`
	LinuxProfile         *LinuxProfilePropertiesResponse                      `pulumi:"linuxProfile"`
	NetworkProfile       *NetworkProfileResponse                              `pulumi:"networkProfile"`
	NodeResourceGroup    *string                                              `pulumi:"nodeResourceGroup"`
	ProvisioningState    string                                               `pulumi:"provisioningState"`
	Status               ProvisionedClustersCommonPropertiesResponseStatus    `pulumi:"status"`
	WindowsProfile       *WindowsProfileResponseResponse                      `pulumi:"windowsProfile"`
}


func (val *ProvisionedClustersResponsePropertiesResponse) Defaults() *ProvisionedClustersResponsePropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ControlPlane = tmp.ControlPlane.Defaults()

	tmp.Features = tmp.Features.Defaults()

	tmp.NetworkProfile = tmp.NetworkProfile.Defaults()

	return &tmp
}

type ProvisionedClustersResponsePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ProvisionedClustersResponsePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisionedClustersResponsePropertiesResponse)(nil)).Elem()
}

func (o ProvisionedClustersResponsePropertiesResponseOutput) ToProvisionedClustersResponsePropertiesResponseOutput() ProvisionedClustersResponsePropertiesResponseOutput {
	return o
}

func (o ProvisionedClustersResponsePropertiesResponseOutput) ToProvisionedClustersResponsePropertiesResponseOutputWithContext(ctx context.Context) ProvisionedClustersResponsePropertiesResponseOutput {
	return o
}

func (o ProvisionedClustersResponsePropertiesResponseOutput) AadProfile() AADProfileResponseResponsePtrOutput {
	return o.ApplyT(func(v ProvisionedClustersResponsePropertiesResponse) *AADProfileResponseResponse { return v.AadProfile }).(AADProfileResponseResponsePtrOutput)
}

func (o ProvisionedClustersResponsePropertiesResponseOutput) AddonProfiles() AddonProfilesResponseMapOutput {
	return o.ApplyT(func(v ProvisionedClustersResponsePropertiesResponse) map[string]AddonProfilesResponse {
		return v.AddonProfiles
	}).(AddonProfilesResponseMapOutput)
}

func (o ProvisionedClustersResponsePropertiesResponseOutput) AgentPoolProfiles() NamedAgentPoolProfileResponseArrayOutput {
	return o.ApplyT(func(v ProvisionedClustersResponsePropertiesResponse) []NamedAgentPoolProfileResponse {
		return v.AgentPoolProfiles
	}).(NamedAgentPoolProfileResponseArrayOutput)
}

func (o ProvisionedClustersResponsePropertiesResponseOutput) CloudProviderProfile() CloudProviderProfileResponsePtrOutput {
	return o.ApplyT(func(v ProvisionedClustersResponsePropertiesResponse) *CloudProviderProfileResponse {
		return v.CloudProviderProfile
	}).(CloudProviderProfileResponsePtrOutput)
}

func (o ProvisionedClustersResponsePropertiesResponseOutput) ControlPlane() ControlPlaneProfileResponsePtrOutput {
	return o.ApplyT(func(v ProvisionedClustersResponsePropertiesResponse) *ControlPlaneProfileResponse {
		return v.ControlPlane
	}).(ControlPlaneProfileResponsePtrOutput)
}

func (o ProvisionedClustersResponsePropertiesResponseOutput) EnableRbac() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersResponsePropertiesResponse) *bool { return v.EnableRbac }).(pulumi.BoolPtrOutput)
}

func (o ProvisionedClustersResponsePropertiesResponseOutput) Features() ProvisionedClustersCommonPropertiesResponseFeaturesPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersResponsePropertiesResponse) *ProvisionedClustersCommonPropertiesResponseFeatures {
		return v.Features
	}).(ProvisionedClustersCommonPropertiesResponseFeaturesPtrOutput)
}

func (o ProvisionedClustersResponsePropertiesResponseOutput) HttpProxyConfig() HttpProxyConfigResponseResponsePtrOutput {
	return o.ApplyT(func(v ProvisionedClustersResponsePropertiesResponse) *HttpProxyConfigResponseResponse {
		return v.HttpProxyConfig
	}).(HttpProxyConfigResponseResponsePtrOutput)
}

func (o ProvisionedClustersResponsePropertiesResponseOutput) KubernetesVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersResponsePropertiesResponse) *string { return v.KubernetesVersion }).(pulumi.StringPtrOutput)
}

func (o ProvisionedClustersResponsePropertiesResponseOutput) LinuxProfile() LinuxProfilePropertiesResponsePtrOutput {
	return o.ApplyT(func(v ProvisionedClustersResponsePropertiesResponse) *LinuxProfilePropertiesResponse {
		return v.LinuxProfile
	}).(LinuxProfilePropertiesResponsePtrOutput)
}

func (o ProvisionedClustersResponsePropertiesResponseOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v ProvisionedClustersResponsePropertiesResponse) *NetworkProfileResponse { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o ProvisionedClustersResponsePropertiesResponseOutput) NodeResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersResponsePropertiesResponse) *string { return v.NodeResourceGroup }).(pulumi.StringPtrOutput)
}

func (o ProvisionedClustersResponsePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ProvisionedClustersResponsePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ProvisionedClustersResponsePropertiesResponseOutput) Status() ProvisionedClustersCommonPropertiesResponseStatusOutput {
	return o.ApplyT(func(v ProvisionedClustersResponsePropertiesResponse) ProvisionedClustersCommonPropertiesResponseStatus {
		return v.Status
	}).(ProvisionedClustersCommonPropertiesResponseStatusOutput)
}

func (o ProvisionedClustersResponsePropertiesResponseOutput) WindowsProfile() WindowsProfileResponseResponsePtrOutput {
	return o.ApplyT(func(v ProvisionedClustersResponsePropertiesResponse) *WindowsProfileResponseResponse {
		return v.WindowsProfile
	}).(WindowsProfileResponseResponsePtrOutput)
}

type ProvisionedClustersResponseResponseExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type ProvisionedClustersResponseResponseExtendedLocationOutput struct{ *pulumi.OutputState }

func (ProvisionedClustersResponseResponseExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisionedClustersResponseResponseExtendedLocation)(nil)).Elem()
}

func (o ProvisionedClustersResponseResponseExtendedLocationOutput) ToProvisionedClustersResponseResponseExtendedLocationOutput() ProvisionedClustersResponseResponseExtendedLocationOutput {
	return o
}

func (o ProvisionedClustersResponseResponseExtendedLocationOutput) ToProvisionedClustersResponseResponseExtendedLocationOutputWithContext(ctx context.Context) ProvisionedClustersResponseResponseExtendedLocationOutput {
	return o
}

func (o ProvisionedClustersResponseResponseExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersResponseResponseExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ProvisionedClustersResponseResponseExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisionedClustersResponseResponseExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ProvisionedClustersResponseResponseExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (ProvisionedClustersResponseResponseExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedClustersResponseResponseExtendedLocation)(nil)).Elem()
}

func (o ProvisionedClustersResponseResponseExtendedLocationPtrOutput) ToProvisionedClustersResponseResponseExtendedLocationPtrOutput() ProvisionedClustersResponseResponseExtendedLocationPtrOutput {
	return o
}

func (o ProvisionedClustersResponseResponseExtendedLocationPtrOutput) ToProvisionedClustersResponseResponseExtendedLocationPtrOutputWithContext(ctx context.Context) ProvisionedClustersResponseResponseExtendedLocationPtrOutput {
	return o
}

func (o ProvisionedClustersResponseResponseExtendedLocationPtrOutput) Elem() ProvisionedClustersResponseResponseExtendedLocationOutput {
	return o.ApplyT(func(v *ProvisionedClustersResponseResponseExtendedLocation) ProvisionedClustersResponseResponseExtendedLocation {
		if v != nil {
			return *v
		}
		var ret ProvisionedClustersResponseResponseExtendedLocation
		return ret
	}).(ProvisionedClustersResponseResponseExtendedLocationOutput)
}

func (o ProvisionedClustersResponseResponseExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersResponseResponseExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ProvisionedClustersResponseResponseExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionedClustersResponseResponseExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type StorageSpacesExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type StorageSpacesExtendedLocationInput interface {
	pulumi.Input

	ToStorageSpacesExtendedLocationOutput() StorageSpacesExtendedLocationOutput
	ToStorageSpacesExtendedLocationOutputWithContext(context.Context) StorageSpacesExtendedLocationOutput
}

type StorageSpacesExtendedLocationArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (StorageSpacesExtendedLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesExtendedLocation)(nil)).Elem()
}

func (i StorageSpacesExtendedLocationArgs) ToStorageSpacesExtendedLocationOutput() StorageSpacesExtendedLocationOutput {
	return i.ToStorageSpacesExtendedLocationOutputWithContext(context.Background())
}

func (i StorageSpacesExtendedLocationArgs) ToStorageSpacesExtendedLocationOutputWithContext(ctx context.Context) StorageSpacesExtendedLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesExtendedLocationOutput)
}

func (i StorageSpacesExtendedLocationArgs) ToStorageSpacesExtendedLocationPtrOutput() StorageSpacesExtendedLocationPtrOutput {
	return i.ToStorageSpacesExtendedLocationPtrOutputWithContext(context.Background())
}

func (i StorageSpacesExtendedLocationArgs) ToStorageSpacesExtendedLocationPtrOutputWithContext(ctx context.Context) StorageSpacesExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesExtendedLocationOutput).ToStorageSpacesExtendedLocationPtrOutputWithContext(ctx)
}









type StorageSpacesExtendedLocationPtrInput interface {
	pulumi.Input

	ToStorageSpacesExtendedLocationPtrOutput() StorageSpacesExtendedLocationPtrOutput
	ToStorageSpacesExtendedLocationPtrOutputWithContext(context.Context) StorageSpacesExtendedLocationPtrOutput
}

type storageSpacesExtendedLocationPtrType StorageSpacesExtendedLocationArgs

func StorageSpacesExtendedLocationPtr(v *StorageSpacesExtendedLocationArgs) StorageSpacesExtendedLocationPtrInput {
	return (*storageSpacesExtendedLocationPtrType)(v)
}

func (*storageSpacesExtendedLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpacesExtendedLocation)(nil)).Elem()
}

func (i *storageSpacesExtendedLocationPtrType) ToStorageSpacesExtendedLocationPtrOutput() StorageSpacesExtendedLocationPtrOutput {
	return i.ToStorageSpacesExtendedLocationPtrOutputWithContext(context.Background())
}

func (i *storageSpacesExtendedLocationPtrType) ToStorageSpacesExtendedLocationPtrOutputWithContext(ctx context.Context) StorageSpacesExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesExtendedLocationPtrOutput)
}

type StorageSpacesExtendedLocationOutput struct{ *pulumi.OutputState }

func (StorageSpacesExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesExtendedLocation)(nil)).Elem()
}

func (o StorageSpacesExtendedLocationOutput) ToStorageSpacesExtendedLocationOutput() StorageSpacesExtendedLocationOutput {
	return o
}

func (o StorageSpacesExtendedLocationOutput) ToStorageSpacesExtendedLocationOutputWithContext(ctx context.Context) StorageSpacesExtendedLocationOutput {
	return o
}

func (o StorageSpacesExtendedLocationOutput) ToStorageSpacesExtendedLocationPtrOutput() StorageSpacesExtendedLocationPtrOutput {
	return o.ToStorageSpacesExtendedLocationPtrOutputWithContext(context.Background())
}

func (o StorageSpacesExtendedLocationOutput) ToStorageSpacesExtendedLocationPtrOutputWithContext(ctx context.Context) StorageSpacesExtendedLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageSpacesExtendedLocation) *StorageSpacesExtendedLocation {
		return &v
	}).(StorageSpacesExtendedLocationPtrOutput)
}

func (o StorageSpacesExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o StorageSpacesExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type StorageSpacesExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (StorageSpacesExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpacesExtendedLocation)(nil)).Elem()
}

func (o StorageSpacesExtendedLocationPtrOutput) ToStorageSpacesExtendedLocationPtrOutput() StorageSpacesExtendedLocationPtrOutput {
	return o
}

func (o StorageSpacesExtendedLocationPtrOutput) ToStorageSpacesExtendedLocationPtrOutputWithContext(ctx context.Context) StorageSpacesExtendedLocationPtrOutput {
	return o
}

func (o StorageSpacesExtendedLocationPtrOutput) Elem() StorageSpacesExtendedLocationOutput {
	return o.ApplyT(func(v *StorageSpacesExtendedLocation) StorageSpacesExtendedLocation {
		if v != nil {
			return *v
		}
		var ret StorageSpacesExtendedLocation
		return ret
	}).(StorageSpacesExtendedLocationOutput)
}

func (o StorageSpacesExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o StorageSpacesExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type StorageSpacesProperties struct {
	HciStorageProfile    *StorageSpacesPropertiesHciStorageProfile    `pulumi:"hciStorageProfile"`
	Status               *StorageSpacesPropertiesStatus               `pulumi:"status"`
	VmwareStorageProfile *StorageSpacesPropertiesVmwareStorageProfile `pulumi:"vmwareStorageProfile"`
}





type StorageSpacesPropertiesInput interface {
	pulumi.Input

	ToStorageSpacesPropertiesOutput() StorageSpacesPropertiesOutput
	ToStorageSpacesPropertiesOutputWithContext(context.Context) StorageSpacesPropertiesOutput
}

type StorageSpacesPropertiesArgs struct {
	HciStorageProfile    StorageSpacesPropertiesHciStorageProfilePtrInput    `pulumi:"hciStorageProfile"`
	Status               StorageSpacesPropertiesStatusPtrInput               `pulumi:"status"`
	VmwareStorageProfile StorageSpacesPropertiesVmwareStorageProfilePtrInput `pulumi:"vmwareStorageProfile"`
}

func (StorageSpacesPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesProperties)(nil)).Elem()
}

func (i StorageSpacesPropertiesArgs) ToStorageSpacesPropertiesOutput() StorageSpacesPropertiesOutput {
	return i.ToStorageSpacesPropertiesOutputWithContext(context.Background())
}

func (i StorageSpacesPropertiesArgs) ToStorageSpacesPropertiesOutputWithContext(ctx context.Context) StorageSpacesPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesPropertiesOutput)
}

func (i StorageSpacesPropertiesArgs) ToStorageSpacesPropertiesPtrOutput() StorageSpacesPropertiesPtrOutput {
	return i.ToStorageSpacesPropertiesPtrOutputWithContext(context.Background())
}

func (i StorageSpacesPropertiesArgs) ToStorageSpacesPropertiesPtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesPropertiesOutput).ToStorageSpacesPropertiesPtrOutputWithContext(ctx)
}









type StorageSpacesPropertiesPtrInput interface {
	pulumi.Input

	ToStorageSpacesPropertiesPtrOutput() StorageSpacesPropertiesPtrOutput
	ToStorageSpacesPropertiesPtrOutputWithContext(context.Context) StorageSpacesPropertiesPtrOutput
}

type storageSpacesPropertiesPtrType StorageSpacesPropertiesArgs

func StorageSpacesPropertiesPtr(v *StorageSpacesPropertiesArgs) StorageSpacesPropertiesPtrInput {
	return (*storageSpacesPropertiesPtrType)(v)
}

func (*storageSpacesPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpacesProperties)(nil)).Elem()
}

func (i *storageSpacesPropertiesPtrType) ToStorageSpacesPropertiesPtrOutput() StorageSpacesPropertiesPtrOutput {
	return i.ToStorageSpacesPropertiesPtrOutputWithContext(context.Background())
}

func (i *storageSpacesPropertiesPtrType) ToStorageSpacesPropertiesPtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesPropertiesPtrOutput)
}

type StorageSpacesPropertiesOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesProperties)(nil)).Elem()
}

func (o StorageSpacesPropertiesOutput) ToStorageSpacesPropertiesOutput() StorageSpacesPropertiesOutput {
	return o
}

func (o StorageSpacesPropertiesOutput) ToStorageSpacesPropertiesOutputWithContext(ctx context.Context) StorageSpacesPropertiesOutput {
	return o
}

func (o StorageSpacesPropertiesOutput) ToStorageSpacesPropertiesPtrOutput() StorageSpacesPropertiesPtrOutput {
	return o.ToStorageSpacesPropertiesPtrOutputWithContext(context.Background())
}

func (o StorageSpacesPropertiesOutput) ToStorageSpacesPropertiesPtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageSpacesProperties) *StorageSpacesProperties {
		return &v
	}).(StorageSpacesPropertiesPtrOutput)
}

func (o StorageSpacesPropertiesOutput) HciStorageProfile() StorageSpacesPropertiesHciStorageProfilePtrOutput {
	return o.ApplyT(func(v StorageSpacesProperties) *StorageSpacesPropertiesHciStorageProfile { return v.HciStorageProfile }).(StorageSpacesPropertiesHciStorageProfilePtrOutput)
}

func (o StorageSpacesPropertiesOutput) Status() StorageSpacesPropertiesStatusPtrOutput {
	return o.ApplyT(func(v StorageSpacesProperties) *StorageSpacesPropertiesStatus { return v.Status }).(StorageSpacesPropertiesStatusPtrOutput)
}

func (o StorageSpacesPropertiesOutput) VmwareStorageProfile() StorageSpacesPropertiesVmwareStorageProfilePtrOutput {
	return o.ApplyT(func(v StorageSpacesProperties) *StorageSpacesPropertiesVmwareStorageProfile {
		return v.VmwareStorageProfile
	}).(StorageSpacesPropertiesVmwareStorageProfilePtrOutput)
}

type StorageSpacesPropertiesPtrOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpacesProperties)(nil)).Elem()
}

func (o StorageSpacesPropertiesPtrOutput) ToStorageSpacesPropertiesPtrOutput() StorageSpacesPropertiesPtrOutput {
	return o
}

func (o StorageSpacesPropertiesPtrOutput) ToStorageSpacesPropertiesPtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesPtrOutput {
	return o
}

func (o StorageSpacesPropertiesPtrOutput) Elem() StorageSpacesPropertiesOutput {
	return o.ApplyT(func(v *StorageSpacesProperties) StorageSpacesProperties {
		if v != nil {
			return *v
		}
		var ret StorageSpacesProperties
		return ret
	}).(StorageSpacesPropertiesOutput)
}

func (o StorageSpacesPropertiesPtrOutput) HciStorageProfile() StorageSpacesPropertiesHciStorageProfilePtrOutput {
	return o.ApplyT(func(v *StorageSpacesProperties) *StorageSpacesPropertiesHciStorageProfile {
		if v == nil {
			return nil
		}
		return v.HciStorageProfile
	}).(StorageSpacesPropertiesHciStorageProfilePtrOutput)
}

func (o StorageSpacesPropertiesPtrOutput) Status() StorageSpacesPropertiesStatusPtrOutput {
	return o.ApplyT(func(v *StorageSpacesProperties) *StorageSpacesPropertiesStatus {
		if v == nil {
			return nil
		}
		return v.Status
	}).(StorageSpacesPropertiesStatusPtrOutput)
}

func (o StorageSpacesPropertiesPtrOutput) VmwareStorageProfile() StorageSpacesPropertiesVmwareStorageProfilePtrOutput {
	return o.ApplyT(func(v *StorageSpacesProperties) *StorageSpacesPropertiesVmwareStorageProfile {
		if v == nil {
			return nil
		}
		return v.VmwareStorageProfile
	}).(StorageSpacesPropertiesVmwareStorageProfilePtrOutput)
}

type StorageSpacesPropertiesError struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}





type StorageSpacesPropertiesErrorInput interface {
	pulumi.Input

	ToStorageSpacesPropertiesErrorOutput() StorageSpacesPropertiesErrorOutput
	ToStorageSpacesPropertiesErrorOutputWithContext(context.Context) StorageSpacesPropertiesErrorOutput
}

type StorageSpacesPropertiesErrorArgs struct {
	Code    pulumi.StringPtrInput `pulumi:"code"`
	Message pulumi.StringPtrInput `pulumi:"message"`
}

func (StorageSpacesPropertiesErrorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesPropertiesError)(nil)).Elem()
}

func (i StorageSpacesPropertiesErrorArgs) ToStorageSpacesPropertiesErrorOutput() StorageSpacesPropertiesErrorOutput {
	return i.ToStorageSpacesPropertiesErrorOutputWithContext(context.Background())
}

func (i StorageSpacesPropertiesErrorArgs) ToStorageSpacesPropertiesErrorOutputWithContext(ctx context.Context) StorageSpacesPropertiesErrorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesPropertiesErrorOutput)
}

func (i StorageSpacesPropertiesErrorArgs) ToStorageSpacesPropertiesErrorPtrOutput() StorageSpacesPropertiesErrorPtrOutput {
	return i.ToStorageSpacesPropertiesErrorPtrOutputWithContext(context.Background())
}

func (i StorageSpacesPropertiesErrorArgs) ToStorageSpacesPropertiesErrorPtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesErrorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesPropertiesErrorOutput).ToStorageSpacesPropertiesErrorPtrOutputWithContext(ctx)
}









type StorageSpacesPropertiesErrorPtrInput interface {
	pulumi.Input

	ToStorageSpacesPropertiesErrorPtrOutput() StorageSpacesPropertiesErrorPtrOutput
	ToStorageSpacesPropertiesErrorPtrOutputWithContext(context.Context) StorageSpacesPropertiesErrorPtrOutput
}

type storageSpacesPropertiesErrorPtrType StorageSpacesPropertiesErrorArgs

func StorageSpacesPropertiesErrorPtr(v *StorageSpacesPropertiesErrorArgs) StorageSpacesPropertiesErrorPtrInput {
	return (*storageSpacesPropertiesErrorPtrType)(v)
}

func (*storageSpacesPropertiesErrorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpacesPropertiesError)(nil)).Elem()
}

func (i *storageSpacesPropertiesErrorPtrType) ToStorageSpacesPropertiesErrorPtrOutput() StorageSpacesPropertiesErrorPtrOutput {
	return i.ToStorageSpacesPropertiesErrorPtrOutputWithContext(context.Background())
}

func (i *storageSpacesPropertiesErrorPtrType) ToStorageSpacesPropertiesErrorPtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesErrorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesPropertiesErrorPtrOutput)
}

type StorageSpacesPropertiesErrorOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesErrorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesPropertiesError)(nil)).Elem()
}

func (o StorageSpacesPropertiesErrorOutput) ToStorageSpacesPropertiesErrorOutput() StorageSpacesPropertiesErrorOutput {
	return o
}

func (o StorageSpacesPropertiesErrorOutput) ToStorageSpacesPropertiesErrorOutputWithContext(ctx context.Context) StorageSpacesPropertiesErrorOutput {
	return o
}

func (o StorageSpacesPropertiesErrorOutput) ToStorageSpacesPropertiesErrorPtrOutput() StorageSpacesPropertiesErrorPtrOutput {
	return o.ToStorageSpacesPropertiesErrorPtrOutputWithContext(context.Background())
}

func (o StorageSpacesPropertiesErrorOutput) ToStorageSpacesPropertiesErrorPtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesErrorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageSpacesPropertiesError) *StorageSpacesPropertiesError {
		return &v
	}).(StorageSpacesPropertiesErrorPtrOutput)
}

func (o StorageSpacesPropertiesErrorOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesError) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesErrorOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesError) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type StorageSpacesPropertiesErrorPtrOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesErrorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpacesPropertiesError)(nil)).Elem()
}

func (o StorageSpacesPropertiesErrorPtrOutput) ToStorageSpacesPropertiesErrorPtrOutput() StorageSpacesPropertiesErrorPtrOutput {
	return o
}

func (o StorageSpacesPropertiesErrorPtrOutput) ToStorageSpacesPropertiesErrorPtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesErrorPtrOutput {
	return o
}

func (o StorageSpacesPropertiesErrorPtrOutput) Elem() StorageSpacesPropertiesErrorOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesError) StorageSpacesPropertiesError {
		if v != nil {
			return *v
		}
		var ret StorageSpacesPropertiesError
		return ret
	}).(StorageSpacesPropertiesErrorOutput)
}

func (o StorageSpacesPropertiesErrorPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesError) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesErrorPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesError) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

type StorageSpacesPropertiesHciStorageProfile struct {
	MocGroup            *string `pulumi:"mocGroup"`
	MocLocation         *string `pulumi:"mocLocation"`
	MocStorageContainer *string `pulumi:"mocStorageContainer"`
}





type StorageSpacesPropertiesHciStorageProfileInput interface {
	pulumi.Input

	ToStorageSpacesPropertiesHciStorageProfileOutput() StorageSpacesPropertiesHciStorageProfileOutput
	ToStorageSpacesPropertiesHciStorageProfileOutputWithContext(context.Context) StorageSpacesPropertiesHciStorageProfileOutput
}

type StorageSpacesPropertiesHciStorageProfileArgs struct {
	MocGroup            pulumi.StringPtrInput `pulumi:"mocGroup"`
	MocLocation         pulumi.StringPtrInput `pulumi:"mocLocation"`
	MocStorageContainer pulumi.StringPtrInput `pulumi:"mocStorageContainer"`
}

func (StorageSpacesPropertiesHciStorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesPropertiesHciStorageProfile)(nil)).Elem()
}

func (i StorageSpacesPropertiesHciStorageProfileArgs) ToStorageSpacesPropertiesHciStorageProfileOutput() StorageSpacesPropertiesHciStorageProfileOutput {
	return i.ToStorageSpacesPropertiesHciStorageProfileOutputWithContext(context.Background())
}

func (i StorageSpacesPropertiesHciStorageProfileArgs) ToStorageSpacesPropertiesHciStorageProfileOutputWithContext(ctx context.Context) StorageSpacesPropertiesHciStorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesPropertiesHciStorageProfileOutput)
}

func (i StorageSpacesPropertiesHciStorageProfileArgs) ToStorageSpacesPropertiesHciStorageProfilePtrOutput() StorageSpacesPropertiesHciStorageProfilePtrOutput {
	return i.ToStorageSpacesPropertiesHciStorageProfilePtrOutputWithContext(context.Background())
}

func (i StorageSpacesPropertiesHciStorageProfileArgs) ToStorageSpacesPropertiesHciStorageProfilePtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesHciStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesPropertiesHciStorageProfileOutput).ToStorageSpacesPropertiesHciStorageProfilePtrOutputWithContext(ctx)
}









type StorageSpacesPropertiesHciStorageProfilePtrInput interface {
	pulumi.Input

	ToStorageSpacesPropertiesHciStorageProfilePtrOutput() StorageSpacesPropertiesHciStorageProfilePtrOutput
	ToStorageSpacesPropertiesHciStorageProfilePtrOutputWithContext(context.Context) StorageSpacesPropertiesHciStorageProfilePtrOutput
}

type storageSpacesPropertiesHciStorageProfilePtrType StorageSpacesPropertiesHciStorageProfileArgs

func StorageSpacesPropertiesHciStorageProfilePtr(v *StorageSpacesPropertiesHciStorageProfileArgs) StorageSpacesPropertiesHciStorageProfilePtrInput {
	return (*storageSpacesPropertiesHciStorageProfilePtrType)(v)
}

func (*storageSpacesPropertiesHciStorageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpacesPropertiesHciStorageProfile)(nil)).Elem()
}

func (i *storageSpacesPropertiesHciStorageProfilePtrType) ToStorageSpacesPropertiesHciStorageProfilePtrOutput() StorageSpacesPropertiesHciStorageProfilePtrOutput {
	return i.ToStorageSpacesPropertiesHciStorageProfilePtrOutputWithContext(context.Background())
}

func (i *storageSpacesPropertiesHciStorageProfilePtrType) ToStorageSpacesPropertiesHciStorageProfilePtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesHciStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesPropertiesHciStorageProfilePtrOutput)
}

type StorageSpacesPropertiesHciStorageProfileOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesHciStorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesPropertiesHciStorageProfile)(nil)).Elem()
}

func (o StorageSpacesPropertiesHciStorageProfileOutput) ToStorageSpacesPropertiesHciStorageProfileOutput() StorageSpacesPropertiesHciStorageProfileOutput {
	return o
}

func (o StorageSpacesPropertiesHciStorageProfileOutput) ToStorageSpacesPropertiesHciStorageProfileOutputWithContext(ctx context.Context) StorageSpacesPropertiesHciStorageProfileOutput {
	return o
}

func (o StorageSpacesPropertiesHciStorageProfileOutput) ToStorageSpacesPropertiesHciStorageProfilePtrOutput() StorageSpacesPropertiesHciStorageProfilePtrOutput {
	return o.ToStorageSpacesPropertiesHciStorageProfilePtrOutputWithContext(context.Background())
}

func (o StorageSpacesPropertiesHciStorageProfileOutput) ToStorageSpacesPropertiesHciStorageProfilePtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesHciStorageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageSpacesPropertiesHciStorageProfile) *StorageSpacesPropertiesHciStorageProfile {
		return &v
	}).(StorageSpacesPropertiesHciStorageProfilePtrOutput)
}

func (o StorageSpacesPropertiesHciStorageProfileOutput) MocGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesHciStorageProfile) *string { return v.MocGroup }).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesHciStorageProfileOutput) MocLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesHciStorageProfile) *string { return v.MocLocation }).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesHciStorageProfileOutput) MocStorageContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesHciStorageProfile) *string { return v.MocStorageContainer }).(pulumi.StringPtrOutput)
}

type StorageSpacesPropertiesHciStorageProfilePtrOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesHciStorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpacesPropertiesHciStorageProfile)(nil)).Elem()
}

func (o StorageSpacesPropertiesHciStorageProfilePtrOutput) ToStorageSpacesPropertiesHciStorageProfilePtrOutput() StorageSpacesPropertiesHciStorageProfilePtrOutput {
	return o
}

func (o StorageSpacesPropertiesHciStorageProfilePtrOutput) ToStorageSpacesPropertiesHciStorageProfilePtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesHciStorageProfilePtrOutput {
	return o
}

func (o StorageSpacesPropertiesHciStorageProfilePtrOutput) Elem() StorageSpacesPropertiesHciStorageProfileOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesHciStorageProfile) StorageSpacesPropertiesHciStorageProfile {
		if v != nil {
			return *v
		}
		var ret StorageSpacesPropertiesHciStorageProfile
		return ret
	}).(StorageSpacesPropertiesHciStorageProfileOutput)
}

func (o StorageSpacesPropertiesHciStorageProfilePtrOutput) MocGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesHciStorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.MocGroup
	}).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesHciStorageProfilePtrOutput) MocLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesHciStorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.MocLocation
	}).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesHciStorageProfilePtrOutput) MocStorageContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesHciStorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.MocStorageContainer
	}).(pulumi.StringPtrOutput)
}

type StorageSpacesPropertiesProvisioningStatus struct {
	Error       *StorageSpacesPropertiesError `pulumi:"error"`
	OperationId *string                       `pulumi:"operationId"`
	Phase       *string                       `pulumi:"phase"`
	Status      *string                       `pulumi:"status"`
}





type StorageSpacesPropertiesProvisioningStatusInput interface {
	pulumi.Input

	ToStorageSpacesPropertiesProvisioningStatusOutput() StorageSpacesPropertiesProvisioningStatusOutput
	ToStorageSpacesPropertiesProvisioningStatusOutputWithContext(context.Context) StorageSpacesPropertiesProvisioningStatusOutput
}

type StorageSpacesPropertiesProvisioningStatusArgs struct {
	Error       StorageSpacesPropertiesErrorPtrInput `pulumi:"error"`
	OperationId pulumi.StringPtrInput                `pulumi:"operationId"`
	Phase       pulumi.StringPtrInput                `pulumi:"phase"`
	Status      pulumi.StringPtrInput                `pulumi:"status"`
}

func (StorageSpacesPropertiesProvisioningStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesPropertiesProvisioningStatus)(nil)).Elem()
}

func (i StorageSpacesPropertiesProvisioningStatusArgs) ToStorageSpacesPropertiesProvisioningStatusOutput() StorageSpacesPropertiesProvisioningStatusOutput {
	return i.ToStorageSpacesPropertiesProvisioningStatusOutputWithContext(context.Background())
}

func (i StorageSpacesPropertiesProvisioningStatusArgs) ToStorageSpacesPropertiesProvisioningStatusOutputWithContext(ctx context.Context) StorageSpacesPropertiesProvisioningStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesPropertiesProvisioningStatusOutput)
}

func (i StorageSpacesPropertiesProvisioningStatusArgs) ToStorageSpacesPropertiesProvisioningStatusPtrOutput() StorageSpacesPropertiesProvisioningStatusPtrOutput {
	return i.ToStorageSpacesPropertiesProvisioningStatusPtrOutputWithContext(context.Background())
}

func (i StorageSpacesPropertiesProvisioningStatusArgs) ToStorageSpacesPropertiesProvisioningStatusPtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesProvisioningStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesPropertiesProvisioningStatusOutput).ToStorageSpacesPropertiesProvisioningStatusPtrOutputWithContext(ctx)
}









type StorageSpacesPropertiesProvisioningStatusPtrInput interface {
	pulumi.Input

	ToStorageSpacesPropertiesProvisioningStatusPtrOutput() StorageSpacesPropertiesProvisioningStatusPtrOutput
	ToStorageSpacesPropertiesProvisioningStatusPtrOutputWithContext(context.Context) StorageSpacesPropertiesProvisioningStatusPtrOutput
}

type storageSpacesPropertiesProvisioningStatusPtrType StorageSpacesPropertiesProvisioningStatusArgs

func StorageSpacesPropertiesProvisioningStatusPtr(v *StorageSpacesPropertiesProvisioningStatusArgs) StorageSpacesPropertiesProvisioningStatusPtrInput {
	return (*storageSpacesPropertiesProvisioningStatusPtrType)(v)
}

func (*storageSpacesPropertiesProvisioningStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpacesPropertiesProvisioningStatus)(nil)).Elem()
}

func (i *storageSpacesPropertiesProvisioningStatusPtrType) ToStorageSpacesPropertiesProvisioningStatusPtrOutput() StorageSpacesPropertiesProvisioningStatusPtrOutput {
	return i.ToStorageSpacesPropertiesProvisioningStatusPtrOutputWithContext(context.Background())
}

func (i *storageSpacesPropertiesProvisioningStatusPtrType) ToStorageSpacesPropertiesProvisioningStatusPtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesProvisioningStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesPropertiesProvisioningStatusPtrOutput)
}

type StorageSpacesPropertiesProvisioningStatusOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesProvisioningStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesPropertiesProvisioningStatus)(nil)).Elem()
}

func (o StorageSpacesPropertiesProvisioningStatusOutput) ToStorageSpacesPropertiesProvisioningStatusOutput() StorageSpacesPropertiesProvisioningStatusOutput {
	return o
}

func (o StorageSpacesPropertiesProvisioningStatusOutput) ToStorageSpacesPropertiesProvisioningStatusOutputWithContext(ctx context.Context) StorageSpacesPropertiesProvisioningStatusOutput {
	return o
}

func (o StorageSpacesPropertiesProvisioningStatusOutput) ToStorageSpacesPropertiesProvisioningStatusPtrOutput() StorageSpacesPropertiesProvisioningStatusPtrOutput {
	return o.ToStorageSpacesPropertiesProvisioningStatusPtrOutputWithContext(context.Background())
}

func (o StorageSpacesPropertiesProvisioningStatusOutput) ToStorageSpacesPropertiesProvisioningStatusPtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesProvisioningStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageSpacesPropertiesProvisioningStatus) *StorageSpacesPropertiesProvisioningStatus {
		return &v
	}).(StorageSpacesPropertiesProvisioningStatusPtrOutput)
}

func (o StorageSpacesPropertiesProvisioningStatusOutput) Error() StorageSpacesPropertiesErrorPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesProvisioningStatus) *StorageSpacesPropertiesError { return v.Error }).(StorageSpacesPropertiesErrorPtrOutput)
}

func (o StorageSpacesPropertiesProvisioningStatusOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesProvisioningStatus) *string { return v.OperationId }).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesProvisioningStatusOutput) Phase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesProvisioningStatus) *string { return v.Phase }).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesProvisioningStatusOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesProvisioningStatus) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type StorageSpacesPropertiesProvisioningStatusPtrOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesProvisioningStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpacesPropertiesProvisioningStatus)(nil)).Elem()
}

func (o StorageSpacesPropertiesProvisioningStatusPtrOutput) ToStorageSpacesPropertiesProvisioningStatusPtrOutput() StorageSpacesPropertiesProvisioningStatusPtrOutput {
	return o
}

func (o StorageSpacesPropertiesProvisioningStatusPtrOutput) ToStorageSpacesPropertiesProvisioningStatusPtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesProvisioningStatusPtrOutput {
	return o
}

func (o StorageSpacesPropertiesProvisioningStatusPtrOutput) Elem() StorageSpacesPropertiesProvisioningStatusOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesProvisioningStatus) StorageSpacesPropertiesProvisioningStatus {
		if v != nil {
			return *v
		}
		var ret StorageSpacesPropertiesProvisioningStatus
		return ret
	}).(StorageSpacesPropertiesProvisioningStatusOutput)
}

func (o StorageSpacesPropertiesProvisioningStatusPtrOutput) Error() StorageSpacesPropertiesErrorPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesProvisioningStatus) *StorageSpacesPropertiesError {
		if v == nil {
			return nil
		}
		return v.Error
	}).(StorageSpacesPropertiesErrorPtrOutput)
}

func (o StorageSpacesPropertiesProvisioningStatusPtrOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.OperationId
	}).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesProvisioningStatusPtrOutput) Phase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.Phase
	}).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesProvisioningStatusPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type StorageSpacesPropertiesResponse struct {
	HciStorageProfile    *StorageSpacesPropertiesResponseHciStorageProfile    `pulumi:"hciStorageProfile"`
	ProvisioningState    string                                               `pulumi:"provisioningState"`
	Status               *StorageSpacesPropertiesResponseStatus               `pulumi:"status"`
	VmwareStorageProfile *StorageSpacesPropertiesResponseVmwareStorageProfile `pulumi:"vmwareStorageProfile"`
}

type StorageSpacesPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesPropertiesResponse)(nil)).Elem()
}

func (o StorageSpacesPropertiesResponseOutput) ToStorageSpacesPropertiesResponseOutput() StorageSpacesPropertiesResponseOutput {
	return o
}

func (o StorageSpacesPropertiesResponseOutput) ToStorageSpacesPropertiesResponseOutputWithContext(ctx context.Context) StorageSpacesPropertiesResponseOutput {
	return o
}

func (o StorageSpacesPropertiesResponseOutput) HciStorageProfile() StorageSpacesPropertiesResponseHciStorageProfilePtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesResponse) *StorageSpacesPropertiesResponseHciStorageProfile {
		return v.HciStorageProfile
	}).(StorageSpacesPropertiesResponseHciStorageProfilePtrOutput)
}

func (o StorageSpacesPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o StorageSpacesPropertiesResponseOutput) Status() StorageSpacesPropertiesResponseStatusPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesResponse) *StorageSpacesPropertiesResponseStatus { return v.Status }).(StorageSpacesPropertiesResponseStatusPtrOutput)
}

func (o StorageSpacesPropertiesResponseOutput) VmwareStorageProfile() StorageSpacesPropertiesResponseVmwareStorageProfilePtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesResponse) *StorageSpacesPropertiesResponseVmwareStorageProfile {
		return v.VmwareStorageProfile
	}).(StorageSpacesPropertiesResponseVmwareStorageProfilePtrOutput)
}

type StorageSpacesPropertiesResponseError struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}

type StorageSpacesPropertiesResponseErrorOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesResponseErrorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesPropertiesResponseError)(nil)).Elem()
}

func (o StorageSpacesPropertiesResponseErrorOutput) ToStorageSpacesPropertiesResponseErrorOutput() StorageSpacesPropertiesResponseErrorOutput {
	return o
}

func (o StorageSpacesPropertiesResponseErrorOutput) ToStorageSpacesPropertiesResponseErrorOutputWithContext(ctx context.Context) StorageSpacesPropertiesResponseErrorOutput {
	return o
}

func (o StorageSpacesPropertiesResponseErrorOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesResponseError) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesResponseErrorOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesResponseError) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type StorageSpacesPropertiesResponseErrorPtrOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesResponseErrorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpacesPropertiesResponseError)(nil)).Elem()
}

func (o StorageSpacesPropertiesResponseErrorPtrOutput) ToStorageSpacesPropertiesResponseErrorPtrOutput() StorageSpacesPropertiesResponseErrorPtrOutput {
	return o
}

func (o StorageSpacesPropertiesResponseErrorPtrOutput) ToStorageSpacesPropertiesResponseErrorPtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesResponseErrorPtrOutput {
	return o
}

func (o StorageSpacesPropertiesResponseErrorPtrOutput) Elem() StorageSpacesPropertiesResponseErrorOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesResponseError) StorageSpacesPropertiesResponseError {
		if v != nil {
			return *v
		}
		var ret StorageSpacesPropertiesResponseError
		return ret
	}).(StorageSpacesPropertiesResponseErrorOutput)
}

func (o StorageSpacesPropertiesResponseErrorPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesResponseError) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesResponseErrorPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesResponseError) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

type StorageSpacesPropertiesResponseHciStorageProfile struct {
	MocGroup            *string `pulumi:"mocGroup"`
	MocLocation         *string `pulumi:"mocLocation"`
	MocStorageContainer *string `pulumi:"mocStorageContainer"`
}

type StorageSpacesPropertiesResponseHciStorageProfileOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesResponseHciStorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesPropertiesResponseHciStorageProfile)(nil)).Elem()
}

func (o StorageSpacesPropertiesResponseHciStorageProfileOutput) ToStorageSpacesPropertiesResponseHciStorageProfileOutput() StorageSpacesPropertiesResponseHciStorageProfileOutput {
	return o
}

func (o StorageSpacesPropertiesResponseHciStorageProfileOutput) ToStorageSpacesPropertiesResponseHciStorageProfileOutputWithContext(ctx context.Context) StorageSpacesPropertiesResponseHciStorageProfileOutput {
	return o
}

func (o StorageSpacesPropertiesResponseHciStorageProfileOutput) MocGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesResponseHciStorageProfile) *string { return v.MocGroup }).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesResponseHciStorageProfileOutput) MocLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesResponseHciStorageProfile) *string { return v.MocLocation }).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesResponseHciStorageProfileOutput) MocStorageContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesResponseHciStorageProfile) *string { return v.MocStorageContainer }).(pulumi.StringPtrOutput)
}

type StorageSpacesPropertiesResponseHciStorageProfilePtrOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesResponseHciStorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpacesPropertiesResponseHciStorageProfile)(nil)).Elem()
}

func (o StorageSpacesPropertiesResponseHciStorageProfilePtrOutput) ToStorageSpacesPropertiesResponseHciStorageProfilePtrOutput() StorageSpacesPropertiesResponseHciStorageProfilePtrOutput {
	return o
}

func (o StorageSpacesPropertiesResponseHciStorageProfilePtrOutput) ToStorageSpacesPropertiesResponseHciStorageProfilePtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesResponseHciStorageProfilePtrOutput {
	return o
}

func (o StorageSpacesPropertiesResponseHciStorageProfilePtrOutput) Elem() StorageSpacesPropertiesResponseHciStorageProfileOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesResponseHciStorageProfile) StorageSpacesPropertiesResponseHciStorageProfile {
		if v != nil {
			return *v
		}
		var ret StorageSpacesPropertiesResponseHciStorageProfile
		return ret
	}).(StorageSpacesPropertiesResponseHciStorageProfileOutput)
}

func (o StorageSpacesPropertiesResponseHciStorageProfilePtrOutput) MocGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesResponseHciStorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.MocGroup
	}).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesResponseHciStorageProfilePtrOutput) MocLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesResponseHciStorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.MocLocation
	}).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesResponseHciStorageProfilePtrOutput) MocStorageContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesResponseHciStorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.MocStorageContainer
	}).(pulumi.StringPtrOutput)
}

type StorageSpacesPropertiesResponseProvisioningStatus struct {
	Error       *StorageSpacesPropertiesResponseError `pulumi:"error"`
	OperationId *string                               `pulumi:"operationId"`
	Phase       *string                               `pulumi:"phase"`
	Status      *string                               `pulumi:"status"`
}

type StorageSpacesPropertiesResponseProvisioningStatusOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesResponseProvisioningStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesPropertiesResponseProvisioningStatus)(nil)).Elem()
}

func (o StorageSpacesPropertiesResponseProvisioningStatusOutput) ToStorageSpacesPropertiesResponseProvisioningStatusOutput() StorageSpacesPropertiesResponseProvisioningStatusOutput {
	return o
}

func (o StorageSpacesPropertiesResponseProvisioningStatusOutput) ToStorageSpacesPropertiesResponseProvisioningStatusOutputWithContext(ctx context.Context) StorageSpacesPropertiesResponseProvisioningStatusOutput {
	return o
}

func (o StorageSpacesPropertiesResponseProvisioningStatusOutput) Error() StorageSpacesPropertiesResponseErrorPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesResponseProvisioningStatus) *StorageSpacesPropertiesResponseError {
		return v.Error
	}).(StorageSpacesPropertiesResponseErrorPtrOutput)
}

func (o StorageSpacesPropertiesResponseProvisioningStatusOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesResponseProvisioningStatus) *string { return v.OperationId }).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesResponseProvisioningStatusOutput) Phase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesResponseProvisioningStatus) *string { return v.Phase }).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesResponseProvisioningStatusOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesResponseProvisioningStatus) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type StorageSpacesPropertiesResponseProvisioningStatusPtrOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesResponseProvisioningStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpacesPropertiesResponseProvisioningStatus)(nil)).Elem()
}

func (o StorageSpacesPropertiesResponseProvisioningStatusPtrOutput) ToStorageSpacesPropertiesResponseProvisioningStatusPtrOutput() StorageSpacesPropertiesResponseProvisioningStatusPtrOutput {
	return o
}

func (o StorageSpacesPropertiesResponseProvisioningStatusPtrOutput) ToStorageSpacesPropertiesResponseProvisioningStatusPtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesResponseProvisioningStatusPtrOutput {
	return o
}

func (o StorageSpacesPropertiesResponseProvisioningStatusPtrOutput) Elem() StorageSpacesPropertiesResponseProvisioningStatusOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesResponseProvisioningStatus) StorageSpacesPropertiesResponseProvisioningStatus {
		if v != nil {
			return *v
		}
		var ret StorageSpacesPropertiesResponseProvisioningStatus
		return ret
	}).(StorageSpacesPropertiesResponseProvisioningStatusOutput)
}

func (o StorageSpacesPropertiesResponseProvisioningStatusPtrOutput) Error() StorageSpacesPropertiesResponseErrorPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesResponseProvisioningStatus) *StorageSpacesPropertiesResponseError {
		if v == nil {
			return nil
		}
		return v.Error
	}).(StorageSpacesPropertiesResponseErrorPtrOutput)
}

func (o StorageSpacesPropertiesResponseProvisioningStatusPtrOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.OperationId
	}).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesResponseProvisioningStatusPtrOutput) Phase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.Phase
	}).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesResponseProvisioningStatusPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type StorageSpacesPropertiesResponseStatus struct {
	ProvisioningStatus *StorageSpacesPropertiesResponseProvisioningStatus `pulumi:"provisioningStatus"`
}

type StorageSpacesPropertiesResponseStatusOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesResponseStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesPropertiesResponseStatus)(nil)).Elem()
}

func (o StorageSpacesPropertiesResponseStatusOutput) ToStorageSpacesPropertiesResponseStatusOutput() StorageSpacesPropertiesResponseStatusOutput {
	return o
}

func (o StorageSpacesPropertiesResponseStatusOutput) ToStorageSpacesPropertiesResponseStatusOutputWithContext(ctx context.Context) StorageSpacesPropertiesResponseStatusOutput {
	return o
}

func (o StorageSpacesPropertiesResponseStatusOutput) ProvisioningStatus() StorageSpacesPropertiesResponseProvisioningStatusPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesResponseStatus) *StorageSpacesPropertiesResponseProvisioningStatus {
		return v.ProvisioningStatus
	}).(StorageSpacesPropertiesResponseProvisioningStatusPtrOutput)
}

type StorageSpacesPropertiesResponseStatusPtrOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesResponseStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpacesPropertiesResponseStatus)(nil)).Elem()
}

func (o StorageSpacesPropertiesResponseStatusPtrOutput) ToStorageSpacesPropertiesResponseStatusPtrOutput() StorageSpacesPropertiesResponseStatusPtrOutput {
	return o
}

func (o StorageSpacesPropertiesResponseStatusPtrOutput) ToStorageSpacesPropertiesResponseStatusPtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesResponseStatusPtrOutput {
	return o
}

func (o StorageSpacesPropertiesResponseStatusPtrOutput) Elem() StorageSpacesPropertiesResponseStatusOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesResponseStatus) StorageSpacesPropertiesResponseStatus {
		if v != nil {
			return *v
		}
		var ret StorageSpacesPropertiesResponseStatus
		return ret
	}).(StorageSpacesPropertiesResponseStatusOutput)
}

func (o StorageSpacesPropertiesResponseStatusPtrOutput) ProvisioningStatus() StorageSpacesPropertiesResponseProvisioningStatusPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesResponseStatus) *StorageSpacesPropertiesResponseProvisioningStatus {
		if v == nil {
			return nil
		}
		return v.ProvisioningStatus
	}).(StorageSpacesPropertiesResponseProvisioningStatusPtrOutput)
}

type StorageSpacesPropertiesResponseVmwareStorageProfile struct {
	Datacenter   *string `pulumi:"datacenter"`
	Datastore    *string `pulumi:"datastore"`
	Folder       *string `pulumi:"folder"`
	ResourcePool *string `pulumi:"resourcePool"`
}

type StorageSpacesPropertiesResponseVmwareStorageProfileOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesResponseVmwareStorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesPropertiesResponseVmwareStorageProfile)(nil)).Elem()
}

func (o StorageSpacesPropertiesResponseVmwareStorageProfileOutput) ToStorageSpacesPropertiesResponseVmwareStorageProfileOutput() StorageSpacesPropertiesResponseVmwareStorageProfileOutput {
	return o
}

func (o StorageSpacesPropertiesResponseVmwareStorageProfileOutput) ToStorageSpacesPropertiesResponseVmwareStorageProfileOutputWithContext(ctx context.Context) StorageSpacesPropertiesResponseVmwareStorageProfileOutput {
	return o
}

func (o StorageSpacesPropertiesResponseVmwareStorageProfileOutput) Datacenter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesResponseVmwareStorageProfile) *string { return v.Datacenter }).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesResponseVmwareStorageProfileOutput) Datastore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesResponseVmwareStorageProfile) *string { return v.Datastore }).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesResponseVmwareStorageProfileOutput) Folder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesResponseVmwareStorageProfile) *string { return v.Folder }).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesResponseVmwareStorageProfileOutput) ResourcePool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesResponseVmwareStorageProfile) *string { return v.ResourcePool }).(pulumi.StringPtrOutput)
}

type StorageSpacesPropertiesResponseVmwareStorageProfilePtrOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesResponseVmwareStorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpacesPropertiesResponseVmwareStorageProfile)(nil)).Elem()
}

func (o StorageSpacesPropertiesResponseVmwareStorageProfilePtrOutput) ToStorageSpacesPropertiesResponseVmwareStorageProfilePtrOutput() StorageSpacesPropertiesResponseVmwareStorageProfilePtrOutput {
	return o
}

func (o StorageSpacesPropertiesResponseVmwareStorageProfilePtrOutput) ToStorageSpacesPropertiesResponseVmwareStorageProfilePtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesResponseVmwareStorageProfilePtrOutput {
	return o
}

func (o StorageSpacesPropertiesResponseVmwareStorageProfilePtrOutput) Elem() StorageSpacesPropertiesResponseVmwareStorageProfileOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesResponseVmwareStorageProfile) StorageSpacesPropertiesResponseVmwareStorageProfile {
		if v != nil {
			return *v
		}
		var ret StorageSpacesPropertiesResponseVmwareStorageProfile
		return ret
	}).(StorageSpacesPropertiesResponseVmwareStorageProfileOutput)
}

func (o StorageSpacesPropertiesResponseVmwareStorageProfilePtrOutput) Datacenter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesResponseVmwareStorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.Datacenter
	}).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesResponseVmwareStorageProfilePtrOutput) Datastore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesResponseVmwareStorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.Datastore
	}).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesResponseVmwareStorageProfilePtrOutput) Folder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesResponseVmwareStorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.Folder
	}).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesResponseVmwareStorageProfilePtrOutput) ResourcePool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesResponseVmwareStorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.ResourcePool
	}).(pulumi.StringPtrOutput)
}

type StorageSpacesPropertiesStatus struct {
	ProvisioningStatus *StorageSpacesPropertiesProvisioningStatus `pulumi:"provisioningStatus"`
}





type StorageSpacesPropertiesStatusInput interface {
	pulumi.Input

	ToStorageSpacesPropertiesStatusOutput() StorageSpacesPropertiesStatusOutput
	ToStorageSpacesPropertiesStatusOutputWithContext(context.Context) StorageSpacesPropertiesStatusOutput
}

type StorageSpacesPropertiesStatusArgs struct {
	ProvisioningStatus StorageSpacesPropertiesProvisioningStatusPtrInput `pulumi:"provisioningStatus"`
}

func (StorageSpacesPropertiesStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesPropertiesStatus)(nil)).Elem()
}

func (i StorageSpacesPropertiesStatusArgs) ToStorageSpacesPropertiesStatusOutput() StorageSpacesPropertiesStatusOutput {
	return i.ToStorageSpacesPropertiesStatusOutputWithContext(context.Background())
}

func (i StorageSpacesPropertiesStatusArgs) ToStorageSpacesPropertiesStatusOutputWithContext(ctx context.Context) StorageSpacesPropertiesStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesPropertiesStatusOutput)
}

func (i StorageSpacesPropertiesStatusArgs) ToStorageSpacesPropertiesStatusPtrOutput() StorageSpacesPropertiesStatusPtrOutput {
	return i.ToStorageSpacesPropertiesStatusPtrOutputWithContext(context.Background())
}

func (i StorageSpacesPropertiesStatusArgs) ToStorageSpacesPropertiesStatusPtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesPropertiesStatusOutput).ToStorageSpacesPropertiesStatusPtrOutputWithContext(ctx)
}









type StorageSpacesPropertiesStatusPtrInput interface {
	pulumi.Input

	ToStorageSpacesPropertiesStatusPtrOutput() StorageSpacesPropertiesStatusPtrOutput
	ToStorageSpacesPropertiesStatusPtrOutputWithContext(context.Context) StorageSpacesPropertiesStatusPtrOutput
}

type storageSpacesPropertiesStatusPtrType StorageSpacesPropertiesStatusArgs

func StorageSpacesPropertiesStatusPtr(v *StorageSpacesPropertiesStatusArgs) StorageSpacesPropertiesStatusPtrInput {
	return (*storageSpacesPropertiesStatusPtrType)(v)
}

func (*storageSpacesPropertiesStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpacesPropertiesStatus)(nil)).Elem()
}

func (i *storageSpacesPropertiesStatusPtrType) ToStorageSpacesPropertiesStatusPtrOutput() StorageSpacesPropertiesStatusPtrOutput {
	return i.ToStorageSpacesPropertiesStatusPtrOutputWithContext(context.Background())
}

func (i *storageSpacesPropertiesStatusPtrType) ToStorageSpacesPropertiesStatusPtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesPropertiesStatusPtrOutput)
}

type StorageSpacesPropertiesStatusOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesPropertiesStatus)(nil)).Elem()
}

func (o StorageSpacesPropertiesStatusOutput) ToStorageSpacesPropertiesStatusOutput() StorageSpacesPropertiesStatusOutput {
	return o
}

func (o StorageSpacesPropertiesStatusOutput) ToStorageSpacesPropertiesStatusOutputWithContext(ctx context.Context) StorageSpacesPropertiesStatusOutput {
	return o
}

func (o StorageSpacesPropertiesStatusOutput) ToStorageSpacesPropertiesStatusPtrOutput() StorageSpacesPropertiesStatusPtrOutput {
	return o.ToStorageSpacesPropertiesStatusPtrOutputWithContext(context.Background())
}

func (o StorageSpacesPropertiesStatusOutput) ToStorageSpacesPropertiesStatusPtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageSpacesPropertiesStatus) *StorageSpacesPropertiesStatus {
		return &v
	}).(StorageSpacesPropertiesStatusPtrOutput)
}

func (o StorageSpacesPropertiesStatusOutput) ProvisioningStatus() StorageSpacesPropertiesProvisioningStatusPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesStatus) *StorageSpacesPropertiesProvisioningStatus {
		return v.ProvisioningStatus
	}).(StorageSpacesPropertiesProvisioningStatusPtrOutput)
}

type StorageSpacesPropertiesStatusPtrOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpacesPropertiesStatus)(nil)).Elem()
}

func (o StorageSpacesPropertiesStatusPtrOutput) ToStorageSpacesPropertiesStatusPtrOutput() StorageSpacesPropertiesStatusPtrOutput {
	return o
}

func (o StorageSpacesPropertiesStatusPtrOutput) ToStorageSpacesPropertiesStatusPtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesStatusPtrOutput {
	return o
}

func (o StorageSpacesPropertiesStatusPtrOutput) Elem() StorageSpacesPropertiesStatusOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesStatus) StorageSpacesPropertiesStatus {
		if v != nil {
			return *v
		}
		var ret StorageSpacesPropertiesStatus
		return ret
	}).(StorageSpacesPropertiesStatusOutput)
}

func (o StorageSpacesPropertiesStatusPtrOutput) ProvisioningStatus() StorageSpacesPropertiesProvisioningStatusPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesStatus) *StorageSpacesPropertiesProvisioningStatus {
		if v == nil {
			return nil
		}
		return v.ProvisioningStatus
	}).(StorageSpacesPropertiesProvisioningStatusPtrOutput)
}

type StorageSpacesPropertiesVmwareStorageProfile struct {
	Datacenter   *string `pulumi:"datacenter"`
	Datastore    *string `pulumi:"datastore"`
	Folder       *string `pulumi:"folder"`
	ResourcePool *string `pulumi:"resourcePool"`
}





type StorageSpacesPropertiesVmwareStorageProfileInput interface {
	pulumi.Input

	ToStorageSpacesPropertiesVmwareStorageProfileOutput() StorageSpacesPropertiesVmwareStorageProfileOutput
	ToStorageSpacesPropertiesVmwareStorageProfileOutputWithContext(context.Context) StorageSpacesPropertiesVmwareStorageProfileOutput
}

type StorageSpacesPropertiesVmwareStorageProfileArgs struct {
	Datacenter   pulumi.StringPtrInput `pulumi:"datacenter"`
	Datastore    pulumi.StringPtrInput `pulumi:"datastore"`
	Folder       pulumi.StringPtrInput `pulumi:"folder"`
	ResourcePool pulumi.StringPtrInput `pulumi:"resourcePool"`
}

func (StorageSpacesPropertiesVmwareStorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesPropertiesVmwareStorageProfile)(nil)).Elem()
}

func (i StorageSpacesPropertiesVmwareStorageProfileArgs) ToStorageSpacesPropertiesVmwareStorageProfileOutput() StorageSpacesPropertiesVmwareStorageProfileOutput {
	return i.ToStorageSpacesPropertiesVmwareStorageProfileOutputWithContext(context.Background())
}

func (i StorageSpacesPropertiesVmwareStorageProfileArgs) ToStorageSpacesPropertiesVmwareStorageProfileOutputWithContext(ctx context.Context) StorageSpacesPropertiesVmwareStorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesPropertiesVmwareStorageProfileOutput)
}

func (i StorageSpacesPropertiesVmwareStorageProfileArgs) ToStorageSpacesPropertiesVmwareStorageProfilePtrOutput() StorageSpacesPropertiesVmwareStorageProfilePtrOutput {
	return i.ToStorageSpacesPropertiesVmwareStorageProfilePtrOutputWithContext(context.Background())
}

func (i StorageSpacesPropertiesVmwareStorageProfileArgs) ToStorageSpacesPropertiesVmwareStorageProfilePtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesVmwareStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesPropertiesVmwareStorageProfileOutput).ToStorageSpacesPropertiesVmwareStorageProfilePtrOutputWithContext(ctx)
}









type StorageSpacesPropertiesVmwareStorageProfilePtrInput interface {
	pulumi.Input

	ToStorageSpacesPropertiesVmwareStorageProfilePtrOutput() StorageSpacesPropertiesVmwareStorageProfilePtrOutput
	ToStorageSpacesPropertiesVmwareStorageProfilePtrOutputWithContext(context.Context) StorageSpacesPropertiesVmwareStorageProfilePtrOutput
}

type storageSpacesPropertiesVmwareStorageProfilePtrType StorageSpacesPropertiesVmwareStorageProfileArgs

func StorageSpacesPropertiesVmwareStorageProfilePtr(v *StorageSpacesPropertiesVmwareStorageProfileArgs) StorageSpacesPropertiesVmwareStorageProfilePtrInput {
	return (*storageSpacesPropertiesVmwareStorageProfilePtrType)(v)
}

func (*storageSpacesPropertiesVmwareStorageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpacesPropertiesVmwareStorageProfile)(nil)).Elem()
}

func (i *storageSpacesPropertiesVmwareStorageProfilePtrType) ToStorageSpacesPropertiesVmwareStorageProfilePtrOutput() StorageSpacesPropertiesVmwareStorageProfilePtrOutput {
	return i.ToStorageSpacesPropertiesVmwareStorageProfilePtrOutputWithContext(context.Background())
}

func (i *storageSpacesPropertiesVmwareStorageProfilePtrType) ToStorageSpacesPropertiesVmwareStorageProfilePtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesVmwareStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpacesPropertiesVmwareStorageProfilePtrOutput)
}

type StorageSpacesPropertiesVmwareStorageProfileOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesVmwareStorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesPropertiesVmwareStorageProfile)(nil)).Elem()
}

func (o StorageSpacesPropertiesVmwareStorageProfileOutput) ToStorageSpacesPropertiesVmwareStorageProfileOutput() StorageSpacesPropertiesVmwareStorageProfileOutput {
	return o
}

func (o StorageSpacesPropertiesVmwareStorageProfileOutput) ToStorageSpacesPropertiesVmwareStorageProfileOutputWithContext(ctx context.Context) StorageSpacesPropertiesVmwareStorageProfileOutput {
	return o
}

func (o StorageSpacesPropertiesVmwareStorageProfileOutput) ToStorageSpacesPropertiesVmwareStorageProfilePtrOutput() StorageSpacesPropertiesVmwareStorageProfilePtrOutput {
	return o.ToStorageSpacesPropertiesVmwareStorageProfilePtrOutputWithContext(context.Background())
}

func (o StorageSpacesPropertiesVmwareStorageProfileOutput) ToStorageSpacesPropertiesVmwareStorageProfilePtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesVmwareStorageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageSpacesPropertiesVmwareStorageProfile) *StorageSpacesPropertiesVmwareStorageProfile {
		return &v
	}).(StorageSpacesPropertiesVmwareStorageProfilePtrOutput)
}

func (o StorageSpacesPropertiesVmwareStorageProfileOutput) Datacenter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesVmwareStorageProfile) *string { return v.Datacenter }).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesVmwareStorageProfileOutput) Datastore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesVmwareStorageProfile) *string { return v.Datastore }).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesVmwareStorageProfileOutput) Folder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesVmwareStorageProfile) *string { return v.Folder }).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesVmwareStorageProfileOutput) ResourcePool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesPropertiesVmwareStorageProfile) *string { return v.ResourcePool }).(pulumi.StringPtrOutput)
}

type StorageSpacesPropertiesVmwareStorageProfilePtrOutput struct{ *pulumi.OutputState }

func (StorageSpacesPropertiesVmwareStorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpacesPropertiesVmwareStorageProfile)(nil)).Elem()
}

func (o StorageSpacesPropertiesVmwareStorageProfilePtrOutput) ToStorageSpacesPropertiesVmwareStorageProfilePtrOutput() StorageSpacesPropertiesVmwareStorageProfilePtrOutput {
	return o
}

func (o StorageSpacesPropertiesVmwareStorageProfilePtrOutput) ToStorageSpacesPropertiesVmwareStorageProfilePtrOutputWithContext(ctx context.Context) StorageSpacesPropertiesVmwareStorageProfilePtrOutput {
	return o
}

func (o StorageSpacesPropertiesVmwareStorageProfilePtrOutput) Elem() StorageSpacesPropertiesVmwareStorageProfileOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesVmwareStorageProfile) StorageSpacesPropertiesVmwareStorageProfile {
		if v != nil {
			return *v
		}
		var ret StorageSpacesPropertiesVmwareStorageProfile
		return ret
	}).(StorageSpacesPropertiesVmwareStorageProfileOutput)
}

func (o StorageSpacesPropertiesVmwareStorageProfilePtrOutput) Datacenter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesVmwareStorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.Datacenter
	}).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesVmwareStorageProfilePtrOutput) Datastore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesVmwareStorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.Datastore
	}).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesVmwareStorageProfilePtrOutput) Folder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesVmwareStorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.Folder
	}).(pulumi.StringPtrOutput)
}

func (o StorageSpacesPropertiesVmwareStorageProfilePtrOutput) ResourcePool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesPropertiesVmwareStorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.ResourcePool
	}).(pulumi.StringPtrOutput)
}

type StorageSpacesResponseExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type StorageSpacesResponseExtendedLocationOutput struct{ *pulumi.OutputState }

func (StorageSpacesResponseExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageSpacesResponseExtendedLocation)(nil)).Elem()
}

func (o StorageSpacesResponseExtendedLocationOutput) ToStorageSpacesResponseExtendedLocationOutput() StorageSpacesResponseExtendedLocationOutput {
	return o
}

func (o StorageSpacesResponseExtendedLocationOutput) ToStorageSpacesResponseExtendedLocationOutputWithContext(ctx context.Context) StorageSpacesResponseExtendedLocationOutput {
	return o
}

func (o StorageSpacesResponseExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesResponseExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o StorageSpacesResponseExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageSpacesResponseExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type StorageSpacesResponseExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (StorageSpacesResponseExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpacesResponseExtendedLocation)(nil)).Elem()
}

func (o StorageSpacesResponseExtendedLocationPtrOutput) ToStorageSpacesResponseExtendedLocationPtrOutput() StorageSpacesResponseExtendedLocationPtrOutput {
	return o
}

func (o StorageSpacesResponseExtendedLocationPtrOutput) ToStorageSpacesResponseExtendedLocationPtrOutputWithContext(ctx context.Context) StorageSpacesResponseExtendedLocationPtrOutput {
	return o
}

func (o StorageSpacesResponseExtendedLocationPtrOutput) Elem() StorageSpacesResponseExtendedLocationOutput {
	return o.ApplyT(func(v *StorageSpacesResponseExtendedLocation) StorageSpacesResponseExtendedLocation {
		if v != nil {
			return *v
		}
		var ret StorageSpacesResponseExtendedLocation
		return ret
	}).(StorageSpacesResponseExtendedLocationOutput)
}

func (o StorageSpacesResponseExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesResponseExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o StorageSpacesResponseExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSpacesResponseExtendedLocation) *string {
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

type VirtualNetworksExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type VirtualNetworksExtendedLocationInput interface {
	pulumi.Input

	ToVirtualNetworksExtendedLocationOutput() VirtualNetworksExtendedLocationOutput
	ToVirtualNetworksExtendedLocationOutputWithContext(context.Context) VirtualNetworksExtendedLocationOutput
}

type VirtualNetworksExtendedLocationArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (VirtualNetworksExtendedLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksExtendedLocation)(nil)).Elem()
}

func (i VirtualNetworksExtendedLocationArgs) ToVirtualNetworksExtendedLocationOutput() VirtualNetworksExtendedLocationOutput {
	return i.ToVirtualNetworksExtendedLocationOutputWithContext(context.Background())
}

func (i VirtualNetworksExtendedLocationArgs) ToVirtualNetworksExtendedLocationOutputWithContext(ctx context.Context) VirtualNetworksExtendedLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksExtendedLocationOutput)
}

func (i VirtualNetworksExtendedLocationArgs) ToVirtualNetworksExtendedLocationPtrOutput() VirtualNetworksExtendedLocationPtrOutput {
	return i.ToVirtualNetworksExtendedLocationPtrOutputWithContext(context.Background())
}

func (i VirtualNetworksExtendedLocationArgs) ToVirtualNetworksExtendedLocationPtrOutputWithContext(ctx context.Context) VirtualNetworksExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksExtendedLocationOutput).ToVirtualNetworksExtendedLocationPtrOutputWithContext(ctx)
}









type VirtualNetworksExtendedLocationPtrInput interface {
	pulumi.Input

	ToVirtualNetworksExtendedLocationPtrOutput() VirtualNetworksExtendedLocationPtrOutput
	ToVirtualNetworksExtendedLocationPtrOutputWithContext(context.Context) VirtualNetworksExtendedLocationPtrOutput
}

type virtualNetworksExtendedLocationPtrType VirtualNetworksExtendedLocationArgs

func VirtualNetworksExtendedLocationPtr(v *VirtualNetworksExtendedLocationArgs) VirtualNetworksExtendedLocationPtrInput {
	return (*virtualNetworksExtendedLocationPtrType)(v)
}

func (*virtualNetworksExtendedLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworksExtendedLocation)(nil)).Elem()
}

func (i *virtualNetworksExtendedLocationPtrType) ToVirtualNetworksExtendedLocationPtrOutput() VirtualNetworksExtendedLocationPtrOutput {
	return i.ToVirtualNetworksExtendedLocationPtrOutputWithContext(context.Background())
}

func (i *virtualNetworksExtendedLocationPtrType) ToVirtualNetworksExtendedLocationPtrOutputWithContext(ctx context.Context) VirtualNetworksExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksExtendedLocationPtrOutput)
}

type VirtualNetworksExtendedLocationOutput struct{ *pulumi.OutputState }

func (VirtualNetworksExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksExtendedLocation)(nil)).Elem()
}

func (o VirtualNetworksExtendedLocationOutput) ToVirtualNetworksExtendedLocationOutput() VirtualNetworksExtendedLocationOutput {
	return o
}

func (o VirtualNetworksExtendedLocationOutput) ToVirtualNetworksExtendedLocationOutputWithContext(ctx context.Context) VirtualNetworksExtendedLocationOutput {
	return o
}

func (o VirtualNetworksExtendedLocationOutput) ToVirtualNetworksExtendedLocationPtrOutput() VirtualNetworksExtendedLocationPtrOutput {
	return o.ToVirtualNetworksExtendedLocationPtrOutputWithContext(context.Background())
}

func (o VirtualNetworksExtendedLocationOutput) ToVirtualNetworksExtendedLocationPtrOutputWithContext(ctx context.Context) VirtualNetworksExtendedLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworksExtendedLocation) *VirtualNetworksExtendedLocation {
		return &v
	}).(VirtualNetworksExtendedLocationPtrOutput)
}

func (o VirtualNetworksExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VirtualNetworksExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworksExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworksExtendedLocation)(nil)).Elem()
}

func (o VirtualNetworksExtendedLocationPtrOutput) ToVirtualNetworksExtendedLocationPtrOutput() VirtualNetworksExtendedLocationPtrOutput {
	return o
}

func (o VirtualNetworksExtendedLocationPtrOutput) ToVirtualNetworksExtendedLocationPtrOutputWithContext(ctx context.Context) VirtualNetworksExtendedLocationPtrOutput {
	return o
}

func (o VirtualNetworksExtendedLocationPtrOutput) Elem() VirtualNetworksExtendedLocationOutput {
	return o.ApplyT(func(v *VirtualNetworksExtendedLocation) VirtualNetworksExtendedLocation {
		if v != nil {
			return *v
		}
		var ret VirtualNetworksExtendedLocation
		return ret
	}).(VirtualNetworksExtendedLocationOutput)
}

func (o VirtualNetworksExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworksProperties struct {
	InfraVnetProfile *VirtualNetworksPropertiesInfraVnetProfile `pulumi:"infraVnetProfile"`
	VipPool          []VirtualNetworksPropertiesVipPool         `pulumi:"vipPool"`
	VmipPool         []VirtualNetworksPropertiesVmipPool        `pulumi:"vmipPool"`
}





type VirtualNetworksPropertiesInput interface {
	pulumi.Input

	ToVirtualNetworksPropertiesOutput() VirtualNetworksPropertiesOutput
	ToVirtualNetworksPropertiesOutputWithContext(context.Context) VirtualNetworksPropertiesOutput
}

type VirtualNetworksPropertiesArgs struct {
	InfraVnetProfile VirtualNetworksPropertiesInfraVnetProfilePtrInput `pulumi:"infraVnetProfile"`
	VipPool          VirtualNetworksPropertiesVipPoolArrayInput        `pulumi:"vipPool"`
	VmipPool         VirtualNetworksPropertiesVmipPoolArrayInput       `pulumi:"vmipPool"`
}

func (VirtualNetworksPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksProperties)(nil)).Elem()
}

func (i VirtualNetworksPropertiesArgs) ToVirtualNetworksPropertiesOutput() VirtualNetworksPropertiesOutput {
	return i.ToVirtualNetworksPropertiesOutputWithContext(context.Background())
}

func (i VirtualNetworksPropertiesArgs) ToVirtualNetworksPropertiesOutputWithContext(ctx context.Context) VirtualNetworksPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksPropertiesOutput)
}

func (i VirtualNetworksPropertiesArgs) ToVirtualNetworksPropertiesPtrOutput() VirtualNetworksPropertiesPtrOutput {
	return i.ToVirtualNetworksPropertiesPtrOutputWithContext(context.Background())
}

func (i VirtualNetworksPropertiesArgs) ToVirtualNetworksPropertiesPtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksPropertiesOutput).ToVirtualNetworksPropertiesPtrOutputWithContext(ctx)
}









type VirtualNetworksPropertiesPtrInput interface {
	pulumi.Input

	ToVirtualNetworksPropertiesPtrOutput() VirtualNetworksPropertiesPtrOutput
	ToVirtualNetworksPropertiesPtrOutputWithContext(context.Context) VirtualNetworksPropertiesPtrOutput
}

type virtualNetworksPropertiesPtrType VirtualNetworksPropertiesArgs

func VirtualNetworksPropertiesPtr(v *VirtualNetworksPropertiesArgs) VirtualNetworksPropertiesPtrInput {
	return (*virtualNetworksPropertiesPtrType)(v)
}

func (*virtualNetworksPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworksProperties)(nil)).Elem()
}

func (i *virtualNetworksPropertiesPtrType) ToVirtualNetworksPropertiesPtrOutput() VirtualNetworksPropertiesPtrOutput {
	return i.ToVirtualNetworksPropertiesPtrOutputWithContext(context.Background())
}

func (i *virtualNetworksPropertiesPtrType) ToVirtualNetworksPropertiesPtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksPropertiesPtrOutput)
}

type VirtualNetworksPropertiesOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksProperties)(nil)).Elem()
}

func (o VirtualNetworksPropertiesOutput) ToVirtualNetworksPropertiesOutput() VirtualNetworksPropertiesOutput {
	return o
}

func (o VirtualNetworksPropertiesOutput) ToVirtualNetworksPropertiesOutputWithContext(ctx context.Context) VirtualNetworksPropertiesOutput {
	return o
}

func (o VirtualNetworksPropertiesOutput) ToVirtualNetworksPropertiesPtrOutput() VirtualNetworksPropertiesPtrOutput {
	return o.ToVirtualNetworksPropertiesPtrOutputWithContext(context.Background())
}

func (o VirtualNetworksPropertiesOutput) ToVirtualNetworksPropertiesPtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworksProperties) *VirtualNetworksProperties {
		return &v
	}).(VirtualNetworksPropertiesPtrOutput)
}

func (o VirtualNetworksPropertiesOutput) InfraVnetProfile() VirtualNetworksPropertiesInfraVnetProfilePtrOutput {
	return o.ApplyT(func(v VirtualNetworksProperties) *VirtualNetworksPropertiesInfraVnetProfile {
		return v.InfraVnetProfile
	}).(VirtualNetworksPropertiesInfraVnetProfilePtrOutput)
}

func (o VirtualNetworksPropertiesOutput) VipPool() VirtualNetworksPropertiesVipPoolArrayOutput {
	return o.ApplyT(func(v VirtualNetworksProperties) []VirtualNetworksPropertiesVipPool { return v.VipPool }).(VirtualNetworksPropertiesVipPoolArrayOutput)
}

func (o VirtualNetworksPropertiesOutput) VmipPool() VirtualNetworksPropertiesVmipPoolArrayOutput {
	return o.ApplyT(func(v VirtualNetworksProperties) []VirtualNetworksPropertiesVmipPool { return v.VmipPool }).(VirtualNetworksPropertiesVmipPoolArrayOutput)
}

type VirtualNetworksPropertiesPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworksProperties)(nil)).Elem()
}

func (o VirtualNetworksPropertiesPtrOutput) ToVirtualNetworksPropertiesPtrOutput() VirtualNetworksPropertiesPtrOutput {
	return o
}

func (o VirtualNetworksPropertiesPtrOutput) ToVirtualNetworksPropertiesPtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesPtrOutput {
	return o
}

func (o VirtualNetworksPropertiesPtrOutput) Elem() VirtualNetworksPropertiesOutput {
	return o.ApplyT(func(v *VirtualNetworksProperties) VirtualNetworksProperties {
		if v != nil {
			return *v
		}
		var ret VirtualNetworksProperties
		return ret
	}).(VirtualNetworksPropertiesOutput)
}

func (o VirtualNetworksPropertiesPtrOutput) InfraVnetProfile() VirtualNetworksPropertiesInfraVnetProfilePtrOutput {
	return o.ApplyT(func(v *VirtualNetworksProperties) *VirtualNetworksPropertiesInfraVnetProfile {
		if v == nil {
			return nil
		}
		return v.InfraVnetProfile
	}).(VirtualNetworksPropertiesInfraVnetProfilePtrOutput)
}

func (o VirtualNetworksPropertiesPtrOutput) VipPool() VirtualNetworksPropertiesVipPoolArrayOutput {
	return o.ApplyT(func(v *VirtualNetworksProperties) []VirtualNetworksPropertiesVipPool {
		if v == nil {
			return nil
		}
		return v.VipPool
	}).(VirtualNetworksPropertiesVipPoolArrayOutput)
}

func (o VirtualNetworksPropertiesPtrOutput) VmipPool() VirtualNetworksPropertiesVmipPoolArrayOutput {
	return o.ApplyT(func(v *VirtualNetworksProperties) []VirtualNetworksPropertiesVmipPool {
		if v == nil {
			return nil
		}
		return v.VmipPool
	}).(VirtualNetworksPropertiesVmipPoolArrayOutput)
}

type VirtualNetworksPropertiesHci struct {
	MocGroup    *string `pulumi:"mocGroup"`
	MocLocation *string `pulumi:"mocLocation"`
	MocVnetName *string `pulumi:"mocVnetName"`
}





type VirtualNetworksPropertiesHciInput interface {
	pulumi.Input

	ToVirtualNetworksPropertiesHciOutput() VirtualNetworksPropertiesHciOutput
	ToVirtualNetworksPropertiesHciOutputWithContext(context.Context) VirtualNetworksPropertiesHciOutput
}

type VirtualNetworksPropertiesHciArgs struct {
	MocGroup    pulumi.StringPtrInput `pulumi:"mocGroup"`
	MocLocation pulumi.StringPtrInput `pulumi:"mocLocation"`
	MocVnetName pulumi.StringPtrInput `pulumi:"mocVnetName"`
}

func (VirtualNetworksPropertiesHciArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesHci)(nil)).Elem()
}

func (i VirtualNetworksPropertiesHciArgs) ToVirtualNetworksPropertiesHciOutput() VirtualNetworksPropertiesHciOutput {
	return i.ToVirtualNetworksPropertiesHciOutputWithContext(context.Background())
}

func (i VirtualNetworksPropertiesHciArgs) ToVirtualNetworksPropertiesHciOutputWithContext(ctx context.Context) VirtualNetworksPropertiesHciOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksPropertiesHciOutput)
}

func (i VirtualNetworksPropertiesHciArgs) ToVirtualNetworksPropertiesHciPtrOutput() VirtualNetworksPropertiesHciPtrOutput {
	return i.ToVirtualNetworksPropertiesHciPtrOutputWithContext(context.Background())
}

func (i VirtualNetworksPropertiesHciArgs) ToVirtualNetworksPropertiesHciPtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesHciPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksPropertiesHciOutput).ToVirtualNetworksPropertiesHciPtrOutputWithContext(ctx)
}









type VirtualNetworksPropertiesHciPtrInput interface {
	pulumi.Input

	ToVirtualNetworksPropertiesHciPtrOutput() VirtualNetworksPropertiesHciPtrOutput
	ToVirtualNetworksPropertiesHciPtrOutputWithContext(context.Context) VirtualNetworksPropertiesHciPtrOutput
}

type virtualNetworksPropertiesHciPtrType VirtualNetworksPropertiesHciArgs

func VirtualNetworksPropertiesHciPtr(v *VirtualNetworksPropertiesHciArgs) VirtualNetworksPropertiesHciPtrInput {
	return (*virtualNetworksPropertiesHciPtrType)(v)
}

func (*virtualNetworksPropertiesHciPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworksPropertiesHci)(nil)).Elem()
}

func (i *virtualNetworksPropertiesHciPtrType) ToVirtualNetworksPropertiesHciPtrOutput() VirtualNetworksPropertiesHciPtrOutput {
	return i.ToVirtualNetworksPropertiesHciPtrOutputWithContext(context.Background())
}

func (i *virtualNetworksPropertiesHciPtrType) ToVirtualNetworksPropertiesHciPtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesHciPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksPropertiesHciPtrOutput)
}

type VirtualNetworksPropertiesHciOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesHciOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesHci)(nil)).Elem()
}

func (o VirtualNetworksPropertiesHciOutput) ToVirtualNetworksPropertiesHciOutput() VirtualNetworksPropertiesHciOutput {
	return o
}

func (o VirtualNetworksPropertiesHciOutput) ToVirtualNetworksPropertiesHciOutputWithContext(ctx context.Context) VirtualNetworksPropertiesHciOutput {
	return o
}

func (o VirtualNetworksPropertiesHciOutput) ToVirtualNetworksPropertiesHciPtrOutput() VirtualNetworksPropertiesHciPtrOutput {
	return o.ToVirtualNetworksPropertiesHciPtrOutputWithContext(context.Background())
}

func (o VirtualNetworksPropertiesHciOutput) ToVirtualNetworksPropertiesHciPtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesHciPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworksPropertiesHci) *VirtualNetworksPropertiesHci {
		return &v
	}).(VirtualNetworksPropertiesHciPtrOutput)
}

func (o VirtualNetworksPropertiesHciOutput) MocGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesHci) *string { return v.MocGroup }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksPropertiesHciOutput) MocLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesHci) *string { return v.MocLocation }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksPropertiesHciOutput) MocVnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesHci) *string { return v.MocVnetName }).(pulumi.StringPtrOutput)
}

type VirtualNetworksPropertiesHciPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesHciPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworksPropertiesHci)(nil)).Elem()
}

func (o VirtualNetworksPropertiesHciPtrOutput) ToVirtualNetworksPropertiesHciPtrOutput() VirtualNetworksPropertiesHciPtrOutput {
	return o
}

func (o VirtualNetworksPropertiesHciPtrOutput) ToVirtualNetworksPropertiesHciPtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesHciPtrOutput {
	return o
}

func (o VirtualNetworksPropertiesHciPtrOutput) Elem() VirtualNetworksPropertiesHciOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesHci) VirtualNetworksPropertiesHci {
		if v != nil {
			return *v
		}
		var ret VirtualNetworksPropertiesHci
		return ret
	}).(VirtualNetworksPropertiesHciOutput)
}

func (o VirtualNetworksPropertiesHciPtrOutput) MocGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesHci) *string {
		if v == nil {
			return nil
		}
		return v.MocGroup
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksPropertiesHciPtrOutput) MocLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesHci) *string {
		if v == nil {
			return nil
		}
		return v.MocLocation
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksPropertiesHciPtrOutput) MocVnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesHci) *string {
		if v == nil {
			return nil
		}
		return v.MocVnetName
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworksPropertiesInfraVnetProfile struct {
	Hci      *VirtualNetworksPropertiesHci      `pulumi:"hci"`
	Kubevirt *VirtualNetworksPropertiesKubevirt `pulumi:"kubevirt"`
	Vmware   *VirtualNetworksPropertiesVmware   `pulumi:"vmware"`
}





type VirtualNetworksPropertiesInfraVnetProfileInput interface {
	pulumi.Input

	ToVirtualNetworksPropertiesInfraVnetProfileOutput() VirtualNetworksPropertiesInfraVnetProfileOutput
	ToVirtualNetworksPropertiesInfraVnetProfileOutputWithContext(context.Context) VirtualNetworksPropertiesInfraVnetProfileOutput
}

type VirtualNetworksPropertiesInfraVnetProfileArgs struct {
	Hci      VirtualNetworksPropertiesHciPtrInput      `pulumi:"hci"`
	Kubevirt VirtualNetworksPropertiesKubevirtPtrInput `pulumi:"kubevirt"`
	Vmware   VirtualNetworksPropertiesVmwarePtrInput   `pulumi:"vmware"`
}

func (VirtualNetworksPropertiesInfraVnetProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesInfraVnetProfile)(nil)).Elem()
}

func (i VirtualNetworksPropertiesInfraVnetProfileArgs) ToVirtualNetworksPropertiesInfraVnetProfileOutput() VirtualNetworksPropertiesInfraVnetProfileOutput {
	return i.ToVirtualNetworksPropertiesInfraVnetProfileOutputWithContext(context.Background())
}

func (i VirtualNetworksPropertiesInfraVnetProfileArgs) ToVirtualNetworksPropertiesInfraVnetProfileOutputWithContext(ctx context.Context) VirtualNetworksPropertiesInfraVnetProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksPropertiesInfraVnetProfileOutput)
}

func (i VirtualNetworksPropertiesInfraVnetProfileArgs) ToVirtualNetworksPropertiesInfraVnetProfilePtrOutput() VirtualNetworksPropertiesInfraVnetProfilePtrOutput {
	return i.ToVirtualNetworksPropertiesInfraVnetProfilePtrOutputWithContext(context.Background())
}

func (i VirtualNetworksPropertiesInfraVnetProfileArgs) ToVirtualNetworksPropertiesInfraVnetProfilePtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesInfraVnetProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksPropertiesInfraVnetProfileOutput).ToVirtualNetworksPropertiesInfraVnetProfilePtrOutputWithContext(ctx)
}









type VirtualNetworksPropertiesInfraVnetProfilePtrInput interface {
	pulumi.Input

	ToVirtualNetworksPropertiesInfraVnetProfilePtrOutput() VirtualNetworksPropertiesInfraVnetProfilePtrOutput
	ToVirtualNetworksPropertiesInfraVnetProfilePtrOutputWithContext(context.Context) VirtualNetworksPropertiesInfraVnetProfilePtrOutput
}

type virtualNetworksPropertiesInfraVnetProfilePtrType VirtualNetworksPropertiesInfraVnetProfileArgs

func VirtualNetworksPropertiesInfraVnetProfilePtr(v *VirtualNetworksPropertiesInfraVnetProfileArgs) VirtualNetworksPropertiesInfraVnetProfilePtrInput {
	return (*virtualNetworksPropertiesInfraVnetProfilePtrType)(v)
}

func (*virtualNetworksPropertiesInfraVnetProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworksPropertiesInfraVnetProfile)(nil)).Elem()
}

func (i *virtualNetworksPropertiesInfraVnetProfilePtrType) ToVirtualNetworksPropertiesInfraVnetProfilePtrOutput() VirtualNetworksPropertiesInfraVnetProfilePtrOutput {
	return i.ToVirtualNetworksPropertiesInfraVnetProfilePtrOutputWithContext(context.Background())
}

func (i *virtualNetworksPropertiesInfraVnetProfilePtrType) ToVirtualNetworksPropertiesInfraVnetProfilePtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesInfraVnetProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksPropertiesInfraVnetProfilePtrOutput)
}

type VirtualNetworksPropertiesInfraVnetProfileOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesInfraVnetProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesInfraVnetProfile)(nil)).Elem()
}

func (o VirtualNetworksPropertiesInfraVnetProfileOutput) ToVirtualNetworksPropertiesInfraVnetProfileOutput() VirtualNetworksPropertiesInfraVnetProfileOutput {
	return o
}

func (o VirtualNetworksPropertiesInfraVnetProfileOutput) ToVirtualNetworksPropertiesInfraVnetProfileOutputWithContext(ctx context.Context) VirtualNetworksPropertiesInfraVnetProfileOutput {
	return o
}

func (o VirtualNetworksPropertiesInfraVnetProfileOutput) ToVirtualNetworksPropertiesInfraVnetProfilePtrOutput() VirtualNetworksPropertiesInfraVnetProfilePtrOutput {
	return o.ToVirtualNetworksPropertiesInfraVnetProfilePtrOutputWithContext(context.Background())
}

func (o VirtualNetworksPropertiesInfraVnetProfileOutput) ToVirtualNetworksPropertiesInfraVnetProfilePtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesInfraVnetProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworksPropertiesInfraVnetProfile) *VirtualNetworksPropertiesInfraVnetProfile {
		return &v
	}).(VirtualNetworksPropertiesInfraVnetProfilePtrOutput)
}

func (o VirtualNetworksPropertiesInfraVnetProfileOutput) Hci() VirtualNetworksPropertiesHciPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesInfraVnetProfile) *VirtualNetworksPropertiesHci { return v.Hci }).(VirtualNetworksPropertiesHciPtrOutput)
}

func (o VirtualNetworksPropertiesInfraVnetProfileOutput) Kubevirt() VirtualNetworksPropertiesKubevirtPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesInfraVnetProfile) *VirtualNetworksPropertiesKubevirt {
		return v.Kubevirt
	}).(VirtualNetworksPropertiesKubevirtPtrOutput)
}

func (o VirtualNetworksPropertiesInfraVnetProfileOutput) Vmware() VirtualNetworksPropertiesVmwarePtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesInfraVnetProfile) *VirtualNetworksPropertiesVmware { return v.Vmware }).(VirtualNetworksPropertiesVmwarePtrOutput)
}

type VirtualNetworksPropertiesInfraVnetProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesInfraVnetProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworksPropertiesInfraVnetProfile)(nil)).Elem()
}

func (o VirtualNetworksPropertiesInfraVnetProfilePtrOutput) ToVirtualNetworksPropertiesInfraVnetProfilePtrOutput() VirtualNetworksPropertiesInfraVnetProfilePtrOutput {
	return o
}

func (o VirtualNetworksPropertiesInfraVnetProfilePtrOutput) ToVirtualNetworksPropertiesInfraVnetProfilePtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesInfraVnetProfilePtrOutput {
	return o
}

func (o VirtualNetworksPropertiesInfraVnetProfilePtrOutput) Elem() VirtualNetworksPropertiesInfraVnetProfileOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesInfraVnetProfile) VirtualNetworksPropertiesInfraVnetProfile {
		if v != nil {
			return *v
		}
		var ret VirtualNetworksPropertiesInfraVnetProfile
		return ret
	}).(VirtualNetworksPropertiesInfraVnetProfileOutput)
}

func (o VirtualNetworksPropertiesInfraVnetProfilePtrOutput) Hci() VirtualNetworksPropertiesHciPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesInfraVnetProfile) *VirtualNetworksPropertiesHci {
		if v == nil {
			return nil
		}
		return v.Hci
	}).(VirtualNetworksPropertiesHciPtrOutput)
}

func (o VirtualNetworksPropertiesInfraVnetProfilePtrOutput) Kubevirt() VirtualNetworksPropertiesKubevirtPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesInfraVnetProfile) *VirtualNetworksPropertiesKubevirt {
		if v == nil {
			return nil
		}
		return v.Kubevirt
	}).(VirtualNetworksPropertiesKubevirtPtrOutput)
}

func (o VirtualNetworksPropertiesInfraVnetProfilePtrOutput) Vmware() VirtualNetworksPropertiesVmwarePtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesInfraVnetProfile) *VirtualNetworksPropertiesVmware {
		if v == nil {
			return nil
		}
		return v.Vmware
	}).(VirtualNetworksPropertiesVmwarePtrOutput)
}

type VirtualNetworksPropertiesKubevirt struct {
	VnetName *string `pulumi:"vnetName"`
}





type VirtualNetworksPropertiesKubevirtInput interface {
	pulumi.Input

	ToVirtualNetworksPropertiesKubevirtOutput() VirtualNetworksPropertiesKubevirtOutput
	ToVirtualNetworksPropertiesKubevirtOutputWithContext(context.Context) VirtualNetworksPropertiesKubevirtOutput
}

type VirtualNetworksPropertiesKubevirtArgs struct {
	VnetName pulumi.StringPtrInput `pulumi:"vnetName"`
}

func (VirtualNetworksPropertiesKubevirtArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesKubevirt)(nil)).Elem()
}

func (i VirtualNetworksPropertiesKubevirtArgs) ToVirtualNetworksPropertiesKubevirtOutput() VirtualNetworksPropertiesKubevirtOutput {
	return i.ToVirtualNetworksPropertiesKubevirtOutputWithContext(context.Background())
}

func (i VirtualNetworksPropertiesKubevirtArgs) ToVirtualNetworksPropertiesKubevirtOutputWithContext(ctx context.Context) VirtualNetworksPropertiesKubevirtOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksPropertiesKubevirtOutput)
}

func (i VirtualNetworksPropertiesKubevirtArgs) ToVirtualNetworksPropertiesKubevirtPtrOutput() VirtualNetworksPropertiesKubevirtPtrOutput {
	return i.ToVirtualNetworksPropertiesKubevirtPtrOutputWithContext(context.Background())
}

func (i VirtualNetworksPropertiesKubevirtArgs) ToVirtualNetworksPropertiesKubevirtPtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesKubevirtPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksPropertiesKubevirtOutput).ToVirtualNetworksPropertiesKubevirtPtrOutputWithContext(ctx)
}









type VirtualNetworksPropertiesKubevirtPtrInput interface {
	pulumi.Input

	ToVirtualNetworksPropertiesKubevirtPtrOutput() VirtualNetworksPropertiesKubevirtPtrOutput
	ToVirtualNetworksPropertiesKubevirtPtrOutputWithContext(context.Context) VirtualNetworksPropertiesKubevirtPtrOutput
}

type virtualNetworksPropertiesKubevirtPtrType VirtualNetworksPropertiesKubevirtArgs

func VirtualNetworksPropertiesKubevirtPtr(v *VirtualNetworksPropertiesKubevirtArgs) VirtualNetworksPropertiesKubevirtPtrInput {
	return (*virtualNetworksPropertiesKubevirtPtrType)(v)
}

func (*virtualNetworksPropertiesKubevirtPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworksPropertiesKubevirt)(nil)).Elem()
}

func (i *virtualNetworksPropertiesKubevirtPtrType) ToVirtualNetworksPropertiesKubevirtPtrOutput() VirtualNetworksPropertiesKubevirtPtrOutput {
	return i.ToVirtualNetworksPropertiesKubevirtPtrOutputWithContext(context.Background())
}

func (i *virtualNetworksPropertiesKubevirtPtrType) ToVirtualNetworksPropertiesKubevirtPtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesKubevirtPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksPropertiesKubevirtPtrOutput)
}

type VirtualNetworksPropertiesKubevirtOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesKubevirtOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesKubevirt)(nil)).Elem()
}

func (o VirtualNetworksPropertiesKubevirtOutput) ToVirtualNetworksPropertiesKubevirtOutput() VirtualNetworksPropertiesKubevirtOutput {
	return o
}

func (o VirtualNetworksPropertiesKubevirtOutput) ToVirtualNetworksPropertiesKubevirtOutputWithContext(ctx context.Context) VirtualNetworksPropertiesKubevirtOutput {
	return o
}

func (o VirtualNetworksPropertiesKubevirtOutput) ToVirtualNetworksPropertiesKubevirtPtrOutput() VirtualNetworksPropertiesKubevirtPtrOutput {
	return o.ToVirtualNetworksPropertiesKubevirtPtrOutputWithContext(context.Background())
}

func (o VirtualNetworksPropertiesKubevirtOutput) ToVirtualNetworksPropertiesKubevirtPtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesKubevirtPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworksPropertiesKubevirt) *VirtualNetworksPropertiesKubevirt {
		return &v
	}).(VirtualNetworksPropertiesKubevirtPtrOutput)
}

func (o VirtualNetworksPropertiesKubevirtOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesKubevirt) *string { return v.VnetName }).(pulumi.StringPtrOutput)
}

type VirtualNetworksPropertiesKubevirtPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesKubevirtPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworksPropertiesKubevirt)(nil)).Elem()
}

func (o VirtualNetworksPropertiesKubevirtPtrOutput) ToVirtualNetworksPropertiesKubevirtPtrOutput() VirtualNetworksPropertiesKubevirtPtrOutput {
	return o
}

func (o VirtualNetworksPropertiesKubevirtPtrOutput) ToVirtualNetworksPropertiesKubevirtPtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesKubevirtPtrOutput {
	return o
}

func (o VirtualNetworksPropertiesKubevirtPtrOutput) Elem() VirtualNetworksPropertiesKubevirtOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesKubevirt) VirtualNetworksPropertiesKubevirt {
		if v != nil {
			return *v
		}
		var ret VirtualNetworksPropertiesKubevirt
		return ret
	}).(VirtualNetworksPropertiesKubevirtOutput)
}

func (o VirtualNetworksPropertiesKubevirtPtrOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesKubevirt) *string {
		if v == nil {
			return nil
		}
		return v.VnetName
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworksPropertiesResponse struct {
	DhcpServers       []string                                           `pulumi:"dhcpServers"`
	DnsServers        []string                                           `pulumi:"dnsServers"`
	Gateway           string                                             `pulumi:"gateway"`
	InfraVnetProfile  *VirtualNetworksPropertiesResponseInfraVnetProfile `pulumi:"infraVnetProfile"`
	IpAddressPrefix   string                                             `pulumi:"ipAddressPrefix"`
	ProvisioningState string                                             `pulumi:"provisioningState"`
	Status            VirtualNetworksPropertiesResponseStatus            `pulumi:"status"`
	VipPool           []VirtualNetworksPropertiesResponseVipPool         `pulumi:"vipPool"`
	VlanID            string                                             `pulumi:"vlanID"`
	VmipPool          []VirtualNetworksPropertiesResponseVmipPool        `pulumi:"vmipPool"`
}

type VirtualNetworksPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesResponse)(nil)).Elem()
}

func (o VirtualNetworksPropertiesResponseOutput) ToVirtualNetworksPropertiesResponseOutput() VirtualNetworksPropertiesResponseOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseOutput) ToVirtualNetworksPropertiesResponseOutputWithContext(ctx context.Context) VirtualNetworksPropertiesResponseOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseOutput) DhcpServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponse) []string { return v.DhcpServers }).(pulumi.StringArrayOutput)
}

func (o VirtualNetworksPropertiesResponseOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponse) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

func (o VirtualNetworksPropertiesResponseOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponse) string { return v.Gateway }).(pulumi.StringOutput)
}

func (o VirtualNetworksPropertiesResponseOutput) InfraVnetProfile() VirtualNetworksPropertiesResponseInfraVnetProfilePtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponse) *VirtualNetworksPropertiesResponseInfraVnetProfile {
		return v.InfraVnetProfile
	}).(VirtualNetworksPropertiesResponseInfraVnetProfilePtrOutput)
}

func (o VirtualNetworksPropertiesResponseOutput) IpAddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponse) string { return v.IpAddressPrefix }).(pulumi.StringOutput)
}

func (o VirtualNetworksPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualNetworksPropertiesResponseOutput) Status() VirtualNetworksPropertiesResponseStatusOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponse) VirtualNetworksPropertiesResponseStatus { return v.Status }).(VirtualNetworksPropertiesResponseStatusOutput)
}

func (o VirtualNetworksPropertiesResponseOutput) VipPool() VirtualNetworksPropertiesResponseVipPoolArrayOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponse) []VirtualNetworksPropertiesResponseVipPool { return v.VipPool }).(VirtualNetworksPropertiesResponseVipPoolArrayOutput)
}

func (o VirtualNetworksPropertiesResponseOutput) VlanID() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponse) string { return v.VlanID }).(pulumi.StringOutput)
}

func (o VirtualNetworksPropertiesResponseOutput) VmipPool() VirtualNetworksPropertiesResponseVmipPoolArrayOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponse) []VirtualNetworksPropertiesResponseVmipPool {
		return v.VmipPool
	}).(VirtualNetworksPropertiesResponseVmipPoolArrayOutput)
}

type VirtualNetworksPropertiesResponseError struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}

type VirtualNetworksPropertiesResponseErrorOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesResponseErrorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesResponseError)(nil)).Elem()
}

func (o VirtualNetworksPropertiesResponseErrorOutput) ToVirtualNetworksPropertiesResponseErrorOutput() VirtualNetworksPropertiesResponseErrorOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseErrorOutput) ToVirtualNetworksPropertiesResponseErrorOutputWithContext(ctx context.Context) VirtualNetworksPropertiesResponseErrorOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseErrorOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponseError) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksPropertiesResponseErrorOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponseError) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type VirtualNetworksPropertiesResponseErrorPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesResponseErrorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworksPropertiesResponseError)(nil)).Elem()
}

func (o VirtualNetworksPropertiesResponseErrorPtrOutput) ToVirtualNetworksPropertiesResponseErrorPtrOutput() VirtualNetworksPropertiesResponseErrorPtrOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseErrorPtrOutput) ToVirtualNetworksPropertiesResponseErrorPtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesResponseErrorPtrOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseErrorPtrOutput) Elem() VirtualNetworksPropertiesResponseErrorOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesResponseError) VirtualNetworksPropertiesResponseError {
		if v != nil {
			return *v
		}
		var ret VirtualNetworksPropertiesResponseError
		return ret
	}).(VirtualNetworksPropertiesResponseErrorOutput)
}

func (o VirtualNetworksPropertiesResponseErrorPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesResponseError) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksPropertiesResponseErrorPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesResponseError) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworksPropertiesResponseHci struct {
	MocGroup    *string `pulumi:"mocGroup"`
	MocLocation *string `pulumi:"mocLocation"`
	MocVnetName *string `pulumi:"mocVnetName"`
}

type VirtualNetworksPropertiesResponseHciOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesResponseHciOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesResponseHci)(nil)).Elem()
}

func (o VirtualNetworksPropertiesResponseHciOutput) ToVirtualNetworksPropertiesResponseHciOutput() VirtualNetworksPropertiesResponseHciOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseHciOutput) ToVirtualNetworksPropertiesResponseHciOutputWithContext(ctx context.Context) VirtualNetworksPropertiesResponseHciOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseHciOutput) MocGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponseHci) *string { return v.MocGroup }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksPropertiesResponseHciOutput) MocLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponseHci) *string { return v.MocLocation }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksPropertiesResponseHciOutput) MocVnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponseHci) *string { return v.MocVnetName }).(pulumi.StringPtrOutput)
}

type VirtualNetworksPropertiesResponseHciPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesResponseHciPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworksPropertiesResponseHci)(nil)).Elem()
}

func (o VirtualNetworksPropertiesResponseHciPtrOutput) ToVirtualNetworksPropertiesResponseHciPtrOutput() VirtualNetworksPropertiesResponseHciPtrOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseHciPtrOutput) ToVirtualNetworksPropertiesResponseHciPtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesResponseHciPtrOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseHciPtrOutput) Elem() VirtualNetworksPropertiesResponseHciOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesResponseHci) VirtualNetworksPropertiesResponseHci {
		if v != nil {
			return *v
		}
		var ret VirtualNetworksPropertiesResponseHci
		return ret
	}).(VirtualNetworksPropertiesResponseHciOutput)
}

func (o VirtualNetworksPropertiesResponseHciPtrOutput) MocGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesResponseHci) *string {
		if v == nil {
			return nil
		}
		return v.MocGroup
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksPropertiesResponseHciPtrOutput) MocLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesResponseHci) *string {
		if v == nil {
			return nil
		}
		return v.MocLocation
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksPropertiesResponseHciPtrOutput) MocVnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesResponseHci) *string {
		if v == nil {
			return nil
		}
		return v.MocVnetName
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworksPropertiesResponseInfraVnetProfile struct {
	Hci      *VirtualNetworksPropertiesResponseHci      `pulumi:"hci"`
	Kubevirt *VirtualNetworksPropertiesResponseKubevirt `pulumi:"kubevirt"`
	Vmware   *VirtualNetworksPropertiesResponseVmware   `pulumi:"vmware"`
}

type VirtualNetworksPropertiesResponseInfraVnetProfileOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesResponseInfraVnetProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesResponseInfraVnetProfile)(nil)).Elem()
}

func (o VirtualNetworksPropertiesResponseInfraVnetProfileOutput) ToVirtualNetworksPropertiesResponseInfraVnetProfileOutput() VirtualNetworksPropertiesResponseInfraVnetProfileOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseInfraVnetProfileOutput) ToVirtualNetworksPropertiesResponseInfraVnetProfileOutputWithContext(ctx context.Context) VirtualNetworksPropertiesResponseInfraVnetProfileOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseInfraVnetProfileOutput) Hci() VirtualNetworksPropertiesResponseHciPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponseInfraVnetProfile) *VirtualNetworksPropertiesResponseHci {
		return v.Hci
	}).(VirtualNetworksPropertiesResponseHciPtrOutput)
}

func (o VirtualNetworksPropertiesResponseInfraVnetProfileOutput) Kubevirt() VirtualNetworksPropertiesResponseKubevirtPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponseInfraVnetProfile) *VirtualNetworksPropertiesResponseKubevirt {
		return v.Kubevirt
	}).(VirtualNetworksPropertiesResponseKubevirtPtrOutput)
}

func (o VirtualNetworksPropertiesResponseInfraVnetProfileOutput) Vmware() VirtualNetworksPropertiesResponseVmwarePtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponseInfraVnetProfile) *VirtualNetworksPropertiesResponseVmware {
		return v.Vmware
	}).(VirtualNetworksPropertiesResponseVmwarePtrOutput)
}

type VirtualNetworksPropertiesResponseInfraVnetProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesResponseInfraVnetProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworksPropertiesResponseInfraVnetProfile)(nil)).Elem()
}

func (o VirtualNetworksPropertiesResponseInfraVnetProfilePtrOutput) ToVirtualNetworksPropertiesResponseInfraVnetProfilePtrOutput() VirtualNetworksPropertiesResponseInfraVnetProfilePtrOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseInfraVnetProfilePtrOutput) ToVirtualNetworksPropertiesResponseInfraVnetProfilePtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesResponseInfraVnetProfilePtrOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseInfraVnetProfilePtrOutput) Elem() VirtualNetworksPropertiesResponseInfraVnetProfileOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesResponseInfraVnetProfile) VirtualNetworksPropertiesResponseInfraVnetProfile {
		if v != nil {
			return *v
		}
		var ret VirtualNetworksPropertiesResponseInfraVnetProfile
		return ret
	}).(VirtualNetworksPropertiesResponseInfraVnetProfileOutput)
}

func (o VirtualNetworksPropertiesResponseInfraVnetProfilePtrOutput) Hci() VirtualNetworksPropertiesResponseHciPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesResponseInfraVnetProfile) *VirtualNetworksPropertiesResponseHci {
		if v == nil {
			return nil
		}
		return v.Hci
	}).(VirtualNetworksPropertiesResponseHciPtrOutput)
}

func (o VirtualNetworksPropertiesResponseInfraVnetProfilePtrOutput) Kubevirt() VirtualNetworksPropertiesResponseKubevirtPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesResponseInfraVnetProfile) *VirtualNetworksPropertiesResponseKubevirt {
		if v == nil {
			return nil
		}
		return v.Kubevirt
	}).(VirtualNetworksPropertiesResponseKubevirtPtrOutput)
}

func (o VirtualNetworksPropertiesResponseInfraVnetProfilePtrOutput) Vmware() VirtualNetworksPropertiesResponseVmwarePtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesResponseInfraVnetProfile) *VirtualNetworksPropertiesResponseVmware {
		if v == nil {
			return nil
		}
		return v.Vmware
	}).(VirtualNetworksPropertiesResponseVmwarePtrOutput)
}

type VirtualNetworksPropertiesResponseKubevirt struct {
	VnetName *string `pulumi:"vnetName"`
}

type VirtualNetworksPropertiesResponseKubevirtOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesResponseKubevirtOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesResponseKubevirt)(nil)).Elem()
}

func (o VirtualNetworksPropertiesResponseKubevirtOutput) ToVirtualNetworksPropertiesResponseKubevirtOutput() VirtualNetworksPropertiesResponseKubevirtOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseKubevirtOutput) ToVirtualNetworksPropertiesResponseKubevirtOutputWithContext(ctx context.Context) VirtualNetworksPropertiesResponseKubevirtOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseKubevirtOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponseKubevirt) *string { return v.VnetName }).(pulumi.StringPtrOutput)
}

type VirtualNetworksPropertiesResponseKubevirtPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesResponseKubevirtPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworksPropertiesResponseKubevirt)(nil)).Elem()
}

func (o VirtualNetworksPropertiesResponseKubevirtPtrOutput) ToVirtualNetworksPropertiesResponseKubevirtPtrOutput() VirtualNetworksPropertiesResponseKubevirtPtrOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseKubevirtPtrOutput) ToVirtualNetworksPropertiesResponseKubevirtPtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesResponseKubevirtPtrOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseKubevirtPtrOutput) Elem() VirtualNetworksPropertiesResponseKubevirtOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesResponseKubevirt) VirtualNetworksPropertiesResponseKubevirt {
		if v != nil {
			return *v
		}
		var ret VirtualNetworksPropertiesResponseKubevirt
		return ret
	}).(VirtualNetworksPropertiesResponseKubevirtOutput)
}

func (o VirtualNetworksPropertiesResponseKubevirtPtrOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesResponseKubevirt) *string {
		if v == nil {
			return nil
		}
		return v.VnetName
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworksPropertiesResponseProvisioningStatus struct {
	Error       *VirtualNetworksPropertiesResponseError `pulumi:"error"`
	OperationId *string                                 `pulumi:"operationId"`
	Phase       *string                                 `pulumi:"phase"`
	Status      *string                                 `pulumi:"status"`
}

type VirtualNetworksPropertiesResponseProvisioningStatusOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesResponseProvisioningStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesResponseProvisioningStatus)(nil)).Elem()
}

func (o VirtualNetworksPropertiesResponseProvisioningStatusOutput) ToVirtualNetworksPropertiesResponseProvisioningStatusOutput() VirtualNetworksPropertiesResponseProvisioningStatusOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseProvisioningStatusOutput) ToVirtualNetworksPropertiesResponseProvisioningStatusOutputWithContext(ctx context.Context) VirtualNetworksPropertiesResponseProvisioningStatusOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseProvisioningStatusOutput) Error() VirtualNetworksPropertiesResponseErrorPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponseProvisioningStatus) *VirtualNetworksPropertiesResponseError {
		return v.Error
	}).(VirtualNetworksPropertiesResponseErrorPtrOutput)
}

func (o VirtualNetworksPropertiesResponseProvisioningStatusOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponseProvisioningStatus) *string { return v.OperationId }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksPropertiesResponseProvisioningStatusOutput) Phase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponseProvisioningStatus) *string { return v.Phase }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksPropertiesResponseProvisioningStatusOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponseProvisioningStatus) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type VirtualNetworksPropertiesResponseProvisioningStatusPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesResponseProvisioningStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworksPropertiesResponseProvisioningStatus)(nil)).Elem()
}

func (o VirtualNetworksPropertiesResponseProvisioningStatusPtrOutput) ToVirtualNetworksPropertiesResponseProvisioningStatusPtrOutput() VirtualNetworksPropertiesResponseProvisioningStatusPtrOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseProvisioningStatusPtrOutput) ToVirtualNetworksPropertiesResponseProvisioningStatusPtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesResponseProvisioningStatusPtrOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseProvisioningStatusPtrOutput) Elem() VirtualNetworksPropertiesResponseProvisioningStatusOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesResponseProvisioningStatus) VirtualNetworksPropertiesResponseProvisioningStatus {
		if v != nil {
			return *v
		}
		var ret VirtualNetworksPropertiesResponseProvisioningStatus
		return ret
	}).(VirtualNetworksPropertiesResponseProvisioningStatusOutput)
}

func (o VirtualNetworksPropertiesResponseProvisioningStatusPtrOutput) Error() VirtualNetworksPropertiesResponseErrorPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesResponseProvisioningStatus) *VirtualNetworksPropertiesResponseError {
		if v == nil {
			return nil
		}
		return v.Error
	}).(VirtualNetworksPropertiesResponseErrorPtrOutput)
}

func (o VirtualNetworksPropertiesResponseProvisioningStatusPtrOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.OperationId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksPropertiesResponseProvisioningStatusPtrOutput) Phase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.Phase
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksPropertiesResponseProvisioningStatusPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworksPropertiesResponseStatus struct {
	ProvisioningStatus *VirtualNetworksPropertiesResponseProvisioningStatus `pulumi:"provisioningStatus"`
}

type VirtualNetworksPropertiesResponseStatusOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesResponseStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesResponseStatus)(nil)).Elem()
}

func (o VirtualNetworksPropertiesResponseStatusOutput) ToVirtualNetworksPropertiesResponseStatusOutput() VirtualNetworksPropertiesResponseStatusOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseStatusOutput) ToVirtualNetworksPropertiesResponseStatusOutputWithContext(ctx context.Context) VirtualNetworksPropertiesResponseStatusOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseStatusOutput) ProvisioningStatus() VirtualNetworksPropertiesResponseProvisioningStatusPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponseStatus) *VirtualNetworksPropertiesResponseProvisioningStatus {
		return v.ProvisioningStatus
	}).(VirtualNetworksPropertiesResponseProvisioningStatusPtrOutput)
}

type VirtualNetworksPropertiesResponseVipPool struct {
	EndIP   *string `pulumi:"endIP"`
	StartIP *string `pulumi:"startIP"`
}

type VirtualNetworksPropertiesResponseVipPoolOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesResponseVipPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesResponseVipPool)(nil)).Elem()
}

func (o VirtualNetworksPropertiesResponseVipPoolOutput) ToVirtualNetworksPropertiesResponseVipPoolOutput() VirtualNetworksPropertiesResponseVipPoolOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseVipPoolOutput) ToVirtualNetworksPropertiesResponseVipPoolOutputWithContext(ctx context.Context) VirtualNetworksPropertiesResponseVipPoolOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseVipPoolOutput) EndIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponseVipPool) *string { return v.EndIP }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksPropertiesResponseVipPoolOutput) StartIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponseVipPool) *string { return v.StartIP }).(pulumi.StringPtrOutput)
}

type VirtualNetworksPropertiesResponseVipPoolArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesResponseVipPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworksPropertiesResponseVipPool)(nil)).Elem()
}

func (o VirtualNetworksPropertiesResponseVipPoolArrayOutput) ToVirtualNetworksPropertiesResponseVipPoolArrayOutput() VirtualNetworksPropertiesResponseVipPoolArrayOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseVipPoolArrayOutput) ToVirtualNetworksPropertiesResponseVipPoolArrayOutputWithContext(ctx context.Context) VirtualNetworksPropertiesResponseVipPoolArrayOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseVipPoolArrayOutput) Index(i pulumi.IntInput) VirtualNetworksPropertiesResponseVipPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworksPropertiesResponseVipPool {
		return vs[0].([]VirtualNetworksPropertiesResponseVipPool)[vs[1].(int)]
	}).(VirtualNetworksPropertiesResponseVipPoolOutput)
}

type VirtualNetworksPropertiesResponseVmipPool struct {
	EndIP   *string `pulumi:"endIP"`
	StartIP *string `pulumi:"startIP"`
}

type VirtualNetworksPropertiesResponseVmipPoolOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesResponseVmipPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesResponseVmipPool)(nil)).Elem()
}

func (o VirtualNetworksPropertiesResponseVmipPoolOutput) ToVirtualNetworksPropertiesResponseVmipPoolOutput() VirtualNetworksPropertiesResponseVmipPoolOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseVmipPoolOutput) ToVirtualNetworksPropertiesResponseVmipPoolOutputWithContext(ctx context.Context) VirtualNetworksPropertiesResponseVmipPoolOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseVmipPoolOutput) EndIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponseVmipPool) *string { return v.EndIP }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksPropertiesResponseVmipPoolOutput) StartIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponseVmipPool) *string { return v.StartIP }).(pulumi.StringPtrOutput)
}

type VirtualNetworksPropertiesResponseVmipPoolArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesResponseVmipPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworksPropertiesResponseVmipPool)(nil)).Elem()
}

func (o VirtualNetworksPropertiesResponseVmipPoolArrayOutput) ToVirtualNetworksPropertiesResponseVmipPoolArrayOutput() VirtualNetworksPropertiesResponseVmipPoolArrayOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseVmipPoolArrayOutput) ToVirtualNetworksPropertiesResponseVmipPoolArrayOutputWithContext(ctx context.Context) VirtualNetworksPropertiesResponseVmipPoolArrayOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseVmipPoolArrayOutput) Index(i pulumi.IntInput) VirtualNetworksPropertiesResponseVmipPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworksPropertiesResponseVmipPool {
		return vs[0].([]VirtualNetworksPropertiesResponseVmipPool)[vs[1].(int)]
	}).(VirtualNetworksPropertiesResponseVmipPoolOutput)
}

type VirtualNetworksPropertiesResponseVmware struct {
	SegmentName *string `pulumi:"segmentName"`
}

type VirtualNetworksPropertiesResponseVmwareOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesResponseVmwareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesResponseVmware)(nil)).Elem()
}

func (o VirtualNetworksPropertiesResponseVmwareOutput) ToVirtualNetworksPropertiesResponseVmwareOutput() VirtualNetworksPropertiesResponseVmwareOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseVmwareOutput) ToVirtualNetworksPropertiesResponseVmwareOutputWithContext(ctx context.Context) VirtualNetworksPropertiesResponseVmwareOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseVmwareOutput) SegmentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesResponseVmware) *string { return v.SegmentName }).(pulumi.StringPtrOutput)
}

type VirtualNetworksPropertiesResponseVmwarePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesResponseVmwarePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworksPropertiesResponseVmware)(nil)).Elem()
}

func (o VirtualNetworksPropertiesResponseVmwarePtrOutput) ToVirtualNetworksPropertiesResponseVmwarePtrOutput() VirtualNetworksPropertiesResponseVmwarePtrOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseVmwarePtrOutput) ToVirtualNetworksPropertiesResponseVmwarePtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesResponseVmwarePtrOutput {
	return o
}

func (o VirtualNetworksPropertiesResponseVmwarePtrOutput) Elem() VirtualNetworksPropertiesResponseVmwareOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesResponseVmware) VirtualNetworksPropertiesResponseVmware {
		if v != nil {
			return *v
		}
		var ret VirtualNetworksPropertiesResponseVmware
		return ret
	}).(VirtualNetworksPropertiesResponseVmwareOutput)
}

func (o VirtualNetworksPropertiesResponseVmwarePtrOutput) SegmentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesResponseVmware) *string {
		if v == nil {
			return nil
		}
		return v.SegmentName
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworksPropertiesVipPool struct {
	EndIP   *string `pulumi:"endIP"`
	StartIP *string `pulumi:"startIP"`
}





type VirtualNetworksPropertiesVipPoolInput interface {
	pulumi.Input

	ToVirtualNetworksPropertiesVipPoolOutput() VirtualNetworksPropertiesVipPoolOutput
	ToVirtualNetworksPropertiesVipPoolOutputWithContext(context.Context) VirtualNetworksPropertiesVipPoolOutput
}

type VirtualNetworksPropertiesVipPoolArgs struct {
	EndIP   pulumi.StringPtrInput `pulumi:"endIP"`
	StartIP pulumi.StringPtrInput `pulumi:"startIP"`
}

func (VirtualNetworksPropertiesVipPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesVipPool)(nil)).Elem()
}

func (i VirtualNetworksPropertiesVipPoolArgs) ToVirtualNetworksPropertiesVipPoolOutput() VirtualNetworksPropertiesVipPoolOutput {
	return i.ToVirtualNetworksPropertiesVipPoolOutputWithContext(context.Background())
}

func (i VirtualNetworksPropertiesVipPoolArgs) ToVirtualNetworksPropertiesVipPoolOutputWithContext(ctx context.Context) VirtualNetworksPropertiesVipPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksPropertiesVipPoolOutput)
}





type VirtualNetworksPropertiesVipPoolArrayInput interface {
	pulumi.Input

	ToVirtualNetworksPropertiesVipPoolArrayOutput() VirtualNetworksPropertiesVipPoolArrayOutput
	ToVirtualNetworksPropertiesVipPoolArrayOutputWithContext(context.Context) VirtualNetworksPropertiesVipPoolArrayOutput
}

type VirtualNetworksPropertiesVipPoolArray []VirtualNetworksPropertiesVipPoolInput

func (VirtualNetworksPropertiesVipPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworksPropertiesVipPool)(nil)).Elem()
}

func (i VirtualNetworksPropertiesVipPoolArray) ToVirtualNetworksPropertiesVipPoolArrayOutput() VirtualNetworksPropertiesVipPoolArrayOutput {
	return i.ToVirtualNetworksPropertiesVipPoolArrayOutputWithContext(context.Background())
}

func (i VirtualNetworksPropertiesVipPoolArray) ToVirtualNetworksPropertiesVipPoolArrayOutputWithContext(ctx context.Context) VirtualNetworksPropertiesVipPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksPropertiesVipPoolArrayOutput)
}

type VirtualNetworksPropertiesVipPoolOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesVipPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesVipPool)(nil)).Elem()
}

func (o VirtualNetworksPropertiesVipPoolOutput) ToVirtualNetworksPropertiesVipPoolOutput() VirtualNetworksPropertiesVipPoolOutput {
	return o
}

func (o VirtualNetworksPropertiesVipPoolOutput) ToVirtualNetworksPropertiesVipPoolOutputWithContext(ctx context.Context) VirtualNetworksPropertiesVipPoolOutput {
	return o
}

func (o VirtualNetworksPropertiesVipPoolOutput) EndIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesVipPool) *string { return v.EndIP }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksPropertiesVipPoolOutput) StartIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesVipPool) *string { return v.StartIP }).(pulumi.StringPtrOutput)
}

type VirtualNetworksPropertiesVipPoolArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesVipPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworksPropertiesVipPool)(nil)).Elem()
}

func (o VirtualNetworksPropertiesVipPoolArrayOutput) ToVirtualNetworksPropertiesVipPoolArrayOutput() VirtualNetworksPropertiesVipPoolArrayOutput {
	return o
}

func (o VirtualNetworksPropertiesVipPoolArrayOutput) ToVirtualNetworksPropertiesVipPoolArrayOutputWithContext(ctx context.Context) VirtualNetworksPropertiesVipPoolArrayOutput {
	return o
}

func (o VirtualNetworksPropertiesVipPoolArrayOutput) Index(i pulumi.IntInput) VirtualNetworksPropertiesVipPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworksPropertiesVipPool {
		return vs[0].([]VirtualNetworksPropertiesVipPool)[vs[1].(int)]
	}).(VirtualNetworksPropertiesVipPoolOutput)
}

type VirtualNetworksPropertiesVmipPool struct {
	EndIP   *string `pulumi:"endIP"`
	StartIP *string `pulumi:"startIP"`
}





type VirtualNetworksPropertiesVmipPoolInput interface {
	pulumi.Input

	ToVirtualNetworksPropertiesVmipPoolOutput() VirtualNetworksPropertiesVmipPoolOutput
	ToVirtualNetworksPropertiesVmipPoolOutputWithContext(context.Context) VirtualNetworksPropertiesVmipPoolOutput
}

type VirtualNetworksPropertiesVmipPoolArgs struct {
	EndIP   pulumi.StringPtrInput `pulumi:"endIP"`
	StartIP pulumi.StringPtrInput `pulumi:"startIP"`
}

func (VirtualNetworksPropertiesVmipPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesVmipPool)(nil)).Elem()
}

func (i VirtualNetworksPropertiesVmipPoolArgs) ToVirtualNetworksPropertiesVmipPoolOutput() VirtualNetworksPropertiesVmipPoolOutput {
	return i.ToVirtualNetworksPropertiesVmipPoolOutputWithContext(context.Background())
}

func (i VirtualNetworksPropertiesVmipPoolArgs) ToVirtualNetworksPropertiesVmipPoolOutputWithContext(ctx context.Context) VirtualNetworksPropertiesVmipPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksPropertiesVmipPoolOutput)
}





type VirtualNetworksPropertiesVmipPoolArrayInput interface {
	pulumi.Input

	ToVirtualNetworksPropertiesVmipPoolArrayOutput() VirtualNetworksPropertiesVmipPoolArrayOutput
	ToVirtualNetworksPropertiesVmipPoolArrayOutputWithContext(context.Context) VirtualNetworksPropertiesVmipPoolArrayOutput
}

type VirtualNetworksPropertiesVmipPoolArray []VirtualNetworksPropertiesVmipPoolInput

func (VirtualNetworksPropertiesVmipPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworksPropertiesVmipPool)(nil)).Elem()
}

func (i VirtualNetworksPropertiesVmipPoolArray) ToVirtualNetworksPropertiesVmipPoolArrayOutput() VirtualNetworksPropertiesVmipPoolArrayOutput {
	return i.ToVirtualNetworksPropertiesVmipPoolArrayOutputWithContext(context.Background())
}

func (i VirtualNetworksPropertiesVmipPoolArray) ToVirtualNetworksPropertiesVmipPoolArrayOutputWithContext(ctx context.Context) VirtualNetworksPropertiesVmipPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksPropertiesVmipPoolArrayOutput)
}

type VirtualNetworksPropertiesVmipPoolOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesVmipPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesVmipPool)(nil)).Elem()
}

func (o VirtualNetworksPropertiesVmipPoolOutput) ToVirtualNetworksPropertiesVmipPoolOutput() VirtualNetworksPropertiesVmipPoolOutput {
	return o
}

func (o VirtualNetworksPropertiesVmipPoolOutput) ToVirtualNetworksPropertiesVmipPoolOutputWithContext(ctx context.Context) VirtualNetworksPropertiesVmipPoolOutput {
	return o
}

func (o VirtualNetworksPropertiesVmipPoolOutput) EndIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesVmipPool) *string { return v.EndIP }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksPropertiesVmipPoolOutput) StartIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesVmipPool) *string { return v.StartIP }).(pulumi.StringPtrOutput)
}

type VirtualNetworksPropertiesVmipPoolArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesVmipPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworksPropertiesVmipPool)(nil)).Elem()
}

func (o VirtualNetworksPropertiesVmipPoolArrayOutput) ToVirtualNetworksPropertiesVmipPoolArrayOutput() VirtualNetworksPropertiesVmipPoolArrayOutput {
	return o
}

func (o VirtualNetworksPropertiesVmipPoolArrayOutput) ToVirtualNetworksPropertiesVmipPoolArrayOutputWithContext(ctx context.Context) VirtualNetworksPropertiesVmipPoolArrayOutput {
	return o
}

func (o VirtualNetworksPropertiesVmipPoolArrayOutput) Index(i pulumi.IntInput) VirtualNetworksPropertiesVmipPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworksPropertiesVmipPool {
		return vs[0].([]VirtualNetworksPropertiesVmipPool)[vs[1].(int)]
	}).(VirtualNetworksPropertiesVmipPoolOutput)
}

type VirtualNetworksPropertiesVmware struct {
	SegmentName *string `pulumi:"segmentName"`
}





type VirtualNetworksPropertiesVmwareInput interface {
	pulumi.Input

	ToVirtualNetworksPropertiesVmwareOutput() VirtualNetworksPropertiesVmwareOutput
	ToVirtualNetworksPropertiesVmwareOutputWithContext(context.Context) VirtualNetworksPropertiesVmwareOutput
}

type VirtualNetworksPropertiesVmwareArgs struct {
	SegmentName pulumi.StringPtrInput `pulumi:"segmentName"`
}

func (VirtualNetworksPropertiesVmwareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesVmware)(nil)).Elem()
}

func (i VirtualNetworksPropertiesVmwareArgs) ToVirtualNetworksPropertiesVmwareOutput() VirtualNetworksPropertiesVmwareOutput {
	return i.ToVirtualNetworksPropertiesVmwareOutputWithContext(context.Background())
}

func (i VirtualNetworksPropertiesVmwareArgs) ToVirtualNetworksPropertiesVmwareOutputWithContext(ctx context.Context) VirtualNetworksPropertiesVmwareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksPropertiesVmwareOutput)
}

func (i VirtualNetworksPropertiesVmwareArgs) ToVirtualNetworksPropertiesVmwarePtrOutput() VirtualNetworksPropertiesVmwarePtrOutput {
	return i.ToVirtualNetworksPropertiesVmwarePtrOutputWithContext(context.Background())
}

func (i VirtualNetworksPropertiesVmwareArgs) ToVirtualNetworksPropertiesVmwarePtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesVmwarePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksPropertiesVmwareOutput).ToVirtualNetworksPropertiesVmwarePtrOutputWithContext(ctx)
}









type VirtualNetworksPropertiesVmwarePtrInput interface {
	pulumi.Input

	ToVirtualNetworksPropertiesVmwarePtrOutput() VirtualNetworksPropertiesVmwarePtrOutput
	ToVirtualNetworksPropertiesVmwarePtrOutputWithContext(context.Context) VirtualNetworksPropertiesVmwarePtrOutput
}

type virtualNetworksPropertiesVmwarePtrType VirtualNetworksPropertiesVmwareArgs

func VirtualNetworksPropertiesVmwarePtr(v *VirtualNetworksPropertiesVmwareArgs) VirtualNetworksPropertiesVmwarePtrInput {
	return (*virtualNetworksPropertiesVmwarePtrType)(v)
}

func (*virtualNetworksPropertiesVmwarePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworksPropertiesVmware)(nil)).Elem()
}

func (i *virtualNetworksPropertiesVmwarePtrType) ToVirtualNetworksPropertiesVmwarePtrOutput() VirtualNetworksPropertiesVmwarePtrOutput {
	return i.ToVirtualNetworksPropertiesVmwarePtrOutputWithContext(context.Background())
}

func (i *virtualNetworksPropertiesVmwarePtrType) ToVirtualNetworksPropertiesVmwarePtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesVmwarePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworksPropertiesVmwarePtrOutput)
}

type VirtualNetworksPropertiesVmwareOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesVmwareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksPropertiesVmware)(nil)).Elem()
}

func (o VirtualNetworksPropertiesVmwareOutput) ToVirtualNetworksPropertiesVmwareOutput() VirtualNetworksPropertiesVmwareOutput {
	return o
}

func (o VirtualNetworksPropertiesVmwareOutput) ToVirtualNetworksPropertiesVmwareOutputWithContext(ctx context.Context) VirtualNetworksPropertiesVmwareOutput {
	return o
}

func (o VirtualNetworksPropertiesVmwareOutput) ToVirtualNetworksPropertiesVmwarePtrOutput() VirtualNetworksPropertiesVmwarePtrOutput {
	return o.ToVirtualNetworksPropertiesVmwarePtrOutputWithContext(context.Background())
}

func (o VirtualNetworksPropertiesVmwareOutput) ToVirtualNetworksPropertiesVmwarePtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesVmwarePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworksPropertiesVmware) *VirtualNetworksPropertiesVmware {
		return &v
	}).(VirtualNetworksPropertiesVmwarePtrOutput)
}

func (o VirtualNetworksPropertiesVmwareOutput) SegmentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksPropertiesVmware) *string { return v.SegmentName }).(pulumi.StringPtrOutput)
}

type VirtualNetworksPropertiesVmwarePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworksPropertiesVmwarePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworksPropertiesVmware)(nil)).Elem()
}

func (o VirtualNetworksPropertiesVmwarePtrOutput) ToVirtualNetworksPropertiesVmwarePtrOutput() VirtualNetworksPropertiesVmwarePtrOutput {
	return o
}

func (o VirtualNetworksPropertiesVmwarePtrOutput) ToVirtualNetworksPropertiesVmwarePtrOutputWithContext(ctx context.Context) VirtualNetworksPropertiesVmwarePtrOutput {
	return o
}

func (o VirtualNetworksPropertiesVmwarePtrOutput) Elem() VirtualNetworksPropertiesVmwareOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesVmware) VirtualNetworksPropertiesVmware {
		if v != nil {
			return *v
		}
		var ret VirtualNetworksPropertiesVmware
		return ret
	}).(VirtualNetworksPropertiesVmwareOutput)
}

func (o VirtualNetworksPropertiesVmwarePtrOutput) SegmentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksPropertiesVmware) *string {
		if v == nil {
			return nil
		}
		return v.SegmentName
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworksResponseExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type VirtualNetworksResponseExtendedLocationOutput struct{ *pulumi.OutputState }

func (VirtualNetworksResponseExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworksResponseExtendedLocation)(nil)).Elem()
}

func (o VirtualNetworksResponseExtendedLocationOutput) ToVirtualNetworksResponseExtendedLocationOutput() VirtualNetworksResponseExtendedLocationOutput {
	return o
}

func (o VirtualNetworksResponseExtendedLocationOutput) ToVirtualNetworksResponseExtendedLocationOutputWithContext(ctx context.Context) VirtualNetworksResponseExtendedLocationOutput {
	return o
}

func (o VirtualNetworksResponseExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksResponseExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksResponseExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworksResponseExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VirtualNetworksResponseExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworksResponseExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworksResponseExtendedLocation)(nil)).Elem()
}

func (o VirtualNetworksResponseExtendedLocationPtrOutput) ToVirtualNetworksResponseExtendedLocationPtrOutput() VirtualNetworksResponseExtendedLocationPtrOutput {
	return o
}

func (o VirtualNetworksResponseExtendedLocationPtrOutput) ToVirtualNetworksResponseExtendedLocationPtrOutputWithContext(ctx context.Context) VirtualNetworksResponseExtendedLocationPtrOutput {
	return o
}

func (o VirtualNetworksResponseExtendedLocationPtrOutput) Elem() VirtualNetworksResponseExtendedLocationOutput {
	return o.ApplyT(func(v *VirtualNetworksResponseExtendedLocation) VirtualNetworksResponseExtendedLocation {
		if v != nil {
			return *v
		}
		var ret VirtualNetworksResponseExtendedLocation
		return ret
	}).(VirtualNetworksResponseExtendedLocationOutput)
}

func (o VirtualNetworksResponseExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksResponseExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworksResponseExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworksResponseExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type WindowsProfile struct {
	AdminPassword  *string `pulumi:"adminPassword"`
	AdminUsername  *string `pulumi:"adminUsername"`
	EnableCsiProxy *bool   `pulumi:"enableCsiProxy"`
	LicenseType    *string `pulumi:"licenseType"`
}





type WindowsProfileInput interface {
	pulumi.Input

	ToWindowsProfileOutput() WindowsProfileOutput
	ToWindowsProfileOutputWithContext(context.Context) WindowsProfileOutput
}

type WindowsProfileArgs struct {
	AdminPassword  pulumi.StringPtrInput `pulumi:"adminPassword"`
	AdminUsername  pulumi.StringPtrInput `pulumi:"adminUsername"`
	EnableCsiProxy pulumi.BoolPtrInput   `pulumi:"enableCsiProxy"`
	LicenseType    pulumi.StringPtrInput `pulumi:"licenseType"`
}

func (WindowsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsProfile)(nil)).Elem()
}

func (i WindowsProfileArgs) ToWindowsProfileOutput() WindowsProfileOutput {
	return i.ToWindowsProfileOutputWithContext(context.Background())
}

func (i WindowsProfileArgs) ToWindowsProfileOutputWithContext(ctx context.Context) WindowsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsProfileOutput)
}

func (i WindowsProfileArgs) ToWindowsProfilePtrOutput() WindowsProfilePtrOutput {
	return i.ToWindowsProfilePtrOutputWithContext(context.Background())
}

func (i WindowsProfileArgs) ToWindowsProfilePtrOutputWithContext(ctx context.Context) WindowsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsProfileOutput).ToWindowsProfilePtrOutputWithContext(ctx)
}









type WindowsProfilePtrInput interface {
	pulumi.Input

	ToWindowsProfilePtrOutput() WindowsProfilePtrOutput
	ToWindowsProfilePtrOutputWithContext(context.Context) WindowsProfilePtrOutput
}

type windowsProfilePtrType WindowsProfileArgs

func WindowsProfilePtr(v *WindowsProfileArgs) WindowsProfilePtrInput {
	return (*windowsProfilePtrType)(v)
}

func (*windowsProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsProfile)(nil)).Elem()
}

func (i *windowsProfilePtrType) ToWindowsProfilePtrOutput() WindowsProfilePtrOutput {
	return i.ToWindowsProfilePtrOutputWithContext(context.Background())
}

func (i *windowsProfilePtrType) ToWindowsProfilePtrOutputWithContext(ctx context.Context) WindowsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsProfilePtrOutput)
}

type WindowsProfileOutput struct{ *pulumi.OutputState }

func (WindowsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsProfile)(nil)).Elem()
}

func (o WindowsProfileOutput) ToWindowsProfileOutput() WindowsProfileOutput {
	return o
}

func (o WindowsProfileOutput) ToWindowsProfileOutputWithContext(ctx context.Context) WindowsProfileOutput {
	return o
}

func (o WindowsProfileOutput) ToWindowsProfilePtrOutput() WindowsProfilePtrOutput {
	return o.ToWindowsProfilePtrOutputWithContext(context.Background())
}

func (o WindowsProfileOutput) ToWindowsProfilePtrOutputWithContext(ctx context.Context) WindowsProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WindowsProfile) *WindowsProfile {
		return &v
	}).(WindowsProfilePtrOutput)
}

func (o WindowsProfileOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsProfile) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o WindowsProfileOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsProfile) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o WindowsProfileOutput) EnableCsiProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WindowsProfile) *bool { return v.EnableCsiProxy }).(pulumi.BoolPtrOutput)
}

func (o WindowsProfileOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsProfile) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

type WindowsProfilePtrOutput struct{ *pulumi.OutputState }

func (WindowsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsProfile)(nil)).Elem()
}

func (o WindowsProfilePtrOutput) ToWindowsProfilePtrOutput() WindowsProfilePtrOutput {
	return o
}

func (o WindowsProfilePtrOutput) ToWindowsProfilePtrOutputWithContext(ctx context.Context) WindowsProfilePtrOutput {
	return o
}

func (o WindowsProfilePtrOutput) Elem() WindowsProfileOutput {
	return o.ApplyT(func(v *WindowsProfile) WindowsProfile {
		if v != nil {
			return *v
		}
		var ret WindowsProfile
		return ret
	}).(WindowsProfileOutput)
}

func (o WindowsProfilePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o WindowsProfilePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o WindowsProfilePtrOutput) EnableCsiProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WindowsProfile) *bool {
		if v == nil {
			return nil
		}
		return v.EnableCsiProxy
	}).(pulumi.BoolPtrOutput)
}

func (o WindowsProfilePtrOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsProfile) *string {
		if v == nil {
			return nil
		}
		return v.LicenseType
	}).(pulumi.StringPtrOutput)
}

type WindowsProfileResponseResponse struct {
	AdminUsername  *string `pulumi:"adminUsername"`
	EnableCsiProxy *bool   `pulumi:"enableCsiProxy"`
	LicenseType    *string `pulumi:"licenseType"`
}

type WindowsProfileResponseResponseOutput struct{ *pulumi.OutputState }

func (WindowsProfileResponseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsProfileResponseResponse)(nil)).Elem()
}

func (o WindowsProfileResponseResponseOutput) ToWindowsProfileResponseResponseOutput() WindowsProfileResponseResponseOutput {
	return o
}

func (o WindowsProfileResponseResponseOutput) ToWindowsProfileResponseResponseOutputWithContext(ctx context.Context) WindowsProfileResponseResponseOutput {
	return o
}

func (o WindowsProfileResponseResponseOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsProfileResponseResponse) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o WindowsProfileResponseResponseOutput) EnableCsiProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WindowsProfileResponseResponse) *bool { return v.EnableCsiProxy }).(pulumi.BoolPtrOutput)
}

func (o WindowsProfileResponseResponseOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsProfileResponseResponse) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

type WindowsProfileResponseResponsePtrOutput struct{ *pulumi.OutputState }

func (WindowsProfileResponseResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsProfileResponseResponse)(nil)).Elem()
}

func (o WindowsProfileResponseResponsePtrOutput) ToWindowsProfileResponseResponsePtrOutput() WindowsProfileResponseResponsePtrOutput {
	return o
}

func (o WindowsProfileResponseResponsePtrOutput) ToWindowsProfileResponseResponsePtrOutputWithContext(ctx context.Context) WindowsProfileResponseResponsePtrOutput {
	return o
}

func (o WindowsProfileResponseResponsePtrOutput) Elem() WindowsProfileResponseResponseOutput {
	return o.ApplyT(func(v *WindowsProfileResponseResponse) WindowsProfileResponseResponse {
		if v != nil {
			return *v
		}
		var ret WindowsProfileResponseResponse
		return ret
	}).(WindowsProfileResponseResponseOutput)
}

func (o WindowsProfileResponseResponsePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsProfileResponseResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o WindowsProfileResponseResponsePtrOutput) EnableCsiProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WindowsProfileResponseResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableCsiProxy
	}).(pulumi.BoolPtrOutput)
}

func (o WindowsProfileResponseResponsePtrOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsProfileResponseResponse) *string {
		if v == nil {
			return nil
		}
		return v.LicenseType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AADProfileOutput{})
	pulumi.RegisterOutputType(AADProfilePtrOutput{})
	pulumi.RegisterOutputType(AADProfileResponseResponseOutput{})
	pulumi.RegisterOutputType(AADProfileResponseResponsePtrOutput{})
	pulumi.RegisterOutputType(AddonProfilesOutput{})
	pulumi.RegisterOutputType(AddonProfilesMapOutput{})
	pulumi.RegisterOutputType(AddonProfilesResponseOutput{})
	pulumi.RegisterOutputType(AddonProfilesResponseMapOutput{})
	pulumi.RegisterOutputType(AddonStatusResponseOutput{})
	pulumi.RegisterOutputType(AddonStatusResponseMapOutput{})
	pulumi.RegisterOutputType(AgentPoolExtendedLocationOutput{})
	pulumi.RegisterOutputType(AgentPoolExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(AgentPoolProvisioningStatusErrorOutput{})
	pulumi.RegisterOutputType(AgentPoolProvisioningStatusErrorPtrOutput{})
	pulumi.RegisterOutputType(AgentPoolProvisioningStatusProvisioningStatusOutput{})
	pulumi.RegisterOutputType(AgentPoolProvisioningStatusProvisioningStatusPtrOutput{})
	pulumi.RegisterOutputType(AgentPoolProvisioningStatusResponseErrorOutput{})
	pulumi.RegisterOutputType(AgentPoolProvisioningStatusResponseErrorPtrOutput{})
	pulumi.RegisterOutputType(AgentPoolProvisioningStatusResponseProvisioningStatusOutput{})
	pulumi.RegisterOutputType(AgentPoolProvisioningStatusResponseProvisioningStatusPtrOutput{})
	pulumi.RegisterOutputType(AgentPoolProvisioningStatusResponseStatusOutput{})
	pulumi.RegisterOutputType(AgentPoolProvisioningStatusResponseStatusPtrOutput{})
	pulumi.RegisterOutputType(AgentPoolProvisioningStatusStatusOutput{})
	pulumi.RegisterOutputType(AgentPoolProvisioningStatusStatusPtrOutput{})
	pulumi.RegisterOutputType(AgentPoolResponseExtendedLocationOutput{})
	pulumi.RegisterOutputType(AgentPoolResponseExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(ArcAgentProfileOutput{})
	pulumi.RegisterOutputType(ArcAgentProfilePtrOutput{})
	pulumi.RegisterOutputType(ArcAgentProfileResponseOutput{})
	pulumi.RegisterOutputType(ArcAgentProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ArcAgentStatusResponseOutput{})
	pulumi.RegisterOutputType(ArcAgentStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(CloudProviderProfileOutput{})
	pulumi.RegisterOutputType(CloudProviderProfilePtrOutput{})
	pulumi.RegisterOutputType(CloudProviderProfileInfraNetworkProfileOutput{})
	pulumi.RegisterOutputType(CloudProviderProfileInfraNetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(CloudProviderProfileInfraStorageProfileOutput{})
	pulumi.RegisterOutputType(CloudProviderProfileInfraStorageProfilePtrOutput{})
	pulumi.RegisterOutputType(CloudProviderProfileResponseOutput{})
	pulumi.RegisterOutputType(CloudProviderProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(CloudProviderProfileResponseInfraNetworkProfileOutput{})
	pulumi.RegisterOutputType(CloudProviderProfileResponseInfraNetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(CloudProviderProfileResponseInfraStorageProfileOutput{})
	pulumi.RegisterOutputType(CloudProviderProfileResponseInfraStorageProfilePtrOutput{})
	pulumi.RegisterOutputType(ControlPlaneEndpointProfileControlPlaneEndpointOutput{})
	pulumi.RegisterOutputType(ControlPlaneEndpointProfileControlPlaneEndpointPtrOutput{})
	pulumi.RegisterOutputType(ControlPlaneEndpointProfileResponseControlPlaneEndpointOutput{})
	pulumi.RegisterOutputType(ControlPlaneEndpointProfileResponseControlPlaneEndpointPtrOutput{})
	pulumi.RegisterOutputType(ControlPlaneProfileOutput{})
	pulumi.RegisterOutputType(ControlPlaneProfilePtrOutput{})
	pulumi.RegisterOutputType(ControlPlaneProfileResponseOutput{})
	pulumi.RegisterOutputType(ControlPlaneProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(HttpProxyConfigOutput{})
	pulumi.RegisterOutputType(HttpProxyConfigPtrOutput{})
	pulumi.RegisterOutputType(HttpProxyConfigResponseResponseOutput{})
	pulumi.RegisterOutputType(HttpProxyConfigResponseResponsePtrOutput{})
	pulumi.RegisterOutputType(LinuxProfilePropertiesOutput{})
	pulumi.RegisterOutputType(LinuxProfilePropertiesPtrOutput{})
	pulumi.RegisterOutputType(LinuxProfilePropertiesPublicKeysOutput{})
	pulumi.RegisterOutputType(LinuxProfilePropertiesPublicKeysArrayOutput{})
	pulumi.RegisterOutputType(LinuxProfilePropertiesResponseOutput{})
	pulumi.RegisterOutputType(LinuxProfilePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LinuxProfilePropertiesResponsePublicKeysOutput{})
	pulumi.RegisterOutputType(LinuxProfilePropertiesResponsePublicKeysArrayOutput{})
	pulumi.RegisterOutputType(LinuxProfilePropertiesResponseSshOutput{})
	pulumi.RegisterOutputType(LinuxProfilePropertiesResponseSshPtrOutput{})
	pulumi.RegisterOutputType(LinuxProfilePropertiesSshOutput{})
	pulumi.RegisterOutputType(LinuxProfilePropertiesSshPtrOutput{})
	pulumi.RegisterOutputType(LoadBalancerProfileOutput{})
	pulumi.RegisterOutputType(LoadBalancerProfilePtrOutput{})
	pulumi.RegisterOutputType(LoadBalancerProfileResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancerProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(NamedAgentPoolProfileOutput{})
	pulumi.RegisterOutputType(NamedAgentPoolProfileArrayOutput{})
	pulumi.RegisterOutputType(NamedAgentPoolProfileResponseOutput{})
	pulumi.RegisterOutputType(NamedAgentPoolProfileResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkProfileOutput{})
	pulumi.RegisterOutputType(NetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ProvisionedClusterIdentityOutput{})
	pulumi.RegisterOutputType(ProvisionedClusterIdentityPtrOutput{})
	pulumi.RegisterOutputType(ProvisionedClusterIdentityResponseOutput{})
	pulumi.RegisterOutputType(ProvisionedClusterIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ProvisionedClustersAllPropertiesOutput{})
	pulumi.RegisterOutputType(ProvisionedClustersAllPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ProvisionedClustersCommonPropertiesFeaturesOutput{})
	pulumi.RegisterOutputType(ProvisionedClustersCommonPropertiesFeaturesPtrOutput{})
	pulumi.RegisterOutputType(ProvisionedClustersCommonPropertiesResponseErrorOutput{})
	pulumi.RegisterOutputType(ProvisionedClustersCommonPropertiesResponseErrorPtrOutput{})
	pulumi.RegisterOutputType(ProvisionedClustersCommonPropertiesResponseFeaturesOutput{})
	pulumi.RegisterOutputType(ProvisionedClustersCommonPropertiesResponseFeaturesPtrOutput{})
	pulumi.RegisterOutputType(ProvisionedClustersCommonPropertiesResponseFeaturesStatusOutput{})
	pulumi.RegisterOutputType(ProvisionedClustersCommonPropertiesResponseFeaturesStatusPtrOutput{})
	pulumi.RegisterOutputType(ProvisionedClustersCommonPropertiesResponseProvisioningStatusOutput{})
	pulumi.RegisterOutputType(ProvisionedClustersCommonPropertiesResponseProvisioningStatusPtrOutput{})
	pulumi.RegisterOutputType(ProvisionedClustersCommonPropertiesResponseStatusOutput{})
	pulumi.RegisterOutputType(ProvisionedClustersExtendedLocationOutput{})
	pulumi.RegisterOutputType(ProvisionedClustersExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(ProvisionedClustersResponsePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ProvisionedClustersResponseResponseExtendedLocationOutput{})
	pulumi.RegisterOutputType(ProvisionedClustersResponseResponseExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(StorageSpacesExtendedLocationOutput{})
	pulumi.RegisterOutputType(StorageSpacesExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesPtrOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesErrorOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesErrorPtrOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesHciStorageProfileOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesHciStorageProfilePtrOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesProvisioningStatusOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesProvisioningStatusPtrOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesResponseErrorOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesResponseErrorPtrOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesResponseHciStorageProfileOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesResponseHciStorageProfilePtrOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesResponseProvisioningStatusOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesResponseProvisioningStatusPtrOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesResponseStatusOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesResponseStatusPtrOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesResponseVmwareStorageProfileOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesResponseVmwareStorageProfilePtrOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesStatusOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesStatusPtrOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesVmwareStorageProfileOutput{})
	pulumi.RegisterOutputType(StorageSpacesPropertiesVmwareStorageProfilePtrOutput{})
	pulumi.RegisterOutputType(StorageSpacesResponseExtendedLocationOutput{})
	pulumi.RegisterOutputType(StorageSpacesResponseExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworksExtendedLocationOutput{})
	pulumi.RegisterOutputType(VirtualNetworksExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesHciOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesHciPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesInfraVnetProfileOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesInfraVnetProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesKubevirtOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesKubevirtPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesResponseErrorOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesResponseErrorPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesResponseHciOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesResponseHciPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesResponseInfraVnetProfileOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesResponseInfraVnetProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesResponseKubevirtOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesResponseKubevirtPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesResponseProvisioningStatusOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesResponseProvisioningStatusPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesResponseStatusOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesResponseVipPoolOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesResponseVipPoolArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesResponseVmipPoolOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesResponseVmipPoolArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesResponseVmwareOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesResponseVmwarePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesVipPoolOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesVipPoolArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesVmipPoolOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesVmipPoolArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesVmwareOutput{})
	pulumi.RegisterOutputType(VirtualNetworksPropertiesVmwarePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworksResponseExtendedLocationOutput{})
	pulumi.RegisterOutputType(VirtualNetworksResponseExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(WindowsProfileOutput{})
	pulumi.RegisterOutputType(WindowsProfilePtrOutput{})
	pulumi.RegisterOutputType(WindowsProfileResponseResponseOutput{})
	pulumi.RegisterOutputType(WindowsProfileResponseResponsePtrOutput{})
}

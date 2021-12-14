


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AgentPoolUpgradeSettings struct {
	MaxSurge *string `pulumi:"maxSurge"`
}





type AgentPoolUpgradeSettingsInput interface {
	pulumi.Input

	ToAgentPoolUpgradeSettingsOutput() AgentPoolUpgradeSettingsOutput
	ToAgentPoolUpgradeSettingsOutputWithContext(context.Context) AgentPoolUpgradeSettingsOutput
}

type AgentPoolUpgradeSettingsArgs struct {
	MaxSurge pulumi.StringPtrInput `pulumi:"maxSurge"`
}

func (AgentPoolUpgradeSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPoolUpgradeSettings)(nil)).Elem()
}

func (i AgentPoolUpgradeSettingsArgs) ToAgentPoolUpgradeSettingsOutput() AgentPoolUpgradeSettingsOutput {
	return i.ToAgentPoolUpgradeSettingsOutputWithContext(context.Background())
}

func (i AgentPoolUpgradeSettingsArgs) ToAgentPoolUpgradeSettingsOutputWithContext(ctx context.Context) AgentPoolUpgradeSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolUpgradeSettingsOutput)
}

func (i AgentPoolUpgradeSettingsArgs) ToAgentPoolUpgradeSettingsPtrOutput() AgentPoolUpgradeSettingsPtrOutput {
	return i.ToAgentPoolUpgradeSettingsPtrOutputWithContext(context.Background())
}

func (i AgentPoolUpgradeSettingsArgs) ToAgentPoolUpgradeSettingsPtrOutputWithContext(ctx context.Context) AgentPoolUpgradeSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolUpgradeSettingsOutput).ToAgentPoolUpgradeSettingsPtrOutputWithContext(ctx)
}









type AgentPoolUpgradeSettingsPtrInput interface {
	pulumi.Input

	ToAgentPoolUpgradeSettingsPtrOutput() AgentPoolUpgradeSettingsPtrOutput
	ToAgentPoolUpgradeSettingsPtrOutputWithContext(context.Context) AgentPoolUpgradeSettingsPtrOutput
}

type agentPoolUpgradeSettingsPtrType AgentPoolUpgradeSettingsArgs

func AgentPoolUpgradeSettingsPtr(v *AgentPoolUpgradeSettingsArgs) AgentPoolUpgradeSettingsPtrInput {
	return (*agentPoolUpgradeSettingsPtrType)(v)
}

func (*agentPoolUpgradeSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPoolUpgradeSettings)(nil)).Elem()
}

func (i *agentPoolUpgradeSettingsPtrType) ToAgentPoolUpgradeSettingsPtrOutput() AgentPoolUpgradeSettingsPtrOutput {
	return i.ToAgentPoolUpgradeSettingsPtrOutputWithContext(context.Background())
}

func (i *agentPoolUpgradeSettingsPtrType) ToAgentPoolUpgradeSettingsPtrOutputWithContext(ctx context.Context) AgentPoolUpgradeSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolUpgradeSettingsPtrOutput)
}

type AgentPoolUpgradeSettingsOutput struct{ *pulumi.OutputState }

func (AgentPoolUpgradeSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPoolUpgradeSettings)(nil)).Elem()
}

func (o AgentPoolUpgradeSettingsOutput) ToAgentPoolUpgradeSettingsOutput() AgentPoolUpgradeSettingsOutput {
	return o
}

func (o AgentPoolUpgradeSettingsOutput) ToAgentPoolUpgradeSettingsOutputWithContext(ctx context.Context) AgentPoolUpgradeSettingsOutput {
	return o
}

func (o AgentPoolUpgradeSettingsOutput) ToAgentPoolUpgradeSettingsPtrOutput() AgentPoolUpgradeSettingsPtrOutput {
	return o.ToAgentPoolUpgradeSettingsPtrOutputWithContext(context.Background())
}

func (o AgentPoolUpgradeSettingsOutput) ToAgentPoolUpgradeSettingsPtrOutputWithContext(ctx context.Context) AgentPoolUpgradeSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AgentPoolUpgradeSettings) *AgentPoolUpgradeSettings {
		return &v
	}).(AgentPoolUpgradeSettingsPtrOutput)
}

func (o AgentPoolUpgradeSettingsOutput) MaxSurge() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPoolUpgradeSettings) *string { return v.MaxSurge }).(pulumi.StringPtrOutput)
}

type AgentPoolUpgradeSettingsPtrOutput struct{ *pulumi.OutputState }

func (AgentPoolUpgradeSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPoolUpgradeSettings)(nil)).Elem()
}

func (o AgentPoolUpgradeSettingsPtrOutput) ToAgentPoolUpgradeSettingsPtrOutput() AgentPoolUpgradeSettingsPtrOutput {
	return o
}

func (o AgentPoolUpgradeSettingsPtrOutput) ToAgentPoolUpgradeSettingsPtrOutputWithContext(ctx context.Context) AgentPoolUpgradeSettingsPtrOutput {
	return o
}

func (o AgentPoolUpgradeSettingsPtrOutput) Elem() AgentPoolUpgradeSettingsOutput {
	return o.ApplyT(func(v *AgentPoolUpgradeSettings) AgentPoolUpgradeSettings {
		if v != nil {
			return *v
		}
		var ret AgentPoolUpgradeSettings
		return ret
	}).(AgentPoolUpgradeSettingsOutput)
}

func (o AgentPoolUpgradeSettingsPtrOutput) MaxSurge() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPoolUpgradeSettings) *string {
		if v == nil {
			return nil
		}
		return v.MaxSurge
	}).(pulumi.StringPtrOutput)
}

type AgentPoolUpgradeSettingsResponse struct {
	MaxSurge *string `pulumi:"maxSurge"`
}





type AgentPoolUpgradeSettingsResponseInput interface {
	pulumi.Input

	ToAgentPoolUpgradeSettingsResponseOutput() AgentPoolUpgradeSettingsResponseOutput
	ToAgentPoolUpgradeSettingsResponseOutputWithContext(context.Context) AgentPoolUpgradeSettingsResponseOutput
}

type AgentPoolUpgradeSettingsResponseArgs struct {
	MaxSurge pulumi.StringPtrInput `pulumi:"maxSurge"`
}

func (AgentPoolUpgradeSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPoolUpgradeSettingsResponse)(nil)).Elem()
}

func (i AgentPoolUpgradeSettingsResponseArgs) ToAgentPoolUpgradeSettingsResponseOutput() AgentPoolUpgradeSettingsResponseOutput {
	return i.ToAgentPoolUpgradeSettingsResponseOutputWithContext(context.Background())
}

func (i AgentPoolUpgradeSettingsResponseArgs) ToAgentPoolUpgradeSettingsResponseOutputWithContext(ctx context.Context) AgentPoolUpgradeSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolUpgradeSettingsResponseOutput)
}

func (i AgentPoolUpgradeSettingsResponseArgs) ToAgentPoolUpgradeSettingsResponsePtrOutput() AgentPoolUpgradeSettingsResponsePtrOutput {
	return i.ToAgentPoolUpgradeSettingsResponsePtrOutputWithContext(context.Background())
}

func (i AgentPoolUpgradeSettingsResponseArgs) ToAgentPoolUpgradeSettingsResponsePtrOutputWithContext(ctx context.Context) AgentPoolUpgradeSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolUpgradeSettingsResponseOutput).ToAgentPoolUpgradeSettingsResponsePtrOutputWithContext(ctx)
}









type AgentPoolUpgradeSettingsResponsePtrInput interface {
	pulumi.Input

	ToAgentPoolUpgradeSettingsResponsePtrOutput() AgentPoolUpgradeSettingsResponsePtrOutput
	ToAgentPoolUpgradeSettingsResponsePtrOutputWithContext(context.Context) AgentPoolUpgradeSettingsResponsePtrOutput
}

type agentPoolUpgradeSettingsResponsePtrType AgentPoolUpgradeSettingsResponseArgs

func AgentPoolUpgradeSettingsResponsePtr(v *AgentPoolUpgradeSettingsResponseArgs) AgentPoolUpgradeSettingsResponsePtrInput {
	return (*agentPoolUpgradeSettingsResponsePtrType)(v)
}

func (*agentPoolUpgradeSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPoolUpgradeSettingsResponse)(nil)).Elem()
}

func (i *agentPoolUpgradeSettingsResponsePtrType) ToAgentPoolUpgradeSettingsResponsePtrOutput() AgentPoolUpgradeSettingsResponsePtrOutput {
	return i.ToAgentPoolUpgradeSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *agentPoolUpgradeSettingsResponsePtrType) ToAgentPoolUpgradeSettingsResponsePtrOutputWithContext(ctx context.Context) AgentPoolUpgradeSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolUpgradeSettingsResponsePtrOutput)
}

type AgentPoolUpgradeSettingsResponseOutput struct{ *pulumi.OutputState }

func (AgentPoolUpgradeSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPoolUpgradeSettingsResponse)(nil)).Elem()
}

func (o AgentPoolUpgradeSettingsResponseOutput) ToAgentPoolUpgradeSettingsResponseOutput() AgentPoolUpgradeSettingsResponseOutput {
	return o
}

func (o AgentPoolUpgradeSettingsResponseOutput) ToAgentPoolUpgradeSettingsResponseOutputWithContext(ctx context.Context) AgentPoolUpgradeSettingsResponseOutput {
	return o
}

func (o AgentPoolUpgradeSettingsResponseOutput) ToAgentPoolUpgradeSettingsResponsePtrOutput() AgentPoolUpgradeSettingsResponsePtrOutput {
	return o.ToAgentPoolUpgradeSettingsResponsePtrOutputWithContext(context.Background())
}

func (o AgentPoolUpgradeSettingsResponseOutput) ToAgentPoolUpgradeSettingsResponsePtrOutputWithContext(ctx context.Context) AgentPoolUpgradeSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AgentPoolUpgradeSettingsResponse) *AgentPoolUpgradeSettingsResponse {
		return &v
	}).(AgentPoolUpgradeSettingsResponsePtrOutput)
}

func (o AgentPoolUpgradeSettingsResponseOutput) MaxSurge() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPoolUpgradeSettingsResponse) *string { return v.MaxSurge }).(pulumi.StringPtrOutput)
}

type AgentPoolUpgradeSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AgentPoolUpgradeSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPoolUpgradeSettingsResponse)(nil)).Elem()
}

func (o AgentPoolUpgradeSettingsResponsePtrOutput) ToAgentPoolUpgradeSettingsResponsePtrOutput() AgentPoolUpgradeSettingsResponsePtrOutput {
	return o
}

func (o AgentPoolUpgradeSettingsResponsePtrOutput) ToAgentPoolUpgradeSettingsResponsePtrOutputWithContext(ctx context.Context) AgentPoolUpgradeSettingsResponsePtrOutput {
	return o
}

func (o AgentPoolUpgradeSettingsResponsePtrOutput) Elem() AgentPoolUpgradeSettingsResponseOutput {
	return o.ApplyT(func(v *AgentPoolUpgradeSettingsResponse) AgentPoolUpgradeSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AgentPoolUpgradeSettingsResponse
		return ret
	}).(AgentPoolUpgradeSettingsResponseOutput)
}

func (o AgentPoolUpgradeSettingsResponsePtrOutput) MaxSurge() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPoolUpgradeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.MaxSurge
	}).(pulumi.StringPtrOutput)
}

type CloudErrorBodyResponse struct {
	Code    *string                  `pulumi:"code"`
	Details []CloudErrorBodyResponse `pulumi:"details"`
	Message *string                  `pulumi:"message"`
	Target  *string                  `pulumi:"target"`
}





type CloudErrorBodyResponseInput interface {
	pulumi.Input

	ToCloudErrorBodyResponseOutput() CloudErrorBodyResponseOutput
	ToCloudErrorBodyResponseOutputWithContext(context.Context) CloudErrorBodyResponseOutput
}

type CloudErrorBodyResponseArgs struct {
	Code    pulumi.StringPtrInput            `pulumi:"code"`
	Details CloudErrorBodyResponseArrayInput `pulumi:"details"`
	Message pulumi.StringPtrInput            `pulumi:"message"`
	Target  pulumi.StringPtrInput            `pulumi:"target"`
}

func (CloudErrorBodyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudErrorBodyResponse)(nil)).Elem()
}

func (i CloudErrorBodyResponseArgs) ToCloudErrorBodyResponseOutput() CloudErrorBodyResponseOutput {
	return i.ToCloudErrorBodyResponseOutputWithContext(context.Background())
}

func (i CloudErrorBodyResponseArgs) ToCloudErrorBodyResponseOutputWithContext(ctx context.Context) CloudErrorBodyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudErrorBodyResponseOutput)
}

func (i CloudErrorBodyResponseArgs) ToCloudErrorBodyResponsePtrOutput() CloudErrorBodyResponsePtrOutput {
	return i.ToCloudErrorBodyResponsePtrOutputWithContext(context.Background())
}

func (i CloudErrorBodyResponseArgs) ToCloudErrorBodyResponsePtrOutputWithContext(ctx context.Context) CloudErrorBodyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudErrorBodyResponseOutput).ToCloudErrorBodyResponsePtrOutputWithContext(ctx)
}









type CloudErrorBodyResponsePtrInput interface {
	pulumi.Input

	ToCloudErrorBodyResponsePtrOutput() CloudErrorBodyResponsePtrOutput
	ToCloudErrorBodyResponsePtrOutputWithContext(context.Context) CloudErrorBodyResponsePtrOutput
}

type cloudErrorBodyResponsePtrType CloudErrorBodyResponseArgs

func CloudErrorBodyResponsePtr(v *CloudErrorBodyResponseArgs) CloudErrorBodyResponsePtrInput {
	return (*cloudErrorBodyResponsePtrType)(v)
}

func (*cloudErrorBodyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudErrorBodyResponse)(nil)).Elem()
}

func (i *cloudErrorBodyResponsePtrType) ToCloudErrorBodyResponsePtrOutput() CloudErrorBodyResponsePtrOutput {
	return i.ToCloudErrorBodyResponsePtrOutputWithContext(context.Background())
}

func (i *cloudErrorBodyResponsePtrType) ToCloudErrorBodyResponsePtrOutputWithContext(ctx context.Context) CloudErrorBodyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudErrorBodyResponsePtrOutput)
}





type CloudErrorBodyResponseArrayInput interface {
	pulumi.Input

	ToCloudErrorBodyResponseArrayOutput() CloudErrorBodyResponseArrayOutput
	ToCloudErrorBodyResponseArrayOutputWithContext(context.Context) CloudErrorBodyResponseArrayOutput
}

type CloudErrorBodyResponseArray []CloudErrorBodyResponseInput

func (CloudErrorBodyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudErrorBodyResponse)(nil)).Elem()
}

func (i CloudErrorBodyResponseArray) ToCloudErrorBodyResponseArrayOutput() CloudErrorBodyResponseArrayOutput {
	return i.ToCloudErrorBodyResponseArrayOutputWithContext(context.Background())
}

func (i CloudErrorBodyResponseArray) ToCloudErrorBodyResponseArrayOutputWithContext(ctx context.Context) CloudErrorBodyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudErrorBodyResponseArrayOutput)
}

type CloudErrorBodyResponseOutput struct{ *pulumi.OutputState }

func (CloudErrorBodyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudErrorBodyResponse)(nil)).Elem()
}

func (o CloudErrorBodyResponseOutput) ToCloudErrorBodyResponseOutput() CloudErrorBodyResponseOutput {
	return o
}

func (o CloudErrorBodyResponseOutput) ToCloudErrorBodyResponseOutputWithContext(ctx context.Context) CloudErrorBodyResponseOutput {
	return o
}

func (o CloudErrorBodyResponseOutput) ToCloudErrorBodyResponsePtrOutput() CloudErrorBodyResponsePtrOutput {
	return o.ToCloudErrorBodyResponsePtrOutputWithContext(context.Background())
}

func (o CloudErrorBodyResponseOutput) ToCloudErrorBodyResponsePtrOutputWithContext(ctx context.Context) CloudErrorBodyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudErrorBodyResponse) *CloudErrorBodyResponse {
		return &v
	}).(CloudErrorBodyResponsePtrOutput)
}

func (o CloudErrorBodyResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudErrorBodyResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o CloudErrorBodyResponseOutput) Details() CloudErrorBodyResponseArrayOutput {
	return o.ApplyT(func(v CloudErrorBodyResponse) []CloudErrorBodyResponse { return v.Details }).(CloudErrorBodyResponseArrayOutput)
}

func (o CloudErrorBodyResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudErrorBodyResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o CloudErrorBodyResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudErrorBodyResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type CloudErrorBodyResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudErrorBodyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudErrorBodyResponse)(nil)).Elem()
}

func (o CloudErrorBodyResponsePtrOutput) ToCloudErrorBodyResponsePtrOutput() CloudErrorBodyResponsePtrOutput {
	return o
}

func (o CloudErrorBodyResponsePtrOutput) ToCloudErrorBodyResponsePtrOutputWithContext(ctx context.Context) CloudErrorBodyResponsePtrOutput {
	return o
}

func (o CloudErrorBodyResponsePtrOutput) Elem() CloudErrorBodyResponseOutput {
	return o.ApplyT(func(v *CloudErrorBodyResponse) CloudErrorBodyResponse {
		if v != nil {
			return *v
		}
		var ret CloudErrorBodyResponse
		return ret
	}).(CloudErrorBodyResponseOutput)
}

func (o CloudErrorBodyResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudErrorBodyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o CloudErrorBodyResponsePtrOutput) Details() CloudErrorBodyResponseArrayOutput {
	return o.ApplyT(func(v *CloudErrorBodyResponse) []CloudErrorBodyResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(CloudErrorBodyResponseArrayOutput)
}

func (o CloudErrorBodyResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudErrorBodyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o CloudErrorBodyResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudErrorBodyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

type CloudErrorBodyResponseArrayOutput struct{ *pulumi.OutputState }

func (CloudErrorBodyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudErrorBodyResponse)(nil)).Elem()
}

func (o CloudErrorBodyResponseArrayOutput) ToCloudErrorBodyResponseArrayOutput() CloudErrorBodyResponseArrayOutput {
	return o
}

func (o CloudErrorBodyResponseArrayOutput) ToCloudErrorBodyResponseArrayOutputWithContext(ctx context.Context) CloudErrorBodyResponseArrayOutput {
	return o
}

func (o CloudErrorBodyResponseArrayOutput) Index(i pulumi.IntInput) CloudErrorBodyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CloudErrorBodyResponse {
		return vs[0].([]CloudErrorBodyResponse)[vs[1].(int)]
	}).(CloudErrorBodyResponseOutput)
}

type CloudErrorResponse struct {
	Error *CloudErrorBodyResponse `pulumi:"error"`
}





type CloudErrorResponseInput interface {
	pulumi.Input

	ToCloudErrorResponseOutput() CloudErrorResponseOutput
	ToCloudErrorResponseOutputWithContext(context.Context) CloudErrorResponseOutput
}

type CloudErrorResponseArgs struct {
	Error CloudErrorBodyResponsePtrInput `pulumi:"error"`
}

func (CloudErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudErrorResponse)(nil)).Elem()
}

func (i CloudErrorResponseArgs) ToCloudErrorResponseOutput() CloudErrorResponseOutput {
	return i.ToCloudErrorResponseOutputWithContext(context.Background())
}

func (i CloudErrorResponseArgs) ToCloudErrorResponseOutputWithContext(ctx context.Context) CloudErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudErrorResponseOutput)
}

func (i CloudErrorResponseArgs) ToCloudErrorResponsePtrOutput() CloudErrorResponsePtrOutput {
	return i.ToCloudErrorResponsePtrOutputWithContext(context.Background())
}

func (i CloudErrorResponseArgs) ToCloudErrorResponsePtrOutputWithContext(ctx context.Context) CloudErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudErrorResponseOutput).ToCloudErrorResponsePtrOutputWithContext(ctx)
}









type CloudErrorResponsePtrInput interface {
	pulumi.Input

	ToCloudErrorResponsePtrOutput() CloudErrorResponsePtrOutput
	ToCloudErrorResponsePtrOutputWithContext(context.Context) CloudErrorResponsePtrOutput
}

type cloudErrorResponsePtrType CloudErrorResponseArgs

func CloudErrorResponsePtr(v *CloudErrorResponseArgs) CloudErrorResponsePtrInput {
	return (*cloudErrorResponsePtrType)(v)
}

func (*cloudErrorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudErrorResponse)(nil)).Elem()
}

func (i *cloudErrorResponsePtrType) ToCloudErrorResponsePtrOutput() CloudErrorResponsePtrOutput {
	return i.ToCloudErrorResponsePtrOutputWithContext(context.Background())
}

func (i *cloudErrorResponsePtrType) ToCloudErrorResponsePtrOutputWithContext(ctx context.Context) CloudErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudErrorResponsePtrOutput)
}

type CloudErrorResponseOutput struct{ *pulumi.OutputState }

func (CloudErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudErrorResponse)(nil)).Elem()
}

func (o CloudErrorResponseOutput) ToCloudErrorResponseOutput() CloudErrorResponseOutput {
	return o
}

func (o CloudErrorResponseOutput) ToCloudErrorResponseOutputWithContext(ctx context.Context) CloudErrorResponseOutput {
	return o
}

func (o CloudErrorResponseOutput) ToCloudErrorResponsePtrOutput() CloudErrorResponsePtrOutput {
	return o.ToCloudErrorResponsePtrOutputWithContext(context.Background())
}

func (o CloudErrorResponseOutput) ToCloudErrorResponsePtrOutputWithContext(ctx context.Context) CloudErrorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudErrorResponse) *CloudErrorResponse {
		return &v
	}).(CloudErrorResponsePtrOutput)
}

func (o CloudErrorResponseOutput) Error() CloudErrorBodyResponsePtrOutput {
	return o.ApplyT(func(v CloudErrorResponse) *CloudErrorBodyResponse { return v.Error }).(CloudErrorBodyResponsePtrOutput)
}

type CloudErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudErrorResponse)(nil)).Elem()
}

func (o CloudErrorResponsePtrOutput) ToCloudErrorResponsePtrOutput() CloudErrorResponsePtrOutput {
	return o
}

func (o CloudErrorResponsePtrOutput) ToCloudErrorResponsePtrOutputWithContext(ctx context.Context) CloudErrorResponsePtrOutput {
	return o
}

func (o CloudErrorResponsePtrOutput) Elem() CloudErrorResponseOutput {
	return o.ApplyT(func(v *CloudErrorResponse) CloudErrorResponse {
		if v != nil {
			return *v
		}
		var ret CloudErrorResponse
		return ret
	}).(CloudErrorResponseOutput)
}

func (o CloudErrorResponsePtrOutput) Error() CloudErrorBodyResponsePtrOutput {
	return o.ApplyT(func(v *CloudErrorResponse) *CloudErrorBodyResponse {
		if v == nil {
			return nil
		}
		return v.Error
	}).(CloudErrorBodyResponsePtrOutput)
}

type ContainerServiceLinuxProfile struct {
	AdminUsername string                           `pulumi:"adminUsername"`
	Ssh           ContainerServiceSshConfiguration `pulumi:"ssh"`
}





type ContainerServiceLinuxProfileInput interface {
	pulumi.Input

	ToContainerServiceLinuxProfileOutput() ContainerServiceLinuxProfileOutput
	ToContainerServiceLinuxProfileOutputWithContext(context.Context) ContainerServiceLinuxProfileOutput
}

type ContainerServiceLinuxProfileArgs struct {
	AdminUsername pulumi.StringInput                    `pulumi:"adminUsername"`
	Ssh           ContainerServiceSshConfigurationInput `pulumi:"ssh"`
}

func (ContainerServiceLinuxProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceLinuxProfile)(nil)).Elem()
}

func (i ContainerServiceLinuxProfileArgs) ToContainerServiceLinuxProfileOutput() ContainerServiceLinuxProfileOutput {
	return i.ToContainerServiceLinuxProfileOutputWithContext(context.Background())
}

func (i ContainerServiceLinuxProfileArgs) ToContainerServiceLinuxProfileOutputWithContext(ctx context.Context) ContainerServiceLinuxProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceLinuxProfileOutput)
}

func (i ContainerServiceLinuxProfileArgs) ToContainerServiceLinuxProfilePtrOutput() ContainerServiceLinuxProfilePtrOutput {
	return i.ToContainerServiceLinuxProfilePtrOutputWithContext(context.Background())
}

func (i ContainerServiceLinuxProfileArgs) ToContainerServiceLinuxProfilePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceLinuxProfileOutput).ToContainerServiceLinuxProfilePtrOutputWithContext(ctx)
}









type ContainerServiceLinuxProfilePtrInput interface {
	pulumi.Input

	ToContainerServiceLinuxProfilePtrOutput() ContainerServiceLinuxProfilePtrOutput
	ToContainerServiceLinuxProfilePtrOutputWithContext(context.Context) ContainerServiceLinuxProfilePtrOutput
}

type containerServiceLinuxProfilePtrType ContainerServiceLinuxProfileArgs

func ContainerServiceLinuxProfilePtr(v *ContainerServiceLinuxProfileArgs) ContainerServiceLinuxProfilePtrInput {
	return (*containerServiceLinuxProfilePtrType)(v)
}

func (*containerServiceLinuxProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceLinuxProfile)(nil)).Elem()
}

func (i *containerServiceLinuxProfilePtrType) ToContainerServiceLinuxProfilePtrOutput() ContainerServiceLinuxProfilePtrOutput {
	return i.ToContainerServiceLinuxProfilePtrOutputWithContext(context.Background())
}

func (i *containerServiceLinuxProfilePtrType) ToContainerServiceLinuxProfilePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceLinuxProfilePtrOutput)
}

type ContainerServiceLinuxProfileOutput struct{ *pulumi.OutputState }

func (ContainerServiceLinuxProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceLinuxProfile)(nil)).Elem()
}

func (o ContainerServiceLinuxProfileOutput) ToContainerServiceLinuxProfileOutput() ContainerServiceLinuxProfileOutput {
	return o
}

func (o ContainerServiceLinuxProfileOutput) ToContainerServiceLinuxProfileOutputWithContext(ctx context.Context) ContainerServiceLinuxProfileOutput {
	return o
}

func (o ContainerServiceLinuxProfileOutput) ToContainerServiceLinuxProfilePtrOutput() ContainerServiceLinuxProfilePtrOutput {
	return o.ToContainerServiceLinuxProfilePtrOutputWithContext(context.Background())
}

func (o ContainerServiceLinuxProfileOutput) ToContainerServiceLinuxProfilePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceLinuxProfile) *ContainerServiceLinuxProfile {
		return &v
	}).(ContainerServiceLinuxProfilePtrOutput)
}

func (o ContainerServiceLinuxProfileOutput) AdminUsername() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceLinuxProfile) string { return v.AdminUsername }).(pulumi.StringOutput)
}

func (o ContainerServiceLinuxProfileOutput) Ssh() ContainerServiceSshConfigurationOutput {
	return o.ApplyT(func(v ContainerServiceLinuxProfile) ContainerServiceSshConfiguration { return v.Ssh }).(ContainerServiceSshConfigurationOutput)
}

type ContainerServiceLinuxProfilePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceLinuxProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceLinuxProfile)(nil)).Elem()
}

func (o ContainerServiceLinuxProfilePtrOutput) ToContainerServiceLinuxProfilePtrOutput() ContainerServiceLinuxProfilePtrOutput {
	return o
}

func (o ContainerServiceLinuxProfilePtrOutput) ToContainerServiceLinuxProfilePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfilePtrOutput {
	return o
}

func (o ContainerServiceLinuxProfilePtrOutput) Elem() ContainerServiceLinuxProfileOutput {
	return o.ApplyT(func(v *ContainerServiceLinuxProfile) ContainerServiceLinuxProfile {
		if v != nil {
			return *v
		}
		var ret ContainerServiceLinuxProfile
		return ret
	}).(ContainerServiceLinuxProfileOutput)
}

func (o ContainerServiceLinuxProfilePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceLinuxProfile) *string {
		if v == nil {
			return nil
		}
		return &v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceLinuxProfilePtrOutput) Ssh() ContainerServiceSshConfigurationPtrOutput {
	return o.ApplyT(func(v *ContainerServiceLinuxProfile) *ContainerServiceSshConfiguration {
		if v == nil {
			return nil
		}
		return &v.Ssh
	}).(ContainerServiceSshConfigurationPtrOutput)
}

type ContainerServiceLinuxProfileResponse struct {
	AdminUsername string                                   `pulumi:"adminUsername"`
	Ssh           ContainerServiceSshConfigurationResponse `pulumi:"ssh"`
}





type ContainerServiceLinuxProfileResponseInput interface {
	pulumi.Input

	ToContainerServiceLinuxProfileResponseOutput() ContainerServiceLinuxProfileResponseOutput
	ToContainerServiceLinuxProfileResponseOutputWithContext(context.Context) ContainerServiceLinuxProfileResponseOutput
}

type ContainerServiceLinuxProfileResponseArgs struct {
	AdminUsername pulumi.StringInput                            `pulumi:"adminUsername"`
	Ssh           ContainerServiceSshConfigurationResponseInput `pulumi:"ssh"`
}

func (ContainerServiceLinuxProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceLinuxProfileResponse)(nil)).Elem()
}

func (i ContainerServiceLinuxProfileResponseArgs) ToContainerServiceLinuxProfileResponseOutput() ContainerServiceLinuxProfileResponseOutput {
	return i.ToContainerServiceLinuxProfileResponseOutputWithContext(context.Background())
}

func (i ContainerServiceLinuxProfileResponseArgs) ToContainerServiceLinuxProfileResponseOutputWithContext(ctx context.Context) ContainerServiceLinuxProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceLinuxProfileResponseOutput)
}

func (i ContainerServiceLinuxProfileResponseArgs) ToContainerServiceLinuxProfileResponsePtrOutput() ContainerServiceLinuxProfileResponsePtrOutput {
	return i.ToContainerServiceLinuxProfileResponsePtrOutputWithContext(context.Background())
}

func (i ContainerServiceLinuxProfileResponseArgs) ToContainerServiceLinuxProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceLinuxProfileResponseOutput).ToContainerServiceLinuxProfileResponsePtrOutputWithContext(ctx)
}









type ContainerServiceLinuxProfileResponsePtrInput interface {
	pulumi.Input

	ToContainerServiceLinuxProfileResponsePtrOutput() ContainerServiceLinuxProfileResponsePtrOutput
	ToContainerServiceLinuxProfileResponsePtrOutputWithContext(context.Context) ContainerServiceLinuxProfileResponsePtrOutput
}

type containerServiceLinuxProfileResponsePtrType ContainerServiceLinuxProfileResponseArgs

func ContainerServiceLinuxProfileResponsePtr(v *ContainerServiceLinuxProfileResponseArgs) ContainerServiceLinuxProfileResponsePtrInput {
	return (*containerServiceLinuxProfileResponsePtrType)(v)
}

func (*containerServiceLinuxProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceLinuxProfileResponse)(nil)).Elem()
}

func (i *containerServiceLinuxProfileResponsePtrType) ToContainerServiceLinuxProfileResponsePtrOutput() ContainerServiceLinuxProfileResponsePtrOutput {
	return i.ToContainerServiceLinuxProfileResponsePtrOutputWithContext(context.Background())
}

func (i *containerServiceLinuxProfileResponsePtrType) ToContainerServiceLinuxProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceLinuxProfileResponsePtrOutput)
}

type ContainerServiceLinuxProfileResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceLinuxProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceLinuxProfileResponse)(nil)).Elem()
}

func (o ContainerServiceLinuxProfileResponseOutput) ToContainerServiceLinuxProfileResponseOutput() ContainerServiceLinuxProfileResponseOutput {
	return o
}

func (o ContainerServiceLinuxProfileResponseOutput) ToContainerServiceLinuxProfileResponseOutputWithContext(ctx context.Context) ContainerServiceLinuxProfileResponseOutput {
	return o
}

func (o ContainerServiceLinuxProfileResponseOutput) ToContainerServiceLinuxProfileResponsePtrOutput() ContainerServiceLinuxProfileResponsePtrOutput {
	return o.ToContainerServiceLinuxProfileResponsePtrOutputWithContext(context.Background())
}

func (o ContainerServiceLinuxProfileResponseOutput) ToContainerServiceLinuxProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceLinuxProfileResponse) *ContainerServiceLinuxProfileResponse {
		return &v
	}).(ContainerServiceLinuxProfileResponsePtrOutput)
}

func (o ContainerServiceLinuxProfileResponseOutput) AdminUsername() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceLinuxProfileResponse) string { return v.AdminUsername }).(pulumi.StringOutput)
}

func (o ContainerServiceLinuxProfileResponseOutput) Ssh() ContainerServiceSshConfigurationResponseOutput {
	return o.ApplyT(func(v ContainerServiceLinuxProfileResponse) ContainerServiceSshConfigurationResponse { return v.Ssh }).(ContainerServiceSshConfigurationResponseOutput)
}

type ContainerServiceLinuxProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceLinuxProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceLinuxProfileResponse)(nil)).Elem()
}

func (o ContainerServiceLinuxProfileResponsePtrOutput) ToContainerServiceLinuxProfileResponsePtrOutput() ContainerServiceLinuxProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceLinuxProfileResponsePtrOutput) ToContainerServiceLinuxProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceLinuxProfileResponsePtrOutput) Elem() ContainerServiceLinuxProfileResponseOutput {
	return o.ApplyT(func(v *ContainerServiceLinuxProfileResponse) ContainerServiceLinuxProfileResponse {
		if v != nil {
			return *v
		}
		var ret ContainerServiceLinuxProfileResponse
		return ret
	}).(ContainerServiceLinuxProfileResponseOutput)
}

func (o ContainerServiceLinuxProfileResponsePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceLinuxProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceLinuxProfileResponsePtrOutput) Ssh() ContainerServiceSshConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ContainerServiceLinuxProfileResponse) *ContainerServiceSshConfigurationResponse {
		if v == nil {
			return nil
		}
		return &v.Ssh
	}).(ContainerServiceSshConfigurationResponsePtrOutput)
}

type ContainerServiceNetworkProfile struct {
	DnsServiceIP        *string                            `pulumi:"dnsServiceIP"`
	DockerBridgeCidr    *string                            `pulumi:"dockerBridgeCidr"`
	LoadBalancerProfile *ManagedClusterLoadBalancerProfile `pulumi:"loadBalancerProfile"`
	LoadBalancerSku     *string                            `pulumi:"loadBalancerSku"`
	NetworkMode         *string                            `pulumi:"networkMode"`
	NetworkPlugin       *string                            `pulumi:"networkPlugin"`
	NetworkPolicy       *string                            `pulumi:"networkPolicy"`
	OutboundType        *string                            `pulumi:"outboundType"`
	PodCidr             *string                            `pulumi:"podCidr"`
	ServiceCidr         *string                            `pulumi:"serviceCidr"`
}


func (val *ContainerServiceNetworkProfile) Defaults() *ContainerServiceNetworkProfile {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DnsServiceIP) {
		dnsServiceIP_ := "10.0.0.10"
		tmp.DnsServiceIP = &dnsServiceIP_
	}
	if isZero(tmp.DockerBridgeCidr) {
		dockerBridgeCidr_ := "172.17.0.1/16"
		tmp.DockerBridgeCidr = &dockerBridgeCidr_
	}
	tmp.LoadBalancerProfile = tmp.LoadBalancerProfile.Defaults()

	if isZero(tmp.NetworkPlugin) {
		networkPlugin_ := "kubenet"
		tmp.NetworkPlugin = &networkPlugin_
	}
	if isZero(tmp.OutboundType) {
		outboundType_ := "loadBalancer"
		tmp.OutboundType = &outboundType_
	}
	if isZero(tmp.PodCidr) {
		podCidr_ := "10.244.0.0/16"
		tmp.PodCidr = &podCidr_
	}
	if isZero(tmp.ServiceCidr) {
		serviceCidr_ := "10.0.0.0/16"
		tmp.ServiceCidr = &serviceCidr_
	}
	return &tmp
}





type ContainerServiceNetworkProfileInput interface {
	pulumi.Input

	ToContainerServiceNetworkProfileOutput() ContainerServiceNetworkProfileOutput
	ToContainerServiceNetworkProfileOutputWithContext(context.Context) ContainerServiceNetworkProfileOutput
}

type ContainerServiceNetworkProfileArgs struct {
	DnsServiceIP        pulumi.StringPtrInput                     `pulumi:"dnsServiceIP"`
	DockerBridgeCidr    pulumi.StringPtrInput                     `pulumi:"dockerBridgeCidr"`
	LoadBalancerProfile ManagedClusterLoadBalancerProfilePtrInput `pulumi:"loadBalancerProfile"`
	LoadBalancerSku     pulumi.StringPtrInput                     `pulumi:"loadBalancerSku"`
	NetworkMode         pulumi.StringPtrInput                     `pulumi:"networkMode"`
	NetworkPlugin       pulumi.StringPtrInput                     `pulumi:"networkPlugin"`
	NetworkPolicy       pulumi.StringPtrInput                     `pulumi:"networkPolicy"`
	OutboundType        pulumi.StringPtrInput                     `pulumi:"outboundType"`
	PodCidr             pulumi.StringPtrInput                     `pulumi:"podCidr"`
	ServiceCidr         pulumi.StringPtrInput                     `pulumi:"serviceCidr"`
}

func (ContainerServiceNetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceNetworkProfile)(nil)).Elem()
}

func (i ContainerServiceNetworkProfileArgs) ToContainerServiceNetworkProfileOutput() ContainerServiceNetworkProfileOutput {
	return i.ToContainerServiceNetworkProfileOutputWithContext(context.Background())
}

func (i ContainerServiceNetworkProfileArgs) ToContainerServiceNetworkProfileOutputWithContext(ctx context.Context) ContainerServiceNetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceNetworkProfileOutput)
}

func (i ContainerServiceNetworkProfileArgs) ToContainerServiceNetworkProfilePtrOutput() ContainerServiceNetworkProfilePtrOutput {
	return i.ToContainerServiceNetworkProfilePtrOutputWithContext(context.Background())
}

func (i ContainerServiceNetworkProfileArgs) ToContainerServiceNetworkProfilePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceNetworkProfileOutput).ToContainerServiceNetworkProfilePtrOutputWithContext(ctx)
}









type ContainerServiceNetworkProfilePtrInput interface {
	pulumi.Input

	ToContainerServiceNetworkProfilePtrOutput() ContainerServiceNetworkProfilePtrOutput
	ToContainerServiceNetworkProfilePtrOutputWithContext(context.Context) ContainerServiceNetworkProfilePtrOutput
}

type containerServiceNetworkProfilePtrType ContainerServiceNetworkProfileArgs

func ContainerServiceNetworkProfilePtr(v *ContainerServiceNetworkProfileArgs) ContainerServiceNetworkProfilePtrInput {
	return (*containerServiceNetworkProfilePtrType)(v)
}

func (*containerServiceNetworkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceNetworkProfile)(nil)).Elem()
}

func (i *containerServiceNetworkProfilePtrType) ToContainerServiceNetworkProfilePtrOutput() ContainerServiceNetworkProfilePtrOutput {
	return i.ToContainerServiceNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *containerServiceNetworkProfilePtrType) ToContainerServiceNetworkProfilePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceNetworkProfilePtrOutput)
}

type ContainerServiceNetworkProfileOutput struct{ *pulumi.OutputState }

func (ContainerServiceNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceNetworkProfile)(nil)).Elem()
}

func (o ContainerServiceNetworkProfileOutput) ToContainerServiceNetworkProfileOutput() ContainerServiceNetworkProfileOutput {
	return o
}

func (o ContainerServiceNetworkProfileOutput) ToContainerServiceNetworkProfileOutputWithContext(ctx context.Context) ContainerServiceNetworkProfileOutput {
	return o
}

func (o ContainerServiceNetworkProfileOutput) ToContainerServiceNetworkProfilePtrOutput() ContainerServiceNetworkProfilePtrOutput {
	return o.ToContainerServiceNetworkProfilePtrOutputWithContext(context.Background())
}

func (o ContainerServiceNetworkProfileOutput) ToContainerServiceNetworkProfilePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceNetworkProfile) *ContainerServiceNetworkProfile {
		return &v
	}).(ContainerServiceNetworkProfilePtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.DnsServiceIP }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.DockerBridgeCidr }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) LoadBalancerProfile() ManagedClusterLoadBalancerProfilePtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *ManagedClusterLoadBalancerProfile {
		return v.LoadBalancerProfile
	}).(ManagedClusterLoadBalancerProfilePtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) LoadBalancerSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.LoadBalancerSku }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) NetworkMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.NetworkMode }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) NetworkPlugin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.NetworkPlugin }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) NetworkPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.NetworkPolicy }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) OutboundType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.OutboundType }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.PodCidr }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.ServiceCidr }).(pulumi.StringPtrOutput)
}

type ContainerServiceNetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceNetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceNetworkProfile)(nil)).Elem()
}

func (o ContainerServiceNetworkProfilePtrOutput) ToContainerServiceNetworkProfilePtrOutput() ContainerServiceNetworkProfilePtrOutput {
	return o
}

func (o ContainerServiceNetworkProfilePtrOutput) ToContainerServiceNetworkProfilePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfilePtrOutput {
	return o
}

func (o ContainerServiceNetworkProfilePtrOutput) Elem() ContainerServiceNetworkProfileOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) ContainerServiceNetworkProfile {
		if v != nil {
			return *v
		}
		var ret ContainerServiceNetworkProfile
		return ret
	}).(ContainerServiceNetworkProfileOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.DnsServiceIP
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.DockerBridgeCidr
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) LoadBalancerProfile() ManagedClusterLoadBalancerProfilePtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *ManagedClusterLoadBalancerProfile {
		if v == nil {
			return nil
		}
		return v.LoadBalancerProfile
	}).(ManagedClusterLoadBalancerProfilePtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) LoadBalancerSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancerSku
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) NetworkMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.NetworkMode
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) NetworkPlugin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.NetworkPlugin
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) NetworkPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.NetworkPolicy
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) OutboundType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.OutboundType
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.PodCidr
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCidr
	}).(pulumi.StringPtrOutput)
}

type ContainerServiceNetworkProfileResponse struct {
	DnsServiceIP        *string                                    `pulumi:"dnsServiceIP"`
	DockerBridgeCidr    *string                                    `pulumi:"dockerBridgeCidr"`
	LoadBalancerProfile *ManagedClusterLoadBalancerProfileResponse `pulumi:"loadBalancerProfile"`
	LoadBalancerSku     *string                                    `pulumi:"loadBalancerSku"`
	NetworkMode         *string                                    `pulumi:"networkMode"`
	NetworkPlugin       *string                                    `pulumi:"networkPlugin"`
	NetworkPolicy       *string                                    `pulumi:"networkPolicy"`
	OutboundType        *string                                    `pulumi:"outboundType"`
	PodCidr             *string                                    `pulumi:"podCidr"`
	ServiceCidr         *string                                    `pulumi:"serviceCidr"`
}


func (val *ContainerServiceNetworkProfileResponse) Defaults() *ContainerServiceNetworkProfileResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DnsServiceIP) {
		dnsServiceIP_ := "10.0.0.10"
		tmp.DnsServiceIP = &dnsServiceIP_
	}
	if isZero(tmp.DockerBridgeCidr) {
		dockerBridgeCidr_ := "172.17.0.1/16"
		tmp.DockerBridgeCidr = &dockerBridgeCidr_
	}
	tmp.LoadBalancerProfile = tmp.LoadBalancerProfile.Defaults()

	if isZero(tmp.NetworkPlugin) {
		networkPlugin_ := "kubenet"
		tmp.NetworkPlugin = &networkPlugin_
	}
	if isZero(tmp.OutboundType) {
		outboundType_ := "loadBalancer"
		tmp.OutboundType = &outboundType_
	}
	if isZero(tmp.PodCidr) {
		podCidr_ := "10.244.0.0/16"
		tmp.PodCidr = &podCidr_
	}
	if isZero(tmp.ServiceCidr) {
		serviceCidr_ := "10.0.0.0/16"
		tmp.ServiceCidr = &serviceCidr_
	}
	return &tmp
}





type ContainerServiceNetworkProfileResponseInput interface {
	pulumi.Input

	ToContainerServiceNetworkProfileResponseOutput() ContainerServiceNetworkProfileResponseOutput
	ToContainerServiceNetworkProfileResponseOutputWithContext(context.Context) ContainerServiceNetworkProfileResponseOutput
}

type ContainerServiceNetworkProfileResponseArgs struct {
	DnsServiceIP        pulumi.StringPtrInput                             `pulumi:"dnsServiceIP"`
	DockerBridgeCidr    pulumi.StringPtrInput                             `pulumi:"dockerBridgeCidr"`
	LoadBalancerProfile ManagedClusterLoadBalancerProfileResponsePtrInput `pulumi:"loadBalancerProfile"`
	LoadBalancerSku     pulumi.StringPtrInput                             `pulumi:"loadBalancerSku"`
	NetworkMode         pulumi.StringPtrInput                             `pulumi:"networkMode"`
	NetworkPlugin       pulumi.StringPtrInput                             `pulumi:"networkPlugin"`
	NetworkPolicy       pulumi.StringPtrInput                             `pulumi:"networkPolicy"`
	OutboundType        pulumi.StringPtrInput                             `pulumi:"outboundType"`
	PodCidr             pulumi.StringPtrInput                             `pulumi:"podCidr"`
	ServiceCidr         pulumi.StringPtrInput                             `pulumi:"serviceCidr"`
}

func (ContainerServiceNetworkProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceNetworkProfileResponse)(nil)).Elem()
}

func (i ContainerServiceNetworkProfileResponseArgs) ToContainerServiceNetworkProfileResponseOutput() ContainerServiceNetworkProfileResponseOutput {
	return i.ToContainerServiceNetworkProfileResponseOutputWithContext(context.Background())
}

func (i ContainerServiceNetworkProfileResponseArgs) ToContainerServiceNetworkProfileResponseOutputWithContext(ctx context.Context) ContainerServiceNetworkProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceNetworkProfileResponseOutput)
}

func (i ContainerServiceNetworkProfileResponseArgs) ToContainerServiceNetworkProfileResponsePtrOutput() ContainerServiceNetworkProfileResponsePtrOutput {
	return i.ToContainerServiceNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (i ContainerServiceNetworkProfileResponseArgs) ToContainerServiceNetworkProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceNetworkProfileResponseOutput).ToContainerServiceNetworkProfileResponsePtrOutputWithContext(ctx)
}









type ContainerServiceNetworkProfileResponsePtrInput interface {
	pulumi.Input

	ToContainerServiceNetworkProfileResponsePtrOutput() ContainerServiceNetworkProfileResponsePtrOutput
	ToContainerServiceNetworkProfileResponsePtrOutputWithContext(context.Context) ContainerServiceNetworkProfileResponsePtrOutput
}

type containerServiceNetworkProfileResponsePtrType ContainerServiceNetworkProfileResponseArgs

func ContainerServiceNetworkProfileResponsePtr(v *ContainerServiceNetworkProfileResponseArgs) ContainerServiceNetworkProfileResponsePtrInput {
	return (*containerServiceNetworkProfileResponsePtrType)(v)
}

func (*containerServiceNetworkProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceNetworkProfileResponse)(nil)).Elem()
}

func (i *containerServiceNetworkProfileResponsePtrType) ToContainerServiceNetworkProfileResponsePtrOutput() ContainerServiceNetworkProfileResponsePtrOutput {
	return i.ToContainerServiceNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (i *containerServiceNetworkProfileResponsePtrType) ToContainerServiceNetworkProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceNetworkProfileResponsePtrOutput)
}

type ContainerServiceNetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceNetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceNetworkProfileResponse)(nil)).Elem()
}

func (o ContainerServiceNetworkProfileResponseOutput) ToContainerServiceNetworkProfileResponseOutput() ContainerServiceNetworkProfileResponseOutput {
	return o
}

func (o ContainerServiceNetworkProfileResponseOutput) ToContainerServiceNetworkProfileResponseOutputWithContext(ctx context.Context) ContainerServiceNetworkProfileResponseOutput {
	return o
}

func (o ContainerServiceNetworkProfileResponseOutput) ToContainerServiceNetworkProfileResponsePtrOutput() ContainerServiceNetworkProfileResponsePtrOutput {
	return o.ToContainerServiceNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (o ContainerServiceNetworkProfileResponseOutput) ToContainerServiceNetworkProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceNetworkProfileResponse) *ContainerServiceNetworkProfileResponse {
		return &v
	}).(ContainerServiceNetworkProfileResponsePtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.DnsServiceIP }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.DockerBridgeCidr }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) LoadBalancerProfile() ManagedClusterLoadBalancerProfileResponsePtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *ManagedClusterLoadBalancerProfileResponse {
		return v.LoadBalancerProfile
	}).(ManagedClusterLoadBalancerProfileResponsePtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) LoadBalancerSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.LoadBalancerSku }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) NetworkMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.NetworkMode }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) NetworkPlugin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.NetworkPlugin }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) NetworkPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.NetworkPolicy }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) OutboundType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.OutboundType }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.PodCidr }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.ServiceCidr }).(pulumi.StringPtrOutput)
}

type ContainerServiceNetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceNetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceNetworkProfileResponse)(nil)).Elem()
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) ToContainerServiceNetworkProfileResponsePtrOutput() ContainerServiceNetworkProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) ToContainerServiceNetworkProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) Elem() ContainerServiceNetworkProfileResponseOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) ContainerServiceNetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret ContainerServiceNetworkProfileResponse
		return ret
	}).(ContainerServiceNetworkProfileResponseOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.DnsServiceIP
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.DockerBridgeCidr
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) LoadBalancerProfile() ManagedClusterLoadBalancerProfileResponsePtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *ManagedClusterLoadBalancerProfileResponse {
		if v == nil {
			return nil
		}
		return v.LoadBalancerProfile
	}).(ManagedClusterLoadBalancerProfileResponsePtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) LoadBalancerSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancerSku
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) NetworkMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.NetworkMode
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) NetworkPlugin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.NetworkPlugin
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) NetworkPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.NetworkPolicy
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) OutboundType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.OutboundType
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.PodCidr
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCidr
	}).(pulumi.StringPtrOutput)
}

type ContainerServiceSshConfiguration struct {
	PublicKeys []ContainerServiceSshPublicKey `pulumi:"publicKeys"`
}





type ContainerServiceSshConfigurationInput interface {
	pulumi.Input

	ToContainerServiceSshConfigurationOutput() ContainerServiceSshConfigurationOutput
	ToContainerServiceSshConfigurationOutputWithContext(context.Context) ContainerServiceSshConfigurationOutput
}

type ContainerServiceSshConfigurationArgs struct {
	PublicKeys ContainerServiceSshPublicKeyArrayInput `pulumi:"publicKeys"`
}

func (ContainerServiceSshConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshConfiguration)(nil)).Elem()
}

func (i ContainerServiceSshConfigurationArgs) ToContainerServiceSshConfigurationOutput() ContainerServiceSshConfigurationOutput {
	return i.ToContainerServiceSshConfigurationOutputWithContext(context.Background())
}

func (i ContainerServiceSshConfigurationArgs) ToContainerServiceSshConfigurationOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshConfigurationOutput)
}

func (i ContainerServiceSshConfigurationArgs) ToContainerServiceSshConfigurationPtrOutput() ContainerServiceSshConfigurationPtrOutput {
	return i.ToContainerServiceSshConfigurationPtrOutputWithContext(context.Background())
}

func (i ContainerServiceSshConfigurationArgs) ToContainerServiceSshConfigurationPtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshConfigurationOutput).ToContainerServiceSshConfigurationPtrOutputWithContext(ctx)
}









type ContainerServiceSshConfigurationPtrInput interface {
	pulumi.Input

	ToContainerServiceSshConfigurationPtrOutput() ContainerServiceSshConfigurationPtrOutput
	ToContainerServiceSshConfigurationPtrOutputWithContext(context.Context) ContainerServiceSshConfigurationPtrOutput
}

type containerServiceSshConfigurationPtrType ContainerServiceSshConfigurationArgs

func ContainerServiceSshConfigurationPtr(v *ContainerServiceSshConfigurationArgs) ContainerServiceSshConfigurationPtrInput {
	return (*containerServiceSshConfigurationPtrType)(v)
}

func (*containerServiceSshConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceSshConfiguration)(nil)).Elem()
}

func (i *containerServiceSshConfigurationPtrType) ToContainerServiceSshConfigurationPtrOutput() ContainerServiceSshConfigurationPtrOutput {
	return i.ToContainerServiceSshConfigurationPtrOutputWithContext(context.Background())
}

func (i *containerServiceSshConfigurationPtrType) ToContainerServiceSshConfigurationPtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshConfigurationPtrOutput)
}

type ContainerServiceSshConfigurationOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshConfiguration)(nil)).Elem()
}

func (o ContainerServiceSshConfigurationOutput) ToContainerServiceSshConfigurationOutput() ContainerServiceSshConfigurationOutput {
	return o
}

func (o ContainerServiceSshConfigurationOutput) ToContainerServiceSshConfigurationOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationOutput {
	return o
}

func (o ContainerServiceSshConfigurationOutput) ToContainerServiceSshConfigurationPtrOutput() ContainerServiceSshConfigurationPtrOutput {
	return o.ToContainerServiceSshConfigurationPtrOutputWithContext(context.Background())
}

func (o ContainerServiceSshConfigurationOutput) ToContainerServiceSshConfigurationPtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceSshConfiguration) *ContainerServiceSshConfiguration {
		return &v
	}).(ContainerServiceSshConfigurationPtrOutput)
}

func (o ContainerServiceSshConfigurationOutput) PublicKeys() ContainerServiceSshPublicKeyArrayOutput {
	return o.ApplyT(func(v ContainerServiceSshConfiguration) []ContainerServiceSshPublicKey { return v.PublicKeys }).(ContainerServiceSshPublicKeyArrayOutput)
}

type ContainerServiceSshConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceSshConfiguration)(nil)).Elem()
}

func (o ContainerServiceSshConfigurationPtrOutput) ToContainerServiceSshConfigurationPtrOutput() ContainerServiceSshConfigurationPtrOutput {
	return o
}

func (o ContainerServiceSshConfigurationPtrOutput) ToContainerServiceSshConfigurationPtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationPtrOutput {
	return o
}

func (o ContainerServiceSshConfigurationPtrOutput) Elem() ContainerServiceSshConfigurationOutput {
	return o.ApplyT(func(v *ContainerServiceSshConfiguration) ContainerServiceSshConfiguration {
		if v != nil {
			return *v
		}
		var ret ContainerServiceSshConfiguration
		return ret
	}).(ContainerServiceSshConfigurationOutput)
}

func (o ContainerServiceSshConfigurationPtrOutput) PublicKeys() ContainerServiceSshPublicKeyArrayOutput {
	return o.ApplyT(func(v *ContainerServiceSshConfiguration) []ContainerServiceSshPublicKey {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(ContainerServiceSshPublicKeyArrayOutput)
}

type ContainerServiceSshConfigurationResponse struct {
	PublicKeys []ContainerServiceSshPublicKeyResponse `pulumi:"publicKeys"`
}





type ContainerServiceSshConfigurationResponseInput interface {
	pulumi.Input

	ToContainerServiceSshConfigurationResponseOutput() ContainerServiceSshConfigurationResponseOutput
	ToContainerServiceSshConfigurationResponseOutputWithContext(context.Context) ContainerServiceSshConfigurationResponseOutput
}

type ContainerServiceSshConfigurationResponseArgs struct {
	PublicKeys ContainerServiceSshPublicKeyResponseArrayInput `pulumi:"publicKeys"`
}

func (ContainerServiceSshConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshConfigurationResponse)(nil)).Elem()
}

func (i ContainerServiceSshConfigurationResponseArgs) ToContainerServiceSshConfigurationResponseOutput() ContainerServiceSshConfigurationResponseOutput {
	return i.ToContainerServiceSshConfigurationResponseOutputWithContext(context.Background())
}

func (i ContainerServiceSshConfigurationResponseArgs) ToContainerServiceSshConfigurationResponseOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshConfigurationResponseOutput)
}

func (i ContainerServiceSshConfigurationResponseArgs) ToContainerServiceSshConfigurationResponsePtrOutput() ContainerServiceSshConfigurationResponsePtrOutput {
	return i.ToContainerServiceSshConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i ContainerServiceSshConfigurationResponseArgs) ToContainerServiceSshConfigurationResponsePtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshConfigurationResponseOutput).ToContainerServiceSshConfigurationResponsePtrOutputWithContext(ctx)
}









type ContainerServiceSshConfigurationResponsePtrInput interface {
	pulumi.Input

	ToContainerServiceSshConfigurationResponsePtrOutput() ContainerServiceSshConfigurationResponsePtrOutput
	ToContainerServiceSshConfigurationResponsePtrOutputWithContext(context.Context) ContainerServiceSshConfigurationResponsePtrOutput
}

type containerServiceSshConfigurationResponsePtrType ContainerServiceSshConfigurationResponseArgs

func ContainerServiceSshConfigurationResponsePtr(v *ContainerServiceSshConfigurationResponseArgs) ContainerServiceSshConfigurationResponsePtrInput {
	return (*containerServiceSshConfigurationResponsePtrType)(v)
}

func (*containerServiceSshConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceSshConfigurationResponse)(nil)).Elem()
}

func (i *containerServiceSshConfigurationResponsePtrType) ToContainerServiceSshConfigurationResponsePtrOutput() ContainerServiceSshConfigurationResponsePtrOutput {
	return i.ToContainerServiceSshConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *containerServiceSshConfigurationResponsePtrType) ToContainerServiceSshConfigurationResponsePtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshConfigurationResponsePtrOutput)
}

type ContainerServiceSshConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshConfigurationResponse)(nil)).Elem()
}

func (o ContainerServiceSshConfigurationResponseOutput) ToContainerServiceSshConfigurationResponseOutput() ContainerServiceSshConfigurationResponseOutput {
	return o
}

func (o ContainerServiceSshConfigurationResponseOutput) ToContainerServiceSshConfigurationResponseOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationResponseOutput {
	return o
}

func (o ContainerServiceSshConfigurationResponseOutput) ToContainerServiceSshConfigurationResponsePtrOutput() ContainerServiceSshConfigurationResponsePtrOutput {
	return o.ToContainerServiceSshConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o ContainerServiceSshConfigurationResponseOutput) ToContainerServiceSshConfigurationResponsePtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceSshConfigurationResponse) *ContainerServiceSshConfigurationResponse {
		return &v
	}).(ContainerServiceSshConfigurationResponsePtrOutput)
}

func (o ContainerServiceSshConfigurationResponseOutput) PublicKeys() ContainerServiceSshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v ContainerServiceSshConfigurationResponse) []ContainerServiceSshPublicKeyResponse {
		return v.PublicKeys
	}).(ContainerServiceSshPublicKeyResponseArrayOutput)
}

type ContainerServiceSshConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceSshConfigurationResponse)(nil)).Elem()
}

func (o ContainerServiceSshConfigurationResponsePtrOutput) ToContainerServiceSshConfigurationResponsePtrOutput() ContainerServiceSshConfigurationResponsePtrOutput {
	return o
}

func (o ContainerServiceSshConfigurationResponsePtrOutput) ToContainerServiceSshConfigurationResponsePtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationResponsePtrOutput {
	return o
}

func (o ContainerServiceSshConfigurationResponsePtrOutput) Elem() ContainerServiceSshConfigurationResponseOutput {
	return o.ApplyT(func(v *ContainerServiceSshConfigurationResponse) ContainerServiceSshConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ContainerServiceSshConfigurationResponse
		return ret
	}).(ContainerServiceSshConfigurationResponseOutput)
}

func (o ContainerServiceSshConfigurationResponsePtrOutput) PublicKeys() ContainerServiceSshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v *ContainerServiceSshConfigurationResponse) []ContainerServiceSshPublicKeyResponse {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(ContainerServiceSshPublicKeyResponseArrayOutput)
}

type ContainerServiceSshPublicKey struct {
	KeyData string `pulumi:"keyData"`
}





type ContainerServiceSshPublicKeyInput interface {
	pulumi.Input

	ToContainerServiceSshPublicKeyOutput() ContainerServiceSshPublicKeyOutput
	ToContainerServiceSshPublicKeyOutputWithContext(context.Context) ContainerServiceSshPublicKeyOutput
}

type ContainerServiceSshPublicKeyArgs struct {
	KeyData pulumi.StringInput `pulumi:"keyData"`
}

func (ContainerServiceSshPublicKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshPublicKey)(nil)).Elem()
}

func (i ContainerServiceSshPublicKeyArgs) ToContainerServiceSshPublicKeyOutput() ContainerServiceSshPublicKeyOutput {
	return i.ToContainerServiceSshPublicKeyOutputWithContext(context.Background())
}

func (i ContainerServiceSshPublicKeyArgs) ToContainerServiceSshPublicKeyOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshPublicKeyOutput)
}





type ContainerServiceSshPublicKeyArrayInput interface {
	pulumi.Input

	ToContainerServiceSshPublicKeyArrayOutput() ContainerServiceSshPublicKeyArrayOutput
	ToContainerServiceSshPublicKeyArrayOutputWithContext(context.Context) ContainerServiceSshPublicKeyArrayOutput
}

type ContainerServiceSshPublicKeyArray []ContainerServiceSshPublicKeyInput

func (ContainerServiceSshPublicKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerServiceSshPublicKey)(nil)).Elem()
}

func (i ContainerServiceSshPublicKeyArray) ToContainerServiceSshPublicKeyArrayOutput() ContainerServiceSshPublicKeyArrayOutput {
	return i.ToContainerServiceSshPublicKeyArrayOutputWithContext(context.Background())
}

func (i ContainerServiceSshPublicKeyArray) ToContainerServiceSshPublicKeyArrayOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshPublicKeyArrayOutput)
}

type ContainerServiceSshPublicKeyOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshPublicKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshPublicKey)(nil)).Elem()
}

func (o ContainerServiceSshPublicKeyOutput) ToContainerServiceSshPublicKeyOutput() ContainerServiceSshPublicKeyOutput {
	return o
}

func (o ContainerServiceSshPublicKeyOutput) ToContainerServiceSshPublicKeyOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyOutput {
	return o
}

func (o ContainerServiceSshPublicKeyOutput) KeyData() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceSshPublicKey) string { return v.KeyData }).(pulumi.StringOutput)
}

type ContainerServiceSshPublicKeyArrayOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshPublicKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerServiceSshPublicKey)(nil)).Elem()
}

func (o ContainerServiceSshPublicKeyArrayOutput) ToContainerServiceSshPublicKeyArrayOutput() ContainerServiceSshPublicKeyArrayOutput {
	return o
}

func (o ContainerServiceSshPublicKeyArrayOutput) ToContainerServiceSshPublicKeyArrayOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyArrayOutput {
	return o
}

func (o ContainerServiceSshPublicKeyArrayOutput) Index(i pulumi.IntInput) ContainerServiceSshPublicKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerServiceSshPublicKey {
		return vs[0].([]ContainerServiceSshPublicKey)[vs[1].(int)]
	}).(ContainerServiceSshPublicKeyOutput)
}

type ContainerServiceSshPublicKeyResponse struct {
	KeyData string `pulumi:"keyData"`
}





type ContainerServiceSshPublicKeyResponseInput interface {
	pulumi.Input

	ToContainerServiceSshPublicKeyResponseOutput() ContainerServiceSshPublicKeyResponseOutput
	ToContainerServiceSshPublicKeyResponseOutputWithContext(context.Context) ContainerServiceSshPublicKeyResponseOutput
}

type ContainerServiceSshPublicKeyResponseArgs struct {
	KeyData pulumi.StringInput `pulumi:"keyData"`
}

func (ContainerServiceSshPublicKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshPublicKeyResponse)(nil)).Elem()
}

func (i ContainerServiceSshPublicKeyResponseArgs) ToContainerServiceSshPublicKeyResponseOutput() ContainerServiceSshPublicKeyResponseOutput {
	return i.ToContainerServiceSshPublicKeyResponseOutputWithContext(context.Background())
}

func (i ContainerServiceSshPublicKeyResponseArgs) ToContainerServiceSshPublicKeyResponseOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshPublicKeyResponseOutput)
}





type ContainerServiceSshPublicKeyResponseArrayInput interface {
	pulumi.Input

	ToContainerServiceSshPublicKeyResponseArrayOutput() ContainerServiceSshPublicKeyResponseArrayOutput
	ToContainerServiceSshPublicKeyResponseArrayOutputWithContext(context.Context) ContainerServiceSshPublicKeyResponseArrayOutput
}

type ContainerServiceSshPublicKeyResponseArray []ContainerServiceSshPublicKeyResponseInput

func (ContainerServiceSshPublicKeyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerServiceSshPublicKeyResponse)(nil)).Elem()
}

func (i ContainerServiceSshPublicKeyResponseArray) ToContainerServiceSshPublicKeyResponseArrayOutput() ContainerServiceSshPublicKeyResponseArrayOutput {
	return i.ToContainerServiceSshPublicKeyResponseArrayOutputWithContext(context.Background())
}

func (i ContainerServiceSshPublicKeyResponseArray) ToContainerServiceSshPublicKeyResponseArrayOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshPublicKeyResponseArrayOutput)
}

type ContainerServiceSshPublicKeyResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshPublicKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshPublicKeyResponse)(nil)).Elem()
}

func (o ContainerServiceSshPublicKeyResponseOutput) ToContainerServiceSshPublicKeyResponseOutput() ContainerServiceSshPublicKeyResponseOutput {
	return o
}

func (o ContainerServiceSshPublicKeyResponseOutput) ToContainerServiceSshPublicKeyResponseOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyResponseOutput {
	return o
}

func (o ContainerServiceSshPublicKeyResponseOutput) KeyData() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceSshPublicKeyResponse) string { return v.KeyData }).(pulumi.StringOutput)
}

type ContainerServiceSshPublicKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshPublicKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerServiceSshPublicKeyResponse)(nil)).Elem()
}

func (o ContainerServiceSshPublicKeyResponseArrayOutput) ToContainerServiceSshPublicKeyResponseArrayOutput() ContainerServiceSshPublicKeyResponseArrayOutput {
	return o
}

func (o ContainerServiceSshPublicKeyResponseArrayOutput) ToContainerServiceSshPublicKeyResponseArrayOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyResponseArrayOutput {
	return o
}

func (o ContainerServiceSshPublicKeyResponseArrayOutput) Index(i pulumi.IntInput) ContainerServiceSshPublicKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerServiceSshPublicKeyResponse {
		return vs[0].([]ContainerServiceSshPublicKeyResponse)[vs[1].(int)]
	}).(ContainerServiceSshPublicKeyResponseOutput)
}

type CredentialResultResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type CredentialResultResponseInput interface {
	pulumi.Input

	ToCredentialResultResponseOutput() CredentialResultResponseOutput
	ToCredentialResultResponseOutputWithContext(context.Context) CredentialResultResponseOutput
}

type CredentialResultResponseArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (CredentialResultResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CredentialResultResponse)(nil)).Elem()
}

func (i CredentialResultResponseArgs) ToCredentialResultResponseOutput() CredentialResultResponseOutput {
	return i.ToCredentialResultResponseOutputWithContext(context.Background())
}

func (i CredentialResultResponseArgs) ToCredentialResultResponseOutputWithContext(ctx context.Context) CredentialResultResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialResultResponseOutput)
}





type CredentialResultResponseArrayInput interface {
	pulumi.Input

	ToCredentialResultResponseArrayOutput() CredentialResultResponseArrayOutput
	ToCredentialResultResponseArrayOutputWithContext(context.Context) CredentialResultResponseArrayOutput
}

type CredentialResultResponseArray []CredentialResultResponseInput

func (CredentialResultResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CredentialResultResponse)(nil)).Elem()
}

func (i CredentialResultResponseArray) ToCredentialResultResponseArrayOutput() CredentialResultResponseArrayOutput {
	return i.ToCredentialResultResponseArrayOutputWithContext(context.Background())
}

func (i CredentialResultResponseArray) ToCredentialResultResponseArrayOutputWithContext(ctx context.Context) CredentialResultResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialResultResponseArrayOutput)
}

type CredentialResultResponseOutput struct{ *pulumi.OutputState }

func (CredentialResultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CredentialResultResponse)(nil)).Elem()
}

func (o CredentialResultResponseOutput) ToCredentialResultResponseOutput() CredentialResultResponseOutput {
	return o
}

func (o CredentialResultResponseOutput) ToCredentialResultResponseOutputWithContext(ctx context.Context) CredentialResultResponseOutput {
	return o
}

func (o CredentialResultResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CredentialResultResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o CredentialResultResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v CredentialResultResponse) string { return v.Value }).(pulumi.StringOutput)
}

type CredentialResultResponseArrayOutput struct{ *pulumi.OutputState }

func (CredentialResultResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CredentialResultResponse)(nil)).Elem()
}

func (o CredentialResultResponseArrayOutput) ToCredentialResultResponseArrayOutput() CredentialResultResponseArrayOutput {
	return o
}

func (o CredentialResultResponseArrayOutput) ToCredentialResultResponseArrayOutputWithContext(ctx context.Context) CredentialResultResponseArrayOutput {
	return o
}

func (o CredentialResultResponseArrayOutput) Index(i pulumi.IntInput) CredentialResultResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CredentialResultResponse {
		return vs[0].([]CredentialResultResponse)[vs[1].(int)]
	}).(CredentialResultResponseOutput)
}

type ExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ExtendedLocationInput interface {
	pulumi.Input

	ToExtendedLocationOutput() ExtendedLocationOutput
	ToExtendedLocationOutputWithContext(context.Context) ExtendedLocationOutput
}

type ExtendedLocationArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ExtendedLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (i ExtendedLocationArgs) ToExtendedLocationOutput() ExtendedLocationOutput {
	return i.ToExtendedLocationOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput)
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput).ToExtendedLocationPtrOutputWithContext(ctx)
}









type ExtendedLocationPtrInput interface {
	pulumi.Input

	ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput
	ToExtendedLocationPtrOutputWithContext(context.Context) ExtendedLocationPtrOutput
}

type extendedLocationPtrType ExtendedLocationArgs

func ExtendedLocationPtr(v *ExtendedLocationArgs) ExtendedLocationPtrInput {
	return (*extendedLocationPtrType)(v)
}

func (*extendedLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationPtrOutput)
}

type ExtendedLocationOutput struct{ *pulumi.OutputState }

func (ExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationOutput) ToExtendedLocationOutput() ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocation) *ExtendedLocation {
		return &v
	}).(ExtendedLocationPtrOutput)
}

func (o ExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) Elem() ExtendedLocationOutput {
	return o.ApplyT(func(v *ExtendedLocation) ExtendedLocation {
		if v != nil {
			return *v
		}
		var ret ExtendedLocation
		return ret
	}).(ExtendedLocationOutput)
}

func (o ExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ExtendedLocationResponseInput interface {
	pulumi.Input

	ToExtendedLocationResponseOutput() ExtendedLocationResponseOutput
	ToExtendedLocationResponseOutputWithContext(context.Context) ExtendedLocationResponseOutput
}

type ExtendedLocationResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ExtendedLocationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationResponse)(nil)).Elem()
}

func (i ExtendedLocationResponseArgs) ToExtendedLocationResponseOutput() ExtendedLocationResponseOutput {
	return i.ToExtendedLocationResponseOutputWithContext(context.Background())
}

func (i ExtendedLocationResponseArgs) ToExtendedLocationResponseOutputWithContext(ctx context.Context) ExtendedLocationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationResponseOutput)
}

func (i ExtendedLocationResponseArgs) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return i.ToExtendedLocationResponsePtrOutputWithContext(context.Background())
}

func (i ExtendedLocationResponseArgs) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationResponseOutput).ToExtendedLocationResponsePtrOutputWithContext(ctx)
}









type ExtendedLocationResponsePtrInput interface {
	pulumi.Input

	ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput
	ToExtendedLocationResponsePtrOutputWithContext(context.Context) ExtendedLocationResponsePtrOutput
}

type extendedLocationResponsePtrType ExtendedLocationResponseArgs

func ExtendedLocationResponsePtr(v *ExtendedLocationResponseArgs) ExtendedLocationResponsePtrInput {
	return (*extendedLocationResponsePtrType)(v)
}

func (*extendedLocationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationResponse)(nil)).Elem()
}

func (i *extendedLocationResponsePtrType) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return i.ToExtendedLocationResponsePtrOutputWithContext(context.Background())
}

func (i *extendedLocationResponsePtrType) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationResponsePtrOutput)
}

type ExtendedLocationResponseOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutput() ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutputWithContext(ctx context.Context) ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return o.ToExtendedLocationResponsePtrOutputWithContext(context.Background())
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocationResponse) *ExtendedLocationResponse {
		return &v
	}).(ExtendedLocationResponsePtrOutput)
}

func (o ExtendedLocationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponsePtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) Elem() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) ExtendedLocationResponse {
		if v != nil {
			return *v
		}
		var ret ExtendedLocationResponse
		return ret
	}).(ExtendedLocationResponseOutput)
}

func (o ExtendedLocationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type KubeletConfig struct {
	AllowedUnsafeSysctls  []string `pulumi:"allowedUnsafeSysctls"`
	ContainerLogMaxFiles  *int     `pulumi:"containerLogMaxFiles"`
	ContainerLogMaxSizeMB *int     `pulumi:"containerLogMaxSizeMB"`
	CpuCfsQuota           *bool    `pulumi:"cpuCfsQuota"`
	CpuCfsQuotaPeriod     *string  `pulumi:"cpuCfsQuotaPeriod"`
	CpuManagerPolicy      *string  `pulumi:"cpuManagerPolicy"`
	FailSwapOn            *bool    `pulumi:"failSwapOn"`
	ImageGcHighThreshold  *int     `pulumi:"imageGcHighThreshold"`
	ImageGcLowThreshold   *int     `pulumi:"imageGcLowThreshold"`
	PodMaxPids            *int     `pulumi:"podMaxPids"`
	TopologyManagerPolicy *string  `pulumi:"topologyManagerPolicy"`
}





type KubeletConfigInput interface {
	pulumi.Input

	ToKubeletConfigOutput() KubeletConfigOutput
	ToKubeletConfigOutputWithContext(context.Context) KubeletConfigOutput
}

type KubeletConfigArgs struct {
	AllowedUnsafeSysctls  pulumi.StringArrayInput `pulumi:"allowedUnsafeSysctls"`
	ContainerLogMaxFiles  pulumi.IntPtrInput      `pulumi:"containerLogMaxFiles"`
	ContainerLogMaxSizeMB pulumi.IntPtrInput      `pulumi:"containerLogMaxSizeMB"`
	CpuCfsQuota           pulumi.BoolPtrInput     `pulumi:"cpuCfsQuota"`
	CpuCfsQuotaPeriod     pulumi.StringPtrInput   `pulumi:"cpuCfsQuotaPeriod"`
	CpuManagerPolicy      pulumi.StringPtrInput   `pulumi:"cpuManagerPolicy"`
	FailSwapOn            pulumi.BoolPtrInput     `pulumi:"failSwapOn"`
	ImageGcHighThreshold  pulumi.IntPtrInput      `pulumi:"imageGcHighThreshold"`
	ImageGcLowThreshold   pulumi.IntPtrInput      `pulumi:"imageGcLowThreshold"`
	PodMaxPids            pulumi.IntPtrInput      `pulumi:"podMaxPids"`
	TopologyManagerPolicy pulumi.StringPtrInput   `pulumi:"topologyManagerPolicy"`
}

func (KubeletConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeletConfig)(nil)).Elem()
}

func (i KubeletConfigArgs) ToKubeletConfigOutput() KubeletConfigOutput {
	return i.ToKubeletConfigOutputWithContext(context.Background())
}

func (i KubeletConfigArgs) ToKubeletConfigOutputWithContext(ctx context.Context) KubeletConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeletConfigOutput)
}

func (i KubeletConfigArgs) ToKubeletConfigPtrOutput() KubeletConfigPtrOutput {
	return i.ToKubeletConfigPtrOutputWithContext(context.Background())
}

func (i KubeletConfigArgs) ToKubeletConfigPtrOutputWithContext(ctx context.Context) KubeletConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeletConfigOutput).ToKubeletConfigPtrOutputWithContext(ctx)
}









type KubeletConfigPtrInput interface {
	pulumi.Input

	ToKubeletConfigPtrOutput() KubeletConfigPtrOutput
	ToKubeletConfigPtrOutputWithContext(context.Context) KubeletConfigPtrOutput
}

type kubeletConfigPtrType KubeletConfigArgs

func KubeletConfigPtr(v *KubeletConfigArgs) KubeletConfigPtrInput {
	return (*kubeletConfigPtrType)(v)
}

func (*kubeletConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeletConfig)(nil)).Elem()
}

func (i *kubeletConfigPtrType) ToKubeletConfigPtrOutput() KubeletConfigPtrOutput {
	return i.ToKubeletConfigPtrOutputWithContext(context.Background())
}

func (i *kubeletConfigPtrType) ToKubeletConfigPtrOutputWithContext(ctx context.Context) KubeletConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeletConfigPtrOutput)
}

type KubeletConfigOutput struct{ *pulumi.OutputState }

func (KubeletConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeletConfig)(nil)).Elem()
}

func (o KubeletConfigOutput) ToKubeletConfigOutput() KubeletConfigOutput {
	return o
}

func (o KubeletConfigOutput) ToKubeletConfigOutputWithContext(ctx context.Context) KubeletConfigOutput {
	return o
}

func (o KubeletConfigOutput) ToKubeletConfigPtrOutput() KubeletConfigPtrOutput {
	return o.ToKubeletConfigPtrOutputWithContext(context.Background())
}

func (o KubeletConfigOutput) ToKubeletConfigPtrOutputWithContext(ctx context.Context) KubeletConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubeletConfig) *KubeletConfig {
		return &v
	}).(KubeletConfigPtrOutput)
}

func (o KubeletConfigOutput) AllowedUnsafeSysctls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v KubeletConfig) []string { return v.AllowedUnsafeSysctls }).(pulumi.StringArrayOutput)
}

func (o KubeletConfigOutput) ContainerLogMaxFiles() pulumi.IntPtrOutput {
	return o.ApplyT(func(v KubeletConfig) *int { return v.ContainerLogMaxFiles }).(pulumi.IntPtrOutput)
}

func (o KubeletConfigOutput) ContainerLogMaxSizeMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v KubeletConfig) *int { return v.ContainerLogMaxSizeMB }).(pulumi.IntPtrOutput)
}

func (o KubeletConfigOutput) CpuCfsQuota() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KubeletConfig) *bool { return v.CpuCfsQuota }).(pulumi.BoolPtrOutput)
}

func (o KubeletConfigOutput) CpuCfsQuotaPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubeletConfig) *string { return v.CpuCfsQuotaPeriod }).(pulumi.StringPtrOutput)
}

func (o KubeletConfigOutput) CpuManagerPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubeletConfig) *string { return v.CpuManagerPolicy }).(pulumi.StringPtrOutput)
}

func (o KubeletConfigOutput) FailSwapOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KubeletConfig) *bool { return v.FailSwapOn }).(pulumi.BoolPtrOutput)
}

func (o KubeletConfigOutput) ImageGcHighThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v KubeletConfig) *int { return v.ImageGcHighThreshold }).(pulumi.IntPtrOutput)
}

func (o KubeletConfigOutput) ImageGcLowThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v KubeletConfig) *int { return v.ImageGcLowThreshold }).(pulumi.IntPtrOutput)
}

func (o KubeletConfigOutput) PodMaxPids() pulumi.IntPtrOutput {
	return o.ApplyT(func(v KubeletConfig) *int { return v.PodMaxPids }).(pulumi.IntPtrOutput)
}

func (o KubeletConfigOutput) TopologyManagerPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubeletConfig) *string { return v.TopologyManagerPolicy }).(pulumi.StringPtrOutput)
}

type KubeletConfigPtrOutput struct{ *pulumi.OutputState }

func (KubeletConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeletConfig)(nil)).Elem()
}

func (o KubeletConfigPtrOutput) ToKubeletConfigPtrOutput() KubeletConfigPtrOutput {
	return o
}

func (o KubeletConfigPtrOutput) ToKubeletConfigPtrOutputWithContext(ctx context.Context) KubeletConfigPtrOutput {
	return o
}

func (o KubeletConfigPtrOutput) Elem() KubeletConfigOutput {
	return o.ApplyT(func(v *KubeletConfig) KubeletConfig {
		if v != nil {
			return *v
		}
		var ret KubeletConfig
		return ret
	}).(KubeletConfigOutput)
}

func (o KubeletConfigPtrOutput) AllowedUnsafeSysctls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KubeletConfig) []string {
		if v == nil {
			return nil
		}
		return v.AllowedUnsafeSysctls
	}).(pulumi.StringArrayOutput)
}

func (o KubeletConfigPtrOutput) ContainerLogMaxFiles() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KubeletConfig) *int {
		if v == nil {
			return nil
		}
		return v.ContainerLogMaxFiles
	}).(pulumi.IntPtrOutput)
}

func (o KubeletConfigPtrOutput) ContainerLogMaxSizeMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KubeletConfig) *int {
		if v == nil {
			return nil
		}
		return v.ContainerLogMaxSizeMB
	}).(pulumi.IntPtrOutput)
}

func (o KubeletConfigPtrOutput) CpuCfsQuota() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KubeletConfig) *bool {
		if v == nil {
			return nil
		}
		return v.CpuCfsQuota
	}).(pulumi.BoolPtrOutput)
}

func (o KubeletConfigPtrOutput) CpuCfsQuotaPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeletConfig) *string {
		if v == nil {
			return nil
		}
		return v.CpuCfsQuotaPeriod
	}).(pulumi.StringPtrOutput)
}

func (o KubeletConfigPtrOutput) CpuManagerPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeletConfig) *string {
		if v == nil {
			return nil
		}
		return v.CpuManagerPolicy
	}).(pulumi.StringPtrOutput)
}

func (o KubeletConfigPtrOutput) FailSwapOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KubeletConfig) *bool {
		if v == nil {
			return nil
		}
		return v.FailSwapOn
	}).(pulumi.BoolPtrOutput)
}

func (o KubeletConfigPtrOutput) ImageGcHighThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KubeletConfig) *int {
		if v == nil {
			return nil
		}
		return v.ImageGcHighThreshold
	}).(pulumi.IntPtrOutput)
}

func (o KubeletConfigPtrOutput) ImageGcLowThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KubeletConfig) *int {
		if v == nil {
			return nil
		}
		return v.ImageGcLowThreshold
	}).(pulumi.IntPtrOutput)
}

func (o KubeletConfigPtrOutput) PodMaxPids() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KubeletConfig) *int {
		if v == nil {
			return nil
		}
		return v.PodMaxPids
	}).(pulumi.IntPtrOutput)
}

func (o KubeletConfigPtrOutput) TopologyManagerPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeletConfig) *string {
		if v == nil {
			return nil
		}
		return v.TopologyManagerPolicy
	}).(pulumi.StringPtrOutput)
}

type KubeletConfigResponse struct {
	AllowedUnsafeSysctls  []string `pulumi:"allowedUnsafeSysctls"`
	ContainerLogMaxFiles  *int     `pulumi:"containerLogMaxFiles"`
	ContainerLogMaxSizeMB *int     `pulumi:"containerLogMaxSizeMB"`
	CpuCfsQuota           *bool    `pulumi:"cpuCfsQuota"`
	CpuCfsQuotaPeriod     *string  `pulumi:"cpuCfsQuotaPeriod"`
	CpuManagerPolicy      *string  `pulumi:"cpuManagerPolicy"`
	FailSwapOn            *bool    `pulumi:"failSwapOn"`
	ImageGcHighThreshold  *int     `pulumi:"imageGcHighThreshold"`
	ImageGcLowThreshold   *int     `pulumi:"imageGcLowThreshold"`
	PodMaxPids            *int     `pulumi:"podMaxPids"`
	TopologyManagerPolicy *string  `pulumi:"topologyManagerPolicy"`
}





type KubeletConfigResponseInput interface {
	pulumi.Input

	ToKubeletConfigResponseOutput() KubeletConfigResponseOutput
	ToKubeletConfigResponseOutputWithContext(context.Context) KubeletConfigResponseOutput
}

type KubeletConfigResponseArgs struct {
	AllowedUnsafeSysctls  pulumi.StringArrayInput `pulumi:"allowedUnsafeSysctls"`
	ContainerLogMaxFiles  pulumi.IntPtrInput      `pulumi:"containerLogMaxFiles"`
	ContainerLogMaxSizeMB pulumi.IntPtrInput      `pulumi:"containerLogMaxSizeMB"`
	CpuCfsQuota           pulumi.BoolPtrInput     `pulumi:"cpuCfsQuota"`
	CpuCfsQuotaPeriod     pulumi.StringPtrInput   `pulumi:"cpuCfsQuotaPeriod"`
	CpuManagerPolicy      pulumi.StringPtrInput   `pulumi:"cpuManagerPolicy"`
	FailSwapOn            pulumi.BoolPtrInput     `pulumi:"failSwapOn"`
	ImageGcHighThreshold  pulumi.IntPtrInput      `pulumi:"imageGcHighThreshold"`
	ImageGcLowThreshold   pulumi.IntPtrInput      `pulumi:"imageGcLowThreshold"`
	PodMaxPids            pulumi.IntPtrInput      `pulumi:"podMaxPids"`
	TopologyManagerPolicy pulumi.StringPtrInput   `pulumi:"topologyManagerPolicy"`
}

func (KubeletConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeletConfigResponse)(nil)).Elem()
}

func (i KubeletConfigResponseArgs) ToKubeletConfigResponseOutput() KubeletConfigResponseOutput {
	return i.ToKubeletConfigResponseOutputWithContext(context.Background())
}

func (i KubeletConfigResponseArgs) ToKubeletConfigResponseOutputWithContext(ctx context.Context) KubeletConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeletConfigResponseOutput)
}

func (i KubeletConfigResponseArgs) ToKubeletConfigResponsePtrOutput() KubeletConfigResponsePtrOutput {
	return i.ToKubeletConfigResponsePtrOutputWithContext(context.Background())
}

func (i KubeletConfigResponseArgs) ToKubeletConfigResponsePtrOutputWithContext(ctx context.Context) KubeletConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeletConfigResponseOutput).ToKubeletConfigResponsePtrOutputWithContext(ctx)
}









type KubeletConfigResponsePtrInput interface {
	pulumi.Input

	ToKubeletConfigResponsePtrOutput() KubeletConfigResponsePtrOutput
	ToKubeletConfigResponsePtrOutputWithContext(context.Context) KubeletConfigResponsePtrOutput
}

type kubeletConfigResponsePtrType KubeletConfigResponseArgs

func KubeletConfigResponsePtr(v *KubeletConfigResponseArgs) KubeletConfigResponsePtrInput {
	return (*kubeletConfigResponsePtrType)(v)
}

func (*kubeletConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeletConfigResponse)(nil)).Elem()
}

func (i *kubeletConfigResponsePtrType) ToKubeletConfigResponsePtrOutput() KubeletConfigResponsePtrOutput {
	return i.ToKubeletConfigResponsePtrOutputWithContext(context.Background())
}

func (i *kubeletConfigResponsePtrType) ToKubeletConfigResponsePtrOutputWithContext(ctx context.Context) KubeletConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeletConfigResponsePtrOutput)
}

type KubeletConfigResponseOutput struct{ *pulumi.OutputState }

func (KubeletConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeletConfigResponse)(nil)).Elem()
}

func (o KubeletConfigResponseOutput) ToKubeletConfigResponseOutput() KubeletConfigResponseOutput {
	return o
}

func (o KubeletConfigResponseOutput) ToKubeletConfigResponseOutputWithContext(ctx context.Context) KubeletConfigResponseOutput {
	return o
}

func (o KubeletConfigResponseOutput) ToKubeletConfigResponsePtrOutput() KubeletConfigResponsePtrOutput {
	return o.ToKubeletConfigResponsePtrOutputWithContext(context.Background())
}

func (o KubeletConfigResponseOutput) ToKubeletConfigResponsePtrOutputWithContext(ctx context.Context) KubeletConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubeletConfigResponse) *KubeletConfigResponse {
		return &v
	}).(KubeletConfigResponsePtrOutput)
}

func (o KubeletConfigResponseOutput) AllowedUnsafeSysctls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v KubeletConfigResponse) []string { return v.AllowedUnsafeSysctls }).(pulumi.StringArrayOutput)
}

func (o KubeletConfigResponseOutput) ContainerLogMaxFiles() pulumi.IntPtrOutput {
	return o.ApplyT(func(v KubeletConfigResponse) *int { return v.ContainerLogMaxFiles }).(pulumi.IntPtrOutput)
}

func (o KubeletConfigResponseOutput) ContainerLogMaxSizeMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v KubeletConfigResponse) *int { return v.ContainerLogMaxSizeMB }).(pulumi.IntPtrOutput)
}

func (o KubeletConfigResponseOutput) CpuCfsQuota() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KubeletConfigResponse) *bool { return v.CpuCfsQuota }).(pulumi.BoolPtrOutput)
}

func (o KubeletConfigResponseOutput) CpuCfsQuotaPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubeletConfigResponse) *string { return v.CpuCfsQuotaPeriod }).(pulumi.StringPtrOutput)
}

func (o KubeletConfigResponseOutput) CpuManagerPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubeletConfigResponse) *string { return v.CpuManagerPolicy }).(pulumi.StringPtrOutput)
}

func (o KubeletConfigResponseOutput) FailSwapOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KubeletConfigResponse) *bool { return v.FailSwapOn }).(pulumi.BoolPtrOutput)
}

func (o KubeletConfigResponseOutput) ImageGcHighThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v KubeletConfigResponse) *int { return v.ImageGcHighThreshold }).(pulumi.IntPtrOutput)
}

func (o KubeletConfigResponseOutput) ImageGcLowThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v KubeletConfigResponse) *int { return v.ImageGcLowThreshold }).(pulumi.IntPtrOutput)
}

func (o KubeletConfigResponseOutput) PodMaxPids() pulumi.IntPtrOutput {
	return o.ApplyT(func(v KubeletConfigResponse) *int { return v.PodMaxPids }).(pulumi.IntPtrOutput)
}

func (o KubeletConfigResponseOutput) TopologyManagerPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubeletConfigResponse) *string { return v.TopologyManagerPolicy }).(pulumi.StringPtrOutput)
}

type KubeletConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (KubeletConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeletConfigResponse)(nil)).Elem()
}

func (o KubeletConfigResponsePtrOutput) ToKubeletConfigResponsePtrOutput() KubeletConfigResponsePtrOutput {
	return o
}

func (o KubeletConfigResponsePtrOutput) ToKubeletConfigResponsePtrOutputWithContext(ctx context.Context) KubeletConfigResponsePtrOutput {
	return o
}

func (o KubeletConfigResponsePtrOutput) Elem() KubeletConfigResponseOutput {
	return o.ApplyT(func(v *KubeletConfigResponse) KubeletConfigResponse {
		if v != nil {
			return *v
		}
		var ret KubeletConfigResponse
		return ret
	}).(KubeletConfigResponseOutput)
}

func (o KubeletConfigResponsePtrOutput) AllowedUnsafeSysctls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KubeletConfigResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedUnsafeSysctls
	}).(pulumi.StringArrayOutput)
}

func (o KubeletConfigResponsePtrOutput) ContainerLogMaxFiles() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KubeletConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.ContainerLogMaxFiles
	}).(pulumi.IntPtrOutput)
}

func (o KubeletConfigResponsePtrOutput) ContainerLogMaxSizeMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KubeletConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.ContainerLogMaxSizeMB
	}).(pulumi.IntPtrOutput)
}

func (o KubeletConfigResponsePtrOutput) CpuCfsQuota() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KubeletConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.CpuCfsQuota
	}).(pulumi.BoolPtrOutput)
}

func (o KubeletConfigResponsePtrOutput) CpuCfsQuotaPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeletConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.CpuCfsQuotaPeriod
	}).(pulumi.StringPtrOutput)
}

func (o KubeletConfigResponsePtrOutput) CpuManagerPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeletConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.CpuManagerPolicy
	}).(pulumi.StringPtrOutput)
}

func (o KubeletConfigResponsePtrOutput) FailSwapOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KubeletConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.FailSwapOn
	}).(pulumi.BoolPtrOutput)
}

func (o KubeletConfigResponsePtrOutput) ImageGcHighThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KubeletConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.ImageGcHighThreshold
	}).(pulumi.IntPtrOutput)
}

func (o KubeletConfigResponsePtrOutput) ImageGcLowThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KubeletConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.ImageGcLowThreshold
	}).(pulumi.IntPtrOutput)
}

func (o KubeletConfigResponsePtrOutput) PodMaxPids() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KubeletConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.PodMaxPids
	}).(pulumi.IntPtrOutput)
}

func (o KubeletConfigResponsePtrOutput) TopologyManagerPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeletConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.TopologyManagerPolicy
	}).(pulumi.StringPtrOutput)
}

type LinuxOSConfig struct {
	SwapFileSizeMB             *int          `pulumi:"swapFileSizeMB"`
	Sysctls                    *SysctlConfig `pulumi:"sysctls"`
	TransparentHugePageDefrag  *string       `pulumi:"transparentHugePageDefrag"`
	TransparentHugePageEnabled *string       `pulumi:"transparentHugePageEnabled"`
}





type LinuxOSConfigInput interface {
	pulumi.Input

	ToLinuxOSConfigOutput() LinuxOSConfigOutput
	ToLinuxOSConfigOutputWithContext(context.Context) LinuxOSConfigOutput
}

type LinuxOSConfigArgs struct {
	SwapFileSizeMB             pulumi.IntPtrInput    `pulumi:"swapFileSizeMB"`
	Sysctls                    SysctlConfigPtrInput  `pulumi:"sysctls"`
	TransparentHugePageDefrag  pulumi.StringPtrInput `pulumi:"transparentHugePageDefrag"`
	TransparentHugePageEnabled pulumi.StringPtrInput `pulumi:"transparentHugePageEnabled"`
}

func (LinuxOSConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxOSConfig)(nil)).Elem()
}

func (i LinuxOSConfigArgs) ToLinuxOSConfigOutput() LinuxOSConfigOutput {
	return i.ToLinuxOSConfigOutputWithContext(context.Background())
}

func (i LinuxOSConfigArgs) ToLinuxOSConfigOutputWithContext(ctx context.Context) LinuxOSConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxOSConfigOutput)
}

func (i LinuxOSConfigArgs) ToLinuxOSConfigPtrOutput() LinuxOSConfigPtrOutput {
	return i.ToLinuxOSConfigPtrOutputWithContext(context.Background())
}

func (i LinuxOSConfigArgs) ToLinuxOSConfigPtrOutputWithContext(ctx context.Context) LinuxOSConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxOSConfigOutput).ToLinuxOSConfigPtrOutputWithContext(ctx)
}









type LinuxOSConfigPtrInput interface {
	pulumi.Input

	ToLinuxOSConfigPtrOutput() LinuxOSConfigPtrOutput
	ToLinuxOSConfigPtrOutputWithContext(context.Context) LinuxOSConfigPtrOutput
}

type linuxOSConfigPtrType LinuxOSConfigArgs

func LinuxOSConfigPtr(v *LinuxOSConfigArgs) LinuxOSConfigPtrInput {
	return (*linuxOSConfigPtrType)(v)
}

func (*linuxOSConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxOSConfig)(nil)).Elem()
}

func (i *linuxOSConfigPtrType) ToLinuxOSConfigPtrOutput() LinuxOSConfigPtrOutput {
	return i.ToLinuxOSConfigPtrOutputWithContext(context.Background())
}

func (i *linuxOSConfigPtrType) ToLinuxOSConfigPtrOutputWithContext(ctx context.Context) LinuxOSConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxOSConfigPtrOutput)
}

type LinuxOSConfigOutput struct{ *pulumi.OutputState }

func (LinuxOSConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxOSConfig)(nil)).Elem()
}

func (o LinuxOSConfigOutput) ToLinuxOSConfigOutput() LinuxOSConfigOutput {
	return o
}

func (o LinuxOSConfigOutput) ToLinuxOSConfigOutputWithContext(ctx context.Context) LinuxOSConfigOutput {
	return o
}

func (o LinuxOSConfigOutput) ToLinuxOSConfigPtrOutput() LinuxOSConfigPtrOutput {
	return o.ToLinuxOSConfigPtrOutputWithContext(context.Background())
}

func (o LinuxOSConfigOutput) ToLinuxOSConfigPtrOutputWithContext(ctx context.Context) LinuxOSConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinuxOSConfig) *LinuxOSConfig {
		return &v
	}).(LinuxOSConfigPtrOutput)
}

func (o LinuxOSConfigOutput) SwapFileSizeMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LinuxOSConfig) *int { return v.SwapFileSizeMB }).(pulumi.IntPtrOutput)
}

func (o LinuxOSConfigOutput) Sysctls() SysctlConfigPtrOutput {
	return o.ApplyT(func(v LinuxOSConfig) *SysctlConfig { return v.Sysctls }).(SysctlConfigPtrOutput)
}

func (o LinuxOSConfigOutput) TransparentHugePageDefrag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxOSConfig) *string { return v.TransparentHugePageDefrag }).(pulumi.StringPtrOutput)
}

func (o LinuxOSConfigOutput) TransparentHugePageEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxOSConfig) *string { return v.TransparentHugePageEnabled }).(pulumi.StringPtrOutput)
}

type LinuxOSConfigPtrOutput struct{ *pulumi.OutputState }

func (LinuxOSConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxOSConfig)(nil)).Elem()
}

func (o LinuxOSConfigPtrOutput) ToLinuxOSConfigPtrOutput() LinuxOSConfigPtrOutput {
	return o
}

func (o LinuxOSConfigPtrOutput) ToLinuxOSConfigPtrOutputWithContext(ctx context.Context) LinuxOSConfigPtrOutput {
	return o
}

func (o LinuxOSConfigPtrOutput) Elem() LinuxOSConfigOutput {
	return o.ApplyT(func(v *LinuxOSConfig) LinuxOSConfig {
		if v != nil {
			return *v
		}
		var ret LinuxOSConfig
		return ret
	}).(LinuxOSConfigOutput)
}

func (o LinuxOSConfigPtrOutput) SwapFileSizeMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LinuxOSConfig) *int {
		if v == nil {
			return nil
		}
		return v.SwapFileSizeMB
	}).(pulumi.IntPtrOutput)
}

func (o LinuxOSConfigPtrOutput) Sysctls() SysctlConfigPtrOutput {
	return o.ApplyT(func(v *LinuxOSConfig) *SysctlConfig {
		if v == nil {
			return nil
		}
		return v.Sysctls
	}).(SysctlConfigPtrOutput)
}

func (o LinuxOSConfigPtrOutput) TransparentHugePageDefrag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxOSConfig) *string {
		if v == nil {
			return nil
		}
		return v.TransparentHugePageDefrag
	}).(pulumi.StringPtrOutput)
}

func (o LinuxOSConfigPtrOutput) TransparentHugePageEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxOSConfig) *string {
		if v == nil {
			return nil
		}
		return v.TransparentHugePageEnabled
	}).(pulumi.StringPtrOutput)
}

type LinuxOSConfigResponse struct {
	SwapFileSizeMB             *int                  `pulumi:"swapFileSizeMB"`
	Sysctls                    *SysctlConfigResponse `pulumi:"sysctls"`
	TransparentHugePageDefrag  *string               `pulumi:"transparentHugePageDefrag"`
	TransparentHugePageEnabled *string               `pulumi:"transparentHugePageEnabled"`
}





type LinuxOSConfigResponseInput interface {
	pulumi.Input

	ToLinuxOSConfigResponseOutput() LinuxOSConfigResponseOutput
	ToLinuxOSConfigResponseOutputWithContext(context.Context) LinuxOSConfigResponseOutput
}

type LinuxOSConfigResponseArgs struct {
	SwapFileSizeMB             pulumi.IntPtrInput           `pulumi:"swapFileSizeMB"`
	Sysctls                    SysctlConfigResponsePtrInput `pulumi:"sysctls"`
	TransparentHugePageDefrag  pulumi.StringPtrInput        `pulumi:"transparentHugePageDefrag"`
	TransparentHugePageEnabled pulumi.StringPtrInput        `pulumi:"transparentHugePageEnabled"`
}

func (LinuxOSConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxOSConfigResponse)(nil)).Elem()
}

func (i LinuxOSConfigResponseArgs) ToLinuxOSConfigResponseOutput() LinuxOSConfigResponseOutput {
	return i.ToLinuxOSConfigResponseOutputWithContext(context.Background())
}

func (i LinuxOSConfigResponseArgs) ToLinuxOSConfigResponseOutputWithContext(ctx context.Context) LinuxOSConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxOSConfigResponseOutput)
}

func (i LinuxOSConfigResponseArgs) ToLinuxOSConfigResponsePtrOutput() LinuxOSConfigResponsePtrOutput {
	return i.ToLinuxOSConfigResponsePtrOutputWithContext(context.Background())
}

func (i LinuxOSConfigResponseArgs) ToLinuxOSConfigResponsePtrOutputWithContext(ctx context.Context) LinuxOSConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxOSConfigResponseOutput).ToLinuxOSConfigResponsePtrOutputWithContext(ctx)
}









type LinuxOSConfigResponsePtrInput interface {
	pulumi.Input

	ToLinuxOSConfigResponsePtrOutput() LinuxOSConfigResponsePtrOutput
	ToLinuxOSConfigResponsePtrOutputWithContext(context.Context) LinuxOSConfigResponsePtrOutput
}

type linuxOSConfigResponsePtrType LinuxOSConfigResponseArgs

func LinuxOSConfigResponsePtr(v *LinuxOSConfigResponseArgs) LinuxOSConfigResponsePtrInput {
	return (*linuxOSConfigResponsePtrType)(v)
}

func (*linuxOSConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxOSConfigResponse)(nil)).Elem()
}

func (i *linuxOSConfigResponsePtrType) ToLinuxOSConfigResponsePtrOutput() LinuxOSConfigResponsePtrOutput {
	return i.ToLinuxOSConfigResponsePtrOutputWithContext(context.Background())
}

func (i *linuxOSConfigResponsePtrType) ToLinuxOSConfigResponsePtrOutputWithContext(ctx context.Context) LinuxOSConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxOSConfigResponsePtrOutput)
}

type LinuxOSConfigResponseOutput struct{ *pulumi.OutputState }

func (LinuxOSConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxOSConfigResponse)(nil)).Elem()
}

func (o LinuxOSConfigResponseOutput) ToLinuxOSConfigResponseOutput() LinuxOSConfigResponseOutput {
	return o
}

func (o LinuxOSConfigResponseOutput) ToLinuxOSConfigResponseOutputWithContext(ctx context.Context) LinuxOSConfigResponseOutput {
	return o
}

func (o LinuxOSConfigResponseOutput) ToLinuxOSConfigResponsePtrOutput() LinuxOSConfigResponsePtrOutput {
	return o.ToLinuxOSConfigResponsePtrOutputWithContext(context.Background())
}

func (o LinuxOSConfigResponseOutput) ToLinuxOSConfigResponsePtrOutputWithContext(ctx context.Context) LinuxOSConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinuxOSConfigResponse) *LinuxOSConfigResponse {
		return &v
	}).(LinuxOSConfigResponsePtrOutput)
}

func (o LinuxOSConfigResponseOutput) SwapFileSizeMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LinuxOSConfigResponse) *int { return v.SwapFileSizeMB }).(pulumi.IntPtrOutput)
}

func (o LinuxOSConfigResponseOutput) Sysctls() SysctlConfigResponsePtrOutput {
	return o.ApplyT(func(v LinuxOSConfigResponse) *SysctlConfigResponse { return v.Sysctls }).(SysctlConfigResponsePtrOutput)
}

func (o LinuxOSConfigResponseOutput) TransparentHugePageDefrag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxOSConfigResponse) *string { return v.TransparentHugePageDefrag }).(pulumi.StringPtrOutput)
}

func (o LinuxOSConfigResponseOutput) TransparentHugePageEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxOSConfigResponse) *string { return v.TransparentHugePageEnabled }).(pulumi.StringPtrOutput)
}

type LinuxOSConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (LinuxOSConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxOSConfigResponse)(nil)).Elem()
}

func (o LinuxOSConfigResponsePtrOutput) ToLinuxOSConfigResponsePtrOutput() LinuxOSConfigResponsePtrOutput {
	return o
}

func (o LinuxOSConfigResponsePtrOutput) ToLinuxOSConfigResponsePtrOutputWithContext(ctx context.Context) LinuxOSConfigResponsePtrOutput {
	return o
}

func (o LinuxOSConfigResponsePtrOutput) Elem() LinuxOSConfigResponseOutput {
	return o.ApplyT(func(v *LinuxOSConfigResponse) LinuxOSConfigResponse {
		if v != nil {
			return *v
		}
		var ret LinuxOSConfigResponse
		return ret
	}).(LinuxOSConfigResponseOutput)
}

func (o LinuxOSConfigResponsePtrOutput) SwapFileSizeMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LinuxOSConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.SwapFileSizeMB
	}).(pulumi.IntPtrOutput)
}

func (o LinuxOSConfigResponsePtrOutput) Sysctls() SysctlConfigResponsePtrOutput {
	return o.ApplyT(func(v *LinuxOSConfigResponse) *SysctlConfigResponse {
		if v == nil {
			return nil
		}
		return v.Sysctls
	}).(SysctlConfigResponsePtrOutput)
}

func (o LinuxOSConfigResponsePtrOutput) TransparentHugePageDefrag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxOSConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.TransparentHugePageDefrag
	}).(pulumi.StringPtrOutput)
}

func (o LinuxOSConfigResponsePtrOutput) TransparentHugePageEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxOSConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.TransparentHugePageEnabled
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterAADProfile struct {
	AdminGroupObjectIDs []string `pulumi:"adminGroupObjectIDs"`
	ClientAppID         *string  `pulumi:"clientAppID"`
	EnableAzureRBAC     *bool    `pulumi:"enableAzureRBAC"`
	Managed             *bool    `pulumi:"managed"`
	ServerAppID         *string  `pulumi:"serverAppID"`
	ServerAppSecret     *string  `pulumi:"serverAppSecret"`
	TenantID            *string  `pulumi:"tenantID"`
}





type ManagedClusterAADProfileInput interface {
	pulumi.Input

	ToManagedClusterAADProfileOutput() ManagedClusterAADProfileOutput
	ToManagedClusterAADProfileOutputWithContext(context.Context) ManagedClusterAADProfileOutput
}

type ManagedClusterAADProfileArgs struct {
	AdminGroupObjectIDs pulumi.StringArrayInput `pulumi:"adminGroupObjectIDs"`
	ClientAppID         pulumi.StringPtrInput   `pulumi:"clientAppID"`
	EnableAzureRBAC     pulumi.BoolPtrInput     `pulumi:"enableAzureRBAC"`
	Managed             pulumi.BoolPtrInput     `pulumi:"managed"`
	ServerAppID         pulumi.StringPtrInput   `pulumi:"serverAppID"`
	ServerAppSecret     pulumi.StringPtrInput   `pulumi:"serverAppSecret"`
	TenantID            pulumi.StringPtrInput   `pulumi:"tenantID"`
}

func (ManagedClusterAADProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAADProfile)(nil)).Elem()
}

func (i ManagedClusterAADProfileArgs) ToManagedClusterAADProfileOutput() ManagedClusterAADProfileOutput {
	return i.ToManagedClusterAADProfileOutputWithContext(context.Background())
}

func (i ManagedClusterAADProfileArgs) ToManagedClusterAADProfileOutputWithContext(ctx context.Context) ManagedClusterAADProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAADProfileOutput)
}

func (i ManagedClusterAADProfileArgs) ToManagedClusterAADProfilePtrOutput() ManagedClusterAADProfilePtrOutput {
	return i.ToManagedClusterAADProfilePtrOutputWithContext(context.Background())
}

func (i ManagedClusterAADProfileArgs) ToManagedClusterAADProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAADProfileOutput).ToManagedClusterAADProfilePtrOutputWithContext(ctx)
}









type ManagedClusterAADProfilePtrInput interface {
	pulumi.Input

	ToManagedClusterAADProfilePtrOutput() ManagedClusterAADProfilePtrOutput
	ToManagedClusterAADProfilePtrOutputWithContext(context.Context) ManagedClusterAADProfilePtrOutput
}

type managedClusterAADProfilePtrType ManagedClusterAADProfileArgs

func ManagedClusterAADProfilePtr(v *ManagedClusterAADProfileArgs) ManagedClusterAADProfilePtrInput {
	return (*managedClusterAADProfilePtrType)(v)
}

func (*managedClusterAADProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAADProfile)(nil)).Elem()
}

func (i *managedClusterAADProfilePtrType) ToManagedClusterAADProfilePtrOutput() ManagedClusterAADProfilePtrOutput {
	return i.ToManagedClusterAADProfilePtrOutputWithContext(context.Background())
}

func (i *managedClusterAADProfilePtrType) ToManagedClusterAADProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAADProfilePtrOutput)
}

type ManagedClusterAADProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterAADProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAADProfile)(nil)).Elem()
}

func (o ManagedClusterAADProfileOutput) ToManagedClusterAADProfileOutput() ManagedClusterAADProfileOutput {
	return o
}

func (o ManagedClusterAADProfileOutput) ToManagedClusterAADProfileOutputWithContext(ctx context.Context) ManagedClusterAADProfileOutput {
	return o
}

func (o ManagedClusterAADProfileOutput) ToManagedClusterAADProfilePtrOutput() ManagedClusterAADProfilePtrOutput {
	return o.ToManagedClusterAADProfilePtrOutputWithContext(context.Background())
}

func (o ManagedClusterAADProfileOutput) ToManagedClusterAADProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterAADProfile) *ManagedClusterAADProfile {
		return &v
	}).(ManagedClusterAADProfilePtrOutput)
}

func (o ManagedClusterAADProfileOutput) AdminGroupObjectIDs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedClusterAADProfile) []string { return v.AdminGroupObjectIDs }).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAADProfileOutput) ClientAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAADProfile) *string { return v.ClientAppID }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfileOutput) EnableAzureRBAC() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAADProfile) *bool { return v.EnableAzureRBAC }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAADProfileOutput) Managed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAADProfile) *bool { return v.Managed }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAADProfileOutput) ServerAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAADProfile) *string { return v.ServerAppID }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfileOutput) ServerAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAADProfile) *string { return v.ServerAppSecret }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfileOutput) TenantID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAADProfile) *string { return v.TenantID }).(pulumi.StringPtrOutput)
}

type ManagedClusterAADProfilePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterAADProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAADProfile)(nil)).Elem()
}

func (o ManagedClusterAADProfilePtrOutput) ToManagedClusterAADProfilePtrOutput() ManagedClusterAADProfilePtrOutput {
	return o
}

func (o ManagedClusterAADProfilePtrOutput) ToManagedClusterAADProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfilePtrOutput {
	return o
}

func (o ManagedClusterAADProfilePtrOutput) Elem() ManagedClusterAADProfileOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfile) ManagedClusterAADProfile {
		if v != nil {
			return *v
		}
		var ret ManagedClusterAADProfile
		return ret
	}).(ManagedClusterAADProfileOutput)
}

func (o ManagedClusterAADProfilePtrOutput) AdminGroupObjectIDs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfile) []string {
		if v == nil {
			return nil
		}
		return v.AdminGroupObjectIDs
	}).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAADProfilePtrOutput) ClientAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfile) *string {
		if v == nil {
			return nil
		}
		return v.ClientAppID
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfilePtrOutput) EnableAzureRBAC() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfile) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAzureRBAC
	}).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAADProfilePtrOutput) Managed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfile) *bool {
		if v == nil {
			return nil
		}
		return v.Managed
	}).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAADProfilePtrOutput) ServerAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfile) *string {
		if v == nil {
			return nil
		}
		return v.ServerAppID
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfilePtrOutput) ServerAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfile) *string {
		if v == nil {
			return nil
		}
		return v.ServerAppSecret
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfilePtrOutput) TenantID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfile) *string {
		if v == nil {
			return nil
		}
		return v.TenantID
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterAADProfileResponse struct {
	AdminGroupObjectIDs []string `pulumi:"adminGroupObjectIDs"`
	ClientAppID         *string  `pulumi:"clientAppID"`
	EnableAzureRBAC     *bool    `pulumi:"enableAzureRBAC"`
	Managed             *bool    `pulumi:"managed"`
	ServerAppID         *string  `pulumi:"serverAppID"`
	ServerAppSecret     *string  `pulumi:"serverAppSecret"`
	TenantID            *string  `pulumi:"tenantID"`
}





type ManagedClusterAADProfileResponseInput interface {
	pulumi.Input

	ToManagedClusterAADProfileResponseOutput() ManagedClusterAADProfileResponseOutput
	ToManagedClusterAADProfileResponseOutputWithContext(context.Context) ManagedClusterAADProfileResponseOutput
}

type ManagedClusterAADProfileResponseArgs struct {
	AdminGroupObjectIDs pulumi.StringArrayInput `pulumi:"adminGroupObjectIDs"`
	ClientAppID         pulumi.StringPtrInput   `pulumi:"clientAppID"`
	EnableAzureRBAC     pulumi.BoolPtrInput     `pulumi:"enableAzureRBAC"`
	Managed             pulumi.BoolPtrInput     `pulumi:"managed"`
	ServerAppID         pulumi.StringPtrInput   `pulumi:"serverAppID"`
	ServerAppSecret     pulumi.StringPtrInput   `pulumi:"serverAppSecret"`
	TenantID            pulumi.StringPtrInput   `pulumi:"tenantID"`
}

func (ManagedClusterAADProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAADProfileResponse)(nil)).Elem()
}

func (i ManagedClusterAADProfileResponseArgs) ToManagedClusterAADProfileResponseOutput() ManagedClusterAADProfileResponseOutput {
	return i.ToManagedClusterAADProfileResponseOutputWithContext(context.Background())
}

func (i ManagedClusterAADProfileResponseArgs) ToManagedClusterAADProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAADProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAADProfileResponseOutput)
}

func (i ManagedClusterAADProfileResponseArgs) ToManagedClusterAADProfileResponsePtrOutput() ManagedClusterAADProfileResponsePtrOutput {
	return i.ToManagedClusterAADProfileResponsePtrOutputWithContext(context.Background())
}

func (i ManagedClusterAADProfileResponseArgs) ToManagedClusterAADProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAADProfileResponseOutput).ToManagedClusterAADProfileResponsePtrOutputWithContext(ctx)
}









type ManagedClusterAADProfileResponsePtrInput interface {
	pulumi.Input

	ToManagedClusterAADProfileResponsePtrOutput() ManagedClusterAADProfileResponsePtrOutput
	ToManagedClusterAADProfileResponsePtrOutputWithContext(context.Context) ManagedClusterAADProfileResponsePtrOutput
}

type managedClusterAADProfileResponsePtrType ManagedClusterAADProfileResponseArgs

func ManagedClusterAADProfileResponsePtr(v *ManagedClusterAADProfileResponseArgs) ManagedClusterAADProfileResponsePtrInput {
	return (*managedClusterAADProfileResponsePtrType)(v)
}

func (*managedClusterAADProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAADProfileResponse)(nil)).Elem()
}

func (i *managedClusterAADProfileResponsePtrType) ToManagedClusterAADProfileResponsePtrOutput() ManagedClusterAADProfileResponsePtrOutput {
	return i.ToManagedClusterAADProfileResponsePtrOutputWithContext(context.Background())
}

func (i *managedClusterAADProfileResponsePtrType) ToManagedClusterAADProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAADProfileResponsePtrOutput)
}

type ManagedClusterAADProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterAADProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAADProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAADProfileResponseOutput) ToManagedClusterAADProfileResponseOutput() ManagedClusterAADProfileResponseOutput {
	return o
}

func (o ManagedClusterAADProfileResponseOutput) ToManagedClusterAADProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAADProfileResponseOutput {
	return o
}

func (o ManagedClusterAADProfileResponseOutput) ToManagedClusterAADProfileResponsePtrOutput() ManagedClusterAADProfileResponsePtrOutput {
	return o.ToManagedClusterAADProfileResponsePtrOutputWithContext(context.Background())
}

func (o ManagedClusterAADProfileResponseOutput) ToManagedClusterAADProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterAADProfileResponse) *ManagedClusterAADProfileResponse {
		return &v
	}).(ManagedClusterAADProfileResponsePtrOutput)
}

func (o ManagedClusterAADProfileResponseOutput) AdminGroupObjectIDs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedClusterAADProfileResponse) []string { return v.AdminGroupObjectIDs }).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAADProfileResponseOutput) ClientAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAADProfileResponse) *string { return v.ClientAppID }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfileResponseOutput) EnableAzureRBAC() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAADProfileResponse) *bool { return v.EnableAzureRBAC }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAADProfileResponseOutput) Managed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAADProfileResponse) *bool { return v.Managed }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAADProfileResponseOutput) ServerAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAADProfileResponse) *string { return v.ServerAppID }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfileResponseOutput) ServerAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAADProfileResponse) *string { return v.ServerAppSecret }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfileResponseOutput) TenantID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAADProfileResponse) *string { return v.TenantID }).(pulumi.StringPtrOutput)
}

type ManagedClusterAADProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterAADProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAADProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAADProfileResponsePtrOutput) ToManagedClusterAADProfileResponsePtrOutput() ManagedClusterAADProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterAADProfileResponsePtrOutput) ToManagedClusterAADProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterAADProfileResponsePtrOutput) Elem() ManagedClusterAADProfileResponseOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfileResponse) ManagedClusterAADProfileResponse {
		if v != nil {
			return *v
		}
		var ret ManagedClusterAADProfileResponse
		return ret
	}).(ManagedClusterAADProfileResponseOutput)
}

func (o ManagedClusterAADProfileResponsePtrOutput) AdminGroupObjectIDs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfileResponse) []string {
		if v == nil {
			return nil
		}
		return v.AdminGroupObjectIDs
	}).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAADProfileResponsePtrOutput) ClientAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientAppID
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfileResponsePtrOutput) EnableAzureRBAC() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfileResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAzureRBAC
	}).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAADProfileResponsePtrOutput) Managed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfileResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Managed
	}).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAADProfileResponsePtrOutput) ServerAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServerAppID
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfileResponsePtrOutput) ServerAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServerAppSecret
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfileResponsePtrOutput) TenantID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantID
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterAPIServerAccessProfile struct {
	AuthorizedIPRanges   []string `pulumi:"authorizedIPRanges"`
	EnablePrivateCluster *bool    `pulumi:"enablePrivateCluster"`
	PrivateDNSZone       *string  `pulumi:"privateDNSZone"`
}





type ManagedClusterAPIServerAccessProfileInput interface {
	pulumi.Input

	ToManagedClusterAPIServerAccessProfileOutput() ManagedClusterAPIServerAccessProfileOutput
	ToManagedClusterAPIServerAccessProfileOutputWithContext(context.Context) ManagedClusterAPIServerAccessProfileOutput
}

type ManagedClusterAPIServerAccessProfileArgs struct {
	AuthorizedIPRanges   pulumi.StringArrayInput `pulumi:"authorizedIPRanges"`
	EnablePrivateCluster pulumi.BoolPtrInput     `pulumi:"enablePrivateCluster"`
	PrivateDNSZone       pulumi.StringPtrInput   `pulumi:"privateDNSZone"`
}

func (ManagedClusterAPIServerAccessProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAPIServerAccessProfile)(nil)).Elem()
}

func (i ManagedClusterAPIServerAccessProfileArgs) ToManagedClusterAPIServerAccessProfileOutput() ManagedClusterAPIServerAccessProfileOutput {
	return i.ToManagedClusterAPIServerAccessProfileOutputWithContext(context.Background())
}

func (i ManagedClusterAPIServerAccessProfileArgs) ToManagedClusterAPIServerAccessProfileOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAPIServerAccessProfileOutput)
}

func (i ManagedClusterAPIServerAccessProfileArgs) ToManagedClusterAPIServerAccessProfilePtrOutput() ManagedClusterAPIServerAccessProfilePtrOutput {
	return i.ToManagedClusterAPIServerAccessProfilePtrOutputWithContext(context.Background())
}

func (i ManagedClusterAPIServerAccessProfileArgs) ToManagedClusterAPIServerAccessProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAPIServerAccessProfileOutput).ToManagedClusterAPIServerAccessProfilePtrOutputWithContext(ctx)
}









type ManagedClusterAPIServerAccessProfilePtrInput interface {
	pulumi.Input

	ToManagedClusterAPIServerAccessProfilePtrOutput() ManagedClusterAPIServerAccessProfilePtrOutput
	ToManagedClusterAPIServerAccessProfilePtrOutputWithContext(context.Context) ManagedClusterAPIServerAccessProfilePtrOutput
}

type managedClusterAPIServerAccessProfilePtrType ManagedClusterAPIServerAccessProfileArgs

func ManagedClusterAPIServerAccessProfilePtr(v *ManagedClusterAPIServerAccessProfileArgs) ManagedClusterAPIServerAccessProfilePtrInput {
	return (*managedClusterAPIServerAccessProfilePtrType)(v)
}

func (*managedClusterAPIServerAccessProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAPIServerAccessProfile)(nil)).Elem()
}

func (i *managedClusterAPIServerAccessProfilePtrType) ToManagedClusterAPIServerAccessProfilePtrOutput() ManagedClusterAPIServerAccessProfilePtrOutput {
	return i.ToManagedClusterAPIServerAccessProfilePtrOutputWithContext(context.Background())
}

func (i *managedClusterAPIServerAccessProfilePtrType) ToManagedClusterAPIServerAccessProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAPIServerAccessProfilePtrOutput)
}

type ManagedClusterAPIServerAccessProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterAPIServerAccessProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAPIServerAccessProfile)(nil)).Elem()
}

func (o ManagedClusterAPIServerAccessProfileOutput) ToManagedClusterAPIServerAccessProfileOutput() ManagedClusterAPIServerAccessProfileOutput {
	return o
}

func (o ManagedClusterAPIServerAccessProfileOutput) ToManagedClusterAPIServerAccessProfileOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfileOutput {
	return o
}

func (o ManagedClusterAPIServerAccessProfileOutput) ToManagedClusterAPIServerAccessProfilePtrOutput() ManagedClusterAPIServerAccessProfilePtrOutput {
	return o.ToManagedClusterAPIServerAccessProfilePtrOutputWithContext(context.Background())
}

func (o ManagedClusterAPIServerAccessProfileOutput) ToManagedClusterAPIServerAccessProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterAPIServerAccessProfile) *ManagedClusterAPIServerAccessProfile {
		return &v
	}).(ManagedClusterAPIServerAccessProfilePtrOutput)
}

func (o ManagedClusterAPIServerAccessProfileOutput) AuthorizedIPRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedClusterAPIServerAccessProfile) []string { return v.AuthorizedIPRanges }).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAPIServerAccessProfileOutput) EnablePrivateCluster() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAPIServerAccessProfile) *bool { return v.EnablePrivateCluster }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAPIServerAccessProfileOutput) PrivateDNSZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAPIServerAccessProfile) *string { return v.PrivateDNSZone }).(pulumi.StringPtrOutput)
}

type ManagedClusterAPIServerAccessProfilePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterAPIServerAccessProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAPIServerAccessProfile)(nil)).Elem()
}

func (o ManagedClusterAPIServerAccessProfilePtrOutput) ToManagedClusterAPIServerAccessProfilePtrOutput() ManagedClusterAPIServerAccessProfilePtrOutput {
	return o
}

func (o ManagedClusterAPIServerAccessProfilePtrOutput) ToManagedClusterAPIServerAccessProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfilePtrOutput {
	return o
}

func (o ManagedClusterAPIServerAccessProfilePtrOutput) Elem() ManagedClusterAPIServerAccessProfileOutput {
	return o.ApplyT(func(v *ManagedClusterAPIServerAccessProfile) ManagedClusterAPIServerAccessProfile {
		if v != nil {
			return *v
		}
		var ret ManagedClusterAPIServerAccessProfile
		return ret
	}).(ManagedClusterAPIServerAccessProfileOutput)
}

func (o ManagedClusterAPIServerAccessProfilePtrOutput) AuthorizedIPRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedClusterAPIServerAccessProfile) []string {
		if v == nil {
			return nil
		}
		return v.AuthorizedIPRanges
	}).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAPIServerAccessProfilePtrOutput) EnablePrivateCluster() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAPIServerAccessProfile) *bool {
		if v == nil {
			return nil
		}
		return v.EnablePrivateCluster
	}).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAPIServerAccessProfilePtrOutput) PrivateDNSZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAPIServerAccessProfile) *string {
		if v == nil {
			return nil
		}
		return v.PrivateDNSZone
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterAPIServerAccessProfileResponse struct {
	AuthorizedIPRanges   []string `pulumi:"authorizedIPRanges"`
	EnablePrivateCluster *bool    `pulumi:"enablePrivateCluster"`
	PrivateDNSZone       *string  `pulumi:"privateDNSZone"`
}





type ManagedClusterAPIServerAccessProfileResponseInput interface {
	pulumi.Input

	ToManagedClusterAPIServerAccessProfileResponseOutput() ManagedClusterAPIServerAccessProfileResponseOutput
	ToManagedClusterAPIServerAccessProfileResponseOutputWithContext(context.Context) ManagedClusterAPIServerAccessProfileResponseOutput
}

type ManagedClusterAPIServerAccessProfileResponseArgs struct {
	AuthorizedIPRanges   pulumi.StringArrayInput `pulumi:"authorizedIPRanges"`
	EnablePrivateCluster pulumi.BoolPtrInput     `pulumi:"enablePrivateCluster"`
	PrivateDNSZone       pulumi.StringPtrInput   `pulumi:"privateDNSZone"`
}

func (ManagedClusterAPIServerAccessProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAPIServerAccessProfileResponse)(nil)).Elem()
}

func (i ManagedClusterAPIServerAccessProfileResponseArgs) ToManagedClusterAPIServerAccessProfileResponseOutput() ManagedClusterAPIServerAccessProfileResponseOutput {
	return i.ToManagedClusterAPIServerAccessProfileResponseOutputWithContext(context.Background())
}

func (i ManagedClusterAPIServerAccessProfileResponseArgs) ToManagedClusterAPIServerAccessProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAPIServerAccessProfileResponseOutput)
}

func (i ManagedClusterAPIServerAccessProfileResponseArgs) ToManagedClusterAPIServerAccessProfileResponsePtrOutput() ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return i.ToManagedClusterAPIServerAccessProfileResponsePtrOutputWithContext(context.Background())
}

func (i ManagedClusterAPIServerAccessProfileResponseArgs) ToManagedClusterAPIServerAccessProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAPIServerAccessProfileResponseOutput).ToManagedClusterAPIServerAccessProfileResponsePtrOutputWithContext(ctx)
}









type ManagedClusterAPIServerAccessProfileResponsePtrInput interface {
	pulumi.Input

	ToManagedClusterAPIServerAccessProfileResponsePtrOutput() ManagedClusterAPIServerAccessProfileResponsePtrOutput
	ToManagedClusterAPIServerAccessProfileResponsePtrOutputWithContext(context.Context) ManagedClusterAPIServerAccessProfileResponsePtrOutput
}

type managedClusterAPIServerAccessProfileResponsePtrType ManagedClusterAPIServerAccessProfileResponseArgs

func ManagedClusterAPIServerAccessProfileResponsePtr(v *ManagedClusterAPIServerAccessProfileResponseArgs) ManagedClusterAPIServerAccessProfileResponsePtrInput {
	return (*managedClusterAPIServerAccessProfileResponsePtrType)(v)
}

func (*managedClusterAPIServerAccessProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAPIServerAccessProfileResponse)(nil)).Elem()
}

func (i *managedClusterAPIServerAccessProfileResponsePtrType) ToManagedClusterAPIServerAccessProfileResponsePtrOutput() ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return i.ToManagedClusterAPIServerAccessProfileResponsePtrOutputWithContext(context.Background())
}

func (i *managedClusterAPIServerAccessProfileResponsePtrType) ToManagedClusterAPIServerAccessProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAPIServerAccessProfileResponsePtrOutput)
}

type ManagedClusterAPIServerAccessProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterAPIServerAccessProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAPIServerAccessProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAPIServerAccessProfileResponseOutput) ToManagedClusterAPIServerAccessProfileResponseOutput() ManagedClusterAPIServerAccessProfileResponseOutput {
	return o
}

func (o ManagedClusterAPIServerAccessProfileResponseOutput) ToManagedClusterAPIServerAccessProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfileResponseOutput {
	return o
}

func (o ManagedClusterAPIServerAccessProfileResponseOutput) ToManagedClusterAPIServerAccessProfileResponsePtrOutput() ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return o.ToManagedClusterAPIServerAccessProfileResponsePtrOutputWithContext(context.Background())
}

func (o ManagedClusterAPIServerAccessProfileResponseOutput) ToManagedClusterAPIServerAccessProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterAPIServerAccessProfileResponse) *ManagedClusterAPIServerAccessProfileResponse {
		return &v
	}).(ManagedClusterAPIServerAccessProfileResponsePtrOutput)
}

func (o ManagedClusterAPIServerAccessProfileResponseOutput) AuthorizedIPRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedClusterAPIServerAccessProfileResponse) []string { return v.AuthorizedIPRanges }).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAPIServerAccessProfileResponseOutput) EnablePrivateCluster() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAPIServerAccessProfileResponse) *bool { return v.EnablePrivateCluster }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAPIServerAccessProfileResponseOutput) PrivateDNSZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAPIServerAccessProfileResponse) *string { return v.PrivateDNSZone }).(pulumi.StringPtrOutput)
}

type ManagedClusterAPIServerAccessProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterAPIServerAccessProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAPIServerAccessProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAPIServerAccessProfileResponsePtrOutput) ToManagedClusterAPIServerAccessProfileResponsePtrOutput() ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterAPIServerAccessProfileResponsePtrOutput) ToManagedClusterAPIServerAccessProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterAPIServerAccessProfileResponsePtrOutput) Elem() ManagedClusterAPIServerAccessProfileResponseOutput {
	return o.ApplyT(func(v *ManagedClusterAPIServerAccessProfileResponse) ManagedClusterAPIServerAccessProfileResponse {
		if v != nil {
			return *v
		}
		var ret ManagedClusterAPIServerAccessProfileResponse
		return ret
	}).(ManagedClusterAPIServerAccessProfileResponseOutput)
}

func (o ManagedClusterAPIServerAccessProfileResponsePtrOutput) AuthorizedIPRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedClusterAPIServerAccessProfileResponse) []string {
		if v == nil {
			return nil
		}
		return v.AuthorizedIPRanges
	}).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAPIServerAccessProfileResponsePtrOutput) EnablePrivateCluster() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAPIServerAccessProfileResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnablePrivateCluster
	}).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAPIServerAccessProfileResponsePtrOutput) PrivateDNSZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAPIServerAccessProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrivateDNSZone
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterAddonProfile struct {
	Config  map[string]string `pulumi:"config"`
	Enabled bool              `pulumi:"enabled"`
}





type ManagedClusterAddonProfileInput interface {
	pulumi.Input

	ToManagedClusterAddonProfileOutput() ManagedClusterAddonProfileOutput
	ToManagedClusterAddonProfileOutputWithContext(context.Context) ManagedClusterAddonProfileOutput
}

type ManagedClusterAddonProfileArgs struct {
	Config  pulumi.StringMapInput `pulumi:"config"`
	Enabled pulumi.BoolInput      `pulumi:"enabled"`
}

func (ManagedClusterAddonProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAddonProfile)(nil)).Elem()
}

func (i ManagedClusterAddonProfileArgs) ToManagedClusterAddonProfileOutput() ManagedClusterAddonProfileOutput {
	return i.ToManagedClusterAddonProfileOutputWithContext(context.Background())
}

func (i ManagedClusterAddonProfileArgs) ToManagedClusterAddonProfileOutputWithContext(ctx context.Context) ManagedClusterAddonProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAddonProfileOutput)
}





type ManagedClusterAddonProfileMapInput interface {
	pulumi.Input

	ToManagedClusterAddonProfileMapOutput() ManagedClusterAddonProfileMapOutput
	ToManagedClusterAddonProfileMapOutputWithContext(context.Context) ManagedClusterAddonProfileMapOutput
}

type ManagedClusterAddonProfileMap map[string]ManagedClusterAddonProfileInput

func (ManagedClusterAddonProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterAddonProfile)(nil)).Elem()
}

func (i ManagedClusterAddonProfileMap) ToManagedClusterAddonProfileMapOutput() ManagedClusterAddonProfileMapOutput {
	return i.ToManagedClusterAddonProfileMapOutputWithContext(context.Background())
}

func (i ManagedClusterAddonProfileMap) ToManagedClusterAddonProfileMapOutputWithContext(ctx context.Context) ManagedClusterAddonProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAddonProfileMapOutput)
}

type ManagedClusterAddonProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterAddonProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAddonProfile)(nil)).Elem()
}

func (o ManagedClusterAddonProfileOutput) ToManagedClusterAddonProfileOutput() ManagedClusterAddonProfileOutput {
	return o
}

func (o ManagedClusterAddonProfileOutput) ToManagedClusterAddonProfileOutputWithContext(ctx context.Context) ManagedClusterAddonProfileOutput {
	return o
}

func (o ManagedClusterAddonProfileOutput) Config() pulumi.StringMapOutput {
	return o.ApplyT(func(v ManagedClusterAddonProfile) map[string]string { return v.Config }).(pulumi.StringMapOutput)
}

func (o ManagedClusterAddonProfileOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ManagedClusterAddonProfile) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type ManagedClusterAddonProfileMapOutput struct{ *pulumi.OutputState }

func (ManagedClusterAddonProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterAddonProfile)(nil)).Elem()
}

func (o ManagedClusterAddonProfileMapOutput) ToManagedClusterAddonProfileMapOutput() ManagedClusterAddonProfileMapOutput {
	return o
}

func (o ManagedClusterAddonProfileMapOutput) ToManagedClusterAddonProfileMapOutputWithContext(ctx context.Context) ManagedClusterAddonProfileMapOutput {
	return o
}

func (o ManagedClusterAddonProfileMapOutput) MapIndex(k pulumi.StringInput) ManagedClusterAddonProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ManagedClusterAddonProfile {
		return vs[0].(map[string]ManagedClusterAddonProfile)[vs[1].(string)]
	}).(ManagedClusterAddonProfileOutput)
}

type ManagedClusterAddonProfileResponse struct {
	Config   map[string]string                          `pulumi:"config"`
	Enabled  bool                                       `pulumi:"enabled"`
	Identity ManagedClusterAddonProfileResponseIdentity `pulumi:"identity"`
}





type ManagedClusterAddonProfileResponseInput interface {
	pulumi.Input

	ToManagedClusterAddonProfileResponseOutput() ManagedClusterAddonProfileResponseOutput
	ToManagedClusterAddonProfileResponseOutputWithContext(context.Context) ManagedClusterAddonProfileResponseOutput
}

type ManagedClusterAddonProfileResponseArgs struct {
	Config   pulumi.StringMapInput                           `pulumi:"config"`
	Enabled  pulumi.BoolInput                                `pulumi:"enabled"`
	Identity ManagedClusterAddonProfileResponseIdentityInput `pulumi:"identity"`
}

func (ManagedClusterAddonProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAddonProfileResponse)(nil)).Elem()
}

func (i ManagedClusterAddonProfileResponseArgs) ToManagedClusterAddonProfileResponseOutput() ManagedClusterAddonProfileResponseOutput {
	return i.ToManagedClusterAddonProfileResponseOutputWithContext(context.Background())
}

func (i ManagedClusterAddonProfileResponseArgs) ToManagedClusterAddonProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAddonProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAddonProfileResponseOutput)
}





type ManagedClusterAddonProfileResponseMapInput interface {
	pulumi.Input

	ToManagedClusterAddonProfileResponseMapOutput() ManagedClusterAddonProfileResponseMapOutput
	ToManagedClusterAddonProfileResponseMapOutputWithContext(context.Context) ManagedClusterAddonProfileResponseMapOutput
}

type ManagedClusterAddonProfileResponseMap map[string]ManagedClusterAddonProfileResponseInput

func (ManagedClusterAddonProfileResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterAddonProfileResponse)(nil)).Elem()
}

func (i ManagedClusterAddonProfileResponseMap) ToManagedClusterAddonProfileResponseMapOutput() ManagedClusterAddonProfileResponseMapOutput {
	return i.ToManagedClusterAddonProfileResponseMapOutputWithContext(context.Background())
}

func (i ManagedClusterAddonProfileResponseMap) ToManagedClusterAddonProfileResponseMapOutputWithContext(ctx context.Context) ManagedClusterAddonProfileResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAddonProfileResponseMapOutput)
}

type ManagedClusterAddonProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterAddonProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAddonProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAddonProfileResponseOutput) ToManagedClusterAddonProfileResponseOutput() ManagedClusterAddonProfileResponseOutput {
	return o
}

func (o ManagedClusterAddonProfileResponseOutput) ToManagedClusterAddonProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAddonProfileResponseOutput {
	return o
}

func (o ManagedClusterAddonProfileResponseOutput) Config() pulumi.StringMapOutput {
	return o.ApplyT(func(v ManagedClusterAddonProfileResponse) map[string]string { return v.Config }).(pulumi.StringMapOutput)
}

func (o ManagedClusterAddonProfileResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ManagedClusterAddonProfileResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o ManagedClusterAddonProfileResponseOutput) Identity() ManagedClusterAddonProfileResponseIdentityOutput {
	return o.ApplyT(func(v ManagedClusterAddonProfileResponse) ManagedClusterAddonProfileResponseIdentity {
		return v.Identity
	}).(ManagedClusterAddonProfileResponseIdentityOutput)
}

type ManagedClusterAddonProfileResponseMapOutput struct{ *pulumi.OutputState }

func (ManagedClusterAddonProfileResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterAddonProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAddonProfileResponseMapOutput) ToManagedClusterAddonProfileResponseMapOutput() ManagedClusterAddonProfileResponseMapOutput {
	return o
}

func (o ManagedClusterAddonProfileResponseMapOutput) ToManagedClusterAddonProfileResponseMapOutputWithContext(ctx context.Context) ManagedClusterAddonProfileResponseMapOutput {
	return o
}

func (o ManagedClusterAddonProfileResponseMapOutput) MapIndex(k pulumi.StringInput) ManagedClusterAddonProfileResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ManagedClusterAddonProfileResponse {
		return vs[0].(map[string]ManagedClusterAddonProfileResponse)[vs[1].(string)]
	}).(ManagedClusterAddonProfileResponseOutput)
}

type ManagedClusterAddonProfileResponseIdentity struct {
	ClientId   *string `pulumi:"clientId"`
	ObjectId   *string `pulumi:"objectId"`
	ResourceId *string `pulumi:"resourceId"`
}





type ManagedClusterAddonProfileResponseIdentityInput interface {
	pulumi.Input

	ToManagedClusterAddonProfileResponseIdentityOutput() ManagedClusterAddonProfileResponseIdentityOutput
	ToManagedClusterAddonProfileResponseIdentityOutputWithContext(context.Context) ManagedClusterAddonProfileResponseIdentityOutput
}

type ManagedClusterAddonProfileResponseIdentityArgs struct {
	ClientId   pulumi.StringPtrInput `pulumi:"clientId"`
	ObjectId   pulumi.StringPtrInput `pulumi:"objectId"`
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (ManagedClusterAddonProfileResponseIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAddonProfileResponseIdentity)(nil)).Elem()
}

func (i ManagedClusterAddonProfileResponseIdentityArgs) ToManagedClusterAddonProfileResponseIdentityOutput() ManagedClusterAddonProfileResponseIdentityOutput {
	return i.ToManagedClusterAddonProfileResponseIdentityOutputWithContext(context.Background())
}

func (i ManagedClusterAddonProfileResponseIdentityArgs) ToManagedClusterAddonProfileResponseIdentityOutputWithContext(ctx context.Context) ManagedClusterAddonProfileResponseIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAddonProfileResponseIdentityOutput)
}

type ManagedClusterAddonProfileResponseIdentityOutput struct{ *pulumi.OutputState }

func (ManagedClusterAddonProfileResponseIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAddonProfileResponseIdentity)(nil)).Elem()
}

func (o ManagedClusterAddonProfileResponseIdentityOutput) ToManagedClusterAddonProfileResponseIdentityOutput() ManagedClusterAddonProfileResponseIdentityOutput {
	return o
}

func (o ManagedClusterAddonProfileResponseIdentityOutput) ToManagedClusterAddonProfileResponseIdentityOutputWithContext(ctx context.Context) ManagedClusterAddonProfileResponseIdentityOutput {
	return o
}

func (o ManagedClusterAddonProfileResponseIdentityOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAddonProfileResponseIdentity) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAddonProfileResponseIdentityOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAddonProfileResponseIdentity) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAddonProfileResponseIdentityOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAddonProfileResponseIdentity) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type ManagedClusterAgentPoolProfile struct {
	AvailabilityZones         []string                  `pulumi:"availabilityZones"`
	Count                     *int                      `pulumi:"count"`
	EnableAutoScaling         *bool                     `pulumi:"enableAutoScaling"`
	EnableEncryptionAtHost    *bool                     `pulumi:"enableEncryptionAtHost"`
	EnableFIPS                *bool                     `pulumi:"enableFIPS"`
	EnableNodePublicIP        *bool                     `pulumi:"enableNodePublicIP"`
	GpuInstanceProfile        *string                   `pulumi:"gpuInstanceProfile"`
	KubeletConfig             *KubeletConfig            `pulumi:"kubeletConfig"`
	KubeletDiskType           *string                   `pulumi:"kubeletDiskType"`
	LinuxOSConfig             *LinuxOSConfig            `pulumi:"linuxOSConfig"`
	MaxCount                  *int                      `pulumi:"maxCount"`
	MaxPods                   *int                      `pulumi:"maxPods"`
	MinCount                  *int                      `pulumi:"minCount"`
	Mode                      *string                   `pulumi:"mode"`
	Name                      string                    `pulumi:"name"`
	NodeLabels                map[string]string         `pulumi:"nodeLabels"`
	NodePublicIPPrefixID      *string                   `pulumi:"nodePublicIPPrefixID"`
	NodeTaints                []string                  `pulumi:"nodeTaints"`
	OrchestratorVersion       *string                   `pulumi:"orchestratorVersion"`
	OsDiskSizeGB              *int                      `pulumi:"osDiskSizeGB"`
	OsDiskType                *string                   `pulumi:"osDiskType"`
	OsSKU                     *string                   `pulumi:"osSKU"`
	OsType                    *string                   `pulumi:"osType"`
	PodSubnetID               *string                   `pulumi:"podSubnetID"`
	ProximityPlacementGroupID *string                   `pulumi:"proximityPlacementGroupID"`
	ScaleSetEvictionPolicy    *string                   `pulumi:"scaleSetEvictionPolicy"`
	ScaleSetPriority          *string                   `pulumi:"scaleSetPriority"`
	SpotMaxPrice              *float64                  `pulumi:"spotMaxPrice"`
	Tags                      map[string]string         `pulumi:"tags"`
	Type                      *string                   `pulumi:"type"`
	UpgradeSettings           *AgentPoolUpgradeSettings `pulumi:"upgradeSettings"`
	VmSize                    *string                   `pulumi:"vmSize"`
	VnetSubnetID              *string                   `pulumi:"vnetSubnetID"`
}





type ManagedClusterAgentPoolProfileInput interface {
	pulumi.Input

	ToManagedClusterAgentPoolProfileOutput() ManagedClusterAgentPoolProfileOutput
	ToManagedClusterAgentPoolProfileOutputWithContext(context.Context) ManagedClusterAgentPoolProfileOutput
}

type ManagedClusterAgentPoolProfileArgs struct {
	AvailabilityZones         pulumi.StringArrayInput          `pulumi:"availabilityZones"`
	Count                     pulumi.IntPtrInput               `pulumi:"count"`
	EnableAutoScaling         pulumi.BoolPtrInput              `pulumi:"enableAutoScaling"`
	EnableEncryptionAtHost    pulumi.BoolPtrInput              `pulumi:"enableEncryptionAtHost"`
	EnableFIPS                pulumi.BoolPtrInput              `pulumi:"enableFIPS"`
	EnableNodePublicIP        pulumi.BoolPtrInput              `pulumi:"enableNodePublicIP"`
	GpuInstanceProfile        pulumi.StringPtrInput            `pulumi:"gpuInstanceProfile"`
	KubeletConfig             KubeletConfigPtrInput            `pulumi:"kubeletConfig"`
	KubeletDiskType           pulumi.StringPtrInput            `pulumi:"kubeletDiskType"`
	LinuxOSConfig             LinuxOSConfigPtrInput            `pulumi:"linuxOSConfig"`
	MaxCount                  pulumi.IntPtrInput               `pulumi:"maxCount"`
	MaxPods                   pulumi.IntPtrInput               `pulumi:"maxPods"`
	MinCount                  pulumi.IntPtrInput               `pulumi:"minCount"`
	Mode                      pulumi.StringPtrInput            `pulumi:"mode"`
	Name                      pulumi.StringInput               `pulumi:"name"`
	NodeLabels                pulumi.StringMapInput            `pulumi:"nodeLabels"`
	NodePublicIPPrefixID      pulumi.StringPtrInput            `pulumi:"nodePublicIPPrefixID"`
	NodeTaints                pulumi.StringArrayInput          `pulumi:"nodeTaints"`
	OrchestratorVersion       pulumi.StringPtrInput            `pulumi:"orchestratorVersion"`
	OsDiskSizeGB              pulumi.IntPtrInput               `pulumi:"osDiskSizeGB"`
	OsDiskType                pulumi.StringPtrInput            `pulumi:"osDiskType"`
	OsSKU                     pulumi.StringPtrInput            `pulumi:"osSKU"`
	OsType                    pulumi.StringPtrInput            `pulumi:"osType"`
	PodSubnetID               pulumi.StringPtrInput            `pulumi:"podSubnetID"`
	ProximityPlacementGroupID pulumi.StringPtrInput            `pulumi:"proximityPlacementGroupID"`
	ScaleSetEvictionPolicy    pulumi.StringPtrInput            `pulumi:"scaleSetEvictionPolicy"`
	ScaleSetPriority          pulumi.StringPtrInput            `pulumi:"scaleSetPriority"`
	SpotMaxPrice              pulumi.Float64PtrInput           `pulumi:"spotMaxPrice"`
	Tags                      pulumi.StringMapInput            `pulumi:"tags"`
	Type                      pulumi.StringPtrInput            `pulumi:"type"`
	UpgradeSettings           AgentPoolUpgradeSettingsPtrInput `pulumi:"upgradeSettings"`
	VmSize                    pulumi.StringPtrInput            `pulumi:"vmSize"`
	VnetSubnetID              pulumi.StringPtrInput            `pulumi:"vnetSubnetID"`
}

func (ManagedClusterAgentPoolProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAgentPoolProfile)(nil)).Elem()
}

func (i ManagedClusterAgentPoolProfileArgs) ToManagedClusterAgentPoolProfileOutput() ManagedClusterAgentPoolProfileOutput {
	return i.ToManagedClusterAgentPoolProfileOutputWithContext(context.Background())
}

func (i ManagedClusterAgentPoolProfileArgs) ToManagedClusterAgentPoolProfileOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAgentPoolProfileOutput)
}





type ManagedClusterAgentPoolProfileArrayInput interface {
	pulumi.Input

	ToManagedClusterAgentPoolProfileArrayOutput() ManagedClusterAgentPoolProfileArrayOutput
	ToManagedClusterAgentPoolProfileArrayOutputWithContext(context.Context) ManagedClusterAgentPoolProfileArrayOutput
}

type ManagedClusterAgentPoolProfileArray []ManagedClusterAgentPoolProfileInput

func (ManagedClusterAgentPoolProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedClusterAgentPoolProfile)(nil)).Elem()
}

func (i ManagedClusterAgentPoolProfileArray) ToManagedClusterAgentPoolProfileArrayOutput() ManagedClusterAgentPoolProfileArrayOutput {
	return i.ToManagedClusterAgentPoolProfileArrayOutputWithContext(context.Background())
}

func (i ManagedClusterAgentPoolProfileArray) ToManagedClusterAgentPoolProfileArrayOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAgentPoolProfileArrayOutput)
}

type ManagedClusterAgentPoolProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterAgentPoolProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAgentPoolProfile)(nil)).Elem()
}

func (o ManagedClusterAgentPoolProfileOutput) ToManagedClusterAgentPoolProfileOutput() ManagedClusterAgentPoolProfileOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileOutput) ToManagedClusterAgentPoolProfileOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) EnableAutoScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *bool { return v.EnableAutoScaling }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) EnableEncryptionAtHost() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *bool { return v.EnableEncryptionAtHost }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) EnableFIPS() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *bool { return v.EnableFIPS }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) EnableNodePublicIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *bool { return v.EnableNodePublicIP }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) GpuInstanceProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.GpuInstanceProfile }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) KubeletConfig() KubeletConfigPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *KubeletConfig { return v.KubeletConfig }).(KubeletConfigPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) KubeletDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.KubeletDiskType }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) LinuxOSConfig() LinuxOSConfigPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *LinuxOSConfig { return v.LinuxOSConfig }).(LinuxOSConfigPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *int { return v.MaxCount }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *int { return v.MaxPods }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *int { return v.MinCount }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) string { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) map[string]string { return v.NodeLabels }).(pulumi.StringMapOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) NodePublicIPPrefixID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.NodePublicIPPrefixID }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) []string { return v.NodeTaints }).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) OrchestratorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.OrchestratorVersion }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *int { return v.OsDiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) OsDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.OsDiskType }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) OsSKU() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.OsSKU }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) PodSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.PodSubnetID }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) ProximityPlacementGroupID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.ProximityPlacementGroupID }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) ScaleSetEvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.ScaleSetEvictionPolicy }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) ScaleSetPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.ScaleSetPriority }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) SpotMaxPrice() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *float64 { return v.SpotMaxPrice }).(pulumi.Float64PtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) UpgradeSettings() AgentPoolUpgradeSettingsPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *AgentPoolUpgradeSettings { return v.UpgradeSettings }).(AgentPoolUpgradeSettingsPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) VnetSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.VnetSubnetID }).(pulumi.StringPtrOutput)
}

type ManagedClusterAgentPoolProfileArrayOutput struct{ *pulumi.OutputState }

func (ManagedClusterAgentPoolProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedClusterAgentPoolProfile)(nil)).Elem()
}

func (o ManagedClusterAgentPoolProfileArrayOutput) ToManagedClusterAgentPoolProfileArrayOutput() ManagedClusterAgentPoolProfileArrayOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileArrayOutput) ToManagedClusterAgentPoolProfileArrayOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileArrayOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileArrayOutput) Index(i pulumi.IntInput) ManagedClusterAgentPoolProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedClusterAgentPoolProfile {
		return vs[0].([]ManagedClusterAgentPoolProfile)[vs[1].(int)]
	}).(ManagedClusterAgentPoolProfileOutput)
}

type ManagedClusterAgentPoolProfileResponse struct {
	AvailabilityZones         []string                          `pulumi:"availabilityZones"`
	Count                     *int                              `pulumi:"count"`
	EnableAutoScaling         *bool                             `pulumi:"enableAutoScaling"`
	EnableEncryptionAtHost    *bool                             `pulumi:"enableEncryptionAtHost"`
	EnableFIPS                *bool                             `pulumi:"enableFIPS"`
	EnableNodePublicIP        *bool                             `pulumi:"enableNodePublicIP"`
	GpuInstanceProfile        *string                           `pulumi:"gpuInstanceProfile"`
	KubeletConfig             *KubeletConfigResponse            `pulumi:"kubeletConfig"`
	KubeletDiskType           *string                           `pulumi:"kubeletDiskType"`
	LinuxOSConfig             *LinuxOSConfigResponse            `pulumi:"linuxOSConfig"`
	MaxCount                  *int                              `pulumi:"maxCount"`
	MaxPods                   *int                              `pulumi:"maxPods"`
	MinCount                  *int                              `pulumi:"minCount"`
	Mode                      *string                           `pulumi:"mode"`
	Name                      string                            `pulumi:"name"`
	NodeImageVersion          string                            `pulumi:"nodeImageVersion"`
	NodeLabels                map[string]string                 `pulumi:"nodeLabels"`
	NodePublicIPPrefixID      *string                           `pulumi:"nodePublicIPPrefixID"`
	NodeTaints                []string                          `pulumi:"nodeTaints"`
	OrchestratorVersion       *string                           `pulumi:"orchestratorVersion"`
	OsDiskSizeGB              *int                              `pulumi:"osDiskSizeGB"`
	OsDiskType                *string                           `pulumi:"osDiskType"`
	OsSKU                     *string                           `pulumi:"osSKU"`
	OsType                    *string                           `pulumi:"osType"`
	PodSubnetID               *string                           `pulumi:"podSubnetID"`
	PowerState                PowerStateResponse                `pulumi:"powerState"`
	ProvisioningState         string                            `pulumi:"provisioningState"`
	ProximityPlacementGroupID *string                           `pulumi:"proximityPlacementGroupID"`
	ScaleSetEvictionPolicy    *string                           `pulumi:"scaleSetEvictionPolicy"`
	ScaleSetPriority          *string                           `pulumi:"scaleSetPriority"`
	SpotMaxPrice              *float64                          `pulumi:"spotMaxPrice"`
	Tags                      map[string]string                 `pulumi:"tags"`
	Type                      *string                           `pulumi:"type"`
	UpgradeSettings           *AgentPoolUpgradeSettingsResponse `pulumi:"upgradeSettings"`
	VmSize                    *string                           `pulumi:"vmSize"`
	VnetSubnetID              *string                           `pulumi:"vnetSubnetID"`
}





type ManagedClusterAgentPoolProfileResponseInput interface {
	pulumi.Input

	ToManagedClusterAgentPoolProfileResponseOutput() ManagedClusterAgentPoolProfileResponseOutput
	ToManagedClusterAgentPoolProfileResponseOutputWithContext(context.Context) ManagedClusterAgentPoolProfileResponseOutput
}

type ManagedClusterAgentPoolProfileResponseArgs struct {
	AvailabilityZones         pulumi.StringArrayInput                  `pulumi:"availabilityZones"`
	Count                     pulumi.IntPtrInput                       `pulumi:"count"`
	EnableAutoScaling         pulumi.BoolPtrInput                      `pulumi:"enableAutoScaling"`
	EnableEncryptionAtHost    pulumi.BoolPtrInput                      `pulumi:"enableEncryptionAtHost"`
	EnableFIPS                pulumi.BoolPtrInput                      `pulumi:"enableFIPS"`
	EnableNodePublicIP        pulumi.BoolPtrInput                      `pulumi:"enableNodePublicIP"`
	GpuInstanceProfile        pulumi.StringPtrInput                    `pulumi:"gpuInstanceProfile"`
	KubeletConfig             KubeletConfigResponsePtrInput            `pulumi:"kubeletConfig"`
	KubeletDiskType           pulumi.StringPtrInput                    `pulumi:"kubeletDiskType"`
	LinuxOSConfig             LinuxOSConfigResponsePtrInput            `pulumi:"linuxOSConfig"`
	MaxCount                  pulumi.IntPtrInput                       `pulumi:"maxCount"`
	MaxPods                   pulumi.IntPtrInput                       `pulumi:"maxPods"`
	MinCount                  pulumi.IntPtrInput                       `pulumi:"minCount"`
	Mode                      pulumi.StringPtrInput                    `pulumi:"mode"`
	Name                      pulumi.StringInput                       `pulumi:"name"`
	NodeImageVersion          pulumi.StringInput                       `pulumi:"nodeImageVersion"`
	NodeLabels                pulumi.StringMapInput                    `pulumi:"nodeLabels"`
	NodePublicIPPrefixID      pulumi.StringPtrInput                    `pulumi:"nodePublicIPPrefixID"`
	NodeTaints                pulumi.StringArrayInput                  `pulumi:"nodeTaints"`
	OrchestratorVersion       pulumi.StringPtrInput                    `pulumi:"orchestratorVersion"`
	OsDiskSizeGB              pulumi.IntPtrInput                       `pulumi:"osDiskSizeGB"`
	OsDiskType                pulumi.StringPtrInput                    `pulumi:"osDiskType"`
	OsSKU                     pulumi.StringPtrInput                    `pulumi:"osSKU"`
	OsType                    pulumi.StringPtrInput                    `pulumi:"osType"`
	PodSubnetID               pulumi.StringPtrInput                    `pulumi:"podSubnetID"`
	PowerState                PowerStateResponseInput                  `pulumi:"powerState"`
	ProvisioningState         pulumi.StringInput                       `pulumi:"provisioningState"`
	ProximityPlacementGroupID pulumi.StringPtrInput                    `pulumi:"proximityPlacementGroupID"`
	ScaleSetEvictionPolicy    pulumi.StringPtrInput                    `pulumi:"scaleSetEvictionPolicy"`
	ScaleSetPriority          pulumi.StringPtrInput                    `pulumi:"scaleSetPriority"`
	SpotMaxPrice              pulumi.Float64PtrInput                   `pulumi:"spotMaxPrice"`
	Tags                      pulumi.StringMapInput                    `pulumi:"tags"`
	Type                      pulumi.StringPtrInput                    `pulumi:"type"`
	UpgradeSettings           AgentPoolUpgradeSettingsResponsePtrInput `pulumi:"upgradeSettings"`
	VmSize                    pulumi.StringPtrInput                    `pulumi:"vmSize"`
	VnetSubnetID              pulumi.StringPtrInput                    `pulumi:"vnetSubnetID"`
}

func (ManagedClusterAgentPoolProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAgentPoolProfileResponse)(nil)).Elem()
}

func (i ManagedClusterAgentPoolProfileResponseArgs) ToManagedClusterAgentPoolProfileResponseOutput() ManagedClusterAgentPoolProfileResponseOutput {
	return i.ToManagedClusterAgentPoolProfileResponseOutputWithContext(context.Background())
}

func (i ManagedClusterAgentPoolProfileResponseArgs) ToManagedClusterAgentPoolProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAgentPoolProfileResponseOutput)
}





type ManagedClusterAgentPoolProfileResponseArrayInput interface {
	pulumi.Input

	ToManagedClusterAgentPoolProfileResponseArrayOutput() ManagedClusterAgentPoolProfileResponseArrayOutput
	ToManagedClusterAgentPoolProfileResponseArrayOutputWithContext(context.Context) ManagedClusterAgentPoolProfileResponseArrayOutput
}

type ManagedClusterAgentPoolProfileResponseArray []ManagedClusterAgentPoolProfileResponseInput

func (ManagedClusterAgentPoolProfileResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedClusterAgentPoolProfileResponse)(nil)).Elem()
}

func (i ManagedClusterAgentPoolProfileResponseArray) ToManagedClusterAgentPoolProfileResponseArrayOutput() ManagedClusterAgentPoolProfileResponseArrayOutput {
	return i.ToManagedClusterAgentPoolProfileResponseArrayOutputWithContext(context.Background())
}

func (i ManagedClusterAgentPoolProfileResponseArray) ToManagedClusterAgentPoolProfileResponseArrayOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAgentPoolProfileResponseArrayOutput)
}

type ManagedClusterAgentPoolProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterAgentPoolProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAgentPoolProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAgentPoolProfileResponseOutput) ToManagedClusterAgentPoolProfileResponseOutput() ManagedClusterAgentPoolProfileResponseOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileResponseOutput) ToManagedClusterAgentPoolProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileResponseOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileResponseOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) EnableAutoScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *bool { return v.EnableAutoScaling }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) EnableEncryptionAtHost() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *bool { return v.EnableEncryptionAtHost }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) EnableFIPS() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *bool { return v.EnableFIPS }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) EnableNodePublicIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *bool { return v.EnableNodePublicIP }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) GpuInstanceProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.GpuInstanceProfile }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) KubeletConfig() KubeletConfigResponsePtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *KubeletConfigResponse { return v.KubeletConfig }).(KubeletConfigResponsePtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) KubeletDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.KubeletDiskType }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) LinuxOSConfig() LinuxOSConfigResponsePtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *LinuxOSConfigResponse { return v.LinuxOSConfig }).(LinuxOSConfigResponsePtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *int { return v.MaxCount }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *int { return v.MaxPods }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *int { return v.MinCount }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) NodeImageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) string { return v.NodeImageVersion }).(pulumi.StringOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) map[string]string { return v.NodeLabels }).(pulumi.StringMapOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) NodePublicIPPrefixID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.NodePublicIPPrefixID }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) []string { return v.NodeTaints }).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) OrchestratorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.OrchestratorVersion }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *int { return v.OsDiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) OsDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.OsDiskType }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) OsSKU() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.OsSKU }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) PodSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.PodSubnetID }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) PowerState() PowerStateResponseOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) PowerStateResponse { return v.PowerState }).(PowerStateResponseOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) ProximityPlacementGroupID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.ProximityPlacementGroupID }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) ScaleSetEvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.ScaleSetEvictionPolicy }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) ScaleSetPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.ScaleSetPriority }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) SpotMaxPrice() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *float64 { return v.SpotMaxPrice }).(pulumi.Float64PtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) UpgradeSettings() AgentPoolUpgradeSettingsResponsePtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *AgentPoolUpgradeSettingsResponse {
		return v.UpgradeSettings
	}).(AgentPoolUpgradeSettingsResponsePtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) VnetSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.VnetSubnetID }).(pulumi.StringPtrOutput)
}

type ManagedClusterAgentPoolProfileResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagedClusterAgentPoolProfileResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedClusterAgentPoolProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAgentPoolProfileResponseArrayOutput) ToManagedClusterAgentPoolProfileResponseArrayOutput() ManagedClusterAgentPoolProfileResponseArrayOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileResponseArrayOutput) ToManagedClusterAgentPoolProfileResponseArrayOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileResponseArrayOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileResponseArrayOutput) Index(i pulumi.IntInput) ManagedClusterAgentPoolProfileResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedClusterAgentPoolProfileResponse {
		return vs[0].([]ManagedClusterAgentPoolProfileResponse)[vs[1].(int)]
	}).(ManagedClusterAgentPoolProfileResponseOutput)
}

type ManagedClusterAutoUpgradeProfile struct {
	UpgradeChannel *string `pulumi:"upgradeChannel"`
}





type ManagedClusterAutoUpgradeProfileInput interface {
	pulumi.Input

	ToManagedClusterAutoUpgradeProfileOutput() ManagedClusterAutoUpgradeProfileOutput
	ToManagedClusterAutoUpgradeProfileOutputWithContext(context.Context) ManagedClusterAutoUpgradeProfileOutput
}

type ManagedClusterAutoUpgradeProfileArgs struct {
	UpgradeChannel pulumi.StringPtrInput `pulumi:"upgradeChannel"`
}

func (ManagedClusterAutoUpgradeProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAutoUpgradeProfile)(nil)).Elem()
}

func (i ManagedClusterAutoUpgradeProfileArgs) ToManagedClusterAutoUpgradeProfileOutput() ManagedClusterAutoUpgradeProfileOutput {
	return i.ToManagedClusterAutoUpgradeProfileOutputWithContext(context.Background())
}

func (i ManagedClusterAutoUpgradeProfileArgs) ToManagedClusterAutoUpgradeProfileOutputWithContext(ctx context.Context) ManagedClusterAutoUpgradeProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAutoUpgradeProfileOutput)
}

func (i ManagedClusterAutoUpgradeProfileArgs) ToManagedClusterAutoUpgradeProfilePtrOutput() ManagedClusterAutoUpgradeProfilePtrOutput {
	return i.ToManagedClusterAutoUpgradeProfilePtrOutputWithContext(context.Background())
}

func (i ManagedClusterAutoUpgradeProfileArgs) ToManagedClusterAutoUpgradeProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAutoUpgradeProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAutoUpgradeProfileOutput).ToManagedClusterAutoUpgradeProfilePtrOutputWithContext(ctx)
}









type ManagedClusterAutoUpgradeProfilePtrInput interface {
	pulumi.Input

	ToManagedClusterAutoUpgradeProfilePtrOutput() ManagedClusterAutoUpgradeProfilePtrOutput
	ToManagedClusterAutoUpgradeProfilePtrOutputWithContext(context.Context) ManagedClusterAutoUpgradeProfilePtrOutput
}

type managedClusterAutoUpgradeProfilePtrType ManagedClusterAutoUpgradeProfileArgs

func ManagedClusterAutoUpgradeProfilePtr(v *ManagedClusterAutoUpgradeProfileArgs) ManagedClusterAutoUpgradeProfilePtrInput {
	return (*managedClusterAutoUpgradeProfilePtrType)(v)
}

func (*managedClusterAutoUpgradeProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAutoUpgradeProfile)(nil)).Elem()
}

func (i *managedClusterAutoUpgradeProfilePtrType) ToManagedClusterAutoUpgradeProfilePtrOutput() ManagedClusterAutoUpgradeProfilePtrOutput {
	return i.ToManagedClusterAutoUpgradeProfilePtrOutputWithContext(context.Background())
}

func (i *managedClusterAutoUpgradeProfilePtrType) ToManagedClusterAutoUpgradeProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAutoUpgradeProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAutoUpgradeProfilePtrOutput)
}

type ManagedClusterAutoUpgradeProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterAutoUpgradeProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAutoUpgradeProfile)(nil)).Elem()
}

func (o ManagedClusterAutoUpgradeProfileOutput) ToManagedClusterAutoUpgradeProfileOutput() ManagedClusterAutoUpgradeProfileOutput {
	return o
}

func (o ManagedClusterAutoUpgradeProfileOutput) ToManagedClusterAutoUpgradeProfileOutputWithContext(ctx context.Context) ManagedClusterAutoUpgradeProfileOutput {
	return o
}

func (o ManagedClusterAutoUpgradeProfileOutput) ToManagedClusterAutoUpgradeProfilePtrOutput() ManagedClusterAutoUpgradeProfilePtrOutput {
	return o.ToManagedClusterAutoUpgradeProfilePtrOutputWithContext(context.Background())
}

func (o ManagedClusterAutoUpgradeProfileOutput) ToManagedClusterAutoUpgradeProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAutoUpgradeProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterAutoUpgradeProfile) *ManagedClusterAutoUpgradeProfile {
		return &v
	}).(ManagedClusterAutoUpgradeProfilePtrOutput)
}

func (o ManagedClusterAutoUpgradeProfileOutput) UpgradeChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAutoUpgradeProfile) *string { return v.UpgradeChannel }).(pulumi.StringPtrOutput)
}

type ManagedClusterAutoUpgradeProfilePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterAutoUpgradeProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAutoUpgradeProfile)(nil)).Elem()
}

func (o ManagedClusterAutoUpgradeProfilePtrOutput) ToManagedClusterAutoUpgradeProfilePtrOutput() ManagedClusterAutoUpgradeProfilePtrOutput {
	return o
}

func (o ManagedClusterAutoUpgradeProfilePtrOutput) ToManagedClusterAutoUpgradeProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAutoUpgradeProfilePtrOutput {
	return o
}

func (o ManagedClusterAutoUpgradeProfilePtrOutput) Elem() ManagedClusterAutoUpgradeProfileOutput {
	return o.ApplyT(func(v *ManagedClusterAutoUpgradeProfile) ManagedClusterAutoUpgradeProfile {
		if v != nil {
			return *v
		}
		var ret ManagedClusterAutoUpgradeProfile
		return ret
	}).(ManagedClusterAutoUpgradeProfileOutput)
}

func (o ManagedClusterAutoUpgradeProfilePtrOutput) UpgradeChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAutoUpgradeProfile) *string {
		if v == nil {
			return nil
		}
		return v.UpgradeChannel
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterAutoUpgradeProfileResponse struct {
	UpgradeChannel *string `pulumi:"upgradeChannel"`
}





type ManagedClusterAutoUpgradeProfileResponseInput interface {
	pulumi.Input

	ToManagedClusterAutoUpgradeProfileResponseOutput() ManagedClusterAutoUpgradeProfileResponseOutput
	ToManagedClusterAutoUpgradeProfileResponseOutputWithContext(context.Context) ManagedClusterAutoUpgradeProfileResponseOutput
}

type ManagedClusterAutoUpgradeProfileResponseArgs struct {
	UpgradeChannel pulumi.StringPtrInput `pulumi:"upgradeChannel"`
}

func (ManagedClusterAutoUpgradeProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAutoUpgradeProfileResponse)(nil)).Elem()
}

func (i ManagedClusterAutoUpgradeProfileResponseArgs) ToManagedClusterAutoUpgradeProfileResponseOutput() ManagedClusterAutoUpgradeProfileResponseOutput {
	return i.ToManagedClusterAutoUpgradeProfileResponseOutputWithContext(context.Background())
}

func (i ManagedClusterAutoUpgradeProfileResponseArgs) ToManagedClusterAutoUpgradeProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAutoUpgradeProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAutoUpgradeProfileResponseOutput)
}

func (i ManagedClusterAutoUpgradeProfileResponseArgs) ToManagedClusterAutoUpgradeProfileResponsePtrOutput() ManagedClusterAutoUpgradeProfileResponsePtrOutput {
	return i.ToManagedClusterAutoUpgradeProfileResponsePtrOutputWithContext(context.Background())
}

func (i ManagedClusterAutoUpgradeProfileResponseArgs) ToManagedClusterAutoUpgradeProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAutoUpgradeProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAutoUpgradeProfileResponseOutput).ToManagedClusterAutoUpgradeProfileResponsePtrOutputWithContext(ctx)
}









type ManagedClusterAutoUpgradeProfileResponsePtrInput interface {
	pulumi.Input

	ToManagedClusterAutoUpgradeProfileResponsePtrOutput() ManagedClusterAutoUpgradeProfileResponsePtrOutput
	ToManagedClusterAutoUpgradeProfileResponsePtrOutputWithContext(context.Context) ManagedClusterAutoUpgradeProfileResponsePtrOutput
}

type managedClusterAutoUpgradeProfileResponsePtrType ManagedClusterAutoUpgradeProfileResponseArgs

func ManagedClusterAutoUpgradeProfileResponsePtr(v *ManagedClusterAutoUpgradeProfileResponseArgs) ManagedClusterAutoUpgradeProfileResponsePtrInput {
	return (*managedClusterAutoUpgradeProfileResponsePtrType)(v)
}

func (*managedClusterAutoUpgradeProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAutoUpgradeProfileResponse)(nil)).Elem()
}

func (i *managedClusterAutoUpgradeProfileResponsePtrType) ToManagedClusterAutoUpgradeProfileResponsePtrOutput() ManagedClusterAutoUpgradeProfileResponsePtrOutput {
	return i.ToManagedClusterAutoUpgradeProfileResponsePtrOutputWithContext(context.Background())
}

func (i *managedClusterAutoUpgradeProfileResponsePtrType) ToManagedClusterAutoUpgradeProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAutoUpgradeProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAutoUpgradeProfileResponsePtrOutput)
}

type ManagedClusterAutoUpgradeProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterAutoUpgradeProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAutoUpgradeProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAutoUpgradeProfileResponseOutput) ToManagedClusterAutoUpgradeProfileResponseOutput() ManagedClusterAutoUpgradeProfileResponseOutput {
	return o
}

func (o ManagedClusterAutoUpgradeProfileResponseOutput) ToManagedClusterAutoUpgradeProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAutoUpgradeProfileResponseOutput {
	return o
}

func (o ManagedClusterAutoUpgradeProfileResponseOutput) ToManagedClusterAutoUpgradeProfileResponsePtrOutput() ManagedClusterAutoUpgradeProfileResponsePtrOutput {
	return o.ToManagedClusterAutoUpgradeProfileResponsePtrOutputWithContext(context.Background())
}

func (o ManagedClusterAutoUpgradeProfileResponseOutput) ToManagedClusterAutoUpgradeProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAutoUpgradeProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterAutoUpgradeProfileResponse) *ManagedClusterAutoUpgradeProfileResponse {
		return &v
	}).(ManagedClusterAutoUpgradeProfileResponsePtrOutput)
}

func (o ManagedClusterAutoUpgradeProfileResponseOutput) UpgradeChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAutoUpgradeProfileResponse) *string { return v.UpgradeChannel }).(pulumi.StringPtrOutput)
}

type ManagedClusterAutoUpgradeProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterAutoUpgradeProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAutoUpgradeProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAutoUpgradeProfileResponsePtrOutput) ToManagedClusterAutoUpgradeProfileResponsePtrOutput() ManagedClusterAutoUpgradeProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterAutoUpgradeProfileResponsePtrOutput) ToManagedClusterAutoUpgradeProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAutoUpgradeProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterAutoUpgradeProfileResponsePtrOutput) Elem() ManagedClusterAutoUpgradeProfileResponseOutput {
	return o.ApplyT(func(v *ManagedClusterAutoUpgradeProfileResponse) ManagedClusterAutoUpgradeProfileResponse {
		if v != nil {
			return *v
		}
		var ret ManagedClusterAutoUpgradeProfileResponse
		return ret
	}).(ManagedClusterAutoUpgradeProfileResponseOutput)
}

func (o ManagedClusterAutoUpgradeProfileResponsePtrOutput) UpgradeChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAutoUpgradeProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.UpgradeChannel
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterHTTPProxyConfig struct {
	HttpProxy  *string  `pulumi:"httpProxy"`
	HttpsProxy *string  `pulumi:"httpsProxy"`
	NoProxy    []string `pulumi:"noProxy"`
	TrustedCa  *string  `pulumi:"trustedCa"`
}





type ManagedClusterHTTPProxyConfigInput interface {
	pulumi.Input

	ToManagedClusterHTTPProxyConfigOutput() ManagedClusterHTTPProxyConfigOutput
	ToManagedClusterHTTPProxyConfigOutputWithContext(context.Context) ManagedClusterHTTPProxyConfigOutput
}

type ManagedClusterHTTPProxyConfigArgs struct {
	HttpProxy  pulumi.StringPtrInput   `pulumi:"httpProxy"`
	HttpsProxy pulumi.StringPtrInput   `pulumi:"httpsProxy"`
	NoProxy    pulumi.StringArrayInput `pulumi:"noProxy"`
	TrustedCa  pulumi.StringPtrInput   `pulumi:"trustedCa"`
}

func (ManagedClusterHTTPProxyConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterHTTPProxyConfig)(nil)).Elem()
}

func (i ManagedClusterHTTPProxyConfigArgs) ToManagedClusterHTTPProxyConfigOutput() ManagedClusterHTTPProxyConfigOutput {
	return i.ToManagedClusterHTTPProxyConfigOutputWithContext(context.Background())
}

func (i ManagedClusterHTTPProxyConfigArgs) ToManagedClusterHTTPProxyConfigOutputWithContext(ctx context.Context) ManagedClusterHTTPProxyConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterHTTPProxyConfigOutput)
}

func (i ManagedClusterHTTPProxyConfigArgs) ToManagedClusterHTTPProxyConfigPtrOutput() ManagedClusterHTTPProxyConfigPtrOutput {
	return i.ToManagedClusterHTTPProxyConfigPtrOutputWithContext(context.Background())
}

func (i ManagedClusterHTTPProxyConfigArgs) ToManagedClusterHTTPProxyConfigPtrOutputWithContext(ctx context.Context) ManagedClusterHTTPProxyConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterHTTPProxyConfigOutput).ToManagedClusterHTTPProxyConfigPtrOutputWithContext(ctx)
}









type ManagedClusterHTTPProxyConfigPtrInput interface {
	pulumi.Input

	ToManagedClusterHTTPProxyConfigPtrOutput() ManagedClusterHTTPProxyConfigPtrOutput
	ToManagedClusterHTTPProxyConfigPtrOutputWithContext(context.Context) ManagedClusterHTTPProxyConfigPtrOutput
}

type managedClusterHTTPProxyConfigPtrType ManagedClusterHTTPProxyConfigArgs

func ManagedClusterHTTPProxyConfigPtr(v *ManagedClusterHTTPProxyConfigArgs) ManagedClusterHTTPProxyConfigPtrInput {
	return (*managedClusterHTTPProxyConfigPtrType)(v)
}

func (*managedClusterHTTPProxyConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterHTTPProxyConfig)(nil)).Elem()
}

func (i *managedClusterHTTPProxyConfigPtrType) ToManagedClusterHTTPProxyConfigPtrOutput() ManagedClusterHTTPProxyConfigPtrOutput {
	return i.ToManagedClusterHTTPProxyConfigPtrOutputWithContext(context.Background())
}

func (i *managedClusterHTTPProxyConfigPtrType) ToManagedClusterHTTPProxyConfigPtrOutputWithContext(ctx context.Context) ManagedClusterHTTPProxyConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterHTTPProxyConfigPtrOutput)
}

type ManagedClusterHTTPProxyConfigOutput struct{ *pulumi.OutputState }

func (ManagedClusterHTTPProxyConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterHTTPProxyConfig)(nil)).Elem()
}

func (o ManagedClusterHTTPProxyConfigOutput) ToManagedClusterHTTPProxyConfigOutput() ManagedClusterHTTPProxyConfigOutput {
	return o
}

func (o ManagedClusterHTTPProxyConfigOutput) ToManagedClusterHTTPProxyConfigOutputWithContext(ctx context.Context) ManagedClusterHTTPProxyConfigOutput {
	return o
}

func (o ManagedClusterHTTPProxyConfigOutput) ToManagedClusterHTTPProxyConfigPtrOutput() ManagedClusterHTTPProxyConfigPtrOutput {
	return o.ToManagedClusterHTTPProxyConfigPtrOutputWithContext(context.Background())
}

func (o ManagedClusterHTTPProxyConfigOutput) ToManagedClusterHTTPProxyConfigPtrOutputWithContext(ctx context.Context) ManagedClusterHTTPProxyConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterHTTPProxyConfig) *ManagedClusterHTTPProxyConfig {
		return &v
	}).(ManagedClusterHTTPProxyConfigPtrOutput)
}

func (o ManagedClusterHTTPProxyConfigOutput) HttpProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterHTTPProxyConfig) *string { return v.HttpProxy }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterHTTPProxyConfigOutput) HttpsProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterHTTPProxyConfig) *string { return v.HttpsProxy }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterHTTPProxyConfigOutput) NoProxy() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedClusterHTTPProxyConfig) []string { return v.NoProxy }).(pulumi.StringArrayOutput)
}

func (o ManagedClusterHTTPProxyConfigOutput) TrustedCa() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterHTTPProxyConfig) *string { return v.TrustedCa }).(pulumi.StringPtrOutput)
}

type ManagedClusterHTTPProxyConfigPtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterHTTPProxyConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterHTTPProxyConfig)(nil)).Elem()
}

func (o ManagedClusterHTTPProxyConfigPtrOutput) ToManagedClusterHTTPProxyConfigPtrOutput() ManagedClusterHTTPProxyConfigPtrOutput {
	return o
}

func (o ManagedClusterHTTPProxyConfigPtrOutput) ToManagedClusterHTTPProxyConfigPtrOutputWithContext(ctx context.Context) ManagedClusterHTTPProxyConfigPtrOutput {
	return o
}

func (o ManagedClusterHTTPProxyConfigPtrOutput) Elem() ManagedClusterHTTPProxyConfigOutput {
	return o.ApplyT(func(v *ManagedClusterHTTPProxyConfig) ManagedClusterHTTPProxyConfig {
		if v != nil {
			return *v
		}
		var ret ManagedClusterHTTPProxyConfig
		return ret
	}).(ManagedClusterHTTPProxyConfigOutput)
}

func (o ManagedClusterHTTPProxyConfigPtrOutput) HttpProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterHTTPProxyConfig) *string {
		if v == nil {
			return nil
		}
		return v.HttpProxy
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterHTTPProxyConfigPtrOutput) HttpsProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterHTTPProxyConfig) *string {
		if v == nil {
			return nil
		}
		return v.HttpsProxy
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterHTTPProxyConfigPtrOutput) NoProxy() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedClusterHTTPProxyConfig) []string {
		if v == nil {
			return nil
		}
		return v.NoProxy
	}).(pulumi.StringArrayOutput)
}

func (o ManagedClusterHTTPProxyConfigPtrOutput) TrustedCa() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterHTTPProxyConfig) *string {
		if v == nil {
			return nil
		}
		return v.TrustedCa
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterHTTPProxyConfigResponse struct {
	HttpProxy  *string  `pulumi:"httpProxy"`
	HttpsProxy *string  `pulumi:"httpsProxy"`
	NoProxy    []string `pulumi:"noProxy"`
	TrustedCa  *string  `pulumi:"trustedCa"`
}





type ManagedClusterHTTPProxyConfigResponseInput interface {
	pulumi.Input

	ToManagedClusterHTTPProxyConfigResponseOutput() ManagedClusterHTTPProxyConfigResponseOutput
	ToManagedClusterHTTPProxyConfigResponseOutputWithContext(context.Context) ManagedClusterHTTPProxyConfigResponseOutput
}

type ManagedClusterHTTPProxyConfigResponseArgs struct {
	HttpProxy  pulumi.StringPtrInput   `pulumi:"httpProxy"`
	HttpsProxy pulumi.StringPtrInput   `pulumi:"httpsProxy"`
	NoProxy    pulumi.StringArrayInput `pulumi:"noProxy"`
	TrustedCa  pulumi.StringPtrInput   `pulumi:"trustedCa"`
}

func (ManagedClusterHTTPProxyConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterHTTPProxyConfigResponse)(nil)).Elem()
}

func (i ManagedClusterHTTPProxyConfigResponseArgs) ToManagedClusterHTTPProxyConfigResponseOutput() ManagedClusterHTTPProxyConfigResponseOutput {
	return i.ToManagedClusterHTTPProxyConfigResponseOutputWithContext(context.Background())
}

func (i ManagedClusterHTTPProxyConfigResponseArgs) ToManagedClusterHTTPProxyConfigResponseOutputWithContext(ctx context.Context) ManagedClusterHTTPProxyConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterHTTPProxyConfigResponseOutput)
}

func (i ManagedClusterHTTPProxyConfigResponseArgs) ToManagedClusterHTTPProxyConfigResponsePtrOutput() ManagedClusterHTTPProxyConfigResponsePtrOutput {
	return i.ToManagedClusterHTTPProxyConfigResponsePtrOutputWithContext(context.Background())
}

func (i ManagedClusterHTTPProxyConfigResponseArgs) ToManagedClusterHTTPProxyConfigResponsePtrOutputWithContext(ctx context.Context) ManagedClusterHTTPProxyConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterHTTPProxyConfigResponseOutput).ToManagedClusterHTTPProxyConfigResponsePtrOutputWithContext(ctx)
}









type ManagedClusterHTTPProxyConfigResponsePtrInput interface {
	pulumi.Input

	ToManagedClusterHTTPProxyConfigResponsePtrOutput() ManagedClusterHTTPProxyConfigResponsePtrOutput
	ToManagedClusterHTTPProxyConfigResponsePtrOutputWithContext(context.Context) ManagedClusterHTTPProxyConfigResponsePtrOutput
}

type managedClusterHTTPProxyConfigResponsePtrType ManagedClusterHTTPProxyConfigResponseArgs

func ManagedClusterHTTPProxyConfigResponsePtr(v *ManagedClusterHTTPProxyConfigResponseArgs) ManagedClusterHTTPProxyConfigResponsePtrInput {
	return (*managedClusterHTTPProxyConfigResponsePtrType)(v)
}

func (*managedClusterHTTPProxyConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterHTTPProxyConfigResponse)(nil)).Elem()
}

func (i *managedClusterHTTPProxyConfigResponsePtrType) ToManagedClusterHTTPProxyConfigResponsePtrOutput() ManagedClusterHTTPProxyConfigResponsePtrOutput {
	return i.ToManagedClusterHTTPProxyConfigResponsePtrOutputWithContext(context.Background())
}

func (i *managedClusterHTTPProxyConfigResponsePtrType) ToManagedClusterHTTPProxyConfigResponsePtrOutputWithContext(ctx context.Context) ManagedClusterHTTPProxyConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterHTTPProxyConfigResponsePtrOutput)
}

type ManagedClusterHTTPProxyConfigResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterHTTPProxyConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterHTTPProxyConfigResponse)(nil)).Elem()
}

func (o ManagedClusterHTTPProxyConfigResponseOutput) ToManagedClusterHTTPProxyConfigResponseOutput() ManagedClusterHTTPProxyConfigResponseOutput {
	return o
}

func (o ManagedClusterHTTPProxyConfigResponseOutput) ToManagedClusterHTTPProxyConfigResponseOutputWithContext(ctx context.Context) ManagedClusterHTTPProxyConfigResponseOutput {
	return o
}

func (o ManagedClusterHTTPProxyConfigResponseOutput) ToManagedClusterHTTPProxyConfigResponsePtrOutput() ManagedClusterHTTPProxyConfigResponsePtrOutput {
	return o.ToManagedClusterHTTPProxyConfigResponsePtrOutputWithContext(context.Background())
}

func (o ManagedClusterHTTPProxyConfigResponseOutput) ToManagedClusterHTTPProxyConfigResponsePtrOutputWithContext(ctx context.Context) ManagedClusterHTTPProxyConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterHTTPProxyConfigResponse) *ManagedClusterHTTPProxyConfigResponse {
		return &v
	}).(ManagedClusterHTTPProxyConfigResponsePtrOutput)
}

func (o ManagedClusterHTTPProxyConfigResponseOutput) HttpProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterHTTPProxyConfigResponse) *string { return v.HttpProxy }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterHTTPProxyConfigResponseOutput) HttpsProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterHTTPProxyConfigResponse) *string { return v.HttpsProxy }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterHTTPProxyConfigResponseOutput) NoProxy() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedClusterHTTPProxyConfigResponse) []string { return v.NoProxy }).(pulumi.StringArrayOutput)
}

func (o ManagedClusterHTTPProxyConfigResponseOutput) TrustedCa() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterHTTPProxyConfigResponse) *string { return v.TrustedCa }).(pulumi.StringPtrOutput)
}

type ManagedClusterHTTPProxyConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterHTTPProxyConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterHTTPProxyConfigResponse)(nil)).Elem()
}

func (o ManagedClusterHTTPProxyConfigResponsePtrOutput) ToManagedClusterHTTPProxyConfigResponsePtrOutput() ManagedClusterHTTPProxyConfigResponsePtrOutput {
	return o
}

func (o ManagedClusterHTTPProxyConfigResponsePtrOutput) ToManagedClusterHTTPProxyConfigResponsePtrOutputWithContext(ctx context.Context) ManagedClusterHTTPProxyConfigResponsePtrOutput {
	return o
}

func (o ManagedClusterHTTPProxyConfigResponsePtrOutput) Elem() ManagedClusterHTTPProxyConfigResponseOutput {
	return o.ApplyT(func(v *ManagedClusterHTTPProxyConfigResponse) ManagedClusterHTTPProxyConfigResponse {
		if v != nil {
			return *v
		}
		var ret ManagedClusterHTTPProxyConfigResponse
		return ret
	}).(ManagedClusterHTTPProxyConfigResponseOutput)
}

func (o ManagedClusterHTTPProxyConfigResponsePtrOutput) HttpProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterHTTPProxyConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.HttpProxy
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterHTTPProxyConfigResponsePtrOutput) HttpsProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterHTTPProxyConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.HttpsProxy
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterHTTPProxyConfigResponsePtrOutput) NoProxy() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedClusterHTTPProxyConfigResponse) []string {
		if v == nil {
			return nil
		}
		return v.NoProxy
	}).(pulumi.StringArrayOutput)
}

func (o ManagedClusterHTTPProxyConfigResponsePtrOutput) TrustedCa() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterHTTPProxyConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.TrustedCa
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterIdentity struct {
	Type                   *ResourceIdentityType  `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ManagedClusterIdentityInput interface {
	pulumi.Input

	ToManagedClusterIdentityOutput() ManagedClusterIdentityOutput
	ToManagedClusterIdentityOutputWithContext(context.Context) ManagedClusterIdentityOutput
}

type ManagedClusterIdentityArgs struct {
	Type                   ResourceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput              `pulumi:"userAssignedIdentities"`
}

func (ManagedClusterIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterIdentity)(nil)).Elem()
}

func (i ManagedClusterIdentityArgs) ToManagedClusterIdentityOutput() ManagedClusterIdentityOutput {
	return i.ToManagedClusterIdentityOutputWithContext(context.Background())
}

func (i ManagedClusterIdentityArgs) ToManagedClusterIdentityOutputWithContext(ctx context.Context) ManagedClusterIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterIdentityOutput)
}

func (i ManagedClusterIdentityArgs) ToManagedClusterIdentityPtrOutput() ManagedClusterIdentityPtrOutput {
	return i.ToManagedClusterIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedClusterIdentityArgs) ToManagedClusterIdentityPtrOutputWithContext(ctx context.Context) ManagedClusterIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterIdentityOutput).ToManagedClusterIdentityPtrOutputWithContext(ctx)
}









type ManagedClusterIdentityPtrInput interface {
	pulumi.Input

	ToManagedClusterIdentityPtrOutput() ManagedClusterIdentityPtrOutput
	ToManagedClusterIdentityPtrOutputWithContext(context.Context) ManagedClusterIdentityPtrOutput
}

type managedClusterIdentityPtrType ManagedClusterIdentityArgs

func ManagedClusterIdentityPtr(v *ManagedClusterIdentityArgs) ManagedClusterIdentityPtrInput {
	return (*managedClusterIdentityPtrType)(v)
}

func (*managedClusterIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterIdentity)(nil)).Elem()
}

func (i *managedClusterIdentityPtrType) ToManagedClusterIdentityPtrOutput() ManagedClusterIdentityPtrOutput {
	return i.ToManagedClusterIdentityPtrOutputWithContext(context.Background())
}

func (i *managedClusterIdentityPtrType) ToManagedClusterIdentityPtrOutputWithContext(ctx context.Context) ManagedClusterIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterIdentityPtrOutput)
}

type ManagedClusterIdentityOutput struct{ *pulumi.OutputState }

func (ManagedClusterIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterIdentity)(nil)).Elem()
}

func (o ManagedClusterIdentityOutput) ToManagedClusterIdentityOutput() ManagedClusterIdentityOutput {
	return o
}

func (o ManagedClusterIdentityOutput) ToManagedClusterIdentityOutputWithContext(ctx context.Context) ManagedClusterIdentityOutput {
	return o
}

func (o ManagedClusterIdentityOutput) ToManagedClusterIdentityPtrOutput() ManagedClusterIdentityPtrOutput {
	return o.ToManagedClusterIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedClusterIdentityOutput) ToManagedClusterIdentityPtrOutputWithContext(ctx context.Context) ManagedClusterIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterIdentity) *ManagedClusterIdentity {
		return &v
	}).(ManagedClusterIdentityPtrOutput)
}

func (o ManagedClusterIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v ManagedClusterIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

func (o ManagedClusterIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ManagedClusterIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ManagedClusterIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterIdentity)(nil)).Elem()
}

func (o ManagedClusterIdentityPtrOutput) ToManagedClusterIdentityPtrOutput() ManagedClusterIdentityPtrOutput {
	return o
}

func (o ManagedClusterIdentityPtrOutput) ToManagedClusterIdentityPtrOutputWithContext(ctx context.Context) ManagedClusterIdentityPtrOutput {
	return o
}

func (o ManagedClusterIdentityPtrOutput) Elem() ManagedClusterIdentityOutput {
	return o.ApplyT(func(v *ManagedClusterIdentity) ManagedClusterIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedClusterIdentity
		return ret
	}).(ManagedClusterIdentityOutput)
}

func (o ManagedClusterIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *ManagedClusterIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

func (o ManagedClusterIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ManagedClusterIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ManagedClusterIdentityResponse struct {
	PrincipalId            string                                                          `pulumi:"principalId"`
	TenantId               string                                                          `pulumi:"tenantId"`
	Type                   *string                                                         `pulumi:"type"`
	UserAssignedIdentities map[string]ManagedClusterIdentityResponseUserAssignedIdentities `pulumi:"userAssignedIdentities"`
}





type ManagedClusterIdentityResponseInput interface {
	pulumi.Input

	ToManagedClusterIdentityResponseOutput() ManagedClusterIdentityResponseOutput
	ToManagedClusterIdentityResponseOutputWithContext(context.Context) ManagedClusterIdentityResponseOutput
}

type ManagedClusterIdentityResponseArgs struct {
	PrincipalId            pulumi.StringInput                                           `pulumi:"principalId"`
	TenantId               pulumi.StringInput                                           `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput                                        `pulumi:"type"`
	UserAssignedIdentities ManagedClusterIdentityResponseUserAssignedIdentitiesMapInput `pulumi:"userAssignedIdentities"`
}

func (ManagedClusterIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterIdentityResponse)(nil)).Elem()
}

func (i ManagedClusterIdentityResponseArgs) ToManagedClusterIdentityResponseOutput() ManagedClusterIdentityResponseOutput {
	return i.ToManagedClusterIdentityResponseOutputWithContext(context.Background())
}

func (i ManagedClusterIdentityResponseArgs) ToManagedClusterIdentityResponseOutputWithContext(ctx context.Context) ManagedClusterIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterIdentityResponseOutput)
}

func (i ManagedClusterIdentityResponseArgs) ToManagedClusterIdentityResponsePtrOutput() ManagedClusterIdentityResponsePtrOutput {
	return i.ToManagedClusterIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ManagedClusterIdentityResponseArgs) ToManagedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedClusterIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterIdentityResponseOutput).ToManagedClusterIdentityResponsePtrOutputWithContext(ctx)
}









type ManagedClusterIdentityResponsePtrInput interface {
	pulumi.Input

	ToManagedClusterIdentityResponsePtrOutput() ManagedClusterIdentityResponsePtrOutput
	ToManagedClusterIdentityResponsePtrOutputWithContext(context.Context) ManagedClusterIdentityResponsePtrOutput
}

type managedClusterIdentityResponsePtrType ManagedClusterIdentityResponseArgs

func ManagedClusterIdentityResponsePtr(v *ManagedClusterIdentityResponseArgs) ManagedClusterIdentityResponsePtrInput {
	return (*managedClusterIdentityResponsePtrType)(v)
}

func (*managedClusterIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterIdentityResponse)(nil)).Elem()
}

func (i *managedClusterIdentityResponsePtrType) ToManagedClusterIdentityResponsePtrOutput() ManagedClusterIdentityResponsePtrOutput {
	return i.ToManagedClusterIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *managedClusterIdentityResponsePtrType) ToManagedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedClusterIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterIdentityResponsePtrOutput)
}

type ManagedClusterIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterIdentityResponse)(nil)).Elem()
}

func (o ManagedClusterIdentityResponseOutput) ToManagedClusterIdentityResponseOutput() ManagedClusterIdentityResponseOutput {
	return o
}

func (o ManagedClusterIdentityResponseOutput) ToManagedClusterIdentityResponseOutputWithContext(ctx context.Context) ManagedClusterIdentityResponseOutput {
	return o
}

func (o ManagedClusterIdentityResponseOutput) ToManagedClusterIdentityResponsePtrOutput() ManagedClusterIdentityResponsePtrOutput {
	return o.ToManagedClusterIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ManagedClusterIdentityResponseOutput) ToManagedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedClusterIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterIdentityResponse) *ManagedClusterIdentityResponse {
		return &v
	}).(ManagedClusterIdentityResponsePtrOutput)
}

func (o ManagedClusterIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedClusterIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedClusterIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterIdentityResponseOutput) UserAssignedIdentities() ManagedClusterIdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v ManagedClusterIdentityResponse) map[string]ManagedClusterIdentityResponseUserAssignedIdentities {
		return v.UserAssignedIdentities
	}).(ManagedClusterIdentityResponseUserAssignedIdentitiesMapOutput)
}

type ManagedClusterIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterIdentityResponse)(nil)).Elem()
}

func (o ManagedClusterIdentityResponsePtrOutput) ToManagedClusterIdentityResponsePtrOutput() ManagedClusterIdentityResponsePtrOutput {
	return o
}

func (o ManagedClusterIdentityResponsePtrOutput) ToManagedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedClusterIdentityResponsePtrOutput {
	return o
}

func (o ManagedClusterIdentityResponsePtrOutput) Elem() ManagedClusterIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedClusterIdentityResponse) ManagedClusterIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedClusterIdentityResponse
		return ret
	}).(ManagedClusterIdentityResponseOutput)
}

func (o ManagedClusterIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterIdentityResponsePtrOutput) UserAssignedIdentities() ManagedClusterIdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v *ManagedClusterIdentityResponse) map[string]ManagedClusterIdentityResponseUserAssignedIdentities {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(ManagedClusterIdentityResponseUserAssignedIdentitiesMapOutput)
}

type ManagedClusterIdentityResponseUserAssignedIdentities struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}





type ManagedClusterIdentityResponseUserAssignedIdentitiesInput interface {
	pulumi.Input

	ToManagedClusterIdentityResponseUserAssignedIdentitiesOutput() ManagedClusterIdentityResponseUserAssignedIdentitiesOutput
	ToManagedClusterIdentityResponseUserAssignedIdentitiesOutputWithContext(context.Context) ManagedClusterIdentityResponseUserAssignedIdentitiesOutput
}

type ManagedClusterIdentityResponseUserAssignedIdentitiesArgs struct {
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (ManagedClusterIdentityResponseUserAssignedIdentitiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (i ManagedClusterIdentityResponseUserAssignedIdentitiesArgs) ToManagedClusterIdentityResponseUserAssignedIdentitiesOutput() ManagedClusterIdentityResponseUserAssignedIdentitiesOutput {
	return i.ToManagedClusterIdentityResponseUserAssignedIdentitiesOutputWithContext(context.Background())
}

func (i ManagedClusterIdentityResponseUserAssignedIdentitiesArgs) ToManagedClusterIdentityResponseUserAssignedIdentitiesOutputWithContext(ctx context.Context) ManagedClusterIdentityResponseUserAssignedIdentitiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterIdentityResponseUserAssignedIdentitiesOutput)
}





type ManagedClusterIdentityResponseUserAssignedIdentitiesMapInput interface {
	pulumi.Input

	ToManagedClusterIdentityResponseUserAssignedIdentitiesMapOutput() ManagedClusterIdentityResponseUserAssignedIdentitiesMapOutput
	ToManagedClusterIdentityResponseUserAssignedIdentitiesMapOutputWithContext(context.Context) ManagedClusterIdentityResponseUserAssignedIdentitiesMapOutput
}

type ManagedClusterIdentityResponseUserAssignedIdentitiesMap map[string]ManagedClusterIdentityResponseUserAssignedIdentitiesInput

func (ManagedClusterIdentityResponseUserAssignedIdentitiesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (i ManagedClusterIdentityResponseUserAssignedIdentitiesMap) ToManagedClusterIdentityResponseUserAssignedIdentitiesMapOutput() ManagedClusterIdentityResponseUserAssignedIdentitiesMapOutput {
	return i.ToManagedClusterIdentityResponseUserAssignedIdentitiesMapOutputWithContext(context.Background())
}

func (i ManagedClusterIdentityResponseUserAssignedIdentitiesMap) ToManagedClusterIdentityResponseUserAssignedIdentitiesMapOutputWithContext(ctx context.Context) ManagedClusterIdentityResponseUserAssignedIdentitiesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterIdentityResponseUserAssignedIdentitiesMapOutput)
}

type ManagedClusterIdentityResponseUserAssignedIdentitiesOutput struct{ *pulumi.OutputState }

func (ManagedClusterIdentityResponseUserAssignedIdentitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o ManagedClusterIdentityResponseUserAssignedIdentitiesOutput) ToManagedClusterIdentityResponseUserAssignedIdentitiesOutput() ManagedClusterIdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o ManagedClusterIdentityResponseUserAssignedIdentitiesOutput) ToManagedClusterIdentityResponseUserAssignedIdentitiesOutputWithContext(ctx context.Context) ManagedClusterIdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o ManagedClusterIdentityResponseUserAssignedIdentitiesOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterIdentityResponseUserAssignedIdentities) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ManagedClusterIdentityResponseUserAssignedIdentitiesOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterIdentityResponseUserAssignedIdentities) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type ManagedClusterIdentityResponseUserAssignedIdentitiesMapOutput struct{ *pulumi.OutputState }

func (ManagedClusterIdentityResponseUserAssignedIdentitiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o ManagedClusterIdentityResponseUserAssignedIdentitiesMapOutput) ToManagedClusterIdentityResponseUserAssignedIdentitiesMapOutput() ManagedClusterIdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o ManagedClusterIdentityResponseUserAssignedIdentitiesMapOutput) ToManagedClusterIdentityResponseUserAssignedIdentitiesMapOutputWithContext(ctx context.Context) ManagedClusterIdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o ManagedClusterIdentityResponseUserAssignedIdentitiesMapOutput) MapIndex(k pulumi.StringInput) ManagedClusterIdentityResponseUserAssignedIdentitiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ManagedClusterIdentityResponseUserAssignedIdentities {
		return vs[0].(map[string]ManagedClusterIdentityResponseUserAssignedIdentities)[vs[1].(string)]
	}).(ManagedClusterIdentityResponseUserAssignedIdentitiesOutput)
}

type ManagedClusterLoadBalancerProfile struct {
	AllocatedOutboundPorts *int                                                 `pulumi:"allocatedOutboundPorts"`
	EffectiveOutboundIPs   []ResourceReference                                  `pulumi:"effectiveOutboundIPs"`
	IdleTimeoutInMinutes   *int                                                 `pulumi:"idleTimeoutInMinutes"`
	ManagedOutboundIPs     *ManagedClusterLoadBalancerProfileManagedOutboundIPs `pulumi:"managedOutboundIPs"`
	OutboundIPPrefixes     *ManagedClusterLoadBalancerProfileOutboundIPPrefixes `pulumi:"outboundIPPrefixes"`
	OutboundIPs            *ManagedClusterLoadBalancerProfileOutboundIPs        `pulumi:"outboundIPs"`
}


func (val *ManagedClusterLoadBalancerProfile) Defaults() *ManagedClusterLoadBalancerProfile {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllocatedOutboundPorts) {
		allocatedOutboundPorts_ := 0
		tmp.AllocatedOutboundPorts = &allocatedOutboundPorts_
	}
	if isZero(tmp.IdleTimeoutInMinutes) {
		idleTimeoutInMinutes_ := 30
		tmp.IdleTimeoutInMinutes = &idleTimeoutInMinutes_
	}
	tmp.ManagedOutboundIPs = tmp.ManagedOutboundIPs.Defaults()

	return &tmp
}





type ManagedClusterLoadBalancerProfileInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileOutput() ManagedClusterLoadBalancerProfileOutput
	ToManagedClusterLoadBalancerProfileOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileOutput
}

type ManagedClusterLoadBalancerProfileArgs struct {
	AllocatedOutboundPorts pulumi.IntPtrInput                                          `pulumi:"allocatedOutboundPorts"`
	EffectiveOutboundIPs   ResourceReferenceArrayInput                                 `pulumi:"effectiveOutboundIPs"`
	IdleTimeoutInMinutes   pulumi.IntPtrInput                                          `pulumi:"idleTimeoutInMinutes"`
	ManagedOutboundIPs     ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrInput `pulumi:"managedOutboundIPs"`
	OutboundIPPrefixes     ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrInput `pulumi:"outboundIPPrefixes"`
	OutboundIPs            ManagedClusterLoadBalancerProfileOutboundIPsPtrInput        `pulumi:"outboundIPs"`
}

func (ManagedClusterLoadBalancerProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfile)(nil)).Elem()
}

func (i ManagedClusterLoadBalancerProfileArgs) ToManagedClusterLoadBalancerProfileOutput() ManagedClusterLoadBalancerProfileOutput {
	return i.ToManagedClusterLoadBalancerProfileOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileArgs) ToManagedClusterLoadBalancerProfileOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileOutput)
}

func (i ManagedClusterLoadBalancerProfileArgs) ToManagedClusterLoadBalancerProfilePtrOutput() ManagedClusterLoadBalancerProfilePtrOutput {
	return i.ToManagedClusterLoadBalancerProfilePtrOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileArgs) ToManagedClusterLoadBalancerProfilePtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileOutput).ToManagedClusterLoadBalancerProfilePtrOutputWithContext(ctx)
}









type ManagedClusterLoadBalancerProfilePtrInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfilePtrOutput() ManagedClusterLoadBalancerProfilePtrOutput
	ToManagedClusterLoadBalancerProfilePtrOutputWithContext(context.Context) ManagedClusterLoadBalancerProfilePtrOutput
}

type managedClusterLoadBalancerProfilePtrType ManagedClusterLoadBalancerProfileArgs

func ManagedClusterLoadBalancerProfilePtr(v *ManagedClusterLoadBalancerProfileArgs) ManagedClusterLoadBalancerProfilePtrInput {
	return (*managedClusterLoadBalancerProfilePtrType)(v)
}

func (*managedClusterLoadBalancerProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfile)(nil)).Elem()
}

func (i *managedClusterLoadBalancerProfilePtrType) ToManagedClusterLoadBalancerProfilePtrOutput() ManagedClusterLoadBalancerProfilePtrOutput {
	return i.ToManagedClusterLoadBalancerProfilePtrOutputWithContext(context.Background())
}

func (i *managedClusterLoadBalancerProfilePtrType) ToManagedClusterLoadBalancerProfilePtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfilePtrOutput)
}

type ManagedClusterLoadBalancerProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfile)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileOutput) ToManagedClusterLoadBalancerProfileOutput() ManagedClusterLoadBalancerProfileOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileOutput) ToManagedClusterLoadBalancerProfileOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileOutput) ToManagedClusterLoadBalancerProfilePtrOutput() ManagedClusterLoadBalancerProfilePtrOutput {
	return o.ToManagedClusterLoadBalancerProfilePtrOutputWithContext(context.Background())
}

func (o ManagedClusterLoadBalancerProfileOutput) ToManagedClusterLoadBalancerProfilePtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterLoadBalancerProfile) *ManagedClusterLoadBalancerProfile {
		return &v
	}).(ManagedClusterLoadBalancerProfilePtrOutput)
}

func (o ManagedClusterLoadBalancerProfileOutput) AllocatedOutboundPorts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfile) *int { return v.AllocatedOutboundPorts }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileOutput) EffectiveOutboundIPs() ResourceReferenceArrayOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfile) []ResourceReference { return v.EffectiveOutboundIPs }).(ResourceReferenceArrayOutput)
}

func (o ManagedClusterLoadBalancerProfileOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfile) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileOutput) ManagedOutboundIPs() ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfile) *ManagedClusterLoadBalancerProfileManagedOutboundIPs {
		return v.ManagedOutboundIPs
	}).(ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileOutput) OutboundIPPrefixes() ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfile) *ManagedClusterLoadBalancerProfileOutboundIPPrefixes {
		return v.OutboundIPPrefixes
	}).(ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileOutput) OutboundIPs() ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfile) *ManagedClusterLoadBalancerProfileOutboundIPs {
		return v.OutboundIPs
	}).(ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput)
}

type ManagedClusterLoadBalancerProfilePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfile)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfilePtrOutput) ToManagedClusterLoadBalancerProfilePtrOutput() ManagedClusterLoadBalancerProfilePtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfilePtrOutput) ToManagedClusterLoadBalancerProfilePtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfilePtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfilePtrOutput) Elem() ManagedClusterLoadBalancerProfileOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfile) ManagedClusterLoadBalancerProfile {
		if v != nil {
			return *v
		}
		var ret ManagedClusterLoadBalancerProfile
		return ret
	}).(ManagedClusterLoadBalancerProfileOutput)
}

func (o ManagedClusterLoadBalancerProfilePtrOutput) AllocatedOutboundPorts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfile) *int {
		if v == nil {
			return nil
		}
		return v.AllocatedOutboundPorts
	}).(pulumi.IntPtrOutput)
}

func (o ManagedClusterLoadBalancerProfilePtrOutput) EffectiveOutboundIPs() ResourceReferenceArrayOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfile) []ResourceReference {
		if v == nil {
			return nil
		}
		return v.EffectiveOutboundIPs
	}).(ResourceReferenceArrayOutput)
}

func (o ManagedClusterLoadBalancerProfilePtrOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfile) *int {
		if v == nil {
			return nil
		}
		return v.IdleTimeoutInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o ManagedClusterLoadBalancerProfilePtrOutput) ManagedOutboundIPs() ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfile) *ManagedClusterLoadBalancerProfileManagedOutboundIPs {
		if v == nil {
			return nil
		}
		return v.ManagedOutboundIPs
	}).(ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput)
}

func (o ManagedClusterLoadBalancerProfilePtrOutput) OutboundIPPrefixes() ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfile) *ManagedClusterLoadBalancerProfileOutboundIPPrefixes {
		if v == nil {
			return nil
		}
		return v.OutboundIPPrefixes
	}).(ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput)
}

func (o ManagedClusterLoadBalancerProfilePtrOutput) OutboundIPs() ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfile) *ManagedClusterLoadBalancerProfileOutboundIPs {
		if v == nil {
			return nil
		}
		return v.OutboundIPs
	}).(ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput)
}

type ManagedClusterLoadBalancerProfileManagedOutboundIPs struct {
	Count *int `pulumi:"count"`
}


func (val *ManagedClusterLoadBalancerProfileManagedOutboundIPs) Defaults() *ManagedClusterLoadBalancerProfileManagedOutboundIPs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		count_ := 1
		tmp.Count = &count_
	}
	return &tmp
}





type ManagedClusterLoadBalancerProfileManagedOutboundIPsInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileManagedOutboundIPsOutput() ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput
	ToManagedClusterLoadBalancerProfileManagedOutboundIPsOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput
}

type ManagedClusterLoadBalancerProfileManagedOutboundIPsArgs struct {
	Count pulumi.IntPtrInput `pulumi:"count"`
}

func (ManagedClusterLoadBalancerProfileManagedOutboundIPsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileManagedOutboundIPs)(nil)).Elem()
}

func (i ManagedClusterLoadBalancerProfileManagedOutboundIPsArgs) ToManagedClusterLoadBalancerProfileManagedOutboundIPsOutput() ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput {
	return i.ToManagedClusterLoadBalancerProfileManagedOutboundIPsOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileManagedOutboundIPsArgs) ToManagedClusterLoadBalancerProfileManagedOutboundIPsOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput)
}

func (i ManagedClusterLoadBalancerProfileManagedOutboundIPsArgs) ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileManagedOutboundIPsArgs) ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput).ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutputWithContext(ctx)
}









type ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput
	ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput
}

type managedClusterLoadBalancerProfileManagedOutboundIPsPtrType ManagedClusterLoadBalancerProfileManagedOutboundIPsArgs

func ManagedClusterLoadBalancerProfileManagedOutboundIPsPtr(v *ManagedClusterLoadBalancerProfileManagedOutboundIPsArgs) ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrInput {
	return (*managedClusterLoadBalancerProfileManagedOutboundIPsPtrType)(v)
}

func (*managedClusterLoadBalancerProfileManagedOutboundIPsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileManagedOutboundIPs)(nil)).Elem()
}

func (i *managedClusterLoadBalancerProfileManagedOutboundIPsPtrType) ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutputWithContext(context.Background())
}

func (i *managedClusterLoadBalancerProfileManagedOutboundIPsPtrType) ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput)
}

type ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileManagedOutboundIPs)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput) ToManagedClusterLoadBalancerProfileManagedOutboundIPsOutput() ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput) ToManagedClusterLoadBalancerProfileManagedOutboundIPsOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput) ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput {
	return o.ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutputWithContext(context.Background())
}

func (o ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput) ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterLoadBalancerProfileManagedOutboundIPs) *ManagedClusterLoadBalancerProfileManagedOutboundIPs {
		return &v
	}).(ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileManagedOutboundIPs) *int { return v.Count }).(pulumi.IntPtrOutput)
}

type ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileManagedOutboundIPs)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput) ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput) ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput) Elem() ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileManagedOutboundIPs) ManagedClusterLoadBalancerProfileManagedOutboundIPs {
		if v != nil {
			return *v
		}
		var ret ManagedClusterLoadBalancerProfileManagedOutboundIPs
		return ret
	}).(ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput)
}

func (o ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileManagedOutboundIPs) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

type ManagedClusterLoadBalancerProfileOutboundIPPrefixes struct {
	PublicIPPrefixes []ResourceReference `pulumi:"publicIPPrefixes"`
}





type ManagedClusterLoadBalancerProfileOutboundIPPrefixesInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput() ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput
	ToManagedClusterLoadBalancerProfileOutboundIPPrefixesOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput
}

type ManagedClusterLoadBalancerProfileOutboundIPPrefixesArgs struct {
	PublicIPPrefixes ResourceReferenceArrayInput `pulumi:"publicIPPrefixes"`
}

func (ManagedClusterLoadBalancerProfileOutboundIPPrefixesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileOutboundIPPrefixes)(nil)).Elem()
}

func (i ManagedClusterLoadBalancerProfileOutboundIPPrefixesArgs) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput() ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput {
	return i.ToManagedClusterLoadBalancerProfileOutboundIPPrefixesOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileOutboundIPPrefixesArgs) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput)
}

func (i ManagedClusterLoadBalancerProfileOutboundIPPrefixesArgs) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput() ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileOutboundIPPrefixesArgs) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput).ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutputWithContext(ctx)
}









type ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput() ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput
	ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput
}

type managedClusterLoadBalancerProfileOutboundIPPrefixesPtrType ManagedClusterLoadBalancerProfileOutboundIPPrefixesArgs

func ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtr(v *ManagedClusterLoadBalancerProfileOutboundIPPrefixesArgs) ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrInput {
	return (*managedClusterLoadBalancerProfileOutboundIPPrefixesPtrType)(v)
}

func (*managedClusterLoadBalancerProfileOutboundIPPrefixesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileOutboundIPPrefixes)(nil)).Elem()
}

func (i *managedClusterLoadBalancerProfileOutboundIPPrefixesPtrType) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput() ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutputWithContext(context.Background())
}

func (i *managedClusterLoadBalancerProfileOutboundIPPrefixesPtrType) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput)
}

type ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileOutboundIPPrefixes)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput() ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput() ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput {
	return o.ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutputWithContext(context.Background())
}

func (o ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterLoadBalancerProfileOutboundIPPrefixes) *ManagedClusterLoadBalancerProfileOutboundIPPrefixes {
		return &v
	}).(ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput) PublicIPPrefixes() ResourceReferenceArrayOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileOutboundIPPrefixes) []ResourceReference {
		return v.PublicIPPrefixes
	}).(ResourceReferenceArrayOutput)
}

type ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileOutboundIPPrefixes)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput() ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput) Elem() ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileOutboundIPPrefixes) ManagedClusterLoadBalancerProfileOutboundIPPrefixes {
		if v != nil {
			return *v
		}
		var ret ManagedClusterLoadBalancerProfileOutboundIPPrefixes
		return ret
	}).(ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput)
}

func (o ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput) PublicIPPrefixes() ResourceReferenceArrayOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileOutboundIPPrefixes) []ResourceReference {
		if v == nil {
			return nil
		}
		return v.PublicIPPrefixes
	}).(ResourceReferenceArrayOutput)
}

type ManagedClusterLoadBalancerProfileOutboundIPs struct {
	PublicIPs []ResourceReference `pulumi:"publicIPs"`
}





type ManagedClusterLoadBalancerProfileOutboundIPsInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileOutboundIPsOutput() ManagedClusterLoadBalancerProfileOutboundIPsOutput
	ToManagedClusterLoadBalancerProfileOutboundIPsOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileOutboundIPsOutput
}

type ManagedClusterLoadBalancerProfileOutboundIPsArgs struct {
	PublicIPs ResourceReferenceArrayInput `pulumi:"publicIPs"`
}

func (ManagedClusterLoadBalancerProfileOutboundIPsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileOutboundIPs)(nil)).Elem()
}

func (i ManagedClusterLoadBalancerProfileOutboundIPsArgs) ToManagedClusterLoadBalancerProfileOutboundIPsOutput() ManagedClusterLoadBalancerProfileOutboundIPsOutput {
	return i.ToManagedClusterLoadBalancerProfileOutboundIPsOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileOutboundIPsArgs) ToManagedClusterLoadBalancerProfileOutboundIPsOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileOutboundIPsOutput)
}

func (i ManagedClusterLoadBalancerProfileOutboundIPsArgs) ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileOutboundIPsArgs) ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileOutboundIPsOutput).ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutputWithContext(ctx)
}









type ManagedClusterLoadBalancerProfileOutboundIPsPtrInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput
	ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput
}

type managedClusterLoadBalancerProfileOutboundIPsPtrType ManagedClusterLoadBalancerProfileOutboundIPsArgs

func ManagedClusterLoadBalancerProfileOutboundIPsPtr(v *ManagedClusterLoadBalancerProfileOutboundIPsArgs) ManagedClusterLoadBalancerProfileOutboundIPsPtrInput {
	return (*managedClusterLoadBalancerProfileOutboundIPsPtrType)(v)
}

func (*managedClusterLoadBalancerProfileOutboundIPsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileOutboundIPs)(nil)).Elem()
}

func (i *managedClusterLoadBalancerProfileOutboundIPsPtrType) ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutputWithContext(context.Background())
}

func (i *managedClusterLoadBalancerProfileOutboundIPsPtrType) ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput)
}

type ManagedClusterLoadBalancerProfileOutboundIPsOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileOutboundIPsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileOutboundIPs)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileOutboundIPsOutput) ToManagedClusterLoadBalancerProfileOutboundIPsOutput() ManagedClusterLoadBalancerProfileOutboundIPsOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileOutboundIPsOutput) ToManagedClusterLoadBalancerProfileOutboundIPsOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPsOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileOutboundIPsOutput) ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput {
	return o.ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutputWithContext(context.Background())
}

func (o ManagedClusterLoadBalancerProfileOutboundIPsOutput) ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterLoadBalancerProfileOutboundIPs) *ManagedClusterLoadBalancerProfileOutboundIPs {
		return &v
	}).(ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileOutboundIPsOutput) PublicIPs() ResourceReferenceArrayOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileOutboundIPs) []ResourceReference { return v.PublicIPs }).(ResourceReferenceArrayOutput)
}

type ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileOutboundIPs)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput) ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput) ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput) Elem() ManagedClusterLoadBalancerProfileOutboundIPsOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileOutboundIPs) ManagedClusterLoadBalancerProfileOutboundIPs {
		if v != nil {
			return *v
		}
		var ret ManagedClusterLoadBalancerProfileOutboundIPs
		return ret
	}).(ManagedClusterLoadBalancerProfileOutboundIPsOutput)
}

func (o ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput) PublicIPs() ResourceReferenceArrayOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileOutboundIPs) []ResourceReference {
		if v == nil {
			return nil
		}
		return v.PublicIPs
	}).(ResourceReferenceArrayOutput)
}

type ManagedClusterLoadBalancerProfileResponse struct {
	AllocatedOutboundPorts *int                                                         `pulumi:"allocatedOutboundPorts"`
	EffectiveOutboundIPs   []ResourceReferenceResponse                                  `pulumi:"effectiveOutboundIPs"`
	IdleTimeoutInMinutes   *int                                                         `pulumi:"idleTimeoutInMinutes"`
	ManagedOutboundIPs     *ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs `pulumi:"managedOutboundIPs"`
	OutboundIPPrefixes     *ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes `pulumi:"outboundIPPrefixes"`
	OutboundIPs            *ManagedClusterLoadBalancerProfileResponseOutboundIPs        `pulumi:"outboundIPs"`
}


func (val *ManagedClusterLoadBalancerProfileResponse) Defaults() *ManagedClusterLoadBalancerProfileResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllocatedOutboundPorts) {
		allocatedOutboundPorts_ := 0
		tmp.AllocatedOutboundPorts = &allocatedOutboundPorts_
	}
	if isZero(tmp.IdleTimeoutInMinutes) {
		idleTimeoutInMinutes_ := 30
		tmp.IdleTimeoutInMinutes = &idleTimeoutInMinutes_
	}
	tmp.ManagedOutboundIPs = tmp.ManagedOutboundIPs.Defaults()

	return &tmp
}





type ManagedClusterLoadBalancerProfileResponseInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileResponseOutput() ManagedClusterLoadBalancerProfileResponseOutput
	ToManagedClusterLoadBalancerProfileResponseOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileResponseOutput
}

type ManagedClusterLoadBalancerProfileResponseArgs struct {
	AllocatedOutboundPorts pulumi.IntPtrInput                                                  `pulumi:"allocatedOutboundPorts"`
	EffectiveOutboundIPs   ResourceReferenceResponseArrayInput                                 `pulumi:"effectiveOutboundIPs"`
	IdleTimeoutInMinutes   pulumi.IntPtrInput                                                  `pulumi:"idleTimeoutInMinutes"`
	ManagedOutboundIPs     ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrInput `pulumi:"managedOutboundIPs"`
	OutboundIPPrefixes     ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrInput `pulumi:"outboundIPPrefixes"`
	OutboundIPs            ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrInput        `pulumi:"outboundIPs"`
}

func (ManagedClusterLoadBalancerProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileResponse)(nil)).Elem()
}

func (i ManagedClusterLoadBalancerProfileResponseArgs) ToManagedClusterLoadBalancerProfileResponseOutput() ManagedClusterLoadBalancerProfileResponseOutput {
	return i.ToManagedClusterLoadBalancerProfileResponseOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileResponseArgs) ToManagedClusterLoadBalancerProfileResponseOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseOutput)
}

func (i ManagedClusterLoadBalancerProfileResponseArgs) ToManagedClusterLoadBalancerProfileResponsePtrOutput() ManagedClusterLoadBalancerProfileResponsePtrOutput {
	return i.ToManagedClusterLoadBalancerProfileResponsePtrOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileResponseArgs) ToManagedClusterLoadBalancerProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseOutput).ToManagedClusterLoadBalancerProfileResponsePtrOutputWithContext(ctx)
}









type ManagedClusterLoadBalancerProfileResponsePtrInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileResponsePtrOutput() ManagedClusterLoadBalancerProfileResponsePtrOutput
	ToManagedClusterLoadBalancerProfileResponsePtrOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileResponsePtrOutput
}

type managedClusterLoadBalancerProfileResponsePtrType ManagedClusterLoadBalancerProfileResponseArgs

func ManagedClusterLoadBalancerProfileResponsePtr(v *ManagedClusterLoadBalancerProfileResponseArgs) ManagedClusterLoadBalancerProfileResponsePtrInput {
	return (*managedClusterLoadBalancerProfileResponsePtrType)(v)
}

func (*managedClusterLoadBalancerProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileResponse)(nil)).Elem()
}

func (i *managedClusterLoadBalancerProfileResponsePtrType) ToManagedClusterLoadBalancerProfileResponsePtrOutput() ManagedClusterLoadBalancerProfileResponsePtrOutput {
	return i.ToManagedClusterLoadBalancerProfileResponsePtrOutputWithContext(context.Background())
}

func (i *managedClusterLoadBalancerProfileResponsePtrType) ToManagedClusterLoadBalancerProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponsePtrOutput)
}

type ManagedClusterLoadBalancerProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileResponse)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileResponseOutput) ToManagedClusterLoadBalancerProfileResponseOutput() ManagedClusterLoadBalancerProfileResponseOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseOutput) ToManagedClusterLoadBalancerProfileResponseOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseOutput) ToManagedClusterLoadBalancerProfileResponsePtrOutput() ManagedClusterLoadBalancerProfileResponsePtrOutput {
	return o.ToManagedClusterLoadBalancerProfileResponsePtrOutputWithContext(context.Background())
}

func (o ManagedClusterLoadBalancerProfileResponseOutput) ToManagedClusterLoadBalancerProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterLoadBalancerProfileResponse) *ManagedClusterLoadBalancerProfileResponse {
		return &v
	}).(ManagedClusterLoadBalancerProfileResponsePtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseOutput) AllocatedOutboundPorts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileResponse) *int { return v.AllocatedOutboundPorts }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseOutput) EffectiveOutboundIPs() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileResponse) []ResourceReferenceResponse {
		return v.EffectiveOutboundIPs
	}).(ResourceReferenceResponseArrayOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileResponse) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseOutput) ManagedOutboundIPs() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileResponse) *ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs {
		return v.ManagedOutboundIPs
	}).(ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseOutput) OutboundIPPrefixes() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileResponse) *ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes {
		return v.OutboundIPPrefixes
	}).(ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseOutput) OutboundIPs() ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileResponse) *ManagedClusterLoadBalancerProfileResponseOutboundIPs {
		return v.OutboundIPs
	}).(ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput)
}

type ManagedClusterLoadBalancerProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileResponse)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileResponsePtrOutput) ToManagedClusterLoadBalancerProfileResponsePtrOutput() ManagedClusterLoadBalancerProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponsePtrOutput) ToManagedClusterLoadBalancerProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponsePtrOutput) Elem() ManagedClusterLoadBalancerProfileResponseOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponse) ManagedClusterLoadBalancerProfileResponse {
		if v != nil {
			return *v
		}
		var ret ManagedClusterLoadBalancerProfileResponse
		return ret
	}).(ManagedClusterLoadBalancerProfileResponseOutput)
}

func (o ManagedClusterLoadBalancerProfileResponsePtrOutput) AllocatedOutboundPorts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.AllocatedOutboundPorts
	}).(pulumi.IntPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponsePtrOutput) EffectiveOutboundIPs() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponse) []ResourceReferenceResponse {
		if v == nil {
			return nil
		}
		return v.EffectiveOutboundIPs
	}).(ResourceReferenceResponseArrayOutput)
}

func (o ManagedClusterLoadBalancerProfileResponsePtrOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.IdleTimeoutInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponsePtrOutput) ManagedOutboundIPs() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponse) *ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs {
		if v == nil {
			return nil
		}
		return v.ManagedOutboundIPs
	}).(ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponsePtrOutput) OutboundIPPrefixes() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponse) *ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes {
		if v == nil {
			return nil
		}
		return v.OutboundIPPrefixes
	}).(ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponsePtrOutput) OutboundIPs() ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponse) *ManagedClusterLoadBalancerProfileResponseOutboundIPs {
		if v == nil {
			return nil
		}
		return v.OutboundIPs
	}).(ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput)
}

type ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs struct {
	Count *int `pulumi:"count"`
}


func (val *ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs) Defaults() *ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		count_ := 1
		tmp.Count = &count_
	}
	return &tmp
}





type ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput
	ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput
}

type ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsArgs struct {
	Count pulumi.IntPtrInput `pulumi:"count"`
}

func (ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs)(nil)).Elem()
}

func (i ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsArgs) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput {
	return i.ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsArgs) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput)
}

func (i ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsArgs) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsArgs) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput).ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutputWithContext(ctx)
}









type ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput
	ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput
}

type managedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrType ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsArgs

func ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtr(v *ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsArgs) ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrInput {
	return (*managedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrType)(v)
}

func (*managedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs)(nil)).Elem()
}

func (i *managedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrType) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutputWithContext(context.Background())
}

func (i *managedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrType) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput)
}

type ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput {
	return o.ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutputWithContext(context.Background())
}

func (o ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs) *ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs {
		return &v
	}).(ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs) *int { return v.Count }).(pulumi.IntPtrOutput)
}

type ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput) Elem() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs) ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs {
		if v != nil {
			return *v
		}
		var ret ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs
		return ret
	}).(ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

type ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes struct {
	PublicIPPrefixes []ResourceReferenceResponse `pulumi:"publicIPPrefixes"`
}





type ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput
	ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput
}

type ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesArgs struct {
	PublicIPPrefixes ResourceReferenceResponseArrayInput `pulumi:"publicIPPrefixes"`
}

func (ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes)(nil)).Elem()
}

func (i ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesArgs) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput {
	return i.ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesArgs) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput)
}

func (i ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesArgs) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesArgs) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput).ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutputWithContext(ctx)
}









type ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput
	ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput
}

type managedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrType ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesArgs

func ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtr(v *ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesArgs) ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrInput {
	return (*managedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrType)(v)
}

func (*managedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes)(nil)).Elem()
}

func (i *managedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrType) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutputWithContext(context.Background())
}

func (i *managedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrType) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput)
}

type ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput {
	return o.ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutputWithContext(context.Background())
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes) *ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes {
		return &v
	}).(ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput) PublicIPPrefixes() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes) []ResourceReferenceResponse {
		return v.PublicIPPrefixes
	}).(ResourceReferenceResponseArrayOutput)
}

type ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput) Elem() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes) ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes {
		if v != nil {
			return *v
		}
		var ret ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes
		return ret
	}).(ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput) PublicIPPrefixes() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes) []ResourceReferenceResponse {
		if v == nil {
			return nil
		}
		return v.PublicIPPrefixes
	}).(ResourceReferenceResponseArrayOutput)
}

type ManagedClusterLoadBalancerProfileResponseOutboundIPs struct {
	PublicIPs []ResourceReferenceResponse `pulumi:"publicIPs"`
}





type ManagedClusterLoadBalancerProfileResponseOutboundIPsInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileResponseOutboundIPsOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput
	ToManagedClusterLoadBalancerProfileResponseOutboundIPsOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput
}

type ManagedClusterLoadBalancerProfileResponseOutboundIPsArgs struct {
	PublicIPs ResourceReferenceResponseArrayInput `pulumi:"publicIPs"`
}

func (ManagedClusterLoadBalancerProfileResponseOutboundIPsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileResponseOutboundIPs)(nil)).Elem()
}

func (i ManagedClusterLoadBalancerProfileResponseOutboundIPsArgs) ToManagedClusterLoadBalancerProfileResponseOutboundIPsOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput {
	return i.ToManagedClusterLoadBalancerProfileResponseOutboundIPsOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileResponseOutboundIPsArgs) ToManagedClusterLoadBalancerProfileResponseOutboundIPsOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput)
}

func (i ManagedClusterLoadBalancerProfileResponseOutboundIPsArgs) ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileResponseOutboundIPsArgs) ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput).ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutputWithContext(ctx)
}









type ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput
	ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput
}

type managedClusterLoadBalancerProfileResponseOutboundIPsPtrType ManagedClusterLoadBalancerProfileResponseOutboundIPsArgs

func ManagedClusterLoadBalancerProfileResponseOutboundIPsPtr(v *ManagedClusterLoadBalancerProfileResponseOutboundIPsArgs) ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrInput {
	return (*managedClusterLoadBalancerProfileResponseOutboundIPsPtrType)(v)
}

func (*managedClusterLoadBalancerProfileResponseOutboundIPsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileResponseOutboundIPs)(nil)).Elem()
}

func (i *managedClusterLoadBalancerProfileResponseOutboundIPsPtrType) ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutputWithContext(context.Background())
}

func (i *managedClusterLoadBalancerProfileResponseOutboundIPsPtrType) ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput)
}

type ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileResponseOutboundIPs)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPsOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPsOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput {
	return o.ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutputWithContext(context.Background())
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterLoadBalancerProfileResponseOutboundIPs) *ManagedClusterLoadBalancerProfileResponseOutboundIPs {
		return &v
	}).(ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput) PublicIPs() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileResponseOutboundIPs) []ResourceReferenceResponse {
		return v.PublicIPs
	}).(ResourceReferenceResponseArrayOutput)
}

type ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileResponseOutboundIPs)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput) Elem() ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponseOutboundIPs) ManagedClusterLoadBalancerProfileResponseOutboundIPs {
		if v != nil {
			return *v
		}
		var ret ManagedClusterLoadBalancerProfileResponseOutboundIPs
		return ret
	}).(ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput) PublicIPs() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponseOutboundIPs) []ResourceReferenceResponse {
		if v == nil {
			return nil
		}
		return v.PublicIPs
	}).(ResourceReferenceResponseArrayOutput)
}

type ManagedClusterPodIdentity struct {
	BindingSelector *string              `pulumi:"bindingSelector"`
	Identity        UserAssignedIdentity `pulumi:"identity"`
	Name            string               `pulumi:"name"`
	Namespace       string               `pulumi:"namespace"`
}





type ManagedClusterPodIdentityInput interface {
	pulumi.Input

	ToManagedClusterPodIdentityOutput() ManagedClusterPodIdentityOutput
	ToManagedClusterPodIdentityOutputWithContext(context.Context) ManagedClusterPodIdentityOutput
}

type ManagedClusterPodIdentityArgs struct {
	BindingSelector pulumi.StringPtrInput     `pulumi:"bindingSelector"`
	Identity        UserAssignedIdentityInput `pulumi:"identity"`
	Name            pulumi.StringInput        `pulumi:"name"`
	Namespace       pulumi.StringInput        `pulumi:"namespace"`
}

func (ManagedClusterPodIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPodIdentity)(nil)).Elem()
}

func (i ManagedClusterPodIdentityArgs) ToManagedClusterPodIdentityOutput() ManagedClusterPodIdentityOutput {
	return i.ToManagedClusterPodIdentityOutputWithContext(context.Background())
}

func (i ManagedClusterPodIdentityArgs) ToManagedClusterPodIdentityOutputWithContext(ctx context.Context) ManagedClusterPodIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPodIdentityOutput)
}





type ManagedClusterPodIdentityArrayInput interface {
	pulumi.Input

	ToManagedClusterPodIdentityArrayOutput() ManagedClusterPodIdentityArrayOutput
	ToManagedClusterPodIdentityArrayOutputWithContext(context.Context) ManagedClusterPodIdentityArrayOutput
}

type ManagedClusterPodIdentityArray []ManagedClusterPodIdentityInput

func (ManagedClusterPodIdentityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedClusterPodIdentity)(nil)).Elem()
}

func (i ManagedClusterPodIdentityArray) ToManagedClusterPodIdentityArrayOutput() ManagedClusterPodIdentityArrayOutput {
	return i.ToManagedClusterPodIdentityArrayOutputWithContext(context.Background())
}

func (i ManagedClusterPodIdentityArray) ToManagedClusterPodIdentityArrayOutputWithContext(ctx context.Context) ManagedClusterPodIdentityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPodIdentityArrayOutput)
}

type ManagedClusterPodIdentityOutput struct{ *pulumi.OutputState }

func (ManagedClusterPodIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPodIdentity)(nil)).Elem()
}

func (o ManagedClusterPodIdentityOutput) ToManagedClusterPodIdentityOutput() ManagedClusterPodIdentityOutput {
	return o
}

func (o ManagedClusterPodIdentityOutput) ToManagedClusterPodIdentityOutputWithContext(ctx context.Context) ManagedClusterPodIdentityOutput {
	return o
}

func (o ManagedClusterPodIdentityOutput) BindingSelector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentity) *string { return v.BindingSelector }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPodIdentityOutput) Identity() UserAssignedIdentityOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentity) UserAssignedIdentity { return v.Identity }).(UserAssignedIdentityOutput)
}

func (o ManagedClusterPodIdentityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentity) string { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedClusterPodIdentityOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentity) string { return v.Namespace }).(pulumi.StringOutput)
}

type ManagedClusterPodIdentityArrayOutput struct{ *pulumi.OutputState }

func (ManagedClusterPodIdentityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedClusterPodIdentity)(nil)).Elem()
}

func (o ManagedClusterPodIdentityArrayOutput) ToManagedClusterPodIdentityArrayOutput() ManagedClusterPodIdentityArrayOutput {
	return o
}

func (o ManagedClusterPodIdentityArrayOutput) ToManagedClusterPodIdentityArrayOutputWithContext(ctx context.Context) ManagedClusterPodIdentityArrayOutput {
	return o
}

func (o ManagedClusterPodIdentityArrayOutput) Index(i pulumi.IntInput) ManagedClusterPodIdentityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedClusterPodIdentity {
		return vs[0].([]ManagedClusterPodIdentity)[vs[1].(int)]
	}).(ManagedClusterPodIdentityOutput)
}

type ManagedClusterPodIdentityException struct {
	Name      string            `pulumi:"name"`
	Namespace string            `pulumi:"namespace"`
	PodLabels map[string]string `pulumi:"podLabels"`
}





type ManagedClusterPodIdentityExceptionInput interface {
	pulumi.Input

	ToManagedClusterPodIdentityExceptionOutput() ManagedClusterPodIdentityExceptionOutput
	ToManagedClusterPodIdentityExceptionOutputWithContext(context.Context) ManagedClusterPodIdentityExceptionOutput
}

type ManagedClusterPodIdentityExceptionArgs struct {
	Name      pulumi.StringInput    `pulumi:"name"`
	Namespace pulumi.StringInput    `pulumi:"namespace"`
	PodLabels pulumi.StringMapInput `pulumi:"podLabels"`
}

func (ManagedClusterPodIdentityExceptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPodIdentityException)(nil)).Elem()
}

func (i ManagedClusterPodIdentityExceptionArgs) ToManagedClusterPodIdentityExceptionOutput() ManagedClusterPodIdentityExceptionOutput {
	return i.ToManagedClusterPodIdentityExceptionOutputWithContext(context.Background())
}

func (i ManagedClusterPodIdentityExceptionArgs) ToManagedClusterPodIdentityExceptionOutputWithContext(ctx context.Context) ManagedClusterPodIdentityExceptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPodIdentityExceptionOutput)
}





type ManagedClusterPodIdentityExceptionArrayInput interface {
	pulumi.Input

	ToManagedClusterPodIdentityExceptionArrayOutput() ManagedClusterPodIdentityExceptionArrayOutput
	ToManagedClusterPodIdentityExceptionArrayOutputWithContext(context.Context) ManagedClusterPodIdentityExceptionArrayOutput
}

type ManagedClusterPodIdentityExceptionArray []ManagedClusterPodIdentityExceptionInput

func (ManagedClusterPodIdentityExceptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedClusterPodIdentityException)(nil)).Elem()
}

func (i ManagedClusterPodIdentityExceptionArray) ToManagedClusterPodIdentityExceptionArrayOutput() ManagedClusterPodIdentityExceptionArrayOutput {
	return i.ToManagedClusterPodIdentityExceptionArrayOutputWithContext(context.Background())
}

func (i ManagedClusterPodIdentityExceptionArray) ToManagedClusterPodIdentityExceptionArrayOutputWithContext(ctx context.Context) ManagedClusterPodIdentityExceptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPodIdentityExceptionArrayOutput)
}

type ManagedClusterPodIdentityExceptionOutput struct{ *pulumi.OutputState }

func (ManagedClusterPodIdentityExceptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPodIdentityException)(nil)).Elem()
}

func (o ManagedClusterPodIdentityExceptionOutput) ToManagedClusterPodIdentityExceptionOutput() ManagedClusterPodIdentityExceptionOutput {
	return o
}

func (o ManagedClusterPodIdentityExceptionOutput) ToManagedClusterPodIdentityExceptionOutputWithContext(ctx context.Context) ManagedClusterPodIdentityExceptionOutput {
	return o
}

func (o ManagedClusterPodIdentityExceptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityException) string { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedClusterPodIdentityExceptionOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityException) string { return v.Namespace }).(pulumi.StringOutput)
}

func (o ManagedClusterPodIdentityExceptionOutput) PodLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityException) map[string]string { return v.PodLabels }).(pulumi.StringMapOutput)
}

type ManagedClusterPodIdentityExceptionArrayOutput struct{ *pulumi.OutputState }

func (ManagedClusterPodIdentityExceptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedClusterPodIdentityException)(nil)).Elem()
}

func (o ManagedClusterPodIdentityExceptionArrayOutput) ToManagedClusterPodIdentityExceptionArrayOutput() ManagedClusterPodIdentityExceptionArrayOutput {
	return o
}

func (o ManagedClusterPodIdentityExceptionArrayOutput) ToManagedClusterPodIdentityExceptionArrayOutputWithContext(ctx context.Context) ManagedClusterPodIdentityExceptionArrayOutput {
	return o
}

func (o ManagedClusterPodIdentityExceptionArrayOutput) Index(i pulumi.IntInput) ManagedClusterPodIdentityExceptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedClusterPodIdentityException {
		return vs[0].([]ManagedClusterPodIdentityException)[vs[1].(int)]
	}).(ManagedClusterPodIdentityExceptionOutput)
}

type ManagedClusterPodIdentityExceptionResponse struct {
	Name      string            `pulumi:"name"`
	Namespace string            `pulumi:"namespace"`
	PodLabels map[string]string `pulumi:"podLabels"`
}





type ManagedClusterPodIdentityExceptionResponseInput interface {
	pulumi.Input

	ToManagedClusterPodIdentityExceptionResponseOutput() ManagedClusterPodIdentityExceptionResponseOutput
	ToManagedClusterPodIdentityExceptionResponseOutputWithContext(context.Context) ManagedClusterPodIdentityExceptionResponseOutput
}

type ManagedClusterPodIdentityExceptionResponseArgs struct {
	Name      pulumi.StringInput    `pulumi:"name"`
	Namespace pulumi.StringInput    `pulumi:"namespace"`
	PodLabels pulumi.StringMapInput `pulumi:"podLabels"`
}

func (ManagedClusterPodIdentityExceptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPodIdentityExceptionResponse)(nil)).Elem()
}

func (i ManagedClusterPodIdentityExceptionResponseArgs) ToManagedClusterPodIdentityExceptionResponseOutput() ManagedClusterPodIdentityExceptionResponseOutput {
	return i.ToManagedClusterPodIdentityExceptionResponseOutputWithContext(context.Background())
}

func (i ManagedClusterPodIdentityExceptionResponseArgs) ToManagedClusterPodIdentityExceptionResponseOutputWithContext(ctx context.Context) ManagedClusterPodIdentityExceptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPodIdentityExceptionResponseOutput)
}





type ManagedClusterPodIdentityExceptionResponseArrayInput interface {
	pulumi.Input

	ToManagedClusterPodIdentityExceptionResponseArrayOutput() ManagedClusterPodIdentityExceptionResponseArrayOutput
	ToManagedClusterPodIdentityExceptionResponseArrayOutputWithContext(context.Context) ManagedClusterPodIdentityExceptionResponseArrayOutput
}

type ManagedClusterPodIdentityExceptionResponseArray []ManagedClusterPodIdentityExceptionResponseInput

func (ManagedClusterPodIdentityExceptionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedClusterPodIdentityExceptionResponse)(nil)).Elem()
}

func (i ManagedClusterPodIdentityExceptionResponseArray) ToManagedClusterPodIdentityExceptionResponseArrayOutput() ManagedClusterPodIdentityExceptionResponseArrayOutput {
	return i.ToManagedClusterPodIdentityExceptionResponseArrayOutputWithContext(context.Background())
}

func (i ManagedClusterPodIdentityExceptionResponseArray) ToManagedClusterPodIdentityExceptionResponseArrayOutputWithContext(ctx context.Context) ManagedClusterPodIdentityExceptionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPodIdentityExceptionResponseArrayOutput)
}

type ManagedClusterPodIdentityExceptionResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterPodIdentityExceptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPodIdentityExceptionResponse)(nil)).Elem()
}

func (o ManagedClusterPodIdentityExceptionResponseOutput) ToManagedClusterPodIdentityExceptionResponseOutput() ManagedClusterPodIdentityExceptionResponseOutput {
	return o
}

func (o ManagedClusterPodIdentityExceptionResponseOutput) ToManagedClusterPodIdentityExceptionResponseOutputWithContext(ctx context.Context) ManagedClusterPodIdentityExceptionResponseOutput {
	return o
}

func (o ManagedClusterPodIdentityExceptionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityExceptionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedClusterPodIdentityExceptionResponseOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityExceptionResponse) string { return v.Namespace }).(pulumi.StringOutput)
}

func (o ManagedClusterPodIdentityExceptionResponseOutput) PodLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityExceptionResponse) map[string]string { return v.PodLabels }).(pulumi.StringMapOutput)
}

type ManagedClusterPodIdentityExceptionResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagedClusterPodIdentityExceptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedClusterPodIdentityExceptionResponse)(nil)).Elem()
}

func (o ManagedClusterPodIdentityExceptionResponseArrayOutput) ToManagedClusterPodIdentityExceptionResponseArrayOutput() ManagedClusterPodIdentityExceptionResponseArrayOutput {
	return o
}

func (o ManagedClusterPodIdentityExceptionResponseArrayOutput) ToManagedClusterPodIdentityExceptionResponseArrayOutputWithContext(ctx context.Context) ManagedClusterPodIdentityExceptionResponseArrayOutput {
	return o
}

func (o ManagedClusterPodIdentityExceptionResponseArrayOutput) Index(i pulumi.IntInput) ManagedClusterPodIdentityExceptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedClusterPodIdentityExceptionResponse {
		return vs[0].([]ManagedClusterPodIdentityExceptionResponse)[vs[1].(int)]
	}).(ManagedClusterPodIdentityExceptionResponseOutput)
}

type ManagedClusterPodIdentityProfile struct {
	AllowNetworkPluginKubenet      *bool                                `pulumi:"allowNetworkPluginKubenet"`
	Enabled                        *bool                                `pulumi:"enabled"`
	UserAssignedIdentities         []ManagedClusterPodIdentity          `pulumi:"userAssignedIdentities"`
	UserAssignedIdentityExceptions []ManagedClusterPodIdentityException `pulumi:"userAssignedIdentityExceptions"`
}





type ManagedClusterPodIdentityProfileInput interface {
	pulumi.Input

	ToManagedClusterPodIdentityProfileOutput() ManagedClusterPodIdentityProfileOutput
	ToManagedClusterPodIdentityProfileOutputWithContext(context.Context) ManagedClusterPodIdentityProfileOutput
}

type ManagedClusterPodIdentityProfileArgs struct {
	AllowNetworkPluginKubenet      pulumi.BoolPtrInput                          `pulumi:"allowNetworkPluginKubenet"`
	Enabled                        pulumi.BoolPtrInput                          `pulumi:"enabled"`
	UserAssignedIdentities         ManagedClusterPodIdentityArrayInput          `pulumi:"userAssignedIdentities"`
	UserAssignedIdentityExceptions ManagedClusterPodIdentityExceptionArrayInput `pulumi:"userAssignedIdentityExceptions"`
}

func (ManagedClusterPodIdentityProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPodIdentityProfile)(nil)).Elem()
}

func (i ManagedClusterPodIdentityProfileArgs) ToManagedClusterPodIdentityProfileOutput() ManagedClusterPodIdentityProfileOutput {
	return i.ToManagedClusterPodIdentityProfileOutputWithContext(context.Background())
}

func (i ManagedClusterPodIdentityProfileArgs) ToManagedClusterPodIdentityProfileOutputWithContext(ctx context.Context) ManagedClusterPodIdentityProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPodIdentityProfileOutput)
}

func (i ManagedClusterPodIdentityProfileArgs) ToManagedClusterPodIdentityProfilePtrOutput() ManagedClusterPodIdentityProfilePtrOutput {
	return i.ToManagedClusterPodIdentityProfilePtrOutputWithContext(context.Background())
}

func (i ManagedClusterPodIdentityProfileArgs) ToManagedClusterPodIdentityProfilePtrOutputWithContext(ctx context.Context) ManagedClusterPodIdentityProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPodIdentityProfileOutput).ToManagedClusterPodIdentityProfilePtrOutputWithContext(ctx)
}









type ManagedClusterPodIdentityProfilePtrInput interface {
	pulumi.Input

	ToManagedClusterPodIdentityProfilePtrOutput() ManagedClusterPodIdentityProfilePtrOutput
	ToManagedClusterPodIdentityProfilePtrOutputWithContext(context.Context) ManagedClusterPodIdentityProfilePtrOutput
}

type managedClusterPodIdentityProfilePtrType ManagedClusterPodIdentityProfileArgs

func ManagedClusterPodIdentityProfilePtr(v *ManagedClusterPodIdentityProfileArgs) ManagedClusterPodIdentityProfilePtrInput {
	return (*managedClusterPodIdentityProfilePtrType)(v)
}

func (*managedClusterPodIdentityProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterPodIdentityProfile)(nil)).Elem()
}

func (i *managedClusterPodIdentityProfilePtrType) ToManagedClusterPodIdentityProfilePtrOutput() ManagedClusterPodIdentityProfilePtrOutput {
	return i.ToManagedClusterPodIdentityProfilePtrOutputWithContext(context.Background())
}

func (i *managedClusterPodIdentityProfilePtrType) ToManagedClusterPodIdentityProfilePtrOutputWithContext(ctx context.Context) ManagedClusterPodIdentityProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPodIdentityProfilePtrOutput)
}

type ManagedClusterPodIdentityProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterPodIdentityProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPodIdentityProfile)(nil)).Elem()
}

func (o ManagedClusterPodIdentityProfileOutput) ToManagedClusterPodIdentityProfileOutput() ManagedClusterPodIdentityProfileOutput {
	return o
}

func (o ManagedClusterPodIdentityProfileOutput) ToManagedClusterPodIdentityProfileOutputWithContext(ctx context.Context) ManagedClusterPodIdentityProfileOutput {
	return o
}

func (o ManagedClusterPodIdentityProfileOutput) ToManagedClusterPodIdentityProfilePtrOutput() ManagedClusterPodIdentityProfilePtrOutput {
	return o.ToManagedClusterPodIdentityProfilePtrOutputWithContext(context.Background())
}

func (o ManagedClusterPodIdentityProfileOutput) ToManagedClusterPodIdentityProfilePtrOutputWithContext(ctx context.Context) ManagedClusterPodIdentityProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterPodIdentityProfile) *ManagedClusterPodIdentityProfile {
		return &v
	}).(ManagedClusterPodIdentityProfilePtrOutput)
}

func (o ManagedClusterPodIdentityProfileOutput) AllowNetworkPluginKubenet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityProfile) *bool { return v.AllowNetworkPluginKubenet }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterPodIdentityProfileOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityProfile) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterPodIdentityProfileOutput) UserAssignedIdentities() ManagedClusterPodIdentityArrayOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityProfile) []ManagedClusterPodIdentity { return v.UserAssignedIdentities }).(ManagedClusterPodIdentityArrayOutput)
}

func (o ManagedClusterPodIdentityProfileOutput) UserAssignedIdentityExceptions() ManagedClusterPodIdentityExceptionArrayOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityProfile) []ManagedClusterPodIdentityException {
		return v.UserAssignedIdentityExceptions
	}).(ManagedClusterPodIdentityExceptionArrayOutput)
}

type ManagedClusterPodIdentityProfilePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterPodIdentityProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterPodIdentityProfile)(nil)).Elem()
}

func (o ManagedClusterPodIdentityProfilePtrOutput) ToManagedClusterPodIdentityProfilePtrOutput() ManagedClusterPodIdentityProfilePtrOutput {
	return o
}

func (o ManagedClusterPodIdentityProfilePtrOutput) ToManagedClusterPodIdentityProfilePtrOutputWithContext(ctx context.Context) ManagedClusterPodIdentityProfilePtrOutput {
	return o
}

func (o ManagedClusterPodIdentityProfilePtrOutput) Elem() ManagedClusterPodIdentityProfileOutput {
	return o.ApplyT(func(v *ManagedClusterPodIdentityProfile) ManagedClusterPodIdentityProfile {
		if v != nil {
			return *v
		}
		var ret ManagedClusterPodIdentityProfile
		return ret
	}).(ManagedClusterPodIdentityProfileOutput)
}

func (o ManagedClusterPodIdentityProfilePtrOutput) AllowNetworkPluginKubenet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPodIdentityProfile) *bool {
		if v == nil {
			return nil
		}
		return v.AllowNetworkPluginKubenet
	}).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterPodIdentityProfilePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPodIdentityProfile) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterPodIdentityProfilePtrOutput) UserAssignedIdentities() ManagedClusterPodIdentityArrayOutput {
	return o.ApplyT(func(v *ManagedClusterPodIdentityProfile) []ManagedClusterPodIdentity {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(ManagedClusterPodIdentityArrayOutput)
}

func (o ManagedClusterPodIdentityProfilePtrOutput) UserAssignedIdentityExceptions() ManagedClusterPodIdentityExceptionArrayOutput {
	return o.ApplyT(func(v *ManagedClusterPodIdentityProfile) []ManagedClusterPodIdentityException {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentityExceptions
	}).(ManagedClusterPodIdentityExceptionArrayOutput)
}

type ManagedClusterPodIdentityProfileResponse struct {
	AllowNetworkPluginKubenet      *bool                                        `pulumi:"allowNetworkPluginKubenet"`
	Enabled                        *bool                                        `pulumi:"enabled"`
	UserAssignedIdentities         []ManagedClusterPodIdentityResponse          `pulumi:"userAssignedIdentities"`
	UserAssignedIdentityExceptions []ManagedClusterPodIdentityExceptionResponse `pulumi:"userAssignedIdentityExceptions"`
}





type ManagedClusterPodIdentityProfileResponseInput interface {
	pulumi.Input

	ToManagedClusterPodIdentityProfileResponseOutput() ManagedClusterPodIdentityProfileResponseOutput
	ToManagedClusterPodIdentityProfileResponseOutputWithContext(context.Context) ManagedClusterPodIdentityProfileResponseOutput
}

type ManagedClusterPodIdentityProfileResponseArgs struct {
	AllowNetworkPluginKubenet      pulumi.BoolPtrInput                                  `pulumi:"allowNetworkPluginKubenet"`
	Enabled                        pulumi.BoolPtrInput                                  `pulumi:"enabled"`
	UserAssignedIdentities         ManagedClusterPodIdentityResponseArrayInput          `pulumi:"userAssignedIdentities"`
	UserAssignedIdentityExceptions ManagedClusterPodIdentityExceptionResponseArrayInput `pulumi:"userAssignedIdentityExceptions"`
}

func (ManagedClusterPodIdentityProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPodIdentityProfileResponse)(nil)).Elem()
}

func (i ManagedClusterPodIdentityProfileResponseArgs) ToManagedClusterPodIdentityProfileResponseOutput() ManagedClusterPodIdentityProfileResponseOutput {
	return i.ToManagedClusterPodIdentityProfileResponseOutputWithContext(context.Background())
}

func (i ManagedClusterPodIdentityProfileResponseArgs) ToManagedClusterPodIdentityProfileResponseOutputWithContext(ctx context.Context) ManagedClusterPodIdentityProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPodIdentityProfileResponseOutput)
}

func (i ManagedClusterPodIdentityProfileResponseArgs) ToManagedClusterPodIdentityProfileResponsePtrOutput() ManagedClusterPodIdentityProfileResponsePtrOutput {
	return i.ToManagedClusterPodIdentityProfileResponsePtrOutputWithContext(context.Background())
}

func (i ManagedClusterPodIdentityProfileResponseArgs) ToManagedClusterPodIdentityProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterPodIdentityProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPodIdentityProfileResponseOutput).ToManagedClusterPodIdentityProfileResponsePtrOutputWithContext(ctx)
}









type ManagedClusterPodIdentityProfileResponsePtrInput interface {
	pulumi.Input

	ToManagedClusterPodIdentityProfileResponsePtrOutput() ManagedClusterPodIdentityProfileResponsePtrOutput
	ToManagedClusterPodIdentityProfileResponsePtrOutputWithContext(context.Context) ManagedClusterPodIdentityProfileResponsePtrOutput
}

type managedClusterPodIdentityProfileResponsePtrType ManagedClusterPodIdentityProfileResponseArgs

func ManagedClusterPodIdentityProfileResponsePtr(v *ManagedClusterPodIdentityProfileResponseArgs) ManagedClusterPodIdentityProfileResponsePtrInput {
	return (*managedClusterPodIdentityProfileResponsePtrType)(v)
}

func (*managedClusterPodIdentityProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterPodIdentityProfileResponse)(nil)).Elem()
}

func (i *managedClusterPodIdentityProfileResponsePtrType) ToManagedClusterPodIdentityProfileResponsePtrOutput() ManagedClusterPodIdentityProfileResponsePtrOutput {
	return i.ToManagedClusterPodIdentityProfileResponsePtrOutputWithContext(context.Background())
}

func (i *managedClusterPodIdentityProfileResponsePtrType) ToManagedClusterPodIdentityProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterPodIdentityProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPodIdentityProfileResponsePtrOutput)
}

type ManagedClusterPodIdentityProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterPodIdentityProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPodIdentityProfileResponse)(nil)).Elem()
}

func (o ManagedClusterPodIdentityProfileResponseOutput) ToManagedClusterPodIdentityProfileResponseOutput() ManagedClusterPodIdentityProfileResponseOutput {
	return o
}

func (o ManagedClusterPodIdentityProfileResponseOutput) ToManagedClusterPodIdentityProfileResponseOutputWithContext(ctx context.Context) ManagedClusterPodIdentityProfileResponseOutput {
	return o
}

func (o ManagedClusterPodIdentityProfileResponseOutput) ToManagedClusterPodIdentityProfileResponsePtrOutput() ManagedClusterPodIdentityProfileResponsePtrOutput {
	return o.ToManagedClusterPodIdentityProfileResponsePtrOutputWithContext(context.Background())
}

func (o ManagedClusterPodIdentityProfileResponseOutput) ToManagedClusterPodIdentityProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterPodIdentityProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterPodIdentityProfileResponse) *ManagedClusterPodIdentityProfileResponse {
		return &v
	}).(ManagedClusterPodIdentityProfileResponsePtrOutput)
}

func (o ManagedClusterPodIdentityProfileResponseOutput) AllowNetworkPluginKubenet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityProfileResponse) *bool { return v.AllowNetworkPluginKubenet }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterPodIdentityProfileResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityProfileResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterPodIdentityProfileResponseOutput) UserAssignedIdentities() ManagedClusterPodIdentityResponseArrayOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityProfileResponse) []ManagedClusterPodIdentityResponse {
		return v.UserAssignedIdentities
	}).(ManagedClusterPodIdentityResponseArrayOutput)
}

func (o ManagedClusterPodIdentityProfileResponseOutput) UserAssignedIdentityExceptions() ManagedClusterPodIdentityExceptionResponseArrayOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityProfileResponse) []ManagedClusterPodIdentityExceptionResponse {
		return v.UserAssignedIdentityExceptions
	}).(ManagedClusterPodIdentityExceptionResponseArrayOutput)
}

type ManagedClusterPodIdentityProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterPodIdentityProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterPodIdentityProfileResponse)(nil)).Elem()
}

func (o ManagedClusterPodIdentityProfileResponsePtrOutput) ToManagedClusterPodIdentityProfileResponsePtrOutput() ManagedClusterPodIdentityProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterPodIdentityProfileResponsePtrOutput) ToManagedClusterPodIdentityProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterPodIdentityProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterPodIdentityProfileResponsePtrOutput) Elem() ManagedClusterPodIdentityProfileResponseOutput {
	return o.ApplyT(func(v *ManagedClusterPodIdentityProfileResponse) ManagedClusterPodIdentityProfileResponse {
		if v != nil {
			return *v
		}
		var ret ManagedClusterPodIdentityProfileResponse
		return ret
	}).(ManagedClusterPodIdentityProfileResponseOutput)
}

func (o ManagedClusterPodIdentityProfileResponsePtrOutput) AllowNetworkPluginKubenet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPodIdentityProfileResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AllowNetworkPluginKubenet
	}).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterPodIdentityProfileResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPodIdentityProfileResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterPodIdentityProfileResponsePtrOutput) UserAssignedIdentities() ManagedClusterPodIdentityResponseArrayOutput {
	return o.ApplyT(func(v *ManagedClusterPodIdentityProfileResponse) []ManagedClusterPodIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(ManagedClusterPodIdentityResponseArrayOutput)
}

func (o ManagedClusterPodIdentityProfileResponsePtrOutput) UserAssignedIdentityExceptions() ManagedClusterPodIdentityExceptionResponseArrayOutput {
	return o.ApplyT(func(v *ManagedClusterPodIdentityProfileResponse) []ManagedClusterPodIdentityExceptionResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentityExceptions
	}).(ManagedClusterPodIdentityExceptionResponseArrayOutput)
}

type ManagedClusterPodIdentityResponse struct {
	BindingSelector   *string                                           `pulumi:"bindingSelector"`
	Identity          UserAssignedIdentityResponse                      `pulumi:"identity"`
	Name              string                                            `pulumi:"name"`
	Namespace         string                                            `pulumi:"namespace"`
	ProvisioningInfo  ManagedClusterPodIdentityResponseProvisioningInfo `pulumi:"provisioningInfo"`
	ProvisioningState string                                            `pulumi:"provisioningState"`
}





type ManagedClusterPodIdentityResponseInput interface {
	pulumi.Input

	ToManagedClusterPodIdentityResponseOutput() ManagedClusterPodIdentityResponseOutput
	ToManagedClusterPodIdentityResponseOutputWithContext(context.Context) ManagedClusterPodIdentityResponseOutput
}

type ManagedClusterPodIdentityResponseArgs struct {
	BindingSelector   pulumi.StringPtrInput                                  `pulumi:"bindingSelector"`
	Identity          UserAssignedIdentityResponseInput                      `pulumi:"identity"`
	Name              pulumi.StringInput                                     `pulumi:"name"`
	Namespace         pulumi.StringInput                                     `pulumi:"namespace"`
	ProvisioningInfo  ManagedClusterPodIdentityResponseProvisioningInfoInput `pulumi:"provisioningInfo"`
	ProvisioningState pulumi.StringInput                                     `pulumi:"provisioningState"`
}

func (ManagedClusterPodIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPodIdentityResponse)(nil)).Elem()
}

func (i ManagedClusterPodIdentityResponseArgs) ToManagedClusterPodIdentityResponseOutput() ManagedClusterPodIdentityResponseOutput {
	return i.ToManagedClusterPodIdentityResponseOutputWithContext(context.Background())
}

func (i ManagedClusterPodIdentityResponseArgs) ToManagedClusterPodIdentityResponseOutputWithContext(ctx context.Context) ManagedClusterPodIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPodIdentityResponseOutput)
}





type ManagedClusterPodIdentityResponseArrayInput interface {
	pulumi.Input

	ToManagedClusterPodIdentityResponseArrayOutput() ManagedClusterPodIdentityResponseArrayOutput
	ToManagedClusterPodIdentityResponseArrayOutputWithContext(context.Context) ManagedClusterPodIdentityResponseArrayOutput
}

type ManagedClusterPodIdentityResponseArray []ManagedClusterPodIdentityResponseInput

func (ManagedClusterPodIdentityResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedClusterPodIdentityResponse)(nil)).Elem()
}

func (i ManagedClusterPodIdentityResponseArray) ToManagedClusterPodIdentityResponseArrayOutput() ManagedClusterPodIdentityResponseArrayOutput {
	return i.ToManagedClusterPodIdentityResponseArrayOutputWithContext(context.Background())
}

func (i ManagedClusterPodIdentityResponseArray) ToManagedClusterPodIdentityResponseArrayOutputWithContext(ctx context.Context) ManagedClusterPodIdentityResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPodIdentityResponseArrayOutput)
}

type ManagedClusterPodIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterPodIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPodIdentityResponse)(nil)).Elem()
}

func (o ManagedClusterPodIdentityResponseOutput) ToManagedClusterPodIdentityResponseOutput() ManagedClusterPodIdentityResponseOutput {
	return o
}

func (o ManagedClusterPodIdentityResponseOutput) ToManagedClusterPodIdentityResponseOutputWithContext(ctx context.Context) ManagedClusterPodIdentityResponseOutput {
	return o
}

func (o ManagedClusterPodIdentityResponseOutput) BindingSelector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityResponse) *string { return v.BindingSelector }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPodIdentityResponseOutput) Identity() UserAssignedIdentityResponseOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityResponse) UserAssignedIdentityResponse { return v.Identity }).(UserAssignedIdentityResponseOutput)
}

func (o ManagedClusterPodIdentityResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedClusterPodIdentityResponseOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityResponse) string { return v.Namespace }).(pulumi.StringOutput)
}

func (o ManagedClusterPodIdentityResponseOutput) ProvisioningInfo() ManagedClusterPodIdentityResponseProvisioningInfoOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityResponse) ManagedClusterPodIdentityResponseProvisioningInfo {
		return v.ProvisioningInfo
	}).(ManagedClusterPodIdentityResponseProvisioningInfoOutput)
}

func (o ManagedClusterPodIdentityResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type ManagedClusterPodIdentityResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagedClusterPodIdentityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedClusterPodIdentityResponse)(nil)).Elem()
}

func (o ManagedClusterPodIdentityResponseArrayOutput) ToManagedClusterPodIdentityResponseArrayOutput() ManagedClusterPodIdentityResponseArrayOutput {
	return o
}

func (o ManagedClusterPodIdentityResponseArrayOutput) ToManagedClusterPodIdentityResponseArrayOutputWithContext(ctx context.Context) ManagedClusterPodIdentityResponseArrayOutput {
	return o
}

func (o ManagedClusterPodIdentityResponseArrayOutput) Index(i pulumi.IntInput) ManagedClusterPodIdentityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedClusterPodIdentityResponse {
		return vs[0].([]ManagedClusterPodIdentityResponse)[vs[1].(int)]
	}).(ManagedClusterPodIdentityResponseOutput)
}

type ManagedClusterPodIdentityResponseProvisioningInfo struct {
	Error *CloudErrorResponse `pulumi:"error"`
}





type ManagedClusterPodIdentityResponseProvisioningInfoInput interface {
	pulumi.Input

	ToManagedClusterPodIdentityResponseProvisioningInfoOutput() ManagedClusterPodIdentityResponseProvisioningInfoOutput
	ToManagedClusterPodIdentityResponseProvisioningInfoOutputWithContext(context.Context) ManagedClusterPodIdentityResponseProvisioningInfoOutput
}

type ManagedClusterPodIdentityResponseProvisioningInfoArgs struct {
	Error CloudErrorResponsePtrInput `pulumi:"error"`
}

func (ManagedClusterPodIdentityResponseProvisioningInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPodIdentityResponseProvisioningInfo)(nil)).Elem()
}

func (i ManagedClusterPodIdentityResponseProvisioningInfoArgs) ToManagedClusterPodIdentityResponseProvisioningInfoOutput() ManagedClusterPodIdentityResponseProvisioningInfoOutput {
	return i.ToManagedClusterPodIdentityResponseProvisioningInfoOutputWithContext(context.Background())
}

func (i ManagedClusterPodIdentityResponseProvisioningInfoArgs) ToManagedClusterPodIdentityResponseProvisioningInfoOutputWithContext(ctx context.Context) ManagedClusterPodIdentityResponseProvisioningInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPodIdentityResponseProvisioningInfoOutput)
}

type ManagedClusterPodIdentityResponseProvisioningInfoOutput struct{ *pulumi.OutputState }

func (ManagedClusterPodIdentityResponseProvisioningInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPodIdentityResponseProvisioningInfo)(nil)).Elem()
}

func (o ManagedClusterPodIdentityResponseProvisioningInfoOutput) ToManagedClusterPodIdentityResponseProvisioningInfoOutput() ManagedClusterPodIdentityResponseProvisioningInfoOutput {
	return o
}

func (o ManagedClusterPodIdentityResponseProvisioningInfoOutput) ToManagedClusterPodIdentityResponseProvisioningInfoOutputWithContext(ctx context.Context) ManagedClusterPodIdentityResponseProvisioningInfoOutput {
	return o
}

func (o ManagedClusterPodIdentityResponseProvisioningInfoOutput) Error() CloudErrorResponsePtrOutput {
	return o.ApplyT(func(v ManagedClusterPodIdentityResponseProvisioningInfo) *CloudErrorResponse { return v.Error }).(CloudErrorResponsePtrOutput)
}

type ManagedClusterPropertiesAutoScalerProfile struct {
	BalanceSimilarNodeGroups      *string `pulumi:"balanceSimilarNodeGroups"`
	Expander                      *string `pulumi:"expander"`
	MaxEmptyBulkDelete            *string `pulumi:"maxEmptyBulkDelete"`
	MaxGracefulTerminationSec     *string `pulumi:"maxGracefulTerminationSec"`
	MaxNodeProvisionTime          *string `pulumi:"maxNodeProvisionTime"`
	MaxTotalUnreadyPercentage     *string `pulumi:"maxTotalUnreadyPercentage"`
	NewPodScaleUpDelay            *string `pulumi:"newPodScaleUpDelay"`
	OkTotalUnreadyCount           *string `pulumi:"okTotalUnreadyCount"`
	ScaleDownDelayAfterAdd        *string `pulumi:"scaleDownDelayAfterAdd"`
	ScaleDownDelayAfterDelete     *string `pulumi:"scaleDownDelayAfterDelete"`
	ScaleDownDelayAfterFailure    *string `pulumi:"scaleDownDelayAfterFailure"`
	ScaleDownUnneededTime         *string `pulumi:"scaleDownUnneededTime"`
	ScaleDownUnreadyTime          *string `pulumi:"scaleDownUnreadyTime"`
	ScaleDownUtilizationThreshold *string `pulumi:"scaleDownUtilizationThreshold"`
	ScanInterval                  *string `pulumi:"scanInterval"`
	SkipNodesWithLocalStorage     *string `pulumi:"skipNodesWithLocalStorage"`
	SkipNodesWithSystemPods       *string `pulumi:"skipNodesWithSystemPods"`
}





type ManagedClusterPropertiesAutoScalerProfileInput interface {
	pulumi.Input

	ToManagedClusterPropertiesAutoScalerProfileOutput() ManagedClusterPropertiesAutoScalerProfileOutput
	ToManagedClusterPropertiesAutoScalerProfileOutputWithContext(context.Context) ManagedClusterPropertiesAutoScalerProfileOutput
}

type ManagedClusterPropertiesAutoScalerProfileArgs struct {
	BalanceSimilarNodeGroups      pulumi.StringPtrInput `pulumi:"balanceSimilarNodeGroups"`
	Expander                      pulumi.StringPtrInput `pulumi:"expander"`
	MaxEmptyBulkDelete            pulumi.StringPtrInput `pulumi:"maxEmptyBulkDelete"`
	MaxGracefulTerminationSec     pulumi.StringPtrInput `pulumi:"maxGracefulTerminationSec"`
	MaxNodeProvisionTime          pulumi.StringPtrInput `pulumi:"maxNodeProvisionTime"`
	MaxTotalUnreadyPercentage     pulumi.StringPtrInput `pulumi:"maxTotalUnreadyPercentage"`
	NewPodScaleUpDelay            pulumi.StringPtrInput `pulumi:"newPodScaleUpDelay"`
	OkTotalUnreadyCount           pulumi.StringPtrInput `pulumi:"okTotalUnreadyCount"`
	ScaleDownDelayAfterAdd        pulumi.StringPtrInput `pulumi:"scaleDownDelayAfterAdd"`
	ScaleDownDelayAfterDelete     pulumi.StringPtrInput `pulumi:"scaleDownDelayAfterDelete"`
	ScaleDownDelayAfterFailure    pulumi.StringPtrInput `pulumi:"scaleDownDelayAfterFailure"`
	ScaleDownUnneededTime         pulumi.StringPtrInput `pulumi:"scaleDownUnneededTime"`
	ScaleDownUnreadyTime          pulumi.StringPtrInput `pulumi:"scaleDownUnreadyTime"`
	ScaleDownUtilizationThreshold pulumi.StringPtrInput `pulumi:"scaleDownUtilizationThreshold"`
	ScanInterval                  pulumi.StringPtrInput `pulumi:"scanInterval"`
	SkipNodesWithLocalStorage     pulumi.StringPtrInput `pulumi:"skipNodesWithLocalStorage"`
	SkipNodesWithSystemPods       pulumi.StringPtrInput `pulumi:"skipNodesWithSystemPods"`
}

func (ManagedClusterPropertiesAutoScalerProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPropertiesAutoScalerProfile)(nil)).Elem()
}

func (i ManagedClusterPropertiesAutoScalerProfileArgs) ToManagedClusterPropertiesAutoScalerProfileOutput() ManagedClusterPropertiesAutoScalerProfileOutput {
	return i.ToManagedClusterPropertiesAutoScalerProfileOutputWithContext(context.Background())
}

func (i ManagedClusterPropertiesAutoScalerProfileArgs) ToManagedClusterPropertiesAutoScalerProfileOutputWithContext(ctx context.Context) ManagedClusterPropertiesAutoScalerProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPropertiesAutoScalerProfileOutput)
}

func (i ManagedClusterPropertiesAutoScalerProfileArgs) ToManagedClusterPropertiesAutoScalerProfilePtrOutput() ManagedClusterPropertiesAutoScalerProfilePtrOutput {
	return i.ToManagedClusterPropertiesAutoScalerProfilePtrOutputWithContext(context.Background())
}

func (i ManagedClusterPropertiesAutoScalerProfileArgs) ToManagedClusterPropertiesAutoScalerProfilePtrOutputWithContext(ctx context.Context) ManagedClusterPropertiesAutoScalerProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPropertiesAutoScalerProfileOutput).ToManagedClusterPropertiesAutoScalerProfilePtrOutputWithContext(ctx)
}









type ManagedClusterPropertiesAutoScalerProfilePtrInput interface {
	pulumi.Input

	ToManagedClusterPropertiesAutoScalerProfilePtrOutput() ManagedClusterPropertiesAutoScalerProfilePtrOutput
	ToManagedClusterPropertiesAutoScalerProfilePtrOutputWithContext(context.Context) ManagedClusterPropertiesAutoScalerProfilePtrOutput
}

type managedClusterPropertiesAutoScalerProfilePtrType ManagedClusterPropertiesAutoScalerProfileArgs

func ManagedClusterPropertiesAutoScalerProfilePtr(v *ManagedClusterPropertiesAutoScalerProfileArgs) ManagedClusterPropertiesAutoScalerProfilePtrInput {
	return (*managedClusterPropertiesAutoScalerProfilePtrType)(v)
}

func (*managedClusterPropertiesAutoScalerProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterPropertiesAutoScalerProfile)(nil)).Elem()
}

func (i *managedClusterPropertiesAutoScalerProfilePtrType) ToManagedClusterPropertiesAutoScalerProfilePtrOutput() ManagedClusterPropertiesAutoScalerProfilePtrOutput {
	return i.ToManagedClusterPropertiesAutoScalerProfilePtrOutputWithContext(context.Background())
}

func (i *managedClusterPropertiesAutoScalerProfilePtrType) ToManagedClusterPropertiesAutoScalerProfilePtrOutputWithContext(ctx context.Context) ManagedClusterPropertiesAutoScalerProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPropertiesAutoScalerProfilePtrOutput)
}

type ManagedClusterPropertiesAutoScalerProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterPropertiesAutoScalerProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPropertiesAutoScalerProfile)(nil)).Elem()
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) ToManagedClusterPropertiesAutoScalerProfileOutput() ManagedClusterPropertiesAutoScalerProfileOutput {
	return o
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) ToManagedClusterPropertiesAutoScalerProfileOutputWithContext(ctx context.Context) ManagedClusterPropertiesAutoScalerProfileOutput {
	return o
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) ToManagedClusterPropertiesAutoScalerProfilePtrOutput() ManagedClusterPropertiesAutoScalerProfilePtrOutput {
	return o.ToManagedClusterPropertiesAutoScalerProfilePtrOutputWithContext(context.Background())
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) ToManagedClusterPropertiesAutoScalerProfilePtrOutputWithContext(ctx context.Context) ManagedClusterPropertiesAutoScalerProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterPropertiesAutoScalerProfile) *ManagedClusterPropertiesAutoScalerProfile {
		return &v
	}).(ManagedClusterPropertiesAutoScalerProfilePtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) BalanceSimilarNodeGroups() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesAutoScalerProfile) *string { return v.BalanceSimilarNodeGroups }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) Expander() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesAutoScalerProfile) *string { return v.Expander }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) MaxEmptyBulkDelete() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesAutoScalerProfile) *string { return v.MaxEmptyBulkDelete }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) MaxGracefulTerminationSec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesAutoScalerProfile) *string { return v.MaxGracefulTerminationSec }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) MaxNodeProvisionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesAutoScalerProfile) *string { return v.MaxNodeProvisionTime }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) MaxTotalUnreadyPercentage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesAutoScalerProfile) *string { return v.MaxTotalUnreadyPercentage }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) NewPodScaleUpDelay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesAutoScalerProfile) *string { return v.NewPodScaleUpDelay }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) OkTotalUnreadyCount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesAutoScalerProfile) *string { return v.OkTotalUnreadyCount }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) ScaleDownDelayAfterAdd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesAutoScalerProfile) *string { return v.ScaleDownDelayAfterAdd }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) ScaleDownDelayAfterDelete() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesAutoScalerProfile) *string { return v.ScaleDownDelayAfterDelete }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) ScaleDownDelayAfterFailure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesAutoScalerProfile) *string { return v.ScaleDownDelayAfterFailure }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) ScaleDownUnneededTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesAutoScalerProfile) *string { return v.ScaleDownUnneededTime }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) ScaleDownUnreadyTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesAutoScalerProfile) *string { return v.ScaleDownUnreadyTime }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) ScaleDownUtilizationThreshold() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesAutoScalerProfile) *string { return v.ScaleDownUtilizationThreshold }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) ScanInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesAutoScalerProfile) *string { return v.ScanInterval }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) SkipNodesWithLocalStorage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesAutoScalerProfile) *string { return v.SkipNodesWithLocalStorage }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfileOutput) SkipNodesWithSystemPods() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesAutoScalerProfile) *string { return v.SkipNodesWithSystemPods }).(pulumi.StringPtrOutput)
}

type ManagedClusterPropertiesAutoScalerProfilePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterPropertiesAutoScalerProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterPropertiesAutoScalerProfile)(nil)).Elem()
}

func (o ManagedClusterPropertiesAutoScalerProfilePtrOutput) ToManagedClusterPropertiesAutoScalerProfilePtrOutput() ManagedClusterPropertiesAutoScalerProfilePtrOutput {
	return o
}

func (o ManagedClusterPropertiesAutoScalerProfilePtrOutput) ToManagedClusterPropertiesAutoScalerProfilePtrOutputWithContext(ctx context.Context) ManagedClusterPropertiesAutoScalerProfilePtrOutput {
	return o
}

func (o ManagedClusterPropertiesAutoScalerProfilePtrOutput) Elem() ManagedClusterPropertiesAutoScalerProfileOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesAutoScalerProfile) ManagedClusterPropertiesAutoScalerProfile {
		if v != nil {
			return *v
		}
		var ret ManagedClusterPropertiesAutoScalerProfile
		return ret
	}).(ManagedClusterPropertiesAutoScalerProfileOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfilePtrOutput) BalanceSimilarNodeGroups() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.BalanceSimilarNodeGroups
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfilePtrOutput) Expander() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.Expander
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfilePtrOutput) MaxEmptyBulkDelete() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.MaxEmptyBulkDelete
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfilePtrOutput) MaxGracefulTerminationSec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.MaxGracefulTerminationSec
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfilePtrOutput) MaxNodeProvisionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.MaxNodeProvisionTime
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfilePtrOutput) MaxTotalUnreadyPercentage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.MaxTotalUnreadyPercentage
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfilePtrOutput) NewPodScaleUpDelay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.NewPodScaleUpDelay
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfilePtrOutput) OkTotalUnreadyCount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.OkTotalUnreadyCount
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfilePtrOutput) ScaleDownDelayAfterAdd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.ScaleDownDelayAfterAdd
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfilePtrOutput) ScaleDownDelayAfterDelete() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.ScaleDownDelayAfterDelete
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfilePtrOutput) ScaleDownDelayAfterFailure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.ScaleDownDelayAfterFailure
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfilePtrOutput) ScaleDownUnneededTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.ScaleDownUnneededTime
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfilePtrOutput) ScaleDownUnreadyTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.ScaleDownUnreadyTime
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfilePtrOutput) ScaleDownUtilizationThreshold() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.ScaleDownUtilizationThreshold
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfilePtrOutput) ScanInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.ScanInterval
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfilePtrOutput) SkipNodesWithLocalStorage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.SkipNodesWithLocalStorage
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesAutoScalerProfilePtrOutput) SkipNodesWithSystemPods() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.SkipNodesWithSystemPods
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterPropertiesIdentityProfile struct {
	ClientId   *string `pulumi:"clientId"`
	ObjectId   *string `pulumi:"objectId"`
	ResourceId *string `pulumi:"resourceId"`
}





type ManagedClusterPropertiesIdentityProfileInput interface {
	pulumi.Input

	ToManagedClusterPropertiesIdentityProfileOutput() ManagedClusterPropertiesIdentityProfileOutput
	ToManagedClusterPropertiesIdentityProfileOutputWithContext(context.Context) ManagedClusterPropertiesIdentityProfileOutput
}

type ManagedClusterPropertiesIdentityProfileArgs struct {
	ClientId   pulumi.StringPtrInput `pulumi:"clientId"`
	ObjectId   pulumi.StringPtrInput `pulumi:"objectId"`
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (ManagedClusterPropertiesIdentityProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPropertiesIdentityProfile)(nil)).Elem()
}

func (i ManagedClusterPropertiesIdentityProfileArgs) ToManagedClusterPropertiesIdentityProfileOutput() ManagedClusterPropertiesIdentityProfileOutput {
	return i.ToManagedClusterPropertiesIdentityProfileOutputWithContext(context.Background())
}

func (i ManagedClusterPropertiesIdentityProfileArgs) ToManagedClusterPropertiesIdentityProfileOutputWithContext(ctx context.Context) ManagedClusterPropertiesIdentityProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPropertiesIdentityProfileOutput)
}





type ManagedClusterPropertiesIdentityProfileMapInput interface {
	pulumi.Input

	ToManagedClusterPropertiesIdentityProfileMapOutput() ManagedClusterPropertiesIdentityProfileMapOutput
	ToManagedClusterPropertiesIdentityProfileMapOutputWithContext(context.Context) ManagedClusterPropertiesIdentityProfileMapOutput
}

type ManagedClusterPropertiesIdentityProfileMap map[string]ManagedClusterPropertiesIdentityProfileInput

func (ManagedClusterPropertiesIdentityProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterPropertiesIdentityProfile)(nil)).Elem()
}

func (i ManagedClusterPropertiesIdentityProfileMap) ToManagedClusterPropertiesIdentityProfileMapOutput() ManagedClusterPropertiesIdentityProfileMapOutput {
	return i.ToManagedClusterPropertiesIdentityProfileMapOutputWithContext(context.Background())
}

func (i ManagedClusterPropertiesIdentityProfileMap) ToManagedClusterPropertiesIdentityProfileMapOutputWithContext(ctx context.Context) ManagedClusterPropertiesIdentityProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPropertiesIdentityProfileMapOutput)
}

type ManagedClusterPropertiesIdentityProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterPropertiesIdentityProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPropertiesIdentityProfile)(nil)).Elem()
}

func (o ManagedClusterPropertiesIdentityProfileOutput) ToManagedClusterPropertiesIdentityProfileOutput() ManagedClusterPropertiesIdentityProfileOutput {
	return o
}

func (o ManagedClusterPropertiesIdentityProfileOutput) ToManagedClusterPropertiesIdentityProfileOutputWithContext(ctx context.Context) ManagedClusterPropertiesIdentityProfileOutput {
	return o
}

func (o ManagedClusterPropertiesIdentityProfileOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesIdentityProfile) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesIdentityProfileOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesIdentityProfile) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesIdentityProfileOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesIdentityProfile) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type ManagedClusterPropertiesIdentityProfileMapOutput struct{ *pulumi.OutputState }

func (ManagedClusterPropertiesIdentityProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterPropertiesIdentityProfile)(nil)).Elem()
}

func (o ManagedClusterPropertiesIdentityProfileMapOutput) ToManagedClusterPropertiesIdentityProfileMapOutput() ManagedClusterPropertiesIdentityProfileMapOutput {
	return o
}

func (o ManagedClusterPropertiesIdentityProfileMapOutput) ToManagedClusterPropertiesIdentityProfileMapOutputWithContext(ctx context.Context) ManagedClusterPropertiesIdentityProfileMapOutput {
	return o
}

func (o ManagedClusterPropertiesIdentityProfileMapOutput) MapIndex(k pulumi.StringInput) ManagedClusterPropertiesIdentityProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ManagedClusterPropertiesIdentityProfile {
		return vs[0].(map[string]ManagedClusterPropertiesIdentityProfile)[vs[1].(string)]
	}).(ManagedClusterPropertiesIdentityProfileOutput)
}

type ManagedClusterPropertiesResponseAutoScalerProfile struct {
	BalanceSimilarNodeGroups      *string `pulumi:"balanceSimilarNodeGroups"`
	Expander                      *string `pulumi:"expander"`
	MaxEmptyBulkDelete            *string `pulumi:"maxEmptyBulkDelete"`
	MaxGracefulTerminationSec     *string `pulumi:"maxGracefulTerminationSec"`
	MaxNodeProvisionTime          *string `pulumi:"maxNodeProvisionTime"`
	MaxTotalUnreadyPercentage     *string `pulumi:"maxTotalUnreadyPercentage"`
	NewPodScaleUpDelay            *string `pulumi:"newPodScaleUpDelay"`
	OkTotalUnreadyCount           *string `pulumi:"okTotalUnreadyCount"`
	ScaleDownDelayAfterAdd        *string `pulumi:"scaleDownDelayAfterAdd"`
	ScaleDownDelayAfterDelete     *string `pulumi:"scaleDownDelayAfterDelete"`
	ScaleDownDelayAfterFailure    *string `pulumi:"scaleDownDelayAfterFailure"`
	ScaleDownUnneededTime         *string `pulumi:"scaleDownUnneededTime"`
	ScaleDownUnreadyTime          *string `pulumi:"scaleDownUnreadyTime"`
	ScaleDownUtilizationThreshold *string `pulumi:"scaleDownUtilizationThreshold"`
	ScanInterval                  *string `pulumi:"scanInterval"`
	SkipNodesWithLocalStorage     *string `pulumi:"skipNodesWithLocalStorage"`
	SkipNodesWithSystemPods       *string `pulumi:"skipNodesWithSystemPods"`
}





type ManagedClusterPropertiesResponseAutoScalerProfileInput interface {
	pulumi.Input

	ToManagedClusterPropertiesResponseAutoScalerProfileOutput() ManagedClusterPropertiesResponseAutoScalerProfileOutput
	ToManagedClusterPropertiesResponseAutoScalerProfileOutputWithContext(context.Context) ManagedClusterPropertiesResponseAutoScalerProfileOutput
}

type ManagedClusterPropertiesResponseAutoScalerProfileArgs struct {
	BalanceSimilarNodeGroups      pulumi.StringPtrInput `pulumi:"balanceSimilarNodeGroups"`
	Expander                      pulumi.StringPtrInput `pulumi:"expander"`
	MaxEmptyBulkDelete            pulumi.StringPtrInput `pulumi:"maxEmptyBulkDelete"`
	MaxGracefulTerminationSec     pulumi.StringPtrInput `pulumi:"maxGracefulTerminationSec"`
	MaxNodeProvisionTime          pulumi.StringPtrInput `pulumi:"maxNodeProvisionTime"`
	MaxTotalUnreadyPercentage     pulumi.StringPtrInput `pulumi:"maxTotalUnreadyPercentage"`
	NewPodScaleUpDelay            pulumi.StringPtrInput `pulumi:"newPodScaleUpDelay"`
	OkTotalUnreadyCount           pulumi.StringPtrInput `pulumi:"okTotalUnreadyCount"`
	ScaleDownDelayAfterAdd        pulumi.StringPtrInput `pulumi:"scaleDownDelayAfterAdd"`
	ScaleDownDelayAfterDelete     pulumi.StringPtrInput `pulumi:"scaleDownDelayAfterDelete"`
	ScaleDownDelayAfterFailure    pulumi.StringPtrInput `pulumi:"scaleDownDelayAfterFailure"`
	ScaleDownUnneededTime         pulumi.StringPtrInput `pulumi:"scaleDownUnneededTime"`
	ScaleDownUnreadyTime          pulumi.StringPtrInput `pulumi:"scaleDownUnreadyTime"`
	ScaleDownUtilizationThreshold pulumi.StringPtrInput `pulumi:"scaleDownUtilizationThreshold"`
	ScanInterval                  pulumi.StringPtrInput `pulumi:"scanInterval"`
	SkipNodesWithLocalStorage     pulumi.StringPtrInput `pulumi:"skipNodesWithLocalStorage"`
	SkipNodesWithSystemPods       pulumi.StringPtrInput `pulumi:"skipNodesWithSystemPods"`
}

func (ManagedClusterPropertiesResponseAutoScalerProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPropertiesResponseAutoScalerProfile)(nil)).Elem()
}

func (i ManagedClusterPropertiesResponseAutoScalerProfileArgs) ToManagedClusterPropertiesResponseAutoScalerProfileOutput() ManagedClusterPropertiesResponseAutoScalerProfileOutput {
	return i.ToManagedClusterPropertiesResponseAutoScalerProfileOutputWithContext(context.Background())
}

func (i ManagedClusterPropertiesResponseAutoScalerProfileArgs) ToManagedClusterPropertiesResponseAutoScalerProfileOutputWithContext(ctx context.Context) ManagedClusterPropertiesResponseAutoScalerProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPropertiesResponseAutoScalerProfileOutput)
}

func (i ManagedClusterPropertiesResponseAutoScalerProfileArgs) ToManagedClusterPropertiesResponseAutoScalerProfilePtrOutput() ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput {
	return i.ToManagedClusterPropertiesResponseAutoScalerProfilePtrOutputWithContext(context.Background())
}

func (i ManagedClusterPropertiesResponseAutoScalerProfileArgs) ToManagedClusterPropertiesResponseAutoScalerProfilePtrOutputWithContext(ctx context.Context) ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPropertiesResponseAutoScalerProfileOutput).ToManagedClusterPropertiesResponseAutoScalerProfilePtrOutputWithContext(ctx)
}









type ManagedClusterPropertiesResponseAutoScalerProfilePtrInput interface {
	pulumi.Input

	ToManagedClusterPropertiesResponseAutoScalerProfilePtrOutput() ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput
	ToManagedClusterPropertiesResponseAutoScalerProfilePtrOutputWithContext(context.Context) ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput
}

type managedClusterPropertiesResponseAutoScalerProfilePtrType ManagedClusterPropertiesResponseAutoScalerProfileArgs

func ManagedClusterPropertiesResponseAutoScalerProfilePtr(v *ManagedClusterPropertiesResponseAutoScalerProfileArgs) ManagedClusterPropertiesResponseAutoScalerProfilePtrInput {
	return (*managedClusterPropertiesResponseAutoScalerProfilePtrType)(v)
}

func (*managedClusterPropertiesResponseAutoScalerProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterPropertiesResponseAutoScalerProfile)(nil)).Elem()
}

func (i *managedClusterPropertiesResponseAutoScalerProfilePtrType) ToManagedClusterPropertiesResponseAutoScalerProfilePtrOutput() ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput {
	return i.ToManagedClusterPropertiesResponseAutoScalerProfilePtrOutputWithContext(context.Background())
}

func (i *managedClusterPropertiesResponseAutoScalerProfilePtrType) ToManagedClusterPropertiesResponseAutoScalerProfilePtrOutputWithContext(ctx context.Context) ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput)
}

type ManagedClusterPropertiesResponseAutoScalerProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterPropertiesResponseAutoScalerProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPropertiesResponseAutoScalerProfile)(nil)).Elem()
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) ToManagedClusterPropertiesResponseAutoScalerProfileOutput() ManagedClusterPropertiesResponseAutoScalerProfileOutput {
	return o
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) ToManagedClusterPropertiesResponseAutoScalerProfileOutputWithContext(ctx context.Context) ManagedClusterPropertiesResponseAutoScalerProfileOutput {
	return o
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) ToManagedClusterPropertiesResponseAutoScalerProfilePtrOutput() ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput {
	return o.ToManagedClusterPropertiesResponseAutoScalerProfilePtrOutputWithContext(context.Background())
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) ToManagedClusterPropertiesResponseAutoScalerProfilePtrOutputWithContext(ctx context.Context) ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterPropertiesResponseAutoScalerProfile) *ManagedClusterPropertiesResponseAutoScalerProfile {
		return &v
	}).(ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) BalanceSimilarNodeGroups() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseAutoScalerProfile) *string { return v.BalanceSimilarNodeGroups }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) Expander() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseAutoScalerProfile) *string { return v.Expander }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) MaxEmptyBulkDelete() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseAutoScalerProfile) *string { return v.MaxEmptyBulkDelete }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) MaxGracefulTerminationSec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseAutoScalerProfile) *string { return v.MaxGracefulTerminationSec }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) MaxNodeProvisionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseAutoScalerProfile) *string { return v.MaxNodeProvisionTime }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) MaxTotalUnreadyPercentage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseAutoScalerProfile) *string { return v.MaxTotalUnreadyPercentage }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) NewPodScaleUpDelay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseAutoScalerProfile) *string { return v.NewPodScaleUpDelay }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) OkTotalUnreadyCount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseAutoScalerProfile) *string { return v.OkTotalUnreadyCount }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) ScaleDownDelayAfterAdd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseAutoScalerProfile) *string { return v.ScaleDownDelayAfterAdd }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) ScaleDownDelayAfterDelete() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseAutoScalerProfile) *string { return v.ScaleDownDelayAfterDelete }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) ScaleDownDelayAfterFailure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseAutoScalerProfile) *string { return v.ScaleDownDelayAfterFailure }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) ScaleDownUnneededTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseAutoScalerProfile) *string { return v.ScaleDownUnneededTime }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) ScaleDownUnreadyTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseAutoScalerProfile) *string { return v.ScaleDownUnreadyTime }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) ScaleDownUtilizationThreshold() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseAutoScalerProfile) *string {
		return v.ScaleDownUtilizationThreshold
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) ScanInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseAutoScalerProfile) *string { return v.ScanInterval }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) SkipNodesWithLocalStorage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseAutoScalerProfile) *string { return v.SkipNodesWithLocalStorage }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfileOutput) SkipNodesWithSystemPods() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseAutoScalerProfile) *string { return v.SkipNodesWithSystemPods }).(pulumi.StringPtrOutput)
}

type ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterPropertiesResponseAutoScalerProfile)(nil)).Elem()
}

func (o ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) ToManagedClusterPropertiesResponseAutoScalerProfilePtrOutput() ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput {
	return o
}

func (o ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) ToManagedClusterPropertiesResponseAutoScalerProfilePtrOutputWithContext(ctx context.Context) ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput {
	return o
}

func (o ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) Elem() ManagedClusterPropertiesResponseAutoScalerProfileOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesResponseAutoScalerProfile) ManagedClusterPropertiesResponseAutoScalerProfile {
		if v != nil {
			return *v
		}
		var ret ManagedClusterPropertiesResponseAutoScalerProfile
		return ret
	}).(ManagedClusterPropertiesResponseAutoScalerProfileOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) BalanceSimilarNodeGroups() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesResponseAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.BalanceSimilarNodeGroups
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) Expander() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesResponseAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.Expander
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) MaxEmptyBulkDelete() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesResponseAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.MaxEmptyBulkDelete
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) MaxGracefulTerminationSec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesResponseAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.MaxGracefulTerminationSec
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) MaxNodeProvisionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesResponseAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.MaxNodeProvisionTime
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) MaxTotalUnreadyPercentage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesResponseAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.MaxTotalUnreadyPercentage
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) NewPodScaleUpDelay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesResponseAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.NewPodScaleUpDelay
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) OkTotalUnreadyCount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesResponseAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.OkTotalUnreadyCount
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) ScaleDownDelayAfterAdd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesResponseAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.ScaleDownDelayAfterAdd
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) ScaleDownDelayAfterDelete() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesResponseAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.ScaleDownDelayAfterDelete
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) ScaleDownDelayAfterFailure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesResponseAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.ScaleDownDelayAfterFailure
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) ScaleDownUnneededTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesResponseAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.ScaleDownUnneededTime
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) ScaleDownUnreadyTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesResponseAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.ScaleDownUnreadyTime
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) ScaleDownUtilizationThreshold() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesResponseAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.ScaleDownUtilizationThreshold
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) ScanInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesResponseAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.ScanInterval
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) SkipNodesWithLocalStorage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesResponseAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.SkipNodesWithLocalStorage
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput) SkipNodesWithSystemPods() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterPropertiesResponseAutoScalerProfile) *string {
		if v == nil {
			return nil
		}
		return v.SkipNodesWithSystemPods
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterPropertiesResponseIdentityProfile struct {
	ClientId   *string `pulumi:"clientId"`
	ObjectId   *string `pulumi:"objectId"`
	ResourceId *string `pulumi:"resourceId"`
}





type ManagedClusterPropertiesResponseIdentityProfileInput interface {
	pulumi.Input

	ToManagedClusterPropertiesResponseIdentityProfileOutput() ManagedClusterPropertiesResponseIdentityProfileOutput
	ToManagedClusterPropertiesResponseIdentityProfileOutputWithContext(context.Context) ManagedClusterPropertiesResponseIdentityProfileOutput
}

type ManagedClusterPropertiesResponseIdentityProfileArgs struct {
	ClientId   pulumi.StringPtrInput `pulumi:"clientId"`
	ObjectId   pulumi.StringPtrInput `pulumi:"objectId"`
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (ManagedClusterPropertiesResponseIdentityProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPropertiesResponseIdentityProfile)(nil)).Elem()
}

func (i ManagedClusterPropertiesResponseIdentityProfileArgs) ToManagedClusterPropertiesResponseIdentityProfileOutput() ManagedClusterPropertiesResponseIdentityProfileOutput {
	return i.ToManagedClusterPropertiesResponseIdentityProfileOutputWithContext(context.Background())
}

func (i ManagedClusterPropertiesResponseIdentityProfileArgs) ToManagedClusterPropertiesResponseIdentityProfileOutputWithContext(ctx context.Context) ManagedClusterPropertiesResponseIdentityProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPropertiesResponseIdentityProfileOutput)
}





type ManagedClusterPropertiesResponseIdentityProfileMapInput interface {
	pulumi.Input

	ToManagedClusterPropertiesResponseIdentityProfileMapOutput() ManagedClusterPropertiesResponseIdentityProfileMapOutput
	ToManagedClusterPropertiesResponseIdentityProfileMapOutputWithContext(context.Context) ManagedClusterPropertiesResponseIdentityProfileMapOutput
}

type ManagedClusterPropertiesResponseIdentityProfileMap map[string]ManagedClusterPropertiesResponseIdentityProfileInput

func (ManagedClusterPropertiesResponseIdentityProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterPropertiesResponseIdentityProfile)(nil)).Elem()
}

func (i ManagedClusterPropertiesResponseIdentityProfileMap) ToManagedClusterPropertiesResponseIdentityProfileMapOutput() ManagedClusterPropertiesResponseIdentityProfileMapOutput {
	return i.ToManagedClusterPropertiesResponseIdentityProfileMapOutputWithContext(context.Background())
}

func (i ManagedClusterPropertiesResponseIdentityProfileMap) ToManagedClusterPropertiesResponseIdentityProfileMapOutputWithContext(ctx context.Context) ManagedClusterPropertiesResponseIdentityProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPropertiesResponseIdentityProfileMapOutput)
}

type ManagedClusterPropertiesResponseIdentityProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterPropertiesResponseIdentityProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPropertiesResponseIdentityProfile)(nil)).Elem()
}

func (o ManagedClusterPropertiesResponseIdentityProfileOutput) ToManagedClusterPropertiesResponseIdentityProfileOutput() ManagedClusterPropertiesResponseIdentityProfileOutput {
	return o
}

func (o ManagedClusterPropertiesResponseIdentityProfileOutput) ToManagedClusterPropertiesResponseIdentityProfileOutputWithContext(ctx context.Context) ManagedClusterPropertiesResponseIdentityProfileOutput {
	return o
}

func (o ManagedClusterPropertiesResponseIdentityProfileOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseIdentityProfile) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseIdentityProfileOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseIdentityProfile) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseIdentityProfileOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseIdentityProfile) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type ManagedClusterPropertiesResponseIdentityProfileMapOutput struct{ *pulumi.OutputState }

func (ManagedClusterPropertiesResponseIdentityProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterPropertiesResponseIdentityProfile)(nil)).Elem()
}

func (o ManagedClusterPropertiesResponseIdentityProfileMapOutput) ToManagedClusterPropertiesResponseIdentityProfileMapOutput() ManagedClusterPropertiesResponseIdentityProfileMapOutput {
	return o
}

func (o ManagedClusterPropertiesResponseIdentityProfileMapOutput) ToManagedClusterPropertiesResponseIdentityProfileMapOutputWithContext(ctx context.Context) ManagedClusterPropertiesResponseIdentityProfileMapOutput {
	return o
}

func (o ManagedClusterPropertiesResponseIdentityProfileMapOutput) MapIndex(k pulumi.StringInput) ManagedClusterPropertiesResponseIdentityProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ManagedClusterPropertiesResponseIdentityProfile {
		return vs[0].(map[string]ManagedClusterPropertiesResponseIdentityProfile)[vs[1].(string)]
	}).(ManagedClusterPropertiesResponseIdentityProfileOutput)
}

type ManagedClusterSKU struct {
	Name *string `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type ManagedClusterSKUInput interface {
	pulumi.Input

	ToManagedClusterSKUOutput() ManagedClusterSKUOutput
	ToManagedClusterSKUOutputWithContext(context.Context) ManagedClusterSKUOutput
}

type ManagedClusterSKUArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (ManagedClusterSKUArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterSKU)(nil)).Elem()
}

func (i ManagedClusterSKUArgs) ToManagedClusterSKUOutput() ManagedClusterSKUOutput {
	return i.ToManagedClusterSKUOutputWithContext(context.Background())
}

func (i ManagedClusterSKUArgs) ToManagedClusterSKUOutputWithContext(ctx context.Context) ManagedClusterSKUOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterSKUOutput)
}

func (i ManagedClusterSKUArgs) ToManagedClusterSKUPtrOutput() ManagedClusterSKUPtrOutput {
	return i.ToManagedClusterSKUPtrOutputWithContext(context.Background())
}

func (i ManagedClusterSKUArgs) ToManagedClusterSKUPtrOutputWithContext(ctx context.Context) ManagedClusterSKUPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterSKUOutput).ToManagedClusterSKUPtrOutputWithContext(ctx)
}









type ManagedClusterSKUPtrInput interface {
	pulumi.Input

	ToManagedClusterSKUPtrOutput() ManagedClusterSKUPtrOutput
	ToManagedClusterSKUPtrOutputWithContext(context.Context) ManagedClusterSKUPtrOutput
}

type managedClusterSKUPtrType ManagedClusterSKUArgs

func ManagedClusterSKUPtr(v *ManagedClusterSKUArgs) ManagedClusterSKUPtrInput {
	return (*managedClusterSKUPtrType)(v)
}

func (*managedClusterSKUPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterSKU)(nil)).Elem()
}

func (i *managedClusterSKUPtrType) ToManagedClusterSKUPtrOutput() ManagedClusterSKUPtrOutput {
	return i.ToManagedClusterSKUPtrOutputWithContext(context.Background())
}

func (i *managedClusterSKUPtrType) ToManagedClusterSKUPtrOutputWithContext(ctx context.Context) ManagedClusterSKUPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterSKUPtrOutput)
}

type ManagedClusterSKUOutput struct{ *pulumi.OutputState }

func (ManagedClusterSKUOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterSKU)(nil)).Elem()
}

func (o ManagedClusterSKUOutput) ToManagedClusterSKUOutput() ManagedClusterSKUOutput {
	return o
}

func (o ManagedClusterSKUOutput) ToManagedClusterSKUOutputWithContext(ctx context.Context) ManagedClusterSKUOutput {
	return o
}

func (o ManagedClusterSKUOutput) ToManagedClusterSKUPtrOutput() ManagedClusterSKUPtrOutput {
	return o.ToManagedClusterSKUPtrOutputWithContext(context.Background())
}

func (o ManagedClusterSKUOutput) ToManagedClusterSKUPtrOutputWithContext(ctx context.Context) ManagedClusterSKUPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterSKU) *ManagedClusterSKU {
		return &v
	}).(ManagedClusterSKUPtrOutput)
}

func (o ManagedClusterSKUOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterSKU) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterSKUOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterSKU) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ManagedClusterSKUPtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterSKUPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterSKU)(nil)).Elem()
}

func (o ManagedClusterSKUPtrOutput) ToManagedClusterSKUPtrOutput() ManagedClusterSKUPtrOutput {
	return o
}

func (o ManagedClusterSKUPtrOutput) ToManagedClusterSKUPtrOutputWithContext(ctx context.Context) ManagedClusterSKUPtrOutput {
	return o
}

func (o ManagedClusterSKUPtrOutput) Elem() ManagedClusterSKUOutput {
	return o.ApplyT(func(v *ManagedClusterSKU) ManagedClusterSKU {
		if v != nil {
			return *v
		}
		var ret ManagedClusterSKU
		return ret
	}).(ManagedClusterSKUOutput)
}

func (o ManagedClusterSKUPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterSKU) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterSKUPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterSKU) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterSKUResponse struct {
	Name *string `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type ManagedClusterSKUResponseInput interface {
	pulumi.Input

	ToManagedClusterSKUResponseOutput() ManagedClusterSKUResponseOutput
	ToManagedClusterSKUResponseOutputWithContext(context.Context) ManagedClusterSKUResponseOutput
}

type ManagedClusterSKUResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (ManagedClusterSKUResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterSKUResponse)(nil)).Elem()
}

func (i ManagedClusterSKUResponseArgs) ToManagedClusterSKUResponseOutput() ManagedClusterSKUResponseOutput {
	return i.ToManagedClusterSKUResponseOutputWithContext(context.Background())
}

func (i ManagedClusterSKUResponseArgs) ToManagedClusterSKUResponseOutputWithContext(ctx context.Context) ManagedClusterSKUResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterSKUResponseOutput)
}

func (i ManagedClusterSKUResponseArgs) ToManagedClusterSKUResponsePtrOutput() ManagedClusterSKUResponsePtrOutput {
	return i.ToManagedClusterSKUResponsePtrOutputWithContext(context.Background())
}

func (i ManagedClusterSKUResponseArgs) ToManagedClusterSKUResponsePtrOutputWithContext(ctx context.Context) ManagedClusterSKUResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterSKUResponseOutput).ToManagedClusterSKUResponsePtrOutputWithContext(ctx)
}









type ManagedClusterSKUResponsePtrInput interface {
	pulumi.Input

	ToManagedClusterSKUResponsePtrOutput() ManagedClusterSKUResponsePtrOutput
	ToManagedClusterSKUResponsePtrOutputWithContext(context.Context) ManagedClusterSKUResponsePtrOutput
}

type managedClusterSKUResponsePtrType ManagedClusterSKUResponseArgs

func ManagedClusterSKUResponsePtr(v *ManagedClusterSKUResponseArgs) ManagedClusterSKUResponsePtrInput {
	return (*managedClusterSKUResponsePtrType)(v)
}

func (*managedClusterSKUResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterSKUResponse)(nil)).Elem()
}

func (i *managedClusterSKUResponsePtrType) ToManagedClusterSKUResponsePtrOutput() ManagedClusterSKUResponsePtrOutput {
	return i.ToManagedClusterSKUResponsePtrOutputWithContext(context.Background())
}

func (i *managedClusterSKUResponsePtrType) ToManagedClusterSKUResponsePtrOutputWithContext(ctx context.Context) ManagedClusterSKUResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterSKUResponsePtrOutput)
}

type ManagedClusterSKUResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterSKUResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterSKUResponse)(nil)).Elem()
}

func (o ManagedClusterSKUResponseOutput) ToManagedClusterSKUResponseOutput() ManagedClusterSKUResponseOutput {
	return o
}

func (o ManagedClusterSKUResponseOutput) ToManagedClusterSKUResponseOutputWithContext(ctx context.Context) ManagedClusterSKUResponseOutput {
	return o
}

func (o ManagedClusterSKUResponseOutput) ToManagedClusterSKUResponsePtrOutput() ManagedClusterSKUResponsePtrOutput {
	return o.ToManagedClusterSKUResponsePtrOutputWithContext(context.Background())
}

func (o ManagedClusterSKUResponseOutput) ToManagedClusterSKUResponsePtrOutputWithContext(ctx context.Context) ManagedClusterSKUResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterSKUResponse) *ManagedClusterSKUResponse {
		return &v
	}).(ManagedClusterSKUResponsePtrOutput)
}

func (o ManagedClusterSKUResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterSKUResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterSKUResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterSKUResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ManagedClusterSKUResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterSKUResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterSKUResponse)(nil)).Elem()
}

func (o ManagedClusterSKUResponsePtrOutput) ToManagedClusterSKUResponsePtrOutput() ManagedClusterSKUResponsePtrOutput {
	return o
}

func (o ManagedClusterSKUResponsePtrOutput) ToManagedClusterSKUResponsePtrOutputWithContext(ctx context.Context) ManagedClusterSKUResponsePtrOutput {
	return o
}

func (o ManagedClusterSKUResponsePtrOutput) Elem() ManagedClusterSKUResponseOutput {
	return o.ApplyT(func(v *ManagedClusterSKUResponse) ManagedClusterSKUResponse {
		if v != nil {
			return *v
		}
		var ret ManagedClusterSKUResponse
		return ret
	}).(ManagedClusterSKUResponseOutput)
}

func (o ManagedClusterSKUResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterSKUResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterSKUResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterSKUResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterServicePrincipalProfile struct {
	ClientId string  `pulumi:"clientId"`
	Secret   *string `pulumi:"secret"`
}





type ManagedClusterServicePrincipalProfileInput interface {
	pulumi.Input

	ToManagedClusterServicePrincipalProfileOutput() ManagedClusterServicePrincipalProfileOutput
	ToManagedClusterServicePrincipalProfileOutputWithContext(context.Context) ManagedClusterServicePrincipalProfileOutput
}

type ManagedClusterServicePrincipalProfileArgs struct {
	ClientId pulumi.StringInput    `pulumi:"clientId"`
	Secret   pulumi.StringPtrInput `pulumi:"secret"`
}

func (ManagedClusterServicePrincipalProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterServicePrincipalProfile)(nil)).Elem()
}

func (i ManagedClusterServicePrincipalProfileArgs) ToManagedClusterServicePrincipalProfileOutput() ManagedClusterServicePrincipalProfileOutput {
	return i.ToManagedClusterServicePrincipalProfileOutputWithContext(context.Background())
}

func (i ManagedClusterServicePrincipalProfileArgs) ToManagedClusterServicePrincipalProfileOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterServicePrincipalProfileOutput)
}

func (i ManagedClusterServicePrincipalProfileArgs) ToManagedClusterServicePrincipalProfilePtrOutput() ManagedClusterServicePrincipalProfilePtrOutput {
	return i.ToManagedClusterServicePrincipalProfilePtrOutputWithContext(context.Background())
}

func (i ManagedClusterServicePrincipalProfileArgs) ToManagedClusterServicePrincipalProfilePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterServicePrincipalProfileOutput).ToManagedClusterServicePrincipalProfilePtrOutputWithContext(ctx)
}









type ManagedClusterServicePrincipalProfilePtrInput interface {
	pulumi.Input

	ToManagedClusterServicePrincipalProfilePtrOutput() ManagedClusterServicePrincipalProfilePtrOutput
	ToManagedClusterServicePrincipalProfilePtrOutputWithContext(context.Context) ManagedClusterServicePrincipalProfilePtrOutput
}

type managedClusterServicePrincipalProfilePtrType ManagedClusterServicePrincipalProfileArgs

func ManagedClusterServicePrincipalProfilePtr(v *ManagedClusterServicePrincipalProfileArgs) ManagedClusterServicePrincipalProfilePtrInput {
	return (*managedClusterServicePrincipalProfilePtrType)(v)
}

func (*managedClusterServicePrincipalProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterServicePrincipalProfile)(nil)).Elem()
}

func (i *managedClusterServicePrincipalProfilePtrType) ToManagedClusterServicePrincipalProfilePtrOutput() ManagedClusterServicePrincipalProfilePtrOutput {
	return i.ToManagedClusterServicePrincipalProfilePtrOutputWithContext(context.Background())
}

func (i *managedClusterServicePrincipalProfilePtrType) ToManagedClusterServicePrincipalProfilePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterServicePrincipalProfilePtrOutput)
}

type ManagedClusterServicePrincipalProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterServicePrincipalProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterServicePrincipalProfile)(nil)).Elem()
}

func (o ManagedClusterServicePrincipalProfileOutput) ToManagedClusterServicePrincipalProfileOutput() ManagedClusterServicePrincipalProfileOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfileOutput) ToManagedClusterServicePrincipalProfileOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfileOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfileOutput) ToManagedClusterServicePrincipalProfilePtrOutput() ManagedClusterServicePrincipalProfilePtrOutput {
	return o.ToManagedClusterServicePrincipalProfilePtrOutputWithContext(context.Background())
}

func (o ManagedClusterServicePrincipalProfileOutput) ToManagedClusterServicePrincipalProfilePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterServicePrincipalProfile) *ManagedClusterServicePrincipalProfile {
		return &v
	}).(ManagedClusterServicePrincipalProfilePtrOutput)
}

func (o ManagedClusterServicePrincipalProfileOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterServicePrincipalProfile) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ManagedClusterServicePrincipalProfileOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterServicePrincipalProfile) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

type ManagedClusterServicePrincipalProfilePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterServicePrincipalProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterServicePrincipalProfile)(nil)).Elem()
}

func (o ManagedClusterServicePrincipalProfilePtrOutput) ToManagedClusterServicePrincipalProfilePtrOutput() ManagedClusterServicePrincipalProfilePtrOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfilePtrOutput) ToManagedClusterServicePrincipalProfilePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfilePtrOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfilePtrOutput) Elem() ManagedClusterServicePrincipalProfileOutput {
	return o.ApplyT(func(v *ManagedClusterServicePrincipalProfile) ManagedClusterServicePrincipalProfile {
		if v != nil {
			return *v
		}
		var ret ManagedClusterServicePrincipalProfile
		return ret
	}).(ManagedClusterServicePrincipalProfileOutput)
}

func (o ManagedClusterServicePrincipalProfilePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterServicePrincipalProfile) *string {
		if v == nil {
			return nil
		}
		return &v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterServicePrincipalProfilePtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterServicePrincipalProfile) *string {
		if v == nil {
			return nil
		}
		return v.Secret
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterServicePrincipalProfileResponse struct {
	ClientId string  `pulumi:"clientId"`
	Secret   *string `pulumi:"secret"`
}





type ManagedClusterServicePrincipalProfileResponseInput interface {
	pulumi.Input

	ToManagedClusterServicePrincipalProfileResponseOutput() ManagedClusterServicePrincipalProfileResponseOutput
	ToManagedClusterServicePrincipalProfileResponseOutputWithContext(context.Context) ManagedClusterServicePrincipalProfileResponseOutput
}

type ManagedClusterServicePrincipalProfileResponseArgs struct {
	ClientId pulumi.StringInput    `pulumi:"clientId"`
	Secret   pulumi.StringPtrInput `pulumi:"secret"`
}

func (ManagedClusterServicePrincipalProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterServicePrincipalProfileResponse)(nil)).Elem()
}

func (i ManagedClusterServicePrincipalProfileResponseArgs) ToManagedClusterServicePrincipalProfileResponseOutput() ManagedClusterServicePrincipalProfileResponseOutput {
	return i.ToManagedClusterServicePrincipalProfileResponseOutputWithContext(context.Background())
}

func (i ManagedClusterServicePrincipalProfileResponseArgs) ToManagedClusterServicePrincipalProfileResponseOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterServicePrincipalProfileResponseOutput)
}

func (i ManagedClusterServicePrincipalProfileResponseArgs) ToManagedClusterServicePrincipalProfileResponsePtrOutput() ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return i.ToManagedClusterServicePrincipalProfileResponsePtrOutputWithContext(context.Background())
}

func (i ManagedClusterServicePrincipalProfileResponseArgs) ToManagedClusterServicePrincipalProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterServicePrincipalProfileResponseOutput).ToManagedClusterServicePrincipalProfileResponsePtrOutputWithContext(ctx)
}









type ManagedClusterServicePrincipalProfileResponsePtrInput interface {
	pulumi.Input

	ToManagedClusterServicePrincipalProfileResponsePtrOutput() ManagedClusterServicePrincipalProfileResponsePtrOutput
	ToManagedClusterServicePrincipalProfileResponsePtrOutputWithContext(context.Context) ManagedClusterServicePrincipalProfileResponsePtrOutput
}

type managedClusterServicePrincipalProfileResponsePtrType ManagedClusterServicePrincipalProfileResponseArgs

func ManagedClusterServicePrincipalProfileResponsePtr(v *ManagedClusterServicePrincipalProfileResponseArgs) ManagedClusterServicePrincipalProfileResponsePtrInput {
	return (*managedClusterServicePrincipalProfileResponsePtrType)(v)
}

func (*managedClusterServicePrincipalProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterServicePrincipalProfileResponse)(nil)).Elem()
}

func (i *managedClusterServicePrincipalProfileResponsePtrType) ToManagedClusterServicePrincipalProfileResponsePtrOutput() ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return i.ToManagedClusterServicePrincipalProfileResponsePtrOutputWithContext(context.Background())
}

func (i *managedClusterServicePrincipalProfileResponsePtrType) ToManagedClusterServicePrincipalProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterServicePrincipalProfileResponsePtrOutput)
}

type ManagedClusterServicePrincipalProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterServicePrincipalProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterServicePrincipalProfileResponse)(nil)).Elem()
}

func (o ManagedClusterServicePrincipalProfileResponseOutput) ToManagedClusterServicePrincipalProfileResponseOutput() ManagedClusterServicePrincipalProfileResponseOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfileResponseOutput) ToManagedClusterServicePrincipalProfileResponseOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfileResponseOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfileResponseOutput) ToManagedClusterServicePrincipalProfileResponsePtrOutput() ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return o.ToManagedClusterServicePrincipalProfileResponsePtrOutputWithContext(context.Background())
}

func (o ManagedClusterServicePrincipalProfileResponseOutput) ToManagedClusterServicePrincipalProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterServicePrincipalProfileResponse) *ManagedClusterServicePrincipalProfileResponse {
		return &v
	}).(ManagedClusterServicePrincipalProfileResponsePtrOutput)
}

func (o ManagedClusterServicePrincipalProfileResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterServicePrincipalProfileResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ManagedClusterServicePrincipalProfileResponseOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterServicePrincipalProfileResponse) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

type ManagedClusterServicePrincipalProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterServicePrincipalProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterServicePrincipalProfileResponse)(nil)).Elem()
}

func (o ManagedClusterServicePrincipalProfileResponsePtrOutput) ToManagedClusterServicePrincipalProfileResponsePtrOutput() ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfileResponsePtrOutput) ToManagedClusterServicePrincipalProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfileResponsePtrOutput) Elem() ManagedClusterServicePrincipalProfileResponseOutput {
	return o.ApplyT(func(v *ManagedClusterServicePrincipalProfileResponse) ManagedClusterServicePrincipalProfileResponse {
		if v != nil {
			return *v
		}
		var ret ManagedClusterServicePrincipalProfileResponse
		return ret
	}).(ManagedClusterServicePrincipalProfileResponseOutput)
}

func (o ManagedClusterServicePrincipalProfileResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterServicePrincipalProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterServicePrincipalProfileResponsePtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterServicePrincipalProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Secret
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterWindowsProfile struct {
	AdminPassword  *string `pulumi:"adminPassword"`
	AdminUsername  string  `pulumi:"adminUsername"`
	EnableCSIProxy *bool   `pulumi:"enableCSIProxy"`
	LicenseType    *string `pulumi:"licenseType"`
}





type ManagedClusterWindowsProfileInput interface {
	pulumi.Input

	ToManagedClusterWindowsProfileOutput() ManagedClusterWindowsProfileOutput
	ToManagedClusterWindowsProfileOutputWithContext(context.Context) ManagedClusterWindowsProfileOutput
}

type ManagedClusterWindowsProfileArgs struct {
	AdminPassword  pulumi.StringPtrInput `pulumi:"adminPassword"`
	AdminUsername  pulumi.StringInput    `pulumi:"adminUsername"`
	EnableCSIProxy pulumi.BoolPtrInput   `pulumi:"enableCSIProxy"`
	LicenseType    pulumi.StringPtrInput `pulumi:"licenseType"`
}

func (ManagedClusterWindowsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterWindowsProfile)(nil)).Elem()
}

func (i ManagedClusterWindowsProfileArgs) ToManagedClusterWindowsProfileOutput() ManagedClusterWindowsProfileOutput {
	return i.ToManagedClusterWindowsProfileOutputWithContext(context.Background())
}

func (i ManagedClusterWindowsProfileArgs) ToManagedClusterWindowsProfileOutputWithContext(ctx context.Context) ManagedClusterWindowsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterWindowsProfileOutput)
}

func (i ManagedClusterWindowsProfileArgs) ToManagedClusterWindowsProfilePtrOutput() ManagedClusterWindowsProfilePtrOutput {
	return i.ToManagedClusterWindowsProfilePtrOutputWithContext(context.Background())
}

func (i ManagedClusterWindowsProfileArgs) ToManagedClusterWindowsProfilePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterWindowsProfileOutput).ToManagedClusterWindowsProfilePtrOutputWithContext(ctx)
}









type ManagedClusterWindowsProfilePtrInput interface {
	pulumi.Input

	ToManagedClusterWindowsProfilePtrOutput() ManagedClusterWindowsProfilePtrOutput
	ToManagedClusterWindowsProfilePtrOutputWithContext(context.Context) ManagedClusterWindowsProfilePtrOutput
}

type managedClusterWindowsProfilePtrType ManagedClusterWindowsProfileArgs

func ManagedClusterWindowsProfilePtr(v *ManagedClusterWindowsProfileArgs) ManagedClusterWindowsProfilePtrInput {
	return (*managedClusterWindowsProfilePtrType)(v)
}

func (*managedClusterWindowsProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterWindowsProfile)(nil)).Elem()
}

func (i *managedClusterWindowsProfilePtrType) ToManagedClusterWindowsProfilePtrOutput() ManagedClusterWindowsProfilePtrOutput {
	return i.ToManagedClusterWindowsProfilePtrOutputWithContext(context.Background())
}

func (i *managedClusterWindowsProfilePtrType) ToManagedClusterWindowsProfilePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterWindowsProfilePtrOutput)
}

type ManagedClusterWindowsProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterWindowsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterWindowsProfile)(nil)).Elem()
}

func (o ManagedClusterWindowsProfileOutput) ToManagedClusterWindowsProfileOutput() ManagedClusterWindowsProfileOutput {
	return o
}

func (o ManagedClusterWindowsProfileOutput) ToManagedClusterWindowsProfileOutputWithContext(ctx context.Context) ManagedClusterWindowsProfileOutput {
	return o
}

func (o ManagedClusterWindowsProfileOutput) ToManagedClusterWindowsProfilePtrOutput() ManagedClusterWindowsProfilePtrOutput {
	return o.ToManagedClusterWindowsProfilePtrOutputWithContext(context.Background())
}

func (o ManagedClusterWindowsProfileOutput) ToManagedClusterWindowsProfilePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterWindowsProfile) *ManagedClusterWindowsProfile {
		return &v
	}).(ManagedClusterWindowsProfilePtrOutput)
}

func (o ManagedClusterWindowsProfileOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterWindowsProfile) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterWindowsProfileOutput) AdminUsername() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterWindowsProfile) string { return v.AdminUsername }).(pulumi.StringOutput)
}

func (o ManagedClusterWindowsProfileOutput) EnableCSIProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterWindowsProfile) *bool { return v.EnableCSIProxy }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterWindowsProfileOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterWindowsProfile) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

type ManagedClusterWindowsProfilePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterWindowsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterWindowsProfile)(nil)).Elem()
}

func (o ManagedClusterWindowsProfilePtrOutput) ToManagedClusterWindowsProfilePtrOutput() ManagedClusterWindowsProfilePtrOutput {
	return o
}

func (o ManagedClusterWindowsProfilePtrOutput) ToManagedClusterWindowsProfilePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfilePtrOutput {
	return o
}

func (o ManagedClusterWindowsProfilePtrOutput) Elem() ManagedClusterWindowsProfileOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfile) ManagedClusterWindowsProfile {
		if v != nil {
			return *v
		}
		var ret ManagedClusterWindowsProfile
		return ret
	}).(ManagedClusterWindowsProfileOutput)
}

func (o ManagedClusterWindowsProfilePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterWindowsProfilePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfile) *string {
		if v == nil {
			return nil
		}
		return &v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterWindowsProfilePtrOutput) EnableCSIProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfile) *bool {
		if v == nil {
			return nil
		}
		return v.EnableCSIProxy
	}).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterWindowsProfilePtrOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfile) *string {
		if v == nil {
			return nil
		}
		return v.LicenseType
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterWindowsProfileResponse struct {
	AdminPassword  *string `pulumi:"adminPassword"`
	AdminUsername  string  `pulumi:"adminUsername"`
	EnableCSIProxy *bool   `pulumi:"enableCSIProxy"`
	LicenseType    *string `pulumi:"licenseType"`
}





type ManagedClusterWindowsProfileResponseInput interface {
	pulumi.Input

	ToManagedClusterWindowsProfileResponseOutput() ManagedClusterWindowsProfileResponseOutput
	ToManagedClusterWindowsProfileResponseOutputWithContext(context.Context) ManagedClusterWindowsProfileResponseOutput
}

type ManagedClusterWindowsProfileResponseArgs struct {
	AdminPassword  pulumi.StringPtrInput `pulumi:"adminPassword"`
	AdminUsername  pulumi.StringInput    `pulumi:"adminUsername"`
	EnableCSIProxy pulumi.BoolPtrInput   `pulumi:"enableCSIProxy"`
	LicenseType    pulumi.StringPtrInput `pulumi:"licenseType"`
}

func (ManagedClusterWindowsProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterWindowsProfileResponse)(nil)).Elem()
}

func (i ManagedClusterWindowsProfileResponseArgs) ToManagedClusterWindowsProfileResponseOutput() ManagedClusterWindowsProfileResponseOutput {
	return i.ToManagedClusterWindowsProfileResponseOutputWithContext(context.Background())
}

func (i ManagedClusterWindowsProfileResponseArgs) ToManagedClusterWindowsProfileResponseOutputWithContext(ctx context.Context) ManagedClusterWindowsProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterWindowsProfileResponseOutput)
}

func (i ManagedClusterWindowsProfileResponseArgs) ToManagedClusterWindowsProfileResponsePtrOutput() ManagedClusterWindowsProfileResponsePtrOutput {
	return i.ToManagedClusterWindowsProfileResponsePtrOutputWithContext(context.Background())
}

func (i ManagedClusterWindowsProfileResponseArgs) ToManagedClusterWindowsProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterWindowsProfileResponseOutput).ToManagedClusterWindowsProfileResponsePtrOutputWithContext(ctx)
}









type ManagedClusterWindowsProfileResponsePtrInput interface {
	pulumi.Input

	ToManagedClusterWindowsProfileResponsePtrOutput() ManagedClusterWindowsProfileResponsePtrOutput
	ToManagedClusterWindowsProfileResponsePtrOutputWithContext(context.Context) ManagedClusterWindowsProfileResponsePtrOutput
}

type managedClusterWindowsProfileResponsePtrType ManagedClusterWindowsProfileResponseArgs

func ManagedClusterWindowsProfileResponsePtr(v *ManagedClusterWindowsProfileResponseArgs) ManagedClusterWindowsProfileResponsePtrInput {
	return (*managedClusterWindowsProfileResponsePtrType)(v)
}

func (*managedClusterWindowsProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterWindowsProfileResponse)(nil)).Elem()
}

func (i *managedClusterWindowsProfileResponsePtrType) ToManagedClusterWindowsProfileResponsePtrOutput() ManagedClusterWindowsProfileResponsePtrOutput {
	return i.ToManagedClusterWindowsProfileResponsePtrOutputWithContext(context.Background())
}

func (i *managedClusterWindowsProfileResponsePtrType) ToManagedClusterWindowsProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterWindowsProfileResponsePtrOutput)
}

type ManagedClusterWindowsProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterWindowsProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterWindowsProfileResponse)(nil)).Elem()
}

func (o ManagedClusterWindowsProfileResponseOutput) ToManagedClusterWindowsProfileResponseOutput() ManagedClusterWindowsProfileResponseOutput {
	return o
}

func (o ManagedClusterWindowsProfileResponseOutput) ToManagedClusterWindowsProfileResponseOutputWithContext(ctx context.Context) ManagedClusterWindowsProfileResponseOutput {
	return o
}

func (o ManagedClusterWindowsProfileResponseOutput) ToManagedClusterWindowsProfileResponsePtrOutput() ManagedClusterWindowsProfileResponsePtrOutput {
	return o.ToManagedClusterWindowsProfileResponsePtrOutputWithContext(context.Background())
}

func (o ManagedClusterWindowsProfileResponseOutput) ToManagedClusterWindowsProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterWindowsProfileResponse) *ManagedClusterWindowsProfileResponse {
		return &v
	}).(ManagedClusterWindowsProfileResponsePtrOutput)
}

func (o ManagedClusterWindowsProfileResponseOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterWindowsProfileResponse) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterWindowsProfileResponseOutput) AdminUsername() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterWindowsProfileResponse) string { return v.AdminUsername }).(pulumi.StringOutput)
}

func (o ManagedClusterWindowsProfileResponseOutput) EnableCSIProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterWindowsProfileResponse) *bool { return v.EnableCSIProxy }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterWindowsProfileResponseOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterWindowsProfileResponse) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

type ManagedClusterWindowsProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterWindowsProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterWindowsProfileResponse)(nil)).Elem()
}

func (o ManagedClusterWindowsProfileResponsePtrOutput) ToManagedClusterWindowsProfileResponsePtrOutput() ManagedClusterWindowsProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterWindowsProfileResponsePtrOutput) ToManagedClusterWindowsProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterWindowsProfileResponsePtrOutput) Elem() ManagedClusterWindowsProfileResponseOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfileResponse) ManagedClusterWindowsProfileResponse {
		if v != nil {
			return *v
		}
		var ret ManagedClusterWindowsProfileResponse
		return ret
	}).(ManagedClusterWindowsProfileResponseOutput)
}

func (o ManagedClusterWindowsProfileResponsePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterWindowsProfileResponsePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterWindowsProfileResponsePtrOutput) EnableCSIProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfileResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableCSIProxy
	}).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterWindowsProfileResponsePtrOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.LicenseType
	}).(pulumi.StringPtrOutput)
}

type PowerStateResponse struct {
	Code *string `pulumi:"code"`
}





type PowerStateResponseInput interface {
	pulumi.Input

	ToPowerStateResponseOutput() PowerStateResponseOutput
	ToPowerStateResponseOutputWithContext(context.Context) PowerStateResponseOutput
}

type PowerStateResponseArgs struct {
	Code pulumi.StringPtrInput `pulumi:"code"`
}

func (PowerStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PowerStateResponse)(nil)).Elem()
}

func (i PowerStateResponseArgs) ToPowerStateResponseOutput() PowerStateResponseOutput {
	return i.ToPowerStateResponseOutputWithContext(context.Background())
}

func (i PowerStateResponseArgs) ToPowerStateResponseOutputWithContext(ctx context.Context) PowerStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PowerStateResponseOutput)
}

func (i PowerStateResponseArgs) ToPowerStateResponsePtrOutput() PowerStateResponsePtrOutput {
	return i.ToPowerStateResponsePtrOutputWithContext(context.Background())
}

func (i PowerStateResponseArgs) ToPowerStateResponsePtrOutputWithContext(ctx context.Context) PowerStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PowerStateResponseOutput).ToPowerStateResponsePtrOutputWithContext(ctx)
}









type PowerStateResponsePtrInput interface {
	pulumi.Input

	ToPowerStateResponsePtrOutput() PowerStateResponsePtrOutput
	ToPowerStateResponsePtrOutputWithContext(context.Context) PowerStateResponsePtrOutput
}

type powerStateResponsePtrType PowerStateResponseArgs

func PowerStateResponsePtr(v *PowerStateResponseArgs) PowerStateResponsePtrInput {
	return (*powerStateResponsePtrType)(v)
}

func (*powerStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PowerStateResponse)(nil)).Elem()
}

func (i *powerStateResponsePtrType) ToPowerStateResponsePtrOutput() PowerStateResponsePtrOutput {
	return i.ToPowerStateResponsePtrOutputWithContext(context.Background())
}

func (i *powerStateResponsePtrType) ToPowerStateResponsePtrOutputWithContext(ctx context.Context) PowerStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PowerStateResponsePtrOutput)
}

type PowerStateResponseOutput struct{ *pulumi.OutputState }

func (PowerStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PowerStateResponse)(nil)).Elem()
}

func (o PowerStateResponseOutput) ToPowerStateResponseOutput() PowerStateResponseOutput {
	return o
}

func (o PowerStateResponseOutput) ToPowerStateResponseOutputWithContext(ctx context.Context) PowerStateResponseOutput {
	return o
}

func (o PowerStateResponseOutput) ToPowerStateResponsePtrOutput() PowerStateResponsePtrOutput {
	return o.ToPowerStateResponsePtrOutputWithContext(context.Background())
}

func (o PowerStateResponseOutput) ToPowerStateResponsePtrOutputWithContext(ctx context.Context) PowerStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PowerStateResponse) *PowerStateResponse {
		return &v
	}).(PowerStateResponsePtrOutput)
}

func (o PowerStateResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PowerStateResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

type PowerStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PowerStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PowerStateResponse)(nil)).Elem()
}

func (o PowerStateResponsePtrOutput) ToPowerStateResponsePtrOutput() PowerStateResponsePtrOutput {
	return o
}

func (o PowerStateResponsePtrOutput) ToPowerStateResponsePtrOutputWithContext(ctx context.Context) PowerStateResponsePtrOutput {
	return o
}

func (o PowerStateResponsePtrOutput) Elem() PowerStateResponseOutput {
	return o.ApplyT(func(v *PowerStateResponse) PowerStateResponse {
		if v != nil {
			return *v
		}
		var ret PowerStateResponse
		return ret
	}).(PowerStateResponseOutput)
}

func (o PowerStateResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PowerStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpoint struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointInput interface {
	pulumi.Input

	ToPrivateEndpointOutput() PrivateEndpointOutput
	ToPrivateEndpointOutputWithContext(context.Context) PrivateEndpointOutput
}

type PrivateEndpointArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return i.ToPrivateEndpointOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput)
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput).ToPrivateEndpointPtrOutputWithContext(ctx)
}









type PrivateEndpointPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput
	ToPrivateEndpointPtrOutputWithContext(context.Context) PrivateEndpointPtrOutput
}

type privateEndpointPtrType PrivateEndpointArgs

func PrivateEndpointPtr(v *PrivateEndpointArgs) PrivateEndpointPtrInput {
	return (*privateEndpointPtrType)(v)
}

func (*privateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPtrOutput)
}

type PrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpoint) *PrivateEndpoint {
		return &v
	}).(PrivateEndpointPtrOutput)
}

func (o PrivateEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) Elem() PrivateEndpointOutput {
	return o.ApplyT(func(v *PrivateEndpoint) PrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret PrivateEndpoint
		return ret
	}).(PrivateEndpointOutput)
}

func (o PrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointResponse struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointResponseInput interface {
	pulumi.Input

	ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput
	ToPrivateEndpointResponseOutputWithContext(context.Context) PrivateEndpointResponseOutput
}

type PrivateEndpointResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return i.ToPrivateEndpointResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput)
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput).ToPrivateEndpointResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput
	ToPrivateEndpointResponsePtrOutputWithContext(context.Context) PrivateEndpointResponsePtrOutput
}

type privateEndpointResponsePtrType PrivateEndpointResponseArgs

func PrivateEndpointResponsePtr(v *PrivateEndpointResponseArgs) PrivateEndpointResponsePtrInput {
	return (*privateEndpointResponsePtrType)(v)
}

func (*privateEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponsePtrOutput)
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

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointResponse) *PrivateEndpointResponse {
		return &v
	}).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
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
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkResource struct {
	GroupId         *string  `pulumi:"groupId"`
	Id              *string  `pulumi:"id"`
	Name            *string  `pulumi:"name"`
	RequiredMembers []string `pulumi:"requiredMembers"`
	Type            *string  `pulumi:"type"`
}





type PrivateLinkResourceInput interface {
	pulumi.Input

	ToPrivateLinkResourceOutput() PrivateLinkResourceOutput
	ToPrivateLinkResourceOutputWithContext(context.Context) PrivateLinkResourceOutput
}

type PrivateLinkResourceArgs struct {
	GroupId         pulumi.StringPtrInput   `pulumi:"groupId"`
	Id              pulumi.StringPtrInput   `pulumi:"id"`
	Name            pulumi.StringPtrInput   `pulumi:"name"`
	RequiredMembers pulumi.StringArrayInput `pulumi:"requiredMembers"`
	Type            pulumi.StringPtrInput   `pulumi:"type"`
}

func (PrivateLinkResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkResource)(nil)).Elem()
}

func (i PrivateLinkResourceArgs) ToPrivateLinkResourceOutput() PrivateLinkResourceOutput {
	return i.ToPrivateLinkResourceOutputWithContext(context.Background())
}

func (i PrivateLinkResourceArgs) ToPrivateLinkResourceOutputWithContext(ctx context.Context) PrivateLinkResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkResourceOutput)
}





type PrivateLinkResourceArrayInput interface {
	pulumi.Input

	ToPrivateLinkResourceArrayOutput() PrivateLinkResourceArrayOutput
	ToPrivateLinkResourceArrayOutputWithContext(context.Context) PrivateLinkResourceArrayOutput
}

type PrivateLinkResourceArray []PrivateLinkResourceInput

func (PrivateLinkResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkResource)(nil)).Elem()
}

func (i PrivateLinkResourceArray) ToPrivateLinkResourceArrayOutput() PrivateLinkResourceArrayOutput {
	return i.ToPrivateLinkResourceArrayOutputWithContext(context.Background())
}

func (i PrivateLinkResourceArray) ToPrivateLinkResourceArrayOutputWithContext(ctx context.Context) PrivateLinkResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkResourceArrayOutput)
}

type PrivateLinkResourceOutput struct{ *pulumi.OutputState }

func (PrivateLinkResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkResource)(nil)).Elem()
}

func (o PrivateLinkResourceOutput) ToPrivateLinkResourceOutput() PrivateLinkResourceOutput {
	return o
}

func (o PrivateLinkResourceOutput) ToPrivateLinkResourceOutputWithContext(ctx context.Context) PrivateLinkResourceOutput {
	return o
}

func (o PrivateLinkResourceOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkResource) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkResourceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkResource) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkResourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkResource) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkResourceOutput) RequiredMembers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateLinkResource) []string { return v.RequiredMembers }).(pulumi.StringArrayOutput)
}

func (o PrivateLinkResourceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkResource) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PrivateLinkResourceArrayOutput struct{ *pulumi.OutputState }

func (PrivateLinkResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkResource)(nil)).Elem()
}

func (o PrivateLinkResourceArrayOutput) ToPrivateLinkResourceArrayOutput() PrivateLinkResourceArrayOutput {
	return o
}

func (o PrivateLinkResourceArrayOutput) ToPrivateLinkResourceArrayOutputWithContext(ctx context.Context) PrivateLinkResourceArrayOutput {
	return o
}

func (o PrivateLinkResourceArrayOutput) Index(i pulumi.IntInput) PrivateLinkResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateLinkResource {
		return vs[0].([]PrivateLinkResource)[vs[1].(int)]
	}).(PrivateLinkResourceOutput)
}

type PrivateLinkResourceResponse struct {
	GroupId              *string  `pulumi:"groupId"`
	Id                   *string  `pulumi:"id"`
	Name                 *string  `pulumi:"name"`
	PrivateLinkServiceID string   `pulumi:"privateLinkServiceID"`
	RequiredMembers      []string `pulumi:"requiredMembers"`
	Type                 *string  `pulumi:"type"`
}





type PrivateLinkResourceResponseInput interface {
	pulumi.Input

	ToPrivateLinkResourceResponseOutput() PrivateLinkResourceResponseOutput
	ToPrivateLinkResourceResponseOutputWithContext(context.Context) PrivateLinkResourceResponseOutput
}

type PrivateLinkResourceResponseArgs struct {
	GroupId              pulumi.StringPtrInput   `pulumi:"groupId"`
	Id                   pulumi.StringPtrInput   `pulumi:"id"`
	Name                 pulumi.StringPtrInput   `pulumi:"name"`
	PrivateLinkServiceID pulumi.StringInput      `pulumi:"privateLinkServiceID"`
	RequiredMembers      pulumi.StringArrayInput `pulumi:"requiredMembers"`
	Type                 pulumi.StringPtrInput   `pulumi:"type"`
}

func (PrivateLinkResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkResourceResponse)(nil)).Elem()
}

func (i PrivateLinkResourceResponseArgs) ToPrivateLinkResourceResponseOutput() PrivateLinkResourceResponseOutput {
	return i.ToPrivateLinkResourceResponseOutputWithContext(context.Background())
}

func (i PrivateLinkResourceResponseArgs) ToPrivateLinkResourceResponseOutputWithContext(ctx context.Context) PrivateLinkResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkResourceResponseOutput)
}





type PrivateLinkResourceResponseArrayInput interface {
	pulumi.Input

	ToPrivateLinkResourceResponseArrayOutput() PrivateLinkResourceResponseArrayOutput
	ToPrivateLinkResourceResponseArrayOutputWithContext(context.Context) PrivateLinkResourceResponseArrayOutput
}

type PrivateLinkResourceResponseArray []PrivateLinkResourceResponseInput

func (PrivateLinkResourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkResourceResponse)(nil)).Elem()
}

func (i PrivateLinkResourceResponseArray) ToPrivateLinkResourceResponseArrayOutput() PrivateLinkResourceResponseArrayOutput {
	return i.ToPrivateLinkResourceResponseArrayOutputWithContext(context.Background())
}

func (i PrivateLinkResourceResponseArray) ToPrivateLinkResourceResponseArrayOutputWithContext(ctx context.Context) PrivateLinkResourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkResourceResponseArrayOutput)
}

type PrivateLinkResourceResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkResourceResponse)(nil)).Elem()
}

func (o PrivateLinkResourceResponseOutput) ToPrivateLinkResourceResponseOutput() PrivateLinkResourceResponseOutput {
	return o
}

func (o PrivateLinkResourceResponseOutput) ToPrivateLinkResourceResponseOutputWithContext(ctx context.Context) PrivateLinkResourceResponseOutput {
	return o
}

func (o PrivateLinkResourceResponseOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkResourceResponse) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkResourceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkResourceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkResourceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkResourceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkResourceResponseOutput) PrivateLinkServiceID() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkResourceResponse) string { return v.PrivateLinkServiceID }).(pulumi.StringOutput)
}

func (o PrivateLinkResourceResponseOutput) RequiredMembers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateLinkResourceResponse) []string { return v.RequiredMembers }).(pulumi.StringArrayOutput)
}

func (o PrivateLinkResourceResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkResourceResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PrivateLinkResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateLinkResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkResourceResponse)(nil)).Elem()
}

func (o PrivateLinkResourceResponseArrayOutput) ToPrivateLinkResourceResponseArrayOutput() PrivateLinkResourceResponseArrayOutput {
	return o
}

func (o PrivateLinkResourceResponseArrayOutput) ToPrivateLinkResourceResponseArrayOutputWithContext(ctx context.Context) PrivateLinkResourceResponseArrayOutput {
	return o
}

func (o PrivateLinkResourceResponseArrayOutput) Index(i pulumi.IntInput) PrivateLinkResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateLinkResourceResponse {
		return vs[0].([]PrivateLinkResourceResponse)[vs[1].(int)]
	}).(PrivateLinkResourceResponseOutput)
}

type PrivateLinkServiceConnectionState struct {
	Description *string `pulumi:"description"`
	Status      *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Status      pulumi.StringPtrInput `pulumi:"status"`
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

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput).ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput
	ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePtrOutput
}

type privateLinkServiceConnectionStatePtrType PrivateLinkServiceConnectionStateArgs

func PrivateLinkServiceConnectionStatePtr(v *PrivateLinkServiceConnectionStateArgs) PrivateLinkServiceConnectionStatePtrInput {
	return (*privateLinkServiceConnectionStatePtrType)(v)
}

func (*privateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePtrOutput)
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

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionState {
		return &v
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) PrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionState
		return ret
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	Description *string `pulumi:"description"`
	Status      *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput
	ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponseOutput
}

type PrivateLinkServiceConnectionStateResponseArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Status      pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return i.ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput).ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStateResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput
	ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput
}

type privateLinkServiceConnectionStateResponsePtrType PrivateLinkServiceConnectionStateResponseArgs

func PrivateLinkServiceConnectionStateResponsePtr(v *PrivateLinkServiceConnectionStateResponseArgs) PrivateLinkServiceConnectionStateResponsePtrInput {
	return (*privateLinkServiceConnectionStateResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponsePtrOutput)
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

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateResponse) *PrivateLinkServiceConnectionStateResponse {
		return &v
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Elem() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) PrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateResponse
		return ret
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ResourceReference struct {
	Id *string `pulumi:"id"`
}





type ResourceReferenceInput interface {
	pulumi.Input

	ToResourceReferenceOutput() ResourceReferenceOutput
	ToResourceReferenceOutputWithContext(context.Context) ResourceReferenceOutput
}

type ResourceReferenceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ResourceReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReference)(nil)).Elem()
}

func (i ResourceReferenceArgs) ToResourceReferenceOutput() ResourceReferenceOutput {
	return i.ToResourceReferenceOutputWithContext(context.Background())
}

func (i ResourceReferenceArgs) ToResourceReferenceOutputWithContext(ctx context.Context) ResourceReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceOutput)
}





type ResourceReferenceArrayInput interface {
	pulumi.Input

	ToResourceReferenceArrayOutput() ResourceReferenceArrayOutput
	ToResourceReferenceArrayOutputWithContext(context.Context) ResourceReferenceArrayOutput
}

type ResourceReferenceArray []ResourceReferenceInput

func (ResourceReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceReference)(nil)).Elem()
}

func (i ResourceReferenceArray) ToResourceReferenceArrayOutput() ResourceReferenceArrayOutput {
	return i.ToResourceReferenceArrayOutputWithContext(context.Background())
}

func (i ResourceReferenceArray) ToResourceReferenceArrayOutputWithContext(ctx context.Context) ResourceReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceArrayOutput)
}

type ResourceReferenceOutput struct{ *pulumi.OutputState }

func (ResourceReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReference)(nil)).Elem()
}

func (o ResourceReferenceOutput) ToResourceReferenceOutput() ResourceReferenceOutput {
	return o
}

func (o ResourceReferenceOutput) ToResourceReferenceOutputWithContext(ctx context.Context) ResourceReferenceOutput {
	return o
}

func (o ResourceReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ResourceReferenceArrayOutput struct{ *pulumi.OutputState }

func (ResourceReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceReference)(nil)).Elem()
}

func (o ResourceReferenceArrayOutput) ToResourceReferenceArrayOutput() ResourceReferenceArrayOutput {
	return o
}

func (o ResourceReferenceArrayOutput) ToResourceReferenceArrayOutputWithContext(ctx context.Context) ResourceReferenceArrayOutput {
	return o
}

func (o ResourceReferenceArrayOutput) Index(i pulumi.IntInput) ResourceReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceReference {
		return vs[0].([]ResourceReference)[vs[1].(int)]
	}).(ResourceReferenceOutput)
}

type ResourceReferenceResponse struct {
	Id *string `pulumi:"id"`
}





type ResourceReferenceResponseInput interface {
	pulumi.Input

	ToResourceReferenceResponseOutput() ResourceReferenceResponseOutput
	ToResourceReferenceResponseOutputWithContext(context.Context) ResourceReferenceResponseOutput
}

type ResourceReferenceResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ResourceReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReferenceResponse)(nil)).Elem()
}

func (i ResourceReferenceResponseArgs) ToResourceReferenceResponseOutput() ResourceReferenceResponseOutput {
	return i.ToResourceReferenceResponseOutputWithContext(context.Background())
}

func (i ResourceReferenceResponseArgs) ToResourceReferenceResponseOutputWithContext(ctx context.Context) ResourceReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceResponseOutput)
}





type ResourceReferenceResponseArrayInput interface {
	pulumi.Input

	ToResourceReferenceResponseArrayOutput() ResourceReferenceResponseArrayOutput
	ToResourceReferenceResponseArrayOutputWithContext(context.Context) ResourceReferenceResponseArrayOutput
}

type ResourceReferenceResponseArray []ResourceReferenceResponseInput

func (ResourceReferenceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceReferenceResponse)(nil)).Elem()
}

func (i ResourceReferenceResponseArray) ToResourceReferenceResponseArrayOutput() ResourceReferenceResponseArrayOutput {
	return i.ToResourceReferenceResponseArrayOutputWithContext(context.Background())
}

func (i ResourceReferenceResponseArray) ToResourceReferenceResponseArrayOutputWithContext(ctx context.Context) ResourceReferenceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceResponseArrayOutput)
}

type ResourceReferenceResponseOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponseOutput() ResourceReferenceResponseOutput {
	return o
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponseOutputWithContext(ctx context.Context) ResourceReferenceResponseOutput {
	return o
}

func (o ResourceReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ResourceReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponseArrayOutput) ToResourceReferenceResponseArrayOutput() ResourceReferenceResponseArrayOutput {
	return o
}

func (o ResourceReferenceResponseArrayOutput) ToResourceReferenceResponseArrayOutputWithContext(ctx context.Context) ResourceReferenceResponseArrayOutput {
	return o
}

func (o ResourceReferenceResponseArrayOutput) Index(i pulumi.IntInput) ResourceReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceReferenceResponse {
		return vs[0].([]ResourceReferenceResponse)[vs[1].(int)]
	}).(ResourceReferenceResponseOutput)
}

type SysctlConfig struct {
	FsAioMaxNr                     *int    `pulumi:"fsAioMaxNr"`
	FsFileMax                      *int    `pulumi:"fsFileMax"`
	FsInotifyMaxUserWatches        *int    `pulumi:"fsInotifyMaxUserWatches"`
	FsNrOpen                       *int    `pulumi:"fsNrOpen"`
	KernelThreadsMax               *int    `pulumi:"kernelThreadsMax"`
	NetCoreNetdevMaxBacklog        *int    `pulumi:"netCoreNetdevMaxBacklog"`
	NetCoreOptmemMax               *int    `pulumi:"netCoreOptmemMax"`
	NetCoreRmemDefault             *int    `pulumi:"netCoreRmemDefault"`
	NetCoreRmemMax                 *int    `pulumi:"netCoreRmemMax"`
	NetCoreSomaxconn               *int    `pulumi:"netCoreSomaxconn"`
	NetCoreWmemDefault             *int    `pulumi:"netCoreWmemDefault"`
	NetCoreWmemMax                 *int    `pulumi:"netCoreWmemMax"`
	NetIpv4IpLocalPortRange        *string `pulumi:"netIpv4IpLocalPortRange"`
	NetIpv4NeighDefaultGcThresh1   *int    `pulumi:"netIpv4NeighDefaultGcThresh1"`
	NetIpv4NeighDefaultGcThresh2   *int    `pulumi:"netIpv4NeighDefaultGcThresh2"`
	NetIpv4NeighDefaultGcThresh3   *int    `pulumi:"netIpv4NeighDefaultGcThresh3"`
	NetIpv4TcpFinTimeout           *int    `pulumi:"netIpv4TcpFinTimeout"`
	NetIpv4TcpKeepaliveProbes      *int    `pulumi:"netIpv4TcpKeepaliveProbes"`
	NetIpv4TcpKeepaliveTime        *int    `pulumi:"netIpv4TcpKeepaliveTime"`
	NetIpv4TcpMaxSynBacklog        *int    `pulumi:"netIpv4TcpMaxSynBacklog"`
	NetIpv4TcpMaxTwBuckets         *int    `pulumi:"netIpv4TcpMaxTwBuckets"`
	NetIpv4TcpTwReuse              *bool   `pulumi:"netIpv4TcpTwReuse"`
	NetIpv4TcpkeepaliveIntvl       *int    `pulumi:"netIpv4TcpkeepaliveIntvl"`
	NetNetfilterNfConntrackBuckets *int    `pulumi:"netNetfilterNfConntrackBuckets"`
	NetNetfilterNfConntrackMax     *int    `pulumi:"netNetfilterNfConntrackMax"`
	VmMaxMapCount                  *int    `pulumi:"vmMaxMapCount"`
	VmSwappiness                   *int    `pulumi:"vmSwappiness"`
	VmVfsCachePressure             *int    `pulumi:"vmVfsCachePressure"`
}





type SysctlConfigInput interface {
	pulumi.Input

	ToSysctlConfigOutput() SysctlConfigOutput
	ToSysctlConfigOutputWithContext(context.Context) SysctlConfigOutput
}

type SysctlConfigArgs struct {
	FsAioMaxNr                     pulumi.IntPtrInput    `pulumi:"fsAioMaxNr"`
	FsFileMax                      pulumi.IntPtrInput    `pulumi:"fsFileMax"`
	FsInotifyMaxUserWatches        pulumi.IntPtrInput    `pulumi:"fsInotifyMaxUserWatches"`
	FsNrOpen                       pulumi.IntPtrInput    `pulumi:"fsNrOpen"`
	KernelThreadsMax               pulumi.IntPtrInput    `pulumi:"kernelThreadsMax"`
	NetCoreNetdevMaxBacklog        pulumi.IntPtrInput    `pulumi:"netCoreNetdevMaxBacklog"`
	NetCoreOptmemMax               pulumi.IntPtrInput    `pulumi:"netCoreOptmemMax"`
	NetCoreRmemDefault             pulumi.IntPtrInput    `pulumi:"netCoreRmemDefault"`
	NetCoreRmemMax                 pulumi.IntPtrInput    `pulumi:"netCoreRmemMax"`
	NetCoreSomaxconn               pulumi.IntPtrInput    `pulumi:"netCoreSomaxconn"`
	NetCoreWmemDefault             pulumi.IntPtrInput    `pulumi:"netCoreWmemDefault"`
	NetCoreWmemMax                 pulumi.IntPtrInput    `pulumi:"netCoreWmemMax"`
	NetIpv4IpLocalPortRange        pulumi.StringPtrInput `pulumi:"netIpv4IpLocalPortRange"`
	NetIpv4NeighDefaultGcThresh1   pulumi.IntPtrInput    `pulumi:"netIpv4NeighDefaultGcThresh1"`
	NetIpv4NeighDefaultGcThresh2   pulumi.IntPtrInput    `pulumi:"netIpv4NeighDefaultGcThresh2"`
	NetIpv4NeighDefaultGcThresh3   pulumi.IntPtrInput    `pulumi:"netIpv4NeighDefaultGcThresh3"`
	NetIpv4TcpFinTimeout           pulumi.IntPtrInput    `pulumi:"netIpv4TcpFinTimeout"`
	NetIpv4TcpKeepaliveProbes      pulumi.IntPtrInput    `pulumi:"netIpv4TcpKeepaliveProbes"`
	NetIpv4TcpKeepaliveTime        pulumi.IntPtrInput    `pulumi:"netIpv4TcpKeepaliveTime"`
	NetIpv4TcpMaxSynBacklog        pulumi.IntPtrInput    `pulumi:"netIpv4TcpMaxSynBacklog"`
	NetIpv4TcpMaxTwBuckets         pulumi.IntPtrInput    `pulumi:"netIpv4TcpMaxTwBuckets"`
	NetIpv4TcpTwReuse              pulumi.BoolPtrInput   `pulumi:"netIpv4TcpTwReuse"`
	NetIpv4TcpkeepaliveIntvl       pulumi.IntPtrInput    `pulumi:"netIpv4TcpkeepaliveIntvl"`
	NetNetfilterNfConntrackBuckets pulumi.IntPtrInput    `pulumi:"netNetfilterNfConntrackBuckets"`
	NetNetfilterNfConntrackMax     pulumi.IntPtrInput    `pulumi:"netNetfilterNfConntrackMax"`
	VmMaxMapCount                  pulumi.IntPtrInput    `pulumi:"vmMaxMapCount"`
	VmSwappiness                   pulumi.IntPtrInput    `pulumi:"vmSwappiness"`
	VmVfsCachePressure             pulumi.IntPtrInput    `pulumi:"vmVfsCachePressure"`
}

func (SysctlConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SysctlConfig)(nil)).Elem()
}

func (i SysctlConfigArgs) ToSysctlConfigOutput() SysctlConfigOutput {
	return i.ToSysctlConfigOutputWithContext(context.Background())
}

func (i SysctlConfigArgs) ToSysctlConfigOutputWithContext(ctx context.Context) SysctlConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SysctlConfigOutput)
}

func (i SysctlConfigArgs) ToSysctlConfigPtrOutput() SysctlConfigPtrOutput {
	return i.ToSysctlConfigPtrOutputWithContext(context.Background())
}

func (i SysctlConfigArgs) ToSysctlConfigPtrOutputWithContext(ctx context.Context) SysctlConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SysctlConfigOutput).ToSysctlConfigPtrOutputWithContext(ctx)
}









type SysctlConfigPtrInput interface {
	pulumi.Input

	ToSysctlConfigPtrOutput() SysctlConfigPtrOutput
	ToSysctlConfigPtrOutputWithContext(context.Context) SysctlConfigPtrOutput
}

type sysctlConfigPtrType SysctlConfigArgs

func SysctlConfigPtr(v *SysctlConfigArgs) SysctlConfigPtrInput {
	return (*sysctlConfigPtrType)(v)
}

func (*sysctlConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SysctlConfig)(nil)).Elem()
}

func (i *sysctlConfigPtrType) ToSysctlConfigPtrOutput() SysctlConfigPtrOutput {
	return i.ToSysctlConfigPtrOutputWithContext(context.Background())
}

func (i *sysctlConfigPtrType) ToSysctlConfigPtrOutputWithContext(ctx context.Context) SysctlConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SysctlConfigPtrOutput)
}

type SysctlConfigOutput struct{ *pulumi.OutputState }

func (SysctlConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SysctlConfig)(nil)).Elem()
}

func (o SysctlConfigOutput) ToSysctlConfigOutput() SysctlConfigOutput {
	return o
}

func (o SysctlConfigOutput) ToSysctlConfigOutputWithContext(ctx context.Context) SysctlConfigOutput {
	return o
}

func (o SysctlConfigOutput) ToSysctlConfigPtrOutput() SysctlConfigPtrOutput {
	return o.ToSysctlConfigPtrOutputWithContext(context.Background())
}

func (o SysctlConfigOutput) ToSysctlConfigPtrOutputWithContext(ctx context.Context) SysctlConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SysctlConfig) *SysctlConfig {
		return &v
	}).(SysctlConfigPtrOutput)
}

func (o SysctlConfigOutput) FsAioMaxNr() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.FsAioMaxNr }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) FsFileMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.FsFileMax }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) FsInotifyMaxUserWatches() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.FsInotifyMaxUserWatches }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) FsNrOpen() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.FsNrOpen }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) KernelThreadsMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.KernelThreadsMax }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) NetCoreNetdevMaxBacklog() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.NetCoreNetdevMaxBacklog }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) NetCoreOptmemMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.NetCoreOptmemMax }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) NetCoreRmemDefault() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.NetCoreRmemDefault }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) NetCoreRmemMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.NetCoreRmemMax }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) NetCoreSomaxconn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.NetCoreSomaxconn }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) NetCoreWmemDefault() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.NetCoreWmemDefault }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) NetCoreWmemMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.NetCoreWmemMax }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) NetIpv4IpLocalPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *string { return v.NetIpv4IpLocalPortRange }).(pulumi.StringPtrOutput)
}

func (o SysctlConfigOutput) NetIpv4NeighDefaultGcThresh1() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.NetIpv4NeighDefaultGcThresh1 }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) NetIpv4NeighDefaultGcThresh2() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.NetIpv4NeighDefaultGcThresh2 }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) NetIpv4NeighDefaultGcThresh3() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.NetIpv4NeighDefaultGcThresh3 }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) NetIpv4TcpFinTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.NetIpv4TcpFinTimeout }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) NetIpv4TcpKeepaliveProbes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.NetIpv4TcpKeepaliveProbes }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) NetIpv4TcpKeepaliveTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.NetIpv4TcpKeepaliveTime }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) NetIpv4TcpMaxSynBacklog() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.NetIpv4TcpMaxSynBacklog }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) NetIpv4TcpMaxTwBuckets() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.NetIpv4TcpMaxTwBuckets }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) NetIpv4TcpTwReuse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *bool { return v.NetIpv4TcpTwReuse }).(pulumi.BoolPtrOutput)
}

func (o SysctlConfigOutput) NetIpv4TcpkeepaliveIntvl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.NetIpv4TcpkeepaliveIntvl }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) NetNetfilterNfConntrackBuckets() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.NetNetfilterNfConntrackBuckets }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) NetNetfilterNfConntrackMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.NetNetfilterNfConntrackMax }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) VmMaxMapCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.VmMaxMapCount }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) VmSwappiness() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.VmSwappiness }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigOutput) VmVfsCachePressure() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfig) *int { return v.VmVfsCachePressure }).(pulumi.IntPtrOutput)
}

type SysctlConfigPtrOutput struct{ *pulumi.OutputState }

func (SysctlConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SysctlConfig)(nil)).Elem()
}

func (o SysctlConfigPtrOutput) ToSysctlConfigPtrOutput() SysctlConfigPtrOutput {
	return o
}

func (o SysctlConfigPtrOutput) ToSysctlConfigPtrOutputWithContext(ctx context.Context) SysctlConfigPtrOutput {
	return o
}

func (o SysctlConfigPtrOutput) Elem() SysctlConfigOutput {
	return o.ApplyT(func(v *SysctlConfig) SysctlConfig {
		if v != nil {
			return *v
		}
		var ret SysctlConfig
		return ret
	}).(SysctlConfigOutput)
}

func (o SysctlConfigPtrOutput) FsAioMaxNr() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.FsAioMaxNr
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) FsFileMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.FsFileMax
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) FsInotifyMaxUserWatches() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.FsInotifyMaxUserWatches
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) FsNrOpen() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.FsNrOpen
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) KernelThreadsMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.KernelThreadsMax
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) NetCoreNetdevMaxBacklog() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.NetCoreNetdevMaxBacklog
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) NetCoreOptmemMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.NetCoreOptmemMax
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) NetCoreRmemDefault() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.NetCoreRmemDefault
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) NetCoreRmemMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.NetCoreRmemMax
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) NetCoreSomaxconn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.NetCoreSomaxconn
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) NetCoreWmemDefault() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.NetCoreWmemDefault
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) NetCoreWmemMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.NetCoreWmemMax
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) NetIpv4IpLocalPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *string {
		if v == nil {
			return nil
		}
		return v.NetIpv4IpLocalPortRange
	}).(pulumi.StringPtrOutput)
}

func (o SysctlConfigPtrOutput) NetIpv4NeighDefaultGcThresh1() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.NetIpv4NeighDefaultGcThresh1
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) NetIpv4NeighDefaultGcThresh2() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.NetIpv4NeighDefaultGcThresh2
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) NetIpv4NeighDefaultGcThresh3() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.NetIpv4NeighDefaultGcThresh3
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) NetIpv4TcpFinTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.NetIpv4TcpFinTimeout
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) NetIpv4TcpKeepaliveProbes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.NetIpv4TcpKeepaliveProbes
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) NetIpv4TcpKeepaliveTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.NetIpv4TcpKeepaliveTime
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) NetIpv4TcpMaxSynBacklog() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.NetIpv4TcpMaxSynBacklog
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) NetIpv4TcpMaxTwBuckets() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.NetIpv4TcpMaxTwBuckets
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) NetIpv4TcpTwReuse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *bool {
		if v == nil {
			return nil
		}
		return v.NetIpv4TcpTwReuse
	}).(pulumi.BoolPtrOutput)
}

func (o SysctlConfigPtrOutput) NetIpv4TcpkeepaliveIntvl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.NetIpv4TcpkeepaliveIntvl
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) NetNetfilterNfConntrackBuckets() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.NetNetfilterNfConntrackBuckets
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) NetNetfilterNfConntrackMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.NetNetfilterNfConntrackMax
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) VmMaxMapCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.VmMaxMapCount
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) VmSwappiness() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.VmSwappiness
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigPtrOutput) VmVfsCachePressure() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfig) *int {
		if v == nil {
			return nil
		}
		return v.VmVfsCachePressure
	}).(pulumi.IntPtrOutput)
}

type SysctlConfigResponse struct {
	FsAioMaxNr                     *int    `pulumi:"fsAioMaxNr"`
	FsFileMax                      *int    `pulumi:"fsFileMax"`
	FsInotifyMaxUserWatches        *int    `pulumi:"fsInotifyMaxUserWatches"`
	FsNrOpen                       *int    `pulumi:"fsNrOpen"`
	KernelThreadsMax               *int    `pulumi:"kernelThreadsMax"`
	NetCoreNetdevMaxBacklog        *int    `pulumi:"netCoreNetdevMaxBacklog"`
	NetCoreOptmemMax               *int    `pulumi:"netCoreOptmemMax"`
	NetCoreRmemDefault             *int    `pulumi:"netCoreRmemDefault"`
	NetCoreRmemMax                 *int    `pulumi:"netCoreRmemMax"`
	NetCoreSomaxconn               *int    `pulumi:"netCoreSomaxconn"`
	NetCoreWmemDefault             *int    `pulumi:"netCoreWmemDefault"`
	NetCoreWmemMax                 *int    `pulumi:"netCoreWmemMax"`
	NetIpv4IpLocalPortRange        *string `pulumi:"netIpv4IpLocalPortRange"`
	NetIpv4NeighDefaultGcThresh1   *int    `pulumi:"netIpv4NeighDefaultGcThresh1"`
	NetIpv4NeighDefaultGcThresh2   *int    `pulumi:"netIpv4NeighDefaultGcThresh2"`
	NetIpv4NeighDefaultGcThresh3   *int    `pulumi:"netIpv4NeighDefaultGcThresh3"`
	NetIpv4TcpFinTimeout           *int    `pulumi:"netIpv4TcpFinTimeout"`
	NetIpv4TcpKeepaliveProbes      *int    `pulumi:"netIpv4TcpKeepaliveProbes"`
	NetIpv4TcpKeepaliveTime        *int    `pulumi:"netIpv4TcpKeepaliveTime"`
	NetIpv4TcpMaxSynBacklog        *int    `pulumi:"netIpv4TcpMaxSynBacklog"`
	NetIpv4TcpMaxTwBuckets         *int    `pulumi:"netIpv4TcpMaxTwBuckets"`
	NetIpv4TcpTwReuse              *bool   `pulumi:"netIpv4TcpTwReuse"`
	NetIpv4TcpkeepaliveIntvl       *int    `pulumi:"netIpv4TcpkeepaliveIntvl"`
	NetNetfilterNfConntrackBuckets *int    `pulumi:"netNetfilterNfConntrackBuckets"`
	NetNetfilterNfConntrackMax     *int    `pulumi:"netNetfilterNfConntrackMax"`
	VmMaxMapCount                  *int    `pulumi:"vmMaxMapCount"`
	VmSwappiness                   *int    `pulumi:"vmSwappiness"`
	VmVfsCachePressure             *int    `pulumi:"vmVfsCachePressure"`
}





type SysctlConfigResponseInput interface {
	pulumi.Input

	ToSysctlConfigResponseOutput() SysctlConfigResponseOutput
	ToSysctlConfigResponseOutputWithContext(context.Context) SysctlConfigResponseOutput
}

type SysctlConfigResponseArgs struct {
	FsAioMaxNr                     pulumi.IntPtrInput    `pulumi:"fsAioMaxNr"`
	FsFileMax                      pulumi.IntPtrInput    `pulumi:"fsFileMax"`
	FsInotifyMaxUserWatches        pulumi.IntPtrInput    `pulumi:"fsInotifyMaxUserWatches"`
	FsNrOpen                       pulumi.IntPtrInput    `pulumi:"fsNrOpen"`
	KernelThreadsMax               pulumi.IntPtrInput    `pulumi:"kernelThreadsMax"`
	NetCoreNetdevMaxBacklog        pulumi.IntPtrInput    `pulumi:"netCoreNetdevMaxBacklog"`
	NetCoreOptmemMax               pulumi.IntPtrInput    `pulumi:"netCoreOptmemMax"`
	NetCoreRmemDefault             pulumi.IntPtrInput    `pulumi:"netCoreRmemDefault"`
	NetCoreRmemMax                 pulumi.IntPtrInput    `pulumi:"netCoreRmemMax"`
	NetCoreSomaxconn               pulumi.IntPtrInput    `pulumi:"netCoreSomaxconn"`
	NetCoreWmemDefault             pulumi.IntPtrInput    `pulumi:"netCoreWmemDefault"`
	NetCoreWmemMax                 pulumi.IntPtrInput    `pulumi:"netCoreWmemMax"`
	NetIpv4IpLocalPortRange        pulumi.StringPtrInput `pulumi:"netIpv4IpLocalPortRange"`
	NetIpv4NeighDefaultGcThresh1   pulumi.IntPtrInput    `pulumi:"netIpv4NeighDefaultGcThresh1"`
	NetIpv4NeighDefaultGcThresh2   pulumi.IntPtrInput    `pulumi:"netIpv4NeighDefaultGcThresh2"`
	NetIpv4NeighDefaultGcThresh3   pulumi.IntPtrInput    `pulumi:"netIpv4NeighDefaultGcThresh3"`
	NetIpv4TcpFinTimeout           pulumi.IntPtrInput    `pulumi:"netIpv4TcpFinTimeout"`
	NetIpv4TcpKeepaliveProbes      pulumi.IntPtrInput    `pulumi:"netIpv4TcpKeepaliveProbes"`
	NetIpv4TcpKeepaliveTime        pulumi.IntPtrInput    `pulumi:"netIpv4TcpKeepaliveTime"`
	NetIpv4TcpMaxSynBacklog        pulumi.IntPtrInput    `pulumi:"netIpv4TcpMaxSynBacklog"`
	NetIpv4TcpMaxTwBuckets         pulumi.IntPtrInput    `pulumi:"netIpv4TcpMaxTwBuckets"`
	NetIpv4TcpTwReuse              pulumi.BoolPtrInput   `pulumi:"netIpv4TcpTwReuse"`
	NetIpv4TcpkeepaliveIntvl       pulumi.IntPtrInput    `pulumi:"netIpv4TcpkeepaliveIntvl"`
	NetNetfilterNfConntrackBuckets pulumi.IntPtrInput    `pulumi:"netNetfilterNfConntrackBuckets"`
	NetNetfilterNfConntrackMax     pulumi.IntPtrInput    `pulumi:"netNetfilterNfConntrackMax"`
	VmMaxMapCount                  pulumi.IntPtrInput    `pulumi:"vmMaxMapCount"`
	VmSwappiness                   pulumi.IntPtrInput    `pulumi:"vmSwappiness"`
	VmVfsCachePressure             pulumi.IntPtrInput    `pulumi:"vmVfsCachePressure"`
}

func (SysctlConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SysctlConfigResponse)(nil)).Elem()
}

func (i SysctlConfigResponseArgs) ToSysctlConfigResponseOutput() SysctlConfigResponseOutput {
	return i.ToSysctlConfigResponseOutputWithContext(context.Background())
}

func (i SysctlConfigResponseArgs) ToSysctlConfigResponseOutputWithContext(ctx context.Context) SysctlConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SysctlConfigResponseOutput)
}

func (i SysctlConfigResponseArgs) ToSysctlConfigResponsePtrOutput() SysctlConfigResponsePtrOutput {
	return i.ToSysctlConfigResponsePtrOutputWithContext(context.Background())
}

func (i SysctlConfigResponseArgs) ToSysctlConfigResponsePtrOutputWithContext(ctx context.Context) SysctlConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SysctlConfigResponseOutput).ToSysctlConfigResponsePtrOutputWithContext(ctx)
}









type SysctlConfigResponsePtrInput interface {
	pulumi.Input

	ToSysctlConfigResponsePtrOutput() SysctlConfigResponsePtrOutput
	ToSysctlConfigResponsePtrOutputWithContext(context.Context) SysctlConfigResponsePtrOutput
}

type sysctlConfigResponsePtrType SysctlConfigResponseArgs

func SysctlConfigResponsePtr(v *SysctlConfigResponseArgs) SysctlConfigResponsePtrInput {
	return (*sysctlConfigResponsePtrType)(v)
}

func (*sysctlConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SysctlConfigResponse)(nil)).Elem()
}

func (i *sysctlConfigResponsePtrType) ToSysctlConfigResponsePtrOutput() SysctlConfigResponsePtrOutput {
	return i.ToSysctlConfigResponsePtrOutputWithContext(context.Background())
}

func (i *sysctlConfigResponsePtrType) ToSysctlConfigResponsePtrOutputWithContext(ctx context.Context) SysctlConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SysctlConfigResponsePtrOutput)
}

type SysctlConfigResponseOutput struct{ *pulumi.OutputState }

func (SysctlConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SysctlConfigResponse)(nil)).Elem()
}

func (o SysctlConfigResponseOutput) ToSysctlConfigResponseOutput() SysctlConfigResponseOutput {
	return o
}

func (o SysctlConfigResponseOutput) ToSysctlConfigResponseOutputWithContext(ctx context.Context) SysctlConfigResponseOutput {
	return o
}

func (o SysctlConfigResponseOutput) ToSysctlConfigResponsePtrOutput() SysctlConfigResponsePtrOutput {
	return o.ToSysctlConfigResponsePtrOutputWithContext(context.Background())
}

func (o SysctlConfigResponseOutput) ToSysctlConfigResponsePtrOutputWithContext(ctx context.Context) SysctlConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SysctlConfigResponse) *SysctlConfigResponse {
		return &v
	}).(SysctlConfigResponsePtrOutput)
}

func (o SysctlConfigResponseOutput) FsAioMaxNr() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.FsAioMaxNr }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) FsFileMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.FsFileMax }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) FsInotifyMaxUserWatches() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.FsInotifyMaxUserWatches }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) FsNrOpen() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.FsNrOpen }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) KernelThreadsMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.KernelThreadsMax }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) NetCoreNetdevMaxBacklog() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.NetCoreNetdevMaxBacklog }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) NetCoreOptmemMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.NetCoreOptmemMax }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) NetCoreRmemDefault() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.NetCoreRmemDefault }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) NetCoreRmemMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.NetCoreRmemMax }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) NetCoreSomaxconn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.NetCoreSomaxconn }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) NetCoreWmemDefault() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.NetCoreWmemDefault }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) NetCoreWmemMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.NetCoreWmemMax }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) NetIpv4IpLocalPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *string { return v.NetIpv4IpLocalPortRange }).(pulumi.StringPtrOutput)
}

func (o SysctlConfigResponseOutput) NetIpv4NeighDefaultGcThresh1() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.NetIpv4NeighDefaultGcThresh1 }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) NetIpv4NeighDefaultGcThresh2() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.NetIpv4NeighDefaultGcThresh2 }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) NetIpv4NeighDefaultGcThresh3() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.NetIpv4NeighDefaultGcThresh3 }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) NetIpv4TcpFinTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.NetIpv4TcpFinTimeout }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) NetIpv4TcpKeepaliveProbes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.NetIpv4TcpKeepaliveProbes }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) NetIpv4TcpKeepaliveTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.NetIpv4TcpKeepaliveTime }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) NetIpv4TcpMaxSynBacklog() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.NetIpv4TcpMaxSynBacklog }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) NetIpv4TcpMaxTwBuckets() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.NetIpv4TcpMaxTwBuckets }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) NetIpv4TcpTwReuse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *bool { return v.NetIpv4TcpTwReuse }).(pulumi.BoolPtrOutput)
}

func (o SysctlConfigResponseOutput) NetIpv4TcpkeepaliveIntvl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.NetIpv4TcpkeepaliveIntvl }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) NetNetfilterNfConntrackBuckets() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.NetNetfilterNfConntrackBuckets }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) NetNetfilterNfConntrackMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.NetNetfilterNfConntrackMax }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) VmMaxMapCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.VmMaxMapCount }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) VmSwappiness() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.VmSwappiness }).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponseOutput) VmVfsCachePressure() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SysctlConfigResponse) *int { return v.VmVfsCachePressure }).(pulumi.IntPtrOutput)
}

type SysctlConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (SysctlConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SysctlConfigResponse)(nil)).Elem()
}

func (o SysctlConfigResponsePtrOutput) ToSysctlConfigResponsePtrOutput() SysctlConfigResponsePtrOutput {
	return o
}

func (o SysctlConfigResponsePtrOutput) ToSysctlConfigResponsePtrOutputWithContext(ctx context.Context) SysctlConfigResponsePtrOutput {
	return o
}

func (o SysctlConfigResponsePtrOutput) Elem() SysctlConfigResponseOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) SysctlConfigResponse {
		if v != nil {
			return *v
		}
		var ret SysctlConfigResponse
		return ret
	}).(SysctlConfigResponseOutput)
}

func (o SysctlConfigResponsePtrOutput) FsAioMaxNr() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.FsAioMaxNr
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) FsFileMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.FsFileMax
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) FsInotifyMaxUserWatches() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.FsInotifyMaxUserWatches
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) FsNrOpen() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.FsNrOpen
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) KernelThreadsMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.KernelThreadsMax
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) NetCoreNetdevMaxBacklog() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NetCoreNetdevMaxBacklog
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) NetCoreOptmemMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NetCoreOptmemMax
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) NetCoreRmemDefault() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NetCoreRmemDefault
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) NetCoreRmemMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NetCoreRmemMax
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) NetCoreSomaxconn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NetCoreSomaxconn
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) NetCoreWmemDefault() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NetCoreWmemDefault
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) NetCoreWmemMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NetCoreWmemMax
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) NetIpv4IpLocalPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.NetIpv4IpLocalPortRange
	}).(pulumi.StringPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) NetIpv4NeighDefaultGcThresh1() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NetIpv4NeighDefaultGcThresh1
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) NetIpv4NeighDefaultGcThresh2() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NetIpv4NeighDefaultGcThresh2
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) NetIpv4NeighDefaultGcThresh3() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NetIpv4NeighDefaultGcThresh3
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) NetIpv4TcpFinTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NetIpv4TcpFinTimeout
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) NetIpv4TcpKeepaliveProbes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NetIpv4TcpKeepaliveProbes
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) NetIpv4TcpKeepaliveTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NetIpv4TcpKeepaliveTime
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) NetIpv4TcpMaxSynBacklog() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NetIpv4TcpMaxSynBacklog
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) NetIpv4TcpMaxTwBuckets() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NetIpv4TcpMaxTwBuckets
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) NetIpv4TcpTwReuse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.NetIpv4TcpTwReuse
	}).(pulumi.BoolPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) NetIpv4TcpkeepaliveIntvl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NetIpv4TcpkeepaliveIntvl
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) NetNetfilterNfConntrackBuckets() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NetNetfilterNfConntrackBuckets
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) NetNetfilterNfConntrackMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NetNetfilterNfConntrackMax
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) VmMaxMapCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.VmMaxMapCount
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) VmSwappiness() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.VmSwappiness
	}).(pulumi.IntPtrOutput)
}

func (o SysctlConfigResponsePtrOutput) VmVfsCachePressure() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SysctlConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.VmVfsCachePressure
	}).(pulumi.IntPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
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
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

type TimeInWeek struct {
	Day       *string `pulumi:"day"`
	HourSlots []int   `pulumi:"hourSlots"`
}





type TimeInWeekInput interface {
	pulumi.Input

	ToTimeInWeekOutput() TimeInWeekOutput
	ToTimeInWeekOutputWithContext(context.Context) TimeInWeekOutput
}

type TimeInWeekArgs struct {
	Day       pulumi.StringPtrInput `pulumi:"day"`
	HourSlots pulumi.IntArrayInput  `pulumi:"hourSlots"`
}

func (TimeInWeekArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeInWeek)(nil)).Elem()
}

func (i TimeInWeekArgs) ToTimeInWeekOutput() TimeInWeekOutput {
	return i.ToTimeInWeekOutputWithContext(context.Background())
}

func (i TimeInWeekArgs) ToTimeInWeekOutputWithContext(ctx context.Context) TimeInWeekOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeInWeekOutput)
}





type TimeInWeekArrayInput interface {
	pulumi.Input

	ToTimeInWeekArrayOutput() TimeInWeekArrayOutput
	ToTimeInWeekArrayOutputWithContext(context.Context) TimeInWeekArrayOutput
}

type TimeInWeekArray []TimeInWeekInput

func (TimeInWeekArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeInWeek)(nil)).Elem()
}

func (i TimeInWeekArray) ToTimeInWeekArrayOutput() TimeInWeekArrayOutput {
	return i.ToTimeInWeekArrayOutputWithContext(context.Background())
}

func (i TimeInWeekArray) ToTimeInWeekArrayOutputWithContext(ctx context.Context) TimeInWeekArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeInWeekArrayOutput)
}

type TimeInWeekOutput struct{ *pulumi.OutputState }

func (TimeInWeekOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeInWeek)(nil)).Elem()
}

func (o TimeInWeekOutput) ToTimeInWeekOutput() TimeInWeekOutput {
	return o
}

func (o TimeInWeekOutput) ToTimeInWeekOutputWithContext(ctx context.Context) TimeInWeekOutput {
	return o
}

func (o TimeInWeekOutput) Day() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TimeInWeek) *string { return v.Day }).(pulumi.StringPtrOutput)
}

func (o TimeInWeekOutput) HourSlots() pulumi.IntArrayOutput {
	return o.ApplyT(func(v TimeInWeek) []int { return v.HourSlots }).(pulumi.IntArrayOutput)
}

type TimeInWeekArrayOutput struct{ *pulumi.OutputState }

func (TimeInWeekArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeInWeek)(nil)).Elem()
}

func (o TimeInWeekArrayOutput) ToTimeInWeekArrayOutput() TimeInWeekArrayOutput {
	return o
}

func (o TimeInWeekArrayOutput) ToTimeInWeekArrayOutputWithContext(ctx context.Context) TimeInWeekArrayOutput {
	return o
}

func (o TimeInWeekArrayOutput) Index(i pulumi.IntInput) TimeInWeekOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimeInWeek {
		return vs[0].([]TimeInWeek)[vs[1].(int)]
	}).(TimeInWeekOutput)
}

type TimeInWeekResponse struct {
	Day       *string `pulumi:"day"`
	HourSlots []int   `pulumi:"hourSlots"`
}





type TimeInWeekResponseInput interface {
	pulumi.Input

	ToTimeInWeekResponseOutput() TimeInWeekResponseOutput
	ToTimeInWeekResponseOutputWithContext(context.Context) TimeInWeekResponseOutput
}

type TimeInWeekResponseArgs struct {
	Day       pulumi.StringPtrInput `pulumi:"day"`
	HourSlots pulumi.IntArrayInput  `pulumi:"hourSlots"`
}

func (TimeInWeekResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeInWeekResponse)(nil)).Elem()
}

func (i TimeInWeekResponseArgs) ToTimeInWeekResponseOutput() TimeInWeekResponseOutput {
	return i.ToTimeInWeekResponseOutputWithContext(context.Background())
}

func (i TimeInWeekResponseArgs) ToTimeInWeekResponseOutputWithContext(ctx context.Context) TimeInWeekResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeInWeekResponseOutput)
}





type TimeInWeekResponseArrayInput interface {
	pulumi.Input

	ToTimeInWeekResponseArrayOutput() TimeInWeekResponseArrayOutput
	ToTimeInWeekResponseArrayOutputWithContext(context.Context) TimeInWeekResponseArrayOutput
}

type TimeInWeekResponseArray []TimeInWeekResponseInput

func (TimeInWeekResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeInWeekResponse)(nil)).Elem()
}

func (i TimeInWeekResponseArray) ToTimeInWeekResponseArrayOutput() TimeInWeekResponseArrayOutput {
	return i.ToTimeInWeekResponseArrayOutputWithContext(context.Background())
}

func (i TimeInWeekResponseArray) ToTimeInWeekResponseArrayOutputWithContext(ctx context.Context) TimeInWeekResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeInWeekResponseArrayOutput)
}

type TimeInWeekResponseOutput struct{ *pulumi.OutputState }

func (TimeInWeekResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeInWeekResponse)(nil)).Elem()
}

func (o TimeInWeekResponseOutput) ToTimeInWeekResponseOutput() TimeInWeekResponseOutput {
	return o
}

func (o TimeInWeekResponseOutput) ToTimeInWeekResponseOutputWithContext(ctx context.Context) TimeInWeekResponseOutput {
	return o
}

func (o TimeInWeekResponseOutput) Day() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TimeInWeekResponse) *string { return v.Day }).(pulumi.StringPtrOutput)
}

func (o TimeInWeekResponseOutput) HourSlots() pulumi.IntArrayOutput {
	return o.ApplyT(func(v TimeInWeekResponse) []int { return v.HourSlots }).(pulumi.IntArrayOutput)
}

type TimeInWeekResponseArrayOutput struct{ *pulumi.OutputState }

func (TimeInWeekResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeInWeekResponse)(nil)).Elem()
}

func (o TimeInWeekResponseArrayOutput) ToTimeInWeekResponseArrayOutput() TimeInWeekResponseArrayOutput {
	return o
}

func (o TimeInWeekResponseArrayOutput) ToTimeInWeekResponseArrayOutputWithContext(ctx context.Context) TimeInWeekResponseArrayOutput {
	return o
}

func (o TimeInWeekResponseArrayOutput) Index(i pulumi.IntInput) TimeInWeekResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimeInWeekResponse {
		return vs[0].([]TimeInWeekResponse)[vs[1].(int)]
	}).(TimeInWeekResponseOutput)
}

type TimeSpan struct {
	End   *string `pulumi:"end"`
	Start *string `pulumi:"start"`
}





type TimeSpanInput interface {
	pulumi.Input

	ToTimeSpanOutput() TimeSpanOutput
	ToTimeSpanOutputWithContext(context.Context) TimeSpanOutput
}

type TimeSpanArgs struct {
	End   pulumi.StringPtrInput `pulumi:"end"`
	Start pulumi.StringPtrInput `pulumi:"start"`
}

func (TimeSpanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSpan)(nil)).Elem()
}

func (i TimeSpanArgs) ToTimeSpanOutput() TimeSpanOutput {
	return i.ToTimeSpanOutputWithContext(context.Background())
}

func (i TimeSpanArgs) ToTimeSpanOutputWithContext(ctx context.Context) TimeSpanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSpanOutput)
}





type TimeSpanArrayInput interface {
	pulumi.Input

	ToTimeSpanArrayOutput() TimeSpanArrayOutput
	ToTimeSpanArrayOutputWithContext(context.Context) TimeSpanArrayOutput
}

type TimeSpanArray []TimeSpanInput

func (TimeSpanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeSpan)(nil)).Elem()
}

func (i TimeSpanArray) ToTimeSpanArrayOutput() TimeSpanArrayOutput {
	return i.ToTimeSpanArrayOutputWithContext(context.Background())
}

func (i TimeSpanArray) ToTimeSpanArrayOutputWithContext(ctx context.Context) TimeSpanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSpanArrayOutput)
}

type TimeSpanOutput struct{ *pulumi.OutputState }

func (TimeSpanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSpan)(nil)).Elem()
}

func (o TimeSpanOutput) ToTimeSpanOutput() TimeSpanOutput {
	return o
}

func (o TimeSpanOutput) ToTimeSpanOutputWithContext(ctx context.Context) TimeSpanOutput {
	return o
}

func (o TimeSpanOutput) End() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TimeSpan) *string { return v.End }).(pulumi.StringPtrOutput)
}

func (o TimeSpanOutput) Start() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TimeSpan) *string { return v.Start }).(pulumi.StringPtrOutput)
}

type TimeSpanArrayOutput struct{ *pulumi.OutputState }

func (TimeSpanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeSpan)(nil)).Elem()
}

func (o TimeSpanArrayOutput) ToTimeSpanArrayOutput() TimeSpanArrayOutput {
	return o
}

func (o TimeSpanArrayOutput) ToTimeSpanArrayOutputWithContext(ctx context.Context) TimeSpanArrayOutput {
	return o
}

func (o TimeSpanArrayOutput) Index(i pulumi.IntInput) TimeSpanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimeSpan {
		return vs[0].([]TimeSpan)[vs[1].(int)]
	}).(TimeSpanOutput)
}

type TimeSpanResponse struct {
	End   *string `pulumi:"end"`
	Start *string `pulumi:"start"`
}





type TimeSpanResponseInput interface {
	pulumi.Input

	ToTimeSpanResponseOutput() TimeSpanResponseOutput
	ToTimeSpanResponseOutputWithContext(context.Context) TimeSpanResponseOutput
}

type TimeSpanResponseArgs struct {
	End   pulumi.StringPtrInput `pulumi:"end"`
	Start pulumi.StringPtrInput `pulumi:"start"`
}

func (TimeSpanResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSpanResponse)(nil)).Elem()
}

func (i TimeSpanResponseArgs) ToTimeSpanResponseOutput() TimeSpanResponseOutput {
	return i.ToTimeSpanResponseOutputWithContext(context.Background())
}

func (i TimeSpanResponseArgs) ToTimeSpanResponseOutputWithContext(ctx context.Context) TimeSpanResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSpanResponseOutput)
}





type TimeSpanResponseArrayInput interface {
	pulumi.Input

	ToTimeSpanResponseArrayOutput() TimeSpanResponseArrayOutput
	ToTimeSpanResponseArrayOutputWithContext(context.Context) TimeSpanResponseArrayOutput
}

type TimeSpanResponseArray []TimeSpanResponseInput

func (TimeSpanResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeSpanResponse)(nil)).Elem()
}

func (i TimeSpanResponseArray) ToTimeSpanResponseArrayOutput() TimeSpanResponseArrayOutput {
	return i.ToTimeSpanResponseArrayOutputWithContext(context.Background())
}

func (i TimeSpanResponseArray) ToTimeSpanResponseArrayOutputWithContext(ctx context.Context) TimeSpanResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSpanResponseArrayOutput)
}

type TimeSpanResponseOutput struct{ *pulumi.OutputState }

func (TimeSpanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSpanResponse)(nil)).Elem()
}

func (o TimeSpanResponseOutput) ToTimeSpanResponseOutput() TimeSpanResponseOutput {
	return o
}

func (o TimeSpanResponseOutput) ToTimeSpanResponseOutputWithContext(ctx context.Context) TimeSpanResponseOutput {
	return o
}

func (o TimeSpanResponseOutput) End() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TimeSpanResponse) *string { return v.End }).(pulumi.StringPtrOutput)
}

func (o TimeSpanResponseOutput) Start() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TimeSpanResponse) *string { return v.Start }).(pulumi.StringPtrOutput)
}

type TimeSpanResponseArrayOutput struct{ *pulumi.OutputState }

func (TimeSpanResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeSpanResponse)(nil)).Elem()
}

func (o TimeSpanResponseArrayOutput) ToTimeSpanResponseArrayOutput() TimeSpanResponseArrayOutput {
	return o
}

func (o TimeSpanResponseArrayOutput) ToTimeSpanResponseArrayOutputWithContext(ctx context.Context) TimeSpanResponseArrayOutput {
	return o
}

func (o TimeSpanResponseArrayOutput) Index(i pulumi.IntInput) TimeSpanResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimeSpanResponse {
		return vs[0].([]TimeSpanResponse)[vs[1].(int)]
	}).(TimeSpanResponseOutput)
}

type UserAssignedIdentity struct {
	ClientId   *string `pulumi:"clientId"`
	ObjectId   *string `pulumi:"objectId"`
	ResourceId *string `pulumi:"resourceId"`
}





type UserAssignedIdentityInput interface {
	pulumi.Input

	ToUserAssignedIdentityOutput() UserAssignedIdentityOutput
	ToUserAssignedIdentityOutputWithContext(context.Context) UserAssignedIdentityOutput
}

type UserAssignedIdentityArgs struct {
	ClientId   pulumi.StringPtrInput `pulumi:"clientId"`
	ObjectId   pulumi.StringPtrInput `pulumi:"objectId"`
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (UserAssignedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentity)(nil)).Elem()
}

func (i UserAssignedIdentityArgs) ToUserAssignedIdentityOutput() UserAssignedIdentityOutput {
	return i.ToUserAssignedIdentityOutputWithContext(context.Background())
}

func (i UserAssignedIdentityArgs) ToUserAssignedIdentityOutputWithContext(ctx context.Context) UserAssignedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityOutput)
}

type UserAssignedIdentityOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentity)(nil)).Elem()
}

func (o UserAssignedIdentityOutput) ToUserAssignedIdentityOutput() UserAssignedIdentityOutput {
	return o
}

func (o UserAssignedIdentityOutput) ToUserAssignedIdentityOutputWithContext(ctx context.Context) UserAssignedIdentityOutput {
	return o
}

func (o UserAssignedIdentityOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentity) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o UserAssignedIdentityOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentity) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o UserAssignedIdentityOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentity) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityResponse struct {
	ClientId   *string `pulumi:"clientId"`
	ObjectId   *string `pulumi:"objectId"`
	ResourceId *string `pulumi:"resourceId"`
}





type UserAssignedIdentityResponseInput interface {
	pulumi.Input

	ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput
	ToUserAssignedIdentityResponseOutputWithContext(context.Context) UserAssignedIdentityResponseOutput
}

type UserAssignedIdentityResponseArgs struct {
	ClientId   pulumi.StringPtrInput `pulumi:"clientId"`
	ObjectId   pulumi.StringPtrInput `pulumi:"objectId"`
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (UserAssignedIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (i UserAssignedIdentityResponseArgs) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return i.ToUserAssignedIdentityResponseOutputWithContext(context.Background())
}

func (i UserAssignedIdentityResponseArgs) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityResponseOutput)
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

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o UserAssignedIdentityResponseOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o UserAssignedIdentityResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AgentPoolUpgradeSettingsOutput{})
	pulumi.RegisterOutputType(AgentPoolUpgradeSettingsPtrOutput{})
	pulumi.RegisterOutputType(AgentPoolUpgradeSettingsResponseOutput{})
	pulumi.RegisterOutputType(AgentPoolUpgradeSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(CloudErrorBodyResponseOutput{})
	pulumi.RegisterOutputType(CloudErrorBodyResponsePtrOutput{})
	pulumi.RegisterOutputType(CloudErrorBodyResponseArrayOutput{})
	pulumi.RegisterOutputType(CloudErrorResponseOutput{})
	pulumi.RegisterOutputType(CloudErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceLinuxProfileOutput{})
	pulumi.RegisterOutputType(ContainerServiceLinuxProfilePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceLinuxProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceLinuxProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceNetworkProfileOutput{})
	pulumi.RegisterOutputType(ContainerServiceNetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceNetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceNetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshConfigurationOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshPublicKeyOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshPublicKeyArrayOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshPublicKeyResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshPublicKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(CredentialResultResponseOutput{})
	pulumi.RegisterOutputType(CredentialResultResponseArrayOutput{})
	pulumi.RegisterOutputType(ExtendedLocationOutput{})
	pulumi.RegisterOutputType(ExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponseOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponsePtrOutput{})
	pulumi.RegisterOutputType(KubeletConfigOutput{})
	pulumi.RegisterOutputType(KubeletConfigPtrOutput{})
	pulumi.RegisterOutputType(KubeletConfigResponseOutput{})
	pulumi.RegisterOutputType(KubeletConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(LinuxOSConfigOutput{})
	pulumi.RegisterOutputType(LinuxOSConfigPtrOutput{})
	pulumi.RegisterOutputType(LinuxOSConfigResponseOutput{})
	pulumi.RegisterOutputType(LinuxOSConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterAADProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterAADProfilePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterAADProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterAADProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterAPIServerAccessProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterAPIServerAccessProfilePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterAPIServerAccessProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterAPIServerAccessProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterAddonProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterAddonProfileMapOutput{})
	pulumi.RegisterOutputType(ManagedClusterAddonProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterAddonProfileResponseMapOutput{})
	pulumi.RegisterOutputType(ManagedClusterAddonProfileResponseIdentityOutput{})
	pulumi.RegisterOutputType(ManagedClusterAgentPoolProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterAgentPoolProfileArrayOutput{})
	pulumi.RegisterOutputType(ManagedClusterAgentPoolProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterAgentPoolProfileResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedClusterAutoUpgradeProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterAutoUpgradeProfilePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterAutoUpgradeProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterAutoUpgradeProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterHTTPProxyConfigOutput{})
	pulumi.RegisterOutputType(ManagedClusterHTTPProxyConfigPtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterHTTPProxyConfigResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterHTTPProxyConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterIdentityOutput{})
	pulumi.RegisterOutputType(ManagedClusterIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterIdentityResponseUserAssignedIdentitiesOutput{})
	pulumi.RegisterOutputType(ManagedClusterIdentityResponseUserAssignedIdentitiesMapOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfilePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileOutboundIPsOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterPodIdentityOutput{})
	pulumi.RegisterOutputType(ManagedClusterPodIdentityArrayOutput{})
	pulumi.RegisterOutputType(ManagedClusterPodIdentityExceptionOutput{})
	pulumi.RegisterOutputType(ManagedClusterPodIdentityExceptionArrayOutput{})
	pulumi.RegisterOutputType(ManagedClusterPodIdentityExceptionResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterPodIdentityExceptionResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedClusterPodIdentityProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterPodIdentityProfilePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterPodIdentityProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterPodIdentityProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterPodIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterPodIdentityResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedClusterPodIdentityResponseProvisioningInfoOutput{})
	pulumi.RegisterOutputType(ManagedClusterPropertiesAutoScalerProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterPropertiesAutoScalerProfilePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterPropertiesIdentityProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterPropertiesIdentityProfileMapOutput{})
	pulumi.RegisterOutputType(ManagedClusterPropertiesResponseAutoScalerProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterPropertiesResponseIdentityProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterPropertiesResponseIdentityProfileMapOutput{})
	pulumi.RegisterOutputType(ManagedClusterSKUOutput{})
	pulumi.RegisterOutputType(ManagedClusterSKUPtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterSKUResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterSKUResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterServicePrincipalProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterServicePrincipalProfilePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterServicePrincipalProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterServicePrincipalProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterWindowsProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterWindowsProfilePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterWindowsProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterWindowsProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(PowerStateResponseOutput{})
	pulumi.RegisterOutputType(PowerStateResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkResourceOutput{})
	pulumi.RegisterOutputType(PrivateLinkResourceArrayOutput{})
	pulumi.RegisterOutputType(PrivateLinkResourceResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceReferenceOutput{})
	pulumi.RegisterOutputType(ResourceReferenceArrayOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponseOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(SysctlConfigOutput{})
	pulumi.RegisterOutputType(SysctlConfigPtrOutput{})
	pulumi.RegisterOutputType(SysctlConfigResponseOutput{})
	pulumi.RegisterOutputType(SysctlConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(TimeInWeekOutput{})
	pulumi.RegisterOutputType(TimeInWeekArrayOutput{})
	pulumi.RegisterOutputType(TimeInWeekResponseOutput{})
	pulumi.RegisterOutputType(TimeInWeekResponseArrayOutput{})
	pulumi.RegisterOutputType(TimeSpanOutput{})
	pulumi.RegisterOutputType(TimeSpanArrayOutput{})
	pulumi.RegisterOutputType(TimeSpanResponseOutput{})
	pulumi.RegisterOutputType(TimeSpanResponseArrayOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
}




package v20180901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AgentProperties struct {
	Cpu *int `pulumi:"cpu"`
}





type AgentPropertiesInput interface {
	pulumi.Input

	ToAgentPropertiesOutput() AgentPropertiesOutput
	ToAgentPropertiesOutputWithContext(context.Context) AgentPropertiesOutput
}

type AgentPropertiesArgs struct {
	Cpu pulumi.IntPtrInput `pulumi:"cpu"`
}

func (AgentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentProperties)(nil)).Elem()
}

func (i AgentPropertiesArgs) ToAgentPropertiesOutput() AgentPropertiesOutput {
	return i.ToAgentPropertiesOutputWithContext(context.Background())
}

func (i AgentPropertiesArgs) ToAgentPropertiesOutputWithContext(ctx context.Context) AgentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPropertiesOutput)
}

func (i AgentPropertiesArgs) ToAgentPropertiesPtrOutput() AgentPropertiesPtrOutput {
	return i.ToAgentPropertiesPtrOutputWithContext(context.Background())
}

func (i AgentPropertiesArgs) ToAgentPropertiesPtrOutputWithContext(ctx context.Context) AgentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPropertiesOutput).ToAgentPropertiesPtrOutputWithContext(ctx)
}









type AgentPropertiesPtrInput interface {
	pulumi.Input

	ToAgentPropertiesPtrOutput() AgentPropertiesPtrOutput
	ToAgentPropertiesPtrOutputWithContext(context.Context) AgentPropertiesPtrOutput
}

type agentPropertiesPtrType AgentPropertiesArgs

func AgentPropertiesPtr(v *AgentPropertiesArgs) AgentPropertiesPtrInput {
	return (*agentPropertiesPtrType)(v)
}

func (*agentPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentProperties)(nil)).Elem()
}

func (i *agentPropertiesPtrType) ToAgentPropertiesPtrOutput() AgentPropertiesPtrOutput {
	return i.ToAgentPropertiesPtrOutputWithContext(context.Background())
}

func (i *agentPropertiesPtrType) ToAgentPropertiesPtrOutputWithContext(ctx context.Context) AgentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPropertiesPtrOutput)
}

type AgentPropertiesOutput struct{ *pulumi.OutputState }

func (AgentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentProperties)(nil)).Elem()
}

func (o AgentPropertiesOutput) ToAgentPropertiesOutput() AgentPropertiesOutput {
	return o
}

func (o AgentPropertiesOutput) ToAgentPropertiesOutputWithContext(ctx context.Context) AgentPropertiesOutput {
	return o
}

func (o AgentPropertiesOutput) ToAgentPropertiesPtrOutput() AgentPropertiesPtrOutput {
	return o.ToAgentPropertiesPtrOutputWithContext(context.Background())
}

func (o AgentPropertiesOutput) ToAgentPropertiesPtrOutputWithContext(ctx context.Context) AgentPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AgentProperties) *AgentProperties {
		return &v
	}).(AgentPropertiesPtrOutput)
}

func (o AgentPropertiesOutput) Cpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AgentProperties) *int { return v.Cpu }).(pulumi.IntPtrOutput)
}

type AgentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (AgentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentProperties)(nil)).Elem()
}

func (o AgentPropertiesPtrOutput) ToAgentPropertiesPtrOutput() AgentPropertiesPtrOutput {
	return o
}

func (o AgentPropertiesPtrOutput) ToAgentPropertiesPtrOutputWithContext(ctx context.Context) AgentPropertiesPtrOutput {
	return o
}

func (o AgentPropertiesPtrOutput) Elem() AgentPropertiesOutput {
	return o.ApplyT(func(v *AgentProperties) AgentProperties {
		if v != nil {
			return *v
		}
		var ret AgentProperties
		return ret
	}).(AgentPropertiesOutput)
}

func (o AgentPropertiesPtrOutput) Cpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentProperties) *int {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.IntPtrOutput)
}

type AgentPropertiesResponse struct {
	Cpu *int `pulumi:"cpu"`
}





type AgentPropertiesResponseInput interface {
	pulumi.Input

	ToAgentPropertiesResponseOutput() AgentPropertiesResponseOutput
	ToAgentPropertiesResponseOutputWithContext(context.Context) AgentPropertiesResponseOutput
}

type AgentPropertiesResponseArgs struct {
	Cpu pulumi.IntPtrInput `pulumi:"cpu"`
}

func (AgentPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPropertiesResponse)(nil)).Elem()
}

func (i AgentPropertiesResponseArgs) ToAgentPropertiesResponseOutput() AgentPropertiesResponseOutput {
	return i.ToAgentPropertiesResponseOutputWithContext(context.Background())
}

func (i AgentPropertiesResponseArgs) ToAgentPropertiesResponseOutputWithContext(ctx context.Context) AgentPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPropertiesResponseOutput)
}

func (i AgentPropertiesResponseArgs) ToAgentPropertiesResponsePtrOutput() AgentPropertiesResponsePtrOutput {
	return i.ToAgentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i AgentPropertiesResponseArgs) ToAgentPropertiesResponsePtrOutputWithContext(ctx context.Context) AgentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPropertiesResponseOutput).ToAgentPropertiesResponsePtrOutputWithContext(ctx)
}









type AgentPropertiesResponsePtrInput interface {
	pulumi.Input

	ToAgentPropertiesResponsePtrOutput() AgentPropertiesResponsePtrOutput
	ToAgentPropertiesResponsePtrOutputWithContext(context.Context) AgentPropertiesResponsePtrOutput
}

type agentPropertiesResponsePtrType AgentPropertiesResponseArgs

func AgentPropertiesResponsePtr(v *AgentPropertiesResponseArgs) AgentPropertiesResponsePtrInput {
	return (*agentPropertiesResponsePtrType)(v)
}

func (*agentPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPropertiesResponse)(nil)).Elem()
}

func (i *agentPropertiesResponsePtrType) ToAgentPropertiesResponsePtrOutput() AgentPropertiesResponsePtrOutput {
	return i.ToAgentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *agentPropertiesResponsePtrType) ToAgentPropertiesResponsePtrOutputWithContext(ctx context.Context) AgentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPropertiesResponsePtrOutput)
}

type AgentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AgentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPropertiesResponse)(nil)).Elem()
}

func (o AgentPropertiesResponseOutput) ToAgentPropertiesResponseOutput() AgentPropertiesResponseOutput {
	return o
}

func (o AgentPropertiesResponseOutput) ToAgentPropertiesResponseOutputWithContext(ctx context.Context) AgentPropertiesResponseOutput {
	return o
}

func (o AgentPropertiesResponseOutput) ToAgentPropertiesResponsePtrOutput() AgentPropertiesResponsePtrOutput {
	return o.ToAgentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o AgentPropertiesResponseOutput) ToAgentPropertiesResponsePtrOutputWithContext(ctx context.Context) AgentPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AgentPropertiesResponse) *AgentPropertiesResponse {
		return &v
	}).(AgentPropertiesResponsePtrOutput)
}

func (o AgentPropertiesResponseOutput) Cpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AgentPropertiesResponse) *int { return v.Cpu }).(pulumi.IntPtrOutput)
}

type AgentPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AgentPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPropertiesResponse)(nil)).Elem()
}

func (o AgentPropertiesResponsePtrOutput) ToAgentPropertiesResponsePtrOutput() AgentPropertiesResponsePtrOutput {
	return o
}

func (o AgentPropertiesResponsePtrOutput) ToAgentPropertiesResponsePtrOutputWithContext(ctx context.Context) AgentPropertiesResponsePtrOutput {
	return o
}

func (o AgentPropertiesResponsePtrOutput) Elem() AgentPropertiesResponseOutput {
	return o.ApplyT(func(v *AgentPropertiesResponse) AgentPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AgentPropertiesResponse
		return ret
	}).(AgentPropertiesResponseOutput)
}

func (o AgentPropertiesResponsePtrOutput) Cpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.IntPtrOutput)
}

type ArgumentResponse struct {
	IsSecret *bool  `pulumi:"isSecret"`
	Name     string `pulumi:"name"`
	Value    string `pulumi:"value"`
}





type ArgumentResponseInput interface {
	pulumi.Input

	ToArgumentResponseOutput() ArgumentResponseOutput
	ToArgumentResponseOutputWithContext(context.Context) ArgumentResponseOutput
}

type ArgumentResponseArgs struct {
	IsSecret pulumi.BoolPtrInput `pulumi:"isSecret"`
	Name     pulumi.StringInput  `pulumi:"name"`
	Value    pulumi.StringInput  `pulumi:"value"`
}

func (ArgumentResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArgumentResponse)(nil)).Elem()
}

func (i ArgumentResponseArgs) ToArgumentResponseOutput() ArgumentResponseOutput {
	return i.ToArgumentResponseOutputWithContext(context.Background())
}

func (i ArgumentResponseArgs) ToArgumentResponseOutputWithContext(ctx context.Context) ArgumentResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArgumentResponseOutput)
}





type ArgumentResponseArrayInput interface {
	pulumi.Input

	ToArgumentResponseArrayOutput() ArgumentResponseArrayOutput
	ToArgumentResponseArrayOutputWithContext(context.Context) ArgumentResponseArrayOutput
}

type ArgumentResponseArray []ArgumentResponseInput

func (ArgumentResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArgumentResponse)(nil)).Elem()
}

func (i ArgumentResponseArray) ToArgumentResponseArrayOutput() ArgumentResponseArrayOutput {
	return i.ToArgumentResponseArrayOutputWithContext(context.Background())
}

func (i ArgumentResponseArray) ToArgumentResponseArrayOutputWithContext(ctx context.Context) ArgumentResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArgumentResponseArrayOutput)
}

type ArgumentResponseOutput struct{ *pulumi.OutputState }

func (ArgumentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArgumentResponse)(nil)).Elem()
}

func (o ArgumentResponseOutput) ToArgumentResponseOutput() ArgumentResponseOutput {
	return o
}

func (o ArgumentResponseOutput) ToArgumentResponseOutputWithContext(ctx context.Context) ArgumentResponseOutput {
	return o
}

func (o ArgumentResponseOutput) GetIsSecret() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ArgumentResponse) *bool { return v.IsSecret }).(pulumi.BoolPtrOutput)
}

func (o ArgumentResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ArgumentResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ArgumentResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ArgumentResponse) string { return v.Value }).(pulumi.StringOutput)
}

type ArgumentResponseArrayOutput struct{ *pulumi.OutputState }

func (ArgumentResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArgumentResponse)(nil)).Elem()
}

func (o ArgumentResponseArrayOutput) ToArgumentResponseArrayOutput() ArgumentResponseArrayOutput {
	return o
}

func (o ArgumentResponseArrayOutput) ToArgumentResponseArrayOutputWithContext(ctx context.Context) ArgumentResponseArrayOutput {
	return o
}

func (o ArgumentResponseArrayOutput) Index(i pulumi.IntInput) ArgumentResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ArgumentResponse {
		return vs[0].([]ArgumentResponse)[vs[1].(int)]
	}).(ArgumentResponseOutput)
}

type AuthInfo struct {
	ExpiresIn    *int    `pulumi:"expiresIn"`
	RefreshToken *string `pulumi:"refreshToken"`
	Scope        *string `pulumi:"scope"`
	Token        string  `pulumi:"token"`
	TokenType    string  `pulumi:"tokenType"`
}





type AuthInfoInput interface {
	pulumi.Input

	ToAuthInfoOutput() AuthInfoOutput
	ToAuthInfoOutputWithContext(context.Context) AuthInfoOutput
}

type AuthInfoArgs struct {
	ExpiresIn    pulumi.IntPtrInput    `pulumi:"expiresIn"`
	RefreshToken pulumi.StringPtrInput `pulumi:"refreshToken"`
	Scope        pulumi.StringPtrInput `pulumi:"scope"`
	Token        pulumi.StringInput    `pulumi:"token"`
	TokenType    pulumi.StringInput    `pulumi:"tokenType"`
}

func (AuthInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthInfo)(nil)).Elem()
}

func (i AuthInfoArgs) ToAuthInfoOutput() AuthInfoOutput {
	return i.ToAuthInfoOutputWithContext(context.Background())
}

func (i AuthInfoArgs) ToAuthInfoOutputWithContext(ctx context.Context) AuthInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthInfoOutput)
}

func (i AuthInfoArgs) ToAuthInfoPtrOutput() AuthInfoPtrOutput {
	return i.ToAuthInfoPtrOutputWithContext(context.Background())
}

func (i AuthInfoArgs) ToAuthInfoPtrOutputWithContext(ctx context.Context) AuthInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthInfoOutput).ToAuthInfoPtrOutputWithContext(ctx)
}









type AuthInfoPtrInput interface {
	pulumi.Input

	ToAuthInfoPtrOutput() AuthInfoPtrOutput
	ToAuthInfoPtrOutputWithContext(context.Context) AuthInfoPtrOutput
}

type authInfoPtrType AuthInfoArgs

func AuthInfoPtr(v *AuthInfoArgs) AuthInfoPtrInput {
	return (*authInfoPtrType)(v)
}

func (*authInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthInfo)(nil)).Elem()
}

func (i *authInfoPtrType) ToAuthInfoPtrOutput() AuthInfoPtrOutput {
	return i.ToAuthInfoPtrOutputWithContext(context.Background())
}

func (i *authInfoPtrType) ToAuthInfoPtrOutputWithContext(ctx context.Context) AuthInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthInfoPtrOutput)
}

type AuthInfoOutput struct{ *pulumi.OutputState }

func (AuthInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthInfo)(nil)).Elem()
}

func (o AuthInfoOutput) ToAuthInfoOutput() AuthInfoOutput {
	return o
}

func (o AuthInfoOutput) ToAuthInfoOutputWithContext(ctx context.Context) AuthInfoOutput {
	return o
}

func (o AuthInfoOutput) ToAuthInfoPtrOutput() AuthInfoPtrOutput {
	return o.ToAuthInfoPtrOutputWithContext(context.Background())
}

func (o AuthInfoOutput) ToAuthInfoPtrOutputWithContext(ctx context.Context) AuthInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthInfo) *AuthInfo {
		return &v
	}).(AuthInfoPtrOutput)
}

func (o AuthInfoOutput) ExpiresIn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AuthInfo) *int { return v.ExpiresIn }).(pulumi.IntPtrOutput)
}

func (o AuthInfoOutput) RefreshToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthInfo) *string { return v.RefreshToken }).(pulumi.StringPtrOutput)
}

func (o AuthInfoOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthInfo) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o AuthInfoOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v AuthInfo) string { return v.Token }).(pulumi.StringOutput)
}

func (o AuthInfoOutput) TokenType() pulumi.StringOutput {
	return o.ApplyT(func(v AuthInfo) string { return v.TokenType }).(pulumi.StringOutput)
}

type AuthInfoPtrOutput struct{ *pulumi.OutputState }

func (AuthInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthInfo)(nil)).Elem()
}

func (o AuthInfoPtrOutput) ToAuthInfoPtrOutput() AuthInfoPtrOutput {
	return o
}

func (o AuthInfoPtrOutput) ToAuthInfoPtrOutputWithContext(ctx context.Context) AuthInfoPtrOutput {
	return o
}

func (o AuthInfoPtrOutput) Elem() AuthInfoOutput {
	return o.ApplyT(func(v *AuthInfo) AuthInfo {
		if v != nil {
			return *v
		}
		var ret AuthInfo
		return ret
	}).(AuthInfoOutput)
}

func (o AuthInfoPtrOutput) ExpiresIn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AuthInfo) *int {
		if v == nil {
			return nil
		}
		return v.ExpiresIn
	}).(pulumi.IntPtrOutput)
}

func (o AuthInfoPtrOutput) RefreshToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthInfo) *string {
		if v == nil {
			return nil
		}
		return v.RefreshToken
	}).(pulumi.StringPtrOutput)
}

func (o AuthInfoPtrOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthInfo) *string {
		if v == nil {
			return nil
		}
		return v.Scope
	}).(pulumi.StringPtrOutput)
}

func (o AuthInfoPtrOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthInfo) *string {
		if v == nil {
			return nil
		}
		return &v.Token
	}).(pulumi.StringPtrOutput)
}

func (o AuthInfoPtrOutput) TokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthInfo) *string {
		if v == nil {
			return nil
		}
		return &v.TokenType
	}).(pulumi.StringPtrOutput)
}

type AuthInfoResponse struct {
	ExpiresIn    *int    `pulumi:"expiresIn"`
	RefreshToken *string `pulumi:"refreshToken"`
	Scope        *string `pulumi:"scope"`
	Token        string  `pulumi:"token"`
	TokenType    string  `pulumi:"tokenType"`
}





type AuthInfoResponseInput interface {
	pulumi.Input

	ToAuthInfoResponseOutput() AuthInfoResponseOutput
	ToAuthInfoResponseOutputWithContext(context.Context) AuthInfoResponseOutput
}

type AuthInfoResponseArgs struct {
	ExpiresIn    pulumi.IntPtrInput    `pulumi:"expiresIn"`
	RefreshToken pulumi.StringPtrInput `pulumi:"refreshToken"`
	Scope        pulumi.StringPtrInput `pulumi:"scope"`
	Token        pulumi.StringInput    `pulumi:"token"`
	TokenType    pulumi.StringInput    `pulumi:"tokenType"`
}

func (AuthInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthInfoResponse)(nil)).Elem()
}

func (i AuthInfoResponseArgs) ToAuthInfoResponseOutput() AuthInfoResponseOutput {
	return i.ToAuthInfoResponseOutputWithContext(context.Background())
}

func (i AuthInfoResponseArgs) ToAuthInfoResponseOutputWithContext(ctx context.Context) AuthInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthInfoResponseOutput)
}

func (i AuthInfoResponseArgs) ToAuthInfoResponsePtrOutput() AuthInfoResponsePtrOutput {
	return i.ToAuthInfoResponsePtrOutputWithContext(context.Background())
}

func (i AuthInfoResponseArgs) ToAuthInfoResponsePtrOutputWithContext(ctx context.Context) AuthInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthInfoResponseOutput).ToAuthInfoResponsePtrOutputWithContext(ctx)
}









type AuthInfoResponsePtrInput interface {
	pulumi.Input

	ToAuthInfoResponsePtrOutput() AuthInfoResponsePtrOutput
	ToAuthInfoResponsePtrOutputWithContext(context.Context) AuthInfoResponsePtrOutput
}

type authInfoResponsePtrType AuthInfoResponseArgs

func AuthInfoResponsePtr(v *AuthInfoResponseArgs) AuthInfoResponsePtrInput {
	return (*authInfoResponsePtrType)(v)
}

func (*authInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthInfoResponse)(nil)).Elem()
}

func (i *authInfoResponsePtrType) ToAuthInfoResponsePtrOutput() AuthInfoResponsePtrOutput {
	return i.ToAuthInfoResponsePtrOutputWithContext(context.Background())
}

func (i *authInfoResponsePtrType) ToAuthInfoResponsePtrOutputWithContext(ctx context.Context) AuthInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthInfoResponsePtrOutput)
}

type AuthInfoResponseOutput struct{ *pulumi.OutputState }

func (AuthInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthInfoResponse)(nil)).Elem()
}

func (o AuthInfoResponseOutput) ToAuthInfoResponseOutput() AuthInfoResponseOutput {
	return o
}

func (o AuthInfoResponseOutput) ToAuthInfoResponseOutputWithContext(ctx context.Context) AuthInfoResponseOutput {
	return o
}

func (o AuthInfoResponseOutput) ToAuthInfoResponsePtrOutput() AuthInfoResponsePtrOutput {
	return o.ToAuthInfoResponsePtrOutputWithContext(context.Background())
}

func (o AuthInfoResponseOutput) ToAuthInfoResponsePtrOutputWithContext(ctx context.Context) AuthInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthInfoResponse) *AuthInfoResponse {
		return &v
	}).(AuthInfoResponsePtrOutput)
}

func (o AuthInfoResponseOutput) ExpiresIn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AuthInfoResponse) *int { return v.ExpiresIn }).(pulumi.IntPtrOutput)
}

func (o AuthInfoResponseOutput) RefreshToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthInfoResponse) *string { return v.RefreshToken }).(pulumi.StringPtrOutput)
}

func (o AuthInfoResponseOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthInfoResponse) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o AuthInfoResponseOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v AuthInfoResponse) string { return v.Token }).(pulumi.StringOutput)
}

func (o AuthInfoResponseOutput) TokenType() pulumi.StringOutput {
	return o.ApplyT(func(v AuthInfoResponse) string { return v.TokenType }).(pulumi.StringOutput)
}

type AuthInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (AuthInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthInfoResponse)(nil)).Elem()
}

func (o AuthInfoResponsePtrOutput) ToAuthInfoResponsePtrOutput() AuthInfoResponsePtrOutput {
	return o
}

func (o AuthInfoResponsePtrOutput) ToAuthInfoResponsePtrOutputWithContext(ctx context.Context) AuthInfoResponsePtrOutput {
	return o
}

func (o AuthInfoResponsePtrOutput) Elem() AuthInfoResponseOutput {
	return o.ApplyT(func(v *AuthInfoResponse) AuthInfoResponse {
		if v != nil {
			return *v
		}
		var ret AuthInfoResponse
		return ret
	}).(AuthInfoResponseOutput)
}

func (o AuthInfoResponsePtrOutput) ExpiresIn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AuthInfoResponse) *int {
		if v == nil {
			return nil
		}
		return v.ExpiresIn
	}).(pulumi.IntPtrOutput)
}

func (o AuthInfoResponsePtrOutput) RefreshToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.RefreshToken
	}).(pulumi.StringPtrOutput)
}

func (o AuthInfoResponsePtrOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Scope
	}).(pulumi.StringPtrOutput)
}

func (o AuthInfoResponsePtrOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Token
	}).(pulumi.StringPtrOutput)
}

func (o AuthInfoResponsePtrOutput) TokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TokenType
	}).(pulumi.StringPtrOutput)
}

type BaseImageDependencyResponse struct {
	Digest     *string `pulumi:"digest"`
	Registry   *string `pulumi:"registry"`
	Repository *string `pulumi:"repository"`
	Tag        *string `pulumi:"tag"`
	Type       *string `pulumi:"type"`
}





type BaseImageDependencyResponseInput interface {
	pulumi.Input

	ToBaseImageDependencyResponseOutput() BaseImageDependencyResponseOutput
	ToBaseImageDependencyResponseOutputWithContext(context.Context) BaseImageDependencyResponseOutput
}

type BaseImageDependencyResponseArgs struct {
	Digest     pulumi.StringPtrInput `pulumi:"digest"`
	Registry   pulumi.StringPtrInput `pulumi:"registry"`
	Repository pulumi.StringPtrInput `pulumi:"repository"`
	Tag        pulumi.StringPtrInput `pulumi:"tag"`
	Type       pulumi.StringPtrInput `pulumi:"type"`
}

func (BaseImageDependencyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BaseImageDependencyResponse)(nil)).Elem()
}

func (i BaseImageDependencyResponseArgs) ToBaseImageDependencyResponseOutput() BaseImageDependencyResponseOutput {
	return i.ToBaseImageDependencyResponseOutputWithContext(context.Background())
}

func (i BaseImageDependencyResponseArgs) ToBaseImageDependencyResponseOutputWithContext(ctx context.Context) BaseImageDependencyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaseImageDependencyResponseOutput)
}





type BaseImageDependencyResponseArrayInput interface {
	pulumi.Input

	ToBaseImageDependencyResponseArrayOutput() BaseImageDependencyResponseArrayOutput
	ToBaseImageDependencyResponseArrayOutputWithContext(context.Context) BaseImageDependencyResponseArrayOutput
}

type BaseImageDependencyResponseArray []BaseImageDependencyResponseInput

func (BaseImageDependencyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BaseImageDependencyResponse)(nil)).Elem()
}

func (i BaseImageDependencyResponseArray) ToBaseImageDependencyResponseArrayOutput() BaseImageDependencyResponseArrayOutput {
	return i.ToBaseImageDependencyResponseArrayOutputWithContext(context.Background())
}

func (i BaseImageDependencyResponseArray) ToBaseImageDependencyResponseArrayOutputWithContext(ctx context.Context) BaseImageDependencyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaseImageDependencyResponseArrayOutput)
}

type BaseImageDependencyResponseOutput struct{ *pulumi.OutputState }

func (BaseImageDependencyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BaseImageDependencyResponse)(nil)).Elem()
}

func (o BaseImageDependencyResponseOutput) ToBaseImageDependencyResponseOutput() BaseImageDependencyResponseOutput {
	return o
}

func (o BaseImageDependencyResponseOutput) ToBaseImageDependencyResponseOutputWithContext(ctx context.Context) BaseImageDependencyResponseOutput {
	return o
}

func (o BaseImageDependencyResponseOutput) Digest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaseImageDependencyResponse) *string { return v.Digest }).(pulumi.StringPtrOutput)
}

func (o BaseImageDependencyResponseOutput) Registry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaseImageDependencyResponse) *string { return v.Registry }).(pulumi.StringPtrOutput)
}

func (o BaseImageDependencyResponseOutput) Repository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaseImageDependencyResponse) *string { return v.Repository }).(pulumi.StringPtrOutput)
}

func (o BaseImageDependencyResponseOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaseImageDependencyResponse) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

func (o BaseImageDependencyResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaseImageDependencyResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type BaseImageDependencyResponseArrayOutput struct{ *pulumi.OutputState }

func (BaseImageDependencyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BaseImageDependencyResponse)(nil)).Elem()
}

func (o BaseImageDependencyResponseArrayOutput) ToBaseImageDependencyResponseArrayOutput() BaseImageDependencyResponseArrayOutput {
	return o
}

func (o BaseImageDependencyResponseArrayOutput) ToBaseImageDependencyResponseArrayOutputWithContext(ctx context.Context) BaseImageDependencyResponseArrayOutput {
	return o
}

func (o BaseImageDependencyResponseArrayOutput) Index(i pulumi.IntInput) BaseImageDependencyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BaseImageDependencyResponse {
		return vs[0].([]BaseImageDependencyResponse)[vs[1].(int)]
	}).(BaseImageDependencyResponseOutput)
}

type BaseImageTrigger struct {
	BaseImageTriggerType string  `pulumi:"baseImageTriggerType"`
	Name                 string  `pulumi:"name"`
	Status               *string `pulumi:"status"`
}





type BaseImageTriggerInput interface {
	pulumi.Input

	ToBaseImageTriggerOutput() BaseImageTriggerOutput
	ToBaseImageTriggerOutputWithContext(context.Context) BaseImageTriggerOutput
}

type BaseImageTriggerArgs struct {
	BaseImageTriggerType pulumi.StringInput    `pulumi:"baseImageTriggerType"`
	Name                 pulumi.StringInput    `pulumi:"name"`
	Status               pulumi.StringPtrInput `pulumi:"status"`
}

func (BaseImageTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BaseImageTrigger)(nil)).Elem()
}

func (i BaseImageTriggerArgs) ToBaseImageTriggerOutput() BaseImageTriggerOutput {
	return i.ToBaseImageTriggerOutputWithContext(context.Background())
}

func (i BaseImageTriggerArgs) ToBaseImageTriggerOutputWithContext(ctx context.Context) BaseImageTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaseImageTriggerOutput)
}

func (i BaseImageTriggerArgs) ToBaseImageTriggerPtrOutput() BaseImageTriggerPtrOutput {
	return i.ToBaseImageTriggerPtrOutputWithContext(context.Background())
}

func (i BaseImageTriggerArgs) ToBaseImageTriggerPtrOutputWithContext(ctx context.Context) BaseImageTriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaseImageTriggerOutput).ToBaseImageTriggerPtrOutputWithContext(ctx)
}









type BaseImageTriggerPtrInput interface {
	pulumi.Input

	ToBaseImageTriggerPtrOutput() BaseImageTriggerPtrOutput
	ToBaseImageTriggerPtrOutputWithContext(context.Context) BaseImageTriggerPtrOutput
}

type baseImageTriggerPtrType BaseImageTriggerArgs

func BaseImageTriggerPtr(v *BaseImageTriggerArgs) BaseImageTriggerPtrInput {
	return (*baseImageTriggerPtrType)(v)
}

func (*baseImageTriggerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BaseImageTrigger)(nil)).Elem()
}

func (i *baseImageTriggerPtrType) ToBaseImageTriggerPtrOutput() BaseImageTriggerPtrOutput {
	return i.ToBaseImageTriggerPtrOutputWithContext(context.Background())
}

func (i *baseImageTriggerPtrType) ToBaseImageTriggerPtrOutputWithContext(ctx context.Context) BaseImageTriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaseImageTriggerPtrOutput)
}

type BaseImageTriggerOutput struct{ *pulumi.OutputState }

func (BaseImageTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BaseImageTrigger)(nil)).Elem()
}

func (o BaseImageTriggerOutput) ToBaseImageTriggerOutput() BaseImageTriggerOutput {
	return o
}

func (o BaseImageTriggerOutput) ToBaseImageTriggerOutputWithContext(ctx context.Context) BaseImageTriggerOutput {
	return o
}

func (o BaseImageTriggerOutput) ToBaseImageTriggerPtrOutput() BaseImageTriggerPtrOutput {
	return o.ToBaseImageTriggerPtrOutputWithContext(context.Background())
}

func (o BaseImageTriggerOutput) ToBaseImageTriggerPtrOutputWithContext(ctx context.Context) BaseImageTriggerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BaseImageTrigger) *BaseImageTrigger {
		return &v
	}).(BaseImageTriggerPtrOutput)
}

func (o BaseImageTriggerOutput) BaseImageTriggerType() pulumi.StringOutput {
	return o.ApplyT(func(v BaseImageTrigger) string { return v.BaseImageTriggerType }).(pulumi.StringOutput)
}

func (o BaseImageTriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v BaseImageTrigger) string { return v.Name }).(pulumi.StringOutput)
}

func (o BaseImageTriggerOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaseImageTrigger) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type BaseImageTriggerPtrOutput struct{ *pulumi.OutputState }

func (BaseImageTriggerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BaseImageTrigger)(nil)).Elem()
}

func (o BaseImageTriggerPtrOutput) ToBaseImageTriggerPtrOutput() BaseImageTriggerPtrOutput {
	return o
}

func (o BaseImageTriggerPtrOutput) ToBaseImageTriggerPtrOutputWithContext(ctx context.Context) BaseImageTriggerPtrOutput {
	return o
}

func (o BaseImageTriggerPtrOutput) Elem() BaseImageTriggerOutput {
	return o.ApplyT(func(v *BaseImageTrigger) BaseImageTrigger {
		if v != nil {
			return *v
		}
		var ret BaseImageTrigger
		return ret
	}).(BaseImageTriggerOutput)
}

func (o BaseImageTriggerPtrOutput) BaseImageTriggerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaseImageTrigger) *string {
		if v == nil {
			return nil
		}
		return &v.BaseImageTriggerType
	}).(pulumi.StringPtrOutput)
}

func (o BaseImageTriggerPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaseImageTrigger) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o BaseImageTriggerPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaseImageTrigger) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type BaseImageTriggerResponse struct {
	BaseImageTriggerType string  `pulumi:"baseImageTriggerType"`
	Name                 string  `pulumi:"name"`
	Status               *string `pulumi:"status"`
}





type BaseImageTriggerResponseInput interface {
	pulumi.Input

	ToBaseImageTriggerResponseOutput() BaseImageTriggerResponseOutput
	ToBaseImageTriggerResponseOutputWithContext(context.Context) BaseImageTriggerResponseOutput
}

type BaseImageTriggerResponseArgs struct {
	BaseImageTriggerType pulumi.StringInput    `pulumi:"baseImageTriggerType"`
	Name                 pulumi.StringInput    `pulumi:"name"`
	Status               pulumi.StringPtrInput `pulumi:"status"`
}

func (BaseImageTriggerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BaseImageTriggerResponse)(nil)).Elem()
}

func (i BaseImageTriggerResponseArgs) ToBaseImageTriggerResponseOutput() BaseImageTriggerResponseOutput {
	return i.ToBaseImageTriggerResponseOutputWithContext(context.Background())
}

func (i BaseImageTriggerResponseArgs) ToBaseImageTriggerResponseOutputWithContext(ctx context.Context) BaseImageTriggerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaseImageTriggerResponseOutput)
}

func (i BaseImageTriggerResponseArgs) ToBaseImageTriggerResponsePtrOutput() BaseImageTriggerResponsePtrOutput {
	return i.ToBaseImageTriggerResponsePtrOutputWithContext(context.Background())
}

func (i BaseImageTriggerResponseArgs) ToBaseImageTriggerResponsePtrOutputWithContext(ctx context.Context) BaseImageTriggerResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaseImageTriggerResponseOutput).ToBaseImageTriggerResponsePtrOutputWithContext(ctx)
}









type BaseImageTriggerResponsePtrInput interface {
	pulumi.Input

	ToBaseImageTriggerResponsePtrOutput() BaseImageTriggerResponsePtrOutput
	ToBaseImageTriggerResponsePtrOutputWithContext(context.Context) BaseImageTriggerResponsePtrOutput
}

type baseImageTriggerResponsePtrType BaseImageTriggerResponseArgs

func BaseImageTriggerResponsePtr(v *BaseImageTriggerResponseArgs) BaseImageTriggerResponsePtrInput {
	return (*baseImageTriggerResponsePtrType)(v)
}

func (*baseImageTriggerResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BaseImageTriggerResponse)(nil)).Elem()
}

func (i *baseImageTriggerResponsePtrType) ToBaseImageTriggerResponsePtrOutput() BaseImageTriggerResponsePtrOutput {
	return i.ToBaseImageTriggerResponsePtrOutputWithContext(context.Background())
}

func (i *baseImageTriggerResponsePtrType) ToBaseImageTriggerResponsePtrOutputWithContext(ctx context.Context) BaseImageTriggerResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaseImageTriggerResponsePtrOutput)
}

type BaseImageTriggerResponseOutput struct{ *pulumi.OutputState }

func (BaseImageTriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BaseImageTriggerResponse)(nil)).Elem()
}

func (o BaseImageTriggerResponseOutput) ToBaseImageTriggerResponseOutput() BaseImageTriggerResponseOutput {
	return o
}

func (o BaseImageTriggerResponseOutput) ToBaseImageTriggerResponseOutputWithContext(ctx context.Context) BaseImageTriggerResponseOutput {
	return o
}

func (o BaseImageTriggerResponseOutput) ToBaseImageTriggerResponsePtrOutput() BaseImageTriggerResponsePtrOutput {
	return o.ToBaseImageTriggerResponsePtrOutputWithContext(context.Background())
}

func (o BaseImageTriggerResponseOutput) ToBaseImageTriggerResponsePtrOutputWithContext(ctx context.Context) BaseImageTriggerResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BaseImageTriggerResponse) *BaseImageTriggerResponse {
		return &v
	}).(BaseImageTriggerResponsePtrOutput)
}

func (o BaseImageTriggerResponseOutput) BaseImageTriggerType() pulumi.StringOutput {
	return o.ApplyT(func(v BaseImageTriggerResponse) string { return v.BaseImageTriggerType }).(pulumi.StringOutput)
}

func (o BaseImageTriggerResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v BaseImageTriggerResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o BaseImageTriggerResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaseImageTriggerResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type BaseImageTriggerResponsePtrOutput struct{ *pulumi.OutputState }

func (BaseImageTriggerResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BaseImageTriggerResponse)(nil)).Elem()
}

func (o BaseImageTriggerResponsePtrOutput) ToBaseImageTriggerResponsePtrOutput() BaseImageTriggerResponsePtrOutput {
	return o
}

func (o BaseImageTriggerResponsePtrOutput) ToBaseImageTriggerResponsePtrOutputWithContext(ctx context.Context) BaseImageTriggerResponsePtrOutput {
	return o
}

func (o BaseImageTriggerResponsePtrOutput) Elem() BaseImageTriggerResponseOutput {
	return o.ApplyT(func(v *BaseImageTriggerResponse) BaseImageTriggerResponse {
		if v != nil {
			return *v
		}
		var ret BaseImageTriggerResponse
		return ret
	}).(BaseImageTriggerResponseOutput)
}

func (o BaseImageTriggerResponsePtrOutput) BaseImageTriggerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaseImageTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return &v.BaseImageTriggerType
	}).(pulumi.StringPtrOutput)
}

func (o BaseImageTriggerResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaseImageTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o BaseImageTriggerResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaseImageTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type Credentials struct {
	CustomRegistries map[string]CustomRegistryCredentials `pulumi:"customRegistries"`
	SourceRegistry   *SourceRegistryCredentials           `pulumi:"sourceRegistry"`
}





type CredentialsInput interface {
	pulumi.Input

	ToCredentialsOutput() CredentialsOutput
	ToCredentialsOutputWithContext(context.Context) CredentialsOutput
}

type CredentialsArgs struct {
	CustomRegistries CustomRegistryCredentialsMapInput `pulumi:"customRegistries"`
	SourceRegistry   SourceRegistryCredentialsPtrInput `pulumi:"sourceRegistry"`
}

func (CredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Credentials)(nil)).Elem()
}

func (i CredentialsArgs) ToCredentialsOutput() CredentialsOutput {
	return i.ToCredentialsOutputWithContext(context.Background())
}

func (i CredentialsArgs) ToCredentialsOutputWithContext(ctx context.Context) CredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialsOutput)
}

func (i CredentialsArgs) ToCredentialsPtrOutput() CredentialsPtrOutput {
	return i.ToCredentialsPtrOutputWithContext(context.Background())
}

func (i CredentialsArgs) ToCredentialsPtrOutputWithContext(ctx context.Context) CredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialsOutput).ToCredentialsPtrOutputWithContext(ctx)
}









type CredentialsPtrInput interface {
	pulumi.Input

	ToCredentialsPtrOutput() CredentialsPtrOutput
	ToCredentialsPtrOutputWithContext(context.Context) CredentialsPtrOutput
}

type credentialsPtrType CredentialsArgs

func CredentialsPtr(v *CredentialsArgs) CredentialsPtrInput {
	return (*credentialsPtrType)(v)
}

func (*credentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Credentials)(nil)).Elem()
}

func (i *credentialsPtrType) ToCredentialsPtrOutput() CredentialsPtrOutput {
	return i.ToCredentialsPtrOutputWithContext(context.Background())
}

func (i *credentialsPtrType) ToCredentialsPtrOutputWithContext(ctx context.Context) CredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialsPtrOutput)
}

type CredentialsOutput struct{ *pulumi.OutputState }

func (CredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Credentials)(nil)).Elem()
}

func (o CredentialsOutput) ToCredentialsOutput() CredentialsOutput {
	return o
}

func (o CredentialsOutput) ToCredentialsOutputWithContext(ctx context.Context) CredentialsOutput {
	return o
}

func (o CredentialsOutput) ToCredentialsPtrOutput() CredentialsPtrOutput {
	return o.ToCredentialsPtrOutputWithContext(context.Background())
}

func (o CredentialsOutput) ToCredentialsPtrOutputWithContext(ctx context.Context) CredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Credentials) *Credentials {
		return &v
	}).(CredentialsPtrOutput)
}

func (o CredentialsOutput) CustomRegistries() CustomRegistryCredentialsMapOutput {
	return o.ApplyT(func(v Credentials) map[string]CustomRegistryCredentials { return v.CustomRegistries }).(CustomRegistryCredentialsMapOutput)
}

func (o CredentialsOutput) SourceRegistry() SourceRegistryCredentialsPtrOutput {
	return o.ApplyT(func(v Credentials) *SourceRegistryCredentials { return v.SourceRegistry }).(SourceRegistryCredentialsPtrOutput)
}

type CredentialsPtrOutput struct{ *pulumi.OutputState }

func (CredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Credentials)(nil)).Elem()
}

func (o CredentialsPtrOutput) ToCredentialsPtrOutput() CredentialsPtrOutput {
	return o
}

func (o CredentialsPtrOutput) ToCredentialsPtrOutputWithContext(ctx context.Context) CredentialsPtrOutput {
	return o
}

func (o CredentialsPtrOutput) Elem() CredentialsOutput {
	return o.ApplyT(func(v *Credentials) Credentials {
		if v != nil {
			return *v
		}
		var ret Credentials
		return ret
	}).(CredentialsOutput)
}

func (o CredentialsPtrOutput) CustomRegistries() CustomRegistryCredentialsMapOutput {
	return o.ApplyT(func(v *Credentials) map[string]CustomRegistryCredentials {
		if v == nil {
			return nil
		}
		return v.CustomRegistries
	}).(CustomRegistryCredentialsMapOutput)
}

func (o CredentialsPtrOutput) SourceRegistry() SourceRegistryCredentialsPtrOutput {
	return o.ApplyT(func(v *Credentials) *SourceRegistryCredentials {
		if v == nil {
			return nil
		}
		return v.SourceRegistry
	}).(SourceRegistryCredentialsPtrOutput)
}

type CredentialsResponse struct {
	CustomRegistries map[string]CustomRegistryCredentialsResponse `pulumi:"customRegistries"`
	SourceRegistry   *SourceRegistryCredentialsResponse           `pulumi:"sourceRegistry"`
}





type CredentialsResponseInput interface {
	pulumi.Input

	ToCredentialsResponseOutput() CredentialsResponseOutput
	ToCredentialsResponseOutputWithContext(context.Context) CredentialsResponseOutput
}

type CredentialsResponseArgs struct {
	CustomRegistries CustomRegistryCredentialsResponseMapInput `pulumi:"customRegistries"`
	SourceRegistry   SourceRegistryCredentialsResponsePtrInput `pulumi:"sourceRegistry"`
}

func (CredentialsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CredentialsResponse)(nil)).Elem()
}

func (i CredentialsResponseArgs) ToCredentialsResponseOutput() CredentialsResponseOutput {
	return i.ToCredentialsResponseOutputWithContext(context.Background())
}

func (i CredentialsResponseArgs) ToCredentialsResponseOutputWithContext(ctx context.Context) CredentialsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialsResponseOutput)
}

func (i CredentialsResponseArgs) ToCredentialsResponsePtrOutput() CredentialsResponsePtrOutput {
	return i.ToCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i CredentialsResponseArgs) ToCredentialsResponsePtrOutputWithContext(ctx context.Context) CredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialsResponseOutput).ToCredentialsResponsePtrOutputWithContext(ctx)
}









type CredentialsResponsePtrInput interface {
	pulumi.Input

	ToCredentialsResponsePtrOutput() CredentialsResponsePtrOutput
	ToCredentialsResponsePtrOutputWithContext(context.Context) CredentialsResponsePtrOutput
}

type credentialsResponsePtrType CredentialsResponseArgs

func CredentialsResponsePtr(v *CredentialsResponseArgs) CredentialsResponsePtrInput {
	return (*credentialsResponsePtrType)(v)
}

func (*credentialsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CredentialsResponse)(nil)).Elem()
}

func (i *credentialsResponsePtrType) ToCredentialsResponsePtrOutput() CredentialsResponsePtrOutput {
	return i.ToCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i *credentialsResponsePtrType) ToCredentialsResponsePtrOutputWithContext(ctx context.Context) CredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialsResponsePtrOutput)
}

type CredentialsResponseOutput struct{ *pulumi.OutputState }

func (CredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CredentialsResponse)(nil)).Elem()
}

func (o CredentialsResponseOutput) ToCredentialsResponseOutput() CredentialsResponseOutput {
	return o
}

func (o CredentialsResponseOutput) ToCredentialsResponseOutputWithContext(ctx context.Context) CredentialsResponseOutput {
	return o
}

func (o CredentialsResponseOutput) ToCredentialsResponsePtrOutput() CredentialsResponsePtrOutput {
	return o.ToCredentialsResponsePtrOutputWithContext(context.Background())
}

func (o CredentialsResponseOutput) ToCredentialsResponsePtrOutputWithContext(ctx context.Context) CredentialsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CredentialsResponse) *CredentialsResponse {
		return &v
	}).(CredentialsResponsePtrOutput)
}

func (o CredentialsResponseOutput) CustomRegistries() CustomRegistryCredentialsResponseMapOutput {
	return o.ApplyT(func(v CredentialsResponse) map[string]CustomRegistryCredentialsResponse { return v.CustomRegistries }).(CustomRegistryCredentialsResponseMapOutput)
}

func (o CredentialsResponseOutput) SourceRegistry() SourceRegistryCredentialsResponsePtrOutput {
	return o.ApplyT(func(v CredentialsResponse) *SourceRegistryCredentialsResponse { return v.SourceRegistry }).(SourceRegistryCredentialsResponsePtrOutput)
}

type CredentialsResponsePtrOutput struct{ *pulumi.OutputState }

func (CredentialsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CredentialsResponse)(nil)).Elem()
}

func (o CredentialsResponsePtrOutput) ToCredentialsResponsePtrOutput() CredentialsResponsePtrOutput {
	return o
}

func (o CredentialsResponsePtrOutput) ToCredentialsResponsePtrOutputWithContext(ctx context.Context) CredentialsResponsePtrOutput {
	return o
}

func (o CredentialsResponsePtrOutput) Elem() CredentialsResponseOutput {
	return o.ApplyT(func(v *CredentialsResponse) CredentialsResponse {
		if v != nil {
			return *v
		}
		var ret CredentialsResponse
		return ret
	}).(CredentialsResponseOutput)
}

func (o CredentialsResponsePtrOutput) CustomRegistries() CustomRegistryCredentialsResponseMapOutput {
	return o.ApplyT(func(v *CredentialsResponse) map[string]CustomRegistryCredentialsResponse {
		if v == nil {
			return nil
		}
		return v.CustomRegistries
	}).(CustomRegistryCredentialsResponseMapOutput)
}

func (o CredentialsResponsePtrOutput) SourceRegistry() SourceRegistryCredentialsResponsePtrOutput {
	return o.ApplyT(func(v *CredentialsResponse) *SourceRegistryCredentialsResponse {
		if v == nil {
			return nil
		}
		return v.SourceRegistry
	}).(SourceRegistryCredentialsResponsePtrOutput)
}

type CustomRegistryCredentials struct {
	Password *SecretObject `pulumi:"password"`
	UserName *SecretObject `pulumi:"userName"`
}





type CustomRegistryCredentialsInput interface {
	pulumi.Input

	ToCustomRegistryCredentialsOutput() CustomRegistryCredentialsOutput
	ToCustomRegistryCredentialsOutputWithContext(context.Context) CustomRegistryCredentialsOutput
}

type CustomRegistryCredentialsArgs struct {
	Password SecretObjectPtrInput `pulumi:"password"`
	UserName SecretObjectPtrInput `pulumi:"userName"`
}

func (CustomRegistryCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRegistryCredentials)(nil)).Elem()
}

func (i CustomRegistryCredentialsArgs) ToCustomRegistryCredentialsOutput() CustomRegistryCredentialsOutput {
	return i.ToCustomRegistryCredentialsOutputWithContext(context.Background())
}

func (i CustomRegistryCredentialsArgs) ToCustomRegistryCredentialsOutputWithContext(ctx context.Context) CustomRegistryCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRegistryCredentialsOutput)
}





type CustomRegistryCredentialsMapInput interface {
	pulumi.Input

	ToCustomRegistryCredentialsMapOutput() CustomRegistryCredentialsMapOutput
	ToCustomRegistryCredentialsMapOutputWithContext(context.Context) CustomRegistryCredentialsMapOutput
}

type CustomRegistryCredentialsMap map[string]CustomRegistryCredentialsInput

func (CustomRegistryCredentialsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomRegistryCredentials)(nil)).Elem()
}

func (i CustomRegistryCredentialsMap) ToCustomRegistryCredentialsMapOutput() CustomRegistryCredentialsMapOutput {
	return i.ToCustomRegistryCredentialsMapOutputWithContext(context.Background())
}

func (i CustomRegistryCredentialsMap) ToCustomRegistryCredentialsMapOutputWithContext(ctx context.Context) CustomRegistryCredentialsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRegistryCredentialsMapOutput)
}

type CustomRegistryCredentialsOutput struct{ *pulumi.OutputState }

func (CustomRegistryCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRegistryCredentials)(nil)).Elem()
}

func (o CustomRegistryCredentialsOutput) ToCustomRegistryCredentialsOutput() CustomRegistryCredentialsOutput {
	return o
}

func (o CustomRegistryCredentialsOutput) ToCustomRegistryCredentialsOutputWithContext(ctx context.Context) CustomRegistryCredentialsOutput {
	return o
}

func (o CustomRegistryCredentialsOutput) Password() SecretObjectPtrOutput {
	return o.ApplyT(func(v CustomRegistryCredentials) *SecretObject { return v.Password }).(SecretObjectPtrOutput)
}

func (o CustomRegistryCredentialsOutput) UserName() SecretObjectPtrOutput {
	return o.ApplyT(func(v CustomRegistryCredentials) *SecretObject { return v.UserName }).(SecretObjectPtrOutput)
}

type CustomRegistryCredentialsMapOutput struct{ *pulumi.OutputState }

func (CustomRegistryCredentialsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomRegistryCredentials)(nil)).Elem()
}

func (o CustomRegistryCredentialsMapOutput) ToCustomRegistryCredentialsMapOutput() CustomRegistryCredentialsMapOutput {
	return o
}

func (o CustomRegistryCredentialsMapOutput) ToCustomRegistryCredentialsMapOutputWithContext(ctx context.Context) CustomRegistryCredentialsMapOutput {
	return o
}

func (o CustomRegistryCredentialsMapOutput) MapIndex(k pulumi.StringInput) CustomRegistryCredentialsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CustomRegistryCredentials {
		return vs[0].(map[string]CustomRegistryCredentials)[vs[1].(string)]
	}).(CustomRegistryCredentialsOutput)
}

type CustomRegistryCredentialsResponse struct {
	Password *SecretObjectResponse `pulumi:"password"`
	UserName *SecretObjectResponse `pulumi:"userName"`
}





type CustomRegistryCredentialsResponseInput interface {
	pulumi.Input

	ToCustomRegistryCredentialsResponseOutput() CustomRegistryCredentialsResponseOutput
	ToCustomRegistryCredentialsResponseOutputWithContext(context.Context) CustomRegistryCredentialsResponseOutput
}

type CustomRegistryCredentialsResponseArgs struct {
	Password SecretObjectResponsePtrInput `pulumi:"password"`
	UserName SecretObjectResponsePtrInput `pulumi:"userName"`
}

func (CustomRegistryCredentialsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRegistryCredentialsResponse)(nil)).Elem()
}

func (i CustomRegistryCredentialsResponseArgs) ToCustomRegistryCredentialsResponseOutput() CustomRegistryCredentialsResponseOutput {
	return i.ToCustomRegistryCredentialsResponseOutputWithContext(context.Background())
}

func (i CustomRegistryCredentialsResponseArgs) ToCustomRegistryCredentialsResponseOutputWithContext(ctx context.Context) CustomRegistryCredentialsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRegistryCredentialsResponseOutput)
}





type CustomRegistryCredentialsResponseMapInput interface {
	pulumi.Input

	ToCustomRegistryCredentialsResponseMapOutput() CustomRegistryCredentialsResponseMapOutput
	ToCustomRegistryCredentialsResponseMapOutputWithContext(context.Context) CustomRegistryCredentialsResponseMapOutput
}

type CustomRegistryCredentialsResponseMap map[string]CustomRegistryCredentialsResponseInput

func (CustomRegistryCredentialsResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomRegistryCredentialsResponse)(nil)).Elem()
}

func (i CustomRegistryCredentialsResponseMap) ToCustomRegistryCredentialsResponseMapOutput() CustomRegistryCredentialsResponseMapOutput {
	return i.ToCustomRegistryCredentialsResponseMapOutputWithContext(context.Background())
}

func (i CustomRegistryCredentialsResponseMap) ToCustomRegistryCredentialsResponseMapOutputWithContext(ctx context.Context) CustomRegistryCredentialsResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRegistryCredentialsResponseMapOutput)
}

type CustomRegistryCredentialsResponseOutput struct{ *pulumi.OutputState }

func (CustomRegistryCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRegistryCredentialsResponse)(nil)).Elem()
}

func (o CustomRegistryCredentialsResponseOutput) ToCustomRegistryCredentialsResponseOutput() CustomRegistryCredentialsResponseOutput {
	return o
}

func (o CustomRegistryCredentialsResponseOutput) ToCustomRegistryCredentialsResponseOutputWithContext(ctx context.Context) CustomRegistryCredentialsResponseOutput {
	return o
}

func (o CustomRegistryCredentialsResponseOutput) Password() SecretObjectResponsePtrOutput {
	return o.ApplyT(func(v CustomRegistryCredentialsResponse) *SecretObjectResponse { return v.Password }).(SecretObjectResponsePtrOutput)
}

func (o CustomRegistryCredentialsResponseOutput) UserName() SecretObjectResponsePtrOutput {
	return o.ApplyT(func(v CustomRegistryCredentialsResponse) *SecretObjectResponse { return v.UserName }).(SecretObjectResponsePtrOutput)
}

type CustomRegistryCredentialsResponseMapOutput struct{ *pulumi.OutputState }

func (CustomRegistryCredentialsResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomRegistryCredentialsResponse)(nil)).Elem()
}

func (o CustomRegistryCredentialsResponseMapOutput) ToCustomRegistryCredentialsResponseMapOutput() CustomRegistryCredentialsResponseMapOutput {
	return o
}

func (o CustomRegistryCredentialsResponseMapOutput) ToCustomRegistryCredentialsResponseMapOutputWithContext(ctx context.Context) CustomRegistryCredentialsResponseMapOutput {
	return o
}

func (o CustomRegistryCredentialsResponseMapOutput) MapIndex(k pulumi.StringInput) CustomRegistryCredentialsResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CustomRegistryCredentialsResponse {
		return vs[0].(map[string]CustomRegistryCredentialsResponse)[vs[1].(string)]
	}).(CustomRegistryCredentialsResponseOutput)
}

type DockerBuildStepResponse struct {
	Arguments             []ArgumentResponse            `pulumi:"arguments"`
	BaseImageDependencies []BaseImageDependencyResponse `pulumi:"baseImageDependencies"`
	ContextAccessToken    *string                       `pulumi:"contextAccessToken"`
	ContextPath           *string                       `pulumi:"contextPath"`
	DockerFilePath        string                        `pulumi:"dockerFilePath"`
	ImageNames            []string                      `pulumi:"imageNames"`
	IsPushEnabled         *bool                         `pulumi:"isPushEnabled"`
	NoCache               *bool                         `pulumi:"noCache"`
	Target                *string                       `pulumi:"target"`
	Type                  string                        `pulumi:"type"`
}





type DockerBuildStepResponseInput interface {
	pulumi.Input

	ToDockerBuildStepResponseOutput() DockerBuildStepResponseOutput
	ToDockerBuildStepResponseOutputWithContext(context.Context) DockerBuildStepResponseOutput
}

type DockerBuildStepResponseArgs struct {
	Arguments             ArgumentResponseArrayInput            `pulumi:"arguments"`
	BaseImageDependencies BaseImageDependencyResponseArrayInput `pulumi:"baseImageDependencies"`
	ContextAccessToken    pulumi.StringPtrInput                 `pulumi:"contextAccessToken"`
	ContextPath           pulumi.StringPtrInput                 `pulumi:"contextPath"`
	DockerFilePath        pulumi.StringInput                    `pulumi:"dockerFilePath"`
	ImageNames            pulumi.StringArrayInput               `pulumi:"imageNames"`
	IsPushEnabled         pulumi.BoolPtrInput                   `pulumi:"isPushEnabled"`
	NoCache               pulumi.BoolPtrInput                   `pulumi:"noCache"`
	Target                pulumi.StringPtrInput                 `pulumi:"target"`
	Type                  pulumi.StringInput                    `pulumi:"type"`
}

func (DockerBuildStepResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DockerBuildStepResponse)(nil)).Elem()
}

func (i DockerBuildStepResponseArgs) ToDockerBuildStepResponseOutput() DockerBuildStepResponseOutput {
	return i.ToDockerBuildStepResponseOutputWithContext(context.Background())
}

func (i DockerBuildStepResponseArgs) ToDockerBuildStepResponseOutputWithContext(ctx context.Context) DockerBuildStepResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerBuildStepResponseOutput)
}

type DockerBuildStepResponseOutput struct{ *pulumi.OutputState }

func (DockerBuildStepResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DockerBuildStepResponse)(nil)).Elem()
}

func (o DockerBuildStepResponseOutput) ToDockerBuildStepResponseOutput() DockerBuildStepResponseOutput {
	return o
}

func (o DockerBuildStepResponseOutput) ToDockerBuildStepResponseOutputWithContext(ctx context.Context) DockerBuildStepResponseOutput {
	return o
}

func (o DockerBuildStepResponseOutput) Arguments() ArgumentResponseArrayOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) []ArgumentResponse { return v.Arguments }).(ArgumentResponseArrayOutput)
}

func (o DockerBuildStepResponseOutput) BaseImageDependencies() BaseImageDependencyResponseArrayOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) []BaseImageDependencyResponse { return v.BaseImageDependencies }).(BaseImageDependencyResponseArrayOutput)
}

func (o DockerBuildStepResponseOutput) ContextAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) *string { return v.ContextAccessToken }).(pulumi.StringPtrOutput)
}

func (o DockerBuildStepResponseOutput) ContextPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) *string { return v.ContextPath }).(pulumi.StringPtrOutput)
}

func (o DockerBuildStepResponseOutput) DockerFilePath() pulumi.StringOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) string { return v.DockerFilePath }).(pulumi.StringOutput)
}

func (o DockerBuildStepResponseOutput) ImageNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) []string { return v.ImageNames }).(pulumi.StringArrayOutput)
}

func (o DockerBuildStepResponseOutput) IsPushEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) *bool { return v.IsPushEnabled }).(pulumi.BoolPtrOutput)
}

func (o DockerBuildStepResponseOutput) NoCache() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) *bool { return v.NoCache }).(pulumi.BoolPtrOutput)
}

func (o DockerBuildStepResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o DockerBuildStepResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) string { return v.Type }).(pulumi.StringOutput)
}

type EncodedTaskStepResponse struct {
	BaseImageDependencies []BaseImageDependencyResponse `pulumi:"baseImageDependencies"`
	ContextAccessToken    *string                       `pulumi:"contextAccessToken"`
	ContextPath           *string                       `pulumi:"contextPath"`
	EncodedTaskContent    string                        `pulumi:"encodedTaskContent"`
	EncodedValuesContent  *string                       `pulumi:"encodedValuesContent"`
	Type                  string                        `pulumi:"type"`
	Values                []SetValueResponse            `pulumi:"values"`
}





type EncodedTaskStepResponseInput interface {
	pulumi.Input

	ToEncodedTaskStepResponseOutput() EncodedTaskStepResponseOutput
	ToEncodedTaskStepResponseOutputWithContext(context.Context) EncodedTaskStepResponseOutput
}

type EncodedTaskStepResponseArgs struct {
	BaseImageDependencies BaseImageDependencyResponseArrayInput `pulumi:"baseImageDependencies"`
	ContextAccessToken    pulumi.StringPtrInput                 `pulumi:"contextAccessToken"`
	ContextPath           pulumi.StringPtrInput                 `pulumi:"contextPath"`
	EncodedTaskContent    pulumi.StringInput                    `pulumi:"encodedTaskContent"`
	EncodedValuesContent  pulumi.StringPtrInput                 `pulumi:"encodedValuesContent"`
	Type                  pulumi.StringInput                    `pulumi:"type"`
	Values                SetValueResponseArrayInput            `pulumi:"values"`
}

func (EncodedTaskStepResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncodedTaskStepResponse)(nil)).Elem()
}

func (i EncodedTaskStepResponseArgs) ToEncodedTaskStepResponseOutput() EncodedTaskStepResponseOutput {
	return i.ToEncodedTaskStepResponseOutputWithContext(context.Background())
}

func (i EncodedTaskStepResponseArgs) ToEncodedTaskStepResponseOutputWithContext(ctx context.Context) EncodedTaskStepResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncodedTaskStepResponseOutput)
}

type EncodedTaskStepResponseOutput struct{ *pulumi.OutputState }

func (EncodedTaskStepResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncodedTaskStepResponse)(nil)).Elem()
}

func (o EncodedTaskStepResponseOutput) ToEncodedTaskStepResponseOutput() EncodedTaskStepResponseOutput {
	return o
}

func (o EncodedTaskStepResponseOutput) ToEncodedTaskStepResponseOutputWithContext(ctx context.Context) EncodedTaskStepResponseOutput {
	return o
}

func (o EncodedTaskStepResponseOutput) BaseImageDependencies() BaseImageDependencyResponseArrayOutput {
	return o.ApplyT(func(v EncodedTaskStepResponse) []BaseImageDependencyResponse { return v.BaseImageDependencies }).(BaseImageDependencyResponseArrayOutput)
}

func (o EncodedTaskStepResponseOutput) ContextAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncodedTaskStepResponse) *string { return v.ContextAccessToken }).(pulumi.StringPtrOutput)
}

func (o EncodedTaskStepResponseOutput) ContextPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncodedTaskStepResponse) *string { return v.ContextPath }).(pulumi.StringPtrOutput)
}

func (o EncodedTaskStepResponseOutput) EncodedTaskContent() pulumi.StringOutput {
	return o.ApplyT(func(v EncodedTaskStepResponse) string { return v.EncodedTaskContent }).(pulumi.StringOutput)
}

func (o EncodedTaskStepResponseOutput) EncodedValuesContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncodedTaskStepResponse) *string { return v.EncodedValuesContent }).(pulumi.StringPtrOutput)
}

func (o EncodedTaskStepResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EncodedTaskStepResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o EncodedTaskStepResponseOutput) Values() SetValueResponseArrayOutput {
	return o.ApplyT(func(v EncodedTaskStepResponse) []SetValueResponse { return v.Values }).(SetValueResponseArrayOutput)
}

type FileTaskStepResponse struct {
	BaseImageDependencies []BaseImageDependencyResponse `pulumi:"baseImageDependencies"`
	ContextAccessToken    *string                       `pulumi:"contextAccessToken"`
	ContextPath           *string                       `pulumi:"contextPath"`
	TaskFilePath          string                        `pulumi:"taskFilePath"`
	Type                  string                        `pulumi:"type"`
	Values                []SetValueResponse            `pulumi:"values"`
	ValuesFilePath        *string                       `pulumi:"valuesFilePath"`
}





type FileTaskStepResponseInput interface {
	pulumi.Input

	ToFileTaskStepResponseOutput() FileTaskStepResponseOutput
	ToFileTaskStepResponseOutputWithContext(context.Context) FileTaskStepResponseOutput
}

type FileTaskStepResponseArgs struct {
	BaseImageDependencies BaseImageDependencyResponseArrayInput `pulumi:"baseImageDependencies"`
	ContextAccessToken    pulumi.StringPtrInput                 `pulumi:"contextAccessToken"`
	ContextPath           pulumi.StringPtrInput                 `pulumi:"contextPath"`
	TaskFilePath          pulumi.StringInput                    `pulumi:"taskFilePath"`
	Type                  pulumi.StringInput                    `pulumi:"type"`
	Values                SetValueResponseArrayInput            `pulumi:"values"`
	ValuesFilePath        pulumi.StringPtrInput                 `pulumi:"valuesFilePath"`
}

func (FileTaskStepResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FileTaskStepResponse)(nil)).Elem()
}

func (i FileTaskStepResponseArgs) ToFileTaskStepResponseOutput() FileTaskStepResponseOutput {
	return i.ToFileTaskStepResponseOutputWithContext(context.Background())
}

func (i FileTaskStepResponseArgs) ToFileTaskStepResponseOutputWithContext(ctx context.Context) FileTaskStepResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileTaskStepResponseOutput)
}

type FileTaskStepResponseOutput struct{ *pulumi.OutputState }

func (FileTaskStepResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileTaskStepResponse)(nil)).Elem()
}

func (o FileTaskStepResponseOutput) ToFileTaskStepResponseOutput() FileTaskStepResponseOutput {
	return o
}

func (o FileTaskStepResponseOutput) ToFileTaskStepResponseOutputWithContext(ctx context.Context) FileTaskStepResponseOutput {
	return o
}

func (o FileTaskStepResponseOutput) BaseImageDependencies() BaseImageDependencyResponseArrayOutput {
	return o.ApplyT(func(v FileTaskStepResponse) []BaseImageDependencyResponse { return v.BaseImageDependencies }).(BaseImageDependencyResponseArrayOutput)
}

func (o FileTaskStepResponseOutput) ContextAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileTaskStepResponse) *string { return v.ContextAccessToken }).(pulumi.StringPtrOutput)
}

func (o FileTaskStepResponseOutput) ContextPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileTaskStepResponse) *string { return v.ContextPath }).(pulumi.StringPtrOutput)
}

func (o FileTaskStepResponseOutput) TaskFilePath() pulumi.StringOutput {
	return o.ApplyT(func(v FileTaskStepResponse) string { return v.TaskFilePath }).(pulumi.StringOutput)
}

func (o FileTaskStepResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v FileTaskStepResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o FileTaskStepResponseOutput) Values() SetValueResponseArrayOutput {
	return o.ApplyT(func(v FileTaskStepResponse) []SetValueResponse { return v.Values }).(SetValueResponseArrayOutput)
}

func (o FileTaskStepResponseOutput) ValuesFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileTaskStepResponse) *string { return v.ValuesFilePath }).(pulumi.StringPtrOutput)
}

type PlatformProperties struct {
	Architecture *string `pulumi:"architecture"`
	Os           string  `pulumi:"os"`
	Variant      *string `pulumi:"variant"`
}





type PlatformPropertiesInput interface {
	pulumi.Input

	ToPlatformPropertiesOutput() PlatformPropertiesOutput
	ToPlatformPropertiesOutputWithContext(context.Context) PlatformPropertiesOutput
}

type PlatformPropertiesArgs struct {
	Architecture pulumi.StringPtrInput `pulumi:"architecture"`
	Os           pulumi.StringInput    `pulumi:"os"`
	Variant      pulumi.StringPtrInput `pulumi:"variant"`
}

func (PlatformPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PlatformProperties)(nil)).Elem()
}

func (i PlatformPropertiesArgs) ToPlatformPropertiesOutput() PlatformPropertiesOutput {
	return i.ToPlatformPropertiesOutputWithContext(context.Background())
}

func (i PlatformPropertiesArgs) ToPlatformPropertiesOutputWithContext(ctx context.Context) PlatformPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformPropertiesOutput)
}

func (i PlatformPropertiesArgs) ToPlatformPropertiesPtrOutput() PlatformPropertiesPtrOutput {
	return i.ToPlatformPropertiesPtrOutputWithContext(context.Background())
}

func (i PlatformPropertiesArgs) ToPlatformPropertiesPtrOutputWithContext(ctx context.Context) PlatformPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformPropertiesOutput).ToPlatformPropertiesPtrOutputWithContext(ctx)
}









type PlatformPropertiesPtrInput interface {
	pulumi.Input

	ToPlatformPropertiesPtrOutput() PlatformPropertiesPtrOutput
	ToPlatformPropertiesPtrOutputWithContext(context.Context) PlatformPropertiesPtrOutput
}

type platformPropertiesPtrType PlatformPropertiesArgs

func PlatformPropertiesPtr(v *PlatformPropertiesArgs) PlatformPropertiesPtrInput {
	return (*platformPropertiesPtrType)(v)
}

func (*platformPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PlatformProperties)(nil)).Elem()
}

func (i *platformPropertiesPtrType) ToPlatformPropertiesPtrOutput() PlatformPropertiesPtrOutput {
	return i.ToPlatformPropertiesPtrOutputWithContext(context.Background())
}

func (i *platformPropertiesPtrType) ToPlatformPropertiesPtrOutputWithContext(ctx context.Context) PlatformPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformPropertiesPtrOutput)
}

type PlatformPropertiesOutput struct{ *pulumi.OutputState }

func (PlatformPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlatformProperties)(nil)).Elem()
}

func (o PlatformPropertiesOutput) ToPlatformPropertiesOutput() PlatformPropertiesOutput {
	return o
}

func (o PlatformPropertiesOutput) ToPlatformPropertiesOutputWithContext(ctx context.Context) PlatformPropertiesOutput {
	return o
}

func (o PlatformPropertiesOutput) ToPlatformPropertiesPtrOutput() PlatformPropertiesPtrOutput {
	return o.ToPlatformPropertiesPtrOutputWithContext(context.Background())
}

func (o PlatformPropertiesOutput) ToPlatformPropertiesPtrOutputWithContext(ctx context.Context) PlatformPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PlatformProperties) *PlatformProperties {
		return &v
	}).(PlatformPropertiesPtrOutput)
}

func (o PlatformPropertiesOutput) Architecture() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlatformProperties) *string { return v.Architecture }).(pulumi.StringPtrOutput)
}

func (o PlatformPropertiesOutput) Os() pulumi.StringOutput {
	return o.ApplyT(func(v PlatformProperties) string { return v.Os }).(pulumi.StringOutput)
}

func (o PlatformPropertiesOutput) Variant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlatformProperties) *string { return v.Variant }).(pulumi.StringPtrOutput)
}

type PlatformPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PlatformPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlatformProperties)(nil)).Elem()
}

func (o PlatformPropertiesPtrOutput) ToPlatformPropertiesPtrOutput() PlatformPropertiesPtrOutput {
	return o
}

func (o PlatformPropertiesPtrOutput) ToPlatformPropertiesPtrOutputWithContext(ctx context.Context) PlatformPropertiesPtrOutput {
	return o
}

func (o PlatformPropertiesPtrOutput) Elem() PlatformPropertiesOutput {
	return o.ApplyT(func(v *PlatformProperties) PlatformProperties {
		if v != nil {
			return *v
		}
		var ret PlatformProperties
		return ret
	}).(PlatformPropertiesOutput)
}

func (o PlatformPropertiesPtrOutput) Architecture() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlatformProperties) *string {
		if v == nil {
			return nil
		}
		return v.Architecture
	}).(pulumi.StringPtrOutput)
}

func (o PlatformPropertiesPtrOutput) Os() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlatformProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Os
	}).(pulumi.StringPtrOutput)
}

func (o PlatformPropertiesPtrOutput) Variant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlatformProperties) *string {
		if v == nil {
			return nil
		}
		return v.Variant
	}).(pulumi.StringPtrOutput)
}

type PlatformPropertiesResponse struct {
	Architecture *string `pulumi:"architecture"`
	Os           string  `pulumi:"os"`
	Variant      *string `pulumi:"variant"`
}





type PlatformPropertiesResponseInput interface {
	pulumi.Input

	ToPlatformPropertiesResponseOutput() PlatformPropertiesResponseOutput
	ToPlatformPropertiesResponseOutputWithContext(context.Context) PlatformPropertiesResponseOutput
}

type PlatformPropertiesResponseArgs struct {
	Architecture pulumi.StringPtrInput `pulumi:"architecture"`
	Os           pulumi.StringInput    `pulumi:"os"`
	Variant      pulumi.StringPtrInput `pulumi:"variant"`
}

func (PlatformPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PlatformPropertiesResponse)(nil)).Elem()
}

func (i PlatformPropertiesResponseArgs) ToPlatformPropertiesResponseOutput() PlatformPropertiesResponseOutput {
	return i.ToPlatformPropertiesResponseOutputWithContext(context.Background())
}

func (i PlatformPropertiesResponseArgs) ToPlatformPropertiesResponseOutputWithContext(ctx context.Context) PlatformPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformPropertiesResponseOutput)
}

func (i PlatformPropertiesResponseArgs) ToPlatformPropertiesResponsePtrOutput() PlatformPropertiesResponsePtrOutput {
	return i.ToPlatformPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i PlatformPropertiesResponseArgs) ToPlatformPropertiesResponsePtrOutputWithContext(ctx context.Context) PlatformPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformPropertiesResponseOutput).ToPlatformPropertiesResponsePtrOutputWithContext(ctx)
}









type PlatformPropertiesResponsePtrInput interface {
	pulumi.Input

	ToPlatformPropertiesResponsePtrOutput() PlatformPropertiesResponsePtrOutput
	ToPlatformPropertiesResponsePtrOutputWithContext(context.Context) PlatformPropertiesResponsePtrOutput
}

type platformPropertiesResponsePtrType PlatformPropertiesResponseArgs

func PlatformPropertiesResponsePtr(v *PlatformPropertiesResponseArgs) PlatformPropertiesResponsePtrInput {
	return (*platformPropertiesResponsePtrType)(v)
}

func (*platformPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PlatformPropertiesResponse)(nil)).Elem()
}

func (i *platformPropertiesResponsePtrType) ToPlatformPropertiesResponsePtrOutput() PlatformPropertiesResponsePtrOutput {
	return i.ToPlatformPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *platformPropertiesResponsePtrType) ToPlatformPropertiesResponsePtrOutputWithContext(ctx context.Context) PlatformPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformPropertiesResponsePtrOutput)
}

type PlatformPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PlatformPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlatformPropertiesResponse)(nil)).Elem()
}

func (o PlatformPropertiesResponseOutput) ToPlatformPropertiesResponseOutput() PlatformPropertiesResponseOutput {
	return o
}

func (o PlatformPropertiesResponseOutput) ToPlatformPropertiesResponseOutputWithContext(ctx context.Context) PlatformPropertiesResponseOutput {
	return o
}

func (o PlatformPropertiesResponseOutput) ToPlatformPropertiesResponsePtrOutput() PlatformPropertiesResponsePtrOutput {
	return o.ToPlatformPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o PlatformPropertiesResponseOutput) ToPlatformPropertiesResponsePtrOutputWithContext(ctx context.Context) PlatformPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PlatformPropertiesResponse) *PlatformPropertiesResponse {
		return &v
	}).(PlatformPropertiesResponsePtrOutput)
}

func (o PlatformPropertiesResponseOutput) Architecture() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlatformPropertiesResponse) *string { return v.Architecture }).(pulumi.StringPtrOutput)
}

func (o PlatformPropertiesResponseOutput) Os() pulumi.StringOutput {
	return o.ApplyT(func(v PlatformPropertiesResponse) string { return v.Os }).(pulumi.StringOutput)
}

func (o PlatformPropertiesResponseOutput) Variant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlatformPropertiesResponse) *string { return v.Variant }).(pulumi.StringPtrOutput)
}

type PlatformPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PlatformPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlatformPropertiesResponse)(nil)).Elem()
}

func (o PlatformPropertiesResponsePtrOutput) ToPlatformPropertiesResponsePtrOutput() PlatformPropertiesResponsePtrOutput {
	return o
}

func (o PlatformPropertiesResponsePtrOutput) ToPlatformPropertiesResponsePtrOutputWithContext(ctx context.Context) PlatformPropertiesResponsePtrOutput {
	return o
}

func (o PlatformPropertiesResponsePtrOutput) Elem() PlatformPropertiesResponseOutput {
	return o.ApplyT(func(v *PlatformPropertiesResponse) PlatformPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PlatformPropertiesResponse
		return ret
	}).(PlatformPropertiesResponseOutput)
}

func (o PlatformPropertiesResponsePtrOutput) Architecture() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlatformPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Architecture
	}).(pulumi.StringPtrOutput)
}

func (o PlatformPropertiesResponsePtrOutput) Os() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlatformPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Os
	}).(pulumi.StringPtrOutput)
}

func (o PlatformPropertiesResponsePtrOutput) Variant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlatformPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Variant
	}).(pulumi.StringPtrOutput)
}

type SecretObject struct {
	Type  *string `pulumi:"type"`
	Value *string `pulumi:"value"`
}





type SecretObjectInput interface {
	pulumi.Input

	ToSecretObjectOutput() SecretObjectOutput
	ToSecretObjectOutputWithContext(context.Context) SecretObjectOutput
}

type SecretObjectArgs struct {
	Type  pulumi.StringPtrInput `pulumi:"type"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (SecretObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretObject)(nil)).Elem()
}

func (i SecretObjectArgs) ToSecretObjectOutput() SecretObjectOutput {
	return i.ToSecretObjectOutputWithContext(context.Background())
}

func (i SecretObjectArgs) ToSecretObjectOutputWithContext(ctx context.Context) SecretObjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretObjectOutput)
}

func (i SecretObjectArgs) ToSecretObjectPtrOutput() SecretObjectPtrOutput {
	return i.ToSecretObjectPtrOutputWithContext(context.Background())
}

func (i SecretObjectArgs) ToSecretObjectPtrOutputWithContext(ctx context.Context) SecretObjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretObjectOutput).ToSecretObjectPtrOutputWithContext(ctx)
}









type SecretObjectPtrInput interface {
	pulumi.Input

	ToSecretObjectPtrOutput() SecretObjectPtrOutput
	ToSecretObjectPtrOutputWithContext(context.Context) SecretObjectPtrOutput
}

type secretObjectPtrType SecretObjectArgs

func SecretObjectPtr(v *SecretObjectArgs) SecretObjectPtrInput {
	return (*secretObjectPtrType)(v)
}

func (*secretObjectPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretObject)(nil)).Elem()
}

func (i *secretObjectPtrType) ToSecretObjectPtrOutput() SecretObjectPtrOutput {
	return i.ToSecretObjectPtrOutputWithContext(context.Background())
}

func (i *secretObjectPtrType) ToSecretObjectPtrOutputWithContext(ctx context.Context) SecretObjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretObjectPtrOutput)
}

type SecretObjectOutput struct{ *pulumi.OutputState }

func (SecretObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretObject)(nil)).Elem()
}

func (o SecretObjectOutput) ToSecretObjectOutput() SecretObjectOutput {
	return o
}

func (o SecretObjectOutput) ToSecretObjectOutputWithContext(ctx context.Context) SecretObjectOutput {
	return o
}

func (o SecretObjectOutput) ToSecretObjectPtrOutput() SecretObjectPtrOutput {
	return o.ToSecretObjectPtrOutputWithContext(context.Background())
}

func (o SecretObjectOutput) ToSecretObjectPtrOutputWithContext(ctx context.Context) SecretObjectPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecretObject) *SecretObject {
		return &v
	}).(SecretObjectPtrOutput)
}

func (o SecretObjectOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretObject) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o SecretObjectOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretObject) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type SecretObjectPtrOutput struct{ *pulumi.OutputState }

func (SecretObjectPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretObject)(nil)).Elem()
}

func (o SecretObjectPtrOutput) ToSecretObjectPtrOutput() SecretObjectPtrOutput {
	return o
}

func (o SecretObjectPtrOutput) ToSecretObjectPtrOutputWithContext(ctx context.Context) SecretObjectPtrOutput {
	return o
}

func (o SecretObjectPtrOutput) Elem() SecretObjectOutput {
	return o.ApplyT(func(v *SecretObject) SecretObject {
		if v != nil {
			return *v
		}
		var ret SecretObject
		return ret
	}).(SecretObjectOutput)
}

func (o SecretObjectPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretObject) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o SecretObjectPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretObject) *string {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.StringPtrOutput)
}

type SecretObjectResponse struct {
	Type  *string `pulumi:"type"`
	Value *string `pulumi:"value"`
}





type SecretObjectResponseInput interface {
	pulumi.Input

	ToSecretObjectResponseOutput() SecretObjectResponseOutput
	ToSecretObjectResponseOutputWithContext(context.Context) SecretObjectResponseOutput
}

type SecretObjectResponseArgs struct {
	Type  pulumi.StringPtrInput `pulumi:"type"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (SecretObjectResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretObjectResponse)(nil)).Elem()
}

func (i SecretObjectResponseArgs) ToSecretObjectResponseOutput() SecretObjectResponseOutput {
	return i.ToSecretObjectResponseOutputWithContext(context.Background())
}

func (i SecretObjectResponseArgs) ToSecretObjectResponseOutputWithContext(ctx context.Context) SecretObjectResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretObjectResponseOutput)
}

func (i SecretObjectResponseArgs) ToSecretObjectResponsePtrOutput() SecretObjectResponsePtrOutput {
	return i.ToSecretObjectResponsePtrOutputWithContext(context.Background())
}

func (i SecretObjectResponseArgs) ToSecretObjectResponsePtrOutputWithContext(ctx context.Context) SecretObjectResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretObjectResponseOutput).ToSecretObjectResponsePtrOutputWithContext(ctx)
}









type SecretObjectResponsePtrInput interface {
	pulumi.Input

	ToSecretObjectResponsePtrOutput() SecretObjectResponsePtrOutput
	ToSecretObjectResponsePtrOutputWithContext(context.Context) SecretObjectResponsePtrOutput
}

type secretObjectResponsePtrType SecretObjectResponseArgs

func SecretObjectResponsePtr(v *SecretObjectResponseArgs) SecretObjectResponsePtrInput {
	return (*secretObjectResponsePtrType)(v)
}

func (*secretObjectResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretObjectResponse)(nil)).Elem()
}

func (i *secretObjectResponsePtrType) ToSecretObjectResponsePtrOutput() SecretObjectResponsePtrOutput {
	return i.ToSecretObjectResponsePtrOutputWithContext(context.Background())
}

func (i *secretObjectResponsePtrType) ToSecretObjectResponsePtrOutputWithContext(ctx context.Context) SecretObjectResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretObjectResponsePtrOutput)
}

type SecretObjectResponseOutput struct{ *pulumi.OutputState }

func (SecretObjectResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretObjectResponse)(nil)).Elem()
}

func (o SecretObjectResponseOutput) ToSecretObjectResponseOutput() SecretObjectResponseOutput {
	return o
}

func (o SecretObjectResponseOutput) ToSecretObjectResponseOutputWithContext(ctx context.Context) SecretObjectResponseOutput {
	return o
}

func (o SecretObjectResponseOutput) ToSecretObjectResponsePtrOutput() SecretObjectResponsePtrOutput {
	return o.ToSecretObjectResponsePtrOutputWithContext(context.Background())
}

func (o SecretObjectResponseOutput) ToSecretObjectResponsePtrOutputWithContext(ctx context.Context) SecretObjectResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecretObjectResponse) *SecretObjectResponse {
		return &v
	}).(SecretObjectResponsePtrOutput)
}

func (o SecretObjectResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretObjectResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o SecretObjectResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretObjectResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type SecretObjectResponsePtrOutput struct{ *pulumi.OutputState }

func (SecretObjectResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretObjectResponse)(nil)).Elem()
}

func (o SecretObjectResponsePtrOutput) ToSecretObjectResponsePtrOutput() SecretObjectResponsePtrOutput {
	return o
}

func (o SecretObjectResponsePtrOutput) ToSecretObjectResponsePtrOutputWithContext(ctx context.Context) SecretObjectResponsePtrOutput {
	return o
}

func (o SecretObjectResponsePtrOutput) Elem() SecretObjectResponseOutput {
	return o.ApplyT(func(v *SecretObjectResponse) SecretObjectResponse {
		if v != nil {
			return *v
		}
		var ret SecretObjectResponse
		return ret
	}).(SecretObjectResponseOutput)
}

func (o SecretObjectResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretObjectResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o SecretObjectResponsePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretObjectResponse) *string {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.StringPtrOutput)
}

type SetValueResponse struct {
	IsSecret *bool  `pulumi:"isSecret"`
	Name     string `pulumi:"name"`
	Value    string `pulumi:"value"`
}





type SetValueResponseInput interface {
	pulumi.Input

	ToSetValueResponseOutput() SetValueResponseOutput
	ToSetValueResponseOutputWithContext(context.Context) SetValueResponseOutput
}

type SetValueResponseArgs struct {
	IsSecret pulumi.BoolPtrInput `pulumi:"isSecret"`
	Name     pulumi.StringInput  `pulumi:"name"`
	Value    pulumi.StringInput  `pulumi:"value"`
}

func (SetValueResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SetValueResponse)(nil)).Elem()
}

func (i SetValueResponseArgs) ToSetValueResponseOutput() SetValueResponseOutput {
	return i.ToSetValueResponseOutputWithContext(context.Background())
}

func (i SetValueResponseArgs) ToSetValueResponseOutputWithContext(ctx context.Context) SetValueResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SetValueResponseOutput)
}





type SetValueResponseArrayInput interface {
	pulumi.Input

	ToSetValueResponseArrayOutput() SetValueResponseArrayOutput
	ToSetValueResponseArrayOutputWithContext(context.Context) SetValueResponseArrayOutput
}

type SetValueResponseArray []SetValueResponseInput

func (SetValueResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SetValueResponse)(nil)).Elem()
}

func (i SetValueResponseArray) ToSetValueResponseArrayOutput() SetValueResponseArrayOutput {
	return i.ToSetValueResponseArrayOutputWithContext(context.Background())
}

func (i SetValueResponseArray) ToSetValueResponseArrayOutputWithContext(ctx context.Context) SetValueResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SetValueResponseArrayOutput)
}

type SetValueResponseOutput struct{ *pulumi.OutputState }

func (SetValueResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SetValueResponse)(nil)).Elem()
}

func (o SetValueResponseOutput) ToSetValueResponseOutput() SetValueResponseOutput {
	return o
}

func (o SetValueResponseOutput) ToSetValueResponseOutputWithContext(ctx context.Context) SetValueResponseOutput {
	return o
}

func (o SetValueResponseOutput) GetIsSecret() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SetValueResponse) *bool { return v.IsSecret }).(pulumi.BoolPtrOutput)
}

func (o SetValueResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SetValueResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SetValueResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SetValueResponse) string { return v.Value }).(pulumi.StringOutput)
}

type SetValueResponseArrayOutput struct{ *pulumi.OutputState }

func (SetValueResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SetValueResponse)(nil)).Elem()
}

func (o SetValueResponseArrayOutput) ToSetValueResponseArrayOutput() SetValueResponseArrayOutput {
	return o
}

func (o SetValueResponseArrayOutput) ToSetValueResponseArrayOutputWithContext(ctx context.Context) SetValueResponseArrayOutput {
	return o
}

func (o SetValueResponseArrayOutput) Index(i pulumi.IntInput) SetValueResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SetValueResponse {
		return vs[0].([]SetValueResponse)[vs[1].(int)]
	}).(SetValueResponseOutput)
}

type SourceProperties struct {
	Branch                      *string   `pulumi:"branch"`
	RepositoryUrl               string    `pulumi:"repositoryUrl"`
	SourceControlAuthProperties *AuthInfo `pulumi:"sourceControlAuthProperties"`
	SourceControlType           string    `pulumi:"sourceControlType"`
}





type SourcePropertiesInput interface {
	pulumi.Input

	ToSourcePropertiesOutput() SourcePropertiesOutput
	ToSourcePropertiesOutputWithContext(context.Context) SourcePropertiesOutput
}

type SourcePropertiesArgs struct {
	Branch                      pulumi.StringPtrInput `pulumi:"branch"`
	RepositoryUrl               pulumi.StringInput    `pulumi:"repositoryUrl"`
	SourceControlAuthProperties AuthInfoPtrInput      `pulumi:"sourceControlAuthProperties"`
	SourceControlType           pulumi.StringInput    `pulumi:"sourceControlType"`
}

func (SourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceProperties)(nil)).Elem()
}

func (i SourcePropertiesArgs) ToSourcePropertiesOutput() SourcePropertiesOutput {
	return i.ToSourcePropertiesOutputWithContext(context.Background())
}

func (i SourcePropertiesArgs) ToSourcePropertiesOutputWithContext(ctx context.Context) SourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePropertiesOutput)
}

type SourcePropertiesOutput struct{ *pulumi.OutputState }

func (SourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceProperties)(nil)).Elem()
}

func (o SourcePropertiesOutput) ToSourcePropertiesOutput() SourcePropertiesOutput {
	return o
}

func (o SourcePropertiesOutput) ToSourcePropertiesOutputWithContext(ctx context.Context) SourcePropertiesOutput {
	return o
}

func (o SourcePropertiesOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceProperties) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o SourcePropertiesOutput) RepositoryUrl() pulumi.StringOutput {
	return o.ApplyT(func(v SourceProperties) string { return v.RepositoryUrl }).(pulumi.StringOutput)
}

func (o SourcePropertiesOutput) SourceControlAuthProperties() AuthInfoPtrOutput {
	return o.ApplyT(func(v SourceProperties) *AuthInfo { return v.SourceControlAuthProperties }).(AuthInfoPtrOutput)
}

func (o SourcePropertiesOutput) SourceControlType() pulumi.StringOutput {
	return o.ApplyT(func(v SourceProperties) string { return v.SourceControlType }).(pulumi.StringOutput)
}

type SourcePropertiesResponse struct {
	Branch                      *string           `pulumi:"branch"`
	RepositoryUrl               string            `pulumi:"repositoryUrl"`
	SourceControlAuthProperties *AuthInfoResponse `pulumi:"sourceControlAuthProperties"`
	SourceControlType           string            `pulumi:"sourceControlType"`
}





type SourcePropertiesResponseInput interface {
	pulumi.Input

	ToSourcePropertiesResponseOutput() SourcePropertiesResponseOutput
	ToSourcePropertiesResponseOutputWithContext(context.Context) SourcePropertiesResponseOutput
}

type SourcePropertiesResponseArgs struct {
	Branch                      pulumi.StringPtrInput    `pulumi:"branch"`
	RepositoryUrl               pulumi.StringInput       `pulumi:"repositoryUrl"`
	SourceControlAuthProperties AuthInfoResponsePtrInput `pulumi:"sourceControlAuthProperties"`
	SourceControlType           pulumi.StringInput       `pulumi:"sourceControlType"`
}

func (SourcePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourcePropertiesResponse)(nil)).Elem()
}

func (i SourcePropertiesResponseArgs) ToSourcePropertiesResponseOutput() SourcePropertiesResponseOutput {
	return i.ToSourcePropertiesResponseOutputWithContext(context.Background())
}

func (i SourcePropertiesResponseArgs) ToSourcePropertiesResponseOutputWithContext(ctx context.Context) SourcePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePropertiesResponseOutput)
}

type SourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (SourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourcePropertiesResponse)(nil)).Elem()
}

func (o SourcePropertiesResponseOutput) ToSourcePropertiesResponseOutput() SourcePropertiesResponseOutput {
	return o
}

func (o SourcePropertiesResponseOutput) ToSourcePropertiesResponseOutputWithContext(ctx context.Context) SourcePropertiesResponseOutput {
	return o
}

func (o SourcePropertiesResponseOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourcePropertiesResponse) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o SourcePropertiesResponseOutput) RepositoryUrl() pulumi.StringOutput {
	return o.ApplyT(func(v SourcePropertiesResponse) string { return v.RepositoryUrl }).(pulumi.StringOutput)
}

func (o SourcePropertiesResponseOutput) SourceControlAuthProperties() AuthInfoResponsePtrOutput {
	return o.ApplyT(func(v SourcePropertiesResponse) *AuthInfoResponse { return v.SourceControlAuthProperties }).(AuthInfoResponsePtrOutput)
}

func (o SourcePropertiesResponseOutput) SourceControlType() pulumi.StringOutput {
	return o.ApplyT(func(v SourcePropertiesResponse) string { return v.SourceControlType }).(pulumi.StringOutput)
}

type SourceRegistryCredentials struct {
	LoginMode *string `pulumi:"loginMode"`
}





type SourceRegistryCredentialsInput interface {
	pulumi.Input

	ToSourceRegistryCredentialsOutput() SourceRegistryCredentialsOutput
	ToSourceRegistryCredentialsOutputWithContext(context.Context) SourceRegistryCredentialsOutput
}

type SourceRegistryCredentialsArgs struct {
	LoginMode pulumi.StringPtrInput `pulumi:"loginMode"`
}

func (SourceRegistryCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceRegistryCredentials)(nil)).Elem()
}

func (i SourceRegistryCredentialsArgs) ToSourceRegistryCredentialsOutput() SourceRegistryCredentialsOutput {
	return i.ToSourceRegistryCredentialsOutputWithContext(context.Background())
}

func (i SourceRegistryCredentialsArgs) ToSourceRegistryCredentialsOutputWithContext(ctx context.Context) SourceRegistryCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRegistryCredentialsOutput)
}

func (i SourceRegistryCredentialsArgs) ToSourceRegistryCredentialsPtrOutput() SourceRegistryCredentialsPtrOutput {
	return i.ToSourceRegistryCredentialsPtrOutputWithContext(context.Background())
}

func (i SourceRegistryCredentialsArgs) ToSourceRegistryCredentialsPtrOutputWithContext(ctx context.Context) SourceRegistryCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRegistryCredentialsOutput).ToSourceRegistryCredentialsPtrOutputWithContext(ctx)
}









type SourceRegistryCredentialsPtrInput interface {
	pulumi.Input

	ToSourceRegistryCredentialsPtrOutput() SourceRegistryCredentialsPtrOutput
	ToSourceRegistryCredentialsPtrOutputWithContext(context.Context) SourceRegistryCredentialsPtrOutput
}

type sourceRegistryCredentialsPtrType SourceRegistryCredentialsArgs

func SourceRegistryCredentialsPtr(v *SourceRegistryCredentialsArgs) SourceRegistryCredentialsPtrInput {
	return (*sourceRegistryCredentialsPtrType)(v)
}

func (*sourceRegistryCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceRegistryCredentials)(nil)).Elem()
}

func (i *sourceRegistryCredentialsPtrType) ToSourceRegistryCredentialsPtrOutput() SourceRegistryCredentialsPtrOutput {
	return i.ToSourceRegistryCredentialsPtrOutputWithContext(context.Background())
}

func (i *sourceRegistryCredentialsPtrType) ToSourceRegistryCredentialsPtrOutputWithContext(ctx context.Context) SourceRegistryCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRegistryCredentialsPtrOutput)
}

type SourceRegistryCredentialsOutput struct{ *pulumi.OutputState }

func (SourceRegistryCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceRegistryCredentials)(nil)).Elem()
}

func (o SourceRegistryCredentialsOutput) ToSourceRegistryCredentialsOutput() SourceRegistryCredentialsOutput {
	return o
}

func (o SourceRegistryCredentialsOutput) ToSourceRegistryCredentialsOutputWithContext(ctx context.Context) SourceRegistryCredentialsOutput {
	return o
}

func (o SourceRegistryCredentialsOutput) ToSourceRegistryCredentialsPtrOutput() SourceRegistryCredentialsPtrOutput {
	return o.ToSourceRegistryCredentialsPtrOutputWithContext(context.Background())
}

func (o SourceRegistryCredentialsOutput) ToSourceRegistryCredentialsPtrOutputWithContext(ctx context.Context) SourceRegistryCredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceRegistryCredentials) *SourceRegistryCredentials {
		return &v
	}).(SourceRegistryCredentialsPtrOutput)
}

func (o SourceRegistryCredentialsOutput) LoginMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceRegistryCredentials) *string { return v.LoginMode }).(pulumi.StringPtrOutput)
}

type SourceRegistryCredentialsPtrOutput struct{ *pulumi.OutputState }

func (SourceRegistryCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceRegistryCredentials)(nil)).Elem()
}

func (o SourceRegistryCredentialsPtrOutput) ToSourceRegistryCredentialsPtrOutput() SourceRegistryCredentialsPtrOutput {
	return o
}

func (o SourceRegistryCredentialsPtrOutput) ToSourceRegistryCredentialsPtrOutputWithContext(ctx context.Context) SourceRegistryCredentialsPtrOutput {
	return o
}

func (o SourceRegistryCredentialsPtrOutput) Elem() SourceRegistryCredentialsOutput {
	return o.ApplyT(func(v *SourceRegistryCredentials) SourceRegistryCredentials {
		if v != nil {
			return *v
		}
		var ret SourceRegistryCredentials
		return ret
	}).(SourceRegistryCredentialsOutput)
}

func (o SourceRegistryCredentialsPtrOutput) LoginMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceRegistryCredentials) *string {
		if v == nil {
			return nil
		}
		return v.LoginMode
	}).(pulumi.StringPtrOutput)
}

type SourceRegistryCredentialsResponse struct {
	LoginMode *string `pulumi:"loginMode"`
}





type SourceRegistryCredentialsResponseInput interface {
	pulumi.Input

	ToSourceRegistryCredentialsResponseOutput() SourceRegistryCredentialsResponseOutput
	ToSourceRegistryCredentialsResponseOutputWithContext(context.Context) SourceRegistryCredentialsResponseOutput
}

type SourceRegistryCredentialsResponseArgs struct {
	LoginMode pulumi.StringPtrInput `pulumi:"loginMode"`
}

func (SourceRegistryCredentialsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceRegistryCredentialsResponse)(nil)).Elem()
}

func (i SourceRegistryCredentialsResponseArgs) ToSourceRegistryCredentialsResponseOutput() SourceRegistryCredentialsResponseOutput {
	return i.ToSourceRegistryCredentialsResponseOutputWithContext(context.Background())
}

func (i SourceRegistryCredentialsResponseArgs) ToSourceRegistryCredentialsResponseOutputWithContext(ctx context.Context) SourceRegistryCredentialsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRegistryCredentialsResponseOutput)
}

func (i SourceRegistryCredentialsResponseArgs) ToSourceRegistryCredentialsResponsePtrOutput() SourceRegistryCredentialsResponsePtrOutput {
	return i.ToSourceRegistryCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i SourceRegistryCredentialsResponseArgs) ToSourceRegistryCredentialsResponsePtrOutputWithContext(ctx context.Context) SourceRegistryCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRegistryCredentialsResponseOutput).ToSourceRegistryCredentialsResponsePtrOutputWithContext(ctx)
}









type SourceRegistryCredentialsResponsePtrInput interface {
	pulumi.Input

	ToSourceRegistryCredentialsResponsePtrOutput() SourceRegistryCredentialsResponsePtrOutput
	ToSourceRegistryCredentialsResponsePtrOutputWithContext(context.Context) SourceRegistryCredentialsResponsePtrOutput
}

type sourceRegistryCredentialsResponsePtrType SourceRegistryCredentialsResponseArgs

func SourceRegistryCredentialsResponsePtr(v *SourceRegistryCredentialsResponseArgs) SourceRegistryCredentialsResponsePtrInput {
	return (*sourceRegistryCredentialsResponsePtrType)(v)
}

func (*sourceRegistryCredentialsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceRegistryCredentialsResponse)(nil)).Elem()
}

func (i *sourceRegistryCredentialsResponsePtrType) ToSourceRegistryCredentialsResponsePtrOutput() SourceRegistryCredentialsResponsePtrOutput {
	return i.ToSourceRegistryCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i *sourceRegistryCredentialsResponsePtrType) ToSourceRegistryCredentialsResponsePtrOutputWithContext(ctx context.Context) SourceRegistryCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRegistryCredentialsResponsePtrOutput)
}

type SourceRegistryCredentialsResponseOutput struct{ *pulumi.OutputState }

func (SourceRegistryCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceRegistryCredentialsResponse)(nil)).Elem()
}

func (o SourceRegistryCredentialsResponseOutput) ToSourceRegistryCredentialsResponseOutput() SourceRegistryCredentialsResponseOutput {
	return o
}

func (o SourceRegistryCredentialsResponseOutput) ToSourceRegistryCredentialsResponseOutputWithContext(ctx context.Context) SourceRegistryCredentialsResponseOutput {
	return o
}

func (o SourceRegistryCredentialsResponseOutput) ToSourceRegistryCredentialsResponsePtrOutput() SourceRegistryCredentialsResponsePtrOutput {
	return o.ToSourceRegistryCredentialsResponsePtrOutputWithContext(context.Background())
}

func (o SourceRegistryCredentialsResponseOutput) ToSourceRegistryCredentialsResponsePtrOutputWithContext(ctx context.Context) SourceRegistryCredentialsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceRegistryCredentialsResponse) *SourceRegistryCredentialsResponse {
		return &v
	}).(SourceRegistryCredentialsResponsePtrOutput)
}

func (o SourceRegistryCredentialsResponseOutput) LoginMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceRegistryCredentialsResponse) *string { return v.LoginMode }).(pulumi.StringPtrOutput)
}

type SourceRegistryCredentialsResponsePtrOutput struct{ *pulumi.OutputState }

func (SourceRegistryCredentialsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceRegistryCredentialsResponse)(nil)).Elem()
}

func (o SourceRegistryCredentialsResponsePtrOutput) ToSourceRegistryCredentialsResponsePtrOutput() SourceRegistryCredentialsResponsePtrOutput {
	return o
}

func (o SourceRegistryCredentialsResponsePtrOutput) ToSourceRegistryCredentialsResponsePtrOutputWithContext(ctx context.Context) SourceRegistryCredentialsResponsePtrOutput {
	return o
}

func (o SourceRegistryCredentialsResponsePtrOutput) Elem() SourceRegistryCredentialsResponseOutput {
	return o.ApplyT(func(v *SourceRegistryCredentialsResponse) SourceRegistryCredentialsResponse {
		if v != nil {
			return *v
		}
		var ret SourceRegistryCredentialsResponse
		return ret
	}).(SourceRegistryCredentialsResponseOutput)
}

func (o SourceRegistryCredentialsResponsePtrOutput) LoginMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceRegistryCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.LoginMode
	}).(pulumi.StringPtrOutput)
}

type SourceTrigger struct {
	Name                string           `pulumi:"name"`
	SourceRepository    SourceProperties `pulumi:"sourceRepository"`
	SourceTriggerEvents []string         `pulumi:"sourceTriggerEvents"`
	Status              *string          `pulumi:"status"`
}





type SourceTriggerInput interface {
	pulumi.Input

	ToSourceTriggerOutput() SourceTriggerOutput
	ToSourceTriggerOutputWithContext(context.Context) SourceTriggerOutput
}

type SourceTriggerArgs struct {
	Name                pulumi.StringInput      `pulumi:"name"`
	SourceRepository    SourcePropertiesInput   `pulumi:"sourceRepository"`
	SourceTriggerEvents pulumi.StringArrayInput `pulumi:"sourceTriggerEvents"`
	Status              pulumi.StringPtrInput   `pulumi:"status"`
}

func (SourceTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceTrigger)(nil)).Elem()
}

func (i SourceTriggerArgs) ToSourceTriggerOutput() SourceTriggerOutput {
	return i.ToSourceTriggerOutputWithContext(context.Background())
}

func (i SourceTriggerArgs) ToSourceTriggerOutputWithContext(ctx context.Context) SourceTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceTriggerOutput)
}





type SourceTriggerArrayInput interface {
	pulumi.Input

	ToSourceTriggerArrayOutput() SourceTriggerArrayOutput
	ToSourceTriggerArrayOutputWithContext(context.Context) SourceTriggerArrayOutput
}

type SourceTriggerArray []SourceTriggerInput

func (SourceTriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SourceTrigger)(nil)).Elem()
}

func (i SourceTriggerArray) ToSourceTriggerArrayOutput() SourceTriggerArrayOutput {
	return i.ToSourceTriggerArrayOutputWithContext(context.Background())
}

func (i SourceTriggerArray) ToSourceTriggerArrayOutputWithContext(ctx context.Context) SourceTriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceTriggerArrayOutput)
}

type SourceTriggerOutput struct{ *pulumi.OutputState }

func (SourceTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceTrigger)(nil)).Elem()
}

func (o SourceTriggerOutput) ToSourceTriggerOutput() SourceTriggerOutput {
	return o
}

func (o SourceTriggerOutput) ToSourceTriggerOutputWithContext(ctx context.Context) SourceTriggerOutput {
	return o
}

func (o SourceTriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SourceTrigger) string { return v.Name }).(pulumi.StringOutput)
}

func (o SourceTriggerOutput) SourceRepository() SourcePropertiesOutput {
	return o.ApplyT(func(v SourceTrigger) SourceProperties { return v.SourceRepository }).(SourcePropertiesOutput)
}

func (o SourceTriggerOutput) SourceTriggerEvents() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SourceTrigger) []string { return v.SourceTriggerEvents }).(pulumi.StringArrayOutput)
}

func (o SourceTriggerOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceTrigger) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type SourceTriggerArrayOutput struct{ *pulumi.OutputState }

func (SourceTriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SourceTrigger)(nil)).Elem()
}

func (o SourceTriggerArrayOutput) ToSourceTriggerArrayOutput() SourceTriggerArrayOutput {
	return o
}

func (o SourceTriggerArrayOutput) ToSourceTriggerArrayOutputWithContext(ctx context.Context) SourceTriggerArrayOutput {
	return o
}

func (o SourceTriggerArrayOutput) Index(i pulumi.IntInput) SourceTriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SourceTrigger {
		return vs[0].([]SourceTrigger)[vs[1].(int)]
	}).(SourceTriggerOutput)
}

type SourceTriggerResponse struct {
	Name                string                   `pulumi:"name"`
	SourceRepository    SourcePropertiesResponse `pulumi:"sourceRepository"`
	SourceTriggerEvents []string                 `pulumi:"sourceTriggerEvents"`
	Status              *string                  `pulumi:"status"`
}





type SourceTriggerResponseInput interface {
	pulumi.Input

	ToSourceTriggerResponseOutput() SourceTriggerResponseOutput
	ToSourceTriggerResponseOutputWithContext(context.Context) SourceTriggerResponseOutput
}

type SourceTriggerResponseArgs struct {
	Name                pulumi.StringInput            `pulumi:"name"`
	SourceRepository    SourcePropertiesResponseInput `pulumi:"sourceRepository"`
	SourceTriggerEvents pulumi.StringArrayInput       `pulumi:"sourceTriggerEvents"`
	Status              pulumi.StringPtrInput         `pulumi:"status"`
}

func (SourceTriggerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceTriggerResponse)(nil)).Elem()
}

func (i SourceTriggerResponseArgs) ToSourceTriggerResponseOutput() SourceTriggerResponseOutput {
	return i.ToSourceTriggerResponseOutputWithContext(context.Background())
}

func (i SourceTriggerResponseArgs) ToSourceTriggerResponseOutputWithContext(ctx context.Context) SourceTriggerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceTriggerResponseOutput)
}





type SourceTriggerResponseArrayInput interface {
	pulumi.Input

	ToSourceTriggerResponseArrayOutput() SourceTriggerResponseArrayOutput
	ToSourceTriggerResponseArrayOutputWithContext(context.Context) SourceTriggerResponseArrayOutput
}

type SourceTriggerResponseArray []SourceTriggerResponseInput

func (SourceTriggerResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SourceTriggerResponse)(nil)).Elem()
}

func (i SourceTriggerResponseArray) ToSourceTriggerResponseArrayOutput() SourceTriggerResponseArrayOutput {
	return i.ToSourceTriggerResponseArrayOutputWithContext(context.Background())
}

func (i SourceTriggerResponseArray) ToSourceTriggerResponseArrayOutputWithContext(ctx context.Context) SourceTriggerResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceTriggerResponseArrayOutput)
}

type SourceTriggerResponseOutput struct{ *pulumi.OutputState }

func (SourceTriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceTriggerResponse)(nil)).Elem()
}

func (o SourceTriggerResponseOutput) ToSourceTriggerResponseOutput() SourceTriggerResponseOutput {
	return o
}

func (o SourceTriggerResponseOutput) ToSourceTriggerResponseOutputWithContext(ctx context.Context) SourceTriggerResponseOutput {
	return o
}

func (o SourceTriggerResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SourceTriggerResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SourceTriggerResponseOutput) SourceRepository() SourcePropertiesResponseOutput {
	return o.ApplyT(func(v SourceTriggerResponse) SourcePropertiesResponse { return v.SourceRepository }).(SourcePropertiesResponseOutput)
}

func (o SourceTriggerResponseOutput) SourceTriggerEvents() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SourceTriggerResponse) []string { return v.SourceTriggerEvents }).(pulumi.StringArrayOutput)
}

func (o SourceTriggerResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceTriggerResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type SourceTriggerResponseArrayOutput struct{ *pulumi.OutputState }

func (SourceTriggerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SourceTriggerResponse)(nil)).Elem()
}

func (o SourceTriggerResponseArrayOutput) ToSourceTriggerResponseArrayOutput() SourceTriggerResponseArrayOutput {
	return o
}

func (o SourceTriggerResponseArrayOutput) ToSourceTriggerResponseArrayOutputWithContext(ctx context.Context) SourceTriggerResponseArrayOutput {
	return o
}

func (o SourceTriggerResponseArrayOutput) Index(i pulumi.IntInput) SourceTriggerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SourceTriggerResponse {
		return vs[0].([]SourceTriggerResponse)[vs[1].(int)]
	}).(SourceTriggerResponseOutput)
}

type TaskStepProperties struct {
	ContextAccessToken *string `pulumi:"contextAccessToken"`
	ContextPath        *string `pulumi:"contextPath"`
}





type TaskStepPropertiesInput interface {
	pulumi.Input

	ToTaskStepPropertiesOutput() TaskStepPropertiesOutput
	ToTaskStepPropertiesOutputWithContext(context.Context) TaskStepPropertiesOutput
}

type TaskStepPropertiesArgs struct {
	ContextAccessToken pulumi.StringPtrInput `pulumi:"contextAccessToken"`
	ContextPath        pulumi.StringPtrInput `pulumi:"contextPath"`
}

func (TaskStepPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskStepProperties)(nil)).Elem()
}

func (i TaskStepPropertiesArgs) ToTaskStepPropertiesOutput() TaskStepPropertiesOutput {
	return i.ToTaskStepPropertiesOutputWithContext(context.Background())
}

func (i TaskStepPropertiesArgs) ToTaskStepPropertiesOutputWithContext(ctx context.Context) TaskStepPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskStepPropertiesOutput)
}

func (i TaskStepPropertiesArgs) ToTaskStepPropertiesPtrOutput() TaskStepPropertiesPtrOutput {
	return i.ToTaskStepPropertiesPtrOutputWithContext(context.Background())
}

func (i TaskStepPropertiesArgs) ToTaskStepPropertiesPtrOutputWithContext(ctx context.Context) TaskStepPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskStepPropertiesOutput).ToTaskStepPropertiesPtrOutputWithContext(ctx)
}









type TaskStepPropertiesPtrInput interface {
	pulumi.Input

	ToTaskStepPropertiesPtrOutput() TaskStepPropertiesPtrOutput
	ToTaskStepPropertiesPtrOutputWithContext(context.Context) TaskStepPropertiesPtrOutput
}

type taskStepPropertiesPtrType TaskStepPropertiesArgs

func TaskStepPropertiesPtr(v *TaskStepPropertiesArgs) TaskStepPropertiesPtrInput {
	return (*taskStepPropertiesPtrType)(v)
}

func (*taskStepPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskStepProperties)(nil)).Elem()
}

func (i *taskStepPropertiesPtrType) ToTaskStepPropertiesPtrOutput() TaskStepPropertiesPtrOutput {
	return i.ToTaskStepPropertiesPtrOutputWithContext(context.Background())
}

func (i *taskStepPropertiesPtrType) ToTaskStepPropertiesPtrOutputWithContext(ctx context.Context) TaskStepPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskStepPropertiesPtrOutput)
}

type TaskStepPropertiesOutput struct{ *pulumi.OutputState }

func (TaskStepPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskStepProperties)(nil)).Elem()
}

func (o TaskStepPropertiesOutput) ToTaskStepPropertiesOutput() TaskStepPropertiesOutput {
	return o
}

func (o TaskStepPropertiesOutput) ToTaskStepPropertiesOutputWithContext(ctx context.Context) TaskStepPropertiesOutput {
	return o
}

func (o TaskStepPropertiesOutput) ToTaskStepPropertiesPtrOutput() TaskStepPropertiesPtrOutput {
	return o.ToTaskStepPropertiesPtrOutputWithContext(context.Background())
}

func (o TaskStepPropertiesOutput) ToTaskStepPropertiesPtrOutputWithContext(ctx context.Context) TaskStepPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TaskStepProperties) *TaskStepProperties {
		return &v
	}).(TaskStepPropertiesPtrOutput)
}

func (o TaskStepPropertiesOutput) ContextAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskStepProperties) *string { return v.ContextAccessToken }).(pulumi.StringPtrOutput)
}

func (o TaskStepPropertiesOutput) ContextPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskStepProperties) *string { return v.ContextPath }).(pulumi.StringPtrOutput)
}

type TaskStepPropertiesPtrOutput struct{ *pulumi.OutputState }

func (TaskStepPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskStepProperties)(nil)).Elem()
}

func (o TaskStepPropertiesPtrOutput) ToTaskStepPropertiesPtrOutput() TaskStepPropertiesPtrOutput {
	return o
}

func (o TaskStepPropertiesPtrOutput) ToTaskStepPropertiesPtrOutputWithContext(ctx context.Context) TaskStepPropertiesPtrOutput {
	return o
}

func (o TaskStepPropertiesPtrOutput) Elem() TaskStepPropertiesOutput {
	return o.ApplyT(func(v *TaskStepProperties) TaskStepProperties {
		if v != nil {
			return *v
		}
		var ret TaskStepProperties
		return ret
	}).(TaskStepPropertiesOutput)
}

func (o TaskStepPropertiesPtrOutput) ContextAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskStepProperties) *string {
		if v == nil {
			return nil
		}
		return v.ContextAccessToken
	}).(pulumi.StringPtrOutput)
}

func (o TaskStepPropertiesPtrOutput) ContextPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskStepProperties) *string {
		if v == nil {
			return nil
		}
		return v.ContextPath
	}).(pulumi.StringPtrOutput)
}

type TriggerProperties struct {
	BaseImageTrigger *BaseImageTrigger `pulumi:"baseImageTrigger"`
	SourceTriggers   []SourceTrigger   `pulumi:"sourceTriggers"`
}





type TriggerPropertiesInput interface {
	pulumi.Input

	ToTriggerPropertiesOutput() TriggerPropertiesOutput
	ToTriggerPropertiesOutputWithContext(context.Context) TriggerPropertiesOutput
}

type TriggerPropertiesArgs struct {
	BaseImageTrigger BaseImageTriggerPtrInput `pulumi:"baseImageTrigger"`
	SourceTriggers   SourceTriggerArrayInput  `pulumi:"sourceTriggers"`
}

func (TriggerPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerProperties)(nil)).Elem()
}

func (i TriggerPropertiesArgs) ToTriggerPropertiesOutput() TriggerPropertiesOutput {
	return i.ToTriggerPropertiesOutputWithContext(context.Background())
}

func (i TriggerPropertiesArgs) ToTriggerPropertiesOutputWithContext(ctx context.Context) TriggerPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerPropertiesOutput)
}

func (i TriggerPropertiesArgs) ToTriggerPropertiesPtrOutput() TriggerPropertiesPtrOutput {
	return i.ToTriggerPropertiesPtrOutputWithContext(context.Background())
}

func (i TriggerPropertiesArgs) ToTriggerPropertiesPtrOutputWithContext(ctx context.Context) TriggerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerPropertiesOutput).ToTriggerPropertiesPtrOutputWithContext(ctx)
}









type TriggerPropertiesPtrInput interface {
	pulumi.Input

	ToTriggerPropertiesPtrOutput() TriggerPropertiesPtrOutput
	ToTriggerPropertiesPtrOutputWithContext(context.Context) TriggerPropertiesPtrOutput
}

type triggerPropertiesPtrType TriggerPropertiesArgs

func TriggerPropertiesPtr(v *TriggerPropertiesArgs) TriggerPropertiesPtrInput {
	return (*triggerPropertiesPtrType)(v)
}

func (*triggerPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerProperties)(nil)).Elem()
}

func (i *triggerPropertiesPtrType) ToTriggerPropertiesPtrOutput() TriggerPropertiesPtrOutput {
	return i.ToTriggerPropertiesPtrOutputWithContext(context.Background())
}

func (i *triggerPropertiesPtrType) ToTriggerPropertiesPtrOutputWithContext(ctx context.Context) TriggerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerPropertiesPtrOutput)
}

type TriggerPropertiesOutput struct{ *pulumi.OutputState }

func (TriggerPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerProperties)(nil)).Elem()
}

func (o TriggerPropertiesOutput) ToTriggerPropertiesOutput() TriggerPropertiesOutput {
	return o
}

func (o TriggerPropertiesOutput) ToTriggerPropertiesOutputWithContext(ctx context.Context) TriggerPropertiesOutput {
	return o
}

func (o TriggerPropertiesOutput) ToTriggerPropertiesPtrOutput() TriggerPropertiesPtrOutput {
	return o.ToTriggerPropertiesPtrOutputWithContext(context.Background())
}

func (o TriggerPropertiesOutput) ToTriggerPropertiesPtrOutputWithContext(ctx context.Context) TriggerPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TriggerProperties) *TriggerProperties {
		return &v
	}).(TriggerPropertiesPtrOutput)
}

func (o TriggerPropertiesOutput) BaseImageTrigger() BaseImageTriggerPtrOutput {
	return o.ApplyT(func(v TriggerProperties) *BaseImageTrigger { return v.BaseImageTrigger }).(BaseImageTriggerPtrOutput)
}

func (o TriggerPropertiesOutput) SourceTriggers() SourceTriggerArrayOutput {
	return o.ApplyT(func(v TriggerProperties) []SourceTrigger { return v.SourceTriggers }).(SourceTriggerArrayOutput)
}

type TriggerPropertiesPtrOutput struct{ *pulumi.OutputState }

func (TriggerPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerProperties)(nil)).Elem()
}

func (o TriggerPropertiesPtrOutput) ToTriggerPropertiesPtrOutput() TriggerPropertiesPtrOutput {
	return o
}

func (o TriggerPropertiesPtrOutput) ToTriggerPropertiesPtrOutputWithContext(ctx context.Context) TriggerPropertiesPtrOutput {
	return o
}

func (o TriggerPropertiesPtrOutput) Elem() TriggerPropertiesOutput {
	return o.ApplyT(func(v *TriggerProperties) TriggerProperties {
		if v != nil {
			return *v
		}
		var ret TriggerProperties
		return ret
	}).(TriggerPropertiesOutput)
}

func (o TriggerPropertiesPtrOutput) BaseImageTrigger() BaseImageTriggerPtrOutput {
	return o.ApplyT(func(v *TriggerProperties) *BaseImageTrigger {
		if v == nil {
			return nil
		}
		return v.BaseImageTrigger
	}).(BaseImageTriggerPtrOutput)
}

func (o TriggerPropertiesPtrOutput) SourceTriggers() SourceTriggerArrayOutput {
	return o.ApplyT(func(v *TriggerProperties) []SourceTrigger {
		if v == nil {
			return nil
		}
		return v.SourceTriggers
	}).(SourceTriggerArrayOutput)
}

type TriggerPropertiesResponse struct {
	BaseImageTrigger *BaseImageTriggerResponse `pulumi:"baseImageTrigger"`
	SourceTriggers   []SourceTriggerResponse   `pulumi:"sourceTriggers"`
}





type TriggerPropertiesResponseInput interface {
	pulumi.Input

	ToTriggerPropertiesResponseOutput() TriggerPropertiesResponseOutput
	ToTriggerPropertiesResponseOutputWithContext(context.Context) TriggerPropertiesResponseOutput
}

type TriggerPropertiesResponseArgs struct {
	BaseImageTrigger BaseImageTriggerResponsePtrInput `pulumi:"baseImageTrigger"`
	SourceTriggers   SourceTriggerResponseArrayInput  `pulumi:"sourceTriggers"`
}

func (TriggerPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerPropertiesResponse)(nil)).Elem()
}

func (i TriggerPropertiesResponseArgs) ToTriggerPropertiesResponseOutput() TriggerPropertiesResponseOutput {
	return i.ToTriggerPropertiesResponseOutputWithContext(context.Background())
}

func (i TriggerPropertiesResponseArgs) ToTriggerPropertiesResponseOutputWithContext(ctx context.Context) TriggerPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerPropertiesResponseOutput)
}

func (i TriggerPropertiesResponseArgs) ToTriggerPropertiesResponsePtrOutput() TriggerPropertiesResponsePtrOutput {
	return i.ToTriggerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i TriggerPropertiesResponseArgs) ToTriggerPropertiesResponsePtrOutputWithContext(ctx context.Context) TriggerPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerPropertiesResponseOutput).ToTriggerPropertiesResponsePtrOutputWithContext(ctx)
}









type TriggerPropertiesResponsePtrInput interface {
	pulumi.Input

	ToTriggerPropertiesResponsePtrOutput() TriggerPropertiesResponsePtrOutput
	ToTriggerPropertiesResponsePtrOutputWithContext(context.Context) TriggerPropertiesResponsePtrOutput
}

type triggerPropertiesResponsePtrType TriggerPropertiesResponseArgs

func TriggerPropertiesResponsePtr(v *TriggerPropertiesResponseArgs) TriggerPropertiesResponsePtrInput {
	return (*triggerPropertiesResponsePtrType)(v)
}

func (*triggerPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerPropertiesResponse)(nil)).Elem()
}

func (i *triggerPropertiesResponsePtrType) ToTriggerPropertiesResponsePtrOutput() TriggerPropertiesResponsePtrOutput {
	return i.ToTriggerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *triggerPropertiesResponsePtrType) ToTriggerPropertiesResponsePtrOutputWithContext(ctx context.Context) TriggerPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerPropertiesResponsePtrOutput)
}

type TriggerPropertiesResponseOutput struct{ *pulumi.OutputState }

func (TriggerPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerPropertiesResponse)(nil)).Elem()
}

func (o TriggerPropertiesResponseOutput) ToTriggerPropertiesResponseOutput() TriggerPropertiesResponseOutput {
	return o
}

func (o TriggerPropertiesResponseOutput) ToTriggerPropertiesResponseOutputWithContext(ctx context.Context) TriggerPropertiesResponseOutput {
	return o
}

func (o TriggerPropertiesResponseOutput) ToTriggerPropertiesResponsePtrOutput() TriggerPropertiesResponsePtrOutput {
	return o.ToTriggerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o TriggerPropertiesResponseOutput) ToTriggerPropertiesResponsePtrOutputWithContext(ctx context.Context) TriggerPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TriggerPropertiesResponse) *TriggerPropertiesResponse {
		return &v
	}).(TriggerPropertiesResponsePtrOutput)
}

func (o TriggerPropertiesResponseOutput) BaseImageTrigger() BaseImageTriggerResponsePtrOutput {
	return o.ApplyT(func(v TriggerPropertiesResponse) *BaseImageTriggerResponse { return v.BaseImageTrigger }).(BaseImageTriggerResponsePtrOutput)
}

func (o TriggerPropertiesResponseOutput) SourceTriggers() SourceTriggerResponseArrayOutput {
	return o.ApplyT(func(v TriggerPropertiesResponse) []SourceTriggerResponse { return v.SourceTriggers }).(SourceTriggerResponseArrayOutput)
}

type TriggerPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (TriggerPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerPropertiesResponse)(nil)).Elem()
}

func (o TriggerPropertiesResponsePtrOutput) ToTriggerPropertiesResponsePtrOutput() TriggerPropertiesResponsePtrOutput {
	return o
}

func (o TriggerPropertiesResponsePtrOutput) ToTriggerPropertiesResponsePtrOutputWithContext(ctx context.Context) TriggerPropertiesResponsePtrOutput {
	return o
}

func (o TriggerPropertiesResponsePtrOutput) Elem() TriggerPropertiesResponseOutput {
	return o.ApplyT(func(v *TriggerPropertiesResponse) TriggerPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret TriggerPropertiesResponse
		return ret
	}).(TriggerPropertiesResponseOutput)
}

func (o TriggerPropertiesResponsePtrOutput) BaseImageTrigger() BaseImageTriggerResponsePtrOutput {
	return o.ApplyT(func(v *TriggerPropertiesResponse) *BaseImageTriggerResponse {
		if v == nil {
			return nil
		}
		return v.BaseImageTrigger
	}).(BaseImageTriggerResponsePtrOutput)
}

func (o TriggerPropertiesResponsePtrOutput) SourceTriggers() SourceTriggerResponseArrayOutput {
	return o.ApplyT(func(v *TriggerPropertiesResponse) []SourceTriggerResponse {
		if v == nil {
			return nil
		}
		return v.SourceTriggers
	}).(SourceTriggerResponseArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AgentPropertiesInput)(nil)).Elem(), AgentPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AgentPropertiesPtrInput)(nil)).Elem(), AgentPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AgentPropertiesResponseInput)(nil)).Elem(), AgentPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AgentPropertiesResponsePtrInput)(nil)).Elem(), AgentPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArgumentResponseInput)(nil)).Elem(), ArgumentResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArgumentResponseArrayInput)(nil)).Elem(), ArgumentResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthInfoInput)(nil)).Elem(), AuthInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthInfoPtrInput)(nil)).Elem(), AuthInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthInfoResponseInput)(nil)).Elem(), AuthInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthInfoResponsePtrInput)(nil)).Elem(), AuthInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BaseImageDependencyResponseInput)(nil)).Elem(), BaseImageDependencyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BaseImageDependencyResponseArrayInput)(nil)).Elem(), BaseImageDependencyResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BaseImageTriggerInput)(nil)).Elem(), BaseImageTriggerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BaseImageTriggerPtrInput)(nil)).Elem(), BaseImageTriggerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BaseImageTriggerResponseInput)(nil)).Elem(), BaseImageTriggerResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BaseImageTriggerResponsePtrInput)(nil)).Elem(), BaseImageTriggerResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CredentialsInput)(nil)).Elem(), CredentialsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CredentialsPtrInput)(nil)).Elem(), CredentialsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CredentialsResponseInput)(nil)).Elem(), CredentialsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CredentialsResponsePtrInput)(nil)).Elem(), CredentialsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomRegistryCredentialsInput)(nil)).Elem(), CustomRegistryCredentialsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomRegistryCredentialsMapInput)(nil)).Elem(), CustomRegistryCredentialsMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomRegistryCredentialsResponseInput)(nil)).Elem(), CustomRegistryCredentialsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomRegistryCredentialsResponseMapInput)(nil)).Elem(), CustomRegistryCredentialsResponseMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*DockerBuildStepResponseInput)(nil)).Elem(), DockerBuildStepResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncodedTaskStepResponseInput)(nil)).Elem(), EncodedTaskStepResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FileTaskStepResponseInput)(nil)).Elem(), FileTaskStepResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlatformPropertiesInput)(nil)).Elem(), PlatformPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlatformPropertiesPtrInput)(nil)).Elem(), PlatformPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlatformPropertiesResponseInput)(nil)).Elem(), PlatformPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlatformPropertiesResponsePtrInput)(nil)).Elem(), PlatformPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretObjectInput)(nil)).Elem(), SecretObjectArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretObjectPtrInput)(nil)).Elem(), SecretObjectArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretObjectResponseInput)(nil)).Elem(), SecretObjectResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretObjectResponsePtrInput)(nil)).Elem(), SecretObjectResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SetValueResponseInput)(nil)).Elem(), SetValueResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SetValueResponseArrayInput)(nil)).Elem(), SetValueResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePropertiesInput)(nil)).Elem(), SourcePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePropertiesResponseInput)(nil)).Elem(), SourcePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRegistryCredentialsInput)(nil)).Elem(), SourceRegistryCredentialsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRegistryCredentialsPtrInput)(nil)).Elem(), SourceRegistryCredentialsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRegistryCredentialsResponseInput)(nil)).Elem(), SourceRegistryCredentialsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRegistryCredentialsResponsePtrInput)(nil)).Elem(), SourceRegistryCredentialsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTriggerInput)(nil)).Elem(), SourceTriggerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTriggerArrayInput)(nil)).Elem(), SourceTriggerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTriggerResponseInput)(nil)).Elem(), SourceTriggerResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTriggerResponseArrayInput)(nil)).Elem(), SourceTriggerResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaskStepPropertiesInput)(nil)).Elem(), TaskStepPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaskStepPropertiesPtrInput)(nil)).Elem(), TaskStepPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerPropertiesInput)(nil)).Elem(), TriggerPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerPropertiesPtrInput)(nil)).Elem(), TriggerPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerPropertiesResponseInput)(nil)).Elem(), TriggerPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerPropertiesResponsePtrInput)(nil)).Elem(), TriggerPropertiesResponseArgs{})
	pulumi.RegisterOutputType(AgentPropertiesOutput{})
	pulumi.RegisterOutputType(AgentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(AgentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AgentPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ArgumentResponseOutput{})
	pulumi.RegisterOutputType(ArgumentResponseArrayOutput{})
	pulumi.RegisterOutputType(AuthInfoOutput{})
	pulumi.RegisterOutputType(AuthInfoPtrOutput{})
	pulumi.RegisterOutputType(AuthInfoResponseOutput{})
	pulumi.RegisterOutputType(AuthInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(BaseImageDependencyResponseOutput{})
	pulumi.RegisterOutputType(BaseImageDependencyResponseArrayOutput{})
	pulumi.RegisterOutputType(BaseImageTriggerOutput{})
	pulumi.RegisterOutputType(BaseImageTriggerPtrOutput{})
	pulumi.RegisterOutputType(BaseImageTriggerResponseOutput{})
	pulumi.RegisterOutputType(BaseImageTriggerResponsePtrOutput{})
	pulumi.RegisterOutputType(CredentialsOutput{})
	pulumi.RegisterOutputType(CredentialsPtrOutput{})
	pulumi.RegisterOutputType(CredentialsResponseOutput{})
	pulumi.RegisterOutputType(CredentialsResponsePtrOutput{})
	pulumi.RegisterOutputType(CustomRegistryCredentialsOutput{})
	pulumi.RegisterOutputType(CustomRegistryCredentialsMapOutput{})
	pulumi.RegisterOutputType(CustomRegistryCredentialsResponseOutput{})
	pulumi.RegisterOutputType(CustomRegistryCredentialsResponseMapOutput{})
	pulumi.RegisterOutputType(DockerBuildStepResponseOutput{})
	pulumi.RegisterOutputType(EncodedTaskStepResponseOutput{})
	pulumi.RegisterOutputType(FileTaskStepResponseOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SecretObjectOutput{})
	pulumi.RegisterOutputType(SecretObjectPtrOutput{})
	pulumi.RegisterOutputType(SecretObjectResponseOutput{})
	pulumi.RegisterOutputType(SecretObjectResponsePtrOutput{})
	pulumi.RegisterOutputType(SetValueResponseOutput{})
	pulumi.RegisterOutputType(SetValueResponseArrayOutput{})
	pulumi.RegisterOutputType(SourcePropertiesOutput{})
	pulumi.RegisterOutputType(SourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SourceRegistryCredentialsOutput{})
	pulumi.RegisterOutputType(SourceRegistryCredentialsPtrOutput{})
	pulumi.RegisterOutputType(SourceRegistryCredentialsResponseOutput{})
	pulumi.RegisterOutputType(SourceRegistryCredentialsResponsePtrOutput{})
	pulumi.RegisterOutputType(SourceTriggerOutput{})
	pulumi.RegisterOutputType(SourceTriggerArrayOutput{})
	pulumi.RegisterOutputType(SourceTriggerResponseOutput{})
	pulumi.RegisterOutputType(SourceTriggerResponseArrayOutput{})
	pulumi.RegisterOutputType(TaskStepPropertiesOutput{})
	pulumi.RegisterOutputType(TaskStepPropertiesPtrOutput{})
	pulumi.RegisterOutputType(TriggerPropertiesOutput{})
	pulumi.RegisterOutputType(TriggerPropertiesPtrOutput{})
	pulumi.RegisterOutputType(TriggerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TriggerPropertiesResponsePtrOutput{})
}

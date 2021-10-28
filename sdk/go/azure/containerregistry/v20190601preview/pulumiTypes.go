


package v20190601preview

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

type Argument struct {
	IsSecret *bool  `pulumi:"isSecret"`
	Name     string `pulumi:"name"`
	Value    string `pulumi:"value"`
}





type ArgumentInput interface {
	pulumi.Input

	ToArgumentOutput() ArgumentOutput
	ToArgumentOutputWithContext(context.Context) ArgumentOutput
}

type ArgumentArgs struct {
	IsSecret pulumi.BoolPtrInput `pulumi:"isSecret"`
	Name     pulumi.StringInput  `pulumi:"name"`
	Value    pulumi.StringInput  `pulumi:"value"`
}

func (ArgumentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Argument)(nil)).Elem()
}

func (i ArgumentArgs) ToArgumentOutput() ArgumentOutput {
	return i.ToArgumentOutputWithContext(context.Background())
}

func (i ArgumentArgs) ToArgumentOutputWithContext(ctx context.Context) ArgumentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArgumentOutput)
}





type ArgumentArrayInput interface {
	pulumi.Input

	ToArgumentArrayOutput() ArgumentArrayOutput
	ToArgumentArrayOutputWithContext(context.Context) ArgumentArrayOutput
}

type ArgumentArray []ArgumentInput

func (ArgumentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Argument)(nil)).Elem()
}

func (i ArgumentArray) ToArgumentArrayOutput() ArgumentArrayOutput {
	return i.ToArgumentArrayOutputWithContext(context.Background())
}

func (i ArgumentArray) ToArgumentArrayOutputWithContext(ctx context.Context) ArgumentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArgumentArrayOutput)
}

type ArgumentOutput struct{ *pulumi.OutputState }

func (ArgumentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Argument)(nil)).Elem()
}

func (o ArgumentOutput) ToArgumentOutput() ArgumentOutput {
	return o
}

func (o ArgumentOutput) ToArgumentOutputWithContext(ctx context.Context) ArgumentOutput {
	return o
}

func (o ArgumentOutput) GetIsSecret() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Argument) *bool { return v.IsSecret }).(pulumi.BoolPtrOutput)
}

func (o ArgumentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Argument) string { return v.Name }).(pulumi.StringOutput)
}

func (o ArgumentOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v Argument) string { return v.Value }).(pulumi.StringOutput)
}

type ArgumentArrayOutput struct{ *pulumi.OutputState }

func (ArgumentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Argument)(nil)).Elem()
}

func (o ArgumentArrayOutput) ToArgumentArrayOutput() ArgumentArrayOutput {
	return o
}

func (o ArgumentArrayOutput) ToArgumentArrayOutputWithContext(ctx context.Context) ArgumentArrayOutput {
	return o
}

func (o ArgumentArrayOutput) Index(i pulumi.IntInput) ArgumentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Argument {
		return vs[0].([]Argument)[vs[1].(int)]
	}).(ArgumentOutput)
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
	BaseImageTriggerType     string  `pulumi:"baseImageTriggerType"`
	Name                     string  `pulumi:"name"`
	Status                   *string `pulumi:"status"`
	UpdateTriggerEndpoint    *string `pulumi:"updateTriggerEndpoint"`
	UpdateTriggerPayloadType *string `pulumi:"updateTriggerPayloadType"`
}





type BaseImageTriggerInput interface {
	pulumi.Input

	ToBaseImageTriggerOutput() BaseImageTriggerOutput
	ToBaseImageTriggerOutputWithContext(context.Context) BaseImageTriggerOutput
}

type BaseImageTriggerArgs struct {
	BaseImageTriggerType     pulumi.StringInput    `pulumi:"baseImageTriggerType"`
	Name                     pulumi.StringInput    `pulumi:"name"`
	Status                   pulumi.StringPtrInput `pulumi:"status"`
	UpdateTriggerEndpoint    pulumi.StringPtrInput `pulumi:"updateTriggerEndpoint"`
	UpdateTriggerPayloadType pulumi.StringPtrInput `pulumi:"updateTriggerPayloadType"`
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

func (o BaseImageTriggerOutput) UpdateTriggerEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaseImageTrigger) *string { return v.UpdateTriggerEndpoint }).(pulumi.StringPtrOutput)
}

func (o BaseImageTriggerOutput) UpdateTriggerPayloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaseImageTrigger) *string { return v.UpdateTriggerPayloadType }).(pulumi.StringPtrOutput)
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

func (o BaseImageTriggerPtrOutput) UpdateTriggerEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaseImageTrigger) *string {
		if v == nil {
			return nil
		}
		return v.UpdateTriggerEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o BaseImageTriggerPtrOutput) UpdateTriggerPayloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaseImageTrigger) *string {
		if v == nil {
			return nil
		}
		return v.UpdateTriggerPayloadType
	}).(pulumi.StringPtrOutput)
}

type BaseImageTriggerResponse struct {
	BaseImageTriggerType     string  `pulumi:"baseImageTriggerType"`
	Name                     string  `pulumi:"name"`
	Status                   *string `pulumi:"status"`
	UpdateTriggerEndpoint    *string `pulumi:"updateTriggerEndpoint"`
	UpdateTriggerPayloadType *string `pulumi:"updateTriggerPayloadType"`
}





type BaseImageTriggerResponseInput interface {
	pulumi.Input

	ToBaseImageTriggerResponseOutput() BaseImageTriggerResponseOutput
	ToBaseImageTriggerResponseOutputWithContext(context.Context) BaseImageTriggerResponseOutput
}

type BaseImageTriggerResponseArgs struct {
	BaseImageTriggerType     pulumi.StringInput    `pulumi:"baseImageTriggerType"`
	Name                     pulumi.StringInput    `pulumi:"name"`
	Status                   pulumi.StringPtrInput `pulumi:"status"`
	UpdateTriggerEndpoint    pulumi.StringPtrInput `pulumi:"updateTriggerEndpoint"`
	UpdateTriggerPayloadType pulumi.StringPtrInput `pulumi:"updateTriggerPayloadType"`
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

func (o BaseImageTriggerResponseOutput) UpdateTriggerEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaseImageTriggerResponse) *string { return v.UpdateTriggerEndpoint }).(pulumi.StringPtrOutput)
}

func (o BaseImageTriggerResponseOutput) UpdateTriggerPayloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaseImageTriggerResponse) *string { return v.UpdateTriggerPayloadType }).(pulumi.StringPtrOutput)
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

func (o BaseImageTriggerResponsePtrOutput) UpdateTriggerEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaseImageTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.UpdateTriggerEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o BaseImageTriggerResponsePtrOutput) UpdateTriggerPayloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaseImageTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.UpdateTriggerPayloadType
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
	Identity *string       `pulumi:"identity"`
	Password *SecretObject `pulumi:"password"`
	UserName *SecretObject `pulumi:"userName"`
}





type CustomRegistryCredentialsInput interface {
	pulumi.Input

	ToCustomRegistryCredentialsOutput() CustomRegistryCredentialsOutput
	ToCustomRegistryCredentialsOutputWithContext(context.Context) CustomRegistryCredentialsOutput
}

type CustomRegistryCredentialsArgs struct {
	Identity pulumi.StringPtrInput `pulumi:"identity"`
	Password SecretObjectPtrInput  `pulumi:"password"`
	UserName SecretObjectPtrInput  `pulumi:"userName"`
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

func (o CustomRegistryCredentialsOutput) Identity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomRegistryCredentials) *string { return v.Identity }).(pulumi.StringPtrOutput)
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
	Identity *string               `pulumi:"identity"`
	Password *SecretObjectResponse `pulumi:"password"`
	UserName *SecretObjectResponse `pulumi:"userName"`
}





type CustomRegistryCredentialsResponseInput interface {
	pulumi.Input

	ToCustomRegistryCredentialsResponseOutput() CustomRegistryCredentialsResponseOutput
	ToCustomRegistryCredentialsResponseOutputWithContext(context.Context) CustomRegistryCredentialsResponseOutput
}

type CustomRegistryCredentialsResponseArgs struct {
	Identity pulumi.StringPtrInput        `pulumi:"identity"`
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

func (o CustomRegistryCredentialsResponseOutput) Identity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomRegistryCredentialsResponse) *string { return v.Identity }).(pulumi.StringPtrOutput)
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

type DockerBuildRequest struct {
	AgentConfiguration *AgentProperties   `pulumi:"agentConfiguration"`
	AgentPoolName      *string            `pulumi:"agentPoolName"`
	Arguments          []Argument         `pulumi:"arguments"`
	Credentials        *Credentials       `pulumi:"credentials"`
	DockerFilePath     string             `pulumi:"dockerFilePath"`
	ImageNames         []string           `pulumi:"imageNames"`
	IsArchiveEnabled   *bool              `pulumi:"isArchiveEnabled"`
	IsPushEnabled      *bool              `pulumi:"isPushEnabled"`
	LogTemplate        *string            `pulumi:"logTemplate"`
	NoCache            *bool              `pulumi:"noCache"`
	Platform           PlatformProperties `pulumi:"platform"`
	SourceLocation     *string            `pulumi:"sourceLocation"`
	Target             *string            `pulumi:"target"`
	Timeout            *int               `pulumi:"timeout"`
	Type               string             `pulumi:"type"`
}





type DockerBuildRequestInput interface {
	pulumi.Input

	ToDockerBuildRequestOutput() DockerBuildRequestOutput
	ToDockerBuildRequestOutputWithContext(context.Context) DockerBuildRequestOutput
}

type DockerBuildRequestArgs struct {
	AgentConfiguration AgentPropertiesPtrInput `pulumi:"agentConfiguration"`
	AgentPoolName      pulumi.StringPtrInput   `pulumi:"agentPoolName"`
	Arguments          ArgumentArrayInput      `pulumi:"arguments"`
	Credentials        CredentialsPtrInput     `pulumi:"credentials"`
	DockerFilePath     pulumi.StringInput      `pulumi:"dockerFilePath"`
	ImageNames         pulumi.StringArrayInput `pulumi:"imageNames"`
	IsArchiveEnabled   pulumi.BoolPtrInput     `pulumi:"isArchiveEnabled"`
	IsPushEnabled      pulumi.BoolPtrInput     `pulumi:"isPushEnabled"`
	LogTemplate        pulumi.StringPtrInput   `pulumi:"logTemplate"`
	NoCache            pulumi.BoolPtrInput     `pulumi:"noCache"`
	Platform           PlatformPropertiesInput `pulumi:"platform"`
	SourceLocation     pulumi.StringPtrInput   `pulumi:"sourceLocation"`
	Target             pulumi.StringPtrInput   `pulumi:"target"`
	Timeout            pulumi.IntPtrInput      `pulumi:"timeout"`
	Type               pulumi.StringInput      `pulumi:"type"`
}

func (DockerBuildRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DockerBuildRequest)(nil)).Elem()
}

func (i DockerBuildRequestArgs) ToDockerBuildRequestOutput() DockerBuildRequestOutput {
	return i.ToDockerBuildRequestOutputWithContext(context.Background())
}

func (i DockerBuildRequestArgs) ToDockerBuildRequestOutputWithContext(ctx context.Context) DockerBuildRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerBuildRequestOutput)
}

type DockerBuildRequestOutput struct{ *pulumi.OutputState }

func (DockerBuildRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DockerBuildRequest)(nil)).Elem()
}

func (o DockerBuildRequestOutput) ToDockerBuildRequestOutput() DockerBuildRequestOutput {
	return o
}

func (o DockerBuildRequestOutput) ToDockerBuildRequestOutputWithContext(ctx context.Context) DockerBuildRequestOutput {
	return o
}

func (o DockerBuildRequestOutput) AgentConfiguration() AgentPropertiesPtrOutput {
	return o.ApplyT(func(v DockerBuildRequest) *AgentProperties { return v.AgentConfiguration }).(AgentPropertiesPtrOutput)
}

func (o DockerBuildRequestOutput) AgentPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DockerBuildRequest) *string { return v.AgentPoolName }).(pulumi.StringPtrOutput)
}

func (o DockerBuildRequestOutput) Arguments() ArgumentArrayOutput {
	return o.ApplyT(func(v DockerBuildRequest) []Argument { return v.Arguments }).(ArgumentArrayOutput)
}

func (o DockerBuildRequestOutput) Credentials() CredentialsPtrOutput {
	return o.ApplyT(func(v DockerBuildRequest) *Credentials { return v.Credentials }).(CredentialsPtrOutput)
}

func (o DockerBuildRequestOutput) DockerFilePath() pulumi.StringOutput {
	return o.ApplyT(func(v DockerBuildRequest) string { return v.DockerFilePath }).(pulumi.StringOutput)
}

func (o DockerBuildRequestOutput) ImageNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DockerBuildRequest) []string { return v.ImageNames }).(pulumi.StringArrayOutput)
}

func (o DockerBuildRequestOutput) IsArchiveEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DockerBuildRequest) *bool { return v.IsArchiveEnabled }).(pulumi.BoolPtrOutput)
}

func (o DockerBuildRequestOutput) IsPushEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DockerBuildRequest) *bool { return v.IsPushEnabled }).(pulumi.BoolPtrOutput)
}

func (o DockerBuildRequestOutput) LogTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DockerBuildRequest) *string { return v.LogTemplate }).(pulumi.StringPtrOutput)
}

func (o DockerBuildRequestOutput) NoCache() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DockerBuildRequest) *bool { return v.NoCache }).(pulumi.BoolPtrOutput)
}

func (o DockerBuildRequestOutput) Platform() PlatformPropertiesOutput {
	return o.ApplyT(func(v DockerBuildRequest) PlatformProperties { return v.Platform }).(PlatformPropertiesOutput)
}

func (o DockerBuildRequestOutput) SourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DockerBuildRequest) *string { return v.SourceLocation }).(pulumi.StringPtrOutput)
}

func (o DockerBuildRequestOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DockerBuildRequest) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o DockerBuildRequestOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DockerBuildRequest) *int { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o DockerBuildRequestOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DockerBuildRequest) string { return v.Type }).(pulumi.StringOutput)
}

type DockerBuildRequestResponse struct {
	AgentConfiguration *AgentPropertiesResponse   `pulumi:"agentConfiguration"`
	AgentPoolName      *string                    `pulumi:"agentPoolName"`
	Arguments          []ArgumentResponse         `pulumi:"arguments"`
	Credentials        *CredentialsResponse       `pulumi:"credentials"`
	DockerFilePath     string                     `pulumi:"dockerFilePath"`
	ImageNames         []string                   `pulumi:"imageNames"`
	IsArchiveEnabled   *bool                      `pulumi:"isArchiveEnabled"`
	IsPushEnabled      *bool                      `pulumi:"isPushEnabled"`
	LogTemplate        *string                    `pulumi:"logTemplate"`
	NoCache            *bool                      `pulumi:"noCache"`
	Platform           PlatformPropertiesResponse `pulumi:"platform"`
	SourceLocation     *string                    `pulumi:"sourceLocation"`
	Target             *string                    `pulumi:"target"`
	Timeout            *int                       `pulumi:"timeout"`
	Type               string                     `pulumi:"type"`
}





type DockerBuildRequestResponseInput interface {
	pulumi.Input

	ToDockerBuildRequestResponseOutput() DockerBuildRequestResponseOutput
	ToDockerBuildRequestResponseOutputWithContext(context.Context) DockerBuildRequestResponseOutput
}

type DockerBuildRequestResponseArgs struct {
	AgentConfiguration AgentPropertiesResponsePtrInput `pulumi:"agentConfiguration"`
	AgentPoolName      pulumi.StringPtrInput           `pulumi:"agentPoolName"`
	Arguments          ArgumentResponseArrayInput      `pulumi:"arguments"`
	Credentials        CredentialsResponsePtrInput     `pulumi:"credentials"`
	DockerFilePath     pulumi.StringInput              `pulumi:"dockerFilePath"`
	ImageNames         pulumi.StringArrayInput         `pulumi:"imageNames"`
	IsArchiveEnabled   pulumi.BoolPtrInput             `pulumi:"isArchiveEnabled"`
	IsPushEnabled      pulumi.BoolPtrInput             `pulumi:"isPushEnabled"`
	LogTemplate        pulumi.StringPtrInput           `pulumi:"logTemplate"`
	NoCache            pulumi.BoolPtrInput             `pulumi:"noCache"`
	Platform           PlatformPropertiesResponseInput `pulumi:"platform"`
	SourceLocation     pulumi.StringPtrInput           `pulumi:"sourceLocation"`
	Target             pulumi.StringPtrInput           `pulumi:"target"`
	Timeout            pulumi.IntPtrInput              `pulumi:"timeout"`
	Type               pulumi.StringInput              `pulumi:"type"`
}

func (DockerBuildRequestResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DockerBuildRequestResponse)(nil)).Elem()
}

func (i DockerBuildRequestResponseArgs) ToDockerBuildRequestResponseOutput() DockerBuildRequestResponseOutput {
	return i.ToDockerBuildRequestResponseOutputWithContext(context.Background())
}

func (i DockerBuildRequestResponseArgs) ToDockerBuildRequestResponseOutputWithContext(ctx context.Context) DockerBuildRequestResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerBuildRequestResponseOutput)
}

type DockerBuildRequestResponseOutput struct{ *pulumi.OutputState }

func (DockerBuildRequestResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DockerBuildRequestResponse)(nil)).Elem()
}

func (o DockerBuildRequestResponseOutput) ToDockerBuildRequestResponseOutput() DockerBuildRequestResponseOutput {
	return o
}

func (o DockerBuildRequestResponseOutput) ToDockerBuildRequestResponseOutputWithContext(ctx context.Context) DockerBuildRequestResponseOutput {
	return o
}

func (o DockerBuildRequestResponseOutput) AgentConfiguration() AgentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v DockerBuildRequestResponse) *AgentPropertiesResponse { return v.AgentConfiguration }).(AgentPropertiesResponsePtrOutput)
}

func (o DockerBuildRequestResponseOutput) AgentPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DockerBuildRequestResponse) *string { return v.AgentPoolName }).(pulumi.StringPtrOutput)
}

func (o DockerBuildRequestResponseOutput) Arguments() ArgumentResponseArrayOutput {
	return o.ApplyT(func(v DockerBuildRequestResponse) []ArgumentResponse { return v.Arguments }).(ArgumentResponseArrayOutput)
}

func (o DockerBuildRequestResponseOutput) Credentials() CredentialsResponsePtrOutput {
	return o.ApplyT(func(v DockerBuildRequestResponse) *CredentialsResponse { return v.Credentials }).(CredentialsResponsePtrOutput)
}

func (o DockerBuildRequestResponseOutput) DockerFilePath() pulumi.StringOutput {
	return o.ApplyT(func(v DockerBuildRequestResponse) string { return v.DockerFilePath }).(pulumi.StringOutput)
}

func (o DockerBuildRequestResponseOutput) ImageNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DockerBuildRequestResponse) []string { return v.ImageNames }).(pulumi.StringArrayOutput)
}

func (o DockerBuildRequestResponseOutput) IsArchiveEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DockerBuildRequestResponse) *bool { return v.IsArchiveEnabled }).(pulumi.BoolPtrOutput)
}

func (o DockerBuildRequestResponseOutput) IsPushEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DockerBuildRequestResponse) *bool { return v.IsPushEnabled }).(pulumi.BoolPtrOutput)
}

func (o DockerBuildRequestResponseOutput) LogTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DockerBuildRequestResponse) *string { return v.LogTemplate }).(pulumi.StringPtrOutput)
}

func (o DockerBuildRequestResponseOutput) NoCache() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DockerBuildRequestResponse) *bool { return v.NoCache }).(pulumi.BoolPtrOutput)
}

func (o DockerBuildRequestResponseOutput) Platform() PlatformPropertiesResponseOutput {
	return o.ApplyT(func(v DockerBuildRequestResponse) PlatformPropertiesResponse { return v.Platform }).(PlatformPropertiesResponseOutput)
}

func (o DockerBuildRequestResponseOutput) SourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DockerBuildRequestResponse) *string { return v.SourceLocation }).(pulumi.StringPtrOutput)
}

func (o DockerBuildRequestResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DockerBuildRequestResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o DockerBuildRequestResponseOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DockerBuildRequestResponse) *int { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o DockerBuildRequestResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DockerBuildRequestResponse) string { return v.Type }).(pulumi.StringOutput)
}

type DockerBuildStep struct {
	Arguments          []Argument `pulumi:"arguments"`
	ContextAccessToken *string    `pulumi:"contextAccessToken"`
	ContextPath        *string    `pulumi:"contextPath"`
	DockerFilePath     string     `pulumi:"dockerFilePath"`
	ImageNames         []string   `pulumi:"imageNames"`
	IsPushEnabled      *bool      `pulumi:"isPushEnabled"`
	NoCache            *bool      `pulumi:"noCache"`
	Target             *string    `pulumi:"target"`
	Type               string     `pulumi:"type"`
}





type DockerBuildStepInput interface {
	pulumi.Input

	ToDockerBuildStepOutput() DockerBuildStepOutput
	ToDockerBuildStepOutputWithContext(context.Context) DockerBuildStepOutput
}

type DockerBuildStepArgs struct {
	Arguments          ArgumentArrayInput      `pulumi:"arguments"`
	ContextAccessToken pulumi.StringPtrInput   `pulumi:"contextAccessToken"`
	ContextPath        pulumi.StringPtrInput   `pulumi:"contextPath"`
	DockerFilePath     pulumi.StringInput      `pulumi:"dockerFilePath"`
	ImageNames         pulumi.StringArrayInput `pulumi:"imageNames"`
	IsPushEnabled      pulumi.BoolPtrInput     `pulumi:"isPushEnabled"`
	NoCache            pulumi.BoolPtrInput     `pulumi:"noCache"`
	Target             pulumi.StringPtrInput   `pulumi:"target"`
	Type               pulumi.StringInput      `pulumi:"type"`
}

func (DockerBuildStepArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DockerBuildStep)(nil)).Elem()
}

func (i DockerBuildStepArgs) ToDockerBuildStepOutput() DockerBuildStepOutput {
	return i.ToDockerBuildStepOutputWithContext(context.Background())
}

func (i DockerBuildStepArgs) ToDockerBuildStepOutputWithContext(ctx context.Context) DockerBuildStepOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerBuildStepOutput)
}

type DockerBuildStepOutput struct{ *pulumi.OutputState }

func (DockerBuildStepOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DockerBuildStep)(nil)).Elem()
}

func (o DockerBuildStepOutput) ToDockerBuildStepOutput() DockerBuildStepOutput {
	return o
}

func (o DockerBuildStepOutput) ToDockerBuildStepOutputWithContext(ctx context.Context) DockerBuildStepOutput {
	return o
}

func (o DockerBuildStepOutput) Arguments() ArgumentArrayOutput {
	return o.ApplyT(func(v DockerBuildStep) []Argument { return v.Arguments }).(ArgumentArrayOutput)
}

func (o DockerBuildStepOutput) ContextAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DockerBuildStep) *string { return v.ContextAccessToken }).(pulumi.StringPtrOutput)
}

func (o DockerBuildStepOutput) ContextPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DockerBuildStep) *string { return v.ContextPath }).(pulumi.StringPtrOutput)
}

func (o DockerBuildStepOutput) DockerFilePath() pulumi.StringOutput {
	return o.ApplyT(func(v DockerBuildStep) string { return v.DockerFilePath }).(pulumi.StringOutput)
}

func (o DockerBuildStepOutput) ImageNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DockerBuildStep) []string { return v.ImageNames }).(pulumi.StringArrayOutput)
}

func (o DockerBuildStepOutput) IsPushEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DockerBuildStep) *bool { return v.IsPushEnabled }).(pulumi.BoolPtrOutput)
}

func (o DockerBuildStepOutput) NoCache() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DockerBuildStep) *bool { return v.NoCache }).(pulumi.BoolPtrOutput)
}

func (o DockerBuildStepOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DockerBuildStep) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o DockerBuildStepOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DockerBuildStep) string { return v.Type }).(pulumi.StringOutput)
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

type EncodedTaskRunRequest struct {
	AgentConfiguration   *AgentProperties   `pulumi:"agentConfiguration"`
	AgentPoolName        *string            `pulumi:"agentPoolName"`
	Credentials          *Credentials       `pulumi:"credentials"`
	EncodedTaskContent   string             `pulumi:"encodedTaskContent"`
	EncodedValuesContent *string            `pulumi:"encodedValuesContent"`
	IsArchiveEnabled     *bool              `pulumi:"isArchiveEnabled"`
	LogTemplate          *string            `pulumi:"logTemplate"`
	Platform             PlatformProperties `pulumi:"platform"`
	SourceLocation       *string            `pulumi:"sourceLocation"`
	Timeout              *int               `pulumi:"timeout"`
	Type                 string             `pulumi:"type"`
	Values               []SetValue         `pulumi:"values"`
}





type EncodedTaskRunRequestInput interface {
	pulumi.Input

	ToEncodedTaskRunRequestOutput() EncodedTaskRunRequestOutput
	ToEncodedTaskRunRequestOutputWithContext(context.Context) EncodedTaskRunRequestOutput
}

type EncodedTaskRunRequestArgs struct {
	AgentConfiguration   AgentPropertiesPtrInput `pulumi:"agentConfiguration"`
	AgentPoolName        pulumi.StringPtrInput   `pulumi:"agentPoolName"`
	Credentials          CredentialsPtrInput     `pulumi:"credentials"`
	EncodedTaskContent   pulumi.StringInput      `pulumi:"encodedTaskContent"`
	EncodedValuesContent pulumi.StringPtrInput   `pulumi:"encodedValuesContent"`
	IsArchiveEnabled     pulumi.BoolPtrInput     `pulumi:"isArchiveEnabled"`
	LogTemplate          pulumi.StringPtrInput   `pulumi:"logTemplate"`
	Platform             PlatformPropertiesInput `pulumi:"platform"`
	SourceLocation       pulumi.StringPtrInput   `pulumi:"sourceLocation"`
	Timeout              pulumi.IntPtrInput      `pulumi:"timeout"`
	Type                 pulumi.StringInput      `pulumi:"type"`
	Values               SetValueArrayInput      `pulumi:"values"`
}

func (EncodedTaskRunRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncodedTaskRunRequest)(nil)).Elem()
}

func (i EncodedTaskRunRequestArgs) ToEncodedTaskRunRequestOutput() EncodedTaskRunRequestOutput {
	return i.ToEncodedTaskRunRequestOutputWithContext(context.Background())
}

func (i EncodedTaskRunRequestArgs) ToEncodedTaskRunRequestOutputWithContext(ctx context.Context) EncodedTaskRunRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncodedTaskRunRequestOutput)
}

type EncodedTaskRunRequestOutput struct{ *pulumi.OutputState }

func (EncodedTaskRunRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncodedTaskRunRequest)(nil)).Elem()
}

func (o EncodedTaskRunRequestOutput) ToEncodedTaskRunRequestOutput() EncodedTaskRunRequestOutput {
	return o
}

func (o EncodedTaskRunRequestOutput) ToEncodedTaskRunRequestOutputWithContext(ctx context.Context) EncodedTaskRunRequestOutput {
	return o
}

func (o EncodedTaskRunRequestOutput) AgentConfiguration() AgentPropertiesPtrOutput {
	return o.ApplyT(func(v EncodedTaskRunRequest) *AgentProperties { return v.AgentConfiguration }).(AgentPropertiesPtrOutput)
}

func (o EncodedTaskRunRequestOutput) AgentPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncodedTaskRunRequest) *string { return v.AgentPoolName }).(pulumi.StringPtrOutput)
}

func (o EncodedTaskRunRequestOutput) Credentials() CredentialsPtrOutput {
	return o.ApplyT(func(v EncodedTaskRunRequest) *Credentials { return v.Credentials }).(CredentialsPtrOutput)
}

func (o EncodedTaskRunRequestOutput) EncodedTaskContent() pulumi.StringOutput {
	return o.ApplyT(func(v EncodedTaskRunRequest) string { return v.EncodedTaskContent }).(pulumi.StringOutput)
}

func (o EncodedTaskRunRequestOutput) EncodedValuesContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncodedTaskRunRequest) *string { return v.EncodedValuesContent }).(pulumi.StringPtrOutput)
}

func (o EncodedTaskRunRequestOutput) IsArchiveEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EncodedTaskRunRequest) *bool { return v.IsArchiveEnabled }).(pulumi.BoolPtrOutput)
}

func (o EncodedTaskRunRequestOutput) LogTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncodedTaskRunRequest) *string { return v.LogTemplate }).(pulumi.StringPtrOutput)
}

func (o EncodedTaskRunRequestOutput) Platform() PlatformPropertiesOutput {
	return o.ApplyT(func(v EncodedTaskRunRequest) PlatformProperties { return v.Platform }).(PlatformPropertiesOutput)
}

func (o EncodedTaskRunRequestOutput) SourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncodedTaskRunRequest) *string { return v.SourceLocation }).(pulumi.StringPtrOutput)
}

func (o EncodedTaskRunRequestOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EncodedTaskRunRequest) *int { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o EncodedTaskRunRequestOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EncodedTaskRunRequest) string { return v.Type }).(pulumi.StringOutput)
}

func (o EncodedTaskRunRequestOutput) Values() SetValueArrayOutput {
	return o.ApplyT(func(v EncodedTaskRunRequest) []SetValue { return v.Values }).(SetValueArrayOutput)
}

type EncodedTaskRunRequestResponse struct {
	AgentConfiguration   *AgentPropertiesResponse   `pulumi:"agentConfiguration"`
	AgentPoolName        *string                    `pulumi:"agentPoolName"`
	Credentials          *CredentialsResponse       `pulumi:"credentials"`
	EncodedTaskContent   string                     `pulumi:"encodedTaskContent"`
	EncodedValuesContent *string                    `pulumi:"encodedValuesContent"`
	IsArchiveEnabled     *bool                      `pulumi:"isArchiveEnabled"`
	LogTemplate          *string                    `pulumi:"logTemplate"`
	Platform             PlatformPropertiesResponse `pulumi:"platform"`
	SourceLocation       *string                    `pulumi:"sourceLocation"`
	Timeout              *int                       `pulumi:"timeout"`
	Type                 string                     `pulumi:"type"`
	Values               []SetValueResponse         `pulumi:"values"`
}





type EncodedTaskRunRequestResponseInput interface {
	pulumi.Input

	ToEncodedTaskRunRequestResponseOutput() EncodedTaskRunRequestResponseOutput
	ToEncodedTaskRunRequestResponseOutputWithContext(context.Context) EncodedTaskRunRequestResponseOutput
}

type EncodedTaskRunRequestResponseArgs struct {
	AgentConfiguration   AgentPropertiesResponsePtrInput `pulumi:"agentConfiguration"`
	AgentPoolName        pulumi.StringPtrInput           `pulumi:"agentPoolName"`
	Credentials          CredentialsResponsePtrInput     `pulumi:"credentials"`
	EncodedTaskContent   pulumi.StringInput              `pulumi:"encodedTaskContent"`
	EncodedValuesContent pulumi.StringPtrInput           `pulumi:"encodedValuesContent"`
	IsArchiveEnabled     pulumi.BoolPtrInput             `pulumi:"isArchiveEnabled"`
	LogTemplate          pulumi.StringPtrInput           `pulumi:"logTemplate"`
	Platform             PlatformPropertiesResponseInput `pulumi:"platform"`
	SourceLocation       pulumi.StringPtrInput           `pulumi:"sourceLocation"`
	Timeout              pulumi.IntPtrInput              `pulumi:"timeout"`
	Type                 pulumi.StringInput              `pulumi:"type"`
	Values               SetValueResponseArrayInput      `pulumi:"values"`
}

func (EncodedTaskRunRequestResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncodedTaskRunRequestResponse)(nil)).Elem()
}

func (i EncodedTaskRunRequestResponseArgs) ToEncodedTaskRunRequestResponseOutput() EncodedTaskRunRequestResponseOutput {
	return i.ToEncodedTaskRunRequestResponseOutputWithContext(context.Background())
}

func (i EncodedTaskRunRequestResponseArgs) ToEncodedTaskRunRequestResponseOutputWithContext(ctx context.Context) EncodedTaskRunRequestResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncodedTaskRunRequestResponseOutput)
}

type EncodedTaskRunRequestResponseOutput struct{ *pulumi.OutputState }

func (EncodedTaskRunRequestResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncodedTaskRunRequestResponse)(nil)).Elem()
}

func (o EncodedTaskRunRequestResponseOutput) ToEncodedTaskRunRequestResponseOutput() EncodedTaskRunRequestResponseOutput {
	return o
}

func (o EncodedTaskRunRequestResponseOutput) ToEncodedTaskRunRequestResponseOutputWithContext(ctx context.Context) EncodedTaskRunRequestResponseOutput {
	return o
}

func (o EncodedTaskRunRequestResponseOutput) AgentConfiguration() AgentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v EncodedTaskRunRequestResponse) *AgentPropertiesResponse { return v.AgentConfiguration }).(AgentPropertiesResponsePtrOutput)
}

func (o EncodedTaskRunRequestResponseOutput) AgentPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncodedTaskRunRequestResponse) *string { return v.AgentPoolName }).(pulumi.StringPtrOutput)
}

func (o EncodedTaskRunRequestResponseOutput) Credentials() CredentialsResponsePtrOutput {
	return o.ApplyT(func(v EncodedTaskRunRequestResponse) *CredentialsResponse { return v.Credentials }).(CredentialsResponsePtrOutput)
}

func (o EncodedTaskRunRequestResponseOutput) EncodedTaskContent() pulumi.StringOutput {
	return o.ApplyT(func(v EncodedTaskRunRequestResponse) string { return v.EncodedTaskContent }).(pulumi.StringOutput)
}

func (o EncodedTaskRunRequestResponseOutput) EncodedValuesContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncodedTaskRunRequestResponse) *string { return v.EncodedValuesContent }).(pulumi.StringPtrOutput)
}

func (o EncodedTaskRunRequestResponseOutput) IsArchiveEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EncodedTaskRunRequestResponse) *bool { return v.IsArchiveEnabled }).(pulumi.BoolPtrOutput)
}

func (o EncodedTaskRunRequestResponseOutput) LogTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncodedTaskRunRequestResponse) *string { return v.LogTemplate }).(pulumi.StringPtrOutput)
}

func (o EncodedTaskRunRequestResponseOutput) Platform() PlatformPropertiesResponseOutput {
	return o.ApplyT(func(v EncodedTaskRunRequestResponse) PlatformPropertiesResponse { return v.Platform }).(PlatformPropertiesResponseOutput)
}

func (o EncodedTaskRunRequestResponseOutput) SourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncodedTaskRunRequestResponse) *string { return v.SourceLocation }).(pulumi.StringPtrOutput)
}

func (o EncodedTaskRunRequestResponseOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EncodedTaskRunRequestResponse) *int { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o EncodedTaskRunRequestResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EncodedTaskRunRequestResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o EncodedTaskRunRequestResponseOutput) Values() SetValueResponseArrayOutput {
	return o.ApplyT(func(v EncodedTaskRunRequestResponse) []SetValueResponse { return v.Values }).(SetValueResponseArrayOutput)
}

type EncodedTaskStep struct {
	ContextAccessToken   *string    `pulumi:"contextAccessToken"`
	ContextPath          *string    `pulumi:"contextPath"`
	EncodedTaskContent   string     `pulumi:"encodedTaskContent"`
	EncodedValuesContent *string    `pulumi:"encodedValuesContent"`
	Type                 string     `pulumi:"type"`
	Values               []SetValue `pulumi:"values"`
}





type EncodedTaskStepInput interface {
	pulumi.Input

	ToEncodedTaskStepOutput() EncodedTaskStepOutput
	ToEncodedTaskStepOutputWithContext(context.Context) EncodedTaskStepOutput
}

type EncodedTaskStepArgs struct {
	ContextAccessToken   pulumi.StringPtrInput `pulumi:"contextAccessToken"`
	ContextPath          pulumi.StringPtrInput `pulumi:"contextPath"`
	EncodedTaskContent   pulumi.StringInput    `pulumi:"encodedTaskContent"`
	EncodedValuesContent pulumi.StringPtrInput `pulumi:"encodedValuesContent"`
	Type                 pulumi.StringInput    `pulumi:"type"`
	Values               SetValueArrayInput    `pulumi:"values"`
}

func (EncodedTaskStepArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncodedTaskStep)(nil)).Elem()
}

func (i EncodedTaskStepArgs) ToEncodedTaskStepOutput() EncodedTaskStepOutput {
	return i.ToEncodedTaskStepOutputWithContext(context.Background())
}

func (i EncodedTaskStepArgs) ToEncodedTaskStepOutputWithContext(ctx context.Context) EncodedTaskStepOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncodedTaskStepOutput)
}

type EncodedTaskStepOutput struct{ *pulumi.OutputState }

func (EncodedTaskStepOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncodedTaskStep)(nil)).Elem()
}

func (o EncodedTaskStepOutput) ToEncodedTaskStepOutput() EncodedTaskStepOutput {
	return o
}

func (o EncodedTaskStepOutput) ToEncodedTaskStepOutputWithContext(ctx context.Context) EncodedTaskStepOutput {
	return o
}

func (o EncodedTaskStepOutput) ContextAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncodedTaskStep) *string { return v.ContextAccessToken }).(pulumi.StringPtrOutput)
}

func (o EncodedTaskStepOutput) ContextPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncodedTaskStep) *string { return v.ContextPath }).(pulumi.StringPtrOutput)
}

func (o EncodedTaskStepOutput) EncodedTaskContent() pulumi.StringOutput {
	return o.ApplyT(func(v EncodedTaskStep) string { return v.EncodedTaskContent }).(pulumi.StringOutput)
}

func (o EncodedTaskStepOutput) EncodedValuesContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncodedTaskStep) *string { return v.EncodedValuesContent }).(pulumi.StringPtrOutput)
}

func (o EncodedTaskStepOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EncodedTaskStep) string { return v.Type }).(pulumi.StringOutput)
}

func (o EncodedTaskStepOutput) Values() SetValueArrayOutput {
	return o.ApplyT(func(v EncodedTaskStep) []SetValue { return v.Values }).(SetValueArrayOutput)
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

type FileTaskRunRequest struct {
	AgentConfiguration *AgentProperties   `pulumi:"agentConfiguration"`
	AgentPoolName      *string            `pulumi:"agentPoolName"`
	Credentials        *Credentials       `pulumi:"credentials"`
	IsArchiveEnabled   *bool              `pulumi:"isArchiveEnabled"`
	LogTemplate        *string            `pulumi:"logTemplate"`
	Platform           PlatformProperties `pulumi:"platform"`
	SourceLocation     *string            `pulumi:"sourceLocation"`
	TaskFilePath       string             `pulumi:"taskFilePath"`
	Timeout            *int               `pulumi:"timeout"`
	Type               string             `pulumi:"type"`
	Values             []SetValue         `pulumi:"values"`
	ValuesFilePath     *string            `pulumi:"valuesFilePath"`
}





type FileTaskRunRequestInput interface {
	pulumi.Input

	ToFileTaskRunRequestOutput() FileTaskRunRequestOutput
	ToFileTaskRunRequestOutputWithContext(context.Context) FileTaskRunRequestOutput
}

type FileTaskRunRequestArgs struct {
	AgentConfiguration AgentPropertiesPtrInput `pulumi:"agentConfiguration"`
	AgentPoolName      pulumi.StringPtrInput   `pulumi:"agentPoolName"`
	Credentials        CredentialsPtrInput     `pulumi:"credentials"`
	IsArchiveEnabled   pulumi.BoolPtrInput     `pulumi:"isArchiveEnabled"`
	LogTemplate        pulumi.StringPtrInput   `pulumi:"logTemplate"`
	Platform           PlatformPropertiesInput `pulumi:"platform"`
	SourceLocation     pulumi.StringPtrInput   `pulumi:"sourceLocation"`
	TaskFilePath       pulumi.StringInput      `pulumi:"taskFilePath"`
	Timeout            pulumi.IntPtrInput      `pulumi:"timeout"`
	Type               pulumi.StringInput      `pulumi:"type"`
	Values             SetValueArrayInput      `pulumi:"values"`
	ValuesFilePath     pulumi.StringPtrInput   `pulumi:"valuesFilePath"`
}

func (FileTaskRunRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FileTaskRunRequest)(nil)).Elem()
}

func (i FileTaskRunRequestArgs) ToFileTaskRunRequestOutput() FileTaskRunRequestOutput {
	return i.ToFileTaskRunRequestOutputWithContext(context.Background())
}

func (i FileTaskRunRequestArgs) ToFileTaskRunRequestOutputWithContext(ctx context.Context) FileTaskRunRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileTaskRunRequestOutput)
}

type FileTaskRunRequestOutput struct{ *pulumi.OutputState }

func (FileTaskRunRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileTaskRunRequest)(nil)).Elem()
}

func (o FileTaskRunRequestOutput) ToFileTaskRunRequestOutput() FileTaskRunRequestOutput {
	return o
}

func (o FileTaskRunRequestOutput) ToFileTaskRunRequestOutputWithContext(ctx context.Context) FileTaskRunRequestOutput {
	return o
}

func (o FileTaskRunRequestOutput) AgentConfiguration() AgentPropertiesPtrOutput {
	return o.ApplyT(func(v FileTaskRunRequest) *AgentProperties { return v.AgentConfiguration }).(AgentPropertiesPtrOutput)
}

func (o FileTaskRunRequestOutput) AgentPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileTaskRunRequest) *string { return v.AgentPoolName }).(pulumi.StringPtrOutput)
}

func (o FileTaskRunRequestOutput) Credentials() CredentialsPtrOutput {
	return o.ApplyT(func(v FileTaskRunRequest) *Credentials { return v.Credentials }).(CredentialsPtrOutput)
}

func (o FileTaskRunRequestOutput) IsArchiveEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FileTaskRunRequest) *bool { return v.IsArchiveEnabled }).(pulumi.BoolPtrOutput)
}

func (o FileTaskRunRequestOutput) LogTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileTaskRunRequest) *string { return v.LogTemplate }).(pulumi.StringPtrOutput)
}

func (o FileTaskRunRequestOutput) Platform() PlatformPropertiesOutput {
	return o.ApplyT(func(v FileTaskRunRequest) PlatformProperties { return v.Platform }).(PlatformPropertiesOutput)
}

func (o FileTaskRunRequestOutput) SourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileTaskRunRequest) *string { return v.SourceLocation }).(pulumi.StringPtrOutput)
}

func (o FileTaskRunRequestOutput) TaskFilePath() pulumi.StringOutput {
	return o.ApplyT(func(v FileTaskRunRequest) string { return v.TaskFilePath }).(pulumi.StringOutput)
}

func (o FileTaskRunRequestOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FileTaskRunRequest) *int { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o FileTaskRunRequestOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v FileTaskRunRequest) string { return v.Type }).(pulumi.StringOutput)
}

func (o FileTaskRunRequestOutput) Values() SetValueArrayOutput {
	return o.ApplyT(func(v FileTaskRunRequest) []SetValue { return v.Values }).(SetValueArrayOutput)
}

func (o FileTaskRunRequestOutput) ValuesFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileTaskRunRequest) *string { return v.ValuesFilePath }).(pulumi.StringPtrOutput)
}

type FileTaskRunRequestResponse struct {
	AgentConfiguration *AgentPropertiesResponse   `pulumi:"agentConfiguration"`
	AgentPoolName      *string                    `pulumi:"agentPoolName"`
	Credentials        *CredentialsResponse       `pulumi:"credentials"`
	IsArchiveEnabled   *bool                      `pulumi:"isArchiveEnabled"`
	LogTemplate        *string                    `pulumi:"logTemplate"`
	Platform           PlatformPropertiesResponse `pulumi:"platform"`
	SourceLocation     *string                    `pulumi:"sourceLocation"`
	TaskFilePath       string                     `pulumi:"taskFilePath"`
	Timeout            *int                       `pulumi:"timeout"`
	Type               string                     `pulumi:"type"`
	Values             []SetValueResponse         `pulumi:"values"`
	ValuesFilePath     *string                    `pulumi:"valuesFilePath"`
}





type FileTaskRunRequestResponseInput interface {
	pulumi.Input

	ToFileTaskRunRequestResponseOutput() FileTaskRunRequestResponseOutput
	ToFileTaskRunRequestResponseOutputWithContext(context.Context) FileTaskRunRequestResponseOutput
}

type FileTaskRunRequestResponseArgs struct {
	AgentConfiguration AgentPropertiesResponsePtrInput `pulumi:"agentConfiguration"`
	AgentPoolName      pulumi.StringPtrInput           `pulumi:"agentPoolName"`
	Credentials        CredentialsResponsePtrInput     `pulumi:"credentials"`
	IsArchiveEnabled   pulumi.BoolPtrInput             `pulumi:"isArchiveEnabled"`
	LogTemplate        pulumi.StringPtrInput           `pulumi:"logTemplate"`
	Platform           PlatformPropertiesResponseInput `pulumi:"platform"`
	SourceLocation     pulumi.StringPtrInput           `pulumi:"sourceLocation"`
	TaskFilePath       pulumi.StringInput              `pulumi:"taskFilePath"`
	Timeout            pulumi.IntPtrInput              `pulumi:"timeout"`
	Type               pulumi.StringInput              `pulumi:"type"`
	Values             SetValueResponseArrayInput      `pulumi:"values"`
	ValuesFilePath     pulumi.StringPtrInput           `pulumi:"valuesFilePath"`
}

func (FileTaskRunRequestResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FileTaskRunRequestResponse)(nil)).Elem()
}

func (i FileTaskRunRequestResponseArgs) ToFileTaskRunRequestResponseOutput() FileTaskRunRequestResponseOutput {
	return i.ToFileTaskRunRequestResponseOutputWithContext(context.Background())
}

func (i FileTaskRunRequestResponseArgs) ToFileTaskRunRequestResponseOutputWithContext(ctx context.Context) FileTaskRunRequestResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileTaskRunRequestResponseOutput)
}

type FileTaskRunRequestResponseOutput struct{ *pulumi.OutputState }

func (FileTaskRunRequestResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileTaskRunRequestResponse)(nil)).Elem()
}

func (o FileTaskRunRequestResponseOutput) ToFileTaskRunRequestResponseOutput() FileTaskRunRequestResponseOutput {
	return o
}

func (o FileTaskRunRequestResponseOutput) ToFileTaskRunRequestResponseOutputWithContext(ctx context.Context) FileTaskRunRequestResponseOutput {
	return o
}

func (o FileTaskRunRequestResponseOutput) AgentConfiguration() AgentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v FileTaskRunRequestResponse) *AgentPropertiesResponse { return v.AgentConfiguration }).(AgentPropertiesResponsePtrOutput)
}

func (o FileTaskRunRequestResponseOutput) AgentPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileTaskRunRequestResponse) *string { return v.AgentPoolName }).(pulumi.StringPtrOutput)
}

func (o FileTaskRunRequestResponseOutput) Credentials() CredentialsResponsePtrOutput {
	return o.ApplyT(func(v FileTaskRunRequestResponse) *CredentialsResponse { return v.Credentials }).(CredentialsResponsePtrOutput)
}

func (o FileTaskRunRequestResponseOutput) IsArchiveEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FileTaskRunRequestResponse) *bool { return v.IsArchiveEnabled }).(pulumi.BoolPtrOutput)
}

func (o FileTaskRunRequestResponseOutput) LogTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileTaskRunRequestResponse) *string { return v.LogTemplate }).(pulumi.StringPtrOutput)
}

func (o FileTaskRunRequestResponseOutput) Platform() PlatformPropertiesResponseOutput {
	return o.ApplyT(func(v FileTaskRunRequestResponse) PlatformPropertiesResponse { return v.Platform }).(PlatformPropertiesResponseOutput)
}

func (o FileTaskRunRequestResponseOutput) SourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileTaskRunRequestResponse) *string { return v.SourceLocation }).(pulumi.StringPtrOutput)
}

func (o FileTaskRunRequestResponseOutput) TaskFilePath() pulumi.StringOutput {
	return o.ApplyT(func(v FileTaskRunRequestResponse) string { return v.TaskFilePath }).(pulumi.StringOutput)
}

func (o FileTaskRunRequestResponseOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FileTaskRunRequestResponse) *int { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o FileTaskRunRequestResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v FileTaskRunRequestResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o FileTaskRunRequestResponseOutput) Values() SetValueResponseArrayOutput {
	return o.ApplyT(func(v FileTaskRunRequestResponse) []SetValueResponse { return v.Values }).(SetValueResponseArrayOutput)
}

func (o FileTaskRunRequestResponseOutput) ValuesFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileTaskRunRequestResponse) *string { return v.ValuesFilePath }).(pulumi.StringPtrOutput)
}

type FileTaskStep struct {
	ContextAccessToken *string    `pulumi:"contextAccessToken"`
	ContextPath        *string    `pulumi:"contextPath"`
	TaskFilePath       string     `pulumi:"taskFilePath"`
	Type               string     `pulumi:"type"`
	Values             []SetValue `pulumi:"values"`
	ValuesFilePath     *string    `pulumi:"valuesFilePath"`
}





type FileTaskStepInput interface {
	pulumi.Input

	ToFileTaskStepOutput() FileTaskStepOutput
	ToFileTaskStepOutputWithContext(context.Context) FileTaskStepOutput
}

type FileTaskStepArgs struct {
	ContextAccessToken pulumi.StringPtrInput `pulumi:"contextAccessToken"`
	ContextPath        pulumi.StringPtrInput `pulumi:"contextPath"`
	TaskFilePath       pulumi.StringInput    `pulumi:"taskFilePath"`
	Type               pulumi.StringInput    `pulumi:"type"`
	Values             SetValueArrayInput    `pulumi:"values"`
	ValuesFilePath     pulumi.StringPtrInput `pulumi:"valuesFilePath"`
}

func (FileTaskStepArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FileTaskStep)(nil)).Elem()
}

func (i FileTaskStepArgs) ToFileTaskStepOutput() FileTaskStepOutput {
	return i.ToFileTaskStepOutputWithContext(context.Background())
}

func (i FileTaskStepArgs) ToFileTaskStepOutputWithContext(ctx context.Context) FileTaskStepOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileTaskStepOutput)
}

type FileTaskStepOutput struct{ *pulumi.OutputState }

func (FileTaskStepOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileTaskStep)(nil)).Elem()
}

func (o FileTaskStepOutput) ToFileTaskStepOutput() FileTaskStepOutput {
	return o
}

func (o FileTaskStepOutput) ToFileTaskStepOutputWithContext(ctx context.Context) FileTaskStepOutput {
	return o
}

func (o FileTaskStepOutput) ContextAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileTaskStep) *string { return v.ContextAccessToken }).(pulumi.StringPtrOutput)
}

func (o FileTaskStepOutput) ContextPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileTaskStep) *string { return v.ContextPath }).(pulumi.StringPtrOutput)
}

func (o FileTaskStepOutput) TaskFilePath() pulumi.StringOutput {
	return o.ApplyT(func(v FileTaskStep) string { return v.TaskFilePath }).(pulumi.StringOutput)
}

func (o FileTaskStepOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v FileTaskStep) string { return v.Type }).(pulumi.StringOutput)
}

func (o FileTaskStepOutput) Values() SetValueArrayOutput {
	return o.ApplyT(func(v FileTaskStep) []SetValue { return v.Values }).(SetValueArrayOutput)
}

func (o FileTaskStepOutput) ValuesFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileTaskStep) *string { return v.ValuesFilePath }).(pulumi.StringPtrOutput)
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

type IdentityProperties struct {
	PrincipalId            *string                           `pulumi:"principalId"`
	TenantId               *string                           `pulumi:"tenantId"`
	Type                   *ResourceIdentityType             `pulumi:"type"`
	UserAssignedIdentities map[string]UserIdentityProperties `pulumi:"userAssignedIdentities"`
}





type IdentityPropertiesInput interface {
	pulumi.Input

	ToIdentityPropertiesOutput() IdentityPropertiesOutput
	ToIdentityPropertiesOutputWithContext(context.Context) IdentityPropertiesOutput
}

type IdentityPropertiesArgs struct {
	PrincipalId            pulumi.StringPtrInput          `pulumi:"principalId"`
	TenantId               pulumi.StringPtrInput          `pulumi:"tenantId"`
	Type                   ResourceIdentityTypePtrInput   `pulumi:"type"`
	UserAssignedIdentities UserIdentityPropertiesMapInput `pulumi:"userAssignedIdentities"`
}

func (IdentityPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProperties)(nil)).Elem()
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesOutput() IdentityPropertiesOutput {
	return i.ToIdentityPropertiesOutputWithContext(context.Background())
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesOutputWithContext(ctx context.Context) IdentityPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesOutput)
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return i.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesOutput).ToIdentityPropertiesPtrOutputWithContext(ctx)
}









type IdentityPropertiesPtrInput interface {
	pulumi.Input

	ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput
	ToIdentityPropertiesPtrOutputWithContext(context.Context) IdentityPropertiesPtrOutput
}

type identityPropertiesPtrType IdentityPropertiesArgs

func IdentityPropertiesPtr(v *IdentityPropertiesArgs) IdentityPropertiesPtrInput {
	return (*identityPropertiesPtrType)(v)
}

func (*identityPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProperties)(nil)).Elem()
}

func (i *identityPropertiesPtrType) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return i.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i *identityPropertiesPtrType) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesPtrOutput)
}

type IdentityPropertiesOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProperties)(nil)).Elem()
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesOutput() IdentityPropertiesOutput {
	return o
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesOutputWithContext(ctx context.Context) IdentityPropertiesOutput {
	return o
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return o.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityProperties) *IdentityProperties {
		return &v
	}).(IdentityPropertiesPtrOutput)
}

func (o IdentityPropertiesOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityProperties) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityProperties) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v IdentityProperties) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

func (o IdentityPropertiesOutput) UserAssignedIdentities() UserIdentityPropertiesMapOutput {
	return o.ApplyT(func(v IdentityProperties) map[string]UserIdentityProperties { return v.UserAssignedIdentities }).(UserIdentityPropertiesMapOutput)
}

type IdentityPropertiesPtrOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProperties)(nil)).Elem()
}

func (o IdentityPropertiesPtrOutput) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return o
}

func (o IdentityPropertiesPtrOutput) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return o
}

func (o IdentityPropertiesPtrOutput) Elem() IdentityPropertiesOutput {
	return o.ApplyT(func(v *IdentityProperties) IdentityProperties {
		if v != nil {
			return *v
		}
		var ret IdentityProperties
		return ret
	}).(IdentityPropertiesOutput)
}

func (o IdentityPropertiesPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *IdentityProperties) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

func (o IdentityPropertiesPtrOutput) UserAssignedIdentities() UserIdentityPropertiesMapOutput {
	return o.ApplyT(func(v *IdentityProperties) map[string]UserIdentityProperties {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserIdentityPropertiesMapOutput)
}

type IdentityPropertiesResponse struct {
	PrincipalId            *string                                   `pulumi:"principalId"`
	TenantId               *string                                   `pulumi:"tenantId"`
	Type                   *string                                   `pulumi:"type"`
	UserAssignedIdentities map[string]UserIdentityPropertiesResponse `pulumi:"userAssignedIdentities"`
}





type IdentityPropertiesResponseInput interface {
	pulumi.Input

	ToIdentityPropertiesResponseOutput() IdentityPropertiesResponseOutput
	ToIdentityPropertiesResponseOutputWithContext(context.Context) IdentityPropertiesResponseOutput
}

type IdentityPropertiesResponseArgs struct {
	PrincipalId            pulumi.StringPtrInput                  `pulumi:"principalId"`
	TenantId               pulumi.StringPtrInput                  `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput                  `pulumi:"type"`
	UserAssignedIdentities UserIdentityPropertiesResponseMapInput `pulumi:"userAssignedIdentities"`
}

func (IdentityPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityPropertiesResponse)(nil)).Elem()
}

func (i IdentityPropertiesResponseArgs) ToIdentityPropertiesResponseOutput() IdentityPropertiesResponseOutput {
	return i.ToIdentityPropertiesResponseOutputWithContext(context.Background())
}

func (i IdentityPropertiesResponseArgs) ToIdentityPropertiesResponseOutputWithContext(ctx context.Context) IdentityPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesResponseOutput)
}

func (i IdentityPropertiesResponseArgs) ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput {
	return i.ToIdentityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i IdentityPropertiesResponseArgs) ToIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) IdentityPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesResponseOutput).ToIdentityPropertiesResponsePtrOutputWithContext(ctx)
}









type IdentityPropertiesResponsePtrInput interface {
	pulumi.Input

	ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput
	ToIdentityPropertiesResponsePtrOutputWithContext(context.Context) IdentityPropertiesResponsePtrOutput
}

type identityPropertiesResponsePtrType IdentityPropertiesResponseArgs

func IdentityPropertiesResponsePtr(v *IdentityPropertiesResponseArgs) IdentityPropertiesResponsePtrInput {
	return (*identityPropertiesResponsePtrType)(v)
}

func (*identityPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityPropertiesResponse)(nil)).Elem()
}

func (i *identityPropertiesResponsePtrType) ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput {
	return i.ToIdentityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *identityPropertiesResponsePtrType) ToIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) IdentityPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesResponsePtrOutput)
}

type IdentityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityPropertiesResponse)(nil)).Elem()
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponseOutput() IdentityPropertiesResponseOutput {
	return o
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponseOutputWithContext(ctx context.Context) IdentityPropertiesResponseOutput {
	return o
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput {
	return o.ToIdentityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) IdentityPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityPropertiesResponse) *IdentityPropertiesResponse {
		return &v
	}).(IdentityPropertiesResponsePtrOutput)
}

func (o IdentityPropertiesResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponseOutput) UserAssignedIdentities() UserIdentityPropertiesResponseMapOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) map[string]UserIdentityPropertiesResponse {
		return v.UserAssignedIdentities
	}).(UserIdentityPropertiesResponseMapOutput)
}

type IdentityPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityPropertiesResponse)(nil)).Elem()
}

func (o IdentityPropertiesResponsePtrOutput) ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput {
	return o
}

func (o IdentityPropertiesResponsePtrOutput) ToIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) IdentityPropertiesResponsePtrOutput {
	return o
}

func (o IdentityPropertiesResponsePtrOutput) Elem() IdentityPropertiesResponseOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) IdentityPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret IdentityPropertiesResponse
		return ret
	}).(IdentityPropertiesResponseOutput)
}

func (o IdentityPropertiesResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponsePtrOutput) UserAssignedIdentities() UserIdentityPropertiesResponseMapOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) map[string]UserIdentityPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserIdentityPropertiesResponseMapOutput)
}

type ImageDescriptorResponse struct {
	Digest     *string `pulumi:"digest"`
	Registry   *string `pulumi:"registry"`
	Repository *string `pulumi:"repository"`
	Tag        *string `pulumi:"tag"`
}





type ImageDescriptorResponseInput interface {
	pulumi.Input

	ToImageDescriptorResponseOutput() ImageDescriptorResponseOutput
	ToImageDescriptorResponseOutputWithContext(context.Context) ImageDescriptorResponseOutput
}

type ImageDescriptorResponseArgs struct {
	Digest     pulumi.StringPtrInput `pulumi:"digest"`
	Registry   pulumi.StringPtrInput `pulumi:"registry"`
	Repository pulumi.StringPtrInput `pulumi:"repository"`
	Tag        pulumi.StringPtrInput `pulumi:"tag"`
}

func (ImageDescriptorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageDescriptorResponse)(nil)).Elem()
}

func (i ImageDescriptorResponseArgs) ToImageDescriptorResponseOutput() ImageDescriptorResponseOutput {
	return i.ToImageDescriptorResponseOutputWithContext(context.Background())
}

func (i ImageDescriptorResponseArgs) ToImageDescriptorResponseOutputWithContext(ctx context.Context) ImageDescriptorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageDescriptorResponseOutput)
}

func (i ImageDescriptorResponseArgs) ToImageDescriptorResponsePtrOutput() ImageDescriptorResponsePtrOutput {
	return i.ToImageDescriptorResponsePtrOutputWithContext(context.Background())
}

func (i ImageDescriptorResponseArgs) ToImageDescriptorResponsePtrOutputWithContext(ctx context.Context) ImageDescriptorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageDescriptorResponseOutput).ToImageDescriptorResponsePtrOutputWithContext(ctx)
}









type ImageDescriptorResponsePtrInput interface {
	pulumi.Input

	ToImageDescriptorResponsePtrOutput() ImageDescriptorResponsePtrOutput
	ToImageDescriptorResponsePtrOutputWithContext(context.Context) ImageDescriptorResponsePtrOutput
}

type imageDescriptorResponsePtrType ImageDescriptorResponseArgs

func ImageDescriptorResponsePtr(v *ImageDescriptorResponseArgs) ImageDescriptorResponsePtrInput {
	return (*imageDescriptorResponsePtrType)(v)
}

func (*imageDescriptorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageDescriptorResponse)(nil)).Elem()
}

func (i *imageDescriptorResponsePtrType) ToImageDescriptorResponsePtrOutput() ImageDescriptorResponsePtrOutput {
	return i.ToImageDescriptorResponsePtrOutputWithContext(context.Background())
}

func (i *imageDescriptorResponsePtrType) ToImageDescriptorResponsePtrOutputWithContext(ctx context.Context) ImageDescriptorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageDescriptorResponsePtrOutput)
}





type ImageDescriptorResponseArrayInput interface {
	pulumi.Input

	ToImageDescriptorResponseArrayOutput() ImageDescriptorResponseArrayOutput
	ToImageDescriptorResponseArrayOutputWithContext(context.Context) ImageDescriptorResponseArrayOutput
}

type ImageDescriptorResponseArray []ImageDescriptorResponseInput

func (ImageDescriptorResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageDescriptorResponse)(nil)).Elem()
}

func (i ImageDescriptorResponseArray) ToImageDescriptorResponseArrayOutput() ImageDescriptorResponseArrayOutput {
	return i.ToImageDescriptorResponseArrayOutputWithContext(context.Background())
}

func (i ImageDescriptorResponseArray) ToImageDescriptorResponseArrayOutputWithContext(ctx context.Context) ImageDescriptorResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageDescriptorResponseArrayOutput)
}

type ImageDescriptorResponseOutput struct{ *pulumi.OutputState }

func (ImageDescriptorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageDescriptorResponse)(nil)).Elem()
}

func (o ImageDescriptorResponseOutput) ToImageDescriptorResponseOutput() ImageDescriptorResponseOutput {
	return o
}

func (o ImageDescriptorResponseOutput) ToImageDescriptorResponseOutputWithContext(ctx context.Context) ImageDescriptorResponseOutput {
	return o
}

func (o ImageDescriptorResponseOutput) ToImageDescriptorResponsePtrOutput() ImageDescriptorResponsePtrOutput {
	return o.ToImageDescriptorResponsePtrOutputWithContext(context.Background())
}

func (o ImageDescriptorResponseOutput) ToImageDescriptorResponsePtrOutputWithContext(ctx context.Context) ImageDescriptorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageDescriptorResponse) *ImageDescriptorResponse {
		return &v
	}).(ImageDescriptorResponsePtrOutput)
}

func (o ImageDescriptorResponseOutput) Digest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageDescriptorResponse) *string { return v.Digest }).(pulumi.StringPtrOutput)
}

func (o ImageDescriptorResponseOutput) Registry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageDescriptorResponse) *string { return v.Registry }).(pulumi.StringPtrOutput)
}

func (o ImageDescriptorResponseOutput) Repository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageDescriptorResponse) *string { return v.Repository }).(pulumi.StringPtrOutput)
}

func (o ImageDescriptorResponseOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageDescriptorResponse) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

type ImageDescriptorResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageDescriptorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageDescriptorResponse)(nil)).Elem()
}

func (o ImageDescriptorResponsePtrOutput) ToImageDescriptorResponsePtrOutput() ImageDescriptorResponsePtrOutput {
	return o
}

func (o ImageDescriptorResponsePtrOutput) ToImageDescriptorResponsePtrOutputWithContext(ctx context.Context) ImageDescriptorResponsePtrOutput {
	return o
}

func (o ImageDescriptorResponsePtrOutput) Elem() ImageDescriptorResponseOutput {
	return o.ApplyT(func(v *ImageDescriptorResponse) ImageDescriptorResponse {
		if v != nil {
			return *v
		}
		var ret ImageDescriptorResponse
		return ret
	}).(ImageDescriptorResponseOutput)
}

func (o ImageDescriptorResponsePtrOutput) Digest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageDescriptorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Digest
	}).(pulumi.StringPtrOutput)
}

func (o ImageDescriptorResponsePtrOutput) Registry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageDescriptorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Registry
	}).(pulumi.StringPtrOutput)
}

func (o ImageDescriptorResponsePtrOutput) Repository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageDescriptorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Repository
	}).(pulumi.StringPtrOutput)
}

func (o ImageDescriptorResponsePtrOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageDescriptorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tag
	}).(pulumi.StringPtrOutput)
}

type ImageDescriptorResponseArrayOutput struct{ *pulumi.OutputState }

func (ImageDescriptorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageDescriptorResponse)(nil)).Elem()
}

func (o ImageDescriptorResponseArrayOutput) ToImageDescriptorResponseArrayOutput() ImageDescriptorResponseArrayOutput {
	return o
}

func (o ImageDescriptorResponseArrayOutput) ToImageDescriptorResponseArrayOutputWithContext(ctx context.Context) ImageDescriptorResponseArrayOutput {
	return o
}

func (o ImageDescriptorResponseArrayOutput) Index(i pulumi.IntInput) ImageDescriptorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ImageDescriptorResponse {
		return vs[0].([]ImageDescriptorResponse)[vs[1].(int)]
	}).(ImageDescriptorResponseOutput)
}

type ImageUpdateTriggerResponse struct {
	Id        *string                   `pulumi:"id"`
	Images    []ImageDescriptorResponse `pulumi:"images"`
	Timestamp *string                   `pulumi:"timestamp"`
}





type ImageUpdateTriggerResponseInput interface {
	pulumi.Input

	ToImageUpdateTriggerResponseOutput() ImageUpdateTriggerResponseOutput
	ToImageUpdateTriggerResponseOutputWithContext(context.Context) ImageUpdateTriggerResponseOutput
}

type ImageUpdateTriggerResponseArgs struct {
	Id        pulumi.StringPtrInput             `pulumi:"id"`
	Images    ImageDescriptorResponseArrayInput `pulumi:"images"`
	Timestamp pulumi.StringPtrInput             `pulumi:"timestamp"`
}

func (ImageUpdateTriggerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageUpdateTriggerResponse)(nil)).Elem()
}

func (i ImageUpdateTriggerResponseArgs) ToImageUpdateTriggerResponseOutput() ImageUpdateTriggerResponseOutput {
	return i.ToImageUpdateTriggerResponseOutputWithContext(context.Background())
}

func (i ImageUpdateTriggerResponseArgs) ToImageUpdateTriggerResponseOutputWithContext(ctx context.Context) ImageUpdateTriggerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageUpdateTriggerResponseOutput)
}

func (i ImageUpdateTriggerResponseArgs) ToImageUpdateTriggerResponsePtrOutput() ImageUpdateTriggerResponsePtrOutput {
	return i.ToImageUpdateTriggerResponsePtrOutputWithContext(context.Background())
}

func (i ImageUpdateTriggerResponseArgs) ToImageUpdateTriggerResponsePtrOutputWithContext(ctx context.Context) ImageUpdateTriggerResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageUpdateTriggerResponseOutput).ToImageUpdateTriggerResponsePtrOutputWithContext(ctx)
}









type ImageUpdateTriggerResponsePtrInput interface {
	pulumi.Input

	ToImageUpdateTriggerResponsePtrOutput() ImageUpdateTriggerResponsePtrOutput
	ToImageUpdateTriggerResponsePtrOutputWithContext(context.Context) ImageUpdateTriggerResponsePtrOutput
}

type imageUpdateTriggerResponsePtrType ImageUpdateTriggerResponseArgs

func ImageUpdateTriggerResponsePtr(v *ImageUpdateTriggerResponseArgs) ImageUpdateTriggerResponsePtrInput {
	return (*imageUpdateTriggerResponsePtrType)(v)
}

func (*imageUpdateTriggerResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageUpdateTriggerResponse)(nil)).Elem()
}

func (i *imageUpdateTriggerResponsePtrType) ToImageUpdateTriggerResponsePtrOutput() ImageUpdateTriggerResponsePtrOutput {
	return i.ToImageUpdateTriggerResponsePtrOutputWithContext(context.Background())
}

func (i *imageUpdateTriggerResponsePtrType) ToImageUpdateTriggerResponsePtrOutputWithContext(ctx context.Context) ImageUpdateTriggerResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageUpdateTriggerResponsePtrOutput)
}

type ImageUpdateTriggerResponseOutput struct{ *pulumi.OutputState }

func (ImageUpdateTriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageUpdateTriggerResponse)(nil)).Elem()
}

func (o ImageUpdateTriggerResponseOutput) ToImageUpdateTriggerResponseOutput() ImageUpdateTriggerResponseOutput {
	return o
}

func (o ImageUpdateTriggerResponseOutput) ToImageUpdateTriggerResponseOutputWithContext(ctx context.Context) ImageUpdateTriggerResponseOutput {
	return o
}

func (o ImageUpdateTriggerResponseOutput) ToImageUpdateTriggerResponsePtrOutput() ImageUpdateTriggerResponsePtrOutput {
	return o.ToImageUpdateTriggerResponsePtrOutputWithContext(context.Background())
}

func (o ImageUpdateTriggerResponseOutput) ToImageUpdateTriggerResponsePtrOutputWithContext(ctx context.Context) ImageUpdateTriggerResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageUpdateTriggerResponse) *ImageUpdateTriggerResponse {
		return &v
	}).(ImageUpdateTriggerResponsePtrOutput)
}

func (o ImageUpdateTriggerResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageUpdateTriggerResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ImageUpdateTriggerResponseOutput) Images() ImageDescriptorResponseArrayOutput {
	return o.ApplyT(func(v ImageUpdateTriggerResponse) []ImageDescriptorResponse { return v.Images }).(ImageDescriptorResponseArrayOutput)
}

func (o ImageUpdateTriggerResponseOutput) Timestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageUpdateTriggerResponse) *string { return v.Timestamp }).(pulumi.StringPtrOutput)
}

type ImageUpdateTriggerResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageUpdateTriggerResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageUpdateTriggerResponse)(nil)).Elem()
}

func (o ImageUpdateTriggerResponsePtrOutput) ToImageUpdateTriggerResponsePtrOutput() ImageUpdateTriggerResponsePtrOutput {
	return o
}

func (o ImageUpdateTriggerResponsePtrOutput) ToImageUpdateTriggerResponsePtrOutputWithContext(ctx context.Context) ImageUpdateTriggerResponsePtrOutput {
	return o
}

func (o ImageUpdateTriggerResponsePtrOutput) Elem() ImageUpdateTriggerResponseOutput {
	return o.ApplyT(func(v *ImageUpdateTriggerResponse) ImageUpdateTriggerResponse {
		if v != nil {
			return *v
		}
		var ret ImageUpdateTriggerResponse
		return ret
	}).(ImageUpdateTriggerResponseOutput)
}

func (o ImageUpdateTriggerResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageUpdateTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ImageUpdateTriggerResponsePtrOutput) Images() ImageDescriptorResponseArrayOutput {
	return o.ApplyT(func(v *ImageUpdateTriggerResponse) []ImageDescriptorResponse {
		if v == nil {
			return nil
		}
		return v.Images
	}).(ImageDescriptorResponseArrayOutput)
}

func (o ImageUpdateTriggerResponsePtrOutput) Timestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageUpdateTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.Timestamp
	}).(pulumi.StringPtrOutput)
}

type OverrideTaskStepProperties struct {
	Arguments          []Argument `pulumi:"arguments"`
	ContextPath        *string    `pulumi:"contextPath"`
	File               *string    `pulumi:"file"`
	Target             *string    `pulumi:"target"`
	UpdateTriggerToken *string    `pulumi:"updateTriggerToken"`
	Values             []SetValue `pulumi:"values"`
}





type OverrideTaskStepPropertiesInput interface {
	pulumi.Input

	ToOverrideTaskStepPropertiesOutput() OverrideTaskStepPropertiesOutput
	ToOverrideTaskStepPropertiesOutputWithContext(context.Context) OverrideTaskStepPropertiesOutput
}

type OverrideTaskStepPropertiesArgs struct {
	Arguments          ArgumentArrayInput    `pulumi:"arguments"`
	ContextPath        pulumi.StringPtrInput `pulumi:"contextPath"`
	File               pulumi.StringPtrInput `pulumi:"file"`
	Target             pulumi.StringPtrInput `pulumi:"target"`
	UpdateTriggerToken pulumi.StringPtrInput `pulumi:"updateTriggerToken"`
	Values             SetValueArrayInput    `pulumi:"values"`
}

func (OverrideTaskStepPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OverrideTaskStepProperties)(nil)).Elem()
}

func (i OverrideTaskStepPropertiesArgs) ToOverrideTaskStepPropertiesOutput() OverrideTaskStepPropertiesOutput {
	return i.ToOverrideTaskStepPropertiesOutputWithContext(context.Background())
}

func (i OverrideTaskStepPropertiesArgs) ToOverrideTaskStepPropertiesOutputWithContext(ctx context.Context) OverrideTaskStepPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverrideTaskStepPropertiesOutput)
}

func (i OverrideTaskStepPropertiesArgs) ToOverrideTaskStepPropertiesPtrOutput() OverrideTaskStepPropertiesPtrOutput {
	return i.ToOverrideTaskStepPropertiesPtrOutputWithContext(context.Background())
}

func (i OverrideTaskStepPropertiesArgs) ToOverrideTaskStepPropertiesPtrOutputWithContext(ctx context.Context) OverrideTaskStepPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverrideTaskStepPropertiesOutput).ToOverrideTaskStepPropertiesPtrOutputWithContext(ctx)
}









type OverrideTaskStepPropertiesPtrInput interface {
	pulumi.Input

	ToOverrideTaskStepPropertiesPtrOutput() OverrideTaskStepPropertiesPtrOutput
	ToOverrideTaskStepPropertiesPtrOutputWithContext(context.Context) OverrideTaskStepPropertiesPtrOutput
}

type overrideTaskStepPropertiesPtrType OverrideTaskStepPropertiesArgs

func OverrideTaskStepPropertiesPtr(v *OverrideTaskStepPropertiesArgs) OverrideTaskStepPropertiesPtrInput {
	return (*overrideTaskStepPropertiesPtrType)(v)
}

func (*overrideTaskStepPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OverrideTaskStepProperties)(nil)).Elem()
}

func (i *overrideTaskStepPropertiesPtrType) ToOverrideTaskStepPropertiesPtrOutput() OverrideTaskStepPropertiesPtrOutput {
	return i.ToOverrideTaskStepPropertiesPtrOutputWithContext(context.Background())
}

func (i *overrideTaskStepPropertiesPtrType) ToOverrideTaskStepPropertiesPtrOutputWithContext(ctx context.Context) OverrideTaskStepPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverrideTaskStepPropertiesPtrOutput)
}

type OverrideTaskStepPropertiesOutput struct{ *pulumi.OutputState }

func (OverrideTaskStepPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OverrideTaskStepProperties)(nil)).Elem()
}

func (o OverrideTaskStepPropertiesOutput) ToOverrideTaskStepPropertiesOutput() OverrideTaskStepPropertiesOutput {
	return o
}

func (o OverrideTaskStepPropertiesOutput) ToOverrideTaskStepPropertiesOutputWithContext(ctx context.Context) OverrideTaskStepPropertiesOutput {
	return o
}

func (o OverrideTaskStepPropertiesOutput) ToOverrideTaskStepPropertiesPtrOutput() OverrideTaskStepPropertiesPtrOutput {
	return o.ToOverrideTaskStepPropertiesPtrOutputWithContext(context.Background())
}

func (o OverrideTaskStepPropertiesOutput) ToOverrideTaskStepPropertiesPtrOutputWithContext(ctx context.Context) OverrideTaskStepPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OverrideTaskStepProperties) *OverrideTaskStepProperties {
		return &v
	}).(OverrideTaskStepPropertiesPtrOutput)
}

func (o OverrideTaskStepPropertiesOutput) Arguments() ArgumentArrayOutput {
	return o.ApplyT(func(v OverrideTaskStepProperties) []Argument { return v.Arguments }).(ArgumentArrayOutput)
}

func (o OverrideTaskStepPropertiesOutput) ContextPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OverrideTaskStepProperties) *string { return v.ContextPath }).(pulumi.StringPtrOutput)
}

func (o OverrideTaskStepPropertiesOutput) File() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OverrideTaskStepProperties) *string { return v.File }).(pulumi.StringPtrOutput)
}

func (o OverrideTaskStepPropertiesOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OverrideTaskStepProperties) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o OverrideTaskStepPropertiesOutput) UpdateTriggerToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OverrideTaskStepProperties) *string { return v.UpdateTriggerToken }).(pulumi.StringPtrOutput)
}

func (o OverrideTaskStepPropertiesOutput) Values() SetValueArrayOutput {
	return o.ApplyT(func(v OverrideTaskStepProperties) []SetValue { return v.Values }).(SetValueArrayOutput)
}

type OverrideTaskStepPropertiesPtrOutput struct{ *pulumi.OutputState }

func (OverrideTaskStepPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OverrideTaskStepProperties)(nil)).Elem()
}

func (o OverrideTaskStepPropertiesPtrOutput) ToOverrideTaskStepPropertiesPtrOutput() OverrideTaskStepPropertiesPtrOutput {
	return o
}

func (o OverrideTaskStepPropertiesPtrOutput) ToOverrideTaskStepPropertiesPtrOutputWithContext(ctx context.Context) OverrideTaskStepPropertiesPtrOutput {
	return o
}

func (o OverrideTaskStepPropertiesPtrOutput) Elem() OverrideTaskStepPropertiesOutput {
	return o.ApplyT(func(v *OverrideTaskStepProperties) OverrideTaskStepProperties {
		if v != nil {
			return *v
		}
		var ret OverrideTaskStepProperties
		return ret
	}).(OverrideTaskStepPropertiesOutput)
}

func (o OverrideTaskStepPropertiesPtrOutput) Arguments() ArgumentArrayOutput {
	return o.ApplyT(func(v *OverrideTaskStepProperties) []Argument {
		if v == nil {
			return nil
		}
		return v.Arguments
	}).(ArgumentArrayOutput)
}

func (o OverrideTaskStepPropertiesPtrOutput) ContextPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OverrideTaskStepProperties) *string {
		if v == nil {
			return nil
		}
		return v.ContextPath
	}).(pulumi.StringPtrOutput)
}

func (o OverrideTaskStepPropertiesPtrOutput) File() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OverrideTaskStepProperties) *string {
		if v == nil {
			return nil
		}
		return v.File
	}).(pulumi.StringPtrOutput)
}

func (o OverrideTaskStepPropertiesPtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OverrideTaskStepProperties) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

func (o OverrideTaskStepPropertiesPtrOutput) UpdateTriggerToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OverrideTaskStepProperties) *string {
		if v == nil {
			return nil
		}
		return v.UpdateTriggerToken
	}).(pulumi.StringPtrOutput)
}

func (o OverrideTaskStepPropertiesPtrOutput) Values() SetValueArrayOutput {
	return o.ApplyT(func(v *OverrideTaskStepProperties) []SetValue {
		if v == nil {
			return nil
		}
		return v.Values
	}).(SetValueArrayOutput)
}

type OverrideTaskStepPropertiesResponse struct {
	Arguments          []ArgumentResponse `pulumi:"arguments"`
	ContextPath        *string            `pulumi:"contextPath"`
	File               *string            `pulumi:"file"`
	Target             *string            `pulumi:"target"`
	UpdateTriggerToken *string            `pulumi:"updateTriggerToken"`
	Values             []SetValueResponse `pulumi:"values"`
}





type OverrideTaskStepPropertiesResponseInput interface {
	pulumi.Input

	ToOverrideTaskStepPropertiesResponseOutput() OverrideTaskStepPropertiesResponseOutput
	ToOverrideTaskStepPropertiesResponseOutputWithContext(context.Context) OverrideTaskStepPropertiesResponseOutput
}

type OverrideTaskStepPropertiesResponseArgs struct {
	Arguments          ArgumentResponseArrayInput `pulumi:"arguments"`
	ContextPath        pulumi.StringPtrInput      `pulumi:"contextPath"`
	File               pulumi.StringPtrInput      `pulumi:"file"`
	Target             pulumi.StringPtrInput      `pulumi:"target"`
	UpdateTriggerToken pulumi.StringPtrInput      `pulumi:"updateTriggerToken"`
	Values             SetValueResponseArrayInput `pulumi:"values"`
}

func (OverrideTaskStepPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OverrideTaskStepPropertiesResponse)(nil)).Elem()
}

func (i OverrideTaskStepPropertiesResponseArgs) ToOverrideTaskStepPropertiesResponseOutput() OverrideTaskStepPropertiesResponseOutput {
	return i.ToOverrideTaskStepPropertiesResponseOutputWithContext(context.Background())
}

func (i OverrideTaskStepPropertiesResponseArgs) ToOverrideTaskStepPropertiesResponseOutputWithContext(ctx context.Context) OverrideTaskStepPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverrideTaskStepPropertiesResponseOutput)
}

func (i OverrideTaskStepPropertiesResponseArgs) ToOverrideTaskStepPropertiesResponsePtrOutput() OverrideTaskStepPropertiesResponsePtrOutput {
	return i.ToOverrideTaskStepPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i OverrideTaskStepPropertiesResponseArgs) ToOverrideTaskStepPropertiesResponsePtrOutputWithContext(ctx context.Context) OverrideTaskStepPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverrideTaskStepPropertiesResponseOutput).ToOverrideTaskStepPropertiesResponsePtrOutputWithContext(ctx)
}









type OverrideTaskStepPropertiesResponsePtrInput interface {
	pulumi.Input

	ToOverrideTaskStepPropertiesResponsePtrOutput() OverrideTaskStepPropertiesResponsePtrOutput
	ToOverrideTaskStepPropertiesResponsePtrOutputWithContext(context.Context) OverrideTaskStepPropertiesResponsePtrOutput
}

type overrideTaskStepPropertiesResponsePtrType OverrideTaskStepPropertiesResponseArgs

func OverrideTaskStepPropertiesResponsePtr(v *OverrideTaskStepPropertiesResponseArgs) OverrideTaskStepPropertiesResponsePtrInput {
	return (*overrideTaskStepPropertiesResponsePtrType)(v)
}

func (*overrideTaskStepPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OverrideTaskStepPropertiesResponse)(nil)).Elem()
}

func (i *overrideTaskStepPropertiesResponsePtrType) ToOverrideTaskStepPropertiesResponsePtrOutput() OverrideTaskStepPropertiesResponsePtrOutput {
	return i.ToOverrideTaskStepPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *overrideTaskStepPropertiesResponsePtrType) ToOverrideTaskStepPropertiesResponsePtrOutputWithContext(ctx context.Context) OverrideTaskStepPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverrideTaskStepPropertiesResponsePtrOutput)
}

type OverrideTaskStepPropertiesResponseOutput struct{ *pulumi.OutputState }

func (OverrideTaskStepPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OverrideTaskStepPropertiesResponse)(nil)).Elem()
}

func (o OverrideTaskStepPropertiesResponseOutput) ToOverrideTaskStepPropertiesResponseOutput() OverrideTaskStepPropertiesResponseOutput {
	return o
}

func (o OverrideTaskStepPropertiesResponseOutput) ToOverrideTaskStepPropertiesResponseOutputWithContext(ctx context.Context) OverrideTaskStepPropertiesResponseOutput {
	return o
}

func (o OverrideTaskStepPropertiesResponseOutput) ToOverrideTaskStepPropertiesResponsePtrOutput() OverrideTaskStepPropertiesResponsePtrOutput {
	return o.ToOverrideTaskStepPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o OverrideTaskStepPropertiesResponseOutput) ToOverrideTaskStepPropertiesResponsePtrOutputWithContext(ctx context.Context) OverrideTaskStepPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OverrideTaskStepPropertiesResponse) *OverrideTaskStepPropertiesResponse {
		return &v
	}).(OverrideTaskStepPropertiesResponsePtrOutput)
}

func (o OverrideTaskStepPropertiesResponseOutput) Arguments() ArgumentResponseArrayOutput {
	return o.ApplyT(func(v OverrideTaskStepPropertiesResponse) []ArgumentResponse { return v.Arguments }).(ArgumentResponseArrayOutput)
}

func (o OverrideTaskStepPropertiesResponseOutput) ContextPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OverrideTaskStepPropertiesResponse) *string { return v.ContextPath }).(pulumi.StringPtrOutput)
}

func (o OverrideTaskStepPropertiesResponseOutput) File() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OverrideTaskStepPropertiesResponse) *string { return v.File }).(pulumi.StringPtrOutput)
}

func (o OverrideTaskStepPropertiesResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OverrideTaskStepPropertiesResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o OverrideTaskStepPropertiesResponseOutput) UpdateTriggerToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OverrideTaskStepPropertiesResponse) *string { return v.UpdateTriggerToken }).(pulumi.StringPtrOutput)
}

func (o OverrideTaskStepPropertiesResponseOutput) Values() SetValueResponseArrayOutput {
	return o.ApplyT(func(v OverrideTaskStepPropertiesResponse) []SetValueResponse { return v.Values }).(SetValueResponseArrayOutput)
}

type OverrideTaskStepPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (OverrideTaskStepPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OverrideTaskStepPropertiesResponse)(nil)).Elem()
}

func (o OverrideTaskStepPropertiesResponsePtrOutput) ToOverrideTaskStepPropertiesResponsePtrOutput() OverrideTaskStepPropertiesResponsePtrOutput {
	return o
}

func (o OverrideTaskStepPropertiesResponsePtrOutput) ToOverrideTaskStepPropertiesResponsePtrOutputWithContext(ctx context.Context) OverrideTaskStepPropertiesResponsePtrOutput {
	return o
}

func (o OverrideTaskStepPropertiesResponsePtrOutput) Elem() OverrideTaskStepPropertiesResponseOutput {
	return o.ApplyT(func(v *OverrideTaskStepPropertiesResponse) OverrideTaskStepPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret OverrideTaskStepPropertiesResponse
		return ret
	}).(OverrideTaskStepPropertiesResponseOutput)
}

func (o OverrideTaskStepPropertiesResponsePtrOutput) Arguments() ArgumentResponseArrayOutput {
	return o.ApplyT(func(v *OverrideTaskStepPropertiesResponse) []ArgumentResponse {
		if v == nil {
			return nil
		}
		return v.Arguments
	}).(ArgumentResponseArrayOutput)
}

func (o OverrideTaskStepPropertiesResponsePtrOutput) ContextPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OverrideTaskStepPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContextPath
	}).(pulumi.StringPtrOutput)
}

func (o OverrideTaskStepPropertiesResponsePtrOutput) File() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OverrideTaskStepPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.File
	}).(pulumi.StringPtrOutput)
}

func (o OverrideTaskStepPropertiesResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OverrideTaskStepPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

func (o OverrideTaskStepPropertiesResponsePtrOutput) UpdateTriggerToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OverrideTaskStepPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.UpdateTriggerToken
	}).(pulumi.StringPtrOutput)
}

func (o OverrideTaskStepPropertiesResponsePtrOutput) Values() SetValueResponseArrayOutput {
	return o.ApplyT(func(v *OverrideTaskStepPropertiesResponse) []SetValueResponse {
		if v == nil {
			return nil
		}
		return v.Values
	}).(SetValueResponseArrayOutput)
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

type RunResponse struct {
	AgentConfiguration *AgentPropertiesResponse         `pulumi:"agentConfiguration"`
	AgentPoolName      *string                          `pulumi:"agentPoolName"`
	CreateTime         *string                          `pulumi:"createTime"`
	CustomRegistries   []string                         `pulumi:"customRegistries"`
	FinishTime         *string                          `pulumi:"finishTime"`
	Id                 string                           `pulumi:"id"`
	ImageUpdateTrigger *ImageUpdateTriggerResponse      `pulumi:"imageUpdateTrigger"`
	IsArchiveEnabled   *bool                            `pulumi:"isArchiveEnabled"`
	LastUpdatedTime    *string                          `pulumi:"lastUpdatedTime"`
	LogArtifact        ImageDescriptorResponse          `pulumi:"logArtifact"`
	Name               string                           `pulumi:"name"`
	OutputImages       []ImageDescriptorResponse        `pulumi:"outputImages"`
	Platform           *PlatformPropertiesResponse      `pulumi:"platform"`
	ProvisioningState  *string                          `pulumi:"provisioningState"`
	RunErrorMessage    string                           `pulumi:"runErrorMessage"`
	RunId              *string                          `pulumi:"runId"`
	RunType            *string                          `pulumi:"runType"`
	SourceRegistryAuth *string                          `pulumi:"sourceRegistryAuth"`
	SourceTrigger      *SourceTriggerDescriptorResponse `pulumi:"sourceTrigger"`
	StartTime          *string                          `pulumi:"startTime"`
	Status             *string                          `pulumi:"status"`
	SystemData         SystemDataResponse               `pulumi:"systemData"`
	Task               *string                          `pulumi:"task"`
	TimerTrigger       *TimerTriggerDescriptorResponse  `pulumi:"timerTrigger"`
	Type               string                           `pulumi:"type"`
	UpdateTriggerToken *string                          `pulumi:"updateTriggerToken"`
}





type RunResponseInput interface {
	pulumi.Input

	ToRunResponseOutput() RunResponseOutput
	ToRunResponseOutputWithContext(context.Context) RunResponseOutput
}

type RunResponseArgs struct {
	AgentConfiguration AgentPropertiesResponsePtrInput         `pulumi:"agentConfiguration"`
	AgentPoolName      pulumi.StringPtrInput                   `pulumi:"agentPoolName"`
	CreateTime         pulumi.StringPtrInput                   `pulumi:"createTime"`
	CustomRegistries   pulumi.StringArrayInput                 `pulumi:"customRegistries"`
	FinishTime         pulumi.StringPtrInput                   `pulumi:"finishTime"`
	Id                 pulumi.StringInput                      `pulumi:"id"`
	ImageUpdateTrigger ImageUpdateTriggerResponsePtrInput      `pulumi:"imageUpdateTrigger"`
	IsArchiveEnabled   pulumi.BoolPtrInput                     `pulumi:"isArchiveEnabled"`
	LastUpdatedTime    pulumi.StringPtrInput                   `pulumi:"lastUpdatedTime"`
	LogArtifact        ImageDescriptorResponseInput            `pulumi:"logArtifact"`
	Name               pulumi.StringInput                      `pulumi:"name"`
	OutputImages       ImageDescriptorResponseArrayInput       `pulumi:"outputImages"`
	Platform           PlatformPropertiesResponsePtrInput      `pulumi:"platform"`
	ProvisioningState  pulumi.StringPtrInput                   `pulumi:"provisioningState"`
	RunErrorMessage    pulumi.StringInput                      `pulumi:"runErrorMessage"`
	RunId              pulumi.StringPtrInput                   `pulumi:"runId"`
	RunType            pulumi.StringPtrInput                   `pulumi:"runType"`
	SourceRegistryAuth pulumi.StringPtrInput                   `pulumi:"sourceRegistryAuth"`
	SourceTrigger      SourceTriggerDescriptorResponsePtrInput `pulumi:"sourceTrigger"`
	StartTime          pulumi.StringPtrInput                   `pulumi:"startTime"`
	Status             pulumi.StringPtrInput                   `pulumi:"status"`
	SystemData         SystemDataResponseInput                 `pulumi:"systemData"`
	Task               pulumi.StringPtrInput                   `pulumi:"task"`
	TimerTrigger       TimerTriggerDescriptorResponsePtrInput  `pulumi:"timerTrigger"`
	Type               pulumi.StringInput                      `pulumi:"type"`
	UpdateTriggerToken pulumi.StringPtrInput                   `pulumi:"updateTriggerToken"`
}

func (RunResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunResponse)(nil)).Elem()
}

func (i RunResponseArgs) ToRunResponseOutput() RunResponseOutput {
	return i.ToRunResponseOutputWithContext(context.Background())
}

func (i RunResponseArgs) ToRunResponseOutputWithContext(ctx context.Context) RunResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunResponseOutput)
}

func (i RunResponseArgs) ToRunResponsePtrOutput() RunResponsePtrOutput {
	return i.ToRunResponsePtrOutputWithContext(context.Background())
}

func (i RunResponseArgs) ToRunResponsePtrOutputWithContext(ctx context.Context) RunResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunResponseOutput).ToRunResponsePtrOutputWithContext(ctx)
}









type RunResponsePtrInput interface {
	pulumi.Input

	ToRunResponsePtrOutput() RunResponsePtrOutput
	ToRunResponsePtrOutputWithContext(context.Context) RunResponsePtrOutput
}

type runResponsePtrType RunResponseArgs

func RunResponsePtr(v *RunResponseArgs) RunResponsePtrInput {
	return (*runResponsePtrType)(v)
}

func (*runResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RunResponse)(nil)).Elem()
}

func (i *runResponsePtrType) ToRunResponsePtrOutput() RunResponsePtrOutput {
	return i.ToRunResponsePtrOutputWithContext(context.Background())
}

func (i *runResponsePtrType) ToRunResponsePtrOutputWithContext(ctx context.Context) RunResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunResponsePtrOutput)
}

type RunResponseOutput struct{ *pulumi.OutputState }

func (RunResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunResponse)(nil)).Elem()
}

func (o RunResponseOutput) ToRunResponseOutput() RunResponseOutput {
	return o
}

func (o RunResponseOutput) ToRunResponseOutputWithContext(ctx context.Context) RunResponseOutput {
	return o
}

func (o RunResponseOutput) ToRunResponsePtrOutput() RunResponsePtrOutput {
	return o.ToRunResponsePtrOutputWithContext(context.Background())
}

func (o RunResponseOutput) ToRunResponsePtrOutputWithContext(ctx context.Context) RunResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RunResponse) *RunResponse {
		return &v
	}).(RunResponsePtrOutput)
}

func (o RunResponseOutput) AgentConfiguration() AgentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v RunResponse) *AgentPropertiesResponse { return v.AgentConfiguration }).(AgentPropertiesResponsePtrOutput)
}

func (o RunResponseOutput) AgentPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunResponse) *string { return v.AgentPoolName }).(pulumi.StringPtrOutput)
}

func (o RunResponseOutput) CreateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunResponse) *string { return v.CreateTime }).(pulumi.StringPtrOutput)
}

func (o RunResponseOutput) CustomRegistries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RunResponse) []string { return v.CustomRegistries }).(pulumi.StringArrayOutput)
}

func (o RunResponseOutput) FinishTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunResponse) *string { return v.FinishTime }).(pulumi.StringPtrOutput)
}

func (o RunResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RunResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o RunResponseOutput) ImageUpdateTrigger() ImageUpdateTriggerResponsePtrOutput {
	return o.ApplyT(func(v RunResponse) *ImageUpdateTriggerResponse { return v.ImageUpdateTrigger }).(ImageUpdateTriggerResponsePtrOutput)
}

func (o RunResponseOutput) IsArchiveEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RunResponse) *bool { return v.IsArchiveEnabled }).(pulumi.BoolPtrOutput)
}

func (o RunResponseOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunResponse) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

func (o RunResponseOutput) LogArtifact() ImageDescriptorResponseOutput {
	return o.ApplyT(func(v RunResponse) ImageDescriptorResponse { return v.LogArtifact }).(ImageDescriptorResponseOutput)
}

func (o RunResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RunResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RunResponseOutput) OutputImages() ImageDescriptorResponseArrayOutput {
	return o.ApplyT(func(v RunResponse) []ImageDescriptorResponse { return v.OutputImages }).(ImageDescriptorResponseArrayOutput)
}

func (o RunResponseOutput) Platform() PlatformPropertiesResponsePtrOutput {
	return o.ApplyT(func(v RunResponse) *PlatformPropertiesResponse { return v.Platform }).(PlatformPropertiesResponsePtrOutput)
}

func (o RunResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o RunResponseOutput) RunErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v RunResponse) string { return v.RunErrorMessage }).(pulumi.StringOutput)
}

func (o RunResponseOutput) RunId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunResponse) *string { return v.RunId }).(pulumi.StringPtrOutput)
}

func (o RunResponseOutput) RunType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunResponse) *string { return v.RunType }).(pulumi.StringPtrOutput)
}

func (o RunResponseOutput) SourceRegistryAuth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunResponse) *string { return v.SourceRegistryAuth }).(pulumi.StringPtrOutput)
}

func (o RunResponseOutput) SourceTrigger() SourceTriggerDescriptorResponsePtrOutput {
	return o.ApplyT(func(v RunResponse) *SourceTriggerDescriptorResponse { return v.SourceTrigger }).(SourceTriggerDescriptorResponsePtrOutput)
}

func (o RunResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o RunResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o RunResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v RunResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o RunResponseOutput) Task() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunResponse) *string { return v.Task }).(pulumi.StringPtrOutput)
}

func (o RunResponseOutput) TimerTrigger() TimerTriggerDescriptorResponsePtrOutput {
	return o.ApplyT(func(v RunResponse) *TimerTriggerDescriptorResponse { return v.TimerTrigger }).(TimerTriggerDescriptorResponsePtrOutput)
}

func (o RunResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RunResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o RunResponseOutput) UpdateTriggerToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunResponse) *string { return v.UpdateTriggerToken }).(pulumi.StringPtrOutput)
}

type RunResponsePtrOutput struct{ *pulumi.OutputState }

func (RunResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunResponse)(nil)).Elem()
}

func (o RunResponsePtrOutput) ToRunResponsePtrOutput() RunResponsePtrOutput {
	return o
}

func (o RunResponsePtrOutput) ToRunResponsePtrOutputWithContext(ctx context.Context) RunResponsePtrOutput {
	return o
}

func (o RunResponsePtrOutput) Elem() RunResponseOutput {
	return o.ApplyT(func(v *RunResponse) RunResponse {
		if v != nil {
			return *v
		}
		var ret RunResponse
		return ret
	}).(RunResponseOutput)
}

func (o RunResponsePtrOutput) AgentConfiguration() AgentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *RunResponse) *AgentPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.AgentConfiguration
	}).(AgentPropertiesResponsePtrOutput)
}

func (o RunResponsePtrOutput) AgentPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunResponse) *string {
		if v == nil {
			return nil
		}
		return v.AgentPoolName
	}).(pulumi.StringPtrOutput)
}

func (o RunResponsePtrOutput) CreateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreateTime
	}).(pulumi.StringPtrOutput)
}

func (o RunResponsePtrOutput) CustomRegistries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RunResponse) []string {
		if v == nil {
			return nil
		}
		return v.CustomRegistries
	}).(pulumi.StringArrayOutput)
}

func (o RunResponsePtrOutput) FinishTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunResponse) *string {
		if v == nil {
			return nil
		}
		return v.FinishTime
	}).(pulumi.StringPtrOutput)
}

func (o RunResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o RunResponsePtrOutput) ImageUpdateTrigger() ImageUpdateTriggerResponsePtrOutput {
	return o.ApplyT(func(v *RunResponse) *ImageUpdateTriggerResponse {
		if v == nil {
			return nil
		}
		return v.ImageUpdateTrigger
	}).(ImageUpdateTriggerResponsePtrOutput)
}

func (o RunResponsePtrOutput) IsArchiveEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RunResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsArchiveEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o RunResponsePtrOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastUpdatedTime
	}).(pulumi.StringPtrOutput)
}

func (o RunResponsePtrOutput) LogArtifact() ImageDescriptorResponsePtrOutput {
	return o.ApplyT(func(v *RunResponse) *ImageDescriptorResponse {
		if v == nil {
			return nil
		}
		return &v.LogArtifact
	}).(ImageDescriptorResponsePtrOutput)
}

func (o RunResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o RunResponsePtrOutput) OutputImages() ImageDescriptorResponseArrayOutput {
	return o.ApplyT(func(v *RunResponse) []ImageDescriptorResponse {
		if v == nil {
			return nil
		}
		return v.OutputImages
	}).(ImageDescriptorResponseArrayOutput)
}

func (o RunResponsePtrOutput) Platform() PlatformPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *RunResponse) *PlatformPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Platform
	}).(PlatformPropertiesResponsePtrOutput)
}

func (o RunResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o RunResponsePtrOutput) RunErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RunErrorMessage
	}).(pulumi.StringPtrOutput)
}

func (o RunResponsePtrOutput) RunId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunResponse) *string {
		if v == nil {
			return nil
		}
		return v.RunId
	}).(pulumi.StringPtrOutput)
}

func (o RunResponsePtrOutput) RunType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunResponse) *string {
		if v == nil {
			return nil
		}
		return v.RunType
	}).(pulumi.StringPtrOutput)
}

func (o RunResponsePtrOutput) SourceRegistryAuth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunResponse) *string {
		if v == nil {
			return nil
		}
		return v.SourceRegistryAuth
	}).(pulumi.StringPtrOutput)
}

func (o RunResponsePtrOutput) SourceTrigger() SourceTriggerDescriptorResponsePtrOutput {
	return o.ApplyT(func(v *RunResponse) *SourceTriggerDescriptorResponse {
		if v == nil {
			return nil
		}
		return v.SourceTrigger
	}).(SourceTriggerDescriptorResponsePtrOutput)
}

func (o RunResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o RunResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o RunResponsePtrOutput) SystemData() SystemDataResponsePtrOutput {
	return o.ApplyT(func(v *RunResponse) *SystemDataResponse {
		if v == nil {
			return nil
		}
		return &v.SystemData
	}).(SystemDataResponsePtrOutput)
}

func (o RunResponsePtrOutput) Task() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunResponse) *string {
		if v == nil {
			return nil
		}
		return v.Task
	}).(pulumi.StringPtrOutput)
}

func (o RunResponsePtrOutput) TimerTrigger() TimerTriggerDescriptorResponsePtrOutput {
	return o.ApplyT(func(v *RunResponse) *TimerTriggerDescriptorResponse {
		if v == nil {
			return nil
		}
		return v.TimerTrigger
	}).(TimerTriggerDescriptorResponsePtrOutput)
}

func (o RunResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o RunResponsePtrOutput) UpdateTriggerToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunResponse) *string {
		if v == nil {
			return nil
		}
		return v.UpdateTriggerToken
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

type SetValue struct {
	IsSecret *bool  `pulumi:"isSecret"`
	Name     string `pulumi:"name"`
	Value    string `pulumi:"value"`
}





type SetValueInput interface {
	pulumi.Input

	ToSetValueOutput() SetValueOutput
	ToSetValueOutputWithContext(context.Context) SetValueOutput
}

type SetValueArgs struct {
	IsSecret pulumi.BoolPtrInput `pulumi:"isSecret"`
	Name     pulumi.StringInput  `pulumi:"name"`
	Value    pulumi.StringInput  `pulumi:"value"`
}

func (SetValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SetValue)(nil)).Elem()
}

func (i SetValueArgs) ToSetValueOutput() SetValueOutput {
	return i.ToSetValueOutputWithContext(context.Background())
}

func (i SetValueArgs) ToSetValueOutputWithContext(ctx context.Context) SetValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SetValueOutput)
}





type SetValueArrayInput interface {
	pulumi.Input

	ToSetValueArrayOutput() SetValueArrayOutput
	ToSetValueArrayOutputWithContext(context.Context) SetValueArrayOutput
}

type SetValueArray []SetValueInput

func (SetValueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SetValue)(nil)).Elem()
}

func (i SetValueArray) ToSetValueArrayOutput() SetValueArrayOutput {
	return i.ToSetValueArrayOutputWithContext(context.Background())
}

func (i SetValueArray) ToSetValueArrayOutputWithContext(ctx context.Context) SetValueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SetValueArrayOutput)
}

type SetValueOutput struct{ *pulumi.OutputState }

func (SetValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SetValue)(nil)).Elem()
}

func (o SetValueOutput) ToSetValueOutput() SetValueOutput {
	return o
}

func (o SetValueOutput) ToSetValueOutputWithContext(ctx context.Context) SetValueOutput {
	return o
}

func (o SetValueOutput) GetIsSecret() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SetValue) *bool { return v.IsSecret }).(pulumi.BoolPtrOutput)
}

func (o SetValueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SetValue) string { return v.Name }).(pulumi.StringOutput)
}

func (o SetValueOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SetValue) string { return v.Value }).(pulumi.StringOutput)
}

type SetValueArrayOutput struct{ *pulumi.OutputState }

func (SetValueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SetValue)(nil)).Elem()
}

func (o SetValueArrayOutput) ToSetValueArrayOutput() SetValueArrayOutput {
	return o
}

func (o SetValueArrayOutput) ToSetValueArrayOutputWithContext(ctx context.Context) SetValueArrayOutput {
	return o
}

func (o SetValueArrayOutput) Index(i pulumi.IntInput) SetValueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SetValue {
		return vs[0].([]SetValue)[vs[1].(int)]
	}).(SetValueOutput)
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

type SourceTriggerDescriptorResponse struct {
	BranchName    *string `pulumi:"branchName"`
	CommitId      *string `pulumi:"commitId"`
	EventType     *string `pulumi:"eventType"`
	Id            *string `pulumi:"id"`
	ProviderType  *string `pulumi:"providerType"`
	PullRequestId *string `pulumi:"pullRequestId"`
	RepositoryUrl *string `pulumi:"repositoryUrl"`
}





type SourceTriggerDescriptorResponseInput interface {
	pulumi.Input

	ToSourceTriggerDescriptorResponseOutput() SourceTriggerDescriptorResponseOutput
	ToSourceTriggerDescriptorResponseOutputWithContext(context.Context) SourceTriggerDescriptorResponseOutput
}

type SourceTriggerDescriptorResponseArgs struct {
	BranchName    pulumi.StringPtrInput `pulumi:"branchName"`
	CommitId      pulumi.StringPtrInput `pulumi:"commitId"`
	EventType     pulumi.StringPtrInput `pulumi:"eventType"`
	Id            pulumi.StringPtrInput `pulumi:"id"`
	ProviderType  pulumi.StringPtrInput `pulumi:"providerType"`
	PullRequestId pulumi.StringPtrInput `pulumi:"pullRequestId"`
	RepositoryUrl pulumi.StringPtrInput `pulumi:"repositoryUrl"`
}

func (SourceTriggerDescriptorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceTriggerDescriptorResponse)(nil)).Elem()
}

func (i SourceTriggerDescriptorResponseArgs) ToSourceTriggerDescriptorResponseOutput() SourceTriggerDescriptorResponseOutput {
	return i.ToSourceTriggerDescriptorResponseOutputWithContext(context.Background())
}

func (i SourceTriggerDescriptorResponseArgs) ToSourceTriggerDescriptorResponseOutputWithContext(ctx context.Context) SourceTriggerDescriptorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceTriggerDescriptorResponseOutput)
}

func (i SourceTriggerDescriptorResponseArgs) ToSourceTriggerDescriptorResponsePtrOutput() SourceTriggerDescriptorResponsePtrOutput {
	return i.ToSourceTriggerDescriptorResponsePtrOutputWithContext(context.Background())
}

func (i SourceTriggerDescriptorResponseArgs) ToSourceTriggerDescriptorResponsePtrOutputWithContext(ctx context.Context) SourceTriggerDescriptorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceTriggerDescriptorResponseOutput).ToSourceTriggerDescriptorResponsePtrOutputWithContext(ctx)
}









type SourceTriggerDescriptorResponsePtrInput interface {
	pulumi.Input

	ToSourceTriggerDescriptorResponsePtrOutput() SourceTriggerDescriptorResponsePtrOutput
	ToSourceTriggerDescriptorResponsePtrOutputWithContext(context.Context) SourceTriggerDescriptorResponsePtrOutput
}

type sourceTriggerDescriptorResponsePtrType SourceTriggerDescriptorResponseArgs

func SourceTriggerDescriptorResponsePtr(v *SourceTriggerDescriptorResponseArgs) SourceTriggerDescriptorResponsePtrInput {
	return (*sourceTriggerDescriptorResponsePtrType)(v)
}

func (*sourceTriggerDescriptorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceTriggerDescriptorResponse)(nil)).Elem()
}

func (i *sourceTriggerDescriptorResponsePtrType) ToSourceTriggerDescriptorResponsePtrOutput() SourceTriggerDescriptorResponsePtrOutput {
	return i.ToSourceTriggerDescriptorResponsePtrOutputWithContext(context.Background())
}

func (i *sourceTriggerDescriptorResponsePtrType) ToSourceTriggerDescriptorResponsePtrOutputWithContext(ctx context.Context) SourceTriggerDescriptorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceTriggerDescriptorResponsePtrOutput)
}

type SourceTriggerDescriptorResponseOutput struct{ *pulumi.OutputState }

func (SourceTriggerDescriptorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceTriggerDescriptorResponse)(nil)).Elem()
}

func (o SourceTriggerDescriptorResponseOutput) ToSourceTriggerDescriptorResponseOutput() SourceTriggerDescriptorResponseOutput {
	return o
}

func (o SourceTriggerDescriptorResponseOutput) ToSourceTriggerDescriptorResponseOutputWithContext(ctx context.Context) SourceTriggerDescriptorResponseOutput {
	return o
}

func (o SourceTriggerDescriptorResponseOutput) ToSourceTriggerDescriptorResponsePtrOutput() SourceTriggerDescriptorResponsePtrOutput {
	return o.ToSourceTriggerDescriptorResponsePtrOutputWithContext(context.Background())
}

func (o SourceTriggerDescriptorResponseOutput) ToSourceTriggerDescriptorResponsePtrOutputWithContext(ctx context.Context) SourceTriggerDescriptorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceTriggerDescriptorResponse) *SourceTriggerDescriptorResponse {
		return &v
	}).(SourceTriggerDescriptorResponsePtrOutput)
}

func (o SourceTriggerDescriptorResponseOutput) BranchName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceTriggerDescriptorResponse) *string { return v.BranchName }).(pulumi.StringPtrOutput)
}

func (o SourceTriggerDescriptorResponseOutput) CommitId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceTriggerDescriptorResponse) *string { return v.CommitId }).(pulumi.StringPtrOutput)
}

func (o SourceTriggerDescriptorResponseOutput) EventType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceTriggerDescriptorResponse) *string { return v.EventType }).(pulumi.StringPtrOutput)
}

func (o SourceTriggerDescriptorResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceTriggerDescriptorResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o SourceTriggerDescriptorResponseOutput) ProviderType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceTriggerDescriptorResponse) *string { return v.ProviderType }).(pulumi.StringPtrOutput)
}

func (o SourceTriggerDescriptorResponseOutput) PullRequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceTriggerDescriptorResponse) *string { return v.PullRequestId }).(pulumi.StringPtrOutput)
}

func (o SourceTriggerDescriptorResponseOutput) RepositoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceTriggerDescriptorResponse) *string { return v.RepositoryUrl }).(pulumi.StringPtrOutput)
}

type SourceTriggerDescriptorResponsePtrOutput struct{ *pulumi.OutputState }

func (SourceTriggerDescriptorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceTriggerDescriptorResponse)(nil)).Elem()
}

func (o SourceTriggerDescriptorResponsePtrOutput) ToSourceTriggerDescriptorResponsePtrOutput() SourceTriggerDescriptorResponsePtrOutput {
	return o
}

func (o SourceTriggerDescriptorResponsePtrOutput) ToSourceTriggerDescriptorResponsePtrOutputWithContext(ctx context.Context) SourceTriggerDescriptorResponsePtrOutput {
	return o
}

func (o SourceTriggerDescriptorResponsePtrOutput) Elem() SourceTriggerDescriptorResponseOutput {
	return o.ApplyT(func(v *SourceTriggerDescriptorResponse) SourceTriggerDescriptorResponse {
		if v != nil {
			return *v
		}
		var ret SourceTriggerDescriptorResponse
		return ret
	}).(SourceTriggerDescriptorResponseOutput)
}

func (o SourceTriggerDescriptorResponsePtrOutput) BranchName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceTriggerDescriptorResponse) *string {
		if v == nil {
			return nil
		}
		return v.BranchName
	}).(pulumi.StringPtrOutput)
}

func (o SourceTriggerDescriptorResponsePtrOutput) CommitId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceTriggerDescriptorResponse) *string {
		if v == nil {
			return nil
		}
		return v.CommitId
	}).(pulumi.StringPtrOutput)
}

func (o SourceTriggerDescriptorResponsePtrOutput) EventType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceTriggerDescriptorResponse) *string {
		if v == nil {
			return nil
		}
		return v.EventType
	}).(pulumi.StringPtrOutput)
}

func (o SourceTriggerDescriptorResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceTriggerDescriptorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SourceTriggerDescriptorResponsePtrOutput) ProviderType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceTriggerDescriptorResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProviderType
	}).(pulumi.StringPtrOutput)
}

func (o SourceTriggerDescriptorResponsePtrOutput) PullRequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceTriggerDescriptorResponse) *string {
		if v == nil {
			return nil
		}
		return v.PullRequestId
	}).(pulumi.StringPtrOutput)
}

func (o SourceTriggerDescriptorResponsePtrOutput) RepositoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceTriggerDescriptorResponse) *string {
		if v == nil {
			return nil
		}
		return v.RepositoryUrl
	}).(pulumi.StringPtrOutput)
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

type TaskRunRequest struct {
	AgentPoolName              *string                     `pulumi:"agentPoolName"`
	IsArchiveEnabled           *bool                       `pulumi:"isArchiveEnabled"`
	LogTemplate                *string                     `pulumi:"logTemplate"`
	OverrideTaskStepProperties *OverrideTaskStepProperties `pulumi:"overrideTaskStepProperties"`
	TaskId                     string                      `pulumi:"taskId"`
	Type                       string                      `pulumi:"type"`
}





type TaskRunRequestInput interface {
	pulumi.Input

	ToTaskRunRequestOutput() TaskRunRequestOutput
	ToTaskRunRequestOutputWithContext(context.Context) TaskRunRequestOutput
}

type TaskRunRequestArgs struct {
	AgentPoolName              pulumi.StringPtrInput              `pulumi:"agentPoolName"`
	IsArchiveEnabled           pulumi.BoolPtrInput                `pulumi:"isArchiveEnabled"`
	LogTemplate                pulumi.StringPtrInput              `pulumi:"logTemplate"`
	OverrideTaskStepProperties OverrideTaskStepPropertiesPtrInput `pulumi:"overrideTaskStepProperties"`
	TaskId                     pulumi.StringInput                 `pulumi:"taskId"`
	Type                       pulumi.StringInput                 `pulumi:"type"`
}

func (TaskRunRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskRunRequest)(nil)).Elem()
}

func (i TaskRunRequestArgs) ToTaskRunRequestOutput() TaskRunRequestOutput {
	return i.ToTaskRunRequestOutputWithContext(context.Background())
}

func (i TaskRunRequestArgs) ToTaskRunRequestOutputWithContext(ctx context.Context) TaskRunRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskRunRequestOutput)
}

type TaskRunRequestOutput struct{ *pulumi.OutputState }

func (TaskRunRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskRunRequest)(nil)).Elem()
}

func (o TaskRunRequestOutput) ToTaskRunRequestOutput() TaskRunRequestOutput {
	return o
}

func (o TaskRunRequestOutput) ToTaskRunRequestOutputWithContext(ctx context.Context) TaskRunRequestOutput {
	return o
}

func (o TaskRunRequestOutput) AgentPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskRunRequest) *string { return v.AgentPoolName }).(pulumi.StringPtrOutput)
}

func (o TaskRunRequestOutput) IsArchiveEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TaskRunRequest) *bool { return v.IsArchiveEnabled }).(pulumi.BoolPtrOutput)
}

func (o TaskRunRequestOutput) LogTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskRunRequest) *string { return v.LogTemplate }).(pulumi.StringPtrOutput)
}

func (o TaskRunRequestOutput) OverrideTaskStepProperties() OverrideTaskStepPropertiesPtrOutput {
	return o.ApplyT(func(v TaskRunRequest) *OverrideTaskStepProperties { return v.OverrideTaskStepProperties }).(OverrideTaskStepPropertiesPtrOutput)
}

func (o TaskRunRequestOutput) TaskId() pulumi.StringOutput {
	return o.ApplyT(func(v TaskRunRequest) string { return v.TaskId }).(pulumi.StringOutput)
}

func (o TaskRunRequestOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v TaskRunRequest) string { return v.Type }).(pulumi.StringOutput)
}

type TaskRunRequestResponse struct {
	AgentPoolName              *string                             `pulumi:"agentPoolName"`
	IsArchiveEnabled           *bool                               `pulumi:"isArchiveEnabled"`
	LogTemplate                *string                             `pulumi:"logTemplate"`
	OverrideTaskStepProperties *OverrideTaskStepPropertiesResponse `pulumi:"overrideTaskStepProperties"`
	TaskId                     string                              `pulumi:"taskId"`
	Type                       string                              `pulumi:"type"`
}





type TaskRunRequestResponseInput interface {
	pulumi.Input

	ToTaskRunRequestResponseOutput() TaskRunRequestResponseOutput
	ToTaskRunRequestResponseOutputWithContext(context.Context) TaskRunRequestResponseOutput
}

type TaskRunRequestResponseArgs struct {
	AgentPoolName              pulumi.StringPtrInput                      `pulumi:"agentPoolName"`
	IsArchiveEnabled           pulumi.BoolPtrInput                        `pulumi:"isArchiveEnabled"`
	LogTemplate                pulumi.StringPtrInput                      `pulumi:"logTemplate"`
	OverrideTaskStepProperties OverrideTaskStepPropertiesResponsePtrInput `pulumi:"overrideTaskStepProperties"`
	TaskId                     pulumi.StringInput                         `pulumi:"taskId"`
	Type                       pulumi.StringInput                         `pulumi:"type"`
}

func (TaskRunRequestResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskRunRequestResponse)(nil)).Elem()
}

func (i TaskRunRequestResponseArgs) ToTaskRunRequestResponseOutput() TaskRunRequestResponseOutput {
	return i.ToTaskRunRequestResponseOutputWithContext(context.Background())
}

func (i TaskRunRequestResponseArgs) ToTaskRunRequestResponseOutputWithContext(ctx context.Context) TaskRunRequestResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskRunRequestResponseOutput)
}

type TaskRunRequestResponseOutput struct{ *pulumi.OutputState }

func (TaskRunRequestResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskRunRequestResponse)(nil)).Elem()
}

func (o TaskRunRequestResponseOutput) ToTaskRunRequestResponseOutput() TaskRunRequestResponseOutput {
	return o
}

func (o TaskRunRequestResponseOutput) ToTaskRunRequestResponseOutputWithContext(ctx context.Context) TaskRunRequestResponseOutput {
	return o
}

func (o TaskRunRequestResponseOutput) AgentPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskRunRequestResponse) *string { return v.AgentPoolName }).(pulumi.StringPtrOutput)
}

func (o TaskRunRequestResponseOutput) IsArchiveEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TaskRunRequestResponse) *bool { return v.IsArchiveEnabled }).(pulumi.BoolPtrOutput)
}

func (o TaskRunRequestResponseOutput) LogTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskRunRequestResponse) *string { return v.LogTemplate }).(pulumi.StringPtrOutput)
}

func (o TaskRunRequestResponseOutput) OverrideTaskStepProperties() OverrideTaskStepPropertiesResponsePtrOutput {
	return o.ApplyT(func(v TaskRunRequestResponse) *OverrideTaskStepPropertiesResponse {
		return v.OverrideTaskStepProperties
	}).(OverrideTaskStepPropertiesResponsePtrOutput)
}

func (o TaskRunRequestResponseOutput) TaskId() pulumi.StringOutput {
	return o.ApplyT(func(v TaskRunRequestResponse) string { return v.TaskId }).(pulumi.StringOutput)
}

func (o TaskRunRequestResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v TaskRunRequestResponse) string { return v.Type }).(pulumi.StringOutput)
}

type TimerTrigger struct {
	Name     string  `pulumi:"name"`
	Schedule string  `pulumi:"schedule"`
	Status   *string `pulumi:"status"`
}





type TimerTriggerInput interface {
	pulumi.Input

	ToTimerTriggerOutput() TimerTriggerOutput
	ToTimerTriggerOutputWithContext(context.Context) TimerTriggerOutput
}

type TimerTriggerArgs struct {
	Name     pulumi.StringInput    `pulumi:"name"`
	Schedule pulumi.StringInput    `pulumi:"schedule"`
	Status   pulumi.StringPtrInput `pulumi:"status"`
}

func (TimerTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimerTrigger)(nil)).Elem()
}

func (i TimerTriggerArgs) ToTimerTriggerOutput() TimerTriggerOutput {
	return i.ToTimerTriggerOutputWithContext(context.Background())
}

func (i TimerTriggerArgs) ToTimerTriggerOutputWithContext(ctx context.Context) TimerTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimerTriggerOutput)
}





type TimerTriggerArrayInput interface {
	pulumi.Input

	ToTimerTriggerArrayOutput() TimerTriggerArrayOutput
	ToTimerTriggerArrayOutputWithContext(context.Context) TimerTriggerArrayOutput
}

type TimerTriggerArray []TimerTriggerInput

func (TimerTriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimerTrigger)(nil)).Elem()
}

func (i TimerTriggerArray) ToTimerTriggerArrayOutput() TimerTriggerArrayOutput {
	return i.ToTimerTriggerArrayOutputWithContext(context.Background())
}

func (i TimerTriggerArray) ToTimerTriggerArrayOutputWithContext(ctx context.Context) TimerTriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimerTriggerArrayOutput)
}

type TimerTriggerOutput struct{ *pulumi.OutputState }

func (TimerTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimerTrigger)(nil)).Elem()
}

func (o TimerTriggerOutput) ToTimerTriggerOutput() TimerTriggerOutput {
	return o
}

func (o TimerTriggerOutput) ToTimerTriggerOutputWithContext(ctx context.Context) TimerTriggerOutput {
	return o
}

func (o TimerTriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TimerTrigger) string { return v.Name }).(pulumi.StringOutput)
}

func (o TimerTriggerOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v TimerTrigger) string { return v.Schedule }).(pulumi.StringOutput)
}

func (o TimerTriggerOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TimerTrigger) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type TimerTriggerArrayOutput struct{ *pulumi.OutputState }

func (TimerTriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimerTrigger)(nil)).Elem()
}

func (o TimerTriggerArrayOutput) ToTimerTriggerArrayOutput() TimerTriggerArrayOutput {
	return o
}

func (o TimerTriggerArrayOutput) ToTimerTriggerArrayOutputWithContext(ctx context.Context) TimerTriggerArrayOutput {
	return o
}

func (o TimerTriggerArrayOutput) Index(i pulumi.IntInput) TimerTriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimerTrigger {
		return vs[0].([]TimerTrigger)[vs[1].(int)]
	}).(TimerTriggerOutput)
}

type TimerTriggerDescriptorResponse struct {
	ScheduleOccurrence *string `pulumi:"scheduleOccurrence"`
	TimerTriggerName   *string `pulumi:"timerTriggerName"`
}





type TimerTriggerDescriptorResponseInput interface {
	pulumi.Input

	ToTimerTriggerDescriptorResponseOutput() TimerTriggerDescriptorResponseOutput
	ToTimerTriggerDescriptorResponseOutputWithContext(context.Context) TimerTriggerDescriptorResponseOutput
}

type TimerTriggerDescriptorResponseArgs struct {
	ScheduleOccurrence pulumi.StringPtrInput `pulumi:"scheduleOccurrence"`
	TimerTriggerName   pulumi.StringPtrInput `pulumi:"timerTriggerName"`
}

func (TimerTriggerDescriptorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimerTriggerDescriptorResponse)(nil)).Elem()
}

func (i TimerTriggerDescriptorResponseArgs) ToTimerTriggerDescriptorResponseOutput() TimerTriggerDescriptorResponseOutput {
	return i.ToTimerTriggerDescriptorResponseOutputWithContext(context.Background())
}

func (i TimerTriggerDescriptorResponseArgs) ToTimerTriggerDescriptorResponseOutputWithContext(ctx context.Context) TimerTriggerDescriptorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimerTriggerDescriptorResponseOutput)
}

func (i TimerTriggerDescriptorResponseArgs) ToTimerTriggerDescriptorResponsePtrOutput() TimerTriggerDescriptorResponsePtrOutput {
	return i.ToTimerTriggerDescriptorResponsePtrOutputWithContext(context.Background())
}

func (i TimerTriggerDescriptorResponseArgs) ToTimerTriggerDescriptorResponsePtrOutputWithContext(ctx context.Context) TimerTriggerDescriptorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimerTriggerDescriptorResponseOutput).ToTimerTriggerDescriptorResponsePtrOutputWithContext(ctx)
}









type TimerTriggerDescriptorResponsePtrInput interface {
	pulumi.Input

	ToTimerTriggerDescriptorResponsePtrOutput() TimerTriggerDescriptorResponsePtrOutput
	ToTimerTriggerDescriptorResponsePtrOutputWithContext(context.Context) TimerTriggerDescriptorResponsePtrOutput
}

type timerTriggerDescriptorResponsePtrType TimerTriggerDescriptorResponseArgs

func TimerTriggerDescriptorResponsePtr(v *TimerTriggerDescriptorResponseArgs) TimerTriggerDescriptorResponsePtrInput {
	return (*timerTriggerDescriptorResponsePtrType)(v)
}

func (*timerTriggerDescriptorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TimerTriggerDescriptorResponse)(nil)).Elem()
}

func (i *timerTriggerDescriptorResponsePtrType) ToTimerTriggerDescriptorResponsePtrOutput() TimerTriggerDescriptorResponsePtrOutput {
	return i.ToTimerTriggerDescriptorResponsePtrOutputWithContext(context.Background())
}

func (i *timerTriggerDescriptorResponsePtrType) ToTimerTriggerDescriptorResponsePtrOutputWithContext(ctx context.Context) TimerTriggerDescriptorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimerTriggerDescriptorResponsePtrOutput)
}

type TimerTriggerDescriptorResponseOutput struct{ *pulumi.OutputState }

func (TimerTriggerDescriptorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimerTriggerDescriptorResponse)(nil)).Elem()
}

func (o TimerTriggerDescriptorResponseOutput) ToTimerTriggerDescriptorResponseOutput() TimerTriggerDescriptorResponseOutput {
	return o
}

func (o TimerTriggerDescriptorResponseOutput) ToTimerTriggerDescriptorResponseOutputWithContext(ctx context.Context) TimerTriggerDescriptorResponseOutput {
	return o
}

func (o TimerTriggerDescriptorResponseOutput) ToTimerTriggerDescriptorResponsePtrOutput() TimerTriggerDescriptorResponsePtrOutput {
	return o.ToTimerTriggerDescriptorResponsePtrOutputWithContext(context.Background())
}

func (o TimerTriggerDescriptorResponseOutput) ToTimerTriggerDescriptorResponsePtrOutputWithContext(ctx context.Context) TimerTriggerDescriptorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TimerTriggerDescriptorResponse) *TimerTriggerDescriptorResponse {
		return &v
	}).(TimerTriggerDescriptorResponsePtrOutput)
}

func (o TimerTriggerDescriptorResponseOutput) ScheduleOccurrence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TimerTriggerDescriptorResponse) *string { return v.ScheduleOccurrence }).(pulumi.StringPtrOutput)
}

func (o TimerTriggerDescriptorResponseOutput) TimerTriggerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TimerTriggerDescriptorResponse) *string { return v.TimerTriggerName }).(pulumi.StringPtrOutput)
}

type TimerTriggerDescriptorResponsePtrOutput struct{ *pulumi.OutputState }

func (TimerTriggerDescriptorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimerTriggerDescriptorResponse)(nil)).Elem()
}

func (o TimerTriggerDescriptorResponsePtrOutput) ToTimerTriggerDescriptorResponsePtrOutput() TimerTriggerDescriptorResponsePtrOutput {
	return o
}

func (o TimerTriggerDescriptorResponsePtrOutput) ToTimerTriggerDescriptorResponsePtrOutputWithContext(ctx context.Context) TimerTriggerDescriptorResponsePtrOutput {
	return o
}

func (o TimerTriggerDescriptorResponsePtrOutput) Elem() TimerTriggerDescriptorResponseOutput {
	return o.ApplyT(func(v *TimerTriggerDescriptorResponse) TimerTriggerDescriptorResponse {
		if v != nil {
			return *v
		}
		var ret TimerTriggerDescriptorResponse
		return ret
	}).(TimerTriggerDescriptorResponseOutput)
}

func (o TimerTriggerDescriptorResponsePtrOutput) ScheduleOccurrence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TimerTriggerDescriptorResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScheduleOccurrence
	}).(pulumi.StringPtrOutput)
}

func (o TimerTriggerDescriptorResponsePtrOutput) TimerTriggerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TimerTriggerDescriptorResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimerTriggerName
	}).(pulumi.StringPtrOutput)
}

type TimerTriggerResponse struct {
	Name     string  `pulumi:"name"`
	Schedule string  `pulumi:"schedule"`
	Status   *string `pulumi:"status"`
}





type TimerTriggerResponseInput interface {
	pulumi.Input

	ToTimerTriggerResponseOutput() TimerTriggerResponseOutput
	ToTimerTriggerResponseOutputWithContext(context.Context) TimerTriggerResponseOutput
}

type TimerTriggerResponseArgs struct {
	Name     pulumi.StringInput    `pulumi:"name"`
	Schedule pulumi.StringInput    `pulumi:"schedule"`
	Status   pulumi.StringPtrInput `pulumi:"status"`
}

func (TimerTriggerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimerTriggerResponse)(nil)).Elem()
}

func (i TimerTriggerResponseArgs) ToTimerTriggerResponseOutput() TimerTriggerResponseOutput {
	return i.ToTimerTriggerResponseOutputWithContext(context.Background())
}

func (i TimerTriggerResponseArgs) ToTimerTriggerResponseOutputWithContext(ctx context.Context) TimerTriggerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimerTriggerResponseOutput)
}





type TimerTriggerResponseArrayInput interface {
	pulumi.Input

	ToTimerTriggerResponseArrayOutput() TimerTriggerResponseArrayOutput
	ToTimerTriggerResponseArrayOutputWithContext(context.Context) TimerTriggerResponseArrayOutput
}

type TimerTriggerResponseArray []TimerTriggerResponseInput

func (TimerTriggerResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimerTriggerResponse)(nil)).Elem()
}

func (i TimerTriggerResponseArray) ToTimerTriggerResponseArrayOutput() TimerTriggerResponseArrayOutput {
	return i.ToTimerTriggerResponseArrayOutputWithContext(context.Background())
}

func (i TimerTriggerResponseArray) ToTimerTriggerResponseArrayOutputWithContext(ctx context.Context) TimerTriggerResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimerTriggerResponseArrayOutput)
}

type TimerTriggerResponseOutput struct{ *pulumi.OutputState }

func (TimerTriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimerTriggerResponse)(nil)).Elem()
}

func (o TimerTriggerResponseOutput) ToTimerTriggerResponseOutput() TimerTriggerResponseOutput {
	return o
}

func (o TimerTriggerResponseOutput) ToTimerTriggerResponseOutputWithContext(ctx context.Context) TimerTriggerResponseOutput {
	return o
}

func (o TimerTriggerResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TimerTriggerResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o TimerTriggerResponseOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v TimerTriggerResponse) string { return v.Schedule }).(pulumi.StringOutput)
}

func (o TimerTriggerResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TimerTriggerResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type TimerTriggerResponseArrayOutput struct{ *pulumi.OutputState }

func (TimerTriggerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimerTriggerResponse)(nil)).Elem()
}

func (o TimerTriggerResponseArrayOutput) ToTimerTriggerResponseArrayOutput() TimerTriggerResponseArrayOutput {
	return o
}

func (o TimerTriggerResponseArrayOutput) ToTimerTriggerResponseArrayOutputWithContext(ctx context.Context) TimerTriggerResponseArrayOutput {
	return o
}

func (o TimerTriggerResponseArrayOutput) Index(i pulumi.IntInput) TimerTriggerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimerTriggerResponse {
		return vs[0].([]TimerTriggerResponse)[vs[1].(int)]
	}).(TimerTriggerResponseOutput)
}

type TriggerProperties struct {
	BaseImageTrigger *BaseImageTrigger `pulumi:"baseImageTrigger"`
	SourceTriggers   []SourceTrigger   `pulumi:"sourceTriggers"`
	TimerTriggers    []TimerTrigger    `pulumi:"timerTriggers"`
}





type TriggerPropertiesInput interface {
	pulumi.Input

	ToTriggerPropertiesOutput() TriggerPropertiesOutput
	ToTriggerPropertiesOutputWithContext(context.Context) TriggerPropertiesOutput
}

type TriggerPropertiesArgs struct {
	BaseImageTrigger BaseImageTriggerPtrInput `pulumi:"baseImageTrigger"`
	SourceTriggers   SourceTriggerArrayInput  `pulumi:"sourceTriggers"`
	TimerTriggers    TimerTriggerArrayInput   `pulumi:"timerTriggers"`
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

func (o TriggerPropertiesOutput) TimerTriggers() TimerTriggerArrayOutput {
	return o.ApplyT(func(v TriggerProperties) []TimerTrigger { return v.TimerTriggers }).(TimerTriggerArrayOutput)
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

func (o TriggerPropertiesPtrOutput) TimerTriggers() TimerTriggerArrayOutput {
	return o.ApplyT(func(v *TriggerProperties) []TimerTrigger {
		if v == nil {
			return nil
		}
		return v.TimerTriggers
	}).(TimerTriggerArrayOutput)
}

type TriggerPropertiesResponse struct {
	BaseImageTrigger *BaseImageTriggerResponse `pulumi:"baseImageTrigger"`
	SourceTriggers   []SourceTriggerResponse   `pulumi:"sourceTriggers"`
	TimerTriggers    []TimerTriggerResponse    `pulumi:"timerTriggers"`
}





type TriggerPropertiesResponseInput interface {
	pulumi.Input

	ToTriggerPropertiesResponseOutput() TriggerPropertiesResponseOutput
	ToTriggerPropertiesResponseOutputWithContext(context.Context) TriggerPropertiesResponseOutput
}

type TriggerPropertiesResponseArgs struct {
	BaseImageTrigger BaseImageTriggerResponsePtrInput `pulumi:"baseImageTrigger"`
	SourceTriggers   SourceTriggerResponseArrayInput  `pulumi:"sourceTriggers"`
	TimerTriggers    TimerTriggerResponseArrayInput   `pulumi:"timerTriggers"`
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

func (o TriggerPropertiesResponseOutput) TimerTriggers() TimerTriggerResponseArrayOutput {
	return o.ApplyT(func(v TriggerPropertiesResponse) []TimerTriggerResponse { return v.TimerTriggers }).(TimerTriggerResponseArrayOutput)
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

func (o TriggerPropertiesResponsePtrOutput) TimerTriggers() TimerTriggerResponseArrayOutput {
	return o.ApplyT(func(v *TriggerPropertiesResponse) []TimerTriggerResponse {
		if v == nil {
			return nil
		}
		return v.TimerTriggers
	}).(TimerTriggerResponseArrayOutput)
}

type UserIdentityProperties struct {
	ClientId    *string `pulumi:"clientId"`
	PrincipalId *string `pulumi:"principalId"`
}





type UserIdentityPropertiesInput interface {
	pulumi.Input

	ToUserIdentityPropertiesOutput() UserIdentityPropertiesOutput
	ToUserIdentityPropertiesOutputWithContext(context.Context) UserIdentityPropertiesOutput
}

type UserIdentityPropertiesArgs struct {
	ClientId    pulumi.StringPtrInput `pulumi:"clientId"`
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
}

func (UserIdentityPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityProperties)(nil)).Elem()
}

func (i UserIdentityPropertiesArgs) ToUserIdentityPropertiesOutput() UserIdentityPropertiesOutput {
	return i.ToUserIdentityPropertiesOutputWithContext(context.Background())
}

func (i UserIdentityPropertiesArgs) ToUserIdentityPropertiesOutputWithContext(ctx context.Context) UserIdentityPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityPropertiesOutput)
}





type UserIdentityPropertiesMapInput interface {
	pulumi.Input

	ToUserIdentityPropertiesMapOutput() UserIdentityPropertiesMapOutput
	ToUserIdentityPropertiesMapOutputWithContext(context.Context) UserIdentityPropertiesMapOutput
}

type UserIdentityPropertiesMap map[string]UserIdentityPropertiesInput

func (UserIdentityPropertiesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityProperties)(nil)).Elem()
}

func (i UserIdentityPropertiesMap) ToUserIdentityPropertiesMapOutput() UserIdentityPropertiesMapOutput {
	return i.ToUserIdentityPropertiesMapOutputWithContext(context.Background())
}

func (i UserIdentityPropertiesMap) ToUserIdentityPropertiesMapOutputWithContext(ctx context.Context) UserIdentityPropertiesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityPropertiesMapOutput)
}

type UserIdentityPropertiesOutput struct{ *pulumi.OutputState }

func (UserIdentityPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityProperties)(nil)).Elem()
}

func (o UserIdentityPropertiesOutput) ToUserIdentityPropertiesOutput() UserIdentityPropertiesOutput {
	return o
}

func (o UserIdentityPropertiesOutput) ToUserIdentityPropertiesOutputWithContext(ctx context.Context) UserIdentityPropertiesOutput {
	return o
}

func (o UserIdentityPropertiesOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityProperties) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o UserIdentityPropertiesOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityProperties) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

type UserIdentityPropertiesMapOutput struct{ *pulumi.OutputState }

func (UserIdentityPropertiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityProperties)(nil)).Elem()
}

func (o UserIdentityPropertiesMapOutput) ToUserIdentityPropertiesMapOutput() UserIdentityPropertiesMapOutput {
	return o
}

func (o UserIdentityPropertiesMapOutput) ToUserIdentityPropertiesMapOutputWithContext(ctx context.Context) UserIdentityPropertiesMapOutput {
	return o
}

func (o UserIdentityPropertiesMapOutput) MapIndex(k pulumi.StringInput) UserIdentityPropertiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserIdentityProperties {
		return vs[0].(map[string]UserIdentityProperties)[vs[1].(string)]
	}).(UserIdentityPropertiesOutput)
}

type UserIdentityPropertiesResponse struct {
	ClientId    *string `pulumi:"clientId"`
	PrincipalId *string `pulumi:"principalId"`
}





type UserIdentityPropertiesResponseInput interface {
	pulumi.Input

	ToUserIdentityPropertiesResponseOutput() UserIdentityPropertiesResponseOutput
	ToUserIdentityPropertiesResponseOutputWithContext(context.Context) UserIdentityPropertiesResponseOutput
}

type UserIdentityPropertiesResponseArgs struct {
	ClientId    pulumi.StringPtrInput `pulumi:"clientId"`
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
}

func (UserIdentityPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityPropertiesResponse)(nil)).Elem()
}

func (i UserIdentityPropertiesResponseArgs) ToUserIdentityPropertiesResponseOutput() UserIdentityPropertiesResponseOutput {
	return i.ToUserIdentityPropertiesResponseOutputWithContext(context.Background())
}

func (i UserIdentityPropertiesResponseArgs) ToUserIdentityPropertiesResponseOutputWithContext(ctx context.Context) UserIdentityPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityPropertiesResponseOutput)
}





type UserIdentityPropertiesResponseMapInput interface {
	pulumi.Input

	ToUserIdentityPropertiesResponseMapOutput() UserIdentityPropertiesResponseMapOutput
	ToUserIdentityPropertiesResponseMapOutputWithContext(context.Context) UserIdentityPropertiesResponseMapOutput
}

type UserIdentityPropertiesResponseMap map[string]UserIdentityPropertiesResponseInput

func (UserIdentityPropertiesResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityPropertiesResponse)(nil)).Elem()
}

func (i UserIdentityPropertiesResponseMap) ToUserIdentityPropertiesResponseMapOutput() UserIdentityPropertiesResponseMapOutput {
	return i.ToUserIdentityPropertiesResponseMapOutputWithContext(context.Background())
}

func (i UserIdentityPropertiesResponseMap) ToUserIdentityPropertiesResponseMapOutputWithContext(ctx context.Context) UserIdentityPropertiesResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityPropertiesResponseMapOutput)
}

type UserIdentityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (UserIdentityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityPropertiesResponse)(nil)).Elem()
}

func (o UserIdentityPropertiesResponseOutput) ToUserIdentityPropertiesResponseOutput() UserIdentityPropertiesResponseOutput {
	return o
}

func (o UserIdentityPropertiesResponseOutput) ToUserIdentityPropertiesResponseOutputWithContext(ctx context.Context) UserIdentityPropertiesResponseOutput {
	return o
}

func (o UserIdentityPropertiesResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityPropertiesResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o UserIdentityPropertiesResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityPropertiesResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

type UserIdentityPropertiesResponseMapOutput struct{ *pulumi.OutputState }

func (UserIdentityPropertiesResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityPropertiesResponse)(nil)).Elem()
}

func (o UserIdentityPropertiesResponseMapOutput) ToUserIdentityPropertiesResponseMapOutput() UserIdentityPropertiesResponseMapOutput {
	return o
}

func (o UserIdentityPropertiesResponseMapOutput) ToUserIdentityPropertiesResponseMapOutputWithContext(ctx context.Context) UserIdentityPropertiesResponseMapOutput {
	return o
}

func (o UserIdentityPropertiesResponseMapOutput) MapIndex(k pulumi.StringInput) UserIdentityPropertiesResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserIdentityPropertiesResponse {
		return vs[0].(map[string]UserIdentityPropertiesResponse)[vs[1].(string)]
	}).(UserIdentityPropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AgentPropertiesOutput{})
	pulumi.RegisterOutputType(AgentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(AgentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AgentPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ArgumentOutput{})
	pulumi.RegisterOutputType(ArgumentArrayOutput{})
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
	pulumi.RegisterOutputType(DockerBuildRequestOutput{})
	pulumi.RegisterOutputType(DockerBuildRequestResponseOutput{})
	pulumi.RegisterOutputType(DockerBuildStepOutput{})
	pulumi.RegisterOutputType(DockerBuildStepResponseOutput{})
	pulumi.RegisterOutputType(EncodedTaskRunRequestOutput{})
	pulumi.RegisterOutputType(EncodedTaskRunRequestResponseOutput{})
	pulumi.RegisterOutputType(EncodedTaskStepOutput{})
	pulumi.RegisterOutputType(EncodedTaskStepResponseOutput{})
	pulumi.RegisterOutputType(FileTaskRunRequestOutput{})
	pulumi.RegisterOutputType(FileTaskRunRequestResponseOutput{})
	pulumi.RegisterOutputType(FileTaskStepOutput{})
	pulumi.RegisterOutputType(FileTaskStepResponseOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesPtrOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageDescriptorResponseOutput{})
	pulumi.RegisterOutputType(ImageDescriptorResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageDescriptorResponseArrayOutput{})
	pulumi.RegisterOutputType(ImageUpdateTriggerResponseOutput{})
	pulumi.RegisterOutputType(ImageUpdateTriggerResponsePtrOutput{})
	pulumi.RegisterOutputType(OverrideTaskStepPropertiesOutput{})
	pulumi.RegisterOutputType(OverrideTaskStepPropertiesPtrOutput{})
	pulumi.RegisterOutputType(OverrideTaskStepPropertiesResponseOutput{})
	pulumi.RegisterOutputType(OverrideTaskStepPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(RunResponseOutput{})
	pulumi.RegisterOutputType(RunResponsePtrOutput{})
	pulumi.RegisterOutputType(SecretObjectOutput{})
	pulumi.RegisterOutputType(SecretObjectPtrOutput{})
	pulumi.RegisterOutputType(SecretObjectResponseOutput{})
	pulumi.RegisterOutputType(SecretObjectResponsePtrOutput{})
	pulumi.RegisterOutputType(SetValueOutput{})
	pulumi.RegisterOutputType(SetValueArrayOutput{})
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
	pulumi.RegisterOutputType(SourceTriggerDescriptorResponseOutput{})
	pulumi.RegisterOutputType(SourceTriggerDescriptorResponsePtrOutput{})
	pulumi.RegisterOutputType(SourceTriggerResponseOutput{})
	pulumi.RegisterOutputType(SourceTriggerResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(TaskRunRequestOutput{})
	pulumi.RegisterOutputType(TaskRunRequestResponseOutput{})
	pulumi.RegisterOutputType(TimerTriggerOutput{})
	pulumi.RegisterOutputType(TimerTriggerArrayOutput{})
	pulumi.RegisterOutputType(TimerTriggerDescriptorResponseOutput{})
	pulumi.RegisterOutputType(TimerTriggerDescriptorResponsePtrOutput{})
	pulumi.RegisterOutputType(TimerTriggerResponseOutput{})
	pulumi.RegisterOutputType(TimerTriggerResponseArrayOutput{})
	pulumi.RegisterOutputType(TriggerPropertiesOutput{})
	pulumi.RegisterOutputType(TriggerPropertiesPtrOutput{})
	pulumi.RegisterOutputType(TriggerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TriggerPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesMapOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesResponseMapOutput{})
}




package containerregistry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActivationPropertiesResponse struct {
	Status string `pulumi:"status"`
}

type ActivationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ActivationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivationPropertiesResponse)(nil)).Elem()
}

func (o ActivationPropertiesResponseOutput) ToActivationPropertiesResponseOutput() ActivationPropertiesResponseOutput {
	return o
}

func (o ActivationPropertiesResponseOutput) ToActivationPropertiesResponseOutputWithContext(ctx context.Context) ActivationPropertiesResponseOutput {
	return o
}

func (o ActivationPropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ActivationPropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

type ActorResponse struct {
	Name *string `pulumi:"name"`
}

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


func (val *Argument) Defaults() *Argument {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsSecret) {
		isSecret_ := false
		tmp.IsSecret = &isSecret_
	}
	return &tmp
}

type ArgumentResponse struct {
	IsSecret *bool  `pulumi:"isSecret"`
	Name     string `pulumi:"name"`
	Value    string `pulumi:"value"`
}


func (val *ArgumentResponse) Defaults() *ArgumentResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsSecret) {
		isSecret_ := false
		tmp.IsSecret = &isSecret_
	}
	return &tmp
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

type BaseImageTrigger struct {
	BaseImageTriggerType     string  `pulumi:"baseImageTriggerType"`
	Name                     string  `pulumi:"name"`
	Status                   *string `pulumi:"status"`
	UpdateTriggerEndpoint    *string `pulumi:"updateTriggerEndpoint"`
	UpdateTriggerPayloadType *string `pulumi:"updateTriggerPayloadType"`
}


func (val *BaseImageTrigger) Defaults() *BaseImageTrigger {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "Enabled"
		tmp.Status = &status_
	}
	return &tmp
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


func (val *BaseImageTriggerResponse) Defaults() *BaseImageTriggerResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "Enabled"
		tmp.Status = &status_
	}
	return &tmp
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

type BuildArgumentResponse struct {
	IsSecret *bool  `pulumi:"isSecret"`
	Name     string `pulumi:"name"`
	Type     string `pulumi:"type"`
	Value    string `pulumi:"value"`
}


func (val *BuildArgumentResponse) Defaults() *BuildArgumentResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsSecret) {
		isSecret_ := false
		tmp.IsSecret = &isSecret_
	}
	return &tmp
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


func (val *DockerBuildRequest) Defaults() *DockerBuildRequest {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchiveEnabled) {
		isArchiveEnabled_ := false
		tmp.IsArchiveEnabled = &isArchiveEnabled_
	}
	if isZero(tmp.IsPushEnabled) {
		isPushEnabled_ := true
		tmp.IsPushEnabled = &isPushEnabled_
	}
	if isZero(tmp.NoCache) {
		noCache_ := false
		tmp.NoCache = &noCache_
	}
	if isZero(tmp.Timeout) {
		timeout_ := 3600
		tmp.Timeout = &timeout_
	}
	return &tmp
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


func (val *DockerBuildRequestResponse) Defaults() *DockerBuildRequestResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchiveEnabled) {
		isArchiveEnabled_ := false
		tmp.IsArchiveEnabled = &isArchiveEnabled_
	}
	if isZero(tmp.IsPushEnabled) {
		isPushEnabled_ := true
		tmp.IsPushEnabled = &isPushEnabled_
	}
	if isZero(tmp.NoCache) {
		noCache_ := false
		tmp.NoCache = &noCache_
	}
	if isZero(tmp.Timeout) {
		timeout_ := 3600
		tmp.Timeout = &timeout_
	}
	return &tmp
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


func (val *DockerBuildStep) Defaults() *DockerBuildStep {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsPushEnabled) {
		isPushEnabled_ := true
		tmp.IsPushEnabled = &isPushEnabled_
	}
	if isZero(tmp.NoCache) {
		noCache_ := false
		tmp.NoCache = &noCache_
	}
	return &tmp
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


func (val *DockerBuildStepResponse) Defaults() *DockerBuildStepResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsPushEnabled) {
		isPushEnabled_ := true
		tmp.IsPushEnabled = &isPushEnabled_
	}
	if isZero(tmp.NoCache) {
		noCache_ := false
		tmp.NoCache = &noCache_
	}
	return &tmp
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


func (val *EncodedTaskRunRequest) Defaults() *EncodedTaskRunRequest {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchiveEnabled) {
		isArchiveEnabled_ := false
		tmp.IsArchiveEnabled = &isArchiveEnabled_
	}
	if isZero(tmp.Timeout) {
		timeout_ := 3600
		tmp.Timeout = &timeout_
	}
	return &tmp
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


func (val *EncodedTaskRunRequestResponse) Defaults() *EncodedTaskRunRequestResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchiveEnabled) {
		isArchiveEnabled_ := false
		tmp.IsArchiveEnabled = &isArchiveEnabled_
	}
	if isZero(tmp.Timeout) {
		timeout_ := 3600
		tmp.Timeout = &timeout_
	}
	return &tmp
}

type EncodedTaskStep struct {
	ContextAccessToken   *string    `pulumi:"contextAccessToken"`
	ContextPath          *string    `pulumi:"contextPath"`
	EncodedTaskContent   string     `pulumi:"encodedTaskContent"`
	EncodedValuesContent *string    `pulumi:"encodedValuesContent"`
	Type                 string     `pulumi:"type"`
	Values               []SetValue `pulumi:"values"`
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

type EventContentResponse struct {
	Action    *string          `pulumi:"action"`
	Actor     *ActorResponse   `pulumi:"actor"`
	Id        *string          `pulumi:"id"`
	Request   *RequestResponse `pulumi:"request"`
	Source    *SourceResponse  `pulumi:"source"`
	Target    *TargetResponse  `pulumi:"target"`
	Timestamp *string          `pulumi:"timestamp"`
}

type EventRequestMessageResponse struct {
	Content    *EventContentResponse `pulumi:"content"`
	Headers    map[string]string     `pulumi:"headers"`
	Method     *string               `pulumi:"method"`
	RequestUri *string               `pulumi:"requestUri"`
	Version    *string               `pulumi:"version"`
}

type EventResponse struct {
	EventRequestMessage  *EventRequestMessageResponse  `pulumi:"eventRequestMessage"`
	EventResponseMessage *EventResponseMessageResponse `pulumi:"eventResponseMessage"`
	Id                   *string                       `pulumi:"id"`
}

type EventResponseMessageResponse struct {
	Content      *string           `pulumi:"content"`
	Headers      map[string]string `pulumi:"headers"`
	ReasonPhrase *string           `pulumi:"reasonPhrase"`
	StatusCode   *string           `pulumi:"statusCode"`
	Version      *string           `pulumi:"version"`
}

type ExportPipelineTargetProperties struct {
	KeyVaultUri string  `pulumi:"keyVaultUri"`
	Type        *string `pulumi:"type"`
	Uri         *string `pulumi:"uri"`
}





type ExportPipelineTargetPropertiesInput interface {
	pulumi.Input

	ToExportPipelineTargetPropertiesOutput() ExportPipelineTargetPropertiesOutput
	ToExportPipelineTargetPropertiesOutputWithContext(context.Context) ExportPipelineTargetPropertiesOutput
}

type ExportPipelineTargetPropertiesArgs struct {
	KeyVaultUri pulumi.StringInput    `pulumi:"keyVaultUri"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
	Uri         pulumi.StringPtrInput `pulumi:"uri"`
}

func (ExportPipelineTargetPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportPipelineTargetProperties)(nil)).Elem()
}

func (i ExportPipelineTargetPropertiesArgs) ToExportPipelineTargetPropertiesOutput() ExportPipelineTargetPropertiesOutput {
	return i.ToExportPipelineTargetPropertiesOutputWithContext(context.Background())
}

func (i ExportPipelineTargetPropertiesArgs) ToExportPipelineTargetPropertiesOutputWithContext(ctx context.Context) ExportPipelineTargetPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportPipelineTargetPropertiesOutput)
}

type ExportPipelineTargetPropertiesOutput struct{ *pulumi.OutputState }

func (ExportPipelineTargetPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportPipelineTargetProperties)(nil)).Elem()
}

func (o ExportPipelineTargetPropertiesOutput) ToExportPipelineTargetPropertiesOutput() ExportPipelineTargetPropertiesOutput {
	return o
}

func (o ExportPipelineTargetPropertiesOutput) ToExportPipelineTargetPropertiesOutputWithContext(ctx context.Context) ExportPipelineTargetPropertiesOutput {
	return o
}

func (o ExportPipelineTargetPropertiesOutput) KeyVaultUri() pulumi.StringOutput {
	return o.ApplyT(func(v ExportPipelineTargetProperties) string { return v.KeyVaultUri }).(pulumi.StringOutput)
}

func (o ExportPipelineTargetPropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportPipelineTargetProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ExportPipelineTargetPropertiesOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportPipelineTargetProperties) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type ExportPipelineTargetPropertiesResponse struct {
	KeyVaultUri string  `pulumi:"keyVaultUri"`
	Type        *string `pulumi:"type"`
	Uri         *string `pulumi:"uri"`
}

type ExportPipelineTargetPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ExportPipelineTargetPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportPipelineTargetPropertiesResponse)(nil)).Elem()
}

func (o ExportPipelineTargetPropertiesResponseOutput) ToExportPipelineTargetPropertiesResponseOutput() ExportPipelineTargetPropertiesResponseOutput {
	return o
}

func (o ExportPipelineTargetPropertiesResponseOutput) ToExportPipelineTargetPropertiesResponseOutputWithContext(ctx context.Context) ExportPipelineTargetPropertiesResponseOutput {
	return o
}

func (o ExportPipelineTargetPropertiesResponseOutput) KeyVaultUri() pulumi.StringOutput {
	return o.ApplyT(func(v ExportPipelineTargetPropertiesResponse) string { return v.KeyVaultUri }).(pulumi.StringOutput)
}

func (o ExportPipelineTargetPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportPipelineTargetPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ExportPipelineTargetPropertiesResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportPipelineTargetPropertiesResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type ExportPipelineTargetPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ExportPipelineTargetPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportPipelineTargetPropertiesResponse)(nil)).Elem()
}

func (o ExportPipelineTargetPropertiesResponsePtrOutput) ToExportPipelineTargetPropertiesResponsePtrOutput() ExportPipelineTargetPropertiesResponsePtrOutput {
	return o
}

func (o ExportPipelineTargetPropertiesResponsePtrOutput) ToExportPipelineTargetPropertiesResponsePtrOutputWithContext(ctx context.Context) ExportPipelineTargetPropertiesResponsePtrOutput {
	return o
}

func (o ExportPipelineTargetPropertiesResponsePtrOutput) Elem() ExportPipelineTargetPropertiesResponseOutput {
	return o.ApplyT(func(v *ExportPipelineTargetPropertiesResponse) ExportPipelineTargetPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ExportPipelineTargetPropertiesResponse
		return ret
	}).(ExportPipelineTargetPropertiesResponseOutput)
}

func (o ExportPipelineTargetPropertiesResponsePtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportPipelineTargetPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o ExportPipelineTargetPropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportPipelineTargetPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ExportPipelineTargetPropertiesResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportPipelineTargetPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
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


func (val *FileTaskRunRequest) Defaults() *FileTaskRunRequest {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchiveEnabled) {
		isArchiveEnabled_ := false
		tmp.IsArchiveEnabled = &isArchiveEnabled_
	}
	if isZero(tmp.Timeout) {
		timeout_ := 3600
		tmp.Timeout = &timeout_
	}
	return &tmp
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


func (val *FileTaskRunRequestResponse) Defaults() *FileTaskRunRequestResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchiveEnabled) {
		isArchiveEnabled_ := false
		tmp.IsArchiveEnabled = &isArchiveEnabled_
	}
	if isZero(tmp.Timeout) {
		timeout_ := 3600
		tmp.Timeout = &timeout_
	}
	return &tmp
}

type FileTaskStep struct {
	ContextAccessToken *string    `pulumi:"contextAccessToken"`
	ContextPath        *string    `pulumi:"contextPath"`
	TaskFilePath       string     `pulumi:"taskFilePath"`
	Type               string     `pulumi:"type"`
	Values             []SetValue `pulumi:"values"`
	ValuesFilePath     *string    `pulumi:"valuesFilePath"`
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

type IPRule struct {
	Action           *string `pulumi:"action"`
	IPAddressOrRange string  `pulumi:"iPAddressOrRange"`
}


func (val *IPRule) Defaults() *IPRule {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}





type IPRuleInput interface {
	pulumi.Input

	ToIPRuleOutput() IPRuleOutput
	ToIPRuleOutputWithContext(context.Context) IPRuleOutput
}

type IPRuleArgs struct {
	Action           pulumi.StringPtrInput `pulumi:"action"`
	IPAddressOrRange pulumi.StringInput    `pulumi:"iPAddressOrRange"`
}

func (IPRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRule)(nil)).Elem()
}

func (i IPRuleArgs) ToIPRuleOutput() IPRuleOutput {
	return i.ToIPRuleOutputWithContext(context.Background())
}

func (i IPRuleArgs) ToIPRuleOutputWithContext(ctx context.Context) IPRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleOutput)
}





type IPRuleArrayInput interface {
	pulumi.Input

	ToIPRuleArrayOutput() IPRuleArrayOutput
	ToIPRuleArrayOutputWithContext(context.Context) IPRuleArrayOutput
}

type IPRuleArray []IPRuleInput

func (IPRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRule)(nil)).Elem()
}

func (i IPRuleArray) ToIPRuleArrayOutput() IPRuleArrayOutput {
	return i.ToIPRuleArrayOutputWithContext(context.Background())
}

func (i IPRuleArray) ToIPRuleArrayOutputWithContext(ctx context.Context) IPRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleArrayOutput)
}

type IPRuleOutput struct{ *pulumi.OutputState }

func (IPRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRule)(nil)).Elem()
}

func (o IPRuleOutput) ToIPRuleOutput() IPRuleOutput {
	return o
}

func (o IPRuleOutput) ToIPRuleOutputWithContext(ctx context.Context) IPRuleOutput {
	return o
}

func (o IPRuleOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRule) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o IPRuleOutput) IPAddressOrRange() pulumi.StringOutput {
	return o.ApplyT(func(v IPRule) string { return v.IPAddressOrRange }).(pulumi.StringOutput)
}

type IPRuleArrayOutput struct{ *pulumi.OutputState }

func (IPRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRule)(nil)).Elem()
}

func (o IPRuleArrayOutput) ToIPRuleArrayOutput() IPRuleArrayOutput {
	return o
}

func (o IPRuleArrayOutput) ToIPRuleArrayOutputWithContext(ctx context.Context) IPRuleArrayOutput {
	return o
}

func (o IPRuleArrayOutput) Index(i pulumi.IntInput) IPRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPRule {
		return vs[0].([]IPRule)[vs[1].(int)]
	}).(IPRuleOutput)
}

type IPRuleResponse struct {
	Action           *string `pulumi:"action"`
	IPAddressOrRange string  `pulumi:"iPAddressOrRange"`
}


func (val *IPRuleResponse) Defaults() *IPRuleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}

type IPRuleResponseOutput struct{ *pulumi.OutputState }

func (IPRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRuleResponse)(nil)).Elem()
}

func (o IPRuleResponseOutput) ToIPRuleResponseOutput() IPRuleResponseOutput {
	return o
}

func (o IPRuleResponseOutput) ToIPRuleResponseOutputWithContext(ctx context.Context) IPRuleResponseOutput {
	return o
}

func (o IPRuleResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRuleResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o IPRuleResponseOutput) IPAddressOrRange() pulumi.StringOutput {
	return o.ApplyT(func(v IPRuleResponse) string { return v.IPAddressOrRange }).(pulumi.StringOutput)
}

type IPRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (IPRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRuleResponse)(nil)).Elem()
}

func (o IPRuleResponseArrayOutput) ToIPRuleResponseArrayOutput() IPRuleResponseArrayOutput {
	return o
}

func (o IPRuleResponseArrayOutput) ToIPRuleResponseArrayOutputWithContext(ctx context.Context) IPRuleResponseArrayOutput {
	return o
}

func (o IPRuleResponseArrayOutput) Index(i pulumi.IntInput) IPRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPRuleResponse {
		return vs[0].([]IPRuleResponse)[vs[1].(int)]
	}).(IPRuleResponseOutput)
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

type ImportPipelineSourceProperties struct {
	KeyVaultUri string  `pulumi:"keyVaultUri"`
	Type        *string `pulumi:"type"`
	Uri         *string `pulumi:"uri"`
}


func (val *ImportPipelineSourceProperties) Defaults() *ImportPipelineSourceProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		type_ := "AzureStorageBlobContainer"
		tmp.Type = &type_
	}
	return &tmp
}





type ImportPipelineSourcePropertiesInput interface {
	pulumi.Input

	ToImportPipelineSourcePropertiesOutput() ImportPipelineSourcePropertiesOutput
	ToImportPipelineSourcePropertiesOutputWithContext(context.Context) ImportPipelineSourcePropertiesOutput
}

type ImportPipelineSourcePropertiesArgs struct {
	KeyVaultUri pulumi.StringInput    `pulumi:"keyVaultUri"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
	Uri         pulumi.StringPtrInput `pulumi:"uri"`
}

func (ImportPipelineSourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImportPipelineSourceProperties)(nil)).Elem()
}

func (i ImportPipelineSourcePropertiesArgs) ToImportPipelineSourcePropertiesOutput() ImportPipelineSourcePropertiesOutput {
	return i.ToImportPipelineSourcePropertiesOutputWithContext(context.Background())
}

func (i ImportPipelineSourcePropertiesArgs) ToImportPipelineSourcePropertiesOutputWithContext(ctx context.Context) ImportPipelineSourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportPipelineSourcePropertiesOutput)
}

type ImportPipelineSourcePropertiesOutput struct{ *pulumi.OutputState }

func (ImportPipelineSourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImportPipelineSourceProperties)(nil)).Elem()
}

func (o ImportPipelineSourcePropertiesOutput) ToImportPipelineSourcePropertiesOutput() ImportPipelineSourcePropertiesOutput {
	return o
}

func (o ImportPipelineSourcePropertiesOutput) ToImportPipelineSourcePropertiesOutputWithContext(ctx context.Context) ImportPipelineSourcePropertiesOutput {
	return o
}

func (o ImportPipelineSourcePropertiesOutput) KeyVaultUri() pulumi.StringOutput {
	return o.ApplyT(func(v ImportPipelineSourceProperties) string { return v.KeyVaultUri }).(pulumi.StringOutput)
}

func (o ImportPipelineSourcePropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImportPipelineSourceProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ImportPipelineSourcePropertiesOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImportPipelineSourceProperties) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type ImportPipelineSourcePropertiesResponse struct {
	KeyVaultUri string  `pulumi:"keyVaultUri"`
	Type        *string `pulumi:"type"`
	Uri         *string `pulumi:"uri"`
}


func (val *ImportPipelineSourcePropertiesResponse) Defaults() *ImportPipelineSourcePropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		type_ := "AzureStorageBlobContainer"
		tmp.Type = &type_
	}
	return &tmp
}

type ImportPipelineSourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ImportPipelineSourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImportPipelineSourcePropertiesResponse)(nil)).Elem()
}

func (o ImportPipelineSourcePropertiesResponseOutput) ToImportPipelineSourcePropertiesResponseOutput() ImportPipelineSourcePropertiesResponseOutput {
	return o
}

func (o ImportPipelineSourcePropertiesResponseOutput) ToImportPipelineSourcePropertiesResponseOutputWithContext(ctx context.Context) ImportPipelineSourcePropertiesResponseOutput {
	return o
}

func (o ImportPipelineSourcePropertiesResponseOutput) KeyVaultUri() pulumi.StringOutput {
	return o.ApplyT(func(v ImportPipelineSourcePropertiesResponse) string { return v.KeyVaultUri }).(pulumi.StringOutput)
}

func (o ImportPipelineSourcePropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImportPipelineSourcePropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ImportPipelineSourcePropertiesResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImportPipelineSourcePropertiesResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type ImportPipelineSourcePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ImportPipelineSourcePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImportPipelineSourcePropertiesResponse)(nil)).Elem()
}

func (o ImportPipelineSourcePropertiesResponsePtrOutput) ToImportPipelineSourcePropertiesResponsePtrOutput() ImportPipelineSourcePropertiesResponsePtrOutput {
	return o
}

func (o ImportPipelineSourcePropertiesResponsePtrOutput) ToImportPipelineSourcePropertiesResponsePtrOutputWithContext(ctx context.Context) ImportPipelineSourcePropertiesResponsePtrOutput {
	return o
}

func (o ImportPipelineSourcePropertiesResponsePtrOutput) Elem() ImportPipelineSourcePropertiesResponseOutput {
	return o.ApplyT(func(v *ImportPipelineSourcePropertiesResponse) ImportPipelineSourcePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ImportPipelineSourcePropertiesResponse
		return ret
	}).(ImportPipelineSourcePropertiesResponseOutput)
}

func (o ImportPipelineSourcePropertiesResponsePtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImportPipelineSourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o ImportPipelineSourcePropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImportPipelineSourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ImportPipelineSourcePropertiesResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImportPipelineSourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type LoggingProperties struct {
	AuditLogStatus *string `pulumi:"auditLogStatus"`
	LogLevel       *string `pulumi:"logLevel"`
}


func (val *LoggingProperties) Defaults() *LoggingProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AuditLogStatus) {
		auditLogStatus_ := "Disabled"
		tmp.AuditLogStatus = &auditLogStatus_
	}
	if isZero(tmp.LogLevel) {
		logLevel_ := "Information"
		tmp.LogLevel = &logLevel_
	}
	return &tmp
}





type LoggingPropertiesInput interface {
	pulumi.Input

	ToLoggingPropertiesOutput() LoggingPropertiesOutput
	ToLoggingPropertiesOutputWithContext(context.Context) LoggingPropertiesOutput
}

type LoggingPropertiesArgs struct {
	AuditLogStatus pulumi.StringPtrInput `pulumi:"auditLogStatus"`
	LogLevel       pulumi.StringPtrInput `pulumi:"logLevel"`
}

func (LoggingPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoggingProperties)(nil)).Elem()
}

func (i LoggingPropertiesArgs) ToLoggingPropertiesOutput() LoggingPropertiesOutput {
	return i.ToLoggingPropertiesOutputWithContext(context.Background())
}

func (i LoggingPropertiesArgs) ToLoggingPropertiesOutputWithContext(ctx context.Context) LoggingPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingPropertiesOutput)
}

func (i LoggingPropertiesArgs) ToLoggingPropertiesPtrOutput() LoggingPropertiesPtrOutput {
	return i.ToLoggingPropertiesPtrOutputWithContext(context.Background())
}

func (i LoggingPropertiesArgs) ToLoggingPropertiesPtrOutputWithContext(ctx context.Context) LoggingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingPropertiesOutput).ToLoggingPropertiesPtrOutputWithContext(ctx)
}









type LoggingPropertiesPtrInput interface {
	pulumi.Input

	ToLoggingPropertiesPtrOutput() LoggingPropertiesPtrOutput
	ToLoggingPropertiesPtrOutputWithContext(context.Context) LoggingPropertiesPtrOutput
}

type loggingPropertiesPtrType LoggingPropertiesArgs

func LoggingPropertiesPtr(v *LoggingPropertiesArgs) LoggingPropertiesPtrInput {
	return (*loggingPropertiesPtrType)(v)
}

func (*loggingPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingProperties)(nil)).Elem()
}

func (i *loggingPropertiesPtrType) ToLoggingPropertiesPtrOutput() LoggingPropertiesPtrOutput {
	return i.ToLoggingPropertiesPtrOutputWithContext(context.Background())
}

func (i *loggingPropertiesPtrType) ToLoggingPropertiesPtrOutputWithContext(ctx context.Context) LoggingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingPropertiesPtrOutput)
}

type LoggingPropertiesOutput struct{ *pulumi.OutputState }

func (LoggingPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoggingProperties)(nil)).Elem()
}

func (o LoggingPropertiesOutput) ToLoggingPropertiesOutput() LoggingPropertiesOutput {
	return o
}

func (o LoggingPropertiesOutput) ToLoggingPropertiesOutputWithContext(ctx context.Context) LoggingPropertiesOutput {
	return o
}

func (o LoggingPropertiesOutput) ToLoggingPropertiesPtrOutput() LoggingPropertiesPtrOutput {
	return o.ToLoggingPropertiesPtrOutputWithContext(context.Background())
}

func (o LoggingPropertiesOutput) ToLoggingPropertiesPtrOutputWithContext(ctx context.Context) LoggingPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoggingProperties) *LoggingProperties {
		return &v
	}).(LoggingPropertiesPtrOutput)
}

func (o LoggingPropertiesOutput) AuditLogStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoggingProperties) *string { return v.AuditLogStatus }).(pulumi.StringPtrOutput)
}

func (o LoggingPropertiesOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoggingProperties) *string { return v.LogLevel }).(pulumi.StringPtrOutput)
}

type LoggingPropertiesPtrOutput struct{ *pulumi.OutputState }

func (LoggingPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingProperties)(nil)).Elem()
}

func (o LoggingPropertiesPtrOutput) ToLoggingPropertiesPtrOutput() LoggingPropertiesPtrOutput {
	return o
}

func (o LoggingPropertiesPtrOutput) ToLoggingPropertiesPtrOutputWithContext(ctx context.Context) LoggingPropertiesPtrOutput {
	return o
}

func (o LoggingPropertiesPtrOutput) Elem() LoggingPropertiesOutput {
	return o.ApplyT(func(v *LoggingProperties) LoggingProperties {
		if v != nil {
			return *v
		}
		var ret LoggingProperties
		return ret
	}).(LoggingPropertiesOutput)
}

func (o LoggingPropertiesPtrOutput) AuditLogStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggingProperties) *string {
		if v == nil {
			return nil
		}
		return v.AuditLogStatus
	}).(pulumi.StringPtrOutput)
}

func (o LoggingPropertiesPtrOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggingProperties) *string {
		if v == nil {
			return nil
		}
		return v.LogLevel
	}).(pulumi.StringPtrOutput)
}

type LoggingPropertiesResponse struct {
	AuditLogStatus *string `pulumi:"auditLogStatus"`
	LogLevel       *string `pulumi:"logLevel"`
}


func (val *LoggingPropertiesResponse) Defaults() *LoggingPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AuditLogStatus) {
		auditLogStatus_ := "Disabled"
		tmp.AuditLogStatus = &auditLogStatus_
	}
	if isZero(tmp.LogLevel) {
		logLevel_ := "Information"
		tmp.LogLevel = &logLevel_
	}
	return &tmp
}

type LoggingPropertiesResponseOutput struct{ *pulumi.OutputState }

func (LoggingPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoggingPropertiesResponse)(nil)).Elem()
}

func (o LoggingPropertiesResponseOutput) ToLoggingPropertiesResponseOutput() LoggingPropertiesResponseOutput {
	return o
}

func (o LoggingPropertiesResponseOutput) ToLoggingPropertiesResponseOutputWithContext(ctx context.Context) LoggingPropertiesResponseOutput {
	return o
}

func (o LoggingPropertiesResponseOutput) AuditLogStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoggingPropertiesResponse) *string { return v.AuditLogStatus }).(pulumi.StringPtrOutput)
}

func (o LoggingPropertiesResponseOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoggingPropertiesResponse) *string { return v.LogLevel }).(pulumi.StringPtrOutput)
}

type LoggingPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (LoggingPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingPropertiesResponse)(nil)).Elem()
}

func (o LoggingPropertiesResponsePtrOutput) ToLoggingPropertiesResponsePtrOutput() LoggingPropertiesResponsePtrOutput {
	return o
}

func (o LoggingPropertiesResponsePtrOutput) ToLoggingPropertiesResponsePtrOutputWithContext(ctx context.Context) LoggingPropertiesResponsePtrOutput {
	return o
}

func (o LoggingPropertiesResponsePtrOutput) Elem() LoggingPropertiesResponseOutput {
	return o.ApplyT(func(v *LoggingPropertiesResponse) LoggingPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret LoggingPropertiesResponse
		return ret
	}).(LoggingPropertiesResponseOutput)
}

func (o LoggingPropertiesResponsePtrOutput) AuditLogStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AuditLogStatus
	}).(pulumi.StringPtrOutput)
}

func (o LoggingPropertiesResponsePtrOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LogLevel
	}).(pulumi.StringPtrOutput)
}

type LoginServerPropertiesResponse struct {
	Host string                `pulumi:"host"`
	Tls  TlsPropertiesResponse `pulumi:"tls"`
}

type LoginServerPropertiesResponseOutput struct{ *pulumi.OutputState }

func (LoginServerPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoginServerPropertiesResponse)(nil)).Elem()
}

func (o LoginServerPropertiesResponseOutput) ToLoginServerPropertiesResponseOutput() LoginServerPropertiesResponseOutput {
	return o
}

func (o LoginServerPropertiesResponseOutput) ToLoginServerPropertiesResponseOutputWithContext(ctx context.Context) LoginServerPropertiesResponseOutput {
	return o
}

func (o LoginServerPropertiesResponseOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v LoginServerPropertiesResponse) string { return v.Host }).(pulumi.StringOutput)
}

func (o LoginServerPropertiesResponseOutput) Tls() TlsPropertiesResponseOutput {
	return o.ApplyT(func(v LoginServerPropertiesResponse) TlsPropertiesResponse { return v.Tls }).(TlsPropertiesResponseOutput)
}

type LoginServerPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (LoginServerPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginServerPropertiesResponse)(nil)).Elem()
}

func (o LoginServerPropertiesResponsePtrOutput) ToLoginServerPropertiesResponsePtrOutput() LoginServerPropertiesResponsePtrOutput {
	return o
}

func (o LoginServerPropertiesResponsePtrOutput) ToLoginServerPropertiesResponsePtrOutputWithContext(ctx context.Context) LoginServerPropertiesResponsePtrOutput {
	return o
}

func (o LoginServerPropertiesResponsePtrOutput) Elem() LoginServerPropertiesResponseOutput {
	return o.ApplyT(func(v *LoginServerPropertiesResponse) LoginServerPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret LoginServerPropertiesResponse
		return ret
	}).(LoginServerPropertiesResponseOutput)
}

func (o LoginServerPropertiesResponsePtrOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoginServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Host
	}).(pulumi.StringPtrOutput)
}

func (o LoginServerPropertiesResponsePtrOutput) Tls() TlsPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *LoginServerPropertiesResponse) *TlsPropertiesResponse {
		if v == nil {
			return nil
		}
		return &v.Tls
	}).(TlsPropertiesResponsePtrOutput)
}

type NetworkRuleSet struct {
	DefaultAction       string               `pulumi:"defaultAction"`
	IpRules             []IPRule             `pulumi:"ipRules"`
	VirtualNetworkRules []VirtualNetworkRule `pulumi:"virtualNetworkRules"`
}


func (val *NetworkRuleSet) Defaults() *NetworkRuleSet {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DefaultAction) {
		tmp.DefaultAction = "Allow"
	}
	return &tmp
}





type NetworkRuleSetInput interface {
	pulumi.Input

	ToNetworkRuleSetOutput() NetworkRuleSetOutput
	ToNetworkRuleSetOutputWithContext(context.Context) NetworkRuleSetOutput
}

type NetworkRuleSetArgs struct {
	DefaultAction       pulumi.StringInput           `pulumi:"defaultAction"`
	IpRules             IPRuleArrayInput             `pulumi:"ipRules"`
	VirtualNetworkRules VirtualNetworkRuleArrayInput `pulumi:"virtualNetworkRules"`
}

func (NetworkRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return i.ToNetworkRuleSetOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput)
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput).ToNetworkRuleSetPtrOutputWithContext(ctx)
}









type NetworkRuleSetPtrInput interface {
	pulumi.Input

	ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput
	ToNetworkRuleSetPtrOutputWithContext(context.Context) NetworkRuleSetPtrOutput
}

type networkRuleSetPtrType NetworkRuleSetArgs

func NetworkRuleSetPtr(v *NetworkRuleSetArgs) NetworkRuleSetPtrInput {
	return (*networkRuleSetPtrType)(v)
}

func (*networkRuleSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetPtrOutput)
}

type NetworkRuleSetOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSet) *NetworkRuleSet {
		return &v
	}).(NetworkRuleSetPtrOutput)
}

func (o NetworkRuleSetOutput) DefaultAction() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkRuleSet) string { return v.DefaultAction }).(pulumi.StringOutput)
}

func (o NetworkRuleSetOutput) IpRules() IPRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []IPRule { return v.IpRules }).(IPRuleArrayOutput)
}

func (o NetworkRuleSetOutput) VirtualNetworkRules() VirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []VirtualNetworkRule { return v.VirtualNetworkRules }).(VirtualNetworkRuleArrayOutput)
}

type NetworkRuleSetPtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) Elem() NetworkRuleSetOutput {
	return o.ApplyT(func(v *NetworkRuleSet) NetworkRuleSet {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSet
		return ret
	}).(NetworkRuleSetOutput)
}

func (o NetworkRuleSetPtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSet) *string {
		if v == nil {
			return nil
		}
		return &v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetPtrOutput) IpRules() IPRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []IPRule {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IPRuleArrayOutput)
}

func (o NetworkRuleSetPtrOutput) VirtualNetworkRules() VirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []VirtualNetworkRule {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(VirtualNetworkRuleArrayOutput)
}

type NetworkRuleSetResponse struct {
	DefaultAction       string                       `pulumi:"defaultAction"`
	IpRules             []IPRuleResponse             `pulumi:"ipRules"`
	VirtualNetworkRules []VirtualNetworkRuleResponse `pulumi:"virtualNetworkRules"`
}


func (val *NetworkRuleSetResponse) Defaults() *NetworkRuleSetResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DefaultAction) {
		tmp.DefaultAction = "Allow"
	}
	return &tmp
}

type NetworkRuleSetResponseOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutputWithContext(ctx context.Context) NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) DefaultAction() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) string { return v.DefaultAction }).(pulumi.StringOutput)
}

func (o NetworkRuleSetResponseOutput) IpRules() IPRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []IPRuleResponse { return v.IpRules }).(IPRuleResponseArrayOutput)
}

func (o NetworkRuleSetResponseOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []VirtualNetworkRuleResponse { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
}

type NetworkRuleSetResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) Elem() NetworkRuleSetResponseOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) NetworkRuleSetResponse {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSetResponse
		return ret
	}).(NetworkRuleSetResponseOutput)
}

func (o NetworkRuleSetResponsePtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponsePtrOutput) IpRules() IPRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) []IPRuleResponse {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IPRuleResponseArrayOutput)
}

func (o NetworkRuleSetResponsePtrOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) []VirtualNetworkRuleResponse {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(VirtualNetworkRuleResponseArrayOutput)
}

type OverrideTaskStepProperties struct {
	Arguments          []Argument `pulumi:"arguments"`
	ContextPath        *string    `pulumi:"contextPath"`
	File               *string    `pulumi:"file"`
	Target             *string    `pulumi:"target"`
	UpdateTriggerToken *string    `pulumi:"updateTriggerToken"`
	Values             []SetValue `pulumi:"values"`
}

type OverrideTaskStepPropertiesResponse struct {
	Arguments          []ArgumentResponse `pulumi:"arguments"`
	ContextPath        *string            `pulumi:"contextPath"`
	File               *string            `pulumi:"file"`
	Target             *string            `pulumi:"target"`
	UpdateTriggerToken *string            `pulumi:"updateTriggerToken"`
	Values             []SetValueResponse `pulumi:"values"`
}

type ParentProperties struct {
	Id             *string        `pulumi:"id"`
	SyncProperties SyncProperties `pulumi:"syncProperties"`
}





type ParentPropertiesInput interface {
	pulumi.Input

	ToParentPropertiesOutput() ParentPropertiesOutput
	ToParentPropertiesOutputWithContext(context.Context) ParentPropertiesOutput
}

type ParentPropertiesArgs struct {
	Id             pulumi.StringPtrInput `pulumi:"id"`
	SyncProperties SyncPropertiesInput   `pulumi:"syncProperties"`
}

func (ParentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParentProperties)(nil)).Elem()
}

func (i ParentPropertiesArgs) ToParentPropertiesOutput() ParentPropertiesOutput {
	return i.ToParentPropertiesOutputWithContext(context.Background())
}

func (i ParentPropertiesArgs) ToParentPropertiesOutputWithContext(ctx context.Context) ParentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParentPropertiesOutput)
}

type ParentPropertiesOutput struct{ *pulumi.OutputState }

func (ParentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParentProperties)(nil)).Elem()
}

func (o ParentPropertiesOutput) ToParentPropertiesOutput() ParentPropertiesOutput {
	return o
}

func (o ParentPropertiesOutput) ToParentPropertiesOutputWithContext(ctx context.Context) ParentPropertiesOutput {
	return o
}

func (o ParentPropertiesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParentProperties) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ParentPropertiesOutput) SyncProperties() SyncPropertiesOutput {
	return o.ApplyT(func(v ParentProperties) SyncProperties { return v.SyncProperties }).(SyncPropertiesOutput)
}

type ParentPropertiesResponse struct {
	Id             *string                `pulumi:"id"`
	SyncProperties SyncPropertiesResponse `pulumi:"syncProperties"`
}

type ParentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ParentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParentPropertiesResponse)(nil)).Elem()
}

func (o ParentPropertiesResponseOutput) ToParentPropertiesResponseOutput() ParentPropertiesResponseOutput {
	return o
}

func (o ParentPropertiesResponseOutput) ToParentPropertiesResponseOutputWithContext(ctx context.Context) ParentPropertiesResponseOutput {
	return o
}

func (o ParentPropertiesResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParentPropertiesResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ParentPropertiesResponseOutput) SyncProperties() SyncPropertiesResponseOutput {
	return o.ApplyT(func(v ParentPropertiesResponse) SyncPropertiesResponse { return v.SyncProperties }).(SyncPropertiesResponseOutput)
}

type PipelineRunRequest struct {
	Artifacts          []string                     `pulumi:"artifacts"`
	CatalogDigest      *string                      `pulumi:"catalogDigest"`
	PipelineResourceId *string                      `pulumi:"pipelineResourceId"`
	Source             *PipelineRunSourceProperties `pulumi:"source"`
	Target             *PipelineRunTargetProperties `pulumi:"target"`
}


func (val *PipelineRunRequest) Defaults() *PipelineRunRequest {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Source = tmp.Source.Defaults()

	tmp.Target = tmp.Target.Defaults()

	return &tmp
}





type PipelineRunRequestInput interface {
	pulumi.Input

	ToPipelineRunRequestOutput() PipelineRunRequestOutput
	ToPipelineRunRequestOutputWithContext(context.Context) PipelineRunRequestOutput
}

type PipelineRunRequestArgs struct {
	Artifacts          pulumi.StringArrayInput             `pulumi:"artifacts"`
	CatalogDigest      pulumi.StringPtrInput               `pulumi:"catalogDigest"`
	PipelineResourceId pulumi.StringPtrInput               `pulumi:"pipelineResourceId"`
	Source             PipelineRunSourcePropertiesPtrInput `pulumi:"source"`
	Target             PipelineRunTargetPropertiesPtrInput `pulumi:"target"`
}

func (PipelineRunRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunRequest)(nil)).Elem()
}

func (i PipelineRunRequestArgs) ToPipelineRunRequestOutput() PipelineRunRequestOutput {
	return i.ToPipelineRunRequestOutputWithContext(context.Background())
}

func (i PipelineRunRequestArgs) ToPipelineRunRequestOutputWithContext(ctx context.Context) PipelineRunRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunRequestOutput)
}

func (i PipelineRunRequestArgs) ToPipelineRunRequestPtrOutput() PipelineRunRequestPtrOutput {
	return i.ToPipelineRunRequestPtrOutputWithContext(context.Background())
}

func (i PipelineRunRequestArgs) ToPipelineRunRequestPtrOutputWithContext(ctx context.Context) PipelineRunRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunRequestOutput).ToPipelineRunRequestPtrOutputWithContext(ctx)
}









type PipelineRunRequestPtrInput interface {
	pulumi.Input

	ToPipelineRunRequestPtrOutput() PipelineRunRequestPtrOutput
	ToPipelineRunRequestPtrOutputWithContext(context.Context) PipelineRunRequestPtrOutput
}

type pipelineRunRequestPtrType PipelineRunRequestArgs

func PipelineRunRequestPtr(v *PipelineRunRequestArgs) PipelineRunRequestPtrInput {
	return (*pipelineRunRequestPtrType)(v)
}

func (*pipelineRunRequestPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunRequest)(nil)).Elem()
}

func (i *pipelineRunRequestPtrType) ToPipelineRunRequestPtrOutput() PipelineRunRequestPtrOutput {
	return i.ToPipelineRunRequestPtrOutputWithContext(context.Background())
}

func (i *pipelineRunRequestPtrType) ToPipelineRunRequestPtrOutputWithContext(ctx context.Context) PipelineRunRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunRequestPtrOutput)
}

type PipelineRunRequestOutput struct{ *pulumi.OutputState }

func (PipelineRunRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunRequest)(nil)).Elem()
}

func (o PipelineRunRequestOutput) ToPipelineRunRequestOutput() PipelineRunRequestOutput {
	return o
}

func (o PipelineRunRequestOutput) ToPipelineRunRequestOutputWithContext(ctx context.Context) PipelineRunRequestOutput {
	return o
}

func (o PipelineRunRequestOutput) ToPipelineRunRequestPtrOutput() PipelineRunRequestPtrOutput {
	return o.ToPipelineRunRequestPtrOutputWithContext(context.Background())
}

func (o PipelineRunRequestOutput) ToPipelineRunRequestPtrOutputWithContext(ctx context.Context) PipelineRunRequestPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineRunRequest) *PipelineRunRequest {
		return &v
	}).(PipelineRunRequestPtrOutput)
}

func (o PipelineRunRequestOutput) Artifacts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PipelineRunRequest) []string { return v.Artifacts }).(pulumi.StringArrayOutput)
}

func (o PipelineRunRequestOutput) CatalogDigest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunRequest) *string { return v.CatalogDigest }).(pulumi.StringPtrOutput)
}

func (o PipelineRunRequestOutput) PipelineResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunRequest) *string { return v.PipelineResourceId }).(pulumi.StringPtrOutput)
}

func (o PipelineRunRequestOutput) Source() PipelineRunSourcePropertiesPtrOutput {
	return o.ApplyT(func(v PipelineRunRequest) *PipelineRunSourceProperties { return v.Source }).(PipelineRunSourcePropertiesPtrOutput)
}

func (o PipelineRunRequestOutput) Target() PipelineRunTargetPropertiesPtrOutput {
	return o.ApplyT(func(v PipelineRunRequest) *PipelineRunTargetProperties { return v.Target }).(PipelineRunTargetPropertiesPtrOutput)
}

type PipelineRunRequestPtrOutput struct{ *pulumi.OutputState }

func (PipelineRunRequestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunRequest)(nil)).Elem()
}

func (o PipelineRunRequestPtrOutput) ToPipelineRunRequestPtrOutput() PipelineRunRequestPtrOutput {
	return o
}

func (o PipelineRunRequestPtrOutput) ToPipelineRunRequestPtrOutputWithContext(ctx context.Context) PipelineRunRequestPtrOutput {
	return o
}

func (o PipelineRunRequestPtrOutput) Elem() PipelineRunRequestOutput {
	return o.ApplyT(func(v *PipelineRunRequest) PipelineRunRequest {
		if v != nil {
			return *v
		}
		var ret PipelineRunRequest
		return ret
	}).(PipelineRunRequestOutput)
}

func (o PipelineRunRequestPtrOutput) Artifacts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PipelineRunRequest) []string {
		if v == nil {
			return nil
		}
		return v.Artifacts
	}).(pulumi.StringArrayOutput)
}

func (o PipelineRunRequestPtrOutput) CatalogDigest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunRequest) *string {
		if v == nil {
			return nil
		}
		return v.CatalogDigest
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunRequestPtrOutput) PipelineResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunRequest) *string {
		if v == nil {
			return nil
		}
		return v.PipelineResourceId
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunRequestPtrOutput) Source() PipelineRunSourcePropertiesPtrOutput {
	return o.ApplyT(func(v *PipelineRunRequest) *PipelineRunSourceProperties {
		if v == nil {
			return nil
		}
		return v.Source
	}).(PipelineRunSourcePropertiesPtrOutput)
}

func (o PipelineRunRequestPtrOutput) Target() PipelineRunTargetPropertiesPtrOutput {
	return o.ApplyT(func(v *PipelineRunRequest) *PipelineRunTargetProperties {
		if v == nil {
			return nil
		}
		return v.Target
	}).(PipelineRunTargetPropertiesPtrOutput)
}

type PipelineRunRequestResponse struct {
	Artifacts          []string                             `pulumi:"artifacts"`
	CatalogDigest      *string                              `pulumi:"catalogDigest"`
	PipelineResourceId *string                              `pulumi:"pipelineResourceId"`
	Source             *PipelineRunSourcePropertiesResponse `pulumi:"source"`
	Target             *PipelineRunTargetPropertiesResponse `pulumi:"target"`
}


func (val *PipelineRunRequestResponse) Defaults() *PipelineRunRequestResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Source = tmp.Source.Defaults()

	tmp.Target = tmp.Target.Defaults()

	return &tmp
}

type PipelineRunRequestResponseOutput struct{ *pulumi.OutputState }

func (PipelineRunRequestResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunRequestResponse)(nil)).Elem()
}

func (o PipelineRunRequestResponseOutput) ToPipelineRunRequestResponseOutput() PipelineRunRequestResponseOutput {
	return o
}

func (o PipelineRunRequestResponseOutput) ToPipelineRunRequestResponseOutputWithContext(ctx context.Context) PipelineRunRequestResponseOutput {
	return o
}

func (o PipelineRunRequestResponseOutput) Artifacts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PipelineRunRequestResponse) []string { return v.Artifacts }).(pulumi.StringArrayOutput)
}

func (o PipelineRunRequestResponseOutput) CatalogDigest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunRequestResponse) *string { return v.CatalogDigest }).(pulumi.StringPtrOutput)
}

func (o PipelineRunRequestResponseOutput) PipelineResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunRequestResponse) *string { return v.PipelineResourceId }).(pulumi.StringPtrOutput)
}

func (o PipelineRunRequestResponseOutput) Source() PipelineRunSourcePropertiesResponsePtrOutput {
	return o.ApplyT(func(v PipelineRunRequestResponse) *PipelineRunSourcePropertiesResponse { return v.Source }).(PipelineRunSourcePropertiesResponsePtrOutput)
}

func (o PipelineRunRequestResponseOutput) Target() PipelineRunTargetPropertiesResponsePtrOutput {
	return o.ApplyT(func(v PipelineRunRequestResponse) *PipelineRunTargetPropertiesResponse { return v.Target }).(PipelineRunTargetPropertiesResponsePtrOutput)
}

type PipelineRunRequestResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelineRunRequestResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunRequestResponse)(nil)).Elem()
}

func (o PipelineRunRequestResponsePtrOutput) ToPipelineRunRequestResponsePtrOutput() PipelineRunRequestResponsePtrOutput {
	return o
}

func (o PipelineRunRequestResponsePtrOutput) ToPipelineRunRequestResponsePtrOutputWithContext(ctx context.Context) PipelineRunRequestResponsePtrOutput {
	return o
}

func (o PipelineRunRequestResponsePtrOutput) Elem() PipelineRunRequestResponseOutput {
	return o.ApplyT(func(v *PipelineRunRequestResponse) PipelineRunRequestResponse {
		if v != nil {
			return *v
		}
		var ret PipelineRunRequestResponse
		return ret
	}).(PipelineRunRequestResponseOutput)
}

func (o PipelineRunRequestResponsePtrOutput) Artifacts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PipelineRunRequestResponse) []string {
		if v == nil {
			return nil
		}
		return v.Artifacts
	}).(pulumi.StringArrayOutput)
}

func (o PipelineRunRequestResponsePtrOutput) CatalogDigest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunRequestResponse) *string {
		if v == nil {
			return nil
		}
		return v.CatalogDigest
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunRequestResponsePtrOutput) PipelineResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunRequestResponse) *string {
		if v == nil {
			return nil
		}
		return v.PipelineResourceId
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunRequestResponsePtrOutput) Source() PipelineRunSourcePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *PipelineRunRequestResponse) *PipelineRunSourcePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Source
	}).(PipelineRunSourcePropertiesResponsePtrOutput)
}

func (o PipelineRunRequestResponsePtrOutput) Target() PipelineRunTargetPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *PipelineRunRequestResponse) *PipelineRunTargetPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Target
	}).(PipelineRunTargetPropertiesResponsePtrOutput)
}

type PipelineRunResponseResponse struct {
	CatalogDigest           *string                                 `pulumi:"catalogDigest"`
	FinishTime              *string                                 `pulumi:"finishTime"`
	ImportedArtifacts       []string                                `pulumi:"importedArtifacts"`
	PipelineRunErrorMessage *string                                 `pulumi:"pipelineRunErrorMessage"`
	Progress                *ProgressPropertiesResponse             `pulumi:"progress"`
	Source                  *ImportPipelineSourcePropertiesResponse `pulumi:"source"`
	StartTime               *string                                 `pulumi:"startTime"`
	Status                  *string                                 `pulumi:"status"`
	Target                  *ExportPipelineTargetPropertiesResponse `pulumi:"target"`
	Trigger                 *PipelineTriggerDescriptorResponse      `pulumi:"trigger"`
}


func (val *PipelineRunResponseResponse) Defaults() *PipelineRunResponseResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Source = tmp.Source.Defaults()

	return &tmp
}

type PipelineRunResponseResponseOutput struct{ *pulumi.OutputState }

func (PipelineRunResponseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunResponseResponse)(nil)).Elem()
}

func (o PipelineRunResponseResponseOutput) ToPipelineRunResponseResponseOutput() PipelineRunResponseResponseOutput {
	return o
}

func (o PipelineRunResponseResponseOutput) ToPipelineRunResponseResponseOutputWithContext(ctx context.Context) PipelineRunResponseResponseOutput {
	return o
}

func (o PipelineRunResponseResponseOutput) CatalogDigest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunResponseResponse) *string { return v.CatalogDigest }).(pulumi.StringPtrOutput)
}

func (o PipelineRunResponseResponseOutput) FinishTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunResponseResponse) *string { return v.FinishTime }).(pulumi.StringPtrOutput)
}

func (o PipelineRunResponseResponseOutput) ImportedArtifacts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PipelineRunResponseResponse) []string { return v.ImportedArtifacts }).(pulumi.StringArrayOutput)
}

func (o PipelineRunResponseResponseOutput) PipelineRunErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunResponseResponse) *string { return v.PipelineRunErrorMessage }).(pulumi.StringPtrOutput)
}

func (o PipelineRunResponseResponseOutput) Progress() ProgressPropertiesResponsePtrOutput {
	return o.ApplyT(func(v PipelineRunResponseResponse) *ProgressPropertiesResponse { return v.Progress }).(ProgressPropertiesResponsePtrOutput)
}

func (o PipelineRunResponseResponseOutput) Source() ImportPipelineSourcePropertiesResponsePtrOutput {
	return o.ApplyT(func(v PipelineRunResponseResponse) *ImportPipelineSourcePropertiesResponse { return v.Source }).(ImportPipelineSourcePropertiesResponsePtrOutput)
}

func (o PipelineRunResponseResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunResponseResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o PipelineRunResponseResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunResponseResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o PipelineRunResponseResponseOutput) Target() ExportPipelineTargetPropertiesResponsePtrOutput {
	return o.ApplyT(func(v PipelineRunResponseResponse) *ExportPipelineTargetPropertiesResponse { return v.Target }).(ExportPipelineTargetPropertiesResponsePtrOutput)
}

func (o PipelineRunResponseResponseOutput) Trigger() PipelineTriggerDescriptorResponsePtrOutput {
	return o.ApplyT(func(v PipelineRunResponseResponse) *PipelineTriggerDescriptorResponse { return v.Trigger }).(PipelineTriggerDescriptorResponsePtrOutput)
}

type PipelineRunSourceProperties struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}


func (val *PipelineRunSourceProperties) Defaults() *PipelineRunSourceProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		type_ := "AzureStorageBlob"
		tmp.Type = &type_
	}
	return &tmp
}





type PipelineRunSourcePropertiesInput interface {
	pulumi.Input

	ToPipelineRunSourcePropertiesOutput() PipelineRunSourcePropertiesOutput
	ToPipelineRunSourcePropertiesOutputWithContext(context.Context) PipelineRunSourcePropertiesOutput
}

type PipelineRunSourcePropertiesArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (PipelineRunSourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunSourceProperties)(nil)).Elem()
}

func (i PipelineRunSourcePropertiesArgs) ToPipelineRunSourcePropertiesOutput() PipelineRunSourcePropertiesOutput {
	return i.ToPipelineRunSourcePropertiesOutputWithContext(context.Background())
}

func (i PipelineRunSourcePropertiesArgs) ToPipelineRunSourcePropertiesOutputWithContext(ctx context.Context) PipelineRunSourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunSourcePropertiesOutput)
}

func (i PipelineRunSourcePropertiesArgs) ToPipelineRunSourcePropertiesPtrOutput() PipelineRunSourcePropertiesPtrOutput {
	return i.ToPipelineRunSourcePropertiesPtrOutputWithContext(context.Background())
}

func (i PipelineRunSourcePropertiesArgs) ToPipelineRunSourcePropertiesPtrOutputWithContext(ctx context.Context) PipelineRunSourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunSourcePropertiesOutput).ToPipelineRunSourcePropertiesPtrOutputWithContext(ctx)
}









type PipelineRunSourcePropertiesPtrInput interface {
	pulumi.Input

	ToPipelineRunSourcePropertiesPtrOutput() PipelineRunSourcePropertiesPtrOutput
	ToPipelineRunSourcePropertiesPtrOutputWithContext(context.Context) PipelineRunSourcePropertiesPtrOutput
}

type pipelineRunSourcePropertiesPtrType PipelineRunSourcePropertiesArgs

func PipelineRunSourcePropertiesPtr(v *PipelineRunSourcePropertiesArgs) PipelineRunSourcePropertiesPtrInput {
	return (*pipelineRunSourcePropertiesPtrType)(v)
}

func (*pipelineRunSourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunSourceProperties)(nil)).Elem()
}

func (i *pipelineRunSourcePropertiesPtrType) ToPipelineRunSourcePropertiesPtrOutput() PipelineRunSourcePropertiesPtrOutput {
	return i.ToPipelineRunSourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *pipelineRunSourcePropertiesPtrType) ToPipelineRunSourcePropertiesPtrOutputWithContext(ctx context.Context) PipelineRunSourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunSourcePropertiesPtrOutput)
}

type PipelineRunSourcePropertiesOutput struct{ *pulumi.OutputState }

func (PipelineRunSourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunSourceProperties)(nil)).Elem()
}

func (o PipelineRunSourcePropertiesOutput) ToPipelineRunSourcePropertiesOutput() PipelineRunSourcePropertiesOutput {
	return o
}

func (o PipelineRunSourcePropertiesOutput) ToPipelineRunSourcePropertiesOutputWithContext(ctx context.Context) PipelineRunSourcePropertiesOutput {
	return o
}

func (o PipelineRunSourcePropertiesOutput) ToPipelineRunSourcePropertiesPtrOutput() PipelineRunSourcePropertiesPtrOutput {
	return o.ToPipelineRunSourcePropertiesPtrOutputWithContext(context.Background())
}

func (o PipelineRunSourcePropertiesOutput) ToPipelineRunSourcePropertiesPtrOutputWithContext(ctx context.Context) PipelineRunSourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineRunSourceProperties) *PipelineRunSourceProperties {
		return &v
	}).(PipelineRunSourcePropertiesPtrOutput)
}

func (o PipelineRunSourcePropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunSourceProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PipelineRunSourcePropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunSourceProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PipelineRunSourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (PipelineRunSourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunSourceProperties)(nil)).Elem()
}

func (o PipelineRunSourcePropertiesPtrOutput) ToPipelineRunSourcePropertiesPtrOutput() PipelineRunSourcePropertiesPtrOutput {
	return o
}

func (o PipelineRunSourcePropertiesPtrOutput) ToPipelineRunSourcePropertiesPtrOutputWithContext(ctx context.Context) PipelineRunSourcePropertiesPtrOutput {
	return o
}

func (o PipelineRunSourcePropertiesPtrOutput) Elem() PipelineRunSourcePropertiesOutput {
	return o.ApplyT(func(v *PipelineRunSourceProperties) PipelineRunSourceProperties {
		if v != nil {
			return *v
		}
		var ret PipelineRunSourceProperties
		return ret
	}).(PipelineRunSourcePropertiesOutput)
}

func (o PipelineRunSourcePropertiesPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunSourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunSourcePropertiesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunSourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type PipelineRunSourcePropertiesResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}


func (val *PipelineRunSourcePropertiesResponse) Defaults() *PipelineRunSourcePropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		type_ := "AzureStorageBlob"
		tmp.Type = &type_
	}
	return &tmp
}

type PipelineRunSourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (PipelineRunSourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunSourcePropertiesResponse)(nil)).Elem()
}

func (o PipelineRunSourcePropertiesResponseOutput) ToPipelineRunSourcePropertiesResponseOutput() PipelineRunSourcePropertiesResponseOutput {
	return o
}

func (o PipelineRunSourcePropertiesResponseOutput) ToPipelineRunSourcePropertiesResponseOutputWithContext(ctx context.Context) PipelineRunSourcePropertiesResponseOutput {
	return o
}

func (o PipelineRunSourcePropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunSourcePropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PipelineRunSourcePropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunSourcePropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PipelineRunSourcePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelineRunSourcePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunSourcePropertiesResponse)(nil)).Elem()
}

func (o PipelineRunSourcePropertiesResponsePtrOutput) ToPipelineRunSourcePropertiesResponsePtrOutput() PipelineRunSourcePropertiesResponsePtrOutput {
	return o
}

func (o PipelineRunSourcePropertiesResponsePtrOutput) ToPipelineRunSourcePropertiesResponsePtrOutputWithContext(ctx context.Context) PipelineRunSourcePropertiesResponsePtrOutput {
	return o
}

func (o PipelineRunSourcePropertiesResponsePtrOutput) Elem() PipelineRunSourcePropertiesResponseOutput {
	return o.ApplyT(func(v *PipelineRunSourcePropertiesResponse) PipelineRunSourcePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PipelineRunSourcePropertiesResponse
		return ret
	}).(PipelineRunSourcePropertiesResponseOutput)
}

func (o PipelineRunSourcePropertiesResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunSourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunSourcePropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunSourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type PipelineRunTargetProperties struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}


func (val *PipelineRunTargetProperties) Defaults() *PipelineRunTargetProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		type_ := "AzureStorageBlob"
		tmp.Type = &type_
	}
	return &tmp
}





type PipelineRunTargetPropertiesInput interface {
	pulumi.Input

	ToPipelineRunTargetPropertiesOutput() PipelineRunTargetPropertiesOutput
	ToPipelineRunTargetPropertiesOutputWithContext(context.Context) PipelineRunTargetPropertiesOutput
}

type PipelineRunTargetPropertiesArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (PipelineRunTargetPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunTargetProperties)(nil)).Elem()
}

func (i PipelineRunTargetPropertiesArgs) ToPipelineRunTargetPropertiesOutput() PipelineRunTargetPropertiesOutput {
	return i.ToPipelineRunTargetPropertiesOutputWithContext(context.Background())
}

func (i PipelineRunTargetPropertiesArgs) ToPipelineRunTargetPropertiesOutputWithContext(ctx context.Context) PipelineRunTargetPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunTargetPropertiesOutput)
}

func (i PipelineRunTargetPropertiesArgs) ToPipelineRunTargetPropertiesPtrOutput() PipelineRunTargetPropertiesPtrOutput {
	return i.ToPipelineRunTargetPropertiesPtrOutputWithContext(context.Background())
}

func (i PipelineRunTargetPropertiesArgs) ToPipelineRunTargetPropertiesPtrOutputWithContext(ctx context.Context) PipelineRunTargetPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunTargetPropertiesOutput).ToPipelineRunTargetPropertiesPtrOutputWithContext(ctx)
}









type PipelineRunTargetPropertiesPtrInput interface {
	pulumi.Input

	ToPipelineRunTargetPropertiesPtrOutput() PipelineRunTargetPropertiesPtrOutput
	ToPipelineRunTargetPropertiesPtrOutputWithContext(context.Context) PipelineRunTargetPropertiesPtrOutput
}

type pipelineRunTargetPropertiesPtrType PipelineRunTargetPropertiesArgs

func PipelineRunTargetPropertiesPtr(v *PipelineRunTargetPropertiesArgs) PipelineRunTargetPropertiesPtrInput {
	return (*pipelineRunTargetPropertiesPtrType)(v)
}

func (*pipelineRunTargetPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunTargetProperties)(nil)).Elem()
}

func (i *pipelineRunTargetPropertiesPtrType) ToPipelineRunTargetPropertiesPtrOutput() PipelineRunTargetPropertiesPtrOutput {
	return i.ToPipelineRunTargetPropertiesPtrOutputWithContext(context.Background())
}

func (i *pipelineRunTargetPropertiesPtrType) ToPipelineRunTargetPropertiesPtrOutputWithContext(ctx context.Context) PipelineRunTargetPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunTargetPropertiesPtrOutput)
}

type PipelineRunTargetPropertiesOutput struct{ *pulumi.OutputState }

func (PipelineRunTargetPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunTargetProperties)(nil)).Elem()
}

func (o PipelineRunTargetPropertiesOutput) ToPipelineRunTargetPropertiesOutput() PipelineRunTargetPropertiesOutput {
	return o
}

func (o PipelineRunTargetPropertiesOutput) ToPipelineRunTargetPropertiesOutputWithContext(ctx context.Context) PipelineRunTargetPropertiesOutput {
	return o
}

func (o PipelineRunTargetPropertiesOutput) ToPipelineRunTargetPropertiesPtrOutput() PipelineRunTargetPropertiesPtrOutput {
	return o.ToPipelineRunTargetPropertiesPtrOutputWithContext(context.Background())
}

func (o PipelineRunTargetPropertiesOutput) ToPipelineRunTargetPropertiesPtrOutputWithContext(ctx context.Context) PipelineRunTargetPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineRunTargetProperties) *PipelineRunTargetProperties {
		return &v
	}).(PipelineRunTargetPropertiesPtrOutput)
}

func (o PipelineRunTargetPropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunTargetProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PipelineRunTargetPropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunTargetProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PipelineRunTargetPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PipelineRunTargetPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunTargetProperties)(nil)).Elem()
}

func (o PipelineRunTargetPropertiesPtrOutput) ToPipelineRunTargetPropertiesPtrOutput() PipelineRunTargetPropertiesPtrOutput {
	return o
}

func (o PipelineRunTargetPropertiesPtrOutput) ToPipelineRunTargetPropertiesPtrOutputWithContext(ctx context.Context) PipelineRunTargetPropertiesPtrOutput {
	return o
}

func (o PipelineRunTargetPropertiesPtrOutput) Elem() PipelineRunTargetPropertiesOutput {
	return o.ApplyT(func(v *PipelineRunTargetProperties) PipelineRunTargetProperties {
		if v != nil {
			return *v
		}
		var ret PipelineRunTargetProperties
		return ret
	}).(PipelineRunTargetPropertiesOutput)
}

func (o PipelineRunTargetPropertiesPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunTargetProperties) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunTargetPropertiesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunTargetProperties) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type PipelineRunTargetPropertiesResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}


func (val *PipelineRunTargetPropertiesResponse) Defaults() *PipelineRunTargetPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		type_ := "AzureStorageBlob"
		tmp.Type = &type_
	}
	return &tmp
}

type PipelineRunTargetPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PipelineRunTargetPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunTargetPropertiesResponse)(nil)).Elem()
}

func (o PipelineRunTargetPropertiesResponseOutput) ToPipelineRunTargetPropertiesResponseOutput() PipelineRunTargetPropertiesResponseOutput {
	return o
}

func (o PipelineRunTargetPropertiesResponseOutput) ToPipelineRunTargetPropertiesResponseOutputWithContext(ctx context.Context) PipelineRunTargetPropertiesResponseOutput {
	return o
}

func (o PipelineRunTargetPropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunTargetPropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PipelineRunTargetPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunTargetPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PipelineRunTargetPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelineRunTargetPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunTargetPropertiesResponse)(nil)).Elem()
}

func (o PipelineRunTargetPropertiesResponsePtrOutput) ToPipelineRunTargetPropertiesResponsePtrOutput() PipelineRunTargetPropertiesResponsePtrOutput {
	return o
}

func (o PipelineRunTargetPropertiesResponsePtrOutput) ToPipelineRunTargetPropertiesResponsePtrOutputWithContext(ctx context.Context) PipelineRunTargetPropertiesResponsePtrOutput {
	return o
}

func (o PipelineRunTargetPropertiesResponsePtrOutput) Elem() PipelineRunTargetPropertiesResponseOutput {
	return o.ApplyT(func(v *PipelineRunTargetPropertiesResponse) PipelineRunTargetPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PipelineRunTargetPropertiesResponse
		return ret
	}).(PipelineRunTargetPropertiesResponseOutput)
}

func (o PipelineRunTargetPropertiesResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunTargetPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunTargetPropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunTargetPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type PipelineSourceTriggerDescriptorResponse struct {
	Timestamp *string `pulumi:"timestamp"`
}

type PipelineSourceTriggerDescriptorResponseOutput struct{ *pulumi.OutputState }

func (PipelineSourceTriggerDescriptorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineSourceTriggerDescriptorResponse)(nil)).Elem()
}

func (o PipelineSourceTriggerDescriptorResponseOutput) ToPipelineSourceTriggerDescriptorResponseOutput() PipelineSourceTriggerDescriptorResponseOutput {
	return o
}

func (o PipelineSourceTriggerDescriptorResponseOutput) ToPipelineSourceTriggerDescriptorResponseOutputWithContext(ctx context.Context) PipelineSourceTriggerDescriptorResponseOutput {
	return o
}

func (o PipelineSourceTriggerDescriptorResponseOutput) Timestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineSourceTriggerDescriptorResponse) *string { return v.Timestamp }).(pulumi.StringPtrOutput)
}

type PipelineSourceTriggerDescriptorResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelineSourceTriggerDescriptorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineSourceTriggerDescriptorResponse)(nil)).Elem()
}

func (o PipelineSourceTriggerDescriptorResponsePtrOutput) ToPipelineSourceTriggerDescriptorResponsePtrOutput() PipelineSourceTriggerDescriptorResponsePtrOutput {
	return o
}

func (o PipelineSourceTriggerDescriptorResponsePtrOutput) ToPipelineSourceTriggerDescriptorResponsePtrOutputWithContext(ctx context.Context) PipelineSourceTriggerDescriptorResponsePtrOutput {
	return o
}

func (o PipelineSourceTriggerDescriptorResponsePtrOutput) Elem() PipelineSourceTriggerDescriptorResponseOutput {
	return o.ApplyT(func(v *PipelineSourceTriggerDescriptorResponse) PipelineSourceTriggerDescriptorResponse {
		if v != nil {
			return *v
		}
		var ret PipelineSourceTriggerDescriptorResponse
		return ret
	}).(PipelineSourceTriggerDescriptorResponseOutput)
}

func (o PipelineSourceTriggerDescriptorResponsePtrOutput) Timestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineSourceTriggerDescriptorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Timestamp
	}).(pulumi.StringPtrOutput)
}

type PipelineSourceTriggerProperties struct {
	Status string `pulumi:"status"`
}


func (val *PipelineSourceTriggerProperties) Defaults() *PipelineSourceTriggerProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		tmp.Status = "Enabled"
	}
	return &tmp
}





type PipelineSourceTriggerPropertiesInput interface {
	pulumi.Input

	ToPipelineSourceTriggerPropertiesOutput() PipelineSourceTriggerPropertiesOutput
	ToPipelineSourceTriggerPropertiesOutputWithContext(context.Context) PipelineSourceTriggerPropertiesOutput
}

type PipelineSourceTriggerPropertiesArgs struct {
	Status pulumi.StringInput `pulumi:"status"`
}

func (PipelineSourceTriggerPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineSourceTriggerProperties)(nil)).Elem()
}

func (i PipelineSourceTriggerPropertiesArgs) ToPipelineSourceTriggerPropertiesOutput() PipelineSourceTriggerPropertiesOutput {
	return i.ToPipelineSourceTriggerPropertiesOutputWithContext(context.Background())
}

func (i PipelineSourceTriggerPropertiesArgs) ToPipelineSourceTriggerPropertiesOutputWithContext(ctx context.Context) PipelineSourceTriggerPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineSourceTriggerPropertiesOutput)
}

func (i PipelineSourceTriggerPropertiesArgs) ToPipelineSourceTriggerPropertiesPtrOutput() PipelineSourceTriggerPropertiesPtrOutput {
	return i.ToPipelineSourceTriggerPropertiesPtrOutputWithContext(context.Background())
}

func (i PipelineSourceTriggerPropertiesArgs) ToPipelineSourceTriggerPropertiesPtrOutputWithContext(ctx context.Context) PipelineSourceTriggerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineSourceTriggerPropertiesOutput).ToPipelineSourceTriggerPropertiesPtrOutputWithContext(ctx)
}









type PipelineSourceTriggerPropertiesPtrInput interface {
	pulumi.Input

	ToPipelineSourceTriggerPropertiesPtrOutput() PipelineSourceTriggerPropertiesPtrOutput
	ToPipelineSourceTriggerPropertiesPtrOutputWithContext(context.Context) PipelineSourceTriggerPropertiesPtrOutput
}

type pipelineSourceTriggerPropertiesPtrType PipelineSourceTriggerPropertiesArgs

func PipelineSourceTriggerPropertiesPtr(v *PipelineSourceTriggerPropertiesArgs) PipelineSourceTriggerPropertiesPtrInput {
	return (*pipelineSourceTriggerPropertiesPtrType)(v)
}

func (*pipelineSourceTriggerPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineSourceTriggerProperties)(nil)).Elem()
}

func (i *pipelineSourceTriggerPropertiesPtrType) ToPipelineSourceTriggerPropertiesPtrOutput() PipelineSourceTriggerPropertiesPtrOutput {
	return i.ToPipelineSourceTriggerPropertiesPtrOutputWithContext(context.Background())
}

func (i *pipelineSourceTriggerPropertiesPtrType) ToPipelineSourceTriggerPropertiesPtrOutputWithContext(ctx context.Context) PipelineSourceTriggerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineSourceTriggerPropertiesPtrOutput)
}

type PipelineSourceTriggerPropertiesOutput struct{ *pulumi.OutputState }

func (PipelineSourceTriggerPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineSourceTriggerProperties)(nil)).Elem()
}

func (o PipelineSourceTriggerPropertiesOutput) ToPipelineSourceTriggerPropertiesOutput() PipelineSourceTriggerPropertiesOutput {
	return o
}

func (o PipelineSourceTriggerPropertiesOutput) ToPipelineSourceTriggerPropertiesOutputWithContext(ctx context.Context) PipelineSourceTriggerPropertiesOutput {
	return o
}

func (o PipelineSourceTriggerPropertiesOutput) ToPipelineSourceTriggerPropertiesPtrOutput() PipelineSourceTriggerPropertiesPtrOutput {
	return o.ToPipelineSourceTriggerPropertiesPtrOutputWithContext(context.Background())
}

func (o PipelineSourceTriggerPropertiesOutput) ToPipelineSourceTriggerPropertiesPtrOutputWithContext(ctx context.Context) PipelineSourceTriggerPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineSourceTriggerProperties) *PipelineSourceTriggerProperties {
		return &v
	}).(PipelineSourceTriggerPropertiesPtrOutput)
}

func (o PipelineSourceTriggerPropertiesOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PipelineSourceTriggerProperties) string { return v.Status }).(pulumi.StringOutput)
}

type PipelineSourceTriggerPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PipelineSourceTriggerPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineSourceTriggerProperties)(nil)).Elem()
}

func (o PipelineSourceTriggerPropertiesPtrOutput) ToPipelineSourceTriggerPropertiesPtrOutput() PipelineSourceTriggerPropertiesPtrOutput {
	return o
}

func (o PipelineSourceTriggerPropertiesPtrOutput) ToPipelineSourceTriggerPropertiesPtrOutputWithContext(ctx context.Context) PipelineSourceTriggerPropertiesPtrOutput {
	return o
}

func (o PipelineSourceTriggerPropertiesPtrOutput) Elem() PipelineSourceTriggerPropertiesOutput {
	return o.ApplyT(func(v *PipelineSourceTriggerProperties) PipelineSourceTriggerProperties {
		if v != nil {
			return *v
		}
		var ret PipelineSourceTriggerProperties
		return ret
	}).(PipelineSourceTriggerPropertiesOutput)
}

func (o PipelineSourceTriggerPropertiesPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineSourceTriggerProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type PipelineSourceTriggerPropertiesResponse struct {
	Status string `pulumi:"status"`
}


func (val *PipelineSourceTriggerPropertiesResponse) Defaults() *PipelineSourceTriggerPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		tmp.Status = "Enabled"
	}
	return &tmp
}

type PipelineSourceTriggerPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PipelineSourceTriggerPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineSourceTriggerPropertiesResponse)(nil)).Elem()
}

func (o PipelineSourceTriggerPropertiesResponseOutput) ToPipelineSourceTriggerPropertiesResponseOutput() PipelineSourceTriggerPropertiesResponseOutput {
	return o
}

func (o PipelineSourceTriggerPropertiesResponseOutput) ToPipelineSourceTriggerPropertiesResponseOutputWithContext(ctx context.Context) PipelineSourceTriggerPropertiesResponseOutput {
	return o
}

func (o PipelineSourceTriggerPropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PipelineSourceTriggerPropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

type PipelineSourceTriggerPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelineSourceTriggerPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineSourceTriggerPropertiesResponse)(nil)).Elem()
}

func (o PipelineSourceTriggerPropertiesResponsePtrOutput) ToPipelineSourceTriggerPropertiesResponsePtrOutput() PipelineSourceTriggerPropertiesResponsePtrOutput {
	return o
}

func (o PipelineSourceTriggerPropertiesResponsePtrOutput) ToPipelineSourceTriggerPropertiesResponsePtrOutputWithContext(ctx context.Context) PipelineSourceTriggerPropertiesResponsePtrOutput {
	return o
}

func (o PipelineSourceTriggerPropertiesResponsePtrOutput) Elem() PipelineSourceTriggerPropertiesResponseOutput {
	return o.ApplyT(func(v *PipelineSourceTriggerPropertiesResponse) PipelineSourceTriggerPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PipelineSourceTriggerPropertiesResponse
		return ret
	}).(PipelineSourceTriggerPropertiesResponseOutput)
}

func (o PipelineSourceTriggerPropertiesResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineSourceTriggerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type PipelineTriggerDescriptorResponse struct {
	SourceTrigger *PipelineSourceTriggerDescriptorResponse `pulumi:"sourceTrigger"`
}

type PipelineTriggerDescriptorResponseOutput struct{ *pulumi.OutputState }

func (PipelineTriggerDescriptorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineTriggerDescriptorResponse)(nil)).Elem()
}

func (o PipelineTriggerDescriptorResponseOutput) ToPipelineTriggerDescriptorResponseOutput() PipelineTriggerDescriptorResponseOutput {
	return o
}

func (o PipelineTriggerDescriptorResponseOutput) ToPipelineTriggerDescriptorResponseOutputWithContext(ctx context.Context) PipelineTriggerDescriptorResponseOutput {
	return o
}

func (o PipelineTriggerDescriptorResponseOutput) SourceTrigger() PipelineSourceTriggerDescriptorResponsePtrOutput {
	return o.ApplyT(func(v PipelineTriggerDescriptorResponse) *PipelineSourceTriggerDescriptorResponse {
		return v.SourceTrigger
	}).(PipelineSourceTriggerDescriptorResponsePtrOutput)
}

type PipelineTriggerDescriptorResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelineTriggerDescriptorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineTriggerDescriptorResponse)(nil)).Elem()
}

func (o PipelineTriggerDescriptorResponsePtrOutput) ToPipelineTriggerDescriptorResponsePtrOutput() PipelineTriggerDescriptorResponsePtrOutput {
	return o
}

func (o PipelineTriggerDescriptorResponsePtrOutput) ToPipelineTriggerDescriptorResponsePtrOutputWithContext(ctx context.Context) PipelineTriggerDescriptorResponsePtrOutput {
	return o
}

func (o PipelineTriggerDescriptorResponsePtrOutput) Elem() PipelineTriggerDescriptorResponseOutput {
	return o.ApplyT(func(v *PipelineTriggerDescriptorResponse) PipelineTriggerDescriptorResponse {
		if v != nil {
			return *v
		}
		var ret PipelineTriggerDescriptorResponse
		return ret
	}).(PipelineTriggerDescriptorResponseOutput)
}

func (o PipelineTriggerDescriptorResponsePtrOutput) SourceTrigger() PipelineSourceTriggerDescriptorResponsePtrOutput {
	return o.ApplyT(func(v *PipelineTriggerDescriptorResponse) *PipelineSourceTriggerDescriptorResponse {
		if v == nil {
			return nil
		}
		return v.SourceTrigger
	}).(PipelineSourceTriggerDescriptorResponsePtrOutput)
}

type PipelineTriggerProperties struct {
	SourceTrigger *PipelineSourceTriggerProperties `pulumi:"sourceTrigger"`
}


func (val *PipelineTriggerProperties) Defaults() *PipelineTriggerProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceTrigger = tmp.SourceTrigger.Defaults()

	return &tmp
}





type PipelineTriggerPropertiesInput interface {
	pulumi.Input

	ToPipelineTriggerPropertiesOutput() PipelineTriggerPropertiesOutput
	ToPipelineTriggerPropertiesOutputWithContext(context.Context) PipelineTriggerPropertiesOutput
}

type PipelineTriggerPropertiesArgs struct {
	SourceTrigger PipelineSourceTriggerPropertiesPtrInput `pulumi:"sourceTrigger"`
}

func (PipelineTriggerPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineTriggerProperties)(nil)).Elem()
}

func (i PipelineTriggerPropertiesArgs) ToPipelineTriggerPropertiesOutput() PipelineTriggerPropertiesOutput {
	return i.ToPipelineTriggerPropertiesOutputWithContext(context.Background())
}

func (i PipelineTriggerPropertiesArgs) ToPipelineTriggerPropertiesOutputWithContext(ctx context.Context) PipelineTriggerPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTriggerPropertiesOutput)
}

func (i PipelineTriggerPropertiesArgs) ToPipelineTriggerPropertiesPtrOutput() PipelineTriggerPropertiesPtrOutput {
	return i.ToPipelineTriggerPropertiesPtrOutputWithContext(context.Background())
}

func (i PipelineTriggerPropertiesArgs) ToPipelineTriggerPropertiesPtrOutputWithContext(ctx context.Context) PipelineTriggerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTriggerPropertiesOutput).ToPipelineTriggerPropertiesPtrOutputWithContext(ctx)
}









type PipelineTriggerPropertiesPtrInput interface {
	pulumi.Input

	ToPipelineTriggerPropertiesPtrOutput() PipelineTriggerPropertiesPtrOutput
	ToPipelineTriggerPropertiesPtrOutputWithContext(context.Context) PipelineTriggerPropertiesPtrOutput
}

type pipelineTriggerPropertiesPtrType PipelineTriggerPropertiesArgs

func PipelineTriggerPropertiesPtr(v *PipelineTriggerPropertiesArgs) PipelineTriggerPropertiesPtrInput {
	return (*pipelineTriggerPropertiesPtrType)(v)
}

func (*pipelineTriggerPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineTriggerProperties)(nil)).Elem()
}

func (i *pipelineTriggerPropertiesPtrType) ToPipelineTriggerPropertiesPtrOutput() PipelineTriggerPropertiesPtrOutput {
	return i.ToPipelineTriggerPropertiesPtrOutputWithContext(context.Background())
}

func (i *pipelineTriggerPropertiesPtrType) ToPipelineTriggerPropertiesPtrOutputWithContext(ctx context.Context) PipelineTriggerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTriggerPropertiesPtrOutput)
}

type PipelineTriggerPropertiesOutput struct{ *pulumi.OutputState }

func (PipelineTriggerPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineTriggerProperties)(nil)).Elem()
}

func (o PipelineTriggerPropertiesOutput) ToPipelineTriggerPropertiesOutput() PipelineTriggerPropertiesOutput {
	return o
}

func (o PipelineTriggerPropertiesOutput) ToPipelineTriggerPropertiesOutputWithContext(ctx context.Context) PipelineTriggerPropertiesOutput {
	return o
}

func (o PipelineTriggerPropertiesOutput) ToPipelineTriggerPropertiesPtrOutput() PipelineTriggerPropertiesPtrOutput {
	return o.ToPipelineTriggerPropertiesPtrOutputWithContext(context.Background())
}

func (o PipelineTriggerPropertiesOutput) ToPipelineTriggerPropertiesPtrOutputWithContext(ctx context.Context) PipelineTriggerPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineTriggerProperties) *PipelineTriggerProperties {
		return &v
	}).(PipelineTriggerPropertiesPtrOutput)
}

func (o PipelineTriggerPropertiesOutput) SourceTrigger() PipelineSourceTriggerPropertiesPtrOutput {
	return o.ApplyT(func(v PipelineTriggerProperties) *PipelineSourceTriggerProperties { return v.SourceTrigger }).(PipelineSourceTriggerPropertiesPtrOutput)
}

type PipelineTriggerPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PipelineTriggerPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineTriggerProperties)(nil)).Elem()
}

func (o PipelineTriggerPropertiesPtrOutput) ToPipelineTriggerPropertiesPtrOutput() PipelineTriggerPropertiesPtrOutput {
	return o
}

func (o PipelineTriggerPropertiesPtrOutput) ToPipelineTriggerPropertiesPtrOutputWithContext(ctx context.Context) PipelineTriggerPropertiesPtrOutput {
	return o
}

func (o PipelineTriggerPropertiesPtrOutput) Elem() PipelineTriggerPropertiesOutput {
	return o.ApplyT(func(v *PipelineTriggerProperties) PipelineTriggerProperties {
		if v != nil {
			return *v
		}
		var ret PipelineTriggerProperties
		return ret
	}).(PipelineTriggerPropertiesOutput)
}

func (o PipelineTriggerPropertiesPtrOutput) SourceTrigger() PipelineSourceTriggerPropertiesPtrOutput {
	return o.ApplyT(func(v *PipelineTriggerProperties) *PipelineSourceTriggerProperties {
		if v == nil {
			return nil
		}
		return v.SourceTrigger
	}).(PipelineSourceTriggerPropertiesPtrOutput)
}

type PipelineTriggerPropertiesResponse struct {
	SourceTrigger *PipelineSourceTriggerPropertiesResponse `pulumi:"sourceTrigger"`
}


func (val *PipelineTriggerPropertiesResponse) Defaults() *PipelineTriggerPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceTrigger = tmp.SourceTrigger.Defaults()

	return &tmp
}

type PipelineTriggerPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PipelineTriggerPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineTriggerPropertiesResponse)(nil)).Elem()
}

func (o PipelineTriggerPropertiesResponseOutput) ToPipelineTriggerPropertiesResponseOutput() PipelineTriggerPropertiesResponseOutput {
	return o
}

func (o PipelineTriggerPropertiesResponseOutput) ToPipelineTriggerPropertiesResponseOutputWithContext(ctx context.Context) PipelineTriggerPropertiesResponseOutput {
	return o
}

func (o PipelineTriggerPropertiesResponseOutput) SourceTrigger() PipelineSourceTriggerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v PipelineTriggerPropertiesResponse) *PipelineSourceTriggerPropertiesResponse {
		return v.SourceTrigger
	}).(PipelineSourceTriggerPropertiesResponsePtrOutput)
}

type PipelineTriggerPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelineTriggerPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineTriggerPropertiesResponse)(nil)).Elem()
}

func (o PipelineTriggerPropertiesResponsePtrOutput) ToPipelineTriggerPropertiesResponsePtrOutput() PipelineTriggerPropertiesResponsePtrOutput {
	return o
}

func (o PipelineTriggerPropertiesResponsePtrOutput) ToPipelineTriggerPropertiesResponsePtrOutputWithContext(ctx context.Context) PipelineTriggerPropertiesResponsePtrOutput {
	return o
}

func (o PipelineTriggerPropertiesResponsePtrOutput) Elem() PipelineTriggerPropertiesResponseOutput {
	return o.ApplyT(func(v *PipelineTriggerPropertiesResponse) PipelineTriggerPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PipelineTriggerPropertiesResponse
		return ret
	}).(PipelineTriggerPropertiesResponseOutput)
}

func (o PipelineTriggerPropertiesResponsePtrOutput) SourceTrigger() PipelineSourceTriggerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *PipelineTriggerPropertiesResponse) *PipelineSourceTriggerPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.SourceTrigger
	}).(PipelineSourceTriggerPropertiesResponsePtrOutput)
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

type Policies struct {
	QuarantinePolicy *QuarantinePolicy `pulumi:"quarantinePolicy"`
	RetentionPolicy  *RetentionPolicy  `pulumi:"retentionPolicy"`
	TrustPolicy      *TrustPolicy      `pulumi:"trustPolicy"`
}


func (val *Policies) Defaults() *Policies {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.QuarantinePolicy = tmp.QuarantinePolicy.Defaults()

	tmp.RetentionPolicy = tmp.RetentionPolicy.Defaults()

	tmp.TrustPolicy = tmp.TrustPolicy.Defaults()

	return &tmp
}





type PoliciesInput interface {
	pulumi.Input

	ToPoliciesOutput() PoliciesOutput
	ToPoliciesOutputWithContext(context.Context) PoliciesOutput
}

type PoliciesArgs struct {
	QuarantinePolicy QuarantinePolicyPtrInput `pulumi:"quarantinePolicy"`
	RetentionPolicy  RetentionPolicyPtrInput  `pulumi:"retentionPolicy"`
	TrustPolicy      TrustPolicyPtrInput      `pulumi:"trustPolicy"`
}

func (PoliciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Policies)(nil)).Elem()
}

func (i PoliciesArgs) ToPoliciesOutput() PoliciesOutput {
	return i.ToPoliciesOutputWithContext(context.Background())
}

func (i PoliciesArgs) ToPoliciesOutputWithContext(ctx context.Context) PoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoliciesOutput)
}

func (i PoliciesArgs) ToPoliciesPtrOutput() PoliciesPtrOutput {
	return i.ToPoliciesPtrOutputWithContext(context.Background())
}

func (i PoliciesArgs) ToPoliciesPtrOutputWithContext(ctx context.Context) PoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoliciesOutput).ToPoliciesPtrOutputWithContext(ctx)
}









type PoliciesPtrInput interface {
	pulumi.Input

	ToPoliciesPtrOutput() PoliciesPtrOutput
	ToPoliciesPtrOutputWithContext(context.Context) PoliciesPtrOutput
}

type policiesPtrType PoliciesArgs

func PoliciesPtr(v *PoliciesArgs) PoliciesPtrInput {
	return (*policiesPtrType)(v)
}

func (*policiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Policies)(nil)).Elem()
}

func (i *policiesPtrType) ToPoliciesPtrOutput() PoliciesPtrOutput {
	return i.ToPoliciesPtrOutputWithContext(context.Background())
}

func (i *policiesPtrType) ToPoliciesPtrOutputWithContext(ctx context.Context) PoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoliciesPtrOutput)
}

type PoliciesOutput struct{ *pulumi.OutputState }

func (PoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Policies)(nil)).Elem()
}

func (o PoliciesOutput) ToPoliciesOutput() PoliciesOutput {
	return o
}

func (o PoliciesOutput) ToPoliciesOutputWithContext(ctx context.Context) PoliciesOutput {
	return o
}

func (o PoliciesOutput) ToPoliciesPtrOutput() PoliciesPtrOutput {
	return o.ToPoliciesPtrOutputWithContext(context.Background())
}

func (o PoliciesOutput) ToPoliciesPtrOutputWithContext(ctx context.Context) PoliciesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Policies) *Policies {
		return &v
	}).(PoliciesPtrOutput)
}

func (o PoliciesOutput) QuarantinePolicy() QuarantinePolicyPtrOutput {
	return o.ApplyT(func(v Policies) *QuarantinePolicy { return v.QuarantinePolicy }).(QuarantinePolicyPtrOutput)
}

func (o PoliciesOutput) RetentionPolicy() RetentionPolicyPtrOutput {
	return o.ApplyT(func(v Policies) *RetentionPolicy { return v.RetentionPolicy }).(RetentionPolicyPtrOutput)
}

func (o PoliciesOutput) TrustPolicy() TrustPolicyPtrOutput {
	return o.ApplyT(func(v Policies) *TrustPolicy { return v.TrustPolicy }).(TrustPolicyPtrOutput)
}

type PoliciesPtrOutput struct{ *pulumi.OutputState }

func (PoliciesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policies)(nil)).Elem()
}

func (o PoliciesPtrOutput) ToPoliciesPtrOutput() PoliciesPtrOutput {
	return o
}

func (o PoliciesPtrOutput) ToPoliciesPtrOutputWithContext(ctx context.Context) PoliciesPtrOutput {
	return o
}

func (o PoliciesPtrOutput) Elem() PoliciesOutput {
	return o.ApplyT(func(v *Policies) Policies {
		if v != nil {
			return *v
		}
		var ret Policies
		return ret
	}).(PoliciesOutput)
}

func (o PoliciesPtrOutput) QuarantinePolicy() QuarantinePolicyPtrOutput {
	return o.ApplyT(func(v *Policies) *QuarantinePolicy {
		if v == nil {
			return nil
		}
		return v.QuarantinePolicy
	}).(QuarantinePolicyPtrOutput)
}

func (o PoliciesPtrOutput) RetentionPolicy() RetentionPolicyPtrOutput {
	return o.ApplyT(func(v *Policies) *RetentionPolicy {
		if v == nil {
			return nil
		}
		return v.RetentionPolicy
	}).(RetentionPolicyPtrOutput)
}

func (o PoliciesPtrOutput) TrustPolicy() TrustPolicyPtrOutput {
	return o.ApplyT(func(v *Policies) *TrustPolicy {
		if v == nil {
			return nil
		}
		return v.TrustPolicy
	}).(TrustPolicyPtrOutput)
}

type PoliciesResponse struct {
	QuarantinePolicy *QuarantinePolicyResponse `pulumi:"quarantinePolicy"`
	RetentionPolicy  *RetentionPolicyResponse  `pulumi:"retentionPolicy"`
	TrustPolicy      *TrustPolicyResponse      `pulumi:"trustPolicy"`
}


func (val *PoliciesResponse) Defaults() *PoliciesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.QuarantinePolicy = tmp.QuarantinePolicy.Defaults()

	tmp.RetentionPolicy = tmp.RetentionPolicy.Defaults()

	tmp.TrustPolicy = tmp.TrustPolicy.Defaults()

	return &tmp
}

type PoliciesResponseOutput struct{ *pulumi.OutputState }

func (PoliciesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PoliciesResponse)(nil)).Elem()
}

func (o PoliciesResponseOutput) ToPoliciesResponseOutput() PoliciesResponseOutput {
	return o
}

func (o PoliciesResponseOutput) ToPoliciesResponseOutputWithContext(ctx context.Context) PoliciesResponseOutput {
	return o
}

func (o PoliciesResponseOutput) QuarantinePolicy() QuarantinePolicyResponsePtrOutput {
	return o.ApplyT(func(v PoliciesResponse) *QuarantinePolicyResponse { return v.QuarantinePolicy }).(QuarantinePolicyResponsePtrOutput)
}

func (o PoliciesResponseOutput) RetentionPolicy() RetentionPolicyResponsePtrOutput {
	return o.ApplyT(func(v PoliciesResponse) *RetentionPolicyResponse { return v.RetentionPolicy }).(RetentionPolicyResponsePtrOutput)
}

func (o PoliciesResponseOutput) TrustPolicy() TrustPolicyResponsePtrOutput {
	return o.ApplyT(func(v PoliciesResponse) *TrustPolicyResponse { return v.TrustPolicy }).(TrustPolicyResponsePtrOutput)
}

type PoliciesResponsePtrOutput struct{ *pulumi.OutputState }

func (PoliciesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PoliciesResponse)(nil)).Elem()
}

func (o PoliciesResponsePtrOutput) ToPoliciesResponsePtrOutput() PoliciesResponsePtrOutput {
	return o
}

func (o PoliciesResponsePtrOutput) ToPoliciesResponsePtrOutputWithContext(ctx context.Context) PoliciesResponsePtrOutput {
	return o
}

func (o PoliciesResponsePtrOutput) Elem() PoliciesResponseOutput {
	return o.ApplyT(func(v *PoliciesResponse) PoliciesResponse {
		if v != nil {
			return *v
		}
		var ret PoliciesResponse
		return ret
	}).(PoliciesResponseOutput)
}

func (o PoliciesResponsePtrOutput) QuarantinePolicy() QuarantinePolicyResponsePtrOutput {
	return o.ApplyT(func(v *PoliciesResponse) *QuarantinePolicyResponse {
		if v == nil {
			return nil
		}
		return v.QuarantinePolicy
	}).(QuarantinePolicyResponsePtrOutput)
}

func (o PoliciesResponsePtrOutput) RetentionPolicy() RetentionPolicyResponsePtrOutput {
	return o.ApplyT(func(v *PoliciesResponse) *RetentionPolicyResponse {
		if v == nil {
			return nil
		}
		return v.RetentionPolicy
	}).(RetentionPolicyResponsePtrOutput)
}

func (o PoliciesResponsePtrOutput) TrustPolicy() TrustPolicyResponsePtrOutput {
	return o.ApplyT(func(v *PoliciesResponse) *TrustPolicyResponse {
		if v == nil {
			return nil
		}
		return v.TrustPolicy
	}).(TrustPolicyResponsePtrOutput)
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

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
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

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
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

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
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
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
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

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
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

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
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

type ProgressPropertiesResponse struct {
	Percentage *string `pulumi:"percentage"`
}

type ProgressPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ProgressPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProgressPropertiesResponse)(nil)).Elem()
}

func (o ProgressPropertiesResponseOutput) ToProgressPropertiesResponseOutput() ProgressPropertiesResponseOutput {
	return o
}

func (o ProgressPropertiesResponseOutput) ToProgressPropertiesResponseOutputWithContext(ctx context.Context) ProgressPropertiesResponseOutput {
	return o
}

func (o ProgressPropertiesResponseOutput) Percentage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProgressPropertiesResponse) *string { return v.Percentage }).(pulumi.StringPtrOutput)
}

type ProgressPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ProgressPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProgressPropertiesResponse)(nil)).Elem()
}

func (o ProgressPropertiesResponsePtrOutput) ToProgressPropertiesResponsePtrOutput() ProgressPropertiesResponsePtrOutput {
	return o
}

func (o ProgressPropertiesResponsePtrOutput) ToProgressPropertiesResponsePtrOutputWithContext(ctx context.Context) ProgressPropertiesResponsePtrOutput {
	return o
}

func (o ProgressPropertiesResponsePtrOutput) Elem() ProgressPropertiesResponseOutput {
	return o.ApplyT(func(v *ProgressPropertiesResponse) ProgressPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ProgressPropertiesResponse
		return ret
	}).(ProgressPropertiesResponseOutput)
}

func (o ProgressPropertiesResponsePtrOutput) Percentage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProgressPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Percentage
	}).(pulumi.StringPtrOutput)
}

type QuarantinePolicy struct {
	Status *string `pulumi:"status"`
}


func (val *QuarantinePolicy) Defaults() *QuarantinePolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "disabled"
		tmp.Status = &status_
	}
	return &tmp
}





type QuarantinePolicyInput interface {
	pulumi.Input

	ToQuarantinePolicyOutput() QuarantinePolicyOutput
	ToQuarantinePolicyOutputWithContext(context.Context) QuarantinePolicyOutput
}

type QuarantinePolicyArgs struct {
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (QuarantinePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QuarantinePolicy)(nil)).Elem()
}

func (i QuarantinePolicyArgs) ToQuarantinePolicyOutput() QuarantinePolicyOutput {
	return i.ToQuarantinePolicyOutputWithContext(context.Background())
}

func (i QuarantinePolicyArgs) ToQuarantinePolicyOutputWithContext(ctx context.Context) QuarantinePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuarantinePolicyOutput)
}

func (i QuarantinePolicyArgs) ToQuarantinePolicyPtrOutput() QuarantinePolicyPtrOutput {
	return i.ToQuarantinePolicyPtrOutputWithContext(context.Background())
}

func (i QuarantinePolicyArgs) ToQuarantinePolicyPtrOutputWithContext(ctx context.Context) QuarantinePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuarantinePolicyOutput).ToQuarantinePolicyPtrOutputWithContext(ctx)
}









type QuarantinePolicyPtrInput interface {
	pulumi.Input

	ToQuarantinePolicyPtrOutput() QuarantinePolicyPtrOutput
	ToQuarantinePolicyPtrOutputWithContext(context.Context) QuarantinePolicyPtrOutput
}

type quarantinePolicyPtrType QuarantinePolicyArgs

func QuarantinePolicyPtr(v *QuarantinePolicyArgs) QuarantinePolicyPtrInput {
	return (*quarantinePolicyPtrType)(v)
}

func (*quarantinePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QuarantinePolicy)(nil)).Elem()
}

func (i *quarantinePolicyPtrType) ToQuarantinePolicyPtrOutput() QuarantinePolicyPtrOutput {
	return i.ToQuarantinePolicyPtrOutputWithContext(context.Background())
}

func (i *quarantinePolicyPtrType) ToQuarantinePolicyPtrOutputWithContext(ctx context.Context) QuarantinePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuarantinePolicyPtrOutput)
}

type QuarantinePolicyOutput struct{ *pulumi.OutputState }

func (QuarantinePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QuarantinePolicy)(nil)).Elem()
}

func (o QuarantinePolicyOutput) ToQuarantinePolicyOutput() QuarantinePolicyOutput {
	return o
}

func (o QuarantinePolicyOutput) ToQuarantinePolicyOutputWithContext(ctx context.Context) QuarantinePolicyOutput {
	return o
}

func (o QuarantinePolicyOutput) ToQuarantinePolicyPtrOutput() QuarantinePolicyPtrOutput {
	return o.ToQuarantinePolicyPtrOutputWithContext(context.Background())
}

func (o QuarantinePolicyOutput) ToQuarantinePolicyPtrOutputWithContext(ctx context.Context) QuarantinePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QuarantinePolicy) *QuarantinePolicy {
		return &v
	}).(QuarantinePolicyPtrOutput)
}

func (o QuarantinePolicyOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QuarantinePolicy) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type QuarantinePolicyPtrOutput struct{ *pulumi.OutputState }

func (QuarantinePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QuarantinePolicy)(nil)).Elem()
}

func (o QuarantinePolicyPtrOutput) ToQuarantinePolicyPtrOutput() QuarantinePolicyPtrOutput {
	return o
}

func (o QuarantinePolicyPtrOutput) ToQuarantinePolicyPtrOutputWithContext(ctx context.Context) QuarantinePolicyPtrOutput {
	return o
}

func (o QuarantinePolicyPtrOutput) Elem() QuarantinePolicyOutput {
	return o.ApplyT(func(v *QuarantinePolicy) QuarantinePolicy {
		if v != nil {
			return *v
		}
		var ret QuarantinePolicy
		return ret
	}).(QuarantinePolicyOutput)
}

func (o QuarantinePolicyPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QuarantinePolicy) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type QuarantinePolicyResponse struct {
	Status *string `pulumi:"status"`
}


func (val *QuarantinePolicyResponse) Defaults() *QuarantinePolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "disabled"
		tmp.Status = &status_
	}
	return &tmp
}

type QuarantinePolicyResponseOutput struct{ *pulumi.OutputState }

func (QuarantinePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QuarantinePolicyResponse)(nil)).Elem()
}

func (o QuarantinePolicyResponseOutput) ToQuarantinePolicyResponseOutput() QuarantinePolicyResponseOutput {
	return o
}

func (o QuarantinePolicyResponseOutput) ToQuarantinePolicyResponseOutputWithContext(ctx context.Context) QuarantinePolicyResponseOutput {
	return o
}

func (o QuarantinePolicyResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QuarantinePolicyResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type QuarantinePolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (QuarantinePolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QuarantinePolicyResponse)(nil)).Elem()
}

func (o QuarantinePolicyResponsePtrOutput) ToQuarantinePolicyResponsePtrOutput() QuarantinePolicyResponsePtrOutput {
	return o
}

func (o QuarantinePolicyResponsePtrOutput) ToQuarantinePolicyResponsePtrOutputWithContext(ctx context.Context) QuarantinePolicyResponsePtrOutput {
	return o
}

func (o QuarantinePolicyResponsePtrOutput) Elem() QuarantinePolicyResponseOutput {
	return o.ApplyT(func(v *QuarantinePolicyResponse) QuarantinePolicyResponse {
		if v != nil {
			return *v
		}
		var ret QuarantinePolicyResponse
		return ret
	}).(QuarantinePolicyResponseOutput)
}

func (o QuarantinePolicyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QuarantinePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type RegistryPasswordResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type RequestResponse struct {
	Addr      *string `pulumi:"addr"`
	Host      *string `pulumi:"host"`
	Id        *string `pulumi:"id"`
	Method    *string `pulumi:"method"`
	Useragent *string `pulumi:"useragent"`
}

type RetentionPolicy struct {
	Days   *int    `pulumi:"days"`
	Status *string `pulumi:"status"`
}


func (val *RetentionPolicy) Defaults() *RetentionPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Days) {
		days_ := 7
		tmp.Days = &days_
	}
	if isZero(tmp.Status) {
		status_ := "disabled"
		tmp.Status = &status_
	}
	return &tmp
}





type RetentionPolicyInput interface {
	pulumi.Input

	ToRetentionPolicyOutput() RetentionPolicyOutput
	ToRetentionPolicyOutputWithContext(context.Context) RetentionPolicyOutput
}

type RetentionPolicyArgs struct {
	Days   pulumi.IntPtrInput    `pulumi:"days"`
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (RetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicy)(nil)).Elem()
}

func (i RetentionPolicyArgs) ToRetentionPolicyOutput() RetentionPolicyOutput {
	return i.ToRetentionPolicyOutputWithContext(context.Background())
}

func (i RetentionPolicyArgs) ToRetentionPolicyOutputWithContext(ctx context.Context) RetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyOutput)
}

func (i RetentionPolicyArgs) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return i.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (i RetentionPolicyArgs) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyOutput).ToRetentionPolicyPtrOutputWithContext(ctx)
}









type RetentionPolicyPtrInput interface {
	pulumi.Input

	ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput
	ToRetentionPolicyPtrOutputWithContext(context.Context) RetentionPolicyPtrOutput
}

type retentionPolicyPtrType RetentionPolicyArgs

func RetentionPolicyPtr(v *RetentionPolicyArgs) RetentionPolicyPtrInput {
	return (*retentionPolicyPtrType)(v)
}

func (*retentionPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicy)(nil)).Elem()
}

func (i *retentionPolicyPtrType) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return i.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (i *retentionPolicyPtrType) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyPtrOutput)
}

type RetentionPolicyOutput struct{ *pulumi.OutputState }

func (RetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicy)(nil)).Elem()
}

func (o RetentionPolicyOutput) ToRetentionPolicyOutput() RetentionPolicyOutput {
	return o
}

func (o RetentionPolicyOutput) ToRetentionPolicyOutputWithContext(ctx context.Context) RetentionPolicyOutput {
	return o
}

func (o RetentionPolicyOutput) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return o.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (o RetentionPolicyOutput) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetentionPolicy) *RetentionPolicy {
		return &v
	}).(RetentionPolicyPtrOutput)
}

func (o RetentionPolicyOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetentionPolicy) *int { return v.Days }).(pulumi.IntPtrOutput)
}

func (o RetentionPolicyOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RetentionPolicy) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type RetentionPolicyPtrOutput struct{ *pulumi.OutputState }

func (RetentionPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicy)(nil)).Elem()
}

func (o RetentionPolicyPtrOutput) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return o
}

func (o RetentionPolicyPtrOutput) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return o
}

func (o RetentionPolicyPtrOutput) Elem() RetentionPolicyOutput {
	return o.ApplyT(func(v *RetentionPolicy) RetentionPolicy {
		if v != nil {
			return *v
		}
		var ret RetentionPolicy
		return ret
	}).(RetentionPolicyOutput)
}

func (o RetentionPolicyPtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetentionPolicy) *int {
		if v == nil {
			return nil
		}
		return v.Days
	}).(pulumi.IntPtrOutput)
}

func (o RetentionPolicyPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RetentionPolicy) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type RetentionPolicyResponse struct {
	Days            *int    `pulumi:"days"`
	LastUpdatedTime string  `pulumi:"lastUpdatedTime"`
	Status          *string `pulumi:"status"`
}


func (val *RetentionPolicyResponse) Defaults() *RetentionPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Days) {
		days_ := 7
		tmp.Days = &days_
	}
	if isZero(tmp.Status) {
		status_ := "disabled"
		tmp.Status = &status_
	}
	return &tmp
}

type RetentionPolicyResponseOutput struct{ *pulumi.OutputState }

func (RetentionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicyResponse)(nil)).Elem()
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponseOutput() RetentionPolicyResponseOutput {
	return o
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponseOutputWithContext(ctx context.Context) RetentionPolicyResponseOutput {
	return o
}

func (o RetentionPolicyResponseOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetentionPolicyResponse) *int { return v.Days }).(pulumi.IntPtrOutput)
}

func (o RetentionPolicyResponseOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v RetentionPolicyResponse) string { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

func (o RetentionPolicyResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RetentionPolicyResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type RetentionPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (RetentionPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicyResponse)(nil)).Elem()
}

func (o RetentionPolicyResponsePtrOutput) ToRetentionPolicyResponsePtrOutput() RetentionPolicyResponsePtrOutput {
	return o
}

func (o RetentionPolicyResponsePtrOutput) ToRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) RetentionPolicyResponsePtrOutput {
	return o
}

func (o RetentionPolicyResponsePtrOutput) Elem() RetentionPolicyResponseOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) RetentionPolicyResponse {
		if v != nil {
			return *v
		}
		var ret RetentionPolicyResponse
		return ret
	}).(RetentionPolicyResponseOutput)
}

func (o RetentionPolicyResponsePtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.Days
	}).(pulumi.IntPtrOutput)
}

func (o RetentionPolicyResponsePtrOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastUpdatedTime
	}).(pulumi.StringPtrOutput)
}

func (o RetentionPolicyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
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


func (val *RunResponse) Defaults() *RunResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchiveEnabled) {
		isArchiveEnabled_ := false
		tmp.IsArchiveEnabled = &isArchiveEnabled_
	}
	return &tmp
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


func (val *SetValue) Defaults() *SetValue {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsSecret) {
		isSecret_ := false
		tmp.IsSecret = &isSecret_
	}
	return &tmp
}

type SetValueResponse struct {
	IsSecret *bool  `pulumi:"isSecret"`
	Name     string `pulumi:"name"`
	Value    string `pulumi:"value"`
}


func (val *SetValueResponse) Defaults() *SetValueResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsSecret) {
		isSecret_ := false
		tmp.IsSecret = &isSecret_
	}
	return &tmp
}

type Sku struct {
	Name string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
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

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
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

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type SourceControlAuthInfoResponse struct {
	ExpiresIn    *int    `pulumi:"expiresIn"`
	RefreshToken *string `pulumi:"refreshToken"`
	Scope        *string `pulumi:"scope"`
	Token        string  `pulumi:"token"`
	TokenType    *string `pulumi:"tokenType"`
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

type SourceResponse struct {
	Addr       *string `pulumi:"addr"`
	InstanceID *string `pulumi:"instanceID"`
}

type SourceTrigger struct {
	Name                string           `pulumi:"name"`
	SourceRepository    SourceProperties `pulumi:"sourceRepository"`
	SourceTriggerEvents []string         `pulumi:"sourceTriggerEvents"`
	Status              *string          `pulumi:"status"`
}


func (val *SourceTrigger) Defaults() *SourceTrigger {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "Enabled"
		tmp.Status = &status_
	}
	return &tmp
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


func (val *SourceTriggerResponse) Defaults() *SourceTriggerResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "Enabled"
		tmp.Status = &status_
	}
	return &tmp
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

type StatusDetailPropertiesResponse struct {
	Code          string `pulumi:"code"`
	CorrelationId string `pulumi:"correlationId"`
	Description   string `pulumi:"description"`
	Timestamp     string `pulumi:"timestamp"`
	Type          string `pulumi:"type"`
}

type StatusDetailPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StatusDetailPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusDetailPropertiesResponse)(nil)).Elem()
}

func (o StatusDetailPropertiesResponseOutput) ToStatusDetailPropertiesResponseOutput() StatusDetailPropertiesResponseOutput {
	return o
}

func (o StatusDetailPropertiesResponseOutput) ToStatusDetailPropertiesResponseOutputWithContext(ctx context.Context) StatusDetailPropertiesResponseOutput {
	return o
}

func (o StatusDetailPropertiesResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v StatusDetailPropertiesResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o StatusDetailPropertiesResponseOutput) CorrelationId() pulumi.StringOutput {
	return o.ApplyT(func(v StatusDetailPropertiesResponse) string { return v.CorrelationId }).(pulumi.StringOutput)
}

func (o StatusDetailPropertiesResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v StatusDetailPropertiesResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o StatusDetailPropertiesResponseOutput) Timestamp() pulumi.StringOutput {
	return o.ApplyT(func(v StatusDetailPropertiesResponse) string { return v.Timestamp }).(pulumi.StringOutput)
}

func (o StatusDetailPropertiesResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v StatusDetailPropertiesResponse) string { return v.Type }).(pulumi.StringOutput)
}

type StatusDetailPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (StatusDetailPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StatusDetailPropertiesResponse)(nil)).Elem()
}

func (o StatusDetailPropertiesResponseArrayOutput) ToStatusDetailPropertiesResponseArrayOutput() StatusDetailPropertiesResponseArrayOutput {
	return o
}

func (o StatusDetailPropertiesResponseArrayOutput) ToStatusDetailPropertiesResponseArrayOutputWithContext(ctx context.Context) StatusDetailPropertiesResponseArrayOutput {
	return o
}

func (o StatusDetailPropertiesResponseArrayOutput) Index(i pulumi.IntInput) StatusDetailPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StatusDetailPropertiesResponse {
		return vs[0].([]StatusDetailPropertiesResponse)[vs[1].(int)]
	}).(StatusDetailPropertiesResponseOutput)
}

type StatusResponse struct {
	DisplayStatus string `pulumi:"displayStatus"`
	Message       string `pulumi:"message"`
	Timestamp     string `pulumi:"timestamp"`
}

type StatusResponseOutput struct{ *pulumi.OutputState }

func (StatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusResponse)(nil)).Elem()
}

func (o StatusResponseOutput) ToStatusResponseOutput() StatusResponseOutput {
	return o
}

func (o StatusResponseOutput) ToStatusResponseOutputWithContext(ctx context.Context) StatusResponseOutput {
	return o
}

func (o StatusResponseOutput) DisplayStatus() pulumi.StringOutput {
	return o.ApplyT(func(v StatusResponse) string { return v.DisplayStatus }).(pulumi.StringOutput)
}

func (o StatusResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v StatusResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o StatusResponseOutput) Timestamp() pulumi.StringOutput {
	return o.ApplyT(func(v StatusResponse) string { return v.Timestamp }).(pulumi.StringOutput)
}

type StorageAccountProperties struct {
	Id string `pulumi:"id"`
}





type StorageAccountPropertiesInput interface {
	pulumi.Input

	ToStorageAccountPropertiesOutput() StorageAccountPropertiesOutput
	ToStorageAccountPropertiesOutputWithContext(context.Context) StorageAccountPropertiesOutput
}

type StorageAccountPropertiesArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (StorageAccountPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountProperties)(nil)).Elem()
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesOutput() StorageAccountPropertiesOutput {
	return i.ToStorageAccountPropertiesOutputWithContext(context.Background())
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesOutputWithContext(ctx context.Context) StorageAccountPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesOutput)
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput {
	return i.ToStorageAccountPropertiesPtrOutputWithContext(context.Background())
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesPtrOutputWithContext(ctx context.Context) StorageAccountPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesOutput).ToStorageAccountPropertiesPtrOutputWithContext(ctx)
}









type StorageAccountPropertiesPtrInput interface {
	pulumi.Input

	ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput
	ToStorageAccountPropertiesPtrOutputWithContext(context.Context) StorageAccountPropertiesPtrOutput
}

type storageAccountPropertiesPtrType StorageAccountPropertiesArgs

func StorageAccountPropertiesPtr(v *StorageAccountPropertiesArgs) StorageAccountPropertiesPtrInput {
	return (*storageAccountPropertiesPtrType)(v)
}

func (*storageAccountPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountProperties)(nil)).Elem()
}

func (i *storageAccountPropertiesPtrType) ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput {
	return i.ToStorageAccountPropertiesPtrOutputWithContext(context.Background())
}

func (i *storageAccountPropertiesPtrType) ToStorageAccountPropertiesPtrOutputWithContext(ctx context.Context) StorageAccountPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesPtrOutput)
}

type StorageAccountPropertiesOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountProperties)(nil)).Elem()
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesOutput() StorageAccountPropertiesOutput {
	return o
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesOutputWithContext(ctx context.Context) StorageAccountPropertiesOutput {
	return o
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput {
	return o.ToStorageAccountPropertiesPtrOutputWithContext(context.Background())
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesPtrOutputWithContext(ctx context.Context) StorageAccountPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccountProperties) *StorageAccountProperties {
		return &v
	}).(StorageAccountPropertiesPtrOutput)
}

func (o StorageAccountPropertiesOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountProperties) string { return v.Id }).(pulumi.StringOutput)
}

type StorageAccountPropertiesPtrOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountProperties)(nil)).Elem()
}

func (o StorageAccountPropertiesPtrOutput) ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput {
	return o
}

func (o StorageAccountPropertiesPtrOutput) ToStorageAccountPropertiesPtrOutputWithContext(ctx context.Context) StorageAccountPropertiesPtrOutput {
	return o
}

func (o StorageAccountPropertiesPtrOutput) Elem() StorageAccountPropertiesOutput {
	return o.ApplyT(func(v *StorageAccountProperties) StorageAccountProperties {
		if v != nil {
			return *v
		}
		var ret StorageAccountProperties
		return ret
	}).(StorageAccountPropertiesOutput)
}

func (o StorageAccountPropertiesPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type StorageAccountPropertiesResponse struct {
	Id string `pulumi:"id"`
}

type StorageAccountPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountPropertiesResponse)(nil)).Elem()
}

func (o StorageAccountPropertiesResponseOutput) ToStorageAccountPropertiesResponseOutput() StorageAccountPropertiesResponseOutput {
	return o
}

func (o StorageAccountPropertiesResponseOutput) ToStorageAccountPropertiesResponseOutputWithContext(ctx context.Context) StorageAccountPropertiesResponseOutput {
	return o
}

func (o StorageAccountPropertiesResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountPropertiesResponse) string { return v.Id }).(pulumi.StringOutput)
}

type StorageAccountPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountPropertiesResponse)(nil)).Elem()
}

func (o StorageAccountPropertiesResponsePtrOutput) ToStorageAccountPropertiesResponsePtrOutput() StorageAccountPropertiesResponsePtrOutput {
	return o
}

func (o StorageAccountPropertiesResponsePtrOutput) ToStorageAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) StorageAccountPropertiesResponsePtrOutput {
	return o
}

func (o StorageAccountPropertiesResponsePtrOutput) Elem() StorageAccountPropertiesResponseOutput {
	return o.ApplyT(func(v *StorageAccountPropertiesResponse) StorageAccountPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret StorageAccountPropertiesResponse
		return ret
	}).(StorageAccountPropertiesResponseOutput)
}

func (o StorageAccountPropertiesResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type SyncProperties struct {
	MessageTtl string  `pulumi:"messageTtl"`
	Schedule   *string `pulumi:"schedule"`
	SyncWindow *string `pulumi:"syncWindow"`
	TokenId    string  `pulumi:"tokenId"`
}





type SyncPropertiesInput interface {
	pulumi.Input

	ToSyncPropertiesOutput() SyncPropertiesOutput
	ToSyncPropertiesOutputWithContext(context.Context) SyncPropertiesOutput
}

type SyncPropertiesArgs struct {
	MessageTtl pulumi.StringInput    `pulumi:"messageTtl"`
	Schedule   pulumi.StringPtrInput `pulumi:"schedule"`
	SyncWindow pulumi.StringPtrInput `pulumi:"syncWindow"`
	TokenId    pulumi.StringInput    `pulumi:"tokenId"`
}

func (SyncPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncProperties)(nil)).Elem()
}

func (i SyncPropertiesArgs) ToSyncPropertiesOutput() SyncPropertiesOutput {
	return i.ToSyncPropertiesOutputWithContext(context.Background())
}

func (i SyncPropertiesArgs) ToSyncPropertiesOutputWithContext(ctx context.Context) SyncPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncPropertiesOutput)
}

type SyncPropertiesOutput struct{ *pulumi.OutputState }

func (SyncPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncProperties)(nil)).Elem()
}

func (o SyncPropertiesOutput) ToSyncPropertiesOutput() SyncPropertiesOutput {
	return o
}

func (o SyncPropertiesOutput) ToSyncPropertiesOutputWithContext(ctx context.Context) SyncPropertiesOutput {
	return o
}

func (o SyncPropertiesOutput) MessageTtl() pulumi.StringOutput {
	return o.ApplyT(func(v SyncProperties) string { return v.MessageTtl }).(pulumi.StringOutput)
}

func (o SyncPropertiesOutput) Schedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncProperties) *string { return v.Schedule }).(pulumi.StringPtrOutput)
}

func (o SyncPropertiesOutput) SyncWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncProperties) *string { return v.SyncWindow }).(pulumi.StringPtrOutput)
}

func (o SyncPropertiesOutput) TokenId() pulumi.StringOutput {
	return o.ApplyT(func(v SyncProperties) string { return v.TokenId }).(pulumi.StringOutput)
}

type SyncPropertiesResponse struct {
	GatewayEndpoint string  `pulumi:"gatewayEndpoint"`
	LastSyncTime    string  `pulumi:"lastSyncTime"`
	MessageTtl      string  `pulumi:"messageTtl"`
	Schedule        *string `pulumi:"schedule"`
	SyncWindow      *string `pulumi:"syncWindow"`
	TokenId         string  `pulumi:"tokenId"`
}

type SyncPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SyncPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncPropertiesResponse)(nil)).Elem()
}

func (o SyncPropertiesResponseOutput) ToSyncPropertiesResponseOutput() SyncPropertiesResponseOutput {
	return o
}

func (o SyncPropertiesResponseOutput) ToSyncPropertiesResponseOutputWithContext(ctx context.Context) SyncPropertiesResponseOutput {
	return o
}

func (o SyncPropertiesResponseOutput) GatewayEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v SyncPropertiesResponse) string { return v.GatewayEndpoint }).(pulumi.StringOutput)
}

func (o SyncPropertiesResponseOutput) LastSyncTime() pulumi.StringOutput {
	return o.ApplyT(func(v SyncPropertiesResponse) string { return v.LastSyncTime }).(pulumi.StringOutput)
}

func (o SyncPropertiesResponseOutput) MessageTtl() pulumi.StringOutput {
	return o.ApplyT(func(v SyncPropertiesResponse) string { return v.MessageTtl }).(pulumi.StringOutput)
}

func (o SyncPropertiesResponseOutput) Schedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncPropertiesResponse) *string { return v.Schedule }).(pulumi.StringPtrOutput)
}

func (o SyncPropertiesResponseOutput) SyncWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncPropertiesResponse) *string { return v.SyncWindow }).(pulumi.StringPtrOutput)
}

func (o SyncPropertiesResponseOutput) TokenId() pulumi.StringOutput {
	return o.ApplyT(func(v SyncPropertiesResponse) string { return v.TokenId }).(pulumi.StringOutput)
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

type TargetResponse struct {
	Digest     *string  `pulumi:"digest"`
	Length     *float64 `pulumi:"length"`
	MediaType  *string  `pulumi:"mediaType"`
	Name       *string  `pulumi:"name"`
	Repository *string  `pulumi:"repository"`
	Size       *float64 `pulumi:"size"`
	Tag        *string  `pulumi:"tag"`
	Url        *string  `pulumi:"url"`
	Version    *string  `pulumi:"version"`
}

type TaskRunRequest struct {
	AgentPoolName              *string                     `pulumi:"agentPoolName"`
	IsArchiveEnabled           *bool                       `pulumi:"isArchiveEnabled"`
	LogTemplate                *string                     `pulumi:"logTemplate"`
	OverrideTaskStepProperties *OverrideTaskStepProperties `pulumi:"overrideTaskStepProperties"`
	TaskId                     string                      `pulumi:"taskId"`
	Type                       string                      `pulumi:"type"`
}


func (val *TaskRunRequest) Defaults() *TaskRunRequest {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchiveEnabled) {
		isArchiveEnabled_ := false
		tmp.IsArchiveEnabled = &isArchiveEnabled_
	}
	return &tmp
}

type TaskRunRequestResponse struct {
	AgentPoolName              *string                             `pulumi:"agentPoolName"`
	IsArchiveEnabled           *bool                               `pulumi:"isArchiveEnabled"`
	LogTemplate                *string                             `pulumi:"logTemplate"`
	OverrideTaskStepProperties *OverrideTaskStepPropertiesResponse `pulumi:"overrideTaskStepProperties"`
	TaskId                     string                              `pulumi:"taskId"`
	Type                       string                              `pulumi:"type"`
}


func (val *TaskRunRequestResponse) Defaults() *TaskRunRequestResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchiveEnabled) {
		isArchiveEnabled_ := false
		tmp.IsArchiveEnabled = &isArchiveEnabled_
	}
	return &tmp
}

type TimerTrigger struct {
	Name     string  `pulumi:"name"`
	Schedule string  `pulumi:"schedule"`
	Status   *string `pulumi:"status"`
}


func (val *TimerTrigger) Defaults() *TimerTrigger {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "Enabled"
		tmp.Status = &status_
	}
	return &tmp
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


func (val *TimerTriggerResponse) Defaults() *TimerTriggerResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "Enabled"
		tmp.Status = &status_
	}
	return &tmp
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

type TlsCertificatePropertiesResponse struct {
	Location string `pulumi:"location"`
	Type     string `pulumi:"type"`
}

type TlsCertificatePropertiesResponseOutput struct{ *pulumi.OutputState }

func (TlsCertificatePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TlsCertificatePropertiesResponse)(nil)).Elem()
}

func (o TlsCertificatePropertiesResponseOutput) ToTlsCertificatePropertiesResponseOutput() TlsCertificatePropertiesResponseOutput {
	return o
}

func (o TlsCertificatePropertiesResponseOutput) ToTlsCertificatePropertiesResponseOutputWithContext(ctx context.Context) TlsCertificatePropertiesResponseOutput {
	return o
}

func (o TlsCertificatePropertiesResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v TlsCertificatePropertiesResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o TlsCertificatePropertiesResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v TlsCertificatePropertiesResponse) string { return v.Type }).(pulumi.StringOutput)
}

type TlsCertificatePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (TlsCertificatePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsCertificatePropertiesResponse)(nil)).Elem()
}

func (o TlsCertificatePropertiesResponsePtrOutput) ToTlsCertificatePropertiesResponsePtrOutput() TlsCertificatePropertiesResponsePtrOutput {
	return o
}

func (o TlsCertificatePropertiesResponsePtrOutput) ToTlsCertificatePropertiesResponsePtrOutputWithContext(ctx context.Context) TlsCertificatePropertiesResponsePtrOutput {
	return o
}

func (o TlsCertificatePropertiesResponsePtrOutput) Elem() TlsCertificatePropertiesResponseOutput {
	return o.ApplyT(func(v *TlsCertificatePropertiesResponse) TlsCertificatePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret TlsCertificatePropertiesResponse
		return ret
	}).(TlsCertificatePropertiesResponseOutput)
}

func (o TlsCertificatePropertiesResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TlsCertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o TlsCertificatePropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TlsCertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type TlsPropertiesResponse struct {
	Certificate TlsCertificatePropertiesResponse `pulumi:"certificate"`
	Status      string                           `pulumi:"status"`
}

type TlsPropertiesResponseOutput struct{ *pulumi.OutputState }

func (TlsPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TlsPropertiesResponse)(nil)).Elem()
}

func (o TlsPropertiesResponseOutput) ToTlsPropertiesResponseOutput() TlsPropertiesResponseOutput {
	return o
}

func (o TlsPropertiesResponseOutput) ToTlsPropertiesResponseOutputWithContext(ctx context.Context) TlsPropertiesResponseOutput {
	return o
}

func (o TlsPropertiesResponseOutput) Certificate() TlsCertificatePropertiesResponseOutput {
	return o.ApplyT(func(v TlsPropertiesResponse) TlsCertificatePropertiesResponse { return v.Certificate }).(TlsCertificatePropertiesResponseOutput)
}

func (o TlsPropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v TlsPropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

type TlsPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (TlsPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsPropertiesResponse)(nil)).Elem()
}

func (o TlsPropertiesResponsePtrOutput) ToTlsPropertiesResponsePtrOutput() TlsPropertiesResponsePtrOutput {
	return o
}

func (o TlsPropertiesResponsePtrOutput) ToTlsPropertiesResponsePtrOutputWithContext(ctx context.Context) TlsPropertiesResponsePtrOutput {
	return o
}

func (o TlsPropertiesResponsePtrOutput) Elem() TlsPropertiesResponseOutput {
	return o.ApplyT(func(v *TlsPropertiesResponse) TlsPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret TlsPropertiesResponse
		return ret
	}).(TlsPropertiesResponseOutput)
}

func (o TlsPropertiesResponsePtrOutput) Certificate() TlsCertificatePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *TlsPropertiesResponse) *TlsCertificatePropertiesResponse {
		if v == nil {
			return nil
		}
		return &v.Certificate
	}).(TlsCertificatePropertiesResponsePtrOutput)
}

func (o TlsPropertiesResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TlsPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type TokenCertificate struct {
	EncodedPemCertificate *string `pulumi:"encodedPemCertificate"`
	Expiry                *string `pulumi:"expiry"`
	Name                  *string `pulumi:"name"`
	Thumbprint            *string `pulumi:"thumbprint"`
}





type TokenCertificateInput interface {
	pulumi.Input

	ToTokenCertificateOutput() TokenCertificateOutput
	ToTokenCertificateOutputWithContext(context.Context) TokenCertificateOutput
}

type TokenCertificateArgs struct {
	EncodedPemCertificate pulumi.StringPtrInput `pulumi:"encodedPemCertificate"`
	Expiry                pulumi.StringPtrInput `pulumi:"expiry"`
	Name                  pulumi.StringPtrInput `pulumi:"name"`
	Thumbprint            pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (TokenCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCertificate)(nil)).Elem()
}

func (i TokenCertificateArgs) ToTokenCertificateOutput() TokenCertificateOutput {
	return i.ToTokenCertificateOutputWithContext(context.Background())
}

func (i TokenCertificateArgs) ToTokenCertificateOutputWithContext(ctx context.Context) TokenCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCertificateOutput)
}





type TokenCertificateArrayInput interface {
	pulumi.Input

	ToTokenCertificateArrayOutput() TokenCertificateArrayOutput
	ToTokenCertificateArrayOutputWithContext(context.Context) TokenCertificateArrayOutput
}

type TokenCertificateArray []TokenCertificateInput

func (TokenCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenCertificate)(nil)).Elem()
}

func (i TokenCertificateArray) ToTokenCertificateArrayOutput() TokenCertificateArrayOutput {
	return i.ToTokenCertificateArrayOutputWithContext(context.Background())
}

func (i TokenCertificateArray) ToTokenCertificateArrayOutputWithContext(ctx context.Context) TokenCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCertificateArrayOutput)
}

type TokenCertificateOutput struct{ *pulumi.OutputState }

func (TokenCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCertificate)(nil)).Elem()
}

func (o TokenCertificateOutput) ToTokenCertificateOutput() TokenCertificateOutput {
	return o
}

func (o TokenCertificateOutput) ToTokenCertificateOutputWithContext(ctx context.Context) TokenCertificateOutput {
	return o
}

func (o TokenCertificateOutput) EncodedPemCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificate) *string { return v.EncodedPemCertificate }).(pulumi.StringPtrOutput)
}

func (o TokenCertificateOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificate) *string { return v.Expiry }).(pulumi.StringPtrOutput)
}

func (o TokenCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TokenCertificateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificate) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type TokenCertificateArrayOutput struct{ *pulumi.OutputState }

func (TokenCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenCertificate)(nil)).Elem()
}

func (o TokenCertificateArrayOutput) ToTokenCertificateArrayOutput() TokenCertificateArrayOutput {
	return o
}

func (o TokenCertificateArrayOutput) ToTokenCertificateArrayOutputWithContext(ctx context.Context) TokenCertificateArrayOutput {
	return o
}

func (o TokenCertificateArrayOutput) Index(i pulumi.IntInput) TokenCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenCertificate {
		return vs[0].([]TokenCertificate)[vs[1].(int)]
	}).(TokenCertificateOutput)
}

type TokenCertificateResponse struct {
	EncodedPemCertificate *string `pulumi:"encodedPemCertificate"`
	Expiry                *string `pulumi:"expiry"`
	Name                  *string `pulumi:"name"`
	Thumbprint            *string `pulumi:"thumbprint"`
}

type TokenCertificateResponseOutput struct{ *pulumi.OutputState }

func (TokenCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCertificateResponse)(nil)).Elem()
}

func (o TokenCertificateResponseOutput) ToTokenCertificateResponseOutput() TokenCertificateResponseOutput {
	return o
}

func (o TokenCertificateResponseOutput) ToTokenCertificateResponseOutputWithContext(ctx context.Context) TokenCertificateResponseOutput {
	return o
}

func (o TokenCertificateResponseOutput) EncodedPemCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificateResponse) *string { return v.EncodedPemCertificate }).(pulumi.StringPtrOutput)
}

func (o TokenCertificateResponseOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificateResponse) *string { return v.Expiry }).(pulumi.StringPtrOutput)
}

func (o TokenCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TokenCertificateResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificateResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type TokenCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (TokenCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenCertificateResponse)(nil)).Elem()
}

func (o TokenCertificateResponseArrayOutput) ToTokenCertificateResponseArrayOutput() TokenCertificateResponseArrayOutput {
	return o
}

func (o TokenCertificateResponseArrayOutput) ToTokenCertificateResponseArrayOutputWithContext(ctx context.Context) TokenCertificateResponseArrayOutput {
	return o
}

func (o TokenCertificateResponseArrayOutput) Index(i pulumi.IntInput) TokenCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenCertificateResponse {
		return vs[0].([]TokenCertificateResponse)[vs[1].(int)]
	}).(TokenCertificateResponseOutput)
}

type TokenCredentialsProperties struct {
	Certificates []TokenCertificate `pulumi:"certificates"`
	Passwords    []TokenPassword    `pulumi:"passwords"`
}





type TokenCredentialsPropertiesInput interface {
	pulumi.Input

	ToTokenCredentialsPropertiesOutput() TokenCredentialsPropertiesOutput
	ToTokenCredentialsPropertiesOutputWithContext(context.Context) TokenCredentialsPropertiesOutput
}

type TokenCredentialsPropertiesArgs struct {
	Certificates TokenCertificateArrayInput `pulumi:"certificates"`
	Passwords    TokenPasswordArrayInput    `pulumi:"passwords"`
}

func (TokenCredentialsPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCredentialsProperties)(nil)).Elem()
}

func (i TokenCredentialsPropertiesArgs) ToTokenCredentialsPropertiesOutput() TokenCredentialsPropertiesOutput {
	return i.ToTokenCredentialsPropertiesOutputWithContext(context.Background())
}

func (i TokenCredentialsPropertiesArgs) ToTokenCredentialsPropertiesOutputWithContext(ctx context.Context) TokenCredentialsPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCredentialsPropertiesOutput)
}

func (i TokenCredentialsPropertiesArgs) ToTokenCredentialsPropertiesPtrOutput() TokenCredentialsPropertiesPtrOutput {
	return i.ToTokenCredentialsPropertiesPtrOutputWithContext(context.Background())
}

func (i TokenCredentialsPropertiesArgs) ToTokenCredentialsPropertiesPtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCredentialsPropertiesOutput).ToTokenCredentialsPropertiesPtrOutputWithContext(ctx)
}









type TokenCredentialsPropertiesPtrInput interface {
	pulumi.Input

	ToTokenCredentialsPropertiesPtrOutput() TokenCredentialsPropertiesPtrOutput
	ToTokenCredentialsPropertiesPtrOutputWithContext(context.Context) TokenCredentialsPropertiesPtrOutput
}

type tokenCredentialsPropertiesPtrType TokenCredentialsPropertiesArgs

func TokenCredentialsPropertiesPtr(v *TokenCredentialsPropertiesArgs) TokenCredentialsPropertiesPtrInput {
	return (*tokenCredentialsPropertiesPtrType)(v)
}

func (*tokenCredentialsPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenCredentialsProperties)(nil)).Elem()
}

func (i *tokenCredentialsPropertiesPtrType) ToTokenCredentialsPropertiesPtrOutput() TokenCredentialsPropertiesPtrOutput {
	return i.ToTokenCredentialsPropertiesPtrOutputWithContext(context.Background())
}

func (i *tokenCredentialsPropertiesPtrType) ToTokenCredentialsPropertiesPtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCredentialsPropertiesPtrOutput)
}

type TokenCredentialsPropertiesOutput struct{ *pulumi.OutputState }

func (TokenCredentialsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCredentialsProperties)(nil)).Elem()
}

func (o TokenCredentialsPropertiesOutput) ToTokenCredentialsPropertiesOutput() TokenCredentialsPropertiesOutput {
	return o
}

func (o TokenCredentialsPropertiesOutput) ToTokenCredentialsPropertiesOutputWithContext(ctx context.Context) TokenCredentialsPropertiesOutput {
	return o
}

func (o TokenCredentialsPropertiesOutput) ToTokenCredentialsPropertiesPtrOutput() TokenCredentialsPropertiesPtrOutput {
	return o.ToTokenCredentialsPropertiesPtrOutputWithContext(context.Background())
}

func (o TokenCredentialsPropertiesOutput) ToTokenCredentialsPropertiesPtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TokenCredentialsProperties) *TokenCredentialsProperties {
		return &v
	}).(TokenCredentialsPropertiesPtrOutput)
}

func (o TokenCredentialsPropertiesOutput) Certificates() TokenCertificateArrayOutput {
	return o.ApplyT(func(v TokenCredentialsProperties) []TokenCertificate { return v.Certificates }).(TokenCertificateArrayOutput)
}

func (o TokenCredentialsPropertiesOutput) Passwords() TokenPasswordArrayOutput {
	return o.ApplyT(func(v TokenCredentialsProperties) []TokenPassword { return v.Passwords }).(TokenPasswordArrayOutput)
}

type TokenCredentialsPropertiesPtrOutput struct{ *pulumi.OutputState }

func (TokenCredentialsPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenCredentialsProperties)(nil)).Elem()
}

func (o TokenCredentialsPropertiesPtrOutput) ToTokenCredentialsPropertiesPtrOutput() TokenCredentialsPropertiesPtrOutput {
	return o
}

func (o TokenCredentialsPropertiesPtrOutput) ToTokenCredentialsPropertiesPtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesPtrOutput {
	return o
}

func (o TokenCredentialsPropertiesPtrOutput) Elem() TokenCredentialsPropertiesOutput {
	return o.ApplyT(func(v *TokenCredentialsProperties) TokenCredentialsProperties {
		if v != nil {
			return *v
		}
		var ret TokenCredentialsProperties
		return ret
	}).(TokenCredentialsPropertiesOutput)
}

func (o TokenCredentialsPropertiesPtrOutput) Certificates() TokenCertificateArrayOutput {
	return o.ApplyT(func(v *TokenCredentialsProperties) []TokenCertificate {
		if v == nil {
			return nil
		}
		return v.Certificates
	}).(TokenCertificateArrayOutput)
}

func (o TokenCredentialsPropertiesPtrOutput) Passwords() TokenPasswordArrayOutput {
	return o.ApplyT(func(v *TokenCredentialsProperties) []TokenPassword {
		if v == nil {
			return nil
		}
		return v.Passwords
	}).(TokenPasswordArrayOutput)
}

type TokenCredentialsPropertiesResponse struct {
	Certificates []TokenCertificateResponse `pulumi:"certificates"`
	Passwords    []TokenPasswordResponse    `pulumi:"passwords"`
}

type TokenCredentialsPropertiesResponseOutput struct{ *pulumi.OutputState }

func (TokenCredentialsPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCredentialsPropertiesResponse)(nil)).Elem()
}

func (o TokenCredentialsPropertiesResponseOutput) ToTokenCredentialsPropertiesResponseOutput() TokenCredentialsPropertiesResponseOutput {
	return o
}

func (o TokenCredentialsPropertiesResponseOutput) ToTokenCredentialsPropertiesResponseOutputWithContext(ctx context.Context) TokenCredentialsPropertiesResponseOutput {
	return o
}

func (o TokenCredentialsPropertiesResponseOutput) Certificates() TokenCertificateResponseArrayOutput {
	return o.ApplyT(func(v TokenCredentialsPropertiesResponse) []TokenCertificateResponse { return v.Certificates }).(TokenCertificateResponseArrayOutput)
}

func (o TokenCredentialsPropertiesResponseOutput) Passwords() TokenPasswordResponseArrayOutput {
	return o.ApplyT(func(v TokenCredentialsPropertiesResponse) []TokenPasswordResponse { return v.Passwords }).(TokenPasswordResponseArrayOutput)
}

type TokenCredentialsPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (TokenCredentialsPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenCredentialsPropertiesResponse)(nil)).Elem()
}

func (o TokenCredentialsPropertiesResponsePtrOutput) ToTokenCredentialsPropertiesResponsePtrOutput() TokenCredentialsPropertiesResponsePtrOutput {
	return o
}

func (o TokenCredentialsPropertiesResponsePtrOutput) ToTokenCredentialsPropertiesResponsePtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesResponsePtrOutput {
	return o
}

func (o TokenCredentialsPropertiesResponsePtrOutput) Elem() TokenCredentialsPropertiesResponseOutput {
	return o.ApplyT(func(v *TokenCredentialsPropertiesResponse) TokenCredentialsPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret TokenCredentialsPropertiesResponse
		return ret
	}).(TokenCredentialsPropertiesResponseOutput)
}

func (o TokenCredentialsPropertiesResponsePtrOutput) Certificates() TokenCertificateResponseArrayOutput {
	return o.ApplyT(func(v *TokenCredentialsPropertiesResponse) []TokenCertificateResponse {
		if v == nil {
			return nil
		}
		return v.Certificates
	}).(TokenCertificateResponseArrayOutput)
}

func (o TokenCredentialsPropertiesResponsePtrOutput) Passwords() TokenPasswordResponseArrayOutput {
	return o.ApplyT(func(v *TokenCredentialsPropertiesResponse) []TokenPasswordResponse {
		if v == nil {
			return nil
		}
		return v.Passwords
	}).(TokenPasswordResponseArrayOutput)
}

type TokenPassword struct {
	CreationTime *string `pulumi:"creationTime"`
	Expiry       *string `pulumi:"expiry"`
	Name         *string `pulumi:"name"`
}





type TokenPasswordInput interface {
	pulumi.Input

	ToTokenPasswordOutput() TokenPasswordOutput
	ToTokenPasswordOutputWithContext(context.Context) TokenPasswordOutput
}

type TokenPasswordArgs struct {
	CreationTime pulumi.StringPtrInput `pulumi:"creationTime"`
	Expiry       pulumi.StringPtrInput `pulumi:"expiry"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
}

func (TokenPasswordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenPassword)(nil)).Elem()
}

func (i TokenPasswordArgs) ToTokenPasswordOutput() TokenPasswordOutput {
	return i.ToTokenPasswordOutputWithContext(context.Background())
}

func (i TokenPasswordArgs) ToTokenPasswordOutputWithContext(ctx context.Context) TokenPasswordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenPasswordOutput)
}





type TokenPasswordArrayInput interface {
	pulumi.Input

	ToTokenPasswordArrayOutput() TokenPasswordArrayOutput
	ToTokenPasswordArrayOutputWithContext(context.Context) TokenPasswordArrayOutput
}

type TokenPasswordArray []TokenPasswordInput

func (TokenPasswordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenPassword)(nil)).Elem()
}

func (i TokenPasswordArray) ToTokenPasswordArrayOutput() TokenPasswordArrayOutput {
	return i.ToTokenPasswordArrayOutputWithContext(context.Background())
}

func (i TokenPasswordArray) ToTokenPasswordArrayOutputWithContext(ctx context.Context) TokenPasswordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenPasswordArrayOutput)
}

type TokenPasswordOutput struct{ *pulumi.OutputState }

func (TokenPasswordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenPassword)(nil)).Elem()
}

func (o TokenPasswordOutput) ToTokenPasswordOutput() TokenPasswordOutput {
	return o
}

func (o TokenPasswordOutput) ToTokenPasswordOutputWithContext(ctx context.Context) TokenPasswordOutput {
	return o
}

func (o TokenPasswordOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenPassword) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o TokenPasswordOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenPassword) *string { return v.Expiry }).(pulumi.StringPtrOutput)
}

func (o TokenPasswordOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenPassword) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type TokenPasswordArrayOutput struct{ *pulumi.OutputState }

func (TokenPasswordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenPassword)(nil)).Elem()
}

func (o TokenPasswordArrayOutput) ToTokenPasswordArrayOutput() TokenPasswordArrayOutput {
	return o
}

func (o TokenPasswordArrayOutput) ToTokenPasswordArrayOutputWithContext(ctx context.Context) TokenPasswordArrayOutput {
	return o
}

func (o TokenPasswordArrayOutput) Index(i pulumi.IntInput) TokenPasswordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenPassword {
		return vs[0].([]TokenPassword)[vs[1].(int)]
	}).(TokenPasswordOutput)
}

type TokenPasswordResponse struct {
	CreationTime *string `pulumi:"creationTime"`
	Expiry       *string `pulumi:"expiry"`
	Name         *string `pulumi:"name"`
	Value        string  `pulumi:"value"`
}

type TokenPasswordResponseOutput struct{ *pulumi.OutputState }

func (TokenPasswordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenPasswordResponse)(nil)).Elem()
}

func (o TokenPasswordResponseOutput) ToTokenPasswordResponseOutput() TokenPasswordResponseOutput {
	return o
}

func (o TokenPasswordResponseOutput) ToTokenPasswordResponseOutputWithContext(ctx context.Context) TokenPasswordResponseOutput {
	return o
}

func (o TokenPasswordResponseOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenPasswordResponse) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o TokenPasswordResponseOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenPasswordResponse) *string { return v.Expiry }).(pulumi.StringPtrOutput)
}

func (o TokenPasswordResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenPasswordResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TokenPasswordResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TokenPasswordResponse) string { return v.Value }).(pulumi.StringOutput)
}

type TokenPasswordResponseArrayOutput struct{ *pulumi.OutputState }

func (TokenPasswordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenPasswordResponse)(nil)).Elem()
}

func (o TokenPasswordResponseArrayOutput) ToTokenPasswordResponseArrayOutput() TokenPasswordResponseArrayOutput {
	return o
}

func (o TokenPasswordResponseArrayOutput) ToTokenPasswordResponseArrayOutputWithContext(ctx context.Context) TokenPasswordResponseArrayOutput {
	return o
}

func (o TokenPasswordResponseArrayOutput) Index(i pulumi.IntInput) TokenPasswordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenPasswordResponse {
		return vs[0].([]TokenPasswordResponse)[vs[1].(int)]
	}).(TokenPasswordResponseOutput)
}

type TriggerProperties struct {
	BaseImageTrigger *BaseImageTrigger `pulumi:"baseImageTrigger"`
	SourceTriggers   []SourceTrigger   `pulumi:"sourceTriggers"`
	TimerTriggers    []TimerTrigger    `pulumi:"timerTriggers"`
}


func (val *TriggerProperties) Defaults() *TriggerProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.BaseImageTrigger = tmp.BaseImageTrigger.Defaults()

	return &tmp
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


func (val *TriggerPropertiesResponse) Defaults() *TriggerPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.BaseImageTrigger = tmp.BaseImageTrigger.Defaults()

	return &tmp
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

type TrustPolicy struct {
	Status *string `pulumi:"status"`
	Type   *string `pulumi:"type"`
}


func (val *TrustPolicy) Defaults() *TrustPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "disabled"
		tmp.Status = &status_
	}
	if isZero(tmp.Type) {
		type_ := "Notary"
		tmp.Type = &type_
	}
	return &tmp
}





type TrustPolicyInput interface {
	pulumi.Input

	ToTrustPolicyOutput() TrustPolicyOutput
	ToTrustPolicyOutputWithContext(context.Context) TrustPolicyOutput
}

type TrustPolicyArgs struct {
	Status pulumi.StringPtrInput `pulumi:"status"`
	Type   pulumi.StringPtrInput `pulumi:"type"`
}

func (TrustPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustPolicy)(nil)).Elem()
}

func (i TrustPolicyArgs) ToTrustPolicyOutput() TrustPolicyOutput {
	return i.ToTrustPolicyOutputWithContext(context.Background())
}

func (i TrustPolicyArgs) ToTrustPolicyOutputWithContext(ctx context.Context) TrustPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustPolicyOutput)
}

func (i TrustPolicyArgs) ToTrustPolicyPtrOutput() TrustPolicyPtrOutput {
	return i.ToTrustPolicyPtrOutputWithContext(context.Background())
}

func (i TrustPolicyArgs) ToTrustPolicyPtrOutputWithContext(ctx context.Context) TrustPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustPolicyOutput).ToTrustPolicyPtrOutputWithContext(ctx)
}









type TrustPolicyPtrInput interface {
	pulumi.Input

	ToTrustPolicyPtrOutput() TrustPolicyPtrOutput
	ToTrustPolicyPtrOutputWithContext(context.Context) TrustPolicyPtrOutput
}

type trustPolicyPtrType TrustPolicyArgs

func TrustPolicyPtr(v *TrustPolicyArgs) TrustPolicyPtrInput {
	return (*trustPolicyPtrType)(v)
}

func (*trustPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustPolicy)(nil)).Elem()
}

func (i *trustPolicyPtrType) ToTrustPolicyPtrOutput() TrustPolicyPtrOutput {
	return i.ToTrustPolicyPtrOutputWithContext(context.Background())
}

func (i *trustPolicyPtrType) ToTrustPolicyPtrOutputWithContext(ctx context.Context) TrustPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustPolicyPtrOutput)
}

type TrustPolicyOutput struct{ *pulumi.OutputState }

func (TrustPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustPolicy)(nil)).Elem()
}

func (o TrustPolicyOutput) ToTrustPolicyOutput() TrustPolicyOutput {
	return o
}

func (o TrustPolicyOutput) ToTrustPolicyOutputWithContext(ctx context.Context) TrustPolicyOutput {
	return o
}

func (o TrustPolicyOutput) ToTrustPolicyPtrOutput() TrustPolicyPtrOutput {
	return o.ToTrustPolicyPtrOutputWithContext(context.Background())
}

func (o TrustPolicyOutput) ToTrustPolicyPtrOutputWithContext(ctx context.Context) TrustPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TrustPolicy) *TrustPolicy {
		return &v
	}).(TrustPolicyPtrOutput)
}

func (o TrustPolicyOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrustPolicy) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o TrustPolicyOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrustPolicy) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type TrustPolicyPtrOutput struct{ *pulumi.OutputState }

func (TrustPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustPolicy)(nil)).Elem()
}

func (o TrustPolicyPtrOutput) ToTrustPolicyPtrOutput() TrustPolicyPtrOutput {
	return o
}

func (o TrustPolicyPtrOutput) ToTrustPolicyPtrOutputWithContext(ctx context.Context) TrustPolicyPtrOutput {
	return o
}

func (o TrustPolicyPtrOutput) Elem() TrustPolicyOutput {
	return o.ApplyT(func(v *TrustPolicy) TrustPolicy {
		if v != nil {
			return *v
		}
		var ret TrustPolicy
		return ret
	}).(TrustPolicyOutput)
}

func (o TrustPolicyPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrustPolicy) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o TrustPolicyPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrustPolicy) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type TrustPolicyResponse struct {
	Status *string `pulumi:"status"`
	Type   *string `pulumi:"type"`
}


func (val *TrustPolicyResponse) Defaults() *TrustPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "disabled"
		tmp.Status = &status_
	}
	if isZero(tmp.Type) {
		type_ := "Notary"
		tmp.Type = &type_
	}
	return &tmp
}

type TrustPolicyResponseOutput struct{ *pulumi.OutputState }

func (TrustPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustPolicyResponse)(nil)).Elem()
}

func (o TrustPolicyResponseOutput) ToTrustPolicyResponseOutput() TrustPolicyResponseOutput {
	return o
}

func (o TrustPolicyResponseOutput) ToTrustPolicyResponseOutputWithContext(ctx context.Context) TrustPolicyResponseOutput {
	return o
}

func (o TrustPolicyResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrustPolicyResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o TrustPolicyResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrustPolicyResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type TrustPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (TrustPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustPolicyResponse)(nil)).Elem()
}

func (o TrustPolicyResponsePtrOutput) ToTrustPolicyResponsePtrOutput() TrustPolicyResponsePtrOutput {
	return o
}

func (o TrustPolicyResponsePtrOutput) ToTrustPolicyResponsePtrOutputWithContext(ctx context.Context) TrustPolicyResponsePtrOutput {
	return o
}

func (o TrustPolicyResponsePtrOutput) Elem() TrustPolicyResponseOutput {
	return o.ApplyT(func(v *TrustPolicyResponse) TrustPolicyResponse {
		if v != nil {
			return *v
		}
		var ret TrustPolicyResponse
		return ret
	}).(TrustPolicyResponseOutput)
}

func (o TrustPolicyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrustPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o TrustPolicyResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrustPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
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

type VirtualNetworkRule struct {
	Action                   *string `pulumi:"action"`
	VirtualNetworkResourceId string  `pulumi:"virtualNetworkResourceId"`
}


func (val *VirtualNetworkRule) Defaults() *VirtualNetworkRule {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}





type VirtualNetworkRuleInput interface {
	pulumi.Input

	ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput
	ToVirtualNetworkRuleOutputWithContext(context.Context) VirtualNetworkRuleOutput
}

type VirtualNetworkRuleArgs struct {
	Action                   pulumi.StringPtrInput `pulumi:"action"`
	VirtualNetworkResourceId pulumi.StringInput    `pulumi:"virtualNetworkResourceId"`
}

func (VirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return i.ToVirtualNetworkRuleOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleOutput)
}





type VirtualNetworkRuleArrayInput interface {
	pulumi.Input

	ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput
	ToVirtualNetworkRuleArrayOutputWithContext(context.Context) VirtualNetworkRuleArrayOutput
}

type VirtualNetworkRuleArray []VirtualNetworkRuleInput

func (VirtualNetworkRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return i.ToVirtualNetworkRuleArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleArrayOutput)
}

type VirtualNetworkRuleOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkRuleOutput) VirtualNetworkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRule) string { return v.VirtualNetworkResourceId }).(pulumi.StringOutput)
}

type VirtualNetworkRuleArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRule {
		return vs[0].([]VirtualNetworkRule)[vs[1].(int)]
	}).(VirtualNetworkRuleOutput)
}

type VirtualNetworkRuleResponse struct {
	Action                   *string `pulumi:"action"`
	VirtualNetworkResourceId string  `pulumi:"virtualNetworkResourceId"`
}


func (val *VirtualNetworkRuleResponse) Defaults() *VirtualNetworkRuleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}

type VirtualNetworkRuleResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkRuleResponseOutput) VirtualNetworkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) string { return v.VirtualNetworkResourceId }).(pulumi.StringOutput)
}

type VirtualNetworkRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRuleResponse {
		return vs[0].([]VirtualNetworkRuleResponse)[vs[1].(int)]
	}).(VirtualNetworkRuleResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ActivationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AgentPropertiesOutput{})
	pulumi.RegisterOutputType(AgentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(AgentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AgentPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(AuthInfoOutput{})
	pulumi.RegisterOutputType(AuthInfoPtrOutput{})
	pulumi.RegisterOutputType(AuthInfoResponseOutput{})
	pulumi.RegisterOutputType(AuthInfoResponsePtrOutput{})
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
	pulumi.RegisterOutputType(ExportPipelineTargetPropertiesOutput{})
	pulumi.RegisterOutputType(ExportPipelineTargetPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ExportPipelineTargetPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(IPRuleOutput{})
	pulumi.RegisterOutputType(IPRuleArrayOutput{})
	pulumi.RegisterOutputType(IPRuleResponseOutput{})
	pulumi.RegisterOutputType(IPRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesPtrOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageDescriptorResponseOutput{})
	pulumi.RegisterOutputType(ImageDescriptorResponseArrayOutput{})
	pulumi.RegisterOutputType(ImageUpdateTriggerResponseOutput{})
	pulumi.RegisterOutputType(ImageUpdateTriggerResponsePtrOutput{})
	pulumi.RegisterOutputType(ImportPipelineSourcePropertiesOutput{})
	pulumi.RegisterOutputType(ImportPipelineSourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ImportPipelineSourcePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LoggingPropertiesOutput{})
	pulumi.RegisterOutputType(LoggingPropertiesPtrOutput{})
	pulumi.RegisterOutputType(LoggingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(LoggingPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LoginServerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(LoginServerPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetPtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponseOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponsePtrOutput{})
	pulumi.RegisterOutputType(ParentPropertiesOutput{})
	pulumi.RegisterOutputType(ParentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PipelineRunRequestOutput{})
	pulumi.RegisterOutputType(PipelineRunRequestPtrOutput{})
	pulumi.RegisterOutputType(PipelineRunRequestResponseOutput{})
	pulumi.RegisterOutputType(PipelineRunRequestResponsePtrOutput{})
	pulumi.RegisterOutputType(PipelineRunResponseResponseOutput{})
	pulumi.RegisterOutputType(PipelineRunSourcePropertiesOutput{})
	pulumi.RegisterOutputType(PipelineRunSourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(PipelineRunSourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(PipelineRunSourcePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PipelineRunTargetPropertiesOutput{})
	pulumi.RegisterOutputType(PipelineRunTargetPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PipelineRunTargetPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PipelineRunTargetPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PipelineSourceTriggerDescriptorResponseOutput{})
	pulumi.RegisterOutputType(PipelineSourceTriggerDescriptorResponsePtrOutput{})
	pulumi.RegisterOutputType(PipelineSourceTriggerPropertiesOutput{})
	pulumi.RegisterOutputType(PipelineSourceTriggerPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PipelineSourceTriggerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PipelineSourceTriggerPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PipelineTriggerDescriptorResponseOutput{})
	pulumi.RegisterOutputType(PipelineTriggerDescriptorResponsePtrOutput{})
	pulumi.RegisterOutputType(PipelineTriggerPropertiesOutput{})
	pulumi.RegisterOutputType(PipelineTriggerPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PipelineTriggerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PipelineTriggerPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PoliciesOutput{})
	pulumi.RegisterOutputType(PoliciesPtrOutput{})
	pulumi.RegisterOutputType(PoliciesResponseOutput{})
	pulumi.RegisterOutputType(PoliciesResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(ProgressPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ProgressPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(QuarantinePolicyOutput{})
	pulumi.RegisterOutputType(QuarantinePolicyPtrOutput{})
	pulumi.RegisterOutputType(QuarantinePolicyResponseOutput{})
	pulumi.RegisterOutputType(QuarantinePolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(RetentionPolicyOutput{})
	pulumi.RegisterOutputType(RetentionPolicyPtrOutput{})
	pulumi.RegisterOutputType(RetentionPolicyResponseOutput{})
	pulumi.RegisterOutputType(RetentionPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(RunResponseOutput{})
	pulumi.RegisterOutputType(SecretObjectOutput{})
	pulumi.RegisterOutputType(SecretObjectPtrOutput{})
	pulumi.RegisterOutputType(SecretObjectResponseOutput{})
	pulumi.RegisterOutputType(SecretObjectResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
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
	pulumi.RegisterOutputType(StatusDetailPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StatusDetailPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(StatusResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesPtrOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SyncPropertiesOutput{})
	pulumi.RegisterOutputType(SyncPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TimerTriggerOutput{})
	pulumi.RegisterOutputType(TimerTriggerArrayOutput{})
	pulumi.RegisterOutputType(TimerTriggerDescriptorResponseOutput{})
	pulumi.RegisterOutputType(TimerTriggerDescriptorResponsePtrOutput{})
	pulumi.RegisterOutputType(TimerTriggerResponseOutput{})
	pulumi.RegisterOutputType(TimerTriggerResponseArrayOutput{})
	pulumi.RegisterOutputType(TlsCertificatePropertiesResponseOutput{})
	pulumi.RegisterOutputType(TlsCertificatePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(TlsPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TlsPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(TokenCertificateOutput{})
	pulumi.RegisterOutputType(TokenCertificateArrayOutput{})
	pulumi.RegisterOutputType(TokenCertificateResponseOutput{})
	pulumi.RegisterOutputType(TokenCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(TokenCredentialsPropertiesOutput{})
	pulumi.RegisterOutputType(TokenCredentialsPropertiesPtrOutput{})
	pulumi.RegisterOutputType(TokenCredentialsPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TokenCredentialsPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(TokenPasswordOutput{})
	pulumi.RegisterOutputType(TokenPasswordArrayOutput{})
	pulumi.RegisterOutputType(TokenPasswordResponseOutput{})
	pulumi.RegisterOutputType(TokenPasswordResponseArrayOutput{})
	pulumi.RegisterOutputType(TriggerPropertiesOutput{})
	pulumi.RegisterOutputType(TriggerPropertiesPtrOutput{})
	pulumi.RegisterOutputType(TriggerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TriggerPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(TrustPolicyOutput{})
	pulumi.RegisterOutputType(TrustPolicyPtrOutput{})
	pulumi.RegisterOutputType(TrustPolicyResponseOutput{})
	pulumi.RegisterOutputType(TrustPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesMapOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesResponseMapOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseArrayOutput{})
}

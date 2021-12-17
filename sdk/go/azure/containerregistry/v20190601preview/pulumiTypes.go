


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

func init() {
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
	pulumi.RegisterOutputType(IdentityPropertiesOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesPtrOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageDescriptorResponseOutput{})
	pulumi.RegisterOutputType(ImageDescriptorResponseArrayOutput{})
	pulumi.RegisterOutputType(ImageUpdateTriggerResponseOutput{})
	pulumi.RegisterOutputType(ImageUpdateTriggerResponsePtrOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(RunResponseOutput{})
	pulumi.RegisterOutputType(SecretObjectOutput{})
	pulumi.RegisterOutputType(SecretObjectPtrOutput{})
	pulumi.RegisterOutputType(SecretObjectResponseOutput{})
	pulumi.RegisterOutputType(SecretObjectResponsePtrOutput{})
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

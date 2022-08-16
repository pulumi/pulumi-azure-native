


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccountEncryption struct {
	Identity           *ResourceIdentity   `pulumi:"identity"`
	KeyVaultProperties *KeyVaultProperties `pulumi:"keyVaultProperties"`
	Type               string              `pulumi:"type"`
}





type AccountEncryptionInput interface {
	pulumi.Input

	ToAccountEncryptionOutput() AccountEncryptionOutput
	ToAccountEncryptionOutputWithContext(context.Context) AccountEncryptionOutput
}

type AccountEncryptionArgs struct {
	Identity           ResourceIdentityPtrInput   `pulumi:"identity"`
	KeyVaultProperties KeyVaultPropertiesPtrInput `pulumi:"keyVaultProperties"`
	Type               pulumi.StringInput         `pulumi:"type"`
}

func (AccountEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountEncryption)(nil)).Elem()
}

func (i AccountEncryptionArgs) ToAccountEncryptionOutput() AccountEncryptionOutput {
	return i.ToAccountEncryptionOutputWithContext(context.Background())
}

func (i AccountEncryptionArgs) ToAccountEncryptionOutputWithContext(ctx context.Context) AccountEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountEncryptionOutput)
}

func (i AccountEncryptionArgs) ToAccountEncryptionPtrOutput() AccountEncryptionPtrOutput {
	return i.ToAccountEncryptionPtrOutputWithContext(context.Background())
}

func (i AccountEncryptionArgs) ToAccountEncryptionPtrOutputWithContext(ctx context.Context) AccountEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountEncryptionOutput).ToAccountEncryptionPtrOutputWithContext(ctx)
}









type AccountEncryptionPtrInput interface {
	pulumi.Input

	ToAccountEncryptionPtrOutput() AccountEncryptionPtrOutput
	ToAccountEncryptionPtrOutputWithContext(context.Context) AccountEncryptionPtrOutput
}

type accountEncryptionPtrType AccountEncryptionArgs

func AccountEncryptionPtr(v *AccountEncryptionArgs) AccountEncryptionPtrInput {
	return (*accountEncryptionPtrType)(v)
}

func (*accountEncryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountEncryption)(nil)).Elem()
}

func (i *accountEncryptionPtrType) ToAccountEncryptionPtrOutput() AccountEncryptionPtrOutput {
	return i.ToAccountEncryptionPtrOutputWithContext(context.Background())
}

func (i *accountEncryptionPtrType) ToAccountEncryptionPtrOutputWithContext(ctx context.Context) AccountEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountEncryptionPtrOutput)
}

type AccountEncryptionOutput struct{ *pulumi.OutputState }

func (AccountEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountEncryption)(nil)).Elem()
}

func (o AccountEncryptionOutput) ToAccountEncryptionOutput() AccountEncryptionOutput {
	return o
}

func (o AccountEncryptionOutput) ToAccountEncryptionOutputWithContext(ctx context.Context) AccountEncryptionOutput {
	return o
}

func (o AccountEncryptionOutput) ToAccountEncryptionPtrOutput() AccountEncryptionPtrOutput {
	return o.ToAccountEncryptionPtrOutputWithContext(context.Background())
}

func (o AccountEncryptionOutput) ToAccountEncryptionPtrOutputWithContext(ctx context.Context) AccountEncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccountEncryption) *AccountEncryption {
		return &v
	}).(AccountEncryptionPtrOutput)
}

func (o AccountEncryptionOutput) Identity() ResourceIdentityPtrOutput {
	return o.ApplyT(func(v AccountEncryption) *ResourceIdentity { return v.Identity }).(ResourceIdentityPtrOutput)
}

func (o AccountEncryptionOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v AccountEncryption) *KeyVaultProperties { return v.KeyVaultProperties }).(KeyVaultPropertiesPtrOutput)
}

func (o AccountEncryptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AccountEncryption) string { return v.Type }).(pulumi.StringOutput)
}

type AccountEncryptionPtrOutput struct{ *pulumi.OutputState }

func (AccountEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountEncryption)(nil)).Elem()
}

func (o AccountEncryptionPtrOutput) ToAccountEncryptionPtrOutput() AccountEncryptionPtrOutput {
	return o
}

func (o AccountEncryptionPtrOutput) ToAccountEncryptionPtrOutputWithContext(ctx context.Context) AccountEncryptionPtrOutput {
	return o
}

func (o AccountEncryptionPtrOutput) Elem() AccountEncryptionOutput {
	return o.ApplyT(func(v *AccountEncryption) AccountEncryption {
		if v != nil {
			return *v
		}
		var ret AccountEncryption
		return ret
	}).(AccountEncryptionOutput)
}

func (o AccountEncryptionPtrOutput) Identity() ResourceIdentityPtrOutput {
	return o.ApplyT(func(v *AccountEncryption) *ResourceIdentity {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(ResourceIdentityPtrOutput)
}

func (o AccountEncryptionPtrOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *AccountEncryption) *KeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesPtrOutput)
}

func (o AccountEncryptionPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountEncryption) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type AccountEncryptionResponse struct {
	Identity           *ResourceIdentityResponse   `pulumi:"identity"`
	KeyVaultProperties *KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	Status             string                      `pulumi:"status"`
	Type               string                      `pulumi:"type"`
}

type AccountEncryptionResponseOutput struct{ *pulumi.OutputState }

func (AccountEncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountEncryptionResponse)(nil)).Elem()
}

func (o AccountEncryptionResponseOutput) ToAccountEncryptionResponseOutput() AccountEncryptionResponseOutput {
	return o
}

func (o AccountEncryptionResponseOutput) ToAccountEncryptionResponseOutputWithContext(ctx context.Context) AccountEncryptionResponseOutput {
	return o
}

func (o AccountEncryptionResponseOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v AccountEncryptionResponse) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o AccountEncryptionResponseOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v AccountEncryptionResponse) *KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponsePtrOutput)
}

func (o AccountEncryptionResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v AccountEncryptionResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o AccountEncryptionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AccountEncryptionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type AccountEncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (AccountEncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountEncryptionResponse)(nil)).Elem()
}

func (o AccountEncryptionResponsePtrOutput) ToAccountEncryptionResponsePtrOutput() AccountEncryptionResponsePtrOutput {
	return o
}

func (o AccountEncryptionResponsePtrOutput) ToAccountEncryptionResponsePtrOutputWithContext(ctx context.Context) AccountEncryptionResponsePtrOutput {
	return o
}

func (o AccountEncryptionResponsePtrOutput) Elem() AccountEncryptionResponseOutput {
	return o.ApplyT(func(v *AccountEncryptionResponse) AccountEncryptionResponse {
		if v != nil {
			return *v
		}
		var ret AccountEncryptionResponse
		return ret
	}).(AccountEncryptionResponseOutput)
}

func (o AccountEncryptionResponsePtrOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *AccountEncryptionResponse) *ResourceIdentityResponse {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(ResourceIdentityResponsePtrOutput)
}

func (o AccountEncryptionResponsePtrOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *AccountEncryptionResponse) *KeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesResponsePtrOutput)
}

func (o AccountEncryptionResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountEncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

func (o AccountEncryptionResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountEncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type AudioEncoderAac struct {
	BitrateKbps *string `pulumi:"bitrateKbps"`
	Type        string  `pulumi:"type"`
}

type AudioEncoderAacResponse struct {
	BitrateKbps *string `pulumi:"bitrateKbps"`
	Type        string  `pulumi:"type"`
}

type EccTokenKey struct {
	Alg  string `pulumi:"alg"`
	Kid  string `pulumi:"kid"`
	Type string `pulumi:"type"`
	X    string `pulumi:"x"`
	Y    string `pulumi:"y"`
}

type EccTokenKeyResponse struct {
	Alg  string `pulumi:"alg"`
	Kid  string `pulumi:"kid"`
	Type string `pulumi:"type"`
	X    string `pulumi:"x"`
	Y    string `pulumi:"y"`
}

type EncoderCustomPreset struct {
	AudioEncoder *AudioEncoderAac  `pulumi:"audioEncoder"`
	Type         string            `pulumi:"type"`
	VideoEncoder *VideoEncoderH264 `pulumi:"videoEncoder"`
}

type EncoderCustomPresetResponse struct {
	AudioEncoder *AudioEncoderAacResponse  `pulumi:"audioEncoder"`
	Type         string                    `pulumi:"type"`
	VideoEncoder *VideoEncoderH264Response `pulumi:"videoEncoder"`
}

type EncoderProcessor struct {
	Inputs []NodeInput `pulumi:"inputs"`
	Name   string      `pulumi:"name"`
	Preset interface{} `pulumi:"preset"`
	Type   string      `pulumi:"type"`
}





type EncoderProcessorInput interface {
	pulumi.Input

	ToEncoderProcessorOutput() EncoderProcessorOutput
	ToEncoderProcessorOutputWithContext(context.Context) EncoderProcessorOutput
}

type EncoderProcessorArgs struct {
	Inputs NodeInputArrayInput `pulumi:"inputs"`
	Name   pulumi.StringInput  `pulumi:"name"`
	Preset pulumi.Input        `pulumi:"preset"`
	Type   pulumi.StringInput  `pulumi:"type"`
}

func (EncoderProcessorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncoderProcessor)(nil)).Elem()
}

func (i EncoderProcessorArgs) ToEncoderProcessorOutput() EncoderProcessorOutput {
	return i.ToEncoderProcessorOutputWithContext(context.Background())
}

func (i EncoderProcessorArgs) ToEncoderProcessorOutputWithContext(ctx context.Context) EncoderProcessorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncoderProcessorOutput)
}





type EncoderProcessorArrayInput interface {
	pulumi.Input

	ToEncoderProcessorArrayOutput() EncoderProcessorArrayOutput
	ToEncoderProcessorArrayOutputWithContext(context.Context) EncoderProcessorArrayOutput
}

type EncoderProcessorArray []EncoderProcessorInput

func (EncoderProcessorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EncoderProcessor)(nil)).Elem()
}

func (i EncoderProcessorArray) ToEncoderProcessorArrayOutput() EncoderProcessorArrayOutput {
	return i.ToEncoderProcessorArrayOutputWithContext(context.Background())
}

func (i EncoderProcessorArray) ToEncoderProcessorArrayOutputWithContext(ctx context.Context) EncoderProcessorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncoderProcessorArrayOutput)
}

type EncoderProcessorOutput struct{ *pulumi.OutputState }

func (EncoderProcessorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncoderProcessor)(nil)).Elem()
}

func (o EncoderProcessorOutput) ToEncoderProcessorOutput() EncoderProcessorOutput {
	return o
}

func (o EncoderProcessorOutput) ToEncoderProcessorOutputWithContext(ctx context.Context) EncoderProcessorOutput {
	return o
}

func (o EncoderProcessorOutput) Inputs() NodeInputArrayOutput {
	return o.ApplyT(func(v EncoderProcessor) []NodeInput { return v.Inputs }).(NodeInputArrayOutput)
}

func (o EncoderProcessorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EncoderProcessor) string { return v.Name }).(pulumi.StringOutput)
}

func (o EncoderProcessorOutput) Preset() pulumi.AnyOutput {
	return o.ApplyT(func(v EncoderProcessor) interface{} { return v.Preset }).(pulumi.AnyOutput)
}

func (o EncoderProcessorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EncoderProcessor) string { return v.Type }).(pulumi.StringOutput)
}

type EncoderProcessorArrayOutput struct{ *pulumi.OutputState }

func (EncoderProcessorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EncoderProcessor)(nil)).Elem()
}

func (o EncoderProcessorArrayOutput) ToEncoderProcessorArrayOutput() EncoderProcessorArrayOutput {
	return o
}

func (o EncoderProcessorArrayOutput) ToEncoderProcessorArrayOutputWithContext(ctx context.Context) EncoderProcessorArrayOutput {
	return o
}

func (o EncoderProcessorArrayOutput) Index(i pulumi.IntInput) EncoderProcessorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EncoderProcessor {
		return vs[0].([]EncoderProcessor)[vs[1].(int)]
	}).(EncoderProcessorOutput)
}

type EncoderProcessorResponse struct {
	Inputs []NodeInputResponse `pulumi:"inputs"`
	Name   string              `pulumi:"name"`
	Preset interface{}         `pulumi:"preset"`
	Type   string              `pulumi:"type"`
}

type EncoderProcessorResponseOutput struct{ *pulumi.OutputState }

func (EncoderProcessorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncoderProcessorResponse)(nil)).Elem()
}

func (o EncoderProcessorResponseOutput) ToEncoderProcessorResponseOutput() EncoderProcessorResponseOutput {
	return o
}

func (o EncoderProcessorResponseOutput) ToEncoderProcessorResponseOutputWithContext(ctx context.Context) EncoderProcessorResponseOutput {
	return o
}

func (o EncoderProcessorResponseOutput) Inputs() NodeInputResponseArrayOutput {
	return o.ApplyT(func(v EncoderProcessorResponse) []NodeInputResponse { return v.Inputs }).(NodeInputResponseArrayOutput)
}

func (o EncoderProcessorResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EncoderProcessorResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o EncoderProcessorResponseOutput) Preset() pulumi.AnyOutput {
	return o.ApplyT(func(v EncoderProcessorResponse) interface{} { return v.Preset }).(pulumi.AnyOutput)
}

func (o EncoderProcessorResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EncoderProcessorResponse) string { return v.Type }).(pulumi.StringOutput)
}

type EncoderProcessorResponseArrayOutput struct{ *pulumi.OutputState }

func (EncoderProcessorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EncoderProcessorResponse)(nil)).Elem()
}

func (o EncoderProcessorResponseArrayOutput) ToEncoderProcessorResponseArrayOutput() EncoderProcessorResponseArrayOutput {
	return o
}

func (o EncoderProcessorResponseArrayOutput) ToEncoderProcessorResponseArrayOutputWithContext(ctx context.Context) EncoderProcessorResponseArrayOutput {
	return o
}

func (o EncoderProcessorResponseArrayOutput) Index(i pulumi.IntInput) EncoderProcessorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EncoderProcessorResponse {
		return vs[0].([]EncoderProcessorResponse)[vs[1].(int)]
	}).(EncoderProcessorResponseOutput)
}

type EncoderSystemPreset struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}

type EncoderSystemPresetResponse struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}

type EndpointResponse struct {
	EndpointUrl *string `pulumi:"endpointUrl"`
	Type        string  `pulumi:"type"`
}

type EndpointResponseOutput struct{ *pulumi.OutputState }

func (EndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointResponse)(nil)).Elem()
}

func (o EndpointResponseOutput) ToEndpointResponseOutput() EndpointResponseOutput {
	return o
}

func (o EndpointResponseOutput) ToEndpointResponseOutputWithContext(ctx context.Context) EndpointResponseOutput {
	return o
}

func (o EndpointResponseOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointResponse) string { return v.Type }).(pulumi.StringOutput)
}

type EndpointResponseArrayOutput struct{ *pulumi.OutputState }

func (EndpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointResponse)(nil)).Elem()
}

func (o EndpointResponseArrayOutput) ToEndpointResponseArrayOutput() EndpointResponseArrayOutput {
	return o
}

func (o EndpointResponseArrayOutput) ToEndpointResponseArrayOutputWithContext(ctx context.Context) EndpointResponseArrayOutput {
	return o
}

func (o EndpointResponseArrayOutput) Index(i pulumi.IntInput) EndpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointResponse {
		return vs[0].([]EndpointResponse)[vs[1].(int)]
	}).(EndpointResponseOutput)
}

type GroupLevelAccessControl struct {
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
}





type GroupLevelAccessControlInput interface {
	pulumi.Input

	ToGroupLevelAccessControlOutput() GroupLevelAccessControlOutput
	ToGroupLevelAccessControlOutputWithContext(context.Context) GroupLevelAccessControlOutput
}

type GroupLevelAccessControlArgs struct {
	PublicNetworkAccess pulumi.StringPtrInput `pulumi:"publicNetworkAccess"`
}

func (GroupLevelAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupLevelAccessControl)(nil)).Elem()
}

func (i GroupLevelAccessControlArgs) ToGroupLevelAccessControlOutput() GroupLevelAccessControlOutput {
	return i.ToGroupLevelAccessControlOutputWithContext(context.Background())
}

func (i GroupLevelAccessControlArgs) ToGroupLevelAccessControlOutputWithContext(ctx context.Context) GroupLevelAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupLevelAccessControlOutput)
}

func (i GroupLevelAccessControlArgs) ToGroupLevelAccessControlPtrOutput() GroupLevelAccessControlPtrOutput {
	return i.ToGroupLevelAccessControlPtrOutputWithContext(context.Background())
}

func (i GroupLevelAccessControlArgs) ToGroupLevelAccessControlPtrOutputWithContext(ctx context.Context) GroupLevelAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupLevelAccessControlOutput).ToGroupLevelAccessControlPtrOutputWithContext(ctx)
}









type GroupLevelAccessControlPtrInput interface {
	pulumi.Input

	ToGroupLevelAccessControlPtrOutput() GroupLevelAccessControlPtrOutput
	ToGroupLevelAccessControlPtrOutputWithContext(context.Context) GroupLevelAccessControlPtrOutput
}

type groupLevelAccessControlPtrType GroupLevelAccessControlArgs

func GroupLevelAccessControlPtr(v *GroupLevelAccessControlArgs) GroupLevelAccessControlPtrInput {
	return (*groupLevelAccessControlPtrType)(v)
}

func (*groupLevelAccessControlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupLevelAccessControl)(nil)).Elem()
}

func (i *groupLevelAccessControlPtrType) ToGroupLevelAccessControlPtrOutput() GroupLevelAccessControlPtrOutput {
	return i.ToGroupLevelAccessControlPtrOutputWithContext(context.Background())
}

func (i *groupLevelAccessControlPtrType) ToGroupLevelAccessControlPtrOutputWithContext(ctx context.Context) GroupLevelAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupLevelAccessControlPtrOutput)
}

type GroupLevelAccessControlOutput struct{ *pulumi.OutputState }

func (GroupLevelAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupLevelAccessControl)(nil)).Elem()
}

func (o GroupLevelAccessControlOutput) ToGroupLevelAccessControlOutput() GroupLevelAccessControlOutput {
	return o
}

func (o GroupLevelAccessControlOutput) ToGroupLevelAccessControlOutputWithContext(ctx context.Context) GroupLevelAccessControlOutput {
	return o
}

func (o GroupLevelAccessControlOutput) ToGroupLevelAccessControlPtrOutput() GroupLevelAccessControlPtrOutput {
	return o.ToGroupLevelAccessControlPtrOutputWithContext(context.Background())
}

func (o GroupLevelAccessControlOutput) ToGroupLevelAccessControlPtrOutputWithContext(ctx context.Context) GroupLevelAccessControlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GroupLevelAccessControl) *GroupLevelAccessControl {
		return &v
	}).(GroupLevelAccessControlPtrOutput)
}

func (o GroupLevelAccessControlOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupLevelAccessControl) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

type GroupLevelAccessControlPtrOutput struct{ *pulumi.OutputState }

func (GroupLevelAccessControlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupLevelAccessControl)(nil)).Elem()
}

func (o GroupLevelAccessControlPtrOutput) ToGroupLevelAccessControlPtrOutput() GroupLevelAccessControlPtrOutput {
	return o
}

func (o GroupLevelAccessControlPtrOutput) ToGroupLevelAccessControlPtrOutputWithContext(ctx context.Context) GroupLevelAccessControlPtrOutput {
	return o
}

func (o GroupLevelAccessControlPtrOutput) Elem() GroupLevelAccessControlOutput {
	return o.ApplyT(func(v *GroupLevelAccessControl) GroupLevelAccessControl {
		if v != nil {
			return *v
		}
		var ret GroupLevelAccessControl
		return ret
	}).(GroupLevelAccessControlOutput)
}

func (o GroupLevelAccessControlPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupLevelAccessControl) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

type GroupLevelAccessControlResponse struct {
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
}

type GroupLevelAccessControlResponseOutput struct{ *pulumi.OutputState }

func (GroupLevelAccessControlResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupLevelAccessControlResponse)(nil)).Elem()
}

func (o GroupLevelAccessControlResponseOutput) ToGroupLevelAccessControlResponseOutput() GroupLevelAccessControlResponseOutput {
	return o
}

func (o GroupLevelAccessControlResponseOutput) ToGroupLevelAccessControlResponseOutputWithContext(ctx context.Context) GroupLevelAccessControlResponseOutput {
	return o
}

func (o GroupLevelAccessControlResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupLevelAccessControlResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

type GroupLevelAccessControlResponsePtrOutput struct{ *pulumi.OutputState }

func (GroupLevelAccessControlResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupLevelAccessControlResponse)(nil)).Elem()
}

func (o GroupLevelAccessControlResponsePtrOutput) ToGroupLevelAccessControlResponsePtrOutput() GroupLevelAccessControlResponsePtrOutput {
	return o
}

func (o GroupLevelAccessControlResponsePtrOutput) ToGroupLevelAccessControlResponsePtrOutputWithContext(ctx context.Context) GroupLevelAccessControlResponsePtrOutput {
	return o
}

func (o GroupLevelAccessControlResponsePtrOutput) Elem() GroupLevelAccessControlResponseOutput {
	return o.ApplyT(func(v *GroupLevelAccessControlResponse) GroupLevelAccessControlResponse {
		if v != nil {
			return *v
		}
		var ret GroupLevelAccessControlResponse
		return ret
	}).(GroupLevelAccessControlResponseOutput)
}

func (o GroupLevelAccessControlResponsePtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupLevelAccessControlResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

type IotHub struct {
	Id       string           `pulumi:"id"`
	Identity ResourceIdentity `pulumi:"identity"`
}





type IotHubInput interface {
	pulumi.Input

	ToIotHubOutput() IotHubOutput
	ToIotHubOutputWithContext(context.Context) IotHubOutput
}

type IotHubArgs struct {
	Id       pulumi.StringInput    `pulumi:"id"`
	Identity ResourceIdentityInput `pulumi:"identity"`
}

func (IotHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHub)(nil)).Elem()
}

func (i IotHubArgs) ToIotHubOutput() IotHubOutput {
	return i.ToIotHubOutputWithContext(context.Background())
}

func (i IotHubArgs) ToIotHubOutputWithContext(ctx context.Context) IotHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubOutput)
}





type IotHubArrayInput interface {
	pulumi.Input

	ToIotHubArrayOutput() IotHubArrayOutput
	ToIotHubArrayOutputWithContext(context.Context) IotHubArrayOutput
}

type IotHubArray []IotHubInput

func (IotHubArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IotHub)(nil)).Elem()
}

func (i IotHubArray) ToIotHubArrayOutput() IotHubArrayOutput {
	return i.ToIotHubArrayOutputWithContext(context.Background())
}

func (i IotHubArray) ToIotHubArrayOutputWithContext(ctx context.Context) IotHubArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubArrayOutput)
}

type IotHubOutput struct{ *pulumi.OutputState }

func (IotHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHub)(nil)).Elem()
}

func (o IotHubOutput) ToIotHubOutput() IotHubOutput {
	return o
}

func (o IotHubOutput) ToIotHubOutputWithContext(ctx context.Context) IotHubOutput {
	return o
}

func (o IotHubOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v IotHub) string { return v.Id }).(pulumi.StringOutput)
}

func (o IotHubOutput) Identity() ResourceIdentityOutput {
	return o.ApplyT(func(v IotHub) ResourceIdentity { return v.Identity }).(ResourceIdentityOutput)
}

type IotHubArrayOutput struct{ *pulumi.OutputState }

func (IotHubArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IotHub)(nil)).Elem()
}

func (o IotHubArrayOutput) ToIotHubArrayOutput() IotHubArrayOutput {
	return o
}

func (o IotHubArrayOutput) ToIotHubArrayOutputWithContext(ctx context.Context) IotHubArrayOutput {
	return o
}

func (o IotHubArrayOutput) Index(i pulumi.IntInput) IotHubOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IotHub {
		return vs[0].([]IotHub)[vs[1].(int)]
	}).(IotHubOutput)
}

type IotHubResponse struct {
	Id       string                   `pulumi:"id"`
	Identity ResourceIdentityResponse `pulumi:"identity"`
	Status   string                   `pulumi:"status"`
}

type IotHubResponseOutput struct{ *pulumi.OutputState }

func (IotHubResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubResponse)(nil)).Elem()
}

func (o IotHubResponseOutput) ToIotHubResponseOutput() IotHubResponseOutput {
	return o
}

func (o IotHubResponseOutput) ToIotHubResponseOutputWithContext(ctx context.Context) IotHubResponseOutput {
	return o
}

func (o IotHubResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o IotHubResponseOutput) Identity() ResourceIdentityResponseOutput {
	return o.ApplyT(func(v IotHubResponse) ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponseOutput)
}

func (o IotHubResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubResponse) string { return v.Status }).(pulumi.StringOutput)
}

type IotHubResponseArrayOutput struct{ *pulumi.OutputState }

func (IotHubResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IotHubResponse)(nil)).Elem()
}

func (o IotHubResponseArrayOutput) ToIotHubResponseArrayOutput() IotHubResponseArrayOutput {
	return o
}

func (o IotHubResponseArrayOutput) ToIotHubResponseArrayOutputWithContext(ctx context.Context) IotHubResponseArrayOutput {
	return o
}

func (o IotHubResponseArrayOutput) Index(i pulumi.IntInput) IotHubResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IotHubResponse {
		return vs[0].([]IotHubResponse)[vs[1].(int)]
	}).(IotHubResponseOutput)
}

type JwtAuthentication struct {
	Audiences []string      `pulumi:"audiences"`
	Claims    []TokenClaim  `pulumi:"claims"`
	Issuers   []string      `pulumi:"issuers"`
	Keys      []interface{} `pulumi:"keys"`
	Type      string        `pulumi:"type"`
}





type JwtAuthenticationInput interface {
	pulumi.Input

	ToJwtAuthenticationOutput() JwtAuthenticationOutput
	ToJwtAuthenticationOutputWithContext(context.Context) JwtAuthenticationOutput
}

type JwtAuthenticationArgs struct {
	Audiences pulumi.StringArrayInput `pulumi:"audiences"`
	Claims    TokenClaimArrayInput    `pulumi:"claims"`
	Issuers   pulumi.StringArrayInput `pulumi:"issuers"`
	Keys      pulumi.ArrayInput       `pulumi:"keys"`
	Type      pulumi.StringInput      `pulumi:"type"`
}

func (JwtAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JwtAuthentication)(nil)).Elem()
}

func (i JwtAuthenticationArgs) ToJwtAuthenticationOutput() JwtAuthenticationOutput {
	return i.ToJwtAuthenticationOutputWithContext(context.Background())
}

func (i JwtAuthenticationArgs) ToJwtAuthenticationOutputWithContext(ctx context.Context) JwtAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JwtAuthenticationOutput)
}

func (i JwtAuthenticationArgs) ToJwtAuthenticationPtrOutput() JwtAuthenticationPtrOutput {
	return i.ToJwtAuthenticationPtrOutputWithContext(context.Background())
}

func (i JwtAuthenticationArgs) ToJwtAuthenticationPtrOutputWithContext(ctx context.Context) JwtAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JwtAuthenticationOutput).ToJwtAuthenticationPtrOutputWithContext(ctx)
}









type JwtAuthenticationPtrInput interface {
	pulumi.Input

	ToJwtAuthenticationPtrOutput() JwtAuthenticationPtrOutput
	ToJwtAuthenticationPtrOutputWithContext(context.Context) JwtAuthenticationPtrOutput
}

type jwtAuthenticationPtrType JwtAuthenticationArgs

func JwtAuthenticationPtr(v *JwtAuthenticationArgs) JwtAuthenticationPtrInput {
	return (*jwtAuthenticationPtrType)(v)
}

func (*jwtAuthenticationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JwtAuthentication)(nil)).Elem()
}

func (i *jwtAuthenticationPtrType) ToJwtAuthenticationPtrOutput() JwtAuthenticationPtrOutput {
	return i.ToJwtAuthenticationPtrOutputWithContext(context.Background())
}

func (i *jwtAuthenticationPtrType) ToJwtAuthenticationPtrOutputWithContext(ctx context.Context) JwtAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JwtAuthenticationPtrOutput)
}

type JwtAuthenticationOutput struct{ *pulumi.OutputState }

func (JwtAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JwtAuthentication)(nil)).Elem()
}

func (o JwtAuthenticationOutput) ToJwtAuthenticationOutput() JwtAuthenticationOutput {
	return o
}

func (o JwtAuthenticationOutput) ToJwtAuthenticationOutputWithContext(ctx context.Context) JwtAuthenticationOutput {
	return o
}

func (o JwtAuthenticationOutput) ToJwtAuthenticationPtrOutput() JwtAuthenticationPtrOutput {
	return o.ToJwtAuthenticationPtrOutputWithContext(context.Background())
}

func (o JwtAuthenticationOutput) ToJwtAuthenticationPtrOutputWithContext(ctx context.Context) JwtAuthenticationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JwtAuthentication) *JwtAuthentication {
		return &v
	}).(JwtAuthenticationPtrOutput)
}

func (o JwtAuthenticationOutput) Audiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JwtAuthentication) []string { return v.Audiences }).(pulumi.StringArrayOutput)
}

func (o JwtAuthenticationOutput) Claims() TokenClaimArrayOutput {
	return o.ApplyT(func(v JwtAuthentication) []TokenClaim { return v.Claims }).(TokenClaimArrayOutput)
}

func (o JwtAuthenticationOutput) Issuers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JwtAuthentication) []string { return v.Issuers }).(pulumi.StringArrayOutput)
}

func (o JwtAuthenticationOutput) Keys() pulumi.ArrayOutput {
	return o.ApplyT(func(v JwtAuthentication) []interface{} { return v.Keys }).(pulumi.ArrayOutput)
}

func (o JwtAuthenticationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v JwtAuthentication) string { return v.Type }).(pulumi.StringOutput)
}

type JwtAuthenticationPtrOutput struct{ *pulumi.OutputState }

func (JwtAuthenticationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JwtAuthentication)(nil)).Elem()
}

func (o JwtAuthenticationPtrOutput) ToJwtAuthenticationPtrOutput() JwtAuthenticationPtrOutput {
	return o
}

func (o JwtAuthenticationPtrOutput) ToJwtAuthenticationPtrOutputWithContext(ctx context.Context) JwtAuthenticationPtrOutput {
	return o
}

func (o JwtAuthenticationPtrOutput) Elem() JwtAuthenticationOutput {
	return o.ApplyT(func(v *JwtAuthentication) JwtAuthentication {
		if v != nil {
			return *v
		}
		var ret JwtAuthentication
		return ret
	}).(JwtAuthenticationOutput)
}

func (o JwtAuthenticationPtrOutput) Audiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JwtAuthentication) []string {
		if v == nil {
			return nil
		}
		return v.Audiences
	}).(pulumi.StringArrayOutput)
}

func (o JwtAuthenticationPtrOutput) Claims() TokenClaimArrayOutput {
	return o.ApplyT(func(v *JwtAuthentication) []TokenClaim {
		if v == nil {
			return nil
		}
		return v.Claims
	}).(TokenClaimArrayOutput)
}

func (o JwtAuthenticationPtrOutput) Issuers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JwtAuthentication) []string {
		if v == nil {
			return nil
		}
		return v.Issuers
	}).(pulumi.StringArrayOutput)
}

func (o JwtAuthenticationPtrOutput) Keys() pulumi.ArrayOutput {
	return o.ApplyT(func(v *JwtAuthentication) []interface{} {
		if v == nil {
			return nil
		}
		return v.Keys
	}).(pulumi.ArrayOutput)
}

func (o JwtAuthenticationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JwtAuthentication) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type JwtAuthenticationResponse struct {
	Audiences []string             `pulumi:"audiences"`
	Claims    []TokenClaimResponse `pulumi:"claims"`
	Issuers   []string             `pulumi:"issuers"`
	Keys      []interface{}        `pulumi:"keys"`
	Type      string               `pulumi:"type"`
}

type JwtAuthenticationResponseOutput struct{ *pulumi.OutputState }

func (JwtAuthenticationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JwtAuthenticationResponse)(nil)).Elem()
}

func (o JwtAuthenticationResponseOutput) ToJwtAuthenticationResponseOutput() JwtAuthenticationResponseOutput {
	return o
}

func (o JwtAuthenticationResponseOutput) ToJwtAuthenticationResponseOutputWithContext(ctx context.Context) JwtAuthenticationResponseOutput {
	return o
}

func (o JwtAuthenticationResponseOutput) Audiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JwtAuthenticationResponse) []string { return v.Audiences }).(pulumi.StringArrayOutput)
}

func (o JwtAuthenticationResponseOutput) Claims() TokenClaimResponseArrayOutput {
	return o.ApplyT(func(v JwtAuthenticationResponse) []TokenClaimResponse { return v.Claims }).(TokenClaimResponseArrayOutput)
}

func (o JwtAuthenticationResponseOutput) Issuers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JwtAuthenticationResponse) []string { return v.Issuers }).(pulumi.StringArrayOutput)
}

func (o JwtAuthenticationResponseOutput) Keys() pulumi.ArrayOutput {
	return o.ApplyT(func(v JwtAuthenticationResponse) []interface{} { return v.Keys }).(pulumi.ArrayOutput)
}

func (o JwtAuthenticationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v JwtAuthenticationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type JwtAuthenticationResponsePtrOutput struct{ *pulumi.OutputState }

func (JwtAuthenticationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JwtAuthenticationResponse)(nil)).Elem()
}

func (o JwtAuthenticationResponsePtrOutput) ToJwtAuthenticationResponsePtrOutput() JwtAuthenticationResponsePtrOutput {
	return o
}

func (o JwtAuthenticationResponsePtrOutput) ToJwtAuthenticationResponsePtrOutputWithContext(ctx context.Context) JwtAuthenticationResponsePtrOutput {
	return o
}

func (o JwtAuthenticationResponsePtrOutput) Elem() JwtAuthenticationResponseOutput {
	return o.ApplyT(func(v *JwtAuthenticationResponse) JwtAuthenticationResponse {
		if v != nil {
			return *v
		}
		var ret JwtAuthenticationResponse
		return ret
	}).(JwtAuthenticationResponseOutput)
}

func (o JwtAuthenticationResponsePtrOutput) Audiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JwtAuthenticationResponse) []string {
		if v == nil {
			return nil
		}
		return v.Audiences
	}).(pulumi.StringArrayOutput)
}

func (o JwtAuthenticationResponsePtrOutput) Claims() TokenClaimResponseArrayOutput {
	return o.ApplyT(func(v *JwtAuthenticationResponse) []TokenClaimResponse {
		if v == nil {
			return nil
		}
		return v.Claims
	}).(TokenClaimResponseArrayOutput)
}

func (o JwtAuthenticationResponsePtrOutput) Issuers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JwtAuthenticationResponse) []string {
		if v == nil {
			return nil
		}
		return v.Issuers
	}).(pulumi.StringArrayOutput)
}

func (o JwtAuthenticationResponsePtrOutput) Keys() pulumi.ArrayOutput {
	return o.ApplyT(func(v *JwtAuthenticationResponse) []interface{} {
		if v == nil {
			return nil
		}
		return v.Keys
	}).(pulumi.ArrayOutput)
}

func (o JwtAuthenticationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JwtAuthenticationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type KeyVaultProperties struct {
	KeyIdentifier string `pulumi:"keyIdentifier"`
}





type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

type KeyVaultPropertiesArgs struct {
	KeyIdentifier pulumi.StringInput `pulumi:"keyIdentifier"`
}

func (KeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return i.ToKeyVaultPropertiesOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput)
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput).ToKeyVaultPropertiesPtrOutputWithContext(ctx)
}









type KeyVaultPropertiesPtrInput interface {
	pulumi.Input

	ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput
	ToKeyVaultPropertiesPtrOutputWithContext(context.Context) KeyVaultPropertiesPtrOutput
}

type keyVaultPropertiesPtrType KeyVaultPropertiesArgs

func KeyVaultPropertiesPtr(v *KeyVaultPropertiesArgs) KeyVaultPropertiesPtrInput {
	return (*keyVaultPropertiesPtrType)(v)
}

func (*keyVaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesPtrOutput)
}

type KeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultProperties) *KeyVaultProperties {
		return &v
	}).(KeyVaultPropertiesPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultProperties) string { return v.KeyIdentifier }).(pulumi.StringOutput)
}

type KeyVaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) Elem() KeyVaultPropertiesOutput {
	return o.ApplyT(func(v *KeyVaultProperties) KeyVaultProperties {
		if v != nil {
			return *v
		}
		var ret KeyVaultProperties
		return ret
	}).(KeyVaultPropertiesOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyIdentifier
	}).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponse struct {
	CurrentKeyIdentifier string `pulumi:"currentKeyIdentifier"`
	KeyIdentifier        string `pulumi:"keyIdentifier"`
}

type KeyVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) CurrentKeyIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) string { return v.CurrentKeyIdentifier }).(pulumi.StringOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) string { return v.KeyIdentifier }).(pulumi.StringOutput)
}

type KeyVaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) Elem() KeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) KeyVaultPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultPropertiesResponse
		return ret
	}).(KeyVaultPropertiesResponseOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) CurrentKeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CurrentKeyIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyIdentifier
	}).(pulumi.StringPtrOutput)
}

type NetworkAccessControl struct {
	Consumption *GroupLevelAccessControl `pulumi:"consumption"`
	Ingestion   *GroupLevelAccessControl `pulumi:"ingestion"`
	Integration *GroupLevelAccessControl `pulumi:"integration"`
}





type NetworkAccessControlInput interface {
	pulumi.Input

	ToNetworkAccessControlOutput() NetworkAccessControlOutput
	ToNetworkAccessControlOutputWithContext(context.Context) NetworkAccessControlOutput
}

type NetworkAccessControlArgs struct {
	Consumption GroupLevelAccessControlPtrInput `pulumi:"consumption"`
	Ingestion   GroupLevelAccessControlPtrInput `pulumi:"ingestion"`
	Integration GroupLevelAccessControlPtrInput `pulumi:"integration"`
}

func (NetworkAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAccessControl)(nil)).Elem()
}

func (i NetworkAccessControlArgs) ToNetworkAccessControlOutput() NetworkAccessControlOutput {
	return i.ToNetworkAccessControlOutputWithContext(context.Background())
}

func (i NetworkAccessControlArgs) ToNetworkAccessControlOutputWithContext(ctx context.Context) NetworkAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAccessControlOutput)
}

func (i NetworkAccessControlArgs) ToNetworkAccessControlPtrOutput() NetworkAccessControlPtrOutput {
	return i.ToNetworkAccessControlPtrOutputWithContext(context.Background())
}

func (i NetworkAccessControlArgs) ToNetworkAccessControlPtrOutputWithContext(ctx context.Context) NetworkAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAccessControlOutput).ToNetworkAccessControlPtrOutputWithContext(ctx)
}









type NetworkAccessControlPtrInput interface {
	pulumi.Input

	ToNetworkAccessControlPtrOutput() NetworkAccessControlPtrOutput
	ToNetworkAccessControlPtrOutputWithContext(context.Context) NetworkAccessControlPtrOutput
}

type networkAccessControlPtrType NetworkAccessControlArgs

func NetworkAccessControlPtr(v *NetworkAccessControlArgs) NetworkAccessControlPtrInput {
	return (*networkAccessControlPtrType)(v)
}

func (*networkAccessControlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkAccessControl)(nil)).Elem()
}

func (i *networkAccessControlPtrType) ToNetworkAccessControlPtrOutput() NetworkAccessControlPtrOutput {
	return i.ToNetworkAccessControlPtrOutputWithContext(context.Background())
}

func (i *networkAccessControlPtrType) ToNetworkAccessControlPtrOutputWithContext(ctx context.Context) NetworkAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAccessControlPtrOutput)
}

type NetworkAccessControlOutput struct{ *pulumi.OutputState }

func (NetworkAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAccessControl)(nil)).Elem()
}

func (o NetworkAccessControlOutput) ToNetworkAccessControlOutput() NetworkAccessControlOutput {
	return o
}

func (o NetworkAccessControlOutput) ToNetworkAccessControlOutputWithContext(ctx context.Context) NetworkAccessControlOutput {
	return o
}

func (o NetworkAccessControlOutput) ToNetworkAccessControlPtrOutput() NetworkAccessControlPtrOutput {
	return o.ToNetworkAccessControlPtrOutputWithContext(context.Background())
}

func (o NetworkAccessControlOutput) ToNetworkAccessControlPtrOutputWithContext(ctx context.Context) NetworkAccessControlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkAccessControl) *NetworkAccessControl {
		return &v
	}).(NetworkAccessControlPtrOutput)
}

func (o NetworkAccessControlOutput) Consumption() GroupLevelAccessControlPtrOutput {
	return o.ApplyT(func(v NetworkAccessControl) *GroupLevelAccessControl { return v.Consumption }).(GroupLevelAccessControlPtrOutput)
}

func (o NetworkAccessControlOutput) Ingestion() GroupLevelAccessControlPtrOutput {
	return o.ApplyT(func(v NetworkAccessControl) *GroupLevelAccessControl { return v.Ingestion }).(GroupLevelAccessControlPtrOutput)
}

func (o NetworkAccessControlOutput) Integration() GroupLevelAccessControlPtrOutput {
	return o.ApplyT(func(v NetworkAccessControl) *GroupLevelAccessControl { return v.Integration }).(GroupLevelAccessControlPtrOutput)
}

type NetworkAccessControlPtrOutput struct{ *pulumi.OutputState }

func (NetworkAccessControlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkAccessControl)(nil)).Elem()
}

func (o NetworkAccessControlPtrOutput) ToNetworkAccessControlPtrOutput() NetworkAccessControlPtrOutput {
	return o
}

func (o NetworkAccessControlPtrOutput) ToNetworkAccessControlPtrOutputWithContext(ctx context.Context) NetworkAccessControlPtrOutput {
	return o
}

func (o NetworkAccessControlPtrOutput) Elem() NetworkAccessControlOutput {
	return o.ApplyT(func(v *NetworkAccessControl) NetworkAccessControl {
		if v != nil {
			return *v
		}
		var ret NetworkAccessControl
		return ret
	}).(NetworkAccessControlOutput)
}

func (o NetworkAccessControlPtrOutput) Consumption() GroupLevelAccessControlPtrOutput {
	return o.ApplyT(func(v *NetworkAccessControl) *GroupLevelAccessControl {
		if v == nil {
			return nil
		}
		return v.Consumption
	}).(GroupLevelAccessControlPtrOutput)
}

func (o NetworkAccessControlPtrOutput) Ingestion() GroupLevelAccessControlPtrOutput {
	return o.ApplyT(func(v *NetworkAccessControl) *GroupLevelAccessControl {
		if v == nil {
			return nil
		}
		return v.Ingestion
	}).(GroupLevelAccessControlPtrOutput)
}

func (o NetworkAccessControlPtrOutput) Integration() GroupLevelAccessControlPtrOutput {
	return o.ApplyT(func(v *NetworkAccessControl) *GroupLevelAccessControl {
		if v == nil {
			return nil
		}
		return v.Integration
	}).(GroupLevelAccessControlPtrOutput)
}

type NetworkAccessControlResponse struct {
	Consumption *GroupLevelAccessControlResponse `pulumi:"consumption"`
	Ingestion   *GroupLevelAccessControlResponse `pulumi:"ingestion"`
	Integration *GroupLevelAccessControlResponse `pulumi:"integration"`
}

type NetworkAccessControlResponseOutput struct{ *pulumi.OutputState }

func (NetworkAccessControlResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAccessControlResponse)(nil)).Elem()
}

func (o NetworkAccessControlResponseOutput) ToNetworkAccessControlResponseOutput() NetworkAccessControlResponseOutput {
	return o
}

func (o NetworkAccessControlResponseOutput) ToNetworkAccessControlResponseOutputWithContext(ctx context.Context) NetworkAccessControlResponseOutput {
	return o
}

func (o NetworkAccessControlResponseOutput) Consumption() GroupLevelAccessControlResponsePtrOutput {
	return o.ApplyT(func(v NetworkAccessControlResponse) *GroupLevelAccessControlResponse { return v.Consumption }).(GroupLevelAccessControlResponsePtrOutput)
}

func (o NetworkAccessControlResponseOutput) Ingestion() GroupLevelAccessControlResponsePtrOutput {
	return o.ApplyT(func(v NetworkAccessControlResponse) *GroupLevelAccessControlResponse { return v.Ingestion }).(GroupLevelAccessControlResponsePtrOutput)
}

func (o NetworkAccessControlResponseOutput) Integration() GroupLevelAccessControlResponsePtrOutput {
	return o.ApplyT(func(v NetworkAccessControlResponse) *GroupLevelAccessControlResponse { return v.Integration }).(GroupLevelAccessControlResponsePtrOutput)
}

type NetworkAccessControlResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkAccessControlResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkAccessControlResponse)(nil)).Elem()
}

func (o NetworkAccessControlResponsePtrOutput) ToNetworkAccessControlResponsePtrOutput() NetworkAccessControlResponsePtrOutput {
	return o
}

func (o NetworkAccessControlResponsePtrOutput) ToNetworkAccessControlResponsePtrOutputWithContext(ctx context.Context) NetworkAccessControlResponsePtrOutput {
	return o
}

func (o NetworkAccessControlResponsePtrOutput) Elem() NetworkAccessControlResponseOutput {
	return o.ApplyT(func(v *NetworkAccessControlResponse) NetworkAccessControlResponse {
		if v != nil {
			return *v
		}
		var ret NetworkAccessControlResponse
		return ret
	}).(NetworkAccessControlResponseOutput)
}

func (o NetworkAccessControlResponsePtrOutput) Consumption() GroupLevelAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *NetworkAccessControlResponse) *GroupLevelAccessControlResponse {
		if v == nil {
			return nil
		}
		return v.Consumption
	}).(GroupLevelAccessControlResponsePtrOutput)
}

func (o NetworkAccessControlResponsePtrOutput) Ingestion() GroupLevelAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *NetworkAccessControlResponse) *GroupLevelAccessControlResponse {
		if v == nil {
			return nil
		}
		return v.Ingestion
	}).(GroupLevelAccessControlResponsePtrOutput)
}

func (o NetworkAccessControlResponsePtrOutput) Integration() GroupLevelAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *NetworkAccessControlResponse) *GroupLevelAccessControlResponse {
		if v == nil {
			return nil
		}
		return v.Integration
	}).(GroupLevelAccessControlResponsePtrOutput)
}

type NodeInput struct {
	NodeName string `pulumi:"nodeName"`
}





type NodeInputInput interface {
	pulumi.Input

	ToNodeInputOutput() NodeInputOutput
	ToNodeInputOutputWithContext(context.Context) NodeInputOutput
}

type NodeInputArgs struct {
	NodeName pulumi.StringInput `pulumi:"nodeName"`
}

func (NodeInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeInput)(nil)).Elem()
}

func (i NodeInputArgs) ToNodeInputOutput() NodeInputOutput {
	return i.ToNodeInputOutputWithContext(context.Background())
}

func (i NodeInputArgs) ToNodeInputOutputWithContext(ctx context.Context) NodeInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeInputOutput)
}





type NodeInputArrayInput interface {
	pulumi.Input

	ToNodeInputArrayOutput() NodeInputArrayOutput
	ToNodeInputArrayOutputWithContext(context.Context) NodeInputArrayOutput
}

type NodeInputArray []NodeInputInput

func (NodeInputArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeInput)(nil)).Elem()
}

func (i NodeInputArray) ToNodeInputArrayOutput() NodeInputArrayOutput {
	return i.ToNodeInputArrayOutputWithContext(context.Background())
}

func (i NodeInputArray) ToNodeInputArrayOutputWithContext(ctx context.Context) NodeInputArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeInputArrayOutput)
}

type NodeInputOutput struct{ *pulumi.OutputState }

func (NodeInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeInput)(nil)).Elem()
}

func (o NodeInputOutput) ToNodeInputOutput() NodeInputOutput {
	return o
}

func (o NodeInputOutput) ToNodeInputOutputWithContext(ctx context.Context) NodeInputOutput {
	return o
}

func (o NodeInputOutput) NodeName() pulumi.StringOutput {
	return o.ApplyT(func(v NodeInput) string { return v.NodeName }).(pulumi.StringOutput)
}

type NodeInputArrayOutput struct{ *pulumi.OutputState }

func (NodeInputArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeInput)(nil)).Elem()
}

func (o NodeInputArrayOutput) ToNodeInputArrayOutput() NodeInputArrayOutput {
	return o
}

func (o NodeInputArrayOutput) ToNodeInputArrayOutputWithContext(ctx context.Context) NodeInputArrayOutput {
	return o
}

func (o NodeInputArrayOutput) Index(i pulumi.IntInput) NodeInputOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NodeInput {
		return vs[0].([]NodeInput)[vs[1].(int)]
	}).(NodeInputOutput)
}

type NodeInputResponse struct {
	NodeName string `pulumi:"nodeName"`
}

type NodeInputResponseOutput struct{ *pulumi.OutputState }

func (NodeInputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeInputResponse)(nil)).Elem()
}

func (o NodeInputResponseOutput) ToNodeInputResponseOutput() NodeInputResponseOutput {
	return o
}

func (o NodeInputResponseOutput) ToNodeInputResponseOutputWithContext(ctx context.Context) NodeInputResponseOutput {
	return o
}

func (o NodeInputResponseOutput) NodeName() pulumi.StringOutput {
	return o.ApplyT(func(v NodeInputResponse) string { return v.NodeName }).(pulumi.StringOutput)
}

type NodeInputResponseArrayOutput struct{ *pulumi.OutputState }

func (NodeInputResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeInputResponse)(nil)).Elem()
}

func (o NodeInputResponseArrayOutput) ToNodeInputResponseArrayOutput() NodeInputResponseArrayOutput {
	return o
}

func (o NodeInputResponseArrayOutput) ToNodeInputResponseArrayOutputWithContext(ctx context.Context) NodeInputResponseArrayOutput {
	return o
}

func (o NodeInputResponseArrayOutput) Index(i pulumi.IntInput) NodeInputResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NodeInputResponse {
		return vs[0].([]NodeInputResponse)[vs[1].(int)]
	}).(NodeInputResponseOutput)
}

type ParameterDeclaration struct {
	Default     *string `pulumi:"default"`
	Description *string `pulumi:"description"`
	Name        string  `pulumi:"name"`
	Type        string  `pulumi:"type"`
}





type ParameterDeclarationInput interface {
	pulumi.Input

	ToParameterDeclarationOutput() ParameterDeclarationOutput
	ToParameterDeclarationOutputWithContext(context.Context) ParameterDeclarationOutput
}

type ParameterDeclarationArgs struct {
	Default     pulumi.StringPtrInput `pulumi:"default"`
	Description pulumi.StringPtrInput `pulumi:"description"`
	Name        pulumi.StringInput    `pulumi:"name"`
	Type        pulumi.StringInput    `pulumi:"type"`
}

func (ParameterDeclarationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDeclaration)(nil)).Elem()
}

func (i ParameterDeclarationArgs) ToParameterDeclarationOutput() ParameterDeclarationOutput {
	return i.ToParameterDeclarationOutputWithContext(context.Background())
}

func (i ParameterDeclarationArgs) ToParameterDeclarationOutputWithContext(ctx context.Context) ParameterDeclarationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDeclarationOutput)
}





type ParameterDeclarationArrayInput interface {
	pulumi.Input

	ToParameterDeclarationArrayOutput() ParameterDeclarationArrayOutput
	ToParameterDeclarationArrayOutputWithContext(context.Context) ParameterDeclarationArrayOutput
}

type ParameterDeclarationArray []ParameterDeclarationInput

func (ParameterDeclarationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterDeclaration)(nil)).Elem()
}

func (i ParameterDeclarationArray) ToParameterDeclarationArrayOutput() ParameterDeclarationArrayOutput {
	return i.ToParameterDeclarationArrayOutputWithContext(context.Background())
}

func (i ParameterDeclarationArray) ToParameterDeclarationArrayOutputWithContext(ctx context.Context) ParameterDeclarationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDeclarationArrayOutput)
}

type ParameterDeclarationOutput struct{ *pulumi.OutputState }

func (ParameterDeclarationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDeclaration)(nil)).Elem()
}

func (o ParameterDeclarationOutput) ToParameterDeclarationOutput() ParameterDeclarationOutput {
	return o
}

func (o ParameterDeclarationOutput) ToParameterDeclarationOutputWithContext(ctx context.Context) ParameterDeclarationOutput {
	return o
}

func (o ParameterDeclarationOutput) Default() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDeclaration) *string { return v.Default }).(pulumi.StringPtrOutput)
}

func (o ParameterDeclarationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDeclaration) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ParameterDeclarationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterDeclaration) string { return v.Name }).(pulumi.StringOutput)
}

func (o ParameterDeclarationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterDeclaration) string { return v.Type }).(pulumi.StringOutput)
}

type ParameterDeclarationArrayOutput struct{ *pulumi.OutputState }

func (ParameterDeclarationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterDeclaration)(nil)).Elem()
}

func (o ParameterDeclarationArrayOutput) ToParameterDeclarationArrayOutput() ParameterDeclarationArrayOutput {
	return o
}

func (o ParameterDeclarationArrayOutput) ToParameterDeclarationArrayOutputWithContext(ctx context.Context) ParameterDeclarationArrayOutput {
	return o
}

func (o ParameterDeclarationArrayOutput) Index(i pulumi.IntInput) ParameterDeclarationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ParameterDeclaration {
		return vs[0].([]ParameterDeclaration)[vs[1].(int)]
	}).(ParameterDeclarationOutput)
}

type ParameterDeclarationResponse struct {
	Default     *string `pulumi:"default"`
	Description *string `pulumi:"description"`
	Name        string  `pulumi:"name"`
	Type        string  `pulumi:"type"`
}

type ParameterDeclarationResponseOutput struct{ *pulumi.OutputState }

func (ParameterDeclarationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDeclarationResponse)(nil)).Elem()
}

func (o ParameterDeclarationResponseOutput) ToParameterDeclarationResponseOutput() ParameterDeclarationResponseOutput {
	return o
}

func (o ParameterDeclarationResponseOutput) ToParameterDeclarationResponseOutputWithContext(ctx context.Context) ParameterDeclarationResponseOutput {
	return o
}

func (o ParameterDeclarationResponseOutput) Default() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDeclarationResponse) *string { return v.Default }).(pulumi.StringPtrOutput)
}

func (o ParameterDeclarationResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDeclarationResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ParameterDeclarationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterDeclarationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ParameterDeclarationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterDeclarationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ParameterDeclarationResponseArrayOutput struct{ *pulumi.OutputState }

func (ParameterDeclarationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterDeclarationResponse)(nil)).Elem()
}

func (o ParameterDeclarationResponseArrayOutput) ToParameterDeclarationResponseArrayOutput() ParameterDeclarationResponseArrayOutput {
	return o
}

func (o ParameterDeclarationResponseArrayOutput) ToParameterDeclarationResponseArrayOutputWithContext(ctx context.Context) ParameterDeclarationResponseArrayOutput {
	return o
}

func (o ParameterDeclarationResponseArrayOutput) Index(i pulumi.IntInput) ParameterDeclarationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ParameterDeclarationResponse {
		return vs[0].([]ParameterDeclarationResponse)[vs[1].(int)]
	}).(ParameterDeclarationResponseOutput)
}

type ParameterDefinition struct {
	Name  string  `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type ParameterDefinitionInput interface {
	pulumi.Input

	ToParameterDefinitionOutput() ParameterDefinitionOutput
	ToParameterDefinitionOutputWithContext(context.Context) ParameterDefinitionOutput
}

type ParameterDefinitionArgs struct {
	Name  pulumi.StringInput    `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (ParameterDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinition)(nil)).Elem()
}

func (i ParameterDefinitionArgs) ToParameterDefinitionOutput() ParameterDefinitionOutput {
	return i.ToParameterDefinitionOutputWithContext(context.Background())
}

func (i ParameterDefinitionArgs) ToParameterDefinitionOutputWithContext(ctx context.Context) ParameterDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDefinitionOutput)
}





type ParameterDefinitionArrayInput interface {
	pulumi.Input

	ToParameterDefinitionArrayOutput() ParameterDefinitionArrayOutput
	ToParameterDefinitionArrayOutputWithContext(context.Context) ParameterDefinitionArrayOutput
}

type ParameterDefinitionArray []ParameterDefinitionInput

func (ParameterDefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterDefinition)(nil)).Elem()
}

func (i ParameterDefinitionArray) ToParameterDefinitionArrayOutput() ParameterDefinitionArrayOutput {
	return i.ToParameterDefinitionArrayOutputWithContext(context.Background())
}

func (i ParameterDefinitionArray) ToParameterDefinitionArrayOutputWithContext(ctx context.Context) ParameterDefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDefinitionArrayOutput)
}

type ParameterDefinitionOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinition)(nil)).Elem()
}

func (o ParameterDefinitionOutput) ToParameterDefinitionOutput() ParameterDefinitionOutput {
	return o
}

func (o ParameterDefinitionOutput) ToParameterDefinitionOutputWithContext(ctx context.Context) ParameterDefinitionOutput {
	return o
}

func (o ParameterDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterDefinition) string { return v.Name }).(pulumi.StringOutput)
}

func (o ParameterDefinitionOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinition) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ParameterDefinitionArrayOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterDefinition)(nil)).Elem()
}

func (o ParameterDefinitionArrayOutput) ToParameterDefinitionArrayOutput() ParameterDefinitionArrayOutput {
	return o
}

func (o ParameterDefinitionArrayOutput) ToParameterDefinitionArrayOutputWithContext(ctx context.Context) ParameterDefinitionArrayOutput {
	return o
}

func (o ParameterDefinitionArrayOutput) Index(i pulumi.IntInput) ParameterDefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ParameterDefinition {
		return vs[0].([]ParameterDefinition)[vs[1].(int)]
	}).(ParameterDefinitionOutput)
}

type ParameterDefinitionResponse struct {
	Name  string  `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type ParameterDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinitionResponse)(nil)).Elem()
}

func (o ParameterDefinitionResponseOutput) ToParameterDefinitionResponseOutput() ParameterDefinitionResponseOutput {
	return o
}

func (o ParameterDefinitionResponseOutput) ToParameterDefinitionResponseOutputWithContext(ctx context.Context) ParameterDefinitionResponseOutput {
	return o
}

func (o ParameterDefinitionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterDefinitionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ParameterDefinitionResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ParameterDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterDefinitionResponse)(nil)).Elem()
}

func (o ParameterDefinitionResponseArrayOutput) ToParameterDefinitionResponseArrayOutput() ParameterDefinitionResponseArrayOutput {
	return o
}

func (o ParameterDefinitionResponseArrayOutput) ToParameterDefinitionResponseArrayOutputWithContext(ctx context.Context) ParameterDefinitionResponseArrayOutput {
	return o
}

func (o ParameterDefinitionResponseArrayOutput) Index(i pulumi.IntInput) ParameterDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ParameterDefinitionResponse {
		return vs[0].([]ParameterDefinitionResponse)[vs[1].(int)]
	}).(ParameterDefinitionResponseOutput)
}

type PemCertificateList struct {
	Certificates []string `pulumi:"certificates"`
	Type         string   `pulumi:"type"`
}

type PemCertificateListResponse struct {
	Certificates []string `pulumi:"certificates"`
	Type         string   `pulumi:"type"`
}

type PipelineJobErrorResponse struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}

type PipelineJobErrorResponseOutput struct{ *pulumi.OutputState }

func (PipelineJobErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineJobErrorResponse)(nil)).Elem()
}

func (o PipelineJobErrorResponseOutput) ToPipelineJobErrorResponseOutput() PipelineJobErrorResponseOutput {
	return o
}

func (o PipelineJobErrorResponseOutput) ToPipelineJobErrorResponseOutputWithContext(ctx context.Context) PipelineJobErrorResponseOutput {
	return o
}

func (o PipelineJobErrorResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineJobErrorResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o PipelineJobErrorResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineJobErrorResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Type                              string                                    `pulumi:"type"`
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
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
		return &v.Id
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

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
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

type ResourceIdentity struct {
	UserAssignedIdentity string `pulumi:"userAssignedIdentity"`
}





type ResourceIdentityInput interface {
	pulumi.Input

	ToResourceIdentityOutput() ResourceIdentityOutput
	ToResourceIdentityOutputWithContext(context.Context) ResourceIdentityOutput
}

type ResourceIdentityArgs struct {
	UserAssignedIdentity pulumi.StringInput `pulumi:"userAssignedIdentity"`
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

func (o ResourceIdentityOutput) UserAssignedIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentity) string { return v.UserAssignedIdentity }).(pulumi.StringOutput)
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

func (o ResourceIdentityPtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityResponse struct {
	UserAssignedIdentity string `pulumi:"userAssignedIdentity"`
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

func (o ResourceIdentityResponseOutput) UserAssignedIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.UserAssignedIdentity }).(pulumi.StringOutput)
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

func (o ResourceIdentityResponsePtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type RsaTokenKey struct {
	Alg  string `pulumi:"alg"`
	E    string `pulumi:"e"`
	Kid  string `pulumi:"kid"`
	N    string `pulumi:"n"`
	Type string `pulumi:"type"`
}

type RsaTokenKeyResponse struct {
	Alg  string `pulumi:"alg"`
	E    string `pulumi:"e"`
	Kid  string `pulumi:"kid"`
	N    string `pulumi:"n"`
	Type string `pulumi:"type"`
}

type RtspSource struct {
	Endpoint  interface{} `pulumi:"endpoint"`
	Name      string      `pulumi:"name"`
	Transport *string     `pulumi:"transport"`
	Type      string      `pulumi:"type"`
}

type RtspSourceResponse struct {
	Endpoint  interface{} `pulumi:"endpoint"`
	Name      string      `pulumi:"name"`
	Transport *string     `pulumi:"transport"`
	Type      string      `pulumi:"type"`
}

type SecureIotDeviceRemoteTunnel struct {
	DeviceId   string `pulumi:"deviceId"`
	IotHubName string `pulumi:"iotHubName"`
	Type       string `pulumi:"type"`
}

type SecureIotDeviceRemoteTunnelResponse struct {
	DeviceId   string `pulumi:"deviceId"`
	IotHubName string `pulumi:"iotHubName"`
	Type       string `pulumi:"type"`
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

type StorageAccount struct {
	Id       string            `pulumi:"id"`
	Identity *ResourceIdentity `pulumi:"identity"`
}





type StorageAccountInput interface {
	pulumi.Input

	ToStorageAccountOutput() StorageAccountOutput
	ToStorageAccountOutputWithContext(context.Context) StorageAccountOutput
}

type StorageAccountArgs struct {
	Id       pulumi.StringInput       `pulumi:"id"`
	Identity ResourceIdentityPtrInput `pulumi:"identity"`
}

func (StorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (i StorageAccountArgs) ToStorageAccountOutput() StorageAccountOutput {
	return i.ToStorageAccountOutputWithContext(context.Background())
}

func (i StorageAccountArgs) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput)
}





type StorageAccountArrayInput interface {
	pulumi.Input

	ToStorageAccountArrayOutput() StorageAccountArrayOutput
	ToStorageAccountArrayOutputWithContext(context.Context) StorageAccountArrayOutput
}

type StorageAccountArray []StorageAccountInput

func (StorageAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccount)(nil)).Elem()
}

func (i StorageAccountArray) ToStorageAccountArrayOutput() StorageAccountArrayOutput {
	return i.ToStorageAccountArrayOutputWithContext(context.Background())
}

func (i StorageAccountArray) ToStorageAccountArrayOutputWithContext(ctx context.Context) StorageAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountArrayOutput)
}

type StorageAccountOutput struct{ *pulumi.OutputState }

func (StorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (o StorageAccountOutput) ToStorageAccountOutput() StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccount) string { return v.Id }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) Identity() ResourceIdentityPtrOutput {
	return o.ApplyT(func(v StorageAccount) *ResourceIdentity { return v.Identity }).(ResourceIdentityPtrOutput)
}

type StorageAccountArrayOutput struct{ *pulumi.OutputState }

func (StorageAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccount)(nil)).Elem()
}

func (o StorageAccountArrayOutput) ToStorageAccountArrayOutput() StorageAccountArrayOutput {
	return o
}

func (o StorageAccountArrayOutput) ToStorageAccountArrayOutputWithContext(ctx context.Context) StorageAccountArrayOutput {
	return o
}

func (o StorageAccountArrayOutput) Index(i pulumi.IntInput) StorageAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageAccount {
		return vs[0].([]StorageAccount)[vs[1].(int)]
	}).(StorageAccountOutput)
}

type StorageAccountResponse struct {
	Id       string                    `pulumi:"id"`
	Identity *ResourceIdentityResponse `pulumi:"identity"`
	Status   string                    `pulumi:"status"`
}

type StorageAccountResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutput() StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutputWithContext(ctx context.Context) StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o StorageAccountResponseOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o StorageAccountResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountResponse) string { return v.Status }).(pulumi.StringOutput)
}

type StorageAccountResponseArrayOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseArrayOutput) ToStorageAccountResponseArrayOutput() StorageAccountResponseArrayOutput {
	return o
}

func (o StorageAccountResponseArrayOutput) ToStorageAccountResponseArrayOutputWithContext(ctx context.Context) StorageAccountResponseArrayOutput {
	return o
}

func (o StorageAccountResponseArrayOutput) Index(i pulumi.IntInput) StorageAccountResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageAccountResponse {
		return vs[0].([]StorageAccountResponse)[vs[1].(int)]
	}).(StorageAccountResponseOutput)
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

type TlsEndpoint struct {
	Credentials         UsernamePasswordCredentials  `pulumi:"credentials"`
	TrustedCertificates *PemCertificateList          `pulumi:"trustedCertificates"`
	Tunnel              *SecureIotDeviceRemoteTunnel `pulumi:"tunnel"`
	Type                string                       `pulumi:"type"`
	Url                 string                       `pulumi:"url"`
	ValidationOptions   *TlsValidationOptions        `pulumi:"validationOptions"`
}

type TlsEndpointResponse struct {
	Credentials         UsernamePasswordCredentialsResponse  `pulumi:"credentials"`
	TrustedCertificates *PemCertificateListResponse          `pulumi:"trustedCertificates"`
	Tunnel              *SecureIotDeviceRemoteTunnelResponse `pulumi:"tunnel"`
	Type                string                               `pulumi:"type"`
	Url                 string                               `pulumi:"url"`
	ValidationOptions   *TlsValidationOptionsResponse        `pulumi:"validationOptions"`
}

type TlsValidationOptions struct {
	IgnoreHostname  *string `pulumi:"ignoreHostname"`
	IgnoreSignature *string `pulumi:"ignoreSignature"`
}

type TlsValidationOptionsResponse struct {
	IgnoreHostname  *string `pulumi:"ignoreHostname"`
	IgnoreSignature *string `pulumi:"ignoreSignature"`
}

type TokenClaim struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type TokenClaimInput interface {
	pulumi.Input

	ToTokenClaimOutput() TokenClaimOutput
	ToTokenClaimOutputWithContext(context.Context) TokenClaimOutput
}

type TokenClaimArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (TokenClaimArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenClaim)(nil)).Elem()
}

func (i TokenClaimArgs) ToTokenClaimOutput() TokenClaimOutput {
	return i.ToTokenClaimOutputWithContext(context.Background())
}

func (i TokenClaimArgs) ToTokenClaimOutputWithContext(ctx context.Context) TokenClaimOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenClaimOutput)
}





type TokenClaimArrayInput interface {
	pulumi.Input

	ToTokenClaimArrayOutput() TokenClaimArrayOutput
	ToTokenClaimArrayOutputWithContext(context.Context) TokenClaimArrayOutput
}

type TokenClaimArray []TokenClaimInput

func (TokenClaimArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenClaim)(nil)).Elem()
}

func (i TokenClaimArray) ToTokenClaimArrayOutput() TokenClaimArrayOutput {
	return i.ToTokenClaimArrayOutputWithContext(context.Background())
}

func (i TokenClaimArray) ToTokenClaimArrayOutputWithContext(ctx context.Context) TokenClaimArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenClaimArrayOutput)
}

type TokenClaimOutput struct{ *pulumi.OutputState }

func (TokenClaimOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenClaim)(nil)).Elem()
}

func (o TokenClaimOutput) ToTokenClaimOutput() TokenClaimOutput {
	return o
}

func (o TokenClaimOutput) ToTokenClaimOutputWithContext(ctx context.Context) TokenClaimOutput {
	return o
}

func (o TokenClaimOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TokenClaim) string { return v.Name }).(pulumi.StringOutput)
}

func (o TokenClaimOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TokenClaim) string { return v.Value }).(pulumi.StringOutput)
}

type TokenClaimArrayOutput struct{ *pulumi.OutputState }

func (TokenClaimArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenClaim)(nil)).Elem()
}

func (o TokenClaimArrayOutput) ToTokenClaimArrayOutput() TokenClaimArrayOutput {
	return o
}

func (o TokenClaimArrayOutput) ToTokenClaimArrayOutputWithContext(ctx context.Context) TokenClaimArrayOutput {
	return o
}

func (o TokenClaimArrayOutput) Index(i pulumi.IntInput) TokenClaimOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenClaim {
		return vs[0].([]TokenClaim)[vs[1].(int)]
	}).(TokenClaimOutput)
}

type TokenClaimResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type TokenClaimResponseOutput struct{ *pulumi.OutputState }

func (TokenClaimResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenClaimResponse)(nil)).Elem()
}

func (o TokenClaimResponseOutput) ToTokenClaimResponseOutput() TokenClaimResponseOutput {
	return o
}

func (o TokenClaimResponseOutput) ToTokenClaimResponseOutputWithContext(ctx context.Context) TokenClaimResponseOutput {
	return o
}

func (o TokenClaimResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TokenClaimResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o TokenClaimResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TokenClaimResponse) string { return v.Value }).(pulumi.StringOutput)
}

type TokenClaimResponseArrayOutput struct{ *pulumi.OutputState }

func (TokenClaimResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenClaimResponse)(nil)).Elem()
}

func (o TokenClaimResponseArrayOutput) ToTokenClaimResponseArrayOutput() TokenClaimResponseArrayOutput {
	return o
}

func (o TokenClaimResponseArrayOutput) ToTokenClaimResponseArrayOutputWithContext(ctx context.Context) TokenClaimResponseArrayOutput {
	return o
}

func (o TokenClaimResponseArrayOutput) Index(i pulumi.IntInput) TokenClaimResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenClaimResponse {
		return vs[0].([]TokenClaimResponse)[vs[1].(int)]
	}).(TokenClaimResponseOutput)
}

type UnsecuredEndpoint struct {
	Credentials UsernamePasswordCredentials  `pulumi:"credentials"`
	Tunnel      *SecureIotDeviceRemoteTunnel `pulumi:"tunnel"`
	Type        string                       `pulumi:"type"`
	Url         string                       `pulumi:"url"`
}

type UnsecuredEndpointResponse struct {
	Credentials UsernamePasswordCredentialsResponse  `pulumi:"credentials"`
	Tunnel      *SecureIotDeviceRemoteTunnelResponse `pulumi:"tunnel"`
	Type        string                               `pulumi:"type"`
	Url         string                               `pulumi:"url"`
}

type UserAssignedManagedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type UserAssignedManagedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedManagedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedManagedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedManagedIdentityResponseOutput) ToUserAssignedManagedIdentityResponseOutput() UserAssignedManagedIdentityResponseOutput {
	return o
}

func (o UserAssignedManagedIdentityResponseOutput) ToUserAssignedManagedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedManagedIdentityResponseOutput {
	return o
}

func (o UserAssignedManagedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedManagedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedManagedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedManagedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedManagedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedManagedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedManagedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedManagedIdentityResponseMapOutput) ToUserAssignedManagedIdentityResponseMapOutput() UserAssignedManagedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedManagedIdentityResponseMapOutput) ToUserAssignedManagedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedManagedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedManagedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedManagedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedManagedIdentityResponse {
		return vs[0].(map[string]UserAssignedManagedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedManagedIdentityResponseOutput)
}

type UsernamePasswordCredentials struct {
	Password string `pulumi:"password"`
	Type     string `pulumi:"type"`
	Username string `pulumi:"username"`
}

type UsernamePasswordCredentialsResponse struct {
	Password string `pulumi:"password"`
	Type     string `pulumi:"type"`
	Username string `pulumi:"username"`
}

type VideoAnalyzerIdentity struct {
	Type                   string                 `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type VideoAnalyzerIdentityInput interface {
	pulumi.Input

	ToVideoAnalyzerIdentityOutput() VideoAnalyzerIdentityOutput
	ToVideoAnalyzerIdentityOutputWithContext(context.Context) VideoAnalyzerIdentityOutput
}

type VideoAnalyzerIdentityArgs struct {
	Type                   pulumi.StringInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput    `pulumi:"userAssignedIdentities"`
}

func (VideoAnalyzerIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoAnalyzerIdentity)(nil)).Elem()
}

func (i VideoAnalyzerIdentityArgs) ToVideoAnalyzerIdentityOutput() VideoAnalyzerIdentityOutput {
	return i.ToVideoAnalyzerIdentityOutputWithContext(context.Background())
}

func (i VideoAnalyzerIdentityArgs) ToVideoAnalyzerIdentityOutputWithContext(ctx context.Context) VideoAnalyzerIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoAnalyzerIdentityOutput)
}

func (i VideoAnalyzerIdentityArgs) ToVideoAnalyzerIdentityPtrOutput() VideoAnalyzerIdentityPtrOutput {
	return i.ToVideoAnalyzerIdentityPtrOutputWithContext(context.Background())
}

func (i VideoAnalyzerIdentityArgs) ToVideoAnalyzerIdentityPtrOutputWithContext(ctx context.Context) VideoAnalyzerIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoAnalyzerIdentityOutput).ToVideoAnalyzerIdentityPtrOutputWithContext(ctx)
}









type VideoAnalyzerIdentityPtrInput interface {
	pulumi.Input

	ToVideoAnalyzerIdentityPtrOutput() VideoAnalyzerIdentityPtrOutput
	ToVideoAnalyzerIdentityPtrOutputWithContext(context.Context) VideoAnalyzerIdentityPtrOutput
}

type videoAnalyzerIdentityPtrType VideoAnalyzerIdentityArgs

func VideoAnalyzerIdentityPtr(v *VideoAnalyzerIdentityArgs) VideoAnalyzerIdentityPtrInput {
	return (*videoAnalyzerIdentityPtrType)(v)
}

func (*videoAnalyzerIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoAnalyzerIdentity)(nil)).Elem()
}

func (i *videoAnalyzerIdentityPtrType) ToVideoAnalyzerIdentityPtrOutput() VideoAnalyzerIdentityPtrOutput {
	return i.ToVideoAnalyzerIdentityPtrOutputWithContext(context.Background())
}

func (i *videoAnalyzerIdentityPtrType) ToVideoAnalyzerIdentityPtrOutputWithContext(ctx context.Context) VideoAnalyzerIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoAnalyzerIdentityPtrOutput)
}

type VideoAnalyzerIdentityOutput struct{ *pulumi.OutputState }

func (VideoAnalyzerIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoAnalyzerIdentity)(nil)).Elem()
}

func (o VideoAnalyzerIdentityOutput) ToVideoAnalyzerIdentityOutput() VideoAnalyzerIdentityOutput {
	return o
}

func (o VideoAnalyzerIdentityOutput) ToVideoAnalyzerIdentityOutputWithContext(ctx context.Context) VideoAnalyzerIdentityOutput {
	return o
}

func (o VideoAnalyzerIdentityOutput) ToVideoAnalyzerIdentityPtrOutput() VideoAnalyzerIdentityPtrOutput {
	return o.ToVideoAnalyzerIdentityPtrOutputWithContext(context.Background())
}

func (o VideoAnalyzerIdentityOutput) ToVideoAnalyzerIdentityPtrOutputWithContext(ctx context.Context) VideoAnalyzerIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VideoAnalyzerIdentity) *VideoAnalyzerIdentity {
		return &v
	}).(VideoAnalyzerIdentityPtrOutput)
}

func (o VideoAnalyzerIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VideoAnalyzerIdentity) string { return v.Type }).(pulumi.StringOutput)
}

func (o VideoAnalyzerIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v VideoAnalyzerIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type VideoAnalyzerIdentityPtrOutput struct{ *pulumi.OutputState }

func (VideoAnalyzerIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoAnalyzerIdentity)(nil)).Elem()
}

func (o VideoAnalyzerIdentityPtrOutput) ToVideoAnalyzerIdentityPtrOutput() VideoAnalyzerIdentityPtrOutput {
	return o
}

func (o VideoAnalyzerIdentityPtrOutput) ToVideoAnalyzerIdentityPtrOutputWithContext(ctx context.Context) VideoAnalyzerIdentityPtrOutput {
	return o
}

func (o VideoAnalyzerIdentityPtrOutput) Elem() VideoAnalyzerIdentityOutput {
	return o.ApplyT(func(v *VideoAnalyzerIdentity) VideoAnalyzerIdentity {
		if v != nil {
			return *v
		}
		var ret VideoAnalyzerIdentity
		return ret
	}).(VideoAnalyzerIdentityOutput)
}

func (o VideoAnalyzerIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoAnalyzerIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o VideoAnalyzerIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *VideoAnalyzerIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type VideoAnalyzerIdentityResponse struct {
	Type                   string                                         `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedManagedIdentityResponse `pulumi:"userAssignedIdentities"`
}

type VideoAnalyzerIdentityResponseOutput struct{ *pulumi.OutputState }

func (VideoAnalyzerIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoAnalyzerIdentityResponse)(nil)).Elem()
}

func (o VideoAnalyzerIdentityResponseOutput) ToVideoAnalyzerIdentityResponseOutput() VideoAnalyzerIdentityResponseOutput {
	return o
}

func (o VideoAnalyzerIdentityResponseOutput) ToVideoAnalyzerIdentityResponseOutputWithContext(ctx context.Context) VideoAnalyzerIdentityResponseOutput {
	return o
}

func (o VideoAnalyzerIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VideoAnalyzerIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VideoAnalyzerIdentityResponseOutput) UserAssignedIdentities() UserAssignedManagedIdentityResponseMapOutput {
	return o.ApplyT(func(v VideoAnalyzerIdentityResponse) map[string]UserAssignedManagedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedManagedIdentityResponseMapOutput)
}

type VideoAnalyzerIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (VideoAnalyzerIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoAnalyzerIdentityResponse)(nil)).Elem()
}

func (o VideoAnalyzerIdentityResponsePtrOutput) ToVideoAnalyzerIdentityResponsePtrOutput() VideoAnalyzerIdentityResponsePtrOutput {
	return o
}

func (o VideoAnalyzerIdentityResponsePtrOutput) ToVideoAnalyzerIdentityResponsePtrOutputWithContext(ctx context.Context) VideoAnalyzerIdentityResponsePtrOutput {
	return o
}

func (o VideoAnalyzerIdentityResponsePtrOutput) Elem() VideoAnalyzerIdentityResponseOutput {
	return o.ApplyT(func(v *VideoAnalyzerIdentityResponse) VideoAnalyzerIdentityResponse {
		if v != nil {
			return *v
		}
		var ret VideoAnalyzerIdentityResponse
		return ret
	}).(VideoAnalyzerIdentityResponseOutput)
}

func (o VideoAnalyzerIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoAnalyzerIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o VideoAnalyzerIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedManagedIdentityResponseMapOutput {
	return o.ApplyT(func(v *VideoAnalyzerIdentityResponse) map[string]UserAssignedManagedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedManagedIdentityResponseMapOutput)
}

type VideoArchival struct {
	RetentionPeriod *string `pulumi:"retentionPeriod"`
}





type VideoArchivalInput interface {
	pulumi.Input

	ToVideoArchivalOutput() VideoArchivalOutput
	ToVideoArchivalOutputWithContext(context.Context) VideoArchivalOutput
}

type VideoArchivalArgs struct {
	RetentionPeriod pulumi.StringPtrInput `pulumi:"retentionPeriod"`
}

func (VideoArchivalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoArchival)(nil)).Elem()
}

func (i VideoArchivalArgs) ToVideoArchivalOutput() VideoArchivalOutput {
	return i.ToVideoArchivalOutputWithContext(context.Background())
}

func (i VideoArchivalArgs) ToVideoArchivalOutputWithContext(ctx context.Context) VideoArchivalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoArchivalOutput)
}

func (i VideoArchivalArgs) ToVideoArchivalPtrOutput() VideoArchivalPtrOutput {
	return i.ToVideoArchivalPtrOutputWithContext(context.Background())
}

func (i VideoArchivalArgs) ToVideoArchivalPtrOutputWithContext(ctx context.Context) VideoArchivalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoArchivalOutput).ToVideoArchivalPtrOutputWithContext(ctx)
}









type VideoArchivalPtrInput interface {
	pulumi.Input

	ToVideoArchivalPtrOutput() VideoArchivalPtrOutput
	ToVideoArchivalPtrOutputWithContext(context.Context) VideoArchivalPtrOutput
}

type videoArchivalPtrType VideoArchivalArgs

func VideoArchivalPtr(v *VideoArchivalArgs) VideoArchivalPtrInput {
	return (*videoArchivalPtrType)(v)
}

func (*videoArchivalPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoArchival)(nil)).Elem()
}

func (i *videoArchivalPtrType) ToVideoArchivalPtrOutput() VideoArchivalPtrOutput {
	return i.ToVideoArchivalPtrOutputWithContext(context.Background())
}

func (i *videoArchivalPtrType) ToVideoArchivalPtrOutputWithContext(ctx context.Context) VideoArchivalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoArchivalPtrOutput)
}

type VideoArchivalOutput struct{ *pulumi.OutputState }

func (VideoArchivalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoArchival)(nil)).Elem()
}

func (o VideoArchivalOutput) ToVideoArchivalOutput() VideoArchivalOutput {
	return o
}

func (o VideoArchivalOutput) ToVideoArchivalOutputWithContext(ctx context.Context) VideoArchivalOutput {
	return o
}

func (o VideoArchivalOutput) ToVideoArchivalPtrOutput() VideoArchivalPtrOutput {
	return o.ToVideoArchivalPtrOutputWithContext(context.Background())
}

func (o VideoArchivalOutput) ToVideoArchivalPtrOutputWithContext(ctx context.Context) VideoArchivalPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VideoArchival) *VideoArchival {
		return &v
	}).(VideoArchivalPtrOutput)
}

func (o VideoArchivalOutput) RetentionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoArchival) *string { return v.RetentionPeriod }).(pulumi.StringPtrOutput)
}

type VideoArchivalPtrOutput struct{ *pulumi.OutputState }

func (VideoArchivalPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoArchival)(nil)).Elem()
}

func (o VideoArchivalPtrOutput) ToVideoArchivalPtrOutput() VideoArchivalPtrOutput {
	return o
}

func (o VideoArchivalPtrOutput) ToVideoArchivalPtrOutputWithContext(ctx context.Context) VideoArchivalPtrOutput {
	return o
}

func (o VideoArchivalPtrOutput) Elem() VideoArchivalOutput {
	return o.ApplyT(func(v *VideoArchival) VideoArchival {
		if v != nil {
			return *v
		}
		var ret VideoArchival
		return ret
	}).(VideoArchivalOutput)
}

func (o VideoArchivalPtrOutput) RetentionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoArchival) *string {
		if v == nil {
			return nil
		}
		return v.RetentionPeriod
	}).(pulumi.StringPtrOutput)
}

type VideoArchivalResponse struct {
	RetentionPeriod *string `pulumi:"retentionPeriod"`
}

type VideoArchivalResponseOutput struct{ *pulumi.OutputState }

func (VideoArchivalResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoArchivalResponse)(nil)).Elem()
}

func (o VideoArchivalResponseOutput) ToVideoArchivalResponseOutput() VideoArchivalResponseOutput {
	return o
}

func (o VideoArchivalResponseOutput) ToVideoArchivalResponseOutputWithContext(ctx context.Context) VideoArchivalResponseOutput {
	return o
}

func (o VideoArchivalResponseOutput) RetentionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoArchivalResponse) *string { return v.RetentionPeriod }).(pulumi.StringPtrOutput)
}

type VideoArchivalResponsePtrOutput struct{ *pulumi.OutputState }

func (VideoArchivalResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoArchivalResponse)(nil)).Elem()
}

func (o VideoArchivalResponsePtrOutput) ToVideoArchivalResponsePtrOutput() VideoArchivalResponsePtrOutput {
	return o
}

func (o VideoArchivalResponsePtrOutput) ToVideoArchivalResponsePtrOutputWithContext(ctx context.Context) VideoArchivalResponsePtrOutput {
	return o
}

func (o VideoArchivalResponsePtrOutput) Elem() VideoArchivalResponseOutput {
	return o.ApplyT(func(v *VideoArchivalResponse) VideoArchivalResponse {
		if v != nil {
			return *v
		}
		var ret VideoArchivalResponse
		return ret
	}).(VideoArchivalResponseOutput)
}

func (o VideoArchivalResponsePtrOutput) RetentionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoArchivalResponse) *string {
		if v == nil {
			return nil
		}
		return v.RetentionPeriod
	}).(pulumi.StringPtrOutput)
}

type VideoContentUrlsResponse struct {
	ArchiveBaseUrl   *string                        `pulumi:"archiveBaseUrl"`
	DownloadUrl      *string                        `pulumi:"downloadUrl"`
	PreviewImageUrls *VideoPreviewImageUrlsResponse `pulumi:"previewImageUrls"`
	RtspTunnelUrl    *string                        `pulumi:"rtspTunnelUrl"`
}

type VideoContentUrlsResponseOutput struct{ *pulumi.OutputState }

func (VideoContentUrlsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoContentUrlsResponse)(nil)).Elem()
}

func (o VideoContentUrlsResponseOutput) ToVideoContentUrlsResponseOutput() VideoContentUrlsResponseOutput {
	return o
}

func (o VideoContentUrlsResponseOutput) ToVideoContentUrlsResponseOutputWithContext(ctx context.Context) VideoContentUrlsResponseOutput {
	return o
}

func (o VideoContentUrlsResponseOutput) ArchiveBaseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoContentUrlsResponse) *string { return v.ArchiveBaseUrl }).(pulumi.StringPtrOutput)
}

func (o VideoContentUrlsResponseOutput) DownloadUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoContentUrlsResponse) *string { return v.DownloadUrl }).(pulumi.StringPtrOutput)
}

func (o VideoContentUrlsResponseOutput) PreviewImageUrls() VideoPreviewImageUrlsResponsePtrOutput {
	return o.ApplyT(func(v VideoContentUrlsResponse) *VideoPreviewImageUrlsResponse { return v.PreviewImageUrls }).(VideoPreviewImageUrlsResponsePtrOutput)
}

func (o VideoContentUrlsResponseOutput) RtspTunnelUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoContentUrlsResponse) *string { return v.RtspTunnelUrl }).(pulumi.StringPtrOutput)
}

type VideoCreationProperties struct {
	Description     *string `pulumi:"description"`
	RetentionPeriod *string `pulumi:"retentionPeriod"`
	SegmentLength   *string `pulumi:"segmentLength"`
	Title           *string `pulumi:"title"`
}





type VideoCreationPropertiesInput interface {
	pulumi.Input

	ToVideoCreationPropertiesOutput() VideoCreationPropertiesOutput
	ToVideoCreationPropertiesOutputWithContext(context.Context) VideoCreationPropertiesOutput
}

type VideoCreationPropertiesArgs struct {
	Description     pulumi.StringPtrInput `pulumi:"description"`
	RetentionPeriod pulumi.StringPtrInput `pulumi:"retentionPeriod"`
	SegmentLength   pulumi.StringPtrInput `pulumi:"segmentLength"`
	Title           pulumi.StringPtrInput `pulumi:"title"`
}

func (VideoCreationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoCreationProperties)(nil)).Elem()
}

func (i VideoCreationPropertiesArgs) ToVideoCreationPropertiesOutput() VideoCreationPropertiesOutput {
	return i.ToVideoCreationPropertiesOutputWithContext(context.Background())
}

func (i VideoCreationPropertiesArgs) ToVideoCreationPropertiesOutputWithContext(ctx context.Context) VideoCreationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoCreationPropertiesOutput)
}

func (i VideoCreationPropertiesArgs) ToVideoCreationPropertiesPtrOutput() VideoCreationPropertiesPtrOutput {
	return i.ToVideoCreationPropertiesPtrOutputWithContext(context.Background())
}

func (i VideoCreationPropertiesArgs) ToVideoCreationPropertiesPtrOutputWithContext(ctx context.Context) VideoCreationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoCreationPropertiesOutput).ToVideoCreationPropertiesPtrOutputWithContext(ctx)
}









type VideoCreationPropertiesPtrInput interface {
	pulumi.Input

	ToVideoCreationPropertiesPtrOutput() VideoCreationPropertiesPtrOutput
	ToVideoCreationPropertiesPtrOutputWithContext(context.Context) VideoCreationPropertiesPtrOutput
}

type videoCreationPropertiesPtrType VideoCreationPropertiesArgs

func VideoCreationPropertiesPtr(v *VideoCreationPropertiesArgs) VideoCreationPropertiesPtrInput {
	return (*videoCreationPropertiesPtrType)(v)
}

func (*videoCreationPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoCreationProperties)(nil)).Elem()
}

func (i *videoCreationPropertiesPtrType) ToVideoCreationPropertiesPtrOutput() VideoCreationPropertiesPtrOutput {
	return i.ToVideoCreationPropertiesPtrOutputWithContext(context.Background())
}

func (i *videoCreationPropertiesPtrType) ToVideoCreationPropertiesPtrOutputWithContext(ctx context.Context) VideoCreationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoCreationPropertiesPtrOutput)
}

type VideoCreationPropertiesOutput struct{ *pulumi.OutputState }

func (VideoCreationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoCreationProperties)(nil)).Elem()
}

func (o VideoCreationPropertiesOutput) ToVideoCreationPropertiesOutput() VideoCreationPropertiesOutput {
	return o
}

func (o VideoCreationPropertiesOutput) ToVideoCreationPropertiesOutputWithContext(ctx context.Context) VideoCreationPropertiesOutput {
	return o
}

func (o VideoCreationPropertiesOutput) ToVideoCreationPropertiesPtrOutput() VideoCreationPropertiesPtrOutput {
	return o.ToVideoCreationPropertiesPtrOutputWithContext(context.Background())
}

func (o VideoCreationPropertiesOutput) ToVideoCreationPropertiesPtrOutputWithContext(ctx context.Context) VideoCreationPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VideoCreationProperties) *VideoCreationProperties {
		return &v
	}).(VideoCreationPropertiesPtrOutput)
}

func (o VideoCreationPropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoCreationProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesOutput) RetentionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoCreationProperties) *string { return v.RetentionPeriod }).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesOutput) SegmentLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoCreationProperties) *string { return v.SegmentLength }).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoCreationProperties) *string { return v.Title }).(pulumi.StringPtrOutput)
}

type VideoCreationPropertiesPtrOutput struct{ *pulumi.OutputState }

func (VideoCreationPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoCreationProperties)(nil)).Elem()
}

func (o VideoCreationPropertiesPtrOutput) ToVideoCreationPropertiesPtrOutput() VideoCreationPropertiesPtrOutput {
	return o
}

func (o VideoCreationPropertiesPtrOutput) ToVideoCreationPropertiesPtrOutputWithContext(ctx context.Context) VideoCreationPropertiesPtrOutput {
	return o
}

func (o VideoCreationPropertiesPtrOutput) Elem() VideoCreationPropertiesOutput {
	return o.ApplyT(func(v *VideoCreationProperties) VideoCreationProperties {
		if v != nil {
			return *v
		}
		var ret VideoCreationProperties
		return ret
	}).(VideoCreationPropertiesOutput)
}

func (o VideoCreationPropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesPtrOutput) RetentionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.RetentionPeriod
	}).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesPtrOutput) SegmentLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.SegmentLength
	}).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesPtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.Title
	}).(pulumi.StringPtrOutput)
}

type VideoCreationPropertiesResponse struct {
	Description     *string `pulumi:"description"`
	RetentionPeriod *string `pulumi:"retentionPeriod"`
	SegmentLength   *string `pulumi:"segmentLength"`
	Title           *string `pulumi:"title"`
}

type VideoCreationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VideoCreationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoCreationPropertiesResponse)(nil)).Elem()
}

func (o VideoCreationPropertiesResponseOutput) ToVideoCreationPropertiesResponseOutput() VideoCreationPropertiesResponseOutput {
	return o
}

func (o VideoCreationPropertiesResponseOutput) ToVideoCreationPropertiesResponseOutputWithContext(ctx context.Context) VideoCreationPropertiesResponseOutput {
	return o
}

func (o VideoCreationPropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoCreationPropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesResponseOutput) RetentionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoCreationPropertiesResponse) *string { return v.RetentionPeriod }).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesResponseOutput) SegmentLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoCreationPropertiesResponse) *string { return v.SegmentLength }).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesResponseOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoCreationPropertiesResponse) *string { return v.Title }).(pulumi.StringPtrOutput)
}

type VideoCreationPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (VideoCreationPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoCreationPropertiesResponse)(nil)).Elem()
}

func (o VideoCreationPropertiesResponsePtrOutput) ToVideoCreationPropertiesResponsePtrOutput() VideoCreationPropertiesResponsePtrOutput {
	return o
}

func (o VideoCreationPropertiesResponsePtrOutput) ToVideoCreationPropertiesResponsePtrOutputWithContext(ctx context.Context) VideoCreationPropertiesResponsePtrOutput {
	return o
}

func (o VideoCreationPropertiesResponsePtrOutput) Elem() VideoCreationPropertiesResponseOutput {
	return o.ApplyT(func(v *VideoCreationPropertiesResponse) VideoCreationPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret VideoCreationPropertiesResponse
		return ret
	}).(VideoCreationPropertiesResponseOutput)
}

func (o VideoCreationPropertiesResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoCreationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesResponsePtrOutput) RetentionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoCreationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RetentionPeriod
	}).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesResponsePtrOutput) SegmentLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoCreationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SegmentLength
	}).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesResponsePtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoCreationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Title
	}).(pulumi.StringPtrOutput)
}

type VideoEncoderH264 struct {
	BitrateKbps *string     `pulumi:"bitrateKbps"`
	FrameRate   *string     `pulumi:"frameRate"`
	Scale       *VideoScale `pulumi:"scale"`
	Type        string      `pulumi:"type"`
}

type VideoEncoderH264Response struct {
	BitrateKbps *string             `pulumi:"bitrateKbps"`
	FrameRate   *string             `pulumi:"frameRate"`
	Scale       *VideoScaleResponse `pulumi:"scale"`
	Type        string              `pulumi:"type"`
}

type VideoFlagsResponse struct {
	CanStream bool `pulumi:"canStream"`
	HasData   bool `pulumi:"hasData"`
	IsInUse   bool `pulumi:"isInUse"`
}

type VideoFlagsResponseOutput struct{ *pulumi.OutputState }

func (VideoFlagsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoFlagsResponse)(nil)).Elem()
}

func (o VideoFlagsResponseOutput) ToVideoFlagsResponseOutput() VideoFlagsResponseOutput {
	return o
}

func (o VideoFlagsResponseOutput) ToVideoFlagsResponseOutputWithContext(ctx context.Context) VideoFlagsResponseOutput {
	return o
}

func (o VideoFlagsResponseOutput) CanStream() pulumi.BoolOutput {
	return o.ApplyT(func(v VideoFlagsResponse) bool { return v.CanStream }).(pulumi.BoolOutput)
}

func (o VideoFlagsResponseOutput) HasData() pulumi.BoolOutput {
	return o.ApplyT(func(v VideoFlagsResponse) bool { return v.HasData }).(pulumi.BoolOutput)
}

func (o VideoFlagsResponseOutput) IsInUse() pulumi.BoolOutput {
	return o.ApplyT(func(v VideoFlagsResponse) bool { return v.IsInUse }).(pulumi.BoolOutput)
}

type VideoMediaInfo struct {
	SegmentLength *string `pulumi:"segmentLength"`
}





type VideoMediaInfoInput interface {
	pulumi.Input

	ToVideoMediaInfoOutput() VideoMediaInfoOutput
	ToVideoMediaInfoOutputWithContext(context.Context) VideoMediaInfoOutput
}

type VideoMediaInfoArgs struct {
	SegmentLength pulumi.StringPtrInput `pulumi:"segmentLength"`
}

func (VideoMediaInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoMediaInfo)(nil)).Elem()
}

func (i VideoMediaInfoArgs) ToVideoMediaInfoOutput() VideoMediaInfoOutput {
	return i.ToVideoMediaInfoOutputWithContext(context.Background())
}

func (i VideoMediaInfoArgs) ToVideoMediaInfoOutputWithContext(ctx context.Context) VideoMediaInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoMediaInfoOutput)
}

func (i VideoMediaInfoArgs) ToVideoMediaInfoPtrOutput() VideoMediaInfoPtrOutput {
	return i.ToVideoMediaInfoPtrOutputWithContext(context.Background())
}

func (i VideoMediaInfoArgs) ToVideoMediaInfoPtrOutputWithContext(ctx context.Context) VideoMediaInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoMediaInfoOutput).ToVideoMediaInfoPtrOutputWithContext(ctx)
}









type VideoMediaInfoPtrInput interface {
	pulumi.Input

	ToVideoMediaInfoPtrOutput() VideoMediaInfoPtrOutput
	ToVideoMediaInfoPtrOutputWithContext(context.Context) VideoMediaInfoPtrOutput
}

type videoMediaInfoPtrType VideoMediaInfoArgs

func VideoMediaInfoPtr(v *VideoMediaInfoArgs) VideoMediaInfoPtrInput {
	return (*videoMediaInfoPtrType)(v)
}

func (*videoMediaInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoMediaInfo)(nil)).Elem()
}

func (i *videoMediaInfoPtrType) ToVideoMediaInfoPtrOutput() VideoMediaInfoPtrOutput {
	return i.ToVideoMediaInfoPtrOutputWithContext(context.Background())
}

func (i *videoMediaInfoPtrType) ToVideoMediaInfoPtrOutputWithContext(ctx context.Context) VideoMediaInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoMediaInfoPtrOutput)
}

type VideoMediaInfoOutput struct{ *pulumi.OutputState }

func (VideoMediaInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoMediaInfo)(nil)).Elem()
}

func (o VideoMediaInfoOutput) ToVideoMediaInfoOutput() VideoMediaInfoOutput {
	return o
}

func (o VideoMediaInfoOutput) ToVideoMediaInfoOutputWithContext(ctx context.Context) VideoMediaInfoOutput {
	return o
}

func (o VideoMediaInfoOutput) ToVideoMediaInfoPtrOutput() VideoMediaInfoPtrOutput {
	return o.ToVideoMediaInfoPtrOutputWithContext(context.Background())
}

func (o VideoMediaInfoOutput) ToVideoMediaInfoPtrOutputWithContext(ctx context.Context) VideoMediaInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VideoMediaInfo) *VideoMediaInfo {
		return &v
	}).(VideoMediaInfoPtrOutput)
}

func (o VideoMediaInfoOutput) SegmentLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoMediaInfo) *string { return v.SegmentLength }).(pulumi.StringPtrOutput)
}

type VideoMediaInfoPtrOutput struct{ *pulumi.OutputState }

func (VideoMediaInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoMediaInfo)(nil)).Elem()
}

func (o VideoMediaInfoPtrOutput) ToVideoMediaInfoPtrOutput() VideoMediaInfoPtrOutput {
	return o
}

func (o VideoMediaInfoPtrOutput) ToVideoMediaInfoPtrOutputWithContext(ctx context.Context) VideoMediaInfoPtrOutput {
	return o
}

func (o VideoMediaInfoPtrOutput) Elem() VideoMediaInfoOutput {
	return o.ApplyT(func(v *VideoMediaInfo) VideoMediaInfo {
		if v != nil {
			return *v
		}
		var ret VideoMediaInfo
		return ret
	}).(VideoMediaInfoOutput)
}

func (o VideoMediaInfoPtrOutput) SegmentLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoMediaInfo) *string {
		if v == nil {
			return nil
		}
		return v.SegmentLength
	}).(pulumi.StringPtrOutput)
}

type VideoMediaInfoResponse struct {
	SegmentLength *string `pulumi:"segmentLength"`
}

type VideoMediaInfoResponseOutput struct{ *pulumi.OutputState }

func (VideoMediaInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoMediaInfoResponse)(nil)).Elem()
}

func (o VideoMediaInfoResponseOutput) ToVideoMediaInfoResponseOutput() VideoMediaInfoResponseOutput {
	return o
}

func (o VideoMediaInfoResponseOutput) ToVideoMediaInfoResponseOutputWithContext(ctx context.Context) VideoMediaInfoResponseOutput {
	return o
}

func (o VideoMediaInfoResponseOutput) SegmentLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoMediaInfoResponse) *string { return v.SegmentLength }).(pulumi.StringPtrOutput)
}

type VideoMediaInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (VideoMediaInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoMediaInfoResponse)(nil)).Elem()
}

func (o VideoMediaInfoResponsePtrOutput) ToVideoMediaInfoResponsePtrOutput() VideoMediaInfoResponsePtrOutput {
	return o
}

func (o VideoMediaInfoResponsePtrOutput) ToVideoMediaInfoResponsePtrOutputWithContext(ctx context.Context) VideoMediaInfoResponsePtrOutput {
	return o
}

func (o VideoMediaInfoResponsePtrOutput) Elem() VideoMediaInfoResponseOutput {
	return o.ApplyT(func(v *VideoMediaInfoResponse) VideoMediaInfoResponse {
		if v != nil {
			return *v
		}
		var ret VideoMediaInfoResponse
		return ret
	}).(VideoMediaInfoResponseOutput)
}

func (o VideoMediaInfoResponsePtrOutput) SegmentLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoMediaInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.SegmentLength
	}).(pulumi.StringPtrOutput)
}

type VideoPreviewImageUrlsResponse struct {
	Large  *string `pulumi:"large"`
	Medium *string `pulumi:"medium"`
	Small  *string `pulumi:"small"`
}

type VideoPreviewImageUrlsResponseOutput struct{ *pulumi.OutputState }

func (VideoPreviewImageUrlsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoPreviewImageUrlsResponse)(nil)).Elem()
}

func (o VideoPreviewImageUrlsResponseOutput) ToVideoPreviewImageUrlsResponseOutput() VideoPreviewImageUrlsResponseOutput {
	return o
}

func (o VideoPreviewImageUrlsResponseOutput) ToVideoPreviewImageUrlsResponseOutputWithContext(ctx context.Context) VideoPreviewImageUrlsResponseOutput {
	return o
}

func (o VideoPreviewImageUrlsResponseOutput) Large() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoPreviewImageUrlsResponse) *string { return v.Large }).(pulumi.StringPtrOutput)
}

func (o VideoPreviewImageUrlsResponseOutput) Medium() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoPreviewImageUrlsResponse) *string { return v.Medium }).(pulumi.StringPtrOutput)
}

func (o VideoPreviewImageUrlsResponseOutput) Small() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoPreviewImageUrlsResponse) *string { return v.Small }).(pulumi.StringPtrOutput)
}

type VideoPreviewImageUrlsResponsePtrOutput struct{ *pulumi.OutputState }

func (VideoPreviewImageUrlsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoPreviewImageUrlsResponse)(nil)).Elem()
}

func (o VideoPreviewImageUrlsResponsePtrOutput) ToVideoPreviewImageUrlsResponsePtrOutput() VideoPreviewImageUrlsResponsePtrOutput {
	return o
}

func (o VideoPreviewImageUrlsResponsePtrOutput) ToVideoPreviewImageUrlsResponsePtrOutputWithContext(ctx context.Context) VideoPreviewImageUrlsResponsePtrOutput {
	return o
}

func (o VideoPreviewImageUrlsResponsePtrOutput) Elem() VideoPreviewImageUrlsResponseOutput {
	return o.ApplyT(func(v *VideoPreviewImageUrlsResponse) VideoPreviewImageUrlsResponse {
		if v != nil {
			return *v
		}
		var ret VideoPreviewImageUrlsResponse
		return ret
	}).(VideoPreviewImageUrlsResponseOutput)
}

func (o VideoPreviewImageUrlsResponsePtrOutput) Large() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoPreviewImageUrlsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Large
	}).(pulumi.StringPtrOutput)
}

func (o VideoPreviewImageUrlsResponsePtrOutput) Medium() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoPreviewImageUrlsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Medium
	}).(pulumi.StringPtrOutput)
}

func (o VideoPreviewImageUrlsResponsePtrOutput) Small() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoPreviewImageUrlsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Small
	}).(pulumi.StringPtrOutput)
}

type VideoPublishingOptions struct {
	DisableArchive        *string `pulumi:"disableArchive"`
	DisableRtspPublishing *string `pulumi:"disableRtspPublishing"`
}





type VideoPublishingOptionsInput interface {
	pulumi.Input

	ToVideoPublishingOptionsOutput() VideoPublishingOptionsOutput
	ToVideoPublishingOptionsOutputWithContext(context.Context) VideoPublishingOptionsOutput
}

type VideoPublishingOptionsArgs struct {
	DisableArchive        pulumi.StringPtrInput `pulumi:"disableArchive"`
	DisableRtspPublishing pulumi.StringPtrInput `pulumi:"disableRtspPublishing"`
}

func (VideoPublishingOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoPublishingOptions)(nil)).Elem()
}

func (i VideoPublishingOptionsArgs) ToVideoPublishingOptionsOutput() VideoPublishingOptionsOutput {
	return i.ToVideoPublishingOptionsOutputWithContext(context.Background())
}

func (i VideoPublishingOptionsArgs) ToVideoPublishingOptionsOutputWithContext(ctx context.Context) VideoPublishingOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoPublishingOptionsOutput)
}

func (i VideoPublishingOptionsArgs) ToVideoPublishingOptionsPtrOutput() VideoPublishingOptionsPtrOutput {
	return i.ToVideoPublishingOptionsPtrOutputWithContext(context.Background())
}

func (i VideoPublishingOptionsArgs) ToVideoPublishingOptionsPtrOutputWithContext(ctx context.Context) VideoPublishingOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoPublishingOptionsOutput).ToVideoPublishingOptionsPtrOutputWithContext(ctx)
}









type VideoPublishingOptionsPtrInput interface {
	pulumi.Input

	ToVideoPublishingOptionsPtrOutput() VideoPublishingOptionsPtrOutput
	ToVideoPublishingOptionsPtrOutputWithContext(context.Context) VideoPublishingOptionsPtrOutput
}

type videoPublishingOptionsPtrType VideoPublishingOptionsArgs

func VideoPublishingOptionsPtr(v *VideoPublishingOptionsArgs) VideoPublishingOptionsPtrInput {
	return (*videoPublishingOptionsPtrType)(v)
}

func (*videoPublishingOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoPublishingOptions)(nil)).Elem()
}

func (i *videoPublishingOptionsPtrType) ToVideoPublishingOptionsPtrOutput() VideoPublishingOptionsPtrOutput {
	return i.ToVideoPublishingOptionsPtrOutputWithContext(context.Background())
}

func (i *videoPublishingOptionsPtrType) ToVideoPublishingOptionsPtrOutputWithContext(ctx context.Context) VideoPublishingOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoPublishingOptionsPtrOutput)
}

type VideoPublishingOptionsOutput struct{ *pulumi.OutputState }

func (VideoPublishingOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoPublishingOptions)(nil)).Elem()
}

func (o VideoPublishingOptionsOutput) ToVideoPublishingOptionsOutput() VideoPublishingOptionsOutput {
	return o
}

func (o VideoPublishingOptionsOutput) ToVideoPublishingOptionsOutputWithContext(ctx context.Context) VideoPublishingOptionsOutput {
	return o
}

func (o VideoPublishingOptionsOutput) ToVideoPublishingOptionsPtrOutput() VideoPublishingOptionsPtrOutput {
	return o.ToVideoPublishingOptionsPtrOutputWithContext(context.Background())
}

func (o VideoPublishingOptionsOutput) ToVideoPublishingOptionsPtrOutputWithContext(ctx context.Context) VideoPublishingOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VideoPublishingOptions) *VideoPublishingOptions {
		return &v
	}).(VideoPublishingOptionsPtrOutput)
}

func (o VideoPublishingOptionsOutput) DisableArchive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoPublishingOptions) *string { return v.DisableArchive }).(pulumi.StringPtrOutput)
}

func (o VideoPublishingOptionsOutput) DisableRtspPublishing() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoPublishingOptions) *string { return v.DisableRtspPublishing }).(pulumi.StringPtrOutput)
}

type VideoPublishingOptionsPtrOutput struct{ *pulumi.OutputState }

func (VideoPublishingOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoPublishingOptions)(nil)).Elem()
}

func (o VideoPublishingOptionsPtrOutput) ToVideoPublishingOptionsPtrOutput() VideoPublishingOptionsPtrOutput {
	return o
}

func (o VideoPublishingOptionsPtrOutput) ToVideoPublishingOptionsPtrOutputWithContext(ctx context.Context) VideoPublishingOptionsPtrOutput {
	return o
}

func (o VideoPublishingOptionsPtrOutput) Elem() VideoPublishingOptionsOutput {
	return o.ApplyT(func(v *VideoPublishingOptions) VideoPublishingOptions {
		if v != nil {
			return *v
		}
		var ret VideoPublishingOptions
		return ret
	}).(VideoPublishingOptionsOutput)
}

func (o VideoPublishingOptionsPtrOutput) DisableArchive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoPublishingOptions) *string {
		if v == nil {
			return nil
		}
		return v.DisableArchive
	}).(pulumi.StringPtrOutput)
}

func (o VideoPublishingOptionsPtrOutput) DisableRtspPublishing() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoPublishingOptions) *string {
		if v == nil {
			return nil
		}
		return v.DisableRtspPublishing
	}).(pulumi.StringPtrOutput)
}

type VideoPublishingOptionsResponse struct {
	DisableArchive        *string `pulumi:"disableArchive"`
	DisableRtspPublishing *string `pulumi:"disableRtspPublishing"`
}

type VideoPublishingOptionsResponseOutput struct{ *pulumi.OutputState }

func (VideoPublishingOptionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoPublishingOptionsResponse)(nil)).Elem()
}

func (o VideoPublishingOptionsResponseOutput) ToVideoPublishingOptionsResponseOutput() VideoPublishingOptionsResponseOutput {
	return o
}

func (o VideoPublishingOptionsResponseOutput) ToVideoPublishingOptionsResponseOutputWithContext(ctx context.Context) VideoPublishingOptionsResponseOutput {
	return o
}

func (o VideoPublishingOptionsResponseOutput) DisableArchive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoPublishingOptionsResponse) *string { return v.DisableArchive }).(pulumi.StringPtrOutput)
}

func (o VideoPublishingOptionsResponseOutput) DisableRtspPublishing() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoPublishingOptionsResponse) *string { return v.DisableRtspPublishing }).(pulumi.StringPtrOutput)
}

type VideoPublishingOptionsResponsePtrOutput struct{ *pulumi.OutputState }

func (VideoPublishingOptionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoPublishingOptionsResponse)(nil)).Elem()
}

func (o VideoPublishingOptionsResponsePtrOutput) ToVideoPublishingOptionsResponsePtrOutput() VideoPublishingOptionsResponsePtrOutput {
	return o
}

func (o VideoPublishingOptionsResponsePtrOutput) ToVideoPublishingOptionsResponsePtrOutputWithContext(ctx context.Context) VideoPublishingOptionsResponsePtrOutput {
	return o
}

func (o VideoPublishingOptionsResponsePtrOutput) Elem() VideoPublishingOptionsResponseOutput {
	return o.ApplyT(func(v *VideoPublishingOptionsResponse) VideoPublishingOptionsResponse {
		if v != nil {
			return *v
		}
		var ret VideoPublishingOptionsResponse
		return ret
	}).(VideoPublishingOptionsResponseOutput)
}

func (o VideoPublishingOptionsResponsePtrOutput) DisableArchive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoPublishingOptionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DisableArchive
	}).(pulumi.StringPtrOutput)
}

func (o VideoPublishingOptionsResponsePtrOutput) DisableRtspPublishing() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoPublishingOptionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DisableRtspPublishing
	}).(pulumi.StringPtrOutput)
}

type VideoScale struct {
	Height *string `pulumi:"height"`
	Mode   *string `pulumi:"mode"`
	Width  *string `pulumi:"width"`
}

type VideoScaleResponse struct {
	Height *string `pulumi:"height"`
	Mode   *string `pulumi:"mode"`
	Width  *string `pulumi:"width"`
}

type VideoSequenceAbsoluteTimeMarkers struct {
	Ranges string `pulumi:"ranges"`
	Type   string `pulumi:"type"`
}

type VideoSequenceAbsoluteTimeMarkersResponse struct {
	Ranges string `pulumi:"ranges"`
	Type   string `pulumi:"type"`
}

type VideoSink struct {
	Inputs                  []NodeInput              `pulumi:"inputs"`
	Name                    string                   `pulumi:"name"`
	Type                    string                   `pulumi:"type"`
	VideoCreationProperties *VideoCreationProperties `pulumi:"videoCreationProperties"`
	VideoName               string                   `pulumi:"videoName"`
	VideoPublishingOptions  *VideoPublishingOptions  `pulumi:"videoPublishingOptions"`
}





type VideoSinkInput interface {
	pulumi.Input

	ToVideoSinkOutput() VideoSinkOutput
	ToVideoSinkOutputWithContext(context.Context) VideoSinkOutput
}

type VideoSinkArgs struct {
	Inputs                  NodeInputArrayInput             `pulumi:"inputs"`
	Name                    pulumi.StringInput              `pulumi:"name"`
	Type                    pulumi.StringInput              `pulumi:"type"`
	VideoCreationProperties VideoCreationPropertiesPtrInput `pulumi:"videoCreationProperties"`
	VideoName               pulumi.StringInput              `pulumi:"videoName"`
	VideoPublishingOptions  VideoPublishingOptionsPtrInput  `pulumi:"videoPublishingOptions"`
}

func (VideoSinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoSink)(nil)).Elem()
}

func (i VideoSinkArgs) ToVideoSinkOutput() VideoSinkOutput {
	return i.ToVideoSinkOutputWithContext(context.Background())
}

func (i VideoSinkArgs) ToVideoSinkOutputWithContext(ctx context.Context) VideoSinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoSinkOutput)
}





type VideoSinkArrayInput interface {
	pulumi.Input

	ToVideoSinkArrayOutput() VideoSinkArrayOutput
	ToVideoSinkArrayOutputWithContext(context.Context) VideoSinkArrayOutput
}

type VideoSinkArray []VideoSinkInput

func (VideoSinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VideoSink)(nil)).Elem()
}

func (i VideoSinkArray) ToVideoSinkArrayOutput() VideoSinkArrayOutput {
	return i.ToVideoSinkArrayOutputWithContext(context.Background())
}

func (i VideoSinkArray) ToVideoSinkArrayOutputWithContext(ctx context.Context) VideoSinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoSinkArrayOutput)
}

type VideoSinkOutput struct{ *pulumi.OutputState }

func (VideoSinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoSink)(nil)).Elem()
}

func (o VideoSinkOutput) ToVideoSinkOutput() VideoSinkOutput {
	return o
}

func (o VideoSinkOutput) ToVideoSinkOutputWithContext(ctx context.Context) VideoSinkOutput {
	return o
}

func (o VideoSinkOutput) Inputs() NodeInputArrayOutput {
	return o.ApplyT(func(v VideoSink) []NodeInput { return v.Inputs }).(NodeInputArrayOutput)
}

func (o VideoSinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSink) string { return v.Name }).(pulumi.StringOutput)
}

func (o VideoSinkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSink) string { return v.Type }).(pulumi.StringOutput)
}

func (o VideoSinkOutput) VideoCreationProperties() VideoCreationPropertiesPtrOutput {
	return o.ApplyT(func(v VideoSink) *VideoCreationProperties { return v.VideoCreationProperties }).(VideoCreationPropertiesPtrOutput)
}

func (o VideoSinkOutput) VideoName() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSink) string { return v.VideoName }).(pulumi.StringOutput)
}

func (o VideoSinkOutput) VideoPublishingOptions() VideoPublishingOptionsPtrOutput {
	return o.ApplyT(func(v VideoSink) *VideoPublishingOptions { return v.VideoPublishingOptions }).(VideoPublishingOptionsPtrOutput)
}

type VideoSinkArrayOutput struct{ *pulumi.OutputState }

func (VideoSinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VideoSink)(nil)).Elem()
}

func (o VideoSinkArrayOutput) ToVideoSinkArrayOutput() VideoSinkArrayOutput {
	return o
}

func (o VideoSinkArrayOutput) ToVideoSinkArrayOutputWithContext(ctx context.Context) VideoSinkArrayOutput {
	return o
}

func (o VideoSinkArrayOutput) Index(i pulumi.IntInput) VideoSinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VideoSink {
		return vs[0].([]VideoSink)[vs[1].(int)]
	}).(VideoSinkOutput)
}

type VideoSinkResponse struct {
	Inputs                  []NodeInputResponse              `pulumi:"inputs"`
	Name                    string                           `pulumi:"name"`
	Type                    string                           `pulumi:"type"`
	VideoCreationProperties *VideoCreationPropertiesResponse `pulumi:"videoCreationProperties"`
	VideoName               string                           `pulumi:"videoName"`
	VideoPublishingOptions  *VideoPublishingOptionsResponse  `pulumi:"videoPublishingOptions"`
}

type VideoSinkResponseOutput struct{ *pulumi.OutputState }

func (VideoSinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoSinkResponse)(nil)).Elem()
}

func (o VideoSinkResponseOutput) ToVideoSinkResponseOutput() VideoSinkResponseOutput {
	return o
}

func (o VideoSinkResponseOutput) ToVideoSinkResponseOutputWithContext(ctx context.Context) VideoSinkResponseOutput {
	return o
}

func (o VideoSinkResponseOutput) Inputs() NodeInputResponseArrayOutput {
	return o.ApplyT(func(v VideoSinkResponse) []NodeInputResponse { return v.Inputs }).(NodeInputResponseArrayOutput)
}

func (o VideoSinkResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSinkResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VideoSinkResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSinkResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VideoSinkResponseOutput) VideoCreationProperties() VideoCreationPropertiesResponsePtrOutput {
	return o.ApplyT(func(v VideoSinkResponse) *VideoCreationPropertiesResponse { return v.VideoCreationProperties }).(VideoCreationPropertiesResponsePtrOutput)
}

func (o VideoSinkResponseOutput) VideoName() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSinkResponse) string { return v.VideoName }).(pulumi.StringOutput)
}

func (o VideoSinkResponseOutput) VideoPublishingOptions() VideoPublishingOptionsResponsePtrOutput {
	return o.ApplyT(func(v VideoSinkResponse) *VideoPublishingOptionsResponse { return v.VideoPublishingOptions }).(VideoPublishingOptionsResponsePtrOutput)
}

type VideoSinkResponseArrayOutput struct{ *pulumi.OutputState }

func (VideoSinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VideoSinkResponse)(nil)).Elem()
}

func (o VideoSinkResponseArrayOutput) ToVideoSinkResponseArrayOutput() VideoSinkResponseArrayOutput {
	return o
}

func (o VideoSinkResponseArrayOutput) ToVideoSinkResponseArrayOutputWithContext(ctx context.Context) VideoSinkResponseArrayOutput {
	return o
}

func (o VideoSinkResponseArrayOutput) Index(i pulumi.IntInput) VideoSinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VideoSinkResponse {
		return vs[0].([]VideoSinkResponse)[vs[1].(int)]
	}).(VideoSinkResponseOutput)
}

type VideoSource struct {
	Name          string                           `pulumi:"name"`
	TimeSequences VideoSequenceAbsoluteTimeMarkers `pulumi:"timeSequences"`
	Type          string                           `pulumi:"type"`
	VideoName     string                           `pulumi:"videoName"`
}

type VideoSourceResponse struct {
	Name          string                                   `pulumi:"name"`
	TimeSequences VideoSequenceAbsoluteTimeMarkersResponse `pulumi:"timeSequences"`
	Type          string                                   `pulumi:"type"`
	VideoName     string                                   `pulumi:"videoName"`
}

func init() {
	pulumi.RegisterOutputType(AccountEncryptionOutput{})
	pulumi.RegisterOutputType(AccountEncryptionPtrOutput{})
	pulumi.RegisterOutputType(AccountEncryptionResponseOutput{})
	pulumi.RegisterOutputType(AccountEncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(EncoderProcessorOutput{})
	pulumi.RegisterOutputType(EncoderProcessorArrayOutput{})
	pulumi.RegisterOutputType(EncoderProcessorResponseOutput{})
	pulumi.RegisterOutputType(EncoderProcessorResponseArrayOutput{})
	pulumi.RegisterOutputType(EndpointResponseOutput{})
	pulumi.RegisterOutputType(EndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(GroupLevelAccessControlOutput{})
	pulumi.RegisterOutputType(GroupLevelAccessControlPtrOutput{})
	pulumi.RegisterOutputType(GroupLevelAccessControlResponseOutput{})
	pulumi.RegisterOutputType(GroupLevelAccessControlResponsePtrOutput{})
	pulumi.RegisterOutputType(IotHubOutput{})
	pulumi.RegisterOutputType(IotHubArrayOutput{})
	pulumi.RegisterOutputType(IotHubResponseOutput{})
	pulumi.RegisterOutputType(IotHubResponseArrayOutput{})
	pulumi.RegisterOutputType(JwtAuthenticationOutput{})
	pulumi.RegisterOutputType(JwtAuthenticationPtrOutput{})
	pulumi.RegisterOutputType(JwtAuthenticationResponseOutput{})
	pulumi.RegisterOutputType(JwtAuthenticationResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkAccessControlOutput{})
	pulumi.RegisterOutputType(NetworkAccessControlPtrOutput{})
	pulumi.RegisterOutputType(NetworkAccessControlResponseOutput{})
	pulumi.RegisterOutputType(NetworkAccessControlResponsePtrOutput{})
	pulumi.RegisterOutputType(NodeInputOutput{})
	pulumi.RegisterOutputType(NodeInputArrayOutput{})
	pulumi.RegisterOutputType(NodeInputResponseOutput{})
	pulumi.RegisterOutputType(NodeInputResponseArrayOutput{})
	pulumi.RegisterOutputType(ParameterDeclarationOutput{})
	pulumi.RegisterOutputType(ParameterDeclarationArrayOutput{})
	pulumi.RegisterOutputType(ParameterDeclarationResponseOutput{})
	pulumi.RegisterOutputType(ParameterDeclarationResponseArrayOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionArrayOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(PipelineJobErrorResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdentityOutput{})
	pulumi.RegisterOutputType(ResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountOutput{})
	pulumi.RegisterOutputType(StorageAccountArrayOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TokenClaimOutput{})
	pulumi.RegisterOutputType(TokenClaimArrayOutput{})
	pulumi.RegisterOutputType(TokenClaimResponseOutput{})
	pulumi.RegisterOutputType(TokenClaimResponseArrayOutput{})
	pulumi.RegisterOutputType(UserAssignedManagedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedManagedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(VideoAnalyzerIdentityOutput{})
	pulumi.RegisterOutputType(VideoAnalyzerIdentityPtrOutput{})
	pulumi.RegisterOutputType(VideoAnalyzerIdentityResponseOutput{})
	pulumi.RegisterOutputType(VideoAnalyzerIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(VideoArchivalOutput{})
	pulumi.RegisterOutputType(VideoArchivalPtrOutput{})
	pulumi.RegisterOutputType(VideoArchivalResponseOutput{})
	pulumi.RegisterOutputType(VideoArchivalResponsePtrOutput{})
	pulumi.RegisterOutputType(VideoContentUrlsResponseOutput{})
	pulumi.RegisterOutputType(VideoCreationPropertiesOutput{})
	pulumi.RegisterOutputType(VideoCreationPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VideoCreationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VideoCreationPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(VideoFlagsResponseOutput{})
	pulumi.RegisterOutputType(VideoMediaInfoOutput{})
	pulumi.RegisterOutputType(VideoMediaInfoPtrOutput{})
	pulumi.RegisterOutputType(VideoMediaInfoResponseOutput{})
	pulumi.RegisterOutputType(VideoMediaInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(VideoPreviewImageUrlsResponseOutput{})
	pulumi.RegisterOutputType(VideoPreviewImageUrlsResponsePtrOutput{})
	pulumi.RegisterOutputType(VideoPublishingOptionsOutput{})
	pulumi.RegisterOutputType(VideoPublishingOptionsPtrOutput{})
	pulumi.RegisterOutputType(VideoPublishingOptionsResponseOutput{})
	pulumi.RegisterOutputType(VideoPublishingOptionsResponsePtrOutput{})
	pulumi.RegisterOutputType(VideoSinkOutput{})
	pulumi.RegisterOutputType(VideoSinkArrayOutput{})
	pulumi.RegisterOutputType(VideoSinkResponseOutput{})
	pulumi.RegisterOutputType(VideoSinkResponseArrayOutput{})
}

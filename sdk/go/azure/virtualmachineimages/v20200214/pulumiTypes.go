


package v20200214

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ImageTemplateFileCustomizer struct {
	Destination    *string `pulumi:"destination"`
	Name           *string `pulumi:"name"`
	Sha256Checksum *string `pulumi:"sha256Checksum"`
	SourceUri      *string `pulumi:"sourceUri"`
	Type           string  `pulumi:"type"`
}





type ImageTemplateFileCustomizerInput interface {
	pulumi.Input

	ToImageTemplateFileCustomizerOutput() ImageTemplateFileCustomizerOutput
	ToImageTemplateFileCustomizerOutputWithContext(context.Context) ImageTemplateFileCustomizerOutput
}

type ImageTemplateFileCustomizerArgs struct {
	Destination    pulumi.StringPtrInput `pulumi:"destination"`
	Name           pulumi.StringPtrInput `pulumi:"name"`
	Sha256Checksum pulumi.StringPtrInput `pulumi:"sha256Checksum"`
	SourceUri      pulumi.StringPtrInput `pulumi:"sourceUri"`
	Type           pulumi.StringInput    `pulumi:"type"`
}

func (ImageTemplateFileCustomizerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateFileCustomizer)(nil)).Elem()
}

func (i ImageTemplateFileCustomizerArgs) ToImageTemplateFileCustomizerOutput() ImageTemplateFileCustomizerOutput {
	return i.ToImageTemplateFileCustomizerOutputWithContext(context.Background())
}

func (i ImageTemplateFileCustomizerArgs) ToImageTemplateFileCustomizerOutputWithContext(ctx context.Context) ImageTemplateFileCustomizerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateFileCustomizerOutput)
}

type ImageTemplateFileCustomizerOutput struct{ *pulumi.OutputState }

func (ImageTemplateFileCustomizerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateFileCustomizer)(nil)).Elem()
}

func (o ImageTemplateFileCustomizerOutput) ToImageTemplateFileCustomizerOutput() ImageTemplateFileCustomizerOutput {
	return o
}

func (o ImageTemplateFileCustomizerOutput) ToImageTemplateFileCustomizerOutputWithContext(ctx context.Context) ImageTemplateFileCustomizerOutput {
	return o
}

func (o ImageTemplateFileCustomizerOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateFileCustomizer) *string { return v.Destination }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateFileCustomizerOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateFileCustomizer) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateFileCustomizerOutput) Sha256Checksum() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateFileCustomizer) *string { return v.Sha256Checksum }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateFileCustomizerOutput) SourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateFileCustomizer) *string { return v.SourceUri }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateFileCustomizerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateFileCustomizer) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateFileCustomizerResponse struct {
	Destination    *string `pulumi:"destination"`
	Name           *string `pulumi:"name"`
	Sha256Checksum *string `pulumi:"sha256Checksum"`
	SourceUri      *string `pulumi:"sourceUri"`
	Type           string  `pulumi:"type"`
}





type ImageTemplateFileCustomizerResponseInput interface {
	pulumi.Input

	ToImageTemplateFileCustomizerResponseOutput() ImageTemplateFileCustomizerResponseOutput
	ToImageTemplateFileCustomizerResponseOutputWithContext(context.Context) ImageTemplateFileCustomizerResponseOutput
}

type ImageTemplateFileCustomizerResponseArgs struct {
	Destination    pulumi.StringPtrInput `pulumi:"destination"`
	Name           pulumi.StringPtrInput `pulumi:"name"`
	Sha256Checksum pulumi.StringPtrInput `pulumi:"sha256Checksum"`
	SourceUri      pulumi.StringPtrInput `pulumi:"sourceUri"`
	Type           pulumi.StringInput    `pulumi:"type"`
}

func (ImageTemplateFileCustomizerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateFileCustomizerResponse)(nil)).Elem()
}

func (i ImageTemplateFileCustomizerResponseArgs) ToImageTemplateFileCustomizerResponseOutput() ImageTemplateFileCustomizerResponseOutput {
	return i.ToImageTemplateFileCustomizerResponseOutputWithContext(context.Background())
}

func (i ImageTemplateFileCustomizerResponseArgs) ToImageTemplateFileCustomizerResponseOutputWithContext(ctx context.Context) ImageTemplateFileCustomizerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateFileCustomizerResponseOutput)
}

type ImageTemplateFileCustomizerResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplateFileCustomizerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateFileCustomizerResponse)(nil)).Elem()
}

func (o ImageTemplateFileCustomizerResponseOutput) ToImageTemplateFileCustomizerResponseOutput() ImageTemplateFileCustomizerResponseOutput {
	return o
}

func (o ImageTemplateFileCustomizerResponseOutput) ToImageTemplateFileCustomizerResponseOutputWithContext(ctx context.Context) ImageTemplateFileCustomizerResponseOutput {
	return o
}

func (o ImageTemplateFileCustomizerResponseOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateFileCustomizerResponse) *string { return v.Destination }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateFileCustomizerResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateFileCustomizerResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateFileCustomizerResponseOutput) Sha256Checksum() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateFileCustomizerResponse) *string { return v.Sha256Checksum }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateFileCustomizerResponseOutput) SourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateFileCustomizerResponse) *string { return v.SourceUri }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateFileCustomizerResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateFileCustomizerResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateIdentity struct {
	Type                   *ResourceIdentityType  `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ImageTemplateIdentityInput interface {
	pulumi.Input

	ToImageTemplateIdentityOutput() ImageTemplateIdentityOutput
	ToImageTemplateIdentityOutputWithContext(context.Context) ImageTemplateIdentityOutput
}

type ImageTemplateIdentityArgs struct {
	Type                   ResourceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput              `pulumi:"userAssignedIdentities"`
}

func (ImageTemplateIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateIdentity)(nil)).Elem()
}

func (i ImageTemplateIdentityArgs) ToImageTemplateIdentityOutput() ImageTemplateIdentityOutput {
	return i.ToImageTemplateIdentityOutputWithContext(context.Background())
}

func (i ImageTemplateIdentityArgs) ToImageTemplateIdentityOutputWithContext(ctx context.Context) ImageTemplateIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateIdentityOutput)
}

func (i ImageTemplateIdentityArgs) ToImageTemplateIdentityPtrOutput() ImageTemplateIdentityPtrOutput {
	return i.ToImageTemplateIdentityPtrOutputWithContext(context.Background())
}

func (i ImageTemplateIdentityArgs) ToImageTemplateIdentityPtrOutputWithContext(ctx context.Context) ImageTemplateIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateIdentityOutput).ToImageTemplateIdentityPtrOutputWithContext(ctx)
}









type ImageTemplateIdentityPtrInput interface {
	pulumi.Input

	ToImageTemplateIdentityPtrOutput() ImageTemplateIdentityPtrOutput
	ToImageTemplateIdentityPtrOutputWithContext(context.Context) ImageTemplateIdentityPtrOutput
}

type imageTemplateIdentityPtrType ImageTemplateIdentityArgs

func ImageTemplateIdentityPtr(v *ImageTemplateIdentityArgs) ImageTemplateIdentityPtrInput {
	return (*imageTemplateIdentityPtrType)(v)
}

func (*imageTemplateIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageTemplateIdentity)(nil)).Elem()
}

func (i *imageTemplateIdentityPtrType) ToImageTemplateIdentityPtrOutput() ImageTemplateIdentityPtrOutput {
	return i.ToImageTemplateIdentityPtrOutputWithContext(context.Background())
}

func (i *imageTemplateIdentityPtrType) ToImageTemplateIdentityPtrOutputWithContext(ctx context.Context) ImageTemplateIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateIdentityPtrOutput)
}

type ImageTemplateIdentityOutput struct{ *pulumi.OutputState }

func (ImageTemplateIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateIdentity)(nil)).Elem()
}

func (o ImageTemplateIdentityOutput) ToImageTemplateIdentityOutput() ImageTemplateIdentityOutput {
	return o
}

func (o ImageTemplateIdentityOutput) ToImageTemplateIdentityOutputWithContext(ctx context.Context) ImageTemplateIdentityOutput {
	return o
}

func (o ImageTemplateIdentityOutput) ToImageTemplateIdentityPtrOutput() ImageTemplateIdentityPtrOutput {
	return o.ToImageTemplateIdentityPtrOutputWithContext(context.Background())
}

func (o ImageTemplateIdentityOutput) ToImageTemplateIdentityPtrOutputWithContext(ctx context.Context) ImageTemplateIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageTemplateIdentity) *ImageTemplateIdentity {
		return &v
	}).(ImageTemplateIdentityPtrOutput)
}

func (o ImageTemplateIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v ImageTemplateIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

func (o ImageTemplateIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ImageTemplateIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ImageTemplateIdentityPtrOutput struct{ *pulumi.OutputState }

func (ImageTemplateIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageTemplateIdentity)(nil)).Elem()
}

func (o ImageTemplateIdentityPtrOutput) ToImageTemplateIdentityPtrOutput() ImageTemplateIdentityPtrOutput {
	return o
}

func (o ImageTemplateIdentityPtrOutput) ToImageTemplateIdentityPtrOutputWithContext(ctx context.Context) ImageTemplateIdentityPtrOutput {
	return o
}

func (o ImageTemplateIdentityPtrOutput) Elem() ImageTemplateIdentityOutput {
	return o.ApplyT(func(v *ImageTemplateIdentity) ImageTemplateIdentity {
		if v != nil {
			return *v
		}
		var ret ImageTemplateIdentity
		return ret
	}).(ImageTemplateIdentityOutput)
}

func (o ImageTemplateIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *ImageTemplateIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

func (o ImageTemplateIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ImageTemplateIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ImageTemplateIdentityResponse struct {
	Type                   *string                                                        `pulumi:"type"`
	UserAssignedIdentities map[string]ImageTemplateIdentityResponseUserAssignedIdentities `pulumi:"userAssignedIdentities"`
}





type ImageTemplateIdentityResponseInput interface {
	pulumi.Input

	ToImageTemplateIdentityResponseOutput() ImageTemplateIdentityResponseOutput
	ToImageTemplateIdentityResponseOutputWithContext(context.Context) ImageTemplateIdentityResponseOutput
}

type ImageTemplateIdentityResponseArgs struct {
	Type                   pulumi.StringPtrInput                                       `pulumi:"type"`
	UserAssignedIdentities ImageTemplateIdentityResponseUserAssignedIdentitiesMapInput `pulumi:"userAssignedIdentities"`
}

func (ImageTemplateIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateIdentityResponse)(nil)).Elem()
}

func (i ImageTemplateIdentityResponseArgs) ToImageTemplateIdentityResponseOutput() ImageTemplateIdentityResponseOutput {
	return i.ToImageTemplateIdentityResponseOutputWithContext(context.Background())
}

func (i ImageTemplateIdentityResponseArgs) ToImageTemplateIdentityResponseOutputWithContext(ctx context.Context) ImageTemplateIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateIdentityResponseOutput)
}

func (i ImageTemplateIdentityResponseArgs) ToImageTemplateIdentityResponsePtrOutput() ImageTemplateIdentityResponsePtrOutput {
	return i.ToImageTemplateIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ImageTemplateIdentityResponseArgs) ToImageTemplateIdentityResponsePtrOutputWithContext(ctx context.Context) ImageTemplateIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateIdentityResponseOutput).ToImageTemplateIdentityResponsePtrOutputWithContext(ctx)
}









type ImageTemplateIdentityResponsePtrInput interface {
	pulumi.Input

	ToImageTemplateIdentityResponsePtrOutput() ImageTemplateIdentityResponsePtrOutput
	ToImageTemplateIdentityResponsePtrOutputWithContext(context.Context) ImageTemplateIdentityResponsePtrOutput
}

type imageTemplateIdentityResponsePtrType ImageTemplateIdentityResponseArgs

func ImageTemplateIdentityResponsePtr(v *ImageTemplateIdentityResponseArgs) ImageTemplateIdentityResponsePtrInput {
	return (*imageTemplateIdentityResponsePtrType)(v)
}

func (*imageTemplateIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageTemplateIdentityResponse)(nil)).Elem()
}

func (i *imageTemplateIdentityResponsePtrType) ToImageTemplateIdentityResponsePtrOutput() ImageTemplateIdentityResponsePtrOutput {
	return i.ToImageTemplateIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *imageTemplateIdentityResponsePtrType) ToImageTemplateIdentityResponsePtrOutputWithContext(ctx context.Context) ImageTemplateIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateIdentityResponsePtrOutput)
}

type ImageTemplateIdentityResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplateIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateIdentityResponse)(nil)).Elem()
}

func (o ImageTemplateIdentityResponseOutput) ToImageTemplateIdentityResponseOutput() ImageTemplateIdentityResponseOutput {
	return o
}

func (o ImageTemplateIdentityResponseOutput) ToImageTemplateIdentityResponseOutputWithContext(ctx context.Context) ImageTemplateIdentityResponseOutput {
	return o
}

func (o ImageTemplateIdentityResponseOutput) ToImageTemplateIdentityResponsePtrOutput() ImageTemplateIdentityResponsePtrOutput {
	return o.ToImageTemplateIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ImageTemplateIdentityResponseOutput) ToImageTemplateIdentityResponsePtrOutputWithContext(ctx context.Context) ImageTemplateIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageTemplateIdentityResponse) *ImageTemplateIdentityResponse {
		return &v
	}).(ImageTemplateIdentityResponsePtrOutput)
}

func (o ImageTemplateIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateIdentityResponseOutput) UserAssignedIdentities() ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v ImageTemplateIdentityResponse) map[string]ImageTemplateIdentityResponseUserAssignedIdentities {
		return v.UserAssignedIdentities
	}).(ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput)
}

type ImageTemplateIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageTemplateIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageTemplateIdentityResponse)(nil)).Elem()
}

func (o ImageTemplateIdentityResponsePtrOutput) ToImageTemplateIdentityResponsePtrOutput() ImageTemplateIdentityResponsePtrOutput {
	return o
}

func (o ImageTemplateIdentityResponsePtrOutput) ToImageTemplateIdentityResponsePtrOutputWithContext(ctx context.Context) ImageTemplateIdentityResponsePtrOutput {
	return o
}

func (o ImageTemplateIdentityResponsePtrOutput) Elem() ImageTemplateIdentityResponseOutput {
	return o.ApplyT(func(v *ImageTemplateIdentityResponse) ImageTemplateIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ImageTemplateIdentityResponse
		return ret
	}).(ImageTemplateIdentityResponseOutput)
}

func (o ImageTemplateIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageTemplateIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ImageTemplateIdentityResponsePtrOutput) UserAssignedIdentities() ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v *ImageTemplateIdentityResponse) map[string]ImageTemplateIdentityResponseUserAssignedIdentities {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput)
}

type ImageTemplateIdentityResponseUserAssignedIdentities struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}





type ImageTemplateIdentityResponseUserAssignedIdentitiesInput interface {
	pulumi.Input

	ToImageTemplateIdentityResponseUserAssignedIdentitiesOutput() ImageTemplateIdentityResponseUserAssignedIdentitiesOutput
	ToImageTemplateIdentityResponseUserAssignedIdentitiesOutputWithContext(context.Context) ImageTemplateIdentityResponseUserAssignedIdentitiesOutput
}

type ImageTemplateIdentityResponseUserAssignedIdentitiesArgs struct {
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (ImageTemplateIdentityResponseUserAssignedIdentitiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (i ImageTemplateIdentityResponseUserAssignedIdentitiesArgs) ToImageTemplateIdentityResponseUserAssignedIdentitiesOutput() ImageTemplateIdentityResponseUserAssignedIdentitiesOutput {
	return i.ToImageTemplateIdentityResponseUserAssignedIdentitiesOutputWithContext(context.Background())
}

func (i ImageTemplateIdentityResponseUserAssignedIdentitiesArgs) ToImageTemplateIdentityResponseUserAssignedIdentitiesOutputWithContext(ctx context.Context) ImageTemplateIdentityResponseUserAssignedIdentitiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateIdentityResponseUserAssignedIdentitiesOutput)
}





type ImageTemplateIdentityResponseUserAssignedIdentitiesMapInput interface {
	pulumi.Input

	ToImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput() ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput
	ToImageTemplateIdentityResponseUserAssignedIdentitiesMapOutputWithContext(context.Context) ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput
}

type ImageTemplateIdentityResponseUserAssignedIdentitiesMap map[string]ImageTemplateIdentityResponseUserAssignedIdentitiesInput

func (ImageTemplateIdentityResponseUserAssignedIdentitiesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ImageTemplateIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (i ImageTemplateIdentityResponseUserAssignedIdentitiesMap) ToImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput() ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput {
	return i.ToImageTemplateIdentityResponseUserAssignedIdentitiesMapOutputWithContext(context.Background())
}

func (i ImageTemplateIdentityResponseUserAssignedIdentitiesMap) ToImageTemplateIdentityResponseUserAssignedIdentitiesMapOutputWithContext(ctx context.Context) ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput)
}

type ImageTemplateIdentityResponseUserAssignedIdentitiesOutput struct{ *pulumi.OutputState }

func (ImageTemplateIdentityResponseUserAssignedIdentitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o ImageTemplateIdentityResponseUserAssignedIdentitiesOutput) ToImageTemplateIdentityResponseUserAssignedIdentitiesOutput() ImageTemplateIdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o ImageTemplateIdentityResponseUserAssignedIdentitiesOutput) ToImageTemplateIdentityResponseUserAssignedIdentitiesOutputWithContext(ctx context.Context) ImageTemplateIdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o ImageTemplateIdentityResponseUserAssignedIdentitiesOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateIdentityResponseUserAssignedIdentities) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ImageTemplateIdentityResponseUserAssignedIdentitiesOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateIdentityResponseUserAssignedIdentities) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput struct{ *pulumi.OutputState }

func (ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ImageTemplateIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput) ToImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput() ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput) ToImageTemplateIdentityResponseUserAssignedIdentitiesMapOutputWithContext(ctx context.Context) ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput) MapIndex(k pulumi.StringInput) ImageTemplateIdentityResponseUserAssignedIdentitiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ImageTemplateIdentityResponseUserAssignedIdentities {
		return vs[0].(map[string]ImageTemplateIdentityResponseUserAssignedIdentities)[vs[1].(string)]
	}).(ImageTemplateIdentityResponseUserAssignedIdentitiesOutput)
}

type ImageTemplateLastRunStatusResponse struct {
	EndTime     *string `pulumi:"endTime"`
	Message     *string `pulumi:"message"`
	RunState    *string `pulumi:"runState"`
	RunSubState *string `pulumi:"runSubState"`
	StartTime   *string `pulumi:"startTime"`
}





type ImageTemplateLastRunStatusResponseInput interface {
	pulumi.Input

	ToImageTemplateLastRunStatusResponseOutput() ImageTemplateLastRunStatusResponseOutput
	ToImageTemplateLastRunStatusResponseOutputWithContext(context.Context) ImageTemplateLastRunStatusResponseOutput
}

type ImageTemplateLastRunStatusResponseArgs struct {
	EndTime     pulumi.StringPtrInput `pulumi:"endTime"`
	Message     pulumi.StringPtrInput `pulumi:"message"`
	RunState    pulumi.StringPtrInput `pulumi:"runState"`
	RunSubState pulumi.StringPtrInput `pulumi:"runSubState"`
	StartTime   pulumi.StringPtrInput `pulumi:"startTime"`
}

func (ImageTemplateLastRunStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateLastRunStatusResponse)(nil)).Elem()
}

func (i ImageTemplateLastRunStatusResponseArgs) ToImageTemplateLastRunStatusResponseOutput() ImageTemplateLastRunStatusResponseOutput {
	return i.ToImageTemplateLastRunStatusResponseOutputWithContext(context.Background())
}

func (i ImageTemplateLastRunStatusResponseArgs) ToImageTemplateLastRunStatusResponseOutputWithContext(ctx context.Context) ImageTemplateLastRunStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateLastRunStatusResponseOutput)
}

func (i ImageTemplateLastRunStatusResponseArgs) ToImageTemplateLastRunStatusResponsePtrOutput() ImageTemplateLastRunStatusResponsePtrOutput {
	return i.ToImageTemplateLastRunStatusResponsePtrOutputWithContext(context.Background())
}

func (i ImageTemplateLastRunStatusResponseArgs) ToImageTemplateLastRunStatusResponsePtrOutputWithContext(ctx context.Context) ImageTemplateLastRunStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateLastRunStatusResponseOutput).ToImageTemplateLastRunStatusResponsePtrOutputWithContext(ctx)
}









type ImageTemplateLastRunStatusResponsePtrInput interface {
	pulumi.Input

	ToImageTemplateLastRunStatusResponsePtrOutput() ImageTemplateLastRunStatusResponsePtrOutput
	ToImageTemplateLastRunStatusResponsePtrOutputWithContext(context.Context) ImageTemplateLastRunStatusResponsePtrOutput
}

type imageTemplateLastRunStatusResponsePtrType ImageTemplateLastRunStatusResponseArgs

func ImageTemplateLastRunStatusResponsePtr(v *ImageTemplateLastRunStatusResponseArgs) ImageTemplateLastRunStatusResponsePtrInput {
	return (*imageTemplateLastRunStatusResponsePtrType)(v)
}

func (*imageTemplateLastRunStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageTemplateLastRunStatusResponse)(nil)).Elem()
}

func (i *imageTemplateLastRunStatusResponsePtrType) ToImageTemplateLastRunStatusResponsePtrOutput() ImageTemplateLastRunStatusResponsePtrOutput {
	return i.ToImageTemplateLastRunStatusResponsePtrOutputWithContext(context.Background())
}

func (i *imageTemplateLastRunStatusResponsePtrType) ToImageTemplateLastRunStatusResponsePtrOutputWithContext(ctx context.Context) ImageTemplateLastRunStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateLastRunStatusResponsePtrOutput)
}

type ImageTemplateLastRunStatusResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplateLastRunStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateLastRunStatusResponse)(nil)).Elem()
}

func (o ImageTemplateLastRunStatusResponseOutput) ToImageTemplateLastRunStatusResponseOutput() ImageTemplateLastRunStatusResponseOutput {
	return o
}

func (o ImageTemplateLastRunStatusResponseOutput) ToImageTemplateLastRunStatusResponseOutputWithContext(ctx context.Context) ImageTemplateLastRunStatusResponseOutput {
	return o
}

func (o ImageTemplateLastRunStatusResponseOutput) ToImageTemplateLastRunStatusResponsePtrOutput() ImageTemplateLastRunStatusResponsePtrOutput {
	return o.ToImageTemplateLastRunStatusResponsePtrOutputWithContext(context.Background())
}

func (o ImageTemplateLastRunStatusResponseOutput) ToImageTemplateLastRunStatusResponsePtrOutputWithContext(ctx context.Context) ImageTemplateLastRunStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageTemplateLastRunStatusResponse) *ImageTemplateLastRunStatusResponse {
		return &v
	}).(ImageTemplateLastRunStatusResponsePtrOutput)
}

func (o ImageTemplateLastRunStatusResponseOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateLastRunStatusResponse) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateLastRunStatusResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateLastRunStatusResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateLastRunStatusResponseOutput) RunState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateLastRunStatusResponse) *string { return v.RunState }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateLastRunStatusResponseOutput) RunSubState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateLastRunStatusResponse) *string { return v.RunSubState }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateLastRunStatusResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateLastRunStatusResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type ImageTemplateLastRunStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageTemplateLastRunStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageTemplateLastRunStatusResponse)(nil)).Elem()
}

func (o ImageTemplateLastRunStatusResponsePtrOutput) ToImageTemplateLastRunStatusResponsePtrOutput() ImageTemplateLastRunStatusResponsePtrOutput {
	return o
}

func (o ImageTemplateLastRunStatusResponsePtrOutput) ToImageTemplateLastRunStatusResponsePtrOutputWithContext(ctx context.Context) ImageTemplateLastRunStatusResponsePtrOutput {
	return o
}

func (o ImageTemplateLastRunStatusResponsePtrOutput) Elem() ImageTemplateLastRunStatusResponseOutput {
	return o.ApplyT(func(v *ImageTemplateLastRunStatusResponse) ImageTemplateLastRunStatusResponse {
		if v != nil {
			return *v
		}
		var ret ImageTemplateLastRunStatusResponse
		return ret
	}).(ImageTemplateLastRunStatusResponseOutput)
}

func (o ImageTemplateLastRunStatusResponsePtrOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageTemplateLastRunStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndTime
	}).(pulumi.StringPtrOutput)
}

func (o ImageTemplateLastRunStatusResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageTemplateLastRunStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ImageTemplateLastRunStatusResponsePtrOutput) RunState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageTemplateLastRunStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.RunState
	}).(pulumi.StringPtrOutput)
}

func (o ImageTemplateLastRunStatusResponsePtrOutput) RunSubState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageTemplateLastRunStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.RunSubState
	}).(pulumi.StringPtrOutput)
}

func (o ImageTemplateLastRunStatusResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageTemplateLastRunStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

type ImageTemplateManagedImageDistributor struct {
	ArtifactTags  map[string]string `pulumi:"artifactTags"`
	ImageId       string            `pulumi:"imageId"`
	Location      string            `pulumi:"location"`
	RunOutputName string            `pulumi:"runOutputName"`
	Type          string            `pulumi:"type"`
}





type ImageTemplateManagedImageDistributorInput interface {
	pulumi.Input

	ToImageTemplateManagedImageDistributorOutput() ImageTemplateManagedImageDistributorOutput
	ToImageTemplateManagedImageDistributorOutputWithContext(context.Context) ImageTemplateManagedImageDistributorOutput
}

type ImageTemplateManagedImageDistributorArgs struct {
	ArtifactTags  pulumi.StringMapInput `pulumi:"artifactTags"`
	ImageId       pulumi.StringInput    `pulumi:"imageId"`
	Location      pulumi.StringInput    `pulumi:"location"`
	RunOutputName pulumi.StringInput    `pulumi:"runOutputName"`
	Type          pulumi.StringInput    `pulumi:"type"`
}

func (ImageTemplateManagedImageDistributorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateManagedImageDistributor)(nil)).Elem()
}

func (i ImageTemplateManagedImageDistributorArgs) ToImageTemplateManagedImageDistributorOutput() ImageTemplateManagedImageDistributorOutput {
	return i.ToImageTemplateManagedImageDistributorOutputWithContext(context.Background())
}

func (i ImageTemplateManagedImageDistributorArgs) ToImageTemplateManagedImageDistributorOutputWithContext(ctx context.Context) ImageTemplateManagedImageDistributorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateManagedImageDistributorOutput)
}

type ImageTemplateManagedImageDistributorOutput struct{ *pulumi.OutputState }

func (ImageTemplateManagedImageDistributorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateManagedImageDistributor)(nil)).Elem()
}

func (o ImageTemplateManagedImageDistributorOutput) ToImageTemplateManagedImageDistributorOutput() ImageTemplateManagedImageDistributorOutput {
	return o
}

func (o ImageTemplateManagedImageDistributorOutput) ToImageTemplateManagedImageDistributorOutputWithContext(ctx context.Context) ImageTemplateManagedImageDistributorOutput {
	return o
}

func (o ImageTemplateManagedImageDistributorOutput) ArtifactTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageDistributor) map[string]string { return v.ArtifactTags }).(pulumi.StringMapOutput)
}

func (o ImageTemplateManagedImageDistributorOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageDistributor) string { return v.ImageId }).(pulumi.StringOutput)
}

func (o ImageTemplateManagedImageDistributorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageDistributor) string { return v.Location }).(pulumi.StringOutput)
}

func (o ImageTemplateManagedImageDistributorOutput) RunOutputName() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageDistributor) string { return v.RunOutputName }).(pulumi.StringOutput)
}

func (o ImageTemplateManagedImageDistributorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageDistributor) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateManagedImageDistributorResponse struct {
	ArtifactTags  map[string]string `pulumi:"artifactTags"`
	ImageId       string            `pulumi:"imageId"`
	Location      string            `pulumi:"location"`
	RunOutputName string            `pulumi:"runOutputName"`
	Type          string            `pulumi:"type"`
}





type ImageTemplateManagedImageDistributorResponseInput interface {
	pulumi.Input

	ToImageTemplateManagedImageDistributorResponseOutput() ImageTemplateManagedImageDistributorResponseOutput
	ToImageTemplateManagedImageDistributorResponseOutputWithContext(context.Context) ImageTemplateManagedImageDistributorResponseOutput
}

type ImageTemplateManagedImageDistributorResponseArgs struct {
	ArtifactTags  pulumi.StringMapInput `pulumi:"artifactTags"`
	ImageId       pulumi.StringInput    `pulumi:"imageId"`
	Location      pulumi.StringInput    `pulumi:"location"`
	RunOutputName pulumi.StringInput    `pulumi:"runOutputName"`
	Type          pulumi.StringInput    `pulumi:"type"`
}

func (ImageTemplateManagedImageDistributorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateManagedImageDistributorResponse)(nil)).Elem()
}

func (i ImageTemplateManagedImageDistributorResponseArgs) ToImageTemplateManagedImageDistributorResponseOutput() ImageTemplateManagedImageDistributorResponseOutput {
	return i.ToImageTemplateManagedImageDistributorResponseOutputWithContext(context.Background())
}

func (i ImageTemplateManagedImageDistributorResponseArgs) ToImageTemplateManagedImageDistributorResponseOutputWithContext(ctx context.Context) ImageTemplateManagedImageDistributorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateManagedImageDistributorResponseOutput)
}

type ImageTemplateManagedImageDistributorResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplateManagedImageDistributorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateManagedImageDistributorResponse)(nil)).Elem()
}

func (o ImageTemplateManagedImageDistributorResponseOutput) ToImageTemplateManagedImageDistributorResponseOutput() ImageTemplateManagedImageDistributorResponseOutput {
	return o
}

func (o ImageTemplateManagedImageDistributorResponseOutput) ToImageTemplateManagedImageDistributorResponseOutputWithContext(ctx context.Context) ImageTemplateManagedImageDistributorResponseOutput {
	return o
}

func (o ImageTemplateManagedImageDistributorResponseOutput) ArtifactTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageDistributorResponse) map[string]string { return v.ArtifactTags }).(pulumi.StringMapOutput)
}

func (o ImageTemplateManagedImageDistributorResponseOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageDistributorResponse) string { return v.ImageId }).(pulumi.StringOutput)
}

func (o ImageTemplateManagedImageDistributorResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageDistributorResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o ImageTemplateManagedImageDistributorResponseOutput) RunOutputName() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageDistributorResponse) string { return v.RunOutputName }).(pulumi.StringOutput)
}

func (o ImageTemplateManagedImageDistributorResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageDistributorResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateManagedImageSource struct {
	ImageId string `pulumi:"imageId"`
	Type    string `pulumi:"type"`
}





type ImageTemplateManagedImageSourceInput interface {
	pulumi.Input

	ToImageTemplateManagedImageSourceOutput() ImageTemplateManagedImageSourceOutput
	ToImageTemplateManagedImageSourceOutputWithContext(context.Context) ImageTemplateManagedImageSourceOutput
}

type ImageTemplateManagedImageSourceArgs struct {
	ImageId pulumi.StringInput `pulumi:"imageId"`
	Type    pulumi.StringInput `pulumi:"type"`
}

func (ImageTemplateManagedImageSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateManagedImageSource)(nil)).Elem()
}

func (i ImageTemplateManagedImageSourceArgs) ToImageTemplateManagedImageSourceOutput() ImageTemplateManagedImageSourceOutput {
	return i.ToImageTemplateManagedImageSourceOutputWithContext(context.Background())
}

func (i ImageTemplateManagedImageSourceArgs) ToImageTemplateManagedImageSourceOutputWithContext(ctx context.Context) ImageTemplateManagedImageSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateManagedImageSourceOutput)
}

type ImageTemplateManagedImageSourceOutput struct{ *pulumi.OutputState }

func (ImageTemplateManagedImageSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateManagedImageSource)(nil)).Elem()
}

func (o ImageTemplateManagedImageSourceOutput) ToImageTemplateManagedImageSourceOutput() ImageTemplateManagedImageSourceOutput {
	return o
}

func (o ImageTemplateManagedImageSourceOutput) ToImageTemplateManagedImageSourceOutputWithContext(ctx context.Context) ImageTemplateManagedImageSourceOutput {
	return o
}

func (o ImageTemplateManagedImageSourceOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageSource) string { return v.ImageId }).(pulumi.StringOutput)
}

func (o ImageTemplateManagedImageSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageSource) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateManagedImageSourceResponse struct {
	ImageId string `pulumi:"imageId"`
	Type    string `pulumi:"type"`
}





type ImageTemplateManagedImageSourceResponseInput interface {
	pulumi.Input

	ToImageTemplateManagedImageSourceResponseOutput() ImageTemplateManagedImageSourceResponseOutput
	ToImageTemplateManagedImageSourceResponseOutputWithContext(context.Context) ImageTemplateManagedImageSourceResponseOutput
}

type ImageTemplateManagedImageSourceResponseArgs struct {
	ImageId pulumi.StringInput `pulumi:"imageId"`
	Type    pulumi.StringInput `pulumi:"type"`
}

func (ImageTemplateManagedImageSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateManagedImageSourceResponse)(nil)).Elem()
}

func (i ImageTemplateManagedImageSourceResponseArgs) ToImageTemplateManagedImageSourceResponseOutput() ImageTemplateManagedImageSourceResponseOutput {
	return i.ToImageTemplateManagedImageSourceResponseOutputWithContext(context.Background())
}

func (i ImageTemplateManagedImageSourceResponseArgs) ToImageTemplateManagedImageSourceResponseOutputWithContext(ctx context.Context) ImageTemplateManagedImageSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateManagedImageSourceResponseOutput)
}

type ImageTemplateManagedImageSourceResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplateManagedImageSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateManagedImageSourceResponse)(nil)).Elem()
}

func (o ImageTemplateManagedImageSourceResponseOutput) ToImageTemplateManagedImageSourceResponseOutput() ImageTemplateManagedImageSourceResponseOutput {
	return o
}

func (o ImageTemplateManagedImageSourceResponseOutput) ToImageTemplateManagedImageSourceResponseOutputWithContext(ctx context.Context) ImageTemplateManagedImageSourceResponseOutput {
	return o
}

func (o ImageTemplateManagedImageSourceResponseOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageSourceResponse) string { return v.ImageId }).(pulumi.StringOutput)
}

func (o ImageTemplateManagedImageSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplatePlatformImageSource struct {
	Offer     *string                    `pulumi:"offer"`
	PlanInfo  *PlatformImagePurchasePlan `pulumi:"planInfo"`
	Publisher *string                    `pulumi:"publisher"`
	Sku       *string                    `pulumi:"sku"`
	Type      string                     `pulumi:"type"`
	Version   *string                    `pulumi:"version"`
}





type ImageTemplatePlatformImageSourceInput interface {
	pulumi.Input

	ToImageTemplatePlatformImageSourceOutput() ImageTemplatePlatformImageSourceOutput
	ToImageTemplatePlatformImageSourceOutputWithContext(context.Context) ImageTemplatePlatformImageSourceOutput
}

type ImageTemplatePlatformImageSourceArgs struct {
	Offer     pulumi.StringPtrInput             `pulumi:"offer"`
	PlanInfo  PlatformImagePurchasePlanPtrInput `pulumi:"planInfo"`
	Publisher pulumi.StringPtrInput             `pulumi:"publisher"`
	Sku       pulumi.StringPtrInput             `pulumi:"sku"`
	Type      pulumi.StringInput                `pulumi:"type"`
	Version   pulumi.StringPtrInput             `pulumi:"version"`
}

func (ImageTemplatePlatformImageSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplatePlatformImageSource)(nil)).Elem()
}

func (i ImageTemplatePlatformImageSourceArgs) ToImageTemplatePlatformImageSourceOutput() ImageTemplatePlatformImageSourceOutput {
	return i.ToImageTemplatePlatformImageSourceOutputWithContext(context.Background())
}

func (i ImageTemplatePlatformImageSourceArgs) ToImageTemplatePlatformImageSourceOutputWithContext(ctx context.Context) ImageTemplatePlatformImageSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplatePlatformImageSourceOutput)
}

type ImageTemplatePlatformImageSourceOutput struct{ *pulumi.OutputState }

func (ImageTemplatePlatformImageSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplatePlatformImageSource)(nil)).Elem()
}

func (o ImageTemplatePlatformImageSourceOutput) ToImageTemplatePlatformImageSourceOutput() ImageTemplatePlatformImageSourceOutput {
	return o
}

func (o ImageTemplatePlatformImageSourceOutput) ToImageTemplatePlatformImageSourceOutputWithContext(ctx context.Context) ImageTemplatePlatformImageSourceOutput {
	return o
}

func (o ImageTemplatePlatformImageSourceOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSource) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o ImageTemplatePlatformImageSourceOutput) PlanInfo() PlatformImagePurchasePlanPtrOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSource) *PlatformImagePurchasePlan { return v.PlanInfo }).(PlatformImagePurchasePlanPtrOutput)
}

func (o ImageTemplatePlatformImageSourceOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSource) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ImageTemplatePlatformImageSourceOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSource) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o ImageTemplatePlatformImageSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSource) string { return v.Type }).(pulumi.StringOutput)
}

func (o ImageTemplatePlatformImageSourceOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSource) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ImageTemplatePlatformImageSourceResponse struct {
	Offer     *string                            `pulumi:"offer"`
	PlanInfo  *PlatformImagePurchasePlanResponse `pulumi:"planInfo"`
	Publisher *string                            `pulumi:"publisher"`
	Sku       *string                            `pulumi:"sku"`
	Type      string                             `pulumi:"type"`
	Version   *string                            `pulumi:"version"`
}





type ImageTemplatePlatformImageSourceResponseInput interface {
	pulumi.Input

	ToImageTemplatePlatformImageSourceResponseOutput() ImageTemplatePlatformImageSourceResponseOutput
	ToImageTemplatePlatformImageSourceResponseOutputWithContext(context.Context) ImageTemplatePlatformImageSourceResponseOutput
}

type ImageTemplatePlatformImageSourceResponseArgs struct {
	Offer     pulumi.StringPtrInput                     `pulumi:"offer"`
	PlanInfo  PlatformImagePurchasePlanResponsePtrInput `pulumi:"planInfo"`
	Publisher pulumi.StringPtrInput                     `pulumi:"publisher"`
	Sku       pulumi.StringPtrInput                     `pulumi:"sku"`
	Type      pulumi.StringInput                        `pulumi:"type"`
	Version   pulumi.StringPtrInput                     `pulumi:"version"`
}

func (ImageTemplatePlatformImageSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplatePlatformImageSourceResponse)(nil)).Elem()
}

func (i ImageTemplatePlatformImageSourceResponseArgs) ToImageTemplatePlatformImageSourceResponseOutput() ImageTemplatePlatformImageSourceResponseOutput {
	return i.ToImageTemplatePlatformImageSourceResponseOutputWithContext(context.Background())
}

func (i ImageTemplatePlatformImageSourceResponseArgs) ToImageTemplatePlatformImageSourceResponseOutputWithContext(ctx context.Context) ImageTemplatePlatformImageSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplatePlatformImageSourceResponseOutput)
}

type ImageTemplatePlatformImageSourceResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplatePlatformImageSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplatePlatformImageSourceResponse)(nil)).Elem()
}

func (o ImageTemplatePlatformImageSourceResponseOutput) ToImageTemplatePlatformImageSourceResponseOutput() ImageTemplatePlatformImageSourceResponseOutput {
	return o
}

func (o ImageTemplatePlatformImageSourceResponseOutput) ToImageTemplatePlatformImageSourceResponseOutputWithContext(ctx context.Context) ImageTemplatePlatformImageSourceResponseOutput {
	return o
}

func (o ImageTemplatePlatformImageSourceResponseOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSourceResponse) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o ImageTemplatePlatformImageSourceResponseOutput) PlanInfo() PlatformImagePurchasePlanResponsePtrOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSourceResponse) *PlatformImagePurchasePlanResponse { return v.PlanInfo }).(PlatformImagePurchasePlanResponsePtrOutput)
}

func (o ImageTemplatePlatformImageSourceResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSourceResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ImageTemplatePlatformImageSourceResponseOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSourceResponse) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o ImageTemplatePlatformImageSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ImageTemplatePlatformImageSourceResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSourceResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ImageTemplatePowerShellCustomizer struct {
	Inline         []string `pulumi:"inline"`
	Name           *string  `pulumi:"name"`
	RunAsSystem    *bool    `pulumi:"runAsSystem"`
	RunElevated    *bool    `pulumi:"runElevated"`
	ScriptUri      *string  `pulumi:"scriptUri"`
	Sha256Checksum *string  `pulumi:"sha256Checksum"`
	Type           string   `pulumi:"type"`
	ValidExitCodes []int    `pulumi:"validExitCodes"`
}





type ImageTemplatePowerShellCustomizerInput interface {
	pulumi.Input

	ToImageTemplatePowerShellCustomizerOutput() ImageTemplatePowerShellCustomizerOutput
	ToImageTemplatePowerShellCustomizerOutputWithContext(context.Context) ImageTemplatePowerShellCustomizerOutput
}

type ImageTemplatePowerShellCustomizerArgs struct {
	Inline         pulumi.StringArrayInput `pulumi:"inline"`
	Name           pulumi.StringPtrInput   `pulumi:"name"`
	RunAsSystem    pulumi.BoolPtrInput     `pulumi:"runAsSystem"`
	RunElevated    pulumi.BoolPtrInput     `pulumi:"runElevated"`
	ScriptUri      pulumi.StringPtrInput   `pulumi:"scriptUri"`
	Sha256Checksum pulumi.StringPtrInput   `pulumi:"sha256Checksum"`
	Type           pulumi.StringInput      `pulumi:"type"`
	ValidExitCodes pulumi.IntArrayInput    `pulumi:"validExitCodes"`
}

func (ImageTemplatePowerShellCustomizerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplatePowerShellCustomizer)(nil)).Elem()
}

func (i ImageTemplatePowerShellCustomizerArgs) ToImageTemplatePowerShellCustomizerOutput() ImageTemplatePowerShellCustomizerOutput {
	return i.ToImageTemplatePowerShellCustomizerOutputWithContext(context.Background())
}

func (i ImageTemplatePowerShellCustomizerArgs) ToImageTemplatePowerShellCustomizerOutputWithContext(ctx context.Context) ImageTemplatePowerShellCustomizerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplatePowerShellCustomizerOutput)
}

type ImageTemplatePowerShellCustomizerOutput struct{ *pulumi.OutputState }

func (ImageTemplatePowerShellCustomizerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplatePowerShellCustomizer)(nil)).Elem()
}

func (o ImageTemplatePowerShellCustomizerOutput) ToImageTemplatePowerShellCustomizerOutput() ImageTemplatePowerShellCustomizerOutput {
	return o
}

func (o ImageTemplatePowerShellCustomizerOutput) ToImageTemplatePowerShellCustomizerOutputWithContext(ctx context.Context) ImageTemplatePowerShellCustomizerOutput {
	return o
}

func (o ImageTemplatePowerShellCustomizerOutput) Inline() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ImageTemplatePowerShellCustomizer) []string { return v.Inline }).(pulumi.StringArrayOutput)
}

func (o ImageTemplatePowerShellCustomizerOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePowerShellCustomizer) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ImageTemplatePowerShellCustomizerOutput) RunAsSystem() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ImageTemplatePowerShellCustomizer) *bool { return v.RunAsSystem }).(pulumi.BoolPtrOutput)
}

func (o ImageTemplatePowerShellCustomizerOutput) RunElevated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ImageTemplatePowerShellCustomizer) *bool { return v.RunElevated }).(pulumi.BoolPtrOutput)
}

func (o ImageTemplatePowerShellCustomizerOutput) ScriptUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePowerShellCustomizer) *string { return v.ScriptUri }).(pulumi.StringPtrOutput)
}

func (o ImageTemplatePowerShellCustomizerOutput) Sha256Checksum() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePowerShellCustomizer) *string { return v.Sha256Checksum }).(pulumi.StringPtrOutput)
}

func (o ImageTemplatePowerShellCustomizerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplatePowerShellCustomizer) string { return v.Type }).(pulumi.StringOutput)
}

func (o ImageTemplatePowerShellCustomizerOutput) ValidExitCodes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v ImageTemplatePowerShellCustomizer) []int { return v.ValidExitCodes }).(pulumi.IntArrayOutput)
}

type ImageTemplatePowerShellCustomizerResponse struct {
	Inline         []string `pulumi:"inline"`
	Name           *string  `pulumi:"name"`
	RunAsSystem    *bool    `pulumi:"runAsSystem"`
	RunElevated    *bool    `pulumi:"runElevated"`
	ScriptUri      *string  `pulumi:"scriptUri"`
	Sha256Checksum *string  `pulumi:"sha256Checksum"`
	Type           string   `pulumi:"type"`
	ValidExitCodes []int    `pulumi:"validExitCodes"`
}





type ImageTemplatePowerShellCustomizerResponseInput interface {
	pulumi.Input

	ToImageTemplatePowerShellCustomizerResponseOutput() ImageTemplatePowerShellCustomizerResponseOutput
	ToImageTemplatePowerShellCustomizerResponseOutputWithContext(context.Context) ImageTemplatePowerShellCustomizerResponseOutput
}

type ImageTemplatePowerShellCustomizerResponseArgs struct {
	Inline         pulumi.StringArrayInput `pulumi:"inline"`
	Name           pulumi.StringPtrInput   `pulumi:"name"`
	RunAsSystem    pulumi.BoolPtrInput     `pulumi:"runAsSystem"`
	RunElevated    pulumi.BoolPtrInput     `pulumi:"runElevated"`
	ScriptUri      pulumi.StringPtrInput   `pulumi:"scriptUri"`
	Sha256Checksum pulumi.StringPtrInput   `pulumi:"sha256Checksum"`
	Type           pulumi.StringInput      `pulumi:"type"`
	ValidExitCodes pulumi.IntArrayInput    `pulumi:"validExitCodes"`
}

func (ImageTemplatePowerShellCustomizerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplatePowerShellCustomizerResponse)(nil)).Elem()
}

func (i ImageTemplatePowerShellCustomizerResponseArgs) ToImageTemplatePowerShellCustomizerResponseOutput() ImageTemplatePowerShellCustomizerResponseOutput {
	return i.ToImageTemplatePowerShellCustomizerResponseOutputWithContext(context.Background())
}

func (i ImageTemplatePowerShellCustomizerResponseArgs) ToImageTemplatePowerShellCustomizerResponseOutputWithContext(ctx context.Context) ImageTemplatePowerShellCustomizerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplatePowerShellCustomizerResponseOutput)
}

type ImageTemplatePowerShellCustomizerResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplatePowerShellCustomizerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplatePowerShellCustomizerResponse)(nil)).Elem()
}

func (o ImageTemplatePowerShellCustomizerResponseOutput) ToImageTemplatePowerShellCustomizerResponseOutput() ImageTemplatePowerShellCustomizerResponseOutput {
	return o
}

func (o ImageTemplatePowerShellCustomizerResponseOutput) ToImageTemplatePowerShellCustomizerResponseOutputWithContext(ctx context.Context) ImageTemplatePowerShellCustomizerResponseOutput {
	return o
}

func (o ImageTemplatePowerShellCustomizerResponseOutput) Inline() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ImageTemplatePowerShellCustomizerResponse) []string { return v.Inline }).(pulumi.StringArrayOutput)
}

func (o ImageTemplatePowerShellCustomizerResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePowerShellCustomizerResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ImageTemplatePowerShellCustomizerResponseOutput) RunAsSystem() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ImageTemplatePowerShellCustomizerResponse) *bool { return v.RunAsSystem }).(pulumi.BoolPtrOutput)
}

func (o ImageTemplatePowerShellCustomizerResponseOutput) RunElevated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ImageTemplatePowerShellCustomizerResponse) *bool { return v.RunElevated }).(pulumi.BoolPtrOutput)
}

func (o ImageTemplatePowerShellCustomizerResponseOutput) ScriptUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePowerShellCustomizerResponse) *string { return v.ScriptUri }).(pulumi.StringPtrOutput)
}

func (o ImageTemplatePowerShellCustomizerResponseOutput) Sha256Checksum() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePowerShellCustomizerResponse) *string { return v.Sha256Checksum }).(pulumi.StringPtrOutput)
}

func (o ImageTemplatePowerShellCustomizerResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplatePowerShellCustomizerResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ImageTemplatePowerShellCustomizerResponseOutput) ValidExitCodes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v ImageTemplatePowerShellCustomizerResponse) []int { return v.ValidExitCodes }).(pulumi.IntArrayOutput)
}

type ImageTemplateRestartCustomizer struct {
	Name                *string `pulumi:"name"`
	RestartCheckCommand *string `pulumi:"restartCheckCommand"`
	RestartCommand      *string `pulumi:"restartCommand"`
	RestartTimeout      *string `pulumi:"restartTimeout"`
	Type                string  `pulumi:"type"`
}





type ImageTemplateRestartCustomizerInput interface {
	pulumi.Input

	ToImageTemplateRestartCustomizerOutput() ImageTemplateRestartCustomizerOutput
	ToImageTemplateRestartCustomizerOutputWithContext(context.Context) ImageTemplateRestartCustomizerOutput
}

type ImageTemplateRestartCustomizerArgs struct {
	Name                pulumi.StringPtrInput `pulumi:"name"`
	RestartCheckCommand pulumi.StringPtrInput `pulumi:"restartCheckCommand"`
	RestartCommand      pulumi.StringPtrInput `pulumi:"restartCommand"`
	RestartTimeout      pulumi.StringPtrInput `pulumi:"restartTimeout"`
	Type                pulumi.StringInput    `pulumi:"type"`
}

func (ImageTemplateRestartCustomizerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateRestartCustomizer)(nil)).Elem()
}

func (i ImageTemplateRestartCustomizerArgs) ToImageTemplateRestartCustomizerOutput() ImageTemplateRestartCustomizerOutput {
	return i.ToImageTemplateRestartCustomizerOutputWithContext(context.Background())
}

func (i ImageTemplateRestartCustomizerArgs) ToImageTemplateRestartCustomizerOutputWithContext(ctx context.Context) ImageTemplateRestartCustomizerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateRestartCustomizerOutput)
}

type ImageTemplateRestartCustomizerOutput struct{ *pulumi.OutputState }

func (ImageTemplateRestartCustomizerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateRestartCustomizer)(nil)).Elem()
}

func (o ImageTemplateRestartCustomizerOutput) ToImageTemplateRestartCustomizerOutput() ImageTemplateRestartCustomizerOutput {
	return o
}

func (o ImageTemplateRestartCustomizerOutput) ToImageTemplateRestartCustomizerOutputWithContext(ctx context.Context) ImageTemplateRestartCustomizerOutput {
	return o
}

func (o ImageTemplateRestartCustomizerOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateRestartCustomizer) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateRestartCustomizerOutput) RestartCheckCommand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateRestartCustomizer) *string { return v.RestartCheckCommand }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateRestartCustomizerOutput) RestartCommand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateRestartCustomizer) *string { return v.RestartCommand }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateRestartCustomizerOutput) RestartTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateRestartCustomizer) *string { return v.RestartTimeout }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateRestartCustomizerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateRestartCustomizer) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateRestartCustomizerResponse struct {
	Name                *string `pulumi:"name"`
	RestartCheckCommand *string `pulumi:"restartCheckCommand"`
	RestartCommand      *string `pulumi:"restartCommand"`
	RestartTimeout      *string `pulumi:"restartTimeout"`
	Type                string  `pulumi:"type"`
}





type ImageTemplateRestartCustomizerResponseInput interface {
	pulumi.Input

	ToImageTemplateRestartCustomizerResponseOutput() ImageTemplateRestartCustomizerResponseOutput
	ToImageTemplateRestartCustomizerResponseOutputWithContext(context.Context) ImageTemplateRestartCustomizerResponseOutput
}

type ImageTemplateRestartCustomizerResponseArgs struct {
	Name                pulumi.StringPtrInput `pulumi:"name"`
	RestartCheckCommand pulumi.StringPtrInput `pulumi:"restartCheckCommand"`
	RestartCommand      pulumi.StringPtrInput `pulumi:"restartCommand"`
	RestartTimeout      pulumi.StringPtrInput `pulumi:"restartTimeout"`
	Type                pulumi.StringInput    `pulumi:"type"`
}

func (ImageTemplateRestartCustomizerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateRestartCustomizerResponse)(nil)).Elem()
}

func (i ImageTemplateRestartCustomizerResponseArgs) ToImageTemplateRestartCustomizerResponseOutput() ImageTemplateRestartCustomizerResponseOutput {
	return i.ToImageTemplateRestartCustomizerResponseOutputWithContext(context.Background())
}

func (i ImageTemplateRestartCustomizerResponseArgs) ToImageTemplateRestartCustomizerResponseOutputWithContext(ctx context.Context) ImageTemplateRestartCustomizerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateRestartCustomizerResponseOutput)
}

type ImageTemplateRestartCustomizerResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplateRestartCustomizerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateRestartCustomizerResponse)(nil)).Elem()
}

func (o ImageTemplateRestartCustomizerResponseOutput) ToImageTemplateRestartCustomizerResponseOutput() ImageTemplateRestartCustomizerResponseOutput {
	return o
}

func (o ImageTemplateRestartCustomizerResponseOutput) ToImageTemplateRestartCustomizerResponseOutputWithContext(ctx context.Context) ImageTemplateRestartCustomizerResponseOutput {
	return o
}

func (o ImageTemplateRestartCustomizerResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateRestartCustomizerResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateRestartCustomizerResponseOutput) RestartCheckCommand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateRestartCustomizerResponse) *string { return v.RestartCheckCommand }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateRestartCustomizerResponseOutput) RestartCommand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateRestartCustomizerResponse) *string { return v.RestartCommand }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateRestartCustomizerResponseOutput) RestartTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateRestartCustomizerResponse) *string { return v.RestartTimeout }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateRestartCustomizerResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateRestartCustomizerResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateSharedImageDistributor struct {
	ArtifactTags       map[string]string `pulumi:"artifactTags"`
	ExcludeFromLatest  *bool             `pulumi:"excludeFromLatest"`
	GalleryImageId     string            `pulumi:"galleryImageId"`
	ReplicationRegions []string          `pulumi:"replicationRegions"`
	RunOutputName      string            `pulumi:"runOutputName"`
	StorageAccountType *string           `pulumi:"storageAccountType"`
	Type               string            `pulumi:"type"`
}





type ImageTemplateSharedImageDistributorInput interface {
	pulumi.Input

	ToImageTemplateSharedImageDistributorOutput() ImageTemplateSharedImageDistributorOutput
	ToImageTemplateSharedImageDistributorOutputWithContext(context.Context) ImageTemplateSharedImageDistributorOutput
}

type ImageTemplateSharedImageDistributorArgs struct {
	ArtifactTags       pulumi.StringMapInput   `pulumi:"artifactTags"`
	ExcludeFromLatest  pulumi.BoolPtrInput     `pulumi:"excludeFromLatest"`
	GalleryImageId     pulumi.StringInput      `pulumi:"galleryImageId"`
	ReplicationRegions pulumi.StringArrayInput `pulumi:"replicationRegions"`
	RunOutputName      pulumi.StringInput      `pulumi:"runOutputName"`
	StorageAccountType pulumi.StringPtrInput   `pulumi:"storageAccountType"`
	Type               pulumi.StringInput      `pulumi:"type"`
}

func (ImageTemplateSharedImageDistributorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateSharedImageDistributor)(nil)).Elem()
}

func (i ImageTemplateSharedImageDistributorArgs) ToImageTemplateSharedImageDistributorOutput() ImageTemplateSharedImageDistributorOutput {
	return i.ToImageTemplateSharedImageDistributorOutputWithContext(context.Background())
}

func (i ImageTemplateSharedImageDistributorArgs) ToImageTemplateSharedImageDistributorOutputWithContext(ctx context.Context) ImageTemplateSharedImageDistributorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateSharedImageDistributorOutput)
}

type ImageTemplateSharedImageDistributorOutput struct{ *pulumi.OutputState }

func (ImageTemplateSharedImageDistributorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateSharedImageDistributor)(nil)).Elem()
}

func (o ImageTemplateSharedImageDistributorOutput) ToImageTemplateSharedImageDistributorOutput() ImageTemplateSharedImageDistributorOutput {
	return o
}

func (o ImageTemplateSharedImageDistributorOutput) ToImageTemplateSharedImageDistributorOutputWithContext(ctx context.Context) ImageTemplateSharedImageDistributorOutput {
	return o
}

func (o ImageTemplateSharedImageDistributorOutput) ArtifactTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributor) map[string]string { return v.ArtifactTags }).(pulumi.StringMapOutput)
}

func (o ImageTemplateSharedImageDistributorOutput) ExcludeFromLatest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributor) *bool { return v.ExcludeFromLatest }).(pulumi.BoolPtrOutput)
}

func (o ImageTemplateSharedImageDistributorOutput) GalleryImageId() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributor) string { return v.GalleryImageId }).(pulumi.StringOutput)
}

func (o ImageTemplateSharedImageDistributorOutput) ReplicationRegions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributor) []string { return v.ReplicationRegions }).(pulumi.StringArrayOutput)
}

func (o ImageTemplateSharedImageDistributorOutput) RunOutputName() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributor) string { return v.RunOutputName }).(pulumi.StringOutput)
}

func (o ImageTemplateSharedImageDistributorOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributor) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateSharedImageDistributorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributor) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateSharedImageDistributorResponse struct {
	ArtifactTags       map[string]string `pulumi:"artifactTags"`
	ExcludeFromLatest  *bool             `pulumi:"excludeFromLatest"`
	GalleryImageId     string            `pulumi:"galleryImageId"`
	ReplicationRegions []string          `pulumi:"replicationRegions"`
	RunOutputName      string            `pulumi:"runOutputName"`
	StorageAccountType *string           `pulumi:"storageAccountType"`
	Type               string            `pulumi:"type"`
}





type ImageTemplateSharedImageDistributorResponseInput interface {
	pulumi.Input

	ToImageTemplateSharedImageDistributorResponseOutput() ImageTemplateSharedImageDistributorResponseOutput
	ToImageTemplateSharedImageDistributorResponseOutputWithContext(context.Context) ImageTemplateSharedImageDistributorResponseOutput
}

type ImageTemplateSharedImageDistributorResponseArgs struct {
	ArtifactTags       pulumi.StringMapInput   `pulumi:"artifactTags"`
	ExcludeFromLatest  pulumi.BoolPtrInput     `pulumi:"excludeFromLatest"`
	GalleryImageId     pulumi.StringInput      `pulumi:"galleryImageId"`
	ReplicationRegions pulumi.StringArrayInput `pulumi:"replicationRegions"`
	RunOutputName      pulumi.StringInput      `pulumi:"runOutputName"`
	StorageAccountType pulumi.StringPtrInput   `pulumi:"storageAccountType"`
	Type               pulumi.StringInput      `pulumi:"type"`
}

func (ImageTemplateSharedImageDistributorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateSharedImageDistributorResponse)(nil)).Elem()
}

func (i ImageTemplateSharedImageDistributorResponseArgs) ToImageTemplateSharedImageDistributorResponseOutput() ImageTemplateSharedImageDistributorResponseOutput {
	return i.ToImageTemplateSharedImageDistributorResponseOutputWithContext(context.Background())
}

func (i ImageTemplateSharedImageDistributorResponseArgs) ToImageTemplateSharedImageDistributorResponseOutputWithContext(ctx context.Context) ImageTemplateSharedImageDistributorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateSharedImageDistributorResponseOutput)
}

type ImageTemplateSharedImageDistributorResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplateSharedImageDistributorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateSharedImageDistributorResponse)(nil)).Elem()
}

func (o ImageTemplateSharedImageDistributorResponseOutput) ToImageTemplateSharedImageDistributorResponseOutput() ImageTemplateSharedImageDistributorResponseOutput {
	return o
}

func (o ImageTemplateSharedImageDistributorResponseOutput) ToImageTemplateSharedImageDistributorResponseOutputWithContext(ctx context.Context) ImageTemplateSharedImageDistributorResponseOutput {
	return o
}

func (o ImageTemplateSharedImageDistributorResponseOutput) ArtifactTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributorResponse) map[string]string { return v.ArtifactTags }).(pulumi.StringMapOutput)
}

func (o ImageTemplateSharedImageDistributorResponseOutput) ExcludeFromLatest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributorResponse) *bool { return v.ExcludeFromLatest }).(pulumi.BoolPtrOutput)
}

func (o ImageTemplateSharedImageDistributorResponseOutput) GalleryImageId() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributorResponse) string { return v.GalleryImageId }).(pulumi.StringOutput)
}

func (o ImageTemplateSharedImageDistributorResponseOutput) ReplicationRegions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributorResponse) []string { return v.ReplicationRegions }).(pulumi.StringArrayOutput)
}

func (o ImageTemplateSharedImageDistributorResponseOutput) RunOutputName() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributorResponse) string { return v.RunOutputName }).(pulumi.StringOutput)
}

func (o ImageTemplateSharedImageDistributorResponseOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributorResponse) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateSharedImageDistributorResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributorResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateSharedImageVersionSource struct {
	ImageVersionId string `pulumi:"imageVersionId"`
	Type           string `pulumi:"type"`
}





type ImageTemplateSharedImageVersionSourceInput interface {
	pulumi.Input

	ToImageTemplateSharedImageVersionSourceOutput() ImageTemplateSharedImageVersionSourceOutput
	ToImageTemplateSharedImageVersionSourceOutputWithContext(context.Context) ImageTemplateSharedImageVersionSourceOutput
}

type ImageTemplateSharedImageVersionSourceArgs struct {
	ImageVersionId pulumi.StringInput `pulumi:"imageVersionId"`
	Type           pulumi.StringInput `pulumi:"type"`
}

func (ImageTemplateSharedImageVersionSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateSharedImageVersionSource)(nil)).Elem()
}

func (i ImageTemplateSharedImageVersionSourceArgs) ToImageTemplateSharedImageVersionSourceOutput() ImageTemplateSharedImageVersionSourceOutput {
	return i.ToImageTemplateSharedImageVersionSourceOutputWithContext(context.Background())
}

func (i ImageTemplateSharedImageVersionSourceArgs) ToImageTemplateSharedImageVersionSourceOutputWithContext(ctx context.Context) ImageTemplateSharedImageVersionSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateSharedImageVersionSourceOutput)
}

type ImageTemplateSharedImageVersionSourceOutput struct{ *pulumi.OutputState }

func (ImageTemplateSharedImageVersionSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateSharedImageVersionSource)(nil)).Elem()
}

func (o ImageTemplateSharedImageVersionSourceOutput) ToImageTemplateSharedImageVersionSourceOutput() ImageTemplateSharedImageVersionSourceOutput {
	return o
}

func (o ImageTemplateSharedImageVersionSourceOutput) ToImageTemplateSharedImageVersionSourceOutputWithContext(ctx context.Context) ImageTemplateSharedImageVersionSourceOutput {
	return o
}

func (o ImageTemplateSharedImageVersionSourceOutput) ImageVersionId() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageVersionSource) string { return v.ImageVersionId }).(pulumi.StringOutput)
}

func (o ImageTemplateSharedImageVersionSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageVersionSource) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateSharedImageVersionSourceResponse struct {
	ImageVersionId string `pulumi:"imageVersionId"`
	Type           string `pulumi:"type"`
}





type ImageTemplateSharedImageVersionSourceResponseInput interface {
	pulumi.Input

	ToImageTemplateSharedImageVersionSourceResponseOutput() ImageTemplateSharedImageVersionSourceResponseOutput
	ToImageTemplateSharedImageVersionSourceResponseOutputWithContext(context.Context) ImageTemplateSharedImageVersionSourceResponseOutput
}

type ImageTemplateSharedImageVersionSourceResponseArgs struct {
	ImageVersionId pulumi.StringInput `pulumi:"imageVersionId"`
	Type           pulumi.StringInput `pulumi:"type"`
}

func (ImageTemplateSharedImageVersionSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateSharedImageVersionSourceResponse)(nil)).Elem()
}

func (i ImageTemplateSharedImageVersionSourceResponseArgs) ToImageTemplateSharedImageVersionSourceResponseOutput() ImageTemplateSharedImageVersionSourceResponseOutput {
	return i.ToImageTemplateSharedImageVersionSourceResponseOutputWithContext(context.Background())
}

func (i ImageTemplateSharedImageVersionSourceResponseArgs) ToImageTemplateSharedImageVersionSourceResponseOutputWithContext(ctx context.Context) ImageTemplateSharedImageVersionSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateSharedImageVersionSourceResponseOutput)
}

type ImageTemplateSharedImageVersionSourceResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplateSharedImageVersionSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateSharedImageVersionSourceResponse)(nil)).Elem()
}

func (o ImageTemplateSharedImageVersionSourceResponseOutput) ToImageTemplateSharedImageVersionSourceResponseOutput() ImageTemplateSharedImageVersionSourceResponseOutput {
	return o
}

func (o ImageTemplateSharedImageVersionSourceResponseOutput) ToImageTemplateSharedImageVersionSourceResponseOutputWithContext(ctx context.Context) ImageTemplateSharedImageVersionSourceResponseOutput {
	return o
}

func (o ImageTemplateSharedImageVersionSourceResponseOutput) ImageVersionId() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageVersionSourceResponse) string { return v.ImageVersionId }).(pulumi.StringOutput)
}

func (o ImageTemplateSharedImageVersionSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageVersionSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateShellCustomizer struct {
	Inline         []string `pulumi:"inline"`
	Name           *string  `pulumi:"name"`
	ScriptUri      *string  `pulumi:"scriptUri"`
	Sha256Checksum *string  `pulumi:"sha256Checksum"`
	Type           string   `pulumi:"type"`
}





type ImageTemplateShellCustomizerInput interface {
	pulumi.Input

	ToImageTemplateShellCustomizerOutput() ImageTemplateShellCustomizerOutput
	ToImageTemplateShellCustomizerOutputWithContext(context.Context) ImageTemplateShellCustomizerOutput
}

type ImageTemplateShellCustomizerArgs struct {
	Inline         pulumi.StringArrayInput `pulumi:"inline"`
	Name           pulumi.StringPtrInput   `pulumi:"name"`
	ScriptUri      pulumi.StringPtrInput   `pulumi:"scriptUri"`
	Sha256Checksum pulumi.StringPtrInput   `pulumi:"sha256Checksum"`
	Type           pulumi.StringInput      `pulumi:"type"`
}

func (ImageTemplateShellCustomizerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateShellCustomizer)(nil)).Elem()
}

func (i ImageTemplateShellCustomizerArgs) ToImageTemplateShellCustomizerOutput() ImageTemplateShellCustomizerOutput {
	return i.ToImageTemplateShellCustomizerOutputWithContext(context.Background())
}

func (i ImageTemplateShellCustomizerArgs) ToImageTemplateShellCustomizerOutputWithContext(ctx context.Context) ImageTemplateShellCustomizerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateShellCustomizerOutput)
}

type ImageTemplateShellCustomizerOutput struct{ *pulumi.OutputState }

func (ImageTemplateShellCustomizerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateShellCustomizer)(nil)).Elem()
}

func (o ImageTemplateShellCustomizerOutput) ToImageTemplateShellCustomizerOutput() ImageTemplateShellCustomizerOutput {
	return o
}

func (o ImageTemplateShellCustomizerOutput) ToImageTemplateShellCustomizerOutputWithContext(ctx context.Context) ImageTemplateShellCustomizerOutput {
	return o
}

func (o ImageTemplateShellCustomizerOutput) Inline() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ImageTemplateShellCustomizer) []string { return v.Inline }).(pulumi.StringArrayOutput)
}

func (o ImageTemplateShellCustomizerOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateShellCustomizer) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateShellCustomizerOutput) ScriptUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateShellCustomizer) *string { return v.ScriptUri }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateShellCustomizerOutput) Sha256Checksum() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateShellCustomizer) *string { return v.Sha256Checksum }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateShellCustomizerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateShellCustomizer) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateShellCustomizerResponse struct {
	Inline         []string `pulumi:"inline"`
	Name           *string  `pulumi:"name"`
	ScriptUri      *string  `pulumi:"scriptUri"`
	Sha256Checksum *string  `pulumi:"sha256Checksum"`
	Type           string   `pulumi:"type"`
}





type ImageTemplateShellCustomizerResponseInput interface {
	pulumi.Input

	ToImageTemplateShellCustomizerResponseOutput() ImageTemplateShellCustomizerResponseOutput
	ToImageTemplateShellCustomizerResponseOutputWithContext(context.Context) ImageTemplateShellCustomizerResponseOutput
}

type ImageTemplateShellCustomizerResponseArgs struct {
	Inline         pulumi.StringArrayInput `pulumi:"inline"`
	Name           pulumi.StringPtrInput   `pulumi:"name"`
	ScriptUri      pulumi.StringPtrInput   `pulumi:"scriptUri"`
	Sha256Checksum pulumi.StringPtrInput   `pulumi:"sha256Checksum"`
	Type           pulumi.StringInput      `pulumi:"type"`
}

func (ImageTemplateShellCustomizerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateShellCustomizerResponse)(nil)).Elem()
}

func (i ImageTemplateShellCustomizerResponseArgs) ToImageTemplateShellCustomizerResponseOutput() ImageTemplateShellCustomizerResponseOutput {
	return i.ToImageTemplateShellCustomizerResponseOutputWithContext(context.Background())
}

func (i ImageTemplateShellCustomizerResponseArgs) ToImageTemplateShellCustomizerResponseOutputWithContext(ctx context.Context) ImageTemplateShellCustomizerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateShellCustomizerResponseOutput)
}

type ImageTemplateShellCustomizerResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplateShellCustomizerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateShellCustomizerResponse)(nil)).Elem()
}

func (o ImageTemplateShellCustomizerResponseOutput) ToImageTemplateShellCustomizerResponseOutput() ImageTemplateShellCustomizerResponseOutput {
	return o
}

func (o ImageTemplateShellCustomizerResponseOutput) ToImageTemplateShellCustomizerResponseOutputWithContext(ctx context.Context) ImageTemplateShellCustomizerResponseOutput {
	return o
}

func (o ImageTemplateShellCustomizerResponseOutput) Inline() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ImageTemplateShellCustomizerResponse) []string { return v.Inline }).(pulumi.StringArrayOutput)
}

func (o ImageTemplateShellCustomizerResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateShellCustomizerResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateShellCustomizerResponseOutput) ScriptUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateShellCustomizerResponse) *string { return v.ScriptUri }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateShellCustomizerResponseOutput) Sha256Checksum() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateShellCustomizerResponse) *string { return v.Sha256Checksum }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateShellCustomizerResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateShellCustomizerResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateVhdDistributor struct {
	ArtifactTags  map[string]string `pulumi:"artifactTags"`
	RunOutputName string            `pulumi:"runOutputName"`
	Type          string            `pulumi:"type"`
}





type ImageTemplateVhdDistributorInput interface {
	pulumi.Input

	ToImageTemplateVhdDistributorOutput() ImageTemplateVhdDistributorOutput
	ToImageTemplateVhdDistributorOutputWithContext(context.Context) ImageTemplateVhdDistributorOutput
}

type ImageTemplateVhdDistributorArgs struct {
	ArtifactTags  pulumi.StringMapInput `pulumi:"artifactTags"`
	RunOutputName pulumi.StringInput    `pulumi:"runOutputName"`
	Type          pulumi.StringInput    `pulumi:"type"`
}

func (ImageTemplateVhdDistributorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateVhdDistributor)(nil)).Elem()
}

func (i ImageTemplateVhdDistributorArgs) ToImageTemplateVhdDistributorOutput() ImageTemplateVhdDistributorOutput {
	return i.ToImageTemplateVhdDistributorOutputWithContext(context.Background())
}

func (i ImageTemplateVhdDistributorArgs) ToImageTemplateVhdDistributorOutputWithContext(ctx context.Context) ImageTemplateVhdDistributorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateVhdDistributorOutput)
}

type ImageTemplateVhdDistributorOutput struct{ *pulumi.OutputState }

func (ImageTemplateVhdDistributorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateVhdDistributor)(nil)).Elem()
}

func (o ImageTemplateVhdDistributorOutput) ToImageTemplateVhdDistributorOutput() ImageTemplateVhdDistributorOutput {
	return o
}

func (o ImageTemplateVhdDistributorOutput) ToImageTemplateVhdDistributorOutputWithContext(ctx context.Context) ImageTemplateVhdDistributorOutput {
	return o
}

func (o ImageTemplateVhdDistributorOutput) ArtifactTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ImageTemplateVhdDistributor) map[string]string { return v.ArtifactTags }).(pulumi.StringMapOutput)
}

func (o ImageTemplateVhdDistributorOutput) RunOutputName() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateVhdDistributor) string { return v.RunOutputName }).(pulumi.StringOutput)
}

func (o ImageTemplateVhdDistributorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateVhdDistributor) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateVhdDistributorResponse struct {
	ArtifactTags  map[string]string `pulumi:"artifactTags"`
	RunOutputName string            `pulumi:"runOutputName"`
	Type          string            `pulumi:"type"`
}





type ImageTemplateVhdDistributorResponseInput interface {
	pulumi.Input

	ToImageTemplateVhdDistributorResponseOutput() ImageTemplateVhdDistributorResponseOutput
	ToImageTemplateVhdDistributorResponseOutputWithContext(context.Context) ImageTemplateVhdDistributorResponseOutput
}

type ImageTemplateVhdDistributorResponseArgs struct {
	ArtifactTags  pulumi.StringMapInput `pulumi:"artifactTags"`
	RunOutputName pulumi.StringInput    `pulumi:"runOutputName"`
	Type          pulumi.StringInput    `pulumi:"type"`
}

func (ImageTemplateVhdDistributorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateVhdDistributorResponse)(nil)).Elem()
}

func (i ImageTemplateVhdDistributorResponseArgs) ToImageTemplateVhdDistributorResponseOutput() ImageTemplateVhdDistributorResponseOutput {
	return i.ToImageTemplateVhdDistributorResponseOutputWithContext(context.Background())
}

func (i ImageTemplateVhdDistributorResponseArgs) ToImageTemplateVhdDistributorResponseOutputWithContext(ctx context.Context) ImageTemplateVhdDistributorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateVhdDistributorResponseOutput)
}

type ImageTemplateVhdDistributorResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplateVhdDistributorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateVhdDistributorResponse)(nil)).Elem()
}

func (o ImageTemplateVhdDistributorResponseOutput) ToImageTemplateVhdDistributorResponseOutput() ImageTemplateVhdDistributorResponseOutput {
	return o
}

func (o ImageTemplateVhdDistributorResponseOutput) ToImageTemplateVhdDistributorResponseOutputWithContext(ctx context.Context) ImageTemplateVhdDistributorResponseOutput {
	return o
}

func (o ImageTemplateVhdDistributorResponseOutput) ArtifactTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ImageTemplateVhdDistributorResponse) map[string]string { return v.ArtifactTags }).(pulumi.StringMapOutput)
}

func (o ImageTemplateVhdDistributorResponseOutput) RunOutputName() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateVhdDistributorResponse) string { return v.RunOutputName }).(pulumi.StringOutput)
}

func (o ImageTemplateVhdDistributorResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateVhdDistributorResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateVmProfile struct {
	OsDiskSizeGB *int                  `pulumi:"osDiskSizeGB"`
	VmSize       *string               `pulumi:"vmSize"`
	VnetConfig   *VirtualNetworkConfig `pulumi:"vnetConfig"`
}





type ImageTemplateVmProfileInput interface {
	pulumi.Input

	ToImageTemplateVmProfileOutput() ImageTemplateVmProfileOutput
	ToImageTemplateVmProfileOutputWithContext(context.Context) ImageTemplateVmProfileOutput
}

type ImageTemplateVmProfileArgs struct {
	OsDiskSizeGB pulumi.IntPtrInput           `pulumi:"osDiskSizeGB"`
	VmSize       pulumi.StringPtrInput        `pulumi:"vmSize"`
	VnetConfig   VirtualNetworkConfigPtrInput `pulumi:"vnetConfig"`
}

func (ImageTemplateVmProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateVmProfile)(nil)).Elem()
}

func (i ImageTemplateVmProfileArgs) ToImageTemplateVmProfileOutput() ImageTemplateVmProfileOutput {
	return i.ToImageTemplateVmProfileOutputWithContext(context.Background())
}

func (i ImageTemplateVmProfileArgs) ToImageTemplateVmProfileOutputWithContext(ctx context.Context) ImageTemplateVmProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateVmProfileOutput)
}

func (i ImageTemplateVmProfileArgs) ToImageTemplateVmProfilePtrOutput() ImageTemplateVmProfilePtrOutput {
	return i.ToImageTemplateVmProfilePtrOutputWithContext(context.Background())
}

func (i ImageTemplateVmProfileArgs) ToImageTemplateVmProfilePtrOutputWithContext(ctx context.Context) ImageTemplateVmProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateVmProfileOutput).ToImageTemplateVmProfilePtrOutputWithContext(ctx)
}









type ImageTemplateVmProfilePtrInput interface {
	pulumi.Input

	ToImageTemplateVmProfilePtrOutput() ImageTemplateVmProfilePtrOutput
	ToImageTemplateVmProfilePtrOutputWithContext(context.Context) ImageTemplateVmProfilePtrOutput
}

type imageTemplateVmProfilePtrType ImageTemplateVmProfileArgs

func ImageTemplateVmProfilePtr(v *ImageTemplateVmProfileArgs) ImageTemplateVmProfilePtrInput {
	return (*imageTemplateVmProfilePtrType)(v)
}

func (*imageTemplateVmProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageTemplateVmProfile)(nil)).Elem()
}

func (i *imageTemplateVmProfilePtrType) ToImageTemplateVmProfilePtrOutput() ImageTemplateVmProfilePtrOutput {
	return i.ToImageTemplateVmProfilePtrOutputWithContext(context.Background())
}

func (i *imageTemplateVmProfilePtrType) ToImageTemplateVmProfilePtrOutputWithContext(ctx context.Context) ImageTemplateVmProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateVmProfilePtrOutput)
}

type ImageTemplateVmProfileOutput struct{ *pulumi.OutputState }

func (ImageTemplateVmProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateVmProfile)(nil)).Elem()
}

func (o ImageTemplateVmProfileOutput) ToImageTemplateVmProfileOutput() ImageTemplateVmProfileOutput {
	return o
}

func (o ImageTemplateVmProfileOutput) ToImageTemplateVmProfileOutputWithContext(ctx context.Context) ImageTemplateVmProfileOutput {
	return o
}

func (o ImageTemplateVmProfileOutput) ToImageTemplateVmProfilePtrOutput() ImageTemplateVmProfilePtrOutput {
	return o.ToImageTemplateVmProfilePtrOutputWithContext(context.Background())
}

func (o ImageTemplateVmProfileOutput) ToImageTemplateVmProfilePtrOutputWithContext(ctx context.Context) ImageTemplateVmProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageTemplateVmProfile) *ImageTemplateVmProfile {
		return &v
	}).(ImageTemplateVmProfilePtrOutput)
}

func (o ImageTemplateVmProfileOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ImageTemplateVmProfile) *int { return v.OsDiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o ImageTemplateVmProfileOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateVmProfile) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateVmProfileOutput) VnetConfig() VirtualNetworkConfigPtrOutput {
	return o.ApplyT(func(v ImageTemplateVmProfile) *VirtualNetworkConfig { return v.VnetConfig }).(VirtualNetworkConfigPtrOutput)
}

type ImageTemplateVmProfilePtrOutput struct{ *pulumi.OutputState }

func (ImageTemplateVmProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageTemplateVmProfile)(nil)).Elem()
}

func (o ImageTemplateVmProfilePtrOutput) ToImageTemplateVmProfilePtrOutput() ImageTemplateVmProfilePtrOutput {
	return o
}

func (o ImageTemplateVmProfilePtrOutput) ToImageTemplateVmProfilePtrOutputWithContext(ctx context.Context) ImageTemplateVmProfilePtrOutput {
	return o
}

func (o ImageTemplateVmProfilePtrOutput) Elem() ImageTemplateVmProfileOutput {
	return o.ApplyT(func(v *ImageTemplateVmProfile) ImageTemplateVmProfile {
		if v != nil {
			return *v
		}
		var ret ImageTemplateVmProfile
		return ret
	}).(ImageTemplateVmProfileOutput)
}

func (o ImageTemplateVmProfilePtrOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ImageTemplateVmProfile) *int {
		if v == nil {
			return nil
		}
		return v.OsDiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o ImageTemplateVmProfilePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageTemplateVmProfile) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

func (o ImageTemplateVmProfilePtrOutput) VnetConfig() VirtualNetworkConfigPtrOutput {
	return o.ApplyT(func(v *ImageTemplateVmProfile) *VirtualNetworkConfig {
		if v == nil {
			return nil
		}
		return v.VnetConfig
	}).(VirtualNetworkConfigPtrOutput)
}

type ImageTemplateVmProfileResponse struct {
	OsDiskSizeGB *int                          `pulumi:"osDiskSizeGB"`
	VmSize       *string                       `pulumi:"vmSize"`
	VnetConfig   *VirtualNetworkConfigResponse `pulumi:"vnetConfig"`
}





type ImageTemplateVmProfileResponseInput interface {
	pulumi.Input

	ToImageTemplateVmProfileResponseOutput() ImageTemplateVmProfileResponseOutput
	ToImageTemplateVmProfileResponseOutputWithContext(context.Context) ImageTemplateVmProfileResponseOutput
}

type ImageTemplateVmProfileResponseArgs struct {
	OsDiskSizeGB pulumi.IntPtrInput                   `pulumi:"osDiskSizeGB"`
	VmSize       pulumi.StringPtrInput                `pulumi:"vmSize"`
	VnetConfig   VirtualNetworkConfigResponsePtrInput `pulumi:"vnetConfig"`
}

func (ImageTemplateVmProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateVmProfileResponse)(nil)).Elem()
}

func (i ImageTemplateVmProfileResponseArgs) ToImageTemplateVmProfileResponseOutput() ImageTemplateVmProfileResponseOutput {
	return i.ToImageTemplateVmProfileResponseOutputWithContext(context.Background())
}

func (i ImageTemplateVmProfileResponseArgs) ToImageTemplateVmProfileResponseOutputWithContext(ctx context.Context) ImageTemplateVmProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateVmProfileResponseOutput)
}

func (i ImageTemplateVmProfileResponseArgs) ToImageTemplateVmProfileResponsePtrOutput() ImageTemplateVmProfileResponsePtrOutput {
	return i.ToImageTemplateVmProfileResponsePtrOutputWithContext(context.Background())
}

func (i ImageTemplateVmProfileResponseArgs) ToImageTemplateVmProfileResponsePtrOutputWithContext(ctx context.Context) ImageTemplateVmProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateVmProfileResponseOutput).ToImageTemplateVmProfileResponsePtrOutputWithContext(ctx)
}









type ImageTemplateVmProfileResponsePtrInput interface {
	pulumi.Input

	ToImageTemplateVmProfileResponsePtrOutput() ImageTemplateVmProfileResponsePtrOutput
	ToImageTemplateVmProfileResponsePtrOutputWithContext(context.Context) ImageTemplateVmProfileResponsePtrOutput
}

type imageTemplateVmProfileResponsePtrType ImageTemplateVmProfileResponseArgs

func ImageTemplateVmProfileResponsePtr(v *ImageTemplateVmProfileResponseArgs) ImageTemplateVmProfileResponsePtrInput {
	return (*imageTemplateVmProfileResponsePtrType)(v)
}

func (*imageTemplateVmProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageTemplateVmProfileResponse)(nil)).Elem()
}

func (i *imageTemplateVmProfileResponsePtrType) ToImageTemplateVmProfileResponsePtrOutput() ImageTemplateVmProfileResponsePtrOutput {
	return i.ToImageTemplateVmProfileResponsePtrOutputWithContext(context.Background())
}

func (i *imageTemplateVmProfileResponsePtrType) ToImageTemplateVmProfileResponsePtrOutputWithContext(ctx context.Context) ImageTemplateVmProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateVmProfileResponsePtrOutput)
}

type ImageTemplateVmProfileResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplateVmProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateVmProfileResponse)(nil)).Elem()
}

func (o ImageTemplateVmProfileResponseOutput) ToImageTemplateVmProfileResponseOutput() ImageTemplateVmProfileResponseOutput {
	return o
}

func (o ImageTemplateVmProfileResponseOutput) ToImageTemplateVmProfileResponseOutputWithContext(ctx context.Context) ImageTemplateVmProfileResponseOutput {
	return o
}

func (o ImageTemplateVmProfileResponseOutput) ToImageTemplateVmProfileResponsePtrOutput() ImageTemplateVmProfileResponsePtrOutput {
	return o.ToImageTemplateVmProfileResponsePtrOutputWithContext(context.Background())
}

func (o ImageTemplateVmProfileResponseOutput) ToImageTemplateVmProfileResponsePtrOutputWithContext(ctx context.Context) ImageTemplateVmProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageTemplateVmProfileResponse) *ImageTemplateVmProfileResponse {
		return &v
	}).(ImageTemplateVmProfileResponsePtrOutput)
}

func (o ImageTemplateVmProfileResponseOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ImageTemplateVmProfileResponse) *int { return v.OsDiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o ImageTemplateVmProfileResponseOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateVmProfileResponse) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateVmProfileResponseOutput) VnetConfig() VirtualNetworkConfigResponsePtrOutput {
	return o.ApplyT(func(v ImageTemplateVmProfileResponse) *VirtualNetworkConfigResponse { return v.VnetConfig }).(VirtualNetworkConfigResponsePtrOutput)
}

type ImageTemplateVmProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageTemplateVmProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageTemplateVmProfileResponse)(nil)).Elem()
}

func (o ImageTemplateVmProfileResponsePtrOutput) ToImageTemplateVmProfileResponsePtrOutput() ImageTemplateVmProfileResponsePtrOutput {
	return o
}

func (o ImageTemplateVmProfileResponsePtrOutput) ToImageTemplateVmProfileResponsePtrOutputWithContext(ctx context.Context) ImageTemplateVmProfileResponsePtrOutput {
	return o
}

func (o ImageTemplateVmProfileResponsePtrOutput) Elem() ImageTemplateVmProfileResponseOutput {
	return o.ApplyT(func(v *ImageTemplateVmProfileResponse) ImageTemplateVmProfileResponse {
		if v != nil {
			return *v
		}
		var ret ImageTemplateVmProfileResponse
		return ret
	}).(ImageTemplateVmProfileResponseOutput)
}

func (o ImageTemplateVmProfileResponsePtrOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ImageTemplateVmProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.OsDiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o ImageTemplateVmProfileResponsePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageTemplateVmProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

func (o ImageTemplateVmProfileResponsePtrOutput) VnetConfig() VirtualNetworkConfigResponsePtrOutput {
	return o.ApplyT(func(v *ImageTemplateVmProfileResponse) *VirtualNetworkConfigResponse {
		if v == nil {
			return nil
		}
		return v.VnetConfig
	}).(VirtualNetworkConfigResponsePtrOutput)
}

type ImageTemplateWindowsUpdateCustomizer struct {
	Filters        []string `pulumi:"filters"`
	Name           *string  `pulumi:"name"`
	SearchCriteria *string  `pulumi:"searchCriteria"`
	Type           string   `pulumi:"type"`
	UpdateLimit    *int     `pulumi:"updateLimit"`
}





type ImageTemplateWindowsUpdateCustomizerInput interface {
	pulumi.Input

	ToImageTemplateWindowsUpdateCustomizerOutput() ImageTemplateWindowsUpdateCustomizerOutput
	ToImageTemplateWindowsUpdateCustomizerOutputWithContext(context.Context) ImageTemplateWindowsUpdateCustomizerOutput
}

type ImageTemplateWindowsUpdateCustomizerArgs struct {
	Filters        pulumi.StringArrayInput `pulumi:"filters"`
	Name           pulumi.StringPtrInput   `pulumi:"name"`
	SearchCriteria pulumi.StringPtrInput   `pulumi:"searchCriteria"`
	Type           pulumi.StringInput      `pulumi:"type"`
	UpdateLimit    pulumi.IntPtrInput      `pulumi:"updateLimit"`
}

func (ImageTemplateWindowsUpdateCustomizerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateWindowsUpdateCustomizer)(nil)).Elem()
}

func (i ImageTemplateWindowsUpdateCustomizerArgs) ToImageTemplateWindowsUpdateCustomizerOutput() ImageTemplateWindowsUpdateCustomizerOutput {
	return i.ToImageTemplateWindowsUpdateCustomizerOutputWithContext(context.Background())
}

func (i ImageTemplateWindowsUpdateCustomizerArgs) ToImageTemplateWindowsUpdateCustomizerOutputWithContext(ctx context.Context) ImageTemplateWindowsUpdateCustomizerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateWindowsUpdateCustomizerOutput)
}

type ImageTemplateWindowsUpdateCustomizerOutput struct{ *pulumi.OutputState }

func (ImageTemplateWindowsUpdateCustomizerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateWindowsUpdateCustomizer)(nil)).Elem()
}

func (o ImageTemplateWindowsUpdateCustomizerOutput) ToImageTemplateWindowsUpdateCustomizerOutput() ImageTemplateWindowsUpdateCustomizerOutput {
	return o
}

func (o ImageTemplateWindowsUpdateCustomizerOutput) ToImageTemplateWindowsUpdateCustomizerOutputWithContext(ctx context.Context) ImageTemplateWindowsUpdateCustomizerOutput {
	return o
}

func (o ImageTemplateWindowsUpdateCustomizerOutput) Filters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ImageTemplateWindowsUpdateCustomizer) []string { return v.Filters }).(pulumi.StringArrayOutput)
}

func (o ImageTemplateWindowsUpdateCustomizerOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateWindowsUpdateCustomizer) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateWindowsUpdateCustomizerOutput) SearchCriteria() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateWindowsUpdateCustomizer) *string { return v.SearchCriteria }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateWindowsUpdateCustomizerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateWindowsUpdateCustomizer) string { return v.Type }).(pulumi.StringOutput)
}

func (o ImageTemplateWindowsUpdateCustomizerOutput) UpdateLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ImageTemplateWindowsUpdateCustomizer) *int { return v.UpdateLimit }).(pulumi.IntPtrOutput)
}

type ImageTemplateWindowsUpdateCustomizerResponse struct {
	Filters        []string `pulumi:"filters"`
	Name           *string  `pulumi:"name"`
	SearchCriteria *string  `pulumi:"searchCriteria"`
	Type           string   `pulumi:"type"`
	UpdateLimit    *int     `pulumi:"updateLimit"`
}





type ImageTemplateWindowsUpdateCustomizerResponseInput interface {
	pulumi.Input

	ToImageTemplateWindowsUpdateCustomizerResponseOutput() ImageTemplateWindowsUpdateCustomizerResponseOutput
	ToImageTemplateWindowsUpdateCustomizerResponseOutputWithContext(context.Context) ImageTemplateWindowsUpdateCustomizerResponseOutput
}

type ImageTemplateWindowsUpdateCustomizerResponseArgs struct {
	Filters        pulumi.StringArrayInput `pulumi:"filters"`
	Name           pulumi.StringPtrInput   `pulumi:"name"`
	SearchCriteria pulumi.StringPtrInput   `pulumi:"searchCriteria"`
	Type           pulumi.StringInput      `pulumi:"type"`
	UpdateLimit    pulumi.IntPtrInput      `pulumi:"updateLimit"`
}

func (ImageTemplateWindowsUpdateCustomizerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateWindowsUpdateCustomizerResponse)(nil)).Elem()
}

func (i ImageTemplateWindowsUpdateCustomizerResponseArgs) ToImageTemplateWindowsUpdateCustomizerResponseOutput() ImageTemplateWindowsUpdateCustomizerResponseOutput {
	return i.ToImageTemplateWindowsUpdateCustomizerResponseOutputWithContext(context.Background())
}

func (i ImageTemplateWindowsUpdateCustomizerResponseArgs) ToImageTemplateWindowsUpdateCustomizerResponseOutputWithContext(ctx context.Context) ImageTemplateWindowsUpdateCustomizerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateWindowsUpdateCustomizerResponseOutput)
}

type ImageTemplateWindowsUpdateCustomizerResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplateWindowsUpdateCustomizerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateWindowsUpdateCustomizerResponse)(nil)).Elem()
}

func (o ImageTemplateWindowsUpdateCustomizerResponseOutput) ToImageTemplateWindowsUpdateCustomizerResponseOutput() ImageTemplateWindowsUpdateCustomizerResponseOutput {
	return o
}

func (o ImageTemplateWindowsUpdateCustomizerResponseOutput) ToImageTemplateWindowsUpdateCustomizerResponseOutputWithContext(ctx context.Context) ImageTemplateWindowsUpdateCustomizerResponseOutput {
	return o
}

func (o ImageTemplateWindowsUpdateCustomizerResponseOutput) Filters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ImageTemplateWindowsUpdateCustomizerResponse) []string { return v.Filters }).(pulumi.StringArrayOutput)
}

func (o ImageTemplateWindowsUpdateCustomizerResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateWindowsUpdateCustomizerResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateWindowsUpdateCustomizerResponseOutput) SearchCriteria() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateWindowsUpdateCustomizerResponse) *string { return v.SearchCriteria }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateWindowsUpdateCustomizerResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateWindowsUpdateCustomizerResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ImageTemplateWindowsUpdateCustomizerResponseOutput) UpdateLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ImageTemplateWindowsUpdateCustomizerResponse) *int { return v.UpdateLimit }).(pulumi.IntPtrOutput)
}

type PlatformImagePurchasePlan struct {
	PlanName      string `pulumi:"planName"`
	PlanProduct   string `pulumi:"planProduct"`
	PlanPublisher string `pulumi:"planPublisher"`
}





type PlatformImagePurchasePlanInput interface {
	pulumi.Input

	ToPlatformImagePurchasePlanOutput() PlatformImagePurchasePlanOutput
	ToPlatformImagePurchasePlanOutputWithContext(context.Context) PlatformImagePurchasePlanOutput
}

type PlatformImagePurchasePlanArgs struct {
	PlanName      pulumi.StringInput `pulumi:"planName"`
	PlanProduct   pulumi.StringInput `pulumi:"planProduct"`
	PlanPublisher pulumi.StringInput `pulumi:"planPublisher"`
}

func (PlatformImagePurchasePlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PlatformImagePurchasePlan)(nil)).Elem()
}

func (i PlatformImagePurchasePlanArgs) ToPlatformImagePurchasePlanOutput() PlatformImagePurchasePlanOutput {
	return i.ToPlatformImagePurchasePlanOutputWithContext(context.Background())
}

func (i PlatformImagePurchasePlanArgs) ToPlatformImagePurchasePlanOutputWithContext(ctx context.Context) PlatformImagePurchasePlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformImagePurchasePlanOutput)
}

func (i PlatformImagePurchasePlanArgs) ToPlatformImagePurchasePlanPtrOutput() PlatformImagePurchasePlanPtrOutput {
	return i.ToPlatformImagePurchasePlanPtrOutputWithContext(context.Background())
}

func (i PlatformImagePurchasePlanArgs) ToPlatformImagePurchasePlanPtrOutputWithContext(ctx context.Context) PlatformImagePurchasePlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformImagePurchasePlanOutput).ToPlatformImagePurchasePlanPtrOutputWithContext(ctx)
}









type PlatformImagePurchasePlanPtrInput interface {
	pulumi.Input

	ToPlatformImagePurchasePlanPtrOutput() PlatformImagePurchasePlanPtrOutput
	ToPlatformImagePurchasePlanPtrOutputWithContext(context.Context) PlatformImagePurchasePlanPtrOutput
}

type platformImagePurchasePlanPtrType PlatformImagePurchasePlanArgs

func PlatformImagePurchasePlanPtr(v *PlatformImagePurchasePlanArgs) PlatformImagePurchasePlanPtrInput {
	return (*platformImagePurchasePlanPtrType)(v)
}

func (*platformImagePurchasePlanPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PlatformImagePurchasePlan)(nil)).Elem()
}

func (i *platformImagePurchasePlanPtrType) ToPlatformImagePurchasePlanPtrOutput() PlatformImagePurchasePlanPtrOutput {
	return i.ToPlatformImagePurchasePlanPtrOutputWithContext(context.Background())
}

func (i *platformImagePurchasePlanPtrType) ToPlatformImagePurchasePlanPtrOutputWithContext(ctx context.Context) PlatformImagePurchasePlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformImagePurchasePlanPtrOutput)
}

type PlatformImagePurchasePlanOutput struct{ *pulumi.OutputState }

func (PlatformImagePurchasePlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlatformImagePurchasePlan)(nil)).Elem()
}

func (o PlatformImagePurchasePlanOutput) ToPlatformImagePurchasePlanOutput() PlatformImagePurchasePlanOutput {
	return o
}

func (o PlatformImagePurchasePlanOutput) ToPlatformImagePurchasePlanOutputWithContext(ctx context.Context) PlatformImagePurchasePlanOutput {
	return o
}

func (o PlatformImagePurchasePlanOutput) ToPlatformImagePurchasePlanPtrOutput() PlatformImagePurchasePlanPtrOutput {
	return o.ToPlatformImagePurchasePlanPtrOutputWithContext(context.Background())
}

func (o PlatformImagePurchasePlanOutput) ToPlatformImagePurchasePlanPtrOutputWithContext(ctx context.Context) PlatformImagePurchasePlanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PlatformImagePurchasePlan) *PlatformImagePurchasePlan {
		return &v
	}).(PlatformImagePurchasePlanPtrOutput)
}

func (o PlatformImagePurchasePlanOutput) PlanName() pulumi.StringOutput {
	return o.ApplyT(func(v PlatformImagePurchasePlan) string { return v.PlanName }).(pulumi.StringOutput)
}

func (o PlatformImagePurchasePlanOutput) PlanProduct() pulumi.StringOutput {
	return o.ApplyT(func(v PlatformImagePurchasePlan) string { return v.PlanProduct }).(pulumi.StringOutput)
}

func (o PlatformImagePurchasePlanOutput) PlanPublisher() pulumi.StringOutput {
	return o.ApplyT(func(v PlatformImagePurchasePlan) string { return v.PlanPublisher }).(pulumi.StringOutput)
}

type PlatformImagePurchasePlanPtrOutput struct{ *pulumi.OutputState }

func (PlatformImagePurchasePlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlatformImagePurchasePlan)(nil)).Elem()
}

func (o PlatformImagePurchasePlanPtrOutput) ToPlatformImagePurchasePlanPtrOutput() PlatformImagePurchasePlanPtrOutput {
	return o
}

func (o PlatformImagePurchasePlanPtrOutput) ToPlatformImagePurchasePlanPtrOutputWithContext(ctx context.Context) PlatformImagePurchasePlanPtrOutput {
	return o
}

func (o PlatformImagePurchasePlanPtrOutput) Elem() PlatformImagePurchasePlanOutput {
	return o.ApplyT(func(v *PlatformImagePurchasePlan) PlatformImagePurchasePlan {
		if v != nil {
			return *v
		}
		var ret PlatformImagePurchasePlan
		return ret
	}).(PlatformImagePurchasePlanOutput)
}

func (o PlatformImagePurchasePlanPtrOutput) PlanName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlatformImagePurchasePlan) *string {
		if v == nil {
			return nil
		}
		return &v.PlanName
	}).(pulumi.StringPtrOutput)
}

func (o PlatformImagePurchasePlanPtrOutput) PlanProduct() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlatformImagePurchasePlan) *string {
		if v == nil {
			return nil
		}
		return &v.PlanProduct
	}).(pulumi.StringPtrOutput)
}

func (o PlatformImagePurchasePlanPtrOutput) PlanPublisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlatformImagePurchasePlan) *string {
		if v == nil {
			return nil
		}
		return &v.PlanPublisher
	}).(pulumi.StringPtrOutput)
}

type PlatformImagePurchasePlanResponse struct {
	PlanName      string `pulumi:"planName"`
	PlanProduct   string `pulumi:"planProduct"`
	PlanPublisher string `pulumi:"planPublisher"`
}





type PlatformImagePurchasePlanResponseInput interface {
	pulumi.Input

	ToPlatformImagePurchasePlanResponseOutput() PlatformImagePurchasePlanResponseOutput
	ToPlatformImagePurchasePlanResponseOutputWithContext(context.Context) PlatformImagePurchasePlanResponseOutput
}

type PlatformImagePurchasePlanResponseArgs struct {
	PlanName      pulumi.StringInput `pulumi:"planName"`
	PlanProduct   pulumi.StringInput `pulumi:"planProduct"`
	PlanPublisher pulumi.StringInput `pulumi:"planPublisher"`
}

func (PlatformImagePurchasePlanResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PlatformImagePurchasePlanResponse)(nil)).Elem()
}

func (i PlatformImagePurchasePlanResponseArgs) ToPlatformImagePurchasePlanResponseOutput() PlatformImagePurchasePlanResponseOutput {
	return i.ToPlatformImagePurchasePlanResponseOutputWithContext(context.Background())
}

func (i PlatformImagePurchasePlanResponseArgs) ToPlatformImagePurchasePlanResponseOutputWithContext(ctx context.Context) PlatformImagePurchasePlanResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformImagePurchasePlanResponseOutput)
}

func (i PlatformImagePurchasePlanResponseArgs) ToPlatformImagePurchasePlanResponsePtrOutput() PlatformImagePurchasePlanResponsePtrOutput {
	return i.ToPlatformImagePurchasePlanResponsePtrOutputWithContext(context.Background())
}

func (i PlatformImagePurchasePlanResponseArgs) ToPlatformImagePurchasePlanResponsePtrOutputWithContext(ctx context.Context) PlatformImagePurchasePlanResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformImagePurchasePlanResponseOutput).ToPlatformImagePurchasePlanResponsePtrOutputWithContext(ctx)
}









type PlatformImagePurchasePlanResponsePtrInput interface {
	pulumi.Input

	ToPlatformImagePurchasePlanResponsePtrOutput() PlatformImagePurchasePlanResponsePtrOutput
	ToPlatformImagePurchasePlanResponsePtrOutputWithContext(context.Context) PlatformImagePurchasePlanResponsePtrOutput
}

type platformImagePurchasePlanResponsePtrType PlatformImagePurchasePlanResponseArgs

func PlatformImagePurchasePlanResponsePtr(v *PlatformImagePurchasePlanResponseArgs) PlatformImagePurchasePlanResponsePtrInput {
	return (*platformImagePurchasePlanResponsePtrType)(v)
}

func (*platformImagePurchasePlanResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PlatformImagePurchasePlanResponse)(nil)).Elem()
}

func (i *platformImagePurchasePlanResponsePtrType) ToPlatformImagePurchasePlanResponsePtrOutput() PlatformImagePurchasePlanResponsePtrOutput {
	return i.ToPlatformImagePurchasePlanResponsePtrOutputWithContext(context.Background())
}

func (i *platformImagePurchasePlanResponsePtrType) ToPlatformImagePurchasePlanResponsePtrOutputWithContext(ctx context.Context) PlatformImagePurchasePlanResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformImagePurchasePlanResponsePtrOutput)
}

type PlatformImagePurchasePlanResponseOutput struct{ *pulumi.OutputState }

func (PlatformImagePurchasePlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlatformImagePurchasePlanResponse)(nil)).Elem()
}

func (o PlatformImagePurchasePlanResponseOutput) ToPlatformImagePurchasePlanResponseOutput() PlatformImagePurchasePlanResponseOutput {
	return o
}

func (o PlatformImagePurchasePlanResponseOutput) ToPlatformImagePurchasePlanResponseOutputWithContext(ctx context.Context) PlatformImagePurchasePlanResponseOutput {
	return o
}

func (o PlatformImagePurchasePlanResponseOutput) ToPlatformImagePurchasePlanResponsePtrOutput() PlatformImagePurchasePlanResponsePtrOutput {
	return o.ToPlatformImagePurchasePlanResponsePtrOutputWithContext(context.Background())
}

func (o PlatformImagePurchasePlanResponseOutput) ToPlatformImagePurchasePlanResponsePtrOutputWithContext(ctx context.Context) PlatformImagePurchasePlanResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PlatformImagePurchasePlanResponse) *PlatformImagePurchasePlanResponse {
		return &v
	}).(PlatformImagePurchasePlanResponsePtrOutput)
}

func (o PlatformImagePurchasePlanResponseOutput) PlanName() pulumi.StringOutput {
	return o.ApplyT(func(v PlatformImagePurchasePlanResponse) string { return v.PlanName }).(pulumi.StringOutput)
}

func (o PlatformImagePurchasePlanResponseOutput) PlanProduct() pulumi.StringOutput {
	return o.ApplyT(func(v PlatformImagePurchasePlanResponse) string { return v.PlanProduct }).(pulumi.StringOutput)
}

func (o PlatformImagePurchasePlanResponseOutput) PlanPublisher() pulumi.StringOutput {
	return o.ApplyT(func(v PlatformImagePurchasePlanResponse) string { return v.PlanPublisher }).(pulumi.StringOutput)
}

type PlatformImagePurchasePlanResponsePtrOutput struct{ *pulumi.OutputState }

func (PlatformImagePurchasePlanResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlatformImagePurchasePlanResponse)(nil)).Elem()
}

func (o PlatformImagePurchasePlanResponsePtrOutput) ToPlatformImagePurchasePlanResponsePtrOutput() PlatformImagePurchasePlanResponsePtrOutput {
	return o
}

func (o PlatformImagePurchasePlanResponsePtrOutput) ToPlatformImagePurchasePlanResponsePtrOutputWithContext(ctx context.Context) PlatformImagePurchasePlanResponsePtrOutput {
	return o
}

func (o PlatformImagePurchasePlanResponsePtrOutput) Elem() PlatformImagePurchasePlanResponseOutput {
	return o.ApplyT(func(v *PlatformImagePurchasePlanResponse) PlatformImagePurchasePlanResponse {
		if v != nil {
			return *v
		}
		var ret PlatformImagePurchasePlanResponse
		return ret
	}).(PlatformImagePurchasePlanResponseOutput)
}

func (o PlatformImagePurchasePlanResponsePtrOutput) PlanName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlatformImagePurchasePlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PlanName
	}).(pulumi.StringPtrOutput)
}

func (o PlatformImagePurchasePlanResponsePtrOutput) PlanProduct() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlatformImagePurchasePlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PlanProduct
	}).(pulumi.StringPtrOutput)
}

func (o PlatformImagePurchasePlanResponsePtrOutput) PlanPublisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlatformImagePurchasePlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PlanPublisher
	}).(pulumi.StringPtrOutput)
}

type ProvisioningErrorResponse struct {
	Message               *string `pulumi:"message"`
	ProvisioningErrorCode *string `pulumi:"provisioningErrorCode"`
}





type ProvisioningErrorResponseInput interface {
	pulumi.Input

	ToProvisioningErrorResponseOutput() ProvisioningErrorResponseOutput
	ToProvisioningErrorResponseOutputWithContext(context.Context) ProvisioningErrorResponseOutput
}

type ProvisioningErrorResponseArgs struct {
	Message               pulumi.StringPtrInput `pulumi:"message"`
	ProvisioningErrorCode pulumi.StringPtrInput `pulumi:"provisioningErrorCode"`
}

func (ProvisioningErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningErrorResponse)(nil)).Elem()
}

func (i ProvisioningErrorResponseArgs) ToProvisioningErrorResponseOutput() ProvisioningErrorResponseOutput {
	return i.ToProvisioningErrorResponseOutputWithContext(context.Background())
}

func (i ProvisioningErrorResponseArgs) ToProvisioningErrorResponseOutputWithContext(ctx context.Context) ProvisioningErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisioningErrorResponseOutput)
}

func (i ProvisioningErrorResponseArgs) ToProvisioningErrorResponsePtrOutput() ProvisioningErrorResponsePtrOutput {
	return i.ToProvisioningErrorResponsePtrOutputWithContext(context.Background())
}

func (i ProvisioningErrorResponseArgs) ToProvisioningErrorResponsePtrOutputWithContext(ctx context.Context) ProvisioningErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisioningErrorResponseOutput).ToProvisioningErrorResponsePtrOutputWithContext(ctx)
}









type ProvisioningErrorResponsePtrInput interface {
	pulumi.Input

	ToProvisioningErrorResponsePtrOutput() ProvisioningErrorResponsePtrOutput
	ToProvisioningErrorResponsePtrOutputWithContext(context.Context) ProvisioningErrorResponsePtrOutput
}

type provisioningErrorResponsePtrType ProvisioningErrorResponseArgs

func ProvisioningErrorResponsePtr(v *ProvisioningErrorResponseArgs) ProvisioningErrorResponsePtrInput {
	return (*provisioningErrorResponsePtrType)(v)
}

func (*provisioningErrorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisioningErrorResponse)(nil)).Elem()
}

func (i *provisioningErrorResponsePtrType) ToProvisioningErrorResponsePtrOutput() ProvisioningErrorResponsePtrOutput {
	return i.ToProvisioningErrorResponsePtrOutputWithContext(context.Background())
}

func (i *provisioningErrorResponsePtrType) ToProvisioningErrorResponsePtrOutputWithContext(ctx context.Context) ProvisioningErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisioningErrorResponsePtrOutput)
}

type ProvisioningErrorResponseOutput struct{ *pulumi.OutputState }

func (ProvisioningErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningErrorResponse)(nil)).Elem()
}

func (o ProvisioningErrorResponseOutput) ToProvisioningErrorResponseOutput() ProvisioningErrorResponseOutput {
	return o
}

func (o ProvisioningErrorResponseOutput) ToProvisioningErrorResponseOutputWithContext(ctx context.Context) ProvisioningErrorResponseOutput {
	return o
}

func (o ProvisioningErrorResponseOutput) ToProvisioningErrorResponsePtrOutput() ProvisioningErrorResponsePtrOutput {
	return o.ToProvisioningErrorResponsePtrOutputWithContext(context.Background())
}

func (o ProvisioningErrorResponseOutput) ToProvisioningErrorResponsePtrOutputWithContext(ctx context.Context) ProvisioningErrorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProvisioningErrorResponse) *ProvisioningErrorResponse {
		return &v
	}).(ProvisioningErrorResponsePtrOutput)
}

func (o ProvisioningErrorResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisioningErrorResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ProvisioningErrorResponseOutput) ProvisioningErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisioningErrorResponse) *string { return v.ProvisioningErrorCode }).(pulumi.StringPtrOutput)
}

type ProvisioningErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (ProvisioningErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisioningErrorResponse)(nil)).Elem()
}

func (o ProvisioningErrorResponsePtrOutput) ToProvisioningErrorResponsePtrOutput() ProvisioningErrorResponsePtrOutput {
	return o
}

func (o ProvisioningErrorResponsePtrOutput) ToProvisioningErrorResponsePtrOutputWithContext(ctx context.Context) ProvisioningErrorResponsePtrOutput {
	return o
}

func (o ProvisioningErrorResponsePtrOutput) Elem() ProvisioningErrorResponseOutput {
	return o.ApplyT(func(v *ProvisioningErrorResponse) ProvisioningErrorResponse {
		if v != nil {
			return *v
		}
		var ret ProvisioningErrorResponse
		return ret
	}).(ProvisioningErrorResponseOutput)
}

func (o ProvisioningErrorResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisioningErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ProvisioningErrorResponsePtrOutput) ProvisioningErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisioningErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningErrorCode
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkConfig struct {
	SubnetId *string `pulumi:"subnetId"`
}





type VirtualNetworkConfigInput interface {
	pulumi.Input

	ToVirtualNetworkConfigOutput() VirtualNetworkConfigOutput
	ToVirtualNetworkConfigOutputWithContext(context.Context) VirtualNetworkConfigOutput
}

type VirtualNetworkConfigArgs struct {
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
}

func (VirtualNetworkConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkConfig)(nil)).Elem()
}

func (i VirtualNetworkConfigArgs) ToVirtualNetworkConfigOutput() VirtualNetworkConfigOutput {
	return i.ToVirtualNetworkConfigOutputWithContext(context.Background())
}

func (i VirtualNetworkConfigArgs) ToVirtualNetworkConfigOutputWithContext(ctx context.Context) VirtualNetworkConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigOutput)
}

func (i VirtualNetworkConfigArgs) ToVirtualNetworkConfigPtrOutput() VirtualNetworkConfigPtrOutput {
	return i.ToVirtualNetworkConfigPtrOutputWithContext(context.Background())
}

func (i VirtualNetworkConfigArgs) ToVirtualNetworkConfigPtrOutputWithContext(ctx context.Context) VirtualNetworkConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigOutput).ToVirtualNetworkConfigPtrOutputWithContext(ctx)
}









type VirtualNetworkConfigPtrInput interface {
	pulumi.Input

	ToVirtualNetworkConfigPtrOutput() VirtualNetworkConfigPtrOutput
	ToVirtualNetworkConfigPtrOutputWithContext(context.Context) VirtualNetworkConfigPtrOutput
}

type virtualNetworkConfigPtrType VirtualNetworkConfigArgs

func VirtualNetworkConfigPtr(v *VirtualNetworkConfigArgs) VirtualNetworkConfigPtrInput {
	return (*virtualNetworkConfigPtrType)(v)
}

func (*virtualNetworkConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkConfig)(nil)).Elem()
}

func (i *virtualNetworkConfigPtrType) ToVirtualNetworkConfigPtrOutput() VirtualNetworkConfigPtrOutput {
	return i.ToVirtualNetworkConfigPtrOutputWithContext(context.Background())
}

func (i *virtualNetworkConfigPtrType) ToVirtualNetworkConfigPtrOutputWithContext(ctx context.Context) VirtualNetworkConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigPtrOutput)
}

type VirtualNetworkConfigOutput struct{ *pulumi.OutputState }

func (VirtualNetworkConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkConfig)(nil)).Elem()
}

func (o VirtualNetworkConfigOutput) ToVirtualNetworkConfigOutput() VirtualNetworkConfigOutput {
	return o
}

func (o VirtualNetworkConfigOutput) ToVirtualNetworkConfigOutputWithContext(ctx context.Context) VirtualNetworkConfigOutput {
	return o
}

func (o VirtualNetworkConfigOutput) ToVirtualNetworkConfigPtrOutput() VirtualNetworkConfigPtrOutput {
	return o.ToVirtualNetworkConfigPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkConfigOutput) ToVirtualNetworkConfigPtrOutputWithContext(ctx context.Context) VirtualNetworkConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkConfig) *VirtualNetworkConfig {
		return &v
	}).(VirtualNetworkConfigPtrOutput)
}

func (o VirtualNetworkConfigOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkConfig) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

type VirtualNetworkConfigPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkConfig)(nil)).Elem()
}

func (o VirtualNetworkConfigPtrOutput) ToVirtualNetworkConfigPtrOutput() VirtualNetworkConfigPtrOutput {
	return o
}

func (o VirtualNetworkConfigPtrOutput) ToVirtualNetworkConfigPtrOutputWithContext(ctx context.Context) VirtualNetworkConfigPtrOutput {
	return o
}

func (o VirtualNetworkConfigPtrOutput) Elem() VirtualNetworkConfigOutput {
	return o.ApplyT(func(v *VirtualNetworkConfig) VirtualNetworkConfig {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkConfig
		return ret
	}).(VirtualNetworkConfigOutput)
}

func (o VirtualNetworkConfigPtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfig) *string {
		if v == nil {
			return nil
		}
		return v.SubnetId
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkConfigResponse struct {
	SubnetId *string `pulumi:"subnetId"`
}





type VirtualNetworkConfigResponseInput interface {
	pulumi.Input

	ToVirtualNetworkConfigResponseOutput() VirtualNetworkConfigResponseOutput
	ToVirtualNetworkConfigResponseOutputWithContext(context.Context) VirtualNetworkConfigResponseOutput
}

type VirtualNetworkConfigResponseArgs struct {
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
}

func (VirtualNetworkConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkConfigResponse)(nil)).Elem()
}

func (i VirtualNetworkConfigResponseArgs) ToVirtualNetworkConfigResponseOutput() VirtualNetworkConfigResponseOutput {
	return i.ToVirtualNetworkConfigResponseOutputWithContext(context.Background())
}

func (i VirtualNetworkConfigResponseArgs) ToVirtualNetworkConfigResponseOutputWithContext(ctx context.Context) VirtualNetworkConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigResponseOutput)
}

func (i VirtualNetworkConfigResponseArgs) ToVirtualNetworkConfigResponsePtrOutput() VirtualNetworkConfigResponsePtrOutput {
	return i.ToVirtualNetworkConfigResponsePtrOutputWithContext(context.Background())
}

func (i VirtualNetworkConfigResponseArgs) ToVirtualNetworkConfigResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigResponseOutput).ToVirtualNetworkConfigResponsePtrOutputWithContext(ctx)
}









type VirtualNetworkConfigResponsePtrInput interface {
	pulumi.Input

	ToVirtualNetworkConfigResponsePtrOutput() VirtualNetworkConfigResponsePtrOutput
	ToVirtualNetworkConfigResponsePtrOutputWithContext(context.Context) VirtualNetworkConfigResponsePtrOutput
}

type virtualNetworkConfigResponsePtrType VirtualNetworkConfigResponseArgs

func VirtualNetworkConfigResponsePtr(v *VirtualNetworkConfigResponseArgs) VirtualNetworkConfigResponsePtrInput {
	return (*virtualNetworkConfigResponsePtrType)(v)
}

func (*virtualNetworkConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkConfigResponse)(nil)).Elem()
}

func (i *virtualNetworkConfigResponsePtrType) ToVirtualNetworkConfigResponsePtrOutput() VirtualNetworkConfigResponsePtrOutput {
	return i.ToVirtualNetworkConfigResponsePtrOutputWithContext(context.Background())
}

func (i *virtualNetworkConfigResponsePtrType) ToVirtualNetworkConfigResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigResponsePtrOutput)
}

type VirtualNetworkConfigResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkConfigResponse)(nil)).Elem()
}

func (o VirtualNetworkConfigResponseOutput) ToVirtualNetworkConfigResponseOutput() VirtualNetworkConfigResponseOutput {
	return o
}

func (o VirtualNetworkConfigResponseOutput) ToVirtualNetworkConfigResponseOutputWithContext(ctx context.Context) VirtualNetworkConfigResponseOutput {
	return o
}

func (o VirtualNetworkConfigResponseOutput) ToVirtualNetworkConfigResponsePtrOutput() VirtualNetworkConfigResponsePtrOutput {
	return o.ToVirtualNetworkConfigResponsePtrOutputWithContext(context.Background())
}

func (o VirtualNetworkConfigResponseOutput) ToVirtualNetworkConfigResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkConfigResponse) *VirtualNetworkConfigResponse {
		return &v
	}).(VirtualNetworkConfigResponsePtrOutput)
}

func (o VirtualNetworkConfigResponseOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkConfigResponse) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

type VirtualNetworkConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkConfigResponse)(nil)).Elem()
}

func (o VirtualNetworkConfigResponsePtrOutput) ToVirtualNetworkConfigResponsePtrOutput() VirtualNetworkConfigResponsePtrOutput {
	return o
}

func (o VirtualNetworkConfigResponsePtrOutput) ToVirtualNetworkConfigResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkConfigResponsePtrOutput {
	return o
}

func (o VirtualNetworkConfigResponsePtrOutput) Elem() VirtualNetworkConfigResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkConfigResponse) VirtualNetworkConfigResponse {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkConfigResponse
		return ret
	}).(VirtualNetworkConfigResponseOutput)
}

func (o VirtualNetworkConfigResponsePtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubnetId
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ImageTemplateFileCustomizerOutput{})
	pulumi.RegisterOutputType(ImageTemplateFileCustomizerResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateIdentityOutput{})
	pulumi.RegisterOutputType(ImageTemplateIdentityPtrOutput{})
	pulumi.RegisterOutputType(ImageTemplateIdentityResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageTemplateIdentityResponseUserAssignedIdentitiesOutput{})
	pulumi.RegisterOutputType(ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput{})
	pulumi.RegisterOutputType(ImageTemplateLastRunStatusResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateLastRunStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageTemplateManagedImageDistributorOutput{})
	pulumi.RegisterOutputType(ImageTemplateManagedImageDistributorResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateManagedImageSourceOutput{})
	pulumi.RegisterOutputType(ImageTemplateManagedImageSourceResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplatePlatformImageSourceOutput{})
	pulumi.RegisterOutputType(ImageTemplatePlatformImageSourceResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplatePowerShellCustomizerOutput{})
	pulumi.RegisterOutputType(ImageTemplatePowerShellCustomizerResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateRestartCustomizerOutput{})
	pulumi.RegisterOutputType(ImageTemplateRestartCustomizerResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateSharedImageDistributorOutput{})
	pulumi.RegisterOutputType(ImageTemplateSharedImageDistributorResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateSharedImageVersionSourceOutput{})
	pulumi.RegisterOutputType(ImageTemplateSharedImageVersionSourceResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateShellCustomizerOutput{})
	pulumi.RegisterOutputType(ImageTemplateShellCustomizerResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateVhdDistributorOutput{})
	pulumi.RegisterOutputType(ImageTemplateVhdDistributorResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateVmProfileOutput{})
	pulumi.RegisterOutputType(ImageTemplateVmProfilePtrOutput{})
	pulumi.RegisterOutputType(ImageTemplateVmProfileResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateVmProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageTemplateWindowsUpdateCustomizerOutput{})
	pulumi.RegisterOutputType(ImageTemplateWindowsUpdateCustomizerResponseOutput{})
	pulumi.RegisterOutputType(PlatformImagePurchasePlanOutput{})
	pulumi.RegisterOutputType(PlatformImagePurchasePlanPtrOutput{})
	pulumi.RegisterOutputType(PlatformImagePurchasePlanResponseOutput{})
	pulumi.RegisterOutputType(PlatformImagePurchasePlanResponsePtrOutput{})
	pulumi.RegisterOutputType(ProvisioningErrorResponseOutput{})
	pulumi.RegisterOutputType(ProvisioningErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigResponsePtrOutput{})
}

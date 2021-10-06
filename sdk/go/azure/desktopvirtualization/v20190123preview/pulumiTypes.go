


package v20190123preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RegistrationInfo struct {
	ExpirationTime *string `pulumi:"expirationTime"`
	ResetToken     *bool   `pulumi:"resetToken"`
	Token          *string `pulumi:"token"`
}





type RegistrationInfoInput interface {
	pulumi.Input

	ToRegistrationInfoOutput() RegistrationInfoOutput
	ToRegistrationInfoOutputWithContext(context.Context) RegistrationInfoOutput
}

type RegistrationInfoArgs struct {
	ExpirationTime pulumi.StringPtrInput `pulumi:"expirationTime"`
	ResetToken     pulumi.BoolPtrInput   `pulumi:"resetToken"`
	Token          pulumi.StringPtrInput `pulumi:"token"`
}

func (RegistrationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationInfo)(nil)).Elem()
}

func (i RegistrationInfoArgs) ToRegistrationInfoOutput() RegistrationInfoOutput {
	return i.ToRegistrationInfoOutputWithContext(context.Background())
}

func (i RegistrationInfoArgs) ToRegistrationInfoOutputWithContext(ctx context.Context) RegistrationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationInfoOutput)
}

func (i RegistrationInfoArgs) ToRegistrationInfoPtrOutput() RegistrationInfoPtrOutput {
	return i.ToRegistrationInfoPtrOutputWithContext(context.Background())
}

func (i RegistrationInfoArgs) ToRegistrationInfoPtrOutputWithContext(ctx context.Context) RegistrationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationInfoOutput).ToRegistrationInfoPtrOutputWithContext(ctx)
}









type RegistrationInfoPtrInput interface {
	pulumi.Input

	ToRegistrationInfoPtrOutput() RegistrationInfoPtrOutput
	ToRegistrationInfoPtrOutputWithContext(context.Context) RegistrationInfoPtrOutput
}

type registrationInfoPtrType RegistrationInfoArgs

func RegistrationInfoPtr(v *RegistrationInfoArgs) RegistrationInfoPtrInput {
	return (*registrationInfoPtrType)(v)
}

func (*registrationInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationInfo)(nil)).Elem()
}

func (i *registrationInfoPtrType) ToRegistrationInfoPtrOutput() RegistrationInfoPtrOutput {
	return i.ToRegistrationInfoPtrOutputWithContext(context.Background())
}

func (i *registrationInfoPtrType) ToRegistrationInfoPtrOutputWithContext(ctx context.Context) RegistrationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationInfoPtrOutput)
}

type RegistrationInfoOutput struct{ *pulumi.OutputState }

func (RegistrationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationInfo)(nil)).Elem()
}

func (o RegistrationInfoOutput) ToRegistrationInfoOutput() RegistrationInfoOutput {
	return o
}

func (o RegistrationInfoOutput) ToRegistrationInfoOutputWithContext(ctx context.Context) RegistrationInfoOutput {
	return o
}

func (o RegistrationInfoOutput) ToRegistrationInfoPtrOutput() RegistrationInfoPtrOutput {
	return o.ToRegistrationInfoPtrOutputWithContext(context.Background())
}

func (o RegistrationInfoOutput) ToRegistrationInfoPtrOutputWithContext(ctx context.Context) RegistrationInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RegistrationInfo) *RegistrationInfo {
		return &v
	}).(RegistrationInfoPtrOutput)
}

func (o RegistrationInfoOutput) ExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfo) *string { return v.ExpirationTime }).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoOutput) ResetToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RegistrationInfo) *bool { return v.ResetToken }).(pulumi.BoolPtrOutput)
}

func (o RegistrationInfoOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfo) *string { return v.Token }).(pulumi.StringPtrOutput)
}

type RegistrationInfoPtrOutput struct{ *pulumi.OutputState }

func (RegistrationInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationInfo)(nil)).Elem()
}

func (o RegistrationInfoPtrOutput) ToRegistrationInfoPtrOutput() RegistrationInfoPtrOutput {
	return o
}

func (o RegistrationInfoPtrOutput) ToRegistrationInfoPtrOutputWithContext(ctx context.Context) RegistrationInfoPtrOutput {
	return o
}

func (o RegistrationInfoPtrOutput) Elem() RegistrationInfoOutput {
	return o.ApplyT(func(v *RegistrationInfo) RegistrationInfo {
		if v != nil {
			return *v
		}
		var ret RegistrationInfo
		return ret
	}).(RegistrationInfoOutput)
}

func (o RegistrationInfoPtrOutput) ExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfo) *string {
		if v == nil {
			return nil
		}
		return v.ExpirationTime
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoPtrOutput) ResetToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RegistrationInfo) *bool {
		if v == nil {
			return nil
		}
		return v.ResetToken
	}).(pulumi.BoolPtrOutput)
}

func (o RegistrationInfoPtrOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfo) *string {
		if v == nil {
			return nil
		}
		return v.Token
	}).(pulumi.StringPtrOutput)
}

type RegistrationInfoResponse struct {
	ExpirationTime *string `pulumi:"expirationTime"`
	ResetToken     *bool   `pulumi:"resetToken"`
	Token          *string `pulumi:"token"`
}





type RegistrationInfoResponseInput interface {
	pulumi.Input

	ToRegistrationInfoResponseOutput() RegistrationInfoResponseOutput
	ToRegistrationInfoResponseOutputWithContext(context.Context) RegistrationInfoResponseOutput
}

type RegistrationInfoResponseArgs struct {
	ExpirationTime pulumi.StringPtrInput `pulumi:"expirationTime"`
	ResetToken     pulumi.BoolPtrInput   `pulumi:"resetToken"`
	Token          pulumi.StringPtrInput `pulumi:"token"`
}

func (RegistrationInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationInfoResponse)(nil)).Elem()
}

func (i RegistrationInfoResponseArgs) ToRegistrationInfoResponseOutput() RegistrationInfoResponseOutput {
	return i.ToRegistrationInfoResponseOutputWithContext(context.Background())
}

func (i RegistrationInfoResponseArgs) ToRegistrationInfoResponseOutputWithContext(ctx context.Context) RegistrationInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationInfoResponseOutput)
}

func (i RegistrationInfoResponseArgs) ToRegistrationInfoResponsePtrOutput() RegistrationInfoResponsePtrOutput {
	return i.ToRegistrationInfoResponsePtrOutputWithContext(context.Background())
}

func (i RegistrationInfoResponseArgs) ToRegistrationInfoResponsePtrOutputWithContext(ctx context.Context) RegistrationInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationInfoResponseOutput).ToRegistrationInfoResponsePtrOutputWithContext(ctx)
}









type RegistrationInfoResponsePtrInput interface {
	pulumi.Input

	ToRegistrationInfoResponsePtrOutput() RegistrationInfoResponsePtrOutput
	ToRegistrationInfoResponsePtrOutputWithContext(context.Context) RegistrationInfoResponsePtrOutput
}

type registrationInfoResponsePtrType RegistrationInfoResponseArgs

func RegistrationInfoResponsePtr(v *RegistrationInfoResponseArgs) RegistrationInfoResponsePtrInput {
	return (*registrationInfoResponsePtrType)(v)
}

func (*registrationInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationInfoResponse)(nil)).Elem()
}

func (i *registrationInfoResponsePtrType) ToRegistrationInfoResponsePtrOutput() RegistrationInfoResponsePtrOutput {
	return i.ToRegistrationInfoResponsePtrOutputWithContext(context.Background())
}

func (i *registrationInfoResponsePtrType) ToRegistrationInfoResponsePtrOutputWithContext(ctx context.Context) RegistrationInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationInfoResponsePtrOutput)
}

type RegistrationInfoResponseOutput struct{ *pulumi.OutputState }

func (RegistrationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationInfoResponse)(nil)).Elem()
}

func (o RegistrationInfoResponseOutput) ToRegistrationInfoResponseOutput() RegistrationInfoResponseOutput {
	return o
}

func (o RegistrationInfoResponseOutput) ToRegistrationInfoResponseOutputWithContext(ctx context.Context) RegistrationInfoResponseOutput {
	return o
}

func (o RegistrationInfoResponseOutput) ToRegistrationInfoResponsePtrOutput() RegistrationInfoResponsePtrOutput {
	return o.ToRegistrationInfoResponsePtrOutputWithContext(context.Background())
}

func (o RegistrationInfoResponseOutput) ToRegistrationInfoResponsePtrOutputWithContext(ctx context.Context) RegistrationInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RegistrationInfoResponse) *RegistrationInfoResponse {
		return &v
	}).(RegistrationInfoResponsePtrOutput)
}

func (o RegistrationInfoResponseOutput) ExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfoResponse) *string { return v.ExpirationTime }).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoResponseOutput) ResetToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RegistrationInfoResponse) *bool { return v.ResetToken }).(pulumi.BoolPtrOutput)
}

func (o RegistrationInfoResponseOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfoResponse) *string { return v.Token }).(pulumi.StringPtrOutput)
}

type RegistrationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (RegistrationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationInfoResponse)(nil)).Elem()
}

func (o RegistrationInfoResponsePtrOutput) ToRegistrationInfoResponsePtrOutput() RegistrationInfoResponsePtrOutput {
	return o
}

func (o RegistrationInfoResponsePtrOutput) ToRegistrationInfoResponsePtrOutputWithContext(ctx context.Context) RegistrationInfoResponsePtrOutput {
	return o
}

func (o RegistrationInfoResponsePtrOutput) Elem() RegistrationInfoResponseOutput {
	return o.ApplyT(func(v *RegistrationInfoResponse) RegistrationInfoResponse {
		if v != nil {
			return *v
		}
		var ret RegistrationInfoResponse
		return ret
	}).(RegistrationInfoResponseOutput)
}

func (o RegistrationInfoResponsePtrOutput) ExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExpirationTime
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoResponsePtrOutput) ResetToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RegistrationInfoResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ResetToken
	}).(pulumi.BoolPtrOutput)
}

func (o RegistrationInfoResponsePtrOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Token
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistrationInfoOutput{})
	pulumi.RegisterOutputType(RegistrationInfoPtrOutput{})
	pulumi.RegisterOutputType(RegistrationInfoResponseOutput{})
	pulumi.RegisterOutputType(RegistrationInfoResponsePtrOutput{})
}

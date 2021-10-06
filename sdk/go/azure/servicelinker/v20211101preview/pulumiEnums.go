


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AuthType string

const (
	AuthTypeSystemAssignedIdentity      = AuthType("systemAssignedIdentity")
	AuthTypeUserAssignedIdentity        = AuthType("userAssignedIdentity")
	AuthTypeServicePrincipalSecret      = AuthType("servicePrincipalSecret")
	AuthTypeServicePrincipalCertificate = AuthType("servicePrincipalCertificate")
	AuthTypeSecret                      = AuthType("secret")
)

func (AuthType) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthType)(nil)).Elem()
}

func (e AuthType) ToAuthTypeOutput() AuthTypeOutput {
	return pulumi.ToOutput(e).(AuthTypeOutput)
}

func (e AuthType) ToAuthTypeOutputWithContext(ctx context.Context) AuthTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AuthTypeOutput)
}

func (e AuthType) ToAuthTypePtrOutput() AuthTypePtrOutput {
	return e.ToAuthTypePtrOutputWithContext(context.Background())
}

func (e AuthType) ToAuthTypePtrOutputWithContext(ctx context.Context) AuthTypePtrOutput {
	return AuthType(e).ToAuthTypeOutputWithContext(ctx).ToAuthTypePtrOutputWithContext(ctx)
}

func (e AuthType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AuthType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AuthTypeOutput struct{ *pulumi.OutputState }

func (AuthTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthType)(nil)).Elem()
}

func (o AuthTypeOutput) ToAuthTypeOutput() AuthTypeOutput {
	return o
}

func (o AuthTypeOutput) ToAuthTypeOutputWithContext(ctx context.Context) AuthTypeOutput {
	return o
}

func (o AuthTypeOutput) ToAuthTypePtrOutput() AuthTypePtrOutput {
	return o.ToAuthTypePtrOutputWithContext(context.Background())
}

func (o AuthTypeOutput) ToAuthTypePtrOutputWithContext(ctx context.Context) AuthTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthType) *AuthType {
		return &v
	}).(AuthTypePtrOutput)
}

func (o AuthTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AuthTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AuthTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AuthTypePtrOutput struct{ *pulumi.OutputState }

func (AuthTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthType)(nil)).Elem()
}

func (o AuthTypePtrOutput) ToAuthTypePtrOutput() AuthTypePtrOutput {
	return o
}

func (o AuthTypePtrOutput) ToAuthTypePtrOutputWithContext(ctx context.Context) AuthTypePtrOutput {
	return o
}

func (o AuthTypePtrOutput) Elem() AuthTypeOutput {
	return o.ApplyT(func(v *AuthType) AuthType {
		if v != nil {
			return *v
		}
		var ret AuthType
		return ret
	}).(AuthTypeOutput)
}

func (o AuthTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AuthType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AuthTypeInput interface {
	pulumi.Input

	ToAuthTypeOutput() AuthTypeOutput
	ToAuthTypeOutputWithContext(context.Context) AuthTypeOutput
}

var authTypePtrType = reflect.TypeOf((**AuthType)(nil)).Elem()

type AuthTypePtrInput interface {
	pulumi.Input

	ToAuthTypePtrOutput() AuthTypePtrOutput
	ToAuthTypePtrOutputWithContext(context.Context) AuthTypePtrOutput
}

type authTypePtr string

func AuthTypePtr(v string) AuthTypePtrInput {
	return (*authTypePtr)(&v)
}

func (*authTypePtr) ElementType() reflect.Type {
	return authTypePtrType
}

func (in *authTypePtr) ToAuthTypePtrOutput() AuthTypePtrOutput {
	return pulumi.ToOutput(in).(AuthTypePtrOutput)
}

func (in *authTypePtr) ToAuthTypePtrOutputWithContext(ctx context.Context) AuthTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AuthTypePtrOutput)
}

type ClientType string

const (
	ClientTypeNone       = ClientType("none")
	ClientTypeDotnet     = ClientType("dotnet")
	ClientTypeJava       = ClientType("java")
	ClientTypePython     = ClientType("python")
	ClientTypeGo         = ClientType("go")
	ClientTypePhp        = ClientType("php")
	ClientTypeRuby       = ClientType("ruby")
	ClientTypeDjango     = ClientType("django")
	ClientTypeNodejs     = ClientType("nodejs")
	ClientTypeSpringBoot = ClientType("springBoot")
)

func (ClientType) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientType)(nil)).Elem()
}

func (e ClientType) ToClientTypeOutput() ClientTypeOutput {
	return pulumi.ToOutput(e).(ClientTypeOutput)
}

func (e ClientType) ToClientTypeOutputWithContext(ctx context.Context) ClientTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ClientTypeOutput)
}

func (e ClientType) ToClientTypePtrOutput() ClientTypePtrOutput {
	return e.ToClientTypePtrOutputWithContext(context.Background())
}

func (e ClientType) ToClientTypePtrOutputWithContext(ctx context.Context) ClientTypePtrOutput {
	return ClientType(e).ToClientTypeOutputWithContext(ctx).ToClientTypePtrOutputWithContext(ctx)
}

func (e ClientType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClientType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClientType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ClientType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ClientTypeOutput struct{ *pulumi.OutputState }

func (ClientTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientType)(nil)).Elem()
}

func (o ClientTypeOutput) ToClientTypeOutput() ClientTypeOutput {
	return o
}

func (o ClientTypeOutput) ToClientTypeOutputWithContext(ctx context.Context) ClientTypeOutput {
	return o
}

func (o ClientTypeOutput) ToClientTypePtrOutput() ClientTypePtrOutput {
	return o.ToClientTypePtrOutputWithContext(context.Background())
}

func (o ClientTypeOutput) ToClientTypePtrOutputWithContext(ctx context.Context) ClientTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClientType) *ClientType {
		return &v
	}).(ClientTypePtrOutput)
}

func (o ClientTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ClientTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClientType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ClientTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClientTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClientType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ClientTypePtrOutput struct{ *pulumi.OutputState }

func (ClientTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientType)(nil)).Elem()
}

func (o ClientTypePtrOutput) ToClientTypePtrOutput() ClientTypePtrOutput {
	return o
}

func (o ClientTypePtrOutput) ToClientTypePtrOutputWithContext(ctx context.Context) ClientTypePtrOutput {
	return o
}

func (o ClientTypePtrOutput) Elem() ClientTypeOutput {
	return o.ApplyT(func(v *ClientType) ClientType {
		if v != nil {
			return *v
		}
		var ret ClientType
		return ret
	}).(ClientTypeOutput)
}

func (o ClientTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClientTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ClientType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ClientTypeInput interface {
	pulumi.Input

	ToClientTypeOutput() ClientTypeOutput
	ToClientTypeOutputWithContext(context.Context) ClientTypeOutput
}

var clientTypePtrType = reflect.TypeOf((**ClientType)(nil)).Elem()

type ClientTypePtrInput interface {
	pulumi.Input

	ToClientTypePtrOutput() ClientTypePtrOutput
	ToClientTypePtrOutputWithContext(context.Context) ClientTypePtrOutput
}

type clientTypePtr string

func ClientTypePtr(v string) ClientTypePtrInput {
	return (*clientTypePtr)(&v)
}

func (*clientTypePtr) ElementType() reflect.Type {
	return clientTypePtrType
}

func (in *clientTypePtr) ToClientTypePtrOutput() ClientTypePtrOutput {
	return pulumi.ToOutput(in).(ClientTypePtrOutput)
}

func (in *clientTypePtr) ToClientTypePtrOutputWithContext(ctx context.Context) ClientTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ClientTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthTypeOutput{})
	pulumi.RegisterOutputType(AuthTypePtrOutput{})
	pulumi.RegisterOutputType(ClientTypeOutput{})
	pulumi.RegisterOutputType(ClientTypePtrOutput{})
}

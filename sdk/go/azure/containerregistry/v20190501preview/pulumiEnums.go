


package v20190501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TokenCertificateName string

const (
	TokenCertificateNameCertificate1 = TokenCertificateName("certificate1")
	TokenCertificateNameCertificate2 = TokenCertificateName("certificate2")
)

func (TokenCertificateName) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCertificateName)(nil)).Elem()
}

func (e TokenCertificateName) ToTokenCertificateNameOutput() TokenCertificateNameOutput {
	return pulumi.ToOutput(e).(TokenCertificateNameOutput)
}

func (e TokenCertificateName) ToTokenCertificateNameOutputWithContext(ctx context.Context) TokenCertificateNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TokenCertificateNameOutput)
}

func (e TokenCertificateName) ToTokenCertificateNamePtrOutput() TokenCertificateNamePtrOutput {
	return e.ToTokenCertificateNamePtrOutputWithContext(context.Background())
}

func (e TokenCertificateName) ToTokenCertificateNamePtrOutputWithContext(ctx context.Context) TokenCertificateNamePtrOutput {
	return TokenCertificateName(e).ToTokenCertificateNameOutputWithContext(ctx).ToTokenCertificateNamePtrOutputWithContext(ctx)
}

func (e TokenCertificateName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TokenCertificateName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TokenCertificateName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TokenCertificateName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TokenCertificateNameOutput struct{ *pulumi.OutputState }

func (TokenCertificateNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCertificateName)(nil)).Elem()
}

func (o TokenCertificateNameOutput) ToTokenCertificateNameOutput() TokenCertificateNameOutput {
	return o
}

func (o TokenCertificateNameOutput) ToTokenCertificateNameOutputWithContext(ctx context.Context) TokenCertificateNameOutput {
	return o
}

func (o TokenCertificateNameOutput) ToTokenCertificateNamePtrOutput() TokenCertificateNamePtrOutput {
	return o.ToTokenCertificateNamePtrOutputWithContext(context.Background())
}

func (o TokenCertificateNameOutput) ToTokenCertificateNamePtrOutputWithContext(ctx context.Context) TokenCertificateNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TokenCertificateName) *TokenCertificateName {
		return &v
	}).(TokenCertificateNamePtrOutput)
}

func (o TokenCertificateNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TokenCertificateNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TokenCertificateName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TokenCertificateNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TokenCertificateNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TokenCertificateName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TokenCertificateNamePtrOutput struct{ *pulumi.OutputState }

func (TokenCertificateNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenCertificateName)(nil)).Elem()
}

func (o TokenCertificateNamePtrOutput) ToTokenCertificateNamePtrOutput() TokenCertificateNamePtrOutput {
	return o
}

func (o TokenCertificateNamePtrOutput) ToTokenCertificateNamePtrOutputWithContext(ctx context.Context) TokenCertificateNamePtrOutput {
	return o
}

func (o TokenCertificateNamePtrOutput) Elem() TokenCertificateNameOutput {
	return o.ApplyT(func(v *TokenCertificateName) TokenCertificateName {
		if v != nil {
			return *v
		}
		var ret TokenCertificateName
		return ret
	}).(TokenCertificateNameOutput)
}

func (o TokenCertificateNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TokenCertificateNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TokenCertificateName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TokenCertificateNameInput interface {
	pulumi.Input

	ToTokenCertificateNameOutput() TokenCertificateNameOutput
	ToTokenCertificateNameOutputWithContext(context.Context) TokenCertificateNameOutput
}

var tokenCertificateNamePtrType = reflect.TypeOf((**TokenCertificateName)(nil)).Elem()

type TokenCertificateNamePtrInput interface {
	pulumi.Input

	ToTokenCertificateNamePtrOutput() TokenCertificateNamePtrOutput
	ToTokenCertificateNamePtrOutputWithContext(context.Context) TokenCertificateNamePtrOutput
}

type tokenCertificateNamePtr string

func TokenCertificateNamePtr(v string) TokenCertificateNamePtrInput {
	return (*tokenCertificateNamePtr)(&v)
}

func (*tokenCertificateNamePtr) ElementType() reflect.Type {
	return tokenCertificateNamePtrType
}

func (in *tokenCertificateNamePtr) ToTokenCertificateNamePtrOutput() TokenCertificateNamePtrOutput {
	return pulumi.ToOutput(in).(TokenCertificateNamePtrOutput)
}

func (in *tokenCertificateNamePtr) ToTokenCertificateNamePtrOutputWithContext(ctx context.Context) TokenCertificateNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TokenCertificateNamePtrOutput)
}

type TokenPasswordName string

const (
	TokenPasswordNamePassword1 = TokenPasswordName("password1")
	TokenPasswordNamePassword2 = TokenPasswordName("password2")
)

func (TokenPasswordName) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenPasswordName)(nil)).Elem()
}

func (e TokenPasswordName) ToTokenPasswordNameOutput() TokenPasswordNameOutput {
	return pulumi.ToOutput(e).(TokenPasswordNameOutput)
}

func (e TokenPasswordName) ToTokenPasswordNameOutputWithContext(ctx context.Context) TokenPasswordNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TokenPasswordNameOutput)
}

func (e TokenPasswordName) ToTokenPasswordNamePtrOutput() TokenPasswordNamePtrOutput {
	return e.ToTokenPasswordNamePtrOutputWithContext(context.Background())
}

func (e TokenPasswordName) ToTokenPasswordNamePtrOutputWithContext(ctx context.Context) TokenPasswordNamePtrOutput {
	return TokenPasswordName(e).ToTokenPasswordNameOutputWithContext(ctx).ToTokenPasswordNamePtrOutputWithContext(ctx)
}

func (e TokenPasswordName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TokenPasswordName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TokenPasswordName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TokenPasswordName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TokenPasswordNameOutput struct{ *pulumi.OutputState }

func (TokenPasswordNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenPasswordName)(nil)).Elem()
}

func (o TokenPasswordNameOutput) ToTokenPasswordNameOutput() TokenPasswordNameOutput {
	return o
}

func (o TokenPasswordNameOutput) ToTokenPasswordNameOutputWithContext(ctx context.Context) TokenPasswordNameOutput {
	return o
}

func (o TokenPasswordNameOutput) ToTokenPasswordNamePtrOutput() TokenPasswordNamePtrOutput {
	return o.ToTokenPasswordNamePtrOutputWithContext(context.Background())
}

func (o TokenPasswordNameOutput) ToTokenPasswordNamePtrOutputWithContext(ctx context.Context) TokenPasswordNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TokenPasswordName) *TokenPasswordName {
		return &v
	}).(TokenPasswordNamePtrOutput)
}

func (o TokenPasswordNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TokenPasswordNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TokenPasswordName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TokenPasswordNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TokenPasswordNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TokenPasswordName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TokenPasswordNamePtrOutput struct{ *pulumi.OutputState }

func (TokenPasswordNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenPasswordName)(nil)).Elem()
}

func (o TokenPasswordNamePtrOutput) ToTokenPasswordNamePtrOutput() TokenPasswordNamePtrOutput {
	return o
}

func (o TokenPasswordNamePtrOutput) ToTokenPasswordNamePtrOutputWithContext(ctx context.Context) TokenPasswordNamePtrOutput {
	return o
}

func (o TokenPasswordNamePtrOutput) Elem() TokenPasswordNameOutput {
	return o.ApplyT(func(v *TokenPasswordName) TokenPasswordName {
		if v != nil {
			return *v
		}
		var ret TokenPasswordName
		return ret
	}).(TokenPasswordNameOutput)
}

func (o TokenPasswordNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TokenPasswordNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TokenPasswordName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TokenPasswordNameInput interface {
	pulumi.Input

	ToTokenPasswordNameOutput() TokenPasswordNameOutput
	ToTokenPasswordNameOutputWithContext(context.Context) TokenPasswordNameOutput
}

var tokenPasswordNamePtrType = reflect.TypeOf((**TokenPasswordName)(nil)).Elem()

type TokenPasswordNamePtrInput interface {
	pulumi.Input

	ToTokenPasswordNamePtrOutput() TokenPasswordNamePtrOutput
	ToTokenPasswordNamePtrOutputWithContext(context.Context) TokenPasswordNamePtrOutput
}

type tokenPasswordNamePtr string

func TokenPasswordNamePtr(v string) TokenPasswordNamePtrInput {
	return (*tokenPasswordNamePtr)(&v)
}

func (*tokenPasswordNamePtr) ElementType() reflect.Type {
	return tokenPasswordNamePtrType
}

func (in *tokenPasswordNamePtr) ToTokenPasswordNamePtrOutput() TokenPasswordNamePtrOutput {
	return pulumi.ToOutput(in).(TokenPasswordNamePtrOutput)
}

func (in *tokenPasswordNamePtr) ToTokenPasswordNamePtrOutputWithContext(ctx context.Context) TokenPasswordNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TokenPasswordNamePtrOutput)
}

type TokenStatus string

const (
	TokenStatusEnabled  = TokenStatus("enabled")
	TokenStatusDisabled = TokenStatus("disabled")
)

func (TokenStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenStatus)(nil)).Elem()
}

func (e TokenStatus) ToTokenStatusOutput() TokenStatusOutput {
	return pulumi.ToOutput(e).(TokenStatusOutput)
}

func (e TokenStatus) ToTokenStatusOutputWithContext(ctx context.Context) TokenStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TokenStatusOutput)
}

func (e TokenStatus) ToTokenStatusPtrOutput() TokenStatusPtrOutput {
	return e.ToTokenStatusPtrOutputWithContext(context.Background())
}

func (e TokenStatus) ToTokenStatusPtrOutputWithContext(ctx context.Context) TokenStatusPtrOutput {
	return TokenStatus(e).ToTokenStatusOutputWithContext(ctx).ToTokenStatusPtrOutputWithContext(ctx)
}

func (e TokenStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TokenStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TokenStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TokenStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TokenStatusOutput struct{ *pulumi.OutputState }

func (TokenStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenStatus)(nil)).Elem()
}

func (o TokenStatusOutput) ToTokenStatusOutput() TokenStatusOutput {
	return o
}

func (o TokenStatusOutput) ToTokenStatusOutputWithContext(ctx context.Context) TokenStatusOutput {
	return o
}

func (o TokenStatusOutput) ToTokenStatusPtrOutput() TokenStatusPtrOutput {
	return o.ToTokenStatusPtrOutputWithContext(context.Background())
}

func (o TokenStatusOutput) ToTokenStatusPtrOutputWithContext(ctx context.Context) TokenStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TokenStatus) *TokenStatus {
		return &v
	}).(TokenStatusPtrOutput)
}

func (o TokenStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TokenStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TokenStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TokenStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TokenStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TokenStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TokenStatusPtrOutput struct{ *pulumi.OutputState }

func (TokenStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenStatus)(nil)).Elem()
}

func (o TokenStatusPtrOutput) ToTokenStatusPtrOutput() TokenStatusPtrOutput {
	return o
}

func (o TokenStatusPtrOutput) ToTokenStatusPtrOutputWithContext(ctx context.Context) TokenStatusPtrOutput {
	return o
}

func (o TokenStatusPtrOutput) Elem() TokenStatusOutput {
	return o.ApplyT(func(v *TokenStatus) TokenStatus {
		if v != nil {
			return *v
		}
		var ret TokenStatus
		return ret
	}).(TokenStatusOutput)
}

func (o TokenStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TokenStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TokenStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TokenStatusInput interface {
	pulumi.Input

	ToTokenStatusOutput() TokenStatusOutput
	ToTokenStatusOutputWithContext(context.Context) TokenStatusOutput
}

var tokenStatusPtrType = reflect.TypeOf((**TokenStatus)(nil)).Elem()

type TokenStatusPtrInput interface {
	pulumi.Input

	ToTokenStatusPtrOutput() TokenStatusPtrOutput
	ToTokenStatusPtrOutputWithContext(context.Context) TokenStatusPtrOutput
}

type tokenStatusPtr string

func TokenStatusPtr(v string) TokenStatusPtrInput {
	return (*tokenStatusPtr)(&v)
}

func (*tokenStatusPtr) ElementType() reflect.Type {
	return tokenStatusPtrType
}

func (in *tokenStatusPtr) ToTokenStatusPtrOutput() TokenStatusPtrOutput {
	return pulumi.ToOutput(in).(TokenStatusPtrOutput)
}

func (in *tokenStatusPtr) ToTokenStatusPtrOutputWithContext(ctx context.Context) TokenStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TokenStatusPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TokenCertificateNameInput)(nil)).Elem(), TokenCertificateName("certificate1"))
	pulumi.RegisterInputType(reflect.TypeOf((*TokenCertificateNamePtrInput)(nil)).Elem(), TokenCertificateName("certificate1"))
	pulumi.RegisterInputType(reflect.TypeOf((*TokenPasswordNameInput)(nil)).Elem(), TokenPasswordName("password1"))
	pulumi.RegisterInputType(reflect.TypeOf((*TokenPasswordNamePtrInput)(nil)).Elem(), TokenPasswordName("password1"))
	pulumi.RegisterInputType(reflect.TypeOf((*TokenStatusInput)(nil)).Elem(), TokenStatus("enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*TokenStatusPtrInput)(nil)).Elem(), TokenStatus("enabled"))
	pulumi.RegisterOutputType(TokenCertificateNameOutput{})
	pulumi.RegisterOutputType(TokenCertificateNamePtrOutput{})
	pulumi.RegisterOutputType(TokenPasswordNameOutput{})
	pulumi.RegisterOutputType(TokenPasswordNamePtrOutput{})
	pulumi.RegisterOutputType(TokenStatusOutput{})
	pulumi.RegisterOutputType(TokenStatusPtrOutput{})
}

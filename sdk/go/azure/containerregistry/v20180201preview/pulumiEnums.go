


package v20180201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BuildTaskStatus string

const (
	BuildTaskStatusDisabled = BuildTaskStatus("Disabled")
	BuildTaskStatusEnabled  = BuildTaskStatus("Enabled")
)

func (BuildTaskStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildTaskStatus)(nil)).Elem()
}

func (e BuildTaskStatus) ToBuildTaskStatusOutput() BuildTaskStatusOutput {
	return pulumi.ToOutput(e).(BuildTaskStatusOutput)
}

func (e BuildTaskStatus) ToBuildTaskStatusOutputWithContext(ctx context.Context) BuildTaskStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BuildTaskStatusOutput)
}

func (e BuildTaskStatus) ToBuildTaskStatusPtrOutput() BuildTaskStatusPtrOutput {
	return e.ToBuildTaskStatusPtrOutputWithContext(context.Background())
}

func (e BuildTaskStatus) ToBuildTaskStatusPtrOutputWithContext(ctx context.Context) BuildTaskStatusPtrOutput {
	return BuildTaskStatus(e).ToBuildTaskStatusOutputWithContext(ctx).ToBuildTaskStatusPtrOutputWithContext(ctx)
}

func (e BuildTaskStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BuildTaskStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BuildTaskStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BuildTaskStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BuildTaskStatusOutput struct{ *pulumi.OutputState }

func (BuildTaskStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildTaskStatus)(nil)).Elem()
}

func (o BuildTaskStatusOutput) ToBuildTaskStatusOutput() BuildTaskStatusOutput {
	return o
}

func (o BuildTaskStatusOutput) ToBuildTaskStatusOutputWithContext(ctx context.Context) BuildTaskStatusOutput {
	return o
}

func (o BuildTaskStatusOutput) ToBuildTaskStatusPtrOutput() BuildTaskStatusPtrOutput {
	return o.ToBuildTaskStatusPtrOutputWithContext(context.Background())
}

func (o BuildTaskStatusOutput) ToBuildTaskStatusPtrOutputWithContext(ctx context.Context) BuildTaskStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BuildTaskStatus) *BuildTaskStatus {
		return &v
	}).(BuildTaskStatusPtrOutput)
}

func (o BuildTaskStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BuildTaskStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BuildTaskStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BuildTaskStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BuildTaskStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BuildTaskStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BuildTaskStatusPtrOutput struct{ *pulumi.OutputState }

func (BuildTaskStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildTaskStatus)(nil)).Elem()
}

func (o BuildTaskStatusPtrOutput) ToBuildTaskStatusPtrOutput() BuildTaskStatusPtrOutput {
	return o
}

func (o BuildTaskStatusPtrOutput) ToBuildTaskStatusPtrOutputWithContext(ctx context.Context) BuildTaskStatusPtrOutput {
	return o
}

func (o BuildTaskStatusPtrOutput) Elem() BuildTaskStatusOutput {
	return o.ApplyT(func(v *BuildTaskStatus) BuildTaskStatus {
		if v != nil {
			return *v
		}
		var ret BuildTaskStatus
		return ret
	}).(BuildTaskStatusOutput)
}

func (o BuildTaskStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BuildTaskStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BuildTaskStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BuildTaskStatusInput interface {
	pulumi.Input

	ToBuildTaskStatusOutput() BuildTaskStatusOutput
	ToBuildTaskStatusOutputWithContext(context.Context) BuildTaskStatusOutput
}

var buildTaskStatusPtrType = reflect.TypeOf((**BuildTaskStatus)(nil)).Elem()

type BuildTaskStatusPtrInput interface {
	pulumi.Input

	ToBuildTaskStatusPtrOutput() BuildTaskStatusPtrOutput
	ToBuildTaskStatusPtrOutputWithContext(context.Context) BuildTaskStatusPtrOutput
}

type buildTaskStatusPtr string

func BuildTaskStatusPtr(v string) BuildTaskStatusPtrInput {
	return (*buildTaskStatusPtr)(&v)
}

func (*buildTaskStatusPtr) ElementType() reflect.Type {
	return buildTaskStatusPtrType
}

func (in *buildTaskStatusPtr) ToBuildTaskStatusPtrOutput() BuildTaskStatusPtrOutput {
	return pulumi.ToOutput(in).(BuildTaskStatusPtrOutput)
}

func (in *buildTaskStatusPtr) ToBuildTaskStatusPtrOutputWithContext(ctx context.Context) BuildTaskStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BuildTaskStatusPtrOutput)
}

type OsType string

const (
	OsTypeWindows = OsType("Windows")
	OsTypeLinux   = OsType("Linux")
)

func (OsType) ElementType() reflect.Type {
	return reflect.TypeOf((*OsType)(nil)).Elem()
}

func (e OsType) ToOsTypeOutput() OsTypeOutput {
	return pulumi.ToOutput(e).(OsTypeOutput)
}

func (e OsType) ToOsTypeOutputWithContext(ctx context.Context) OsTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OsTypeOutput)
}

func (e OsType) ToOsTypePtrOutput() OsTypePtrOutput {
	return e.ToOsTypePtrOutputWithContext(context.Background())
}

func (e OsType) ToOsTypePtrOutputWithContext(ctx context.Context) OsTypePtrOutput {
	return OsType(e).ToOsTypeOutputWithContext(ctx).ToOsTypePtrOutputWithContext(ctx)
}

func (e OsType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OsType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OsType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OsType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OsTypeOutput struct{ *pulumi.OutputState }

func (OsTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OsType)(nil)).Elem()
}

func (o OsTypeOutput) ToOsTypeOutput() OsTypeOutput {
	return o
}

func (o OsTypeOutput) ToOsTypeOutputWithContext(ctx context.Context) OsTypeOutput {
	return o
}

func (o OsTypeOutput) ToOsTypePtrOutput() OsTypePtrOutput {
	return o.ToOsTypePtrOutputWithContext(context.Background())
}

func (o OsTypeOutput) ToOsTypePtrOutputWithContext(ctx context.Context) OsTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OsType) *OsType {
		return &v
	}).(OsTypePtrOutput)
}

func (o OsTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OsTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OsType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OsTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OsTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OsType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OsTypePtrOutput struct{ *pulumi.OutputState }

func (OsTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OsType)(nil)).Elem()
}

func (o OsTypePtrOutput) ToOsTypePtrOutput() OsTypePtrOutput {
	return o
}

func (o OsTypePtrOutput) ToOsTypePtrOutputWithContext(ctx context.Context) OsTypePtrOutput {
	return o
}

func (o OsTypePtrOutput) Elem() OsTypeOutput {
	return o.ApplyT(func(v *OsType) OsType {
		if v != nil {
			return *v
		}
		var ret OsType
		return ret
	}).(OsTypeOutput)
}

func (o OsTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OsTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OsType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OsTypeInput interface {
	pulumi.Input

	ToOsTypeOutput() OsTypeOutput
	ToOsTypeOutputWithContext(context.Context) OsTypeOutput
}

var osTypePtrType = reflect.TypeOf((**OsType)(nil)).Elem()

type OsTypePtrInput interface {
	pulumi.Input

	ToOsTypePtrOutput() OsTypePtrOutput
	ToOsTypePtrOutputWithContext(context.Context) OsTypePtrOutput
}

type osTypePtr string

func OsTypePtr(v string) OsTypePtrInput {
	return (*osTypePtr)(&v)
}

func (*osTypePtr) ElementType() reflect.Type {
	return osTypePtrType
}

func (in *osTypePtr) ToOsTypePtrOutput() OsTypePtrOutput {
	return pulumi.ToOutput(in).(OsTypePtrOutput)
}

func (in *osTypePtr) ToOsTypePtrOutputWithContext(ctx context.Context) OsTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OsTypePtrOutput)
}

type SourceControlType string

const (
	SourceControlTypeGithub                  = SourceControlType("Github")
	SourceControlTypeVisualStudioTeamService = SourceControlType("VisualStudioTeamService")
)

func (SourceControlType) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceControlType)(nil)).Elem()
}

func (e SourceControlType) ToSourceControlTypeOutput() SourceControlTypeOutput {
	return pulumi.ToOutput(e).(SourceControlTypeOutput)
}

func (e SourceControlType) ToSourceControlTypeOutputWithContext(ctx context.Context) SourceControlTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SourceControlTypeOutput)
}

func (e SourceControlType) ToSourceControlTypePtrOutput() SourceControlTypePtrOutput {
	return e.ToSourceControlTypePtrOutputWithContext(context.Background())
}

func (e SourceControlType) ToSourceControlTypePtrOutputWithContext(ctx context.Context) SourceControlTypePtrOutput {
	return SourceControlType(e).ToSourceControlTypeOutputWithContext(ctx).ToSourceControlTypePtrOutputWithContext(ctx)
}

func (e SourceControlType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceControlType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceControlType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SourceControlType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SourceControlTypeOutput struct{ *pulumi.OutputState }

func (SourceControlTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceControlType)(nil)).Elem()
}

func (o SourceControlTypeOutput) ToSourceControlTypeOutput() SourceControlTypeOutput {
	return o
}

func (o SourceControlTypeOutput) ToSourceControlTypeOutputWithContext(ctx context.Context) SourceControlTypeOutput {
	return o
}

func (o SourceControlTypeOutput) ToSourceControlTypePtrOutput() SourceControlTypePtrOutput {
	return o.ToSourceControlTypePtrOutputWithContext(context.Background())
}

func (o SourceControlTypeOutput) ToSourceControlTypePtrOutputWithContext(ctx context.Context) SourceControlTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceControlType) *SourceControlType {
		return &v
	}).(SourceControlTypePtrOutput)
}

func (o SourceControlTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SourceControlTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceControlType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SourceControlTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceControlTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceControlType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SourceControlTypePtrOutput struct{ *pulumi.OutputState }

func (SourceControlTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceControlType)(nil)).Elem()
}

func (o SourceControlTypePtrOutput) ToSourceControlTypePtrOutput() SourceControlTypePtrOutput {
	return o
}

func (o SourceControlTypePtrOutput) ToSourceControlTypePtrOutputWithContext(ctx context.Context) SourceControlTypePtrOutput {
	return o
}

func (o SourceControlTypePtrOutput) Elem() SourceControlTypeOutput {
	return o.ApplyT(func(v *SourceControlType) SourceControlType {
		if v != nil {
			return *v
		}
		var ret SourceControlType
		return ret
	}).(SourceControlTypeOutput)
}

func (o SourceControlTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceControlTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SourceControlType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SourceControlTypeInput interface {
	pulumi.Input

	ToSourceControlTypeOutput() SourceControlTypeOutput
	ToSourceControlTypeOutputWithContext(context.Context) SourceControlTypeOutput
}

var sourceControlTypePtrType = reflect.TypeOf((**SourceControlType)(nil)).Elem()

type SourceControlTypePtrInput interface {
	pulumi.Input

	ToSourceControlTypePtrOutput() SourceControlTypePtrOutput
	ToSourceControlTypePtrOutputWithContext(context.Context) SourceControlTypePtrOutput
}

type sourceControlTypePtr string

func SourceControlTypePtr(v string) SourceControlTypePtrInput {
	return (*sourceControlTypePtr)(&v)
}

func (*sourceControlTypePtr) ElementType() reflect.Type {
	return sourceControlTypePtrType
}

func (in *sourceControlTypePtr) ToSourceControlTypePtrOutput() SourceControlTypePtrOutput {
	return pulumi.ToOutput(in).(SourceControlTypePtrOutput)
}

func (in *sourceControlTypePtr) ToSourceControlTypePtrOutputWithContext(ctx context.Context) SourceControlTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SourceControlTypePtrOutput)
}

type TokenType string

const (
	TokenTypePAT   = TokenType("PAT")
	TokenTypeOAuth = TokenType("OAuth")
)

func (TokenType) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenType)(nil)).Elem()
}

func (e TokenType) ToTokenTypeOutput() TokenTypeOutput {
	return pulumi.ToOutput(e).(TokenTypeOutput)
}

func (e TokenType) ToTokenTypeOutputWithContext(ctx context.Context) TokenTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TokenTypeOutput)
}

func (e TokenType) ToTokenTypePtrOutput() TokenTypePtrOutput {
	return e.ToTokenTypePtrOutputWithContext(context.Background())
}

func (e TokenType) ToTokenTypePtrOutputWithContext(ctx context.Context) TokenTypePtrOutput {
	return TokenType(e).ToTokenTypeOutputWithContext(ctx).ToTokenTypePtrOutputWithContext(ctx)
}

func (e TokenType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TokenType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TokenType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TokenType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TokenTypeOutput struct{ *pulumi.OutputState }

func (TokenTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenType)(nil)).Elem()
}

func (o TokenTypeOutput) ToTokenTypeOutput() TokenTypeOutput {
	return o
}

func (o TokenTypeOutput) ToTokenTypeOutputWithContext(ctx context.Context) TokenTypeOutput {
	return o
}

func (o TokenTypeOutput) ToTokenTypePtrOutput() TokenTypePtrOutput {
	return o.ToTokenTypePtrOutputWithContext(context.Background())
}

func (o TokenTypeOutput) ToTokenTypePtrOutputWithContext(ctx context.Context) TokenTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TokenType) *TokenType {
		return &v
	}).(TokenTypePtrOutput)
}

func (o TokenTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TokenTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TokenType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TokenTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TokenTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TokenType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TokenTypePtrOutput struct{ *pulumi.OutputState }

func (TokenTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenType)(nil)).Elem()
}

func (o TokenTypePtrOutput) ToTokenTypePtrOutput() TokenTypePtrOutput {
	return o
}

func (o TokenTypePtrOutput) ToTokenTypePtrOutputWithContext(ctx context.Context) TokenTypePtrOutput {
	return o
}

func (o TokenTypePtrOutput) Elem() TokenTypeOutput {
	return o.ApplyT(func(v *TokenType) TokenType {
		if v != nil {
			return *v
		}
		var ret TokenType
		return ret
	}).(TokenTypeOutput)
}

func (o TokenTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TokenTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TokenType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TokenTypeInput interface {
	pulumi.Input

	ToTokenTypeOutput() TokenTypeOutput
	ToTokenTypeOutputWithContext(context.Context) TokenTypeOutput
}

var tokenTypePtrType = reflect.TypeOf((**TokenType)(nil)).Elem()

type TokenTypePtrInput interface {
	pulumi.Input

	ToTokenTypePtrOutput() TokenTypePtrOutput
	ToTokenTypePtrOutputWithContext(context.Context) TokenTypePtrOutput
}

type tokenTypePtr string

func TokenTypePtr(v string) TokenTypePtrInput {
	return (*tokenTypePtr)(&v)
}

func (*tokenTypePtr) ElementType() reflect.Type {
	return tokenTypePtrType
}

func (in *tokenTypePtr) ToTokenTypePtrOutput() TokenTypePtrOutput {
	return pulumi.ToOutput(in).(TokenTypePtrOutput)
}

func (in *tokenTypePtr) ToTokenTypePtrOutputWithContext(ctx context.Context) TokenTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TokenTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BuildTaskStatusOutput{})
	pulumi.RegisterOutputType(BuildTaskStatusPtrOutput{})
	pulumi.RegisterOutputType(OsTypeOutput{})
	pulumi.RegisterOutputType(OsTypePtrOutput{})
	pulumi.RegisterOutputType(SourceControlTypeOutput{})
	pulumi.RegisterOutputType(SourceControlTypePtrOutput{})
	pulumi.RegisterOutputType(TokenTypeOutput{})
	pulumi.RegisterOutputType(TokenTypePtrOutput{})
}

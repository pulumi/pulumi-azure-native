


package confidentialledger

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LedgerRoleName string

const (
	LedgerRoleNameReader        = LedgerRoleName("Reader")
	LedgerRoleNameContributor   = LedgerRoleName("Contributor")
	LedgerRoleNameAdministrator = LedgerRoleName("Administrator")
)

func (LedgerRoleName) ElementType() reflect.Type {
	return reflect.TypeOf((*LedgerRoleName)(nil)).Elem()
}

func (e LedgerRoleName) ToLedgerRoleNameOutput() LedgerRoleNameOutput {
	return pulumi.ToOutput(e).(LedgerRoleNameOutput)
}

func (e LedgerRoleName) ToLedgerRoleNameOutputWithContext(ctx context.Context) LedgerRoleNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LedgerRoleNameOutput)
}

func (e LedgerRoleName) ToLedgerRoleNamePtrOutput() LedgerRoleNamePtrOutput {
	return e.ToLedgerRoleNamePtrOutputWithContext(context.Background())
}

func (e LedgerRoleName) ToLedgerRoleNamePtrOutputWithContext(ctx context.Context) LedgerRoleNamePtrOutput {
	return LedgerRoleName(e).ToLedgerRoleNameOutputWithContext(ctx).ToLedgerRoleNamePtrOutputWithContext(ctx)
}

func (e LedgerRoleName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LedgerRoleName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LedgerRoleName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LedgerRoleName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LedgerRoleNameOutput struct{ *pulumi.OutputState }

func (LedgerRoleNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LedgerRoleName)(nil)).Elem()
}

func (o LedgerRoleNameOutput) ToLedgerRoleNameOutput() LedgerRoleNameOutput {
	return o
}

func (o LedgerRoleNameOutput) ToLedgerRoleNameOutputWithContext(ctx context.Context) LedgerRoleNameOutput {
	return o
}

func (o LedgerRoleNameOutput) ToLedgerRoleNamePtrOutput() LedgerRoleNamePtrOutput {
	return o.ToLedgerRoleNamePtrOutputWithContext(context.Background())
}

func (o LedgerRoleNameOutput) ToLedgerRoleNamePtrOutputWithContext(ctx context.Context) LedgerRoleNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LedgerRoleName) *LedgerRoleName {
		return &v
	}).(LedgerRoleNamePtrOutput)
}

func (o LedgerRoleNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LedgerRoleNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LedgerRoleName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LedgerRoleNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LedgerRoleNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LedgerRoleName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LedgerRoleNamePtrOutput struct{ *pulumi.OutputState }

func (LedgerRoleNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LedgerRoleName)(nil)).Elem()
}

func (o LedgerRoleNamePtrOutput) ToLedgerRoleNamePtrOutput() LedgerRoleNamePtrOutput {
	return o
}

func (o LedgerRoleNamePtrOutput) ToLedgerRoleNamePtrOutputWithContext(ctx context.Context) LedgerRoleNamePtrOutput {
	return o
}

func (o LedgerRoleNamePtrOutput) Elem() LedgerRoleNameOutput {
	return o.ApplyT(func(v *LedgerRoleName) LedgerRoleName {
		if v != nil {
			return *v
		}
		var ret LedgerRoleName
		return ret
	}).(LedgerRoleNameOutput)
}

func (o LedgerRoleNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LedgerRoleNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LedgerRoleName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LedgerRoleNameInput interface {
	pulumi.Input

	ToLedgerRoleNameOutput() LedgerRoleNameOutput
	ToLedgerRoleNameOutputWithContext(context.Context) LedgerRoleNameOutput
}

var ledgerRoleNamePtrType = reflect.TypeOf((**LedgerRoleName)(nil)).Elem()

type LedgerRoleNamePtrInput interface {
	pulumi.Input

	ToLedgerRoleNamePtrOutput() LedgerRoleNamePtrOutput
	ToLedgerRoleNamePtrOutputWithContext(context.Context) LedgerRoleNamePtrOutput
}

type ledgerRoleNamePtr string

func LedgerRoleNamePtr(v string) LedgerRoleNamePtrInput {
	return (*ledgerRoleNamePtr)(&v)
}

func (*ledgerRoleNamePtr) ElementType() reflect.Type {
	return ledgerRoleNamePtrType
}

func (in *ledgerRoleNamePtr) ToLedgerRoleNamePtrOutput() LedgerRoleNamePtrOutput {
	return pulumi.ToOutput(in).(LedgerRoleNamePtrOutput)
}

func (in *ledgerRoleNamePtr) ToLedgerRoleNamePtrOutputWithContext(ctx context.Context) LedgerRoleNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LedgerRoleNamePtrOutput)
}

type LedgerType string

const (
	LedgerTypeUnknown = LedgerType("Unknown")
	LedgerTypePublic  = LedgerType("Public")
	LedgerTypePrivate = LedgerType("Private")
)

func (LedgerType) ElementType() reflect.Type {
	return reflect.TypeOf((*LedgerType)(nil)).Elem()
}

func (e LedgerType) ToLedgerTypeOutput() LedgerTypeOutput {
	return pulumi.ToOutput(e).(LedgerTypeOutput)
}

func (e LedgerType) ToLedgerTypeOutputWithContext(ctx context.Context) LedgerTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LedgerTypeOutput)
}

func (e LedgerType) ToLedgerTypePtrOutput() LedgerTypePtrOutput {
	return e.ToLedgerTypePtrOutputWithContext(context.Background())
}

func (e LedgerType) ToLedgerTypePtrOutputWithContext(ctx context.Context) LedgerTypePtrOutput {
	return LedgerType(e).ToLedgerTypeOutputWithContext(ctx).ToLedgerTypePtrOutputWithContext(ctx)
}

func (e LedgerType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LedgerType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LedgerType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LedgerType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LedgerTypeOutput struct{ *pulumi.OutputState }

func (LedgerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LedgerType)(nil)).Elem()
}

func (o LedgerTypeOutput) ToLedgerTypeOutput() LedgerTypeOutput {
	return o
}

func (o LedgerTypeOutput) ToLedgerTypeOutputWithContext(ctx context.Context) LedgerTypeOutput {
	return o
}

func (o LedgerTypeOutput) ToLedgerTypePtrOutput() LedgerTypePtrOutput {
	return o.ToLedgerTypePtrOutputWithContext(context.Background())
}

func (o LedgerTypeOutput) ToLedgerTypePtrOutputWithContext(ctx context.Context) LedgerTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LedgerType) *LedgerType {
		return &v
	}).(LedgerTypePtrOutput)
}

func (o LedgerTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LedgerTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LedgerType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LedgerTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LedgerTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LedgerType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LedgerTypePtrOutput struct{ *pulumi.OutputState }

func (LedgerTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LedgerType)(nil)).Elem()
}

func (o LedgerTypePtrOutput) ToLedgerTypePtrOutput() LedgerTypePtrOutput {
	return o
}

func (o LedgerTypePtrOutput) ToLedgerTypePtrOutputWithContext(ctx context.Context) LedgerTypePtrOutput {
	return o
}

func (o LedgerTypePtrOutput) Elem() LedgerTypeOutput {
	return o.ApplyT(func(v *LedgerType) LedgerType {
		if v != nil {
			return *v
		}
		var ret LedgerType
		return ret
	}).(LedgerTypeOutput)
}

func (o LedgerTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LedgerTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LedgerType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LedgerTypeInput interface {
	pulumi.Input

	ToLedgerTypeOutput() LedgerTypeOutput
	ToLedgerTypeOutputWithContext(context.Context) LedgerTypeOutput
}

var ledgerTypePtrType = reflect.TypeOf((**LedgerType)(nil)).Elem()

type LedgerTypePtrInput interface {
	pulumi.Input

	ToLedgerTypePtrOutput() LedgerTypePtrOutput
	ToLedgerTypePtrOutputWithContext(context.Context) LedgerTypePtrOutput
}

type ledgerTypePtr string

func LedgerTypePtr(v string) LedgerTypePtrInput {
	return (*ledgerTypePtr)(&v)
}

func (*ledgerTypePtr) ElementType() reflect.Type {
	return ledgerTypePtrType
}

func (in *ledgerTypePtr) ToLedgerTypePtrOutput() LedgerTypePtrOutput {
	return pulumi.ToOutput(in).(LedgerTypePtrOutput)
}

func (in *ledgerTypePtr) ToLedgerTypePtrOutputWithContext(ctx context.Context) LedgerTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LedgerTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LedgerRoleNameInput)(nil)).Elem(), LedgerRoleName("Reader"))
	pulumi.RegisterInputType(reflect.TypeOf((*LedgerRoleNamePtrInput)(nil)).Elem(), LedgerRoleName("Reader"))
	pulumi.RegisterInputType(reflect.TypeOf((*LedgerTypeInput)(nil)).Elem(), LedgerType("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*LedgerTypePtrInput)(nil)).Elem(), LedgerType("Unknown"))
	pulumi.RegisterOutputType(LedgerRoleNameOutput{})
	pulumi.RegisterOutputType(LedgerRoleNamePtrOutput{})
	pulumi.RegisterOutputType(LedgerTypeOutput{})
	pulumi.RegisterOutputType(LedgerTypePtrOutput{})
}

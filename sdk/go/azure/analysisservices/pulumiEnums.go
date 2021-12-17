


package analysisservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectionMode string

const (
	ConnectionModeAll      = ConnectionMode("All")
	ConnectionModeReadOnly = ConnectionMode("ReadOnly")
)

func (ConnectionMode) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionMode)(nil)).Elem()
}

func (e ConnectionMode) ToConnectionModeOutput() ConnectionModeOutput {
	return pulumi.ToOutput(e).(ConnectionModeOutput)
}

func (e ConnectionMode) ToConnectionModeOutputWithContext(ctx context.Context) ConnectionModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConnectionModeOutput)
}

func (e ConnectionMode) ToConnectionModePtrOutput() ConnectionModePtrOutput {
	return e.ToConnectionModePtrOutputWithContext(context.Background())
}

func (e ConnectionMode) ToConnectionModePtrOutputWithContext(ctx context.Context) ConnectionModePtrOutput {
	return ConnectionMode(e).ToConnectionModeOutputWithContext(ctx).ToConnectionModePtrOutputWithContext(ctx)
}

func (e ConnectionMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectionMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectionMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConnectionMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConnectionModeOutput struct{ *pulumi.OutputState }

func (ConnectionModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionMode)(nil)).Elem()
}

func (o ConnectionModeOutput) ToConnectionModeOutput() ConnectionModeOutput {
	return o
}

func (o ConnectionModeOutput) ToConnectionModeOutputWithContext(ctx context.Context) ConnectionModeOutput {
	return o
}

func (o ConnectionModeOutput) ToConnectionModePtrOutput() ConnectionModePtrOutput {
	return o.ToConnectionModePtrOutputWithContext(context.Background())
}

func (o ConnectionModeOutput) ToConnectionModePtrOutputWithContext(ctx context.Context) ConnectionModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionMode) *ConnectionMode {
		return &v
	}).(ConnectionModePtrOutput)
}

func (o ConnectionModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConnectionModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConnectionModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConnectionModePtrOutput struct{ *pulumi.OutputState }

func (ConnectionModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionMode)(nil)).Elem()
}

func (o ConnectionModePtrOutput) ToConnectionModePtrOutput() ConnectionModePtrOutput {
	return o
}

func (o ConnectionModePtrOutput) ToConnectionModePtrOutputWithContext(ctx context.Context) ConnectionModePtrOutput {
	return o
}

func (o ConnectionModePtrOutput) Elem() ConnectionModeOutput {
	return o.ApplyT(func(v *ConnectionMode) ConnectionMode {
		if v != nil {
			return *v
		}
		var ret ConnectionMode
		return ret
	}).(ConnectionModeOutput)
}

func (o ConnectionModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConnectionMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConnectionModeInput interface {
	pulumi.Input

	ToConnectionModeOutput() ConnectionModeOutput
	ToConnectionModeOutputWithContext(context.Context) ConnectionModeOutput
}

var connectionModePtrType = reflect.TypeOf((**ConnectionMode)(nil)).Elem()

type ConnectionModePtrInput interface {
	pulumi.Input

	ToConnectionModePtrOutput() ConnectionModePtrOutput
	ToConnectionModePtrOutputWithContext(context.Context) ConnectionModePtrOutput
}

type connectionModePtr string

func ConnectionModePtr(v string) ConnectionModePtrInput {
	return (*connectionModePtr)(&v)
}

func (*connectionModePtr) ElementType() reflect.Type {
	return connectionModePtrType
}

func (in *connectionModePtr) ToConnectionModePtrOutput() ConnectionModePtrOutput {
	return pulumi.ToOutput(in).(ConnectionModePtrOutput)
}

func (in *connectionModePtr) ToConnectionModePtrOutputWithContext(ctx context.Context) ConnectionModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConnectionModePtrOutput)
}

type SkuTier string

const (
	SkuTierDevelopment = SkuTier("Development")
	SkuTierBasic       = SkuTier("Basic")
	SkuTierStandard    = SkuTier("Standard")
)

func init() {
	pulumi.RegisterOutputType(ConnectionModeOutput{})
	pulumi.RegisterOutputType(ConnectionModePtrOutput{})
}

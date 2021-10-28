


package storagepool

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DiskPoolTier string

const (
	DiskPoolTierBasic    = DiskPoolTier("Basic")
	DiskPoolTierStandard = DiskPoolTier("Standard")
	DiskPoolTierPremium  = DiskPoolTier("Premium")
)

func (DiskPoolTier) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskPoolTier)(nil)).Elem()
}

func (e DiskPoolTier) ToDiskPoolTierOutput() DiskPoolTierOutput {
	return pulumi.ToOutput(e).(DiskPoolTierOutput)
}

func (e DiskPoolTier) ToDiskPoolTierOutputWithContext(ctx context.Context) DiskPoolTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DiskPoolTierOutput)
}

func (e DiskPoolTier) ToDiskPoolTierPtrOutput() DiskPoolTierPtrOutput {
	return e.ToDiskPoolTierPtrOutputWithContext(context.Background())
}

func (e DiskPoolTier) ToDiskPoolTierPtrOutputWithContext(ctx context.Context) DiskPoolTierPtrOutput {
	return DiskPoolTier(e).ToDiskPoolTierOutputWithContext(ctx).ToDiskPoolTierPtrOutputWithContext(ctx)
}

func (e DiskPoolTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskPoolTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskPoolTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DiskPoolTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DiskPoolTierOutput struct{ *pulumi.OutputState }

func (DiskPoolTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskPoolTier)(nil)).Elem()
}

func (o DiskPoolTierOutput) ToDiskPoolTierOutput() DiskPoolTierOutput {
	return o
}

func (o DiskPoolTierOutput) ToDiskPoolTierOutputWithContext(ctx context.Context) DiskPoolTierOutput {
	return o
}

func (o DiskPoolTierOutput) ToDiskPoolTierPtrOutput() DiskPoolTierPtrOutput {
	return o.ToDiskPoolTierPtrOutputWithContext(context.Background())
}

func (o DiskPoolTierOutput) ToDiskPoolTierPtrOutputWithContext(ctx context.Context) DiskPoolTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskPoolTier) *DiskPoolTier {
		return &v
	}).(DiskPoolTierPtrOutput)
}

func (o DiskPoolTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DiskPoolTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskPoolTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DiskPoolTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskPoolTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskPoolTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DiskPoolTierPtrOutput struct{ *pulumi.OutputState }

func (DiskPoolTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskPoolTier)(nil)).Elem()
}

func (o DiskPoolTierPtrOutput) ToDiskPoolTierPtrOutput() DiskPoolTierPtrOutput {
	return o
}

func (o DiskPoolTierPtrOutput) ToDiskPoolTierPtrOutputWithContext(ctx context.Context) DiskPoolTierPtrOutput {
	return o
}

func (o DiskPoolTierPtrOutput) Elem() DiskPoolTierOutput {
	return o.ApplyT(func(v *DiskPoolTier) DiskPoolTier {
		if v != nil {
			return *v
		}
		var ret DiskPoolTier
		return ret
	}).(DiskPoolTierOutput)
}

func (o DiskPoolTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskPoolTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DiskPoolTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DiskPoolTierInput interface {
	pulumi.Input

	ToDiskPoolTierOutput() DiskPoolTierOutput
	ToDiskPoolTierOutputWithContext(context.Context) DiskPoolTierOutput
}

var diskPoolTierPtrType = reflect.TypeOf((**DiskPoolTier)(nil)).Elem()

type DiskPoolTierPtrInput interface {
	pulumi.Input

	ToDiskPoolTierPtrOutput() DiskPoolTierPtrOutput
	ToDiskPoolTierPtrOutputWithContext(context.Context) DiskPoolTierPtrOutput
}

type diskPoolTierPtr string

func DiskPoolTierPtr(v string) DiskPoolTierPtrInput {
	return (*diskPoolTierPtr)(&v)
}

func (*diskPoolTierPtr) ElementType() reflect.Type {
	return diskPoolTierPtrType
}

func (in *diskPoolTierPtr) ToDiskPoolTierPtrOutput() DiskPoolTierPtrOutput {
	return pulumi.ToOutput(in).(DiskPoolTierPtrOutput)
}

func (in *diskPoolTierPtr) ToDiskPoolTierPtrOutputWithContext(ctx context.Context) DiskPoolTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DiskPoolTierPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DiskPoolTierInput)(nil)).Elem(), DiskPoolTier("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*DiskPoolTierPtrInput)(nil)).Elem(), DiskPoolTier("Basic"))
	pulumi.RegisterOutputType(DiskPoolTierOutput{})
	pulumi.RegisterOutputType(DiskPoolTierPtrOutput{})
}

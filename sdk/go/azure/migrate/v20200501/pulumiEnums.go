


package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Status string

const (
	StatusApproved     = Status("Approved")
	StatusPending      = Status("Pending")
	StatusRejected     = Status("Rejected")
	StatusDisconnected = Status("Disconnected")
)

func (Status) ElementType() reflect.Type {
	return reflect.TypeOf((*Status)(nil)).Elem()
}

func (e Status) ToStatusOutput() StatusOutput {
	return pulumi.ToOutput(e).(StatusOutput)
}

func (e Status) ToStatusOutputWithContext(ctx context.Context) StatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StatusOutput)
}

func (e Status) ToStatusPtrOutput() StatusPtrOutput {
	return e.ToStatusPtrOutputWithContext(context.Background())
}

func (e Status) ToStatusPtrOutputWithContext(ctx context.Context) StatusPtrOutput {
	return Status(e).ToStatusOutputWithContext(ctx).ToStatusPtrOutputWithContext(ctx)
}

func (e Status) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Status) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Status) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Status) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StatusOutput struct{ *pulumi.OutputState }

func (StatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Status)(nil)).Elem()
}

func (o StatusOutput) ToStatusOutput() StatusOutput {
	return o
}

func (o StatusOutput) ToStatusOutputWithContext(ctx context.Context) StatusOutput {
	return o
}

func (o StatusOutput) ToStatusPtrOutput() StatusPtrOutput {
	return o.ToStatusPtrOutputWithContext(context.Background())
}

func (o StatusOutput) ToStatusPtrOutputWithContext(ctx context.Context) StatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Status) *Status {
		return &v
	}).(StatusPtrOutput)
}

func (o StatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Status) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Status) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StatusPtrOutput struct{ *pulumi.OutputState }

func (StatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Status)(nil)).Elem()
}

func (o StatusPtrOutput) ToStatusPtrOutput() StatusPtrOutput {
	return o
}

func (o StatusPtrOutput) ToStatusPtrOutputWithContext(ctx context.Context) StatusPtrOutput {
	return o
}

func (o StatusPtrOutput) Elem() StatusOutput {
	return o.ApplyT(func(v *Status) Status {
		if v != nil {
			return *v
		}
		var ret Status
		return ret
	}).(StatusOutput)
}

func (o StatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Status) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StatusInput interface {
	pulumi.Input

	ToStatusOutput() StatusOutput
	ToStatusOutputWithContext(context.Context) StatusOutput
}

var statusPtrType = reflect.TypeOf((**Status)(nil)).Elem()

type StatusPtrInput interface {
	pulumi.Input

	ToStatusPtrOutput() StatusPtrOutput
	ToStatusPtrOutputWithContext(context.Context) StatusPtrOutput
}

type statusPtr string

func StatusPtr(v string) StatusPtrInput {
	return (*statusPtr)(&v)
}

func (*statusPtr) ElementType() reflect.Type {
	return statusPtrType
}

func (in *statusPtr) ToStatusPtrOutput() StatusPtrOutput {
	return pulumi.ToOutput(in).(StatusPtrOutput)
}

func (in *statusPtr) ToStatusPtrOutputWithContext(ctx context.Context) StatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StatusPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StatusInput)(nil)).Elem(), Status("Approved"))
	pulumi.RegisterInputType(reflect.TypeOf((*StatusPtrInput)(nil)).Elem(), Status("Approved"))
	pulumi.RegisterOutputType(StatusOutput{})
	pulumi.RegisterOutputType(StatusPtrOutput{})
}

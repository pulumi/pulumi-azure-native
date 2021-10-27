


package v20200201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MediaGraphRtspTransport string

const (
	// HTTP/HTTPS transport. This should be used when HTTP tunneling is desired.
	MediaGraphRtspTransportHttp = MediaGraphRtspTransport("Http")
	// TCP transport. This should be used when HTTP tunneling is not desired.
	MediaGraphRtspTransportTcp = MediaGraphRtspTransport("Tcp")
)

func (MediaGraphRtspTransport) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphRtspTransport)(nil)).Elem()
}

func (e MediaGraphRtspTransport) ToMediaGraphRtspTransportOutput() MediaGraphRtspTransportOutput {
	return pulumi.ToOutput(e).(MediaGraphRtspTransportOutput)
}

func (e MediaGraphRtspTransport) ToMediaGraphRtspTransportOutputWithContext(ctx context.Context) MediaGraphRtspTransportOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MediaGraphRtspTransportOutput)
}

func (e MediaGraphRtspTransport) ToMediaGraphRtspTransportPtrOutput() MediaGraphRtspTransportPtrOutput {
	return e.ToMediaGraphRtspTransportPtrOutputWithContext(context.Background())
}

func (e MediaGraphRtspTransport) ToMediaGraphRtspTransportPtrOutputWithContext(ctx context.Context) MediaGraphRtspTransportPtrOutput {
	return MediaGraphRtspTransport(e).ToMediaGraphRtspTransportOutputWithContext(ctx).ToMediaGraphRtspTransportPtrOutputWithContext(ctx)
}

func (e MediaGraphRtspTransport) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MediaGraphRtspTransport) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MediaGraphRtspTransport) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MediaGraphRtspTransport) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MediaGraphRtspTransportOutput struct{ *pulumi.OutputState }

func (MediaGraphRtspTransportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphRtspTransport)(nil)).Elem()
}

func (o MediaGraphRtspTransportOutput) ToMediaGraphRtspTransportOutput() MediaGraphRtspTransportOutput {
	return o
}

func (o MediaGraphRtspTransportOutput) ToMediaGraphRtspTransportOutputWithContext(ctx context.Context) MediaGraphRtspTransportOutput {
	return o
}

func (o MediaGraphRtspTransportOutput) ToMediaGraphRtspTransportPtrOutput() MediaGraphRtspTransportPtrOutput {
	return o.ToMediaGraphRtspTransportPtrOutputWithContext(context.Background())
}

func (o MediaGraphRtspTransportOutput) ToMediaGraphRtspTransportPtrOutputWithContext(ctx context.Context) MediaGraphRtspTransportPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MediaGraphRtspTransport) *MediaGraphRtspTransport {
		return &v
	}).(MediaGraphRtspTransportPtrOutput)
}

func (o MediaGraphRtspTransportOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MediaGraphRtspTransportOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MediaGraphRtspTransport) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MediaGraphRtspTransportOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MediaGraphRtspTransportOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MediaGraphRtspTransport) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MediaGraphRtspTransportPtrOutput struct{ *pulumi.OutputState }

func (MediaGraphRtspTransportPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaGraphRtspTransport)(nil)).Elem()
}

func (o MediaGraphRtspTransportPtrOutput) ToMediaGraphRtspTransportPtrOutput() MediaGraphRtspTransportPtrOutput {
	return o
}

func (o MediaGraphRtspTransportPtrOutput) ToMediaGraphRtspTransportPtrOutputWithContext(ctx context.Context) MediaGraphRtspTransportPtrOutput {
	return o
}

func (o MediaGraphRtspTransportPtrOutput) Elem() MediaGraphRtspTransportOutput {
	return o.ApplyT(func(v *MediaGraphRtspTransport) MediaGraphRtspTransport {
		if v != nil {
			return *v
		}
		var ret MediaGraphRtspTransport
		return ret
	}).(MediaGraphRtspTransportOutput)
}

func (o MediaGraphRtspTransportPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MediaGraphRtspTransportPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MediaGraphRtspTransport) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MediaGraphRtspTransportInput interface {
	pulumi.Input

	ToMediaGraphRtspTransportOutput() MediaGraphRtspTransportOutput
	ToMediaGraphRtspTransportOutputWithContext(context.Context) MediaGraphRtspTransportOutput
}

var mediaGraphRtspTransportPtrType = reflect.TypeOf((**MediaGraphRtspTransport)(nil)).Elem()

type MediaGraphRtspTransportPtrInput interface {
	pulumi.Input

	ToMediaGraphRtspTransportPtrOutput() MediaGraphRtspTransportPtrOutput
	ToMediaGraphRtspTransportPtrOutputWithContext(context.Context) MediaGraphRtspTransportPtrOutput
}

type mediaGraphRtspTransportPtr string

func MediaGraphRtspTransportPtr(v string) MediaGraphRtspTransportPtrInput {
	return (*mediaGraphRtspTransportPtr)(&v)
}

func (*mediaGraphRtspTransportPtr) ElementType() reflect.Type {
	return mediaGraphRtspTransportPtrType
}

func (in *mediaGraphRtspTransportPtr) ToMediaGraphRtspTransportPtrOutput() MediaGraphRtspTransportPtrOutput {
	return pulumi.ToOutput(in).(MediaGraphRtspTransportPtrOutput)
}

func (in *mediaGraphRtspTransportPtr) ToMediaGraphRtspTransportPtrOutputWithContext(ctx context.Context) MediaGraphRtspTransportPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MediaGraphRtspTransportPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MediaGraphRtspTransportInput)(nil)).Elem(), MediaGraphRtspTransport("Http"))
	pulumi.RegisterInputType(reflect.TypeOf((*MediaGraphRtspTransportPtrInput)(nil)).Elem(), MediaGraphRtspTransport("Http"))
	pulumi.RegisterOutputType(MediaGraphRtspTransportOutput{})
	pulumi.RegisterOutputType(MediaGraphRtspTransportPtrOutput{})
}

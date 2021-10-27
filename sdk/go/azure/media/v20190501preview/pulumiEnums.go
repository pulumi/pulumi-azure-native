


package v20190501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LiveEventEncodingType string

const (
	LiveEventEncodingTypeNone         = LiveEventEncodingType("None")
	LiveEventEncodingTypeBasic        = LiveEventEncodingType("Basic")
	LiveEventEncodingTypeStandard     = LiveEventEncodingType("Standard")
	LiveEventEncodingTypePremium1080p = LiveEventEncodingType("Premium1080p")
)

func (LiveEventEncodingType) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEncodingType)(nil)).Elem()
}

func (e LiveEventEncodingType) ToLiveEventEncodingTypeOutput() LiveEventEncodingTypeOutput {
	return pulumi.ToOutput(e).(LiveEventEncodingTypeOutput)
}

func (e LiveEventEncodingType) ToLiveEventEncodingTypeOutputWithContext(ctx context.Context) LiveEventEncodingTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LiveEventEncodingTypeOutput)
}

func (e LiveEventEncodingType) ToLiveEventEncodingTypePtrOutput() LiveEventEncodingTypePtrOutput {
	return e.ToLiveEventEncodingTypePtrOutputWithContext(context.Background())
}

func (e LiveEventEncodingType) ToLiveEventEncodingTypePtrOutputWithContext(ctx context.Context) LiveEventEncodingTypePtrOutput {
	return LiveEventEncodingType(e).ToLiveEventEncodingTypeOutputWithContext(ctx).ToLiveEventEncodingTypePtrOutputWithContext(ctx)
}

func (e LiveEventEncodingType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LiveEventEncodingType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LiveEventEncodingType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LiveEventEncodingType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LiveEventEncodingTypeOutput struct{ *pulumi.OutputState }

func (LiveEventEncodingTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEncodingType)(nil)).Elem()
}

func (o LiveEventEncodingTypeOutput) ToLiveEventEncodingTypeOutput() LiveEventEncodingTypeOutput {
	return o
}

func (o LiveEventEncodingTypeOutput) ToLiveEventEncodingTypeOutputWithContext(ctx context.Context) LiveEventEncodingTypeOutput {
	return o
}

func (o LiveEventEncodingTypeOutput) ToLiveEventEncodingTypePtrOutput() LiveEventEncodingTypePtrOutput {
	return o.ToLiveEventEncodingTypePtrOutputWithContext(context.Background())
}

func (o LiveEventEncodingTypeOutput) ToLiveEventEncodingTypePtrOutputWithContext(ctx context.Context) LiveEventEncodingTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventEncodingType) *LiveEventEncodingType {
		return &v
	}).(LiveEventEncodingTypePtrOutput)
}

func (o LiveEventEncodingTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LiveEventEncodingTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LiveEventEncodingType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LiveEventEncodingTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LiveEventEncodingTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LiveEventEncodingType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LiveEventEncodingTypePtrOutput struct{ *pulumi.OutputState }

func (LiveEventEncodingTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventEncodingType)(nil)).Elem()
}

func (o LiveEventEncodingTypePtrOutput) ToLiveEventEncodingTypePtrOutput() LiveEventEncodingTypePtrOutput {
	return o
}

func (o LiveEventEncodingTypePtrOutput) ToLiveEventEncodingTypePtrOutputWithContext(ctx context.Context) LiveEventEncodingTypePtrOutput {
	return o
}

func (o LiveEventEncodingTypePtrOutput) Elem() LiveEventEncodingTypeOutput {
	return o.ApplyT(func(v *LiveEventEncodingType) LiveEventEncodingType {
		if v != nil {
			return *v
		}
		var ret LiveEventEncodingType
		return ret
	}).(LiveEventEncodingTypeOutput)
}

func (o LiveEventEncodingTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LiveEventEncodingTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LiveEventEncodingType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LiveEventEncodingTypeInput interface {
	pulumi.Input

	ToLiveEventEncodingTypeOutput() LiveEventEncodingTypeOutput
	ToLiveEventEncodingTypeOutputWithContext(context.Context) LiveEventEncodingTypeOutput
}

var liveEventEncodingTypePtrType = reflect.TypeOf((**LiveEventEncodingType)(nil)).Elem()

type LiveEventEncodingTypePtrInput interface {
	pulumi.Input

	ToLiveEventEncodingTypePtrOutput() LiveEventEncodingTypePtrOutput
	ToLiveEventEncodingTypePtrOutputWithContext(context.Context) LiveEventEncodingTypePtrOutput
}

type liveEventEncodingTypePtr string

func LiveEventEncodingTypePtr(v string) LiveEventEncodingTypePtrInput {
	return (*liveEventEncodingTypePtr)(&v)
}

func (*liveEventEncodingTypePtr) ElementType() reflect.Type {
	return liveEventEncodingTypePtrType
}

func (in *liveEventEncodingTypePtr) ToLiveEventEncodingTypePtrOutput() LiveEventEncodingTypePtrOutput {
	return pulumi.ToOutput(in).(LiveEventEncodingTypePtrOutput)
}

func (in *liveEventEncodingTypePtr) ToLiveEventEncodingTypePtrOutputWithContext(ctx context.Context) LiveEventEncodingTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LiveEventEncodingTypePtrOutput)
}

type LiveEventInputProtocol string

const (
	LiveEventInputProtocolFragmentedMP4 = LiveEventInputProtocol("FragmentedMP4")
	LiveEventInputProtocolRTMP          = LiveEventInputProtocol("RTMP")
)

func (LiveEventInputProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputProtocol)(nil)).Elem()
}

func (e LiveEventInputProtocol) ToLiveEventInputProtocolOutput() LiveEventInputProtocolOutput {
	return pulumi.ToOutput(e).(LiveEventInputProtocolOutput)
}

func (e LiveEventInputProtocol) ToLiveEventInputProtocolOutputWithContext(ctx context.Context) LiveEventInputProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LiveEventInputProtocolOutput)
}

func (e LiveEventInputProtocol) ToLiveEventInputProtocolPtrOutput() LiveEventInputProtocolPtrOutput {
	return e.ToLiveEventInputProtocolPtrOutputWithContext(context.Background())
}

func (e LiveEventInputProtocol) ToLiveEventInputProtocolPtrOutputWithContext(ctx context.Context) LiveEventInputProtocolPtrOutput {
	return LiveEventInputProtocol(e).ToLiveEventInputProtocolOutputWithContext(ctx).ToLiveEventInputProtocolPtrOutputWithContext(ctx)
}

func (e LiveEventInputProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LiveEventInputProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LiveEventInputProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LiveEventInputProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LiveEventInputProtocolOutput struct{ *pulumi.OutputState }

func (LiveEventInputProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputProtocol)(nil)).Elem()
}

func (o LiveEventInputProtocolOutput) ToLiveEventInputProtocolOutput() LiveEventInputProtocolOutput {
	return o
}

func (o LiveEventInputProtocolOutput) ToLiveEventInputProtocolOutputWithContext(ctx context.Context) LiveEventInputProtocolOutput {
	return o
}

func (o LiveEventInputProtocolOutput) ToLiveEventInputProtocolPtrOutput() LiveEventInputProtocolPtrOutput {
	return o.ToLiveEventInputProtocolPtrOutputWithContext(context.Background())
}

func (o LiveEventInputProtocolOutput) ToLiveEventInputProtocolPtrOutputWithContext(ctx context.Context) LiveEventInputProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventInputProtocol) *LiveEventInputProtocol {
		return &v
	}).(LiveEventInputProtocolPtrOutput)
}

func (o LiveEventInputProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LiveEventInputProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LiveEventInputProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LiveEventInputProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LiveEventInputProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LiveEventInputProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LiveEventInputProtocolPtrOutput struct{ *pulumi.OutputState }

func (LiveEventInputProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventInputProtocol)(nil)).Elem()
}

func (o LiveEventInputProtocolPtrOutput) ToLiveEventInputProtocolPtrOutput() LiveEventInputProtocolPtrOutput {
	return o
}

func (o LiveEventInputProtocolPtrOutput) ToLiveEventInputProtocolPtrOutputWithContext(ctx context.Context) LiveEventInputProtocolPtrOutput {
	return o
}

func (o LiveEventInputProtocolPtrOutput) Elem() LiveEventInputProtocolOutput {
	return o.ApplyT(func(v *LiveEventInputProtocol) LiveEventInputProtocol {
		if v != nil {
			return *v
		}
		var ret LiveEventInputProtocol
		return ret
	}).(LiveEventInputProtocolOutput)
}

func (o LiveEventInputProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LiveEventInputProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LiveEventInputProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LiveEventInputProtocolInput interface {
	pulumi.Input

	ToLiveEventInputProtocolOutput() LiveEventInputProtocolOutput
	ToLiveEventInputProtocolOutputWithContext(context.Context) LiveEventInputProtocolOutput
}

var liveEventInputProtocolPtrType = reflect.TypeOf((**LiveEventInputProtocol)(nil)).Elem()

type LiveEventInputProtocolPtrInput interface {
	pulumi.Input

	ToLiveEventInputProtocolPtrOutput() LiveEventInputProtocolPtrOutput
	ToLiveEventInputProtocolPtrOutputWithContext(context.Context) LiveEventInputProtocolPtrOutput
}

type liveEventInputProtocolPtr string

func LiveEventInputProtocolPtr(v string) LiveEventInputProtocolPtrInput {
	return (*liveEventInputProtocolPtr)(&v)
}

func (*liveEventInputProtocolPtr) ElementType() reflect.Type {
	return liveEventInputProtocolPtrType
}

func (in *liveEventInputProtocolPtr) ToLiveEventInputProtocolPtrOutput() LiveEventInputProtocolPtrOutput {
	return pulumi.ToOutput(in).(LiveEventInputProtocolPtrOutput)
}

func (in *liveEventInputProtocolPtr) ToLiveEventInputProtocolPtrOutputWithContext(ctx context.Context) LiveEventInputProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LiveEventInputProtocolPtrOutput)
}

type StreamOptionsFlag string

const (
	StreamOptionsFlagDefault    = StreamOptionsFlag("Default")
	StreamOptionsFlagLowLatency = StreamOptionsFlag("LowLatency")
)

func (StreamOptionsFlag) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamOptionsFlag)(nil)).Elem()
}

func (e StreamOptionsFlag) ToStreamOptionsFlagOutput() StreamOptionsFlagOutput {
	return pulumi.ToOutput(e).(StreamOptionsFlagOutput)
}

func (e StreamOptionsFlag) ToStreamOptionsFlagOutputWithContext(ctx context.Context) StreamOptionsFlagOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StreamOptionsFlagOutput)
}

func (e StreamOptionsFlag) ToStreamOptionsFlagPtrOutput() StreamOptionsFlagPtrOutput {
	return e.ToStreamOptionsFlagPtrOutputWithContext(context.Background())
}

func (e StreamOptionsFlag) ToStreamOptionsFlagPtrOutputWithContext(ctx context.Context) StreamOptionsFlagPtrOutput {
	return StreamOptionsFlag(e).ToStreamOptionsFlagOutputWithContext(ctx).ToStreamOptionsFlagPtrOutputWithContext(ctx)
}

func (e StreamOptionsFlag) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StreamOptionsFlag) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StreamOptionsFlag) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StreamOptionsFlag) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StreamOptionsFlagOutput struct{ *pulumi.OutputState }

func (StreamOptionsFlagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamOptionsFlag)(nil)).Elem()
}

func (o StreamOptionsFlagOutput) ToStreamOptionsFlagOutput() StreamOptionsFlagOutput {
	return o
}

func (o StreamOptionsFlagOutput) ToStreamOptionsFlagOutputWithContext(ctx context.Context) StreamOptionsFlagOutput {
	return o
}

func (o StreamOptionsFlagOutput) ToStreamOptionsFlagPtrOutput() StreamOptionsFlagPtrOutput {
	return o.ToStreamOptionsFlagPtrOutputWithContext(context.Background())
}

func (o StreamOptionsFlagOutput) ToStreamOptionsFlagPtrOutputWithContext(ctx context.Context) StreamOptionsFlagPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamOptionsFlag) *StreamOptionsFlag {
		return &v
	}).(StreamOptionsFlagPtrOutput)
}

func (o StreamOptionsFlagOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StreamOptionsFlagOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StreamOptionsFlag) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StreamOptionsFlagOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StreamOptionsFlagOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StreamOptionsFlag) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StreamOptionsFlagPtrOutput struct{ *pulumi.OutputState }

func (StreamOptionsFlagPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamOptionsFlag)(nil)).Elem()
}

func (o StreamOptionsFlagPtrOutput) ToStreamOptionsFlagPtrOutput() StreamOptionsFlagPtrOutput {
	return o
}

func (o StreamOptionsFlagPtrOutput) ToStreamOptionsFlagPtrOutputWithContext(ctx context.Context) StreamOptionsFlagPtrOutput {
	return o
}

func (o StreamOptionsFlagPtrOutput) Elem() StreamOptionsFlagOutput {
	return o.ApplyT(func(v *StreamOptionsFlag) StreamOptionsFlag {
		if v != nil {
			return *v
		}
		var ret StreamOptionsFlag
		return ret
	}).(StreamOptionsFlagOutput)
}

func (o StreamOptionsFlagPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StreamOptionsFlagPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StreamOptionsFlag) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StreamOptionsFlagInput interface {
	pulumi.Input

	ToStreamOptionsFlagOutput() StreamOptionsFlagOutput
	ToStreamOptionsFlagOutputWithContext(context.Context) StreamOptionsFlagOutput
}

var streamOptionsFlagPtrType = reflect.TypeOf((**StreamOptionsFlag)(nil)).Elem()

type StreamOptionsFlagPtrInput interface {
	pulumi.Input

	ToStreamOptionsFlagPtrOutput() StreamOptionsFlagPtrOutput
	ToStreamOptionsFlagPtrOutputWithContext(context.Context) StreamOptionsFlagPtrOutput
}

type streamOptionsFlagPtr string

func StreamOptionsFlagPtr(v string) StreamOptionsFlagPtrInput {
	return (*streamOptionsFlagPtr)(&v)
}

func (*streamOptionsFlagPtr) ElementType() reflect.Type {
	return streamOptionsFlagPtrType
}

func (in *streamOptionsFlagPtr) ToStreamOptionsFlagPtrOutput() StreamOptionsFlagPtrOutput {
	return pulumi.ToOutput(in).(StreamOptionsFlagPtrOutput)
}

func (in *streamOptionsFlagPtr) ToStreamOptionsFlagPtrOutputWithContext(ctx context.Context) StreamOptionsFlagPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StreamOptionsFlagPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventEncodingTypeInput)(nil)).Elem(), LiveEventEncodingType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventEncodingTypePtrInput)(nil)).Elem(), LiveEventEncodingType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputProtocolInput)(nil)).Elem(), LiveEventInputProtocol("FragmentedMP4"))
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputProtocolPtrInput)(nil)).Elem(), LiveEventInputProtocol("FragmentedMP4"))
	pulumi.RegisterInputType(reflect.TypeOf((*StreamOptionsFlagInput)(nil)).Elem(), StreamOptionsFlag("Default"))
	pulumi.RegisterInputType(reflect.TypeOf((*StreamOptionsFlagPtrInput)(nil)).Elem(), StreamOptionsFlag("Default"))
	pulumi.RegisterOutputType(LiveEventEncodingTypeOutput{})
	pulumi.RegisterOutputType(LiveEventEncodingTypePtrOutput{})
	pulumi.RegisterOutputType(LiveEventInputProtocolOutput{})
	pulumi.RegisterOutputType(LiveEventInputProtocolPtrOutput{})
	pulumi.RegisterOutputType(StreamOptionsFlagOutput{})
	pulumi.RegisterOutputType(StreamOptionsFlagPtrOutput{})
}

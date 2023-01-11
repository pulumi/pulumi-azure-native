


package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CommunicationsPlatform string

const (
	CommunicationsPlatformOperatorConnect  = CommunicationsPlatform("OperatorConnect")
	CommunicationsPlatformTeamsPhoneMobile = CommunicationsPlatform("TeamsPhoneMobile")
)

type Connectivity string

const (
	// This deployment connects to the operator network using a Public IP address, e.g. when using MAPS
	ConnectivityPublicAddress = Connectivity("PublicAddress")
)

type E911Type string

const (
	// Emergency calls are not handled different from other calls
	E911TypeStandard = E911Type("Standard")
	// Emergency calls are routed directly to the ESRP
	E911TypeDirectToEsrp = E911Type("DirectToEsrp")
)

type TeamsCodecs string

const (
	TeamsCodecsPCMA     = TeamsCodecs("PCMA")
	TeamsCodecsPCMU     = TeamsCodecs("PCMU")
	TeamsCodecsG722     = TeamsCodecs("G722")
	TeamsCodecs_G722_2  = TeamsCodecs("G722_2")
	TeamsCodecs_SILK_8  = TeamsCodecs("SILK_8")
	TeamsCodecs_SILK_16 = TeamsCodecs("SILK_16")
)

type TestLinePurpose string

const (
	TestLinePurposeManual    = TestLinePurpose("Manual")
	TestLinePurposeAutomated = TestLinePurpose("Automated")
)

func (TestLinePurpose) ElementType() reflect.Type {
	return reflect.TypeOf((*TestLinePurpose)(nil)).Elem()
}

func (e TestLinePurpose) ToTestLinePurposeOutput() TestLinePurposeOutput {
	return pulumi.ToOutput(e).(TestLinePurposeOutput)
}

func (e TestLinePurpose) ToTestLinePurposeOutputWithContext(ctx context.Context) TestLinePurposeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TestLinePurposeOutput)
}

func (e TestLinePurpose) ToTestLinePurposePtrOutput() TestLinePurposePtrOutput {
	return e.ToTestLinePurposePtrOutputWithContext(context.Background())
}

func (e TestLinePurpose) ToTestLinePurposePtrOutputWithContext(ctx context.Context) TestLinePurposePtrOutput {
	return TestLinePurpose(e).ToTestLinePurposeOutputWithContext(ctx).ToTestLinePurposePtrOutputWithContext(ctx)
}

func (e TestLinePurpose) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TestLinePurpose) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TestLinePurpose) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TestLinePurpose) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TestLinePurposeOutput struct{ *pulumi.OutputState }

func (TestLinePurposeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TestLinePurpose)(nil)).Elem()
}

func (o TestLinePurposeOutput) ToTestLinePurposeOutput() TestLinePurposeOutput {
	return o
}

func (o TestLinePurposeOutput) ToTestLinePurposeOutputWithContext(ctx context.Context) TestLinePurposeOutput {
	return o
}

func (o TestLinePurposeOutput) ToTestLinePurposePtrOutput() TestLinePurposePtrOutput {
	return o.ToTestLinePurposePtrOutputWithContext(context.Background())
}

func (o TestLinePurposeOutput) ToTestLinePurposePtrOutputWithContext(ctx context.Context) TestLinePurposePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TestLinePurpose) *TestLinePurpose {
		return &v
	}).(TestLinePurposePtrOutput)
}

func (o TestLinePurposeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TestLinePurposeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TestLinePurpose) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TestLinePurposeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TestLinePurposeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TestLinePurpose) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TestLinePurposePtrOutput struct{ *pulumi.OutputState }

func (TestLinePurposePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TestLinePurpose)(nil)).Elem()
}

func (o TestLinePurposePtrOutput) ToTestLinePurposePtrOutput() TestLinePurposePtrOutput {
	return o
}

func (o TestLinePurposePtrOutput) ToTestLinePurposePtrOutputWithContext(ctx context.Context) TestLinePurposePtrOutput {
	return o
}

func (o TestLinePurposePtrOutput) Elem() TestLinePurposeOutput {
	return o.ApplyT(func(v *TestLinePurpose) TestLinePurpose {
		if v != nil {
			return *v
		}
		var ret TestLinePurpose
		return ret
	}).(TestLinePurposeOutput)
}

func (o TestLinePurposePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TestLinePurposePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TestLinePurpose) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TestLinePurposeInput interface {
	pulumi.Input

	ToTestLinePurposeOutput() TestLinePurposeOutput
	ToTestLinePurposeOutputWithContext(context.Context) TestLinePurposeOutput
}

var testLinePurposePtrType = reflect.TypeOf((**TestLinePurpose)(nil)).Elem()

type TestLinePurposePtrInput interface {
	pulumi.Input

	ToTestLinePurposePtrOutput() TestLinePurposePtrOutput
	ToTestLinePurposePtrOutputWithContext(context.Context) TestLinePurposePtrOutput
}

type testLinePurposePtr string

func TestLinePurposePtr(v string) TestLinePurposePtrInput {
	return (*testLinePurposePtr)(&v)
}

func (*testLinePurposePtr) ElementType() reflect.Type {
	return testLinePurposePtrType
}

func (in *testLinePurposePtr) ToTestLinePurposePtrOutput() TestLinePurposePtrOutput {
	return pulumi.ToOutput(in).(TestLinePurposePtrOutput)
}

func (in *testLinePurposePtr) ToTestLinePurposePtrOutputWithContext(ctx context.Context) TestLinePurposePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TestLinePurposePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(TestLinePurposeOutput{})
	pulumi.RegisterOutputType(TestLinePurposePtrOutput{})
}

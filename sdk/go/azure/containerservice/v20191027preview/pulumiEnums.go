


package v20191027preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OSType string

const (
	OSTypeLinux   = OSType("Linux")
	OSTypeWindows = OSType("Windows")
)

func (OSType) ElementType() reflect.Type {
	return reflect.TypeOf((*OSType)(nil)).Elem()
}

func (e OSType) ToOSTypeOutput() OSTypeOutput {
	return pulumi.ToOutput(e).(OSTypeOutput)
}

func (e OSType) ToOSTypeOutputWithContext(ctx context.Context) OSTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OSTypeOutput)
}

func (e OSType) ToOSTypePtrOutput() OSTypePtrOutput {
	return e.ToOSTypePtrOutputWithContext(context.Background())
}

func (e OSType) ToOSTypePtrOutputWithContext(ctx context.Context) OSTypePtrOutput {
	return OSType(e).ToOSTypeOutputWithContext(ctx).ToOSTypePtrOutputWithContext(ctx)
}

func (e OSType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OSType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OSType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OSType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OSTypeOutput struct{ *pulumi.OutputState }

func (OSTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSType)(nil)).Elem()
}

func (o OSTypeOutput) ToOSTypeOutput() OSTypeOutput {
	return o
}

func (o OSTypeOutput) ToOSTypeOutputWithContext(ctx context.Context) OSTypeOutput {
	return o
}

func (o OSTypeOutput) ToOSTypePtrOutput() OSTypePtrOutput {
	return o.ToOSTypePtrOutputWithContext(context.Background())
}

func (o OSTypeOutput) ToOSTypePtrOutputWithContext(ctx context.Context) OSTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSType) *OSType {
		return &v
	}).(OSTypePtrOutput)
}

func (o OSTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OSTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OSType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OSTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OSTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OSType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OSTypePtrOutput struct{ *pulumi.OutputState }

func (OSTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSType)(nil)).Elem()
}

func (o OSTypePtrOutput) ToOSTypePtrOutput() OSTypePtrOutput {
	return o
}

func (o OSTypePtrOutput) ToOSTypePtrOutputWithContext(ctx context.Context) OSTypePtrOutput {
	return o
}

func (o OSTypePtrOutput) Elem() OSTypeOutput {
	return o.ApplyT(func(v *OSType) OSType {
		if v != nil {
			return *v
		}
		var ret OSType
		return ret
	}).(OSTypeOutput)
}

func (o OSTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OSTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OSType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OSTypeInput interface {
	pulumi.Input

	ToOSTypeOutput() OSTypeOutput
	ToOSTypeOutputWithContext(context.Context) OSTypeOutput
}

var ostypePtrType = reflect.TypeOf((**OSType)(nil)).Elem()

type OSTypePtrInput interface {
	pulumi.Input

	ToOSTypePtrOutput() OSTypePtrOutput
	ToOSTypePtrOutputWithContext(context.Context) OSTypePtrOutput
}

type ostypePtr string

func OSTypePtr(v string) OSTypePtrInput {
	return (*ostypePtr)(&v)
}

func (*ostypePtr) ElementType() reflect.Type {
	return ostypePtrType
}

func (in *ostypePtr) ToOSTypePtrOutput() OSTypePtrOutput {
	return pulumi.ToOutput(in).(OSTypePtrOutput)
}

func (in *ostypePtr) ToOSTypePtrOutputWithContext(ctx context.Context) OSTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OSTypePtrOutput)
}

type OpenShiftAgentPoolProfileRole string

const (
	OpenShiftAgentPoolProfileRoleCompute = OpenShiftAgentPoolProfileRole("compute")
	OpenShiftAgentPoolProfileRoleInfra   = OpenShiftAgentPoolProfileRole("infra")
)

func (OpenShiftAgentPoolProfileRole) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftAgentPoolProfileRole)(nil)).Elem()
}

func (e OpenShiftAgentPoolProfileRole) ToOpenShiftAgentPoolProfileRoleOutput() OpenShiftAgentPoolProfileRoleOutput {
	return pulumi.ToOutput(e).(OpenShiftAgentPoolProfileRoleOutput)
}

func (e OpenShiftAgentPoolProfileRole) ToOpenShiftAgentPoolProfileRoleOutputWithContext(ctx context.Context) OpenShiftAgentPoolProfileRoleOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OpenShiftAgentPoolProfileRoleOutput)
}

func (e OpenShiftAgentPoolProfileRole) ToOpenShiftAgentPoolProfileRolePtrOutput() OpenShiftAgentPoolProfileRolePtrOutput {
	return e.ToOpenShiftAgentPoolProfileRolePtrOutputWithContext(context.Background())
}

func (e OpenShiftAgentPoolProfileRole) ToOpenShiftAgentPoolProfileRolePtrOutputWithContext(ctx context.Context) OpenShiftAgentPoolProfileRolePtrOutput {
	return OpenShiftAgentPoolProfileRole(e).ToOpenShiftAgentPoolProfileRoleOutputWithContext(ctx).ToOpenShiftAgentPoolProfileRolePtrOutputWithContext(ctx)
}

func (e OpenShiftAgentPoolProfileRole) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OpenShiftAgentPoolProfileRole) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OpenShiftAgentPoolProfileRole) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OpenShiftAgentPoolProfileRole) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OpenShiftAgentPoolProfileRoleOutput struct{ *pulumi.OutputState }

func (OpenShiftAgentPoolProfileRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftAgentPoolProfileRole)(nil)).Elem()
}

func (o OpenShiftAgentPoolProfileRoleOutput) ToOpenShiftAgentPoolProfileRoleOutput() OpenShiftAgentPoolProfileRoleOutput {
	return o
}

func (o OpenShiftAgentPoolProfileRoleOutput) ToOpenShiftAgentPoolProfileRoleOutputWithContext(ctx context.Context) OpenShiftAgentPoolProfileRoleOutput {
	return o
}

func (o OpenShiftAgentPoolProfileRoleOutput) ToOpenShiftAgentPoolProfileRolePtrOutput() OpenShiftAgentPoolProfileRolePtrOutput {
	return o.ToOpenShiftAgentPoolProfileRolePtrOutputWithContext(context.Background())
}

func (o OpenShiftAgentPoolProfileRoleOutput) ToOpenShiftAgentPoolProfileRolePtrOutputWithContext(ctx context.Context) OpenShiftAgentPoolProfileRolePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OpenShiftAgentPoolProfileRole) *OpenShiftAgentPoolProfileRole {
		return &v
	}).(OpenShiftAgentPoolProfileRolePtrOutput)
}

func (o OpenShiftAgentPoolProfileRoleOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OpenShiftAgentPoolProfileRoleOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OpenShiftAgentPoolProfileRole) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OpenShiftAgentPoolProfileRoleOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OpenShiftAgentPoolProfileRoleOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OpenShiftAgentPoolProfileRole) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OpenShiftAgentPoolProfileRolePtrOutput struct{ *pulumi.OutputState }

func (OpenShiftAgentPoolProfileRolePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenShiftAgentPoolProfileRole)(nil)).Elem()
}

func (o OpenShiftAgentPoolProfileRolePtrOutput) ToOpenShiftAgentPoolProfileRolePtrOutput() OpenShiftAgentPoolProfileRolePtrOutput {
	return o
}

func (o OpenShiftAgentPoolProfileRolePtrOutput) ToOpenShiftAgentPoolProfileRolePtrOutputWithContext(ctx context.Context) OpenShiftAgentPoolProfileRolePtrOutput {
	return o
}

func (o OpenShiftAgentPoolProfileRolePtrOutput) Elem() OpenShiftAgentPoolProfileRoleOutput {
	return o.ApplyT(func(v *OpenShiftAgentPoolProfileRole) OpenShiftAgentPoolProfileRole {
		if v != nil {
			return *v
		}
		var ret OpenShiftAgentPoolProfileRole
		return ret
	}).(OpenShiftAgentPoolProfileRoleOutput)
}

func (o OpenShiftAgentPoolProfileRolePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OpenShiftAgentPoolProfileRolePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OpenShiftAgentPoolProfileRole) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OpenShiftAgentPoolProfileRoleInput interface {
	pulumi.Input

	ToOpenShiftAgentPoolProfileRoleOutput() OpenShiftAgentPoolProfileRoleOutput
	ToOpenShiftAgentPoolProfileRoleOutputWithContext(context.Context) OpenShiftAgentPoolProfileRoleOutput
}

var openShiftAgentPoolProfileRolePtrType = reflect.TypeOf((**OpenShiftAgentPoolProfileRole)(nil)).Elem()

type OpenShiftAgentPoolProfileRolePtrInput interface {
	pulumi.Input

	ToOpenShiftAgentPoolProfileRolePtrOutput() OpenShiftAgentPoolProfileRolePtrOutput
	ToOpenShiftAgentPoolProfileRolePtrOutputWithContext(context.Context) OpenShiftAgentPoolProfileRolePtrOutput
}

type openShiftAgentPoolProfileRolePtr string

func OpenShiftAgentPoolProfileRolePtr(v string) OpenShiftAgentPoolProfileRolePtrInput {
	return (*openShiftAgentPoolProfileRolePtr)(&v)
}

func (*openShiftAgentPoolProfileRolePtr) ElementType() reflect.Type {
	return openShiftAgentPoolProfileRolePtrType
}

func (in *openShiftAgentPoolProfileRolePtr) ToOpenShiftAgentPoolProfileRolePtrOutput() OpenShiftAgentPoolProfileRolePtrOutput {
	return pulumi.ToOutput(in).(OpenShiftAgentPoolProfileRolePtrOutput)
}

func (in *openShiftAgentPoolProfileRolePtr) ToOpenShiftAgentPoolProfileRolePtrOutputWithContext(ctx context.Context) OpenShiftAgentPoolProfileRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OpenShiftAgentPoolProfileRolePtrOutput)
}

type OpenShiftContainerServiceVMSize string

const (
	OpenShiftContainerServiceVMSize_Standard_D2s_v3  = OpenShiftContainerServiceVMSize("Standard_D2s_v3")
	OpenShiftContainerServiceVMSize_Standard_D4s_v3  = OpenShiftContainerServiceVMSize("Standard_D4s_v3")
	OpenShiftContainerServiceVMSize_Standard_D8s_v3  = OpenShiftContainerServiceVMSize("Standard_D8s_v3")
	OpenShiftContainerServiceVMSize_Standard_D16s_v3 = OpenShiftContainerServiceVMSize("Standard_D16s_v3")
	OpenShiftContainerServiceVMSize_Standard_D32s_v3 = OpenShiftContainerServiceVMSize("Standard_D32s_v3")
	OpenShiftContainerServiceVMSize_Standard_D64s_v3 = OpenShiftContainerServiceVMSize("Standard_D64s_v3")
	OpenShiftContainerServiceVMSize_Standard_DS4_v2  = OpenShiftContainerServiceVMSize("Standard_DS4_v2")
	OpenShiftContainerServiceVMSize_Standard_DS5_v2  = OpenShiftContainerServiceVMSize("Standard_DS5_v2")
	OpenShiftContainerServiceVMSize_Standard_F8s_v2  = OpenShiftContainerServiceVMSize("Standard_F8s_v2")
	OpenShiftContainerServiceVMSize_Standard_F16s_v2 = OpenShiftContainerServiceVMSize("Standard_F16s_v2")
	OpenShiftContainerServiceVMSize_Standard_F32s_v2 = OpenShiftContainerServiceVMSize("Standard_F32s_v2")
	OpenShiftContainerServiceVMSize_Standard_F64s_v2 = OpenShiftContainerServiceVMSize("Standard_F64s_v2")
	OpenShiftContainerServiceVMSize_Standard_F72s_v2 = OpenShiftContainerServiceVMSize("Standard_F72s_v2")
	OpenShiftContainerServiceVMSize_Standard_F8s     = OpenShiftContainerServiceVMSize("Standard_F8s")
	OpenShiftContainerServiceVMSize_Standard_F16s    = OpenShiftContainerServiceVMSize("Standard_F16s")
	OpenShiftContainerServiceVMSize_Standard_E4s_v3  = OpenShiftContainerServiceVMSize("Standard_E4s_v3")
	OpenShiftContainerServiceVMSize_Standard_E8s_v3  = OpenShiftContainerServiceVMSize("Standard_E8s_v3")
	OpenShiftContainerServiceVMSize_Standard_E16s_v3 = OpenShiftContainerServiceVMSize("Standard_E16s_v3")
	OpenShiftContainerServiceVMSize_Standard_E20s_v3 = OpenShiftContainerServiceVMSize("Standard_E20s_v3")
	OpenShiftContainerServiceVMSize_Standard_E32s_v3 = OpenShiftContainerServiceVMSize("Standard_E32s_v3")
	OpenShiftContainerServiceVMSize_Standard_E64s_v3 = OpenShiftContainerServiceVMSize("Standard_E64s_v3")
	OpenShiftContainerServiceVMSize_Standard_GS2     = OpenShiftContainerServiceVMSize("Standard_GS2")
	OpenShiftContainerServiceVMSize_Standard_GS3     = OpenShiftContainerServiceVMSize("Standard_GS3")
	OpenShiftContainerServiceVMSize_Standard_GS4     = OpenShiftContainerServiceVMSize("Standard_GS4")
	OpenShiftContainerServiceVMSize_Standard_GS5     = OpenShiftContainerServiceVMSize("Standard_GS5")
	OpenShiftContainerServiceVMSize_Standard_DS12_v2 = OpenShiftContainerServiceVMSize("Standard_DS12_v2")
	OpenShiftContainerServiceVMSize_Standard_DS13_v2 = OpenShiftContainerServiceVMSize("Standard_DS13_v2")
	OpenShiftContainerServiceVMSize_Standard_DS14_v2 = OpenShiftContainerServiceVMSize("Standard_DS14_v2")
	OpenShiftContainerServiceVMSize_Standard_DS15_v2 = OpenShiftContainerServiceVMSize("Standard_DS15_v2")
	OpenShiftContainerServiceVMSize_Standard_L4s     = OpenShiftContainerServiceVMSize("Standard_L4s")
	OpenShiftContainerServiceVMSize_Standard_L8s     = OpenShiftContainerServiceVMSize("Standard_L8s")
	OpenShiftContainerServiceVMSize_Standard_L16s    = OpenShiftContainerServiceVMSize("Standard_L16s")
	OpenShiftContainerServiceVMSize_Standard_L32s    = OpenShiftContainerServiceVMSize("Standard_L32s")
)

func (OpenShiftContainerServiceVMSize) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftContainerServiceVMSize)(nil)).Elem()
}

func (e OpenShiftContainerServiceVMSize) ToOpenShiftContainerServiceVMSizeOutput() OpenShiftContainerServiceVMSizeOutput {
	return pulumi.ToOutput(e).(OpenShiftContainerServiceVMSizeOutput)
}

func (e OpenShiftContainerServiceVMSize) ToOpenShiftContainerServiceVMSizeOutputWithContext(ctx context.Context) OpenShiftContainerServiceVMSizeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OpenShiftContainerServiceVMSizeOutput)
}

func (e OpenShiftContainerServiceVMSize) ToOpenShiftContainerServiceVMSizePtrOutput() OpenShiftContainerServiceVMSizePtrOutput {
	return e.ToOpenShiftContainerServiceVMSizePtrOutputWithContext(context.Background())
}

func (e OpenShiftContainerServiceVMSize) ToOpenShiftContainerServiceVMSizePtrOutputWithContext(ctx context.Context) OpenShiftContainerServiceVMSizePtrOutput {
	return OpenShiftContainerServiceVMSize(e).ToOpenShiftContainerServiceVMSizeOutputWithContext(ctx).ToOpenShiftContainerServiceVMSizePtrOutputWithContext(ctx)
}

func (e OpenShiftContainerServiceVMSize) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OpenShiftContainerServiceVMSize) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OpenShiftContainerServiceVMSize) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OpenShiftContainerServiceVMSize) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OpenShiftContainerServiceVMSizeOutput struct{ *pulumi.OutputState }

func (OpenShiftContainerServiceVMSizeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftContainerServiceVMSize)(nil)).Elem()
}

func (o OpenShiftContainerServiceVMSizeOutput) ToOpenShiftContainerServiceVMSizeOutput() OpenShiftContainerServiceVMSizeOutput {
	return o
}

func (o OpenShiftContainerServiceVMSizeOutput) ToOpenShiftContainerServiceVMSizeOutputWithContext(ctx context.Context) OpenShiftContainerServiceVMSizeOutput {
	return o
}

func (o OpenShiftContainerServiceVMSizeOutput) ToOpenShiftContainerServiceVMSizePtrOutput() OpenShiftContainerServiceVMSizePtrOutput {
	return o.ToOpenShiftContainerServiceVMSizePtrOutputWithContext(context.Background())
}

func (o OpenShiftContainerServiceVMSizeOutput) ToOpenShiftContainerServiceVMSizePtrOutputWithContext(ctx context.Context) OpenShiftContainerServiceVMSizePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OpenShiftContainerServiceVMSize) *OpenShiftContainerServiceVMSize {
		return &v
	}).(OpenShiftContainerServiceVMSizePtrOutput)
}

func (o OpenShiftContainerServiceVMSizeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OpenShiftContainerServiceVMSizeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OpenShiftContainerServiceVMSize) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OpenShiftContainerServiceVMSizeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OpenShiftContainerServiceVMSizeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OpenShiftContainerServiceVMSize) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OpenShiftContainerServiceVMSizePtrOutput struct{ *pulumi.OutputState }

func (OpenShiftContainerServiceVMSizePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenShiftContainerServiceVMSize)(nil)).Elem()
}

func (o OpenShiftContainerServiceVMSizePtrOutput) ToOpenShiftContainerServiceVMSizePtrOutput() OpenShiftContainerServiceVMSizePtrOutput {
	return o
}

func (o OpenShiftContainerServiceVMSizePtrOutput) ToOpenShiftContainerServiceVMSizePtrOutputWithContext(ctx context.Context) OpenShiftContainerServiceVMSizePtrOutput {
	return o
}

func (o OpenShiftContainerServiceVMSizePtrOutput) Elem() OpenShiftContainerServiceVMSizeOutput {
	return o.ApplyT(func(v *OpenShiftContainerServiceVMSize) OpenShiftContainerServiceVMSize {
		if v != nil {
			return *v
		}
		var ret OpenShiftContainerServiceVMSize
		return ret
	}).(OpenShiftContainerServiceVMSizeOutput)
}

func (o OpenShiftContainerServiceVMSizePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OpenShiftContainerServiceVMSizePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OpenShiftContainerServiceVMSize) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OpenShiftContainerServiceVMSizeInput interface {
	pulumi.Input

	ToOpenShiftContainerServiceVMSizeOutput() OpenShiftContainerServiceVMSizeOutput
	ToOpenShiftContainerServiceVMSizeOutputWithContext(context.Context) OpenShiftContainerServiceVMSizeOutput
}

var openShiftContainerServiceVMSizePtrType = reflect.TypeOf((**OpenShiftContainerServiceVMSize)(nil)).Elem()

type OpenShiftContainerServiceVMSizePtrInput interface {
	pulumi.Input

	ToOpenShiftContainerServiceVMSizePtrOutput() OpenShiftContainerServiceVMSizePtrOutput
	ToOpenShiftContainerServiceVMSizePtrOutputWithContext(context.Context) OpenShiftContainerServiceVMSizePtrOutput
}

type openShiftContainerServiceVMSizePtr string

func OpenShiftContainerServiceVMSizePtr(v string) OpenShiftContainerServiceVMSizePtrInput {
	return (*openShiftContainerServiceVMSizePtr)(&v)
}

func (*openShiftContainerServiceVMSizePtr) ElementType() reflect.Type {
	return openShiftContainerServiceVMSizePtrType
}

func (in *openShiftContainerServiceVMSizePtr) ToOpenShiftContainerServiceVMSizePtrOutput() OpenShiftContainerServiceVMSizePtrOutput {
	return pulumi.ToOutput(in).(OpenShiftContainerServiceVMSizePtrOutput)
}

func (in *openShiftContainerServiceVMSizePtr) ToOpenShiftContainerServiceVMSizePtrOutputWithContext(ctx context.Context) OpenShiftContainerServiceVMSizePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OpenShiftContainerServiceVMSizePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OSTypeInput)(nil)).Elem(), OSType("Linux"))
	pulumi.RegisterInputType(reflect.TypeOf((*OSTypePtrInput)(nil)).Elem(), OSType("Linux"))
	pulumi.RegisterInputType(reflect.TypeOf((*OpenShiftAgentPoolProfileRoleInput)(nil)).Elem(), OpenShiftAgentPoolProfileRole("compute"))
	pulumi.RegisterInputType(reflect.TypeOf((*OpenShiftAgentPoolProfileRolePtrInput)(nil)).Elem(), OpenShiftAgentPoolProfileRole("compute"))
	pulumi.RegisterInputType(reflect.TypeOf((*OpenShiftContainerServiceVMSizeInput)(nil)).Elem(), OpenShiftContainerServiceVMSize("Standard_D2s_v3"))
	pulumi.RegisterInputType(reflect.TypeOf((*OpenShiftContainerServiceVMSizePtrInput)(nil)).Elem(), OpenShiftContainerServiceVMSize("Standard_D2s_v3"))
	pulumi.RegisterOutputType(OSTypeOutput{})
	pulumi.RegisterOutputType(OSTypePtrOutput{})
	pulumi.RegisterOutputType(OpenShiftAgentPoolProfileRoleOutput{})
	pulumi.RegisterOutputType(OpenShiftAgentPoolProfileRolePtrOutput{})
	pulumi.RegisterOutputType(OpenShiftContainerServiceVMSizeOutput{})
	pulumi.RegisterOutputType(OpenShiftContainerServiceVMSizePtrOutput{})
}

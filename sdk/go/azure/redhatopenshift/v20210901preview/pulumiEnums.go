


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EncryptionAtHost string

const (
	EncryptionAtHostDisabled = EncryptionAtHost("Disabled")
	EncryptionAtHostEnabled  = EncryptionAtHost("Enabled")
)

func (EncryptionAtHost) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionAtHost)(nil)).Elem()
}

func (e EncryptionAtHost) ToEncryptionAtHostOutput() EncryptionAtHostOutput {
	return pulumi.ToOutput(e).(EncryptionAtHostOutput)
}

func (e EncryptionAtHost) ToEncryptionAtHostOutputWithContext(ctx context.Context) EncryptionAtHostOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EncryptionAtHostOutput)
}

func (e EncryptionAtHost) ToEncryptionAtHostPtrOutput() EncryptionAtHostPtrOutput {
	return e.ToEncryptionAtHostPtrOutputWithContext(context.Background())
}

func (e EncryptionAtHost) ToEncryptionAtHostPtrOutputWithContext(ctx context.Context) EncryptionAtHostPtrOutput {
	return EncryptionAtHost(e).ToEncryptionAtHostOutputWithContext(ctx).ToEncryptionAtHostPtrOutputWithContext(ctx)
}

func (e EncryptionAtHost) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncryptionAtHost) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncryptionAtHost) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EncryptionAtHost) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EncryptionAtHostOutput struct{ *pulumi.OutputState }

func (EncryptionAtHostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionAtHost)(nil)).Elem()
}

func (o EncryptionAtHostOutput) ToEncryptionAtHostOutput() EncryptionAtHostOutput {
	return o
}

func (o EncryptionAtHostOutput) ToEncryptionAtHostOutputWithContext(ctx context.Context) EncryptionAtHostOutput {
	return o
}

func (o EncryptionAtHostOutput) ToEncryptionAtHostPtrOutput() EncryptionAtHostPtrOutput {
	return o.ToEncryptionAtHostPtrOutputWithContext(context.Background())
}

func (o EncryptionAtHostOutput) ToEncryptionAtHostPtrOutputWithContext(ctx context.Context) EncryptionAtHostPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionAtHost) *EncryptionAtHost {
		return &v
	}).(EncryptionAtHostPtrOutput)
}

func (o EncryptionAtHostOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EncryptionAtHostOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncryptionAtHost) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EncryptionAtHostOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncryptionAtHostOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncryptionAtHost) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EncryptionAtHostPtrOutput struct{ *pulumi.OutputState }

func (EncryptionAtHostPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionAtHost)(nil)).Elem()
}

func (o EncryptionAtHostPtrOutput) ToEncryptionAtHostPtrOutput() EncryptionAtHostPtrOutput {
	return o
}

func (o EncryptionAtHostPtrOutput) ToEncryptionAtHostPtrOutputWithContext(ctx context.Context) EncryptionAtHostPtrOutput {
	return o
}

func (o EncryptionAtHostPtrOutput) Elem() EncryptionAtHostOutput {
	return o.ApplyT(func(v *EncryptionAtHost) EncryptionAtHost {
		if v != nil {
			return *v
		}
		var ret EncryptionAtHost
		return ret
	}).(EncryptionAtHostOutput)
}

func (o EncryptionAtHostPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncryptionAtHostPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EncryptionAtHost) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EncryptionAtHostInput interface {
	pulumi.Input

	ToEncryptionAtHostOutput() EncryptionAtHostOutput
	ToEncryptionAtHostOutputWithContext(context.Context) EncryptionAtHostOutput
}

var encryptionAtHostPtrType = reflect.TypeOf((**EncryptionAtHost)(nil)).Elem()

type EncryptionAtHostPtrInput interface {
	pulumi.Input

	ToEncryptionAtHostPtrOutput() EncryptionAtHostPtrOutput
	ToEncryptionAtHostPtrOutputWithContext(context.Context) EncryptionAtHostPtrOutput
}

type encryptionAtHostPtr string

func EncryptionAtHostPtr(v string) EncryptionAtHostPtrInput {
	return (*encryptionAtHostPtr)(&v)
}

func (*encryptionAtHostPtr) ElementType() reflect.Type {
	return encryptionAtHostPtrType
}

func (in *encryptionAtHostPtr) ToEncryptionAtHostPtrOutput() EncryptionAtHostPtrOutput {
	return pulumi.ToOutput(in).(EncryptionAtHostPtrOutput)
}

func (in *encryptionAtHostPtr) ToEncryptionAtHostPtrOutputWithContext(ctx context.Context) EncryptionAtHostPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EncryptionAtHostPtrOutput)
}

type SoftwareDefinedNetwork string

const (
	SoftwareDefinedNetworkOVNKubernetes = SoftwareDefinedNetwork("OVNKubernetes")
	SoftwareDefinedNetworkOpenShiftSDN  = SoftwareDefinedNetwork("OpenShiftSDN")
)

func (SoftwareDefinedNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((*SoftwareDefinedNetwork)(nil)).Elem()
}

func (e SoftwareDefinedNetwork) ToSoftwareDefinedNetworkOutput() SoftwareDefinedNetworkOutput {
	return pulumi.ToOutput(e).(SoftwareDefinedNetworkOutput)
}

func (e SoftwareDefinedNetwork) ToSoftwareDefinedNetworkOutputWithContext(ctx context.Context) SoftwareDefinedNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SoftwareDefinedNetworkOutput)
}

func (e SoftwareDefinedNetwork) ToSoftwareDefinedNetworkPtrOutput() SoftwareDefinedNetworkPtrOutput {
	return e.ToSoftwareDefinedNetworkPtrOutputWithContext(context.Background())
}

func (e SoftwareDefinedNetwork) ToSoftwareDefinedNetworkPtrOutputWithContext(ctx context.Context) SoftwareDefinedNetworkPtrOutput {
	return SoftwareDefinedNetwork(e).ToSoftwareDefinedNetworkOutputWithContext(ctx).ToSoftwareDefinedNetworkPtrOutputWithContext(ctx)
}

func (e SoftwareDefinedNetwork) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SoftwareDefinedNetwork) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SoftwareDefinedNetwork) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SoftwareDefinedNetwork) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SoftwareDefinedNetworkOutput struct{ *pulumi.OutputState }

func (SoftwareDefinedNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SoftwareDefinedNetwork)(nil)).Elem()
}

func (o SoftwareDefinedNetworkOutput) ToSoftwareDefinedNetworkOutput() SoftwareDefinedNetworkOutput {
	return o
}

func (o SoftwareDefinedNetworkOutput) ToSoftwareDefinedNetworkOutputWithContext(ctx context.Context) SoftwareDefinedNetworkOutput {
	return o
}

func (o SoftwareDefinedNetworkOutput) ToSoftwareDefinedNetworkPtrOutput() SoftwareDefinedNetworkPtrOutput {
	return o.ToSoftwareDefinedNetworkPtrOutputWithContext(context.Background())
}

func (o SoftwareDefinedNetworkOutput) ToSoftwareDefinedNetworkPtrOutputWithContext(ctx context.Context) SoftwareDefinedNetworkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SoftwareDefinedNetwork) *SoftwareDefinedNetwork {
		return &v
	}).(SoftwareDefinedNetworkPtrOutput)
}

func (o SoftwareDefinedNetworkOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SoftwareDefinedNetworkOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SoftwareDefinedNetwork) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SoftwareDefinedNetworkOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SoftwareDefinedNetworkOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SoftwareDefinedNetwork) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SoftwareDefinedNetworkPtrOutput struct{ *pulumi.OutputState }

func (SoftwareDefinedNetworkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SoftwareDefinedNetwork)(nil)).Elem()
}

func (o SoftwareDefinedNetworkPtrOutput) ToSoftwareDefinedNetworkPtrOutput() SoftwareDefinedNetworkPtrOutput {
	return o
}

func (o SoftwareDefinedNetworkPtrOutput) ToSoftwareDefinedNetworkPtrOutputWithContext(ctx context.Context) SoftwareDefinedNetworkPtrOutput {
	return o
}

func (o SoftwareDefinedNetworkPtrOutput) Elem() SoftwareDefinedNetworkOutput {
	return o.ApplyT(func(v *SoftwareDefinedNetwork) SoftwareDefinedNetwork {
		if v != nil {
			return *v
		}
		var ret SoftwareDefinedNetwork
		return ret
	}).(SoftwareDefinedNetworkOutput)
}

func (o SoftwareDefinedNetworkPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SoftwareDefinedNetworkPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SoftwareDefinedNetwork) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SoftwareDefinedNetworkInput interface {
	pulumi.Input

	ToSoftwareDefinedNetworkOutput() SoftwareDefinedNetworkOutput
	ToSoftwareDefinedNetworkOutputWithContext(context.Context) SoftwareDefinedNetworkOutput
}

var softwareDefinedNetworkPtrType = reflect.TypeOf((**SoftwareDefinedNetwork)(nil)).Elem()

type SoftwareDefinedNetworkPtrInput interface {
	pulumi.Input

	ToSoftwareDefinedNetworkPtrOutput() SoftwareDefinedNetworkPtrOutput
	ToSoftwareDefinedNetworkPtrOutputWithContext(context.Context) SoftwareDefinedNetworkPtrOutput
}

type softwareDefinedNetworkPtr string

func SoftwareDefinedNetworkPtr(v string) SoftwareDefinedNetworkPtrInput {
	return (*softwareDefinedNetworkPtr)(&v)
}

func (*softwareDefinedNetworkPtr) ElementType() reflect.Type {
	return softwareDefinedNetworkPtrType
}

func (in *softwareDefinedNetworkPtr) ToSoftwareDefinedNetworkPtrOutput() SoftwareDefinedNetworkPtrOutput {
	return pulumi.ToOutput(in).(SoftwareDefinedNetworkPtrOutput)
}

func (in *softwareDefinedNetworkPtr) ToSoftwareDefinedNetworkPtrOutputWithContext(ctx context.Context) SoftwareDefinedNetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SoftwareDefinedNetworkPtrOutput)
}

type VMSize string

const (
	VMSize_Standard_D16as_v4 = VMSize("Standard_D16as_v4")
	VMSize_Standard_D16s_v3  = VMSize("Standard_D16s_v3")
	VMSize_Standard_D2s_v3   = VMSize("Standard_D2s_v3")
	VMSize_Standard_D32as_v4 = VMSize("Standard_D32as_v4")
	VMSize_Standard_D32s_v3  = VMSize("Standard_D32s_v3")
	VMSize_Standard_D4as_v4  = VMSize("Standard_D4as_v4")
	VMSize_Standard_D4s_v3   = VMSize("Standard_D4s_v3")
	VMSize_Standard_D8as_v4  = VMSize("Standard_D8as_v4")
	VMSize_Standard_D8s_v3   = VMSize("Standard_D8s_v3")
	VMSize_Standard_E16s_v3  = VMSize("Standard_E16s_v3")
	VMSize_Standard_E32s_v3  = VMSize("Standard_E32s_v3")
	VMSize_Standard_E4s_v3   = VMSize("Standard_E4s_v3")
	VMSize_Standard_E64i_v3  = VMSize("Standard_E64i_v3")
	VMSize_Standard_E64is_v3 = VMSize("Standard_E64is_v3")
	VMSize_Standard_E8s_v3   = VMSize("Standard_E8s_v3")
	VMSize_Standard_F16s_v2  = VMSize("Standard_F16s_v2")
	VMSize_Standard_F32s_v2  = VMSize("Standard_F32s_v2")
	VMSize_Standard_F4s_v2   = VMSize("Standard_F4s_v2")
	VMSize_Standard_F72s_v2  = VMSize("Standard_F72s_v2")
	VMSize_Standard_F8s_v2   = VMSize("Standard_F8s_v2")
	VMSize_Standard_G5       = VMSize("Standard_G5")
	VMSize_Standard_GS5      = VMSize("Standard_GS5")
	VMSize_Standard_M128ms   = VMSize("Standard_M128ms")
)

func (VMSize) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSize)(nil)).Elem()
}

func (e VMSize) ToVMSizeOutput() VMSizeOutput {
	return pulumi.ToOutput(e).(VMSizeOutput)
}

func (e VMSize) ToVMSizeOutputWithContext(ctx context.Context) VMSizeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VMSizeOutput)
}

func (e VMSize) ToVMSizePtrOutput() VMSizePtrOutput {
	return e.ToVMSizePtrOutputWithContext(context.Background())
}

func (e VMSize) ToVMSizePtrOutputWithContext(ctx context.Context) VMSizePtrOutput {
	return VMSize(e).ToVMSizeOutputWithContext(ctx).ToVMSizePtrOutputWithContext(ctx)
}

func (e VMSize) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VMSize) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VMSize) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VMSize) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VMSizeOutput struct{ *pulumi.OutputState }

func (VMSizeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSize)(nil)).Elem()
}

func (o VMSizeOutput) ToVMSizeOutput() VMSizeOutput {
	return o
}

func (o VMSizeOutput) ToVMSizeOutputWithContext(ctx context.Context) VMSizeOutput {
	return o
}

func (o VMSizeOutput) ToVMSizePtrOutput() VMSizePtrOutput {
	return o.ToVMSizePtrOutputWithContext(context.Background())
}

func (o VMSizeOutput) ToVMSizePtrOutputWithContext(ctx context.Context) VMSizePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VMSize) *VMSize {
		return &v
	}).(VMSizePtrOutput)
}

func (o VMSizeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VMSizeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VMSize) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VMSizeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VMSizeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VMSize) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VMSizePtrOutput struct{ *pulumi.OutputState }

func (VMSizePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMSize)(nil)).Elem()
}

func (o VMSizePtrOutput) ToVMSizePtrOutput() VMSizePtrOutput {
	return o
}

func (o VMSizePtrOutput) ToVMSizePtrOutputWithContext(ctx context.Context) VMSizePtrOutput {
	return o
}

func (o VMSizePtrOutput) Elem() VMSizeOutput {
	return o.ApplyT(func(v *VMSize) VMSize {
		if v != nil {
			return *v
		}
		var ret VMSize
		return ret
	}).(VMSizeOutput)
}

func (o VMSizePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VMSizePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VMSize) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VMSizeInput interface {
	pulumi.Input

	ToVMSizeOutput() VMSizeOutput
	ToVMSizeOutputWithContext(context.Context) VMSizeOutput
}

var vmsizePtrType = reflect.TypeOf((**VMSize)(nil)).Elem()

type VMSizePtrInput interface {
	pulumi.Input

	ToVMSizePtrOutput() VMSizePtrOutput
	ToVMSizePtrOutputWithContext(context.Context) VMSizePtrOutput
}

type vmsizePtr string

func VMSizePtr(v string) VMSizePtrInput {
	return (*vmsizePtr)(&v)
}

func (*vmsizePtr) ElementType() reflect.Type {
	return vmsizePtrType
}

func (in *vmsizePtr) ToVMSizePtrOutput() VMSizePtrOutput {
	return pulumi.ToOutput(in).(VMSizePtrOutput)
}

func (in *vmsizePtr) ToVMSizePtrOutputWithContext(ctx context.Context) VMSizePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VMSizePtrOutput)
}

type Visibility string

const (
	VisibilityPrivate = Visibility("Private")
	VisibilityPublic  = Visibility("Public")
)

func (Visibility) ElementType() reflect.Type {
	return reflect.TypeOf((*Visibility)(nil)).Elem()
}

func (e Visibility) ToVisibilityOutput() VisibilityOutput {
	return pulumi.ToOutput(e).(VisibilityOutput)
}

func (e Visibility) ToVisibilityOutputWithContext(ctx context.Context) VisibilityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VisibilityOutput)
}

func (e Visibility) ToVisibilityPtrOutput() VisibilityPtrOutput {
	return e.ToVisibilityPtrOutputWithContext(context.Background())
}

func (e Visibility) ToVisibilityPtrOutputWithContext(ctx context.Context) VisibilityPtrOutput {
	return Visibility(e).ToVisibilityOutputWithContext(ctx).ToVisibilityPtrOutputWithContext(ctx)
}

func (e Visibility) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Visibility) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Visibility) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Visibility) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VisibilityOutput struct{ *pulumi.OutputState }

func (VisibilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Visibility)(nil)).Elem()
}

func (o VisibilityOutput) ToVisibilityOutput() VisibilityOutput {
	return o
}

func (o VisibilityOutput) ToVisibilityOutputWithContext(ctx context.Context) VisibilityOutput {
	return o
}

func (o VisibilityOutput) ToVisibilityPtrOutput() VisibilityPtrOutput {
	return o.ToVisibilityPtrOutputWithContext(context.Background())
}

func (o VisibilityOutput) ToVisibilityPtrOutputWithContext(ctx context.Context) VisibilityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Visibility) *Visibility {
		return &v
	}).(VisibilityPtrOutput)
}

func (o VisibilityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VisibilityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Visibility) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VisibilityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VisibilityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Visibility) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VisibilityPtrOutput struct{ *pulumi.OutputState }

func (VisibilityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Visibility)(nil)).Elem()
}

func (o VisibilityPtrOutput) ToVisibilityPtrOutput() VisibilityPtrOutput {
	return o
}

func (o VisibilityPtrOutput) ToVisibilityPtrOutputWithContext(ctx context.Context) VisibilityPtrOutput {
	return o
}

func (o VisibilityPtrOutput) Elem() VisibilityOutput {
	return o.ApplyT(func(v *Visibility) Visibility {
		if v != nil {
			return *v
		}
		var ret Visibility
		return ret
	}).(VisibilityOutput)
}

func (o VisibilityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VisibilityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Visibility) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VisibilityInput interface {
	pulumi.Input

	ToVisibilityOutput() VisibilityOutput
	ToVisibilityOutputWithContext(context.Context) VisibilityOutput
}

var visibilityPtrType = reflect.TypeOf((**Visibility)(nil)).Elem()

type VisibilityPtrInput interface {
	pulumi.Input

	ToVisibilityPtrOutput() VisibilityPtrOutput
	ToVisibilityPtrOutputWithContext(context.Context) VisibilityPtrOutput
}

type visibilityPtr string

func VisibilityPtr(v string) VisibilityPtrInput {
	return (*visibilityPtr)(&v)
}

func (*visibilityPtr) ElementType() reflect.Type {
	return visibilityPtrType
}

func (in *visibilityPtr) ToVisibilityPtrOutput() VisibilityPtrOutput {
	return pulumi.ToOutput(in).(VisibilityPtrOutput)
}

func (in *visibilityPtr) ToVisibilityPtrOutputWithContext(ctx context.Context) VisibilityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VisibilityPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EncryptionAtHostOutput{})
	pulumi.RegisterOutputType(EncryptionAtHostPtrOutput{})
	pulumi.RegisterOutputType(SoftwareDefinedNetworkOutput{})
	pulumi.RegisterOutputType(SoftwareDefinedNetworkPtrOutput{})
	pulumi.RegisterOutputType(VMSizeOutput{})
	pulumi.RegisterOutputType(VMSizePtrOutput{})
	pulumi.RegisterOutputType(VisibilityOutput{})
	pulumi.RegisterOutputType(VisibilityPtrOutput{})
}

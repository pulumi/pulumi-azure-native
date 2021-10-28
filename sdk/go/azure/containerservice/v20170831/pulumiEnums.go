


package v20170831

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContainerServiceStorageProfileTypes string

const (
	ContainerServiceStorageProfileTypesStorageAccount = ContainerServiceStorageProfileTypes("StorageAccount")
	ContainerServiceStorageProfileTypesManagedDisks   = ContainerServiceStorageProfileTypes("ManagedDisks")
)

func (ContainerServiceStorageProfileTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceStorageProfileTypes)(nil)).Elem()
}

func (e ContainerServiceStorageProfileTypes) ToContainerServiceStorageProfileTypesOutput() ContainerServiceStorageProfileTypesOutput {
	return pulumi.ToOutput(e).(ContainerServiceStorageProfileTypesOutput)
}

func (e ContainerServiceStorageProfileTypes) ToContainerServiceStorageProfileTypesOutputWithContext(ctx context.Context) ContainerServiceStorageProfileTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContainerServiceStorageProfileTypesOutput)
}

func (e ContainerServiceStorageProfileTypes) ToContainerServiceStorageProfileTypesPtrOutput() ContainerServiceStorageProfileTypesPtrOutput {
	return e.ToContainerServiceStorageProfileTypesPtrOutputWithContext(context.Background())
}

func (e ContainerServiceStorageProfileTypes) ToContainerServiceStorageProfileTypesPtrOutputWithContext(ctx context.Context) ContainerServiceStorageProfileTypesPtrOutput {
	return ContainerServiceStorageProfileTypes(e).ToContainerServiceStorageProfileTypesOutputWithContext(ctx).ToContainerServiceStorageProfileTypesPtrOutputWithContext(ctx)
}

func (e ContainerServiceStorageProfileTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerServiceStorageProfileTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerServiceStorageProfileTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContainerServiceStorageProfileTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContainerServiceStorageProfileTypesOutput struct{ *pulumi.OutputState }

func (ContainerServiceStorageProfileTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceStorageProfileTypes)(nil)).Elem()
}

func (o ContainerServiceStorageProfileTypesOutput) ToContainerServiceStorageProfileTypesOutput() ContainerServiceStorageProfileTypesOutput {
	return o
}

func (o ContainerServiceStorageProfileTypesOutput) ToContainerServiceStorageProfileTypesOutputWithContext(ctx context.Context) ContainerServiceStorageProfileTypesOutput {
	return o
}

func (o ContainerServiceStorageProfileTypesOutput) ToContainerServiceStorageProfileTypesPtrOutput() ContainerServiceStorageProfileTypesPtrOutput {
	return o.ToContainerServiceStorageProfileTypesPtrOutputWithContext(context.Background())
}

func (o ContainerServiceStorageProfileTypesOutput) ToContainerServiceStorageProfileTypesPtrOutputWithContext(ctx context.Context) ContainerServiceStorageProfileTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceStorageProfileTypes) *ContainerServiceStorageProfileTypes {
		return &v
	}).(ContainerServiceStorageProfileTypesPtrOutput)
}

func (o ContainerServiceStorageProfileTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContainerServiceStorageProfileTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerServiceStorageProfileTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContainerServiceStorageProfileTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerServiceStorageProfileTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerServiceStorageProfileTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContainerServiceStorageProfileTypesPtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceStorageProfileTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceStorageProfileTypes)(nil)).Elem()
}

func (o ContainerServiceStorageProfileTypesPtrOutput) ToContainerServiceStorageProfileTypesPtrOutput() ContainerServiceStorageProfileTypesPtrOutput {
	return o
}

func (o ContainerServiceStorageProfileTypesPtrOutput) ToContainerServiceStorageProfileTypesPtrOutputWithContext(ctx context.Context) ContainerServiceStorageProfileTypesPtrOutput {
	return o
}

func (o ContainerServiceStorageProfileTypesPtrOutput) Elem() ContainerServiceStorageProfileTypesOutput {
	return o.ApplyT(func(v *ContainerServiceStorageProfileTypes) ContainerServiceStorageProfileTypes {
		if v != nil {
			return *v
		}
		var ret ContainerServiceStorageProfileTypes
		return ret
	}).(ContainerServiceStorageProfileTypesOutput)
}

func (o ContainerServiceStorageProfileTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerServiceStorageProfileTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContainerServiceStorageProfileTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ContainerServiceStorageProfileTypesInput interface {
	pulumi.Input

	ToContainerServiceStorageProfileTypesOutput() ContainerServiceStorageProfileTypesOutput
	ToContainerServiceStorageProfileTypesOutputWithContext(context.Context) ContainerServiceStorageProfileTypesOutput
}

var containerServiceStorageProfileTypesPtrType = reflect.TypeOf((**ContainerServiceStorageProfileTypes)(nil)).Elem()

type ContainerServiceStorageProfileTypesPtrInput interface {
	pulumi.Input

	ToContainerServiceStorageProfileTypesPtrOutput() ContainerServiceStorageProfileTypesPtrOutput
	ToContainerServiceStorageProfileTypesPtrOutputWithContext(context.Context) ContainerServiceStorageProfileTypesPtrOutput
}

type containerServiceStorageProfileTypesPtr string

func ContainerServiceStorageProfileTypesPtr(v string) ContainerServiceStorageProfileTypesPtrInput {
	return (*containerServiceStorageProfileTypesPtr)(&v)
}

func (*containerServiceStorageProfileTypesPtr) ElementType() reflect.Type {
	return containerServiceStorageProfileTypesPtrType
}

func (in *containerServiceStorageProfileTypesPtr) ToContainerServiceStorageProfileTypesPtrOutput() ContainerServiceStorageProfileTypesPtrOutput {
	return pulumi.ToOutput(in).(ContainerServiceStorageProfileTypesPtrOutput)
}

func (in *containerServiceStorageProfileTypesPtr) ToContainerServiceStorageProfileTypesPtrOutputWithContext(ctx context.Context) ContainerServiceStorageProfileTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContainerServiceStorageProfileTypesPtrOutput)
}

type ContainerServiceVMSizeTypes string

const (
	ContainerServiceVMSizeTypes_Standard_A1            = ContainerServiceVMSizeTypes("Standard_A1")
	ContainerServiceVMSizeTypes_Standard_A10           = ContainerServiceVMSizeTypes("Standard_A10")
	ContainerServiceVMSizeTypes_Standard_A11           = ContainerServiceVMSizeTypes("Standard_A11")
	ContainerServiceVMSizeTypes_Standard_A1_v2         = ContainerServiceVMSizeTypes("Standard_A1_v2")
	ContainerServiceVMSizeTypes_Standard_A2            = ContainerServiceVMSizeTypes("Standard_A2")
	ContainerServiceVMSizeTypes_Standard_A2_v2         = ContainerServiceVMSizeTypes("Standard_A2_v2")
	ContainerServiceVMSizeTypes_Standard_A2m_v2        = ContainerServiceVMSizeTypes("Standard_A2m_v2")
	ContainerServiceVMSizeTypes_Standard_A3            = ContainerServiceVMSizeTypes("Standard_A3")
	ContainerServiceVMSizeTypes_Standard_A4            = ContainerServiceVMSizeTypes("Standard_A4")
	ContainerServiceVMSizeTypes_Standard_A4_v2         = ContainerServiceVMSizeTypes("Standard_A4_v2")
	ContainerServiceVMSizeTypes_Standard_A4m_v2        = ContainerServiceVMSizeTypes("Standard_A4m_v2")
	ContainerServiceVMSizeTypes_Standard_A5            = ContainerServiceVMSizeTypes("Standard_A5")
	ContainerServiceVMSizeTypes_Standard_A6            = ContainerServiceVMSizeTypes("Standard_A6")
	ContainerServiceVMSizeTypes_Standard_A7            = ContainerServiceVMSizeTypes("Standard_A7")
	ContainerServiceVMSizeTypes_Standard_A8            = ContainerServiceVMSizeTypes("Standard_A8")
	ContainerServiceVMSizeTypes_Standard_A8_v2         = ContainerServiceVMSizeTypes("Standard_A8_v2")
	ContainerServiceVMSizeTypes_Standard_A8m_v2        = ContainerServiceVMSizeTypes("Standard_A8m_v2")
	ContainerServiceVMSizeTypes_Standard_A9            = ContainerServiceVMSizeTypes("Standard_A9")
	ContainerServiceVMSizeTypes_Standard_B2ms          = ContainerServiceVMSizeTypes("Standard_B2ms")
	ContainerServiceVMSizeTypes_Standard_B2s           = ContainerServiceVMSizeTypes("Standard_B2s")
	ContainerServiceVMSizeTypes_Standard_B4ms          = ContainerServiceVMSizeTypes("Standard_B4ms")
	ContainerServiceVMSizeTypes_Standard_B8ms          = ContainerServiceVMSizeTypes("Standard_B8ms")
	ContainerServiceVMSizeTypes_Standard_D1            = ContainerServiceVMSizeTypes("Standard_D1")
	ContainerServiceVMSizeTypes_Standard_D11           = ContainerServiceVMSizeTypes("Standard_D11")
	ContainerServiceVMSizeTypes_Standard_D11_v2        = ContainerServiceVMSizeTypes("Standard_D11_v2")
	ContainerServiceVMSizeTypes_Standard_D11_v2_Promo  = ContainerServiceVMSizeTypes("Standard_D11_v2_Promo")
	ContainerServiceVMSizeTypes_Standard_D12           = ContainerServiceVMSizeTypes("Standard_D12")
	ContainerServiceVMSizeTypes_Standard_D12_v2        = ContainerServiceVMSizeTypes("Standard_D12_v2")
	ContainerServiceVMSizeTypes_Standard_D12_v2_Promo  = ContainerServiceVMSizeTypes("Standard_D12_v2_Promo")
	ContainerServiceVMSizeTypes_Standard_D13           = ContainerServiceVMSizeTypes("Standard_D13")
	ContainerServiceVMSizeTypes_Standard_D13_v2        = ContainerServiceVMSizeTypes("Standard_D13_v2")
	ContainerServiceVMSizeTypes_Standard_D13_v2_Promo  = ContainerServiceVMSizeTypes("Standard_D13_v2_Promo")
	ContainerServiceVMSizeTypes_Standard_D14           = ContainerServiceVMSizeTypes("Standard_D14")
	ContainerServiceVMSizeTypes_Standard_D14_v2        = ContainerServiceVMSizeTypes("Standard_D14_v2")
	ContainerServiceVMSizeTypes_Standard_D14_v2_Promo  = ContainerServiceVMSizeTypes("Standard_D14_v2_Promo")
	ContainerServiceVMSizeTypes_Standard_D15_v2        = ContainerServiceVMSizeTypes("Standard_D15_v2")
	ContainerServiceVMSizeTypes_Standard_D16_v3        = ContainerServiceVMSizeTypes("Standard_D16_v3")
	ContainerServiceVMSizeTypes_Standard_D16s_v3       = ContainerServiceVMSizeTypes("Standard_D16s_v3")
	ContainerServiceVMSizeTypes_Standard_D1_v2         = ContainerServiceVMSizeTypes("Standard_D1_v2")
	ContainerServiceVMSizeTypes_Standard_D2            = ContainerServiceVMSizeTypes("Standard_D2")
	ContainerServiceVMSizeTypes_Standard_D2_v2         = ContainerServiceVMSizeTypes("Standard_D2_v2")
	ContainerServiceVMSizeTypes_Standard_D2_v2_Promo   = ContainerServiceVMSizeTypes("Standard_D2_v2_Promo")
	ContainerServiceVMSizeTypes_Standard_D2_v3         = ContainerServiceVMSizeTypes("Standard_D2_v3")
	ContainerServiceVMSizeTypes_Standard_D2s_v3        = ContainerServiceVMSizeTypes("Standard_D2s_v3")
	ContainerServiceVMSizeTypes_Standard_D3            = ContainerServiceVMSizeTypes("Standard_D3")
	ContainerServiceVMSizeTypes_Standard_D32_v3        = ContainerServiceVMSizeTypes("Standard_D32_v3")
	ContainerServiceVMSizeTypes_Standard_D32s_v3       = ContainerServiceVMSizeTypes("Standard_D32s_v3")
	ContainerServiceVMSizeTypes_Standard_D3_v2         = ContainerServiceVMSizeTypes("Standard_D3_v2")
	ContainerServiceVMSizeTypes_Standard_D3_v2_Promo   = ContainerServiceVMSizeTypes("Standard_D3_v2_Promo")
	ContainerServiceVMSizeTypes_Standard_D4            = ContainerServiceVMSizeTypes("Standard_D4")
	ContainerServiceVMSizeTypes_Standard_D4_v2         = ContainerServiceVMSizeTypes("Standard_D4_v2")
	ContainerServiceVMSizeTypes_Standard_D4_v2_Promo   = ContainerServiceVMSizeTypes("Standard_D4_v2_Promo")
	ContainerServiceVMSizeTypes_Standard_D4_v3         = ContainerServiceVMSizeTypes("Standard_D4_v3")
	ContainerServiceVMSizeTypes_Standard_D4s_v3        = ContainerServiceVMSizeTypes("Standard_D4s_v3")
	ContainerServiceVMSizeTypes_Standard_D5_v2         = ContainerServiceVMSizeTypes("Standard_D5_v2")
	ContainerServiceVMSizeTypes_Standard_D5_v2_Promo   = ContainerServiceVMSizeTypes("Standard_D5_v2_Promo")
	ContainerServiceVMSizeTypes_Standard_D64_v3        = ContainerServiceVMSizeTypes("Standard_D64_v3")
	ContainerServiceVMSizeTypes_Standard_D64s_v3       = ContainerServiceVMSizeTypes("Standard_D64s_v3")
	ContainerServiceVMSizeTypes_Standard_D8_v3         = ContainerServiceVMSizeTypes("Standard_D8_v3")
	ContainerServiceVMSizeTypes_Standard_D8s_v3        = ContainerServiceVMSizeTypes("Standard_D8s_v3")
	ContainerServiceVMSizeTypes_Standard_DS1           = ContainerServiceVMSizeTypes("Standard_DS1")
	ContainerServiceVMSizeTypes_Standard_DS11          = ContainerServiceVMSizeTypes("Standard_DS11")
	ContainerServiceVMSizeTypes_Standard_DS11_v2       = ContainerServiceVMSizeTypes("Standard_DS11_v2")
	ContainerServiceVMSizeTypes_Standard_DS11_v2_Promo = ContainerServiceVMSizeTypes("Standard_DS11_v2_Promo")
	ContainerServiceVMSizeTypes_Standard_DS12          = ContainerServiceVMSizeTypes("Standard_DS12")
	ContainerServiceVMSizeTypes_Standard_DS12_v2       = ContainerServiceVMSizeTypes("Standard_DS12_v2")
	ContainerServiceVMSizeTypes_Standard_DS12_v2_Promo = ContainerServiceVMSizeTypes("Standard_DS12_v2_Promo")
	ContainerServiceVMSizeTypes_Standard_DS13          = ContainerServiceVMSizeTypes("Standard_DS13")
	ContainerServiceVMSizeTypes_Standard_DS13_2_v2     = ContainerServiceVMSizeTypes("Standard_DS13-2_v2")
	ContainerServiceVMSizeTypes_Standard_DS13_4_v2     = ContainerServiceVMSizeTypes("Standard_DS13-4_v2")
	ContainerServiceVMSizeTypes_Standard_DS13_v2       = ContainerServiceVMSizeTypes("Standard_DS13_v2")
	ContainerServiceVMSizeTypes_Standard_DS13_v2_Promo = ContainerServiceVMSizeTypes("Standard_DS13_v2_Promo")
	ContainerServiceVMSizeTypes_Standard_DS14          = ContainerServiceVMSizeTypes("Standard_DS14")
	ContainerServiceVMSizeTypes_Standard_DS14_4_v2     = ContainerServiceVMSizeTypes("Standard_DS14-4_v2")
	ContainerServiceVMSizeTypes_Standard_DS14_8_v2     = ContainerServiceVMSizeTypes("Standard_DS14-8_v2")
	ContainerServiceVMSizeTypes_Standard_DS14_v2       = ContainerServiceVMSizeTypes("Standard_DS14_v2")
	ContainerServiceVMSizeTypes_Standard_DS14_v2_Promo = ContainerServiceVMSizeTypes("Standard_DS14_v2_Promo")
	ContainerServiceVMSizeTypes_Standard_DS15_v2       = ContainerServiceVMSizeTypes("Standard_DS15_v2")
	ContainerServiceVMSizeTypes_Standard_DS1_v2        = ContainerServiceVMSizeTypes("Standard_DS1_v2")
	ContainerServiceVMSizeTypes_Standard_DS2           = ContainerServiceVMSizeTypes("Standard_DS2")
	ContainerServiceVMSizeTypes_Standard_DS2_v2        = ContainerServiceVMSizeTypes("Standard_DS2_v2")
	ContainerServiceVMSizeTypes_Standard_DS2_v2_Promo  = ContainerServiceVMSizeTypes("Standard_DS2_v2_Promo")
	ContainerServiceVMSizeTypes_Standard_DS3           = ContainerServiceVMSizeTypes("Standard_DS3")
	ContainerServiceVMSizeTypes_Standard_DS3_v2        = ContainerServiceVMSizeTypes("Standard_DS3_v2")
	ContainerServiceVMSizeTypes_Standard_DS3_v2_Promo  = ContainerServiceVMSizeTypes("Standard_DS3_v2_Promo")
	ContainerServiceVMSizeTypes_Standard_DS4           = ContainerServiceVMSizeTypes("Standard_DS4")
	ContainerServiceVMSizeTypes_Standard_DS4_v2        = ContainerServiceVMSizeTypes("Standard_DS4_v2")
	ContainerServiceVMSizeTypes_Standard_DS4_v2_Promo  = ContainerServiceVMSizeTypes("Standard_DS4_v2_Promo")
	ContainerServiceVMSizeTypes_Standard_DS5_v2        = ContainerServiceVMSizeTypes("Standard_DS5_v2")
	ContainerServiceVMSizeTypes_Standard_DS5_v2_Promo  = ContainerServiceVMSizeTypes("Standard_DS5_v2_Promo")
	ContainerServiceVMSizeTypes_Standard_E16_v3        = ContainerServiceVMSizeTypes("Standard_E16_v3")
	ContainerServiceVMSizeTypes_Standard_E16s_v3       = ContainerServiceVMSizeTypes("Standard_E16s_v3")
	ContainerServiceVMSizeTypes_Standard_E2_v3         = ContainerServiceVMSizeTypes("Standard_E2_v3")
	ContainerServiceVMSizeTypes_Standard_E2s_v3        = ContainerServiceVMSizeTypes("Standard_E2s_v3")
	ContainerServiceVMSizeTypes_Standard_E32_16s_v3    = ContainerServiceVMSizeTypes("Standard_E32-16s_v3")
	ContainerServiceVMSizeTypes_Standard_E32_8s_v3     = ContainerServiceVMSizeTypes("Standard_E32-8s_v3")
	ContainerServiceVMSizeTypes_Standard_E32_v3        = ContainerServiceVMSizeTypes("Standard_E32_v3")
	ContainerServiceVMSizeTypes_Standard_E32s_v3       = ContainerServiceVMSizeTypes("Standard_E32s_v3")
	ContainerServiceVMSizeTypes_Standard_E4_v3         = ContainerServiceVMSizeTypes("Standard_E4_v3")
	ContainerServiceVMSizeTypes_Standard_E4s_v3        = ContainerServiceVMSizeTypes("Standard_E4s_v3")
	ContainerServiceVMSizeTypes_Standard_E64_16s_v3    = ContainerServiceVMSizeTypes("Standard_E64-16s_v3")
	ContainerServiceVMSizeTypes_Standard_E64_32s_v3    = ContainerServiceVMSizeTypes("Standard_E64-32s_v3")
	ContainerServiceVMSizeTypes_Standard_E64_v3        = ContainerServiceVMSizeTypes("Standard_E64_v3")
	ContainerServiceVMSizeTypes_Standard_E64s_v3       = ContainerServiceVMSizeTypes("Standard_E64s_v3")
	ContainerServiceVMSizeTypes_Standard_E8_v3         = ContainerServiceVMSizeTypes("Standard_E8_v3")
	ContainerServiceVMSizeTypes_Standard_E8s_v3        = ContainerServiceVMSizeTypes("Standard_E8s_v3")
	ContainerServiceVMSizeTypes_Standard_F1            = ContainerServiceVMSizeTypes("Standard_F1")
	ContainerServiceVMSizeTypes_Standard_F16           = ContainerServiceVMSizeTypes("Standard_F16")
	ContainerServiceVMSizeTypes_Standard_F16s          = ContainerServiceVMSizeTypes("Standard_F16s")
	ContainerServiceVMSizeTypes_Standard_F16s_v2       = ContainerServiceVMSizeTypes("Standard_F16s_v2")
	ContainerServiceVMSizeTypes_Standard_F1s           = ContainerServiceVMSizeTypes("Standard_F1s")
	ContainerServiceVMSizeTypes_Standard_F2            = ContainerServiceVMSizeTypes("Standard_F2")
	ContainerServiceVMSizeTypes_Standard_F2s           = ContainerServiceVMSizeTypes("Standard_F2s")
	ContainerServiceVMSizeTypes_Standard_F2s_v2        = ContainerServiceVMSizeTypes("Standard_F2s_v2")
	ContainerServiceVMSizeTypes_Standard_F32s_v2       = ContainerServiceVMSizeTypes("Standard_F32s_v2")
	ContainerServiceVMSizeTypes_Standard_F4            = ContainerServiceVMSizeTypes("Standard_F4")
	ContainerServiceVMSizeTypes_Standard_F4s           = ContainerServiceVMSizeTypes("Standard_F4s")
	ContainerServiceVMSizeTypes_Standard_F4s_v2        = ContainerServiceVMSizeTypes("Standard_F4s_v2")
	ContainerServiceVMSizeTypes_Standard_F64s_v2       = ContainerServiceVMSizeTypes("Standard_F64s_v2")
	ContainerServiceVMSizeTypes_Standard_F72s_v2       = ContainerServiceVMSizeTypes("Standard_F72s_v2")
	ContainerServiceVMSizeTypes_Standard_F8            = ContainerServiceVMSizeTypes("Standard_F8")
	ContainerServiceVMSizeTypes_Standard_F8s           = ContainerServiceVMSizeTypes("Standard_F8s")
	ContainerServiceVMSizeTypes_Standard_F8s_v2        = ContainerServiceVMSizeTypes("Standard_F8s_v2")
	ContainerServiceVMSizeTypes_Standard_G1            = ContainerServiceVMSizeTypes("Standard_G1")
	ContainerServiceVMSizeTypes_Standard_G2            = ContainerServiceVMSizeTypes("Standard_G2")
	ContainerServiceVMSizeTypes_Standard_G3            = ContainerServiceVMSizeTypes("Standard_G3")
	ContainerServiceVMSizeTypes_Standard_G4            = ContainerServiceVMSizeTypes("Standard_G4")
	ContainerServiceVMSizeTypes_Standard_G5            = ContainerServiceVMSizeTypes("Standard_G5")
	ContainerServiceVMSizeTypes_Standard_GS1           = ContainerServiceVMSizeTypes("Standard_GS1")
	ContainerServiceVMSizeTypes_Standard_GS2           = ContainerServiceVMSizeTypes("Standard_GS2")
	ContainerServiceVMSizeTypes_Standard_GS3           = ContainerServiceVMSizeTypes("Standard_GS3")
	ContainerServiceVMSizeTypes_Standard_GS4           = ContainerServiceVMSizeTypes("Standard_GS4")
	ContainerServiceVMSizeTypes_Standard_GS4_4         = ContainerServiceVMSizeTypes("Standard_GS4-4")
	ContainerServiceVMSizeTypes_Standard_GS4_8         = ContainerServiceVMSizeTypes("Standard_GS4-8")
	ContainerServiceVMSizeTypes_Standard_GS5           = ContainerServiceVMSizeTypes("Standard_GS5")
	ContainerServiceVMSizeTypes_Standard_GS5_16        = ContainerServiceVMSizeTypes("Standard_GS5-16")
	ContainerServiceVMSizeTypes_Standard_GS5_8         = ContainerServiceVMSizeTypes("Standard_GS5-8")
	ContainerServiceVMSizeTypes_Standard_H16           = ContainerServiceVMSizeTypes("Standard_H16")
	ContainerServiceVMSizeTypes_Standard_H16m          = ContainerServiceVMSizeTypes("Standard_H16m")
	ContainerServiceVMSizeTypes_Standard_H16mr         = ContainerServiceVMSizeTypes("Standard_H16mr")
	ContainerServiceVMSizeTypes_Standard_H16r          = ContainerServiceVMSizeTypes("Standard_H16r")
	ContainerServiceVMSizeTypes_Standard_H8            = ContainerServiceVMSizeTypes("Standard_H8")
	ContainerServiceVMSizeTypes_Standard_H8m           = ContainerServiceVMSizeTypes("Standard_H8m")
	ContainerServiceVMSizeTypes_Standard_L16s          = ContainerServiceVMSizeTypes("Standard_L16s")
	ContainerServiceVMSizeTypes_Standard_L32s          = ContainerServiceVMSizeTypes("Standard_L32s")
	ContainerServiceVMSizeTypes_Standard_L4s           = ContainerServiceVMSizeTypes("Standard_L4s")
	ContainerServiceVMSizeTypes_Standard_L8s           = ContainerServiceVMSizeTypes("Standard_L8s")
	ContainerServiceVMSizeTypes_Standard_M128_32ms     = ContainerServiceVMSizeTypes("Standard_M128-32ms")
	ContainerServiceVMSizeTypes_Standard_M128_64ms     = ContainerServiceVMSizeTypes("Standard_M128-64ms")
	ContainerServiceVMSizeTypes_Standard_M128ms        = ContainerServiceVMSizeTypes("Standard_M128ms")
	ContainerServiceVMSizeTypes_Standard_M128s         = ContainerServiceVMSizeTypes("Standard_M128s")
	ContainerServiceVMSizeTypes_Standard_M64_16ms      = ContainerServiceVMSizeTypes("Standard_M64-16ms")
	ContainerServiceVMSizeTypes_Standard_M64_32ms      = ContainerServiceVMSizeTypes("Standard_M64-32ms")
	ContainerServiceVMSizeTypes_Standard_M64ms         = ContainerServiceVMSizeTypes("Standard_M64ms")
	ContainerServiceVMSizeTypes_Standard_M64s          = ContainerServiceVMSizeTypes("Standard_M64s")
	ContainerServiceVMSizeTypes_Standard_NC12          = ContainerServiceVMSizeTypes("Standard_NC12")
	ContainerServiceVMSizeTypes_Standard_NC12s_v2      = ContainerServiceVMSizeTypes("Standard_NC12s_v2")
	ContainerServiceVMSizeTypes_Standard_NC12s_v3      = ContainerServiceVMSizeTypes("Standard_NC12s_v3")
	ContainerServiceVMSizeTypes_Standard_NC24          = ContainerServiceVMSizeTypes("Standard_NC24")
	ContainerServiceVMSizeTypes_Standard_NC24r         = ContainerServiceVMSizeTypes("Standard_NC24r")
	ContainerServiceVMSizeTypes_Standard_NC24rs_v2     = ContainerServiceVMSizeTypes("Standard_NC24rs_v2")
	ContainerServiceVMSizeTypes_Standard_NC24rs_v3     = ContainerServiceVMSizeTypes("Standard_NC24rs_v3")
	ContainerServiceVMSizeTypes_Standard_NC24s_v2      = ContainerServiceVMSizeTypes("Standard_NC24s_v2")
	ContainerServiceVMSizeTypes_Standard_NC24s_v3      = ContainerServiceVMSizeTypes("Standard_NC24s_v3")
	ContainerServiceVMSizeTypes_Standard_NC6           = ContainerServiceVMSizeTypes("Standard_NC6")
	ContainerServiceVMSizeTypes_Standard_NC6s_v2       = ContainerServiceVMSizeTypes("Standard_NC6s_v2")
	ContainerServiceVMSizeTypes_Standard_NC6s_v3       = ContainerServiceVMSizeTypes("Standard_NC6s_v3")
	ContainerServiceVMSizeTypes_Standard_ND12s         = ContainerServiceVMSizeTypes("Standard_ND12s")
	ContainerServiceVMSizeTypes_Standard_ND24rs        = ContainerServiceVMSizeTypes("Standard_ND24rs")
	ContainerServiceVMSizeTypes_Standard_ND24s         = ContainerServiceVMSizeTypes("Standard_ND24s")
	ContainerServiceVMSizeTypes_Standard_ND6s          = ContainerServiceVMSizeTypes("Standard_ND6s")
	ContainerServiceVMSizeTypes_Standard_NV12          = ContainerServiceVMSizeTypes("Standard_NV12")
	ContainerServiceVMSizeTypes_Standard_NV24          = ContainerServiceVMSizeTypes("Standard_NV24")
	ContainerServiceVMSizeTypes_Standard_NV6           = ContainerServiceVMSizeTypes("Standard_NV6")
)

func (ContainerServiceVMSizeTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceVMSizeTypes)(nil)).Elem()
}

func (e ContainerServiceVMSizeTypes) ToContainerServiceVMSizeTypesOutput() ContainerServiceVMSizeTypesOutput {
	return pulumi.ToOutput(e).(ContainerServiceVMSizeTypesOutput)
}

func (e ContainerServiceVMSizeTypes) ToContainerServiceVMSizeTypesOutputWithContext(ctx context.Context) ContainerServiceVMSizeTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContainerServiceVMSizeTypesOutput)
}

func (e ContainerServiceVMSizeTypes) ToContainerServiceVMSizeTypesPtrOutput() ContainerServiceVMSizeTypesPtrOutput {
	return e.ToContainerServiceVMSizeTypesPtrOutputWithContext(context.Background())
}

func (e ContainerServiceVMSizeTypes) ToContainerServiceVMSizeTypesPtrOutputWithContext(ctx context.Context) ContainerServiceVMSizeTypesPtrOutput {
	return ContainerServiceVMSizeTypes(e).ToContainerServiceVMSizeTypesOutputWithContext(ctx).ToContainerServiceVMSizeTypesPtrOutputWithContext(ctx)
}

func (e ContainerServiceVMSizeTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerServiceVMSizeTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerServiceVMSizeTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContainerServiceVMSizeTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContainerServiceVMSizeTypesOutput struct{ *pulumi.OutputState }

func (ContainerServiceVMSizeTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceVMSizeTypes)(nil)).Elem()
}

func (o ContainerServiceVMSizeTypesOutput) ToContainerServiceVMSizeTypesOutput() ContainerServiceVMSizeTypesOutput {
	return o
}

func (o ContainerServiceVMSizeTypesOutput) ToContainerServiceVMSizeTypesOutputWithContext(ctx context.Context) ContainerServiceVMSizeTypesOutput {
	return o
}

func (o ContainerServiceVMSizeTypesOutput) ToContainerServiceVMSizeTypesPtrOutput() ContainerServiceVMSizeTypesPtrOutput {
	return o.ToContainerServiceVMSizeTypesPtrOutputWithContext(context.Background())
}

func (o ContainerServiceVMSizeTypesOutput) ToContainerServiceVMSizeTypesPtrOutputWithContext(ctx context.Context) ContainerServiceVMSizeTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceVMSizeTypes) *ContainerServiceVMSizeTypes {
		return &v
	}).(ContainerServiceVMSizeTypesPtrOutput)
}

func (o ContainerServiceVMSizeTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContainerServiceVMSizeTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerServiceVMSizeTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContainerServiceVMSizeTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerServiceVMSizeTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerServiceVMSizeTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContainerServiceVMSizeTypesPtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceVMSizeTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceVMSizeTypes)(nil)).Elem()
}

func (o ContainerServiceVMSizeTypesPtrOutput) ToContainerServiceVMSizeTypesPtrOutput() ContainerServiceVMSizeTypesPtrOutput {
	return o
}

func (o ContainerServiceVMSizeTypesPtrOutput) ToContainerServiceVMSizeTypesPtrOutputWithContext(ctx context.Context) ContainerServiceVMSizeTypesPtrOutput {
	return o
}

func (o ContainerServiceVMSizeTypesPtrOutput) Elem() ContainerServiceVMSizeTypesOutput {
	return o.ApplyT(func(v *ContainerServiceVMSizeTypes) ContainerServiceVMSizeTypes {
		if v != nil {
			return *v
		}
		var ret ContainerServiceVMSizeTypes
		return ret
	}).(ContainerServiceVMSizeTypesOutput)
}

func (o ContainerServiceVMSizeTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerServiceVMSizeTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContainerServiceVMSizeTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ContainerServiceVMSizeTypesInput interface {
	pulumi.Input

	ToContainerServiceVMSizeTypesOutput() ContainerServiceVMSizeTypesOutput
	ToContainerServiceVMSizeTypesOutputWithContext(context.Context) ContainerServiceVMSizeTypesOutput
}

var containerServiceVMSizeTypesPtrType = reflect.TypeOf((**ContainerServiceVMSizeTypes)(nil)).Elem()

type ContainerServiceVMSizeTypesPtrInput interface {
	pulumi.Input

	ToContainerServiceVMSizeTypesPtrOutput() ContainerServiceVMSizeTypesPtrOutput
	ToContainerServiceVMSizeTypesPtrOutputWithContext(context.Context) ContainerServiceVMSizeTypesPtrOutput
}

type containerServiceVMSizeTypesPtr string

func ContainerServiceVMSizeTypesPtr(v string) ContainerServiceVMSizeTypesPtrInput {
	return (*containerServiceVMSizeTypesPtr)(&v)
}

func (*containerServiceVMSizeTypesPtr) ElementType() reflect.Type {
	return containerServiceVMSizeTypesPtrType
}

func (in *containerServiceVMSizeTypesPtr) ToContainerServiceVMSizeTypesPtrOutput() ContainerServiceVMSizeTypesPtrOutput {
	return pulumi.ToOutput(in).(ContainerServiceVMSizeTypesPtrOutput)
}

func (in *containerServiceVMSizeTypesPtr) ToContainerServiceVMSizeTypesPtrOutputWithContext(ctx context.Context) ContainerServiceVMSizeTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContainerServiceVMSizeTypesPtrOutput)
}

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

func init() {
	pulumi.RegisterOutputType(ContainerServiceStorageProfileTypesOutput{})
	pulumi.RegisterOutputType(ContainerServiceStorageProfileTypesPtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceVMSizeTypesOutput{})
	pulumi.RegisterOutputType(ContainerServiceVMSizeTypesPtrOutput{})
	pulumi.RegisterOutputType(OSTypeOutput{})
	pulumi.RegisterOutputType(OSTypePtrOutput{})
}




package hanaonazure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Disk struct {
	DiskSizeGB *int    `pulumi:"diskSizeGB"`
	Name       *string `pulumi:"name"`
}





type DiskInput interface {
	pulumi.Input

	ToDiskOutput() DiskOutput
	ToDiskOutputWithContext(context.Context) DiskOutput
}

type DiskArgs struct {
	DiskSizeGB pulumi.IntPtrInput    `pulumi:"diskSizeGB"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
}

func (DiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Disk)(nil)).Elem()
}

func (i DiskArgs) ToDiskOutput() DiskOutput {
	return i.ToDiskOutputWithContext(context.Background())
}

func (i DiskArgs) ToDiskOutputWithContext(ctx context.Context) DiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskOutput)
}





type DiskArrayInput interface {
	pulumi.Input

	ToDiskArrayOutput() DiskArrayOutput
	ToDiskArrayOutputWithContext(context.Context) DiskArrayOutput
}

type DiskArray []DiskInput

func (DiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Disk)(nil)).Elem()
}

func (i DiskArray) ToDiskArrayOutput() DiskArrayOutput {
	return i.ToDiskArrayOutputWithContext(context.Background())
}

func (i DiskArray) ToDiskArrayOutputWithContext(ctx context.Context) DiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskArrayOutput)
}

type DiskOutput struct{ *pulumi.OutputState }

func (DiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Disk)(nil)).Elem()
}

func (o DiskOutput) ToDiskOutput() DiskOutput {
	return o
}

func (o DiskOutput) ToDiskOutputWithContext(ctx context.Context) DiskOutput {
	return o
}

func (o DiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Disk) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o DiskOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Disk) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DiskArrayOutput struct{ *pulumi.OutputState }

func (DiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Disk)(nil)).Elem()
}

func (o DiskArrayOutput) ToDiskArrayOutput() DiskArrayOutput {
	return o
}

func (o DiskArrayOutput) ToDiskArrayOutputWithContext(ctx context.Context) DiskArrayOutput {
	return o
}

func (o DiskArrayOutput) Index(i pulumi.IntInput) DiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Disk {
		return vs[0].([]Disk)[vs[1].(int)]
	}).(DiskOutput)
}

type DiskResponse struct {
	DiskSizeGB *int    `pulumi:"diskSizeGB"`
	Lun        int     `pulumi:"lun"`
	Name       *string `pulumi:"name"`
}

type DiskResponseOutput struct{ *pulumi.OutputState }

func (DiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskResponse)(nil)).Elem()
}

func (o DiskResponseOutput) ToDiskResponseOutput() DiskResponseOutput {
	return o
}

func (o DiskResponseOutput) ToDiskResponseOutputWithContext(ctx context.Context) DiskResponseOutput {
	return o
}

func (o DiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o DiskResponseOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v DiskResponse) int { return v.Lun }).(pulumi.IntOutput)
}

func (o DiskResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DiskResponseArrayOutput struct{ *pulumi.OutputState }

func (DiskResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskResponse)(nil)).Elem()
}

func (o DiskResponseArrayOutput) ToDiskResponseArrayOutput() DiskResponseArrayOutput {
	return o
}

func (o DiskResponseArrayOutput) ToDiskResponseArrayOutputWithContext(ctx context.Context) DiskResponseArrayOutput {
	return o
}

func (o DiskResponseArrayOutput) Index(i pulumi.IntInput) DiskResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiskResponse {
		return vs[0].([]DiskResponse)[vs[1].(int)]
	}).(DiskResponseOutput)
}

type HardwareProfileResponse struct {
	HanaInstanceSize string `pulumi:"hanaInstanceSize"`
	HardwareType     string `pulumi:"hardwareType"`
}

type HardwareProfileResponseOutput struct{ *pulumi.OutputState }

func (HardwareProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HardwareProfileResponse)(nil)).Elem()
}

func (o HardwareProfileResponseOutput) ToHardwareProfileResponseOutput() HardwareProfileResponseOutput {
	return o
}

func (o HardwareProfileResponseOutput) ToHardwareProfileResponseOutputWithContext(ctx context.Context) HardwareProfileResponseOutput {
	return o
}

func (o HardwareProfileResponseOutput) HanaInstanceSize() pulumi.StringOutput {
	return o.ApplyT(func(v HardwareProfileResponse) string { return v.HanaInstanceSize }).(pulumi.StringOutput)
}

func (o HardwareProfileResponseOutput) HardwareType() pulumi.StringOutput {
	return o.ApplyT(func(v HardwareProfileResponse) string { return v.HardwareType }).(pulumi.StringOutput)
}

type HardwareProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (HardwareProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardwareProfileResponse)(nil)).Elem()
}

func (o HardwareProfileResponsePtrOutput) ToHardwareProfileResponsePtrOutput() HardwareProfileResponsePtrOutput {
	return o
}

func (o HardwareProfileResponsePtrOutput) ToHardwareProfileResponsePtrOutputWithContext(ctx context.Context) HardwareProfileResponsePtrOutput {
	return o
}

func (o HardwareProfileResponsePtrOutput) Elem() HardwareProfileResponseOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) HardwareProfileResponse {
		if v != nil {
			return *v
		}
		var ret HardwareProfileResponse
		return ret
	}).(HardwareProfileResponseOutput)
}

func (o HardwareProfileResponsePtrOutput) HanaInstanceSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HanaInstanceSize
	}).(pulumi.StringPtrOutput)
}

func (o HardwareProfileResponsePtrOutput) HardwareType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HardwareType
	}).(pulumi.StringPtrOutput)
}

type IpAddress struct {
	IpAddress *string `pulumi:"ipAddress"`
}





type IpAddressInput interface {
	pulumi.Input

	ToIpAddressOutput() IpAddressOutput
	ToIpAddressOutputWithContext(context.Context) IpAddressOutput
}

type IpAddressArgs struct {
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
}

func (IpAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddress)(nil)).Elem()
}

func (i IpAddressArgs) ToIpAddressOutput() IpAddressOutput {
	return i.ToIpAddressOutputWithContext(context.Background())
}

func (i IpAddressArgs) ToIpAddressOutputWithContext(ctx context.Context) IpAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressOutput)
}





type IpAddressArrayInput interface {
	pulumi.Input

	ToIpAddressArrayOutput() IpAddressArrayOutput
	ToIpAddressArrayOutputWithContext(context.Context) IpAddressArrayOutput
}

type IpAddressArray []IpAddressInput

func (IpAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpAddress)(nil)).Elem()
}

func (i IpAddressArray) ToIpAddressArrayOutput() IpAddressArrayOutput {
	return i.ToIpAddressArrayOutputWithContext(context.Background())
}

func (i IpAddressArray) ToIpAddressArrayOutputWithContext(ctx context.Context) IpAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressArrayOutput)
}

type IpAddressOutput struct{ *pulumi.OutputState }

func (IpAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddress)(nil)).Elem()
}

func (o IpAddressOutput) ToIpAddressOutput() IpAddressOutput {
	return o
}

func (o IpAddressOutput) ToIpAddressOutputWithContext(ctx context.Context) IpAddressOutput {
	return o
}

func (o IpAddressOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpAddress) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

type IpAddressArrayOutput struct{ *pulumi.OutputState }

func (IpAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpAddress)(nil)).Elem()
}

func (o IpAddressArrayOutput) ToIpAddressArrayOutput() IpAddressArrayOutput {
	return o
}

func (o IpAddressArrayOutput) ToIpAddressArrayOutputWithContext(ctx context.Context) IpAddressArrayOutput {
	return o
}

func (o IpAddressArrayOutput) Index(i pulumi.IntInput) IpAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpAddress {
		return vs[0].([]IpAddress)[vs[1].(int)]
	}).(IpAddressOutput)
}

type IpAddressResponse struct {
	IpAddress *string `pulumi:"ipAddress"`
}

type IpAddressResponseOutput struct{ *pulumi.OutputState }

func (IpAddressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddressResponse)(nil)).Elem()
}

func (o IpAddressResponseOutput) ToIpAddressResponseOutput() IpAddressResponseOutput {
	return o
}

func (o IpAddressResponseOutput) ToIpAddressResponseOutputWithContext(ctx context.Context) IpAddressResponseOutput {
	return o
}

func (o IpAddressResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpAddressResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

type IpAddressResponseArrayOutput struct{ *pulumi.OutputState }

func (IpAddressResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpAddressResponse)(nil)).Elem()
}

func (o IpAddressResponseArrayOutput) ToIpAddressResponseArrayOutput() IpAddressResponseArrayOutput {
	return o
}

func (o IpAddressResponseArrayOutput) ToIpAddressResponseArrayOutputWithContext(ctx context.Context) IpAddressResponseArrayOutput {
	return o
}

func (o IpAddressResponseArrayOutput) Index(i pulumi.IntInput) IpAddressResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpAddressResponse {
		return vs[0].([]IpAddressResponse)[vs[1].(int)]
	}).(IpAddressResponseOutput)
}

type NetworkProfile struct {
	NetworkInterfaces []IpAddress `pulumi:"networkInterfaces"`
}





type NetworkProfileInput interface {
	pulumi.Input

	ToNetworkProfileOutput() NetworkProfileOutput
	ToNetworkProfileOutputWithContext(context.Context) NetworkProfileOutput
}

type NetworkProfileArgs struct {
	NetworkInterfaces IpAddressArrayInput `pulumi:"networkInterfaces"`
}

func (NetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (i NetworkProfileArgs) ToNetworkProfileOutput() NetworkProfileOutput {
	return i.ToNetworkProfileOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput)
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput).ToNetworkProfilePtrOutputWithContext(ctx)
}









type NetworkProfilePtrInput interface {
	pulumi.Input

	ToNetworkProfilePtrOutput() NetworkProfilePtrOutput
	ToNetworkProfilePtrOutputWithContext(context.Context) NetworkProfilePtrOutput
}

type networkProfilePtrType NetworkProfileArgs

func NetworkProfilePtr(v *NetworkProfileArgs) NetworkProfilePtrInput {
	return (*networkProfilePtrType)(v)
}

func (*networkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfilePtrOutput)
}

type NetworkProfileOutput struct{ *pulumi.OutputState }

func (NetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (o NetworkProfileOutput) ToNetworkProfileOutput() NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkProfile) *NetworkProfile {
		return &v
	}).(NetworkProfilePtrOutput)
}

func (o NetworkProfileOutput) NetworkInterfaces() IpAddressArrayOutput {
	return o.ApplyT(func(v NetworkProfile) []IpAddress { return v.NetworkInterfaces }).(IpAddressArrayOutput)
}

type NetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) Elem() NetworkProfileOutput {
	return o.ApplyT(func(v *NetworkProfile) NetworkProfile {
		if v != nil {
			return *v
		}
		var ret NetworkProfile
		return ret
	}).(NetworkProfileOutput)
}

func (o NetworkProfilePtrOutput) NetworkInterfaces() IpAddressArrayOutput {
	return o.ApplyT(func(v *NetworkProfile) []IpAddress {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(IpAddressArrayOutput)
}

type NetworkProfileResponse struct {
	CircuitId         string              `pulumi:"circuitId"`
	NetworkInterfaces []IpAddressResponse `pulumi:"networkInterfaces"`
}

type NetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutput() NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutputWithContext(ctx context.Context) NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) CircuitId() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkProfileResponse) string { return v.CircuitId }).(pulumi.StringOutput)
}

func (o NetworkProfileResponseOutput) NetworkInterfaces() IpAddressResponseArrayOutput {
	return o.ApplyT(func(v NetworkProfileResponse) []IpAddressResponse { return v.NetworkInterfaces }).(IpAddressResponseArrayOutput)
}

type NetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutputWithContext(ctx context.Context) NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) Elem() NetworkProfileResponseOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) NetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret NetworkProfileResponse
		return ret
	}).(NetworkProfileResponseOutput)
}

func (o NetworkProfileResponsePtrOutput) CircuitId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CircuitId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) NetworkInterfaces() IpAddressResponseArrayOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) []IpAddressResponse {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(IpAddressResponseArrayOutput)
}

type OSProfile struct {
	ComputerName *string `pulumi:"computerName"`
	SshPublicKey *string `pulumi:"sshPublicKey"`
}





type OSProfileInput interface {
	pulumi.Input

	ToOSProfileOutput() OSProfileOutput
	ToOSProfileOutputWithContext(context.Context) OSProfileOutput
}

type OSProfileArgs struct {
	ComputerName pulumi.StringPtrInput `pulumi:"computerName"`
	SshPublicKey pulumi.StringPtrInput `pulumi:"sshPublicKey"`
}

func (OSProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfile)(nil)).Elem()
}

func (i OSProfileArgs) ToOSProfileOutput() OSProfileOutput {
	return i.ToOSProfileOutputWithContext(context.Background())
}

func (i OSProfileArgs) ToOSProfileOutputWithContext(ctx context.Context) OSProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileOutput)
}

func (i OSProfileArgs) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return i.ToOSProfilePtrOutputWithContext(context.Background())
}

func (i OSProfileArgs) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileOutput).ToOSProfilePtrOutputWithContext(ctx)
}









type OSProfilePtrInput interface {
	pulumi.Input

	ToOSProfilePtrOutput() OSProfilePtrOutput
	ToOSProfilePtrOutputWithContext(context.Context) OSProfilePtrOutput
}

type osprofilePtrType OSProfileArgs

func OSProfilePtr(v *OSProfileArgs) OSProfilePtrInput {
	return (*osprofilePtrType)(v)
}

func (*osprofilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfile)(nil)).Elem()
}

func (i *osprofilePtrType) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return i.ToOSProfilePtrOutputWithContext(context.Background())
}

func (i *osprofilePtrType) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfilePtrOutput)
}

type OSProfileOutput struct{ *pulumi.OutputState }

func (OSProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfile)(nil)).Elem()
}

func (o OSProfileOutput) ToOSProfileOutput() OSProfileOutput {
	return o
}

func (o OSProfileOutput) ToOSProfileOutputWithContext(ctx context.Context) OSProfileOutput {
	return o
}

func (o OSProfileOutput) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return o.ToOSProfilePtrOutputWithContext(context.Background())
}

func (o OSProfileOutput) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSProfile) *OSProfile {
		return &v
	}).(OSProfilePtrOutput)
}

func (o OSProfileOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfile) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

func (o OSProfileOutput) SshPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfile) *string { return v.SshPublicKey }).(pulumi.StringPtrOutput)
}

type OSProfilePtrOutput struct{ *pulumi.OutputState }

func (OSProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfile)(nil)).Elem()
}

func (o OSProfilePtrOutput) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return o
}

func (o OSProfilePtrOutput) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return o
}

func (o OSProfilePtrOutput) Elem() OSProfileOutput {
	return o.ApplyT(func(v *OSProfile) OSProfile {
		if v != nil {
			return *v
		}
		var ret OSProfile
		return ret
	}).(OSProfileOutput)
}

func (o OSProfilePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfile) *string {
		if v == nil {
			return nil
		}
		return v.ComputerName
	}).(pulumi.StringPtrOutput)
}

func (o OSProfilePtrOutput) SshPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfile) *string {
		if v == nil {
			return nil
		}
		return v.SshPublicKey
	}).(pulumi.StringPtrOutput)
}

type OSProfileResponse struct {
	ComputerName *string `pulumi:"computerName"`
	OsType       string  `pulumi:"osType"`
	SshPublicKey *string `pulumi:"sshPublicKey"`
	Version      string  `pulumi:"version"`
}

type OSProfileResponseOutput struct{ *pulumi.OutputState }

func (OSProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfileResponse)(nil)).Elem()
}

func (o OSProfileResponseOutput) ToOSProfileResponseOutput() OSProfileResponseOutput {
	return o
}

func (o OSProfileResponseOutput) ToOSProfileResponseOutputWithContext(ctx context.Context) OSProfileResponseOutput {
	return o
}

func (o OSProfileResponseOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

func (o OSProfileResponseOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v OSProfileResponse) string { return v.OsType }).(pulumi.StringOutput)
}

func (o OSProfileResponseOutput) SshPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *string { return v.SshPublicKey }).(pulumi.StringPtrOutput)
}

func (o OSProfileResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v OSProfileResponse) string { return v.Version }).(pulumi.StringOutput)
}

type OSProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (OSProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfileResponse)(nil)).Elem()
}

func (o OSProfileResponsePtrOutput) ToOSProfileResponsePtrOutput() OSProfileResponsePtrOutput {
	return o
}

func (o OSProfileResponsePtrOutput) ToOSProfileResponsePtrOutputWithContext(ctx context.Context) OSProfileResponsePtrOutput {
	return o
}

func (o OSProfileResponsePtrOutput) Elem() OSProfileResponseOutput {
	return o.ApplyT(func(v *OSProfileResponse) OSProfileResponse {
		if v != nil {
			return *v
		}
		var ret OSProfileResponse
		return ret
	}).(OSProfileResponseOutput)
}

func (o OSProfileResponsePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ComputerName
	}).(pulumi.StringPtrOutput)
}

func (o OSProfileResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o OSProfileResponsePtrOutput) SshPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.SshPublicKey
	}).(pulumi.StringPtrOutput)
}

func (o OSProfileResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

type SAPSystemID struct {
	Gid      *string `pulumi:"gid"`
	Sid      *string `pulumi:"sid"`
	Uid      *string `pulumi:"uid"`
	Username *string `pulumi:"username"`
}





type SAPSystemIDInput interface {
	pulumi.Input

	ToSAPSystemIDOutput() SAPSystemIDOutput
	ToSAPSystemIDOutputWithContext(context.Context) SAPSystemIDOutput
}

type SAPSystemIDArgs struct {
	Gid      pulumi.StringPtrInput `pulumi:"gid"`
	Sid      pulumi.StringPtrInput `pulumi:"sid"`
	Uid      pulumi.StringPtrInput `pulumi:"uid"`
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (SAPSystemIDArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SAPSystemID)(nil)).Elem()
}

func (i SAPSystemIDArgs) ToSAPSystemIDOutput() SAPSystemIDOutput {
	return i.ToSAPSystemIDOutputWithContext(context.Background())
}

func (i SAPSystemIDArgs) ToSAPSystemIDOutputWithContext(ctx context.Context) SAPSystemIDOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SAPSystemIDOutput)
}





type SAPSystemIDArrayInput interface {
	pulumi.Input

	ToSAPSystemIDArrayOutput() SAPSystemIDArrayOutput
	ToSAPSystemIDArrayOutputWithContext(context.Context) SAPSystemIDArrayOutput
}

type SAPSystemIDArray []SAPSystemIDInput

func (SAPSystemIDArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SAPSystemID)(nil)).Elem()
}

func (i SAPSystemIDArray) ToSAPSystemIDArrayOutput() SAPSystemIDArrayOutput {
	return i.ToSAPSystemIDArrayOutputWithContext(context.Background())
}

func (i SAPSystemIDArray) ToSAPSystemIDArrayOutputWithContext(ctx context.Context) SAPSystemIDArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SAPSystemIDArrayOutput)
}

type SAPSystemIDOutput struct{ *pulumi.OutputState }

func (SAPSystemIDOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SAPSystemID)(nil)).Elem()
}

func (o SAPSystemIDOutput) ToSAPSystemIDOutput() SAPSystemIDOutput {
	return o
}

func (o SAPSystemIDOutput) ToSAPSystemIDOutputWithContext(ctx context.Context) SAPSystemIDOutput {
	return o
}

func (o SAPSystemIDOutput) Gid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SAPSystemID) *string { return v.Gid }).(pulumi.StringPtrOutput)
}

func (o SAPSystemIDOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SAPSystemID) *string { return v.Sid }).(pulumi.StringPtrOutput)
}

func (o SAPSystemIDOutput) Uid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SAPSystemID) *string { return v.Uid }).(pulumi.StringPtrOutput)
}

func (o SAPSystemIDOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SAPSystemID) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type SAPSystemIDArrayOutput struct{ *pulumi.OutputState }

func (SAPSystemIDArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SAPSystemID)(nil)).Elem()
}

func (o SAPSystemIDArrayOutput) ToSAPSystemIDArrayOutput() SAPSystemIDArrayOutput {
	return o
}

func (o SAPSystemIDArrayOutput) ToSAPSystemIDArrayOutputWithContext(ctx context.Context) SAPSystemIDArrayOutput {
	return o
}

func (o SAPSystemIDArrayOutput) Index(i pulumi.IntInput) SAPSystemIDOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SAPSystemID {
		return vs[0].([]SAPSystemID)[vs[1].(int)]
	}).(SAPSystemIDOutput)
}

type SAPSystemIDResponse struct {
	Gid              *string `pulumi:"gid"`
	MemoryAllocation string  `pulumi:"memoryAllocation"`
	Sid              *string `pulumi:"sid"`
	Uid              *string `pulumi:"uid"`
	Username         *string `pulumi:"username"`
}

type SAPSystemIDResponseOutput struct{ *pulumi.OutputState }

func (SAPSystemIDResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SAPSystemIDResponse)(nil)).Elem()
}

func (o SAPSystemIDResponseOutput) ToSAPSystemIDResponseOutput() SAPSystemIDResponseOutput {
	return o
}

func (o SAPSystemIDResponseOutput) ToSAPSystemIDResponseOutputWithContext(ctx context.Context) SAPSystemIDResponseOutput {
	return o
}

func (o SAPSystemIDResponseOutput) Gid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SAPSystemIDResponse) *string { return v.Gid }).(pulumi.StringPtrOutput)
}

func (o SAPSystemIDResponseOutput) MemoryAllocation() pulumi.StringOutput {
	return o.ApplyT(func(v SAPSystemIDResponse) string { return v.MemoryAllocation }).(pulumi.StringOutput)
}

func (o SAPSystemIDResponseOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SAPSystemIDResponse) *string { return v.Sid }).(pulumi.StringPtrOutput)
}

func (o SAPSystemIDResponseOutput) Uid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SAPSystemIDResponse) *string { return v.Uid }).(pulumi.StringPtrOutput)
}

func (o SAPSystemIDResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SAPSystemIDResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type SAPSystemIDResponseArrayOutput struct{ *pulumi.OutputState }

func (SAPSystemIDResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SAPSystemIDResponse)(nil)).Elem()
}

func (o SAPSystemIDResponseArrayOutput) ToSAPSystemIDResponseArrayOutput() SAPSystemIDResponseArrayOutput {
	return o
}

func (o SAPSystemIDResponseArrayOutput) ToSAPSystemIDResponseArrayOutputWithContext(ctx context.Context) SAPSystemIDResponseArrayOutput {
	return o
}

func (o SAPSystemIDResponseArrayOutput) Index(i pulumi.IntInput) SAPSystemIDResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SAPSystemIDResponse {
		return vs[0].([]SAPSystemIDResponse)[vs[1].(int)]
	}).(SAPSystemIDResponseOutput)
}

type StorageProfile struct {
	HanaSids []SAPSystemID `pulumi:"hanaSids"`
	OsDisks  []Disk        `pulumi:"osDisks"`
}





type StorageProfileInput interface {
	pulumi.Input

	ToStorageProfileOutput() StorageProfileOutput
	ToStorageProfileOutputWithContext(context.Context) StorageProfileOutput
}

type StorageProfileArgs struct {
	HanaSids SAPSystemIDArrayInput `pulumi:"hanaSids"`
	OsDisks  DiskArrayInput        `pulumi:"osDisks"`
}

func (StorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (i StorageProfileArgs) ToStorageProfileOutput() StorageProfileOutput {
	return i.ToStorageProfileOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput)
}

func (i StorageProfileArgs) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput).ToStorageProfilePtrOutputWithContext(ctx)
}









type StorageProfilePtrInput interface {
	pulumi.Input

	ToStorageProfilePtrOutput() StorageProfilePtrOutput
	ToStorageProfilePtrOutputWithContext(context.Context) StorageProfilePtrOutput
}

type storageProfilePtrType StorageProfileArgs

func StorageProfilePtr(v *StorageProfileArgs) StorageProfilePtrInput {
	return (*storageProfilePtrType)(v)
}

func (*storageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfilePtrOutput)
}

type StorageProfileOutput struct{ *pulumi.OutputState }

func (StorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (o StorageProfileOutput) ToStorageProfileOutput() StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return o.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (o StorageProfileOutput) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageProfile) *StorageProfile {
		return &v
	}).(StorageProfilePtrOutput)
}

func (o StorageProfileOutput) HanaSids() SAPSystemIDArrayOutput {
	return o.ApplyT(func(v StorageProfile) []SAPSystemID { return v.HanaSids }).(SAPSystemIDArrayOutput)
}

func (o StorageProfileOutput) OsDisks() DiskArrayOutput {
	return o.ApplyT(func(v StorageProfile) []Disk { return v.OsDisks }).(DiskArrayOutput)
}

type StorageProfilePtrOutput struct{ *pulumi.OutputState }

func (StorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (o StorageProfilePtrOutput) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return o
}

func (o StorageProfilePtrOutput) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return o
}

func (o StorageProfilePtrOutput) Elem() StorageProfileOutput {
	return o.ApplyT(func(v *StorageProfile) StorageProfile {
		if v != nil {
			return *v
		}
		var ret StorageProfile
		return ret
	}).(StorageProfileOutput)
}

func (o StorageProfilePtrOutput) HanaSids() SAPSystemIDArrayOutput {
	return o.ApplyT(func(v *StorageProfile) []SAPSystemID {
		if v == nil {
			return nil
		}
		return v.HanaSids
	}).(SAPSystemIDArrayOutput)
}

func (o StorageProfilePtrOutput) OsDisks() DiskArrayOutput {
	return o.ApplyT(func(v *StorageProfile) []Disk {
		if v == nil {
			return nil
		}
		return v.OsDisks
	}).(DiskArrayOutput)
}

type StorageProfileResponse struct {
	HanaSids     []SAPSystemIDResponse `pulumi:"hanaSids"`
	NfsIpAddress string                `pulumi:"nfsIpAddress"`
	OsDisks      []DiskResponse        `pulumi:"osDisks"`
}

type StorageProfileResponseOutput struct{ *pulumi.OutputState }

func (StorageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutput() StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutputWithContext(ctx context.Context) StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) HanaSids() SAPSystemIDResponseArrayOutput {
	return o.ApplyT(func(v StorageProfileResponse) []SAPSystemIDResponse { return v.HanaSids }).(SAPSystemIDResponseArrayOutput)
}

func (o StorageProfileResponseOutput) NfsIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v StorageProfileResponse) string { return v.NfsIpAddress }).(pulumi.StringOutput)
}

func (o StorageProfileResponseOutput) OsDisks() DiskResponseArrayOutput {
	return o.ApplyT(func(v StorageProfileResponse) []DiskResponse { return v.OsDisks }).(DiskResponseArrayOutput)
}

type StorageProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) Elem() StorageProfileResponseOutput {
	return o.ApplyT(func(v *StorageProfileResponse) StorageProfileResponse {
		if v != nil {
			return *v
		}
		var ret StorageProfileResponse
		return ret
	}).(StorageProfileResponseOutput)
}

func (o StorageProfileResponsePtrOutput) HanaSids() SAPSystemIDResponseArrayOutput {
	return o.ApplyT(func(v *StorageProfileResponse) []SAPSystemIDResponse {
		if v == nil {
			return nil
		}
		return v.HanaSids
	}).(SAPSystemIDResponseArrayOutput)
}

func (o StorageProfileResponsePtrOutput) NfsIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.NfsIpAddress
	}).(pulumi.StringPtrOutput)
}

func (o StorageProfileResponsePtrOutput) OsDisks() DiskResponseArrayOutput {
	return o.ApplyT(func(v *StorageProfileResponse) []DiskResponse {
		if v == nil {
			return nil
		}
		return v.OsDisks
	}).(DiskResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(DiskOutput{})
	pulumi.RegisterOutputType(DiskArrayOutput{})
	pulumi.RegisterOutputType(DiskResponseOutput{})
	pulumi.RegisterOutputType(DiskResponseArrayOutput{})
	pulumi.RegisterOutputType(HardwareProfileResponseOutput{})
	pulumi.RegisterOutputType(HardwareProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(IpAddressOutput{})
	pulumi.RegisterOutputType(IpAddressArrayOutput{})
	pulumi.RegisterOutputType(IpAddressResponseOutput{})
	pulumi.RegisterOutputType(IpAddressResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkProfileOutput{})
	pulumi.RegisterOutputType(NetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(OSProfileOutput{})
	pulumi.RegisterOutputType(OSProfilePtrOutput{})
	pulumi.RegisterOutputType(OSProfileResponseOutput{})
	pulumi.RegisterOutputType(OSProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SAPSystemIDOutput{})
	pulumi.RegisterOutputType(SAPSystemIDArrayOutput{})
	pulumi.RegisterOutputType(SAPSystemIDResponseOutput{})
	pulumi.RegisterOutputType(SAPSystemIDResponseArrayOutput{})
	pulumi.RegisterOutputType(StorageProfileOutput{})
	pulumi.RegisterOutputType(StorageProfilePtrOutput{})
	pulumi.RegisterOutputType(StorageProfileResponseOutput{})
	pulumi.RegisterOutputType(StorageProfileResponsePtrOutput{})
}




package vmwarecloudsimple

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GuestOSCustomization struct {
	DnsServers []string `pulumi:"dnsServers"`
	HostName   *string  `pulumi:"hostName"`
	Password   *string  `pulumi:"password"`
	PolicyId   *string  `pulumi:"policyId"`
	Username   *string  `pulumi:"username"`
}





type GuestOSCustomizationInput interface {
	pulumi.Input

	ToGuestOSCustomizationOutput() GuestOSCustomizationOutput
	ToGuestOSCustomizationOutputWithContext(context.Context) GuestOSCustomizationOutput
}

type GuestOSCustomizationArgs struct {
	DnsServers pulumi.StringArrayInput `pulumi:"dnsServers"`
	HostName   pulumi.StringPtrInput   `pulumi:"hostName"`
	Password   pulumi.StringPtrInput   `pulumi:"password"`
	PolicyId   pulumi.StringPtrInput   `pulumi:"policyId"`
	Username   pulumi.StringPtrInput   `pulumi:"username"`
}

func (GuestOSCustomizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestOSCustomization)(nil)).Elem()
}

func (i GuestOSCustomizationArgs) ToGuestOSCustomizationOutput() GuestOSCustomizationOutput {
	return i.ToGuestOSCustomizationOutputWithContext(context.Background())
}

func (i GuestOSCustomizationArgs) ToGuestOSCustomizationOutputWithContext(ctx context.Context) GuestOSCustomizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestOSCustomizationOutput)
}

func (i GuestOSCustomizationArgs) ToGuestOSCustomizationPtrOutput() GuestOSCustomizationPtrOutput {
	return i.ToGuestOSCustomizationPtrOutputWithContext(context.Background())
}

func (i GuestOSCustomizationArgs) ToGuestOSCustomizationPtrOutputWithContext(ctx context.Context) GuestOSCustomizationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestOSCustomizationOutput).ToGuestOSCustomizationPtrOutputWithContext(ctx)
}









type GuestOSCustomizationPtrInput interface {
	pulumi.Input

	ToGuestOSCustomizationPtrOutput() GuestOSCustomizationPtrOutput
	ToGuestOSCustomizationPtrOutputWithContext(context.Context) GuestOSCustomizationPtrOutput
}

type guestOSCustomizationPtrType GuestOSCustomizationArgs

func GuestOSCustomizationPtr(v *GuestOSCustomizationArgs) GuestOSCustomizationPtrInput {
	return (*guestOSCustomizationPtrType)(v)
}

func (*guestOSCustomizationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestOSCustomization)(nil)).Elem()
}

func (i *guestOSCustomizationPtrType) ToGuestOSCustomizationPtrOutput() GuestOSCustomizationPtrOutput {
	return i.ToGuestOSCustomizationPtrOutputWithContext(context.Background())
}

func (i *guestOSCustomizationPtrType) ToGuestOSCustomizationPtrOutputWithContext(ctx context.Context) GuestOSCustomizationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestOSCustomizationPtrOutput)
}

type GuestOSCustomizationOutput struct{ *pulumi.OutputState }

func (GuestOSCustomizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestOSCustomization)(nil)).Elem()
}

func (o GuestOSCustomizationOutput) ToGuestOSCustomizationOutput() GuestOSCustomizationOutput {
	return o
}

func (o GuestOSCustomizationOutput) ToGuestOSCustomizationOutputWithContext(ctx context.Context) GuestOSCustomizationOutput {
	return o
}

func (o GuestOSCustomizationOutput) ToGuestOSCustomizationPtrOutput() GuestOSCustomizationPtrOutput {
	return o.ToGuestOSCustomizationPtrOutputWithContext(context.Background())
}

func (o GuestOSCustomizationOutput) ToGuestOSCustomizationPtrOutputWithContext(ctx context.Context) GuestOSCustomizationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GuestOSCustomization) *GuestOSCustomization {
		return &v
	}).(GuestOSCustomizationPtrOutput)
}

func (o GuestOSCustomizationOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GuestOSCustomization) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

func (o GuestOSCustomizationOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestOSCustomization) *string { return v.HostName }).(pulumi.StringPtrOutput)
}

func (o GuestOSCustomizationOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestOSCustomization) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o GuestOSCustomizationOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestOSCustomization) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o GuestOSCustomizationOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestOSCustomization) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type GuestOSCustomizationPtrOutput struct{ *pulumi.OutputState }

func (GuestOSCustomizationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestOSCustomization)(nil)).Elem()
}

func (o GuestOSCustomizationPtrOutput) ToGuestOSCustomizationPtrOutput() GuestOSCustomizationPtrOutput {
	return o
}

func (o GuestOSCustomizationPtrOutput) ToGuestOSCustomizationPtrOutputWithContext(ctx context.Context) GuestOSCustomizationPtrOutput {
	return o
}

func (o GuestOSCustomizationPtrOutput) Elem() GuestOSCustomizationOutput {
	return o.ApplyT(func(v *GuestOSCustomization) GuestOSCustomization {
		if v != nil {
			return *v
		}
		var ret GuestOSCustomization
		return ret
	}).(GuestOSCustomizationOutput)
}

func (o GuestOSCustomizationPtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GuestOSCustomization) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

func (o GuestOSCustomizationPtrOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestOSCustomization) *string {
		if v == nil {
			return nil
		}
		return v.HostName
	}).(pulumi.StringPtrOutput)
}

func (o GuestOSCustomizationPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestOSCustomization) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o GuestOSCustomizationPtrOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestOSCustomization) *string {
		if v == nil {
			return nil
		}
		return v.PolicyId
	}).(pulumi.StringPtrOutput)
}

func (o GuestOSCustomizationPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestOSCustomization) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type GuestOSCustomizationResponse struct {
	DnsServers []string `pulumi:"dnsServers"`
	HostName   *string  `pulumi:"hostName"`
	Password   *string  `pulumi:"password"`
	PolicyId   *string  `pulumi:"policyId"`
	Username   *string  `pulumi:"username"`
}





type GuestOSCustomizationResponseInput interface {
	pulumi.Input

	ToGuestOSCustomizationResponseOutput() GuestOSCustomizationResponseOutput
	ToGuestOSCustomizationResponseOutputWithContext(context.Context) GuestOSCustomizationResponseOutput
}

type GuestOSCustomizationResponseArgs struct {
	DnsServers pulumi.StringArrayInput `pulumi:"dnsServers"`
	HostName   pulumi.StringPtrInput   `pulumi:"hostName"`
	Password   pulumi.StringPtrInput   `pulumi:"password"`
	PolicyId   pulumi.StringPtrInput   `pulumi:"policyId"`
	Username   pulumi.StringPtrInput   `pulumi:"username"`
}

func (GuestOSCustomizationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestOSCustomizationResponse)(nil)).Elem()
}

func (i GuestOSCustomizationResponseArgs) ToGuestOSCustomizationResponseOutput() GuestOSCustomizationResponseOutput {
	return i.ToGuestOSCustomizationResponseOutputWithContext(context.Background())
}

func (i GuestOSCustomizationResponseArgs) ToGuestOSCustomizationResponseOutputWithContext(ctx context.Context) GuestOSCustomizationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestOSCustomizationResponseOutput)
}

func (i GuestOSCustomizationResponseArgs) ToGuestOSCustomizationResponsePtrOutput() GuestOSCustomizationResponsePtrOutput {
	return i.ToGuestOSCustomizationResponsePtrOutputWithContext(context.Background())
}

func (i GuestOSCustomizationResponseArgs) ToGuestOSCustomizationResponsePtrOutputWithContext(ctx context.Context) GuestOSCustomizationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestOSCustomizationResponseOutput).ToGuestOSCustomizationResponsePtrOutputWithContext(ctx)
}









type GuestOSCustomizationResponsePtrInput interface {
	pulumi.Input

	ToGuestOSCustomizationResponsePtrOutput() GuestOSCustomizationResponsePtrOutput
	ToGuestOSCustomizationResponsePtrOutputWithContext(context.Context) GuestOSCustomizationResponsePtrOutput
}

type guestOSCustomizationResponsePtrType GuestOSCustomizationResponseArgs

func GuestOSCustomizationResponsePtr(v *GuestOSCustomizationResponseArgs) GuestOSCustomizationResponsePtrInput {
	return (*guestOSCustomizationResponsePtrType)(v)
}

func (*guestOSCustomizationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestOSCustomizationResponse)(nil)).Elem()
}

func (i *guestOSCustomizationResponsePtrType) ToGuestOSCustomizationResponsePtrOutput() GuestOSCustomizationResponsePtrOutput {
	return i.ToGuestOSCustomizationResponsePtrOutputWithContext(context.Background())
}

func (i *guestOSCustomizationResponsePtrType) ToGuestOSCustomizationResponsePtrOutputWithContext(ctx context.Context) GuestOSCustomizationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestOSCustomizationResponsePtrOutput)
}

type GuestOSCustomizationResponseOutput struct{ *pulumi.OutputState }

func (GuestOSCustomizationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestOSCustomizationResponse)(nil)).Elem()
}

func (o GuestOSCustomizationResponseOutput) ToGuestOSCustomizationResponseOutput() GuestOSCustomizationResponseOutput {
	return o
}

func (o GuestOSCustomizationResponseOutput) ToGuestOSCustomizationResponseOutputWithContext(ctx context.Context) GuestOSCustomizationResponseOutput {
	return o
}

func (o GuestOSCustomizationResponseOutput) ToGuestOSCustomizationResponsePtrOutput() GuestOSCustomizationResponsePtrOutput {
	return o.ToGuestOSCustomizationResponsePtrOutputWithContext(context.Background())
}

func (o GuestOSCustomizationResponseOutput) ToGuestOSCustomizationResponsePtrOutputWithContext(ctx context.Context) GuestOSCustomizationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GuestOSCustomizationResponse) *GuestOSCustomizationResponse {
		return &v
	}).(GuestOSCustomizationResponsePtrOutput)
}

func (o GuestOSCustomizationResponseOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GuestOSCustomizationResponse) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

func (o GuestOSCustomizationResponseOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestOSCustomizationResponse) *string { return v.HostName }).(pulumi.StringPtrOutput)
}

func (o GuestOSCustomizationResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestOSCustomizationResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o GuestOSCustomizationResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestOSCustomizationResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o GuestOSCustomizationResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestOSCustomizationResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type GuestOSCustomizationResponsePtrOutput struct{ *pulumi.OutputState }

func (GuestOSCustomizationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestOSCustomizationResponse)(nil)).Elem()
}

func (o GuestOSCustomizationResponsePtrOutput) ToGuestOSCustomizationResponsePtrOutput() GuestOSCustomizationResponsePtrOutput {
	return o
}

func (o GuestOSCustomizationResponsePtrOutput) ToGuestOSCustomizationResponsePtrOutputWithContext(ctx context.Context) GuestOSCustomizationResponsePtrOutput {
	return o
}

func (o GuestOSCustomizationResponsePtrOutput) Elem() GuestOSCustomizationResponseOutput {
	return o.ApplyT(func(v *GuestOSCustomizationResponse) GuestOSCustomizationResponse {
		if v != nil {
			return *v
		}
		var ret GuestOSCustomizationResponse
		return ret
	}).(GuestOSCustomizationResponseOutput)
}

func (o GuestOSCustomizationResponsePtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GuestOSCustomizationResponse) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

func (o GuestOSCustomizationResponsePtrOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestOSCustomizationResponse) *string {
		if v == nil {
			return nil
		}
		return v.HostName
	}).(pulumi.StringPtrOutput)
}

func (o GuestOSCustomizationResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestOSCustomizationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o GuestOSCustomizationResponsePtrOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestOSCustomizationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PolicyId
	}).(pulumi.StringPtrOutput)
}

func (o GuestOSCustomizationResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestOSCustomizationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type GuestOSNICCustomization struct {
	Allocation          *string  `pulumi:"allocation"`
	DnsServers          []string `pulumi:"dnsServers"`
	Gateway             []string `pulumi:"gateway"`
	IpAddress           *string  `pulumi:"ipAddress"`
	Mask                *string  `pulumi:"mask"`
	PrimaryWinsServer   *string  `pulumi:"primaryWinsServer"`
	SecondaryWinsServer *string  `pulumi:"secondaryWinsServer"`
}





type GuestOSNICCustomizationInput interface {
	pulumi.Input

	ToGuestOSNICCustomizationOutput() GuestOSNICCustomizationOutput
	ToGuestOSNICCustomizationOutputWithContext(context.Context) GuestOSNICCustomizationOutput
}

type GuestOSNICCustomizationArgs struct {
	Allocation          pulumi.StringPtrInput   `pulumi:"allocation"`
	DnsServers          pulumi.StringArrayInput `pulumi:"dnsServers"`
	Gateway             pulumi.StringArrayInput `pulumi:"gateway"`
	IpAddress           pulumi.StringPtrInput   `pulumi:"ipAddress"`
	Mask                pulumi.StringPtrInput   `pulumi:"mask"`
	PrimaryWinsServer   pulumi.StringPtrInput   `pulumi:"primaryWinsServer"`
	SecondaryWinsServer pulumi.StringPtrInput   `pulumi:"secondaryWinsServer"`
}

func (GuestOSNICCustomizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestOSNICCustomization)(nil)).Elem()
}

func (i GuestOSNICCustomizationArgs) ToGuestOSNICCustomizationOutput() GuestOSNICCustomizationOutput {
	return i.ToGuestOSNICCustomizationOutputWithContext(context.Background())
}

func (i GuestOSNICCustomizationArgs) ToGuestOSNICCustomizationOutputWithContext(ctx context.Context) GuestOSNICCustomizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestOSNICCustomizationOutput)
}

func (i GuestOSNICCustomizationArgs) ToGuestOSNICCustomizationPtrOutput() GuestOSNICCustomizationPtrOutput {
	return i.ToGuestOSNICCustomizationPtrOutputWithContext(context.Background())
}

func (i GuestOSNICCustomizationArgs) ToGuestOSNICCustomizationPtrOutputWithContext(ctx context.Context) GuestOSNICCustomizationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestOSNICCustomizationOutput).ToGuestOSNICCustomizationPtrOutputWithContext(ctx)
}









type GuestOSNICCustomizationPtrInput interface {
	pulumi.Input

	ToGuestOSNICCustomizationPtrOutput() GuestOSNICCustomizationPtrOutput
	ToGuestOSNICCustomizationPtrOutputWithContext(context.Context) GuestOSNICCustomizationPtrOutput
}

type guestOSNICCustomizationPtrType GuestOSNICCustomizationArgs

func GuestOSNICCustomizationPtr(v *GuestOSNICCustomizationArgs) GuestOSNICCustomizationPtrInput {
	return (*guestOSNICCustomizationPtrType)(v)
}

func (*guestOSNICCustomizationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestOSNICCustomization)(nil)).Elem()
}

func (i *guestOSNICCustomizationPtrType) ToGuestOSNICCustomizationPtrOutput() GuestOSNICCustomizationPtrOutput {
	return i.ToGuestOSNICCustomizationPtrOutputWithContext(context.Background())
}

func (i *guestOSNICCustomizationPtrType) ToGuestOSNICCustomizationPtrOutputWithContext(ctx context.Context) GuestOSNICCustomizationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestOSNICCustomizationPtrOutput)
}

type GuestOSNICCustomizationOutput struct{ *pulumi.OutputState }

func (GuestOSNICCustomizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestOSNICCustomization)(nil)).Elem()
}

func (o GuestOSNICCustomizationOutput) ToGuestOSNICCustomizationOutput() GuestOSNICCustomizationOutput {
	return o
}

func (o GuestOSNICCustomizationOutput) ToGuestOSNICCustomizationOutputWithContext(ctx context.Context) GuestOSNICCustomizationOutput {
	return o
}

func (o GuestOSNICCustomizationOutput) ToGuestOSNICCustomizationPtrOutput() GuestOSNICCustomizationPtrOutput {
	return o.ToGuestOSNICCustomizationPtrOutputWithContext(context.Background())
}

func (o GuestOSNICCustomizationOutput) ToGuestOSNICCustomizationPtrOutputWithContext(ctx context.Context) GuestOSNICCustomizationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GuestOSNICCustomization) *GuestOSNICCustomization {
		return &v
	}).(GuestOSNICCustomizationPtrOutput)
}

func (o GuestOSNICCustomizationOutput) Allocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestOSNICCustomization) *string { return v.Allocation }).(pulumi.StringPtrOutput)
}

func (o GuestOSNICCustomizationOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GuestOSNICCustomization) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

func (o GuestOSNICCustomizationOutput) Gateway() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GuestOSNICCustomization) []string { return v.Gateway }).(pulumi.StringArrayOutput)
}

func (o GuestOSNICCustomizationOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestOSNICCustomization) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o GuestOSNICCustomizationOutput) Mask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestOSNICCustomization) *string { return v.Mask }).(pulumi.StringPtrOutput)
}

func (o GuestOSNICCustomizationOutput) PrimaryWinsServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestOSNICCustomization) *string { return v.PrimaryWinsServer }).(pulumi.StringPtrOutput)
}

func (o GuestOSNICCustomizationOutput) SecondaryWinsServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestOSNICCustomization) *string { return v.SecondaryWinsServer }).(pulumi.StringPtrOutput)
}

type GuestOSNICCustomizationPtrOutput struct{ *pulumi.OutputState }

func (GuestOSNICCustomizationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestOSNICCustomization)(nil)).Elem()
}

func (o GuestOSNICCustomizationPtrOutput) ToGuestOSNICCustomizationPtrOutput() GuestOSNICCustomizationPtrOutput {
	return o
}

func (o GuestOSNICCustomizationPtrOutput) ToGuestOSNICCustomizationPtrOutputWithContext(ctx context.Context) GuestOSNICCustomizationPtrOutput {
	return o
}

func (o GuestOSNICCustomizationPtrOutput) Elem() GuestOSNICCustomizationOutput {
	return o.ApplyT(func(v *GuestOSNICCustomization) GuestOSNICCustomization {
		if v != nil {
			return *v
		}
		var ret GuestOSNICCustomization
		return ret
	}).(GuestOSNICCustomizationOutput)
}

func (o GuestOSNICCustomizationPtrOutput) Allocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestOSNICCustomization) *string {
		if v == nil {
			return nil
		}
		return v.Allocation
	}).(pulumi.StringPtrOutput)
}

func (o GuestOSNICCustomizationPtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GuestOSNICCustomization) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

func (o GuestOSNICCustomizationPtrOutput) Gateway() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GuestOSNICCustomization) []string {
		if v == nil {
			return nil
		}
		return v.Gateway
	}).(pulumi.StringArrayOutput)
}

func (o GuestOSNICCustomizationPtrOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestOSNICCustomization) *string {
		if v == nil {
			return nil
		}
		return v.IpAddress
	}).(pulumi.StringPtrOutput)
}

func (o GuestOSNICCustomizationPtrOutput) Mask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestOSNICCustomization) *string {
		if v == nil {
			return nil
		}
		return v.Mask
	}).(pulumi.StringPtrOutput)
}

func (o GuestOSNICCustomizationPtrOutput) PrimaryWinsServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestOSNICCustomization) *string {
		if v == nil {
			return nil
		}
		return v.PrimaryWinsServer
	}).(pulumi.StringPtrOutput)
}

func (o GuestOSNICCustomizationPtrOutput) SecondaryWinsServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestOSNICCustomization) *string {
		if v == nil {
			return nil
		}
		return v.SecondaryWinsServer
	}).(pulumi.StringPtrOutput)
}

type GuestOSNICCustomizationResponse struct {
	Allocation          *string  `pulumi:"allocation"`
	DnsServers          []string `pulumi:"dnsServers"`
	Gateway             []string `pulumi:"gateway"`
	IpAddress           *string  `pulumi:"ipAddress"`
	Mask                *string  `pulumi:"mask"`
	PrimaryWinsServer   *string  `pulumi:"primaryWinsServer"`
	SecondaryWinsServer *string  `pulumi:"secondaryWinsServer"`
}





type GuestOSNICCustomizationResponseInput interface {
	pulumi.Input

	ToGuestOSNICCustomizationResponseOutput() GuestOSNICCustomizationResponseOutput
	ToGuestOSNICCustomizationResponseOutputWithContext(context.Context) GuestOSNICCustomizationResponseOutput
}

type GuestOSNICCustomizationResponseArgs struct {
	Allocation          pulumi.StringPtrInput   `pulumi:"allocation"`
	DnsServers          pulumi.StringArrayInput `pulumi:"dnsServers"`
	Gateway             pulumi.StringArrayInput `pulumi:"gateway"`
	IpAddress           pulumi.StringPtrInput   `pulumi:"ipAddress"`
	Mask                pulumi.StringPtrInput   `pulumi:"mask"`
	PrimaryWinsServer   pulumi.StringPtrInput   `pulumi:"primaryWinsServer"`
	SecondaryWinsServer pulumi.StringPtrInput   `pulumi:"secondaryWinsServer"`
}

func (GuestOSNICCustomizationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestOSNICCustomizationResponse)(nil)).Elem()
}

func (i GuestOSNICCustomizationResponseArgs) ToGuestOSNICCustomizationResponseOutput() GuestOSNICCustomizationResponseOutput {
	return i.ToGuestOSNICCustomizationResponseOutputWithContext(context.Background())
}

func (i GuestOSNICCustomizationResponseArgs) ToGuestOSNICCustomizationResponseOutputWithContext(ctx context.Context) GuestOSNICCustomizationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestOSNICCustomizationResponseOutput)
}

func (i GuestOSNICCustomizationResponseArgs) ToGuestOSNICCustomizationResponsePtrOutput() GuestOSNICCustomizationResponsePtrOutput {
	return i.ToGuestOSNICCustomizationResponsePtrOutputWithContext(context.Background())
}

func (i GuestOSNICCustomizationResponseArgs) ToGuestOSNICCustomizationResponsePtrOutputWithContext(ctx context.Context) GuestOSNICCustomizationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestOSNICCustomizationResponseOutput).ToGuestOSNICCustomizationResponsePtrOutputWithContext(ctx)
}









type GuestOSNICCustomizationResponsePtrInput interface {
	pulumi.Input

	ToGuestOSNICCustomizationResponsePtrOutput() GuestOSNICCustomizationResponsePtrOutput
	ToGuestOSNICCustomizationResponsePtrOutputWithContext(context.Context) GuestOSNICCustomizationResponsePtrOutput
}

type guestOSNICCustomizationResponsePtrType GuestOSNICCustomizationResponseArgs

func GuestOSNICCustomizationResponsePtr(v *GuestOSNICCustomizationResponseArgs) GuestOSNICCustomizationResponsePtrInput {
	return (*guestOSNICCustomizationResponsePtrType)(v)
}

func (*guestOSNICCustomizationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestOSNICCustomizationResponse)(nil)).Elem()
}

func (i *guestOSNICCustomizationResponsePtrType) ToGuestOSNICCustomizationResponsePtrOutput() GuestOSNICCustomizationResponsePtrOutput {
	return i.ToGuestOSNICCustomizationResponsePtrOutputWithContext(context.Background())
}

func (i *guestOSNICCustomizationResponsePtrType) ToGuestOSNICCustomizationResponsePtrOutputWithContext(ctx context.Context) GuestOSNICCustomizationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestOSNICCustomizationResponsePtrOutput)
}

type GuestOSNICCustomizationResponseOutput struct{ *pulumi.OutputState }

func (GuestOSNICCustomizationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestOSNICCustomizationResponse)(nil)).Elem()
}

func (o GuestOSNICCustomizationResponseOutput) ToGuestOSNICCustomizationResponseOutput() GuestOSNICCustomizationResponseOutput {
	return o
}

func (o GuestOSNICCustomizationResponseOutput) ToGuestOSNICCustomizationResponseOutputWithContext(ctx context.Context) GuestOSNICCustomizationResponseOutput {
	return o
}

func (o GuestOSNICCustomizationResponseOutput) ToGuestOSNICCustomizationResponsePtrOutput() GuestOSNICCustomizationResponsePtrOutput {
	return o.ToGuestOSNICCustomizationResponsePtrOutputWithContext(context.Background())
}

func (o GuestOSNICCustomizationResponseOutput) ToGuestOSNICCustomizationResponsePtrOutputWithContext(ctx context.Context) GuestOSNICCustomizationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GuestOSNICCustomizationResponse) *GuestOSNICCustomizationResponse {
		return &v
	}).(GuestOSNICCustomizationResponsePtrOutput)
}

func (o GuestOSNICCustomizationResponseOutput) Allocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestOSNICCustomizationResponse) *string { return v.Allocation }).(pulumi.StringPtrOutput)
}

func (o GuestOSNICCustomizationResponseOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GuestOSNICCustomizationResponse) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

func (o GuestOSNICCustomizationResponseOutput) Gateway() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GuestOSNICCustomizationResponse) []string { return v.Gateway }).(pulumi.StringArrayOutput)
}

func (o GuestOSNICCustomizationResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestOSNICCustomizationResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o GuestOSNICCustomizationResponseOutput) Mask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestOSNICCustomizationResponse) *string { return v.Mask }).(pulumi.StringPtrOutput)
}

func (o GuestOSNICCustomizationResponseOutput) PrimaryWinsServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestOSNICCustomizationResponse) *string { return v.PrimaryWinsServer }).(pulumi.StringPtrOutput)
}

func (o GuestOSNICCustomizationResponseOutput) SecondaryWinsServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestOSNICCustomizationResponse) *string { return v.SecondaryWinsServer }).(pulumi.StringPtrOutput)
}

type GuestOSNICCustomizationResponsePtrOutput struct{ *pulumi.OutputState }

func (GuestOSNICCustomizationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestOSNICCustomizationResponse)(nil)).Elem()
}

func (o GuestOSNICCustomizationResponsePtrOutput) ToGuestOSNICCustomizationResponsePtrOutput() GuestOSNICCustomizationResponsePtrOutput {
	return o
}

func (o GuestOSNICCustomizationResponsePtrOutput) ToGuestOSNICCustomizationResponsePtrOutputWithContext(ctx context.Context) GuestOSNICCustomizationResponsePtrOutput {
	return o
}

func (o GuestOSNICCustomizationResponsePtrOutput) Elem() GuestOSNICCustomizationResponseOutput {
	return o.ApplyT(func(v *GuestOSNICCustomizationResponse) GuestOSNICCustomizationResponse {
		if v != nil {
			return *v
		}
		var ret GuestOSNICCustomizationResponse
		return ret
	}).(GuestOSNICCustomizationResponseOutput)
}

func (o GuestOSNICCustomizationResponsePtrOutput) Allocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestOSNICCustomizationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Allocation
	}).(pulumi.StringPtrOutput)
}

func (o GuestOSNICCustomizationResponsePtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GuestOSNICCustomizationResponse) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

func (o GuestOSNICCustomizationResponsePtrOutput) Gateway() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GuestOSNICCustomizationResponse) []string {
		if v == nil {
			return nil
		}
		return v.Gateway
	}).(pulumi.StringArrayOutput)
}

func (o GuestOSNICCustomizationResponsePtrOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestOSNICCustomizationResponse) *string {
		if v == nil {
			return nil
		}
		return v.IpAddress
	}).(pulumi.StringPtrOutput)
}

func (o GuestOSNICCustomizationResponsePtrOutput) Mask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestOSNICCustomizationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mask
	}).(pulumi.StringPtrOutput)
}

func (o GuestOSNICCustomizationResponsePtrOutput) PrimaryWinsServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestOSNICCustomizationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrimaryWinsServer
	}).(pulumi.StringPtrOutput)
}

func (o GuestOSNICCustomizationResponsePtrOutput) SecondaryWinsServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestOSNICCustomizationResponse) *string {
		if v == nil {
			return nil
		}
		return v.SecondaryWinsServer
	}).(pulumi.StringPtrOutput)
}

type ResourcePool struct {
	Id string `pulumi:"id"`
}





type ResourcePoolInput interface {
	pulumi.Input

	ToResourcePoolOutput() ResourcePoolOutput
	ToResourcePoolOutputWithContext(context.Context) ResourcePoolOutput
}

type ResourcePoolArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (ResourcePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourcePool)(nil)).Elem()
}

func (i ResourcePoolArgs) ToResourcePoolOutput() ResourcePoolOutput {
	return i.ToResourcePoolOutputWithContext(context.Background())
}

func (i ResourcePoolArgs) ToResourcePoolOutputWithContext(ctx context.Context) ResourcePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePoolOutput)
}

func (i ResourcePoolArgs) ToResourcePoolPtrOutput() ResourcePoolPtrOutput {
	return i.ToResourcePoolPtrOutputWithContext(context.Background())
}

func (i ResourcePoolArgs) ToResourcePoolPtrOutputWithContext(ctx context.Context) ResourcePoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePoolOutput).ToResourcePoolPtrOutputWithContext(ctx)
}









type ResourcePoolPtrInput interface {
	pulumi.Input

	ToResourcePoolPtrOutput() ResourcePoolPtrOutput
	ToResourcePoolPtrOutputWithContext(context.Context) ResourcePoolPtrOutput
}

type resourcePoolPtrType ResourcePoolArgs

func ResourcePoolPtr(v *ResourcePoolArgs) ResourcePoolPtrInput {
	return (*resourcePoolPtrType)(v)
}

func (*resourcePoolPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePool)(nil)).Elem()
}

func (i *resourcePoolPtrType) ToResourcePoolPtrOutput() ResourcePoolPtrOutput {
	return i.ToResourcePoolPtrOutputWithContext(context.Background())
}

func (i *resourcePoolPtrType) ToResourcePoolPtrOutputWithContext(ctx context.Context) ResourcePoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePoolPtrOutput)
}

type ResourcePoolOutput struct{ *pulumi.OutputState }

func (ResourcePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourcePool)(nil)).Elem()
}

func (o ResourcePoolOutput) ToResourcePoolOutput() ResourcePoolOutput {
	return o
}

func (o ResourcePoolOutput) ToResourcePoolOutputWithContext(ctx context.Context) ResourcePoolOutput {
	return o
}

func (o ResourcePoolOutput) ToResourcePoolPtrOutput() ResourcePoolPtrOutput {
	return o.ToResourcePoolPtrOutputWithContext(context.Background())
}

func (o ResourcePoolOutput) ToResourcePoolPtrOutputWithContext(ctx context.Context) ResourcePoolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourcePool) *ResourcePool {
		return &v
	}).(ResourcePoolPtrOutput)
}

func (o ResourcePoolOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ResourcePool) string { return v.Id }).(pulumi.StringOutput)
}

type ResourcePoolPtrOutput struct{ *pulumi.OutputState }

func (ResourcePoolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePool)(nil)).Elem()
}

func (o ResourcePoolPtrOutput) ToResourcePoolPtrOutput() ResourcePoolPtrOutput {
	return o
}

func (o ResourcePoolPtrOutput) ToResourcePoolPtrOutputWithContext(ctx context.Context) ResourcePoolPtrOutput {
	return o
}

func (o ResourcePoolPtrOutput) Elem() ResourcePoolOutput {
	return o.ApplyT(func(v *ResourcePool) ResourcePool {
		if v != nil {
			return *v
		}
		var ret ResourcePool
		return ret
	}).(ResourcePoolOutput)
}

func (o ResourcePoolPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePool) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type ResourcePoolResponse struct {
	FullName       string `pulumi:"fullName"`
	Id             string `pulumi:"id"`
	Location       string `pulumi:"location"`
	Name           string `pulumi:"name"`
	PrivateCloudId string `pulumi:"privateCloudId"`
	Type           string `pulumi:"type"`
}





type ResourcePoolResponseInput interface {
	pulumi.Input

	ToResourcePoolResponseOutput() ResourcePoolResponseOutput
	ToResourcePoolResponseOutputWithContext(context.Context) ResourcePoolResponseOutput
}

type ResourcePoolResponseArgs struct {
	FullName       pulumi.StringInput `pulumi:"fullName"`
	Id             pulumi.StringInput `pulumi:"id"`
	Location       pulumi.StringInput `pulumi:"location"`
	Name           pulumi.StringInput `pulumi:"name"`
	PrivateCloudId pulumi.StringInput `pulumi:"privateCloudId"`
	Type           pulumi.StringInput `pulumi:"type"`
}

func (ResourcePoolResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourcePoolResponse)(nil)).Elem()
}

func (i ResourcePoolResponseArgs) ToResourcePoolResponseOutput() ResourcePoolResponseOutput {
	return i.ToResourcePoolResponseOutputWithContext(context.Background())
}

func (i ResourcePoolResponseArgs) ToResourcePoolResponseOutputWithContext(ctx context.Context) ResourcePoolResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePoolResponseOutput)
}

func (i ResourcePoolResponseArgs) ToResourcePoolResponsePtrOutput() ResourcePoolResponsePtrOutput {
	return i.ToResourcePoolResponsePtrOutputWithContext(context.Background())
}

func (i ResourcePoolResponseArgs) ToResourcePoolResponsePtrOutputWithContext(ctx context.Context) ResourcePoolResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePoolResponseOutput).ToResourcePoolResponsePtrOutputWithContext(ctx)
}









type ResourcePoolResponsePtrInput interface {
	pulumi.Input

	ToResourcePoolResponsePtrOutput() ResourcePoolResponsePtrOutput
	ToResourcePoolResponsePtrOutputWithContext(context.Context) ResourcePoolResponsePtrOutput
}

type resourcePoolResponsePtrType ResourcePoolResponseArgs

func ResourcePoolResponsePtr(v *ResourcePoolResponseArgs) ResourcePoolResponsePtrInput {
	return (*resourcePoolResponsePtrType)(v)
}

func (*resourcePoolResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePoolResponse)(nil)).Elem()
}

func (i *resourcePoolResponsePtrType) ToResourcePoolResponsePtrOutput() ResourcePoolResponsePtrOutput {
	return i.ToResourcePoolResponsePtrOutputWithContext(context.Background())
}

func (i *resourcePoolResponsePtrType) ToResourcePoolResponsePtrOutputWithContext(ctx context.Context) ResourcePoolResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePoolResponsePtrOutput)
}

type ResourcePoolResponseOutput struct{ *pulumi.OutputState }

func (ResourcePoolResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourcePoolResponse)(nil)).Elem()
}

func (o ResourcePoolResponseOutput) ToResourcePoolResponseOutput() ResourcePoolResponseOutput {
	return o
}

func (o ResourcePoolResponseOutput) ToResourcePoolResponseOutputWithContext(ctx context.Context) ResourcePoolResponseOutput {
	return o
}

func (o ResourcePoolResponseOutput) ToResourcePoolResponsePtrOutput() ResourcePoolResponsePtrOutput {
	return o.ToResourcePoolResponsePtrOutputWithContext(context.Background())
}

func (o ResourcePoolResponseOutput) ToResourcePoolResponsePtrOutputWithContext(ctx context.Context) ResourcePoolResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourcePoolResponse) *ResourcePoolResponse {
		return &v
	}).(ResourcePoolResponsePtrOutput)
}

func (o ResourcePoolResponseOutput) FullName() pulumi.StringOutput {
	return o.ApplyT(func(v ResourcePoolResponse) string { return v.FullName }).(pulumi.StringOutput)
}

func (o ResourcePoolResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ResourcePoolResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ResourcePoolResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ResourcePoolResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o ResourcePoolResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourcePoolResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourcePoolResponseOutput) PrivateCloudId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourcePoolResponse) string { return v.PrivateCloudId }).(pulumi.StringOutput)
}

func (o ResourcePoolResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ResourcePoolResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ResourcePoolResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourcePoolResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePoolResponse)(nil)).Elem()
}

func (o ResourcePoolResponsePtrOutput) ToResourcePoolResponsePtrOutput() ResourcePoolResponsePtrOutput {
	return o
}

func (o ResourcePoolResponsePtrOutput) ToResourcePoolResponsePtrOutputWithContext(ctx context.Context) ResourcePoolResponsePtrOutput {
	return o
}

func (o ResourcePoolResponsePtrOutput) Elem() ResourcePoolResponseOutput {
	return o.ApplyT(func(v *ResourcePoolResponse) ResourcePoolResponse {
		if v != nil {
			return *v
		}
		var ret ResourcePoolResponse
		return ret
	}).(ResourcePoolResponseOutput)
}

func (o ResourcePoolResponsePtrOutput) FullName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePoolResponse) *string {
		if v == nil {
			return nil
		}
		return &v.FullName
	}).(pulumi.StringPtrOutput)
}

func (o ResourcePoolResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePoolResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ResourcePoolResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePoolResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o ResourcePoolResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePoolResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourcePoolResponsePtrOutput) PrivateCloudId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePoolResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrivateCloudId
	}).(pulumi.StringPtrOutput)
}

func (o ResourcePoolResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePoolResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Capacity    *string `pulumi:"capacity"`
	Description *string `pulumi:"description"`
	Family      *string `pulumi:"family"`
	Name        string  `pulumi:"name"`
	Tier        *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity    pulumi.StringPtrInput `pulumi:"capacity"`
	Description pulumi.StringPtrInput `pulumi:"description"`
	Family      pulumi.StringPtrInput `pulumi:"family"`
	Name        pulumi.StringInput    `pulumi:"name"`
	Tier        pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Capacity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Capacity }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Capacity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity    *string `pulumi:"capacity"`
	Description *string `pulumi:"description"`
	Family      *string `pulumi:"family"`
	Name        string  `pulumi:"name"`
	Tier        *string `pulumi:"tier"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Capacity    pulumi.StringPtrInput `pulumi:"capacity"`
	Description pulumi.StringPtrInput `pulumi:"description"`
	Family      pulumi.StringPtrInput `pulumi:"family"`
	Name        pulumi.StringInput    `pulumi:"name"`
	Tier        pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Capacity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Capacity }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Capacity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type VirtualDisk struct {
	ControllerId     string               `pulumi:"controllerId"`
	IndependenceMode DiskIndependenceMode `pulumi:"independenceMode"`
	TotalSize        int                  `pulumi:"totalSize"`
	VirtualDiskId    *string              `pulumi:"virtualDiskId"`
}





type VirtualDiskInput interface {
	pulumi.Input

	ToVirtualDiskOutput() VirtualDiskOutput
	ToVirtualDiskOutputWithContext(context.Context) VirtualDiskOutput
}

type VirtualDiskArgs struct {
	ControllerId     pulumi.StringInput        `pulumi:"controllerId"`
	IndependenceMode DiskIndependenceModeInput `pulumi:"independenceMode"`
	TotalSize        pulumi.IntInput           `pulumi:"totalSize"`
	VirtualDiskId    pulumi.StringPtrInput     `pulumi:"virtualDiskId"`
}

func (VirtualDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDisk)(nil)).Elem()
}

func (i VirtualDiskArgs) ToVirtualDiskOutput() VirtualDiskOutput {
	return i.ToVirtualDiskOutputWithContext(context.Background())
}

func (i VirtualDiskArgs) ToVirtualDiskOutputWithContext(ctx context.Context) VirtualDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDiskOutput)
}





type VirtualDiskArrayInput interface {
	pulumi.Input

	ToVirtualDiskArrayOutput() VirtualDiskArrayOutput
	ToVirtualDiskArrayOutputWithContext(context.Context) VirtualDiskArrayOutput
}

type VirtualDiskArray []VirtualDiskInput

func (VirtualDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDisk)(nil)).Elem()
}

func (i VirtualDiskArray) ToVirtualDiskArrayOutput() VirtualDiskArrayOutput {
	return i.ToVirtualDiskArrayOutputWithContext(context.Background())
}

func (i VirtualDiskArray) ToVirtualDiskArrayOutputWithContext(ctx context.Context) VirtualDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDiskArrayOutput)
}

type VirtualDiskOutput struct{ *pulumi.OutputState }

func (VirtualDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDisk)(nil)).Elem()
}

func (o VirtualDiskOutput) ToVirtualDiskOutput() VirtualDiskOutput {
	return o
}

func (o VirtualDiskOutput) ToVirtualDiskOutputWithContext(ctx context.Context) VirtualDiskOutput {
	return o
}

func (o VirtualDiskOutput) ControllerId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualDisk) string { return v.ControllerId }).(pulumi.StringOutput)
}

func (o VirtualDiskOutput) IndependenceMode() DiskIndependenceModeOutput {
	return o.ApplyT(func(v VirtualDisk) DiskIndependenceMode { return v.IndependenceMode }).(DiskIndependenceModeOutput)
}

func (o VirtualDiskOutput) TotalSize() pulumi.IntOutput {
	return o.ApplyT(func(v VirtualDisk) int { return v.TotalSize }).(pulumi.IntOutput)
}

func (o VirtualDiskOutput) VirtualDiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDisk) *string { return v.VirtualDiskId }).(pulumi.StringPtrOutput)
}

type VirtualDiskArrayOutput struct{ *pulumi.OutputState }

func (VirtualDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDisk)(nil)).Elem()
}

func (o VirtualDiskArrayOutput) ToVirtualDiskArrayOutput() VirtualDiskArrayOutput {
	return o
}

func (o VirtualDiskArrayOutput) ToVirtualDiskArrayOutputWithContext(ctx context.Context) VirtualDiskArrayOutput {
	return o
}

func (o VirtualDiskArrayOutput) Index(i pulumi.IntInput) VirtualDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualDisk {
		return vs[0].([]VirtualDisk)[vs[1].(int)]
	}).(VirtualDiskOutput)
}

type VirtualDiskControllerResponse struct {
	Id      string `pulumi:"id"`
	Name    string `pulumi:"name"`
	SubType string `pulumi:"subType"`
	Type    string `pulumi:"type"`
}





type VirtualDiskControllerResponseInput interface {
	pulumi.Input

	ToVirtualDiskControllerResponseOutput() VirtualDiskControllerResponseOutput
	ToVirtualDiskControllerResponseOutputWithContext(context.Context) VirtualDiskControllerResponseOutput
}

type VirtualDiskControllerResponseArgs struct {
	Id      pulumi.StringInput `pulumi:"id"`
	Name    pulumi.StringInput `pulumi:"name"`
	SubType pulumi.StringInput `pulumi:"subType"`
	Type    pulumi.StringInput `pulumi:"type"`
}

func (VirtualDiskControllerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDiskControllerResponse)(nil)).Elem()
}

func (i VirtualDiskControllerResponseArgs) ToVirtualDiskControllerResponseOutput() VirtualDiskControllerResponseOutput {
	return i.ToVirtualDiskControllerResponseOutputWithContext(context.Background())
}

func (i VirtualDiskControllerResponseArgs) ToVirtualDiskControllerResponseOutputWithContext(ctx context.Context) VirtualDiskControllerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDiskControllerResponseOutput)
}





type VirtualDiskControllerResponseArrayInput interface {
	pulumi.Input

	ToVirtualDiskControllerResponseArrayOutput() VirtualDiskControllerResponseArrayOutput
	ToVirtualDiskControllerResponseArrayOutputWithContext(context.Context) VirtualDiskControllerResponseArrayOutput
}

type VirtualDiskControllerResponseArray []VirtualDiskControllerResponseInput

func (VirtualDiskControllerResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDiskControllerResponse)(nil)).Elem()
}

func (i VirtualDiskControllerResponseArray) ToVirtualDiskControllerResponseArrayOutput() VirtualDiskControllerResponseArrayOutput {
	return i.ToVirtualDiskControllerResponseArrayOutputWithContext(context.Background())
}

func (i VirtualDiskControllerResponseArray) ToVirtualDiskControllerResponseArrayOutputWithContext(ctx context.Context) VirtualDiskControllerResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDiskControllerResponseArrayOutput)
}

type VirtualDiskControllerResponseOutput struct{ *pulumi.OutputState }

func (VirtualDiskControllerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDiskControllerResponse)(nil)).Elem()
}

func (o VirtualDiskControllerResponseOutput) ToVirtualDiskControllerResponseOutput() VirtualDiskControllerResponseOutput {
	return o
}

func (o VirtualDiskControllerResponseOutput) ToVirtualDiskControllerResponseOutputWithContext(ctx context.Context) VirtualDiskControllerResponseOutput {
	return o
}

func (o VirtualDiskControllerResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualDiskControllerResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o VirtualDiskControllerResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualDiskControllerResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualDiskControllerResponseOutput) SubType() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualDiskControllerResponse) string { return v.SubType }).(pulumi.StringOutput)
}

func (o VirtualDiskControllerResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualDiskControllerResponse) string { return v.Type }).(pulumi.StringOutput)
}

type VirtualDiskControllerResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualDiskControllerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDiskControllerResponse)(nil)).Elem()
}

func (o VirtualDiskControllerResponseArrayOutput) ToVirtualDiskControllerResponseArrayOutput() VirtualDiskControllerResponseArrayOutput {
	return o
}

func (o VirtualDiskControllerResponseArrayOutput) ToVirtualDiskControllerResponseArrayOutputWithContext(ctx context.Context) VirtualDiskControllerResponseArrayOutput {
	return o
}

func (o VirtualDiskControllerResponseArrayOutput) Index(i pulumi.IntInput) VirtualDiskControllerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualDiskControllerResponse {
		return vs[0].([]VirtualDiskControllerResponse)[vs[1].(int)]
	}).(VirtualDiskControllerResponseOutput)
}

type VirtualDiskResponse struct {
	ControllerId     string  `pulumi:"controllerId"`
	IndependenceMode string  `pulumi:"independenceMode"`
	TotalSize        int     `pulumi:"totalSize"`
	VirtualDiskId    *string `pulumi:"virtualDiskId"`
	VirtualDiskName  string  `pulumi:"virtualDiskName"`
}





type VirtualDiskResponseInput interface {
	pulumi.Input

	ToVirtualDiskResponseOutput() VirtualDiskResponseOutput
	ToVirtualDiskResponseOutputWithContext(context.Context) VirtualDiskResponseOutput
}

type VirtualDiskResponseArgs struct {
	ControllerId     pulumi.StringInput    `pulumi:"controllerId"`
	IndependenceMode pulumi.StringInput    `pulumi:"independenceMode"`
	TotalSize        pulumi.IntInput       `pulumi:"totalSize"`
	VirtualDiskId    pulumi.StringPtrInput `pulumi:"virtualDiskId"`
	VirtualDiskName  pulumi.StringInput    `pulumi:"virtualDiskName"`
}

func (VirtualDiskResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDiskResponse)(nil)).Elem()
}

func (i VirtualDiskResponseArgs) ToVirtualDiskResponseOutput() VirtualDiskResponseOutput {
	return i.ToVirtualDiskResponseOutputWithContext(context.Background())
}

func (i VirtualDiskResponseArgs) ToVirtualDiskResponseOutputWithContext(ctx context.Context) VirtualDiskResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDiskResponseOutput)
}





type VirtualDiskResponseArrayInput interface {
	pulumi.Input

	ToVirtualDiskResponseArrayOutput() VirtualDiskResponseArrayOutput
	ToVirtualDiskResponseArrayOutputWithContext(context.Context) VirtualDiskResponseArrayOutput
}

type VirtualDiskResponseArray []VirtualDiskResponseInput

func (VirtualDiskResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDiskResponse)(nil)).Elem()
}

func (i VirtualDiskResponseArray) ToVirtualDiskResponseArrayOutput() VirtualDiskResponseArrayOutput {
	return i.ToVirtualDiskResponseArrayOutputWithContext(context.Background())
}

func (i VirtualDiskResponseArray) ToVirtualDiskResponseArrayOutputWithContext(ctx context.Context) VirtualDiskResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDiskResponseArrayOutput)
}

type VirtualDiskResponseOutput struct{ *pulumi.OutputState }

func (VirtualDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDiskResponse)(nil)).Elem()
}

func (o VirtualDiskResponseOutput) ToVirtualDiskResponseOutput() VirtualDiskResponseOutput {
	return o
}

func (o VirtualDiskResponseOutput) ToVirtualDiskResponseOutputWithContext(ctx context.Context) VirtualDiskResponseOutput {
	return o
}

func (o VirtualDiskResponseOutput) ControllerId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualDiskResponse) string { return v.ControllerId }).(pulumi.StringOutput)
}

func (o VirtualDiskResponseOutput) IndependenceMode() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualDiskResponse) string { return v.IndependenceMode }).(pulumi.StringOutput)
}

func (o VirtualDiskResponseOutput) TotalSize() pulumi.IntOutput {
	return o.ApplyT(func(v VirtualDiskResponse) int { return v.TotalSize }).(pulumi.IntOutput)
}

func (o VirtualDiskResponseOutput) VirtualDiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDiskResponse) *string { return v.VirtualDiskId }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskResponseOutput) VirtualDiskName() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualDiskResponse) string { return v.VirtualDiskName }).(pulumi.StringOutput)
}

type VirtualDiskResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualDiskResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDiskResponse)(nil)).Elem()
}

func (o VirtualDiskResponseArrayOutput) ToVirtualDiskResponseArrayOutput() VirtualDiskResponseArrayOutput {
	return o
}

func (o VirtualDiskResponseArrayOutput) ToVirtualDiskResponseArrayOutputWithContext(ctx context.Context) VirtualDiskResponseArrayOutput {
	return o
}

func (o VirtualDiskResponseArrayOutput) Index(i pulumi.IntInput) VirtualDiskResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualDiskResponse {
		return vs[0].([]VirtualDiskResponse)[vs[1].(int)]
	}).(VirtualDiskResponseOutput)
}

type VirtualNetwork struct {
	Id string `pulumi:"id"`
}





type VirtualNetworkInput interface {
	pulumi.Input

	ToVirtualNetworkOutput() VirtualNetworkOutput
	ToVirtualNetworkOutputWithContext(context.Context) VirtualNetworkOutput
}

type VirtualNetworkArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (VirtualNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetwork)(nil)).Elem()
}

func (i VirtualNetworkArgs) ToVirtualNetworkOutput() VirtualNetworkOutput {
	return i.ToVirtualNetworkOutputWithContext(context.Background())
}

func (i VirtualNetworkArgs) ToVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkOutput)
}

type VirtualNetworkOutput struct{ *pulumi.OutputState }

func (VirtualNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetwork)(nil)).Elem()
}

func (o VirtualNetworkOutput) ToVirtualNetworkOutput() VirtualNetworkOutput {
	return o
}

func (o VirtualNetworkOutput) ToVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkOutput {
	return o
}

func (o VirtualNetworkOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetwork) string { return v.Id }).(pulumi.StringOutput)
}

type VirtualNetworkResponse struct {
	Assignable     bool   `pulumi:"assignable"`
	Id             string `pulumi:"id"`
	Location       string `pulumi:"location"`
	Name           string `pulumi:"name"`
	PrivateCloudId string `pulumi:"privateCloudId"`
	Type           string `pulumi:"type"`
}





type VirtualNetworkResponseInput interface {
	pulumi.Input

	ToVirtualNetworkResponseOutput() VirtualNetworkResponseOutput
	ToVirtualNetworkResponseOutputWithContext(context.Context) VirtualNetworkResponseOutput
}

type VirtualNetworkResponseArgs struct {
	Assignable     pulumi.BoolInput   `pulumi:"assignable"`
	Id             pulumi.StringInput `pulumi:"id"`
	Location       pulumi.StringInput `pulumi:"location"`
	Name           pulumi.StringInput `pulumi:"name"`
	PrivateCloudId pulumi.StringInput `pulumi:"privateCloudId"`
	Type           pulumi.StringInput `pulumi:"type"`
}

func (VirtualNetworkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkResponse)(nil)).Elem()
}

func (i VirtualNetworkResponseArgs) ToVirtualNetworkResponseOutput() VirtualNetworkResponseOutput {
	return i.ToVirtualNetworkResponseOutputWithContext(context.Background())
}

func (i VirtualNetworkResponseArgs) ToVirtualNetworkResponseOutputWithContext(ctx context.Context) VirtualNetworkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkResponseOutput)
}

type VirtualNetworkResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkResponse)(nil)).Elem()
}

func (o VirtualNetworkResponseOutput) ToVirtualNetworkResponseOutput() VirtualNetworkResponseOutput {
	return o
}

func (o VirtualNetworkResponseOutput) ToVirtualNetworkResponseOutputWithContext(ctx context.Context) VirtualNetworkResponseOutput {
	return o
}

func (o VirtualNetworkResponseOutput) Assignable() pulumi.BoolOutput {
	return o.ApplyT(func(v VirtualNetworkResponse) bool { return v.Assignable }).(pulumi.BoolOutput)
}

func (o VirtualNetworkResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o VirtualNetworkResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualNetworkResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkResponseOutput) PrivateCloudId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkResponse) string { return v.PrivateCloudId }).(pulumi.StringOutput)
}

func (o VirtualNetworkResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkResponse) string { return v.Type }).(pulumi.StringOutput)
}

type VirtualNic struct {
	Customization *GuestOSNICCustomization `pulumi:"customization"`
	IpAddresses   []string                 `pulumi:"ipAddresses"`
	MacAddress    *string                  `pulumi:"macAddress"`
	Network       VirtualNetwork           `pulumi:"network"`
	NicType       NICType                  `pulumi:"nicType"`
	PowerOnBoot   *bool                    `pulumi:"powerOnBoot"`
	VirtualNicId  *string                  `pulumi:"virtualNicId"`
}





type VirtualNicInput interface {
	pulumi.Input

	ToVirtualNicOutput() VirtualNicOutput
	ToVirtualNicOutputWithContext(context.Context) VirtualNicOutput
}

type VirtualNicArgs struct {
	Customization GuestOSNICCustomizationPtrInput `pulumi:"customization"`
	IpAddresses   pulumi.StringArrayInput         `pulumi:"ipAddresses"`
	MacAddress    pulumi.StringPtrInput           `pulumi:"macAddress"`
	Network       VirtualNetworkInput             `pulumi:"network"`
	NicType       NICTypeInput                    `pulumi:"nicType"`
	PowerOnBoot   pulumi.BoolPtrInput             `pulumi:"powerOnBoot"`
	VirtualNicId  pulumi.StringPtrInput           `pulumi:"virtualNicId"`
}

func (VirtualNicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNic)(nil)).Elem()
}

func (i VirtualNicArgs) ToVirtualNicOutput() VirtualNicOutput {
	return i.ToVirtualNicOutputWithContext(context.Background())
}

func (i VirtualNicArgs) ToVirtualNicOutputWithContext(ctx context.Context) VirtualNicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNicOutput)
}





type VirtualNicArrayInput interface {
	pulumi.Input

	ToVirtualNicArrayOutput() VirtualNicArrayOutput
	ToVirtualNicArrayOutputWithContext(context.Context) VirtualNicArrayOutput
}

type VirtualNicArray []VirtualNicInput

func (VirtualNicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNic)(nil)).Elem()
}

func (i VirtualNicArray) ToVirtualNicArrayOutput() VirtualNicArrayOutput {
	return i.ToVirtualNicArrayOutputWithContext(context.Background())
}

func (i VirtualNicArray) ToVirtualNicArrayOutputWithContext(ctx context.Context) VirtualNicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNicArrayOutput)
}

type VirtualNicOutput struct{ *pulumi.OutputState }

func (VirtualNicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNic)(nil)).Elem()
}

func (o VirtualNicOutput) ToVirtualNicOutput() VirtualNicOutput {
	return o
}

func (o VirtualNicOutput) ToVirtualNicOutputWithContext(ctx context.Context) VirtualNicOutput {
	return o
}

func (o VirtualNicOutput) Customization() GuestOSNICCustomizationPtrOutput {
	return o.ApplyT(func(v VirtualNic) *GuestOSNICCustomization { return v.Customization }).(GuestOSNICCustomizationPtrOutput)
}

func (o VirtualNicOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualNic) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o VirtualNicOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNic) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o VirtualNicOutput) Network() VirtualNetworkOutput {
	return o.ApplyT(func(v VirtualNic) VirtualNetwork { return v.Network }).(VirtualNetworkOutput)
}

func (o VirtualNicOutput) NicType() NICTypeOutput {
	return o.ApplyT(func(v VirtualNic) NICType { return v.NicType }).(NICTypeOutput)
}

func (o VirtualNicOutput) PowerOnBoot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNic) *bool { return v.PowerOnBoot }).(pulumi.BoolPtrOutput)
}

func (o VirtualNicOutput) VirtualNicId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNic) *string { return v.VirtualNicId }).(pulumi.StringPtrOutput)
}

type VirtualNicArrayOutput struct{ *pulumi.OutputState }

func (VirtualNicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNic)(nil)).Elem()
}

func (o VirtualNicArrayOutput) ToVirtualNicArrayOutput() VirtualNicArrayOutput {
	return o
}

func (o VirtualNicArrayOutput) ToVirtualNicArrayOutputWithContext(ctx context.Context) VirtualNicArrayOutput {
	return o
}

func (o VirtualNicArrayOutput) Index(i pulumi.IntInput) VirtualNicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNic {
		return vs[0].([]VirtualNic)[vs[1].(int)]
	}).(VirtualNicOutput)
}

type VirtualNicResponse struct {
	Customization  *GuestOSNICCustomizationResponse `pulumi:"customization"`
	IpAddresses    []string                         `pulumi:"ipAddresses"`
	MacAddress     *string                          `pulumi:"macAddress"`
	Network        VirtualNetworkResponse           `pulumi:"network"`
	NicType        string                           `pulumi:"nicType"`
	PowerOnBoot    *bool                            `pulumi:"powerOnBoot"`
	VirtualNicId   *string                          `pulumi:"virtualNicId"`
	VirtualNicName string                           `pulumi:"virtualNicName"`
}





type VirtualNicResponseInput interface {
	pulumi.Input

	ToVirtualNicResponseOutput() VirtualNicResponseOutput
	ToVirtualNicResponseOutputWithContext(context.Context) VirtualNicResponseOutput
}

type VirtualNicResponseArgs struct {
	Customization  GuestOSNICCustomizationResponsePtrInput `pulumi:"customization"`
	IpAddresses    pulumi.StringArrayInput                 `pulumi:"ipAddresses"`
	MacAddress     pulumi.StringPtrInput                   `pulumi:"macAddress"`
	Network        VirtualNetworkResponseInput             `pulumi:"network"`
	NicType        pulumi.StringInput                      `pulumi:"nicType"`
	PowerOnBoot    pulumi.BoolPtrInput                     `pulumi:"powerOnBoot"`
	VirtualNicId   pulumi.StringPtrInput                   `pulumi:"virtualNicId"`
	VirtualNicName pulumi.StringInput                      `pulumi:"virtualNicName"`
}

func (VirtualNicResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNicResponse)(nil)).Elem()
}

func (i VirtualNicResponseArgs) ToVirtualNicResponseOutput() VirtualNicResponseOutput {
	return i.ToVirtualNicResponseOutputWithContext(context.Background())
}

func (i VirtualNicResponseArgs) ToVirtualNicResponseOutputWithContext(ctx context.Context) VirtualNicResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNicResponseOutput)
}





type VirtualNicResponseArrayInput interface {
	pulumi.Input

	ToVirtualNicResponseArrayOutput() VirtualNicResponseArrayOutput
	ToVirtualNicResponseArrayOutputWithContext(context.Context) VirtualNicResponseArrayOutput
}

type VirtualNicResponseArray []VirtualNicResponseInput

func (VirtualNicResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNicResponse)(nil)).Elem()
}

func (i VirtualNicResponseArray) ToVirtualNicResponseArrayOutput() VirtualNicResponseArrayOutput {
	return i.ToVirtualNicResponseArrayOutputWithContext(context.Background())
}

func (i VirtualNicResponseArray) ToVirtualNicResponseArrayOutputWithContext(ctx context.Context) VirtualNicResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNicResponseArrayOutput)
}

type VirtualNicResponseOutput struct{ *pulumi.OutputState }

func (VirtualNicResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNicResponse)(nil)).Elem()
}

func (o VirtualNicResponseOutput) ToVirtualNicResponseOutput() VirtualNicResponseOutput {
	return o
}

func (o VirtualNicResponseOutput) ToVirtualNicResponseOutputWithContext(ctx context.Context) VirtualNicResponseOutput {
	return o
}

func (o VirtualNicResponseOutput) Customization() GuestOSNICCustomizationResponsePtrOutput {
	return o.ApplyT(func(v VirtualNicResponse) *GuestOSNICCustomizationResponse { return v.Customization }).(GuestOSNICCustomizationResponsePtrOutput)
}

func (o VirtualNicResponseOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualNicResponse) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o VirtualNicResponseOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNicResponse) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o VirtualNicResponseOutput) Network() VirtualNetworkResponseOutput {
	return o.ApplyT(func(v VirtualNicResponse) VirtualNetworkResponse { return v.Network }).(VirtualNetworkResponseOutput)
}

func (o VirtualNicResponseOutput) NicType() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNicResponse) string { return v.NicType }).(pulumi.StringOutput)
}

func (o VirtualNicResponseOutput) PowerOnBoot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNicResponse) *bool { return v.PowerOnBoot }).(pulumi.BoolPtrOutput)
}

func (o VirtualNicResponseOutput) VirtualNicId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNicResponse) *string { return v.VirtualNicId }).(pulumi.StringPtrOutput)
}

func (o VirtualNicResponseOutput) VirtualNicName() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNicResponse) string { return v.VirtualNicName }).(pulumi.StringOutput)
}

type VirtualNicResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNicResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNicResponse)(nil)).Elem()
}

func (o VirtualNicResponseArrayOutput) ToVirtualNicResponseArrayOutput() VirtualNicResponseArrayOutput {
	return o
}

func (o VirtualNicResponseArrayOutput) ToVirtualNicResponseArrayOutputWithContext(ctx context.Context) VirtualNicResponseArrayOutput {
	return o
}

func (o VirtualNicResponseArrayOutput) Index(i pulumi.IntInput) VirtualNicResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNicResponse {
		return vs[0].([]VirtualNicResponse)[vs[1].(int)]
	}).(VirtualNicResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(GuestOSCustomizationOutput{})
	pulumi.RegisterOutputType(GuestOSCustomizationPtrOutput{})
	pulumi.RegisterOutputType(GuestOSCustomizationResponseOutput{})
	pulumi.RegisterOutputType(GuestOSCustomizationResponsePtrOutput{})
	pulumi.RegisterOutputType(GuestOSNICCustomizationOutput{})
	pulumi.RegisterOutputType(GuestOSNICCustomizationPtrOutput{})
	pulumi.RegisterOutputType(GuestOSNICCustomizationResponseOutput{})
	pulumi.RegisterOutputType(GuestOSNICCustomizationResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourcePoolOutput{})
	pulumi.RegisterOutputType(ResourcePoolPtrOutput{})
	pulumi.RegisterOutputType(ResourcePoolResponseOutput{})
	pulumi.RegisterOutputType(ResourcePoolResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualDiskOutput{})
	pulumi.RegisterOutputType(VirtualDiskArrayOutput{})
	pulumi.RegisterOutputType(VirtualDiskControllerResponseOutput{})
	pulumi.RegisterOutputType(VirtualDiskControllerResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualDiskResponseOutput{})
	pulumi.RegisterOutputType(VirtualDiskResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkOutput{})
	pulumi.RegisterOutputType(VirtualNetworkResponseOutput{})
	pulumi.RegisterOutputType(VirtualNicOutput{})
	pulumi.RegisterOutputType(VirtualNicArrayOutput{})
	pulumi.RegisterOutputType(VirtualNicResponseOutput{})
	pulumi.RegisterOutputType(VirtualNicResponseArrayOutput{})
}

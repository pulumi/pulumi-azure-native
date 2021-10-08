


package peering

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BgpSession struct {
	MaxPrefixesAdvertisedV4     *int    `pulumi:"maxPrefixesAdvertisedV4"`
	MaxPrefixesAdvertisedV6     *int    `pulumi:"maxPrefixesAdvertisedV6"`
	Md5AuthenticationKey        *string `pulumi:"md5AuthenticationKey"`
	MicrosoftSessionIPv4Address *string `pulumi:"microsoftSessionIPv4Address"`
	MicrosoftSessionIPv6Address *string `pulumi:"microsoftSessionIPv6Address"`
	PeerSessionIPv4Address      *string `pulumi:"peerSessionIPv4Address"`
	PeerSessionIPv6Address      *string `pulumi:"peerSessionIPv6Address"`
	SessionPrefixV4             *string `pulumi:"sessionPrefixV4"`
	SessionPrefixV6             *string `pulumi:"sessionPrefixV6"`
}





type BgpSessionInput interface {
	pulumi.Input

	ToBgpSessionOutput() BgpSessionOutput
	ToBgpSessionOutputWithContext(context.Context) BgpSessionOutput
}

type BgpSessionArgs struct {
	MaxPrefixesAdvertisedV4     pulumi.IntPtrInput    `pulumi:"maxPrefixesAdvertisedV4"`
	MaxPrefixesAdvertisedV6     pulumi.IntPtrInput    `pulumi:"maxPrefixesAdvertisedV6"`
	Md5AuthenticationKey        pulumi.StringPtrInput `pulumi:"md5AuthenticationKey"`
	MicrosoftSessionIPv4Address pulumi.StringPtrInput `pulumi:"microsoftSessionIPv4Address"`
	MicrosoftSessionIPv6Address pulumi.StringPtrInput `pulumi:"microsoftSessionIPv6Address"`
	PeerSessionIPv4Address      pulumi.StringPtrInput `pulumi:"peerSessionIPv4Address"`
	PeerSessionIPv6Address      pulumi.StringPtrInput `pulumi:"peerSessionIPv6Address"`
	SessionPrefixV4             pulumi.StringPtrInput `pulumi:"sessionPrefixV4"`
	SessionPrefixV6             pulumi.StringPtrInput `pulumi:"sessionPrefixV6"`
}

func (BgpSessionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BgpSession)(nil)).Elem()
}

func (i BgpSessionArgs) ToBgpSessionOutput() BgpSessionOutput {
	return i.ToBgpSessionOutputWithContext(context.Background())
}

func (i BgpSessionArgs) ToBgpSessionOutputWithContext(ctx context.Context) BgpSessionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpSessionOutput)
}

func (i BgpSessionArgs) ToBgpSessionPtrOutput() BgpSessionPtrOutput {
	return i.ToBgpSessionPtrOutputWithContext(context.Background())
}

func (i BgpSessionArgs) ToBgpSessionPtrOutputWithContext(ctx context.Context) BgpSessionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpSessionOutput).ToBgpSessionPtrOutputWithContext(ctx)
}









type BgpSessionPtrInput interface {
	pulumi.Input

	ToBgpSessionPtrOutput() BgpSessionPtrOutput
	ToBgpSessionPtrOutputWithContext(context.Context) BgpSessionPtrOutput
}

type bgpSessionPtrType BgpSessionArgs

func BgpSessionPtr(v *BgpSessionArgs) BgpSessionPtrInput {
	return (*bgpSessionPtrType)(v)
}

func (*bgpSessionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpSession)(nil)).Elem()
}

func (i *bgpSessionPtrType) ToBgpSessionPtrOutput() BgpSessionPtrOutput {
	return i.ToBgpSessionPtrOutputWithContext(context.Background())
}

func (i *bgpSessionPtrType) ToBgpSessionPtrOutputWithContext(ctx context.Context) BgpSessionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpSessionPtrOutput)
}

type BgpSessionOutput struct{ *pulumi.OutputState }

func (BgpSessionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BgpSession)(nil)).Elem()
}

func (o BgpSessionOutput) ToBgpSessionOutput() BgpSessionOutput {
	return o
}

func (o BgpSessionOutput) ToBgpSessionOutputWithContext(ctx context.Context) BgpSessionOutput {
	return o
}

func (o BgpSessionOutput) ToBgpSessionPtrOutput() BgpSessionPtrOutput {
	return o.ToBgpSessionPtrOutputWithContext(context.Background())
}

func (o BgpSessionOutput) ToBgpSessionPtrOutputWithContext(ctx context.Context) BgpSessionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BgpSession) *BgpSession {
		return &v
	}).(BgpSessionPtrOutput)
}

func (o BgpSessionOutput) MaxPrefixesAdvertisedV4() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BgpSession) *int { return v.MaxPrefixesAdvertisedV4 }).(pulumi.IntPtrOutput)
}

func (o BgpSessionOutput) MaxPrefixesAdvertisedV6() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BgpSession) *int { return v.MaxPrefixesAdvertisedV6 }).(pulumi.IntPtrOutput)
}

func (o BgpSessionOutput) Md5AuthenticationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpSession) *string { return v.Md5AuthenticationKey }).(pulumi.StringPtrOutput)
}

func (o BgpSessionOutput) MicrosoftSessionIPv4Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpSession) *string { return v.MicrosoftSessionIPv4Address }).(pulumi.StringPtrOutput)
}

func (o BgpSessionOutput) MicrosoftSessionIPv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpSession) *string { return v.MicrosoftSessionIPv6Address }).(pulumi.StringPtrOutput)
}

func (o BgpSessionOutput) PeerSessionIPv4Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpSession) *string { return v.PeerSessionIPv4Address }).(pulumi.StringPtrOutput)
}

func (o BgpSessionOutput) PeerSessionIPv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpSession) *string { return v.PeerSessionIPv6Address }).(pulumi.StringPtrOutput)
}

func (o BgpSessionOutput) SessionPrefixV4() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpSession) *string { return v.SessionPrefixV4 }).(pulumi.StringPtrOutput)
}

func (o BgpSessionOutput) SessionPrefixV6() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpSession) *string { return v.SessionPrefixV6 }).(pulumi.StringPtrOutput)
}

type BgpSessionPtrOutput struct{ *pulumi.OutputState }

func (BgpSessionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpSession)(nil)).Elem()
}

func (o BgpSessionPtrOutput) ToBgpSessionPtrOutput() BgpSessionPtrOutput {
	return o
}

func (o BgpSessionPtrOutput) ToBgpSessionPtrOutputWithContext(ctx context.Context) BgpSessionPtrOutput {
	return o
}

func (o BgpSessionPtrOutput) Elem() BgpSessionOutput {
	return o.ApplyT(func(v *BgpSession) BgpSession {
		if v != nil {
			return *v
		}
		var ret BgpSession
		return ret
	}).(BgpSessionOutput)
}

func (o BgpSessionPtrOutput) MaxPrefixesAdvertisedV4() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BgpSession) *int {
		if v == nil {
			return nil
		}
		return v.MaxPrefixesAdvertisedV4
	}).(pulumi.IntPtrOutput)
}

func (o BgpSessionPtrOutput) MaxPrefixesAdvertisedV6() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BgpSession) *int {
		if v == nil {
			return nil
		}
		return v.MaxPrefixesAdvertisedV6
	}).(pulumi.IntPtrOutput)
}

func (o BgpSessionPtrOutput) Md5AuthenticationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSession) *string {
		if v == nil {
			return nil
		}
		return v.Md5AuthenticationKey
	}).(pulumi.StringPtrOutput)
}

func (o BgpSessionPtrOutput) MicrosoftSessionIPv4Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSession) *string {
		if v == nil {
			return nil
		}
		return v.MicrosoftSessionIPv4Address
	}).(pulumi.StringPtrOutput)
}

func (o BgpSessionPtrOutput) MicrosoftSessionIPv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSession) *string {
		if v == nil {
			return nil
		}
		return v.MicrosoftSessionIPv6Address
	}).(pulumi.StringPtrOutput)
}

func (o BgpSessionPtrOutput) PeerSessionIPv4Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSession) *string {
		if v == nil {
			return nil
		}
		return v.PeerSessionIPv4Address
	}).(pulumi.StringPtrOutput)
}

func (o BgpSessionPtrOutput) PeerSessionIPv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSession) *string {
		if v == nil {
			return nil
		}
		return v.PeerSessionIPv6Address
	}).(pulumi.StringPtrOutput)
}

func (o BgpSessionPtrOutput) SessionPrefixV4() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSession) *string {
		if v == nil {
			return nil
		}
		return v.SessionPrefixV4
	}).(pulumi.StringPtrOutput)
}

func (o BgpSessionPtrOutput) SessionPrefixV6() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSession) *string {
		if v == nil {
			return nil
		}
		return v.SessionPrefixV6
	}).(pulumi.StringPtrOutput)
}

type BgpSessionResponse struct {
	MaxPrefixesAdvertisedV4     *int    `pulumi:"maxPrefixesAdvertisedV4"`
	MaxPrefixesAdvertisedV6     *int    `pulumi:"maxPrefixesAdvertisedV6"`
	Md5AuthenticationKey        *string `pulumi:"md5AuthenticationKey"`
	MicrosoftSessionIPv4Address *string `pulumi:"microsoftSessionIPv4Address"`
	MicrosoftSessionIPv6Address *string `pulumi:"microsoftSessionIPv6Address"`
	PeerSessionIPv4Address      *string `pulumi:"peerSessionIPv4Address"`
	PeerSessionIPv6Address      *string `pulumi:"peerSessionIPv6Address"`
	SessionPrefixV4             *string `pulumi:"sessionPrefixV4"`
	SessionPrefixV6             *string `pulumi:"sessionPrefixV6"`
	SessionStateV4              string  `pulumi:"sessionStateV4"`
	SessionStateV6              string  `pulumi:"sessionStateV6"`
}





type BgpSessionResponseInput interface {
	pulumi.Input

	ToBgpSessionResponseOutput() BgpSessionResponseOutput
	ToBgpSessionResponseOutputWithContext(context.Context) BgpSessionResponseOutput
}

type BgpSessionResponseArgs struct {
	MaxPrefixesAdvertisedV4     pulumi.IntPtrInput    `pulumi:"maxPrefixesAdvertisedV4"`
	MaxPrefixesAdvertisedV6     pulumi.IntPtrInput    `pulumi:"maxPrefixesAdvertisedV6"`
	Md5AuthenticationKey        pulumi.StringPtrInput `pulumi:"md5AuthenticationKey"`
	MicrosoftSessionIPv4Address pulumi.StringPtrInput `pulumi:"microsoftSessionIPv4Address"`
	MicrosoftSessionIPv6Address pulumi.StringPtrInput `pulumi:"microsoftSessionIPv6Address"`
	PeerSessionIPv4Address      pulumi.StringPtrInput `pulumi:"peerSessionIPv4Address"`
	PeerSessionIPv6Address      pulumi.StringPtrInput `pulumi:"peerSessionIPv6Address"`
	SessionPrefixV4             pulumi.StringPtrInput `pulumi:"sessionPrefixV4"`
	SessionPrefixV6             pulumi.StringPtrInput `pulumi:"sessionPrefixV6"`
	SessionStateV4              pulumi.StringInput    `pulumi:"sessionStateV4"`
	SessionStateV6              pulumi.StringInput    `pulumi:"sessionStateV6"`
}

func (BgpSessionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BgpSessionResponse)(nil)).Elem()
}

func (i BgpSessionResponseArgs) ToBgpSessionResponseOutput() BgpSessionResponseOutput {
	return i.ToBgpSessionResponseOutputWithContext(context.Background())
}

func (i BgpSessionResponseArgs) ToBgpSessionResponseOutputWithContext(ctx context.Context) BgpSessionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpSessionResponseOutput)
}

func (i BgpSessionResponseArgs) ToBgpSessionResponsePtrOutput() BgpSessionResponsePtrOutput {
	return i.ToBgpSessionResponsePtrOutputWithContext(context.Background())
}

func (i BgpSessionResponseArgs) ToBgpSessionResponsePtrOutputWithContext(ctx context.Context) BgpSessionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpSessionResponseOutput).ToBgpSessionResponsePtrOutputWithContext(ctx)
}









type BgpSessionResponsePtrInput interface {
	pulumi.Input

	ToBgpSessionResponsePtrOutput() BgpSessionResponsePtrOutput
	ToBgpSessionResponsePtrOutputWithContext(context.Context) BgpSessionResponsePtrOutput
}

type bgpSessionResponsePtrType BgpSessionResponseArgs

func BgpSessionResponsePtr(v *BgpSessionResponseArgs) BgpSessionResponsePtrInput {
	return (*bgpSessionResponsePtrType)(v)
}

func (*bgpSessionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpSessionResponse)(nil)).Elem()
}

func (i *bgpSessionResponsePtrType) ToBgpSessionResponsePtrOutput() BgpSessionResponsePtrOutput {
	return i.ToBgpSessionResponsePtrOutputWithContext(context.Background())
}

func (i *bgpSessionResponsePtrType) ToBgpSessionResponsePtrOutputWithContext(ctx context.Context) BgpSessionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpSessionResponsePtrOutput)
}

type BgpSessionResponseOutput struct{ *pulumi.OutputState }

func (BgpSessionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BgpSessionResponse)(nil)).Elem()
}

func (o BgpSessionResponseOutput) ToBgpSessionResponseOutput() BgpSessionResponseOutput {
	return o
}

func (o BgpSessionResponseOutput) ToBgpSessionResponseOutputWithContext(ctx context.Context) BgpSessionResponseOutput {
	return o
}

func (o BgpSessionResponseOutput) ToBgpSessionResponsePtrOutput() BgpSessionResponsePtrOutput {
	return o.ToBgpSessionResponsePtrOutputWithContext(context.Background())
}

func (o BgpSessionResponseOutput) ToBgpSessionResponsePtrOutputWithContext(ctx context.Context) BgpSessionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BgpSessionResponse) *BgpSessionResponse {
		return &v
	}).(BgpSessionResponsePtrOutput)
}

func (o BgpSessionResponseOutput) MaxPrefixesAdvertisedV4() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BgpSessionResponse) *int { return v.MaxPrefixesAdvertisedV4 }).(pulumi.IntPtrOutput)
}

func (o BgpSessionResponseOutput) MaxPrefixesAdvertisedV6() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BgpSessionResponse) *int { return v.MaxPrefixesAdvertisedV6 }).(pulumi.IntPtrOutput)
}

func (o BgpSessionResponseOutput) Md5AuthenticationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpSessionResponse) *string { return v.Md5AuthenticationKey }).(pulumi.StringPtrOutput)
}

func (o BgpSessionResponseOutput) MicrosoftSessionIPv4Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpSessionResponse) *string { return v.MicrosoftSessionIPv4Address }).(pulumi.StringPtrOutput)
}

func (o BgpSessionResponseOutput) MicrosoftSessionIPv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpSessionResponse) *string { return v.MicrosoftSessionIPv6Address }).(pulumi.StringPtrOutput)
}

func (o BgpSessionResponseOutput) PeerSessionIPv4Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpSessionResponse) *string { return v.PeerSessionIPv4Address }).(pulumi.StringPtrOutput)
}

func (o BgpSessionResponseOutput) PeerSessionIPv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpSessionResponse) *string { return v.PeerSessionIPv6Address }).(pulumi.StringPtrOutput)
}

func (o BgpSessionResponseOutput) SessionPrefixV4() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpSessionResponse) *string { return v.SessionPrefixV4 }).(pulumi.StringPtrOutput)
}

func (o BgpSessionResponseOutput) SessionPrefixV6() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpSessionResponse) *string { return v.SessionPrefixV6 }).(pulumi.StringPtrOutput)
}

func (o BgpSessionResponseOutput) SessionStateV4() pulumi.StringOutput {
	return o.ApplyT(func(v BgpSessionResponse) string { return v.SessionStateV4 }).(pulumi.StringOutput)
}

func (o BgpSessionResponseOutput) SessionStateV6() pulumi.StringOutput {
	return o.ApplyT(func(v BgpSessionResponse) string { return v.SessionStateV6 }).(pulumi.StringOutput)
}

type BgpSessionResponsePtrOutput struct{ *pulumi.OutputState }

func (BgpSessionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpSessionResponse)(nil)).Elem()
}

func (o BgpSessionResponsePtrOutput) ToBgpSessionResponsePtrOutput() BgpSessionResponsePtrOutput {
	return o
}

func (o BgpSessionResponsePtrOutput) ToBgpSessionResponsePtrOutputWithContext(ctx context.Context) BgpSessionResponsePtrOutput {
	return o
}

func (o BgpSessionResponsePtrOutput) Elem() BgpSessionResponseOutput {
	return o.ApplyT(func(v *BgpSessionResponse) BgpSessionResponse {
		if v != nil {
			return *v
		}
		var ret BgpSessionResponse
		return ret
	}).(BgpSessionResponseOutput)
}

func (o BgpSessionResponsePtrOutput) MaxPrefixesAdvertisedV4() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BgpSessionResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxPrefixesAdvertisedV4
	}).(pulumi.IntPtrOutput)
}

func (o BgpSessionResponsePtrOutput) MaxPrefixesAdvertisedV6() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BgpSessionResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxPrefixesAdvertisedV6
	}).(pulumi.IntPtrOutput)
}

func (o BgpSessionResponsePtrOutput) Md5AuthenticationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSessionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Md5AuthenticationKey
	}).(pulumi.StringPtrOutput)
}

func (o BgpSessionResponsePtrOutput) MicrosoftSessionIPv4Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSessionResponse) *string {
		if v == nil {
			return nil
		}
		return v.MicrosoftSessionIPv4Address
	}).(pulumi.StringPtrOutput)
}

func (o BgpSessionResponsePtrOutput) MicrosoftSessionIPv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSessionResponse) *string {
		if v == nil {
			return nil
		}
		return v.MicrosoftSessionIPv6Address
	}).(pulumi.StringPtrOutput)
}

func (o BgpSessionResponsePtrOutput) PeerSessionIPv4Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSessionResponse) *string {
		if v == nil {
			return nil
		}
		return v.PeerSessionIPv4Address
	}).(pulumi.StringPtrOutput)
}

func (o BgpSessionResponsePtrOutput) PeerSessionIPv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSessionResponse) *string {
		if v == nil {
			return nil
		}
		return v.PeerSessionIPv6Address
	}).(pulumi.StringPtrOutput)
}

func (o BgpSessionResponsePtrOutput) SessionPrefixV4() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSessionResponse) *string {
		if v == nil {
			return nil
		}
		return v.SessionPrefixV4
	}).(pulumi.StringPtrOutput)
}

func (o BgpSessionResponsePtrOutput) SessionPrefixV6() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSessionResponse) *string {
		if v == nil {
			return nil
		}
		return v.SessionPrefixV6
	}).(pulumi.StringPtrOutput)
}

func (o BgpSessionResponsePtrOutput) SessionStateV4() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSessionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SessionStateV4
	}).(pulumi.StringPtrOutput)
}

func (o BgpSessionResponsePtrOutput) SessionStateV6() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSessionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SessionStateV6
	}).(pulumi.StringPtrOutput)
}

type ContactDetail struct {
	Email *string `pulumi:"email"`
	Phone *string `pulumi:"phone"`
	Role  *string `pulumi:"role"`
}





type ContactDetailInput interface {
	pulumi.Input

	ToContactDetailOutput() ContactDetailOutput
	ToContactDetailOutputWithContext(context.Context) ContactDetailOutput
}

type ContactDetailArgs struct {
	Email pulumi.StringPtrInput `pulumi:"email"`
	Phone pulumi.StringPtrInput `pulumi:"phone"`
	Role  pulumi.StringPtrInput `pulumi:"role"`
}

func (ContactDetailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactDetail)(nil)).Elem()
}

func (i ContactDetailArgs) ToContactDetailOutput() ContactDetailOutput {
	return i.ToContactDetailOutputWithContext(context.Background())
}

func (i ContactDetailArgs) ToContactDetailOutputWithContext(ctx context.Context) ContactDetailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactDetailOutput)
}





type ContactDetailArrayInput interface {
	pulumi.Input

	ToContactDetailArrayOutput() ContactDetailArrayOutput
	ToContactDetailArrayOutputWithContext(context.Context) ContactDetailArrayOutput
}

type ContactDetailArray []ContactDetailInput

func (ContactDetailArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContactDetail)(nil)).Elem()
}

func (i ContactDetailArray) ToContactDetailArrayOutput() ContactDetailArrayOutput {
	return i.ToContactDetailArrayOutputWithContext(context.Background())
}

func (i ContactDetailArray) ToContactDetailArrayOutputWithContext(ctx context.Context) ContactDetailArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactDetailArrayOutput)
}

type ContactDetailOutput struct{ *pulumi.OutputState }

func (ContactDetailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactDetail)(nil)).Elem()
}

func (o ContactDetailOutput) ToContactDetailOutput() ContactDetailOutput {
	return o
}

func (o ContactDetailOutput) ToContactDetailOutputWithContext(ctx context.Context) ContactDetailOutput {
	return o
}

func (o ContactDetailOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetail) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o ContactDetailOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetail) *string { return v.Phone }).(pulumi.StringPtrOutput)
}

func (o ContactDetailOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetail) *string { return v.Role }).(pulumi.StringPtrOutput)
}

type ContactDetailArrayOutput struct{ *pulumi.OutputState }

func (ContactDetailArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContactDetail)(nil)).Elem()
}

func (o ContactDetailArrayOutput) ToContactDetailArrayOutput() ContactDetailArrayOutput {
	return o
}

func (o ContactDetailArrayOutput) ToContactDetailArrayOutputWithContext(ctx context.Context) ContactDetailArrayOutput {
	return o
}

func (o ContactDetailArrayOutput) Index(i pulumi.IntInput) ContactDetailOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContactDetail {
		return vs[0].([]ContactDetail)[vs[1].(int)]
	}).(ContactDetailOutput)
}

type ContactDetailResponse struct {
	Email *string `pulumi:"email"`
	Phone *string `pulumi:"phone"`
	Role  *string `pulumi:"role"`
}





type ContactDetailResponseInput interface {
	pulumi.Input

	ToContactDetailResponseOutput() ContactDetailResponseOutput
	ToContactDetailResponseOutputWithContext(context.Context) ContactDetailResponseOutput
}

type ContactDetailResponseArgs struct {
	Email pulumi.StringPtrInput `pulumi:"email"`
	Phone pulumi.StringPtrInput `pulumi:"phone"`
	Role  pulumi.StringPtrInput `pulumi:"role"`
}

func (ContactDetailResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactDetailResponse)(nil)).Elem()
}

func (i ContactDetailResponseArgs) ToContactDetailResponseOutput() ContactDetailResponseOutput {
	return i.ToContactDetailResponseOutputWithContext(context.Background())
}

func (i ContactDetailResponseArgs) ToContactDetailResponseOutputWithContext(ctx context.Context) ContactDetailResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactDetailResponseOutput)
}





type ContactDetailResponseArrayInput interface {
	pulumi.Input

	ToContactDetailResponseArrayOutput() ContactDetailResponseArrayOutput
	ToContactDetailResponseArrayOutputWithContext(context.Context) ContactDetailResponseArrayOutput
}

type ContactDetailResponseArray []ContactDetailResponseInput

func (ContactDetailResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContactDetailResponse)(nil)).Elem()
}

func (i ContactDetailResponseArray) ToContactDetailResponseArrayOutput() ContactDetailResponseArrayOutput {
	return i.ToContactDetailResponseArrayOutputWithContext(context.Background())
}

func (i ContactDetailResponseArray) ToContactDetailResponseArrayOutputWithContext(ctx context.Context) ContactDetailResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactDetailResponseArrayOutput)
}

type ContactDetailResponseOutput struct{ *pulumi.OutputState }

func (ContactDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactDetailResponse)(nil)).Elem()
}

func (o ContactDetailResponseOutput) ToContactDetailResponseOutput() ContactDetailResponseOutput {
	return o
}

func (o ContactDetailResponseOutput) ToContactDetailResponseOutputWithContext(ctx context.Context) ContactDetailResponseOutput {
	return o
}

func (o ContactDetailResponseOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetailResponse) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o ContactDetailResponseOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetailResponse) *string { return v.Phone }).(pulumi.StringPtrOutput)
}

func (o ContactDetailResponseOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetailResponse) *string { return v.Role }).(pulumi.StringPtrOutput)
}

type ContactDetailResponseArrayOutput struct{ *pulumi.OutputState }

func (ContactDetailResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContactDetailResponse)(nil)).Elem()
}

func (o ContactDetailResponseArrayOutput) ToContactDetailResponseArrayOutput() ContactDetailResponseArrayOutput {
	return o
}

func (o ContactDetailResponseArrayOutput) ToContactDetailResponseArrayOutputWithContext(ctx context.Context) ContactDetailResponseArrayOutput {
	return o
}

func (o ContactDetailResponseArrayOutput) Index(i pulumi.IntInput) ContactDetailResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContactDetailResponse {
		return vs[0].([]ContactDetailResponse)[vs[1].(int)]
	}).(ContactDetailResponseOutput)
}

type DirectConnection struct {
	BandwidthInMbps        *int        `pulumi:"bandwidthInMbps"`
	BgpSession             *BgpSession `pulumi:"bgpSession"`
	ConnectionIdentifier   *string     `pulumi:"connectionIdentifier"`
	PeeringDBFacilityId    *int        `pulumi:"peeringDBFacilityId"`
	SessionAddressProvider *string     `pulumi:"sessionAddressProvider"`
	UseForPeeringService   *bool       `pulumi:"useForPeeringService"`
}





type DirectConnectionInput interface {
	pulumi.Input

	ToDirectConnectionOutput() DirectConnectionOutput
	ToDirectConnectionOutputWithContext(context.Context) DirectConnectionOutput
}

type DirectConnectionArgs struct {
	BandwidthInMbps        pulumi.IntPtrInput    `pulumi:"bandwidthInMbps"`
	BgpSession             BgpSessionPtrInput    `pulumi:"bgpSession"`
	ConnectionIdentifier   pulumi.StringPtrInput `pulumi:"connectionIdentifier"`
	PeeringDBFacilityId    pulumi.IntPtrInput    `pulumi:"peeringDBFacilityId"`
	SessionAddressProvider pulumi.StringPtrInput `pulumi:"sessionAddressProvider"`
	UseForPeeringService   pulumi.BoolPtrInput   `pulumi:"useForPeeringService"`
}

func (DirectConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectConnection)(nil)).Elem()
}

func (i DirectConnectionArgs) ToDirectConnectionOutput() DirectConnectionOutput {
	return i.ToDirectConnectionOutputWithContext(context.Background())
}

func (i DirectConnectionArgs) ToDirectConnectionOutputWithContext(ctx context.Context) DirectConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectConnectionOutput)
}





type DirectConnectionArrayInput interface {
	pulumi.Input

	ToDirectConnectionArrayOutput() DirectConnectionArrayOutput
	ToDirectConnectionArrayOutputWithContext(context.Context) DirectConnectionArrayOutput
}

type DirectConnectionArray []DirectConnectionInput

func (DirectConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DirectConnection)(nil)).Elem()
}

func (i DirectConnectionArray) ToDirectConnectionArrayOutput() DirectConnectionArrayOutput {
	return i.ToDirectConnectionArrayOutputWithContext(context.Background())
}

func (i DirectConnectionArray) ToDirectConnectionArrayOutputWithContext(ctx context.Context) DirectConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectConnectionArrayOutput)
}

type DirectConnectionOutput struct{ *pulumi.OutputState }

func (DirectConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectConnection)(nil)).Elem()
}

func (o DirectConnectionOutput) ToDirectConnectionOutput() DirectConnectionOutput {
	return o
}

func (o DirectConnectionOutput) ToDirectConnectionOutputWithContext(ctx context.Context) DirectConnectionOutput {
	return o
}

func (o DirectConnectionOutput) BandwidthInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DirectConnection) *int { return v.BandwidthInMbps }).(pulumi.IntPtrOutput)
}

func (o DirectConnectionOutput) BgpSession() BgpSessionPtrOutput {
	return o.ApplyT(func(v DirectConnection) *BgpSession { return v.BgpSession }).(BgpSessionPtrOutput)
}

func (o DirectConnectionOutput) ConnectionIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DirectConnection) *string { return v.ConnectionIdentifier }).(pulumi.StringPtrOutput)
}

func (o DirectConnectionOutput) PeeringDBFacilityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DirectConnection) *int { return v.PeeringDBFacilityId }).(pulumi.IntPtrOutput)
}

func (o DirectConnectionOutput) SessionAddressProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DirectConnection) *string { return v.SessionAddressProvider }).(pulumi.StringPtrOutput)
}

func (o DirectConnectionOutput) UseForPeeringService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DirectConnection) *bool { return v.UseForPeeringService }).(pulumi.BoolPtrOutput)
}

type DirectConnectionArrayOutput struct{ *pulumi.OutputState }

func (DirectConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DirectConnection)(nil)).Elem()
}

func (o DirectConnectionArrayOutput) ToDirectConnectionArrayOutput() DirectConnectionArrayOutput {
	return o
}

func (o DirectConnectionArrayOutput) ToDirectConnectionArrayOutputWithContext(ctx context.Context) DirectConnectionArrayOutput {
	return o
}

func (o DirectConnectionArrayOutput) Index(i pulumi.IntInput) DirectConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DirectConnection {
		return vs[0].([]DirectConnection)[vs[1].(int)]
	}).(DirectConnectionOutput)
}

type DirectConnectionResponse struct {
	BandwidthInMbps            *int                `pulumi:"bandwidthInMbps"`
	BgpSession                 *BgpSessionResponse `pulumi:"bgpSession"`
	ConnectionIdentifier       *string             `pulumi:"connectionIdentifier"`
	ConnectionState            string              `pulumi:"connectionState"`
	ErrorMessage               string              `pulumi:"errorMessage"`
	MicrosoftTrackingId        string              `pulumi:"microsoftTrackingId"`
	PeeringDBFacilityId        *int                `pulumi:"peeringDBFacilityId"`
	ProvisionedBandwidthInMbps int                 `pulumi:"provisionedBandwidthInMbps"`
	SessionAddressProvider     *string             `pulumi:"sessionAddressProvider"`
	UseForPeeringService       *bool               `pulumi:"useForPeeringService"`
}





type DirectConnectionResponseInput interface {
	pulumi.Input

	ToDirectConnectionResponseOutput() DirectConnectionResponseOutput
	ToDirectConnectionResponseOutputWithContext(context.Context) DirectConnectionResponseOutput
}

type DirectConnectionResponseArgs struct {
	BandwidthInMbps            pulumi.IntPtrInput         `pulumi:"bandwidthInMbps"`
	BgpSession                 BgpSessionResponsePtrInput `pulumi:"bgpSession"`
	ConnectionIdentifier       pulumi.StringPtrInput      `pulumi:"connectionIdentifier"`
	ConnectionState            pulumi.StringInput         `pulumi:"connectionState"`
	ErrorMessage               pulumi.StringInput         `pulumi:"errorMessage"`
	MicrosoftTrackingId        pulumi.StringInput         `pulumi:"microsoftTrackingId"`
	PeeringDBFacilityId        pulumi.IntPtrInput         `pulumi:"peeringDBFacilityId"`
	ProvisionedBandwidthInMbps pulumi.IntInput            `pulumi:"provisionedBandwidthInMbps"`
	SessionAddressProvider     pulumi.StringPtrInput      `pulumi:"sessionAddressProvider"`
	UseForPeeringService       pulumi.BoolPtrInput        `pulumi:"useForPeeringService"`
}

func (DirectConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectConnectionResponse)(nil)).Elem()
}

func (i DirectConnectionResponseArgs) ToDirectConnectionResponseOutput() DirectConnectionResponseOutput {
	return i.ToDirectConnectionResponseOutputWithContext(context.Background())
}

func (i DirectConnectionResponseArgs) ToDirectConnectionResponseOutputWithContext(ctx context.Context) DirectConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectConnectionResponseOutput)
}





type DirectConnectionResponseArrayInput interface {
	pulumi.Input

	ToDirectConnectionResponseArrayOutput() DirectConnectionResponseArrayOutput
	ToDirectConnectionResponseArrayOutputWithContext(context.Context) DirectConnectionResponseArrayOutput
}

type DirectConnectionResponseArray []DirectConnectionResponseInput

func (DirectConnectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DirectConnectionResponse)(nil)).Elem()
}

func (i DirectConnectionResponseArray) ToDirectConnectionResponseArrayOutput() DirectConnectionResponseArrayOutput {
	return i.ToDirectConnectionResponseArrayOutputWithContext(context.Background())
}

func (i DirectConnectionResponseArray) ToDirectConnectionResponseArrayOutputWithContext(ctx context.Context) DirectConnectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectConnectionResponseArrayOutput)
}

type DirectConnectionResponseOutput struct{ *pulumi.OutputState }

func (DirectConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectConnectionResponse)(nil)).Elem()
}

func (o DirectConnectionResponseOutput) ToDirectConnectionResponseOutput() DirectConnectionResponseOutput {
	return o
}

func (o DirectConnectionResponseOutput) ToDirectConnectionResponseOutputWithContext(ctx context.Context) DirectConnectionResponseOutput {
	return o
}

func (o DirectConnectionResponseOutput) BandwidthInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DirectConnectionResponse) *int { return v.BandwidthInMbps }).(pulumi.IntPtrOutput)
}

func (o DirectConnectionResponseOutput) BgpSession() BgpSessionResponsePtrOutput {
	return o.ApplyT(func(v DirectConnectionResponse) *BgpSessionResponse { return v.BgpSession }).(BgpSessionResponsePtrOutput)
}

func (o DirectConnectionResponseOutput) ConnectionIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DirectConnectionResponse) *string { return v.ConnectionIdentifier }).(pulumi.StringPtrOutput)
}

func (o DirectConnectionResponseOutput) ConnectionState() pulumi.StringOutput {
	return o.ApplyT(func(v DirectConnectionResponse) string { return v.ConnectionState }).(pulumi.StringOutput)
}

func (o DirectConnectionResponseOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v DirectConnectionResponse) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o DirectConnectionResponseOutput) MicrosoftTrackingId() pulumi.StringOutput {
	return o.ApplyT(func(v DirectConnectionResponse) string { return v.MicrosoftTrackingId }).(pulumi.StringOutput)
}

func (o DirectConnectionResponseOutput) PeeringDBFacilityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DirectConnectionResponse) *int { return v.PeeringDBFacilityId }).(pulumi.IntPtrOutput)
}

func (o DirectConnectionResponseOutput) ProvisionedBandwidthInMbps() pulumi.IntOutput {
	return o.ApplyT(func(v DirectConnectionResponse) int { return v.ProvisionedBandwidthInMbps }).(pulumi.IntOutput)
}

func (o DirectConnectionResponseOutput) SessionAddressProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DirectConnectionResponse) *string { return v.SessionAddressProvider }).(pulumi.StringPtrOutput)
}

func (o DirectConnectionResponseOutput) UseForPeeringService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DirectConnectionResponse) *bool { return v.UseForPeeringService }).(pulumi.BoolPtrOutput)
}

type DirectConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (DirectConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DirectConnectionResponse)(nil)).Elem()
}

func (o DirectConnectionResponseArrayOutput) ToDirectConnectionResponseArrayOutput() DirectConnectionResponseArrayOutput {
	return o
}

func (o DirectConnectionResponseArrayOutput) ToDirectConnectionResponseArrayOutputWithContext(ctx context.Context) DirectConnectionResponseArrayOutput {
	return o
}

func (o DirectConnectionResponseArrayOutput) Index(i pulumi.IntInput) DirectConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DirectConnectionResponse {
		return vs[0].([]DirectConnectionResponse)[vs[1].(int)]
	}).(DirectConnectionResponseOutput)
}

type ExchangeConnection struct {
	BgpSession           *BgpSession `pulumi:"bgpSession"`
	ConnectionIdentifier *string     `pulumi:"connectionIdentifier"`
	PeeringDBFacilityId  *int        `pulumi:"peeringDBFacilityId"`
}





type ExchangeConnectionInput interface {
	pulumi.Input

	ToExchangeConnectionOutput() ExchangeConnectionOutput
	ToExchangeConnectionOutputWithContext(context.Context) ExchangeConnectionOutput
}

type ExchangeConnectionArgs struct {
	BgpSession           BgpSessionPtrInput    `pulumi:"bgpSession"`
	ConnectionIdentifier pulumi.StringPtrInput `pulumi:"connectionIdentifier"`
	PeeringDBFacilityId  pulumi.IntPtrInput    `pulumi:"peeringDBFacilityId"`
}

func (ExchangeConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExchangeConnection)(nil)).Elem()
}

func (i ExchangeConnectionArgs) ToExchangeConnectionOutput() ExchangeConnectionOutput {
	return i.ToExchangeConnectionOutputWithContext(context.Background())
}

func (i ExchangeConnectionArgs) ToExchangeConnectionOutputWithContext(ctx context.Context) ExchangeConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExchangeConnectionOutput)
}





type ExchangeConnectionArrayInput interface {
	pulumi.Input

	ToExchangeConnectionArrayOutput() ExchangeConnectionArrayOutput
	ToExchangeConnectionArrayOutputWithContext(context.Context) ExchangeConnectionArrayOutput
}

type ExchangeConnectionArray []ExchangeConnectionInput

func (ExchangeConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExchangeConnection)(nil)).Elem()
}

func (i ExchangeConnectionArray) ToExchangeConnectionArrayOutput() ExchangeConnectionArrayOutput {
	return i.ToExchangeConnectionArrayOutputWithContext(context.Background())
}

func (i ExchangeConnectionArray) ToExchangeConnectionArrayOutputWithContext(ctx context.Context) ExchangeConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExchangeConnectionArrayOutput)
}

type ExchangeConnectionOutput struct{ *pulumi.OutputState }

func (ExchangeConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExchangeConnection)(nil)).Elem()
}

func (o ExchangeConnectionOutput) ToExchangeConnectionOutput() ExchangeConnectionOutput {
	return o
}

func (o ExchangeConnectionOutput) ToExchangeConnectionOutputWithContext(ctx context.Context) ExchangeConnectionOutput {
	return o
}

func (o ExchangeConnectionOutput) BgpSession() BgpSessionPtrOutput {
	return o.ApplyT(func(v ExchangeConnection) *BgpSession { return v.BgpSession }).(BgpSessionPtrOutput)
}

func (o ExchangeConnectionOutput) ConnectionIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExchangeConnection) *string { return v.ConnectionIdentifier }).(pulumi.StringPtrOutput)
}

func (o ExchangeConnectionOutput) PeeringDBFacilityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExchangeConnection) *int { return v.PeeringDBFacilityId }).(pulumi.IntPtrOutput)
}

type ExchangeConnectionArrayOutput struct{ *pulumi.OutputState }

func (ExchangeConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExchangeConnection)(nil)).Elem()
}

func (o ExchangeConnectionArrayOutput) ToExchangeConnectionArrayOutput() ExchangeConnectionArrayOutput {
	return o
}

func (o ExchangeConnectionArrayOutput) ToExchangeConnectionArrayOutputWithContext(ctx context.Context) ExchangeConnectionArrayOutput {
	return o
}

func (o ExchangeConnectionArrayOutput) Index(i pulumi.IntInput) ExchangeConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExchangeConnection {
		return vs[0].([]ExchangeConnection)[vs[1].(int)]
	}).(ExchangeConnectionOutput)
}

type ExchangeConnectionResponse struct {
	BgpSession           *BgpSessionResponse `pulumi:"bgpSession"`
	ConnectionIdentifier *string             `pulumi:"connectionIdentifier"`
	ConnectionState      string              `pulumi:"connectionState"`
	ErrorMessage         string              `pulumi:"errorMessage"`
	PeeringDBFacilityId  *int                `pulumi:"peeringDBFacilityId"`
}





type ExchangeConnectionResponseInput interface {
	pulumi.Input

	ToExchangeConnectionResponseOutput() ExchangeConnectionResponseOutput
	ToExchangeConnectionResponseOutputWithContext(context.Context) ExchangeConnectionResponseOutput
}

type ExchangeConnectionResponseArgs struct {
	BgpSession           BgpSessionResponsePtrInput `pulumi:"bgpSession"`
	ConnectionIdentifier pulumi.StringPtrInput      `pulumi:"connectionIdentifier"`
	ConnectionState      pulumi.StringInput         `pulumi:"connectionState"`
	ErrorMessage         pulumi.StringInput         `pulumi:"errorMessage"`
	PeeringDBFacilityId  pulumi.IntPtrInput         `pulumi:"peeringDBFacilityId"`
}

func (ExchangeConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExchangeConnectionResponse)(nil)).Elem()
}

func (i ExchangeConnectionResponseArgs) ToExchangeConnectionResponseOutput() ExchangeConnectionResponseOutput {
	return i.ToExchangeConnectionResponseOutputWithContext(context.Background())
}

func (i ExchangeConnectionResponseArgs) ToExchangeConnectionResponseOutputWithContext(ctx context.Context) ExchangeConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExchangeConnectionResponseOutput)
}





type ExchangeConnectionResponseArrayInput interface {
	pulumi.Input

	ToExchangeConnectionResponseArrayOutput() ExchangeConnectionResponseArrayOutput
	ToExchangeConnectionResponseArrayOutputWithContext(context.Context) ExchangeConnectionResponseArrayOutput
}

type ExchangeConnectionResponseArray []ExchangeConnectionResponseInput

func (ExchangeConnectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExchangeConnectionResponse)(nil)).Elem()
}

func (i ExchangeConnectionResponseArray) ToExchangeConnectionResponseArrayOutput() ExchangeConnectionResponseArrayOutput {
	return i.ToExchangeConnectionResponseArrayOutputWithContext(context.Background())
}

func (i ExchangeConnectionResponseArray) ToExchangeConnectionResponseArrayOutputWithContext(ctx context.Context) ExchangeConnectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExchangeConnectionResponseArrayOutput)
}

type ExchangeConnectionResponseOutput struct{ *pulumi.OutputState }

func (ExchangeConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExchangeConnectionResponse)(nil)).Elem()
}

func (o ExchangeConnectionResponseOutput) ToExchangeConnectionResponseOutput() ExchangeConnectionResponseOutput {
	return o
}

func (o ExchangeConnectionResponseOutput) ToExchangeConnectionResponseOutputWithContext(ctx context.Context) ExchangeConnectionResponseOutput {
	return o
}

func (o ExchangeConnectionResponseOutput) BgpSession() BgpSessionResponsePtrOutput {
	return o.ApplyT(func(v ExchangeConnectionResponse) *BgpSessionResponse { return v.BgpSession }).(BgpSessionResponsePtrOutput)
}

func (o ExchangeConnectionResponseOutput) ConnectionIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExchangeConnectionResponse) *string { return v.ConnectionIdentifier }).(pulumi.StringPtrOutput)
}

func (o ExchangeConnectionResponseOutput) ConnectionState() pulumi.StringOutput {
	return o.ApplyT(func(v ExchangeConnectionResponse) string { return v.ConnectionState }).(pulumi.StringOutput)
}

func (o ExchangeConnectionResponseOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v ExchangeConnectionResponse) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o ExchangeConnectionResponseOutput) PeeringDBFacilityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExchangeConnectionResponse) *int { return v.PeeringDBFacilityId }).(pulumi.IntPtrOutput)
}

type ExchangeConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (ExchangeConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExchangeConnectionResponse)(nil)).Elem()
}

func (o ExchangeConnectionResponseArrayOutput) ToExchangeConnectionResponseArrayOutput() ExchangeConnectionResponseArrayOutput {
	return o
}

func (o ExchangeConnectionResponseArrayOutput) ToExchangeConnectionResponseArrayOutputWithContext(ctx context.Context) ExchangeConnectionResponseArrayOutput {
	return o
}

func (o ExchangeConnectionResponseArrayOutput) Index(i pulumi.IntInput) ExchangeConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExchangeConnectionResponse {
		return vs[0].([]ExchangeConnectionResponse)[vs[1].(int)]
	}).(ExchangeConnectionResponseOutput)
}

type PeeringPropertiesDirect struct {
	Connections       []DirectConnection `pulumi:"connections"`
	DirectPeeringType *string            `pulumi:"directPeeringType"`
	PeerAsn           *SubResource       `pulumi:"peerAsn"`
}





type PeeringPropertiesDirectInput interface {
	pulumi.Input

	ToPeeringPropertiesDirectOutput() PeeringPropertiesDirectOutput
	ToPeeringPropertiesDirectOutputWithContext(context.Context) PeeringPropertiesDirectOutput
}

type PeeringPropertiesDirectArgs struct {
	Connections       DirectConnectionArrayInput `pulumi:"connections"`
	DirectPeeringType pulumi.StringPtrInput      `pulumi:"directPeeringType"`
	PeerAsn           SubResourcePtrInput        `pulumi:"peerAsn"`
}

func (PeeringPropertiesDirectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PeeringPropertiesDirect)(nil)).Elem()
}

func (i PeeringPropertiesDirectArgs) ToPeeringPropertiesDirectOutput() PeeringPropertiesDirectOutput {
	return i.ToPeeringPropertiesDirectOutputWithContext(context.Background())
}

func (i PeeringPropertiesDirectArgs) ToPeeringPropertiesDirectOutputWithContext(ctx context.Context) PeeringPropertiesDirectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringPropertiesDirectOutput)
}

func (i PeeringPropertiesDirectArgs) ToPeeringPropertiesDirectPtrOutput() PeeringPropertiesDirectPtrOutput {
	return i.ToPeeringPropertiesDirectPtrOutputWithContext(context.Background())
}

func (i PeeringPropertiesDirectArgs) ToPeeringPropertiesDirectPtrOutputWithContext(ctx context.Context) PeeringPropertiesDirectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringPropertiesDirectOutput).ToPeeringPropertiesDirectPtrOutputWithContext(ctx)
}









type PeeringPropertiesDirectPtrInput interface {
	pulumi.Input

	ToPeeringPropertiesDirectPtrOutput() PeeringPropertiesDirectPtrOutput
	ToPeeringPropertiesDirectPtrOutputWithContext(context.Context) PeeringPropertiesDirectPtrOutput
}

type peeringPropertiesDirectPtrType PeeringPropertiesDirectArgs

func PeeringPropertiesDirectPtr(v *PeeringPropertiesDirectArgs) PeeringPropertiesDirectPtrInput {
	return (*peeringPropertiesDirectPtrType)(v)
}

func (*peeringPropertiesDirectPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringPropertiesDirect)(nil)).Elem()
}

func (i *peeringPropertiesDirectPtrType) ToPeeringPropertiesDirectPtrOutput() PeeringPropertiesDirectPtrOutput {
	return i.ToPeeringPropertiesDirectPtrOutputWithContext(context.Background())
}

func (i *peeringPropertiesDirectPtrType) ToPeeringPropertiesDirectPtrOutputWithContext(ctx context.Context) PeeringPropertiesDirectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringPropertiesDirectPtrOutput)
}

type PeeringPropertiesDirectOutput struct{ *pulumi.OutputState }

func (PeeringPropertiesDirectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PeeringPropertiesDirect)(nil)).Elem()
}

func (o PeeringPropertiesDirectOutput) ToPeeringPropertiesDirectOutput() PeeringPropertiesDirectOutput {
	return o
}

func (o PeeringPropertiesDirectOutput) ToPeeringPropertiesDirectOutputWithContext(ctx context.Context) PeeringPropertiesDirectOutput {
	return o
}

func (o PeeringPropertiesDirectOutput) ToPeeringPropertiesDirectPtrOutput() PeeringPropertiesDirectPtrOutput {
	return o.ToPeeringPropertiesDirectPtrOutputWithContext(context.Background())
}

func (o PeeringPropertiesDirectOutput) ToPeeringPropertiesDirectPtrOutputWithContext(ctx context.Context) PeeringPropertiesDirectPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PeeringPropertiesDirect) *PeeringPropertiesDirect {
		return &v
	}).(PeeringPropertiesDirectPtrOutput)
}

func (o PeeringPropertiesDirectOutput) Connections() DirectConnectionArrayOutput {
	return o.ApplyT(func(v PeeringPropertiesDirect) []DirectConnection { return v.Connections }).(DirectConnectionArrayOutput)
}

func (o PeeringPropertiesDirectOutput) DirectPeeringType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PeeringPropertiesDirect) *string { return v.DirectPeeringType }).(pulumi.StringPtrOutput)
}

func (o PeeringPropertiesDirectOutput) PeerAsn() SubResourcePtrOutput {
	return o.ApplyT(func(v PeeringPropertiesDirect) *SubResource { return v.PeerAsn }).(SubResourcePtrOutput)
}

type PeeringPropertiesDirectPtrOutput struct{ *pulumi.OutputState }

func (PeeringPropertiesDirectPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringPropertiesDirect)(nil)).Elem()
}

func (o PeeringPropertiesDirectPtrOutput) ToPeeringPropertiesDirectPtrOutput() PeeringPropertiesDirectPtrOutput {
	return o
}

func (o PeeringPropertiesDirectPtrOutput) ToPeeringPropertiesDirectPtrOutputWithContext(ctx context.Context) PeeringPropertiesDirectPtrOutput {
	return o
}

func (o PeeringPropertiesDirectPtrOutput) Elem() PeeringPropertiesDirectOutput {
	return o.ApplyT(func(v *PeeringPropertiesDirect) PeeringPropertiesDirect {
		if v != nil {
			return *v
		}
		var ret PeeringPropertiesDirect
		return ret
	}).(PeeringPropertiesDirectOutput)
}

func (o PeeringPropertiesDirectPtrOutput) Connections() DirectConnectionArrayOutput {
	return o.ApplyT(func(v *PeeringPropertiesDirect) []DirectConnection {
		if v == nil {
			return nil
		}
		return v.Connections
	}).(DirectConnectionArrayOutput)
}

func (o PeeringPropertiesDirectPtrOutput) DirectPeeringType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeeringPropertiesDirect) *string {
		if v == nil {
			return nil
		}
		return v.DirectPeeringType
	}).(pulumi.StringPtrOutput)
}

func (o PeeringPropertiesDirectPtrOutput) PeerAsn() SubResourcePtrOutput {
	return o.ApplyT(func(v *PeeringPropertiesDirect) *SubResource {
		if v == nil {
			return nil
		}
		return v.PeerAsn
	}).(SubResourcePtrOutput)
}

type PeeringPropertiesDirectResponse struct {
	Connections          []DirectConnectionResponse `pulumi:"connections"`
	DirectPeeringType    *string                    `pulumi:"directPeeringType"`
	PeerAsn              *SubResourceResponse       `pulumi:"peerAsn"`
	UseForPeeringService bool                       `pulumi:"useForPeeringService"`
}





type PeeringPropertiesDirectResponseInput interface {
	pulumi.Input

	ToPeeringPropertiesDirectResponseOutput() PeeringPropertiesDirectResponseOutput
	ToPeeringPropertiesDirectResponseOutputWithContext(context.Context) PeeringPropertiesDirectResponseOutput
}

type PeeringPropertiesDirectResponseArgs struct {
	Connections          DirectConnectionResponseArrayInput `pulumi:"connections"`
	DirectPeeringType    pulumi.StringPtrInput              `pulumi:"directPeeringType"`
	PeerAsn              SubResourceResponsePtrInput        `pulumi:"peerAsn"`
	UseForPeeringService pulumi.BoolInput                   `pulumi:"useForPeeringService"`
}

func (PeeringPropertiesDirectResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PeeringPropertiesDirectResponse)(nil)).Elem()
}

func (i PeeringPropertiesDirectResponseArgs) ToPeeringPropertiesDirectResponseOutput() PeeringPropertiesDirectResponseOutput {
	return i.ToPeeringPropertiesDirectResponseOutputWithContext(context.Background())
}

func (i PeeringPropertiesDirectResponseArgs) ToPeeringPropertiesDirectResponseOutputWithContext(ctx context.Context) PeeringPropertiesDirectResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringPropertiesDirectResponseOutput)
}

func (i PeeringPropertiesDirectResponseArgs) ToPeeringPropertiesDirectResponsePtrOutput() PeeringPropertiesDirectResponsePtrOutput {
	return i.ToPeeringPropertiesDirectResponsePtrOutputWithContext(context.Background())
}

func (i PeeringPropertiesDirectResponseArgs) ToPeeringPropertiesDirectResponsePtrOutputWithContext(ctx context.Context) PeeringPropertiesDirectResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringPropertiesDirectResponseOutput).ToPeeringPropertiesDirectResponsePtrOutputWithContext(ctx)
}









type PeeringPropertiesDirectResponsePtrInput interface {
	pulumi.Input

	ToPeeringPropertiesDirectResponsePtrOutput() PeeringPropertiesDirectResponsePtrOutput
	ToPeeringPropertiesDirectResponsePtrOutputWithContext(context.Context) PeeringPropertiesDirectResponsePtrOutput
}

type peeringPropertiesDirectResponsePtrType PeeringPropertiesDirectResponseArgs

func PeeringPropertiesDirectResponsePtr(v *PeeringPropertiesDirectResponseArgs) PeeringPropertiesDirectResponsePtrInput {
	return (*peeringPropertiesDirectResponsePtrType)(v)
}

func (*peeringPropertiesDirectResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringPropertiesDirectResponse)(nil)).Elem()
}

func (i *peeringPropertiesDirectResponsePtrType) ToPeeringPropertiesDirectResponsePtrOutput() PeeringPropertiesDirectResponsePtrOutput {
	return i.ToPeeringPropertiesDirectResponsePtrOutputWithContext(context.Background())
}

func (i *peeringPropertiesDirectResponsePtrType) ToPeeringPropertiesDirectResponsePtrOutputWithContext(ctx context.Context) PeeringPropertiesDirectResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringPropertiesDirectResponsePtrOutput)
}

type PeeringPropertiesDirectResponseOutput struct{ *pulumi.OutputState }

func (PeeringPropertiesDirectResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PeeringPropertiesDirectResponse)(nil)).Elem()
}

func (o PeeringPropertiesDirectResponseOutput) ToPeeringPropertiesDirectResponseOutput() PeeringPropertiesDirectResponseOutput {
	return o
}

func (o PeeringPropertiesDirectResponseOutput) ToPeeringPropertiesDirectResponseOutputWithContext(ctx context.Context) PeeringPropertiesDirectResponseOutput {
	return o
}

func (o PeeringPropertiesDirectResponseOutput) ToPeeringPropertiesDirectResponsePtrOutput() PeeringPropertiesDirectResponsePtrOutput {
	return o.ToPeeringPropertiesDirectResponsePtrOutputWithContext(context.Background())
}

func (o PeeringPropertiesDirectResponseOutput) ToPeeringPropertiesDirectResponsePtrOutputWithContext(ctx context.Context) PeeringPropertiesDirectResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PeeringPropertiesDirectResponse) *PeeringPropertiesDirectResponse {
		return &v
	}).(PeeringPropertiesDirectResponsePtrOutput)
}

func (o PeeringPropertiesDirectResponseOutput) Connections() DirectConnectionResponseArrayOutput {
	return o.ApplyT(func(v PeeringPropertiesDirectResponse) []DirectConnectionResponse { return v.Connections }).(DirectConnectionResponseArrayOutput)
}

func (o PeeringPropertiesDirectResponseOutput) DirectPeeringType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PeeringPropertiesDirectResponse) *string { return v.DirectPeeringType }).(pulumi.StringPtrOutput)
}

func (o PeeringPropertiesDirectResponseOutput) PeerAsn() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v PeeringPropertiesDirectResponse) *SubResourceResponse { return v.PeerAsn }).(SubResourceResponsePtrOutput)
}

func (o PeeringPropertiesDirectResponseOutput) UseForPeeringService() pulumi.BoolOutput {
	return o.ApplyT(func(v PeeringPropertiesDirectResponse) bool { return v.UseForPeeringService }).(pulumi.BoolOutput)
}

type PeeringPropertiesDirectResponsePtrOutput struct{ *pulumi.OutputState }

func (PeeringPropertiesDirectResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringPropertiesDirectResponse)(nil)).Elem()
}

func (o PeeringPropertiesDirectResponsePtrOutput) ToPeeringPropertiesDirectResponsePtrOutput() PeeringPropertiesDirectResponsePtrOutput {
	return o
}

func (o PeeringPropertiesDirectResponsePtrOutput) ToPeeringPropertiesDirectResponsePtrOutputWithContext(ctx context.Context) PeeringPropertiesDirectResponsePtrOutput {
	return o
}

func (o PeeringPropertiesDirectResponsePtrOutput) Elem() PeeringPropertiesDirectResponseOutput {
	return o.ApplyT(func(v *PeeringPropertiesDirectResponse) PeeringPropertiesDirectResponse {
		if v != nil {
			return *v
		}
		var ret PeeringPropertiesDirectResponse
		return ret
	}).(PeeringPropertiesDirectResponseOutput)
}

func (o PeeringPropertiesDirectResponsePtrOutput) Connections() DirectConnectionResponseArrayOutput {
	return o.ApplyT(func(v *PeeringPropertiesDirectResponse) []DirectConnectionResponse {
		if v == nil {
			return nil
		}
		return v.Connections
	}).(DirectConnectionResponseArrayOutput)
}

func (o PeeringPropertiesDirectResponsePtrOutput) DirectPeeringType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeeringPropertiesDirectResponse) *string {
		if v == nil {
			return nil
		}
		return v.DirectPeeringType
	}).(pulumi.StringPtrOutput)
}

func (o PeeringPropertiesDirectResponsePtrOutput) PeerAsn() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *PeeringPropertiesDirectResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return v.PeerAsn
	}).(SubResourceResponsePtrOutput)
}

func (o PeeringPropertiesDirectResponsePtrOutput) UseForPeeringService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PeeringPropertiesDirectResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.UseForPeeringService
	}).(pulumi.BoolPtrOutput)
}

type PeeringPropertiesExchange struct {
	Connections []ExchangeConnection `pulumi:"connections"`
	PeerAsn     *SubResource         `pulumi:"peerAsn"`
}





type PeeringPropertiesExchangeInput interface {
	pulumi.Input

	ToPeeringPropertiesExchangeOutput() PeeringPropertiesExchangeOutput
	ToPeeringPropertiesExchangeOutputWithContext(context.Context) PeeringPropertiesExchangeOutput
}

type PeeringPropertiesExchangeArgs struct {
	Connections ExchangeConnectionArrayInput `pulumi:"connections"`
	PeerAsn     SubResourcePtrInput          `pulumi:"peerAsn"`
}

func (PeeringPropertiesExchangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PeeringPropertiesExchange)(nil)).Elem()
}

func (i PeeringPropertiesExchangeArgs) ToPeeringPropertiesExchangeOutput() PeeringPropertiesExchangeOutput {
	return i.ToPeeringPropertiesExchangeOutputWithContext(context.Background())
}

func (i PeeringPropertiesExchangeArgs) ToPeeringPropertiesExchangeOutputWithContext(ctx context.Context) PeeringPropertiesExchangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringPropertiesExchangeOutput)
}

func (i PeeringPropertiesExchangeArgs) ToPeeringPropertiesExchangePtrOutput() PeeringPropertiesExchangePtrOutput {
	return i.ToPeeringPropertiesExchangePtrOutputWithContext(context.Background())
}

func (i PeeringPropertiesExchangeArgs) ToPeeringPropertiesExchangePtrOutputWithContext(ctx context.Context) PeeringPropertiesExchangePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringPropertiesExchangeOutput).ToPeeringPropertiesExchangePtrOutputWithContext(ctx)
}









type PeeringPropertiesExchangePtrInput interface {
	pulumi.Input

	ToPeeringPropertiesExchangePtrOutput() PeeringPropertiesExchangePtrOutput
	ToPeeringPropertiesExchangePtrOutputWithContext(context.Context) PeeringPropertiesExchangePtrOutput
}

type peeringPropertiesExchangePtrType PeeringPropertiesExchangeArgs

func PeeringPropertiesExchangePtr(v *PeeringPropertiesExchangeArgs) PeeringPropertiesExchangePtrInput {
	return (*peeringPropertiesExchangePtrType)(v)
}

func (*peeringPropertiesExchangePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringPropertiesExchange)(nil)).Elem()
}

func (i *peeringPropertiesExchangePtrType) ToPeeringPropertiesExchangePtrOutput() PeeringPropertiesExchangePtrOutput {
	return i.ToPeeringPropertiesExchangePtrOutputWithContext(context.Background())
}

func (i *peeringPropertiesExchangePtrType) ToPeeringPropertiesExchangePtrOutputWithContext(ctx context.Context) PeeringPropertiesExchangePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringPropertiesExchangePtrOutput)
}

type PeeringPropertiesExchangeOutput struct{ *pulumi.OutputState }

func (PeeringPropertiesExchangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PeeringPropertiesExchange)(nil)).Elem()
}

func (o PeeringPropertiesExchangeOutput) ToPeeringPropertiesExchangeOutput() PeeringPropertiesExchangeOutput {
	return o
}

func (o PeeringPropertiesExchangeOutput) ToPeeringPropertiesExchangeOutputWithContext(ctx context.Context) PeeringPropertiesExchangeOutput {
	return o
}

func (o PeeringPropertiesExchangeOutput) ToPeeringPropertiesExchangePtrOutput() PeeringPropertiesExchangePtrOutput {
	return o.ToPeeringPropertiesExchangePtrOutputWithContext(context.Background())
}

func (o PeeringPropertiesExchangeOutput) ToPeeringPropertiesExchangePtrOutputWithContext(ctx context.Context) PeeringPropertiesExchangePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PeeringPropertiesExchange) *PeeringPropertiesExchange {
		return &v
	}).(PeeringPropertiesExchangePtrOutput)
}

func (o PeeringPropertiesExchangeOutput) Connections() ExchangeConnectionArrayOutput {
	return o.ApplyT(func(v PeeringPropertiesExchange) []ExchangeConnection { return v.Connections }).(ExchangeConnectionArrayOutput)
}

func (o PeeringPropertiesExchangeOutput) PeerAsn() SubResourcePtrOutput {
	return o.ApplyT(func(v PeeringPropertiesExchange) *SubResource { return v.PeerAsn }).(SubResourcePtrOutput)
}

type PeeringPropertiesExchangePtrOutput struct{ *pulumi.OutputState }

func (PeeringPropertiesExchangePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringPropertiesExchange)(nil)).Elem()
}

func (o PeeringPropertiesExchangePtrOutput) ToPeeringPropertiesExchangePtrOutput() PeeringPropertiesExchangePtrOutput {
	return o
}

func (o PeeringPropertiesExchangePtrOutput) ToPeeringPropertiesExchangePtrOutputWithContext(ctx context.Context) PeeringPropertiesExchangePtrOutput {
	return o
}

func (o PeeringPropertiesExchangePtrOutput) Elem() PeeringPropertiesExchangeOutput {
	return o.ApplyT(func(v *PeeringPropertiesExchange) PeeringPropertiesExchange {
		if v != nil {
			return *v
		}
		var ret PeeringPropertiesExchange
		return ret
	}).(PeeringPropertiesExchangeOutput)
}

func (o PeeringPropertiesExchangePtrOutput) Connections() ExchangeConnectionArrayOutput {
	return o.ApplyT(func(v *PeeringPropertiesExchange) []ExchangeConnection {
		if v == nil {
			return nil
		}
		return v.Connections
	}).(ExchangeConnectionArrayOutput)
}

func (o PeeringPropertiesExchangePtrOutput) PeerAsn() SubResourcePtrOutput {
	return o.ApplyT(func(v *PeeringPropertiesExchange) *SubResource {
		if v == nil {
			return nil
		}
		return v.PeerAsn
	}).(SubResourcePtrOutput)
}

type PeeringPropertiesExchangeResponse struct {
	Connections []ExchangeConnectionResponse `pulumi:"connections"`
	PeerAsn     *SubResourceResponse         `pulumi:"peerAsn"`
}





type PeeringPropertiesExchangeResponseInput interface {
	pulumi.Input

	ToPeeringPropertiesExchangeResponseOutput() PeeringPropertiesExchangeResponseOutput
	ToPeeringPropertiesExchangeResponseOutputWithContext(context.Context) PeeringPropertiesExchangeResponseOutput
}

type PeeringPropertiesExchangeResponseArgs struct {
	Connections ExchangeConnectionResponseArrayInput `pulumi:"connections"`
	PeerAsn     SubResourceResponsePtrInput          `pulumi:"peerAsn"`
}

func (PeeringPropertiesExchangeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PeeringPropertiesExchangeResponse)(nil)).Elem()
}

func (i PeeringPropertiesExchangeResponseArgs) ToPeeringPropertiesExchangeResponseOutput() PeeringPropertiesExchangeResponseOutput {
	return i.ToPeeringPropertiesExchangeResponseOutputWithContext(context.Background())
}

func (i PeeringPropertiesExchangeResponseArgs) ToPeeringPropertiesExchangeResponseOutputWithContext(ctx context.Context) PeeringPropertiesExchangeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringPropertiesExchangeResponseOutput)
}

func (i PeeringPropertiesExchangeResponseArgs) ToPeeringPropertiesExchangeResponsePtrOutput() PeeringPropertiesExchangeResponsePtrOutput {
	return i.ToPeeringPropertiesExchangeResponsePtrOutputWithContext(context.Background())
}

func (i PeeringPropertiesExchangeResponseArgs) ToPeeringPropertiesExchangeResponsePtrOutputWithContext(ctx context.Context) PeeringPropertiesExchangeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringPropertiesExchangeResponseOutput).ToPeeringPropertiesExchangeResponsePtrOutputWithContext(ctx)
}









type PeeringPropertiesExchangeResponsePtrInput interface {
	pulumi.Input

	ToPeeringPropertiesExchangeResponsePtrOutput() PeeringPropertiesExchangeResponsePtrOutput
	ToPeeringPropertiesExchangeResponsePtrOutputWithContext(context.Context) PeeringPropertiesExchangeResponsePtrOutput
}

type peeringPropertiesExchangeResponsePtrType PeeringPropertiesExchangeResponseArgs

func PeeringPropertiesExchangeResponsePtr(v *PeeringPropertiesExchangeResponseArgs) PeeringPropertiesExchangeResponsePtrInput {
	return (*peeringPropertiesExchangeResponsePtrType)(v)
}

func (*peeringPropertiesExchangeResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringPropertiesExchangeResponse)(nil)).Elem()
}

func (i *peeringPropertiesExchangeResponsePtrType) ToPeeringPropertiesExchangeResponsePtrOutput() PeeringPropertiesExchangeResponsePtrOutput {
	return i.ToPeeringPropertiesExchangeResponsePtrOutputWithContext(context.Background())
}

func (i *peeringPropertiesExchangeResponsePtrType) ToPeeringPropertiesExchangeResponsePtrOutputWithContext(ctx context.Context) PeeringPropertiesExchangeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringPropertiesExchangeResponsePtrOutput)
}

type PeeringPropertiesExchangeResponseOutput struct{ *pulumi.OutputState }

func (PeeringPropertiesExchangeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PeeringPropertiesExchangeResponse)(nil)).Elem()
}

func (o PeeringPropertiesExchangeResponseOutput) ToPeeringPropertiesExchangeResponseOutput() PeeringPropertiesExchangeResponseOutput {
	return o
}

func (o PeeringPropertiesExchangeResponseOutput) ToPeeringPropertiesExchangeResponseOutputWithContext(ctx context.Context) PeeringPropertiesExchangeResponseOutput {
	return o
}

func (o PeeringPropertiesExchangeResponseOutput) ToPeeringPropertiesExchangeResponsePtrOutput() PeeringPropertiesExchangeResponsePtrOutput {
	return o.ToPeeringPropertiesExchangeResponsePtrOutputWithContext(context.Background())
}

func (o PeeringPropertiesExchangeResponseOutput) ToPeeringPropertiesExchangeResponsePtrOutputWithContext(ctx context.Context) PeeringPropertiesExchangeResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PeeringPropertiesExchangeResponse) *PeeringPropertiesExchangeResponse {
		return &v
	}).(PeeringPropertiesExchangeResponsePtrOutput)
}

func (o PeeringPropertiesExchangeResponseOutput) Connections() ExchangeConnectionResponseArrayOutput {
	return o.ApplyT(func(v PeeringPropertiesExchangeResponse) []ExchangeConnectionResponse { return v.Connections }).(ExchangeConnectionResponseArrayOutput)
}

func (o PeeringPropertiesExchangeResponseOutput) PeerAsn() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v PeeringPropertiesExchangeResponse) *SubResourceResponse { return v.PeerAsn }).(SubResourceResponsePtrOutput)
}

type PeeringPropertiesExchangeResponsePtrOutput struct{ *pulumi.OutputState }

func (PeeringPropertiesExchangeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringPropertiesExchangeResponse)(nil)).Elem()
}

func (o PeeringPropertiesExchangeResponsePtrOutput) ToPeeringPropertiesExchangeResponsePtrOutput() PeeringPropertiesExchangeResponsePtrOutput {
	return o
}

func (o PeeringPropertiesExchangeResponsePtrOutput) ToPeeringPropertiesExchangeResponsePtrOutputWithContext(ctx context.Context) PeeringPropertiesExchangeResponsePtrOutput {
	return o
}

func (o PeeringPropertiesExchangeResponsePtrOutput) Elem() PeeringPropertiesExchangeResponseOutput {
	return o.ApplyT(func(v *PeeringPropertiesExchangeResponse) PeeringPropertiesExchangeResponse {
		if v != nil {
			return *v
		}
		var ret PeeringPropertiesExchangeResponse
		return ret
	}).(PeeringPropertiesExchangeResponseOutput)
}

func (o PeeringPropertiesExchangeResponsePtrOutput) Connections() ExchangeConnectionResponseArrayOutput {
	return o.ApplyT(func(v *PeeringPropertiesExchangeResponse) []ExchangeConnectionResponse {
		if v == nil {
			return nil
		}
		return v.Connections
	}).(ExchangeConnectionResponseArrayOutput)
}

func (o PeeringPropertiesExchangeResponsePtrOutput) PeerAsn() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *PeeringPropertiesExchangeResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return v.PeerAsn
	}).(SubResourceResponsePtrOutput)
}

type PeeringServicePrefixEventResponse struct {
	EventDescription string `pulumi:"eventDescription"`
	EventLevel       string `pulumi:"eventLevel"`
	EventSummary     string `pulumi:"eventSummary"`
	EventTimestamp   string `pulumi:"eventTimestamp"`
	EventType        string `pulumi:"eventType"`
}





type PeeringServicePrefixEventResponseInput interface {
	pulumi.Input

	ToPeeringServicePrefixEventResponseOutput() PeeringServicePrefixEventResponseOutput
	ToPeeringServicePrefixEventResponseOutputWithContext(context.Context) PeeringServicePrefixEventResponseOutput
}

type PeeringServicePrefixEventResponseArgs struct {
	EventDescription pulumi.StringInput `pulumi:"eventDescription"`
	EventLevel       pulumi.StringInput `pulumi:"eventLevel"`
	EventSummary     pulumi.StringInput `pulumi:"eventSummary"`
	EventTimestamp   pulumi.StringInput `pulumi:"eventTimestamp"`
	EventType        pulumi.StringInput `pulumi:"eventType"`
}

func (PeeringServicePrefixEventResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PeeringServicePrefixEventResponse)(nil)).Elem()
}

func (i PeeringServicePrefixEventResponseArgs) ToPeeringServicePrefixEventResponseOutput() PeeringServicePrefixEventResponseOutput {
	return i.ToPeeringServicePrefixEventResponseOutputWithContext(context.Background())
}

func (i PeeringServicePrefixEventResponseArgs) ToPeeringServicePrefixEventResponseOutputWithContext(ctx context.Context) PeeringServicePrefixEventResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringServicePrefixEventResponseOutput)
}





type PeeringServicePrefixEventResponseArrayInput interface {
	pulumi.Input

	ToPeeringServicePrefixEventResponseArrayOutput() PeeringServicePrefixEventResponseArrayOutput
	ToPeeringServicePrefixEventResponseArrayOutputWithContext(context.Context) PeeringServicePrefixEventResponseArrayOutput
}

type PeeringServicePrefixEventResponseArray []PeeringServicePrefixEventResponseInput

func (PeeringServicePrefixEventResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PeeringServicePrefixEventResponse)(nil)).Elem()
}

func (i PeeringServicePrefixEventResponseArray) ToPeeringServicePrefixEventResponseArrayOutput() PeeringServicePrefixEventResponseArrayOutput {
	return i.ToPeeringServicePrefixEventResponseArrayOutputWithContext(context.Background())
}

func (i PeeringServicePrefixEventResponseArray) ToPeeringServicePrefixEventResponseArrayOutputWithContext(ctx context.Context) PeeringServicePrefixEventResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringServicePrefixEventResponseArrayOutput)
}

type PeeringServicePrefixEventResponseOutput struct{ *pulumi.OutputState }

func (PeeringServicePrefixEventResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PeeringServicePrefixEventResponse)(nil)).Elem()
}

func (o PeeringServicePrefixEventResponseOutput) ToPeeringServicePrefixEventResponseOutput() PeeringServicePrefixEventResponseOutput {
	return o
}

func (o PeeringServicePrefixEventResponseOutput) ToPeeringServicePrefixEventResponseOutputWithContext(ctx context.Context) PeeringServicePrefixEventResponseOutput {
	return o
}

func (o PeeringServicePrefixEventResponseOutput) EventDescription() pulumi.StringOutput {
	return o.ApplyT(func(v PeeringServicePrefixEventResponse) string { return v.EventDescription }).(pulumi.StringOutput)
}

func (o PeeringServicePrefixEventResponseOutput) EventLevel() pulumi.StringOutput {
	return o.ApplyT(func(v PeeringServicePrefixEventResponse) string { return v.EventLevel }).(pulumi.StringOutput)
}

func (o PeeringServicePrefixEventResponseOutput) EventSummary() pulumi.StringOutput {
	return o.ApplyT(func(v PeeringServicePrefixEventResponse) string { return v.EventSummary }).(pulumi.StringOutput)
}

func (o PeeringServicePrefixEventResponseOutput) EventTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v PeeringServicePrefixEventResponse) string { return v.EventTimestamp }).(pulumi.StringOutput)
}

func (o PeeringServicePrefixEventResponseOutput) EventType() pulumi.StringOutput {
	return o.ApplyT(func(v PeeringServicePrefixEventResponse) string { return v.EventType }).(pulumi.StringOutput)
}

type PeeringServicePrefixEventResponseArrayOutput struct{ *pulumi.OutputState }

func (PeeringServicePrefixEventResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PeeringServicePrefixEventResponse)(nil)).Elem()
}

func (o PeeringServicePrefixEventResponseArrayOutput) ToPeeringServicePrefixEventResponseArrayOutput() PeeringServicePrefixEventResponseArrayOutput {
	return o
}

func (o PeeringServicePrefixEventResponseArrayOutput) ToPeeringServicePrefixEventResponseArrayOutputWithContext(ctx context.Context) PeeringServicePrefixEventResponseArrayOutput {
	return o
}

func (o PeeringServicePrefixEventResponseArrayOutput) Index(i pulumi.IntInput) PeeringServicePrefixEventResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PeeringServicePrefixEventResponse {
		return vs[0].([]PeeringServicePrefixEventResponse)[vs[1].(int)]
	}).(PeeringServicePrefixEventResponseOutput)
}

type PeeringServiceSku struct {
	Name *string `pulumi:"name"`
}





type PeeringServiceSkuInput interface {
	pulumi.Input

	ToPeeringServiceSkuOutput() PeeringServiceSkuOutput
	ToPeeringServiceSkuOutputWithContext(context.Context) PeeringServiceSkuOutput
}

type PeeringServiceSkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (PeeringServiceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PeeringServiceSku)(nil)).Elem()
}

func (i PeeringServiceSkuArgs) ToPeeringServiceSkuOutput() PeeringServiceSkuOutput {
	return i.ToPeeringServiceSkuOutputWithContext(context.Background())
}

func (i PeeringServiceSkuArgs) ToPeeringServiceSkuOutputWithContext(ctx context.Context) PeeringServiceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringServiceSkuOutput)
}

func (i PeeringServiceSkuArgs) ToPeeringServiceSkuPtrOutput() PeeringServiceSkuPtrOutput {
	return i.ToPeeringServiceSkuPtrOutputWithContext(context.Background())
}

func (i PeeringServiceSkuArgs) ToPeeringServiceSkuPtrOutputWithContext(ctx context.Context) PeeringServiceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringServiceSkuOutput).ToPeeringServiceSkuPtrOutputWithContext(ctx)
}









type PeeringServiceSkuPtrInput interface {
	pulumi.Input

	ToPeeringServiceSkuPtrOutput() PeeringServiceSkuPtrOutput
	ToPeeringServiceSkuPtrOutputWithContext(context.Context) PeeringServiceSkuPtrOutput
}

type peeringServiceSkuPtrType PeeringServiceSkuArgs

func PeeringServiceSkuPtr(v *PeeringServiceSkuArgs) PeeringServiceSkuPtrInput {
	return (*peeringServiceSkuPtrType)(v)
}

func (*peeringServiceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringServiceSku)(nil)).Elem()
}

func (i *peeringServiceSkuPtrType) ToPeeringServiceSkuPtrOutput() PeeringServiceSkuPtrOutput {
	return i.ToPeeringServiceSkuPtrOutputWithContext(context.Background())
}

func (i *peeringServiceSkuPtrType) ToPeeringServiceSkuPtrOutputWithContext(ctx context.Context) PeeringServiceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringServiceSkuPtrOutput)
}

type PeeringServiceSkuOutput struct{ *pulumi.OutputState }

func (PeeringServiceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PeeringServiceSku)(nil)).Elem()
}

func (o PeeringServiceSkuOutput) ToPeeringServiceSkuOutput() PeeringServiceSkuOutput {
	return o
}

func (o PeeringServiceSkuOutput) ToPeeringServiceSkuOutputWithContext(ctx context.Context) PeeringServiceSkuOutput {
	return o
}

func (o PeeringServiceSkuOutput) ToPeeringServiceSkuPtrOutput() PeeringServiceSkuPtrOutput {
	return o.ToPeeringServiceSkuPtrOutputWithContext(context.Background())
}

func (o PeeringServiceSkuOutput) ToPeeringServiceSkuPtrOutputWithContext(ctx context.Context) PeeringServiceSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PeeringServiceSku) *PeeringServiceSku {
		return &v
	}).(PeeringServiceSkuPtrOutput)
}

func (o PeeringServiceSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PeeringServiceSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type PeeringServiceSkuPtrOutput struct{ *pulumi.OutputState }

func (PeeringServiceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringServiceSku)(nil)).Elem()
}

func (o PeeringServiceSkuPtrOutput) ToPeeringServiceSkuPtrOutput() PeeringServiceSkuPtrOutput {
	return o
}

func (o PeeringServiceSkuPtrOutput) ToPeeringServiceSkuPtrOutputWithContext(ctx context.Context) PeeringServiceSkuPtrOutput {
	return o
}

func (o PeeringServiceSkuPtrOutput) Elem() PeeringServiceSkuOutput {
	return o.ApplyT(func(v *PeeringServiceSku) PeeringServiceSku {
		if v != nil {
			return *v
		}
		var ret PeeringServiceSku
		return ret
	}).(PeeringServiceSkuOutput)
}

func (o PeeringServiceSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeeringServiceSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type PeeringServiceSkuResponse struct {
	Name *string `pulumi:"name"`
}





type PeeringServiceSkuResponseInput interface {
	pulumi.Input

	ToPeeringServiceSkuResponseOutput() PeeringServiceSkuResponseOutput
	ToPeeringServiceSkuResponseOutputWithContext(context.Context) PeeringServiceSkuResponseOutput
}

type PeeringServiceSkuResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (PeeringServiceSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PeeringServiceSkuResponse)(nil)).Elem()
}

func (i PeeringServiceSkuResponseArgs) ToPeeringServiceSkuResponseOutput() PeeringServiceSkuResponseOutput {
	return i.ToPeeringServiceSkuResponseOutputWithContext(context.Background())
}

func (i PeeringServiceSkuResponseArgs) ToPeeringServiceSkuResponseOutputWithContext(ctx context.Context) PeeringServiceSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringServiceSkuResponseOutput)
}

func (i PeeringServiceSkuResponseArgs) ToPeeringServiceSkuResponsePtrOutput() PeeringServiceSkuResponsePtrOutput {
	return i.ToPeeringServiceSkuResponsePtrOutputWithContext(context.Background())
}

func (i PeeringServiceSkuResponseArgs) ToPeeringServiceSkuResponsePtrOutputWithContext(ctx context.Context) PeeringServiceSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringServiceSkuResponseOutput).ToPeeringServiceSkuResponsePtrOutputWithContext(ctx)
}









type PeeringServiceSkuResponsePtrInput interface {
	pulumi.Input

	ToPeeringServiceSkuResponsePtrOutput() PeeringServiceSkuResponsePtrOutput
	ToPeeringServiceSkuResponsePtrOutputWithContext(context.Context) PeeringServiceSkuResponsePtrOutput
}

type peeringServiceSkuResponsePtrType PeeringServiceSkuResponseArgs

func PeeringServiceSkuResponsePtr(v *PeeringServiceSkuResponseArgs) PeeringServiceSkuResponsePtrInput {
	return (*peeringServiceSkuResponsePtrType)(v)
}

func (*peeringServiceSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringServiceSkuResponse)(nil)).Elem()
}

func (i *peeringServiceSkuResponsePtrType) ToPeeringServiceSkuResponsePtrOutput() PeeringServiceSkuResponsePtrOutput {
	return i.ToPeeringServiceSkuResponsePtrOutputWithContext(context.Background())
}

func (i *peeringServiceSkuResponsePtrType) ToPeeringServiceSkuResponsePtrOutputWithContext(ctx context.Context) PeeringServiceSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringServiceSkuResponsePtrOutput)
}

type PeeringServiceSkuResponseOutput struct{ *pulumi.OutputState }

func (PeeringServiceSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PeeringServiceSkuResponse)(nil)).Elem()
}

func (o PeeringServiceSkuResponseOutput) ToPeeringServiceSkuResponseOutput() PeeringServiceSkuResponseOutput {
	return o
}

func (o PeeringServiceSkuResponseOutput) ToPeeringServiceSkuResponseOutputWithContext(ctx context.Context) PeeringServiceSkuResponseOutput {
	return o
}

func (o PeeringServiceSkuResponseOutput) ToPeeringServiceSkuResponsePtrOutput() PeeringServiceSkuResponsePtrOutput {
	return o.ToPeeringServiceSkuResponsePtrOutputWithContext(context.Background())
}

func (o PeeringServiceSkuResponseOutput) ToPeeringServiceSkuResponsePtrOutputWithContext(ctx context.Context) PeeringServiceSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PeeringServiceSkuResponse) *PeeringServiceSkuResponse {
		return &v
	}).(PeeringServiceSkuResponsePtrOutput)
}

func (o PeeringServiceSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PeeringServiceSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type PeeringServiceSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (PeeringServiceSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringServiceSkuResponse)(nil)).Elem()
}

func (o PeeringServiceSkuResponsePtrOutput) ToPeeringServiceSkuResponsePtrOutput() PeeringServiceSkuResponsePtrOutput {
	return o
}

func (o PeeringServiceSkuResponsePtrOutput) ToPeeringServiceSkuResponsePtrOutputWithContext(ctx context.Context) PeeringServiceSkuResponsePtrOutput {
	return o
}

func (o PeeringServiceSkuResponsePtrOutput) Elem() PeeringServiceSkuResponseOutput {
	return o.ApplyT(func(v *PeeringServiceSkuResponse) PeeringServiceSkuResponse {
		if v != nil {
			return *v
		}
		var ret PeeringServiceSkuResponse
		return ret
	}).(PeeringServiceSkuResponseOutput)
}

func (o PeeringServiceSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeeringServiceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type PeeringSku struct {
	Family *string `pulumi:"family"`
	Name   *string `pulumi:"name"`
	Size   *string `pulumi:"size"`
	Tier   *string `pulumi:"tier"`
}





type PeeringSkuInput interface {
	pulumi.Input

	ToPeeringSkuOutput() PeeringSkuOutput
	ToPeeringSkuOutputWithContext(context.Context) PeeringSkuOutput
}

type PeeringSkuArgs struct {
	Family pulumi.StringPtrInput `pulumi:"family"`
	Name   pulumi.StringPtrInput `pulumi:"name"`
	Size   pulumi.StringPtrInput `pulumi:"size"`
	Tier   pulumi.StringPtrInput `pulumi:"tier"`
}

func (PeeringSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PeeringSku)(nil)).Elem()
}

func (i PeeringSkuArgs) ToPeeringSkuOutput() PeeringSkuOutput {
	return i.ToPeeringSkuOutputWithContext(context.Background())
}

func (i PeeringSkuArgs) ToPeeringSkuOutputWithContext(ctx context.Context) PeeringSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringSkuOutput)
}

func (i PeeringSkuArgs) ToPeeringSkuPtrOutput() PeeringSkuPtrOutput {
	return i.ToPeeringSkuPtrOutputWithContext(context.Background())
}

func (i PeeringSkuArgs) ToPeeringSkuPtrOutputWithContext(ctx context.Context) PeeringSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringSkuOutput).ToPeeringSkuPtrOutputWithContext(ctx)
}









type PeeringSkuPtrInput interface {
	pulumi.Input

	ToPeeringSkuPtrOutput() PeeringSkuPtrOutput
	ToPeeringSkuPtrOutputWithContext(context.Context) PeeringSkuPtrOutput
}

type peeringSkuPtrType PeeringSkuArgs

func PeeringSkuPtr(v *PeeringSkuArgs) PeeringSkuPtrInput {
	return (*peeringSkuPtrType)(v)
}

func (*peeringSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringSku)(nil)).Elem()
}

func (i *peeringSkuPtrType) ToPeeringSkuPtrOutput() PeeringSkuPtrOutput {
	return i.ToPeeringSkuPtrOutputWithContext(context.Background())
}

func (i *peeringSkuPtrType) ToPeeringSkuPtrOutputWithContext(ctx context.Context) PeeringSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringSkuPtrOutput)
}

type PeeringSkuOutput struct{ *pulumi.OutputState }

func (PeeringSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PeeringSku)(nil)).Elem()
}

func (o PeeringSkuOutput) ToPeeringSkuOutput() PeeringSkuOutput {
	return o
}

func (o PeeringSkuOutput) ToPeeringSkuOutputWithContext(ctx context.Context) PeeringSkuOutput {
	return o
}

func (o PeeringSkuOutput) ToPeeringSkuPtrOutput() PeeringSkuPtrOutput {
	return o.ToPeeringSkuPtrOutputWithContext(context.Background())
}

func (o PeeringSkuOutput) ToPeeringSkuPtrOutputWithContext(ctx context.Context) PeeringSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PeeringSku) *PeeringSku {
		return &v
	}).(PeeringSkuPtrOutput)
}

func (o PeeringSkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PeeringSku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o PeeringSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PeeringSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PeeringSkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PeeringSku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o PeeringSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PeeringSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type PeeringSkuPtrOutput struct{ *pulumi.OutputState }

func (PeeringSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringSku)(nil)).Elem()
}

func (o PeeringSkuPtrOutput) ToPeeringSkuPtrOutput() PeeringSkuPtrOutput {
	return o
}

func (o PeeringSkuPtrOutput) ToPeeringSkuPtrOutputWithContext(ctx context.Context) PeeringSkuPtrOutput {
	return o
}

func (o PeeringSkuPtrOutput) Elem() PeeringSkuOutput {
	return o.ApplyT(func(v *PeeringSku) PeeringSku {
		if v != nil {
			return *v
		}
		var ret PeeringSku
		return ret
	}).(PeeringSkuOutput)
}

func (o PeeringSkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeeringSku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o PeeringSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeeringSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PeeringSkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeeringSku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o PeeringSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeeringSku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type PeeringSkuResponse struct {
	Family *string `pulumi:"family"`
	Name   *string `pulumi:"name"`
	Size   *string `pulumi:"size"`
	Tier   *string `pulumi:"tier"`
}





type PeeringSkuResponseInput interface {
	pulumi.Input

	ToPeeringSkuResponseOutput() PeeringSkuResponseOutput
	ToPeeringSkuResponseOutputWithContext(context.Context) PeeringSkuResponseOutput
}

type PeeringSkuResponseArgs struct {
	Family pulumi.StringPtrInput `pulumi:"family"`
	Name   pulumi.StringPtrInput `pulumi:"name"`
	Size   pulumi.StringPtrInput `pulumi:"size"`
	Tier   pulumi.StringPtrInput `pulumi:"tier"`
}

func (PeeringSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PeeringSkuResponse)(nil)).Elem()
}

func (i PeeringSkuResponseArgs) ToPeeringSkuResponseOutput() PeeringSkuResponseOutput {
	return i.ToPeeringSkuResponseOutputWithContext(context.Background())
}

func (i PeeringSkuResponseArgs) ToPeeringSkuResponseOutputWithContext(ctx context.Context) PeeringSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringSkuResponseOutput)
}

func (i PeeringSkuResponseArgs) ToPeeringSkuResponsePtrOutput() PeeringSkuResponsePtrOutput {
	return i.ToPeeringSkuResponsePtrOutputWithContext(context.Background())
}

func (i PeeringSkuResponseArgs) ToPeeringSkuResponsePtrOutputWithContext(ctx context.Context) PeeringSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringSkuResponseOutput).ToPeeringSkuResponsePtrOutputWithContext(ctx)
}









type PeeringSkuResponsePtrInput interface {
	pulumi.Input

	ToPeeringSkuResponsePtrOutput() PeeringSkuResponsePtrOutput
	ToPeeringSkuResponsePtrOutputWithContext(context.Context) PeeringSkuResponsePtrOutput
}

type peeringSkuResponsePtrType PeeringSkuResponseArgs

func PeeringSkuResponsePtr(v *PeeringSkuResponseArgs) PeeringSkuResponsePtrInput {
	return (*peeringSkuResponsePtrType)(v)
}

func (*peeringSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringSkuResponse)(nil)).Elem()
}

func (i *peeringSkuResponsePtrType) ToPeeringSkuResponsePtrOutput() PeeringSkuResponsePtrOutput {
	return i.ToPeeringSkuResponsePtrOutputWithContext(context.Background())
}

func (i *peeringSkuResponsePtrType) ToPeeringSkuResponsePtrOutputWithContext(ctx context.Context) PeeringSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringSkuResponsePtrOutput)
}

type PeeringSkuResponseOutput struct{ *pulumi.OutputState }

func (PeeringSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PeeringSkuResponse)(nil)).Elem()
}

func (o PeeringSkuResponseOutput) ToPeeringSkuResponseOutput() PeeringSkuResponseOutput {
	return o
}

func (o PeeringSkuResponseOutput) ToPeeringSkuResponseOutputWithContext(ctx context.Context) PeeringSkuResponseOutput {
	return o
}

func (o PeeringSkuResponseOutput) ToPeeringSkuResponsePtrOutput() PeeringSkuResponsePtrOutput {
	return o.ToPeeringSkuResponsePtrOutputWithContext(context.Background())
}

func (o PeeringSkuResponseOutput) ToPeeringSkuResponsePtrOutputWithContext(ctx context.Context) PeeringSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PeeringSkuResponse) *PeeringSkuResponse {
		return &v
	}).(PeeringSkuResponsePtrOutput)
}

func (o PeeringSkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PeeringSkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o PeeringSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PeeringSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PeeringSkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PeeringSkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o PeeringSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PeeringSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type PeeringSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (PeeringSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringSkuResponse)(nil)).Elem()
}

func (o PeeringSkuResponsePtrOutput) ToPeeringSkuResponsePtrOutput() PeeringSkuResponsePtrOutput {
	return o
}

func (o PeeringSkuResponsePtrOutput) ToPeeringSkuResponsePtrOutputWithContext(ctx context.Context) PeeringSkuResponsePtrOutput {
	return o
}

func (o PeeringSkuResponsePtrOutput) Elem() PeeringSkuResponseOutput {
	return o.ApplyT(func(v *PeeringSkuResponse) PeeringSkuResponse {
		if v != nil {
			return *v
		}
		var ret PeeringSkuResponse
		return ret
	}).(PeeringSkuResponseOutput)
}

func (o PeeringSkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeeringSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o PeeringSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeeringSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PeeringSkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeeringSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o PeeringSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeeringSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SubResource struct {
	Id *string `pulumi:"id"`
}





type SubResourceInput interface {
	pulumi.Input

	ToSubResourceOutput() SubResourceOutput
	ToSubResourceOutputWithContext(context.Context) SubResourceOutput
}

type SubResourceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (i SubResourceArgs) ToSubResourceOutput() SubResourceOutput {
	return i.ToSubResourceOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput)
}

func (i SubResourceArgs) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput).ToSubResourcePtrOutputWithContext(ctx)
}









type SubResourcePtrInput interface {
	pulumi.Input

	ToSubResourcePtrOutput() SubResourcePtrOutput
	ToSubResourcePtrOutputWithContext(context.Context) SubResourcePtrOutput
}

type subResourcePtrType SubResourceArgs

func SubResourcePtr(v *SubResourceArgs) SubResourcePtrInput {
	return (*subResourcePtrType)(v)
}

func (*subResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (i *subResourcePtrType) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i *subResourcePtrType) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourcePtrOutput)
}

type SubResourceOutput struct{ *pulumi.OutputState }

func (SubResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (o SubResourceOutput) ToSubResourceOutput() SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o.ToSubResourcePtrOutputWithContext(context.Background())
}

func (o SubResourceOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubResource) *SubResource {
		return &v
	}).(SubResourcePtrOutput)
}

func (o SubResourceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResource) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourcePtrOutput struct{ *pulumi.OutputState }

func (SubResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) Elem() SubResourceOutput {
	return o.ApplyT(func(v *SubResource) SubResource {
		if v != nil {
			return *v
		}
		var ret SubResource
		return ret
	}).(SubResourceOutput)
}

func (o SubResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResource) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceResponse struct {
	Id *string `pulumi:"id"`
}





type SubResourceResponseInput interface {
	pulumi.Input

	ToSubResourceResponseOutput() SubResourceResponseOutput
	ToSubResourceResponseOutputWithContext(context.Context) SubResourceResponseOutput
}

type SubResourceResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (i SubResourceResponseArgs) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return i.ToSubResourceResponseOutputWithContext(context.Background())
}

func (i SubResourceResponseArgs) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceResponseOutput)
}

func (i SubResourceResponseArgs) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return i.ToSubResourceResponsePtrOutputWithContext(context.Background())
}

func (i SubResourceResponseArgs) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceResponseOutput).ToSubResourceResponsePtrOutputWithContext(ctx)
}









type SubResourceResponsePtrInput interface {
	pulumi.Input

	ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput
	ToSubResourceResponsePtrOutputWithContext(context.Context) SubResourceResponsePtrOutput
}

type subResourceResponsePtrType SubResourceResponseArgs

func SubResourceResponsePtr(v *SubResourceResponseArgs) SubResourceResponsePtrInput {
	return (*subResourceResponsePtrType)(v)
}

func (*subResourceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResourceResponse)(nil)).Elem()
}

func (i *subResourceResponsePtrType) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return i.ToSubResourceResponsePtrOutputWithContext(context.Background())
}

func (i *subResourceResponsePtrType) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceResponsePtrOutput)
}

type SubResourceResponseOutput struct{ *pulumi.OutputState }

func (SubResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return o.ToSubResourceResponsePtrOutputWithContext(context.Background())
}

func (o SubResourceResponseOutput) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubResourceResponse) *SubResourceResponse {
		return &v
	}).(SubResourceResponsePtrOutput)
}

func (o SubResourceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResourceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourceResponsePtrOutput struct{ *pulumi.OutputState }

func (SubResourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) Elem() SubResourceResponseOutput {
	return o.ApplyT(func(v *SubResourceResponse) SubResourceResponse {
		if v != nil {
			return *v
		}
		var ret SubResourceResponse
		return ret
	}).(SubResourceResponseOutput)
}

func (o SubResourceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BgpSessionInput)(nil)).Elem(), BgpSessionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpSessionPtrInput)(nil)).Elem(), BgpSessionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpSessionResponseInput)(nil)).Elem(), BgpSessionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpSessionResponsePtrInput)(nil)).Elem(), BgpSessionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContactDetailInput)(nil)).Elem(), ContactDetailArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContactDetailArrayInput)(nil)).Elem(), ContactDetailArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContactDetailResponseInput)(nil)).Elem(), ContactDetailResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContactDetailResponseArrayInput)(nil)).Elem(), ContactDetailResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectConnectionInput)(nil)).Elem(), DirectConnectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectConnectionArrayInput)(nil)).Elem(), DirectConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectConnectionResponseInput)(nil)).Elem(), DirectConnectionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectConnectionResponseArrayInput)(nil)).Elem(), DirectConnectionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExchangeConnectionInput)(nil)).Elem(), ExchangeConnectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExchangeConnectionArrayInput)(nil)).Elem(), ExchangeConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExchangeConnectionResponseInput)(nil)).Elem(), ExchangeConnectionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExchangeConnectionResponseArrayInput)(nil)).Elem(), ExchangeConnectionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringPropertiesDirectInput)(nil)).Elem(), PeeringPropertiesDirectArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringPropertiesDirectPtrInput)(nil)).Elem(), PeeringPropertiesDirectArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringPropertiesDirectResponseInput)(nil)).Elem(), PeeringPropertiesDirectResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringPropertiesDirectResponsePtrInput)(nil)).Elem(), PeeringPropertiesDirectResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringPropertiesExchangeInput)(nil)).Elem(), PeeringPropertiesExchangeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringPropertiesExchangePtrInput)(nil)).Elem(), PeeringPropertiesExchangeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringPropertiesExchangeResponseInput)(nil)).Elem(), PeeringPropertiesExchangeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringPropertiesExchangeResponsePtrInput)(nil)).Elem(), PeeringPropertiesExchangeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringServicePrefixEventResponseInput)(nil)).Elem(), PeeringServicePrefixEventResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringServicePrefixEventResponseArrayInput)(nil)).Elem(), PeeringServicePrefixEventResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringServiceSkuInput)(nil)).Elem(), PeeringServiceSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringServiceSkuPtrInput)(nil)).Elem(), PeeringServiceSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringServiceSkuResponseInput)(nil)).Elem(), PeeringServiceSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringServiceSkuResponsePtrInput)(nil)).Elem(), PeeringServiceSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringSkuInput)(nil)).Elem(), PeeringSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringSkuPtrInput)(nil)).Elem(), PeeringSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringSkuResponseInput)(nil)).Elem(), PeeringSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringSkuResponsePtrInput)(nil)).Elem(), PeeringSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubResourceInput)(nil)).Elem(), SubResourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubResourcePtrInput)(nil)).Elem(), SubResourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubResourceResponseInput)(nil)).Elem(), SubResourceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubResourceResponsePtrInput)(nil)).Elem(), SubResourceResponseArgs{})
	pulumi.RegisterOutputType(BgpSessionOutput{})
	pulumi.RegisterOutputType(BgpSessionPtrOutput{})
	pulumi.RegisterOutputType(BgpSessionResponseOutput{})
	pulumi.RegisterOutputType(BgpSessionResponsePtrOutput{})
	pulumi.RegisterOutputType(ContactDetailOutput{})
	pulumi.RegisterOutputType(ContactDetailArrayOutput{})
	pulumi.RegisterOutputType(ContactDetailResponseOutput{})
	pulumi.RegisterOutputType(ContactDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(DirectConnectionOutput{})
	pulumi.RegisterOutputType(DirectConnectionArrayOutput{})
	pulumi.RegisterOutputType(DirectConnectionResponseOutput{})
	pulumi.RegisterOutputType(DirectConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(ExchangeConnectionOutput{})
	pulumi.RegisterOutputType(ExchangeConnectionArrayOutput{})
	pulumi.RegisterOutputType(ExchangeConnectionResponseOutput{})
	pulumi.RegisterOutputType(ExchangeConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PeeringPropertiesDirectOutput{})
	pulumi.RegisterOutputType(PeeringPropertiesDirectPtrOutput{})
	pulumi.RegisterOutputType(PeeringPropertiesDirectResponseOutput{})
	pulumi.RegisterOutputType(PeeringPropertiesDirectResponsePtrOutput{})
	pulumi.RegisterOutputType(PeeringPropertiesExchangeOutput{})
	pulumi.RegisterOutputType(PeeringPropertiesExchangePtrOutput{})
	pulumi.RegisterOutputType(PeeringPropertiesExchangeResponseOutput{})
	pulumi.RegisterOutputType(PeeringPropertiesExchangeResponsePtrOutput{})
	pulumi.RegisterOutputType(PeeringServicePrefixEventResponseOutput{})
	pulumi.RegisterOutputType(PeeringServicePrefixEventResponseArrayOutput{})
	pulumi.RegisterOutputType(PeeringServiceSkuOutput{})
	pulumi.RegisterOutputType(PeeringServiceSkuPtrOutput{})
	pulumi.RegisterOutputType(PeeringServiceSkuResponseOutput{})
	pulumi.RegisterOutputType(PeeringServiceSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(PeeringSkuOutput{})
	pulumi.RegisterOutputType(PeeringSkuPtrOutput{})
	pulumi.RegisterOutputType(PeeringSkuResponseOutput{})
	pulumi.RegisterOutputType(PeeringSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourcePtrOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SubResourceResponsePtrOutput{})
}

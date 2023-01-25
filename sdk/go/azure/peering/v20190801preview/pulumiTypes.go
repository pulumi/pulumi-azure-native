


package v20190801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BgpSession struct {
	MaxPrefixesAdvertisedV4 *int    `pulumi:"maxPrefixesAdvertisedV4"`
	MaxPrefixesAdvertisedV6 *int    `pulumi:"maxPrefixesAdvertisedV6"`
	Md5AuthenticationKey    *string `pulumi:"md5AuthenticationKey"`
	PeerSessionIPv4Address  *string `pulumi:"peerSessionIPv4Address"`
	PeerSessionIPv6Address  *string `pulumi:"peerSessionIPv6Address"`
	SessionPrefixV4         *string `pulumi:"sessionPrefixV4"`
	SessionPrefixV6         *string `pulumi:"sessionPrefixV6"`
}





type BgpSessionInput interface {
	pulumi.Input

	ToBgpSessionOutput() BgpSessionOutput
	ToBgpSessionOutputWithContext(context.Context) BgpSessionOutput
}

type BgpSessionArgs struct {
	MaxPrefixesAdvertisedV4 pulumi.IntPtrInput    `pulumi:"maxPrefixesAdvertisedV4"`
	MaxPrefixesAdvertisedV6 pulumi.IntPtrInput    `pulumi:"maxPrefixesAdvertisedV6"`
	Md5AuthenticationKey    pulumi.StringPtrInput `pulumi:"md5AuthenticationKey"`
	PeerSessionIPv4Address  pulumi.StringPtrInput `pulumi:"peerSessionIPv4Address"`
	PeerSessionIPv6Address  pulumi.StringPtrInput `pulumi:"peerSessionIPv6Address"`
	SessionPrefixV4         pulumi.StringPtrInput `pulumi:"sessionPrefixV4"`
	SessionPrefixV6         pulumi.StringPtrInput `pulumi:"sessionPrefixV6"`
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
	MicrosoftSessionIPv4Address string  `pulumi:"microsoftSessionIPv4Address"`
	MicrosoftSessionIPv6Address string  `pulumi:"microsoftSessionIPv6Address"`
	PeerSessionIPv4Address      *string `pulumi:"peerSessionIPv4Address"`
	PeerSessionIPv6Address      *string `pulumi:"peerSessionIPv6Address"`
	SessionPrefixV4             *string `pulumi:"sessionPrefixV4"`
	SessionPrefixV6             *string `pulumi:"sessionPrefixV6"`
	SessionStateV4              string  `pulumi:"sessionStateV4"`
	SessionStateV6              string  `pulumi:"sessionStateV6"`
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

func (o BgpSessionResponseOutput) MaxPrefixesAdvertisedV4() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BgpSessionResponse) *int { return v.MaxPrefixesAdvertisedV4 }).(pulumi.IntPtrOutput)
}

func (o BgpSessionResponseOutput) MaxPrefixesAdvertisedV6() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BgpSessionResponse) *int { return v.MaxPrefixesAdvertisedV6 }).(pulumi.IntPtrOutput)
}

func (o BgpSessionResponseOutput) Md5AuthenticationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpSessionResponse) *string { return v.Md5AuthenticationKey }).(pulumi.StringPtrOutput)
}

func (o BgpSessionResponseOutput) MicrosoftSessionIPv4Address() pulumi.StringOutput {
	return o.ApplyT(func(v BgpSessionResponse) string { return v.MicrosoftSessionIPv4Address }).(pulumi.StringOutput)
}

func (o BgpSessionResponseOutput) MicrosoftSessionIPv6Address() pulumi.StringOutput {
	return o.ApplyT(func(v BgpSessionResponse) string { return v.MicrosoftSessionIPv6Address }).(pulumi.StringOutput)
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
		return &v.MicrosoftSessionIPv4Address
	}).(pulumi.StringPtrOutput)
}

func (o BgpSessionResponsePtrOutput) MicrosoftSessionIPv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSessionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.MicrosoftSessionIPv6Address
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

type ContactInfo struct {
	Emails []string `pulumi:"emails"`
	Phone  []string `pulumi:"phone"`
}





type ContactInfoInput interface {
	pulumi.Input

	ToContactInfoOutput() ContactInfoOutput
	ToContactInfoOutputWithContext(context.Context) ContactInfoOutput
}

type ContactInfoArgs struct {
	Emails pulumi.StringArrayInput `pulumi:"emails"`
	Phone  pulumi.StringArrayInput `pulumi:"phone"`
}

func (ContactInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactInfo)(nil)).Elem()
}

func (i ContactInfoArgs) ToContactInfoOutput() ContactInfoOutput {
	return i.ToContactInfoOutputWithContext(context.Background())
}

func (i ContactInfoArgs) ToContactInfoOutputWithContext(ctx context.Context) ContactInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactInfoOutput)
}

func (i ContactInfoArgs) ToContactInfoPtrOutput() ContactInfoPtrOutput {
	return i.ToContactInfoPtrOutputWithContext(context.Background())
}

func (i ContactInfoArgs) ToContactInfoPtrOutputWithContext(ctx context.Context) ContactInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactInfoOutput).ToContactInfoPtrOutputWithContext(ctx)
}









type ContactInfoPtrInput interface {
	pulumi.Input

	ToContactInfoPtrOutput() ContactInfoPtrOutput
	ToContactInfoPtrOutputWithContext(context.Context) ContactInfoPtrOutput
}

type contactInfoPtrType ContactInfoArgs

func ContactInfoPtr(v *ContactInfoArgs) ContactInfoPtrInput {
	return (*contactInfoPtrType)(v)
}

func (*contactInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactInfo)(nil)).Elem()
}

func (i *contactInfoPtrType) ToContactInfoPtrOutput() ContactInfoPtrOutput {
	return i.ToContactInfoPtrOutputWithContext(context.Background())
}

func (i *contactInfoPtrType) ToContactInfoPtrOutputWithContext(ctx context.Context) ContactInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactInfoPtrOutput)
}

type ContactInfoOutput struct{ *pulumi.OutputState }

func (ContactInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactInfo)(nil)).Elem()
}

func (o ContactInfoOutput) ToContactInfoOutput() ContactInfoOutput {
	return o
}

func (o ContactInfoOutput) ToContactInfoOutputWithContext(ctx context.Context) ContactInfoOutput {
	return o
}

func (o ContactInfoOutput) ToContactInfoPtrOutput() ContactInfoPtrOutput {
	return o.ToContactInfoPtrOutputWithContext(context.Background())
}

func (o ContactInfoOutput) ToContactInfoPtrOutputWithContext(ctx context.Context) ContactInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContactInfo) *ContactInfo {
		return &v
	}).(ContactInfoPtrOutput)
}

func (o ContactInfoOutput) Emails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContactInfo) []string { return v.Emails }).(pulumi.StringArrayOutput)
}

func (o ContactInfoOutput) Phone() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContactInfo) []string { return v.Phone }).(pulumi.StringArrayOutput)
}

type ContactInfoPtrOutput struct{ *pulumi.OutputState }

func (ContactInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactInfo)(nil)).Elem()
}

func (o ContactInfoPtrOutput) ToContactInfoPtrOutput() ContactInfoPtrOutput {
	return o
}

func (o ContactInfoPtrOutput) ToContactInfoPtrOutputWithContext(ctx context.Context) ContactInfoPtrOutput {
	return o
}

func (o ContactInfoPtrOutput) Elem() ContactInfoOutput {
	return o.ApplyT(func(v *ContactInfo) ContactInfo {
		if v != nil {
			return *v
		}
		var ret ContactInfo
		return ret
	}).(ContactInfoOutput)
}

func (o ContactInfoPtrOutput) Emails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContactInfo) []string {
		if v == nil {
			return nil
		}
		return v.Emails
	}).(pulumi.StringArrayOutput)
}

func (o ContactInfoPtrOutput) Phone() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContactInfo) []string {
		if v == nil {
			return nil
		}
		return v.Phone
	}).(pulumi.StringArrayOutput)
}

type ContactInfoResponse struct {
	Emails []string `pulumi:"emails"`
	Phone  []string `pulumi:"phone"`
}

type ContactInfoResponseOutput struct{ *pulumi.OutputState }

func (ContactInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactInfoResponse)(nil)).Elem()
}

func (o ContactInfoResponseOutput) ToContactInfoResponseOutput() ContactInfoResponseOutput {
	return o
}

func (o ContactInfoResponseOutput) ToContactInfoResponseOutputWithContext(ctx context.Context) ContactInfoResponseOutput {
	return o
}

func (o ContactInfoResponseOutput) Emails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContactInfoResponse) []string { return v.Emails }).(pulumi.StringArrayOutput)
}

func (o ContactInfoResponseOutput) Phone() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContactInfoResponse) []string { return v.Phone }).(pulumi.StringArrayOutput)
}

type ContactInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ContactInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactInfoResponse)(nil)).Elem()
}

func (o ContactInfoResponsePtrOutput) ToContactInfoResponsePtrOutput() ContactInfoResponsePtrOutput {
	return o
}

func (o ContactInfoResponsePtrOutput) ToContactInfoResponsePtrOutputWithContext(ctx context.Context) ContactInfoResponsePtrOutput {
	return o
}

func (o ContactInfoResponsePtrOutput) Elem() ContactInfoResponseOutput {
	return o.ApplyT(func(v *ContactInfoResponse) ContactInfoResponse {
		if v != nil {
			return *v
		}
		var ret ContactInfoResponse
		return ret
	}).(ContactInfoResponseOutput)
}

func (o ContactInfoResponsePtrOutput) Emails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContactInfoResponse) []string {
		if v == nil {
			return nil
		}
		return v.Emails
	}).(pulumi.StringArrayOutput)
}

func (o ContactInfoResponsePtrOutput) Phone() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContactInfoResponse) []string {
		if v == nil {
			return nil
		}
		return v.Phone
	}).(pulumi.StringArrayOutput)
}

type DirectConnection struct {
	BandwidthInMbps            *int        `pulumi:"bandwidthInMbps"`
	BgpSession                 *BgpSession `pulumi:"bgpSession"`
	ConnectionIdentifier       *string     `pulumi:"connectionIdentifier"`
	PeeringDBFacilityId        *int        `pulumi:"peeringDBFacilityId"`
	ProvisionedBandwidthInMbps *int        `pulumi:"provisionedBandwidthInMbps"`
	SessionAddressProvider     *string     `pulumi:"sessionAddressProvider"`
	UseForPeeringService       *bool       `pulumi:"useForPeeringService"`
}





type DirectConnectionInput interface {
	pulumi.Input

	ToDirectConnectionOutput() DirectConnectionOutput
	ToDirectConnectionOutputWithContext(context.Context) DirectConnectionOutput
}

type DirectConnectionArgs struct {
	BandwidthInMbps            pulumi.IntPtrInput    `pulumi:"bandwidthInMbps"`
	BgpSession                 BgpSessionPtrInput    `pulumi:"bgpSession"`
	ConnectionIdentifier       pulumi.StringPtrInput `pulumi:"connectionIdentifier"`
	PeeringDBFacilityId        pulumi.IntPtrInput    `pulumi:"peeringDBFacilityId"`
	ProvisionedBandwidthInMbps pulumi.IntPtrInput    `pulumi:"provisionedBandwidthInMbps"`
	SessionAddressProvider     pulumi.StringPtrInput `pulumi:"sessionAddressProvider"`
	UseForPeeringService       pulumi.BoolPtrInput   `pulumi:"useForPeeringService"`
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

func (o DirectConnectionOutput) ProvisionedBandwidthInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DirectConnection) *int { return v.ProvisionedBandwidthInMbps }).(pulumi.IntPtrOutput)
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
	PeeringDBFacilityId        *int                `pulumi:"peeringDBFacilityId"`
	ProvisionedBandwidthInMbps *int                `pulumi:"provisionedBandwidthInMbps"`
	SessionAddressProvider     *string             `pulumi:"sessionAddressProvider"`
	UseForPeeringService       *bool               `pulumi:"useForPeeringService"`
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

func (o DirectConnectionResponseOutput) PeeringDBFacilityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DirectConnectionResponse) *int { return v.PeeringDBFacilityId }).(pulumi.IntPtrOutput)
}

func (o DirectConnectionResponseOutput) ProvisionedBandwidthInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DirectConnectionResponse) *int { return v.ProvisionedBandwidthInMbps }).(pulumi.IntPtrOutput)
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
	PeeringDBFacilityId  *int                `pulumi:"peeringDBFacilityId"`
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
	Connections          []DirectConnection `pulumi:"connections"`
	DirectPeeringType    *string            `pulumi:"directPeeringType"`
	PeerAsn              *SubResource       `pulumi:"peerAsn"`
	UseForPeeringService *bool              `pulumi:"useForPeeringService"`
}





type PeeringPropertiesDirectInput interface {
	pulumi.Input

	ToPeeringPropertiesDirectOutput() PeeringPropertiesDirectOutput
	ToPeeringPropertiesDirectOutputWithContext(context.Context) PeeringPropertiesDirectOutput
}

type PeeringPropertiesDirectArgs struct {
	Connections          DirectConnectionArrayInput `pulumi:"connections"`
	DirectPeeringType    pulumi.StringPtrInput      `pulumi:"directPeeringType"`
	PeerAsn              SubResourcePtrInput        `pulumi:"peerAsn"`
	UseForPeeringService pulumi.BoolPtrInput        `pulumi:"useForPeeringService"`
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

func (o PeeringPropertiesDirectOutput) UseForPeeringService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PeeringPropertiesDirect) *bool { return v.UseForPeeringService }).(pulumi.BoolPtrOutput)
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

func (o PeeringPropertiesDirectPtrOutput) UseForPeeringService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PeeringPropertiesDirect) *bool {
		if v == nil {
			return nil
		}
		return v.UseForPeeringService
	}).(pulumi.BoolPtrOutput)
}

type PeeringPropertiesDirectResponse struct {
	Connections          []DirectConnectionResponse `pulumi:"connections"`
	DirectPeeringType    *string                    `pulumi:"directPeeringType"`
	PeerAsn              *SubResourceResponse       `pulumi:"peerAsn"`
	UseForPeeringService *bool                      `pulumi:"useForPeeringService"`
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

func (o PeeringPropertiesDirectResponseOutput) Connections() DirectConnectionResponseArrayOutput {
	return o.ApplyT(func(v PeeringPropertiesDirectResponse) []DirectConnectionResponse { return v.Connections }).(DirectConnectionResponseArrayOutput)
}

func (o PeeringPropertiesDirectResponseOutput) DirectPeeringType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PeeringPropertiesDirectResponse) *string { return v.DirectPeeringType }).(pulumi.StringPtrOutput)
}

func (o PeeringPropertiesDirectResponseOutput) PeerAsn() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v PeeringPropertiesDirectResponse) *SubResourceResponse { return v.PeerAsn }).(SubResourceResponsePtrOutput)
}

func (o PeeringPropertiesDirectResponseOutput) UseForPeeringService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PeeringPropertiesDirectResponse) *bool { return v.UseForPeeringService }).(pulumi.BoolPtrOutput)
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
		return v.UseForPeeringService
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

type PeeringSkuResponse struct {
	Family *string `pulumi:"family"`
	Name   *string `pulumi:"name"`
	Size   *string `pulumi:"size"`
	Tier   *string `pulumi:"tier"`
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
	pulumi.RegisterOutputType(BgpSessionOutput{})
	pulumi.RegisterOutputType(BgpSessionPtrOutput{})
	pulumi.RegisterOutputType(BgpSessionResponseOutput{})
	pulumi.RegisterOutputType(BgpSessionResponsePtrOutput{})
	pulumi.RegisterOutputType(ContactInfoOutput{})
	pulumi.RegisterOutputType(ContactInfoPtrOutput{})
	pulumi.RegisterOutputType(ContactInfoResponseOutput{})
	pulumi.RegisterOutputType(ContactInfoResponsePtrOutput{})
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
	pulumi.RegisterOutputType(PeeringSkuOutput{})
	pulumi.RegisterOutputType(PeeringSkuResponseOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourcePtrOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SubResourceResponsePtrOutput{})
}

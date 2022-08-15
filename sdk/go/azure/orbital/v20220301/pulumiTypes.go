


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AuthorizedGroundstationResponse struct {
	ExpirationDate *string `pulumi:"expirationDate"`
	GroundStation  *string `pulumi:"groundStation"`
}

type AuthorizedGroundstationResponseOutput struct{ *pulumi.OutputState }

func (AuthorizedGroundstationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizedGroundstationResponse)(nil)).Elem()
}

func (o AuthorizedGroundstationResponseOutput) ToAuthorizedGroundstationResponseOutput() AuthorizedGroundstationResponseOutput {
	return o
}

func (o AuthorizedGroundstationResponseOutput) ToAuthorizedGroundstationResponseOutputWithContext(ctx context.Context) AuthorizedGroundstationResponseOutput {
	return o
}

func (o AuthorizedGroundstationResponseOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthorizedGroundstationResponse) *string { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

func (o AuthorizedGroundstationResponseOutput) GroundStation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthorizedGroundstationResponse) *string { return v.GroundStation }).(pulumi.StringPtrOutput)
}

type AuthorizedGroundstationResponseArrayOutput struct{ *pulumi.OutputState }

func (AuthorizedGroundstationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthorizedGroundstationResponse)(nil)).Elem()
}

func (o AuthorizedGroundstationResponseArrayOutput) ToAuthorizedGroundstationResponseArrayOutput() AuthorizedGroundstationResponseArrayOutput {
	return o
}

func (o AuthorizedGroundstationResponseArrayOutput) ToAuthorizedGroundstationResponseArrayOutputWithContext(ctx context.Context) AuthorizedGroundstationResponseArrayOutput {
	return o
}

func (o AuthorizedGroundstationResponseArrayOutput) Index(i pulumi.IntInput) AuthorizedGroundstationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AuthorizedGroundstationResponse {
		return vs[0].([]AuthorizedGroundstationResponse)[vs[1].(int)]
	}).(AuthorizedGroundstationResponseOutput)
}

type AvailableContactsResponse struct {
	EndAzimuthDegrees       float64                              `pulumi:"endAzimuthDegrees"`
	EndElevationDegrees     float64                              `pulumi:"endElevationDegrees"`
	GroundStationName       string                               `pulumi:"groundStationName"`
	MaximumElevationDegrees float64                              `pulumi:"maximumElevationDegrees"`
	RxEndTime               string                               `pulumi:"rxEndTime"`
	RxStartTime             string                               `pulumi:"rxStartTime"`
	Spacecraft              *AvailableContactsResponseSpacecraft `pulumi:"spacecraft"`
	StartAzimuthDegrees     float64                              `pulumi:"startAzimuthDegrees"`
	StartElevationDegrees   float64                              `pulumi:"startElevationDegrees"`
	TxEndTime               string                               `pulumi:"txEndTime"`
	TxStartTime             string                               `pulumi:"txStartTime"`
}

type AvailableContactsResponseOutput struct{ *pulumi.OutputState }

func (AvailableContactsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailableContactsResponse)(nil)).Elem()
}

func (o AvailableContactsResponseOutput) ToAvailableContactsResponseOutput() AvailableContactsResponseOutput {
	return o
}

func (o AvailableContactsResponseOutput) ToAvailableContactsResponseOutputWithContext(ctx context.Context) AvailableContactsResponseOutput {
	return o
}

func (o AvailableContactsResponseOutput) EndAzimuthDegrees() pulumi.Float64Output {
	return o.ApplyT(func(v AvailableContactsResponse) float64 { return v.EndAzimuthDegrees }).(pulumi.Float64Output)
}

func (o AvailableContactsResponseOutput) EndElevationDegrees() pulumi.Float64Output {
	return o.ApplyT(func(v AvailableContactsResponse) float64 { return v.EndElevationDegrees }).(pulumi.Float64Output)
}

func (o AvailableContactsResponseOutput) GroundStationName() pulumi.StringOutput {
	return o.ApplyT(func(v AvailableContactsResponse) string { return v.GroundStationName }).(pulumi.StringOutput)
}

func (o AvailableContactsResponseOutput) MaximumElevationDegrees() pulumi.Float64Output {
	return o.ApplyT(func(v AvailableContactsResponse) float64 { return v.MaximumElevationDegrees }).(pulumi.Float64Output)
}

func (o AvailableContactsResponseOutput) RxEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v AvailableContactsResponse) string { return v.RxEndTime }).(pulumi.StringOutput)
}

func (o AvailableContactsResponseOutput) RxStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v AvailableContactsResponse) string { return v.RxStartTime }).(pulumi.StringOutput)
}

func (o AvailableContactsResponseOutput) Spacecraft() AvailableContactsResponseSpacecraftPtrOutput {
	return o.ApplyT(func(v AvailableContactsResponse) *AvailableContactsResponseSpacecraft { return v.Spacecraft }).(AvailableContactsResponseSpacecraftPtrOutput)
}

func (o AvailableContactsResponseOutput) StartAzimuthDegrees() pulumi.Float64Output {
	return o.ApplyT(func(v AvailableContactsResponse) float64 { return v.StartAzimuthDegrees }).(pulumi.Float64Output)
}

func (o AvailableContactsResponseOutput) StartElevationDegrees() pulumi.Float64Output {
	return o.ApplyT(func(v AvailableContactsResponse) float64 { return v.StartElevationDegrees }).(pulumi.Float64Output)
}

func (o AvailableContactsResponseOutput) TxEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v AvailableContactsResponse) string { return v.TxEndTime }).(pulumi.StringOutput)
}

func (o AvailableContactsResponseOutput) TxStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v AvailableContactsResponse) string { return v.TxStartTime }).(pulumi.StringOutput)
}

type AvailableContactsResponseArrayOutput struct{ *pulumi.OutputState }

func (AvailableContactsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AvailableContactsResponse)(nil)).Elem()
}

func (o AvailableContactsResponseArrayOutput) ToAvailableContactsResponseArrayOutput() AvailableContactsResponseArrayOutput {
	return o
}

func (o AvailableContactsResponseArrayOutput) ToAvailableContactsResponseArrayOutputWithContext(ctx context.Context) AvailableContactsResponseArrayOutput {
	return o
}

func (o AvailableContactsResponseArrayOutput) Index(i pulumi.IntInput) AvailableContactsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AvailableContactsResponse {
		return vs[0].([]AvailableContactsResponse)[vs[1].(int)]
	}).(AvailableContactsResponseOutput)
}

type AvailableContactsResponseSpacecraft struct {
	Id *string `pulumi:"id"`
}

type AvailableContactsResponseSpacecraftOutput struct{ *pulumi.OutputState }

func (AvailableContactsResponseSpacecraftOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailableContactsResponseSpacecraft)(nil)).Elem()
}

func (o AvailableContactsResponseSpacecraftOutput) ToAvailableContactsResponseSpacecraftOutput() AvailableContactsResponseSpacecraftOutput {
	return o
}

func (o AvailableContactsResponseSpacecraftOutput) ToAvailableContactsResponseSpacecraftOutputWithContext(ctx context.Context) AvailableContactsResponseSpacecraftOutput {
	return o
}

func (o AvailableContactsResponseSpacecraftOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AvailableContactsResponseSpacecraft) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type AvailableContactsResponseSpacecraftPtrOutput struct{ *pulumi.OutputState }

func (AvailableContactsResponseSpacecraftPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AvailableContactsResponseSpacecraft)(nil)).Elem()
}

func (o AvailableContactsResponseSpacecraftPtrOutput) ToAvailableContactsResponseSpacecraftPtrOutput() AvailableContactsResponseSpacecraftPtrOutput {
	return o
}

func (o AvailableContactsResponseSpacecraftPtrOutput) ToAvailableContactsResponseSpacecraftPtrOutputWithContext(ctx context.Context) AvailableContactsResponseSpacecraftPtrOutput {
	return o
}

func (o AvailableContactsResponseSpacecraftPtrOutput) Elem() AvailableContactsResponseSpacecraftOutput {
	return o.ApplyT(func(v *AvailableContactsResponseSpacecraft) AvailableContactsResponseSpacecraft {
		if v != nil {
			return *v
		}
		var ret AvailableContactsResponseSpacecraft
		return ret
	}).(AvailableContactsResponseSpacecraftOutput)
}

func (o AvailableContactsResponseSpacecraftPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvailableContactsResponseSpacecraft) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type ContactParametersContactProfile struct {
	Id *string `pulumi:"id"`
}





type ContactParametersContactProfileInput interface {
	pulumi.Input

	ToContactParametersContactProfileOutput() ContactParametersContactProfileOutput
	ToContactParametersContactProfileOutputWithContext(context.Context) ContactParametersContactProfileOutput
}

type ContactParametersContactProfileArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ContactParametersContactProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactParametersContactProfile)(nil)).Elem()
}

func (i ContactParametersContactProfileArgs) ToContactParametersContactProfileOutput() ContactParametersContactProfileOutput {
	return i.ToContactParametersContactProfileOutputWithContext(context.Background())
}

func (i ContactParametersContactProfileArgs) ToContactParametersContactProfileOutputWithContext(ctx context.Context) ContactParametersContactProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactParametersContactProfileOutput)
}

type ContactParametersContactProfileOutput struct{ *pulumi.OutputState }

func (ContactParametersContactProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactParametersContactProfile)(nil)).Elem()
}

func (o ContactParametersContactProfileOutput) ToContactParametersContactProfileOutput() ContactParametersContactProfileOutput {
	return o
}

func (o ContactParametersContactProfileOutput) ToContactParametersContactProfileOutputWithContext(ctx context.Context) ContactParametersContactProfileOutput {
	return o
}

func (o ContactParametersContactProfileOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactParametersContactProfile) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ContactProfileLink struct {
	Channels            []ContactProfileLinkChannel `pulumi:"channels"`
	Direction           string                      `pulumi:"direction"`
	EirpdBW             *float64                    `pulumi:"eirpdBW"`
	GainOverTemperature *float64                    `pulumi:"gainOverTemperature"`
	Name                string                      `pulumi:"name"`
	Polarization        string                      `pulumi:"polarization"`
}





type ContactProfileLinkInput interface {
	pulumi.Input

	ToContactProfileLinkOutput() ContactProfileLinkOutput
	ToContactProfileLinkOutputWithContext(context.Context) ContactProfileLinkOutput
}

type ContactProfileLinkArgs struct {
	Channels            ContactProfileLinkChannelArrayInput `pulumi:"channels"`
	Direction           pulumi.StringInput                  `pulumi:"direction"`
	EirpdBW             pulumi.Float64PtrInput              `pulumi:"eirpdBW"`
	GainOverTemperature pulumi.Float64PtrInput              `pulumi:"gainOverTemperature"`
	Name                pulumi.StringInput                  `pulumi:"name"`
	Polarization        pulumi.StringInput                  `pulumi:"polarization"`
}

func (ContactProfileLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactProfileLink)(nil)).Elem()
}

func (i ContactProfileLinkArgs) ToContactProfileLinkOutput() ContactProfileLinkOutput {
	return i.ToContactProfileLinkOutputWithContext(context.Background())
}

func (i ContactProfileLinkArgs) ToContactProfileLinkOutputWithContext(ctx context.Context) ContactProfileLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactProfileLinkOutput)
}





type ContactProfileLinkArrayInput interface {
	pulumi.Input

	ToContactProfileLinkArrayOutput() ContactProfileLinkArrayOutput
	ToContactProfileLinkArrayOutputWithContext(context.Context) ContactProfileLinkArrayOutput
}

type ContactProfileLinkArray []ContactProfileLinkInput

func (ContactProfileLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContactProfileLink)(nil)).Elem()
}

func (i ContactProfileLinkArray) ToContactProfileLinkArrayOutput() ContactProfileLinkArrayOutput {
	return i.ToContactProfileLinkArrayOutputWithContext(context.Background())
}

func (i ContactProfileLinkArray) ToContactProfileLinkArrayOutputWithContext(ctx context.Context) ContactProfileLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactProfileLinkArrayOutput)
}

type ContactProfileLinkOutput struct{ *pulumi.OutputState }

func (ContactProfileLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactProfileLink)(nil)).Elem()
}

func (o ContactProfileLinkOutput) ToContactProfileLinkOutput() ContactProfileLinkOutput {
	return o
}

func (o ContactProfileLinkOutput) ToContactProfileLinkOutputWithContext(ctx context.Context) ContactProfileLinkOutput {
	return o
}

func (o ContactProfileLinkOutput) Channels() ContactProfileLinkChannelArrayOutput {
	return o.ApplyT(func(v ContactProfileLink) []ContactProfileLinkChannel { return v.Channels }).(ContactProfileLinkChannelArrayOutput)
}

func (o ContactProfileLinkOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v ContactProfileLink) string { return v.Direction }).(pulumi.StringOutput)
}

func (o ContactProfileLinkOutput) EirpdBW() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ContactProfileLink) *float64 { return v.EirpdBW }).(pulumi.Float64PtrOutput)
}

func (o ContactProfileLinkOutput) GainOverTemperature() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ContactProfileLink) *float64 { return v.GainOverTemperature }).(pulumi.Float64PtrOutput)
}

func (o ContactProfileLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContactProfileLink) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContactProfileLinkOutput) Polarization() pulumi.StringOutput {
	return o.ApplyT(func(v ContactProfileLink) string { return v.Polarization }).(pulumi.StringOutput)
}

type ContactProfileLinkArrayOutput struct{ *pulumi.OutputState }

func (ContactProfileLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContactProfileLink)(nil)).Elem()
}

func (o ContactProfileLinkArrayOutput) ToContactProfileLinkArrayOutput() ContactProfileLinkArrayOutput {
	return o
}

func (o ContactProfileLinkArrayOutput) ToContactProfileLinkArrayOutputWithContext(ctx context.Context) ContactProfileLinkArrayOutput {
	return o
}

func (o ContactProfileLinkArrayOutput) Index(i pulumi.IntInput) ContactProfileLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContactProfileLink {
		return vs[0].([]ContactProfileLink)[vs[1].(int)]
	}).(ContactProfileLinkOutput)
}

type ContactProfileLinkChannel struct {
	BandwidthMHz              float64  `pulumi:"bandwidthMHz"`
	CenterFrequencyMHz        float64  `pulumi:"centerFrequencyMHz"`
	DecodingConfiguration     *string  `pulumi:"decodingConfiguration"`
	DemodulationConfiguration *string  `pulumi:"demodulationConfiguration"`
	EncodingConfiguration     *string  `pulumi:"encodingConfiguration"`
	EndPoint                  EndPoint `pulumi:"endPoint"`
	ModulationConfiguration   *string  `pulumi:"modulationConfiguration"`
	Name                      string   `pulumi:"name"`
}





type ContactProfileLinkChannelInput interface {
	pulumi.Input

	ToContactProfileLinkChannelOutput() ContactProfileLinkChannelOutput
	ToContactProfileLinkChannelOutputWithContext(context.Context) ContactProfileLinkChannelOutput
}

type ContactProfileLinkChannelArgs struct {
	BandwidthMHz              pulumi.Float64Input   `pulumi:"bandwidthMHz"`
	CenterFrequencyMHz        pulumi.Float64Input   `pulumi:"centerFrequencyMHz"`
	DecodingConfiguration     pulumi.StringPtrInput `pulumi:"decodingConfiguration"`
	DemodulationConfiguration pulumi.StringPtrInput `pulumi:"demodulationConfiguration"`
	EncodingConfiguration     pulumi.StringPtrInput `pulumi:"encodingConfiguration"`
	EndPoint                  EndPointInput         `pulumi:"endPoint"`
	ModulationConfiguration   pulumi.StringPtrInput `pulumi:"modulationConfiguration"`
	Name                      pulumi.StringInput    `pulumi:"name"`
}

func (ContactProfileLinkChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactProfileLinkChannel)(nil)).Elem()
}

func (i ContactProfileLinkChannelArgs) ToContactProfileLinkChannelOutput() ContactProfileLinkChannelOutput {
	return i.ToContactProfileLinkChannelOutputWithContext(context.Background())
}

func (i ContactProfileLinkChannelArgs) ToContactProfileLinkChannelOutputWithContext(ctx context.Context) ContactProfileLinkChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactProfileLinkChannelOutput)
}





type ContactProfileLinkChannelArrayInput interface {
	pulumi.Input

	ToContactProfileLinkChannelArrayOutput() ContactProfileLinkChannelArrayOutput
	ToContactProfileLinkChannelArrayOutputWithContext(context.Context) ContactProfileLinkChannelArrayOutput
}

type ContactProfileLinkChannelArray []ContactProfileLinkChannelInput

func (ContactProfileLinkChannelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContactProfileLinkChannel)(nil)).Elem()
}

func (i ContactProfileLinkChannelArray) ToContactProfileLinkChannelArrayOutput() ContactProfileLinkChannelArrayOutput {
	return i.ToContactProfileLinkChannelArrayOutputWithContext(context.Background())
}

func (i ContactProfileLinkChannelArray) ToContactProfileLinkChannelArrayOutputWithContext(ctx context.Context) ContactProfileLinkChannelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactProfileLinkChannelArrayOutput)
}

type ContactProfileLinkChannelOutput struct{ *pulumi.OutputState }

func (ContactProfileLinkChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactProfileLinkChannel)(nil)).Elem()
}

func (o ContactProfileLinkChannelOutput) ToContactProfileLinkChannelOutput() ContactProfileLinkChannelOutput {
	return o
}

func (o ContactProfileLinkChannelOutput) ToContactProfileLinkChannelOutputWithContext(ctx context.Context) ContactProfileLinkChannelOutput {
	return o
}

func (o ContactProfileLinkChannelOutput) BandwidthMHz() pulumi.Float64Output {
	return o.ApplyT(func(v ContactProfileLinkChannel) float64 { return v.BandwidthMHz }).(pulumi.Float64Output)
}

func (o ContactProfileLinkChannelOutput) CenterFrequencyMHz() pulumi.Float64Output {
	return o.ApplyT(func(v ContactProfileLinkChannel) float64 { return v.CenterFrequencyMHz }).(pulumi.Float64Output)
}

func (o ContactProfileLinkChannelOutput) DecodingConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactProfileLinkChannel) *string { return v.DecodingConfiguration }).(pulumi.StringPtrOutput)
}

func (o ContactProfileLinkChannelOutput) DemodulationConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactProfileLinkChannel) *string { return v.DemodulationConfiguration }).(pulumi.StringPtrOutput)
}

func (o ContactProfileLinkChannelOutput) EncodingConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactProfileLinkChannel) *string { return v.EncodingConfiguration }).(pulumi.StringPtrOutput)
}

func (o ContactProfileLinkChannelOutput) EndPoint() EndPointOutput {
	return o.ApplyT(func(v ContactProfileLinkChannel) EndPoint { return v.EndPoint }).(EndPointOutput)
}

func (o ContactProfileLinkChannelOutput) ModulationConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactProfileLinkChannel) *string { return v.ModulationConfiguration }).(pulumi.StringPtrOutput)
}

func (o ContactProfileLinkChannelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContactProfileLinkChannel) string { return v.Name }).(pulumi.StringOutput)
}

type ContactProfileLinkChannelArrayOutput struct{ *pulumi.OutputState }

func (ContactProfileLinkChannelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContactProfileLinkChannel)(nil)).Elem()
}

func (o ContactProfileLinkChannelArrayOutput) ToContactProfileLinkChannelArrayOutput() ContactProfileLinkChannelArrayOutput {
	return o
}

func (o ContactProfileLinkChannelArrayOutput) ToContactProfileLinkChannelArrayOutputWithContext(ctx context.Context) ContactProfileLinkChannelArrayOutput {
	return o
}

func (o ContactProfileLinkChannelArrayOutput) Index(i pulumi.IntInput) ContactProfileLinkChannelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContactProfileLinkChannel {
		return vs[0].([]ContactProfileLinkChannel)[vs[1].(int)]
	}).(ContactProfileLinkChannelOutput)
}

type ContactProfileLinkChannelResponse struct {
	BandwidthMHz              float64          `pulumi:"bandwidthMHz"`
	CenterFrequencyMHz        float64          `pulumi:"centerFrequencyMHz"`
	DecodingConfiguration     *string          `pulumi:"decodingConfiguration"`
	DemodulationConfiguration *string          `pulumi:"demodulationConfiguration"`
	EncodingConfiguration     *string          `pulumi:"encodingConfiguration"`
	EndPoint                  EndPointResponse `pulumi:"endPoint"`
	ModulationConfiguration   *string          `pulumi:"modulationConfiguration"`
	Name                      string           `pulumi:"name"`
}

type ContactProfileLinkChannelResponseOutput struct{ *pulumi.OutputState }

func (ContactProfileLinkChannelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactProfileLinkChannelResponse)(nil)).Elem()
}

func (o ContactProfileLinkChannelResponseOutput) ToContactProfileLinkChannelResponseOutput() ContactProfileLinkChannelResponseOutput {
	return o
}

func (o ContactProfileLinkChannelResponseOutput) ToContactProfileLinkChannelResponseOutputWithContext(ctx context.Context) ContactProfileLinkChannelResponseOutput {
	return o
}

func (o ContactProfileLinkChannelResponseOutput) BandwidthMHz() pulumi.Float64Output {
	return o.ApplyT(func(v ContactProfileLinkChannelResponse) float64 { return v.BandwidthMHz }).(pulumi.Float64Output)
}

func (o ContactProfileLinkChannelResponseOutput) CenterFrequencyMHz() pulumi.Float64Output {
	return o.ApplyT(func(v ContactProfileLinkChannelResponse) float64 { return v.CenterFrequencyMHz }).(pulumi.Float64Output)
}

func (o ContactProfileLinkChannelResponseOutput) DecodingConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactProfileLinkChannelResponse) *string { return v.DecodingConfiguration }).(pulumi.StringPtrOutput)
}

func (o ContactProfileLinkChannelResponseOutput) DemodulationConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactProfileLinkChannelResponse) *string { return v.DemodulationConfiguration }).(pulumi.StringPtrOutput)
}

func (o ContactProfileLinkChannelResponseOutput) EncodingConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactProfileLinkChannelResponse) *string { return v.EncodingConfiguration }).(pulumi.StringPtrOutput)
}

func (o ContactProfileLinkChannelResponseOutput) EndPoint() EndPointResponseOutput {
	return o.ApplyT(func(v ContactProfileLinkChannelResponse) EndPointResponse { return v.EndPoint }).(EndPointResponseOutput)
}

func (o ContactProfileLinkChannelResponseOutput) ModulationConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactProfileLinkChannelResponse) *string { return v.ModulationConfiguration }).(pulumi.StringPtrOutput)
}

func (o ContactProfileLinkChannelResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContactProfileLinkChannelResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ContactProfileLinkChannelResponseArrayOutput struct{ *pulumi.OutputState }

func (ContactProfileLinkChannelResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContactProfileLinkChannelResponse)(nil)).Elem()
}

func (o ContactProfileLinkChannelResponseArrayOutput) ToContactProfileLinkChannelResponseArrayOutput() ContactProfileLinkChannelResponseArrayOutput {
	return o
}

func (o ContactProfileLinkChannelResponseArrayOutput) ToContactProfileLinkChannelResponseArrayOutputWithContext(ctx context.Context) ContactProfileLinkChannelResponseArrayOutput {
	return o
}

func (o ContactProfileLinkChannelResponseArrayOutput) Index(i pulumi.IntInput) ContactProfileLinkChannelResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContactProfileLinkChannelResponse {
		return vs[0].([]ContactProfileLinkChannelResponse)[vs[1].(int)]
	}).(ContactProfileLinkChannelResponseOutput)
}

type ContactProfileLinkResponse struct {
	Channels            []ContactProfileLinkChannelResponse `pulumi:"channels"`
	Direction           string                              `pulumi:"direction"`
	EirpdBW             *float64                            `pulumi:"eirpdBW"`
	GainOverTemperature *float64                            `pulumi:"gainOverTemperature"`
	Name                string                              `pulumi:"name"`
	Polarization        string                              `pulumi:"polarization"`
}

type ContactProfileLinkResponseOutput struct{ *pulumi.OutputState }

func (ContactProfileLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactProfileLinkResponse)(nil)).Elem()
}

func (o ContactProfileLinkResponseOutput) ToContactProfileLinkResponseOutput() ContactProfileLinkResponseOutput {
	return o
}

func (o ContactProfileLinkResponseOutput) ToContactProfileLinkResponseOutputWithContext(ctx context.Context) ContactProfileLinkResponseOutput {
	return o
}

func (o ContactProfileLinkResponseOutput) Channels() ContactProfileLinkChannelResponseArrayOutput {
	return o.ApplyT(func(v ContactProfileLinkResponse) []ContactProfileLinkChannelResponse { return v.Channels }).(ContactProfileLinkChannelResponseArrayOutput)
}

func (o ContactProfileLinkResponseOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v ContactProfileLinkResponse) string { return v.Direction }).(pulumi.StringOutput)
}

func (o ContactProfileLinkResponseOutput) EirpdBW() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ContactProfileLinkResponse) *float64 { return v.EirpdBW }).(pulumi.Float64PtrOutput)
}

func (o ContactProfileLinkResponseOutput) GainOverTemperature() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ContactProfileLinkResponse) *float64 { return v.GainOverTemperature }).(pulumi.Float64PtrOutput)
}

func (o ContactProfileLinkResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContactProfileLinkResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContactProfileLinkResponseOutput) Polarization() pulumi.StringOutput {
	return o.ApplyT(func(v ContactProfileLinkResponse) string { return v.Polarization }).(pulumi.StringOutput)
}

type ContactProfileLinkResponseArrayOutput struct{ *pulumi.OutputState }

func (ContactProfileLinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContactProfileLinkResponse)(nil)).Elem()
}

func (o ContactProfileLinkResponseArrayOutput) ToContactProfileLinkResponseArrayOutput() ContactProfileLinkResponseArrayOutput {
	return o
}

func (o ContactProfileLinkResponseArrayOutput) ToContactProfileLinkResponseArrayOutputWithContext(ctx context.Context) ContactProfileLinkResponseArrayOutput {
	return o
}

func (o ContactProfileLinkResponseArrayOutput) Index(i pulumi.IntInput) ContactProfileLinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContactProfileLinkResponse {
		return vs[0].([]ContactProfileLinkResponse)[vs[1].(int)]
	}).(ContactProfileLinkResponseOutput)
}

type ContactProfilesPropertiesNetworkConfiguration struct {
	SubnetId string `pulumi:"subnetId"`
}





type ContactProfilesPropertiesNetworkConfigurationInput interface {
	pulumi.Input

	ToContactProfilesPropertiesNetworkConfigurationOutput() ContactProfilesPropertiesNetworkConfigurationOutput
	ToContactProfilesPropertiesNetworkConfigurationOutputWithContext(context.Context) ContactProfilesPropertiesNetworkConfigurationOutput
}

type ContactProfilesPropertiesNetworkConfigurationArgs struct {
	SubnetId pulumi.StringInput `pulumi:"subnetId"`
}

func (ContactProfilesPropertiesNetworkConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactProfilesPropertiesNetworkConfiguration)(nil)).Elem()
}

func (i ContactProfilesPropertiesNetworkConfigurationArgs) ToContactProfilesPropertiesNetworkConfigurationOutput() ContactProfilesPropertiesNetworkConfigurationOutput {
	return i.ToContactProfilesPropertiesNetworkConfigurationOutputWithContext(context.Background())
}

func (i ContactProfilesPropertiesNetworkConfigurationArgs) ToContactProfilesPropertiesNetworkConfigurationOutputWithContext(ctx context.Context) ContactProfilesPropertiesNetworkConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactProfilesPropertiesNetworkConfigurationOutput)
}

type ContactProfilesPropertiesNetworkConfigurationOutput struct{ *pulumi.OutputState }

func (ContactProfilesPropertiesNetworkConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactProfilesPropertiesNetworkConfiguration)(nil)).Elem()
}

func (o ContactProfilesPropertiesNetworkConfigurationOutput) ToContactProfilesPropertiesNetworkConfigurationOutput() ContactProfilesPropertiesNetworkConfigurationOutput {
	return o
}

func (o ContactProfilesPropertiesNetworkConfigurationOutput) ToContactProfilesPropertiesNetworkConfigurationOutputWithContext(ctx context.Context) ContactProfilesPropertiesNetworkConfigurationOutput {
	return o
}

func (o ContactProfilesPropertiesNetworkConfigurationOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v ContactProfilesPropertiesNetworkConfiguration) string { return v.SubnetId }).(pulumi.StringOutput)
}

type ContactProfilesPropertiesResponseNetworkConfiguration struct {
	SubnetId string `pulumi:"subnetId"`
}

type ContactProfilesPropertiesResponseNetworkConfigurationOutput struct{ *pulumi.OutputState }

func (ContactProfilesPropertiesResponseNetworkConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactProfilesPropertiesResponseNetworkConfiguration)(nil)).Elem()
}

func (o ContactProfilesPropertiesResponseNetworkConfigurationOutput) ToContactProfilesPropertiesResponseNetworkConfigurationOutput() ContactProfilesPropertiesResponseNetworkConfigurationOutput {
	return o
}

func (o ContactProfilesPropertiesResponseNetworkConfigurationOutput) ToContactProfilesPropertiesResponseNetworkConfigurationOutputWithContext(ctx context.Context) ContactProfilesPropertiesResponseNetworkConfigurationOutput {
	return o
}

func (o ContactProfilesPropertiesResponseNetworkConfigurationOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v ContactProfilesPropertiesResponseNetworkConfiguration) string { return v.SubnetId }).(pulumi.StringOutput)
}

type ContactsPropertiesContactProfile struct {
	Id *string `pulumi:"id"`
}





type ContactsPropertiesContactProfileInput interface {
	pulumi.Input

	ToContactsPropertiesContactProfileOutput() ContactsPropertiesContactProfileOutput
	ToContactsPropertiesContactProfileOutputWithContext(context.Context) ContactsPropertiesContactProfileOutput
}

type ContactsPropertiesContactProfileArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ContactsPropertiesContactProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactsPropertiesContactProfile)(nil)).Elem()
}

func (i ContactsPropertiesContactProfileArgs) ToContactsPropertiesContactProfileOutput() ContactsPropertiesContactProfileOutput {
	return i.ToContactsPropertiesContactProfileOutputWithContext(context.Background())
}

func (i ContactsPropertiesContactProfileArgs) ToContactsPropertiesContactProfileOutputWithContext(ctx context.Context) ContactsPropertiesContactProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactsPropertiesContactProfileOutput)
}

type ContactsPropertiesContactProfileOutput struct{ *pulumi.OutputState }

func (ContactsPropertiesContactProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactsPropertiesContactProfile)(nil)).Elem()
}

func (o ContactsPropertiesContactProfileOutput) ToContactsPropertiesContactProfileOutput() ContactsPropertiesContactProfileOutput {
	return o
}

func (o ContactsPropertiesContactProfileOutput) ToContactsPropertiesContactProfileOutputWithContext(ctx context.Context) ContactsPropertiesContactProfileOutput {
	return o
}

func (o ContactsPropertiesContactProfileOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactsPropertiesContactProfile) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ContactsPropertiesResponseAntennaConfiguration struct {
	DestinationIp *string  `pulumi:"destinationIp"`
	SourceIps     []string `pulumi:"sourceIps"`
}

type ContactsPropertiesResponseAntennaConfigurationOutput struct{ *pulumi.OutputState }

func (ContactsPropertiesResponseAntennaConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactsPropertiesResponseAntennaConfiguration)(nil)).Elem()
}

func (o ContactsPropertiesResponseAntennaConfigurationOutput) ToContactsPropertiesResponseAntennaConfigurationOutput() ContactsPropertiesResponseAntennaConfigurationOutput {
	return o
}

func (o ContactsPropertiesResponseAntennaConfigurationOutput) ToContactsPropertiesResponseAntennaConfigurationOutputWithContext(ctx context.Context) ContactsPropertiesResponseAntennaConfigurationOutput {
	return o
}

func (o ContactsPropertiesResponseAntennaConfigurationOutput) DestinationIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactsPropertiesResponseAntennaConfiguration) *string { return v.DestinationIp }).(pulumi.StringPtrOutput)
}

func (o ContactsPropertiesResponseAntennaConfigurationOutput) SourceIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContactsPropertiesResponseAntennaConfiguration) []string { return v.SourceIps }).(pulumi.StringArrayOutput)
}

type ContactsPropertiesResponseContactProfile struct {
	Id *string `pulumi:"id"`
}

type ContactsPropertiesResponseContactProfileOutput struct{ *pulumi.OutputState }

func (ContactsPropertiesResponseContactProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactsPropertiesResponseContactProfile)(nil)).Elem()
}

func (o ContactsPropertiesResponseContactProfileOutput) ToContactsPropertiesResponseContactProfileOutput() ContactsPropertiesResponseContactProfileOutput {
	return o
}

func (o ContactsPropertiesResponseContactProfileOutput) ToContactsPropertiesResponseContactProfileOutputWithContext(ctx context.Context) ContactsPropertiesResponseContactProfileOutput {
	return o
}

func (o ContactsPropertiesResponseContactProfileOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactsPropertiesResponseContactProfile) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type EndPoint struct {
	EndPointName string `pulumi:"endPointName"`
	IpAddress    string `pulumi:"ipAddress"`
	Port         string `pulumi:"port"`
	Protocol     string `pulumi:"protocol"`
}





type EndPointInput interface {
	pulumi.Input

	ToEndPointOutput() EndPointOutput
	ToEndPointOutputWithContext(context.Context) EndPointOutput
}

type EndPointArgs struct {
	EndPointName pulumi.StringInput `pulumi:"endPointName"`
	IpAddress    pulumi.StringInput `pulumi:"ipAddress"`
	Port         pulumi.StringInput `pulumi:"port"`
	Protocol     pulumi.StringInput `pulumi:"protocol"`
}

func (EndPointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndPoint)(nil)).Elem()
}

func (i EndPointArgs) ToEndPointOutput() EndPointOutput {
	return i.ToEndPointOutputWithContext(context.Background())
}

func (i EndPointArgs) ToEndPointOutputWithContext(ctx context.Context) EndPointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndPointOutput)
}

type EndPointOutput struct{ *pulumi.OutputState }

func (EndPointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndPoint)(nil)).Elem()
}

func (o EndPointOutput) ToEndPointOutput() EndPointOutput {
	return o
}

func (o EndPointOutput) ToEndPointOutputWithContext(ctx context.Context) EndPointOutput {
	return o
}

func (o EndPointOutput) EndPointName() pulumi.StringOutput {
	return o.ApplyT(func(v EndPoint) string { return v.EndPointName }).(pulumi.StringOutput)
}

func (o EndPointOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v EndPoint) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o EndPointOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v EndPoint) string { return v.Port }).(pulumi.StringOutput)
}

func (o EndPointOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v EndPoint) string { return v.Protocol }).(pulumi.StringOutput)
}

type EndPointResponse struct {
	EndPointName string `pulumi:"endPointName"`
	IpAddress    string `pulumi:"ipAddress"`
	Port         string `pulumi:"port"`
	Protocol     string `pulumi:"protocol"`
}

type EndPointResponseOutput struct{ *pulumi.OutputState }

func (EndPointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndPointResponse)(nil)).Elem()
}

func (o EndPointResponseOutput) ToEndPointResponseOutput() EndPointResponseOutput {
	return o
}

func (o EndPointResponseOutput) ToEndPointResponseOutputWithContext(ctx context.Context) EndPointResponseOutput {
	return o
}

func (o EndPointResponseOutput) EndPointName() pulumi.StringOutput {
	return o.ApplyT(func(v EndPointResponse) string { return v.EndPointName }).(pulumi.StringOutput)
}

func (o EndPointResponseOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v EndPointResponse) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o EndPointResponseOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v EndPointResponse) string { return v.Port }).(pulumi.StringOutput)
}

func (o EndPointResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v EndPointResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

type SpacecraftLink struct {
	BandwidthMHz       float64 `pulumi:"bandwidthMHz"`
	CenterFrequencyMHz float64 `pulumi:"centerFrequencyMHz"`
	Direction          string  `pulumi:"direction"`
	Name               string  `pulumi:"name"`
	Polarization       string  `pulumi:"polarization"`
}





type SpacecraftLinkInput interface {
	pulumi.Input

	ToSpacecraftLinkOutput() SpacecraftLinkOutput
	ToSpacecraftLinkOutputWithContext(context.Context) SpacecraftLinkOutput
}

type SpacecraftLinkArgs struct {
	BandwidthMHz       pulumi.Float64Input `pulumi:"bandwidthMHz"`
	CenterFrequencyMHz pulumi.Float64Input `pulumi:"centerFrequencyMHz"`
	Direction          pulumi.StringInput  `pulumi:"direction"`
	Name               pulumi.StringInput  `pulumi:"name"`
	Polarization       pulumi.StringInput  `pulumi:"polarization"`
}

func (SpacecraftLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SpacecraftLink)(nil)).Elem()
}

func (i SpacecraftLinkArgs) ToSpacecraftLinkOutput() SpacecraftLinkOutput {
	return i.ToSpacecraftLinkOutputWithContext(context.Background())
}

func (i SpacecraftLinkArgs) ToSpacecraftLinkOutputWithContext(ctx context.Context) SpacecraftLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpacecraftLinkOutput)
}





type SpacecraftLinkArrayInput interface {
	pulumi.Input

	ToSpacecraftLinkArrayOutput() SpacecraftLinkArrayOutput
	ToSpacecraftLinkArrayOutputWithContext(context.Context) SpacecraftLinkArrayOutput
}

type SpacecraftLinkArray []SpacecraftLinkInput

func (SpacecraftLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SpacecraftLink)(nil)).Elem()
}

func (i SpacecraftLinkArray) ToSpacecraftLinkArrayOutput() SpacecraftLinkArrayOutput {
	return i.ToSpacecraftLinkArrayOutputWithContext(context.Background())
}

func (i SpacecraftLinkArray) ToSpacecraftLinkArrayOutputWithContext(ctx context.Context) SpacecraftLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpacecraftLinkArrayOutput)
}

type SpacecraftLinkOutput struct{ *pulumi.OutputState }

func (SpacecraftLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SpacecraftLink)(nil)).Elem()
}

func (o SpacecraftLinkOutput) ToSpacecraftLinkOutput() SpacecraftLinkOutput {
	return o
}

func (o SpacecraftLinkOutput) ToSpacecraftLinkOutputWithContext(ctx context.Context) SpacecraftLinkOutput {
	return o
}

func (o SpacecraftLinkOutput) BandwidthMHz() pulumi.Float64Output {
	return o.ApplyT(func(v SpacecraftLink) float64 { return v.BandwidthMHz }).(pulumi.Float64Output)
}

func (o SpacecraftLinkOutput) CenterFrequencyMHz() pulumi.Float64Output {
	return o.ApplyT(func(v SpacecraftLink) float64 { return v.CenterFrequencyMHz }).(pulumi.Float64Output)
}

func (o SpacecraftLinkOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v SpacecraftLink) string { return v.Direction }).(pulumi.StringOutput)
}

func (o SpacecraftLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SpacecraftLink) string { return v.Name }).(pulumi.StringOutput)
}

func (o SpacecraftLinkOutput) Polarization() pulumi.StringOutput {
	return o.ApplyT(func(v SpacecraftLink) string { return v.Polarization }).(pulumi.StringOutput)
}

type SpacecraftLinkArrayOutput struct{ *pulumi.OutputState }

func (SpacecraftLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SpacecraftLink)(nil)).Elem()
}

func (o SpacecraftLinkArrayOutput) ToSpacecraftLinkArrayOutput() SpacecraftLinkArrayOutput {
	return o
}

func (o SpacecraftLinkArrayOutput) ToSpacecraftLinkArrayOutputWithContext(ctx context.Context) SpacecraftLinkArrayOutput {
	return o
}

func (o SpacecraftLinkArrayOutput) Index(i pulumi.IntInput) SpacecraftLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SpacecraftLink {
		return vs[0].([]SpacecraftLink)[vs[1].(int)]
	}).(SpacecraftLinkOutput)
}

type SpacecraftLinkResponse struct {
	Authorizations     []AuthorizedGroundstationResponse `pulumi:"authorizations"`
	BandwidthMHz       float64                           `pulumi:"bandwidthMHz"`
	CenterFrequencyMHz float64                           `pulumi:"centerFrequencyMHz"`
	Direction          string                            `pulumi:"direction"`
	Name               string                            `pulumi:"name"`
	Polarization       string                            `pulumi:"polarization"`
}

type SpacecraftLinkResponseOutput struct{ *pulumi.OutputState }

func (SpacecraftLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SpacecraftLinkResponse)(nil)).Elem()
}

func (o SpacecraftLinkResponseOutput) ToSpacecraftLinkResponseOutput() SpacecraftLinkResponseOutput {
	return o
}

func (o SpacecraftLinkResponseOutput) ToSpacecraftLinkResponseOutputWithContext(ctx context.Context) SpacecraftLinkResponseOutput {
	return o
}

func (o SpacecraftLinkResponseOutput) Authorizations() AuthorizedGroundstationResponseArrayOutput {
	return o.ApplyT(func(v SpacecraftLinkResponse) []AuthorizedGroundstationResponse { return v.Authorizations }).(AuthorizedGroundstationResponseArrayOutput)
}

func (o SpacecraftLinkResponseOutput) BandwidthMHz() pulumi.Float64Output {
	return o.ApplyT(func(v SpacecraftLinkResponse) float64 { return v.BandwidthMHz }).(pulumi.Float64Output)
}

func (o SpacecraftLinkResponseOutput) CenterFrequencyMHz() pulumi.Float64Output {
	return o.ApplyT(func(v SpacecraftLinkResponse) float64 { return v.CenterFrequencyMHz }).(pulumi.Float64Output)
}

func (o SpacecraftLinkResponseOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v SpacecraftLinkResponse) string { return v.Direction }).(pulumi.StringOutput)
}

func (o SpacecraftLinkResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SpacecraftLinkResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SpacecraftLinkResponseOutput) Polarization() pulumi.StringOutput {
	return o.ApplyT(func(v SpacecraftLinkResponse) string { return v.Polarization }).(pulumi.StringOutput)
}

type SpacecraftLinkResponseArrayOutput struct{ *pulumi.OutputState }

func (SpacecraftLinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SpacecraftLinkResponse)(nil)).Elem()
}

func (o SpacecraftLinkResponseArrayOutput) ToSpacecraftLinkResponseArrayOutput() SpacecraftLinkResponseArrayOutput {
	return o
}

func (o SpacecraftLinkResponseArrayOutput) ToSpacecraftLinkResponseArrayOutputWithContext(ctx context.Context) SpacecraftLinkResponseArrayOutput {
	return o
}

func (o SpacecraftLinkResponseArrayOutput) Index(i pulumi.IntInput) SpacecraftLinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SpacecraftLinkResponse {
		return vs[0].([]SpacecraftLinkResponse)[vs[1].(int)]
	}).(SpacecraftLinkResponseOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorizedGroundstationResponseOutput{})
	pulumi.RegisterOutputType(AuthorizedGroundstationResponseArrayOutput{})
	pulumi.RegisterOutputType(AvailableContactsResponseOutput{})
	pulumi.RegisterOutputType(AvailableContactsResponseArrayOutput{})
	pulumi.RegisterOutputType(AvailableContactsResponseSpacecraftOutput{})
	pulumi.RegisterOutputType(AvailableContactsResponseSpacecraftPtrOutput{})
	pulumi.RegisterOutputType(ContactParametersContactProfileOutput{})
	pulumi.RegisterOutputType(ContactProfileLinkOutput{})
	pulumi.RegisterOutputType(ContactProfileLinkArrayOutput{})
	pulumi.RegisterOutputType(ContactProfileLinkChannelOutput{})
	pulumi.RegisterOutputType(ContactProfileLinkChannelArrayOutput{})
	pulumi.RegisterOutputType(ContactProfileLinkChannelResponseOutput{})
	pulumi.RegisterOutputType(ContactProfileLinkChannelResponseArrayOutput{})
	pulumi.RegisterOutputType(ContactProfileLinkResponseOutput{})
	pulumi.RegisterOutputType(ContactProfileLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(ContactProfilesPropertiesNetworkConfigurationOutput{})
	pulumi.RegisterOutputType(ContactProfilesPropertiesResponseNetworkConfigurationOutput{})
	pulumi.RegisterOutputType(ContactsPropertiesContactProfileOutput{})
	pulumi.RegisterOutputType(ContactsPropertiesResponseAntennaConfigurationOutput{})
	pulumi.RegisterOutputType(ContactsPropertiesResponseContactProfileOutput{})
	pulumi.RegisterOutputType(EndPointOutput{})
	pulumi.RegisterOutputType(EndPointResponseOutput{})
	pulumi.RegisterOutputType(SpacecraftLinkOutput{})
	pulumi.RegisterOutputType(SpacecraftLinkArrayOutput{})
	pulumi.RegisterOutputType(SpacecraftLinkResponseOutput{})
	pulumi.RegisterOutputType(SpacecraftLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}

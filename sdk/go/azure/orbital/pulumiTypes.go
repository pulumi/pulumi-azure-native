


package orbital

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AvailableContactsResponse struct {
	EndAzimuthDegrees       float64                    `pulumi:"endAzimuthDegrees"`
	EndElevationDegrees     float64                    `pulumi:"endElevationDegrees"`
	GroundStationName       string                     `pulumi:"groundStationName"`
	MaximumElevationDegrees float64                    `pulumi:"maximumElevationDegrees"`
	RxEndTime               string                     `pulumi:"rxEndTime"`
	RxStartTime             string                     `pulumi:"rxStartTime"`
	Spacecraft              *ResourceReferenceResponse `pulumi:"spacecraft"`
	StartAzimuthDegrees     float64                    `pulumi:"startAzimuthDegrees"`
	StartElevationDegrees   float64                    `pulumi:"startElevationDegrees"`
	TxEndTime               string                     `pulumi:"txEndTime"`
	TxStartTime             string                     `pulumi:"txStartTime"`
}





type AvailableContactsResponseInput interface {
	pulumi.Input

	ToAvailableContactsResponseOutput() AvailableContactsResponseOutput
	ToAvailableContactsResponseOutputWithContext(context.Context) AvailableContactsResponseOutput
}

type AvailableContactsResponseArgs struct {
	EndAzimuthDegrees       pulumi.Float64Input               `pulumi:"endAzimuthDegrees"`
	EndElevationDegrees     pulumi.Float64Input               `pulumi:"endElevationDegrees"`
	GroundStationName       pulumi.StringInput                `pulumi:"groundStationName"`
	MaximumElevationDegrees pulumi.Float64Input               `pulumi:"maximumElevationDegrees"`
	RxEndTime               pulumi.StringInput                `pulumi:"rxEndTime"`
	RxStartTime             pulumi.StringInput                `pulumi:"rxStartTime"`
	Spacecraft              ResourceReferenceResponsePtrInput `pulumi:"spacecraft"`
	StartAzimuthDegrees     pulumi.Float64Input               `pulumi:"startAzimuthDegrees"`
	StartElevationDegrees   pulumi.Float64Input               `pulumi:"startElevationDegrees"`
	TxEndTime               pulumi.StringInput                `pulumi:"txEndTime"`
	TxStartTime             pulumi.StringInput                `pulumi:"txStartTime"`
}

func (AvailableContactsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailableContactsResponse)(nil)).Elem()
}

func (i AvailableContactsResponseArgs) ToAvailableContactsResponseOutput() AvailableContactsResponseOutput {
	return i.ToAvailableContactsResponseOutputWithContext(context.Background())
}

func (i AvailableContactsResponseArgs) ToAvailableContactsResponseOutputWithContext(ctx context.Context) AvailableContactsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailableContactsResponseOutput)
}





type AvailableContactsResponseArrayInput interface {
	pulumi.Input

	ToAvailableContactsResponseArrayOutput() AvailableContactsResponseArrayOutput
	ToAvailableContactsResponseArrayOutputWithContext(context.Context) AvailableContactsResponseArrayOutput
}

type AvailableContactsResponseArray []AvailableContactsResponseInput

func (AvailableContactsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AvailableContactsResponse)(nil)).Elem()
}

func (i AvailableContactsResponseArray) ToAvailableContactsResponseArrayOutput() AvailableContactsResponseArrayOutput {
	return i.ToAvailableContactsResponseArrayOutputWithContext(context.Background())
}

func (i AvailableContactsResponseArray) ToAvailableContactsResponseArrayOutputWithContext(ctx context.Context) AvailableContactsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailableContactsResponseArrayOutput)
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

func (o AvailableContactsResponseOutput) Spacecraft() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v AvailableContactsResponse) *ResourceReferenceResponse { return v.Spacecraft }).(ResourceReferenceResponsePtrOutput)
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

type ContactProfileLink struct {
	Channels            []ContactProfileLinkChannel `pulumi:"channels"`
	Direction           string                      `pulumi:"direction"`
	EirpdBW             *float64                    `pulumi:"eirpdBW"`
	GainOverTemperature *float64                    `pulumi:"gainOverTemperature"`
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
}





type ContactProfileLinkChannelResponseInput interface {
	pulumi.Input

	ToContactProfileLinkChannelResponseOutput() ContactProfileLinkChannelResponseOutput
	ToContactProfileLinkChannelResponseOutputWithContext(context.Context) ContactProfileLinkChannelResponseOutput
}

type ContactProfileLinkChannelResponseArgs struct {
	BandwidthMHz              pulumi.Float64Input   `pulumi:"bandwidthMHz"`
	CenterFrequencyMHz        pulumi.Float64Input   `pulumi:"centerFrequencyMHz"`
	DecodingConfiguration     pulumi.StringPtrInput `pulumi:"decodingConfiguration"`
	DemodulationConfiguration pulumi.StringPtrInput `pulumi:"demodulationConfiguration"`
	EncodingConfiguration     pulumi.StringPtrInput `pulumi:"encodingConfiguration"`
	EndPoint                  EndPointResponseInput `pulumi:"endPoint"`
	ModulationConfiguration   pulumi.StringPtrInput `pulumi:"modulationConfiguration"`
}

func (ContactProfileLinkChannelResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactProfileLinkChannelResponse)(nil)).Elem()
}

func (i ContactProfileLinkChannelResponseArgs) ToContactProfileLinkChannelResponseOutput() ContactProfileLinkChannelResponseOutput {
	return i.ToContactProfileLinkChannelResponseOutputWithContext(context.Background())
}

func (i ContactProfileLinkChannelResponseArgs) ToContactProfileLinkChannelResponseOutputWithContext(ctx context.Context) ContactProfileLinkChannelResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactProfileLinkChannelResponseOutput)
}





type ContactProfileLinkChannelResponseArrayInput interface {
	pulumi.Input

	ToContactProfileLinkChannelResponseArrayOutput() ContactProfileLinkChannelResponseArrayOutput
	ToContactProfileLinkChannelResponseArrayOutputWithContext(context.Context) ContactProfileLinkChannelResponseArrayOutput
}

type ContactProfileLinkChannelResponseArray []ContactProfileLinkChannelResponseInput

func (ContactProfileLinkChannelResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContactProfileLinkChannelResponse)(nil)).Elem()
}

func (i ContactProfileLinkChannelResponseArray) ToContactProfileLinkChannelResponseArrayOutput() ContactProfileLinkChannelResponseArrayOutput {
	return i.ToContactProfileLinkChannelResponseArrayOutputWithContext(context.Background())
}

func (i ContactProfileLinkChannelResponseArray) ToContactProfileLinkChannelResponseArrayOutputWithContext(ctx context.Context) ContactProfileLinkChannelResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactProfileLinkChannelResponseArrayOutput)
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
	Polarization        string                              `pulumi:"polarization"`
}





type ContactProfileLinkResponseInput interface {
	pulumi.Input

	ToContactProfileLinkResponseOutput() ContactProfileLinkResponseOutput
	ToContactProfileLinkResponseOutputWithContext(context.Context) ContactProfileLinkResponseOutput
}

type ContactProfileLinkResponseArgs struct {
	Channels            ContactProfileLinkChannelResponseArrayInput `pulumi:"channels"`
	Direction           pulumi.StringInput                          `pulumi:"direction"`
	EirpdBW             pulumi.Float64PtrInput                      `pulumi:"eirpdBW"`
	GainOverTemperature pulumi.Float64PtrInput                      `pulumi:"gainOverTemperature"`
	Polarization        pulumi.StringInput                          `pulumi:"polarization"`
}

func (ContactProfileLinkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactProfileLinkResponse)(nil)).Elem()
}

func (i ContactProfileLinkResponseArgs) ToContactProfileLinkResponseOutput() ContactProfileLinkResponseOutput {
	return i.ToContactProfileLinkResponseOutputWithContext(context.Background())
}

func (i ContactProfileLinkResponseArgs) ToContactProfileLinkResponseOutputWithContext(ctx context.Context) ContactProfileLinkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactProfileLinkResponseOutput)
}





type ContactProfileLinkResponseArrayInput interface {
	pulumi.Input

	ToContactProfileLinkResponseArrayOutput() ContactProfileLinkResponseArrayOutput
	ToContactProfileLinkResponseArrayOutputWithContext(context.Context) ContactProfileLinkResponseArrayOutput
}

type ContactProfileLinkResponseArray []ContactProfileLinkResponseInput

func (ContactProfileLinkResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContactProfileLinkResponse)(nil)).Elem()
}

func (i ContactProfileLinkResponseArray) ToContactProfileLinkResponseArrayOutput() ContactProfileLinkResponseArrayOutput {
	return i.ToContactProfileLinkResponseArrayOutputWithContext(context.Background())
}

func (i ContactProfileLinkResponseArray) ToContactProfileLinkResponseArrayOutputWithContext(ctx context.Context) ContactProfileLinkResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactProfileLinkResponseArrayOutput)
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





type EndPointResponseInput interface {
	pulumi.Input

	ToEndPointResponseOutput() EndPointResponseOutput
	ToEndPointResponseOutputWithContext(context.Context) EndPointResponseOutput
}

type EndPointResponseArgs struct {
	EndPointName pulumi.StringInput `pulumi:"endPointName"`
	IpAddress    pulumi.StringInput `pulumi:"ipAddress"`
	Port         pulumi.StringInput `pulumi:"port"`
	Protocol     pulumi.StringInput `pulumi:"protocol"`
}

func (EndPointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndPointResponse)(nil)).Elem()
}

func (i EndPointResponseArgs) ToEndPointResponseOutput() EndPointResponseOutput {
	return i.ToEndPointResponseOutputWithContext(context.Background())
}

func (i EndPointResponseArgs) ToEndPointResponseOutputWithContext(ctx context.Context) EndPointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndPointResponseOutput)
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

type ResourceReference struct {
	Id *string `pulumi:"id"`
}





type ResourceReferenceInput interface {
	pulumi.Input

	ToResourceReferenceOutput() ResourceReferenceOutput
	ToResourceReferenceOutputWithContext(context.Context) ResourceReferenceOutput
}

type ResourceReferenceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ResourceReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReference)(nil)).Elem()
}

func (i ResourceReferenceArgs) ToResourceReferenceOutput() ResourceReferenceOutput {
	return i.ToResourceReferenceOutputWithContext(context.Background())
}

func (i ResourceReferenceArgs) ToResourceReferenceOutputWithContext(ctx context.Context) ResourceReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceOutput)
}

func (i ResourceReferenceArgs) ToResourceReferencePtrOutput() ResourceReferencePtrOutput {
	return i.ToResourceReferencePtrOutputWithContext(context.Background())
}

func (i ResourceReferenceArgs) ToResourceReferencePtrOutputWithContext(ctx context.Context) ResourceReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceOutput).ToResourceReferencePtrOutputWithContext(ctx)
}









type ResourceReferencePtrInput interface {
	pulumi.Input

	ToResourceReferencePtrOutput() ResourceReferencePtrOutput
	ToResourceReferencePtrOutputWithContext(context.Context) ResourceReferencePtrOutput
}

type resourceReferencePtrType ResourceReferenceArgs

func ResourceReferencePtr(v *ResourceReferenceArgs) ResourceReferencePtrInput {
	return (*resourceReferencePtrType)(v)
}

func (*resourceReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceReference)(nil)).Elem()
}

func (i *resourceReferencePtrType) ToResourceReferencePtrOutput() ResourceReferencePtrOutput {
	return i.ToResourceReferencePtrOutputWithContext(context.Background())
}

func (i *resourceReferencePtrType) ToResourceReferencePtrOutputWithContext(ctx context.Context) ResourceReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferencePtrOutput)
}

type ResourceReferenceOutput struct{ *pulumi.OutputState }

func (ResourceReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReference)(nil)).Elem()
}

func (o ResourceReferenceOutput) ToResourceReferenceOutput() ResourceReferenceOutput {
	return o
}

func (o ResourceReferenceOutput) ToResourceReferenceOutputWithContext(ctx context.Context) ResourceReferenceOutput {
	return o
}

func (o ResourceReferenceOutput) ToResourceReferencePtrOutput() ResourceReferencePtrOutput {
	return o.ToResourceReferencePtrOutputWithContext(context.Background())
}

func (o ResourceReferenceOutput) ToResourceReferencePtrOutputWithContext(ctx context.Context) ResourceReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceReference) *ResourceReference {
		return &v
	}).(ResourceReferencePtrOutput)
}

func (o ResourceReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ResourceReferencePtrOutput struct{ *pulumi.OutputState }

func (ResourceReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceReference)(nil)).Elem()
}

func (o ResourceReferencePtrOutput) ToResourceReferencePtrOutput() ResourceReferencePtrOutput {
	return o
}

func (o ResourceReferencePtrOutput) ToResourceReferencePtrOutputWithContext(ctx context.Context) ResourceReferencePtrOutput {
	return o
}

func (o ResourceReferencePtrOutput) Elem() ResourceReferenceOutput {
	return o.ApplyT(func(v *ResourceReference) ResourceReference {
		if v != nil {
			return *v
		}
		var ret ResourceReference
		return ret
	}).(ResourceReferenceOutput)
}

func (o ResourceReferencePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceReference) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type ResourceReferenceResponse struct {
	Id *string `pulumi:"id"`
}





type ResourceReferenceResponseInput interface {
	pulumi.Input

	ToResourceReferenceResponseOutput() ResourceReferenceResponseOutput
	ToResourceReferenceResponseOutputWithContext(context.Context) ResourceReferenceResponseOutput
}

type ResourceReferenceResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ResourceReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReferenceResponse)(nil)).Elem()
}

func (i ResourceReferenceResponseArgs) ToResourceReferenceResponseOutput() ResourceReferenceResponseOutput {
	return i.ToResourceReferenceResponseOutputWithContext(context.Background())
}

func (i ResourceReferenceResponseArgs) ToResourceReferenceResponseOutputWithContext(ctx context.Context) ResourceReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceResponseOutput)
}

func (i ResourceReferenceResponseArgs) ToResourceReferenceResponsePtrOutput() ResourceReferenceResponsePtrOutput {
	return i.ToResourceReferenceResponsePtrOutputWithContext(context.Background())
}

func (i ResourceReferenceResponseArgs) ToResourceReferenceResponsePtrOutputWithContext(ctx context.Context) ResourceReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceResponseOutput).ToResourceReferenceResponsePtrOutputWithContext(ctx)
}









type ResourceReferenceResponsePtrInput interface {
	pulumi.Input

	ToResourceReferenceResponsePtrOutput() ResourceReferenceResponsePtrOutput
	ToResourceReferenceResponsePtrOutputWithContext(context.Context) ResourceReferenceResponsePtrOutput
}

type resourceReferenceResponsePtrType ResourceReferenceResponseArgs

func ResourceReferenceResponsePtr(v *ResourceReferenceResponseArgs) ResourceReferenceResponsePtrInput {
	return (*resourceReferenceResponsePtrType)(v)
}

func (*resourceReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceReferenceResponse)(nil)).Elem()
}

func (i *resourceReferenceResponsePtrType) ToResourceReferenceResponsePtrOutput() ResourceReferenceResponsePtrOutput {
	return i.ToResourceReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *resourceReferenceResponsePtrType) ToResourceReferenceResponsePtrOutputWithContext(ctx context.Context) ResourceReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceResponsePtrOutput)
}

type ResourceReferenceResponseOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponseOutput() ResourceReferenceResponseOutput {
	return o
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponseOutputWithContext(ctx context.Context) ResourceReferenceResponseOutput {
	return o
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponsePtrOutput() ResourceReferenceResponsePtrOutput {
	return o.ToResourceReferenceResponsePtrOutputWithContext(context.Background())
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponsePtrOutputWithContext(ctx context.Context) ResourceReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceReferenceResponse) *ResourceReferenceResponse {
		return &v
	}).(ResourceReferenceResponsePtrOutput)
}

func (o ResourceReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ResourceReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponsePtrOutput) ToResourceReferenceResponsePtrOutput() ResourceReferenceResponsePtrOutput {
	return o
}

func (o ResourceReferenceResponsePtrOutput) ToResourceReferenceResponsePtrOutputWithContext(ctx context.Context) ResourceReferenceResponsePtrOutput {
	return o
}

func (o ResourceReferenceResponsePtrOutput) Elem() ResourceReferenceResponseOutput {
	return o.ApplyT(func(v *ResourceReferenceResponse) ResourceReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ResourceReferenceResponse
		return ret
	}).(ResourceReferenceResponseOutput)
}

func (o ResourceReferenceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SpacecraftLink struct {
	BandwidthMHz       float64 `pulumi:"bandwidthMHz"`
	CenterFrequencyMHz float64 `pulumi:"centerFrequencyMHz"`
	Direction          string  `pulumi:"direction"`
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
	BandwidthMHz       float64 `pulumi:"bandwidthMHz"`
	CenterFrequencyMHz float64 `pulumi:"centerFrequencyMHz"`
	Direction          string  `pulumi:"direction"`
	Polarization       string  `pulumi:"polarization"`
}





type SpacecraftLinkResponseInput interface {
	pulumi.Input

	ToSpacecraftLinkResponseOutput() SpacecraftLinkResponseOutput
	ToSpacecraftLinkResponseOutputWithContext(context.Context) SpacecraftLinkResponseOutput
}

type SpacecraftLinkResponseArgs struct {
	BandwidthMHz       pulumi.Float64Input `pulumi:"bandwidthMHz"`
	CenterFrequencyMHz pulumi.Float64Input `pulumi:"centerFrequencyMHz"`
	Direction          pulumi.StringInput  `pulumi:"direction"`
	Polarization       pulumi.StringInput  `pulumi:"polarization"`
}

func (SpacecraftLinkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SpacecraftLinkResponse)(nil)).Elem()
}

func (i SpacecraftLinkResponseArgs) ToSpacecraftLinkResponseOutput() SpacecraftLinkResponseOutput {
	return i.ToSpacecraftLinkResponseOutputWithContext(context.Background())
}

func (i SpacecraftLinkResponseArgs) ToSpacecraftLinkResponseOutputWithContext(ctx context.Context) SpacecraftLinkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpacecraftLinkResponseOutput)
}





type SpacecraftLinkResponseArrayInput interface {
	pulumi.Input

	ToSpacecraftLinkResponseArrayOutput() SpacecraftLinkResponseArrayOutput
	ToSpacecraftLinkResponseArrayOutputWithContext(context.Context) SpacecraftLinkResponseArrayOutput
}

type SpacecraftLinkResponseArray []SpacecraftLinkResponseInput

func (SpacecraftLinkResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SpacecraftLinkResponse)(nil)).Elem()
}

func (i SpacecraftLinkResponseArray) ToSpacecraftLinkResponseArrayOutput() SpacecraftLinkResponseArrayOutput {
	return i.ToSpacecraftLinkResponseArrayOutputWithContext(context.Background())
}

func (i SpacecraftLinkResponseArray) ToSpacecraftLinkResponseArrayOutputWithContext(ctx context.Context) SpacecraftLinkResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpacecraftLinkResponseArrayOutput)
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

func (o SpacecraftLinkResponseOutput) BandwidthMHz() pulumi.Float64Output {
	return o.ApplyT(func(v SpacecraftLinkResponse) float64 { return v.BandwidthMHz }).(pulumi.Float64Output)
}

func (o SpacecraftLinkResponseOutput) CenterFrequencyMHz() pulumi.Float64Output {
	return o.ApplyT(func(v SpacecraftLinkResponse) float64 { return v.CenterFrequencyMHz }).(pulumi.Float64Output)
}

func (o SpacecraftLinkResponseOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v SpacecraftLinkResponse) string { return v.Direction }).(pulumi.StringOutput)
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





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
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

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
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

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AvailableContactsResponseOutput{})
	pulumi.RegisterOutputType(AvailableContactsResponseArrayOutput{})
	pulumi.RegisterOutputType(ContactProfileLinkOutput{})
	pulumi.RegisterOutputType(ContactProfileLinkArrayOutput{})
	pulumi.RegisterOutputType(ContactProfileLinkChannelOutput{})
	pulumi.RegisterOutputType(ContactProfileLinkChannelArrayOutput{})
	pulumi.RegisterOutputType(ContactProfileLinkChannelResponseOutput{})
	pulumi.RegisterOutputType(ContactProfileLinkChannelResponseArrayOutput{})
	pulumi.RegisterOutputType(ContactProfileLinkResponseOutput{})
	pulumi.RegisterOutputType(ContactProfileLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(EndPointOutput{})
	pulumi.RegisterOutputType(EndPointResponseOutput{})
	pulumi.RegisterOutputType(ResourceReferenceOutput{})
	pulumi.RegisterOutputType(ResourceReferencePtrOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponseOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(SpacecraftLinkOutput{})
	pulumi.RegisterOutputType(SpacecraftLinkArrayOutput{})
	pulumi.RegisterOutputType(SpacecraftLinkResponseOutput{})
	pulumi.RegisterOutputType(SpacecraftLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}

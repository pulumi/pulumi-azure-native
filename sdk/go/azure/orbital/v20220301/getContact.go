


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContact(ctx *pulumi.Context, args *LookupContactArgs, opts ...pulumi.InvokeOption) (*LookupContactResult, error) {
	var rv LookupContactResult
	err := ctx.Invoke("azure-native:orbital/v20220301:getContact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContactArgs struct {
	ContactName       string `pulumi:"contactName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SpacecraftName    string `pulumi:"spacecraftName"`
}


type LookupContactResult struct {
	AntennaConfiguration    ContactsPropertiesResponseAntennaConfiguration `pulumi:"antennaConfiguration"`
	ContactProfile          ContactsPropertiesResponseContactProfile       `pulumi:"contactProfile"`
	EndAzimuthDegrees       float64                                        `pulumi:"endAzimuthDegrees"`
	EndElevationDegrees     float64                                        `pulumi:"endElevationDegrees"`
	ErrorMessage            string                                         `pulumi:"errorMessage"`
	Etag                    string                                         `pulumi:"etag"`
	GroundStationName       string                                         `pulumi:"groundStationName"`
	Id                      string                                         `pulumi:"id"`
	MaximumElevationDegrees float64                                        `pulumi:"maximumElevationDegrees"`
	Name                    string                                         `pulumi:"name"`
	ReservationEndTime      string                                         `pulumi:"reservationEndTime"`
	ReservationStartTime    string                                         `pulumi:"reservationStartTime"`
	RxEndTime               string                                         `pulumi:"rxEndTime"`
	RxStartTime             string                                         `pulumi:"rxStartTime"`
	StartAzimuthDegrees     float64                                        `pulumi:"startAzimuthDegrees"`
	StartElevationDegrees   float64                                        `pulumi:"startElevationDegrees"`
	Status                  string                                         `pulumi:"status"`
	SystemData              SystemDataResponse                             `pulumi:"systemData"`
	TxEndTime               string                                         `pulumi:"txEndTime"`
	TxStartTime             string                                         `pulumi:"txStartTime"`
	Type                    string                                         `pulumi:"type"`
}

func LookupContactOutput(ctx *pulumi.Context, args LookupContactOutputArgs, opts ...pulumi.InvokeOption) LookupContactResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContactResult, error) {
			args := v.(LookupContactArgs)
			r, err := LookupContact(ctx, &args, opts...)
			var s LookupContactResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContactResultOutput)
}

type LookupContactOutputArgs struct {
	ContactName       pulumi.StringInput `pulumi:"contactName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SpacecraftName    pulumi.StringInput `pulumi:"spacecraftName"`
}

func (LookupContactOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContactArgs)(nil)).Elem()
}


type LookupContactResultOutput struct{ *pulumi.OutputState }

func (LookupContactResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContactResult)(nil)).Elem()
}

func (o LookupContactResultOutput) ToLookupContactResultOutput() LookupContactResultOutput {
	return o
}

func (o LookupContactResultOutput) ToLookupContactResultOutputWithContext(ctx context.Context) LookupContactResultOutput {
	return o
}

func (o LookupContactResultOutput) AntennaConfiguration() ContactsPropertiesResponseAntennaConfigurationOutput {
	return o.ApplyT(func(v LookupContactResult) ContactsPropertiesResponseAntennaConfiguration {
		return v.AntennaConfiguration
	}).(ContactsPropertiesResponseAntennaConfigurationOutput)
}

func (o LookupContactResultOutput) ContactProfile() ContactsPropertiesResponseContactProfileOutput {
	return o.ApplyT(func(v LookupContactResult) ContactsPropertiesResponseContactProfile { return v.ContactProfile }).(ContactsPropertiesResponseContactProfileOutput)
}

func (o LookupContactResultOutput) EndAzimuthDegrees() pulumi.Float64Output {
	return o.ApplyT(func(v LookupContactResult) float64 { return v.EndAzimuthDegrees }).(pulumi.Float64Output)
}

func (o LookupContactResultOutput) EndElevationDegrees() pulumi.Float64Output {
	return o.ApplyT(func(v LookupContactResult) float64 { return v.EndElevationDegrees }).(pulumi.Float64Output)
}

func (o LookupContactResultOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o LookupContactResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupContactResultOutput) GroundStationName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.GroundStationName }).(pulumi.StringOutput)
}

func (o LookupContactResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupContactResultOutput) MaximumElevationDegrees() pulumi.Float64Output {
	return o.ApplyT(func(v LookupContactResult) float64 { return v.MaximumElevationDegrees }).(pulumi.Float64Output)
}

func (o LookupContactResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupContactResultOutput) ReservationEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.ReservationEndTime }).(pulumi.StringOutput)
}

func (o LookupContactResultOutput) ReservationStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.ReservationStartTime }).(pulumi.StringOutput)
}

func (o LookupContactResultOutput) RxEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.RxEndTime }).(pulumi.StringOutput)
}

func (o LookupContactResultOutput) RxStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.RxStartTime }).(pulumi.StringOutput)
}

func (o LookupContactResultOutput) StartAzimuthDegrees() pulumi.Float64Output {
	return o.ApplyT(func(v LookupContactResult) float64 { return v.StartAzimuthDegrees }).(pulumi.Float64Output)
}

func (o LookupContactResultOutput) StartElevationDegrees() pulumi.Float64Output {
	return o.ApplyT(func(v LookupContactResult) float64 { return v.StartElevationDegrees }).(pulumi.Float64Output)
}

func (o LookupContactResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupContactResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupContactResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupContactResultOutput) TxEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.TxEndTime }).(pulumi.StringOutput)
}

func (o LookupContactResultOutput) TxStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.TxStartTime }).(pulumi.StringOutput)
}

func (o LookupContactResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContactResultOutput{})
}

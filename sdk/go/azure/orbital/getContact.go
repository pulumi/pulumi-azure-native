


package orbital

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContact(ctx *pulumi.Context, args *LookupContactArgs, opts ...pulumi.InvokeOption) (*LookupContactResult, error) {
	var rv LookupContactResult
	err := ctx.Invoke("azure-native:orbital:getContact", args, &rv, opts...)
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
	ContactProfile          ResourceReferenceResponse `pulumi:"contactProfile"`
	EndAzimuthDegrees       float64                   `pulumi:"endAzimuthDegrees"`
	EndElevationDegrees     float64                   `pulumi:"endElevationDegrees"`
	ErrorMessage            string                    `pulumi:"errorMessage"`
	Etag                    string                    `pulumi:"etag"`
	GroundStationName       string                    `pulumi:"groundStationName"`
	Id                      string                    `pulumi:"id"`
	MaximumElevationDegrees float64                   `pulumi:"maximumElevationDegrees"`
	Name                    string                    `pulumi:"name"`
	ReservationEndTime      string                    `pulumi:"reservationEndTime"`
	ReservationStartTime    string                    `pulumi:"reservationStartTime"`
	RxEndTime               string                    `pulumi:"rxEndTime"`
	RxStartTime             string                    `pulumi:"rxStartTime"`
	StartAzimuthDegrees     float64                   `pulumi:"startAzimuthDegrees"`
	StartElevationDegrees   float64                   `pulumi:"startElevationDegrees"`
	Status                  string                    `pulumi:"status"`
	SystemData              SystemDataResponse        `pulumi:"systemData"`
	TxEndTime               string                    `pulumi:"txEndTime"`
	TxStartTime             string                    `pulumi:"txStartTime"`
	Type                    string                    `pulumi:"type"`
}

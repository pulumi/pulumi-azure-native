


package v20170901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupZone(ctx *pulumi.Context, args *LookupZoneArgs, opts ...pulumi.InvokeOption) (*LookupZoneResult, error) {
	var rv LookupZoneResult
	err := ctx.Invoke("azure-native:network/v20170901:getZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupZoneArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ZoneName          string `pulumi:"zoneName"`
}


type LookupZoneResult struct {
	Etag                           *string           `pulumi:"etag"`
	Id                             string            `pulumi:"id"`
	Location                       string            `pulumi:"location"`
	MaxNumberOfRecordSets          float64           `pulumi:"maxNumberOfRecordSets"`
	MaxNumberOfRecordsPerRecordSet float64           `pulumi:"maxNumberOfRecordsPerRecordSet"`
	Name                           string            `pulumi:"name"`
	NameServers                    []string          `pulumi:"nameServers"`
	NumberOfRecordSets             float64           `pulumi:"numberOfRecordSets"`
	Tags                           map[string]string `pulumi:"tags"`
	Type                           string            `pulumi:"type"`
	ZoneType                       *string           `pulumi:"zoneType"`
}


func (val *LookupZoneResult) Defaults() *LookupZoneResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ZoneType) {
		zoneType_ := "Public"
		tmp.ZoneType = &zoneType_
	}
	return &tmp
}




package v20171001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPacketCapture(ctx *pulumi.Context, args *LookupPacketCaptureArgs, opts ...pulumi.InvokeOption) (*LookupPacketCaptureResult, error) {
	var rv LookupPacketCaptureResult
	err := ctx.Invoke("azure-native:network/v20171001:getPacketCapture", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPacketCaptureArgs struct {
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	PacketCaptureName  string `pulumi:"packetCaptureName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupPacketCaptureResult struct {
	BytesToCapturePerPacket *int                                 `pulumi:"bytesToCapturePerPacket"`
	Etag                    *string                              `pulumi:"etag"`
	Filters                 []PacketCaptureFilterResponse        `pulumi:"filters"`
	Id                      string                               `pulumi:"id"`
	Name                    string                               `pulumi:"name"`
	ProvisioningState       *string                              `pulumi:"provisioningState"`
	StorageLocation         PacketCaptureStorageLocationResponse `pulumi:"storageLocation"`
	Target                  string                               `pulumi:"target"`
	TimeLimitInSeconds      *int                                 `pulumi:"timeLimitInSeconds"`
	TotalBytesPerSession    *int                                 `pulumi:"totalBytesPerSession"`
}


func (val *LookupPacketCaptureResult) Defaults() *LookupPacketCaptureResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.BytesToCapturePerPacket) {
		bytesToCapturePerPacket_ := 0
		tmp.BytesToCapturePerPacket = &bytesToCapturePerPacket_
	}
	if isZero(tmp.Etag) {
		etag_ := "A unique read-only string that changes whenever the resource is updated."
		tmp.Etag = &etag_
	}
	if isZero(tmp.TimeLimitInSeconds) {
		timeLimitInSeconds_ := 18000
		tmp.TimeLimitInSeconds = &timeLimitInSeconds_
	}
	if isZero(tmp.TotalBytesPerSession) {
		totalBytesPerSession_ := 1073741824
		tmp.TotalBytesPerSession = &totalBytesPerSession_
	}
	return &tmp
}

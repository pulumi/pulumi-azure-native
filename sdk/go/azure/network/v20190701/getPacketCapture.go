


package v20190701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPacketCapture(ctx *pulumi.Context, args *LookupPacketCaptureArgs, opts ...pulumi.InvokeOption) (*LookupPacketCaptureResult, error) {
	var rv LookupPacketCaptureResult
	err := ctx.Invoke("azure-native:network/v20190701:getPacketCapture", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
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
	ProvisioningState       string                               `pulumi:"provisioningState"`
	StorageLocation         PacketCaptureStorageLocationResponse `pulumi:"storageLocation"`
	Target                  string                               `pulumi:"target"`
	TimeLimitInSeconds      *int                                 `pulumi:"timeLimitInSeconds"`
	TotalBytesPerSession    *int                                 `pulumi:"totalBytesPerSession"`
}

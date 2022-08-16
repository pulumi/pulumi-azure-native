


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPacketCapture(ctx *pulumi.Context, args *LookupPacketCaptureArgs, opts ...pulumi.InvokeOption) (*LookupPacketCaptureResult, error) {
	var rv LookupPacketCaptureResult
	err := ctx.Invoke("azure-native:network/v20210801:getPacketCapture", args, &rv, opts...)
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
	BytesToCapturePerPacket *float64                             `pulumi:"bytesToCapturePerPacket"`
	Etag                    string                               `pulumi:"etag"`
	Filters                 []PacketCaptureFilterResponse        `pulumi:"filters"`
	Id                      string                               `pulumi:"id"`
	Name                    string                               `pulumi:"name"`
	ProvisioningState       string                               `pulumi:"provisioningState"`
	StorageLocation         PacketCaptureStorageLocationResponse `pulumi:"storageLocation"`
	Target                  string                               `pulumi:"target"`
	TimeLimitInSeconds      *int                                 `pulumi:"timeLimitInSeconds"`
	TotalBytesPerSession    *float64                             `pulumi:"totalBytesPerSession"`
}


func (val *LookupPacketCaptureResult) Defaults() *LookupPacketCaptureResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.BytesToCapturePerPacket) {
		bytesToCapturePerPacket_ := 0.0
		tmp.BytesToCapturePerPacket = &bytesToCapturePerPacket_
	}
	if isZero(tmp.TimeLimitInSeconds) {
		timeLimitInSeconds_ := 18000
		tmp.TimeLimitInSeconds = &timeLimitInSeconds_
	}
	if isZero(tmp.TotalBytesPerSession) {
		totalBytesPerSession_ := 1073741824.0
		tmp.TotalBytesPerSession = &totalBytesPerSession_
	}
	return &tmp
}

func LookupPacketCaptureOutput(ctx *pulumi.Context, args LookupPacketCaptureOutputArgs, opts ...pulumi.InvokeOption) LookupPacketCaptureResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPacketCaptureResult, error) {
			args := v.(LookupPacketCaptureArgs)
			r, err := LookupPacketCapture(ctx, &args, opts...)
			var s LookupPacketCaptureResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPacketCaptureResultOutput)
}

type LookupPacketCaptureOutputArgs struct {
	NetworkWatcherName pulumi.StringInput `pulumi:"networkWatcherName"`
	PacketCaptureName  pulumi.StringInput `pulumi:"packetCaptureName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPacketCaptureOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPacketCaptureArgs)(nil)).Elem()
}


type LookupPacketCaptureResultOutput struct{ *pulumi.OutputState }

func (LookupPacketCaptureResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPacketCaptureResult)(nil)).Elem()
}

func (o LookupPacketCaptureResultOutput) ToLookupPacketCaptureResultOutput() LookupPacketCaptureResultOutput {
	return o
}

func (o LookupPacketCaptureResultOutput) ToLookupPacketCaptureResultOutputWithContext(ctx context.Context) LookupPacketCaptureResultOutput {
	return o
}

func (o LookupPacketCaptureResultOutput) BytesToCapturePerPacket() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) *float64 { return v.BytesToCapturePerPacket }).(pulumi.Float64PtrOutput)
}

func (o LookupPacketCaptureResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupPacketCaptureResultOutput) Filters() PacketCaptureFilterResponseArrayOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) []PacketCaptureFilterResponse { return v.Filters }).(PacketCaptureFilterResponseArrayOutput)
}

func (o LookupPacketCaptureResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPacketCaptureResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPacketCaptureResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPacketCaptureResultOutput) StorageLocation() PacketCaptureStorageLocationResponseOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) PacketCaptureStorageLocationResponse { return v.StorageLocation }).(PacketCaptureStorageLocationResponseOutput)
}

func (o LookupPacketCaptureResultOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) string { return v.Target }).(pulumi.StringOutput)
}

func (o LookupPacketCaptureResultOutput) TimeLimitInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) *int { return v.TimeLimitInSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupPacketCaptureResultOutput) TotalBytesPerSession() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) *float64 { return v.TotalBytesPerSession }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPacketCaptureResultOutput{})
}

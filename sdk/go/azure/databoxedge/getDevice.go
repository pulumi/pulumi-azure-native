


package databoxedge

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDevice(ctx *pulumi.Context, args *LookupDeviceArgs, opts ...pulumi.InvokeOption) (*LookupDeviceResult, error) {
	var rv LookupDeviceResult
	err := ctx.Invoke("azure-native:databoxedge:getDevice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeviceArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDeviceResult struct {
	ConfiguredRoleTypes     []string                    `pulumi:"configuredRoleTypes"`
	Culture                 string                      `pulumi:"culture"`
	DataBoxEdgeDeviceStatus *string                     `pulumi:"dataBoxEdgeDeviceStatus"`
	Description             string                      `pulumi:"description"`
	DeviceHcsVersion        string                      `pulumi:"deviceHcsVersion"`
	DeviceLocalCapacity     float64                     `pulumi:"deviceLocalCapacity"`
	DeviceModel             string                      `pulumi:"deviceModel"`
	DeviceSoftwareVersion   string                      `pulumi:"deviceSoftwareVersion"`
	DeviceType              string                      `pulumi:"deviceType"`
	EdgeProfile             EdgeProfileResponse         `pulumi:"edgeProfile"`
	Etag                    *string                     `pulumi:"etag"`
	FriendlyName            string                      `pulumi:"friendlyName"`
	Id                      string                      `pulumi:"id"`
	Identity                *ResourceIdentityResponse   `pulumi:"identity"`
	Kind                    string                      `pulumi:"kind"`
	Location                string                      `pulumi:"location"`
	ModelDescription        string                      `pulumi:"modelDescription"`
	Name                    string                      `pulumi:"name"`
	NodeCount               int                         `pulumi:"nodeCount"`
	ResourceMoveDetails     ResourceMoveDetailsResponse `pulumi:"resourceMoveDetails"`
	SerialNumber            string                      `pulumi:"serialNumber"`
	Sku                     *SkuResponse                `pulumi:"sku"`
	SystemData              SystemDataResponse          `pulumi:"systemData"`
	Tags                    map[string]string           `pulumi:"tags"`
	TimeZone                string                      `pulumi:"timeZone"`
	Type                    string                      `pulumi:"type"`
}

func LookupDeviceOutput(ctx *pulumi.Context, args LookupDeviceOutputArgs, opts ...pulumi.InvokeOption) LookupDeviceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeviceResult, error) {
			args := v.(LookupDeviceArgs)
			r, err := LookupDevice(ctx, &args, opts...)
			var s LookupDeviceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeviceResultOutput)
}

type LookupDeviceOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDeviceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceArgs)(nil)).Elem()
}


type LookupDeviceResultOutput struct{ *pulumi.OutputState }

func (LookupDeviceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceResult)(nil)).Elem()
}

func (o LookupDeviceResultOutput) ToLookupDeviceResultOutput() LookupDeviceResultOutput {
	return o
}

func (o LookupDeviceResultOutput) ToLookupDeviceResultOutputWithContext(ctx context.Context) LookupDeviceResultOutput {
	return o
}

func (o LookupDeviceResultOutput) ConfiguredRoleTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDeviceResult) []string { return v.ConfiguredRoleTypes }).(pulumi.StringArrayOutput)
}

func (o LookupDeviceResultOutput) Culture() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Culture }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) DataBoxEdgeDeviceStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeviceResult) *string { return v.DataBoxEdgeDeviceStatus }).(pulumi.StringPtrOutput)
}

func (o LookupDeviceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) DeviceHcsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.DeviceHcsVersion }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) DeviceLocalCapacity() pulumi.Float64Output {
	return o.ApplyT(func(v LookupDeviceResult) float64 { return v.DeviceLocalCapacity }).(pulumi.Float64Output)
}

func (o LookupDeviceResultOutput) DeviceModel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.DeviceModel }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) DeviceSoftwareVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.DeviceSoftwareVersion }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) DeviceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.DeviceType }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) EdgeProfile() EdgeProfileResponseOutput {
	return o.ApplyT(func(v LookupDeviceResult) EdgeProfileResponse { return v.EdgeProfile }).(EdgeProfileResponseOutput)
}

func (o LookupDeviceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeviceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupDeviceResultOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.FriendlyName }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupDeviceResult) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o LookupDeviceResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) ModelDescription() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.ModelDescription }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) NodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDeviceResult) int { return v.NodeCount }).(pulumi.IntOutput)
}

func (o LookupDeviceResultOutput) ResourceMoveDetails() ResourceMoveDetailsResponseOutput {
	return o.ApplyT(func(v LookupDeviceResult) ResourceMoveDetailsResponse { return v.ResourceMoveDetails }).(ResourceMoveDetailsResponseOutput)
}

func (o LookupDeviceResultOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.SerialNumber }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupDeviceResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupDeviceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDeviceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDeviceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDeviceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDeviceResultOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.TimeZone }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeviceResultOutput{})
}

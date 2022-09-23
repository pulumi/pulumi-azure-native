


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPacketCoreControlPlane(ctx *pulumi.Context, args *LookupPacketCoreControlPlaneArgs, opts ...pulumi.InvokeOption) (*LookupPacketCoreControlPlaneResult, error) {
	var rv LookupPacketCoreControlPlaneResult
	err := ctx.Invoke("azure-native:mobilenetwork/v20220401preview:getPacketCoreControlPlane", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPacketCoreControlPlaneArgs struct {
	PacketCoreControlPlaneName string `pulumi:"packetCoreControlPlaneName"`
	ResourceGroupName          string `pulumi:"resourceGroupName"`
}


type LookupPacketCoreControlPlaneResult struct {
	ControlPlaneAccessInterface InterfacePropertiesResponse                  `pulumi:"controlPlaneAccessInterface"`
	CoreNetworkTechnology       *string                                      `pulumi:"coreNetworkTechnology"`
	CreatedAt                   *string                                      `pulumi:"createdAt"`
	CreatedBy                   *string                                      `pulumi:"createdBy"`
	CreatedByType               *string                                      `pulumi:"createdByType"`
	Id                          string                                       `pulumi:"id"`
	Identity                    *ManagedServiceIdentityResponse              `pulumi:"identity"`
	InteropSettings             interface{}                                  `pulumi:"interopSettings"`
	LastModifiedAt              *string                                      `pulumi:"lastModifiedAt"`
	LastModifiedBy              *string                                      `pulumi:"lastModifiedBy"`
	LastModifiedByType          *string                                      `pulumi:"lastModifiedByType"`
	LocalDiagnosticsAccess      *LocalDiagnosticsAccessConfigurationResponse `pulumi:"localDiagnosticsAccess"`
	Location                    string                                       `pulumi:"location"`
	MobileNetwork               MobileNetworkResourceIdResponse              `pulumi:"mobileNetwork"`
	Name                        string                                       `pulumi:"name"`
	Platform                    *PlatformConfigurationResponse               `pulumi:"platform"`
	ProvisioningState           string                                       `pulumi:"provisioningState"`
	Sku                         string                                       `pulumi:"sku"`
	SystemData                  SystemDataResponse                           `pulumi:"systemData"`
	Tags                        map[string]string                            `pulumi:"tags"`
	Type                        string                                       `pulumi:"type"`
	Version                     *string                                      `pulumi:"version"`
}

func LookupPacketCoreControlPlaneOutput(ctx *pulumi.Context, args LookupPacketCoreControlPlaneOutputArgs, opts ...pulumi.InvokeOption) LookupPacketCoreControlPlaneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPacketCoreControlPlaneResult, error) {
			args := v.(LookupPacketCoreControlPlaneArgs)
			r, err := LookupPacketCoreControlPlane(ctx, &args, opts...)
			var s LookupPacketCoreControlPlaneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPacketCoreControlPlaneResultOutput)
}

type LookupPacketCoreControlPlaneOutputArgs struct {
	PacketCoreControlPlaneName pulumi.StringInput `pulumi:"packetCoreControlPlaneName"`
	ResourceGroupName          pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPacketCoreControlPlaneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPacketCoreControlPlaneArgs)(nil)).Elem()
}


type LookupPacketCoreControlPlaneResultOutput struct{ *pulumi.OutputState }

func (LookupPacketCoreControlPlaneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPacketCoreControlPlaneResult)(nil)).Elem()
}

func (o LookupPacketCoreControlPlaneResultOutput) ToLookupPacketCoreControlPlaneResultOutput() LookupPacketCoreControlPlaneResultOutput {
	return o
}

func (o LookupPacketCoreControlPlaneResultOutput) ToLookupPacketCoreControlPlaneResultOutputWithContext(ctx context.Context) LookupPacketCoreControlPlaneResultOutput {
	return o
}

func (o LookupPacketCoreControlPlaneResultOutput) ControlPlaneAccessInterface() InterfacePropertiesResponseOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) InterfacePropertiesResponse {
		return v.ControlPlaneAccessInterface
	}).(InterfacePropertiesResponseOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) CoreNetworkTechnology() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) *string { return v.CoreNetworkTechnology }).(pulumi.StringPtrOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) InteropSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) interface{} { return v.InteropSettings }).(pulumi.AnyOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) LocalDiagnosticsAccess() LocalDiagnosticsAccessConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) *LocalDiagnosticsAccessConfigurationResponse {
		return v.LocalDiagnosticsAccess
	}).(LocalDiagnosticsAccessConfigurationResponsePtrOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) MobileNetwork() MobileNetworkResourceIdResponseOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) MobileNetworkResourceIdResponse { return v.MobileNetwork }).(MobileNetworkResourceIdResponseOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) Platform() PlatformConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) *PlatformConfigurationResponse { return v.Platform }).(PlatformConfigurationResponsePtrOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) Sku() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) string { return v.Sku }).(pulumi.StringOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPacketCoreControlPlaneResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPacketCoreControlPlaneResultOutput{})
}

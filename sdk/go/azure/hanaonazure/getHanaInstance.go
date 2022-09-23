


package hanaonazure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHanaInstance(ctx *pulumi.Context, args *LookupHanaInstanceArgs, opts ...pulumi.InvokeOption) (*LookupHanaInstanceResult, error) {
	var rv LookupHanaInstanceResult
	err := ctx.Invoke("azure-native:hanaonazure:getHanaInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHanaInstanceArgs struct {
	HanaInstanceName  string `pulumi:"hanaInstanceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupHanaInstanceResult struct {
	HanaInstanceId          string                   `pulumi:"hanaInstanceId"`
	HardwareProfile         *HardwareProfileResponse `pulumi:"hardwareProfile"`
	HwRevision              string                   `pulumi:"hwRevision"`
	Id                      string                   `pulumi:"id"`
	Location                *string                  `pulumi:"location"`
	Name                    string                   `pulumi:"name"`
	NetworkProfile          *NetworkProfileResponse  `pulumi:"networkProfile"`
	OsProfile               *OSProfileResponse       `pulumi:"osProfile"`
	PartnerNodeId           *string                  `pulumi:"partnerNodeId"`
	PowerState              string                   `pulumi:"powerState"`
	ProvisioningState       string                   `pulumi:"provisioningState"`
	ProximityPlacementGroup string                   `pulumi:"proximityPlacementGroup"`
	StorageProfile          *StorageProfileResponse  `pulumi:"storageProfile"`
	Tags                    map[string]string        `pulumi:"tags"`
	Type                    string                   `pulumi:"type"`
}

func LookupHanaInstanceOutput(ctx *pulumi.Context, args LookupHanaInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupHanaInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHanaInstanceResult, error) {
			args := v.(LookupHanaInstanceArgs)
			r, err := LookupHanaInstance(ctx, &args, opts...)
			var s LookupHanaInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHanaInstanceResultOutput)
}

type LookupHanaInstanceOutputArgs struct {
	HanaInstanceName  pulumi.StringInput `pulumi:"hanaInstanceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupHanaInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHanaInstanceArgs)(nil)).Elem()
}


type LookupHanaInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupHanaInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHanaInstanceResult)(nil)).Elem()
}

func (o LookupHanaInstanceResultOutput) ToLookupHanaInstanceResultOutput() LookupHanaInstanceResultOutput {
	return o
}

func (o LookupHanaInstanceResultOutput) ToLookupHanaInstanceResultOutputWithContext(ctx context.Context) LookupHanaInstanceResultOutput {
	return o
}

func (o LookupHanaInstanceResultOutput) HanaInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHanaInstanceResult) string { return v.HanaInstanceId }).(pulumi.StringOutput)
}

func (o LookupHanaInstanceResultOutput) HardwareProfile() HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupHanaInstanceResult) *HardwareProfileResponse { return v.HardwareProfile }).(HardwareProfileResponsePtrOutput)
}

func (o LookupHanaInstanceResultOutput) HwRevision() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHanaInstanceResult) string { return v.HwRevision }).(pulumi.StringOutput)
}

func (o LookupHanaInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHanaInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHanaInstanceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHanaInstanceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupHanaInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHanaInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupHanaInstanceResultOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupHanaInstanceResult) *NetworkProfileResponse { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o LookupHanaInstanceResultOutput) OsProfile() OSProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupHanaInstanceResult) *OSProfileResponse { return v.OsProfile }).(OSProfileResponsePtrOutput)
}

func (o LookupHanaInstanceResultOutput) PartnerNodeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHanaInstanceResult) *string { return v.PartnerNodeId }).(pulumi.StringPtrOutput)
}

func (o LookupHanaInstanceResultOutput) PowerState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHanaInstanceResult) string { return v.PowerState }).(pulumi.StringOutput)
}

func (o LookupHanaInstanceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHanaInstanceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupHanaInstanceResultOutput) ProximityPlacementGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHanaInstanceResult) string { return v.ProximityPlacementGroup }).(pulumi.StringOutput)
}

func (o LookupHanaInstanceResultOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupHanaInstanceResult) *StorageProfileResponse { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

func (o LookupHanaInstanceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupHanaInstanceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupHanaInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHanaInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHanaInstanceResultOutput{})
}

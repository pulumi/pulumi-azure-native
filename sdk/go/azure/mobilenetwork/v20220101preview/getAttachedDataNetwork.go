


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAttachedDataNetwork(ctx *pulumi.Context, args *LookupAttachedDataNetworkArgs, opts ...pulumi.InvokeOption) (*LookupAttachedDataNetworkResult, error) {
	var rv LookupAttachedDataNetworkResult
	err := ctx.Invoke("azure-native:mobilenetwork/v20220101preview:getAttachedDataNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAttachedDataNetworkArgs struct {
	AttachedDataNetworkName    string `pulumi:"attachedDataNetworkName"`
	PacketCoreControlPlaneName string `pulumi:"packetCoreControlPlaneName"`
	PacketCoreDataPlaneName    string `pulumi:"packetCoreDataPlaneName"`
	ResourceGroupName          string `pulumi:"resourceGroupName"`
}


type LookupAttachedDataNetworkResult struct {
	CreatedAt                            *string                     `pulumi:"createdAt"`
	CreatedBy                            *string                     `pulumi:"createdBy"`
	CreatedByType                        *string                     `pulumi:"createdByType"`
	Id                                   string                      `pulumi:"id"`
	LastModifiedAt                       *string                     `pulumi:"lastModifiedAt"`
	LastModifiedBy                       *string                     `pulumi:"lastModifiedBy"`
	LastModifiedByType                   *string                     `pulumi:"lastModifiedByType"`
	Location                             string                      `pulumi:"location"`
	Name                                 string                      `pulumi:"name"`
	NaptConfiguration                    *NaptConfigurationResponse  `pulumi:"naptConfiguration"`
	ProvisioningState                    string                      `pulumi:"provisioningState"`
	Tags                                 map[string]string           `pulumi:"tags"`
	Type                                 string                      `pulumi:"type"`
	UserEquipmentAddressPoolPrefix       []string                    `pulumi:"userEquipmentAddressPoolPrefix"`
	UserEquipmentStaticAddressPoolPrefix []string                    `pulumi:"userEquipmentStaticAddressPoolPrefix"`
	UserPlaneDataInterface               InterfacePropertiesResponse `pulumi:"userPlaneDataInterface"`
}


func (val *LookupAttachedDataNetworkResult) Defaults() *LookupAttachedDataNetworkResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.NaptConfiguration = tmp.NaptConfiguration.Defaults()

	return &tmp
}

func LookupAttachedDataNetworkOutput(ctx *pulumi.Context, args LookupAttachedDataNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupAttachedDataNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAttachedDataNetworkResult, error) {
			args := v.(LookupAttachedDataNetworkArgs)
			r, err := LookupAttachedDataNetwork(ctx, &args, opts...)
			var s LookupAttachedDataNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAttachedDataNetworkResultOutput)
}

type LookupAttachedDataNetworkOutputArgs struct {
	AttachedDataNetworkName    pulumi.StringInput `pulumi:"attachedDataNetworkName"`
	PacketCoreControlPlaneName pulumi.StringInput `pulumi:"packetCoreControlPlaneName"`
	PacketCoreDataPlaneName    pulumi.StringInput `pulumi:"packetCoreDataPlaneName"`
	ResourceGroupName          pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAttachedDataNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttachedDataNetworkArgs)(nil)).Elem()
}


type LookupAttachedDataNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupAttachedDataNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttachedDataNetworkResult)(nil)).Elem()
}

func (o LookupAttachedDataNetworkResultOutput) ToLookupAttachedDataNetworkResultOutput() LookupAttachedDataNetworkResultOutput {
	return o
}

func (o LookupAttachedDataNetworkResultOutput) ToLookupAttachedDataNetworkResultOutputWithContext(ctx context.Context) LookupAttachedDataNetworkResultOutput {
	return o
}

func (o LookupAttachedDataNetworkResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupAttachedDataNetworkResultOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o LookupAttachedDataNetworkResultOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o LookupAttachedDataNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAttachedDataNetworkResultOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o LookupAttachedDataNetworkResultOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o LookupAttachedDataNetworkResultOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o LookupAttachedDataNetworkResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAttachedDataNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAttachedDataNetworkResultOutput) NaptConfiguration() NaptConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) *NaptConfigurationResponse { return v.NaptConfiguration }).(NaptConfigurationResponsePtrOutput)
}

func (o LookupAttachedDataNetworkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAttachedDataNetworkResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAttachedDataNetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAttachedDataNetworkResultOutput) UserEquipmentAddressPoolPrefix() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) []string { return v.UserEquipmentAddressPoolPrefix }).(pulumi.StringArrayOutput)
}

func (o LookupAttachedDataNetworkResultOutput) UserEquipmentStaticAddressPoolPrefix() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) []string { return v.UserEquipmentStaticAddressPoolPrefix }).(pulumi.StringArrayOutput)
}

func (o LookupAttachedDataNetworkResultOutput) UserPlaneDataInterface() InterfacePropertiesResponseOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) InterfacePropertiesResponse { return v.UserPlaneDataInterface }).(InterfacePropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAttachedDataNetworkResultOutput{})
}

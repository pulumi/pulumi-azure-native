


package fluidrelay

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FluidRelayEndpointsResponse struct {
	OrdererEndpoints []string `pulumi:"ordererEndpoints"`
	StorageEndpoints []string `pulumi:"storageEndpoints"`
}

type FluidRelayEndpointsResponseOutput struct{ *pulumi.OutputState }

func (FluidRelayEndpointsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FluidRelayEndpointsResponse)(nil)).Elem()
}

func (o FluidRelayEndpointsResponseOutput) ToFluidRelayEndpointsResponseOutput() FluidRelayEndpointsResponseOutput {
	return o
}

func (o FluidRelayEndpointsResponseOutput) ToFluidRelayEndpointsResponseOutputWithContext(ctx context.Context) FluidRelayEndpointsResponseOutput {
	return o
}

func (o FluidRelayEndpointsResponseOutput) OrdererEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FluidRelayEndpointsResponse) []string { return v.OrdererEndpoints }).(pulumi.StringArrayOutput)
}

func (o FluidRelayEndpointsResponseOutput) StorageEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FluidRelayEndpointsResponse) []string { return v.StorageEndpoints }).(pulumi.StringArrayOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(FluidRelayEndpointsResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}

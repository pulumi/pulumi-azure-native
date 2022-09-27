


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetwork(ctx *pulumi.Context, args *LookupVirtualNetworkArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkResult, error) {
	var rv LookupVirtualNetworkResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere/v20201001preview:getVirtualNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkArgs struct {
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	VirtualNetworkName string `pulumi:"virtualNetworkName"`
}


type LookupVirtualNetworkResult struct {
	CustomResourceName string                    `pulumi:"customResourceName"`
	ExtendedLocation   *ExtendedLocationResponse `pulumi:"extendedLocation"`
	Id                 string                    `pulumi:"id"`
	InventoryItemId    *string                   `pulumi:"inventoryItemId"`
	Kind               *string                   `pulumi:"kind"`
	Location           string                    `pulumi:"location"`
	MoName             string                    `pulumi:"moName"`
	MoRefId            *string                   `pulumi:"moRefId"`
	Name               string                    `pulumi:"name"`
	ProvisioningState  string                    `pulumi:"provisioningState"`
	Statuses           []ResourceStatusResponse  `pulumi:"statuses"`
	SystemData         SystemDataResponse        `pulumi:"systemData"`
	Tags               map[string]string         `pulumi:"tags"`
	Type               string                    `pulumi:"type"`
	Uuid               string                    `pulumi:"uuid"`
	VCenterId          *string                   `pulumi:"vCenterId"`
}

func LookupVirtualNetworkOutput(ctx *pulumi.Context, args LookupVirtualNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNetworkResult, error) {
			args := v.(LookupVirtualNetworkArgs)
			r, err := LookupVirtualNetwork(ctx, &args, opts...)
			var s LookupVirtualNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualNetworkResultOutput)
}

type LookupVirtualNetworkOutputArgs struct {
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualNetworkName pulumi.StringInput `pulumi:"virtualNetworkName"`
}

func (LookupVirtualNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkArgs)(nil)).Elem()
}


type LookupVirtualNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkResult)(nil)).Elem()
}

func (o LookupVirtualNetworkResultOutput) ToLookupVirtualNetworkResultOutput() LookupVirtualNetworkResultOutput {
	return o
}

func (o LookupVirtualNetworkResultOutput) ToLookupVirtualNetworkResultOutputWithContext(ctx context.Context) LookupVirtualNetworkResultOutput {
	return o
}

func (o LookupVirtualNetworkResultOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.CustomResourceName }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupVirtualNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkResultOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *string { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkResultOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.MoName }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkResultOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *string { return v.MoRefId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkResultOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) []ResourceStatusResponse { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

func (o LookupVirtualNetworkResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupVirtualNetworkResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualNetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkResultOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *string { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkResultOutput{})
}

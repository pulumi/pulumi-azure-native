


package v20200605preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetwork(ctx *pulumi.Context, args *LookupVirtualNetworkArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkResult, error) {
	var rv LookupVirtualNetworkResult
	err := ctx.Invoke("azure-native:scvmm/v20200605preview:getVirtualNetwork", args, &rv, opts...)
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
	ExtendedLocation  ExtendedLocationResponse `pulumi:"extendedLocation"`
	Id                string                   `pulumi:"id"`
	InventoryItemId   *string                  `pulumi:"inventoryItemId"`
	Location          string                   `pulumi:"location"`
	Name              string                   `pulumi:"name"`
	NetworkName       string                   `pulumi:"networkName"`
	ProvisioningState string                   `pulumi:"provisioningState"`
	SystemData        SystemDataResponse       `pulumi:"systemData"`
	Tags              map[string]string        `pulumi:"tags"`
	Type              string                   `pulumi:"type"`
	Uuid              *string                  `pulumi:"uuid"`
	VmmServerId       *string                  `pulumi:"vmmServerId"`
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

func (o LookupVirtualNetworkResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

func (o LookupVirtualNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkResultOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *string { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkResultOutput) NetworkName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.NetworkName }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
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

func (o LookupVirtualNetworkResultOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *string { return v.Uuid }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkResultOutput) VmmServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *string { return v.VmmServerId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkResultOutput{})
}

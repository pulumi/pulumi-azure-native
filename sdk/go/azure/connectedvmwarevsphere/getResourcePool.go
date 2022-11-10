


package connectedvmwarevsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupResourcePool(ctx *pulumi.Context, args *LookupResourcePoolArgs, opts ...pulumi.InvokeOption) (*LookupResourcePoolResult, error) {
	var rv LookupResourcePoolResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere:getResourcePool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourcePoolArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourcePoolName  string `pulumi:"resourcePoolName"`
}


type LookupResourcePoolResult struct {
	CpuLimitMHz        float64                   `pulumi:"cpuLimitMHz"`
	CpuReservationMHz  float64                   `pulumi:"cpuReservationMHz"`
	CpuSharesLevel     string                    `pulumi:"cpuSharesLevel"`
	CustomResourceName string                    `pulumi:"customResourceName"`
	ExtendedLocation   *ExtendedLocationResponse `pulumi:"extendedLocation"`
	Id                 string                    `pulumi:"id"`
	InventoryItemId    *string                   `pulumi:"inventoryItemId"`
	Kind               *string                   `pulumi:"kind"`
	Location           string                    `pulumi:"location"`
	MemLimitMB         float64                   `pulumi:"memLimitMB"`
	MemReservationMB   float64                   `pulumi:"memReservationMB"`
	MemSharesLevel     string                    `pulumi:"memSharesLevel"`
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

func LookupResourcePoolOutput(ctx *pulumi.Context, args LookupResourcePoolOutputArgs, opts ...pulumi.InvokeOption) LookupResourcePoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourcePoolResult, error) {
			args := v.(LookupResourcePoolArgs)
			r, err := LookupResourcePool(ctx, &args, opts...)
			var s LookupResourcePoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourcePoolResultOutput)
}

type LookupResourcePoolOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourcePoolName  pulumi.StringInput `pulumi:"resourcePoolName"`
}

func (LookupResourcePoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourcePoolArgs)(nil)).Elem()
}


type LookupResourcePoolResultOutput struct{ *pulumi.OutputState }

func (LookupResourcePoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourcePoolResult)(nil)).Elem()
}

func (o LookupResourcePoolResultOutput) ToLookupResourcePoolResultOutput() LookupResourcePoolResultOutput {
	return o
}

func (o LookupResourcePoolResultOutput) ToLookupResourcePoolResultOutputWithContext(ctx context.Context) LookupResourcePoolResultOutput {
	return o
}

func (o LookupResourcePoolResultOutput) CpuLimitMHz() pulumi.Float64Output {
	return o.ApplyT(func(v LookupResourcePoolResult) float64 { return v.CpuLimitMHz }).(pulumi.Float64Output)
}

func (o LookupResourcePoolResultOutput) CpuReservationMHz() pulumi.Float64Output {
	return o.ApplyT(func(v LookupResourcePoolResult) float64 { return v.CpuReservationMHz }).(pulumi.Float64Output)
}

func (o LookupResourcePoolResultOutput) CpuSharesLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.CpuSharesLevel }).(pulumi.StringOutput)
}

func (o LookupResourcePoolResultOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.CustomResourceName }).(pulumi.StringOutput)
}

func (o LookupResourcePoolResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupResourcePoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupResourcePoolResultOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) *string { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

func (o LookupResourcePoolResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupResourcePoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupResourcePoolResultOutput) MemLimitMB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupResourcePoolResult) float64 { return v.MemLimitMB }).(pulumi.Float64Output)
}

func (o LookupResourcePoolResultOutput) MemReservationMB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupResourcePoolResult) float64 { return v.MemReservationMB }).(pulumi.Float64Output)
}

func (o LookupResourcePoolResultOutput) MemSharesLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.MemSharesLevel }).(pulumi.StringOutput)
}

func (o LookupResourcePoolResultOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.MoName }).(pulumi.StringOutput)
}

func (o LookupResourcePoolResultOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) *string { return v.MoRefId }).(pulumi.StringPtrOutput)
}

func (o LookupResourcePoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupResourcePoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupResourcePoolResultOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) []ResourceStatusResponse { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

func (o LookupResourcePoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupResourcePoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupResourcePoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupResourcePoolResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupResourcePoolResultOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) *string { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourcePoolResultOutput{})
}

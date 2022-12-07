


package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCapacityReservation(ctx *pulumi.Context, args *LookupCapacityReservationArgs, opts ...pulumi.InvokeOption) (*LookupCapacityReservationResult, error) {
	var rv LookupCapacityReservationResult
	err := ctx.Invoke("azure-native:compute/v20211101:getCapacityReservation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCapacityReservationArgs struct {
	CapacityReservationGroupName string  `pulumi:"capacityReservationGroupName"`
	CapacityReservationName      string  `pulumi:"capacityReservationName"`
	Expand                       *string `pulumi:"expand"`
	ResourceGroupName            string  `pulumi:"resourceGroupName"`
}


type LookupCapacityReservationResult struct {
	Id                        string                                  `pulumi:"id"`
	InstanceView              CapacityReservationInstanceViewResponse `pulumi:"instanceView"`
	Location                  string                                  `pulumi:"location"`
	Name                      string                                  `pulumi:"name"`
	ProvisioningState         string                                  `pulumi:"provisioningState"`
	ProvisioningTime          string                                  `pulumi:"provisioningTime"`
	ReservationId             string                                  `pulumi:"reservationId"`
	Sku                       SkuResponse                             `pulumi:"sku"`
	Tags                      map[string]string                       `pulumi:"tags"`
	TimeCreated               string                                  `pulumi:"timeCreated"`
	Type                      string                                  `pulumi:"type"`
	VirtualMachinesAssociated []SubResourceReadOnlyResponse           `pulumi:"virtualMachinesAssociated"`
	Zones                     []string                                `pulumi:"zones"`
}

func LookupCapacityReservationOutput(ctx *pulumi.Context, args LookupCapacityReservationOutputArgs, opts ...pulumi.InvokeOption) LookupCapacityReservationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCapacityReservationResult, error) {
			args := v.(LookupCapacityReservationArgs)
			r, err := LookupCapacityReservation(ctx, &args, opts...)
			var s LookupCapacityReservationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCapacityReservationResultOutput)
}

type LookupCapacityReservationOutputArgs struct {
	CapacityReservationGroupName pulumi.StringInput    `pulumi:"capacityReservationGroupName"`
	CapacityReservationName      pulumi.StringInput    `pulumi:"capacityReservationName"`
	Expand                       pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName            pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupCapacityReservationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityReservationArgs)(nil)).Elem()
}


type LookupCapacityReservationResultOutput struct{ *pulumi.OutputState }

func (LookupCapacityReservationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityReservationResult)(nil)).Elem()
}

func (o LookupCapacityReservationResultOutput) ToLookupCapacityReservationResultOutput() LookupCapacityReservationResultOutput {
	return o
}

func (o LookupCapacityReservationResultOutput) ToLookupCapacityReservationResultOutputWithContext(ctx context.Context) LookupCapacityReservationResultOutput {
	return o
}

func (o LookupCapacityReservationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCapacityReservationResultOutput) InstanceView() CapacityReservationInstanceViewResponseOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) CapacityReservationInstanceViewResponse { return v.InstanceView }).(CapacityReservationInstanceViewResponseOutput)
}

func (o LookupCapacityReservationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupCapacityReservationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCapacityReservationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupCapacityReservationResultOutput) ProvisioningTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) string { return v.ProvisioningTime }).(pulumi.StringOutput)
}

func (o LookupCapacityReservationResultOutput) ReservationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) string { return v.ReservationId }).(pulumi.StringOutput)
}

func (o LookupCapacityReservationResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupCapacityReservationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCapacityReservationResultOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) string { return v.TimeCreated }).(pulumi.StringOutput)
}

func (o LookupCapacityReservationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupCapacityReservationResultOutput) VirtualMachinesAssociated() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) []SubResourceReadOnlyResponse {
		return v.VirtualMachinesAssociated
	}).(SubResourceReadOnlyResponseArrayOutput)
}

func (o LookupCapacityReservationResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCapacityReservationResultOutput{})
}

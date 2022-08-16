


package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCapacityReservationGroup(ctx *pulumi.Context, args *LookupCapacityReservationGroupArgs, opts ...pulumi.InvokeOption) (*LookupCapacityReservationGroupResult, error) {
	var rv LookupCapacityReservationGroupResult
	err := ctx.Invoke("azure-native:compute:getCapacityReservationGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCapacityReservationGroupArgs struct {
	CapacityReservationGroupName string  `pulumi:"capacityReservationGroupName"`
	Expand                       *string `pulumi:"expand"`
	ResourceGroupName            string  `pulumi:"resourceGroupName"`
}


type LookupCapacityReservationGroupResult struct {
	CapacityReservations      []SubResourceReadOnlyResponse                `pulumi:"capacityReservations"`
	Id                        string                                       `pulumi:"id"`
	InstanceView              CapacityReservationGroupInstanceViewResponse `pulumi:"instanceView"`
	Location                  string                                       `pulumi:"location"`
	Name                      string                                       `pulumi:"name"`
	Tags                      map[string]string                            `pulumi:"tags"`
	Type                      string                                       `pulumi:"type"`
	VirtualMachinesAssociated []SubResourceReadOnlyResponse                `pulumi:"virtualMachinesAssociated"`
	Zones                     []string                                     `pulumi:"zones"`
}

func LookupCapacityReservationGroupOutput(ctx *pulumi.Context, args LookupCapacityReservationGroupOutputArgs, opts ...pulumi.InvokeOption) LookupCapacityReservationGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCapacityReservationGroupResult, error) {
			args := v.(LookupCapacityReservationGroupArgs)
			r, err := LookupCapacityReservationGroup(ctx, &args, opts...)
			var s LookupCapacityReservationGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCapacityReservationGroupResultOutput)
}

type LookupCapacityReservationGroupOutputArgs struct {
	CapacityReservationGroupName pulumi.StringInput    `pulumi:"capacityReservationGroupName"`
	Expand                       pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName            pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupCapacityReservationGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityReservationGroupArgs)(nil)).Elem()
}


type LookupCapacityReservationGroupResultOutput struct{ *pulumi.OutputState }

func (LookupCapacityReservationGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityReservationGroupResult)(nil)).Elem()
}

func (o LookupCapacityReservationGroupResultOutput) ToLookupCapacityReservationGroupResultOutput() LookupCapacityReservationGroupResultOutput {
	return o
}

func (o LookupCapacityReservationGroupResultOutput) ToLookupCapacityReservationGroupResultOutputWithContext(ctx context.Context) LookupCapacityReservationGroupResultOutput {
	return o
}

func (o LookupCapacityReservationGroupResultOutput) CapacityReservations() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) []SubResourceReadOnlyResponse {
		return v.CapacityReservations
	}).(SubResourceReadOnlyResponseArrayOutput)
}

func (o LookupCapacityReservationGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCapacityReservationGroupResultOutput) InstanceView() CapacityReservationGroupInstanceViewResponseOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) CapacityReservationGroupInstanceViewResponse {
		return v.InstanceView
	}).(CapacityReservationGroupInstanceViewResponseOutput)
}

func (o LookupCapacityReservationGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupCapacityReservationGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCapacityReservationGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCapacityReservationGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupCapacityReservationGroupResultOutput) VirtualMachinesAssociated() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) []SubResourceReadOnlyResponse {
		return v.VirtualMachinesAssociated
	}).(SubResourceReadOnlyResponseArrayOutput)
}

func (o LookupCapacityReservationGroupResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCapacityReservationGroupResultOutput{})
}

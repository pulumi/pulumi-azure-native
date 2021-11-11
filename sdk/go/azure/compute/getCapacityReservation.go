


package compute

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCapacityReservation(ctx *pulumi.Context, args *LookupCapacityReservationArgs, opts ...pulumi.InvokeOption) (*LookupCapacityReservationResult, error) {
	var rv LookupCapacityReservationResult
	err := ctx.Invoke("azure-native:compute:getCapacityReservation", args, &rv, opts...)
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
	Type                      string                                  `pulumi:"type"`
	VirtualMachinesAssociated []SubResourceReadOnlyResponse           `pulumi:"virtualMachinesAssociated"`
	Zones                     []string                                `pulumi:"zones"`
}

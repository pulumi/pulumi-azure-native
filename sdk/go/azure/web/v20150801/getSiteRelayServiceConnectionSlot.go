


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSiteRelayServiceConnectionSlot(ctx *pulumi.Context, args *LookupSiteRelayServiceConnectionSlotArgs, opts ...pulumi.InvokeOption) (*LookupSiteRelayServiceConnectionSlotResult, error) {
	var rv LookupSiteRelayServiceConnectionSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteRelayServiceConnectionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteRelayServiceConnectionSlotArgs struct {
	EntityName        string `pulumi:"entityName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupSiteRelayServiceConnectionSlotResult struct {
	BiztalkUri               *string           `pulumi:"biztalkUri"`
	EntityConnectionString   *string           `pulumi:"entityConnectionString"`
	EntityName               *string           `pulumi:"entityName"`
	Hostname                 *string           `pulumi:"hostname"`
	Id                       *string           `pulumi:"id"`
	Kind                     *string           `pulumi:"kind"`
	Location                 string            `pulumi:"location"`
	Name                     *string           `pulumi:"name"`
	Port                     *int              `pulumi:"port"`
	ResourceConnectionString *string           `pulumi:"resourceConnectionString"`
	ResourceType             *string           `pulumi:"resourceType"`
	Tags                     map[string]string `pulumi:"tags"`
	Type                     *string           `pulumi:"type"`
}

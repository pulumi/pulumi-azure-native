


package v20201001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppRelayServiceConnectionSlot(ctx *pulumi.Context, args *LookupWebAppRelayServiceConnectionSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppRelayServiceConnectionSlotResult, error) {
	var rv LookupWebAppRelayServiceConnectionSlotResult
	err := ctx.Invoke("azure-native:web/v20201001:getWebAppRelayServiceConnectionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppRelayServiceConnectionSlotArgs struct {
	EntityName        string `pulumi:"entityName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupWebAppRelayServiceConnectionSlotResult struct {
	BiztalkUri               *string            `pulumi:"biztalkUri"`
	EntityConnectionString   *string            `pulumi:"entityConnectionString"`
	EntityName               *string            `pulumi:"entityName"`
	Hostname                 *string            `pulumi:"hostname"`
	Id                       string             `pulumi:"id"`
	Kind                     *string            `pulumi:"kind"`
	Name                     string             `pulumi:"name"`
	Port                     *int               `pulumi:"port"`
	ResourceConnectionString *string            `pulumi:"resourceConnectionString"`
	ResourceType             *string            `pulumi:"resourceType"`
	SystemData               SystemDataResponse `pulumi:"systemData"`
	Type                     string             `pulumi:"type"`
}

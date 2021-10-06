


package v20210101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDisasterRecoveryConfig(ctx *pulumi.Context, args *LookupDisasterRecoveryConfigArgs, opts ...pulumi.InvokeOption) (*LookupDisasterRecoveryConfigResult, error) {
	var rv LookupDisasterRecoveryConfigResult
	err := ctx.Invoke("azure-native:servicebus/v20210101preview:getDisasterRecoveryConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDisasterRecoveryConfigArgs struct {
	Alias             string `pulumi:"alias"`
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDisasterRecoveryConfigResult struct {
	AlternateName                     *string            `pulumi:"alternateName"`
	Id                                string             `pulumi:"id"`
	Name                              string             `pulumi:"name"`
	PartnerNamespace                  *string            `pulumi:"partnerNamespace"`
	PendingReplicationOperationsCount float64            `pulumi:"pendingReplicationOperationsCount"`
	ProvisioningState                 string             `pulumi:"provisioningState"`
	Role                              string             `pulumi:"role"`
	SystemData                        SystemDataResponse `pulumi:"systemData"`
	Type                              string             `pulumi:"type"`
}

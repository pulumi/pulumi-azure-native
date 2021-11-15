


package v20210401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplicationRecoveryPlan(ctx *pulumi.Context, args *LookupReplicationRecoveryPlanArgs, opts ...pulumi.InvokeOption) (*LookupReplicationRecoveryPlanResult, error) {
	var rv LookupReplicationRecoveryPlanResult
	err := ctx.Invoke("azure-native:recoveryservices/v20210401:getReplicationRecoveryPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationRecoveryPlanArgs struct {
	RecoveryPlanName  string `pulumi:"recoveryPlanName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupReplicationRecoveryPlanResult struct {
	Id         string                         `pulumi:"id"`
	Location   *string                        `pulumi:"location"`
	Name       string                         `pulumi:"name"`
	Properties RecoveryPlanPropertiesResponse `pulumi:"properties"`
	Type       string                         `pulumi:"type"`
}

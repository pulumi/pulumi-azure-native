


package recoveryservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplicationRecoveryServicesProvider(ctx *pulumi.Context, args *LookupReplicationRecoveryServicesProviderArgs, opts ...pulumi.InvokeOption) (*LookupReplicationRecoveryServicesProviderResult, error) {
	var rv LookupReplicationRecoveryServicesProviderResult
	err := ctx.Invoke("azure-native:recoveryservices:getReplicationRecoveryServicesProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationRecoveryServicesProviderArgs struct {
	FabricName        string `pulumi:"fabricName"`
	ProviderName      string `pulumi:"providerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupReplicationRecoveryServicesProviderResult struct {
	Id         string                                     `pulumi:"id"`
	Location   *string                                    `pulumi:"location"`
	Name       string                                     `pulumi:"name"`
	Properties RecoveryServicesProviderPropertiesResponse `pulumi:"properties"`
	Type       string                                     `pulumi:"type"`
}

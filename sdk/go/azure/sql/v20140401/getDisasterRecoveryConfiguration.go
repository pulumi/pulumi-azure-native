


package v20140401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDisasterRecoveryConfiguration(ctx *pulumi.Context, args *LookupDisasterRecoveryConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupDisasterRecoveryConfigurationResult, error) {
	var rv LookupDisasterRecoveryConfigurationResult
	err := ctx.Invoke("azure-native:sql/v20140401:getDisasterRecoveryConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDisasterRecoveryConfigurationArgs struct {
	DisasterRecoveryConfigurationName string `pulumi:"disasterRecoveryConfigurationName"`
	ResourceGroupName                 string `pulumi:"resourceGroupName"`
	ServerName                        string `pulumi:"serverName"`
}


type LookupDisasterRecoveryConfigurationResult struct {
	AutoFailover             string `pulumi:"autoFailover"`
	FailoverPolicy           string `pulumi:"failoverPolicy"`
	Id                       string `pulumi:"id"`
	Location                 string `pulumi:"location"`
	LogicalServerName        string `pulumi:"logicalServerName"`
	Name                     string `pulumi:"name"`
	PartnerLogicalServerName string `pulumi:"partnerLogicalServerName"`
	PartnerServerId          string `pulumi:"partnerServerId"`
	Role                     string `pulumi:"role"`
	Status                   string `pulumi:"status"`
	Type                     string `pulumi:"type"`
}




package sql

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataMaskingPolicy(ctx *pulumi.Context, args *LookupDataMaskingPolicyArgs, opts ...pulumi.InvokeOption) (*LookupDataMaskingPolicyResult, error) {
	var rv LookupDataMaskingPolicyResult
	err := ctx.Invoke("azure-native:sql:getDataMaskingPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataMaskingPolicyArgs struct {
	DataMaskingPolicyName string `pulumi:"dataMaskingPolicyName"`
	DatabaseName          string `pulumi:"databaseName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ServerName            string `pulumi:"serverName"`
}


type LookupDataMaskingPolicyResult struct {
	ApplicationPrincipals string  `pulumi:"applicationPrincipals"`
	DataMaskingState      string  `pulumi:"dataMaskingState"`
	ExemptPrincipals      *string `pulumi:"exemptPrincipals"`
	Id                    string  `pulumi:"id"`
	Kind                  string  `pulumi:"kind"`
	Location              string  `pulumi:"location"`
	MaskingLevel          string  `pulumi:"maskingLevel"`
	Name                  string  `pulumi:"name"`
	Type                  string  `pulumi:"type"`
}

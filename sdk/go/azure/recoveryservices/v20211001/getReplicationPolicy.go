


package v20211001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplicationPolicy(ctx *pulumi.Context, args *LookupReplicationPolicyArgs, opts ...pulumi.InvokeOption) (*LookupReplicationPolicyResult, error) {
	var rv LookupReplicationPolicyResult
	err := ctx.Invoke("azure-native:recoveryservices/v20211001:getReplicationPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationPolicyArgs struct {
	PolicyName        string `pulumi:"policyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupReplicationPolicyResult struct {
	Id         string                   `pulumi:"id"`
	Location   *string                  `pulumi:"location"`
	Name       string                   `pulumi:"name"`
	Properties PolicyPropertiesResponse `pulumi:"properties"`
	Type       string                   `pulumi:"type"`
}

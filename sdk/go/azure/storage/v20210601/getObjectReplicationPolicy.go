


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupObjectReplicationPolicy(ctx *pulumi.Context, args *LookupObjectReplicationPolicyArgs, opts ...pulumi.InvokeOption) (*LookupObjectReplicationPolicyResult, error) {
	var rv LookupObjectReplicationPolicyResult
	err := ctx.Invoke("azure-native:storage/v20210601:getObjectReplicationPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupObjectReplicationPolicyArgs struct {
	AccountName               string `pulumi:"accountName"`
	ObjectReplicationPolicyId string `pulumi:"objectReplicationPolicyId"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
}


type LookupObjectReplicationPolicyResult struct {
	DestinationAccount string                                `pulumi:"destinationAccount"`
	EnabledTime        string                                `pulumi:"enabledTime"`
	Id                 string                                `pulumi:"id"`
	Name               string                                `pulumi:"name"`
	PolicyId           string                                `pulumi:"policyId"`
	Rules              []ObjectReplicationPolicyRuleResponse `pulumi:"rules"`
	SourceAccount      string                                `pulumi:"sourceAccount"`
	Type               string                                `pulumi:"type"`
}

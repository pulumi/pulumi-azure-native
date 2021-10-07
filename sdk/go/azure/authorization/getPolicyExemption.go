


package authorization

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicyExemption(ctx *pulumi.Context, args *LookupPolicyExemptionArgs, opts ...pulumi.InvokeOption) (*LookupPolicyExemptionResult, error) {
	var rv LookupPolicyExemptionResult
	err := ctx.Invoke("azure-native:authorization:getPolicyExemption", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyExemptionArgs struct {
	PolicyExemptionName string `pulumi:"policyExemptionName"`
	Scope               string `pulumi:"scope"`
}


type LookupPolicyExemptionResult struct {
	Description                  *string            `pulumi:"description"`
	DisplayName                  *string            `pulumi:"displayName"`
	ExemptionCategory            string             `pulumi:"exemptionCategory"`
	ExpiresOn                    *string            `pulumi:"expiresOn"`
	Id                           string             `pulumi:"id"`
	Metadata                     interface{}        `pulumi:"metadata"`
	Name                         string             `pulumi:"name"`
	PolicyAssignmentId           string             `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceIds []string           `pulumi:"policyDefinitionReferenceIds"`
	SystemData                   SystemDataResponse `pulumi:"systemData"`
	Type                         string             `pulumi:"type"`
}

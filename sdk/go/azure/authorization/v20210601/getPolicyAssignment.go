


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicyAssignment(ctx *pulumi.Context, args *LookupPolicyAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupPolicyAssignmentResult, error) {
	var rv LookupPolicyAssignmentResult
	err := ctx.Invoke("azure-native:authorization/v20210601:getPolicyAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPolicyAssignmentArgs struct {
	PolicyAssignmentName string `pulumi:"policyAssignmentName"`
	Scope                string `pulumi:"scope"`
}


type LookupPolicyAssignmentResult struct {
	Description           *string                                 `pulumi:"description"`
	DisplayName           *string                                 `pulumi:"displayName"`
	EnforcementMode       *string                                 `pulumi:"enforcementMode"`
	Id                    string                                  `pulumi:"id"`
	Identity              *IdentityResponse                       `pulumi:"identity"`
	Location              *string                                 `pulumi:"location"`
	Metadata              interface{}                             `pulumi:"metadata"`
	Name                  string                                  `pulumi:"name"`
	NonComplianceMessages []NonComplianceMessageResponse          `pulumi:"nonComplianceMessages"`
	NotScopes             []string                                `pulumi:"notScopes"`
	Parameters            map[string]ParameterValuesValueResponse `pulumi:"parameters"`
	PolicyDefinitionId    *string                                 `pulumi:"policyDefinitionId"`
	Scope                 string                                  `pulumi:"scope"`
	SystemData            SystemDataResponse                      `pulumi:"systemData"`
	Type                  string                                  `pulumi:"type"`
}


func (val *LookupPolicyAssignmentResult) Defaults() *LookupPolicyAssignmentResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnforcementMode) {
		enforcementMode_ := "Default"
		tmp.EnforcementMode = &enforcementMode_
	}
	return &tmp
}

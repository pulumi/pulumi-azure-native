


package v20210101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAttestationAtResourceGroup(ctx *pulumi.Context, args *LookupAttestationAtResourceGroupArgs, opts ...pulumi.InvokeOption) (*LookupAttestationAtResourceGroupResult, error) {
	var rv LookupAttestationAtResourceGroupResult
	err := ctx.Invoke("azure-native:policyinsights/v20210101:getAttestationAtResourceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAttestationAtResourceGroupArgs struct {
	AttestationName   string `pulumi:"attestationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAttestationAtResourceGroupResult struct {
	Comments                    *string                       `pulumi:"comments"`
	ComplianceState             *string                       `pulumi:"complianceState"`
	Evidence                    []AttestationEvidenceResponse `pulumi:"evidence"`
	ExpiresOn                   *string                       `pulumi:"expiresOn"`
	Id                          string                        `pulumi:"id"`
	LastComplianceStateChangeAt string                        `pulumi:"lastComplianceStateChangeAt"`
	Name                        string                        `pulumi:"name"`
	Owner                       *string                       `pulumi:"owner"`
	PolicyAssignmentId          string                        `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceId *string                       `pulumi:"policyDefinitionReferenceId"`
	ProvisioningState           string                        `pulumi:"provisioningState"`
	SystemData                  SystemDataResponse            `pulumi:"systemData"`
	Type                        string                        `pulumi:"type"`
}

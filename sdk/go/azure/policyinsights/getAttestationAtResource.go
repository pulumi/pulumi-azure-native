


package policyinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAttestationAtResource(ctx *pulumi.Context, args *LookupAttestationAtResourceArgs, opts ...pulumi.InvokeOption) (*LookupAttestationAtResourceResult, error) {
	var rv LookupAttestationAtResourceResult
	err := ctx.Invoke("azure-native:policyinsights:getAttestationAtResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAttestationAtResourceArgs struct {
	AttestationName string `pulumi:"attestationName"`
	ResourceId      string `pulumi:"resourceId"`
}


type LookupAttestationAtResourceResult struct {
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

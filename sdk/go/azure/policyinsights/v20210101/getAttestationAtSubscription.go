


package v20210101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAttestationAtSubscription(ctx *pulumi.Context, args *LookupAttestationAtSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupAttestationAtSubscriptionResult, error) {
	var rv LookupAttestationAtSubscriptionResult
	err := ctx.Invoke("azure-native:policyinsights/v20210101:getAttestationAtSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAttestationAtSubscriptionArgs struct {
	AttestationName string `pulumi:"attestationName"`
}


type LookupAttestationAtSubscriptionResult struct {
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

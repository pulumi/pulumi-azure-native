


package v20211001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRemediationAtSubscription(ctx *pulumi.Context, args *LookupRemediationAtSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupRemediationAtSubscriptionResult, error) {
	var rv LookupRemediationAtSubscriptionResult
	err := ctx.Invoke("azure-native:policyinsights/v20211001:getRemediationAtSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRemediationAtSubscriptionArgs struct {
	RemediationName string `pulumi:"remediationName"`
}


type LookupRemediationAtSubscriptionResult struct {
	CorrelationId               string                                         `pulumi:"correlationId"`
	CreatedOn                   string                                         `pulumi:"createdOn"`
	DeploymentStatus            RemediationDeploymentSummaryResponse           `pulumi:"deploymentStatus"`
	FailureThreshold            *RemediationPropertiesResponseFailureThreshold `pulumi:"failureThreshold"`
	Filters                     *RemediationFiltersResponse                    `pulumi:"filters"`
	Id                          string                                         `pulumi:"id"`
	LastUpdatedOn               string                                         `pulumi:"lastUpdatedOn"`
	Name                        string                                         `pulumi:"name"`
	ParallelDeployments         *int                                           `pulumi:"parallelDeployments"`
	PolicyAssignmentId          *string                                        `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceId *string                                        `pulumi:"policyDefinitionReferenceId"`
	ProvisioningState           string                                         `pulumi:"provisioningState"`
	ResourceCount               *int                                           `pulumi:"resourceCount"`
	ResourceDiscoveryMode       *string                                        `pulumi:"resourceDiscoveryMode"`
	StatusMessage               string                                         `pulumi:"statusMessage"`
	SystemData                  SystemDataResponse                             `pulumi:"systemData"`
	Type                        string                                         `pulumi:"type"`
}

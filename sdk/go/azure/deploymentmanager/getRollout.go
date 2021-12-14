


package deploymentmanager

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRollout(ctx *pulumi.Context, args *LookupRolloutArgs, opts ...pulumi.InvokeOption) (*LookupRolloutResult, error) {
	var rv LookupRolloutResult
	err := ctx.Invoke("azure-native:deploymentmanager:getRollout", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRolloutArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RetryAttempt      *int   `pulumi:"retryAttempt"`
	RolloutName       string `pulumi:"rolloutName"`
}


type LookupRolloutResult struct {
	ArtifactSourceId        *string                      `pulumi:"artifactSourceId"`
	BuildVersion            string                       `pulumi:"buildVersion"`
	Id                      string                       `pulumi:"id"`
	Identity                *IdentityResponse            `pulumi:"identity"`
	Location                string                       `pulumi:"location"`
	Name                    string                       `pulumi:"name"`
	OperationInfo           RolloutOperationInfoResponse `pulumi:"operationInfo"`
	Services                []ServiceResponse            `pulumi:"services"`
	Status                  string                       `pulumi:"status"`
	StepGroups              []StepGroupResponse          `pulumi:"stepGroups"`
	Tags                    map[string]string            `pulumi:"tags"`
	TargetServiceTopologyId string                       `pulumi:"targetServiceTopologyId"`
	TotalRetryAttempts      int                          `pulumi:"totalRetryAttempts"`
	Type                    string                       `pulumi:"type"`
}

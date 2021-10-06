


package v20191101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataCollectionRule(ctx *pulumi.Context, args *LookupDataCollectionRuleArgs, opts ...pulumi.InvokeOption) (*LookupDataCollectionRuleResult, error) {
	var rv LookupDataCollectionRuleResult
	err := ctx.Invoke("azure-native:insights/v20191101preview:getDataCollectionRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataCollectionRuleArgs struct {
	DataCollectionRuleName string `pulumi:"dataCollectionRuleName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupDataCollectionRuleResult struct {
	DataFlows         []DataFlowResponse                      `pulumi:"dataFlows"`
	DataSources       *DataCollectionRuleResponseDataSources  `pulumi:"dataSources"`
	Description       *string                                 `pulumi:"description"`
	Destinations      *DataCollectionRuleResponseDestinations `pulumi:"destinations"`
	Etag              string                                  `pulumi:"etag"`
	Id                string                                  `pulumi:"id"`
	ImmutableId       string                                  `pulumi:"immutableId"`
	Kind              *string                                 `pulumi:"kind"`
	Location          string                                  `pulumi:"location"`
	Name              string                                  `pulumi:"name"`
	ProvisioningState string                                  `pulumi:"provisioningState"`
	Tags              map[string]string                       `pulumi:"tags"`
	Type              string                                  `pulumi:"type"`
}

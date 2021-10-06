


package v20210401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupThreatIntelligenceIndicator(ctx *pulumi.Context, args *LookupThreatIntelligenceIndicatorArgs, opts ...pulumi.InvokeOption) (*LookupThreatIntelligenceIndicatorResult, error) {
	var rv LookupThreatIntelligenceIndicatorResult
	err := ctx.Invoke("azure-native:securityinsights/v20210401:getThreatIntelligenceIndicator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupThreatIntelligenceIndicatorArgs struct {
	Name                                string `pulumi:"name"`
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupThreatIntelligenceIndicatorResult struct {
	Etag       *string            `pulumi:"etag"`
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

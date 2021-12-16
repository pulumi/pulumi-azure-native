


package v20200202preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabaseAdvisor(ctx *pulumi.Context, args *LookupDatabaseAdvisorArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAdvisorResult, error) {
	var rv LookupDatabaseAdvisorResult
	err := ctx.Invoke("azure-native:sql/v20200202preview:getDatabaseAdvisor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseAdvisorArgs struct {
	AdvisorName       string `pulumi:"advisorName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupDatabaseAdvisorResult struct {
	AdvisorStatus                  string                      `pulumi:"advisorStatus"`
	AutoExecuteStatus              string                      `pulumi:"autoExecuteStatus"`
	AutoExecuteStatusInheritedFrom string                      `pulumi:"autoExecuteStatusInheritedFrom"`
	Id                             string                      `pulumi:"id"`
	Kind                           string                      `pulumi:"kind"`
	LastChecked                    string                      `pulumi:"lastChecked"`
	Location                       string                      `pulumi:"location"`
	Name                           string                      `pulumi:"name"`
	RecommendationsStatus          string                      `pulumi:"recommendationsStatus"`
	RecommendedActions             []RecommendedActionResponse `pulumi:"recommendedActions"`
	Type                           string                      `pulumi:"type"`
}

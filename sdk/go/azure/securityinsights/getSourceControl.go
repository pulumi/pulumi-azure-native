


package securityinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceControl(ctx *pulumi.Context, args *LookupSourceControlArgs, opts ...pulumi.InvokeOption) (*LookupSourceControlResult, error) {
	var rv LookupSourceControlResult
	err := ctx.Invoke("azure-native:securityinsights:getSourceControl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSourceControlArgs struct {
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	SourceControlId                     string `pulumi:"sourceControlId"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupSourceControlResult struct {
	ContentTypes    []string           `pulumi:"contentTypes"`
	Description     *string            `pulumi:"description"`
	DisplayName     string             `pulumi:"displayName"`
	Etag            *string            `pulumi:"etag"`
	Id              string             `pulumi:"id"`
	Name            string             `pulumi:"name"`
	RepoType        string             `pulumi:"repoType"`
	Repository      RepositoryResponse `pulumi:"repository"`
	SourceControlId *string            `pulumi:"sourceControlId"`
	SystemData      SystemDataResponse `pulumi:"systemData"`
	Type            string             `pulumi:"type"`
}

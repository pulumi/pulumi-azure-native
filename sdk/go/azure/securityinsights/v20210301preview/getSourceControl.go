


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceControl(ctx *pulumi.Context, args *LookupSourceControlArgs, opts ...pulumi.InvokeOption) (*LookupSourceControlResult, error) {
	var rv LookupSourceControlResult
	err := ctx.Invoke("azure-native:securityinsights/v20210301preview:getSourceControl", args, &rv, opts...)
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
	ContentTypes       []string           `pulumi:"contentTypes"`
	CreatedAt          *string            `pulumi:"createdAt"`
	CreatedBy          *string            `pulumi:"createdBy"`
	CreatedByType      *string            `pulumi:"createdByType"`
	Description        *string            `pulumi:"description"`
	DisplayName        string             `pulumi:"displayName"`
	Etag               *string            `pulumi:"etag"`
	Id                 string             `pulumi:"id"`
	LastModifiedAt     *string            `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string            `pulumi:"lastModifiedBy"`
	LastModifiedByType *string            `pulumi:"lastModifiedByType"`
	Name               string             `pulumi:"name"`
	RepoType           string             `pulumi:"repoType"`
	Repository         RepositoryResponse `pulumi:"repository"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	Type               string             `pulumi:"type"`
}

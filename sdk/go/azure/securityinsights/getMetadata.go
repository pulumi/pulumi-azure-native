


package securityinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMetadata(ctx *pulumi.Context, args *LookupMetadataArgs, opts ...pulumi.InvokeOption) (*LookupMetadataResult, error) {
	var rv LookupMetadataResult
	err := ctx.Invoke("azure-native:securityinsights:getMetadata", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMetadataArgs struct {
	MetadataName                        string `pulumi:"metadataName"`
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupMetadataResult struct {
	Author           *MetadataAuthorResponse       `pulumi:"author"`
	Categories       *MetadataCategoriesResponse   `pulumi:"categories"`
	ContentId        *string                       `pulumi:"contentId"`
	Dependencies     *MetadataDependenciesResponse `pulumi:"dependencies"`
	Etag             *string                       `pulumi:"etag"`
	FirstPublishDate *string                       `pulumi:"firstPublishDate"`
	Id               string                        `pulumi:"id"`
	Kind             string                        `pulumi:"kind"`
	LastPublishDate  *string                       `pulumi:"lastPublishDate"`
	Name             string                        `pulumi:"name"`
	ParentId         string                        `pulumi:"parentId"`
	Providers        []string                      `pulumi:"providers"`
	Source           *MetadataSourceResponse       `pulumi:"source"`
	Support          *MetadataSupportResponse      `pulumi:"support"`
	SystemData       SystemDataResponse            `pulumi:"systemData"`
	Type             string                        `pulumi:"type"`
	Version          *string                       `pulumi:"version"`
}

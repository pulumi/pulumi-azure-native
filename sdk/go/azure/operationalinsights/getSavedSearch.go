


package operationalinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSavedSearch(ctx *pulumi.Context, args *LookupSavedSearchArgs, opts ...pulumi.InvokeOption) (*LookupSavedSearchResult, error) {
	var rv LookupSavedSearchResult
	err := ctx.Invoke("azure-native:operationalinsights:getSavedSearch", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSavedSearchArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SavedSearchId     string `pulumi:"savedSearchId"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupSavedSearchResult struct {
	Category           string        `pulumi:"category"`
	DisplayName        string        `pulumi:"displayName"`
	Etag               *string       `pulumi:"etag"`
	FunctionAlias      *string       `pulumi:"functionAlias"`
	FunctionParameters *string       `pulumi:"functionParameters"`
	Id                 string        `pulumi:"id"`
	Name               string        `pulumi:"name"`
	Query              string        `pulumi:"query"`
	Tags               []TagResponse `pulumi:"tags"`
	Type               string        `pulumi:"type"`
	Version            *float64      `pulumi:"version"`
}




package v20190901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupQuery(ctx *pulumi.Context, args *LookupQueryArgs, opts ...pulumi.InvokeOption) (*LookupQueryResult, error) {
	var rv LookupQueryResult
	err := ctx.Invoke("azure-native:operationalinsights/v20190901preview:getQuery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueryArgs struct {
	Id                string `pulumi:"id"`
	QueryPackName     string `pulumi:"queryPackName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupQueryResult struct {
	Author       string                                               `pulumi:"author"`
	Body         string                                               `pulumi:"body"`
	Description  *string                                              `pulumi:"description"`
	DisplayName  string                                               `pulumi:"displayName"`
	Id           string                                               `pulumi:"id"`
	Name         string                                               `pulumi:"name"`
	Properties   interface{}                                          `pulumi:"properties"`
	Related      *LogAnalyticsQueryPackQueryPropertiesResponseRelated `pulumi:"related"`
	SystemData   SystemDataResponse                                   `pulumi:"systemData"`
	Tags         map[string][]string                                  `pulumi:"tags"`
	TimeCreated  string                                               `pulumi:"timeCreated"`
	TimeModified string                                               `pulumi:"timeModified"`
	Type         string                                               `pulumi:"type"`
}

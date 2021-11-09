


package v20200301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGremlinResourceGremlinGraph(ctx *pulumi.Context, args *LookupGremlinResourceGremlinGraphArgs, opts ...pulumi.InvokeOption) (*LookupGremlinResourceGremlinGraphResult, error) {
	var rv LookupGremlinResourceGremlinGraphResult
	err := ctx.Invoke("azure-native:documentdb/v20200301:getGremlinResourceGremlinGraph", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGremlinResourceGremlinGraphArgs struct {
	AccountName       string `pulumi:"accountName"`
	DatabaseName      string `pulumi:"databaseName"`
	GraphName         string `pulumi:"graphName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupGremlinResourceGremlinGraphResult struct {
	Id       string                                     `pulumi:"id"`
	Location *string                                    `pulumi:"location"`
	Name     string                                     `pulumi:"name"`
	Options  *GremlinGraphGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *GremlinGraphGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                          `pulumi:"tags"`
	Type     string                                     `pulumi:"type"`
}

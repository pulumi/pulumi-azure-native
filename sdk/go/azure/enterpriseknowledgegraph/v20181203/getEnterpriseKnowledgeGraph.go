


package v20181203

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnterpriseKnowledgeGraph(ctx *pulumi.Context, args *LookupEnterpriseKnowledgeGraphArgs, opts ...pulumi.InvokeOption) (*LookupEnterpriseKnowledgeGraphResult, error) {
	var rv LookupEnterpriseKnowledgeGraphResult
	err := ctx.Invoke("azure-native:enterpriseknowledgegraph/v20181203:getEnterpriseKnowledgeGraph", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnterpriseKnowledgeGraphArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupEnterpriseKnowledgeGraphResult struct {
	Id         string                                     `pulumi:"id"`
	Location   *string                                    `pulumi:"location"`
	Name       string                                     `pulumi:"name"`
	Properties EnterpriseKnowledgeGraphPropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse                               `pulumi:"sku"`
	Tags       map[string]string                          `pulumi:"tags"`
	Type       string                                     `pulumi:"type"`
}

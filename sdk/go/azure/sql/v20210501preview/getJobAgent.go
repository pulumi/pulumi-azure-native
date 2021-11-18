


package v20210501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJobAgent(ctx *pulumi.Context, args *LookupJobAgentArgs, opts ...pulumi.InvokeOption) (*LookupJobAgentResult, error) {
	var rv LookupJobAgentResult
	err := ctx.Invoke("azure-native:sql/v20210501preview:getJobAgent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobAgentArgs struct {
	JobAgentName      string `pulumi:"jobAgentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupJobAgentResult struct {
	DatabaseId string            `pulumi:"databaseId"`
	Id         string            `pulumi:"id"`
	Location   string            `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Sku        *SkuResponse      `pulumi:"sku"`
	State      string            `pulumi:"state"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}

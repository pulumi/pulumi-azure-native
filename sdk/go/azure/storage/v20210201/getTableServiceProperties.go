


package v20210201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTableServiceProperties(ctx *pulumi.Context, args *LookupTableServicePropertiesArgs, opts ...pulumi.InvokeOption) (*LookupTableServicePropertiesResult, error) {
	var rv LookupTableServicePropertiesResult
	err := ctx.Invoke("azure-native:storage/v20210201:getTableServiceProperties", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTableServicePropertiesArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TableServiceName  string `pulumi:"tableServiceName"`
}


type LookupTableServicePropertiesResult struct {
	Cors *CorsRulesResponse `pulumi:"cors"`
	Id   string             `pulumi:"id"`
	Name string             `pulumi:"name"`
	Type string             `pulumi:"type"`
}

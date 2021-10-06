


package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDataSetMapping(ctx *pulumi.Context, args *LookupDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupDataSetMappingResult, error) {
	var rv LookupDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupDataSetMappingResult struct {
	Id   string `pulumi:"id"`
	Kind string `pulumi:"kind"`
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}

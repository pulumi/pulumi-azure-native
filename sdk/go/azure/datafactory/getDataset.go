


package datafactory

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataset(ctx *pulumi.Context, args *LookupDatasetArgs, opts ...pulumi.InvokeOption) (*LookupDatasetResult, error) {
	var rv LookupDatasetResult
	err := ctx.Invoke("azure-native:datafactory:getDataset", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatasetArgs struct {
	DatasetName       string `pulumi:"datasetName"`
	FactoryName       string `pulumi:"factoryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatasetResult struct {
	Etag       string      `pulumi:"etag"`
	Id         string      `pulumi:"id"`
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}

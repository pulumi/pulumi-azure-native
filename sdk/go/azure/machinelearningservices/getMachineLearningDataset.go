


package machinelearningservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMachineLearningDataset(ctx *pulumi.Context, args *LookupMachineLearningDatasetArgs, opts ...pulumi.InvokeOption) (*LookupMachineLearningDatasetResult, error) {
	var rv LookupMachineLearningDatasetResult
	err := ctx.Invoke("azure-native:machinelearningservices:getMachineLearningDataset", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMachineLearningDatasetArgs struct {
	DatasetName       string `pulumi:"datasetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupMachineLearningDatasetResult struct {
	Id         string            `pulumi:"id"`
	Identity   *IdentityResponse `pulumi:"identity"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties DatasetResponse   `pulumi:"properties"`
	Sku        *SkuResponse      `pulumi:"sku"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}

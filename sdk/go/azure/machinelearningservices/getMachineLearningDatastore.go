


package machinelearningservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMachineLearningDatastore(ctx *pulumi.Context, args *LookupMachineLearningDatastoreArgs, opts ...pulumi.InvokeOption) (*LookupMachineLearningDatastoreResult, error) {
	var rv LookupMachineLearningDatastoreResult
	err := ctx.Invoke("azure-native:machinelearningservices:getMachineLearningDatastore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupMachineLearningDatastoreArgs struct {
	DatastoreName     string `pulumi:"datastoreName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupMachineLearningDatastoreResult struct {
	Id         string            `pulumi:"id"`
	Identity   *IdentityResponse `pulumi:"identity"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties DatastoreResponse `pulumi:"properties"`
	Sku        *SkuResponse      `pulumi:"sku"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}


func (val *LookupMachineLearningDatastoreResult) Defaults() *LookupMachineLearningDatastoreResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

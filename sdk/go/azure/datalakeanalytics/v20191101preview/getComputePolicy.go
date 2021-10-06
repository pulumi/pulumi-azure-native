


package v20191101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComputePolicy(ctx *pulumi.Context, args *LookupComputePolicyArgs, opts ...pulumi.InvokeOption) (*LookupComputePolicyResult, error) {
	var rv LookupComputePolicyResult
	err := ctx.Invoke("azure-native:datalakeanalytics/v20191101preview:getComputePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupComputePolicyArgs struct {
	AccountName       string `pulumi:"accountName"`
	ComputePolicyName string `pulumi:"computePolicyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupComputePolicyResult struct {
	Id                           string `pulumi:"id"`
	MaxDegreeOfParallelismPerJob int    `pulumi:"maxDegreeOfParallelismPerJob"`
	MinPriorityPerJob            int    `pulumi:"minPriorityPerJob"`
	Name                         string `pulumi:"name"`
	ObjectId                     string `pulumi:"objectId"`
	ObjectType                   string `pulumi:"objectType"`
	Type                         string `pulumi:"type"`
}




package v20200515preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupMachineLearningService(ctx *pulumi.Context, args *LookupMachineLearningServiceArgs, opts ...pulumi.InvokeOption) (*LookupMachineLearningServiceResult, error) {
	var rv LookupMachineLearningServiceResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200515preview:getMachineLearningService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMachineLearningServiceArgs struct {
	Expand            *bool  `pulumi:"expand"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupMachineLearningServiceResult struct {
	Id         string            `pulumi:"id"`
	Identity   *IdentityResponse `pulumi:"identity"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Sku        *SkuResponse      `pulumi:"sku"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}

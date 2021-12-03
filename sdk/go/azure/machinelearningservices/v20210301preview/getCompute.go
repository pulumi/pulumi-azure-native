


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCompute(ctx *pulumi.Context, args *LookupComputeArgs, opts ...pulumi.InvokeOption) (*LookupComputeResult, error) {
	var rv LookupComputeResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:getCompute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupComputeArgs struct {
	ComputeName       string `pulumi:"computeName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupComputeResult struct {
	Id         string             `pulumi:"id"`
	Identity   *IdentityResponse  `pulumi:"identity"`
	Location   *string            `pulumi:"location"`
	Name       string             `pulumi:"name"`
	Properties interface{}        `pulumi:"properties"`
	Sku        *SkuResponse       `pulumi:"sku"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Tags       map[string]string  `pulumi:"tags"`
	Type       string             `pulumi:"type"`
}

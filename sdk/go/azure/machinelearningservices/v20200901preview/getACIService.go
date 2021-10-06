


package v20200901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupACIService(ctx *pulumi.Context, args *LookupACIServiceArgs, opts ...pulumi.InvokeOption) (*LookupACIServiceResult, error) {
	var rv LookupACIServiceResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200901preview:getACIService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupACIServiceArgs struct {
	Expand            *bool  `pulumi:"expand"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupACIServiceResult struct {
	Id         string            `pulumi:"id"`
	Identity   *IdentityResponse `pulumi:"identity"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Sku        *SkuResponse      `pulumi:"sku"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}

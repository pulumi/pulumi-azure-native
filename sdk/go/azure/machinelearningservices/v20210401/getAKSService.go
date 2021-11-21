


package v20210401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAKSService(ctx *pulumi.Context, args *LookupAKSServiceArgs, opts ...pulumi.InvokeOption) (*LookupAKSServiceResult, error) {
	var rv LookupAKSServiceResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210401:getAKSService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAKSServiceArgs struct {
	Expand            *bool  `pulumi:"expand"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupAKSServiceResult struct {
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




package v20210701preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGraphResourceGraph(ctx *pulumi.Context, args *LookupGraphResourceGraphArgs, opts ...pulumi.InvokeOption) (*LookupGraphResourceGraphResult, error) {
	var rv LookupGraphResourceGraphResult
	err := ctx.Invoke("azure-native:documentdb/v20210701preview:getGraphResourceGraph", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGraphResourceGraphArgs struct {
	AccountName       string `pulumi:"accountName"`
	GraphName         string `pulumi:"graphName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupGraphResourceGraphResult struct {
	Id       string                                      `pulumi:"id"`
	Identity *ManagedServiceIdentityResponse             `pulumi:"identity"`
	Location *string                                     `pulumi:"location"`
	Name     string                                      `pulumi:"name"`
	Options  *GraphResourceGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *GraphResourceGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                           `pulumi:"tags"`
	Type     string                                      `pulumi:"type"`
}

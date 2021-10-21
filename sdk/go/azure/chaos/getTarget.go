


package chaos

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTarget(ctx *pulumi.Context, args *LookupTargetArgs, opts ...pulumi.InvokeOption) (*LookupTargetResult, error) {
	var rv LookupTargetResult
	err := ctx.Invoke("azure-native:chaos:getTarget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTargetArgs struct {
	ParentProviderNamespace string `pulumi:"parentProviderNamespace"`
	ParentResourceName      string `pulumi:"parentResourceName"`
	ParentResourceType      string `pulumi:"parentResourceType"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	TargetName              string `pulumi:"targetName"`
}


type LookupTargetResult struct {
	Id         string             `pulumi:"id"`
	Location   *string            `pulumi:"location"`
	Name       string             `pulumi:"name"`
	Properties interface{}        `pulumi:"properties"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

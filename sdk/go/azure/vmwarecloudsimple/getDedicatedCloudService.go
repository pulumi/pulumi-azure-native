


package vmwarecloudsimple

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDedicatedCloudService(ctx *pulumi.Context, args *LookupDedicatedCloudServiceArgs, opts ...pulumi.InvokeOption) (*LookupDedicatedCloudServiceResult, error) {
	var rv LookupDedicatedCloudServiceResult
	err := ctx.Invoke("azure-native:vmwarecloudsimple:getDedicatedCloudService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDedicatedCloudServiceArgs struct {
	DedicatedCloudServiceName string `pulumi:"dedicatedCloudServiceName"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
}


type LookupDedicatedCloudServiceResult struct {
	GatewaySubnet      string            `pulumi:"gatewaySubnet"`
	Id                 string            `pulumi:"id"`
	IsAccountOnboarded string            `pulumi:"isAccountOnboarded"`
	Location           string            `pulumi:"location"`
	Name               string            `pulumi:"name"`
	Nodes              int               `pulumi:"nodes"`
	ServiceURL         string            `pulumi:"serviceURL"`
	Tags               map[string]string `pulumi:"tags"`
	Type               string            `pulumi:"type"`
}

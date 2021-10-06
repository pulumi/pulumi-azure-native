


package v20160402

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOrigin(ctx *pulumi.Context, args *LookupOriginArgs, opts ...pulumi.InvokeOption) (*LookupOriginResult, error) {
	var rv LookupOriginResult
	err := ctx.Invoke("azure-native:cdn/v20160402:getOrigin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOriginArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	OriginName        string `pulumi:"originName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupOriginResult struct {
	HostName          string `pulumi:"hostName"`
	HttpPort          *int   `pulumi:"httpPort"`
	HttpsPort         *int   `pulumi:"httpsPort"`
	Id                string `pulumi:"id"`
	Name              string `pulumi:"name"`
	ProvisioningState string `pulumi:"provisioningState"`
	ResourceState     string `pulumi:"resourceState"`
	Type              string `pulumi:"type"`
}

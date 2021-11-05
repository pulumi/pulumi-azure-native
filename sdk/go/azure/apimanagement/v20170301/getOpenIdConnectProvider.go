


package v20170301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOpenIdConnectProvider(ctx *pulumi.Context, args *LookupOpenIdConnectProviderArgs, opts ...pulumi.InvokeOption) (*LookupOpenIdConnectProviderResult, error) {
	var rv LookupOpenIdConnectProviderResult
	err := ctx.Invoke("azure-native:apimanagement/v20170301:getOpenIdConnectProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOpenIdConnectProviderArgs struct {
	Opid              string `pulumi:"opid"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupOpenIdConnectProviderResult struct {
	ClientId         string  `pulumi:"clientId"`
	ClientSecret     *string `pulumi:"clientSecret"`
	Description      *string `pulumi:"description"`
	DisplayName      string  `pulumi:"displayName"`
	Id               string  `pulumi:"id"`
	MetadataEndpoint string  `pulumi:"metadataEndpoint"`
	Name             string  `pulumi:"name"`
	Type             string  `pulumi:"type"`
}

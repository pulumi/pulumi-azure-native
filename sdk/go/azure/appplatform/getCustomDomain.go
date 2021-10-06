


package appplatform

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomDomain(ctx *pulumi.Context, args *LookupCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupCustomDomainResult, error) {
	var rv LookupCustomDomainResult
	err := ctx.Invoke("azure-native:appplatform:getCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomDomainArgs struct {
	AppName           string `pulumi:"appName"`
	DomainName        string `pulumi:"domainName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupCustomDomainResult struct {
	Id         string                         `pulumi:"id"`
	Name       string                         `pulumi:"name"`
	Properties CustomDomainPropertiesResponse `pulumi:"properties"`
	Type       string                         `pulumi:"type"`
}

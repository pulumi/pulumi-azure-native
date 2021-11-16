


package apimanagement

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProperty(ctx *pulumi.Context, args *LookupPropertyArgs, opts ...pulumi.InvokeOption) (*LookupPropertyResult, error) {
	var rv LookupPropertyResult
	err := ctx.Invoke("azure-native:apimanagement:getProperty", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPropertyArgs struct {
	PropId            string `pulumi:"propId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupPropertyResult struct {
	DisplayName string   `pulumi:"displayName"`
	Id          string   `pulumi:"id"`
	Name        string   `pulumi:"name"`
	Secret      *bool    `pulumi:"secret"`
	Tags        []string `pulumi:"tags"`
	Type        string   `pulumi:"type"`
	Value       string   `pulumi:"value"`
}

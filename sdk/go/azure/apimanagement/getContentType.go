


package apimanagement

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContentType(ctx *pulumi.Context, args *LookupContentTypeArgs, opts ...pulumi.InvokeOption) (*LookupContentTypeResult, error) {
	var rv LookupContentTypeResult
	err := ctx.Invoke("azure-native:apimanagement:getContentType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContentTypeArgs struct {
	ContentTypeId     string `pulumi:"contentTypeId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupContentTypeResult struct {
	Description *string     `pulumi:"description"`
	Id          string      `pulumi:"id"`
	Name        string      `pulumi:"name"`
	Schema      interface{} `pulumi:"schema"`
	Type        string      `pulumi:"type"`
	Version     *string     `pulumi:"version"`
}

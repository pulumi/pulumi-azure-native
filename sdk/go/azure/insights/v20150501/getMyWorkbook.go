


package v20150501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMyWorkbook(ctx *pulumi.Context, args *LookupMyWorkbookArgs, opts ...pulumi.InvokeOption) (*LookupMyWorkbookResult, error) {
	var rv LookupMyWorkbookResult
	err := ctx.Invoke("azure-native:insights/v20150501:getMyWorkbook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMyWorkbookArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupMyWorkbookResult struct {
	Category       string            `pulumi:"category"`
	DisplayName    string            `pulumi:"displayName"`
	Id             *string           `pulumi:"id"`
	Kind           *string           `pulumi:"kind"`
	Location       *string           `pulumi:"location"`
	Name           *string           `pulumi:"name"`
	SerializedData string            `pulumi:"serializedData"`
	SourceId       *string           `pulumi:"sourceId"`
	Tags           map[string]string `pulumi:"tags"`
	TimeModified   string            `pulumi:"timeModified"`
	Type           *string           `pulumi:"type"`
	UserId         string            `pulumi:"userId"`
	Version        *string           `pulumi:"version"`
}

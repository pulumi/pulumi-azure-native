


package v20190901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMediaGraph(ctx *pulumi.Context, args *LookupMediaGraphArgs, opts ...pulumi.InvokeOption) (*LookupMediaGraphResult, error) {
	var rv LookupMediaGraphResult
	err := ctx.Invoke("azure-native:media/v20190901preview:getMediaGraph", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMediaGraphArgs struct {
	AccountName       string `pulumi:"accountName"`
	MediaGraphName    string `pulumi:"mediaGraphName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupMediaGraphResult struct {
	Created      string                         `pulumi:"created"`
	Description  *string                        `pulumi:"description"`
	Id           string                         `pulumi:"id"`
	LastModified string                         `pulumi:"lastModified"`
	Name         string                         `pulumi:"name"`
	Sinks        []MediaGraphAssetSinkResponse  `pulumi:"sinks"`
	Sources      []MediaGraphRtspSourceResponse `pulumi:"sources"`
	State        string                         `pulumi:"state"`
	Type         string                         `pulumi:"type"`
}

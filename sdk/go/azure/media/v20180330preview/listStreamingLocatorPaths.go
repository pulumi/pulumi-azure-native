


package v20180330preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStreamingLocatorPaths(ctx *pulumi.Context, args *ListStreamingLocatorPathsArgs, opts ...pulumi.InvokeOption) (*ListStreamingLocatorPathsResult, error) {
	var rv ListStreamingLocatorPathsResult
	err := ctx.Invoke("azure-native:media/v20180330preview:listStreamingLocatorPaths", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStreamingLocatorPathsArgs struct {
	AccountName          string `pulumi:"accountName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	StreamingLocatorName string `pulumi:"streamingLocatorName"`
}


type ListStreamingLocatorPathsResult struct {
	DownloadPaths  []string                `pulumi:"downloadPaths"`
	StreamingPaths []StreamingPathResponse `pulumi:"streamingPaths"`
}

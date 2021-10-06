


package v20180330preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStreamingLocatorContentKeys(ctx *pulumi.Context, args *ListStreamingLocatorContentKeysArgs, opts ...pulumi.InvokeOption) (*ListStreamingLocatorContentKeysResult, error) {
	var rv ListStreamingLocatorContentKeysResult
	err := ctx.Invoke("azure-native:media/v20180330preview:listStreamingLocatorContentKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStreamingLocatorContentKeysArgs struct {
	AccountName          string `pulumi:"accountName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	StreamingLocatorName string `pulumi:"streamingLocatorName"`
}


type ListStreamingLocatorContentKeysResult struct {
	ContentKeys []StreamingLocatorContentKeyResponse `pulumi:"contentKeys"`
}

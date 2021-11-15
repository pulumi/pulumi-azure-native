


package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStreamingLocator(ctx *pulumi.Context, args *LookupStreamingLocatorArgs, opts ...pulumi.InvokeOption) (*LookupStreamingLocatorResult, error) {
	var rv LookupStreamingLocatorResult
	err := ctx.Invoke("azure-native:media/v20180601preview:getStreamingLocator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStreamingLocatorArgs struct {
	AccountName          string `pulumi:"accountName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	StreamingLocatorName string `pulumi:"streamingLocatorName"`
}


type LookupStreamingLocatorResult struct {
	AlternativeMediaId          *string                              `pulumi:"alternativeMediaId"`
	AssetName                   string                               `pulumi:"assetName"`
	ContentKeys                 []StreamingLocatorContentKeyResponse `pulumi:"contentKeys"`
	Created                     string                               `pulumi:"created"`
	DefaultContentKeyPolicyName *string                              `pulumi:"defaultContentKeyPolicyName"`
	EndTime                     *string                              `pulumi:"endTime"`
	Id                          string                               `pulumi:"id"`
	Name                        string                               `pulumi:"name"`
	StartTime                   *string                              `pulumi:"startTime"`
	StreamingLocatorId          *string                              `pulumi:"streamingLocatorId"`
	StreamingPolicyName         string                               `pulumi:"streamingPolicyName"`
	Type                        string                               `pulumi:"type"`
}

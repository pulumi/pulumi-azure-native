


package videoanalyzer

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListVideoStreamingToken(ctx *pulumi.Context, args *ListVideoStreamingTokenArgs, opts ...pulumi.InvokeOption) (*ListVideoStreamingTokenResult, error) {
	var rv ListVideoStreamingTokenResult
	err := ctx.Invoke("azure-native:videoanalyzer:listVideoStreamingToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListVideoStreamingTokenArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VideoName         string `pulumi:"videoName"`
}


type ListVideoStreamingTokenResult struct {
	ExpirationDate string `pulumi:"expirationDate"`
	Token          string `pulumi:"token"`
}

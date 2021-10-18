


package v20211101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVideo(ctx *pulumi.Context, args *LookupVideoArgs, opts ...pulumi.InvokeOption) (*LookupVideoResult, error) {
	var rv LookupVideoResult
	err := ctx.Invoke("azure-native:videoanalyzer/v20211101preview:getVideo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVideoArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VideoName         string `pulumi:"videoName"`
}


type LookupVideoResult struct {
	Archival    *VideoArchivalResponse   `pulumi:"archival"`
	ContentUrls VideoContentUrlsResponse `pulumi:"contentUrls"`
	Description *string                  `pulumi:"description"`
	Flags       VideoFlagsResponse       `pulumi:"flags"`
	Id          string                   `pulumi:"id"`
	MediaInfo   VideoMediaInfoResponse   `pulumi:"mediaInfo"`
	Name        string                   `pulumi:"name"`
	SystemData  SystemDataResponse       `pulumi:"systemData"`
	Title       *string                  `pulumi:"title"`
	Type        string                   `pulumi:"type"`
}

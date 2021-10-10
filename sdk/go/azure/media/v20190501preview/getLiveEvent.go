


package v20190501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLiveEvent(ctx *pulumi.Context, args *LookupLiveEventArgs, opts ...pulumi.InvokeOption) (*LookupLiveEventResult, error) {
	var rv LookupLiveEventResult
	err := ctx.Invoke("azure-native:media/v20190501preview:getLiveEvent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLiveEventArgs struct {
	AccountName       string `pulumi:"accountName"`
	LiveEventName     string `pulumi:"liveEventName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLiveEventResult struct {
	Created                 string                           `pulumi:"created"`
	CrossSiteAccessPolicies *CrossSiteAccessPoliciesResponse `pulumi:"crossSiteAccessPolicies"`
	Description             *string                          `pulumi:"description"`
	Encoding                *LiveEventEncodingResponse       `pulumi:"encoding"`
	Id                      string                           `pulumi:"id"`
	Input                   LiveEventInputResponse           `pulumi:"input"`
	LastModified            string                           `pulumi:"lastModified"`
	Location                *string                          `pulumi:"location"`
	Name                    string                           `pulumi:"name"`
	Preview                 *LiveEventPreviewResponse        `pulumi:"preview"`
	ProvisioningState       string                           `pulumi:"provisioningState"`
	ResourceState           string                           `pulumi:"resourceState"`
	StreamOptions           []string                         `pulumi:"streamOptions"`
	Tags                    map[string]string                `pulumi:"tags"`
	Transcriptions          []LiveEventTranscriptionResponse `pulumi:"transcriptions"`
	Type                    string                           `pulumi:"type"`
	VanityUrl               *bool                            `pulumi:"vanityUrl"`
}

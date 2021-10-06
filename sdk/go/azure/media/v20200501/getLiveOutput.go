


package v20200501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLiveOutput(ctx *pulumi.Context, args *LookupLiveOutputArgs, opts ...pulumi.InvokeOption) (*LookupLiveOutputResult, error) {
	var rv LookupLiveOutputResult
	err := ctx.Invoke("azure-native:media/v20200501:getLiveOutput", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLiveOutputArgs struct {
	AccountName       string `pulumi:"accountName"`
	LiveEventName     string `pulumi:"liveEventName"`
	LiveOutputName    string `pulumi:"liveOutputName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLiveOutputResult struct {
	ArchiveWindowLength string       `pulumi:"archiveWindowLength"`
	AssetName           string       `pulumi:"assetName"`
	Created             string       `pulumi:"created"`
	Description         *string      `pulumi:"description"`
	Hls                 *HlsResponse `pulumi:"hls"`
	Id                  string       `pulumi:"id"`
	LastModified        string       `pulumi:"lastModified"`
	ManifestName        *string      `pulumi:"manifestName"`
	Name                string       `pulumi:"name"`
	OutputSnapTime      *float64     `pulumi:"outputSnapTime"`
	ProvisioningState   string       `pulumi:"provisioningState"`
	ResourceState       string       `pulumi:"resourceState"`
	Type                string       `pulumi:"type"`
}

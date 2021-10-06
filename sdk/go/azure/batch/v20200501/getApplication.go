


package v20200501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("azure-native:batch/v20200501:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	AccountName       string `pulumi:"accountName"`
	ApplicationName   string `pulumi:"applicationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupApplicationResult struct {
	AllowUpdates   *bool   `pulumi:"allowUpdates"`
	DefaultVersion *string `pulumi:"defaultVersion"`
	DisplayName    *string `pulumi:"displayName"`
	Etag           string  `pulumi:"etag"`
	Id             string  `pulumi:"id"`
	Name           string  `pulumi:"name"`
	Type           string  `pulumi:"type"`
}

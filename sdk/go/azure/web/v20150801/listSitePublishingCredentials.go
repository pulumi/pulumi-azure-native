


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSitePublishingCredentials(ctx *pulumi.Context, args *ListSitePublishingCredentialsArgs, opts ...pulumi.InvokeOption) (*ListSitePublishingCredentialsResult, error) {
	var rv ListSitePublishingCredentialsResult
	err := ctx.Invoke("azure-native:web/v20150801:listSitePublishingCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSitePublishingCredentialsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListSitePublishingCredentialsResult struct {
	Id                 *string           `pulumi:"id"`
	Kind               *string           `pulumi:"kind"`
	Location           string            `pulumi:"location"`
	Name               *string           `pulumi:"name"`
	PublishingPassword *string           `pulumi:"publishingPassword"`
	PublishingUserName *string           `pulumi:"publishingUserName"`
	ScmUri             *string           `pulumi:"scmUri"`
	Tags               map[string]string `pulumi:"tags"`
	Type               *string           `pulumi:"type"`
}

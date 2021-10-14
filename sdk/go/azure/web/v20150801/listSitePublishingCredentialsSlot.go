


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSitePublishingCredentialsSlot(ctx *pulumi.Context, args *ListSitePublishingCredentialsSlotArgs, opts ...pulumi.InvokeOption) (*ListSitePublishingCredentialsSlotResult, error) {
	var rv ListSitePublishingCredentialsSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:listSitePublishingCredentialsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSitePublishingCredentialsSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListSitePublishingCredentialsSlotResult struct {
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

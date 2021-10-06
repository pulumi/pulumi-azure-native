


package v20210101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppPublishingCredentials(ctx *pulumi.Context, args *ListWebAppPublishingCredentialsArgs, opts ...pulumi.InvokeOption) (*ListWebAppPublishingCredentialsResult, error) {
	var rv ListWebAppPublishingCredentialsResult
	err := ctx.Invoke("azure-native:web/v20210101:listWebAppPublishingCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppPublishingCredentialsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppPublishingCredentialsResult struct {
	Id                         string  `pulumi:"id"`
	Kind                       *string `pulumi:"kind"`
	Name                       string  `pulumi:"name"`
	PublishingPassword         *string `pulumi:"publishingPassword"`
	PublishingPasswordHash     *string `pulumi:"publishingPasswordHash"`
	PublishingPasswordHashSalt *string `pulumi:"publishingPasswordHashSalt"`
	PublishingUserName         string  `pulumi:"publishingUserName"`
	ScmUri                     *string `pulumi:"scmUri"`
	Type                       string  `pulumi:"type"`
}

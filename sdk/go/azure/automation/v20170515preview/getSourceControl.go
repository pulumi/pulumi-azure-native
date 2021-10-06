


package v20170515preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceControl(ctx *pulumi.Context, args *LookupSourceControlArgs, opts ...pulumi.InvokeOption) (*LookupSourceControlResult, error) {
	var rv LookupSourceControlResult
	err := ctx.Invoke("azure-native:automation/v20170515preview:getSourceControl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSourceControlArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	SourceControlName     string `pulumi:"sourceControlName"`
}


type LookupSourceControlResult struct {
	AutoSync         *bool   `pulumi:"autoSync"`
	Branch           *string `pulumi:"branch"`
	CreationTime     *string `pulumi:"creationTime"`
	Description      *string `pulumi:"description"`
	FolderPath       *string `pulumi:"folderPath"`
	Id               string  `pulumi:"id"`
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	Name             string  `pulumi:"name"`
	PublishRunbook   *bool   `pulumi:"publishRunbook"`
	RepoUrl          *string `pulumi:"repoUrl"`
	SourceType       *string `pulumi:"sourceType"`
	Type             string  `pulumi:"type"`
}

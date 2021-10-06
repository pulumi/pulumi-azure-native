


package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppSourceControl(ctx *pulumi.Context, args *LookupWebAppSourceControlArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSourceControlResult, error) {
	var rv LookupWebAppSourceControlResult
	err := ctx.Invoke("azure-native:web/v20201201:getWebAppSourceControl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppSourceControlArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWebAppSourceControlResult struct {
	Branch                    *string                            `pulumi:"branch"`
	DeploymentRollbackEnabled *bool                              `pulumi:"deploymentRollbackEnabled"`
	GitHubActionConfiguration *GitHubActionConfigurationResponse `pulumi:"gitHubActionConfiguration"`
	Id                        string                             `pulumi:"id"`
	IsGitHubAction            *bool                              `pulumi:"isGitHubAction"`
	IsManualIntegration       *bool                              `pulumi:"isManualIntegration"`
	IsMercurial               *bool                              `pulumi:"isMercurial"`
	Kind                      *string                            `pulumi:"kind"`
	Name                      string                             `pulumi:"name"`
	RepoUrl                   *string                            `pulumi:"repoUrl"`
	Type                      string                             `pulumi:"type"`
}

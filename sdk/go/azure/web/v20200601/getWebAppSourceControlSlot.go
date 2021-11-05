


package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppSourceControlSlot(ctx *pulumi.Context, args *LookupWebAppSourceControlSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSourceControlSlotResult, error) {
	var rv LookupWebAppSourceControlSlotResult
	err := ctx.Invoke("azure-native:web/v20200601:getWebAppSourceControlSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppSourceControlSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupWebAppSourceControlSlotResult struct {
	Branch                    *string `pulumi:"branch"`
	DeploymentRollbackEnabled *bool   `pulumi:"deploymentRollbackEnabled"`
	Id                        string  `pulumi:"id"`
	IsGitHubAction            *bool   `pulumi:"isGitHubAction"`
	IsManualIntegration       *bool   `pulumi:"isManualIntegration"`
	IsMercurial               *bool   `pulumi:"isMercurial"`
	Kind                      *string `pulumi:"kind"`
	Name                      string  `pulumi:"name"`
	RepoUrl                   *string `pulumi:"repoUrl"`
	Type                      string  `pulumi:"type"`
}

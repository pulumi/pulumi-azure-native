


package v20190801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppSourceControl(ctx *pulumi.Context, args *LookupWebAppSourceControlArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSourceControlResult, error) {
	var rv LookupWebAppSourceControlResult
	err := ctx.Invoke("azure-native:web/v20190801:getWebAppSourceControl", args, &rv, opts...)
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
	Branch                    *string `pulumi:"branch"`
	DeploymentRollbackEnabled *bool   `pulumi:"deploymentRollbackEnabled"`
	Id                        string  `pulumi:"id"`
	IsManualIntegration       *bool   `pulumi:"isManualIntegration"`
	IsMercurial               *bool   `pulumi:"isMercurial"`
	Kind                      *string `pulumi:"kind"`
	Name                      string  `pulumi:"name"`
	RepoUrl                   *string `pulumi:"repoUrl"`
	Type                      string  `pulumi:"type"`
}

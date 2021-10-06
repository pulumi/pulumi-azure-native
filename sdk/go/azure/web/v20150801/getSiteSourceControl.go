


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSiteSourceControl(ctx *pulumi.Context, args *LookupSiteSourceControlArgs, opts ...pulumi.InvokeOption) (*LookupSiteSourceControlResult, error) {
	var rv LookupSiteSourceControlResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteSourceControl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteSourceControlArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupSiteSourceControlResult struct {
	Branch                    *string           `pulumi:"branch"`
	DeploymentRollbackEnabled *bool             `pulumi:"deploymentRollbackEnabled"`
	Id                        *string           `pulumi:"id"`
	IsManualIntegration       *bool             `pulumi:"isManualIntegration"`
	IsMercurial               *bool             `pulumi:"isMercurial"`
	Kind                      *string           `pulumi:"kind"`
	Location                  string            `pulumi:"location"`
	Name                      *string           `pulumi:"name"`
	RepoUrl                   *string           `pulumi:"repoUrl"`
	Tags                      map[string]string `pulumi:"tags"`
	Type                      *string           `pulumi:"type"`
}

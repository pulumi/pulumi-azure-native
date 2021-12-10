


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSiteSourceControlSlot(ctx *pulumi.Context, args *LookupSiteSourceControlSlotArgs, opts ...pulumi.InvokeOption) (*LookupSiteSourceControlSlotResult, error) {
	var rv LookupSiteSourceControlSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteSourceControlSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteSourceControlSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupSiteSourceControlSlotResult struct {
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

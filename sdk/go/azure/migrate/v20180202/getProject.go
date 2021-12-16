


package v20180202

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	var rv LookupProjectResult
	err := ctx.Invoke("azure-native:migrate/v20180202:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectArgs struct {
	ProjectName       string `pulumi:"projectName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupProjectResult struct {
	CreatedTimestamp          string      `pulumi:"createdTimestamp"`
	CustomerWorkspaceId       *string     `pulumi:"customerWorkspaceId"`
	CustomerWorkspaceLocation *string     `pulumi:"customerWorkspaceLocation"`
	DiscoveryStatus           string      `pulumi:"discoveryStatus"`
	ETag                      *string     `pulumi:"eTag"`
	Id                        string      `pulumi:"id"`
	LastAssessmentTimestamp   string      `pulumi:"lastAssessmentTimestamp"`
	LastDiscoverySessionId    string      `pulumi:"lastDiscoverySessionId"`
	LastDiscoveryTimestamp    string      `pulumi:"lastDiscoveryTimestamp"`
	Location                  *string     `pulumi:"location"`
	Name                      string      `pulumi:"name"`
	NumberOfAssessments       int         `pulumi:"numberOfAssessments"`
	NumberOfGroups            int         `pulumi:"numberOfGroups"`
	NumberOfMachines          int         `pulumi:"numberOfMachines"`
	ProvisioningState         *string     `pulumi:"provisioningState"`
	Tags                      interface{} `pulumi:"tags"`
	Type                      string      `pulumi:"type"`
	UpdatedTimestamp          string      `pulumi:"updatedTimestamp"`
}

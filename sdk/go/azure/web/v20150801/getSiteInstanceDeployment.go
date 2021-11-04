


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSiteInstanceDeployment(ctx *pulumi.Context, args *LookupSiteInstanceDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupSiteInstanceDeploymentResult, error) {
	var rv LookupSiteInstanceDeploymentResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteInstanceDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteInstanceDeploymentArgs struct {
	Id                string `pulumi:"id"`
	InstanceId        string `pulumi:"instanceId"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupSiteInstanceDeploymentResult struct {
	Active      *bool             `pulumi:"active"`
	Author      *string           `pulumi:"author"`
	AuthorEmail *string           `pulumi:"authorEmail"`
	Deployer    *string           `pulumi:"deployer"`
	Details     *string           `pulumi:"details"`
	EndTime     *string           `pulumi:"endTime"`
	Id          *string           `pulumi:"id"`
	Kind        *string           `pulumi:"kind"`
	Location    string            `pulumi:"location"`
	Message     *string           `pulumi:"message"`
	Name        *string           `pulumi:"name"`
	StartTime   *string           `pulumi:"startTime"`
	Status      *int              `pulumi:"status"`
	Tags        map[string]string `pulumi:"tags"`
	Type        *string           `pulumi:"type"`
}

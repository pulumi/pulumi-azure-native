


package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppDeployment(ctx *pulumi.Context, args *LookupWebAppDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupWebAppDeploymentResult, error) {
	var rv LookupWebAppDeploymentResult
	err := ctx.Invoke("azure-native:web/v20201201:getWebAppDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppDeploymentArgs struct {
	Id                string `pulumi:"id"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWebAppDeploymentResult struct {
	Active      *bool   `pulumi:"active"`
	Author      *string `pulumi:"author"`
	AuthorEmail *string `pulumi:"authorEmail"`
	Deployer    *string `pulumi:"deployer"`
	Details     *string `pulumi:"details"`
	EndTime     *string `pulumi:"endTime"`
	Id          string  `pulumi:"id"`
	Kind        *string `pulumi:"kind"`
	Message     *string `pulumi:"message"`
	Name        string  `pulumi:"name"`
	StartTime   *string `pulumi:"startTime"`
	Status      *int    `pulumi:"status"`
	Type        string  `pulumi:"type"`
}

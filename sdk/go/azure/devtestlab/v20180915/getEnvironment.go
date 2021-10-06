


package v20180915

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnvironment(ctx *pulumi.Context, args *LookupEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentResult, error) {
	var rv LookupEnvironmentResult
	err := ctx.Invoke("azure-native:devtestlab/v20180915:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	UserName          string  `pulumi:"userName"`
}


type LookupEnvironmentResult struct {
	ArmTemplateDisplayName *string                                  `pulumi:"armTemplateDisplayName"`
	CreatedByUser          string                                   `pulumi:"createdByUser"`
	DeploymentProperties   *EnvironmentDeploymentPropertiesResponse `pulumi:"deploymentProperties"`
	Id                     string                                   `pulumi:"id"`
	Location               *string                                  `pulumi:"location"`
	Name                   string                                   `pulumi:"name"`
	ProvisioningState      string                                   `pulumi:"provisioningState"`
	ResourceGroupId        string                                   `pulumi:"resourceGroupId"`
	Tags                   map[string]string                        `pulumi:"tags"`
	Type                   string                                   `pulumi:"type"`
	UniqueIdentifier       string                                   `pulumi:"uniqueIdentifier"`
}

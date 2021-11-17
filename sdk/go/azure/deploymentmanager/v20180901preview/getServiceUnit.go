


package v20180901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServiceUnit(ctx *pulumi.Context, args *LookupServiceUnitArgs, opts ...pulumi.InvokeOption) (*LookupServiceUnitResult, error) {
	var rv LookupServiceUnitResult
	err := ctx.Invoke("azure-native:deploymentmanager/v20180901preview:getServiceUnit", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceUnitArgs struct {
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	ServiceName         string `pulumi:"serviceName"`
	ServiceTopologyName string `pulumi:"serviceTopologyName"`
	ServiceUnitName     string `pulumi:"serviceUnitName"`
}


type LookupServiceUnitResult struct {
	Artifacts           *ServiceUnitArtifactsResponse `pulumi:"artifacts"`
	DeploymentMode      string                        `pulumi:"deploymentMode"`
	Id                  string                        `pulumi:"id"`
	Location            string                        `pulumi:"location"`
	Name                string                        `pulumi:"name"`
	Tags                map[string]string             `pulumi:"tags"`
	TargetResourceGroup string                        `pulumi:"targetResourceGroup"`
	Type                string                        `pulumi:"type"`
}

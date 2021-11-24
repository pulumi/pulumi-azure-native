


package v20210901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDeploymentLogFileUrl(ctx *pulumi.Context, args *GetDeploymentLogFileUrlArgs, opts ...pulumi.InvokeOption) (*GetDeploymentLogFileUrlResult, error) {
	var rv GetDeploymentLogFileUrlResult
	err := ctx.Invoke("azure-native:appplatform/v20210901preview:getDeploymentLogFileUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDeploymentLogFileUrlArgs struct {
	AppName           string `pulumi:"appName"`
	DeploymentName    string `pulumi:"deploymentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type GetDeploymentLogFileUrlResult struct {
	Url string `pulumi:"url"`
}




package v20201101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDeploymentLogFileUrl(ctx *pulumi.Context, args *GetDeploymentLogFileUrlArgs, opts ...pulumi.InvokeOption) (*GetDeploymentLogFileUrlResult, error) {
	var rv GetDeploymentLogFileUrlResult
	err := ctx.Invoke("azure-native:appplatform/v20201101preview:getDeploymentLogFileUrl", args, &rv, opts...)
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

func GetDeploymentLogFileUrlOutput(ctx *pulumi.Context, args GetDeploymentLogFileUrlOutputArgs, opts ...pulumi.InvokeOption) GetDeploymentLogFileUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDeploymentLogFileUrlResult, error) {
			args := v.(GetDeploymentLogFileUrlArgs)
			r, err := GetDeploymentLogFileUrl(ctx, &args, opts...)
			var s GetDeploymentLogFileUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDeploymentLogFileUrlResultOutput)
}

type GetDeploymentLogFileUrlOutputArgs struct {
	AppName           pulumi.StringInput `pulumi:"appName"`
	DeploymentName    pulumi.StringInput `pulumi:"deploymentName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (GetDeploymentLogFileUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeploymentLogFileUrlArgs)(nil)).Elem()
}


type GetDeploymentLogFileUrlResultOutput struct{ *pulumi.OutputState }

func (GetDeploymentLogFileUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeploymentLogFileUrlResult)(nil)).Elem()
}

func (o GetDeploymentLogFileUrlResultOutput) ToGetDeploymentLogFileUrlResultOutput() GetDeploymentLogFileUrlResultOutput {
	return o
}

func (o GetDeploymentLogFileUrlResultOutput) ToGetDeploymentLogFileUrlResultOutputWithContext(ctx context.Context) GetDeploymentLogFileUrlResultOutput {
	return o
}

func (o GetDeploymentLogFileUrlResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeploymentLogFileUrlResult) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDeploymentLogFileUrlResultOutput{})
}

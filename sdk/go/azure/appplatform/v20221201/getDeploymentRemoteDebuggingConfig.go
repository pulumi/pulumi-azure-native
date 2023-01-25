


package v20221201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDeploymentRemoteDebuggingConfig(ctx *pulumi.Context, args *GetDeploymentRemoteDebuggingConfigArgs, opts ...pulumi.InvokeOption) (*GetDeploymentRemoteDebuggingConfigResult, error) {
	var rv GetDeploymentRemoteDebuggingConfigResult
	err := ctx.Invoke("azure-native:appplatform/v20221201:getDeploymentRemoteDebuggingConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDeploymentRemoteDebuggingConfigArgs struct {
	AppName           string `pulumi:"appName"`
	DeploymentName    string `pulumi:"deploymentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type GetDeploymentRemoteDebuggingConfigResult struct {
	Enabled *bool `pulumi:"enabled"`
	Port    *int  `pulumi:"port"`
}

func GetDeploymentRemoteDebuggingConfigOutput(ctx *pulumi.Context, args GetDeploymentRemoteDebuggingConfigOutputArgs, opts ...pulumi.InvokeOption) GetDeploymentRemoteDebuggingConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDeploymentRemoteDebuggingConfigResult, error) {
			args := v.(GetDeploymentRemoteDebuggingConfigArgs)
			r, err := GetDeploymentRemoteDebuggingConfig(ctx, &args, opts...)
			var s GetDeploymentRemoteDebuggingConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDeploymentRemoteDebuggingConfigResultOutput)
}

type GetDeploymentRemoteDebuggingConfigOutputArgs struct {
	AppName           pulumi.StringInput `pulumi:"appName"`
	DeploymentName    pulumi.StringInput `pulumi:"deploymentName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (GetDeploymentRemoteDebuggingConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeploymentRemoteDebuggingConfigArgs)(nil)).Elem()
}


type GetDeploymentRemoteDebuggingConfigResultOutput struct{ *pulumi.OutputState }

func (GetDeploymentRemoteDebuggingConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeploymentRemoteDebuggingConfigResult)(nil)).Elem()
}

func (o GetDeploymentRemoteDebuggingConfigResultOutput) ToGetDeploymentRemoteDebuggingConfigResultOutput() GetDeploymentRemoteDebuggingConfigResultOutput {
	return o
}

func (o GetDeploymentRemoteDebuggingConfigResultOutput) ToGetDeploymentRemoteDebuggingConfigResultOutputWithContext(ctx context.Context) GetDeploymentRemoteDebuggingConfigResultOutput {
	return o
}

func (o GetDeploymentRemoteDebuggingConfigResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetDeploymentRemoteDebuggingConfigResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o GetDeploymentRemoteDebuggingConfigResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDeploymentRemoteDebuggingConfigResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDeploymentRemoteDebuggingConfigResultOutput{})
}

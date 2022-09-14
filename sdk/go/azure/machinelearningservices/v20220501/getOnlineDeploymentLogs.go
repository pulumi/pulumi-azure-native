


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOnlineDeploymentLogs(ctx *pulumi.Context, args *GetOnlineDeploymentLogsArgs, opts ...pulumi.InvokeOption) (*GetOnlineDeploymentLogsResult, error) {
	var rv GetOnlineDeploymentLogsResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20220501:getOnlineDeploymentLogs", args.Defaults(), &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetOnlineDeploymentLogsArgs struct {
	ContainerType     *string `pulumi:"containerType"`
	DeploymentName    string  `pulumi:"deploymentName"`
	EndpointName      string  `pulumi:"endpointName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	Tail              *int    `pulumi:"tail"`
	WorkspaceName     string  `pulumi:"workspaceName"`
}


func (val *GetOnlineDeploymentLogsArgs) Defaults() *GetOnlineDeploymentLogsArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ContainerType) {
		containerType_ := "InferenceServer"
		tmp.ContainerType = &containerType_
	}
	return &tmp
}

type GetOnlineDeploymentLogsResult struct {
	Content *string `pulumi:"content"`
}

func GetOnlineDeploymentLogsOutput(ctx *pulumi.Context, args GetOnlineDeploymentLogsOutputArgs, opts ...pulumi.InvokeOption) GetOnlineDeploymentLogsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOnlineDeploymentLogsResult, error) {
			args := v.(GetOnlineDeploymentLogsArgs)
			r, err := GetOnlineDeploymentLogs(ctx, &args, opts...)
			var s GetOnlineDeploymentLogsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOnlineDeploymentLogsResultOutput)
}

type GetOnlineDeploymentLogsOutputArgs struct {
	ContainerType     pulumi.StringPtrInput `pulumi:"containerType"`
	DeploymentName    pulumi.StringInput    `pulumi:"deploymentName"`
	EndpointName      pulumi.StringInput    `pulumi:"endpointName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	Tail              pulumi.IntPtrInput    `pulumi:"tail"`
	WorkspaceName     pulumi.StringInput    `pulumi:"workspaceName"`
}

func (GetOnlineDeploymentLogsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOnlineDeploymentLogsArgs)(nil)).Elem()
}

type GetOnlineDeploymentLogsResultOutput struct{ *pulumi.OutputState }

func (GetOnlineDeploymentLogsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOnlineDeploymentLogsResult)(nil)).Elem()
}

func (o GetOnlineDeploymentLogsResultOutput) ToGetOnlineDeploymentLogsResultOutput() GetOnlineDeploymentLogsResultOutput {
	return o
}

func (o GetOnlineDeploymentLogsResultOutput) ToGetOnlineDeploymentLogsResultOutputWithContext(ctx context.Context) GetOnlineDeploymentLogsResultOutput {
	return o
}

func (o GetOnlineDeploymentLogsResultOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOnlineDeploymentLogsResult) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOnlineDeploymentLogsResultOutput{})
}

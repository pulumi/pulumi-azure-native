


package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOnlineDeployment(ctx *pulumi.Context, args *LookupOnlineDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupOnlineDeploymentResult, error) {
	var rv LookupOnlineDeploymentResult
	err := ctx.Invoke("azure-native:machinelearningservices:getOnlineDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOnlineDeploymentArgs struct {
	DeploymentName    string `pulumi:"deploymentName"`
	EndpointName      string `pulumi:"endpointName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

type LookupOnlineDeploymentResult struct {
	Id         string                    `pulumi:"id"`
	Identity   *ResourceIdentityResponse `pulumi:"identity"`
	Kind       *string                   `pulumi:"kind"`
	Location   string                    `pulumi:"location"`
	Name       string                    `pulumi:"name"`
	Properties interface{}               `pulumi:"properties"`
	SystemData SystemDataResponse        `pulumi:"systemData"`
	Tags       map[string]string         `pulumi:"tags"`
	Type       string                    `pulumi:"type"`
}

func LookupOnlineDeploymentOutput(ctx *pulumi.Context, args LookupOnlineDeploymentOutputArgs, opts ...pulumi.InvokeOption) LookupOnlineDeploymentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOnlineDeploymentResult, error) {
			args := v.(LookupOnlineDeploymentArgs)
			r, err := LookupOnlineDeployment(ctx, &args, opts...)
			var s LookupOnlineDeploymentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOnlineDeploymentResultOutput)
}

type LookupOnlineDeploymentOutputArgs struct {
	DeploymentName    pulumi.StringInput `pulumi:"deploymentName"`
	EndpointName      pulumi.StringInput `pulumi:"endpointName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupOnlineDeploymentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOnlineDeploymentArgs)(nil)).Elem()
}

type LookupOnlineDeploymentResultOutput struct{ *pulumi.OutputState }

func (LookupOnlineDeploymentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOnlineDeploymentResult)(nil)).Elem()
}

func (o LookupOnlineDeploymentResultOutput) ToLookupOnlineDeploymentResultOutput() LookupOnlineDeploymentResultOutput {
	return o
}

func (o LookupOnlineDeploymentResultOutput) ToLookupOnlineDeploymentResultOutputWithContext(ctx context.Context) LookupOnlineDeploymentResultOutput {
	return o
}

func (o LookupOnlineDeploymentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineDeploymentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOnlineDeploymentResultOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupOnlineDeploymentResult) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o LookupOnlineDeploymentResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOnlineDeploymentResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupOnlineDeploymentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineDeploymentResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupOnlineDeploymentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineDeploymentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOnlineDeploymentResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupOnlineDeploymentResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupOnlineDeploymentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOnlineDeploymentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupOnlineDeploymentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOnlineDeploymentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupOnlineDeploymentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineDeploymentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOnlineDeploymentResultOutput{})
}

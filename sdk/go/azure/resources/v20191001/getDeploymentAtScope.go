


package v20191001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDeploymentAtScope(ctx *pulumi.Context, args *LookupDeploymentAtScopeArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentAtScopeResult, error) {
	var rv LookupDeploymentAtScopeResult
	err := ctx.Invoke("azure-native:resources/v20191001:getDeploymentAtScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentAtScopeArgs struct {
	DeploymentName string `pulumi:"deploymentName"`
	Scope          string `pulumi:"scope"`
}


type LookupDeploymentAtScopeResult struct {
	Id         string                               `pulumi:"id"`
	Location   *string                              `pulumi:"location"`
	Name       string                               `pulumi:"name"`
	Properties DeploymentPropertiesExtendedResponse `pulumi:"properties"`
	Tags       map[string]string                    `pulumi:"tags"`
	Type       string                               `pulumi:"type"`
}

func LookupDeploymentAtScopeOutput(ctx *pulumi.Context, args LookupDeploymentAtScopeOutputArgs, opts ...pulumi.InvokeOption) LookupDeploymentAtScopeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeploymentAtScopeResult, error) {
			args := v.(LookupDeploymentAtScopeArgs)
			r, err := LookupDeploymentAtScope(ctx, &args, opts...)
			var s LookupDeploymentAtScopeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeploymentAtScopeResultOutput)
}

type LookupDeploymentAtScopeOutputArgs struct {
	DeploymentName pulumi.StringInput `pulumi:"deploymentName"`
	Scope          pulumi.StringInput `pulumi:"scope"`
}

func (LookupDeploymentAtScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentAtScopeArgs)(nil)).Elem()
}


type LookupDeploymentAtScopeResultOutput struct{ *pulumi.OutputState }

func (LookupDeploymentAtScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentAtScopeResult)(nil)).Elem()
}

func (o LookupDeploymentAtScopeResultOutput) ToLookupDeploymentAtScopeResultOutput() LookupDeploymentAtScopeResultOutput {
	return o
}

func (o LookupDeploymentAtScopeResultOutput) ToLookupDeploymentAtScopeResultOutputWithContext(ctx context.Context) LookupDeploymentAtScopeResultOutput {
	return o
}

func (o LookupDeploymentAtScopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentAtScopeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDeploymentAtScopeResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeploymentAtScopeResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDeploymentAtScopeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentAtScopeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDeploymentAtScopeResultOutput) Properties() DeploymentPropertiesExtendedResponseOutput {
	return o.ApplyT(func(v LookupDeploymentAtScopeResult) DeploymentPropertiesExtendedResponse { return v.Properties }).(DeploymentPropertiesExtendedResponseOutput)
}

func (o LookupDeploymentAtScopeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDeploymentAtScopeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDeploymentAtScopeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentAtScopeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeploymentAtScopeResultOutput{})
}

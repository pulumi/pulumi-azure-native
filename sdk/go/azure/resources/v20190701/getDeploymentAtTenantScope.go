


package v20190701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDeploymentAtTenantScope(ctx *pulumi.Context, args *LookupDeploymentAtTenantScopeArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentAtTenantScopeResult, error) {
	var rv LookupDeploymentAtTenantScopeResult
	err := ctx.Invoke("azure-native:resources/v20190701:getDeploymentAtTenantScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentAtTenantScopeArgs struct {
	DeploymentName string `pulumi:"deploymentName"`
}


type LookupDeploymentAtTenantScopeResult struct {
	Id         string                               `pulumi:"id"`
	Location   *string                              `pulumi:"location"`
	Name       string                               `pulumi:"name"`
	Properties DeploymentPropertiesExtendedResponse `pulumi:"properties"`
	Type       string                               `pulumi:"type"`
}

func LookupDeploymentAtTenantScopeOutput(ctx *pulumi.Context, args LookupDeploymentAtTenantScopeOutputArgs, opts ...pulumi.InvokeOption) LookupDeploymentAtTenantScopeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeploymentAtTenantScopeResult, error) {
			args := v.(LookupDeploymentAtTenantScopeArgs)
			r, err := LookupDeploymentAtTenantScope(ctx, &args, opts...)
			var s LookupDeploymentAtTenantScopeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeploymentAtTenantScopeResultOutput)
}

type LookupDeploymentAtTenantScopeOutputArgs struct {
	DeploymentName pulumi.StringInput `pulumi:"deploymentName"`
}

func (LookupDeploymentAtTenantScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentAtTenantScopeArgs)(nil)).Elem()
}


type LookupDeploymentAtTenantScopeResultOutput struct{ *pulumi.OutputState }

func (LookupDeploymentAtTenantScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentAtTenantScopeResult)(nil)).Elem()
}

func (o LookupDeploymentAtTenantScopeResultOutput) ToLookupDeploymentAtTenantScopeResultOutput() LookupDeploymentAtTenantScopeResultOutput {
	return o
}

func (o LookupDeploymentAtTenantScopeResultOutput) ToLookupDeploymentAtTenantScopeResultOutputWithContext(ctx context.Context) LookupDeploymentAtTenantScopeResultOutput {
	return o
}

func (o LookupDeploymentAtTenantScopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentAtTenantScopeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDeploymentAtTenantScopeResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeploymentAtTenantScopeResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDeploymentAtTenantScopeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentAtTenantScopeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDeploymentAtTenantScopeResultOutput) Properties() DeploymentPropertiesExtendedResponseOutput {
	return o.ApplyT(func(v LookupDeploymentAtTenantScopeResult) DeploymentPropertiesExtendedResponse { return v.Properties }).(DeploymentPropertiesExtendedResponseOutput)
}

func (o LookupDeploymentAtTenantScopeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentAtTenantScopeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeploymentAtTenantScopeResultOutput{})
}

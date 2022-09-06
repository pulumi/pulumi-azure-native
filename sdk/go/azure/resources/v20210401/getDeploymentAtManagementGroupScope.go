


package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDeploymentAtManagementGroupScope(ctx *pulumi.Context, args *LookupDeploymentAtManagementGroupScopeArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentAtManagementGroupScopeResult, error) {
	var rv LookupDeploymentAtManagementGroupScopeResult
	err := ctx.Invoke("azure-native:resources/v20210401:getDeploymentAtManagementGroupScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentAtManagementGroupScopeArgs struct {
	DeploymentName string `pulumi:"deploymentName"`
	GroupId        string `pulumi:"groupId"`
}


type LookupDeploymentAtManagementGroupScopeResult struct {
	Id         string                               `pulumi:"id"`
	Location   *string                              `pulumi:"location"`
	Name       string                               `pulumi:"name"`
	Properties DeploymentPropertiesExtendedResponse `pulumi:"properties"`
	Tags       map[string]string                    `pulumi:"tags"`
	Type       string                               `pulumi:"type"`
}

func LookupDeploymentAtManagementGroupScopeOutput(ctx *pulumi.Context, args LookupDeploymentAtManagementGroupScopeOutputArgs, opts ...pulumi.InvokeOption) LookupDeploymentAtManagementGroupScopeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeploymentAtManagementGroupScopeResult, error) {
			args := v.(LookupDeploymentAtManagementGroupScopeArgs)
			r, err := LookupDeploymentAtManagementGroupScope(ctx, &args, opts...)
			var s LookupDeploymentAtManagementGroupScopeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeploymentAtManagementGroupScopeResultOutput)
}

type LookupDeploymentAtManagementGroupScopeOutputArgs struct {
	DeploymentName pulumi.StringInput `pulumi:"deploymentName"`
	GroupId        pulumi.StringInput `pulumi:"groupId"`
}

func (LookupDeploymentAtManagementGroupScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentAtManagementGroupScopeArgs)(nil)).Elem()
}


type LookupDeploymentAtManagementGroupScopeResultOutput struct{ *pulumi.OutputState }

func (LookupDeploymentAtManagementGroupScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentAtManagementGroupScopeResult)(nil)).Elem()
}

func (o LookupDeploymentAtManagementGroupScopeResultOutput) ToLookupDeploymentAtManagementGroupScopeResultOutput() LookupDeploymentAtManagementGroupScopeResultOutput {
	return o
}

func (o LookupDeploymentAtManagementGroupScopeResultOutput) ToLookupDeploymentAtManagementGroupScopeResultOutputWithContext(ctx context.Context) LookupDeploymentAtManagementGroupScopeResultOutput {
	return o
}

func (o LookupDeploymentAtManagementGroupScopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentAtManagementGroupScopeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDeploymentAtManagementGroupScopeResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeploymentAtManagementGroupScopeResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDeploymentAtManagementGroupScopeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentAtManagementGroupScopeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDeploymentAtManagementGroupScopeResultOutput) Properties() DeploymentPropertiesExtendedResponseOutput {
	return o.ApplyT(func(v LookupDeploymentAtManagementGroupScopeResult) DeploymentPropertiesExtendedResponse {
		return v.Properties
	}).(DeploymentPropertiesExtendedResponseOutput)
}

func (o LookupDeploymentAtManagementGroupScopeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDeploymentAtManagementGroupScopeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDeploymentAtManagementGroupScopeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentAtManagementGroupScopeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeploymentAtManagementGroupScopeResultOutput{})
}

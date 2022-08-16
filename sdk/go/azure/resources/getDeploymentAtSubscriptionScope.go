


package resources

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDeploymentAtSubscriptionScope(ctx *pulumi.Context, args *LookupDeploymentAtSubscriptionScopeArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentAtSubscriptionScopeResult, error) {
	var rv LookupDeploymentAtSubscriptionScopeResult
	err := ctx.Invoke("azure-native:resources:getDeploymentAtSubscriptionScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentAtSubscriptionScopeArgs struct {
	DeploymentName string `pulumi:"deploymentName"`
}


type LookupDeploymentAtSubscriptionScopeResult struct {
	Id         string                               `pulumi:"id"`
	Location   *string                              `pulumi:"location"`
	Name       string                               `pulumi:"name"`
	Properties DeploymentPropertiesExtendedResponse `pulumi:"properties"`
	Tags       map[string]string                    `pulumi:"tags"`
	Type       string                               `pulumi:"type"`
}

func LookupDeploymentAtSubscriptionScopeOutput(ctx *pulumi.Context, args LookupDeploymentAtSubscriptionScopeOutputArgs, opts ...pulumi.InvokeOption) LookupDeploymentAtSubscriptionScopeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeploymentAtSubscriptionScopeResult, error) {
			args := v.(LookupDeploymentAtSubscriptionScopeArgs)
			r, err := LookupDeploymentAtSubscriptionScope(ctx, &args, opts...)
			var s LookupDeploymentAtSubscriptionScopeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeploymentAtSubscriptionScopeResultOutput)
}

type LookupDeploymentAtSubscriptionScopeOutputArgs struct {
	DeploymentName pulumi.StringInput `pulumi:"deploymentName"`
}

func (LookupDeploymentAtSubscriptionScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentAtSubscriptionScopeArgs)(nil)).Elem()
}


type LookupDeploymentAtSubscriptionScopeResultOutput struct{ *pulumi.OutputState }

func (LookupDeploymentAtSubscriptionScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentAtSubscriptionScopeResult)(nil)).Elem()
}

func (o LookupDeploymentAtSubscriptionScopeResultOutput) ToLookupDeploymentAtSubscriptionScopeResultOutput() LookupDeploymentAtSubscriptionScopeResultOutput {
	return o
}

func (o LookupDeploymentAtSubscriptionScopeResultOutput) ToLookupDeploymentAtSubscriptionScopeResultOutputWithContext(ctx context.Context) LookupDeploymentAtSubscriptionScopeResultOutput {
	return o
}

func (o LookupDeploymentAtSubscriptionScopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentAtSubscriptionScopeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDeploymentAtSubscriptionScopeResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeploymentAtSubscriptionScopeResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDeploymentAtSubscriptionScopeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentAtSubscriptionScopeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDeploymentAtSubscriptionScopeResultOutput) Properties() DeploymentPropertiesExtendedResponseOutput {
	return o.ApplyT(func(v LookupDeploymentAtSubscriptionScopeResult) DeploymentPropertiesExtendedResponse {
		return v.Properties
	}).(DeploymentPropertiesExtendedResponseOutput)
}

func (o LookupDeploymentAtSubscriptionScopeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDeploymentAtSubscriptionScopeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDeploymentAtSubscriptionScopeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentAtSubscriptionScopeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeploymentAtSubscriptionScopeResultOutput{})
}

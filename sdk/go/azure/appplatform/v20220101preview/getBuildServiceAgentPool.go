


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBuildServiceAgentPool(ctx *pulumi.Context, args *LookupBuildServiceAgentPoolArgs, opts ...pulumi.InvokeOption) (*LookupBuildServiceAgentPoolResult, error) {
	var rv LookupBuildServiceAgentPoolResult
	err := ctx.Invoke("azure-native:appplatform/v20220101preview:getBuildServiceAgentPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBuildServiceAgentPoolArgs struct {
	AgentPoolName     string `pulumi:"agentPoolName"`
	BuildServiceName  string `pulumi:"buildServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupBuildServiceAgentPoolResult struct {
	Id         string                                  `pulumi:"id"`
	Name       string                                  `pulumi:"name"`
	Properties BuildServiceAgentPoolPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                      `pulumi:"systemData"`
	Type       string                                  `pulumi:"type"`
}

func LookupBuildServiceAgentPoolOutput(ctx *pulumi.Context, args LookupBuildServiceAgentPoolOutputArgs, opts ...pulumi.InvokeOption) LookupBuildServiceAgentPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBuildServiceAgentPoolResult, error) {
			args := v.(LookupBuildServiceAgentPoolArgs)
			r, err := LookupBuildServiceAgentPool(ctx, &args, opts...)
			var s LookupBuildServiceAgentPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBuildServiceAgentPoolResultOutput)
}

type LookupBuildServiceAgentPoolOutputArgs struct {
	AgentPoolName     pulumi.StringInput `pulumi:"agentPoolName"`
	BuildServiceName  pulumi.StringInput `pulumi:"buildServiceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupBuildServiceAgentPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBuildServiceAgentPoolArgs)(nil)).Elem()
}


type LookupBuildServiceAgentPoolResultOutput struct{ *pulumi.OutputState }

func (LookupBuildServiceAgentPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBuildServiceAgentPoolResult)(nil)).Elem()
}

func (o LookupBuildServiceAgentPoolResultOutput) ToLookupBuildServiceAgentPoolResultOutput() LookupBuildServiceAgentPoolResultOutput {
	return o
}

func (o LookupBuildServiceAgentPoolResultOutput) ToLookupBuildServiceAgentPoolResultOutputWithContext(ctx context.Context) LookupBuildServiceAgentPoolResultOutput {
	return o
}

func (o LookupBuildServiceAgentPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildServiceAgentPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBuildServiceAgentPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildServiceAgentPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBuildServiceAgentPoolResultOutput) Properties() BuildServiceAgentPoolPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBuildServiceAgentPoolResult) BuildServiceAgentPoolPropertiesResponse { return v.Properties }).(BuildServiceAgentPoolPropertiesResponseOutput)
}

func (o LookupBuildServiceAgentPoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBuildServiceAgentPoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupBuildServiceAgentPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildServiceAgentPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBuildServiceAgentPoolResultOutput{})
}

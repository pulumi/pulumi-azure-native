


package v20220401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBuildServiceBuilder(ctx *pulumi.Context, args *LookupBuildServiceBuilderArgs, opts ...pulumi.InvokeOption) (*LookupBuildServiceBuilderResult, error) {
	var rv LookupBuildServiceBuilderResult
	err := ctx.Invoke("azure-native:appplatform/v20220401:getBuildServiceBuilder", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBuildServiceBuilderArgs struct {
	BuildServiceName  string `pulumi:"buildServiceName"`
	BuilderName       string `pulumi:"builderName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupBuildServiceBuilderResult struct {
	Id         string                    `pulumi:"id"`
	Name       string                    `pulumi:"name"`
	Properties BuilderPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse        `pulumi:"systemData"`
	Type       string                    `pulumi:"type"`
}

func LookupBuildServiceBuilderOutput(ctx *pulumi.Context, args LookupBuildServiceBuilderOutputArgs, opts ...pulumi.InvokeOption) LookupBuildServiceBuilderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBuildServiceBuilderResult, error) {
			args := v.(LookupBuildServiceBuilderArgs)
			r, err := LookupBuildServiceBuilder(ctx, &args, opts...)
			var s LookupBuildServiceBuilderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBuildServiceBuilderResultOutput)
}

type LookupBuildServiceBuilderOutputArgs struct {
	BuildServiceName  pulumi.StringInput `pulumi:"buildServiceName"`
	BuilderName       pulumi.StringInput `pulumi:"builderName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupBuildServiceBuilderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBuildServiceBuilderArgs)(nil)).Elem()
}


type LookupBuildServiceBuilderResultOutput struct{ *pulumi.OutputState }

func (LookupBuildServiceBuilderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBuildServiceBuilderResult)(nil)).Elem()
}

func (o LookupBuildServiceBuilderResultOutput) ToLookupBuildServiceBuilderResultOutput() LookupBuildServiceBuilderResultOutput {
	return o
}

func (o LookupBuildServiceBuilderResultOutput) ToLookupBuildServiceBuilderResultOutputWithContext(ctx context.Context) LookupBuildServiceBuilderResultOutput {
	return o
}

func (o LookupBuildServiceBuilderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildServiceBuilderResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBuildServiceBuilderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildServiceBuilderResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBuildServiceBuilderResultOutput) Properties() BuilderPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBuildServiceBuilderResult) BuilderPropertiesResponse { return v.Properties }).(BuilderPropertiesResponseOutput)
}

func (o LookupBuildServiceBuilderResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBuildServiceBuilderResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupBuildServiceBuilderResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildServiceBuilderResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBuildServiceBuilderResultOutput{})
}

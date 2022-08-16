


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBuildpackBinding(ctx *pulumi.Context, args *LookupBuildpackBindingArgs, opts ...pulumi.InvokeOption) (*LookupBuildpackBindingResult, error) {
	var rv LookupBuildpackBindingResult
	err := ctx.Invoke("azure-native:appplatform/v20220101preview:getBuildpackBinding", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBuildpackBindingArgs struct {
	BuildServiceName     string `pulumi:"buildServiceName"`
	BuilderName          string `pulumi:"builderName"`
	BuildpackBindingName string `pulumi:"buildpackBindingName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	ServiceName          string `pulumi:"serviceName"`
}


type LookupBuildpackBindingResult struct {
	Id         string                             `pulumi:"id"`
	Name       string                             `pulumi:"name"`
	Properties BuildpackBindingPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                 `pulumi:"systemData"`
	Type       string                             `pulumi:"type"`
}

func LookupBuildpackBindingOutput(ctx *pulumi.Context, args LookupBuildpackBindingOutputArgs, opts ...pulumi.InvokeOption) LookupBuildpackBindingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBuildpackBindingResult, error) {
			args := v.(LookupBuildpackBindingArgs)
			r, err := LookupBuildpackBinding(ctx, &args, opts...)
			var s LookupBuildpackBindingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBuildpackBindingResultOutput)
}

type LookupBuildpackBindingOutputArgs struct {
	BuildServiceName     pulumi.StringInput `pulumi:"buildServiceName"`
	BuilderName          pulumi.StringInput `pulumi:"builderName"`
	BuildpackBindingName pulumi.StringInput `pulumi:"buildpackBindingName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName          pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupBuildpackBindingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBuildpackBindingArgs)(nil)).Elem()
}


type LookupBuildpackBindingResultOutput struct{ *pulumi.OutputState }

func (LookupBuildpackBindingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBuildpackBindingResult)(nil)).Elem()
}

func (o LookupBuildpackBindingResultOutput) ToLookupBuildpackBindingResultOutput() LookupBuildpackBindingResultOutput {
	return o
}

func (o LookupBuildpackBindingResultOutput) ToLookupBuildpackBindingResultOutputWithContext(ctx context.Context) LookupBuildpackBindingResultOutput {
	return o
}

func (o LookupBuildpackBindingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildpackBindingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBuildpackBindingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildpackBindingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBuildpackBindingResultOutput) Properties() BuildpackBindingPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBuildpackBindingResult) BuildpackBindingPropertiesResponse { return v.Properties }).(BuildpackBindingPropertiesResponseOutput)
}

func (o LookupBuildpackBindingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBuildpackBindingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupBuildpackBindingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildpackBindingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBuildpackBindingResultOutput{})
}

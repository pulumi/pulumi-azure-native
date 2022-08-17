


package v20220701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExperiment(ctx *pulumi.Context, args *LookupExperimentArgs, opts ...pulumi.InvokeOption) (*LookupExperimentResult, error) {
	var rv LookupExperimentResult
	err := ctx.Invoke("azure-native:chaos/v20220701preview:getExperiment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExperimentArgs struct {
	ExperimentName    string `pulumi:"experimentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupExperimentResult struct {
	Id         string                       `pulumi:"id"`
	Identity   *ResourceIdentityResponse    `pulumi:"identity"`
	Location   string                       `pulumi:"location"`
	Name       string                       `pulumi:"name"`
	Properties ExperimentPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse           `pulumi:"systemData"`
	Tags       map[string]string            `pulumi:"tags"`
	Type       string                       `pulumi:"type"`
}

func LookupExperimentOutput(ctx *pulumi.Context, args LookupExperimentOutputArgs, opts ...pulumi.InvokeOption) LookupExperimentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExperimentResult, error) {
			args := v.(LookupExperimentArgs)
			r, err := LookupExperiment(ctx, &args, opts...)
			var s LookupExperimentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExperimentResultOutput)
}

type LookupExperimentOutputArgs struct {
	ExperimentName    pulumi.StringInput `pulumi:"experimentName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExperimentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExperimentArgs)(nil)).Elem()
}


type LookupExperimentResultOutput struct{ *pulumi.OutputState }

func (LookupExperimentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExperimentResult)(nil)).Elem()
}

func (o LookupExperimentResultOutput) ToLookupExperimentResultOutput() LookupExperimentResultOutput {
	return o
}

func (o LookupExperimentResultOutput) ToLookupExperimentResultOutputWithContext(ctx context.Context) LookupExperimentResultOutput {
	return o
}

func (o LookupExperimentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupExperimentResultOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o LookupExperimentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupExperimentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupExperimentResultOutput) Properties() ExperimentPropertiesResponseOutput {
	return o.ApplyT(func(v LookupExperimentResult) ExperimentPropertiesResponse { return v.Properties }).(ExperimentPropertiesResponseOutput)
}

func (o LookupExperimentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupExperimentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupExperimentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExperimentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupExperimentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExperimentResultOutput{})
}

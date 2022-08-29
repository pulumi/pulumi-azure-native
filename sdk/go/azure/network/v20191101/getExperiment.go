


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExperiment(ctx *pulumi.Context, args *LookupExperimentArgs, opts ...pulumi.InvokeOption) (*LookupExperimentResult, error) {
	var rv LookupExperimentResult
	err := ctx.Invoke("azure-native:network/v20191101:getExperiment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExperimentArgs struct {
	ExperimentName    string `pulumi:"experimentName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupExperimentResult struct {
	Description   *string                     `pulumi:"description"`
	EnabledState  *string                     `pulumi:"enabledState"`
	EndpointA     *ExperimentEndpointResponse `pulumi:"endpointA"`
	EndpointB     *ExperimentEndpointResponse `pulumi:"endpointB"`
	Id            string                      `pulumi:"id"`
	Location      *string                     `pulumi:"location"`
	Name          string                      `pulumi:"name"`
	ResourceState string                      `pulumi:"resourceState"`
	ScriptFileUri string                      `pulumi:"scriptFileUri"`
	Status        string                      `pulumi:"status"`
	Tags          map[string]string           `pulumi:"tags"`
	Type          string                      `pulumi:"type"`
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
	ProfileName       pulumi.StringInput `pulumi:"profileName"`
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

func (o LookupExperimentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupExperimentResultOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o LookupExperimentResultOutput) EndpointA() ExperimentEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *ExperimentEndpointResponse { return v.EndpointA }).(ExperimentEndpointResponsePtrOutput)
}

func (o LookupExperimentResultOutput) EndpointB() ExperimentEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *ExperimentEndpointResponse { return v.EndpointB }).(ExperimentEndpointResponsePtrOutput)
}

func (o LookupExperimentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupExperimentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupExperimentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupExperimentResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o LookupExperimentResultOutput) ScriptFileUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.ScriptFileUri }).(pulumi.StringOutput)
}

func (o LookupExperimentResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.Status }).(pulumi.StringOutput)
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

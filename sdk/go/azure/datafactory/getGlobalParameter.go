


package datafactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGlobalParameter(ctx *pulumi.Context, args *LookupGlobalParameterArgs, opts ...pulumi.InvokeOption) (*LookupGlobalParameterResult, error) {
	var rv LookupGlobalParameterResult
	err := ctx.Invoke("azure-native:datafactory:getGlobalParameter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGlobalParameterArgs struct {
	FactoryName         string `pulumi:"factoryName"`
	GlobalParameterName string `pulumi:"globalParameterName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupGlobalParameterResult struct {
	Etag       string                                          `pulumi:"etag"`
	Id         string                                          `pulumi:"id"`
	Name       string                                          `pulumi:"name"`
	Properties map[string]GlobalParameterSpecificationResponse `pulumi:"properties"`
	Type       string                                          `pulumi:"type"`
}

func LookupGlobalParameterOutput(ctx *pulumi.Context, args LookupGlobalParameterOutputArgs, opts ...pulumi.InvokeOption) LookupGlobalParameterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGlobalParameterResult, error) {
			args := v.(LookupGlobalParameterArgs)
			r, err := LookupGlobalParameter(ctx, &args, opts...)
			var s LookupGlobalParameterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGlobalParameterResultOutput)
}

type LookupGlobalParameterOutputArgs struct {
	FactoryName         pulumi.StringInput `pulumi:"factoryName"`
	GlobalParameterName pulumi.StringInput `pulumi:"globalParameterName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGlobalParameterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalParameterArgs)(nil)).Elem()
}


type LookupGlobalParameterResultOutput struct{ *pulumi.OutputState }

func (LookupGlobalParameterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalParameterResult)(nil)).Elem()
}

func (o LookupGlobalParameterResultOutput) ToLookupGlobalParameterResultOutput() LookupGlobalParameterResultOutput {
	return o
}

func (o LookupGlobalParameterResultOutput) ToLookupGlobalParameterResultOutputWithContext(ctx context.Context) LookupGlobalParameterResultOutput {
	return o
}

func (o LookupGlobalParameterResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalParameterResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupGlobalParameterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalParameterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGlobalParameterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalParameterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGlobalParameterResultOutput) Properties() GlobalParameterSpecificationResponseMapOutput {
	return o.ApplyT(func(v LookupGlobalParameterResult) map[string]GlobalParameterSpecificationResponse {
		return v.Properties
	}).(GlobalParameterSpecificationResponseMapOutput)
}

func (o LookupGlobalParameterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalParameterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGlobalParameterResultOutput{})
}

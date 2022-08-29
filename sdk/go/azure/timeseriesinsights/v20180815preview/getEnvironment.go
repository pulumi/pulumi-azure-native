


package v20180815preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupEnvironment(ctx *pulumi.Context, args *LookupEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentResult, error) {
	var rv LookupEnvironmentResult
	err := ctx.Invoke("azure-native:timeseriesinsights/v20180815preview:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentArgs struct {
	EnvironmentName   string  `pulumi:"environmentName"`
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupEnvironmentResult struct {
	Id       string            `pulumi:"id"`
	Kind     string            `pulumi:"kind"`
	Location string            `pulumi:"location"`
	Name     string            `pulumi:"name"`
	Sku      SkuResponse       `pulumi:"sku"`
	Tags     map[string]string `pulumi:"tags"`
	Type     string            `pulumi:"type"`
}

func LookupEnvironmentOutput(ctx *pulumi.Context, args LookupEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentResult, error) {
			args := v.(LookupEnvironmentArgs)
			r, err := LookupEnvironment(ctx, &args, opts...)
			var s LookupEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentResultOutput)
}

type LookupEnvironmentOutputArgs struct {
	EnvironmentName   pulumi.StringInput    `pulumi:"environmentName"`
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentArgs)(nil)).Elem()
}


type LookupEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentResult)(nil)).Elem()
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutput() LookupEnvironmentResultOutput {
	return o
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutputWithContext(ctx context.Context) LookupEnvironmentResultOutput {
	return o
}

func (o LookupEnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupEnvironmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupEnvironmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentResultOutput{})
}

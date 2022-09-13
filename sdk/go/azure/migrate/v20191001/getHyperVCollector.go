


package v20191001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHyperVCollector(ctx *pulumi.Context, args *LookupHyperVCollectorArgs, opts ...pulumi.InvokeOption) (*LookupHyperVCollectorResult, error) {
	var rv LookupHyperVCollectorResult
	err := ctx.Invoke("azure-native:migrate/v20191001:getHyperVCollector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHyperVCollectorArgs struct {
	HyperVCollectorName string `pulumi:"hyperVCollectorName"`
	ProjectName         string `pulumi:"projectName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}

type LookupHyperVCollectorResult struct {
	ETag       *string                     `pulumi:"eTag"`
	Id         string                      `pulumi:"id"`
	Name       string                      `pulumi:"name"`
	Properties CollectorPropertiesResponse `pulumi:"properties"`
	Type       string                      `pulumi:"type"`
}

func LookupHyperVCollectorOutput(ctx *pulumi.Context, args LookupHyperVCollectorOutputArgs, opts ...pulumi.InvokeOption) LookupHyperVCollectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHyperVCollectorResult, error) {
			args := v.(LookupHyperVCollectorArgs)
			r, err := LookupHyperVCollector(ctx, &args, opts...)
			var s LookupHyperVCollectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHyperVCollectorResultOutput)
}

type LookupHyperVCollectorOutputArgs struct {
	HyperVCollectorName pulumi.StringInput `pulumi:"hyperVCollectorName"`
	ProjectName         pulumi.StringInput `pulumi:"projectName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupHyperVCollectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHyperVCollectorArgs)(nil)).Elem()
}

type LookupHyperVCollectorResultOutput struct{ *pulumi.OutputState }

func (LookupHyperVCollectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHyperVCollectorResult)(nil)).Elem()
}

func (o LookupHyperVCollectorResultOutput) ToLookupHyperVCollectorResultOutput() LookupHyperVCollectorResultOutput {
	return o
}

func (o LookupHyperVCollectorResultOutput) ToLookupHyperVCollectorResultOutputWithContext(ctx context.Context) LookupHyperVCollectorResultOutput {
	return o
}

func (o LookupHyperVCollectorResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHyperVCollectorResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupHyperVCollectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHyperVCollectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHyperVCollectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHyperVCollectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupHyperVCollectorResultOutput) Properties() CollectorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupHyperVCollectorResult) CollectorPropertiesResponse { return v.Properties }).(CollectorPropertiesResponseOutput)
}

func (o LookupHyperVCollectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHyperVCollectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHyperVCollectorResultOutput{})
}

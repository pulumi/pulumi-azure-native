


package v20191001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupImportCollector(ctx *pulumi.Context, args *LookupImportCollectorArgs, opts ...pulumi.InvokeOption) (*LookupImportCollectorResult, error) {
	var rv LookupImportCollectorResult
	err := ctx.Invoke("azure-native:migrate/v20191001:getImportCollector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupImportCollectorArgs struct {
	ImportCollectorName string `pulumi:"importCollectorName"`
	ProjectName         string `pulumi:"projectName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}

type LookupImportCollectorResult struct {
	ETag       *string                           `pulumi:"eTag"`
	Id         string                            `pulumi:"id"`
	Name       string                            `pulumi:"name"`
	Properties ImportCollectorPropertiesResponse `pulumi:"properties"`
	Type       string                            `pulumi:"type"`
}

func LookupImportCollectorOutput(ctx *pulumi.Context, args LookupImportCollectorOutputArgs, opts ...pulumi.InvokeOption) LookupImportCollectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupImportCollectorResult, error) {
			args := v.(LookupImportCollectorArgs)
			r, err := LookupImportCollector(ctx, &args, opts...)
			var s LookupImportCollectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupImportCollectorResultOutput)
}

type LookupImportCollectorOutputArgs struct {
	ImportCollectorName pulumi.StringInput `pulumi:"importCollectorName"`
	ProjectName         pulumi.StringInput `pulumi:"projectName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupImportCollectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImportCollectorArgs)(nil)).Elem()
}

type LookupImportCollectorResultOutput struct{ *pulumi.OutputState }

func (LookupImportCollectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImportCollectorResult)(nil)).Elem()
}

func (o LookupImportCollectorResultOutput) ToLookupImportCollectorResultOutput() LookupImportCollectorResultOutput {
	return o
}

func (o LookupImportCollectorResultOutput) ToLookupImportCollectorResultOutputWithContext(ctx context.Context) LookupImportCollectorResultOutput {
	return o
}

func (o LookupImportCollectorResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImportCollectorResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupImportCollectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportCollectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupImportCollectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportCollectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupImportCollectorResultOutput) Properties() ImportCollectorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupImportCollectorResult) ImportCollectorPropertiesResponse { return v.Properties }).(ImportCollectorPropertiesResponseOutput)
}

func (o LookupImportCollectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportCollectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImportCollectorResultOutput{})
}

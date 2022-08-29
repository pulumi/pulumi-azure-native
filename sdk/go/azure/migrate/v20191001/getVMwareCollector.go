


package v20191001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVMwareCollector(ctx *pulumi.Context, args *LookupVMwareCollectorArgs, opts ...pulumi.InvokeOption) (*LookupVMwareCollectorResult, error) {
	var rv LookupVMwareCollectorResult
	err := ctx.Invoke("azure-native:migrate/v20191001:getVMwareCollector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVMwareCollectorArgs struct {
	ProjectName         string `pulumi:"projectName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	VmWareCollectorName string `pulumi:"vmWareCollectorName"`
}

type LookupVMwareCollectorResult struct {
	ETag       *string                     `pulumi:"eTag"`
	Id         string                      `pulumi:"id"`
	Name       string                      `pulumi:"name"`
	Properties CollectorPropertiesResponse `pulumi:"properties"`
	Type       string                      `pulumi:"type"`
}

func LookupVMwareCollectorOutput(ctx *pulumi.Context, args LookupVMwareCollectorOutputArgs, opts ...pulumi.InvokeOption) LookupVMwareCollectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVMwareCollectorResult, error) {
			args := v.(LookupVMwareCollectorArgs)
			r, err := LookupVMwareCollector(ctx, &args, opts...)
			var s LookupVMwareCollectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVMwareCollectorResultOutput)
}

type LookupVMwareCollectorOutputArgs struct {
	ProjectName         pulumi.StringInput `pulumi:"projectName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	VmWareCollectorName pulumi.StringInput `pulumi:"vmWareCollectorName"`
}

func (LookupVMwareCollectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVMwareCollectorArgs)(nil)).Elem()
}

type LookupVMwareCollectorResultOutput struct{ *pulumi.OutputState }

func (LookupVMwareCollectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVMwareCollectorResult)(nil)).Elem()
}

func (o LookupVMwareCollectorResultOutput) ToLookupVMwareCollectorResultOutput() LookupVMwareCollectorResultOutput {
	return o
}

func (o LookupVMwareCollectorResultOutput) ToLookupVMwareCollectorResultOutputWithContext(ctx context.Context) LookupVMwareCollectorResultOutput {
	return o
}

func (o LookupVMwareCollectorResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVMwareCollectorResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupVMwareCollectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVMwareCollectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVMwareCollectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVMwareCollectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVMwareCollectorResultOutput) Properties() CollectorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupVMwareCollectorResult) CollectorPropertiesResponse { return v.Properties }).(CollectorPropertiesResponseOutput)
}

func (o LookupVMwareCollectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVMwareCollectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVMwareCollectorResultOutput{})
}

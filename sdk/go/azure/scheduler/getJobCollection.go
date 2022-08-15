


package scheduler

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJobCollection(ctx *pulumi.Context, args *LookupJobCollectionArgs, opts ...pulumi.InvokeOption) (*LookupJobCollectionResult, error) {
	var rv LookupJobCollectionResult
	err := ctx.Invoke("azure-native:scheduler:getJobCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobCollectionArgs struct {
	JobCollectionName string `pulumi:"jobCollectionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupJobCollectionResult struct {
	Id         string                          `pulumi:"id"`
	Location   *string                         `pulumi:"location"`
	Name       *string                         `pulumi:"name"`
	Properties JobCollectionPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string               `pulumi:"tags"`
	Type       string                          `pulumi:"type"`
}

func LookupJobCollectionOutput(ctx *pulumi.Context, args LookupJobCollectionOutputArgs, opts ...pulumi.InvokeOption) LookupJobCollectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobCollectionResult, error) {
			args := v.(LookupJobCollectionArgs)
			r, err := LookupJobCollection(ctx, &args, opts...)
			var s LookupJobCollectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobCollectionResultOutput)
}

type LookupJobCollectionOutputArgs struct {
	JobCollectionName pulumi.StringInput `pulumi:"jobCollectionName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupJobCollectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobCollectionArgs)(nil)).Elem()
}

type LookupJobCollectionResultOutput struct{ *pulumi.OutputState }

func (LookupJobCollectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobCollectionResult)(nil)).Elem()
}

func (o LookupJobCollectionResultOutput) ToLookupJobCollectionResultOutput() LookupJobCollectionResultOutput {
	return o
}

func (o LookupJobCollectionResultOutput) ToLookupJobCollectionResultOutputWithContext(ctx context.Context) LookupJobCollectionResultOutput {
	return o
}

func (o LookupJobCollectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobCollectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupJobCollectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobCollectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupJobCollectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobCollectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupJobCollectionResultOutput) Properties() JobCollectionPropertiesResponseOutput {
	return o.ApplyT(func(v LookupJobCollectionResult) JobCollectionPropertiesResponse { return v.Properties }).(JobCollectionPropertiesResponseOutput)
}

func (o LookupJobCollectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupJobCollectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupJobCollectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobCollectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobCollectionResultOutput{})
}

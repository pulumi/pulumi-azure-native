


package v20160712preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSuppression(ctx *pulumi.Context, args *LookupSuppressionArgs, opts ...pulumi.InvokeOption) (*LookupSuppressionResult, error) {
	var rv LookupSuppressionResult
	err := ctx.Invoke("azure-native:advisor/v20160712preview:getSuppression", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSuppressionArgs struct {
	Name             string `pulumi:"name"`
	RecommendationId string `pulumi:"recommendationId"`
	ResourceUri      string `pulumi:"resourceUri"`
}


type LookupSuppressionResult struct {
	Id            string            `pulumi:"id"`
	Location      *string           `pulumi:"location"`
	Name          string            `pulumi:"name"`
	SuppressionId *string           `pulumi:"suppressionId"`
	Tags          map[string]string `pulumi:"tags"`
	Ttl           *string           `pulumi:"ttl"`
	Type          string            `pulumi:"type"`
}

func LookupSuppressionOutput(ctx *pulumi.Context, args LookupSuppressionOutputArgs, opts ...pulumi.InvokeOption) LookupSuppressionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSuppressionResult, error) {
			args := v.(LookupSuppressionArgs)
			r, err := LookupSuppression(ctx, &args, opts...)
			var s LookupSuppressionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSuppressionResultOutput)
}

type LookupSuppressionOutputArgs struct {
	Name             pulumi.StringInput `pulumi:"name"`
	RecommendationId pulumi.StringInput `pulumi:"recommendationId"`
	ResourceUri      pulumi.StringInput `pulumi:"resourceUri"`
}

func (LookupSuppressionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSuppressionArgs)(nil)).Elem()
}


type LookupSuppressionResultOutput struct{ *pulumi.OutputState }

func (LookupSuppressionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSuppressionResult)(nil)).Elem()
}

func (o LookupSuppressionResultOutput) ToLookupSuppressionResultOutput() LookupSuppressionResultOutput {
	return o
}

func (o LookupSuppressionResultOutput) ToLookupSuppressionResultOutputWithContext(ctx context.Context) LookupSuppressionResultOutput {
	return o
}

func (o LookupSuppressionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSuppressionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSuppressionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSuppressionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupSuppressionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSuppressionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSuppressionResultOutput) SuppressionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSuppressionResult) *string { return v.SuppressionId }).(pulumi.StringPtrOutput)
}

func (o LookupSuppressionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSuppressionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSuppressionResultOutput) Ttl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSuppressionResult) *string { return v.Ttl }).(pulumi.StringPtrOutput)
}

func (o LookupSuppressionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSuppressionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSuppressionResultOutput{})
}

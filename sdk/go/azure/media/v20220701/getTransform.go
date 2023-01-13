


package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTransform(ctx *pulumi.Context, args *LookupTransformArgs, opts ...pulumi.InvokeOption) (*LookupTransformResult, error) {
	var rv LookupTransformResult
	err := ctx.Invoke("azure-native:media/v20220701:getTransform", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTransformArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TransformName     string `pulumi:"transformName"`
}


type LookupTransformResult struct {
	Created      string                    `pulumi:"created"`
	Description  *string                   `pulumi:"description"`
	Id           string                    `pulumi:"id"`
	LastModified string                    `pulumi:"lastModified"`
	Name         string                    `pulumi:"name"`
	Outputs      []TransformOutputResponse `pulumi:"outputs"`
	SystemData   SystemDataResponse        `pulumi:"systemData"`
	Type         string                    `pulumi:"type"`
}

func LookupTransformOutput(ctx *pulumi.Context, args LookupTransformOutputArgs, opts ...pulumi.InvokeOption) LookupTransformResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTransformResult, error) {
			args := v.(LookupTransformArgs)
			r, err := LookupTransform(ctx, &args, opts...)
			var s LookupTransformResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTransformResultOutput)
}

type LookupTransformOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TransformName     pulumi.StringInput `pulumi:"transformName"`
}

func (LookupTransformOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTransformArgs)(nil)).Elem()
}


type LookupTransformResultOutput struct{ *pulumi.OutputState }

func (LookupTransformResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTransformResult)(nil)).Elem()
}

func (o LookupTransformResultOutput) ToLookupTransformResultOutput() LookupTransformResultOutput {
	return o
}

func (o LookupTransformResultOutput) ToLookupTransformResultOutputWithContext(ctx context.Context) LookupTransformResultOutput {
	return o
}

func (o LookupTransformResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransformResult) string { return v.Created }).(pulumi.StringOutput)
}

func (o LookupTransformResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTransformResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupTransformResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransformResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTransformResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransformResult) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o LookupTransformResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransformResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTransformResultOutput) Outputs() TransformOutputResponseArrayOutput {
	return o.ApplyT(func(v LookupTransformResult) []TransformOutputResponse { return v.Outputs }).(TransformOutputResponseArrayOutput)
}

func (o LookupTransformResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTransformResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupTransformResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransformResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTransformResultOutput{})
}

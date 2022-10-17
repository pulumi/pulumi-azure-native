


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTagInheritanceSetting(ctx *pulumi.Context, args *LookupTagInheritanceSettingArgs, opts ...pulumi.InvokeOption) (*LookupTagInheritanceSettingResult, error) {
	var rv LookupTagInheritanceSettingResult
	err := ctx.Invoke("azure-native:costmanagement/v20221001preview:getTagInheritanceSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagInheritanceSettingArgs struct {
	Scope string `pulumi:"scope"`
	Type  string `pulumi:"type"`
}


type LookupTagInheritanceSettingResult struct {
	ETag       *string                          `pulumi:"eTag"`
	Id         string                           `pulumi:"id"`
	Kind       string                           `pulumi:"kind"`
	Name       string                           `pulumi:"name"`
	Properties TagInheritancePropertiesResponse `pulumi:"properties"`
	Type       string                           `pulumi:"type"`
}

func LookupTagInheritanceSettingOutput(ctx *pulumi.Context, args LookupTagInheritanceSettingOutputArgs, opts ...pulumi.InvokeOption) LookupTagInheritanceSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagInheritanceSettingResult, error) {
			args := v.(LookupTagInheritanceSettingArgs)
			r, err := LookupTagInheritanceSetting(ctx, &args, opts...)
			var s LookupTagInheritanceSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTagInheritanceSettingResultOutput)
}

type LookupTagInheritanceSettingOutputArgs struct {
	Scope pulumi.StringInput `pulumi:"scope"`
	Type  pulumi.StringInput `pulumi:"type"`
}

func (LookupTagInheritanceSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagInheritanceSettingArgs)(nil)).Elem()
}


type LookupTagInheritanceSettingResultOutput struct{ *pulumi.OutputState }

func (LookupTagInheritanceSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagInheritanceSettingResult)(nil)).Elem()
}

func (o LookupTagInheritanceSettingResultOutput) ToLookupTagInheritanceSettingResultOutput() LookupTagInheritanceSettingResultOutput {
	return o
}

func (o LookupTagInheritanceSettingResultOutput) ToLookupTagInheritanceSettingResultOutputWithContext(ctx context.Context) LookupTagInheritanceSettingResultOutput {
	return o
}

func (o LookupTagInheritanceSettingResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTagInheritanceSettingResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupTagInheritanceSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagInheritanceSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTagInheritanceSettingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagInheritanceSettingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupTagInheritanceSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagInheritanceSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTagInheritanceSettingResultOutput) Properties() TagInheritancePropertiesResponseOutput {
	return o.ApplyT(func(v LookupTagInheritanceSettingResult) TagInheritancePropertiesResponse { return v.Properties }).(TagInheritancePropertiesResponseOutput)
}

func (o LookupTagInheritanceSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagInheritanceSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagInheritanceSettingResultOutput{})
}

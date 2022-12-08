


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSettingByScope(ctx *pulumi.Context, args *LookupSettingByScopeArgs, opts ...pulumi.InvokeOption) (*LookupSettingByScopeResult, error) {
	var rv LookupSettingByScopeResult
	err := ctx.Invoke("azure-native:costmanagement/v20221001preview:getSettingByScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSettingByScopeArgs struct {
	Scope string `pulumi:"scope"`
	Type  string `pulumi:"type"`
}


type LookupSettingByScopeResult struct {
	ETag *string `pulumi:"eTag"`
	Id   string  `pulumi:"id"`
	Kind string  `pulumi:"kind"`
	Name string  `pulumi:"name"`
	Type string  `pulumi:"type"`
}

func LookupSettingByScopeOutput(ctx *pulumi.Context, args LookupSettingByScopeOutputArgs, opts ...pulumi.InvokeOption) LookupSettingByScopeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSettingByScopeResult, error) {
			args := v.(LookupSettingByScopeArgs)
			r, err := LookupSettingByScope(ctx, &args, opts...)
			var s LookupSettingByScopeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSettingByScopeResultOutput)
}

type LookupSettingByScopeOutputArgs struct {
	Scope pulumi.StringInput `pulumi:"scope"`
	Type  pulumi.StringInput `pulumi:"type"`
}

func (LookupSettingByScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSettingByScopeArgs)(nil)).Elem()
}


type LookupSettingByScopeResultOutput struct{ *pulumi.OutputState }

func (LookupSettingByScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSettingByScopeResult)(nil)).Elem()
}

func (o LookupSettingByScopeResultOutput) ToLookupSettingByScopeResultOutput() LookupSettingByScopeResultOutput {
	return o
}

func (o LookupSettingByScopeResultOutput) ToLookupSettingByScopeResultOutputWithContext(ctx context.Context) LookupSettingByScopeResultOutput {
	return o
}

func (o LookupSettingByScopeResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSettingByScopeResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupSettingByScopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSettingByScopeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSettingByScopeResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSettingByScopeResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupSettingByScopeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSettingByScopeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSettingByScopeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSettingByScopeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSettingByScopeResultOutput{})
}

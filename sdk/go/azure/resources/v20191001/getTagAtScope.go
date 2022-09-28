


package v20191001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTagAtScope(ctx *pulumi.Context, args *LookupTagAtScopeArgs, opts ...pulumi.InvokeOption) (*LookupTagAtScopeResult, error) {
	var rv LookupTagAtScopeResult
	err := ctx.Invoke("azure-native:resources/v20191001:getTagAtScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagAtScopeArgs struct {
	Scope string `pulumi:"scope"`
}


type LookupTagAtScopeResult struct {
	Id         string       `pulumi:"id"`
	Name       string       `pulumi:"name"`
	Properties TagsResponse `pulumi:"properties"`
	Type       string       `pulumi:"type"`
}

func LookupTagAtScopeOutput(ctx *pulumi.Context, args LookupTagAtScopeOutputArgs, opts ...pulumi.InvokeOption) LookupTagAtScopeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagAtScopeResult, error) {
			args := v.(LookupTagAtScopeArgs)
			r, err := LookupTagAtScope(ctx, &args, opts...)
			var s LookupTagAtScopeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTagAtScopeResultOutput)
}

type LookupTagAtScopeOutputArgs struct {
	Scope pulumi.StringInput `pulumi:"scope"`
}

func (LookupTagAtScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagAtScopeArgs)(nil)).Elem()
}


type LookupTagAtScopeResultOutput struct{ *pulumi.OutputState }

func (LookupTagAtScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagAtScopeResult)(nil)).Elem()
}

func (o LookupTagAtScopeResultOutput) ToLookupTagAtScopeResultOutput() LookupTagAtScopeResultOutput {
	return o
}

func (o LookupTagAtScopeResultOutput) ToLookupTagAtScopeResultOutputWithContext(ctx context.Context) LookupTagAtScopeResultOutput {
	return o
}

func (o LookupTagAtScopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagAtScopeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTagAtScopeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagAtScopeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTagAtScopeResultOutput) Properties() TagsResponseOutput {
	return o.ApplyT(func(v LookupTagAtScopeResult) TagsResponse { return v.Properties }).(TagsResponseOutput)
}

func (o LookupTagAtScopeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagAtScopeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagAtScopeResultOutput{})
}

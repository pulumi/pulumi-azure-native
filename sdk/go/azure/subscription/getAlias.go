


package subscription

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAlias(ctx *pulumi.Context, args *LookupAliasArgs, opts ...pulumi.InvokeOption) (*LookupAliasResult, error) {
	var rv LookupAliasResult
	err := ctx.Invoke("azure-native:subscription:getAlias", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAliasArgs struct {
	AliasName string `pulumi:"aliasName"`
}


type LookupAliasResult struct {
	Id         string                             `pulumi:"id"`
	Name       string                             `pulumi:"name"`
	Properties PutAliasResponsePropertiesResponse `pulumi:"properties"`
	Type       string                             `pulumi:"type"`
}

func LookupAliasOutput(ctx *pulumi.Context, args LookupAliasOutputArgs, opts ...pulumi.InvokeOption) LookupAliasResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAliasResult, error) {
			args := v.(LookupAliasArgs)
			r, err := LookupAlias(ctx, &args, opts...)
			var s LookupAliasResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAliasResultOutput)
}

type LookupAliasOutputArgs struct {
	AliasName pulumi.StringInput `pulumi:"aliasName"`
}

func (LookupAliasOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAliasArgs)(nil)).Elem()
}


type LookupAliasResultOutput struct{ *pulumi.OutputState }

func (LookupAliasResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAliasResult)(nil)).Elem()
}

func (o LookupAliasResultOutput) ToLookupAliasResultOutput() LookupAliasResultOutput {
	return o
}

func (o LookupAliasResultOutput) ToLookupAliasResultOutputWithContext(ctx context.Context) LookupAliasResultOutput {
	return o
}

func (o LookupAliasResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAliasResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAliasResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAliasResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAliasResultOutput) Properties() PutAliasResponsePropertiesResponseOutput {
	return o.ApplyT(func(v LookupAliasResult) PutAliasResponsePropertiesResponse { return v.Properties }).(PutAliasResponsePropertiesResponseOutput)
}

func (o LookupAliasResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAliasResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAliasResultOutput{})
}

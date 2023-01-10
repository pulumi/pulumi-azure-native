


package v20191001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSubscriptionAlias(ctx *pulumi.Context, args *LookupSubscriptionAliasArgs, opts ...pulumi.InvokeOption) (*LookupSubscriptionAliasResult, error) {
	var rv LookupSubscriptionAliasResult
	err := ctx.Invoke("azure-native:subscription/v20191001preview:getSubscriptionAlias", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubscriptionAliasArgs struct {
	AliasName string `pulumi:"aliasName"`
}


type LookupSubscriptionAliasResult struct {
	Id         string                             `pulumi:"id"`
	Name       string                             `pulumi:"name"`
	Properties PutAliasResponsePropertiesResponse `pulumi:"properties"`
	Type       string                             `pulumi:"type"`
}

func LookupSubscriptionAliasOutput(ctx *pulumi.Context, args LookupSubscriptionAliasOutputArgs, opts ...pulumi.InvokeOption) LookupSubscriptionAliasResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubscriptionAliasResult, error) {
			args := v.(LookupSubscriptionAliasArgs)
			r, err := LookupSubscriptionAlias(ctx, &args, opts...)
			var s LookupSubscriptionAliasResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubscriptionAliasResultOutput)
}

type LookupSubscriptionAliasOutputArgs struct {
	AliasName pulumi.StringInput `pulumi:"aliasName"`
}

func (LookupSubscriptionAliasOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionAliasArgs)(nil)).Elem()
}


type LookupSubscriptionAliasResultOutput struct{ *pulumi.OutputState }

func (LookupSubscriptionAliasResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionAliasResult)(nil)).Elem()
}

func (o LookupSubscriptionAliasResultOutput) ToLookupSubscriptionAliasResultOutput() LookupSubscriptionAliasResultOutput {
	return o
}

func (o LookupSubscriptionAliasResultOutput) ToLookupSubscriptionAliasResultOutputWithContext(ctx context.Context) LookupSubscriptionAliasResultOutput {
	return o
}

func (o LookupSubscriptionAliasResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionAliasResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSubscriptionAliasResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionAliasResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSubscriptionAliasResultOutput) Properties() PutAliasResponsePropertiesResponseOutput {
	return o.ApplyT(func(v LookupSubscriptionAliasResult) PutAliasResponsePropertiesResponse { return v.Properties }).(PutAliasResponsePropertiesResponseOutput)
}

func (o LookupSubscriptionAliasResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionAliasResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubscriptionAliasResultOutput{})
}

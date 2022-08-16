


package v20201120

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDefaultRollout(ctx *pulumi.Context, args *LookupDefaultRolloutArgs, opts ...pulumi.InvokeOption) (*LookupDefaultRolloutResult, error) {
	var rv LookupDefaultRolloutResult
	err := ctx.Invoke("azure-native:providerhub/v20201120:getDefaultRollout", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDefaultRolloutArgs struct {
	ProviderNamespace string `pulumi:"providerNamespace"`
	RolloutName       string `pulumi:"rolloutName"`
}


type LookupDefaultRolloutResult struct {
	Id         string                           `pulumi:"id"`
	Name       string                           `pulumi:"name"`
	Properties DefaultRolloutResponseProperties `pulumi:"properties"`
	Type       string                           `pulumi:"type"`
}

func LookupDefaultRolloutOutput(ctx *pulumi.Context, args LookupDefaultRolloutOutputArgs, opts ...pulumi.InvokeOption) LookupDefaultRolloutResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDefaultRolloutResult, error) {
			args := v.(LookupDefaultRolloutArgs)
			r, err := LookupDefaultRollout(ctx, &args, opts...)
			var s LookupDefaultRolloutResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDefaultRolloutResultOutput)
}

type LookupDefaultRolloutOutputArgs struct {
	ProviderNamespace pulumi.StringInput `pulumi:"providerNamespace"`
	RolloutName       pulumi.StringInput `pulumi:"rolloutName"`
}

func (LookupDefaultRolloutOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultRolloutArgs)(nil)).Elem()
}


type LookupDefaultRolloutResultOutput struct{ *pulumi.OutputState }

func (LookupDefaultRolloutResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultRolloutResult)(nil)).Elem()
}

func (o LookupDefaultRolloutResultOutput) ToLookupDefaultRolloutResultOutput() LookupDefaultRolloutResultOutput {
	return o
}

func (o LookupDefaultRolloutResultOutput) ToLookupDefaultRolloutResultOutputWithContext(ctx context.Context) LookupDefaultRolloutResultOutput {
	return o
}

func (o LookupDefaultRolloutResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultRolloutResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDefaultRolloutResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultRolloutResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDefaultRolloutResultOutput) Properties() DefaultRolloutResponsePropertiesOutput {
	return o.ApplyT(func(v LookupDefaultRolloutResult) DefaultRolloutResponseProperties { return v.Properties }).(DefaultRolloutResponsePropertiesOutput)
}

func (o LookupDefaultRolloutResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultRolloutResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDefaultRolloutResultOutput{})
}

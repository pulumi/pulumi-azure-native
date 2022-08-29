


package v20201120

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProviderRegistration(ctx *pulumi.Context, args *LookupProviderRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupProviderRegistrationResult, error) {
	var rv LookupProviderRegistrationResult
	err := ctx.Invoke("azure-native:providerhub/v20201120:getProviderRegistration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProviderRegistrationArgs struct {
	ProviderNamespace string `pulumi:"providerNamespace"`
}

type LookupProviderRegistrationResult struct {
	Id         string                                 `pulumi:"id"`
	Name       string                                 `pulumi:"name"`
	Properties ProviderRegistrationResponseProperties `pulumi:"properties"`
	Type       string                                 `pulumi:"type"`
}

func LookupProviderRegistrationOutput(ctx *pulumi.Context, args LookupProviderRegistrationOutputArgs, opts ...pulumi.InvokeOption) LookupProviderRegistrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProviderRegistrationResult, error) {
			args := v.(LookupProviderRegistrationArgs)
			r, err := LookupProviderRegistration(ctx, &args, opts...)
			var s LookupProviderRegistrationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProviderRegistrationResultOutput)
}

type LookupProviderRegistrationOutputArgs struct {
	ProviderNamespace pulumi.StringInput `pulumi:"providerNamespace"`
}

func (LookupProviderRegistrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProviderRegistrationArgs)(nil)).Elem()
}

type LookupProviderRegistrationResultOutput struct{ *pulumi.OutputState }

func (LookupProviderRegistrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProviderRegistrationResult)(nil)).Elem()
}

func (o LookupProviderRegistrationResultOutput) ToLookupProviderRegistrationResultOutput() LookupProviderRegistrationResultOutput {
	return o
}

func (o LookupProviderRegistrationResultOutput) ToLookupProviderRegistrationResultOutputWithContext(ctx context.Context) LookupProviderRegistrationResultOutput {
	return o
}

func (o LookupProviderRegistrationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderRegistrationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProviderRegistrationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderRegistrationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProviderRegistrationResultOutput) Properties() ProviderRegistrationResponsePropertiesOutput {
	return o.ApplyT(func(v LookupProviderRegistrationResult) ProviderRegistrationResponseProperties { return v.Properties }).(ProviderRegistrationResponsePropertiesOutput)
}

func (o LookupProviderRegistrationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderRegistrationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProviderRegistrationResultOutput{})
}

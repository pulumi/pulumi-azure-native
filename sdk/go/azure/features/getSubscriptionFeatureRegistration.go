


package features

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSubscriptionFeatureRegistration(ctx *pulumi.Context, args *LookupSubscriptionFeatureRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupSubscriptionFeatureRegistrationResult, error) {
	var rv LookupSubscriptionFeatureRegistrationResult
	err := ctx.Invoke("azure-native:features:getSubscriptionFeatureRegistration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSubscriptionFeatureRegistrationArgs struct {
	FeatureName       string `pulumi:"featureName"`
	ProviderNamespace string `pulumi:"providerNamespace"`
}


type LookupSubscriptionFeatureRegistrationResult struct {
	Id         string                                            `pulumi:"id"`
	Name       string                                            `pulumi:"name"`
	Properties SubscriptionFeatureRegistrationResponseProperties `pulumi:"properties"`
	Type       string                                            `pulumi:"type"`
}


func (val *LookupSubscriptionFeatureRegistrationResult) Defaults() *LookupSubscriptionFeatureRegistrationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupSubscriptionFeatureRegistrationOutput(ctx *pulumi.Context, args LookupSubscriptionFeatureRegistrationOutputArgs, opts ...pulumi.InvokeOption) LookupSubscriptionFeatureRegistrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubscriptionFeatureRegistrationResult, error) {
			args := v.(LookupSubscriptionFeatureRegistrationArgs)
			r, err := LookupSubscriptionFeatureRegistration(ctx, &args, opts...)
			var s LookupSubscriptionFeatureRegistrationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubscriptionFeatureRegistrationResultOutput)
}

type LookupSubscriptionFeatureRegistrationOutputArgs struct {
	FeatureName       pulumi.StringInput `pulumi:"featureName"`
	ProviderNamespace pulumi.StringInput `pulumi:"providerNamespace"`
}

func (LookupSubscriptionFeatureRegistrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionFeatureRegistrationArgs)(nil)).Elem()
}


type LookupSubscriptionFeatureRegistrationResultOutput struct{ *pulumi.OutputState }

func (LookupSubscriptionFeatureRegistrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionFeatureRegistrationResult)(nil)).Elem()
}

func (o LookupSubscriptionFeatureRegistrationResultOutput) ToLookupSubscriptionFeatureRegistrationResultOutput() LookupSubscriptionFeatureRegistrationResultOutput {
	return o
}

func (o LookupSubscriptionFeatureRegistrationResultOutput) ToLookupSubscriptionFeatureRegistrationResultOutputWithContext(ctx context.Context) LookupSubscriptionFeatureRegistrationResultOutput {
	return o
}

func (o LookupSubscriptionFeatureRegistrationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionFeatureRegistrationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSubscriptionFeatureRegistrationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionFeatureRegistrationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSubscriptionFeatureRegistrationResultOutput) Properties() SubscriptionFeatureRegistrationResponsePropertiesOutput {
	return o.ApplyT(func(v LookupSubscriptionFeatureRegistrationResult) SubscriptionFeatureRegistrationResponseProperties {
		return v.Properties
	}).(SubscriptionFeatureRegistrationResponsePropertiesOutput)
}

func (o LookupSubscriptionFeatureRegistrationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionFeatureRegistrationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubscriptionFeatureRegistrationResultOutput{})
}

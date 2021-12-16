


package features

import (
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

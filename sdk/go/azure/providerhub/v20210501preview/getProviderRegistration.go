


package v20210501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProviderRegistration(ctx *pulumi.Context, args *LookupProviderRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupProviderRegistrationResult, error) {
	var rv LookupProviderRegistrationResult
	err := ctx.Invoke("azure-native:providerhub/v20210501preview:getProviderRegistration", args, &rv, opts...)
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

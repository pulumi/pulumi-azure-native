


package v20210501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupResourceTypeRegistration(ctx *pulumi.Context, args *LookupResourceTypeRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupResourceTypeRegistrationResult, error) {
	var rv LookupResourceTypeRegistrationResult
	err := ctx.Invoke("azure-native:providerhub/v20210501preview:getResourceTypeRegistration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceTypeRegistrationArgs struct {
	ProviderNamespace string `pulumi:"providerNamespace"`
	ResourceType      string `pulumi:"resourceType"`
}

type LookupResourceTypeRegistrationResult struct {
	Id         string                                     `pulumi:"id"`
	Name       string                                     `pulumi:"name"`
	Properties ResourceTypeRegistrationResponseProperties `pulumi:"properties"`
	Type       string                                     `pulumi:"type"`
}

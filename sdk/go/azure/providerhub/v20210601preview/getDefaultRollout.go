


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDefaultRollout(ctx *pulumi.Context, args *LookupDefaultRolloutArgs, opts ...pulumi.InvokeOption) (*LookupDefaultRolloutResult, error) {
	var rv LookupDefaultRolloutResult
	err := ctx.Invoke("azure-native:providerhub/v20210601preview:getDefaultRollout", args, &rv, opts...)
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
	SystemData SystemDataResponse               `pulumi:"systemData"`
	Type       string                           `pulumi:"type"`
}

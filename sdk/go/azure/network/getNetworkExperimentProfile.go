


package network

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkExperimentProfile(ctx *pulumi.Context, args *LookupNetworkExperimentProfileArgs, opts ...pulumi.InvokeOption) (*LookupNetworkExperimentProfileResult, error) {
	var rv LookupNetworkExperimentProfileResult
	err := ctx.Invoke("azure-native:network:getNetworkExperimentProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkExperimentProfileArgs struct {
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupNetworkExperimentProfileResult struct {
	EnabledState  *string           `pulumi:"enabledState"`
	Etag          *string           `pulumi:"etag"`
	Id            string            `pulumi:"id"`
	Location      *string           `pulumi:"location"`
	Name          string            `pulumi:"name"`
	ResourceState string            `pulumi:"resourceState"`
	Tags          map[string]string `pulumi:"tags"`
	Type          string            `pulumi:"type"`
}

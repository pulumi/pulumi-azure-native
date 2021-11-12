


package v20211201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListConfigurations(ctx *pulumi.Context, args *ListConfigurationsArgs, opts ...pulumi.InvokeOption) (*ListConfigurationsResult, error) {
	var rv ListConfigurationsResult
	err := ctx.Invoke("azure-native:edgeorder/v20211201:listConfigurations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConfigurationsArgs struct {
	ConfigurationFilters        []ConfigurationFilters       `pulumi:"configurationFilters"`
	CustomerSubscriptionDetails *CustomerSubscriptionDetails `pulumi:"customerSubscriptionDetails"`
	SkipToken                   *string                      `pulumi:"skipToken"`
}


type ListConfigurationsResult struct {
	NextLink *string                 `pulumi:"nextLink"`
	Value    []ConfigurationResponse `pulumi:"value"`
}




package portal

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListListTenantConfigurationViolation(ctx *pulumi.Context, args *ListListTenantConfigurationViolationArgs, opts ...pulumi.InvokeOption) (*ListListTenantConfigurationViolationResult, error) {
	var rv ListListTenantConfigurationViolationResult
	err := ctx.Invoke("azure-native:portal:listListTenantConfigurationViolation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListListTenantConfigurationViolationArgs struct {
}


type ListListTenantConfigurationViolationResult struct {
	NextLink *string             `pulumi:"nextLink"`
	Value    []ViolationResponse `pulumi:"value"`
}

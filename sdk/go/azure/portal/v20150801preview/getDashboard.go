


package v20150801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDashboard(ctx *pulumi.Context, args *LookupDashboardArgs, opts ...pulumi.InvokeOption) (*LookupDashboardResult, error) {
	var rv LookupDashboardResult
	err := ctx.Invoke("azure-native:portal/v20150801preview:getDashboard", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDashboardArgs struct {
	DashboardName     string `pulumi:"dashboardName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDashboardResult struct {
	Id       string                           `pulumi:"id"`
	Lenses   map[string]DashboardLensResponse `pulumi:"lenses"`
	Location string                           `pulumi:"location"`
	Metadata map[string]interface{}           `pulumi:"metadata"`
	Name     string                           `pulumi:"name"`
	Tags     map[string]string                `pulumi:"tags"`
	Type     string                           `pulumi:"type"`
}

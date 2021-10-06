


package v20200901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDashboard(ctx *pulumi.Context, args *LookupDashboardArgs, opts ...pulumi.InvokeOption) (*LookupDashboardResult, error) {
	var rv LookupDashboardResult
	err := ctx.Invoke("azure-native:portal/v20200901preview:getDashboard", args, &rv, opts...)
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
	Id       string                  `pulumi:"id"`
	Lenses   []DashboardLensResponse `pulumi:"lenses"`
	Location string                  `pulumi:"location"`
	Metadata map[string]interface{}  `pulumi:"metadata"`
	Name     string                  `pulumi:"name"`
	Tags     map[string]string       `pulumi:"tags"`
	Type     string                  `pulumi:"type"`
}




package v20170101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupView(ctx *pulumi.Context, args *LookupViewArgs, opts ...pulumi.InvokeOption) (*LookupViewResult, error) {
	var rv LookupViewResult
	err := ctx.Invoke("azure-native:customerinsights/v20170101:getView", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupViewArgs struct {
	HubName           string `pulumi:"hubName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	UserId            string `pulumi:"userId"`
	ViewName          string `pulumi:"viewName"`
}


type LookupViewResult struct {
	Changed     string            `pulumi:"changed"`
	Created     string            `pulumi:"created"`
	Definition  string            `pulumi:"definition"`
	DisplayName map[string]string `pulumi:"displayName"`
	Id          string            `pulumi:"id"`
	Name        string            `pulumi:"name"`
	TenantId    string            `pulumi:"tenantId"`
	Type        string            `pulumi:"type"`
	UserId      *string           `pulumi:"userId"`
	ViewName    string            `pulumi:"viewName"`
}




package v20170301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedInstanceAdministrator(ctx *pulumi.Context, args *LookupManagedInstanceAdministratorArgs, opts ...pulumi.InvokeOption) (*LookupManagedInstanceAdministratorResult, error) {
	var rv LookupManagedInstanceAdministratorResult
	err := ctx.Invoke("azure-native:sql/v20170301preview:getManagedInstanceAdministrator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedInstanceAdministratorArgs struct {
	AdministratorName   string `pulumi:"administratorName"`
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupManagedInstanceAdministratorResult struct {
	AdministratorType string  `pulumi:"administratorType"`
	Id                string  `pulumi:"id"`
	Login             string  `pulumi:"login"`
	Name              string  `pulumi:"name"`
	Sid               string  `pulumi:"sid"`
	TenantId          *string `pulumi:"tenantId"`
	Type              string  `pulumi:"type"`
}

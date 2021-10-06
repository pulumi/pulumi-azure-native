


package v20201001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMonitorUserRoles(ctx *pulumi.Context, args *ListMonitorUserRolesArgs, opts ...pulumi.InvokeOption) (*ListMonitorUserRolesResult, error) {
	var rv ListMonitorUserRolesResult
	err := ctx.Invoke("azure-native:logz/v20201001:listMonitorUserRoles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMonitorUserRolesArgs struct {
	EmailAddress      *string `pulumi:"emailAddress"`
	MonitorName       string  `pulumi:"monitorName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type ListMonitorUserRolesResult struct {
	NextLink *string                    `pulumi:"nextLink"`
	Value    []UserRoleResponseResponse `pulumi:"value"`
}

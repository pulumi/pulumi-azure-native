


package v20201001

import (
	"context"
	"reflect"

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

func ListMonitorUserRolesOutput(ctx *pulumi.Context, args ListMonitorUserRolesOutputArgs, opts ...pulumi.InvokeOption) ListMonitorUserRolesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListMonitorUserRolesResult, error) {
			args := v.(ListMonitorUserRolesArgs)
			r, err := ListMonitorUserRoles(ctx, &args, opts...)
			var s ListMonitorUserRolesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListMonitorUserRolesResultOutput)
}

type ListMonitorUserRolesOutputArgs struct {
	EmailAddress      pulumi.StringPtrInput `pulumi:"emailAddress"`
	MonitorName       pulumi.StringInput    `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (ListMonitorUserRolesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitorUserRolesArgs)(nil)).Elem()
}


type ListMonitorUserRolesResultOutput struct{ *pulumi.OutputState }

func (ListMonitorUserRolesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitorUserRolesResult)(nil)).Elem()
}

func (o ListMonitorUserRolesResultOutput) ToListMonitorUserRolesResultOutput() ListMonitorUserRolesResultOutput {
	return o
}

func (o ListMonitorUserRolesResultOutput) ToListMonitorUserRolesResultOutputWithContext(ctx context.Context) ListMonitorUserRolesResultOutput {
	return o
}

func (o ListMonitorUserRolesResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListMonitorUserRolesResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListMonitorUserRolesResultOutput) Value() UserRoleResponseResponseArrayOutput {
	return o.ApplyT(func(v ListMonitorUserRolesResult) []UserRoleResponseResponse { return v.Value }).(UserRoleResponseResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListMonitorUserRolesResultOutput{})
}




package v20230101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCloudEdgeManagementRole(ctx *pulumi.Context, args *LookupCloudEdgeManagementRoleArgs, opts ...pulumi.InvokeOption) (*LookupCloudEdgeManagementRoleResult, error) {
	var rv LookupCloudEdgeManagementRoleResult
	err := ctx.Invoke("azure-native:databoxedge/v20230101preview:getCloudEdgeManagementRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCloudEdgeManagementRoleArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}




type LookupCloudEdgeManagementRoleResult struct {
	EdgeProfile           EdgeProfileResponse `pulumi:"edgeProfile"`
	Id                    string              `pulumi:"id"`
	Kind                  string              `pulumi:"kind"`
	LocalManagementStatus string              `pulumi:"localManagementStatus"`
	Name                  string              `pulumi:"name"`
	RoleStatus            string              `pulumi:"roleStatus"`
	SystemData            SystemDataResponse  `pulumi:"systemData"`
	Type                  string              `pulumi:"type"`
}

func LookupCloudEdgeManagementRoleOutput(ctx *pulumi.Context, args LookupCloudEdgeManagementRoleOutputArgs, opts ...pulumi.InvokeOption) LookupCloudEdgeManagementRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudEdgeManagementRoleResult, error) {
			args := v.(LookupCloudEdgeManagementRoleArgs)
			r, err := LookupCloudEdgeManagementRole(ctx, &args, opts...)
			var s LookupCloudEdgeManagementRoleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudEdgeManagementRoleResultOutput)
}

type LookupCloudEdgeManagementRoleOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCloudEdgeManagementRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudEdgeManagementRoleArgs)(nil)).Elem()
}




type LookupCloudEdgeManagementRoleResultOutput struct{ *pulumi.OutputState }

func (LookupCloudEdgeManagementRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudEdgeManagementRoleResult)(nil)).Elem()
}

func (o LookupCloudEdgeManagementRoleResultOutput) ToLookupCloudEdgeManagementRoleResultOutput() LookupCloudEdgeManagementRoleResultOutput {
	return o
}

func (o LookupCloudEdgeManagementRoleResultOutput) ToLookupCloudEdgeManagementRoleResultOutputWithContext(ctx context.Context) LookupCloudEdgeManagementRoleResultOutput {
	return o
}

func (o LookupCloudEdgeManagementRoleResultOutput) EdgeProfile() EdgeProfileResponseOutput {
	return o.ApplyT(func(v LookupCloudEdgeManagementRoleResult) EdgeProfileResponse { return v.EdgeProfile }).(EdgeProfileResponseOutput)
}

func (o LookupCloudEdgeManagementRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudEdgeManagementRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCloudEdgeManagementRoleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudEdgeManagementRoleResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupCloudEdgeManagementRoleResultOutput) LocalManagementStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudEdgeManagementRoleResult) string { return v.LocalManagementStatus }).(pulumi.StringOutput)
}

func (o LookupCloudEdgeManagementRoleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudEdgeManagementRoleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCloudEdgeManagementRoleResultOutput) RoleStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudEdgeManagementRoleResult) string { return v.RoleStatus }).(pulumi.StringOutput)
}

func (o LookupCloudEdgeManagementRoleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCloudEdgeManagementRoleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCloudEdgeManagementRoleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudEdgeManagementRoleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudEdgeManagementRoleResultOutput{})
}

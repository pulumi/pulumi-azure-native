


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProjectEnvironmentType(ctx *pulumi.Context, args *LookupProjectEnvironmentTypeArgs, opts ...pulumi.InvokeOption) (*LookupProjectEnvironmentTypeResult, error) {
	var rv LookupProjectEnvironmentTypeResult
	err := ctx.Invoke("azure-native:devcenter/v20220901preview:getProjectEnvironmentType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectEnvironmentTypeArgs struct {
	EnvironmentTypeName string `pulumi:"environmentTypeName"`
	ProjectName         string `pulumi:"projectName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupProjectEnvironmentTypeResult struct {
	CreatorRoleAssignment *ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignment `pulumi:"creatorRoleAssignment"`
	DeploymentTargetId    *string                                                              `pulumi:"deploymentTargetId"`
	Id                    string                                                               `pulumi:"id"`
	Identity              *ManagedServiceIdentityResponse                                      `pulumi:"identity"`
	Location              *string                                                              `pulumi:"location"`
	Name                  string                                                               `pulumi:"name"`
	ProvisioningState     string                                                               `pulumi:"provisioningState"`
	Status                *string                                                              `pulumi:"status"`
	SystemData            SystemDataResponse                                                   `pulumi:"systemData"`
	Tags                  map[string]string                                                    `pulumi:"tags"`
	Type                  string                                                               `pulumi:"type"`
	UserRoleAssignments   map[string]UserRoleAssignmentResponse                                `pulumi:"userRoleAssignments"`
}

func LookupProjectEnvironmentTypeOutput(ctx *pulumi.Context, args LookupProjectEnvironmentTypeOutputArgs, opts ...pulumi.InvokeOption) LookupProjectEnvironmentTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectEnvironmentTypeResult, error) {
			args := v.(LookupProjectEnvironmentTypeArgs)
			r, err := LookupProjectEnvironmentType(ctx, &args, opts...)
			var s LookupProjectEnvironmentTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectEnvironmentTypeResultOutput)
}

type LookupProjectEnvironmentTypeOutputArgs struct {
	EnvironmentTypeName pulumi.StringInput `pulumi:"environmentTypeName"`
	ProjectName         pulumi.StringInput `pulumi:"projectName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupProjectEnvironmentTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectEnvironmentTypeArgs)(nil)).Elem()
}


type LookupProjectEnvironmentTypeResultOutput struct{ *pulumi.OutputState }

func (LookupProjectEnvironmentTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectEnvironmentTypeResult)(nil)).Elem()
}

func (o LookupProjectEnvironmentTypeResultOutput) ToLookupProjectEnvironmentTypeResultOutput() LookupProjectEnvironmentTypeResultOutput {
	return o
}

func (o LookupProjectEnvironmentTypeResultOutput) ToLookupProjectEnvironmentTypeResultOutputWithContext(ctx context.Context) LookupProjectEnvironmentTypeResultOutput {
	return o
}

func (o LookupProjectEnvironmentTypeResultOutput) CreatorRoleAssignment() ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) *ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignment {
		return v.CreatorRoleAssignment
	}).(ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput)
}

func (o LookupProjectEnvironmentTypeResultOutput) DeploymentTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) *string { return v.DeploymentTargetId }).(pulumi.StringPtrOutput)
}

func (o LookupProjectEnvironmentTypeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProjectEnvironmentTypeResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupProjectEnvironmentTypeResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupProjectEnvironmentTypeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProjectEnvironmentTypeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupProjectEnvironmentTypeResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupProjectEnvironmentTypeResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupProjectEnvironmentTypeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupProjectEnvironmentTypeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupProjectEnvironmentTypeResultOutput) UserRoleAssignments() UserRoleAssignmentResponseMapOutput {
	return o.ApplyT(func(v LookupProjectEnvironmentTypeResult) map[string]UserRoleAssignmentResponse {
		return v.UserRoleAssignments
	}).(UserRoleAssignmentResponseMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectEnvironmentTypeResultOutput{})
}

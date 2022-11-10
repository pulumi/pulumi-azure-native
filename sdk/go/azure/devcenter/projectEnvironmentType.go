


package devcenter

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProjectEnvironmentType struct {
	pulumi.CustomResourceState

	CreatorRoleAssignment ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput `pulumi:"creatorRoleAssignment"`
	DeploymentTargetId    pulumi.StringPtrOutput                                                       `pulumi:"deploymentTargetId"`
	Identity              ManagedServiceIdentityResponsePtrOutput                                      `pulumi:"identity"`
	Location              pulumi.StringPtrOutput                                                       `pulumi:"location"`
	Name                  pulumi.StringOutput                                                          `pulumi:"name"`
	ProvisioningState     pulumi.StringOutput                                                          `pulumi:"provisioningState"`
	Status                pulumi.StringPtrOutput                                                       `pulumi:"status"`
	SystemData            SystemDataResponseOutput                                                     `pulumi:"systemData"`
	Tags                  pulumi.StringMapOutput                                                       `pulumi:"tags"`
	Type                  pulumi.StringOutput                                                          `pulumi:"type"`
	UserRoleAssignments   UserRoleAssignmentResponseMapOutput                                          `pulumi:"userRoleAssignments"`
}


func NewProjectEnvironmentType(ctx *pulumi.Context,
	name string, args *ProjectEnvironmentTypeArgs, opts ...pulumi.ResourceOption) (*ProjectEnvironmentType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter/v20220801preview:ProjectEnvironmentType"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220901preview:ProjectEnvironmentType"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221012preview:ProjectEnvironmentType"),
		},
	})
	opts = append(opts, aliases)
	var resource ProjectEnvironmentType
	err := ctx.RegisterResource("azure-native:devcenter:ProjectEnvironmentType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProjectEnvironmentType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectEnvironmentTypeState, opts ...pulumi.ResourceOption) (*ProjectEnvironmentType, error) {
	var resource ProjectEnvironmentType
	err := ctx.ReadResource("azure-native:devcenter:ProjectEnvironmentType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type projectEnvironmentTypeState struct {
}

type ProjectEnvironmentTypeState struct {
}

func (ProjectEnvironmentTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectEnvironmentTypeState)(nil)).Elem()
}

type projectEnvironmentTypeArgs struct {
	CreatorRoleAssignment *ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignment `pulumi:"creatorRoleAssignment"`
	DeploymentTargetId    *string                                                      `pulumi:"deploymentTargetId"`
	EnvironmentTypeName   *string                                                      `pulumi:"environmentTypeName"`
	Identity              *ManagedServiceIdentity                                      `pulumi:"identity"`
	Location              *string                                                      `pulumi:"location"`
	ProjectName           string                                                       `pulumi:"projectName"`
	ResourceGroupName     string                                                       `pulumi:"resourceGroupName"`
	Status                *string                                                      `pulumi:"status"`
	Tags                  map[string]string                                            `pulumi:"tags"`
	UserRoleAssignments   map[string]UserRoleAssignment                                `pulumi:"userRoleAssignments"`
}


type ProjectEnvironmentTypeArgs struct {
	CreatorRoleAssignment ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrInput
	DeploymentTargetId    pulumi.StringPtrInput
	EnvironmentTypeName   pulumi.StringPtrInput
	Identity              ManagedServiceIdentityPtrInput
	Location              pulumi.StringPtrInput
	ProjectName           pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	Status                pulumi.StringPtrInput
	Tags                  pulumi.StringMapInput
	UserRoleAssignments   UserRoleAssignmentMapInput
}

func (ProjectEnvironmentTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectEnvironmentTypeArgs)(nil)).Elem()
}

type ProjectEnvironmentTypeInput interface {
	pulumi.Input

	ToProjectEnvironmentTypeOutput() ProjectEnvironmentTypeOutput
	ToProjectEnvironmentTypeOutputWithContext(ctx context.Context) ProjectEnvironmentTypeOutput
}

func (*ProjectEnvironmentType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectEnvironmentType)(nil)).Elem()
}

func (i *ProjectEnvironmentType) ToProjectEnvironmentTypeOutput() ProjectEnvironmentTypeOutput {
	return i.ToProjectEnvironmentTypeOutputWithContext(context.Background())
}

func (i *ProjectEnvironmentType) ToProjectEnvironmentTypeOutputWithContext(ctx context.Context) ProjectEnvironmentTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectEnvironmentTypeOutput)
}

type ProjectEnvironmentTypeOutput struct{ *pulumi.OutputState }

func (ProjectEnvironmentTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectEnvironmentType)(nil)).Elem()
}

func (o ProjectEnvironmentTypeOutput) ToProjectEnvironmentTypeOutput() ProjectEnvironmentTypeOutput {
	return o
}

func (o ProjectEnvironmentTypeOutput) ToProjectEnvironmentTypeOutputWithContext(ctx context.Context) ProjectEnvironmentTypeOutput {
	return o
}

func (o ProjectEnvironmentTypeOutput) CreatorRoleAssignment() ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput {
		return v.CreatorRoleAssignment
	}).(ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput)
}

func (o ProjectEnvironmentTypeOutput) DeploymentTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) pulumi.StringPtrOutput { return v.DeploymentTargetId }).(pulumi.StringPtrOutput)
}

func (o ProjectEnvironmentTypeOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o ProjectEnvironmentTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ProjectEnvironmentTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProjectEnvironmentTypeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ProjectEnvironmentTypeOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ProjectEnvironmentTypeOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ProjectEnvironmentTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ProjectEnvironmentTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ProjectEnvironmentTypeOutput) UserRoleAssignments() UserRoleAssignmentResponseMapOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) UserRoleAssignmentResponseMapOutput { return v.UserRoleAssignments }).(UserRoleAssignmentResponseMapOutput)
}

func init() {
	pulumi.RegisterOutputType(ProjectEnvironmentTypeOutput{})
}

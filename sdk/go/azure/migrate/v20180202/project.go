


package v20180202

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Project struct {
	pulumi.CustomResourceState

	CreatedTimestamp          pulumi.StringOutput    `pulumi:"createdTimestamp"`
	CustomerWorkspaceId       pulumi.StringPtrOutput `pulumi:"customerWorkspaceId"`
	CustomerWorkspaceLocation pulumi.StringPtrOutput `pulumi:"customerWorkspaceLocation"`
	DiscoveryStatus           pulumi.StringOutput    `pulumi:"discoveryStatus"`
	ETag                      pulumi.StringPtrOutput `pulumi:"eTag"`
	LastAssessmentTimestamp   pulumi.StringOutput    `pulumi:"lastAssessmentTimestamp"`
	LastDiscoverySessionId    pulumi.StringOutput    `pulumi:"lastDiscoverySessionId"`
	LastDiscoveryTimestamp    pulumi.StringOutput    `pulumi:"lastDiscoveryTimestamp"`
	Location                  pulumi.StringPtrOutput `pulumi:"location"`
	Name                      pulumi.StringOutput    `pulumi:"name"`
	NumberOfAssessments       pulumi.IntOutput       `pulumi:"numberOfAssessments"`
	NumberOfGroups            pulumi.IntOutput       `pulumi:"numberOfGroups"`
	NumberOfMachines          pulumi.IntOutput       `pulumi:"numberOfMachines"`
	ProvisioningState         pulumi.StringPtrOutput `pulumi:"provisioningState"`
	Tags                      pulumi.AnyOutput       `pulumi:"tags"`
	Type                      pulumi.StringOutput    `pulumi:"type"`
	UpdatedTimestamp          pulumi.StringOutput    `pulumi:"updatedTimestamp"`
}


func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate/v20171111preview:Project"),
		},
	})
	opts = append(opts, aliases)
	var resource Project
	err := ctx.RegisterResource("azure-native:migrate/v20180202:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("azure-native:migrate/v20180202:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type projectState struct {
}

type ProjectState struct {
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	CustomerWorkspaceId       *string     `pulumi:"customerWorkspaceId"`
	CustomerWorkspaceLocation *string     `pulumi:"customerWorkspaceLocation"`
	ETag                      *string     `pulumi:"eTag"`
	Location                  *string     `pulumi:"location"`
	ProjectName               *string     `pulumi:"projectName"`
	ProvisioningState         *string     `pulumi:"provisioningState"`
	ResourceGroupName         string      `pulumi:"resourceGroupName"`
	Tags                      interface{} `pulumi:"tags"`
}


type ProjectArgs struct {
	CustomerWorkspaceId       pulumi.StringPtrInput
	CustomerWorkspaceLocation pulumi.StringPtrInput
	ETag                      pulumi.StringPtrInput
	Location                  pulumi.StringPtrInput
	ProjectName               pulumi.StringPtrInput
	ProvisioningState         pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	Tags                      pulumi.Input
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProjectOutput{})
}

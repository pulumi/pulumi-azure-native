


package v20171111preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Project struct {
	pulumi.CustomResourceState

	CreatedTimestamp    pulumi.StringOutput    `pulumi:"createdTimestamp"`
	CustomerWorkspaceId pulumi.StringPtrOutput `pulumi:"customerWorkspaceId"`
	DiscoveryStatus     pulumi.StringOutput    `pulumi:"discoveryStatus"`
	ETag                pulumi.StringPtrOutput `pulumi:"eTag"`
	Location            pulumi.StringPtrOutput `pulumi:"location"`
	Name                pulumi.StringOutput    `pulumi:"name"`
	NumberOfAssessments pulumi.IntOutput       `pulumi:"numberOfAssessments"`
	NumberOfGroups      pulumi.IntOutput       `pulumi:"numberOfGroups"`
	NumberOfMachines    pulumi.IntOutput       `pulumi:"numberOfMachines"`
	ProvisioningState   pulumi.StringPtrOutput `pulumi:"provisioningState"`
	Tags                pulumi.AnyOutput       `pulumi:"tags"`
	Type                pulumi.StringOutput    `pulumi:"type"`
	UpdatedTimestamp    pulumi.StringOutput    `pulumi:"updatedTimestamp"`
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
			Type: pulumi.String("azure-native:migrate/v20180202:Project"),
		},
	})
	opts = append(opts, aliases)
	var resource Project
	err := ctx.RegisterResource("azure-native:migrate/v20171111preview:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("azure-native:migrate/v20171111preview:Project", name, id, state, &resource, opts...)
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
	CustomerWorkspaceId *string     `pulumi:"customerWorkspaceId"`
	ETag                *string     `pulumi:"eTag"`
	Location            *string     `pulumi:"location"`
	ProjectName         *string     `pulumi:"projectName"`
	ProvisioningState   *string     `pulumi:"provisioningState"`
	ResourceGroupName   string      `pulumi:"resourceGroupName"`
	Tags                interface{} `pulumi:"tags"`
}


type ProjectArgs struct {
	CustomerWorkspaceId pulumi.StringPtrInput
	ETag                pulumi.StringPtrInput
	Location            pulumi.StringPtrInput
	ProjectName         pulumi.StringPtrInput
	ProvisioningState   pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	Tags                pulumi.Input
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

func (o ProjectOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

func (o ProjectOutput) CustomerWorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.CustomerWorkspaceId }).(pulumi.StringPtrOutput)
}

func (o ProjectOutput) DiscoveryStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.DiscoveryStatus }).(pulumi.StringOutput)
}

func (o ProjectOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o ProjectOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProjectOutput) NumberOfAssessments() pulumi.IntOutput {
	return o.ApplyT(func(v *Project) pulumi.IntOutput { return v.NumberOfAssessments }).(pulumi.IntOutput)
}

func (o ProjectOutput) NumberOfGroups() pulumi.IntOutput {
	return o.ApplyT(func(v *Project) pulumi.IntOutput { return v.NumberOfGroups }).(pulumi.IntOutput)
}

func (o ProjectOutput) NumberOfMachines() pulumi.IntOutput {
	return o.ApplyT(func(v *Project) pulumi.IntOutput { return v.NumberOfMachines }).(pulumi.IntOutput)
}

func (o ProjectOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ProjectOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *Project) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func (o ProjectOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ProjectOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProjectOutput{})
}

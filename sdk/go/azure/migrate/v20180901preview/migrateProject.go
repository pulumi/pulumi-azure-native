


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MigrateProject struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput                 `pulumi:"eTag"`
	Location   pulumi.StringPtrOutput                 `pulumi:"location"`
	Name       pulumi.StringOutput                    `pulumi:"name"`
	Properties MigrateProjectPropertiesResponseOutput `pulumi:"properties"`
	Tags       MigrateProjectResponseTagsPtrOutput    `pulumi:"tags"`
	Type       pulumi.StringOutput                    `pulumi:"type"`
}


func NewMigrateProject(ctx *pulumi.Context,
	name string, args *MigrateProjectArgs, opts ...pulumi.ResourceOption) (*MigrateProject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate:MigrateProject"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20200501:MigrateProject"),
		},
	})
	opts = append(opts, aliases)
	var resource MigrateProject
	err := ctx.RegisterResource("azure-native:migrate/v20180901preview:MigrateProject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMigrateProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MigrateProjectState, opts ...pulumi.ResourceOption) (*MigrateProject, error) {
	var resource MigrateProject
	err := ctx.ReadResource("azure-native:migrate/v20180901preview:MigrateProject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type migrateProjectState struct {
}

type MigrateProjectState struct {
}

func (MigrateProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*migrateProjectState)(nil)).Elem()
}

type migrateProjectArgs struct {
	ETag               *string                   `pulumi:"eTag"`
	Location           *string                   `pulumi:"location"`
	MigrateProjectName *string                   `pulumi:"migrateProjectName"`
	Properties         *MigrateProjectProperties `pulumi:"properties"`
	ResourceGroupName  string                    `pulumi:"resourceGroupName"`
	Tags               *MigrateProjectTags       `pulumi:"tags"`
}


type MigrateProjectArgs struct {
	ETag               pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	MigrateProjectName pulumi.StringPtrInput
	Properties         MigrateProjectPropertiesPtrInput
	ResourceGroupName  pulumi.StringInput
	Tags               MigrateProjectTagsPtrInput
}

func (MigrateProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*migrateProjectArgs)(nil)).Elem()
}

type MigrateProjectInput interface {
	pulumi.Input

	ToMigrateProjectOutput() MigrateProjectOutput
	ToMigrateProjectOutputWithContext(ctx context.Context) MigrateProjectOutput
}

func (*MigrateProject) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateProject)(nil)).Elem()
}

func (i *MigrateProject) ToMigrateProjectOutput() MigrateProjectOutput {
	return i.ToMigrateProjectOutputWithContext(context.Background())
}

func (i *MigrateProject) ToMigrateProjectOutputWithContext(ctx context.Context) MigrateProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectOutput)
}

type MigrateProjectOutput struct{ *pulumi.OutputState }

func (MigrateProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateProject)(nil)).Elem()
}

func (o MigrateProjectOutput) ToMigrateProjectOutput() MigrateProjectOutput {
	return o
}

func (o MigrateProjectOutput) ToMigrateProjectOutputWithContext(ctx context.Context) MigrateProjectOutput {
	return o
}

func (o MigrateProjectOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrateProject) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o MigrateProjectOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrateProject) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o MigrateProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MigrateProject) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MigrateProjectOutput) Properties() MigrateProjectPropertiesResponseOutput {
	return o.ApplyT(func(v *MigrateProject) MigrateProjectPropertiesResponseOutput { return v.Properties }).(MigrateProjectPropertiesResponseOutput)
}

func (o MigrateProjectOutput) Tags() MigrateProjectResponseTagsPtrOutput {
	return o.ApplyT(func(v *MigrateProject) MigrateProjectResponseTagsPtrOutput { return v.Tags }).(MigrateProjectResponseTagsPtrOutput)
}

func (o MigrateProjectOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MigrateProject) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MigrateProjectOutput{})
}

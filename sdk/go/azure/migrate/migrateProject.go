


package migrate

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
			Type: pulumi.String("azure-native:migrate/v20180901preview:MigrateProject"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20200501:MigrateProject"),
		},
	})
	opts = append(opts, aliases)
	var resource MigrateProject
	err := ctx.RegisterResource("azure-native:migrate:MigrateProject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMigrateProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MigrateProjectState, opts ...pulumi.ResourceOption) (*MigrateProject, error) {
	var resource MigrateProject
	err := ctx.ReadResource("azure-native:migrate:MigrateProject", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*MigrateProject)(nil))
}

func (i *MigrateProject) ToMigrateProjectOutput() MigrateProjectOutput {
	return i.ToMigrateProjectOutputWithContext(context.Background())
}

func (i *MigrateProject) ToMigrateProjectOutputWithContext(ctx context.Context) MigrateProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectOutput)
}

type MigrateProjectOutput struct{ *pulumi.OutputState }

func (MigrateProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrateProject)(nil))
}

func (o MigrateProjectOutput) ToMigrateProjectOutput() MigrateProjectOutput {
	return o
}

func (o MigrateProjectOutput) ToMigrateProjectOutputWithContext(ctx context.Context) MigrateProjectOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MigrateProjectOutput{})
}

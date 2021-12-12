


package v20200501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MigrateProjectsControllerMigrateProject struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput                 `pulumi:"eTag"`
	Location   pulumi.StringPtrOutput                 `pulumi:"location"`
	Name       pulumi.StringOutput                    `pulumi:"name"`
	Properties MigrateProjectPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput               `pulumi:"systemData"`
	Type       pulumi.StringOutput                    `pulumi:"type"`
}


func NewMigrateProjectsControllerMigrateProject(ctx *pulumi.Context,
	name string, args *MigrateProjectsControllerMigrateProjectArgs, opts ...pulumi.ResourceOption) (*MigrateProjectsControllerMigrateProject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate:MigrateProjectsControllerMigrateProject"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20180901preview:MigrateProjectsControllerMigrateProject"),
		},
	})
	opts = append(opts, aliases)
	var resource MigrateProjectsControllerMigrateProject
	err := ctx.RegisterResource("azure-native:migrate/v20200501:MigrateProjectsControllerMigrateProject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMigrateProjectsControllerMigrateProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MigrateProjectsControllerMigrateProjectState, opts ...pulumi.ResourceOption) (*MigrateProjectsControllerMigrateProject, error) {
	var resource MigrateProjectsControllerMigrateProject
	err := ctx.ReadResource("azure-native:migrate/v20200501:MigrateProjectsControllerMigrateProject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type migrateProjectsControllerMigrateProjectState struct {
}

type MigrateProjectsControllerMigrateProjectState struct {
}

func (MigrateProjectsControllerMigrateProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*migrateProjectsControllerMigrateProjectState)(nil)).Elem()
}

type migrateProjectsControllerMigrateProjectArgs struct {
	ETag               *string                   `pulumi:"eTag"`
	Location           *string                   `pulumi:"location"`
	MigrateProjectName *string                   `pulumi:"migrateProjectName"`
	Properties         *MigrateProjectProperties `pulumi:"properties"`
	ResourceGroupName  string                    `pulumi:"resourceGroupName"`
}


type MigrateProjectsControllerMigrateProjectArgs struct {
	ETag               pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	MigrateProjectName pulumi.StringPtrInput
	Properties         MigrateProjectPropertiesPtrInput
	ResourceGroupName  pulumi.StringInput
}

func (MigrateProjectsControllerMigrateProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*migrateProjectsControllerMigrateProjectArgs)(nil)).Elem()
}

type MigrateProjectsControllerMigrateProjectInput interface {
	pulumi.Input

	ToMigrateProjectsControllerMigrateProjectOutput() MigrateProjectsControllerMigrateProjectOutput
	ToMigrateProjectsControllerMigrateProjectOutputWithContext(ctx context.Context) MigrateProjectsControllerMigrateProjectOutput
}

func (*MigrateProjectsControllerMigrateProject) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrateProjectsControllerMigrateProject)(nil))
}

func (i *MigrateProjectsControllerMigrateProject) ToMigrateProjectsControllerMigrateProjectOutput() MigrateProjectsControllerMigrateProjectOutput {
	return i.ToMigrateProjectsControllerMigrateProjectOutputWithContext(context.Background())
}

func (i *MigrateProjectsControllerMigrateProject) ToMigrateProjectsControllerMigrateProjectOutputWithContext(ctx context.Context) MigrateProjectsControllerMigrateProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectsControllerMigrateProjectOutput)
}

type MigrateProjectsControllerMigrateProjectOutput struct{ *pulumi.OutputState }

func (MigrateProjectsControllerMigrateProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrateProjectsControllerMigrateProject)(nil))
}

func (o MigrateProjectsControllerMigrateProjectOutput) ToMigrateProjectsControllerMigrateProjectOutput() MigrateProjectsControllerMigrateProjectOutput {
	return o
}

func (o MigrateProjectsControllerMigrateProjectOutput) ToMigrateProjectsControllerMigrateProjectOutputWithContext(ctx context.Context) MigrateProjectsControllerMigrateProjectOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MigrateProjectsControllerMigrateProjectOutput{})
}

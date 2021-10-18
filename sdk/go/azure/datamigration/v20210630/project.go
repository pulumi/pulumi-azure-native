


package v20210630

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Project struct {
	pulumi.CustomResourceState

	CreationTime         pulumi.StringOutput             `pulumi:"creationTime"`
	DatabasesInfo        DatabaseInfoResponseArrayOutput `pulumi:"databasesInfo"`
	Location             pulumi.StringOutput             `pulumi:"location"`
	Name                 pulumi.StringOutput             `pulumi:"name"`
	ProvisioningState    pulumi.StringOutput             `pulumi:"provisioningState"`
	SourceConnectionInfo pulumi.AnyOutput                `pulumi:"sourceConnectionInfo"`
	SourcePlatform       pulumi.StringOutput             `pulumi:"sourcePlatform"`
	SystemData           SystemDataResponseOutput        `pulumi:"systemData"`
	Tags                 pulumi.StringMapOutput          `pulumi:"tags"`
	TargetConnectionInfo pulumi.AnyOutput                `pulumi:"targetConnectionInfo"`
	TargetPlatform       pulumi.StringOutput             `pulumi:"targetPlatform"`
	Type                 pulumi.StringOutput             `pulumi:"type"`
}


func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.SourcePlatform == nil {
		return nil, errors.New("invalid value for required argument 'SourcePlatform'")
	}
	if args.TargetPlatform == nil {
		return nil, errors.New("invalid value for required argument 'TargetPlatform'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:datamigration/v20210630:Project"),
		},
		{
			Type: pulumi.String("azure-native:datamigration:Project"),
		},
		{
			Type: pulumi.String("azure-nextgen:datamigration:Project"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20171115preview:Project"),
		},
		{
			Type: pulumi.String("azure-nextgen:datamigration/v20171115preview:Project"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180315preview:Project"),
		},
		{
			Type: pulumi.String("azure-nextgen:datamigration/v20180315preview:Project"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180331preview:Project"),
		},
		{
			Type: pulumi.String("azure-nextgen:datamigration/v20180331preview:Project"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180419:Project"),
		},
		{
			Type: pulumi.String("azure-nextgen:datamigration/v20180419:Project"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180715preview:Project"),
		},
		{
			Type: pulumi.String("azure-nextgen:datamigration/v20180715preview:Project"),
		},
	})
	opts = append(opts, aliases)
	var resource Project
	err := ctx.RegisterResource("azure-native:datamigration/v20210630:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("azure-native:datamigration/v20210630:Project", name, id, state, &resource, opts...)
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
	DatabasesInfo        []DatabaseInfo    `pulumi:"databasesInfo"`
	GroupName            string            `pulumi:"groupName"`
	Location             *string           `pulumi:"location"`
	ProjectName          *string           `pulumi:"projectName"`
	ServiceName          string            `pulumi:"serviceName"`
	SourceConnectionInfo interface{}       `pulumi:"sourceConnectionInfo"`
	SourcePlatform       string            `pulumi:"sourcePlatform"`
	Tags                 map[string]string `pulumi:"tags"`
	TargetConnectionInfo interface{}       `pulumi:"targetConnectionInfo"`
	TargetPlatform       string            `pulumi:"targetPlatform"`
}


type ProjectArgs struct {
	DatabasesInfo        DatabaseInfoArrayInput
	GroupName            pulumi.StringInput
	Location             pulumi.StringPtrInput
	ProjectName          pulumi.StringPtrInput
	ServiceName          pulumi.StringInput
	SourceConnectionInfo pulumi.Input
	SourcePlatform       pulumi.StringInput
	Tags                 pulumi.StringMapInput
	TargetConnectionInfo pulumi.Input
	TargetPlatform       pulumi.StringInput
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
	return reflect.TypeOf((*Project)(nil))
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Project)(nil))
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




package v20180715preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type File struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput              `pulumi:"etag"`
	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties ProjectFilePropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                 `pulumi:"type"`
}


func NewFile(ctx *pulumi.Context,
	name string, args *FileArgs, opts ...pulumi.ResourceOption) (*File, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datamigration:File"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20210630:File"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20211030preview:File"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20220130preview:File"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20220330preview:File"),
		},
	})
	opts = append(opts, aliases)
	var resource File
	err := ctx.RegisterResource("azure-native:datamigration/v20180715preview:File", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileState, opts ...pulumi.ResourceOption) (*File, error) {
	var resource File
	err := ctx.ReadResource("azure-native:datamigration/v20180715preview:File", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type fileState struct {
}

type FileState struct {
}

func (FileState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileState)(nil)).Elem()
}

type fileArgs struct {
	FileName    *string                `pulumi:"fileName"`
	GroupName   string                 `pulumi:"groupName"`
	ProjectName string                 `pulumi:"projectName"`
	Properties  *ProjectFileProperties `pulumi:"properties"`
	ServiceName string                 `pulumi:"serviceName"`
}


type FileArgs struct {
	FileName    pulumi.StringPtrInput
	GroupName   pulumi.StringInput
	ProjectName pulumi.StringInput
	Properties  ProjectFilePropertiesPtrInput
	ServiceName pulumi.StringInput
}

func (FileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileArgs)(nil)).Elem()
}

type FileInput interface {
	pulumi.Input

	ToFileOutput() FileOutput
	ToFileOutputWithContext(ctx context.Context) FileOutput
}

func (*File) ElementType() reflect.Type {
	return reflect.TypeOf((**File)(nil)).Elem()
}

func (i *File) ToFileOutput() FileOutput {
	return i.ToFileOutputWithContext(context.Background())
}

func (i *File) ToFileOutputWithContext(ctx context.Context) FileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileOutput)
}

type FileOutput struct{ *pulumi.OutputState }

func (FileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**File)(nil)).Elem()
}

func (o FileOutput) ToFileOutput() FileOutput {
	return o
}

func (o FileOutput) ToFileOutputWithContext(ctx context.Context) FileOutput {
	return o
}

func (o FileOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *File) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o FileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FileOutput) Properties() ProjectFilePropertiesResponseOutput {
	return o.ApplyT(func(v *File) ProjectFilePropertiesResponseOutput { return v.Properties }).(ProjectFilePropertiesResponseOutput)
}

func (o FileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FileOutput{})
}

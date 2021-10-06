


package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ModelVersion struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput        `pulumi:"name"`
	Properties ModelVersionResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput   `pulumi:"systemData"`
	Type       pulumi.StringOutput        `pulumi:"type"`
}


func NewModelVersion(ctx *pulumi.Context,
	name string, args *ModelVersionArgs, opts ...pulumi.ResourceOption) (*ModelVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices:ModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:ModelVersion"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210301preview:ModelVersion"),
		},
	})
	opts = append(opts, aliases)
	var resource ModelVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices:ModelVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetModelVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModelVersionState, opts ...pulumi.ResourceOption) (*ModelVersion, error) {
	var resource ModelVersion
	err := ctx.ReadResource("azure-native:machinelearningservices:ModelVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type modelVersionState struct {
}

type ModelVersionState struct {
}

func (ModelVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*modelVersionState)(nil)).Elem()
}

type modelVersionArgs struct {
	Name              string           `pulumi:"name"`
	Properties        ModelVersionType `pulumi:"properties"`
	ResourceGroupName string           `pulumi:"resourceGroupName"`
	Version           *string          `pulumi:"version"`
	WorkspaceName     string           `pulumi:"workspaceName"`
}


type ModelVersionArgs struct {
	Name              pulumi.StringInput
	Properties        ModelVersionTypeInput
	ResourceGroupName pulumi.StringInput
	Version           pulumi.StringPtrInput
	WorkspaceName     pulumi.StringInput
}

func (ModelVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modelVersionArgs)(nil)).Elem()
}

type ModelVersionInput interface {
	pulumi.Input

	ToModelVersionOutput() ModelVersionOutput
	ToModelVersionOutputWithContext(ctx context.Context) ModelVersionOutput
}

func (*ModelVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelVersion)(nil))
}

func (i *ModelVersion) ToModelVersionOutput() ModelVersionOutput {
	return i.ToModelVersionOutputWithContext(context.Background())
}

func (i *ModelVersion) ToModelVersionOutputWithContext(ctx context.Context) ModelVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelVersionOutput)
}

type ModelVersionOutput struct{ *pulumi.OutputState }

func (ModelVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelVersion)(nil))
}

func (o ModelVersionOutput) ToModelVersionOutput() ModelVersionOutput {
	return o
}

func (o ModelVersionOutput) ToModelVersionOutputWithContext(ctx context.Context) ModelVersionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ModelVersionOutput{})
}

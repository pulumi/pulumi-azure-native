


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnvironmentSpecificationVersion struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                           `pulumi:"name"`
	Properties EnvironmentSpecificationVersionResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                      `pulumi:"systemData"`
	Type       pulumi.StringOutput                           `pulumi:"type"`
}


func NewEnvironmentSpecificationVersion(ctx *pulumi.Context,
	name string, args *EnvironmentSpecificationVersionArgs, opts ...pulumi.ResourceOption) (*EnvironmentSpecificationVersion, error) {
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
			Type: pulumi.String("azure-native:machinelearningservices:EnvironmentSpecificationVersion"),
		},
	})
	opts = append(opts, aliases)
	var resource EnvironmentSpecificationVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20210301preview:EnvironmentSpecificationVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEnvironmentSpecificationVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentSpecificationVersionState, opts ...pulumi.ResourceOption) (*EnvironmentSpecificationVersion, error) {
	var resource EnvironmentSpecificationVersion
	err := ctx.ReadResource("azure-native:machinelearningservices/v20210301preview:EnvironmentSpecificationVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type environmentSpecificationVersionState struct {
}

type EnvironmentSpecificationVersionState struct {
}

func (EnvironmentSpecificationVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentSpecificationVersionState)(nil)).Elem()
}

type environmentSpecificationVersionArgs struct {
	Name              string                              `pulumi:"name"`
	Properties        EnvironmentSpecificationVersionType `pulumi:"properties"`
	ResourceGroupName string                              `pulumi:"resourceGroupName"`
	Version           *string                             `pulumi:"version"`
	WorkspaceName     string                              `pulumi:"workspaceName"`
}


type EnvironmentSpecificationVersionArgs struct {
	Name              pulumi.StringInput
	Properties        EnvironmentSpecificationVersionTypeInput
	ResourceGroupName pulumi.StringInput
	Version           pulumi.StringPtrInput
	WorkspaceName     pulumi.StringInput
}

func (EnvironmentSpecificationVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentSpecificationVersionArgs)(nil)).Elem()
}

type EnvironmentSpecificationVersionInput interface {
	pulumi.Input

	ToEnvironmentSpecificationVersionOutput() EnvironmentSpecificationVersionOutput
	ToEnvironmentSpecificationVersionOutputWithContext(ctx context.Context) EnvironmentSpecificationVersionOutput
}

func (*EnvironmentSpecificationVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentSpecificationVersion)(nil))
}

func (i *EnvironmentSpecificationVersion) ToEnvironmentSpecificationVersionOutput() EnvironmentSpecificationVersionOutput {
	return i.ToEnvironmentSpecificationVersionOutputWithContext(context.Background())
}

func (i *EnvironmentSpecificationVersion) ToEnvironmentSpecificationVersionOutputWithContext(ctx context.Context) EnvironmentSpecificationVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentSpecificationVersionOutput)
}

type EnvironmentSpecificationVersionOutput struct{ *pulumi.OutputState }

func (EnvironmentSpecificationVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentSpecificationVersion)(nil))
}

func (o EnvironmentSpecificationVersionOutput) ToEnvironmentSpecificationVersionOutput() EnvironmentSpecificationVersionOutput {
	return o
}

func (o EnvironmentSpecificationVersionOutput) ToEnvironmentSpecificationVersionOutputWithContext(ctx context.Context) EnvironmentSpecificationVersionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EnvironmentSpecificationVersionOutput{})
}

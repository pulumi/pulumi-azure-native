


package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CodeVersion struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput       `pulumi:"name"`
	Properties CodeVersionResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput  `pulumi:"systemData"`
	Type       pulumi.StringOutput       `pulumi:"type"`
}


func NewCodeVersion(ctx *pulumi.Context,
	name string, args *CodeVersionArgs, opts ...pulumi.ResourceOption) (*CodeVersion, error) {
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
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:CodeVersion"),
		},
	})
	opts = append(opts, aliases)
	var resource CodeVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices:CodeVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCodeVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CodeVersionState, opts ...pulumi.ResourceOption) (*CodeVersion, error) {
	var resource CodeVersion
	err := ctx.ReadResource("azure-native:machinelearningservices:CodeVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type codeVersionState struct {
}

type CodeVersionState struct {
}

func (CodeVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*codeVersionState)(nil)).Elem()
}

type codeVersionArgs struct {
	Name              string          `pulumi:"name"`
	Properties        CodeVersionType `pulumi:"properties"`
	ResourceGroupName string          `pulumi:"resourceGroupName"`
	Version           *string         `pulumi:"version"`
	WorkspaceName     string          `pulumi:"workspaceName"`
}


type CodeVersionArgs struct {
	Name              pulumi.StringInput
	Properties        CodeVersionTypeInput
	ResourceGroupName pulumi.StringInput
	Version           pulumi.StringPtrInput
	WorkspaceName     pulumi.StringInput
}

func (CodeVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*codeVersionArgs)(nil)).Elem()
}

type CodeVersionInput interface {
	pulumi.Input

	ToCodeVersionOutput() CodeVersionOutput
	ToCodeVersionOutputWithContext(ctx context.Context) CodeVersionOutput
}

func (*CodeVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**CodeVersion)(nil)).Elem()
}

func (i *CodeVersion) ToCodeVersionOutput() CodeVersionOutput {
	return i.ToCodeVersionOutputWithContext(context.Background())
}

func (i *CodeVersion) ToCodeVersionOutputWithContext(ctx context.Context) CodeVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodeVersionOutput)
}

type CodeVersionOutput struct{ *pulumi.OutputState }

func (CodeVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodeVersion)(nil)).Elem()
}

func (o CodeVersionOutput) ToCodeVersionOutput() CodeVersionOutput {
	return o
}

func (o CodeVersionOutput) ToCodeVersionOutputWithContext(ctx context.Context) CodeVersionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CodeVersionOutput{})
}

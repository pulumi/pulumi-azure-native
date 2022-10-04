


package v20220501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComponentVersion struct {
	pulumi.CustomResourceState

	ComponentVersionProperties ComponentVersionResponseOutput `pulumi:"componentVersionProperties"`
	Name                       pulumi.StringOutput            `pulumi:"name"`
	SystemData                 SystemDataResponseOutput       `pulumi:"systemData"`
	Type                       pulumi.StringOutput            `pulumi:"type"`
}


func NewComponentVersion(ctx *pulumi.Context,
	name string, args *ComponentVersionArgs, opts ...pulumi.ResourceOption) (*ComponentVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComponentVersionProperties == nil {
		return nil, errors.New("invalid value for required argument 'ComponentVersionProperties'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.ComponentVersionProperties = args.ComponentVersionProperties.ToComponentVersionTypeOutput().ApplyT(func(v ComponentVersionType) ComponentVersionType { return *v.Defaults() }).(ComponentVersionTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:ComponentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:ComponentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:ComponentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:ComponentVersion"),
		},
	})
	opts = append(opts, aliases)
	var resource ComponentVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20220501:ComponentVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetComponentVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComponentVersionState, opts ...pulumi.ResourceOption) (*ComponentVersion, error) {
	var resource ComponentVersion
	err := ctx.ReadResource("azure-native:machinelearningservices/v20220501:ComponentVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type componentVersionState struct {
}

type ComponentVersionState struct {
}

func (ComponentVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*componentVersionState)(nil)).Elem()
}

type componentVersionArgs struct {
	ComponentVersionProperties ComponentVersionType `pulumi:"componentVersionProperties"`
	Name                       string               `pulumi:"name"`
	ResourceGroupName          string               `pulumi:"resourceGroupName"`
	Version                    *string              `pulumi:"version"`
	WorkspaceName              string               `pulumi:"workspaceName"`
}


type ComponentVersionArgs struct {
	ComponentVersionProperties ComponentVersionTypeInput
	Name                       pulumi.StringInput
	ResourceGroupName          pulumi.StringInput
	Version                    pulumi.StringPtrInput
	WorkspaceName              pulumi.StringInput
}

func (ComponentVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentVersionArgs)(nil)).Elem()
}

type ComponentVersionInput interface {
	pulumi.Input

	ToComponentVersionOutput() ComponentVersionOutput
	ToComponentVersionOutputWithContext(ctx context.Context) ComponentVersionOutput
}

func (*ComponentVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**ComponentVersion)(nil)).Elem()
}

func (i *ComponentVersion) ToComponentVersionOutput() ComponentVersionOutput {
	return i.ToComponentVersionOutputWithContext(context.Background())
}

func (i *ComponentVersion) ToComponentVersionOutputWithContext(ctx context.Context) ComponentVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentVersionOutput)
}

type ComponentVersionOutput struct{ *pulumi.OutputState }

func (ComponentVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComponentVersion)(nil)).Elem()
}

func (o ComponentVersionOutput) ToComponentVersionOutput() ComponentVersionOutput {
	return o
}

func (o ComponentVersionOutput) ToComponentVersionOutputWithContext(ctx context.Context) ComponentVersionOutput {
	return o
}

func (o ComponentVersionOutput) ComponentVersionProperties() ComponentVersionResponseOutput {
	return o.ApplyT(func(v *ComponentVersion) ComponentVersionResponseOutput { return v.ComponentVersionProperties }).(ComponentVersionResponseOutput)
}

func (o ComponentVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComponentVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ComponentVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ComponentVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ComponentVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ComponentVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ComponentVersionOutput{})
}

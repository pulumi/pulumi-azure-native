


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnvironmentVersion struct {
	pulumi.CustomResourceState

	EnvironmentVersionDetails EnvironmentVersionResponseOutput `pulumi:"environmentVersionDetails"`
	Name                      pulumi.StringOutput              `pulumi:"name"`
	SystemData                SystemDataResponseOutput         `pulumi:"systemData"`
	Type                      pulumi.StringOutput              `pulumi:"type"`
}


func NewEnvironmentVersion(ctx *pulumi.Context,
	name string, args *EnvironmentVersionArgs, opts ...pulumi.ResourceOption) (*EnvironmentVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentVersionDetails == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentVersionDetails'")
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
	args.EnvironmentVersionDetails = args.EnvironmentVersionDetails.ToEnvironmentVersionTypeOutput().ApplyT(func(v EnvironmentVersionType) EnvironmentVersionType { return *v.Defaults() }).(EnvironmentVersionTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:EnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:EnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:EnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:EnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:EnvironmentVersion"),
		},
	})
	opts = append(opts, aliases)
	var resource EnvironmentVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20220201preview:EnvironmentVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEnvironmentVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentVersionState, opts ...pulumi.ResourceOption) (*EnvironmentVersion, error) {
	var resource EnvironmentVersion
	err := ctx.ReadResource("azure-native:machinelearningservices/v20220201preview:EnvironmentVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type environmentVersionState struct {
}

type EnvironmentVersionState struct {
}

func (EnvironmentVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentVersionState)(nil)).Elem()
}

type environmentVersionArgs struct {
	EnvironmentVersionDetails EnvironmentVersionType `pulumi:"environmentVersionDetails"`
	Name                      string                 `pulumi:"name"`
	ResourceGroupName         string                 `pulumi:"resourceGroupName"`
	Version                   *string                `pulumi:"version"`
	WorkspaceName             string                 `pulumi:"workspaceName"`
}


type EnvironmentVersionArgs struct {
	EnvironmentVersionDetails EnvironmentVersionTypeInput
	Name                      pulumi.StringInput
	ResourceGroupName         pulumi.StringInput
	Version                   pulumi.StringPtrInput
	WorkspaceName             pulumi.StringInput
}

func (EnvironmentVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentVersionArgs)(nil)).Elem()
}

type EnvironmentVersionInput interface {
	pulumi.Input

	ToEnvironmentVersionOutput() EnvironmentVersionOutput
	ToEnvironmentVersionOutputWithContext(ctx context.Context) EnvironmentVersionOutput
}

func (*EnvironmentVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentVersion)(nil)).Elem()
}

func (i *EnvironmentVersion) ToEnvironmentVersionOutput() EnvironmentVersionOutput {
	return i.ToEnvironmentVersionOutputWithContext(context.Background())
}

func (i *EnvironmentVersion) ToEnvironmentVersionOutputWithContext(ctx context.Context) EnvironmentVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentVersionOutput)
}

type EnvironmentVersionOutput struct{ *pulumi.OutputState }

func (EnvironmentVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentVersion)(nil)).Elem()
}

func (o EnvironmentVersionOutput) ToEnvironmentVersionOutput() EnvironmentVersionOutput {
	return o
}

func (o EnvironmentVersionOutput) ToEnvironmentVersionOutputWithContext(ctx context.Context) EnvironmentVersionOutput {
	return o
}

func (o EnvironmentVersionOutput) EnvironmentVersionDetails() EnvironmentVersionResponseOutput {
	return o.ApplyT(func(v *EnvironmentVersion) EnvironmentVersionResponseOutput { return v.EnvironmentVersionDetails }).(EnvironmentVersionResponseOutput)
}

func (o EnvironmentVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EnvironmentVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EnvironmentVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o EnvironmentVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EnvironmentVersionOutput{})
}

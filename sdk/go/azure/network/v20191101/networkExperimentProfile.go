


package v20191101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkExperimentProfile struct {
	pulumi.CustomResourceState

	EnabledState  pulumi.StringPtrOutput `pulumi:"enabledState"`
	Etag          pulumi.StringPtrOutput `pulumi:"etag"`
	Location      pulumi.StringPtrOutput `pulumi:"location"`
	Name          pulumi.StringOutput    `pulumi:"name"`
	ResourceState pulumi.StringOutput    `pulumi:"resourceState"`
	Tags          pulumi.StringMapOutput `pulumi:"tags"`
	Type          pulumi.StringOutput    `pulumi:"type"`
}


func NewNetworkExperimentProfile(ctx *pulumi.Context,
	name string, args *NetworkExperimentProfileArgs, opts ...pulumi.ResourceOption) (*NetworkExperimentProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NetworkExperimentProfile"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkExperimentProfile
	err := ctx.RegisterResource("azure-native:network/v20191101:NetworkExperimentProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkExperimentProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkExperimentProfileState, opts ...pulumi.ResourceOption) (*NetworkExperimentProfile, error) {
	var resource NetworkExperimentProfile
	err := ctx.ReadResource("azure-native:network/v20191101:NetworkExperimentProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type networkExperimentProfileState struct {
}

type NetworkExperimentProfileState struct {
}

func (NetworkExperimentProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkExperimentProfileState)(nil)).Elem()
}

type networkExperimentProfileArgs struct {
	EnabledState      *string           `pulumi:"enabledState"`
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	ProfileName       *string           `pulumi:"profileName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type NetworkExperimentProfileArgs struct {
	EnabledState      pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	ProfileName       pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (NetworkExperimentProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkExperimentProfileArgs)(nil)).Elem()
}

type NetworkExperimentProfileInput interface {
	pulumi.Input

	ToNetworkExperimentProfileOutput() NetworkExperimentProfileOutput
	ToNetworkExperimentProfileOutputWithContext(ctx context.Context) NetworkExperimentProfileOutput
}

func (*NetworkExperimentProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkExperimentProfile)(nil)).Elem()
}

func (i *NetworkExperimentProfile) ToNetworkExperimentProfileOutput() NetworkExperimentProfileOutput {
	return i.ToNetworkExperimentProfileOutputWithContext(context.Background())
}

func (i *NetworkExperimentProfile) ToNetworkExperimentProfileOutputWithContext(ctx context.Context) NetworkExperimentProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkExperimentProfileOutput)
}

type NetworkExperimentProfileOutput struct{ *pulumi.OutputState }

func (NetworkExperimentProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkExperimentProfile)(nil)).Elem()
}

func (o NetworkExperimentProfileOutput) ToNetworkExperimentProfileOutput() NetworkExperimentProfileOutput {
	return o
}

func (o NetworkExperimentProfileOutput) ToNetworkExperimentProfileOutputWithContext(ctx context.Context) NetworkExperimentProfileOutput {
	return o
}

func (o NetworkExperimentProfileOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkExperimentProfile) pulumi.StringPtrOutput { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o NetworkExperimentProfileOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkExperimentProfile) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o NetworkExperimentProfileOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkExperimentProfile) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NetworkExperimentProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkExperimentProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkExperimentProfileOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkExperimentProfile) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

func (o NetworkExperimentProfileOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkExperimentProfile) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkExperimentProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkExperimentProfile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkExperimentProfileOutput{})
}

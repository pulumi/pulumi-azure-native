


package network

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
			Type: pulumi.String("azure-nextgen:network:NetworkExperimentProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:NetworkExperimentProfile"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:NetworkExperimentProfile"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkExperimentProfile
	err := ctx.RegisterResource("azure-native:network:NetworkExperimentProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkExperimentProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkExperimentProfileState, opts ...pulumi.ResourceOption) (*NetworkExperimentProfile, error) {
	var resource NetworkExperimentProfile
	err := ctx.ReadResource("azure-native:network:NetworkExperimentProfile", name, id, state, &resource, opts...)
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
	Etag              *string           `pulumi:"etag"`
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	ProfileName       *string           `pulumi:"profileName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type NetworkExperimentProfileArgs struct {
	EnabledState      pulumi.StringPtrInput
	Etag              pulumi.StringPtrInput
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
	return reflect.TypeOf((*NetworkExperimentProfile)(nil))
}

func (i *NetworkExperimentProfile) ToNetworkExperimentProfileOutput() NetworkExperimentProfileOutput {
	return i.ToNetworkExperimentProfileOutputWithContext(context.Background())
}

func (i *NetworkExperimentProfile) ToNetworkExperimentProfileOutputWithContext(ctx context.Context) NetworkExperimentProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkExperimentProfileOutput)
}

type NetworkExperimentProfileOutput struct{ *pulumi.OutputState }

func (NetworkExperimentProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkExperimentProfile)(nil))
}

func (o NetworkExperimentProfileOutput) ToNetworkExperimentProfileOutput() NetworkExperimentProfileOutput {
	return o
}

func (o NetworkExperimentProfileOutput) ToNetworkExperimentProfileOutputWithContext(ctx context.Context) NetworkExperimentProfileOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkExperimentProfileInput)(nil)).Elem(), &NetworkExperimentProfile{})
	pulumi.RegisterOutputType(NetworkExperimentProfileOutput{})
}

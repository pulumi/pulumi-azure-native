


package v20180601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedPrivateEndpoint struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringOutput                  `pulumi:"etag"`
	Name       pulumi.StringOutput                  `pulumi:"name"`
	Properties ManagedPrivateEndpointResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                  `pulumi:"type"`
}


func NewManagedPrivateEndpoint(ctx *pulumi.Context,
	name string, args *ManagedPrivateEndpointArgs, opts ...pulumi.ResourceOption) (*ManagedPrivateEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FactoryName == nil {
		return nil, errors.New("invalid value for required argument 'FactoryName'")
	}
	if args.ManagedVirtualNetworkName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedVirtualNetworkName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datafactory:ManagedPrivateEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedPrivateEndpoint
	err := ctx.RegisterResource("azure-native:datafactory/v20180601:ManagedPrivateEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedPrivateEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedPrivateEndpointState, opts ...pulumi.ResourceOption) (*ManagedPrivateEndpoint, error) {
	var resource ManagedPrivateEndpoint
	err := ctx.ReadResource("azure-native:datafactory/v20180601:ManagedPrivateEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedPrivateEndpointState struct {
}

type ManagedPrivateEndpointState struct {
}

func (ManagedPrivateEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPrivateEndpointState)(nil)).Elem()
}

type managedPrivateEndpointArgs struct {
	FactoryName                string                     `pulumi:"factoryName"`
	ManagedPrivateEndpointName *string                    `pulumi:"managedPrivateEndpointName"`
	ManagedVirtualNetworkName  string                     `pulumi:"managedVirtualNetworkName"`
	Properties                 ManagedPrivateEndpointType `pulumi:"properties"`
	ResourceGroupName          string                     `pulumi:"resourceGroupName"`
}


type ManagedPrivateEndpointArgs struct {
	FactoryName                pulumi.StringInput
	ManagedPrivateEndpointName pulumi.StringPtrInput
	ManagedVirtualNetworkName  pulumi.StringInput
	Properties                 ManagedPrivateEndpointTypeInput
	ResourceGroupName          pulumi.StringInput
}

func (ManagedPrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPrivateEndpointArgs)(nil)).Elem()
}

type ManagedPrivateEndpointInput interface {
	pulumi.Input

	ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput
	ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput
}

func (*ManagedPrivateEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPrivateEndpoint)(nil)).Elem()
}

func (i *ManagedPrivateEndpoint) ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput {
	return i.ToManagedPrivateEndpointOutputWithContext(context.Background())
}

func (i *ManagedPrivateEndpoint) ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPrivateEndpointOutput)
}

type ManagedPrivateEndpointOutput struct{ *pulumi.OutputState }

func (ManagedPrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPrivateEndpoint)(nil)).Elem()
}

func (o ManagedPrivateEndpointOutput) ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput {
	return o
}

func (o ManagedPrivateEndpointOutput) ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput {
	return o
}

func (o ManagedPrivateEndpointOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ManagedPrivateEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedPrivateEndpointOutput) Properties() ManagedPrivateEndpointResponseOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) ManagedPrivateEndpointResponseOutput { return v.Properties }).(ManagedPrivateEndpointResponseOutput)
}

func (o ManagedPrivateEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedPrivateEndpointOutput{})
}

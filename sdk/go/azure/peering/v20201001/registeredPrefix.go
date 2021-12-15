


package v20201001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RegisteredPrefix struct {
	pulumi.CustomResourceState

	ErrorMessage            pulumi.StringOutput    `pulumi:"errorMessage"`
	Name                    pulumi.StringOutput    `pulumi:"name"`
	PeeringServicePrefixKey pulumi.StringOutput    `pulumi:"peeringServicePrefixKey"`
	Prefix                  pulumi.StringPtrOutput `pulumi:"prefix"`
	PrefixValidationState   pulumi.StringOutput    `pulumi:"prefixValidationState"`
	ProvisioningState       pulumi.StringOutput    `pulumi:"provisioningState"`
	Type                    pulumi.StringOutput    `pulumi:"type"`
}


func NewRegisteredPrefix(ctx *pulumi.Context,
	name string, args *RegisteredPrefixArgs, opts ...pulumi.ResourceOption) (*RegisteredPrefix, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PeeringName == nil {
		return nil, errors.New("invalid value for required argument 'PeeringName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:peering:RegisteredPrefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200101preview:RegisteredPrefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200401:RegisteredPrefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210101:RegisteredPrefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210601:RegisteredPrefix"),
		},
	})
	opts = append(opts, aliases)
	var resource RegisteredPrefix
	err := ctx.RegisterResource("azure-native:peering/v20201001:RegisteredPrefix", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRegisteredPrefix(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegisteredPrefixState, opts ...pulumi.ResourceOption) (*RegisteredPrefix, error) {
	var resource RegisteredPrefix
	err := ctx.ReadResource("azure-native:peering/v20201001:RegisteredPrefix", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type registeredPrefixState struct {
}

type RegisteredPrefixState struct {
}

func (RegisteredPrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*registeredPrefixState)(nil)).Elem()
}

type registeredPrefixArgs struct {
	PeeringName          string  `pulumi:"peeringName"`
	Prefix               *string `pulumi:"prefix"`
	RegisteredPrefixName *string `pulumi:"registeredPrefixName"`
	ResourceGroupName    string  `pulumi:"resourceGroupName"`
}


type RegisteredPrefixArgs struct {
	PeeringName          pulumi.StringInput
	Prefix               pulumi.StringPtrInput
	RegisteredPrefixName pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
}

func (RegisteredPrefixArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registeredPrefixArgs)(nil)).Elem()
}

type RegisteredPrefixInput interface {
	pulumi.Input

	ToRegisteredPrefixOutput() RegisteredPrefixOutput
	ToRegisteredPrefixOutputWithContext(ctx context.Context) RegisteredPrefixOutput
}

func (*RegisteredPrefix) ElementType() reflect.Type {
	return reflect.TypeOf((**RegisteredPrefix)(nil)).Elem()
}

func (i *RegisteredPrefix) ToRegisteredPrefixOutput() RegisteredPrefixOutput {
	return i.ToRegisteredPrefixOutputWithContext(context.Background())
}

func (i *RegisteredPrefix) ToRegisteredPrefixOutputWithContext(ctx context.Context) RegisteredPrefixOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegisteredPrefixOutput)
}

type RegisteredPrefixOutput struct{ *pulumi.OutputState }

func (RegisteredPrefixOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegisteredPrefix)(nil)).Elem()
}

func (o RegisteredPrefixOutput) ToRegisteredPrefixOutput() RegisteredPrefixOutput {
	return o
}

func (o RegisteredPrefixOutput) ToRegisteredPrefixOutputWithContext(ctx context.Context) RegisteredPrefixOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegisteredPrefixOutput{})
}

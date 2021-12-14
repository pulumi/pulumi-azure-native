


package v20211001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProtectionContainer struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput `pulumi:"eTag"`
	Location   pulumi.StringPtrOutput `pulumi:"location"`
	Name       pulumi.StringOutput    `pulumi:"name"`
	Properties pulumi.AnyOutput       `pulumi:"properties"`
	Tags       pulumi.StringMapOutput `pulumi:"tags"`
	Type       pulumi.StringOutput    `pulumi:"type"`
}


func NewProtectionContainer(ctx *pulumi.Context,
	name string, args *ProtectionContainerArgs, opts ...pulumi.ResourceOption) (*ProtectionContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FabricName == nil {
		return nil, errors.New("invalid value for required argument 'FabricName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VaultName == nil {
		return nil, errors.New("invalid value for required argument 'VaultName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20161201:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20201001:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20201201:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210101:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210201:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210201preview:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ProtectionContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource ProtectionContainer
	err := ctx.RegisterResource("azure-native:recoveryservices/v20211001:ProtectionContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProtectionContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProtectionContainerState, opts ...pulumi.ResourceOption) (*ProtectionContainer, error) {
	var resource ProtectionContainer
	err := ctx.ReadResource("azure-native:recoveryservices/v20211001:ProtectionContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type protectionContainerState struct {
}

type ProtectionContainerState struct {
}

func (ProtectionContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*protectionContainerState)(nil)).Elem()
}

type protectionContainerArgs struct {
	ContainerName     *string           `pulumi:"containerName"`
	ETag              *string           `pulumi:"eTag"`
	FabricName        string            `pulumi:"fabricName"`
	Location          *string           `pulumi:"location"`
	Properties        interface{}       `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	VaultName         string            `pulumi:"vaultName"`
}


type ProtectionContainerArgs struct {
	ContainerName     pulumi.StringPtrInput
	ETag              pulumi.StringPtrInput
	FabricName        pulumi.StringInput
	Location          pulumi.StringPtrInput
	Properties        pulumi.Input
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	VaultName         pulumi.StringInput
}

func (ProtectionContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*protectionContainerArgs)(nil)).Elem()
}

type ProtectionContainerInput interface {
	pulumi.Input

	ToProtectionContainerOutput() ProtectionContainerOutput
	ToProtectionContainerOutputWithContext(ctx context.Context) ProtectionContainerOutput
}

func (*ProtectionContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionContainer)(nil)).Elem()
}

func (i *ProtectionContainer) ToProtectionContainerOutput() ProtectionContainerOutput {
	return i.ToProtectionContainerOutputWithContext(context.Background())
}

func (i *ProtectionContainer) ToProtectionContainerOutputWithContext(ctx context.Context) ProtectionContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionContainerOutput)
}

type ProtectionContainerOutput struct{ *pulumi.OutputState }

func (ProtectionContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionContainer)(nil)).Elem()
}

func (o ProtectionContainerOutput) ToProtectionContainerOutput() ProtectionContainerOutput {
	return o
}

func (o ProtectionContainerOutput) ToProtectionContainerOutputWithContext(ctx context.Context) ProtectionContainerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProtectionContainerOutput{})
}




package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProtectedItem struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput `pulumi:"eTag"`
	Location   pulumi.StringPtrOutput `pulumi:"location"`
	Name       pulumi.StringOutput    `pulumi:"name"`
	Properties pulumi.AnyOutput       `pulumi:"properties"`
	Tags       pulumi.StringMapOutput `pulumi:"tags"`
	Type       pulumi.StringOutput    `pulumi:"type"`
}


func NewProtectedItem(ctx *pulumi.Context,
	name string, args *ProtectedItemArgs, opts ...pulumi.ResourceOption) (*ProtectedItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContainerName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerName'")
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
			Type: pulumi.String("azure-nextgen:recoveryservices/v20210801:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20160601:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20160601:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20190513:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20190513:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20190615:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20190615:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20201001:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20201001:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20201201:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20201201:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210101:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20210101:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210201:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20210201:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210201preview:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20210201preview:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20210210:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20210301:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20210401:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20210601:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20210701:ProtectedItem"),
		},
	})
	opts = append(opts, aliases)
	var resource ProtectedItem
	err := ctx.RegisterResource("azure-native:recoveryservices/v20210801:ProtectedItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProtectedItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProtectedItemState, opts ...pulumi.ResourceOption) (*ProtectedItem, error) {
	var resource ProtectedItem
	err := ctx.ReadResource("azure-native:recoveryservices/v20210801:ProtectedItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type protectedItemState struct {
}

type ProtectedItemState struct {
}

func (ProtectedItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*protectedItemState)(nil)).Elem()
}

type protectedItemArgs struct {
	ContainerName     string            `pulumi:"containerName"`
	ETag              *string           `pulumi:"eTag"`
	FabricName        string            `pulumi:"fabricName"`
	Location          *string           `pulumi:"location"`
	Properties        interface{}       `pulumi:"properties"`
	ProtectedItemName *string           `pulumi:"protectedItemName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	VaultName         string            `pulumi:"vaultName"`
}


type ProtectedItemArgs struct {
	ContainerName     pulumi.StringInput
	ETag              pulumi.StringPtrInput
	FabricName        pulumi.StringInput
	Location          pulumi.StringPtrInput
	Properties        pulumi.Input
	ProtectedItemName pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	VaultName         pulumi.StringInput
}

func (ProtectedItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*protectedItemArgs)(nil)).Elem()
}

type ProtectedItemInput interface {
	pulumi.Input

	ToProtectedItemOutput() ProtectedItemOutput
	ToProtectedItemOutputWithContext(ctx context.Context) ProtectedItemOutput
}

func (*ProtectedItem) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectedItem)(nil))
}

func (i *ProtectedItem) ToProtectedItemOutput() ProtectedItemOutput {
	return i.ToProtectedItemOutputWithContext(context.Background())
}

func (i *ProtectedItem) ToProtectedItemOutputWithContext(ctx context.Context) ProtectedItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectedItemOutput)
}

type ProtectedItemOutput struct{ *pulumi.OutputState }

func (ProtectedItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectedItem)(nil))
}

func (o ProtectedItemOutput) ToProtectedItemOutput() ProtectedItemOutput {
	return o
}

func (o ProtectedItemOutput) ToProtectedItemOutputWithContext(ctx context.Context) ProtectedItemOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProtectedItemOutput{})
}

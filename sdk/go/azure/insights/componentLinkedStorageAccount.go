


package insights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComponentLinkedStorageAccount struct {
	pulumi.CustomResourceState

	LinkedStorageAccount pulumi.StringPtrOutput `pulumi:"linkedStorageAccount"`
	Name                 pulumi.StringOutput    `pulumi:"name"`
	Type                 pulumi.StringOutput    `pulumi:"type"`
}


func NewComponentLinkedStorageAccount(ctx *pulumi.Context,
	name string, args *ComponentLinkedStorageAccountArgs, opts ...pulumi.ResourceOption) (*ComponentLinkedStorageAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:insights:ComponentLinkedStorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20200301preview:ComponentLinkedStorageAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights/v20200301preview:ComponentLinkedStorageAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource ComponentLinkedStorageAccount
	err := ctx.RegisterResource("azure-native:insights:ComponentLinkedStorageAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetComponentLinkedStorageAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComponentLinkedStorageAccountState, opts ...pulumi.ResourceOption) (*ComponentLinkedStorageAccount, error) {
	var resource ComponentLinkedStorageAccount
	err := ctx.ReadResource("azure-native:insights:ComponentLinkedStorageAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type componentLinkedStorageAccountState struct {
}

type ComponentLinkedStorageAccountState struct {
}

func (ComponentLinkedStorageAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*componentLinkedStorageAccountState)(nil)).Elem()
}

type componentLinkedStorageAccountArgs struct {
	LinkedStorageAccount *string `pulumi:"linkedStorageAccount"`
	ResourceGroupName    string  `pulumi:"resourceGroupName"`
	ResourceName         string  `pulumi:"resourceName"`
	StorageType          *string `pulumi:"storageType"`
}


type ComponentLinkedStorageAccountArgs struct {
	LinkedStorageAccount pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	ResourceName         pulumi.StringInput
	StorageType          pulumi.StringPtrInput
}

func (ComponentLinkedStorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentLinkedStorageAccountArgs)(nil)).Elem()
}

type ComponentLinkedStorageAccountInput interface {
	pulumi.Input

	ToComponentLinkedStorageAccountOutput() ComponentLinkedStorageAccountOutput
	ToComponentLinkedStorageAccountOutputWithContext(ctx context.Context) ComponentLinkedStorageAccountOutput
}

func (*ComponentLinkedStorageAccount) ElementType() reflect.Type {
	return reflect.TypeOf((*ComponentLinkedStorageAccount)(nil))
}

func (i *ComponentLinkedStorageAccount) ToComponentLinkedStorageAccountOutput() ComponentLinkedStorageAccountOutput {
	return i.ToComponentLinkedStorageAccountOutputWithContext(context.Background())
}

func (i *ComponentLinkedStorageAccount) ToComponentLinkedStorageAccountOutputWithContext(ctx context.Context) ComponentLinkedStorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentLinkedStorageAccountOutput)
}

type ComponentLinkedStorageAccountOutput struct{ *pulumi.OutputState }

func (ComponentLinkedStorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComponentLinkedStorageAccount)(nil))
}

func (o ComponentLinkedStorageAccountOutput) ToComponentLinkedStorageAccountOutput() ComponentLinkedStorageAccountOutput {
	return o
}

func (o ComponentLinkedStorageAccountOutput) ToComponentLinkedStorageAccountOutputWithContext(ctx context.Context) ComponentLinkedStorageAccountOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComponentLinkedStorageAccountInput)(nil)).Elem(), &ComponentLinkedStorageAccount{})
	pulumi.RegisterOutputType(ComponentLinkedStorageAccountOutput{})
}

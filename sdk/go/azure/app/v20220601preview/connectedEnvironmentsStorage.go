


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectedEnvironmentsStorage struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                                 `pulumi:"name"`
	Properties ConnectedEnvironmentStorageResponsePropertiesOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                            `pulumi:"systemData"`
	Type       pulumi.StringOutput                                 `pulumi:"type"`
}


func NewConnectedEnvironmentsStorage(ctx *pulumi.Context,
	name string, args *ConnectedEnvironmentsStorageArgs, opts ...pulumi.ResourceOption) (*ConnectedEnvironmentsStorage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectedEnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'ConnectedEnvironmentName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ConnectedEnvironmentsStorage
	err := ctx.RegisterResource("azure-native:app/v20220601preview:ConnectedEnvironmentsStorage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnectedEnvironmentsStorage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectedEnvironmentsStorageState, opts ...pulumi.ResourceOption) (*ConnectedEnvironmentsStorage, error) {
	var resource ConnectedEnvironmentsStorage
	err := ctx.ReadResource("azure-native:app/v20220601preview:ConnectedEnvironmentsStorage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type connectedEnvironmentsStorageState struct {
}

type ConnectedEnvironmentsStorageState struct {
}

func (ConnectedEnvironmentsStorageState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedEnvironmentsStorageState)(nil)).Elem()
}

type connectedEnvironmentsStorageArgs struct {
	ConnectedEnvironmentName string                                 `pulumi:"connectedEnvironmentName"`
	Properties               *ConnectedEnvironmentStorageProperties `pulumi:"properties"`
	ResourceGroupName        string                                 `pulumi:"resourceGroupName"`
	StorageName              *string                                `pulumi:"storageName"`
}


type ConnectedEnvironmentsStorageArgs struct {
	ConnectedEnvironmentName pulumi.StringInput
	Properties               ConnectedEnvironmentStoragePropertiesPtrInput
	ResourceGroupName        pulumi.StringInput
	StorageName              pulumi.StringPtrInput
}

func (ConnectedEnvironmentsStorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedEnvironmentsStorageArgs)(nil)).Elem()
}

type ConnectedEnvironmentsStorageInput interface {
	pulumi.Input

	ToConnectedEnvironmentsStorageOutput() ConnectedEnvironmentsStorageOutput
	ToConnectedEnvironmentsStorageOutputWithContext(ctx context.Context) ConnectedEnvironmentsStorageOutput
}

func (*ConnectedEnvironmentsStorage) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedEnvironmentsStorage)(nil)).Elem()
}

func (i *ConnectedEnvironmentsStorage) ToConnectedEnvironmentsStorageOutput() ConnectedEnvironmentsStorageOutput {
	return i.ToConnectedEnvironmentsStorageOutputWithContext(context.Background())
}

func (i *ConnectedEnvironmentsStorage) ToConnectedEnvironmentsStorageOutputWithContext(ctx context.Context) ConnectedEnvironmentsStorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedEnvironmentsStorageOutput)
}

type ConnectedEnvironmentsStorageOutput struct{ *pulumi.OutputState }

func (ConnectedEnvironmentsStorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedEnvironmentsStorage)(nil)).Elem()
}

func (o ConnectedEnvironmentsStorageOutput) ToConnectedEnvironmentsStorageOutput() ConnectedEnvironmentsStorageOutput {
	return o
}

func (o ConnectedEnvironmentsStorageOutput) ToConnectedEnvironmentsStorageOutputWithContext(ctx context.Context) ConnectedEnvironmentsStorageOutput {
	return o
}

func (o ConnectedEnvironmentsStorageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsStorage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConnectedEnvironmentsStorageOutput) Properties() ConnectedEnvironmentStorageResponsePropertiesOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsStorage) ConnectedEnvironmentStorageResponsePropertiesOutput {
		return v.Properties
	}).(ConnectedEnvironmentStorageResponsePropertiesOutput)
}

func (o ConnectedEnvironmentsStorageOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsStorage) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ConnectedEnvironmentsStorageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsStorage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectedEnvironmentsStorageOutput{})
}

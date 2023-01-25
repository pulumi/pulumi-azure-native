


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedEnvironmentsStorage struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                               `pulumi:"name"`
	Properties ManagedEnvironmentStorageResponsePropertiesOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                          `pulumi:"systemData"`
	Type       pulumi.StringOutput                               `pulumi:"type"`
}


func NewManagedEnvironmentsStorage(ctx *pulumi.Context,
	name string, args *ManagedEnvironmentsStorageArgs, opts ...pulumi.ResourceOption) (*ManagedEnvironmentsStorage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvName == nil {
		return nil, errors.New("invalid value for required argument 'EnvName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app:ManagedEnvironmentsStorage"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220301:ManagedEnvironmentsStorage"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220601preview:ManagedEnvironmentsStorage"),
		},
		{
			Type: pulumi.String("azure-native:app/v20221001:ManagedEnvironmentsStorage"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedEnvironmentsStorage
	err := ctx.RegisterResource("azure-native:app/v20220101preview:ManagedEnvironmentsStorage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedEnvironmentsStorage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedEnvironmentsStorageState, opts ...pulumi.ResourceOption) (*ManagedEnvironmentsStorage, error) {
	var resource ManagedEnvironmentsStorage
	err := ctx.ReadResource("azure-native:app/v20220101preview:ManagedEnvironmentsStorage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedEnvironmentsStorageState struct {
}

type ManagedEnvironmentsStorageState struct {
}

func (ManagedEnvironmentsStorageState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedEnvironmentsStorageState)(nil)).Elem()
}

type managedEnvironmentsStorageArgs struct {
	EnvName           string                               `pulumi:"envName"`
	Name              *string                              `pulumi:"name"`
	Properties        *ManagedEnvironmentStorageProperties `pulumi:"properties"`
	ResourceGroupName string                               `pulumi:"resourceGroupName"`
}


type ManagedEnvironmentsStorageArgs struct {
	EnvName           pulumi.StringInput
	Name              pulumi.StringPtrInput
	Properties        ManagedEnvironmentStoragePropertiesPtrInput
	ResourceGroupName pulumi.StringInput
}

func (ManagedEnvironmentsStorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedEnvironmentsStorageArgs)(nil)).Elem()
}

type ManagedEnvironmentsStorageInput interface {
	pulumi.Input

	ToManagedEnvironmentsStorageOutput() ManagedEnvironmentsStorageOutput
	ToManagedEnvironmentsStorageOutputWithContext(ctx context.Context) ManagedEnvironmentsStorageOutput
}

func (*ManagedEnvironmentsStorage) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedEnvironmentsStorage)(nil)).Elem()
}

func (i *ManagedEnvironmentsStorage) ToManagedEnvironmentsStorageOutput() ManagedEnvironmentsStorageOutput {
	return i.ToManagedEnvironmentsStorageOutputWithContext(context.Background())
}

func (i *ManagedEnvironmentsStorage) ToManagedEnvironmentsStorageOutputWithContext(ctx context.Context) ManagedEnvironmentsStorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedEnvironmentsStorageOutput)
}

type ManagedEnvironmentsStorageOutput struct{ *pulumi.OutputState }

func (ManagedEnvironmentsStorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedEnvironmentsStorage)(nil)).Elem()
}

func (o ManagedEnvironmentsStorageOutput) ToManagedEnvironmentsStorageOutput() ManagedEnvironmentsStorageOutput {
	return o
}

func (o ManagedEnvironmentsStorageOutput) ToManagedEnvironmentsStorageOutputWithContext(ctx context.Context) ManagedEnvironmentsStorageOutput {
	return o
}

func (o ManagedEnvironmentsStorageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedEnvironmentsStorage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedEnvironmentsStorageOutput) Properties() ManagedEnvironmentStorageResponsePropertiesOutput {
	return o.ApplyT(func(v *ManagedEnvironmentsStorage) ManagedEnvironmentStorageResponsePropertiesOutput {
		return v.Properties
	}).(ManagedEnvironmentStorageResponsePropertiesOutput)
}

func (o ManagedEnvironmentsStorageOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagedEnvironmentsStorage) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ManagedEnvironmentsStorageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedEnvironmentsStorage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedEnvironmentsStorageOutput{})
}

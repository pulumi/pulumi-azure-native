


package v20201101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedInstanceKey struct {
	pulumi.CustomResourceState

	AutoRotationEnabled pulumi.BoolOutput   `pulumi:"autoRotationEnabled"`
	CreationDate        pulumi.StringOutput `pulumi:"creationDate"`
	Kind                pulumi.StringOutput `pulumi:"kind"`
	Name                pulumi.StringOutput `pulumi:"name"`
	Thumbprint          pulumi.StringOutput `pulumi:"thumbprint"`
	Type                pulumi.StringOutput `pulumi:"type"`
}


func NewManagedInstanceKey(ctx *pulumi.Context,
	name string, args *ManagedInstanceKeyArgs, opts ...pulumi.ResourceOption) (*ManagedInstanceKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerKeyType == nil {
		return nil, errors.New("invalid value for required argument 'ServerKeyType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:sql/v20201101preview:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-native:sql:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20171001preview:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20171001preview:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200202preview:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200801preview:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20210201preview:ManagedInstanceKey"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedInstanceKey
	err := ctx.RegisterResource("azure-native:sql/v20201101preview:ManagedInstanceKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedInstanceKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedInstanceKeyState, opts ...pulumi.ResourceOption) (*ManagedInstanceKey, error) {
	var resource ManagedInstanceKey
	err := ctx.ReadResource("azure-native:sql/v20201101preview:ManagedInstanceKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedInstanceKeyState struct {
}

type ManagedInstanceKeyState struct {
}

func (ManagedInstanceKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceKeyState)(nil)).Elem()
}

type managedInstanceKeyArgs struct {
	KeyName             *string `pulumi:"keyName"`
	ManagedInstanceName string  `pulumi:"managedInstanceName"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
	ServerKeyType       string  `pulumi:"serverKeyType"`
	Uri                 *string `pulumi:"uri"`
}


type ManagedInstanceKeyArgs struct {
	KeyName             pulumi.StringPtrInput
	ManagedInstanceName pulumi.StringInput
	ResourceGroupName   pulumi.StringInput
	ServerKeyType       pulumi.StringInput
	Uri                 pulumi.StringPtrInput
}

func (ManagedInstanceKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceKeyArgs)(nil)).Elem()
}

type ManagedInstanceKeyInput interface {
	pulumi.Input

	ToManagedInstanceKeyOutput() ManagedInstanceKeyOutput
	ToManagedInstanceKeyOutputWithContext(ctx context.Context) ManagedInstanceKeyOutput
}

func (*ManagedInstanceKey) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstanceKey)(nil))
}

func (i *ManagedInstanceKey) ToManagedInstanceKeyOutput() ManagedInstanceKeyOutput {
	return i.ToManagedInstanceKeyOutputWithContext(context.Background())
}

func (i *ManagedInstanceKey) ToManagedInstanceKeyOutputWithContext(ctx context.Context) ManagedInstanceKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceKeyOutput)
}

type ManagedInstanceKeyOutput struct{ *pulumi.OutputState }

func (ManagedInstanceKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstanceKey)(nil))
}

func (o ManagedInstanceKeyOutput) ToManagedInstanceKeyOutput() ManagedInstanceKeyOutput {
	return o
}

func (o ManagedInstanceKeyOutput) ToManagedInstanceKeyOutputWithContext(ctx context.Context) ManagedInstanceKeyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagedInstanceKeyOutput{})
}

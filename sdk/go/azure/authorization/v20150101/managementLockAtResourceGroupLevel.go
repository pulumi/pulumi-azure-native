


package v20150101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagementLockAtResourceGroupLevel struct {
	pulumi.CustomResourceState

	Level pulumi.StringPtrOutput `pulumi:"level"`
	Name  pulumi.StringPtrOutput `pulumi:"name"`
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	Type  pulumi.StringOutput    `pulumi:"type"`
}


func NewManagementLockAtResourceGroupLevel(ctx *pulumi.Context,
	name string, args *ManagementLockAtResourceGroupLevelArgs, opts ...pulumi.ResourceOption) (*ManagementLockAtResourceGroupLevel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization:ManagementLockAtResourceGroupLevel"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20160901:ManagementLockAtResourceGroupLevel"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20170401:ManagementLockAtResourceGroupLevel"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200501:ManagementLockAtResourceGroupLevel"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagementLockAtResourceGroupLevel
	err := ctx.RegisterResource("azure-native:authorization/v20150101:ManagementLockAtResourceGroupLevel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagementLockAtResourceGroupLevel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementLockAtResourceGroupLevelState, opts ...pulumi.ResourceOption) (*ManagementLockAtResourceGroupLevel, error) {
	var resource ManagementLockAtResourceGroupLevel
	err := ctx.ReadResource("azure-native:authorization/v20150101:ManagementLockAtResourceGroupLevel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managementLockAtResourceGroupLevelState struct {
}

type ManagementLockAtResourceGroupLevelState struct {
}

func (ManagementLockAtResourceGroupLevelState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementLockAtResourceGroupLevelState)(nil)).Elem()
}

type managementLockAtResourceGroupLevelArgs struct {
	Level             *string `pulumi:"level"`
	LockName          *string `pulumi:"lockName"`
	Name              *string `pulumi:"name"`
	Notes             *string `pulumi:"notes"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type ManagementLockAtResourceGroupLevelArgs struct {
	Level             pulumi.StringPtrInput
	LockName          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	Notes             pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (ManagementLockAtResourceGroupLevelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementLockAtResourceGroupLevelArgs)(nil)).Elem()
}

type ManagementLockAtResourceGroupLevelInput interface {
	pulumi.Input

	ToManagementLockAtResourceGroupLevelOutput() ManagementLockAtResourceGroupLevelOutput
	ToManagementLockAtResourceGroupLevelOutputWithContext(ctx context.Context) ManagementLockAtResourceGroupLevelOutput
}

func (*ManagementLockAtResourceGroupLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockAtResourceGroupLevel)(nil))
}

func (i *ManagementLockAtResourceGroupLevel) ToManagementLockAtResourceGroupLevelOutput() ManagementLockAtResourceGroupLevelOutput {
	return i.ToManagementLockAtResourceGroupLevelOutputWithContext(context.Background())
}

func (i *ManagementLockAtResourceGroupLevel) ToManagementLockAtResourceGroupLevelOutputWithContext(ctx context.Context) ManagementLockAtResourceGroupLevelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockAtResourceGroupLevelOutput)
}

type ManagementLockAtResourceGroupLevelOutput struct{ *pulumi.OutputState }

func (ManagementLockAtResourceGroupLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockAtResourceGroupLevel)(nil))
}

func (o ManagementLockAtResourceGroupLevelOutput) ToManagementLockAtResourceGroupLevelOutput() ManagementLockAtResourceGroupLevelOutput {
	return o
}

func (o ManagementLockAtResourceGroupLevelOutput) ToManagementLockAtResourceGroupLevelOutputWithContext(ctx context.Context) ManagementLockAtResourceGroupLevelOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagementLockAtResourceGroupLevelOutput{})
}




package authorization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagementLockAtSubscriptionLevel struct {
	pulumi.CustomResourceState

	Level  pulumi.StringOutput                    `pulumi:"level"`
	Name   pulumi.StringOutput                    `pulumi:"name"`
	Notes  pulumi.StringPtrOutput                 `pulumi:"notes"`
	Owners ManagementLockOwnerResponseArrayOutput `pulumi:"owners"`
	Type   pulumi.StringOutput                    `pulumi:"type"`
}


func NewManagementLockAtSubscriptionLevel(ctx *pulumi.Context,
	name string, args *ManagementLockAtSubscriptionLevelArgs, opts ...pulumi.ResourceOption) (*ManagementLockAtSubscriptionLevel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Level == nil {
		return nil, errors.New("invalid value for required argument 'Level'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization/v20150101:ManagementLockAtSubscriptionLevel"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20160901:ManagementLockAtSubscriptionLevel"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20170401:ManagementLockAtSubscriptionLevel"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200501:ManagementLockAtSubscriptionLevel"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagementLockAtSubscriptionLevel
	err := ctx.RegisterResource("azure-native:authorization:ManagementLockAtSubscriptionLevel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagementLockAtSubscriptionLevel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementLockAtSubscriptionLevelState, opts ...pulumi.ResourceOption) (*ManagementLockAtSubscriptionLevel, error) {
	var resource ManagementLockAtSubscriptionLevel
	err := ctx.ReadResource("azure-native:authorization:ManagementLockAtSubscriptionLevel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managementLockAtSubscriptionLevelState struct {
}

type ManagementLockAtSubscriptionLevelState struct {
}

func (ManagementLockAtSubscriptionLevelState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementLockAtSubscriptionLevelState)(nil)).Elem()
}

type managementLockAtSubscriptionLevelArgs struct {
	Level    string                `pulumi:"level"`
	LockName *string               `pulumi:"lockName"`
	Notes    *string               `pulumi:"notes"`
	Owners   []ManagementLockOwner `pulumi:"owners"`
}


type ManagementLockAtSubscriptionLevelArgs struct {
	Level    pulumi.StringInput
	LockName pulumi.StringPtrInput
	Notes    pulumi.StringPtrInput
	Owners   ManagementLockOwnerArrayInput
}

func (ManagementLockAtSubscriptionLevelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementLockAtSubscriptionLevelArgs)(nil)).Elem()
}

type ManagementLockAtSubscriptionLevelInput interface {
	pulumi.Input

	ToManagementLockAtSubscriptionLevelOutput() ManagementLockAtSubscriptionLevelOutput
	ToManagementLockAtSubscriptionLevelOutputWithContext(ctx context.Context) ManagementLockAtSubscriptionLevelOutput
}

func (*ManagementLockAtSubscriptionLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockAtSubscriptionLevel)(nil))
}

func (i *ManagementLockAtSubscriptionLevel) ToManagementLockAtSubscriptionLevelOutput() ManagementLockAtSubscriptionLevelOutput {
	return i.ToManagementLockAtSubscriptionLevelOutputWithContext(context.Background())
}

func (i *ManagementLockAtSubscriptionLevel) ToManagementLockAtSubscriptionLevelOutputWithContext(ctx context.Context) ManagementLockAtSubscriptionLevelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockAtSubscriptionLevelOutput)
}

type ManagementLockAtSubscriptionLevelOutput struct{ *pulumi.OutputState }

func (ManagementLockAtSubscriptionLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockAtSubscriptionLevel)(nil))
}

func (o ManagementLockAtSubscriptionLevelOutput) ToManagementLockAtSubscriptionLevelOutput() ManagementLockAtSubscriptionLevelOutput {
	return o
}

func (o ManagementLockAtSubscriptionLevelOutput) ToManagementLockAtSubscriptionLevelOutputWithContext(ctx context.Context) ManagementLockAtSubscriptionLevelOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagementLockAtSubscriptionLevelOutput{})
}

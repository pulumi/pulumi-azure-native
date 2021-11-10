


package v20170401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagementLockByScope struct {
	pulumi.CustomResourceState

	Level  pulumi.StringOutput                    `pulumi:"level"`
	Name   pulumi.StringOutput                    `pulumi:"name"`
	Notes  pulumi.StringPtrOutput                 `pulumi:"notes"`
	Owners ManagementLockOwnerResponseArrayOutput `pulumi:"owners"`
	Type   pulumi.StringOutput                    `pulumi:"type"`
}


func NewManagementLockByScope(ctx *pulumi.Context,
	name string, args *ManagementLockByScopeArgs, opts ...pulumi.ResourceOption) (*ManagementLockByScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Level == nil {
		return nil, errors.New("invalid value for required argument 'Level'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization:ManagementLockByScope"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20160901:ManagementLockByScope"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200501:ManagementLockByScope"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagementLockByScope
	err := ctx.RegisterResource("azure-native:authorization/v20170401:ManagementLockByScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagementLockByScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementLockByScopeState, opts ...pulumi.ResourceOption) (*ManagementLockByScope, error) {
	var resource ManagementLockByScope
	err := ctx.ReadResource("azure-native:authorization/v20170401:ManagementLockByScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managementLockByScopeState struct {
}

type ManagementLockByScopeState struct {
}

func (ManagementLockByScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementLockByScopeState)(nil)).Elem()
}

type managementLockByScopeArgs struct {
	Level    string                `pulumi:"level"`
	LockName *string               `pulumi:"lockName"`
	Notes    *string               `pulumi:"notes"`
	Owners   []ManagementLockOwner `pulumi:"owners"`
	Scope    string                `pulumi:"scope"`
}


type ManagementLockByScopeArgs struct {
	Level    pulumi.StringInput
	LockName pulumi.StringPtrInput
	Notes    pulumi.StringPtrInput
	Owners   ManagementLockOwnerArrayInput
	Scope    pulumi.StringInput
}

func (ManagementLockByScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementLockByScopeArgs)(nil)).Elem()
}

type ManagementLockByScopeInput interface {
	pulumi.Input

	ToManagementLockByScopeOutput() ManagementLockByScopeOutput
	ToManagementLockByScopeOutputWithContext(ctx context.Context) ManagementLockByScopeOutput
}

func (*ManagementLockByScope) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockByScope)(nil))
}

func (i *ManagementLockByScope) ToManagementLockByScopeOutput() ManagementLockByScopeOutput {
	return i.ToManagementLockByScopeOutputWithContext(context.Background())
}

func (i *ManagementLockByScope) ToManagementLockByScopeOutputWithContext(ctx context.Context) ManagementLockByScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockByScopeOutput)
}

type ManagementLockByScopeOutput struct{ *pulumi.OutputState }

func (ManagementLockByScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockByScope)(nil))
}

func (o ManagementLockByScopeOutput) ToManagementLockByScopeOutput() ManagementLockByScopeOutput {
	return o
}

func (o ManagementLockByScopeOutput) ToManagementLockByScopeOutputWithContext(ctx context.Context) ManagementLockByScopeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagementLockByScopeOutput{})
}




package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SettingByScope struct {
	pulumi.CustomResourceState

	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	Kind pulumi.StringOutput    `pulumi:"kind"`
	Name pulumi.StringOutput    `pulumi:"name"`
	Type pulumi.StringOutput    `pulumi:"type"`
}


func NewSettingByScope(ctx *pulumi.Context,
	name string, args *SettingByScopeArgs, opts ...pulumi.ResourceOption) (*SettingByScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	var resource SettingByScope
	err := ctx.RegisterResource("azure-native:costmanagement/v20221001preview:SettingByScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSettingByScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SettingByScopeState, opts ...pulumi.ResourceOption) (*SettingByScope, error) {
	var resource SettingByScope
	err := ctx.ReadResource("azure-native:costmanagement/v20221001preview:SettingByScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type settingByScopeState struct {
}

type SettingByScopeState struct {
}

func (SettingByScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*settingByScopeState)(nil)).Elem()
}

type settingByScopeArgs struct {
	ETag  *string `pulumi:"eTag"`
	Kind  string  `pulumi:"kind"`
	Scope string  `pulumi:"scope"`
	Type  *string `pulumi:"type"`
}


type SettingByScopeArgs struct {
	ETag  pulumi.StringPtrInput
	Kind  pulumi.StringInput
	Scope pulumi.StringInput
	Type  pulumi.StringPtrInput
}

func (SettingByScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*settingByScopeArgs)(nil)).Elem()
}

type SettingByScopeInput interface {
	pulumi.Input

	ToSettingByScopeOutput() SettingByScopeOutput
	ToSettingByScopeOutputWithContext(ctx context.Context) SettingByScopeOutput
}

func (*SettingByScope) ElementType() reflect.Type {
	return reflect.TypeOf((**SettingByScope)(nil)).Elem()
}

func (i *SettingByScope) ToSettingByScopeOutput() SettingByScopeOutput {
	return i.ToSettingByScopeOutputWithContext(context.Background())
}

func (i *SettingByScope) ToSettingByScopeOutputWithContext(ctx context.Context) SettingByScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingByScopeOutput)
}

type SettingByScopeOutput struct{ *pulumi.OutputState }

func (SettingByScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SettingByScope)(nil)).Elem()
}

func (o SettingByScopeOutput) ToSettingByScopeOutput() SettingByScopeOutput {
	return o
}

func (o SettingByScopeOutput) ToSettingByScopeOutputWithContext(ctx context.Context) SettingByScopeOutput {
	return o
}

func (o SettingByScopeOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SettingByScope) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o SettingByScopeOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *SettingByScope) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o SettingByScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SettingByScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SettingByScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SettingByScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SettingByScopeOutput{})
}




package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectedEnvironmentsDaprComponent struct {
	pulumi.CustomResourceState

	ComponentType        pulumi.StringPtrOutput          `pulumi:"componentType"`
	IgnoreErrors         pulumi.BoolPtrOutput            `pulumi:"ignoreErrors"`
	InitTimeout          pulumi.StringPtrOutput          `pulumi:"initTimeout"`
	Metadata             DaprMetadataResponseArrayOutput `pulumi:"metadata"`
	Name                 pulumi.StringOutput             `pulumi:"name"`
	Scopes               pulumi.StringArrayOutput        `pulumi:"scopes"`
	SecretStoreComponent pulumi.StringPtrOutput          `pulumi:"secretStoreComponent"`
	Secrets              SecretResponseArrayOutput       `pulumi:"secrets"`
	SystemData           SystemDataResponseOutput        `pulumi:"systemData"`
	Type                 pulumi.StringOutput             `pulumi:"type"`
	Version              pulumi.StringPtrOutput          `pulumi:"version"`
}


func NewConnectedEnvironmentsDaprComponent(ctx *pulumi.Context,
	name string, args *ConnectedEnvironmentsDaprComponentArgs, opts ...pulumi.ResourceOption) (*ConnectedEnvironmentsDaprComponent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectedEnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'ConnectedEnvironmentName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ConnectedEnvironmentsDaprComponent
	err := ctx.RegisterResource("azure-native:app/v20220601preview:ConnectedEnvironmentsDaprComponent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnectedEnvironmentsDaprComponent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectedEnvironmentsDaprComponentState, opts ...pulumi.ResourceOption) (*ConnectedEnvironmentsDaprComponent, error) {
	var resource ConnectedEnvironmentsDaprComponent
	err := ctx.ReadResource("azure-native:app/v20220601preview:ConnectedEnvironmentsDaprComponent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type connectedEnvironmentsDaprComponentState struct {
}

type ConnectedEnvironmentsDaprComponentState struct {
}

func (ConnectedEnvironmentsDaprComponentState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedEnvironmentsDaprComponentState)(nil)).Elem()
}

type connectedEnvironmentsDaprComponentArgs struct {
	ComponentName            *string        `pulumi:"componentName"`
	ComponentType            *string        `pulumi:"componentType"`
	ConnectedEnvironmentName string         `pulumi:"connectedEnvironmentName"`
	IgnoreErrors             *bool          `pulumi:"ignoreErrors"`
	InitTimeout              *string        `pulumi:"initTimeout"`
	Metadata                 []DaprMetadata `pulumi:"metadata"`
	ResourceGroupName        string         `pulumi:"resourceGroupName"`
	Scopes                   []string       `pulumi:"scopes"`
	SecretStoreComponent     *string        `pulumi:"secretStoreComponent"`
	Secrets                  []Secret       `pulumi:"secrets"`
	Version                  *string        `pulumi:"version"`
}


type ConnectedEnvironmentsDaprComponentArgs struct {
	ComponentName            pulumi.StringPtrInput
	ComponentType            pulumi.StringPtrInput
	ConnectedEnvironmentName pulumi.StringInput
	IgnoreErrors             pulumi.BoolPtrInput
	InitTimeout              pulumi.StringPtrInput
	Metadata                 DaprMetadataArrayInput
	ResourceGroupName        pulumi.StringInput
	Scopes                   pulumi.StringArrayInput
	SecretStoreComponent     pulumi.StringPtrInput
	Secrets                  SecretArrayInput
	Version                  pulumi.StringPtrInput
}

func (ConnectedEnvironmentsDaprComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedEnvironmentsDaprComponentArgs)(nil)).Elem()
}

type ConnectedEnvironmentsDaprComponentInput interface {
	pulumi.Input

	ToConnectedEnvironmentsDaprComponentOutput() ConnectedEnvironmentsDaprComponentOutput
	ToConnectedEnvironmentsDaprComponentOutputWithContext(ctx context.Context) ConnectedEnvironmentsDaprComponentOutput
}

func (*ConnectedEnvironmentsDaprComponent) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedEnvironmentsDaprComponent)(nil)).Elem()
}

func (i *ConnectedEnvironmentsDaprComponent) ToConnectedEnvironmentsDaprComponentOutput() ConnectedEnvironmentsDaprComponentOutput {
	return i.ToConnectedEnvironmentsDaprComponentOutputWithContext(context.Background())
}

func (i *ConnectedEnvironmentsDaprComponent) ToConnectedEnvironmentsDaprComponentOutputWithContext(ctx context.Context) ConnectedEnvironmentsDaprComponentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedEnvironmentsDaprComponentOutput)
}

type ConnectedEnvironmentsDaprComponentOutput struct{ *pulumi.OutputState }

func (ConnectedEnvironmentsDaprComponentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedEnvironmentsDaprComponent)(nil)).Elem()
}

func (o ConnectedEnvironmentsDaprComponentOutput) ToConnectedEnvironmentsDaprComponentOutput() ConnectedEnvironmentsDaprComponentOutput {
	return o
}

func (o ConnectedEnvironmentsDaprComponentOutput) ToConnectedEnvironmentsDaprComponentOutputWithContext(ctx context.Context) ConnectedEnvironmentsDaprComponentOutput {
	return o
}

func (o ConnectedEnvironmentsDaprComponentOutput) ComponentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) pulumi.StringPtrOutput { return v.ComponentType }).(pulumi.StringPtrOutput)
}

func (o ConnectedEnvironmentsDaprComponentOutput) IgnoreErrors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) pulumi.BoolPtrOutput { return v.IgnoreErrors }).(pulumi.BoolPtrOutput)
}

func (o ConnectedEnvironmentsDaprComponentOutput) InitTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) pulumi.StringPtrOutput { return v.InitTimeout }).(pulumi.StringPtrOutput)
}

func (o ConnectedEnvironmentsDaprComponentOutput) Metadata() DaprMetadataResponseArrayOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) DaprMetadataResponseArrayOutput { return v.Metadata }).(DaprMetadataResponseArrayOutput)
}

func (o ConnectedEnvironmentsDaprComponentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConnectedEnvironmentsDaprComponentOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

func (o ConnectedEnvironmentsDaprComponentOutput) SecretStoreComponent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) pulumi.StringPtrOutput { return v.SecretStoreComponent }).(pulumi.StringPtrOutput)
}

func (o ConnectedEnvironmentsDaprComponentOutput) Secrets() SecretResponseArrayOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) SecretResponseArrayOutput { return v.Secrets }).(SecretResponseArrayOutput)
}

func (o ConnectedEnvironmentsDaprComponentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ConnectedEnvironmentsDaprComponentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ConnectedEnvironmentsDaprComponentOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsDaprComponent) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectedEnvironmentsDaprComponentOutput{})
}

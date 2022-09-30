


package app

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DaprComponent struct {
	pulumi.CustomResourceState

	ComponentType pulumi.StringPtrOutput          `pulumi:"componentType"`
	IgnoreErrors  pulumi.BoolPtrOutput            `pulumi:"ignoreErrors"`
	InitTimeout   pulumi.StringPtrOutput          `pulumi:"initTimeout"`
	Metadata      DaprMetadataResponseArrayOutput `pulumi:"metadata"`
	Name          pulumi.StringOutput             `pulumi:"name"`
	Scopes        pulumi.StringArrayOutput        `pulumi:"scopes"`
	Secrets       SecretResponseArrayOutput       `pulumi:"secrets"`
	SystemData    SystemDataResponseOutput        `pulumi:"systemData"`
	Type          pulumi.StringOutput             `pulumi:"type"`
	Version       pulumi.StringPtrOutput          `pulumi:"version"`
}


func NewDaprComponent(ctx *pulumi.Context,
	name string, args *DaprComponentArgs, opts ...pulumi.ResourceOption) (*DaprComponent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app/v20220101preview:DaprComponent"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220301:DaprComponent"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220601preview:DaprComponent"),
		},
	})
	opts = append(opts, aliases)
	var resource DaprComponent
	err := ctx.RegisterResource("azure-native:app:DaprComponent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDaprComponent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DaprComponentState, opts ...pulumi.ResourceOption) (*DaprComponent, error) {
	var resource DaprComponent
	err := ctx.ReadResource("azure-native:app:DaprComponent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type daprComponentState struct {
}

type DaprComponentState struct {
}

func (DaprComponentState) ElementType() reflect.Type {
	return reflect.TypeOf((*daprComponentState)(nil)).Elem()
}

type daprComponentArgs struct {
	ComponentName     *string        `pulumi:"componentName"`
	ComponentType     *string        `pulumi:"componentType"`
	EnvironmentName   string         `pulumi:"environmentName"`
	IgnoreErrors      *bool          `pulumi:"ignoreErrors"`
	InitTimeout       *string        `pulumi:"initTimeout"`
	Metadata          []DaprMetadata `pulumi:"metadata"`
	ResourceGroupName string         `pulumi:"resourceGroupName"`
	Scopes            []string       `pulumi:"scopes"`
	Secrets           []Secret       `pulumi:"secrets"`
	Version           *string        `pulumi:"version"`
}


type DaprComponentArgs struct {
	ComponentName     pulumi.StringPtrInput
	ComponentType     pulumi.StringPtrInput
	EnvironmentName   pulumi.StringInput
	IgnoreErrors      pulumi.BoolPtrInput
	InitTimeout       pulumi.StringPtrInput
	Metadata          DaprMetadataArrayInput
	ResourceGroupName pulumi.StringInput
	Scopes            pulumi.StringArrayInput
	Secrets           SecretArrayInput
	Version           pulumi.StringPtrInput
}

func (DaprComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*daprComponentArgs)(nil)).Elem()
}

type DaprComponentInput interface {
	pulumi.Input

	ToDaprComponentOutput() DaprComponentOutput
	ToDaprComponentOutputWithContext(ctx context.Context) DaprComponentOutput
}

func (*DaprComponent) ElementType() reflect.Type {
	return reflect.TypeOf((**DaprComponent)(nil)).Elem()
}

func (i *DaprComponent) ToDaprComponentOutput() DaprComponentOutput {
	return i.ToDaprComponentOutputWithContext(context.Background())
}

func (i *DaprComponent) ToDaprComponentOutputWithContext(ctx context.Context) DaprComponentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaprComponentOutput)
}

type DaprComponentOutput struct{ *pulumi.OutputState }

func (DaprComponentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DaprComponent)(nil)).Elem()
}

func (o DaprComponentOutput) ToDaprComponentOutput() DaprComponentOutput {
	return o
}

func (o DaprComponentOutput) ToDaprComponentOutputWithContext(ctx context.Context) DaprComponentOutput {
	return o
}

func (o DaprComponentOutput) ComponentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DaprComponent) pulumi.StringPtrOutput { return v.ComponentType }).(pulumi.StringPtrOutput)
}

func (o DaprComponentOutput) IgnoreErrors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DaprComponent) pulumi.BoolPtrOutput { return v.IgnoreErrors }).(pulumi.BoolPtrOutput)
}

func (o DaprComponentOutput) InitTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DaprComponent) pulumi.StringPtrOutput { return v.InitTimeout }).(pulumi.StringPtrOutput)
}

func (o DaprComponentOutput) Metadata() DaprMetadataResponseArrayOutput {
	return o.ApplyT(func(v *DaprComponent) DaprMetadataResponseArrayOutput { return v.Metadata }).(DaprMetadataResponseArrayOutput)
}

func (o DaprComponentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DaprComponent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DaprComponentOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DaprComponent) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

func (o DaprComponentOutput) Secrets() SecretResponseArrayOutput {
	return o.ApplyT(func(v *DaprComponent) SecretResponseArrayOutput { return v.Secrets }).(SecretResponseArrayOutput)
}

func (o DaprComponentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DaprComponent) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DaprComponentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DaprComponent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o DaprComponentOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DaprComponent) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DaprComponentOutput{})
}

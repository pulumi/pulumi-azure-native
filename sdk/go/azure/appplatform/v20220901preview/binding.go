


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Binding struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                     `pulumi:"name"`
	Properties BindingResourcePropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                `pulumi:"systemData"`
	Type       pulumi.StringOutput                     `pulumi:"type"`
}


func NewBinding(ctx *pulumi.Context,
	name string, args *BindingArgs, opts ...pulumi.ResourceOption) (*Binding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppName == nil {
		return nil, errors.New("invalid value for required argument 'AppName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20200701:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20201101preview:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210601preview:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210901preview:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220401:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:Binding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:Binding"),
		},
	})
	opts = append(opts, aliases)
	var resource Binding
	err := ctx.RegisterResource("azure-native:appplatform/v20220901preview:Binding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BindingState, opts ...pulumi.ResourceOption) (*Binding, error) {
	var resource Binding
	err := ctx.ReadResource("azure-native:appplatform/v20220901preview:Binding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type bindingState struct {
}

type BindingState struct {
}

func (BindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*bindingState)(nil)).Elem()
}

type bindingArgs struct {
	AppName           string                     `pulumi:"appName"`
	BindingName       *string                    `pulumi:"bindingName"`
	Properties        *BindingResourceProperties `pulumi:"properties"`
	ResourceGroupName string                     `pulumi:"resourceGroupName"`
	ServiceName       string                     `pulumi:"serviceName"`
}


type BindingArgs struct {
	AppName           pulumi.StringInput
	BindingName       pulumi.StringPtrInput
	Properties        BindingResourcePropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (BindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bindingArgs)(nil)).Elem()
}

type BindingInput interface {
	pulumi.Input

	ToBindingOutput() BindingOutput
	ToBindingOutputWithContext(ctx context.Context) BindingOutput
}

func (*Binding) ElementType() reflect.Type {
	return reflect.TypeOf((**Binding)(nil)).Elem()
}

func (i *Binding) ToBindingOutput() BindingOutput {
	return i.ToBindingOutputWithContext(context.Background())
}

func (i *Binding) ToBindingOutputWithContext(ctx context.Context) BindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingOutput)
}

type BindingOutput struct{ *pulumi.OutputState }

func (BindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Binding)(nil)).Elem()
}

func (o BindingOutput) ToBindingOutput() BindingOutput {
	return o
}

func (o BindingOutput) ToBindingOutputWithContext(ctx context.Context) BindingOutput {
	return o
}

func (o BindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Binding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BindingOutput) Properties() BindingResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *Binding) BindingResourcePropertiesResponseOutput { return v.Properties }).(BindingResourcePropertiesResponseOutput)
}

func (o BindingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Binding) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o BindingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Binding) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BindingOutput{})
}

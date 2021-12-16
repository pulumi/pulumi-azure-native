


package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IngestionSetting struct {
	pulumi.CustomResourceState

	Name pulumi.StringOutput `pulumi:"name"`
	Type pulumi.StringOutput `pulumi:"type"`
}


func NewIngestionSetting(ctx *pulumi.Context,
	name string, args *IngestionSettingArgs, opts ...pulumi.ResourceOption) (*IngestionSetting, error) {
	if args == nil {
		args = &IngestionSettingArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security/v20210115preview:IngestionSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource IngestionSetting
	err := ctx.RegisterResource("azure-native:security:IngestionSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIngestionSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IngestionSettingState, opts ...pulumi.ResourceOption) (*IngestionSetting, error) {
	var resource IngestionSetting
	err := ctx.ReadResource("azure-native:security:IngestionSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type ingestionSettingState struct {
}

type IngestionSettingState struct {
}

func (IngestionSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*ingestionSettingState)(nil)).Elem()
}

type ingestionSettingArgs struct {
	IngestionSettingName *string `pulumi:"ingestionSettingName"`
}


type IngestionSettingArgs struct {
	IngestionSettingName pulumi.StringPtrInput
}

func (IngestionSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ingestionSettingArgs)(nil)).Elem()
}

type IngestionSettingInput interface {
	pulumi.Input

	ToIngestionSettingOutput() IngestionSettingOutput
	ToIngestionSettingOutputWithContext(ctx context.Context) IngestionSettingOutput
}

func (*IngestionSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**IngestionSetting)(nil)).Elem()
}

func (i *IngestionSetting) ToIngestionSettingOutput() IngestionSettingOutput {
	return i.ToIngestionSettingOutputWithContext(context.Background())
}

func (i *IngestionSetting) ToIngestionSettingOutputWithContext(ctx context.Context) IngestionSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngestionSettingOutput)
}

type IngestionSettingOutput struct{ *pulumi.OutputState }

func (IngestionSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IngestionSetting)(nil)).Elem()
}

func (o IngestionSettingOutput) ToIngestionSettingOutput() IngestionSettingOutput {
	return o
}

func (o IngestionSettingOutput) ToIngestionSettingOutputWithContext(ctx context.Context) IngestionSettingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IngestionSettingOutput{})
}

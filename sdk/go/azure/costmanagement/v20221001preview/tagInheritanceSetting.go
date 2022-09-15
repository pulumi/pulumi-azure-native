


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TagInheritanceSetting struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput                 `pulumi:"eTag"`
	Kind       pulumi.StringOutput                    `pulumi:"kind"`
	Name       pulumi.StringOutput                    `pulumi:"name"`
	Properties TagInheritancePropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                    `pulumi:"type"`
}


func NewTagInheritanceSetting(ctx *pulumi.Context,
	name string, args *TagInheritanceSettingArgs, opts ...pulumi.ResourceOption) (*TagInheritanceSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	args.Kind = pulumi.String("taginheritance")
	var resource TagInheritanceSetting
	err := ctx.RegisterResource("azure-native:costmanagement/v20221001preview:TagInheritanceSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTagInheritanceSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagInheritanceSettingState, opts ...pulumi.ResourceOption) (*TagInheritanceSetting, error) {
	var resource TagInheritanceSetting
	err := ctx.ReadResource("azure-native:costmanagement/v20221001preview:TagInheritanceSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type tagInheritanceSettingState struct {
}

type TagInheritanceSettingState struct {
}

func (TagInheritanceSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagInheritanceSettingState)(nil)).Elem()
}

type tagInheritanceSettingArgs struct {
	ETag       *string                   `pulumi:"eTag"`
	Kind       string                    `pulumi:"kind"`
	Properties *TagInheritanceProperties `pulumi:"properties"`
	Scope      string                    `pulumi:"scope"`
	Type       *string                   `pulumi:"type"`
}


type TagInheritanceSettingArgs struct {
	ETag       pulumi.StringPtrInput
	Kind       pulumi.StringInput
	Properties TagInheritancePropertiesPtrInput
	Scope      pulumi.StringInput
	Type       pulumi.StringPtrInput
}

func (TagInheritanceSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagInheritanceSettingArgs)(nil)).Elem()
}

type TagInheritanceSettingInput interface {
	pulumi.Input

	ToTagInheritanceSettingOutput() TagInheritanceSettingOutput
	ToTagInheritanceSettingOutputWithContext(ctx context.Context) TagInheritanceSettingOutput
}

func (*TagInheritanceSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**TagInheritanceSetting)(nil)).Elem()
}

func (i *TagInheritanceSetting) ToTagInheritanceSettingOutput() TagInheritanceSettingOutput {
	return i.ToTagInheritanceSettingOutputWithContext(context.Background())
}

func (i *TagInheritanceSetting) ToTagInheritanceSettingOutputWithContext(ctx context.Context) TagInheritanceSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagInheritanceSettingOutput)
}

type TagInheritanceSettingOutput struct{ *pulumi.OutputState }

func (TagInheritanceSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagInheritanceSetting)(nil)).Elem()
}

func (o TagInheritanceSettingOutput) ToTagInheritanceSettingOutput() TagInheritanceSettingOutput {
	return o
}

func (o TagInheritanceSettingOutput) ToTagInheritanceSettingOutputWithContext(ctx context.Context) TagInheritanceSettingOutput {
	return o
}

func (o TagInheritanceSettingOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagInheritanceSetting) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o TagInheritanceSettingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *TagInheritanceSetting) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o TagInheritanceSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TagInheritanceSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TagInheritanceSettingOutput) Properties() TagInheritancePropertiesResponseOutput {
	return o.ApplyT(func(v *TagInheritanceSetting) TagInheritancePropertiesResponseOutput { return v.Properties }).(TagInheritancePropertiesResponseOutput)
}

func (o TagInheritanceSettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TagInheritanceSetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TagInheritanceSettingOutput{})
}

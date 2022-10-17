


package v20170301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TagDescription struct {
	pulumi.CustomResourceState

	Description             pulumi.StringPtrOutput `pulumi:"description"`
	DisplayName             pulumi.StringPtrOutput `pulumi:"displayName"`
	ExternalDocsDescription pulumi.StringPtrOutput `pulumi:"externalDocsDescription"`
	ExternalDocsUrl         pulumi.StringPtrOutput `pulumi:"externalDocsUrl"`
	Name                    pulumi.StringOutput    `pulumi:"name"`
	Type                    pulumi.StringOutput    `pulumi:"type"`
}


func NewTagDescription(ctx *pulumi.Context,
	name string, args *TagDescriptionArgs, opts ...pulumi.ResourceOption) (*TagDescription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:TagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:TagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:TagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:TagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:TagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:TagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:TagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:TagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:TagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:TagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:TagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:TagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:TagDescription"),
		},
	})
	opts = append(opts, aliases)
	var resource TagDescription
	err := ctx.RegisterResource("azure-native:apimanagement/v20170301:TagDescription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTagDescription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagDescriptionState, opts ...pulumi.ResourceOption) (*TagDescription, error) {
	var resource TagDescription
	err := ctx.ReadResource("azure-native:apimanagement/v20170301:TagDescription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type tagDescriptionState struct {
}

type TagDescriptionState struct {
}

func (TagDescriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagDescriptionState)(nil)).Elem()
}

type tagDescriptionArgs struct {
	ApiId                   string  `pulumi:"apiId"`
	Description             *string `pulumi:"description"`
	ExternalDocsDescription *string `pulumi:"externalDocsDescription"`
	ExternalDocsUrl         *string `pulumi:"externalDocsUrl"`
	ResourceGroupName       string  `pulumi:"resourceGroupName"`
	ServiceName             string  `pulumi:"serviceName"`
	TagId                   *string `pulumi:"tagId"`
}


type TagDescriptionArgs struct {
	ApiId                   pulumi.StringInput
	Description             pulumi.StringPtrInput
	ExternalDocsDescription pulumi.StringPtrInput
	ExternalDocsUrl         pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	ServiceName             pulumi.StringInput
	TagId                   pulumi.StringPtrInput
}

func (TagDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagDescriptionArgs)(nil)).Elem()
}

type TagDescriptionInput interface {
	pulumi.Input

	ToTagDescriptionOutput() TagDescriptionOutput
	ToTagDescriptionOutputWithContext(ctx context.Context) TagDescriptionOutput
}

func (*TagDescription) ElementType() reflect.Type {
	return reflect.TypeOf((**TagDescription)(nil)).Elem()
}

func (i *TagDescription) ToTagDescriptionOutput() TagDescriptionOutput {
	return i.ToTagDescriptionOutputWithContext(context.Background())
}

func (i *TagDescription) ToTagDescriptionOutputWithContext(ctx context.Context) TagDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagDescriptionOutput)
}

type TagDescriptionOutput struct{ *pulumi.OutputState }

func (TagDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagDescription)(nil)).Elem()
}

func (o TagDescriptionOutput) ToTagDescriptionOutput() TagDescriptionOutput {
	return o
}

func (o TagDescriptionOutput) ToTagDescriptionOutputWithContext(ctx context.Context) TagDescriptionOutput {
	return o
}

func (o TagDescriptionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagDescription) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o TagDescriptionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagDescription) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o TagDescriptionOutput) ExternalDocsDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagDescription) pulumi.StringPtrOutput { return v.ExternalDocsDescription }).(pulumi.StringPtrOutput)
}

func (o TagDescriptionOutput) ExternalDocsUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagDescription) pulumi.StringPtrOutput { return v.ExternalDocsUrl }).(pulumi.StringPtrOutput)
}

func (o TagDescriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TagDescription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TagDescriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TagDescription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TagDescriptionOutput{})
}

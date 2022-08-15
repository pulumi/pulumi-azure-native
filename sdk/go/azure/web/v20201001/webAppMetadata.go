


package v20201001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppMetadata struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringPtrOutput   `pulumi:"kind"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	Properties pulumi.StringMapOutput   `pulumi:"properties"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewWebAppMetadata(ctx *pulumi.Context,
	name string, args *WebAppMetadataArgs, opts ...pulumi.ResourceOption) (*WebAppMetadata, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppMetadata"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppMetadata
	err := ctx.RegisterResource("azure-native:web/v20201001:WebAppMetadata", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppMetadata(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppMetadataState, opts ...pulumi.ResourceOption) (*WebAppMetadata, error) {
	var resource WebAppMetadata
	err := ctx.ReadResource("azure-native:web/v20201001:WebAppMetadata", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppMetadataState struct {
}

type WebAppMetadataState struct {
}

func (WebAppMetadataState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppMetadataState)(nil)).Elem()
}

type webAppMetadataArgs struct {
	Kind              *string           `pulumi:"kind"`
	Name              string            `pulumi:"name"`
	Properties        map[string]string `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
}


type WebAppMetadataArgs struct {
	Kind              pulumi.StringPtrInput
	Name              pulumi.StringInput
	Properties        pulumi.StringMapInput
	ResourceGroupName pulumi.StringInput
}

func (WebAppMetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppMetadataArgs)(nil)).Elem()
}

type WebAppMetadataInput interface {
	pulumi.Input

	ToWebAppMetadataOutput() WebAppMetadataOutput
	ToWebAppMetadataOutputWithContext(ctx context.Context) WebAppMetadataOutput
}

func (*WebAppMetadata) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppMetadata)(nil)).Elem()
}

func (i *WebAppMetadata) ToWebAppMetadataOutput() WebAppMetadataOutput {
	return i.ToWebAppMetadataOutputWithContext(context.Background())
}

func (i *WebAppMetadata) ToWebAppMetadataOutputWithContext(ctx context.Context) WebAppMetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppMetadataOutput)
}

type WebAppMetadataOutput struct{ *pulumi.OutputState }

func (WebAppMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppMetadata)(nil)).Elem()
}

func (o WebAppMetadataOutput) ToWebAppMetadataOutput() WebAppMetadataOutput {
	return o
}

func (o WebAppMetadataOutput) ToWebAppMetadataOutputWithContext(ctx context.Context) WebAppMetadataOutput {
	return o
}

func (o WebAppMetadataOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppMetadata) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppMetadataOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppMetadata) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppMetadataOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebAppMetadata) pulumi.StringMapOutput { return v.Properties }).(pulumi.StringMapOutput)
}

func (o WebAppMetadataOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebAppMetadata) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o WebAppMetadataOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppMetadata) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppMetadataOutput{})
}

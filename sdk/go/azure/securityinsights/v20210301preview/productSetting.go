


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ProductSetting struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput   `pulumi:"etag"`
	Kind       pulumi.StringOutput      `pulumi:"kind"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewProductSetting(ctx *pulumi.Context,
	name string, args *ProductSettingArgs, opts ...pulumi.ResourceOption) (*ProductSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("invalid value for required argument 'OperationalInsightsResourceProvider'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:ProductSetting"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:ProductSetting"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:ProductSetting"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:ProductSetting"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:ProductSetting"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:ProductSetting"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:ProductSetting"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:ProductSetting"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:ProductSetting"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:ProductSetting"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:ProductSetting"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:ProductSetting"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:ProductSetting"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:ProductSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource ProductSetting
	err := ctx.RegisterResource("azure-native:securityinsights/v20210301preview:ProductSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProductSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductSettingState, opts ...pulumi.ResourceOption) (*ProductSetting, error) {
	var resource ProductSetting
	err := ctx.ReadResource("azure-native:securityinsights/v20210301preview:ProductSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type productSettingState struct {
}

type ProductSettingState struct {
}

func (ProductSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*productSettingState)(nil)).Elem()
}

type productSettingArgs struct {
	Kind                                string  `pulumi:"kind"`
	OperationalInsightsResourceProvider string  `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string  `pulumi:"resourceGroupName"`
	SettingsName                        *string `pulumi:"settingsName"`
	WorkspaceName                       string  `pulumi:"workspaceName"`
}


type ProductSettingArgs struct {
	Kind                                pulumi.StringInput
	OperationalInsightsResourceProvider pulumi.StringInput
	ResourceGroupName                   pulumi.StringInput
	SettingsName                        pulumi.StringPtrInput
	WorkspaceName                       pulumi.StringInput
}

func (ProductSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productSettingArgs)(nil)).Elem()
}

type ProductSettingInput interface {
	pulumi.Input

	ToProductSettingOutput() ProductSettingOutput
	ToProductSettingOutputWithContext(ctx context.Context) ProductSettingOutput
}

func (*ProductSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductSetting)(nil)).Elem()
}

func (i *ProductSetting) ToProductSettingOutput() ProductSettingOutput {
	return i.ToProductSettingOutputWithContext(context.Background())
}

func (i *ProductSetting) ToProductSettingOutputWithContext(ctx context.Context) ProductSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductSettingOutput)
}

type ProductSettingOutput struct{ *pulumi.OutputState }

func (ProductSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductSetting)(nil)).Elem()
}

func (o ProductSettingOutput) ToProductSettingOutput() ProductSettingOutput {
	return o
}

func (o ProductSettingOutput) ToProductSettingOutputWithContext(ctx context.Context) ProductSettingOutput {
	return o
}

func (o ProductSettingOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductSetting) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ProductSettingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductSetting) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o ProductSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProductSettingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ProductSetting) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ProductSettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductSetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProductSettingOutput{})
}

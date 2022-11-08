


package v20210702preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IotHubResource struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput         `pulumi:"etag"`
	Identity   ArmIdentityResponsePtrOutput   `pulumi:"identity"`
	Location   pulumi.StringOutput            `pulumi:"location"`
	Name       pulumi.StringOutput            `pulumi:"name"`
	Properties IotHubPropertiesResponseOutput `pulumi:"properties"`
	Sku        IotHubSkuInfoResponseOutput    `pulumi:"sku"`
	SystemData SystemDataResponseOutput       `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput         `pulumi:"tags"`
	Type       pulumi.StringOutput            `pulumi:"type"`
}


func NewIotHubResource(ctx *pulumi.Context,
	name string, args *IotHubResourceArgs, opts ...pulumi.ResourceOption) (*IotHubResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToIotHubPropertiesPtrOutput().ApplyT(func(v *IotHubProperties) *IotHubProperties { return v.Defaults() }).(IotHubPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devices:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20160203:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20170119:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20170701:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20180122:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20180401:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20181201preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20190322:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20190322preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20190701preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20191104:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200301:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200401:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200615:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200710preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200801:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200831:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200831preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210201preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210303preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210331:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210701:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210701preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210702:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20220430preview:IotHubResource"),
		},
	})
	opts = append(opts, aliases)
	var resource IotHubResource
	err := ctx.RegisterResource("azure-native:devices/v20210702preview:IotHubResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIotHubResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotHubResourceState, opts ...pulumi.ResourceOption) (*IotHubResource, error) {
	var resource IotHubResource
	err := ctx.ReadResource("azure-native:devices/v20210702preview:IotHubResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type iotHubResourceState struct {
}

type IotHubResourceState struct {
}

func (IotHubResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubResourceState)(nil)).Elem()
}

type iotHubResourceArgs struct {
	Identity          *ArmIdentity      `pulumi:"identity"`
	Location          *string           `pulumi:"location"`
	Properties        *IotHubProperties `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceName      *string           `pulumi:"resourceName"`
	Sku               IotHubSkuInfo     `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
}


type IotHubResourceArgs struct {
	Identity          ArmIdentityPtrInput
	Location          pulumi.StringPtrInput
	Properties        IotHubPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Sku               IotHubSkuInfoInput
	Tags              pulumi.StringMapInput
}

func (IotHubResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubResourceArgs)(nil)).Elem()
}

type IotHubResourceInput interface {
	pulumi.Input

	ToIotHubResourceOutput() IotHubResourceOutput
	ToIotHubResourceOutputWithContext(ctx context.Context) IotHubResourceOutput
}

func (*IotHubResource) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubResource)(nil)).Elem()
}

func (i *IotHubResource) ToIotHubResourceOutput() IotHubResourceOutput {
	return i.ToIotHubResourceOutputWithContext(context.Background())
}

func (i *IotHubResource) ToIotHubResourceOutputWithContext(ctx context.Context) IotHubResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubResourceOutput)
}

type IotHubResourceOutput struct{ *pulumi.OutputState }

func (IotHubResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubResource)(nil)).Elem()
}

func (o IotHubResourceOutput) ToIotHubResourceOutput() IotHubResourceOutput {
	return o
}

func (o IotHubResourceOutput) ToIotHubResourceOutputWithContext(ctx context.Context) IotHubResourceOutput {
	return o
}

func (o IotHubResourceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotHubResource) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o IotHubResourceOutput) Identity() ArmIdentityResponsePtrOutput {
	return o.ApplyT(func(v *IotHubResource) ArmIdentityResponsePtrOutput { return v.Identity }).(ArmIdentityResponsePtrOutput)
}

func (o IotHubResourceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *IotHubResource) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o IotHubResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IotHubResource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IotHubResourceOutput) Properties() IotHubPropertiesResponseOutput {
	return o.ApplyT(func(v *IotHubResource) IotHubPropertiesResponseOutput { return v.Properties }).(IotHubPropertiesResponseOutput)
}

func (o IotHubResourceOutput) Sku() IotHubSkuInfoResponseOutput {
	return o.ApplyT(func(v *IotHubResource) IotHubSkuInfoResponseOutput { return v.Sku }).(IotHubSkuInfoResponseOutput)
}

func (o IotHubResourceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *IotHubResource) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o IotHubResourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IotHubResource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o IotHubResourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IotHubResource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IotHubResourceOutput{})
}

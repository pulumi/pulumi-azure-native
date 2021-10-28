


package devices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IotHubResource struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput         `pulumi:"etag"`
	Location   pulumi.StringOutput            `pulumi:"location"`
	Name       pulumi.StringOutput            `pulumi:"name"`
	Properties IotHubPropertiesResponseOutput `pulumi:"properties"`
	Sku        IotHubSkuInfoResponseOutput    `pulumi:"sku"`
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:devices:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20160203:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20160203:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20170119:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20170119:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20170701:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20170701:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20180122:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20180122:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20180401:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20180401:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20181201preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20181201preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20190322:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20190322:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20190322preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20190322preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20190701preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20190701preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20191104:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20191104:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200301:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20200301:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200401:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20200401:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200615:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20200615:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200710preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20200710preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200801:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20200801:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200831:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20200831:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200831preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20200831preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210201preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20210201preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210303preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20210303preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210331:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20210331:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210701:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20210701:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210701preview:IotHubResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20210701preview:IotHubResource"),
		},
	})
	opts = append(opts, aliases)
	var resource IotHubResource
	err := ctx.RegisterResource("azure-native:devices:IotHubResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIotHubResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotHubResourceState, opts ...pulumi.ResourceOption) (*IotHubResource, error) {
	var resource IotHubResource
	err := ctx.ReadResource("azure-native:devices:IotHubResource", name, id, state, &resource, opts...)
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
	Etag              *string           `pulumi:"etag"`
	Location          *string           `pulumi:"location"`
	Properties        *IotHubProperties `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceName      *string           `pulumi:"resourceName"`
	Sku               IotHubSkuInfo     `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
}


type IotHubResourceArgs struct {
	Etag              pulumi.StringPtrInput
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
	return reflect.TypeOf((*IotHubResource)(nil))
}

func (i *IotHubResource) ToIotHubResourceOutput() IotHubResourceOutput {
	return i.ToIotHubResourceOutputWithContext(context.Background())
}

func (i *IotHubResource) ToIotHubResourceOutputWithContext(ctx context.Context) IotHubResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubResourceOutput)
}

type IotHubResourceOutput struct{ *pulumi.OutputState }

func (IotHubResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubResource)(nil))
}

func (o IotHubResourceOutput) ToIotHubResourceOutput() IotHubResourceOutput {
	return o
}

func (o IotHubResourceOutput) ToIotHubResourceOutputWithContext(ctx context.Context) IotHubResourceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IotHubResourceOutput{})
}

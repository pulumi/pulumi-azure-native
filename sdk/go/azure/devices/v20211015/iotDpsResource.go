


package v20211015

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IotDpsResource struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput                    `pulumi:"etag"`
	Location   pulumi.StringOutput                       `pulumi:"location"`
	Name       pulumi.StringOutput                       `pulumi:"name"`
	Properties IotDpsPropertiesDescriptionResponseOutput `pulumi:"properties"`
	Sku        IotDpsSkuInfoResponseOutput               `pulumi:"sku"`
	SystemData SystemDataResponseOutput                  `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                    `pulumi:"tags"`
	Type       pulumi.StringOutput                       `pulumi:"type"`
}


func NewIotDpsResource(ctx *pulumi.Context,
	name string, args *IotDpsResourceArgs, opts ...pulumi.ResourceOption) (*IotDpsResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devices:IotDpsResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20170821preview:IotDpsResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20171115:IotDpsResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20180122:IotDpsResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200101:IotDpsResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200301:IotDpsResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200901preview:IotDpsResource"),
		},
	})
	opts = append(opts, aliases)
	var resource IotDpsResource
	err := ctx.RegisterResource("azure-native:devices/v20211015:IotDpsResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIotDpsResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotDpsResourceState, opts ...pulumi.ResourceOption) (*IotDpsResource, error) {
	var resource IotDpsResource
	err := ctx.ReadResource("azure-native:devices/v20211015:IotDpsResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type iotDpsResourceState struct {
}

type IotDpsResourceState struct {
}

func (IotDpsResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotDpsResourceState)(nil)).Elem()
}

type iotDpsResourceArgs struct {
	Etag                    *string                     `pulumi:"etag"`
	Location                *string                     `pulumi:"location"`
	Properties              IotDpsPropertiesDescription `pulumi:"properties"`
	ProvisioningServiceName *string                     `pulumi:"provisioningServiceName"`
	ResourceGroupName       string                      `pulumi:"resourceGroupName"`
	Sku                     IotDpsSkuInfo               `pulumi:"sku"`
	Tags                    map[string]string           `pulumi:"tags"`
}


type IotDpsResourceArgs struct {
	Etag                    pulumi.StringPtrInput
	Location                pulumi.StringPtrInput
	Properties              IotDpsPropertiesDescriptionInput
	ProvisioningServiceName pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	Sku                     IotDpsSkuInfoInput
	Tags                    pulumi.StringMapInput
}

func (IotDpsResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotDpsResourceArgs)(nil)).Elem()
}

type IotDpsResourceInput interface {
	pulumi.Input

	ToIotDpsResourceOutput() IotDpsResourceOutput
	ToIotDpsResourceOutputWithContext(ctx context.Context) IotDpsResourceOutput
}

func (*IotDpsResource) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsResource)(nil))
}

func (i *IotDpsResource) ToIotDpsResourceOutput() IotDpsResourceOutput {
	return i.ToIotDpsResourceOutputWithContext(context.Background())
}

func (i *IotDpsResource) ToIotDpsResourceOutputWithContext(ctx context.Context) IotDpsResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDpsResourceOutput)
}

type IotDpsResourceOutput struct{ *pulumi.OutputState }

func (IotDpsResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsResource)(nil))
}

func (o IotDpsResourceOutput) ToIotDpsResourceOutput() IotDpsResourceOutput {
	return o
}

func (o IotDpsResourceOutput) ToIotDpsResourceOutputWithContext(ctx context.Context) IotDpsResourceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IotDpsResourceOutput{})
}

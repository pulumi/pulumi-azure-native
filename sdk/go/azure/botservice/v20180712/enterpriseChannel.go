


package v20180712

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnterpriseChannel struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput                    `pulumi:"etag"`
	Kind       pulumi.StringPtrOutput                    `pulumi:"kind"`
	Location   pulumi.StringPtrOutput                    `pulumi:"location"`
	Name       pulumi.StringOutput                       `pulumi:"name"`
	Properties EnterpriseChannelPropertiesResponseOutput `pulumi:"properties"`
	Sku        SkuResponsePtrOutput                      `pulumi:"sku"`
	Tags       pulumi.StringMapOutput                    `pulumi:"tags"`
	Type       pulumi.StringOutput                       `pulumi:"type"`
}


func NewEnterpriseChannel(ctx *pulumi.Context,
	name string, args *EnterpriseChannelArgs, opts ...pulumi.ResourceOption) (*EnterpriseChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:botservice/v20180712:EnterpriseChannel"),
		},
		{
			Type: pulumi.String("azure-native:botservice:EnterpriseChannel"),
		},
		{
			Type: pulumi.String("azure-nextgen:botservice:EnterpriseChannel"),
		},
	})
	opts = append(opts, aliases)
	var resource EnterpriseChannel
	err := ctx.RegisterResource("azure-native:botservice/v20180712:EnterpriseChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEnterpriseChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnterpriseChannelState, opts ...pulumi.ResourceOption) (*EnterpriseChannel, error) {
	var resource EnterpriseChannel
	err := ctx.ReadResource("azure-native:botservice/v20180712:EnterpriseChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type enterpriseChannelState struct {
}

type EnterpriseChannelState struct {
}

func (EnterpriseChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseChannelState)(nil)).Elem()
}

type enterpriseChannelArgs struct {
	Etag              *string                      `pulumi:"etag"`
	Kind              *string                      `pulumi:"kind"`
	Location          *string                      `pulumi:"location"`
	Properties        *EnterpriseChannelProperties `pulumi:"properties"`
	ResourceGroupName string                       `pulumi:"resourceGroupName"`
	ResourceName      *string                      `pulumi:"resourceName"`
	Sku               *Sku                         `pulumi:"sku"`
	Tags              map[string]string            `pulumi:"tags"`
}


type EnterpriseChannelArgs struct {
	Etag              pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Properties        EnterpriseChannelPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Sku               SkuPtrInput
	Tags              pulumi.StringMapInput
}

func (EnterpriseChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseChannelArgs)(nil)).Elem()
}

type EnterpriseChannelInput interface {
	pulumi.Input

	ToEnterpriseChannelOutput() EnterpriseChannelOutput
	ToEnterpriseChannelOutputWithContext(ctx context.Context) EnterpriseChannelOutput
}

func (*EnterpriseChannel) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseChannel)(nil))
}

func (i *EnterpriseChannel) ToEnterpriseChannelOutput() EnterpriseChannelOutput {
	return i.ToEnterpriseChannelOutputWithContext(context.Background())
}

func (i *EnterpriseChannel) ToEnterpriseChannelOutputWithContext(ctx context.Context) EnterpriseChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseChannelOutput)
}

type EnterpriseChannelOutput struct{ *pulumi.OutputState }

func (EnterpriseChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseChannel)(nil))
}

func (o EnterpriseChannelOutput) ToEnterpriseChannelOutput() EnterpriseChannelOutput {
	return o
}

func (o EnterpriseChannelOutput) ToEnterpriseChannelOutputWithContext(ctx context.Context) EnterpriseChannelOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseChannelInput)(nil)).Elem(), &EnterpriseChannel{})
	pulumi.RegisterOutputType(EnterpriseChannelOutput{})
}

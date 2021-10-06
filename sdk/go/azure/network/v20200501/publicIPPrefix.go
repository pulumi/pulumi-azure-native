


package v20200501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PublicIPPrefix struct {
	pulumi.CustomResourceState

	Etag                                pulumi.StringOutput                          `pulumi:"etag"`
	IpPrefix                            pulumi.StringOutput                          `pulumi:"ipPrefix"`
	IpTags                              IpTagResponseArrayOutput                     `pulumi:"ipTags"`
	LoadBalancerFrontendIpConfiguration SubResourceResponseOutput                    `pulumi:"loadBalancerFrontendIpConfiguration"`
	Location                            pulumi.StringPtrOutput                       `pulumi:"location"`
	Name                                pulumi.StringOutput                          `pulumi:"name"`
	PrefixLength                        pulumi.IntPtrOutput                          `pulumi:"prefixLength"`
	ProvisioningState                   pulumi.StringOutput                          `pulumi:"provisioningState"`
	PublicIPAddressVersion              pulumi.StringPtrOutput                       `pulumi:"publicIPAddressVersion"`
	PublicIPAddresses                   ReferencedPublicIpAddressResponseArrayOutput `pulumi:"publicIPAddresses"`
	ResourceGuid                        pulumi.StringOutput                          `pulumi:"resourceGuid"`
	Sku                                 PublicIPPrefixSkuResponsePtrOutput           `pulumi:"sku"`
	Tags                                pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                                pulumi.StringOutput                          `pulumi:"type"`
	Zones                               pulumi.StringArrayOutput                     `pulumi:"zones"`
}


func NewPublicIPPrefix(ctx *pulumi.Context,
	name string, args *PublicIPPrefixArgs, opts ...pulumi.ResourceOption) (*PublicIPPrefix, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180701:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20201101:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210201:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210301:PublicIPPrefix"),
		},
	})
	opts = append(opts, aliases)
	var resource PublicIPPrefix
	err := ctx.RegisterResource("azure-native:network/v20200501:PublicIPPrefix", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPublicIPPrefix(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicIPPrefixState, opts ...pulumi.ResourceOption) (*PublicIPPrefix, error) {
	var resource PublicIPPrefix
	err := ctx.ReadResource("azure-native:network/v20200501:PublicIPPrefix", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type publicIPPrefixState struct {
}

type PublicIPPrefixState struct {
}

func (PublicIPPrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIPPrefixState)(nil)).Elem()
}

type publicIPPrefixArgs struct {
	Id                     *string            `pulumi:"id"`
	IpTags                 []IpTag            `pulumi:"ipTags"`
	Location               *string            `pulumi:"location"`
	PrefixLength           *int               `pulumi:"prefixLength"`
	PublicIPAddressVersion *string            `pulumi:"publicIPAddressVersion"`
	PublicIpPrefixName     *string            `pulumi:"publicIpPrefixName"`
	ResourceGroupName      string             `pulumi:"resourceGroupName"`
	Sku                    *PublicIPPrefixSku `pulumi:"sku"`
	Tags                   map[string]string  `pulumi:"tags"`
	Zones                  []string           `pulumi:"zones"`
}


type PublicIPPrefixArgs struct {
	Id                     pulumi.StringPtrInput
	IpTags                 IpTagArrayInput
	Location               pulumi.StringPtrInput
	PrefixLength           pulumi.IntPtrInput
	PublicIPAddressVersion pulumi.StringPtrInput
	PublicIpPrefixName     pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	Sku                    PublicIPPrefixSkuPtrInput
	Tags                   pulumi.StringMapInput
	Zones                  pulumi.StringArrayInput
}

func (PublicIPPrefixArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIPPrefixArgs)(nil)).Elem()
}

type PublicIPPrefixInput interface {
	pulumi.Input

	ToPublicIPPrefixOutput() PublicIPPrefixOutput
	ToPublicIPPrefixOutputWithContext(ctx context.Context) PublicIPPrefixOutput
}

func (*PublicIPPrefix) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPPrefix)(nil))
}

func (i *PublicIPPrefix) ToPublicIPPrefixOutput() PublicIPPrefixOutput {
	return i.ToPublicIPPrefixOutputWithContext(context.Background())
}

func (i *PublicIPPrefix) ToPublicIPPrefixOutputWithContext(ctx context.Context) PublicIPPrefixOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPPrefixOutput)
}

type PublicIPPrefixOutput struct{ *pulumi.OutputState }

func (PublicIPPrefixOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPPrefix)(nil))
}

func (o PublicIPPrefixOutput) ToPublicIPPrefixOutput() PublicIPPrefixOutput {
	return o
}

func (o PublicIPPrefixOutput) ToPublicIPPrefixOutputWithContext(ctx context.Context) PublicIPPrefixOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PublicIPPrefixOutput{})
}




package v20180701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualWAN struct {
	pulumi.CustomResourceState

	DisableVpnEncryption pulumi.BoolPtrOutput           `pulumi:"disableVpnEncryption"`
	Etag                 pulumi.StringOutput            `pulumi:"etag"`
	Location             pulumi.StringOutput            `pulumi:"location"`
	Name                 pulumi.StringOutput            `pulumi:"name"`
	ProvisioningState    pulumi.StringOutput            `pulumi:"provisioningState"`
	Tags                 pulumi.StringMapOutput         `pulumi:"tags"`
	Type                 pulumi.StringOutput            `pulumi:"type"`
	VirtualHubs          SubResourceResponseArrayOutput `pulumi:"virtualHubs"`
	VpnSites             SubResourceResponseArrayOutput `pulumi:"vpnSites"`
}


func NewVirtualWAN(ctx *pulumi.Context,
	name string, args *VirtualWANArgs, opts ...pulumi.ResourceOption) (*VirtualWAN, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/v20180701:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180401:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180601:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20201101:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210201:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210301:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualWAN"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210501:VirtualWAN"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualWAN
	err := ctx.RegisterResource("azure-native:network/v20180701:VirtualWAN", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualWAN(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualWANState, opts ...pulumi.ResourceOption) (*VirtualWAN, error) {
	var resource VirtualWAN
	err := ctx.ReadResource("azure-native:network/v20180701:VirtualWAN", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualWANState struct {
}

type VirtualWANState struct {
}

func (VirtualWANState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualWANState)(nil)).Elem()
}

type virtualWANArgs struct {
	DisableVpnEncryption *bool             `pulumi:"disableVpnEncryption"`
	Id                   *string           `pulumi:"id"`
	Location             *string           `pulumi:"location"`
	ResourceGroupName    string            `pulumi:"resourceGroupName"`
	Tags                 map[string]string `pulumi:"tags"`
	VirtualWANName       *string           `pulumi:"virtualWANName"`
}


type VirtualWANArgs struct {
	DisableVpnEncryption pulumi.BoolPtrInput
	Id                   pulumi.StringPtrInput
	Location             pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	Tags                 pulumi.StringMapInput
	VirtualWANName       pulumi.StringPtrInput
}

func (VirtualWANArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualWANArgs)(nil)).Elem()
}

type VirtualWANInput interface {
	pulumi.Input

	ToVirtualWANOutput() VirtualWANOutput
	ToVirtualWANOutputWithContext(ctx context.Context) VirtualWANOutput
}

func (*VirtualWAN) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualWAN)(nil))
}

func (i *VirtualWAN) ToVirtualWANOutput() VirtualWANOutput {
	return i.ToVirtualWANOutputWithContext(context.Background())
}

func (i *VirtualWAN) ToVirtualWANOutputWithContext(ctx context.Context) VirtualWANOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualWANOutput)
}

type VirtualWANOutput struct{ *pulumi.OutputState }

func (VirtualWANOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualWAN)(nil))
}

func (o VirtualWANOutput) ToVirtualWANOutput() VirtualWANOutput {
	return o
}

func (o VirtualWANOutput) ToVirtualWANOutputWithContext(ctx context.Context) VirtualWANOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualWANInput)(nil)).Elem(), &VirtualWAN{})
	pulumi.RegisterOutputType(VirtualWANOutput{})
}

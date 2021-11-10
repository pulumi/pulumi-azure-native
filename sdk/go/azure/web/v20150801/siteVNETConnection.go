


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SiteVNETConnection struct {
	pulumi.CustomResourceState

	CertBlob       pulumi.StringPtrOutput       `pulumi:"certBlob"`
	CertThumbprint pulumi.StringPtrOutput       `pulumi:"certThumbprint"`
	DnsServers     pulumi.StringPtrOutput       `pulumi:"dnsServers"`
	Kind           pulumi.StringPtrOutput       `pulumi:"kind"`
	Location       pulumi.StringOutput          `pulumi:"location"`
	Name           pulumi.StringPtrOutput       `pulumi:"name"`
	ResyncRequired pulumi.BoolPtrOutput         `pulumi:"resyncRequired"`
	Routes         VnetRouteResponseArrayOutput `pulumi:"routes"`
	Tags           pulumi.StringMapOutput       `pulumi:"tags"`
	Type           pulumi.StringPtrOutput       `pulumi:"type"`
	VnetResourceId pulumi.StringPtrOutput       `pulumi:"vnetResourceId"`
}


func NewSiteVNETConnection(ctx *pulumi.Context,
	name string, args *SiteVNETConnectionArgs, opts ...pulumi.ResourceOption) (*SiteVNETConnection, error) {
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
			Type: pulumi.String("azure-native:web:SiteVNETConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteVNETConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteVNETConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteVNETConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteVNETConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteVNETConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteVNETConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteVNETConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteVNETConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteVNETConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteVNETConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteVNETConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteVNETConnection
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteVNETConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteVNETConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteVNETConnectionState, opts ...pulumi.ResourceOption) (*SiteVNETConnection, error) {
	var resource SiteVNETConnection
	err := ctx.ReadResource("azure-native:web/v20150801:SiteVNETConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteVNETConnectionState struct {
}

type SiteVNETConnectionState struct {
}

func (SiteVNETConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteVNETConnectionState)(nil)).Elem()
}

type siteVNETConnectionArgs struct {
	CertBlob          *string           `pulumi:"certBlob"`
	CertThumbprint    *string           `pulumi:"certThumbprint"`
	DnsServers        *string           `pulumi:"dnsServers"`
	Id                *string           `pulumi:"id"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResyncRequired    *bool             `pulumi:"resyncRequired"`
	Routes            []VnetRoute       `pulumi:"routes"`
	Tags              map[string]string `pulumi:"tags"`
	Type              *string           `pulumi:"type"`
	VnetName          *string           `pulumi:"vnetName"`
	VnetResourceId    *string           `pulumi:"vnetResourceId"`
}


type SiteVNETConnectionArgs struct {
	CertBlob          pulumi.StringPtrInput
	CertThumbprint    pulumi.StringPtrInput
	DnsServers        pulumi.StringPtrInput
	Id                pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ResyncRequired    pulumi.BoolPtrInput
	Routes            VnetRouteArrayInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringPtrInput
	VnetName          pulumi.StringPtrInput
	VnetResourceId    pulumi.StringPtrInput
}

func (SiteVNETConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteVNETConnectionArgs)(nil)).Elem()
}

type SiteVNETConnectionInput interface {
	pulumi.Input

	ToSiteVNETConnectionOutput() SiteVNETConnectionOutput
	ToSiteVNETConnectionOutputWithContext(ctx context.Context) SiteVNETConnectionOutput
}

func (*SiteVNETConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteVNETConnection)(nil))
}

func (i *SiteVNETConnection) ToSiteVNETConnectionOutput() SiteVNETConnectionOutput {
	return i.ToSiteVNETConnectionOutputWithContext(context.Background())
}

func (i *SiteVNETConnection) ToSiteVNETConnectionOutputWithContext(ctx context.Context) SiteVNETConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteVNETConnectionOutput)
}

type SiteVNETConnectionOutput struct{ *pulumi.OutputState }

func (SiteVNETConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteVNETConnection)(nil))
}

func (o SiteVNETConnectionOutput) ToSiteVNETConnectionOutput() SiteVNETConnectionOutput {
	return o
}

func (o SiteVNETConnectionOutput) ToSiteVNETConnectionOutputWithContext(ctx context.Context) SiteVNETConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SiteVNETConnectionOutput{})
}

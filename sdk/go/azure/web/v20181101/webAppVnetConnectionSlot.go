


package v20181101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppVnetConnectionSlot struct {
	pulumi.CustomResourceState

	CertBlob       pulumi.StringPtrOutput       `pulumi:"certBlob"`
	CertThumbprint pulumi.StringOutput          `pulumi:"certThumbprint"`
	DnsServers     pulumi.StringPtrOutput       `pulumi:"dnsServers"`
	IsSwift        pulumi.BoolPtrOutput         `pulumi:"isSwift"`
	Kind           pulumi.StringPtrOutput       `pulumi:"kind"`
	Name           pulumi.StringOutput          `pulumi:"name"`
	ResyncRequired pulumi.BoolOutput            `pulumi:"resyncRequired"`
	Routes         VnetRouteResponseArrayOutput `pulumi:"routes"`
	Type           pulumi.StringOutput          `pulumi:"type"`
	VnetResourceId pulumi.StringPtrOutput       `pulumi:"vnetResourceId"`
}


func NewWebAppVnetConnectionSlot(ctx *pulumi.Context,
	name string, args *WebAppVnetConnectionSlotArgs, opts ...pulumi.ResourceOption) (*WebAppVnetConnectionSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppVnetConnectionSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppVnetConnectionSlot
	err := ctx.RegisterResource("azure-native:web/v20181101:WebAppVnetConnectionSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppVnetConnectionSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppVnetConnectionSlotState, opts ...pulumi.ResourceOption) (*WebAppVnetConnectionSlot, error) {
	var resource WebAppVnetConnectionSlot
	err := ctx.ReadResource("azure-native:web/v20181101:WebAppVnetConnectionSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppVnetConnectionSlotState struct {
}

type WebAppVnetConnectionSlotState struct {
}

func (WebAppVnetConnectionSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppVnetConnectionSlotState)(nil)).Elem()
}

type webAppVnetConnectionSlotArgs struct {
	CertBlob          *string `pulumi:"certBlob"`
	DnsServers        *string `pulumi:"dnsServers"`
	IsSwift           *bool   `pulumi:"isSwift"`
	Kind              *string `pulumi:"kind"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	Slot              string  `pulumi:"slot"`
	VnetName          *string `pulumi:"vnetName"`
	VnetResourceId    *string `pulumi:"vnetResourceId"`
}


type WebAppVnetConnectionSlotArgs struct {
	CertBlob          pulumi.StringPtrInput
	DnsServers        pulumi.StringPtrInput
	IsSwift           pulumi.BoolPtrInput
	Kind              pulumi.StringPtrInput
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Slot              pulumi.StringInput
	VnetName          pulumi.StringPtrInput
	VnetResourceId    pulumi.StringPtrInput
}

func (WebAppVnetConnectionSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppVnetConnectionSlotArgs)(nil)).Elem()
}

type WebAppVnetConnectionSlotInput interface {
	pulumi.Input

	ToWebAppVnetConnectionSlotOutput() WebAppVnetConnectionSlotOutput
	ToWebAppVnetConnectionSlotOutputWithContext(ctx context.Context) WebAppVnetConnectionSlotOutput
}

func (*WebAppVnetConnectionSlot) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppVnetConnectionSlot)(nil))
}

func (i *WebAppVnetConnectionSlot) ToWebAppVnetConnectionSlotOutput() WebAppVnetConnectionSlotOutput {
	return i.ToWebAppVnetConnectionSlotOutputWithContext(context.Background())
}

func (i *WebAppVnetConnectionSlot) ToWebAppVnetConnectionSlotOutputWithContext(ctx context.Context) WebAppVnetConnectionSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppVnetConnectionSlotOutput)
}

type WebAppVnetConnectionSlotOutput struct{ *pulumi.OutputState }

func (WebAppVnetConnectionSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppVnetConnectionSlot)(nil))
}

func (o WebAppVnetConnectionSlotOutput) ToWebAppVnetConnectionSlotOutput() WebAppVnetConnectionSlotOutput {
	return o
}

func (o WebAppVnetConnectionSlotOutput) ToWebAppVnetConnectionSlotOutputWithContext(ctx context.Context) WebAppVnetConnectionSlotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppVnetConnectionSlotOutput{})
}




package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OpenIdConnectProvider struct {
	pulumi.CustomResourceState

	ClientId         pulumi.StringOutput    `pulumi:"clientId"`
	ClientSecret     pulumi.StringPtrOutput `pulumi:"clientSecret"`
	Description      pulumi.StringPtrOutput `pulumi:"description"`
	DisplayName      pulumi.StringOutput    `pulumi:"displayName"`
	MetadataEndpoint pulumi.StringOutput    `pulumi:"metadataEndpoint"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	Type             pulumi.StringOutput    `pulumi:"type"`
}


func NewOpenIdConnectProvider(ctx *pulumi.Context,
	name string, args *OpenIdConnectProviderArgs, opts ...pulumi.ResourceOption) (*OpenIdConnectProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.MetadataEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'MetadataEndpoint'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210401preview:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20160707:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20161010:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20170301:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180101:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180601preview:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20190101:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20201201:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210101preview:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210801:OpenIdConnectProvider"),
		},
	})
	opts = append(opts, aliases)
	var resource OpenIdConnectProvider
	err := ctx.RegisterResource("azure-native:apimanagement/v20210401preview:OpenIdConnectProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOpenIdConnectProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OpenIdConnectProviderState, opts ...pulumi.ResourceOption) (*OpenIdConnectProvider, error) {
	var resource OpenIdConnectProvider
	err := ctx.ReadResource("azure-native:apimanagement/v20210401preview:OpenIdConnectProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type openIdConnectProviderState struct {
}

type OpenIdConnectProviderState struct {
}

func (OpenIdConnectProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*openIdConnectProviderState)(nil)).Elem()
}

type openIdConnectProviderArgs struct {
	ClientId          string  `pulumi:"clientId"`
	ClientSecret      *string `pulumi:"clientSecret"`
	Description       *string `pulumi:"description"`
	DisplayName       string  `pulumi:"displayName"`
	MetadataEndpoint  string  `pulumi:"metadataEndpoint"`
	Opid              *string `pulumi:"opid"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type OpenIdConnectProviderArgs struct {
	ClientId          pulumi.StringInput
	ClientSecret      pulumi.StringPtrInput
	Description       pulumi.StringPtrInput
	DisplayName       pulumi.StringInput
	MetadataEndpoint  pulumi.StringInput
	Opid              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (OpenIdConnectProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*openIdConnectProviderArgs)(nil)).Elem()
}

type OpenIdConnectProviderInput interface {
	pulumi.Input

	ToOpenIdConnectProviderOutput() OpenIdConnectProviderOutput
	ToOpenIdConnectProviderOutputWithContext(ctx context.Context) OpenIdConnectProviderOutput
}

func (*OpenIdConnectProvider) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectProvider)(nil))
}

func (i *OpenIdConnectProvider) ToOpenIdConnectProviderOutput() OpenIdConnectProviderOutput {
	return i.ToOpenIdConnectProviderOutputWithContext(context.Background())
}

func (i *OpenIdConnectProvider) ToOpenIdConnectProviderOutputWithContext(ctx context.Context) OpenIdConnectProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectProviderOutput)
}

type OpenIdConnectProviderOutput struct{ *pulumi.OutputState }

func (OpenIdConnectProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectProvider)(nil))
}

func (o OpenIdConnectProviderOutput) ToOpenIdConnectProviderOutput() OpenIdConnectProviderOutput {
	return o
}

func (o OpenIdConnectProviderOutput) ToOpenIdConnectProviderOutputWithContext(ctx context.Context) OpenIdConnectProviderOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OpenIdConnectProviderInput)(nil)).Elem(), &OpenIdConnectProvider{})
	pulumi.RegisterOutputType(OpenIdConnectProviderOutput{})
}

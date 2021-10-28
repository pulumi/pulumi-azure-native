


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerDnsAlias struct {
	pulumi.CustomResourceState

	AzureDnsRecord pulumi.StringOutput `pulumi:"azureDnsRecord"`
	Name           pulumi.StringOutput `pulumi:"name"`
	Type           pulumi.StringOutput `pulumi:"type"`
}


func NewServerDnsAlias(ctx *pulumi.Context,
	name string, args *ServerDnsAliasArgs, opts ...pulumi.ResourceOption) (*ServerDnsAlias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:sql/v20210501preview:ServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-native:sql:ServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql:ServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:ServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20170301preview:ServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200202preview:ServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200801preview:ServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20201101preview:ServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20210201preview:ServerDnsAlias"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerDnsAlias
	err := ctx.RegisterResource("azure-native:sql/v20210501preview:ServerDnsAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerDnsAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerDnsAliasState, opts ...pulumi.ResourceOption) (*ServerDnsAlias, error) {
	var resource ServerDnsAlias
	err := ctx.ReadResource("azure-native:sql/v20210501preview:ServerDnsAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serverDnsAliasState struct {
}

type ServerDnsAliasState struct {
}

func (ServerDnsAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverDnsAliasState)(nil)).Elem()
}

type serverDnsAliasArgs struct {
	DnsAliasName      *string `pulumi:"dnsAliasName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServerName        string  `pulumi:"serverName"`
}


type ServerDnsAliasArgs struct {
	DnsAliasName      pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
}

func (ServerDnsAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverDnsAliasArgs)(nil)).Elem()
}

type ServerDnsAliasInput interface {
	pulumi.Input

	ToServerDnsAliasOutput() ServerDnsAliasOutput
	ToServerDnsAliasOutputWithContext(ctx context.Context) ServerDnsAliasOutput
}

func (*ServerDnsAlias) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerDnsAlias)(nil))
}

func (i *ServerDnsAlias) ToServerDnsAliasOutput() ServerDnsAliasOutput {
	return i.ToServerDnsAliasOutputWithContext(context.Background())
}

func (i *ServerDnsAlias) ToServerDnsAliasOutputWithContext(ctx context.Context) ServerDnsAliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerDnsAliasOutput)
}

type ServerDnsAliasOutput struct{ *pulumi.OutputState }

func (ServerDnsAliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerDnsAlias)(nil))
}

func (o ServerDnsAliasOutput) ToServerDnsAliasOutput() ServerDnsAliasOutput {
	return o
}

func (o ServerDnsAliasOutput) ToServerDnsAliasOutputWithContext(ctx context.Context) ServerDnsAliasOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServerDnsAliasOutput{})
}

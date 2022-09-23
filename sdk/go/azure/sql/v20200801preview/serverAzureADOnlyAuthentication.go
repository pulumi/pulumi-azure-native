


package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerAzureADOnlyAuthentication struct {
	pulumi.CustomResourceState

	AzureADOnlyAuthentication pulumi.BoolOutput   `pulumi:"azureADOnlyAuthentication"`
	Name                      pulumi.StringOutput `pulumi:"name"`
	Type                      pulumi.StringOutput `pulumi:"type"`
}


func NewServerAzureADOnlyAuthentication(ctx *pulumi.Context,
	name string, args *ServerAzureADOnlyAuthenticationArgs, opts ...pulumi.ResourceOption) (*ServerAzureADOnlyAuthentication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AzureADOnlyAuthentication == nil {
		return nil, errors.New("invalid value for required argument 'AzureADOnlyAuthentication'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ServerAzureADOnlyAuthentication"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ServerAzureADOnlyAuthentication"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerAzureADOnlyAuthentication
	err := ctx.RegisterResource("azure-native:sql/v20200801preview:ServerAzureADOnlyAuthentication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerAzureADOnlyAuthentication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerAzureADOnlyAuthenticationState, opts ...pulumi.ResourceOption) (*ServerAzureADOnlyAuthentication, error) {
	var resource ServerAzureADOnlyAuthentication
	err := ctx.ReadResource("azure-native:sql/v20200801preview:ServerAzureADOnlyAuthentication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serverAzureADOnlyAuthenticationState struct {
}

type ServerAzureADOnlyAuthenticationState struct {
}

func (ServerAzureADOnlyAuthenticationState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverAzureADOnlyAuthenticationState)(nil)).Elem()
}

type serverAzureADOnlyAuthenticationArgs struct {
	AuthenticationName        *string `pulumi:"authenticationName"`
	AzureADOnlyAuthentication bool    `pulumi:"azureADOnlyAuthentication"`
	ResourceGroupName         string  `pulumi:"resourceGroupName"`
	ServerName                string  `pulumi:"serverName"`
}


type ServerAzureADOnlyAuthenticationArgs struct {
	AuthenticationName        pulumi.StringPtrInput
	AzureADOnlyAuthentication pulumi.BoolInput
	ResourceGroupName         pulumi.StringInput
	ServerName                pulumi.StringInput
}

func (ServerAzureADOnlyAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverAzureADOnlyAuthenticationArgs)(nil)).Elem()
}

type ServerAzureADOnlyAuthenticationInput interface {
	pulumi.Input

	ToServerAzureADOnlyAuthenticationOutput() ServerAzureADOnlyAuthenticationOutput
	ToServerAzureADOnlyAuthenticationOutputWithContext(ctx context.Context) ServerAzureADOnlyAuthenticationOutput
}

func (*ServerAzureADOnlyAuthentication) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAzureADOnlyAuthentication)(nil)).Elem()
}

func (i *ServerAzureADOnlyAuthentication) ToServerAzureADOnlyAuthenticationOutput() ServerAzureADOnlyAuthenticationOutput {
	return i.ToServerAzureADOnlyAuthenticationOutputWithContext(context.Background())
}

func (i *ServerAzureADOnlyAuthentication) ToServerAzureADOnlyAuthenticationOutputWithContext(ctx context.Context) ServerAzureADOnlyAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAzureADOnlyAuthenticationOutput)
}

type ServerAzureADOnlyAuthenticationOutput struct{ *pulumi.OutputState }

func (ServerAzureADOnlyAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAzureADOnlyAuthentication)(nil)).Elem()
}

func (o ServerAzureADOnlyAuthenticationOutput) ToServerAzureADOnlyAuthenticationOutput() ServerAzureADOnlyAuthenticationOutput {
	return o
}

func (o ServerAzureADOnlyAuthenticationOutput) ToServerAzureADOnlyAuthenticationOutputWithContext(ctx context.Context) ServerAzureADOnlyAuthenticationOutput {
	return o
}

func (o ServerAzureADOnlyAuthenticationOutput) AzureADOnlyAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServerAzureADOnlyAuthentication) pulumi.BoolOutput { return v.AzureADOnlyAuthentication }).(pulumi.BoolOutput)
}

func (o ServerAzureADOnlyAuthenticationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAzureADOnlyAuthentication) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServerAzureADOnlyAuthenticationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAzureADOnlyAuthentication) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerAzureADOnlyAuthenticationOutput{})
}

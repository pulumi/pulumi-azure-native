


package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerAzureADAdministrator struct {
	pulumi.CustomResourceState

	AdministratorType         pulumi.StringOutput    `pulumi:"administratorType"`
	AzureADOnlyAuthentication pulumi.BoolOutput      `pulumi:"azureADOnlyAuthentication"`
	Login                     pulumi.StringOutput    `pulumi:"login"`
	Name                      pulumi.StringOutput    `pulumi:"name"`
	Sid                       pulumi.StringOutput    `pulumi:"sid"`
	TenantId                  pulumi.StringPtrOutput `pulumi:"tenantId"`
	Type                      pulumi.StringOutput    `pulumi:"type"`
}


func NewServerAzureADAdministrator(ctx *pulumi.Context,
	name string, args *ServerAzureADAdministratorArgs, opts ...pulumi.ResourceOption) (*ServerAzureADAdministrator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdministratorType == nil {
		return nil, errors.New("invalid value for required argument 'AdministratorType'")
	}
	if args.Login == nil {
		return nil, errors.New("invalid value for required argument 'Login'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.Sid == nil {
		return nil, errors.New("invalid value for required argument 'Sid'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:sql/v20190601preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20140401:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20140401:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20180601preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20180601preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200202preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200801preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20201101preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20210201preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20210501preview:ServerAzureADAdministrator"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerAzureADAdministrator
	err := ctx.RegisterResource("azure-native:sql/v20190601preview:ServerAzureADAdministrator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerAzureADAdministrator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerAzureADAdministratorState, opts ...pulumi.ResourceOption) (*ServerAzureADAdministrator, error) {
	var resource ServerAzureADAdministrator
	err := ctx.ReadResource("azure-native:sql/v20190601preview:ServerAzureADAdministrator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serverAzureADAdministratorState struct {
}

type ServerAzureADAdministratorState struct {
}

func (ServerAzureADAdministratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverAzureADAdministratorState)(nil)).Elem()
}

type serverAzureADAdministratorArgs struct {
	AdministratorName *string `pulumi:"administratorName"`
	AdministratorType string  `pulumi:"administratorType"`
	Login             string  `pulumi:"login"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServerName        string  `pulumi:"serverName"`
	Sid               string  `pulumi:"sid"`
	TenantId          *string `pulumi:"tenantId"`
}


type ServerAzureADAdministratorArgs struct {
	AdministratorName pulumi.StringPtrInput
	AdministratorType pulumi.StringInput
	Login             pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
	Sid               pulumi.StringInput
	TenantId          pulumi.StringPtrInput
}

func (ServerAzureADAdministratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverAzureADAdministratorArgs)(nil)).Elem()
}

type ServerAzureADAdministratorInput interface {
	pulumi.Input

	ToServerAzureADAdministratorOutput() ServerAzureADAdministratorOutput
	ToServerAzureADAdministratorOutputWithContext(ctx context.Context) ServerAzureADAdministratorOutput
}

func (*ServerAzureADAdministrator) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerAzureADAdministrator)(nil))
}

func (i *ServerAzureADAdministrator) ToServerAzureADAdministratorOutput() ServerAzureADAdministratorOutput {
	return i.ToServerAzureADAdministratorOutputWithContext(context.Background())
}

func (i *ServerAzureADAdministrator) ToServerAzureADAdministratorOutputWithContext(ctx context.Context) ServerAzureADAdministratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAzureADAdministratorOutput)
}

type ServerAzureADAdministratorOutput struct{ *pulumi.OutputState }

func (ServerAzureADAdministratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerAzureADAdministrator)(nil))
}

func (o ServerAzureADAdministratorOutput) ToServerAzureADAdministratorOutput() ServerAzureADAdministratorOutput {
	return o
}

func (o ServerAzureADAdministratorOutput) ToServerAzureADAdministratorOutputWithContext(ctx context.Context) ServerAzureADAdministratorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServerAzureADAdministratorOutput{})
}




package v20171201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerAdministrator struct {
	pulumi.CustomResourceState

	AdministratorType pulumi.StringOutput `pulumi:"administratorType"`
	Login             pulumi.StringOutput `pulumi:"login"`
	Name              pulumi.StringOutput `pulumi:"name"`
	Sid               pulumi.StringOutput `pulumi:"sid"`
	TenantId          pulumi.StringOutput `pulumi:"tenantId"`
	Type              pulumi.StringOutput `pulumi:"type"`
}


func NewServerAdministrator(ctx *pulumi.Context,
	name string, args *ServerAdministratorArgs, opts ...pulumi.ResourceOption) (*ServerAdministrator, error) {
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
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:dbformysql/v20171201:ServerAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql:ServerAdministrator"),
		},
		{
			Type: pulumi.String("azure-nextgen:dbformysql:ServerAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20171201preview:ServerAdministrator"),
		},
		{
			Type: pulumi.String("azure-nextgen:dbformysql/v20171201preview:ServerAdministrator"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerAdministrator
	err := ctx.RegisterResource("azure-native:dbformysql/v20171201:ServerAdministrator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerAdministrator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerAdministratorState, opts ...pulumi.ResourceOption) (*ServerAdministrator, error) {
	var resource ServerAdministrator
	err := ctx.ReadResource("azure-native:dbformysql/v20171201:ServerAdministrator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serverAdministratorState struct {
}

type ServerAdministratorState struct {
}

func (ServerAdministratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverAdministratorState)(nil)).Elem()
}

type serverAdministratorArgs struct {
	AdministratorType string `pulumi:"administratorType"`
	Login             string `pulumi:"login"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
	Sid               string `pulumi:"sid"`
	TenantId          string `pulumi:"tenantId"`
}


type ServerAdministratorArgs struct {
	AdministratorType pulumi.StringInput
	Login             pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
	Sid               pulumi.StringInput
	TenantId          pulumi.StringInput
}

func (ServerAdministratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverAdministratorArgs)(nil)).Elem()
}

type ServerAdministratorInput interface {
	pulumi.Input

	ToServerAdministratorOutput() ServerAdministratorOutput
	ToServerAdministratorOutputWithContext(ctx context.Context) ServerAdministratorOutput
}

func (*ServerAdministrator) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerAdministrator)(nil))
}

func (i *ServerAdministrator) ToServerAdministratorOutput() ServerAdministratorOutput {
	return i.ToServerAdministratorOutputWithContext(context.Background())
}

func (i *ServerAdministrator) ToServerAdministratorOutputWithContext(ctx context.Context) ServerAdministratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAdministratorOutput)
}

type ServerAdministratorOutput struct{ *pulumi.OutputState }

func (ServerAdministratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerAdministrator)(nil))
}

func (o ServerAdministratorOutput) ToServerAdministratorOutput() ServerAdministratorOutput {
	return o
}

func (o ServerAdministratorOutput) ToServerAdministratorOutputWithContext(ctx context.Context) ServerAdministratorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServerAdministratorOutput{})
}




package v20170801beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerDetails struct {
	pulumi.CustomResourceState

	AsAdministrators        ServerAdministratorsResponsePtrOutput `pulumi:"asAdministrators"`
	BackupBlobContainerUri  pulumi.StringPtrOutput                `pulumi:"backupBlobContainerUri"`
	GatewayDetails          GatewayDetailsResponsePtrOutput       `pulumi:"gatewayDetails"`
	IpV4FirewallSettings    IPv4FirewallSettingsResponsePtrOutput `pulumi:"ipV4FirewallSettings"`
	Location                pulumi.StringOutput                   `pulumi:"location"`
	ManagedMode             pulumi.IntPtrOutput                   `pulumi:"managedMode"`
	Name                    pulumi.StringOutput                   `pulumi:"name"`
	ProvisioningState       pulumi.StringOutput                   `pulumi:"provisioningState"`
	QuerypoolConnectionMode pulumi.StringPtrOutput                `pulumi:"querypoolConnectionMode"`
	ServerFullName          pulumi.StringOutput                   `pulumi:"serverFullName"`
	ServerMonitorMode       pulumi.IntPtrOutput                   `pulumi:"serverMonitorMode"`
	Sku                     ResourceSkuResponseOutput             `pulumi:"sku"`
	State                   pulumi.StringOutput                   `pulumi:"state"`
	Tags                    pulumi.StringMapOutput                `pulumi:"tags"`
	Type                    pulumi.StringOutput                   `pulumi:"type"`
}


func NewServerDetails(ctx *pulumi.Context,
	name string, args *ServerDetailsArgs, opts ...pulumi.ResourceOption) (*ServerDetails, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if isZero(args.ManagedMode) {
		args.ManagedMode = pulumi.IntPtr(1)
	}
	if isZero(args.QuerypoolConnectionMode) {
		args.QuerypoolConnectionMode = ConnectionMode("All")
	}
	if isZero(args.ServerMonitorMode) {
		args.ServerMonitorMode = pulumi.IntPtr(1)
	}
	skuApplier := func(v ResourceSku) *ResourceSku { return v.Defaults() }
	args.Sku = args.Sku.ToResourceSkuOutput().ApplyT(skuApplier).(ResourceSkuPtrOutput).Elem()
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:analysisservices:ServerDetails"),
		},
		{
			Type: pulumi.String("azure-native:analysisservices/v20160516:ServerDetails"),
		},
		{
			Type: pulumi.String("azure-native:analysisservices/v20170714:ServerDetails"),
		},
		{
			Type: pulumi.String("azure-native:analysisservices/v20170801:ServerDetails"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerDetails
	err := ctx.RegisterResource("azure-native:analysisservices/v20170801beta:ServerDetails", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerDetails(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerDetailsState, opts ...pulumi.ResourceOption) (*ServerDetails, error) {
	var resource ServerDetails
	err := ctx.ReadResource("azure-native:analysisservices/v20170801beta:ServerDetails", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serverDetailsState struct {
}

type ServerDetailsState struct {
}

func (ServerDetailsState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverDetailsState)(nil)).Elem()
}

type serverDetailsArgs struct {
	AsAdministrators        *ServerAdministrators `pulumi:"asAdministrators"`
	BackupBlobContainerUri  *string               `pulumi:"backupBlobContainerUri"`
	GatewayDetails          *GatewayDetails       `pulumi:"gatewayDetails"`
	IpV4FirewallSettings    *IPv4FirewallSettings `pulumi:"ipV4FirewallSettings"`
	Location                *string               `pulumi:"location"`
	ManagedMode             *int                  `pulumi:"managedMode"`
	QuerypoolConnectionMode *ConnectionMode       `pulumi:"querypoolConnectionMode"`
	ResourceGroupName       string                `pulumi:"resourceGroupName"`
	ServerMonitorMode       *int                  `pulumi:"serverMonitorMode"`
	ServerName              *string               `pulumi:"serverName"`
	Sku                     ResourceSku           `pulumi:"sku"`
	Tags                    map[string]string     `pulumi:"tags"`
}


type ServerDetailsArgs struct {
	AsAdministrators        ServerAdministratorsPtrInput
	BackupBlobContainerUri  pulumi.StringPtrInput
	GatewayDetails          GatewayDetailsPtrInput
	IpV4FirewallSettings    IPv4FirewallSettingsPtrInput
	Location                pulumi.StringPtrInput
	ManagedMode             pulumi.IntPtrInput
	QuerypoolConnectionMode ConnectionModePtrInput
	ResourceGroupName       pulumi.StringInput
	ServerMonitorMode       pulumi.IntPtrInput
	ServerName              pulumi.StringPtrInput
	Sku                     ResourceSkuInput
	Tags                    pulumi.StringMapInput
}

func (ServerDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverDetailsArgs)(nil)).Elem()
}

type ServerDetailsInput interface {
	pulumi.Input

	ToServerDetailsOutput() ServerDetailsOutput
	ToServerDetailsOutputWithContext(ctx context.Context) ServerDetailsOutput
}

func (*ServerDetails) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerDetails)(nil)).Elem()
}

func (i *ServerDetails) ToServerDetailsOutput() ServerDetailsOutput {
	return i.ToServerDetailsOutputWithContext(context.Background())
}

func (i *ServerDetails) ToServerDetailsOutputWithContext(ctx context.Context) ServerDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerDetailsOutput)
}

type ServerDetailsOutput struct{ *pulumi.OutputState }

func (ServerDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerDetails)(nil)).Elem()
}

func (o ServerDetailsOutput) ToServerDetailsOutput() ServerDetailsOutput {
	return o
}

func (o ServerDetailsOutput) ToServerDetailsOutputWithContext(ctx context.Context) ServerDetailsOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServerDetailsOutput{})
}

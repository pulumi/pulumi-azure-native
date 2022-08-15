


package v20170801

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
	args.Sku = args.Sku.ToResourceSkuOutput().ApplyT(func(v ResourceSku) ResourceSku { return *v.Defaults() }).(ResourceSkuOutput)
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
			Type: pulumi.String("azure-native:analysisservices/v20170801beta:ServerDetails"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerDetails
	err := ctx.RegisterResource("azure-native:analysisservices/v20170801:ServerDetails", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerDetails(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerDetailsState, opts ...pulumi.ResourceOption) (*ServerDetails, error) {
	var resource ServerDetails
	err := ctx.ReadResource("azure-native:analysisservices/v20170801:ServerDetails", name, id, state, &resource, opts...)
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

func (o ServerDetailsOutput) AsAdministrators() ServerAdministratorsResponsePtrOutput {
	return o.ApplyT(func(v *ServerDetails) ServerAdministratorsResponsePtrOutput { return v.AsAdministrators }).(ServerAdministratorsResponsePtrOutput)
}

func (o ServerDetailsOutput) BackupBlobContainerUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.StringPtrOutput { return v.BackupBlobContainerUri }).(pulumi.StringPtrOutput)
}

func (o ServerDetailsOutput) GatewayDetails() GatewayDetailsResponsePtrOutput {
	return o.ApplyT(func(v *ServerDetails) GatewayDetailsResponsePtrOutput { return v.GatewayDetails }).(GatewayDetailsResponsePtrOutput)
}

func (o ServerDetailsOutput) IpV4FirewallSettings() IPv4FirewallSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ServerDetails) IPv4FirewallSettingsResponsePtrOutput { return v.IpV4FirewallSettings }).(IPv4FirewallSettingsResponsePtrOutput)
}

func (o ServerDetailsOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ServerDetailsOutput) ManagedMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.IntPtrOutput { return v.ManagedMode }).(pulumi.IntPtrOutput)
}

func (o ServerDetailsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServerDetailsOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ServerDetailsOutput) QuerypoolConnectionMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.StringPtrOutput { return v.QuerypoolConnectionMode }).(pulumi.StringPtrOutput)
}

func (o ServerDetailsOutput) ServerFullName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.StringOutput { return v.ServerFullName }).(pulumi.StringOutput)
}

func (o ServerDetailsOutput) ServerMonitorMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.IntPtrOutput { return v.ServerMonitorMode }).(pulumi.IntPtrOutput)
}

func (o ServerDetailsOutput) Sku() ResourceSkuResponseOutput {
	return o.ApplyT(func(v *ServerDetails) ResourceSkuResponseOutput { return v.Sku }).(ResourceSkuResponseOutput)
}

func (o ServerDetailsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o ServerDetailsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ServerDetailsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerDetails) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerDetailsOutput{})
}

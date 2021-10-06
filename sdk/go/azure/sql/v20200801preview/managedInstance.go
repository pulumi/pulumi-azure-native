


package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedInstance struct {
	pulumi.CustomResourceState

	AdministratorLogin         pulumi.StringPtrOutput                        `pulumi:"administratorLogin"`
	Collation                  pulumi.StringPtrOutput                        `pulumi:"collation"`
	DnsZone                    pulumi.StringOutput                           `pulumi:"dnsZone"`
	FullyQualifiedDomainName   pulumi.StringOutput                           `pulumi:"fullyQualifiedDomainName"`
	Identity                   ResourceIdentityResponsePtrOutput             `pulumi:"identity"`
	InstancePoolId             pulumi.StringPtrOutput                        `pulumi:"instancePoolId"`
	LicenseType                pulumi.StringPtrOutput                        `pulumi:"licenseType"`
	Location                   pulumi.StringOutput                           `pulumi:"location"`
	MaintenanceConfigurationId pulumi.StringPtrOutput                        `pulumi:"maintenanceConfigurationId"`
	MinimalTlsVersion          pulumi.StringPtrOutput                        `pulumi:"minimalTlsVersion"`
	Name                       pulumi.StringOutput                           `pulumi:"name"`
	PrivateEndpointConnections ManagedInstancePecPropertyResponseArrayOutput `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringOutput                           `pulumi:"provisioningState"`
	ProxyOverride              pulumi.StringPtrOutput                        `pulumi:"proxyOverride"`
	PublicDataEndpointEnabled  pulumi.BoolPtrOutput                          `pulumi:"publicDataEndpointEnabled"`
	Sku                        SkuResponsePtrOutput                          `pulumi:"sku"`
	State                      pulumi.StringOutput                           `pulumi:"state"`
	StorageAccountType         pulumi.StringPtrOutput                        `pulumi:"storageAccountType"`
	StorageSizeInGB            pulumi.IntPtrOutput                           `pulumi:"storageSizeInGB"`
	SubnetId                   pulumi.StringPtrOutput                        `pulumi:"subnetId"`
	Tags                       pulumi.StringMapOutput                        `pulumi:"tags"`
	TimezoneId                 pulumi.StringPtrOutput                        `pulumi:"timezoneId"`
	Type                       pulumi.StringOutput                           `pulumi:"type"`
	VCores                     pulumi.IntPtrOutput                           `pulumi:"vCores"`
	ZoneRedundant              pulumi.BoolPtrOutput                          `pulumi:"zoneRedundant"`
}


func NewManagedInstance(ctx *pulumi.Context,
	name string, args *ManagedInstanceArgs, opts ...pulumi.ResourceOption) (*ManagedInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:sql/v20200801preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20150501preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20180601preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20180601preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200202preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20201101preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20210201preview:ManagedInstance"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedInstance
	err := ctx.RegisterResource("azure-native:sql/v20200801preview:ManagedInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedInstanceState, opts ...pulumi.ResourceOption) (*ManagedInstance, error) {
	var resource ManagedInstance
	err := ctx.ReadResource("azure-native:sql/v20200801preview:ManagedInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedInstanceState struct {
}

type ManagedInstanceState struct {
}

func (ManagedInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceState)(nil)).Elem()
}

type managedInstanceArgs struct {
	AdministratorLogin         *string           `pulumi:"administratorLogin"`
	AdministratorLoginPassword *string           `pulumi:"administratorLoginPassword"`
	Collation                  *string           `pulumi:"collation"`
	DnsZonePartner             *string           `pulumi:"dnsZonePartner"`
	Identity                   *ResourceIdentity `pulumi:"identity"`
	InstancePoolId             *string           `pulumi:"instancePoolId"`
	LicenseType                *string           `pulumi:"licenseType"`
	Location                   *string           `pulumi:"location"`
	MaintenanceConfigurationId *string           `pulumi:"maintenanceConfigurationId"`
	ManagedInstanceCreateMode  *string           `pulumi:"managedInstanceCreateMode"`
	ManagedInstanceName        *string           `pulumi:"managedInstanceName"`
	MinimalTlsVersion          *string           `pulumi:"minimalTlsVersion"`
	ProxyOverride              *string           `pulumi:"proxyOverride"`
	PublicDataEndpointEnabled  *bool             `pulumi:"publicDataEndpointEnabled"`
	ResourceGroupName          string            `pulumi:"resourceGroupName"`
	RestorePointInTime         *string           `pulumi:"restorePointInTime"`
	Sku                        *Sku              `pulumi:"sku"`
	SourceManagedInstanceId    *string           `pulumi:"sourceManagedInstanceId"`
	StorageAccountType         *string           `pulumi:"storageAccountType"`
	StorageSizeInGB            *int              `pulumi:"storageSizeInGB"`
	SubnetId                   *string           `pulumi:"subnetId"`
	Tags                       map[string]string `pulumi:"tags"`
	TimezoneId                 *string           `pulumi:"timezoneId"`
	VCores                     *int              `pulumi:"vCores"`
	ZoneRedundant              *bool             `pulumi:"zoneRedundant"`
}


type ManagedInstanceArgs struct {
	AdministratorLogin         pulumi.StringPtrInput
	AdministratorLoginPassword pulumi.StringPtrInput
	Collation                  pulumi.StringPtrInput
	DnsZonePartner             pulumi.StringPtrInput
	Identity                   ResourceIdentityPtrInput
	InstancePoolId             pulumi.StringPtrInput
	LicenseType                pulumi.StringPtrInput
	Location                   pulumi.StringPtrInput
	MaintenanceConfigurationId pulumi.StringPtrInput
	ManagedInstanceCreateMode  pulumi.StringPtrInput
	ManagedInstanceName        pulumi.StringPtrInput
	MinimalTlsVersion          pulumi.StringPtrInput
	ProxyOverride              pulumi.StringPtrInput
	PublicDataEndpointEnabled  pulumi.BoolPtrInput
	ResourceGroupName          pulumi.StringInput
	RestorePointInTime         pulumi.StringPtrInput
	Sku                        SkuPtrInput
	SourceManagedInstanceId    pulumi.StringPtrInput
	StorageAccountType         pulumi.StringPtrInput
	StorageSizeInGB            pulumi.IntPtrInput
	SubnetId                   pulumi.StringPtrInput
	Tags                       pulumi.StringMapInput
	TimezoneId                 pulumi.StringPtrInput
	VCores                     pulumi.IntPtrInput
	ZoneRedundant              pulumi.BoolPtrInput
}

func (ManagedInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceArgs)(nil)).Elem()
}

type ManagedInstanceInput interface {
	pulumi.Input

	ToManagedInstanceOutput() ManagedInstanceOutput
	ToManagedInstanceOutputWithContext(ctx context.Context) ManagedInstanceOutput
}

func (*ManagedInstance) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstance)(nil))
}

func (i *ManagedInstance) ToManagedInstanceOutput() ManagedInstanceOutput {
	return i.ToManagedInstanceOutputWithContext(context.Background())
}

func (i *ManagedInstance) ToManagedInstanceOutputWithContext(ctx context.Context) ManagedInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceOutput)
}

type ManagedInstanceOutput struct{ *pulumi.OutputState }

func (ManagedInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstance)(nil))
}

func (o ManagedInstanceOutput) ToManagedInstanceOutput() ManagedInstanceOutput {
	return o
}

func (o ManagedInstanceOutput) ToManagedInstanceOutputWithContext(ctx context.Context) ManagedInstanceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagedInstanceOutput{})
}

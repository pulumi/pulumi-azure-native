


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedInstance struct {
	pulumi.CustomResourceState

	AdministratorLogin         pulumi.StringPtrOutput            `pulumi:"administratorLogin"`
	Collation                  pulumi.StringPtrOutput            `pulumi:"collation"`
	DnsZone                    pulumi.StringOutput               `pulumi:"dnsZone"`
	FullyQualifiedDomainName   pulumi.StringOutput               `pulumi:"fullyQualifiedDomainName"`
	Identity                   ResourceIdentityResponsePtrOutput `pulumi:"identity"`
	InstancePoolId             pulumi.StringPtrOutput            `pulumi:"instancePoolId"`
	LicenseType                pulumi.StringPtrOutput            `pulumi:"licenseType"`
	Location                   pulumi.StringOutput               `pulumi:"location"`
	MaintenanceConfigurationId pulumi.StringPtrOutput            `pulumi:"maintenanceConfigurationId"`
	MinimalTlsVersion          pulumi.StringPtrOutput            `pulumi:"minimalTlsVersion"`
	Name                       pulumi.StringOutput               `pulumi:"name"`
	ProxyOverride              pulumi.StringPtrOutput            `pulumi:"proxyOverride"`
	PublicDataEndpointEnabled  pulumi.BoolPtrOutput              `pulumi:"publicDataEndpointEnabled"`
	Sku                        SkuResponsePtrOutput              `pulumi:"sku"`
	State                      pulumi.StringOutput               `pulumi:"state"`
	StorageSizeInGB            pulumi.IntPtrOutput               `pulumi:"storageSizeInGB"`
	SubnetId                   pulumi.StringPtrOutput            `pulumi:"subnetId"`
	Tags                       pulumi.StringMapOutput            `pulumi:"tags"`
	TimezoneId                 pulumi.StringPtrOutput            `pulumi:"timezoneId"`
	Type                       pulumi.StringOutput               `pulumi:"type"`
	VCores                     pulumi.IntPtrOutput               `pulumi:"vCores"`
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
			Type: pulumi.String("azure-native:sql:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ManagedInstance"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedInstance
	err := ctx.RegisterResource("azure-native:sql/v20180601preview:ManagedInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedInstanceState, opts ...pulumi.ResourceOption) (*ManagedInstance, error) {
	var resource ManagedInstance
	err := ctx.ReadResource("azure-native:sql/v20180601preview:ManagedInstance", name, id, state, &resource, opts...)
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
	StorageSizeInGB            *int              `pulumi:"storageSizeInGB"`
	SubnetId                   *string           `pulumi:"subnetId"`
	Tags                       map[string]string `pulumi:"tags"`
	TimezoneId                 *string           `pulumi:"timezoneId"`
	VCores                     *int              `pulumi:"vCores"`
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
	StorageSizeInGB            pulumi.IntPtrInput
	SubnetId                   pulumi.StringPtrInput
	Tags                       pulumi.StringMapInput
	TimezoneId                 pulumi.StringPtrInput
	VCores                     pulumi.IntPtrInput
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
	return reflect.TypeOf((**ManagedInstance)(nil)).Elem()
}

func (i *ManagedInstance) ToManagedInstanceOutput() ManagedInstanceOutput {
	return i.ToManagedInstanceOutputWithContext(context.Background())
}

func (i *ManagedInstance) ToManagedInstanceOutputWithContext(ctx context.Context) ManagedInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceOutput)
}

type ManagedInstanceOutput struct{ *pulumi.OutputState }

func (ManagedInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstance)(nil)).Elem()
}

func (o ManagedInstanceOutput) ToManagedInstanceOutput() ManagedInstanceOutput {
	return o
}

func (o ManagedInstanceOutput) ToManagedInstanceOutputWithContext(ctx context.Context) ManagedInstanceOutput {
	return o
}

func (o ManagedInstanceOutput) AdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.AdministratorLogin }).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.Collation }).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceOutput) DnsZone() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringOutput { return v.DnsZone }).(pulumi.StringOutput)
}

func (o ManagedInstanceOutput) FullyQualifiedDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringOutput { return v.FullyQualifiedDomainName }).(pulumi.StringOutput)
}

func (o ManagedInstanceOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ManagedInstance) ResourceIdentityResponsePtrOutput { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o ManagedInstanceOutput) InstancePoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.InstancePoolId }).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ManagedInstanceOutput) MaintenanceConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.MaintenanceConfigurationId }).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceOutput) MinimalTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.MinimalTlsVersion }).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedInstanceOutput) ProxyOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.ProxyOverride }).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceOutput) PublicDataEndpointEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.BoolPtrOutput { return v.PublicDataEndpointEnabled }).(pulumi.BoolPtrOutput)
}

func (o ManagedInstanceOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *ManagedInstance) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o ManagedInstanceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o ManagedInstanceOutput) StorageSizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.IntPtrOutput { return v.StorageSizeInGB }).(pulumi.IntPtrOutput)
}

func (o ManagedInstanceOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ManagedInstanceOutput) TimezoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.TimezoneId }).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedInstanceOutput) VCores() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.IntPtrOutput { return v.VCores }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedInstanceOutput{})
}




package v20170601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type DomainService struct {
	pulumi.CustomResourceState

	DeploymentId              pulumi.StringOutput                     `pulumi:"deploymentId"`
	DomainConfigurationType   pulumi.StringPtrOutput                  `pulumi:"domainConfigurationType"`
	DomainControllerIpAddress pulumi.StringArrayOutput                `pulumi:"domainControllerIpAddress"`
	DomainName                pulumi.StringPtrOutput                  `pulumi:"domainName"`
	DomainSecuritySettings    DomainSecuritySettingsResponsePtrOutput `pulumi:"domainSecuritySettings"`
	Etag                      pulumi.StringPtrOutput                  `pulumi:"etag"`
	FilteredSync              pulumi.StringPtrOutput                  `pulumi:"filteredSync"`
	HealthAlerts              HealthAlertResponseArrayOutput          `pulumi:"healthAlerts"`
	HealthLastEvaluated       pulumi.StringOutput                     `pulumi:"healthLastEvaluated"`
	HealthMonitors            HealthMonitorResponseArrayOutput        `pulumi:"healthMonitors"`
	LdapsSettings             LdapsSettingsResponsePtrOutput          `pulumi:"ldapsSettings"`
	Location                  pulumi.StringPtrOutput                  `pulumi:"location"`
	MigrationProperties       MigrationPropertiesResponseOutput       `pulumi:"migrationProperties"`
	Name                      pulumi.StringOutput                     `pulumi:"name"`
	NotificationSettings      NotificationSettingsResponsePtrOutput   `pulumi:"notificationSettings"`
	ProvisioningState         pulumi.StringOutput                     `pulumi:"provisioningState"`
	ResourceForestSettings    ResourceForestSettingsResponsePtrOutput `pulumi:"resourceForestSettings"`
	ServiceStatus             pulumi.StringOutput                     `pulumi:"serviceStatus"`
	Sku                       pulumi.StringPtrOutput                  `pulumi:"sku"`
	SubnetId                  pulumi.StringPtrOutput                  `pulumi:"subnetId"`
	Tags                      pulumi.StringMapOutput                  `pulumi:"tags"`
	TenantId                  pulumi.StringOutput                     `pulumi:"tenantId"`
	Type                      pulumi.StringOutput                     `pulumi:"type"`
	Version                   pulumi.IntOutput                        `pulumi:"version"`
	VnetSiteId                pulumi.StringOutput                     `pulumi:"vnetSiteId"`
}


func NewDomainService(ctx *pulumi.Context,
	name string, args *DomainServiceArgs, opts ...pulumi.ResourceOption) (*DomainService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.DomainSecuritySettings != nil {
		args.DomainSecuritySettings = args.DomainSecuritySettings.ToDomainSecuritySettingsPtrOutput().ApplyT(func(v *DomainSecuritySettings) *DomainSecuritySettings { return v.Defaults() }).(DomainSecuritySettingsPtrOutput)
	}
	if args.LdapsSettings != nil {
		args.LdapsSettings = args.LdapsSettings.ToLdapsSettingsPtrOutput().ApplyT(func(v *LdapsSettings) *LdapsSettings { return v.Defaults() }).(LdapsSettingsPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:aad:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20170101:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20200101:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20210301:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20210501:DomainService"),
		},
	})
	opts = append(opts, aliases)
	var resource DomainService
	err := ctx.RegisterResource("azure-native:aad/v20170601:DomainService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDomainService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainServiceState, opts ...pulumi.ResourceOption) (*DomainService, error) {
	var resource DomainService
	err := ctx.ReadResource("azure-native:aad/v20170601:DomainService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type domainServiceState struct {
}

type DomainServiceState struct {
}

func (DomainServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainServiceState)(nil)).Elem()
}

type domainServiceArgs struct {
	DomainConfigurationType *string                 `pulumi:"domainConfigurationType"`
	DomainName              *string                 `pulumi:"domainName"`
	DomainSecuritySettings  *DomainSecuritySettings `pulumi:"domainSecuritySettings"`
	DomainServiceName       *string                 `pulumi:"domainServiceName"`
	FilteredSync            *string                 `pulumi:"filteredSync"`
	LdapsSettings           *LdapsSettings          `pulumi:"ldapsSettings"`
	Location                *string                 `pulumi:"location"`
	NotificationSettings    *NotificationSettings   `pulumi:"notificationSettings"`
	ResourceForestSettings  *ResourceForestSettings `pulumi:"resourceForestSettings"`
	ResourceGroupName       string                  `pulumi:"resourceGroupName"`
	Sku                     *string                 `pulumi:"sku"`
	SubnetId                *string                 `pulumi:"subnetId"`
	Tags                    map[string]string       `pulumi:"tags"`
}


type DomainServiceArgs struct {
	DomainConfigurationType pulumi.StringPtrInput
	DomainName              pulumi.StringPtrInput
	DomainSecuritySettings  DomainSecuritySettingsPtrInput
	DomainServiceName       pulumi.StringPtrInput
	FilteredSync            pulumi.StringPtrInput
	LdapsSettings           LdapsSettingsPtrInput
	Location                pulumi.StringPtrInput
	NotificationSettings    NotificationSettingsPtrInput
	ResourceForestSettings  ResourceForestSettingsPtrInput
	ResourceGroupName       pulumi.StringInput
	Sku                     pulumi.StringPtrInput
	SubnetId                pulumi.StringPtrInput
	Tags                    pulumi.StringMapInput
}

func (DomainServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainServiceArgs)(nil)).Elem()
}

type DomainServiceInput interface {
	pulumi.Input

	ToDomainServiceOutput() DomainServiceOutput
	ToDomainServiceOutputWithContext(ctx context.Context) DomainServiceOutput
}

func (*DomainService) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainService)(nil)).Elem()
}

func (i *DomainService) ToDomainServiceOutput() DomainServiceOutput {
	return i.ToDomainServiceOutputWithContext(context.Background())
}

func (i *DomainService) ToDomainServiceOutputWithContext(ctx context.Context) DomainServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainServiceOutput)
}

type DomainServiceOutput struct{ *pulumi.OutputState }

func (DomainServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainService)(nil)).Elem()
}

func (o DomainServiceOutput) ToDomainServiceOutput() DomainServiceOutput {
	return o
}

func (o DomainServiceOutput) ToDomainServiceOutputWithContext(ctx context.Context) DomainServiceOutput {
	return o
}

func (o DomainServiceOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.DeploymentId }).(pulumi.StringOutput)
}

func (o DomainServiceOutput) DomainConfigurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringPtrOutput { return v.DomainConfigurationType }).(pulumi.StringPtrOutput)
}

func (o DomainServiceOutput) DomainControllerIpAddress() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringArrayOutput { return v.DomainControllerIpAddress }).(pulumi.StringArrayOutput)
}

func (o DomainServiceOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringPtrOutput { return v.DomainName }).(pulumi.StringPtrOutput)
}

func (o DomainServiceOutput) DomainSecuritySettings() DomainSecuritySettingsResponsePtrOutput {
	return o.ApplyT(func(v *DomainService) DomainSecuritySettingsResponsePtrOutput { return v.DomainSecuritySettings }).(DomainSecuritySettingsResponsePtrOutput)
}

func (o DomainServiceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o DomainServiceOutput) FilteredSync() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringPtrOutput { return v.FilteredSync }).(pulumi.StringPtrOutput)
}

func (o DomainServiceOutput) HealthAlerts() HealthAlertResponseArrayOutput {
	return o.ApplyT(func(v *DomainService) HealthAlertResponseArrayOutput { return v.HealthAlerts }).(HealthAlertResponseArrayOutput)
}

func (o DomainServiceOutput) HealthLastEvaluated() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.HealthLastEvaluated }).(pulumi.StringOutput)
}

func (o DomainServiceOutput) HealthMonitors() HealthMonitorResponseArrayOutput {
	return o.ApplyT(func(v *DomainService) HealthMonitorResponseArrayOutput { return v.HealthMonitors }).(HealthMonitorResponseArrayOutput)
}

func (o DomainServiceOutput) LdapsSettings() LdapsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *DomainService) LdapsSettingsResponsePtrOutput { return v.LdapsSettings }).(LdapsSettingsResponsePtrOutput)
}

func (o DomainServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o DomainServiceOutput) MigrationProperties() MigrationPropertiesResponseOutput {
	return o.ApplyT(func(v *DomainService) MigrationPropertiesResponseOutput { return v.MigrationProperties }).(MigrationPropertiesResponseOutput)
}

func (o DomainServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DomainServiceOutput) NotificationSettings() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v *DomainService) NotificationSettingsResponsePtrOutput { return v.NotificationSettings }).(NotificationSettingsResponsePtrOutput)
}

func (o DomainServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DomainServiceOutput) ResourceForestSettings() ResourceForestSettingsResponsePtrOutput {
	return o.ApplyT(func(v *DomainService) ResourceForestSettingsResponsePtrOutput { return v.ResourceForestSettings }).(ResourceForestSettingsResponsePtrOutput)
}

func (o DomainServiceOutput) ServiceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.ServiceStatus }).(pulumi.StringOutput)
}

func (o DomainServiceOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringPtrOutput { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o DomainServiceOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o DomainServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DomainServiceOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o DomainServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o DomainServiceOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *DomainService) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

func (o DomainServiceOutput) VnetSiteId() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.VnetSiteId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainServiceOutput{})
}

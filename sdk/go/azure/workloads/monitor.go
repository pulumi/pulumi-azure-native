


package workloads

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Monitor struct {
	pulumi.CustomResourceState

	AppLocation                       pulumi.StringPtrOutput                       `pulumi:"appLocation"`
	Errors                            MonitorPropertiesResponseErrorsOutput        `pulumi:"errors"`
	Identity                          UserAssignedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	Location                          pulumi.StringOutput                          `pulumi:"location"`
	LogAnalyticsWorkspaceArmId        pulumi.StringPtrOutput                       `pulumi:"logAnalyticsWorkspaceArmId"`
	ManagedResourceGroupConfiguration ManagedRGConfigurationResponsePtrOutput      `pulumi:"managedResourceGroupConfiguration"`
	MonitorSubnet                     pulumi.StringPtrOutput                       `pulumi:"monitorSubnet"`
	MsiArmId                          pulumi.StringOutput                          `pulumi:"msiArmId"`
	Name                              pulumi.StringOutput                          `pulumi:"name"`
	ProvisioningState                 pulumi.StringOutput                          `pulumi:"provisioningState"`
	RoutingPreference                 pulumi.StringPtrOutput                       `pulumi:"routingPreference"`
	SystemData                        SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags                              pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                              pulumi.StringOutput                          `pulumi:"type"`
	ZoneRedundancyPreference          pulumi.StringPtrOutput                       `pulumi:"zoneRedundancyPreference"`
}


func NewMonitor(ctx *pulumi.Context,
	name string, args *MonitorArgs, opts ...pulumi.ResourceOption) (*Monitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:workloads/v20211201preview:monitor"),
		},
	})
	opts = append(opts, aliases)
	var resource Monitor
	err := ctx.RegisterResource("azure-native:workloads:monitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitorState, opts ...pulumi.ResourceOption) (*Monitor, error) {
	var resource Monitor
	err := ctx.ReadResource("azure-native:workloads:monitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type monitorState struct {
}

type MonitorState struct {
}

func (MonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorState)(nil)).Elem()
}

type monitorArgs struct {
	AppLocation                       *string                      `pulumi:"appLocation"`
	Identity                          *UserAssignedServiceIdentity `pulumi:"identity"`
	Location                          *string                      `pulumi:"location"`
	LogAnalyticsWorkspaceArmId        *string                      `pulumi:"logAnalyticsWorkspaceArmId"`
	ManagedResourceGroupConfiguration *ManagedRGConfiguration      `pulumi:"managedResourceGroupConfiguration"`
	MonitorName                       *string                      `pulumi:"monitorName"`
	MonitorSubnet                     *string                      `pulumi:"monitorSubnet"`
	ResourceGroupName                 string                       `pulumi:"resourceGroupName"`
	RoutingPreference                 *string                      `pulumi:"routingPreference"`
	Tags                              map[string]string            `pulumi:"tags"`
	ZoneRedundancyPreference          *string                      `pulumi:"zoneRedundancyPreference"`
}


type MonitorArgs struct {
	AppLocation                       pulumi.StringPtrInput
	Identity                          UserAssignedServiceIdentityPtrInput
	Location                          pulumi.StringPtrInput
	LogAnalyticsWorkspaceArmId        pulumi.StringPtrInput
	ManagedResourceGroupConfiguration ManagedRGConfigurationPtrInput
	MonitorName                       pulumi.StringPtrInput
	MonitorSubnet                     pulumi.StringPtrInput
	ResourceGroupName                 pulumi.StringInput
	RoutingPreference                 pulumi.StringPtrInput
	Tags                              pulumi.StringMapInput
	ZoneRedundancyPreference          pulumi.StringPtrInput
}

func (MonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorArgs)(nil)).Elem()
}

type MonitorInput interface {
	pulumi.Input

	ToMonitorOutput() MonitorOutput
	ToMonitorOutputWithContext(ctx context.Context) MonitorOutput
}

func (*Monitor) ElementType() reflect.Type {
	return reflect.TypeOf((**Monitor)(nil)).Elem()
}

func (i *Monitor) ToMonitorOutput() MonitorOutput {
	return i.ToMonitorOutputWithContext(context.Background())
}

func (i *Monitor) ToMonitorOutputWithContext(ctx context.Context) MonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorOutput)
}

type MonitorOutput struct{ *pulumi.OutputState }

func (MonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Monitor)(nil)).Elem()
}

func (o MonitorOutput) ToMonitorOutput() MonitorOutput {
	return o
}

func (o MonitorOutput) ToMonitorOutputWithContext(ctx context.Context) MonitorOutput {
	return o
}

func (o MonitorOutput) AppLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringPtrOutput { return v.AppLocation }).(pulumi.StringPtrOutput)
}

func (o MonitorOutput) Errors() MonitorPropertiesResponseErrorsOutput {
	return o.ApplyT(func(v *Monitor) MonitorPropertiesResponseErrorsOutput { return v.Errors }).(MonitorPropertiesResponseErrorsOutput)
}

func (o MonitorOutput) Identity() UserAssignedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Monitor) UserAssignedServiceIdentityResponsePtrOutput { return v.Identity }).(UserAssignedServiceIdentityResponsePtrOutput)
}

func (o MonitorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o MonitorOutput) LogAnalyticsWorkspaceArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringPtrOutput { return v.LogAnalyticsWorkspaceArmId }).(pulumi.StringPtrOutput)
}

func (o MonitorOutput) ManagedResourceGroupConfiguration() ManagedRGConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Monitor) ManagedRGConfigurationResponsePtrOutput { return v.ManagedResourceGroupConfiguration }).(ManagedRGConfigurationResponsePtrOutput)
}

func (o MonitorOutput) MonitorSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringPtrOutput { return v.MonitorSubnet }).(pulumi.StringPtrOutput)
}

func (o MonitorOutput) MsiArmId() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.MsiArmId }).(pulumi.StringOutput)
}

func (o MonitorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MonitorOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MonitorOutput) RoutingPreference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringPtrOutput { return v.RoutingPreference }).(pulumi.StringPtrOutput)
}

func (o MonitorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Monitor) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MonitorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o MonitorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o MonitorOutput) ZoneRedundancyPreference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringPtrOutput { return v.ZoneRedundancyPreference }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MonitorOutput{})
}

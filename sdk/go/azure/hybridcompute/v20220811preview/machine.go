


package v20220811preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Machine struct {
	pulumi.CustomResourceState

	AdFqdn                     pulumi.StringOutput                 `pulumi:"adFqdn"`
	AgentConfiguration         AgentConfigurationResponseOutput    `pulumi:"agentConfiguration"`
	AgentVersion               pulumi.StringOutput                 `pulumi:"agentVersion"`
	ClientPublicKey            pulumi.StringPtrOutput              `pulumi:"clientPublicKey"`
	CloudMetadata              CloudMetadataResponsePtrOutput      `pulumi:"cloudMetadata"`
	DetectedProperties         pulumi.StringMapOutput              `pulumi:"detectedProperties"`
	DisplayName                pulumi.StringOutput                 `pulumi:"displayName"`
	DnsFqdn                    pulumi.StringOutput                 `pulumi:"dnsFqdn"`
	DomainName                 pulumi.StringOutput                 `pulumi:"domainName"`
	ErrorDetails               ErrorDetailResponseArrayOutput      `pulumi:"errorDetails"`
	Identity                   IdentityResponsePtrOutput           `pulumi:"identity"`
	LastStatusChange           pulumi.StringOutput                 `pulumi:"lastStatusChange"`
	Location                   pulumi.StringOutput                 `pulumi:"location"`
	LocationData               LocationDataResponsePtrOutput       `pulumi:"locationData"`
	MachineFqdn                pulumi.StringOutput                 `pulumi:"machineFqdn"`
	MssqlDiscovered            pulumi.StringPtrOutput              `pulumi:"mssqlDiscovered"`
	Name                       pulumi.StringOutput                 `pulumi:"name"`
	OsName                     pulumi.StringOutput                 `pulumi:"osName"`
	OsProfile                  OSProfileResponsePtrOutput          `pulumi:"osProfile"`
	OsSku                      pulumi.StringOutput                 `pulumi:"osSku"`
	OsType                     pulumi.StringPtrOutput              `pulumi:"osType"`
	OsVersion                  pulumi.StringOutput                 `pulumi:"osVersion"`
	ParentClusterResourceId    pulumi.StringPtrOutput              `pulumi:"parentClusterResourceId"`
	PrivateLinkScopeResourceId pulumi.StringPtrOutput              `pulumi:"privateLinkScopeResourceId"`
	ProvisioningState          pulumi.StringOutput                 `pulumi:"provisioningState"`
	Resources                  MachineExtensionResponseArrayOutput `pulumi:"resources"`
	ServiceStatuses            ServiceStatusesResponsePtrOutput    `pulumi:"serviceStatuses"`
	Status                     pulumi.StringOutput                 `pulumi:"status"`
	SystemData                 SystemDataResponseOutput            `pulumi:"systemData"`
	Tags                       pulumi.StringMapOutput              `pulumi:"tags"`
	Type                       pulumi.StringOutput                 `pulumi:"type"`
	VmId                       pulumi.StringPtrOutput              `pulumi:"vmId"`
	VmUuid                     pulumi.StringOutput                 `pulumi:"vmUuid"`
}


func NewMachine(ctx *pulumi.Context,
	name string, args *MachineArgs, opts ...pulumi.ResourceOption) (*Machine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcompute:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20190318preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20190802preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20191212:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200730preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200802:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200815preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210128preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210325preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210422preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210517preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210520:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210610preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20211210preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20220310:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20220510preview:Machine"),
		},
	})
	opts = append(opts, aliases)
	var resource Machine
	err := ctx.RegisterResource("azure-native:hybridcompute/v20220811preview:Machine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineState, opts ...pulumi.ResourceOption) (*Machine, error) {
	var resource Machine
	err := ctx.ReadResource("azure-native:hybridcompute/v20220811preview:Machine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type machineState struct {
}

type MachineState struct {
}

func (MachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineState)(nil)).Elem()
}

type machineArgs struct {
	ClientPublicKey            *string           `pulumi:"clientPublicKey"`
	Identity                   *Identity         `pulumi:"identity"`
	Location                   *string           `pulumi:"location"`
	LocationData               *LocationData     `pulumi:"locationData"`
	MachineName                *string           `pulumi:"machineName"`
	MssqlDiscovered            *string           `pulumi:"mssqlDiscovered"`
	OsProfile                  *OSProfile        `pulumi:"osProfile"`
	OsType                     *string           `pulumi:"osType"`
	ParentClusterResourceId    *string           `pulumi:"parentClusterResourceId"`
	PrivateLinkScopeResourceId *string           `pulumi:"privateLinkScopeResourceId"`
	ResourceGroupName          string            `pulumi:"resourceGroupName"`
	ServiceStatuses            *ServiceStatuses  `pulumi:"serviceStatuses"`
	Tags                       map[string]string `pulumi:"tags"`
	VmId                       *string           `pulumi:"vmId"`
}


type MachineArgs struct {
	ClientPublicKey            pulumi.StringPtrInput
	Identity                   IdentityPtrInput
	Location                   pulumi.StringPtrInput
	LocationData               LocationDataPtrInput
	MachineName                pulumi.StringPtrInput
	MssqlDiscovered            pulumi.StringPtrInput
	OsProfile                  OSProfilePtrInput
	OsType                     pulumi.StringPtrInput
	ParentClusterResourceId    pulumi.StringPtrInput
	PrivateLinkScopeResourceId pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	ServiceStatuses            ServiceStatusesPtrInput
	Tags                       pulumi.StringMapInput
	VmId                       pulumi.StringPtrInput
}

func (MachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineArgs)(nil)).Elem()
}

type MachineInput interface {
	pulumi.Input

	ToMachineOutput() MachineOutput
	ToMachineOutputWithContext(ctx context.Context) MachineOutput
}

func (*Machine) ElementType() reflect.Type {
	return reflect.TypeOf((**Machine)(nil)).Elem()
}

func (i *Machine) ToMachineOutput() MachineOutput {
	return i.ToMachineOutputWithContext(context.Background())
}

func (i *Machine) ToMachineOutputWithContext(ctx context.Context) MachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineOutput)
}

type MachineOutput struct{ *pulumi.OutputState }

func (MachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Machine)(nil)).Elem()
}

func (o MachineOutput) ToMachineOutput() MachineOutput {
	return o
}

func (o MachineOutput) ToMachineOutputWithContext(ctx context.Context) MachineOutput {
	return o
}

func (o MachineOutput) AdFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.AdFqdn }).(pulumi.StringOutput)
}

func (o MachineOutput) AgentConfiguration() AgentConfigurationResponseOutput {
	return o.ApplyT(func(v *Machine) AgentConfigurationResponseOutput { return v.AgentConfiguration }).(AgentConfigurationResponseOutput)
}

func (o MachineOutput) AgentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.AgentVersion }).(pulumi.StringOutput)
}

func (o MachineOutput) ClientPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.ClientPublicKey }).(pulumi.StringPtrOutput)
}

func (o MachineOutput) CloudMetadata() CloudMetadataResponsePtrOutput {
	return o.ApplyT(func(v *Machine) CloudMetadataResponsePtrOutput { return v.CloudMetadata }).(CloudMetadataResponsePtrOutput)
}

func (o MachineOutput) DetectedProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringMapOutput { return v.DetectedProperties }).(pulumi.StringMapOutput)
}

func (o MachineOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o MachineOutput) DnsFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.DnsFqdn }).(pulumi.StringOutput)
}

func (o MachineOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

func (o MachineOutput) ErrorDetails() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v *Machine) ErrorDetailResponseArrayOutput { return v.ErrorDetails }).(ErrorDetailResponseArrayOutput)
}

func (o MachineOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Machine) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o MachineOutput) LastStatusChange() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.LastStatusChange }).(pulumi.StringOutput)
}

func (o MachineOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o MachineOutput) LocationData() LocationDataResponsePtrOutput {
	return o.ApplyT(func(v *Machine) LocationDataResponsePtrOutput { return v.LocationData }).(LocationDataResponsePtrOutput)
}

func (o MachineOutput) MachineFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.MachineFqdn }).(pulumi.StringOutput)
}

func (o MachineOutput) MssqlDiscovered() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.MssqlDiscovered }).(pulumi.StringPtrOutput)
}

func (o MachineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MachineOutput) OsName() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.OsName }).(pulumi.StringOutput)
}

func (o MachineOutput) OsProfile() OSProfileResponsePtrOutput {
	return o.ApplyT(func(v *Machine) OSProfileResponsePtrOutput { return v.OsProfile }).(OSProfileResponsePtrOutput)
}

func (o MachineOutput) OsSku() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.OsSku }).(pulumi.StringOutput)
}

func (o MachineOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o MachineOutput) OsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.OsVersion }).(pulumi.StringOutput)
}

func (o MachineOutput) ParentClusterResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.ParentClusterResourceId }).(pulumi.StringPtrOutput)
}

func (o MachineOutput) PrivateLinkScopeResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.PrivateLinkScopeResourceId }).(pulumi.StringPtrOutput)
}

func (o MachineOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MachineOutput) Resources() MachineExtensionResponseArrayOutput {
	return o.ApplyT(func(v *Machine) MachineExtensionResponseArrayOutput { return v.Resources }).(MachineExtensionResponseArrayOutput)
}

func (o MachineOutput) ServiceStatuses() ServiceStatusesResponsePtrOutput {
	return o.ApplyT(func(v *Machine) ServiceStatusesResponsePtrOutput { return v.ServiceStatuses }).(ServiceStatusesResponsePtrOutput)
}

func (o MachineOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o MachineOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Machine) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MachineOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o MachineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o MachineOutput) VmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.VmId }).(pulumi.StringPtrOutput)
}

func (o MachineOutput) VmUuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.VmUuid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MachineOutput{})
}

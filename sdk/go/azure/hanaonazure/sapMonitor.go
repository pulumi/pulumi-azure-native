


package hanaonazure

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SapMonitor struct {
	pulumi.CustomResourceState

	EnableCustomerAnalytics        pulumi.BoolPtrOutput   `pulumi:"enableCustomerAnalytics"`
	Location                       pulumi.StringOutput    `pulumi:"location"`
	LogAnalyticsWorkspaceArmId     pulumi.StringPtrOutput `pulumi:"logAnalyticsWorkspaceArmId"`
	LogAnalyticsWorkspaceId        pulumi.StringPtrOutput `pulumi:"logAnalyticsWorkspaceId"`
	LogAnalyticsWorkspaceSharedKey pulumi.StringPtrOutput `pulumi:"logAnalyticsWorkspaceSharedKey"`
	ManagedResourceGroupName       pulumi.StringOutput    `pulumi:"managedResourceGroupName"`
	MonitorSubnet                  pulumi.StringPtrOutput `pulumi:"monitorSubnet"`
	Name                           pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState              pulumi.StringOutput    `pulumi:"provisioningState"`
	SapMonitorCollectorVersion     pulumi.StringOutput    `pulumi:"sapMonitorCollectorVersion"`
	Tags                           pulumi.StringMapOutput `pulumi:"tags"`
	Type                           pulumi.StringOutput    `pulumi:"type"`
}


func NewSapMonitor(ctx *pulumi.Context,
	name string, args *SapMonitorArgs, opts ...pulumi.ResourceOption) (*SapMonitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hanaonazure/v20200207preview:SapMonitor"),
		},
	})
	opts = append(opts, aliases)
	var resource SapMonitor
	err := ctx.RegisterResource("azure-native:hanaonazure:SapMonitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSapMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SapMonitorState, opts ...pulumi.ResourceOption) (*SapMonitor, error) {
	var resource SapMonitor
	err := ctx.ReadResource("azure-native:hanaonazure:SapMonitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sapMonitorState struct {
}

type SapMonitorState struct {
}

func (SapMonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*sapMonitorState)(nil)).Elem()
}

type sapMonitorArgs struct {
	EnableCustomerAnalytics        *bool             `pulumi:"enableCustomerAnalytics"`
	Location                       *string           `pulumi:"location"`
	LogAnalyticsWorkspaceArmId     *string           `pulumi:"logAnalyticsWorkspaceArmId"`
	LogAnalyticsWorkspaceId        *string           `pulumi:"logAnalyticsWorkspaceId"`
	LogAnalyticsWorkspaceSharedKey *string           `pulumi:"logAnalyticsWorkspaceSharedKey"`
	MonitorSubnet                  *string           `pulumi:"monitorSubnet"`
	ResourceGroupName              string            `pulumi:"resourceGroupName"`
	SapMonitorName                 *string           `pulumi:"sapMonitorName"`
	Tags                           map[string]string `pulumi:"tags"`
}


type SapMonitorArgs struct {
	EnableCustomerAnalytics        pulumi.BoolPtrInput
	Location                       pulumi.StringPtrInput
	LogAnalyticsWorkspaceArmId     pulumi.StringPtrInput
	LogAnalyticsWorkspaceId        pulumi.StringPtrInput
	LogAnalyticsWorkspaceSharedKey pulumi.StringPtrInput
	MonitorSubnet                  pulumi.StringPtrInput
	ResourceGroupName              pulumi.StringInput
	SapMonitorName                 pulumi.StringPtrInput
	Tags                           pulumi.StringMapInput
}

func (SapMonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sapMonitorArgs)(nil)).Elem()
}

type SapMonitorInput interface {
	pulumi.Input

	ToSapMonitorOutput() SapMonitorOutput
	ToSapMonitorOutputWithContext(ctx context.Context) SapMonitorOutput
}

func (*SapMonitor) ElementType() reflect.Type {
	return reflect.TypeOf((**SapMonitor)(nil)).Elem()
}

func (i *SapMonitor) ToSapMonitorOutput() SapMonitorOutput {
	return i.ToSapMonitorOutputWithContext(context.Background())
}

func (i *SapMonitor) ToSapMonitorOutputWithContext(ctx context.Context) SapMonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SapMonitorOutput)
}

type SapMonitorOutput struct{ *pulumi.OutputState }

func (SapMonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SapMonitor)(nil)).Elem()
}

func (o SapMonitorOutput) ToSapMonitorOutput() SapMonitorOutput {
	return o
}

func (o SapMonitorOutput) ToSapMonitorOutputWithContext(ctx context.Context) SapMonitorOutput {
	return o
}

func (o SapMonitorOutput) EnableCustomerAnalytics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.BoolPtrOutput { return v.EnableCustomerAnalytics }).(pulumi.BoolPtrOutput)
}

func (o SapMonitorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SapMonitorOutput) LogAnalyticsWorkspaceArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringPtrOutput { return v.LogAnalyticsWorkspaceArmId }).(pulumi.StringPtrOutput)
}

func (o SapMonitorOutput) LogAnalyticsWorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringPtrOutput { return v.LogAnalyticsWorkspaceId }).(pulumi.StringPtrOutput)
}

func (o SapMonitorOutput) LogAnalyticsWorkspaceSharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringPtrOutput { return v.LogAnalyticsWorkspaceSharedKey }).(pulumi.StringPtrOutput)
}

func (o SapMonitorOutput) ManagedResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringOutput { return v.ManagedResourceGroupName }).(pulumi.StringOutput)
}

func (o SapMonitorOutput) MonitorSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringPtrOutput { return v.MonitorSubnet }).(pulumi.StringPtrOutput)
}

func (o SapMonitorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SapMonitorOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SapMonitorOutput) SapMonitorCollectorVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringOutput { return v.SapMonitorCollectorVersion }).(pulumi.StringOutput)
}

func (o SapMonitorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SapMonitorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SapMonitorOutput{})
}

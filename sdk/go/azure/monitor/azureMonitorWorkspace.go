


package monitor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureMonitorWorkspace struct {
	pulumi.CustomResourceState

	AccountId                pulumi.StringOutput                                     `pulumi:"accountId"`
	DefaultIngestionSettings MonitoringAccountResponseDefaultIngestionSettingsOutput `pulumi:"defaultIngestionSettings"`
	Etag                     pulumi.StringOutput                                     `pulumi:"etag"`
	Location                 pulumi.StringOutput                                     `pulumi:"location"`
	Metrics                  MonitoringAccountResponseMetricsOutput                  `pulumi:"metrics"`
	Name                     pulumi.StringOutput                                     `pulumi:"name"`
	ProvisioningState        pulumi.StringOutput                                     `pulumi:"provisioningState"`
	SystemData               SystemDataResponseOutput                                `pulumi:"systemData"`
	Tags                     pulumi.StringMapOutput                                  `pulumi:"tags"`
	Type                     pulumi.StringOutput                                     `pulumi:"type"`
}


func NewAzureMonitorWorkspace(ctx *pulumi.Context,
	name string, args *AzureMonitorWorkspaceArgs, opts ...pulumi.ResourceOption) (*AzureMonitorWorkspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:monitor/v20210603preview:AzureMonitorWorkspace"),
		},
	})
	opts = append(opts, aliases)
	var resource AzureMonitorWorkspace
	err := ctx.RegisterResource("azure-native:monitor:AzureMonitorWorkspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAzureMonitorWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureMonitorWorkspaceState, opts ...pulumi.ResourceOption) (*AzureMonitorWorkspace, error) {
	var resource AzureMonitorWorkspace
	err := ctx.ReadResource("azure-native:monitor:AzureMonitorWorkspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type azureMonitorWorkspaceState struct {
}

type AzureMonitorWorkspaceState struct {
}

func (AzureMonitorWorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureMonitorWorkspaceState)(nil)).Elem()
}

type azureMonitorWorkspaceArgs struct {
	Location              *string           `pulumi:"location"`
	MonitoringAccountName *string           `pulumi:"monitoringAccountName"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	Tags                  map[string]string `pulumi:"tags"`
}


type AzureMonitorWorkspaceArgs struct {
	Location              pulumi.StringPtrInput
	MonitoringAccountName pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	Tags                  pulumi.StringMapInput
}

func (AzureMonitorWorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureMonitorWorkspaceArgs)(nil)).Elem()
}

type AzureMonitorWorkspaceInput interface {
	pulumi.Input

	ToAzureMonitorWorkspaceOutput() AzureMonitorWorkspaceOutput
	ToAzureMonitorWorkspaceOutputWithContext(ctx context.Context) AzureMonitorWorkspaceOutput
}

func (*AzureMonitorWorkspace) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureMonitorWorkspace)(nil)).Elem()
}

func (i *AzureMonitorWorkspace) ToAzureMonitorWorkspaceOutput() AzureMonitorWorkspaceOutput {
	return i.ToAzureMonitorWorkspaceOutputWithContext(context.Background())
}

func (i *AzureMonitorWorkspace) ToAzureMonitorWorkspaceOutputWithContext(ctx context.Context) AzureMonitorWorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMonitorWorkspaceOutput)
}

type AzureMonitorWorkspaceOutput struct{ *pulumi.OutputState }

func (AzureMonitorWorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureMonitorWorkspace)(nil)).Elem()
}

func (o AzureMonitorWorkspaceOutput) ToAzureMonitorWorkspaceOutput() AzureMonitorWorkspaceOutput {
	return o
}

func (o AzureMonitorWorkspaceOutput) ToAzureMonitorWorkspaceOutputWithContext(ctx context.Context) AzureMonitorWorkspaceOutput {
	return o
}

func (o AzureMonitorWorkspaceOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureMonitorWorkspace) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

func (o AzureMonitorWorkspaceOutput) DefaultIngestionSettings() MonitoringAccountResponseDefaultIngestionSettingsOutput {
	return o.ApplyT(func(v *AzureMonitorWorkspace) MonitoringAccountResponseDefaultIngestionSettingsOutput {
		return v.DefaultIngestionSettings
	}).(MonitoringAccountResponseDefaultIngestionSettingsOutput)
}

func (o AzureMonitorWorkspaceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureMonitorWorkspace) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o AzureMonitorWorkspaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureMonitorWorkspace) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AzureMonitorWorkspaceOutput) Metrics() MonitoringAccountResponseMetricsOutput {
	return o.ApplyT(func(v *AzureMonitorWorkspace) MonitoringAccountResponseMetricsOutput { return v.Metrics }).(MonitoringAccountResponseMetricsOutput)
}

func (o AzureMonitorWorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureMonitorWorkspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AzureMonitorWorkspaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureMonitorWorkspace) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AzureMonitorWorkspaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AzureMonitorWorkspace) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AzureMonitorWorkspaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AzureMonitorWorkspace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AzureMonitorWorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureMonitorWorkspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureMonitorWorkspaceOutput{})
}

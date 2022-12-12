


package v20210603preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MonitoringAccount struct {
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


func NewMonitoringAccount(ctx *pulumi.Context,
	name string, args *MonitoringAccountArgs, opts ...pulumi.ResourceOption) (*MonitoringAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:monitor:MonitoringAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource MonitoringAccount
	err := ctx.RegisterResource("azure-native:monitor/v20210603preview:MonitoringAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMonitoringAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitoringAccountState, opts ...pulumi.ResourceOption) (*MonitoringAccount, error) {
	var resource MonitoringAccount
	err := ctx.ReadResource("azure-native:monitor/v20210603preview:MonitoringAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type monitoringAccountState struct {
}

type MonitoringAccountState struct {
}

func (MonitoringAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringAccountState)(nil)).Elem()
}

type monitoringAccountArgs struct {
	Location              *string           `pulumi:"location"`
	MonitoringAccountName *string           `pulumi:"monitoringAccountName"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	Tags                  map[string]string `pulumi:"tags"`
}


type MonitoringAccountArgs struct {
	Location              pulumi.StringPtrInput
	MonitoringAccountName pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	Tags                  pulumi.StringMapInput
}

func (MonitoringAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringAccountArgs)(nil)).Elem()
}

type MonitoringAccountInput interface {
	pulumi.Input

	ToMonitoringAccountOutput() MonitoringAccountOutput
	ToMonitoringAccountOutputWithContext(ctx context.Context) MonitoringAccountOutput
}

func (*MonitoringAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringAccount)(nil)).Elem()
}

func (i *MonitoringAccount) ToMonitoringAccountOutput() MonitoringAccountOutput {
	return i.ToMonitoringAccountOutputWithContext(context.Background())
}

func (i *MonitoringAccount) ToMonitoringAccountOutputWithContext(ctx context.Context) MonitoringAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringAccountOutput)
}

type MonitoringAccountOutput struct{ *pulumi.OutputState }

func (MonitoringAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringAccount)(nil)).Elem()
}

func (o MonitoringAccountOutput) ToMonitoringAccountOutput() MonitoringAccountOutput {
	return o
}

func (o MonitoringAccountOutput) ToMonitoringAccountOutputWithContext(ctx context.Context) MonitoringAccountOutput {
	return o
}

func (o MonitoringAccountOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringAccount) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

func (o MonitoringAccountOutput) DefaultIngestionSettings() MonitoringAccountResponseDefaultIngestionSettingsOutput {
	return o.ApplyT(func(v *MonitoringAccount) MonitoringAccountResponseDefaultIngestionSettingsOutput {
		return v.DefaultIngestionSettings
	}).(MonitoringAccountResponseDefaultIngestionSettingsOutput)
}

func (o MonitoringAccountOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringAccount) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o MonitoringAccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringAccount) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o MonitoringAccountOutput) Metrics() MonitoringAccountResponseMetricsOutput {
	return o.ApplyT(func(v *MonitoringAccount) MonitoringAccountResponseMetricsOutput { return v.Metrics }).(MonitoringAccountResponseMetricsOutput)
}

func (o MonitoringAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MonitoringAccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringAccount) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MonitoringAccountOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MonitoringAccount) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MonitoringAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MonitoringAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o MonitoringAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MonitoringAccountOutput{})
}




package v20200501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectionMonitor struct {
	pulumi.CustomResourceState

	AutoStart                   pulumi.BoolPtrOutput                                  `pulumi:"autoStart"`
	ConnectionMonitorType       pulumi.StringOutput                                   `pulumi:"connectionMonitorType"`
	Destination                 ConnectionMonitorDestinationResponsePtrOutput         `pulumi:"destination"`
	Endpoints                   ConnectionMonitorEndpointResponseArrayOutput          `pulumi:"endpoints"`
	Etag                        pulumi.StringOutput                                   `pulumi:"etag"`
	Location                    pulumi.StringPtrOutput                                `pulumi:"location"`
	MonitoringIntervalInSeconds pulumi.IntPtrOutput                                   `pulumi:"monitoringIntervalInSeconds"`
	MonitoringStatus            pulumi.StringOutput                                   `pulumi:"monitoringStatus"`
	Name                        pulumi.StringOutput                                   `pulumi:"name"`
	Notes                       pulumi.StringPtrOutput                                `pulumi:"notes"`
	Outputs                     ConnectionMonitorOutputResponseArrayOutput            `pulumi:"outputs"`
	ProvisioningState           pulumi.StringOutput                                   `pulumi:"provisioningState"`
	Source                      ConnectionMonitorSourceResponsePtrOutput              `pulumi:"source"`
	StartTime                   pulumi.StringOutput                                   `pulumi:"startTime"`
	Tags                        pulumi.StringMapOutput                                `pulumi:"tags"`
	TestConfigurations          ConnectionMonitorTestConfigurationResponseArrayOutput `pulumi:"testConfigurations"`
	TestGroups                  ConnectionMonitorTestGroupResponseArrayOutput         `pulumi:"testGroups"`
	Type                        pulumi.StringOutput                                   `pulumi:"type"`
}


func NewConnectionMonitor(ctx *pulumi.Context,
	name string, args *ConnectionMonitorArgs, opts ...pulumi.ResourceOption) (*ConnectionMonitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkWatcherName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkWatcherName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.AutoStart) {
		args.AutoStart = pulumi.BoolPtr(true)
	}
	if isZero(args.MonitoringIntervalInSeconds) {
		args.MonitoringIntervalInSeconds = pulumi.IntPtr(60)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ConnectionMonitor"),
		},
	})
	opts = append(opts, aliases)
	var resource ConnectionMonitor
	err := ctx.RegisterResource("azure-native:network/v20200501:ConnectionMonitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnectionMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionMonitorState, opts ...pulumi.ResourceOption) (*ConnectionMonitor, error) {
	var resource ConnectionMonitor
	err := ctx.ReadResource("azure-native:network/v20200501:ConnectionMonitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type connectionMonitorState struct {
}

type ConnectionMonitorState struct {
}

func (ConnectionMonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionMonitorState)(nil)).Elem()
}

type connectionMonitorArgs struct {
	AutoStart                   *bool                                `pulumi:"autoStart"`
	ConnectionMonitorName       *string                              `pulumi:"connectionMonitorName"`
	Destination                 *ConnectionMonitorDestination        `pulumi:"destination"`
	Endpoints                   []ConnectionMonitorEndpoint          `pulumi:"endpoints"`
	Location                    *string                              `pulumi:"location"`
	MonitoringIntervalInSeconds *int                                 `pulumi:"monitoringIntervalInSeconds"`
	NetworkWatcherName          string                               `pulumi:"networkWatcherName"`
	Notes                       *string                              `pulumi:"notes"`
	Outputs                     []ConnectionMonitorOutputType        `pulumi:"outputs"`
	ResourceGroupName           string                               `pulumi:"resourceGroupName"`
	Source                      *ConnectionMonitorSource             `pulumi:"source"`
	Tags                        map[string]string                    `pulumi:"tags"`
	TestConfigurations          []ConnectionMonitorTestConfiguration `pulumi:"testConfigurations"`
	TestGroups                  []ConnectionMonitorTestGroup         `pulumi:"testGroups"`
}


type ConnectionMonitorArgs struct {
	AutoStart                   pulumi.BoolPtrInput
	ConnectionMonitorName       pulumi.StringPtrInput
	Destination                 ConnectionMonitorDestinationPtrInput
	Endpoints                   ConnectionMonitorEndpointArrayInput
	Location                    pulumi.StringPtrInput
	MonitoringIntervalInSeconds pulumi.IntPtrInput
	NetworkWatcherName          pulumi.StringInput
	Notes                       pulumi.StringPtrInput
	Outputs                     ConnectionMonitorOutputTypeArrayInput
	ResourceGroupName           pulumi.StringInput
	Source                      ConnectionMonitorSourcePtrInput
	Tags                        pulumi.StringMapInput
	TestConfigurations          ConnectionMonitorTestConfigurationArrayInput
	TestGroups                  ConnectionMonitorTestGroupArrayInput
}

func (ConnectionMonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionMonitorArgs)(nil)).Elem()
}

type ConnectionMonitorInput interface {
	pulumi.Input

	ToConnectionMonitorOutput() ConnectionMonitorOutput
	ToConnectionMonitorOutputWithContext(ctx context.Context) ConnectionMonitorOutput
}

func (*ConnectionMonitor) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionMonitor)(nil)).Elem()
}

func (i *ConnectionMonitor) ToConnectionMonitorOutput() ConnectionMonitorOutput {
	return i.ToConnectionMonitorOutputWithContext(context.Background())
}

func (i *ConnectionMonitor) ToConnectionMonitorOutputWithContext(ctx context.Context) ConnectionMonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMonitorOutput)
}

type ConnectionMonitorOutput struct{ *pulumi.OutputState }

func (ConnectionMonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionMonitor)(nil)).Elem()
}

func (o ConnectionMonitorOutput) ToConnectionMonitorOutput() ConnectionMonitorOutput {
	return o
}

func (o ConnectionMonitorOutput) ToConnectionMonitorOutputWithContext(ctx context.Context) ConnectionMonitorOutput {
	return o
}

func (o ConnectionMonitorOutput) AutoStart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.BoolPtrOutput { return v.AutoStart }).(pulumi.BoolPtrOutput)
}

func (o ConnectionMonitorOutput) ConnectionMonitorType() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.StringOutput { return v.ConnectionMonitorType }).(pulumi.StringOutput)
}

func (o ConnectionMonitorOutput) Destination() ConnectionMonitorDestinationResponsePtrOutput {
	return o.ApplyT(func(v *ConnectionMonitor) ConnectionMonitorDestinationResponsePtrOutput { return v.Destination }).(ConnectionMonitorDestinationResponsePtrOutput)
}

func (o ConnectionMonitorOutput) Endpoints() ConnectionMonitorEndpointResponseArrayOutput {
	return o.ApplyT(func(v *ConnectionMonitor) ConnectionMonitorEndpointResponseArrayOutput { return v.Endpoints }).(ConnectionMonitorEndpointResponseArrayOutput)
}

func (o ConnectionMonitorOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ConnectionMonitorOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ConnectionMonitorOutput) MonitoringIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.IntPtrOutput { return v.MonitoringIntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o ConnectionMonitorOutput) MonitoringStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.StringOutput { return v.MonitoringStatus }).(pulumi.StringOutput)
}

func (o ConnectionMonitorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConnectionMonitorOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o ConnectionMonitorOutput) Outputs() ConnectionMonitorOutputResponseArrayOutput {
	return o.ApplyT(func(v *ConnectionMonitor) ConnectionMonitorOutputResponseArrayOutput { return v.Outputs }).(ConnectionMonitorOutputResponseArrayOutput)
}

func (o ConnectionMonitorOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ConnectionMonitorOutput) Source() ConnectionMonitorSourceResponsePtrOutput {
	return o.ApplyT(func(v *ConnectionMonitor) ConnectionMonitorSourceResponsePtrOutput { return v.Source }).(ConnectionMonitorSourceResponsePtrOutput)
}

func (o ConnectionMonitorOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

func (o ConnectionMonitorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ConnectionMonitorOutput) TestConfigurations() ConnectionMonitorTestConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *ConnectionMonitor) ConnectionMonitorTestConfigurationResponseArrayOutput {
		return v.TestConfigurations
	}).(ConnectionMonitorTestConfigurationResponseArrayOutput)
}

func (o ConnectionMonitorOutput) TestGroups() ConnectionMonitorTestGroupResponseArrayOutput {
	return o.ApplyT(func(v *ConnectionMonitor) ConnectionMonitorTestGroupResponseArrayOutput { return v.TestGroups }).(ConnectionMonitorTestGroupResponseArrayOutput)
}

func (o ConnectionMonitorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectionMonitorOutput{})
}

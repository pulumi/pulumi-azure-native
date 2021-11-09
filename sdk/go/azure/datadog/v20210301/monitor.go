


package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Monitor struct {
	pulumi.CustomResourceState

	Identity   IdentityPropertiesResponsePtrOutput `pulumi:"identity"`
	Location   pulumi.StringOutput                 `pulumi:"location"`
	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties MonitorPropertiesResponseOutput     `pulumi:"properties"`
	Sku        ResourceSkuResponsePtrOutput        `pulumi:"sku"`
	SystemData SystemDataResponseOutput            `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput              `pulumi:"tags"`
	Type       pulumi.StringOutput                 `pulumi:"type"`
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
			Type: pulumi.String("azure-nextgen:datadog/v20210301:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:datadog:Monitor"),
		},
		{
			Type: pulumi.String("azure-nextgen:datadog:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:datadog/v20200201preview:Monitor"),
		},
		{
			Type: pulumi.String("azure-nextgen:datadog/v20200201preview:Monitor"),
		},
	})
	opts = append(opts, aliases)
	var resource Monitor
	err := ctx.RegisterResource("azure-native:datadog/v20210301:Monitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitorState, opts ...pulumi.ResourceOption) (*Monitor, error) {
	var resource Monitor
	err := ctx.ReadResource("azure-native:datadog/v20210301:Monitor", name, id, state, &resource, opts...)
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
	Identity          *IdentityProperties `pulumi:"identity"`
	Location          *string             `pulumi:"location"`
	MonitorName       *string             `pulumi:"monitorName"`
	Properties        *MonitorProperties  `pulumi:"properties"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	Sku               *ResourceSku        `pulumi:"sku"`
	Tags              map[string]string   `pulumi:"tags"`
}


type MonitorArgs struct {
	Identity          IdentityPropertiesPtrInput
	Location          pulumi.StringPtrInput
	MonitorName       pulumi.StringPtrInput
	Properties        MonitorPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               ResourceSkuPtrInput
	Tags              pulumi.StringMapInput
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
	return reflect.TypeOf((*Monitor)(nil))
}

func (i *Monitor) ToMonitorOutput() MonitorOutput {
	return i.ToMonitorOutputWithContext(context.Background())
}

func (i *Monitor) ToMonitorOutputWithContext(ctx context.Context) MonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorOutput)
}

type MonitorOutput struct{ *pulumi.OutputState }

func (MonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Monitor)(nil))
}

func (o MonitorOutput) ToMonitorOutput() MonitorOutput {
	return o
}

func (o MonitorOutput) ToMonitorOutputWithContext(ctx context.Context) MonitorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MonitorOutput{})
}

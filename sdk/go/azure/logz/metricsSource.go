


package logz

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MetricsSource struct {
	pulumi.CustomResourceState

	Identity   IdentityPropertiesResponsePtrOutput `pulumi:"identity"`
	Location   pulumi.StringOutput                 `pulumi:"location"`
	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties MonitorPropertiesResponseOutput     `pulumi:"properties"`
	SystemData SystemDataResponseOutput            `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput              `pulumi:"tags"`
	Type       pulumi.StringOutput                 `pulumi:"type"`
}


func NewMetricsSource(ctx *pulumi.Context,
	name string, args *MetricsSourceArgs, opts ...pulumi.ResourceOption) (*MetricsSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MonitorName == nil {
		return nil, errors.New("invalid value for required argument 'MonitorName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logz/v20220101preview:MetricsSource"),
		},
	})
	opts = append(opts, aliases)
	var resource MetricsSource
	err := ctx.RegisterResource("azure-native:logz:MetricsSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMetricsSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetricsSourceState, opts ...pulumi.ResourceOption) (*MetricsSource, error) {
	var resource MetricsSource
	err := ctx.ReadResource("azure-native:logz:MetricsSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type metricsSourceState struct {
}

type MetricsSourceState struct {
}

func (MetricsSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*metricsSourceState)(nil)).Elem()
}

type metricsSourceArgs struct {
	Identity          *IdentityProperties `pulumi:"identity"`
	Location          *string             `pulumi:"location"`
	MetricsSourceName *string             `pulumi:"metricsSourceName"`
	MonitorName       string              `pulumi:"monitorName"`
	Properties        *MonitorProperties  `pulumi:"properties"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	Tags              map[string]string   `pulumi:"tags"`
}


type MetricsSourceArgs struct {
	Identity          IdentityPropertiesPtrInput
	Location          pulumi.StringPtrInput
	MetricsSourceName pulumi.StringPtrInput
	MonitorName       pulumi.StringInput
	Properties        MonitorPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (MetricsSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metricsSourceArgs)(nil)).Elem()
}

type MetricsSourceInput interface {
	pulumi.Input

	ToMetricsSourceOutput() MetricsSourceOutput
	ToMetricsSourceOutputWithContext(ctx context.Context) MetricsSourceOutput
}

func (*MetricsSource) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricsSource)(nil)).Elem()
}

func (i *MetricsSource) ToMetricsSourceOutput() MetricsSourceOutput {
	return i.ToMetricsSourceOutputWithContext(context.Background())
}

func (i *MetricsSource) ToMetricsSourceOutputWithContext(ctx context.Context) MetricsSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricsSourceOutput)
}

type MetricsSourceOutput struct{ *pulumi.OutputState }

func (MetricsSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricsSource)(nil)).Elem()
}

func (o MetricsSourceOutput) ToMetricsSourceOutput() MetricsSourceOutput {
	return o
}

func (o MetricsSourceOutput) ToMetricsSourceOutputWithContext(ctx context.Context) MetricsSourceOutput {
	return o
}

func (o MetricsSourceOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *MetricsSource) IdentityPropertiesResponsePtrOutput { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o MetricsSourceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricsSource) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o MetricsSourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricsSource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MetricsSourceOutput) Properties() MonitorPropertiesResponseOutput {
	return o.ApplyT(func(v *MetricsSource) MonitorPropertiesResponseOutput { return v.Properties }).(MonitorPropertiesResponseOutput)
}

func (o MetricsSourceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MetricsSource) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MetricsSourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MetricsSource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o MetricsSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricsSource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MetricsSourceOutput{})
}

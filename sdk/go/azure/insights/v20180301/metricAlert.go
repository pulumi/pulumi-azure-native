


package v20180301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MetricAlert struct {
	pulumi.CustomResourceState

	Actions              MetricAlertActionResponseArrayOutput `pulumi:"actions"`
	AutoMitigate         pulumi.BoolPtrOutput                 `pulumi:"autoMitigate"`
	Criteria             pulumi.AnyOutput                     `pulumi:"criteria"`
	Description          pulumi.StringPtrOutput               `pulumi:"description"`
	Enabled              pulumi.BoolOutput                    `pulumi:"enabled"`
	EvaluationFrequency  pulumi.StringOutput                  `pulumi:"evaluationFrequency"`
	IsMigrated           pulumi.BoolOutput                    `pulumi:"isMigrated"`
	LastUpdatedTime      pulumi.StringOutput                  `pulumi:"lastUpdatedTime"`
	Location             pulumi.StringOutput                  `pulumi:"location"`
	Name                 pulumi.StringOutput                  `pulumi:"name"`
	Scopes               pulumi.StringArrayOutput             `pulumi:"scopes"`
	Severity             pulumi.IntOutput                     `pulumi:"severity"`
	Tags                 pulumi.StringMapOutput               `pulumi:"tags"`
	TargetResourceRegion pulumi.StringPtrOutput               `pulumi:"targetResourceRegion"`
	TargetResourceType   pulumi.StringPtrOutput               `pulumi:"targetResourceType"`
	Type                 pulumi.StringOutput                  `pulumi:"type"`
	WindowSize           pulumi.StringOutput                  `pulumi:"windowSize"`
}


func NewMetricAlert(ctx *pulumi.Context,
	name string, args *MetricAlertArgs, opts ...pulumi.ResourceOption) (*MetricAlert, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Criteria == nil {
		return nil, errors.New("invalid value for required argument 'Criteria'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.EvaluationFrequency == nil {
		return nil, errors.New("invalid value for required argument 'EvaluationFrequency'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Scopes == nil {
		return nil, errors.New("invalid value for required argument 'Scopes'")
	}
	if args.Severity == nil {
		return nil, errors.New("invalid value for required argument 'Severity'")
	}
	if args.WindowSize == nil {
		return nil, errors.New("invalid value for required argument 'WindowSize'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:MetricAlert"),
		},
	})
	opts = append(opts, aliases)
	var resource MetricAlert
	err := ctx.RegisterResource("azure-native:insights/v20180301:MetricAlert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMetricAlert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetricAlertState, opts ...pulumi.ResourceOption) (*MetricAlert, error) {
	var resource MetricAlert
	err := ctx.ReadResource("azure-native:insights/v20180301:MetricAlert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type metricAlertState struct {
}

type MetricAlertState struct {
}

func (MetricAlertState) ElementType() reflect.Type {
	return reflect.TypeOf((*metricAlertState)(nil)).Elem()
}

type metricAlertArgs struct {
	Actions              []MetricAlertAction `pulumi:"actions"`
	AutoMitigate         *bool               `pulumi:"autoMitigate"`
	Criteria             interface{}         `pulumi:"criteria"`
	Description          *string             `pulumi:"description"`
	Enabled              bool                `pulumi:"enabled"`
	EvaluationFrequency  string              `pulumi:"evaluationFrequency"`
	Location             *string             `pulumi:"location"`
	ResourceGroupName    string              `pulumi:"resourceGroupName"`
	RuleName             *string             `pulumi:"ruleName"`
	Scopes               []string            `pulumi:"scopes"`
	Severity             int                 `pulumi:"severity"`
	Tags                 map[string]string   `pulumi:"tags"`
	TargetResourceRegion *string             `pulumi:"targetResourceRegion"`
	TargetResourceType   *string             `pulumi:"targetResourceType"`
	WindowSize           string              `pulumi:"windowSize"`
}


type MetricAlertArgs struct {
	Actions              MetricAlertActionArrayInput
	AutoMitigate         pulumi.BoolPtrInput
	Criteria             pulumi.Input
	Description          pulumi.StringPtrInput
	Enabled              pulumi.BoolInput
	EvaluationFrequency  pulumi.StringInput
	Location             pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	RuleName             pulumi.StringPtrInput
	Scopes               pulumi.StringArrayInput
	Severity             pulumi.IntInput
	Tags                 pulumi.StringMapInput
	TargetResourceRegion pulumi.StringPtrInput
	TargetResourceType   pulumi.StringPtrInput
	WindowSize           pulumi.StringInput
}

func (MetricAlertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metricAlertArgs)(nil)).Elem()
}

type MetricAlertInput interface {
	pulumi.Input

	ToMetricAlertOutput() MetricAlertOutput
	ToMetricAlertOutputWithContext(ctx context.Context) MetricAlertOutput
}

func (*MetricAlert) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricAlert)(nil)).Elem()
}

func (i *MetricAlert) ToMetricAlertOutput() MetricAlertOutput {
	return i.ToMetricAlertOutputWithContext(context.Background())
}

func (i *MetricAlert) ToMetricAlertOutputWithContext(ctx context.Context) MetricAlertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricAlertOutput)
}

type MetricAlertOutput struct{ *pulumi.OutputState }

func (MetricAlertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricAlert)(nil)).Elem()
}

func (o MetricAlertOutput) ToMetricAlertOutput() MetricAlertOutput {
	return o
}

func (o MetricAlertOutput) ToMetricAlertOutputWithContext(ctx context.Context) MetricAlertOutput {
	return o
}

func (o MetricAlertOutput) Actions() MetricAlertActionResponseArrayOutput {
	return o.ApplyT(func(v *MetricAlert) MetricAlertActionResponseArrayOutput { return v.Actions }).(MetricAlertActionResponseArrayOutput)
}

func (o MetricAlertOutput) AutoMitigate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MetricAlert) pulumi.BoolPtrOutput { return v.AutoMitigate }).(pulumi.BoolPtrOutput)
}

func (o MetricAlertOutput) Criteria() pulumi.AnyOutput {
	return o.ApplyT(func(v *MetricAlert) pulumi.AnyOutput { return v.Criteria }).(pulumi.AnyOutput)
}

func (o MetricAlertOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetricAlert) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o MetricAlertOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *MetricAlert) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o MetricAlertOutput) EvaluationFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricAlert) pulumi.StringOutput { return v.EvaluationFrequency }).(pulumi.StringOutput)
}

func (o MetricAlertOutput) IsMigrated() pulumi.BoolOutput {
	return o.ApplyT(func(v *MetricAlert) pulumi.BoolOutput { return v.IsMigrated }).(pulumi.BoolOutput)
}

func (o MetricAlertOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricAlert) pulumi.StringOutput { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

func (o MetricAlertOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricAlert) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o MetricAlertOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricAlert) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MetricAlertOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MetricAlert) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

func (o MetricAlertOutput) Severity() pulumi.IntOutput {
	return o.ApplyT(func(v *MetricAlert) pulumi.IntOutput { return v.Severity }).(pulumi.IntOutput)
}

func (o MetricAlertOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MetricAlert) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o MetricAlertOutput) TargetResourceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetricAlert) pulumi.StringPtrOutput { return v.TargetResourceRegion }).(pulumi.StringPtrOutput)
}

func (o MetricAlertOutput) TargetResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetricAlert) pulumi.StringPtrOutput { return v.TargetResourceType }).(pulumi.StringPtrOutput)
}

func (o MetricAlertOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricAlert) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o MetricAlertOutput) WindowSize() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricAlert) pulumi.StringOutput { return v.WindowSize }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MetricAlertOutput{})
}

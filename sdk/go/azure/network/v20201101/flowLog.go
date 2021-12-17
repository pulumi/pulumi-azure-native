


package v20201101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FlowLog struct {
	pulumi.CustomResourceState

	Enabled                    pulumi.BoolPtrOutput                        `pulumi:"enabled"`
	Etag                       pulumi.StringOutput                         `pulumi:"etag"`
	FlowAnalyticsConfiguration TrafficAnalyticsPropertiesResponsePtrOutput `pulumi:"flowAnalyticsConfiguration"`
	Format                     FlowLogFormatParametersResponsePtrOutput    `pulumi:"format"`
	Location                   pulumi.StringPtrOutput                      `pulumi:"location"`
	Name                       pulumi.StringOutput                         `pulumi:"name"`
	ProvisioningState          pulumi.StringOutput                         `pulumi:"provisioningState"`
	RetentionPolicy            RetentionPolicyParametersResponsePtrOutput  `pulumi:"retentionPolicy"`
	StorageId                  pulumi.StringOutput                         `pulumi:"storageId"`
	Tags                       pulumi.StringMapOutput                      `pulumi:"tags"`
	TargetResourceGuid         pulumi.StringOutput                         `pulumi:"targetResourceGuid"`
	TargetResourceId           pulumi.StringOutput                         `pulumi:"targetResourceId"`
	Type                       pulumi.StringOutput                         `pulumi:"type"`
}


func NewFlowLog(ctx *pulumi.Context,
	name string, args *FlowLogArgs, opts ...pulumi.ResourceOption) (*FlowLog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkWatcherName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkWatcherName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageId == nil {
		return nil, errors.New("invalid value for required argument 'StorageId'")
	}
	if args.TargetResourceId == nil {
		return nil, errors.New("invalid value for required argument 'TargetResourceId'")
	}
	formatApplier := func(v FlowLogFormatParameters) *FlowLogFormatParameters { return v.Defaults() }
	if args.Format != nil {
		args.Format = args.Format.ToFlowLogFormatParametersPtrOutput().Elem().ApplyT(formatApplier).(FlowLogFormatParametersPtrOutput)
	}
	retentionPolicyApplier := func(v RetentionPolicyParameters) *RetentionPolicyParameters { return v.Defaults() }
	if args.RetentionPolicy != nil {
		args.RetentionPolicy = args.RetentionPolicy.ToRetentionPolicyParametersPtrOutput().Elem().ApplyT(retentionPolicyApplier).(RetentionPolicyParametersPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:FlowLog"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:FlowLog"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:FlowLog"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:FlowLog"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:FlowLog"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:FlowLog"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:FlowLog"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:FlowLog"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:FlowLog"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:FlowLog"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:FlowLog"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:FlowLog"),
		},
	})
	opts = append(opts, aliases)
	var resource FlowLog
	err := ctx.RegisterResource("azure-native:network/v20201101:FlowLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFlowLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowLogState, opts ...pulumi.ResourceOption) (*FlowLog, error) {
	var resource FlowLog
	err := ctx.ReadResource("azure-native:network/v20201101:FlowLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type flowLogState struct {
}

type FlowLogState struct {
}

func (FlowLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowLogState)(nil)).Elem()
}

type flowLogArgs struct {
	Enabled                    *bool                       `pulumi:"enabled"`
	FlowAnalyticsConfiguration *TrafficAnalyticsProperties `pulumi:"flowAnalyticsConfiguration"`
	FlowLogName                *string                     `pulumi:"flowLogName"`
	Format                     *FlowLogFormatParameters    `pulumi:"format"`
	Id                         *string                     `pulumi:"id"`
	Location                   *string                     `pulumi:"location"`
	NetworkWatcherName         string                      `pulumi:"networkWatcherName"`
	ResourceGroupName          string                      `pulumi:"resourceGroupName"`
	RetentionPolicy            *RetentionPolicyParameters  `pulumi:"retentionPolicy"`
	StorageId                  string                      `pulumi:"storageId"`
	Tags                       map[string]string           `pulumi:"tags"`
	TargetResourceId           string                      `pulumi:"targetResourceId"`
}


type FlowLogArgs struct {
	Enabled                    pulumi.BoolPtrInput
	FlowAnalyticsConfiguration TrafficAnalyticsPropertiesPtrInput
	FlowLogName                pulumi.StringPtrInput
	Format                     FlowLogFormatParametersPtrInput
	Id                         pulumi.StringPtrInput
	Location                   pulumi.StringPtrInput
	NetworkWatcherName         pulumi.StringInput
	ResourceGroupName          pulumi.StringInput
	RetentionPolicy            RetentionPolicyParametersPtrInput
	StorageId                  pulumi.StringInput
	Tags                       pulumi.StringMapInput
	TargetResourceId           pulumi.StringInput
}

func (FlowLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowLogArgs)(nil)).Elem()
}

type FlowLogInput interface {
	pulumi.Input

	ToFlowLogOutput() FlowLogOutput
	ToFlowLogOutputWithContext(ctx context.Context) FlowLogOutput
}

func (*FlowLog) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowLog)(nil)).Elem()
}

func (i *FlowLog) ToFlowLogOutput() FlowLogOutput {
	return i.ToFlowLogOutputWithContext(context.Background())
}

func (i *FlowLog) ToFlowLogOutputWithContext(ctx context.Context) FlowLogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowLogOutput)
}

type FlowLogOutput struct{ *pulumi.OutputState }

func (FlowLogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowLog)(nil)).Elem()
}

func (o FlowLogOutput) ToFlowLogOutput() FlowLogOutput {
	return o
}

func (o FlowLogOutput) ToFlowLogOutputWithContext(ctx context.Context) FlowLogOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FlowLogOutput{})
}

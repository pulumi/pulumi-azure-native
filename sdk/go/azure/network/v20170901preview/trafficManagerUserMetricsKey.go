


package v20170901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TrafficManagerUserMetricsKey struct {
	pulumi.CustomResourceState

	Key  pulumi.StringPtrOutput `pulumi:"key"`
	Name pulumi.StringOutput    `pulumi:"name"`
	Type pulumi.StringOutput    `pulumi:"type"`
}


func NewTrafficManagerUserMetricsKey(ctx *pulumi.Context,
	name string, args *TrafficManagerUserMetricsKeyArgs, opts ...pulumi.ResourceOption) (*TrafficManagerUserMetricsKey, error) {
	if args == nil {
		args = &TrafficManagerUserMetricsKeyArgs{}
	}

	var resource TrafficManagerUserMetricsKey
	err := ctx.RegisterResource("azure-native:network/v20170901preview:TrafficManagerUserMetricsKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTrafficManagerUserMetricsKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficManagerUserMetricsKeyState, opts ...pulumi.ResourceOption) (*TrafficManagerUserMetricsKey, error) {
	var resource TrafficManagerUserMetricsKey
	err := ctx.ReadResource("azure-native:network/v20170901preview:TrafficManagerUserMetricsKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type trafficManagerUserMetricsKeyState struct {
}

type TrafficManagerUserMetricsKeyState struct {
}

func (TrafficManagerUserMetricsKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficManagerUserMetricsKeyState)(nil)).Elem()
}

type trafficManagerUserMetricsKeyArgs struct {
}


type TrafficManagerUserMetricsKeyArgs struct {
}

func (TrafficManagerUserMetricsKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficManagerUserMetricsKeyArgs)(nil)).Elem()
}

type TrafficManagerUserMetricsKeyInput interface {
	pulumi.Input

	ToTrafficManagerUserMetricsKeyOutput() TrafficManagerUserMetricsKeyOutput
	ToTrafficManagerUserMetricsKeyOutputWithContext(ctx context.Context) TrafficManagerUserMetricsKeyOutput
}

func (*TrafficManagerUserMetricsKey) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficManagerUserMetricsKey)(nil))
}

func (i *TrafficManagerUserMetricsKey) ToTrafficManagerUserMetricsKeyOutput() TrafficManagerUserMetricsKeyOutput {
	return i.ToTrafficManagerUserMetricsKeyOutputWithContext(context.Background())
}

func (i *TrafficManagerUserMetricsKey) ToTrafficManagerUserMetricsKeyOutputWithContext(ctx context.Context) TrafficManagerUserMetricsKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficManagerUserMetricsKeyOutput)
}

type TrafficManagerUserMetricsKeyOutput struct{ *pulumi.OutputState }

func (TrafficManagerUserMetricsKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficManagerUserMetricsKey)(nil))
}

func (o TrafficManagerUserMetricsKeyOutput) ToTrafficManagerUserMetricsKeyOutput() TrafficManagerUserMetricsKeyOutput {
	return o
}

func (o TrafficManagerUserMetricsKeyOutput) ToTrafficManagerUserMetricsKeyOutputWithContext(ctx context.Context) TrafficManagerUserMetricsKeyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TrafficManagerUserMetricsKeyOutput{})
}

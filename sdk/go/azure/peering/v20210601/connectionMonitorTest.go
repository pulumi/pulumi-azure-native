


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectionMonitorTest struct {
	pulumi.CustomResourceState

	Destination        pulumi.StringPtrOutput   `pulumi:"destination"`
	DestinationPort    pulumi.IntPtrOutput      `pulumi:"destinationPort"`
	IsTestSuccessful   pulumi.BoolOutput        `pulumi:"isTestSuccessful"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	Path               pulumi.StringArrayOutput `pulumi:"path"`
	ProvisioningState  pulumi.StringOutput      `pulumi:"provisioningState"`
	SourceAgent        pulumi.StringPtrOutput   `pulumi:"sourceAgent"`
	TestFrequencyInSec pulumi.IntPtrOutput      `pulumi:"testFrequencyInSec"`
	Type               pulumi.StringOutput      `pulumi:"type"`
}


func NewConnectionMonitorTest(ctx *pulumi.Context,
	name string, args *ConnectionMonitorTestArgs, opts ...pulumi.ResourceOption) (*ConnectionMonitorTest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PeeringServiceName == nil {
		return nil, errors.New("invalid value for required argument 'PeeringServiceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:peering:ConnectionMonitorTest"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20220101:ConnectionMonitorTest"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20220601:ConnectionMonitorTest"),
		},
	})
	opts = append(opts, aliases)
	var resource ConnectionMonitorTest
	err := ctx.RegisterResource("azure-native:peering/v20210601:ConnectionMonitorTest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnectionMonitorTest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionMonitorTestState, opts ...pulumi.ResourceOption) (*ConnectionMonitorTest, error) {
	var resource ConnectionMonitorTest
	err := ctx.ReadResource("azure-native:peering/v20210601:ConnectionMonitorTest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type connectionMonitorTestState struct {
}

type ConnectionMonitorTestState struct {
}

func (ConnectionMonitorTestState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionMonitorTestState)(nil)).Elem()
}

type connectionMonitorTestArgs struct {
	ConnectionMonitorTestName *string `pulumi:"connectionMonitorTestName"`
	Destination               *string `pulumi:"destination"`
	DestinationPort           *int    `pulumi:"destinationPort"`
	PeeringServiceName        string  `pulumi:"peeringServiceName"`
	ResourceGroupName         string  `pulumi:"resourceGroupName"`
	SourceAgent               *string `pulumi:"sourceAgent"`
	TestFrequencyInSec        *int    `pulumi:"testFrequencyInSec"`
}


type ConnectionMonitorTestArgs struct {
	ConnectionMonitorTestName pulumi.StringPtrInput
	Destination               pulumi.StringPtrInput
	DestinationPort           pulumi.IntPtrInput
	PeeringServiceName        pulumi.StringInput
	ResourceGroupName         pulumi.StringInput
	SourceAgent               pulumi.StringPtrInput
	TestFrequencyInSec        pulumi.IntPtrInput
}

func (ConnectionMonitorTestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionMonitorTestArgs)(nil)).Elem()
}

type ConnectionMonitorTestInput interface {
	pulumi.Input

	ToConnectionMonitorTestOutput() ConnectionMonitorTestOutput
	ToConnectionMonitorTestOutputWithContext(ctx context.Context) ConnectionMonitorTestOutput
}

func (*ConnectionMonitorTest) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionMonitorTest)(nil)).Elem()
}

func (i *ConnectionMonitorTest) ToConnectionMonitorTestOutput() ConnectionMonitorTestOutput {
	return i.ToConnectionMonitorTestOutputWithContext(context.Background())
}

func (i *ConnectionMonitorTest) ToConnectionMonitorTestOutputWithContext(ctx context.Context) ConnectionMonitorTestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMonitorTestOutput)
}

type ConnectionMonitorTestOutput struct{ *pulumi.OutputState }

func (ConnectionMonitorTestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionMonitorTest)(nil)).Elem()
}

func (o ConnectionMonitorTestOutput) ToConnectionMonitorTestOutput() ConnectionMonitorTestOutput {
	return o
}

func (o ConnectionMonitorTestOutput) ToConnectionMonitorTestOutputWithContext(ctx context.Context) ConnectionMonitorTestOutput {
	return o
}

func (o ConnectionMonitorTestOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitorTest) pulumi.StringPtrOutput { return v.Destination }).(pulumi.StringPtrOutput)
}

func (o ConnectionMonitorTestOutput) DestinationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitorTest) pulumi.IntPtrOutput { return v.DestinationPort }).(pulumi.IntPtrOutput)
}

func (o ConnectionMonitorTestOutput) IsTestSuccessful() pulumi.BoolOutput {
	return o.ApplyT(func(v *ConnectionMonitorTest) pulumi.BoolOutput { return v.IsTestSuccessful }).(pulumi.BoolOutput)
}

func (o ConnectionMonitorTestOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionMonitorTest) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConnectionMonitorTestOutput) Path() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ConnectionMonitorTest) pulumi.StringArrayOutput { return v.Path }).(pulumi.StringArrayOutput)
}

func (o ConnectionMonitorTestOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionMonitorTest) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ConnectionMonitorTestOutput) SourceAgent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitorTest) pulumi.StringPtrOutput { return v.SourceAgent }).(pulumi.StringPtrOutput)
}

func (o ConnectionMonitorTestOutput) TestFrequencyInSec() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitorTest) pulumi.IntPtrOutput { return v.TestFrequencyInSec }).(pulumi.IntPtrOutput)
}

func (o ConnectionMonitorTestOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionMonitorTest) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectionMonitorTestOutput{})
}

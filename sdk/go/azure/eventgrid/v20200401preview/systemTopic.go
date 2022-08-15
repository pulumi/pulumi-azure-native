


package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemTopic struct {
	pulumi.CustomResourceState

	Location          pulumi.StringOutput      `pulumi:"location"`
	MetricResourceId  pulumi.StringOutput      `pulumi:"metricResourceId"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	Source            pulumi.StringPtrOutput   `pulumi:"source"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput   `pulumi:"tags"`
	TopicType         pulumi.StringPtrOutput   `pulumi:"topicType"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewSystemTopic(ctx *pulumi.Context,
	name string, args *SystemTopicArgs, opts ...pulumi.ResourceOption) (*SystemTopic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:SystemTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20201015preview:SystemTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20210601preview:SystemTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:SystemTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211201:SystemTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20220615:SystemTopic"),
		},
	})
	opts = append(opts, aliases)
	var resource SystemTopic
	err := ctx.RegisterResource("azure-native:eventgrid/v20200401preview:SystemTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSystemTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemTopicState, opts ...pulumi.ResourceOption) (*SystemTopic, error) {
	var resource SystemTopic
	err := ctx.ReadResource("azure-native:eventgrid/v20200401preview:SystemTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type systemTopicState struct {
}

type SystemTopicState struct {
}

func (SystemTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemTopicState)(nil)).Elem()
}

type systemTopicArgs struct {
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Source            *string           `pulumi:"source"`
	SystemTopicName   *string           `pulumi:"systemTopicName"`
	Tags              map[string]string `pulumi:"tags"`
	TopicType         *string           `pulumi:"topicType"`
}


type SystemTopicArgs struct {
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Source            pulumi.StringPtrInput
	SystemTopicName   pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	TopicType         pulumi.StringPtrInput
}

func (SystemTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemTopicArgs)(nil)).Elem()
}

type SystemTopicInput interface {
	pulumi.Input

	ToSystemTopicOutput() SystemTopicOutput
	ToSystemTopicOutputWithContext(ctx context.Context) SystemTopicOutput
}

func (*SystemTopic) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemTopic)(nil)).Elem()
}

func (i *SystemTopic) ToSystemTopicOutput() SystemTopicOutput {
	return i.ToSystemTopicOutputWithContext(context.Background())
}

func (i *SystemTopic) ToSystemTopicOutputWithContext(ctx context.Context) SystemTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemTopicOutput)
}

type SystemTopicOutput struct{ *pulumi.OutputState }

func (SystemTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemTopic)(nil)).Elem()
}

func (o SystemTopicOutput) ToSystemTopicOutput() SystemTopicOutput {
	return o
}

func (o SystemTopicOutput) ToSystemTopicOutputWithContext(ctx context.Context) SystemTopicOutput {
	return o
}

func (o SystemTopicOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemTopic) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SystemTopicOutput) MetricResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemTopic) pulumi.StringOutput { return v.MetricResourceId }).(pulumi.StringOutput)
}

func (o SystemTopicOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemTopic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SystemTopicOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemTopic) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SystemTopicOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemTopic) pulumi.StringPtrOutput { return v.Source }).(pulumi.StringPtrOutput)
}

func (o SystemTopicOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemTopic) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SystemTopicOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SystemTopic) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SystemTopicOutput) TopicType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemTopic) pulumi.StringPtrOutput { return v.TopicType }).(pulumi.StringPtrOutput)
}

func (o SystemTopicOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemTopic) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemTopicOutput{})
}

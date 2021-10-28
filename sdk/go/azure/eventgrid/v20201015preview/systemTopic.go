


package v20201015preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemTopic struct {
	pulumi.CustomResourceState

	Identity          IdentityInfoResponsePtrOutput `pulumi:"identity"`
	Location          pulumi.StringOutput           `pulumi:"location"`
	MetricResourceId  pulumi.StringOutput           `pulumi:"metricResourceId"`
	Name              pulumi.StringOutput           `pulumi:"name"`
	ProvisioningState pulumi.StringOutput           `pulumi:"provisioningState"`
	Source            pulumi.StringPtrOutput        `pulumi:"source"`
	SystemData        SystemDataResponseOutput      `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput        `pulumi:"tags"`
	TopicType         pulumi.StringPtrOutput        `pulumi:"topicType"`
	Type              pulumi.StringOutput           `pulumi:"type"`
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
			Type: pulumi.String("azure-nextgen:eventgrid/v20201015preview:SystemTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid:SystemTopic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid:SystemTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200401preview:SystemTopic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20200401preview:SystemTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20210601preview:SystemTopic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20210601preview:SystemTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211201:SystemTopic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20211201:SystemTopic"),
		},
	})
	opts = append(opts, aliases)
	var resource SystemTopic
	err := ctx.RegisterResource("azure-native:eventgrid/v20201015preview:SystemTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSystemTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemTopicState, opts ...pulumi.ResourceOption) (*SystemTopic, error) {
	var resource SystemTopic
	err := ctx.ReadResource("azure-native:eventgrid/v20201015preview:SystemTopic", name, id, state, &resource, opts...)
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
	Identity          *IdentityInfo     `pulumi:"identity"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Source            *string           `pulumi:"source"`
	SystemTopicName   *string           `pulumi:"systemTopicName"`
	Tags              map[string]string `pulumi:"tags"`
	TopicType         *string           `pulumi:"topicType"`
}


type SystemTopicArgs struct {
	Identity          IdentityInfoPtrInput
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
	return reflect.TypeOf((*SystemTopic)(nil))
}

func (i *SystemTopic) ToSystemTopicOutput() SystemTopicOutput {
	return i.ToSystemTopicOutputWithContext(context.Background())
}

func (i *SystemTopic) ToSystemTopicOutputWithContext(ctx context.Context) SystemTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemTopicOutput)
}

type SystemTopicOutput struct{ *pulumi.OutputState }

func (SystemTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemTopic)(nil))
}

func (o SystemTopicOutput) ToSystemTopicOutput() SystemTopicOutput {
	return o
}

func (o SystemTopicOutput) ToSystemTopicOutputWithContext(ctx context.Context) SystemTopicOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemTopicInput)(nil)).Elem(), &SystemTopic{})
	pulumi.RegisterOutputType(SystemTopicOutput{})
}

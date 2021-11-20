


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DomainTopic struct {
	pulumi.CustomResourceState

	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewDomainTopic(ctx *pulumi.Context,
	name string, args *DomainTopicArgs, opts ...pulumi.ResourceOption) (*DomainTopic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:DomainTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190201preview:DomainTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190601:DomainTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200101preview:DomainTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200401preview:DomainTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200601:DomainTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20201015preview:DomainTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211201:DomainTopic"),
		},
	})
	opts = append(opts, aliases)
	var resource DomainTopic
	err := ctx.RegisterResource("azure-native:eventgrid/v20210601preview:DomainTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDomainTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainTopicState, opts ...pulumi.ResourceOption) (*DomainTopic, error) {
	var resource DomainTopic
	err := ctx.ReadResource("azure-native:eventgrid/v20210601preview:DomainTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type domainTopicState struct {
}

type DomainTopicState struct {
}

func (DomainTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainTopicState)(nil)).Elem()
}

type domainTopicArgs struct {
	DomainName        string  `pulumi:"domainName"`
	DomainTopicName   *string `pulumi:"domainTopicName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type DomainTopicArgs struct {
	DomainName        pulumi.StringInput
	DomainTopicName   pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (DomainTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainTopicArgs)(nil)).Elem()
}

type DomainTopicInput interface {
	pulumi.Input

	ToDomainTopicOutput() DomainTopicOutput
	ToDomainTopicOutputWithContext(ctx context.Context) DomainTopicOutput
}

func (*DomainTopic) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainTopic)(nil))
}

func (i *DomainTopic) ToDomainTopicOutput() DomainTopicOutput {
	return i.ToDomainTopicOutputWithContext(context.Background())
}

func (i *DomainTopic) ToDomainTopicOutputWithContext(ctx context.Context) DomainTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainTopicOutput)
}

type DomainTopicOutput struct{ *pulumi.OutputState }

func (DomainTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainTopic)(nil))
}

func (o DomainTopicOutput) ToDomainTopicOutput() DomainTopicOutput {
	return o
}

func (o DomainTopicOutput) ToDomainTopicOutputWithContext(ctx context.Context) DomainTopicOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DomainTopicOutput{})
}

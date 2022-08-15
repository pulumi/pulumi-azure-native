


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type QueueServiceProperties struct {
	pulumi.CustomResourceState

	Cors CorsRulesResponsePtrOutput `pulumi:"cors"`
	Name pulumi.StringOutput        `pulumi:"name"`
	Type pulumi.StringOutput        `pulumi:"type"`
}


func NewQueueServiceProperties(ctx *pulumi.Context,
	name string, args *QueueServicePropertiesArgs, opts ...pulumi.ResourceOption) (*QueueServiceProperties, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage:QueueServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:QueueServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:QueueServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:QueueServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:QueueServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:QueueServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:QueueServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:QueueServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:QueueServiceProperties"),
		},
	})
	opts = append(opts, aliases)
	var resource QueueServiceProperties
	err := ctx.RegisterResource("azure-native:storage/v20210601:QueueServiceProperties", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetQueueServiceProperties(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueServicePropertiesState, opts ...pulumi.ResourceOption) (*QueueServiceProperties, error) {
	var resource QueueServiceProperties
	err := ctx.ReadResource("azure-native:storage/v20210601:QueueServiceProperties", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type queueServicePropertiesState struct {
}

type QueueServicePropertiesState struct {
}

func (QueueServicePropertiesState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueServicePropertiesState)(nil)).Elem()
}

type queueServicePropertiesArgs struct {
	AccountName       string     `pulumi:"accountName"`
	Cors              *CorsRules `pulumi:"cors"`
	QueueServiceName  *string    `pulumi:"queueServiceName"`
	ResourceGroupName string     `pulumi:"resourceGroupName"`
}


type QueueServicePropertiesArgs struct {
	AccountName       pulumi.StringInput
	Cors              CorsRulesPtrInput
	QueueServiceName  pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (QueueServicePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueServicePropertiesArgs)(nil)).Elem()
}

type QueueServicePropertiesInput interface {
	pulumi.Input

	ToQueueServicePropertiesOutput() QueueServicePropertiesOutput
	ToQueueServicePropertiesOutputWithContext(ctx context.Context) QueueServicePropertiesOutput
}

func (*QueueServiceProperties) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueServiceProperties)(nil)).Elem()
}

func (i *QueueServiceProperties) ToQueueServicePropertiesOutput() QueueServicePropertiesOutput {
	return i.ToQueueServicePropertiesOutputWithContext(context.Background())
}

func (i *QueueServiceProperties) ToQueueServicePropertiesOutputWithContext(ctx context.Context) QueueServicePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueServicePropertiesOutput)
}

type QueueServicePropertiesOutput struct{ *pulumi.OutputState }

func (QueueServicePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueServiceProperties)(nil)).Elem()
}

func (o QueueServicePropertiesOutput) ToQueueServicePropertiesOutput() QueueServicePropertiesOutput {
	return o
}

func (o QueueServicePropertiesOutput) ToQueueServicePropertiesOutputWithContext(ctx context.Context) QueueServicePropertiesOutput {
	return o
}

func (o QueueServicePropertiesOutput) Cors() CorsRulesResponsePtrOutput {
	return o.ApplyT(func(v *QueueServiceProperties) CorsRulesResponsePtrOutput { return v.Cors }).(CorsRulesResponsePtrOutput)
}

func (o QueueServicePropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueServiceProperties) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o QueueServicePropertiesOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueServiceProperties) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(QueueServicePropertiesOutput{})
}

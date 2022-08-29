


package v20160601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Connection struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput                          `pulumi:"etag"`
	Location   pulumi.StringPtrOutput                          `pulumi:"location"`
	Name       pulumi.StringOutput                             `pulumi:"name"`
	Properties ApiConnectionDefinitionResponsePropertiesOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                          `pulumi:"tags"`
	Type       pulumi.StringOutput                             `pulumi:"type"`
}


func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:Connection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801preview:Connection"),
		},
	})
	opts = append(opts, aliases)
	var resource Connection
	err := ctx.RegisterResource("azure-native:web/v20160601:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("azure-native:web/v20160601:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type connectionState struct {
}

type ConnectionState struct {
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	ConnectionName    *string                            `pulumi:"connectionName"`
	Location          *string                            `pulumi:"location"`
	Properties        *ApiConnectionDefinitionProperties `pulumi:"properties"`
	ResourceGroupName string                             `pulumi:"resourceGroupName"`
	SubscriptionId    *string                            `pulumi:"subscriptionId"`
	Tags              map[string]string                  `pulumi:"tags"`
}


type ConnectionArgs struct {
	ConnectionName    pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Properties        ApiConnectionDefinitionPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	SubscriptionId    pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}

type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput
}

func (*Connection) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (i *Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

type ConnectionOutput struct{ *pulumi.OutputState }

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

func (o ConnectionOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ConnectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConnectionOutput) Properties() ApiConnectionDefinitionResponsePropertiesOutput {
	return o.ApplyT(func(v *Connection) ApiConnectionDefinitionResponsePropertiesOutput { return v.Properties }).(ApiConnectionDefinitionResponsePropertiesOutput)
}

func (o ConnectionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectionOutput{})
}

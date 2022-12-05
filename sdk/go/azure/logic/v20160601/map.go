


package v20160601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Map struct {
	pulumi.CustomResourceState

	ChangedTime      pulumi.StringOutput                                              `pulumi:"changedTime"`
	Content          pulumi.StringPtrOutput                                           `pulumi:"content"`
	ContentLink      ContentLinkResponseOutput                                        `pulumi:"contentLink"`
	ContentType      pulumi.StringPtrOutput                                           `pulumi:"contentType"`
	CreatedTime      pulumi.StringOutput                                              `pulumi:"createdTime"`
	Location         pulumi.StringPtrOutput                                           `pulumi:"location"`
	MapType          pulumi.StringOutput                                              `pulumi:"mapType"`
	Metadata         pulumi.AnyOutput                                                 `pulumi:"metadata"`
	Name             pulumi.StringOutput                                              `pulumi:"name"`
	ParametersSchema IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput `pulumi:"parametersSchema"`
	Tags             pulumi.StringMapOutput                                           `pulumi:"tags"`
	Type             pulumi.StringOutput                                              `pulumi:"type"`
}


func NewMap(ctx *pulumi.Context,
	name string, args *MapArgs, opts ...pulumi.ResourceOption) (*Map, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.MapType == nil {
		return nil, errors.New("invalid value for required argument 'MapType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic:Map"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20150801preview:Map"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:Map"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:Map"),
		},
	})
	opts = append(opts, aliases)
	var resource Map
	err := ctx.RegisterResource("azure-native:logic/v20160601:Map", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MapState, opts ...pulumi.ResourceOption) (*Map, error) {
	var resource Map
	err := ctx.ReadResource("azure-native:logic/v20160601:Map", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type mapState struct {
}

type MapState struct {
}

func (MapState) ElementType() reflect.Type {
	return reflect.TypeOf((*mapState)(nil)).Elem()
}

type mapArgs struct {
	Content                *string                                          `pulumi:"content"`
	ContentType            *string                                          `pulumi:"contentType"`
	IntegrationAccountName string                                           `pulumi:"integrationAccountName"`
	Location               *string                                          `pulumi:"location"`
	MapName                *string                                          `pulumi:"mapName"`
	MapType                MapType                                          `pulumi:"mapType"`
	Metadata               interface{}                                      `pulumi:"metadata"`
	ParametersSchema       *IntegrationAccountMapPropertiesParametersSchema `pulumi:"parametersSchema"`
	ResourceGroupName      string                                           `pulumi:"resourceGroupName"`
	Tags                   map[string]string                                `pulumi:"tags"`
}


type MapArgs struct {
	Content                pulumi.StringPtrInput
	ContentType            pulumi.StringPtrInput
	IntegrationAccountName pulumi.StringInput
	Location               pulumi.StringPtrInput
	MapName                pulumi.StringPtrInput
	MapType                MapTypeInput
	Metadata               pulumi.Input
	ParametersSchema       IntegrationAccountMapPropertiesParametersSchemaPtrInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
}

func (MapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mapArgs)(nil)).Elem()
}

type MapInput interface {
	pulumi.Input

	ToMapOutput() MapOutput
	ToMapOutputWithContext(ctx context.Context) MapOutput
}

func (*Map) ElementType() reflect.Type {
	return reflect.TypeOf((**Map)(nil)).Elem()
}

func (i *Map) ToMapOutput() MapOutput {
	return i.ToMapOutputWithContext(context.Background())
}

func (i *Map) ToMapOutputWithContext(ctx context.Context) MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MapOutput)
}

type MapOutput struct{ *pulumi.OutputState }

func (MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Map)(nil)).Elem()
}

func (o MapOutput) ToMapOutput() MapOutput {
	return o
}

func (o MapOutput) ToMapOutputWithContext(ctx context.Context) MapOutput {
	return o
}

func (o MapOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Map) pulumi.StringOutput { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o MapOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Map) pulumi.StringPtrOutput { return v.Content }).(pulumi.StringPtrOutput)
}

func (o MapOutput) ContentLink() ContentLinkResponseOutput {
	return o.ApplyT(func(v *Map) ContentLinkResponseOutput { return v.ContentLink }).(ContentLinkResponseOutput)
}

func (o MapOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Map) pulumi.StringPtrOutput { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o MapOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Map) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o MapOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Map) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o MapOutput) MapType() pulumi.StringOutput {
	return o.ApplyT(func(v *Map) pulumi.StringOutput { return v.MapType }).(pulumi.StringOutput)
}

func (o MapOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *Map) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

func (o MapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Map) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MapOutput) ParametersSchema() IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput {
	return o.ApplyT(func(v *Map) IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput {
		return v.ParametersSchema
	}).(IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput)
}

func (o MapOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Map) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o MapOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Map) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MapOutput{})
}

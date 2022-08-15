


package v20150801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationAccountMap struct {
	pulumi.CustomResourceState

	ChangedTime pulumi.StringOutput                         `pulumi:"changedTime"`
	Content     pulumi.AnyOutput                            `pulumi:"content"`
	ContentLink IntegrationAccountContentLinkResponseOutput `pulumi:"contentLink"`
	ContentType pulumi.StringPtrOutput                      `pulumi:"contentType"`
	CreatedTime pulumi.StringOutput                         `pulumi:"createdTime"`
	Location    pulumi.StringPtrOutput                      `pulumi:"location"`
	MapType     pulumi.StringPtrOutput                      `pulumi:"mapType"`
	Metadata    pulumi.AnyOutput                            `pulumi:"metadata"`
	Name        pulumi.StringPtrOutput                      `pulumi:"name"`
	Tags        pulumi.StringMapOutput                      `pulumi:"tags"`
	Type        pulumi.StringPtrOutput                      `pulumi:"type"`
}


func NewIntegrationAccountMap(ctx *pulumi.Context,
	name string, args *IntegrationAccountMapArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountMap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic:IntegrationAccountMap"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:IntegrationAccountMap"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:IntegrationAccountMap"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:IntegrationAccountMap"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationAccountMap
	err := ctx.RegisterResource("azure-native:logic/v20150801preview:IntegrationAccountMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIntegrationAccountMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountMapState, opts ...pulumi.ResourceOption) (*IntegrationAccountMap, error) {
	var resource IntegrationAccountMap
	err := ctx.ReadResource("azure-native:logic/v20150801preview:IntegrationAccountMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type integrationAccountMapState struct {
}

type IntegrationAccountMapState struct {
}

func (IntegrationAccountMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountMapState)(nil)).Elem()
}

type integrationAccountMapArgs struct {
	Content                interface{}       `pulumi:"content"`
	ContentType            *string           `pulumi:"contentType"`
	Id                     *string           `pulumi:"id"`
	IntegrationAccountName string            `pulumi:"integrationAccountName"`
	Location               *string           `pulumi:"location"`
	MapName                *string           `pulumi:"mapName"`
	MapType                *MapType          `pulumi:"mapType"`
	Metadata               interface{}       `pulumi:"metadata"`
	Name                   *string           `pulumi:"name"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	Tags                   map[string]string `pulumi:"tags"`
	Type                   *string           `pulumi:"type"`
}


type IntegrationAccountMapArgs struct {
	Content                pulumi.Input
	ContentType            pulumi.StringPtrInput
	Id                     pulumi.StringPtrInput
	IntegrationAccountName pulumi.StringInput
	Location               pulumi.StringPtrInput
	MapName                pulumi.StringPtrInput
	MapType                MapTypePtrInput
	Metadata               pulumi.Input
	Name                   pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
	Type                   pulumi.StringPtrInput
}

func (IntegrationAccountMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountMapArgs)(nil)).Elem()
}

type IntegrationAccountMapInput interface {
	pulumi.Input

	ToIntegrationAccountMapOutput() IntegrationAccountMapOutput
	ToIntegrationAccountMapOutputWithContext(ctx context.Context) IntegrationAccountMapOutput
}

func (*IntegrationAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountMap)(nil)).Elem()
}

func (i *IntegrationAccountMap) ToIntegrationAccountMapOutput() IntegrationAccountMapOutput {
	return i.ToIntegrationAccountMapOutputWithContext(context.Background())
}

func (i *IntegrationAccountMap) ToIntegrationAccountMapOutputWithContext(ctx context.Context) IntegrationAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountMapOutput)
}

type IntegrationAccountMapOutput struct{ *pulumi.OutputState }

func (IntegrationAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountMap)(nil)).Elem()
}

func (o IntegrationAccountMapOutput) ToIntegrationAccountMapOutput() IntegrationAccountMapOutput {
	return o
}

func (o IntegrationAccountMapOutput) ToIntegrationAccountMapOutputWithContext(ctx context.Context) IntegrationAccountMapOutput {
	return o
}

func (o IntegrationAccountMapOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) pulumi.StringOutput { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o IntegrationAccountMapOutput) Content() pulumi.AnyOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) pulumi.AnyOutput { return v.Content }).(pulumi.AnyOutput)
}

func (o IntegrationAccountMapOutput) ContentLink() IntegrationAccountContentLinkResponseOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) IntegrationAccountContentLinkResponseOutput { return v.ContentLink }).(IntegrationAccountContentLinkResponseOutput)
}

func (o IntegrationAccountMapOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) pulumi.StringPtrOutput { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o IntegrationAccountMapOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o IntegrationAccountMapOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o IntegrationAccountMapOutput) MapType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) pulumi.StringPtrOutput { return v.MapType }).(pulumi.StringPtrOutput)
}

func (o IntegrationAccountMapOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

func (o IntegrationAccountMapOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IntegrationAccountMapOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o IntegrationAccountMapOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountMapOutput{})
}




package v20190901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MediaGraph struct {
	pulumi.CustomResourceState

	Created      pulumi.StringOutput                     `pulumi:"created"`
	Description  pulumi.StringPtrOutput                  `pulumi:"description"`
	LastModified pulumi.StringOutput                     `pulumi:"lastModified"`
	Name         pulumi.StringOutput                     `pulumi:"name"`
	Sinks        MediaGraphAssetSinkResponseArrayOutput  `pulumi:"sinks"`
	Sources      MediaGraphRtspSourceResponseArrayOutput `pulumi:"sources"`
	State        pulumi.StringOutput                     `pulumi:"state"`
	Type         pulumi.StringOutput                     `pulumi:"type"`
}


func NewMediaGraph(ctx *pulumi.Context,
	name string, args *MediaGraphArgs, opts ...pulumi.ResourceOption) (*MediaGraph, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sinks == nil {
		return nil, errors.New("invalid value for required argument 'Sinks'")
	}
	if args.Sources == nil {
		return nil, errors.New("invalid value for required argument 'Sources'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:media:MediaGraph"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200201preview:MediaGraph"),
		},
	})
	opts = append(opts, aliases)
	var resource MediaGraph
	err := ctx.RegisterResource("azure-native:media/v20190901preview:MediaGraph", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMediaGraph(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MediaGraphState, opts ...pulumi.ResourceOption) (*MediaGraph, error) {
	var resource MediaGraph
	err := ctx.ReadResource("azure-native:media/v20190901preview:MediaGraph", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type mediaGraphState struct {
}

type MediaGraphState struct {
}

func (MediaGraphState) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaGraphState)(nil)).Elem()
}

type mediaGraphArgs struct {
	AccountName       string                 `pulumi:"accountName"`
	Description       *string                `pulumi:"description"`
	MediaGraphName    *string                `pulumi:"mediaGraphName"`
	ResourceGroupName string                 `pulumi:"resourceGroupName"`
	Sinks             []MediaGraphAssetSink  `pulumi:"sinks"`
	Sources           []MediaGraphRtspSource `pulumi:"sources"`
}


type MediaGraphArgs struct {
	AccountName       pulumi.StringInput
	Description       pulumi.StringPtrInput
	MediaGraphName    pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sinks             MediaGraphAssetSinkArrayInput
	Sources           MediaGraphRtspSourceArrayInput
}

func (MediaGraphArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaGraphArgs)(nil)).Elem()
}

type MediaGraphInput interface {
	pulumi.Input

	ToMediaGraphOutput() MediaGraphOutput
	ToMediaGraphOutputWithContext(ctx context.Context) MediaGraphOutput
}

func (*MediaGraph) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaGraph)(nil)).Elem()
}

func (i *MediaGraph) ToMediaGraphOutput() MediaGraphOutput {
	return i.ToMediaGraphOutputWithContext(context.Background())
}

func (i *MediaGraph) ToMediaGraphOutputWithContext(ctx context.Context) MediaGraphOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphOutput)
}

type MediaGraphOutput struct{ *pulumi.OutputState }

func (MediaGraphOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaGraph)(nil)).Elem()
}

func (o MediaGraphOutput) ToMediaGraphOutput() MediaGraphOutput {
	return o
}

func (o MediaGraphOutput) ToMediaGraphOutputWithContext(ctx context.Context) MediaGraphOutput {
	return o
}

func (o MediaGraphOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaGraph) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

func (o MediaGraphOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaGraph) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o MediaGraphOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaGraph) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

func (o MediaGraphOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaGraph) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MediaGraphOutput) Sinks() MediaGraphAssetSinkResponseArrayOutput {
	return o.ApplyT(func(v *MediaGraph) MediaGraphAssetSinkResponseArrayOutput { return v.Sinks }).(MediaGraphAssetSinkResponseArrayOutput)
}

func (o MediaGraphOutput) Sources() MediaGraphRtspSourceResponseArrayOutput {
	return o.ApplyT(func(v *MediaGraph) MediaGraphRtspSourceResponseArrayOutput { return v.Sources }).(MediaGraphRtspSourceResponseArrayOutput)
}

func (o MediaGraphOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaGraph) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o MediaGraphOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaGraph) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MediaGraphOutput{})
}

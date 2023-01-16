


package v20220915preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DocumentProcessor struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                       `pulumi:"location"`
	Name       pulumi.StringOutput                       `pulumi:"name"`
	Properties DocumentProcessorPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                  `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                    `pulumi:"tags"`
	Type       pulumi.StringOutput                       `pulumi:"type"`
}


func NewDocumentProcessor(ctx *pulumi.Context,
	name string, args *DocumentProcessorArgs, opts ...pulumi.ResourceOption) (*DocumentProcessor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:syntex:DocumentProcessor"),
		},
	})
	opts = append(opts, aliases)
	var resource DocumentProcessor
	err := ctx.RegisterResource("azure-native:syntex/v20220915preview:DocumentProcessor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDocumentProcessor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentProcessorState, opts ...pulumi.ResourceOption) (*DocumentProcessor, error) {
	var resource DocumentProcessor
	err := ctx.ReadResource("azure-native:syntex/v20220915preview:DocumentProcessor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type documentProcessorState struct {
}

type DocumentProcessorState struct {
}

func (DocumentProcessorState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentProcessorState)(nil)).Elem()
}

type documentProcessorArgs struct {
	Location          *string                      `pulumi:"location"`
	ProcessorName     *string                      `pulumi:"processorName"`
	Properties        *DocumentProcessorProperties `pulumi:"properties"`
	ResourceGroupName string                       `pulumi:"resourceGroupName"`
	Tags              map[string]string            `pulumi:"tags"`
}


type DocumentProcessorArgs struct {
	Location          pulumi.StringPtrInput
	ProcessorName     pulumi.StringPtrInput
	Properties        DocumentProcessorPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (DocumentProcessorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentProcessorArgs)(nil)).Elem()
}

type DocumentProcessorInput interface {
	pulumi.Input

	ToDocumentProcessorOutput() DocumentProcessorOutput
	ToDocumentProcessorOutputWithContext(ctx context.Context) DocumentProcessorOutput
}

func (*DocumentProcessor) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentProcessor)(nil)).Elem()
}

func (i *DocumentProcessor) ToDocumentProcessorOutput() DocumentProcessorOutput {
	return i.ToDocumentProcessorOutputWithContext(context.Background())
}

func (i *DocumentProcessor) ToDocumentProcessorOutputWithContext(ctx context.Context) DocumentProcessorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentProcessorOutput)
}

type DocumentProcessorOutput struct{ *pulumi.OutputState }

func (DocumentProcessorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentProcessor)(nil)).Elem()
}

func (o DocumentProcessorOutput) ToDocumentProcessorOutput() DocumentProcessorOutput {
	return o
}

func (o DocumentProcessorOutput) ToDocumentProcessorOutputWithContext(ctx context.Context) DocumentProcessorOutput {
	return o
}

func (o DocumentProcessorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentProcessor) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DocumentProcessorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentProcessor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DocumentProcessorOutput) Properties() DocumentProcessorPropertiesResponseOutput {
	return o.ApplyT(func(v *DocumentProcessor) DocumentProcessorPropertiesResponseOutput { return v.Properties }).(DocumentProcessorPropertiesResponseOutput)
}

func (o DocumentProcessorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DocumentProcessor) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DocumentProcessorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DocumentProcessor) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DocumentProcessorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentProcessor) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DocumentProcessorOutput{})
}

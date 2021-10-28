


package v20191017preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkbookTemplate struct {
	pulumi.CustomResourceState

	Author       pulumi.StringPtrOutput                                 `pulumi:"author"`
	Galleries    WorkbookTemplateGalleryResponseArrayOutput             `pulumi:"galleries"`
	Localized    WorkbookTemplateLocalizedGalleryResponseArrayMapOutput `pulumi:"localized"`
	Location     pulumi.StringOutput                                    `pulumi:"location"`
	Name         pulumi.StringOutput                                    `pulumi:"name"`
	Priority     pulumi.IntPtrOutput                                    `pulumi:"priority"`
	Tags         pulumi.StringMapOutput                                 `pulumi:"tags"`
	TemplateData pulumi.AnyOutput                                       `pulumi:"templateData"`
	Type         pulumi.StringOutput                                    `pulumi:"type"`
}


func NewWorkbookTemplate(ctx *pulumi.Context,
	name string, args *WorkbookTemplateArgs, opts ...pulumi.ResourceOption) (*WorkbookTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Galleries == nil {
		return nil, errors.New("invalid value for required argument 'Galleries'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TemplateData == nil {
		return nil, errors.New("invalid value for required argument 'TemplateData'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:insights/v20191017preview:WorkbookTemplate"),
		},
		{
			Type: pulumi.String("azure-native:insights:WorkbookTemplate"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights:WorkbookTemplate"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20201120:WorkbookTemplate"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights/v20201120:WorkbookTemplate"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkbookTemplate
	err := ctx.RegisterResource("azure-native:insights/v20191017preview:WorkbookTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkbookTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkbookTemplateState, opts ...pulumi.ResourceOption) (*WorkbookTemplate, error) {
	var resource WorkbookTemplate
	err := ctx.ReadResource("azure-native:insights/v20191017preview:WorkbookTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workbookTemplateState struct {
}

type WorkbookTemplateState struct {
}

func (WorkbookTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*workbookTemplateState)(nil)).Elem()
}

type workbookTemplateArgs struct {
	Author            *string                                       `pulumi:"author"`
	Galleries         []WorkbookTemplateGallery                     `pulumi:"galleries"`
	Localized         map[string][]WorkbookTemplateLocalizedGallery `pulumi:"localized"`
	Location          *string                                       `pulumi:"location"`
	Priority          *int                                          `pulumi:"priority"`
	ResourceGroupName string                                        `pulumi:"resourceGroupName"`
	ResourceName      *string                                       `pulumi:"resourceName"`
	Tags              map[string]string                             `pulumi:"tags"`
	TemplateData      interface{}                                   `pulumi:"templateData"`
}


type WorkbookTemplateArgs struct {
	Author            pulumi.StringPtrInput
	Galleries         WorkbookTemplateGalleryArrayInput
	Localized         WorkbookTemplateLocalizedGalleryArrayMapInput
	Location          pulumi.StringPtrInput
	Priority          pulumi.IntPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	TemplateData      pulumi.Input
}

func (WorkbookTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workbookTemplateArgs)(nil)).Elem()
}

type WorkbookTemplateInput interface {
	pulumi.Input

	ToWorkbookTemplateOutput() WorkbookTemplateOutput
	ToWorkbookTemplateOutputWithContext(ctx context.Context) WorkbookTemplateOutput
}

func (*WorkbookTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplate)(nil))
}

func (i *WorkbookTemplate) ToWorkbookTemplateOutput() WorkbookTemplateOutput {
	return i.ToWorkbookTemplateOutputWithContext(context.Background())
}

func (i *WorkbookTemplate) ToWorkbookTemplateOutputWithContext(ctx context.Context) WorkbookTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateOutput)
}

type WorkbookTemplateOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplate)(nil))
}

func (o WorkbookTemplateOutput) ToWorkbookTemplateOutput() WorkbookTemplateOutput {
	return o
}

func (o WorkbookTemplateOutput) ToWorkbookTemplateOutputWithContext(ctx context.Context) WorkbookTemplateOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateInput)(nil)).Elem(), &WorkbookTemplate{})
	pulumi.RegisterOutputType(WorkbookTemplateOutput{})
}

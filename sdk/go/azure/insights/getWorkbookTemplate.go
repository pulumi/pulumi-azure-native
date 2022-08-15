


package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkbookTemplate(ctx *pulumi.Context, args *LookupWorkbookTemplateArgs, opts ...pulumi.InvokeOption) (*LookupWorkbookTemplateResult, error) {
	var rv LookupWorkbookTemplateResult
	err := ctx.Invoke("azure-native:insights:getWorkbookTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkbookTemplateArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupWorkbookTemplateResult struct {
	Author       *string                                               `pulumi:"author"`
	Galleries    []WorkbookTemplateGalleryResponse                     `pulumi:"galleries"`
	Id           string                                                `pulumi:"id"`
	Localized    map[string][]WorkbookTemplateLocalizedGalleryResponse `pulumi:"localized"`
	Location     string                                                `pulumi:"location"`
	Name         string                                                `pulumi:"name"`
	Priority     *int                                                  `pulumi:"priority"`
	Tags         map[string]string                                     `pulumi:"tags"`
	TemplateData interface{}                                           `pulumi:"templateData"`
	Type         string                                                `pulumi:"type"`
}

func LookupWorkbookTemplateOutput(ctx *pulumi.Context, args LookupWorkbookTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupWorkbookTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkbookTemplateResult, error) {
			args := v.(LookupWorkbookTemplateArgs)
			r, err := LookupWorkbookTemplate(ctx, &args, opts...)
			var s LookupWorkbookTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkbookTemplateResultOutput)
}

type LookupWorkbookTemplateOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupWorkbookTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkbookTemplateArgs)(nil)).Elem()
}


type LookupWorkbookTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupWorkbookTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkbookTemplateResult)(nil)).Elem()
}

func (o LookupWorkbookTemplateResultOutput) ToLookupWorkbookTemplateResultOutput() LookupWorkbookTemplateResultOutput {
	return o
}

func (o LookupWorkbookTemplateResultOutput) ToLookupWorkbookTemplateResultOutputWithContext(ctx context.Context) LookupWorkbookTemplateResultOutput {
	return o
}

func (o LookupWorkbookTemplateResultOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookTemplateResult) *string { return v.Author }).(pulumi.StringPtrOutput)
}

func (o LookupWorkbookTemplateResultOutput) Galleries() WorkbookTemplateGalleryResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkbookTemplateResult) []WorkbookTemplateGalleryResponse { return v.Galleries }).(WorkbookTemplateGalleryResponseArrayOutput)
}

func (o LookupWorkbookTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkbookTemplateResultOutput) Localized() WorkbookTemplateLocalizedGalleryResponseArrayMapOutput {
	return o.ApplyT(func(v LookupWorkbookTemplateResult) map[string][]WorkbookTemplateLocalizedGalleryResponse {
		return v.Localized
	}).(WorkbookTemplateLocalizedGalleryResponseArrayMapOutput)
}

func (o LookupWorkbookTemplateResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookTemplateResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupWorkbookTemplateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookTemplateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkbookTemplateResultOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWorkbookTemplateResult) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o LookupWorkbookTemplateResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkbookTemplateResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWorkbookTemplateResultOutput) TemplateData() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupWorkbookTemplateResult) interface{} { return v.TemplateData }).(pulumi.AnyOutput)
}

func (o LookupWorkbookTemplateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookTemplateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkbookTemplateResultOutput{})
}

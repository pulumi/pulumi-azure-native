


package v20201120

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkbookTemplateGallery struct {
	Category     *string `pulumi:"category"`
	Name         *string `pulumi:"name"`
	Order        *int    `pulumi:"order"`
	ResourceType *string `pulumi:"resourceType"`
	Type         *string `pulumi:"type"`
}





type WorkbookTemplateGalleryInput interface {
	pulumi.Input

	ToWorkbookTemplateGalleryOutput() WorkbookTemplateGalleryOutput
	ToWorkbookTemplateGalleryOutputWithContext(context.Context) WorkbookTemplateGalleryOutput
}

type WorkbookTemplateGalleryArgs struct {
	Category     pulumi.StringPtrInput `pulumi:"category"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
	Order        pulumi.IntPtrInput    `pulumi:"order"`
	ResourceType pulumi.StringPtrInput `pulumi:"resourceType"`
	Type         pulumi.StringPtrInput `pulumi:"type"`
}

func (WorkbookTemplateGalleryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateGallery)(nil)).Elem()
}

func (i WorkbookTemplateGalleryArgs) ToWorkbookTemplateGalleryOutput() WorkbookTemplateGalleryOutput {
	return i.ToWorkbookTemplateGalleryOutputWithContext(context.Background())
}

func (i WorkbookTemplateGalleryArgs) ToWorkbookTemplateGalleryOutputWithContext(ctx context.Context) WorkbookTemplateGalleryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateGalleryOutput)
}





type WorkbookTemplateGalleryArrayInput interface {
	pulumi.Input

	ToWorkbookTemplateGalleryArrayOutput() WorkbookTemplateGalleryArrayOutput
	ToWorkbookTemplateGalleryArrayOutputWithContext(context.Context) WorkbookTemplateGalleryArrayOutput
}

type WorkbookTemplateGalleryArray []WorkbookTemplateGalleryInput

func (WorkbookTemplateGalleryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateGallery)(nil)).Elem()
}

func (i WorkbookTemplateGalleryArray) ToWorkbookTemplateGalleryArrayOutput() WorkbookTemplateGalleryArrayOutput {
	return i.ToWorkbookTemplateGalleryArrayOutputWithContext(context.Background())
}

func (i WorkbookTemplateGalleryArray) ToWorkbookTemplateGalleryArrayOutputWithContext(ctx context.Context) WorkbookTemplateGalleryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateGalleryArrayOutput)
}

type WorkbookTemplateGalleryOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateGalleryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateGallery)(nil)).Elem()
}

func (o WorkbookTemplateGalleryOutput) ToWorkbookTemplateGalleryOutput() WorkbookTemplateGalleryOutput {
	return o
}

func (o WorkbookTemplateGalleryOutput) ToWorkbookTemplateGalleryOutputWithContext(ctx context.Context) WorkbookTemplateGalleryOutput {
	return o
}

func (o WorkbookTemplateGalleryOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGallery) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o WorkbookTemplateGalleryOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGallery) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WorkbookTemplateGalleryOutput) Order() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGallery) *int { return v.Order }).(pulumi.IntPtrOutput)
}

func (o WorkbookTemplateGalleryOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGallery) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o WorkbookTemplateGalleryOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGallery) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type WorkbookTemplateGalleryArrayOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateGalleryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateGallery)(nil)).Elem()
}

func (o WorkbookTemplateGalleryArrayOutput) ToWorkbookTemplateGalleryArrayOutput() WorkbookTemplateGalleryArrayOutput {
	return o
}

func (o WorkbookTemplateGalleryArrayOutput) ToWorkbookTemplateGalleryArrayOutputWithContext(ctx context.Context) WorkbookTemplateGalleryArrayOutput {
	return o
}

func (o WorkbookTemplateGalleryArrayOutput) Index(i pulumi.IntInput) WorkbookTemplateGalleryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkbookTemplateGallery {
		return vs[0].([]WorkbookTemplateGallery)[vs[1].(int)]
	}).(WorkbookTemplateGalleryOutput)
}

type WorkbookTemplateGalleryResponse struct {
	Category     *string `pulumi:"category"`
	Name         *string `pulumi:"name"`
	Order        *int    `pulumi:"order"`
	ResourceType *string `pulumi:"resourceType"`
	Type         *string `pulumi:"type"`
}





type WorkbookTemplateGalleryResponseInput interface {
	pulumi.Input

	ToWorkbookTemplateGalleryResponseOutput() WorkbookTemplateGalleryResponseOutput
	ToWorkbookTemplateGalleryResponseOutputWithContext(context.Context) WorkbookTemplateGalleryResponseOutput
}

type WorkbookTemplateGalleryResponseArgs struct {
	Category     pulumi.StringPtrInput `pulumi:"category"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
	Order        pulumi.IntPtrInput    `pulumi:"order"`
	ResourceType pulumi.StringPtrInput `pulumi:"resourceType"`
	Type         pulumi.StringPtrInput `pulumi:"type"`
}

func (WorkbookTemplateGalleryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateGalleryResponse)(nil)).Elem()
}

func (i WorkbookTemplateGalleryResponseArgs) ToWorkbookTemplateGalleryResponseOutput() WorkbookTemplateGalleryResponseOutput {
	return i.ToWorkbookTemplateGalleryResponseOutputWithContext(context.Background())
}

func (i WorkbookTemplateGalleryResponseArgs) ToWorkbookTemplateGalleryResponseOutputWithContext(ctx context.Context) WorkbookTemplateGalleryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateGalleryResponseOutput)
}





type WorkbookTemplateGalleryResponseArrayInput interface {
	pulumi.Input

	ToWorkbookTemplateGalleryResponseArrayOutput() WorkbookTemplateGalleryResponseArrayOutput
	ToWorkbookTemplateGalleryResponseArrayOutputWithContext(context.Context) WorkbookTemplateGalleryResponseArrayOutput
}

type WorkbookTemplateGalleryResponseArray []WorkbookTemplateGalleryResponseInput

func (WorkbookTemplateGalleryResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateGalleryResponse)(nil)).Elem()
}

func (i WorkbookTemplateGalleryResponseArray) ToWorkbookTemplateGalleryResponseArrayOutput() WorkbookTemplateGalleryResponseArrayOutput {
	return i.ToWorkbookTemplateGalleryResponseArrayOutputWithContext(context.Background())
}

func (i WorkbookTemplateGalleryResponseArray) ToWorkbookTemplateGalleryResponseArrayOutputWithContext(ctx context.Context) WorkbookTemplateGalleryResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateGalleryResponseArrayOutput)
}

type WorkbookTemplateGalleryResponseOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateGalleryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateGalleryResponse)(nil)).Elem()
}

func (o WorkbookTemplateGalleryResponseOutput) ToWorkbookTemplateGalleryResponseOutput() WorkbookTemplateGalleryResponseOutput {
	return o
}

func (o WorkbookTemplateGalleryResponseOutput) ToWorkbookTemplateGalleryResponseOutputWithContext(ctx context.Context) WorkbookTemplateGalleryResponseOutput {
	return o
}

func (o WorkbookTemplateGalleryResponseOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGalleryResponse) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o WorkbookTemplateGalleryResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGalleryResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WorkbookTemplateGalleryResponseOutput) Order() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGalleryResponse) *int { return v.Order }).(pulumi.IntPtrOutput)
}

func (o WorkbookTemplateGalleryResponseOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGalleryResponse) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o WorkbookTemplateGalleryResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGalleryResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type WorkbookTemplateGalleryResponseArrayOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateGalleryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateGalleryResponse)(nil)).Elem()
}

func (o WorkbookTemplateGalleryResponseArrayOutput) ToWorkbookTemplateGalleryResponseArrayOutput() WorkbookTemplateGalleryResponseArrayOutput {
	return o
}

func (o WorkbookTemplateGalleryResponseArrayOutput) ToWorkbookTemplateGalleryResponseArrayOutputWithContext(ctx context.Context) WorkbookTemplateGalleryResponseArrayOutput {
	return o
}

func (o WorkbookTemplateGalleryResponseArrayOutput) Index(i pulumi.IntInput) WorkbookTemplateGalleryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkbookTemplateGalleryResponse {
		return vs[0].([]WorkbookTemplateGalleryResponse)[vs[1].(int)]
	}).(WorkbookTemplateGalleryResponseOutput)
}

type WorkbookTemplateLocalizedGallery struct {
	Galleries    []WorkbookTemplateGallery `pulumi:"galleries"`
	TemplateData interface{}               `pulumi:"templateData"`
}





type WorkbookTemplateLocalizedGalleryInput interface {
	pulumi.Input

	ToWorkbookTemplateLocalizedGalleryOutput() WorkbookTemplateLocalizedGalleryOutput
	ToWorkbookTemplateLocalizedGalleryOutputWithContext(context.Context) WorkbookTemplateLocalizedGalleryOutput
}

type WorkbookTemplateLocalizedGalleryArgs struct {
	Galleries    WorkbookTemplateGalleryArrayInput `pulumi:"galleries"`
	TemplateData pulumi.Input                      `pulumi:"templateData"`
}

func (WorkbookTemplateLocalizedGalleryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateLocalizedGallery)(nil)).Elem()
}

func (i WorkbookTemplateLocalizedGalleryArgs) ToWorkbookTemplateLocalizedGalleryOutput() WorkbookTemplateLocalizedGalleryOutput {
	return i.ToWorkbookTemplateLocalizedGalleryOutputWithContext(context.Background())
}

func (i WorkbookTemplateLocalizedGalleryArgs) ToWorkbookTemplateLocalizedGalleryOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateLocalizedGalleryOutput)
}





type WorkbookTemplateLocalizedGalleryArrayInput interface {
	pulumi.Input

	ToWorkbookTemplateLocalizedGalleryArrayOutput() WorkbookTemplateLocalizedGalleryArrayOutput
	ToWorkbookTemplateLocalizedGalleryArrayOutputWithContext(context.Context) WorkbookTemplateLocalizedGalleryArrayOutput
}

type WorkbookTemplateLocalizedGalleryArray []WorkbookTemplateLocalizedGalleryInput

func (WorkbookTemplateLocalizedGalleryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateLocalizedGallery)(nil)).Elem()
}

func (i WorkbookTemplateLocalizedGalleryArray) ToWorkbookTemplateLocalizedGalleryArrayOutput() WorkbookTemplateLocalizedGalleryArrayOutput {
	return i.ToWorkbookTemplateLocalizedGalleryArrayOutputWithContext(context.Background())
}

func (i WorkbookTemplateLocalizedGalleryArray) ToWorkbookTemplateLocalizedGalleryArrayOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateLocalizedGalleryArrayOutput)
}

type WorkbookTemplateLocalizedGalleryOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateLocalizedGalleryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateLocalizedGallery)(nil)).Elem()
}

func (o WorkbookTemplateLocalizedGalleryOutput) ToWorkbookTemplateLocalizedGalleryOutput() WorkbookTemplateLocalizedGalleryOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryOutput) ToWorkbookTemplateLocalizedGalleryOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryOutput) Galleries() WorkbookTemplateGalleryArrayOutput {
	return o.ApplyT(func(v WorkbookTemplateLocalizedGallery) []WorkbookTemplateGallery { return v.Galleries }).(WorkbookTemplateGalleryArrayOutput)
}

func (o WorkbookTemplateLocalizedGalleryOutput) TemplateData() pulumi.AnyOutput {
	return o.ApplyT(func(v WorkbookTemplateLocalizedGallery) interface{} { return v.TemplateData }).(pulumi.AnyOutput)
}

type WorkbookTemplateLocalizedGalleryArrayOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateLocalizedGalleryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateLocalizedGallery)(nil)).Elem()
}

func (o WorkbookTemplateLocalizedGalleryArrayOutput) ToWorkbookTemplateLocalizedGalleryArrayOutput() WorkbookTemplateLocalizedGalleryArrayOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryArrayOutput) ToWorkbookTemplateLocalizedGalleryArrayOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryArrayOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryArrayOutput) Index(i pulumi.IntInput) WorkbookTemplateLocalizedGalleryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkbookTemplateLocalizedGallery {
		return vs[0].([]WorkbookTemplateLocalizedGallery)[vs[1].(int)]
	}).(WorkbookTemplateLocalizedGalleryOutput)
}

type WorkbookTemplateLocalizedGalleryResponse struct {
	Galleries    []WorkbookTemplateGalleryResponse `pulumi:"galleries"`
	TemplateData interface{}                       `pulumi:"templateData"`
}





type WorkbookTemplateLocalizedGalleryResponseInput interface {
	pulumi.Input

	ToWorkbookTemplateLocalizedGalleryResponseOutput() WorkbookTemplateLocalizedGalleryResponseOutput
	ToWorkbookTemplateLocalizedGalleryResponseOutputWithContext(context.Context) WorkbookTemplateLocalizedGalleryResponseOutput
}

type WorkbookTemplateLocalizedGalleryResponseArgs struct {
	Galleries    WorkbookTemplateGalleryResponseArrayInput `pulumi:"galleries"`
	TemplateData pulumi.Input                              `pulumi:"templateData"`
}

func (WorkbookTemplateLocalizedGalleryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateLocalizedGalleryResponse)(nil)).Elem()
}

func (i WorkbookTemplateLocalizedGalleryResponseArgs) ToWorkbookTemplateLocalizedGalleryResponseOutput() WorkbookTemplateLocalizedGalleryResponseOutput {
	return i.ToWorkbookTemplateLocalizedGalleryResponseOutputWithContext(context.Background())
}

func (i WorkbookTemplateLocalizedGalleryResponseArgs) ToWorkbookTemplateLocalizedGalleryResponseOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateLocalizedGalleryResponseOutput)
}





type WorkbookTemplateLocalizedGalleryResponseArrayInput interface {
	pulumi.Input

	ToWorkbookTemplateLocalizedGalleryResponseArrayOutput() WorkbookTemplateLocalizedGalleryResponseArrayOutput
	ToWorkbookTemplateLocalizedGalleryResponseArrayOutputWithContext(context.Context) WorkbookTemplateLocalizedGalleryResponseArrayOutput
}

type WorkbookTemplateLocalizedGalleryResponseArray []WorkbookTemplateLocalizedGalleryResponseInput

func (WorkbookTemplateLocalizedGalleryResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateLocalizedGalleryResponse)(nil)).Elem()
}

func (i WorkbookTemplateLocalizedGalleryResponseArray) ToWorkbookTemplateLocalizedGalleryResponseArrayOutput() WorkbookTemplateLocalizedGalleryResponseArrayOutput {
	return i.ToWorkbookTemplateLocalizedGalleryResponseArrayOutputWithContext(context.Background())
}

func (i WorkbookTemplateLocalizedGalleryResponseArray) ToWorkbookTemplateLocalizedGalleryResponseArrayOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateLocalizedGalleryResponseArrayOutput)
}

type WorkbookTemplateLocalizedGalleryResponseOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateLocalizedGalleryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateLocalizedGalleryResponse)(nil)).Elem()
}

func (o WorkbookTemplateLocalizedGalleryResponseOutput) ToWorkbookTemplateLocalizedGalleryResponseOutput() WorkbookTemplateLocalizedGalleryResponseOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryResponseOutput) ToWorkbookTemplateLocalizedGalleryResponseOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryResponseOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryResponseOutput) Galleries() WorkbookTemplateGalleryResponseArrayOutput {
	return o.ApplyT(func(v WorkbookTemplateLocalizedGalleryResponse) []WorkbookTemplateGalleryResponse { return v.Galleries }).(WorkbookTemplateGalleryResponseArrayOutput)
}

func (o WorkbookTemplateLocalizedGalleryResponseOutput) TemplateData() pulumi.AnyOutput {
	return o.ApplyT(func(v WorkbookTemplateLocalizedGalleryResponse) interface{} { return v.TemplateData }).(pulumi.AnyOutput)
}

type WorkbookTemplateLocalizedGalleryResponseArrayOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateLocalizedGalleryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateLocalizedGalleryResponse)(nil)).Elem()
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayOutput) ToWorkbookTemplateLocalizedGalleryResponseArrayOutput() WorkbookTemplateLocalizedGalleryResponseArrayOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayOutput) ToWorkbookTemplateLocalizedGalleryResponseArrayOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryResponseArrayOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayOutput) Index(i pulumi.IntInput) WorkbookTemplateLocalizedGalleryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkbookTemplateLocalizedGalleryResponse {
		return vs[0].([]WorkbookTemplateLocalizedGalleryResponse)[vs[1].(int)]
	}).(WorkbookTemplateLocalizedGalleryResponseOutput)
}

type WorkbookTemplateLocalizedGalleryArrayMap map[string]WorkbookTemplateLocalizedGalleryArrayInput

func (WorkbookTemplateLocalizedGalleryArrayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateLocalizedGalleryArray)(nil)).Elem()
}

func (i WorkbookTemplateLocalizedGalleryArrayMap) ToWorkbookTemplateLocalizedGalleryArrayMapOutput() WorkbookTemplateLocalizedGalleryArrayMapOutput {
	return i.ToWorkbookTemplateLocalizedGalleryArrayMapOutputWithContext(context.Background())
}

func (i WorkbookTemplateLocalizedGalleryArrayMap) ToWorkbookTemplateLocalizedGalleryArrayMapOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryArrayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateLocalizedGalleryArrayMapOutput)
}

type WorkbookTemplateLocalizedGalleryArrayMapOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateLocalizedGalleryArrayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WorkbookTemplateLocalizedGalleryArray)(nil)).Elem()
}

func (o WorkbookTemplateLocalizedGalleryArrayMapOutput) ToWorkbookTemplateLocalizedGalleryArrayMapOutput() WorkbookTemplateLocalizedGalleryArrayMapOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryArrayMapOutput) ToWorkbookTemplateLocalizedGalleryArrayMapOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryArrayMapOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryArrayMapOutput) MapIndex(k pulumi.StringInput) WorkbookTemplateLocalizedGalleryArrayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WorkbookTemplateLocalizedGalleryArray {
		return vs[0].(map[string]WorkbookTemplateLocalizedGalleryArray)[vs[1].(string)]
	}).(WorkbookTemplateLocalizedGalleryArrayOutput)
}





type WorkbookTemplateLocalizedGalleryArrayMapInput interface {
	pulumi.Input

	ToWorkbookTemplateLocalizedGalleryArrayMapOutput() WorkbookTemplateLocalizedGalleryArrayMapOutput
	ToWorkbookTemplateLocalizedGalleryArrayMapOutputWithContext(context.Context) WorkbookTemplateLocalizedGalleryArrayMapOutput
}

type WorkbookTemplateLocalizedGalleryResponseArrayMap map[string]WorkbookTemplateLocalizedGalleryResponseArrayInput

func (WorkbookTemplateLocalizedGalleryResponseArrayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateLocalizedGalleryResponseArray)(nil)).Elem()
}

func (i WorkbookTemplateLocalizedGalleryResponseArrayMap) ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutput() WorkbookTemplateLocalizedGalleryResponseArrayMapOutput {
	return i.ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutputWithContext(context.Background())
}

func (i WorkbookTemplateLocalizedGalleryResponseArrayMap) ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryResponseArrayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateLocalizedGalleryResponseArrayMapOutput)
}

type WorkbookTemplateLocalizedGalleryResponseArrayMapOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateLocalizedGalleryResponseArrayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WorkbookTemplateLocalizedGalleryResponseArray)(nil)).Elem()
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayMapOutput) ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutput() WorkbookTemplateLocalizedGalleryResponseArrayMapOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayMapOutput) ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryResponseArrayMapOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayMapOutput) MapIndex(k pulumi.StringInput) WorkbookTemplateLocalizedGalleryResponseArrayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WorkbookTemplateLocalizedGalleryResponseArray {
		return vs[0].(map[string]WorkbookTemplateLocalizedGalleryResponseArray)[vs[1].(string)]
	}).(WorkbookTemplateLocalizedGalleryResponseArrayOutput)
}





type WorkbookTemplateLocalizedGalleryResponseArrayMapInput interface {
	pulumi.Input

	ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutput() WorkbookTemplateLocalizedGalleryResponseArrayMapOutput
	ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutputWithContext(context.Context) WorkbookTemplateLocalizedGalleryResponseArrayMapOutput
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateGalleryInput)(nil)).Elem(), WorkbookTemplateGalleryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateGalleryArrayInput)(nil)).Elem(), WorkbookTemplateGalleryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateGalleryResponseInput)(nil)).Elem(), WorkbookTemplateGalleryResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateGalleryResponseArrayInput)(nil)).Elem(), WorkbookTemplateGalleryResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateLocalizedGalleryInput)(nil)).Elem(), WorkbookTemplateLocalizedGalleryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateLocalizedGalleryArrayInput)(nil)).Elem(), WorkbookTemplateLocalizedGalleryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateLocalizedGalleryResponseInput)(nil)).Elem(), WorkbookTemplateLocalizedGalleryResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateLocalizedGalleryResponseArrayInput)(nil)).Elem(), WorkbookTemplateLocalizedGalleryResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateLocalizedGalleryArrayMapInput)(nil)).Elem(), WorkbookTemplateLocalizedGalleryArrayMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateLocalizedGalleryResponseArrayMapInput)(nil)).Elem(), WorkbookTemplateLocalizedGalleryResponseArrayMap{})
	pulumi.RegisterOutputType(WorkbookTemplateGalleryOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateGalleryArrayOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateGalleryResponseOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateGalleryResponseArrayOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateLocalizedGalleryOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateLocalizedGalleryArrayOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateLocalizedGalleryResponseOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateLocalizedGalleryResponseArrayOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateLocalizedGalleryArrayMapOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateLocalizedGalleryResponseArrayMapOutput{})
}

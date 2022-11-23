


package v20220701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceSelector struct {
	Name      *string    `pulumi:"name"`
	Selectors []Selector `pulumi:"selectors"`
}





type ResourceSelectorInput interface {
	pulumi.Input

	ToResourceSelectorOutput() ResourceSelectorOutput
	ToResourceSelectorOutputWithContext(context.Context) ResourceSelectorOutput
}

type ResourceSelectorArgs struct {
	Name      pulumi.StringPtrInput `pulumi:"name"`
	Selectors SelectorArrayInput    `pulumi:"selectors"`
}

func (ResourceSelectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSelector)(nil)).Elem()
}

func (i ResourceSelectorArgs) ToResourceSelectorOutput() ResourceSelectorOutput {
	return i.ToResourceSelectorOutputWithContext(context.Background())
}

func (i ResourceSelectorArgs) ToResourceSelectorOutputWithContext(ctx context.Context) ResourceSelectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSelectorOutput)
}





type ResourceSelectorArrayInput interface {
	pulumi.Input

	ToResourceSelectorArrayOutput() ResourceSelectorArrayOutput
	ToResourceSelectorArrayOutputWithContext(context.Context) ResourceSelectorArrayOutput
}

type ResourceSelectorArray []ResourceSelectorInput

func (ResourceSelectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceSelector)(nil)).Elem()
}

func (i ResourceSelectorArray) ToResourceSelectorArrayOutput() ResourceSelectorArrayOutput {
	return i.ToResourceSelectorArrayOutputWithContext(context.Background())
}

func (i ResourceSelectorArray) ToResourceSelectorArrayOutputWithContext(ctx context.Context) ResourceSelectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSelectorArrayOutput)
}

type ResourceSelectorOutput struct{ *pulumi.OutputState }

func (ResourceSelectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSelector)(nil)).Elem()
}

func (o ResourceSelectorOutput) ToResourceSelectorOutput() ResourceSelectorOutput {
	return o
}

func (o ResourceSelectorOutput) ToResourceSelectorOutputWithContext(ctx context.Context) ResourceSelectorOutput {
	return o
}

func (o ResourceSelectorOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSelector) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ResourceSelectorOutput) Selectors() SelectorArrayOutput {
	return o.ApplyT(func(v ResourceSelector) []Selector { return v.Selectors }).(SelectorArrayOutput)
}

type ResourceSelectorArrayOutput struct{ *pulumi.OutputState }

func (ResourceSelectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceSelector)(nil)).Elem()
}

func (o ResourceSelectorArrayOutput) ToResourceSelectorArrayOutput() ResourceSelectorArrayOutput {
	return o
}

func (o ResourceSelectorArrayOutput) ToResourceSelectorArrayOutputWithContext(ctx context.Context) ResourceSelectorArrayOutput {
	return o
}

func (o ResourceSelectorArrayOutput) Index(i pulumi.IntInput) ResourceSelectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceSelector {
		return vs[0].([]ResourceSelector)[vs[1].(int)]
	}).(ResourceSelectorOutput)
}

type ResourceSelectorResponse struct {
	Name      *string            `pulumi:"name"`
	Selectors []SelectorResponse `pulumi:"selectors"`
}

type ResourceSelectorResponseOutput struct{ *pulumi.OutputState }

func (ResourceSelectorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSelectorResponse)(nil)).Elem()
}

func (o ResourceSelectorResponseOutput) ToResourceSelectorResponseOutput() ResourceSelectorResponseOutput {
	return o
}

func (o ResourceSelectorResponseOutput) ToResourceSelectorResponseOutputWithContext(ctx context.Context) ResourceSelectorResponseOutput {
	return o
}

func (o ResourceSelectorResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSelectorResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ResourceSelectorResponseOutput) Selectors() SelectorResponseArrayOutput {
	return o.ApplyT(func(v ResourceSelectorResponse) []SelectorResponse { return v.Selectors }).(SelectorResponseArrayOutput)
}

type ResourceSelectorResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceSelectorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceSelectorResponse)(nil)).Elem()
}

func (o ResourceSelectorResponseArrayOutput) ToResourceSelectorResponseArrayOutput() ResourceSelectorResponseArrayOutput {
	return o
}

func (o ResourceSelectorResponseArrayOutput) ToResourceSelectorResponseArrayOutputWithContext(ctx context.Context) ResourceSelectorResponseArrayOutput {
	return o
}

func (o ResourceSelectorResponseArrayOutput) Index(i pulumi.IntInput) ResourceSelectorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceSelectorResponse {
		return vs[0].([]ResourceSelectorResponse)[vs[1].(int)]
	}).(ResourceSelectorResponseOutput)
}

type Selector struct {
	In    []string `pulumi:"in"`
	Kind  *string  `pulumi:"kind"`
	NotIn []string `pulumi:"notIn"`
}





type SelectorInput interface {
	pulumi.Input

	ToSelectorOutput() SelectorOutput
	ToSelectorOutputWithContext(context.Context) SelectorOutput
}

type SelectorArgs struct {
	In    pulumi.StringArrayInput `pulumi:"in"`
	Kind  pulumi.StringPtrInput   `pulumi:"kind"`
	NotIn pulumi.StringArrayInput `pulumi:"notIn"`
}

func (SelectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Selector)(nil)).Elem()
}

func (i SelectorArgs) ToSelectorOutput() SelectorOutput {
	return i.ToSelectorOutputWithContext(context.Background())
}

func (i SelectorArgs) ToSelectorOutputWithContext(ctx context.Context) SelectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectorOutput)
}





type SelectorArrayInput interface {
	pulumi.Input

	ToSelectorArrayOutput() SelectorArrayOutput
	ToSelectorArrayOutputWithContext(context.Context) SelectorArrayOutput
}

type SelectorArray []SelectorInput

func (SelectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Selector)(nil)).Elem()
}

func (i SelectorArray) ToSelectorArrayOutput() SelectorArrayOutput {
	return i.ToSelectorArrayOutputWithContext(context.Background())
}

func (i SelectorArray) ToSelectorArrayOutputWithContext(ctx context.Context) SelectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectorArrayOutput)
}

type SelectorOutput struct{ *pulumi.OutputState }

func (SelectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Selector)(nil)).Elem()
}

func (o SelectorOutput) ToSelectorOutput() SelectorOutput {
	return o
}

func (o SelectorOutput) ToSelectorOutputWithContext(ctx context.Context) SelectorOutput {
	return o
}

func (o SelectorOutput) In() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Selector) []string { return v.In }).(pulumi.StringArrayOutput)
}

func (o SelectorOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Selector) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SelectorOutput) NotIn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Selector) []string { return v.NotIn }).(pulumi.StringArrayOutput)
}

type SelectorArrayOutput struct{ *pulumi.OutputState }

func (SelectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Selector)(nil)).Elem()
}

func (o SelectorArrayOutput) ToSelectorArrayOutput() SelectorArrayOutput {
	return o
}

func (o SelectorArrayOutput) ToSelectorArrayOutputWithContext(ctx context.Context) SelectorArrayOutput {
	return o
}

func (o SelectorArrayOutput) Index(i pulumi.IntInput) SelectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Selector {
		return vs[0].([]Selector)[vs[1].(int)]
	}).(SelectorOutput)
}

type SelectorResponse struct {
	In    []string `pulumi:"in"`
	Kind  *string  `pulumi:"kind"`
	NotIn []string `pulumi:"notIn"`
}

type SelectorResponseOutput struct{ *pulumi.OutputState }

func (SelectorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectorResponse)(nil)).Elem()
}

func (o SelectorResponseOutput) ToSelectorResponseOutput() SelectorResponseOutput {
	return o
}

func (o SelectorResponseOutput) ToSelectorResponseOutputWithContext(ctx context.Context) SelectorResponseOutput {
	return o
}

func (o SelectorResponseOutput) In() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SelectorResponse) []string { return v.In }).(pulumi.StringArrayOutput)
}

func (o SelectorResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SelectorResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SelectorResponseOutput) NotIn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SelectorResponse) []string { return v.NotIn }).(pulumi.StringArrayOutput)
}

type SelectorResponseArrayOutput struct{ *pulumi.OutputState }

func (SelectorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SelectorResponse)(nil)).Elem()
}

func (o SelectorResponseArrayOutput) ToSelectorResponseArrayOutput() SelectorResponseArrayOutput {
	return o
}

func (o SelectorResponseArrayOutput) ToSelectorResponseArrayOutputWithContext(ctx context.Context) SelectorResponseArrayOutput {
	return o
}

func (o SelectorResponseArrayOutput) Index(i pulumi.IntInput) SelectorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SelectorResponse {
		return vs[0].([]SelectorResponse)[vs[1].(int)]
	}).(SelectorResponseOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceSelectorOutput{})
	pulumi.RegisterOutputType(ResourceSelectorArrayOutput{})
	pulumi.RegisterOutputType(ResourceSelectorResponseOutput{})
	pulumi.RegisterOutputType(ResourceSelectorResponseArrayOutput{})
	pulumi.RegisterOutputType(SelectorOutput{})
	pulumi.RegisterOutputType(SelectorArrayOutput{})
	pulumi.RegisterOutputType(SelectorResponseOutput{})
	pulumi.RegisterOutputType(SelectorResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}

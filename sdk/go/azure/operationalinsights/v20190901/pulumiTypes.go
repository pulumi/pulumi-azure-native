


package v20190901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LogAnalyticsQueryPackQueryPropertiesRelated struct {
	Categories    []string `pulumi:"categories"`
	ResourceTypes []string `pulumi:"resourceTypes"`
	Solutions     []string `pulumi:"solutions"`
}





type LogAnalyticsQueryPackQueryPropertiesRelatedInput interface {
	pulumi.Input

	ToLogAnalyticsQueryPackQueryPropertiesRelatedOutput() LogAnalyticsQueryPackQueryPropertiesRelatedOutput
	ToLogAnalyticsQueryPackQueryPropertiesRelatedOutputWithContext(context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedOutput
}

type LogAnalyticsQueryPackQueryPropertiesRelatedArgs struct {
	Categories    pulumi.StringArrayInput `pulumi:"categories"`
	ResourceTypes pulumi.StringArrayInput `pulumi:"resourceTypes"`
	Solutions     pulumi.StringArrayInput `pulumi:"solutions"`
}

func (LogAnalyticsQueryPackQueryPropertiesRelatedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsQueryPackQueryPropertiesRelated)(nil)).Elem()
}

func (i LogAnalyticsQueryPackQueryPropertiesRelatedArgs) ToLogAnalyticsQueryPackQueryPropertiesRelatedOutput() LogAnalyticsQueryPackQueryPropertiesRelatedOutput {
	return i.ToLogAnalyticsQueryPackQueryPropertiesRelatedOutputWithContext(context.Background())
}

func (i LogAnalyticsQueryPackQueryPropertiesRelatedArgs) ToLogAnalyticsQueryPackQueryPropertiesRelatedOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsQueryPackQueryPropertiesRelatedOutput)
}

func (i LogAnalyticsQueryPackQueryPropertiesRelatedArgs) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput() LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return i.ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(context.Background())
}

func (i LogAnalyticsQueryPackQueryPropertiesRelatedArgs) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsQueryPackQueryPropertiesRelatedOutput).ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(ctx)
}









type LogAnalyticsQueryPackQueryPropertiesRelatedPtrInput interface {
	pulumi.Input

	ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput() LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput
	ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput
}

type logAnalyticsQueryPackQueryPropertiesRelatedPtrType LogAnalyticsQueryPackQueryPropertiesRelatedArgs

func LogAnalyticsQueryPackQueryPropertiesRelatedPtr(v *LogAnalyticsQueryPackQueryPropertiesRelatedArgs) LogAnalyticsQueryPackQueryPropertiesRelatedPtrInput {
	return (*logAnalyticsQueryPackQueryPropertiesRelatedPtrType)(v)
}

func (*logAnalyticsQueryPackQueryPropertiesRelatedPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsQueryPackQueryPropertiesRelated)(nil)).Elem()
}

func (i *logAnalyticsQueryPackQueryPropertiesRelatedPtrType) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput() LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return i.ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(context.Background())
}

func (i *logAnalyticsQueryPackQueryPropertiesRelatedPtrType) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput)
}

type LogAnalyticsQueryPackQueryPropertiesRelatedOutput struct{ *pulumi.OutputState }

func (LogAnalyticsQueryPackQueryPropertiesRelatedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsQueryPackQueryPropertiesRelated)(nil)).Elem()
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) ToLogAnalyticsQueryPackQueryPropertiesRelatedOutput() LogAnalyticsQueryPackQueryPropertiesRelatedOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) ToLogAnalyticsQueryPackQueryPropertiesRelatedOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput() LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return o.ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(context.Background())
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogAnalyticsQueryPackQueryPropertiesRelated) *LogAnalyticsQueryPackQueryPropertiesRelated {
		return &v
	}).(LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogAnalyticsQueryPackQueryPropertiesRelated) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) ResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogAnalyticsQueryPackQueryPropertiesRelated) []string { return v.ResourceTypes }).(pulumi.StringArrayOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) Solutions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogAnalyticsQueryPackQueryPropertiesRelated) []string { return v.Solutions }).(pulumi.StringArrayOutput)
}

type LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput struct{ *pulumi.OutputState }

func (LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsQueryPackQueryPropertiesRelated)(nil)).Elem()
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput() LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) Elem() LogAnalyticsQueryPackQueryPropertiesRelatedOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesRelated) LogAnalyticsQueryPackQueryPropertiesRelated {
		if v != nil {
			return *v
		}
		var ret LogAnalyticsQueryPackQueryPropertiesRelated
		return ret
	}).(LogAnalyticsQueryPackQueryPropertiesRelatedOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesRelated) []string {
		if v == nil {
			return nil
		}
		return v.Categories
	}).(pulumi.StringArrayOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) ResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesRelated) []string {
		if v == nil {
			return nil
		}
		return v.ResourceTypes
	}).(pulumi.StringArrayOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) Solutions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesRelated) []string {
		if v == nil {
			return nil
		}
		return v.Solutions
	}).(pulumi.StringArrayOutput)
}

type LogAnalyticsQueryPackQueryPropertiesResponseRelated struct {
	Categories    []string `pulumi:"categories"`
	ResourceTypes []string `pulumi:"resourceTypes"`
	Solutions     []string `pulumi:"solutions"`
}

type LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput struct{ *pulumi.OutputState }

func (LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsQueryPackQueryPropertiesResponseRelated)(nil)).Elem()
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput) ToLogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput() LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput) ToLogAnalyticsQueryPackQueryPropertiesResponseRelatedOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogAnalyticsQueryPackQueryPropertiesResponseRelated) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput) ResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogAnalyticsQueryPackQueryPropertiesResponseRelated) []string { return v.ResourceTypes }).(pulumi.StringArrayOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput) Solutions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogAnalyticsQueryPackQueryPropertiesResponseRelated) []string { return v.Solutions }).(pulumi.StringArrayOutput)
}

type LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput struct{ *pulumi.OutputState }

func (LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsQueryPackQueryPropertiesResponseRelated)(nil)).Elem()
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) ToLogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput() LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) ToLogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) Elem() LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesResponseRelated) LogAnalyticsQueryPackQueryPropertiesResponseRelated {
		if v != nil {
			return *v
		}
		var ret LogAnalyticsQueryPackQueryPropertiesResponseRelated
		return ret
	}).(LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesResponseRelated) []string {
		if v == nil {
			return nil
		}
		return v.Categories
	}).(pulumi.StringArrayOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) ResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesResponseRelated) []string {
		if v == nil {
			return nil
		}
		return v.ResourceTypes
	}).(pulumi.StringArrayOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) Solutions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesResponseRelated) []string {
		if v == nil {
			return nil
		}
		return v.Solutions
	}).(pulumi.StringArrayOutput)
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
	pulumi.RegisterOutputType(LogAnalyticsQueryPackQueryPropertiesRelatedOutput{})
	pulumi.RegisterOutputType(LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput{})
	pulumi.RegisterOutputType(LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}

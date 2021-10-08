


package v20200701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataPoolLocation struct {
	Name string `pulumi:"name"`
}





type DataPoolLocationInput interface {
	pulumi.Input

	ToDataPoolLocationOutput() DataPoolLocationOutput
	ToDataPoolLocationOutputWithContext(context.Context) DataPoolLocationOutput
}

type DataPoolLocationArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (DataPoolLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPoolLocation)(nil)).Elem()
}

func (i DataPoolLocationArgs) ToDataPoolLocationOutput() DataPoolLocationOutput {
	return i.ToDataPoolLocationOutputWithContext(context.Background())
}

func (i DataPoolLocationArgs) ToDataPoolLocationOutputWithContext(ctx context.Context) DataPoolLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPoolLocationOutput)
}





type DataPoolLocationArrayInput interface {
	pulumi.Input

	ToDataPoolLocationArrayOutput() DataPoolLocationArrayOutput
	ToDataPoolLocationArrayOutputWithContext(context.Context) DataPoolLocationArrayOutput
}

type DataPoolLocationArray []DataPoolLocationInput

func (DataPoolLocationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataPoolLocation)(nil)).Elem()
}

func (i DataPoolLocationArray) ToDataPoolLocationArrayOutput() DataPoolLocationArrayOutput {
	return i.ToDataPoolLocationArrayOutputWithContext(context.Background())
}

func (i DataPoolLocationArray) ToDataPoolLocationArrayOutputWithContext(ctx context.Context) DataPoolLocationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPoolLocationArrayOutput)
}

type DataPoolLocationOutput struct{ *pulumi.OutputState }

func (DataPoolLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPoolLocation)(nil)).Elem()
}

func (o DataPoolLocationOutput) ToDataPoolLocationOutput() DataPoolLocationOutput {
	return o
}

func (o DataPoolLocationOutput) ToDataPoolLocationOutputWithContext(ctx context.Context) DataPoolLocationOutput {
	return o
}

func (o DataPoolLocationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DataPoolLocation) string { return v.Name }).(pulumi.StringOutput)
}

type DataPoolLocationArrayOutput struct{ *pulumi.OutputState }

func (DataPoolLocationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataPoolLocation)(nil)).Elem()
}

func (o DataPoolLocationArrayOutput) ToDataPoolLocationArrayOutput() DataPoolLocationArrayOutput {
	return o
}

func (o DataPoolLocationArrayOutput) ToDataPoolLocationArrayOutputWithContext(ctx context.Context) DataPoolLocationArrayOutput {
	return o
}

func (o DataPoolLocationArrayOutput) Index(i pulumi.IntInput) DataPoolLocationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataPoolLocation {
		return vs[0].([]DataPoolLocation)[vs[1].(int)]
	}).(DataPoolLocationOutput)
}

type DataPoolLocationResponse struct {
	Name string `pulumi:"name"`
}





type DataPoolLocationResponseInput interface {
	pulumi.Input

	ToDataPoolLocationResponseOutput() DataPoolLocationResponseOutput
	ToDataPoolLocationResponseOutputWithContext(context.Context) DataPoolLocationResponseOutput
}

type DataPoolLocationResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (DataPoolLocationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPoolLocationResponse)(nil)).Elem()
}

func (i DataPoolLocationResponseArgs) ToDataPoolLocationResponseOutput() DataPoolLocationResponseOutput {
	return i.ToDataPoolLocationResponseOutputWithContext(context.Background())
}

func (i DataPoolLocationResponseArgs) ToDataPoolLocationResponseOutputWithContext(ctx context.Context) DataPoolLocationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPoolLocationResponseOutput)
}





type DataPoolLocationResponseArrayInput interface {
	pulumi.Input

	ToDataPoolLocationResponseArrayOutput() DataPoolLocationResponseArrayOutput
	ToDataPoolLocationResponseArrayOutputWithContext(context.Context) DataPoolLocationResponseArrayOutput
}

type DataPoolLocationResponseArray []DataPoolLocationResponseInput

func (DataPoolLocationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataPoolLocationResponse)(nil)).Elem()
}

func (i DataPoolLocationResponseArray) ToDataPoolLocationResponseArrayOutput() DataPoolLocationResponseArrayOutput {
	return i.ToDataPoolLocationResponseArrayOutputWithContext(context.Background())
}

func (i DataPoolLocationResponseArray) ToDataPoolLocationResponseArrayOutputWithContext(ctx context.Context) DataPoolLocationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPoolLocationResponseArrayOutput)
}

type DataPoolLocationResponseOutput struct{ *pulumi.OutputState }

func (DataPoolLocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPoolLocationResponse)(nil)).Elem()
}

func (o DataPoolLocationResponseOutput) ToDataPoolLocationResponseOutput() DataPoolLocationResponseOutput {
	return o
}

func (o DataPoolLocationResponseOutput) ToDataPoolLocationResponseOutputWithContext(ctx context.Context) DataPoolLocationResponseOutput {
	return o
}

func (o DataPoolLocationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DataPoolLocationResponse) string { return v.Name }).(pulumi.StringOutput)
}

type DataPoolLocationResponseArrayOutput struct{ *pulumi.OutputState }

func (DataPoolLocationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataPoolLocationResponse)(nil)).Elem()
}

func (o DataPoolLocationResponseArrayOutput) ToDataPoolLocationResponseArrayOutput() DataPoolLocationResponseArrayOutput {
	return o
}

func (o DataPoolLocationResponseArrayOutput) ToDataPoolLocationResponseArrayOutputWithContext(ctx context.Context) DataPoolLocationResponseArrayOutput {
	return o
}

func (o DataPoolLocationResponseArrayOutput) Index(i pulumi.IntInput) DataPoolLocationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataPoolLocationResponse {
		return vs[0].([]DataPoolLocationResponse)[vs[1].(int)]
	}).(DataPoolLocationResponseOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
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

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
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

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataPoolLocationInput)(nil)).Elem(), DataPoolLocationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataPoolLocationArrayInput)(nil)).Elem(), DataPoolLocationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataPoolLocationResponseInput)(nil)).Elem(), DataPoolLocationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataPoolLocationResponseArrayInput)(nil)).Elem(), DataPoolLocationResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterOutputType(DataPoolLocationOutput{})
	pulumi.RegisterOutputType(DataPoolLocationArrayOutput{})
	pulumi.RegisterOutputType(DataPoolLocationResponseOutput{})
	pulumi.RegisterOutputType(DataPoolLocationResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}

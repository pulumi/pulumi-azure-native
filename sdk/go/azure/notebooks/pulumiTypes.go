


package notebooks

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NotebookResourceSystemData struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type NotebookResourceSystemDataInput interface {
	pulumi.Input

	ToNotebookResourceSystemDataOutput() NotebookResourceSystemDataOutput
	ToNotebookResourceSystemDataOutputWithContext(context.Context) NotebookResourceSystemDataOutput
}

type NotebookResourceSystemDataArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (NotebookResourceSystemDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookResourceSystemData)(nil)).Elem()
}

func (i NotebookResourceSystemDataArgs) ToNotebookResourceSystemDataOutput() NotebookResourceSystemDataOutput {
	return i.ToNotebookResourceSystemDataOutputWithContext(context.Background())
}

func (i NotebookResourceSystemDataArgs) ToNotebookResourceSystemDataOutputWithContext(ctx context.Context) NotebookResourceSystemDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookResourceSystemDataOutput)
}

func (i NotebookResourceSystemDataArgs) ToNotebookResourceSystemDataPtrOutput() NotebookResourceSystemDataPtrOutput {
	return i.ToNotebookResourceSystemDataPtrOutputWithContext(context.Background())
}

func (i NotebookResourceSystemDataArgs) ToNotebookResourceSystemDataPtrOutputWithContext(ctx context.Context) NotebookResourceSystemDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookResourceSystemDataOutput).ToNotebookResourceSystemDataPtrOutputWithContext(ctx)
}









type NotebookResourceSystemDataPtrInput interface {
	pulumi.Input

	ToNotebookResourceSystemDataPtrOutput() NotebookResourceSystemDataPtrOutput
	ToNotebookResourceSystemDataPtrOutputWithContext(context.Context) NotebookResourceSystemDataPtrOutput
}

type notebookResourceSystemDataPtrType NotebookResourceSystemDataArgs

func NotebookResourceSystemDataPtr(v *NotebookResourceSystemDataArgs) NotebookResourceSystemDataPtrInput {
	return (*notebookResourceSystemDataPtrType)(v)
}

func (*notebookResourceSystemDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookResourceSystemData)(nil)).Elem()
}

func (i *notebookResourceSystemDataPtrType) ToNotebookResourceSystemDataPtrOutput() NotebookResourceSystemDataPtrOutput {
	return i.ToNotebookResourceSystemDataPtrOutputWithContext(context.Background())
}

func (i *notebookResourceSystemDataPtrType) ToNotebookResourceSystemDataPtrOutputWithContext(ctx context.Context) NotebookResourceSystemDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookResourceSystemDataPtrOutput)
}

type NotebookResourceSystemDataOutput struct{ *pulumi.OutputState }

func (NotebookResourceSystemDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookResourceSystemData)(nil)).Elem()
}

func (o NotebookResourceSystemDataOutput) ToNotebookResourceSystemDataOutput() NotebookResourceSystemDataOutput {
	return o
}

func (o NotebookResourceSystemDataOutput) ToNotebookResourceSystemDataOutputWithContext(ctx context.Context) NotebookResourceSystemDataOutput {
	return o
}

func (o NotebookResourceSystemDataOutput) ToNotebookResourceSystemDataPtrOutput() NotebookResourceSystemDataPtrOutput {
	return o.ToNotebookResourceSystemDataPtrOutputWithContext(context.Background())
}

func (o NotebookResourceSystemDataOutput) ToNotebookResourceSystemDataPtrOutputWithContext(ctx context.Context) NotebookResourceSystemDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotebookResourceSystemData) *NotebookResourceSystemData {
		return &v
	}).(NotebookResourceSystemDataPtrOutput)
}

func (o NotebookResourceSystemDataOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotebookResourceSystemData) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o NotebookResourceSystemDataOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotebookResourceSystemData) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o NotebookResourceSystemDataOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotebookResourceSystemData) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o NotebookResourceSystemDataOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotebookResourceSystemData) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o NotebookResourceSystemDataOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotebookResourceSystemData) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o NotebookResourceSystemDataOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotebookResourceSystemData) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type NotebookResourceSystemDataPtrOutput struct{ *pulumi.OutputState }

func (NotebookResourceSystemDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookResourceSystemData)(nil)).Elem()
}

func (o NotebookResourceSystemDataPtrOutput) ToNotebookResourceSystemDataPtrOutput() NotebookResourceSystemDataPtrOutput {
	return o
}

func (o NotebookResourceSystemDataPtrOutput) ToNotebookResourceSystemDataPtrOutputWithContext(ctx context.Context) NotebookResourceSystemDataPtrOutput {
	return o
}

func (o NotebookResourceSystemDataPtrOutput) Elem() NotebookResourceSystemDataOutput {
	return o.ApplyT(func(v *NotebookResourceSystemData) NotebookResourceSystemData {
		if v != nil {
			return *v
		}
		var ret NotebookResourceSystemData
		return ret
	}).(NotebookResourceSystemDataOutput)
}

func (o NotebookResourceSystemDataPtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookResourceSystemData) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o NotebookResourceSystemDataPtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookResourceSystemData) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o NotebookResourceSystemDataPtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookResourceSystemData) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o NotebookResourceSystemDataPtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookResourceSystemData) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o NotebookResourceSystemDataPtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookResourceSystemData) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o NotebookResourceSystemDataPtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookResourceSystemData) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

type NotebookResourceSystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type NotebookResourceSystemDataResponseInput interface {
	pulumi.Input

	ToNotebookResourceSystemDataResponseOutput() NotebookResourceSystemDataResponseOutput
	ToNotebookResourceSystemDataResponseOutputWithContext(context.Context) NotebookResourceSystemDataResponseOutput
}

type NotebookResourceSystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (NotebookResourceSystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookResourceSystemDataResponse)(nil)).Elem()
}

func (i NotebookResourceSystemDataResponseArgs) ToNotebookResourceSystemDataResponseOutput() NotebookResourceSystemDataResponseOutput {
	return i.ToNotebookResourceSystemDataResponseOutputWithContext(context.Background())
}

func (i NotebookResourceSystemDataResponseArgs) ToNotebookResourceSystemDataResponseOutputWithContext(ctx context.Context) NotebookResourceSystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookResourceSystemDataResponseOutput)
}

func (i NotebookResourceSystemDataResponseArgs) ToNotebookResourceSystemDataResponsePtrOutput() NotebookResourceSystemDataResponsePtrOutput {
	return i.ToNotebookResourceSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i NotebookResourceSystemDataResponseArgs) ToNotebookResourceSystemDataResponsePtrOutputWithContext(ctx context.Context) NotebookResourceSystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookResourceSystemDataResponseOutput).ToNotebookResourceSystemDataResponsePtrOutputWithContext(ctx)
}









type NotebookResourceSystemDataResponsePtrInput interface {
	pulumi.Input

	ToNotebookResourceSystemDataResponsePtrOutput() NotebookResourceSystemDataResponsePtrOutput
	ToNotebookResourceSystemDataResponsePtrOutputWithContext(context.Context) NotebookResourceSystemDataResponsePtrOutput
}

type notebookResourceSystemDataResponsePtrType NotebookResourceSystemDataResponseArgs

func NotebookResourceSystemDataResponsePtr(v *NotebookResourceSystemDataResponseArgs) NotebookResourceSystemDataResponsePtrInput {
	return (*notebookResourceSystemDataResponsePtrType)(v)
}

func (*notebookResourceSystemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookResourceSystemDataResponse)(nil)).Elem()
}

func (i *notebookResourceSystemDataResponsePtrType) ToNotebookResourceSystemDataResponsePtrOutput() NotebookResourceSystemDataResponsePtrOutput {
	return i.ToNotebookResourceSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *notebookResourceSystemDataResponsePtrType) ToNotebookResourceSystemDataResponsePtrOutputWithContext(ctx context.Context) NotebookResourceSystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookResourceSystemDataResponsePtrOutput)
}

type NotebookResourceSystemDataResponseOutput struct{ *pulumi.OutputState }

func (NotebookResourceSystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookResourceSystemDataResponse)(nil)).Elem()
}

func (o NotebookResourceSystemDataResponseOutput) ToNotebookResourceSystemDataResponseOutput() NotebookResourceSystemDataResponseOutput {
	return o
}

func (o NotebookResourceSystemDataResponseOutput) ToNotebookResourceSystemDataResponseOutputWithContext(ctx context.Context) NotebookResourceSystemDataResponseOutput {
	return o
}

func (o NotebookResourceSystemDataResponseOutput) ToNotebookResourceSystemDataResponsePtrOutput() NotebookResourceSystemDataResponsePtrOutput {
	return o.ToNotebookResourceSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o NotebookResourceSystemDataResponseOutput) ToNotebookResourceSystemDataResponsePtrOutputWithContext(ctx context.Context) NotebookResourceSystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotebookResourceSystemDataResponse) *NotebookResourceSystemDataResponse {
		return &v
	}).(NotebookResourceSystemDataResponsePtrOutput)
}

func (o NotebookResourceSystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotebookResourceSystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o NotebookResourceSystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotebookResourceSystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o NotebookResourceSystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotebookResourceSystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o NotebookResourceSystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotebookResourceSystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o NotebookResourceSystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotebookResourceSystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o NotebookResourceSystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotebookResourceSystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type NotebookResourceSystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (NotebookResourceSystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookResourceSystemDataResponse)(nil)).Elem()
}

func (o NotebookResourceSystemDataResponsePtrOutput) ToNotebookResourceSystemDataResponsePtrOutput() NotebookResourceSystemDataResponsePtrOutput {
	return o
}

func (o NotebookResourceSystemDataResponsePtrOutput) ToNotebookResourceSystemDataResponsePtrOutputWithContext(ctx context.Context) NotebookResourceSystemDataResponsePtrOutput {
	return o
}

func (o NotebookResourceSystemDataResponsePtrOutput) Elem() NotebookResourceSystemDataResponseOutput {
	return o.ApplyT(func(v *NotebookResourceSystemDataResponse) NotebookResourceSystemDataResponse {
		if v != nil {
			return *v
		}
		var ret NotebookResourceSystemDataResponse
		return ret
	}).(NotebookResourceSystemDataResponseOutput)
}

func (o NotebookResourceSystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookResourceSystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o NotebookResourceSystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookResourceSystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o NotebookResourceSystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookResourceSystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o NotebookResourceSystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookResourceSystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o NotebookResourceSystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookResourceSystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o NotebookResourceSystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookResourceSystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NotebookResourceSystemDataOutput{})
	pulumi.RegisterOutputType(NotebookResourceSystemDataPtrOutput{})
	pulumi.RegisterOutputType(NotebookResourceSystemDataResponseOutput{})
	pulumi.RegisterOutputType(NotebookResourceSystemDataResponsePtrOutput{})
}

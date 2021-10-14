


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AttestationEvidence struct {
	Description *string `pulumi:"description"`
	SourceUri   *string `pulumi:"sourceUri"`
}





type AttestationEvidenceInput interface {
	pulumi.Input

	ToAttestationEvidenceOutput() AttestationEvidenceOutput
	ToAttestationEvidenceOutputWithContext(context.Context) AttestationEvidenceOutput
}

type AttestationEvidenceArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	SourceUri   pulumi.StringPtrInput `pulumi:"sourceUri"`
}

func (AttestationEvidenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AttestationEvidence)(nil)).Elem()
}

func (i AttestationEvidenceArgs) ToAttestationEvidenceOutput() AttestationEvidenceOutput {
	return i.ToAttestationEvidenceOutputWithContext(context.Background())
}

func (i AttestationEvidenceArgs) ToAttestationEvidenceOutputWithContext(ctx context.Context) AttestationEvidenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestationEvidenceOutput)
}





type AttestationEvidenceArrayInput interface {
	pulumi.Input

	ToAttestationEvidenceArrayOutput() AttestationEvidenceArrayOutput
	ToAttestationEvidenceArrayOutputWithContext(context.Context) AttestationEvidenceArrayOutput
}

type AttestationEvidenceArray []AttestationEvidenceInput

func (AttestationEvidenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AttestationEvidence)(nil)).Elem()
}

func (i AttestationEvidenceArray) ToAttestationEvidenceArrayOutput() AttestationEvidenceArrayOutput {
	return i.ToAttestationEvidenceArrayOutputWithContext(context.Background())
}

func (i AttestationEvidenceArray) ToAttestationEvidenceArrayOutputWithContext(ctx context.Context) AttestationEvidenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestationEvidenceArrayOutput)
}

type AttestationEvidenceOutput struct{ *pulumi.OutputState }

func (AttestationEvidenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttestationEvidence)(nil)).Elem()
}

func (o AttestationEvidenceOutput) ToAttestationEvidenceOutput() AttestationEvidenceOutput {
	return o
}

func (o AttestationEvidenceOutput) ToAttestationEvidenceOutputWithContext(ctx context.Context) AttestationEvidenceOutput {
	return o
}

func (o AttestationEvidenceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AttestationEvidence) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AttestationEvidenceOutput) SourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AttestationEvidence) *string { return v.SourceUri }).(pulumi.StringPtrOutput)
}

type AttestationEvidenceArrayOutput struct{ *pulumi.OutputState }

func (AttestationEvidenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AttestationEvidence)(nil)).Elem()
}

func (o AttestationEvidenceArrayOutput) ToAttestationEvidenceArrayOutput() AttestationEvidenceArrayOutput {
	return o
}

func (o AttestationEvidenceArrayOutput) ToAttestationEvidenceArrayOutputWithContext(ctx context.Context) AttestationEvidenceArrayOutput {
	return o
}

func (o AttestationEvidenceArrayOutput) Index(i pulumi.IntInput) AttestationEvidenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AttestationEvidence {
		return vs[0].([]AttestationEvidence)[vs[1].(int)]
	}).(AttestationEvidenceOutput)
}

type AttestationEvidenceResponse struct {
	Description *string `pulumi:"description"`
	SourceUri   *string `pulumi:"sourceUri"`
}





type AttestationEvidenceResponseInput interface {
	pulumi.Input

	ToAttestationEvidenceResponseOutput() AttestationEvidenceResponseOutput
	ToAttestationEvidenceResponseOutputWithContext(context.Context) AttestationEvidenceResponseOutput
}

type AttestationEvidenceResponseArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	SourceUri   pulumi.StringPtrInput `pulumi:"sourceUri"`
}

func (AttestationEvidenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AttestationEvidenceResponse)(nil)).Elem()
}

func (i AttestationEvidenceResponseArgs) ToAttestationEvidenceResponseOutput() AttestationEvidenceResponseOutput {
	return i.ToAttestationEvidenceResponseOutputWithContext(context.Background())
}

func (i AttestationEvidenceResponseArgs) ToAttestationEvidenceResponseOutputWithContext(ctx context.Context) AttestationEvidenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestationEvidenceResponseOutput)
}





type AttestationEvidenceResponseArrayInput interface {
	pulumi.Input

	ToAttestationEvidenceResponseArrayOutput() AttestationEvidenceResponseArrayOutput
	ToAttestationEvidenceResponseArrayOutputWithContext(context.Context) AttestationEvidenceResponseArrayOutput
}

type AttestationEvidenceResponseArray []AttestationEvidenceResponseInput

func (AttestationEvidenceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AttestationEvidenceResponse)(nil)).Elem()
}

func (i AttestationEvidenceResponseArray) ToAttestationEvidenceResponseArrayOutput() AttestationEvidenceResponseArrayOutput {
	return i.ToAttestationEvidenceResponseArrayOutputWithContext(context.Background())
}

func (i AttestationEvidenceResponseArray) ToAttestationEvidenceResponseArrayOutputWithContext(ctx context.Context) AttestationEvidenceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestationEvidenceResponseArrayOutput)
}

type AttestationEvidenceResponseOutput struct{ *pulumi.OutputState }

func (AttestationEvidenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttestationEvidenceResponse)(nil)).Elem()
}

func (o AttestationEvidenceResponseOutput) ToAttestationEvidenceResponseOutput() AttestationEvidenceResponseOutput {
	return o
}

func (o AttestationEvidenceResponseOutput) ToAttestationEvidenceResponseOutputWithContext(ctx context.Context) AttestationEvidenceResponseOutput {
	return o
}

func (o AttestationEvidenceResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AttestationEvidenceResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AttestationEvidenceResponseOutput) SourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AttestationEvidenceResponse) *string { return v.SourceUri }).(pulumi.StringPtrOutput)
}

type AttestationEvidenceResponseArrayOutput struct{ *pulumi.OutputState }

func (AttestationEvidenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AttestationEvidenceResponse)(nil)).Elem()
}

func (o AttestationEvidenceResponseArrayOutput) ToAttestationEvidenceResponseArrayOutput() AttestationEvidenceResponseArrayOutput {
	return o
}

func (o AttestationEvidenceResponseArrayOutput) ToAttestationEvidenceResponseArrayOutputWithContext(ctx context.Context) AttestationEvidenceResponseArrayOutput {
	return o
}

func (o AttestationEvidenceResponseArrayOutput) Index(i pulumi.IntInput) AttestationEvidenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AttestationEvidenceResponse {
		return vs[0].([]AttestationEvidenceResponse)[vs[1].(int)]
	}).(AttestationEvidenceResponseOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*AttestationEvidenceInput)(nil)).Elem(), AttestationEvidenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AttestationEvidenceArrayInput)(nil)).Elem(), AttestationEvidenceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AttestationEvidenceResponseInput)(nil)).Elem(), AttestationEvidenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AttestationEvidenceResponseArrayInput)(nil)).Elem(), AttestationEvidenceResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterOutputType(AttestationEvidenceOutput{})
	pulumi.RegisterOutputType(AttestationEvidenceArrayOutput{})
	pulumi.RegisterOutputType(AttestationEvidenceResponseOutput{})
	pulumi.RegisterOutputType(AttestationEvidenceResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}

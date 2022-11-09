


package syntex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DocumentProcessorProperties struct {
	SpoTenantId  string `pulumi:"spoTenantId"`
	SpoTenantUrl string `pulumi:"spoTenantUrl"`
}





type DocumentProcessorPropertiesInput interface {
	pulumi.Input

	ToDocumentProcessorPropertiesOutput() DocumentProcessorPropertiesOutput
	ToDocumentProcessorPropertiesOutputWithContext(context.Context) DocumentProcessorPropertiesOutput
}

type DocumentProcessorPropertiesArgs struct {
	SpoTenantId  pulumi.StringInput `pulumi:"spoTenantId"`
	SpoTenantUrl pulumi.StringInput `pulumi:"spoTenantUrl"`
}

func (DocumentProcessorPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DocumentProcessorProperties)(nil)).Elem()
}

func (i DocumentProcessorPropertiesArgs) ToDocumentProcessorPropertiesOutput() DocumentProcessorPropertiesOutput {
	return i.ToDocumentProcessorPropertiesOutputWithContext(context.Background())
}

func (i DocumentProcessorPropertiesArgs) ToDocumentProcessorPropertiesOutputWithContext(ctx context.Context) DocumentProcessorPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentProcessorPropertiesOutput)
}

func (i DocumentProcessorPropertiesArgs) ToDocumentProcessorPropertiesPtrOutput() DocumentProcessorPropertiesPtrOutput {
	return i.ToDocumentProcessorPropertiesPtrOutputWithContext(context.Background())
}

func (i DocumentProcessorPropertiesArgs) ToDocumentProcessorPropertiesPtrOutputWithContext(ctx context.Context) DocumentProcessorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentProcessorPropertiesOutput).ToDocumentProcessorPropertiesPtrOutputWithContext(ctx)
}









type DocumentProcessorPropertiesPtrInput interface {
	pulumi.Input

	ToDocumentProcessorPropertiesPtrOutput() DocumentProcessorPropertiesPtrOutput
	ToDocumentProcessorPropertiesPtrOutputWithContext(context.Context) DocumentProcessorPropertiesPtrOutput
}

type documentProcessorPropertiesPtrType DocumentProcessorPropertiesArgs

func DocumentProcessorPropertiesPtr(v *DocumentProcessorPropertiesArgs) DocumentProcessorPropertiesPtrInput {
	return (*documentProcessorPropertiesPtrType)(v)
}

func (*documentProcessorPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentProcessorProperties)(nil)).Elem()
}

func (i *documentProcessorPropertiesPtrType) ToDocumentProcessorPropertiesPtrOutput() DocumentProcessorPropertiesPtrOutput {
	return i.ToDocumentProcessorPropertiesPtrOutputWithContext(context.Background())
}

func (i *documentProcessorPropertiesPtrType) ToDocumentProcessorPropertiesPtrOutputWithContext(ctx context.Context) DocumentProcessorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentProcessorPropertiesPtrOutput)
}

type DocumentProcessorPropertiesOutput struct{ *pulumi.OutputState }

func (DocumentProcessorPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DocumentProcessorProperties)(nil)).Elem()
}

func (o DocumentProcessorPropertiesOutput) ToDocumentProcessorPropertiesOutput() DocumentProcessorPropertiesOutput {
	return o
}

func (o DocumentProcessorPropertiesOutput) ToDocumentProcessorPropertiesOutputWithContext(ctx context.Context) DocumentProcessorPropertiesOutput {
	return o
}

func (o DocumentProcessorPropertiesOutput) ToDocumentProcessorPropertiesPtrOutput() DocumentProcessorPropertiesPtrOutput {
	return o.ToDocumentProcessorPropertiesPtrOutputWithContext(context.Background())
}

func (o DocumentProcessorPropertiesOutput) ToDocumentProcessorPropertiesPtrOutputWithContext(ctx context.Context) DocumentProcessorPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DocumentProcessorProperties) *DocumentProcessorProperties {
		return &v
	}).(DocumentProcessorPropertiesPtrOutput)
}

func (o DocumentProcessorPropertiesOutput) SpoTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v DocumentProcessorProperties) string { return v.SpoTenantId }).(pulumi.StringOutput)
}

func (o DocumentProcessorPropertiesOutput) SpoTenantUrl() pulumi.StringOutput {
	return o.ApplyT(func(v DocumentProcessorProperties) string { return v.SpoTenantUrl }).(pulumi.StringOutput)
}

type DocumentProcessorPropertiesPtrOutput struct{ *pulumi.OutputState }

func (DocumentProcessorPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentProcessorProperties)(nil)).Elem()
}

func (o DocumentProcessorPropertiesPtrOutput) ToDocumentProcessorPropertiesPtrOutput() DocumentProcessorPropertiesPtrOutput {
	return o
}

func (o DocumentProcessorPropertiesPtrOutput) ToDocumentProcessorPropertiesPtrOutputWithContext(ctx context.Context) DocumentProcessorPropertiesPtrOutput {
	return o
}

func (o DocumentProcessorPropertiesPtrOutput) Elem() DocumentProcessorPropertiesOutput {
	return o.ApplyT(func(v *DocumentProcessorProperties) DocumentProcessorProperties {
		if v != nil {
			return *v
		}
		var ret DocumentProcessorProperties
		return ret
	}).(DocumentProcessorPropertiesOutput)
}

func (o DocumentProcessorPropertiesPtrOutput) SpoTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DocumentProcessorProperties) *string {
		if v == nil {
			return nil
		}
		return &v.SpoTenantId
	}).(pulumi.StringPtrOutput)
}

func (o DocumentProcessorPropertiesPtrOutput) SpoTenantUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DocumentProcessorProperties) *string {
		if v == nil {
			return nil
		}
		return &v.SpoTenantUrl
	}).(pulumi.StringPtrOutput)
}

type DocumentProcessorPropertiesResponse struct {
	ProvisioningState string `pulumi:"provisioningState"`
	SpoTenantId       string `pulumi:"spoTenantId"`
	SpoTenantUrl      string `pulumi:"spoTenantUrl"`
}

type DocumentProcessorPropertiesResponseOutput struct{ *pulumi.OutputState }

func (DocumentProcessorPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DocumentProcessorPropertiesResponse)(nil)).Elem()
}

func (o DocumentProcessorPropertiesResponseOutput) ToDocumentProcessorPropertiesResponseOutput() DocumentProcessorPropertiesResponseOutput {
	return o
}

func (o DocumentProcessorPropertiesResponseOutput) ToDocumentProcessorPropertiesResponseOutputWithContext(ctx context.Context) DocumentProcessorPropertiesResponseOutput {
	return o
}

func (o DocumentProcessorPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v DocumentProcessorPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DocumentProcessorPropertiesResponseOutput) SpoTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v DocumentProcessorPropertiesResponse) string { return v.SpoTenantId }).(pulumi.StringOutput)
}

func (o DocumentProcessorPropertiesResponseOutput) SpoTenantUrl() pulumi.StringOutput {
	return o.ApplyT(func(v DocumentProcessorPropertiesResponse) string { return v.SpoTenantUrl }).(pulumi.StringOutput)
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
	pulumi.RegisterOutputType(DocumentProcessorPropertiesOutput{})
	pulumi.RegisterOutputType(DocumentProcessorPropertiesPtrOutput{})
	pulumi.RegisterOutputType(DocumentProcessorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}

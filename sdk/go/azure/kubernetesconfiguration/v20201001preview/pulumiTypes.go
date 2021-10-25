


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComplianceStatusResponse struct {
	ComplianceState   string  `pulumi:"complianceState"`
	LastConfigApplied *string `pulumi:"lastConfigApplied"`
	Message           *string `pulumi:"message"`
	MessageLevel      *string `pulumi:"messageLevel"`
}





type ComplianceStatusResponseInput interface {
	pulumi.Input

	ToComplianceStatusResponseOutput() ComplianceStatusResponseOutput
	ToComplianceStatusResponseOutputWithContext(context.Context) ComplianceStatusResponseOutput
}

type ComplianceStatusResponseArgs struct {
	ComplianceState   pulumi.StringInput    `pulumi:"complianceState"`
	LastConfigApplied pulumi.StringPtrInput `pulumi:"lastConfigApplied"`
	Message           pulumi.StringPtrInput `pulumi:"message"`
	MessageLevel      pulumi.StringPtrInput `pulumi:"messageLevel"`
}

func (ComplianceStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComplianceStatusResponse)(nil)).Elem()
}

func (i ComplianceStatusResponseArgs) ToComplianceStatusResponseOutput() ComplianceStatusResponseOutput {
	return i.ToComplianceStatusResponseOutputWithContext(context.Background())
}

func (i ComplianceStatusResponseArgs) ToComplianceStatusResponseOutputWithContext(ctx context.Context) ComplianceStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComplianceStatusResponseOutput)
}

func (i ComplianceStatusResponseArgs) ToComplianceStatusResponsePtrOutput() ComplianceStatusResponsePtrOutput {
	return i.ToComplianceStatusResponsePtrOutputWithContext(context.Background())
}

func (i ComplianceStatusResponseArgs) ToComplianceStatusResponsePtrOutputWithContext(ctx context.Context) ComplianceStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComplianceStatusResponseOutput).ToComplianceStatusResponsePtrOutputWithContext(ctx)
}









type ComplianceStatusResponsePtrInput interface {
	pulumi.Input

	ToComplianceStatusResponsePtrOutput() ComplianceStatusResponsePtrOutput
	ToComplianceStatusResponsePtrOutputWithContext(context.Context) ComplianceStatusResponsePtrOutput
}

type complianceStatusResponsePtrType ComplianceStatusResponseArgs

func ComplianceStatusResponsePtr(v *ComplianceStatusResponseArgs) ComplianceStatusResponsePtrInput {
	return (*complianceStatusResponsePtrType)(v)
}

func (*complianceStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ComplianceStatusResponse)(nil)).Elem()
}

func (i *complianceStatusResponsePtrType) ToComplianceStatusResponsePtrOutput() ComplianceStatusResponsePtrOutput {
	return i.ToComplianceStatusResponsePtrOutputWithContext(context.Background())
}

func (i *complianceStatusResponsePtrType) ToComplianceStatusResponsePtrOutputWithContext(ctx context.Context) ComplianceStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComplianceStatusResponsePtrOutput)
}

type ComplianceStatusResponseOutput struct{ *pulumi.OutputState }

func (ComplianceStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComplianceStatusResponse)(nil)).Elem()
}

func (o ComplianceStatusResponseOutput) ToComplianceStatusResponseOutput() ComplianceStatusResponseOutput {
	return o
}

func (o ComplianceStatusResponseOutput) ToComplianceStatusResponseOutputWithContext(ctx context.Context) ComplianceStatusResponseOutput {
	return o
}

func (o ComplianceStatusResponseOutput) ToComplianceStatusResponsePtrOutput() ComplianceStatusResponsePtrOutput {
	return o.ToComplianceStatusResponsePtrOutputWithContext(context.Background())
}

func (o ComplianceStatusResponseOutput) ToComplianceStatusResponsePtrOutputWithContext(ctx context.Context) ComplianceStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComplianceStatusResponse) *ComplianceStatusResponse {
		return &v
	}).(ComplianceStatusResponsePtrOutput)
}

func (o ComplianceStatusResponseOutput) ComplianceState() pulumi.StringOutput {
	return o.ApplyT(func(v ComplianceStatusResponse) string { return v.ComplianceState }).(pulumi.StringOutput)
}

func (o ComplianceStatusResponseOutput) LastConfigApplied() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComplianceStatusResponse) *string { return v.LastConfigApplied }).(pulumi.StringPtrOutput)
}

func (o ComplianceStatusResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComplianceStatusResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ComplianceStatusResponseOutput) MessageLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComplianceStatusResponse) *string { return v.MessageLevel }).(pulumi.StringPtrOutput)
}

type ComplianceStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (ComplianceStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComplianceStatusResponse)(nil)).Elem()
}

func (o ComplianceStatusResponsePtrOutput) ToComplianceStatusResponsePtrOutput() ComplianceStatusResponsePtrOutput {
	return o
}

func (o ComplianceStatusResponsePtrOutput) ToComplianceStatusResponsePtrOutputWithContext(ctx context.Context) ComplianceStatusResponsePtrOutput {
	return o
}

func (o ComplianceStatusResponsePtrOutput) Elem() ComplianceStatusResponseOutput {
	return o.ApplyT(func(v *ComplianceStatusResponse) ComplianceStatusResponse {
		if v != nil {
			return *v
		}
		var ret ComplianceStatusResponse
		return ret
	}).(ComplianceStatusResponseOutput)
}

func (o ComplianceStatusResponsePtrOutput) ComplianceState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComplianceStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ComplianceState
	}).(pulumi.StringPtrOutput)
}

func (o ComplianceStatusResponsePtrOutput) LastConfigApplied() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComplianceStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastConfigApplied
	}).(pulumi.StringPtrOutput)
}

func (o ComplianceStatusResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComplianceStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ComplianceStatusResponsePtrOutput) MessageLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComplianceStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.MessageLevel
	}).(pulumi.StringPtrOutput)
}

type HelmOperatorProperties struct {
	ChartValues  *string `pulumi:"chartValues"`
	ChartVersion *string `pulumi:"chartVersion"`
}





type HelmOperatorPropertiesInput interface {
	pulumi.Input

	ToHelmOperatorPropertiesOutput() HelmOperatorPropertiesOutput
	ToHelmOperatorPropertiesOutputWithContext(context.Context) HelmOperatorPropertiesOutput
}

type HelmOperatorPropertiesArgs struct {
	ChartValues  pulumi.StringPtrInput `pulumi:"chartValues"`
	ChartVersion pulumi.StringPtrInput `pulumi:"chartVersion"`
}

func (HelmOperatorPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HelmOperatorProperties)(nil)).Elem()
}

func (i HelmOperatorPropertiesArgs) ToHelmOperatorPropertiesOutput() HelmOperatorPropertiesOutput {
	return i.ToHelmOperatorPropertiesOutputWithContext(context.Background())
}

func (i HelmOperatorPropertiesArgs) ToHelmOperatorPropertiesOutputWithContext(ctx context.Context) HelmOperatorPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmOperatorPropertiesOutput)
}

func (i HelmOperatorPropertiesArgs) ToHelmOperatorPropertiesPtrOutput() HelmOperatorPropertiesPtrOutput {
	return i.ToHelmOperatorPropertiesPtrOutputWithContext(context.Background())
}

func (i HelmOperatorPropertiesArgs) ToHelmOperatorPropertiesPtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmOperatorPropertiesOutput).ToHelmOperatorPropertiesPtrOutputWithContext(ctx)
}









type HelmOperatorPropertiesPtrInput interface {
	pulumi.Input

	ToHelmOperatorPropertiesPtrOutput() HelmOperatorPropertiesPtrOutput
	ToHelmOperatorPropertiesPtrOutputWithContext(context.Context) HelmOperatorPropertiesPtrOutput
}

type helmOperatorPropertiesPtrType HelmOperatorPropertiesArgs

func HelmOperatorPropertiesPtr(v *HelmOperatorPropertiesArgs) HelmOperatorPropertiesPtrInput {
	return (*helmOperatorPropertiesPtrType)(v)
}

func (*helmOperatorPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HelmOperatorProperties)(nil)).Elem()
}

func (i *helmOperatorPropertiesPtrType) ToHelmOperatorPropertiesPtrOutput() HelmOperatorPropertiesPtrOutput {
	return i.ToHelmOperatorPropertiesPtrOutputWithContext(context.Background())
}

func (i *helmOperatorPropertiesPtrType) ToHelmOperatorPropertiesPtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmOperatorPropertiesPtrOutput)
}

type HelmOperatorPropertiesOutput struct{ *pulumi.OutputState }

func (HelmOperatorPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HelmOperatorProperties)(nil)).Elem()
}

func (o HelmOperatorPropertiesOutput) ToHelmOperatorPropertiesOutput() HelmOperatorPropertiesOutput {
	return o
}

func (o HelmOperatorPropertiesOutput) ToHelmOperatorPropertiesOutputWithContext(ctx context.Context) HelmOperatorPropertiesOutput {
	return o
}

func (o HelmOperatorPropertiesOutput) ToHelmOperatorPropertiesPtrOutput() HelmOperatorPropertiesPtrOutput {
	return o.ToHelmOperatorPropertiesPtrOutputWithContext(context.Background())
}

func (o HelmOperatorPropertiesOutput) ToHelmOperatorPropertiesPtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HelmOperatorProperties) *HelmOperatorProperties {
		return &v
	}).(HelmOperatorPropertiesPtrOutput)
}

func (o HelmOperatorPropertiesOutput) ChartValues() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmOperatorProperties) *string { return v.ChartValues }).(pulumi.StringPtrOutput)
}

func (o HelmOperatorPropertiesOutput) ChartVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmOperatorProperties) *string { return v.ChartVersion }).(pulumi.StringPtrOutput)
}

type HelmOperatorPropertiesPtrOutput struct{ *pulumi.OutputState }

func (HelmOperatorPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HelmOperatorProperties)(nil)).Elem()
}

func (o HelmOperatorPropertiesPtrOutput) ToHelmOperatorPropertiesPtrOutput() HelmOperatorPropertiesPtrOutput {
	return o
}

func (o HelmOperatorPropertiesPtrOutput) ToHelmOperatorPropertiesPtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesPtrOutput {
	return o
}

func (o HelmOperatorPropertiesPtrOutput) Elem() HelmOperatorPropertiesOutput {
	return o.ApplyT(func(v *HelmOperatorProperties) HelmOperatorProperties {
		if v != nil {
			return *v
		}
		var ret HelmOperatorProperties
		return ret
	}).(HelmOperatorPropertiesOutput)
}

func (o HelmOperatorPropertiesPtrOutput) ChartValues() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmOperatorProperties) *string {
		if v == nil {
			return nil
		}
		return v.ChartValues
	}).(pulumi.StringPtrOutput)
}

func (o HelmOperatorPropertiesPtrOutput) ChartVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmOperatorProperties) *string {
		if v == nil {
			return nil
		}
		return v.ChartVersion
	}).(pulumi.StringPtrOutput)
}

type HelmOperatorPropertiesResponse struct {
	ChartValues  *string `pulumi:"chartValues"`
	ChartVersion *string `pulumi:"chartVersion"`
}





type HelmOperatorPropertiesResponseInput interface {
	pulumi.Input

	ToHelmOperatorPropertiesResponseOutput() HelmOperatorPropertiesResponseOutput
	ToHelmOperatorPropertiesResponseOutputWithContext(context.Context) HelmOperatorPropertiesResponseOutput
}

type HelmOperatorPropertiesResponseArgs struct {
	ChartValues  pulumi.StringPtrInput `pulumi:"chartValues"`
	ChartVersion pulumi.StringPtrInput `pulumi:"chartVersion"`
}

func (HelmOperatorPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HelmOperatorPropertiesResponse)(nil)).Elem()
}

func (i HelmOperatorPropertiesResponseArgs) ToHelmOperatorPropertiesResponseOutput() HelmOperatorPropertiesResponseOutput {
	return i.ToHelmOperatorPropertiesResponseOutputWithContext(context.Background())
}

func (i HelmOperatorPropertiesResponseArgs) ToHelmOperatorPropertiesResponseOutputWithContext(ctx context.Context) HelmOperatorPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmOperatorPropertiesResponseOutput)
}

func (i HelmOperatorPropertiesResponseArgs) ToHelmOperatorPropertiesResponsePtrOutput() HelmOperatorPropertiesResponsePtrOutput {
	return i.ToHelmOperatorPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i HelmOperatorPropertiesResponseArgs) ToHelmOperatorPropertiesResponsePtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmOperatorPropertiesResponseOutput).ToHelmOperatorPropertiesResponsePtrOutputWithContext(ctx)
}









type HelmOperatorPropertiesResponsePtrInput interface {
	pulumi.Input

	ToHelmOperatorPropertiesResponsePtrOutput() HelmOperatorPropertiesResponsePtrOutput
	ToHelmOperatorPropertiesResponsePtrOutputWithContext(context.Context) HelmOperatorPropertiesResponsePtrOutput
}

type helmOperatorPropertiesResponsePtrType HelmOperatorPropertiesResponseArgs

func HelmOperatorPropertiesResponsePtr(v *HelmOperatorPropertiesResponseArgs) HelmOperatorPropertiesResponsePtrInput {
	return (*helmOperatorPropertiesResponsePtrType)(v)
}

func (*helmOperatorPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HelmOperatorPropertiesResponse)(nil)).Elem()
}

func (i *helmOperatorPropertiesResponsePtrType) ToHelmOperatorPropertiesResponsePtrOutput() HelmOperatorPropertiesResponsePtrOutput {
	return i.ToHelmOperatorPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *helmOperatorPropertiesResponsePtrType) ToHelmOperatorPropertiesResponsePtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmOperatorPropertiesResponsePtrOutput)
}

type HelmOperatorPropertiesResponseOutput struct{ *pulumi.OutputState }

func (HelmOperatorPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HelmOperatorPropertiesResponse)(nil)).Elem()
}

func (o HelmOperatorPropertiesResponseOutput) ToHelmOperatorPropertiesResponseOutput() HelmOperatorPropertiesResponseOutput {
	return o
}

func (o HelmOperatorPropertiesResponseOutput) ToHelmOperatorPropertiesResponseOutputWithContext(ctx context.Context) HelmOperatorPropertiesResponseOutput {
	return o
}

func (o HelmOperatorPropertiesResponseOutput) ToHelmOperatorPropertiesResponsePtrOutput() HelmOperatorPropertiesResponsePtrOutput {
	return o.ToHelmOperatorPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o HelmOperatorPropertiesResponseOutput) ToHelmOperatorPropertiesResponsePtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HelmOperatorPropertiesResponse) *HelmOperatorPropertiesResponse {
		return &v
	}).(HelmOperatorPropertiesResponsePtrOutput)
}

func (o HelmOperatorPropertiesResponseOutput) ChartValues() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmOperatorPropertiesResponse) *string { return v.ChartValues }).(pulumi.StringPtrOutput)
}

func (o HelmOperatorPropertiesResponseOutput) ChartVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmOperatorPropertiesResponse) *string { return v.ChartVersion }).(pulumi.StringPtrOutput)
}

type HelmOperatorPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (HelmOperatorPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HelmOperatorPropertiesResponse)(nil)).Elem()
}

func (o HelmOperatorPropertiesResponsePtrOutput) ToHelmOperatorPropertiesResponsePtrOutput() HelmOperatorPropertiesResponsePtrOutput {
	return o
}

func (o HelmOperatorPropertiesResponsePtrOutput) ToHelmOperatorPropertiesResponsePtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesResponsePtrOutput {
	return o
}

func (o HelmOperatorPropertiesResponsePtrOutput) Elem() HelmOperatorPropertiesResponseOutput {
	return o.ApplyT(func(v *HelmOperatorPropertiesResponse) HelmOperatorPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret HelmOperatorPropertiesResponse
		return ret
	}).(HelmOperatorPropertiesResponseOutput)
}

func (o HelmOperatorPropertiesResponsePtrOutput) ChartValues() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmOperatorPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ChartValues
	}).(pulumi.StringPtrOutput)
}

func (o HelmOperatorPropertiesResponsePtrOutput) ChartVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmOperatorPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ChartVersion
	}).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          string `pulumi:"createdAt"`
	CreatedBy          string `pulumi:"createdBy"`
	CreatedByType      string `pulumi:"createdByType"`
	LastModifiedAt     string `pulumi:"lastModifiedAt"`
	LastModifiedBy     string `pulumi:"lastModifiedBy"`
	LastModifiedByType string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringInput `pulumi:"lastModifiedByType"`
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

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.CreatedBy }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.CreatedByType }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.LastModifiedAt }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.LastModifiedBy }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.LastModifiedByType }).(pulumi.StringOutput)
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
		return &v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComplianceStatusResponseInput)(nil)).Elem(), ComplianceStatusResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComplianceStatusResponsePtrInput)(nil)).Elem(), ComplianceStatusResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HelmOperatorPropertiesInput)(nil)).Elem(), HelmOperatorPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HelmOperatorPropertiesPtrInput)(nil)).Elem(), HelmOperatorPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HelmOperatorPropertiesResponseInput)(nil)).Elem(), HelmOperatorPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HelmOperatorPropertiesResponsePtrInput)(nil)).Elem(), HelmOperatorPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterOutputType(ComplianceStatusResponseOutput{})
	pulumi.RegisterOutputType(ComplianceStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(HelmOperatorPropertiesOutput{})
	pulumi.RegisterOutputType(HelmOperatorPropertiesPtrOutput{})
	pulumi.RegisterOutputType(HelmOperatorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(HelmOperatorPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}

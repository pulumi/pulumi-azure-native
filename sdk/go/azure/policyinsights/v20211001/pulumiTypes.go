


package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ErrorDefinitionResponse struct {
	AdditionalInfo []TypedErrorInfoResponse  `pulumi:"additionalInfo"`
	Code           string                    `pulumi:"code"`
	Details        []ErrorDefinitionResponse `pulumi:"details"`
	Message        string                    `pulumi:"message"`
	Target         string                    `pulumi:"target"`
}

type RemediationDeploymentResponse struct {
	CreatedOn            string                  `pulumi:"createdOn"`
	DeploymentId         string                  `pulumi:"deploymentId"`
	Error                ErrorDefinitionResponse `pulumi:"error"`
	LastUpdatedOn        string                  `pulumi:"lastUpdatedOn"`
	RemediatedResourceId string                  `pulumi:"remediatedResourceId"`
	ResourceLocation     string                  `pulumi:"resourceLocation"`
	Status               string                  `pulumi:"status"`
}

type RemediationDeploymentSummaryResponse struct {
	FailedDeployments     int `pulumi:"failedDeployments"`
	SuccessfulDeployments int `pulumi:"successfulDeployments"`
	TotalDeployments      int `pulumi:"totalDeployments"`
}

type RemediationDeploymentSummaryResponseOutput struct{ *pulumi.OutputState }

func (RemediationDeploymentSummaryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationDeploymentSummaryResponse)(nil)).Elem()
}

func (o RemediationDeploymentSummaryResponseOutput) ToRemediationDeploymentSummaryResponseOutput() RemediationDeploymentSummaryResponseOutput {
	return o
}

func (o RemediationDeploymentSummaryResponseOutput) ToRemediationDeploymentSummaryResponseOutputWithContext(ctx context.Context) RemediationDeploymentSummaryResponseOutput {
	return o
}

func (o RemediationDeploymentSummaryResponseOutput) FailedDeployments() pulumi.IntOutput {
	return o.ApplyT(func(v RemediationDeploymentSummaryResponse) int { return v.FailedDeployments }).(pulumi.IntOutput)
}

func (o RemediationDeploymentSummaryResponseOutput) SuccessfulDeployments() pulumi.IntOutput {
	return o.ApplyT(func(v RemediationDeploymentSummaryResponse) int { return v.SuccessfulDeployments }).(pulumi.IntOutput)
}

func (o RemediationDeploymentSummaryResponseOutput) TotalDeployments() pulumi.IntOutput {
	return o.ApplyT(func(v RemediationDeploymentSummaryResponse) int { return v.TotalDeployments }).(pulumi.IntOutput)
}

type RemediationFilters struct {
	Locations []string `pulumi:"locations"`
}





type RemediationFiltersInput interface {
	pulumi.Input

	ToRemediationFiltersOutput() RemediationFiltersOutput
	ToRemediationFiltersOutputWithContext(context.Context) RemediationFiltersOutput
}

type RemediationFiltersArgs struct {
	Locations pulumi.StringArrayInput `pulumi:"locations"`
}

func (RemediationFiltersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationFilters)(nil)).Elem()
}

func (i RemediationFiltersArgs) ToRemediationFiltersOutput() RemediationFiltersOutput {
	return i.ToRemediationFiltersOutputWithContext(context.Background())
}

func (i RemediationFiltersArgs) ToRemediationFiltersOutputWithContext(ctx context.Context) RemediationFiltersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationFiltersOutput)
}

func (i RemediationFiltersArgs) ToRemediationFiltersPtrOutput() RemediationFiltersPtrOutput {
	return i.ToRemediationFiltersPtrOutputWithContext(context.Background())
}

func (i RemediationFiltersArgs) ToRemediationFiltersPtrOutputWithContext(ctx context.Context) RemediationFiltersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationFiltersOutput).ToRemediationFiltersPtrOutputWithContext(ctx)
}









type RemediationFiltersPtrInput interface {
	pulumi.Input

	ToRemediationFiltersPtrOutput() RemediationFiltersPtrOutput
	ToRemediationFiltersPtrOutputWithContext(context.Context) RemediationFiltersPtrOutput
}

type remediationFiltersPtrType RemediationFiltersArgs

func RemediationFiltersPtr(v *RemediationFiltersArgs) RemediationFiltersPtrInput {
	return (*remediationFiltersPtrType)(v)
}

func (*remediationFiltersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationFilters)(nil)).Elem()
}

func (i *remediationFiltersPtrType) ToRemediationFiltersPtrOutput() RemediationFiltersPtrOutput {
	return i.ToRemediationFiltersPtrOutputWithContext(context.Background())
}

func (i *remediationFiltersPtrType) ToRemediationFiltersPtrOutputWithContext(ctx context.Context) RemediationFiltersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationFiltersPtrOutput)
}

type RemediationFiltersOutput struct{ *pulumi.OutputState }

func (RemediationFiltersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationFilters)(nil)).Elem()
}

func (o RemediationFiltersOutput) ToRemediationFiltersOutput() RemediationFiltersOutput {
	return o
}

func (o RemediationFiltersOutput) ToRemediationFiltersOutputWithContext(ctx context.Context) RemediationFiltersOutput {
	return o
}

func (o RemediationFiltersOutput) ToRemediationFiltersPtrOutput() RemediationFiltersPtrOutput {
	return o.ToRemediationFiltersPtrOutputWithContext(context.Background())
}

func (o RemediationFiltersOutput) ToRemediationFiltersPtrOutputWithContext(ctx context.Context) RemediationFiltersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RemediationFilters) *RemediationFilters {
		return &v
	}).(RemediationFiltersPtrOutput)
}

func (o RemediationFiltersOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RemediationFilters) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

type RemediationFiltersPtrOutput struct{ *pulumi.OutputState }

func (RemediationFiltersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationFilters)(nil)).Elem()
}

func (o RemediationFiltersPtrOutput) ToRemediationFiltersPtrOutput() RemediationFiltersPtrOutput {
	return o
}

func (o RemediationFiltersPtrOutput) ToRemediationFiltersPtrOutputWithContext(ctx context.Context) RemediationFiltersPtrOutput {
	return o
}

func (o RemediationFiltersPtrOutput) Elem() RemediationFiltersOutput {
	return o.ApplyT(func(v *RemediationFilters) RemediationFilters {
		if v != nil {
			return *v
		}
		var ret RemediationFilters
		return ret
	}).(RemediationFiltersOutput)
}

func (o RemediationFiltersPtrOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RemediationFilters) []string {
		if v == nil {
			return nil
		}
		return v.Locations
	}).(pulumi.StringArrayOutput)
}

type RemediationFiltersResponse struct {
	Locations []string `pulumi:"locations"`
}

type RemediationFiltersResponseOutput struct{ *pulumi.OutputState }

func (RemediationFiltersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationFiltersResponse)(nil)).Elem()
}

func (o RemediationFiltersResponseOutput) ToRemediationFiltersResponseOutput() RemediationFiltersResponseOutput {
	return o
}

func (o RemediationFiltersResponseOutput) ToRemediationFiltersResponseOutputWithContext(ctx context.Context) RemediationFiltersResponseOutput {
	return o
}

func (o RemediationFiltersResponseOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RemediationFiltersResponse) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

type RemediationFiltersResponsePtrOutput struct{ *pulumi.OutputState }

func (RemediationFiltersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationFiltersResponse)(nil)).Elem()
}

func (o RemediationFiltersResponsePtrOutput) ToRemediationFiltersResponsePtrOutput() RemediationFiltersResponsePtrOutput {
	return o
}

func (o RemediationFiltersResponsePtrOutput) ToRemediationFiltersResponsePtrOutputWithContext(ctx context.Context) RemediationFiltersResponsePtrOutput {
	return o
}

func (o RemediationFiltersResponsePtrOutput) Elem() RemediationFiltersResponseOutput {
	return o.ApplyT(func(v *RemediationFiltersResponse) RemediationFiltersResponse {
		if v != nil {
			return *v
		}
		var ret RemediationFiltersResponse
		return ret
	}).(RemediationFiltersResponseOutput)
}

func (o RemediationFiltersResponsePtrOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RemediationFiltersResponse) []string {
		if v == nil {
			return nil
		}
		return v.Locations
	}).(pulumi.StringArrayOutput)
}

type RemediationPropertiesFailureThreshold struct {
	Percentage *float64 `pulumi:"percentage"`
}





type RemediationPropertiesFailureThresholdInput interface {
	pulumi.Input

	ToRemediationPropertiesFailureThresholdOutput() RemediationPropertiesFailureThresholdOutput
	ToRemediationPropertiesFailureThresholdOutputWithContext(context.Context) RemediationPropertiesFailureThresholdOutput
}

type RemediationPropertiesFailureThresholdArgs struct {
	Percentage pulumi.Float64PtrInput `pulumi:"percentage"`
}

func (RemediationPropertiesFailureThresholdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationPropertiesFailureThreshold)(nil)).Elem()
}

func (i RemediationPropertiesFailureThresholdArgs) ToRemediationPropertiesFailureThresholdOutput() RemediationPropertiesFailureThresholdOutput {
	return i.ToRemediationPropertiesFailureThresholdOutputWithContext(context.Background())
}

func (i RemediationPropertiesFailureThresholdArgs) ToRemediationPropertiesFailureThresholdOutputWithContext(ctx context.Context) RemediationPropertiesFailureThresholdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationPropertiesFailureThresholdOutput)
}

func (i RemediationPropertiesFailureThresholdArgs) ToRemediationPropertiesFailureThresholdPtrOutput() RemediationPropertiesFailureThresholdPtrOutput {
	return i.ToRemediationPropertiesFailureThresholdPtrOutputWithContext(context.Background())
}

func (i RemediationPropertiesFailureThresholdArgs) ToRemediationPropertiesFailureThresholdPtrOutputWithContext(ctx context.Context) RemediationPropertiesFailureThresholdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationPropertiesFailureThresholdOutput).ToRemediationPropertiesFailureThresholdPtrOutputWithContext(ctx)
}









type RemediationPropertiesFailureThresholdPtrInput interface {
	pulumi.Input

	ToRemediationPropertiesFailureThresholdPtrOutput() RemediationPropertiesFailureThresholdPtrOutput
	ToRemediationPropertiesFailureThresholdPtrOutputWithContext(context.Context) RemediationPropertiesFailureThresholdPtrOutput
}

type remediationPropertiesFailureThresholdPtrType RemediationPropertiesFailureThresholdArgs

func RemediationPropertiesFailureThresholdPtr(v *RemediationPropertiesFailureThresholdArgs) RemediationPropertiesFailureThresholdPtrInput {
	return (*remediationPropertiesFailureThresholdPtrType)(v)
}

func (*remediationPropertiesFailureThresholdPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationPropertiesFailureThreshold)(nil)).Elem()
}

func (i *remediationPropertiesFailureThresholdPtrType) ToRemediationPropertiesFailureThresholdPtrOutput() RemediationPropertiesFailureThresholdPtrOutput {
	return i.ToRemediationPropertiesFailureThresholdPtrOutputWithContext(context.Background())
}

func (i *remediationPropertiesFailureThresholdPtrType) ToRemediationPropertiesFailureThresholdPtrOutputWithContext(ctx context.Context) RemediationPropertiesFailureThresholdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationPropertiesFailureThresholdPtrOutput)
}

type RemediationPropertiesFailureThresholdOutput struct{ *pulumi.OutputState }

func (RemediationPropertiesFailureThresholdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationPropertiesFailureThreshold)(nil)).Elem()
}

func (o RemediationPropertiesFailureThresholdOutput) ToRemediationPropertiesFailureThresholdOutput() RemediationPropertiesFailureThresholdOutput {
	return o
}

func (o RemediationPropertiesFailureThresholdOutput) ToRemediationPropertiesFailureThresholdOutputWithContext(ctx context.Context) RemediationPropertiesFailureThresholdOutput {
	return o
}

func (o RemediationPropertiesFailureThresholdOutput) ToRemediationPropertiesFailureThresholdPtrOutput() RemediationPropertiesFailureThresholdPtrOutput {
	return o.ToRemediationPropertiesFailureThresholdPtrOutputWithContext(context.Background())
}

func (o RemediationPropertiesFailureThresholdOutput) ToRemediationPropertiesFailureThresholdPtrOutputWithContext(ctx context.Context) RemediationPropertiesFailureThresholdPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RemediationPropertiesFailureThreshold) *RemediationPropertiesFailureThreshold {
		return &v
	}).(RemediationPropertiesFailureThresholdPtrOutput)
}

func (o RemediationPropertiesFailureThresholdOutput) Percentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RemediationPropertiesFailureThreshold) *float64 { return v.Percentage }).(pulumi.Float64PtrOutput)
}

type RemediationPropertiesFailureThresholdPtrOutput struct{ *pulumi.OutputState }

func (RemediationPropertiesFailureThresholdPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationPropertiesFailureThreshold)(nil)).Elem()
}

func (o RemediationPropertiesFailureThresholdPtrOutput) ToRemediationPropertiesFailureThresholdPtrOutput() RemediationPropertiesFailureThresholdPtrOutput {
	return o
}

func (o RemediationPropertiesFailureThresholdPtrOutput) ToRemediationPropertiesFailureThresholdPtrOutputWithContext(ctx context.Context) RemediationPropertiesFailureThresholdPtrOutput {
	return o
}

func (o RemediationPropertiesFailureThresholdPtrOutput) Elem() RemediationPropertiesFailureThresholdOutput {
	return o.ApplyT(func(v *RemediationPropertiesFailureThreshold) RemediationPropertiesFailureThreshold {
		if v != nil {
			return *v
		}
		var ret RemediationPropertiesFailureThreshold
		return ret
	}).(RemediationPropertiesFailureThresholdOutput)
}

func (o RemediationPropertiesFailureThresholdPtrOutput) Percentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *RemediationPropertiesFailureThreshold) *float64 {
		if v == nil {
			return nil
		}
		return v.Percentage
	}).(pulumi.Float64PtrOutput)
}

type RemediationPropertiesResponseFailureThreshold struct {
	Percentage *float64 `pulumi:"percentage"`
}

type RemediationPropertiesResponseFailureThresholdOutput struct{ *pulumi.OutputState }

func (RemediationPropertiesResponseFailureThresholdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationPropertiesResponseFailureThreshold)(nil)).Elem()
}

func (o RemediationPropertiesResponseFailureThresholdOutput) ToRemediationPropertiesResponseFailureThresholdOutput() RemediationPropertiesResponseFailureThresholdOutput {
	return o
}

func (o RemediationPropertiesResponseFailureThresholdOutput) ToRemediationPropertiesResponseFailureThresholdOutputWithContext(ctx context.Context) RemediationPropertiesResponseFailureThresholdOutput {
	return o
}

func (o RemediationPropertiesResponseFailureThresholdOutput) Percentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RemediationPropertiesResponseFailureThreshold) *float64 { return v.Percentage }).(pulumi.Float64PtrOutput)
}

type RemediationPropertiesResponseFailureThresholdPtrOutput struct{ *pulumi.OutputState }

func (RemediationPropertiesResponseFailureThresholdPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationPropertiesResponseFailureThreshold)(nil)).Elem()
}

func (o RemediationPropertiesResponseFailureThresholdPtrOutput) ToRemediationPropertiesResponseFailureThresholdPtrOutput() RemediationPropertiesResponseFailureThresholdPtrOutput {
	return o
}

func (o RemediationPropertiesResponseFailureThresholdPtrOutput) ToRemediationPropertiesResponseFailureThresholdPtrOutputWithContext(ctx context.Context) RemediationPropertiesResponseFailureThresholdPtrOutput {
	return o
}

func (o RemediationPropertiesResponseFailureThresholdPtrOutput) Elem() RemediationPropertiesResponseFailureThresholdOutput {
	return o.ApplyT(func(v *RemediationPropertiesResponseFailureThreshold) RemediationPropertiesResponseFailureThreshold {
		if v != nil {
			return *v
		}
		var ret RemediationPropertiesResponseFailureThreshold
		return ret
	}).(RemediationPropertiesResponseFailureThresholdOutput)
}

func (o RemediationPropertiesResponseFailureThresholdPtrOutput) Percentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *RemediationPropertiesResponseFailureThreshold) *float64 {
		if v == nil {
			return nil
		}
		return v.Percentage
	}).(pulumi.Float64PtrOutput)
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

type TypedErrorInfoResponse struct {
	Info interface{} `pulumi:"info"`
	Type string      `pulumi:"type"`
}

func init() {
	pulumi.RegisterOutputType(RemediationDeploymentSummaryResponseOutput{})
	pulumi.RegisterOutputType(RemediationFiltersOutput{})
	pulumi.RegisterOutputType(RemediationFiltersPtrOutput{})
	pulumi.RegisterOutputType(RemediationFiltersResponseOutput{})
	pulumi.RegisterOutputType(RemediationFiltersResponsePtrOutput{})
	pulumi.RegisterOutputType(RemediationPropertiesFailureThresholdOutput{})
	pulumi.RegisterOutputType(RemediationPropertiesFailureThresholdPtrOutput{})
	pulumi.RegisterOutputType(RemediationPropertiesResponseFailureThresholdOutput{})
	pulumi.RegisterOutputType(RemediationPropertiesResponseFailureThresholdPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}

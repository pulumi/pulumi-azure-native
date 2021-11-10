


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





type ErrorDefinitionResponseInput interface {
	pulumi.Input

	ToErrorDefinitionResponseOutput() ErrorDefinitionResponseOutput
	ToErrorDefinitionResponseOutputWithContext(context.Context) ErrorDefinitionResponseOutput
}

type ErrorDefinitionResponseArgs struct {
	AdditionalInfo TypedErrorInfoResponseArrayInput  `pulumi:"additionalInfo"`
	Code           pulumi.StringInput                `pulumi:"code"`
	Details        ErrorDefinitionResponseArrayInput `pulumi:"details"`
	Message        pulumi.StringInput                `pulumi:"message"`
	Target         pulumi.StringInput                `pulumi:"target"`
}

func (ErrorDefinitionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDefinitionResponse)(nil)).Elem()
}

func (i ErrorDefinitionResponseArgs) ToErrorDefinitionResponseOutput() ErrorDefinitionResponseOutput {
	return i.ToErrorDefinitionResponseOutputWithContext(context.Background())
}

func (i ErrorDefinitionResponseArgs) ToErrorDefinitionResponseOutputWithContext(ctx context.Context) ErrorDefinitionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDefinitionResponseOutput)
}





type ErrorDefinitionResponseArrayInput interface {
	pulumi.Input

	ToErrorDefinitionResponseArrayOutput() ErrorDefinitionResponseArrayOutput
	ToErrorDefinitionResponseArrayOutputWithContext(context.Context) ErrorDefinitionResponseArrayOutput
}

type ErrorDefinitionResponseArray []ErrorDefinitionResponseInput

func (ErrorDefinitionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDefinitionResponse)(nil)).Elem()
}

func (i ErrorDefinitionResponseArray) ToErrorDefinitionResponseArrayOutput() ErrorDefinitionResponseArrayOutput {
	return i.ToErrorDefinitionResponseArrayOutputWithContext(context.Background())
}

func (i ErrorDefinitionResponseArray) ToErrorDefinitionResponseArrayOutputWithContext(ctx context.Context) ErrorDefinitionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDefinitionResponseArrayOutput)
}

type ErrorDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ErrorDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDefinitionResponse)(nil)).Elem()
}

func (o ErrorDefinitionResponseOutput) ToErrorDefinitionResponseOutput() ErrorDefinitionResponseOutput {
	return o
}

func (o ErrorDefinitionResponseOutput) ToErrorDefinitionResponseOutputWithContext(ctx context.Context) ErrorDefinitionResponseOutput {
	return o
}

func (o ErrorDefinitionResponseOutput) AdditionalInfo() TypedErrorInfoResponseArrayOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) []TypedErrorInfoResponse { return v.AdditionalInfo }).(TypedErrorInfoResponseArrayOutput)
}

func (o ErrorDefinitionResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorDefinitionResponseOutput) Details() ErrorDefinitionResponseArrayOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) []ErrorDefinitionResponse { return v.Details }).(ErrorDefinitionResponseArrayOutput)
}

func (o ErrorDefinitionResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ErrorDefinitionResponseOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) string { return v.Target }).(pulumi.StringOutput)
}

type ErrorDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDefinitionResponse)(nil)).Elem()
}

func (o ErrorDefinitionResponseArrayOutput) ToErrorDefinitionResponseArrayOutput() ErrorDefinitionResponseArrayOutput {
	return o
}

func (o ErrorDefinitionResponseArrayOutput) ToErrorDefinitionResponseArrayOutputWithContext(ctx context.Context) ErrorDefinitionResponseArrayOutput {
	return o
}

func (o ErrorDefinitionResponseArrayOutput) Index(i pulumi.IntInput) ErrorDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorDefinitionResponse {
		return vs[0].([]ErrorDefinitionResponse)[vs[1].(int)]
	}).(ErrorDefinitionResponseOutput)
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





type RemediationDeploymentResponseInput interface {
	pulumi.Input

	ToRemediationDeploymentResponseOutput() RemediationDeploymentResponseOutput
	ToRemediationDeploymentResponseOutputWithContext(context.Context) RemediationDeploymentResponseOutput
}

type RemediationDeploymentResponseArgs struct {
	CreatedOn            pulumi.StringInput           `pulumi:"createdOn"`
	DeploymentId         pulumi.StringInput           `pulumi:"deploymentId"`
	Error                ErrorDefinitionResponseInput `pulumi:"error"`
	LastUpdatedOn        pulumi.StringInput           `pulumi:"lastUpdatedOn"`
	RemediatedResourceId pulumi.StringInput           `pulumi:"remediatedResourceId"`
	ResourceLocation     pulumi.StringInput           `pulumi:"resourceLocation"`
	Status               pulumi.StringInput           `pulumi:"status"`
}

func (RemediationDeploymentResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationDeploymentResponse)(nil)).Elem()
}

func (i RemediationDeploymentResponseArgs) ToRemediationDeploymentResponseOutput() RemediationDeploymentResponseOutput {
	return i.ToRemediationDeploymentResponseOutputWithContext(context.Background())
}

func (i RemediationDeploymentResponseArgs) ToRemediationDeploymentResponseOutputWithContext(ctx context.Context) RemediationDeploymentResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationDeploymentResponseOutput)
}





type RemediationDeploymentResponseArrayInput interface {
	pulumi.Input

	ToRemediationDeploymentResponseArrayOutput() RemediationDeploymentResponseArrayOutput
	ToRemediationDeploymentResponseArrayOutputWithContext(context.Context) RemediationDeploymentResponseArrayOutput
}

type RemediationDeploymentResponseArray []RemediationDeploymentResponseInput

func (RemediationDeploymentResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RemediationDeploymentResponse)(nil)).Elem()
}

func (i RemediationDeploymentResponseArray) ToRemediationDeploymentResponseArrayOutput() RemediationDeploymentResponseArrayOutput {
	return i.ToRemediationDeploymentResponseArrayOutputWithContext(context.Background())
}

func (i RemediationDeploymentResponseArray) ToRemediationDeploymentResponseArrayOutputWithContext(ctx context.Context) RemediationDeploymentResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationDeploymentResponseArrayOutput)
}

type RemediationDeploymentResponseOutput struct{ *pulumi.OutputState }

func (RemediationDeploymentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationDeploymentResponse)(nil)).Elem()
}

func (o RemediationDeploymentResponseOutput) ToRemediationDeploymentResponseOutput() RemediationDeploymentResponseOutput {
	return o
}

func (o RemediationDeploymentResponseOutput) ToRemediationDeploymentResponseOutputWithContext(ctx context.Context) RemediationDeploymentResponseOutput {
	return o
}

func (o RemediationDeploymentResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o RemediationDeploymentResponseOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) string { return v.DeploymentId }).(pulumi.StringOutput)
}

func (o RemediationDeploymentResponseOutput) Error() ErrorDefinitionResponseOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) ErrorDefinitionResponse { return v.Error }).(ErrorDefinitionResponseOutput)
}

func (o RemediationDeploymentResponseOutput) LastUpdatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) string { return v.LastUpdatedOn }).(pulumi.StringOutput)
}

func (o RemediationDeploymentResponseOutput) RemediatedResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) string { return v.RemediatedResourceId }).(pulumi.StringOutput)
}

func (o RemediationDeploymentResponseOutput) ResourceLocation() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) string { return v.ResourceLocation }).(pulumi.StringOutput)
}

func (o RemediationDeploymentResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) string { return v.Status }).(pulumi.StringOutput)
}

type RemediationDeploymentResponseArrayOutput struct{ *pulumi.OutputState }

func (RemediationDeploymentResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RemediationDeploymentResponse)(nil)).Elem()
}

func (o RemediationDeploymentResponseArrayOutput) ToRemediationDeploymentResponseArrayOutput() RemediationDeploymentResponseArrayOutput {
	return o
}

func (o RemediationDeploymentResponseArrayOutput) ToRemediationDeploymentResponseArrayOutputWithContext(ctx context.Context) RemediationDeploymentResponseArrayOutput {
	return o
}

func (o RemediationDeploymentResponseArrayOutput) Index(i pulumi.IntInput) RemediationDeploymentResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RemediationDeploymentResponse {
		return vs[0].([]RemediationDeploymentResponse)[vs[1].(int)]
	}).(RemediationDeploymentResponseOutput)
}

type RemediationDeploymentSummaryResponse struct {
	FailedDeployments     int `pulumi:"failedDeployments"`
	SuccessfulDeployments int `pulumi:"successfulDeployments"`
	TotalDeployments      int `pulumi:"totalDeployments"`
}





type RemediationDeploymentSummaryResponseInput interface {
	pulumi.Input

	ToRemediationDeploymentSummaryResponseOutput() RemediationDeploymentSummaryResponseOutput
	ToRemediationDeploymentSummaryResponseOutputWithContext(context.Context) RemediationDeploymentSummaryResponseOutput
}

type RemediationDeploymentSummaryResponseArgs struct {
	FailedDeployments     pulumi.IntInput `pulumi:"failedDeployments"`
	SuccessfulDeployments pulumi.IntInput `pulumi:"successfulDeployments"`
	TotalDeployments      pulumi.IntInput `pulumi:"totalDeployments"`
}

func (RemediationDeploymentSummaryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationDeploymentSummaryResponse)(nil)).Elem()
}

func (i RemediationDeploymentSummaryResponseArgs) ToRemediationDeploymentSummaryResponseOutput() RemediationDeploymentSummaryResponseOutput {
	return i.ToRemediationDeploymentSummaryResponseOutputWithContext(context.Background())
}

func (i RemediationDeploymentSummaryResponseArgs) ToRemediationDeploymentSummaryResponseOutputWithContext(ctx context.Context) RemediationDeploymentSummaryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationDeploymentSummaryResponseOutput)
}

func (i RemediationDeploymentSummaryResponseArgs) ToRemediationDeploymentSummaryResponsePtrOutput() RemediationDeploymentSummaryResponsePtrOutput {
	return i.ToRemediationDeploymentSummaryResponsePtrOutputWithContext(context.Background())
}

func (i RemediationDeploymentSummaryResponseArgs) ToRemediationDeploymentSummaryResponsePtrOutputWithContext(ctx context.Context) RemediationDeploymentSummaryResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationDeploymentSummaryResponseOutput).ToRemediationDeploymentSummaryResponsePtrOutputWithContext(ctx)
}









type RemediationDeploymentSummaryResponsePtrInput interface {
	pulumi.Input

	ToRemediationDeploymentSummaryResponsePtrOutput() RemediationDeploymentSummaryResponsePtrOutput
	ToRemediationDeploymentSummaryResponsePtrOutputWithContext(context.Context) RemediationDeploymentSummaryResponsePtrOutput
}

type remediationDeploymentSummaryResponsePtrType RemediationDeploymentSummaryResponseArgs

func RemediationDeploymentSummaryResponsePtr(v *RemediationDeploymentSummaryResponseArgs) RemediationDeploymentSummaryResponsePtrInput {
	return (*remediationDeploymentSummaryResponsePtrType)(v)
}

func (*remediationDeploymentSummaryResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationDeploymentSummaryResponse)(nil)).Elem()
}

func (i *remediationDeploymentSummaryResponsePtrType) ToRemediationDeploymentSummaryResponsePtrOutput() RemediationDeploymentSummaryResponsePtrOutput {
	return i.ToRemediationDeploymentSummaryResponsePtrOutputWithContext(context.Background())
}

func (i *remediationDeploymentSummaryResponsePtrType) ToRemediationDeploymentSummaryResponsePtrOutputWithContext(ctx context.Context) RemediationDeploymentSummaryResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationDeploymentSummaryResponsePtrOutput)
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

func (o RemediationDeploymentSummaryResponseOutput) ToRemediationDeploymentSummaryResponsePtrOutput() RemediationDeploymentSummaryResponsePtrOutput {
	return o.ToRemediationDeploymentSummaryResponsePtrOutputWithContext(context.Background())
}

func (o RemediationDeploymentSummaryResponseOutput) ToRemediationDeploymentSummaryResponsePtrOutputWithContext(ctx context.Context) RemediationDeploymentSummaryResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RemediationDeploymentSummaryResponse) *RemediationDeploymentSummaryResponse {
		return &v
	}).(RemediationDeploymentSummaryResponsePtrOutput)
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

type RemediationDeploymentSummaryResponsePtrOutput struct{ *pulumi.OutputState }

func (RemediationDeploymentSummaryResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationDeploymentSummaryResponse)(nil)).Elem()
}

func (o RemediationDeploymentSummaryResponsePtrOutput) ToRemediationDeploymentSummaryResponsePtrOutput() RemediationDeploymentSummaryResponsePtrOutput {
	return o
}

func (o RemediationDeploymentSummaryResponsePtrOutput) ToRemediationDeploymentSummaryResponsePtrOutputWithContext(ctx context.Context) RemediationDeploymentSummaryResponsePtrOutput {
	return o
}

func (o RemediationDeploymentSummaryResponsePtrOutput) Elem() RemediationDeploymentSummaryResponseOutput {
	return o.ApplyT(func(v *RemediationDeploymentSummaryResponse) RemediationDeploymentSummaryResponse {
		if v != nil {
			return *v
		}
		var ret RemediationDeploymentSummaryResponse
		return ret
	}).(RemediationDeploymentSummaryResponseOutput)
}

func (o RemediationDeploymentSummaryResponsePtrOutput) FailedDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RemediationDeploymentSummaryResponse) *int {
		if v == nil {
			return nil
		}
		return &v.FailedDeployments
	}).(pulumi.IntPtrOutput)
}

func (o RemediationDeploymentSummaryResponsePtrOutput) SuccessfulDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RemediationDeploymentSummaryResponse) *int {
		if v == nil {
			return nil
		}
		return &v.SuccessfulDeployments
	}).(pulumi.IntPtrOutput)
}

func (o RemediationDeploymentSummaryResponsePtrOutput) TotalDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RemediationDeploymentSummaryResponse) *int {
		if v == nil {
			return nil
		}
		return &v.TotalDeployments
	}).(pulumi.IntPtrOutput)
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





type RemediationFiltersResponseInput interface {
	pulumi.Input

	ToRemediationFiltersResponseOutput() RemediationFiltersResponseOutput
	ToRemediationFiltersResponseOutputWithContext(context.Context) RemediationFiltersResponseOutput
}

type RemediationFiltersResponseArgs struct {
	Locations pulumi.StringArrayInput `pulumi:"locations"`
}

func (RemediationFiltersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationFiltersResponse)(nil)).Elem()
}

func (i RemediationFiltersResponseArgs) ToRemediationFiltersResponseOutput() RemediationFiltersResponseOutput {
	return i.ToRemediationFiltersResponseOutputWithContext(context.Background())
}

func (i RemediationFiltersResponseArgs) ToRemediationFiltersResponseOutputWithContext(ctx context.Context) RemediationFiltersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationFiltersResponseOutput)
}

func (i RemediationFiltersResponseArgs) ToRemediationFiltersResponsePtrOutput() RemediationFiltersResponsePtrOutput {
	return i.ToRemediationFiltersResponsePtrOutputWithContext(context.Background())
}

func (i RemediationFiltersResponseArgs) ToRemediationFiltersResponsePtrOutputWithContext(ctx context.Context) RemediationFiltersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationFiltersResponseOutput).ToRemediationFiltersResponsePtrOutputWithContext(ctx)
}









type RemediationFiltersResponsePtrInput interface {
	pulumi.Input

	ToRemediationFiltersResponsePtrOutput() RemediationFiltersResponsePtrOutput
	ToRemediationFiltersResponsePtrOutputWithContext(context.Context) RemediationFiltersResponsePtrOutput
}

type remediationFiltersResponsePtrType RemediationFiltersResponseArgs

func RemediationFiltersResponsePtr(v *RemediationFiltersResponseArgs) RemediationFiltersResponsePtrInput {
	return (*remediationFiltersResponsePtrType)(v)
}

func (*remediationFiltersResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationFiltersResponse)(nil)).Elem()
}

func (i *remediationFiltersResponsePtrType) ToRemediationFiltersResponsePtrOutput() RemediationFiltersResponsePtrOutput {
	return i.ToRemediationFiltersResponsePtrOutputWithContext(context.Background())
}

func (i *remediationFiltersResponsePtrType) ToRemediationFiltersResponsePtrOutputWithContext(ctx context.Context) RemediationFiltersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationFiltersResponsePtrOutput)
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

func (o RemediationFiltersResponseOutput) ToRemediationFiltersResponsePtrOutput() RemediationFiltersResponsePtrOutput {
	return o.ToRemediationFiltersResponsePtrOutputWithContext(context.Background())
}

func (o RemediationFiltersResponseOutput) ToRemediationFiltersResponsePtrOutputWithContext(ctx context.Context) RemediationFiltersResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RemediationFiltersResponse) *RemediationFiltersResponse {
		return &v
	}).(RemediationFiltersResponsePtrOutput)
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





type RemediationPropertiesResponseFailureThresholdInput interface {
	pulumi.Input

	ToRemediationPropertiesResponseFailureThresholdOutput() RemediationPropertiesResponseFailureThresholdOutput
	ToRemediationPropertiesResponseFailureThresholdOutputWithContext(context.Context) RemediationPropertiesResponseFailureThresholdOutput
}

type RemediationPropertiesResponseFailureThresholdArgs struct {
	Percentage pulumi.Float64PtrInput `pulumi:"percentage"`
}

func (RemediationPropertiesResponseFailureThresholdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationPropertiesResponseFailureThreshold)(nil)).Elem()
}

func (i RemediationPropertiesResponseFailureThresholdArgs) ToRemediationPropertiesResponseFailureThresholdOutput() RemediationPropertiesResponseFailureThresholdOutput {
	return i.ToRemediationPropertiesResponseFailureThresholdOutputWithContext(context.Background())
}

func (i RemediationPropertiesResponseFailureThresholdArgs) ToRemediationPropertiesResponseFailureThresholdOutputWithContext(ctx context.Context) RemediationPropertiesResponseFailureThresholdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationPropertiesResponseFailureThresholdOutput)
}

func (i RemediationPropertiesResponseFailureThresholdArgs) ToRemediationPropertiesResponseFailureThresholdPtrOutput() RemediationPropertiesResponseFailureThresholdPtrOutput {
	return i.ToRemediationPropertiesResponseFailureThresholdPtrOutputWithContext(context.Background())
}

func (i RemediationPropertiesResponseFailureThresholdArgs) ToRemediationPropertiesResponseFailureThresholdPtrOutputWithContext(ctx context.Context) RemediationPropertiesResponseFailureThresholdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationPropertiesResponseFailureThresholdOutput).ToRemediationPropertiesResponseFailureThresholdPtrOutputWithContext(ctx)
}









type RemediationPropertiesResponseFailureThresholdPtrInput interface {
	pulumi.Input

	ToRemediationPropertiesResponseFailureThresholdPtrOutput() RemediationPropertiesResponseFailureThresholdPtrOutput
	ToRemediationPropertiesResponseFailureThresholdPtrOutputWithContext(context.Context) RemediationPropertiesResponseFailureThresholdPtrOutput
}

type remediationPropertiesResponseFailureThresholdPtrType RemediationPropertiesResponseFailureThresholdArgs

func RemediationPropertiesResponseFailureThresholdPtr(v *RemediationPropertiesResponseFailureThresholdArgs) RemediationPropertiesResponseFailureThresholdPtrInput {
	return (*remediationPropertiesResponseFailureThresholdPtrType)(v)
}

func (*remediationPropertiesResponseFailureThresholdPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationPropertiesResponseFailureThreshold)(nil)).Elem()
}

func (i *remediationPropertiesResponseFailureThresholdPtrType) ToRemediationPropertiesResponseFailureThresholdPtrOutput() RemediationPropertiesResponseFailureThresholdPtrOutput {
	return i.ToRemediationPropertiesResponseFailureThresholdPtrOutputWithContext(context.Background())
}

func (i *remediationPropertiesResponseFailureThresholdPtrType) ToRemediationPropertiesResponseFailureThresholdPtrOutputWithContext(ctx context.Context) RemediationPropertiesResponseFailureThresholdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationPropertiesResponseFailureThresholdPtrOutput)
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

func (o RemediationPropertiesResponseFailureThresholdOutput) ToRemediationPropertiesResponseFailureThresholdPtrOutput() RemediationPropertiesResponseFailureThresholdPtrOutput {
	return o.ToRemediationPropertiesResponseFailureThresholdPtrOutputWithContext(context.Background())
}

func (o RemediationPropertiesResponseFailureThresholdOutput) ToRemediationPropertiesResponseFailureThresholdPtrOutputWithContext(ctx context.Context) RemediationPropertiesResponseFailureThresholdPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RemediationPropertiesResponseFailureThreshold) *RemediationPropertiesResponseFailureThreshold {
		return &v
	}).(RemediationPropertiesResponseFailureThresholdPtrOutput)
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

type TypedErrorInfoResponse struct {
	Info interface{} `pulumi:"info"`
	Type string      `pulumi:"type"`
}





type TypedErrorInfoResponseInput interface {
	pulumi.Input

	ToTypedErrorInfoResponseOutput() TypedErrorInfoResponseOutput
	ToTypedErrorInfoResponseOutputWithContext(context.Context) TypedErrorInfoResponseOutput
}

type TypedErrorInfoResponseArgs struct {
	Info pulumi.Input       `pulumi:"info"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (TypedErrorInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TypedErrorInfoResponse)(nil)).Elem()
}

func (i TypedErrorInfoResponseArgs) ToTypedErrorInfoResponseOutput() TypedErrorInfoResponseOutput {
	return i.ToTypedErrorInfoResponseOutputWithContext(context.Background())
}

func (i TypedErrorInfoResponseArgs) ToTypedErrorInfoResponseOutputWithContext(ctx context.Context) TypedErrorInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TypedErrorInfoResponseOutput)
}





type TypedErrorInfoResponseArrayInput interface {
	pulumi.Input

	ToTypedErrorInfoResponseArrayOutput() TypedErrorInfoResponseArrayOutput
	ToTypedErrorInfoResponseArrayOutputWithContext(context.Context) TypedErrorInfoResponseArrayOutput
}

type TypedErrorInfoResponseArray []TypedErrorInfoResponseInput

func (TypedErrorInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TypedErrorInfoResponse)(nil)).Elem()
}

func (i TypedErrorInfoResponseArray) ToTypedErrorInfoResponseArrayOutput() TypedErrorInfoResponseArrayOutput {
	return i.ToTypedErrorInfoResponseArrayOutputWithContext(context.Background())
}

func (i TypedErrorInfoResponseArray) ToTypedErrorInfoResponseArrayOutputWithContext(ctx context.Context) TypedErrorInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TypedErrorInfoResponseArrayOutput)
}

type TypedErrorInfoResponseOutput struct{ *pulumi.OutputState }

func (TypedErrorInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TypedErrorInfoResponse)(nil)).Elem()
}

func (o TypedErrorInfoResponseOutput) ToTypedErrorInfoResponseOutput() TypedErrorInfoResponseOutput {
	return o
}

func (o TypedErrorInfoResponseOutput) ToTypedErrorInfoResponseOutputWithContext(ctx context.Context) TypedErrorInfoResponseOutput {
	return o
}

func (o TypedErrorInfoResponseOutput) Info() pulumi.AnyOutput {
	return o.ApplyT(func(v TypedErrorInfoResponse) interface{} { return v.Info }).(pulumi.AnyOutput)
}

func (o TypedErrorInfoResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v TypedErrorInfoResponse) string { return v.Type }).(pulumi.StringOutput)
}

type TypedErrorInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (TypedErrorInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TypedErrorInfoResponse)(nil)).Elem()
}

func (o TypedErrorInfoResponseArrayOutput) ToTypedErrorInfoResponseArrayOutput() TypedErrorInfoResponseArrayOutput {
	return o
}

func (o TypedErrorInfoResponseArrayOutput) ToTypedErrorInfoResponseArrayOutputWithContext(ctx context.Context) TypedErrorInfoResponseArrayOutput {
	return o
}

func (o TypedErrorInfoResponseArrayOutput) Index(i pulumi.IntInput) TypedErrorInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TypedErrorInfoResponse {
		return vs[0].([]TypedErrorInfoResponse)[vs[1].(int)]
	}).(TypedErrorInfoResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ErrorDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ErrorDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(RemediationDeploymentResponseOutput{})
	pulumi.RegisterOutputType(RemediationDeploymentResponseArrayOutput{})
	pulumi.RegisterOutputType(RemediationDeploymentSummaryResponseOutput{})
	pulumi.RegisterOutputType(RemediationDeploymentSummaryResponsePtrOutput{})
	pulumi.RegisterOutputType(RemediationFiltersOutput{})
	pulumi.RegisterOutputType(RemediationFiltersPtrOutput{})
	pulumi.RegisterOutputType(RemediationFiltersResponseOutput{})
	pulumi.RegisterOutputType(RemediationFiltersResponsePtrOutput{})
	pulumi.RegisterOutputType(RemediationPropertiesFailureThresholdOutput{})
	pulumi.RegisterOutputType(RemediationPropertiesFailureThresholdPtrOutput{})
	pulumi.RegisterOutputType(RemediationPropertiesResponseFailureThresholdOutput{})
	pulumi.RegisterOutputType(RemediationPropertiesResponseFailureThresholdPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(TypedErrorInfoResponseOutput{})
	pulumi.RegisterOutputType(TypedErrorInfoResponseArrayOutput{})
}

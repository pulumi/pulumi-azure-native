


package v20180701preview

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

type RemediationDeploymentSummary struct {
	FailedDeployments     *int `pulumi:"failedDeployments"`
	SuccessfulDeployments *int `pulumi:"successfulDeployments"`
	TotalDeployments      *int `pulumi:"totalDeployments"`
}





type RemediationDeploymentSummaryInput interface {
	pulumi.Input

	ToRemediationDeploymentSummaryOutput() RemediationDeploymentSummaryOutput
	ToRemediationDeploymentSummaryOutputWithContext(context.Context) RemediationDeploymentSummaryOutput
}

type RemediationDeploymentSummaryArgs struct {
	FailedDeployments     pulumi.IntPtrInput `pulumi:"failedDeployments"`
	SuccessfulDeployments pulumi.IntPtrInput `pulumi:"successfulDeployments"`
	TotalDeployments      pulumi.IntPtrInput `pulumi:"totalDeployments"`
}

func (RemediationDeploymentSummaryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationDeploymentSummary)(nil)).Elem()
}

func (i RemediationDeploymentSummaryArgs) ToRemediationDeploymentSummaryOutput() RemediationDeploymentSummaryOutput {
	return i.ToRemediationDeploymentSummaryOutputWithContext(context.Background())
}

func (i RemediationDeploymentSummaryArgs) ToRemediationDeploymentSummaryOutputWithContext(ctx context.Context) RemediationDeploymentSummaryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationDeploymentSummaryOutput)
}

func (i RemediationDeploymentSummaryArgs) ToRemediationDeploymentSummaryPtrOutput() RemediationDeploymentSummaryPtrOutput {
	return i.ToRemediationDeploymentSummaryPtrOutputWithContext(context.Background())
}

func (i RemediationDeploymentSummaryArgs) ToRemediationDeploymentSummaryPtrOutputWithContext(ctx context.Context) RemediationDeploymentSummaryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationDeploymentSummaryOutput).ToRemediationDeploymentSummaryPtrOutputWithContext(ctx)
}









type RemediationDeploymentSummaryPtrInput interface {
	pulumi.Input

	ToRemediationDeploymentSummaryPtrOutput() RemediationDeploymentSummaryPtrOutput
	ToRemediationDeploymentSummaryPtrOutputWithContext(context.Context) RemediationDeploymentSummaryPtrOutput
}

type remediationDeploymentSummaryPtrType RemediationDeploymentSummaryArgs

func RemediationDeploymentSummaryPtr(v *RemediationDeploymentSummaryArgs) RemediationDeploymentSummaryPtrInput {
	return (*remediationDeploymentSummaryPtrType)(v)
}

func (*remediationDeploymentSummaryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationDeploymentSummary)(nil)).Elem()
}

func (i *remediationDeploymentSummaryPtrType) ToRemediationDeploymentSummaryPtrOutput() RemediationDeploymentSummaryPtrOutput {
	return i.ToRemediationDeploymentSummaryPtrOutputWithContext(context.Background())
}

func (i *remediationDeploymentSummaryPtrType) ToRemediationDeploymentSummaryPtrOutputWithContext(ctx context.Context) RemediationDeploymentSummaryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationDeploymentSummaryPtrOutput)
}

type RemediationDeploymentSummaryOutput struct{ *pulumi.OutputState }

func (RemediationDeploymentSummaryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationDeploymentSummary)(nil)).Elem()
}

func (o RemediationDeploymentSummaryOutput) ToRemediationDeploymentSummaryOutput() RemediationDeploymentSummaryOutput {
	return o
}

func (o RemediationDeploymentSummaryOutput) ToRemediationDeploymentSummaryOutputWithContext(ctx context.Context) RemediationDeploymentSummaryOutput {
	return o
}

func (o RemediationDeploymentSummaryOutput) ToRemediationDeploymentSummaryPtrOutput() RemediationDeploymentSummaryPtrOutput {
	return o.ToRemediationDeploymentSummaryPtrOutputWithContext(context.Background())
}

func (o RemediationDeploymentSummaryOutput) ToRemediationDeploymentSummaryPtrOutputWithContext(ctx context.Context) RemediationDeploymentSummaryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RemediationDeploymentSummary) *RemediationDeploymentSummary {
		return &v
	}).(RemediationDeploymentSummaryPtrOutput)
}

func (o RemediationDeploymentSummaryOutput) FailedDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RemediationDeploymentSummary) *int { return v.FailedDeployments }).(pulumi.IntPtrOutput)
}

func (o RemediationDeploymentSummaryOutput) SuccessfulDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RemediationDeploymentSummary) *int { return v.SuccessfulDeployments }).(pulumi.IntPtrOutput)
}

func (o RemediationDeploymentSummaryOutput) TotalDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RemediationDeploymentSummary) *int { return v.TotalDeployments }).(pulumi.IntPtrOutput)
}

type RemediationDeploymentSummaryPtrOutput struct{ *pulumi.OutputState }

func (RemediationDeploymentSummaryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationDeploymentSummary)(nil)).Elem()
}

func (o RemediationDeploymentSummaryPtrOutput) ToRemediationDeploymentSummaryPtrOutput() RemediationDeploymentSummaryPtrOutput {
	return o
}

func (o RemediationDeploymentSummaryPtrOutput) ToRemediationDeploymentSummaryPtrOutputWithContext(ctx context.Context) RemediationDeploymentSummaryPtrOutput {
	return o
}

func (o RemediationDeploymentSummaryPtrOutput) Elem() RemediationDeploymentSummaryOutput {
	return o.ApplyT(func(v *RemediationDeploymentSummary) RemediationDeploymentSummary {
		if v != nil {
			return *v
		}
		var ret RemediationDeploymentSummary
		return ret
	}).(RemediationDeploymentSummaryOutput)
}

func (o RemediationDeploymentSummaryPtrOutput) FailedDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RemediationDeploymentSummary) *int {
		if v == nil {
			return nil
		}
		return v.FailedDeployments
	}).(pulumi.IntPtrOutput)
}

func (o RemediationDeploymentSummaryPtrOutput) SuccessfulDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RemediationDeploymentSummary) *int {
		if v == nil {
			return nil
		}
		return v.SuccessfulDeployments
	}).(pulumi.IntPtrOutput)
}

func (o RemediationDeploymentSummaryPtrOutput) TotalDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RemediationDeploymentSummary) *int {
		if v == nil {
			return nil
		}
		return v.TotalDeployments
	}).(pulumi.IntPtrOutput)
}

type RemediationDeploymentSummaryResponse struct {
	FailedDeployments     *int `pulumi:"failedDeployments"`
	SuccessfulDeployments *int `pulumi:"successfulDeployments"`
	TotalDeployments      *int `pulumi:"totalDeployments"`
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

func (o RemediationDeploymentSummaryResponseOutput) FailedDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RemediationDeploymentSummaryResponse) *int { return v.FailedDeployments }).(pulumi.IntPtrOutput)
}

func (o RemediationDeploymentSummaryResponseOutput) SuccessfulDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RemediationDeploymentSummaryResponse) *int { return v.SuccessfulDeployments }).(pulumi.IntPtrOutput)
}

func (o RemediationDeploymentSummaryResponseOutput) TotalDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RemediationDeploymentSummaryResponse) *int { return v.TotalDeployments }).(pulumi.IntPtrOutput)
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
		return v.FailedDeployments
	}).(pulumi.IntPtrOutput)
}

func (o RemediationDeploymentSummaryResponsePtrOutput) SuccessfulDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RemediationDeploymentSummaryResponse) *int {
		if v == nil {
			return nil
		}
		return v.SuccessfulDeployments
	}).(pulumi.IntPtrOutput)
}

func (o RemediationDeploymentSummaryResponsePtrOutput) TotalDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RemediationDeploymentSummaryResponse) *int {
		if v == nil {
			return nil
		}
		return v.TotalDeployments
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

type TypedErrorInfoResponse struct {
	Info interface{} `pulumi:"info"`
	Type string      `pulumi:"type"`
}

func init() {
	pulumi.RegisterOutputType(RemediationDeploymentSummaryOutput{})
	pulumi.RegisterOutputType(RemediationDeploymentSummaryPtrOutput{})
	pulumi.RegisterOutputType(RemediationDeploymentSummaryResponseOutput{})
	pulumi.RegisterOutputType(RemediationDeploymentSummaryResponsePtrOutput{})
	pulumi.RegisterOutputType(RemediationFiltersOutput{})
	pulumi.RegisterOutputType(RemediationFiltersPtrOutput{})
	pulumi.RegisterOutputType(RemediationFiltersResponseOutput{})
	pulumi.RegisterOutputType(RemediationFiltersResponsePtrOutput{})
}




package v20221116preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OverviewStatusResponse struct {
	FailedCount *int `pulumi:"failedCount"`
	ManualCount *int `pulumi:"manualCount"`
	PassedCount *int `pulumi:"passedCount"`
}

type OverviewStatusResponseOutput struct{ *pulumi.OutputState }

func (OverviewStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OverviewStatusResponse)(nil)).Elem()
}

func (o OverviewStatusResponseOutput) ToOverviewStatusResponseOutput() OverviewStatusResponseOutput {
	return o
}

func (o OverviewStatusResponseOutput) ToOverviewStatusResponseOutputWithContext(ctx context.Context) OverviewStatusResponseOutput {
	return o
}

func (o OverviewStatusResponseOutput) FailedCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OverviewStatusResponse) *int { return v.FailedCount }).(pulumi.IntPtrOutput)
}

func (o OverviewStatusResponseOutput) ManualCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OverviewStatusResponse) *int { return v.ManualCount }).(pulumi.IntPtrOutput)
}

func (o OverviewStatusResponseOutput) PassedCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OverviewStatusResponse) *int { return v.PassedCount }).(pulumi.IntPtrOutput)
}

type OverviewStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (OverviewStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OverviewStatusResponse)(nil)).Elem()
}

func (o OverviewStatusResponsePtrOutput) ToOverviewStatusResponsePtrOutput() OverviewStatusResponsePtrOutput {
	return o
}

func (o OverviewStatusResponsePtrOutput) ToOverviewStatusResponsePtrOutputWithContext(ctx context.Context) OverviewStatusResponsePtrOutput {
	return o
}

func (o OverviewStatusResponsePtrOutput) Elem() OverviewStatusResponseOutput {
	return o.ApplyT(func(v *OverviewStatusResponse) OverviewStatusResponse {
		if v != nil {
			return *v
		}
		var ret OverviewStatusResponse
		return ret
	}).(OverviewStatusResponseOutput)
}

func (o OverviewStatusResponsePtrOutput) FailedCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OverviewStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.FailedCount
	}).(pulumi.IntPtrOutput)
}

func (o OverviewStatusResponsePtrOutput) ManualCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OverviewStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.ManualCount
	}).(pulumi.IntPtrOutput)
}

func (o OverviewStatusResponsePtrOutput) PassedCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OverviewStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.PassedCount
	}).(pulumi.IntPtrOutput)
}

type ReportComplianceStatusResponse struct {
	M365 *OverviewStatusResponse `pulumi:"m365"`
}

type ReportComplianceStatusResponseOutput struct{ *pulumi.OutputState }

func (ReportComplianceStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportComplianceStatusResponse)(nil)).Elem()
}

func (o ReportComplianceStatusResponseOutput) ToReportComplianceStatusResponseOutput() ReportComplianceStatusResponseOutput {
	return o
}

func (o ReportComplianceStatusResponseOutput) ToReportComplianceStatusResponseOutputWithContext(ctx context.Context) ReportComplianceStatusResponseOutput {
	return o
}

func (o ReportComplianceStatusResponseOutput) M365() OverviewStatusResponsePtrOutput {
	return o.ApplyT(func(v ReportComplianceStatusResponse) *OverviewStatusResponse { return v.M365 }).(OverviewStatusResponsePtrOutput)
}

type ReportProperties struct {
	OfferGuid   *string            `pulumi:"offerGuid"`
	Resources   []ResourceMetadata `pulumi:"resources"`
	TimeZone    string             `pulumi:"timeZone"`
	TriggerTime string             `pulumi:"triggerTime"`
}





type ReportPropertiesInput interface {
	pulumi.Input

	ToReportPropertiesOutput() ReportPropertiesOutput
	ToReportPropertiesOutputWithContext(context.Context) ReportPropertiesOutput
}

type ReportPropertiesArgs struct {
	OfferGuid   pulumi.StringPtrInput      `pulumi:"offerGuid"`
	Resources   ResourceMetadataArrayInput `pulumi:"resources"`
	TimeZone    pulumi.StringInput         `pulumi:"timeZone"`
	TriggerTime pulumi.StringInput         `pulumi:"triggerTime"`
}

func (ReportPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportProperties)(nil)).Elem()
}

func (i ReportPropertiesArgs) ToReportPropertiesOutput() ReportPropertiesOutput {
	return i.ToReportPropertiesOutputWithContext(context.Background())
}

func (i ReportPropertiesArgs) ToReportPropertiesOutputWithContext(ctx context.Context) ReportPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportPropertiesOutput)
}

type ReportPropertiesOutput struct{ *pulumi.OutputState }

func (ReportPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportProperties)(nil)).Elem()
}

func (o ReportPropertiesOutput) ToReportPropertiesOutput() ReportPropertiesOutput {
	return o
}

func (o ReportPropertiesOutput) ToReportPropertiesOutputWithContext(ctx context.Context) ReportPropertiesOutput {
	return o
}

func (o ReportPropertiesOutput) OfferGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportProperties) *string { return v.OfferGuid }).(pulumi.StringPtrOutput)
}

func (o ReportPropertiesOutput) Resources() ResourceMetadataArrayOutput {
	return o.ApplyT(func(v ReportProperties) []ResourceMetadata { return v.Resources }).(ResourceMetadataArrayOutput)
}

func (o ReportPropertiesOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v ReportProperties) string { return v.TimeZone }).(pulumi.StringOutput)
}

func (o ReportPropertiesOutput) TriggerTime() pulumi.StringOutput {
	return o.ApplyT(func(v ReportProperties) string { return v.TriggerTime }).(pulumi.StringOutput)
}

type ReportPropertiesResponse struct {
	ComplianceStatus  ReportComplianceStatusResponse `pulumi:"complianceStatus"`
	Id                string                         `pulumi:"id"`
	LastTriggerTime   string                         `pulumi:"lastTriggerTime"`
	NextTriggerTime   string                         `pulumi:"nextTriggerTime"`
	OfferGuid         *string                        `pulumi:"offerGuid"`
	ProvisioningState string                         `pulumi:"provisioningState"`
	ReportName        string                         `pulumi:"reportName"`
	Resources         []ResourceMetadataResponse     `pulumi:"resources"`
	Status            string                         `pulumi:"status"`
	Subscriptions     []string                       `pulumi:"subscriptions"`
	TenantId          string                         `pulumi:"tenantId"`
	TimeZone          string                         `pulumi:"timeZone"`
	TriggerTime       string                         `pulumi:"triggerTime"`
}

type ReportPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ReportPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportPropertiesResponse)(nil)).Elem()
}

func (o ReportPropertiesResponseOutput) ToReportPropertiesResponseOutput() ReportPropertiesResponseOutput {
	return o
}

func (o ReportPropertiesResponseOutput) ToReportPropertiesResponseOutputWithContext(ctx context.Context) ReportPropertiesResponseOutput {
	return o
}

func (o ReportPropertiesResponseOutput) ComplianceStatus() ReportComplianceStatusResponseOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) ReportComplianceStatusResponse { return v.ComplianceStatus }).(ReportComplianceStatusResponseOutput)
}

func (o ReportPropertiesResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ReportPropertiesResponseOutput) LastTriggerTime() pulumi.StringOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) string { return v.LastTriggerTime }).(pulumi.StringOutput)
}

func (o ReportPropertiesResponseOutput) NextTriggerTime() pulumi.StringOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) string { return v.NextTriggerTime }).(pulumi.StringOutput)
}

func (o ReportPropertiesResponseOutput) OfferGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) *string { return v.OfferGuid }).(pulumi.StringPtrOutput)
}

func (o ReportPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ReportPropertiesResponseOutput) ReportName() pulumi.StringOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) string { return v.ReportName }).(pulumi.StringOutput)
}

func (o ReportPropertiesResponseOutput) Resources() ResourceMetadataResponseArrayOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) []ResourceMetadataResponse { return v.Resources }).(ResourceMetadataResponseArrayOutput)
}

func (o ReportPropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o ReportPropertiesResponseOutput) Subscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) []string { return v.Subscriptions }).(pulumi.StringArrayOutput)
}

func (o ReportPropertiesResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ReportPropertiesResponseOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) string { return v.TimeZone }).(pulumi.StringOutput)
}

func (o ReportPropertiesResponseOutput) TriggerTime() pulumi.StringOutput {
	return o.ApplyT(func(v ReportPropertiesResponse) string { return v.TriggerTime }).(pulumi.StringOutput)
}

type ResourceMetadata struct {
	ResourceId   string            `pulumi:"resourceId"`
	ResourceKind *string           `pulumi:"resourceKind"`
	ResourceName *string           `pulumi:"resourceName"`
	ResourceType *string           `pulumi:"resourceType"`
	Tags         map[string]string `pulumi:"tags"`
}





type ResourceMetadataInput interface {
	pulumi.Input

	ToResourceMetadataOutput() ResourceMetadataOutput
	ToResourceMetadataOutputWithContext(context.Context) ResourceMetadataOutput
}

type ResourceMetadataArgs struct {
	ResourceId   pulumi.StringInput    `pulumi:"resourceId"`
	ResourceKind pulumi.StringPtrInput `pulumi:"resourceKind"`
	ResourceName pulumi.StringPtrInput `pulumi:"resourceName"`
	ResourceType pulumi.StringPtrInput `pulumi:"resourceType"`
	Tags         pulumi.StringMapInput `pulumi:"tags"`
}

func (ResourceMetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceMetadata)(nil)).Elem()
}

func (i ResourceMetadataArgs) ToResourceMetadataOutput() ResourceMetadataOutput {
	return i.ToResourceMetadataOutputWithContext(context.Background())
}

func (i ResourceMetadataArgs) ToResourceMetadataOutputWithContext(ctx context.Context) ResourceMetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceMetadataOutput)
}





type ResourceMetadataArrayInput interface {
	pulumi.Input

	ToResourceMetadataArrayOutput() ResourceMetadataArrayOutput
	ToResourceMetadataArrayOutputWithContext(context.Context) ResourceMetadataArrayOutput
}

type ResourceMetadataArray []ResourceMetadataInput

func (ResourceMetadataArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceMetadata)(nil)).Elem()
}

func (i ResourceMetadataArray) ToResourceMetadataArrayOutput() ResourceMetadataArrayOutput {
	return i.ToResourceMetadataArrayOutputWithContext(context.Background())
}

func (i ResourceMetadataArray) ToResourceMetadataArrayOutputWithContext(ctx context.Context) ResourceMetadataArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceMetadataArrayOutput)
}

type ResourceMetadataOutput struct{ *pulumi.OutputState }

func (ResourceMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceMetadata)(nil)).Elem()
}

func (o ResourceMetadataOutput) ToResourceMetadataOutput() ResourceMetadataOutput {
	return o
}

func (o ResourceMetadataOutput) ToResourceMetadataOutputWithContext(ctx context.Context) ResourceMetadataOutput {
	return o
}

func (o ResourceMetadataOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceMetadata) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o ResourceMetadataOutput) ResourceKind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceMetadata) *string { return v.ResourceKind }).(pulumi.StringPtrOutput)
}

func (o ResourceMetadataOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceMetadata) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o ResourceMetadataOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceMetadata) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o ResourceMetadataOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ResourceMetadata) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type ResourceMetadataArrayOutput struct{ *pulumi.OutputState }

func (ResourceMetadataArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceMetadata)(nil)).Elem()
}

func (o ResourceMetadataArrayOutput) ToResourceMetadataArrayOutput() ResourceMetadataArrayOutput {
	return o
}

func (o ResourceMetadataArrayOutput) ToResourceMetadataArrayOutputWithContext(ctx context.Context) ResourceMetadataArrayOutput {
	return o
}

func (o ResourceMetadataArrayOutput) Index(i pulumi.IntInput) ResourceMetadataOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceMetadata {
		return vs[0].([]ResourceMetadata)[vs[1].(int)]
	}).(ResourceMetadataOutput)
}

type ResourceMetadataResponse struct {
	ResourceId   string            `pulumi:"resourceId"`
	ResourceKind *string           `pulumi:"resourceKind"`
	ResourceName *string           `pulumi:"resourceName"`
	ResourceType *string           `pulumi:"resourceType"`
	Tags         map[string]string `pulumi:"tags"`
}

type ResourceMetadataResponseOutput struct{ *pulumi.OutputState }

func (ResourceMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceMetadataResponse)(nil)).Elem()
}

func (o ResourceMetadataResponseOutput) ToResourceMetadataResponseOutput() ResourceMetadataResponseOutput {
	return o
}

func (o ResourceMetadataResponseOutput) ToResourceMetadataResponseOutputWithContext(ctx context.Context) ResourceMetadataResponseOutput {
	return o
}

func (o ResourceMetadataResponseOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceMetadataResponse) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o ResourceMetadataResponseOutput) ResourceKind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceMetadataResponse) *string { return v.ResourceKind }).(pulumi.StringPtrOutput)
}

func (o ResourceMetadataResponseOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceMetadataResponse) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o ResourceMetadataResponseOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceMetadataResponse) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o ResourceMetadataResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ResourceMetadataResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type ResourceMetadataResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceMetadataResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceMetadataResponse)(nil)).Elem()
}

func (o ResourceMetadataResponseArrayOutput) ToResourceMetadataResponseArrayOutput() ResourceMetadataResponseArrayOutput {
	return o
}

func (o ResourceMetadataResponseArrayOutput) ToResourceMetadataResponseArrayOutputWithContext(ctx context.Context) ResourceMetadataResponseArrayOutput {
	return o
}

func (o ResourceMetadataResponseArrayOutput) Index(i pulumi.IntInput) ResourceMetadataResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceMetadataResponse {
		return vs[0].([]ResourceMetadataResponse)[vs[1].(int)]
	}).(ResourceMetadataResponseOutput)
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
	pulumi.RegisterOutputType(OverviewStatusResponseOutput{})
	pulumi.RegisterOutputType(OverviewStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportComplianceStatusResponseOutput{})
	pulumi.RegisterOutputType(ReportPropertiesOutput{})
	pulumi.RegisterOutputType(ReportPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ResourceMetadataOutput{})
	pulumi.RegisterOutputType(ResourceMetadataArrayOutput{})
	pulumi.RegisterOutputType(ResourceMetadataResponseOutput{})
	pulumi.RegisterOutputType(ResourceMetadataResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}

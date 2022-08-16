


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssessmentLinksResponse struct {
	AzurePortalUri string `pulumi:"azurePortalUri"`
}

type AssessmentLinksResponseOutput struct{ *pulumi.OutputState }

func (AssessmentLinksResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentLinksResponse)(nil)).Elem()
}

func (o AssessmentLinksResponseOutput) ToAssessmentLinksResponseOutput() AssessmentLinksResponseOutput {
	return o
}

func (o AssessmentLinksResponseOutput) ToAssessmentLinksResponseOutputWithContext(ctx context.Context) AssessmentLinksResponseOutput {
	return o
}

func (o AssessmentLinksResponseOutput) AzurePortalUri() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentLinksResponse) string { return v.AzurePortalUri }).(pulumi.StringOutput)
}

type AssessmentStatus struct {
	Cause       *string `pulumi:"cause"`
	Code        string  `pulumi:"code"`
	Description *string `pulumi:"description"`
}





type AssessmentStatusInput interface {
	pulumi.Input

	ToAssessmentStatusOutput() AssessmentStatusOutput
	ToAssessmentStatusOutputWithContext(context.Context) AssessmentStatusOutput
}

type AssessmentStatusArgs struct {
	Cause       pulumi.StringPtrInput `pulumi:"cause"`
	Code        pulumi.StringInput    `pulumi:"code"`
	Description pulumi.StringPtrInput `pulumi:"description"`
}

func (AssessmentStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentStatus)(nil)).Elem()
}

func (i AssessmentStatusArgs) ToAssessmentStatusOutput() AssessmentStatusOutput {
	return i.ToAssessmentStatusOutputWithContext(context.Background())
}

func (i AssessmentStatusArgs) ToAssessmentStatusOutputWithContext(ctx context.Context) AssessmentStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentStatusOutput)
}

type AssessmentStatusOutput struct{ *pulumi.OutputState }

func (AssessmentStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentStatus)(nil)).Elem()
}

func (o AssessmentStatusOutput) ToAssessmentStatusOutput() AssessmentStatusOutput {
	return o
}

func (o AssessmentStatusOutput) ToAssessmentStatusOutputWithContext(ctx context.Context) AssessmentStatusOutput {
	return o
}

func (o AssessmentStatusOutput) Cause() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssessmentStatus) *string { return v.Cause }).(pulumi.StringPtrOutput)
}

func (o AssessmentStatusOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentStatus) string { return v.Code }).(pulumi.StringOutput)
}

func (o AssessmentStatusOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssessmentStatus) *string { return v.Description }).(pulumi.StringPtrOutput)
}

type AssessmentStatusResponseResponse struct {
	Cause               *string `pulumi:"cause"`
	Code                string  `pulumi:"code"`
	Description         *string `pulumi:"description"`
	FirstEvaluationDate string  `pulumi:"firstEvaluationDate"`
	StatusChangeDate    string  `pulumi:"statusChangeDate"`
}

type AssessmentStatusResponseResponseOutput struct{ *pulumi.OutputState }

func (AssessmentStatusResponseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentStatusResponseResponse)(nil)).Elem()
}

func (o AssessmentStatusResponseResponseOutput) ToAssessmentStatusResponseResponseOutput() AssessmentStatusResponseResponseOutput {
	return o
}

func (o AssessmentStatusResponseResponseOutput) ToAssessmentStatusResponseResponseOutputWithContext(ctx context.Context) AssessmentStatusResponseResponseOutput {
	return o
}

func (o AssessmentStatusResponseResponseOutput) Cause() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssessmentStatusResponseResponse) *string { return v.Cause }).(pulumi.StringPtrOutput)
}

func (o AssessmentStatusResponseResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentStatusResponseResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o AssessmentStatusResponseResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssessmentStatusResponseResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AssessmentStatusResponseResponseOutput) FirstEvaluationDate() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentStatusResponseResponse) string { return v.FirstEvaluationDate }).(pulumi.StringOutput)
}

func (o AssessmentStatusResponseResponseOutput) StatusChangeDate() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentStatusResponseResponse) string { return v.StatusChangeDate }).(pulumi.StringOutput)
}

type AzureResourceDetails struct {
	Source string `pulumi:"source"`
}

type AzureResourceDetailsResponse struct {
	Id     string `pulumi:"id"`
	Source string `pulumi:"source"`
}

type OnPremiseResourceDetails struct {
	MachineName      string `pulumi:"machineName"`
	Source           string `pulumi:"source"`
	SourceComputerId string `pulumi:"sourceComputerId"`
	Vmuuid           string `pulumi:"vmuuid"`
	WorkspaceId      string `pulumi:"workspaceId"`
}

type OnPremiseResourceDetailsResponse struct {
	MachineName      string `pulumi:"machineName"`
	Source           string `pulumi:"source"`
	SourceComputerId string `pulumi:"sourceComputerId"`
	Vmuuid           string `pulumi:"vmuuid"`
	WorkspaceId      string `pulumi:"workspaceId"`
}

type OnPremiseSqlResourceDetails struct {
	DatabaseName     string `pulumi:"databaseName"`
	MachineName      string `pulumi:"machineName"`
	ServerName       string `pulumi:"serverName"`
	Source           string `pulumi:"source"`
	SourceComputerId string `pulumi:"sourceComputerId"`
	Vmuuid           string `pulumi:"vmuuid"`
	WorkspaceId      string `pulumi:"workspaceId"`
}

type OnPremiseSqlResourceDetailsResponse struct {
	DatabaseName     string `pulumi:"databaseName"`
	MachineName      string `pulumi:"machineName"`
	ServerName       string `pulumi:"serverName"`
	Source           string `pulumi:"source"`
	SourceComputerId string `pulumi:"sourceComputerId"`
	Vmuuid           string `pulumi:"vmuuid"`
	WorkspaceId      string `pulumi:"workspaceId"`
}

type SecurityAssessmentMetadataPartnerData struct {
	PartnerName string  `pulumi:"partnerName"`
	ProductName *string `pulumi:"productName"`
	Secret      string  `pulumi:"secret"`
}





type SecurityAssessmentMetadataPartnerDataInput interface {
	pulumi.Input

	ToSecurityAssessmentMetadataPartnerDataOutput() SecurityAssessmentMetadataPartnerDataOutput
	ToSecurityAssessmentMetadataPartnerDataOutputWithContext(context.Context) SecurityAssessmentMetadataPartnerDataOutput
}

type SecurityAssessmentMetadataPartnerDataArgs struct {
	PartnerName pulumi.StringInput    `pulumi:"partnerName"`
	ProductName pulumi.StringPtrInput `pulumi:"productName"`
	Secret      pulumi.StringInput    `pulumi:"secret"`
}

func (SecurityAssessmentMetadataPartnerDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataPartnerData)(nil)).Elem()
}

func (i SecurityAssessmentMetadataPartnerDataArgs) ToSecurityAssessmentMetadataPartnerDataOutput() SecurityAssessmentMetadataPartnerDataOutput {
	return i.ToSecurityAssessmentMetadataPartnerDataOutputWithContext(context.Background())
}

func (i SecurityAssessmentMetadataPartnerDataArgs) ToSecurityAssessmentMetadataPartnerDataOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPartnerDataOutput)
}

func (i SecurityAssessmentMetadataPartnerDataArgs) ToSecurityAssessmentMetadataPartnerDataPtrOutput() SecurityAssessmentMetadataPartnerDataPtrOutput {
	return i.ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(context.Background())
}

func (i SecurityAssessmentMetadataPartnerDataArgs) ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPartnerDataOutput).ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(ctx)
}









type SecurityAssessmentMetadataPartnerDataPtrInput interface {
	pulumi.Input

	ToSecurityAssessmentMetadataPartnerDataPtrOutput() SecurityAssessmentMetadataPartnerDataPtrOutput
	ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(context.Context) SecurityAssessmentMetadataPartnerDataPtrOutput
}

type securityAssessmentMetadataPartnerDataPtrType SecurityAssessmentMetadataPartnerDataArgs

func SecurityAssessmentMetadataPartnerDataPtr(v *SecurityAssessmentMetadataPartnerDataArgs) SecurityAssessmentMetadataPartnerDataPtrInput {
	return (*securityAssessmentMetadataPartnerDataPtrType)(v)
}

func (*securityAssessmentMetadataPartnerDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataPartnerData)(nil)).Elem()
}

func (i *securityAssessmentMetadataPartnerDataPtrType) ToSecurityAssessmentMetadataPartnerDataPtrOutput() SecurityAssessmentMetadataPartnerDataPtrOutput {
	return i.ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(context.Background())
}

func (i *securityAssessmentMetadataPartnerDataPtrType) ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPartnerDataPtrOutput)
}

type SecurityAssessmentMetadataPartnerDataOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPartnerDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataPartnerData)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPartnerDataOutput) ToSecurityAssessmentMetadataPartnerDataOutput() SecurityAssessmentMetadataPartnerDataOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataOutput) ToSecurityAssessmentMetadataPartnerDataOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataOutput) ToSecurityAssessmentMetadataPartnerDataPtrOutput() SecurityAssessmentMetadataPartnerDataPtrOutput {
	return o.ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(context.Background())
}

func (o SecurityAssessmentMetadataPartnerDataOutput) ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityAssessmentMetadataPartnerData) *SecurityAssessmentMetadataPartnerData {
		return &v
	}).(SecurityAssessmentMetadataPartnerDataPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataOutput) PartnerName() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPartnerData) string { return v.PartnerName }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPartnerDataOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPartnerData) *string { return v.ProductName }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPartnerData) string { return v.Secret }).(pulumi.StringOutput)
}

type SecurityAssessmentMetadataPartnerDataPtrOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPartnerDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataPartnerData)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPartnerDataPtrOutput) ToSecurityAssessmentMetadataPartnerDataPtrOutput() SecurityAssessmentMetadataPartnerDataPtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataPtrOutput) ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataPtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataPtrOutput) Elem() SecurityAssessmentMetadataPartnerDataOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerData) SecurityAssessmentMetadataPartnerData {
		if v != nil {
			return *v
		}
		var ret SecurityAssessmentMetadataPartnerData
		return ret
	}).(SecurityAssessmentMetadataPartnerDataOutput)
}

func (o SecurityAssessmentMetadataPartnerDataPtrOutput) PartnerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerData) *string {
		if v == nil {
			return nil
		}
		return &v.PartnerName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataPtrOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerData) *string {
		if v == nil {
			return nil
		}
		return v.ProductName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataPtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerData) *string {
		if v == nil {
			return nil
		}
		return &v.Secret
	}).(pulumi.StringPtrOutput)
}

type SecurityAssessmentMetadataPartnerDataResponse struct {
	PartnerName string  `pulumi:"partnerName"`
	ProductName *string `pulumi:"productName"`
	Secret      string  `pulumi:"secret"`
}

type SecurityAssessmentMetadataPartnerDataResponseOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPartnerDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataPartnerDataResponse)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPartnerDataResponseOutput) ToSecurityAssessmentMetadataPartnerDataResponseOutput() SecurityAssessmentMetadataPartnerDataResponseOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataResponseOutput) ToSecurityAssessmentMetadataPartnerDataResponseOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataResponseOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataResponseOutput) PartnerName() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPartnerDataResponse) string { return v.PartnerName }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPartnerDataResponseOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPartnerDataResponse) *string { return v.ProductName }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataResponseOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPartnerDataResponse) string { return v.Secret }).(pulumi.StringOutput)
}

type SecurityAssessmentMetadataPartnerDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPartnerDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataPartnerDataResponse)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPartnerDataResponsePtrOutput) ToSecurityAssessmentMetadataPartnerDataResponsePtrOutput() SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataResponsePtrOutput) ToSecurityAssessmentMetadataPartnerDataResponsePtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataResponsePtrOutput) Elem() SecurityAssessmentMetadataPartnerDataResponseOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerDataResponse) SecurityAssessmentMetadataPartnerDataResponse {
		if v != nil {
			return *v
		}
		var ret SecurityAssessmentMetadataPartnerDataResponse
		return ret
	}).(SecurityAssessmentMetadataPartnerDataResponseOutput)
}

func (o SecurityAssessmentMetadataPartnerDataResponsePtrOutput) PartnerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PartnerName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataResponsePtrOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProductName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataResponsePtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Secret
	}).(pulumi.StringPtrOutput)
}

type SecurityAssessmentMetadataProperties struct {
	AssessmentType         string                                 `pulumi:"assessmentType"`
	Categories             []string                               `pulumi:"categories"`
	Description            *string                                `pulumi:"description"`
	DisplayName            string                                 `pulumi:"displayName"`
	ImplementationEffort   *string                                `pulumi:"implementationEffort"`
	PartnerData            *SecurityAssessmentMetadataPartnerData `pulumi:"partnerData"`
	Preview                *bool                                  `pulumi:"preview"`
	RemediationDescription *string                                `pulumi:"remediationDescription"`
	Severity               string                                 `pulumi:"severity"`
	Threats                []string                               `pulumi:"threats"`
	UserImpact             *string                                `pulumi:"userImpact"`
}





type SecurityAssessmentMetadataPropertiesInput interface {
	pulumi.Input

	ToSecurityAssessmentMetadataPropertiesOutput() SecurityAssessmentMetadataPropertiesOutput
	ToSecurityAssessmentMetadataPropertiesOutputWithContext(context.Context) SecurityAssessmentMetadataPropertiesOutput
}

type SecurityAssessmentMetadataPropertiesArgs struct {
	AssessmentType         pulumi.StringInput                            `pulumi:"assessmentType"`
	Categories             pulumi.StringArrayInput                       `pulumi:"categories"`
	Description            pulumi.StringPtrInput                         `pulumi:"description"`
	DisplayName            pulumi.StringInput                            `pulumi:"displayName"`
	ImplementationEffort   pulumi.StringPtrInput                         `pulumi:"implementationEffort"`
	PartnerData            SecurityAssessmentMetadataPartnerDataPtrInput `pulumi:"partnerData"`
	Preview                pulumi.BoolPtrInput                           `pulumi:"preview"`
	RemediationDescription pulumi.StringPtrInput                         `pulumi:"remediationDescription"`
	Severity               pulumi.StringInput                            `pulumi:"severity"`
	Threats                pulumi.StringArrayInput                       `pulumi:"threats"`
	UserImpact             pulumi.StringPtrInput                         `pulumi:"userImpact"`
}

func (SecurityAssessmentMetadataPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataProperties)(nil)).Elem()
}

func (i SecurityAssessmentMetadataPropertiesArgs) ToSecurityAssessmentMetadataPropertiesOutput() SecurityAssessmentMetadataPropertiesOutput {
	return i.ToSecurityAssessmentMetadataPropertiesOutputWithContext(context.Background())
}

func (i SecurityAssessmentMetadataPropertiesArgs) ToSecurityAssessmentMetadataPropertiesOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPropertiesOutput)
}

func (i SecurityAssessmentMetadataPropertiesArgs) ToSecurityAssessmentMetadataPropertiesPtrOutput() SecurityAssessmentMetadataPropertiesPtrOutput {
	return i.ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(context.Background())
}

func (i SecurityAssessmentMetadataPropertiesArgs) ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPropertiesOutput).ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(ctx)
}









type SecurityAssessmentMetadataPropertiesPtrInput interface {
	pulumi.Input

	ToSecurityAssessmentMetadataPropertiesPtrOutput() SecurityAssessmentMetadataPropertiesPtrOutput
	ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(context.Context) SecurityAssessmentMetadataPropertiesPtrOutput
}

type securityAssessmentMetadataPropertiesPtrType SecurityAssessmentMetadataPropertiesArgs

func SecurityAssessmentMetadataPropertiesPtr(v *SecurityAssessmentMetadataPropertiesArgs) SecurityAssessmentMetadataPropertiesPtrInput {
	return (*securityAssessmentMetadataPropertiesPtrType)(v)
}

func (*securityAssessmentMetadataPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataProperties)(nil)).Elem()
}

func (i *securityAssessmentMetadataPropertiesPtrType) ToSecurityAssessmentMetadataPropertiesPtrOutput() SecurityAssessmentMetadataPropertiesPtrOutput {
	return i.ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(context.Background())
}

func (i *securityAssessmentMetadataPropertiesPtrType) ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPropertiesPtrOutput)
}

type SecurityAssessmentMetadataPropertiesOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataProperties)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPropertiesOutput) ToSecurityAssessmentMetadataPropertiesOutput() SecurityAssessmentMetadataPropertiesOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesOutput) ToSecurityAssessmentMetadataPropertiesOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesOutput) ToSecurityAssessmentMetadataPropertiesPtrOutput() SecurityAssessmentMetadataPropertiesPtrOutput {
	return o.ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(context.Background())
}

func (o SecurityAssessmentMetadataPropertiesOutput) ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityAssessmentMetadataProperties) *SecurityAssessmentMetadataProperties {
		return &v
	}).(SecurityAssessmentMetadataPropertiesPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) AssessmentType() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) string { return v.AssessmentType }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) ImplementationEffort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) *string { return v.ImplementationEffort }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) PartnerData() SecurityAssessmentMetadataPartnerDataPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) *SecurityAssessmentMetadataPartnerData {
		return v.PartnerData
	}).(SecurityAssessmentMetadataPartnerDataPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) *bool { return v.Preview }).(pulumi.BoolPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) RemediationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) *string { return v.RemediationDescription }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) string { return v.Severity }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) Threats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) []string { return v.Threats }).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) UserImpact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) *string { return v.UserImpact }).(pulumi.StringPtrOutput)
}

type SecurityAssessmentMetadataPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataProperties)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) ToSecurityAssessmentMetadataPropertiesPtrOutput() SecurityAssessmentMetadataPropertiesPtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesPtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) Elem() SecurityAssessmentMetadataPropertiesOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) SecurityAssessmentMetadataProperties {
		if v != nil {
			return *v
		}
		var ret SecurityAssessmentMetadataProperties
		return ret
	}).(SecurityAssessmentMetadataPropertiesOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) AssessmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AssessmentType
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) []string {
		if v == nil {
			return nil
		}
		return v.Categories
	}).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) ImplementationEffort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return v.ImplementationEffort
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) PartnerData() SecurityAssessmentMetadataPartnerDataPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *SecurityAssessmentMetadataPartnerData {
		if v == nil {
			return nil
		}
		return v.PartnerData
	}).(SecurityAssessmentMetadataPartnerDataPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Preview
	}).(pulumi.BoolPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) RemediationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return v.RemediationDescription
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) Severity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Severity
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) Threats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) []string {
		if v == nil {
			return nil
		}
		return v.Threats
	}).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) UserImpact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return v.UserImpact
	}).(pulumi.StringPtrOutput)
}

type SecurityAssessmentMetadataPropertiesResponse struct {
	AssessmentType         string                                         `pulumi:"assessmentType"`
	Categories             []string                                       `pulumi:"categories"`
	Description            *string                                        `pulumi:"description"`
	DisplayName            string                                         `pulumi:"displayName"`
	ImplementationEffort   *string                                        `pulumi:"implementationEffort"`
	PartnerData            *SecurityAssessmentMetadataPartnerDataResponse `pulumi:"partnerData"`
	PolicyDefinitionId     string                                         `pulumi:"policyDefinitionId"`
	Preview                *bool                                          `pulumi:"preview"`
	RemediationDescription *string                                        `pulumi:"remediationDescription"`
	Severity               string                                         `pulumi:"severity"`
	Threats                []string                                       `pulumi:"threats"`
	UserImpact             *string                                        `pulumi:"userImpact"`
}

type SecurityAssessmentMetadataPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataPropertiesResponse)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) ToSecurityAssessmentMetadataPropertiesResponseOutput() SecurityAssessmentMetadataPropertiesResponseOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) ToSecurityAssessmentMetadataPropertiesResponseOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesResponseOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) AssessmentType() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) string { return v.AssessmentType }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) ImplementationEffort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) *string { return v.ImplementationEffort }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) PartnerData() SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) *SecurityAssessmentMetadataPartnerDataResponse {
		return v.PartnerData
	}).(SecurityAssessmentMetadataPartnerDataResponsePtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) PolicyDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) string { return v.PolicyDefinitionId }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) *bool { return v.Preview }).(pulumi.BoolPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) RemediationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) *string { return v.RemediationDescription }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) string { return v.Severity }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) Threats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) []string { return v.Threats }).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) UserImpact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) *string { return v.UserImpact }).(pulumi.StringPtrOutput)
}

type SecurityAssessmentMetadataPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataPropertiesResponse)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) ToSecurityAssessmentMetadataPropertiesResponsePtrOutput() SecurityAssessmentMetadataPropertiesResponsePtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) ToSecurityAssessmentMetadataPropertiesResponsePtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesResponsePtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) Elem() SecurityAssessmentMetadataPropertiesResponseOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) SecurityAssessmentMetadataPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret SecurityAssessmentMetadataPropertiesResponse
		return ret
	}).(SecurityAssessmentMetadataPropertiesResponseOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) AssessmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AssessmentType
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.Categories
	}).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) ImplementationEffort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ImplementationEffort
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) PartnerData() SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *SecurityAssessmentMetadataPartnerDataResponse {
		if v == nil {
			return nil
		}
		return v.PartnerData
	}).(SecurityAssessmentMetadataPartnerDataResponsePtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) PolicyDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PolicyDefinitionId
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Preview
	}).(pulumi.BoolPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) RemediationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RemediationDescription
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) Severity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Severity
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) Threats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.Threats
	}).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) UserImpact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserImpact
	}).(pulumi.StringPtrOutput)
}

type SecurityAssessmentMetadataPropertiesResponsePublishDates struct {
	GA     *string `pulumi:"gA"`
	Public string  `pulumi:"public"`
}





type SecurityAssessmentMetadataPropertiesResponsePublishDatesInput interface {
	pulumi.Input

	ToSecurityAssessmentMetadataPropertiesResponsePublishDatesOutput() SecurityAssessmentMetadataPropertiesResponsePublishDatesOutput
	ToSecurityAssessmentMetadataPropertiesResponsePublishDatesOutputWithContext(context.Context) SecurityAssessmentMetadataPropertiesResponsePublishDatesOutput
}

type SecurityAssessmentMetadataPropertiesResponsePublishDatesArgs struct {
	GA     pulumi.StringPtrInput `pulumi:"gA"`
	Public pulumi.StringInput    `pulumi:"public"`
}

func (SecurityAssessmentMetadataPropertiesResponsePublishDatesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataPropertiesResponsePublishDates)(nil)).Elem()
}

func (i SecurityAssessmentMetadataPropertiesResponsePublishDatesArgs) ToSecurityAssessmentMetadataPropertiesResponsePublishDatesOutput() SecurityAssessmentMetadataPropertiesResponsePublishDatesOutput {
	return i.ToSecurityAssessmentMetadataPropertiesResponsePublishDatesOutputWithContext(context.Background())
}

func (i SecurityAssessmentMetadataPropertiesResponsePublishDatesArgs) ToSecurityAssessmentMetadataPropertiesResponsePublishDatesOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesResponsePublishDatesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPropertiesResponsePublishDatesOutput)
}

func (i SecurityAssessmentMetadataPropertiesResponsePublishDatesArgs) ToSecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput() SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput {
	return i.ToSecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutputWithContext(context.Background())
}

func (i SecurityAssessmentMetadataPropertiesResponsePublishDatesArgs) ToSecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPropertiesResponsePublishDatesOutput).ToSecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutputWithContext(ctx)
}









type SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrInput interface {
	pulumi.Input

	ToSecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput() SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput
	ToSecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutputWithContext(context.Context) SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput
}

type securityAssessmentMetadataPropertiesResponsePublishDatesPtrType SecurityAssessmentMetadataPropertiesResponsePublishDatesArgs

func SecurityAssessmentMetadataPropertiesResponsePublishDatesPtr(v *SecurityAssessmentMetadataPropertiesResponsePublishDatesArgs) SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrInput {
	return (*securityAssessmentMetadataPropertiesResponsePublishDatesPtrType)(v)
}

func (*securityAssessmentMetadataPropertiesResponsePublishDatesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataPropertiesResponsePublishDates)(nil)).Elem()
}

func (i *securityAssessmentMetadataPropertiesResponsePublishDatesPtrType) ToSecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput() SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput {
	return i.ToSecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutputWithContext(context.Background())
}

func (i *securityAssessmentMetadataPropertiesResponsePublishDatesPtrType) ToSecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput)
}

type SecurityAssessmentMetadataPropertiesResponsePublishDatesOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPropertiesResponsePublishDatesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataPropertiesResponsePublishDates)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPropertiesResponsePublishDatesOutput) ToSecurityAssessmentMetadataPropertiesResponsePublishDatesOutput() SecurityAssessmentMetadataPropertiesResponsePublishDatesOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesResponsePublishDatesOutput) ToSecurityAssessmentMetadataPropertiesResponsePublishDatesOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesResponsePublishDatesOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesResponsePublishDatesOutput) ToSecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput() SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput {
	return o.ToSecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutputWithContext(context.Background())
}

func (o SecurityAssessmentMetadataPropertiesResponsePublishDatesOutput) ToSecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityAssessmentMetadataPropertiesResponsePublishDates) *SecurityAssessmentMetadataPropertiesResponsePublishDates {
		return &v
	}).(SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePublishDatesOutput) GA() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponsePublishDates) *string { return v.GA }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePublishDatesOutput) Public() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponsePublishDates) string { return v.Public }).(pulumi.StringOutput)
}

type SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataPropertiesResponsePublishDates)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput) ToSecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput() SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput) ToSecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput) Elem() SecurityAssessmentMetadataPropertiesResponsePublishDatesOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponsePublishDates) SecurityAssessmentMetadataPropertiesResponsePublishDates {
		if v != nil {
			return *v
		}
		var ret SecurityAssessmentMetadataPropertiesResponsePublishDates
		return ret
	}).(SecurityAssessmentMetadataPropertiesResponsePublishDatesOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput) GA() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponsePublishDates) *string {
		if v == nil {
			return nil
		}
		return v.GA
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput) Public() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponsePublishDates) *string {
		if v == nil {
			return nil
		}
		return &v.Public
	}).(pulumi.StringPtrOutput)
}

type SecurityAssessmentMetadataPropertiesResponseResponsePublishDates struct {
	GA     *string `pulumi:"gA"`
	Public string  `pulumi:"public"`
}

type SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataPropertiesResponseResponsePublishDates)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesOutput) ToSecurityAssessmentMetadataPropertiesResponseResponsePublishDatesOutput() SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesOutput) ToSecurityAssessmentMetadataPropertiesResponseResponsePublishDatesOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesOutput) GA() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponseResponsePublishDates) *string { return v.GA }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesOutput) Public() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponseResponsePublishDates) string { return v.Public }).(pulumi.StringOutput)
}

type SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesPtrOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataPropertiesResponseResponsePublishDates)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesPtrOutput) ToSecurityAssessmentMetadataPropertiesResponseResponsePublishDatesPtrOutput() SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesPtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesPtrOutput) ToSecurityAssessmentMetadataPropertiesResponseResponsePublishDatesPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesPtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesPtrOutput) Elem() SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponseResponsePublishDates) SecurityAssessmentMetadataPropertiesResponseResponsePublishDates {
		if v != nil {
			return *v
		}
		var ret SecurityAssessmentMetadataPropertiesResponseResponsePublishDates
		return ret
	}).(SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesPtrOutput) GA() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponseResponsePublishDates) *string {
		if v == nil {
			return nil
		}
		return v.GA
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesPtrOutput) Public() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponseResponsePublishDates) *string {
		if v == nil {
			return nil
		}
		return &v.Public
	}).(pulumi.StringPtrOutput)
}

type SecurityAssessmentPartnerData struct {
	PartnerName string `pulumi:"partnerName"`
	Secret      string `pulumi:"secret"`
}





type SecurityAssessmentPartnerDataInput interface {
	pulumi.Input

	ToSecurityAssessmentPartnerDataOutput() SecurityAssessmentPartnerDataOutput
	ToSecurityAssessmentPartnerDataOutputWithContext(context.Context) SecurityAssessmentPartnerDataOutput
}

type SecurityAssessmentPartnerDataArgs struct {
	PartnerName pulumi.StringInput `pulumi:"partnerName"`
	Secret      pulumi.StringInput `pulumi:"secret"`
}

func (SecurityAssessmentPartnerDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentPartnerData)(nil)).Elem()
}

func (i SecurityAssessmentPartnerDataArgs) ToSecurityAssessmentPartnerDataOutput() SecurityAssessmentPartnerDataOutput {
	return i.ToSecurityAssessmentPartnerDataOutputWithContext(context.Background())
}

func (i SecurityAssessmentPartnerDataArgs) ToSecurityAssessmentPartnerDataOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentPartnerDataOutput)
}

func (i SecurityAssessmentPartnerDataArgs) ToSecurityAssessmentPartnerDataPtrOutput() SecurityAssessmentPartnerDataPtrOutput {
	return i.ToSecurityAssessmentPartnerDataPtrOutputWithContext(context.Background())
}

func (i SecurityAssessmentPartnerDataArgs) ToSecurityAssessmentPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentPartnerDataOutput).ToSecurityAssessmentPartnerDataPtrOutputWithContext(ctx)
}









type SecurityAssessmentPartnerDataPtrInput interface {
	pulumi.Input

	ToSecurityAssessmentPartnerDataPtrOutput() SecurityAssessmentPartnerDataPtrOutput
	ToSecurityAssessmentPartnerDataPtrOutputWithContext(context.Context) SecurityAssessmentPartnerDataPtrOutput
}

type securityAssessmentPartnerDataPtrType SecurityAssessmentPartnerDataArgs

func SecurityAssessmentPartnerDataPtr(v *SecurityAssessmentPartnerDataArgs) SecurityAssessmentPartnerDataPtrInput {
	return (*securityAssessmentPartnerDataPtrType)(v)
}

func (*securityAssessmentPartnerDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentPartnerData)(nil)).Elem()
}

func (i *securityAssessmentPartnerDataPtrType) ToSecurityAssessmentPartnerDataPtrOutput() SecurityAssessmentPartnerDataPtrOutput {
	return i.ToSecurityAssessmentPartnerDataPtrOutputWithContext(context.Background())
}

func (i *securityAssessmentPartnerDataPtrType) ToSecurityAssessmentPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentPartnerDataPtrOutput)
}

type SecurityAssessmentPartnerDataOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentPartnerDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentPartnerData)(nil)).Elem()
}

func (o SecurityAssessmentPartnerDataOutput) ToSecurityAssessmentPartnerDataOutput() SecurityAssessmentPartnerDataOutput {
	return o
}

func (o SecurityAssessmentPartnerDataOutput) ToSecurityAssessmentPartnerDataOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataOutput {
	return o
}

func (o SecurityAssessmentPartnerDataOutput) ToSecurityAssessmentPartnerDataPtrOutput() SecurityAssessmentPartnerDataPtrOutput {
	return o.ToSecurityAssessmentPartnerDataPtrOutputWithContext(context.Background())
}

func (o SecurityAssessmentPartnerDataOutput) ToSecurityAssessmentPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityAssessmentPartnerData) *SecurityAssessmentPartnerData {
		return &v
	}).(SecurityAssessmentPartnerDataPtrOutput)
}

func (o SecurityAssessmentPartnerDataOutput) PartnerName() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentPartnerData) string { return v.PartnerName }).(pulumi.StringOutput)
}

func (o SecurityAssessmentPartnerDataOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentPartnerData) string { return v.Secret }).(pulumi.StringOutput)
}

type SecurityAssessmentPartnerDataPtrOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentPartnerDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentPartnerData)(nil)).Elem()
}

func (o SecurityAssessmentPartnerDataPtrOutput) ToSecurityAssessmentPartnerDataPtrOutput() SecurityAssessmentPartnerDataPtrOutput {
	return o
}

func (o SecurityAssessmentPartnerDataPtrOutput) ToSecurityAssessmentPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataPtrOutput {
	return o
}

func (o SecurityAssessmentPartnerDataPtrOutput) Elem() SecurityAssessmentPartnerDataOutput {
	return o.ApplyT(func(v *SecurityAssessmentPartnerData) SecurityAssessmentPartnerData {
		if v != nil {
			return *v
		}
		var ret SecurityAssessmentPartnerData
		return ret
	}).(SecurityAssessmentPartnerDataOutput)
}

func (o SecurityAssessmentPartnerDataPtrOutput) PartnerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentPartnerData) *string {
		if v == nil {
			return nil
		}
		return &v.PartnerName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentPartnerDataPtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentPartnerData) *string {
		if v == nil {
			return nil
		}
		return &v.Secret
	}).(pulumi.StringPtrOutput)
}

type SecurityAssessmentPartnerDataResponse struct {
	PartnerName string `pulumi:"partnerName"`
	Secret      string `pulumi:"secret"`
}

type SecurityAssessmentPartnerDataResponseOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentPartnerDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentPartnerDataResponse)(nil)).Elem()
}

func (o SecurityAssessmentPartnerDataResponseOutput) ToSecurityAssessmentPartnerDataResponseOutput() SecurityAssessmentPartnerDataResponseOutput {
	return o
}

func (o SecurityAssessmentPartnerDataResponseOutput) ToSecurityAssessmentPartnerDataResponseOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataResponseOutput {
	return o
}

func (o SecurityAssessmentPartnerDataResponseOutput) PartnerName() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentPartnerDataResponse) string { return v.PartnerName }).(pulumi.StringOutput)
}

func (o SecurityAssessmentPartnerDataResponseOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentPartnerDataResponse) string { return v.Secret }).(pulumi.StringOutput)
}

type SecurityAssessmentPartnerDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentPartnerDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentPartnerDataResponse)(nil)).Elem()
}

func (o SecurityAssessmentPartnerDataResponsePtrOutput) ToSecurityAssessmentPartnerDataResponsePtrOutput() SecurityAssessmentPartnerDataResponsePtrOutput {
	return o
}

func (o SecurityAssessmentPartnerDataResponsePtrOutput) ToSecurityAssessmentPartnerDataResponsePtrOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataResponsePtrOutput {
	return o
}

func (o SecurityAssessmentPartnerDataResponsePtrOutput) Elem() SecurityAssessmentPartnerDataResponseOutput {
	return o.ApplyT(func(v *SecurityAssessmentPartnerDataResponse) SecurityAssessmentPartnerDataResponse {
		if v != nil {
			return *v
		}
		var ret SecurityAssessmentPartnerDataResponse
		return ret
	}).(SecurityAssessmentPartnerDataResponseOutput)
}

func (o SecurityAssessmentPartnerDataResponsePtrOutput) PartnerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentPartnerDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PartnerName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentPartnerDataResponsePtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentPartnerDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Secret
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AssessmentLinksResponseOutput{})
	pulumi.RegisterOutputType(AssessmentStatusOutput{})
	pulumi.RegisterOutputType(AssessmentStatusResponseResponseOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPartnerDataOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPartnerDataPtrOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPartnerDataResponseOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPartnerDataResponsePtrOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPropertiesOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPropertiesResponsePublishDatesOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesPtrOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentPartnerDataOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentPartnerDataPtrOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentPartnerDataResponseOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentPartnerDataResponsePtrOutput{})
}

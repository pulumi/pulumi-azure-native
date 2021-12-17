


package alertsmanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActionGroup struct {
	ActionGroupId string      `pulumi:"actionGroupId"`
	Conditions    *Conditions `pulumi:"conditions"`
	Description   *string     `pulumi:"description"`
	Scope         *Scope      `pulumi:"scope"`
	Status        *string     `pulumi:"status"`
	Type          string      `pulumi:"type"`
}

type ActionGroupResponse struct {
	ActionGroupId  string              `pulumi:"actionGroupId"`
	Conditions     *ConditionsResponse `pulumi:"conditions"`
	CreatedAt      string              `pulumi:"createdAt"`
	CreatedBy      string              `pulumi:"createdBy"`
	Description    *string             `pulumi:"description"`
	LastModifiedAt string              `pulumi:"lastModifiedAt"`
	LastModifiedBy string              `pulumi:"lastModifiedBy"`
	Scope          *ScopeResponse      `pulumi:"scope"`
	Status         *string             `pulumi:"status"`
	Type           string              `pulumi:"type"`
}

type ActionGroupsInformation struct {
	CustomEmailSubject   *string  `pulumi:"customEmailSubject"`
	CustomWebhookPayload *string  `pulumi:"customWebhookPayload"`
	GroupIds             []string `pulumi:"groupIds"`
}





type ActionGroupsInformationInput interface {
	pulumi.Input

	ToActionGroupsInformationOutput() ActionGroupsInformationOutput
	ToActionGroupsInformationOutputWithContext(context.Context) ActionGroupsInformationOutput
}

type ActionGroupsInformationArgs struct {
	CustomEmailSubject   pulumi.StringPtrInput   `pulumi:"customEmailSubject"`
	CustomWebhookPayload pulumi.StringPtrInput   `pulumi:"customWebhookPayload"`
	GroupIds             pulumi.StringArrayInput `pulumi:"groupIds"`
}

func (ActionGroupsInformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroupsInformation)(nil)).Elem()
}

func (i ActionGroupsInformationArgs) ToActionGroupsInformationOutput() ActionGroupsInformationOutput {
	return i.ToActionGroupsInformationOutputWithContext(context.Background())
}

func (i ActionGroupsInformationArgs) ToActionGroupsInformationOutputWithContext(ctx context.Context) ActionGroupsInformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupsInformationOutput)
}

type ActionGroupsInformationOutput struct{ *pulumi.OutputState }

func (ActionGroupsInformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroupsInformation)(nil)).Elem()
}

func (o ActionGroupsInformationOutput) ToActionGroupsInformationOutput() ActionGroupsInformationOutput {
	return o
}

func (o ActionGroupsInformationOutput) ToActionGroupsInformationOutputWithContext(ctx context.Context) ActionGroupsInformationOutput {
	return o
}

func (o ActionGroupsInformationOutput) CustomEmailSubject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionGroupsInformation) *string { return v.CustomEmailSubject }).(pulumi.StringPtrOutput)
}

func (o ActionGroupsInformationOutput) CustomWebhookPayload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionGroupsInformation) *string { return v.CustomWebhookPayload }).(pulumi.StringPtrOutput)
}

func (o ActionGroupsInformationOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActionGroupsInformation) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

type ActionGroupsInformationResponse struct {
	CustomEmailSubject   *string  `pulumi:"customEmailSubject"`
	CustomWebhookPayload *string  `pulumi:"customWebhookPayload"`
	GroupIds             []string `pulumi:"groupIds"`
}

type ActionGroupsInformationResponseOutput struct{ *pulumi.OutputState }

func (ActionGroupsInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroupsInformationResponse)(nil)).Elem()
}

func (o ActionGroupsInformationResponseOutput) ToActionGroupsInformationResponseOutput() ActionGroupsInformationResponseOutput {
	return o
}

func (o ActionGroupsInformationResponseOutput) ToActionGroupsInformationResponseOutputWithContext(ctx context.Context) ActionGroupsInformationResponseOutput {
	return o
}

func (o ActionGroupsInformationResponseOutput) CustomEmailSubject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionGroupsInformationResponse) *string { return v.CustomEmailSubject }).(pulumi.StringPtrOutput)
}

func (o ActionGroupsInformationResponseOutput) CustomWebhookPayload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionGroupsInformationResponse) *string { return v.CustomWebhookPayload }).(pulumi.StringPtrOutput)
}

func (o ActionGroupsInformationResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActionGroupsInformationResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

type Condition struct {
	Operator *string  `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}

type ConditionResponse struct {
	Operator *string  `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}

type Conditions struct {
	AlertContext       *Condition `pulumi:"alertContext"`
	AlertRuleId        *Condition `pulumi:"alertRuleId"`
	Description        *Condition `pulumi:"description"`
	MonitorCondition   *Condition `pulumi:"monitorCondition"`
	MonitorService     *Condition `pulumi:"monitorService"`
	Severity           *Condition `pulumi:"severity"`
	TargetResourceType *Condition `pulumi:"targetResourceType"`
}

type ConditionsResponse struct {
	AlertContext       *ConditionResponse `pulumi:"alertContext"`
	AlertRuleId        *ConditionResponse `pulumi:"alertRuleId"`
	Description        *ConditionResponse `pulumi:"description"`
	MonitorCondition   *ConditionResponse `pulumi:"monitorCondition"`
	MonitorService     *ConditionResponse `pulumi:"monitorService"`
	Severity           *ConditionResponse `pulumi:"severity"`
	TargetResourceType *ConditionResponse `pulumi:"targetResourceType"`
}

type Detector struct {
	Description            *string                `pulumi:"description"`
	Id                     string                 `pulumi:"id"`
	ImagePaths             []string               `pulumi:"imagePaths"`
	Name                   *string                `pulumi:"name"`
	Parameters             map[string]interface{} `pulumi:"parameters"`
	SupportedResourceTypes []string               `pulumi:"supportedResourceTypes"`
}





type DetectorInput interface {
	pulumi.Input

	ToDetectorOutput() DetectorOutput
	ToDetectorOutputWithContext(context.Context) DetectorOutput
}

type DetectorArgs struct {
	Description            pulumi.StringPtrInput   `pulumi:"description"`
	Id                     pulumi.StringInput      `pulumi:"id"`
	ImagePaths             pulumi.StringArrayInput `pulumi:"imagePaths"`
	Name                   pulumi.StringPtrInput   `pulumi:"name"`
	Parameters             pulumi.MapInput         `pulumi:"parameters"`
	SupportedResourceTypes pulumi.StringArrayInput `pulumi:"supportedResourceTypes"`
}

func (DetectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Detector)(nil)).Elem()
}

func (i DetectorArgs) ToDetectorOutput() DetectorOutput {
	return i.ToDetectorOutputWithContext(context.Background())
}

func (i DetectorArgs) ToDetectorOutputWithContext(ctx context.Context) DetectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorOutput)
}

type DetectorOutput struct{ *pulumi.OutputState }

func (DetectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Detector)(nil)).Elem()
}

func (o DetectorOutput) ToDetectorOutput() DetectorOutput {
	return o
}

func (o DetectorOutput) ToDetectorOutputWithContext(ctx context.Context) DetectorOutput {
	return o
}

func (o DetectorOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Detector) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DetectorOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v Detector) string { return v.Id }).(pulumi.StringOutput)
}

func (o DetectorOutput) ImagePaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Detector) []string { return v.ImagePaths }).(pulumi.StringArrayOutput)
}

func (o DetectorOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Detector) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DetectorOutput) Parameters() pulumi.MapOutput {
	return o.ApplyT(func(v Detector) map[string]interface{} { return v.Parameters }).(pulumi.MapOutput)
}

func (o DetectorOutput) SupportedResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Detector) []string { return v.SupportedResourceTypes }).(pulumi.StringArrayOutput)
}

type DetectorResponse struct {
	Description            *string                `pulumi:"description"`
	Id                     string                 `pulumi:"id"`
	ImagePaths             []string               `pulumi:"imagePaths"`
	Name                   *string                `pulumi:"name"`
	Parameters             map[string]interface{} `pulumi:"parameters"`
	SupportedResourceTypes []string               `pulumi:"supportedResourceTypes"`
}

type DetectorResponseOutput struct{ *pulumi.OutputState }

func (DetectorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DetectorResponse)(nil)).Elem()
}

func (o DetectorResponseOutput) ToDetectorResponseOutput() DetectorResponseOutput {
	return o
}

func (o DetectorResponseOutput) ToDetectorResponseOutputWithContext(ctx context.Context) DetectorResponseOutput {
	return o
}

func (o DetectorResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DetectorResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DetectorResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v DetectorResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o DetectorResponseOutput) ImagePaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DetectorResponse) []string { return v.ImagePaths }).(pulumi.StringArrayOutput)
}

func (o DetectorResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DetectorResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DetectorResponseOutput) Parameters() pulumi.MapOutput {
	return o.ApplyT(func(v DetectorResponse) map[string]interface{} { return v.Parameters }).(pulumi.MapOutput)
}

func (o DetectorResponseOutput) SupportedResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DetectorResponse) []string { return v.SupportedResourceTypes }).(pulumi.StringArrayOutput)
}

type Diagnostics struct {
	Conditions  *Conditions `pulumi:"conditions"`
	Description *string     `pulumi:"description"`
	Scope       *Scope      `pulumi:"scope"`
	Status      *string     `pulumi:"status"`
	Type        string      `pulumi:"type"`
}

type DiagnosticsResponse struct {
	Conditions     *ConditionsResponse `pulumi:"conditions"`
	CreatedAt      string              `pulumi:"createdAt"`
	CreatedBy      string              `pulumi:"createdBy"`
	Description    *string             `pulumi:"description"`
	LastModifiedAt string              `pulumi:"lastModifiedAt"`
	LastModifiedBy string              `pulumi:"lastModifiedBy"`
	Scope          *ScopeResponse      `pulumi:"scope"`
	Status         *string             `pulumi:"status"`
	Type           string              `pulumi:"type"`
}

type HealthAlertAction struct {
	ActionGroupId     *string           `pulumi:"actionGroupId"`
	WebHookProperties map[string]string `pulumi:"webHookProperties"`
}





type HealthAlertActionInput interface {
	pulumi.Input

	ToHealthAlertActionOutput() HealthAlertActionOutput
	ToHealthAlertActionOutputWithContext(context.Context) HealthAlertActionOutput
}

type HealthAlertActionArgs struct {
	ActionGroupId     pulumi.StringPtrInput `pulumi:"actionGroupId"`
	WebHookProperties pulumi.StringMapInput `pulumi:"webHookProperties"`
}

func (HealthAlertActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertAction)(nil)).Elem()
}

func (i HealthAlertActionArgs) ToHealthAlertActionOutput() HealthAlertActionOutput {
	return i.ToHealthAlertActionOutputWithContext(context.Background())
}

func (i HealthAlertActionArgs) ToHealthAlertActionOutputWithContext(ctx context.Context) HealthAlertActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertActionOutput)
}





type HealthAlertActionArrayInput interface {
	pulumi.Input

	ToHealthAlertActionArrayOutput() HealthAlertActionArrayOutput
	ToHealthAlertActionArrayOutputWithContext(context.Context) HealthAlertActionArrayOutput
}

type HealthAlertActionArray []HealthAlertActionInput

func (HealthAlertActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthAlertAction)(nil)).Elem()
}

func (i HealthAlertActionArray) ToHealthAlertActionArrayOutput() HealthAlertActionArrayOutput {
	return i.ToHealthAlertActionArrayOutputWithContext(context.Background())
}

func (i HealthAlertActionArray) ToHealthAlertActionArrayOutputWithContext(ctx context.Context) HealthAlertActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertActionArrayOutput)
}

type HealthAlertActionOutput struct{ *pulumi.OutputState }

func (HealthAlertActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertAction)(nil)).Elem()
}

func (o HealthAlertActionOutput) ToHealthAlertActionOutput() HealthAlertActionOutput {
	return o
}

func (o HealthAlertActionOutput) ToHealthAlertActionOutputWithContext(ctx context.Context) HealthAlertActionOutput {
	return o
}

func (o HealthAlertActionOutput) ActionGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthAlertAction) *string { return v.ActionGroupId }).(pulumi.StringPtrOutput)
}

func (o HealthAlertActionOutput) WebHookProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v HealthAlertAction) map[string]string { return v.WebHookProperties }).(pulumi.StringMapOutput)
}

type HealthAlertActionArrayOutput struct{ *pulumi.OutputState }

func (HealthAlertActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthAlertAction)(nil)).Elem()
}

func (o HealthAlertActionArrayOutput) ToHealthAlertActionArrayOutput() HealthAlertActionArrayOutput {
	return o
}

func (o HealthAlertActionArrayOutput) ToHealthAlertActionArrayOutputWithContext(ctx context.Context) HealthAlertActionArrayOutput {
	return o
}

func (o HealthAlertActionArrayOutput) Index(i pulumi.IntInput) HealthAlertActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HealthAlertAction {
		return vs[0].([]HealthAlertAction)[vs[1].(int)]
	}).(HealthAlertActionOutput)
}

type HealthAlertActionResponse struct {
	ActionGroupId     *string           `pulumi:"actionGroupId"`
	WebHookProperties map[string]string `pulumi:"webHookProperties"`
}

type HealthAlertActionResponseOutput struct{ *pulumi.OutputState }

func (HealthAlertActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertActionResponse)(nil)).Elem()
}

func (o HealthAlertActionResponseOutput) ToHealthAlertActionResponseOutput() HealthAlertActionResponseOutput {
	return o
}

func (o HealthAlertActionResponseOutput) ToHealthAlertActionResponseOutputWithContext(ctx context.Context) HealthAlertActionResponseOutput {
	return o
}

func (o HealthAlertActionResponseOutput) ActionGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthAlertActionResponse) *string { return v.ActionGroupId }).(pulumi.StringPtrOutput)
}

func (o HealthAlertActionResponseOutput) WebHookProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v HealthAlertActionResponse) map[string]string { return v.WebHookProperties }).(pulumi.StringMapOutput)
}

type HealthAlertActionResponseArrayOutput struct{ *pulumi.OutputState }

func (HealthAlertActionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthAlertActionResponse)(nil)).Elem()
}

func (o HealthAlertActionResponseArrayOutput) ToHealthAlertActionResponseArrayOutput() HealthAlertActionResponseArrayOutput {
	return o
}

func (o HealthAlertActionResponseArrayOutput) ToHealthAlertActionResponseArrayOutputWithContext(ctx context.Context) HealthAlertActionResponseArrayOutput {
	return o
}

func (o HealthAlertActionResponseArrayOutput) Index(i pulumi.IntInput) HealthAlertActionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HealthAlertActionResponse {
		return vs[0].([]HealthAlertActionResponse)[vs[1].(int)]
	}).(HealthAlertActionResponseOutput)
}

type HealthAlertCriteria struct {
	AllOf []VmGuestHealthAlertCriterion `pulumi:"allOf"`
}





type HealthAlertCriteriaInput interface {
	pulumi.Input

	ToHealthAlertCriteriaOutput() HealthAlertCriteriaOutput
	ToHealthAlertCriteriaOutputWithContext(context.Context) HealthAlertCriteriaOutput
}

type HealthAlertCriteriaArgs struct {
	AllOf VmGuestHealthAlertCriterionArrayInput `pulumi:"allOf"`
}

func (HealthAlertCriteriaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertCriteria)(nil)).Elem()
}

func (i HealthAlertCriteriaArgs) ToHealthAlertCriteriaOutput() HealthAlertCriteriaOutput {
	return i.ToHealthAlertCriteriaOutputWithContext(context.Background())
}

func (i HealthAlertCriteriaArgs) ToHealthAlertCriteriaOutputWithContext(ctx context.Context) HealthAlertCriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertCriteriaOutput)
}

type HealthAlertCriteriaOutput struct{ *pulumi.OutputState }

func (HealthAlertCriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertCriteria)(nil)).Elem()
}

func (o HealthAlertCriteriaOutput) ToHealthAlertCriteriaOutput() HealthAlertCriteriaOutput {
	return o
}

func (o HealthAlertCriteriaOutput) ToHealthAlertCriteriaOutputWithContext(ctx context.Context) HealthAlertCriteriaOutput {
	return o
}

func (o HealthAlertCriteriaOutput) AllOf() VmGuestHealthAlertCriterionArrayOutput {
	return o.ApplyT(func(v HealthAlertCriteria) []VmGuestHealthAlertCriterion { return v.AllOf }).(VmGuestHealthAlertCriterionArrayOutput)
}

type HealthAlertCriteriaResponse struct {
	AllOf []VmGuestHealthAlertCriterionResponse `pulumi:"allOf"`
}

type HealthAlertCriteriaResponseOutput struct{ *pulumi.OutputState }

func (HealthAlertCriteriaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertCriteriaResponse)(nil)).Elem()
}

func (o HealthAlertCriteriaResponseOutput) ToHealthAlertCriteriaResponseOutput() HealthAlertCriteriaResponseOutput {
	return o
}

func (o HealthAlertCriteriaResponseOutput) ToHealthAlertCriteriaResponseOutputWithContext(ctx context.Context) HealthAlertCriteriaResponseOutput {
	return o
}

func (o HealthAlertCriteriaResponseOutput) AllOf() VmGuestHealthAlertCriterionResponseArrayOutput {
	return o.ApplyT(func(v HealthAlertCriteriaResponse) []VmGuestHealthAlertCriterionResponse { return v.AllOf }).(VmGuestHealthAlertCriterionResponseArrayOutput)
}

type HealthState struct {
	HealthStateName string  `pulumi:"healthStateName"`
	Severity        float64 `pulumi:"severity"`
}





type HealthStateInput interface {
	pulumi.Input

	ToHealthStateOutput() HealthStateOutput
	ToHealthStateOutputWithContext(context.Context) HealthStateOutput
}

type HealthStateArgs struct {
	HealthStateName pulumi.StringInput  `pulumi:"healthStateName"`
	Severity        pulumi.Float64Input `pulumi:"severity"`
}

func (HealthStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthState)(nil)).Elem()
}

func (i HealthStateArgs) ToHealthStateOutput() HealthStateOutput {
	return i.ToHealthStateOutputWithContext(context.Background())
}

func (i HealthStateArgs) ToHealthStateOutputWithContext(ctx context.Context) HealthStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthStateOutput)
}





type HealthStateArrayInput interface {
	pulumi.Input

	ToHealthStateArrayOutput() HealthStateArrayOutput
	ToHealthStateArrayOutputWithContext(context.Context) HealthStateArrayOutput
}

type HealthStateArray []HealthStateInput

func (HealthStateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthState)(nil)).Elem()
}

func (i HealthStateArray) ToHealthStateArrayOutput() HealthStateArrayOutput {
	return i.ToHealthStateArrayOutputWithContext(context.Background())
}

func (i HealthStateArray) ToHealthStateArrayOutputWithContext(ctx context.Context) HealthStateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthStateArrayOutput)
}

type HealthStateOutput struct{ *pulumi.OutputState }

func (HealthStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthState)(nil)).Elem()
}

func (o HealthStateOutput) ToHealthStateOutput() HealthStateOutput {
	return o
}

func (o HealthStateOutput) ToHealthStateOutputWithContext(ctx context.Context) HealthStateOutput {
	return o
}

func (o HealthStateOutput) HealthStateName() pulumi.StringOutput {
	return o.ApplyT(func(v HealthState) string { return v.HealthStateName }).(pulumi.StringOutput)
}

func (o HealthStateOutput) Severity() pulumi.Float64Output {
	return o.ApplyT(func(v HealthState) float64 { return v.Severity }).(pulumi.Float64Output)
}

type HealthStateArrayOutput struct{ *pulumi.OutputState }

func (HealthStateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthState)(nil)).Elem()
}

func (o HealthStateArrayOutput) ToHealthStateArrayOutput() HealthStateArrayOutput {
	return o
}

func (o HealthStateArrayOutput) ToHealthStateArrayOutputWithContext(ctx context.Context) HealthStateArrayOutput {
	return o
}

func (o HealthStateArrayOutput) Index(i pulumi.IntInput) HealthStateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HealthState {
		return vs[0].([]HealthState)[vs[1].(int)]
	}).(HealthStateOutput)
}

type HealthStateResponse struct {
	HealthStateName string  `pulumi:"healthStateName"`
	Severity        float64 `pulumi:"severity"`
}

type HealthStateResponseOutput struct{ *pulumi.OutputState }

func (HealthStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthStateResponse)(nil)).Elem()
}

func (o HealthStateResponseOutput) ToHealthStateResponseOutput() HealthStateResponseOutput {
	return o
}

func (o HealthStateResponseOutput) ToHealthStateResponseOutputWithContext(ctx context.Context) HealthStateResponseOutput {
	return o
}

func (o HealthStateResponseOutput) HealthStateName() pulumi.StringOutput {
	return o.ApplyT(func(v HealthStateResponse) string { return v.HealthStateName }).(pulumi.StringOutput)
}

func (o HealthStateResponseOutput) Severity() pulumi.Float64Output {
	return o.ApplyT(func(v HealthStateResponse) float64 { return v.Severity }).(pulumi.Float64Output)
}

type HealthStateResponseArrayOutput struct{ *pulumi.OutputState }

func (HealthStateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthStateResponse)(nil)).Elem()
}

func (o HealthStateResponseArrayOutput) ToHealthStateResponseArrayOutput() HealthStateResponseArrayOutput {
	return o
}

func (o HealthStateResponseArrayOutput) ToHealthStateResponseArrayOutputWithContext(ctx context.Context) HealthStateResponseArrayOutput {
	return o
}

func (o HealthStateResponseArrayOutput) Index(i pulumi.IntInput) HealthStateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HealthStateResponse {
		return vs[0].([]HealthStateResponse)[vs[1].(int)]
	}).(HealthStateResponseOutput)
}

type Scope struct {
	ScopeType *string  `pulumi:"scopeType"`
	Values    []string `pulumi:"values"`
}

type ScopeResponse struct {
	ScopeType *string  `pulumi:"scopeType"`
	Values    []string `pulumi:"values"`
}

type Suppression struct {
	Conditions        *Conditions       `pulumi:"conditions"`
	Description       *string           `pulumi:"description"`
	Scope             *Scope            `pulumi:"scope"`
	Status            *string           `pulumi:"status"`
	SuppressionConfig SuppressionConfig `pulumi:"suppressionConfig"`
	Type              string            `pulumi:"type"`
}

type SuppressionConfig struct {
	RecurrenceType string               `pulumi:"recurrenceType"`
	Schedule       *SuppressionSchedule `pulumi:"schedule"`
}

type SuppressionConfigResponse struct {
	RecurrenceType string                       `pulumi:"recurrenceType"`
	Schedule       *SuppressionScheduleResponse `pulumi:"schedule"`
}

type SuppressionResponse struct {
	Conditions        *ConditionsResponse       `pulumi:"conditions"`
	CreatedAt         string                    `pulumi:"createdAt"`
	CreatedBy         string                    `pulumi:"createdBy"`
	Description       *string                   `pulumi:"description"`
	LastModifiedAt    string                    `pulumi:"lastModifiedAt"`
	LastModifiedBy    string                    `pulumi:"lastModifiedBy"`
	Scope             *ScopeResponse            `pulumi:"scope"`
	Status            *string                   `pulumi:"status"`
	SuppressionConfig SuppressionConfigResponse `pulumi:"suppressionConfig"`
	Type              string                    `pulumi:"type"`
}

type SuppressionSchedule struct {
	EndDate          *string `pulumi:"endDate"`
	EndTime          *string `pulumi:"endTime"`
	RecurrenceValues []int   `pulumi:"recurrenceValues"`
	StartDate        *string `pulumi:"startDate"`
	StartTime        *string `pulumi:"startTime"`
}

type SuppressionScheduleResponse struct {
	EndDate          *string `pulumi:"endDate"`
	EndTime          *string `pulumi:"endTime"`
	RecurrenceValues []int   `pulumi:"recurrenceValues"`
	StartDate        *string `pulumi:"startDate"`
	StartTime        *string `pulumi:"startTime"`
}

type ThrottlingInformation struct {
	Duration *string `pulumi:"duration"`
}





type ThrottlingInformationInput interface {
	pulumi.Input

	ToThrottlingInformationOutput() ThrottlingInformationOutput
	ToThrottlingInformationOutputWithContext(context.Context) ThrottlingInformationOutput
}

type ThrottlingInformationArgs struct {
	Duration pulumi.StringPtrInput `pulumi:"duration"`
}

func (ThrottlingInformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThrottlingInformation)(nil)).Elem()
}

func (i ThrottlingInformationArgs) ToThrottlingInformationOutput() ThrottlingInformationOutput {
	return i.ToThrottlingInformationOutputWithContext(context.Background())
}

func (i ThrottlingInformationArgs) ToThrottlingInformationOutputWithContext(ctx context.Context) ThrottlingInformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThrottlingInformationOutput)
}

func (i ThrottlingInformationArgs) ToThrottlingInformationPtrOutput() ThrottlingInformationPtrOutput {
	return i.ToThrottlingInformationPtrOutputWithContext(context.Background())
}

func (i ThrottlingInformationArgs) ToThrottlingInformationPtrOutputWithContext(ctx context.Context) ThrottlingInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThrottlingInformationOutput).ToThrottlingInformationPtrOutputWithContext(ctx)
}









type ThrottlingInformationPtrInput interface {
	pulumi.Input

	ToThrottlingInformationPtrOutput() ThrottlingInformationPtrOutput
	ToThrottlingInformationPtrOutputWithContext(context.Context) ThrottlingInformationPtrOutput
}

type throttlingInformationPtrType ThrottlingInformationArgs

func ThrottlingInformationPtr(v *ThrottlingInformationArgs) ThrottlingInformationPtrInput {
	return (*throttlingInformationPtrType)(v)
}

func (*throttlingInformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ThrottlingInformation)(nil)).Elem()
}

func (i *throttlingInformationPtrType) ToThrottlingInformationPtrOutput() ThrottlingInformationPtrOutput {
	return i.ToThrottlingInformationPtrOutputWithContext(context.Background())
}

func (i *throttlingInformationPtrType) ToThrottlingInformationPtrOutputWithContext(ctx context.Context) ThrottlingInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThrottlingInformationPtrOutput)
}

type ThrottlingInformationOutput struct{ *pulumi.OutputState }

func (ThrottlingInformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThrottlingInformation)(nil)).Elem()
}

func (o ThrottlingInformationOutput) ToThrottlingInformationOutput() ThrottlingInformationOutput {
	return o
}

func (o ThrottlingInformationOutput) ToThrottlingInformationOutputWithContext(ctx context.Context) ThrottlingInformationOutput {
	return o
}

func (o ThrottlingInformationOutput) ToThrottlingInformationPtrOutput() ThrottlingInformationPtrOutput {
	return o.ToThrottlingInformationPtrOutputWithContext(context.Background())
}

func (o ThrottlingInformationOutput) ToThrottlingInformationPtrOutputWithContext(ctx context.Context) ThrottlingInformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ThrottlingInformation) *ThrottlingInformation {
		return &v
	}).(ThrottlingInformationPtrOutput)
}

func (o ThrottlingInformationOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThrottlingInformation) *string { return v.Duration }).(pulumi.StringPtrOutput)
}

type ThrottlingInformationPtrOutput struct{ *pulumi.OutputState }

func (ThrottlingInformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ThrottlingInformation)(nil)).Elem()
}

func (o ThrottlingInformationPtrOutput) ToThrottlingInformationPtrOutput() ThrottlingInformationPtrOutput {
	return o
}

func (o ThrottlingInformationPtrOutput) ToThrottlingInformationPtrOutputWithContext(ctx context.Context) ThrottlingInformationPtrOutput {
	return o
}

func (o ThrottlingInformationPtrOutput) Elem() ThrottlingInformationOutput {
	return o.ApplyT(func(v *ThrottlingInformation) ThrottlingInformation {
		if v != nil {
			return *v
		}
		var ret ThrottlingInformation
		return ret
	}).(ThrottlingInformationOutput)
}

func (o ThrottlingInformationPtrOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThrottlingInformation) *string {
		if v == nil {
			return nil
		}
		return v.Duration
	}).(pulumi.StringPtrOutput)
}

type ThrottlingInformationResponse struct {
	Duration *string `pulumi:"duration"`
}

type ThrottlingInformationResponseOutput struct{ *pulumi.OutputState }

func (ThrottlingInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThrottlingInformationResponse)(nil)).Elem()
}

func (o ThrottlingInformationResponseOutput) ToThrottlingInformationResponseOutput() ThrottlingInformationResponseOutput {
	return o
}

func (o ThrottlingInformationResponseOutput) ToThrottlingInformationResponseOutputWithContext(ctx context.Context) ThrottlingInformationResponseOutput {
	return o
}

func (o ThrottlingInformationResponseOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThrottlingInformationResponse) *string { return v.Duration }).(pulumi.StringPtrOutput)
}

type ThrottlingInformationResponsePtrOutput struct{ *pulumi.OutputState }

func (ThrottlingInformationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ThrottlingInformationResponse)(nil)).Elem()
}

func (o ThrottlingInformationResponsePtrOutput) ToThrottlingInformationResponsePtrOutput() ThrottlingInformationResponsePtrOutput {
	return o
}

func (o ThrottlingInformationResponsePtrOutput) ToThrottlingInformationResponsePtrOutputWithContext(ctx context.Context) ThrottlingInformationResponsePtrOutput {
	return o
}

func (o ThrottlingInformationResponsePtrOutput) Elem() ThrottlingInformationResponseOutput {
	return o.ApplyT(func(v *ThrottlingInformationResponse) ThrottlingInformationResponse {
		if v != nil {
			return *v
		}
		var ret ThrottlingInformationResponse
		return ret
	}).(ThrottlingInformationResponseOutput)
}

func (o ThrottlingInformationResponsePtrOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThrottlingInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Duration
	}).(pulumi.StringPtrOutput)
}

type VmGuestHealthAlertCriterion struct {
	HealthStates []HealthState `pulumi:"healthStates"`
	MonitorNames []string      `pulumi:"monitorNames"`
	MonitorTypes []string      `pulumi:"monitorTypes"`
	Namespace    string        `pulumi:"namespace"`
}





type VmGuestHealthAlertCriterionInput interface {
	pulumi.Input

	ToVmGuestHealthAlertCriterionOutput() VmGuestHealthAlertCriterionOutput
	ToVmGuestHealthAlertCriterionOutputWithContext(context.Context) VmGuestHealthAlertCriterionOutput
}

type VmGuestHealthAlertCriterionArgs struct {
	HealthStates HealthStateArrayInput   `pulumi:"healthStates"`
	MonitorNames pulumi.StringArrayInput `pulumi:"monitorNames"`
	MonitorTypes pulumi.StringArrayInput `pulumi:"monitorTypes"`
	Namespace    pulumi.StringInput      `pulumi:"namespace"`
}

func (VmGuestHealthAlertCriterionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmGuestHealthAlertCriterion)(nil)).Elem()
}

func (i VmGuestHealthAlertCriterionArgs) ToVmGuestHealthAlertCriterionOutput() VmGuestHealthAlertCriterionOutput {
	return i.ToVmGuestHealthAlertCriterionOutputWithContext(context.Background())
}

func (i VmGuestHealthAlertCriterionArgs) ToVmGuestHealthAlertCriterionOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmGuestHealthAlertCriterionOutput)
}





type VmGuestHealthAlertCriterionArrayInput interface {
	pulumi.Input

	ToVmGuestHealthAlertCriterionArrayOutput() VmGuestHealthAlertCriterionArrayOutput
	ToVmGuestHealthAlertCriterionArrayOutputWithContext(context.Context) VmGuestHealthAlertCriterionArrayOutput
}

type VmGuestHealthAlertCriterionArray []VmGuestHealthAlertCriterionInput

func (VmGuestHealthAlertCriterionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmGuestHealthAlertCriterion)(nil)).Elem()
}

func (i VmGuestHealthAlertCriterionArray) ToVmGuestHealthAlertCriterionArrayOutput() VmGuestHealthAlertCriterionArrayOutput {
	return i.ToVmGuestHealthAlertCriterionArrayOutputWithContext(context.Background())
}

func (i VmGuestHealthAlertCriterionArray) ToVmGuestHealthAlertCriterionArrayOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmGuestHealthAlertCriterionArrayOutput)
}

type VmGuestHealthAlertCriterionOutput struct{ *pulumi.OutputState }

func (VmGuestHealthAlertCriterionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmGuestHealthAlertCriterion)(nil)).Elem()
}

func (o VmGuestHealthAlertCriterionOutput) ToVmGuestHealthAlertCriterionOutput() VmGuestHealthAlertCriterionOutput {
	return o
}

func (o VmGuestHealthAlertCriterionOutput) ToVmGuestHealthAlertCriterionOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionOutput {
	return o
}

func (o VmGuestHealthAlertCriterionOutput) HealthStates() HealthStateArrayOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterion) []HealthState { return v.HealthStates }).(HealthStateArrayOutput)
}

func (o VmGuestHealthAlertCriterionOutput) MonitorNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterion) []string { return v.MonitorNames }).(pulumi.StringArrayOutput)
}

func (o VmGuestHealthAlertCriterionOutput) MonitorTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterion) []string { return v.MonitorTypes }).(pulumi.StringArrayOutput)
}

func (o VmGuestHealthAlertCriterionOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterion) string { return v.Namespace }).(pulumi.StringOutput)
}

type VmGuestHealthAlertCriterionArrayOutput struct{ *pulumi.OutputState }

func (VmGuestHealthAlertCriterionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmGuestHealthAlertCriterion)(nil)).Elem()
}

func (o VmGuestHealthAlertCriterionArrayOutput) ToVmGuestHealthAlertCriterionArrayOutput() VmGuestHealthAlertCriterionArrayOutput {
	return o
}

func (o VmGuestHealthAlertCriterionArrayOutput) ToVmGuestHealthAlertCriterionArrayOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionArrayOutput {
	return o
}

func (o VmGuestHealthAlertCriterionArrayOutput) Index(i pulumi.IntInput) VmGuestHealthAlertCriterionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VmGuestHealthAlertCriterion {
		return vs[0].([]VmGuestHealthAlertCriterion)[vs[1].(int)]
	}).(VmGuestHealthAlertCriterionOutput)
}

type VmGuestHealthAlertCriterionResponse struct {
	HealthStates []HealthStateResponse `pulumi:"healthStates"`
	MonitorNames []string              `pulumi:"monitorNames"`
	MonitorTypes []string              `pulumi:"monitorTypes"`
	Namespace    string                `pulumi:"namespace"`
}

type VmGuestHealthAlertCriterionResponseOutput struct{ *pulumi.OutputState }

func (VmGuestHealthAlertCriterionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmGuestHealthAlertCriterionResponse)(nil)).Elem()
}

func (o VmGuestHealthAlertCriterionResponseOutput) ToVmGuestHealthAlertCriterionResponseOutput() VmGuestHealthAlertCriterionResponseOutput {
	return o
}

func (o VmGuestHealthAlertCriterionResponseOutput) ToVmGuestHealthAlertCriterionResponseOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionResponseOutput {
	return o
}

func (o VmGuestHealthAlertCriterionResponseOutput) HealthStates() HealthStateResponseArrayOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterionResponse) []HealthStateResponse { return v.HealthStates }).(HealthStateResponseArrayOutput)
}

func (o VmGuestHealthAlertCriterionResponseOutput) MonitorNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterionResponse) []string { return v.MonitorNames }).(pulumi.StringArrayOutput)
}

func (o VmGuestHealthAlertCriterionResponseOutput) MonitorTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterionResponse) []string { return v.MonitorTypes }).(pulumi.StringArrayOutput)
}

func (o VmGuestHealthAlertCriterionResponseOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterionResponse) string { return v.Namespace }).(pulumi.StringOutput)
}

type VmGuestHealthAlertCriterionResponseArrayOutput struct{ *pulumi.OutputState }

func (VmGuestHealthAlertCriterionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmGuestHealthAlertCriterionResponse)(nil)).Elem()
}

func (o VmGuestHealthAlertCriterionResponseArrayOutput) ToVmGuestHealthAlertCriterionResponseArrayOutput() VmGuestHealthAlertCriterionResponseArrayOutput {
	return o
}

func (o VmGuestHealthAlertCriterionResponseArrayOutput) ToVmGuestHealthAlertCriterionResponseArrayOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionResponseArrayOutput {
	return o
}

func (o VmGuestHealthAlertCriterionResponseArrayOutput) Index(i pulumi.IntInput) VmGuestHealthAlertCriterionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VmGuestHealthAlertCriterionResponse {
		return vs[0].([]VmGuestHealthAlertCriterionResponse)[vs[1].(int)]
	}).(VmGuestHealthAlertCriterionResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionGroupsInformationOutput{})
	pulumi.RegisterOutputType(ActionGroupsInformationResponseOutput{})
	pulumi.RegisterOutputType(DetectorOutput{})
	pulumi.RegisterOutputType(DetectorResponseOutput{})
	pulumi.RegisterOutputType(HealthAlertActionOutput{})
	pulumi.RegisterOutputType(HealthAlertActionArrayOutput{})
	pulumi.RegisterOutputType(HealthAlertActionResponseOutput{})
	pulumi.RegisterOutputType(HealthAlertActionResponseArrayOutput{})
	pulumi.RegisterOutputType(HealthAlertCriteriaOutput{})
	pulumi.RegisterOutputType(HealthAlertCriteriaResponseOutput{})
	pulumi.RegisterOutputType(HealthStateOutput{})
	pulumi.RegisterOutputType(HealthStateArrayOutput{})
	pulumi.RegisterOutputType(HealthStateResponseOutput{})
	pulumi.RegisterOutputType(HealthStateResponseArrayOutput{})
	pulumi.RegisterOutputType(ThrottlingInformationOutput{})
	pulumi.RegisterOutputType(ThrottlingInformationPtrOutput{})
	pulumi.RegisterOutputType(ThrottlingInformationResponseOutput{})
	pulumi.RegisterOutputType(ThrottlingInformationResponsePtrOutput{})
	pulumi.RegisterOutputType(VmGuestHealthAlertCriterionOutput{})
	pulumi.RegisterOutputType(VmGuestHealthAlertCriterionArrayOutput{})
	pulumi.RegisterOutputType(VmGuestHealthAlertCriterionResponseOutput{})
	pulumi.RegisterOutputType(VmGuestHealthAlertCriterionResponseArrayOutput{})
}

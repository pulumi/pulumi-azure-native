


package v20160301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LocationThresholdRuleCondition struct {
	DataSource          interface{} `pulumi:"dataSource"`
	FailedLocationCount int         `pulumi:"failedLocationCount"`
	OdataType           string      `pulumi:"odataType"`
	WindowSize          *string     `pulumi:"windowSize"`
}

type LocationThresholdRuleConditionResponse struct {
	DataSource          interface{} `pulumi:"dataSource"`
	FailedLocationCount int         `pulumi:"failedLocationCount"`
	OdataType           string      `pulumi:"odataType"`
	WindowSize          *string     `pulumi:"windowSize"`
}

type ManagementEventAggregationCondition struct {
	Operator   *ConditionOperator `pulumi:"operator"`
	Threshold  *float64           `pulumi:"threshold"`
	WindowSize *string            `pulumi:"windowSize"`
}

type ManagementEventAggregationConditionResponse struct {
	Operator   *string  `pulumi:"operator"`
	Threshold  *float64 `pulumi:"threshold"`
	WindowSize *string  `pulumi:"windowSize"`
}

type ManagementEventRuleCondition struct {
	Aggregation *ManagementEventAggregationCondition `pulumi:"aggregation"`
	DataSource  interface{}                          `pulumi:"dataSource"`
	OdataType   string                               `pulumi:"odataType"`
}

type ManagementEventRuleConditionResponse struct {
	Aggregation *ManagementEventAggregationConditionResponse `pulumi:"aggregation"`
	DataSource  interface{}                                  `pulumi:"dataSource"`
	OdataType   string                                       `pulumi:"odataType"`
}

type RetentionPolicy struct {
	Days    int  `pulumi:"days"`
	Enabled bool `pulumi:"enabled"`
}





type RetentionPolicyInput interface {
	pulumi.Input

	ToRetentionPolicyOutput() RetentionPolicyOutput
	ToRetentionPolicyOutputWithContext(context.Context) RetentionPolicyOutput
}

type RetentionPolicyArgs struct {
	Days    pulumi.IntInput  `pulumi:"days"`
	Enabled pulumi.BoolInput `pulumi:"enabled"`
}

func (RetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicy)(nil)).Elem()
}

func (i RetentionPolicyArgs) ToRetentionPolicyOutput() RetentionPolicyOutput {
	return i.ToRetentionPolicyOutputWithContext(context.Background())
}

func (i RetentionPolicyArgs) ToRetentionPolicyOutputWithContext(ctx context.Context) RetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyOutput)
}

type RetentionPolicyOutput struct{ *pulumi.OutputState }

func (RetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicy)(nil)).Elem()
}

func (o RetentionPolicyOutput) ToRetentionPolicyOutput() RetentionPolicyOutput {
	return o
}

func (o RetentionPolicyOutput) ToRetentionPolicyOutputWithContext(ctx context.Context) RetentionPolicyOutput {
	return o
}

func (o RetentionPolicyOutput) Days() pulumi.IntOutput {
	return o.ApplyT(func(v RetentionPolicy) int { return v.Days }).(pulumi.IntOutput)
}

func (o RetentionPolicyOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v RetentionPolicy) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type RetentionPolicyResponse struct {
	Days    int  `pulumi:"days"`
	Enabled bool `pulumi:"enabled"`
}

type RetentionPolicyResponseOutput struct{ *pulumi.OutputState }

func (RetentionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicyResponse)(nil)).Elem()
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponseOutput() RetentionPolicyResponseOutput {
	return o
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponseOutputWithContext(ctx context.Context) RetentionPolicyResponseOutput {
	return o
}

func (o RetentionPolicyResponseOutput) Days() pulumi.IntOutput {
	return o.ApplyT(func(v RetentionPolicyResponse) int { return v.Days }).(pulumi.IntOutput)
}

func (o RetentionPolicyResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v RetentionPolicyResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type RuleEmailAction struct {
	CustomEmails        []string `pulumi:"customEmails"`
	OdataType           string   `pulumi:"odataType"`
	SendToServiceOwners *bool    `pulumi:"sendToServiceOwners"`
}

type RuleEmailActionResponse struct {
	CustomEmails        []string `pulumi:"customEmails"`
	OdataType           string   `pulumi:"odataType"`
	SendToServiceOwners *bool    `pulumi:"sendToServiceOwners"`
}

type RuleManagementEventClaimsDataSource struct {
	EmailAddress *string `pulumi:"emailAddress"`
}

type RuleManagementEventClaimsDataSourceResponse struct {
	EmailAddress *string `pulumi:"emailAddress"`
}

type RuleManagementEventDataSource struct {
	Claims               *RuleManagementEventClaimsDataSource `pulumi:"claims"`
	EventName            *string                              `pulumi:"eventName"`
	EventSource          *string                              `pulumi:"eventSource"`
	LegacyResourceId     *string                              `pulumi:"legacyResourceId"`
	Level                *string                              `pulumi:"level"`
	MetricNamespace      *string                              `pulumi:"metricNamespace"`
	OdataType            string                               `pulumi:"odataType"`
	OperationName        *string                              `pulumi:"operationName"`
	ResourceGroupName    *string                              `pulumi:"resourceGroupName"`
	ResourceLocation     *string                              `pulumi:"resourceLocation"`
	ResourceProviderName *string                              `pulumi:"resourceProviderName"`
	ResourceUri          *string                              `pulumi:"resourceUri"`
	Status               *string                              `pulumi:"status"`
	SubStatus            *string                              `pulumi:"subStatus"`
}

type RuleManagementEventDataSourceResponse struct {
	Claims               *RuleManagementEventClaimsDataSourceResponse `pulumi:"claims"`
	EventName            *string                                      `pulumi:"eventName"`
	EventSource          *string                                      `pulumi:"eventSource"`
	LegacyResourceId     *string                                      `pulumi:"legacyResourceId"`
	Level                *string                                      `pulumi:"level"`
	MetricNamespace      *string                                      `pulumi:"metricNamespace"`
	OdataType            string                                       `pulumi:"odataType"`
	OperationName        *string                                      `pulumi:"operationName"`
	ResourceGroupName    *string                                      `pulumi:"resourceGroupName"`
	ResourceLocation     *string                                      `pulumi:"resourceLocation"`
	ResourceProviderName *string                                      `pulumi:"resourceProviderName"`
	ResourceUri          *string                                      `pulumi:"resourceUri"`
	Status               *string                                      `pulumi:"status"`
	SubStatus            *string                                      `pulumi:"subStatus"`
}

type RuleMetricDataSource struct {
	LegacyResourceId *string `pulumi:"legacyResourceId"`
	MetricName       *string `pulumi:"metricName"`
	MetricNamespace  *string `pulumi:"metricNamespace"`
	OdataType        string  `pulumi:"odataType"`
	ResourceLocation *string `pulumi:"resourceLocation"`
	ResourceUri      *string `pulumi:"resourceUri"`
}

type RuleMetricDataSourceResponse struct {
	LegacyResourceId *string `pulumi:"legacyResourceId"`
	MetricName       *string `pulumi:"metricName"`
	MetricNamespace  *string `pulumi:"metricNamespace"`
	OdataType        string  `pulumi:"odataType"`
	ResourceLocation *string `pulumi:"resourceLocation"`
	ResourceUri      *string `pulumi:"resourceUri"`
}

type RuleWebhookAction struct {
	OdataType  string            `pulumi:"odataType"`
	Properties map[string]string `pulumi:"properties"`
	ServiceUri *string           `pulumi:"serviceUri"`
}

type RuleWebhookActionResponse struct {
	OdataType  string            `pulumi:"odataType"`
	Properties map[string]string `pulumi:"properties"`
	ServiceUri *string           `pulumi:"serviceUri"`
}

type ThresholdRuleCondition struct {
	DataSource      interface{}              `pulumi:"dataSource"`
	OdataType       string                   `pulumi:"odataType"`
	Operator        ConditionOperator        `pulumi:"operator"`
	Threshold       float64                  `pulumi:"threshold"`
	TimeAggregation *TimeAggregationOperator `pulumi:"timeAggregation"`
	WindowSize      *string                  `pulumi:"windowSize"`
}

type ThresholdRuleConditionResponse struct {
	DataSource      interface{} `pulumi:"dataSource"`
	OdataType       string      `pulumi:"odataType"`
	Operator        string      `pulumi:"operator"`
	Threshold       float64     `pulumi:"threshold"`
	TimeAggregation *string     `pulumi:"timeAggregation"`
	WindowSize      *string     `pulumi:"windowSize"`
}

func init() {
	pulumi.RegisterOutputType(RetentionPolicyOutput{})
	pulumi.RegisterOutputType(RetentionPolicyResponseOutput{})
}

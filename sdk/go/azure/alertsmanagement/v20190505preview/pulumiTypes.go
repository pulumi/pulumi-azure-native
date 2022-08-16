


package v20190505preview

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

func init() {
}

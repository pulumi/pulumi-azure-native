


package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssessmentLinksResponse struct {
	AzurePortalUri string `pulumi:"azurePortalUri"`
}





type AssessmentLinksResponseInput interface {
	pulumi.Input

	ToAssessmentLinksResponseOutput() AssessmentLinksResponseOutput
	ToAssessmentLinksResponseOutputWithContext(context.Context) AssessmentLinksResponseOutput
}

type AssessmentLinksResponseArgs struct {
	AzurePortalUri pulumi.StringInput `pulumi:"azurePortalUri"`
}

func (AssessmentLinksResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentLinksResponse)(nil)).Elem()
}

func (i AssessmentLinksResponseArgs) ToAssessmentLinksResponseOutput() AssessmentLinksResponseOutput {
	return i.ToAssessmentLinksResponseOutputWithContext(context.Background())
}

func (i AssessmentLinksResponseArgs) ToAssessmentLinksResponseOutputWithContext(ctx context.Context) AssessmentLinksResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentLinksResponseOutput)
}

func (i AssessmentLinksResponseArgs) ToAssessmentLinksResponsePtrOutput() AssessmentLinksResponsePtrOutput {
	return i.ToAssessmentLinksResponsePtrOutputWithContext(context.Background())
}

func (i AssessmentLinksResponseArgs) ToAssessmentLinksResponsePtrOutputWithContext(ctx context.Context) AssessmentLinksResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentLinksResponseOutput).ToAssessmentLinksResponsePtrOutputWithContext(ctx)
}









type AssessmentLinksResponsePtrInput interface {
	pulumi.Input

	ToAssessmentLinksResponsePtrOutput() AssessmentLinksResponsePtrOutput
	ToAssessmentLinksResponsePtrOutputWithContext(context.Context) AssessmentLinksResponsePtrOutput
}

type assessmentLinksResponsePtrType AssessmentLinksResponseArgs

func AssessmentLinksResponsePtr(v *AssessmentLinksResponseArgs) AssessmentLinksResponsePtrInput {
	return (*assessmentLinksResponsePtrType)(v)
}

func (*assessmentLinksResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentLinksResponse)(nil)).Elem()
}

func (i *assessmentLinksResponsePtrType) ToAssessmentLinksResponsePtrOutput() AssessmentLinksResponsePtrOutput {
	return i.ToAssessmentLinksResponsePtrOutputWithContext(context.Background())
}

func (i *assessmentLinksResponsePtrType) ToAssessmentLinksResponsePtrOutputWithContext(ctx context.Context) AssessmentLinksResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentLinksResponsePtrOutput)
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

func (o AssessmentLinksResponseOutput) ToAssessmentLinksResponsePtrOutput() AssessmentLinksResponsePtrOutput {
	return o.ToAssessmentLinksResponsePtrOutputWithContext(context.Background())
}

func (o AssessmentLinksResponseOutput) ToAssessmentLinksResponsePtrOutputWithContext(ctx context.Context) AssessmentLinksResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssessmentLinksResponse) *AssessmentLinksResponse {
		return &v
	}).(AssessmentLinksResponsePtrOutput)
}

func (o AssessmentLinksResponseOutput) AzurePortalUri() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentLinksResponse) string { return v.AzurePortalUri }).(pulumi.StringOutput)
}

type AssessmentLinksResponsePtrOutput struct{ *pulumi.OutputState }

func (AssessmentLinksResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentLinksResponse)(nil)).Elem()
}

func (o AssessmentLinksResponsePtrOutput) ToAssessmentLinksResponsePtrOutput() AssessmentLinksResponsePtrOutput {
	return o
}

func (o AssessmentLinksResponsePtrOutput) ToAssessmentLinksResponsePtrOutputWithContext(ctx context.Context) AssessmentLinksResponsePtrOutput {
	return o
}

func (o AssessmentLinksResponsePtrOutput) Elem() AssessmentLinksResponseOutput {
	return o.ApplyT(func(v *AssessmentLinksResponse) AssessmentLinksResponse {
		if v != nil {
			return *v
		}
		var ret AssessmentLinksResponse
		return ret
	}).(AssessmentLinksResponseOutput)
}

func (o AssessmentLinksResponsePtrOutput) AzurePortalUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentLinksResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AzurePortalUri
	}).(pulumi.StringPtrOutput)
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

func (i AssessmentStatusArgs) ToAssessmentStatusPtrOutput() AssessmentStatusPtrOutput {
	return i.ToAssessmentStatusPtrOutputWithContext(context.Background())
}

func (i AssessmentStatusArgs) ToAssessmentStatusPtrOutputWithContext(ctx context.Context) AssessmentStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentStatusOutput).ToAssessmentStatusPtrOutputWithContext(ctx)
}









type AssessmentStatusPtrInput interface {
	pulumi.Input

	ToAssessmentStatusPtrOutput() AssessmentStatusPtrOutput
	ToAssessmentStatusPtrOutputWithContext(context.Context) AssessmentStatusPtrOutput
}

type assessmentStatusPtrType AssessmentStatusArgs

func AssessmentStatusPtr(v *AssessmentStatusArgs) AssessmentStatusPtrInput {
	return (*assessmentStatusPtrType)(v)
}

func (*assessmentStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentStatus)(nil)).Elem()
}

func (i *assessmentStatusPtrType) ToAssessmentStatusPtrOutput() AssessmentStatusPtrOutput {
	return i.ToAssessmentStatusPtrOutputWithContext(context.Background())
}

func (i *assessmentStatusPtrType) ToAssessmentStatusPtrOutputWithContext(ctx context.Context) AssessmentStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentStatusPtrOutput)
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

func (o AssessmentStatusOutput) ToAssessmentStatusPtrOutput() AssessmentStatusPtrOutput {
	return o.ToAssessmentStatusPtrOutputWithContext(context.Background())
}

func (o AssessmentStatusOutput) ToAssessmentStatusPtrOutputWithContext(ctx context.Context) AssessmentStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssessmentStatus) *AssessmentStatus {
		return &v
	}).(AssessmentStatusPtrOutput)
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

type AssessmentStatusPtrOutput struct{ *pulumi.OutputState }

func (AssessmentStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentStatus)(nil)).Elem()
}

func (o AssessmentStatusPtrOutput) ToAssessmentStatusPtrOutput() AssessmentStatusPtrOutput {
	return o
}

func (o AssessmentStatusPtrOutput) ToAssessmentStatusPtrOutputWithContext(ctx context.Context) AssessmentStatusPtrOutput {
	return o
}

func (o AssessmentStatusPtrOutput) Elem() AssessmentStatusOutput {
	return o.ApplyT(func(v *AssessmentStatus) AssessmentStatus {
		if v != nil {
			return *v
		}
		var ret AssessmentStatus
		return ret
	}).(AssessmentStatusOutput)
}

func (o AssessmentStatusPtrOutput) Cause() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentStatus) *string {
		if v == nil {
			return nil
		}
		return v.Cause
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentStatusPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentStatus) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentStatusPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentStatus) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

type AssessmentStatusResponse struct {
	Cause       *string `pulumi:"cause"`
	Code        string  `pulumi:"code"`
	Description *string `pulumi:"description"`
}





type AssessmentStatusResponseInput interface {
	pulumi.Input

	ToAssessmentStatusResponseOutput() AssessmentStatusResponseOutput
	ToAssessmentStatusResponseOutputWithContext(context.Context) AssessmentStatusResponseOutput
}

type AssessmentStatusResponseArgs struct {
	Cause       pulumi.StringPtrInput `pulumi:"cause"`
	Code        pulumi.StringInput    `pulumi:"code"`
	Description pulumi.StringPtrInput `pulumi:"description"`
}

func (AssessmentStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentStatusResponse)(nil)).Elem()
}

func (i AssessmentStatusResponseArgs) ToAssessmentStatusResponseOutput() AssessmentStatusResponseOutput {
	return i.ToAssessmentStatusResponseOutputWithContext(context.Background())
}

func (i AssessmentStatusResponseArgs) ToAssessmentStatusResponseOutputWithContext(ctx context.Context) AssessmentStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentStatusResponseOutput)
}

func (i AssessmentStatusResponseArgs) ToAssessmentStatusResponsePtrOutput() AssessmentStatusResponsePtrOutput {
	return i.ToAssessmentStatusResponsePtrOutputWithContext(context.Background())
}

func (i AssessmentStatusResponseArgs) ToAssessmentStatusResponsePtrOutputWithContext(ctx context.Context) AssessmentStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentStatusResponseOutput).ToAssessmentStatusResponsePtrOutputWithContext(ctx)
}









type AssessmentStatusResponsePtrInput interface {
	pulumi.Input

	ToAssessmentStatusResponsePtrOutput() AssessmentStatusResponsePtrOutput
	ToAssessmentStatusResponsePtrOutputWithContext(context.Context) AssessmentStatusResponsePtrOutput
}

type assessmentStatusResponsePtrType AssessmentStatusResponseArgs

func AssessmentStatusResponsePtr(v *AssessmentStatusResponseArgs) AssessmentStatusResponsePtrInput {
	return (*assessmentStatusResponsePtrType)(v)
}

func (*assessmentStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentStatusResponse)(nil)).Elem()
}

func (i *assessmentStatusResponsePtrType) ToAssessmentStatusResponsePtrOutput() AssessmentStatusResponsePtrOutput {
	return i.ToAssessmentStatusResponsePtrOutputWithContext(context.Background())
}

func (i *assessmentStatusResponsePtrType) ToAssessmentStatusResponsePtrOutputWithContext(ctx context.Context) AssessmentStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentStatusResponsePtrOutput)
}

type AssessmentStatusResponseOutput struct{ *pulumi.OutputState }

func (AssessmentStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentStatusResponse)(nil)).Elem()
}

func (o AssessmentStatusResponseOutput) ToAssessmentStatusResponseOutput() AssessmentStatusResponseOutput {
	return o
}

func (o AssessmentStatusResponseOutput) ToAssessmentStatusResponseOutputWithContext(ctx context.Context) AssessmentStatusResponseOutput {
	return o
}

func (o AssessmentStatusResponseOutput) ToAssessmentStatusResponsePtrOutput() AssessmentStatusResponsePtrOutput {
	return o.ToAssessmentStatusResponsePtrOutputWithContext(context.Background())
}

func (o AssessmentStatusResponseOutput) ToAssessmentStatusResponsePtrOutputWithContext(ctx context.Context) AssessmentStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssessmentStatusResponse) *AssessmentStatusResponse {
		return &v
	}).(AssessmentStatusResponsePtrOutput)
}

func (o AssessmentStatusResponseOutput) Cause() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssessmentStatusResponse) *string { return v.Cause }).(pulumi.StringPtrOutput)
}

func (o AssessmentStatusResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentStatusResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o AssessmentStatusResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssessmentStatusResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

type AssessmentStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (AssessmentStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentStatusResponse)(nil)).Elem()
}

func (o AssessmentStatusResponsePtrOutput) ToAssessmentStatusResponsePtrOutput() AssessmentStatusResponsePtrOutput {
	return o
}

func (o AssessmentStatusResponsePtrOutput) ToAssessmentStatusResponsePtrOutputWithContext(ctx context.Context) AssessmentStatusResponsePtrOutput {
	return o
}

func (o AssessmentStatusResponsePtrOutput) Elem() AssessmentStatusResponseOutput {
	return o.ApplyT(func(v *AssessmentStatusResponse) AssessmentStatusResponse {
		if v != nil {
			return *v
		}
		var ret AssessmentStatusResponse
		return ret
	}).(AssessmentStatusResponseOutput)
}

func (o AssessmentStatusResponsePtrOutput) Cause() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Cause
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentStatusResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentStatusResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

type AutomationActionEventHub struct {
	ActionType         string  `pulumi:"actionType"`
	ConnectionString   *string `pulumi:"connectionString"`
	EventHubResourceId *string `pulumi:"eventHubResourceId"`
}





type AutomationActionEventHubInput interface {
	pulumi.Input

	ToAutomationActionEventHubOutput() AutomationActionEventHubOutput
	ToAutomationActionEventHubOutputWithContext(context.Context) AutomationActionEventHubOutput
}

type AutomationActionEventHubArgs struct {
	ActionType         pulumi.StringInput    `pulumi:"actionType"`
	ConnectionString   pulumi.StringPtrInput `pulumi:"connectionString"`
	EventHubResourceId pulumi.StringPtrInput `pulumi:"eventHubResourceId"`
}

func (AutomationActionEventHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionEventHub)(nil)).Elem()
}

func (i AutomationActionEventHubArgs) ToAutomationActionEventHubOutput() AutomationActionEventHubOutput {
	return i.ToAutomationActionEventHubOutputWithContext(context.Background())
}

func (i AutomationActionEventHubArgs) ToAutomationActionEventHubOutputWithContext(ctx context.Context) AutomationActionEventHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationActionEventHubOutput)
}

type AutomationActionEventHubOutput struct{ *pulumi.OutputState }

func (AutomationActionEventHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionEventHub)(nil)).Elem()
}

func (o AutomationActionEventHubOutput) ToAutomationActionEventHubOutput() AutomationActionEventHubOutput {
	return o
}

func (o AutomationActionEventHubOutput) ToAutomationActionEventHubOutputWithContext(ctx context.Context) AutomationActionEventHubOutput {
	return o
}

func (o AutomationActionEventHubOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationActionEventHub) string { return v.ActionType }).(pulumi.StringOutput)
}

func (o AutomationActionEventHubOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationActionEventHub) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o AutomationActionEventHubOutput) EventHubResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationActionEventHub) *string { return v.EventHubResourceId }).(pulumi.StringPtrOutput)
}

type AutomationActionEventHubResponse struct {
	ActionType         string  `pulumi:"actionType"`
	ConnectionString   *string `pulumi:"connectionString"`
	EventHubResourceId *string `pulumi:"eventHubResourceId"`
	SasPolicyName      string  `pulumi:"sasPolicyName"`
}





type AutomationActionEventHubResponseInput interface {
	pulumi.Input

	ToAutomationActionEventHubResponseOutput() AutomationActionEventHubResponseOutput
	ToAutomationActionEventHubResponseOutputWithContext(context.Context) AutomationActionEventHubResponseOutput
}

type AutomationActionEventHubResponseArgs struct {
	ActionType         pulumi.StringInput    `pulumi:"actionType"`
	ConnectionString   pulumi.StringPtrInput `pulumi:"connectionString"`
	EventHubResourceId pulumi.StringPtrInput `pulumi:"eventHubResourceId"`
	SasPolicyName      pulumi.StringInput    `pulumi:"sasPolicyName"`
}

func (AutomationActionEventHubResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionEventHubResponse)(nil)).Elem()
}

func (i AutomationActionEventHubResponseArgs) ToAutomationActionEventHubResponseOutput() AutomationActionEventHubResponseOutput {
	return i.ToAutomationActionEventHubResponseOutputWithContext(context.Background())
}

func (i AutomationActionEventHubResponseArgs) ToAutomationActionEventHubResponseOutputWithContext(ctx context.Context) AutomationActionEventHubResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationActionEventHubResponseOutput)
}

type AutomationActionEventHubResponseOutput struct{ *pulumi.OutputState }

func (AutomationActionEventHubResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionEventHubResponse)(nil)).Elem()
}

func (o AutomationActionEventHubResponseOutput) ToAutomationActionEventHubResponseOutput() AutomationActionEventHubResponseOutput {
	return o
}

func (o AutomationActionEventHubResponseOutput) ToAutomationActionEventHubResponseOutputWithContext(ctx context.Context) AutomationActionEventHubResponseOutput {
	return o
}

func (o AutomationActionEventHubResponseOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationActionEventHubResponse) string { return v.ActionType }).(pulumi.StringOutput)
}

func (o AutomationActionEventHubResponseOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationActionEventHubResponse) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o AutomationActionEventHubResponseOutput) EventHubResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationActionEventHubResponse) *string { return v.EventHubResourceId }).(pulumi.StringPtrOutput)
}

func (o AutomationActionEventHubResponseOutput) SasPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationActionEventHubResponse) string { return v.SasPolicyName }).(pulumi.StringOutput)
}

type AutomationActionLogicApp struct {
	ActionType         string  `pulumi:"actionType"`
	LogicAppResourceId *string `pulumi:"logicAppResourceId"`
	Uri                *string `pulumi:"uri"`
}





type AutomationActionLogicAppInput interface {
	pulumi.Input

	ToAutomationActionLogicAppOutput() AutomationActionLogicAppOutput
	ToAutomationActionLogicAppOutputWithContext(context.Context) AutomationActionLogicAppOutput
}

type AutomationActionLogicAppArgs struct {
	ActionType         pulumi.StringInput    `pulumi:"actionType"`
	LogicAppResourceId pulumi.StringPtrInput `pulumi:"logicAppResourceId"`
	Uri                pulumi.StringPtrInput `pulumi:"uri"`
}

func (AutomationActionLogicAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionLogicApp)(nil)).Elem()
}

func (i AutomationActionLogicAppArgs) ToAutomationActionLogicAppOutput() AutomationActionLogicAppOutput {
	return i.ToAutomationActionLogicAppOutputWithContext(context.Background())
}

func (i AutomationActionLogicAppArgs) ToAutomationActionLogicAppOutputWithContext(ctx context.Context) AutomationActionLogicAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationActionLogicAppOutput)
}

type AutomationActionLogicAppOutput struct{ *pulumi.OutputState }

func (AutomationActionLogicAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionLogicApp)(nil)).Elem()
}

func (o AutomationActionLogicAppOutput) ToAutomationActionLogicAppOutput() AutomationActionLogicAppOutput {
	return o
}

func (o AutomationActionLogicAppOutput) ToAutomationActionLogicAppOutputWithContext(ctx context.Context) AutomationActionLogicAppOutput {
	return o
}

func (o AutomationActionLogicAppOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationActionLogicApp) string { return v.ActionType }).(pulumi.StringOutput)
}

func (o AutomationActionLogicAppOutput) LogicAppResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationActionLogicApp) *string { return v.LogicAppResourceId }).(pulumi.StringPtrOutput)
}

func (o AutomationActionLogicAppOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationActionLogicApp) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type AutomationActionLogicAppResponse struct {
	ActionType         string  `pulumi:"actionType"`
	LogicAppResourceId *string `pulumi:"logicAppResourceId"`
	Uri                *string `pulumi:"uri"`
}





type AutomationActionLogicAppResponseInput interface {
	pulumi.Input

	ToAutomationActionLogicAppResponseOutput() AutomationActionLogicAppResponseOutput
	ToAutomationActionLogicAppResponseOutputWithContext(context.Context) AutomationActionLogicAppResponseOutput
}

type AutomationActionLogicAppResponseArgs struct {
	ActionType         pulumi.StringInput    `pulumi:"actionType"`
	LogicAppResourceId pulumi.StringPtrInput `pulumi:"logicAppResourceId"`
	Uri                pulumi.StringPtrInput `pulumi:"uri"`
}

func (AutomationActionLogicAppResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionLogicAppResponse)(nil)).Elem()
}

func (i AutomationActionLogicAppResponseArgs) ToAutomationActionLogicAppResponseOutput() AutomationActionLogicAppResponseOutput {
	return i.ToAutomationActionLogicAppResponseOutputWithContext(context.Background())
}

func (i AutomationActionLogicAppResponseArgs) ToAutomationActionLogicAppResponseOutputWithContext(ctx context.Context) AutomationActionLogicAppResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationActionLogicAppResponseOutput)
}

type AutomationActionLogicAppResponseOutput struct{ *pulumi.OutputState }

func (AutomationActionLogicAppResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionLogicAppResponse)(nil)).Elem()
}

func (o AutomationActionLogicAppResponseOutput) ToAutomationActionLogicAppResponseOutput() AutomationActionLogicAppResponseOutput {
	return o
}

func (o AutomationActionLogicAppResponseOutput) ToAutomationActionLogicAppResponseOutputWithContext(ctx context.Context) AutomationActionLogicAppResponseOutput {
	return o
}

func (o AutomationActionLogicAppResponseOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationActionLogicAppResponse) string { return v.ActionType }).(pulumi.StringOutput)
}

func (o AutomationActionLogicAppResponseOutput) LogicAppResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationActionLogicAppResponse) *string { return v.LogicAppResourceId }).(pulumi.StringPtrOutput)
}

func (o AutomationActionLogicAppResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationActionLogicAppResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type AutomationActionWorkspace struct {
	ActionType          string  `pulumi:"actionType"`
	WorkspaceResourceId *string `pulumi:"workspaceResourceId"`
}





type AutomationActionWorkspaceInput interface {
	pulumi.Input

	ToAutomationActionWorkspaceOutput() AutomationActionWorkspaceOutput
	ToAutomationActionWorkspaceOutputWithContext(context.Context) AutomationActionWorkspaceOutput
}

type AutomationActionWorkspaceArgs struct {
	ActionType          pulumi.StringInput    `pulumi:"actionType"`
	WorkspaceResourceId pulumi.StringPtrInput `pulumi:"workspaceResourceId"`
}

func (AutomationActionWorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionWorkspace)(nil)).Elem()
}

func (i AutomationActionWorkspaceArgs) ToAutomationActionWorkspaceOutput() AutomationActionWorkspaceOutput {
	return i.ToAutomationActionWorkspaceOutputWithContext(context.Background())
}

func (i AutomationActionWorkspaceArgs) ToAutomationActionWorkspaceOutputWithContext(ctx context.Context) AutomationActionWorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationActionWorkspaceOutput)
}

type AutomationActionWorkspaceOutput struct{ *pulumi.OutputState }

func (AutomationActionWorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionWorkspace)(nil)).Elem()
}

func (o AutomationActionWorkspaceOutput) ToAutomationActionWorkspaceOutput() AutomationActionWorkspaceOutput {
	return o
}

func (o AutomationActionWorkspaceOutput) ToAutomationActionWorkspaceOutputWithContext(ctx context.Context) AutomationActionWorkspaceOutput {
	return o
}

func (o AutomationActionWorkspaceOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationActionWorkspace) string { return v.ActionType }).(pulumi.StringOutput)
}

func (o AutomationActionWorkspaceOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationActionWorkspace) *string { return v.WorkspaceResourceId }).(pulumi.StringPtrOutput)
}

type AutomationActionWorkspaceResponse struct {
	ActionType          string  `pulumi:"actionType"`
	WorkspaceResourceId *string `pulumi:"workspaceResourceId"`
}





type AutomationActionWorkspaceResponseInput interface {
	pulumi.Input

	ToAutomationActionWorkspaceResponseOutput() AutomationActionWorkspaceResponseOutput
	ToAutomationActionWorkspaceResponseOutputWithContext(context.Context) AutomationActionWorkspaceResponseOutput
}

type AutomationActionWorkspaceResponseArgs struct {
	ActionType          pulumi.StringInput    `pulumi:"actionType"`
	WorkspaceResourceId pulumi.StringPtrInput `pulumi:"workspaceResourceId"`
}

func (AutomationActionWorkspaceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionWorkspaceResponse)(nil)).Elem()
}

func (i AutomationActionWorkspaceResponseArgs) ToAutomationActionWorkspaceResponseOutput() AutomationActionWorkspaceResponseOutput {
	return i.ToAutomationActionWorkspaceResponseOutputWithContext(context.Background())
}

func (i AutomationActionWorkspaceResponseArgs) ToAutomationActionWorkspaceResponseOutputWithContext(ctx context.Context) AutomationActionWorkspaceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationActionWorkspaceResponseOutput)
}

type AutomationActionWorkspaceResponseOutput struct{ *pulumi.OutputState }

func (AutomationActionWorkspaceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionWorkspaceResponse)(nil)).Elem()
}

func (o AutomationActionWorkspaceResponseOutput) ToAutomationActionWorkspaceResponseOutput() AutomationActionWorkspaceResponseOutput {
	return o
}

func (o AutomationActionWorkspaceResponseOutput) ToAutomationActionWorkspaceResponseOutputWithContext(ctx context.Context) AutomationActionWorkspaceResponseOutput {
	return o
}

func (o AutomationActionWorkspaceResponseOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationActionWorkspaceResponse) string { return v.ActionType }).(pulumi.StringOutput)
}

func (o AutomationActionWorkspaceResponseOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationActionWorkspaceResponse) *string { return v.WorkspaceResourceId }).(pulumi.StringPtrOutput)
}

type AutomationRuleSet struct {
	Rules []AutomationTriggeringRule `pulumi:"rules"`
}





type AutomationRuleSetInput interface {
	pulumi.Input

	ToAutomationRuleSetOutput() AutomationRuleSetOutput
	ToAutomationRuleSetOutputWithContext(context.Context) AutomationRuleSetOutput
}

type AutomationRuleSetArgs struct {
	Rules AutomationTriggeringRuleArrayInput `pulumi:"rules"`
}

func (AutomationRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleSet)(nil)).Elem()
}

func (i AutomationRuleSetArgs) ToAutomationRuleSetOutput() AutomationRuleSetOutput {
	return i.ToAutomationRuleSetOutputWithContext(context.Background())
}

func (i AutomationRuleSetArgs) ToAutomationRuleSetOutputWithContext(ctx context.Context) AutomationRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleSetOutput)
}





type AutomationRuleSetArrayInput interface {
	pulumi.Input

	ToAutomationRuleSetArrayOutput() AutomationRuleSetArrayOutput
	ToAutomationRuleSetArrayOutputWithContext(context.Context) AutomationRuleSetArrayOutput
}

type AutomationRuleSetArray []AutomationRuleSetInput

func (AutomationRuleSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationRuleSet)(nil)).Elem()
}

func (i AutomationRuleSetArray) ToAutomationRuleSetArrayOutput() AutomationRuleSetArrayOutput {
	return i.ToAutomationRuleSetArrayOutputWithContext(context.Background())
}

func (i AutomationRuleSetArray) ToAutomationRuleSetArrayOutputWithContext(ctx context.Context) AutomationRuleSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleSetArrayOutput)
}

type AutomationRuleSetOutput struct{ *pulumi.OutputState }

func (AutomationRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleSet)(nil)).Elem()
}

func (o AutomationRuleSetOutput) ToAutomationRuleSetOutput() AutomationRuleSetOutput {
	return o
}

func (o AutomationRuleSetOutput) ToAutomationRuleSetOutputWithContext(ctx context.Context) AutomationRuleSetOutput {
	return o
}

func (o AutomationRuleSetOutput) Rules() AutomationTriggeringRuleArrayOutput {
	return o.ApplyT(func(v AutomationRuleSet) []AutomationTriggeringRule { return v.Rules }).(AutomationTriggeringRuleArrayOutput)
}

type AutomationRuleSetArrayOutput struct{ *pulumi.OutputState }

func (AutomationRuleSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationRuleSet)(nil)).Elem()
}

func (o AutomationRuleSetArrayOutput) ToAutomationRuleSetArrayOutput() AutomationRuleSetArrayOutput {
	return o
}

func (o AutomationRuleSetArrayOutput) ToAutomationRuleSetArrayOutputWithContext(ctx context.Context) AutomationRuleSetArrayOutput {
	return o
}

func (o AutomationRuleSetArrayOutput) Index(i pulumi.IntInput) AutomationRuleSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutomationRuleSet {
		return vs[0].([]AutomationRuleSet)[vs[1].(int)]
	}).(AutomationRuleSetOutput)
}

type AutomationRuleSetResponse struct {
	Rules []AutomationTriggeringRuleResponse `pulumi:"rules"`
}





type AutomationRuleSetResponseInput interface {
	pulumi.Input

	ToAutomationRuleSetResponseOutput() AutomationRuleSetResponseOutput
	ToAutomationRuleSetResponseOutputWithContext(context.Context) AutomationRuleSetResponseOutput
}

type AutomationRuleSetResponseArgs struct {
	Rules AutomationTriggeringRuleResponseArrayInput `pulumi:"rules"`
}

func (AutomationRuleSetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleSetResponse)(nil)).Elem()
}

func (i AutomationRuleSetResponseArgs) ToAutomationRuleSetResponseOutput() AutomationRuleSetResponseOutput {
	return i.ToAutomationRuleSetResponseOutputWithContext(context.Background())
}

func (i AutomationRuleSetResponseArgs) ToAutomationRuleSetResponseOutputWithContext(ctx context.Context) AutomationRuleSetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleSetResponseOutput)
}





type AutomationRuleSetResponseArrayInput interface {
	pulumi.Input

	ToAutomationRuleSetResponseArrayOutput() AutomationRuleSetResponseArrayOutput
	ToAutomationRuleSetResponseArrayOutputWithContext(context.Context) AutomationRuleSetResponseArrayOutput
}

type AutomationRuleSetResponseArray []AutomationRuleSetResponseInput

func (AutomationRuleSetResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationRuleSetResponse)(nil)).Elem()
}

func (i AutomationRuleSetResponseArray) ToAutomationRuleSetResponseArrayOutput() AutomationRuleSetResponseArrayOutput {
	return i.ToAutomationRuleSetResponseArrayOutputWithContext(context.Background())
}

func (i AutomationRuleSetResponseArray) ToAutomationRuleSetResponseArrayOutputWithContext(ctx context.Context) AutomationRuleSetResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleSetResponseArrayOutput)
}

type AutomationRuleSetResponseOutput struct{ *pulumi.OutputState }

func (AutomationRuleSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleSetResponse)(nil)).Elem()
}

func (o AutomationRuleSetResponseOutput) ToAutomationRuleSetResponseOutput() AutomationRuleSetResponseOutput {
	return o
}

func (o AutomationRuleSetResponseOutput) ToAutomationRuleSetResponseOutputWithContext(ctx context.Context) AutomationRuleSetResponseOutput {
	return o
}

func (o AutomationRuleSetResponseOutput) Rules() AutomationTriggeringRuleResponseArrayOutput {
	return o.ApplyT(func(v AutomationRuleSetResponse) []AutomationTriggeringRuleResponse { return v.Rules }).(AutomationTriggeringRuleResponseArrayOutput)
}

type AutomationRuleSetResponseArrayOutput struct{ *pulumi.OutputState }

func (AutomationRuleSetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationRuleSetResponse)(nil)).Elem()
}

func (o AutomationRuleSetResponseArrayOutput) ToAutomationRuleSetResponseArrayOutput() AutomationRuleSetResponseArrayOutput {
	return o
}

func (o AutomationRuleSetResponseArrayOutput) ToAutomationRuleSetResponseArrayOutputWithContext(ctx context.Context) AutomationRuleSetResponseArrayOutput {
	return o
}

func (o AutomationRuleSetResponseArrayOutput) Index(i pulumi.IntInput) AutomationRuleSetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutomationRuleSetResponse {
		return vs[0].([]AutomationRuleSetResponse)[vs[1].(int)]
	}).(AutomationRuleSetResponseOutput)
}

type AutomationScope struct {
	Description *string `pulumi:"description"`
	ScopePath   *string `pulumi:"scopePath"`
}





type AutomationScopeInput interface {
	pulumi.Input

	ToAutomationScopeOutput() AutomationScopeOutput
	ToAutomationScopeOutputWithContext(context.Context) AutomationScopeOutput
}

type AutomationScopeArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	ScopePath   pulumi.StringPtrInput `pulumi:"scopePath"`
}

func (AutomationScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationScope)(nil)).Elem()
}

func (i AutomationScopeArgs) ToAutomationScopeOutput() AutomationScopeOutput {
	return i.ToAutomationScopeOutputWithContext(context.Background())
}

func (i AutomationScopeArgs) ToAutomationScopeOutputWithContext(ctx context.Context) AutomationScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationScopeOutput)
}





type AutomationScopeArrayInput interface {
	pulumi.Input

	ToAutomationScopeArrayOutput() AutomationScopeArrayOutput
	ToAutomationScopeArrayOutputWithContext(context.Context) AutomationScopeArrayOutput
}

type AutomationScopeArray []AutomationScopeInput

func (AutomationScopeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationScope)(nil)).Elem()
}

func (i AutomationScopeArray) ToAutomationScopeArrayOutput() AutomationScopeArrayOutput {
	return i.ToAutomationScopeArrayOutputWithContext(context.Background())
}

func (i AutomationScopeArray) ToAutomationScopeArrayOutputWithContext(ctx context.Context) AutomationScopeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationScopeArrayOutput)
}

type AutomationScopeOutput struct{ *pulumi.OutputState }

func (AutomationScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationScope)(nil)).Elem()
}

func (o AutomationScopeOutput) ToAutomationScopeOutput() AutomationScopeOutput {
	return o
}

func (o AutomationScopeOutput) ToAutomationScopeOutputWithContext(ctx context.Context) AutomationScopeOutput {
	return o
}

func (o AutomationScopeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationScope) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AutomationScopeOutput) ScopePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationScope) *string { return v.ScopePath }).(pulumi.StringPtrOutput)
}

type AutomationScopeArrayOutput struct{ *pulumi.OutputState }

func (AutomationScopeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationScope)(nil)).Elem()
}

func (o AutomationScopeArrayOutput) ToAutomationScopeArrayOutput() AutomationScopeArrayOutput {
	return o
}

func (o AutomationScopeArrayOutput) ToAutomationScopeArrayOutputWithContext(ctx context.Context) AutomationScopeArrayOutput {
	return o
}

func (o AutomationScopeArrayOutput) Index(i pulumi.IntInput) AutomationScopeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutomationScope {
		return vs[0].([]AutomationScope)[vs[1].(int)]
	}).(AutomationScopeOutput)
}

type AutomationScopeResponse struct {
	Description *string `pulumi:"description"`
	ScopePath   *string `pulumi:"scopePath"`
}





type AutomationScopeResponseInput interface {
	pulumi.Input

	ToAutomationScopeResponseOutput() AutomationScopeResponseOutput
	ToAutomationScopeResponseOutputWithContext(context.Context) AutomationScopeResponseOutput
}

type AutomationScopeResponseArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	ScopePath   pulumi.StringPtrInput `pulumi:"scopePath"`
}

func (AutomationScopeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationScopeResponse)(nil)).Elem()
}

func (i AutomationScopeResponseArgs) ToAutomationScopeResponseOutput() AutomationScopeResponseOutput {
	return i.ToAutomationScopeResponseOutputWithContext(context.Background())
}

func (i AutomationScopeResponseArgs) ToAutomationScopeResponseOutputWithContext(ctx context.Context) AutomationScopeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationScopeResponseOutput)
}





type AutomationScopeResponseArrayInput interface {
	pulumi.Input

	ToAutomationScopeResponseArrayOutput() AutomationScopeResponseArrayOutput
	ToAutomationScopeResponseArrayOutputWithContext(context.Context) AutomationScopeResponseArrayOutput
}

type AutomationScopeResponseArray []AutomationScopeResponseInput

func (AutomationScopeResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationScopeResponse)(nil)).Elem()
}

func (i AutomationScopeResponseArray) ToAutomationScopeResponseArrayOutput() AutomationScopeResponseArrayOutput {
	return i.ToAutomationScopeResponseArrayOutputWithContext(context.Background())
}

func (i AutomationScopeResponseArray) ToAutomationScopeResponseArrayOutputWithContext(ctx context.Context) AutomationScopeResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationScopeResponseArrayOutput)
}

type AutomationScopeResponseOutput struct{ *pulumi.OutputState }

func (AutomationScopeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationScopeResponse)(nil)).Elem()
}

func (o AutomationScopeResponseOutput) ToAutomationScopeResponseOutput() AutomationScopeResponseOutput {
	return o
}

func (o AutomationScopeResponseOutput) ToAutomationScopeResponseOutputWithContext(ctx context.Context) AutomationScopeResponseOutput {
	return o
}

func (o AutomationScopeResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationScopeResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AutomationScopeResponseOutput) ScopePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationScopeResponse) *string { return v.ScopePath }).(pulumi.StringPtrOutput)
}

type AutomationScopeResponseArrayOutput struct{ *pulumi.OutputState }

func (AutomationScopeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationScopeResponse)(nil)).Elem()
}

func (o AutomationScopeResponseArrayOutput) ToAutomationScopeResponseArrayOutput() AutomationScopeResponseArrayOutput {
	return o
}

func (o AutomationScopeResponseArrayOutput) ToAutomationScopeResponseArrayOutputWithContext(ctx context.Context) AutomationScopeResponseArrayOutput {
	return o
}

func (o AutomationScopeResponseArrayOutput) Index(i pulumi.IntInput) AutomationScopeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutomationScopeResponse {
		return vs[0].([]AutomationScopeResponse)[vs[1].(int)]
	}).(AutomationScopeResponseOutput)
}

type AutomationSource struct {
	EventSource *string             `pulumi:"eventSource"`
	RuleSets    []AutomationRuleSet `pulumi:"ruleSets"`
}





type AutomationSourceInput interface {
	pulumi.Input

	ToAutomationSourceOutput() AutomationSourceOutput
	ToAutomationSourceOutputWithContext(context.Context) AutomationSourceOutput
}

type AutomationSourceArgs struct {
	EventSource pulumi.StringPtrInput       `pulumi:"eventSource"`
	RuleSets    AutomationRuleSetArrayInput `pulumi:"ruleSets"`
}

func (AutomationSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationSource)(nil)).Elem()
}

func (i AutomationSourceArgs) ToAutomationSourceOutput() AutomationSourceOutput {
	return i.ToAutomationSourceOutputWithContext(context.Background())
}

func (i AutomationSourceArgs) ToAutomationSourceOutputWithContext(ctx context.Context) AutomationSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationSourceOutput)
}





type AutomationSourceArrayInput interface {
	pulumi.Input

	ToAutomationSourceArrayOutput() AutomationSourceArrayOutput
	ToAutomationSourceArrayOutputWithContext(context.Context) AutomationSourceArrayOutput
}

type AutomationSourceArray []AutomationSourceInput

func (AutomationSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationSource)(nil)).Elem()
}

func (i AutomationSourceArray) ToAutomationSourceArrayOutput() AutomationSourceArrayOutput {
	return i.ToAutomationSourceArrayOutputWithContext(context.Background())
}

func (i AutomationSourceArray) ToAutomationSourceArrayOutputWithContext(ctx context.Context) AutomationSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationSourceArrayOutput)
}

type AutomationSourceOutput struct{ *pulumi.OutputState }

func (AutomationSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationSource)(nil)).Elem()
}

func (o AutomationSourceOutput) ToAutomationSourceOutput() AutomationSourceOutput {
	return o
}

func (o AutomationSourceOutput) ToAutomationSourceOutputWithContext(ctx context.Context) AutomationSourceOutput {
	return o
}

func (o AutomationSourceOutput) EventSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationSource) *string { return v.EventSource }).(pulumi.StringPtrOutput)
}

func (o AutomationSourceOutput) RuleSets() AutomationRuleSetArrayOutput {
	return o.ApplyT(func(v AutomationSource) []AutomationRuleSet { return v.RuleSets }).(AutomationRuleSetArrayOutput)
}

type AutomationSourceArrayOutput struct{ *pulumi.OutputState }

func (AutomationSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationSource)(nil)).Elem()
}

func (o AutomationSourceArrayOutput) ToAutomationSourceArrayOutput() AutomationSourceArrayOutput {
	return o
}

func (o AutomationSourceArrayOutput) ToAutomationSourceArrayOutputWithContext(ctx context.Context) AutomationSourceArrayOutput {
	return o
}

func (o AutomationSourceArrayOutput) Index(i pulumi.IntInput) AutomationSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutomationSource {
		return vs[0].([]AutomationSource)[vs[1].(int)]
	}).(AutomationSourceOutput)
}

type AutomationSourceResponse struct {
	EventSource *string                     `pulumi:"eventSource"`
	RuleSets    []AutomationRuleSetResponse `pulumi:"ruleSets"`
}





type AutomationSourceResponseInput interface {
	pulumi.Input

	ToAutomationSourceResponseOutput() AutomationSourceResponseOutput
	ToAutomationSourceResponseOutputWithContext(context.Context) AutomationSourceResponseOutput
}

type AutomationSourceResponseArgs struct {
	EventSource pulumi.StringPtrInput               `pulumi:"eventSource"`
	RuleSets    AutomationRuleSetResponseArrayInput `pulumi:"ruleSets"`
}

func (AutomationSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationSourceResponse)(nil)).Elem()
}

func (i AutomationSourceResponseArgs) ToAutomationSourceResponseOutput() AutomationSourceResponseOutput {
	return i.ToAutomationSourceResponseOutputWithContext(context.Background())
}

func (i AutomationSourceResponseArgs) ToAutomationSourceResponseOutputWithContext(ctx context.Context) AutomationSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationSourceResponseOutput)
}





type AutomationSourceResponseArrayInput interface {
	pulumi.Input

	ToAutomationSourceResponseArrayOutput() AutomationSourceResponseArrayOutput
	ToAutomationSourceResponseArrayOutputWithContext(context.Context) AutomationSourceResponseArrayOutput
}

type AutomationSourceResponseArray []AutomationSourceResponseInput

func (AutomationSourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationSourceResponse)(nil)).Elem()
}

func (i AutomationSourceResponseArray) ToAutomationSourceResponseArrayOutput() AutomationSourceResponseArrayOutput {
	return i.ToAutomationSourceResponseArrayOutputWithContext(context.Background())
}

func (i AutomationSourceResponseArray) ToAutomationSourceResponseArrayOutputWithContext(ctx context.Context) AutomationSourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationSourceResponseArrayOutput)
}

type AutomationSourceResponseOutput struct{ *pulumi.OutputState }

func (AutomationSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationSourceResponse)(nil)).Elem()
}

func (o AutomationSourceResponseOutput) ToAutomationSourceResponseOutput() AutomationSourceResponseOutput {
	return o
}

func (o AutomationSourceResponseOutput) ToAutomationSourceResponseOutputWithContext(ctx context.Context) AutomationSourceResponseOutput {
	return o
}

func (o AutomationSourceResponseOutput) EventSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationSourceResponse) *string { return v.EventSource }).(pulumi.StringPtrOutput)
}

func (o AutomationSourceResponseOutput) RuleSets() AutomationRuleSetResponseArrayOutput {
	return o.ApplyT(func(v AutomationSourceResponse) []AutomationRuleSetResponse { return v.RuleSets }).(AutomationRuleSetResponseArrayOutput)
}

type AutomationSourceResponseArrayOutput struct{ *pulumi.OutputState }

func (AutomationSourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationSourceResponse)(nil)).Elem()
}

func (o AutomationSourceResponseArrayOutput) ToAutomationSourceResponseArrayOutput() AutomationSourceResponseArrayOutput {
	return o
}

func (o AutomationSourceResponseArrayOutput) ToAutomationSourceResponseArrayOutputWithContext(ctx context.Context) AutomationSourceResponseArrayOutput {
	return o
}

func (o AutomationSourceResponseArrayOutput) Index(i pulumi.IntInput) AutomationSourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutomationSourceResponse {
		return vs[0].([]AutomationSourceResponse)[vs[1].(int)]
	}).(AutomationSourceResponseOutput)
}

type AutomationTriggeringRule struct {
	ExpectedValue *string `pulumi:"expectedValue"`
	Operator      *string `pulumi:"operator"`
	PropertyJPath *string `pulumi:"propertyJPath"`
	PropertyType  *string `pulumi:"propertyType"`
}





type AutomationTriggeringRuleInput interface {
	pulumi.Input

	ToAutomationTriggeringRuleOutput() AutomationTriggeringRuleOutput
	ToAutomationTriggeringRuleOutputWithContext(context.Context) AutomationTriggeringRuleOutput
}

type AutomationTriggeringRuleArgs struct {
	ExpectedValue pulumi.StringPtrInput `pulumi:"expectedValue"`
	Operator      pulumi.StringPtrInput `pulumi:"operator"`
	PropertyJPath pulumi.StringPtrInput `pulumi:"propertyJPath"`
	PropertyType  pulumi.StringPtrInput `pulumi:"propertyType"`
}

func (AutomationTriggeringRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationTriggeringRule)(nil)).Elem()
}

func (i AutomationTriggeringRuleArgs) ToAutomationTriggeringRuleOutput() AutomationTriggeringRuleOutput {
	return i.ToAutomationTriggeringRuleOutputWithContext(context.Background())
}

func (i AutomationTriggeringRuleArgs) ToAutomationTriggeringRuleOutputWithContext(ctx context.Context) AutomationTriggeringRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationTriggeringRuleOutput)
}





type AutomationTriggeringRuleArrayInput interface {
	pulumi.Input

	ToAutomationTriggeringRuleArrayOutput() AutomationTriggeringRuleArrayOutput
	ToAutomationTriggeringRuleArrayOutputWithContext(context.Context) AutomationTriggeringRuleArrayOutput
}

type AutomationTriggeringRuleArray []AutomationTriggeringRuleInput

func (AutomationTriggeringRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationTriggeringRule)(nil)).Elem()
}

func (i AutomationTriggeringRuleArray) ToAutomationTriggeringRuleArrayOutput() AutomationTriggeringRuleArrayOutput {
	return i.ToAutomationTriggeringRuleArrayOutputWithContext(context.Background())
}

func (i AutomationTriggeringRuleArray) ToAutomationTriggeringRuleArrayOutputWithContext(ctx context.Context) AutomationTriggeringRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationTriggeringRuleArrayOutput)
}

type AutomationTriggeringRuleOutput struct{ *pulumi.OutputState }

func (AutomationTriggeringRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationTriggeringRule)(nil)).Elem()
}

func (o AutomationTriggeringRuleOutput) ToAutomationTriggeringRuleOutput() AutomationTriggeringRuleOutput {
	return o
}

func (o AutomationTriggeringRuleOutput) ToAutomationTriggeringRuleOutputWithContext(ctx context.Context) AutomationTriggeringRuleOutput {
	return o
}

func (o AutomationTriggeringRuleOutput) ExpectedValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationTriggeringRule) *string { return v.ExpectedValue }).(pulumi.StringPtrOutput)
}

func (o AutomationTriggeringRuleOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationTriggeringRule) *string { return v.Operator }).(pulumi.StringPtrOutput)
}

func (o AutomationTriggeringRuleOutput) PropertyJPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationTriggeringRule) *string { return v.PropertyJPath }).(pulumi.StringPtrOutput)
}

func (o AutomationTriggeringRuleOutput) PropertyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationTriggeringRule) *string { return v.PropertyType }).(pulumi.StringPtrOutput)
}

type AutomationTriggeringRuleArrayOutput struct{ *pulumi.OutputState }

func (AutomationTriggeringRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationTriggeringRule)(nil)).Elem()
}

func (o AutomationTriggeringRuleArrayOutput) ToAutomationTriggeringRuleArrayOutput() AutomationTriggeringRuleArrayOutput {
	return o
}

func (o AutomationTriggeringRuleArrayOutput) ToAutomationTriggeringRuleArrayOutputWithContext(ctx context.Context) AutomationTriggeringRuleArrayOutput {
	return o
}

func (o AutomationTriggeringRuleArrayOutput) Index(i pulumi.IntInput) AutomationTriggeringRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutomationTriggeringRule {
		return vs[0].([]AutomationTriggeringRule)[vs[1].(int)]
	}).(AutomationTriggeringRuleOutput)
}

type AutomationTriggeringRuleResponse struct {
	ExpectedValue *string `pulumi:"expectedValue"`
	Operator      *string `pulumi:"operator"`
	PropertyJPath *string `pulumi:"propertyJPath"`
	PropertyType  *string `pulumi:"propertyType"`
}





type AutomationTriggeringRuleResponseInput interface {
	pulumi.Input

	ToAutomationTriggeringRuleResponseOutput() AutomationTriggeringRuleResponseOutput
	ToAutomationTriggeringRuleResponseOutputWithContext(context.Context) AutomationTriggeringRuleResponseOutput
}

type AutomationTriggeringRuleResponseArgs struct {
	ExpectedValue pulumi.StringPtrInput `pulumi:"expectedValue"`
	Operator      pulumi.StringPtrInput `pulumi:"operator"`
	PropertyJPath pulumi.StringPtrInput `pulumi:"propertyJPath"`
	PropertyType  pulumi.StringPtrInput `pulumi:"propertyType"`
}

func (AutomationTriggeringRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationTriggeringRuleResponse)(nil)).Elem()
}

func (i AutomationTriggeringRuleResponseArgs) ToAutomationTriggeringRuleResponseOutput() AutomationTriggeringRuleResponseOutput {
	return i.ToAutomationTriggeringRuleResponseOutputWithContext(context.Background())
}

func (i AutomationTriggeringRuleResponseArgs) ToAutomationTriggeringRuleResponseOutputWithContext(ctx context.Context) AutomationTriggeringRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationTriggeringRuleResponseOutput)
}





type AutomationTriggeringRuleResponseArrayInput interface {
	pulumi.Input

	ToAutomationTriggeringRuleResponseArrayOutput() AutomationTriggeringRuleResponseArrayOutput
	ToAutomationTriggeringRuleResponseArrayOutputWithContext(context.Context) AutomationTriggeringRuleResponseArrayOutput
}

type AutomationTriggeringRuleResponseArray []AutomationTriggeringRuleResponseInput

func (AutomationTriggeringRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationTriggeringRuleResponse)(nil)).Elem()
}

func (i AutomationTriggeringRuleResponseArray) ToAutomationTriggeringRuleResponseArrayOutput() AutomationTriggeringRuleResponseArrayOutput {
	return i.ToAutomationTriggeringRuleResponseArrayOutputWithContext(context.Background())
}

func (i AutomationTriggeringRuleResponseArray) ToAutomationTriggeringRuleResponseArrayOutputWithContext(ctx context.Context) AutomationTriggeringRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationTriggeringRuleResponseArrayOutput)
}

type AutomationTriggeringRuleResponseOutput struct{ *pulumi.OutputState }

func (AutomationTriggeringRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationTriggeringRuleResponse)(nil)).Elem()
}

func (o AutomationTriggeringRuleResponseOutput) ToAutomationTriggeringRuleResponseOutput() AutomationTriggeringRuleResponseOutput {
	return o
}

func (o AutomationTriggeringRuleResponseOutput) ToAutomationTriggeringRuleResponseOutputWithContext(ctx context.Context) AutomationTriggeringRuleResponseOutput {
	return o
}

func (o AutomationTriggeringRuleResponseOutput) ExpectedValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationTriggeringRuleResponse) *string { return v.ExpectedValue }).(pulumi.StringPtrOutput)
}

func (o AutomationTriggeringRuleResponseOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationTriggeringRuleResponse) *string { return v.Operator }).(pulumi.StringPtrOutput)
}

func (o AutomationTriggeringRuleResponseOutput) PropertyJPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationTriggeringRuleResponse) *string { return v.PropertyJPath }).(pulumi.StringPtrOutput)
}

func (o AutomationTriggeringRuleResponseOutput) PropertyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationTriggeringRuleResponse) *string { return v.PropertyType }).(pulumi.StringPtrOutput)
}

type AutomationTriggeringRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (AutomationTriggeringRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationTriggeringRuleResponse)(nil)).Elem()
}

func (o AutomationTriggeringRuleResponseArrayOutput) ToAutomationTriggeringRuleResponseArrayOutput() AutomationTriggeringRuleResponseArrayOutput {
	return o
}

func (o AutomationTriggeringRuleResponseArrayOutput) ToAutomationTriggeringRuleResponseArrayOutputWithContext(ctx context.Context) AutomationTriggeringRuleResponseArrayOutput {
	return o
}

func (o AutomationTriggeringRuleResponseArrayOutput) Index(i pulumi.IntInput) AutomationTriggeringRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutomationTriggeringRuleResponse {
		return vs[0].([]AutomationTriggeringRuleResponse)[vs[1].(int)]
	}).(AutomationTriggeringRuleResponseOutput)
}

type AzureResourceDetails struct {
	Source string `pulumi:"source"`
}





type AzureResourceDetailsInput interface {
	pulumi.Input

	ToAzureResourceDetailsOutput() AzureResourceDetailsOutput
	ToAzureResourceDetailsOutputWithContext(context.Context) AzureResourceDetailsOutput
}

type AzureResourceDetailsArgs struct {
	Source pulumi.StringInput `pulumi:"source"`
}

func (AzureResourceDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceDetails)(nil)).Elem()
}

func (i AzureResourceDetailsArgs) ToAzureResourceDetailsOutput() AzureResourceDetailsOutput {
	return i.ToAzureResourceDetailsOutputWithContext(context.Background())
}

func (i AzureResourceDetailsArgs) ToAzureResourceDetailsOutputWithContext(ctx context.Context) AzureResourceDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureResourceDetailsOutput)
}

type AzureResourceDetailsOutput struct{ *pulumi.OutputState }

func (AzureResourceDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceDetails)(nil)).Elem()
}

func (o AzureResourceDetailsOutput) ToAzureResourceDetailsOutput() AzureResourceDetailsOutput {
	return o
}

func (o AzureResourceDetailsOutput) ToAzureResourceDetailsOutputWithContext(ctx context.Context) AzureResourceDetailsOutput {
	return o
}

func (o AzureResourceDetailsOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceDetails) string { return v.Source }).(pulumi.StringOutput)
}

type AzureResourceDetailsResponse struct {
	Id     string `pulumi:"id"`
	Source string `pulumi:"source"`
}





type AzureResourceDetailsResponseInput interface {
	pulumi.Input

	ToAzureResourceDetailsResponseOutput() AzureResourceDetailsResponseOutput
	ToAzureResourceDetailsResponseOutputWithContext(context.Context) AzureResourceDetailsResponseOutput
}

type AzureResourceDetailsResponseArgs struct {
	Id     pulumi.StringInput `pulumi:"id"`
	Source pulumi.StringInput `pulumi:"source"`
}

func (AzureResourceDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceDetailsResponse)(nil)).Elem()
}

func (i AzureResourceDetailsResponseArgs) ToAzureResourceDetailsResponseOutput() AzureResourceDetailsResponseOutput {
	return i.ToAzureResourceDetailsResponseOutputWithContext(context.Background())
}

func (i AzureResourceDetailsResponseArgs) ToAzureResourceDetailsResponseOutputWithContext(ctx context.Context) AzureResourceDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureResourceDetailsResponseOutput)
}

type AzureResourceDetailsResponseOutput struct{ *pulumi.OutputState }

func (AzureResourceDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceDetailsResponse)(nil)).Elem()
}

func (o AzureResourceDetailsResponseOutput) ToAzureResourceDetailsResponseOutput() AzureResourceDetailsResponseOutput {
	return o
}

func (o AzureResourceDetailsResponseOutput) ToAzureResourceDetailsResponseOutputWithContext(ctx context.Context) AzureResourceDetailsResponseOutput {
	return o
}

func (o AzureResourceDetailsResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceDetailsResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o AzureResourceDetailsResponseOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceDetailsResponse) string { return v.Source }).(pulumi.StringOutput)
}

type OnPremiseResourceDetails struct {
	MachineName      string `pulumi:"machineName"`
	Source           string `pulumi:"source"`
	SourceComputerId string `pulumi:"sourceComputerId"`
	Vmuuid           string `pulumi:"vmuuid"`
	WorkspaceId      string `pulumi:"workspaceId"`
}





type OnPremiseResourceDetailsInput interface {
	pulumi.Input

	ToOnPremiseResourceDetailsOutput() OnPremiseResourceDetailsOutput
	ToOnPremiseResourceDetailsOutputWithContext(context.Context) OnPremiseResourceDetailsOutput
}

type OnPremiseResourceDetailsArgs struct {
	MachineName      pulumi.StringInput `pulumi:"machineName"`
	Source           pulumi.StringInput `pulumi:"source"`
	SourceComputerId pulumi.StringInput `pulumi:"sourceComputerId"`
	Vmuuid           pulumi.StringInput `pulumi:"vmuuid"`
	WorkspaceId      pulumi.StringInput `pulumi:"workspaceId"`
}

func (OnPremiseResourceDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseResourceDetails)(nil)).Elem()
}

func (i OnPremiseResourceDetailsArgs) ToOnPremiseResourceDetailsOutput() OnPremiseResourceDetailsOutput {
	return i.ToOnPremiseResourceDetailsOutputWithContext(context.Background())
}

func (i OnPremiseResourceDetailsArgs) ToOnPremiseResourceDetailsOutputWithContext(ctx context.Context) OnPremiseResourceDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnPremiseResourceDetailsOutput)
}

type OnPremiseResourceDetailsOutput struct{ *pulumi.OutputState }

func (OnPremiseResourceDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseResourceDetails)(nil)).Elem()
}

func (o OnPremiseResourceDetailsOutput) ToOnPremiseResourceDetailsOutput() OnPremiseResourceDetailsOutput {
	return o
}

func (o OnPremiseResourceDetailsOutput) ToOnPremiseResourceDetailsOutputWithContext(ctx context.Context) OnPremiseResourceDetailsOutput {
	return o
}

func (o OnPremiseResourceDetailsOutput) MachineName() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseResourceDetails) string { return v.MachineName }).(pulumi.StringOutput)
}

func (o OnPremiseResourceDetailsOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseResourceDetails) string { return v.Source }).(pulumi.StringOutput)
}

func (o OnPremiseResourceDetailsOutput) SourceComputerId() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseResourceDetails) string { return v.SourceComputerId }).(pulumi.StringOutput)
}

func (o OnPremiseResourceDetailsOutput) Vmuuid() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseResourceDetails) string { return v.Vmuuid }).(pulumi.StringOutput)
}

func (o OnPremiseResourceDetailsOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseResourceDetails) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

type OnPremiseResourceDetailsResponse struct {
	MachineName      string `pulumi:"machineName"`
	Source           string `pulumi:"source"`
	SourceComputerId string `pulumi:"sourceComputerId"`
	Vmuuid           string `pulumi:"vmuuid"`
	WorkspaceId      string `pulumi:"workspaceId"`
}





type OnPremiseResourceDetailsResponseInput interface {
	pulumi.Input

	ToOnPremiseResourceDetailsResponseOutput() OnPremiseResourceDetailsResponseOutput
	ToOnPremiseResourceDetailsResponseOutputWithContext(context.Context) OnPremiseResourceDetailsResponseOutput
}

type OnPremiseResourceDetailsResponseArgs struct {
	MachineName      pulumi.StringInput `pulumi:"machineName"`
	Source           pulumi.StringInput `pulumi:"source"`
	SourceComputerId pulumi.StringInput `pulumi:"sourceComputerId"`
	Vmuuid           pulumi.StringInput `pulumi:"vmuuid"`
	WorkspaceId      pulumi.StringInput `pulumi:"workspaceId"`
}

func (OnPremiseResourceDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseResourceDetailsResponse)(nil)).Elem()
}

func (i OnPremiseResourceDetailsResponseArgs) ToOnPremiseResourceDetailsResponseOutput() OnPremiseResourceDetailsResponseOutput {
	return i.ToOnPremiseResourceDetailsResponseOutputWithContext(context.Background())
}

func (i OnPremiseResourceDetailsResponseArgs) ToOnPremiseResourceDetailsResponseOutputWithContext(ctx context.Context) OnPremiseResourceDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnPremiseResourceDetailsResponseOutput)
}

type OnPremiseResourceDetailsResponseOutput struct{ *pulumi.OutputState }

func (OnPremiseResourceDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseResourceDetailsResponse)(nil)).Elem()
}

func (o OnPremiseResourceDetailsResponseOutput) ToOnPremiseResourceDetailsResponseOutput() OnPremiseResourceDetailsResponseOutput {
	return o
}

func (o OnPremiseResourceDetailsResponseOutput) ToOnPremiseResourceDetailsResponseOutputWithContext(ctx context.Context) OnPremiseResourceDetailsResponseOutput {
	return o
}

func (o OnPremiseResourceDetailsResponseOutput) MachineName() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseResourceDetailsResponse) string { return v.MachineName }).(pulumi.StringOutput)
}

func (o OnPremiseResourceDetailsResponseOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseResourceDetailsResponse) string { return v.Source }).(pulumi.StringOutput)
}

func (o OnPremiseResourceDetailsResponseOutput) SourceComputerId() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseResourceDetailsResponse) string { return v.SourceComputerId }).(pulumi.StringOutput)
}

func (o OnPremiseResourceDetailsResponseOutput) Vmuuid() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseResourceDetailsResponse) string { return v.Vmuuid }).(pulumi.StringOutput)
}

func (o OnPremiseResourceDetailsResponseOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseResourceDetailsResponse) string { return v.WorkspaceId }).(pulumi.StringOutput)
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





type OnPremiseSqlResourceDetailsInput interface {
	pulumi.Input

	ToOnPremiseSqlResourceDetailsOutput() OnPremiseSqlResourceDetailsOutput
	ToOnPremiseSqlResourceDetailsOutputWithContext(context.Context) OnPremiseSqlResourceDetailsOutput
}

type OnPremiseSqlResourceDetailsArgs struct {
	DatabaseName     pulumi.StringInput `pulumi:"databaseName"`
	MachineName      pulumi.StringInput `pulumi:"machineName"`
	ServerName       pulumi.StringInput `pulumi:"serverName"`
	Source           pulumi.StringInput `pulumi:"source"`
	SourceComputerId pulumi.StringInput `pulumi:"sourceComputerId"`
	Vmuuid           pulumi.StringInput `pulumi:"vmuuid"`
	WorkspaceId      pulumi.StringInput `pulumi:"workspaceId"`
}

func (OnPremiseSqlResourceDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseSqlResourceDetails)(nil)).Elem()
}

func (i OnPremiseSqlResourceDetailsArgs) ToOnPremiseSqlResourceDetailsOutput() OnPremiseSqlResourceDetailsOutput {
	return i.ToOnPremiseSqlResourceDetailsOutputWithContext(context.Background())
}

func (i OnPremiseSqlResourceDetailsArgs) ToOnPremiseSqlResourceDetailsOutputWithContext(ctx context.Context) OnPremiseSqlResourceDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnPremiseSqlResourceDetailsOutput)
}

type OnPremiseSqlResourceDetailsOutput struct{ *pulumi.OutputState }

func (OnPremiseSqlResourceDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseSqlResourceDetails)(nil)).Elem()
}

func (o OnPremiseSqlResourceDetailsOutput) ToOnPremiseSqlResourceDetailsOutput() OnPremiseSqlResourceDetailsOutput {
	return o
}

func (o OnPremiseSqlResourceDetailsOutput) ToOnPremiseSqlResourceDetailsOutputWithContext(ctx context.Context) OnPremiseSqlResourceDetailsOutput {
	return o
}

func (o OnPremiseSqlResourceDetailsOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetails) string { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsOutput) MachineName() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetails) string { return v.MachineName }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetails) string { return v.ServerName }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetails) string { return v.Source }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsOutput) SourceComputerId() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetails) string { return v.SourceComputerId }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsOutput) Vmuuid() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetails) string { return v.Vmuuid }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetails) string { return v.WorkspaceId }).(pulumi.StringOutput)
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





type OnPremiseSqlResourceDetailsResponseInput interface {
	pulumi.Input

	ToOnPremiseSqlResourceDetailsResponseOutput() OnPremiseSqlResourceDetailsResponseOutput
	ToOnPremiseSqlResourceDetailsResponseOutputWithContext(context.Context) OnPremiseSqlResourceDetailsResponseOutput
}

type OnPremiseSqlResourceDetailsResponseArgs struct {
	DatabaseName     pulumi.StringInput `pulumi:"databaseName"`
	MachineName      pulumi.StringInput `pulumi:"machineName"`
	ServerName       pulumi.StringInput `pulumi:"serverName"`
	Source           pulumi.StringInput `pulumi:"source"`
	SourceComputerId pulumi.StringInput `pulumi:"sourceComputerId"`
	Vmuuid           pulumi.StringInput `pulumi:"vmuuid"`
	WorkspaceId      pulumi.StringInput `pulumi:"workspaceId"`
}

func (OnPremiseSqlResourceDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseSqlResourceDetailsResponse)(nil)).Elem()
}

func (i OnPremiseSqlResourceDetailsResponseArgs) ToOnPremiseSqlResourceDetailsResponseOutput() OnPremiseSqlResourceDetailsResponseOutput {
	return i.ToOnPremiseSqlResourceDetailsResponseOutputWithContext(context.Background())
}

func (i OnPremiseSqlResourceDetailsResponseArgs) ToOnPremiseSqlResourceDetailsResponseOutputWithContext(ctx context.Context) OnPremiseSqlResourceDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnPremiseSqlResourceDetailsResponseOutput)
}

type OnPremiseSqlResourceDetailsResponseOutput struct{ *pulumi.OutputState }

func (OnPremiseSqlResourceDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseSqlResourceDetailsResponse)(nil)).Elem()
}

func (o OnPremiseSqlResourceDetailsResponseOutput) ToOnPremiseSqlResourceDetailsResponseOutput() OnPremiseSqlResourceDetailsResponseOutput {
	return o
}

func (o OnPremiseSqlResourceDetailsResponseOutput) ToOnPremiseSqlResourceDetailsResponseOutputWithContext(ctx context.Context) OnPremiseSqlResourceDetailsResponseOutput {
	return o
}

func (o OnPremiseSqlResourceDetailsResponseOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetailsResponse) string { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsResponseOutput) MachineName() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetailsResponse) string { return v.MachineName }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsResponseOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetailsResponse) string { return v.ServerName }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsResponseOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetailsResponse) string { return v.Source }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsResponseOutput) SourceComputerId() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetailsResponse) string { return v.SourceComputerId }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsResponseOutput) Vmuuid() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetailsResponse) string { return v.Vmuuid }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsResponseOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetailsResponse) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

type ScopeElement struct {
	Field *string `pulumi:"field"`
}





type ScopeElementInput interface {
	pulumi.Input

	ToScopeElementOutput() ScopeElementOutput
	ToScopeElementOutputWithContext(context.Context) ScopeElementOutput
}

type ScopeElementArgs struct {
	Field pulumi.StringPtrInput `pulumi:"field"`
}

func (ScopeElementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeElement)(nil)).Elem()
}

func (i ScopeElementArgs) ToScopeElementOutput() ScopeElementOutput {
	return i.ToScopeElementOutputWithContext(context.Background())
}

func (i ScopeElementArgs) ToScopeElementOutputWithContext(ctx context.Context) ScopeElementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeElementOutput)
}





type ScopeElementArrayInput interface {
	pulumi.Input

	ToScopeElementArrayOutput() ScopeElementArrayOutput
	ToScopeElementArrayOutputWithContext(context.Context) ScopeElementArrayOutput
}

type ScopeElementArray []ScopeElementInput

func (ScopeElementArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScopeElement)(nil)).Elem()
}

func (i ScopeElementArray) ToScopeElementArrayOutput() ScopeElementArrayOutput {
	return i.ToScopeElementArrayOutputWithContext(context.Background())
}

func (i ScopeElementArray) ToScopeElementArrayOutputWithContext(ctx context.Context) ScopeElementArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeElementArrayOutput)
}

type ScopeElementOutput struct{ *pulumi.OutputState }

func (ScopeElementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeElement)(nil)).Elem()
}

func (o ScopeElementOutput) ToScopeElementOutput() ScopeElementOutput {
	return o
}

func (o ScopeElementOutput) ToScopeElementOutputWithContext(ctx context.Context) ScopeElementOutput {
	return o
}

func (o ScopeElementOutput) Field() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScopeElement) *string { return v.Field }).(pulumi.StringPtrOutput)
}

type ScopeElementArrayOutput struct{ *pulumi.OutputState }

func (ScopeElementArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScopeElement)(nil)).Elem()
}

func (o ScopeElementArrayOutput) ToScopeElementArrayOutput() ScopeElementArrayOutput {
	return o
}

func (o ScopeElementArrayOutput) ToScopeElementArrayOutputWithContext(ctx context.Context) ScopeElementArrayOutput {
	return o
}

func (o ScopeElementArrayOutput) Index(i pulumi.IntInput) ScopeElementOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScopeElement {
		return vs[0].([]ScopeElement)[vs[1].(int)]
	}).(ScopeElementOutput)
}

type ScopeElementResponse struct {
	Field *string `pulumi:"field"`
}





type ScopeElementResponseInput interface {
	pulumi.Input

	ToScopeElementResponseOutput() ScopeElementResponseOutput
	ToScopeElementResponseOutputWithContext(context.Context) ScopeElementResponseOutput
}

type ScopeElementResponseArgs struct {
	Field pulumi.StringPtrInput `pulumi:"field"`
}

func (ScopeElementResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeElementResponse)(nil)).Elem()
}

func (i ScopeElementResponseArgs) ToScopeElementResponseOutput() ScopeElementResponseOutput {
	return i.ToScopeElementResponseOutputWithContext(context.Background())
}

func (i ScopeElementResponseArgs) ToScopeElementResponseOutputWithContext(ctx context.Context) ScopeElementResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeElementResponseOutput)
}





type ScopeElementResponseArrayInput interface {
	pulumi.Input

	ToScopeElementResponseArrayOutput() ScopeElementResponseArrayOutput
	ToScopeElementResponseArrayOutputWithContext(context.Context) ScopeElementResponseArrayOutput
}

type ScopeElementResponseArray []ScopeElementResponseInput

func (ScopeElementResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScopeElementResponse)(nil)).Elem()
}

func (i ScopeElementResponseArray) ToScopeElementResponseArrayOutput() ScopeElementResponseArrayOutput {
	return i.ToScopeElementResponseArrayOutputWithContext(context.Background())
}

func (i ScopeElementResponseArray) ToScopeElementResponseArrayOutputWithContext(ctx context.Context) ScopeElementResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeElementResponseArrayOutput)
}

type ScopeElementResponseOutput struct{ *pulumi.OutputState }

func (ScopeElementResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeElementResponse)(nil)).Elem()
}

func (o ScopeElementResponseOutput) ToScopeElementResponseOutput() ScopeElementResponseOutput {
	return o
}

func (o ScopeElementResponseOutput) ToScopeElementResponseOutputWithContext(ctx context.Context) ScopeElementResponseOutput {
	return o
}

func (o ScopeElementResponseOutput) Field() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScopeElementResponse) *string { return v.Field }).(pulumi.StringPtrOutput)
}

type ScopeElementResponseArrayOutput struct{ *pulumi.OutputState }

func (ScopeElementResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScopeElementResponse)(nil)).Elem()
}

func (o ScopeElementResponseArrayOutput) ToScopeElementResponseArrayOutput() ScopeElementResponseArrayOutput {
	return o
}

func (o ScopeElementResponseArrayOutput) ToScopeElementResponseArrayOutputWithContext(ctx context.Context) ScopeElementResponseArrayOutput {
	return o
}

func (o ScopeElementResponseArrayOutput) Index(i pulumi.IntInput) ScopeElementResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScopeElementResponse {
		return vs[0].([]ScopeElementResponse)[vs[1].(int)]
	}).(ScopeElementResponseOutput)
}

type SuppressionAlertsScope struct {
	AllOf []ScopeElement `pulumi:"allOf"`
}





type SuppressionAlertsScopeInput interface {
	pulumi.Input

	ToSuppressionAlertsScopeOutput() SuppressionAlertsScopeOutput
	ToSuppressionAlertsScopeOutputWithContext(context.Context) SuppressionAlertsScopeOutput
}

type SuppressionAlertsScopeArgs struct {
	AllOf ScopeElementArrayInput `pulumi:"allOf"`
}

func (SuppressionAlertsScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SuppressionAlertsScope)(nil)).Elem()
}

func (i SuppressionAlertsScopeArgs) ToSuppressionAlertsScopeOutput() SuppressionAlertsScopeOutput {
	return i.ToSuppressionAlertsScopeOutputWithContext(context.Background())
}

func (i SuppressionAlertsScopeArgs) ToSuppressionAlertsScopeOutputWithContext(ctx context.Context) SuppressionAlertsScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionAlertsScopeOutput)
}

func (i SuppressionAlertsScopeArgs) ToSuppressionAlertsScopePtrOutput() SuppressionAlertsScopePtrOutput {
	return i.ToSuppressionAlertsScopePtrOutputWithContext(context.Background())
}

func (i SuppressionAlertsScopeArgs) ToSuppressionAlertsScopePtrOutputWithContext(ctx context.Context) SuppressionAlertsScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionAlertsScopeOutput).ToSuppressionAlertsScopePtrOutputWithContext(ctx)
}









type SuppressionAlertsScopePtrInput interface {
	pulumi.Input

	ToSuppressionAlertsScopePtrOutput() SuppressionAlertsScopePtrOutput
	ToSuppressionAlertsScopePtrOutputWithContext(context.Context) SuppressionAlertsScopePtrOutput
}

type suppressionAlertsScopePtrType SuppressionAlertsScopeArgs

func SuppressionAlertsScopePtr(v *SuppressionAlertsScopeArgs) SuppressionAlertsScopePtrInput {
	return (*suppressionAlertsScopePtrType)(v)
}

func (*suppressionAlertsScopePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SuppressionAlertsScope)(nil)).Elem()
}

func (i *suppressionAlertsScopePtrType) ToSuppressionAlertsScopePtrOutput() SuppressionAlertsScopePtrOutput {
	return i.ToSuppressionAlertsScopePtrOutputWithContext(context.Background())
}

func (i *suppressionAlertsScopePtrType) ToSuppressionAlertsScopePtrOutputWithContext(ctx context.Context) SuppressionAlertsScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionAlertsScopePtrOutput)
}

type SuppressionAlertsScopeOutput struct{ *pulumi.OutputState }

func (SuppressionAlertsScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SuppressionAlertsScope)(nil)).Elem()
}

func (o SuppressionAlertsScopeOutput) ToSuppressionAlertsScopeOutput() SuppressionAlertsScopeOutput {
	return o
}

func (o SuppressionAlertsScopeOutput) ToSuppressionAlertsScopeOutputWithContext(ctx context.Context) SuppressionAlertsScopeOutput {
	return o
}

func (o SuppressionAlertsScopeOutput) ToSuppressionAlertsScopePtrOutput() SuppressionAlertsScopePtrOutput {
	return o.ToSuppressionAlertsScopePtrOutputWithContext(context.Background())
}

func (o SuppressionAlertsScopeOutput) ToSuppressionAlertsScopePtrOutputWithContext(ctx context.Context) SuppressionAlertsScopePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SuppressionAlertsScope) *SuppressionAlertsScope {
		return &v
	}).(SuppressionAlertsScopePtrOutput)
}

func (o SuppressionAlertsScopeOutput) AllOf() ScopeElementArrayOutput {
	return o.ApplyT(func(v SuppressionAlertsScope) []ScopeElement { return v.AllOf }).(ScopeElementArrayOutput)
}

type SuppressionAlertsScopePtrOutput struct{ *pulumi.OutputState }

func (SuppressionAlertsScopePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SuppressionAlertsScope)(nil)).Elem()
}

func (o SuppressionAlertsScopePtrOutput) ToSuppressionAlertsScopePtrOutput() SuppressionAlertsScopePtrOutput {
	return o
}

func (o SuppressionAlertsScopePtrOutput) ToSuppressionAlertsScopePtrOutputWithContext(ctx context.Context) SuppressionAlertsScopePtrOutput {
	return o
}

func (o SuppressionAlertsScopePtrOutput) Elem() SuppressionAlertsScopeOutput {
	return o.ApplyT(func(v *SuppressionAlertsScope) SuppressionAlertsScope {
		if v != nil {
			return *v
		}
		var ret SuppressionAlertsScope
		return ret
	}).(SuppressionAlertsScopeOutput)
}

func (o SuppressionAlertsScopePtrOutput) AllOf() ScopeElementArrayOutput {
	return o.ApplyT(func(v *SuppressionAlertsScope) []ScopeElement {
		if v == nil {
			return nil
		}
		return v.AllOf
	}).(ScopeElementArrayOutput)
}

type SuppressionAlertsScopeResponse struct {
	AllOf []ScopeElementResponse `pulumi:"allOf"`
}





type SuppressionAlertsScopeResponseInput interface {
	pulumi.Input

	ToSuppressionAlertsScopeResponseOutput() SuppressionAlertsScopeResponseOutput
	ToSuppressionAlertsScopeResponseOutputWithContext(context.Context) SuppressionAlertsScopeResponseOutput
}

type SuppressionAlertsScopeResponseArgs struct {
	AllOf ScopeElementResponseArrayInput `pulumi:"allOf"`
}

func (SuppressionAlertsScopeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SuppressionAlertsScopeResponse)(nil)).Elem()
}

func (i SuppressionAlertsScopeResponseArgs) ToSuppressionAlertsScopeResponseOutput() SuppressionAlertsScopeResponseOutput {
	return i.ToSuppressionAlertsScopeResponseOutputWithContext(context.Background())
}

func (i SuppressionAlertsScopeResponseArgs) ToSuppressionAlertsScopeResponseOutputWithContext(ctx context.Context) SuppressionAlertsScopeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionAlertsScopeResponseOutput)
}

func (i SuppressionAlertsScopeResponseArgs) ToSuppressionAlertsScopeResponsePtrOutput() SuppressionAlertsScopeResponsePtrOutput {
	return i.ToSuppressionAlertsScopeResponsePtrOutputWithContext(context.Background())
}

func (i SuppressionAlertsScopeResponseArgs) ToSuppressionAlertsScopeResponsePtrOutputWithContext(ctx context.Context) SuppressionAlertsScopeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionAlertsScopeResponseOutput).ToSuppressionAlertsScopeResponsePtrOutputWithContext(ctx)
}









type SuppressionAlertsScopeResponsePtrInput interface {
	pulumi.Input

	ToSuppressionAlertsScopeResponsePtrOutput() SuppressionAlertsScopeResponsePtrOutput
	ToSuppressionAlertsScopeResponsePtrOutputWithContext(context.Context) SuppressionAlertsScopeResponsePtrOutput
}

type suppressionAlertsScopeResponsePtrType SuppressionAlertsScopeResponseArgs

func SuppressionAlertsScopeResponsePtr(v *SuppressionAlertsScopeResponseArgs) SuppressionAlertsScopeResponsePtrInput {
	return (*suppressionAlertsScopeResponsePtrType)(v)
}

func (*suppressionAlertsScopeResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SuppressionAlertsScopeResponse)(nil)).Elem()
}

func (i *suppressionAlertsScopeResponsePtrType) ToSuppressionAlertsScopeResponsePtrOutput() SuppressionAlertsScopeResponsePtrOutput {
	return i.ToSuppressionAlertsScopeResponsePtrOutputWithContext(context.Background())
}

func (i *suppressionAlertsScopeResponsePtrType) ToSuppressionAlertsScopeResponsePtrOutputWithContext(ctx context.Context) SuppressionAlertsScopeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionAlertsScopeResponsePtrOutput)
}

type SuppressionAlertsScopeResponseOutput struct{ *pulumi.OutputState }

func (SuppressionAlertsScopeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SuppressionAlertsScopeResponse)(nil)).Elem()
}

func (o SuppressionAlertsScopeResponseOutput) ToSuppressionAlertsScopeResponseOutput() SuppressionAlertsScopeResponseOutput {
	return o
}

func (o SuppressionAlertsScopeResponseOutput) ToSuppressionAlertsScopeResponseOutputWithContext(ctx context.Context) SuppressionAlertsScopeResponseOutput {
	return o
}

func (o SuppressionAlertsScopeResponseOutput) ToSuppressionAlertsScopeResponsePtrOutput() SuppressionAlertsScopeResponsePtrOutput {
	return o.ToSuppressionAlertsScopeResponsePtrOutputWithContext(context.Background())
}

func (o SuppressionAlertsScopeResponseOutput) ToSuppressionAlertsScopeResponsePtrOutputWithContext(ctx context.Context) SuppressionAlertsScopeResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SuppressionAlertsScopeResponse) *SuppressionAlertsScopeResponse {
		return &v
	}).(SuppressionAlertsScopeResponsePtrOutput)
}

func (o SuppressionAlertsScopeResponseOutput) AllOf() ScopeElementResponseArrayOutput {
	return o.ApplyT(func(v SuppressionAlertsScopeResponse) []ScopeElementResponse { return v.AllOf }).(ScopeElementResponseArrayOutput)
}

type SuppressionAlertsScopeResponsePtrOutput struct{ *pulumi.OutputState }

func (SuppressionAlertsScopeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SuppressionAlertsScopeResponse)(nil)).Elem()
}

func (o SuppressionAlertsScopeResponsePtrOutput) ToSuppressionAlertsScopeResponsePtrOutput() SuppressionAlertsScopeResponsePtrOutput {
	return o
}

func (o SuppressionAlertsScopeResponsePtrOutput) ToSuppressionAlertsScopeResponsePtrOutputWithContext(ctx context.Context) SuppressionAlertsScopeResponsePtrOutput {
	return o
}

func (o SuppressionAlertsScopeResponsePtrOutput) Elem() SuppressionAlertsScopeResponseOutput {
	return o.ApplyT(func(v *SuppressionAlertsScopeResponse) SuppressionAlertsScopeResponse {
		if v != nil {
			return *v
		}
		var ret SuppressionAlertsScopeResponse
		return ret
	}).(SuppressionAlertsScopeResponseOutput)
}

func (o SuppressionAlertsScopeResponsePtrOutput) AllOf() ScopeElementResponseArrayOutput {
	return o.ApplyT(func(v *SuppressionAlertsScopeResponse) []ScopeElementResponse {
		if v == nil {
			return nil
		}
		return v.AllOf
	}).(ScopeElementResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(AssessmentLinksResponseOutput{})
	pulumi.RegisterOutputType(AssessmentLinksResponsePtrOutput{})
	pulumi.RegisterOutputType(AssessmentStatusOutput{})
	pulumi.RegisterOutputType(AssessmentStatusPtrOutput{})
	pulumi.RegisterOutputType(AssessmentStatusResponseOutput{})
	pulumi.RegisterOutputType(AssessmentStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(AutomationActionEventHubOutput{})
	pulumi.RegisterOutputType(AutomationActionEventHubResponseOutput{})
	pulumi.RegisterOutputType(AutomationActionLogicAppOutput{})
	pulumi.RegisterOutputType(AutomationActionLogicAppResponseOutput{})
	pulumi.RegisterOutputType(AutomationActionWorkspaceOutput{})
	pulumi.RegisterOutputType(AutomationActionWorkspaceResponseOutput{})
	pulumi.RegisterOutputType(AutomationRuleSetOutput{})
	pulumi.RegisterOutputType(AutomationRuleSetArrayOutput{})
	pulumi.RegisterOutputType(AutomationRuleSetResponseOutput{})
	pulumi.RegisterOutputType(AutomationRuleSetResponseArrayOutput{})
	pulumi.RegisterOutputType(AutomationScopeOutput{})
	pulumi.RegisterOutputType(AutomationScopeArrayOutput{})
	pulumi.RegisterOutputType(AutomationScopeResponseOutput{})
	pulumi.RegisterOutputType(AutomationScopeResponseArrayOutput{})
	pulumi.RegisterOutputType(AutomationSourceOutput{})
	pulumi.RegisterOutputType(AutomationSourceArrayOutput{})
	pulumi.RegisterOutputType(AutomationSourceResponseOutput{})
	pulumi.RegisterOutputType(AutomationSourceResponseArrayOutput{})
	pulumi.RegisterOutputType(AutomationTriggeringRuleOutput{})
	pulumi.RegisterOutputType(AutomationTriggeringRuleArrayOutput{})
	pulumi.RegisterOutputType(AutomationTriggeringRuleResponseOutput{})
	pulumi.RegisterOutputType(AutomationTriggeringRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(AzureResourceDetailsOutput{})
	pulumi.RegisterOutputType(AzureResourceDetailsResponseOutput{})
	pulumi.RegisterOutputType(OnPremiseResourceDetailsOutput{})
	pulumi.RegisterOutputType(OnPremiseResourceDetailsResponseOutput{})
	pulumi.RegisterOutputType(OnPremiseSqlResourceDetailsOutput{})
	pulumi.RegisterOutputType(OnPremiseSqlResourceDetailsResponseOutput{})
	pulumi.RegisterOutputType(ScopeElementOutput{})
	pulumi.RegisterOutputType(ScopeElementArrayOutput{})
	pulumi.RegisterOutputType(ScopeElementResponseOutput{})
	pulumi.RegisterOutputType(ScopeElementResponseArrayOutput{})
	pulumi.RegisterOutputType(SuppressionAlertsScopeOutput{})
	pulumi.RegisterOutputType(SuppressionAlertsScopePtrOutput{})
	pulumi.RegisterOutputType(SuppressionAlertsScopeResponseOutput{})
	pulumi.RegisterOutputType(SuppressionAlertsScopeResponsePtrOutput{})
}

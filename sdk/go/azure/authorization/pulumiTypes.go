


package authorization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessReviewHistoryInstance struct {
	DisplayName                      *string `pulumi:"displayName"`
	Expiration                       *string `pulumi:"expiration"`
	FulfilledDateTime                *string `pulumi:"fulfilledDateTime"`
	ReviewHistoryPeriodEndDateTime   *string `pulumi:"reviewHistoryPeriodEndDateTime"`
	ReviewHistoryPeriodStartDateTime *string `pulumi:"reviewHistoryPeriodStartDateTime"`
	RunDateTime                      *string `pulumi:"runDateTime"`
}





type AccessReviewHistoryInstanceInput interface {
	pulumi.Input

	ToAccessReviewHistoryInstanceOutput() AccessReviewHistoryInstanceOutput
	ToAccessReviewHistoryInstanceOutputWithContext(context.Context) AccessReviewHistoryInstanceOutput
}

type AccessReviewHistoryInstanceArgs struct {
	DisplayName                      pulumi.StringPtrInput `pulumi:"displayName"`
	Expiration                       pulumi.StringPtrInput `pulumi:"expiration"`
	FulfilledDateTime                pulumi.StringPtrInput `pulumi:"fulfilledDateTime"`
	ReviewHistoryPeriodEndDateTime   pulumi.StringPtrInput `pulumi:"reviewHistoryPeriodEndDateTime"`
	ReviewHistoryPeriodStartDateTime pulumi.StringPtrInput `pulumi:"reviewHistoryPeriodStartDateTime"`
	RunDateTime                      pulumi.StringPtrInput `pulumi:"runDateTime"`
}

func (AccessReviewHistoryInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewHistoryInstance)(nil)).Elem()
}

func (i AccessReviewHistoryInstanceArgs) ToAccessReviewHistoryInstanceOutput() AccessReviewHistoryInstanceOutput {
	return i.ToAccessReviewHistoryInstanceOutputWithContext(context.Background())
}

func (i AccessReviewHistoryInstanceArgs) ToAccessReviewHistoryInstanceOutputWithContext(ctx context.Context) AccessReviewHistoryInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessReviewHistoryInstanceOutput)
}





type AccessReviewHistoryInstanceArrayInput interface {
	pulumi.Input

	ToAccessReviewHistoryInstanceArrayOutput() AccessReviewHistoryInstanceArrayOutput
	ToAccessReviewHistoryInstanceArrayOutputWithContext(context.Context) AccessReviewHistoryInstanceArrayOutput
}

type AccessReviewHistoryInstanceArray []AccessReviewHistoryInstanceInput

func (AccessReviewHistoryInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessReviewHistoryInstance)(nil)).Elem()
}

func (i AccessReviewHistoryInstanceArray) ToAccessReviewHistoryInstanceArrayOutput() AccessReviewHistoryInstanceArrayOutput {
	return i.ToAccessReviewHistoryInstanceArrayOutputWithContext(context.Background())
}

func (i AccessReviewHistoryInstanceArray) ToAccessReviewHistoryInstanceArrayOutputWithContext(ctx context.Context) AccessReviewHistoryInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessReviewHistoryInstanceArrayOutput)
}

type AccessReviewHistoryInstanceOutput struct{ *pulumi.OutputState }

func (AccessReviewHistoryInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewHistoryInstance)(nil)).Elem()
}

func (o AccessReviewHistoryInstanceOutput) ToAccessReviewHistoryInstanceOutput() AccessReviewHistoryInstanceOutput {
	return o
}

func (o AccessReviewHistoryInstanceOutput) ToAccessReviewHistoryInstanceOutputWithContext(ctx context.Context) AccessReviewHistoryInstanceOutput {
	return o
}

func (o AccessReviewHistoryInstanceOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewHistoryInstance) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o AccessReviewHistoryInstanceOutput) Expiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewHistoryInstance) *string { return v.Expiration }).(pulumi.StringPtrOutput)
}

func (o AccessReviewHistoryInstanceOutput) FulfilledDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewHistoryInstance) *string { return v.FulfilledDateTime }).(pulumi.StringPtrOutput)
}

func (o AccessReviewHistoryInstanceOutput) ReviewHistoryPeriodEndDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewHistoryInstance) *string { return v.ReviewHistoryPeriodEndDateTime }).(pulumi.StringPtrOutput)
}

func (o AccessReviewHistoryInstanceOutput) ReviewHistoryPeriodStartDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewHistoryInstance) *string { return v.ReviewHistoryPeriodStartDateTime }).(pulumi.StringPtrOutput)
}

func (o AccessReviewHistoryInstanceOutput) RunDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewHistoryInstance) *string { return v.RunDateTime }).(pulumi.StringPtrOutput)
}

type AccessReviewHistoryInstanceArrayOutput struct{ *pulumi.OutputState }

func (AccessReviewHistoryInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessReviewHistoryInstance)(nil)).Elem()
}

func (o AccessReviewHistoryInstanceArrayOutput) ToAccessReviewHistoryInstanceArrayOutput() AccessReviewHistoryInstanceArrayOutput {
	return o
}

func (o AccessReviewHistoryInstanceArrayOutput) ToAccessReviewHistoryInstanceArrayOutputWithContext(ctx context.Context) AccessReviewHistoryInstanceArrayOutput {
	return o
}

func (o AccessReviewHistoryInstanceArrayOutput) Index(i pulumi.IntInput) AccessReviewHistoryInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessReviewHistoryInstance {
		return vs[0].([]AccessReviewHistoryInstance)[vs[1].(int)]
	}).(AccessReviewHistoryInstanceOutput)
}

type AccessReviewHistoryInstanceResponse struct {
	DisplayName                      *string `pulumi:"displayName"`
	DownloadUri                      string  `pulumi:"downloadUri"`
	Expiration                       *string `pulumi:"expiration"`
	FulfilledDateTime                *string `pulumi:"fulfilledDateTime"`
	Id                               string  `pulumi:"id"`
	Name                             string  `pulumi:"name"`
	ReviewHistoryPeriodEndDateTime   *string `pulumi:"reviewHistoryPeriodEndDateTime"`
	ReviewHistoryPeriodStartDateTime *string `pulumi:"reviewHistoryPeriodStartDateTime"`
	RunDateTime                      *string `pulumi:"runDateTime"`
	Status                           string  `pulumi:"status"`
	Type                             string  `pulumi:"type"`
}

type AccessReviewHistoryInstanceResponseOutput struct{ *pulumi.OutputState }

func (AccessReviewHistoryInstanceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewHistoryInstanceResponse)(nil)).Elem()
}

func (o AccessReviewHistoryInstanceResponseOutput) ToAccessReviewHistoryInstanceResponseOutput() AccessReviewHistoryInstanceResponseOutput {
	return o
}

func (o AccessReviewHistoryInstanceResponseOutput) ToAccessReviewHistoryInstanceResponseOutputWithContext(ctx context.Context) AccessReviewHistoryInstanceResponseOutput {
	return o
}

func (o AccessReviewHistoryInstanceResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewHistoryInstanceResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o AccessReviewHistoryInstanceResponseOutput) DownloadUri() pulumi.StringOutput {
	return o.ApplyT(func(v AccessReviewHistoryInstanceResponse) string { return v.DownloadUri }).(pulumi.StringOutput)
}

func (o AccessReviewHistoryInstanceResponseOutput) Expiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewHistoryInstanceResponse) *string { return v.Expiration }).(pulumi.StringPtrOutput)
}

func (o AccessReviewHistoryInstanceResponseOutput) FulfilledDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewHistoryInstanceResponse) *string { return v.FulfilledDateTime }).(pulumi.StringPtrOutput)
}

func (o AccessReviewHistoryInstanceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AccessReviewHistoryInstanceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o AccessReviewHistoryInstanceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AccessReviewHistoryInstanceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o AccessReviewHistoryInstanceResponseOutput) ReviewHistoryPeriodEndDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewHistoryInstanceResponse) *string { return v.ReviewHistoryPeriodEndDateTime }).(pulumi.StringPtrOutput)
}

func (o AccessReviewHistoryInstanceResponseOutput) ReviewHistoryPeriodStartDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewHistoryInstanceResponse) *string { return v.ReviewHistoryPeriodStartDateTime }).(pulumi.StringPtrOutput)
}

func (o AccessReviewHistoryInstanceResponseOutput) RunDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewHistoryInstanceResponse) *string { return v.RunDateTime }).(pulumi.StringPtrOutput)
}

func (o AccessReviewHistoryInstanceResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v AccessReviewHistoryInstanceResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o AccessReviewHistoryInstanceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AccessReviewHistoryInstanceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type AccessReviewHistoryInstanceResponseArrayOutput struct{ *pulumi.OutputState }

func (AccessReviewHistoryInstanceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessReviewHistoryInstanceResponse)(nil)).Elem()
}

func (o AccessReviewHistoryInstanceResponseArrayOutput) ToAccessReviewHistoryInstanceResponseArrayOutput() AccessReviewHistoryInstanceResponseArrayOutput {
	return o
}

func (o AccessReviewHistoryInstanceResponseArrayOutput) ToAccessReviewHistoryInstanceResponseArrayOutputWithContext(ctx context.Context) AccessReviewHistoryInstanceResponseArrayOutput {
	return o
}

func (o AccessReviewHistoryInstanceResponseArrayOutput) Index(i pulumi.IntInput) AccessReviewHistoryInstanceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessReviewHistoryInstanceResponse {
		return vs[0].([]AccessReviewHistoryInstanceResponse)[vs[1].(int)]
	}).(AccessReviewHistoryInstanceResponseOutput)
}

type AccessReviewInstance struct {
	BackupReviewers []AccessReviewReviewer `pulumi:"backupReviewers"`
	EndDateTime     *string                `pulumi:"endDateTime"`
	Reviewers       []AccessReviewReviewer `pulumi:"reviewers"`
	StartDateTime   *string                `pulumi:"startDateTime"`
}





type AccessReviewInstanceInput interface {
	pulumi.Input

	ToAccessReviewInstanceOutput() AccessReviewInstanceOutput
	ToAccessReviewInstanceOutputWithContext(context.Context) AccessReviewInstanceOutput
}

type AccessReviewInstanceArgs struct {
	BackupReviewers AccessReviewReviewerArrayInput `pulumi:"backupReviewers"`
	EndDateTime     pulumi.StringPtrInput          `pulumi:"endDateTime"`
	Reviewers       AccessReviewReviewerArrayInput `pulumi:"reviewers"`
	StartDateTime   pulumi.StringPtrInput          `pulumi:"startDateTime"`
}

func (AccessReviewInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewInstance)(nil)).Elem()
}

func (i AccessReviewInstanceArgs) ToAccessReviewInstanceOutput() AccessReviewInstanceOutput {
	return i.ToAccessReviewInstanceOutputWithContext(context.Background())
}

func (i AccessReviewInstanceArgs) ToAccessReviewInstanceOutputWithContext(ctx context.Context) AccessReviewInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessReviewInstanceOutput)
}





type AccessReviewInstanceArrayInput interface {
	pulumi.Input

	ToAccessReviewInstanceArrayOutput() AccessReviewInstanceArrayOutput
	ToAccessReviewInstanceArrayOutputWithContext(context.Context) AccessReviewInstanceArrayOutput
}

type AccessReviewInstanceArray []AccessReviewInstanceInput

func (AccessReviewInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessReviewInstance)(nil)).Elem()
}

func (i AccessReviewInstanceArray) ToAccessReviewInstanceArrayOutput() AccessReviewInstanceArrayOutput {
	return i.ToAccessReviewInstanceArrayOutputWithContext(context.Background())
}

func (i AccessReviewInstanceArray) ToAccessReviewInstanceArrayOutputWithContext(ctx context.Context) AccessReviewInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessReviewInstanceArrayOutput)
}

type AccessReviewInstanceOutput struct{ *pulumi.OutputState }

func (AccessReviewInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewInstance)(nil)).Elem()
}

func (o AccessReviewInstanceOutput) ToAccessReviewInstanceOutput() AccessReviewInstanceOutput {
	return o
}

func (o AccessReviewInstanceOutput) ToAccessReviewInstanceOutputWithContext(ctx context.Context) AccessReviewInstanceOutput {
	return o
}

func (o AccessReviewInstanceOutput) BackupReviewers() AccessReviewReviewerArrayOutput {
	return o.ApplyT(func(v AccessReviewInstance) []AccessReviewReviewer { return v.BackupReviewers }).(AccessReviewReviewerArrayOutput)
}

func (o AccessReviewInstanceOutput) EndDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewInstance) *string { return v.EndDateTime }).(pulumi.StringPtrOutput)
}

func (o AccessReviewInstanceOutput) Reviewers() AccessReviewReviewerArrayOutput {
	return o.ApplyT(func(v AccessReviewInstance) []AccessReviewReviewer { return v.Reviewers }).(AccessReviewReviewerArrayOutput)
}

func (o AccessReviewInstanceOutput) StartDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewInstance) *string { return v.StartDateTime }).(pulumi.StringPtrOutput)
}

type AccessReviewInstanceArrayOutput struct{ *pulumi.OutputState }

func (AccessReviewInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessReviewInstance)(nil)).Elem()
}

func (o AccessReviewInstanceArrayOutput) ToAccessReviewInstanceArrayOutput() AccessReviewInstanceArrayOutput {
	return o
}

func (o AccessReviewInstanceArrayOutput) ToAccessReviewInstanceArrayOutputWithContext(ctx context.Context) AccessReviewInstanceArrayOutput {
	return o
}

func (o AccessReviewInstanceArrayOutput) Index(i pulumi.IntInput) AccessReviewInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessReviewInstance {
		return vs[0].([]AccessReviewInstance)[vs[1].(int)]
	}).(AccessReviewInstanceOutput)
}

type AccessReviewInstanceResponse struct {
	BackupReviewers []AccessReviewReviewerResponse `pulumi:"backupReviewers"`
	EndDateTime     *string                        `pulumi:"endDateTime"`
	Id              string                         `pulumi:"id"`
	Name            string                         `pulumi:"name"`
	Reviewers       []AccessReviewReviewerResponse `pulumi:"reviewers"`
	ReviewersType   string                         `pulumi:"reviewersType"`
	StartDateTime   *string                        `pulumi:"startDateTime"`
	Status          string                         `pulumi:"status"`
	Type            string                         `pulumi:"type"`
}

type AccessReviewInstanceResponseOutput struct{ *pulumi.OutputState }

func (AccessReviewInstanceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewInstanceResponse)(nil)).Elem()
}

func (o AccessReviewInstanceResponseOutput) ToAccessReviewInstanceResponseOutput() AccessReviewInstanceResponseOutput {
	return o
}

func (o AccessReviewInstanceResponseOutput) ToAccessReviewInstanceResponseOutputWithContext(ctx context.Context) AccessReviewInstanceResponseOutput {
	return o
}

func (o AccessReviewInstanceResponseOutput) BackupReviewers() AccessReviewReviewerResponseArrayOutput {
	return o.ApplyT(func(v AccessReviewInstanceResponse) []AccessReviewReviewerResponse { return v.BackupReviewers }).(AccessReviewReviewerResponseArrayOutput)
}

func (o AccessReviewInstanceResponseOutput) EndDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewInstanceResponse) *string { return v.EndDateTime }).(pulumi.StringPtrOutput)
}

func (o AccessReviewInstanceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AccessReviewInstanceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o AccessReviewInstanceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AccessReviewInstanceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o AccessReviewInstanceResponseOutput) Reviewers() AccessReviewReviewerResponseArrayOutput {
	return o.ApplyT(func(v AccessReviewInstanceResponse) []AccessReviewReviewerResponse { return v.Reviewers }).(AccessReviewReviewerResponseArrayOutput)
}

func (o AccessReviewInstanceResponseOutput) ReviewersType() pulumi.StringOutput {
	return o.ApplyT(func(v AccessReviewInstanceResponse) string { return v.ReviewersType }).(pulumi.StringOutput)
}

func (o AccessReviewInstanceResponseOutput) StartDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewInstanceResponse) *string { return v.StartDateTime }).(pulumi.StringPtrOutput)
}

func (o AccessReviewInstanceResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v AccessReviewInstanceResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o AccessReviewInstanceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AccessReviewInstanceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type AccessReviewInstanceResponseArrayOutput struct{ *pulumi.OutputState }

func (AccessReviewInstanceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessReviewInstanceResponse)(nil)).Elem()
}

func (o AccessReviewInstanceResponseArrayOutput) ToAccessReviewInstanceResponseArrayOutput() AccessReviewInstanceResponseArrayOutput {
	return o
}

func (o AccessReviewInstanceResponseArrayOutput) ToAccessReviewInstanceResponseArrayOutputWithContext(ctx context.Context) AccessReviewInstanceResponseArrayOutput {
	return o
}

func (o AccessReviewInstanceResponseArrayOutput) Index(i pulumi.IntInput) AccessReviewInstanceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessReviewInstanceResponse {
		return vs[0].([]AccessReviewInstanceResponse)[vs[1].(int)]
	}).(AccessReviewInstanceResponseOutput)
}

type AccessReviewReviewer struct {
	PrincipalId *string `pulumi:"principalId"`
}





type AccessReviewReviewerInput interface {
	pulumi.Input

	ToAccessReviewReviewerOutput() AccessReviewReviewerOutput
	ToAccessReviewReviewerOutputWithContext(context.Context) AccessReviewReviewerOutput
}

type AccessReviewReviewerArgs struct {
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
}

func (AccessReviewReviewerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewReviewer)(nil)).Elem()
}

func (i AccessReviewReviewerArgs) ToAccessReviewReviewerOutput() AccessReviewReviewerOutput {
	return i.ToAccessReviewReviewerOutputWithContext(context.Background())
}

func (i AccessReviewReviewerArgs) ToAccessReviewReviewerOutputWithContext(ctx context.Context) AccessReviewReviewerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessReviewReviewerOutput)
}





type AccessReviewReviewerArrayInput interface {
	pulumi.Input

	ToAccessReviewReviewerArrayOutput() AccessReviewReviewerArrayOutput
	ToAccessReviewReviewerArrayOutputWithContext(context.Context) AccessReviewReviewerArrayOutput
}

type AccessReviewReviewerArray []AccessReviewReviewerInput

func (AccessReviewReviewerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessReviewReviewer)(nil)).Elem()
}

func (i AccessReviewReviewerArray) ToAccessReviewReviewerArrayOutput() AccessReviewReviewerArrayOutput {
	return i.ToAccessReviewReviewerArrayOutputWithContext(context.Background())
}

func (i AccessReviewReviewerArray) ToAccessReviewReviewerArrayOutputWithContext(ctx context.Context) AccessReviewReviewerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessReviewReviewerArrayOutput)
}

type AccessReviewReviewerOutput struct{ *pulumi.OutputState }

func (AccessReviewReviewerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewReviewer)(nil)).Elem()
}

func (o AccessReviewReviewerOutput) ToAccessReviewReviewerOutput() AccessReviewReviewerOutput {
	return o
}

func (o AccessReviewReviewerOutput) ToAccessReviewReviewerOutputWithContext(ctx context.Context) AccessReviewReviewerOutput {
	return o
}

func (o AccessReviewReviewerOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewReviewer) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

type AccessReviewReviewerArrayOutput struct{ *pulumi.OutputState }

func (AccessReviewReviewerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessReviewReviewer)(nil)).Elem()
}

func (o AccessReviewReviewerArrayOutput) ToAccessReviewReviewerArrayOutput() AccessReviewReviewerArrayOutput {
	return o
}

func (o AccessReviewReviewerArrayOutput) ToAccessReviewReviewerArrayOutputWithContext(ctx context.Context) AccessReviewReviewerArrayOutput {
	return o
}

func (o AccessReviewReviewerArrayOutput) Index(i pulumi.IntInput) AccessReviewReviewerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessReviewReviewer {
		return vs[0].([]AccessReviewReviewer)[vs[1].(int)]
	}).(AccessReviewReviewerOutput)
}

type AccessReviewReviewerResponse struct {
	PrincipalId   *string `pulumi:"principalId"`
	PrincipalType string  `pulumi:"principalType"`
}

type AccessReviewReviewerResponseOutput struct{ *pulumi.OutputState }

func (AccessReviewReviewerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewReviewerResponse)(nil)).Elem()
}

func (o AccessReviewReviewerResponseOutput) ToAccessReviewReviewerResponseOutput() AccessReviewReviewerResponseOutput {
	return o
}

func (o AccessReviewReviewerResponseOutput) ToAccessReviewReviewerResponseOutputWithContext(ctx context.Context) AccessReviewReviewerResponseOutput {
	return o
}

func (o AccessReviewReviewerResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewReviewerResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o AccessReviewReviewerResponseOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v AccessReviewReviewerResponse) string { return v.PrincipalType }).(pulumi.StringOutput)
}

type AccessReviewReviewerResponseArrayOutput struct{ *pulumi.OutputState }

func (AccessReviewReviewerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessReviewReviewerResponse)(nil)).Elem()
}

func (o AccessReviewReviewerResponseArrayOutput) ToAccessReviewReviewerResponseArrayOutput() AccessReviewReviewerResponseArrayOutput {
	return o
}

func (o AccessReviewReviewerResponseArrayOutput) ToAccessReviewReviewerResponseArrayOutputWithContext(ctx context.Context) AccessReviewReviewerResponseArrayOutput {
	return o
}

func (o AccessReviewReviewerResponseArrayOutput) Index(i pulumi.IntInput) AccessReviewReviewerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessReviewReviewerResponse {
		return vs[0].([]AccessReviewReviewerResponse)[vs[1].(int)]
	}).(AccessReviewReviewerResponseOutput)
}

type AccessReviewScope struct {
	ExcludeResourceId          *string `pulumi:"excludeResourceId"`
	ExcludeRoleDefinitionId    *string `pulumi:"excludeRoleDefinitionId"`
	ExpandNestedMemberships    *bool   `pulumi:"expandNestedMemberships"`
	InactiveDuration           *string `pulumi:"inactiveDuration"`
	IncludeAccessBelowResource *bool   `pulumi:"includeAccessBelowResource"`
	IncludeInheritedAccess     *bool   `pulumi:"includeInheritedAccess"`
}





type AccessReviewScopeInput interface {
	pulumi.Input

	ToAccessReviewScopeOutput() AccessReviewScopeOutput
	ToAccessReviewScopeOutputWithContext(context.Context) AccessReviewScopeOutput
}

type AccessReviewScopeArgs struct {
	ExcludeResourceId          pulumi.StringPtrInput `pulumi:"excludeResourceId"`
	ExcludeRoleDefinitionId    pulumi.StringPtrInput `pulumi:"excludeRoleDefinitionId"`
	ExpandNestedMemberships    pulumi.BoolPtrInput   `pulumi:"expandNestedMemberships"`
	InactiveDuration           pulumi.StringPtrInput `pulumi:"inactiveDuration"`
	IncludeAccessBelowResource pulumi.BoolPtrInput   `pulumi:"includeAccessBelowResource"`
	IncludeInheritedAccess     pulumi.BoolPtrInput   `pulumi:"includeInheritedAccess"`
}

func (AccessReviewScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewScope)(nil)).Elem()
}

func (i AccessReviewScopeArgs) ToAccessReviewScopeOutput() AccessReviewScopeOutput {
	return i.ToAccessReviewScopeOutputWithContext(context.Background())
}

func (i AccessReviewScopeArgs) ToAccessReviewScopeOutputWithContext(ctx context.Context) AccessReviewScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessReviewScopeOutput)
}





type AccessReviewScopeArrayInput interface {
	pulumi.Input

	ToAccessReviewScopeArrayOutput() AccessReviewScopeArrayOutput
	ToAccessReviewScopeArrayOutputWithContext(context.Context) AccessReviewScopeArrayOutput
}

type AccessReviewScopeArray []AccessReviewScopeInput

func (AccessReviewScopeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessReviewScope)(nil)).Elem()
}

func (i AccessReviewScopeArray) ToAccessReviewScopeArrayOutput() AccessReviewScopeArrayOutput {
	return i.ToAccessReviewScopeArrayOutputWithContext(context.Background())
}

func (i AccessReviewScopeArray) ToAccessReviewScopeArrayOutputWithContext(ctx context.Context) AccessReviewScopeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessReviewScopeArrayOutput)
}

type AccessReviewScopeOutput struct{ *pulumi.OutputState }

func (AccessReviewScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewScope)(nil)).Elem()
}

func (o AccessReviewScopeOutput) ToAccessReviewScopeOutput() AccessReviewScopeOutput {
	return o
}

func (o AccessReviewScopeOutput) ToAccessReviewScopeOutputWithContext(ctx context.Context) AccessReviewScopeOutput {
	return o
}

func (o AccessReviewScopeOutput) ExcludeResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewScope) *string { return v.ExcludeResourceId }).(pulumi.StringPtrOutput)
}

func (o AccessReviewScopeOutput) ExcludeRoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewScope) *string { return v.ExcludeRoleDefinitionId }).(pulumi.StringPtrOutput)
}

func (o AccessReviewScopeOutput) ExpandNestedMemberships() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AccessReviewScope) *bool { return v.ExpandNestedMemberships }).(pulumi.BoolPtrOutput)
}

func (o AccessReviewScopeOutput) InactiveDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewScope) *string { return v.InactiveDuration }).(pulumi.StringPtrOutput)
}

func (o AccessReviewScopeOutput) IncludeAccessBelowResource() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AccessReviewScope) *bool { return v.IncludeAccessBelowResource }).(pulumi.BoolPtrOutput)
}

func (o AccessReviewScopeOutput) IncludeInheritedAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AccessReviewScope) *bool { return v.IncludeInheritedAccess }).(pulumi.BoolPtrOutput)
}

type AccessReviewScopeArrayOutput struct{ *pulumi.OutputState }

func (AccessReviewScopeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessReviewScope)(nil)).Elem()
}

func (o AccessReviewScopeArrayOutput) ToAccessReviewScopeArrayOutput() AccessReviewScopeArrayOutput {
	return o
}

func (o AccessReviewScopeArrayOutput) ToAccessReviewScopeArrayOutputWithContext(ctx context.Context) AccessReviewScopeArrayOutput {
	return o
}

func (o AccessReviewScopeArrayOutput) Index(i pulumi.IntInput) AccessReviewScopeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessReviewScope {
		return vs[0].([]AccessReviewScope)[vs[1].(int)]
	}).(AccessReviewScopeOutput)
}

type AccessReviewScopeResponse struct {
	AssignmentState            string  `pulumi:"assignmentState"`
	ExcludeResourceId          *string `pulumi:"excludeResourceId"`
	ExcludeRoleDefinitionId    *string `pulumi:"excludeRoleDefinitionId"`
	ExpandNestedMemberships    *bool   `pulumi:"expandNestedMemberships"`
	InactiveDuration           *string `pulumi:"inactiveDuration"`
	IncludeAccessBelowResource *bool   `pulumi:"includeAccessBelowResource"`
	IncludeInheritedAccess     *bool   `pulumi:"includeInheritedAccess"`
	PrincipalType              string  `pulumi:"principalType"`
	ResourceId                 string  `pulumi:"resourceId"`
	RoleDefinitionId           string  `pulumi:"roleDefinitionId"`
}

type AccessReviewScopeResponseOutput struct{ *pulumi.OutputState }

func (AccessReviewScopeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewScopeResponse)(nil)).Elem()
}

func (o AccessReviewScopeResponseOutput) ToAccessReviewScopeResponseOutput() AccessReviewScopeResponseOutput {
	return o
}

func (o AccessReviewScopeResponseOutput) ToAccessReviewScopeResponseOutputWithContext(ctx context.Context) AccessReviewScopeResponseOutput {
	return o
}

func (o AccessReviewScopeResponseOutput) AssignmentState() pulumi.StringOutput {
	return o.ApplyT(func(v AccessReviewScopeResponse) string { return v.AssignmentState }).(pulumi.StringOutput)
}

func (o AccessReviewScopeResponseOutput) ExcludeResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewScopeResponse) *string { return v.ExcludeResourceId }).(pulumi.StringPtrOutput)
}

func (o AccessReviewScopeResponseOutput) ExcludeRoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewScopeResponse) *string { return v.ExcludeRoleDefinitionId }).(pulumi.StringPtrOutput)
}

func (o AccessReviewScopeResponseOutput) ExpandNestedMemberships() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AccessReviewScopeResponse) *bool { return v.ExpandNestedMemberships }).(pulumi.BoolPtrOutput)
}

func (o AccessReviewScopeResponseOutput) InactiveDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewScopeResponse) *string { return v.InactiveDuration }).(pulumi.StringPtrOutput)
}

func (o AccessReviewScopeResponseOutput) IncludeAccessBelowResource() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AccessReviewScopeResponse) *bool { return v.IncludeAccessBelowResource }).(pulumi.BoolPtrOutput)
}

func (o AccessReviewScopeResponseOutput) IncludeInheritedAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AccessReviewScopeResponse) *bool { return v.IncludeInheritedAccess }).(pulumi.BoolPtrOutput)
}

func (o AccessReviewScopeResponseOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v AccessReviewScopeResponse) string { return v.PrincipalType }).(pulumi.StringOutput)
}

func (o AccessReviewScopeResponseOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v AccessReviewScopeResponse) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o AccessReviewScopeResponseOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v AccessReviewScopeResponse) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type AccessReviewScopeResponseArrayOutput struct{ *pulumi.OutputState }

func (AccessReviewScopeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessReviewScopeResponse)(nil)).Elem()
}

func (o AccessReviewScopeResponseArrayOutput) ToAccessReviewScopeResponseArrayOutput() AccessReviewScopeResponseArrayOutput {
	return o
}

func (o AccessReviewScopeResponseArrayOutput) ToAccessReviewScopeResponseArrayOutputWithContext(ctx context.Context) AccessReviewScopeResponseArrayOutput {
	return o
}

func (o AccessReviewScopeResponseArrayOutput) Index(i pulumi.IntInput) AccessReviewScopeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessReviewScopeResponse {
		return vs[0].([]AccessReviewScopeResponse)[vs[1].(int)]
	}).(AccessReviewScopeResponseOutput)
}

type ApprovalSettingsResponse struct {
	ApprovalMode                     *string                 `pulumi:"approvalMode"`
	ApprovalStages                   []ApprovalStageResponse `pulumi:"approvalStages"`
	IsApprovalRequired               *bool                   `pulumi:"isApprovalRequired"`
	IsApprovalRequiredForExtension   *bool                   `pulumi:"isApprovalRequiredForExtension"`
	IsRequestorJustificationRequired *bool                   `pulumi:"isRequestorJustificationRequired"`
}

type ApprovalStageResponse struct {
	ApprovalStageTimeOutInDays      *int              `pulumi:"approvalStageTimeOutInDays"`
	EscalationApprovers             []UserSetResponse `pulumi:"escalationApprovers"`
	EscalationTimeInMinutes         *int              `pulumi:"escalationTimeInMinutes"`
	IsApproverJustificationRequired *bool             `pulumi:"isApproverJustificationRequired"`
	IsEscalationEnabled             *bool             `pulumi:"isEscalationEnabled"`
	PrimaryApprovers                []UserSetResponse `pulumi:"primaryApprovers"`
}

type Identity struct {
	Type *ResourceIdentityType `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type ResourceIdentityTypePtrInput `pulumi:"type"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v Identity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type IdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ManagementLockOwner struct {
	ApplicationId *string `pulumi:"applicationId"`
}





type ManagementLockOwnerInput interface {
	pulumi.Input

	ToManagementLockOwnerOutput() ManagementLockOwnerOutput
	ToManagementLockOwnerOutputWithContext(context.Context) ManagementLockOwnerOutput
}

type ManagementLockOwnerArgs struct {
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
}

func (ManagementLockOwnerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockOwner)(nil)).Elem()
}

func (i ManagementLockOwnerArgs) ToManagementLockOwnerOutput() ManagementLockOwnerOutput {
	return i.ToManagementLockOwnerOutputWithContext(context.Background())
}

func (i ManagementLockOwnerArgs) ToManagementLockOwnerOutputWithContext(ctx context.Context) ManagementLockOwnerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockOwnerOutput)
}





type ManagementLockOwnerArrayInput interface {
	pulumi.Input

	ToManagementLockOwnerArrayOutput() ManagementLockOwnerArrayOutput
	ToManagementLockOwnerArrayOutputWithContext(context.Context) ManagementLockOwnerArrayOutput
}

type ManagementLockOwnerArray []ManagementLockOwnerInput

func (ManagementLockOwnerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementLockOwner)(nil)).Elem()
}

func (i ManagementLockOwnerArray) ToManagementLockOwnerArrayOutput() ManagementLockOwnerArrayOutput {
	return i.ToManagementLockOwnerArrayOutputWithContext(context.Background())
}

func (i ManagementLockOwnerArray) ToManagementLockOwnerArrayOutputWithContext(ctx context.Context) ManagementLockOwnerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockOwnerArrayOutput)
}

type ManagementLockOwnerOutput struct{ *pulumi.OutputState }

func (ManagementLockOwnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockOwner)(nil)).Elem()
}

func (o ManagementLockOwnerOutput) ToManagementLockOwnerOutput() ManagementLockOwnerOutput {
	return o
}

func (o ManagementLockOwnerOutput) ToManagementLockOwnerOutputWithContext(ctx context.Context) ManagementLockOwnerOutput {
	return o
}

func (o ManagementLockOwnerOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementLockOwner) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

type ManagementLockOwnerArrayOutput struct{ *pulumi.OutputState }

func (ManagementLockOwnerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementLockOwner)(nil)).Elem()
}

func (o ManagementLockOwnerArrayOutput) ToManagementLockOwnerArrayOutput() ManagementLockOwnerArrayOutput {
	return o
}

func (o ManagementLockOwnerArrayOutput) ToManagementLockOwnerArrayOutputWithContext(ctx context.Context) ManagementLockOwnerArrayOutput {
	return o
}

func (o ManagementLockOwnerArrayOutput) Index(i pulumi.IntInput) ManagementLockOwnerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagementLockOwner {
		return vs[0].([]ManagementLockOwner)[vs[1].(int)]
	}).(ManagementLockOwnerOutput)
}

type ManagementLockOwnerResponse struct {
	ApplicationId *string `pulumi:"applicationId"`
}

type ManagementLockOwnerResponseOutput struct{ *pulumi.OutputState }

func (ManagementLockOwnerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockOwnerResponse)(nil)).Elem()
}

func (o ManagementLockOwnerResponseOutput) ToManagementLockOwnerResponseOutput() ManagementLockOwnerResponseOutput {
	return o
}

func (o ManagementLockOwnerResponseOutput) ToManagementLockOwnerResponseOutputWithContext(ctx context.Context) ManagementLockOwnerResponseOutput {
	return o
}

func (o ManagementLockOwnerResponseOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementLockOwnerResponse) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

type ManagementLockOwnerResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagementLockOwnerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementLockOwnerResponse)(nil)).Elem()
}

func (o ManagementLockOwnerResponseArrayOutput) ToManagementLockOwnerResponseArrayOutput() ManagementLockOwnerResponseArrayOutput {
	return o
}

func (o ManagementLockOwnerResponseArrayOutput) ToManagementLockOwnerResponseArrayOutputWithContext(ctx context.Context) ManagementLockOwnerResponseArrayOutput {
	return o
}

func (o ManagementLockOwnerResponseArrayOutput) Index(i pulumi.IntInput) ManagementLockOwnerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagementLockOwnerResponse {
		return vs[0].([]ManagementLockOwnerResponse)[vs[1].(int)]
	}).(ManagementLockOwnerResponseOutput)
}

type NonComplianceMessage struct {
	Message                     string  `pulumi:"message"`
	PolicyDefinitionReferenceId *string `pulumi:"policyDefinitionReferenceId"`
}





type NonComplianceMessageInput interface {
	pulumi.Input

	ToNonComplianceMessageOutput() NonComplianceMessageOutput
	ToNonComplianceMessageOutputWithContext(context.Context) NonComplianceMessageOutput
}

type NonComplianceMessageArgs struct {
	Message                     pulumi.StringInput    `pulumi:"message"`
	PolicyDefinitionReferenceId pulumi.StringPtrInput `pulumi:"policyDefinitionReferenceId"`
}

func (NonComplianceMessageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NonComplianceMessage)(nil)).Elem()
}

func (i NonComplianceMessageArgs) ToNonComplianceMessageOutput() NonComplianceMessageOutput {
	return i.ToNonComplianceMessageOutputWithContext(context.Background())
}

func (i NonComplianceMessageArgs) ToNonComplianceMessageOutputWithContext(ctx context.Context) NonComplianceMessageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NonComplianceMessageOutput)
}





type NonComplianceMessageArrayInput interface {
	pulumi.Input

	ToNonComplianceMessageArrayOutput() NonComplianceMessageArrayOutput
	ToNonComplianceMessageArrayOutputWithContext(context.Context) NonComplianceMessageArrayOutput
}

type NonComplianceMessageArray []NonComplianceMessageInput

func (NonComplianceMessageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NonComplianceMessage)(nil)).Elem()
}

func (i NonComplianceMessageArray) ToNonComplianceMessageArrayOutput() NonComplianceMessageArrayOutput {
	return i.ToNonComplianceMessageArrayOutputWithContext(context.Background())
}

func (i NonComplianceMessageArray) ToNonComplianceMessageArrayOutputWithContext(ctx context.Context) NonComplianceMessageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NonComplianceMessageArrayOutput)
}

type NonComplianceMessageOutput struct{ *pulumi.OutputState }

func (NonComplianceMessageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NonComplianceMessage)(nil)).Elem()
}

func (o NonComplianceMessageOutput) ToNonComplianceMessageOutput() NonComplianceMessageOutput {
	return o
}

func (o NonComplianceMessageOutput) ToNonComplianceMessageOutputWithContext(ctx context.Context) NonComplianceMessageOutput {
	return o
}

func (o NonComplianceMessageOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v NonComplianceMessage) string { return v.Message }).(pulumi.StringOutput)
}

func (o NonComplianceMessageOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NonComplianceMessage) *string { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

type NonComplianceMessageArrayOutput struct{ *pulumi.OutputState }

func (NonComplianceMessageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NonComplianceMessage)(nil)).Elem()
}

func (o NonComplianceMessageArrayOutput) ToNonComplianceMessageArrayOutput() NonComplianceMessageArrayOutput {
	return o
}

func (o NonComplianceMessageArrayOutput) ToNonComplianceMessageArrayOutputWithContext(ctx context.Context) NonComplianceMessageArrayOutput {
	return o
}

func (o NonComplianceMessageArrayOutput) Index(i pulumi.IntInput) NonComplianceMessageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NonComplianceMessage {
		return vs[0].([]NonComplianceMessage)[vs[1].(int)]
	}).(NonComplianceMessageOutput)
}

type NonComplianceMessageResponse struct {
	Message                     string  `pulumi:"message"`
	PolicyDefinitionReferenceId *string `pulumi:"policyDefinitionReferenceId"`
}

type NonComplianceMessageResponseOutput struct{ *pulumi.OutputState }

func (NonComplianceMessageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NonComplianceMessageResponse)(nil)).Elem()
}

func (o NonComplianceMessageResponseOutput) ToNonComplianceMessageResponseOutput() NonComplianceMessageResponseOutput {
	return o
}

func (o NonComplianceMessageResponseOutput) ToNonComplianceMessageResponseOutputWithContext(ctx context.Context) NonComplianceMessageResponseOutput {
	return o
}

func (o NonComplianceMessageResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v NonComplianceMessageResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o NonComplianceMessageResponseOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NonComplianceMessageResponse) *string { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

type NonComplianceMessageResponseArrayOutput struct{ *pulumi.OutputState }

func (NonComplianceMessageResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NonComplianceMessageResponse)(nil)).Elem()
}

func (o NonComplianceMessageResponseArrayOutput) ToNonComplianceMessageResponseArrayOutput() NonComplianceMessageResponseArrayOutput {
	return o
}

func (o NonComplianceMessageResponseArrayOutput) ToNonComplianceMessageResponseArrayOutputWithContext(ctx context.Context) NonComplianceMessageResponseArrayOutput {
	return o
}

func (o NonComplianceMessageResponseArrayOutput) Index(i pulumi.IntInput) NonComplianceMessageResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NonComplianceMessageResponse {
		return vs[0].([]NonComplianceMessageResponse)[vs[1].(int)]
	}).(NonComplianceMessageResponseOutput)
}

type ParameterDefinitionsValue struct {
	AllowedValues []interface{}                      `pulumi:"allowedValues"`
	DefaultValue  interface{}                        `pulumi:"defaultValue"`
	Metadata      *ParameterDefinitionsValueMetadata `pulumi:"metadata"`
	Type          *string                            `pulumi:"type"`
}





type ParameterDefinitionsValueInput interface {
	pulumi.Input

	ToParameterDefinitionsValueOutput() ParameterDefinitionsValueOutput
	ToParameterDefinitionsValueOutputWithContext(context.Context) ParameterDefinitionsValueOutput
}

type ParameterDefinitionsValueArgs struct {
	AllowedValues pulumi.ArrayInput                         `pulumi:"allowedValues"`
	DefaultValue  pulumi.Input                              `pulumi:"defaultValue"`
	Metadata      ParameterDefinitionsValueMetadataPtrInput `pulumi:"metadata"`
	Type          pulumi.StringPtrInput                     `pulumi:"type"`
}

func (ParameterDefinitionsValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinitionsValue)(nil)).Elem()
}

func (i ParameterDefinitionsValueArgs) ToParameterDefinitionsValueOutput() ParameterDefinitionsValueOutput {
	return i.ToParameterDefinitionsValueOutputWithContext(context.Background())
}

func (i ParameterDefinitionsValueArgs) ToParameterDefinitionsValueOutputWithContext(ctx context.Context) ParameterDefinitionsValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDefinitionsValueOutput)
}





type ParameterDefinitionsValueMapInput interface {
	pulumi.Input

	ToParameterDefinitionsValueMapOutput() ParameterDefinitionsValueMapOutput
	ToParameterDefinitionsValueMapOutputWithContext(context.Context) ParameterDefinitionsValueMapOutput
}

type ParameterDefinitionsValueMap map[string]ParameterDefinitionsValueInput

func (ParameterDefinitionsValueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterDefinitionsValue)(nil)).Elem()
}

func (i ParameterDefinitionsValueMap) ToParameterDefinitionsValueMapOutput() ParameterDefinitionsValueMapOutput {
	return i.ToParameterDefinitionsValueMapOutputWithContext(context.Background())
}

func (i ParameterDefinitionsValueMap) ToParameterDefinitionsValueMapOutputWithContext(ctx context.Context) ParameterDefinitionsValueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDefinitionsValueMapOutput)
}

type ParameterDefinitionsValueOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionsValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinitionsValue)(nil)).Elem()
}

func (o ParameterDefinitionsValueOutput) ToParameterDefinitionsValueOutput() ParameterDefinitionsValueOutput {
	return o
}

func (o ParameterDefinitionsValueOutput) ToParameterDefinitionsValueOutputWithContext(ctx context.Context) ParameterDefinitionsValueOutput {
	return o
}

func (o ParameterDefinitionsValueOutput) AllowedValues() pulumi.ArrayOutput {
	return o.ApplyT(func(v ParameterDefinitionsValue) []interface{} { return v.AllowedValues }).(pulumi.ArrayOutput)
}

func (o ParameterDefinitionsValueOutput) DefaultValue() pulumi.AnyOutput {
	return o.ApplyT(func(v ParameterDefinitionsValue) interface{} { return v.DefaultValue }).(pulumi.AnyOutput)
}

func (o ParameterDefinitionsValueOutput) Metadata() ParameterDefinitionsValueMetadataPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionsValue) *ParameterDefinitionsValueMetadata { return v.Metadata }).(ParameterDefinitionsValueMetadataPtrOutput)
}

func (o ParameterDefinitionsValueOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionsValue) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ParameterDefinitionsValueMapOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionsValueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterDefinitionsValue)(nil)).Elem()
}

func (o ParameterDefinitionsValueMapOutput) ToParameterDefinitionsValueMapOutput() ParameterDefinitionsValueMapOutput {
	return o
}

func (o ParameterDefinitionsValueMapOutput) ToParameterDefinitionsValueMapOutputWithContext(ctx context.Context) ParameterDefinitionsValueMapOutput {
	return o
}

func (o ParameterDefinitionsValueMapOutput) MapIndex(k pulumi.StringInput) ParameterDefinitionsValueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterDefinitionsValue {
		return vs[0].(map[string]ParameterDefinitionsValue)[vs[1].(string)]
	}).(ParameterDefinitionsValueOutput)
}

type ParameterDefinitionsValueMetadata struct {
	AssignPermissions *bool   `pulumi:"assignPermissions"`
	Description       *string `pulumi:"description"`
	DisplayName       *string `pulumi:"displayName"`
	StrongType        *string `pulumi:"strongType"`
}





type ParameterDefinitionsValueMetadataInput interface {
	pulumi.Input

	ToParameterDefinitionsValueMetadataOutput() ParameterDefinitionsValueMetadataOutput
	ToParameterDefinitionsValueMetadataOutputWithContext(context.Context) ParameterDefinitionsValueMetadataOutput
}

type ParameterDefinitionsValueMetadataArgs struct {
	AssignPermissions pulumi.BoolPtrInput   `pulumi:"assignPermissions"`
	Description       pulumi.StringPtrInput `pulumi:"description"`
	DisplayName       pulumi.StringPtrInput `pulumi:"displayName"`
	StrongType        pulumi.StringPtrInput `pulumi:"strongType"`
}

func (ParameterDefinitionsValueMetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinitionsValueMetadata)(nil)).Elem()
}

func (i ParameterDefinitionsValueMetadataArgs) ToParameterDefinitionsValueMetadataOutput() ParameterDefinitionsValueMetadataOutput {
	return i.ToParameterDefinitionsValueMetadataOutputWithContext(context.Background())
}

func (i ParameterDefinitionsValueMetadataArgs) ToParameterDefinitionsValueMetadataOutputWithContext(ctx context.Context) ParameterDefinitionsValueMetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDefinitionsValueMetadataOutput)
}

func (i ParameterDefinitionsValueMetadataArgs) ToParameterDefinitionsValueMetadataPtrOutput() ParameterDefinitionsValueMetadataPtrOutput {
	return i.ToParameterDefinitionsValueMetadataPtrOutputWithContext(context.Background())
}

func (i ParameterDefinitionsValueMetadataArgs) ToParameterDefinitionsValueMetadataPtrOutputWithContext(ctx context.Context) ParameterDefinitionsValueMetadataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDefinitionsValueMetadataOutput).ToParameterDefinitionsValueMetadataPtrOutputWithContext(ctx)
}









type ParameterDefinitionsValueMetadataPtrInput interface {
	pulumi.Input

	ToParameterDefinitionsValueMetadataPtrOutput() ParameterDefinitionsValueMetadataPtrOutput
	ToParameterDefinitionsValueMetadataPtrOutputWithContext(context.Context) ParameterDefinitionsValueMetadataPtrOutput
}

type parameterDefinitionsValueMetadataPtrType ParameterDefinitionsValueMetadataArgs

func ParameterDefinitionsValueMetadataPtr(v *ParameterDefinitionsValueMetadataArgs) ParameterDefinitionsValueMetadataPtrInput {
	return (*parameterDefinitionsValueMetadataPtrType)(v)
}

func (*parameterDefinitionsValueMetadataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterDefinitionsValueMetadata)(nil)).Elem()
}

func (i *parameterDefinitionsValueMetadataPtrType) ToParameterDefinitionsValueMetadataPtrOutput() ParameterDefinitionsValueMetadataPtrOutput {
	return i.ToParameterDefinitionsValueMetadataPtrOutputWithContext(context.Background())
}

func (i *parameterDefinitionsValueMetadataPtrType) ToParameterDefinitionsValueMetadataPtrOutputWithContext(ctx context.Context) ParameterDefinitionsValueMetadataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDefinitionsValueMetadataPtrOutput)
}

type ParameterDefinitionsValueMetadataOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionsValueMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinitionsValueMetadata)(nil)).Elem()
}

func (o ParameterDefinitionsValueMetadataOutput) ToParameterDefinitionsValueMetadataOutput() ParameterDefinitionsValueMetadataOutput {
	return o
}

func (o ParameterDefinitionsValueMetadataOutput) ToParameterDefinitionsValueMetadataOutputWithContext(ctx context.Context) ParameterDefinitionsValueMetadataOutput {
	return o
}

func (o ParameterDefinitionsValueMetadataOutput) ToParameterDefinitionsValueMetadataPtrOutput() ParameterDefinitionsValueMetadataPtrOutput {
	return o.ToParameterDefinitionsValueMetadataPtrOutputWithContext(context.Background())
}

func (o ParameterDefinitionsValueMetadataOutput) ToParameterDefinitionsValueMetadataPtrOutputWithContext(ctx context.Context) ParameterDefinitionsValueMetadataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ParameterDefinitionsValueMetadata) *ParameterDefinitionsValueMetadata {
		return &v
	}).(ParameterDefinitionsValueMetadataPtrOutput)
}

func (o ParameterDefinitionsValueMetadataOutput) AssignPermissions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionsValueMetadata) *bool { return v.AssignPermissions }).(pulumi.BoolPtrOutput)
}

func (o ParameterDefinitionsValueMetadataOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionsValueMetadata) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionsValueMetadataOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionsValueMetadata) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionsValueMetadataOutput) StrongType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionsValueMetadata) *string { return v.StrongType }).(pulumi.StringPtrOutput)
}

type ParameterDefinitionsValueMetadataPtrOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionsValueMetadataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterDefinitionsValueMetadata)(nil)).Elem()
}

func (o ParameterDefinitionsValueMetadataPtrOutput) ToParameterDefinitionsValueMetadataPtrOutput() ParameterDefinitionsValueMetadataPtrOutput {
	return o
}

func (o ParameterDefinitionsValueMetadataPtrOutput) ToParameterDefinitionsValueMetadataPtrOutputWithContext(ctx context.Context) ParameterDefinitionsValueMetadataPtrOutput {
	return o
}

func (o ParameterDefinitionsValueMetadataPtrOutput) Elem() ParameterDefinitionsValueMetadataOutput {
	return o.ApplyT(func(v *ParameterDefinitionsValueMetadata) ParameterDefinitionsValueMetadata {
		if v != nil {
			return *v
		}
		var ret ParameterDefinitionsValueMetadata
		return ret
	}).(ParameterDefinitionsValueMetadataOutput)
}

func (o ParameterDefinitionsValueMetadataPtrOutput) AssignPermissions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ParameterDefinitionsValueMetadata) *bool {
		if v == nil {
			return nil
		}
		return v.AssignPermissions
	}).(pulumi.BoolPtrOutput)
}

func (o ParameterDefinitionsValueMetadataPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParameterDefinitionsValueMetadata) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionsValueMetadataPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParameterDefinitionsValueMetadata) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionsValueMetadataPtrOutput) StrongType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParameterDefinitionsValueMetadata) *string {
		if v == nil {
			return nil
		}
		return v.StrongType
	}).(pulumi.StringPtrOutput)
}

type ParameterDefinitionsValueResponse struct {
	AllowedValues []interface{}                              `pulumi:"allowedValues"`
	DefaultValue  interface{}                                `pulumi:"defaultValue"`
	Metadata      *ParameterDefinitionsValueResponseMetadata `pulumi:"metadata"`
	Type          *string                                    `pulumi:"type"`
}

type ParameterDefinitionsValueResponseOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionsValueResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinitionsValueResponse)(nil)).Elem()
}

func (o ParameterDefinitionsValueResponseOutput) ToParameterDefinitionsValueResponseOutput() ParameterDefinitionsValueResponseOutput {
	return o
}

func (o ParameterDefinitionsValueResponseOutput) ToParameterDefinitionsValueResponseOutputWithContext(ctx context.Context) ParameterDefinitionsValueResponseOutput {
	return o
}

func (o ParameterDefinitionsValueResponseOutput) AllowedValues() pulumi.ArrayOutput {
	return o.ApplyT(func(v ParameterDefinitionsValueResponse) []interface{} { return v.AllowedValues }).(pulumi.ArrayOutput)
}

func (o ParameterDefinitionsValueResponseOutput) DefaultValue() pulumi.AnyOutput {
	return o.ApplyT(func(v ParameterDefinitionsValueResponse) interface{} { return v.DefaultValue }).(pulumi.AnyOutput)
}

func (o ParameterDefinitionsValueResponseOutput) Metadata() ParameterDefinitionsValueResponseMetadataPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionsValueResponse) *ParameterDefinitionsValueResponseMetadata {
		return v.Metadata
	}).(ParameterDefinitionsValueResponseMetadataPtrOutput)
}

func (o ParameterDefinitionsValueResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionsValueResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ParameterDefinitionsValueResponseMapOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionsValueResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterDefinitionsValueResponse)(nil)).Elem()
}

func (o ParameterDefinitionsValueResponseMapOutput) ToParameterDefinitionsValueResponseMapOutput() ParameterDefinitionsValueResponseMapOutput {
	return o
}

func (o ParameterDefinitionsValueResponseMapOutput) ToParameterDefinitionsValueResponseMapOutputWithContext(ctx context.Context) ParameterDefinitionsValueResponseMapOutput {
	return o
}

func (o ParameterDefinitionsValueResponseMapOutput) MapIndex(k pulumi.StringInput) ParameterDefinitionsValueResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterDefinitionsValueResponse {
		return vs[0].(map[string]ParameterDefinitionsValueResponse)[vs[1].(string)]
	}).(ParameterDefinitionsValueResponseOutput)
}

type ParameterDefinitionsValueResponseMetadata struct {
	AssignPermissions *bool   `pulumi:"assignPermissions"`
	Description       *string `pulumi:"description"`
	DisplayName       *string `pulumi:"displayName"`
	StrongType        *string `pulumi:"strongType"`
}

type ParameterDefinitionsValueResponseMetadataOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionsValueResponseMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinitionsValueResponseMetadata)(nil)).Elem()
}

func (o ParameterDefinitionsValueResponseMetadataOutput) ToParameterDefinitionsValueResponseMetadataOutput() ParameterDefinitionsValueResponseMetadataOutput {
	return o
}

func (o ParameterDefinitionsValueResponseMetadataOutput) ToParameterDefinitionsValueResponseMetadataOutputWithContext(ctx context.Context) ParameterDefinitionsValueResponseMetadataOutput {
	return o
}

func (o ParameterDefinitionsValueResponseMetadataOutput) AssignPermissions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionsValueResponseMetadata) *bool { return v.AssignPermissions }).(pulumi.BoolPtrOutput)
}

func (o ParameterDefinitionsValueResponseMetadataOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionsValueResponseMetadata) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionsValueResponseMetadataOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionsValueResponseMetadata) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionsValueResponseMetadataOutput) StrongType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionsValueResponseMetadata) *string { return v.StrongType }).(pulumi.StringPtrOutput)
}

type ParameterDefinitionsValueResponseMetadataPtrOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionsValueResponseMetadataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterDefinitionsValueResponseMetadata)(nil)).Elem()
}

func (o ParameterDefinitionsValueResponseMetadataPtrOutput) ToParameterDefinitionsValueResponseMetadataPtrOutput() ParameterDefinitionsValueResponseMetadataPtrOutput {
	return o
}

func (o ParameterDefinitionsValueResponseMetadataPtrOutput) ToParameterDefinitionsValueResponseMetadataPtrOutputWithContext(ctx context.Context) ParameterDefinitionsValueResponseMetadataPtrOutput {
	return o
}

func (o ParameterDefinitionsValueResponseMetadataPtrOutput) Elem() ParameterDefinitionsValueResponseMetadataOutput {
	return o.ApplyT(func(v *ParameterDefinitionsValueResponseMetadata) ParameterDefinitionsValueResponseMetadata {
		if v != nil {
			return *v
		}
		var ret ParameterDefinitionsValueResponseMetadata
		return ret
	}).(ParameterDefinitionsValueResponseMetadataOutput)
}

func (o ParameterDefinitionsValueResponseMetadataPtrOutput) AssignPermissions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ParameterDefinitionsValueResponseMetadata) *bool {
		if v == nil {
			return nil
		}
		return v.AssignPermissions
	}).(pulumi.BoolPtrOutput)
}

func (o ParameterDefinitionsValueResponseMetadataPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParameterDefinitionsValueResponseMetadata) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionsValueResponseMetadataPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParameterDefinitionsValueResponseMetadata) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionsValueResponseMetadataPtrOutput) StrongType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParameterDefinitionsValueResponseMetadata) *string {
		if v == nil {
			return nil
		}
		return v.StrongType
	}).(pulumi.StringPtrOutput)
}

type ParameterValuesValue struct {
	Value interface{} `pulumi:"value"`
}





type ParameterValuesValueInput interface {
	pulumi.Input

	ToParameterValuesValueOutput() ParameterValuesValueOutput
	ToParameterValuesValueOutputWithContext(context.Context) ParameterValuesValueOutput
}

type ParameterValuesValueArgs struct {
	Value pulumi.Input `pulumi:"value"`
}

func (ParameterValuesValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterValuesValue)(nil)).Elem()
}

func (i ParameterValuesValueArgs) ToParameterValuesValueOutput() ParameterValuesValueOutput {
	return i.ToParameterValuesValueOutputWithContext(context.Background())
}

func (i ParameterValuesValueArgs) ToParameterValuesValueOutputWithContext(ctx context.Context) ParameterValuesValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterValuesValueOutput)
}





type ParameterValuesValueMapInput interface {
	pulumi.Input

	ToParameterValuesValueMapOutput() ParameterValuesValueMapOutput
	ToParameterValuesValueMapOutputWithContext(context.Context) ParameterValuesValueMapOutput
}

type ParameterValuesValueMap map[string]ParameterValuesValueInput

func (ParameterValuesValueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterValuesValue)(nil)).Elem()
}

func (i ParameterValuesValueMap) ToParameterValuesValueMapOutput() ParameterValuesValueMapOutput {
	return i.ToParameterValuesValueMapOutputWithContext(context.Background())
}

func (i ParameterValuesValueMap) ToParameterValuesValueMapOutputWithContext(ctx context.Context) ParameterValuesValueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterValuesValueMapOutput)
}

type ParameterValuesValueOutput struct{ *pulumi.OutputState }

func (ParameterValuesValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterValuesValue)(nil)).Elem()
}

func (o ParameterValuesValueOutput) ToParameterValuesValueOutput() ParameterValuesValueOutput {
	return o
}

func (o ParameterValuesValueOutput) ToParameterValuesValueOutputWithContext(ctx context.Context) ParameterValuesValueOutput {
	return o
}

func (o ParameterValuesValueOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v ParameterValuesValue) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type ParameterValuesValueMapOutput struct{ *pulumi.OutputState }

func (ParameterValuesValueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterValuesValue)(nil)).Elem()
}

func (o ParameterValuesValueMapOutput) ToParameterValuesValueMapOutput() ParameterValuesValueMapOutput {
	return o
}

func (o ParameterValuesValueMapOutput) ToParameterValuesValueMapOutputWithContext(ctx context.Context) ParameterValuesValueMapOutput {
	return o
}

func (o ParameterValuesValueMapOutput) MapIndex(k pulumi.StringInput) ParameterValuesValueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterValuesValue {
		return vs[0].(map[string]ParameterValuesValue)[vs[1].(string)]
	}).(ParameterValuesValueOutput)
}

type ParameterValuesValueResponse struct {
	Value interface{} `pulumi:"value"`
}

type ParameterValuesValueResponseOutput struct{ *pulumi.OutputState }

func (ParameterValuesValueResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterValuesValueResponse)(nil)).Elem()
}

func (o ParameterValuesValueResponseOutput) ToParameterValuesValueResponseOutput() ParameterValuesValueResponseOutput {
	return o
}

func (o ParameterValuesValueResponseOutput) ToParameterValuesValueResponseOutputWithContext(ctx context.Context) ParameterValuesValueResponseOutput {
	return o
}

func (o ParameterValuesValueResponseOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v ParameterValuesValueResponse) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type ParameterValuesValueResponseMapOutput struct{ *pulumi.OutputState }

func (ParameterValuesValueResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterValuesValueResponse)(nil)).Elem()
}

func (o ParameterValuesValueResponseMapOutput) ToParameterValuesValueResponseMapOutput() ParameterValuesValueResponseMapOutput {
	return o
}

func (o ParameterValuesValueResponseMapOutput) ToParameterValuesValueResponseMapOutputWithContext(ctx context.Context) ParameterValuesValueResponseMapOutput {
	return o
}

func (o ParameterValuesValueResponseMapOutput) MapIndex(k pulumi.StringInput) ParameterValuesValueResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterValuesValueResponse {
		return vs[0].(map[string]ParameterValuesValueResponse)[vs[1].(string)]
	}).(ParameterValuesValueResponseOutput)
}

type Permission struct {
	Actions        []string `pulumi:"actions"`
	DataActions    []string `pulumi:"dataActions"`
	NotActions     []string `pulumi:"notActions"`
	NotDataActions []string `pulumi:"notDataActions"`
}





type PermissionInput interface {
	pulumi.Input

	ToPermissionOutput() PermissionOutput
	ToPermissionOutputWithContext(context.Context) PermissionOutput
}

type PermissionArgs struct {
	Actions        pulumi.StringArrayInput `pulumi:"actions"`
	DataActions    pulumi.StringArrayInput `pulumi:"dataActions"`
	NotActions     pulumi.StringArrayInput `pulumi:"notActions"`
	NotDataActions pulumi.StringArrayInput `pulumi:"notDataActions"`
}

func (PermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Permission)(nil)).Elem()
}

func (i PermissionArgs) ToPermissionOutput() PermissionOutput {
	return i.ToPermissionOutputWithContext(context.Background())
}

func (i PermissionArgs) ToPermissionOutputWithContext(ctx context.Context) PermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionOutput)
}





type PermissionArrayInput interface {
	pulumi.Input

	ToPermissionArrayOutput() PermissionArrayOutput
	ToPermissionArrayOutputWithContext(context.Context) PermissionArrayOutput
}

type PermissionArray []PermissionInput

func (PermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Permission)(nil)).Elem()
}

func (i PermissionArray) ToPermissionArrayOutput() PermissionArrayOutput {
	return i.ToPermissionArrayOutputWithContext(context.Background())
}

func (i PermissionArray) ToPermissionArrayOutputWithContext(ctx context.Context) PermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionArrayOutput)
}

type PermissionOutput struct{ *pulumi.OutputState }

func (PermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Permission)(nil)).Elem()
}

func (o PermissionOutput) ToPermissionOutput() PermissionOutput {
	return o
}

func (o PermissionOutput) ToPermissionOutputWithContext(ctx context.Context) PermissionOutput {
	return o
}

func (o PermissionOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Permission) []string { return v.Actions }).(pulumi.StringArrayOutput)
}

func (o PermissionOutput) DataActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Permission) []string { return v.DataActions }).(pulumi.StringArrayOutput)
}

func (o PermissionOutput) NotActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Permission) []string { return v.NotActions }).(pulumi.StringArrayOutput)
}

func (o PermissionOutput) NotDataActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Permission) []string { return v.NotDataActions }).(pulumi.StringArrayOutput)
}

type PermissionArrayOutput struct{ *pulumi.OutputState }

func (PermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Permission)(nil)).Elem()
}

func (o PermissionArrayOutput) ToPermissionArrayOutput() PermissionArrayOutput {
	return o
}

func (o PermissionArrayOutput) ToPermissionArrayOutputWithContext(ctx context.Context) PermissionArrayOutput {
	return o
}

func (o PermissionArrayOutput) Index(i pulumi.IntInput) PermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Permission {
		return vs[0].([]Permission)[vs[1].(int)]
	}).(PermissionOutput)
}

type PermissionResponse struct {
	Actions        []string `pulumi:"actions"`
	DataActions    []string `pulumi:"dataActions"`
	NotActions     []string `pulumi:"notActions"`
	NotDataActions []string `pulumi:"notDataActions"`
}

type PermissionResponseOutput struct{ *pulumi.OutputState }

func (PermissionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionResponse)(nil)).Elem()
}

func (o PermissionResponseOutput) ToPermissionResponseOutput() PermissionResponseOutput {
	return o
}

func (o PermissionResponseOutput) ToPermissionResponseOutputWithContext(ctx context.Context) PermissionResponseOutput {
	return o
}

func (o PermissionResponseOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PermissionResponse) []string { return v.Actions }).(pulumi.StringArrayOutput)
}

func (o PermissionResponseOutput) DataActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PermissionResponse) []string { return v.DataActions }).(pulumi.StringArrayOutput)
}

func (o PermissionResponseOutput) NotActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PermissionResponse) []string { return v.NotActions }).(pulumi.StringArrayOutput)
}

func (o PermissionResponseOutput) NotDataActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PermissionResponse) []string { return v.NotDataActions }).(pulumi.StringArrayOutput)
}

type PermissionResponseArrayOutput struct{ *pulumi.OutputState }

func (PermissionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PermissionResponse)(nil)).Elem()
}

func (o PermissionResponseArrayOutput) ToPermissionResponseArrayOutput() PermissionResponseArrayOutput {
	return o
}

func (o PermissionResponseArrayOutput) ToPermissionResponseArrayOutputWithContext(ctx context.Context) PermissionResponseArrayOutput {
	return o
}

func (o PermissionResponseArrayOutput) Index(i pulumi.IntInput) PermissionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PermissionResponse {
		return vs[0].([]PermissionResponse)[vs[1].(int)]
	}).(PermissionResponseOutput)
}

type PolicyAssignmentPropertiesResponse struct {
	Policy         *PolicyAssignmentPropertiesResponsePolicy         `pulumi:"policy"`
	RoleDefinition *PolicyAssignmentPropertiesResponseRoleDefinition `pulumi:"roleDefinition"`
	Scope          *PolicyAssignmentPropertiesResponseScope          `pulumi:"scope"`
}

type PolicyAssignmentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PolicyAssignmentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAssignmentPropertiesResponse)(nil)).Elem()
}

func (o PolicyAssignmentPropertiesResponseOutput) ToPolicyAssignmentPropertiesResponseOutput() PolicyAssignmentPropertiesResponseOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponseOutput) ToPolicyAssignmentPropertiesResponseOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponseOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponseOutput) Policy() PolicyAssignmentPropertiesResponsePolicyPtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponse) *PolicyAssignmentPropertiesResponsePolicy { return v.Policy }).(PolicyAssignmentPropertiesResponsePolicyPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseOutput) RoleDefinition() PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponse) *PolicyAssignmentPropertiesResponseRoleDefinition {
		return v.RoleDefinition
	}).(PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseOutput) Scope() PolicyAssignmentPropertiesResponseScopePtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponse) *PolicyAssignmentPropertiesResponseScope { return v.Scope }).(PolicyAssignmentPropertiesResponseScopePtrOutput)
}

type PolicyAssignmentPropertiesResponsePolicy struct {
	Id                   *string           `pulumi:"id"`
	LastModifiedBy       PrincipalResponse `pulumi:"lastModifiedBy"`
	LastModifiedDateTime *string           `pulumi:"lastModifiedDateTime"`
}

type PolicyAssignmentPropertiesResponsePolicyOutput struct{ *pulumi.OutputState }

func (PolicyAssignmentPropertiesResponsePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAssignmentPropertiesResponsePolicy)(nil)).Elem()
}

func (o PolicyAssignmentPropertiesResponsePolicyOutput) ToPolicyAssignmentPropertiesResponsePolicyOutput() PolicyAssignmentPropertiesResponsePolicyOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponsePolicyOutput) ToPolicyAssignmentPropertiesResponsePolicyOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponsePolicyOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponsePolicyOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponsePolicy) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentPropertiesResponsePolicyOutput) LastModifiedBy() PrincipalResponseOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponsePolicy) PrincipalResponse { return v.LastModifiedBy }).(PrincipalResponseOutput)
}

func (o PolicyAssignmentPropertiesResponsePolicyOutput) LastModifiedDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponsePolicy) *string { return v.LastModifiedDateTime }).(pulumi.StringPtrOutput)
}

type PolicyAssignmentPropertiesResponsePolicyPtrOutput struct{ *pulumi.OutputState }

func (PolicyAssignmentPropertiesResponsePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAssignmentPropertiesResponsePolicy)(nil)).Elem()
}

func (o PolicyAssignmentPropertiesResponsePolicyPtrOutput) ToPolicyAssignmentPropertiesResponsePolicyPtrOutput() PolicyAssignmentPropertiesResponsePolicyPtrOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponsePolicyPtrOutput) ToPolicyAssignmentPropertiesResponsePolicyPtrOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponsePolicyPtrOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponsePolicyPtrOutput) Elem() PolicyAssignmentPropertiesResponsePolicyOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponsePolicy) PolicyAssignmentPropertiesResponsePolicy {
		if v != nil {
			return *v
		}
		var ret PolicyAssignmentPropertiesResponsePolicy
		return ret
	}).(PolicyAssignmentPropertiesResponsePolicyOutput)
}

func (o PolicyAssignmentPropertiesResponsePolicyPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponsePolicy) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentPropertiesResponsePolicyPtrOutput) LastModifiedBy() PrincipalResponsePtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponsePolicy) *PrincipalResponse {
		if v == nil {
			return nil
		}
		return &v.LastModifiedBy
	}).(PrincipalResponsePtrOutput)
}

func (o PolicyAssignmentPropertiesResponsePolicyPtrOutput) LastModifiedDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponsePolicy) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedDateTime
	}).(pulumi.StringPtrOutput)
}

type PolicyAssignmentPropertiesResponseRoleDefinition struct {
	DisplayName *string `pulumi:"displayName"`
	Id          *string `pulumi:"id"`
	Type        *string `pulumi:"type"`
}

type PolicyAssignmentPropertiesResponseRoleDefinitionOutput struct{ *pulumi.OutputState }

func (PolicyAssignmentPropertiesResponseRoleDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAssignmentPropertiesResponseRoleDefinition)(nil)).Elem()
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionOutput) ToPolicyAssignmentPropertiesResponseRoleDefinitionOutput() PolicyAssignmentPropertiesResponseRoleDefinitionOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionOutput) ToPolicyAssignmentPropertiesResponseRoleDefinitionOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponseRoleDefinitionOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponseRoleDefinition) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponseRoleDefinition) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponseRoleDefinition) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput struct{ *pulumi.OutputState }

func (PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAssignmentPropertiesResponseRoleDefinition)(nil)).Elem()
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput) ToPolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput() PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput) ToPolicyAssignmentPropertiesResponseRoleDefinitionPtrOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput) Elem() PolicyAssignmentPropertiesResponseRoleDefinitionOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponseRoleDefinition) PolicyAssignmentPropertiesResponseRoleDefinition {
		if v != nil {
			return *v
		}
		var ret PolicyAssignmentPropertiesResponseRoleDefinition
		return ret
	}).(PolicyAssignmentPropertiesResponseRoleDefinitionOutput)
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponseRoleDefinition) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponseRoleDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponseRoleDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type PolicyAssignmentPropertiesResponseScope struct {
	DisplayName *string `pulumi:"displayName"`
	Id          *string `pulumi:"id"`
	Type        *string `pulumi:"type"`
}

type PolicyAssignmentPropertiesResponseScopeOutput struct{ *pulumi.OutputState }

func (PolicyAssignmentPropertiesResponseScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAssignmentPropertiesResponseScope)(nil)).Elem()
}

func (o PolicyAssignmentPropertiesResponseScopeOutput) ToPolicyAssignmentPropertiesResponseScopeOutput() PolicyAssignmentPropertiesResponseScopeOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponseScopeOutput) ToPolicyAssignmentPropertiesResponseScopeOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponseScopeOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponseScopeOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponseScope) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseScopeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponseScope) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseScopeOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponseScope) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PolicyAssignmentPropertiesResponseScopePtrOutput struct{ *pulumi.OutputState }

func (PolicyAssignmentPropertiesResponseScopePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAssignmentPropertiesResponseScope)(nil)).Elem()
}

func (o PolicyAssignmentPropertiesResponseScopePtrOutput) ToPolicyAssignmentPropertiesResponseScopePtrOutput() PolicyAssignmentPropertiesResponseScopePtrOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponseScopePtrOutput) ToPolicyAssignmentPropertiesResponseScopePtrOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponseScopePtrOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponseScopePtrOutput) Elem() PolicyAssignmentPropertiesResponseScopeOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponseScope) PolicyAssignmentPropertiesResponseScope {
		if v != nil {
			return *v
		}
		var ret PolicyAssignmentPropertiesResponseScope
		return ret
	}).(PolicyAssignmentPropertiesResponseScopeOutput)
}

func (o PolicyAssignmentPropertiesResponseScopePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponseScope) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseScopePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponseScope) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseScopePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponseScope) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type PolicyDefinitionGroup struct {
	AdditionalMetadataId *string `pulumi:"additionalMetadataId"`
	Category             *string `pulumi:"category"`
	Description          *string `pulumi:"description"`
	DisplayName          *string `pulumi:"displayName"`
	Name                 string  `pulumi:"name"`
}





type PolicyDefinitionGroupInput interface {
	pulumi.Input

	ToPolicyDefinitionGroupOutput() PolicyDefinitionGroupOutput
	ToPolicyDefinitionGroupOutputWithContext(context.Context) PolicyDefinitionGroupOutput
}

type PolicyDefinitionGroupArgs struct {
	AdditionalMetadataId pulumi.StringPtrInput `pulumi:"additionalMetadataId"`
	Category             pulumi.StringPtrInput `pulumi:"category"`
	Description          pulumi.StringPtrInput `pulumi:"description"`
	DisplayName          pulumi.StringPtrInput `pulumi:"displayName"`
	Name                 pulumi.StringInput    `pulumi:"name"`
}

func (PolicyDefinitionGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyDefinitionGroup)(nil)).Elem()
}

func (i PolicyDefinitionGroupArgs) ToPolicyDefinitionGroupOutput() PolicyDefinitionGroupOutput {
	return i.ToPolicyDefinitionGroupOutputWithContext(context.Background())
}

func (i PolicyDefinitionGroupArgs) ToPolicyDefinitionGroupOutputWithContext(ctx context.Context) PolicyDefinitionGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyDefinitionGroupOutput)
}





type PolicyDefinitionGroupArrayInput interface {
	pulumi.Input

	ToPolicyDefinitionGroupArrayOutput() PolicyDefinitionGroupArrayOutput
	ToPolicyDefinitionGroupArrayOutputWithContext(context.Context) PolicyDefinitionGroupArrayOutput
}

type PolicyDefinitionGroupArray []PolicyDefinitionGroupInput

func (PolicyDefinitionGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyDefinitionGroup)(nil)).Elem()
}

func (i PolicyDefinitionGroupArray) ToPolicyDefinitionGroupArrayOutput() PolicyDefinitionGroupArrayOutput {
	return i.ToPolicyDefinitionGroupArrayOutputWithContext(context.Background())
}

func (i PolicyDefinitionGroupArray) ToPolicyDefinitionGroupArrayOutputWithContext(ctx context.Context) PolicyDefinitionGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyDefinitionGroupArrayOutput)
}

type PolicyDefinitionGroupOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyDefinitionGroup)(nil)).Elem()
}

func (o PolicyDefinitionGroupOutput) ToPolicyDefinitionGroupOutput() PolicyDefinitionGroupOutput {
	return o
}

func (o PolicyDefinitionGroupOutput) ToPolicyDefinitionGroupOutputWithContext(ctx context.Context) PolicyDefinitionGroupOutput {
	return o
}

func (o PolicyDefinitionGroupOutput) AdditionalMetadataId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionGroup) *string { return v.AdditionalMetadataId }).(pulumi.StringPtrOutput)
}

func (o PolicyDefinitionGroupOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionGroup) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o PolicyDefinitionGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionGroup) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PolicyDefinitionGroupOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionGroup) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o PolicyDefinitionGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyDefinitionGroup) string { return v.Name }).(pulumi.StringOutput)
}

type PolicyDefinitionGroupArrayOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyDefinitionGroup)(nil)).Elem()
}

func (o PolicyDefinitionGroupArrayOutput) ToPolicyDefinitionGroupArrayOutput() PolicyDefinitionGroupArrayOutput {
	return o
}

func (o PolicyDefinitionGroupArrayOutput) ToPolicyDefinitionGroupArrayOutputWithContext(ctx context.Context) PolicyDefinitionGroupArrayOutput {
	return o
}

func (o PolicyDefinitionGroupArrayOutput) Index(i pulumi.IntInput) PolicyDefinitionGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyDefinitionGroup {
		return vs[0].([]PolicyDefinitionGroup)[vs[1].(int)]
	}).(PolicyDefinitionGroupOutput)
}

type PolicyDefinitionGroupResponse struct {
	AdditionalMetadataId *string `pulumi:"additionalMetadataId"`
	Category             *string `pulumi:"category"`
	Description          *string `pulumi:"description"`
	DisplayName          *string `pulumi:"displayName"`
	Name                 string  `pulumi:"name"`
}

type PolicyDefinitionGroupResponseOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyDefinitionGroupResponse)(nil)).Elem()
}

func (o PolicyDefinitionGroupResponseOutput) ToPolicyDefinitionGroupResponseOutput() PolicyDefinitionGroupResponseOutput {
	return o
}

func (o PolicyDefinitionGroupResponseOutput) ToPolicyDefinitionGroupResponseOutputWithContext(ctx context.Context) PolicyDefinitionGroupResponseOutput {
	return o
}

func (o PolicyDefinitionGroupResponseOutput) AdditionalMetadataId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionGroupResponse) *string { return v.AdditionalMetadataId }).(pulumi.StringPtrOutput)
}

func (o PolicyDefinitionGroupResponseOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionGroupResponse) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o PolicyDefinitionGroupResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionGroupResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PolicyDefinitionGroupResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionGroupResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o PolicyDefinitionGroupResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyDefinitionGroupResponse) string { return v.Name }).(pulumi.StringOutput)
}

type PolicyDefinitionGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyDefinitionGroupResponse)(nil)).Elem()
}

func (o PolicyDefinitionGroupResponseArrayOutput) ToPolicyDefinitionGroupResponseArrayOutput() PolicyDefinitionGroupResponseArrayOutput {
	return o
}

func (o PolicyDefinitionGroupResponseArrayOutput) ToPolicyDefinitionGroupResponseArrayOutputWithContext(ctx context.Context) PolicyDefinitionGroupResponseArrayOutput {
	return o
}

func (o PolicyDefinitionGroupResponseArrayOutput) Index(i pulumi.IntInput) PolicyDefinitionGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyDefinitionGroupResponse {
		return vs[0].([]PolicyDefinitionGroupResponse)[vs[1].(int)]
	}).(PolicyDefinitionGroupResponseOutput)
}

type PolicyDefinitionReference struct {
	GroupNames                  []string                        `pulumi:"groupNames"`
	Parameters                  map[string]ParameterValuesValue `pulumi:"parameters"`
	PolicyDefinitionId          string                          `pulumi:"policyDefinitionId"`
	PolicyDefinitionReferenceId *string                         `pulumi:"policyDefinitionReferenceId"`
}





type PolicyDefinitionReferenceInput interface {
	pulumi.Input

	ToPolicyDefinitionReferenceOutput() PolicyDefinitionReferenceOutput
	ToPolicyDefinitionReferenceOutputWithContext(context.Context) PolicyDefinitionReferenceOutput
}

type PolicyDefinitionReferenceArgs struct {
	GroupNames                  pulumi.StringArrayInput      `pulumi:"groupNames"`
	Parameters                  ParameterValuesValueMapInput `pulumi:"parameters"`
	PolicyDefinitionId          pulumi.StringInput           `pulumi:"policyDefinitionId"`
	PolicyDefinitionReferenceId pulumi.StringPtrInput        `pulumi:"policyDefinitionReferenceId"`
}

func (PolicyDefinitionReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyDefinitionReference)(nil)).Elem()
}

func (i PolicyDefinitionReferenceArgs) ToPolicyDefinitionReferenceOutput() PolicyDefinitionReferenceOutput {
	return i.ToPolicyDefinitionReferenceOutputWithContext(context.Background())
}

func (i PolicyDefinitionReferenceArgs) ToPolicyDefinitionReferenceOutputWithContext(ctx context.Context) PolicyDefinitionReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyDefinitionReferenceOutput)
}





type PolicyDefinitionReferenceArrayInput interface {
	pulumi.Input

	ToPolicyDefinitionReferenceArrayOutput() PolicyDefinitionReferenceArrayOutput
	ToPolicyDefinitionReferenceArrayOutputWithContext(context.Context) PolicyDefinitionReferenceArrayOutput
}

type PolicyDefinitionReferenceArray []PolicyDefinitionReferenceInput

func (PolicyDefinitionReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyDefinitionReference)(nil)).Elem()
}

func (i PolicyDefinitionReferenceArray) ToPolicyDefinitionReferenceArrayOutput() PolicyDefinitionReferenceArrayOutput {
	return i.ToPolicyDefinitionReferenceArrayOutputWithContext(context.Background())
}

func (i PolicyDefinitionReferenceArray) ToPolicyDefinitionReferenceArrayOutputWithContext(ctx context.Context) PolicyDefinitionReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyDefinitionReferenceArrayOutput)
}

type PolicyDefinitionReferenceOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyDefinitionReference)(nil)).Elem()
}

func (o PolicyDefinitionReferenceOutput) ToPolicyDefinitionReferenceOutput() PolicyDefinitionReferenceOutput {
	return o
}

func (o PolicyDefinitionReferenceOutput) ToPolicyDefinitionReferenceOutputWithContext(ctx context.Context) PolicyDefinitionReferenceOutput {
	return o
}

func (o PolicyDefinitionReferenceOutput) GroupNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PolicyDefinitionReference) []string { return v.GroupNames }).(pulumi.StringArrayOutput)
}

func (o PolicyDefinitionReferenceOutput) Parameters() ParameterValuesValueMapOutput {
	return o.ApplyT(func(v PolicyDefinitionReference) map[string]ParameterValuesValue { return v.Parameters }).(ParameterValuesValueMapOutput)
}

func (o PolicyDefinitionReferenceOutput) PolicyDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyDefinitionReference) string { return v.PolicyDefinitionId }).(pulumi.StringOutput)
}

func (o PolicyDefinitionReferenceOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionReference) *string { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

type PolicyDefinitionReferenceArrayOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyDefinitionReference)(nil)).Elem()
}

func (o PolicyDefinitionReferenceArrayOutput) ToPolicyDefinitionReferenceArrayOutput() PolicyDefinitionReferenceArrayOutput {
	return o
}

func (o PolicyDefinitionReferenceArrayOutput) ToPolicyDefinitionReferenceArrayOutputWithContext(ctx context.Context) PolicyDefinitionReferenceArrayOutput {
	return o
}

func (o PolicyDefinitionReferenceArrayOutput) Index(i pulumi.IntInput) PolicyDefinitionReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyDefinitionReference {
		return vs[0].([]PolicyDefinitionReference)[vs[1].(int)]
	}).(PolicyDefinitionReferenceOutput)
}

type PolicyDefinitionReferenceResponse struct {
	GroupNames                  []string                                `pulumi:"groupNames"`
	Parameters                  map[string]ParameterValuesValueResponse `pulumi:"parameters"`
	PolicyDefinitionId          string                                  `pulumi:"policyDefinitionId"`
	PolicyDefinitionReferenceId *string                                 `pulumi:"policyDefinitionReferenceId"`
}

type PolicyDefinitionReferenceResponseOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyDefinitionReferenceResponse)(nil)).Elem()
}

func (o PolicyDefinitionReferenceResponseOutput) ToPolicyDefinitionReferenceResponseOutput() PolicyDefinitionReferenceResponseOutput {
	return o
}

func (o PolicyDefinitionReferenceResponseOutput) ToPolicyDefinitionReferenceResponseOutputWithContext(ctx context.Context) PolicyDefinitionReferenceResponseOutput {
	return o
}

func (o PolicyDefinitionReferenceResponseOutput) GroupNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PolicyDefinitionReferenceResponse) []string { return v.GroupNames }).(pulumi.StringArrayOutput)
}

func (o PolicyDefinitionReferenceResponseOutput) Parameters() ParameterValuesValueResponseMapOutput {
	return o.ApplyT(func(v PolicyDefinitionReferenceResponse) map[string]ParameterValuesValueResponse { return v.Parameters }).(ParameterValuesValueResponseMapOutput)
}

func (o PolicyDefinitionReferenceResponseOutput) PolicyDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyDefinitionReferenceResponse) string { return v.PolicyDefinitionId }).(pulumi.StringOutput)
}

func (o PolicyDefinitionReferenceResponseOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionReferenceResponse) *string { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

type PolicyDefinitionReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyDefinitionReferenceResponse)(nil)).Elem()
}

func (o PolicyDefinitionReferenceResponseArrayOutput) ToPolicyDefinitionReferenceResponseArrayOutput() PolicyDefinitionReferenceResponseArrayOutput {
	return o
}

func (o PolicyDefinitionReferenceResponseArrayOutput) ToPolicyDefinitionReferenceResponseArrayOutputWithContext(ctx context.Context) PolicyDefinitionReferenceResponseArrayOutput {
	return o
}

func (o PolicyDefinitionReferenceResponseArrayOutput) Index(i pulumi.IntInput) PolicyDefinitionReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyDefinitionReferenceResponse {
		return vs[0].([]PolicyDefinitionReferenceResponse)[vs[1].(int)]
	}).(PolicyDefinitionReferenceResponseOutput)
}

type PrincipalResponse struct {
	DisplayName *string `pulumi:"displayName"`
	Email       *string `pulumi:"email"`
	Id          *string `pulumi:"id"`
	Type        *string `pulumi:"type"`
}

type PrincipalResponseOutput struct{ *pulumi.OutputState }

func (PrincipalResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrincipalResponse)(nil)).Elem()
}

func (o PrincipalResponseOutput) ToPrincipalResponseOutput() PrincipalResponseOutput {
	return o
}

func (o PrincipalResponseOutput) ToPrincipalResponseOutputWithContext(ctx context.Context) PrincipalResponseOutput {
	return o
}

func (o PrincipalResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrincipalResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o PrincipalResponseOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrincipalResponse) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o PrincipalResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrincipalResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o PrincipalResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrincipalResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PrincipalResponsePtrOutput struct{ *pulumi.OutputState }

func (PrincipalResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrincipalResponse)(nil)).Elem()
}

func (o PrincipalResponsePtrOutput) ToPrincipalResponsePtrOutput() PrincipalResponsePtrOutput {
	return o
}

func (o PrincipalResponsePtrOutput) ToPrincipalResponsePtrOutputWithContext(ctx context.Context) PrincipalResponsePtrOutput {
	return o
}

func (o PrincipalResponsePtrOutput) Elem() PrincipalResponseOutput {
	return o.ApplyT(func(v *PrincipalResponse) PrincipalResponse {
		if v != nil {
			return *v
		}
		var ret PrincipalResponse
		return ret
	}).(PrincipalResponseOutput)
}

func (o PrincipalResponsePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrincipalResponse) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o PrincipalResponsePtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrincipalResponse) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o PrincipalResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrincipalResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o PrincipalResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrincipalResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkAssociationProperties struct {
	PrivateLink         *string `pulumi:"privateLink"`
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
}





type PrivateLinkAssociationPropertiesInput interface {
	pulumi.Input

	ToPrivateLinkAssociationPropertiesOutput() PrivateLinkAssociationPropertiesOutput
	ToPrivateLinkAssociationPropertiesOutputWithContext(context.Context) PrivateLinkAssociationPropertiesOutput
}

type PrivateLinkAssociationPropertiesArgs struct {
	PrivateLink         pulumi.StringPtrInput `pulumi:"privateLink"`
	PublicNetworkAccess pulumi.StringPtrInput `pulumi:"publicNetworkAccess"`
}

func (PrivateLinkAssociationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkAssociationProperties)(nil)).Elem()
}

func (i PrivateLinkAssociationPropertiesArgs) ToPrivateLinkAssociationPropertiesOutput() PrivateLinkAssociationPropertiesOutput {
	return i.ToPrivateLinkAssociationPropertiesOutputWithContext(context.Background())
}

func (i PrivateLinkAssociationPropertiesArgs) ToPrivateLinkAssociationPropertiesOutputWithContext(ctx context.Context) PrivateLinkAssociationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkAssociationPropertiesOutput)
}

func (i PrivateLinkAssociationPropertiesArgs) ToPrivateLinkAssociationPropertiesPtrOutput() PrivateLinkAssociationPropertiesPtrOutput {
	return i.ToPrivateLinkAssociationPropertiesPtrOutputWithContext(context.Background())
}

func (i PrivateLinkAssociationPropertiesArgs) ToPrivateLinkAssociationPropertiesPtrOutputWithContext(ctx context.Context) PrivateLinkAssociationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkAssociationPropertiesOutput).ToPrivateLinkAssociationPropertiesPtrOutputWithContext(ctx)
}









type PrivateLinkAssociationPropertiesPtrInput interface {
	pulumi.Input

	ToPrivateLinkAssociationPropertiesPtrOutput() PrivateLinkAssociationPropertiesPtrOutput
	ToPrivateLinkAssociationPropertiesPtrOutputWithContext(context.Context) PrivateLinkAssociationPropertiesPtrOutput
}

type privateLinkAssociationPropertiesPtrType PrivateLinkAssociationPropertiesArgs

func PrivateLinkAssociationPropertiesPtr(v *PrivateLinkAssociationPropertiesArgs) PrivateLinkAssociationPropertiesPtrInput {
	return (*privateLinkAssociationPropertiesPtrType)(v)
}

func (*privateLinkAssociationPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkAssociationProperties)(nil)).Elem()
}

func (i *privateLinkAssociationPropertiesPtrType) ToPrivateLinkAssociationPropertiesPtrOutput() PrivateLinkAssociationPropertiesPtrOutput {
	return i.ToPrivateLinkAssociationPropertiesPtrOutputWithContext(context.Background())
}

func (i *privateLinkAssociationPropertiesPtrType) ToPrivateLinkAssociationPropertiesPtrOutputWithContext(ctx context.Context) PrivateLinkAssociationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkAssociationPropertiesPtrOutput)
}

type PrivateLinkAssociationPropertiesOutput struct{ *pulumi.OutputState }

func (PrivateLinkAssociationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkAssociationProperties)(nil)).Elem()
}

func (o PrivateLinkAssociationPropertiesOutput) ToPrivateLinkAssociationPropertiesOutput() PrivateLinkAssociationPropertiesOutput {
	return o
}

func (o PrivateLinkAssociationPropertiesOutput) ToPrivateLinkAssociationPropertiesOutputWithContext(ctx context.Context) PrivateLinkAssociationPropertiesOutput {
	return o
}

func (o PrivateLinkAssociationPropertiesOutput) ToPrivateLinkAssociationPropertiesPtrOutput() PrivateLinkAssociationPropertiesPtrOutput {
	return o.ToPrivateLinkAssociationPropertiesPtrOutputWithContext(context.Background())
}

func (o PrivateLinkAssociationPropertiesOutput) ToPrivateLinkAssociationPropertiesPtrOutputWithContext(ctx context.Context) PrivateLinkAssociationPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkAssociationProperties) *PrivateLinkAssociationProperties {
		return &v
	}).(PrivateLinkAssociationPropertiesPtrOutput)
}

func (o PrivateLinkAssociationPropertiesOutput) PrivateLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkAssociationProperties) *string { return v.PrivateLink }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkAssociationPropertiesOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkAssociationProperties) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

type PrivateLinkAssociationPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkAssociationPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkAssociationProperties)(nil)).Elem()
}

func (o PrivateLinkAssociationPropertiesPtrOutput) ToPrivateLinkAssociationPropertiesPtrOutput() PrivateLinkAssociationPropertiesPtrOutput {
	return o
}

func (o PrivateLinkAssociationPropertiesPtrOutput) ToPrivateLinkAssociationPropertiesPtrOutputWithContext(ctx context.Context) PrivateLinkAssociationPropertiesPtrOutput {
	return o
}

func (o PrivateLinkAssociationPropertiesPtrOutput) Elem() PrivateLinkAssociationPropertiesOutput {
	return o.ApplyT(func(v *PrivateLinkAssociationProperties) PrivateLinkAssociationProperties {
		if v != nil {
			return *v
		}
		var ret PrivateLinkAssociationProperties
		return ret
	}).(PrivateLinkAssociationPropertiesOutput)
}

func (o PrivateLinkAssociationPropertiesPtrOutput) PrivateLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkAssociationProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrivateLink
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkAssociationPropertiesPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkAssociationProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkAssociationPropertiesExpandedResponse struct {
	PrivateLink         *string `pulumi:"privateLink"`
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	Scope               *string `pulumi:"scope"`
	TenantID            *string `pulumi:"tenantID"`
}

type PrivateLinkAssociationPropertiesExpandedResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkAssociationPropertiesExpandedResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkAssociationPropertiesExpandedResponse)(nil)).Elem()
}

func (o PrivateLinkAssociationPropertiesExpandedResponseOutput) ToPrivateLinkAssociationPropertiesExpandedResponseOutput() PrivateLinkAssociationPropertiesExpandedResponseOutput {
	return o
}

func (o PrivateLinkAssociationPropertiesExpandedResponseOutput) ToPrivateLinkAssociationPropertiesExpandedResponseOutputWithContext(ctx context.Context) PrivateLinkAssociationPropertiesExpandedResponseOutput {
	return o
}

func (o PrivateLinkAssociationPropertiesExpandedResponseOutput) PrivateLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkAssociationPropertiesExpandedResponse) *string { return v.PrivateLink }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkAssociationPropertiesExpandedResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkAssociationPropertiesExpandedResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkAssociationPropertiesExpandedResponseOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkAssociationPropertiesExpandedResponse) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkAssociationPropertiesExpandedResponseOutput) TenantID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkAssociationPropertiesExpandedResponse) *string { return v.TenantID }).(pulumi.StringPtrOutput)
}

type ResourceManagementPrivateLinkEndpointConnectionsResponse struct {
	PrivateEndpointConnections []string `pulumi:"privateEndpointConnections"`
}

type ResourceManagementPrivateLinkEndpointConnectionsResponseOutput struct{ *pulumi.OutputState }

func (ResourceManagementPrivateLinkEndpointConnectionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceManagementPrivateLinkEndpointConnectionsResponse)(nil)).Elem()
}

func (o ResourceManagementPrivateLinkEndpointConnectionsResponseOutput) ToResourceManagementPrivateLinkEndpointConnectionsResponseOutput() ResourceManagementPrivateLinkEndpointConnectionsResponseOutput {
	return o
}

func (o ResourceManagementPrivateLinkEndpointConnectionsResponseOutput) ToResourceManagementPrivateLinkEndpointConnectionsResponseOutputWithContext(ctx context.Context) ResourceManagementPrivateLinkEndpointConnectionsResponseOutput {
	return o
}

func (o ResourceManagementPrivateLinkEndpointConnectionsResponseOutput) PrivateEndpointConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceManagementPrivateLinkEndpointConnectionsResponse) []string {
		return v.PrivateEndpointConnections
	}).(pulumi.StringArrayOutput)
}

type RoleManagementPolicyApprovalRuleResponse struct {
	Id       *string                                 `pulumi:"id"`
	RuleType string                                  `pulumi:"ruleType"`
	Setting  *ApprovalSettingsResponse               `pulumi:"setting"`
	Target   *RoleManagementPolicyRuleTargetResponse `pulumi:"target"`
}

type RoleManagementPolicyAuthenticationContextRuleResponse struct {
	ClaimValue *string                                 `pulumi:"claimValue"`
	Id         *string                                 `pulumi:"id"`
	IsEnabled  *bool                                   `pulumi:"isEnabled"`
	RuleType   string                                  `pulumi:"ruleType"`
	Target     *RoleManagementPolicyRuleTargetResponse `pulumi:"target"`
}

type RoleManagementPolicyEnablementRuleResponse struct {
	EnabledRules []string                                `pulumi:"enabledRules"`
	Id           *string                                 `pulumi:"id"`
	RuleType     string                                  `pulumi:"ruleType"`
	Target       *RoleManagementPolicyRuleTargetResponse `pulumi:"target"`
}

type RoleManagementPolicyExpirationRuleResponse struct {
	Id                   *string                                 `pulumi:"id"`
	IsExpirationRequired *bool                                   `pulumi:"isExpirationRequired"`
	MaximumDuration      *string                                 `pulumi:"maximumDuration"`
	RuleType             string                                  `pulumi:"ruleType"`
	Target               *RoleManagementPolicyRuleTargetResponse `pulumi:"target"`
}

type RoleManagementPolicyNotificationRuleResponse struct {
	Id                         *string                                 `pulumi:"id"`
	IsDefaultRecipientsEnabled *bool                                   `pulumi:"isDefaultRecipientsEnabled"`
	NotificationLevel          *string                                 `pulumi:"notificationLevel"`
	NotificationRecipients     []string                                `pulumi:"notificationRecipients"`
	NotificationType           *string                                 `pulumi:"notificationType"`
	RecipientType              *string                                 `pulumi:"recipientType"`
	RuleType                   string                                  `pulumi:"ruleType"`
	Target                     *RoleManagementPolicyRuleTargetResponse `pulumi:"target"`
}

type RoleManagementPolicyRuleTargetResponse struct {
	Caller              *string  `pulumi:"caller"`
	EnforcedSettings    []string `pulumi:"enforcedSettings"`
	InheritableSettings []string `pulumi:"inheritableSettings"`
	Level               *string  `pulumi:"level"`
	Operations          []string `pulumi:"operations"`
	TargetObjects       []string `pulumi:"targetObjects"`
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

type UserSetResponse struct {
	Description *string `pulumi:"description"`
	Id          *string `pulumi:"id"`
	IsBackup    *bool   `pulumi:"isBackup"`
	UserType    *string `pulumi:"userType"`
}

func init() {
	pulumi.RegisterOutputType(AccessReviewHistoryInstanceOutput{})
	pulumi.RegisterOutputType(AccessReviewHistoryInstanceArrayOutput{})
	pulumi.RegisterOutputType(AccessReviewHistoryInstanceResponseOutput{})
	pulumi.RegisterOutputType(AccessReviewHistoryInstanceResponseArrayOutput{})
	pulumi.RegisterOutputType(AccessReviewInstanceOutput{})
	pulumi.RegisterOutputType(AccessReviewInstanceArrayOutput{})
	pulumi.RegisterOutputType(AccessReviewInstanceResponseOutput{})
	pulumi.RegisterOutputType(AccessReviewInstanceResponseArrayOutput{})
	pulumi.RegisterOutputType(AccessReviewReviewerOutput{})
	pulumi.RegisterOutputType(AccessReviewReviewerArrayOutput{})
	pulumi.RegisterOutputType(AccessReviewReviewerResponseOutput{})
	pulumi.RegisterOutputType(AccessReviewReviewerResponseArrayOutput{})
	pulumi.RegisterOutputType(AccessReviewScopeOutput{})
	pulumi.RegisterOutputType(AccessReviewScopeArrayOutput{})
	pulumi.RegisterOutputType(AccessReviewScopeResponseOutput{})
	pulumi.RegisterOutputType(AccessReviewScopeResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagementLockOwnerOutput{})
	pulumi.RegisterOutputType(ManagementLockOwnerArrayOutput{})
	pulumi.RegisterOutputType(ManagementLockOwnerResponseOutput{})
	pulumi.RegisterOutputType(ManagementLockOwnerResponseArrayOutput{})
	pulumi.RegisterOutputType(NonComplianceMessageOutput{})
	pulumi.RegisterOutputType(NonComplianceMessageArrayOutput{})
	pulumi.RegisterOutputType(NonComplianceMessageResponseOutput{})
	pulumi.RegisterOutputType(NonComplianceMessageResponseArrayOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionsValueOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionsValueMapOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionsValueMetadataOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionsValueMetadataPtrOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionsValueResponseOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionsValueResponseMapOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionsValueResponseMetadataOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionsValueResponseMetadataPtrOutput{})
	pulumi.RegisterOutputType(ParameterValuesValueOutput{})
	pulumi.RegisterOutputType(ParameterValuesValueMapOutput{})
	pulumi.RegisterOutputType(ParameterValuesValueResponseOutput{})
	pulumi.RegisterOutputType(ParameterValuesValueResponseMapOutput{})
	pulumi.RegisterOutputType(PermissionOutput{})
	pulumi.RegisterOutputType(PermissionArrayOutput{})
	pulumi.RegisterOutputType(PermissionResponseOutput{})
	pulumi.RegisterOutputType(PermissionResponseArrayOutput{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponsePolicyOutput{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponsePolicyPtrOutput{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponseRoleDefinitionOutput{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponseScopeOutput{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponseScopePtrOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionGroupOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionGroupArrayOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionGroupResponseOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionReferenceOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionReferenceArrayOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionReferenceResponseOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(PrincipalResponseOutput{})
	pulumi.RegisterOutputType(PrincipalResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkAssociationPropertiesOutput{})
	pulumi.RegisterOutputType(PrivateLinkAssociationPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkAssociationPropertiesExpandedResponseOutput{})
	pulumi.RegisterOutputType(ResourceManagementPrivateLinkEndpointConnectionsResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}

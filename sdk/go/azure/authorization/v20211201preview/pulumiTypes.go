


package v20211201preview

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
}

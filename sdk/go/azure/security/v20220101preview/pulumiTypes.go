


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GovernanceAssignmentAdditionalData struct {
	TicketLink   *string `pulumi:"ticketLink"`
	TicketNumber *int    `pulumi:"ticketNumber"`
	TicketStatus *string `pulumi:"ticketStatus"`
}





type GovernanceAssignmentAdditionalDataInput interface {
	pulumi.Input

	ToGovernanceAssignmentAdditionalDataOutput() GovernanceAssignmentAdditionalDataOutput
	ToGovernanceAssignmentAdditionalDataOutputWithContext(context.Context) GovernanceAssignmentAdditionalDataOutput
}

type GovernanceAssignmentAdditionalDataArgs struct {
	TicketLink   pulumi.StringPtrInput `pulumi:"ticketLink"`
	TicketNumber pulumi.IntPtrInput    `pulumi:"ticketNumber"`
	TicketStatus pulumi.StringPtrInput `pulumi:"ticketStatus"`
}

func (GovernanceAssignmentAdditionalDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GovernanceAssignmentAdditionalData)(nil)).Elem()
}

func (i GovernanceAssignmentAdditionalDataArgs) ToGovernanceAssignmentAdditionalDataOutput() GovernanceAssignmentAdditionalDataOutput {
	return i.ToGovernanceAssignmentAdditionalDataOutputWithContext(context.Background())
}

func (i GovernanceAssignmentAdditionalDataArgs) ToGovernanceAssignmentAdditionalDataOutputWithContext(ctx context.Context) GovernanceAssignmentAdditionalDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GovernanceAssignmentAdditionalDataOutput)
}

func (i GovernanceAssignmentAdditionalDataArgs) ToGovernanceAssignmentAdditionalDataPtrOutput() GovernanceAssignmentAdditionalDataPtrOutput {
	return i.ToGovernanceAssignmentAdditionalDataPtrOutputWithContext(context.Background())
}

func (i GovernanceAssignmentAdditionalDataArgs) ToGovernanceAssignmentAdditionalDataPtrOutputWithContext(ctx context.Context) GovernanceAssignmentAdditionalDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GovernanceAssignmentAdditionalDataOutput).ToGovernanceAssignmentAdditionalDataPtrOutputWithContext(ctx)
}









type GovernanceAssignmentAdditionalDataPtrInput interface {
	pulumi.Input

	ToGovernanceAssignmentAdditionalDataPtrOutput() GovernanceAssignmentAdditionalDataPtrOutput
	ToGovernanceAssignmentAdditionalDataPtrOutputWithContext(context.Context) GovernanceAssignmentAdditionalDataPtrOutput
}

type governanceAssignmentAdditionalDataPtrType GovernanceAssignmentAdditionalDataArgs

func GovernanceAssignmentAdditionalDataPtr(v *GovernanceAssignmentAdditionalDataArgs) GovernanceAssignmentAdditionalDataPtrInput {
	return (*governanceAssignmentAdditionalDataPtrType)(v)
}

func (*governanceAssignmentAdditionalDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GovernanceAssignmentAdditionalData)(nil)).Elem()
}

func (i *governanceAssignmentAdditionalDataPtrType) ToGovernanceAssignmentAdditionalDataPtrOutput() GovernanceAssignmentAdditionalDataPtrOutput {
	return i.ToGovernanceAssignmentAdditionalDataPtrOutputWithContext(context.Background())
}

func (i *governanceAssignmentAdditionalDataPtrType) ToGovernanceAssignmentAdditionalDataPtrOutputWithContext(ctx context.Context) GovernanceAssignmentAdditionalDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GovernanceAssignmentAdditionalDataPtrOutput)
}

type GovernanceAssignmentAdditionalDataOutput struct{ *pulumi.OutputState }

func (GovernanceAssignmentAdditionalDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GovernanceAssignmentAdditionalData)(nil)).Elem()
}

func (o GovernanceAssignmentAdditionalDataOutput) ToGovernanceAssignmentAdditionalDataOutput() GovernanceAssignmentAdditionalDataOutput {
	return o
}

func (o GovernanceAssignmentAdditionalDataOutput) ToGovernanceAssignmentAdditionalDataOutputWithContext(ctx context.Context) GovernanceAssignmentAdditionalDataOutput {
	return o
}

func (o GovernanceAssignmentAdditionalDataOutput) ToGovernanceAssignmentAdditionalDataPtrOutput() GovernanceAssignmentAdditionalDataPtrOutput {
	return o.ToGovernanceAssignmentAdditionalDataPtrOutputWithContext(context.Background())
}

func (o GovernanceAssignmentAdditionalDataOutput) ToGovernanceAssignmentAdditionalDataPtrOutputWithContext(ctx context.Context) GovernanceAssignmentAdditionalDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GovernanceAssignmentAdditionalData) *GovernanceAssignmentAdditionalData {
		return &v
	}).(GovernanceAssignmentAdditionalDataPtrOutput)
}

func (o GovernanceAssignmentAdditionalDataOutput) TicketLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GovernanceAssignmentAdditionalData) *string { return v.TicketLink }).(pulumi.StringPtrOutput)
}

func (o GovernanceAssignmentAdditionalDataOutput) TicketNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GovernanceAssignmentAdditionalData) *int { return v.TicketNumber }).(pulumi.IntPtrOutput)
}

func (o GovernanceAssignmentAdditionalDataOutput) TicketStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GovernanceAssignmentAdditionalData) *string { return v.TicketStatus }).(pulumi.StringPtrOutput)
}

type GovernanceAssignmentAdditionalDataPtrOutput struct{ *pulumi.OutputState }

func (GovernanceAssignmentAdditionalDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GovernanceAssignmentAdditionalData)(nil)).Elem()
}

func (o GovernanceAssignmentAdditionalDataPtrOutput) ToGovernanceAssignmentAdditionalDataPtrOutput() GovernanceAssignmentAdditionalDataPtrOutput {
	return o
}

func (o GovernanceAssignmentAdditionalDataPtrOutput) ToGovernanceAssignmentAdditionalDataPtrOutputWithContext(ctx context.Context) GovernanceAssignmentAdditionalDataPtrOutput {
	return o
}

func (o GovernanceAssignmentAdditionalDataPtrOutput) Elem() GovernanceAssignmentAdditionalDataOutput {
	return o.ApplyT(func(v *GovernanceAssignmentAdditionalData) GovernanceAssignmentAdditionalData {
		if v != nil {
			return *v
		}
		var ret GovernanceAssignmentAdditionalData
		return ret
	}).(GovernanceAssignmentAdditionalDataOutput)
}

func (o GovernanceAssignmentAdditionalDataPtrOutput) TicketLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GovernanceAssignmentAdditionalData) *string {
		if v == nil {
			return nil
		}
		return v.TicketLink
	}).(pulumi.StringPtrOutput)
}

func (o GovernanceAssignmentAdditionalDataPtrOutput) TicketNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GovernanceAssignmentAdditionalData) *int {
		if v == nil {
			return nil
		}
		return v.TicketNumber
	}).(pulumi.IntPtrOutput)
}

func (o GovernanceAssignmentAdditionalDataPtrOutput) TicketStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GovernanceAssignmentAdditionalData) *string {
		if v == nil {
			return nil
		}
		return v.TicketStatus
	}).(pulumi.StringPtrOutput)
}

type GovernanceAssignmentAdditionalDataResponse struct {
	TicketLink   *string `pulumi:"ticketLink"`
	TicketNumber *int    `pulumi:"ticketNumber"`
	TicketStatus *string `pulumi:"ticketStatus"`
}

type GovernanceAssignmentAdditionalDataResponseOutput struct{ *pulumi.OutputState }

func (GovernanceAssignmentAdditionalDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GovernanceAssignmentAdditionalDataResponse)(nil)).Elem()
}

func (o GovernanceAssignmentAdditionalDataResponseOutput) ToGovernanceAssignmentAdditionalDataResponseOutput() GovernanceAssignmentAdditionalDataResponseOutput {
	return o
}

func (o GovernanceAssignmentAdditionalDataResponseOutput) ToGovernanceAssignmentAdditionalDataResponseOutputWithContext(ctx context.Context) GovernanceAssignmentAdditionalDataResponseOutput {
	return o
}

func (o GovernanceAssignmentAdditionalDataResponseOutput) TicketLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GovernanceAssignmentAdditionalDataResponse) *string { return v.TicketLink }).(pulumi.StringPtrOutput)
}

func (o GovernanceAssignmentAdditionalDataResponseOutput) TicketNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GovernanceAssignmentAdditionalDataResponse) *int { return v.TicketNumber }).(pulumi.IntPtrOutput)
}

func (o GovernanceAssignmentAdditionalDataResponseOutput) TicketStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GovernanceAssignmentAdditionalDataResponse) *string { return v.TicketStatus }).(pulumi.StringPtrOutput)
}

type GovernanceAssignmentAdditionalDataResponsePtrOutput struct{ *pulumi.OutputState }

func (GovernanceAssignmentAdditionalDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GovernanceAssignmentAdditionalDataResponse)(nil)).Elem()
}

func (o GovernanceAssignmentAdditionalDataResponsePtrOutput) ToGovernanceAssignmentAdditionalDataResponsePtrOutput() GovernanceAssignmentAdditionalDataResponsePtrOutput {
	return o
}

func (o GovernanceAssignmentAdditionalDataResponsePtrOutput) ToGovernanceAssignmentAdditionalDataResponsePtrOutputWithContext(ctx context.Context) GovernanceAssignmentAdditionalDataResponsePtrOutput {
	return o
}

func (o GovernanceAssignmentAdditionalDataResponsePtrOutput) Elem() GovernanceAssignmentAdditionalDataResponseOutput {
	return o.ApplyT(func(v *GovernanceAssignmentAdditionalDataResponse) GovernanceAssignmentAdditionalDataResponse {
		if v != nil {
			return *v
		}
		var ret GovernanceAssignmentAdditionalDataResponse
		return ret
	}).(GovernanceAssignmentAdditionalDataResponseOutput)
}

func (o GovernanceAssignmentAdditionalDataResponsePtrOutput) TicketLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GovernanceAssignmentAdditionalDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.TicketLink
	}).(pulumi.StringPtrOutput)
}

func (o GovernanceAssignmentAdditionalDataResponsePtrOutput) TicketNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GovernanceAssignmentAdditionalDataResponse) *int {
		if v == nil {
			return nil
		}
		return v.TicketNumber
	}).(pulumi.IntPtrOutput)
}

func (o GovernanceAssignmentAdditionalDataResponsePtrOutput) TicketStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GovernanceAssignmentAdditionalDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.TicketStatus
	}).(pulumi.StringPtrOutput)
}

type GovernanceEmailNotification struct {
	DisableManagerEmailNotification *bool `pulumi:"disableManagerEmailNotification"`
	DisableOwnerEmailNotification   *bool `pulumi:"disableOwnerEmailNotification"`
}





type GovernanceEmailNotificationInput interface {
	pulumi.Input

	ToGovernanceEmailNotificationOutput() GovernanceEmailNotificationOutput
	ToGovernanceEmailNotificationOutputWithContext(context.Context) GovernanceEmailNotificationOutput
}

type GovernanceEmailNotificationArgs struct {
	DisableManagerEmailNotification pulumi.BoolPtrInput `pulumi:"disableManagerEmailNotification"`
	DisableOwnerEmailNotification   pulumi.BoolPtrInput `pulumi:"disableOwnerEmailNotification"`
}

func (GovernanceEmailNotificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GovernanceEmailNotification)(nil)).Elem()
}

func (i GovernanceEmailNotificationArgs) ToGovernanceEmailNotificationOutput() GovernanceEmailNotificationOutput {
	return i.ToGovernanceEmailNotificationOutputWithContext(context.Background())
}

func (i GovernanceEmailNotificationArgs) ToGovernanceEmailNotificationOutputWithContext(ctx context.Context) GovernanceEmailNotificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GovernanceEmailNotificationOutput)
}

func (i GovernanceEmailNotificationArgs) ToGovernanceEmailNotificationPtrOutput() GovernanceEmailNotificationPtrOutput {
	return i.ToGovernanceEmailNotificationPtrOutputWithContext(context.Background())
}

func (i GovernanceEmailNotificationArgs) ToGovernanceEmailNotificationPtrOutputWithContext(ctx context.Context) GovernanceEmailNotificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GovernanceEmailNotificationOutput).ToGovernanceEmailNotificationPtrOutputWithContext(ctx)
}









type GovernanceEmailNotificationPtrInput interface {
	pulumi.Input

	ToGovernanceEmailNotificationPtrOutput() GovernanceEmailNotificationPtrOutput
	ToGovernanceEmailNotificationPtrOutputWithContext(context.Context) GovernanceEmailNotificationPtrOutput
}

type governanceEmailNotificationPtrType GovernanceEmailNotificationArgs

func GovernanceEmailNotificationPtr(v *GovernanceEmailNotificationArgs) GovernanceEmailNotificationPtrInput {
	return (*governanceEmailNotificationPtrType)(v)
}

func (*governanceEmailNotificationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GovernanceEmailNotification)(nil)).Elem()
}

func (i *governanceEmailNotificationPtrType) ToGovernanceEmailNotificationPtrOutput() GovernanceEmailNotificationPtrOutput {
	return i.ToGovernanceEmailNotificationPtrOutputWithContext(context.Background())
}

func (i *governanceEmailNotificationPtrType) ToGovernanceEmailNotificationPtrOutputWithContext(ctx context.Context) GovernanceEmailNotificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GovernanceEmailNotificationPtrOutput)
}

type GovernanceEmailNotificationOutput struct{ *pulumi.OutputState }

func (GovernanceEmailNotificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GovernanceEmailNotification)(nil)).Elem()
}

func (o GovernanceEmailNotificationOutput) ToGovernanceEmailNotificationOutput() GovernanceEmailNotificationOutput {
	return o
}

func (o GovernanceEmailNotificationOutput) ToGovernanceEmailNotificationOutputWithContext(ctx context.Context) GovernanceEmailNotificationOutput {
	return o
}

func (o GovernanceEmailNotificationOutput) ToGovernanceEmailNotificationPtrOutput() GovernanceEmailNotificationPtrOutput {
	return o.ToGovernanceEmailNotificationPtrOutputWithContext(context.Background())
}

func (o GovernanceEmailNotificationOutput) ToGovernanceEmailNotificationPtrOutputWithContext(ctx context.Context) GovernanceEmailNotificationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GovernanceEmailNotification) *GovernanceEmailNotification {
		return &v
	}).(GovernanceEmailNotificationPtrOutput)
}

func (o GovernanceEmailNotificationOutput) DisableManagerEmailNotification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GovernanceEmailNotification) *bool { return v.DisableManagerEmailNotification }).(pulumi.BoolPtrOutput)
}

func (o GovernanceEmailNotificationOutput) DisableOwnerEmailNotification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GovernanceEmailNotification) *bool { return v.DisableOwnerEmailNotification }).(pulumi.BoolPtrOutput)
}

type GovernanceEmailNotificationPtrOutput struct{ *pulumi.OutputState }

func (GovernanceEmailNotificationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GovernanceEmailNotification)(nil)).Elem()
}

func (o GovernanceEmailNotificationPtrOutput) ToGovernanceEmailNotificationPtrOutput() GovernanceEmailNotificationPtrOutput {
	return o
}

func (o GovernanceEmailNotificationPtrOutput) ToGovernanceEmailNotificationPtrOutputWithContext(ctx context.Context) GovernanceEmailNotificationPtrOutput {
	return o
}

func (o GovernanceEmailNotificationPtrOutput) Elem() GovernanceEmailNotificationOutput {
	return o.ApplyT(func(v *GovernanceEmailNotification) GovernanceEmailNotification {
		if v != nil {
			return *v
		}
		var ret GovernanceEmailNotification
		return ret
	}).(GovernanceEmailNotificationOutput)
}

func (o GovernanceEmailNotificationPtrOutput) DisableManagerEmailNotification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GovernanceEmailNotification) *bool {
		if v == nil {
			return nil
		}
		return v.DisableManagerEmailNotification
	}).(pulumi.BoolPtrOutput)
}

func (o GovernanceEmailNotificationPtrOutput) DisableOwnerEmailNotification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GovernanceEmailNotification) *bool {
		if v == nil {
			return nil
		}
		return v.DisableOwnerEmailNotification
	}).(pulumi.BoolPtrOutput)
}

type GovernanceEmailNotificationResponse struct {
	DisableManagerEmailNotification *bool `pulumi:"disableManagerEmailNotification"`
	DisableOwnerEmailNotification   *bool `pulumi:"disableOwnerEmailNotification"`
}

type GovernanceEmailNotificationResponseOutput struct{ *pulumi.OutputState }

func (GovernanceEmailNotificationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GovernanceEmailNotificationResponse)(nil)).Elem()
}

func (o GovernanceEmailNotificationResponseOutput) ToGovernanceEmailNotificationResponseOutput() GovernanceEmailNotificationResponseOutput {
	return o
}

func (o GovernanceEmailNotificationResponseOutput) ToGovernanceEmailNotificationResponseOutputWithContext(ctx context.Context) GovernanceEmailNotificationResponseOutput {
	return o
}

func (o GovernanceEmailNotificationResponseOutput) DisableManagerEmailNotification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GovernanceEmailNotificationResponse) *bool { return v.DisableManagerEmailNotification }).(pulumi.BoolPtrOutput)
}

func (o GovernanceEmailNotificationResponseOutput) DisableOwnerEmailNotification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GovernanceEmailNotificationResponse) *bool { return v.DisableOwnerEmailNotification }).(pulumi.BoolPtrOutput)
}

type GovernanceEmailNotificationResponsePtrOutput struct{ *pulumi.OutputState }

func (GovernanceEmailNotificationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GovernanceEmailNotificationResponse)(nil)).Elem()
}

func (o GovernanceEmailNotificationResponsePtrOutput) ToGovernanceEmailNotificationResponsePtrOutput() GovernanceEmailNotificationResponsePtrOutput {
	return o
}

func (o GovernanceEmailNotificationResponsePtrOutput) ToGovernanceEmailNotificationResponsePtrOutputWithContext(ctx context.Context) GovernanceEmailNotificationResponsePtrOutput {
	return o
}

func (o GovernanceEmailNotificationResponsePtrOutput) Elem() GovernanceEmailNotificationResponseOutput {
	return o.ApplyT(func(v *GovernanceEmailNotificationResponse) GovernanceEmailNotificationResponse {
		if v != nil {
			return *v
		}
		var ret GovernanceEmailNotificationResponse
		return ret
	}).(GovernanceEmailNotificationResponseOutput)
}

func (o GovernanceEmailNotificationResponsePtrOutput) DisableManagerEmailNotification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GovernanceEmailNotificationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DisableManagerEmailNotification
	}).(pulumi.BoolPtrOutput)
}

func (o GovernanceEmailNotificationResponsePtrOutput) DisableOwnerEmailNotification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GovernanceEmailNotificationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DisableOwnerEmailNotification
	}).(pulumi.BoolPtrOutput)
}

type GovernanceRuleEmailNotification struct {
	DisableManagerEmailNotification *bool `pulumi:"disableManagerEmailNotification"`
	DisableOwnerEmailNotification   *bool `pulumi:"disableOwnerEmailNotification"`
}





type GovernanceRuleEmailNotificationInput interface {
	pulumi.Input

	ToGovernanceRuleEmailNotificationOutput() GovernanceRuleEmailNotificationOutput
	ToGovernanceRuleEmailNotificationOutputWithContext(context.Context) GovernanceRuleEmailNotificationOutput
}

type GovernanceRuleEmailNotificationArgs struct {
	DisableManagerEmailNotification pulumi.BoolPtrInput `pulumi:"disableManagerEmailNotification"`
	DisableOwnerEmailNotification   pulumi.BoolPtrInput `pulumi:"disableOwnerEmailNotification"`
}

func (GovernanceRuleEmailNotificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GovernanceRuleEmailNotification)(nil)).Elem()
}

func (i GovernanceRuleEmailNotificationArgs) ToGovernanceRuleEmailNotificationOutput() GovernanceRuleEmailNotificationOutput {
	return i.ToGovernanceRuleEmailNotificationOutputWithContext(context.Background())
}

func (i GovernanceRuleEmailNotificationArgs) ToGovernanceRuleEmailNotificationOutputWithContext(ctx context.Context) GovernanceRuleEmailNotificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GovernanceRuleEmailNotificationOutput)
}

func (i GovernanceRuleEmailNotificationArgs) ToGovernanceRuleEmailNotificationPtrOutput() GovernanceRuleEmailNotificationPtrOutput {
	return i.ToGovernanceRuleEmailNotificationPtrOutputWithContext(context.Background())
}

func (i GovernanceRuleEmailNotificationArgs) ToGovernanceRuleEmailNotificationPtrOutputWithContext(ctx context.Context) GovernanceRuleEmailNotificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GovernanceRuleEmailNotificationOutput).ToGovernanceRuleEmailNotificationPtrOutputWithContext(ctx)
}









type GovernanceRuleEmailNotificationPtrInput interface {
	pulumi.Input

	ToGovernanceRuleEmailNotificationPtrOutput() GovernanceRuleEmailNotificationPtrOutput
	ToGovernanceRuleEmailNotificationPtrOutputWithContext(context.Context) GovernanceRuleEmailNotificationPtrOutput
}

type governanceRuleEmailNotificationPtrType GovernanceRuleEmailNotificationArgs

func GovernanceRuleEmailNotificationPtr(v *GovernanceRuleEmailNotificationArgs) GovernanceRuleEmailNotificationPtrInput {
	return (*governanceRuleEmailNotificationPtrType)(v)
}

func (*governanceRuleEmailNotificationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GovernanceRuleEmailNotification)(nil)).Elem()
}

func (i *governanceRuleEmailNotificationPtrType) ToGovernanceRuleEmailNotificationPtrOutput() GovernanceRuleEmailNotificationPtrOutput {
	return i.ToGovernanceRuleEmailNotificationPtrOutputWithContext(context.Background())
}

func (i *governanceRuleEmailNotificationPtrType) ToGovernanceRuleEmailNotificationPtrOutputWithContext(ctx context.Context) GovernanceRuleEmailNotificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GovernanceRuleEmailNotificationPtrOutput)
}

type GovernanceRuleEmailNotificationOutput struct{ *pulumi.OutputState }

func (GovernanceRuleEmailNotificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GovernanceRuleEmailNotification)(nil)).Elem()
}

func (o GovernanceRuleEmailNotificationOutput) ToGovernanceRuleEmailNotificationOutput() GovernanceRuleEmailNotificationOutput {
	return o
}

func (o GovernanceRuleEmailNotificationOutput) ToGovernanceRuleEmailNotificationOutputWithContext(ctx context.Context) GovernanceRuleEmailNotificationOutput {
	return o
}

func (o GovernanceRuleEmailNotificationOutput) ToGovernanceRuleEmailNotificationPtrOutput() GovernanceRuleEmailNotificationPtrOutput {
	return o.ToGovernanceRuleEmailNotificationPtrOutputWithContext(context.Background())
}

func (o GovernanceRuleEmailNotificationOutput) ToGovernanceRuleEmailNotificationPtrOutputWithContext(ctx context.Context) GovernanceRuleEmailNotificationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GovernanceRuleEmailNotification) *GovernanceRuleEmailNotification {
		return &v
	}).(GovernanceRuleEmailNotificationPtrOutput)
}

func (o GovernanceRuleEmailNotificationOutput) DisableManagerEmailNotification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GovernanceRuleEmailNotification) *bool { return v.DisableManagerEmailNotification }).(pulumi.BoolPtrOutput)
}

func (o GovernanceRuleEmailNotificationOutput) DisableOwnerEmailNotification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GovernanceRuleEmailNotification) *bool { return v.DisableOwnerEmailNotification }).(pulumi.BoolPtrOutput)
}

type GovernanceRuleEmailNotificationPtrOutput struct{ *pulumi.OutputState }

func (GovernanceRuleEmailNotificationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GovernanceRuleEmailNotification)(nil)).Elem()
}

func (o GovernanceRuleEmailNotificationPtrOutput) ToGovernanceRuleEmailNotificationPtrOutput() GovernanceRuleEmailNotificationPtrOutput {
	return o
}

func (o GovernanceRuleEmailNotificationPtrOutput) ToGovernanceRuleEmailNotificationPtrOutputWithContext(ctx context.Context) GovernanceRuleEmailNotificationPtrOutput {
	return o
}

func (o GovernanceRuleEmailNotificationPtrOutput) Elem() GovernanceRuleEmailNotificationOutput {
	return o.ApplyT(func(v *GovernanceRuleEmailNotification) GovernanceRuleEmailNotification {
		if v != nil {
			return *v
		}
		var ret GovernanceRuleEmailNotification
		return ret
	}).(GovernanceRuleEmailNotificationOutput)
}

func (o GovernanceRuleEmailNotificationPtrOutput) DisableManagerEmailNotification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GovernanceRuleEmailNotification) *bool {
		if v == nil {
			return nil
		}
		return v.DisableManagerEmailNotification
	}).(pulumi.BoolPtrOutput)
}

func (o GovernanceRuleEmailNotificationPtrOutput) DisableOwnerEmailNotification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GovernanceRuleEmailNotification) *bool {
		if v == nil {
			return nil
		}
		return v.DisableOwnerEmailNotification
	}).(pulumi.BoolPtrOutput)
}

type GovernanceRuleEmailNotificationResponse struct {
	DisableManagerEmailNotification *bool `pulumi:"disableManagerEmailNotification"`
	DisableOwnerEmailNotification   *bool `pulumi:"disableOwnerEmailNotification"`
}

type GovernanceRuleEmailNotificationResponseOutput struct{ *pulumi.OutputState }

func (GovernanceRuleEmailNotificationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GovernanceRuleEmailNotificationResponse)(nil)).Elem()
}

func (o GovernanceRuleEmailNotificationResponseOutput) ToGovernanceRuleEmailNotificationResponseOutput() GovernanceRuleEmailNotificationResponseOutput {
	return o
}

func (o GovernanceRuleEmailNotificationResponseOutput) ToGovernanceRuleEmailNotificationResponseOutputWithContext(ctx context.Context) GovernanceRuleEmailNotificationResponseOutput {
	return o
}

func (o GovernanceRuleEmailNotificationResponseOutput) DisableManagerEmailNotification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GovernanceRuleEmailNotificationResponse) *bool { return v.DisableManagerEmailNotification }).(pulumi.BoolPtrOutput)
}

func (o GovernanceRuleEmailNotificationResponseOutput) DisableOwnerEmailNotification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GovernanceRuleEmailNotificationResponse) *bool { return v.DisableOwnerEmailNotification }).(pulumi.BoolPtrOutput)
}

type GovernanceRuleEmailNotificationResponsePtrOutput struct{ *pulumi.OutputState }

func (GovernanceRuleEmailNotificationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GovernanceRuleEmailNotificationResponse)(nil)).Elem()
}

func (o GovernanceRuleEmailNotificationResponsePtrOutput) ToGovernanceRuleEmailNotificationResponsePtrOutput() GovernanceRuleEmailNotificationResponsePtrOutput {
	return o
}

func (o GovernanceRuleEmailNotificationResponsePtrOutput) ToGovernanceRuleEmailNotificationResponsePtrOutputWithContext(ctx context.Context) GovernanceRuleEmailNotificationResponsePtrOutput {
	return o
}

func (o GovernanceRuleEmailNotificationResponsePtrOutput) Elem() GovernanceRuleEmailNotificationResponseOutput {
	return o.ApplyT(func(v *GovernanceRuleEmailNotificationResponse) GovernanceRuleEmailNotificationResponse {
		if v != nil {
			return *v
		}
		var ret GovernanceRuleEmailNotificationResponse
		return ret
	}).(GovernanceRuleEmailNotificationResponseOutput)
}

func (o GovernanceRuleEmailNotificationResponsePtrOutput) DisableManagerEmailNotification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GovernanceRuleEmailNotificationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DisableManagerEmailNotification
	}).(pulumi.BoolPtrOutput)
}

func (o GovernanceRuleEmailNotificationResponsePtrOutput) DisableOwnerEmailNotification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GovernanceRuleEmailNotificationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DisableOwnerEmailNotification
	}).(pulumi.BoolPtrOutput)
}

type GovernanceRuleMetadataResponse struct {
	CreatedBy string `pulumi:"createdBy"`
	CreatedOn string `pulumi:"createdOn"`
	UpdatedBy string `pulumi:"updatedBy"`
	UpdatedOn string `pulumi:"updatedOn"`
}

type GovernanceRuleMetadataResponseOutput struct{ *pulumi.OutputState }

func (GovernanceRuleMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GovernanceRuleMetadataResponse)(nil)).Elem()
}

func (o GovernanceRuleMetadataResponseOutput) ToGovernanceRuleMetadataResponseOutput() GovernanceRuleMetadataResponseOutput {
	return o
}

func (o GovernanceRuleMetadataResponseOutput) ToGovernanceRuleMetadataResponseOutputWithContext(ctx context.Context) GovernanceRuleMetadataResponseOutput {
	return o
}

func (o GovernanceRuleMetadataResponseOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v GovernanceRuleMetadataResponse) string { return v.CreatedBy }).(pulumi.StringOutput)
}

func (o GovernanceRuleMetadataResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v GovernanceRuleMetadataResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o GovernanceRuleMetadataResponseOutput) UpdatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v GovernanceRuleMetadataResponse) string { return v.UpdatedBy }).(pulumi.StringOutput)
}

func (o GovernanceRuleMetadataResponseOutput) UpdatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v GovernanceRuleMetadataResponse) string { return v.UpdatedOn }).(pulumi.StringOutput)
}

type GovernanceRuleMetadataResponsePtrOutput struct{ *pulumi.OutputState }

func (GovernanceRuleMetadataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GovernanceRuleMetadataResponse)(nil)).Elem()
}

func (o GovernanceRuleMetadataResponsePtrOutput) ToGovernanceRuleMetadataResponsePtrOutput() GovernanceRuleMetadataResponsePtrOutput {
	return o
}

func (o GovernanceRuleMetadataResponsePtrOutput) ToGovernanceRuleMetadataResponsePtrOutputWithContext(ctx context.Context) GovernanceRuleMetadataResponsePtrOutput {
	return o
}

func (o GovernanceRuleMetadataResponsePtrOutput) Elem() GovernanceRuleMetadataResponseOutput {
	return o.ApplyT(func(v *GovernanceRuleMetadataResponse) GovernanceRuleMetadataResponse {
		if v != nil {
			return *v
		}
		var ret GovernanceRuleMetadataResponse
		return ret
	}).(GovernanceRuleMetadataResponseOutput)
}

func (o GovernanceRuleMetadataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GovernanceRuleMetadataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o GovernanceRuleMetadataResponsePtrOutput) CreatedOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GovernanceRuleMetadataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedOn
	}).(pulumi.StringPtrOutput)
}

func (o GovernanceRuleMetadataResponsePtrOutput) UpdatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GovernanceRuleMetadataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpdatedBy
	}).(pulumi.StringPtrOutput)
}

func (o GovernanceRuleMetadataResponsePtrOutput) UpdatedOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GovernanceRuleMetadataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpdatedOn
	}).(pulumi.StringPtrOutput)
}

type GovernanceRuleOwnerSource struct {
	Type  *string `pulumi:"type"`
	Value *string `pulumi:"value"`
}





type GovernanceRuleOwnerSourceInput interface {
	pulumi.Input

	ToGovernanceRuleOwnerSourceOutput() GovernanceRuleOwnerSourceOutput
	ToGovernanceRuleOwnerSourceOutputWithContext(context.Context) GovernanceRuleOwnerSourceOutput
}

type GovernanceRuleOwnerSourceArgs struct {
	Type  pulumi.StringPtrInput `pulumi:"type"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (GovernanceRuleOwnerSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GovernanceRuleOwnerSource)(nil)).Elem()
}

func (i GovernanceRuleOwnerSourceArgs) ToGovernanceRuleOwnerSourceOutput() GovernanceRuleOwnerSourceOutput {
	return i.ToGovernanceRuleOwnerSourceOutputWithContext(context.Background())
}

func (i GovernanceRuleOwnerSourceArgs) ToGovernanceRuleOwnerSourceOutputWithContext(ctx context.Context) GovernanceRuleOwnerSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GovernanceRuleOwnerSourceOutput)
}

type GovernanceRuleOwnerSourceOutput struct{ *pulumi.OutputState }

func (GovernanceRuleOwnerSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GovernanceRuleOwnerSource)(nil)).Elem()
}

func (o GovernanceRuleOwnerSourceOutput) ToGovernanceRuleOwnerSourceOutput() GovernanceRuleOwnerSourceOutput {
	return o
}

func (o GovernanceRuleOwnerSourceOutput) ToGovernanceRuleOwnerSourceOutputWithContext(ctx context.Context) GovernanceRuleOwnerSourceOutput {
	return o
}

func (o GovernanceRuleOwnerSourceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GovernanceRuleOwnerSource) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o GovernanceRuleOwnerSourceOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GovernanceRuleOwnerSource) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type GovernanceRuleOwnerSourceResponse struct {
	Type  *string `pulumi:"type"`
	Value *string `pulumi:"value"`
}

type GovernanceRuleOwnerSourceResponseOutput struct{ *pulumi.OutputState }

func (GovernanceRuleOwnerSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GovernanceRuleOwnerSourceResponse)(nil)).Elem()
}

func (o GovernanceRuleOwnerSourceResponseOutput) ToGovernanceRuleOwnerSourceResponseOutput() GovernanceRuleOwnerSourceResponseOutput {
	return o
}

func (o GovernanceRuleOwnerSourceResponseOutput) ToGovernanceRuleOwnerSourceResponseOutputWithContext(ctx context.Context) GovernanceRuleOwnerSourceResponseOutput {
	return o
}

func (o GovernanceRuleOwnerSourceResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GovernanceRuleOwnerSourceResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o GovernanceRuleOwnerSourceResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GovernanceRuleOwnerSourceResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type RemediationEta struct {
	Eta           string `pulumi:"eta"`
	Justification string `pulumi:"justification"`
}





type RemediationEtaInput interface {
	pulumi.Input

	ToRemediationEtaOutput() RemediationEtaOutput
	ToRemediationEtaOutputWithContext(context.Context) RemediationEtaOutput
}

type RemediationEtaArgs struct {
	Eta           pulumi.StringInput `pulumi:"eta"`
	Justification pulumi.StringInput `pulumi:"justification"`
}

func (RemediationEtaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationEta)(nil)).Elem()
}

func (i RemediationEtaArgs) ToRemediationEtaOutput() RemediationEtaOutput {
	return i.ToRemediationEtaOutputWithContext(context.Background())
}

func (i RemediationEtaArgs) ToRemediationEtaOutputWithContext(ctx context.Context) RemediationEtaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationEtaOutput)
}

func (i RemediationEtaArgs) ToRemediationEtaPtrOutput() RemediationEtaPtrOutput {
	return i.ToRemediationEtaPtrOutputWithContext(context.Background())
}

func (i RemediationEtaArgs) ToRemediationEtaPtrOutputWithContext(ctx context.Context) RemediationEtaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationEtaOutput).ToRemediationEtaPtrOutputWithContext(ctx)
}









type RemediationEtaPtrInput interface {
	pulumi.Input

	ToRemediationEtaPtrOutput() RemediationEtaPtrOutput
	ToRemediationEtaPtrOutputWithContext(context.Context) RemediationEtaPtrOutput
}

type remediationEtaPtrType RemediationEtaArgs

func RemediationEtaPtr(v *RemediationEtaArgs) RemediationEtaPtrInput {
	return (*remediationEtaPtrType)(v)
}

func (*remediationEtaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationEta)(nil)).Elem()
}

func (i *remediationEtaPtrType) ToRemediationEtaPtrOutput() RemediationEtaPtrOutput {
	return i.ToRemediationEtaPtrOutputWithContext(context.Background())
}

func (i *remediationEtaPtrType) ToRemediationEtaPtrOutputWithContext(ctx context.Context) RemediationEtaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationEtaPtrOutput)
}

type RemediationEtaOutput struct{ *pulumi.OutputState }

func (RemediationEtaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationEta)(nil)).Elem()
}

func (o RemediationEtaOutput) ToRemediationEtaOutput() RemediationEtaOutput {
	return o
}

func (o RemediationEtaOutput) ToRemediationEtaOutputWithContext(ctx context.Context) RemediationEtaOutput {
	return o
}

func (o RemediationEtaOutput) ToRemediationEtaPtrOutput() RemediationEtaPtrOutput {
	return o.ToRemediationEtaPtrOutputWithContext(context.Background())
}

func (o RemediationEtaOutput) ToRemediationEtaPtrOutputWithContext(ctx context.Context) RemediationEtaPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RemediationEta) *RemediationEta {
		return &v
	}).(RemediationEtaPtrOutput)
}

func (o RemediationEtaOutput) Eta() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationEta) string { return v.Eta }).(pulumi.StringOutput)
}

func (o RemediationEtaOutput) Justification() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationEta) string { return v.Justification }).(pulumi.StringOutput)
}

type RemediationEtaPtrOutput struct{ *pulumi.OutputState }

func (RemediationEtaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationEta)(nil)).Elem()
}

func (o RemediationEtaPtrOutput) ToRemediationEtaPtrOutput() RemediationEtaPtrOutput {
	return o
}

func (o RemediationEtaPtrOutput) ToRemediationEtaPtrOutputWithContext(ctx context.Context) RemediationEtaPtrOutput {
	return o
}

func (o RemediationEtaPtrOutput) Elem() RemediationEtaOutput {
	return o.ApplyT(func(v *RemediationEta) RemediationEta {
		if v != nil {
			return *v
		}
		var ret RemediationEta
		return ret
	}).(RemediationEtaOutput)
}

func (o RemediationEtaPtrOutput) Eta() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemediationEta) *string {
		if v == nil {
			return nil
		}
		return &v.Eta
	}).(pulumi.StringPtrOutput)
}

func (o RemediationEtaPtrOutput) Justification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemediationEta) *string {
		if v == nil {
			return nil
		}
		return &v.Justification
	}).(pulumi.StringPtrOutput)
}

type RemediationEtaResponse struct {
	Eta           string `pulumi:"eta"`
	Justification string `pulumi:"justification"`
}

type RemediationEtaResponseOutput struct{ *pulumi.OutputState }

func (RemediationEtaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationEtaResponse)(nil)).Elem()
}

func (o RemediationEtaResponseOutput) ToRemediationEtaResponseOutput() RemediationEtaResponseOutput {
	return o
}

func (o RemediationEtaResponseOutput) ToRemediationEtaResponseOutputWithContext(ctx context.Context) RemediationEtaResponseOutput {
	return o
}

func (o RemediationEtaResponseOutput) Eta() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationEtaResponse) string { return v.Eta }).(pulumi.StringOutput)
}

func (o RemediationEtaResponseOutput) Justification() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationEtaResponse) string { return v.Justification }).(pulumi.StringOutput)
}

type RemediationEtaResponsePtrOutput struct{ *pulumi.OutputState }

func (RemediationEtaResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationEtaResponse)(nil)).Elem()
}

func (o RemediationEtaResponsePtrOutput) ToRemediationEtaResponsePtrOutput() RemediationEtaResponsePtrOutput {
	return o
}

func (o RemediationEtaResponsePtrOutput) ToRemediationEtaResponsePtrOutputWithContext(ctx context.Context) RemediationEtaResponsePtrOutput {
	return o
}

func (o RemediationEtaResponsePtrOutput) Elem() RemediationEtaResponseOutput {
	return o.ApplyT(func(v *RemediationEtaResponse) RemediationEtaResponse {
		if v != nil {
			return *v
		}
		var ret RemediationEtaResponse
		return ret
	}).(RemediationEtaResponseOutput)
}

func (o RemediationEtaResponsePtrOutput) Eta() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemediationEtaResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Eta
	}).(pulumi.StringPtrOutput)
}

func (o RemediationEtaResponsePtrOutput) Justification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemediationEtaResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Justification
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GovernanceAssignmentAdditionalDataOutput{})
	pulumi.RegisterOutputType(GovernanceAssignmentAdditionalDataPtrOutput{})
	pulumi.RegisterOutputType(GovernanceAssignmentAdditionalDataResponseOutput{})
	pulumi.RegisterOutputType(GovernanceAssignmentAdditionalDataResponsePtrOutput{})
	pulumi.RegisterOutputType(GovernanceEmailNotificationOutput{})
	pulumi.RegisterOutputType(GovernanceEmailNotificationPtrOutput{})
	pulumi.RegisterOutputType(GovernanceEmailNotificationResponseOutput{})
	pulumi.RegisterOutputType(GovernanceEmailNotificationResponsePtrOutput{})
	pulumi.RegisterOutputType(GovernanceRuleEmailNotificationOutput{})
	pulumi.RegisterOutputType(GovernanceRuleEmailNotificationPtrOutput{})
	pulumi.RegisterOutputType(GovernanceRuleEmailNotificationResponseOutput{})
	pulumi.RegisterOutputType(GovernanceRuleEmailNotificationResponsePtrOutput{})
	pulumi.RegisterOutputType(GovernanceRuleMetadataResponseOutput{})
	pulumi.RegisterOutputType(GovernanceRuleMetadataResponsePtrOutput{})
	pulumi.RegisterOutputType(GovernanceRuleOwnerSourceOutput{})
	pulumi.RegisterOutputType(GovernanceRuleOwnerSourceResponseOutput{})
	pulumi.RegisterOutputType(RemediationEtaOutput{})
	pulumi.RegisterOutputType(RemediationEtaPtrOutput{})
	pulumi.RegisterOutputType(RemediationEtaResponseOutput{})
	pulumi.RegisterOutputType(RemediationEtaResponsePtrOutput{})
}

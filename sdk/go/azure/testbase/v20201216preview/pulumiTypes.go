


package v20201216preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Command struct {
	Action            string `pulumi:"action"`
	AlwaysRun         *bool  `pulumi:"alwaysRun"`
	ApplyUpdateBefore *bool  `pulumi:"applyUpdateBefore"`
	Content           string `pulumi:"content"`
	ContentType       string `pulumi:"contentType"`
	MaxRunTime        *int   `pulumi:"maxRunTime"`
	Name              string `pulumi:"name"`
	RestartAfter      *bool  `pulumi:"restartAfter"`
	RunAsInteractive  *bool  `pulumi:"runAsInteractive"`
	RunElevated       *bool  `pulumi:"runElevated"`
}





type CommandInput interface {
	pulumi.Input

	ToCommandOutput() CommandOutput
	ToCommandOutputWithContext(context.Context) CommandOutput
}

type CommandArgs struct {
	Action            pulumi.StringInput  `pulumi:"action"`
	AlwaysRun         pulumi.BoolPtrInput `pulumi:"alwaysRun"`
	ApplyUpdateBefore pulumi.BoolPtrInput `pulumi:"applyUpdateBefore"`
	Content           pulumi.StringInput  `pulumi:"content"`
	ContentType       pulumi.StringInput  `pulumi:"contentType"`
	MaxRunTime        pulumi.IntPtrInput  `pulumi:"maxRunTime"`
	Name              pulumi.StringInput  `pulumi:"name"`
	RestartAfter      pulumi.BoolPtrInput `pulumi:"restartAfter"`
	RunAsInteractive  pulumi.BoolPtrInput `pulumi:"runAsInteractive"`
	RunElevated       pulumi.BoolPtrInput `pulumi:"runElevated"`
}

func (CommandArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Command)(nil)).Elem()
}

func (i CommandArgs) ToCommandOutput() CommandOutput {
	return i.ToCommandOutputWithContext(context.Background())
}

func (i CommandArgs) ToCommandOutputWithContext(ctx context.Context) CommandOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommandOutput)
}





type CommandArrayInput interface {
	pulumi.Input

	ToCommandArrayOutput() CommandArrayOutput
	ToCommandArrayOutputWithContext(context.Context) CommandArrayOutput
}

type CommandArray []CommandInput

func (CommandArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Command)(nil)).Elem()
}

func (i CommandArray) ToCommandArrayOutput() CommandArrayOutput {
	return i.ToCommandArrayOutputWithContext(context.Background())
}

func (i CommandArray) ToCommandArrayOutputWithContext(ctx context.Context) CommandArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommandArrayOutput)
}

type CommandOutput struct{ *pulumi.OutputState }

func (CommandOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Command)(nil)).Elem()
}

func (o CommandOutput) ToCommandOutput() CommandOutput {
	return o
}

func (o CommandOutput) ToCommandOutputWithContext(ctx context.Context) CommandOutput {
	return o
}

func (o CommandOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v Command) string { return v.Action }).(pulumi.StringOutput)
}

func (o CommandOutput) AlwaysRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Command) *bool { return v.AlwaysRun }).(pulumi.BoolPtrOutput)
}

func (o CommandOutput) ApplyUpdateBefore() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Command) *bool { return v.ApplyUpdateBefore }).(pulumi.BoolPtrOutput)
}

func (o CommandOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v Command) string { return v.Content }).(pulumi.StringOutput)
}

func (o CommandOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v Command) string { return v.ContentType }).(pulumi.StringOutput)
}

func (o CommandOutput) MaxRunTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Command) *int { return v.MaxRunTime }).(pulumi.IntPtrOutput)
}

func (o CommandOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Command) string { return v.Name }).(pulumi.StringOutput)
}

func (o CommandOutput) RestartAfter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Command) *bool { return v.RestartAfter }).(pulumi.BoolPtrOutput)
}

func (o CommandOutput) RunAsInteractive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Command) *bool { return v.RunAsInteractive }).(pulumi.BoolPtrOutput)
}

func (o CommandOutput) RunElevated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Command) *bool { return v.RunElevated }).(pulumi.BoolPtrOutput)
}

type CommandArrayOutput struct{ *pulumi.OutputState }

func (CommandArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Command)(nil)).Elem()
}

func (o CommandArrayOutput) ToCommandArrayOutput() CommandArrayOutput {
	return o
}

func (o CommandArrayOutput) ToCommandArrayOutputWithContext(ctx context.Context) CommandArrayOutput {
	return o
}

func (o CommandArrayOutput) Index(i pulumi.IntInput) CommandOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Command {
		return vs[0].([]Command)[vs[1].(int)]
	}).(CommandOutput)
}

type CommandResponse struct {
	Action            string `pulumi:"action"`
	AlwaysRun         *bool  `pulumi:"alwaysRun"`
	ApplyUpdateBefore *bool  `pulumi:"applyUpdateBefore"`
	Content           string `pulumi:"content"`
	ContentType       string `pulumi:"contentType"`
	MaxRunTime        *int   `pulumi:"maxRunTime"`
	Name              string `pulumi:"name"`
	RestartAfter      *bool  `pulumi:"restartAfter"`
	RunAsInteractive  *bool  `pulumi:"runAsInteractive"`
	RunElevated       *bool  `pulumi:"runElevated"`
}





type CommandResponseInput interface {
	pulumi.Input

	ToCommandResponseOutput() CommandResponseOutput
	ToCommandResponseOutputWithContext(context.Context) CommandResponseOutput
}

type CommandResponseArgs struct {
	Action            pulumi.StringInput  `pulumi:"action"`
	AlwaysRun         pulumi.BoolPtrInput `pulumi:"alwaysRun"`
	ApplyUpdateBefore pulumi.BoolPtrInput `pulumi:"applyUpdateBefore"`
	Content           pulumi.StringInput  `pulumi:"content"`
	ContentType       pulumi.StringInput  `pulumi:"contentType"`
	MaxRunTime        pulumi.IntPtrInput  `pulumi:"maxRunTime"`
	Name              pulumi.StringInput  `pulumi:"name"`
	RestartAfter      pulumi.BoolPtrInput `pulumi:"restartAfter"`
	RunAsInteractive  pulumi.BoolPtrInput `pulumi:"runAsInteractive"`
	RunElevated       pulumi.BoolPtrInput `pulumi:"runElevated"`
}

func (CommandResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CommandResponse)(nil)).Elem()
}

func (i CommandResponseArgs) ToCommandResponseOutput() CommandResponseOutput {
	return i.ToCommandResponseOutputWithContext(context.Background())
}

func (i CommandResponseArgs) ToCommandResponseOutputWithContext(ctx context.Context) CommandResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommandResponseOutput)
}





type CommandResponseArrayInput interface {
	pulumi.Input

	ToCommandResponseArrayOutput() CommandResponseArrayOutput
	ToCommandResponseArrayOutputWithContext(context.Context) CommandResponseArrayOutput
}

type CommandResponseArray []CommandResponseInput

func (CommandResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CommandResponse)(nil)).Elem()
}

func (i CommandResponseArray) ToCommandResponseArrayOutput() CommandResponseArrayOutput {
	return i.ToCommandResponseArrayOutputWithContext(context.Background())
}

func (i CommandResponseArray) ToCommandResponseArrayOutputWithContext(ctx context.Context) CommandResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommandResponseArrayOutput)
}

type CommandResponseOutput struct{ *pulumi.OutputState }

func (CommandResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommandResponse)(nil)).Elem()
}

func (o CommandResponseOutput) ToCommandResponseOutput() CommandResponseOutput {
	return o
}

func (o CommandResponseOutput) ToCommandResponseOutputWithContext(ctx context.Context) CommandResponseOutput {
	return o
}

func (o CommandResponseOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v CommandResponse) string { return v.Action }).(pulumi.StringOutput)
}

func (o CommandResponseOutput) AlwaysRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CommandResponse) *bool { return v.AlwaysRun }).(pulumi.BoolPtrOutput)
}

func (o CommandResponseOutput) ApplyUpdateBefore() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CommandResponse) *bool { return v.ApplyUpdateBefore }).(pulumi.BoolPtrOutput)
}

func (o CommandResponseOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v CommandResponse) string { return v.Content }).(pulumi.StringOutput)
}

func (o CommandResponseOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v CommandResponse) string { return v.ContentType }).(pulumi.StringOutput)
}

func (o CommandResponseOutput) MaxRunTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CommandResponse) *int { return v.MaxRunTime }).(pulumi.IntPtrOutput)
}

func (o CommandResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CommandResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o CommandResponseOutput) RestartAfter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CommandResponse) *bool { return v.RestartAfter }).(pulumi.BoolPtrOutput)
}

func (o CommandResponseOutput) RunAsInteractive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CommandResponse) *bool { return v.RunAsInteractive }).(pulumi.BoolPtrOutput)
}

func (o CommandResponseOutput) RunElevated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CommandResponse) *bool { return v.RunElevated }).(pulumi.BoolPtrOutput)
}

type CommandResponseArrayOutput struct{ *pulumi.OutputState }

func (CommandResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CommandResponse)(nil)).Elem()
}

func (o CommandResponseArrayOutput) ToCommandResponseArrayOutput() CommandResponseArrayOutput {
	return o
}

func (o CommandResponseArrayOutput) ToCommandResponseArrayOutputWithContext(ctx context.Context) CommandResponseArrayOutput {
	return o
}

func (o CommandResponseArrayOutput) Index(i pulumi.IntInput) CommandResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CommandResponse {
		return vs[0].([]CommandResponse)[vs[1].(int)]
	}).(CommandResponseOutput)
}

type DistributionGroupListReceiverValue struct {
	DistributionGroups []string `pulumi:"distributionGroups"`
}





type DistributionGroupListReceiverValueInput interface {
	pulumi.Input

	ToDistributionGroupListReceiverValueOutput() DistributionGroupListReceiverValueOutput
	ToDistributionGroupListReceiverValueOutputWithContext(context.Context) DistributionGroupListReceiverValueOutput
}

type DistributionGroupListReceiverValueArgs struct {
	DistributionGroups pulumi.StringArrayInput `pulumi:"distributionGroups"`
}

func (DistributionGroupListReceiverValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DistributionGroupListReceiverValue)(nil)).Elem()
}

func (i DistributionGroupListReceiverValueArgs) ToDistributionGroupListReceiverValueOutput() DistributionGroupListReceiverValueOutput {
	return i.ToDistributionGroupListReceiverValueOutputWithContext(context.Background())
}

func (i DistributionGroupListReceiverValueArgs) ToDistributionGroupListReceiverValueOutputWithContext(ctx context.Context) DistributionGroupListReceiverValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributionGroupListReceiverValueOutput)
}

func (i DistributionGroupListReceiverValueArgs) ToDistributionGroupListReceiverValuePtrOutput() DistributionGroupListReceiverValuePtrOutput {
	return i.ToDistributionGroupListReceiverValuePtrOutputWithContext(context.Background())
}

func (i DistributionGroupListReceiverValueArgs) ToDistributionGroupListReceiverValuePtrOutputWithContext(ctx context.Context) DistributionGroupListReceiverValuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributionGroupListReceiverValueOutput).ToDistributionGroupListReceiverValuePtrOutputWithContext(ctx)
}









type DistributionGroupListReceiverValuePtrInput interface {
	pulumi.Input

	ToDistributionGroupListReceiverValuePtrOutput() DistributionGroupListReceiverValuePtrOutput
	ToDistributionGroupListReceiverValuePtrOutputWithContext(context.Context) DistributionGroupListReceiverValuePtrOutput
}

type distributionGroupListReceiverValuePtrType DistributionGroupListReceiverValueArgs

func DistributionGroupListReceiverValuePtr(v *DistributionGroupListReceiverValueArgs) DistributionGroupListReceiverValuePtrInput {
	return (*distributionGroupListReceiverValuePtrType)(v)
}

func (*distributionGroupListReceiverValuePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DistributionGroupListReceiverValue)(nil)).Elem()
}

func (i *distributionGroupListReceiverValuePtrType) ToDistributionGroupListReceiverValuePtrOutput() DistributionGroupListReceiverValuePtrOutput {
	return i.ToDistributionGroupListReceiverValuePtrOutputWithContext(context.Background())
}

func (i *distributionGroupListReceiverValuePtrType) ToDistributionGroupListReceiverValuePtrOutputWithContext(ctx context.Context) DistributionGroupListReceiverValuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributionGroupListReceiverValuePtrOutput)
}

type DistributionGroupListReceiverValueOutput struct{ *pulumi.OutputState }

func (DistributionGroupListReceiverValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DistributionGroupListReceiverValue)(nil)).Elem()
}

func (o DistributionGroupListReceiverValueOutput) ToDistributionGroupListReceiverValueOutput() DistributionGroupListReceiverValueOutput {
	return o
}

func (o DistributionGroupListReceiverValueOutput) ToDistributionGroupListReceiverValueOutputWithContext(ctx context.Context) DistributionGroupListReceiverValueOutput {
	return o
}

func (o DistributionGroupListReceiverValueOutput) ToDistributionGroupListReceiverValuePtrOutput() DistributionGroupListReceiverValuePtrOutput {
	return o.ToDistributionGroupListReceiverValuePtrOutputWithContext(context.Background())
}

func (o DistributionGroupListReceiverValueOutput) ToDistributionGroupListReceiverValuePtrOutputWithContext(ctx context.Context) DistributionGroupListReceiverValuePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DistributionGroupListReceiverValue) *DistributionGroupListReceiverValue {
		return &v
	}).(DistributionGroupListReceiverValuePtrOutput)
}

func (o DistributionGroupListReceiverValueOutput) DistributionGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DistributionGroupListReceiverValue) []string { return v.DistributionGroups }).(pulumi.StringArrayOutput)
}

type DistributionGroupListReceiverValuePtrOutput struct{ *pulumi.OutputState }

func (DistributionGroupListReceiverValuePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DistributionGroupListReceiverValue)(nil)).Elem()
}

func (o DistributionGroupListReceiverValuePtrOutput) ToDistributionGroupListReceiverValuePtrOutput() DistributionGroupListReceiverValuePtrOutput {
	return o
}

func (o DistributionGroupListReceiverValuePtrOutput) ToDistributionGroupListReceiverValuePtrOutputWithContext(ctx context.Context) DistributionGroupListReceiverValuePtrOutput {
	return o
}

func (o DistributionGroupListReceiverValuePtrOutput) Elem() DistributionGroupListReceiverValueOutput {
	return o.ApplyT(func(v *DistributionGroupListReceiverValue) DistributionGroupListReceiverValue {
		if v != nil {
			return *v
		}
		var ret DistributionGroupListReceiverValue
		return ret
	}).(DistributionGroupListReceiverValueOutput)
}

func (o DistributionGroupListReceiverValuePtrOutput) DistributionGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DistributionGroupListReceiverValue) []string {
		if v == nil {
			return nil
		}
		return v.DistributionGroups
	}).(pulumi.StringArrayOutput)
}

type DistributionGroupListReceiverValueResponse struct {
	DistributionGroups []string `pulumi:"distributionGroups"`
}





type DistributionGroupListReceiverValueResponseInput interface {
	pulumi.Input

	ToDistributionGroupListReceiverValueResponseOutput() DistributionGroupListReceiverValueResponseOutput
	ToDistributionGroupListReceiverValueResponseOutputWithContext(context.Context) DistributionGroupListReceiverValueResponseOutput
}

type DistributionGroupListReceiverValueResponseArgs struct {
	DistributionGroups pulumi.StringArrayInput `pulumi:"distributionGroups"`
}

func (DistributionGroupListReceiverValueResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DistributionGroupListReceiverValueResponse)(nil)).Elem()
}

func (i DistributionGroupListReceiverValueResponseArgs) ToDistributionGroupListReceiverValueResponseOutput() DistributionGroupListReceiverValueResponseOutput {
	return i.ToDistributionGroupListReceiverValueResponseOutputWithContext(context.Background())
}

func (i DistributionGroupListReceiverValueResponseArgs) ToDistributionGroupListReceiverValueResponseOutputWithContext(ctx context.Context) DistributionGroupListReceiverValueResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributionGroupListReceiverValueResponseOutput)
}

func (i DistributionGroupListReceiverValueResponseArgs) ToDistributionGroupListReceiverValueResponsePtrOutput() DistributionGroupListReceiverValueResponsePtrOutput {
	return i.ToDistributionGroupListReceiverValueResponsePtrOutputWithContext(context.Background())
}

func (i DistributionGroupListReceiverValueResponseArgs) ToDistributionGroupListReceiverValueResponsePtrOutputWithContext(ctx context.Context) DistributionGroupListReceiverValueResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributionGroupListReceiverValueResponseOutput).ToDistributionGroupListReceiverValueResponsePtrOutputWithContext(ctx)
}









type DistributionGroupListReceiverValueResponsePtrInput interface {
	pulumi.Input

	ToDistributionGroupListReceiverValueResponsePtrOutput() DistributionGroupListReceiverValueResponsePtrOutput
	ToDistributionGroupListReceiverValueResponsePtrOutputWithContext(context.Context) DistributionGroupListReceiverValueResponsePtrOutput
}

type distributionGroupListReceiverValueResponsePtrType DistributionGroupListReceiverValueResponseArgs

func DistributionGroupListReceiverValueResponsePtr(v *DistributionGroupListReceiverValueResponseArgs) DistributionGroupListReceiverValueResponsePtrInput {
	return (*distributionGroupListReceiverValueResponsePtrType)(v)
}

func (*distributionGroupListReceiverValueResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DistributionGroupListReceiverValueResponse)(nil)).Elem()
}

func (i *distributionGroupListReceiverValueResponsePtrType) ToDistributionGroupListReceiverValueResponsePtrOutput() DistributionGroupListReceiverValueResponsePtrOutput {
	return i.ToDistributionGroupListReceiverValueResponsePtrOutputWithContext(context.Background())
}

func (i *distributionGroupListReceiverValueResponsePtrType) ToDistributionGroupListReceiverValueResponsePtrOutputWithContext(ctx context.Context) DistributionGroupListReceiverValueResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributionGroupListReceiverValueResponsePtrOutput)
}

type DistributionGroupListReceiverValueResponseOutput struct{ *pulumi.OutputState }

func (DistributionGroupListReceiverValueResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DistributionGroupListReceiverValueResponse)(nil)).Elem()
}

func (o DistributionGroupListReceiverValueResponseOutput) ToDistributionGroupListReceiverValueResponseOutput() DistributionGroupListReceiverValueResponseOutput {
	return o
}

func (o DistributionGroupListReceiverValueResponseOutput) ToDistributionGroupListReceiverValueResponseOutputWithContext(ctx context.Context) DistributionGroupListReceiverValueResponseOutput {
	return o
}

func (o DistributionGroupListReceiverValueResponseOutput) ToDistributionGroupListReceiverValueResponsePtrOutput() DistributionGroupListReceiverValueResponsePtrOutput {
	return o.ToDistributionGroupListReceiverValueResponsePtrOutputWithContext(context.Background())
}

func (o DistributionGroupListReceiverValueResponseOutput) ToDistributionGroupListReceiverValueResponsePtrOutputWithContext(ctx context.Context) DistributionGroupListReceiverValueResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DistributionGroupListReceiverValueResponse) *DistributionGroupListReceiverValueResponse {
		return &v
	}).(DistributionGroupListReceiverValueResponsePtrOutput)
}

func (o DistributionGroupListReceiverValueResponseOutput) DistributionGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DistributionGroupListReceiverValueResponse) []string { return v.DistributionGroups }).(pulumi.StringArrayOutput)
}

type DistributionGroupListReceiverValueResponsePtrOutput struct{ *pulumi.OutputState }

func (DistributionGroupListReceiverValueResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DistributionGroupListReceiverValueResponse)(nil)).Elem()
}

func (o DistributionGroupListReceiverValueResponsePtrOutput) ToDistributionGroupListReceiverValueResponsePtrOutput() DistributionGroupListReceiverValueResponsePtrOutput {
	return o
}

func (o DistributionGroupListReceiverValueResponsePtrOutput) ToDistributionGroupListReceiverValueResponsePtrOutputWithContext(ctx context.Context) DistributionGroupListReceiverValueResponsePtrOutput {
	return o
}

func (o DistributionGroupListReceiverValueResponsePtrOutput) Elem() DistributionGroupListReceiverValueResponseOutput {
	return o.ApplyT(func(v *DistributionGroupListReceiverValueResponse) DistributionGroupListReceiverValueResponse {
		if v != nil {
			return *v
		}
		var ret DistributionGroupListReceiverValueResponse
		return ret
	}).(DistributionGroupListReceiverValueResponseOutput)
}

func (o DistributionGroupListReceiverValueResponsePtrOutput) DistributionGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DistributionGroupListReceiverValueResponse) []string {
		if v == nil {
			return nil
		}
		return v.DistributionGroups
	}).(pulumi.StringArrayOutput)
}

type NotificationEventReceiver struct {
	ReceiverType  *string                    `pulumi:"receiverType"`
	ReceiverValue *NotificationReceiverValue `pulumi:"receiverValue"`
}





type NotificationEventReceiverInput interface {
	pulumi.Input

	ToNotificationEventReceiverOutput() NotificationEventReceiverOutput
	ToNotificationEventReceiverOutputWithContext(context.Context) NotificationEventReceiverOutput
}

type NotificationEventReceiverArgs struct {
	ReceiverType  pulumi.StringPtrInput             `pulumi:"receiverType"`
	ReceiverValue NotificationReceiverValuePtrInput `pulumi:"receiverValue"`
}

func (NotificationEventReceiverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationEventReceiver)(nil)).Elem()
}

func (i NotificationEventReceiverArgs) ToNotificationEventReceiverOutput() NotificationEventReceiverOutput {
	return i.ToNotificationEventReceiverOutputWithContext(context.Background())
}

func (i NotificationEventReceiverArgs) ToNotificationEventReceiverOutputWithContext(ctx context.Context) NotificationEventReceiverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationEventReceiverOutput)
}





type NotificationEventReceiverArrayInput interface {
	pulumi.Input

	ToNotificationEventReceiverArrayOutput() NotificationEventReceiverArrayOutput
	ToNotificationEventReceiverArrayOutputWithContext(context.Context) NotificationEventReceiverArrayOutput
}

type NotificationEventReceiverArray []NotificationEventReceiverInput

func (NotificationEventReceiverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotificationEventReceiver)(nil)).Elem()
}

func (i NotificationEventReceiverArray) ToNotificationEventReceiverArrayOutput() NotificationEventReceiverArrayOutput {
	return i.ToNotificationEventReceiverArrayOutputWithContext(context.Background())
}

func (i NotificationEventReceiverArray) ToNotificationEventReceiverArrayOutputWithContext(ctx context.Context) NotificationEventReceiverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationEventReceiverArrayOutput)
}

type NotificationEventReceiverOutput struct{ *pulumi.OutputState }

func (NotificationEventReceiverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationEventReceiver)(nil)).Elem()
}

func (o NotificationEventReceiverOutput) ToNotificationEventReceiverOutput() NotificationEventReceiverOutput {
	return o
}

func (o NotificationEventReceiverOutput) ToNotificationEventReceiverOutputWithContext(ctx context.Context) NotificationEventReceiverOutput {
	return o
}

func (o NotificationEventReceiverOutput) ReceiverType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationEventReceiver) *string { return v.ReceiverType }).(pulumi.StringPtrOutput)
}

func (o NotificationEventReceiverOutput) ReceiverValue() NotificationReceiverValuePtrOutput {
	return o.ApplyT(func(v NotificationEventReceiver) *NotificationReceiverValue { return v.ReceiverValue }).(NotificationReceiverValuePtrOutput)
}

type NotificationEventReceiverArrayOutput struct{ *pulumi.OutputState }

func (NotificationEventReceiverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotificationEventReceiver)(nil)).Elem()
}

func (o NotificationEventReceiverArrayOutput) ToNotificationEventReceiverArrayOutput() NotificationEventReceiverArrayOutput {
	return o
}

func (o NotificationEventReceiverArrayOutput) ToNotificationEventReceiverArrayOutputWithContext(ctx context.Context) NotificationEventReceiverArrayOutput {
	return o
}

func (o NotificationEventReceiverArrayOutput) Index(i pulumi.IntInput) NotificationEventReceiverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NotificationEventReceiver {
		return vs[0].([]NotificationEventReceiver)[vs[1].(int)]
	}).(NotificationEventReceiverOutput)
}

type NotificationEventReceiverResponse struct {
	ReceiverType  *string                            `pulumi:"receiverType"`
	ReceiverValue *NotificationReceiverValueResponse `pulumi:"receiverValue"`
}





type NotificationEventReceiverResponseInput interface {
	pulumi.Input

	ToNotificationEventReceiverResponseOutput() NotificationEventReceiverResponseOutput
	ToNotificationEventReceiverResponseOutputWithContext(context.Context) NotificationEventReceiverResponseOutput
}

type NotificationEventReceiverResponseArgs struct {
	ReceiverType  pulumi.StringPtrInput                     `pulumi:"receiverType"`
	ReceiverValue NotificationReceiverValueResponsePtrInput `pulumi:"receiverValue"`
}

func (NotificationEventReceiverResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationEventReceiverResponse)(nil)).Elem()
}

func (i NotificationEventReceiverResponseArgs) ToNotificationEventReceiverResponseOutput() NotificationEventReceiverResponseOutput {
	return i.ToNotificationEventReceiverResponseOutputWithContext(context.Background())
}

func (i NotificationEventReceiverResponseArgs) ToNotificationEventReceiverResponseOutputWithContext(ctx context.Context) NotificationEventReceiverResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationEventReceiverResponseOutput)
}





type NotificationEventReceiverResponseArrayInput interface {
	pulumi.Input

	ToNotificationEventReceiverResponseArrayOutput() NotificationEventReceiverResponseArrayOutput
	ToNotificationEventReceiverResponseArrayOutputWithContext(context.Context) NotificationEventReceiverResponseArrayOutput
}

type NotificationEventReceiverResponseArray []NotificationEventReceiverResponseInput

func (NotificationEventReceiverResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotificationEventReceiverResponse)(nil)).Elem()
}

func (i NotificationEventReceiverResponseArray) ToNotificationEventReceiverResponseArrayOutput() NotificationEventReceiverResponseArrayOutput {
	return i.ToNotificationEventReceiverResponseArrayOutputWithContext(context.Background())
}

func (i NotificationEventReceiverResponseArray) ToNotificationEventReceiverResponseArrayOutputWithContext(ctx context.Context) NotificationEventReceiverResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationEventReceiverResponseArrayOutput)
}

type NotificationEventReceiverResponseOutput struct{ *pulumi.OutputState }

func (NotificationEventReceiverResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationEventReceiverResponse)(nil)).Elem()
}

func (o NotificationEventReceiverResponseOutput) ToNotificationEventReceiverResponseOutput() NotificationEventReceiverResponseOutput {
	return o
}

func (o NotificationEventReceiverResponseOutput) ToNotificationEventReceiverResponseOutputWithContext(ctx context.Context) NotificationEventReceiverResponseOutput {
	return o
}

func (o NotificationEventReceiverResponseOutput) ReceiverType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationEventReceiverResponse) *string { return v.ReceiverType }).(pulumi.StringPtrOutput)
}

func (o NotificationEventReceiverResponseOutput) ReceiverValue() NotificationReceiverValueResponsePtrOutput {
	return o.ApplyT(func(v NotificationEventReceiverResponse) *NotificationReceiverValueResponse { return v.ReceiverValue }).(NotificationReceiverValueResponsePtrOutput)
}

type NotificationEventReceiverResponseArrayOutput struct{ *pulumi.OutputState }

func (NotificationEventReceiverResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotificationEventReceiverResponse)(nil)).Elem()
}

func (o NotificationEventReceiverResponseArrayOutput) ToNotificationEventReceiverResponseArrayOutput() NotificationEventReceiverResponseArrayOutput {
	return o
}

func (o NotificationEventReceiverResponseArrayOutput) ToNotificationEventReceiverResponseArrayOutputWithContext(ctx context.Context) NotificationEventReceiverResponseArrayOutput {
	return o
}

func (o NotificationEventReceiverResponseArrayOutput) Index(i pulumi.IntInput) NotificationEventReceiverResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NotificationEventReceiverResponse {
		return vs[0].([]NotificationEventReceiverResponse)[vs[1].(int)]
	}).(NotificationEventReceiverResponseOutput)
}

type NotificationReceiverValue struct {
	DistributionGroupListReceiverValue *DistributionGroupListReceiverValue `pulumi:"distributionGroupListReceiverValue"`
	SubscriptionReceiverValue          *SubscriptionReceiverValue          `pulumi:"subscriptionReceiverValue"`
	UserObjectReceiverValue            *UserObjectReceiverValue            `pulumi:"userObjectReceiverValue"`
}





type NotificationReceiverValueInput interface {
	pulumi.Input

	ToNotificationReceiverValueOutput() NotificationReceiverValueOutput
	ToNotificationReceiverValueOutputWithContext(context.Context) NotificationReceiverValueOutput
}

type NotificationReceiverValueArgs struct {
	DistributionGroupListReceiverValue DistributionGroupListReceiverValuePtrInput `pulumi:"distributionGroupListReceiverValue"`
	SubscriptionReceiverValue          SubscriptionReceiverValuePtrInput          `pulumi:"subscriptionReceiverValue"`
	UserObjectReceiverValue            UserObjectReceiverValuePtrInput            `pulumi:"userObjectReceiverValue"`
}

func (NotificationReceiverValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationReceiverValue)(nil)).Elem()
}

func (i NotificationReceiverValueArgs) ToNotificationReceiverValueOutput() NotificationReceiverValueOutput {
	return i.ToNotificationReceiverValueOutputWithContext(context.Background())
}

func (i NotificationReceiverValueArgs) ToNotificationReceiverValueOutputWithContext(ctx context.Context) NotificationReceiverValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationReceiverValueOutput)
}

func (i NotificationReceiverValueArgs) ToNotificationReceiverValuePtrOutput() NotificationReceiverValuePtrOutput {
	return i.ToNotificationReceiverValuePtrOutputWithContext(context.Background())
}

func (i NotificationReceiverValueArgs) ToNotificationReceiverValuePtrOutputWithContext(ctx context.Context) NotificationReceiverValuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationReceiverValueOutput).ToNotificationReceiverValuePtrOutputWithContext(ctx)
}









type NotificationReceiverValuePtrInput interface {
	pulumi.Input

	ToNotificationReceiverValuePtrOutput() NotificationReceiverValuePtrOutput
	ToNotificationReceiverValuePtrOutputWithContext(context.Context) NotificationReceiverValuePtrOutput
}

type notificationReceiverValuePtrType NotificationReceiverValueArgs

func NotificationReceiverValuePtr(v *NotificationReceiverValueArgs) NotificationReceiverValuePtrInput {
	return (*notificationReceiverValuePtrType)(v)
}

func (*notificationReceiverValuePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationReceiverValue)(nil)).Elem()
}

func (i *notificationReceiverValuePtrType) ToNotificationReceiverValuePtrOutput() NotificationReceiverValuePtrOutput {
	return i.ToNotificationReceiverValuePtrOutputWithContext(context.Background())
}

func (i *notificationReceiverValuePtrType) ToNotificationReceiverValuePtrOutputWithContext(ctx context.Context) NotificationReceiverValuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationReceiverValuePtrOutput)
}

type NotificationReceiverValueOutput struct{ *pulumi.OutputState }

func (NotificationReceiverValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationReceiverValue)(nil)).Elem()
}

func (o NotificationReceiverValueOutput) ToNotificationReceiverValueOutput() NotificationReceiverValueOutput {
	return o
}

func (o NotificationReceiverValueOutput) ToNotificationReceiverValueOutputWithContext(ctx context.Context) NotificationReceiverValueOutput {
	return o
}

func (o NotificationReceiverValueOutput) ToNotificationReceiverValuePtrOutput() NotificationReceiverValuePtrOutput {
	return o.ToNotificationReceiverValuePtrOutputWithContext(context.Background())
}

func (o NotificationReceiverValueOutput) ToNotificationReceiverValuePtrOutputWithContext(ctx context.Context) NotificationReceiverValuePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotificationReceiverValue) *NotificationReceiverValue {
		return &v
	}).(NotificationReceiverValuePtrOutput)
}

func (o NotificationReceiverValueOutput) DistributionGroupListReceiverValue() DistributionGroupListReceiverValuePtrOutput {
	return o.ApplyT(func(v NotificationReceiverValue) *DistributionGroupListReceiverValue {
		return v.DistributionGroupListReceiverValue
	}).(DistributionGroupListReceiverValuePtrOutput)
}

func (o NotificationReceiverValueOutput) SubscriptionReceiverValue() SubscriptionReceiverValuePtrOutput {
	return o.ApplyT(func(v NotificationReceiverValue) *SubscriptionReceiverValue { return v.SubscriptionReceiverValue }).(SubscriptionReceiverValuePtrOutput)
}

func (o NotificationReceiverValueOutput) UserObjectReceiverValue() UserObjectReceiverValuePtrOutput {
	return o.ApplyT(func(v NotificationReceiverValue) *UserObjectReceiverValue { return v.UserObjectReceiverValue }).(UserObjectReceiverValuePtrOutput)
}

type NotificationReceiverValuePtrOutput struct{ *pulumi.OutputState }

func (NotificationReceiverValuePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationReceiverValue)(nil)).Elem()
}

func (o NotificationReceiverValuePtrOutput) ToNotificationReceiverValuePtrOutput() NotificationReceiverValuePtrOutput {
	return o
}

func (o NotificationReceiverValuePtrOutput) ToNotificationReceiverValuePtrOutputWithContext(ctx context.Context) NotificationReceiverValuePtrOutput {
	return o
}

func (o NotificationReceiverValuePtrOutput) Elem() NotificationReceiverValueOutput {
	return o.ApplyT(func(v *NotificationReceiverValue) NotificationReceiverValue {
		if v != nil {
			return *v
		}
		var ret NotificationReceiverValue
		return ret
	}).(NotificationReceiverValueOutput)
}

func (o NotificationReceiverValuePtrOutput) DistributionGroupListReceiverValue() DistributionGroupListReceiverValuePtrOutput {
	return o.ApplyT(func(v *NotificationReceiverValue) *DistributionGroupListReceiverValue {
		if v == nil {
			return nil
		}
		return v.DistributionGroupListReceiverValue
	}).(DistributionGroupListReceiverValuePtrOutput)
}

func (o NotificationReceiverValuePtrOutput) SubscriptionReceiverValue() SubscriptionReceiverValuePtrOutput {
	return o.ApplyT(func(v *NotificationReceiverValue) *SubscriptionReceiverValue {
		if v == nil {
			return nil
		}
		return v.SubscriptionReceiverValue
	}).(SubscriptionReceiverValuePtrOutput)
}

func (o NotificationReceiverValuePtrOutput) UserObjectReceiverValue() UserObjectReceiverValuePtrOutput {
	return o.ApplyT(func(v *NotificationReceiverValue) *UserObjectReceiverValue {
		if v == nil {
			return nil
		}
		return v.UserObjectReceiverValue
	}).(UserObjectReceiverValuePtrOutput)
}

type NotificationReceiverValueResponse struct {
	DistributionGroupListReceiverValue *DistributionGroupListReceiverValueResponse `pulumi:"distributionGroupListReceiverValue"`
	SubscriptionReceiverValue          *SubscriptionReceiverValueResponse          `pulumi:"subscriptionReceiverValue"`
	UserObjectReceiverValue            *UserObjectReceiverValueResponse            `pulumi:"userObjectReceiverValue"`
}





type NotificationReceiverValueResponseInput interface {
	pulumi.Input

	ToNotificationReceiverValueResponseOutput() NotificationReceiverValueResponseOutput
	ToNotificationReceiverValueResponseOutputWithContext(context.Context) NotificationReceiverValueResponseOutput
}

type NotificationReceiverValueResponseArgs struct {
	DistributionGroupListReceiverValue DistributionGroupListReceiverValueResponsePtrInput `pulumi:"distributionGroupListReceiverValue"`
	SubscriptionReceiverValue          SubscriptionReceiverValueResponsePtrInput          `pulumi:"subscriptionReceiverValue"`
	UserObjectReceiverValue            UserObjectReceiverValueResponsePtrInput            `pulumi:"userObjectReceiverValue"`
}

func (NotificationReceiverValueResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationReceiverValueResponse)(nil)).Elem()
}

func (i NotificationReceiverValueResponseArgs) ToNotificationReceiverValueResponseOutput() NotificationReceiverValueResponseOutput {
	return i.ToNotificationReceiverValueResponseOutputWithContext(context.Background())
}

func (i NotificationReceiverValueResponseArgs) ToNotificationReceiverValueResponseOutputWithContext(ctx context.Context) NotificationReceiverValueResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationReceiverValueResponseOutput)
}

func (i NotificationReceiverValueResponseArgs) ToNotificationReceiverValueResponsePtrOutput() NotificationReceiverValueResponsePtrOutput {
	return i.ToNotificationReceiverValueResponsePtrOutputWithContext(context.Background())
}

func (i NotificationReceiverValueResponseArgs) ToNotificationReceiverValueResponsePtrOutputWithContext(ctx context.Context) NotificationReceiverValueResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationReceiverValueResponseOutput).ToNotificationReceiverValueResponsePtrOutputWithContext(ctx)
}









type NotificationReceiverValueResponsePtrInput interface {
	pulumi.Input

	ToNotificationReceiverValueResponsePtrOutput() NotificationReceiverValueResponsePtrOutput
	ToNotificationReceiverValueResponsePtrOutputWithContext(context.Context) NotificationReceiverValueResponsePtrOutput
}

type notificationReceiverValueResponsePtrType NotificationReceiverValueResponseArgs

func NotificationReceiverValueResponsePtr(v *NotificationReceiverValueResponseArgs) NotificationReceiverValueResponsePtrInput {
	return (*notificationReceiverValueResponsePtrType)(v)
}

func (*notificationReceiverValueResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationReceiverValueResponse)(nil)).Elem()
}

func (i *notificationReceiverValueResponsePtrType) ToNotificationReceiverValueResponsePtrOutput() NotificationReceiverValueResponsePtrOutput {
	return i.ToNotificationReceiverValueResponsePtrOutputWithContext(context.Background())
}

func (i *notificationReceiverValueResponsePtrType) ToNotificationReceiverValueResponsePtrOutputWithContext(ctx context.Context) NotificationReceiverValueResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationReceiverValueResponsePtrOutput)
}

type NotificationReceiverValueResponseOutput struct{ *pulumi.OutputState }

func (NotificationReceiverValueResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationReceiverValueResponse)(nil)).Elem()
}

func (o NotificationReceiverValueResponseOutput) ToNotificationReceiverValueResponseOutput() NotificationReceiverValueResponseOutput {
	return o
}

func (o NotificationReceiverValueResponseOutput) ToNotificationReceiverValueResponseOutputWithContext(ctx context.Context) NotificationReceiverValueResponseOutput {
	return o
}

func (o NotificationReceiverValueResponseOutput) ToNotificationReceiverValueResponsePtrOutput() NotificationReceiverValueResponsePtrOutput {
	return o.ToNotificationReceiverValueResponsePtrOutputWithContext(context.Background())
}

func (o NotificationReceiverValueResponseOutput) ToNotificationReceiverValueResponsePtrOutputWithContext(ctx context.Context) NotificationReceiverValueResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotificationReceiverValueResponse) *NotificationReceiverValueResponse {
		return &v
	}).(NotificationReceiverValueResponsePtrOutput)
}

func (o NotificationReceiverValueResponseOutput) DistributionGroupListReceiverValue() DistributionGroupListReceiverValueResponsePtrOutput {
	return o.ApplyT(func(v NotificationReceiverValueResponse) *DistributionGroupListReceiverValueResponse {
		return v.DistributionGroupListReceiverValue
	}).(DistributionGroupListReceiverValueResponsePtrOutput)
}

func (o NotificationReceiverValueResponseOutput) SubscriptionReceiverValue() SubscriptionReceiverValueResponsePtrOutput {
	return o.ApplyT(func(v NotificationReceiverValueResponse) *SubscriptionReceiverValueResponse {
		return v.SubscriptionReceiverValue
	}).(SubscriptionReceiverValueResponsePtrOutput)
}

func (o NotificationReceiverValueResponseOutput) UserObjectReceiverValue() UserObjectReceiverValueResponsePtrOutput {
	return o.ApplyT(func(v NotificationReceiverValueResponse) *UserObjectReceiverValueResponse {
		return v.UserObjectReceiverValue
	}).(UserObjectReceiverValueResponsePtrOutput)
}

type NotificationReceiverValueResponsePtrOutput struct{ *pulumi.OutputState }

func (NotificationReceiverValueResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationReceiverValueResponse)(nil)).Elem()
}

func (o NotificationReceiverValueResponsePtrOutput) ToNotificationReceiverValueResponsePtrOutput() NotificationReceiverValueResponsePtrOutput {
	return o
}

func (o NotificationReceiverValueResponsePtrOutput) ToNotificationReceiverValueResponsePtrOutputWithContext(ctx context.Context) NotificationReceiverValueResponsePtrOutput {
	return o
}

func (o NotificationReceiverValueResponsePtrOutput) Elem() NotificationReceiverValueResponseOutput {
	return o.ApplyT(func(v *NotificationReceiverValueResponse) NotificationReceiverValueResponse {
		if v != nil {
			return *v
		}
		var ret NotificationReceiverValueResponse
		return ret
	}).(NotificationReceiverValueResponseOutput)
}

func (o NotificationReceiverValueResponsePtrOutput) DistributionGroupListReceiverValue() DistributionGroupListReceiverValueResponsePtrOutput {
	return o.ApplyT(func(v *NotificationReceiverValueResponse) *DistributionGroupListReceiverValueResponse {
		if v == nil {
			return nil
		}
		return v.DistributionGroupListReceiverValue
	}).(DistributionGroupListReceiverValueResponsePtrOutput)
}

func (o NotificationReceiverValueResponsePtrOutput) SubscriptionReceiverValue() SubscriptionReceiverValueResponsePtrOutput {
	return o.ApplyT(func(v *NotificationReceiverValueResponse) *SubscriptionReceiverValueResponse {
		if v == nil {
			return nil
		}
		return v.SubscriptionReceiverValue
	}).(SubscriptionReceiverValueResponsePtrOutput)
}

func (o NotificationReceiverValueResponsePtrOutput) UserObjectReceiverValue() UserObjectReceiverValueResponsePtrOutput {
	return o.ApplyT(func(v *NotificationReceiverValueResponse) *UserObjectReceiverValueResponse {
		if v == nil {
			return nil
		}
		return v.UserObjectReceiverValue
	}).(UserObjectReceiverValueResponsePtrOutput)
}

type PackageValidationResultResponse struct {
	Errors         []string `pulumi:"errors"`
	IsValid        bool     `pulumi:"isValid"`
	ValidationName string   `pulumi:"validationName"`
}





type PackageValidationResultResponseInput interface {
	pulumi.Input

	ToPackageValidationResultResponseOutput() PackageValidationResultResponseOutput
	ToPackageValidationResultResponseOutputWithContext(context.Context) PackageValidationResultResponseOutput
}

type PackageValidationResultResponseArgs struct {
	Errors         pulumi.StringArrayInput `pulumi:"errors"`
	IsValid        pulumi.BoolInput        `pulumi:"isValid"`
	ValidationName pulumi.StringInput      `pulumi:"validationName"`
}

func (PackageValidationResultResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageValidationResultResponse)(nil)).Elem()
}

func (i PackageValidationResultResponseArgs) ToPackageValidationResultResponseOutput() PackageValidationResultResponseOutput {
	return i.ToPackageValidationResultResponseOutputWithContext(context.Background())
}

func (i PackageValidationResultResponseArgs) ToPackageValidationResultResponseOutputWithContext(ctx context.Context) PackageValidationResultResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageValidationResultResponseOutput)
}





type PackageValidationResultResponseArrayInput interface {
	pulumi.Input

	ToPackageValidationResultResponseArrayOutput() PackageValidationResultResponseArrayOutput
	ToPackageValidationResultResponseArrayOutputWithContext(context.Context) PackageValidationResultResponseArrayOutput
}

type PackageValidationResultResponseArray []PackageValidationResultResponseInput

func (PackageValidationResultResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PackageValidationResultResponse)(nil)).Elem()
}

func (i PackageValidationResultResponseArray) ToPackageValidationResultResponseArrayOutput() PackageValidationResultResponseArrayOutput {
	return i.ToPackageValidationResultResponseArrayOutputWithContext(context.Background())
}

func (i PackageValidationResultResponseArray) ToPackageValidationResultResponseArrayOutputWithContext(ctx context.Context) PackageValidationResultResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageValidationResultResponseArrayOutput)
}

type PackageValidationResultResponseOutput struct{ *pulumi.OutputState }

func (PackageValidationResultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageValidationResultResponse)(nil)).Elem()
}

func (o PackageValidationResultResponseOutput) ToPackageValidationResultResponseOutput() PackageValidationResultResponseOutput {
	return o
}

func (o PackageValidationResultResponseOutput) ToPackageValidationResultResponseOutputWithContext(ctx context.Context) PackageValidationResultResponseOutput {
	return o
}

func (o PackageValidationResultResponseOutput) Errors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PackageValidationResultResponse) []string { return v.Errors }).(pulumi.StringArrayOutput)
}

func (o PackageValidationResultResponseOutput) IsValid() pulumi.BoolOutput {
	return o.ApplyT(func(v PackageValidationResultResponse) bool { return v.IsValid }).(pulumi.BoolOutput)
}

func (o PackageValidationResultResponseOutput) ValidationName() pulumi.StringOutput {
	return o.ApplyT(func(v PackageValidationResultResponse) string { return v.ValidationName }).(pulumi.StringOutput)
}

type PackageValidationResultResponseArrayOutput struct{ *pulumi.OutputState }

func (PackageValidationResultResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PackageValidationResultResponse)(nil)).Elem()
}

func (o PackageValidationResultResponseArrayOutput) ToPackageValidationResultResponseArrayOutput() PackageValidationResultResponseArrayOutput {
	return o
}

func (o PackageValidationResultResponseArrayOutput) ToPackageValidationResultResponseArrayOutputWithContext(ctx context.Context) PackageValidationResultResponseArrayOutput {
	return o
}

func (o PackageValidationResultResponseArrayOutput) Index(i pulumi.IntInput) PackageValidationResultResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PackageValidationResultResponse {
		return vs[0].([]PackageValidationResultResponse)[vs[1].(int)]
	}).(PackageValidationResultResponseOutput)
}

type SubscriptionReceiverValue struct {
	Role             *string `pulumi:"role"`
	SubscriptionId   *string `pulumi:"subscriptionId"`
	SubscriptionName *string `pulumi:"subscriptionName"`
}





type SubscriptionReceiverValueInput interface {
	pulumi.Input

	ToSubscriptionReceiverValueOutput() SubscriptionReceiverValueOutput
	ToSubscriptionReceiverValueOutputWithContext(context.Context) SubscriptionReceiverValueOutput
}

type SubscriptionReceiverValueArgs struct {
	Role             pulumi.StringPtrInput `pulumi:"role"`
	SubscriptionId   pulumi.StringPtrInput `pulumi:"subscriptionId"`
	SubscriptionName pulumi.StringPtrInput `pulumi:"subscriptionName"`
}

func (SubscriptionReceiverValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionReceiverValue)(nil)).Elem()
}

func (i SubscriptionReceiverValueArgs) ToSubscriptionReceiverValueOutput() SubscriptionReceiverValueOutput {
	return i.ToSubscriptionReceiverValueOutputWithContext(context.Background())
}

func (i SubscriptionReceiverValueArgs) ToSubscriptionReceiverValueOutputWithContext(ctx context.Context) SubscriptionReceiverValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionReceiverValueOutput)
}

func (i SubscriptionReceiverValueArgs) ToSubscriptionReceiverValuePtrOutput() SubscriptionReceiverValuePtrOutput {
	return i.ToSubscriptionReceiverValuePtrOutputWithContext(context.Background())
}

func (i SubscriptionReceiverValueArgs) ToSubscriptionReceiverValuePtrOutputWithContext(ctx context.Context) SubscriptionReceiverValuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionReceiverValueOutput).ToSubscriptionReceiverValuePtrOutputWithContext(ctx)
}









type SubscriptionReceiverValuePtrInput interface {
	pulumi.Input

	ToSubscriptionReceiverValuePtrOutput() SubscriptionReceiverValuePtrOutput
	ToSubscriptionReceiverValuePtrOutputWithContext(context.Context) SubscriptionReceiverValuePtrOutput
}

type subscriptionReceiverValuePtrType SubscriptionReceiverValueArgs

func SubscriptionReceiverValuePtr(v *SubscriptionReceiverValueArgs) SubscriptionReceiverValuePtrInput {
	return (*subscriptionReceiverValuePtrType)(v)
}

func (*subscriptionReceiverValuePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionReceiverValue)(nil)).Elem()
}

func (i *subscriptionReceiverValuePtrType) ToSubscriptionReceiverValuePtrOutput() SubscriptionReceiverValuePtrOutput {
	return i.ToSubscriptionReceiverValuePtrOutputWithContext(context.Background())
}

func (i *subscriptionReceiverValuePtrType) ToSubscriptionReceiverValuePtrOutputWithContext(ctx context.Context) SubscriptionReceiverValuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionReceiverValuePtrOutput)
}

type SubscriptionReceiverValueOutput struct{ *pulumi.OutputState }

func (SubscriptionReceiverValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionReceiverValue)(nil)).Elem()
}

func (o SubscriptionReceiverValueOutput) ToSubscriptionReceiverValueOutput() SubscriptionReceiverValueOutput {
	return o
}

func (o SubscriptionReceiverValueOutput) ToSubscriptionReceiverValueOutputWithContext(ctx context.Context) SubscriptionReceiverValueOutput {
	return o
}

func (o SubscriptionReceiverValueOutput) ToSubscriptionReceiverValuePtrOutput() SubscriptionReceiverValuePtrOutput {
	return o.ToSubscriptionReceiverValuePtrOutputWithContext(context.Background())
}

func (o SubscriptionReceiverValueOutput) ToSubscriptionReceiverValuePtrOutputWithContext(ctx context.Context) SubscriptionReceiverValuePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubscriptionReceiverValue) *SubscriptionReceiverValue {
		return &v
	}).(SubscriptionReceiverValuePtrOutput)
}

func (o SubscriptionReceiverValueOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionReceiverValue) *string { return v.Role }).(pulumi.StringPtrOutput)
}

func (o SubscriptionReceiverValueOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionReceiverValue) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o SubscriptionReceiverValueOutput) SubscriptionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionReceiverValue) *string { return v.SubscriptionName }).(pulumi.StringPtrOutput)
}

type SubscriptionReceiverValuePtrOutput struct{ *pulumi.OutputState }

func (SubscriptionReceiverValuePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionReceiverValue)(nil)).Elem()
}

func (o SubscriptionReceiverValuePtrOutput) ToSubscriptionReceiverValuePtrOutput() SubscriptionReceiverValuePtrOutput {
	return o
}

func (o SubscriptionReceiverValuePtrOutput) ToSubscriptionReceiverValuePtrOutputWithContext(ctx context.Context) SubscriptionReceiverValuePtrOutput {
	return o
}

func (o SubscriptionReceiverValuePtrOutput) Elem() SubscriptionReceiverValueOutput {
	return o.ApplyT(func(v *SubscriptionReceiverValue) SubscriptionReceiverValue {
		if v != nil {
			return *v
		}
		var ret SubscriptionReceiverValue
		return ret
	}).(SubscriptionReceiverValueOutput)
}

func (o SubscriptionReceiverValuePtrOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionReceiverValue) *string {
		if v == nil {
			return nil
		}
		return v.Role
	}).(pulumi.StringPtrOutput)
}

func (o SubscriptionReceiverValuePtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionReceiverValue) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o SubscriptionReceiverValuePtrOutput) SubscriptionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionReceiverValue) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionName
	}).(pulumi.StringPtrOutput)
}

type SubscriptionReceiverValueResponse struct {
	Role             *string `pulumi:"role"`
	SubscriptionId   *string `pulumi:"subscriptionId"`
	SubscriptionName *string `pulumi:"subscriptionName"`
}





type SubscriptionReceiverValueResponseInput interface {
	pulumi.Input

	ToSubscriptionReceiverValueResponseOutput() SubscriptionReceiverValueResponseOutput
	ToSubscriptionReceiverValueResponseOutputWithContext(context.Context) SubscriptionReceiverValueResponseOutput
}

type SubscriptionReceiverValueResponseArgs struct {
	Role             pulumi.StringPtrInput `pulumi:"role"`
	SubscriptionId   pulumi.StringPtrInput `pulumi:"subscriptionId"`
	SubscriptionName pulumi.StringPtrInput `pulumi:"subscriptionName"`
}

func (SubscriptionReceiverValueResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionReceiverValueResponse)(nil)).Elem()
}

func (i SubscriptionReceiverValueResponseArgs) ToSubscriptionReceiverValueResponseOutput() SubscriptionReceiverValueResponseOutput {
	return i.ToSubscriptionReceiverValueResponseOutputWithContext(context.Background())
}

func (i SubscriptionReceiverValueResponseArgs) ToSubscriptionReceiverValueResponseOutputWithContext(ctx context.Context) SubscriptionReceiverValueResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionReceiverValueResponseOutput)
}

func (i SubscriptionReceiverValueResponseArgs) ToSubscriptionReceiverValueResponsePtrOutput() SubscriptionReceiverValueResponsePtrOutput {
	return i.ToSubscriptionReceiverValueResponsePtrOutputWithContext(context.Background())
}

func (i SubscriptionReceiverValueResponseArgs) ToSubscriptionReceiverValueResponsePtrOutputWithContext(ctx context.Context) SubscriptionReceiverValueResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionReceiverValueResponseOutput).ToSubscriptionReceiverValueResponsePtrOutputWithContext(ctx)
}









type SubscriptionReceiverValueResponsePtrInput interface {
	pulumi.Input

	ToSubscriptionReceiverValueResponsePtrOutput() SubscriptionReceiverValueResponsePtrOutput
	ToSubscriptionReceiverValueResponsePtrOutputWithContext(context.Context) SubscriptionReceiverValueResponsePtrOutput
}

type subscriptionReceiverValueResponsePtrType SubscriptionReceiverValueResponseArgs

func SubscriptionReceiverValueResponsePtr(v *SubscriptionReceiverValueResponseArgs) SubscriptionReceiverValueResponsePtrInput {
	return (*subscriptionReceiverValueResponsePtrType)(v)
}

func (*subscriptionReceiverValueResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionReceiverValueResponse)(nil)).Elem()
}

func (i *subscriptionReceiverValueResponsePtrType) ToSubscriptionReceiverValueResponsePtrOutput() SubscriptionReceiverValueResponsePtrOutput {
	return i.ToSubscriptionReceiverValueResponsePtrOutputWithContext(context.Background())
}

func (i *subscriptionReceiverValueResponsePtrType) ToSubscriptionReceiverValueResponsePtrOutputWithContext(ctx context.Context) SubscriptionReceiverValueResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionReceiverValueResponsePtrOutput)
}

type SubscriptionReceiverValueResponseOutput struct{ *pulumi.OutputState }

func (SubscriptionReceiverValueResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionReceiverValueResponse)(nil)).Elem()
}

func (o SubscriptionReceiverValueResponseOutput) ToSubscriptionReceiverValueResponseOutput() SubscriptionReceiverValueResponseOutput {
	return o
}

func (o SubscriptionReceiverValueResponseOutput) ToSubscriptionReceiverValueResponseOutputWithContext(ctx context.Context) SubscriptionReceiverValueResponseOutput {
	return o
}

func (o SubscriptionReceiverValueResponseOutput) ToSubscriptionReceiverValueResponsePtrOutput() SubscriptionReceiverValueResponsePtrOutput {
	return o.ToSubscriptionReceiverValueResponsePtrOutputWithContext(context.Background())
}

func (o SubscriptionReceiverValueResponseOutput) ToSubscriptionReceiverValueResponsePtrOutputWithContext(ctx context.Context) SubscriptionReceiverValueResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubscriptionReceiverValueResponse) *SubscriptionReceiverValueResponse {
		return &v
	}).(SubscriptionReceiverValueResponsePtrOutput)
}

func (o SubscriptionReceiverValueResponseOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionReceiverValueResponse) *string { return v.Role }).(pulumi.StringPtrOutput)
}

func (o SubscriptionReceiverValueResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionReceiverValueResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o SubscriptionReceiverValueResponseOutput) SubscriptionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionReceiverValueResponse) *string { return v.SubscriptionName }).(pulumi.StringPtrOutput)
}

type SubscriptionReceiverValueResponsePtrOutput struct{ *pulumi.OutputState }

func (SubscriptionReceiverValueResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionReceiverValueResponse)(nil)).Elem()
}

func (o SubscriptionReceiverValueResponsePtrOutput) ToSubscriptionReceiverValueResponsePtrOutput() SubscriptionReceiverValueResponsePtrOutput {
	return o
}

func (o SubscriptionReceiverValueResponsePtrOutput) ToSubscriptionReceiverValueResponsePtrOutputWithContext(ctx context.Context) SubscriptionReceiverValueResponsePtrOutput {
	return o
}

func (o SubscriptionReceiverValueResponsePtrOutput) Elem() SubscriptionReceiverValueResponseOutput {
	return o.ApplyT(func(v *SubscriptionReceiverValueResponse) SubscriptionReceiverValueResponse {
		if v != nil {
			return *v
		}
		var ret SubscriptionReceiverValueResponse
		return ret
	}).(SubscriptionReceiverValueResponseOutput)
}

func (o SubscriptionReceiverValueResponsePtrOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionReceiverValueResponse) *string {
		if v == nil {
			return nil
		}
		return v.Role
	}).(pulumi.StringPtrOutput)
}

func (o SubscriptionReceiverValueResponsePtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionReceiverValueResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o SubscriptionReceiverValueResponsePtrOutput) SubscriptionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionReceiverValueResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionName
	}).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
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

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
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

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

type TargetOSInfo struct {
	OsUpdateType string   `pulumi:"osUpdateType"`
	TargetOSs    []string `pulumi:"targetOSs"`
}





type TargetOSInfoInput interface {
	pulumi.Input

	ToTargetOSInfoOutput() TargetOSInfoOutput
	ToTargetOSInfoOutputWithContext(context.Context) TargetOSInfoOutput
}

type TargetOSInfoArgs struct {
	OsUpdateType pulumi.StringInput      `pulumi:"osUpdateType"`
	TargetOSs    pulumi.StringArrayInput `pulumi:"targetOSs"`
}

func (TargetOSInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetOSInfo)(nil)).Elem()
}

func (i TargetOSInfoArgs) ToTargetOSInfoOutput() TargetOSInfoOutput {
	return i.ToTargetOSInfoOutputWithContext(context.Background())
}

func (i TargetOSInfoArgs) ToTargetOSInfoOutputWithContext(ctx context.Context) TargetOSInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetOSInfoOutput)
}





type TargetOSInfoArrayInput interface {
	pulumi.Input

	ToTargetOSInfoArrayOutput() TargetOSInfoArrayOutput
	ToTargetOSInfoArrayOutputWithContext(context.Context) TargetOSInfoArrayOutput
}

type TargetOSInfoArray []TargetOSInfoInput

func (TargetOSInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetOSInfo)(nil)).Elem()
}

func (i TargetOSInfoArray) ToTargetOSInfoArrayOutput() TargetOSInfoArrayOutput {
	return i.ToTargetOSInfoArrayOutputWithContext(context.Background())
}

func (i TargetOSInfoArray) ToTargetOSInfoArrayOutputWithContext(ctx context.Context) TargetOSInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetOSInfoArrayOutput)
}

type TargetOSInfoOutput struct{ *pulumi.OutputState }

func (TargetOSInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetOSInfo)(nil)).Elem()
}

func (o TargetOSInfoOutput) ToTargetOSInfoOutput() TargetOSInfoOutput {
	return o
}

func (o TargetOSInfoOutput) ToTargetOSInfoOutputWithContext(ctx context.Context) TargetOSInfoOutput {
	return o
}

func (o TargetOSInfoOutput) OsUpdateType() pulumi.StringOutput {
	return o.ApplyT(func(v TargetOSInfo) string { return v.OsUpdateType }).(pulumi.StringOutput)
}

func (o TargetOSInfoOutput) TargetOSs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TargetOSInfo) []string { return v.TargetOSs }).(pulumi.StringArrayOutput)
}

type TargetOSInfoArrayOutput struct{ *pulumi.OutputState }

func (TargetOSInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetOSInfo)(nil)).Elem()
}

func (o TargetOSInfoArrayOutput) ToTargetOSInfoArrayOutput() TargetOSInfoArrayOutput {
	return o
}

func (o TargetOSInfoArrayOutput) ToTargetOSInfoArrayOutputWithContext(ctx context.Context) TargetOSInfoArrayOutput {
	return o
}

func (o TargetOSInfoArrayOutput) Index(i pulumi.IntInput) TargetOSInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetOSInfo {
		return vs[0].([]TargetOSInfo)[vs[1].(int)]
	}).(TargetOSInfoOutput)
}

type TargetOSInfoResponse struct {
	OsUpdateType string   `pulumi:"osUpdateType"`
	TargetOSs    []string `pulumi:"targetOSs"`
}





type TargetOSInfoResponseInput interface {
	pulumi.Input

	ToTargetOSInfoResponseOutput() TargetOSInfoResponseOutput
	ToTargetOSInfoResponseOutputWithContext(context.Context) TargetOSInfoResponseOutput
}

type TargetOSInfoResponseArgs struct {
	OsUpdateType pulumi.StringInput      `pulumi:"osUpdateType"`
	TargetOSs    pulumi.StringArrayInput `pulumi:"targetOSs"`
}

func (TargetOSInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetOSInfoResponse)(nil)).Elem()
}

func (i TargetOSInfoResponseArgs) ToTargetOSInfoResponseOutput() TargetOSInfoResponseOutput {
	return i.ToTargetOSInfoResponseOutputWithContext(context.Background())
}

func (i TargetOSInfoResponseArgs) ToTargetOSInfoResponseOutputWithContext(ctx context.Context) TargetOSInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetOSInfoResponseOutput)
}





type TargetOSInfoResponseArrayInput interface {
	pulumi.Input

	ToTargetOSInfoResponseArrayOutput() TargetOSInfoResponseArrayOutput
	ToTargetOSInfoResponseArrayOutputWithContext(context.Context) TargetOSInfoResponseArrayOutput
}

type TargetOSInfoResponseArray []TargetOSInfoResponseInput

func (TargetOSInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetOSInfoResponse)(nil)).Elem()
}

func (i TargetOSInfoResponseArray) ToTargetOSInfoResponseArrayOutput() TargetOSInfoResponseArrayOutput {
	return i.ToTargetOSInfoResponseArrayOutputWithContext(context.Background())
}

func (i TargetOSInfoResponseArray) ToTargetOSInfoResponseArrayOutputWithContext(ctx context.Context) TargetOSInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetOSInfoResponseArrayOutput)
}

type TargetOSInfoResponseOutput struct{ *pulumi.OutputState }

func (TargetOSInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetOSInfoResponse)(nil)).Elem()
}

func (o TargetOSInfoResponseOutput) ToTargetOSInfoResponseOutput() TargetOSInfoResponseOutput {
	return o
}

func (o TargetOSInfoResponseOutput) ToTargetOSInfoResponseOutputWithContext(ctx context.Context) TargetOSInfoResponseOutput {
	return o
}

func (o TargetOSInfoResponseOutput) OsUpdateType() pulumi.StringOutput {
	return o.ApplyT(func(v TargetOSInfoResponse) string { return v.OsUpdateType }).(pulumi.StringOutput)
}

func (o TargetOSInfoResponseOutput) TargetOSs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TargetOSInfoResponse) []string { return v.TargetOSs }).(pulumi.StringArrayOutput)
}

type TargetOSInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (TargetOSInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetOSInfoResponse)(nil)).Elem()
}

func (o TargetOSInfoResponseArrayOutput) ToTargetOSInfoResponseArrayOutput() TargetOSInfoResponseArrayOutput {
	return o
}

func (o TargetOSInfoResponseArrayOutput) ToTargetOSInfoResponseArrayOutputWithContext(ctx context.Context) TargetOSInfoResponseArrayOutput {
	return o
}

func (o TargetOSInfoResponseArrayOutput) Index(i pulumi.IntInput) TargetOSInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetOSInfoResponse {
		return vs[0].([]TargetOSInfoResponse)[vs[1].(int)]
	}).(TargetOSInfoResponseOutput)
}

type Test struct {
	Commands []Command `pulumi:"commands"`
	IsActive *bool     `pulumi:"isActive"`
	TestType string    `pulumi:"testType"`
}





type TestInput interface {
	pulumi.Input

	ToTestOutput() TestOutput
	ToTestOutputWithContext(context.Context) TestOutput
}

type TestArgs struct {
	Commands CommandArrayInput   `pulumi:"commands"`
	IsActive pulumi.BoolPtrInput `pulumi:"isActive"`
	TestType pulumi.StringInput  `pulumi:"testType"`
}

func (TestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Test)(nil)).Elem()
}

func (i TestArgs) ToTestOutput() TestOutput {
	return i.ToTestOutputWithContext(context.Background())
}

func (i TestArgs) ToTestOutputWithContext(ctx context.Context) TestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TestOutput)
}





type TestArrayInput interface {
	pulumi.Input

	ToTestArrayOutput() TestArrayOutput
	ToTestArrayOutputWithContext(context.Context) TestArrayOutput
}

type TestArray []TestInput

func (TestArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Test)(nil)).Elem()
}

func (i TestArray) ToTestArrayOutput() TestArrayOutput {
	return i.ToTestArrayOutputWithContext(context.Background())
}

func (i TestArray) ToTestArrayOutputWithContext(ctx context.Context) TestArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TestArrayOutput)
}

type TestOutput struct{ *pulumi.OutputState }

func (TestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Test)(nil)).Elem()
}

func (o TestOutput) ToTestOutput() TestOutput {
	return o
}

func (o TestOutput) ToTestOutputWithContext(ctx context.Context) TestOutput {
	return o
}

func (o TestOutput) Commands() CommandArrayOutput {
	return o.ApplyT(func(v Test) []Command { return v.Commands }).(CommandArrayOutput)
}

func (o TestOutput) IsActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Test) *bool { return v.IsActive }).(pulumi.BoolPtrOutput)
}

func (o TestOutput) TestType() pulumi.StringOutput {
	return o.ApplyT(func(v Test) string { return v.TestType }).(pulumi.StringOutput)
}

type TestArrayOutput struct{ *pulumi.OutputState }

func (TestArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Test)(nil)).Elem()
}

func (o TestArrayOutput) ToTestArrayOutput() TestArrayOutput {
	return o
}

func (o TestArrayOutput) ToTestArrayOutputWithContext(ctx context.Context) TestArrayOutput {
	return o
}

func (o TestArrayOutput) Index(i pulumi.IntInput) TestOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Test {
		return vs[0].([]Test)[vs[1].(int)]
	}).(TestOutput)
}

type TestBaseAccountSKU struct {
	Locations    []string `pulumi:"locations"`
	Name         string   `pulumi:"name"`
	ResourceType *string  `pulumi:"resourceType"`
	Tier         string   `pulumi:"tier"`
}





type TestBaseAccountSKUInput interface {
	pulumi.Input

	ToTestBaseAccountSKUOutput() TestBaseAccountSKUOutput
	ToTestBaseAccountSKUOutputWithContext(context.Context) TestBaseAccountSKUOutput
}

type TestBaseAccountSKUArgs struct {
	Locations    pulumi.StringArrayInput `pulumi:"locations"`
	Name         pulumi.StringInput      `pulumi:"name"`
	ResourceType pulumi.StringPtrInput   `pulumi:"resourceType"`
	Tier         pulumi.StringInput      `pulumi:"tier"`
}

func (TestBaseAccountSKUArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TestBaseAccountSKU)(nil)).Elem()
}

func (i TestBaseAccountSKUArgs) ToTestBaseAccountSKUOutput() TestBaseAccountSKUOutput {
	return i.ToTestBaseAccountSKUOutputWithContext(context.Background())
}

func (i TestBaseAccountSKUArgs) ToTestBaseAccountSKUOutputWithContext(ctx context.Context) TestBaseAccountSKUOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TestBaseAccountSKUOutput)
}

func (i TestBaseAccountSKUArgs) ToTestBaseAccountSKUPtrOutput() TestBaseAccountSKUPtrOutput {
	return i.ToTestBaseAccountSKUPtrOutputWithContext(context.Background())
}

func (i TestBaseAccountSKUArgs) ToTestBaseAccountSKUPtrOutputWithContext(ctx context.Context) TestBaseAccountSKUPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TestBaseAccountSKUOutput).ToTestBaseAccountSKUPtrOutputWithContext(ctx)
}









type TestBaseAccountSKUPtrInput interface {
	pulumi.Input

	ToTestBaseAccountSKUPtrOutput() TestBaseAccountSKUPtrOutput
	ToTestBaseAccountSKUPtrOutputWithContext(context.Context) TestBaseAccountSKUPtrOutput
}

type testBaseAccountSKUPtrType TestBaseAccountSKUArgs

func TestBaseAccountSKUPtr(v *TestBaseAccountSKUArgs) TestBaseAccountSKUPtrInput {
	return (*testBaseAccountSKUPtrType)(v)
}

func (*testBaseAccountSKUPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TestBaseAccountSKU)(nil)).Elem()
}

func (i *testBaseAccountSKUPtrType) ToTestBaseAccountSKUPtrOutput() TestBaseAccountSKUPtrOutput {
	return i.ToTestBaseAccountSKUPtrOutputWithContext(context.Background())
}

func (i *testBaseAccountSKUPtrType) ToTestBaseAccountSKUPtrOutputWithContext(ctx context.Context) TestBaseAccountSKUPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TestBaseAccountSKUPtrOutput)
}

type TestBaseAccountSKUOutput struct{ *pulumi.OutputState }

func (TestBaseAccountSKUOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TestBaseAccountSKU)(nil)).Elem()
}

func (o TestBaseAccountSKUOutput) ToTestBaseAccountSKUOutput() TestBaseAccountSKUOutput {
	return o
}

func (o TestBaseAccountSKUOutput) ToTestBaseAccountSKUOutputWithContext(ctx context.Context) TestBaseAccountSKUOutput {
	return o
}

func (o TestBaseAccountSKUOutput) ToTestBaseAccountSKUPtrOutput() TestBaseAccountSKUPtrOutput {
	return o.ToTestBaseAccountSKUPtrOutputWithContext(context.Background())
}

func (o TestBaseAccountSKUOutput) ToTestBaseAccountSKUPtrOutputWithContext(ctx context.Context) TestBaseAccountSKUPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TestBaseAccountSKU) *TestBaseAccountSKU {
		return &v
	}).(TestBaseAccountSKUPtrOutput)
}

func (o TestBaseAccountSKUOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TestBaseAccountSKU) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o TestBaseAccountSKUOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TestBaseAccountSKU) string { return v.Name }).(pulumi.StringOutput)
}

func (o TestBaseAccountSKUOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TestBaseAccountSKU) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o TestBaseAccountSKUOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v TestBaseAccountSKU) string { return v.Tier }).(pulumi.StringOutput)
}

type TestBaseAccountSKUPtrOutput struct{ *pulumi.OutputState }

func (TestBaseAccountSKUPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TestBaseAccountSKU)(nil)).Elem()
}

func (o TestBaseAccountSKUPtrOutput) ToTestBaseAccountSKUPtrOutput() TestBaseAccountSKUPtrOutput {
	return o
}

func (o TestBaseAccountSKUPtrOutput) ToTestBaseAccountSKUPtrOutputWithContext(ctx context.Context) TestBaseAccountSKUPtrOutput {
	return o
}

func (o TestBaseAccountSKUPtrOutput) Elem() TestBaseAccountSKUOutput {
	return o.ApplyT(func(v *TestBaseAccountSKU) TestBaseAccountSKU {
		if v != nil {
			return *v
		}
		var ret TestBaseAccountSKU
		return ret
	}).(TestBaseAccountSKUOutput)
}

func (o TestBaseAccountSKUPtrOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TestBaseAccountSKU) []string {
		if v == nil {
			return nil
		}
		return v.Locations
	}).(pulumi.StringArrayOutput)
}

func (o TestBaseAccountSKUPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TestBaseAccountSKU) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o TestBaseAccountSKUPtrOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TestBaseAccountSKU) *string {
		if v == nil {
			return nil
		}
		return v.ResourceType
	}).(pulumi.StringPtrOutput)
}

func (o TestBaseAccountSKUPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TestBaseAccountSKU) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type TestBaseAccountSKUCapabilityResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type TestBaseAccountSKUCapabilityResponseInput interface {
	pulumi.Input

	ToTestBaseAccountSKUCapabilityResponseOutput() TestBaseAccountSKUCapabilityResponseOutput
	ToTestBaseAccountSKUCapabilityResponseOutputWithContext(context.Context) TestBaseAccountSKUCapabilityResponseOutput
}

type TestBaseAccountSKUCapabilityResponseArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (TestBaseAccountSKUCapabilityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TestBaseAccountSKUCapabilityResponse)(nil)).Elem()
}

func (i TestBaseAccountSKUCapabilityResponseArgs) ToTestBaseAccountSKUCapabilityResponseOutput() TestBaseAccountSKUCapabilityResponseOutput {
	return i.ToTestBaseAccountSKUCapabilityResponseOutputWithContext(context.Background())
}

func (i TestBaseAccountSKUCapabilityResponseArgs) ToTestBaseAccountSKUCapabilityResponseOutputWithContext(ctx context.Context) TestBaseAccountSKUCapabilityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TestBaseAccountSKUCapabilityResponseOutput)
}





type TestBaseAccountSKUCapabilityResponseArrayInput interface {
	pulumi.Input

	ToTestBaseAccountSKUCapabilityResponseArrayOutput() TestBaseAccountSKUCapabilityResponseArrayOutput
	ToTestBaseAccountSKUCapabilityResponseArrayOutputWithContext(context.Context) TestBaseAccountSKUCapabilityResponseArrayOutput
}

type TestBaseAccountSKUCapabilityResponseArray []TestBaseAccountSKUCapabilityResponseInput

func (TestBaseAccountSKUCapabilityResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TestBaseAccountSKUCapabilityResponse)(nil)).Elem()
}

func (i TestBaseAccountSKUCapabilityResponseArray) ToTestBaseAccountSKUCapabilityResponseArrayOutput() TestBaseAccountSKUCapabilityResponseArrayOutput {
	return i.ToTestBaseAccountSKUCapabilityResponseArrayOutputWithContext(context.Background())
}

func (i TestBaseAccountSKUCapabilityResponseArray) ToTestBaseAccountSKUCapabilityResponseArrayOutputWithContext(ctx context.Context) TestBaseAccountSKUCapabilityResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TestBaseAccountSKUCapabilityResponseArrayOutput)
}

type TestBaseAccountSKUCapabilityResponseOutput struct{ *pulumi.OutputState }

func (TestBaseAccountSKUCapabilityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TestBaseAccountSKUCapabilityResponse)(nil)).Elem()
}

func (o TestBaseAccountSKUCapabilityResponseOutput) ToTestBaseAccountSKUCapabilityResponseOutput() TestBaseAccountSKUCapabilityResponseOutput {
	return o
}

func (o TestBaseAccountSKUCapabilityResponseOutput) ToTestBaseAccountSKUCapabilityResponseOutputWithContext(ctx context.Context) TestBaseAccountSKUCapabilityResponseOutput {
	return o
}

func (o TestBaseAccountSKUCapabilityResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TestBaseAccountSKUCapabilityResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o TestBaseAccountSKUCapabilityResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TestBaseAccountSKUCapabilityResponse) string { return v.Value }).(pulumi.StringOutput)
}

type TestBaseAccountSKUCapabilityResponseArrayOutput struct{ *pulumi.OutputState }

func (TestBaseAccountSKUCapabilityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TestBaseAccountSKUCapabilityResponse)(nil)).Elem()
}

func (o TestBaseAccountSKUCapabilityResponseArrayOutput) ToTestBaseAccountSKUCapabilityResponseArrayOutput() TestBaseAccountSKUCapabilityResponseArrayOutput {
	return o
}

func (o TestBaseAccountSKUCapabilityResponseArrayOutput) ToTestBaseAccountSKUCapabilityResponseArrayOutputWithContext(ctx context.Context) TestBaseAccountSKUCapabilityResponseArrayOutput {
	return o
}

func (o TestBaseAccountSKUCapabilityResponseArrayOutput) Index(i pulumi.IntInput) TestBaseAccountSKUCapabilityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TestBaseAccountSKUCapabilityResponse {
		return vs[0].([]TestBaseAccountSKUCapabilityResponse)[vs[1].(int)]
	}).(TestBaseAccountSKUCapabilityResponseOutput)
}

type TestBaseAccountSKUResponse struct {
	Capabilities []TestBaseAccountSKUCapabilityResponse `pulumi:"capabilities"`
	Locations    []string                               `pulumi:"locations"`
	Name         string                                 `pulumi:"name"`
	ResourceType *string                                `pulumi:"resourceType"`
	Tier         string                                 `pulumi:"tier"`
}





type TestBaseAccountSKUResponseInput interface {
	pulumi.Input

	ToTestBaseAccountSKUResponseOutput() TestBaseAccountSKUResponseOutput
	ToTestBaseAccountSKUResponseOutputWithContext(context.Context) TestBaseAccountSKUResponseOutput
}

type TestBaseAccountSKUResponseArgs struct {
	Capabilities TestBaseAccountSKUCapabilityResponseArrayInput `pulumi:"capabilities"`
	Locations    pulumi.StringArrayInput                        `pulumi:"locations"`
	Name         pulumi.StringInput                             `pulumi:"name"`
	ResourceType pulumi.StringPtrInput                          `pulumi:"resourceType"`
	Tier         pulumi.StringInput                             `pulumi:"tier"`
}

func (TestBaseAccountSKUResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TestBaseAccountSKUResponse)(nil)).Elem()
}

func (i TestBaseAccountSKUResponseArgs) ToTestBaseAccountSKUResponseOutput() TestBaseAccountSKUResponseOutput {
	return i.ToTestBaseAccountSKUResponseOutputWithContext(context.Background())
}

func (i TestBaseAccountSKUResponseArgs) ToTestBaseAccountSKUResponseOutputWithContext(ctx context.Context) TestBaseAccountSKUResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TestBaseAccountSKUResponseOutput)
}

func (i TestBaseAccountSKUResponseArgs) ToTestBaseAccountSKUResponsePtrOutput() TestBaseAccountSKUResponsePtrOutput {
	return i.ToTestBaseAccountSKUResponsePtrOutputWithContext(context.Background())
}

func (i TestBaseAccountSKUResponseArgs) ToTestBaseAccountSKUResponsePtrOutputWithContext(ctx context.Context) TestBaseAccountSKUResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TestBaseAccountSKUResponseOutput).ToTestBaseAccountSKUResponsePtrOutputWithContext(ctx)
}









type TestBaseAccountSKUResponsePtrInput interface {
	pulumi.Input

	ToTestBaseAccountSKUResponsePtrOutput() TestBaseAccountSKUResponsePtrOutput
	ToTestBaseAccountSKUResponsePtrOutputWithContext(context.Context) TestBaseAccountSKUResponsePtrOutput
}

type testBaseAccountSKUResponsePtrType TestBaseAccountSKUResponseArgs

func TestBaseAccountSKUResponsePtr(v *TestBaseAccountSKUResponseArgs) TestBaseAccountSKUResponsePtrInput {
	return (*testBaseAccountSKUResponsePtrType)(v)
}

func (*testBaseAccountSKUResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TestBaseAccountSKUResponse)(nil)).Elem()
}

func (i *testBaseAccountSKUResponsePtrType) ToTestBaseAccountSKUResponsePtrOutput() TestBaseAccountSKUResponsePtrOutput {
	return i.ToTestBaseAccountSKUResponsePtrOutputWithContext(context.Background())
}

func (i *testBaseAccountSKUResponsePtrType) ToTestBaseAccountSKUResponsePtrOutputWithContext(ctx context.Context) TestBaseAccountSKUResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TestBaseAccountSKUResponsePtrOutput)
}

type TestBaseAccountSKUResponseOutput struct{ *pulumi.OutputState }

func (TestBaseAccountSKUResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TestBaseAccountSKUResponse)(nil)).Elem()
}

func (o TestBaseAccountSKUResponseOutput) ToTestBaseAccountSKUResponseOutput() TestBaseAccountSKUResponseOutput {
	return o
}

func (o TestBaseAccountSKUResponseOutput) ToTestBaseAccountSKUResponseOutputWithContext(ctx context.Context) TestBaseAccountSKUResponseOutput {
	return o
}

func (o TestBaseAccountSKUResponseOutput) ToTestBaseAccountSKUResponsePtrOutput() TestBaseAccountSKUResponsePtrOutput {
	return o.ToTestBaseAccountSKUResponsePtrOutputWithContext(context.Background())
}

func (o TestBaseAccountSKUResponseOutput) ToTestBaseAccountSKUResponsePtrOutputWithContext(ctx context.Context) TestBaseAccountSKUResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TestBaseAccountSKUResponse) *TestBaseAccountSKUResponse {
		return &v
	}).(TestBaseAccountSKUResponsePtrOutput)
}

func (o TestBaseAccountSKUResponseOutput) Capabilities() TestBaseAccountSKUCapabilityResponseArrayOutput {
	return o.ApplyT(func(v TestBaseAccountSKUResponse) []TestBaseAccountSKUCapabilityResponse { return v.Capabilities }).(TestBaseAccountSKUCapabilityResponseArrayOutput)
}

func (o TestBaseAccountSKUResponseOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TestBaseAccountSKUResponse) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o TestBaseAccountSKUResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TestBaseAccountSKUResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o TestBaseAccountSKUResponseOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TestBaseAccountSKUResponse) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o TestBaseAccountSKUResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v TestBaseAccountSKUResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type TestBaseAccountSKUResponsePtrOutput struct{ *pulumi.OutputState }

func (TestBaseAccountSKUResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TestBaseAccountSKUResponse)(nil)).Elem()
}

func (o TestBaseAccountSKUResponsePtrOutput) ToTestBaseAccountSKUResponsePtrOutput() TestBaseAccountSKUResponsePtrOutput {
	return o
}

func (o TestBaseAccountSKUResponsePtrOutput) ToTestBaseAccountSKUResponsePtrOutputWithContext(ctx context.Context) TestBaseAccountSKUResponsePtrOutput {
	return o
}

func (o TestBaseAccountSKUResponsePtrOutput) Elem() TestBaseAccountSKUResponseOutput {
	return o.ApplyT(func(v *TestBaseAccountSKUResponse) TestBaseAccountSKUResponse {
		if v != nil {
			return *v
		}
		var ret TestBaseAccountSKUResponse
		return ret
	}).(TestBaseAccountSKUResponseOutput)
}

func (o TestBaseAccountSKUResponsePtrOutput) Capabilities() TestBaseAccountSKUCapabilityResponseArrayOutput {
	return o.ApplyT(func(v *TestBaseAccountSKUResponse) []TestBaseAccountSKUCapabilityResponse {
		if v == nil {
			return nil
		}
		return v.Capabilities
	}).(TestBaseAccountSKUCapabilityResponseArrayOutput)
}

func (o TestBaseAccountSKUResponsePtrOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TestBaseAccountSKUResponse) []string {
		if v == nil {
			return nil
		}
		return v.Locations
	}).(pulumi.StringArrayOutput)
}

func (o TestBaseAccountSKUResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TestBaseAccountSKUResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o TestBaseAccountSKUResponsePtrOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TestBaseAccountSKUResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceType
	}).(pulumi.StringPtrOutput)
}

func (o TestBaseAccountSKUResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TestBaseAccountSKUResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type TestResponse struct {
	Commands            []CommandResponse `pulumi:"commands"`
	IsActive            *bool             `pulumi:"isActive"`
	TestType            string            `pulumi:"testType"`
	ValidationRunStatus string            `pulumi:"validationRunStatus"`
}





type TestResponseInput interface {
	pulumi.Input

	ToTestResponseOutput() TestResponseOutput
	ToTestResponseOutputWithContext(context.Context) TestResponseOutput
}

type TestResponseArgs struct {
	Commands            CommandResponseArrayInput `pulumi:"commands"`
	IsActive            pulumi.BoolPtrInput       `pulumi:"isActive"`
	TestType            pulumi.StringInput        `pulumi:"testType"`
	ValidationRunStatus pulumi.StringInput        `pulumi:"validationRunStatus"`
}

func (TestResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TestResponse)(nil)).Elem()
}

func (i TestResponseArgs) ToTestResponseOutput() TestResponseOutput {
	return i.ToTestResponseOutputWithContext(context.Background())
}

func (i TestResponseArgs) ToTestResponseOutputWithContext(ctx context.Context) TestResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TestResponseOutput)
}





type TestResponseArrayInput interface {
	pulumi.Input

	ToTestResponseArrayOutput() TestResponseArrayOutput
	ToTestResponseArrayOutputWithContext(context.Context) TestResponseArrayOutput
}

type TestResponseArray []TestResponseInput

func (TestResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TestResponse)(nil)).Elem()
}

func (i TestResponseArray) ToTestResponseArrayOutput() TestResponseArrayOutput {
	return i.ToTestResponseArrayOutputWithContext(context.Background())
}

func (i TestResponseArray) ToTestResponseArrayOutputWithContext(ctx context.Context) TestResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TestResponseArrayOutput)
}

type TestResponseOutput struct{ *pulumi.OutputState }

func (TestResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TestResponse)(nil)).Elem()
}

func (o TestResponseOutput) ToTestResponseOutput() TestResponseOutput {
	return o
}

func (o TestResponseOutput) ToTestResponseOutputWithContext(ctx context.Context) TestResponseOutput {
	return o
}

func (o TestResponseOutput) Commands() CommandResponseArrayOutput {
	return o.ApplyT(func(v TestResponse) []CommandResponse { return v.Commands }).(CommandResponseArrayOutput)
}

func (o TestResponseOutput) IsActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TestResponse) *bool { return v.IsActive }).(pulumi.BoolPtrOutput)
}

func (o TestResponseOutput) TestType() pulumi.StringOutput {
	return o.ApplyT(func(v TestResponse) string { return v.TestType }).(pulumi.StringOutput)
}

func (o TestResponseOutput) ValidationRunStatus() pulumi.StringOutput {
	return o.ApplyT(func(v TestResponse) string { return v.ValidationRunStatus }).(pulumi.StringOutput)
}

type TestResponseArrayOutput struct{ *pulumi.OutputState }

func (TestResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TestResponse)(nil)).Elem()
}

func (o TestResponseArrayOutput) ToTestResponseArrayOutput() TestResponseArrayOutput {
	return o
}

func (o TestResponseArrayOutput) ToTestResponseArrayOutputWithContext(ctx context.Context) TestResponseArrayOutput {
	return o
}

func (o TestResponseArrayOutput) Index(i pulumi.IntInput) TestResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TestResponse {
		return vs[0].([]TestResponse)[vs[1].(int)]
	}).(TestResponseOutput)
}

type UserObjectReceiverValue struct {
	UserObjectIds []string `pulumi:"userObjectIds"`
}





type UserObjectReceiverValueInput interface {
	pulumi.Input

	ToUserObjectReceiverValueOutput() UserObjectReceiverValueOutput
	ToUserObjectReceiverValueOutputWithContext(context.Context) UserObjectReceiverValueOutput
}

type UserObjectReceiverValueArgs struct {
	UserObjectIds pulumi.StringArrayInput `pulumi:"userObjectIds"`
}

func (UserObjectReceiverValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserObjectReceiverValue)(nil)).Elem()
}

func (i UserObjectReceiverValueArgs) ToUserObjectReceiverValueOutput() UserObjectReceiverValueOutput {
	return i.ToUserObjectReceiverValueOutputWithContext(context.Background())
}

func (i UserObjectReceiverValueArgs) ToUserObjectReceiverValueOutputWithContext(ctx context.Context) UserObjectReceiverValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserObjectReceiverValueOutput)
}

func (i UserObjectReceiverValueArgs) ToUserObjectReceiverValuePtrOutput() UserObjectReceiverValuePtrOutput {
	return i.ToUserObjectReceiverValuePtrOutputWithContext(context.Background())
}

func (i UserObjectReceiverValueArgs) ToUserObjectReceiverValuePtrOutputWithContext(ctx context.Context) UserObjectReceiverValuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserObjectReceiverValueOutput).ToUserObjectReceiverValuePtrOutputWithContext(ctx)
}









type UserObjectReceiverValuePtrInput interface {
	pulumi.Input

	ToUserObjectReceiverValuePtrOutput() UserObjectReceiverValuePtrOutput
	ToUserObjectReceiverValuePtrOutputWithContext(context.Context) UserObjectReceiverValuePtrOutput
}

type userObjectReceiverValuePtrType UserObjectReceiverValueArgs

func UserObjectReceiverValuePtr(v *UserObjectReceiverValueArgs) UserObjectReceiverValuePtrInput {
	return (*userObjectReceiverValuePtrType)(v)
}

func (*userObjectReceiverValuePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserObjectReceiverValue)(nil)).Elem()
}

func (i *userObjectReceiverValuePtrType) ToUserObjectReceiverValuePtrOutput() UserObjectReceiverValuePtrOutput {
	return i.ToUserObjectReceiverValuePtrOutputWithContext(context.Background())
}

func (i *userObjectReceiverValuePtrType) ToUserObjectReceiverValuePtrOutputWithContext(ctx context.Context) UserObjectReceiverValuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserObjectReceiverValuePtrOutput)
}

type UserObjectReceiverValueOutput struct{ *pulumi.OutputState }

func (UserObjectReceiverValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserObjectReceiverValue)(nil)).Elem()
}

func (o UserObjectReceiverValueOutput) ToUserObjectReceiverValueOutput() UserObjectReceiverValueOutput {
	return o
}

func (o UserObjectReceiverValueOutput) ToUserObjectReceiverValueOutputWithContext(ctx context.Context) UserObjectReceiverValueOutput {
	return o
}

func (o UserObjectReceiverValueOutput) ToUserObjectReceiverValuePtrOutput() UserObjectReceiverValuePtrOutput {
	return o.ToUserObjectReceiverValuePtrOutputWithContext(context.Background())
}

func (o UserObjectReceiverValueOutput) ToUserObjectReceiverValuePtrOutputWithContext(ctx context.Context) UserObjectReceiverValuePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserObjectReceiverValue) *UserObjectReceiverValue {
		return &v
	}).(UserObjectReceiverValuePtrOutput)
}

func (o UserObjectReceiverValueOutput) UserObjectIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UserObjectReceiverValue) []string { return v.UserObjectIds }).(pulumi.StringArrayOutput)
}

type UserObjectReceiverValuePtrOutput struct{ *pulumi.OutputState }

func (UserObjectReceiverValuePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserObjectReceiverValue)(nil)).Elem()
}

func (o UserObjectReceiverValuePtrOutput) ToUserObjectReceiverValuePtrOutput() UserObjectReceiverValuePtrOutput {
	return o
}

func (o UserObjectReceiverValuePtrOutput) ToUserObjectReceiverValuePtrOutputWithContext(ctx context.Context) UserObjectReceiverValuePtrOutput {
	return o
}

func (o UserObjectReceiverValuePtrOutput) Elem() UserObjectReceiverValueOutput {
	return o.ApplyT(func(v *UserObjectReceiverValue) UserObjectReceiverValue {
		if v != nil {
			return *v
		}
		var ret UserObjectReceiverValue
		return ret
	}).(UserObjectReceiverValueOutput)
}

func (o UserObjectReceiverValuePtrOutput) UserObjectIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UserObjectReceiverValue) []string {
		if v == nil {
			return nil
		}
		return v.UserObjectIds
	}).(pulumi.StringArrayOutput)
}

type UserObjectReceiverValueResponse struct {
	UserObjectIds []string `pulumi:"userObjectIds"`
}





type UserObjectReceiverValueResponseInput interface {
	pulumi.Input

	ToUserObjectReceiverValueResponseOutput() UserObjectReceiverValueResponseOutput
	ToUserObjectReceiverValueResponseOutputWithContext(context.Context) UserObjectReceiverValueResponseOutput
}

type UserObjectReceiverValueResponseArgs struct {
	UserObjectIds pulumi.StringArrayInput `pulumi:"userObjectIds"`
}

func (UserObjectReceiverValueResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserObjectReceiverValueResponse)(nil)).Elem()
}

func (i UserObjectReceiverValueResponseArgs) ToUserObjectReceiverValueResponseOutput() UserObjectReceiverValueResponseOutput {
	return i.ToUserObjectReceiverValueResponseOutputWithContext(context.Background())
}

func (i UserObjectReceiverValueResponseArgs) ToUserObjectReceiverValueResponseOutputWithContext(ctx context.Context) UserObjectReceiverValueResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserObjectReceiverValueResponseOutput)
}

func (i UserObjectReceiverValueResponseArgs) ToUserObjectReceiverValueResponsePtrOutput() UserObjectReceiverValueResponsePtrOutput {
	return i.ToUserObjectReceiverValueResponsePtrOutputWithContext(context.Background())
}

func (i UserObjectReceiverValueResponseArgs) ToUserObjectReceiverValueResponsePtrOutputWithContext(ctx context.Context) UserObjectReceiverValueResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserObjectReceiverValueResponseOutput).ToUserObjectReceiverValueResponsePtrOutputWithContext(ctx)
}









type UserObjectReceiverValueResponsePtrInput interface {
	pulumi.Input

	ToUserObjectReceiverValueResponsePtrOutput() UserObjectReceiverValueResponsePtrOutput
	ToUserObjectReceiverValueResponsePtrOutputWithContext(context.Context) UserObjectReceiverValueResponsePtrOutput
}

type userObjectReceiverValueResponsePtrType UserObjectReceiverValueResponseArgs

func UserObjectReceiverValueResponsePtr(v *UserObjectReceiverValueResponseArgs) UserObjectReceiverValueResponsePtrInput {
	return (*userObjectReceiverValueResponsePtrType)(v)
}

func (*userObjectReceiverValueResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserObjectReceiverValueResponse)(nil)).Elem()
}

func (i *userObjectReceiverValueResponsePtrType) ToUserObjectReceiverValueResponsePtrOutput() UserObjectReceiverValueResponsePtrOutput {
	return i.ToUserObjectReceiverValueResponsePtrOutputWithContext(context.Background())
}

func (i *userObjectReceiverValueResponsePtrType) ToUserObjectReceiverValueResponsePtrOutputWithContext(ctx context.Context) UserObjectReceiverValueResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserObjectReceiverValueResponsePtrOutput)
}

type UserObjectReceiverValueResponseOutput struct{ *pulumi.OutputState }

func (UserObjectReceiverValueResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserObjectReceiverValueResponse)(nil)).Elem()
}

func (o UserObjectReceiverValueResponseOutput) ToUserObjectReceiverValueResponseOutput() UserObjectReceiverValueResponseOutput {
	return o
}

func (o UserObjectReceiverValueResponseOutput) ToUserObjectReceiverValueResponseOutputWithContext(ctx context.Context) UserObjectReceiverValueResponseOutput {
	return o
}

func (o UserObjectReceiverValueResponseOutput) ToUserObjectReceiverValueResponsePtrOutput() UserObjectReceiverValueResponsePtrOutput {
	return o.ToUserObjectReceiverValueResponsePtrOutputWithContext(context.Background())
}

func (o UserObjectReceiverValueResponseOutput) ToUserObjectReceiverValueResponsePtrOutputWithContext(ctx context.Context) UserObjectReceiverValueResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserObjectReceiverValueResponse) *UserObjectReceiverValueResponse {
		return &v
	}).(UserObjectReceiverValueResponsePtrOutput)
}

func (o UserObjectReceiverValueResponseOutput) UserObjectIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UserObjectReceiverValueResponse) []string { return v.UserObjectIds }).(pulumi.StringArrayOutput)
}

type UserObjectReceiverValueResponsePtrOutput struct{ *pulumi.OutputState }

func (UserObjectReceiverValueResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserObjectReceiverValueResponse)(nil)).Elem()
}

func (o UserObjectReceiverValueResponsePtrOutput) ToUserObjectReceiverValueResponsePtrOutput() UserObjectReceiverValueResponsePtrOutput {
	return o
}

func (o UserObjectReceiverValueResponsePtrOutput) ToUserObjectReceiverValueResponsePtrOutputWithContext(ctx context.Context) UserObjectReceiverValueResponsePtrOutput {
	return o
}

func (o UserObjectReceiverValueResponsePtrOutput) Elem() UserObjectReceiverValueResponseOutput {
	return o.ApplyT(func(v *UserObjectReceiverValueResponse) UserObjectReceiverValueResponse {
		if v != nil {
			return *v
		}
		var ret UserObjectReceiverValueResponse
		return ret
	}).(UserObjectReceiverValueResponseOutput)
}

func (o UserObjectReceiverValueResponsePtrOutput) UserObjectIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UserObjectReceiverValueResponse) []string {
		if v == nil {
			return nil
		}
		return v.UserObjectIds
	}).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(CommandOutput{})
	pulumi.RegisterOutputType(CommandArrayOutput{})
	pulumi.RegisterOutputType(CommandResponseOutput{})
	pulumi.RegisterOutputType(CommandResponseArrayOutput{})
	pulumi.RegisterOutputType(DistributionGroupListReceiverValueOutput{})
	pulumi.RegisterOutputType(DistributionGroupListReceiverValuePtrOutput{})
	pulumi.RegisterOutputType(DistributionGroupListReceiverValueResponseOutput{})
	pulumi.RegisterOutputType(DistributionGroupListReceiverValueResponsePtrOutput{})
	pulumi.RegisterOutputType(NotificationEventReceiverOutput{})
	pulumi.RegisterOutputType(NotificationEventReceiverArrayOutput{})
	pulumi.RegisterOutputType(NotificationEventReceiverResponseOutput{})
	pulumi.RegisterOutputType(NotificationEventReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(NotificationReceiverValueOutput{})
	pulumi.RegisterOutputType(NotificationReceiverValuePtrOutput{})
	pulumi.RegisterOutputType(NotificationReceiverValueResponseOutput{})
	pulumi.RegisterOutputType(NotificationReceiverValueResponsePtrOutput{})
	pulumi.RegisterOutputType(PackageValidationResultResponseOutput{})
	pulumi.RegisterOutputType(PackageValidationResultResponseArrayOutput{})
	pulumi.RegisterOutputType(SubscriptionReceiverValueOutput{})
	pulumi.RegisterOutputType(SubscriptionReceiverValuePtrOutput{})
	pulumi.RegisterOutputType(SubscriptionReceiverValueResponseOutput{})
	pulumi.RegisterOutputType(SubscriptionReceiverValueResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(TargetOSInfoOutput{})
	pulumi.RegisterOutputType(TargetOSInfoArrayOutput{})
	pulumi.RegisterOutputType(TargetOSInfoResponseOutput{})
	pulumi.RegisterOutputType(TargetOSInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(TestOutput{})
	pulumi.RegisterOutputType(TestArrayOutput{})
	pulumi.RegisterOutputType(TestBaseAccountSKUOutput{})
	pulumi.RegisterOutputType(TestBaseAccountSKUPtrOutput{})
	pulumi.RegisterOutputType(TestBaseAccountSKUCapabilityResponseOutput{})
	pulumi.RegisterOutputType(TestBaseAccountSKUCapabilityResponseArrayOutput{})
	pulumi.RegisterOutputType(TestBaseAccountSKUResponseOutput{})
	pulumi.RegisterOutputType(TestBaseAccountSKUResponsePtrOutput{})
	pulumi.RegisterOutputType(TestResponseOutput{})
	pulumi.RegisterOutputType(TestResponseArrayOutput{})
	pulumi.RegisterOutputType(UserObjectReceiverValueOutput{})
	pulumi.RegisterOutputType(UserObjectReceiverValuePtrOutput{})
	pulumi.RegisterOutputType(UserObjectReceiverValueResponseOutput{})
	pulumi.RegisterOutputType(UserObjectReceiverValueResponsePtrOutput{})
}
